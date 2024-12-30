package compiler

import (
	"blom/ast"
	"blom/compiler/qbe"
)

type Backend int

const (
	QBE Backend = iota
	LLVM
)

type Compiler struct {
	Backend Backend
}

func New(backend Backend) *Compiler {
	return &Compiler{Backend: backend}
}

func (c *Compiler) Compile(program *ast.Program) string {
	switch c.Backend {
	case QBE:
		return c.compileQBE(program)
	case LLVM:
		return c.compileLLVM()
	}

	panic("Unknown backend")
}

func (c *Compiler) compileQBE(program *ast.Program) string {
	qbeCompiler := qbe.New()
	qbeCompiler.Compile(program)

	return qbeCompiler.Emit()
}

func (c *Compiler) compileLLVM() string {
	panic("Not implemented")
}
