package parser

import (
	"blom/ast"
)

func (p *Parser) parseNamespaceAccess(left ast.Statement) *ast.FunctionCall {
	p.Consume()

	member := p.parseLiteral()
	call := p.parseFunctionCall(member)

	call.Name = left.(*ast.IdentifierLiteral).Value + "." + member.(*ast.IdentifierLiteral).Value

	return call
}
