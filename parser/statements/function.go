package statements

import (
	"blom/ast"
	"blom/compiler"
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

	fn := ast.FunctionDeclaration{
		Name:        name.Value,
		Annotations: annotations,
		Loc:         name.Location,
	}

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		arg, location := parseArgument(p, &fn)
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

	var returnType compiler.Type = compiler.Void

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
		returnType, err = compiler.ParseType(returnTypeToken.Value)

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

	block := expressions.ParseBlock(p)

	fn.Body = block
	return &fn
}

func parseAnnotation(p Parser) ast.Annotation {
	p.Consume()

	name := p.Consume()
	typ := ast.ParseAnnotation(name.Value)

	if typ == -1 {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Annotation \"%s\" is not recognized", name.Value), true)
	}

	return ast.Annotation{
		Type: typ,
		Loc:  name.Location,
	}
}

func parseArgument(p Parser, fun *ast.FunctionDeclaration) (*ast.FunctionArgument, *tokens.Location) {
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

	if typToken.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), typToken.Location.Row, typToken.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Argument type must be a valid type, got \"%s\"", typToken.Value), true)
	}

	typ, err := compiler.ParseType(typToken.Value)

	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), typToken.Location.Row, typToken.Location.Column)
		dbg.ThrowError(err.Error(), true)
	}

	return &ast.FunctionArgument{
		Name: name.Value,
		Type: typ,
	}, &typToken.Location
}
