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

func TestArtDecode(t *testing.T) {
	assertInvalid := func(art ArtEncoded) {
		_, err := art.Decode(artSpec)
		assert.Equal(t, ErrInvalidArt, err)
	}

	t.Run("Invalid width", func(t *testing.T) {
		assertInvalid(ArtEncoded(make([]byte, 0)))
		assertInvalid(ArtEncoded(hexToByteSlice("01")))
		assertInvalid(ArtEncoded(hexToByteSlice("0010")))
	})
	t.Run("Invalid height", func(t *testing.T) {
		assertInvalid(ArtEncoded(hexToByteSlice("0003")))
		assertInvalid(ArtEncoded(hexToByteSlice("000303")))
		assertInvalid(ArtEncoded(hexToByteSlice("00030300")))
	})
	t.Run("Invalid color amount", func(t *testing.T) {
		assertInvalid(ArtEncoded(hexToByteSlice("00030003")))
		assertInvalid(ArtEncoded(hexToByteSlice("0003000304")))
	})
	t.Run("Invalid color table", func(t *testing.T) {
		assertInvalid(ArtEncoded(hexToByteSlice("0003000303FF000000")))
	})
	t.Run("Invalid size", func(t *testing.T) {
		assertInvalid(ArtEncoded(hexToByteSlice("0003000303FF000000FF000000FF0001020001020001")))
		assertInvalid(ArtEncoded(hexToByteSlice("0003000303FF000000FF000000FF00010200010200010201")))
	})
	t.Run("Color index out of range", func(t *testing.T) {
		assertInvalid(ArtEncoded(hexToByteSlice("0003000303FF000000FF000000FF000102000102000103")))
	})
	t.Run("Valid", func(t *testing.T) {
		decoded, err := ArtEncoded(hexToByteSlice("0003000303FF000000FF000000FF000102000102000102")).Decode(artSpec)

		assert.NoError(t, err)
		assert.Equal(t, Art{
			Width:     3,
			Height:    3,
			TableSize: 3,
			Body:      []uint8{0, 1, 2, 0, 1, 2, 0, 1, 2},
		}, decoded)
	})
}

func TestGenerateRandomArt(t *testing.T) {
	t.Run("Is valid", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			art, err := GenerateRandomArt(artSpec)
			assert.Nil(t, err)

			if _, err = art.Decode(artSpec); err != nil {
				t.Errorf("GenerateRandomArt failed to create valid art. Generated %v", art)
				return
			}
		}
	})

	t.Run("Expected return", func(t *testing.T) {
		rand.Seed(0)
		art, err := GenerateRandomArt(artSpec)

		assert.Nil(t, err)

		assert.Equal(t, ArtEncoded(hexToByteSlice("00030003030194fdc2fa2ffcc041010200000000000000")), art)
	})
}
