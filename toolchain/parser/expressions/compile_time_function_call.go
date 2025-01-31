package expressions

import (
	"blom/ast"
	"blom/debug"
	"blom/tokens"
)

func ParseCompileTimeFunctionCall(p Parser) *ast.BuiltinFunctionCall {
	p.Consume()

	identifier := p.Consume()
	name := identifier.Value
	parameters := make([]ast.Expression, 0)

	p.Consume() // consume (

	for p.Current().Kind != tokens.RightParenthesis {
		exp, err := p.ParseExpression()
		if err != nil {
			dbg := debug.NewSourceLocation(p.Source(), identifier.Location.Row, identifier.Location.Column+2)
			dbg.ThrowError(
				err.Error(),
				true,
				debug.NewHint("Did you forget to close a function call?", ")"),
			)
		}

		parameters = append(parameters, exp)

		if p.Current().Kind != tokens.Comma {
			if p.Current().Kind != tokens.RightParenthesis {
				dbg := debug.NewSourceLocationFromExpression(p.Source(), parameters[len(parameters)-1])
				dbg.ThrowError(
					"Expected comma or right parenthesis",
					true,
					debug.NewHint("Add comma for more parameters", ","),
					debug.NewHint("Add closing parenthesis to end function call", ")"),
				)
			}

			break
		}

		p.Consume()
	}

	last := p.Consume()

	return &ast.BuiltinFunctionCall{
		Name:       name,
		Parameters: parameters,
		Loc:        last.Location,
	}
}
