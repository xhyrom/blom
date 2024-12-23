package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseBlock(p Parser) ast.BlockStatement {
	p.Consume()

	current := p.Current()
	body := []ast.Statement{}

	for !p.IsEof() && current.Kind != tokens.RightCurlyBracket {
		body = append(body, p.ParseStatement())
		current = p.Current()
	}

	p.Consume()
	return ast.BlockStatement{Body: body}
}
