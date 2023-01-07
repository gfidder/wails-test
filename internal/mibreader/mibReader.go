package mibreader

import (
	"fmt"
	"log"

	"github.com/sleepinggenius2/gosmi/parser"
)

func ReadMib(fileName string) {
	imports := []parser.Import{}

	module, err := parser.ParseFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// need information on what mibs we have loaded in, just in case we already have the proper information
	for _, moduleImport := range module.Body.Imports {
		duplicateModule := false
		for i := range imports {
			if imports[i].Module == moduleImport.Module {
				duplicateModule = true
			}
		}

		// probably have to some recursive search for all imports to make sure we have them all

		if !duplicateModule {
			imports = append(imports, moduleImport)
		}
	}

	// TODO : if the imports are not located in the same folder as the mib, need to throw an error
	// by emitting an event for Vue to deal with

	fmt.Println("Did it")
}
