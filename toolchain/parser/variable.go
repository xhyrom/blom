package parser

import (
	"blom/ast"
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
			Loc:     name.Location(),
		}
	}

	assign := p.Consume()
	value := p.parseExpression()

	return &ast.VariableDeclaration{
		Name:    name,
		Type:    ty,
		Init:    value,
		Mutable: mutable,
		Loc:     assign.Location,
	}
}

func (p *Parser) parseAssignment(left ast.Node) ast.Node {
	p.Consume()

	right := p.parseExpression()

	return &ast.Assignment{
		Left:  left,
		Right: right,
		Loc:   right.Location(),
	}
}
