package auth

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/pkg/argon2"
	"github.com/ocboogie/pixel-art/repositories"
	"github.com/ocboogie/pixel-art/services/testutils"
	"github.com/stretchr/testify/assert"
)

var cfg = Config{
	HashConfig:      argon2.DefaultParams(),
	SessionLifetime: 7 * 24 * 60 * 60 * 1000,
	Secret:          "SECRET",
}

func TestSignUp(t *testing.T) {
	t.Run("Expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mocks.NewRepositoryUser(ctrl)
		sessionRepo := mocks.NewRepositorySession(ctrl)
		avatarService := mocks.NewServiceAvatar(ctrl)
		s := &service{
			log:           testutils.NullLogger(),
			userRepo:      userRepo,
			sessionRepo:   sessionRepo,
			avatarService: avatarService,
			config:        cfg,
		}

		avatarService.EXPECT().Validate(gomock.Any()).Return(true)
		userRepo.EXPECT().ExistsEmail(gomock.Any()).Return(false, nil)
		userRepo.EXPECT().Save(gomock.AssignableToTypeOf(&models.User{})).Return(nil)
		sessionRepo.EXPECT().Save(gomock.AssignableToTypeOf(&models.Session{})).Return(nil)

		id, err := s.SignUp(&models.UserNew{Name: "Boogie", Avatar: "1010101010101011100101011#2ecc71", Email: "foo@bar.com", Password: "password"})

		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("Email in use", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mocks.NewRepositoryUser(ctrl)
		sessionRepo := mocks.NewRepositorySession(ctrl)
		avatarService := mocks.NewServiceAvatar(ctrl)
		s := &service{
			log:           testutils.NullLogger(),
			userRepo:      userRepo,
			sessionRepo:   sessionRepo,
			avatarService: avatarService,
			config:        cfg,
		}

		avatarService.EXPECT().Validate(gomock.Any()).Return(true)
		userRepo.EXPECT().ExistsEmail(gomock.Any()).Return(true, nil)

		_, err := s.SignUp(&models.UserNew{Name: "Boogie", Avatar: "1010101010101011100101011#2ecc71", Email: "foo@bar.com", Password: "password"})

		assert.Equal(t, err, ErrEmailAlreadyInUse)
	})

	t.Run("Invalid avatar", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mocks.NewRepositoryUser(ctrl)
		sessionRepo := mocks.NewRepositorySession(ctrl)
		avatarService := mocks.NewServiceAvatar(ctrl)
		s := &service{
			log:           testutils.NullLogger(),
			userRepo:      userRepo,
			sessionRepo:   sessionRepo,
			avatarService: avatarService,
			config:        cfg,
		}

		avatarService.EXPECT().Validate(gomock.Any()).Return(false)

		_, err := s.SignUp(&models.UserNew{Name: "Boogie", Avatar: "1010101010101011100101011#2ecc71", Email: "foo@bar.com", Password: "password"})

		assert.Equal(t, err, ErrInvalidAvatar)
	})
}

func TestLogin(t *testing.T) {
	// TODO: This takes way too long for a test
	hashed, err := argon2.Hash("correct battery horse staple", cfg.HashConfig)
	assert.NoError(t, err)

	user := &models.User{
		ID:       "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
		Email:    "foo@bar.com",
		Password: hashed,
	}

	t.Run("Expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mocks.NewRepositoryUser(ctrl)
		sessionRepo := mocks.NewRepositorySession(ctrl)
		s := &service{
			log:         testutils.NullLogger(),
			userRepo:    userRepo,
			sessionRepo: sessionRepo,
			config:      cfg,
		}

		userRepo.EXPECT().FindByEmail(gomock.Eq("foo@bar.com")).Return(user, nil)
		sessionRepo.EXPECT().Save(gomock.AssignableToTypeOf(&models.Session{})).Return(nil)

		session, err := s.Login(&models.UserCredentials{
			Email:    "foo@bar.com",
			Password: "correct battery horse staple",
		})

		assert.NoError(t, err)
		assert.IsType(t, &models.Session{}, session)
		assert.Regexp(t, `^[A-Fa-f0-9]+$`, session.ID)
	})

	t.Run("Email not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mocks.NewRepositoryUser(ctrl)
		sessionRepo := mocks.NewRepositorySession(ctrl)
		s := &service{
			log:         testutils.NullLogger(),
			userRepo:    userRepo,
			sessionRepo: sessionRepo,
			config:      cfg,
		}

		userRepo.EXPECT().FindByEmail(gomock.Eq("foo@bar.com")).Return(nil, repositories.ErrUserNotFound)

		_, err := s.Login(&models.UserCredentials{
			Email: "foo@bar.com",
		})
		assert.Equal(t, ErrInvalidCredentials, err)
	})

	t.Run("Invalid credentials", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userRepo := mocks.NewRepositoryUser(ctrl)
		sessionRepo := mocks.NewRepositorySession(ctrl)
		s := &service{
			log:         testutils.NullLogger(),
			userRepo:    userRepo,
			sessionRepo: sessionRepo,
			config:      cfg,
		}

		userRepo.EXPECT().FindByEmail(gomock.Eq("foo@bar.com")).Return(user, nil)

		_, err := s.Login(&models.UserCredentials{
			Email:    "foo@bar.com",
			Password: "test",
		})
		assert.Equal(t, ErrInvalidCredentials, err)
	})
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sessionRepo := mocks.NewRepositorySession(ctrl)
	s := &service{
		log:         testutils.NullLogger(),
		sessionRepo: sessionRepo,
		config:      cfg,
	}

	sessionRepo.EXPECT().Delete("the session id").Return(nil)

	err := s.Logout("the session id")

	assert.NoError(t, err)
}

func TestCreateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositorySession(ctrl)
	s := &service{
		log:         testutils.NullLogger(),
		sessionRepo: repo,
		config:      cfg,
	}

	// TODO: Test that the session IDs are correct
	repo.EXPECT().Save(gomock.AssignableToTypeOf(&models.Session{})).Return(nil)
	session, err := s.CreateSession("60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.NoError(t, err)
	assert.IsType(t, &models.Session{}, session)
	assert.Regexp(t, `^[A-Fa-f0-9]+$`, session.ID)
	assert.Len(t, session.ID, SessionIDLength)
}

func TestVerifySession(t *testing.T) {
	t.Run("Expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks.NewRepositorySession(ctrl)
		s := &service{
			log:         testutils.NullLogger(),
			sessionRepo: repo,
			config:      cfg,
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
		repo := mocks.NewRepositorySession(ctrl)
		s := &service{
			log:         testutils.NullLogger(),
			sessionRepo: repo,
			config:      cfg,
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
