package graphics

type Color int

const (
	Undefined Color = iota
	RED
	DARK_ORCHID
	DARK_ORANGE
	DEEP_PINK
	KHAKI
)

func (c Color) String() string { // https://www.color-hex.com/color-names.html
	switch c {
	case Undefined:
		return "undefined"
	case RED:
		return "#f43f1a"
	case DARK_ORCHID:
		return "#9932cc"
	case DARK_ORANGE:
		return "#ff8c00"
	case DEEP_PINK:
		return "#ee1289"
	case KHAKI:
		return "cdc673"
	}
	return "how did you get here?"
}