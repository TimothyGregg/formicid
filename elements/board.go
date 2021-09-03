package elements

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	. "github.com/TimothyGregg/Antmound/graph"
)

type gameboard interface {
	update() error
}

type Board struct {
	graph       *Graph
	Nodes       []*Node
	Paths       []*Path
	Size_x      float64
	Size_y      float64
	node_radius int
}

type Node struct {
	vertex          *Vertex
	population_cap  float32
	generation_rate float32
	radius          float32
}

type Path struct {
	edge *Edge
}

func (b *Board) GetSize() [2]float64 {
	return [2]float64{b.Size_x, b.Size_y}
}

func (b Board) String() string {
	outstr := ""
	for i, node := range b.Nodes {
		outstr = outstr + "[" + fmt.Sprint(i) + "]: " + node.String() + "\n"
	}
	for _, path := range b.Paths {
		var n1, n2 int
		v1, v2 := path.edge.Get()
		for i, node := range b.Nodes {
			if v1 == node.vertex {
				n1 = i
			} else if v2 == node.vertex {
				n2 = i
			}
		}
		outstr = outstr + fmt.Sprint(n1) + " to " + fmt.Sprint(n2) + "; "
	}
	return outstr
}

func (n Node) String() string {
	return n.vertex.String()
}

func (p Path) String() string {
	return p.edge.String()
}

func NewBoard() *Board {
	b := &Board{}
	b.graph = &Graph{}
	return b
}

func (b *Board) SetSize(dims [2]float64) error {
	if dims[0] < 1 || dims[1] < 1 {
		return errors.New("Dimensions for a board cannot be less than 1")
	}
	b.Size_x = dims[0]
	b.Size_y = dims[1]
	return nil
}

func (b *Board) Set_Node_Radius(val int) error {
	if val <= 0 {
		return errors.New("Node radius must be greater than 0")
	}
	b.node_radius = val
	return nil
}

func (b *Board) has(n *Node) bool {
	for _, n_test := range b.Nodes {
		if n == n_test {
			return true
		}
	}
	return false
}

func (b *Board) Add_Node(x, y float64) error {
	if x < 0 || x > b.Size_x-1 {
		return errors.New("X-position outside board boundaries")
	} else if y < 0 || y > b.Size_y-1 {
		return errors.New("Y-position outside board boundaries")
	}
	v, err := b.graph.Add_Vertex(x, y)
	b.Nodes = append(b.Nodes, &Node{vertex: v})
	return err
}

func (b *Board) Connect_Nodes(n1 *Node, n2 *Node) error {
	if !b.has(n1) || !b.has(n2) {
		return errors.New("One or more nodes do not exist on the board")
	}
	e, err := b.graph.Add_Edge(n1.vertex, n2.vertex)
	if err != nil {
		return err
	}
	b.Paths = append(b.Paths, &Path{edge: e})
	return err
}

func (node *Node) node_distance(x, y float64) float64 {
	nx, ny := node.vertex.Get()
	val := math.Pow(float64(nx)-x, 2)
	return math.Sqrt(val + math.Pow(float64(ny)-y, 2))
}

func (b *Board) Naive_Fill() error {
	for i := 0; i < 10; i++ {
		for {
			if !b.add_random_node() {
				break
			}
		}
	}
	fmt.Println(len(b.Nodes))
	return nil
}

func (b *Board) add_random_node() bool {
	guess_x := float64(rand.Intn(int(b.Size_x)))
	guess_y := float64(rand.Intn(int(b.Size_y)))
	good := true
	for _, node := range b.Nodes {
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

func (b *Board) Connect_Delaunay() error {
	triangulation, err := b.graph.Delaunay_Triangulate()
	if err != nil {
		return err
	}
	// fmt.Println(triangulation.Triangles)
	for it := 0; it < len(triangulation.Triangles)/3; it++ {
		for jt := 0; jt < 3; jt++ {
			// fmt.Println(fmt.Sprint(triangulation.Triangles[3*it + jt]) + "-" + fmt.Sprint(triangulation.Triangles[3*it + (1 + jt) % 3]))
			err = b.Connect_Nodes(b.Nodes[triangulation.Triangles[3*it+jt]], b.Nodes[triangulation.Triangles[3*it+(1+jt)%3]])
			//if err != nil {
			//	fmt.Println(err)
			//}
		}
	}
	return nil
}
