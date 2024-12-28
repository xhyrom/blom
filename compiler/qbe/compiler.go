package qbe

import (
	"blom/ast"
	"blom/env"
	"fmt"
	"strconv"
	"strings"
)

type Compiler struct {
	Source      string
	data        []string
	dataCounter int
	Environment *env.Environment[*Variable]
}

func New(file string) Compiler {
	return Compiler{
		Source:      file,
		Environment: env.New[*Variable](),
	}
}

func (c *Compiler) Compile(program *ast.Program) (string, error) {
	result := ""

	block := c.CompileBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, 0)

	for _, data := range c.data {
		result += data + "\n"
	}

	result += block

	return result, nil
}

func (c *Compiler) CompileBlock(block ast.BlockStatement, indent int) string {
	result := ""
	indentation := strings.Repeat("    ", indent)

	c.Environment.TempCounter = 0

	for _, stmt := range block.Body {
		compiled := c.CompileStatement(stmt, indent)
		if len(compiled) > 0 {
			result += indentation + compiled + "\n"
		}
	}

	return result
}

func (c *Compiler) CompileStatement(stmt ast.Statement, indent int) string {
	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		return strconv.FormatInt(int64(stmt.Value), 10)
	case *ast.FloatLiteralStatement:
		return c.CompileFloatLiteralStatement(stmt, indent+1)
	case *ast.StringLiteralStatement:
		id := c.dataCounter

		c.data = append(c.data, fmt.Sprintf("data $%s.%d = { b \"%s\", b 0 }", c.Environment.CurrentFunction.Name, id, stmt.Value))
		c.dataCounter += 1

		return fmt.Sprintf("l $%s.%d", c.Environment.CurrentFunction.Name, id)
	case *ast.DeclarationStatement:
		return c.CompileDeclarationStatement(stmt, indent+1)
	case *ast.IdentifierLiteralStatement:
		variable := c.Environment.Get(stmt.Value)
		return fmt.Sprintf("%s %%%s.%d", c.StoreType(variable.declaration.Type), stmt.Value, variable.id)
	case *ast.FunctionCall:
		return c.CompileFunctionCall(stmt, indent+1)
	case *ast.FunctionDeclaration:
		return c.CompileFunctionDeclaration(stmt, indent+1)
	case *ast.ReturnStatement:
		return c.CompileReturnStatement(stmt, indent+1)
	case *ast.BlockStatement:
		return c.CompileBlock(*stmt, indent+1)
	case *ast.BinaryExpression:
		return c.CompileBinaryExpression(stmt, indent+1)
	}

	fmt.Printf("Unknown statement: %T\n", stmt)

	return ""
}
