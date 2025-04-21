package mangling

import (
	"blom/ast"
	"strings"
)

func shouldSkipMangling(fn *ast.FunctionDeclaration) bool {
	return fn.Path.Dotify() == "main" || fn.HasAnnotation(ast.NoMangle) || fn.HasAnnotation(ast.Native)
}

func GetOriginalFunctionName(mangledName string) string {
	// If it's not a mangled name, return as is
	if !strings.HasPrefix(mangledName, "_Z") {
		return mangledName
	}

	// If it's a simple namespace path
	if strings.Contains(mangledName, ".") {
		return strings.Split(mangledName, ".")[0]
	}

	return mangledName
}
