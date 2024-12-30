package qbe

import (
	"fmt"
)

type InstructionType int

const (
	Add InstructionType = iota
	Subtract
	Multiply
	Divide
	Modulus
	BitwiseAnd
	BitwiseOr
	BitwiseXor
	BitwiseNot
	Negate
	Compare
	Copy
	Return
	JumpNonZero
	Jump
	Call
	VAArg
	VAStart
	Alloc8
	Store
	Load
	Conversion
	Extension
	Truncate
	ShiftLeft
	ArithmeticShiftRight
	Comment
)

type Instruction interface {
	InstructionType() InstructionType
}

// AddInstruction represents an add instruction. (add {left}, {right})
type AddInstruction struct {
	Left  Value
	Right Value
}

func NewAddInstruction(left Value, right Value) AddInstruction {
	return AddInstruction{
		Left:  left,
		Right: right,
	}
}

func (i AddInstruction) InstructionType() InstructionType {
	return Add
}

func (i AddInstruction) String() string {
	return fmt.Sprintf("add %s, %s", i.Left, i.Right)
}

// SubtractInstruction represents a subtract instruction. (sub {left}, {right})
type SubtractInstruction struct {
	Left  Value
	Right Value
}

func NewSubtractInstruction(left Value, right Value) SubtractInstruction {
	return SubtractInstruction{
		Left:  left,
		Right: right,
	}
}

func (i SubtractInstruction) InstructionType() InstructionType {
	return Subtract
}

func (i SubtractInstruction) String() string {
	return fmt.Sprintf("sub %s, %s", i.Left, i.Right)
}

// MultiplyInstruction represents a multiply instruction. (mul {left}, {right})
type MultiplyInstruction struct {
	Left  Value
	Right Value
}

func NewMultiplyInstruction(left Value, right Value) MultiplyInstruction {
	return MultiplyInstruction{
		Left:  left,
		Right: right,
	}
}

func (i MultiplyInstruction) InstructionType() InstructionType {
	return Multiply
}

func (i MultiplyInstruction) String() string {
	return fmt.Sprintf("mul %s, %s", i.Left, i.Right)
}

// DivideInstruction represents a divide instruction. (div {left}, {right})
type DivideInstruction struct {
	Left  Value
	Right Value
}

func NewDivideInstruction(left Value, right Value) DivideInstruction {
	return DivideInstruction{
		Left:  left,
		Right: right,
	}
}

func (i DivideInstruction) InstructionType() InstructionType {
	return Divide
}

func (i DivideInstruction) String() string {
	return fmt.Sprintf("div %s, %s", i.Left, i.Right)
}

// ModulusInstruction represents a modulus instruction. (rem {left}, {right})
type ModulusInstruction struct {
	Left  Value
	Right Value
}

func NewModulusInstruction(left Value, right Value) ModulusInstruction {
	return ModulusInstruction{
		Left:  left,
		Right: right,
	}
}

func (i ModulusInstruction) InstructionType() InstructionType {
	return Modulus
}

func (i ModulusInstruction) String() string {
	return fmt.Sprintf("rem %s, %s", i.Left, i.Right)
}

// BitwiseAndInstruction represents a bitwise and instruction. (and {left}, {right})
type BitwiseAndInstruction struct {
	Left  Value
	Right Value
}

func NewBitwiseAndInstruction(left Value, right Value) BitwiseAndInstruction {
	return BitwiseAndInstruction{
		Left:  left,
		Right: right,
	}
}

func (i BitwiseAndInstruction) InstructionType() InstructionType {
	return BitwiseAnd
}

func (i BitwiseAndInstruction) String() string {
	return fmt.Sprintf("and %s, %s", i.Left, i.Right)
}

// BitwiseOrInstruction represents a bitwise or instruction. (or {left}, {right})
type BitwiseOrInstruction struct {
	Left  Value
	Right Value
}

func NewBitwiseOrInstruction(left Value, right Value) BitwiseOrInstruction {
	return BitwiseOrInstruction{
		Left:  left,
		Right: right,
	}
}

func (i BitwiseOrInstruction) InstructionType() InstructionType {
	return BitwiseOr
}

func (i BitwiseOrInstruction) String() string {
	return fmt.Sprintf("or %s, %s", i.Left, i.Right)
}

// BitwiseXorInstruction represents a bitwise xor instruction. (xor {left}, {right})
type BitwiseXorInstruction struct {
	Left  Value
	Right Value
}

func NewBitwiseXorInstruction(left Value, right Value) BitwiseXorInstruction {
	return BitwiseXorInstruction{
		Left:  left,
		Right: right,
	}
}

func (i BitwiseXorInstruction) InstructionType() InstructionType {
	return BitwiseXor
}

