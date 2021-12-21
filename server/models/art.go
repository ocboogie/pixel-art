package models

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"image/color"
	"io"
	"math/rand"
)

type Art struct {
	Width      uint16
	Height     uint16
	TableSize  uint8
	ColorTable []color.Color
	Body       []uint8
}

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
type ArtEncoded []byte

// ArtSpec is the info needed to create a correctly formatted art for a post.
type ArtSpec struct {
	Size   uint16 `json:"size"`
	Colors uint8  `json:"colors"`
}

func GenerateRandomArt(spec ArtSpec) (ArtEncoded, error) {
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

	return ArtEncoded(data.Bytes()), nil
}

// func (art Art) ToPaletted() *image.Paletted {
// 	decoder := bufio.NewReader(bytes.NewBuffer(art))
//
// 	widthB := make([]byte, 2)
// 	var width uint16
// 	if _, err := decoder.Read(widthB); err != nil {
// 		return nil
// 	}
// 	width = binary.BigEndian.Uint16(widthB)
//
// }

func (art ArtEncoded) Decode() (ArtDecoded, error) {

}

func (art Art) Decode(spec ArtSpec) (ArtDecoded, error) {
	decoder := bufio.NewReader(bytes.NewBuffer(art))
	var artDecoded Art

	// Check for invalid width
	widthB := make([]byte, 2)
	if _, err := decoder.Read(widthB); err != nil {
		return artDecoded, ErrInvalidArt
	}
	artDecoded.Width = binary.BigEndian.Uint16(widthB)
	if artDecoded.Width != spec.Size {
		return artDecoded, ErrInvalidArt
	}

	// Check for invalid height
	heightB := make([]byte, 2)
	if _, err := decoder.Read(heightB); err != nil {
		return artDecoded, ErrInvalidArt
	}
	artDecoded.Height = binary.BigEndian.Uint16(heightB)
	if artDecoded.Height != spec.Size {
		return artDecoded, ErrInvalidArt
	}

	// Get and check color amount
	colorAmount, err := decoder.ReadByte()
	if err != nil {
		return artDecoded, ErrInvalidArt
	}
	if colorAmount != spec.Colors {
		return artDecoded, ErrInvalidArt
	}
	artDecoded.TableSize = colorAmount

	// Check and pass the color table
	if _, err := decoder.Discard(int(colorAmount) * 3); err != nil {
		return artDecoded, ErrInvalidArt
	}

	// Get and check the body
	artDecoded.Body = make([]uint8, int(artDecoded.Width)*int(artDecoded.Height))
	if _, err = io.ReadFull(decoder, artDecoded.Body); err != nil {
		return artDecoded, ErrInvalidArt
	}
	if decoder.Buffered() != 0 {
		return artDecoded, ErrInvalidArt
	}

	// Check if any of the pixel indices are out of range of the color table
	for _, colorIndex := range artDecoded.Body {
		if colorIndex >= colorAmount {
			return artDecoded, ErrInvalidArt
		}
	}

	return artDecoded, nil
}
