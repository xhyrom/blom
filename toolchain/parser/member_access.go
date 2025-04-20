package parser

import (
	"blom/ast"
)

func (p *Parser) parseMemberAccess(left ast.Statement) ast.Expression {
	p.Consume()

	member := p.parseLiteral()
	return &ast.MemberAccess{
		Left:  left,
		Right: member,
		Loc:   left.Location(),
	}
}
