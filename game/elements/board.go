package elements

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	util "github.com/TimothyGregg/formicid/game/util"
	graph "github.com/TimothyGregg/formicid/game/util/graph"
)

type Board struct {
	Element
	graph.Graph
	Nodes              map[int]*Node `json:"-"`
	Paths              map[int]*Path `json:"-"`
	node_connections   map[*Node][]*Node
	Size_x             int `json:"size_x"`
	Size_y             int `json:"size_y"`
	radius_channel     chan int
	node_uid_generator *util.UID_Generator
	edge_uid_generator *util.UID_Generator
}

func (b *Board) MarshalJSON() ([]byte, error) {
	type Alias Board
	node_array := make([]*Node, 0, len(b.Nodes))
	for _, node := range b.Nodes {
		node_array = append(node_array, node)
	}
	path_array := make([]*Path, 0, len(b.Paths))
	for _, path := range b.Paths {
		path_array = append(path_array, path)
	}
	return json.Marshal(&struct {
		Nodes []*Node `json:"nodes"`
		Paths []*Path `json:"paths"`
		*Alias
	}{
		Nodes: node_array,
		Paths: path_array,
		Alias: (*Alias)(b),
	})
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

func (b *Board) Get_Size() [2]int {
	return [2]int{b.Size_x, b.Size_y}
}

// Temp
func (b *Board) Get_node_connections() map[*Node][]*Node {
	return b.node_connections
}

func (b Board) String() string {
	outstr := ""
	for i, node := range b.Nodes {
		outstr = outstr + "[" + fmt.Sprint(i) + "]: " + node.String() + "\n"
	}
	for _, path := range b.Paths {
		var n1, n2 int
		v1, v2 := path.Vertices()
		for i, node := range b.Nodes {
			if v1 == node {
				n1 = i
			} else if v2 == node {
				n2 = i
			}
		}
		outstr = outstr + fmt.Sprint(n1) + " to " + fmt.Sprint(n2) + "; "
	}
	return outstr
}

func New_Board() *Board {
	b := &Board{}
	b.Graph = *graph.New_Graph()
	b.Element.New(0)
	b.Nodes = make(map[int]*Node)
	b.Paths = make(map[int]*Path)
	b.node_connections = map[*Node][]*Node{}
	b.radius_channel = make(chan int)
	b.node_uid_generator = util.New_UID_Generator()
	b.edge_uid_generator = util.New_UID_Generator()
	go b.init_radii_generation()
	return b
}

func (b *Board) init_radii_generation() {
	for {
		b.radius_channel <- 10 // rand.Intn(10) + 10
	}
}

func (b *Board) Set_Size(dims [2]int) error {
	if dims[0] < 1 || dims[1] < 1 {
		return errors.New("dimensions for a board cannot be less than 1")
	}
	b.Size_x = dims[0]
	b.Size_y = dims[1]
	return nil
}

func (b *Board) has(n *Node) bool {
	_, ok := b.Nodes[n.UID]
	return ok
}

func (b *Board) add_node(x, y int, radius int) error {
	if x < 0 || x > b.Size_x-1 {
		return errors.New("x-position outside board boundaries")
	} else if y < 0 || y > b.Size_y-1 {
		return errors.New("y-position outside board boundaries")
	}
	_, err := b.Add_Vertex(x, y) // This is gonna be wrong
	if err != nil {
		return err
	}
	next_uid := b.node_uid_generator.Next()
	b.Nodes[next_uid] = New_Node(next_uid, x, y, <-b.radius_channel)
	return nil
}

func (b *Board) connect_nodes(n1 *Node, n2 *Node) error {
	if !b.has(n1) || !b.has(n2) {
		return errors.New("one or more nodes do not exist on the board")
	}
	already, err := b.check_connected(n1, n2)
	if err != nil {
		return err
	} else if already {
		return errors.New("connection already exists on the board")
	}
	e, err := b.graph.Add_Edge(n1.vertex, n2.vertex) // This line WILL error quite often
	_, ok := err.(*graph.EdgeAlreadyExistsError)     // This checks to make sure it's the correct kind of error
	if err != nil && !ok {
		return err
	}
	next_uid := b.edge_uid_generator.Next()
	b.Paths[next_uid] = New_Path(next_uid, e)
	b.node_connections[n1] = append(b.node_connections[n1], n2)
	b.node_connections[n2] = append(b.node_connections[n2], n1)
	return err
}

func (b *Board) disconnect_path(uid int) error {
	_, ok := b.Paths[uid]
	if !ok {
		return errors.New("Path not found")
	}
	// Ensure the nodes on the path are connected
	v1, v2 := b.Paths[uid].Vertices()
	n1, err := b.get_node_from_vertex(v1)
	if err != nil {
		return err
	}
	n2, err := b.get_node_from_vertex(v2)
	if err != nil {
		return err
	}
	already, err := b.check_connected(n1, n2)
	if err != nil {
		return err
	} else if !already {
		return errors.New("cannot disconnect non-connected nodes")
	}
	// Remove connections from the adjacency map
	for i, n_test := range b.node_connections[n1] {
		if n_test == n2 {
			b.node_connections[n1] = append(b.node_connections[n1][:i], b.node_connections[n1][i+1:]...)
			break
		}
	}
	for i, n_test := range b.node_connections[n2] {
		if n_test == n1 {
			b.node_connections[n2] = append(b.node_connections[n2][:i], b.node_connections[n2][i+1:]...)
			break
		}
	}
	// Delete Path
	delete(b.Paths, uid)
	return nil
}

func (b *Board) check_connected(n1 *Node, n2 *Node) (bool, error) {
	n1c := b.node_connections[n1]
	n2c := b.node_connections[n2]
	n1ton2, n2ton1 := false, false
	for _, n_test := range n2c {
		if n1 == n_test {
			n2ton1 = true
		}
	}
	for _, n_test := range n1c {
		if n2 == n_test {
			n1ton2 = true
		}
	}
	if n1ton2 == n2ton1 {
		return n1ton2, nil
	}
	return false, errors.New("disagreement in connection lists")
}

func (b *Board) get_node_from_vertex(v *graph.Vertex) (*Node, error) {
	n, ok := b.node_vertices[v]
	if !ok {
		return nil, errors.New("vertex not in node-vertex map")
	}
	return n, nil
}

func (b *Board) Fill() error {
	err := b.naive_fill(100)
	if err != nil {
		return err
	}
	return nil
}

func (b *Board) naive_fill(tries int) error {
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
	guess_x := rand.Intn(int(b.Size_x))
	guess_y := rand.Intn(int(b.Size_y))
	next_radius := <-b.radius_channel
	good := true
	for _, node := range b.Nodes {
		if node.node_distance(guess_x, guess_y) < float64(node.Radius+next_radius) {
			good = false
			break
		}
	}
	if good {
		b.add_node(guess_x, guess_y, next_radius)
	}
	return good
}

func (b *Board) Connect() error {
	err := b.connect_delaunay()
	if err != nil {
		return err
	}
	return nil
}

func (b *Board) connect_delaunay() error {
	b.graph.Connect_Delaunay()
	avg := 0.0
	for it, e := range b.graph.Edges {
		v1, v2 := e.Vertices()
		n1, err := b.get_node_from_vertex(v1)
		if err != nil {
			return err
		}
		n2, err := b.get_node_from_vertex(v2)
		if err != nil {
			return err
		}
		b.connect_nodes(n1, n2)
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
