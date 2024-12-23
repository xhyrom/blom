package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseFunction(p Parser) ast.Statement {
	p.Consume()

	annotations := make([]ast.Annotation, 0)
	for p.Current().Kind == tokens.AtMark {
		annotations = append(annotations, parseAnnotation(p))
	}

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

	if len(arguments) == 0 {
		p.Consume()
	}

	if p.Current().Kind == tokens.Minus {
		p.Consume()

		current = p.Consume()
		if current.Kind != tokens.GreaterThan {
			panic("Expected '->'")
		}

		returnType = p.Consume().Kind
		current = p.Current()
	}

	if p.Current().Kind != tokens.LeftCurlyBracket {
		panic("Expected '{'")
	}

	fn := ast.FunctionDeclaration{
		Name:        name.Value,
		Arguments:   arguments,
		Annotations: annotations,
		ReturnType:  int(returnType),
		Body:        ParseBlock(p),
		Loc:         name.Location,
	}

	return fn
}

func parseAnnotation(p Parser) ast.Annotation {
	p.Consume()

	name := p.Consume()

	return ast.Annotation{
		Name: name.Value,
		Loc:  name.Location,
	}
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
