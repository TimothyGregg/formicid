package elements

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/TimothyGregg/formicid/game/util/graph"
	"github.com/TimothyGregg/formicid/game/util/uid"
)

type Board struct {
	Element
	Nodes           map[*uid.UID]*Node `json:"-"`
	Paths           map[*uid.UID]*Path `json:"-"`
	NodeConnections map[int][]*Node    `json:"-"`
	Size            [2]int             `json:"size"`
}

func New_Board(board_uid *uid.UID, size_x, size_y int) (*Board, error) {
	b := &Board{}
	b.UID = board_uid
	b.Size = [2]int{size_x, size_y}
	g := *graph.New_Graph()
	b.Nodes = make(map[*uid.UID]*Node)
	b.Paths = make(map[*uid.UID]*Path)

	b.NodeConnections = make(map[int][]*Node)

	node_uid_generator := uid.New_UID_Generator()
	edge_uid_generator := uid.New_UID_Generator()

	// Radius generation
	radius_channel := make(chan int)
	go func() {
		for {
			radius_channel <- 10
		}
	}()

	// Generate Nodes
	node_vertices := make(map[*graph.Vertex]*Node)
	add_node := func(x, y, radius int) error {
		if x < 0 || x > size_x-1 {
			return errors.New("x-position outside board boundaries")
		} else if y < 0 || y > size_y-1 {
			return errors.New("y-position outside board boundaries")
		}
		v, err := g.Add_Vertex(x, y)
		if err != nil {
			return err
		}
		next_uid := node_uid_generator.Next()
		b.Nodes[next_uid] = NewNode(next_uid, x, y, <-radius_channel)
		node_vertices[v] = b.Nodes[next_uid]
		return nil
	}

	fill_tries := 100
	for i := 0; i < fill_tries; i++ {
		for {
			guess_x := rand.Intn(int(size_x))
			guess_y := rand.Intn(int(size_y))
			next_radius := <-radius_channel
			good := true
			for _, node := range b.Nodes {
				if func(node *Node, x, y int) float64 {
					val := math.Pow(float64(node.X-x), 2)
					return math.Sqrt(val + math.Pow(float64(node.Y-y), 2))
				}(node, guess_x, guess_y) < float64(node.Radius+next_radius) {
					good = false
					break
				}
			}
			if good {
				add_node(guess_x, guess_y, next_radius)
			} else {
				break
			}
		}
	}

	// Connect Nodes
	node_from_vertex := func(v *graph.Vertex) (*Node, error) {
		n, ok := node_vertices[v]
		if !ok {
			return nil, errors.New("vertex not in node-vertex map")
		}
		return n, nil
	}

	g.Connect_Delaunay()
	avg := 0.0
	for it, e := range g.Edges {
		v1, v2 := e.Vertices()
		n1, _ := node_from_vertex(v1)
		n2, _ := node_from_vertex(v2)
		// Make connection
		e, err := g.Add_Edge(v1, v2)                 // This line WILL error quite often
		_, ok := err.(*graph.EdgeAlreadyExistsError) // This checks to make sure it's the correct kind of error
		if err != nil && !ok {
			return nil, err
		}
		func(n1, n2 *Node) error {
			next_uid := edge_uid_generator.Next()
			b.Paths[next_uid] = New_Path(next_uid, n1, n2)
			b.NodeConnections[n1.UID.Value()] = append(b.NodeConnections[n1.UID.Value()], n2)
			b.NodeConnections[n2.UID.Value()] = append(b.NodeConnections[n2.UID.Value()], n1)
			return err
		}(n1, n2)
		if it > 0 {
			avg = avg*(float64(it)-1)/float64(it) + e.Length()/float64(it)
		} else {
			avg = e.Length()
		}
	}
	// Prune unwanted connections
	var to_disconnect []*uid.UID
	for _, p := range b.Paths {
		if p.Length > 2.5*avg {
			to_disconnect = append(to_disconnect, p.UID)
		}
	}
	for _, uid_to_be_rid_of := range to_disconnect {
		func(uid *uid.UID) error {
			_, ok := b.Paths[uid_to_be_rid_of]
			if !ok {
				return errors.New("Path not found")
			}
			// Remove connections from the adjacency map
			n1, n2 := b.Paths[uid_to_be_rid_of].Nodes[0], b.Paths[uid_to_be_rid_of].Nodes[1]
			for i, n_test := range b.NodeConnections[n1] {
				if n_test.UID.Value() == n2 {
					b.NodeConnections[n1] = append(b.NodeConnections[n1][:i], b.NodeConnections[n1][i+1:]...)
					break
				}
			}
			for i, n_test := range b.NodeConnections[n2] {
				if n_test.UID.Value() == n1 {
					b.NodeConnections[n2] = append(b.NodeConnections[n2][:i], b.NodeConnections[n2][i+1:]...)
					break
				}
			}
			// Delete Path
			delete(b.Paths, uid_to_be_rid_of)
			return nil
		}(uid_to_be_rid_of)
	}

	return b, nil
}

func (b *Board) NodeByID(id int) (*Node, error) {
	for uid, node := range b.Nodes {
		if id == uid.Value() {
			return node, nil
		}
	}
	return nil, errors.New("no node found with id " + fmt.Sprint(id))
}

func (b Board) String() string {
	outstr := fmt.Sprintf("Board %d:", b.UID.Value())
	for _, n := range b.Nodes {
		outstr += n.String() + ", "
	}
	outstr = outstr[:len(outstr)-2]
	for _, p := range b.Paths {
		outstr += p.String() + ", "
	}
	return outstr[:len(outstr)-2]
}

func (b *Board) MarshalJSON() ([]byte, error) {
	type Alias Board
	node_array := make([]int, 0, len(b.Nodes))
	for id := range b.Nodes {
		node_array = append(node_array, id.Value())
	}
	path_array := make([]int, 0, len(b.Paths))
	for id := range b.Paths {
		path_array = append(path_array, id.Value())
	}
	return json.Marshal(&struct {
		Nodes []int `json:"nodes"`
		Paths []int `json:"paths"`
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
