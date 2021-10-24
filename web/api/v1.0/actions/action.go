package actions

import "encoding/json"

type ActionType int

const (
	Undefined ActionType = iota
	META
)

type Action struct {
	action_type ActionType `json:"action_type"`
	action_details      []byte     `json:"details"`
}

func New_Action(action_bytes []byte) *Action {
	a := &Action{}
	json.Unmarshal(action_bytes, a)
	return a
}

func (a *Action) IsType(T ActionType) bool {
	return a.action_type == T
}