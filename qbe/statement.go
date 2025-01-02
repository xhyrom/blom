package qbe

import (
	"fmt"
)

type Block struct {
	Label      string
	Statements []Statement
}

func (b *Block) AddInstruction(instruction Instruction) {
	b.AddStatement(VolatileStatement{Instruction: instruction})
}

func (b *Block) AddAssign(name Value, t Type, instruction Instruction) {
	b.AddStatement(AssignStatement{Name: name, Type: t.IntoAbi(), Instruction: instruction})
}

func (b *Block) AddStatement(statement Statement) {
	b.Statements = append(b.Statements, statement)
}

func (b *Block) IsLastStatement(instruction InstructionType) bool {
	if len(b.Statements) == 0 {
		return false
	}

	lastStatement := b.Statements[len(b.Statements)-1]

	if lastStatement.StatementType() != VolatileStatementType {
		return false
	}

	vs := lastStatement.(VolatileStatement)
	return vs.Instruction.InstructionType() == instruction
}

func (b Block) String() string {
	result := fmt.Sprintf("@%s\n", b.Label)

	for i, stmt := range b.Statements {
		if i == len(b.Statements)-1 {
			result += fmt.Sprintf("\t%s", stmt)
		} else {
			result += fmt.Sprintf("\t%s\n", stmt)
		}
	}

	return result
}

type StatementType int

const (
	AssignStatementType StatementType = iota
	VolatileStatementType
)

type Statement interface {
	StatementType() StatementType
	String() string
}

// AssignStatement represents an assignment statement. ({name} ={type} {instruction})
type AssignStatement struct {
	Name        Value
	Type        Type
	Instruction Instruction
}

func (s AssignStatement) StatementType() StatementType {
	return AssignStatementType
}

func (s AssignStatement) String() string {
	return fmt.Sprintf("%s =%s %s", s.Name, s.Type, s.Instruction)
}

// VolatileStatement represents a volatile assignment statement. ({instruction})
type VolatileStatement struct {
	Instruction Instruction
}

func (s VolatileStatement) StatementType() StatementType {
	return VolatileStatementType
}

func (s VolatileStatement) String() string {
	return fmt.Sprintf("%s", s.Instruction)
}
