package parser

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
	"strconv"
)

func (p *Parser) parseLiteral() ast.Expression {
	switch p.Current().Kind {
	case tokens.CharLiteral:
		token := p.Consume()
		value := []rune(token.Value)[0]
		return &ast.CharLiteral{
			Value: value,
			Loc:   token.Location,
		}
	case tokens.StringLiteral:
		token := p.Consume()
		value := token.Value
		return &ast.StringLiteral{
			Value: value,
			Loc:   token.Location,
		}
	case tokens.IntLiteral:
		token := p.Consume()
		value, _ := strconv.ParseInt(token.Value, 10, 64)
		return &ast.IntLiteral{
			Value: int64(value),
			Loc:   token.Location,
		}
	case tokens.FloatLiteral:
		token := p.Consume()
		value, _ := strconv.ParseFloat(token.Value, 64)
		return &ast.FloatLiteral{
			Value: float64(value),
			Loc:   token.Location,
		}
	case tokens.BooleanLiteral:
		token := p.Consume()
		value, _ := strconv.ParseBool(token.Value)
		return &ast.BooleanLiteral{
			Value: value,
			Loc:   token.Location,
		}
	case tokens.Identifier:
		return parseIdentifier(p)
	}

	panic(fmt.Sprintf("unexpected literal %T", p.Current()))
}

func parseIdentifier(p *Parser) ast.Statement {
	token := p.Consume()

	return &ast.IdentifierLiteral{
		Value: token.Value,
		Loc:   token.Location,
	}
}
