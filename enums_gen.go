// This file was generated. DO NOT EDIT!

package spirv

type ImageOperands uint32

func (v ImageOperands) Verify() error { return nil }

const (
	ImageOperandsNone                  ImageOperands = 0x0000
	ImageOperandsBias                  ImageOperands = 0x0001
	ImageOperandsLod                   ImageOperands = 0x0002
	ImageOperandsGrad                  ImageOperands = 0x0004
	ImageOperandsConstOffset           ImageOperands = 0x0008
	ImageOperandsOffset                ImageOperands = 0x0010
	ImageOperandsConstOffsets          ImageOperands = 0x0020
	ImageOperandsSample                ImageOperands = 0x0040
	ImageOperandsMinLod                ImageOperands = 0x0080
	ImageOperandsMakeTexelAvailable    ImageOperands = 0x0100
	ImageOperandsMakeTexelAvailableKHR ImageOperands = 0x0100
	ImageOperandsMakeTexelVisible      ImageOperands = 0x0200
	ImageOperandsMakeTexelVisibleKHR   ImageOperands = 0x0200
	ImageOperandsNonPrivateTexel       ImageOperands = 0x0400
	ImageOperandsNonPrivateTexelKHR    ImageOperands = 0x0400
	ImageOperandsVolatileTexel         ImageOperands = 0x0800
	ImageOperandsVolatileTexelKHR      ImageOperands = 0x0800
	ImageOperandsSignExtend            ImageOperands = 0x1000
	ImageOperandsZeroExtend            ImageOperands = 0x2000
)

type FPFastMathMode uint32

func (v FPFastMathMode) Verify() error { return nil }

const (
	FPFastMathModeNone       FPFastMathMode = 0x0000
	FPFastMathModeNotNaN     FPFastMathMode = 0x0001
	FPFastMathModeNotInf     FPFastMathMode = 0x0002
	FPFastMathModeNSZ        FPFastMathMode = 0x0004
	FPFastMathModeAllowRecip FPFastMathMode = 0x0008
	FPFastMathModeFast       FPFastMathMode = 0x0010
)

type SelectionControl uint32

func (v SelectionControl) Verify() error { return nil }

const (
	SelectionControlNone        SelectionControl = 0x0000
	SelectionControlFlatten     SelectionControl = 0x0001
	SelectionControlDontFlatten SelectionControl = 0x0002
)

type LoopControl uint32

func (v LoopControl) Verify() error { return nil }

const (
	LoopControlNone               LoopControl = 0x0000
	LoopControlUnroll             LoopControl = 0x0001
	LoopControlDontUnroll         LoopControl = 0x0002
	LoopControlDependencyInfinite LoopControl = 0x0004
	LoopControlDependencyLength   LoopControl = 0x0008
	LoopControlMinIterations      LoopControl = 0x0010
	LoopControlMaxIterations      LoopControl = 0x0020
	LoopControlIterationMultiple  LoopControl = 0x0040
	LoopControlPeelCount          LoopControl = 0x0080
	LoopControlPartialCount       LoopControl = 0x0100
)

type FunctionControl uint32

func (v FunctionControl) Verify() error { return nil }

const (
	FunctionControlNone       FunctionControl = 0x0000
	FunctionControlInline     FunctionControl = 0x0001
	FunctionControlDontInline FunctionControl = 0x0002
	FunctionControlPure       FunctionControl = 0x0004
	FunctionControlConst      FunctionControl = 0x0008
)

type MemorySemantics uint32

func (v MemorySemantics) Verify() error { return nil }

const (
	MemorySemanticsRelaxed                MemorySemantics = 0x0000
	MemorySemanticsNone                   MemorySemantics = 0x0000
	MemorySemanticsAcquire                MemorySemantics = 0x0002
	MemorySemanticsRelease                MemorySemantics = 0x0004
	MemorySemanticsAcquireRelease         MemorySemantics = 0x0008
	MemorySemanticsSequentiallyConsistent MemorySemantics = 0x0010
	MemorySemanticsUniformMemory          MemorySemantics = 0x0040
	MemorySemanticsSubgroupMemory         MemorySemantics = 0x0080
	MemorySemanticsWorkgroupMemory        MemorySemantics = 0x0100
	MemorySemanticsCrossWorkgroupMemory   MemorySemantics = 0x0200
	MemorySemanticsAtomicCounterMemory    MemorySemantics = 0x0400
	MemorySemanticsImageMemory            MemorySemantics = 0x0800
	MemorySemanticsOutputMemory           MemorySemantics = 0x1000
	MemorySemanticsOutputMemoryKHR        MemorySemantics = 0x1000
	MemorySemanticsMakeAvailable          MemorySemantics = 0x2000
	MemorySemanticsMakeAvailableKHR       MemorySemantics = 0x2000
	MemorySemanticsMakeVisible            MemorySemantics = 0x4000
	MemorySemanticsMakeVisibleKHR         MemorySemantics = 0x4000
	MemorySemanticsVolatile               MemorySemantics = 0x8000
)

type MemoryAccess uint32

func (v MemoryAccess) Verify() error { return nil }

const (
	MemoryAccessNone                    MemoryAccess = 0x0000
	MemoryAccessVolatile                MemoryAccess = 0x0001
	MemoryAccessAligned                 MemoryAccess = 0x0002
	MemoryAccessNontemporal             MemoryAccess = 0x0004
	MemoryAccessMakePointerAvailable    MemoryAccess = 0x0008
	MemoryAccessMakePointerAvailableKHR MemoryAccess = 0x0008
	MemoryAccessMakePointerVisible      MemoryAccess = 0x0010
	MemoryAccessMakePointerVisibleKHR   MemoryAccess = 0x0010
	MemoryAccessNonPrivatePointer       MemoryAccess = 0x0020
	MemoryAccessNonPrivatePointerKHR    MemoryAccess = 0x0020
)

