package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclarationStatement {
	ty := p.parseType()
	name := p.parseLiteral().(*ast.IdentifierLiteral).Value

	assign := p.Consume()

	if assign.Kind != tokens.Assign {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Current())
		dbg.ThrowError("Expected assignment", true, debug.NewHint("Did you forget to add an assignment?", "="))
	}

	value := p.parseExpression()

	return &ast.VariableDeclarationStatement{
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
