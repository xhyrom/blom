package qbe

import (
	"blom/ast"
	"blom/env"
	"blom/qbe"
	"fmt"
)

type Compiler struct {
	TempCounter int
	Module      qbe.Module
	Scopes      []*env.Environment[*qbe.TypedValue]
}

func New() *Compiler {
	return &Compiler{
		TempCounter: 0,
		Module:      qbe.NewModule(),
		Scopes:      make([]*env.Environment[*qbe.TypedValue], 0),
	}
}

func (c *Compiler) Compile(program *ast.Program) {
	for _, primitive := range program.Body {
		c.compilePrimitive(primitive)
	}
}

func (c *Compiler) Emit() string {
	return c.Module.String()
}

func (c *Compiler) compilePrimitive(primitive ast.Statement) {
	switch primitive := primitive.(type) {
	case *ast.FunctionDeclaration:
		c.Module.AddFunction(c.compileFunction(primitive))
	default:
		panic(fmt.Sprintf("'%T' is not a valid primitive", primitive))
	}
}

func (c *Compiler) assignNameToValue() string {
	return c.assignNameToValueWithPrefix("")
}

func (c *Compiler) assignNameToValueWithPrefix(prefix string) string {
	c.TempCounter += 1
	return fmt.Sprintf("%s.%d", prefix, c.TempCounter)
}

func (c *Compiler) newTemporaryValue(name *string) *qbe.TemporaryValue {
	var prefix string
	if name != nil {
		prefix = *name
	} else {
		prefix = "tmp"
	}

	return &qbe.TemporaryValue{
		Name: c.assignNameToValueWithPrefix(prefix),
	}
}

func (c *Compiler) getVariable(name string) *qbe.TypedValue {
	for i := len(c.Scopes) - 1; i >= 0; i-- {
		value := c.Scopes[i].Get(name)
		if value != nil {
			return value
		}
	}

	return nil
}

func (c *Compiler) getVariableOr(name string, defaultValue *qbe.TypedValue) *qbe.TypedValue {
	value := c.getVariable(name)
	if value == nil {
		return defaultValue
	}

	return value
}
