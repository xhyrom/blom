package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parseBlock() *ast.Block {
	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected opening bracket", true, debug.NewHint("Did you forget to add an opening bracket?", " {"))
	}

	p.Consume() // consume the opening bracket

	current := p.Current()

	body := []ast.Node{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		stmt := p.parseStatement()

		body = append(body, stmt)

		current = p.Current()
	}

	if p.Current().Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected closing bracket", true, debug.NewHint("Did you forget to add a closing bracket?", "}"))
	}

	p.Consume() // consume the closing bracket

	return &ast.Block{
		Body: body,
		Loc:  current.Location,
	}
}
