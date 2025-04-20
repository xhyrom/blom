package ast

import "blom/tokens"

type Path struct {
	Segments []IdentifierLiteral
	Loc      tokens.Location
}
