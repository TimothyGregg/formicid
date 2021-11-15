package storage

import (
	"github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/game/util/uid"
)

type Store struct {
	Games         []*game.Game
	UID_Generator *uid.UID_Generator
}

func NewStore() *Store {
	s := &Store{}
	s.UID_Generator = uid.New_UID_Generator()
	return s
}

func (s *Store) New_Game(size_x, size_y int) {
	s.Games = append(s.Games, game.New_Game(s.UID_Generator.Next(), size_x, size_y))
}
