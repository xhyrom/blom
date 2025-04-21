package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a function declaration that can have a form:
// fun ?(<annotations>) <path> (<parameters>) ?(-> <return type>) { <body> }
//
// where:
// - <annotations> is a list of annotations
// - <path> is the path (name) of the function
// - <parameters> is a list of parameters separated by commas
// - <return type> is the type of the return value
// - <body> is the body of the function
//
// Example:
// fun add(a: i32, b: i32) -> i32 {}
// fun a::b::c::d::e::f(a: i32, b: i32) -> i32 {}
// fun main() {}
func (p *Parser) parseFunction() *ast.FunctionDeclaration {
	p.Consume()
	p.collectAnnotations()

	if p.Current().Kind != tokens.Identifier {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected identifier", true, debug.NewHint("Did you forget to add a function name?", " <name>"))
	}

	path := p.parsePath(p.parseLiteral())

	if p.Consume().Kind != tokens.LeftParenthesis {
		dbg := debug.NewSourceLocationFromNode(p.Source(), path)
		dbg.ThrowError("Expected opening parenthesis", true, debug.NewHint("Did you forget to add an opening parenthesis?", "("))
	}

	fun := &ast.FunctionDeclaration{
		Path: *path,
		Loc:  path.Location(),
	}

	params := make([]ast.Argument, 0)
	for p.Current().Kind != tokens.RightParenthesis {
		param := p.parseArgument()
		if param == nil {
			break
		}

		params = append(params, *param)

		if p.Current().Kind == tokens.Comma {
			p.Consume()
		} else if p.Current().Kind != tokens.RightParenthesis {
			dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
			dbg.ThrowError("Expected parameter or closing parenthesis", true, debug.NewHint("Did you forget to add a parameter?", ", <name>: <type>"), debug.NewHint("Did you forget to add a closing parenthesis?", ")"))
		}
	}

	p.Consume() // consume the closing parenthesis

	if p.Current().Kind == tokens.Minus {
		p.Consume()

		if p.Consume().Kind != tokens.GreaterThan {
			dbg := debug.NewSourceLocation(p.Source(), path.Location().Row, path.Location().Column+3)
			dbg.ThrowError("Expected arrow", true, debug.NewHint("Did you forget to add an arrow?", " ->"))
		}

		fun.Return, _ = p.parseType()
	} else {
		fun.Return = ast.Int32
	}

	fun.Params = params
	fun.Annotations = p.extractAnnotations()

	if !fun.HasAnnotation(ast.Native) {
		fun.Block = p.parseBlock()
	} else {
		if p.Current().Kind == tokens.LeftCurlyBracket {
			dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
			dbg.ThrowError("Native functions cannot have a body", true)
		}

		if p.Current().Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		}

		p.Consume() // consume the semicolon
	}

	return fun
}
