package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileBuiltinFunctionCall(call *ast.BuiltinFunctionCall, function *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	switch call.Name {
	case "cast":
		return compileCastFunctionCall(c, call, function, vtype)
	}

	panic("Unknown builtin function call")
}

func compileCastFunctionCall(c *Compiler, call *ast.BuiltinFunctionCall, function *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	if len(call.Parameters) < 2 {
		panic("Cast function requires at least two parameters")
	}

	firstParam := call.Parameters[0]
	if firstParam.Kind() != ast.IdentifierLiteralNode {
		panic("First parameter of cast function must be an type")
	}

	typeName := firstParam.(*ast.IdentifierLiteral).Value
	astType, err := ast.ParseType(typeName, map[string]ast.Type{})
	if err != nil {
		panic("Invalid type name")
	}

	t := qbe.RemapAstType(astType)

	stmt := c.compileStatement(call.Parameters[1], function, vtype, false)

	return c.convertToType(
		stmt.Type,
		t,
		stmt.Value,
		function,
	)
}
