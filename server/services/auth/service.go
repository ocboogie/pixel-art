package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/pkg/argon2"
	"github.com/ocboogie/pixel-art/repositories"
	"github.com/ocboogie/pixel-art/services/avatar"
)

//go:generate mockgen -destination=../../mocks/service_auth.go -package mocks -mock_names Service=ServiceAuth github.com/ocboogie/pixel-art/services/auth Service

// The number of hex characters in a session ID
const SessionIDLength = 32

type Service interface {
	SignUp(user *models.UserNew) (*models.Session, error)
	Login(credentials *models.UserCredentials) (*models.Session, error)
	Logout(sessionID string) error
	CreateSession(userID string) (*models.Session, error)
	VerifySession(sessionID string) (string, error)
}

type service struct {
	userRepo      repositories.User
	sessionRepo   repositories.Session
	avatarService avatar.Service
	config        Config
}

func New(config Config, userRepo repositories.User, sessionRepo repositories.Session, avatarService avatar.Service) Service {
	return &service{
		userRepo:      userRepo,
		sessionRepo:   sessionRepo,
		avatarService: avatarService,
		config:        config,
	}
}

func (s *service) SignUp(user *models.UserNew) (*models.Session, error) {
	if valid := s.avatarService.Validate(user.Avatar); !valid {
		return nil, ErrInvalidAvatar
	}

	exists, err := s.userRepo.ExistsEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyInUse
	}

	idBytes, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	id := idBytes.String()

	hashedPassword, err := argon2.Hash(user.Password, s.config.HashConfig)
	if err != nil {
		return nil, err
	}

	userHashed := &models.User{
		ID:        id,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := s.userRepo.Save(userHashed); err != nil {
		return nil, err
	}

	return s.CreateSession(id)
}

func (s *service) Login(credentials *models.UserCredentials) (*models.Session, error) {
	user, err := s.userRepo.FindByEmail(credentials.Email)
	if err != nil {
		// TODO: Doc why this is here
		argon2.Hash(credentials.Password, s.config.HashConfig)
		if err == repositories.ErrUserNotFound {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	match, err := argon2.Verify(credentials.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, ErrInvalidCredentials
	}

	return s.CreateSession(user.ID)
}

func (s *service) Logout(sessionID string) error {
	return s.sessionRepo.Delete(sessionID)
}

func (s *service) hashSessionID(id string) string {
	hasher := sha256.New()
	hasher.Write([]byte(id))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *service) CreateSession(userID string) (*models.Session, error) {
	// Dividing by 2 because for each byte we will have two hex characters
	b := make([]byte, SessionIDLength/2)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	id := hex.EncodeToString(b)

	hashedID := s.hashSessionID(id)
	expiresAt := time.Now().Add(time.Duration(s.config.SessionLifetime) * time.Second)

	session := &models.Session{
		ID:        hashedID,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}

	// Save the session with the hased ID
	if err := s.sessionRepo.Save(session); err != nil {
		return nil, err
	}

	// We want to return the session with a non hashed ID
	session.ID = id

	return session, nil
}

func (s *service) VerifySession(sessionID string) (string, error) {
	hashedID := s.hashSessionID(sessionID)

	session, err := s.sessionRepo.Find(hashedID)

	if err != nil {
		return "", err
	}

	if session.ExpiresAt.Before(time.Now()) {
		s.sessionRepo.Delete(hashedID)

		return "", ErrExpiredSession
	}

	return session.UserID, nil
}
