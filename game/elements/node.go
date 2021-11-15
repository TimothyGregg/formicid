package elements

import (
	"encoding/json"
	"fmt"

	"github.com/TimothyGregg/formicid/game/util/uid"
)

type Node struct {
	Element
	X      int `json:"x"`
	Y      int `json:"y"`
	Radius int `json:"radius"`
	// Population int `json:"population"`
}

func NewNode(uid *uid.UID, x, y, radius int) *Node {
	n := &Node{Radius: radius}
	n.UID = uid
	n.X = x
	n.Y = y
	return n
}

func (n *Node) MarshalJSON() ([]byte, error) {
	type Alias Node
	return json.Marshal(&struct {
		Position [2]int `json:"position"`
		*Alias
	}{
		Position: [2]int{n.X, n.Y},
		Alias:    (*Alias)(n),
	})
}

func (n *Node) UnmarshalJSON(data []byte) error {
	type Alias Node
	aux := struct {
		Position [2]int `json:"position"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	n.X = aux.Position[0]
	n.Y = aux.Position[1]
	return nil
}

func (n Node) String() string {
	return fmt.Sprintf("Node %d @ (%d, %d)", n.UID.Value(), n.X, n.Y)
}

func (n *Node) update() error {
	n.Element.update()
	return nil
}
