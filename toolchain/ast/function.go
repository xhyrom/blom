package ast

import (
	"blom/tokens"
	"fmt"
	"strings"
)

type FunctionDeclaration struct {
	Name        tokens.Token
	Params      []Argument
	Annotations []Annotation
	Block       *Block

	Return Type

	Loc tokens.Location
}

func (f FunctionDeclaration) Kind() NodeKind {
	return FunctionDeclarationNode
}

func (f FunctionDeclaration) Location() tokens.Location {
	return f.Loc
}

func (f FunctionDeclaration) HasAnnotation(ty AnnotationType) bool {
	for _, annotation := range f.Annotations {
		if annotation.Type == ty {
			return true
		}
	}

	return false
}

func (f FunctionDeclaration) String() string {
	args := make([]string, len(f.Params))
	for i, arg := range f.Params {
		args[i] = arg.String()
	}
	return fmt.Sprintf("fn(%s) -> %s", strings.Join(args, ", "), f.Return)
}
