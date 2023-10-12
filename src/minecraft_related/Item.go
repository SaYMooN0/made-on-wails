package minecraftrelated

import (
	"encoding/json"
)

type Item struct {
	Id            string `json:"Id"`
	InGameName    string `json:"InGameName"`
	MaxStackSize  int    `json:"MaxStackSize"`
	BurnTime      *int   `json:"BurnTime,omitempty"`
	FireResistant bool   `json:"FireResistant"`
}

const pngExtension = ".png"

func NewItem(id string) *Item {
	return &Item{
		Id:            id,
		InGameName:    "",
		MaxStackSize:  64,
		BurnTime:      nil,
		FireResistant: false,
	}
}

func NewItemWithName(id, inGameName string) *Item {
	return &Item{
		Id:            id,
		InGameName:    inGameName,
		MaxStackSize:  64,
		BurnTime:      nil,
		FireResistant: false,
	}
}

func (i *Item) GetImagePath() string {
	return i.Id + pngExtension
}

func (i *Item) UnmarshalJSON(data []byte) error {
	type Alias Item
	aux := &struct {
		MaxStackSize  int  `json:"MaxStackSize"`
		BurnTime      *int `json:"BurnTime"`
		FireResistant bool `json:"FireResistant"`
		*Alias
	}{
		MaxStackSize:  64,
		FireResistant: false,
		Alias:         (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	i.MaxStackSize = aux.MaxStackSize
	i.BurnTime = aux.BurnTime
	i.FireResistant = aux.FireResistant
	return nil
}
