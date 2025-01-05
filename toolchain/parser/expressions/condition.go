package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseIf(p Parser) *ast.If {
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

	thenBlock := ParseBlock(p)
	var elseBlock *ast.BlockStatement = &ast.BlockStatement{Body: []ast.Statement{}}

	loc := thenBlock.Loc

	maybe_else := p.Current()
	if maybe_else.Kind == tokens.Else {
		p.Consume()

		if p.Current().Kind != tokens.LeftCurlyBracket {
			dbg := debug.NewSourceLocation(p.Source(), maybe_else.Location.Row, maybe_else.Location.Column+1)
			dbg.ThrowError("Missing block", true, debug.NewHint("Add '{'", " {"))
		}

		block := ParseBlock(p)

		elseBlock = block
		loc = elseBlock.Loc
	}

	return &ast.If{
		Condition: condition,
		Then:      thenBlock.Body,
		Else:      elseBlock.Body,
		Loc:       loc,
	}
}
