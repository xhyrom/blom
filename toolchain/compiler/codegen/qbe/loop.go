package qbe

import (
	"blom/ast"
	"blom/qbe"
	"fmt"
)

func (c *Codegen) compileLoop(loopStatement *ast.WhileLoop, function *qbe.Function, vtype qbe.Type, isReturn bool) *qbe.TypedValue {
	c.Scopes.Append()

	c.TempCounter += 1

	conditionLabel := fmt.Sprintf("loop.%d.cond", c.TempCounter)
	loopLabel := fmt.Sprintf("loop.%d.body", c.TempCounter)
	endLabel := fmt.Sprintf("loop.%d.end", c.TempCounter)

	function.AddBlock(conditionLabel)

	// Loop condition
	condition := c.compileStatement(loopStatement.Condition, function, vtype, isReturn)

	function.LastBlock().AddInstruction(
		qbe.NewJumpNonZeroInstruction(
			condition.Value,
			loopLabel,
			endLabel,
		),
	)

	// Loop body
	function.AddBlock(loopLabel)

	for _, statement := range loopStatement.Block.Body {
		c.compileStatement(statement, function, nil, isReturn)
	}

	if !function.LastBlock().IsLastStatement(qbe.Jump) &&
		!function.LastBlock().IsLastStatement(qbe.Return) &&
		!function.LastBlock().IsLastStatement(qbe.JumpNonZero) {
		function.LastBlock().AddInstruction(
			qbe.NewJumpInstruction(conditionLabel),
		)
	}

	// End of loop
	function.AddBlock(endLabel)

	c.Scopes.Pop()

	return nil
}
