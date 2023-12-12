package projectrelated

import (
	"encoding/json"
	"fmt"
	src "made/src"
	"strings"
	"time"
)

func HandleActionWithoutId(actionType src.ActionType, arguments map[string]string, projectFolderPath, version string) *HistoryItem {
	actionId := GenerateRecipeId()
	return HandleAction(actionType, arguments, projectFolderPath, actionId, version)
}

func HandleAction(actionType src.ActionType, arguments map[string]string, projectFolderPath, actionId, version string) *HistoryItem {
	for key, value := range arguments {
		arguments[key] = strings.Replace(value, "'", "", -1)
	}

	var contentToWrite string
	var historyItem *HistoryItem

	switch actionType {
	case src.RecipeRemoved:
		contentToWrite = contentToWriteForRecipeDeleting(arguments)
	case src.StoneCutterAdd:
		contentToWrite = fmt.Sprintf("event.stonecutting('%s','%s')", returnItemWithCount(arguments["output"], arguments["outputCount"]), arguments["input"])
	case src.FurnaceOnlyAdd:
		contentToWrite = fmt.Sprintf("event.smelting('%s', '%s')", arguments["output"], arguments["input"])
	case src.FurnaceAndBlastAdd:
		contentToWrite = fmt.Sprintf("event.smelting('%s', '%s'); event.blasting('%s', '%s');", arguments["output"], arguments["input"], arguments["output"], arguments["input"])
	case src.FurnaceAndSmokerAdd:
		contentToWrite = fmt.Sprintf("event.smelting('%s', '%s'); event.smoking('%s', '%s');", arguments["output"], arguments["input"], arguments["output"], arguments["input"])
	case src.CraftingTableAdd:
		var temp []struct {
			Letter string `json:"letter"`
			Value  string `json:"value"`
		}
		err := json.Unmarshal([]byte(arguments["letterItemDictionary"]), &temp)
		if err != nil {
			panic(err)
		}
		letterItemDictionary := make(map[rune]string)
		for _, item := range temp {
			if len(item.Letter) == 1 {
				letterItemDictionary[rune(item.Letter[0])] = item.Value
			} else {
				fmt.Printf("Invalid letter: %s\n", item.Letter)
			}
		}
		fmt.Println(letterItemDictionary)
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

			var gridInputStrings []string
			err := json.Unmarshal([]byte(arguments["lettersInputGrid"]), &gridInputStrings)
			if err != nil {
				fmt.Println("Error unmarshalling JSON string:", err)
			}
			gridInput := "['"

			for i, char := range gridInputStrings {
				if char == "" {
					char = " "
				}

				if (i+1)%3 == 0 {
					gridInput += char
					if i != len(gridInputStrings)-1 {
						gridInput += "', '"
					}
				} else {
					gridInput += char
				}
			}
			gridInput += "']"
			contentToWrite = fmt.Sprintf("event.shaped('%s', %s, {%s})", returnItemWithCount(arguments["output"], arguments["outputCount"]), gridInput, letterItemString)
		}
	default:
		return nil
	}

	src.WriteRecipeEvent(contentToWrite, projectFolderPath, actionId, version)
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
func contentToWriteForRecipeDeleting(arguments map[string]string) string {
	item, hasItem := arguments["item"]
	asInput := arguments["asInput"] == "true"
	asOutput := arguments["asOutput"] == "true"
	typesStr, hasTypes := arguments["types"]

	var contentToWrite string

	if hasItem {
		if hasTypes && typesStr != "" {
			var types []string
			if err := json.Unmarshal([]byte(typesStr), &types); err == nil {
				var typeStatements []string
				for _, t := range types {
					if asOutput {
						typeStatements = append(typeStatements, fmt.Sprintf("{ type: '%s', output: '%s' }", t, item))
					}
					if asInput {
						typeStatements = append(typeStatements, fmt.Sprintf("{ type: '%s', input: '%s' }", t, item))
					}
				}
				contentToWrite = fmt.Sprintf("event.remove([%s])", strings.Join(typeStatements, ", "))
			}
		} else {
			if asOutput {
				contentToWrite = fmt.Sprintf("event.remove({ output: '%s' })", item)
			}
			if asInput {
				contentToWrite = fmt.Sprintf("event.remove({ input: '%s' })", item)
			}
		}
	}

	return contentToWrite
}
