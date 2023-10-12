package main

import (
	"embed"
	themeRelated "made/src/theme_related"

	// projectRelated "made/src/project_related"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	app := NewApp()
	themeCollection := themeRelated.InitializeThemeCollection()
	err := wails.Run(&options.App{
		Title:     "Made",
		Width:     900,
		Height:    600,
		MinWidth:  700,
		MinHeight: 500,
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
