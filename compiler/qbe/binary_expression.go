package qbe

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"blom/tokens"
	"fmt"
)

func (c *Compiler) CompileBinaryExpression(stmt *ast.BinaryExpression, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	name := fmt.Sprintf("%%tmp.%d", c.tempCounter)

	left, leftVar := c.CompileStatement(stmt.Left, scope)
	right, rightVar := c.CompileStatement(stmt.Right, scope)

	if leftVar.Type != rightVar.Type {
		dbg := debug.NewSourceLocation(c.source, stmt.OperatorLoc.Row, stmt.OperatorLoc.Column)
		dbg.ThrowError(fmt.Sprintf("Cannot perform binary operation on two different types \"%s\" and \"%s\"!", leftVar.Type.Inspect(), rightVar.Type.Inspect()), true)
	}

	result := make([]string, 0)

	for _, l := range left {
		result = append(result, l)
	}

	for _, r := range right {
		result = append(result, r)
	}

	exp := fmt.Sprintf("%s =%s ", name, c.StoreType(leftVar.Type))

	switch stmt.Operator {
	case tokens.Plus:
		exp += "add"
	case tokens.Minus:
		exp += "sub"
	case tokens.Asterisk:
		exp += "mul"
	case tokens.Slash:
		exp += "div"
	case tokens.PercentSign:
		exp += "rem"
	case tokens.Ampersand:
		exp += "and"
	case tokens.VerticalLine:
		exp += "or"
	case tokens.CircumflexAccent:
		exp += "xor"
	case tokens.DoubleLessThan:
		exp += "shl"
	case tokens.DoubleGreaterThan:
		exp += "shr"
	case tokens.Equals:
		exp += "ceq" + c.StoreType(leftVar.Type).String()
	case tokens.LessThan:
		exp += "cslt" + c.StoreType(leftVar.Type).String()
	case tokens.LessThanOrEqual:
		exp += "csle" + c.StoreType(leftVar.Type).String()
	case tokens.GreaterThan:
		exp += "csgt" + c.StoreType(leftVar.Type).String()
	case tokens.GreaterThanOrEqual:
		exp += "csge" + c.StoreType(leftVar.Type).String()
	}

	exp += " " + c.StoreVal(leftVar) + ", " + c.StoreVal(rightVar)

	result = append(result, exp)

	result = append(result, "# ^ binary expression\n")

	return result, &QbeIdentifier{
		Name: name,
		Type: leftVar.Type,
	}
}
