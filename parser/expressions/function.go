package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseFunction(p Parser) ast.Statement {
	p.Consume()

	name := p.Consume()
	current := p.Consume()

	if current.Kind != tokens.LeftParenthesis {
		panic("Expected '('")
	}

	arguments := make([]ast.FunctionArgument, 0)

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		arg := parseArgument(p)
		arguments = append(arguments, arg)

		current = p.Consume()

		if current.Kind != tokens.Comma && current.Kind != tokens.RightParenthesis {
			panic("Expected ',' or ')'")
		}
	}

	var returnType tokens.TokenKind

	current = p.Consume()
	if current.Kind == tokens.RightParenthesis { // If there are no arguments, otherwise already consumed
		current = p.Consume()
	}

	if current.Kind == tokens.Minus {
		current = p.Consume()
		if current.Kind != tokens.GreaterThan {
			panic("Expected '->'")
		}

		returnType = p.Consume().Kind
		current = p.Consume()
	}

	if current.Kind != tokens.LeftCurlyBracket {
		panic("Expected '{'")
	}

	fn := ast.FunctionDeclaration{
		Name:       name.Value,
		Arguments:  arguments,
		ReturnType: int(returnType),
		Body:       make([]ast.Statement, 0),
		Loc:        name.Location,
	}

	current = p.Current()
	for current.Kind != tokens.RightCurlyBracket {
		fn.Body = append(fn.Body, p.ParseStatement())
		current = p.Current()
	}

	p.Consume()

	return fn
}

func parseArgument(p Parser) ast.FunctionArgument {
	name := p.Consume()
	if p.Consume().Kind != tokens.Colon {
		panic("Expected ':'")
	}

	typ := p.Consume()

	return ast.FunctionArgument{
		Name: name.Value,
		Type: int(typ.Kind),
	}
}
