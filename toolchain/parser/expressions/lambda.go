package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

// Parses a lambda expression that can have form:
// fun () { <body> }
// fun (<identifier>: <type>) { <body> }
// there can be any number of arguments inside ()
func ParseLambda(p Parser) *ast.LambdaDeclaration {
	fun := p.Consume()
	current := p.Consume()

	if current.Kind != tokens.LeftParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), fun.Location.Row, fun.Location.Column+1)
		dbg.ThrowError("Function arguments must be enclosed in parentheses", true, debug.NewHint("Did you forget to add parentheses?", "()"))
	}

	arguments := make([]ast.FunctionArgument, 0)
	locationBeforeBlock := p.Current().Location

	lambda := ast.LambdaDeclaration{
		Loc: fun.Location,
	}

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		arg, location := parseArgument(p, &lambda)
		if arg == nil {
			p.Consume()
			break
		}

		arguments = append(arguments, *arg)

		current = p.Consume()

		if current.Kind != tokens.Comma && current.Kind != tokens.RightParenthesis {
			dbg := debug.NewSourceLocation(p.Source(), location.Row, location.Column+1)
			dbg.ThrowError(
				"Expected comma or right parenthesis",
				true,
				debug.NewHint("Arguments must be separated by commas", ","),
				debug.NewHint("Did you forget to close the parentheses?", ")"),
			)
		}

		locationBeforeBlock = tokens.Location{
			Row:    location.Row,
			Column: location.Column + 1,
		}
	}

	lambda.Arguments = arguments

	var returnType ast.Type = ast.Int32

	if len(arguments) == 0 {
		p.Consume()
	}

	minusLocation := p.Current().Location
	if p.Current().Kind == tokens.Minus {
		p.Consume()

		current = p.Consume()
		if current.Kind != tokens.GreaterThan {
			dbg := debug.NewSourceLocation(p.Source(), minusLocation.Row, minusLocation.Column+1)
			dbg.ThrowError("Return type must be preceded by a dash", true, debug.NewHint("Did you forget to add a dash?", ">"))
		}

		returnTypeToken := p.Consume()
		if returnTypeToken.Kind != tokens.Identifier {
			dbg := debug.NewSourceLocation(p.Source(), returnTypeToken.Location.Row, returnTypeToken.Location.Column)
			dbg.ThrowError(fmt.Sprintf("Return type must be a valid type, got \"%s\"", returnTypeToken.Value), true)
		}

		var err error
		returnType, err = ast.ParseType(returnTypeToken.Value)

		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), returnTypeToken.Location.Row, returnTypeToken.Location.Column)
			dbg.ThrowError(err.Error(), true)
		}

		locationBeforeBlock = returnTypeToken.Location
		current = p.Current()
	}

	lambda.ReturnType = returnType

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), locationBeforeBlock.Row, locationBeforeBlock.Column+1)
		dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
	}

	block := ParseBlock(p)
	// TOOD: move to analyzer
	hasReturn := false
	for _, stmt := range block.Body {
		if stmt.Kind() == ast.ReturnNode {
			hasReturn = true
			break
		}
	}

	if !hasReturn {
		block.Body = append(block.Body, &ast.ReturnStatement{
			Loc: block.Loc,
			Value: &ast.IntLiteral{
				Value: 0,
			},
		})
	}

	lambda.Body = block.Body

	return &lambda
}

func parseArgument(p Parser, fun *ast.LambdaDeclaration) (*ast.FunctionArgument, *tokens.Location) {
	name := p.Consume()
	if name.Kind == tokens.Ellipsis {
		fun.Variadic = true
		return nil, nil
	}

	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Argument name must be valid identifier, got \"%s\"", name.Value), true)
	}

	if p.Consume().Kind != tokens.Colon {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+1)
		dbg.ThrowError("Argument type must be preceded by a colon", true, debug.NewHint("Did you forget to add a colon?", ":"))
	}

	typToken := p.Consume()
	typStr := typToken.Value

	if typToken.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), typToken.Location.Row, typToken.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Argument type must be a valid type, got \"%s\"", typToken.Value), true)
	}

	if p.Current().Kind == tokens.Asterisk {
		typStr = typToken.Value + "*"
		typToken = p.Consume()
	}

	typ, err := ast.ParseType(typStr)

	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), typToken.Location.Row, typToken.Location.Column)
		dbg.ThrowError(err.Error(), true)
	}

	return &ast.FunctionArgument{
		Name: name.Value,
		Type: typ,
	}, &typToken.Location
}
