package posting

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

		repo.EXPECT().Create(gomock.AssignableToTypeOf(&models.Post{})).Return(nil)

		id, err := s.Post(models.PostInput{
			UserID: "60aaf13d-8ddc-403b-ba42-960e18a22f6a",
			Title:  "Yup",
			Data:   make([]byte, 0),
		})

		assert.NoError(t, err)
		assert.Regexp(t, `^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`, id)
	})
	t.Run("Invalid post", func(t *testing.T) {
		s := &service{}

		_, err := s.Post(models.PostInput{
			UserID: "a",
			Title:  "",
			Data:   make([]byte, 0),
		})

		assert.EqualError(t, err, `Invalid post: Key: 'PostInput.UserID' Error:Field validation for 'UserID' failed on the 'uuid' tag
Key: 'PostInput.Title' Error:Field validation for 'Title' failed on the 'required' tag`)
	})
}
