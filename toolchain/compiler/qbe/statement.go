package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileStatement(statement ast.Node, function *qbe.Function, vtype qbe.Type, isReturn bool) *qbe.TypedValue {
	switch statement := statement.(type) {
	case *ast.VariableDeclaration:
		return c.compileVariableDeclaration(statement, function, isReturn)
	case *ast.Assignment:
		return c.compileAssignmentStatement(statement, function, isReturn)
	case *ast.IdentifierLiteral, *ast.IntLiteral, *ast.FloatLiteral, *ast.CharLiteral, *ast.StringLiteral, *ast.BooleanLiteral:
		return c.compileLiteral(statement, function, vtype, isReturn)
	case *ast.FunctionCall, *ast.MethodCall, *ast.InfixCall:
		return c.compileCall(statement.(ast.Call), function, vtype)
	case *ast.If:
		return c.compileCondition(statement, function, vtype, isReturn)
	case *ast.WhileLoop:
		return c.compileLoop(statement, function, vtype, isReturn)
	case *ast.Return:
		return c.compileReturnStatement(statement, function, vtype)
	case *ast.BinaryExpression:
		return c.compileBinaryExpression(statement, function, vtype, isReturn)
	case *ast.UnaryExpression:
		return c.compileUnaryExpression(statement, function, vtype, isReturn)
	case *ast.Block:
		return c.compileBlock(statement, function, vtype, isReturn)
	case *ast.GroupedExpression:
		return c.compileStatement(statement.Expression, function, vtype, isReturn)
	}

	return nil
}
