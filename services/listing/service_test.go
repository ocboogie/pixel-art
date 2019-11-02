package listing

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ocboogie/pixel-art/mocks"
	"github.com/ocboogie/pixel-art/models"
	"github.com/stretchr/testify/assert"
)

func TestLatest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewPostRepository(ctrl)
	s := &Service{
		PostRepo: repo,
	}

	mockLatestPosts := []*models.Post{
		&models.Post{
			UserID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
			Title:  "Yup",
			Data:   make([]byte, 0),
		},
	}

	repo.EXPECT().Latest(gomock.Any()).Return(mockLatestPosts, nil)

	LatestPosts, err := s.Latest()

	assert.NoError(t, err)
	assert.Equal(t, LatestPosts, LatestPosts)
}
