package main

import (
	"errors"
)

type board struct {
	graph *Graph
	nodes []*node
	paths []*path
	size_x int
	size_y int
}

type node struct {
	vertex *Vertex
}

type path struct {
	edge *Edge
}

func NewBoard() *board {
	b := &board{}
	b.graph = &Graph{}
	return b
}

func (b *board) has(n *node) bool {
	for _, n_test := range b.nodes {
		if n == n_test {
			return true
		}
	}
	return false
}

func (b *board) Add_Node(position [2]int) error {
	if position[0] < 0 || position[0] > b.size_x - 1 {
		return errors.New("X-position outside board boundaries")
	} else if position[1] < 0 || position[1] > b.size_y - 1 {
		return errors.New("Y-position outside board boundaries")
	}
	v, err := b.graph.Add_Vertex(position)
	b.nodes = append(b.nodes, &node{vertex: v})
	return err
}

func (b *board) Connect_Nodes(n1 *node, n2 *node) error {
	if !b.has(n1) || !b.has(n2) {
		return errors.New("One or more nodes do not exist on the board")
	}
	e, err := b.graph.Add_Edge(n1.vertex, n2.vertex)
	if err != nil {
		return err
	}
	b.paths = append(b.paths, &path{edge: e})
	return err
}

func (b *board) Nice_Fill() error {
	return nil
}