package mangling

import (
	"blom/analyzer/types"
	"blom/ast"
	"fmt"
	"strings"
)

func mangleFunctionDeclaration(fn *ast.FunctionDeclaration, ptrMap FunctionPtrMap) {
	if shouldSkipMangling(fn) {
		return
	}

	if mangledName, exists := ptrMap[fn]; exists {
		fn.Name.Value = mangledName
	}
}

func mangleCall(call ast.Call, ptrMap FunctionPtrMap, nameMap FunctionNameMap, analysisContext *types.AnalysisContext) {
	switch c := call.(type) {
	case *ast.FunctionCall:
		mangleFunctionCall(c, ptrMap, nameMap, analysisContext)
	case *ast.MethodCall:
		mangleMethodCall(c, ptrMap, nameMap, analysisContext)
	case *ast.InfixCall:
		mangleInfixCall(c, ptrMap, nameMap, analysisContext)
	default:
		panic("unknown call type")
	}
}

func mangleFunctionCall(call *ast.FunctionCall, ptrMap FunctionPtrMap, nameMap FunctionNameMap, analysisContext *types.AnalysisContext) {
	for _, arg := range call.Args {
		mangleNode(arg, ptrMap, nameMap, analysisContext)
	}

	if fn, exists := analysisContext.FunctionCalls[call]; exists {
		if !shouldSkipMangling(fn) {
			if mangledName, exists := ptrMap[fn]; exists {
				updateCallPath(call, mangledName)
				return
			}
		}
	}

	originalName := call.Path.Dotify()

	if strings.Contains(originalName, ".") {
		parts := strings.Split(originalName, ".")
		originalName = parts[0]
	}

	for key, mangledName := range nameMap {
		prefix := originalName + "#"
		if strings.HasPrefix(key, prefix) {
			updateCallPath(call, mangledName)
			return
		}
	}
}

func mangleMethodCall(call *ast.MethodCall, ptrMap FunctionPtrMap, nameMap FunctionNameMap, analysisContext *types.AnalysisContext) {
	panic("not implemented method call mangling")
}

func mangleInfixCall(call *ast.InfixCall, ptrMap FunctionPtrMap, nameMap FunctionNameMap, analysisContext *types.AnalysisContext) {
	if call.FunctionCall != nil {
		mangleFunctionCall(call.FunctionCall, ptrMap, nameMap, analysisContext)
		return
	}

	mangleMethodCall(call.MethodCall, ptrMap, nameMap, analysisContext)
}

func updateCallPath(call *ast.FunctionCall, mangledName string) {
	parts := strings.Split(mangledName, ".")
	segments := make([]ast.IdentifierLiteral, len(parts))

	for i, part := range parts {
		segments[i] = ast.IdentifierLiteral{
			Value: part,
			Loc:   call.Path.Loc,
		}
	}

	call.Path.Segments = segments
}

func GenerateFunctionMangledName(fn *ast.FunctionDeclaration) string {
	types := make([]string, len(fn.Params))
	for i, param := range fn.Params {
		types[i] = typeToCode(param.Type)
	}

	returnTypeCode := typeToCode(fn.Return)

	return fmt.Sprintf("%s.%s.%s",
		fn.Name.Value,
		strings.Join(types, ""),
		returnTypeCode)
}
