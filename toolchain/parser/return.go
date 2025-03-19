package parser

import (
	"blom/ast"
)

func (p *Parser) parseReturn() *ast.Return {
	p.Consume()

	value := p.parseExpression()

	return &ast.Return{
		Value: value,
		Loc:   value.Location(),
	}
}
