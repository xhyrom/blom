package ast

import "blom/tokens"

type WhileLoopStatement struct {
	Condition Expression
	Body      []Statement
	Loc       tokens.Location
}

func (w WhileLoopStatement) Kind() NodeKind {
	return WhileLoopNode
}

func (w WhileLoopStatement) Location() tokens.Location {
	return w.Loc
}

func (w *WhileLoopStatement) SetLocation(row uint64, column uint64) {
	w.Loc.Row = row
	w.Loc.Column = column
}
