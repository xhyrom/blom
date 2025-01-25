package manager

import (
	"blom/ast"
	"fmt"
)

type FunctionManager struct {
	functions map[string][]*ast.FunctionDeclaration
}

func NewFunctionManager() *FunctionManager {
	return &FunctionManager{
		functions: make(map[string][]*ast.FunctionDeclaration),
	}
}

func (m *FunctionManager) Register(fun *ast.FunctionDeclaration) {
	if _, ok := m.functions[fun.Name]; !ok {
		m.functions[fun.Name] = make([]*ast.FunctionDeclaration, 0)
	}

	m.functions[fun.Name] = append(m.functions[fun.Name], fun)
}

func (m *FunctionManager) Get(name string, arguments []ast.Type) (*ast.FunctionDeclaration, bool) {
	if functions, ok := m.functions[name]; ok {
		for _, fun := range functions {
			if len(fun.Arguments) == len(arguments) {
				matches := true
				for i, arg := range fun.Arguments {
					if !arg.Type.Equal(arguments[i]) {
						matches = false
						break
					}
				}

				if matches {
					return fun, true
				}
			}
		}
	}

	return nil, false
}

func (m *FunctionManager) GetAllNamed(name string) []*ast.FunctionDeclaration {
	return m.functions[name]
}

func (m *FunctionManager) GetByDeclaration(fun *ast.FunctionDeclaration) (*ast.FunctionDeclaration, bool) {
	arguments := make([]ast.Type, len(fun.Arguments))
	for i, arg := range fun.Arguments {
		arguments[i] = arg.Type
	}

	return m.Get(fun.Name, arguments)
}

func (m *FunctionManager) GetNewName(fun *ast.FunctionDeclaration) string {
	functions := m.functions[fun.Name]

	if functions == nil || len(functions) == 1 {
		return fun.Name
	}

	index := 1
	for _, f := range functions {
		if f == fun {
			break
		}

		index++
	}

	return fmt.Sprintf("%s.%d", fun.Name, index)
}

func (m *FunctionManager) Has(name string) bool {
	_, ok := m.functions[name]
	return ok
}
