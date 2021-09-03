package elements

import (
	"gopkg.in/go-playground/colors.v1"
)

type Team struct {
	Color                      colors.RGBColor
	generation_rate_multiplier float32
	population_cap_multiplier  float32
}

func NewTeam(color colors.RGBColor) *Team {
	return &Team{Color: color}
}
