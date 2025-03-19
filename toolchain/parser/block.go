package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	p.Consume()

	current := p.Current()

	body := []ast.Statement{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		stmt := p.parseStatement()

		body = append(body, stmt)

		current = p.Current()
	}

	if p.Consume().Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), current.Location.Row, current.Location.Column)
		dbg.ThrowError("Expected closing bracket", true, debug.NewHint("Did you forget to add a closing bracket?", "}"))
	}

	return &ast.BlockStatement{
		Body: body,
		Loc:  current.Location,
	}
}

func (p *Parser) parseBlockExpression() *ast.BlockExpression {
	p.Consume()

	current := p.Current()

	body := []ast.Statement{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		expr := p.parseStatement()

		body = append(body, expr)

		current = p.Current()
	}

	if p.Consume().Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), current.Location.Row, current.Location.Column)
		dbg.ThrowError("Expected closing bracket", true, debug.NewHint("Did you forget to add a closing bracket?", "}"))
	}

	return &ast.BlockExpression{
		Body: body,
		Loc:  current.Location,
	}
}
