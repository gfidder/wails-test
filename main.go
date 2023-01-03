package main

import (
	"changeme/internal/mibreader"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(cd *menu.CallbackData) {
		file, err := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{
			DefaultDirectory: "./",
			Title:            "Select File",
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Text (*.txt)",
					Pattern:     "*.txt",
				},
			},
		})

		if err != nil {
			log.Fatal(err)
		}

		if file != "" {
			mibreader.ReadMib(file)
		}
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(cd *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "test-wails",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Menu:             AppMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
