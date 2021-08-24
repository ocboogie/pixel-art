package models

import (
	"encoding/hex"
	"math/rand"
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

func TestGenerateRandomArt(t *testing.T) {
	t.Run("Is valid", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			art, err := GenerateRandomArt(artSpec)
			assert.Nil(t, err)

			if err = art.Validate(artSpec); err != nil {
				t.Errorf("GenerateRandomArt failed to create valid art. Generated %v", art)
				return
			}
		}
	})

	t.Run("Expected return", func(t *testing.T) {
		rand.Seed(0)
		art, err := GenerateRandomArt(artSpec)

		assert.Nil(t, err)

		assert.Equal(t, Art(hexToByteSlice("00030003030194fdc2fa2ffcc041010200000000000000")), art)
	})
}
