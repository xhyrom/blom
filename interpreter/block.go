package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretBlock(block *ast.BlockStatement) objects.Object {
	t.Scopes.Append()

	for _, statement := range block.Body {
		t.interpretStatement(statement)
	}

	t.Scopes.Pop()

	return nil // TODO: closures
}
