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
			a.populateFunction(statement)
		case *ast.TypeDefinition:
			a.populateTypeDefinition(statement)
		default:
			dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
			dbg.ThrowError(
				"Top-level statements must be function declarations",
				true,
			)
		}
	}

	if !a.FunctionManager.Has("main") {
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

func (a *Analyzer) populateFunction(statement *ast.FunctionDeclaration) {
	fun, exists := a.FunctionManager.GetByDeclaration(statement)

	if exists {
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

	a.FunctionManager.Register(statement)
}

func (a *Analyzer) populateTypeDefinition(statement *ast.TypeDefinition) {
	if statement.Name == "main" {
		dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
		dbg.ThrowError(
			"Type 'main' is reserved and cannot be redefined.",
			true,
			debug.NewHint(
				"Consider renaming the type to avoid conflicts.",
				"",
			),
		)
	}

	_, err := ast.ParseType(statement.Name, map[string]ast.Type{})
	if err == nil {
		dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
		dbg.ThrowError(
			fmt.Sprintf("Type '%s' is a primitive type and cannot be redefined.", statement.Name),
			true,
			debug.NewHint(
				"Consider renaming the type to avoid conflicts.",
				"",
			),
		)
	}

	if a.TypeManager.Has(statement.Name) {
		dbg := debug.NewSourceLocation(a.Source, statement.Location().Row, statement.Location().Column)
		dbg.ThrowError(
			fmt.Sprintf("Type '%s' has already been defined.", statement.Name),
			true,
			debug.NewHint(
				"Consider renaming the type to avoid conflicts.",
				"",
			),
		)
	}

	a.TypeManager.Register(statement.Name, statement.Type)
}
