package post

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/services/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks.NewRepositoryPost(ctrl)
		s := &service{
			log:      testutils.NullLogger(),
			postRepo: repo,
		}

		repo.EXPECT().Save(gomock.AssignableToTypeOf(&models.Post{})).Return(nil)

		id, err := s.Create(models.PostNew{
			UserID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
			Title:  "Yup",
			Data:   "",
		})

		assert.NoError(t, err)
		assert.Regexp(t, `^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`, id)
	})
}

func TestFind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryPost(ctrl)
	s := &service{
		postRepo: repo,
	}

	mockPost := &models.Post{
		Author: models.User{ID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a"},
		Title:  "Yup",
		Data:   make([]byte, 0),
	}

	repo.EXPECT().Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(mockPost, nil)

	post, err := s.Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.NoError(t, err)
	assert.Equal(t, mockPost, post)
}

func TestLatest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryPost(ctrl)
	s := &service{
		postRepo: repo,
	}

	mockLatestPosts := []*models.Post{
		&models.Post{
			Author: models.User{ID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
			Data:   make([]byte, 0),
		},
		&models.Post{
			Author: models.User{ID: "6caaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
			Data:   make([]byte, 0),
		},
		&models.Post{
			Author: models.User{ID: "6aaaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
			Data:   make([]byte, 0),
		},
	}
	after := time.Now()

	repo.EXPECT().Latest(20, &after).Return(mockLatestPosts, nil)

	LatestPosts, err := s.Latest(20, &after)

	assert.NoError(t, err)
	assert.Equal(t, LatestPosts, LatestPosts)
}

func TestLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryLike(ctrl)
	s := &service{
		log:      testutils.NullLogger(),
		likeRepo: repo,
	}

	mockErr := errors.New("mock error")

	repo.EXPECT().Save("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(mockErr)

	err := s.Like("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.Equal(t, mockErr, err)
}

func TestUnlike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryLike(ctrl)
	s := &service{
		log:      testutils.NullLogger(),
		likeRepo: repo,
	}

	mockErr := errors.New("mock error")

	repo.EXPECT().Delete("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(mockErr)

	err := s.Unlike("107ed308-3984-11ea-a137-2e728ce88125", "60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.Equal(t, mockErr, err)
}
