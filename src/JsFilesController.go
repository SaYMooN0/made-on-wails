package src

import (
	"bufio"
	"fmt"
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

func DeleteAction(actionId, filePath string) error {
	tempFilePath := filepath.Join(os.TempDir(), "tempfile")
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	skipNextLine := false
	for scanner.Scan() {
		line := scanner.Text()
		if !skipNextLine && !strings.Contains(line, fmt.Sprintf("//Made:%s;", actionId)) {
			_, err := tempFile.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("error writing to temp file: %v", err)
			}
		} else if strings.Contains(line, fmt.Sprintf("//Made:%s;", actionId)) {
			skipNextLine = true
			continue
		} else if skipNextLine {
			skipNextLine = false
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	tempFile.Close()
	file.Close()
	err = os.Rename(tempFilePath, filePath)
	if err != nil {
		return fmt.Errorf("error replacing the file: %v", err)
	}

	return nil
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
