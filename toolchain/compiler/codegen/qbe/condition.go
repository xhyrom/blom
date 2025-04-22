package qbe

import (
	"blom/ast"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileCondition(conditionStatement *ast.If, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
	c.Scopes.Append()

	c.TempCounter += 1

	thenLabel := fmt.Sprintf("ift.%d", c.TempCounter)
	elseLabel := fmt.Sprintf("iff.%d", c.TempCounter)
	endLabel := fmt.Sprintf("ife.%d", c.TempCounter)

	var ifZero string
	if conditionStatement.Else != nil && len(conditionStatement.Else.Body) > 0 {
		ifZero = elseLabel
	} else {
		ifZero = endLabel
	}

	// If condition
	condition := c.compileStatement(conditionStatement.Condition, function, vtype, isReturn)

	function.LastBlock().AddInstruction(
		ir.NewJumpNonZeroInstruction(
			condition.Value,
			thenLabel,
			ifZero,
		),
	)

	// Then block
	function.AddBlock(thenLabel)

	for _, statement := range conditionStatement.Then.Body {
		c.compileStatement(statement, function, nil, isReturn)
	}

	// Else block
	if conditionStatement.Else != nil && len(conditionStatement.Else.Body) > 0 {
		if !function.LastBlock().IsLastStatement(ir.Jump) &&
			!function.LastBlock().IsLastStatement(ir.Return) &&
			!function.LastBlock().IsLastStatement(ir.JumpNonZero) {
			function.LastBlock().AddInstruction(
				ir.NewJumpInstruction(endLabel),
			)
		}

		function.AddBlock(elseLabel)

		for _, statement := range conditionStatement.Else.Body {
			c.compileStatement(statement, function, nil, isReturn)
		}
	}

	// End of if
	function.AddBlock(endLabel)

	c.Scopes.Pop()

	return nil
}
