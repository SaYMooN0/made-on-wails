package projectrelated

import src "made/src"

func (pm *ProjectManager) CurrentProjectGetItemsTypeSuggestion(input string) []string {
	return pm.currentProject.GetItemsTypeSuggestion(input)
}
func (pm *ProjectManager) CurrentProjectAddNewRecipe(actionType src.ActionType, arguments map[string]string) *HistoryItem {
	return pm.currentProject.AddNewRecipe(actionType, arguments)
}
