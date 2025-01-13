package expressions

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
	"strconv"
)

func ParseLiteral(p Parser) (ast.Expression, error) {
	switch p.Current().Kind {
	case tokens.CharLiteral:
		token := p.Consume()
		value := []rune(token.Value)[0]
		return &ast.CharLiteral{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.StringLiteral:
		token := p.Consume()
		value := token.Value
		return &ast.StringLiteral{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.IntLiteral:
		token := p.Consume()
		value, _ := strconv.ParseInt(token.Value, 10, 64)
		return &ast.IntLiteral{
			Value: int64(value),
			Loc:   token.Location,
		}, nil
	case tokens.FloatLiteral:
		token := p.Consume()
		value, _ := strconv.ParseFloat(token.Value, 64)
		return &ast.FloatLiteral{
			Value: float64(value),
			Loc:   token.Location,
		}, nil
	case tokens.BooleanLiteral:
		token := p.Consume()
		value, _ := strconv.ParseBool(token.Value)
		return &ast.BooleanLiteral{
			Value: value,
			Loc:   token.Location,
		}, nil
	case tokens.Identifier:
		return ParseIdentifier(p), nil
	}

	panic(fmt.Sprintf("unexpected literal %T", p.Current()))
}
