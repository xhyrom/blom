package consumers

import "blom/tokens"

func ConsumeCharacter(lex Lexer) *tokens.Token {
	lex.Advance()

	startLocation := lex.Location()
	value := lex.CurrentChar()

	lex.Advance()

	// check if there's no ' next to the character
	if value == '\\' {
		switch lex.CurrentChar() {
		case 'a':
			value = '\x07'
		case 'b':
			value = '\x08'
		case 'f':
			value = '\x0C'
		case 'n':
			value = '\x0A'
		case 'r':
			value = '\x0D'
		case 't':
			value = '\x09'
		case 'v':
			value = '\x0B'
		case '0':
			value = '\x00'
		case '\'':
			value = '\''
		default:
			panic("Invalid escape sequence")
		}

		lex.Advance()
	}

	if lex.CurrentChar() != '\'' {
		panic("Character literal must contain only one character")
	}

	return &tokens.Token{
		Kind:     tokens.CharLiteral,
		Value:    string(value),
		Location: startLocation.Copy(),
	}
}
