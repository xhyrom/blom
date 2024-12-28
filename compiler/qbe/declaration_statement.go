package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/debug"
	"fmt"
)

type Variable struct {
	declaration *ast.DeclarationStatement
	id          int
}

func (c *Compiler) CompileDeclarationStatement(stmt *ast.DeclarationStatement, indent int) ([]string, *Additional) {
	env := c.Environment

	result := make([]string, 0)

	id := env.TempCounter

	result = append(result, fmt.Sprintf("%%%s.addr.%d =l alloc8 %d", stmt.Name, id, c.AllocSize(stmt.Type)))

	stat, statIdentifier := c.CompileStatement(stmt.Value, indent, &stmt.Type)

	if stmt.Type != statIdentifier.Type {
		dbg := debug.NewSourceLocation(c.Source, stmt.Loc.Row, stmt.Loc.Column)
		dbg.ThrowError(fmt.Sprintf("Type mismatch in declaration %s != %s", stmt.Type.Inspect(), statIdentifier.Type.Inspect()), true)
	}

	for _, stat := range stat {
		result = append(result, stat)
	}

	result = append(result, fmt.Sprintf("store%s %s, %%%s.addr.%d", c.StoreType(stmt.Type), statIdentifier.Name, stmt.Name, id))

	name := fmt.Sprintf("%%%s.%d", stmt.Name, id)
	result = append(result, fmt.Sprintf("%s =%s load%s %%%s.addr.%d", name, c.StoreType(stmt.Type), stmt.Type, stmt.Name, id))

	env.Set(stmt.Name, &Variable{
		declaration: stmt,
		id:          id,
	})

	result = append(result, "# ^ declaration statement\n")

	return result, &Additional{
		Name: name,
		Type: stmt.Type,
	}
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
