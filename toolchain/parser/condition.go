package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

// Parses a condition that can have a form:
// if <condition> { <then> } else { <else> }
func (p *Parser) parseCondition() *ast.If {
	p.Consume()

	condition := p.parseExpression()
	if condition == nil {
		dbg := debug.NewSourceLocation(p.Source(), p.Current().Location.Row, p.Current().Location.Column)
		dbg.ThrowError("Expected condition", true, debug.NewHint("Did you forget to add a condition?", "if"))
	}

	if p.Current().Kind != tokens.LeftCurlyBracket {
		dbg := debug.NewSourceLocation(p.Source(), condition.Location().Row, condition.Location().Column)
		dbg.ThrowError("Expected opening bracket", true, debug.NewHint("Did you forget to add an opening bracket?", "{"))
	}

	thenBlock := p.parseBlock()
	var elseBlock *ast.Block = nil

	loc := thenBlock.Loc

	if p.Current().Kind == tokens.Else {
		p.Consume()

		if p.Current().Kind != tokens.LeftCurlyBracket {
			dbg := debug.NewSourceLocation(p.Source(), condition.Location().Row, condition.Location().Column)
			dbg.ThrowError("Expected opening bracket", true, debug.NewHint("Did you forget to add an opening bracket", "{"))
		}

		elseBlock = p.parseBlock()
		loc = elseBlock.Loc
	}

	return &ast.If{
		Condition: condition,
		Then:      *thenBlock,
		Else:      elseBlock,
		Loc:       loc,
	}
}
