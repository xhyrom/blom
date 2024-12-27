package expressions

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

func ParseIf(p Parser) *ast.IfStatement {
	p.Consume()

	condition := p.ParseExpression()

	if p.Current().Kind != tokens.LeftCurlyBracket {
		fmt.Println("Expected {")
	}

	then_block := ParseBlock(p, true)
	else_block := &ast.BlockStatement{}

	if p.Current().Kind == tokens.Else {
		p.Consume()
		else_block = ParseBlock(p, true)
	}

	return &ast.IfStatement{
		Condition: condition,
		Then:      then_block,
		Else:      else_block,
	}
}
