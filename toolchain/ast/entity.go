package ast

import (
	"blom/tokens"
)

type Entity struct {
	Name        string
	Fields      []*VariableDeclarationStatement
	Annotations []Annotation
	Loc         tokens.Location
}

func (e Entity) Kind() NodeKind {
	return EntityNode
}

func (e Entity) Location() tokens.Location {
	return e.Loc
}

func (e *Entity) SetLocation(row uint64, column uint64) {
	e.Loc.Row = row
	e.Loc.Column = column
}

func (e Entity) Equal(other Type) bool {
	if otherEntity, ok := other.(*Entity); ok {
		if e.Name != otherEntity.Name {
			return false
		}

		if len(e.Fields) != len(otherEntity.Fields) {
			return false
		}

		return true
	}

	return false
}

func (e Entity) String() string {
	return e.Name
}

func (e Entity) IsPointer() bool {
	return false
}

func (e Entity) IsFunction() bool {
	return false
}

func (e Entity) IsEntity() bool {
	return true
}

func (e Entity) IsNumeric() bool {
	return false
}

func (e Entity) IsInteger() bool {
	return false
}

func (e Entity) IsFloatingPoint() bool {
	return false
}

func (e Entity) IsMapToInt() bool {
	return false
}

func (e Entity) Weight() uint8 {
	return 0
}

type EntityConstruction struct {
	Name   string
	Values map[string]Expression
	Loc    tokens.Location
}

func (e EntityConstruction) Kind() NodeKind {
	return EntityConstructionNode
}

func (e EntityConstruction) Location() tokens.Location {
	return e.Loc
}

func (e *EntityConstruction) SetLocation(row uint64, column uint64) {
	e.Loc.Row = row
	e.Loc.Column = column
}
