package parser

import (
	"blom/ast"
	"blom/lexer"
	"blom/parser/expressions"
	"blom/tokens"
	"errors"
	"strconv"
)

type Parser struct {
	tokens []tokens.Token
	source string
}

func New(file string) *Parser {
	return &Parser{
		tokens: make([]tokens.Token, 0),
		source: file,
	}
}

func (p *Parser) AST(file string, code string) *ast.Program {
	lexer := lexer.New(file, code)

	for {
		token := lexer.Next()

		p.tokens = append(p.tokens, *token)

		if token.Kind == tokens.Eof {
			break
		}
	}

	prog := &ast.Program{
		Loc: tokens.Location{
			Row:    1,
			Column: 0,
		},
	}

	for !p.IsEof() {
		stmt, _ := p.ParseStatement()
		prog.Body = append(prog.Body, stmt)
	}

	return prog
}

func (p *Parser) Source() string {
	return p.source
}

func (p *Parser) IsEof() bool {
	if len(p.tokens) == 0 {
		return true
	}

	return p.tokens[0].Kind == tokens.Eof
}

func (p *Parser) Current() tokens.Token {
	return p.tokens[0]
}

func (p *Parser) Next() tokens.Token {
	return p.tokens[1]
}

func (p *Parser) Consume() tokens.Token {
	prev := p.tokens[0]
	p.Advance()

	return prev
}

func (p *Parser) Advance() {
	p.tokens = p.tokens[1:]
}

func (p *Parser) ParseStatement() (ast.Statement, error) {
	switch p.Current().Kind {
	case tokens.Fun:
		return expressions.ParseFunction(p)
	case tokens.Return:
		return expressions.ParseReturn(p)
	case tokens.For:
		return expressions.ParseForLoop(p)
	case tokens.While:
		return expressions.ParseWhileLoop(p)
	case tokens.Identifier:
		if p.Next().Kind == tokens.Identifier {

			return expressions.ParseAssignment(p, false)
		}

		if p.Next().Kind == tokens.Assign {
			return expressions.ParseAssignment(p, true)
		}
	}

	return p.ParseExpression()
}

func (p *Parser) ParseExpression() (ast.Expression, error) {
	return p.parseExpressionWithPrecedence(tokens.LowestPrecedence)
}

func (p *Parser) parseExpressionWithPrecedence(precedence tokens.Precedence) (ast.Expression, error) {
	left, err := p.ParsePrimaryExpression()
	if err != nil {
		return nil, err
	}

	for !p.IsEof() && precedence < p.Current().Kind.Precedence() {
		op := p.Current()
		p.Consume()

		right, err := p.parseExpressionWithPrecedence(op.Kind.Precedence())
		if err != nil {
			return nil, err
		}

		left = &ast.BinaryExpression{
			Left:     left,
			Operator: op.Kind,
			Right:    right,
			Loc:      right.Location(),
		}
	}

	return left, nil
}

func (p *Parser) ParsePrimaryExpression() (ast.Expression, error) {
	switch p.Current().Kind {
	case tokens.CharLiteral:
		token := p.Consume()
		value := []rune(token.Value)[0]
		return &ast.CharLiteralStatement{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.StringLiteral:
		token := p.Consume()
		value := token.Value
		return &ast.StringLiteralStatement{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.IntLiteral:
		token := p.Consume()
		value, _ := strconv.ParseInt(token.Value, 10, 64)
		return &ast.IntLiteralStatement{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.FloatLiteral:
		token := p.Consume()
		value, _ := strconv.ParseFloat(token.Value, 64)
		return &ast.FloatLiteralStatement{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.Identifier:
		return expressions.ParseIdentifier(p)
	case tokens.If:
		return expressions.ParseIf(p)
	case tokens.LeftParenthesis:
		p.Consume() // consume '('
		expr, err := p.ParseExpression()
		p.Consume() // consume ')'

		expr.SetLocation(expr.Location().Row, expr.Location().Column+1)
		return expr, err
	case tokens.Plus, tokens.Minus, tokens.Tilde:
		return expressions.ParseUnary(p)
	case tokens.LeftCurlyBracket:
		return expressions.ParseBlock(p, true)
	}

	return nil, errors.New("Unexpected token " + p.Current().Kind.String())
}
