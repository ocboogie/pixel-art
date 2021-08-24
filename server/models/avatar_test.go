package models

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var avatarSpec = AvatarSpec{
	Size: 3,
	Palette: []string{
		"#1abc9c",
		"#e74c3c",
	},
}

func TestValidate(t *testing.T) {
	t.Run("Invalid size", func(t *testing.T) {
		assert.Equal(t, ErrInvalidAvatar, Avatar("0010001110000001010111011#3498db").Validate(avatarSpec))
		assert.Equal(t, ErrInvalidAvatar, Avatar("1011#3498db").Validate(avatarSpec))
	})
	t.Run("No color", func(t *testing.T) {
		assert.Equal(t, ErrInvalidAvatar, Avatar("101010101").Validate(avatarSpec))
	})
	t.Run("Invalid color", func(t *testing.T) {
		assert.Equal(t, ErrInvalidAvatar, Avatar("101010101#1b").Validate(avatarSpec))
		assert.Equal(t, ErrInvalidAvatar, Avatar("101010101#1abc9cf0").Validate(avatarSpec))
	})
	// Testing colors that aren't in the colors palette
	t.Run("Not in color palette", func(t *testing.T) {
		assert.Equal(t, ErrInvalidAvatar, Avatar("101010101#569186").Validate(avatarSpec))
	})
	t.Run("Extra junk", func(t *testing.T) {
		assert.Equal(t, ErrInvalidAvatar, Avatar("101010101#569186#arstars").Validate(avatarSpec))
	})
	t.Run("Invalid image data", func(t *testing.T) {
		assert.Equal(t, ErrInvalidAvatar, Avatar("aiersntet#569186#arstars").Validate(avatarSpec))
	})

	t.Run("Valid", func(t *testing.T) {
		assert.Nil(t, Avatar("101010101#1abc9c").Validate(avatarSpec))
		assert.Nil(t, Avatar("101010101#e74c3c").Validate(avatarSpec))
	})
}

func TestGenerateRandom(t *testing.T) {
	t.Run("Is valid", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			avatar := GenerateRandomAvatar(avatarSpec)

			if err := avatar.Validate(avatarSpec); err != nil {
				t.Errorf("GenerateRandomAvatar failed to create a valid avatar. Generated %v", avatar)
				return
			}
		}
	})

	// This may fail if the implementation changes
	t.Run("Expected return", func(t *testing.T) {
		rand.Seed(0)
		avatar := GenerateRandomAvatar(avatarSpec)
		assert.Equal(t, Avatar("111010010#e74c3c"), avatar)
	})

}
