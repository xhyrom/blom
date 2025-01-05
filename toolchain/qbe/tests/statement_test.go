package tests

import (
	"blom/qbe"
	"testing"
)

func TestAssignStatement(t *testing.T) {
	name := qbe.TemporaryValue{Name: "name"}
	instr := qbe.AddInstruction{
		Left:  qbe.TemporaryValue{Name: "left"},
		Right: qbe.TemporaryValue{Name: "right"},
	}
	stmt := qbe.AssignStatement{Name: name, Type: qbe.Word, Instruction: instr}

	if stmt.StatementType() != qbe.AssignStatementType {
		t.Errorf("Expected AssignStatementType, got %v", stmt.StatementType())
	}

	expected := "%name =w add %left, %right"
	if stmt.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stmt.String())
	}
}

func TestVolatileStatement(t *testing.T) {
	instr := qbe.CallInstruction{
		Name: qbe.NewGlobalValue("func"),
		Parameters: []qbe.TypedValue{
			{Value: qbe.TemporaryValue{Name: "param1"}, Type: qbe.Word},
			{Value: qbe.TemporaryValue{Name: "param2"}, Type: qbe.Byte},
		},
	}
	stmt := qbe.VolatileStatement{Instruction: instr}

	if stmt.StatementType() != qbe.VolatileStatementType {
		t.Errorf("Expected VolatileStatementType, got %v", stmt.StatementType())
	}

	expected := "call $func(w %param1, w %param2)"
	if stmt.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stmt.String())
	}
}

func TestBlock(t *testing.T) {
	assignStmt := qbe.AssignStatement{
		Name: qbe.TemporaryValue{Name: "name"},
		Type: qbe.Word,
		Instruction: qbe.AddInstruction{
			Left:  qbe.TemporaryValue{Name: "left"},
			Right: qbe.TemporaryValue{Name: "right"},
		},
	}
	volatileStmt := qbe.VolatileStatement{
		Instruction: qbe.CallInstruction{
			Name: qbe.NewGlobalValue("func"),
			Parameters: []qbe.TypedValue{
				{Value: qbe.TemporaryValue{Name: "param1"}, Type: qbe.Word},
				{Value: qbe.TemporaryValue{Name: "param2"}, Type: qbe.Byte},
			},
		},
	}
	block := qbe.Block{
		Label:      "block_label",
		Statements: []qbe.Statement{assignStmt, volatileStmt},
	}

	expected := "@block_label\n\t%name =w add %left, %right\n\tcall $func(w %param1, w %param2)"
	result := block.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
