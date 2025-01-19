package ast

import "blom/tokens"

type NodeKind int

const (
	ProgramNode NodeKind = iota
	IdentifierLiteralNode
	CharLiteralNode
	StringLiteralNode
	IntLiteralNode
	FloatLiteralNode
	BooleanLiteralNode
	BinaryExpressionNode
	UnaryExpressionNode
	VariableDeclarationNode
	AssignmentNode
	ReturnNode
	BlockNode
	IfNode
	WhileLoopNode

	AnnotationNode
	FunctionDeclarationNode
	FunctionCallNode
	CompileTimeFunctionCallNode
	TypeDefinitionNode
)

type Statement interface {
	Kind() NodeKind
	Location() tokens.Location
	SetLocation(row uint64, column uint64)
}

type Program BlockStatement
