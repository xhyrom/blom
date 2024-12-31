package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretFunctionDeclaration(interpreter Interpreter, environment *env.Scope[objects.Object], declaration *ast.FunctionDeclaration) {
	environment.SetFunction(declaration.Name, declaration)
}
