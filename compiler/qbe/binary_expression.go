package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, indent int, expectedType *compiler.Type) ([]string, *Additional) {
	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)

	left, leftVar := c.CompileStatement(stmt.Left, indent, expectedType)
	right, rightVar := c.CompileStatement(stmt.Right, indent, expectedType)

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

	exp += " " + leftVar.Name + ", " + rightVar.Name

	result = append(result, exp)

	result = append(result, "# ^ binary expression\n")

	return result, &Additional{
		Name: name,
		Type: leftVar.Type,
	} //fmt.Sprintf("l %s", name)
}
