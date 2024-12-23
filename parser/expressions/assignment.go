package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseAssignment(p Parser, var_type tokens.Token) ast.Statement {
	name := p.Consume().Value

	if p.Consume().Kind != tokens.Assign {
		panic("Expected assignment operator")
	}

	value := p.ParseExpression()

	if p.Consume().Kind != tokens.Semicolon {
		panic("Expected semicolon")
	}

	return ast.DeclarationStatement{
		Name:  name,
		Value: value,
		Type:  int(var_type.Kind),
	}
}
