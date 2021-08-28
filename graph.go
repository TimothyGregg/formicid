package main

import "errors"

type Vertex struct {
	position [2]int
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

func (g *Graph) Add_Vertex(position [2]int) (*Vertex, error) {
	v := &Vertex{position: position}
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
	return v1.position[0] == v2.position[0] && v1.position[1] == v2.position[1]
}
