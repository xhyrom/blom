package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a type, for example:
// - i32
// - i64
func (p *Parser) parseType() (ast.Type, tokens.Location) {
	if p.Current().Kind != tokens.Identifier {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add a type name?", " i32"))
	}

	token := p.Consume()

	str := token.Value

	if p.Current().Kind == tokens.Asterisk {
		str += "*"
		token = p.Consume()
	}

	ty, err := ast.ParseType(str)
	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column)
		dbg.ThrowError(err.Error(), true)
	}

	return ty, token.Location
}
