package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileWhileLoopStatement(stmt *ast.WhileLoopStatement, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	result := make([]string, 0)

	id := c.tempCounter

	result = append(result, fmt.Sprintf("@loop.cond.%d", id))

	condition, conditionIdentifier := c.CompileStatement(stmt.Condition, scope)

	for _, data := range condition {
		result = append(result, data)
	}

	result = append(result, fmt.Sprintf("jnz %s, @loop.body.%d, @loop.end.%d", conditionIdentifier.Name, id, id))

	// loop body
	result = append(result, fmt.Sprintf("@loop.body.%d", id))

	body, _ := c.CompileBlock(*stmt.Body, scope, false)

	for _, data := range body {
		result = append(result, data)
	}

	result = append(result, fmt.Sprintf("jmp @loop.cond.%d", id))

	// loop end
	result = append(result, fmt.Sprintf("@loop.end.%d", id))

	return result, &QbeIdentifier{
		Id: id,
	}
}
