package game

import (
	"encoding/json"

	elements "github.com/TimothyGregg/formicid/game/elements"
	"github.com/TimothyGregg/formicid/game/util/uid"
)

// An similar concept: https://np.ironhelmet.com/

type Game struct {
	UID   *uid.UID        `json:"-"`
	Board *elements.Board `json:"-"`
}

func (g *Game) MarshalJSON() ([]byte, error) {
	type Alias Game
	return json.Marshal(&struct {
		UID     int `json:"uid"`
		BoardID int `json:"board_id"`
		*Alias
	}{
		UID:     g.UID.Value(),
		BoardID: g.Board.UID.Value(),
		Alias:   (*Alias)(g),
	})
}

func (g *Game) UnmarshalJSON(data []byte) error {
	return nil
}

func New_Game(game_uid *uid.UID, size_x, size_y int) *Game {
	g := &Game{}
	g.UID = game_uid
	g.Board, _ = elements.New_Board(uid.NewUID(0), size_x, size_y)
	return g
}

func (g *Game) Update() error {
	return nil
}
