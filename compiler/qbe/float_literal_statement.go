package qbe

import (
	"blom/ast"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *Compiler) CompileFloatLiteralStatement(stmt *ast.FloatLiteralStatement, ident int) ([]string, string) {
	name := fmt.Sprintf("%%tmp.%d", c.Environment.TempCounter)
	c.Environment.TempCounter += 1

	floatStr := strconv.FormatFloat(stmt.Value, 'f', -1, 64)
	parts := strings.Split(floatStr, ".")
	fracDigits := 0
	if len(parts) > 1 {
		fracDigits = len(parts[1])
	}

	wholeNumber := int(stmt.Value * math.Pow(10, float64(fracDigits)))

	result := fmt.Sprintf("%s =d div d_%d, d_%d", name, wholeNumber, int(math.Pow(10, float64(fracDigits))))

	return []string{result, "# ^ float literal statement\n"}, name
}
