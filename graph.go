package main

import (
	"errors"
	"fmt"

	"github.com/fogleman/delaunay"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) String() string {
	return "(" + fmt.Sprint(v.x) + ", " + fmt.Sprint(v.y) + ")"
}

func (v1 *Vertex) same_as(v2 *Vertex) bool {
	return v1.x == v2.x && v1.y == v2.y
}

type Edge struct {
	v1, v2 *Vertex
}

func (e Edge) String() string {
	return fmt.Sprint(e.v1) + " to " + fmt.Sprint(e.v2)
}

func (e1 *Edge) same_as(e2 *Edge) bool {
	return (e1.v1 == e2.v1 && e1.v2 == e2.v2) || (e1.v1 == e2.v2 && e1.v2 == e2.v1)
}

type Graph struct {
	vertices []*Vertex
	edges    []*Edge
}

func (g Graph) String() string {
	outstr := ""
	for i, vertex := range g.vertices {
		outstr = outstr + "[" + fmt.Sprint(i) + "]: " + vertex.String() + "\n"
	}
	for _, edge:= range g.edges {
		outstr = outstr + edge.String() + "; "
	}
	return outstr
}

func (g *Graph) has(v *Vertex) bool {
	for _, v_test := range g.vertices {
		if v == v_test {
			return true
		}
	}
	return false
}

func (g *Graph) Add_Vertex(x, y float64) (*Vertex, error) {
	v := &Vertex{x: x, y: y}
	for _, v_test := range g.vertices {
		if v.same_as(v_test) {
			return nil, errors.New("Vertex already exists")
		}
	}
	g.vertices = append(g.vertices, v)
	return v, nil
}

func (g *Graph) Add_Edge(v1 *Vertex, v2 *Vertex) (*Edge, error) {
	if !g.has(v1) || !g.has(v2) {
		return nil, errors.New("One or more vertices do not exist in the graph")
	}
	e := &Edge{v1: v1, v2: v2}
	for _, e_test := range g.edges {
		if e.same_as(e_test) {
			return nil, errors.New("Edge already exists")
		}
	}
	g.edges = append(g.edges, e)
	return e, nil
}

func (g *Graph) Delaunay_Triangulate() (*delaunay.Triangulation, error) {
	var points []delaunay.Point
	for _, v := range g.vertices {
		points = append(points, delaunay.Point{X: v.x, Y: v.y})
	}
	return delaunay.Triangulate(points)
}

func (g *Graph) Connect_Delaunay() error { // https://mapbox.github.io/delaunator/
	triangulation, err := g.Delaunay_Triangulate()
	for it := 1; it <= len(triangulation.Triangles) / 3; it++ {
		for jt := 0; jt < 3; jt++ {
			g.Add_Edge(g.vertices[triangulation.Triangles[3*it-3 + jt]], g.vertices[triangulation.Triangles[(3*it-2 + jt) % 3]])
		}
	}
	return err
}
