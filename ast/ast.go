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
)

var names = []string{
	ProgramNode:                 "program",
	IdentifierLiteralNode:       "identifier literal",
	CharLiteralNode:             "char literal",
	StringLiteralNode:           "string literal",
	IntLiteralNode:              "int literal",
	FloatLiteralNode:            "float literal",
	BinaryExpressionNode:        "binary expression",
	UnaryExpressionNode:         "unary expression",
	VariableDeclarationNode:     "variable declaration",
	AssignmentNode:              "assignment",
	ReturnNode:                  "return",
	BlockNode:                   "block",
	IfNode:                      "if",
	WhileLoopNode:               "while loop",
	AnnotationNode:              "annotation",
	FunctionDeclarationNode:     "function declaration",
	FunctionCallNode:            "function call",
	CompileTimeFunctionCallNode: "compile time function call",
}

type Statement interface {
	Kind() NodeKind
	Location() tokens.Location
	SetLocation(row uint64, column uint64)
}

type Program BlockStatement

func (k NodeKind) String() string {
	return names[k]
}
