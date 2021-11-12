package elements

import (
	"encoding/json"

	graph "github.com/TimothyGregg/formicid/game/util/graph"
)

type Path struct {
	Element
	graph.Edge
}

func (p *Path) MarshalJSON() ([]byte, error) {
	type Alias Path
	v1, v2 := p.Vertices()
	return json.Marshal(&struct {
		Vertices [2]*Node `json:"nodes"`
		*Alias
	}{
		Vertices: [2]*Node{v1, v2},
		Alias:    (*Alias)(p),
	})
}

func New_Path(uid int) *Path {
	p := &Path{}
	p.New(uid)
	return p
}

func (p Path) String() string {
	return p.String()
}

func (p *Path) Vertices() (*Node, *Node) {
	return p.Vertices()
}

func (p *Path) update() error {
	p.Element.tick()
	return nil
}
