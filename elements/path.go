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

func (p *Path) Vertices() (*graph.Vertex, *graph.Vertex) {
	return p.edge.Vertices()
}

func (p *Path) update() error {
	p.Element.tick()
	return nil
}
