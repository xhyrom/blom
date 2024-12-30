package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileStatement(statement ast.Statement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	switch statement := statement.(type) {
	case *ast.IdentifierLiteral, *ast.IntLiteral, *ast.FloatLiteral, *ast.StringLiteral, *ast.BooleanLiteral:
		return c.compileLiteral(statement, function, vtype, isReturn)
	case *ast.FunctionCall:
		return c.compileFunctionCall(statement, function)
	case *ast.ReturnStatement:
		return c.compileReturnStatement(statement, function, vtype)
	}

	return nil
}
