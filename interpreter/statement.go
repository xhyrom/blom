package interpreter

import (
	"blom/ast"
	"blom/interpreter/objects"
)

func (t *Interpreter) interpretStatement(statement ast.Statement, function *ast.FunctionDeclaration, vtype *ast.Type, isReturn bool) objects.Object {
	switch statement := statement.(type) {
	case *ast.VariableDeclarationStatement:
		t.interpretVariableDeclaration(statement, function, isReturn)
	case *ast.Assignment:
		return t.interpretAssignment(statement, function, isReturn)
	case *ast.IdentifierLiteral, *ast.IntLiteral, *ast.FloatLiteral, *ast.CharLiteral, *ast.StringLiteral, *ast.BooleanLiteral:
		return t.interpretLiteral(statement, function, vtype, isReturn)
	case *ast.FunctionCall:
		return t.interpretFunctionCall(statement, function, vtype)
	case *ast.BuiltinFunctionCall:
		return t.interpretBuiltinFunctionCall(statement, function, vtype)
	case *ast.If:
		return t.interpretCondition(statement, function, vtype, isReturn)
	case *ast.WhileLoopStatement:
		return t.interpretLoop(statement, function, vtype, isReturn)
	case *ast.ReturnStatement:
		return t.interpretReturnStatement(statement, function, vtype)
	case *ast.BinaryExpression:
		return t.interpretBinaryExpression(statement, function, vtype, isReturn)
	case *ast.UnaryExpression:
		return t.interpretUnaryExpression(statement, function, vtype, isReturn)
	case *ast.BlockStatement:
		return t.interpretBlock(statement, function, vtype, isReturn)
	}

	return nil
}
