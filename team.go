package main

import (
	"gopkg.in/go-playground/colors.v1"
)

type team struct {
	color colors.RGBColor
}

func NewTeam(color colors.RGBColor) *team {
	return &team{color: color}
}
