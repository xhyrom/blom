package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileCompileTimeFunctionCall(call *ast.CompileTimeFunctionCall, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	exp := call.Parameters[0]
	castType, _ := compiler.ParseType(call.Parameters[1].(*ast.IdentifierLiteral).Value)

	stmt, stmtAdd := c.CompileStatement(exp, scope)

	if exp.Kind() == ast.IdentifierLiteralNode {
		name := fmt.Sprintf("%%conv.%d", c.tempCounter)

		result := make([]string, 0)
		for _, s := range stmt {
			result = append(result, s)
		}

		result = append(result, fmt.Sprintf("%s =%s %s %s", name, c.StoreType(castType), convertOperation(c, stmtAdd.Type, castType), stmtAdd.Name))

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

func convertOperation(c *Compiler, original compiler.Type, new compiler.Type) string {
	if !original.IsInteger() && !new.IsInteger() {
		return "exts" // TODO: finish
	}

	var first, second string
	if original.IsFloatingPoint() {
		first = original.String()
	} else {
		first = "s" + c.StoreType(original).String()
	}

	if new.IsFloatingPoint() {
		second = "f"
	} else {
		second = "si"
	}

	return fmt.Sprintf(
		"%sto%s",
		first,
		second,
	)
}
