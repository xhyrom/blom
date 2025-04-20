package manager

import (
	"blom/ast"
)

type FunctionWithSuffix struct {
	*ast.FunctionDeclaration
	Suffix string
}

type FunctionManager struct {
	functions map[string][]*ast.FunctionDeclaration
}

func NewFunctionManager() *FunctionManager {
	return &FunctionManager{
		functions: make(map[string][]*ast.FunctionDeclaration),
	}
}

func (m *FunctionManager) Register(fun *ast.FunctionDeclaration) {
	if _, ok := m.functions[fun.Name.Value]; !ok {
		m.functions[fun.Name.Value] = make([]*ast.FunctionDeclaration, 0)
	}

	m.functions[fun.Name.Value] = append(m.functions[fun.Name.Value], fun)
}

func (m *FunctionManager) GetWithIndex(name string, params []ast.Type) (*ast.FunctionDeclaration, int, bool) {
	if functions, ok := m.functions[name]; ok {
		for i, fun := range functions {
			if len(fun.Params) == len(params) {
				matches := true
				for j, param := range fun.Params {
					if !param.Type.Equal(params[j]) {
						matches = false
						break
					}
				}

				if matches {
					return fun, i, true
				}
			}
		}
	}

	return nil, -1, false
}

func (m *FunctionManager) GetByDeclarationWithIndex(fun *ast.FunctionDeclaration) (*ast.FunctionDeclaration, int, bool) {
	params := make([]ast.Type, len(fun.Params))
	for i, arg := range fun.Params {
		params[i] = arg.Type
	}

	return m.GetWithIndex(fun.Name.Value, params)
}

func (m *FunctionManager) Get(name string, params []ast.Type) (*ast.FunctionDeclaration, bool) {
	if fun, _, ok := m.GetWithIndex(name, params); ok {
		return fun, true
	}

	return nil, false
}

func (m *FunctionManager) GetByDeclaration(fun *ast.FunctionDeclaration) (*ast.FunctionDeclaration, bool) {
	if fun, _, ok := m.GetByDeclarationWithIndex(fun); ok {
		return fun, true
	}

	return nil, false
}

func (m *FunctionManager) Has(name string) bool {
	_, ok := m.functions[name]
	return ok
}
