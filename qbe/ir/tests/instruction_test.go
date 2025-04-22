package tests

import (
	"testing"

	"github.com/xhyrom/blom/qbe/ir"
)

func TestAddInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.AddInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.Add {
		t.Errorf("Expected Add, got %v", instr.InstructionType())
	}

	expected := "add %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestSubtractInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.SubtractInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.Subtract {
		t.Errorf("Expected Subtract, got %v", instr.InstructionType())
	}

	expected := "sub %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestMultiplyInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.MultiplyInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.Multiply {
		t.Errorf("Expected Multiply, got %v", instr.InstructionType())
	}

	expected := "mul %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestDivideInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.DivideInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.Divide {
		t.Errorf("Expected Divide, got %v", instr.InstructionType())
	}

	expected := "div %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestModulusInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.ModulusInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.Modulus {
		t.Errorf("Expected Modulus, got %v", instr.InstructionType())
	}

	expected := "rem %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseAndInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.BitwiseAndInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.BitwiseAnd {
		t.Errorf("Expected BitwiseAnd, got %v", instr.InstructionType())
	}

	expected := "and %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseOrInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.BitwiseOrInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.BitwiseOr {
		t.Errorf("Expected BitwiseOr, got %v", instr.InstructionType())
	}

	expected := "or %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseXorInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.BitwiseXorInstruction{Left: left, Right: right}

	if instr.InstructionType() != ir.BitwiseXor {
		t.Errorf("Expected BitwiseXor, got %v", instr.InstructionType())
	}

	expected := "xor %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseNotInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.BitwiseNotInstruction{Value: value}

	if instr.InstructionType() != ir.BitwiseNot {
		t.Errorf("Expected BitwiseNot, got %v", instr.InstructionType())
	}

	expected := "xor %value, -1"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestNegateInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.NegateInstruction{Value: value}

	if instr.InstructionType() != ir.Negate {
		t.Errorf("Expected Negate, got %v", instr.InstructionType())
	}

	expected := "neg %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCompareInstruction(t *testing.T) {
	left := ir.TemporaryValue{Name: "left"}
	right := ir.TemporaryValue{Name: "right"}
	instr := ir.CompareInstruction{
		Type:       ir.Word,
		Comparison: ir.Equal,
		Left:       left,
		Right:      right,
	}

	if instr.InstructionType() != ir.Compare {
		t.Errorf("Expected Compare, got %v", instr.InstructionType())
	}

	expected := "ceqw %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCopyInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.CopyInstruction{Value: value}

	if instr.InstructionType() != ir.Copy {
		t.Errorf("Expected Copy, got %v", instr.InstructionType())
	}

	expected := "copy %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestReturnInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ReturnInstruction{Value: &value}

	if instr.InstructionType() != ir.Return {
		t.Errorf("Expected Return, got %v", instr.InstructionType())
	}

	expected := "ret %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}

	instr = ir.ReturnInstruction{Value: nil}
	expected = "ret"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestJumpNonZeroInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.JumpNonZeroInstruction{
		Value:     value,
		IfNonZero: "nonzero",
		IfZero:    "zero",
	}

	if instr.InstructionType() != ir.JumpNonZero {
		t.Errorf("Expected JumpNonZero, got %v", instr.InstructionType())
	}

	expected := "jnz %value, @nonzero, @zero"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestJumpInstruction(t *testing.T) {
	instr := ir.JumpInstruction{Label: "label"}

	if instr.InstructionType() != ir.Jump {
		t.Errorf("Expected Jump, got %v", instr.InstructionType())
	}

	expected := "jmp @label"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCallInstruction(t *testing.T) {
	params := []ir.TypedValue{
		{Value: ir.TemporaryValue{Name: "param1"}, Type: ir.Word},
		{Value: ir.TemporaryValue{Name: "param2"}, Type: ir.Byte},
	}
	instr := ir.CallInstruction{Name: ir.NewGlobalValue("func"), Parameters: params}

	if instr.InstructionType() != ir.Call {
		t.Errorf("Expected Call, got %v", instr.InstructionType())
	}

	expected := "call $func(w %param1, w %param2)"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestVAArgInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.VAArgInstruction{Value: value}

	if instr.InstructionType() != ir.VAArg {
		t.Errorf("Expected VAArg, got %v", instr.InstructionType())
	}

	expected := "vaarg %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestVAStartInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.VAStartInstruction{Value: value}

	if instr.InstructionType() != ir.VAStart {
		t.Errorf("Expected VAStart, got %v", instr.InstructionType())
	}

	expected := "vastart %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestAlloc8Instruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.Alloc8Instruction{Value: value}

	if instr.InstructionType() != ir.Alloc8 {
		t.Errorf("Expected Alloc8, got %v", instr.InstructionType())
	}

	expected := "alloc8 %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestStoreInstruction(t *testing.T) {
	dest := ir.TemporaryValue{Name: "dest"}
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.StoreInstruction{Type: ir.Word, Destination: dest, Value: value}

	if instr.InstructionType() != ir.Store {
		t.Errorf("Expected Store, got %v", instr.InstructionType())
	}

	expected := "storew %value, %dest"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestLoadInstruction(t *testing.T) {
	source := ir.TemporaryValue{Name: "source"}
	instr := ir.LoadInstruction{Type: ir.Word, Source: source}

	if instr.InstructionType() != ir.Load {
		t.Errorf("Expected Load, got %v", instr.InstructionType())
	}

	expected := "loadw %source"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSingleToSignedInteger(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Single, To: ir.Word, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "stosi %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSingleToUnsignedInteger(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Single, To: ir.UnsignedWord, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "stoui %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionDoubleToUnsignedInteger(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Double, To: ir.UnsignedWord, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "dtoui %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedWordToSingle(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Word, To: ir.Single, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "swtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedWordToDouble(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Word, To: ir.Double, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "swtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedWordToSingle(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.UnsignedWord, To: ir.Single, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "uwtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedWordToDouble(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.UnsignedWord, To: ir.Double, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "uwtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedLongToSingle(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Long, To: ir.Single, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "sltof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedLongToDouble(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.Long, To: ir.Double, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "sltof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedLongToSingle(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.UnsignedLong, To: ir.Single, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "ultof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedLongToDouble(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ConversionInstruction{From: ir.UnsignedLong, To: ir.Double, Value: value}

	if instr.InstructionType() != ir.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "ultof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestExtensionInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.ExtensionInstruction{Type: ir.Word, Value: value}

	if instr.InstructionType() != ir.Extension {
		t.Errorf("Expected Extension, got %v", instr.InstructionType())
	}

	expected := "extsw %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestTruncateInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	instr := ir.TruncateInstruction{Value: value}

	if instr.InstructionType() != ir.Truncate {
		t.Errorf("Expected Truncate, got %v", instr.InstructionType())
	}

	expected := "truncd %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestShiftLeftInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	shift := ir.TemporaryValue{Name: "shift"}
	instr := ir.ShiftLeftInstruction{Value: value, Shift: shift}

	if instr.InstructionType() != ir.ShiftLeft {
		t.Errorf("Expected ShiftLeft, got %v", instr.InstructionType())
	}

	expected := "shl %value, %shift"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestArithmeticShiftRightInstruction(t *testing.T) {
	value := ir.TemporaryValue{Name: "value"}
	shift := ir.TemporaryValue{Name: "shift"}
	instr := ir.ArithmeticShiftRightInstruction{Value: value, Shift: shift}

	if instr.InstructionType() != ir.ArithmeticShiftRight {
		t.Errorf("Expected ArithmeticShiftRight, got %v", instr.InstructionType())
	}

	expected := "sar %value, %shift"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCommentInstruction(t *testing.T) {
	instr := ir.CommentInstruction{Comment: "This is a comment"}

	if instr.InstructionType() != ir.Comment {
		t.Errorf("Expected Comment, got %v", instr.InstructionType())
	}

	expected := "# This is a comment"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}
