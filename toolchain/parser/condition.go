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
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected condition", true, debug.NewHint("Did you forget to add a condition?", " <condition>"))
	}

	thenBlock := p.parseBlock()
	var elseBlock *ast.Block = nil

	loc := thenBlock.Loc

	if p.Current().Kind == tokens.Else {
		p.Consume()

		if p.Current().Kind == tokens.If {
			return &ast.If{
				Condition: condition,
				Then:      *thenBlock,
				Else: &ast.Block{
					Body: []ast.Node{
						p.parseCondition(),
					},
				},
				Loc: loc,
			}
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
