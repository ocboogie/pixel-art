package avatar

import (
	"strings"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../../mocks/service_avatar.go -package mocks -mock_names Service=ServiceAvatar github.com/ocboogie/pixel-art/services/avatar Service

type Service interface {
	Validate(data string) bool
	Format() models.Format
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
