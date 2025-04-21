package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"strings"
)

func (p *Parser) parseImport() *ast.Import {
	p.Consume()

	var path strings.Builder
	var alias string

	for !p.IsEof() &&
		p.Current().Kind == tokens.Identifier ||
		p.Current().Kind == tokens.Dot ||
		p.Current().Kind == tokens.Slash {

		switch p.Current().Kind {
		case tokens.Identifier:
			path.WriteString(p.Consume().Value)
		case tokens.Dot:
			path.WriteString(".")
			p.Consume()
		case tokens.Slash:
			path.WriteString("/")
			p.Consume()
		}
	}

	if path.Len() == 0 {
		dbg := debug.NewSourceLocationFromToken(p.Source(), p.Previous())
		dbg.ThrowError("Expected import path", true,
			debug.NewHint("Import requires a path", " path/to/module"))
	}

	return &ast.Import{
		Path:  path.String(),
		Alias: alias,
		Loc:   p.Previous().Location,
	}
}
