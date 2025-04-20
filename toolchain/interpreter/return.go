package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretReturnStatement(statement *ast.Return, function *ast.FunctionDeclaration, vtype *ast.Type) objects.Object {
	return t.interpretStatement(statement.Value, function, vtype, true)
}
