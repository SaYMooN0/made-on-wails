package src

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const (
	kubejs             = "kubejs"
	serverScripts      = "server_scripts"
	startupScripts     = "startup_scripts"
	clientScripts      = "client_scripts"
	assets             = "assets"
	vanillaRecipesFile = "vanilla"
	textures           = "textures"
)

func WriteVanillaRecipe(contentToWrite, projectFolderPath, actionId string) {
	fullDirPath := filepath.Join(projectFolderPath, kubejs, serverScripts)
	os.MkdirAll(fullDirPath, 0755)
	fullFilePath := filepath.Join(fullDirPath, vanillaRecipesFile+".js")
	contentToWrite = "\n" + GenerateMadeComment(actionId) + "\n" + WrapInOnEventRecipes(contentToWrite)
	file, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(contentToWrite)
}

func DeleteAction(actionId, filePath string) {
	tempFilePath := filepath.Join(os.TempDir(), "tempfile")
	skipNextLine := false

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)

	for scanner.Scan() {
		line := scanner.Text()
		if !skipNextLine && !strings.Contains(line, "//Made:"+actionId) {
			writer.WriteString(line + "\n")
		} else if strings.Contains(line, "//Made:"+actionId) {
			skipNextLine = true
			continue
		} else if skipNextLine {
			skipNextLine = false
		}
	}
	writer.Flush()

	os.Remove(filePath)
	os.Rename(tempFilePath, filePath)
}

func GetFullVanillaPath() string {
	return filepath.Join(kubejs, serverScripts, vanillaRecipesFile+".js")
}

func GetFullTextureItemPath() string {
	return filepath.Join(kubejs, assets, kubejs, textures, "item")
}

func WrapInOnEventRecipes(input string) string {
	return "onEvent('recipes', event => {" + input + "}); "
}

func GenerateMadeComment(actionId string) string {
	return "//Made:" + actionId + ";"
}
