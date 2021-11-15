package graph

import (
	"fmt"
	"math"
)

type Edge struct {
	v1, v2 *Vertex
	length float64
}

func (e *Edge) Length() float64 {
	return e.length
}

func NewEdge(v1, v2 *Vertex) *Edge {
	e := &Edge{v1: v1, v2: v2}
	e.length = float64(math.Sqrt(math.Pow(float64(v2.X-v1.X), 2) + math.Pow(float64(v2.Y-v1.Y), 2)))
	return e
}

func (e *Edge) Vertices() (*Vertex, *Vertex) {
	return e.v1, e.v2
}

func (e Edge) String() string {
	return fmt.Sprint(e.v1) + " to " + fmt.Sprint(e.v2)
}

func (e1 *Edge) same_as(e2 *Edge) bool {
	return (e1.v1 == e2.v1 && e1.v2 == e2.v2) || (e1.v1 == e2.v2 && e1.v2 == e2.v1)
}
