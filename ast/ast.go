package ast

import "blom/tokens"

type NodeKind int

const (
	ProgramNode NodeKind = iota
	IntLiteralNode
	FloatLiteralNode
	BinaryExpressionNode
	DeclarationNode
	ReturnNode

	FunctionDeclarationNode
	FunctionCallNode
)

type Statement interface {
	Kind() NodeKind
	Location() tokens.Location
}

type Program struct {
	Body []Statement
}
