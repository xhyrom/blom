package expressions

import (
	"blom/ast"
	"blom/tokens"
)

func ParseMemberAccess(p Parser, left ast.Expression) ast.Expression {
	p.Consume()

	right := ParseIdentifier(p)
	if right, ok := right.(*ast.FunctionCall); ok {
		right.MemberAccess = true
		right.Parameters = append([]ast.Expression{left}, right.Parameters...)

		var exp ast.Expression = right
		if p.Current().Kind == tokens.Dot {
			exp = ParseMemberAccess(p, exp)
		}

		return exp
	}

	// todo: member access
	return left
}
