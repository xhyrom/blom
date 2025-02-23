package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/lexer"
	"blom/tokens"
	"errors"
	"fmt"
)

type Parser struct {
	tokens      []tokens.Token
	source      string
	customTypes map[string]ast.Type
	annotations []ast.Annotation
}

type Context int

const (
	Statement Context = iota
	Expression
)

func New(file string) *Parser {
	return &Parser{
		tokens:      make([]tokens.Token, 0),
		source:      file,
		customTypes: make(map[string]ast.Type),
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

func (p *Parser) CustomTypes() map[string]ast.Type {
	return p.customTypes
}

func (p *Parser) AddCustomType(name string, ty ast.Type) {
	p.customTypes[name] = ty
}

func (p *Parser) ParseStatement() (ast.Statement, error) {
	p.collectAnnotations()

	switch p.Current().Kind {
	case tokens.Fun:
		return p.parseFunction(), nil
	case tokens.Return:
		return p.parseReturn(), nil
	case tokens.Type:
		return p.parseTypeDefinition(), nil
	case tokens.Entity:
		return p.parseEntity(), nil
	case tokens.For:
		return p.ParseForLoop(), nil
	case tokens.While:
		return p.ParseWhileLoop(), nil
	case tokens.Identifier:
		if p.Next().Kind == tokens.Identifier ||
			(p.Next().Kind == tokens.Asterisk && p.Peek(2).Kind == tokens.Identifier) {
			return p.parseVariableDeclaration(), nil
		}

		if p.Next().Kind == tokens.LeftParenthesis {
			return p.parseFunctionCall(p.Consume(), true), nil
		}

		if p.Next().Kind == tokens.DoubleColon {
			token := p.Consume()

			_, err := ast.ParseType(token.Value, p.CustomTypes())
			if err != nil {
				dbg := debug.NewSourceLocation(p.Source(), token.Location.Row, token.Location.Column)
				dbg.ThrowError(
					fmt.Sprintf(
						"Cannot extend type \"%s\" because it isn't a primitive type",
						token.Value,
					),
					true,
				)
			}

			p.Consume()
			identifier := p.Consume()
			if identifier.Kind != tokens.Identifier {
				dbg := debug.NewSourceLocation(p.Source(), identifier.Location.Row, identifier.Location.Column)
				dbg.ThrowError(fmt.Sprintf("Function name must be valid identifier, got \"%s\"", token.Value), true)
			}

			token = tokens.Token{
				Kind:     tokens.Identifier,
				Location: token.Location,
				Value:    token.Value + "." + identifier.Value,
			}

			return p.parseFunctionCall(token, true), nil
		}

		if p.Next().Kind == tokens.Dot {
			exp := p.parseMemberAccess(p.parseLiteral())

			if p.Consume().Kind != tokens.Semicolon {
				dbg := debug.NewSourceLocation(p.Source(), exp.Location().Row, exp.Location().Column+1)
				dbg.ThrowError("Expected semicolon", true, debug.NewHint("Did you forget to add a semicolon?", ";"))
			}

			return exp, nil
		}
	}

	exp, err := p.ParseExpression()
	return exp, err
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
			Left:        left,
			Operator:    op.Kind,
			Right:       right,
			Loc:         right.Location(),
			OperatorLoc: op.Location,
		}
	}

	return left, nil
}

func (p *Parser) ParsePrimaryExpression() (ast.Expression, error) {
	// parse cases that can't be infix
	switch p.Current().Kind {
	case tokens.LeftCurlyBracket:
		return p.parseBlock(), nil
	case tokens.If:
		return p.parseCondition(), nil
	case tokens.AtMark:
		return p.parseBuiltinFunctionCall(), nil
	case tokens.LeftParenthesis:
		p.Consume() // consume '('
		expr, err := p.ParseExpression()
		p.Consume() // consume ')'
		expr.SetLocation(expr.Location().Row, expr.Location().Column+1)

		if p.Current().Kind == tokens.Dot {
			expr = p.parseMemberAccess(expr)
		}

		return expr, err
	}

	left, err := p.parseSingleExpression()
	if err != nil {
		return nil, err
	}

	if !p.IsEof() && p.Current().Kind == tokens.Assign {
		return p.parseAssignment(left), nil
	}

	if !p.IsEof() && p.Current().Kind == tokens.Identifier && p.Next().Kind != tokens.Asterisk {
		op := p.Current()

		p.Consume()
		if !p.IsEof() && p.Current().Kind != tokens.LeftCurlyBracket && p.Current().Kind != tokens.Dot &&
			p.Current().Kind != tokens.If && p.Current().Kind != tokens.LeftParenthesis {

			right, err := p.ParsePrimaryExpression()

			if err == nil {
				return &ast.FunctionCall{
					Name: op.Value,
					Parameters: []ast.Expression{
						left,
						right,
					},
					Infix: true,
					Loc:   op.Location,
				}, nil
			}
		}

		// restore
		p.tokens = append([]tokens.Token{op}, p.tokens...)
	}

	return left, nil
}

func (p *Parser) parseSingleExpression() (ast.Expression, error) {
	var exp ast.Expression

	switch p.Current().Kind {
	case tokens.CharLiteral,
		tokens.StringLiteral,
		tokens.IntLiteral,
		tokens.FloatLiteral,
		tokens.BooleanLiteral,
		tokens.Identifier:
		exp = p.parseLiteral()
	case tokens.AtMark:
		exp = p.parseBuiltinFunctionCall()
	case tokens.Plus, tokens.Minus, tokens.Tilde, tokens.Ampersand, tokens.Asterisk:
		exp = p.parseUnaryExpression()
	case tokens.Fun:
		exp = p.parseLambda()
	}

	if p.Current().Kind == tokens.Dot {
		exp = p.parseMemberAccess(exp)
	}

	if exp != nil {
		return exp, nil
	}

	return nil, errors.New("Unexpected token " + p.Current().Kind.String())
}
