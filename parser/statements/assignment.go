package statements

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"blom/tokens"
)

// Parses an assignment statement that can have form:
// <type> <identifier> = <expression>;
// <type> <identifier>;
// <identifier> = <expression>;
func ParseAssignment(p Parser, redeclaration bool) *ast.DeclarationStatement {
	if redeclaration {
		name := p.Consume()

		eq := p.Consume()

		value, _ := p.ParseExpression()

		if p.Consume().Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		}

		return &ast.DeclarationStatement{
			Name:          name.Value,
			Value:         value,
			Redeclaration: true,
			Type:          nil,
			Loc:           eq.Location,
		}
	}

	valueTypeToken := p.Consume()

	valueType, err := compiler.ParseType(valueTypeToken.Value)
	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), valueTypeToken.Location.Row, valueTypeToken.Location.Column)
		dbg.ThrowError(err.Error(), true)
	}

	name := p.Consume()
	var value ast.Expression = nil

	right := p.Consume()

	if right.Kind != tokens.Assign && right.Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+1)
		dbg.ThrowError("Expected assignment or semicolon", true, debug.NewHint("Add semicolon", ";"), debug.NewHint("Initialize variable", " = 0;"))
	}

	if right.Kind == tokens.Assign {
		exp, _ := p.ParseExpression()
		value = exp
	}

	if p.Consume().Kind != tokens.Semicolon {
		if value != nil {
			dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		} else {
			dbg := debug.NewSourceLocation(p.Source(), right.Location.Row, right.Location.Column+1)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		}
	}

	return &ast.DeclarationStatement{
		Name:          name.Value,
		Value:         value,
		Redeclaration: false,
		Type:          &valueType,
		Loc:           right.Location,
	}
}
