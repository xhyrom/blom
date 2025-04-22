package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a loop that can have a form:
// while <condition> { <block> }
func (p *Parser) parseWhileLoop() *ast.WhileLoop {
	p.Consume()

	condition := p.parseExpression()
	if condition == nil {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError("Expected condition", true, debug.NewHint("Did you forget to add a condition?", "<condition>"))
	}

	block := p.parseBlock()

	return &ast.WhileLoop{
		Condition: condition,
		Block:     block,
	}
}

// Parses a loop that can have a form:
// for <declaration>; <condition>; <step> { <block> }
// for <condition>; <step> { <block> }
func (p *Parser) parseForLoop() ast.Node {
	p.Consume()

	var declaration *ast.VariableDeclaration
	var condition ast.Node
	var step ast.Node

	stmt := p.parseStatement()
	if stmt == nil {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError("Expected declaration or condition", true, debug.NewHint("Did you forget to add a declaration, condition and step?", " <declaration>; <condition>; <step>"), debug.NewHint("Did you forget to add a condition and step?", " <condition>; <step>"))
	}

	if stmt.Kind() == ast.VariableDeclarationNode {
		declaration = stmt.(*ast.VariableDeclaration)
	} else {
		condition = stmt
	}

	if declaration != nil {
		condition = p.parseExpression()

		if p.Consume().Kind != tokens.Semicolon {
			dbg := debug.NewSourceLocationFromNode(p.Source(), condition)
			dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
		}
	}

	step = p.parseExpression()

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

	return &ast.WhileLoop{
		Condition: condition,
		Block:     block,
		Loc:       condition.Location(),
	}
}
