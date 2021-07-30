package models

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type Art []byte

// ArtSpec is the info needed to create a correctly formatted art for a post.
type ArtSpec struct {
	Size   uint16 `json:"size"`
	Colors uint8  `json:"colors"`
}

func (art Art) Validate(spec ArtSpec) error {
	decoder := bufio.NewReader(bytes.NewBuffer(art))

	// Check for invalid width
	widthB := make([]byte, 2)
	var width uint16
	if _, err := decoder.Read(widthB); err != nil {
		return ErrInvalidArt
	}
	width = binary.BigEndian.Uint16(widthB)
	if width != spec.Size {
		return ErrInvalidArt
	}

	// Check for invalid height
	heightB := make([]byte, 2)
	var height uint16
	if _, err := decoder.Read(heightB); err != nil {
		return ErrInvalidArt
	}
	height = binary.BigEndian.Uint16(heightB)
	if height != spec.Size {
		return ErrInvalidArt
	}

	// Get and check color amount
	colorAmount, err := decoder.ReadByte()
	if err != nil {
		return ErrInvalidArt
	}
	if colorAmount != spec.Colors {
		return ErrInvalidArt
	}

	// Check and pass the color table
	if _, err := decoder.Discard(int(colorAmount) * 3); err != nil {
		return ErrInvalidArt
	}

	// Get and check the body
	artBody := make([]byte, int(width)*int(height))
	if _, err = io.ReadFull(decoder, artBody); err != nil {
		return ErrInvalidArt
	}
	if decoder.Buffered() != 0 {
		return ErrInvalidArt
	}

	// Check if any of the pixel indices are out of range of the color table
	for _, colorIndex := range artBody {
		if colorIndex >= colorAmount {
			return ErrInvalidArt
		}
	}

	return nil
}
