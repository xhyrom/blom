package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileStatement(statement ast.Statement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	switch statement := statement.(type) {
	case *ast.VariableDeclarationStatement:
		return c.compileVariableDeclaration(statement, function, isReturn)
	case *ast.AssignmentStatement:
		return c.compileAssignmentStatement(statement, function, isReturn)
	case *ast.IdentifierLiteral, *ast.IntLiteral, *ast.FloatLiteral, *ast.CharLiteral, *ast.StringLiteral, *ast.BooleanLiteral:
		return c.compileLiteral(statement, function, vtype, isReturn)
	case *ast.FunctionCall:
		return c.compileFunctionCall(statement, function, vtype)
	case *ast.IfStatement:
		return c.compileCondition(statement, function, vtype, isReturn)
	case *ast.WhileLoopStatement:
		return c.compileLoop(statement, function, vtype, isReturn)
	case *ast.ReturnStatement:
		return c.compileReturnStatement(statement, function, vtype)
	case *ast.BinaryExpression:
		return c.compileBinaryExpression(statement, function, vtype, isReturn)
	case *ast.UnaryExpression:
		return c.compileUnaryExpression(statement, function, vtype, isReturn)
	case *ast.BlockStatement:
		return c.compileBlock(statement, function, vtype, isReturn)
	}

	return nil
}
