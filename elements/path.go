package elements

import graph "github.com/TimothyGregg/Antmound/graph"

type Path struct {
	Element
	UID  int
	edge *graph.Edge
}

func (p Path) String() string {
	return p.edge.String()
}

func (p *Path) Get() (*graph.Vertex, *graph.Vertex) {
	return p.edge.Get()
}

func (p *Path) update() error {
	p.Element.tick()
	return nil
}