type KernelProfilingInfo uint32

func (v KernelProfilingInfo) Verify() error { return nil }

const (
	KernelProfilingInfoNone        KernelProfilingInfo = 0x0000
	KernelProfilingInfoCmdExecTime KernelProfilingInfo = 0x0001
)

type SourceLanguage uint32

func (v SourceLanguage) Verify() error { return nil }

const (
	SourceLanguageUnknown    SourceLanguage = 0
	SourceLanguageESSL       SourceLanguage = 1
	SourceLanguageGLSL       SourceLanguage = 2
	SourceLanguageOpenCL_C   SourceLanguage = 3
	SourceLanguageOpenCL_CPP SourceLanguage = 4
	SourceLanguageHLSL       SourceLanguage = 5
)

type ExecutionModel uint32

func (v ExecutionModel) Verify() error { return nil }

const (
	ExecutionModelVertex                 ExecutionModel = 0
	ExecutionModelTessellationControl    ExecutionModel = 1
	ExecutionModelTessellationEvaluation ExecutionModel = 2
	ExecutionModelGeometry               ExecutionModel = 3
	ExecutionModelFragment               ExecutionModel = 4
	ExecutionModelGLCompute              ExecutionModel = 5
	ExecutionModelKernel                 ExecutionModel = 6
	ExecutionModelTaskNV                 ExecutionModel = 5267
	ExecutionModelMeshNV                 ExecutionModel = 5268
	ExecutionModelRayGenerationNV        ExecutionModel = 5313
	ExecutionModelIntersectionNV         ExecutionModel = 5314
	ExecutionModelAnyHitNV               ExecutionModel = 5315
	ExecutionModelClosestHitNV           ExecutionModel = 5316
	ExecutionModelMissNV                 ExecutionModel = 5317
	ExecutionModelCallableNV             ExecutionModel = 5318
)

type AddressingModel uint32

func (v AddressingModel) Verify() error { return nil }

const (
	AddressingModelLogical                    AddressingModel = 0
	AddressingModelPhysical32                 AddressingModel = 1
	AddressingModelPhysical64                 AddressingModel = 2
	AddressingModelPhysicalStorageBuffer64    AddressingModel = 5348
	AddressingModelPhysicalStorageBuffer64EXT AddressingModel = 5348
)

type MemoryModel uint32

func (v MemoryModel) Verify() error { return nil }

const (
	MemoryModelSimple    MemoryModel = 0
	MemoryModelGLSL450   MemoryModel = 1
	MemoryModelOpenCL    MemoryModel = 2
	MemoryModelVulkan    MemoryModel = 3
	MemoryModelVulkanKHR MemoryModel = 3
)

type ExecutionMode uint32

func (v ExecutionMode) Verify() error { return nil }

const (
	ExecutionModeInvocations                      ExecutionMode = 0
	ExecutionModeSpacingEqual                     ExecutionMode = 1
	ExecutionModeSpacingFractionalEven            ExecutionMode = 2
	ExecutionModeSpacingFractionalOdd             ExecutionMode = 3
	ExecutionModeVertexOrderCw                    ExecutionMode = 4
	ExecutionModeVertexOrderCcw                   ExecutionMode = 5
	ExecutionModePixelCenterInteger               ExecutionMode = 6
	ExecutionModeOriginUpperLeft                  ExecutionMode = 7
	ExecutionModeOriginLowerLeft                  ExecutionMode = 8
	ExecutionModeEarlyFragmentTests               ExecutionMode = 9
	ExecutionModePointMode                        ExecutionMode = 10
	ExecutionModeXfb                              ExecutionMode = 11
	ExecutionModeDepthReplacing                   ExecutionMode = 12
	ExecutionModeDepthGreater                     ExecutionMode = 14
	ExecutionModeDepthLess                        ExecutionMode = 15
	ExecutionModeDepthUnchanged                   ExecutionMode = 16
	ExecutionModeLocalSize                        ExecutionMode = 17
	ExecutionModeLocalSizeHint                    ExecutionMode = 18
	ExecutionModeInputPoints                      ExecutionMode = 19
	ExecutionModeInputLines                       ExecutionMode = 20
	ExecutionModeInputLinesAdjacency              ExecutionMode = 21
	ExecutionModeTriangles                        ExecutionMode = 22
	ExecutionModeInputTrianglesAdjacency          ExecutionMode = 23
	ExecutionModeQuads                            ExecutionMode = 24
	ExecutionModeIsolines                         ExecutionMode = 25
	ExecutionModeOutputVertices                   ExecutionMode = 26
	ExecutionModeOutputPoints                     ExecutionMode = 27
	ExecutionModeOutputLineStrip                  ExecutionMode = 28
	ExecutionModeOutputTriangleStrip              ExecutionMode = 29
	ExecutionModeVecTypeHint                      ExecutionMode = 30
	ExecutionModeContractionOff                   ExecutionMode = 31
	ExecutionModeInitializer                      ExecutionMode = 33
	ExecutionModeFinalizer                        ExecutionMode = 34
	ExecutionModeSubgroupSize                     ExecutionMode = 35
	ExecutionModeSubgroupsPerWorkgroup            ExecutionMode = 36
	ExecutionModeSubgroupsPerWorkgroupId          ExecutionMode = 37
	ExecutionModeLocalSizeId                      ExecutionMode = 38
	ExecutionModeLocalSizeHintId                  ExecutionMode = 39
	ExecutionModePostDepthCoverage                ExecutionMode = 4446
	ExecutionModeDenormPreserve                   ExecutionMode = 4459
	ExecutionModeDenormFlushToZero                ExecutionMode = 4460
	ExecutionModeSignedZeroInfNanPreserve         ExecutionMode = 4461
	ExecutionModeRoundingModeRTE                  ExecutionMode = 4462
	ExecutionModeRoundingModeRTZ                  ExecutionMode = 4463
	ExecutionModeStencilRefReplacingEXT           ExecutionMode = 5027
	ExecutionModeOutputLinesNV                    ExecutionMode = 5269
	ExecutionModeOutputPrimitivesNV               ExecutionMode = 5270
	ExecutionModeDerivativeGroupQuadsNV           ExecutionMode = 5289
	ExecutionModeDerivativeGroupLinearNV          ExecutionMode = 5290
	ExecutionModeOutputTrianglesNV                ExecutionMode = 5298
	ExecutionModePixelInterlockOrderedEXT         ExecutionMode = 5366
	ExecutionModePixelInterlockUnorderedEXT       ExecutionMode = 5367
	ExecutionModeSampleInterlockOrderedEXT        ExecutionMode = 5368
	ExecutionModeSampleInterlockUnorderedEXT      ExecutionMode = 5369
	ExecutionModeShadingRateInterlockOrderedEXT   ExecutionMode = 5370
	ExecutionModeShadingRateInterlockUnorderedEXT ExecutionMode = 5371
)

