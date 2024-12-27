package ast

import "blom/tokens"

type ForLoopStatement struct {
	Declaration *DeclarationStatement
	Condition   *BinaryExpression
	Step        *DeclarationStatement
	Body        *BlockStatement
	Loc         tokens.Location
}

func (f ForLoopStatement) Kind() NodeKind {
	return ForLoopNode
}

func (f ForLoopStatement) Location() tokens.Location {
	return f.Loc
}

func (f *ForLoopStatement) SetLocation(row uint64, column uint64) {
	f.Loc.Row = row
	f.Loc.Column = column
}

type WhileLoopStatement struct {
	Condition Expression
	Body      *BlockStatement
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
