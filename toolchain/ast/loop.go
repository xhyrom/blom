package ast

import "blom/tokens"

type WhileLoop struct {
	Condition Expression
	Body      []Statement
	Cat       Category
	Loc       tokens.Location
}

func (w WhileLoop) Kind() NodeKind {
	return WhileLoopNode
}

func (w WhileLoop) Location() tokens.Location {
	return w.Loc
}

func (w WhileLoop) Category() Category {
	return w.Cat
}

func (w *WhileLoop) SetLocation(row uint64, column uint64) {
	w.Loc.Row = row
	w.Loc.Column = column
}
