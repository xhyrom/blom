package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseReturn(p Parser) *ast.ReturnStatement {
	returnLoc := p.Consume().Location

	current := p.Current()
	if current.Kind == tokens.Semicolon {
		p.Consume()

		return &ast.ReturnStatement{
			Value: nil,
			Loc:   current.Location,
		}
	}

	value, err := p.ParseExpression()
	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), returnLoc.Row, returnLoc.Column+1)
		dbg.ThrowError("Expected expression", true, debug.NewHint("Did you forget to return a value?", " 0;"))
	}

	if p.Consume().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
	}

	return &ast.ReturnStatement{
		Value: value,
		Loc:   value.Location(),
	}
}
