package expressions

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
)

// Parses a for loop statement that can has form:
// for <declaration>; <condition>; <step> { <body> }
// for <condition>; <step> { <body> }
func ParseForLoop(p Parser) ast.Statement {
	p.Consume()

	var declaration *ast.DeclarationStatement
	var condition ast.BinaryExpression

	if p.Current().Kind == tokens.Identifier {
		stmt := p.ParseExpression()
		if decl, ok := stmt.(ast.DeclarationStatement); ok {
			declaration = &decl
		} else if bin, ok := stmt.(*ast.BinaryExpression); ok {
			condition = *bin
		} else {
			panic(fmt.Sprintf("expected declaration or binary expression, got %T", stmt))
		}
	}

	p.Consume() // consume the semicolon

	if declaration != nil {
		stmt := p.ParseExpression()
		if bin, ok := stmt.(*ast.BinaryExpression); ok {
			condition = *bin
		} else {
			panic(fmt.Sprintf("expected binary expression, got %T", stmt))
		}

		p.Consume() // consume the semicolon
	}

	var step ast.DeclarationStatement
	stmt := p.ParseExpression()
	if decl, ok := stmt.(ast.DeclarationStatement); ok {
		step = decl
	} else {
		panic(fmt.Sprintf("expected declaration statement, got %T", stmt))
	}

	body := ParseBlock(p, true)

	return ast.ForLoopStatement{
		Declaration: declaration,
		Condition:   condition,
		Step:        step,
		Body:        body,
		Loc:         condition.Location(),
	}
}

func ParseWhileLoop(p Parser) ast.Statement {
	p.Consume()

	condition := p.ParseExpression()
	body := ParseBlock(p, true)

	return ast.WhileLoopStatement{
		Condition: condition,
		Body:      body,
		Loc:       condition.Location(),
	}
}
