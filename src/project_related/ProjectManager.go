package projectrelated

import (
	"context"
	"encoding/json"
	"fmt"
	src "made/src"
	"os"
	"os/exec"
	"path/filepath"
	goRuntime "runtime"
	"strings"
	"unicode"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const linksFileName = "links.madeLinks"

type ProjectManager struct {
	ProjectLinks       []string        `json:"_projectLinks"`
	PinnedProjectLinks []string        `json:"_pinnedProjectLinks"`
	Projects           []MadeProject   `json:"-"`
	currentProject     MadeProject     `json:"-"`
	ctx                context.Context `json:"-"`
}

func (pm *ProjectManager) CurrentProject() MadeProject {
	return pm.currentProject
}
func InitializeProjectManager() *ProjectManager {
	pm := &ProjectManager{}
	if _, err := os.Stat(linksFileName); os.IsNotExist(err) {
		defaultData, _ := json.MarshalIndent(pm, "", "  ")
		os.WriteFile(linksFileName, defaultData, 0644)
	} else {
		data, _ := os.ReadFile(linksFileName)
		json.Unmarshal(data, pm)
		pm.Projects = pm.GetProjects()
	}

	return pm
}

func (pm *ProjectManager) SetStartupOnContext(ctx context.Context) {
	pm.ctx = ctx
}
func (pm *ProjectManager) GetProjects() []MadeProject {
	var projects []MadeProject

	for _, projectString := range pm.ProjectLinks {
		project, err := CreateFromFile(projectString)
		if err == nil && project != nil {
			projects = append(projects, *project)
		}
	}
	return projects
}
func (pm *ProjectManager) DeleteProjectLink(linkPath string) string {
	for i, link := range pm.ProjectLinks {
		if link == linkPath {
			pm.ProjectLinks = append(pm.ProjectLinks[:i], pm.ProjectLinks[i+1:]...)
			pm.SaveToFile()
			return ""
		}
	}
	return "Error: Link path not found in Project Links."
}
func (pm *ProjectManager) PinProjectLink(linkPath string) string {
	exists := false
	for _, link := range pm.ProjectLinks {
		if link == linkPath {
			exists = true
			break
		}
	}
	if !exists {
		return "Error: Link path not found in Project Links."
	}
	for _, pinnedLink := range pm.PinnedProjectLinks {
		if pinnedLink == linkPath {
			return "Error: Project already pinned."
		}
	}

	pm.PinnedProjectLinks = append(pm.PinnedProjectLinks, linkPath)
	pm.SaveToFile()
	return ""
}
func (pm *ProjectManager) UnpinProjectLink(linkPath string) string {
	for i, pinnedLink := range pm.PinnedProjectLinks {
		if pinnedLink == linkPath {
			pm.PinnedProjectLinks = append(pm.PinnedProjectLinks[:i], pm.PinnedProjectLinks[i+1:]...)
			pm.SaveToFile()
			return ""
		}
	}
	return "Error: Link path not found in Pinned Project Links."
}

// func (pm *ProjectManager) getPinnedProjects() []MadeProject {
// 	var pinnedProjects []MadeProject
// 	for _, pinnedProjectString := range pm.PinnedProjectLinks {
// 		pinnedProject, err := CreateFromFile(pinnedProjectString)
// 		if err == nil && pinnedProject != nil {
// 			pinnedProjects = append(pinnedProjects, *pinnedProject)
// 		}
// 	}
// 	return pinnedProjects
// }

func (pm *ProjectManager) SaveToFile() {
	data, _ := json.MarshalIndent(pm, "", "  ")
	os.WriteFile(linksFileName, data, 0644)

	for _, project := range pm.GetProjects() {
		project.SaveToFile()
	}
}

func (pm *ProjectManager) CreateProject(name, pathToFolder, version string, loader src.Loader) string {
	fullPath := filepath.Join(pathToFolder, name+MadeProjectFileExt)

	if _, err := os.Stat(pathToFolder); err != nil {
		if os.IsNotExist(err) {
			return fmt.Sprintf("path does not exist: %s", pathToFolder)
		}
		return fmt.Sprintf("error checking path: %s", err)
	}

	project := NewMadeProject(name, fullPath, pathToFolder, version, loader)
	if err := project.SaveToFile(); err != nil {
		return fmt.Sprintf("error saving project to file: %s", err)
	}
	pm.ProjectLinks = append(pm.ProjectLinks, project.FullPath)
	pm.SaveToFile()
	return ""
}

func (pm *ProjectManager) SetCurrentProjectByFolder(folderPath string) {

	for _, project := range pm.GetProjects() {
		if project.PathToFolder == folderPath {
			pm.currentProject = project
			break
		}
	}

}

func (pm *ProjectManager) AnyMadeProjectFilesInFolder(pathToFolder string) bool {
	if _, err := os.Stat(pathToFolder); os.IsNotExist(err) {
		return false
	}

	err := filepath.Walk(pathToFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), MadeProjectFileExt) {
			return fmt.Errorf("found")
		}

		return nil
	})

	return err != nil && err.Error() == "found"
}
func (pm *ProjectManager) AddProjectToCollectionIfNeeded(pathToFile string) {
	for _, link := range pm.ProjectLinks {
		if link == pathToFile {
			return
		}
	}
	project, err := CreateFromFile(pathToFile)
	if err != nil {
		return
	}
	pm.Projects = append(pm.Projects, *project)
	pm.ProjectLinks = append(pm.ProjectLinks, project.FullPath)
	pm.SaveToFile()
}

