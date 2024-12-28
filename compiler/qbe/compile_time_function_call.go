package qbe

import (
	"blom/ast"
	"blom/compiler"
)

func (c *Compiler) CompileCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall, indent int) ([]string, *Additional) {
	exp := call.Parameters[0]
	castType, _ := compiler.ParseType(call.Parameters[1].(*ast.IdentifierLiteralStatement).Value)

	stmt, stmtAdd := c.CompileStatement(exp, indent, &castType)

	ad := &Additional{
		Name: stmtAdd.Name,
		Type: castType,
		Raw:  stmtAdd.Raw,
	}

	return stmt, ad
}
