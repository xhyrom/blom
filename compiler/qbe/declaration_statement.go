package qbe

import (
	"blom/ast"
	"blom/compiler"
	"fmt"
	"strings"
)

type Variable struct {
	declaration *ast.DeclarationStatement
	id          int
}

func (c *Compiler) CompileDeclarationStatement(stmt *ast.DeclarationStatement, indent int) string {
	indentation := strings.Repeat("    ", indent-1)

	env := c.Environment

	if stmt.Value.Kind() == ast.BinaryExpressionNode || stmt.Value.Kind() == ast.FunctionCallNode || stmt.Value.Kind() == ast.FloatLiteralNode {
		result := fmt.Sprintf("%%%s.addr.%d =l alloc8 %d\n", stmt.Name, env.TempCounter, c.AllocSize(stmt.Type))
		result += fmt.Sprintf("%s%s\n", indentation, c.CompileStatement(stmt.Value, indent))
		result += fmt.Sprintf("%sstore%s %%tmp.%d, %%%s.addr.%d\n", indentation, c.StoreType(stmt.Type), env.TempCounter, stmt.Name, env.TempCounter)
		result += fmt.Sprintf("%s%%%s.%d =%s load%s %%%s.addr.%d", indentation, stmt.Name, env.TempCounter, c.StoreType(stmt.Type), stmt.Type, stmt.Name, env.TempCounter)

		env.Set(stmt.Name, &Variable{
			declaration: stmt,
			id:          env.TempCounter,
		})

		env.TempCounter += 1

		return result
	}

	result := fmt.Sprintf("%%%s.addr.%d =l alloc8 %d\n", stmt.Name, env.TempCounter, c.AllocSize(stmt.Type))
	result += fmt.Sprintf("%sstore%s %s, %%%s.addr.%d\n", indentation, c.StoreType(stmt.Type), c.CompileStatement(stmt.Value, indent), stmt.Name, env.TempCounter)
	result += fmt.Sprintf("%s%%%s.%d =%s load%s %%%s.addr.%d", indentation, stmt.Name, env.TempCounter, c.StoreType(stmt.Type), stmt.Type, stmt.Name, env.TempCounter)

	env.Set(stmt.Name, &Variable{
		declaration: stmt,
		id:          env.TempCounter,
	})

	env.TempCounter += 1

	return result
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
