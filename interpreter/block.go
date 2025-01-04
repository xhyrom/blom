package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretBlock(block *ast.BlockStatement, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	t.Scopes.Append()

	for _, statement := range block.Body {
		t.interpretStatement(statement, function, vtype, isReturn)
	}

	t.Scopes.Pop()

	return nil // TODO: closures
}
