package src

import (
	"fmt"
	"strings"
)

type ActionType int
type CollectionItemType int
type Loader int

const (
	StonecutterAdd ActionType = iota
	FurnaceOnlyAdd
	FurnaceAndSmokerAdd
	FurnaceAndBlastAdd
	CraftingTableAdd
	RecipeRemoved
)
const (
	Item CollectionItemType = iota
	Block
	Tag
	Mod
	Type
)
const (
	Forge Loader = iota
	Fabric
)

func StringToLoader(loaderString string) (Loader, error) {
	switch strings.ToLower(loaderString) {
	case "fabric":
		return Fabric, nil
	case "forge":
		return Forge, nil
	default:
		return -1, fmt.Errorf("unknown loader: %s", loaderString)
	}
}
