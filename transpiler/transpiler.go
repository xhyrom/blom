package transpiler

import "blom/ast"

type Transpiler interface {
	Transpile(program *ast.Program) (string, error)
}
