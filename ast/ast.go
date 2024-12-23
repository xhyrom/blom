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
	BlockNode
	IfNode

	AnnotationNode
	FunctionDeclarationNode
	FunctionCallNode
)

type Statement interface {
	Kind() NodeKind
	Location() tokens.Location
}

type Program BlockStatement
