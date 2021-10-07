package elements

import (
	"errors"
	"math"

	graph "github.com/TimothyGregg/Antmound/graph"
)

type Node struct {
	Element
	uid             int
	vertex          *graph.Vertex
	radius          float64
	population      int
	Team            Team
}

func (n *Node) ID() int {
	return n.uid
}

func (n Node) String() string {
	return n.vertex.String()
}

func (n *Node) Get() (int, int, float64) {
	x, y := n.vertex.Position()
	return x, y, n.radius
}

func (node *Node) node_distance(x, y float64) float64 {
	nx, ny := node.vertex.Position()
	val := math.Pow(float64(nx)-x, 2)
	return math.Sqrt(val + math.Pow(float64(ny)-y, 2))
}

func (n *Node) change_population(val int) (int, error) {
	rem := -1 * (n.population + val)
	n.population += val
	if n.population < 0 {
		n.population = 0
		return rem, nil
	}
	return 0, nil
}

func (n *Node) New_Unit(val int) (*Unit, error) {
	if val > n.population {
		return nil, errors.New("cannot create a unit with more population than the node")
	} else if val <= 0 {
		return nil, errors.New("unit created with power less than 1")
	}
	u := &Unit{population: val}
	u.Team = n.Team
	u.UID = u.Team.Unit_UID_Generator.Next()
	n.change_population(-1 * val)
	return u, nil
}

func (n *Node) interact(u Unit) {
	if n.Team == u.Team {
		n.change_population(u.population)
	} else {
		rem, _ := n.change_population(-1 * u.population)
		if rem > 0 {
			n.population = rem
			n.Team = u.Team
		}
	}
	u.Destroy()
}

func (n *Node) update() error {
	n.Element.update()
	return nil
}