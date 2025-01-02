package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/scope"
	"fmt"
)

type Compiler struct {
	TempCounter int
	Module      qbe.Module
	Scopes      []scope.Scope[*qbe.TypedValue]
}

func New() *Compiler {
	return &Compiler{
		TempCounter: 0,
		Module:      qbe.NewModule(),
		Scopes:      make([]scope.Scope[*qbe.TypedValue], 0),
	}
}

func (c *Compiler) Compile(program *ast.Program) {
	for _, primitive := range program.Body {
		c.compilePrimitive(primitive, true)
	}

	for _, primitive := range program.Body {
		c.compilePrimitive(primitive, false)
	}
}

func (c *Compiler) Emit() string {
	return c.Module.String()
}

func (c *Compiler) compilePrimitive(primitive ast.Statement, populate bool) {
	switch primitive := primitive.(type) {
	case *ast.FunctionDeclaration:
		if populate {
			c.Module.AddFunction(qbe.RemapAstFunction(*primitive))
		} else {
			c.compileFunction(primitive)
		}
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

func (c *Compiler) getTemporaryValue(name *string) *qbe.TemporaryValue {
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

func (c *Compiler) createVariable(t qbe.Type, name string) *qbe.TemporaryValue {
	tmp := c.getTemporaryValue(&name)

	c.Scopes[len(c.Scopes)-1].Set(name, &qbe.TypedValue{
		Type:  t,
		Value: tmp,
	})

	return tmp
}

func (c *Compiler) getVariable(name string) *qbe.TypedValue {
	for i := len(c.Scopes) - 1; i >= 0; i-- {
		value, exists := c.Scopes[i].Get(name)
		if exists {
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

func (c *Compiler) convertToType(first qbe.Type, second qbe.Type, value qbe.Value, function *qbe.Function) *qbe.TypedValue {
	if first.IsPointer() && second.IsPointer() && (first.(qbe.PointerBox).Inner == qbe.Void || second.(qbe.PointerBox).Inner == qbe.Void) {
		return &qbe.TypedValue{
			Value: value,
			Type:  second,
		}
	}

	if first.Weight() == second.Weight() {
		return &qbe.TypedValue{
			Value: value,
			Type:  second,
		}
	} else if (first.IsInteger() && second.IsInteger()) || (first.IsFloatingPoint() && second.IsFloatingPoint()) {
		name := "conv"
		conv := c.getTemporaryValue(&name)

		var instruction qbe.Instruction
		if first.Weight() > second.Weight() {
			if first.IsFloatingPoint() {
				instruction = qbe.NewTruncateInstruction(value)
			} else {
				instruction = qbe.NewCopyInstruction(value)
			}
		} else {
			instruction = qbe.NewExtensionInstruction(first, value)
		}

		function.LastBlock().AddAssign(
			conv,
			second,
			instruction,
		)

		return &qbe.TypedValue{
			Value: conv,
			Type:  second,
		}
	} else {
		name := "conv"
		conv := c.getTemporaryValue(&name)

		function.LastBlock().AddAssign(
			conv,
			second,
			qbe.NewConversionInstruction(first, second, value),
		)

		return &qbe.TypedValue{
			Value: conv,
			Type:  second,
		}
	}
}
