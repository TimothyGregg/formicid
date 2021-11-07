package storage

import (
	"github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/game/tools"
)

type Store struct {
	Games         []*game.Game
	UID_Generator *tools.UID_Generator
}

func NewStore() *Store {
	s := &Store{}
	s.UID_Generator = tools.New_UID_Generator()
	return s
}

func (s *Store) New_Game(size_x, size_y int) {
	s.Games = append(s.Games, game.New_Game(s.UID_Generator.Next(), size_x, size_y))
}