package consumers

import (
	"blom/tokens"
	"unicode"
)

func ConsumeIdentifier(lex Lexer) *tokens.Token {
	startLocation := lex.Location()
	value := ""

	char := lex.CurrentChar()
	for unicode.IsDigit(char) || unicode.IsLetter(char) || char == '_' {
		value += string(char)

		err := lex.Advance()
		if err != nil {
			break
		}
		char = lex.CurrentChar()
	}

	lex.Rewind()

	if token := tokens.FromIdentifier(value); token == tokens.Illegal {
		return &tokens.Token{
			Kind:     tokens.Identifier,
			Value:    value,
			Location: startLocation.Copy(),
		}
	} else {
		return &tokens.Token{
			Kind:     token,
			Value:    value,
			Location: startLocation.Copy(),
		}
	}
}
