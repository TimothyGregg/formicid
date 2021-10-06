package elements

import (
	"errors"
	"fmt"
	"math/rand"

	graph "github.com/TimothyGregg/Antmound/graph"
	tools "github.com/TimothyGregg/Antmound/tools"
)

type Board struct {
	Element
	graph              *graph.Graph
	Nodes              map[int]*Node
	Paths              map[int]*Path
	Size_x             float64
	Size_y             float64
	radius_channel     chan int
	node_uid_generator *tools.UID_Generator
	edge_uid_generator *tools.UID_Generator
}

func (b *Board) Update() error {
	b.Element.tick()
	for _, n := range b.Nodes {
		n.update()
	}
	for _, e := range b.Paths {
		e.update()
	}
	return nil
}

func (b *Board) Get_Size() [2]float64 {
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

func New_Board() *Board {
	b := &Board{}
	b.graph = graph.NewGraph()
	b.Nodes = make(map[int]*Node)
	b.Paths = make(map[int]*Path)
	b.radius_channel = make(chan int)
	b.node_uid_generator = tools.New_UID_Generator()
	b.edge_uid_generator = tools.New_UID_Generator()
	go b.radii_generation()

	return b
}

func (b *Board) radii_generation() {
	for {
		b.radius_channel <- 10 // rand.Intn(10) + 10
	}
}

func (b *Board) SetSize(dims [2]float64) error {
	if dims[0] < 1 || dims[1] < 1 {
		return errors.New("dimensions for a board cannot be less than 1")
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
		return errors.New("x-position outside board boundaries")
	} else if y < 0 || y > b.Size_y-1 {
		return errors.New("y-position outside board boundaries")
	}
	v, err := b.graph.Add_Vertex(x, y)
	next_uid := b.node_uid_generator.Next()
	b.Nodes[next_uid] = &Node{vertex: v, radius: float64(radius), UID: next_uid}
	return err
}

func (b *Board) connect_nodes(n1 *Node, n2 *Node) error {
	if !b.has(n1) || !b.has(n2) {
		return errors.New("one or more nodes do not exist on the board")
	}
	e, err := b.graph.Add_Edge(n1.vertex, n2.vertex)
	_, ok := err.(*graph.EdgeAlreadyExistsError)
	if err != nil && !ok {
		return err
	}
	next_uid := b.edge_uid_generator.Next()
	b.Paths[next_uid] = &Path{edge: e, UID: next_uid}
	return err
}

func (b *Board) disconnect_path(uid int) error {
	for _, p_test := range b.Paths {
		if p_test.UID == uid {
			b.graph.Remove_Edge(p_test.edge)
			delete(b.Paths, uid)
			return nil
		}
	}
	return errors.New("Path not found")
}

func (b *Board) find_node(v *graph.Vertex) *Node {
	for _, n := range b.Nodes {
		if n.vertex.Same_As(v) {
			return n
		}
	}
	return nil
}

func (b *Board) Naive_Fill(tries int) error {
	for i := 0; i < tries; i++ {
		for {
			if !b.add_random_node() {
				break
			}
		}
	}
	return nil
}

func (b *Board) add_random_node() bool {
	guess_x := float64(rand.Intn(int(b.Size_x)))
	guess_y := float64(rand.Intn(int(b.Size_y)))
	next_radius := <-b.radius_channel
	good := true
	for _, node := range b.Nodes {
		if node.node_distance(guess_x, guess_y) < float64(node.radius+float64(next_radius)) {
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
	b.graph.Connect_Delaunay()
	avg := 0.0
	for it, e := range b.graph.Edges {
		v1, v2 := e.Get()
		b.connect_nodes(b.find_node(v1), b.find_node(v2))
		if it > 0 {
			avg = avg*(float64(it)-1)/float64(it) + e.Length()/float64(it)
		} else {
			avg = e.Length()
		}
	}
	var to_disconnect []int
	for _, p := range b.Paths {
		if p.edge.Length() > 2.5*avg {
			to_disconnect = append(to_disconnect, p.UID)
		}
	}
	for _, uid := range to_disconnect {
		b.disconnect_path(uid)
	}
	return nil
}
