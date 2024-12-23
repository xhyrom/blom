package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseAssignment(p Parser, token tokens.Token, redeclaration bool) ast.Statement {
	if redeclaration {
		p.Consume()

		value := p.ParseExpression()

		if p.Consume().Kind != tokens.Semicolon {
			panic("Expected semicolon")
		}

		return ast.DeclarationStatement{
			Name:          token.Value,
			Value:         value,
			Redeclaration: true,
		}
	}

	name := p.Consume().Value

	if p.Consume().Kind != tokens.Assign {
		panic("Expected assignment operator")
	}

	value := p.ParseExpression()

	if p.Consume().Kind != tokens.Semicolon {
		panic("Expected semicolon")
	}

	return ast.DeclarationStatement{
		Name:          name,
		Value:         value,
		Redeclaration: false,
		Type:          int(token.Kind),
	}
}
