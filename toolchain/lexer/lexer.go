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
			Row:    1,
			Column: 0,
		},
	}
}

func (lex *Lexer) Next() *tokens.Token {
	char, err := lex.Reader.Read()
	if err != nil {
		return &tokens.Token{
			Kind: tokens.Eof,
			Location: tokens.Location{
				Row:    lex.Location().Row,
				Column: lex.Location().Column + 1,
			},
		}
	}

	lex.location.Column++

	kind := tokens.Illegal

	switch char {
	case '\n':
		{
			lex.NewLine()
			return lex.Next()
		}
	case '!':
		{
			lex.Advance()

			switch lex.CurrentChar() {
			case '=':
				kind = tokens.NotEquals
			default:
				// TODO: implement
				lex.Rewind()
			}
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
		kind = tokens.PercentSign
	case '&':
		kind = tokens.Ampersand
	case '|':
		kind = tokens.VerticalLine
	case '^':
		kind = tokens.CircumflexAccent
	case '~':
		kind = tokens.Tilde
	case '<':
		{
			lex.Advance()

			switch lex.CurrentChar() {
			case '=':
				kind = tokens.LessThanOrEqual
			case '<':
				kind = tokens.DoubleLessThan
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
			case '>':
				kind = tokens.DoubleGreaterThan
			default:
				kind = tokens.GreaterThan
				lex.Rewind()
			}
		}
	case '.':
		lex.Advance()

		switch lex.CurrentChar() {
		case '.':
			lex.Advance()

			switch lex.CurrentChar() {
			case '.':
				kind = tokens.Ellipsis
			default:
				lex.Rewind()
			}
		default:
			kind = tokens.Dot
			lex.Rewind()
		}
	case ',':
		kind = tokens.Comma
	case ':':
		kind = tokens.Colon
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
	case '\'':
		return consumers.ConsumeCharacter(lex)
	case '"':
		return consumers.ConsumeString(lex)
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
	lex.location.Column = 0
}

func (lex *Lexer) Advance() error {
	_, err := lex.Reader.Peek()
	if err != nil {
		return err
	}

	lex.location.Column++
	lex.Reader.Read()

	return nil
}

func (lex *Lexer) Rewind() {
	lex.Reader.Rewind()
	lex.location.Column--
}

func (lex *Lexer) CurrentChar() rune {
	char, err := lex.Reader.Current()

	if err != nil {
		panic(err)
	}

	return char
}

func (lex *Lexer) PreviousChar() rune {
	char, err := lex.Reader.Previous()

	if err != nil {
		panic(err)
	}

	return char
}

func (lex *Lexer) IsEof() bool {
	_, err := lex.Reader.Peek()

	return err != nil
}

func (lex *Lexer) Location() *tokens.Location {
	return lex.location
}
