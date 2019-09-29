// This file was generated. DO NOT EDIT!

package spirv

type OpNop struct {
}

func (op *OpNop) Opcode() uint32 { return opcodeNop }
func (op *OpNop) Optional() bool { return false }
func (op *OpNop) Verify() error  { return nil }

type OpUndef struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpUndef) Opcode() uint32 { return opcodeUndef }
func (op *OpUndef) Optional() bool { return false }
func (op *OpUndef) Verify() error  { return nil }

type OpSourceContinued struct {
	ContinuedSource String
}

func (op *OpSourceContinued) Opcode() uint32 { return opcodeSourceContinued }
func (op *OpSourceContinued) Optional() bool { return true }
func (op *OpSourceContinued) Verify() error  { return nil }

type OpSource struct {
	SourceLanguage SourceLanguage
	Version        uint32
	File           Id     `spirv:"optional"`
	Source         String `spirv:"optional"`
}

func (op *OpSource) Opcode() uint32 { return opcodeSource }
func (op *OpSource) Optional() bool { return true }
func (op *OpSource) Verify() error  { return nil }

type OpSourceExtension struct {
	Extension String
}

func (op *OpSourceExtension) Opcode() uint32 { return opcodeSourceExtension }
func (op *OpSourceExtension) Optional() bool { return true }
func (op *OpSourceExtension) Verify() error  { return nil }

type OpName struct {
	Target Id
	Name   String
}

func (op *OpName) Opcode() uint32 { return opcodeName }
func (op *OpName) Optional() bool { return true }
func (op *OpName) Verify() error  { return nil }

type OpMemberName struct {
	Type   Id
	Member uint32
	Name   String
}

func (op *OpMemberName) Opcode() uint32 { return opcodeMemberName }
func (op *OpMemberName) Optional() bool { return true }
func (op *OpMemberName) Verify() error  { return nil }

type OpString struct {
	IdResult Id
	String   String
}

func (op *OpString) Opcode() uint32 { return opcodeString }
func (op *OpString) Optional() bool { return true }
func (op *OpString) Verify() error  { return nil }

type OpLine struct {
	File   Id
	Line   uint32
	Column uint32
}

func (op *OpLine) Opcode() uint32 { return opcodeLine }
func (op *OpLine) Optional() bool { return true }
func (op *OpLine) Verify() error  { return nil }

type OpExtension struct {
	Name String
}

func (op *OpExtension) Opcode() uint32 { return opcodeExtension }
func (op *OpExtension) Optional() bool { return false }
func (op *OpExtension) Verify() error  { return nil }

type OpExtInstImport struct {
	IdResult Id
	Name     String
}

func (op *OpExtInstImport) Opcode() uint32 { return opcodeExtInstImport }
func (op *OpExtInstImport) Optional() bool { return false }
func (op *OpExtInstImport) Verify() error  { return nil }

type OpExtInst struct {
	IdResultType Id
	IdResult     Id
	Set          Id
	Instruction  uint32
	Argv         []Id `spirv:"optional"`
}

func (op *OpExtInst) Opcode() uint32 { return opcodeExtInst }
func (op *OpExtInst) Optional() bool { return false }
func (op *OpExtInst) Verify() error  { return nil }

type OpMemoryModel struct {
	AddressingModel AddressingModel
	MemoryModel     MemoryModel
}

func (op *OpMemoryModel) Opcode() uint32 { return opcodeMemoryModel }
func (op *OpMemoryModel) Optional() bool { return false }
func (op *OpMemoryModel) Verify() error  { return nil }

type OpEntryPoint struct {
	ExecutionModel ExecutionModel
	EntryPoint     Id
	Name           String
	Interface      []Id `spirv:"optional"`
}

func (op *OpEntryPoint) Opcode() uint32 { return opcodeEntryPoint }
func (op *OpEntryPoint) Optional() bool { return false }
func (op *OpEntryPoint) Verify() error  { return nil }

type OpExecutionMode struct {
	EntryPoint Id
	Mode       ExecutionMode
	Argv       []uint32
}

func (op *OpExecutionMode) Opcode() uint32 { return opcodeExecutionMode }
func (op *OpExecutionMode) Optional() bool { return false }
func (op *OpExecutionMode) Verify() error  { return nil }

type OpCapability struct {
	Capability Capability
}

func (op *OpCapability) Opcode() uint32 { return opcodeCapability }
func (op *OpCapability) Optional() bool { return false }
func (op *OpCapability) Verify() error  { return nil }

type OpTypeVoid struct {
	IdResult Id
}

func (op *OpTypeVoid) Opcode() uint32 { return opcodeTypeVoid }
func (op *OpTypeVoid) Optional() bool { return false }
func (op *OpTypeVoid) Verify() error  { return nil }

type OpTypeBool struct {
	IdResult Id
}

func (op *OpTypeBool) Opcode() uint32 { return opcodeTypeBool }
func (op *OpTypeBool) Optional() bool { return false }
func (op *OpTypeBool) Verify() error  { return nil }

type OpTypeInt struct {
	IdResult   Id
	Width      uint32
	Signedness uint32
}

func (op *OpTypeInt) Opcode() uint32 { return opcodeTypeInt }
func (op *OpTypeInt) Optional() bool { return false }
func (op *OpTypeInt) Verify() error  { return nil }

type OpTypeFloat struct {
	IdResult Id
	Width    uint32
}

func (op *OpTypeFloat) Opcode() uint32 { return opcodeTypeFloat }
func (op *OpTypeFloat) Optional() bool { return false }
func (op *OpTypeFloat) Verify() error  { return nil }

type OpTypeVector struct {
	IdResult       Id
	ComponentType  Id
	ComponentCount uint32
}

func (op *OpTypeVector) Opcode() uint32 { return opcodeTypeVector }
func (op *OpTypeVector) Optional() bool { return false }
func (op *OpTypeVector) Verify() error  { return nil }

type OpTypeMatrix struct {
	IdResult    Id
	ColumnType  Id
	ColumnCount uint32
}

func (op *OpTypeMatrix) Opcode() uint32 { return opcodeTypeMatrix }
func (op *OpTypeMatrix) Optional() bool { return false }
func (op *OpTypeMatrix) Verify() error  { return nil }

type OpTypeImage struct {
	IdResult        Id
	SampledType     Id
	Dim             Dim
	Depth           uint32
	Arrayed         uint32
	MS              uint32
	Sampled         uint32
	ImageFormat     ImageFormat
	AccessQualifier AccessQualifier `spirv:"optional"`
}

func (op *OpTypeImage) Opcode() uint32 { return opcodeTypeImage }
func (op *OpTypeImage) Optional() bool { return false }
func (op *OpTypeImage) Verify() error  { return nil }

type OpTypeSampler struct {
	IdResult Id
}

func (op *OpTypeSampler) Opcode() uint32 { return opcodeTypeSampler }
func (op *OpTypeSampler) Optional() bool { return false }
func (op *OpTypeSampler) Verify() error  { return nil }

type OpTypeSampledImage struct {
	IdResult  Id
	ImageType Id
}

func (op *OpTypeSampledImage) Opcode() uint32 { return opcodeTypeSampledImage }
func (op *OpTypeSampledImage) Optional() bool { return false }
func (op *OpTypeSampledImage) Verify() error  { return nil }

type OpTypeArray struct {
	IdResult    Id
	ElementType Id
	Length      Id
}

func (op *OpTypeArray) Opcode() uint32 { return opcodeTypeArray }
func (op *OpTypeArray) Optional() bool { return false }
func (op *OpTypeArray) Verify() error  { return nil }

type OpTypeRuntimeArray struct {
	IdResult    Id
	ElementType Id
}

func (op *OpTypeRuntimeArray) Opcode() uint32 { return opcodeTypeRuntimeArray }
func (op *OpTypeRuntimeArray) Optional() bool { return false }
func (op *OpTypeRuntimeArray) Verify() error  { return nil }

type OpTypeStruct struct {
	IdResult Id
	Argv     []Id `spirv:"optional"`
}

func (op *OpTypeStruct) Opcode() uint32 { return opcodeTypeStruct }
func (op *OpTypeStruct) Optional() bool { return false }
func (op *OpTypeStruct) Verify() error  { return nil }

type OpTypeOpaque struct {
	IdResult                Id
	TheNameOfTheOpaqueType_ String
}

func (op *OpTypeOpaque) Opcode() uint32 { return opcodeTypeOpaque }
func (op *OpTypeOpaque) Optional() bool { return false }
func (op *OpTypeOpaque) Verify() error  { return nil }

type OpTypePointer struct {
	IdResult     Id
	StorageClass StorageClass
	Type         Id
}

func (op *OpTypePointer) Opcode() uint32 { return opcodeTypePointer }
func (op *OpTypePointer) Optional() bool { return false }
func (op *OpTypePointer) Verify() error  { return nil }

type OpTypeFunction struct {
	IdResult   Id
	ReturnType Id
	Argv       []Id `spirv:"optional"`
}

func (op *OpTypeFunction) Opcode() uint32 { return opcodeTypeFunction }
func (op *OpTypeFunction) Optional() bool { return false }
func (op *OpTypeFunction) Verify() error  { return nil }

type OpTypeEvent struct {
	IdResult Id
}

func (op *OpTypeEvent) Opcode() uint32 { return opcodeTypeEvent }
func (op *OpTypeEvent) Optional() bool { return false }
func (op *OpTypeEvent) Verify() error  { return nil }

type OpTypeDeviceEvent struct {
	IdResult Id
}

func (op *OpTypeDeviceEvent) Opcode() uint32 { return opcodeTypeDeviceEvent }
func (op *OpTypeDeviceEvent) Optional() bool { return false }
func (op *OpTypeDeviceEvent) Verify() error  { return nil }

type OpTypeReserveId struct {
	IdResult Id
}

func (op *OpTypeReserveId) Opcode() uint32 { return opcodeTypeReserveId }
func (op *OpTypeReserveId) Optional() bool { return false }
func (op *OpTypeReserveId) Verify() error  { return nil }

type OpTypeQueue struct {
	IdResult Id
}

func (op *OpTypeQueue) Opcode() uint32 { return opcodeTypeQueue }
func (op *OpTypeQueue) Optional() bool { return false }
func (op *OpTypeQueue) Verify() error  { return nil }

type OpTypePipe struct {
	IdResult  Id
	Qualifier AccessQualifier
}

func (op *OpTypePipe) Opcode() uint32 { return opcodeTypePipe }
func (op *OpTypePipe) Optional() bool { return false }
func (op *OpTypePipe) Verify() error  { return nil }

type OpTypeForwardPointer struct {
	PointerType  Id
	StorageClass StorageClass
}

func (op *OpTypeForwardPointer) Opcode() uint32 { return opcodeTypeForwardPointer }
func (op *OpTypeForwardPointer) Optional() bool { return false }
func (op *OpTypeForwardPointer) Verify() error  { return nil }

type OpConstantTrue struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpConstantTrue) Opcode() uint32 { return opcodeConstantTrue }
func (op *OpConstantTrue) Optional() bool { return false }
func (op *OpConstantTrue) Verify() error  { return nil }

type OpConstantFalse struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpConstantFalse) Opcode() uint32 { return opcodeConstantFalse }
func (op *OpConstantFalse) Optional() bool { return false }
func (op *OpConstantFalse) Verify() error  { return nil }

type OpConstant struct {
	IdResultType Id
	IdResult     Id
	Value        uint32
}

func (op *OpConstant) Opcode() uint32 { return opcodeConstant }
func (op *OpConstant) Optional() bool { return false }
func (op *OpConstant) Verify() error  { return nil }

type OpConstantComposite struct {
	IdResultType Id
	IdResult     Id
	Constituents []Id `spirv:"optional"`
}

func (op *OpConstantComposite) Opcode() uint32 { return opcodeConstantComposite }
func (op *OpConstantComposite) Optional() bool { return false }
func (op *OpConstantComposite) Verify() error  { return nil }

type OpConstantSampler struct {
	IdResultType          Id
	IdResult              Id
	SamplerAddressingMode SamplerAddressingMode
	Param                 uint32
	SamplerFilterMode     SamplerFilterMode
}

func (op *OpConstantSampler) Opcode() uint32 { return opcodeConstantSampler }
func (op *OpConstantSampler) Optional() bool { return false }
func (op *OpConstantSampler) Verify() error  { return nil }

type OpConstantNull struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpConstantNull) Opcode() uint32 { return opcodeConstantNull }
func (op *OpConstantNull) Optional() bool { return false }
func (op *OpConstantNull) Verify() error  { return nil }

type OpSpecConstantTrue struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpSpecConstantTrue) Opcode() uint32 { return opcodeSpecConstantTrue }
func (op *OpSpecConstantTrue) Optional() bool { return false }
func (op *OpSpecConstantTrue) Verify() error  { return nil }

type OpSpecConstantFalse struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpSpecConstantFalse) Opcode() uint32 { return opcodeSpecConstantFalse }
func (op *OpSpecConstantFalse) Optional() bool { return false }
func (op *OpSpecConstantFalse) Verify() error  { return nil }

type OpSpecConstant struct {
	IdResultType Id
	IdResult     Id
	Value        uint32
}

func (op *OpSpecConstant) Opcode() uint32 { return opcodeSpecConstant }
func (op *OpSpecConstant) Optional() bool { return false }
func (op *OpSpecConstant) Verify() error  { return nil }

type OpSpecConstantComposite struct {
	IdResultType Id
	IdResult     Id
	Constituents []Id `spirv:"optional"`
}

func (op *OpSpecConstantComposite) Opcode() uint32 { return opcodeSpecConstantComposite }
func (op *OpSpecConstantComposite) Optional() bool { return false }
func (op *OpSpecConstantComposite) Verify() error  { return nil }

type OpSpecConstantOp struct {
	IdResultType Id
	IdResult     Id
	OpcodeParam  uint32
}

func (op *OpSpecConstantOp) Opcode() uint32 { return opcodeSpecConstantOp }
func (op *OpSpecConstantOp) Optional() bool { return false }
func (op *OpSpecConstantOp) Verify() error  { return nil }

type OpFunction struct {
	IdResultType    Id
	IdResult        Id
	FunctionControl FunctionControl
	FunctionType    Id
}

func (op *OpFunction) Opcode() uint32 { return opcodeFunction }
func (op *OpFunction) Optional() bool { return false }
func (op *OpFunction) Verify() error  { return nil }

type OpFunctionParameter struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpFunctionParameter) Opcode() uint32 { return opcodeFunctionParameter }
func (op *OpFunctionParameter) Optional() bool { return false }
func (op *OpFunctionParameter) Verify() error  { return nil }

type OpFunctionEnd struct {
}

func (op *OpFunctionEnd) Opcode() uint32 { return opcodeFunctionEnd }
func (op *OpFunctionEnd) Optional() bool { return false }
func (op *OpFunctionEnd) Verify() error  { return nil }

type OpFunctionCall struct {
	IdResultType Id
	IdResult     Id
	Function     Id
	Argv         []Id `spirv:"optional"`
}

func (op *OpFunctionCall) Opcode() uint32 { return opcodeFunctionCall }
func (op *OpFunctionCall) Optional() bool { return false }
func (op *OpFunctionCall) Verify() error  { return nil }

type OpVariable struct {
	IdResultType Id
	IdResult     Id
	StorageClass StorageClass
	Initializer  Id `spirv:"optional"`
}

func (op *OpVariable) Opcode() uint32 { return opcodeVariable }
func (op *OpVariable) Optional() bool { return false }
func (op *OpVariable) Verify() error  { return nil }

type OpImageTexelPointer struct {
	IdResultType Id
	IdResult     Id
	Image        Id
	Coordinate   Id
	Sample       Id
}

func (op *OpImageTexelPointer) Opcode() uint32 { return opcodeImageTexelPointer }
func (op *OpImageTexelPointer) Optional() bool { return false }
func (op *OpImageTexelPointer) Verify() error  { return nil }

type OpLoad struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	MemoryAccess MemoryAccess `spirv:"optional"`
}

func (op *OpLoad) Opcode() uint32 { return opcodeLoad }
func (op *OpLoad) Optional() bool { return false }
func (op *OpLoad) Verify() error  { return nil }

type OpStore struct {
	Pointer      Id
	Object       Id
	MemoryAccess MemoryAccess `spirv:"optional"`
}

func (op *OpStore) Opcode() uint32 { return opcodeStore }
func (op *OpStore) Optional() bool { return false }
func (op *OpStore) Verify() error  { return nil }

type OpCopyMemory struct {
	Target        Id
	Source        Id
	MemoryAccess  MemoryAccess `spirv:"optional"`
	MemoryAccess2 MemoryAccess `spirv:"optional"`
}

func (op *OpCopyMemory) Opcode() uint32 { return opcodeCopyMemory }
func (op *OpCopyMemory) Optional() bool { return false }
func (op *OpCopyMemory) Verify() error  { return nil }

type OpCopyMemorySized struct {
	Target        Id
	Source        Id
	Size          Id
	MemoryAccess  MemoryAccess `spirv:"optional"`
	MemoryAccess2 MemoryAccess `spirv:"optional"`
}

