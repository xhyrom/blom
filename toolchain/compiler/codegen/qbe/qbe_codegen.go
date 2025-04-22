package qbe

import (
	"blom/ast"
	"blom/qbe"
	"blom/scope"
	"fmt"
)

type Codegen struct {
	TempCounter int
	Module      qbe.Module
	Scopes      *scope.Scopes[*qbe.TypedValue]
}

func New() *Codegen {
	return &Codegen{
		TempCounter: 0,
		Module:      qbe.NewModule(),
		Scopes:      scope.NewScopes[*qbe.TypedValue](),
	}
}

func (c *Codegen) Generate(program *ast.Program) {
	for _, primitive := range program.Body {
		c.generatePrimitive(primitive, true)
	}

	for _, primitive := range program.Body {
		c.generatePrimitive(primitive, false)
	}
}

func (c *Codegen) Emit() string {
	return c.Module.String()
}

func (c *Codegen) generatePrimitive(primitive ast.Node, populate bool) {
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

func (c *Codegen) assignNameToValue() string {
	return c.assignNameToValueWithPrefix("")
}

func (c *Codegen) assignNameToValueWithPrefix(prefix string) string {
	c.TempCounter += 1
	return fmt.Sprintf("%s.%d", prefix, c.TempCounter)
}

func (c *Codegen) getTemporaryValue(name *string) *qbe.TemporaryValue {
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

func (c *Codegen) getGlobalValue(name *string) *qbe.GlobalValue {
	return &qbe.GlobalValue{
		Name: c.assignNameToValueWithPrefix(*name),
	}
}

func (c *Codegen) createVariable(t qbe.Type, name string) *qbe.TemporaryValue {
	tmp := c.getTemporaryValue(&name)

	c.Scopes.Set(name, &qbe.TypedValue{
		Type:  t,
		Value: tmp,
	})

	return tmp
}

func (c *Codegen) createGlobalVariable(t qbe.Type, name string) *qbe.GlobalValue {
	tmp := c.getGlobalValue(&name)

	c.Scopes.Set(name, &qbe.TypedValue{
		Type:  t,
		Value: tmp,
	})

	return tmp
}

func (c *Codegen) convertToType(first qbe.Type, second qbe.Type, value qbe.Value, function *qbe.Function) *qbe.TypedValue {
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
