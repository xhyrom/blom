package qbe

import "fmt"

type Block struct {
	Label      string
	Statements []Statement
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
