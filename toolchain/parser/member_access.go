package parser

import (
	"blom/ast"
	"blom/tokens"
)

func (p *Parser) parseMemberAccess(left ast.Expression) ast.Expression {
	p.Consume()

	right := p.parseLiteral()
	if right, ok := right.(*ast.FunctionCall); ok {
		right.MemberAccess = true
		right.Parameters = append([]ast.Expression{left}, right.Parameters...)

		var exp ast.Expression = right
		if p.Current().Kind == tokens.Dot {
			exp = p.parseMemberAccess(exp)
		}

		return exp
	}

	exp := &ast.MemberAccess{
		Left:  left,
		Right: right,
		Loc:   left.Location(),
	}

	if p.Current().Kind == tokens.Assign {
		return p.parseAssignment(exp)
	}

	return exp
}
