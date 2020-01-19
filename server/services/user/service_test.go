package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/services/testutils"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryUser(ctrl)
	s := &service{
		userRepo: repo,
	}

	mockUser := &models.User{
		ID:       "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
		Email:    "foo@bar.com",
		Password: "correct battery horse staple but this should been hashed",
	}

	repo.EXPECT().Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(mockUser, nil)

	user, err := s.Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryUser(ctrl)
	s := &service{
		log:      testutils.NullLogger(),
		userRepo: repo,
	}

	mockUpdate := &models.User{
		ID:   "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
		Name: "test",
	}

	repo.EXPECT().Update(mockUpdate).Return(nil)

	err := s.Update(mockUpdate)

	assert.NoError(t, err)
}