func (op *OpCopyMemorySized) Opcode() uint32 { return opcodeCopyMemorySized }
func (op *OpCopyMemorySized) Optional() bool { return false }
func (op *OpCopyMemorySized) Verify() error  { return nil }

type OpAccessChain struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Indexes      []Id `spirv:"optional"`
}

func (op *OpAccessChain) Opcode() uint32 { return opcodeAccessChain }
func (op *OpAccessChain) Optional() bool { return false }
func (op *OpAccessChain) Verify() error  { return nil }

type OpInBoundsAccessChain struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Indexes      []Id `spirv:"optional"`
}

func (op *OpInBoundsAccessChain) Opcode() uint32 { return opcodeInBoundsAccessChain }
func (op *OpInBoundsAccessChain) Optional() bool { return false }
func (op *OpInBoundsAccessChain) Verify() error  { return nil }

type OpPtrAccessChain struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Element      Id
	Indexes      []Id `spirv:"optional"`
}

func (op *OpPtrAccessChain) Opcode() uint32 { return opcodePtrAccessChain }
func (op *OpPtrAccessChain) Optional() bool { return false }
func (op *OpPtrAccessChain) Verify() error  { return nil }

type OpArrayLength struct {
	IdResultType Id
	IdResult     Id
	Structure    Id
	ArrayMember  uint32
}

func (op *OpArrayLength) Opcode() uint32 { return opcodeArrayLength }
func (op *OpArrayLength) Optional() bool { return false }
func (op *OpArrayLength) Verify() error  { return nil }

type OpGenericPtrMemSemantics struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
}

func (op *OpGenericPtrMemSemantics) Opcode() uint32 { return opcodeGenericPtrMemSemantics }
func (op *OpGenericPtrMemSemantics) Optional() bool { return false }
func (op *OpGenericPtrMemSemantics) Verify() error  { return nil }

type OpInBoundsPtrAccessChain struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Element      Id
	Indexes      []Id `spirv:"optional"`
}

func (op *OpInBoundsPtrAccessChain) Opcode() uint32 { return opcodeInBoundsPtrAccessChain }
func (op *OpInBoundsPtrAccessChain) Optional() bool { return false }
func (op *OpInBoundsPtrAccessChain) Verify() error  { return nil }

type OpDecorate struct {
	Target     Id
	Decoration Decoration
	Argv       []uint32
}

func (op *OpDecorate) Opcode() uint32 { return opcodeDecorate }
func (op *OpDecorate) Optional() bool { return false }
func (op *OpDecorate) Verify() error  { return nil }

type OpMemberDecorate struct {
	StructureType Id
	Member        uint32
	Decoration    Decoration
	Argv          []uint32
}

func (op *OpMemberDecorate) Opcode() uint32 { return opcodeMemberDecorate }
func (op *OpMemberDecorate) Optional() bool { return false }
func (op *OpMemberDecorate) Verify() error  { return nil }

type OpDecorationGroup struct {
	IdResult Id
}

func (op *OpDecorationGroup) Opcode() uint32 { return opcodeDecorationGroup }
func (op *OpDecorationGroup) Optional() bool { return false }
func (op *OpDecorationGroup) Verify() error  { return nil }

type OpGroupDecorate struct {
	DecorationGroup Id
	Targets         []Id `spirv:"optional"`
}

func (op *OpGroupDecorate) Opcode() uint32 { return opcodeGroupDecorate }
func (op *OpGroupDecorate) Optional() bool { return false }
func (op *OpGroupDecorate) Verify() error  { return nil }

type OpGroupMemberDecorate struct {
	DecorationGroup Id
	Targets         []PairIdRefLiteralInteger `spirv:"optional"`
}

func (op *OpGroupMemberDecorate) Opcode() uint32 { return opcodeGroupMemberDecorate }
func (op *OpGroupMemberDecorate) Optional() bool { return false }
func (op *OpGroupMemberDecorate) Verify() error  { return nil }

type OpVectorExtractDynamic struct {
	IdResultType Id
	IdResult     Id
	Vector       Id
	Index        Id
}

func (op *OpVectorExtractDynamic) Opcode() uint32 { return opcodeVectorExtractDynamic }
func (op *OpVectorExtractDynamic) Optional() bool { return false }
func (op *OpVectorExtractDynamic) Verify() error  { return nil }

type OpVectorInsertDynamic struct {
	IdResultType Id
	IdResult     Id
	Vector       Id
	Component    Id
	Index        Id
}

func (op *OpVectorInsertDynamic) Opcode() uint32 { return opcodeVectorInsertDynamic }
func (op *OpVectorInsertDynamic) Optional() bool { return false }
func (op *OpVectorInsertDynamic) Verify() error  { return nil }

type OpVectorShuffle struct {
	IdResultType Id
	IdResult     Id
	Vector1      Id
	Vector2      Id
	Components   []uint32 `spirv:"optional"`
}

func (op *OpVectorShuffle) Opcode() uint32 { return opcodeVectorShuffle }
func (op *OpVectorShuffle) Optional() bool { return false }
func (op *OpVectorShuffle) Verify() error  { return nil }

type OpCompositeConstruct struct {
	IdResultType Id
	IdResult     Id
	Constituents []Id `spirv:"optional"`
}

func (op *OpCompositeConstruct) Opcode() uint32 { return opcodeCompositeConstruct }
func (op *OpCompositeConstruct) Optional() bool { return false }
func (op *OpCompositeConstruct) Verify() error  { return nil }

type OpCompositeExtract struct {
	IdResultType Id
	IdResult     Id
	Composite    Id
	Indexes      []uint32 `spirv:"optional"`
}

func (op *OpCompositeExtract) Opcode() uint32 { return opcodeCompositeExtract }
func (op *OpCompositeExtract) Optional() bool { return false }
func (op *OpCompositeExtract) Verify() error  { return nil }

type OpCompositeInsert struct {
	IdResultType Id
	IdResult     Id
	Object       Id
	Composite    Id
	Indexes      []uint32 `spirv:"optional"`
}

func (op *OpCompositeInsert) Opcode() uint32 { return opcodeCompositeInsert }
func (op *OpCompositeInsert) Optional() bool { return false }
func (op *OpCompositeInsert) Verify() error  { return nil }

type OpCopyObject struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpCopyObject) Opcode() uint32 { return opcodeCopyObject }
func (op *OpCopyObject) Optional() bool { return false }
func (op *OpCopyObject) Verify() error  { return nil }

type OpTranspose struct {
	IdResultType Id
	IdResult     Id
	Matrix       Id
}

func (op *OpTranspose) Opcode() uint32 { return opcodeTranspose }
func (op *OpTranspose) Optional() bool { return false }
func (op *OpTranspose) Verify() error  { return nil }

type OpSampledImage struct {
	IdResultType Id
	IdResult     Id
	Image        Id
	Sampler      Id
}

func (op *OpSampledImage) Opcode() uint32 { return opcodeSampledImage }
func (op *OpSampledImage) Optional() bool { return false }
func (op *OpSampledImage) Verify() error  { return nil }

type OpImageSampleImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSampleImplicitLod) Opcode() uint32 { return opcodeImageSampleImplicitLod }
func (op *OpImageSampleImplicitLod) Optional() bool { return false }
func (op *OpImageSampleImplicitLod) Verify() error  { return nil }

type OpImageSampleExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands
}

func (op *OpImageSampleExplicitLod) Opcode() uint32 { return opcodeImageSampleExplicitLod }
func (op *OpImageSampleExplicitLod) Optional() bool { return false }
func (op *OpImageSampleExplicitLod) Verify() error  { return nil }

type OpImageSampleDrefImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSampleDrefImplicitLod) Opcode() uint32 { return opcodeImageSampleDrefImplicitLod }
func (op *OpImageSampleDrefImplicitLod) Optional() bool { return false }
func (op *OpImageSampleDrefImplicitLod) Verify() error  { return nil }

type OpImageSampleDrefExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands
}

func (op *OpImageSampleDrefExplicitLod) Opcode() uint32 { return opcodeImageSampleDrefExplicitLod }
func (op *OpImageSampleDrefExplicitLod) Optional() bool { return false }
func (op *OpImageSampleDrefExplicitLod) Verify() error  { return nil }

type OpImageSampleProjImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSampleProjImplicitLod) Opcode() uint32 { return opcodeImageSampleProjImplicitLod }
func (op *OpImageSampleProjImplicitLod) Optional() bool { return false }
func (op *OpImageSampleProjImplicitLod) Verify() error  { return nil }

type OpImageSampleProjExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands
}

func (op *OpImageSampleProjExplicitLod) Opcode() uint32 { return opcodeImageSampleProjExplicitLod }
func (op *OpImageSampleProjExplicitLod) Optional() bool { return false }
func (op *OpImageSampleProjExplicitLod) Verify() error  { return nil }

type OpImageSampleProjDrefImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSampleProjDrefImplicitLod) Opcode() uint32 {
	return opcodeImageSampleProjDrefImplicitLod
}
func (op *OpImageSampleProjDrefImplicitLod) Optional() bool { return false }
func (op *OpImageSampleProjDrefImplicitLod) Verify() error  { return nil }

type OpImageSampleProjDrefExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands
}

func (op *OpImageSampleProjDrefExplicitLod) Opcode() uint32 {
	return opcodeImageSampleProjDrefExplicitLod
}
func (op *OpImageSampleProjDrefExplicitLod) Optional() bool { return false }
func (op *OpImageSampleProjDrefExplicitLod) Verify() error  { return nil }

type OpImageFetch struct {
	IdResultType  Id
	IdResult      Id
	Image         Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageFetch) Opcode() uint32 { return opcodeImageFetch }
func (op *OpImageFetch) Optional() bool { return false }
func (op *OpImageFetch) Verify() error  { return nil }

type OpImageGather struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	Component     Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageGather) Opcode() uint32 { return opcodeImageGather }
func (op *OpImageGather) Optional() bool { return false }
func (op *OpImageGather) Verify() error  { return nil }

type OpImageDrefGather struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageDrefGather) Opcode() uint32 { return opcodeImageDrefGather }
func (op *OpImageDrefGather) Optional() bool { return false }
func (op *OpImageDrefGather) Verify() error  { return nil }

type OpImageRead struct {
	IdResultType  Id
	IdResult      Id
	Image         Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageRead) Opcode() uint32 { return opcodeImageRead }
func (op *OpImageRead) Optional() bool { return false }
func (op *OpImageRead) Verify() error  { return nil }

type OpImageWrite struct {
	Image         Id
	Coordinate    Id
	Texel         Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageWrite) Opcode() uint32 { return opcodeImageWrite }
func (op *OpImageWrite) Optional() bool { return false }
func (op *OpImageWrite) Verify() error  { return nil }

type OpImage struct {
	IdResultType Id
	IdResult     Id
	SampledImage Id
}

func (op *OpImage) Opcode() uint32 { return opcodeImage }
func (op *OpImage) Optional() bool { return false }
func (op *OpImage) Verify() error  { return nil }

type OpImageQueryFormat struct {
	IdResultType Id
	IdResult     Id
	Image        Id
}

func (op *OpImageQueryFormat) Opcode() uint32 { return opcodeImageQueryFormat }
func (op *OpImageQueryFormat) Optional() bool { return false }
func (op *OpImageQueryFormat) Verify() error  { return nil }

type OpImageQueryOrder struct {
	IdResultType Id
	IdResult     Id
	Image        Id
}

func (op *OpImageQueryOrder) Opcode() uint32 { return opcodeImageQueryOrder }
func (op *OpImageQueryOrder) Optional() bool { return false }
func (op *OpImageQueryOrder) Verify() error  { return nil }

type OpImageQuerySizeLod struct {
	IdResultType  Id
	IdResult      Id
	Image         Id
	LevelOfDetail Id
}

func (op *OpImageQuerySizeLod) Opcode() uint32 { return opcodeImageQuerySizeLod }
func (op *OpImageQuerySizeLod) Optional() bool { return false }
func (op *OpImageQuerySizeLod) Verify() error  { return nil }

type OpImageQuerySize struct {
	IdResultType Id
	IdResult     Id
	Image        Id
}

func (op *OpImageQuerySize) Opcode() uint32 { return opcodeImageQuerySize }
func (op *OpImageQuerySize) Optional() bool { return false }
func (op *OpImageQuerySize) Verify() error  { return nil }

type OpImageQueryLod struct {
	IdResultType Id
	IdResult     Id
	SampledImage Id
	Coordinate   Id
}

func (op *OpImageQueryLod) Opcode() uint32 { return opcodeImageQueryLod }
func (op *OpImageQueryLod) Optional() bool { return false }
func (op *OpImageQueryLod) Verify() error  { return nil }

type OpImageQueryLevels struct {
	IdResultType Id
	IdResult     Id
	Image        Id
}

func (op *OpImageQueryLevels) Opcode() uint32 { return opcodeImageQueryLevels }
func (op *OpImageQueryLevels) Optional() bool { return false }
func (op *OpImageQueryLevels) Verify() error  { return nil }

type OpImageQuerySamples struct {
	IdResultType Id
	IdResult     Id
	Image        Id
}

func (op *OpImageQuerySamples) Opcode() uint32 { return opcodeImageQuerySamples }
func (op *OpImageQuerySamples) Optional() bool { return false }
func (op *OpImageQuerySamples) Verify() error  { return nil }

type OpConvertFToU struct {
	IdResultType Id
	IdResult     Id
	FloatValue   Id
}

func (op *OpConvertFToU) Opcode() uint32 { return opcodeConvertFToU }
func (op *OpConvertFToU) Optional() bool { return false }
func (op *OpConvertFToU) Verify() error  { return nil }

type OpConvertFToS struct {
	IdResultType Id
	IdResult     Id
	FloatValue   Id
}

func (op *OpConvertFToS) Opcode() uint32 { return opcodeConvertFToS }
func (op *OpConvertFToS) Optional() bool { return false }
func (op *OpConvertFToS) Verify() error  { return nil }

type OpConvertSToF struct {
	IdResultType Id
	IdResult     Id
	SignedValue  Id
}

func (op *OpConvertSToF) Opcode() uint32 { return opcodeConvertSToF }
func (op *OpConvertSToF) Optional() bool { return false }
func (op *OpConvertSToF) Verify() error  { return nil }

type OpConvertUToF struct {
	IdResultType  Id
	IdResult      Id
	UnsignedValue Id
}

func (op *OpConvertUToF) Opcode() uint32 { return opcodeConvertUToF }
func (op *OpConvertUToF) Optional() bool { return false }
func (op *OpConvertUToF) Verify() error  { return nil }

type OpUConvert struct {
	IdResultType  Id
	IdResult      Id
	UnsignedValue Id
}

func (op *OpUConvert) Opcode() uint32 { return opcodeUConvert }
func (op *OpUConvert) Optional() bool { return false }
func (op *OpUConvert) Verify() error  { return nil }

type OpSConvert struct {
	IdResultType Id
	IdResult     Id
	SignedValue  Id
}

func (op *OpSConvert) Opcode() uint32 { return opcodeSConvert }
func (op *OpSConvert) Optional() bool { return false }
func (op *OpSConvert) Verify() error  { return nil }

type OpFConvert struct {
	IdResultType Id
	IdResult     Id
	FloatValue   Id
}

func (op *OpFConvert) Opcode() uint32 { return opcodeFConvert }
func (op *OpFConvert) Optional() bool { return false }
func (op *OpFConvert) Verify() error  { return nil }

type OpQuantizeToF16 struct {
	IdResultType Id
	IdResult     Id
	Value        Id
}

func (op *OpQuantizeToF16) Opcode() uint32 { return opcodeQuantizeToF16 }
func (op *OpQuantizeToF16) Optional() bool { return false }
func (op *OpQuantizeToF16) Verify() error  { return nil }

type OpConvertPtrToU struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
}

func (op *OpConvertPtrToU) Opcode() uint32 { return opcodeConvertPtrToU }
func (op *OpConvertPtrToU) Optional() bool { return false }
func (op *OpConvertPtrToU) Verify() error  { return nil }

type OpSatConvertSToU struct {
	IdResultType Id
	IdResult     Id
	SignedValue  Id
}

func (op *OpSatConvertSToU) Opcode() uint32 { return opcodeSatConvertSToU }
func (op *OpSatConvertSToU) Optional() bool { return false }
func (op *OpSatConvertSToU) Verify() error  { return nil }

type OpSatConvertUToS struct {
	IdResultType  Id
	IdResult      Id
	UnsignedValue Id
}

func (op *OpSatConvertUToS) Opcode() uint32 { return opcodeSatConvertUToS }
func (op *OpSatConvertUToS) Optional() bool { return false }
func (op *OpSatConvertUToS) Verify() error  { return nil }

type OpConvertUToPtr struct {
	IdResultType Id
	IdResult     Id
	IntegerValue Id
}

func (op *OpConvertUToPtr) Opcode() uint32 { return opcodeConvertUToPtr }
func (op *OpConvertUToPtr) Optional() bool { return false }
func (op *OpConvertUToPtr) Verify() error  { return nil }

