package consumers

import "blom/tokens"

func ConsumeString(lex Lexer) *tokens.Token {
	lex.Advance()

	startLocation := lex.Location()
	value := string(lex.CurrentChar())

	lex.Advance()

	for !lex.IsEof() {
		if lex.CurrentChar() == '"' && lex.PreviousChar() != '\\' {
			break
		}

		value += string(lex.CurrentChar())
		lex.Advance()
	}

	return &tokens.Token{
		Kind:     tokens.StringLiteral,
		Value:    string(value),
		Location: startLocation.Copy(),
	}
}
