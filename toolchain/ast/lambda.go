package ast

import (
	"blom/tokens"
)

type LambdaDeclaration struct {
	Arguments  []FunctionArgument
	ReturnType Type
	Body       []Statement
	Variadic   bool
	Loc        tokens.Location
}

func (f LambdaDeclaration) Kind() NodeKind {
	return FunctionDeclarationNode
}

func (f LambdaDeclaration) Location() tokens.Location {
	return f.Loc
}

func (f *LambdaDeclaration) SetLocation(row uint64, column uint64) {
	f.Loc.Row = row
	f.Loc.Column = column
}
