package profile

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
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

	repo.EXPECT().Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a", UserIncludes{Following: "c41270f0-0555-4923-999e-4c798bd47f01"}).Return(mockUser, nil)

	user, err := s.Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a", UserIncludes{Following: "c41270f0-0555-4923-999e-4c798bd47f01"})

	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryUser(ctrl)
	s := &service{
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

func TestFollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryFollow(ctrl)
	s := &service{
		followsRepo: repo,
	}

	mockErr := errors.New("mock error")

	repo.EXPECT().Save("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(mockErr)

	err := s.Follow("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.Equal(t, mockErr, err)
}

func TestUnfollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryFollow(ctrl)
	s := &service{
		followsRepo: repo,
	}

	mockErr := errors.New("mock error")

	repo.EXPECT().Delete("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(mockErr)

	err := s.Unfollow("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.Equal(t, mockErr, err)
}
