package avatar

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cfg = Config{
	Size: 3,
	Palette: []string{
		"#1abc9c",
		"#e74c3c",
	},
}

func TestValidate(t *testing.T) {
	s := &service{
		config: cfg,
	}

	t.Run("Invalid size", func(t *testing.T) {
		assert.False(t, s.Validate("0010001110000001010111011#3498db"))
		assert.False(t, s.Validate("1011#3498db"))
	})
	t.Run("No color", func(t *testing.T) {
		assert.False(t, s.Validate("101010101"))
	})
	t.Run("Invalid color", func(t *testing.T) {
		assert.False(t, s.Validate("101010101#1b"))
		assert.False(t, s.Validate("101010101#1abc9cf0"))
	})
	// Testing colors that aren't in the colors palette
	t.Run("Not in color palette", func(t *testing.T) {
		assert.False(t, s.Validate("101010101#569186"))
	})
	t.Run("Extra junk", func(t *testing.T) {
		assert.False(t, s.Validate("101010101#569186#arstars"))
	})
	t.Run("Invalid image data", func(t *testing.T) {
		assert.False(t, s.Validate("aiersntet#569186#arstars"))
	})

	t.Run("Valid", func(t *testing.T) {
		assert.True(t, s.Validate("101010101#1abc9c"))
		assert.True(t, s.Validate("101010101#e74c3c"))
	})
}

func TestFormat(t *testing.T) {
	s := &service{
		config: cfg,
	}

	format := s.Format()
	assert.Equal(t, cfg.Size, format.Size)
	assert.Equal(t, cfg.Palette, format.Palette)
}

func TestGenerateRandom(t *testing.T) {
	s := &service{
		config: cfg,
	}

	t.Run("Is valid", func(t *testing.T) {
		assert.True(t, s.Validate(s.GenerateRandom()))
	})
	// This may fail if the implementation changes
	t.Run("Expected return", func(t *testing.T) {
		rand.Seed(0)
		avatar := s.GenerateRandom()
		assert.Equal(t, "111010010#e74c3c", avatar)
	})

}
