package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/env"
	"blom/env/objects"
	"fmt"
	"strings"
)

func InterpretFunctionCall(interpreter Interpreter, environment *env.Environment[objects.Object], call *ast.FunctionCall) objects.Object {
	function := environment.FindFunction(call.Name)

	if function.IsNative() {
		return nativeInterpretFunctionCall(interpreter, environment, call)
	}

	env := env.New(*environment)

	for i, param := range call.Parameters {
		arg := function.Arguments[i]
		obj := interpreter.InterpretStatement(param, env)

		if arg.Type != obj.Type() {
			dbg := debug.NewSourceLocation(interpreter.Source(), call.Loc.Row, call.Loc.Column)
			dbg.ThrowError(fmt.Sprintf("Expected type %s, got %s", arg.Type, obj.Type()), true)
		}

		env.Set(function.Arguments[i].Name, obj)
	}

	return interpreter.InterpretBlock(function.Body, env)
}

func nativeInterpretFunctionCall(interpreter Interpreter, environment *env.Environment[objects.Object], call *ast.FunctionCall) objects.Object {
	switch call.Name {
	case "printf":
		return interpretPrintf(interpreter, environment, call)
	default:
		return nil
	}
}

func interpretPrintf(interpreter Interpreter, environment *env.Environment[objects.Object], call *ast.FunctionCall) objects.Object {
	format := call.Parameters[0].(*ast.StringLiteral).Value

	args := make([]interface{}, 0)

	for _, param := range call.Parameters[1:] {
		obj := interpreter.InterpretStatement(param, environment)

		switch obj := obj.(type) {
		case *objects.IntObject:
			args = append(args, obj.Value)
		case *objects.FloatObject:
			args = append(args, obj.Value)
		case *objects.DoubleObject:
			args = append(args, obj.Value)
		case *objects.StringObject:
			args = append(args, obj.Value)
		case *objects.BooleanObject:
			args = append(args, obj.Value)
		default:
			panic(fmt.Sprintf("Unsupported type %T", obj))
		}
	}

	format = strings.ReplaceAll(format, `\n`, "\n")
	format = strings.ReplaceAll(format, `\t`, "\t")
	format = strings.ReplaceAll(format, `\"`, "\"")
	format = strings.ReplaceAll(format, `\\`, "\\")

	fmt.Printf(format, args...)

	return nil
}
