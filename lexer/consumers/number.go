package consumers

import (
	"blom/tokens"
	"unicode"
)

func ConsumeNumber(lex Lexer) *tokens.Token {
	startLocation := lex.Location()
	value := ""
	radix := 10
	isFloat := false

	char := lex.CurrentChar()

	if char == '0' {
		err := lex.Advance()
		if err != nil {
			return nil
		}

		char = lex.CurrentChar()

		switch char {
		case 'x':
			radix = 16
		case 'o':
			radix = 8
		case 'b':
			radix = 2
		default:
			lex.Rewind()
			lex.Rewind()
		}

		err = lex.Advance()
		if err != nil {
			return nil
		}

		char = lex.CurrentChar()
	}

	for unicode.IsDigit(char) || (radix == 16 && unicode.Is(unicode.Hex_Digit, char)) || char == '.' && !isFloat || char == '_' {
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

	lex.Rewind()

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
