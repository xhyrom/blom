package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	mutable := false
	if p.Consume().Kind == tokens.Val {
		mutable = true
	}

	argument := p.parseArgument()
	name := argument.Name
	ty := argument.Type

	if p.Current().Kind != tokens.Assign {
		return &ast.VariableDeclaration{
			Name:    name,
			Type:    ast.Int32,
			Mutable: mutable,
			Loc:     argument.Location(),
		}
	}

	p.Consume()
	value := p.parseExpression()

	if value == nil {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected value for variable declaration", true, debug.NewHint("Did you forget to add a value?", ""))
	}

	return &ast.VariableDeclaration{
		Name:    name,
		Type:    ty,
		Init:    value,
		Mutable: mutable,
		Loc:     value.Location(),
	}
}

func (p *Parser) parseAssignment(left ast.Node) ast.Node {
	operator := p.Consume()

	right := p.parseExpression()

	if right == nil {
		dbg := debug.NewSourceLocationFromToken(p.Source(), operator)
		dbg.ThrowError("Expected right operand", true, debug.NewHint("Did you forget to add a right operand?", ""))
	}

	return &ast.Assignment{
		Left:  left,
		Right: right,
		Loc:   right.Location(),
	}
}
