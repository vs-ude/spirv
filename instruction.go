// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"reflect"
)

// Instruction defines a generic instruction.
type Instruction interface {
	Verifiable

	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32

	// Optional returns true if the instruction has no semantic meaning.
	// Its presence is mostly for debugging purposes.
	Optional() bool
}

// instructionResultId returns the value of the instruction's result Id,
// provided it defines one.
func instructionResultId(i Instruction) (Id, bool) {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)

	field := rv.FieldByName("ResultId")
	if field.Kind() == reflect.Invalid {
		return 0, false
	}

	id := field.Interface().(Id)
	return id, true
}

// instructionName returns the name for the given instruction.
// This is the type name, minus some package cruft.
func instructionName(i Instruction) string {
	name := fmt.Sprintf("%T", i)
	return name[len("*spirv."):]
}
