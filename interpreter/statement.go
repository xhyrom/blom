package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretStatement(statement ast.Statement) objects.Object {
	switch statement := statement.(type) {
	case *ast.VariableDeclarationStatement:
		t.interpretVariableDeclaration(statement)
	case *ast.Assignment:
		return t.interpretAssignment(statement)
	case *ast.IdentifierLiteral, *ast.IntLiteral, *ast.FloatLiteral, *ast.CharLiteral, *ast.StringLiteral, *ast.BooleanLiteral:
		return t.interpretLiteral(statement)
	case *ast.FunctionCall:
		return t.interpretFunctionCall(statement)
	case *ast.BuiltinFunctionCall:
		return t.interpretBuiltinFunctionCall(statement)
	case *ast.BlockStatement:
		return t.interpretBlock(statement)
	}

	return nil
}
