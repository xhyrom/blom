package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/env"
	"fmt"
	"strings"
)

func (c *Compiler) CompileVariableDeclarationStatement(stmt *ast.VariableDeclarationStatement, scope *env.Environment[*Variable], id int) ([]string, *QbeIdentifier) {
	result := make([]string, 0)

	result = append(result, fmt.Sprintf("%%%s.addr.%d =l alloc8 %d", stmt.Name, id, c.AllocSize(stmt.Type)))

	stat, statIdentifier := c.CompileStatement(stmt.Value, scope)

	if stmt.Value.Kind() == ast.IfNode {
		for _, stat := range stat {
			stat = strings.TrimSpace(stat)

			if strings.HasPrefix(stat, "ret") {
				returnValue := strings.Split(stat, " ")[1]
				stat = fmt.Sprintf("store%s %s, %%%s.addr.%d", c.StoreType(stmt.Type), returnValue, stmt.Name, id)
			}

			result = append(result, stat)

			if strings.HasPrefix(stat, "ret") {
				result = append(result, fmt.Sprintf("jnz @end.%d", statIdentifier.Id))
			}
		}
	} else {
		for _, stat := range stat {
			result = append(result, stat)
		}

		result = append(result, fmt.Sprintf("store%s %s, %%%s.addr.%d", c.StoreType(stmt.Type), c.StoreVal(statIdentifier), stmt.Name, id))
	}

	name := fmt.Sprintf("%%%s.%d", stmt.Name, id)
	result = append(result, fmt.Sprintf("%s =%s load%s %%%s.addr.%d", name, c.StoreType(stmt.Type), stmt.Type, stmt.Name, id))

	scope.Set(stmt.Name, &Variable{
		Type: stmt.Type,
		Id:   id,
	})

	result = append(result, "# ^ declaration statement\n")

	return result, &QbeIdentifier{
		Name: name,
		Type: stmt.Type,
	}
}

func (c *Compiler) CompileAssignmentStatement(stmt *ast.AssignmentStatement, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	original, exists := scope.FindVariable(stmt.Name)
	if !exists {
		panic("VARIABLE NOT FOUND IN COMPILER IN ASSIGNMENT STATEMENT, SHOULD HAPPEN IN ANALYZER")
	}

	return c.CompileVariableDeclarationStatement(
		&ast.VariableDeclarationStatement{
			Name:  stmt.Name,
			Type:  original.Type,
			Value: stmt.Value,
			Loc:   stmt.Loc,
		},
		scope,
		original.Id,
	)
}

func (c *Compiler) StoreVal(additional *QbeIdentifier) string {
	if additional.Raw {
		switch additional.Type {
		case compiler.Double, compiler.Single:
			return fmt.Sprintf("d_%s", additional.Name)
		}
	}

	return additional.Name
}

func (c *Compiler) AllocSize(t compiler.Type) int {
	switch t {
	case compiler.Long, compiler.UnsignedLong, compiler.Double:
		return 8
	case compiler.Word, compiler.UnsignedWord, compiler.Single:
		return 4
	case compiler.Halfword, compiler.UnsignedHalfword:
		return 2
	default:
		return 1
	}
}

func (c *Compiler) StoreType(t compiler.Type) compiler.Type {
	switch t {
	case compiler.Long, compiler.UnsignedLong, compiler.String:
		return compiler.Long
	case compiler.Double:
		return compiler.Double
	case compiler.Single:
		return compiler.Single
	default:
		return compiler.Word
	}
}