type OpPtrCastToGeneric struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
}

func (op *OpPtrCastToGeneric) Opcode() uint32 { return opcodePtrCastToGeneric }
func (op *OpPtrCastToGeneric) Optional() bool { return false }
func (op *OpPtrCastToGeneric) Verify() error  { return nil }

type OpGenericCastToPtr struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
}

func (op *OpGenericCastToPtr) Opcode() uint32 { return opcodeGenericCastToPtr }
func (op *OpGenericCastToPtr) Optional() bool { return false }
func (op *OpGenericCastToPtr) Verify() error  { return nil }

type OpGenericCastToPtrExplicit struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Storage      StorageClass
}

func (op *OpGenericCastToPtrExplicit) Opcode() uint32 { return opcodeGenericCastToPtrExplicit }
func (op *OpGenericCastToPtrExplicit) Optional() bool { return false }
func (op *OpGenericCastToPtrExplicit) Verify() error  { return nil }

type OpBitcast struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpBitcast) Opcode() uint32 { return opcodeBitcast }
func (op *OpBitcast) Optional() bool { return false }
func (op *OpBitcast) Verify() error  { return nil }

type OpSNegate struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpSNegate) Opcode() uint32 { return opcodeSNegate }
func (op *OpSNegate) Optional() bool { return false }
func (op *OpSNegate) Verify() error  { return nil }

type OpFNegate struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpFNegate) Opcode() uint32 { return opcodeFNegate }
func (op *OpFNegate) Optional() bool { return false }
func (op *OpFNegate) Verify() error  { return nil }

type OpIAdd struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpIAdd) Opcode() uint32 { return opcodeIAdd }
func (op *OpIAdd) Optional() bool { return false }
func (op *OpIAdd) Verify() error  { return nil }

type OpFAdd struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFAdd) Opcode() uint32 { return opcodeFAdd }
func (op *OpFAdd) Optional() bool { return false }
func (op *OpFAdd) Verify() error  { return nil }

type OpISub struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpISub) Opcode() uint32 { return opcodeISub }
func (op *OpISub) Optional() bool { return false }
func (op *OpISub) Verify() error  { return nil }

type OpFSub struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFSub) Opcode() uint32 { return opcodeFSub }
func (op *OpFSub) Optional() bool { return false }
func (op *OpFSub) Verify() error  { return nil }

type OpIMul struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpIMul) Opcode() uint32 { return opcodeIMul }
func (op *OpIMul) Optional() bool { return false }
func (op *OpIMul) Verify() error  { return nil }

type OpFMul struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFMul) Opcode() uint32 { return opcodeFMul }
func (op *OpFMul) Optional() bool { return false }
func (op *OpFMul) Verify() error  { return nil }

type OpUDiv struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpUDiv) Opcode() uint32 { return opcodeUDiv }
func (op *OpUDiv) Optional() bool { return false }
func (op *OpUDiv) Verify() error  { return nil }

type OpSDiv struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSDiv) Opcode() uint32 { return opcodeSDiv }
func (op *OpSDiv) Optional() bool { return false }
func (op *OpSDiv) Verify() error  { return nil }

type OpFDiv struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFDiv) Opcode() uint32 { return opcodeFDiv }
func (op *OpFDiv) Optional() bool { return false }
func (op *OpFDiv) Verify() error  { return nil }

type OpUMod struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpUMod) Opcode() uint32 { return opcodeUMod }
func (op *OpUMod) Optional() bool { return false }
func (op *OpUMod) Verify() error  { return nil }

type OpSRem struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSRem) Opcode() uint32 { return opcodeSRem }
func (op *OpSRem) Optional() bool { return false }
func (op *OpSRem) Verify() error  { return nil }

type OpSMod struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSMod) Opcode() uint32 { return opcodeSMod }
func (op *OpSMod) Optional() bool { return false }
func (op *OpSMod) Verify() error  { return nil }

type OpFRem struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFRem) Opcode() uint32 { return opcodeFRem }
func (op *OpFRem) Optional() bool { return false }
func (op *OpFRem) Verify() error  { return nil }

type OpFMod struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFMod) Opcode() uint32 { return opcodeFMod }
func (op *OpFMod) Optional() bool { return false }
func (op *OpFMod) Verify() error  { return nil }

type OpVectorTimesScalar struct {
	IdResultType Id
	IdResult     Id
	Vector       Id
	Scalar       Id
}

func (op *OpVectorTimesScalar) Opcode() uint32 { return opcodeVectorTimesScalar }
func (op *OpVectorTimesScalar) Optional() bool { return false }
func (op *OpVectorTimesScalar) Verify() error  { return nil }

type OpMatrixTimesScalar struct {
	IdResultType Id
	IdResult     Id
	Matrix       Id
	Scalar       Id
}

func (op *OpMatrixTimesScalar) Opcode() uint32 { return opcodeMatrixTimesScalar }
func (op *OpMatrixTimesScalar) Optional() bool { return false }
func (op *OpMatrixTimesScalar) Verify() error  { return nil }

type OpVectorTimesMatrix struct {
	IdResultType Id
	IdResult     Id
	Vector       Id
	Matrix       Id
}

func (op *OpVectorTimesMatrix) Opcode() uint32 { return opcodeVectorTimesMatrix }
func (op *OpVectorTimesMatrix) Optional() bool { return false }
func (op *OpVectorTimesMatrix) Verify() error  { return nil }

type OpMatrixTimesVector struct {
	IdResultType Id
	IdResult     Id
	Matrix       Id
	Vector       Id
}

func (op *OpMatrixTimesVector) Opcode() uint32 { return opcodeMatrixTimesVector }
func (op *OpMatrixTimesVector) Optional() bool { return false }
func (op *OpMatrixTimesVector) Verify() error  { return nil }

type OpMatrixTimesMatrix struct {
	IdResultType Id
	IdResult     Id
	LeftMatrix   Id
	RightMatrix  Id
}

func (op *OpMatrixTimesMatrix) Opcode() uint32 { return opcodeMatrixTimesMatrix }
func (op *OpMatrixTimesMatrix) Optional() bool { return false }
func (op *OpMatrixTimesMatrix) Verify() error  { return nil }

type OpOuterProduct struct {
	IdResultType Id
	IdResult     Id
	Vector1      Id
	Vector2      Id
}

func (op *OpOuterProduct) Opcode() uint32 { return opcodeOuterProduct }
func (op *OpOuterProduct) Optional() bool { return false }
func (op *OpOuterProduct) Verify() error  { return nil }

type OpDot struct {
	IdResultType Id
	IdResult     Id
	Vector1      Id
	Vector2      Id
}

func (op *OpDot) Opcode() uint32 { return opcodeDot }
func (op *OpDot) Optional() bool { return false }
func (op *OpDot) Verify() error  { return nil }

type OpIAddCarry struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpIAddCarry) Opcode() uint32 { return opcodeIAddCarry }
func (op *OpIAddCarry) Optional() bool { return false }
func (op *OpIAddCarry) Verify() error  { return nil }

type OpISubBorrow struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpISubBorrow) Opcode() uint32 { return opcodeISubBorrow }
func (op *OpISubBorrow) Optional() bool { return false }
func (op *OpISubBorrow) Verify() error  { return nil }

type OpUMulExtended struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpUMulExtended) Opcode() uint32 { return opcodeUMulExtended }
func (op *OpUMulExtended) Optional() bool { return false }
func (op *OpUMulExtended) Verify() error  { return nil }

type OpSMulExtended struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSMulExtended) Opcode() uint32 { return opcodeSMulExtended }
func (op *OpSMulExtended) Optional() bool { return false }
func (op *OpSMulExtended) Verify() error  { return nil }

type OpAny struct {
	IdResultType Id
	IdResult     Id
	Vector       Id
}

func (op *OpAny) Opcode() uint32 { return opcodeAny }
func (op *OpAny) Optional() bool { return false }
func (op *OpAny) Verify() error  { return nil }

type OpAll struct {
	IdResultType Id
	IdResult     Id
	Vector       Id
}

func (op *OpAll) Opcode() uint32 { return opcodeAll }
func (op *OpAll) Optional() bool { return false }
func (op *OpAll) Verify() error  { return nil }

type OpIsNan struct {
	IdResultType Id
	IdResult     Id
	x            Id
}

func (op *OpIsNan) Opcode() uint32 { return opcodeIsNan }
func (op *OpIsNan) Optional() bool { return false }
func (op *OpIsNan) Verify() error  { return nil }

type OpIsInf struct {
	IdResultType Id
	IdResult     Id
	x            Id
}

func (op *OpIsInf) Opcode() uint32 { return opcodeIsInf }
func (op *OpIsInf) Optional() bool { return false }
func (op *OpIsInf) Verify() error  { return nil }

type OpIsFinite struct {
	IdResultType Id
	IdResult     Id
	x            Id
}

func (op *OpIsFinite) Opcode() uint32 { return opcodeIsFinite }
func (op *OpIsFinite) Optional() bool { return false }
func (op *OpIsFinite) Verify() error  { return nil }

type OpIsNormal struct {
	IdResultType Id
	IdResult     Id
	x            Id
}

func (op *OpIsNormal) Opcode() uint32 { return opcodeIsNormal }
func (op *OpIsNormal) Optional() bool { return false }
func (op *OpIsNormal) Verify() error  { return nil }

type OpSignBitSet struct {
	IdResultType Id
	IdResult     Id
	x            Id
}

func (op *OpSignBitSet) Opcode() uint32 { return opcodeSignBitSet }
func (op *OpSignBitSet) Optional() bool { return false }
func (op *OpSignBitSet) Verify() error  { return nil }

type OpLessOrGreater struct {
	IdResultType Id
	IdResult     Id
	x            Id
	y            Id
}

func (op *OpLessOrGreater) Opcode() uint32 { return opcodeLessOrGreater }
func (op *OpLessOrGreater) Optional() bool { return false }
func (op *OpLessOrGreater) Verify() error  { return nil }

type OpOrdered struct {
	IdResultType Id
	IdResult     Id
	x            Id
	y            Id
}

func (op *OpOrdered) Opcode() uint32 { return opcodeOrdered }
func (op *OpOrdered) Optional() bool { return false }
func (op *OpOrdered) Verify() error  { return nil }

type OpUnordered struct {
	IdResultType Id
	IdResult     Id
	x            Id
	y            Id
}

func (op *OpUnordered) Opcode() uint32 { return opcodeUnordered }
func (op *OpUnordered) Optional() bool { return false }
func (op *OpUnordered) Verify() error  { return nil }

type OpLogicalEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpLogicalEqual) Opcode() uint32 { return opcodeLogicalEqual }
func (op *OpLogicalEqual) Optional() bool { return false }
func (op *OpLogicalEqual) Verify() error  { return nil }

type OpLogicalNotEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpLogicalNotEqual) Opcode() uint32 { return opcodeLogicalNotEqual }
func (op *OpLogicalNotEqual) Optional() bool { return false }
func (op *OpLogicalNotEqual) Verify() error  { return nil }

type OpLogicalOr struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpLogicalOr) Opcode() uint32 { return opcodeLogicalOr }
func (op *OpLogicalOr) Optional() bool { return false }
func (op *OpLogicalOr) Verify() error  { return nil }

type OpLogicalAnd struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpLogicalAnd) Opcode() uint32 { return opcodeLogicalAnd }
func (op *OpLogicalAnd) Optional() bool { return false }
func (op *OpLogicalAnd) Verify() error  { return nil }

type OpLogicalNot struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpLogicalNot) Opcode() uint32 { return opcodeLogicalNot }
func (op *OpLogicalNot) Optional() bool { return false }
func (op *OpLogicalNot) Verify() error  { return nil }

type OpSelect struct {
	IdResultType Id
	IdResult     Id
	Condition    Id
	Object1      Id
	Object2      Id
}

func (op *OpSelect) Opcode() uint32 { return opcodeSelect }
func (op *OpSelect) Optional() bool { return false }
func (op *OpSelect) Verify() error  { return nil }

type OpIEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpIEqual) Opcode() uint32 { return opcodeIEqual }
func (op *OpIEqual) Optional() bool { return false }
func (op *OpIEqual) Verify() error  { return nil }

type OpINotEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpINotEqual) Opcode() uint32 { return opcodeINotEqual }
func (op *OpINotEqual) Optional() bool { return false }
func (op *OpINotEqual) Verify() error  { return nil }

type OpUGreaterThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpUGreaterThan) Opcode() uint32 { return opcodeUGreaterThan }
func (op *OpUGreaterThan) Optional() bool { return false }
func (op *OpUGreaterThan) Verify() error  { return nil }

type OpSGreaterThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSGreaterThan) Opcode() uint32 { return opcodeSGreaterThan }
func (op *OpSGreaterThan) Optional() bool { return false }
func (op *OpSGreaterThan) Verify() error  { return nil }

type OpUGreaterThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpUGreaterThanEqual) Opcode() uint32 { return opcodeUGreaterThanEqual }
func (op *OpUGreaterThanEqual) Optional() bool { return false }
func (op *OpUGreaterThanEqual) Verify() error  { return nil }

type OpSGreaterThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSGreaterThanEqual) Opcode() uint32 { return opcodeSGreaterThanEqual }
func (op *OpSGreaterThanEqual) Optional() bool { return false }
func (op *OpSGreaterThanEqual) Verify() error  { return nil }

type OpULessThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpULessThan) Opcode() uint32 { return opcodeULessThan }
func (op *OpULessThan) Optional() bool { return false }
func (op *OpULessThan) Verify() error  { return nil }

type OpSLessThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSLessThan) Opcode() uint32 { return opcodeSLessThan }
func (op *OpSLessThan) Optional() bool { return false }
func (op *OpSLessThan) Verify() error  { return nil }

type OpULessThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpULessThanEqual) Opcode() uint32 { return opcodeULessThanEqual }
func (op *OpULessThanEqual) Optional() bool { return false }
func (op *OpULessThanEqual) Verify() error  { return nil }

type OpSLessThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpSLessThanEqual) Opcode() uint32 { return opcodeSLessThanEqual }
func (op *OpSLessThanEqual) Optional() bool { return false }
func (op *OpSLessThanEqual) Verify() error  { return nil }

type OpFOrdEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFOrdEqual) Opcode() uint32 { return opcodeFOrdEqual }
func (op *OpFOrdEqual) Optional() bool { return false }
func (op *OpFOrdEqual) Verify() error  { return nil }

type OpFUnordEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFUnordEqual) Opcode() uint32 { return opcodeFUnordEqual }
func (op *OpFUnordEqual) Optional() bool { return false }
func (op *OpFUnordEqual) Verify() error  { return nil }

type OpFOrdNotEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFOrdNotEqual) Opcode() uint32 { return opcodeFOrdNotEqual }
func (op *OpFOrdNotEqual) Optional() bool { return false }
func (op *OpFOrdNotEqual) Verify() error  { return nil }

type OpFUnordNotEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFUnordNotEqual) Opcode() uint32 { return opcodeFUnordNotEqual }
func (op *OpFUnordNotEqual) Optional() bool { return false }
func (op *OpFUnordNotEqual) Verify() error  { return nil }

type OpFOrdLessThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFOrdLessThan) Opcode() uint32 { return opcodeFOrdLessThan }
func (op *OpFOrdLessThan) Optional() bool { return false }
func (op *OpFOrdLessThan) Verify() error  { return nil }

type OpFUnordLessThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFUnordLessThan) Opcode() uint32 { return opcodeFUnordLessThan }
func (op *OpFUnordLessThan) Optional() bool { return false }
func (op *OpFUnordLessThan) Verify() error  { return nil }

type OpFOrdGreaterThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFOrdGreaterThan) Opcode() uint32 { return opcodeFOrdGreaterThan }
func (op *OpFOrdGreaterThan) Optional() bool { return false }
func (op *OpFOrdGreaterThan) Verify() error  { return nil }

type OpFUnordGreaterThan struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFUnordGreaterThan) Opcode() uint32 { return opcodeFUnordGreaterThan }
func (op *OpFUnordGreaterThan) Optional() bool { return false }
func (op *OpFUnordGreaterThan) Verify() error  { return nil }

type OpFOrdLessThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFOrdLessThanEqual) Opcode() uint32 { return opcodeFOrdLessThanEqual }
func (op *OpFOrdLessThanEqual) Optional() bool { return false }
func (op *OpFOrdLessThanEqual) Verify() error  { return nil }

type OpFUnordLessThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFUnordLessThanEqual) Opcode() uint32 { return opcodeFUnordLessThanEqual }
func (op *OpFUnordLessThanEqual) Optional() bool { return false }
func (op *OpFUnordLessThanEqual) Verify() error  { return nil }

type OpFOrdGreaterThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFOrdGreaterThanEqual) Opcode() uint32 { return opcodeFOrdGreaterThanEqual }
func (op *OpFOrdGreaterThanEqual) Optional() bool { return false }
func (op *OpFOrdGreaterThanEqual) Verify() error  { return nil }

type OpFUnordGreaterThanEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpFUnordGreaterThanEqual) Opcode() uint32 { return opcodeFUnordGreaterThanEqual }
func (op *OpFUnordGreaterThanEqual) Optional() bool { return false }
func (op *OpFUnordGreaterThanEqual) Verify() error  { return nil }

type OpShiftRightLogical struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Shift        Id
}

func (op *OpShiftRightLogical) Opcode() uint32 { return opcodeShiftRightLogical }
func (op *OpShiftRightLogical) Optional() bool { return false }
func (op *OpShiftRightLogical) Verify() error  { return nil }

type OpShiftRightArithmetic struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Shift        Id
}

func (op *OpShiftRightArithmetic) Opcode() uint32 { return opcodeShiftRightArithmetic }
func (op *OpShiftRightArithmetic) Optional() bool { return false }
func (op *OpShiftRightArithmetic) Verify() error  { return nil }

type OpShiftLeftLogical struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Shift        Id
}

func (op *OpShiftLeftLogical) Opcode() uint32 { return opcodeShiftLeftLogical }
func (op *OpShiftLeftLogical) Optional() bool { return false }
func (op *OpShiftLeftLogical) Verify() error  { return nil }

type OpBitwiseOr struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpBitwiseOr) Opcode() uint32 { return opcodeBitwiseOr }
func (op *OpBitwiseOr) Optional() bool { return false }
func (op *OpBitwiseOr) Verify() error  { return nil }

type OpBitwiseXor struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpBitwiseXor) Opcode() uint32 { return opcodeBitwiseXor }
func (op *OpBitwiseXor) Optional() bool { return false }
func (op *OpBitwiseXor) Verify() error  { return nil }

type OpBitwiseAnd struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpBitwiseAnd) Opcode() uint32 { return opcodeBitwiseAnd }
func (op *OpBitwiseAnd) Optional() bool { return false }
func (op *OpBitwiseAnd) Verify() error  { return nil }

type OpNot struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpNot) Opcode() uint32 { return opcodeNot }
func (op *OpNot) Optional() bool { return false }
func (op *OpNot) Verify() error  { return nil }

type OpBitFieldInsert struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Insert       Id
	Offset       Id
	Count        Id
}

func (op *OpBitFieldInsert) Opcode() uint32 { return opcodeBitFieldInsert }
func (op *OpBitFieldInsert) Optional() bool { return false }
func (op *OpBitFieldInsert) Verify() error  { return nil }

type OpBitFieldSExtract struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Offset       Id
	Count        Id
}

func (op *OpBitFieldSExtract) Opcode() uint32 { return opcodeBitFieldSExtract }
func (op *OpBitFieldSExtract) Optional() bool { return false }
func (op *OpBitFieldSExtract) Verify() error  { return nil }

type OpBitFieldUExtract struct {
	IdResultType Id
	IdResult     Id
	Base         Id
	Offset       Id
	Count        Id
}

func (op *OpBitFieldUExtract) Opcode() uint32 { return opcodeBitFieldUExtract }
func (op *OpBitFieldUExtract) Optional() bool { return false }
func (op *OpBitFieldUExtract) Verify() error  { return nil }

type OpBitReverse struct {
	IdResultType Id
	IdResult     Id
	Base         Id
}

func (op *OpBitReverse) Opcode() uint32 { return opcodeBitReverse }
func (op *OpBitReverse) Optional() bool { return false }
func (op *OpBitReverse) Verify() error  { return nil }

type OpBitCount struct {
	IdResultType Id
	IdResult     Id
	Base         Id
}

func (op *OpBitCount) Opcode() uint32 { return opcodeBitCount }
func (op *OpBitCount) Optional() bool { return false }
func (op *OpBitCount) Verify() error  { return nil }

type OpDPdx struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpDPdx) Opcode() uint32 { return opcodeDPdx }
func (op *OpDPdx) Optional() bool { return false }
func (op *OpDPdx) Verify() error  { return nil }

type OpDPdy struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpDPdy) Opcode() uint32 { return opcodeDPdy }
func (op *OpDPdy) Optional() bool { return false }
func (op *OpDPdy) Verify() error  { return nil }

type OpFwidth struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpFwidth) Opcode() uint32 { return opcodeFwidth }
func (op *OpFwidth) Optional() bool { return false }
func (op *OpFwidth) Verify() error  { return nil }

type OpDPdxFine struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpDPdxFine) Opcode() uint32 { return opcodeDPdxFine }
func (op *OpDPdxFine) Optional() bool { return false }
func (op *OpDPdxFine) Verify() error  { return nil }

type OpDPdyFine struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpDPdyFine) Opcode() uint32 { return opcodeDPdyFine }
func (op *OpDPdyFine) Optional() bool { return false }
func (op *OpDPdyFine) Verify() error  { return nil }

type OpFwidthFine struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpFwidthFine) Opcode() uint32 { return opcodeFwidthFine }
func (op *OpFwidthFine) Optional() bool { return false }
func (op *OpFwidthFine) Verify() error  { return nil }

type OpDPdxCoarse struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpDPdxCoarse) Opcode() uint32 { return opcodeDPdxCoarse }
func (op *OpDPdxCoarse) Optional() bool { return false }
func (op *OpDPdxCoarse) Verify() error  { return nil }

type OpDPdyCoarse struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpDPdyCoarse) Opcode() uint32 { return opcodeDPdyCoarse }
func (op *OpDPdyCoarse) Optional() bool { return false }
func (op *OpDPdyCoarse) Verify() error  { return nil }

type OpFwidthCoarse struct {
	IdResultType Id
	IdResult     Id
	P            Id
}

func (op *OpFwidthCoarse) Opcode() uint32 { return opcodeFwidthCoarse }
func (op *OpFwidthCoarse) Optional() bool { return false }
func (op *OpFwidthCoarse) Verify() error  { return nil }

type OpEmitVertex struct {
}

func (op *OpEmitVertex) Opcode() uint32 { return opcodeEmitVertex }
func (op *OpEmitVertex) Optional() bool { return false }
func (op *OpEmitVertex) Verify() error  { return nil }

type OpEndPrimitive struct {
}

func (op *OpEndPrimitive) Opcode() uint32 { return opcodeEndPrimitive }
func (op *OpEndPrimitive) Optional() bool { return false }
func (op *OpEndPrimitive) Verify() error  { return nil }

type OpEmitStreamVertex struct {
	Stream Id
}

func (op *OpEmitStreamVertex) Opcode() uint32 { return opcodeEmitStreamVertex }
func (op *OpEmitStreamVertex) Optional() bool { return false }
func (op *OpEmitStreamVertex) Verify() error  { return nil }

type OpEndStreamPrimitive struct {
	Stream Id
}

func (op *OpEndStreamPrimitive) Opcode() uint32 { return opcodeEndStreamPrimitive }
func (op *OpEndStreamPrimitive) Optional() bool { return false }
func (op *OpEndStreamPrimitive) Verify() error  { return nil }

type OpControlBarrier struct {
	Execution Id
	Memory    Id
	Semantics Id
}

func (op *OpControlBarrier) Opcode() uint32 { return opcodeControlBarrier }
func (op *OpControlBarrier) Optional() bool { return false }
func (op *OpControlBarrier) Verify() error  { return nil }

type OpMemoryBarrier struct {
	Memory    Id
	Semantics Id
}

func (op *OpMemoryBarrier) Opcode() uint32 { return opcodeMemoryBarrier }
func (op *OpMemoryBarrier) Optional() bool { return false }
func (op *OpMemoryBarrier) Verify() error  { return nil }

type OpAtomicLoad struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
}

func (op *OpAtomicLoad) Opcode() uint32 { return opcodeAtomicLoad }
func (op *OpAtomicLoad) Optional() bool { return false }
func (op *OpAtomicLoad) Verify() error  { return nil }

type OpAtomicStore struct {
	Pointer   Id
	Memory    Id
	Semantics Id
	Value     Id
}

func (op *OpAtomicStore) Opcode() uint32 { return opcodeAtomicStore }
func (op *OpAtomicStore) Optional() bool { return false }
func (op *OpAtomicStore) Verify() error  { return nil }

type OpAtomicExchange struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicExchange) Opcode() uint32 { return opcodeAtomicExchange }
func (op *OpAtomicExchange) Optional() bool { return false }
func (op *OpAtomicExchange) Verify() error  { return nil }

type OpAtomicCompareExchange struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Equal        Id
	Unequal      Id
	Value        Id
	Comparator   Id
}

func (op *OpAtomicCompareExchange) Opcode() uint32 { return opcodeAtomicCompareExchange }
func (op *OpAtomicCompareExchange) Optional() bool { return false }
func (op *OpAtomicCompareExchange) Verify() error  { return nil }

type OpAtomicCompareExchangeWeak struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Equal        Id
	Unequal      Id
	Value        Id
	Comparator   Id
}

func (op *OpAtomicCompareExchangeWeak) Opcode() uint32 { return opcodeAtomicCompareExchangeWeak }
func (op *OpAtomicCompareExchangeWeak) Optional() bool { return false }
func (op *OpAtomicCompareExchangeWeak) Verify() error  { return nil }

type OpAtomicIIncrement struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
}

func (op *OpAtomicIIncrement) Opcode() uint32 { return opcodeAtomicIIncrement }
func (op *OpAtomicIIncrement) Optional() bool { return false }
func (op *OpAtomicIIncrement) Verify() error  { return nil }

type OpAtomicIDecrement struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
}

func (op *OpAtomicIDecrement) Opcode() uint32 { return opcodeAtomicIDecrement }
func (op *OpAtomicIDecrement) Optional() bool { return false }
func (op *OpAtomicIDecrement) Verify() error  { return nil }

type OpAtomicIAdd struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicIAdd) Opcode() uint32 { return opcodeAtomicIAdd }
func (op *OpAtomicIAdd) Optional() bool { return false }
func (op *OpAtomicIAdd) Verify() error  { return nil }

type OpAtomicISub struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicISub) Opcode() uint32 { return opcodeAtomicISub }
func (op *OpAtomicISub) Optional() bool { return false }
func (op *OpAtomicISub) Verify() error  { return nil }

type OpAtomicSMin struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicSMin) Opcode() uint32 { return opcodeAtomicSMin }
func (op *OpAtomicSMin) Optional() bool { return false }
func (op *OpAtomicSMin) Verify() error  { return nil }

type OpAtomicUMin struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicUMin) Opcode() uint32 { return opcodeAtomicUMin }
func (op *OpAtomicUMin) Optional() bool { return false }
func (op *OpAtomicUMin) Verify() error  { return nil }

type OpAtomicSMax struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicSMax) Opcode() uint32 { return opcodeAtomicSMax }
func (op *OpAtomicSMax) Optional() bool { return false }
func (op *OpAtomicSMax) Verify() error  { return nil }

type OpAtomicUMax struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicUMax) Opcode() uint32 { return opcodeAtomicUMax }
func (op *OpAtomicUMax) Optional() bool { return false }
func (op *OpAtomicUMax) Verify() error  { return nil }

type OpAtomicAnd struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicAnd) Opcode() uint32 { return opcodeAtomicAnd }
func (op *OpAtomicAnd) Optional() bool { return false }
func (op *OpAtomicAnd) Verify() error  { return nil }

type OpAtomicOr struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicOr) Opcode() uint32 { return opcodeAtomicOr }
func (op *OpAtomicOr) Optional() bool { return false }
func (op *OpAtomicOr) Verify() error  { return nil }

type OpAtomicXor struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
	Value        Id
}

func (op *OpAtomicXor) Opcode() uint32 { return opcodeAtomicXor }
func (op *OpAtomicXor) Optional() bool { return false }
func (op *OpAtomicXor) Verify() error  { return nil }

type OpPhi struct {
	IdResultType Id
	IdResult     Id
	Argv         []PairIdRefIdRef `spirv:"optional"`
}

func (op *OpPhi) Opcode() uint32 { return opcodePhi }
func (op *OpPhi) Optional() bool { return false }
func (op *OpPhi) Verify() error  { return nil }

type OpLoopMerge struct {
	MergeBlock     Id
	ContinueTarget Id
	LoopControl    LoopControl
	Argv           []uint32
}

func (op *OpLoopMerge) Opcode() uint32 { return opcodeLoopMerge }
func (op *OpLoopMerge) Optional() bool { return false }
func (op *OpLoopMerge) Verify() error  { return nil }

type OpSelectionMerge struct {
	MergeBlock       Id
	SelectionControl SelectionControl
}

func (op *OpSelectionMerge) Opcode() uint32 { return opcodeSelectionMerge }
func (op *OpSelectionMerge) Optional() bool { return false }
func (op *OpSelectionMerge) Verify() error  { return nil }

type OpLabel struct {
	IdResult Id
}

func (op *OpLabel) Opcode() uint32 { return opcodeLabel }
func (op *OpLabel) Optional() bool { return false }
func (op *OpLabel) Verify() error  { return nil }

type OpBranch struct {
	TargetLabel Id
}

func (op *OpBranch) Opcode() uint32 { return opcodeBranch }
func (op *OpBranch) Optional() bool { return false }
func (op *OpBranch) Verify() error  { return nil }

type OpBranchConditional struct {
	Condition     Id
	TrueLabel     Id
	FalseLabel    Id
	BranchWeights []uint32 `spirv:"optional"`
}

func (op *OpBranchConditional) Opcode() uint32 { return opcodeBranchConditional }
func (op *OpBranchConditional) Optional() bool { return false }
func (op *OpBranchConditional) Verify() error  { return nil }

type OpSwitch struct {
	Selector Id
	Default  Id
	Target   []PairLiteralIntegerIdRef `spirv:"optional"`
}

func (op *OpSwitch) Opcode() uint32 { return opcodeSwitch }
func (op *OpSwitch) Optional() bool { return false }
func (op *OpSwitch) Verify() error  { return nil }

type OpKill struct {
}

func (op *OpKill) Opcode() uint32 { return opcodeKill }
func (op *OpKill) Optional() bool { return false }
func (op *OpKill) Verify() error  { return nil }

type OpReturn struct {
}

func (op *OpReturn) Opcode() uint32 { return opcodeReturn }
func (op *OpReturn) Optional() bool { return false }
func (op *OpReturn) Verify() error  { return nil }

type OpReturnValue struct {
	Value Id
}

func (op *OpReturnValue) Opcode() uint32 { return opcodeReturnValue }
func (op *OpReturnValue) Optional() bool { return false }
func (op *OpReturnValue) Verify() error  { return nil }

type OpUnreachable struct {
}

func (op *OpUnreachable) Opcode() uint32 { return opcodeUnreachable }
func (op *OpUnreachable) Optional() bool { return false }
func (op *OpUnreachable) Verify() error  { return nil }

type OpLifetimeStart struct {
	Pointer Id
	Size    uint32
}

func (op *OpLifetimeStart) Opcode() uint32 { return opcodeLifetimeStart }
func (op *OpLifetimeStart) Optional() bool { return false }
func (op *OpLifetimeStart) Verify() error  { return nil }

type OpLifetimeStop struct {
	Pointer Id
	Size    uint32
}

func (op *OpLifetimeStop) Opcode() uint32 { return opcodeLifetimeStop }
func (op *OpLifetimeStop) Optional() bool { return false }
func (op *OpLifetimeStop) Verify() error  { return nil }

type OpGroupAsyncCopy struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Destination  Id
	Source       Id
	NumElements  Id
	Stride       Id
	Event        Id
}

func (op *OpGroupAsyncCopy) Opcode() uint32 { return opcodeGroupAsyncCopy }
func (op *OpGroupAsyncCopy) Optional() bool { return false }
func (op *OpGroupAsyncCopy) Verify() error  { return nil }

type OpGroupWaitEvents struct {
	Execution  Id
	NumEvents  Id
	EventsList Id
}

func (op *OpGroupWaitEvents) Opcode() uint32 { return opcodeGroupWaitEvents }
func (op *OpGroupWaitEvents) Optional() bool { return false }
func (op *OpGroupWaitEvents) Verify() error  { return nil }

type OpGroupAll struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Predicate    Id
}

func (op *OpGroupAll) Opcode() uint32 { return opcodeGroupAll }
func (op *OpGroupAll) Optional() bool { return false }
func (op *OpGroupAll) Verify() error  { return nil }

type OpGroupAny struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Predicate    Id
}

func (op *OpGroupAny) Opcode() uint32 { return opcodeGroupAny }
func (op *OpGroupAny) Optional() bool { return false }
func (op *OpGroupAny) Verify() error  { return nil }

type OpGroupBroadcast struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	LocalId      Id
}

func (op *OpGroupBroadcast) Opcode() uint32 { return opcodeGroupBroadcast }
func (op *OpGroupBroadcast) Optional() bool { return false }
func (op *OpGroupBroadcast) Verify() error  { return nil }

type OpGroupIAdd struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupIAdd) Opcode() uint32 { return opcodeGroupIAdd }
func (op *OpGroupIAdd) Optional() bool { return false }
func (op *OpGroupIAdd) Verify() error  { return nil }

