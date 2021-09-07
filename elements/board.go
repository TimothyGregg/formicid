package elements

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	. "github.com/TimothyGregg/Antmound/graph"
)

type Gameboard interface {
	update() error
}

type Board struct {
	graph          *Graph
	Nodes          []*Node
	Paths          []*Path
	Size_x         float64
	Size_y         float64
	radius_channel chan int
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

func (n *Node) Get() (int, int, float32, float32, float32) {
	x, y := n.vertex.Get()
	return x, y, n.population_cap, n.generation_rate, n.radius
}

func (p *Path) Get() (*Vertex, *Vertex) {
	return p.edge.Get()
}

func NewBoard() *Board {
	b := &Board{}
	b.graph = &Graph{}
	b.radius_channel = make(chan int)
	fmt.Println("here")
	go b.generate_radii()
	fmt.Println("There")
	return b
}

func (b *Board) generate_radii() {
	for {
		b.radius_channel <- rand.Intn(20)
	}
}

func (b *Board) SetSize(dims [2]float64) error {
	if dims[0] < 1 || dims[1] < 1 {
		return errors.New("Dimensions for a board cannot be less than 1")
	}
	b.Size_x = dims[0]
	b.Size_y = dims[1]
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

func (b *Board) add_node(x, y float64, radius int) error {
	if x < 0 || x > b.Size_x-1 {
		return errors.New("X-position outside board boundaries")
	} else if y < 0 || y > b.Size_y-1 {
		return errors.New("Y-position outside board boundaries")
	}
	v, err := b.graph.Add_Vertex(x, y)
	b.Nodes = append(b.Nodes, &Node{vertex: v, radius: float32(radius)})
	return err
}

func (b *Board) connect_nodes(n1 *Node, n2 *Node) error {
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
	for i := 0; i < 100; i++ {
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
	next_radius := <-b.radius_channel
	good := true
	for _, node := range b.Nodes {
		if node.node_distance(guess_x, guess_y) < float64(node.radius+float32(next_radius)) {
			good = false
			break
		}
	}
	if good {
		b.add_node(guess_x, guess_y, next_radius)
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
			err = b.connect_nodes(b.Nodes[triangulation.Triangles[3*it+jt]], b.Nodes[triangulation.Triangles[3*it+(1+jt)%3]])
			//if err != nil {
			//	fmt.Println(err)
			//}
		}
	}
	return nil
}
