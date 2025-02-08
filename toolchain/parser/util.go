package parser

import (
	"blom/ast"
	"blom/debug"
	"fmt"
)

func parseAnnotation(p *Parser) ast.Annotation {
	p.Consume()

	name := p.Consume()
	typ := ast.ParseAnnotation(name.Value)

	if typ == -1 {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Annotation \"%s\" is not recognized", name.Value), true)
	}

	return ast.Annotation{
		Type: typ,
		Loc:  name.Location,
	}
}
