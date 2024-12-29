package analyzer

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *Analyzer) populate() {
	for _, statement := range a.Program.Body {
		switch statement := statement.(type) {
		case *ast.FunctionDeclaration:
			fun := a.Functions[statement.Name]
			if fun != nil {
				dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
				dbg.ThrowError(
					fmt.Sprintf(
						"Function '%s' has already been declared at line %d",
						statement.Name,
						fun.Location().Row,
					),
					true,
					debug.NewHint(
						"Consider renaming the function to avoid conflicts.",
						"",
					),
				)
			}

			a.Functions[statement.Name] = statement
		default:
			dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
			dbg.ThrowError(
				"Top-level statements must be function declarations",
				true,
			)
		}
	}
}
