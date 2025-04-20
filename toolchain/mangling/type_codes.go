package mangling

import (
	"blom/ast"
	"strings"
)

var typeCodeMap = map[string]string{
	"i8":     "a",
	"u8":     "b",
	"i16":    "c",
	"u16":    "d",
	"i32":    "e",
	"u32":    "f",
	"i64":    "g",
	"u64":    "h",
	"f32":    "i",
	"f64":    "j",
	"bool":   "k",
	"char":   "l",
	"string": "m",
	"void":   "n",
	"null":   "o",
}

func typeToCode(t ast.Type) string {
	if t.IsPointer() {
		ptrType := t.(ast.PointerType)
		return "p" + typeToCode(ptrType.Dereference())
	}

	if t.IsFunction() {
		fnType := t.(ast.FunctionType)
		params := make([]string, len(fnType.Arguments))
		for i, arg := range fnType.Arguments {
			params[i] = typeToCode(arg)
		}

		return "q" + strings.Join(params, "") + typeToCode(fnType.ReturnType)
	}

	if code, ok := typeCodeMap[t.String()]; ok {
		return code
	}

	return "x" + t.String()
}
