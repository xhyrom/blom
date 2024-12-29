package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseIdentifier(p Parser) ast.Statement {
	token := p.Consume()

	if p.Current().Kind == tokens.LeftParenthesis {
		return ParseFunctionCall(p, token, false)
	}

	return &ast.IdentifierLiteralStatement{
		Value: token.Value,
		Loc:   token.Location,
	}
}
