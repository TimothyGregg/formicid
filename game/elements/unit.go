package elements

import "encoding/json"

type Unit struct {
	Element
	Population int
	Team       *Team `json:"-"`
}

func (u *Unit) MarshalJSON() ([]byte, error) {
	type Alias Unit
	return json.Marshal(&struct {
		Team int `json:"teamID"`
		*Alias
	}{
		Team: u.Team.UID,
		Alias: (*Alias)(u),
	})
}

func New_Unit(uid, population int) *Unit {
	u := &Unit{Population: population}
	u.New(uid)
	return u
}

func (u *Unit) Destroy() {
	u.Team.Recycle_UID(u.UID)
}