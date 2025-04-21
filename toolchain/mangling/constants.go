package mangling

type TypeCode string

const (
	// basic types
	TypeCodeVoid    TypeCode = "v"
	TypeCodeBool    TypeCode = "b"
	TypeCodeChar    TypeCode = "c"
	TypeCodeInt8    TypeCode = "a"
	TypeCodeUInt8   TypeCode = "h"
	TypeCodeInt16   TypeCode = "s"
	TypeCodeUInt16  TypeCode = "t"
	TypeCodeInt32   TypeCode = "i"
	TypeCodeUInt32  TypeCode = "j"
	TypeCodeInt64   TypeCode = "x"
	TypeCodeUInt64  TypeCode = "y"
	TypeCodeFloat32 TypeCode = "f"
	TypeCodeFloat64 TypeCode = "d"
	TypeCodeString  TypeCode = "S"
	TypeCodeNull    TypeCode = "n"

	// compound types
	TypeCodePointer  TypeCode = "P" // followed by pointer type
	TypeCodeFunction TypeCode = "F" // followed by return type and parameter types, ended with E
)

var TypeCodeMap = map[string]TypeCode{
	"void":   TypeCodeVoid,
	"bool":   TypeCodeBool,
	"char":   TypeCodeChar,
	"i8":     TypeCodeInt8,
	"u8":     TypeCodeUInt8,
	"i16":    TypeCodeInt16,
	"u16":    TypeCodeUInt16,
	"i32":    TypeCodeInt32,
	"u32":    TypeCodeUInt32,
	"i64":    TypeCodeInt64,
	"u64":    TypeCodeUInt64,
	"f32":    TypeCodeFloat32,
	"f64":    TypeCodeFloat64,
	"string": TypeCodeString,
	"null":   TypeCodeNull,
}

var ReverseTypeCodeMap = map[TypeCode]string{
	TypeCodeVoid:    "void",
	TypeCodeBool:    "bool",
	TypeCodeChar:    "char",
	TypeCodeInt8:    "i8",
	TypeCodeUInt8:   "u8",
	TypeCodeInt16:   "i16",
	TypeCodeUInt16:  "u16",
	TypeCodeInt32:   "i32",
	TypeCodeUInt32:  "u32",
	TypeCodeInt64:   "i64",
	TypeCodeUInt64:  "u64",
	TypeCodeFloat32: "f32",
	TypeCodeFloat64: "f64",
	TypeCodeString:  "string",
	TypeCodeNull:    "null",
}
