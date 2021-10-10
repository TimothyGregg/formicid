package elements

import (
	"encoding/json"
	"errors"
	"math"

	graph "github.com/TimothyGregg/Antmound/game/graph"
)

type Node struct {
	Element
	vertex          *graph.Vertex
	Radius          int `json:"radius"`
	Population      int `json:"popuation"`
	Team            *Team `json:"-"`
}

func (n *Node) MarshalJSON() ([]byte, error) {
	type Alias Node
	x_pos, y_pos, _ := n.Get()
	var team int
	if n.Team == nil {
		team = -1
	} else {
		team = n.Team.UID
	}
	return json.Marshal(&struct {
		Team int `json:"teamID"`
		Position [2]int `json:"position"`
		*Alias
	}{
		Team: team,
		Position: [2]int{x_pos, y_pos},
		Alias: (*Alias)(n),
	})
}

func New_Node(uid, radius int, v *graph.Vertex) *Node {
	n := &Node{vertex: v, Radius: radius}
	n.New(uid)
	return n
}

func (n Node) String() string {
	return n.vertex.String()
}

func (n *Node) Get() (int, int, int) {
	x, y := n.vertex.Position()
	return x, y, n.Radius
}

func (node *Node) node_distance(x, y int) float64 {
	nx, ny := node.vertex.Position()
	val := math.Pow(float64(nx-x), 2)
	return math.Sqrt(val + math.Pow(float64(ny-y), 2))
}

func (n *Node) change_population(val int) (int, error) {
	rem := -1 * (n.Population + val)
	n.Population += val
	if n.Population < 0 {
		n.Population = 0
		return rem, nil
	}
	return 0, nil
}

func (n *Node) New_Unit(val int) (*Unit, error) {
	if val > n.Population {
		return nil, errors.New("cannot create a unit with more population than the node")
	} else if val <= 0 {
		return nil, errors.New("unit created with power less than 1")
	}
	u := New_Unit(n.Team.Next_Unit_UID(), val)
	u.Team = n.Team
	n.change_population(-1 * val)
	return u, nil
}

/*
func (n *Node) interact(u Unit) {
	if n.Team == u.Team {
		n.change_population(u.Population)
	} else {
		rem, _ := n.change_population(-1 * u.Population)
		if rem > 0 {
			n.Population = rem
			n.Team = u.Team
		}
	}
	u.Destroy()
}
*/

func (n *Node) update() error {
	n.Element.update()
	return nil
}