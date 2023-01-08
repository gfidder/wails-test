package mibreader

import (
	"fmt"
	"log"
	"strings"

	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/sleepinggenius2/gosmi/types"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

type MibReader struct {
	imports    []parser.Import
	types      []parser.Type
	loadedOids oidstorage.LoadedOids
	newOids    []oidstorage.Oid
}

func NewMibReader(loadedOids oidstorage.LoadedOids) *MibReader {
	// maybe when we return the new mibReader, load in the already loaded Oids

	return &MibReader{
		imports:    []parser.Import{},
		types:      []parser.Type{},
		loadedOids: loadedOids,
		newOids:    []oidstorage.Oid{},
	}
}

func (m *MibReader) ReadMib(fileName string) {

	module, err := parser.ParseFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// load in stuff from already involved OIDs and types for the parsing process

	m.readNewImports(module)

	// need to import the new stuff before going through types
	m.readNewTypes(module)

	m.addIdentity(module.Body.Identity, module.Name.String())
	m.readNewOids(module)

	// TODO : if the imports are not located in the same folder as the mib, need to throw an error
	// by emitting an event for Vue to deal with

	fmt.Println("Did it")
}

func (m *MibReader) readNewImports(module *parser.Module) {
	// need information on what mibs we have loaded in, just in case we already have the proper information
	// worry about this last. Don't even know how we are going to store it yet
	for _, moduleImport := range module.Body.Imports {
		duplicateModule := false
		for i := range m.imports {
			if m.imports[i].Module == moduleImport.Module {
				duplicateModule = true
			}
		}

		// probably have to some recursive search for all imports to make sure we have them all

		if !duplicateModule {
			m.imports = append(m.imports, moduleImport)
		}
	}

	for _, newImport := range m.imports {
		fmt.Printf("Import is %s\n", newImport.Module.String())
	}
}

func (m *MibReader) readNewTypes(module *parser.Module) {
	for _, moduleType := range module.Body.Types {
		duplicateType := false
		for i := range m.types {
			if m.types[i].Name == moduleType.Name {
				duplicateType = true
			}
		}

		// probably have to some recursive search for all imports to make sure we have them all

		if !duplicateType {
			m.types = append(m.types, moduleType)
		}
	}

	for _, newType := range m.types {
		fmt.Printf("New type is %s\n", newType.Name.String())
	}
}

func (m *MibReader) addIdentity(moduleIdentity *parser.ModuleIdentity, mibName string) {
	parentOid := m.loadedOids.FindDirectParent(moduleIdentity.Oid.SubIdentifiers[0].Name.String())
	oidNum := appendOidNumber(parentOid.OID, *moduleIdentity.Oid.SubIdentifiers[1].Number)

	newStoredOid := oidstorage.CreateNewOid(moduleIdentity.Name.String(), oidNum, mibName)
	m.newOids = append(m.newOids, newStoredOid)
	parentOid.AddChildren(&newStoredOid)
}

func (m *MibReader) readNewOids(module *parser.Module) {
	// TODO : fix this up
	newStoredOid := oidstorage.CreateNewOid("", "", "")

	for _, newOid := range module.Body.Nodes {
		fmt.Println(newOid.Name.String())

		if newOid.ObjectIdentifier {
			// parentOid := m.loadedOids.FindDirectParent(newOid.Oid.SubIdentifiers[0].Name.String())
			newStoredOid.Name = newOid.Name.String()
		} else {

		}
	}
}

func appendOidNumber(oidNumStr string, newOidNum types.SmiSubId) string {
	var sb strings.Builder

	sb.WriteString(oidNumStr)
	sb.WriteString(fmt.Sprintf(".%d", newOidNum))

	return sb.String()
}
