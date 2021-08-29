package main

import (
	"fmt"
	"errors"
	"github.com/fogleman/delaunay"
)

type Vertex struct {
	x, y float64
}

type Edge struct {
	vertices [2]*Vertex
}

type Graph struct {
	vertices []*Vertex
	edges    []*Edge
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
	e := &Edge{vertices: [2]*Vertex{v1, v2}}
	g.edges = append(g.edges, e)
	return e, nil
}

func (v1 *Vertex) same_as(v2 *Vertex) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func (g *Graph) Delaunay() error { // https://mapbox.github.io/delaunator/
	var points []delaunay.Point
	for _, v := range g.vertices {
		points = append(points, delaunay.Point{X: v.x, Y: v.y})
	}
	triangulation, err := delaunay.Triangulate(points)
	fmt.Println(triangulation)
	return err
}
