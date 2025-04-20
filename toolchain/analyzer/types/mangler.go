package types

import (
	"blom/ast"
	"strconv"
)

func (a *TypeAnalyzer) mangleFunctionDeclaration(decl *ast.FunctionDeclaration) {
	fn, i, _ := a.FunctionManager.GetByDeclarationWithIndex(decl)

	fn.Name.Value = mangledFunctionName(decl, i)
}

func mangledFunctionName(decl *ast.FunctionDeclaration, index int) string {
	if decl.Name.Value == "main" {
		return decl.Name.Value
	}

	return decl.Name.Value + "." + strconv.Itoa(index)
}

func modifyPathLikeFunName(decl *ast.FunctionDeclaration, path ast.Path) ast.Path {
	if decl.Name.Value == "main" {
		return path
	}

	var use int

	previousLastDot := -1
	lastDot := -1
	for i, c := range decl.Name.Value {
		if c == '.' {
			previousLastDot = lastDot
			lastDot = i
		}
	}

	if previousLastDot != -1 && lastDot != -1 {
		namespace := decl.Name.Value[previousLastDot+1 : lastDot]

		for i := 0; i < len(path.Segments); i++ {
			if path.Segments[i].Value == namespace {
				use = i
				break
			}
		}
	}

	path.Segments = path.Segments[:use+1]

	if lastDot != -1 && lastDot+1 < len(decl.Name.Value) {
		path.Segments = append(path.Segments, ast.IdentifierLiteral{
			Value: decl.Name.Value[lastDot+1:],
		})
	}

	return path
}
