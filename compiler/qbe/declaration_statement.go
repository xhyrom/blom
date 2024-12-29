package qbe

import (
	"blom/ast"
	"blom/compiler"
	"fmt"
	"strings"
)

type Variable struct {
	Type compiler.Type
	Id   int
}

func (c *Compiler) CompileDeclarationStatement(stmt *ast.VariableDeclarationStatement) ([]string, *Additional) {
	env := c.Environment

	result := make([]string, 0)

	id := env.TempCounter

	var stmtType compiler.Type
	if stmt.Redeclaration {
		original := env.Get(stmt.Name)
		stmtType = original.Type
	} else {
		stmtType = stmt.Type
	}

	result = append(result, fmt.Sprintf("%%%s.addr.%d =l alloc8 %d", stmt.Name, id, c.AllocSize(stmtType)))

	stat, statIdentifier := c.CompileStatement(stmt.Value, &stmt.Type)

	if stmt.Redeclaration {
		original := env.Get(stmt.Name)

		//if original.Type != statIdentifier.Type {
		//	dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		//	dbg.ThrowError(fmt.Sprintf("Type mismatch in declaration %s != %s", original.Type.Inspect(), statIdentifier.Type.Inspect()), true)
		//}

		id = original.Id // reuse the same id
	} else if stmt.Value.Kind() != ast.IfNode && stmtType != statIdentifier.Type {
		//dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		//dbg.ThrowError(fmt.Sprintf("Type mismatch in declaration %s != %s", stmtType.Inspect(), statIdentifier.Type.Inspect()), true)
	}

	if stmt.Value.Kind() == ast.IfNode {
		for _, stat := range stat {
			stat = strings.TrimSpace(stat)

			if strings.HasPrefix(stat, "ret") {
				returnValue := strings.Split(stat, " ")[1]
				stat = fmt.Sprintf("store%s %s, %%%s.addr.%d", c.StoreType(stmtType), returnValue, stmt.Name, id)
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

		result = append(result, fmt.Sprintf("store%s %s, %%%s.addr.%d", c.StoreType(stmtType), c.StoreVal(statIdentifier), stmt.Name, id))
	}

	name := fmt.Sprintf("%%%s.%d", stmt.Name, id)
	result = append(result, fmt.Sprintf("%s =%s load%s %%%s.addr.%d", name, c.StoreType(stmtType), stmtType, stmt.Name, id))

	env.Set(stmt.Name, &Variable{
		Type: stmtType,
		Id:   id,
	})

	result = append(result, "# ^ declaration statement\n")

	return result, &Additional{
		Name: name,
		Type: stmtType,
	}
}

func (c *Compiler) StoreVal(additional *Additional) string {
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
