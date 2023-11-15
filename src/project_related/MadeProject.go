package projectrelated

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	src "made/src"
	minecraft "made/src/minecraft_related"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	MadeProjectFileExt = ".madeProject"
	KubejsModId        = "kubejs"
)
const TagChar byte = '#'

type MadeProject struct {
	Name         string          `json:"Name"`
	FullPath     string          `json:"FullPath"`
	PathToFolder string          `json:"PathToFolder"`
	Version      string          `json:"Version"`
	Loader       src.Loader      `json:"Loader"`
	CreationDate time.Time       `json:"CreationDate"`
	LastUpdated  time.Time       `json:"LastUpdated"`
	Settings     ProjectSettings `json:"Settings"`
	History      []HistoryItem   `json:"History"`
	Mods         []minecraft.Mod `json:"Mods"`
}

func (mp *MadeProject) SaveToFile() error {
	data, err := json.MarshalIndent(mp, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(mp.FullPath, data, 0644)
}

func CreateFromFile(filePath string) (*MadeProject, error) {
	data, err := os.ReadFile(filePath)
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
		*minecraft.NewItemWithName("redstone", "Redstone"),
	)
	minecraftMod.Blocks = append(minecraftMod.Blocks,
		*minecraft.NewBlockWithName("granite", "Granite"),
		*minecraft.NewBlockWithName("dirt", "Dirt"),
		*minecraft.NewBlockWithName("andesite", "Andesite"),
		*minecraft.NewBlockWithName("sand", "Sand"))
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
		Name:         name,
		FullPath:     fullPath,
		PathToFolder: pathToFolder,
		Version:      version,
		Loader:       loader,
		CreationDate: currentTime,
		LastUpdated:  currentTime,
		Settings:     ProjectSettings{ShowWarningWhenDeletingAction: true},
		History:      []HistoryItem{},
		Mods:         modList,
	}
}

func NewMadeProjectWithParams(name, fullPath, pathToFolder, version string, loader src.Loader, creationDate, lastUpdated time.Time, settings ProjectSettings, history []HistoryItem, mods []minecraft.Mod) *MadeProject {
	return &MadeProject{
		Name:         name,
		FullPath:     fullPath,
		PathToFolder: pathToFolder,
		Version:      version,
		Loader:       loader,
		CreationDate: creationDate,
		LastUpdated:  lastUpdated,
		Settings:     settings,
		History:      history,
		Mods:         mods,
	}
}
func (mp *MadeProject) AddNewRecipe(actionType src.ActionType, arguments map[string]string) *HistoryItem {
	mp.LastUpdated = time.Now()
	historyItem := HandleActionWithoutId(actionType, arguments, mp.PathToFolder)
	if historyItem != nil {
		mp.History = append(mp.History, *historyItem)
	}
	mp.SaveToFile()
	return historyItem
}

func (mp *MadeProject) ChangeAction(actionId, filePath string, actionType src.ActionType, arguments map[string]string) {
	DeleteAction(actionId, filePath)
	mp.LastUpdated = time.Now()
	historyItem := HandleAction(actionType, arguments, mp.PathToFolder, actionId)
	if historyItem != nil {
		mp.History = append(mp.History, *historyItem)
	}
	mp.SaveToFile()
}
func (mp *MadeProject) DeleteAction(actionId, filePath string) {
	fullPathToFile := filepath.Join(mp.PathToFolder, filepath.FromSlash(filePath))
	if mp.tryDeleteHistoryItemByActionId(actionId) {
		src.DeleteAction(actionId, fullPathToFile)
	}
	mp.SaveToFile()
}

func (mp *MadeProject) tryDeleteHistoryItemByActionId(actionId string) bool {
	for i, item := range mp.History {
		if item.ActionID == actionId {
			mp.History = append(mp.History[:i], mp.History[i+1:]...)
			return true
		}
	}
	return false
}
func (mp *MadeProject) GetAllItems() []src.IdNamePair {
	var results []src.IdNamePair
	for _, mod := range mp.Mods {
		for _, item := range mod.Items {
			results = append(results, src.IdNamePair{Id: mod.Id + ":" + item.Id, InGameName: item.InGameName})
		}
	}
	return results
}

func (mp *MadeProject) GetAllBlocks() []src.IdNamePair {
	var results []src.IdNamePair
	for _, mod := range mp.Mods {
		for _, block := range mod.Blocks {
			results = append(results, src.IdNamePair{Id: mod.Id + ":" + block.Id, InGameName: block.InGameName})
		}
	}
	return results
}

func (mp *MadeProject) GetAllTags() []src.IdNamePair {
	var results []src.IdNamePair
	for _, mod := range mp.Mods {
		for _, tag := range mod.Tags {
			results = append(results, src.IdNamePair{Id: "#" + mod.Id + ":" + tag.Id, InGameName: tag.InGameName})
		}
	}
	return results
}

func (mp *MadeProject) GetAllMods() []src.IdNamePair {
	var results []src.IdNamePair
	for _, mod := range mp.Mods {
		results = append(results, src.IdNamePair{Id: mod.Id, InGameName: mod.InGameName})
	}
	return results
}

