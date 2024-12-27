package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
	"strconv"
	"strings"
)

type Compiler struct {
	Source string
	data   []string
	env    *env.Environment
}

func New(file string) Compiler {
	return Compiler{
		Source: file,
	}
}

func (c *Compiler) Compile(program *ast.Program) (string, error) {
	result := ""

	block := c.CompileBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, env.New(), 0)

	for _, data := range c.data {
		result += data + "\n"
	}

	result += block

	return result, nil
}

func (c *Compiler) CompileBlock(block ast.BlockStatement, env *env.Environment, indent int) string {
	result := ""
	indentation := strings.Repeat("    ", indent)

	for _, stmt := range block.Body {
		compiled := c.CompileStatement(stmt, env, indent)
		if len(compiled) > 0 {
			result += indentation + compiled + "\n"
		}
	}

	return result
}

func (c *Compiler) CompileStatement(stmt ast.Statement, env *env.Environment, indent int) string {
	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		return strconv.FormatInt(int64(stmt.Value), 10)
	case *ast.StringLiteralStatement:
		c.data = append(c.data, fmt.Sprintf("data $%s.%d = { b \"%s\", b 0 }", "s", indent, stmt.Value))
		return fmt.Sprintf("$%s.%d", "s", indent)
	case *ast.DeclarationStatement:
		return c.CompileDeclarationStatement(stmt, env, indent+1)
	case *ast.IdentifierLiteralStatement:
		return "%" + stmt.Value
	case *ast.FunctionCall:
		return c.CompileFunctionCall(stmt, env, indent+1)
	case *ast.FunctionDeclaration:
		return c.CompileFunctionDeclaration(stmt, env, indent+1)
	case *ast.ReturnStatement:
		return c.CompileReturnStatement(stmt, env, indent+1)
	case *ast.BlockStatement:
		return c.CompileBlock(*stmt, env, indent+1)
	case *ast.BinaryExpression:
		return c.CompileBinaryExpression(stmt, env, indent+1)
	}

	fmt.Printf("Unknown statement: %T\n", stmt)

	return ""
}
