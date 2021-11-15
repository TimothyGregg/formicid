package graph

import "fmt"

type Vertex struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (v *Vertex) Position() (int, int) {
	return v.X, v.Y
}

func (v Vertex) String() string {
	return "(" + fmt.Sprint(v.X) + ", " + fmt.Sprint(v.Y) + ")"
}

func (v1 *Vertex) Same_As(v2 *Vertex) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}
