package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses an assignment statement that can have form:
// <expression> = <expression>;
func ParseAssignment(p Parser, left ast.Expression) *ast.Assignment {
	if left == nil {
		left, _ = p.ParseExpression()
	}

	eq := p.Consume()

	right, _ := p.ParseExpression()

	if p.Consume().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), right)
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
	}

	return &ast.Assignment{
		Left:  left,
		Right: right,
		Loc:   eq.Location,
	}
}
