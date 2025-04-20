package manager

import (
	"blom/ast"
	"strconv"
)

type Function struct {
	*ast.FunctionDeclaration
	Suffix string
}

type FunctionManager struct {
	functions map[string][]*Function
}

func NewFunctionManager() *FunctionManager {
	return &FunctionManager{
		functions: make(map[string][]*Function),
	}
}

func (m *FunctionManager) Register(fun *ast.FunctionDeclaration) {
	if _, ok := m.functions[fun.Name.Value]; !ok {
		m.functions[fun.Name.Value] = make([]*Function, 0)
	}

	m.functions[fun.Name.Value] = append(m.functions[fun.Name.Value], &Function{
		FunctionDeclaration: fun,
		Suffix:              "",
	})
}

func (m *FunctionManager) Get(name string, params []ast.Type) (*Function, bool) {
	if functions, ok := m.functions[name]; ok {
		for _, fun := range functions {
			if len(fun.Params) == len(params) {
				matches := true
				for i, param := range fun.Params {
					if !param.Type.Equal(params[i]) {
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

func (m *FunctionManager) GetDeclaration(name string, params []ast.Type) (*ast.FunctionDeclaration, bool) {
	function, ok := m.Get(name, params)
	if !ok {
		return nil, false
	}

	return function.FunctionDeclaration, true
}

func (m *FunctionManager) GetByDeclaration(fun *ast.FunctionDeclaration) (*Function, bool) {
	params := make([]ast.Type, len(fun.Params))
	for i, arg := range fun.Params {
		params[i] = arg.Type
	}

	return m.Get(fun.Name.Value, params)
}

func (m *FunctionManager) GetDeclarationByDeclaration(fun *ast.FunctionDeclaration) (*ast.FunctionDeclaration, bool) {
	function, ok := m.GetByDeclaration(fun)
	if !ok {
		return nil, false
	}

	return function.FunctionDeclaration, true
}

func (m *FunctionManager) GetFuncNameSuffix(fun *ast.FunctionDeclaration) string {
	functions := m.functions[fun.Name.Value]

	if functions == nil || len(functions) == 1 {
		return ""
	}

	index := 1
	for _, f := range functions {
		if f.FunctionDeclaration == fun {
			break
		}

		index++
	}

	return strconv.Itoa(index)
}

func (m *FunctionManager) Has(name string) bool {
	_, ok := m.functions[name]
	return ok
}
