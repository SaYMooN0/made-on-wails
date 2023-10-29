package projectrelated

func (pm *ProjectManager) CurrentProjectGetItemsTypeSuggestion(input string) []string {
	return pm.currentProject.GetItemsTypeSuggestion(input)
}
