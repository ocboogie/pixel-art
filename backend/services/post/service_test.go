package post

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	t.Run("Expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mocks.NewRepositoryPost(ctrl)
		s := &service{
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

	repo.EXPECT().Latest(gomock.Any()).Return(mockLatestPosts, nil)

	LatestPosts, err := s.Latest(20)

	assert.NoError(t, err)
	assert.Equal(t, LatestPosts, LatestPosts)
}
