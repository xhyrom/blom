package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

// Parses a function statement that can have form:
// fun <identifier>() { <body> }
// fun <identifier>(<identifier>: <type>) { <body> }
// fun <identifier>() -> <type> { <body> }
// fun <identifier>(<identifier>: <type>) -> <type> { <body> }
// there can be any number of arguments inside ()
func (p *Parser) parseFunction() *ast.FunctionDeclaration {
	p.Consume()

	p.collectAnnotations()

	name := p.Consume()
	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Function name must be valid identifier, got \"%s\"", name.Value), true)
	}

	if p.Current().Kind == tokens.DoubleColon {
		_, err := ast.ParseType(name.Value, p.CustomTypes())
		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Cannot extend type \"%s\" because it isn't a primitive type",
					name.Value,
				),
				true,
			)
		}

		p.Consume()
		identifier := p.Consume()
		if identifier.Kind != tokens.Identifier {
			dbg := debug.NewSourceLocation(p.Source(), identifier.Location.Row, identifier.Location.Column)
			dbg.ThrowError(fmt.Sprintf("Function name must be valid identifier, got \"%s\"", name.Value), true)
		}

		name = tokens.Token{
			Kind:     tokens.Identifier,
			Location: name.Location,
			Value:    name.Value + "." + identifier.Value,
		}
	}

	current := p.Consume()

	if current.Kind != tokens.LeftParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+1)
		dbg.ThrowError("Function arguments must be enclosed in parentheses", true, debug.NewHint("Did you forget to add parentheses?", "()"))
	}

	arguments := make([]ast.FunctionArgument, 0)
	locationBeforeBlock := p.Current().Location

	fn := ast.FunctionDeclaration{
		Name:        name.Value,
		Annotations: p.extractAnnotations(),
		Loc:         name.Location,
	}

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		arg, location := parseFunctionArgument(p, &fn)
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

	fn.Arguments = arguments

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
		returnType, err = ast.ParseType(returnTypeToken.Value, p.CustomTypes())

		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), returnTypeToken.Location.Row, returnTypeToken.Location.Column)
			dbg.ThrowError(err.Error(), true)
		}

		locationBeforeBlock = returnTypeToken.Location
		current = p.Current()
	}

	fn.ReturnType = returnType

	if fn.IsNative() {
		if p.Current().Kind == tokens.LeftCurlyBracket {
			dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
			dbg.ThrowError("Native function must not have a body", true, debug.NewHint("Remove '{' to make function native", ""))
		}

		if current.Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocation(p.Source(), locationBeforeBlock.Row, locationBeforeBlock.Column+1)
			dbg.ThrowError("Did you forget to add a semicolon?", true, debug.NewHint("Add ';'", ";"))
		}

		p.Consume()

		return &fn
	}

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), locationBeforeBlock.Row, locationBeforeBlock.Column+1)
		dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
	}

	block := p.parseBlock()

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

	fn.Body = block.Body
	return &fn
}

// Parses an function argument that can have form:
// <identifier>: <type>
func parseFunctionArgument(p *Parser, fun *ast.FunctionDeclaration) (*ast.FunctionArgument, *tokens.Location) {
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

	typ, err := ast.ParseType(typStr, p.CustomTypes())

	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), typToken.Location.Row, typToken.Location.Column)
		dbg.ThrowError(err.Error(), true)
	}

	return &ast.FunctionArgument{
		Name: name.Value,
		Type: typ,
	}, &typToken.Location
}

// Parses a function call that can have form:
// <identifier>(<expression>, <expression>, ...)
func (p *Parser) parseFunctionCall(identifier tokens.Token, requiresSemicolon bool) *ast.FunctionCall {
	p.Consume()

	name := identifier.Value
	parameters := make([]ast.Expression, 0)

	for p.Current().Kind != tokens.RightParenthesis {
		exp, err := p.ParseExpression()
		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), identifier.Location.Row, identifier.Location.Column+2)
			dbg.ThrowError(
				err.Error(),
				true,
				debug.NewHint("Did you forget to close a function call?", ")"),
			)
		}

		parameters = append(parameters, exp)

		if p.Current().Kind != tokens.Comma {
			if p.Current().Kind != tokens.RightParenthesis {
				dbg := debug.NewSourceLocationFromExpression(p.Source(), parameters[len(parameters)-1])
				dbg.ThrowError(
					"Expected comma or right parenthesis",
					true,
					debug.NewHint("Add comma for more parameters", ","),
					debug.NewHint("Add closing parenthesis to end function call", ")"),
				)
			}

			break
		}

		p.Consume()
	}

	last := p.Consume()

	if requiresSemicolon {
		if p.Consume().Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocation(p.Source(), last.Location.Row, last.Location.Column+1)
			dbg.ThrowError(
				"Expected semicolon",
				true,
				debug.NewHint("Did you forget to add a semicolon?", ";"),
			)
		} else {
			last = p.Current()
		}
	}

	return &ast.FunctionCall{
		Name:       name,
		Parameters: parameters,
		Loc:        last.Location,
	}
}
