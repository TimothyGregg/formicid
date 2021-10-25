package actions

import "encoding/json"

type ActionType int

const (
	Undefined ActionType = iota
	META
)

type Action struct {
	Action_type ActionType `json:"action_type"`
	Action_name string `json:"action_name"`
	Action_details      []byte     `json:"details"`
}

func NewAction(action_bytes []byte) (*Action, error) {
	a := &Action{}
	err := json.Unmarshal(action_bytes, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Action) IsType(T ActionType) bool {
	return a.Action_type == T
}