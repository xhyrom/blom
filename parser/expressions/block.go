package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseBlock(p Parser) *ast.BlockStatement {
	p.Consume()

	current := p.Current()

	body := []ast.Statement{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		stmt, err := p.ParseStatement()
		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), current.Location.Row, current.Location.Column)
			dbg.ThrowError(err.Error(), true)
		}

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
