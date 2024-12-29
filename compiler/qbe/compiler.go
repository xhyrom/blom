package qbe

import (
	"blom/ast"
	"blom/compiler"
	"blom/env"
	"fmt"
	"strconv"
	"strings"
)

type Compiler struct {
	source      string
	data        []string
	dataCounter int
	tempCounter int
	globalScope *env.Environment[*Variable]
}

type Variable struct {
	Type compiler.Type
	Id   int
}

type QbeIdentifier struct {
	Name string
	Id   int
	Raw  bool
	Type compiler.Type
}

func (qbeIdentifier *QbeIdentifier) String() string {
	return fmt.Sprintf("%s %s", qbeIdentifier.Type, qbeIdentifier.Name)
}

func New(file string, functions map[string]*ast.FunctionDeclaration) Compiler {
	globalScope := env.New[*Variable]()

	for _, function := range functions {
		globalScope.SetFunction(function.Name, function)
	}

	return Compiler{
		source:      file,
		data:        make([]string, 0),
		dataCounter: 0,
		tempCounter: 0,
		globalScope: globalScope,
	}
}

func (c *Compiler) Compile(program *ast.Program) (string, error) {
	result := ""

	block, _ := c.CompileBlock(ast.BlockStatement{
		Body: program.Body,
		Loc:  program.Loc,
	}, c.globalScope, false)

	for _, data := range c.data {
		result += data + "\n"
	}

	for _, block := range block {
		result += block
	}

	return result, nil
}

func (c *Compiler) CompileBlock(block ast.BlockStatement, scope *env.Environment[*Variable], labels bool) ([]string, *QbeIdentifier) {
	result := make([]string, 0)

	id := c.tempCounter
	if labels {
		result = append(result, fmt.Sprintf("@block.start.%d", id))
	}

	newScope := env.New[*Variable](*scope)

	for _, stmt := range block.Body {
		indentation := strings.Repeat("    ", 1)

		if stmt.Kind() == ast.FunctionDeclarationNode || labels {
			indentation = ""
		}

		compiled, _ := c.CompileStatement(stmt, newScope)
		for _, compiled := range compiled {
			if strings.HasPrefix(compiled, "@") {
				result = append(result, strings.TrimSpace(compiled)+"\n")
			} else {
				result = append(result, indentation+strings.TrimSpace(compiled)+"\n")
			}
		}
	}

	if labels {
		result = append(result, fmt.Sprintf("@block.end.%d", id))

		c.tempCounter += 1
	}

	return result, &QbeIdentifier{
		Id: id,
	}
}

func (c *Compiler) CompileStatement(stmt ast.Statement, scope *env.Environment[*Variable]) ([]string, *QbeIdentifier) {
	c.tempCounter += 1

	switch stmt := stmt.(type) {
	case *ast.IntLiteralStatement:
		val := strconv.FormatInt(int64(stmt.Value), 10)
		return []string{}, &QbeIdentifier{
			Name: val,
			Raw:  true,
			Type: stmt.Type,
		}
	case *ast.FloatLiteralStatement:
		return c.CompileFloatLiteralStatement(stmt)
	case *ast.BooleanLiteralStatement:
		return c.CompileBooleanLiteralStatement(stmt, scope)
	case *ast.StringLiteralStatement:
		id := c.dataCounter

		c.data = append(c.data, fmt.Sprintf("data $.%d = { b \"%s\", b 0 }", id, stmt.Value))
		c.dataCounter += 1

		return []string{}, &QbeIdentifier{
			Name: fmt.Sprintf("$.%d", id),
			Type: compiler.String,
		}
	case *ast.VariableDeclarationStatement:
		return c.CompileVariableDeclarationStatement(stmt, scope, c.tempCounter)
	case *ast.AssignmentStatement:
		return c.CompileAssignmentStatement(stmt, scope)
	case *ast.IdentifierLiteralStatement:
		variable, exists := scope.FindVariable(stmt.Value)
		if !exists {
			panic("VARIABLE NOT FOUND IN COMPILER, SHOULD HAPPEN IN ANALYZER")
		}

		return []string{}, &QbeIdentifier{
			Name: fmt.Sprintf("%%%s.%d", stmt.Value, variable.Id),
			Type: variable.Type,
		}
	case *ast.FunctionCall:
		return c.CompileFunctionCall(stmt, scope)
	case *ast.FunctionDeclaration:
		return c.CompileFunctionDeclaration(stmt, scope), nil
	case *ast.ReturnStatement:
		return c.CompileReturnStatement(stmt, scope)
	case *ast.BlockStatement:
		return c.CompileBlock(*stmt, scope, true)
	case *ast.BinaryExpression:
		return c.CompileBinaryExpression(stmt, scope)
	case *ast.UnaryExpression:
		return c.CompileUnaryExpression(stmt, scope)
	case *ast.IfStatement:
		return c.CompileIfStatement(stmt, scope)
	case *ast.WhileLoopStatement:
		return c.CompileWhileLoopStatement(stmt, scope)
	case *ast.CompileTimeFunctionCall:
		return c.CompileCompileTimeFunctionCall(stmt, scope)
	}

	panic(fmt.Sprintf("Unknown statement: %T\n", stmt))
}