type StorageClass uint32

func (v StorageClass) Verify() error { return nil }

const (
	StorageClassUniformConstant          StorageClass = 0
	StorageClassInput                    StorageClass = 1
	StorageClassUniform                  StorageClass = 2
	StorageClassOutput                   StorageClass = 3
	StorageClassWorkgroup                StorageClass = 4
	StorageClassCrossWorkgroup           StorageClass = 5
	StorageClassPrivate                  StorageClass = 6
	StorageClassFunction                 StorageClass = 7
	StorageClassGeneric                  StorageClass = 8
	StorageClassPushConstant             StorageClass = 9
	StorageClassAtomicCounter            StorageClass = 10
	StorageClassImage                    StorageClass = 11
	StorageClassStorageBuffer            StorageClass = 12
	StorageClassCallableDataNV           StorageClass = 5328
	StorageClassIncomingCallableDataNV   StorageClass = 5329
	StorageClassRayPayloadNV             StorageClass = 5338
	StorageClassHitAttributeNV           StorageClass = 5339
	StorageClassIncomingRayPayloadNV     StorageClass = 5342
	StorageClassShaderRecordBufferNV     StorageClass = 5343
	StorageClassPhysicalStorageBuffer    StorageClass = 5349
	StorageClassPhysicalStorageBufferEXT StorageClass = 5349
)

type Dim uint32

func (v Dim) Verify() error { return nil }

const (
	Dim1D          Dim = 0
	Dim2D          Dim = 1
	Dim3D          Dim = 2
	DimCube        Dim = 3
	DimRect        Dim = 4
	DimBuffer      Dim = 5
	DimSubpassData Dim = 6
)

type SamplerAddressingMode uint32

func (v SamplerAddressingMode) Verify() error { return nil }

const (
	SamplerAddressingModeNone           SamplerAddressingMode = 0
	SamplerAddressingModeClampToEdge    SamplerAddressingMode = 1
	SamplerAddressingModeClamp          SamplerAddressingMode = 2
	SamplerAddressingModeRepeat         SamplerAddressingMode = 3
	SamplerAddressingModeRepeatMirrored SamplerAddressingMode = 4
)

type SamplerFilterMode uint32

func (v SamplerFilterMode) Verify() error { return nil }

const (
	SamplerFilterModeNearest SamplerFilterMode = 0
	SamplerFilterModeLinear  SamplerFilterMode = 1
)

type ImageFormat uint32

func (v ImageFormat) Verify() error { return nil }

const (
	ImageFormatUnknown      ImageFormat = 0
	ImageFormatRgba32f      ImageFormat = 1
	ImageFormatRgba16f      ImageFormat = 2
	ImageFormatR32f         ImageFormat = 3
	ImageFormatRgba8        ImageFormat = 4
	ImageFormatRgba8Snorm   ImageFormat = 5
	ImageFormatRg32f        ImageFormat = 6
	ImageFormatRg16f        ImageFormat = 7
	ImageFormatR11fG11fB10f ImageFormat = 8
	ImageFormatR16f         ImageFormat = 9
	ImageFormatRgba16       ImageFormat = 10
	ImageFormatRgb10A2      ImageFormat = 11
	ImageFormatRg16         ImageFormat = 12
	ImageFormatRg8          ImageFormat = 13
	ImageFormatR16          ImageFormat = 14
	ImageFormatR8           ImageFormat = 15
	ImageFormatRgba16Snorm  ImageFormat = 16
	ImageFormatRg16Snorm    ImageFormat = 17
	ImageFormatRg8Snorm     ImageFormat = 18
	ImageFormatR16Snorm     ImageFormat = 19
	ImageFormatR8Snorm      ImageFormat = 20
	ImageFormatRgba32i      ImageFormat = 21
	ImageFormatRgba16i      ImageFormat = 22
	ImageFormatRgba8i       ImageFormat = 23
	ImageFormatR32i         ImageFormat = 24
	ImageFormatRg32i        ImageFormat = 25
	ImageFormatRg16i        ImageFormat = 26
	ImageFormatRg8i         ImageFormat = 27
	ImageFormatR16i         ImageFormat = 28
	ImageFormatR8i          ImageFormat = 29
	ImageFormatRgba32ui     ImageFormat = 30
	ImageFormatRgba16ui     ImageFormat = 31
	ImageFormatRgba8ui      ImageFormat = 32
	ImageFormatR32ui        ImageFormat = 33
	ImageFormatRgb10a2ui    ImageFormat = 34
	ImageFormatRg32ui       ImageFormat = 35
	ImageFormatRg16ui       ImageFormat = 36
	ImageFormatRg8ui        ImageFormat = 37
	ImageFormatR16ui        ImageFormat = 38
	ImageFormatR8ui         ImageFormat = 39
)

