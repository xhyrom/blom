package qbe

import (
	"blom/ast"
	"blom/compiler"
)

func (c *Compiler) CompileCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall) ([]string, *Additional) {
	exp := call.Parameters[0]
	castType, _ := compiler.ParseType(call.Parameters[1].(*ast.IdentifierLiteralStatement).Value)

	stmt, stmtAdd := c.CompileStatement(exp, &castType)

	ad := &Additional{
		Name: stmtAdd.Name,
		Type: castType,
		Raw:  stmtAdd.Raw,
	}

	return stmt, ad
}
