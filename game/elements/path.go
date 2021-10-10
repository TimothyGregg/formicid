package elements

import (
	"encoding/json"

	graph "github.com/TimothyGregg/Antmound/game/graph"
)

type Path struct {
	Element
	edge *graph.Edge
}

func (p *Path) MarshalJSON() ([]byte, error) {
	type Alias Path
	v1, v2 := p.Vertices()
	return json.Marshal(&struct {
		Vertices [2]*graph.Vertex `json:"vertices"`
		*Alias
	}{
		Vertices: [2]*graph.Vertex{v1, v2},
		Alias: (*Alias)(p),
	})
}

func New_Path(uid int, edge *graph.Edge) *Path {
	p := &Path{edge: edge}
	p.New(uid)
	return p
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
