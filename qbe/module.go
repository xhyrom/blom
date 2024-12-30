package qbe

import "strings"

// A complete QBE module
type Module struct {
	Functions []Function
	Types     []TypeDefinition
	Data      []Data
}

func NewModule() Module {
	return Module{
		Functions: make([]Function, 0),
		Types:     make([]TypeDefinition, 0),
		Data:      make([]Data, 0),
	}
}

func (m *Module) AddFunction(f Function) {
	m.Functions = append(m.Functions, f)
}

func (m *Module) GetFunctionByName(name string) *Function {
	for i, f := range m.Functions {
		if f.Name == name {
			return &m.Functions[i]
		}
	}

	return nil
}

func (m *Module) AddType(t TypeDefinition) {
	m.Types = append(m.Types, t)
}

func (m *Module) AddData(d Data) {
	m.Data = append(m.Data, d)
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