type OpGroupFAdd struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupFAdd) Opcode() uint32 { return opcodeGroupFAdd }
func (op *OpGroupFAdd) Optional() bool { return false }
func (op *OpGroupFAdd) Verify() error  { return nil }

type OpGroupFMin struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupFMin) Opcode() uint32 { return opcodeGroupFMin }
func (op *OpGroupFMin) Optional() bool { return false }
func (op *OpGroupFMin) Verify() error  { return nil }

type OpGroupUMin struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupUMin) Opcode() uint32 { return opcodeGroupUMin }
func (op *OpGroupUMin) Optional() bool { return false }
func (op *OpGroupUMin) Verify() error  { return nil }

type OpGroupSMin struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupSMin) Opcode() uint32 { return opcodeGroupSMin }
func (op *OpGroupSMin) Optional() bool { return false }
func (op *OpGroupSMin) Verify() error  { return nil }

type OpGroupFMax struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupFMax) Opcode() uint32 { return opcodeGroupFMax }
func (op *OpGroupFMax) Optional() bool { return false }
func (op *OpGroupFMax) Verify() error  { return nil }

type OpGroupUMax struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupUMax) Opcode() uint32 { return opcodeGroupUMax }
func (op *OpGroupUMax) Optional() bool { return false }
func (op *OpGroupUMax) Verify() error  { return nil }

type OpGroupSMax struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupSMax) Opcode() uint32 { return opcodeGroupSMax }
func (op *OpGroupSMax) Optional() bool { return false }
func (op *OpGroupSMax) Verify() error  { return nil }

type OpReadPipe struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	Pointer         Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpReadPipe) Opcode() uint32 { return opcodeReadPipe }
func (op *OpReadPipe) Optional() bool { return false }
func (op *OpReadPipe) Verify() error  { return nil }

type OpWritePipe struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	Pointer         Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpWritePipe) Opcode() uint32 { return opcodeWritePipe }
func (op *OpWritePipe) Optional() bool { return false }
func (op *OpWritePipe) Verify() error  { return nil }

type OpReservedReadPipe struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	ReserveId       Id
	Index           Id
	Pointer         Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpReservedReadPipe) Opcode() uint32 { return opcodeReservedReadPipe }
func (op *OpReservedReadPipe) Optional() bool { return false }
func (op *OpReservedReadPipe) Verify() error  { return nil }

type OpReservedWritePipe struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	ReserveId       Id
	Index           Id
	Pointer         Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpReservedWritePipe) Opcode() uint32 { return opcodeReservedWritePipe }
func (op *OpReservedWritePipe) Optional() bool { return false }
func (op *OpReservedWritePipe) Verify() error  { return nil }

type OpReserveReadPipePackets struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	NumPackets      Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpReserveReadPipePackets) Opcode() uint32 { return opcodeReserveReadPipePackets }
func (op *OpReserveReadPipePackets) Optional() bool { return false }
func (op *OpReserveReadPipePackets) Verify() error  { return nil }

type OpReserveWritePipePackets struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	NumPackets      Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpReserveWritePipePackets) Opcode() uint32 { return opcodeReserveWritePipePackets }
func (op *OpReserveWritePipePackets) Optional() bool { return false }
func (op *OpReserveWritePipePackets) Verify() error  { return nil }

type OpCommitReadPipe struct {
	Pipe            Id
	ReserveId       Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpCommitReadPipe) Opcode() uint32 { return opcodeCommitReadPipe }
func (op *OpCommitReadPipe) Optional() bool { return false }
func (op *OpCommitReadPipe) Verify() error  { return nil }

type OpCommitWritePipe struct {
	Pipe            Id
	ReserveId       Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpCommitWritePipe) Opcode() uint32 { return opcodeCommitWritePipe }
func (op *OpCommitWritePipe) Optional() bool { return false }
func (op *OpCommitWritePipe) Verify() error  { return nil }

type OpIsValidReserveId struct {
	IdResultType Id
	IdResult     Id
	ReserveId    Id
}

func (op *OpIsValidReserveId) Opcode() uint32 { return opcodeIsValidReserveId }
func (op *OpIsValidReserveId) Optional() bool { return false }
func (op *OpIsValidReserveId) Verify() error  { return nil }

type OpGetNumPipePackets struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpGetNumPipePackets) Opcode() uint32 { return opcodeGetNumPipePackets }
func (op *OpGetNumPipePackets) Optional() bool { return false }
func (op *OpGetNumPipePackets) Verify() error  { return nil }

type OpGetMaxPipePackets struct {
	IdResultType    Id
	IdResult        Id
	Pipe            Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpGetMaxPipePackets) Opcode() uint32 { return opcodeGetMaxPipePackets }
func (op *OpGetMaxPipePackets) Optional() bool { return false }
func (op *OpGetMaxPipePackets) Verify() error  { return nil }

type OpGroupReserveReadPipePackets struct {
	IdResultType    Id
	IdResult        Id
	Execution       Id
	Pipe            Id
	NumPackets      Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpGroupReserveReadPipePackets) Opcode() uint32 { return opcodeGroupReserveReadPipePackets }
func (op *OpGroupReserveReadPipePackets) Optional() bool { return false }
func (op *OpGroupReserveReadPipePackets) Verify() error  { return nil }

type OpGroupReserveWritePipePackets struct {
	IdResultType    Id
	IdResult        Id
	Execution       Id
	Pipe            Id
	NumPackets      Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpGroupReserveWritePipePackets) Opcode() uint32 { return opcodeGroupReserveWritePipePackets }
func (op *OpGroupReserveWritePipePackets) Optional() bool { return false }
func (op *OpGroupReserveWritePipePackets) Verify() error  { return nil }

type OpGroupCommitReadPipe struct {
	Execution       Id
	Pipe            Id
	ReserveId       Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpGroupCommitReadPipe) Opcode() uint32 { return opcodeGroupCommitReadPipe }
func (op *OpGroupCommitReadPipe) Optional() bool { return false }
func (op *OpGroupCommitReadPipe) Verify() error  { return nil }

type OpGroupCommitWritePipe struct {
	Execution       Id
	Pipe            Id
	ReserveId       Id
	PacketSize      Id
	PacketAlignment Id
}

func (op *OpGroupCommitWritePipe) Opcode() uint32 { return opcodeGroupCommitWritePipe }
func (op *OpGroupCommitWritePipe) Optional() bool { return false }
func (op *OpGroupCommitWritePipe) Verify() error  { return nil }

type OpEnqueueMarker struct {
	IdResultType Id
	IdResult     Id
	Queue        Id
	NumEvents    Id
	WaitEvents   Id
	RetEvent     Id
}

func (op *OpEnqueueMarker) Opcode() uint32 { return opcodeEnqueueMarker }
func (op *OpEnqueueMarker) Optional() bool { return false }
func (op *OpEnqueueMarker) Verify() error  { return nil }

type OpEnqueueKernel struct {
	IdResultType Id
	IdResult     Id
	Queue        Id
	Flags        Id
	NDRange      Id
	NumEvents    Id
	WaitEvents   Id
	RetEvent     Id
	Invoke       Id
	Param        Id
	ParamSize    Id
	ParamAlign   Id
	LocalSize    []Id `spirv:"optional"`
}

func (op *OpEnqueueKernel) Opcode() uint32 { return opcodeEnqueueKernel }
func (op *OpEnqueueKernel) Optional() bool { return false }
func (op *OpEnqueueKernel) Verify() error  { return nil }

type OpGetKernelNDrangeSubGroupCount struct {
	IdResultType Id
	IdResult     Id
	NDRange      Id
	Invoke       Id
	Param        Id
	ParamSize    Id
	ParamAlign   Id
}

func (op *OpGetKernelNDrangeSubGroupCount) Opcode() uint32 { return opcodeGetKernelNDrangeSubGroupCount }
func (op *OpGetKernelNDrangeSubGroupCount) Optional() bool { return false }
func (op *OpGetKernelNDrangeSubGroupCount) Verify() error  { return nil }

type OpGetKernelNDrangeMaxSubGroupSize struct {
	IdResultType Id
	IdResult     Id
	NDRange      Id
	Invoke       Id
	Param        Id
	ParamSize    Id
	ParamAlign   Id
}

func (op *OpGetKernelNDrangeMaxSubGroupSize) Opcode() uint32 {
	return opcodeGetKernelNDrangeMaxSubGroupSize
}
func (op *OpGetKernelNDrangeMaxSubGroupSize) Optional() bool { return false }
func (op *OpGetKernelNDrangeMaxSubGroupSize) Verify() error  { return nil }

type OpGetKernelWorkGroupSize struct {
	IdResultType Id
	IdResult     Id
	Invoke       Id
	Param        Id
	ParamSize    Id
	ParamAlign   Id
}

func (op *OpGetKernelWorkGroupSize) Opcode() uint32 { return opcodeGetKernelWorkGroupSize }
func (op *OpGetKernelWorkGroupSize) Optional() bool { return false }
func (op *OpGetKernelWorkGroupSize) Verify() error  { return nil }

type OpGetKernelPreferredWorkGroupSizeMultiple struct {
	IdResultType Id
	IdResult     Id
	Invoke       Id
	Param        Id
	ParamSize    Id
	ParamAlign   Id
}

func (op *OpGetKernelPreferredWorkGroupSizeMultiple) Opcode() uint32 {
	return opcodeGetKernelPreferredWorkGroupSizeMultiple
}
func (op *OpGetKernelPreferredWorkGroupSizeMultiple) Optional() bool { return false }
func (op *OpGetKernelPreferredWorkGroupSizeMultiple) Verify() error  { return nil }

type OpRetainEvent struct {
	Event Id
}

func (op *OpRetainEvent) Opcode() uint32 { return opcodeRetainEvent }
func (op *OpRetainEvent) Optional() bool { return false }
func (op *OpRetainEvent) Verify() error  { return nil }

type OpReleaseEvent struct {
	Event Id
}

func (op *OpReleaseEvent) Opcode() uint32 { return opcodeReleaseEvent }
func (op *OpReleaseEvent) Optional() bool { return false }
func (op *OpReleaseEvent) Verify() error  { return nil }

type OpCreateUserEvent struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpCreateUserEvent) Opcode() uint32 { return opcodeCreateUserEvent }
func (op *OpCreateUserEvent) Optional() bool { return false }
func (op *OpCreateUserEvent) Verify() error  { return nil }

type OpIsValidEvent struct {
	IdResultType Id
	IdResult     Id
	Event        Id
}

func (op *OpIsValidEvent) Opcode() uint32 { return opcodeIsValidEvent }
func (op *OpIsValidEvent) Optional() bool { return false }
func (op *OpIsValidEvent) Verify() error  { return nil }

type OpSetUserEventStatus struct {
	Event  Id
	Status Id
}

func (op *OpSetUserEventStatus) Opcode() uint32 { return opcodeSetUserEventStatus }
func (op *OpSetUserEventStatus) Optional() bool { return false }
func (op *OpSetUserEventStatus) Verify() error  { return nil }

type OpCaptureEventProfilingInfo struct {
	Event         Id
	ProfilingInfo Id
	Value         Id
}

func (op *OpCaptureEventProfilingInfo) Opcode() uint32 { return opcodeCaptureEventProfilingInfo }
func (op *OpCaptureEventProfilingInfo) Optional() bool { return false }
func (op *OpCaptureEventProfilingInfo) Verify() error  { return nil }

type OpGetDefaultQueue struct {
	IdResultType Id
	IdResult     Id
}

func (op *OpGetDefaultQueue) Opcode() uint32 { return opcodeGetDefaultQueue }
func (op *OpGetDefaultQueue) Optional() bool { return false }
func (op *OpGetDefaultQueue) Verify() error  { return nil }

type OpBuildNDRange struct {
	IdResultType     Id
	IdResult         Id
	GlobalWorkSize   Id
	LocalWorkSize    Id
	GlobalWorkOffset Id
}

func (op *OpBuildNDRange) Opcode() uint32 { return opcodeBuildNDRange }
func (op *OpBuildNDRange) Optional() bool { return false }
func (op *OpBuildNDRange) Verify() error  { return nil }

type OpImageSparseSampleImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseSampleImplicitLod) Opcode() uint32 { return opcodeImageSparseSampleImplicitLod }
func (op *OpImageSparseSampleImplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleImplicitLod) Verify() error  { return nil }

type OpImageSparseSampleExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands
}

func (op *OpImageSparseSampleExplicitLod) Opcode() uint32 { return opcodeImageSparseSampleExplicitLod }
func (op *OpImageSparseSampleExplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleExplicitLod) Verify() error  { return nil }

type OpImageSparseSampleDrefImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseSampleDrefImplicitLod) Opcode() uint32 {
	return opcodeImageSparseSampleDrefImplicitLod
}
func (op *OpImageSparseSampleDrefImplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleDrefImplicitLod) Verify() error  { return nil }

type OpImageSparseSampleDrefExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands
}

func (op *OpImageSparseSampleDrefExplicitLod) Opcode() uint32 {
	return opcodeImageSparseSampleDrefExplicitLod
}
func (op *OpImageSparseSampleDrefExplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleDrefExplicitLod) Verify() error  { return nil }

type OpImageSparseSampleProjImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseSampleProjImplicitLod) Opcode() uint32 {
	return opcodeImageSparseSampleProjImplicitLod
}
func (op *OpImageSparseSampleProjImplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleProjImplicitLod) Verify() error  { return nil }

type OpImageSparseSampleProjExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	ImageOperands ImageOperands
}

func (op *OpImageSparseSampleProjExplicitLod) Opcode() uint32 {
	return opcodeImageSparseSampleProjExplicitLod
}
func (op *OpImageSparseSampleProjExplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleProjExplicitLod) Verify() error  { return nil }

type OpImageSparseSampleProjDrefImplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseSampleProjDrefImplicitLod) Opcode() uint32 {
	return opcodeImageSparseSampleProjDrefImplicitLod
}
func (op *OpImageSparseSampleProjDrefImplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleProjDrefImplicitLod) Verify() error  { return nil }

type OpImageSparseSampleProjDrefExplicitLod struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands
}

func (op *OpImageSparseSampleProjDrefExplicitLod) Opcode() uint32 {
	return opcodeImageSparseSampleProjDrefExplicitLod
}
func (op *OpImageSparseSampleProjDrefExplicitLod) Optional() bool { return false }
func (op *OpImageSparseSampleProjDrefExplicitLod) Verify() error  { return nil }

type OpImageSparseFetch struct {
	IdResultType  Id
	IdResult      Id
	Image         Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseFetch) Opcode() uint32 { return opcodeImageSparseFetch }
func (op *OpImageSparseFetch) Optional() bool { return false }
func (op *OpImageSparseFetch) Verify() error  { return nil }

type OpImageSparseGather struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	Component     Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseGather) Opcode() uint32 { return opcodeImageSparseGather }
func (op *OpImageSparseGather) Optional() bool { return false }
func (op *OpImageSparseGather) Verify() error  { return nil }

type OpImageSparseDrefGather struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	D_ref_        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseDrefGather) Opcode() uint32 { return opcodeImageSparseDrefGather }
func (op *OpImageSparseDrefGather) Optional() bool { return false }
func (op *OpImageSparseDrefGather) Verify() error  { return nil }

type OpImageSparseTexelsResident struct {
	IdResultType Id
	IdResult     Id
	ResidentCode Id
}

func (op *OpImageSparseTexelsResident) Opcode() uint32 { return opcodeImageSparseTexelsResident }
func (op *OpImageSparseTexelsResident) Optional() bool { return false }
func (op *OpImageSparseTexelsResident) Verify() error  { return nil }

type OpNoLine struct {
}

func (op *OpNoLine) Opcode() uint32 { return opcodeNoLine }
func (op *OpNoLine) Optional() bool { return true }
func (op *OpNoLine) Verify() error  { return nil }

type OpAtomicFlagTestAndSet struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
	Memory       Id
	Semantics    Id
}

func (op *OpAtomicFlagTestAndSet) Opcode() uint32 { return opcodeAtomicFlagTestAndSet }
func (op *OpAtomicFlagTestAndSet) Optional() bool { return false }
func (op *OpAtomicFlagTestAndSet) Verify() error  { return nil }

type OpAtomicFlagClear struct {
	Pointer   Id
	Memory    Id
	Semantics Id
}

func (op *OpAtomicFlagClear) Opcode() uint32 { return opcodeAtomicFlagClear }
func (op *OpAtomicFlagClear) Optional() bool { return false }
func (op *OpAtomicFlagClear) Verify() error  { return nil }

type OpImageSparseRead struct {
	IdResultType  Id
	IdResult      Id
	Image         Id
	Coordinate    Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSparseRead) Opcode() uint32 { return opcodeImageSparseRead }
func (op *OpImageSparseRead) Optional() bool { return false }
func (op *OpImageSparseRead) Verify() error  { return nil }

type OpSizeOf struct {
	IdResultType Id
	IdResult     Id
	Pointer      Id
}

