package types

import (
	"blom/ast"
)

type FunctionCallMap map[*ast.FunctionCall]*ast.FunctionDeclaration

type AnalysisContext struct {
	FunctionCalls FunctionCallMap
}

func NewAnalysisContext() *AnalysisContext {
	return &AnalysisContext{
		FunctionCalls: make(FunctionCallMap),
	}
}
