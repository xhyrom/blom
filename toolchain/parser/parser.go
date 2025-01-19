package parser

import (
	"blom/ast"
	"blom/lexer"
	"blom/parser/expressions"
	"blom/parser/statements"
	"blom/tokens"
	"errors"
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
		stmts, _ := p.ParseStatement()
		for _, stmt := range stmts {
			prog.Body = append(prog.Body, stmt)
		}
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

func (p *Parser) ParseStatement() ([]ast.Statement, error) {
	switch p.Current().Kind {
	case tokens.Fun:
		return []ast.Statement{statements.ParseFunction(p)}, nil
	case tokens.Return:
		return []ast.Statement{statements.ParseReturn(p)}, nil
	case tokens.Type:
		return []ast.Statement{statements.ParseTypeDefinition(p)}, nil
	case tokens.For:
		decl, while := statements.ParseForLoop(p)
		if decl != nil {
			return []ast.Statement{decl, while}, nil
		}

		return []ast.Statement{while}, nil
	case tokens.While:
		return []ast.Statement{statements.ParseWhileLoop(p)}, nil
	case tokens.Identifier:
		if p.Next().Kind == tokens.Identifier ||
			(p.Next().Kind == tokens.Asterisk && p.Peek(2).Kind == tokens.Identifier) {
			return []ast.Statement{statements.ParseVariableDeclaration(p)}, nil
		}

		if p.Next().Kind == tokens.LeftParenthesis {
			return []ast.Statement{expressions.ParseFunctionCall(p, p.Consume(), true)}, nil
		}
	}

	exp, err := p.ParseExpression()
	return []ast.Statement{exp}, err
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
		return expressions.ParseBlock(p), nil
	case tokens.If:
		return expressions.ParseIf(p), nil
	case tokens.AtMark:
		return expressions.ParseCompileTimeFunctionCall(p), nil
	case tokens.LeftParenthesis:
		p.Consume() // consume '('
		expr, err := p.ParseExpression()
		p.Consume() // consume ')'
		expr.SetLocation(expr.Location().Row, expr.Location().Column+1)
		return expr, err
	}

	left, err := p.parseSingleExpression()
	if err != nil {
		return nil, err
	}

	if !p.IsEof() && p.Current().Kind == tokens.Assign {
		return statements.ParseAssignment(p, left), nil
	}

	if !p.IsEof() && p.Current().Kind == tokens.Identifier {
		op := p.Current()

		p.Consume()
		if !p.IsEof() && p.Current().Kind != tokens.LeftCurlyBracket &&
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
		exp, _ = expressions.ParseLiteral(p)
	case tokens.AtMark:
		exp = expressions.ParseCompileTimeFunctionCall(p)
	case tokens.Plus, tokens.Minus, tokens.Tilde, tokens.Ampersand, tokens.Asterisk:
		exp = expressions.ParseUnary(p)
	case tokens.Fun:
		exp = expressions.ParseLambda(p)
	}

	if p.Current().Kind == tokens.Dot {
		exp = expressions.ParseMemberAccess(p, exp)
	}

	if exp != nil {
		return exp, nil
	}

	return nil, errors.New("Unexpected token " + p.Current().Kind.String())
}
