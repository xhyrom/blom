package qbe

import (
	"blom/ast"
	"fmt"
	"math"
)

func (c *Compiler) CompileFloatLiteralStatement(stmt *ast.FloatLiteralStatement, ident int) string {
	result := fmt.Sprintf("%%tmp.%d =%s ", c.Environment.TempCounter, "d")

	_, fracPart := math.Modf(float64(stmt.Value))

	fracDigits := 0
	for fracPart != 0 {
		fracPart *= 10
		fracPart = fracPart - math.Floor(fracPart)
		fracDigits++
	}

	wholeNumber := int(float64(stmt.Value) * math.Pow(10, float64(fracDigits)))
	result += fmt.Sprintf("div d_%d, d_%d", wholeNumber, fracDigits)

	return result
}
