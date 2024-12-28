package ast

import "blom/tokens"

type IfStatement struct {
	Condition Expression
	Then      *BlockStatement
	Else      *BlockStatement
	Loc       tokens.Location
}

func (i IfStatement) Kind() NodeKind {
	return IfNode
}

func (i IfStatement) Location() tokens.Location {
	return i.Loc
}

func (i *IfStatement) SetLocation(row uint64, column uint64) {
	i.Loc.Row = row
	i.Loc.Column = column
}

func (i *IfStatement) HasElse() bool {
	if i.Else != nil {
		return true
	}

	return false
}
