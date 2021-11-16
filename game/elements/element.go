package elements

import (
	"github.com/TimothyGregg/formicid/game/util/uid"
)

type Element struct {
	UID   *uid.UID `json:"uid"`
	Timer uint8    `json:"timer"`
}

func (e *Element) tick() {
	e.Timer = (e.Timer + 1) % 60
}

func (e *Element) update() {
	e.tick()
}
