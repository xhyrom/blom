package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a loop that can have a form:
// for <declaration>; <condition>; <step> { <block> }
// for <condition>; <step> { <block> }
func (p *Parser) parseForLoop() *ast.Block {
	p.Consume()

	var declaration *ast.VariableDeclaration
	var condition ast.Node
	var step ast.Node

	stmt := p.parseStatement()
	if stmt.Kind() == ast.VariableDeclarationNode {
		declaration = stmt.(*ast.VariableDeclaration)
	} else {
		condition = stmt
	}

	if declaration != nil {
		condition = p.parseExpression()

		if p.Consume().Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocation(p.Source(), condition.Location().Row, condition.Location().Column)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		}
	}

	step = p.parseExpression()

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError("Expected opening bracket", true, debug.NewHint("Did you forget to add an opening bracket?", "{"))
	}

	block := p.parseBlock()
	block.Body = append(block.Body, step)

	if declaration != nil {
		return &ast.Block{
			Body: []ast.Node{
				declaration,
				&ast.WhileLoop{
					Condition: condition,
					Block:     block,
					Loc:       condition.Location(),
				},
			},
			Loc: block.Location(),
		}
	}

	return &ast.Block{
		Body: []ast.Node{
			&ast.WhileLoop{
				Condition: condition,
				Block:     block,
				Loc:       condition.Location(),
			},
		},
		Loc: block.Location(),
	}
}
