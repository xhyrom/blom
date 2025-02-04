package ast

import "blom/tokens"

type MemberAccess struct {
	Left  Expression
	Right Expression
	Loc   tokens.Location
}

func (m *MemberAccess) Kind() NodeKind {
	return MemberAccessNode
}

func (m *MemberAccess) Location() tokens.Location {
	return m.Loc
}

func (m *MemberAccess) SetLocation(row uint64, column uint64) {
	m.Loc.Row = row
	m.Loc.Column = column
}
