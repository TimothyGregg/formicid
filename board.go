package main

import (
	"fmt"
	"errors"
	"math"
	"math/rand"
)

type board struct {
	graph       *Graph
	nodes       []*node
	paths       []*path
	size_x      int
	size_y      int
	node_radius int
}

func (b board) String() string {
	outstr := ""
	for _, node := range b.nodes {
		outstr = outstr + node.String() + ", "
	}
	return outstr[:len(outstr) - 2]
}

type node struct {
	vertex *Vertex
}

func (n node) String() string {
	return "[" + fmt.Sprint(n.vertex.position[0]) + ", " + fmt.Sprint(n.vertex.position[1]) + "]"
}

type path struct {
	edge *Edge
}

func NewBoard() *board {
	b := &board{}
	b.graph = &Graph{}
	return b
}

func (b *board) SetSize(dims [2]int) error {
	if dims[0] < 1 || dims[1] < 1 {
		return errors.New("Dimensions for a board cannot be less than 1")
	}
	b.size_x = dims[0]
	b.size_y = dims[1]
	return nil
}

func (b *board) Set_Node_Radius(val int) error {
	if val <= 0 {
		return errors.New("Node radius must be greater than 0")
	}
	b.node_radius = val
	return nil
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
	if position[0] < 0 || position[0] > b.size_x-1 {
		return errors.New("X-position outside board boundaries")
	} else if position[1] < 0 || position[1] > b.size_y-1 {
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

func (node *node) node_distance(x int, y int) float64 {
	val := math.Pow(float64(node.vertex.position[0]-x), 2)
	return math.Sqrt(val + math.Pow(float64(node.vertex.position[1]-y), 2))
}

func (b *board) Naive_Fill() error {
	for i := 0; i < 10; i++ {
		for {
			if !b.add_random_node() {
				break
			}
		}
	}
	return nil
}

func (b *board) add_random_node() bool {
	guess_x := rand.Intn(b.size_x)
	guess_y := rand.Intn(b.size_y)
	good := true
	for _, node := range b.nodes {
		if node.node_distance(guess_x, guess_y) < float64(2*b.node_radius) {
			good = false
			break
		}
	}
	if good {
		b.Add_Node([2]int{guess_x, guess_y})
	}
	return good
}
