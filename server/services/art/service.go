package art

import (
	"bufio"
	"encoding/base64"
	"encoding/binary"
	"io"
	"strings"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../../mocks/service_art.go -package mocks -mock_names Service=ServiceArt github.com/ocboogie/pixel-art/services/art Service

type Service interface {
	Validate(data string) bool
	Format() models.ArtFormat
}

type service struct {
	config Config
}

func New(config Config) Service {
	return &service{
		config: config,
	}
}

func (s *service) Validate(data string) bool {
	decoder := bufio.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(data)))

	// Check for invalid width
	widthB := make([]byte, 2)
	var width uint16
	if _, err := decoder.Read(widthB); err != nil {
		return false
	}
	width = binary.BigEndian.Uint16(widthB)
	if width != s.config.Size {
		return false
	}

	// Check for invalid height
	heightB := make([]byte, 2)
	var height uint16
	if _, err := decoder.Read(heightB); err != nil {
		return false
	}
	height = binary.BigEndian.Uint16(heightB)
	if height != s.config.Size {
		return false
	}

	// Get and check color amount
	colorAmount, err := decoder.ReadByte()
	if err != nil {
		return false
	}
	if colorAmount != s.config.Colors {
		return false
	}

	// Check and pass the color table
	if _, err := decoder.Discard(int(colorAmount) * 3); err != nil {
		return false
	}

	// Get and check the body
	artBody := make([]byte, int(width)*int(height))
	if _, err = io.ReadFull(decoder, artBody); err != nil {
		return false
	}
	if decoder.Buffered() != 0 {
		return false
	}

	// Check if any of the pixel indices are out of range of the color table
	for _, colorIndex := range artBody {
		if colorIndex >= colorAmount {
			return false
		}
	}

	return true
}

func (s *service) Format() models.ArtFormat {
	return models.ArtFormat{
		Size:   s.config.Size,
		Colors: s.config.Colors,
	}
}
