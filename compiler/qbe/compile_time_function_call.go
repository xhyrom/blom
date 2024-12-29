package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/env"
)

func (c *Compiler) CompileCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	exp := call.Parameters[0]
	castType, _ := compiler.ParseType(call.Parameters[1].(*ast.IdentifierLiteralStatement).Value)

	stmt, stmtAdd := c.CompileStatement(exp, scope)

	ad := &QbeIdentifier{
		Name: stmtAdd.Name,
		Type: castType,
		Raw:  stmtAdd.Raw,
	}

	return stmt, ad
}
