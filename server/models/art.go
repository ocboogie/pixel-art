package models

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"math/rand"
)

// Form:
// Width|Height|Color Amount|Colors Table      |Body
// 0003  0003   03           FF000000FF000000FF 000102000102000103
//
// Name         | Size                        | Type
// -------------|-----------------------------|---------------------------
// Width        | 16 bits                     | Big endian unsigned int
// Height       | 16 bits                     | Big endian unsigned int
// Color Amount | 8 bits                      | Unsigned int
// Color Table  | 24 * {Color Amount} bits    | Hex color
// Body         | 8 * {Width} * {Height} bits | Indexed color based on the table
type Art []byte

// ArtSpec is the info needed to create a correctly formatted art for a post.
type ArtSpec struct {
	Size   uint16 `json:"size"`
	Colors uint8  `json:"colors"`
}

type Writer int

func GenerateRandomArt(spec ArtSpec) (Art, error) {
	width := 2
	height := 2
	colorAmount := 1
	colorTable := 3 * int(spec.Colors)
	body := int(spec.Size) * int(spec.Size)

	data := bytes.NewBuffer(make([]byte, 0, width+height+colorAmount+colorTable+body))

	// Width
	if err := binary.Write(data, binary.BigEndian, spec.Size); err != nil {
		return nil, err
	}

	// Height
	if err := binary.Write(data, binary.BigEndian, spec.Size); err != nil {
		return nil, err
	}

	// Color amount
	if err := data.WriteByte(spec.Colors); err != nil {
		return nil, err
	}

	// Color table
	for i := 0; i < int(spec.Colors); i++ {
		color := make([]byte, 3)

		if _, err := rand.Read(color); err != nil {
			return nil, err
		}

		if _, err := data.Write(color); err != nil {
			return nil, err
		}
	}

	// Body
	for i := 0; i < int(spec.Size)*int(spec.Size); i++ {
		data.WriteByte(byte(rand.Int63n(int64(spec.Colors))))
	}

	return Art(data.Bytes()), nil
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
