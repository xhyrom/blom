package lexer

import (
	"blom/lexer/consumers"
	"blom/reader"
	"blom/tokens"
	"unicode"
)

type Lexer struct {
	Reader   *reader.Reader
	location *tokens.Location
}

func New(file string, content string) *Lexer {
	return &Lexer{
		Reader: reader.New(content),
		location: &tokens.Location{
			File: file,
			Row:  1,
			Col:  0,
		},
	}
}

func (lex *Lexer) Next() *tokens.Token {
	char, err := lex.Reader.Read()
	if err != nil {
		return nil
	}

	lex.location.Col++

	kind := tokens.Illegal

	switch char {
	case '\n':
		{
			lex.NewLine()
			return lex.Next()
		}
	case '=':
		{
			lex.Advance()

			switch lex.CurrentChar() {
			case '=':
				kind = tokens.Equals
			default:
				kind = tokens.Assign
				lex.Rewind()
			}
		}
	case '+':
		kind = tokens.Plus
	case '-':
		kind = tokens.Minus
	case '*':
		kind = tokens.Asterisk
	case '/':
		{
			lex.Advance()

			switch lex.CurrentChar() {
			case '/':
				consumers.ConsumeComment(lex)
				return lex.Next()
			default:
				kind = tokens.Slash
				lex.Rewind()
			}
		}
	case '%':
		kind = tokens.Modulo
	case '<':
		{
			lex.Advance()

			switch lex.CurrentChar() {
			case '=':
				kind = tokens.LessThanOrEqual
			default:
				kind = tokens.LessThan
				lex.Rewind()
			}
		}
	case '>':
		{
			lex.Advance()

			switch lex.CurrentChar() {
			case '=':
				kind = tokens.GreaterThanOrEqual
			default:
				kind = tokens.GreaterThan
				lex.Rewind()
			}
		}
	case '.':
		kind = tokens.Dot
	case ',':
		kind = tokens.Comma
	case ';':
		kind = tokens.Semicolon
	case '@':
		kind = tokens.AtMark
	case '(':
		kind = tokens.LeftParenthesis
	case ')':
		kind = tokens.RightParenthesis
	case '[':
		kind = tokens.LeftSquareBracket
	case ']':
		kind = tokens.RightSquareBracket
	case '{':
		kind = tokens.LeftCurlyBracket
	case '}':
		kind = tokens.RightCurlyBracket
	default:
		if unicode.IsSpace(char) {
			return lex.Next()
		} else if unicode.IsDigit(char) {
			return consumers.ConsumeNumber(lex)
		} else if unicode.IsLetter(char) {
			return consumers.ConsumeIdentifier(lex)
		}

		kind = tokens.Illegal
	}

	if kind == tokens.Illegal {
		return nil
	}

	return &tokens.Token{
		Kind:     kind,
		Location: lex.location.Copy(),
	}
}

func (lex *Lexer) NewLine() {
	lex.location.Row++
	lex.location.Col = 0
}

func (lex *Lexer) Advance() error {
	_, err := lex.Reader.Peek()
	if err != nil {
		return err
	}

	lex.location.Col++
	lex.Reader.Read()

	return nil
}

func (lex *Lexer) Rewind() {
	lex.Reader.Rewind()
	lex.location.Col--
}

func (lex *Lexer) CurrentChar() rune {
	char, err := lex.Reader.Current()

	if err != nil {
		panic(err)
	}

	return char
}

func (lex *Lexer) Location() *tokens.Location {
	return lex.location
}
