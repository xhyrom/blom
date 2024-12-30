package qbe

import "strings"

// A complete QBE module
type Module struct {
	Functions []Function
	Types     []TypeDefinition
	Data      []Data
}

func (m Module) String() string {
	var parts []string

	for _, typ := range m.Types {
		parts = append(parts, typ.String())
	}

	for _, data := range m.Data {
		parts = append(parts, data.String())
	}

	for _, function := range m.Functions {
		parts = append(parts, function.String())
	}

	return strings.Join(parts, "\n")
}
