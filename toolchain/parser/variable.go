package parser

import (
	"blom/ast"
	"blom/tokens"
)

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	ty := p.parseType()
	literal := p.parseLiteral()
	name := literal.(*ast.IdentifierLiteral)

	if p.Current().Kind != tokens.Assign {
		return &ast.VariableDeclaration{
			Id:   name,
			Type: ast.Int32,
			Loc:  literal.Location(),
		}
	}

	assign := p.Consume()
	value := p.parseExpression()

	return &ast.VariableDeclaration{
		Id:   name,
		Type: ty,
		Init: value,
		Loc:  assign.Location,
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
