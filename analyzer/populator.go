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

			if statement.Name == "main" && !statement.HasAnnotation(ast.Public) {
				dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column-4)
				dbg.ThrowError(
					"The 'main' function must be declared as public since it's the program's entry point.",
					true,
					debug.NewHint(
						"Add the '@public' annotation to the function declaration.",
						" @public",
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

	if a.Functions["main"] == nil {
		dbg := debug.NewSourceLocation(a.Source, 1, 1)
		dbg.ThrowError(
			"A public 'main' function is required as the program's entry point.",
			true,
			debug.NewHint(
				"Add a public function called 'main' to your program",
				"fun @public main() {}\n",
			),
		)
	}
}
