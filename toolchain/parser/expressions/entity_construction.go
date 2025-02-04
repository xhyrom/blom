package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseEntityConstruction(p Parser, identifier tokens.Token) ast.Statement {
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

func parseEntityConstructionValue(p Parser) (string, ast.Expression) {
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
