package codegen

import (
	"blom/ast"
	"blom/compiler/codegen/qbe"
)

type Backend int

const (
	QBE Backend = iota
	LLVM
)

type Codegen struct {
	Backend Backend
}

func New(backend Backend) *Codegen {
	return &Codegen{Backend: backend}
}

func (c Codegen) Generate(program *ast.Program) string {
	switch c.Backend {
	case QBE:
		qbe := qbe.New()
		qbe.Generate(program)

		return qbe.Emit()
	}

	panic("unsupported codegen backend")
}
