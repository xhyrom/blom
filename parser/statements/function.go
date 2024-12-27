package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/parser/expressions"
	"blom/tokens"
	"fmt"
)

// Parses a function statement that can have form:
// fun <identifier>() { <body> }
// fun <identifier>(<identifier>: <type>) { <body> }
// fun <identifier>() -> <type> { <body> }
// fun <identifier>(<identifier>: <type>) -> <type> { <body> }
// there can be any number of arguments inside ()
func ParseFunction(p Parser) *ast.FunctionDeclaration {
	p.Consume()

	annotations := make([]ast.Annotation, 0)
	for p.Current().Kind == tokens.AtMark {
		annotations = append(annotations, parseAnnotation(p))
	}

	name := p.Consume()
	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Function name must be valid identifier, got \"%s\"", name.Value), true)
	}

	current := p.Consume()

	if current.Kind != tokens.LeftParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+1)
		dbg.ThrowError("Function arguments must be enclosed in parentheses", true, debug.NewHint("Did you forget to add parentheses?", "()"))
	}

	arguments := make([]ast.FunctionArgument, 0)
	locationBeforeBlock := p.Current().Location

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		arg, location := parseArgument(p)
		arguments = append(arguments, arg)

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

	var returnType tokens.Token

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

		returnType = p.Consume()
		if returnType.Kind != tokens.Identifier {
			dbg := debug.NewSourceLocation(p.Source(), returnType.Location.Row, returnType.Location.Column)
			dbg.ThrowError(fmt.Sprintf("Return type must be a valid type, got \"%s\"", returnType.Value), true)
		}

		locationBeforeBlock = returnType.Location
		current = p.Current()
	}

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), locationBeforeBlock.Row, locationBeforeBlock.Column+1)
		dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
	}

	block := expressions.ParseBlock(p)

	fn := &ast.FunctionDeclaration{
		Name:        name.Value,
		Arguments:   arguments,
		Annotations: annotations,
		ReturnType:  int(returnType.Kind),
		Body:        block,
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

func parseArgument(p Parser) (ast.FunctionArgument, tokens.Location) {
	name := p.Consume()
	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Argument name must be valid identifier, got \"%s\"", name.Value), true)
	}

	if p.Consume().Kind != tokens.Colon {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+1)
		dbg.ThrowError("Argument type must be preceded by a colon", true, debug.NewHint("Did you forget to add a colon?", ":"))
	}

	typ := p.Consume()

	if typ.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), typ.Location.Row, typ.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Argument type must be a valid type, got \"%s\"", typ.Value), true)
	}

	return ast.FunctionArgument{
		Name: name.Value,
		Type: int(typ.Kind),
	}, typ.Location
}
