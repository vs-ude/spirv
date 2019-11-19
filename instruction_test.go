// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestInstructionResultId(t *testing.T) {
	var instr Instruction
	instr = &OpAtomicXor{1, 2, 3, 4, 5, 6}
	want := instr.(*OpAtomicXor).ResultId
	have, ok := InstructionResultId(instr)
	if have != want {
		t.Fatalf("ResultId mismatch:\nHave: %d\nWant: %d", have, want)
	}

	instr = &OpLoopMerge{1, 2, 3, []uint32{}}
	_, ok = InstructionResultId(instr)
	if ok {
		t.Fatalf("Expected failure for %+v", instr)
	}
}

func TestSetInstructionResultId(t *testing.T) {
	var instr Instruction
	want := Id(3)
	instr = &OpAtomicXor{1, 2, 3, 4, 5, 6}
	SetInstructionResultId(instr, want)
	have := instr.(*OpAtomicXor).ResultId
	if have != want {
		t.Fatalf("ResultId mismatch:\nHave: %d\nWant: %d", have, want)
	}

	instr = &OpLoopMerge{1, 2, 3, []uint32{}}
	ok := SetInstructionResultId(instr, 0)
	if ok {
		t.Fatalf("Expected failure for %+v", instr)
	}
}

type InstructionEqualsTest struct {
	a, b Instruction
	want, wantWithout bool
}

func TestInstructionEquals(t *testing.T) {
	for i, st := range []InstructionEqualsTest{
		{&OpAtomicXor{1, 2, 3, 4, 5, 6}, &OpAtomicXor{1, 2, 3, 4, 5, 6}, true, true},
		{&OpAtomicXor{1, 2, 3, 4, 5, 6}, &OpAtomicXor{1, 12, 3, 4, 5, 6}, false, true},
		{&OpAtomicXor{1, 2, 3, 4, 5, 6}, &OpAtomicAnd{1, 2, 3, 4, 5, 6}, false, false},
		{&OpAtomicXor{1, 2, 3, 4, 5, 6}, &OpGetDefaultQueue{0, 0}, false, false},
	} {
		have := InstructionEquals(st.a, st.b, true)
		if have != st.want {
			t.Fatalf("case %d: result mismatch: %+v; %+v; true\nHave: %t\nWant: %t",
				i, st.a, st.b, have, st.want)
		}
		have = InstructionEquals(st.a, st.b, false)
		if have != st.wantWithout {
			t.Fatalf("case %d: result mismatch: %+v; %+v; false\nHave: %t\nWant: %t",
				i, st.a, st.b, have, st.wantWithout)
		}
	}
}

func TestInstructionName(t *testing.T) {
	var instr Instruction = &OpAtomicXor{1, 2, 3, 4, 5, 6}

	want := "OpAtomicXor"
	have := InstructionName(instr)
	if want != have {
		t.Fatalf("result mismatch\nHave: %s\nWant: %s", have, want)
	}
}
