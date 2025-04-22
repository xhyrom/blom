package consumers

import (
	"blom/tokens"
	"unicode"
)

func ConsumeNumber(lex Lexer) *tokens.Token {
	startLocation := lex.Location()
	value := ""
	isFloat := false

	if lex.CurrentChar() == '0' {
		value += string(lex.CurrentChar())
		lex.Advance()

		if lex.IsEof() {
			return &tokens.Token{
				Kind:     tokens.IntLiteral,
				Value:    value,
				Location: startLocation.Copy(),
			}
		}

		switch lex.CurrentChar() {
		case 'x', 'X': // hexadecimal
			value += string(lex.CurrentChar())
			lex.Advance()

			for !lex.IsEof() && (unicode.IsDigit(lex.CurrentChar()) ||
				('a' <= lex.CurrentChar() && lex.CurrentChar() <= 'f') ||
				('A' <= lex.CurrentChar() && lex.CurrentChar() <= 'F') ||
				lex.CurrentChar() == '_') {
				if lex.CurrentChar() != '_' {
					value += string(lex.CurrentChar())
				}
				lex.Advance()
			}

			lex.Rewind()
			return &tokens.Token{
				Kind:     tokens.IntLiteral,
				Value:    value,
				Location: startLocation.Copy(),
			}

		case 'b', 'B': // binary
			value += string(lex.CurrentChar())
			lex.Advance()

			for !lex.IsEof() && (lex.CurrentChar() == '0' || lex.CurrentChar() == '1' || lex.CurrentChar() == '_') {
				if lex.CurrentChar() != '_' {
					value += string(lex.CurrentChar())
				}
				lex.Advance()
			}

			lex.Rewind()
			return &tokens.Token{
				Kind:     tokens.IntLiteral,
				Value:    value,
				Location: startLocation.Copy(),
			}

		case 'o', 'O': // octal
			value += string(lex.CurrentChar())
			lex.Advance()

			for !lex.IsEof() && ('0' <= lex.CurrentChar() && lex.CurrentChar() <= '7' || lex.CurrentChar() == '_') {
				if lex.CurrentChar() != '_' {
					value += string(lex.CurrentChar())
				}
				lex.Advance()
			}

			lex.Rewind()
			return &tokens.Token{
				Kind:     tokens.IntLiteral,
				Value:    value,
				Location: startLocation.Copy(),
			}

		default:
			if !unicode.IsDigit(lex.CurrentChar()) {
				lex.Rewind()
				return &tokens.Token{
					Kind:     tokens.IntLiteral,
					Value:    value,
					Location: startLocation.Copy(),
				}
			}
		}
	}

	for !lex.IsEof() && (unicode.IsDigit(lex.CurrentChar()) || lex.CurrentChar() == '.' || lex.CurrentChar() == '_') {
		if lex.CurrentChar() == '.' {
			if isFloat {
				break
			}

			isFloat = true
		}

		if lex.CurrentChar() != '_' {
			value += string(lex.CurrentChar())
		}

		lex.Advance()
	}

	// handle scientific notation
	if !lex.IsEof() && (lex.CurrentChar() == 'e' || lex.CurrentChar() == 'E') {
		value += string(lex.CurrentChar())
		lex.Advance()

		if !lex.IsEof() && (lex.CurrentChar() == '+' || lex.CurrentChar() == '-') {
			value += string(lex.CurrentChar())
			lex.Advance()
		}

		hasExponentDigits := false

		for !lex.IsEof() && (unicode.IsDigit(lex.CurrentChar()) || lex.CurrentChar() == '_') {
			if lex.CurrentChar() != '_' {
				value += string(lex.CurrentChar())
				hasExponentDigits = true
			}
			lex.Advance()
		}

		if !hasExponentDigits {
			for value[len(value)-1] == 'e' || value[len(value)-1] == 'E' ||
				value[len(value)-1] == '+' || value[len(value)-1] == '-' {
				value = value[:len(value)-1]
				lex.Rewind()
			}
		}
	} else {
		lex.Rewind()
	}

	// handle trailing dot
	if len(value) > 0 && value[len(value)-1] == '.' {
		value = value[:len(value)-1]
		isFloat = false
		lex.Rewind()
	}

	kind := tokens.IntLiteral
	if isFloat {
		kind = tokens.FloatLiteral
	}

	loc := startLocation.Copy()
	if lex.CurrentChar() == '\n' {
		loc.Column-- // remove the newline character from the column
	}

	return &tokens.Token{
		Kind:     kind,
		Value:    value,
		Location: loc,
	}
}