func (op *OpSizeOf) Opcode() uint32 { return opcodeSizeOf }
func (op *OpSizeOf) Optional() bool { return false }
func (op *OpSizeOf) Verify() error  { return nil }

type OpTypePipeStorage struct {
	IdResult Id
}

func (op *OpTypePipeStorage) Opcode() uint32 { return opcodeTypePipeStorage }
func (op *OpTypePipeStorage) Optional() bool { return false }
func (op *OpTypePipeStorage) Verify() error  { return nil }

type OpConstantPipeStorage struct {
	IdResultType    Id
	IdResult        Id
	PacketSize      uint32
	PacketAlignment uint32
	Capacity        uint32
}

func (op *OpConstantPipeStorage) Opcode() uint32 { return opcodeConstantPipeStorage }
func (op *OpConstantPipeStorage) Optional() bool { return false }
func (op *OpConstantPipeStorage) Verify() error  { return nil }

type OpCreatePipeFromPipeStorage struct {
	IdResultType Id
	IdResult     Id
	PipeStorage  Id
}

func (op *OpCreatePipeFromPipeStorage) Opcode() uint32 { return opcodeCreatePipeFromPipeStorage }
func (op *OpCreatePipeFromPipeStorage) Optional() bool { return false }
func (op *OpCreatePipeFromPipeStorage) Verify() error  { return nil }

type OpGetKernelLocalSizeForSubgroupCount struct {
	IdResultType  Id
	IdResult      Id
	SubgroupCount Id
	Invoke        Id
	Param         Id
	ParamSize     Id
	ParamAlign    Id
}

func (op *OpGetKernelLocalSizeForSubgroupCount) Opcode() uint32 {
	return opcodeGetKernelLocalSizeForSubgroupCount
}
func (op *OpGetKernelLocalSizeForSubgroupCount) Optional() bool { return false }
func (op *OpGetKernelLocalSizeForSubgroupCount) Verify() error  { return nil }

type OpGetKernelMaxNumSubgroups struct {
	IdResultType Id
	IdResult     Id
	Invoke       Id
	Param        Id
	ParamSize    Id
	ParamAlign   Id
}

func (op *OpGetKernelMaxNumSubgroups) Opcode() uint32 { return opcodeGetKernelMaxNumSubgroups }
func (op *OpGetKernelMaxNumSubgroups) Optional() bool { return false }
func (op *OpGetKernelMaxNumSubgroups) Verify() error  { return nil }

type OpTypeNamedBarrier struct {
	IdResult Id
}

func (op *OpTypeNamedBarrier) Opcode() uint32 { return opcodeTypeNamedBarrier }
func (op *OpTypeNamedBarrier) Optional() bool { return false }
func (op *OpTypeNamedBarrier) Verify() error  { return nil }

type OpNamedBarrierInitialize struct {
	IdResultType  Id
	IdResult      Id
	SubgroupCount Id
}

func (op *OpNamedBarrierInitialize) Opcode() uint32 { return opcodeNamedBarrierInitialize }
func (op *OpNamedBarrierInitialize) Optional() bool { return false }
func (op *OpNamedBarrierInitialize) Verify() error  { return nil }

type OpMemoryNamedBarrier struct {
	NamedBarrier Id
	Memory       Id
	Semantics    Id
}

func (op *OpMemoryNamedBarrier) Opcode() uint32 { return opcodeMemoryNamedBarrier }
func (op *OpMemoryNamedBarrier) Optional() bool { return false }
func (op *OpMemoryNamedBarrier) Verify() error  { return nil }

type OpModuleProcessed struct {
	Process String
}

func (op *OpModuleProcessed) Opcode() uint32 { return opcodeModuleProcessed }
func (op *OpModuleProcessed) Optional() bool { return true }
func (op *OpModuleProcessed) Verify() error  { return nil }

type OpExecutionModeId struct {
	EntryPoint Id
	Mode       ExecutionMode
}

func (op *OpExecutionModeId) Opcode() uint32 { return opcodeExecutionModeId }
func (op *OpExecutionModeId) Optional() bool { return false }
func (op *OpExecutionModeId) Verify() error  { return nil }

type OpDecorateId struct {
	Target     Id
	Decoration Decoration
}

func (op *OpDecorateId) Opcode() uint32 { return opcodeDecorateId }
func (op *OpDecorateId) Optional() bool { return false }
func (op *OpDecorateId) Verify() error  { return nil }

type OpGroupNonUniformElect struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
}

func (op *OpGroupNonUniformElect) Opcode() uint32 { return opcodeGroupNonUniformElect }
func (op *OpGroupNonUniformElect) Optional() bool { return false }
func (op *OpGroupNonUniformElect) Verify() error  { return nil }

type OpGroupNonUniformAll struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Predicate    Id
}

func (op *OpGroupNonUniformAll) Opcode() uint32 { return opcodeGroupNonUniformAll }
func (op *OpGroupNonUniformAll) Optional() bool { return false }
func (op *OpGroupNonUniformAll) Verify() error  { return nil }

type OpGroupNonUniformAny struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Predicate    Id
}

func (op *OpGroupNonUniformAny) Opcode() uint32 { return opcodeGroupNonUniformAny }
func (op *OpGroupNonUniformAny) Optional() bool { return false }
func (op *OpGroupNonUniformAny) Verify() error  { return nil }

type OpGroupNonUniformAllEqual struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
}

func (op *OpGroupNonUniformAllEqual) Opcode() uint32 { return opcodeGroupNonUniformAllEqual }
func (op *OpGroupNonUniformAllEqual) Optional() bool { return false }
func (op *OpGroupNonUniformAllEqual) Verify() error  { return nil }

type OpGroupNonUniformBroadcast struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Id           Id
}

func (op *OpGroupNonUniformBroadcast) Opcode() uint32 { return opcodeGroupNonUniformBroadcast }
func (op *OpGroupNonUniformBroadcast) Optional() bool { return false }
func (op *OpGroupNonUniformBroadcast) Verify() error  { return nil }

type OpGroupNonUniformBroadcastFirst struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
}

func (op *OpGroupNonUniformBroadcastFirst) Opcode() uint32 { return opcodeGroupNonUniformBroadcastFirst }
func (op *OpGroupNonUniformBroadcastFirst) Optional() bool { return false }
func (op *OpGroupNonUniformBroadcastFirst) Verify() error  { return nil }

type OpGroupNonUniformBallot struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Predicate    Id
}

func (op *OpGroupNonUniformBallot) Opcode() uint32 { return opcodeGroupNonUniformBallot }
func (op *OpGroupNonUniformBallot) Optional() bool { return false }
func (op *OpGroupNonUniformBallot) Verify() error  { return nil }

type OpGroupNonUniformInverseBallot struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
}

func (op *OpGroupNonUniformInverseBallot) Opcode() uint32 { return opcodeGroupNonUniformInverseBallot }
func (op *OpGroupNonUniformInverseBallot) Optional() bool { return false }
func (op *OpGroupNonUniformInverseBallot) Verify() error  { return nil }

type OpGroupNonUniformBallotBitExtract struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Index        Id
}

func (op *OpGroupNonUniformBallotBitExtract) Opcode() uint32 {
	return opcodeGroupNonUniformBallotBitExtract
}
func (op *OpGroupNonUniformBallotBitExtract) Optional() bool { return false }
func (op *OpGroupNonUniformBallotBitExtract) Verify() error  { return nil }

type OpGroupNonUniformBallotBitCount struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
}

func (op *OpGroupNonUniformBallotBitCount) Opcode() uint32 { return opcodeGroupNonUniformBallotBitCount }
func (op *OpGroupNonUniformBallotBitCount) Optional() bool { return false }
func (op *OpGroupNonUniformBallotBitCount) Verify() error  { return nil }

type OpGroupNonUniformBallotFindLSB struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
}

func (op *OpGroupNonUniformBallotFindLSB) Opcode() uint32 { return opcodeGroupNonUniformBallotFindLSB }
func (op *OpGroupNonUniformBallotFindLSB) Optional() bool { return false }
func (op *OpGroupNonUniformBallotFindLSB) Verify() error  { return nil }

type OpGroupNonUniformBallotFindMSB struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
}

func (op *OpGroupNonUniformBallotFindMSB) Opcode() uint32 { return opcodeGroupNonUniformBallotFindMSB }
func (op *OpGroupNonUniformBallotFindMSB) Optional() bool { return false }
func (op *OpGroupNonUniformBallotFindMSB) Verify() error  { return nil }

type OpGroupNonUniformShuffle struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Id           Id
}

func (op *OpGroupNonUniformShuffle) Opcode() uint32 { return opcodeGroupNonUniformShuffle }
func (op *OpGroupNonUniformShuffle) Optional() bool { return false }
func (op *OpGroupNonUniformShuffle) Verify() error  { return nil }

type OpGroupNonUniformShuffleXor struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Mask         Id
}

func (op *OpGroupNonUniformShuffleXor) Opcode() uint32 { return opcodeGroupNonUniformShuffleXor }
func (op *OpGroupNonUniformShuffleXor) Optional() bool { return false }
func (op *OpGroupNonUniformShuffleXor) Verify() error  { return nil }

type OpGroupNonUniformShuffleUp struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Delta        Id
}

func (op *OpGroupNonUniformShuffleUp) Opcode() uint32 { return opcodeGroupNonUniformShuffleUp }
func (op *OpGroupNonUniformShuffleUp) Optional() bool { return false }
func (op *OpGroupNonUniformShuffleUp) Verify() error  { return nil }

type OpGroupNonUniformShuffleDown struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Delta        Id
}

func (op *OpGroupNonUniformShuffleDown) Opcode() uint32 { return opcodeGroupNonUniformShuffleDown }
func (op *OpGroupNonUniformShuffleDown) Optional() bool { return false }
func (op *OpGroupNonUniformShuffleDown) Verify() error  { return nil }

type OpGroupNonUniformIAdd struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformIAdd) Opcode() uint32 { return opcodeGroupNonUniformIAdd }
func (op *OpGroupNonUniformIAdd) Optional() bool { return false }
func (op *OpGroupNonUniformIAdd) Verify() error  { return nil }

type OpGroupNonUniformFAdd struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformFAdd) Opcode() uint32 { return opcodeGroupNonUniformFAdd }
func (op *OpGroupNonUniformFAdd) Optional() bool { return false }
func (op *OpGroupNonUniformFAdd) Verify() error  { return nil }

type OpGroupNonUniformIMul struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformIMul) Opcode() uint32 { return opcodeGroupNonUniformIMul }
func (op *OpGroupNonUniformIMul) Optional() bool { return false }
func (op *OpGroupNonUniformIMul) Verify() error  { return nil }

type OpGroupNonUniformFMul struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformFMul) Opcode() uint32 { return opcodeGroupNonUniformFMul }
func (op *OpGroupNonUniformFMul) Optional() bool { return false }
func (op *OpGroupNonUniformFMul) Verify() error  { return nil }

type OpGroupNonUniformSMin struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformSMin) Opcode() uint32 { return opcodeGroupNonUniformSMin }
func (op *OpGroupNonUniformSMin) Optional() bool { return false }
func (op *OpGroupNonUniformSMin) Verify() error  { return nil }

type OpGroupNonUniformUMin struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformUMin) Opcode() uint32 { return opcodeGroupNonUniformUMin }
func (op *OpGroupNonUniformUMin) Optional() bool { return false }
func (op *OpGroupNonUniformUMin) Verify() error  { return nil }

type OpGroupNonUniformFMin struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformFMin) Opcode() uint32 { return opcodeGroupNonUniformFMin }
func (op *OpGroupNonUniformFMin) Optional() bool { return false }
func (op *OpGroupNonUniformFMin) Verify() error  { return nil }

type OpGroupNonUniformSMax struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformSMax) Opcode() uint32 { return opcodeGroupNonUniformSMax }
func (op *OpGroupNonUniformSMax) Optional() bool { return false }
func (op *OpGroupNonUniformSMax) Verify() error  { return nil }

type OpGroupNonUniformUMax struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformUMax) Opcode() uint32 { return opcodeGroupNonUniformUMax }
func (op *OpGroupNonUniformUMax) Optional() bool { return false }
func (op *OpGroupNonUniformUMax) Verify() error  { return nil }

type OpGroupNonUniformFMax struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformFMax) Opcode() uint32 { return opcodeGroupNonUniformFMax }
func (op *OpGroupNonUniformFMax) Optional() bool { return false }
func (op *OpGroupNonUniformFMax) Verify() error  { return nil }

type OpGroupNonUniformBitwiseAnd struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformBitwiseAnd) Opcode() uint32 { return opcodeGroupNonUniformBitwiseAnd }
func (op *OpGroupNonUniformBitwiseAnd) Optional() bool { return false }
func (op *OpGroupNonUniformBitwiseAnd) Verify() error  { return nil }

type OpGroupNonUniformBitwiseOr struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformBitwiseOr) Opcode() uint32 { return opcodeGroupNonUniformBitwiseOr }
func (op *OpGroupNonUniformBitwiseOr) Optional() bool { return false }
func (op *OpGroupNonUniformBitwiseOr) Verify() error  { return nil }

type OpGroupNonUniformBitwiseXor struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformBitwiseXor) Opcode() uint32 { return opcodeGroupNonUniformBitwiseXor }
func (op *OpGroupNonUniformBitwiseXor) Optional() bool { return false }
func (op *OpGroupNonUniformBitwiseXor) Verify() error  { return nil }

type OpGroupNonUniformLogicalAnd struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformLogicalAnd) Opcode() uint32 { return opcodeGroupNonUniformLogicalAnd }
func (op *OpGroupNonUniformLogicalAnd) Optional() bool { return false }
func (op *OpGroupNonUniformLogicalAnd) Verify() error  { return nil }

type OpGroupNonUniformLogicalOr struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformLogicalOr) Opcode() uint32 { return opcodeGroupNonUniformLogicalOr }
func (op *OpGroupNonUniformLogicalOr) Optional() bool { return false }
func (op *OpGroupNonUniformLogicalOr) Verify() error  { return nil }

type OpGroupNonUniformLogicalXor struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	Value        Id
	ClusterSize  Id `spirv:"optional"`
}

func (op *OpGroupNonUniformLogicalXor) Opcode() uint32 { return opcodeGroupNonUniformLogicalXor }
func (op *OpGroupNonUniformLogicalXor) Optional() bool { return false }
func (op *OpGroupNonUniformLogicalXor) Verify() error  { return nil }

type OpGroupNonUniformQuadBroadcast struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Index        Id
}

func (op *OpGroupNonUniformQuadBroadcast) Opcode() uint32 { return opcodeGroupNonUniformQuadBroadcast }
func (op *OpGroupNonUniformQuadBroadcast) Optional() bool { return false }
func (op *OpGroupNonUniformQuadBroadcast) Verify() error  { return nil }

type OpGroupNonUniformQuadSwap struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Value        Id
	Direction    Id
}

func (op *OpGroupNonUniformQuadSwap) Opcode() uint32 { return opcodeGroupNonUniformQuadSwap }
func (op *OpGroupNonUniformQuadSwap) Optional() bool { return false }
func (op *OpGroupNonUniformQuadSwap) Verify() error  { return nil }

type OpCopyLogical struct {
	IdResultType Id
	IdResult     Id
	Operand      Id
}

func (op *OpCopyLogical) Opcode() uint32 { return opcodeCopyLogical }
func (op *OpCopyLogical) Optional() bool { return false }
func (op *OpCopyLogical) Verify() error  { return nil }

type OpPtrEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpPtrEqual) Opcode() uint32 { return opcodePtrEqual }
func (op *OpPtrEqual) Optional() bool { return false }
func (op *OpPtrEqual) Verify() error  { return nil }

type OpPtrNotEqual struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpPtrNotEqual) Opcode() uint32 { return opcodePtrNotEqual }
func (op *OpPtrNotEqual) Optional() bool { return false }
func (op *OpPtrNotEqual) Verify() error  { return nil }

type OpPtrDiff struct {
	IdResultType Id
	IdResult     Id
	Operand1     Id
	Operand2     Id
}

func (op *OpPtrDiff) Opcode() uint32 { return opcodePtrDiff }
func (op *OpPtrDiff) Optional() bool { return false }
func (op *OpPtrDiff) Verify() error  { return nil }

type OpSubgroupBallotKHR struct {
	IdResultType Id
	IdResult     Id
	Predicate    Id
}

func (op *OpSubgroupBallotKHR) Opcode() uint32 { return opcodeSubgroupBallotKHR }
func (op *OpSubgroupBallotKHR) Optional() bool { return false }
func (op *OpSubgroupBallotKHR) Verify() error  { return nil }

type OpSubgroupFirstInvocationKHR struct {
	IdResultType Id
	IdResult     Id
	Value        Id
}

func (op *OpSubgroupFirstInvocationKHR) Opcode() uint32 { return opcodeSubgroupFirstInvocationKHR }
func (op *OpSubgroupFirstInvocationKHR) Optional() bool { return false }
func (op *OpSubgroupFirstInvocationKHR) Verify() error  { return nil }

