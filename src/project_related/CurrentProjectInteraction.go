package projectrelated

import (
	"fmt"
	src "made/src"
)

func (pm *ProjectManager) CurrentProjectGetItemSuggestion(input string) []string {
	return pm.currentProject.GetItemSuggestion(input)
}
func (pm *ProjectManager) CurrentProjectGetProcessingTypeSuggestion(input string) []string {
	return pm.currentProject.GetProcessingTypeSuggestion(input)
}
func (pm *ProjectManager) CurrentProjectDeleteAction(actionId, filePath string) {
	pm.currentProject.DeleteAction(actionId, filePath)
}
func (pm *ProjectManager) CurrentProjectHistory() []HistoryItem {
	return pm.currentProject.History
}
func (pm *ProjectManager) DoShowWarningWhenDeletingActionForCurrentProjectHistory() bool {
	return pm.currentProject.Settings.ShowWarningWhenDeletingAction
}
func (pm *ProjectManager) SetDoShowWarningWhenDeletingActionForCurrentProjectHistory(doShowWarningWhenDeletingAction bool) {
	pm.currentProject.Settings.ShowWarningWhenDeletingAction = doShowWarningWhenDeletingAction
}
func (pm *ProjectManager) CurrentProjectRemoveRecipe(arguments map[string]string) *HistoryItem {
	actionType, _ := src.StringToActionType("RecipeRemoved")
	return pm.currentProject.HandleRecipesChanges(actionType, arguments)
}
func (pm *ProjectManager) CurrentProjectAddNewRecipe(actionTypeStr string, arguments map[string]string) *HistoryItem {
	actionType, _ := src.StringToActionType(actionTypeStr)
	return pm.currentProject.HandleRecipesChanges(actionType, arguments)
}
func (pm *ProjectManager) CurrentProjectChangeAction(actionId, filePath string, actionTypeStr string, arguments map[string]string) {
	fmt.Println(actionId, filePath, actionTypeStr, arguments)
	actionType, err := src.StringToActionType(actionTypeStr)
	if err != nil {
		fmt.Println("Error converting string to ActionType:", err)
	}
	pm.currentProject.ChangeAction(actionId, filePath, actionType, arguments)
}
