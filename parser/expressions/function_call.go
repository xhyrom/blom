package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseFunctionCall(p Parser, identifier tokens.Token) *ast.FunctionCall {
	p.Consume()

	name := identifier.Value
	parameters := make([]ast.Expression, 0)

	for p.Current().Kind != tokens.RightParenthesis {
		parameters = append(parameters, p.ParseExpression())

		if p.Current().Kind != tokens.Comma {
			break
		}

		p.Consume()
	}

	p.Consume()

	return &ast.FunctionCall{
		Name:       name,
		Parameters: parameters,
	}
}
