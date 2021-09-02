package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

//Temp
func (b *board) Get_Nodes() []*node {
	return b.nodes
}

//Temp
func (b *board) Get_Paths() []*path {
	return b.paths
}

type gameboard interface {
	update() error
}

type board struct {
	graph       *Graph
	nodes       []*node
	paths       []*path
	size_x      float64
	size_y      float64
	node_radius int
}

func (b *board) GetSize() [2]float64 {
	return [2]float64{b.size_x, b.size_y}
}

func (b board) String() string {
	outstr := ""
	for i, node := range b.nodes {
		outstr = outstr + "[" + fmt.Sprint(i) + "]: " + node.String() + "\n"
	}
	for _, path := range b.paths {
		var n1, n2 int
		for i, node := range b.nodes {
			if path.edge.v1 == node.vertex {
				n1 = i
			} else if path.edge.v2 == node.vertex {
				n2 = i
			}
		}
		outstr = outstr + fmt.Sprint(n1) + " to " + fmt.Sprint(n2) + "; "
	}
	return outstr
}

type node struct {
	vertex *Vertex
}

//Temp
func (n *node) Get() (int, int) {
	return n.vertex.Get()
}

func (n node) String() string {
	return n.vertex.String()
}

type path struct {
	edge *Edge
}

//Temp
func (p *path) Get() (*Vertex, *Vertex) {
	return p.edge.Get()
}

func (p path) String() string {
	return p.edge.String()
}

func NewBoard() *board {
	b := &board{}
	b.graph = &Graph{}
	return b
}

func (b *board) SetSize(dims [2]float64) error {
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

func (b *board) Add_Node(x, y float64) error {
	if x < 0 || x > b.size_x-1 {
		return errors.New("X-position outside board boundaries")
	} else if y < 0 || y > b.size_y-1 {
		return errors.New("Y-position outside board boundaries")
	}
	v, err := b.graph.Add_Vertex(x, y)
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

func (node *node) node_distance(x, y float64) float64 {
	val := math.Pow(float64(node.vertex.x-x), 2)
	return math.Sqrt(val + math.Pow(float64(node.vertex.y-y), 2))
}

func (b *board) Naive_Fill() error {
	for i := 0; i < 10; i++ {
		for {
			if !b.add_random_node() {
				break
			}
		}
	}
	fmt.Println(len(b.nodes))
	return nil
}

func (b *board) add_random_node() bool {
	guess_x := float64(rand.Intn(int(b.size_x)))
	guess_y := float64(rand.Intn(int(b.size_y)))
	good := true
	for _, node := range b.nodes {
		if node.node_distance(guess_x, guess_y) < float64(2*b.node_radius) {
			good = false
			break
		}
	}
	if good {
		b.Add_Node(guess_x, guess_y)
	}
	return good
}

func (b *board) Connect_Delaunay() error {
	triangulation, err := b.graph.Delaunay_Triangulate()
	if err != nil {
		return err
	}
	fmt.Println(triangulation.Triangles)
	for it := 0; it < len(triangulation.Triangles)/3; it++ {
		for jt := 0; jt < 3; jt++ {
			// fmt.Println(fmt.Sprint(triangulation.Triangles[3*it + jt]) + "-" + fmt.Sprint(triangulation.Triangles[3*it + (1 + jt) % 3]))
			err = b.Connect_Nodes(b.nodes[triangulation.Triangles[3*it+jt]], b.nodes[triangulation.Triangles[3*it+(1+jt)%3]])
			//if err != nil {
			//	fmt.Println(err)
			//}
		}
	}
	return nil
}
