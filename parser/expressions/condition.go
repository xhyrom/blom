package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseIf(p Parser) *ast.IfStatement {
	p.Consume()

	condition, err := p.ParseExpression()
	if err != nil {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError(err.Error(), true, debug.NewHint("Add condition", "true "))
	}

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocationFromExpression(p.Source(), condition)
		dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
	}

	then_block := ParseBlock(p)
	else_block := &ast.BlockStatement{}

	loc := then_block.Loc

	maybe_else := p.Current()
	if maybe_else.Kind == tokens.Else {
		p.Consume()

		if p.Current().Kind != tokens.LeftCurlyBracket {
			dbg := debug.NewSourceLocation(p.Source(), maybe_else.Location.Row, maybe_else.Location.Column+1)
			dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
		}

		block := ParseBlock(p)

		else_block = block
		loc = else_block.Loc
	}

	return &ast.IfStatement{
		Condition: condition,
		Then:      then_block,
		Else:      else_block,
		Loc:       loc,
	}
}
