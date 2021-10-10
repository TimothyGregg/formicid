package elements

import (
	"encoding/json"
	"fmt"

	colors "github.com/TimothyGregg/Antmound/game/graphics"
	tools "github.com/TimothyGregg/Antmound/game/tools"
)

type Team struct {
	Element
	Color                      colors.Color `json:"-"`
	unit_UID_Generator         *tools.UID_Generator
}

func (t *Team) MarshalJSON() ([]byte, error) {
	type Alias Team
	return json.Marshal(&struct {
		Color string `json:"color"`
		*Alias
	}{
		Color: fmt.Sprint(t.Color),
		Alias: (*Alias)(t),
	})
}

func New_Team(color colors.Color, uid int) *Team {
	t := &Team{Color: color}
	t.New(uid)
	t.unit_UID_Generator = tools.New_UID_Generator()
	return t
}

func (t *Team) Next_Unit_UID() int {
	return t.unit_UID_Generator.Next()
}

func (t *Team) Recycle_UID(uid int) {
	t.unit_UID_Generator.Recycle(uid)
}