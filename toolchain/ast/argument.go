package ast

import (
	"blom/tokens"
	"fmt"
)

type Argument struct {
	Name tokens.Token
	Type Type
}

func (arg Argument) String() string {
	return fmt.Sprintf("%s: %s", arg.Name, arg.Type)
}
