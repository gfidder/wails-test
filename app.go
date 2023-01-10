package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alecthomas/repr"
	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/willowbrowser/snmpmibbrowser/internal/mibreader"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

// App struct
type App struct {
	ctx        context.Context
	mibReader  *mibreader.MibReader
	loadedOids *oidstorage.LoadedOids
}

// NewApp creates a new App application struct
func NewApp() *App {
	loadedOids := oidstorage.NewLoadedOids()
	mibReader := mibreader.NewMibReader(loadedOids)

	fmt.Printf("Address of loadedOids:\t%p\n", loadedOids)
	fmt.Printf("Address of mibReader:\t%p\n", mibReader)

	mibReader.PrintLoadedOidsAddress()

	return &App{
		mibReader:  mibReader,
		loadedOids: loadedOids,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	fmt.Printf("Address of a.loadedOids:\t%p\n", a.loadedOids)
	fmt.Printf("Address of a.mibReader:\t%p\n", a.mibReader)
}

func (a *App) ParseMib() {
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: "./",
		Title:            "Select a mib",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Mib Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	if result != "" {
		module, err := parser.ParseFile(result)
		if err != nil {
			log.Fatalln(err)
		}

		repr.Println(module)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetCurrentOids() []oidstorage.Oid {
	// apparently this loaded oids isn't the same as the one in the mib reader class
	return a.loadedOids.GetLoadedOids()
}
