package qbe

import (
	"blom/ast"
	"blom/scope"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

type Codegen struct {
	TempCounter int
	Module      ir.Module
	Scopes      *scope.Scopes[*ir.TypedValue]
}

func New() *Codegen {
	return &Codegen{
		TempCounter: 0,
		Module:      ir.NewModule(),
		Scopes:      scope.NewScopes[*ir.TypedValue](),
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
			c.Module.AddFunction(RemapAstFunction(*primitive))
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

func (c *Codegen) getTemporaryValue(name *string) *ir.TemporaryValue {
	var prefix string
	if name != nil {
		prefix = *name
	} else {
		prefix = "tmp"
	}

	return &ir.TemporaryValue{
		Name: c.assignNameToValueWithPrefix(prefix),
	}
}

func (c *Codegen) getGlobalValue(name *string) *ir.GlobalValue {
	return &ir.GlobalValue{
		Name: c.assignNameToValueWithPrefix(*name),
	}
}

func (c *Codegen) createVariable(t ir.Type, name string) *ir.TemporaryValue {
	tmp := c.getTemporaryValue(&name)

	c.Scopes.Set(name, &ir.TypedValue{
		Type:  t,
		Value: tmp,
	})

	return tmp
}

func (c *Codegen) createGlobalVariable(t ir.Type, name string) *ir.GlobalValue {
	tmp := c.getGlobalValue(&name)

	c.Scopes.Set(name, &ir.TypedValue{
		Type:  t,
		Value: tmp,
	})

	return tmp
}

func (c *Codegen) convertToType(first ir.Type, second ir.Type, value ir.Value, function *ir.Function) *ir.TypedValue {
	if first.IsPointer() && second.IsPointer() && (first.(ir.PointerBox).Inner == ir.Void || second.(ir.PointerBox).Inner == ir.Void) {
		return &ir.TypedValue{
			Value: value,
			Type:  second,
		}
	}

	if first.Weight() == second.Weight() {
		return &ir.TypedValue{
			Value: value,
			Type:  second,
		}
	} else if (first.IsInteger() && second.IsInteger()) || (first.IsFloatingPoint() && second.IsFloatingPoint()) {
		name := "conv"
		conv := c.getTemporaryValue(&name)

		var instruction ir.Instruction
		if first.Weight() > second.Weight() {
			if first.IsFloatingPoint() {
				instruction = ir.NewTruncateInstruction(value)
			} else {
				instruction = ir.NewCopyInstruction(value)
			}
		} else {
			instruction = ir.NewExtensionInstruction(first, value)
		}

		function.LastBlock().AddAssign(
			conv,
			second,
			instruction,
		)

		return &ir.TypedValue{
			Value: conv,
			Type:  second,
		}
	} else {
		name := "conv"
		conv := c.getTemporaryValue(&name)

		function.LastBlock().AddAssign(
			conv,
			second,
			ir.NewConversionInstruction(first, second, value),
		)

		return &ir.TypedValue{
			Value: conv,
			Type:  second,
		}
	}
}

func RemapAstType(t ast.Type) ir.Type {
	switch t {
	case ast.Int8:
		return ir.Byte
	case ast.UnsignedInt8:
		return ir.UnsignedByte
	case ast.Int16:
		return ir.Halfword
	case ast.UnsignedInt16:
		return ir.UnsignedHalfword
	case ast.Int32:
		return ir.Word
	case ast.UnsignedInt32:
		return ir.UnsignedWord
	case ast.Int64:
		return ir.Long
	case ast.UnsignedInt64:
		return ir.UnsignedLong
	case ast.Float32:
		return ir.Single
	case ast.Float64:
		return ir.Double
	case ast.Boolean:
		return ir.Boolean
	case ast.Char:
		return ir.Char
	case ast.String:
		return ir.PointerBox{Inner: ir.Char}
	case ast.Void:
		return ir.Void
	case ast.Null:
		return ir.Null
	}

	if t.IsPointer() {
		return ir.PointerBox{Inner: RemapAstType(t.(ast.PointerType).Inner)}
	}

	if t.IsFunction() {
		fnType := t.(ast.FunctionType)

		lambda := ir.Function{
			Linkage:    ir.NewLinkage(false),
			Params:     make([]ir.TypedValue, len(fnType.Arguments)),
			ReturnType: RemapAstType(fnType.ReturnType),
		}

		return ir.FunctionBox{Inner: lambda}
	}

	panic(fmt.Sprintf("Unknown type '%s'", t))
}

func RemapAstFunction(fun ast.FunctionDeclaration) ir.Function {
	params := make([]ir.TypedValue, len(fun.Params))

	for i, param := range fun.Params {
		params[i] = ir.TypedValue{
			Type:  RemapAstType(param.Type),
			Value: ir.NewTemporaryValue(param.Name.Value),
		}
	}

	return ir.Function{
		Linkage:    ir.NewLinkage(fun.HasAnnotation(ast.Public)),
		Name:       fun.Path.Dotify(),
		Params:     params,
		ReturnType: RemapAstType(fun.Return),
		Variadic:   fun.Variadic,
		External:   fun.HasAnnotation(ast.Native),
		Blocks:     make([]ir.Block, 0),
	}
}
