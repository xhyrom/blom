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
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column+1)
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add an argument name?", "<name>"))
	}

	argument := p.parseLiteral().(*ast.IdentifierLiteral)

	if p.Consume().Kind != tokens.Colon {
		dbg := debug.NewSourceLocation(p.Source(), argument.Location().Row, argument.Location().Column+1)
		dbg.ThrowError("Expected colon and type", true, debug.NewHint("Did you forget to add a colon and type?", ": i32"))
	}

	return &ast.Argument{
		Name: *argument,
		Type: p.parseType(),
	}
}
