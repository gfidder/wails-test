package mibreader

import (
	"log"

	"github.com/alecthomas/repr"
	"github.com/sleepinggenius2/gosmi/parser"
)

func ReadMib(fileName string) {
	module, err := parser.ParseFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	repr.Println(module)
}
