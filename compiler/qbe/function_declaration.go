package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileFunctionDeclaration(stmt *ast.FunctionDeclaration, scope *env.Environment[*Variable]) []string {
	if stmt.IsNative() {
		return []string{}
	}

	functionScope := env.New[*Variable]()

	result := ""

	if stmt.Name == "main" {
		result += "export "
	}

	result += "function " + stmt.ReturnType.String() + " $" + stmt.Name + "("

	for i, param := range stmt.Arguments {
		if i > 0 {
			result += ", "
		}

		result += fmt.Sprintf("%s %%%s.%d", param.Type, param.Name, i)

		functionScope.Set(param.Name, &Variable{
			Type: param.Type,
			Id:   i,
		})
	}

	result += ") {\n"
	result += "@start\n"

	block, _ := c.CompileBlock(*stmt.Body, functionScope, false)
	for _, b := range block {
		result += b
	}

	result += "}"

	return []string{result, "# ^ function declaration\n"}
}
