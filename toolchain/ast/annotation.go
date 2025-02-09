package ast

import (
	"blom/tokens"
	"slices"
)

type Annotation struct {
	Type   AnnotationType
	Values map[string]Expression
	Loc    tokens.Location
}

func GetAnnotationValueOrDefault[T any](a *Annotation, key string, def T) T {
	if a == nil {
		return def
	}

	if val, ok := a.Values[key]; ok {
		switch v := val.(type) {
		case *StringLiteral:
			if _, ok := any(def).(string); ok {
				return any(v.Value).(T)
			}
		case *IntLiteral:
			if _, ok := any(def).(int64); ok {
				return any(v.Value).(T)
			}
		case *FloatLiteral:
			if _, ok := any(def).(float64); ok {
				return any(v.Value).(T)
			}
		case *BooleanLiteral:
			if _, ok := any(def).(bool); ok {
				return any(v.Value).(T)
			}
		case *IdentifierLiteral:
			if _, ok := any(def).(string); ok {
				return any(v.Value).(T)
			}
		}
	}

	return def
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
	Public
	Infix
)

var annotations = []string{
	Native: "native",
	Public: "public",
	Infix:  "infix",
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
