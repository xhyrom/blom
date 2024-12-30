package qbe

import (
	"blom/ast"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *Compiler) CompileFloatLiteralStatement(stmt *ast.FloatLiteral) ([]string, *QbeIdentifier) {
	name := fmt.Sprintf("%%tmp.%d", c.tempCounter)

	floatStr := strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	parts := strings.Split(floatStr, ".")
	fracDigits := 0
	if len(parts) > 1 {
		fracDigits = len(parts[1])
	}

	wholeNumber := int(stmt.Value * math.Pow(10, float64(fracDigits)))

	result := fmt.Sprintf("%s =%s div %s_%d, %s_%d", name, stmt.Type, stmt.Type, wholeNumber, stmt.Type, int(math.Pow(10, float64(fracDigits))))

	return []string{result, "# ^ float literal statement\n"}, &QbeIdentifier{
		Name: name,
		Type: stmt.Type,
	}
}
