package manager

import (
	"blom/ast"
	"fmt"
	"strconv"
	"strings"
)

type FunctionSignature struct {
	name       string
	parameters []ast.Type
}

func (f *FunctionSignature) Key() string {
	types := make([]string, len(f.parameters))
	for i, t := range f.parameters {
		types[i] = strconv.Itoa(int(t))
	}

	return fmt.Sprintf("%s(%s)", f.name, strings.Join(types, ","))
}

type FunctionManager struct {
	functions map[string]*ast.FunctionDeclaration
}

func NewFunctionManager() *FunctionManager {
	return &FunctionManager{
		functions: make(map[string]*ast.FunctionDeclaration),
	}
}

func (m *FunctionManager) Register(fun *ast.FunctionDeclaration) {
	sig := FunctionSignature{
		name:       fun.Name,
		parameters: make([]ast.Type, len(fun.Arguments)),
	}

	for i, arg := range fun.Arguments {
		sig.parameters[i] = arg.Type
	}

	m.functions[sig.Key()] = fun
}

func (m *FunctionManager) Get(name string, arguments []ast.Type) (*ast.FunctionDeclaration, bool) {
	sig := FunctionSignature{
		name:       name,
		parameters: arguments,
	}

	fun, ok := m.functions[sig.Key()]
	return fun, ok
}

func (m *FunctionManager) GetByDeclaration(fun *ast.FunctionDeclaration) (*ast.FunctionDeclaration, bool) {
	arguments := make([]ast.Type, len(fun.Arguments))
	for i, arg := range fun.Arguments {
		arguments[i] = arg.Type
	}

	return m.Get(fun.Name, arguments)
}

func (m *FunctionManager) ContainsByName(name string) bool {
	for key := range m.functions {
		if strings.HasPrefix(key, name+"(") {
			return true
		}
	}

	return false
}
