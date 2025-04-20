package mangling

import (
	"blom/ast"
	"fmt"
	"strings"
)

func shouldSkipMangling(fn *ast.FunctionDeclaration) bool {
	return fn.Name.Value == "main" || fn.HasAnnotation(ast.NoMangle) || fn.HasAnnotation(ast.Native)
}

func getFunctionKey(fn *ast.FunctionDeclaration) string {
	paramTypes := make([]string, len(fn.Params))
	for i, param := range fn.Params {
		paramTypes[i] = param.Type.String()
	}

	originalName := getOriginalFunctionName(fn.Name.Value)

	return fmt.Sprintf("%s_%s", originalName, strings.Join(paramTypes, "_"))
}

func getOriginalFunctionName(mangledName string) string {
	if strings.Contains(mangledName, ".") {
		return strings.Split(mangledName, ".")[0]
	}

	return mangledName
}