type ImageChannelOrder uint32

func (v ImageChannelOrder) Verify() error { return nil }

const (
	ImageChannelOrderR            ImageChannelOrder = 0
	ImageChannelOrderA            ImageChannelOrder = 1
	ImageChannelOrderRG           ImageChannelOrder = 2
	ImageChannelOrderRA           ImageChannelOrder = 3
	ImageChannelOrderRGB          ImageChannelOrder = 4
	ImageChannelOrderRGBA         ImageChannelOrder = 5
	ImageChannelOrderBGRA         ImageChannelOrder = 6
	ImageChannelOrderARGB         ImageChannelOrder = 7
	ImageChannelOrderIntensity    ImageChannelOrder = 8
	ImageChannelOrderLuminance    ImageChannelOrder = 9
	ImageChannelOrderRx           ImageChannelOrder = 10
	ImageChannelOrderRGx          ImageChannelOrder = 11
	ImageChannelOrderRGBx         ImageChannelOrder = 12
	ImageChannelOrderDepth        ImageChannelOrder = 13
	ImageChannelOrderDepthStencil ImageChannelOrder = 14
	ImageChannelOrdersRGB         ImageChannelOrder = 15
	ImageChannelOrdersRGBx        ImageChannelOrder = 16
	ImageChannelOrdersRGBA        ImageChannelOrder = 17
	ImageChannelOrdersBGRA        ImageChannelOrder = 18
	ImageChannelOrderABGR         ImageChannelOrder = 19
)

type ImageChannelDataType uint32

func (v ImageChannelDataType) Verify() error { return nil }

const (
	ImageChannelDataTypeSnormInt8        ImageChannelDataType = 0
	ImageChannelDataTypeSnormInt16       ImageChannelDataType = 1
	ImageChannelDataTypeUnormInt8        ImageChannelDataType = 2
	ImageChannelDataTypeUnormInt16       ImageChannelDataType = 3
	ImageChannelDataTypeUnormShort565    ImageChannelDataType = 4
	ImageChannelDataTypeUnormShort555    ImageChannelDataType = 5
	ImageChannelDataTypeUnormInt101010   ImageChannelDataType = 6
	ImageChannelDataTypeSignedInt8       ImageChannelDataType = 7
	ImageChannelDataTypeSignedInt16      ImageChannelDataType = 8
	ImageChannelDataTypeSignedInt32      ImageChannelDataType = 9
	ImageChannelDataTypeUnsignedInt8     ImageChannelDataType = 10
	ImageChannelDataTypeUnsignedInt16    ImageChannelDataType = 11
	ImageChannelDataTypeUnsignedInt32    ImageChannelDataType = 12
	ImageChannelDataTypeHalfFloat        ImageChannelDataType = 13
	ImageChannelDataTypeFloat            ImageChannelDataType = 14
	ImageChannelDataTypeUnormInt24       ImageChannelDataType = 15
	ImageChannelDataTypeUnormInt101010_2 ImageChannelDataType = 16
)

type FPRoundingMode uint32

func (v FPRoundingMode) Verify() error { return nil }

const (
	FPRoundingModeRTE FPRoundingMode = 0
	FPRoundingModeRTZ FPRoundingMode = 1
	FPRoundingModeRTP FPRoundingMode = 2
	FPRoundingModeRTN FPRoundingMode = 3
)

type LinkageType uint32

func (v LinkageType) Verify() error { return nil }

const (
	LinkageTypeExport LinkageType = 0
	LinkageTypeImport LinkageType = 1
)

type AccessQualifier uint32

func (v AccessQualifier) Verify() error { return nil }

const (
	AccessQualifierReadOnly  AccessQualifier = 0
	AccessQualifierWriteOnly AccessQualifier = 1
	AccessQualifierReadWrite AccessQualifier = 2
)

type FunctionParameterAttribute uint32

func (v FunctionParameterAttribute) Verify() error { return nil }

const (
	FunctionParameterAttributeZext        FunctionParameterAttribute = 0
	FunctionParameterAttributeSext        FunctionParameterAttribute = 1
	FunctionParameterAttributeByVal       FunctionParameterAttribute = 2
	FunctionParameterAttributeSret        FunctionParameterAttribute = 3
	FunctionParameterAttributeNoAlias     FunctionParameterAttribute = 4
	FunctionParameterAttributeNoCapture   FunctionParameterAttribute = 5
	FunctionParameterAttributeNoWrite     FunctionParameterAttribute = 6
	FunctionParameterAttributeNoReadWrite FunctionParameterAttribute = 7
)

type Decoration uint32

func (v Decoration) Verify() error { return nil }

