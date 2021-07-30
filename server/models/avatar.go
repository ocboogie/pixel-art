package models

import (
	"math"
	"math/rand"
	"strings"
)

type Avatar string

// AvatarSpec is the info needed to create a correctly formatted avatar.
type AvatarSpec struct {
	Size    uint8    `json:"size"`
	Palette []string `json:"palette"`
}

// TODO: Probably a much better way of doing this but this should work fine for
//       now
func GenerateRandomAvatar(spec AvatarSpec) Avatar {
	pixels := make([]bool, spec.Size*spec.Size)
	sizeHalf := int(math.Ceil(float64(spec.Size) / 2))
	data := ""

	for y := 0; y < int(spec.Size); y++ {
		for x := 0; x < sizeHalf; x++ {
			pos := y*int(spec.Size) + x
			mirroredPos := y*int(spec.Size) + (int(spec.Size) - 1 - x)

			active := rand.Int31()%2 == 0

			pixels[pos] = active
			pixels[mirroredPos] = active
		}
	}

	for _, active := range pixels {
		if active {
			data += "1"
		} else {
			data += "0"
		}
	}

	data += spec.Palette[rand.Intn(len(spec.Palette))]
	return Avatar(data)
}

func (a Avatar) Validate(spec AvatarSpec) error {
	sections := strings.Split(string(a), "#")
	if len(sections) != 2 {
		return ErrInvalidAvatar
	}
	cellsString := sections[0]

	if len(cellsString) != int(spec.Size)*int(spec.Size) {
		return ErrInvalidAvatar
	}
	for _, cell := range cellsString {
		if cell != '1' && cell != '0' {
			return ErrInvalidAvatar
		}
	}
	color := sections[1]
	color = "#" + color

	for _, paletteColor := range spec.Palette {
		if color == paletteColor {
			return nil
		}
	}

	return ErrInvalidAvatar
}
