package elements

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/TimothyGregg/formicid/game/util/uid"
)

type Path struct {
	Element
	Nodes  [2]int  `json:"nodes"`
	Length float64 `json:"length"`
}

func (p *Path) MarshalJSON() ([]byte, error) {
	type Alias Path
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	})
}

func (p *Path) UnmarshalJSON(data []byte) error {
	type Alias Path
	aux := struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func New_Path(uid *uid.UID, n1, n2 *Node) *Path {
	p := &Path{Nodes: [2]int{n1.UID.Value(), n2.UID.Value()}}
	p.UID = uid
	val := math.Pow(float64(n1.X-n2.X), 2)
	p.Length = math.Sqrt(val + math.Pow(float64(n1.Y-n2.Y), 2))
	return p
}

func (p Path) String() string {
	return fmt.Sprintf("Path from Node %d to Node %d", p.Nodes[0], p.Nodes[1])
}

func (p *Path) update() error {
	p.Element.tick()
	return nil
}
