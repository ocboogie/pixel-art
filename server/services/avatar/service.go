package avatar

import (
	"strings"
)

//go:generate mockgen -destination=../../mocks/service_avatar.go -package mocks -mock_names Service=ServiceAvatar github.com/ocboogie/pixel-art/services/avatar Service

type Service interface {
	Validate(data string) bool
}

type service struct {
	config Config
}

func New() Service {
	return &service{}
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

	for _, paletteColor := range s.config.Colors {
		if color == paletteColor {
			return true
		}
	}

	return false
}
