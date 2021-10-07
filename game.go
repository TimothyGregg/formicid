package main

import (
	elements "github.com/TimothyGregg/Antmound/elements"
	graphics "github.com/TimothyGregg/Antmound/graphics"
	"gopkg.in/go-playground/colors.v1"
)

type Game struct {
	Board *elements.Board
	Teams []*elements.Team
}

func (g *Game) generate_board(size_x, size_y int) {
	g.Board = elements.New_Board()
	g.Board.Set_Size([2]float64{float64(size_x), float64(size_y)}) // TODO Make Variable
	g.Board.Fill()
	g.Board.Connect()
}

func (g *Game) generate_teams() {
	for i := 0; i < 6; i++ {
		color, _ := colors.ParseHEX(graphics.Red)
		g.Teams = append(g.Teams, elements.New_Team(*color.ToRGB()))
	}
}

func NewGame(size_x, size_y, fill_tries int) *Game {
	g := &Game{}
	g.generate_board(size_x, size_y)
	g.generate_teams()
	return g
}

func (g *Game) Update() error {
	return nil
}
