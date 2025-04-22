package qbe

import (
	"blom/ast"
	"fmt"

	"github.com/xhyrom/blom/qbe/ir"
)

func (c *Codegen) compileLoop(loopStatement *ast.WhileLoop, function *ir.Function, vtype ir.Type, isReturn bool) *ir.TypedValue {
	c.Scopes.Append()

	c.TempCounter += 1

	conditionLabel := fmt.Sprintf("loop.%d.cond", c.TempCounter)
	loopLabel := fmt.Sprintf("loop.%d.body", c.TempCounter)
	endLabel := fmt.Sprintf("loop.%d.end", c.TempCounter)

	function.AddBlock(conditionLabel)

	// Loop condition
	condition := c.compileStatement(loopStatement.Condition, function, vtype, isReturn)

	function.LastBlock().AddInstruction(
		ir.NewJumpNonZeroInstruction(
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

	if !function.LastBlock().IsLastStatement(ir.Jump) &&
		!function.LastBlock().IsLastStatement(ir.Return) &&
		!function.LastBlock().IsLastStatement(ir.JumpNonZero) {
		function.LastBlock().AddInstruction(
			ir.NewJumpInstruction(conditionLabel),
		)
	}

	// End of loop
	function.AddBlock(endLabel)

	c.Scopes.Pop()

	return nil
}
