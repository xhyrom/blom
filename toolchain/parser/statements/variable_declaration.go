package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a variable declaration statement that can have form:
// <type> <identifier> = <expression>;
// <type> <identifier>;
func ParseVariableDeclaration(p Parser) *ast.VariableDeclarationStatement {
	valueTypeToken := p.Consume()

	var typeStr string
	if valueTypeToken.Kind == tokens.Identifier {
		typeStr = valueTypeToken.Value
		for p.Current().Kind == tokens.Asterisk {
			typeStr += "*"
			p.Consume()
		}
	} else {
		dbg := debug.NewSourceLocation(p.Source(), valueTypeToken.Location.Row, valueTypeToken.Location.Column)
		dbg.ThrowError("Expected type identifier", true)
	}

	valueType, err := ast.ParseType(typeStr)
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

	return &ast.VariableDeclarationStatement{
		Name:  name.Value,
		Value: value,
		Type:  valueType,
		Loc:   right.Location,
	}
}
