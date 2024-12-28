package qbe

import (
	"blom/ast"
	"fmt"
)

func (c *Compiler) CompileWhileLoopStatement(stmt *ast.WhileLoopStatement) ([]string, *Additional) {
	result := make([]string, 0)

	id := c.Environment.TempCounter
	c.Environment.TempCounter += 1

	result = append(result, fmt.Sprintf("@loop.cond.%d", id))

	condition, conditionIdentifier := c.CompileStatement(stmt.Condition, nil)

	for _, data := range condition {
		result = append(result, data)
	}

	result = append(result, fmt.Sprintf("jnz %s, @loop.body.%d, @loop.end.%d", conditionIdentifier.Name, id, id))

	// loop body
	result = append(result, fmt.Sprintf("@loop.body.%d", id))

	body, _ := c.CompileBlock(*stmt.Body, false)

	for _, data := range body {
		result = append(result, data)
	}

	result = append(result, fmt.Sprintf("jmp @loop.cond.%d", id))

	// loop end
	result = append(result, fmt.Sprintf("@loop.end.%d", id))

	return result, &Additional{
		Id: id,
	}
}
