package tokens

import (
	"fmt"
	"slices"
)

type TokenKind int

const (
	Eof TokenKind = iota
	Illegal

	// Identifiers and literals
	Identifier
	CharLiteral
	StringLiteral
	IntLiteral
	FloatLiteral
	BooleanLiteral

	// Statements
	Assign
	If
	Else
	For
	While

	// Operators
	Equals
	NotEquals
	Plus
	Minus
	Asterisk
	Slash
	PercentSign
	Ampersand
	VerticalLine
	CircumflexAccent
	Tilde
	And
	Or
	LessThan
	DoubleLessThan
	LessThanOrEqual
	GreaterThan
	DoubleGreaterThan
	GreaterThanOrEqual

	// Delimiters
	Dot
	Range
	Ellipsis
	Comma
	Colon
	DoubleColon
	Semicolon
	AtMark
	LeftParenthesis
	RightParenthesis
	LeftSquareBracket
	RightSquareBracket
	LeftCurlyBracket
	RightCurlyBracket

	// Keywords
	Fun
	Return
	Val
	Var
	Import
)

var tokens = []string{
	Eof:                "EOF",
	Illegal:            "Illegal",
	Identifier:         "Identifier",
	CharLiteral:        "CharLiteral",
	StringLiteral:      "StringLiteral",
	IntLiteral:         "IntLiteral",
	FloatLiteral:       "FloatLiteral",
	BooleanLiteral:     "BooleanLiteral",
	Assign:             "=",
	If:                 "if",
	Else:               "else",
	For:                "for",
	While:              "while",
	Equals:             "==",
	NotEquals:          "!=",
	Plus:               "+",
	Minus:              "-",
	Asterisk:           "*",
	Slash:              "/",
	PercentSign:        "%",
	Ampersand:          "&",
	VerticalLine:       "|",
	CircumflexAccent:   "^",
	Tilde:              "~",
	And:                "and",
	Or:                 "or",
	LessThan:           "<",
	DoubleLessThan:     "<<",
	LessThanOrEqual:    "<=",
	GreaterThan:        ">",
	DoubleGreaterThan:  ">>",
	GreaterThanOrEqual: ">=",
	Dot:                ".",
	Range:              "..",
	Ellipsis:           "...",
	Comma:              ",",
	Colon:              ":",
	DoubleColon:        "::",
	Semicolon:          ";",
	AtMark:             "@",
	LeftParenthesis:    "(",
	RightParenthesis:   ")",
	LeftSquareBracket:  "[",
	RightSquareBracket: "]",
	LeftCurlyBracket:   "{",
	RightCurlyBracket:  "}",
	Fun:                "fun",
	Return:             "return",
	Val:                "val",
	Var:                "var",
	Import:             "import",
}

var reserved = []string{
	If:     "if",
	Else:   "else",
	For:    "for",
	While:  "while",
	Fun:    "fun",
	Return: "return",
	Val:    "val",
	Var:    "var",
	Import: "import",
}

func (t TokenKind) String() string {
	return tokens[t]
}

func FromIdentifier(identifier string) TokenKind {
	if identifier == "true" || identifier == "false" {
		return BooleanLiteral
	}

	index := slices.Index(reserved, identifier)
	if index == -1 {
		return Illegal
	}

	return TokenKind(index)
}

type Precedence int

const (
	LowestPrecedence         Precedence = iota
	AssignPrecedence                    // =
	OrPrecedence                        // or (logical OR)
	AndPrecedence                       // and (logical AND)
	BitwiseOrPrecedence                 // |
	BitwiseXorPrecedence                // ^
	BitwiseAndPrecedence                // &
	EqualityPrecedence                  // == !=
	RelationalPrecedence                // < <= > >=
	ShiftPrecedence                     // << >>
	AdditivePrecedence                  // + -
	MultiplicativePrecedence            // * / %
	UnaryPrecedence                     // unary + - ~ (not)
	MemberAccessPrecedence              // . :: (highest: member access, namespace)
	HighestPrecedence                   // parentheses, function calls, array indexing
)

func (kind TokenKind) Precedence() Precedence {
	switch kind {
	case Assign:
		return AssignPrecedence
	case Or:
		return OrPrecedence
	case And:
		return AndPrecedence
	case VerticalLine:
		return BitwiseOrPrecedence
	case CircumflexAccent:
		return BitwiseXorPrecedence
	case Ampersand:
		return BitwiseAndPrecedence
	case Equals, NotEquals:
		return EqualityPrecedence
	case LessThan, LessThanOrEqual, GreaterThan, GreaterThanOrEqual:
		return RelationalPrecedence
	case DoubleLessThan, DoubleGreaterThan:
		return ShiftPrecedence
	case Plus, Minus:
		return AdditivePrecedence
	case Asterisk, Slash, PercentSign:
		return MultiplicativePrecedence
	case Tilde:
		return UnaryPrecedence
	case Dot, DoubleColon:
		return MemberAccessPrecedence
	case LeftParenthesis, LeftSquareBracket:
		return HighestPrecedence
	default:
		return LowestPrecedence
	}
}

type Location struct {
	Row    uint64
	Column uint64
}

func (l *Location) Copy() Location {
	return Location{
		Row:    l.Row,
		Column: l.Column,
	}
}

func (l Location) String() string {
	return fmt.Sprintf("%d:%d", l.Row, l.Column)
}

type Token struct {
	Kind     TokenKind
	Location Location
	Value    string
}
