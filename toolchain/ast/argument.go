package ast

import "blom/tokens"

type Argument struct {
	Name tokens.Token
	Type Type
}
