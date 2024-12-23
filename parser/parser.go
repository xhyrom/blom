package parser

import (
	"blom/ast"
	"blom/lexer"
	"blom/parser/expressions"
	"blom/tokens"
	"fmt"
	"strconv"
)

type Parser struct {
	tokens   []tokens.Token
	location *tokens.Location
}

func New(file string) *Parser {
	return &Parser{
		tokens: make([]tokens.Token, 0),
		location: &tokens.Location{
			File: file,
			Row:  1,
			Col:  0,
		},
	}
}

func (p *Parser) AST(file string, code string) *ast.Program {
	lexer := lexer.New(file, code)

	for {
		token := lexer.Next()
		fmt.Printf("Parsing token %s\n", token.Kind)

		p.tokens = append(p.tokens, *token)

		if token.Kind == tokens.Eof {
			break
		}
	}

	prog := &ast.Program{
		Loc: tokens.Location{
			File: file,
			Row:  1,
			Col:  0,
		},
	}

	for !p.IsEof() {
		prog.Body = append(prog.Body, p.ParseStatement())
	}

	return prog
}

func (p *Parser) IsEof() bool {
	return p.tokens[0].Kind == tokens.Eof
}

func (p *Parser) Current() tokens.Token {
	return p.tokens[0]
}

func (p *Parser) Consume() tokens.Token {
	prev := p.tokens[0]
	p.Advance()

	return prev
}

func (p *Parser) Advance() {
	p.tokens = p.tokens[1:]
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.Current().Kind {
	case tokens.Fun:
		return expressions.ParseFunction(p)
	case tokens.Return:
		return expressions.ParseReturn(p)
	default:
		return p.ParseExpression()
	}
}

func (p *Parser) ParseExpression() ast.Expression {
	return p.parseExpressionWithPrecedence(tokens.LowestPrecedence)
}

func (p *Parser) parseExpressionWithPrecedence(precedence tokens.Precedence) ast.Expression {
	left := p.ParsePrimaryExpression()

	for !p.IsEof() && precedence < p.Current().Kind.Precedence() {
		op := p.Current()
		p.Consume()

		right := p.parseExpressionWithPrecedence(op.Kind.Precedence())

		left = &ast.BinaryExpression{
			Left:     left,
			Operator: op.Kind,
			Right:    right,
			Loc:      op.Location,
		}
	}

	return left
}

func (p *Parser) ParsePrimaryExpression() ast.Expression {
	switch p.Current().Kind {
	case tokens.IntLiteral:
		value, _ := strconv.ParseInt(p.Consume().Value, 10, 64)
		return &ast.IntLiteralStatement{
			Value: value,
			Loc:   p.Current().Location,
		}
	case tokens.FloatLiteral:
		value, _ := strconv.ParseFloat(p.Consume().Value, 64)
		return &ast.FloatLiteralStatement{
			Value: value,
			Loc:   p.Current().Location,
		}
	case tokens.Identifier:
		return expressions.ParseIdentifier(p)
	case tokens.If:
		return expressions.ParseIf(p)
	case tokens.LeftParenthesis:
		p.Consume() // Consume '('
		expr := p.ParseExpression()
		p.Consume() // Consume ')'
		return expr
	case tokens.Plus, tokens.Minus:
		return expressions.ParseUnary(p)
	case tokens.LeftCurlyBracket:
		return expressions.ParseBlock(p, true)
	}

	return nil
}
