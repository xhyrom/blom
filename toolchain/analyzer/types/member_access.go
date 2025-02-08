package types

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func (a *TypeAnalyzer) analyzeMemberAccess(access *ast.MemberAccess) ast.Type {
	left := a.analyzeExpression(access.Left)
	right := access.Right.(*ast.IdentifierLiteral).Value

	if !left.IsEntity() {
		dbg := debug.NewSourceLocation(a.Source, access.Location().Row, access.Location().Column)
		dbg.ThrowError(
			fmt.Sprintf(
				"Cannot access member '%s' on non-entity type '%s'.",
				right,
				left,
			),
			true,
			debug.NewHint(
				"Ensure that the left-hand side of the member access is an entity type.",
				"",
			),
		)
	}

	var field *ast.VariableDeclarationStatement
	for _, entityField := range left.(*ast.Entity).Fields {
		if entityField.Name == right {
			field = entityField
			break
		}
	}

	if field == nil {
		dbg := debug.NewSourceLocation(a.Source, access.Location().Row, access.Location().Column)
		dbg.ThrowError(
			fmt.Sprintf(
				"Entity type '%s' does not have a member '%s'.",
				left,
				right,
			),
			true,
		)
	}

	return field.Type
}
