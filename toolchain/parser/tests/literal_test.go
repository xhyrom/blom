package tests

import (
	"blom/ast"
	"blom/parser"
	"blom/tokens"
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
)

func TestIdentifierLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "i32;")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.IdentifierLiteral{
				Value: "i32",
				Loc: tokens.Location{
					Row:    1,
					Column: 3,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}

func TestCharLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "'a';")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.CharLiteral{
				Value: 'a',
				Loc: tokens.Location{
					Row:    1,
					Column: 3,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}

func TestStringLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "\"a\";")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.StringLiteral{
				Value: "a",
				Loc: tokens.Location{
					Row:    1,
					Column: 3,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}

func TestIntLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "1;")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.IntLiteral{
				Value: 1,
				Loc: tokens.Location{
					Row:    1,
					Column: 1,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}

func TestFloatLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "1.5;")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.FloatLiteral{
				Value: 1.5,
				Loc: tokens.Location{
					Row:    1,
					Column: 3,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}

func TestBooleanTrueLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "true;")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.BooleanLiteral{
				Value: true,
				Loc: tokens.Location{
					Row:    1,
					Column: 4,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}

func TestBooleanFalseLiteral(t *testing.T) {
	parser := parser.New("test.blom")
	program := *parser.AST("test.blom", "false;")
	expected := ast.Program{
		Body: []ast.Node{
			&ast.BooleanLiteral{
				Value: false,
				Loc: tokens.Location{
					Row:    1,
					Column: 5,
				},
			},
		},
	}

	if !reflect.DeepEqual(program, expected) {
		t.Errorf("expected %v, got %v", dump.Format(expected), dump.Format(program))
	}
}
