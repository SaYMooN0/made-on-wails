package src

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
	Fabric Loader = iota
	Forge
)