func (mp *MadeProject) GetAllProcessingTypes() []src.IdNamePair {
	var results []src.IdNamePair
	seen := make(map[src.IdNamePair]bool)
	for _, mod := range mp.Mods {
		for _, supportedType := range mod.SupportedTypes {
			pair := src.IdNamePair{Id: supportedType.Id, InGameName: supportedType.InGameName}
			if _, exists := seen[pair]; !exists {
				results = append(results, pair)
				seen[pair] = true
			}
		}
	}
	return results
}
func (mp *MadeProject) AddNewItem(itemIdString, inGameName string) string {
	parts := strings.Split(itemIdString, ":")
	modString := parts[0]
	itemId := parts[1]

	var mod *minecraft.Mod
	for _, m := range mp.Mods {
		if m.Id == modString {
			mod = &m
			break
		}
	}

	if mod == nil {
		mod = &minecraft.Mod{Id: modString}
		mod.Items = append(mod.Items, minecraft.Item{Id: itemId, InGameName: inGameName})
		mp.Mods = append(mp.Mods, *mod)
	} else {
		for _, item := range mod.Items {
			if item.Id == inGameName {
				return fmt.Sprintf("Mod with id %s already contains this item", mod.Id)
			}
		}
		mod.Items = append(mod.Items, minecraft.Item{Id: itemId, InGameName: inGameName})
	}
	mp.SaveToFile()
	return ""
}

func (mp *MadeProject) EditCollectionItem(oldItem, newItem, newInGameName string) string {
	processResult := mp.AddNewItem(newItem, newInGameName)
	if processResult != "" {
		return processResult
	}
	processResult = mp.DeleteItemFromCollection(oldItem)
	if processResult != "" {
		return processResult
	}
	return ""
}

func (mp *MadeProject) DeleteItemFromCollection(itemToDelete string) string {
	item := mp.GetItemById(itemToDelete)
	if item == nil {
		return fmt.Sprintf("No item with id %s found", itemToDelete)
	}

	modString := strings.Split(itemToDelete, ":")[0]
	var mod *minecraft.Mod
	for _, m := range mp.Mods {
		if m.Id == modString {
			mod = &m
			break
		}
	}

	if mod != nil {
		for i, itm := range mod.Items {
			if itm.Id == item.Id {
				mod.Items = append(mod.Items[:i], mod.Items[i+1:]...)
				break
			}
		}
		mp.SaveToFile()
		return ""
	}
	return fmt.Sprintf("No mod with id %s found", modString)
}

func (mp *MadeProject) GetItemById(itemIdString string) *minecraft.Item {
	parts := strings.Split(itemIdString, ":")
	modString := parts[0]
	itemId := parts[1]

	for _, mod := range mp.Mods {
		if mod.Id == modString {
			for _, item := range mod.Items {
				if item.Id == itemId {
					return &item
				}
			}
		}
	}
	return nil
}
func (mp *MadeProject) GetItemImgPathById(id string) string {
	i := mp.GetItemById(id)
	if i == nil || strings.Split(id, ":")[0] != KubejsModId {
		return ""
	}
	texturePath := src.GetFullTextureItemPath()
	return filepath.Join(mp.PathToFolder, texturePath, i.GetImagePath())
}
func (mp *MadeProject) GetItemImgInBase64(itemId string) string {
	imagePath := mp.GetItemImgPathById(itemId)
	if imagePath == "" {
		return ""
	}

	data, err := os.ReadFile(imagePath)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(data)
}
func (mp *MadeProject) GetItemsTypeSuggestion(input string) []string {
	if input == "" {
		return []string{}
	}

	if !strings.Contains(input, ":") {
		var suggestions []string
		for _, mod := range mp.Mods {
			if strings.HasPrefix(input, string(TagChar)) {
				if strings.Contains(mod.Id, strings.TrimPrefix(input, string(TagChar))) {
					suggestions = append(suggestions, string(TagChar)+mod.Id)
				}
			} else {
				if strings.Contains(mod.Id, input) {
					suggestions = append(suggestions, mod.Id)
				}
			}
		}
		return suggestions
	}

	mod := getModFromInput(mp.Mods, input)
	if mod == nil {
		return []string{}
	}

	afterModIDString := strings.Split(input, ":")[1]

	if input[0] == TagChar {
		var tagSuggestions []string
		for _, tag := range mod.Tags {
			if strings.Contains(tag.Id, afterModIDString) {
				tagSuggestions = append(tagSuggestions, string(TagChar)+mod.Id+":"+tag.Id)
			}
		}
		return tagSuggestions
	}

	var itemSuggestions []string
	for _, item := range mod.Items {
		if strings.Contains(item.Id, afterModIDString) {
			itemSuggestions = append(itemSuggestions, mod.Id+":"+item.Id)
		}
	}
	for _, block := range mod.Blocks {
		if strings.Contains(block.Id, afterModIDString) {
			itemSuggestions = append(itemSuggestions, mod.Id+":"+block.Id)
		}
	}
	return itemSuggestions
}
func getModFromInput(mods []minecraft.Mod, input string) *minecraft.Mod {
	if len(input) > 0 && input[0] == TagChar {
		input = input[1:]
	}
	modID := strings.Split(input, ":")[0]
	for _, mod := range mods {
		if mod.Id == modID {
			return &mod
		}
	}
	return nil
}
