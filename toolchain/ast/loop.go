package ast

import "blom/tokens"

type WhileLoop struct {
	Condition Node
	Block     Block
	Loc       tokens.Location
}

func (w WhileLoop) Kind() NodeKind {
	return WhileLoopNode
}

func (w WhileLoop) Location() tokens.Location {
	return w.Loc
}
