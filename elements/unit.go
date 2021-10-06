package elements

import "errors"

type Unit struct {
	Element
	UID        int
	population int
	Team       Team
}

func (n *Node) New_Unit(val int) (*Unit, error) {
	if val <= 0 {
		return nil, errors.New("unit created with power less than 1")
	}
	u := &Unit{population: val}
	u.Team = n.Team
	u.UID = u.Team.Unit_UID_Generator.Next()
	return u, nil
}

func (u *Unit) Destroy() {
	u.Team.Unit_UID_Generator.Recycle(u.UID)
}
