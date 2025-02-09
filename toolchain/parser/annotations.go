package parser

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
	"fmt"
)

func (p *Parser) collectAnnotations() {
	for p.Current().Kind == tokens.AtMark {
		p.annotations = append(p.annotations, parseAnnotation(p))
	}
}

func (p *Parser) extractAnnotations() []ast.Annotation {
	annotations := p.annotations
	p.annotations = make([]ast.Annotation, 0)

	return annotations
}

func parseAnnotation(p *Parser) ast.Annotation {
	p.Consume()

	name := p.Consume()
	ty := ast.ParseAnnotation(name.Value)

	if ty == -1 {
		dbg := debug.NewSourceLocation(p.Source(), name.Location.Row, name.Location.Column)
		dbg.ThrowError(fmt.Sprintf("Annotation \"%s\" is not recognized", name.Value), true)
	}

	return ast.Annotation{
		Type: ty,
		Loc:  name.Location,
	}
}
