package elements

import (
	"github.com/TimothyGregg/Antmound/tools"
	"gopkg.in/go-playground/colors.v1"
)

type Team struct {
	Element
	Color                      colors.RGBColor
	generation_rate_multiplier float32
	population_cap_multiplier  float32
	Unit_UID_Generator         *tools.UID_Generator
}

func New_Team(color colors.RGBColor) *Team {
	t := &Team{Color: color}
	t.Unit_UID_Generator = tools.New_UID_Generator()
	return t
}
