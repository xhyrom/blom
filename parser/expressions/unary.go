package expressions

import "blom/ast"

func ParseUnary(p Parser) (*ast.UnaryExpression, error) {
	operator := p.Consume()
	operand, _ := p.ParsePrimaryExpression()

	return &ast.UnaryExpression{
		Operator: operator.Kind,
		Operand:  operand,
		Loc:      operator.Location,
	}, nil
}
