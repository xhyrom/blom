package overloading

import (
	"blom/analyzer/manager"
	"blom/ast"
)

type Overloader struct {
	Program         *ast.Program
	FunctionManager *manager.FunctionManager
}

func New(program *ast.Program, functionManager *manager.FunctionManager) *Overloader {
	return &Overloader{
		Program:         program,
		FunctionManager: functionManager,
	}
}

func (o *Overloader) Overload() {
	for _, statement := range o.Program.Body {
		switch statement := statement.(type) {
		case *ast.FunctionDeclaration:
			o.overloadFunctionDeclaration(statement)
		case *ast.FunctionCall:
			o.overloadFunctionCall(statement)
		}
	}
}
