package uid

import "encoding/json"

type UID struct {
	value int
}

func (u *UID) MarshalJSON() ([]byte, error) {
	type Alias UID
	return json.Marshal(&struct {
		Value int `json:"value"`
		*Alias
	}{
		Value: u.value,
		Alias: (*Alias)(u),
	})
}

func (u *UID) UnmarshalJSON(data []byte) error {
	type Alias UID
	aux := struct {
		Value int `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.value = aux.value
	return nil
}

func NewUID(value int) *UID {
	return &UID{value: value}
}

func (u *UID) Value() int {
	return u.value
}
