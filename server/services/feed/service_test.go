package feed

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
			Art:    models.Art(""),
		})

		assert.NoError(t, err)
		assert.Regexp(t, `^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`, id)
	})
}
func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryPost(ctrl)
	s := &service{
		postRepo: repo,
	}

	repo.EXPECT().Delete("60aaf13d-8ddc-403b-ba42-960e18a22f6a").Return(nil)

	err := s.Delete("60aaf13d-8ddc-403b-ba42-960e18a22f6a")

	assert.NoError(t, err)
}

func TestFind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryPost(ctrl)
	s := &service{
		postRepo: repo,
	}

	mockPost := &models.Post{
		Author: &models.PublicUser{ID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a"},
		Title:  "Yup",
		Art:    models.Art(""),
	}

	repo.EXPECT().Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a", PostIncludes{Author: true}).Return(mockPost, nil)

	post, err := s.Find("60aaf13d-8ddc-403b-ba42-960e18a22f6a", PostIncludes{Author: true})

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
		{
			Author: &models.PublicUser{ID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
			Art:    models.Art(""),
		},
		{
			Author: &models.PublicUser{ID: "6caaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
			Art:    models.Art(""),
		},
		{
			Author: &models.PublicUser{ID: "6aaaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
			Art:    models.Art(""),
		},
	}
	after := time.Now()

	repo.EXPECT().Latest(20, &after, PostIncludes{Author: true}).Return(mockLatestPosts, nil)

	latestPosts, err := s.Latest(20, &after, PostIncludes{Author: true})

	assert.NoError(t, err)
	assert.Equal(t, mockLatestPosts, latestPosts)
}

func TestPostsByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewRepositoryPost(ctrl)
	s := &service{
		postRepo: repo,
	}

	mockUsersPosts := []*models.Post{
		{
			Author: &models.PublicUser{ID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
		},
		{
			Author: &models.PublicUser{ID: "6caaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
		},
		{
			Author: &models.PublicUser{ID: "6aaaf13d-8ddc-403b-ba42-960e18a22f6a"},
			Title:  "Yup",
		},
	}
	after := time.Now()

	repo.EXPECT().PostsByUser("60aaf13d-8ddc-403b-ba42-960e18a22f6a", 20, &after, PostIncludes{Author: true}).Return(mockUsersPosts, nil)

	usersPosts, err := s.PostsByUser("60aaf13d-8ddc-403b-ba42-960e18a22f6a", 20, &after, PostIncludes{Author: true})

	assert.NoError(t, err)
	assert.Equal(t, mockUsersPosts, usersPosts)
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
