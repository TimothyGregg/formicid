package elements

import (
	"math"

	graph "github.com/TimothyGregg/Antmound/graph"
)

type Node struct {
	Element
	UID             int
	vertex          *graph.Vertex
	population_cap  float64
	generation_rate float64
	radius          float64
	population      int
	Team            Team
}

func (n Node) String() string {
	return n.vertex.String()
}

func (n *Node) Get() (int, int, float64, float64, float64) {
	x, y := n.vertex.Get()
	return x, y, n.population_cap, n.generation_rate, n.radius
}

func (node *Node) node_distance(x, y float64) float64 {
	nx, ny := node.vertex.Get()
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