type OpSubgroupAllKHR struct {
	IdResultType Id
	IdResult     Id
	Predicate    Id
}

func (op *OpSubgroupAllKHR) Opcode() uint32 { return opcodeSubgroupAllKHR }
func (op *OpSubgroupAllKHR) Optional() bool { return false }
func (op *OpSubgroupAllKHR) Verify() error  { return nil }

type OpSubgroupAnyKHR struct {
	IdResultType Id
	IdResult     Id
	Predicate    Id
}

func (op *OpSubgroupAnyKHR) Opcode() uint32 { return opcodeSubgroupAnyKHR }
func (op *OpSubgroupAnyKHR) Optional() bool { return false }
func (op *OpSubgroupAnyKHR) Verify() error  { return nil }

type OpSubgroupAllEqualKHR struct {
	IdResultType Id
	IdResult     Id
	Predicate    Id
}

func (op *OpSubgroupAllEqualKHR) Opcode() uint32 { return opcodeSubgroupAllEqualKHR }
func (op *OpSubgroupAllEqualKHR) Optional() bool { return false }
func (op *OpSubgroupAllEqualKHR) Verify() error  { return nil }

type OpSubgroupReadInvocationKHR struct {
	IdResultType Id
	IdResult     Id
	Value        Id
	Index        Id
}

func (op *OpSubgroupReadInvocationKHR) Opcode() uint32 { return opcodeSubgroupReadInvocationKHR }
func (op *OpSubgroupReadInvocationKHR) Optional() bool { return false }
func (op *OpSubgroupReadInvocationKHR) Verify() error  { return nil }

type OpGroupIAddNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupIAddNonUniformAMD) Opcode() uint32 { return opcodeGroupIAddNonUniformAMD }
func (op *OpGroupIAddNonUniformAMD) Optional() bool { return false }
func (op *OpGroupIAddNonUniformAMD) Verify() error  { return nil }

type OpGroupFAddNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupFAddNonUniformAMD) Opcode() uint32 { return opcodeGroupFAddNonUniformAMD }
func (op *OpGroupFAddNonUniformAMD) Optional() bool { return false }
func (op *OpGroupFAddNonUniformAMD) Verify() error  { return nil }

type OpGroupFMinNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupFMinNonUniformAMD) Opcode() uint32 { return opcodeGroupFMinNonUniformAMD }
func (op *OpGroupFMinNonUniformAMD) Optional() bool { return false }
func (op *OpGroupFMinNonUniformAMD) Verify() error  { return nil }

type OpGroupUMinNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupUMinNonUniformAMD) Opcode() uint32 { return opcodeGroupUMinNonUniformAMD }
func (op *OpGroupUMinNonUniformAMD) Optional() bool { return false }
func (op *OpGroupUMinNonUniformAMD) Verify() error  { return nil }

type OpGroupSMinNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupSMinNonUniformAMD) Opcode() uint32 { return opcodeGroupSMinNonUniformAMD }
func (op *OpGroupSMinNonUniformAMD) Optional() bool { return false }
func (op *OpGroupSMinNonUniformAMD) Verify() error  { return nil }

type OpGroupFMaxNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupFMaxNonUniformAMD) Opcode() uint32 { return opcodeGroupFMaxNonUniformAMD }
func (op *OpGroupFMaxNonUniformAMD) Optional() bool { return false }
func (op *OpGroupFMaxNonUniformAMD) Verify() error  { return nil }

type OpGroupUMaxNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupUMaxNonUniformAMD) Opcode() uint32 { return opcodeGroupUMaxNonUniformAMD }
func (op *OpGroupUMaxNonUniformAMD) Optional() bool { return false }
func (op *OpGroupUMaxNonUniformAMD) Verify() error  { return nil }

type OpGroupSMaxNonUniformAMD struct {
	IdResultType Id
	IdResult     Id
	Execution    Id
	Operation    GroupOperation
	X            Id
}

func (op *OpGroupSMaxNonUniformAMD) Opcode() uint32 { return opcodeGroupSMaxNonUniformAMD }
func (op *OpGroupSMaxNonUniformAMD) Optional() bool { return false }
func (op *OpGroupSMaxNonUniformAMD) Verify() error  { return nil }

type OpImageSampleFootprintNV struct {
	IdResultType  Id
	IdResult      Id
	SampledImage  Id
	Coordinate    Id
	Granularity   Id
	Coarse        Id
	ImageOperands ImageOperands `spirv:"optional"`
}

func (op *OpImageSampleFootprintNV) Opcode() uint32 { return opcodeImageSampleFootprintNV }
func (op *OpImageSampleFootprintNV) Optional() bool { return false }
func (op *OpImageSampleFootprintNV) Verify() error  { return nil }

type OpGroupNonUniformPartitionNV struct {
	IdResultType Id
	IdResult     Id
	Value        Id
}

func (op *OpGroupNonUniformPartitionNV) Opcode() uint32 { return opcodeGroupNonUniformPartitionNV }
func (op *OpGroupNonUniformPartitionNV) Optional() bool { return false }
func (op *OpGroupNonUniformPartitionNV) Verify() error  { return nil }

type OpSubgroupShuffleINTEL struct {
	IdResultType Id
	IdResult     Id
	Data         Id
	InvocationId Id
}

func (op *OpSubgroupShuffleINTEL) Opcode() uint32 { return opcodeSubgroupShuffleINTEL }
func (op *OpSubgroupShuffleINTEL) Optional() bool { return false }
func (op *OpSubgroupShuffleINTEL) Verify() error  { return nil }

type OpSubgroupShuffleDownINTEL struct {
	IdResultType Id
	IdResult     Id
	Current      Id
	Next         Id
	Delta        Id
}

func (op *OpSubgroupShuffleDownINTEL) Opcode() uint32 { return opcodeSubgroupShuffleDownINTEL }
func (op *OpSubgroupShuffleDownINTEL) Optional() bool { return false }
func (op *OpSubgroupShuffleDownINTEL) Verify() error  { return nil }

type OpSubgroupShuffleUpINTEL struct {
	IdResultType Id
	IdResult     Id
	Previous     Id
	Current      Id
	Delta        Id
}

func (op *OpSubgroupShuffleUpINTEL) Opcode() uint32 { return opcodeSubgroupShuffleUpINTEL }
func (op *OpSubgroupShuffleUpINTEL) Optional() bool { return false }
func (op *OpSubgroupShuffleUpINTEL) Verify() error  { return nil }

type OpSubgroupShuffleXorINTEL struct {
	IdResultType Id
	IdResult     Id
	Data         Id
	Value        Id
}

func (op *OpSubgroupShuffleXorINTEL) Opcode() uint32 { return opcodeSubgroupShuffleXorINTEL }
func (op *OpSubgroupShuffleXorINTEL) Optional() bool { return false }
func (op *OpSubgroupShuffleXorINTEL) Verify() error  { return nil }

type OpSubgroupBlockReadINTEL struct {
	IdResultType Id
	IdResult     Id
	Ptr          Id
}

func (op *OpSubgroupBlockReadINTEL) Opcode() uint32 { return opcodeSubgroupBlockReadINTEL }
func (op *OpSubgroupBlockReadINTEL) Optional() bool { return false }
func (op *OpSubgroupBlockReadINTEL) Verify() error  { return nil }

type OpSubgroupBlockWriteINTEL struct {
	Ptr  Id
	Data Id
}

func (op *OpSubgroupBlockWriteINTEL) Opcode() uint32 { return opcodeSubgroupBlockWriteINTEL }
func (op *OpSubgroupBlockWriteINTEL) Optional() bool { return false }
func (op *OpSubgroupBlockWriteINTEL) Verify() error  { return nil }

type OpSubgroupImageBlockReadINTEL struct {
	IdResultType Id
	IdResult     Id
	Image        Id
	Coordinate   Id
}

func (op *OpSubgroupImageBlockReadINTEL) Opcode() uint32 { return opcodeSubgroupImageBlockReadINTEL }
func (op *OpSubgroupImageBlockReadINTEL) Optional() bool { return false }
func (op *OpSubgroupImageBlockReadINTEL) Verify() error  { return nil }

type OpSubgroupImageBlockWriteINTEL struct {
	Image      Id
	Coordinate Id
	Data       Id
}

func (op *OpSubgroupImageBlockWriteINTEL) Opcode() uint32 { return opcodeSubgroupImageBlockWriteINTEL }
func (op *OpSubgroupImageBlockWriteINTEL) Optional() bool { return false }
func (op *OpSubgroupImageBlockWriteINTEL) Verify() error  { return nil }

type OpSubgroupImageMediaBlockReadINTEL struct {
	IdResultType Id
	IdResult     Id
	Image        Id
	Coordinate   Id
	Width        Id
	Height       Id
}

func (op *OpSubgroupImageMediaBlockReadINTEL) Opcode() uint32 {
	return opcodeSubgroupImageMediaBlockReadINTEL
}
func (op *OpSubgroupImageMediaBlockReadINTEL) Optional() bool { return false }
func (op *OpSubgroupImageMediaBlockReadINTEL) Verify() error  { return nil }

type OpSubgroupImageMediaBlockWriteINTEL struct {
	Image      Id
	Coordinate Id
	Width      Id
	Height     Id
	Data       Id
}

func (op *OpSubgroupImageMediaBlockWriteINTEL) Opcode() uint32 {
	return opcodeSubgroupImageMediaBlockWriteINTEL
}
func (op *OpSubgroupImageMediaBlockWriteINTEL) Optional() bool { return false }
func (op *OpSubgroupImageMediaBlockWriteINTEL) Verify() error  { return nil }

type OpDecorateString struct {
	Target     Id
	Decoration Decoration
}

func (op *OpDecorateString) Opcode() uint32 { return opcodeDecorateString }
func (op *OpDecorateString) Optional() bool { return false }
func (op *OpDecorateString) Verify() error  { return nil }

type OpMemberDecorateString struct {
	StructType Id
	Member     uint32
	Decoration Decoration
}

