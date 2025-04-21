package parser

import (
	"blom/ast"
	"blom/debug"
)

func (p *Parser) parseReturn() *ast.Return {
	p.Consume()

	value := p.parseExpression()

	if value == nil {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected expression after return", true, debug.NewHint("Did you forget to add an expression?", " <expression>"))
	}

	return &ast.Return{
		Value: value,
		Loc:   value.Location(),
	}
}
