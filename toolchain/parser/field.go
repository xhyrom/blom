package parser

import (
	"blom/ast"
)

func (p *Parser) parseField(base ast.Node) *ast.Field {
	p.Consume()

	member := p.parseLiteral()

	return &ast.Field{
		Base:   base,
		Member: member,
		Loc:    member.Location(),
	}
}
