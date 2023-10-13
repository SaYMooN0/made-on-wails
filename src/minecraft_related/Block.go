package minecraftrelated

import (
	"encoding/json"
)

type Block struct {
	Id            string  `json:"id"`
	InGameName    string  `json:"inGameName"`
	MaxStackSize  int     `json:"maxStackSize"`
	BurnTime      *int    `json:"burnTime"`
	FireResistant bool    `json:"fireResistant"`
	Hardness      float32 `json:"hardness"`
	Resistance    float32 `json:"resistance"`
}

func NewBlock(id string) *Block {
	return &Block{
		Id:         id,
		InGameName: "",
	}
}

func NewBlockWithName(id, inGameName string) *Block {
	return &Block{
		Id:         id,
		InGameName: inGameName,
	}
}

func NewBlockDetailed(id, inGameName string, maxStackSize int, burnTime *int, fireResistant bool, hardness, resistance float32) *Block {
	return &Block{
		Id:            id,
		InGameName:    inGameName,
		MaxStackSize:  maxStackSize,
		BurnTime:      burnTime,
		FireResistant: fireResistant,
		Hardness:      hardness,
		Resistance:    resistance,
	}
}
func (b *Block) UnmarshalJSON(data []byte) error {
	type Alias Block
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
