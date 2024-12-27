package ast

import (
	"blom/tokens"
	"slices"
)

type Annotation struct {
	Type AnnotationType
	Loc  tokens.Location
}

func (a Annotation) Kind() NodeKind {
	return AnnotationNode
}

func (a Annotation) Location() tokens.Location {
	return a.Loc
}

type AnnotationType int

const (
	Native AnnotationType = iota
	Unknown
)

var annotations = []string{
	Native: "native",
}

func (a AnnotationType) String() string {
	return annotations[a]
}

func ParseAnnotation(annotation string) AnnotationType {
	index := slices.Index(annotations, annotation)
	if index == -1 {
		return -1
	}

	return AnnotationType(index)
}
