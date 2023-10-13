package projectrelated

import (
	"encoding/json"
	"fmt"
	src "made/src"
	"strings"
	"time"
)

func HandleActionWithoutId(actionType src.ActionType, arguments map[string]string, projectFolderPath string) *HistoryItem {
	actionId := GenerateRecipeId()
	return HandleAction(actionType, arguments, projectFolderPath, actionId)
}

func HandleAction(actionType src.ActionType, arguments map[string]string, projectFolderPath, actionId string) *HistoryItem {
	for key, value := range arguments {
		arguments[key] = strings.Replace(value, "'", "", -1)
	}

	var contentToWrite string
	var historyItem *HistoryItem

	switch actionType {
	case src.StonecutterAdd:
		contentToWrite = fmt.Sprintf("event.stonecutting('%s','%s')", returnItemWithCount(arguments["output"], arguments["outputCount"]), arguments["input"])
	case src.FurnaceOnlyAdd:
		contentToWrite = fmt.Sprintf("event.smelting('%s', '%s')", arguments["output"], arguments["input"])
	case src.FurnaceAndBlastAdd:
		contentToWrite = fmt.Sprintf("event.smelting('%s', '%s'); event.blasting('%s', '%s');", arguments["output"], arguments["input"], arguments["output"], arguments["input"])
	case src.FurnaceAndSmokerAdd:
		contentToWrite = fmt.Sprintf("event.smelting('%s', '%s'); event.smoking('%s', '%s');", arguments["output"], arguments["input"], arguments["output"], arguments["input"])
	case src.CraftingTableAdd:
		letterItemDictionary := make(map[rune]string)
		err := json.Unmarshal([]byte(arguments["letterItemDictionary"]), &letterItemDictionary)
		if err != nil {
			panic(err)
		}

		if len(letterItemDictionary) < 1 {
			return nil
		}

		isShapeless := strings.ToLower(arguments["isShapeless"]) == "true"
		if isShapeless {
			itemsString := ""
			for _, letter := range arguments["lettersInputGrid"] {
				if item, ok := letterItemDictionary[letter]; ok {
					itemsString += fmt.Sprintf("'%s',", item)
				}
			}
			itemsString = strings.TrimSuffix(itemsString, ",")
			contentToWrite = fmt.Sprintf("event.shapeless('%s', [%s])", returnItemWithCount(arguments["output"], arguments["outputCount"]), itemsString)
		} else {
			letterItemString := ""
			for letter, item := range letterItemDictionary {
				letterItemString += fmt.Sprintf("%c: '%s',", letter, item)
			}
			letterItemString = strings.TrimSuffix(letterItemString, ",")
			gridInputStrings := strings.Split(arguments["lettersInputGrid"], ",")
			gridInput := fmt.Sprintf("'%s','%s','%s'", gridInputStrings[0], gridInputStrings[1], gridInputStrings[2])
			contentToWrite = fmt.Sprintf("event.shaped('%s', [%s], {%s})", returnItemWithCount(arguments["output"], arguments["outputCount"]), gridInput, letterItemString)
		}
	default:
		return nil
	}

	src.WriteVanillaRecipe(contentToWrite, projectFolderPath, actionId)
	historyItem = NewHistoryItem(arguments, src.GetFullVanillaPath(), actionId, actionType)
	return historyItem
}

func DeleteAction(actionId, fullPathToFile string) {
	src.DeleteAction(actionId, fullPathToFile)
}

func GenerateRecipeId() string {
	return time.Now().Format("060102150405")
}

func returnItemWithCount(item, itemCount string) string {
	if itemCount == "1" {
		return item
	}
	return fmt.Sprintf("%sx %s", itemCount, item)
}
