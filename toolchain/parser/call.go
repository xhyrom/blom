package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parseCall(left ast.Node) ast.Call {
	switch left.Kind() {
	case ast.IdentifierNode, ast.PathNode:
		return p.parseFunctionCall(left)
	case ast.FieldNode:
		return p.parseMethodCall(left.(*ast.Field))
	}

	dbg := debug.NewSourceLocation(p.Source(), left.Location().Row, left.Location().Column)
	dbg.ThrowError("Expected identifier or path", true, debug.NewHint("Did you forget to add a function name?", "fn"))

	return nil
}

// Parses a function call that can have a form:
// <name>(<arguments>)
func (p *Parser) parseFunctionCall(left ast.Node) *ast.FunctionCall {
	p.Consume()

	var path *ast.Path
	if left.Kind() == ast.IdentifierNode {
		path = &ast.Path{
			Segments: []ast.IdentifierLiteral{*left.(*ast.IdentifierLiteral)},
		}
	} else {
		path = left.(*ast.Path)
	}

	args := make([]ast.Node, 0)

	for p.Current().Kind != tokens.RightParenthesis {
		arg := p.parseExpression()
		if arg == nil {
			break
		}

		args = append(args, arg)

		if p.Current().Kind == tokens.Comma {
			p.Consume()
		}
	}

	last := p.Consume()

	if last.Kind != tokens.RightParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), left.Location().Row, left.Location().Column)
		dbg.ThrowError("Expected closing parenthesis", true, debug.NewHint("Did you forget to add a closing parenthesis?", ")"))
	}

	return &ast.FunctionCall{
		Path: *path,
		Args: args,
		Loc:  last.Location,
	}
}

// Parses a method call that can have a form:
// <receiver>.<name>(<arguments>)
func (p *Parser) parseMethodCall(left *ast.Field) *ast.MethodCall {
	p.Consume()

	args := make([]ast.Node, 0)

	for p.Current().Kind != tokens.RightParenthesis {
		arg := p.parseExpression()
		if arg == nil {
			break
		}

		args = append(args, arg)

		if p.Current().Kind == tokens.Comma {
			p.Consume()
		}
	}

	last := p.Consume()

	if last.Kind != tokens.RightParenthesis {
		dbg := debug.NewSourceLocation(p.Source(), left.Location().Row, left.Location().Column)
		dbg.ThrowError("Expected closing parenthesis", true, debug.NewHint("Did you forget to add a closing parenthesis?", ")"))
	}

	return &ast.MethodCall{
		Callee: left,
		Args:   args,
		Loc:    last.Location,
	}
}

// Parses an infix call that can have a form:
// <left> <name> <right>
func (p *Parser) parseInfixCall(left ast.Node) *ast.InfixCall {
	function := p.parseExpression()
	right := p.parseExpressionWithPrecedence(tokens.HighestPrecedence)

	switch function.Kind() {
	case ast.IdentifierNode:
		return &ast.InfixCall{
			FunctionCall: &ast.FunctionCall{
				Path: ast.Path{
					Segments: []ast.IdentifierLiteral{*function.(*ast.IdentifierLiteral)},
				},
				Args: []ast.Node{left, right},
				Loc:  function.Location(),
			},
		}
	case ast.PathNode:
		return &ast.InfixCall{
			FunctionCall: &ast.FunctionCall{
				Path: *function.(*ast.Path),
				Args: []ast.Node{left, right},
				Loc:  function.Location(),
			},
		}
	case ast.FieldNode:
		return &ast.InfixCall{
			MethodCall: &ast.MethodCall{
				Callee: function,
				Args:   []ast.Node{left, right},
				Loc:    function.Location(),
			},
		}
	default:
		dbg := debug.NewSourceLocation(p.Source(), function.Location().Row, function.Location().Column)
		dbg.ThrowError("Expected identifier or path", true, debug.NewHint("Did you forget to add a function name?", "fn"))
	}

	return nil
}
