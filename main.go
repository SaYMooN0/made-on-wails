package main

import (
	"embed"
	made "made/structs/theme_related"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	app := NewApp()
	themeCollection := made.InitializeThemeCollection()
	defer themeCollection.SaveToFile()
	err := wails.Run(&options.App{
		Title:  "Made",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			themeCollection,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}

}
