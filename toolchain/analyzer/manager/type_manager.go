package manager

import "blom/ast"

type TypeManager struct {
	types map[string]ast.Type
}

func NewTypeManager() *TypeManager {
	return &TypeManager{
		types: make(map[string]ast.Type),
	}
}

func (m *TypeManager) Register(name string, t ast.Type) {
	m.types[name] = t
}

func (m *TypeManager) Get(name string) (ast.Type, bool) {
	t, ok := m.types[name]
	return t, ok
}

func (m *TypeManager) Has(name string) bool {
	_, ok := m.types[name]
	return ok
}

func (m *TypeManager) Types() map[string]ast.Type {
	return m.types
}
