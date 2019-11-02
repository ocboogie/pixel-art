package authenticating

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/pkg/argon2"
	"github.com/ocboogie/pixel-art/repositories"
)

// The number of hex characters in a session ID
const SessionIDLength = 32

type Service struct {
	UserRepo    repositories.User
	SessionRepo repositories.Session
	Config      *config.Config
}

func (s *Service) SignUp(user *models.UserInput) (string, error) {
	if err := user.Validate(); err != nil {
		return "", &ErrInvalidUser{Err: err}
	}

	exists, err := s.UserRepo.ExistsEmail(user.Email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", ErrEmailAlreadyInUse
	}

	idBytes, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := idBytes.String()

	hashedPassword, err := argon2.Hash(user.Password, s.Config.HashConfig)
	if err != nil {
		return "", err
	}

	userHashed := &models.User{
		ID:        id,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := s.UserRepo.Create(userHashed); err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) Login(email string, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		// TODO: Doc why this is here
		argon2.Hash(password, s.Config.HashConfig)
		return "", err
	}

	match, err := argon2.Verify(password, user.Password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", ErrInvalidCredentials
	}

	return s.CreateSession(user.ID)
}

func (s *Service) hashSessionID(id string) string {
	hasher := sha256.New()
	hasher.Write([]byte(id))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *Service) CreateSession(userID string) (string, error) {
	// Dividing by 2 because for each byte we will have two hex characters
	b := make([]byte, SessionIDLength/2)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	id := hex.EncodeToString(b)

	hashedID := s.hashSessionID(id)
	expiresAt := time.Now().Add(time.Duration(s.Config.SessionLifetime) * time.Second)

	session := &models.Session{
		ID:        hashedID,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}

	if err := s.SessionRepo.Create(session); err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) VerifySession(sessionID string) (string, error) {
	hashedID := s.hashSessionID(sessionID)

	session, err := s.SessionRepo.Find(hashedID)

	if err != nil {
		return "", err
	}

	if session.ExpiresAt.Before(time.Now()) {
		s.SessionRepo.Delete(hashedID)

		return "", ErrExpiredSession
	}

	return session.UserID, nil
}
