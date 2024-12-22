package consumers

import (
	"blom/tokens"
	"unicode"
)

func ConsumeNumber(lex Lexer) *tokens.Token {
	startLocation := lex.Location()
	value := ""
	isFloat := false

	char := lex.CurrentChar()
	for unicode.IsDigit(char) || char == '.' && !isFloat || char == '_' {
		if char == '.' {
			isFloat = true
		}
		if char != '_' {
			value += string(char)
		}
		err := lex.Advance()
		if err != nil {
			break
		}
		char = lex.CurrentChar()
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
