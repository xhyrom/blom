package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseReturn(p Parser) (*ast.ReturnStatement, error) {
	p.Consume()

	current := p.Current()
	if current.Kind == tokens.Semicolon {
		p.Consume()

		return &ast.ReturnStatement{
			Value: nil,
			Loc:   current.Location,
		}, nil
	}

	value, _ := p.ParseExpression()

	if p.Consume().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Add semicolon", ";"))
	}

	return &ast.ReturnStatement{
		Value: value,
		Loc:   value.Location(),
	}, nil
}
