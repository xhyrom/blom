package mangling

import (
	"blom/ast"
	"fmt"
	"strconv"
	"strings"
)

type MangledName string
type Encoder struct{}

func NewEncoder() *Encoder {
	return &Encoder{}
}

// format: _Z + nameLength + name + paramTypes + _R + returnType
func (e *Encoder) EncodeFunctionName(name string, params []ast.Type, returnType ast.Type) MangledName {
	var sb strings.Builder

	sb.WriteString("_Z")

	components := strings.Split(name, ".")
	if len(components) > 1 {
		sb.WriteString("N")
		for _, component := range components {
			sb.WriteString(strconv.Itoa(len(component)))
			sb.WriteString(component)
		}
		sb.WriteString("E")
	} else {
		sb.WriteString(strconv.Itoa(len(name)))
		sb.WriteString(name)
	}

	if len(params) == 0 {
		sb.WriteString("v") // void parameter list
	} else {
		for _, param := range params {
			sb.WriteString(string(e.encodeType(param)))
		}
	}

	sb.WriteString("_R")
	sb.WriteString(string(e.encodeType(returnType)))

	return MangledName(sb.String())
}

func (e *Encoder) encodeType(t ast.Type) string {
	if t.IsPointer() {
		ptrType := t.(ast.PointerType)
		return "P" + e.encodeType(ptrType.Inner)
	}

	if t.IsFunction() {
		fnType := t.(ast.FunctionType)
		var sb strings.Builder
		sb.WriteString("F")

		// return type
		sb.WriteString(e.encodeType(fnType.ReturnType))

		// parameter types
		for _, arg := range fnType.Arguments {
			sb.WriteString(e.encodeType(arg))
		}

		sb.WriteString("E") // end of function type
		return sb.String()
	}

	// basic type
	if code, ok := TypeCodeMap[t.String()]; ok {
		return string(code)
	}

	// unknown type - prefix with X and length
	typeStr := t.String()
	return fmt.Sprintf("X%d%s", len(typeStr), typeStr)
}
