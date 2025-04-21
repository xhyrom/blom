package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a argument that can have a form:
// <argument> : <type>
//
// where:
// - <argument> is the name of the argument
// - <type> is the type of the argument
func (p *Parser) parseArgument() *ast.Argument {
	if p.Current().Kind != tokens.Identifier {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add an argument name?", "<name>"))
	}

	argument := p.parseLiteral().(*ast.IdentifierLiteral)

	if p.Consume().Kind != tokens.Colon {
		dbg := debug.NewSourceLocationFromNode(p.Source(), argument)
		dbg.ThrowError("Expected colon and type", true, debug.NewHint("Did you forget to add a colon and type?", ": i32"))
	}

	ty, loc := p.parseType()

	return &ast.Argument{
		Name: *argument,
		Type: ty,
		Loc:  loc,
	}
}
