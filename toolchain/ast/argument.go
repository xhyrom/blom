package ast

import (
	"fmt"
)

type Argument struct {
	Name IdentifierLiteral
	Type Type
}

func (arg Argument) String() string {
	return fmt.Sprintf("%s: %s", arg.Name.Value, arg.Type)
}
