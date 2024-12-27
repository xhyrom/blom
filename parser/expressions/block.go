package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseBlock(p Parser, consumeFirst bool) (*ast.BlockStatement, error) {
	if consumeFirst {
		p.Consume()
	}

	current := p.Current()

	body := []ast.Statement{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		stmt, _ := p.ParseStatement()

		body = append(body, stmt)
		current = p.Current()
	}

	if p.Consume().Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), current.Location.Row, current.Location.Column)
		dbg.ThrowError("Expected closing bracket", true, debug.NewHint("Add closing bracket", "}"))
	}

	return &ast.BlockStatement{
		Body: body,
		Loc:  current.Location,
	}, nil
}
