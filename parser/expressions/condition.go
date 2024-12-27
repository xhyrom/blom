package expressions

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

func ParseIf(p Parser) (*ast.IfStatement, error) {
	p.Consume()

	condition, _ := p.ParseExpression()

	if p.Current().Kind != tokens.LeftCurlyBracket {
		fmt.Println("Expected {")
	}

	then_block, _ := ParseBlock(p, true)
	else_block := &ast.BlockStatement{}

	loc := then_block.Loc

	if p.Current().Kind == tokens.Else {
		p.Consume()
		block, _ := ParseBlock(p, true)

		else_block = block
		loc = else_block.Loc
	}

	return &ast.IfStatement{
		Condition: condition,
		Then:      then_block,
		Else:      else_block,
		Loc:       loc,
	}, nil
}
