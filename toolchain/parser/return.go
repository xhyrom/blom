package parser

import (
	"blom/ast"
)

func (p *Parser) parseReturn() *ast.ReturnStatement {
	p.Consume()

	value := p.parseExpression()

	return &ast.ReturnStatement{
		Value: value,
		Loc:   value.Location(),
	}
}
