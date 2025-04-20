package qbe

import (
	"blom/ast"
	"blom/qbe"
)

func (c *Compiler) compileCall(call ast.Call, function *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
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

func (c *Compiler) compileFunctionCall(call *ast.FunctionCall, currentFunction *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	function := c.Module.GetFunctionByName(call.Path.Dotify())
	name := qbe.NewGlobalValue(function.Name)

	arguments := make([]qbe.TypedValue, 0)
	for i, arg := range call.Args {
		var paramType qbe.Type
		if i < len(function.Params) {
			paramType = function.Params[i].Type
		} else {
			paramType = vtype
		}

		value := *c.compileStatement(arg, currentFunction, paramType, false)

		// Promotes f32 to f64 acording to the ISO C standard
		// https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
		if value.Type == qbe.Single && i >= len(function.Params) {
			value = *c.convertToType(value.Type, qbe.Double, value.Value, currentFunction)
		}

		if len(function.Params) == i && function.Variadic {
			arguments = append(arguments, qbe.TypedValue{
				Value: qbe.NewLiteralValue("..."),
				Type:  qbe.Null,
			})
		}

		arguments = append(arguments, value)
	}

	tempValue := c.getTemporaryValue(nil)

	currentFunction.LastBlock().AddAssign(
		tempValue,
		function.ReturnType,
		qbe.NewCallInstruction(name, arguments...),
	)

	return &qbe.TypedValue{
		Type:  function.ReturnType,
		Value: tempValue,
	}
}

func (c *Compiler) compileMethodCall(call *ast.MethodCall, currentFunction *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	// TODO: implement method call
	panic("not implemented method call")
}

func (c *Compiler) compileInfixCall(call *ast.InfixCall, currentFunction *qbe.Function, vtype qbe.Type) *qbe.TypedValue {
	if call.FunctionCall != nil {
		return c.compileFunctionCall(call.FunctionCall, currentFunction, vtype)
	}

	return c.compileMethodCall(call.MethodCall, currentFunction, vtype)
}
