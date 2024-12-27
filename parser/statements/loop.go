package statements

import (
	"blom/ast"
	"blom/debug"
	"blom/parser/expressions"
	"blom/tokens"
	"fmt"
)

// Parses a for loop statement that can have form:
// for <declaration>; <condition>; <step>; { <body> }
// for <condition>; <step>; { <body> }
func ParseForLoop(p Parser) *ast.ForLoopStatement {
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
			dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
			dbg.ThrowError(fmt.Sprintf("Expected declaration or binary expression, got %T", stmt), true)
		}
	}

	if declaration != nil {
		stmt, _ := p.ParseExpression()
		if bin, ok := stmt.(*ast.BinaryExpression); ok {
			condition = bin
		} else {
			dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
			dbg.ThrowError(fmt.Sprintf("Expected binary expression, got %T", stmt), true)
		}

		p.Consume() // consume the semicolon
	}

	var step *ast.DeclarationStatement
	location := p.Current().Location
	stmt, _ := p.ParseStatement()
	if decl, ok := stmt.(*ast.DeclarationStatement); ok {
		step = decl
	} else {
		dbg := debug.NewSourceLocation(p.Source(), location.Row, location.Column)
		dbg.ThrowError(fmt.Sprintf("Expected declaration, got %T", stmt), true)
	}

	body := expressions.ParseBlock(p)

	return &ast.ForLoopStatement{
		Declaration: declaration,
		Condition:   condition,
		Step:        step,
		Body:        body,
		Loc:         condition.Location(),
	}
}

func ParseWhileLoop(p Parser) *ast.WhileLoopStatement {
	p.Consume()

	condition, err := p.ParseExpression()
	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError(err.Error(), true, debug.NewHint("Add condition", "true "))
	}

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), condition)
		dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
	}

	body := expressions.ParseBlock(p)

	return &ast.WhileLoopStatement{
		Condition: condition,
		Body:      body,
		Loc:       condition.Location(),
	}
}
