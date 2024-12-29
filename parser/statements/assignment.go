package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses an assignment statement that can have form:
// <identifier> = <expression>;
func ParseAssignment(p Parser) *ast.AssignmentStatement {
	name := p.Consume()

	eq := p.Consume()

	value, _ := p.ParseExpression()

	if p.Consume().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
	}

	return &ast.AssignmentStatement{
		Name:  name.Value,
		Value: value,
		Loc:   eq.Location,
	}
}