const (
	DecorationRelaxedPrecision            Decoration = 0
	DecorationSpecId                      Decoration = 1
	DecorationBlock                       Decoration = 2
	DecorationBufferBlock                 Decoration = 3
	DecorationRowMajor                    Decoration = 4
	DecorationColMajor                    Decoration = 5
	DecorationArrayStride                 Decoration = 6
	DecorationMatrixStride                Decoration = 7
	DecorationGLSLShared                  Decoration = 8
	DecorationGLSLPacked                  Decoration = 9
	DecorationCPacked                     Decoration = 10
	DecorationBuiltIn                     Decoration = 11
	DecorationNoPerspective               Decoration = 13
	DecorationFlat                        Decoration = 14
	DecorationPatch                       Decoration = 15
	DecorationCentroid                    Decoration = 16
	DecorationSample                      Decoration = 17
	DecorationInvariant                   Decoration = 18
	DecorationRestrict                    Decoration = 19
	DecorationAliased                     Decoration = 20
	DecorationVolatile                    Decoration = 21
	DecorationConstant                    Decoration = 22
	DecorationCoherent                    Decoration = 23
	DecorationNonWritable                 Decoration = 24
	DecorationNonReadable                 Decoration = 25
	DecorationUniform                     Decoration = 26
	DecorationUniformId                   Decoration = 27
	DecorationSaturatedConversion         Decoration = 28
	DecorationStream                      Decoration = 29
	DecorationLocation                    Decoration = 30
	DecorationComponent                   Decoration = 31
	DecorationIndex                       Decoration = 32
	DecorationBinding                     Decoration = 33
	DecorationDescriptorSet               Decoration = 34
	DecorationOffset                      Decoration = 35
	DecorationXfbBuffer                   Decoration = 36
	DecorationXfbStride                   Decoration = 37
	DecorationFuncParamAttr               Decoration = 38
	DecorationFPRoundingMode              Decoration = 39
	DecorationFPFastMathMode              Decoration = 40
	DecorationLinkageAttributes           Decoration = 41
	DecorationNoContraction               Decoration = 42
	DecorationInputAttachmentIndex        Decoration = 43
	DecorationAlignment                   Decoration = 44
	DecorationMaxByteOffset               Decoration = 45
	DecorationAlignmentId                 Decoration = 46
	DecorationMaxByteOffsetId             Decoration = 47
	DecorationNoSignedWrap                Decoration = 4469
	DecorationNoUnsignedWrap              Decoration = 4470
	DecorationExplicitInterpAMD           Decoration = 4999
	DecorationOverrideCoverageNV          Decoration = 5248
	DecorationPassthroughNV               Decoration = 5250
	DecorationViewportRelativeNV          Decoration = 5252
	DecorationSecondaryViewportRelativeNV Decoration = 5256
	DecorationPerPrimitiveNV              Decoration = 5271
	DecorationPerViewNV                   Decoration = 5272
	DecorationPerTaskNV                   Decoration = 5273
	DecorationPerVertexNV                 Decoration = 5285
	DecorationNonUniform                  Decoration = 5300
	DecorationNonUniformEXT               Decoration = 5300
	DecorationRestrictPointer             Decoration = 5355
	DecorationRestrictPointerEXT          Decoration = 5355
	DecorationAliasedPointer              Decoration = 5356
	DecorationAliasedPointerEXT           Decoration = 5356
	DecorationCounterBuffer               Decoration = 5634
	DecorationHlslCounterBufferGOOGLE     Decoration = 5634
	DecorationUserSemantic                Decoration = 5635
	DecorationHlslSemanticGOOGLE          Decoration = 5635
	DecorationUserTypeGOOGLE              Decoration = 5636
)

type BuiltIn uint32

func (v BuiltIn) Verify() error { return nil }

