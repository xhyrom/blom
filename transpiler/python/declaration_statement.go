package python

import (
	"blom/ast"
	"blom/env"
	"fmt"
)

// TODO: fix for more nested environments
func (t PythonTranspiler) TranspileDeclarationStatement(declaration *ast.DeclarationStatement, environment *env.Environment, indent int) string {
	skip := false

	if environment.Parent != nil && environment.Parent.FindVariable(declaration.Name) != nil {
		if !declaration.Redeclaration && environment.Get(declaration.Name) == nil {
			environment.Set(declaration.Name, &env.BooleanObject{})

			return fmt.Sprintf("%s_%d = %s\n", declaration.Name, indent, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
		}
	} else {
		skip = true

		environment.Set(declaration.Name, &env.BooleanObject{
			Value: false,
		})
	}

	if environment.Get(declaration.Name) != nil && !skip {
		return fmt.Sprintf("%s_%d = %s\n", declaration.Name, indent, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
	}

	return fmt.Sprintf("%s = %s\n", declaration.Name, t.TranspileAndFunctionifyStatement(declaration.Value, environment, indent))
}
