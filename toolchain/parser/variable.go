package parser

import (
	"blom/ast"
	"blom/tokens"
)

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	ty := p.parseType()
	literal := p.parseLiteral()
	name := literal.(*ast.IdentifierLiteral).Value

	if p.Current().Kind != tokens.Assign {
		return &ast.VariableDeclaration{
			Name:        name,
			Type:        ast.Int32,
			Value:       &ast.IntLiteral{Value: 0},
			Annotations: p.extractAnnotations(),
			Loc:         literal.Location(),
		}
	}

	assign := p.Consume()
	value := p.parseExpression()

	return &ast.VariableDeclaration{
		Name:        name,
		Type:        ty,
		Value:       value,
		Annotations: p.extractAnnotations(),
		Loc:         assign.Location,
	}
}

func (p *Parser) parseAssignment(left ast.Expression) ast.Expression {
	p.Consume()

	right := p.parseExpression()

	return &ast.Assignment{
		Left:  left,
		Right: right,
		Loc:   right.Location(),
	}
}
