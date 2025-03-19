package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a function declaration that can have a form:
// fun ?(<annotations>) <name> (<parameters>) ?(-> <return type>) { <body> }
//
// where:
// - <annotations> is a list of annotations
// - <name> is the name of the function
// - <parameters> is a list of parameters separated by commas
// - <return type> is the type of the return value
// - <body> is the body of the function
//
// Example:
// fun add(a: i32, b: i32) -> i32 {}
// fun main() {}
func (p *Parser) parseFunction() *ast.FunctionDeclaration {
	p.Consume()
	p.collectAnnotations()

	name := p.Consume()
	if name.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add a function name?", "fn"))
	}

	if p.Consume().Kind != tokens.LeftParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError("Expected opening parenthesis", true, debug.NewHint("Did you forget to add an opening parenthesis?", "("))
	}

	fun := &ast.FunctionDeclaration{
		Name: name.Value,
		Loc:  name.Location,
	}

	arguments := make([]ast.FunctionArgument, 0)
	for p.Current().Kind != tokens.RightParenthesis {
		argument := parseFunctionArgument(p, fun)
		if argument == nil {
			p.Consume()
			break
		}

		arguments = append(arguments, *argument)

		if p.Current().Kind == tokens.Comma {
			p.Consume()
		}
	}

	if p.Consume().Kind != tokens.RightParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError("Expected closing parenthesis", true, debug.NewHint("Did you forget to add a closing parenthesis?", ")"))
	}

	if p.Current().Kind == tokens.Minus {
		p.Consume()

		if p.Consume().Kind != tokens.GreaterThan {
			dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column+3)
			dbg.ThrowError("Expected arrow", true, debug.NewHint("Did you forget to add an arrow?", " ->"))
		}

		fun.ReturnType = p.parseType()
	} else {
		fun.ReturnType = ast.Int32
	}

	fun.Arguments = arguments
	fun.Annotations = p.extractAnnotations()
	fun.Body = p.parseBlock(ast.StatementCategory).Body

	return fun
}

// Parses a list of function arguments that can have a form:
// (<argument> : <type>, <argument> : <type>, ...)
//
// where:
// - <argument> is the name of the argument
// - <type> is the type of the argument
func parseFunctionArgument(p *Parser, fun *ast.FunctionDeclaration) *ast.FunctionArgument {
	argument := p.Consume()
	if argument.Kind == tokens.Ellipsis {
		fun.Variadic = true
		return nil
	}

	if argument.Kind != tokens.Identifier {
		dbg := debug.NewSourceLocation(p.Source(), argument.Location.Row, argument.Location.Column)
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add an argument name?", "<name>"))
	}

	if p.Consume().Kind != tokens.Colon {
		dbg := debug.NewSourceLocation(p.Source(), argument.Location.Row, argument.Location.Column)
		dbg.ThrowError("Expected colon", true, debug.NewHint("Did you forget to add a colon?", ":"))
	}

	return &ast.FunctionArgument{
		Name: argument.Value,
		Type: p.parseType(),
	}
}
