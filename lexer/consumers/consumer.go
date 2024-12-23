package consumers

import "blom/tokens"

type Lexer interface {
	CurrentChar() rune
	PreviousChar() rune
	Advance() error
	Location() *tokens.Location
	NewLine()
	Rewind()
	IsEof() bool
}
