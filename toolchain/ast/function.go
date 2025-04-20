package ast

import (
	"blom/tokens"
)

type FunctionDeclaration struct {
	Id          tokens.Token
	Params      []Argument
	Annotations []Annotation
	Block       *Block

	ReturnType Type

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

type FunctionCall struct {
	Path Path
	Args []Node
	Loc  tokens.Location
}

func (f FunctionCall) Kind() NodeKind {
	return FunctionCallNode
}

func (f FunctionCall) Location() tokens.Location {
	return f.Loc
}
