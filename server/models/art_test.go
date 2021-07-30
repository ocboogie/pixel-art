package models

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var artSpec = ArtSpec{
	Size:   3,
	Colors: 3,
}

func hexToByteSlice(hexCode string) []byte {
	bytes, err := hex.DecodeString(hexCode)
	if err != nil {
		panic(err)
	}

	return bytes
}

func TestArtValidate(t *testing.T) {
	t.Run("Invalid width", func(t *testing.T) {
		assert.Equal(t, ErrInvalidArt, Art(make([]byte, 0)).Validate(artSpec))
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("01")).Validate(artSpec))
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0010")).Validate(artSpec))
	})
	t.Run("Invalid height", func(t *testing.T) {
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0003")).Validate(artSpec))
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("000303")).Validate(artSpec))
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("00030300")).Validate(artSpec))
	})
	t.Run("Invalid color amount", func(t *testing.T) {
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("00030003")).Validate(artSpec))
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0003000304")).Validate(artSpec))
	})
	t.Run("Invalid color table", func(t *testing.T) {
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0003000303FF000000")).Validate(artSpec))
	})
	t.Run("Invalid size", func(t *testing.T) {
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0003000303FF000000FF000000FF0001020001020001")).Validate(artSpec))
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0003000303FF000000FF000000FF00010200010200010201")).Validate(artSpec))
	})
	t.Run("Color index out of range", func(t *testing.T) {
		assert.Equal(t, ErrInvalidArt, Art(hexToByteSlice("0003000303FF000000FF000000FF000102000102000103")).Validate(artSpec))
	})
	t.Run("Valid", func(t *testing.T) {
		assert.Nil(t, Art(hexToByteSlice("0003000303FF000000FF000000FF000102000102000102")).Validate(artSpec))
	})
}
