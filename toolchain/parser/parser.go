package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/lexer"
	"blom/tokens"
)

type Parser struct {
	tokens      []tokens.Token
	previous    tokens.Token
	source      string
	annotations []ast.Annotation
}

func New(file string) *Parser {
	return &Parser{
		tokens:      make([]tokens.Token, 0),
		previous:    tokens.Token{},
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

	prog := ast.NewProgram()
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

func (p *Parser) Previous() tokens.Token {
	return p.previous
}

func (p *Parser) Peek(i int) tokens.Token {
	return p.tokens[i]
}

func (p *Parser) Consume() tokens.Token {
	prev := p.tokens[0]
	p.Advance()

	p.previous = prev

	return prev
}

func (p *Parser) Advance() {
	p.previous = p.tokens[0]
	p.tokens = p.tokens[1:]
}

func (p *Parser) parseStatement() ast.Node {
	p.collectAnnotations()

	var statement ast.Node

	switch p.Current().Kind {
	case tokens.Fun:
		return p.parseFunction()
	case tokens.Return:
		statement = p.parseReturn()
	case tokens.If:
		return p.parseCondition()
	case tokens.For:
		return p.parseForLoop()
	case tokens.Val, tokens.Var:
		statement = p.parseVariableDeclaration()
	case tokens.LeftCurlyBracket:
		return p.parseBlock()
	}

	if statement == nil {
		statement = p.parseExpression()
	}

	if statement != nil && (p.IsEof() || p.Consume().Kind != tokens.Semicolon) {
		dbg := debug.NewSourceLocationFromNode(p.Source(), statement)
		dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
	}

	return statement
}

func (p *Parser) parseExpression() ast.Node {
	return p.parseExpressionWithPrecedence(tokens.LowestPrecedence)
}

func (p *Parser) parseExpressionWithPrecedence(precedence tokens.Precedence) ast.Node {
	// prefix
	var left ast.Node

	switch p.Current().Kind {
	case tokens.Identifier, tokens.IntLiteral, tokens.FloatLiteral, tokens.StringLiteral, tokens.CharLiteral, tokens.BooleanLiteral:
		left = p.parseLiteral()
	case tokens.LeftParenthesis:
		left = p.parseGroupedExpression()
	case tokens.If:
		left = p.parseCondition()
	case tokens.LeftCurlyBracket:
		left = p.parseBlock()
	case tokens.Plus, tokens.Minus, tokens.Ampersand, tokens.Tilde, tokens.Asterisk:
		left = p.parseUnaryExpression()
	}

	for p.Current().Kind != tokens.Semicolon && precedence < p.Current().Kind.Precedence() {
		// infix
		switch p.Current().Kind {
		case tokens.Equals,
			tokens.NotEquals,
			tokens.Plus,
			tokens.Minus,
			tokens.Asterisk,
			tokens.Slash,
			tokens.PercentSign,
			tokens.VerticalLine,
			tokens.CircumflexAccent,
			tokens.LessThan,
			tokens.DoubleLessThan,
			tokens.GreaterThan,
			tokens.DoubleGreaterThan:
			left = p.parseBinaryExpression(left)
		case tokens.Identifier:
			left = p.parseInfixCall(left)
		case tokens.Dot:
			left = p.parseField(left)
		case tokens.DoubleColon:
			left = p.parsePath(left)
		case tokens.Assign:
			left = p.parseAssignment(left)
		case tokens.LeftParenthesis:
			left = p.parseCall(left)
		default:
			return left
		}
	}

	return left
}

func (p *Parser) parseGroupedExpression() ast.Node {
	p.Consume() // consume "("

	exp := p.parseExpression()

	token := p.Consume()
	if token.Kind != tokens.RightParenthesis {
		dbg := debug.NewSourceLocationFromNode(p.Source(), exp)
		dbg.ThrowError("Expected closing parenthesis", true, debug.NewHint("Did you forget to add a closing parenthesis?", ")"))
	}

	return &ast.GroupedExpression{
		Expression: exp,
		Loc:        token.Location,
	}
}

func (p *Parser) parseUnaryExpression() ast.Node {
	operator := p.Consume()
	operand := p.parseExpressionWithPrecedence(operator.Kind.Precedence())

	return &ast.UnaryExpression{
		Operator: operator.Kind,
		Operand:  operand,
		Loc:      operator.Location,
	}
}

func (p *Parser) parseBinaryExpression(left ast.Node) ast.Node {
	operator := p.Consume()
	right := p.parseExpressionWithPrecedence(operator.Kind.Precedence())

	if right == nil {
		dbg := debug.NewSourceLocationFromToken(p.Source(), operator)
		dbg.ThrowError("Expected right operand", true, debug.NewHint("Did you forget to add a right operand?", " a"))
	}

	return &ast.BinaryExpression{
		Left:     left,
		Operator: operator.Kind,
		Right:    right,
		Loc:      right.Location(),
	}
}
