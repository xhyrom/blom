package expressions

import (
	"blom/ast"
	"blom/debug"
)

func ParseUnary(p Parser) *ast.UnaryExpression {
	operator := p.Consume()
	operand, err := p.ParsePrimaryExpression()

	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), operator.Location.Row, operator.Location.Column+1)
		dbg.ThrowError(err.Error(), true, debug.NewHint("Did you forget to add an operand?", "0"))
	}

	return &ast.UnaryExpression{
		Operator: operator.Kind,
		Operand:  operand,
		Loc:      operator.Location,
	}
}
