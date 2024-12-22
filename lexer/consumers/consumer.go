package consumers

import "blom/tokens"

type Lexer interface {
	CurrentChar() rune
	Advance() error
	Location() *tokens.Location
}
