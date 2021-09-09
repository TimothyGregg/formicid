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

func (g *Game) generate_board(size_x, size_y, fill_tries int) {
	g.Board = NewBoard()
	g.Board.SetSize([2]float64{float64(size_x), float64(size_y)}) // TODO Make Variable
	g.Board.Naive_Fill(fill_tries)
	g.Board.Connect_Delaunay()
}

func (g *Game) generate_teams() {
	for i := 0; i < 6; i++ {
		color, _ := colors.ParseHEX(Red)
		g.Teams = append(g.Teams, NewTeam(*color.ToRGB()))
	}
}

func NewGame(size_x, size_y, fill_tries int) *Game {
	g := &Game{}
	g.generate_board(size_x, size_y, fill_tries)
	g.generate_teams()
	return g
}

func (g *Game) Update() error {
	return nil
}
