package avatar

import (
	"math"
	"math/rand"
	"strings"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../../mocks/service_avatar.go -package mocks -mock_names Service=ServiceAvatar github.com/ocboogie/pixel-art/services/avatar Service

type Service interface {
	Validate(data string) bool
	Format() models.Format
	GenerateRandom() string
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
	sections := strings.Split(data, "#")
	if len(sections) != 2 {
		return false
	}
	cellsString := sections[0]

	if len(cellsString) != int(s.config.Size)*int(s.config.Size) {
		return false
	}
	for _, cell := range cellsString {
		if cell != '1' && cell != '0' {
			return false
		}
	}
	color := sections[1]
	color = "#" + color

	for _, paletteColor := range s.config.Palette {
		if color == paletteColor {
			return true
		}
	}

	return false
}

func (s *service) Format() models.Format {
	return models.Format{
		Size:    s.config.Size,
		Palette: s.config.Palette,
	}
}

// TODO: Probably a much better way of doing this but this should work fine for
//       now
func (s *service) GenerateRandom() string {
	pixels := make([]bool, s.config.Size*s.config.Size)
	sizeHalf := int(math.Ceil(float64(s.config.Size) / 2))
	data := ""

	for y := 0; y < int(s.config.Size); y++ {
		for x := 0; x < sizeHalf; x++ {
			pos := y*int(s.config.Size) + x
			mirroredPos := y*int(s.config.Size) + (int(s.config.Size) - 1 - x)

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

	data += s.config.Palette[rand.Intn(len(s.config.Palette))]
	return data
}
