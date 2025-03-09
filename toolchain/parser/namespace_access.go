package parser

import "blom/ast"

func (p *Parser) parseNamespaceAccess(left ast.Statement) *ast.FunctionCall {
	p.Consume()

	member := p.parseLiteral()

	return &ast.FunctionCall{
		Name: left.(*ast.IdentifierLiteral).Value + "." + member.(*ast.IdentifierLiteral).Value,
		Loc:  left.Location(),
	}
}
