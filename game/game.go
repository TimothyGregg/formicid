package game

import (
	elements "github.com/TimothyGregg/Antmound/game/elements"
	graphics "github.com/TimothyGregg/Antmound/game/graphics"
	"github.com/TimothyGregg/Antmound/game/tools"
)

type Game struct {
	Board *elements.Board
	Teams []*elements.Team
	team_uid_generator *tools.UID_Generator
}

func (g *Game) generate_board(size_x, size_y int) {
	g.Board = elements.New_Board()
	g.Board.Set_Size([2]int{size_x, size_y}) // TODO Make Variable
	g.Board.Fill()
	g.Board.Connect()
}

func (g *Game) generate_teams() {
	for i := 0; i < 6; i++ {
		g.Teams = append(g.Teams, elements.New_Team(graphics.Red, g.team_uid_generator.Next()))
	}
}

func NewGame(size_x, size_y, fill_tries int) *Game {
	g := &Game{}
	g.team_uid_generator = tools.New_UID_Generator()
	g.generate_board(size_x, size_y)
	g.generate_teams()
	return g
}

func (g *Game) Update() error {
	return nil
}