const (
	BuiltInPosition                    BuiltIn = 0
	BuiltInPointSize                   BuiltIn = 1
	BuiltInClipDistance                BuiltIn = 3
	BuiltInCullDistance                BuiltIn = 4
	BuiltInVertexId                    BuiltIn = 5
	BuiltInInstanceId                  BuiltIn = 6
	BuiltInPrimitiveId                 BuiltIn = 7
	BuiltInInvocationId                BuiltIn = 8
	BuiltInLayer                       BuiltIn = 9
	BuiltInViewportIndex               BuiltIn = 10
	BuiltInTessLevelOuter              BuiltIn = 11
	BuiltInTessLevelInner              BuiltIn = 12
	BuiltInTessCoord                   BuiltIn = 13
	BuiltInPatchVertices               BuiltIn = 14
	BuiltInFragCoord                   BuiltIn = 15
	BuiltInPointCoord                  BuiltIn = 16
	BuiltInFrontFacing                 BuiltIn = 17
	BuiltInSampleId                    BuiltIn = 18
	BuiltInSamplePosition              BuiltIn = 19
	BuiltInSampleMask                  BuiltIn = 20
	BuiltInFragDepth                   BuiltIn = 22
	BuiltInHelperInvocation            BuiltIn = 23
	BuiltInNumWorkgroups               BuiltIn = 24
	BuiltInWorkgroupSize               BuiltIn = 25
	BuiltInWorkgroupId                 BuiltIn = 26
	BuiltInLocalInvocationId           BuiltIn = 27
	BuiltInGlobalInvocationId          BuiltIn = 28
	BuiltInLocalInvocationIndex        BuiltIn = 29
	BuiltInWorkDim                     BuiltIn = 30
	BuiltInGlobalSize                  BuiltIn = 31
	BuiltInEnqueuedWorkgroupSize       BuiltIn = 32
	BuiltInGlobalOffset                BuiltIn = 33
	BuiltInGlobalLinearId              BuiltIn = 34
	BuiltInSubgroupSize                BuiltIn = 36
	BuiltInSubgroupMaxSize             BuiltIn = 37
	BuiltInNumSubgroups                BuiltIn = 38
	BuiltInNumEnqueuedSubgroups        BuiltIn = 39
	BuiltInSubgroupId                  BuiltIn = 40
	BuiltInSubgroupLocalInvocationId   BuiltIn = 41
	BuiltInVertexIndex                 BuiltIn = 42
	BuiltInInstanceIndex               BuiltIn = 43
	BuiltInSubgroupEqMask              BuiltIn = 4416
	BuiltInSubgroupGeMask              BuiltIn = 4417
	BuiltInSubgroupGtMask              BuiltIn = 4418
	BuiltInSubgroupLeMask              BuiltIn = 4419
	BuiltInSubgroupLtMask              BuiltIn = 4420
	BuiltInSubgroupEqMaskKHR           BuiltIn = 4416
	BuiltInSubgroupGeMaskKHR           BuiltIn = 4417
	BuiltInSubgroupGtMaskKHR           BuiltIn = 4418
	BuiltInSubgroupLeMaskKHR           BuiltIn = 4419
	BuiltInSubgroupLtMaskKHR           BuiltIn = 4420
	BuiltInBaseVertex                  BuiltIn = 4424
	BuiltInBaseInstance                BuiltIn = 4425
	BuiltInDrawIndex                   BuiltIn = 4426
	BuiltInDeviceIndex                 BuiltIn = 4438
	BuiltInViewIndex                   BuiltIn = 4440
	BuiltInBaryCoordNoPerspAMD         BuiltIn = 4992
	BuiltInBaryCoordNoPerspCentroidAMD BuiltIn = 4993
	BuiltInBaryCoordNoPerspSampleAMD   BuiltIn = 4994
	BuiltInBaryCoordSmoothAMD          BuiltIn = 4995
	BuiltInBaryCoordSmoothCentroidAMD  BuiltIn = 4996
	BuiltInBaryCoordSmoothSampleAMD    BuiltIn = 4997
	BuiltInBaryCoordPullModelAMD       BuiltIn = 4998
	BuiltInFragStencilRefEXT           BuiltIn = 5014
	BuiltInViewportMaskNV              BuiltIn = 5253
	BuiltInSecondaryPositionNV         BuiltIn = 5257
	BuiltInSecondaryViewportMaskNV     BuiltIn = 5258
	BuiltInPositionPerViewNV           BuiltIn = 5261
	BuiltInViewportMaskPerViewNV       BuiltIn = 5262
	BuiltInFullyCoveredEXT             BuiltIn = 5264
	BuiltInTaskCountNV                 BuiltIn = 5274
	BuiltInPrimitiveCountNV            BuiltIn = 5275
	BuiltInPrimitiveIndicesNV          BuiltIn = 5276
	BuiltInClipDistancePerViewNV       BuiltIn = 5277
	BuiltInCullDistancePerViewNV       BuiltIn = 5278
	BuiltInLayerPerViewNV              BuiltIn = 5279
	BuiltInMeshViewCountNV             BuiltIn = 5280
	BuiltInMeshViewIndicesNV           BuiltIn = 5281
	BuiltInBaryCoordNV                 BuiltIn = 5286
	BuiltInBaryCoordNoPerspNV          BuiltIn = 5287
	BuiltInFragSizeEXT                 BuiltIn = 5292
	BuiltInFragmentSizeNV              BuiltIn = 5292
	BuiltInFragInvocationCountEXT      BuiltIn = 5293
	BuiltInInvocationsPerPixelNV       BuiltIn = 5293
	BuiltInLaunchIdNV                  BuiltIn = 5319
	BuiltInLaunchSizeNV                BuiltIn = 5320
	BuiltInWorldRayOriginNV            BuiltIn = 5321
	BuiltInWorldRayDirectionNV         BuiltIn = 5322
	BuiltInObjectRayOriginNV           BuiltIn = 5323
	BuiltInObjectRayDirectionNV        BuiltIn = 5324
	BuiltInRayTminNV                   BuiltIn = 5325
	BuiltInRayTmaxNV                   BuiltIn = 5326
	BuiltInInstanceCustomIndexNV       BuiltIn = 5327
	BuiltInObjectToWorldNV             BuiltIn = 5330
	BuiltInWorldToObjectNV             BuiltIn = 5331
	BuiltInHitTNV                      BuiltIn = 5332
	BuiltInHitKindNV                   BuiltIn = 5333
	BuiltInIncomingRayFlagsNV          BuiltIn = 5351
	BuiltInWarpsPerSMNV                BuiltIn = 5374
	BuiltInSMCountNV                   BuiltIn = 5375
	BuiltInWarpIDNV                    BuiltIn = 5376
	BuiltInSMIDNV                      BuiltIn = 5377
)

type Scope uint32

func (v Scope) Verify() error { return nil }

const (
	ScopeCrossDevice    Scope = 0
	ScopeDevice         Scope = 1
	ScopeWorkgroup      Scope = 2
	ScopeSubgroup       Scope = 3
	ScopeInvocation     Scope = 4
	ScopeQueueFamily    Scope = 5
	ScopeQueueFamilyKHR Scope = 5
)

type GroupOperation uint32

func (v GroupOperation) Verify() error { return nil }

const (
	GroupOperationReduce                     GroupOperation = 0
	GroupOperationInclusiveScan              GroupOperation = 1
	GroupOperationExclusiveScan              GroupOperation = 2
	GroupOperationClusteredReduce            GroupOperation = 3
	GroupOperationPartitionedReduceNV        GroupOperation = 6
	GroupOperationPartitionedInclusiveScanNV GroupOperation = 7
	GroupOperationPartitionedExclusiveScanNV GroupOperation = 8
)

