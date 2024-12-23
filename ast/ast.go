package ast

import "blom/tokens"

type NodeKind int

const (
	ProgramNode NodeKind = iota
	IdentifierLiteralNode
	IntLiteralNode
	FloatLiteralNode
	BinaryExpressionNode
	UnaryExpressionNode
	DeclarationNode
	ReturnNode

	AnnotationNode
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
