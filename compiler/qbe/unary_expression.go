package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileUnaryExpression(stmt *ast.UnaryExpression, indent int, expectedType *compiler.Type) ([]string, *Additional) {
	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)

	operand, operandVar := c.CompileStatement(stmt.Operand, indent, expectedType)
	result := make([]string, 0)

	for _, r := range operand {
		result = append(result, r)
	}

	exp := fmt.Sprintf("%s =%s mul", name, c.StoreType(operandVar.Type))

	multiplicator := ""
	switch stmt.Operator {
	case tokens.Plus:
		multiplicator = "1"
	case tokens.Minus:
		multiplicator = "-1"
	}

	exp += " " + c.StoreVal(operandVar) + ", " + multiplicator

	result = append(result, exp)

	result = append(result, "# ^ unary expression\n")

	return result, &Additional{
		Name: name,
		Type: operandVar.Type,
	}
}
