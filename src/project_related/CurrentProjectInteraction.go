package projectrelated

import (
	"fmt"
	src "made/src"
)

func (pm *ProjectManager) CurrentProjectGetItemsTypeSuggestion(input string) []string {
	return pm.currentProject.GetItemsTypeSuggestion(input)
}
func (pm *ProjectManager) CurrentProjectAddNewRecipe(actionTypeStr string, arguments map[string]string) *HistoryItem {
	actionType, err := src.StringToActionType(actionTypeStr)
	fmt.Println(actionType, actionTypeStr, arguments)
	if err != nil {
		fmt.Println("Error converting string to ActionType:", err)
	}
	return pm.currentProject.AddNewRecipe(actionType, arguments)
}
func (pm *ProjectManager) CurrentProjectChangeAction(actionId, filePath string, actionTypeStr string, arguments map[string]string) {
	fmt.Println(actionId, filePath, actionTypeStr, arguments)
	actionType, err := src.StringToActionType(actionTypeStr)
	if err != nil {
		fmt.Println("Error converting string to ActionType:", err)
	}

	pm.currentProject.ChangeAction(actionId, filePath, actionType, arguments)
}
