package projectrelated

import (
	"encoding/json"
	src "made/src"
	minecraft "made/src/minecraft_related"
	"os"
	"time"
)

const (
	FileExtension = ".madeProject"
	kubejsModId   = "kubejs"
)

type MadeProject struct {
	Name                  string                 `json:"Name"`
	FullPath              string                 `json:"FullPath"`
	PathToFolder          string                 `json:"PathToFolder"`
	Version               string                 `json:"Version"`
	Loader                src.Loader             `json:"Loader"`
	CreationDate          time.Time              `json:"CreationDate"`
	LastUpdated           time.Time              `json:"LastUpdated"`
	Settings              ProjectSettings        `json:"Settings"`
	History               []HistoryItem          `json:"History"`
	Mods                  []minecraft.Mod        `json:"Mods"`
	SuggestionsCollection *SuggestionsCollection `json:"-"`
}

func (mp *MadeProject) SaveToFile() error {
	data, err := json.MarshalIndent(mp, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(mp.FullPath, data, 0644)
}

func CreateFromFile(filePath string) (*MadeProject, error) {
	data, err := os.ReadFile(filePath) // Update this line as well
	if err != nil {
		return nil, err
	}

	var project MadeProject
	err = json.Unmarshal(data, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func NewMadeProject(name, fullPath, pathToFolder, version string, loader src.Loader) *MadeProject {
	currentTime := time.Now()

	minecraftMod := minecraft.NewMod("minecraft", "Minecraft")
	minecraftMod.Items = append(minecraftMod.Items,
		*minecraft.NewItemWithName("stone", "Stone"),
		*minecraft.NewItemWithName("granite", "Granite"),
		*minecraft.NewItemWithName("dirt", "Dirt"),
		*minecraft.NewItemWithName("andesite", "Andesite"),
		*minecraft.NewItemWithName("sand", "Sand"),
	)
	minecraftMod.SupportedTypes = append(minecraftMod.SupportedTypes,
		*minecraft.NewProcessingType("shapeless", "Shapeless crafting", true),
		*minecraft.NewProcessingType("minecraft:campfire_cooking", "Campfire cooking", true),
	)

	modList := []minecraft.Mod{*minecraftMod}

	switch loader {
	case src.Forge:
		forgeMod := minecraft.NewMod("forge", "Forge")
		forgeMod.Tags = append(forgeMod.Tags,
			*minecraft.NewTag("ores", "Forge ores"),
			*minecraft.NewTag("ores/copper", "Forge copper ores"),
		)
		modList = append(modList, *forgeMod)
	case src.Fabric:
		fabricMod := minecraft.NewMod("fabric", "Fabric")
		modList = append(modList, *fabricMod)
	}

	return &MadeProject{
		Name:                  name,
		FullPath:              fullPath,
		PathToFolder:          pathToFolder,
		Version:               version,
		Loader:                loader,
		CreationDate:          currentTime,
		LastUpdated:           currentTime,
		Settings:              ProjectSettings{}, // Assuming default initialization is fine
		History:               []HistoryItem{},
		Mods:                  modList,
		SuggestionsCollection: NewSuggestionsCollection(loader, modList), // Assuming you have a constructor for this
	}
}

func NewMadeProjectWithParams(name, fullPath, pathToFolder, version string, loader src.Loader, creationDate, lastUpdated time.Time, settings ProjectSettings, history []HistoryItem, mods []minecraft.Mod) *MadeProject {
	return &MadeProject{
		Name:                  name,
		FullPath:              fullPath,
		PathToFolder:          pathToFolder,
		Version:               version,
		Loader:                loader,
		CreationDate:          creationDate,
		LastUpdated:           lastUpdated,
		Settings:              settings,
		History:               history,
		Mods:                  mods,
		SuggestionsCollection: NewSuggestionsCollection(loader, mods), // Assuming you have a constructor for this
	}
}
