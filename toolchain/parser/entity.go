package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses an entity that can have form:
//
//	entity <identifier> {
//	   <variable_declaration>
//	   <function_declaration>
//	}
func (p *Parser) parseEntity() *ast.Entity {
	p.Consume()

	name := p.Consume()
	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected identifier", true)
	}

	if p.Consume().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected '{'", true)
	}

	fields := make([]*ast.VariableDeclarationStatement, 0)

	for !p.IsEof() && p.Current().Kind == tokens.Identifier {
		fields = append(fields, p.parseVariableDeclaration())
	}

	token := p.Consume()
	if token.Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected '}'", true)
	}

	if p.Consume().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column+1)
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
	}

	entity := &ast.Entity{
		Name:   name.Value,
		Fields: fields,
		Loc:    name.Location,
	}

	p.AddCustomType(name.Value, entity)

	return entity
}

// Parses an entity construction that can have form:
//
//	<identifier> {
//	 <identifier> = <expression>,
//	 <identifier> = <expression>,
//	 ...
//	}
func (p *Parser) parseEntityConstruction(identifier tokens.Token) ast.Statement {
	p.Consume()

	name := identifier.Value

	values := make(map[string]ast.Expression)

	for !p.IsEof() && p.Current().Kind != tokens.RightCurlyBracket {
		key, value := parseEntityConstructionValue(p)

		values[key] = value

		if p.Current().Kind == tokens.Comma {
			p.Consume()
		}
	}

	p.Consume()

	return &ast.EntityConstruction{
		Name:   name,
		Values: values,
		Loc:    identifier.Location,
	}
}

func parseEntityConstructionValue(p *Parser) (string, ast.Expression) {
	name := p.Consume()

	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected identifier", true)
	}

	if p.Current().Kind != tokens.Assign {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected assignment", true)
	}

	p.Consume()

	value, _ := p.ParseExpression()

	return name.Value, value
}
