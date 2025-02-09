package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a variable declaration statement that can have form:
// <type> <identifier> = <expression>;
// <type> <identifier>;
func (p *Parser) parseVariableDeclaration() *ast.VariableDeclarationStatement {
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

	valueType, err := ast.ParseType(typeStr, p.CustomTypes())
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

	if right.Kind != tokens.Semicolon && p.Consume().Kind != tokens.Semicolon {
		if value != nil {
			dbg := debug.NewSourceLocationFromExpression(p.Source(), value)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		} else {
			dbg := debug.NewSourceLocation(p.Source(), right.Location.Row, right.Location.Column+1)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		}
	}

	return &ast.VariableDeclarationStatement{
		Name:        name.Value,
		Value:       value,
		Type:        valueType,
		Annotations: p.extractAnnotations(),
		Loc:         right.Location,
	}
}

// Parses an assignment statement that can have form:
// <expression> = <expression>;
func (p *Parser) parseAssignment(left ast.Expression) *ast.Assignment {
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
