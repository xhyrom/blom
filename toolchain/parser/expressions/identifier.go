package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

func ParseIdentifier(p Parser) ast.Statement {
	token := p.Consume()

	if p.Current().Kind == tokens.LeftParenthesis {
		return ParseFunctionCall(p, token, true)
	}

	if p.Current().Kind == tokens.DoubleColon {
		_, err := ast.ParseType(token.Value, p.CustomTypes())
		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Cannot extend type \"%s\" because it isn't a primitive type",
					token.Value,
				),
				true,
			)
		}

		p.Consume()
		identifier := p.Consume()
		if identifier.Kind != tokens.Identifier {
			dbg := debug.NewSourceLocation(p.Source(), identifier.Location.Row, identifier.Location.Column)
			dbg.ThrowError(fmt.Sprintf("Function name must be valid identifier, got \"%s\"", token.Value), true)
		}

		token = tokens.Token{
			Kind:     tokens.Identifier,
			Location: token.Location,
			Value:    token.Value + "." + identifier.Value,
		}

		return ParseFunctionCall(p, token, true)
	}

	if p.Current().Kind == tokens.LeftCurlyBracket {
		return ParseEntityConstruction(p, token)
	}

	return &ast.IdentifierLiteral{
		Value: token.Value,
		Loc:   token.Location,
	}
}
