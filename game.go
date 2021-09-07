package main

import (
	. "github.com/TimothyGregg/Antmound/elements"
	. "github.com/TimothyGregg/Antmound/graphics"
	"gopkg.in/go-playground/colors.v1"
)

type Game struct {
	Board *Board
	Teams []*Team
}

func (g *Game) generate_board() {
	g.Board = NewBoard()
	g.Board.SetSize([2]float64{500, 500}) // TODO Make Variable
	g.Board.Naive_Fill()
	g.Board.Connect_Delaunay()
}

func (g *Game) generate_teams() {
	for i := 0; i < 6; i++ {
		color, _ := colors.ParseHEX(Red)
		g.Teams = append(g.Teams, NewTeam(*color.ToRGB()))
	}
}

func NewGame() *Game {
	g := &Game{}
	g.generate_board()
	g.generate_teams()
	return g
}

func (g *Game) Update() error {
	return nil
}
