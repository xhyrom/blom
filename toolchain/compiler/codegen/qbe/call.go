package qbe

import (
	"blom/ast"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileCall(call ast.Call, function *ir.Function, vtype ir.Type) *ir.TypedValue {
	switch call := call.(type) {
	case *ast.FunctionCall:
		return c.compileFunctionCall(call, function, vtype)
	case *ast.MethodCall:
		return c.compileMethodCall(call, function, vtype)
	case *ast.InfixCall:
		return c.compileInfixCall(call, function, vtype)
	}

	panic("invalid call")
}

func (c *Codegen) compileFunctionCall(call *ast.FunctionCall, currentFunction *ir.Function, vtype ir.Type) *ir.TypedValue {
	function := c.Module.GetFunctionByName(call.Path.Dotify())
	name := ir.NewGlobalValue(function.Name)

	arguments := make([]ir.TypedValue, 0)
	for i, arg := range call.Args {
		var paramType ir.Type
		if i < len(function.Params) {
			paramType = function.Params[i].Type
		} else {
			paramType = vtype
		}

		value := *c.compileStatement(arg, currentFunction, paramType, false)

		// Promotes f32 to f64 acording to the ISO C standard
		// https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
		if value.Type == ir.Single && i >= len(function.Params) {
			value = *c.convertToType(value.Type, ir.Double, value.Value, currentFunction)
		}

		if len(function.Params) == i && function.Variadic {
			arguments = append(arguments, ir.TypedValue{
				Value: ir.NewLiteralValue("..."),
				Type:  ir.Null,
			})
		}

		arguments = append(arguments, value)
	}

	tempValue := c.getTemporaryValue(nil)

	currentFunction.LastBlock().AddAssign(
		tempValue,
		function.ReturnType,
		ir.NewCallInstruction(name, arguments...),
	)

	return &ir.TypedValue{
		Type:  function.ReturnType,
		Value: tempValue,
	}
}

func (c *Codegen) compileMethodCall(call *ast.MethodCall, currentFunction *ir.Function, vtype ir.Type) *ir.TypedValue {
	// TODO: implement method call
	panic("not implemented method call")
}

func (c *Codegen) compileInfixCall(call *ast.InfixCall, currentFunction *ir.Function, vtype ir.Type) *ir.TypedValue {
	if call.FunctionCall != nil {
		return c.compileFunctionCall(call.FunctionCall, currentFunction, vtype)
	}

	return c.compileMethodCall(call.MethodCall, currentFunction, vtype)
}
