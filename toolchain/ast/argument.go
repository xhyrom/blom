package ast

import (
	"blom/tokens"
	"fmt"
)

type Argument struct {
	Name IdentifierLiteral
	Type Type
	Loc  tokens.Location
}

func (arg Argument) Location() tokens.Location {
	return arg.Loc
}

func (arg Argument) String() string {
	return fmt.Sprintf("%s: %s", arg.Name.Value, arg.Type)
}