type KernelEnqueueFlags uint32

func (v KernelEnqueueFlags) Verify() error { return nil }

const (
	KernelEnqueueFlagsNoWait        KernelEnqueueFlags = 0
	KernelEnqueueFlagsWaitKernel    KernelEnqueueFlags = 1
	KernelEnqueueFlagsWaitWorkGroup KernelEnqueueFlags = 2
)

type Capability uint32

func (v Capability) Verify() error { return nil }

const (
	CapabilityMatrix                                       Capability = 0
	CapabilityShader                                       Capability = 1
	CapabilityGeometry                                     Capability = 2
	CapabilityTessellation                                 Capability = 3
	CapabilityAddresses                                    Capability = 4
	CapabilityLinkage                                      Capability = 5
	CapabilityKernel                                       Capability = 6
	CapabilityVector16                                     Capability = 7
	CapabilityFloat16Buffer                                Capability = 8
	CapabilityFloat16                                      Capability = 9
	CapabilityFloat64                                      Capability = 10
	CapabilityInt64                                        Capability = 11
	CapabilityInt64Atomics                                 Capability = 12
	CapabilityImageBasic                                   Capability = 13
	CapabilityImageReadWrite                               Capability = 14
	CapabilityImageMipmap                                  Capability = 15
	CapabilityPipes                                        Capability = 17
	CapabilityGroups                                       Capability = 18
	CapabilityDeviceEnqueue                                Capability = 19
	CapabilityLiteralSampler                               Capability = 20
	CapabilityAtomicStorage                                Capability = 21
	CapabilityInt16                                        Capability = 22
	CapabilityTessellationPointSize                        Capability = 23
	CapabilityGeometryPointSize                            Capability = 24
	CapabilityImageGatherExtended                          Capability = 25
	CapabilityStorageImageMultisample                      Capability = 27
	CapabilityUniformBufferArrayDynamicIndexing            Capability = 28
	CapabilitySampledImageArrayDynamicIndexing             Capability = 29
	CapabilityStorageBufferArrayDynamicIndexing            Capability = 30
	CapabilityStorageImageArrayDynamicIndexing             Capability = 31
	CapabilityClipDistance                                 Capability = 32
	CapabilityCullDistance                                 Capability = 33
	CapabilityImageCubeArray                               Capability = 34
	CapabilitySampleRateShading                            Capability = 35
	CapabilityImageRect                                    Capability = 36
	CapabilitySampledRect                                  Capability = 37
	CapabilityGenericPointer                               Capability = 38
	CapabilityInt8                                         Capability = 39
	CapabilityInputAttachment                              Capability = 40
	CapabilitySparseResidency                              Capability = 41
	CapabilityMinLod                                       Capability = 42
	CapabilitySampled1D                                    Capability = 43
	CapabilityImage1D                                      Capability = 44
	CapabilitySampledCubeArray                             Capability = 45
	CapabilitySampledBuffer                                Capability = 46
	CapabilityImageBuffer                                  Capability = 47
	CapabilityImageMSArray                                 Capability = 48
	CapabilityStorageImageExtendedFormats                  Capability = 49
	CapabilityImageQuery                                   Capability = 50
	CapabilityDerivativeControl                            Capability = 51
	CapabilityInterpolationFunction                        Capability = 52
	CapabilityTransformFeedback                            Capability = 53
	CapabilityGeometryStreams                              Capability = 54
	CapabilityStorageImageReadWithoutFormat                Capability = 55
	CapabilityStorageImageWriteWithoutFormat               Capability = 56
	CapabilityMultiViewport                                Capability = 57
	CapabilitySubgroupDispatch                             Capability = 58
	CapabilityNamedBarrier                                 Capability = 59
	CapabilityPipeStorage                                  Capability = 60
	CapabilityGroupNonUniform                              Capability = 61
	CapabilityGroupNonUniformVote                          Capability = 62
	CapabilityGroupNonUniformArithmetic                    Capability = 63
	CapabilityGroupNonUniformBallot                        Capability = 64
	CapabilityGroupNonUniformShuffle                       Capability = 65
	CapabilityGroupNonUniformShuffleRelative               Capability = 66
	CapabilityGroupNonUniformClustered                     Capability = 67
	CapabilityGroupNonUniformQuad                          Capability = 68
	CapabilityShaderLayer                                  Capability = 69
	CapabilityShaderViewportIndex                          Capability = 70
	CapabilitySubgroupBallotKHR                            Capability = 4423
	CapabilityDrawParameters                               Capability = 4427
	CapabilitySubgroupVoteKHR                              Capability = 4431
	CapabilityStorageBuffer16BitAccess                     Capability = 4433
	CapabilityStorageUniformBufferBlock16                  Capability = 4433
	CapabilityUniformAndStorageBuffer16BitAccess           Capability = 4434
	CapabilityStorageUniform16                             Capability = 4434
	CapabilityStoragePushConstant16                        Capability = 4435
	CapabilityStorageInputOutput16                         Capability = 4436
	CapabilityDeviceGroup                                  Capability = 4437
	CapabilityMultiView                                    Capability = 4439
	CapabilityVariablePointersStorageBuffer                Capability = 4441
	CapabilityVariablePointers                             Capability = 4442
	CapabilityAtomicStorageOps                             Capability = 4445
	CapabilitySampleMaskPostDepthCoverage                  Capability = 4447
	CapabilityStorageBuffer8BitAccess                      Capability = 4448
	CapabilityUniformAndStorageBuffer8BitAccess            Capability = 4449
	CapabilityStoragePushConstant8                         Capability = 4450
	CapabilityDenormPreserve                               Capability = 4464
	CapabilityDenormFlushToZero                            Capability = 4465
	CapabilitySignedZeroInfNanPreserve                     Capability = 4466
	CapabilityRoundingModeRTE                              Capability = 4467
	CapabilityRoundingModeRTZ                              Capability = 4468
	CapabilityFloat16ImageAMD                              Capability = 5008
	CapabilityImageGatherBiasLodAMD                        Capability = 5009
	CapabilityFragmentMaskAMD                              Capability = 5010
	CapabilityStencilExportEXT                             Capability = 5013
	CapabilityImageReadWriteLodAMD                         Capability = 5015
	CapabilityShaderClockKHR                               Capability = 5055
	CapabilitySampleMaskOverrideCoverageNV                 Capability = 5249
	CapabilityGeometryShaderPassthroughNV                  Capability = 5251
	CapabilityShaderViewportIndexLayerEXT                  Capability = 5254
	CapabilityShaderViewportIndexLayerNV                   Capability = 5254
	CapabilityShaderViewportMaskNV                         Capability = 5255
	CapabilityShaderStereoViewNV                           Capability = 5259
	CapabilityPerViewAttributesNV                          Capability = 5260
	CapabilityFragmentFullyCoveredEXT                      Capability = 5265
	CapabilityMeshShadingNV                                Capability = 5266
	CapabilityImageFootprintNV                             Capability = 5282
	CapabilityFragmentBarycentricNV                        Capability = 5284
	CapabilityComputeDerivativeGroupQuadsNV                Capability = 5288
	CapabilityFragmentDensityEXT                           Capability = 5291
	CapabilityShadingRateNV                                Capability = 5291
	CapabilityGroupNonUniformPartitionedNV                 Capability = 5297
	CapabilityShaderNonUniform                             Capability = 5301
	CapabilityShaderNonUniformEXT                          Capability = 5301
	CapabilityRuntimeDescriptorArray                       Capability = 5302
	CapabilityRuntimeDescriptorArrayEXT                    Capability = 5302
	CapabilityInputAttachmentArrayDynamicIndexing          Capability = 5303
	CapabilityInputAttachmentArrayDynamicIndexingEXT       Capability = 5303
	CapabilityUniformTexelBufferArrayDynamicIndexing       Capability = 5304
	CapabilityUniformTexelBufferArrayDynamicIndexingEXT    Capability = 5304
	CapabilityStorageTexelBufferArrayDynamicIndexing       Capability = 5305
	CapabilityStorageTexelBufferArrayDynamicIndexingEXT    Capability = 5305
	CapabilityUniformBufferArrayNonUniformIndexing         Capability = 5306
	CapabilityUniformBufferArrayNonUniformIndexingEXT      Capability = 5306
	CapabilitySampledImageArrayNonUniformIndexing          Capability = 5307
	CapabilitySampledImageArrayNonUniformIndexingEXT       Capability = 5307
	CapabilityStorageBufferArrayNonUniformIndexing         Capability = 5308
	CapabilityStorageBufferArrayNonUniformIndexingEXT      Capability = 5308
	CapabilityStorageImageArrayNonUniformIndexing          Capability = 5309
	CapabilityStorageImageArrayNonUniformIndexingEXT       Capability = 5309
	CapabilityInputAttachmentArrayNonUniformIndexing       Capability = 5310
	CapabilityInputAttachmentArrayNonUniformIndexingEXT    Capability = 5310
	CapabilityUniformTexelBufferArrayNonUniformIndexing    Capability = 5311
	CapabilityUniformTexelBufferArrayNonUniformIndexingEXT Capability = 5311
	CapabilityStorageTexelBufferArrayNonUniformIndexing    Capability = 5312
	CapabilityStorageTexelBufferArrayNonUniformIndexingEXT Capability = 5312
	CapabilityRayTracingNV                                 Capability = 5340
	CapabilityVulkanMemoryModel                            Capability = 5345
	CapabilityVulkanMemoryModelKHR                         Capability = 5345
	CapabilityVulkanMemoryModelDeviceScope                 Capability = 5346
	CapabilityVulkanMemoryModelDeviceScopeKHR              Capability = 5346
	CapabilityPhysicalStorageBufferAddresses               Capability = 5347
	CapabilityPhysicalStorageBufferAddressesEXT            Capability = 5347
	CapabilityComputeDerivativeGroupLinearNV               Capability = 5350
	CapabilityCooperativeMatrixNV                          Capability = 5357
	CapabilityFragmentShaderSampleInterlockEXT             Capability = 5363
	CapabilityFragmentShaderShadingRateInterlockEXT        Capability = 5372
	CapabilityShaderSMBuiltinsNV                           Capability = 5373
	CapabilityFragmentShaderPixelInterlockEXT              Capability = 5378
	CapabilityDemoteToHelperInvocationEXT                  Capability = 5379
	CapabilitySubgroupShuffleINTEL                         Capability = 5568
	CapabilitySubgroupBufferBlockIOINTEL                   Capability = 5569
	CapabilitySubgroupImageBlockIOINTEL                    Capability = 5570
	CapabilitySubgroupImageMediaBlockIOINTEL               Capability = 5579
	CapabilityIntegerFunctions2INTEL                       Capability = 5584
	CapabilitySubgroupAvcMotionEstimationINTEL             Capability = 5696
	CapabilitySubgroupAvcMotionEstimationIntraINTEL        Capability = 5697
	CapabilitySubgroupAvcMotionEstimationChromaINTEL       Capability = 5698
)
