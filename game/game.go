package game

import (
	"encoding/json"

	elements "github.com/TimothyGregg/formicid/game/elements"
	graphics "github.com/TimothyGregg/formicid/game/graphics"
	util "github.com/TimothyGregg/formicid/game/util"
)

// An similar concept: https://np.ironhelmet.com/

type Game struct {
	UID                int              `json:"uid"`
	Board              *elements.Board  `json:"board"`
	Teams              []*elements.Team `json:"teams"`
	team_uid_generator *util.UID_Generator
}

func (g *Game) MarshalJSON() ([]byte, error) {
	type Alias Game
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(g),
	})
}

func (g *Game) generate_board(size_x, size_y int) {
	g.Board = elements.New_Board()
	g.Board.Set_Size([2]int{size_x, size_y}) // TODO Make Variable
	g.Board.Fill()
	g.Board.Connect()
}

func (g *Game) generate_teams() { // Improve to randomize colors
	for i := 0; i < 6; i++ {
		new_number := g.team_uid_generator.Next()
		color := graphics.Color(new_number)
		g.Teams = append(g.Teams, elements.New_Team(color, new_number))
	}
}

func New_Game(uid, size_x, size_y int) *Game {
	g := &Game{UID: uid}
	g.team_uid_generator = util.New_UID_Generator()
	g.generate_board(size_x, size_y)
	g.generate_teams()
	return g
}

func (g *Game) Update() error {
	return nil
}
