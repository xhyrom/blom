package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseReturn(p Parser) ast.Statement {
	p.Consume()

	current := p.Current()
	if current.Kind == tokens.Semicolon {
		p.Consume()

		return ast.ReturnStatement{
			Value: nil,
			Loc:   current.Location,
		}
	}

	value := p.ParseExpression()

	if p.Consume().Kind != tokens.Semicolon {
		panic("Expected semicolon")
	}

	return ast.ReturnStatement{
		Value: value,
		Loc:   value.Location(),
	}
}
