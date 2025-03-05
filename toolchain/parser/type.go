package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a type, for example:
// - i32
// - i64
func (p *Parser) parseType() ast.Type {
	token := p.Consume()

	if token.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column)
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add a type name?", "i32"))
	}

	ty, err := ast.ParseType(token.Value, map[string]ast.Type{})
	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column)
		dbg.ThrowError(err.Error(), true)
	}

	return ty
}
