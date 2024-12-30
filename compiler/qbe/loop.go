package qbe

import (
	"blom/ast"
	"blom/env"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileLoop(loopStatement *ast.WhileLoopStatement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	c.Scopes = append(c.Scopes, env.New[*qbe.TypedValue]())

	c.TempCounter += 1

	conditionLabel := fmt.Sprintf("loopc.%d", c.TempCounter)
	loopLabel := fmt.Sprintf("loop.%d", c.TempCounter)
	endLabel := fmt.Sprintf("end.%d", c.TempCounter)

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

	for _, statement := range loopStatement.Body {
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

	c.Scopes = c.Scopes[:len(c.Scopes)-1]

	return nil
}
