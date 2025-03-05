package parser

import "blom/ast"

func (p *Parser) parseMemberAccess(left ast.Statement) *ast.MemberAccess {
	p.Consume()

	member := p.parseLiteral()

	return &ast.MemberAccess{
		Left:  left,
		Right: member,
		Loc:   left.Location(),
	}
}
