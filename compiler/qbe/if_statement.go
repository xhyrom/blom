package qbe

import (
	"blom/ast"
	"fmt"

	"github.com/gookit/goutil/dump"
)

func (c *Compiler) CompileIfStatement(stmt *ast.IfStatement, indent int) []string {
	dump.P(stmt)

	result := make([]string, 0)

	condition, conditionIdentifier := c.CompileStatement(stmt.Condition, indent, nil)

	for _, data := range condition {
		result = append(result, data)
	}

	id := c.Environment.TempCounter
	c.Environment.TempCounter += 1

	result = append(result, fmt.Sprintf("jnz %s, @if.%d, @else.%d", conditionIdentifier.Name, id, id))

	// if block
	result = append(result, fmt.Sprintf("@if.%d", id))

	thenBlock, _ := c.CompileStatement(stmt.Then, indent-2, nil)
	for _, data := range thenBlock {
		result = append(result, data)
	}

	// else block
	result = append(result, fmt.Sprintf("@else.%d", id))

	if stmt.Else != nil {
		elseBlock, _ := c.CompileStatement(stmt.Else, indent-2, nil)
		for _, data := range elseBlock {
			result = append(result, data)
		}
	}

	// end if
	result = append(result, fmt.Sprintf("@end.%d", id))

	result = append(result, "# ^ if statement")

	return result
}
