package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

// Parses a type definition statement that can have form:
// type <identifier> = <primitive type>
// type <identifier> = <function signature>
func ParseTypeDefinition(p Parser) *ast.TypeDefinition {
	p.Consume()

	name := p.Consume()
	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Type name must be valid identifier, got \"%s\"", name.Value), true)
	}

	if p.Current().Kind != tokens.Assign {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError(fmt.Sprintf("Expected '=', got %s", p.Current().Kind), true)
	}

	p.Consume()

	var ty ast.Type
	if p.Current().Kind == tokens.Fun {
		ty = parseFunctionSignature(p)
	} else {
		t, err := ast.ParseType(p.Current().Value, p.CustomTypes())

		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Cannot assign type \"%s\" to type \"%s\" because it isn't a primitive type",
					p.Current().Value,
					name.Value,
				),
				true,
			)
		}

		ty = t
		p.Consume()
	}

	if p.Current().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError(fmt.Sprintf("Expected ';', got %s", p.Current().Kind), true)
	}

	p.Consume()

	p.AddCustomType(name.Value, ty)

	return &ast.TypeDefinition{
		Name: name.Value,
		Type: ty,
		Loc:  name.Location,
	}
}

func parseFunctionSignature(p Parser) ast.Type {
	p.Consume()

	current := p.Consume()

	if current.Kind != tokens.LeftParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError("Function arguments must be enclosed in parentheses", true, debug.NewHint("Did you forget to add parentheses?", "()"))
	}

	arguments := make([]ast.Type, 0)

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		token := p.Consume()
		ty, err := ast.ParseType(token.Value, p.CustomTypes())

		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
			dbg.ThrowError(
				fmt.Sprintf(
					"Cannot assign type \"%s\" to function argument because it isn't a primitive type",
					p.Current().Value,
				),
				true,
			)
		}

		arguments = append(arguments, ty)
		current = p.Consume()

		if current.Kind != tokens.Comma && current.Kind != tokens.RightParenthesis {
			dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column+1)
			dbg.ThrowError(
				"Expected comma or right parenthesis",
				true,
				debug.NewHint("Arguments must be separated by commas", ","),
				debug.NewHint("Did you forget to close the parentheses?", ")"),
			)
		}
	}

	var returnType ast.Type

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

		current = p.Current()
	}

	return ast.NewFunctionType(
		arguments,
		returnType,
	)
}
