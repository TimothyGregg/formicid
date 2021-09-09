package graph

import (
	"errors"
	"fmt"
	"math"

	"github.com/fogleman/delaunay"
)

type Vertex struct {
	x, y float64
}

func (v *Vertex) Get() (int, int) {
	return int(v.x), int(v.y)
}

func (v Vertex) String() string {
	return "(" + fmt.Sprint(v.x) + ", " + fmt.Sprint(v.y) + ")"
}

func (v1 *Vertex) Same_As(v2 *Vertex) bool {
	return v1.x == v2.x && v1.y == v2.y
}

type Edge struct {
	v1, v2 *Vertex
	length float64
}

func (e *Edge) Length() float64 {
	return e.length
}

func NewEdge(v1, v2 *Vertex) *Edge {
	e := &Edge{v1: v1, v2: v2}
	e.length = float64(math.Sqrt(math.Pow(v2.x-v1.x, 2) + math.Pow(v2.y-v1.y, 2)))
	return e
}

func (e *Edge) Get() (*Vertex, *Vertex) {
	return e.v1, e.v2
}

func (e Edge) String() string {
	return fmt.Sprint(e.v1) + " to " + fmt.Sprint(e.v2)
}

func (e1 *Edge) same_as(e2 *Edge) bool {
	return (e1.v1 == e2.v1 && e1.v2 == e2.v2) || (e1.v1 == e2.v2 && e1.v2 == e2.v1)
}

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

func (g Graph) String() string {
	outstr := ""
	for i, vertex := range g.Vertices {
		outstr = outstr + "[" + fmt.Sprint(i) + "]: " + vertex.String() + "\n"
	}
	for _, edge := range g.Edges {
		outstr = outstr + edge.String() + "; "
	}
	return outstr
}

func (g *Graph) has(v *Vertex) bool {
	for _, v_test := range g.Vertices {
		if v == v_test {
			return true
		}
	}
	return false
}

func (g *Graph) Add_Vertex(x, y float64) (*Vertex, error) {
	v := &Vertex{x: x, y: y}
	for _, v_test := range g.Vertices {
		if v.Same_As(v_test) {
			return v_test, errors.New("Vertex already exists")
		}
	}
	g.Vertices = append(g.Vertices, v)
	return v, nil
}

func (g *Graph) Add_Edge(v1 *Vertex, v2 *Vertex) (*Edge, error) {
	if !g.has(v1) || !g.has(v2) {
		return nil, errors.New("One or more vertices do not exist in the graph")
	}
	e := NewEdge(v1, v2)
	for _, e_test := range g.Edges {
		if e.same_as(e_test) {
			return e_test, errors.New("Edge already exists")
		}
	}
	g.Edges = append(g.Edges, e)
	return e, nil
}

func (g *Graph) Remove_Edge(e *Edge) error {
	for it, edge := range g.Edges {
		if e.same_as(edge) {
			g.Edges = append(g.Edges[:it], g.Edges[it+1:]...)
			return nil
		}
	}
	return errors.New("Edge not found")
}

/*
func (g *Graph) Get_Edge(v1, v2 *Vertex) (*Edge, error) {
	test_edge := NewEdge(v1, v2)
	for _, e := range g.Edges {
		if e.same_as(test_edge) {
			return e, nil
		}
	}
	return nil, errors.New("Edge not found")
}
*/

func (g *Graph) Delaunay_Triangulate() (*delaunay.Triangulation, error) {
	var points []delaunay.Point
	for _, v := range g.Vertices {
		points = append(points, delaunay.Point{X: v.x, Y: v.y})
	}
	return delaunay.Triangulate(points)
}

func (g *Graph) Connect_Delaunay() error { // https://mapbox.github.io/delaunator/
	triangulation, err := g.Delaunay_Triangulate()
	for it := 0; it < len(triangulation.Triangles)/3; it++ {
		for jt := 0; jt < 3; jt++ {
			g.Add_Edge(g.Vertices[triangulation.Triangles[3*it+jt]], g.Vertices[triangulation.Triangles[3*it+(jt+1)%3]])
		}
	}
	return err
}
