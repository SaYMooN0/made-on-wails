package minecraftrelated

import (
	"encoding/json"
)

type Mod struct {
	Id             string           `json:"id"`
	InGameName     string           `json:"inGameName"`
	Items          []Item           `json:"items"`
	Blocks         []Block          `json:"blocks"`
	Tags           []Tag            `json:"tags"`
	SupportedTypes []ProcessingType `json:"supportedTypes"`
}

func NewMod(id, inGameName string) *Mod {
	return &Mod{
		Id:             id,
		InGameName:     inGameName,
		Items:          []Item{},
		Blocks:         []Block{},
		Tags:           []Tag{},
		SupportedTypes: []ProcessingType{},
	}
}

func NewModWithId(id string) *Mod {
	return NewMod(id, "")
}
func (m *Mod) UnmarshalJSON(data []byte) error {
	type Alias Mod
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
