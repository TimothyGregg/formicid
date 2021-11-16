package elements

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/TimothyGregg/formicid/game/util/uid"
)

type Path struct {
	Element
	Nodes  [2]*Node `json:"nodes"`
	Length float64  `json:"length"`
}

func (p *Path) MarshalJSON() ([]byte, error) {
	type Alias Path
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	})
}

func New_Path(uid *uid.UID, n1, n2 *Node) *Path {
	p := &Path{Nodes: [2]*Node{n1, n2}}
	p.UID = uid
	val := math.Pow(float64(n1.X-n2.X), 2)
	p.Length = math.Sqrt(val + math.Pow(float64(n1.Y-n2.Y), 2))
	return p
}

func (p Path) String() string {
	return fmt.Sprintf("Path from Node %d to Node %d", p.Nodes[0].UID.Value(), p.Nodes[1].UID.Value())
}

func (p *Path) update() error {
	p.Element.tick()
	return nil
}
