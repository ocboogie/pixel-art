package art

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cfg = &Config{
	Size:   3,
	Colors: 3,
}

func hexToBase64(hexCode string) string {
	bytes, err := hex.DecodeString(hexCode)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

func TestValidate(t *testing.T) {
	s := &service{
		config: cfg,
	}

	t.Run("Invalid width", func(t *testing.T) {
		assert.False(t, s.Validate(hexToBase64("01")))
		assert.False(t, s.Validate(hexToBase64("0010")))
	})
	t.Run("Invalid height", func(t *testing.T) {
		assert.False(t, s.Validate(hexToBase64("000303")))
		assert.False(t, s.Validate(hexToBase64("00030300")))
	})
	t.Run("Invalid color amount", func(t *testing.T) {
		assert.False(t, s.Validate(hexToBase64("0003000304")))
	})
	t.Run("Invalid color table", func(t *testing.T) {
		assert.False(t, s.Validate(hexToBase64("0003000303FF000000")))
	})
	t.Run("Invalid size", func(t *testing.T) {
		assert.False(t, s.Validate(hexToBase64("0003000303FF000000FF000000FF0001020001020001")))
		assert.False(t, s.Validate(hexToBase64("0003000303FF000000FF000000FF00010200010200010201")))
	})
	t.Run("Color index out of range", func(t *testing.T) {
		assert.False(t, s.Validate(hexToBase64("0003000303FF000000FF000000FF000102000102000103")))
	})
	t.Run("Valid", func(t *testing.T) {
		assert.True(t, s.Validate(hexToBase64("0003000303FF000000FF000000FF000102000102000102")))
	})
}
