package expressions

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

// Parses a for loop statement that can have form:
// for <declaration>; <condition>; <step>; { <body> }
// for <condition>; <step>; { <body> }
func ParseForLoop(p Parser) (*ast.ForLoopStatement, error) {
	p.Consume()

	var declaration *ast.DeclarationStatement
	var condition *ast.BinaryExpression

	if p.Current().Kind == tokens.Identifier {
		stmt, _ := p.ParseStatement()
		if decl, ok := stmt.(*ast.DeclarationStatement); ok {
			declaration = decl
		} else if bin, ok := stmt.(*ast.BinaryExpression); ok {
			condition = bin

			p.Consume() // consume the semicolon
		} else {
			panic(fmt.Sprintf("expected declaration or binary expression, got %T", stmt))
		}
	}

	if declaration != nil {
		stmt, _ := p.ParseExpression()
		if bin, ok := stmt.(*ast.BinaryExpression); ok {
			condition = bin
		} else {
			panic(fmt.Sprintf("expected binary expression, got %T", stmt))
		}

		p.Consume() // consume the semicolon
	}

	var step *ast.DeclarationStatement
	stmt, _ := p.ParseStatement()
	if decl, ok := stmt.(*ast.DeclarationStatement); ok {
		step = decl
	} else {
		panic(fmt.Sprintf("expected declaration statement, got %T", stmt))
	}

	body, _ := ParseBlock(p, true)

	return &ast.ForLoopStatement{
		Declaration: declaration,
		Condition:   condition,
		Step:        step,
		Body:        body,
		Loc:         condition.Location(),
	}, nil
}

func ParseWhileLoop(p Parser) (*ast.WhileLoopStatement, error) {
	p.Consume()

	condition, _ := p.ParseExpression()
	body, _ := ParseBlock(p, true)

	return &ast.WhileLoopStatement{
		Condition: condition,
		Body:      body,
		Loc:       condition.Location(),
	}, nil
}
