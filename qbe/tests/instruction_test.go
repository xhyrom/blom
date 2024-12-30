package tests

import (
	"blom/qbe"
	"testing"
)

func TestAddInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.AddInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.Add {
		t.Errorf("Expected Add, got %v", instr.InstructionType())
	}

	expected := "add %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestSubtractInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.SubtractInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.Subtract {
		t.Errorf("Expected Subtract, got %v", instr.InstructionType())
	}

	expected := "sub %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestMultiplyInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.MultiplyInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.Multiply {
		t.Errorf("Expected Multiply, got %v", instr.InstructionType())
	}

	expected := "mul %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestDivideInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.DivideInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.Divide {
		t.Errorf("Expected Divide, got %v", instr.InstructionType())
	}

	expected := "div %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestModulusInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.ModulusInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.Modulus {
		t.Errorf("Expected Modulus, got %v", instr.InstructionType())
	}

	expected := "rem %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseAndInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.BitwiseAndInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.BitwiseAnd {
		t.Errorf("Expected BitwiseAnd, got %v", instr.InstructionType())
	}

	expected := "and %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseOrInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.BitwiseOrInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.BitwiseOr {
		t.Errorf("Expected BitwiseOr, got %v", instr.InstructionType())
	}

	expected := "or %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseXorInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.BitwiseXorInstruction{Left: left, Right: right}

	if instr.InstructionType() != qbe.BitwiseXor {
		t.Errorf("Expected BitwiseXor, got %v", instr.InstructionType())
	}

	expected := "xor %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestBitwiseNotInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.BitwiseNotInstruction{Value: value}

	if instr.InstructionType() != qbe.BitwiseNot {
		t.Errorf("Expected BitwiseNot, got %v", instr.InstructionType())
	}

	expected := "xor %value, -1"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestNegateInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.NegateInstruction{Value: value}

	if instr.InstructionType() != qbe.Negate {
		t.Errorf("Expected Negate, got %v", instr.InstructionType())
	}

	expected := "neg %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCompareInstruction(t *testing.T) {
	left := qbe.TemporaryValue{Name: "left"}
	right := qbe.TemporaryValue{Name: "right"}
	instr := qbe.CompareInstruction{
		Type:       qbe.Word,
		Comparison: qbe.Equal,
		Left:       left,
		Right:      right,
	}

	if instr.InstructionType() != qbe.Compare {
		t.Errorf("Expected Compare, got %v", instr.InstructionType())
	}

	expected := "cseqw %left, %right"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCopyInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.CopyInstruction{Value: value}

	if instr.InstructionType() != qbe.Copy {
		t.Errorf("Expected Copy, got %v", instr.InstructionType())
	}

	expected := "copy %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestReturnInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ReturnInstruction{Value: &value}

	if instr.InstructionType() != qbe.Return {
		t.Errorf("Expected Return, got %v", instr.InstructionType())
	}

	expected := "ret %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}

	instr = qbe.ReturnInstruction{Value: nil}
	expected = "ret"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestJumpNonZeroInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.JumpNonZeroInstruction{
		Value:     value,
		IfNonZero: "nonzero",
		IfZero:    "zero",
	}

	if instr.InstructionType() != qbe.JumpNonZero {
		t.Errorf("Expected JumpNonZero, got %v", instr.InstructionType())
	}

	expected := "jnz %value, @nonzero, @zero"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestJumpInstruction(t *testing.T) {
	instr := qbe.JumpInstruction{Label: "label"}

	if instr.InstructionType() != qbe.Jump {
		t.Errorf("Expected Jump, got %v", instr.InstructionType())
	}

	expected := "jmp @label"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCallInstruction(t *testing.T) {
	params := []qbe.TypedValue{
		{Value: qbe.TemporaryValue{Name: "param1"}, Type: qbe.Word},
		{Value: qbe.TemporaryValue{Name: "param2"}, Type: qbe.Byte},
	}
	instr := qbe.CallInstruction{Name: qbe.NewGlobalValue("func"), Parameters: params}

	if instr.InstructionType() != qbe.Call {
		t.Errorf("Expected Call, got %v", instr.InstructionType())
	}

	expected := "call $func(w %param1, w %param2)"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestVAArgInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.VAArgInstruction{Value: value}

	if instr.InstructionType() != qbe.VAArg {
		t.Errorf("Expected VAArg, got %v", instr.InstructionType())
	}

	expected := "vaarg %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestVAStartInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.VAStartInstruction{Value: value}

	if instr.InstructionType() != qbe.VAStart {
		t.Errorf("Expected VAStart, got %v", instr.InstructionType())
	}

	expected := "vastart %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestAlloc8Instruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.Alloc8Instruction{Value: value}

	if instr.InstructionType() != qbe.Alloc8 {
		t.Errorf("Expected Alloc8, got %v", instr.InstructionType())
	}

	expected := "alloc8 %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestStoreInstruction(t *testing.T) {
	dest := qbe.TemporaryValue{Name: "dest"}
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.StoreInstruction{Type: qbe.Word, Destination: dest, Value: value}

	if instr.InstructionType() != qbe.Store {
		t.Errorf("Expected Store, got %v", instr.InstructionType())
	}

	expected := "storew %dest, %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestLoadInstruction(t *testing.T) {
	source := qbe.TemporaryValue{Name: "source"}
	instr := qbe.LoadInstruction{Type: qbe.Word, Source: source}

	if instr.InstructionType() != qbe.Load {
		t.Errorf("Expected Load, got %v", instr.InstructionType())
	}

	expected := "loadw %source"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSingleToSignedInteger(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Single, To: qbe.Word, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "stosi %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSingleToUnsignedInteger(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Single, To: qbe.UnsignedWord, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "stoui %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionDoubleToUnsignedInteger(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Double, To: qbe.UnsignedWord, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "dtoui %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedWordToSingle(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Word, To: qbe.Single, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "swtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedWordToDouble(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Word, To: qbe.Double, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "swtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedWordToSingle(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.UnsignedWord, To: qbe.Single, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "uwtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedWordToDouble(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.UnsignedWord, To: qbe.Double, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "uwtof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedLongToSingle(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Long, To: qbe.Single, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "sltof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionSignedLongToDouble(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.Long, To: qbe.Double, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "sltof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedLongToSingle(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.UnsignedLong, To: qbe.Single, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "ultof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestConversionInstructionUnsignedLongToDouble(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ConversionInstruction{From: qbe.UnsignedLong, To: qbe.Double, Value: value}

	if instr.InstructionType() != qbe.Conversion {
		t.Errorf("Expected Conversion, got %v", instr.InstructionType())
	}

	expected := "ultof %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestExtensionInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.ExtensionInstruction{Type: qbe.Word, Value: value}

	if instr.InstructionType() != qbe.Extension {
		t.Errorf("Expected Extension, got %v", instr.InstructionType())
	}

	expected := "extsw %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestTruncateInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	instr := qbe.TruncateInstruction{Value: value}

	if instr.InstructionType() != qbe.Truncate {
		t.Errorf("Expected Truncate, got %v", instr.InstructionType())
	}

	expected := "truncd %value"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestShiftLeftInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	shift := qbe.TemporaryValue{Name: "shift"}
	instr := qbe.ShiftLeftInstruction{Value: value, Shift: shift}

	if instr.InstructionType() != qbe.ShiftLeft {
		t.Errorf("Expected ShiftLeft, got %v", instr.InstructionType())
	}

	expected := "shl %value, %shift"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestArithmeticShiftRightInstruction(t *testing.T) {
	value := qbe.TemporaryValue{Name: "value"}
	shift := qbe.TemporaryValue{Name: "shift"}
	instr := qbe.ArithmeticShiftRightInstruction{Value: value, Shift: shift}

	if instr.InstructionType() != qbe.ArithmeticShiftRight {
		t.Errorf("Expected ArithmeticShiftRight, got %v", instr.InstructionType())
	}

	expected := "sar %value, %shift"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}

func TestCommentInstruction(t *testing.T) {
	instr := qbe.CommentInstruction{Comment: "This is a comment"}

	if instr.InstructionType() != qbe.Comment {
		t.Errorf("Expected Comment, got %v", instr.InstructionType())
	}

	expected := "# This is a comment"
	if instr.String() != expected {
		t.Errorf("Expected %s, got %s", expected, instr.String())
	}
}
