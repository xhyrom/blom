package expressions

import (
	"blom/ast"
	"blom/env"
	"blom/env/objects"
)

func InterpretCompileTimeFunctionCall(interpreter Interpreter, environment *env.Scope[objects.Object], call *ast.CompileTimeFunctionCall) objects.Object {
	return runFunction(interpreter, environment, call)
}

func runFunction(interpreter Interpreter, environment *env.Scope[objects.Object], call *ast.CompileTimeFunctionCall) objects.Object {
	switch call.Name {
	case "cast":
		return cast(interpreter, environment, call)
	}

	panic("Unknown compile time function call")
}

func cast(interpreter Interpreter, environment *env.Scope[objects.Object], call *ast.CompileTimeFunctionCall) objects.Object {
	stmt := interpreter.InterpretStatement(call.Parameters[0], environment)
	requiredType, _ := ast.ParseType(call.Parameters[1].(*ast.IdentifierLiteral).Value)

	switch stmt := stmt.(type) {
	case *objects.UnsignedLongObject, *objects.UnsignedByteObject, *objects.ShortObject, *objects.LongObject, *objects.FloatObject, *objects.ByteObject, *objects.DoubleObject, *objects.IntObject:
		var value interface{}
		switch v := stmt.(type) {
		case *objects.UnsignedLongObject:
			value = v.Value
		case *objects.UnsignedByteObject:
			value = v.Value
		case *objects.ShortObject:
			value = v.Value
		case *objects.LongObject:
			value = v.Value
		case *objects.FloatObject:
			value = v.Value
		case *objects.ByteObject:
			value = v.Value
		case *objects.DoubleObject:
			value = v.Value
		case *objects.IntObject:
			value = v.Value
		}

		switch requiredType {
		case ast.Float64:
			switch v := value.(type) {
			case int64:
				return &objects.DoubleObject{Value: float64(v)}
			case int32:
				return &objects.DoubleObject{Value: float64(v)}
			case int16:
				return &objects.DoubleObject{Value: float64(v)}
			case int8:
				return &objects.DoubleObject{Value: float64(v)}
			case uint64:
				return &objects.DoubleObject{Value: float64(v)}
			case uint32:
				return &objects.DoubleObject{Value: float64(v)}
			case uint16:
				return &objects.DoubleObject{Value: float64(v)}
			case uint8:
				return &objects.DoubleObject{Value: float64(v)}
			case float32:
				return &objects.DoubleObject{Value: float64(v)}
			case float64:
				return &objects.DoubleObject{Value: float64(v)}
			}
		case ast.Float32:
			switch v := value.(type) {
			case int64:
				return &objects.FloatObject{Value: float32(v)}
			case int32:
				return &objects.FloatObject{Value: float32(v)}
			case int16:
				return &objects.FloatObject{Value: float32(v)}
			case int8:
				return &objects.FloatObject{Value: float32(v)}
			case uint64:
				return &objects.FloatObject{Value: float32(v)}
			case uint32:
				return &objects.FloatObject{Value: float32(v)}
			case uint16:
				return &objects.FloatObject{Value: float32(v)}
			case uint8:
				return &objects.FloatObject{Value: float32(v)}
			case float32:
				return &objects.FloatObject{Value: float32(v)}
			case float64:
				return &objects.FloatObject{Value: float32(v)}
			}
		case ast.Int64:
			switch v := value.(type) {
			case int64:
				return &objects.LongObject{Value: int64(v)}
			case int32:
				return &objects.LongObject{Value: int64(v)}
			case int16:
				return &objects.LongObject{Value: int64(v)}
			case int8:
				return &objects.LongObject{Value: int64(v)}
			case uint64:
				return &objects.LongObject{Value: int64(v)}
			case uint32:
				return &objects.LongObject{Value: int64(v)}
			case uint16:
				return &objects.LongObject{Value: int64(v)}
			case uint8:
				return &objects.LongObject{Value: int64(v)}
			case float32:
				return &objects.LongObject{Value: int64(v)}
			case float64:
				return &objects.LongObject{Value: int64(v)}
			}
		case ast.Int32:
			switch v := value.(type) {
			case int64:
				return &objects.IntObject{Value: int32(v)}
			case int32:
				return &objects.IntObject{Value: int32(v)}
			case int16:
				return &objects.IntObject{Value: int32(v)}
			case int8:
				return &objects.IntObject{Value: int32(v)}
			case uint64:
				return &objects.IntObject{Value: int32(v)}
			case uint32:
				return &objects.IntObject{Value: int32(v)}
			case uint16:
				return &objects.IntObject{Value: int32(v)}
			case uint8:
				return &objects.IntObject{Value: int32(v)}
			case float32:
				return &objects.IntObject{Value: int32(v)}
			case float64:
				return &objects.IntObject{Value: int32(v)}
			}
		case ast.Int16:
			switch v := value.(type) {
			case int64:
				return &objects.ShortObject{Value: int16(v)}
			case int32:
				return &objects.ShortObject{Value: int16(v)}
			case int16:
				return &objects.ShortObject{Value: int16(v)}
			case int8:
				return &objects.ShortObject{Value: int16(v)}
			case uint64:
				return &objects.ShortObject{Value: int16(v)}
			case uint32:
				return &objects.ShortObject{Value: int16(v)}
			case uint16:
				return &objects.ShortObject{Value: int16(v)}
			case uint8:
				return &objects.ShortObject{Value: int16(v)}
			case float32:
				return &objects.ShortObject{Value: int16(v)}
			case float64:
				return &objects.ShortObject{Value: int16(v)}
			}
		case ast.Int8:
			switch v := value.(type) {
			case int64:
				return &objects.ByteObject{Value: int8(v)}
			case int32:
				return &objects.ByteObject{Value: int8(v)}
			case int16:
				return &objects.ByteObject{Value: int8(v)}
			case int8:
				return &objects.ByteObject{Value: int8(v)}
			case uint64:
				return &objects.ByteObject{Value: int8(v)}
			case uint32:
				return &objects.ByteObject{Value: int8(v)}
			case uint16:
				return &objects.ByteObject{Value: int8(v)}
			case uint8:
				return &objects.ByteObject{Value: int8(v)}
			case float32:
				return &objects.ByteObject{Value: int8(v)}
			case float64:
				return &objects.ByteObject{Value: int8(v)}
			}
		case ast.UnsignedInt64:
			switch v := value.(type) {
			case int64:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case int32:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case int16:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case int8:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case uint64:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case uint32:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case uint16:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case uint8:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case float32:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			case float64:
				return &objects.UnsignedLongObject{Value: uint64(v)}
			}
		case ast.UnsignedInt32:
			switch v := value.(type) {
			case int64:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case int32:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case int16:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case int8:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case uint64:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case uint32:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case uint16:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case uint8:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case float32:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			case float64:
				return &objects.UnsignedIntObject{Value: uint32(v)}
			}
		case ast.UnsignedInt16:
			switch v := value.(type) {
			case int64:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case int32:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case int16:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case int8:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case uint64:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case uint32:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case uint16:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case uint8:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case float32:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			case float64:
				return &objects.UnsignedShortObject{Value: uint16(v)}
			}
		case ast.UnsignedInt8:
			switch v := value.(type) {
			case int64:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case int32:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case int16:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case int8:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case uint64:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case uint32:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case uint16:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case uint8:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case float32:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			case float64:
				return &objects.UnsignedByteObject{Value: uint8(v)}
			}
		}
	}

	panic("Unknown cast")
}
