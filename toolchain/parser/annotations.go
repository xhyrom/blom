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

	if p.Current().Kind != tokens.LeftParenthesis {
		return ast.Annotation{
			Type: ty,
			Loc:  name.Location,
		}
	}

	current := p.Consume()

	if p.Current().Kind == tokens.RightParenthesis {
		p.Consume()

		return ast.Annotation{
			Type: ty,
			Loc:  name.Location,
		}
	}

	values := make(map[string]ast.Expression)

	for current.Kind != tokens.RightParenthesis && p.Current().Kind != tokens.RightParenthesis {
		key := p.Consume()

		if p.Consume().Kind != tokens.Assign {
			dbg := debug.NewSourceLocation(p.Source(), key.Location.Row, key.Location.Column)
			dbg.ThrowError("Expected '='", true)
		}

		value := p.parseLiteral()

		current = p.Consume()

		if current.Kind != tokens.Comma && current.Kind != tokens.RightParenthesis {
			dbg := debug.NewSourceLocation(p.Source(), value.Location().Row, value.Location().Column)
			dbg.ThrowError(
				"Expected comma or right parenthesis",
				true,
				debug.NewHint("Arguments must be separated by commas", ","),
				debug.NewHint("Did you forget to close the parentheses?", ")"),
			)
		}

		values[key.Value] = value
	}

	return ast.Annotation{
		Type:   ty,
		Values: values,
		Loc:    name.Location,
	}
}
