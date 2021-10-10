package graphics

type Color int

const (
	Undefined Color = iota
	Red
)

func (c Color) String() string {
	switch c {
	case Undefined:
		return "undefined"
	case Red:
		return "#f43f1a"
	}
	return "how did you get here?"
}