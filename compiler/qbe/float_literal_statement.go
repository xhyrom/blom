package qbe

import (
	"blom/ast"
	"blom/compiler"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *Compiler) CompileFloatLiteralStatement(stmt *ast.FloatLiteralStatement, ident int, expectedType *compiler.Type) ([]string, *Additional) {
	if expectedType == nil {
		double := compiler.Double
		expectedType = &double
	}

	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)

	floatStr := strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	parts := strings.Split(floatStr, ".")
	fracDigits := 0
	if len(parts) > 1 {
		fracDigits = len(parts[1])
	}

	wholeNumber := int(stmt.Value * math.Pow(10, float64(fracDigits)))

	result := fmt.Sprintf("%s =%s div %s_%d, %s_%d", name, expectedType, expectedType, wholeNumber, expectedType, int(math.Pow(10, float64(fracDigits))))

	return []string{result, "# ^ float literal statement\n"}, &Additional{
		Name: name,
		Type: *expectedType,
	}
}
