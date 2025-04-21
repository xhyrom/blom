package parser

import (
	"blom/ast"
	"blom/tokens"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (p *Parser) parseLiteral() ast.Node {
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
		return parseInt(p)
	case tokens.FloatLiteral:
		return parseFloat(p)
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

func parseInt(p *Parser) ast.Node {
	token := p.Consume()
	value := token.Value
	cleanValue := strings.Replace(value, "_", "", -1)

	var val int64
	var err error

	if strings.HasPrefix(cleanValue, "0x") || strings.HasPrefix(cleanValue, "0X") {
		// hexadecimal
		val, err = strconv.ParseInt(cleanValue[2:], 16, 64)
	} else if strings.HasPrefix(cleanValue, "0b") || strings.HasPrefix(cleanValue, "0B") {
		// binary
		val, err = strconv.ParseInt(cleanValue[2:], 2, 64)
	} else if strings.HasPrefix(cleanValue, "0o") || strings.HasPrefix(cleanValue, "0O") {
		// octal
		val, err = strconv.ParseInt(cleanValue[2:], 8, 64)
	} else if strings.ContainsAny(cleanValue, "eE") {
		// scientific notation
		f, err := strconv.ParseFloat(cleanValue, 64)
		if err != nil {
			panic(fmt.Sprintf("invalid scientific notation number: %s", value))
		}

		// check if it's a whole number
		if f == math.Trunc(f) {
			val = int64(f)
		} else {
			panic(fmt.Sprintf("scientific notation number is not an integer: %s", value))
		}

		return &ast.IntLiteral{
			Value: val,
			Loc:   token.Location,
		}
	} else {
		// regular decimal integer
		val, err = strconv.ParseInt(cleanValue, 10, 64)
	}

	if err != nil {
		panic(fmt.Sprintf("invalid integer: %s (%s)", value, err))
	}

	return &ast.IntLiteral{
		Value: val,
		Loc:   token.Location,
	}
}

func parseFloat(p *Parser) ast.Node {
	token := p.Consume()
	value := token.Value
	cleanValue := strings.Replace(value, "_", "", -1)
	val, err := strconv.ParseFloat(cleanValue, 64)

	if err != nil {
		panic(fmt.Sprintf("invalid floating-point number: %s (%s)", value, err))
	}

	return &ast.FloatLiteral{
		Value: val,
		Loc:   token.Location,
	}
}

func parseIdentifier(p *Parser) ast.Node {
	token := p.Consume()

	return &ast.IdentifierLiteral{
		Value: token.Value,
		Loc:   token.Location,
	}
}
