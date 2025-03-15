package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/lexer"
	"blom/tokens"
)

type Parser struct {
	tokens      []tokens.Token
	source      string
	annotations []ast.Annotation
}

func New(file string) *Parser {
	return &Parser{
		tokens:      make([]tokens.Token, 0),
		source:      file,
		annotations: make([]ast.Annotation, 0),
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
		stmt := p.parseStatement()
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

func (p *Parser) Peek(i int) tokens.Token {
	return p.tokens[i]
}

func (p *Parser) Consume() tokens.Token {
	prev := p.tokens[0]
	p.Advance()

	return prev
}

func (p *Parser) Advance() {
	p.tokens = p.tokens[1:]
}

func (p *Parser) parseStatement() ast.Statement {
	p.collectAnnotations()

	var statement ast.Statement

	switch p.Current().Kind {
	case tokens.Fun:
		return p.parseFunction()
	case tokens.Return:
		statement = p.parseReturn()
	case tokens.Identifier:
		statement = p.parseVariableDeclaration()
	}

	if statement == nil {
		statement = p.parseExpression()
	}

	if p.Consume().Kind != tokens.Semicolon {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Current())
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
	}

	return statement
}

func (p *Parser) parseExpression() ast.Statement {
	return p.parseExpressionWithPrecedence(tokens.LowestPrecedence)
}

func (p *Parser) parseExpressionWithPrecedence(precedence tokens.Precedence) ast.Statement {
	// prefix
	var left ast.Statement

	switch p.Current().Kind {
	case tokens.Identifier, tokens.IntLiteral, tokens.FloatLiteral, tokens.StringLiteral, tokens.CharLiteral, tokens.BooleanLiteral:
		left = p.parseLiteral()
	case tokens.LeftParenthesis:
		left = p.parseGroupedExpression()
	}

	for p.Current().Kind != tokens.Semicolon && precedence < p.Current().Kind.Precedence() {
		// infix
		switch p.Current().Kind {
		case tokens.Plus,
			tokens.Minus,
			tokens.Asterisk,
			tokens.Slash,
			tokens.PercentSign,
			tokens.VerticalLine,
			tokens.CircumflexAccent:
			left = p.parseInfixExpression(left)
		case tokens.Dot:
			left = p.parseMemberAccess(left)
		case tokens.DoubleColon:
			left = p.parseNamespaceAccess(left)
		default:
			return left
		}
	}

	return left
}

func (p *Parser) parseGroupedExpression() ast.Statement {
	p.Consume() // consume "("

	exp := p.parseExpression()

	if p.Consume().Kind != tokens.RightParenthesis {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), exp)
		dbg.ThrowError("Expected closing parenthesis", true, debug.NewHint("Did you forget to add a closing parenthesis?", ")"))
	}

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Statement) ast.Statement {
	operator := p.Consume()
	right := p.parseExpressionWithPrecedence(operator.Kind.Precedence())

	return &ast.BinaryExpression{
		Left:     left,
		Operator: operator.Kind,
		Right:    right,
		Loc:      operator.Location,
	}
}
