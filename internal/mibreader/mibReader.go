package mibreader

import (
	"fmt"
	"strings"

	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/sleepinggenius2/gosmi/types"
	"github.com/willowbrowser/snmpmibbrowser/internal/log"
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
		log.Error(fmt.Sprintf("Error parsing mib: %v", err))
	}

	// load in stuff from already involved OIDs and types for the parsing process

	m.readNewImports(module)

	// need to import the new stuff before going through types
	m.readNewTypes(module)

	m.addIdentity(module.Body.Identity, module.Name.String())
	m.readNewOids(module)

	m.loadedOids.AddNewOids(m.newOids)

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
		log.Info(fmt.Sprintf("Import is %s\n", newImport.Module.String()))
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
		log.Info(fmt.Sprintf("New type is %s\n", newType.Name.String()))
	}
}

func (m *MibReader) addIdentity(moduleIdentity *parser.ModuleIdentity, mibName string) {
	parentOid := m.loadedOids.FindDirectParent(moduleIdentity.Oid.SubIdentifiers[0].Name.String())
	oidNum := appendOidNumber(parentOid.OID, *moduleIdentity.Oid.SubIdentifiers[1].Number)

	newStoredOid := oidstorage.CreateNewOid(moduleIdentity.Name.String(), oidNum, mibName)
	newStoredOid.Description = moduleIdentity.Description
	m.newOids = append(m.newOids, newStoredOid)
	parentOid.AddChildren(&newStoredOid)
}

func (m *MibReader) readNewOids(module *parser.Module) {
	// TODO : fix this up
	mibName := module.Name.String()
	newStoredOid := oidstorage.CreateNewOid("", "", "")

	for _, newOid := range module.Body.Nodes {
		parentName := newOid.Oid.SubIdentifiers[0].Name.String()
		parentOid := m.loadedOids.FindDirectParent(parentName)

		if parentOid == nil {
			parentOid = m.findParentInNewOids(parentName)

			if parentOid == nil {
				log.Warn(fmt.Sprintf("No suitable parent found for oid: %s\n", newOid.Name.String()))
			}
		}

		if newOid.ObjectIdentifier {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newStoredOid.Name = newOid.Name.String()
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			m.newOids = append(m.newOids, newOidStore)

		} else if newOid.ObjectIdentity != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.ObjectIdentity.Description
			newOidStore.Status = newOid.ObjectIdentity.Status.ToSmi().String()
			newOidStore.Type = oidstorage.ObjectIdentity
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else {
			fmt.Printf("Still need some work to properly add oid: %s\n", newOid.Name.String())
		}
	}
}

func (m *MibReader) findParentInNewOids(parentName string) *oidstorage.Oid {
	var parentOid *oidstorage.Oid

	for _, oid := range m.newOids {
		if oid.Name == parentName {
			parentOid = &oid
			break
		}
	}

	return parentOid
}

// TODO : some way to overwrite duplicates. Assuming the mib is an updated mib

func appendOidNumber(oidNumStr string, newOidNum types.SmiSubId) string {
	var sb strings.Builder

	sb.WriteString(oidNumStr)
	sb.WriteString(fmt.Sprintf(".%d", newOidNum))

	return sb.String()
}