func (op *OpMemberDecorateString) Opcode() uint32 { return opcodeMemberDecorateString }
func (op *OpMemberDecorateString) Optional() bool { return false }
func (op *OpMemberDecorateString) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpNop{} })
	bind(func() Instruction { return &OpUndef{} })
	bind(func() Instruction { return &OpSourceContinued{} })
	bind(func() Instruction { return &OpSource{} })
	bind(func() Instruction { return &OpSourceExtension{} })
	bind(func() Instruction { return &OpName{} })
	bind(func() Instruction { return &OpMemberName{} })
	bind(func() Instruction { return &OpString{} })
	bind(func() Instruction { return &OpLine{} })
	bind(func() Instruction { return &OpExtension{} })
	bind(func() Instruction { return &OpExtInstImport{} })
	bind(func() Instruction { return &OpExtInst{} })
	bind(func() Instruction { return &OpMemoryModel{} })
	bind(func() Instruction { return &OpEntryPoint{} })
	bind(func() Instruction { return &OpExecutionMode{} })
	bind(func() Instruction { return &OpCapability{} })
	bind(func() Instruction { return &OpTypeVoid{} })
	bind(func() Instruction { return &OpTypeBool{} })
	bind(func() Instruction { return &OpTypeInt{} })
	bind(func() Instruction { return &OpTypeFloat{} })
	bind(func() Instruction { return &OpTypeVector{} })
	bind(func() Instruction { return &OpTypeMatrix{} })
	bind(func() Instruction { return &OpTypeImage{} })
	bind(func() Instruction { return &OpTypeSampler{} })
	bind(func() Instruction { return &OpTypeSampledImage{} })
	bind(func() Instruction { return &OpTypeArray{} })
	bind(func() Instruction { return &OpTypeRuntimeArray{} })
	bind(func() Instruction { return &OpTypeStruct{} })
	bind(func() Instruction { return &OpTypeOpaque{} })
	bind(func() Instruction { return &OpTypePointer{} })
	bind(func() Instruction { return &OpTypeFunction{} })
	bind(func() Instruction { return &OpTypeEvent{} })
	bind(func() Instruction { return &OpTypeDeviceEvent{} })
	bind(func() Instruction { return &OpTypeReserveId{} })
	bind(func() Instruction { return &OpTypeQueue{} })
	bind(func() Instruction { return &OpTypePipe{} })
	bind(func() Instruction { return &OpTypeForwardPointer{} })
	bind(func() Instruction { return &OpConstantTrue{} })
	bind(func() Instruction { return &OpConstantFalse{} })
	bind(func() Instruction { return &OpConstant{} })
	bind(func() Instruction { return &OpConstantComposite{} })
	bind(func() Instruction { return &OpConstantSampler{} })
	bind(func() Instruction { return &OpConstantNull{} })
	bind(func() Instruction { return &OpSpecConstantTrue{} })
	bind(func() Instruction { return &OpSpecConstantFalse{} })
	bind(func() Instruction { return &OpSpecConstant{} })
	bind(func() Instruction { return &OpSpecConstantComposite{} })
	bind(func() Instruction { return &OpSpecConstantOp{} })
	bind(func() Instruction { return &OpFunction{} })
	bind(func() Instruction { return &OpFunctionParameter{} })
	bind(func() Instruction { return &OpFunctionEnd{} })
	bind(func() Instruction { return &OpFunctionCall{} })
	bind(func() Instruction { return &OpVariable{} })
	bind(func() Instruction { return &OpImageTexelPointer{} })
	bind(func() Instruction { return &OpLoad{} })
	bind(func() Instruction { return &OpStore{} })
	bind(func() Instruction { return &OpCopyMemory{} })
	bind(func() Instruction { return &OpCopyMemorySized{} })
	bind(func() Instruction { return &OpAccessChain{} })
	bind(func() Instruction { return &OpInBoundsAccessChain{} })
	bind(func() Instruction { return &OpPtrAccessChain{} })
	bind(func() Instruction { return &OpArrayLength{} })
	bind(func() Instruction { return &OpGenericPtrMemSemantics{} })
	bind(func() Instruction { return &OpInBoundsPtrAccessChain{} })
	bind(func() Instruction { return &OpDecorate{} })
	bind(func() Instruction { return &OpMemberDecorate{} })
	bind(func() Instruction { return &OpDecorationGroup{} })
	bind(func() Instruction { return &OpGroupDecorate{} })
	bind(func() Instruction { return &OpGroupMemberDecorate{} })
	bind(func() Instruction { return &OpVectorExtractDynamic{} })
	bind(func() Instruction { return &OpVectorInsertDynamic{} })
	bind(func() Instruction { return &OpVectorShuffle{} })
	bind(func() Instruction { return &OpCompositeConstruct{} })
	bind(func() Instruction { return &OpCompositeExtract{} })
	bind(func() Instruction { return &OpCompositeInsert{} })
	bind(func() Instruction { return &OpCopyObject{} })
	bind(func() Instruction { return &OpTranspose{} })
	bind(func() Instruction { return &OpSampledImage{} })
	bind(func() Instruction { return &OpImageSampleImplicitLod{} })
	bind(func() Instruction { return &OpImageSampleExplicitLod{} })
	bind(func() Instruction { return &OpImageSampleDrefImplicitLod{} })
	bind(func() Instruction { return &OpImageSampleDrefExplicitLod{} })
	bind(func() Instruction { return &OpImageSampleProjImplicitLod{} })
	bind(func() Instruction { return &OpImageSampleProjExplicitLod{} })
	bind(func() Instruction { return &OpImageSampleProjDrefImplicitLod{} })
	bind(func() Instruction { return &OpImageSampleProjDrefExplicitLod{} })
	bind(func() Instruction { return &OpImageFetch{} })
	bind(func() Instruction { return &OpImageGather{} })
	bind(func() Instruction { return &OpImageDrefGather{} })
	bind(func() Instruction { return &OpImageRead{} })
	bind(func() Instruction { return &OpImageWrite{} })
	bind(func() Instruction { return &OpImage{} })
	bind(func() Instruction { return &OpImageQueryFormat{} })
	bind(func() Instruction { return &OpImageQueryOrder{} })
	bind(func() Instruction { return &OpImageQuerySizeLod{} })
	bind(func() Instruction { return &OpImageQuerySize{} })
	bind(func() Instruction { return &OpImageQueryLod{} })
	bind(func() Instruction { return &OpImageQueryLevels{} })
	bind(func() Instruction { return &OpImageQuerySamples{} })
	bind(func() Instruction { return &OpConvertFToU{} })
	bind(func() Instruction { return &OpConvertFToS{} })
	bind(func() Instruction { return &OpConvertSToF{} })
	bind(func() Instruction { return &OpConvertUToF{} })
	bind(func() Instruction { return &OpUConvert{} })
	bind(func() Instruction { return &OpSConvert{} })
	bind(func() Instruction { return &OpFConvert{} })
	bind(func() Instruction { return &OpQuantizeToF16{} })
	bind(func() Instruction { return &OpConvertPtrToU{} })
	bind(func() Instruction { return &OpSatConvertSToU{} })
	bind(func() Instruction { return &OpSatConvertUToS{} })
	bind(func() Instruction { return &OpConvertUToPtr{} })
	bind(func() Instruction { return &OpPtrCastToGeneric{} })
	bind(func() Instruction { return &OpGenericCastToPtr{} })
	bind(func() Instruction { return &OpGenericCastToPtrExplicit{} })
	bind(func() Instruction { return &OpBitcast{} })
	bind(func() Instruction { return &OpSNegate{} })
	bind(func() Instruction { return &OpFNegate{} })
	bind(func() Instruction { return &OpIAdd{} })
	bind(func() Instruction { return &OpFAdd{} })
	bind(func() Instruction { return &OpISub{} })
	bind(func() Instruction { return &OpFSub{} })
	bind(func() Instruction { return &OpIMul{} })
	bind(func() Instruction { return &OpFMul{} })
	bind(func() Instruction { return &OpUDiv{} })
	bind(func() Instruction { return &OpSDiv{} })
	bind(func() Instruction { return &OpFDiv{} })
	bind(func() Instruction { return &OpUMod{} })
	bind(func() Instruction { return &OpSRem{} })
	bind(func() Instruction { return &OpSMod{} })
	bind(func() Instruction { return &OpFRem{} })
	bind(func() Instruction { return &OpFMod{} })
	bind(func() Instruction { return &OpVectorTimesScalar{} })
	bind(func() Instruction { return &OpMatrixTimesScalar{} })
	bind(func() Instruction { return &OpVectorTimesMatrix{} })
	bind(func() Instruction { return &OpMatrixTimesVector{} })
	bind(func() Instruction { return &OpMatrixTimesMatrix{} })
	bind(func() Instruction { return &OpOuterProduct{} })
	bind(func() Instruction { return &OpDot{} })
	bind(func() Instruction { return &OpIAddCarry{} })
	bind(func() Instruction { return &OpISubBorrow{} })
	bind(func() Instruction { return &OpUMulExtended{} })
	bind(func() Instruction { return &OpSMulExtended{} })
	bind(func() Instruction { return &OpAny{} })
	bind(func() Instruction { return &OpAll{} })
	bind(func() Instruction { return &OpIsNan{} })
	bind(func() Instruction { return &OpIsInf{} })
	bind(func() Instruction { return &OpIsFinite{} })
	bind(func() Instruction { return &OpIsNormal{} })
	bind(func() Instruction { return &OpSignBitSet{} })
	bind(func() Instruction { return &OpLessOrGreater{} })
	bind(func() Instruction { return &OpOrdered{} })
	bind(func() Instruction { return &OpUnordered{} })
	bind(func() Instruction { return &OpLogicalEqual{} })
	bind(func() Instruction { return &OpLogicalNotEqual{} })
	bind(func() Instruction { return &OpLogicalOr{} })
	bind(func() Instruction { return &OpLogicalAnd{} })
	bind(func() Instruction { return &OpLogicalNot{} })
	bind(func() Instruction { return &OpSelect{} })
	bind(func() Instruction { return &OpIEqual{} })
	bind(func() Instruction { return &OpINotEqual{} })
	bind(func() Instruction { return &OpUGreaterThan{} })
	bind(func() Instruction { return &OpSGreaterThan{} })
	bind(func() Instruction { return &OpUGreaterThanEqual{} })
	bind(func() Instruction { return &OpSGreaterThanEqual{} })
	bind(func() Instruction { return &OpULessThan{} })
	bind(func() Instruction { return &OpSLessThan{} })
	bind(func() Instruction { return &OpULessThanEqual{} })
	bind(func() Instruction { return &OpSLessThanEqual{} })
	bind(func() Instruction { return &OpFOrdEqual{} })
	bind(func() Instruction { return &OpFUnordEqual{} })
	bind(func() Instruction { return &OpFOrdNotEqual{} })
	bind(func() Instruction { return &OpFUnordNotEqual{} })
	bind(func() Instruction { return &OpFOrdLessThan{} })
	bind(func() Instruction { return &OpFUnordLessThan{} })
	bind(func() Instruction { return &OpFOrdGreaterThan{} })
	bind(func() Instruction { return &OpFUnordGreaterThan{} })
	bind(func() Instruction { return &OpFOrdLessThanEqual{} })
	bind(func() Instruction { return &OpFUnordLessThanEqual{} })
	bind(func() Instruction { return &OpFOrdGreaterThanEqual{} })
	bind(func() Instruction { return &OpFUnordGreaterThanEqual{} })
	bind(func() Instruction { return &OpShiftRightLogical{} })
	bind(func() Instruction { return &OpShiftRightArithmetic{} })
	bind(func() Instruction { return &OpShiftLeftLogical{} })
	bind(func() Instruction { return &OpBitwiseOr{} })
	bind(func() Instruction { return &OpBitwiseXor{} })
	bind(func() Instruction { return &OpBitwiseAnd{} })
	bind(func() Instruction { return &OpNot{} })
	bind(func() Instruction { return &OpBitFieldInsert{} })
	bind(func() Instruction { return &OpBitFieldSExtract{} })
	bind(func() Instruction { return &OpBitFieldUExtract{} })
	bind(func() Instruction { return &OpBitReverse{} })
	bind(func() Instruction { return &OpBitCount{} })
	bind(func() Instruction { return &OpDPdx{} })
	bind(func() Instruction { return &OpDPdy{} })
	bind(func() Instruction { return &OpFwidth{} })
	bind(func() Instruction { return &OpDPdxFine{} })
	bind(func() Instruction { return &OpDPdyFine{} })
	bind(func() Instruction { return &OpFwidthFine{} })
	bind(func() Instruction { return &OpDPdxCoarse{} })
	bind(func() Instruction { return &OpDPdyCoarse{} })
	bind(func() Instruction { return &OpFwidthCoarse{} })
	bind(func() Instruction { return &OpEmitVertex{} })
	bind(func() Instruction { return &OpEndPrimitive{} })
	bind(func() Instruction { return &OpEmitStreamVertex{} })
	bind(func() Instruction { return &OpEndStreamPrimitive{} })
	bind(func() Instruction { return &OpControlBarrier{} })
	bind(func() Instruction { return &OpMemoryBarrier{} })
	bind(func() Instruction { return &OpAtomicLoad{} })
	bind(func() Instruction { return &OpAtomicStore{} })
	bind(func() Instruction { return &OpAtomicExchange{} })
	bind(func() Instruction { return &OpAtomicCompareExchange{} })
	bind(func() Instruction { return &OpAtomicCompareExchangeWeak{} })
	bind(func() Instruction { return &OpAtomicIIncrement{} })
	bind(func() Instruction { return &OpAtomicIDecrement{} })
	bind(func() Instruction { return &OpAtomicIAdd{} })
	bind(func() Instruction { return &OpAtomicISub{} })
	bind(func() Instruction { return &OpAtomicSMin{} })
	bind(func() Instruction { return &OpAtomicUMin{} })
	bind(func() Instruction { return &OpAtomicSMax{} })
	bind(func() Instruction { return &OpAtomicUMax{} })
	bind(func() Instruction { return &OpAtomicAnd{} })
	bind(func() Instruction { return &OpAtomicOr{} })
	bind(func() Instruction { return &OpAtomicXor{} })
	bind(func() Instruction { return &OpPhi{} })
	bind(func() Instruction { return &OpLoopMerge{} })
	bind(func() Instruction { return &OpSelectionMerge{} })
	bind(func() Instruction { return &OpLabel{} })
	bind(func() Instruction { return &OpBranch{} })
	bind(func() Instruction { return &OpBranchConditional{} })
	bind(func() Instruction { return &OpSwitch{} })
	bind(func() Instruction { return &OpKill{} })
	bind(func() Instruction { return &OpReturn{} })
	bind(func() Instruction { return &OpReturnValue{} })
	bind(func() Instruction { return &OpUnreachable{} })
	bind(func() Instruction { return &OpLifetimeStart{} })
	bind(func() Instruction { return &OpLifetimeStop{} })
	bind(func() Instruction { return &OpGroupAsyncCopy{} })
	bind(func() Instruction { return &OpGroupWaitEvents{} })
	bind(func() Instruction { return &OpGroupAll{} })
	bind(func() Instruction { return &OpGroupAny{} })
	bind(func() Instruction { return &OpGroupBroadcast{} })
	bind(func() Instruction { return &OpGroupIAdd{} })
	bind(func() Instruction { return &OpGroupFAdd{} })
	bind(func() Instruction { return &OpGroupFMin{} })
	bind(func() Instruction { return &OpGroupUMin{} })
	bind(func() Instruction { return &OpGroupSMin{} })
	bind(func() Instruction { return &OpGroupFMax{} })
	bind(func() Instruction { return &OpGroupUMax{} })
	bind(func() Instruction { return &OpGroupSMax{} })
	bind(func() Instruction { return &OpReadPipe{} })
	bind(func() Instruction { return &OpWritePipe{} })
	bind(func() Instruction { return &OpReservedReadPipe{} })
	bind(func() Instruction { return &OpReservedWritePipe{} })
	bind(func() Instruction { return &OpReserveReadPipePackets{} })
	bind(func() Instruction { return &OpReserveWritePipePackets{} })
	bind(func() Instruction { return &OpCommitReadPipe{} })
	bind(func() Instruction { return &OpCommitWritePipe{} })
	bind(func() Instruction { return &OpIsValidReserveId{} })
	bind(func() Instruction { return &OpGetNumPipePackets{} })
	bind(func() Instruction { return &OpGetMaxPipePackets{} })
	bind(func() Instruction { return &OpGroupReserveReadPipePackets{} })
	bind(func() Instruction { return &OpGroupReserveWritePipePackets{} })
	bind(func() Instruction { return &OpGroupCommitReadPipe{} })
	bind(func() Instruction { return &OpGroupCommitWritePipe{} })
	bind(func() Instruction { return &OpEnqueueMarker{} })
	bind(func() Instruction { return &OpEnqueueKernel{} })
	bind(func() Instruction { return &OpGetKernelNDrangeSubGroupCount{} })
	bind(func() Instruction { return &OpGetKernelNDrangeMaxSubGroupSize{} })
	bind(func() Instruction { return &OpGetKernelWorkGroupSize{} })
	bind(func() Instruction { return &OpGetKernelPreferredWorkGroupSizeMultiple{} })
	bind(func() Instruction { return &OpRetainEvent{} })
	bind(func() Instruction { return &OpReleaseEvent{} })
	bind(func() Instruction { return &OpCreateUserEvent{} })
	bind(func() Instruction { return &OpIsValidEvent{} })
	bind(func() Instruction { return &OpSetUserEventStatus{} })
	bind(func() Instruction { return &OpCaptureEventProfilingInfo{} })
	bind(func() Instruction { return &OpGetDefaultQueue{} })
	bind(func() Instruction { return &OpBuildNDRange{} })
	bind(func() Instruction { return &OpImageSparseSampleImplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleExplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleDrefImplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleDrefExplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleProjImplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleProjExplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleProjDrefImplicitLod{} })
	bind(func() Instruction { return &OpImageSparseSampleProjDrefExplicitLod{} })
	bind(func() Instruction { return &OpImageSparseFetch{} })
	bind(func() Instruction { return &OpImageSparseGather{} })
	bind(func() Instruction { return &OpImageSparseDrefGather{} })
	bind(func() Instruction { return &OpImageSparseTexelsResident{} })
	bind(func() Instruction { return &OpNoLine{} })
	bind(func() Instruction { return &OpAtomicFlagTestAndSet{} })
	bind(func() Instruction { return &OpAtomicFlagClear{} })
	bind(func() Instruction { return &OpImageSparseRead{} })
	bind(func() Instruction { return &OpSizeOf{} })
	bind(func() Instruction { return &OpTypePipeStorage{} })
	bind(func() Instruction { return &OpConstantPipeStorage{} })
	bind(func() Instruction { return &OpCreatePipeFromPipeStorage{} })
	bind(func() Instruction { return &OpGetKernelLocalSizeForSubgroupCount{} })
	bind(func() Instruction { return &OpGetKernelMaxNumSubgroups{} })
	bind(func() Instruction { return &OpTypeNamedBarrier{} })
	bind(func() Instruction { return &OpNamedBarrierInitialize{} })
	bind(func() Instruction { return &OpMemoryNamedBarrier{} })
	bind(func() Instruction { return &OpModuleProcessed{} })
	bind(func() Instruction { return &OpExecutionModeId{} })
	bind(func() Instruction { return &OpDecorateId{} })
	bind(func() Instruction { return &OpGroupNonUniformElect{} })
	bind(func() Instruction { return &OpGroupNonUniformAll{} })
	bind(func() Instruction { return &OpGroupNonUniformAny{} })
	bind(func() Instruction { return &OpGroupNonUniformAllEqual{} })
	bind(func() Instruction { return &OpGroupNonUniformBroadcast{} })
	bind(func() Instruction { return &OpGroupNonUniformBroadcastFirst{} })
	bind(func() Instruction { return &OpGroupNonUniformBallot{} })
	bind(func() Instruction { return &OpGroupNonUniformInverseBallot{} })
	bind(func() Instruction { return &OpGroupNonUniformBallotBitExtract{} })
	bind(func() Instruction { return &OpGroupNonUniformBallotBitCount{} })
	bind(func() Instruction { return &OpGroupNonUniformBallotFindLSB{} })
	bind(func() Instruction { return &OpGroupNonUniformBallotFindMSB{} })
	bind(func() Instruction { return &OpGroupNonUniformShuffle{} })
	bind(func() Instruction { return &OpGroupNonUniformShuffleXor{} })
	bind(func() Instruction { return &OpGroupNonUniformShuffleUp{} })
	bind(func() Instruction { return &OpGroupNonUniformShuffleDown{} })
	bind(func() Instruction { return &OpGroupNonUniformIAdd{} })
	bind(func() Instruction { return &OpGroupNonUniformFAdd{} })
	bind(func() Instruction { return &OpGroupNonUniformIMul{} })
	bind(func() Instruction { return &OpGroupNonUniformFMul{} })
	bind(func() Instruction { return &OpGroupNonUniformSMin{} })
	bind(func() Instruction { return &OpGroupNonUniformUMin{} })
	bind(func() Instruction { return &OpGroupNonUniformFMin{} })
	bind(func() Instruction { return &OpGroupNonUniformSMax{} })
	bind(func() Instruction { return &OpGroupNonUniformUMax{} })
	bind(func() Instruction { return &OpGroupNonUniformFMax{} })
	bind(func() Instruction { return &OpGroupNonUniformBitwiseAnd{} })
	bind(func() Instruction { return &OpGroupNonUniformBitwiseOr{} })
	bind(func() Instruction { return &OpGroupNonUniformBitwiseXor{} })
	bind(func() Instruction { return &OpGroupNonUniformLogicalAnd{} })
	bind(func() Instruction { return &OpGroupNonUniformLogicalOr{} })
	bind(func() Instruction { return &OpGroupNonUniformLogicalXor{} })
	bind(func() Instruction { return &OpGroupNonUniformQuadBroadcast{} })
	bind(func() Instruction { return &OpGroupNonUniformQuadSwap{} })
	bind(func() Instruction { return &OpCopyLogical{} })
	bind(func() Instruction { return &OpPtrEqual{} })
	bind(func() Instruction { return &OpPtrNotEqual{} })
	bind(func() Instruction { return &OpPtrDiff{} })
	bind(func() Instruction { return &OpSubgroupBallotKHR{} })
	bind(func() Instruction { return &OpSubgroupFirstInvocationKHR{} })
	bind(func() Instruction { return &OpSubgroupAllKHR{} })
	bind(func() Instruction { return &OpSubgroupAnyKHR{} })
	bind(func() Instruction { return &OpSubgroupAllEqualKHR{} })
	bind(func() Instruction { return &OpSubgroupReadInvocationKHR{} })
	bind(func() Instruction { return &OpGroupIAddNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupFAddNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupFMinNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupUMinNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupSMinNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupFMaxNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupUMaxNonUniformAMD{} })
	bind(func() Instruction { return &OpGroupSMaxNonUniformAMD{} })
	bind(func() Instruction { return &OpImageSampleFootprintNV{} })
	bind(func() Instruction { return &OpGroupNonUniformPartitionNV{} })
	bind(func() Instruction { return &OpSubgroupShuffleINTEL{} })
	bind(func() Instruction { return &OpSubgroupShuffleDownINTEL{} })
	bind(func() Instruction { return &OpSubgroupShuffleUpINTEL{} })
	bind(func() Instruction { return &OpSubgroupShuffleXorINTEL{} })
	bind(func() Instruction { return &OpSubgroupBlockReadINTEL{} })
	bind(func() Instruction { return &OpSubgroupBlockWriteINTEL{} })
	bind(func() Instruction { return &OpSubgroupImageBlockReadINTEL{} })
	bind(func() Instruction { return &OpSubgroupImageBlockWriteINTEL{} })
	bind(func() Instruction { return &OpSubgroupImageMediaBlockReadINTEL{} })
	bind(func() Instruction { return &OpSubgroupImageMediaBlockWriteINTEL{} })
	bind(func() Instruction { return &OpDecorateString{} })
	bind(func() Instruction { return &OpMemberDecorateString{} })
}
