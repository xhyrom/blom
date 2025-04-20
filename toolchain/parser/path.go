package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func (p *Parser) parsePath(left ast.Node) *ast.Path {
	if left.Kind() != ast.IdentifierNode {
		dbg := debug.NewSourceLocationFromNode(p.Source(), left)
		dbg.ThrowError("Left side of path must be an identifier", true)
	}

	next := left.(*ast.IdentifierLiteral)
	segments := []ast.IdentifierLiteral{*next}

	for p.Current().Kind == tokens.DoubleColon {
		p.Consume()

		if p.Current().Kind != tokens.Identifier {
			dbg := debug.NewSourceLocationFromToken(p.Source(), p.Current())
			dbg.ThrowError("Expected identifier after '::'", true)
		}

		next := p.parseLiteral().(*ast.IdentifierLiteral)
		segments = append(segments, *next)
	}

	return &ast.Path{
		Segments: segments,
		Loc:      next.Loc,
	}
}
