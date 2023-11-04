package main

import (
	"embed"
	projectRelated "made/src/project_related"
	themeRelated "made/src/theme_related"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	themeCollection := themeRelated.InitializeThemeCollection()
	projectManager := projectRelated.InitializeProjectManager()
	defer themeCollection.SaveToFile()
	defer projectManager.SaveToFile()
	err := wails.Run(&options.App{
		Title:     "Made",
		Width:     900,
		Height:    600,
		MinWidth:  720,
		MinHeight: 540,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        projectManager.SetStartupOnContext,
		Bind: []interface{}{
			app,
			themeCollection,
			projectManager,
		},
	})
	if err != nil {
		errMessage := "Error: " + err.Error() + "\n"
		err := os.WriteFile("error.txt", []byte(errMessage), 0644)
		if err != nil {
			println("Failed to write to error.txt:", err.Error())
		}
	}

}
