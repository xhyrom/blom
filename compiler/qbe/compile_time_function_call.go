package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	exp := call.Parameters[0]
	castType, _ := compiler.ParseType(call.Parameters[1].(*ast.IdentifierLiteralStatement).Value)

	stmt, stmtAdd := c.CompileStatement(exp, scope)

	if exp.Kind() == ast.IdentifierLiteralNode {
		name := fmt.Sprintf("%%conv.%d", c.tempCounter)

		result := make([]string, 0)
		for _, s := range stmt {
			result = append(result, s)
		}

		result = append(result, fmt.Sprintf("%s =%s exts %s", name, c.StoreType(castType), stmtAdd.Name))

		ad := &QbeIdentifier{
			Name: name,
			Type: castType,
			Raw:  stmtAdd.Raw,
		}

		return result, ad
	}

	ad := &QbeIdentifier{
		Name: stmtAdd.Name,
		Type: castType,
		Raw:  stmtAdd.Raw,
	}

	return stmt, ad
}
