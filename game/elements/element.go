package elements

type Element struct {
	UID   int `json:"e_uid"`
	timer uint8
}

func (e *Element) New(uid int) {
	e.UID = uid
}

func (e *Element) tick() {
	e.timer = (e.timer + 1) % 60
}

func (e *Element) update() {
	e.tick()
}

func (e *Element) Timer() uint8 {
	return e.timer
}
