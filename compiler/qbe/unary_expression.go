package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileUnaryExpression(stmt *ast.UnaryExpression, expectedType *compiler.Type) ([]string, *Additional) {
	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)

	operand, operandVar := c.CompileStatement(stmt.Operand, expectedType)
	result := make([]string, 0)

	for _, r := range operand {
		result = append(result, r)
	}

	operation := ""
	right := ""
	switch stmt.Operator {
	case tokens.Plus:
		operation = "mul"
		right = "1"
	case tokens.Minus:
		operation = "mul"
		right = "-1"
	case tokens.Tilde:
		operation = "xor"
		right = "-1"
	}

	exp := fmt.Sprintf("%s =%s %s", name, c.StoreType(operandVar.Type), operation)

	exp += " " + c.StoreVal(operandVar) + ", " + right

	result = append(result, exp)

	result = append(result, "# ^ unary expression\n")

	return result, &Additional{
		Name: name,
		Type: operandVar.Type,
	}
}
