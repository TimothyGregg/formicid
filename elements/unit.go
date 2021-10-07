package elements

type Unit struct {
	Element
	UID        int
	population int
	Team       Team
}

func (u *Unit) Destroy() {
	u.Team.Unit_UID_Generator.Recycle(u.UID)
}