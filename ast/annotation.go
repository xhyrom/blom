package ast

import "blom/tokens"

type Annotation struct {
	Name string
	Loc  tokens.Location
}

func (a Annotation) Kind() NodeKind {
	return AnnotationNode
}

func (a Annotation) Location() tokens.Location {
	return a.Loc
}
