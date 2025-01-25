package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
	"blom/scope"
	"fmt"
)

type Interpreter struct {
	Functions map[string]*ast.FunctionDeclaration
	Scopes    *scope.Scopes[objects.Object]
}

func New() *Interpreter {
	return &Interpreter{
		Functions: make(map[string]*ast.FunctionDeclaration),
		Scopes:    scope.NewScopes[objects.Object](),
	}
}

func (t *Interpreter) Interpret(program *ast.Program) {
	for _, primitive := range program.Body {
		t.populateFunctions(primitive)
	}

	t.Scopes.Append()
	t.interpretStatement(&ast.FunctionCall{
		Name:       "main",
		Parameters: []ast.Expression{},
		Loc:        program.Loc,
	}, t.Functions["main"], &t.Functions["main"].ReturnType, false)
	t.Scopes.Pop()
}

func (t *Interpreter) populateFunctions(primitive ast.Statement) {
	switch primitive := primitive.(type) {
	case *ast.FunctionDeclaration:
		t.Functions[primitive.Name] = primitive
	case *ast.TypeDefinition:
	default:
		panic(fmt.Sprintf("'%T' is not a valid primitive", primitive))
	}
}

func (t *Interpreter) convertToType(first ast.Type, second ast.Type, value objects.Object) objects.Object {
	if first.IsFunction() && second.IsPointer() && second.(ast.PointerType).Dereference() == ast.Void {
		return value
	}

	obj := objects.FromType(second)
	obj.SetValue(value.Value())
	return obj
}
