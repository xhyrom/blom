package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parseBlock() *ast.Block {
	p.Consume()

	current := p.Current()

	body := []ast.Node{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		stmt := p.parseStatement()

		body = append(body, stmt)

		current = p.Current()
	}

	if p.Consume().Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), current.Location.Row, current.Location.Column)
		dbg.ThrowError("Expected closing bracket", true, debug.NewHint("Did you forget to add a closing bracket?", "}"))
	}

	return &ast.Block{
		Body: body,
		Loc:  current.Location,
	}
}
