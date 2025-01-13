package consumers

import (
	"blom/tokens"
	"unicode"
)

func ConsumeNumber(lex Lexer) *tokens.Token {
	startLocation := lex.Location()
	value := ""
	isFloat := false

	var err error

	char := lex.CurrentChar()
	for unicode.IsDigit(char) || char == '.' && !isFloat || char == '_' {
		if char == '.' {
			isFloat = true
		}

		if char != '_' {
			value += string(char)
		}

		err = lex.Advance()
		if err != nil {
			break
		}

		char = lex.CurrentChar()
	}

	// check if last character is an dot, if so, rewind
	if value[len(value)-1] == '.' {
		value = value[:len(value)-1]
		isFloat = false

		lex.Rewind()
	}

	if err == nil {
		lex.Rewind()
	}

	kind := tokens.IntLiteral
	if isFloat {
		kind = tokens.FloatLiteral
	}

	return &tokens.Token{
		Kind:     kind,
		Value:    value,
		Location: startLocation.Copy(),
	}
}
