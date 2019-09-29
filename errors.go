// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "errors"

var (
	ErrInvalidAccessQualifier       = errors.New("invalid AccessQualifier value")
	ErrInvalidAddressingModel       = errors.New("invalid AddressingModel value")
	ErrInvalidDimensionality        = errors.New("invalid Dimensionality value")
	ErrInvalidExecutionMode         = errors.New("invalid ExecutionMode value")
	ErrInvalidExecutionModel        = errors.New("invalid ExecutionModel value")
	ErrInvalidFPFastMathMode        = errors.New("invalid FPFastMathMode value")
	ErrInvalidKernelProfilingInfo   = errors.New("invalid KernelProfilingInfo value")
	ErrInvalidKernelEnqueueFlag     = errors.New("invalid KernelEnqueueFlag value")
	ErrInvalidGroupOperation        = errors.New("invalid GroupOperation value")
	ErrInvalidExecutionScope        = errors.New("invalid ExecutionScope value")
	ErrInvalidMemoryAccess          = errors.New("invalid MemoryAccess value")
	ErrInvalidMemorySemantic        = errors.New("invalid MemorySemantic value")
	ErrInvalidFunctionControlMask   = errors.New("invalid FunctionControlMask value")
	ErrInvalidLoopControl           = errors.New("invalid LoopControl value")
	ErrInvalidSelectionControl      = errors.New("invalid SelectionControl value")
	ErrInvalidBuiltin               = errors.New("invalid Builtin value")
	ErrInvalidDecoration            = errors.New("invalid Decoration value")
	ErrInvalidFunctionParameter     = errors.New("invalid FunctionParameter value")
	ErrInvalidStorageClass          = errors.New("invalid StorageClass value")
	ErrInvalidSourceLanguage        = errors.New("invalid SourceLanguage value")
	ErrInvalidSamplerFilterMode     = errors.New("invalid SamplerFilterMode value")
	ErrInvalidSamplerAddressingMode = errors.New("invalid SamplerAddressingMode value")
	ErrInvalidMemoryModel           = errors.New("invalid MemoryModel value")
	ErrInvalidLinkageType           = errors.New("invalid LinkageType value")
	ErrInvalidFPRoundingMode        = errors.New("invalid FPRoundingMode value")
)
