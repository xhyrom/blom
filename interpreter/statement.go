package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretStatement(statement ast.Statement) objects.Object {
	switch statement := statement.(type) {
	case *ast.IdentifierLiteral, *ast.IntLiteral, *ast.FloatLiteral, *ast.CharLiteral, *ast.StringLiteral, *ast.BooleanLiteral:
		return t.interpretLiteral(statement)
	case *ast.FunctionCall:
		return t.interpretFunctionCall(statement)
	}

	return nil
}