func (i BitwiseXorInstruction) String() string {
	return fmt.Sprintf("xor %s, %s", i.Left, i.Right)
}

// BitwiseNotInstruction represents a bitwise not instruction. (xor {value}, -1)
type BitwiseNotInstruction struct {
	Value Value
}

func NewBitwiseNotInstruction(value Value) BitwiseNotInstruction {
	return BitwiseNotInstruction{
		Value: value,
	}
}

func (i BitwiseNotInstruction) InstructionType() InstructionType {
	return BitwiseNot
}

func (i BitwiseNotInstruction) String() string {
	return fmt.Sprintf("xor %s, -1", i.Value)
}

// NegateInstruction represents a negate instruction. (neg {value})
type NegateInstruction struct {
	Value Value
}

func NewNegateInstruction(value Value) NegateInstruction {
	return NegateInstruction{
		Value: value,
	}
}

func (i NegateInstruction) InstructionType() InstructionType {
	return Negate
}

func (i NegateInstruction) String() string {
	return fmt.Sprintf("neg %s", i.Value)
}

// CompareInstruction represents a compare instruction. ({comparison}{type} {left}, {right})
type CompareInstruction struct {
	Type       Type
	Comparison ComparisonType
	Left       Value
	Right      Value
}

func NewCompareInstruction(t Type, comp ComparisonType, left Value, right Value) CompareInstruction {
	return CompareInstruction{
		Type:       t,
		Comparison: comp,
		Left:       left,
		Right:      right,
	}
}

func (i CompareInstruction) InstructionType() InstructionType {
	return Compare
}

func (i CompareInstruction) String() string {
	base := "c"
	if !i.Type.IsFloatingPoint() && i.Type.IsSigned() {
		base += "s"
	}
	base += i.Comparison.String()

	return fmt.Sprintf("%s%s %s, %s", base, i.Type.IntoAbi(), i.Left, i.Right)
}

// CopyInstruction represents a copy instruction. (copy {value})
type CopyInstruction struct {
	Value Value
}

func (i CopyInstruction) InstructionType() InstructionType {
	return Copy
}

func (i CopyInstruction) String() string {
	return fmt.Sprintf("copy %s", i.Value)
}

// ReturnInstruction represents a return instruction. (ret {value})
type ReturnInstruction struct {
	Value Value
}

func (i ReturnInstruction) InstructionType() InstructionType {
	return Return
}

func (i ReturnInstruction) String() string {
	if i.Value != nil {
		return fmt.Sprintf("ret %s", i.Value)
	}

	return "ret"
}

// JumpNonZeroInstruction represents a jump non-zero instruction. (jnz {if_nonzero}, {if_zero})
type JumpNonZeroInstruction struct {
	Value     Value
	IfNonZero string
	IfZero    string
}

func (i JumpNonZeroInstruction) InstructionType() InstructionType {
	return JumpNonZero
}

func (i JumpNonZeroInstruction) String() string {
	return fmt.Sprintf("jnz %s, @%s, @%s", i.Value, i.IfNonZero, i.IfZero)
}

// JumpInstruction represents a jump instruction. (jmp {label})
type JumpInstruction struct {
	Label string
}

func (i JumpInstruction) InstructionType() InstructionType {
	return Jump
}

func (i JumpInstruction) String() string {
	return fmt.Sprintf("jmp @%s", i.Label)
}

// CallInstruction represents a call instruction. (call {name}({parameters}))
type CallInstruction struct {
	Name       Value
	Parameters []TypedValue
}

func NewCallInstruction(name Value, parameters ...TypedValue) CallInstruction {
	return CallInstruction{
		Name:       name,
		Parameters: parameters,
	}
}

func (i CallInstruction) InstructionType() InstructionType {
	return Call
}

func (i CallInstruction) String() string {
	params := ""
	for j, param := range i.Parameters {
		params += fmt.Sprintf("%s", param.AbiString())

		if j < len(i.Parameters)-1 {
			params += ", "
		}
	}

	return fmt.Sprintf("call %s(%s)", i.Name, params)
}

// VAArgInstruction represents a va_arg instruction. (vaarg {value})
type VAArgInstruction struct {
	Value Value
}

func (i VAArgInstruction) InstructionType() InstructionType {
	return VAArg
}

func (i VAArgInstruction) String() string {
	return fmt.Sprintf("vaarg %s", i.Value)
}

// VAStartInstruction represents a va_start instruction. (vastart {value})
type VAStartInstruction struct {
	Value Value
}

func (i VAStartInstruction) InstructionType() InstructionType {
	return VAStart
}

func (i VAStartInstruction) String() string {
	return fmt.Sprintf("vastart %s", i.Value)
}

