package qbe

import (
	"blom/ast"
	"blom/env"
	"blom/qbe"
	"fmt"
)

func (c *Compiler) compileCondition(conditionStatement *ast.IfStatement, function *qbe.Function, vtype *qbe.Type, isReturn bool) *qbe.TypedValue {
	c.Scopes = append(c.Scopes, env.New[*qbe.TypedValue]())

	condition := c.compileStatement(conditionStatement.Condition, function, vtype, isReturn)

	c.TempCounter += 1

	thenLabel := fmt.Sprintf("ift.%d", c.TempCounter)
	elseLabel := fmt.Sprintf("iff.%d", c.TempCounter)
	endLabel := fmt.Sprintf("iend.%d", c.TempCounter)

	var ifZero string
	if conditionStatement.Else != nil && len(conditionStatement.Else) > 0 {
		ifZero = elseLabel
	} else {
		ifZero = endLabel
	}

	function.LastBlock().AddInstruction(
		qbe.NewJumpNonZeroInstruction(
			condition.Value,
			thenLabel,
			ifZero,
		),
	)

	function.AddBlock(thenLabel)

	for _, statement := range conditionStatement.Then {
		c.compileStatement(statement, function, nil, isReturn)
	}

	if conditionStatement.Else != nil && len(conditionStatement.Else) > 0 {
		if !function.LastBlock().IsLastStatement(qbe.Jump) &&
			!function.LastBlock().IsLastStatement(qbe.Return) &&
			!function.LastBlock().IsLastStatement(qbe.JumpNonZero) {
			function.LastBlock().AddInstruction(
				qbe.NewJumpInstruction(endLabel),
			)
		}

		function.AddBlock(elseLabel)

		for _, statement := range conditionStatement.Else {
			c.compileStatement(statement, function, nil, isReturn)
		}
	}

	function.AddBlock(endLabel)

	c.Scopes = c.Scopes[:len(c.Scopes)-1]

	return nil
}
