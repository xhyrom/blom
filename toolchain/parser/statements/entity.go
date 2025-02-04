package statements

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
func ParseEntity(p Parser) *ast.Entity {
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
		fields = append(fields, ParseVariableDeclaration(p))
	}

	if p.Consume().Kind != tokens.RightCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected '}'", true)
	}

	entity := &ast.Entity{
		Name:   name.Value,
		Fields: fields,
		Loc:    name.Location,
	}

	p.AddCustomType(name.Value, entity)

	return entity
}