// Alloc8Instruction represents an alloc8 instruction. (alloc8 {value})
type Alloc8Instruction struct {
	Value Value
}

func (i Alloc8Instruction) InstructionType() InstructionType {
	return Alloc8
}

func (i Alloc8Instruction) String() string {
	return fmt.Sprintf("alloc8 %s", i.Value)
}

// StoreInstruction represents a store instruction. (store{type} {value}, {destination})
type StoreInstruction struct {
	Type        Type
	Value       Value
	Destination Value
}

func NewStoreInstruction(t Type, value Value, destination Value) StoreInstruction {
	return StoreInstruction{
		Type:        t,
		Value:       value,
		Destination: destination,
	}
}

func (i StoreInstruction) InstructionType() InstructionType {
	return Store
}

func (i StoreInstruction) String() string {
	return fmt.Sprintf("store%s %s, %s", i.Type, i.Value, i.Destination)
}

// LoadInstruction represents a load instruction. (load{type} {source})
type LoadInstruction struct {
	Type   Type
	Source Value
}

func NewLoadInstruction(t Type, s Value) LoadInstruction {
	return LoadInstruction{
		Type:   t,
		Source: s,
	}
}

func (i LoadInstruction) InstructionType() InstructionType {
	return Load
}

func (i LoadInstruction) String() string {
	return fmt.Sprintf("load%s %s", i.Type, i.Source)
}

// ConversionInstruction represents a conversion instruction. ({from}to{to} {value})
type ConversionInstruction struct {
	From  Type
	To    Type
	Value Value
}

func (i ConversionInstruction) InstructionType() InstructionType {
	return Conversion
}

func (i ConversionInstruction) String() string {
	from := ""
	if i.From.IsFloatingPoint() {
		from = i.From.String()
	} else {
		if i.From.IsSigned() {
			from = "s"
		} else if i.From.IntoAbi().IsSigned() {
			from = "u"
		}

		from += i.From.IntoAbi().String()
	}

	to := ""
	if i.To.IsFloatingPoint() {
		to = "f"
	} else {
		if i.To.IsSigned() {
			to = "s"
		} else {
			to = "u"
		}

		to += "i"
	}

	return fmt.Sprintf("%sto%s %s", from, to, i.Value)
}

// ExtensionInstruction represents an extension instruction. (ext{type} {value})
type ExtensionInstruction struct {
	Type  Type
	Value Value
}

func (i ExtensionInstruction) InstructionType() InstructionType {
	return Extension
}

func (i ExtensionInstruction) String() string {
	typ := ""
	if i.Type.IsFloatingPoint() {
		typ = i.Type.String()
	} else {
		typ = "s" + i.Type.String()
	}

	return fmt.Sprintf("ext%s %s", typ, i.Value)
}

// TruncateInstruction represents a truncate instruction. (truncd {value})
type TruncateInstruction struct {
	Value Value
}

func (i TruncateInstruction) InstructionType() InstructionType {
	return Truncate
}

func (i TruncateInstruction) String() string {
	return fmt.Sprintf("truncd %s", i.Value)
}

// ShiftLeftInstruction represents a shift left instruction. (shl {value}, {shift})
type ShiftLeftInstruction struct {
	Value Value
	Shift Value
}

func NewShiftLeftInstruction(value Value, shift Value) ShiftLeftInstruction {
	return ShiftLeftInstruction{
		Value: value,
		Shift: shift,
	}
}

func (i ShiftLeftInstruction) InstructionType() InstructionType {
	return ShiftLeft
}

func (i ShiftLeftInstruction) String() string {
	return fmt.Sprintf("shl %s, %s", i.Value, i.Shift)
}

// ArithmeticShiftRightInstruction represents an arithmetic shift right instruction. (sar {value}, {shift})
type ArithmeticShiftRightInstruction struct {
	Value Value
	Shift Value
}

func NewArithmeticShiftRightInstruction(value Value, shift Value) ArithmeticShiftRightInstruction {
	return ArithmeticShiftRightInstruction{
		Value: value,
		Shift: shift,
	}
}

func (i ArithmeticShiftRightInstruction) InstructionType() InstructionType {
	return ArithmeticShiftRight
}

func (i ArithmeticShiftRightInstruction) String() string {
	return fmt.Sprintf("sar %s, %s", i.Value, i.Shift)
}

// CommentInstruction represents a comment instruction. (# {comment})
type CommentInstruction struct {
	Comment string
}

func (i CommentInstruction) InstructionType() InstructionType {
	return Comment
}

func (i CommentInstruction) String() string {
	return fmt.Sprintf("# %s", i.Comment)
}
