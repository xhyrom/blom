package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
)

func (c *Compiler) CompileIfStatement(stmt *ast.IfStatement, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	result := make([]string, 0)

	condition, conditionIdentifier := c.CompileStatement(stmt.Condition, scope)

	for _, data := range condition {
		result = append(result, data)
	}

	id := c.tempCounter

	if stmt.HasElse() {
		result = append(result, fmt.Sprintf("jnz %s, @if.%d, @else.%d", conditionIdentifier.Name, id, id))
	} else {
		result = append(result, fmt.Sprintf("jnz %s, @if.%d, @end.%d", conditionIdentifier.Name, id, id))
	}

	// if block
	result = append(result, fmt.Sprintf("@if.%d", id))

	thenBlock, _ := c.CompileBlock(*stmt.Then, scope, false)
	for _, data := range thenBlock {
		result = append(result, data)
	}

	// check if result doesn't contain return statement
	containsRet := false
	for _, data := range stmt.Then.Body {
		if data.Kind() == ast.ReturnNode {
			containsRet = true
		}
	}

	if !containsRet {
		result = append(result, fmt.Sprintf("jmp @end.%d", id))
	}

	// else block
	if stmt.HasElse() {
		result = append(result, fmt.Sprintf("@else.%d", id))

		elseBlock, _ := c.CompileBlock(*stmt.Else, scope, false)
		for _, data := range elseBlock {
			result = append(result, data)
		}

		// check if result doesn't contain return statement
		containsRet := false
		for _, data := range stmt.Else.Body {
			if data.Kind() == ast.ReturnNode {
				containsRet = true
			}
		}

		if !containsRet {
			result = append(result, fmt.Sprintf("jmp @end.%d", id))
		}
	}

	// end if
	result = append(result, fmt.Sprintf("@end.%d", id))

	result = append(result, "# ^ if statement")

	return result, &QbeIdentifier{
		Id: id,
	}
}
