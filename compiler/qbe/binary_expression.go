package qbe

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, indent int) ([]string, string) {
	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)
	c.Environment.TempCounter += 1

	left, leftVar := c.CompileStatement(stmt.Left, indent)
	right, rightVar := c.CompileStatement(stmt.Right, indent)

	result := make([]string, 0)

	for _, l := range left {
		result = append(result, l)
	}

	for _, r := range right {
		result = append(result, r)
	}

	exp := fmt.Sprintf("%s =%s ", name, "l")

	switch stmt.Operator {
	case tokens.Plus:
		exp += "add"
	}

	exp += " " + leftVar[2:] + ", " + rightVar[2:]

	result = append(result, exp)

	result = append(result, "# ^ binary expression\n")

	return result, fmt.Sprintf("l %s", name)
}