func (pm *ProjectManager) ChooseFolderForNewProject() string {
	dialogOptions := runtime.OpenDialogOptions{
		Title: "Select a directory",
	}
	folder, err := runtime.OpenDirectoryDialog(pm.ctx, dialogOptions)
	if err != nil {
		return ""
	}
	return folder
}
func (pm *ProjectManager) OpenExistingProject() string {
	fileFilter := runtime.FileFilter{
		Pattern: "*.madeProject",
	}

	dialogOptions := runtime.OpenDialogOptions{
		Title:   "Select the project file",
		Filters: []runtime.FileFilter{fileFilter},
	}

	file, err := runtime.OpenFileDialog(pm.ctx, dialogOptions)
	if err != nil || file == "" {
		return "No file selected or error opening file dialog"
	}
	pm.AddProjectToCollectionIfNeeded(file)
	proj, err := CreateFromFile(file)
	if err != nil {
		return "An error occurred while reading the project file"
	}
	pm.currentProject = *proj
	return ""
}

func (pm *ProjectManager) OpenProjectInFileManager(folderPath string) string {
	var cmd *exec.Cmd

	switch goRuntime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", folderPath)
	case "linux":
		cmd = exec.Command("xdg-open", folderPath)
	case "darwin":
		cmd = exec.Command("open", folderPath)
	default:
		return "Error: Unsupported platform"
	}

	err := cmd.Start()
	if err != nil {
		return err.Error()
	}
	return ""
}

func (pm *ProjectManager) ChooseProject() string {
	fileFilter := runtime.FileFilter{
		Pattern: fmt.Sprintf("*.%s", MadeProjectFileExt),
	}
	dialogOptions := runtime.OpenDialogOptions{
		Title:   "Choose a MADE project file",
		Filters: []runtime.FileFilter{fileFilter},
	}

	projectFile, err := runtime.OpenFileDialog(context.Background(), dialogOptions)
	if err != nil {
		return ""
	}
	pm.AddProjectToCollectionIfNeeded(projectFile)
	return projectFile
}

func (pm *ProjectManager) GetInformationToFillCreationForm(folderPath string) ProjectCreationInformation {
	fullPath := filepath.Join(folderPath, "minecraftinstance.json")
	info := ProjectCreationInformation{
		FolderPath: folderPath,
	}
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return info
	}
	contents, err := os.ReadFile(fullPath)
	if err != nil {
		return info
	}
	info.Name = filepath.Base(folderPath)
	version := getInJsonValueOf("minecraftVersion", string(contents))
	if version != "" {
		info.Version = strings.Map(func(r rune) rune {
			if unicode.IsDigit(r) || r == '.' {
				return r
			}
			return -1
		}, version)
	}

	loaderString := getInJsonValueOf("name", string(contents))

	if loaderString != "" {
		loaderString := getPureLoaderValue(loaderString)
		if loaderString != "" {
			loader, err := src.StringToLoader(loaderString)
			if err == nil {
				info.ModLoader = &loader
			}
		}
	}
	return info
}
func getPureLoaderValue(input string) string {
	endIndex := strings.Index(input, "-")
	if endIndex == -1 {
		return ""
	}
	return input[:endIndex]
}
func getInJsonValueOf(key string, jsonData string) string {
	searchFor := "\"" + key + "\": \""
	startIndex := strings.Index(jsonData, searchFor)
	if startIndex == -1 {
		searchFor = "\"" + key + "\":\""
		startIndex = strings.Index(jsonData, searchFor)
		if startIndex == -1 {
			return ""
		}
	}
	jsonData = jsonData[startIndex+len(searchFor):]
	endIndex := strings.Index(jsonData, "\"")
	if endIndex == -1 {
		return ""
	}
	return jsonData[:endIndex]
}
