package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestAssignStatement(t *testing.T) {
	name := ir.TemporaryValue{Name: "name"}
	instr := ir.AddInstruction{
		Left:  ir.TemporaryValue{Name: "left"},
		Right: ir.TemporaryValue{Name: "right"},
	}
	stmt := ir.AssignStatement{Name: name, Type: ir.Word, Instruction: instr}

	if stmt.StatementType() != ir.AssignStatementType {
		t.Errorf("Expected AssignStatementType, got %v", stmt.StatementType())
	}

	expected := "%name =w add %left, %right"
	if stmt.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stmt.String())
	}
}

func TestVolatileStatement(t *testing.T) {
	instr := ir.CallInstruction{
		Name: ir.NewGlobalValue("func"),
		Parameters: []ir.TypedValue{
			{Value: ir.TemporaryValue{Name: "param1"}, Type: ir.Word},
			{Value: ir.TemporaryValue{Name: "param2"}, Type: ir.Byte},
		},
	}
	stmt := ir.VolatileStatement{Instruction: instr}

	if stmt.StatementType() != ir.VolatileStatementType {
		t.Errorf("Expected VolatileStatementType, got %v", stmt.StatementType())
	}

	expected := "call $func(w %param1, w %param2)"
	if stmt.String() != expected {
		t.Errorf("Expected %s, got %s", expected, stmt.String())
	}
}

func TestBlock(t *testing.T) {
	assignStmt := ir.AssignStatement{
		Name: ir.TemporaryValue{Name: "name"},
		Type: ir.Word,
		Instruction: ir.AddInstruction{
			Left:  ir.TemporaryValue{Name: "left"},
			Right: ir.TemporaryValue{Name: "right"},
		},
	}
	volatileStmt := ir.VolatileStatement{
		Instruction: ir.CallInstruction{
			Name: ir.NewGlobalValue("func"),
			Parameters: []ir.TypedValue{
				{Value: ir.TemporaryValue{Name: "param1"}, Type: ir.Word},
				{Value: ir.TemporaryValue{Name: "param2"}, Type: ir.Byte},
			},
		},
	}
	block := ir.Block{
		Label:      "block_label",
		Statements: []ir.Statement{assignStmt, volatileStmt},
	}

	expected := "@block_label\n\t%name =w add %left, %right\n\tcall $func(w %param1, w %param2)"
	result := block.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
