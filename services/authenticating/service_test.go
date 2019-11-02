package authenticating

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/pkg/argon2"
	"github.com/stretchr/testify/assert"
)

var cfg = &config.Config{
	HashConfig:      argon2.DefaultParams(),
	SessionLifetime: 7 * 24 * 60 * 60 * 1000,
	Secret:          "SECRET",
}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewUserRepository(ctrl)
	s := &Service{
		UserRepo: repo,
		Config:   cfg,
	}

	repo.EXPECT().ExistsEmail(gomock.Any()).Return(false, nil)
	repo.EXPECT().Create(gomock.AssignableToTypeOf(&models.User{})).Return(nil)

	id, err := s.SignUp(&models.UserInput{Email: "foo@bar.com", Password: "password"})

	assert.NotEmpty(t, id)
	assert.NoError(t, err)

	// TODO: Test when email exists
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := mocks.NewUserRepository(ctrl)
	sessionRepo := mocks.NewSessionRepository(ctrl)
	s := &Service{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		Config:      cfg,
	}

	hashed, err := argon2.Hash("correct battery horse staple", s.Config.HashConfig)
	assert.NoError(t, err)

	user := &models.User{
		ID:       "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
		Email:    "foo@bar.com",
		Password: hashed,
	}

	userRepo.EXPECT().FindByEmail(gomock.Eq("foo@bar.com")).Return(user, nil)
	sessionRepo.EXPECT().Create(gomock.AssignableToTypeOf(&models.Session{})).Return(nil)

	sessionID, err := s.Login("foo@bar.com", "correct battery horse staple")
	assert.NoError(t, err)
	assert.Regexp(t, `^[A-Fa-f0-9]+$`, sessionID)

	// TODO: Test when email not found
}

func TestCreateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewSessionRepository(ctrl)
	s := &Service{
		SessionRepo: repo,
		Config:      cfg,
	}

	// TODO: Test that the session IDs are correct
	repo.EXPECT().Create(gomock.AssignableToTypeOf(&models.Session{})).Return(nil)
	sessionID, err := s.CreateSession("60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.NoError(t, err)
	assert.Regexp(t, `^[A-Fa-f0-9]+$`, sessionID)
	assert.Len(t, sessionID, SessionIDLength)
}

func TestVerifySession(t *testing.T) {
	t.Run("Expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks.NewSessionRepository(ctrl)
		s := &Service{
			SessionRepo: repo,
			Config:      cfg,
		}

		session := &models.Session{
			ID:        "f76261e75790651404e62c3f8224206eb7ca85ce646746a7815b8eaffb7b0d78",
			UserID:    "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
			ExpiresAt: time.Now().Add(time.Hour),
		}

		repo.EXPECT().Find(session.ID).Return(session, nil)

		userID, err := s.VerifySession("6dcc66e6faf5b409d5b007b12a1bfd7e")

		assert.NoError(t, err)
		assert.Equal(t, session.UserID, userID)
	})
	t.Run("Expired", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks.NewSessionRepository(ctrl)
		s := &Service{
			SessionRepo: repo,
			Config:      cfg,
		}

		session := &models.Session{
			ID:        "f76261e75790651404e62c3f8224206eb7ca85ce646746a7815b8eaffb7b0d78",
			UserID:    "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
			ExpiresAt: time.Now().Add(-time.Hour),
		}

		repo.EXPECT().Find(session.ID).Return(session, nil)
		repo.EXPECT().Delete(session.ID).Return(nil)

		_, err := s.VerifySession("6dcc66e6faf5b409d5b007b12a1bfd7e")

		assert.Equal(t, ErrExpiredSession, err)
	})
}
