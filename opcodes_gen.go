// This file was generated. DO NOT EDIT!

package spirv

const (
	opcodeNop                                     = 0
	opcodeUndef                                   = 1
	opcodeSourceContinued                         = 2
	opcodeSource                                  = 3
	opcodeSourceExtension                         = 4
	opcodeName                                    = 5
	opcodeMemberName                              = 6
	opcodeString                                  = 7
	opcodeLine                                    = 8
	opcodeExtension                               = 10
	opcodeExtInstImport                           = 11
	opcodeExtInst                                 = 12
	opcodeMemoryModel                             = 14
	opcodeEntryPoint                              = 15
	opcodeExecutionMode                           = 16
	opcodeCapability                              = 17
	opcodeTypeVoid                                = 19
	opcodeTypeBool                                = 20
	opcodeTypeInt                                 = 21
	opcodeTypeFloat                               = 22
	opcodeTypeVector                              = 23
	opcodeTypeMatrix                              = 24
	opcodeTypeImage                               = 25
	opcodeTypeSampler                             = 26
	opcodeTypeSampledImage                        = 27
	opcodeTypeArray                               = 28
	opcodeTypeRuntimeArray                        = 29
	opcodeTypeStruct                              = 30
	opcodeTypeOpaque                              = 31
	opcodeTypePointer                             = 32
	opcodeTypeFunction                            = 33
	opcodeTypeEvent                               = 34
	opcodeTypeDeviceEvent                         = 35
	opcodeTypeReserveId                           = 36
	opcodeTypeQueue                               = 37
	opcodeTypePipe                                = 38
	opcodeTypeForwardPointer                      = 39
	opcodeConstantTrue                            = 41
	opcodeConstantFalse                           = 42
	opcodeConstant                                = 43
	opcodeConstantComposite                       = 44
	opcodeConstantSampler                         = 45
	opcodeConstantNull                            = 46
	opcodeSpecConstantTrue                        = 48
	opcodeSpecConstantFalse                       = 49
	opcodeSpecConstant                            = 50
	opcodeSpecConstantComposite                   = 51
	opcodeSpecConstantOp                          = 52
	opcodeFunction                                = 54
	opcodeFunctionParameter                       = 55
	opcodeFunctionEnd                             = 56
	opcodeFunctionCall                            = 57
	opcodeVariable                                = 59
	opcodeImageTexelPointer                       = 60
	opcodeLoad                                    = 61
	opcodeStore                                   = 62
	opcodeCopyMemory                              = 63
	opcodeCopyMemorySized                         = 64
	opcodeAccessChain                             = 65
	opcodeInBoundsAccessChain                     = 66
	opcodePtrAccessChain                          = 67
	opcodeArrayLength                             = 68
	opcodeGenericPtrMemSemantics                  = 69
	opcodeInBoundsPtrAccessChain                  = 70
	opcodeDecorate                                = 71
	opcodeMemberDecorate                          = 72
	opcodeDecorationGroup                         = 73
	opcodeGroupDecorate                           = 74
	opcodeGroupMemberDecorate                     = 75
	opcodeVectorExtractDynamic                    = 77
	opcodeVectorInsertDynamic                     = 78
	opcodeVectorShuffle                           = 79
	opcodeCompositeConstruct                      = 80
	opcodeCompositeExtract                        = 81
	opcodeCompositeInsert                         = 82
	opcodeCopyObject                              = 83
	opcodeTranspose                               = 84
	opcodeSampledImage                            = 86
	opcodeImageSampleImplicitLod                  = 87
	opcodeImageSampleExplicitLod                  = 88
	opcodeImageSampleDrefImplicitLod              = 89
	opcodeImageSampleDrefExplicitLod              = 90
	opcodeImageSampleProjImplicitLod              = 91
	opcodeImageSampleProjExplicitLod              = 92
	opcodeImageSampleProjDrefImplicitLod          = 93
	opcodeImageSampleProjDrefExplicitLod          = 94
	opcodeImageFetch                              = 95
	opcodeImageGather                             = 96
	opcodeImageDrefGather                         = 97
	opcodeImageRead                               = 98
	opcodeImageWrite                              = 99
	opcodeImage                                   = 100
	opcodeImageQueryFormat                        = 101
	opcodeImageQueryOrder                         = 102
	opcodeImageQuerySizeLod                       = 103
	opcodeImageQuerySize                          = 104
	opcodeImageQueryLod                           = 105
	opcodeImageQueryLevels                        = 106
	opcodeImageQuerySamples                       = 107
	opcodeConvertFToU                             = 109
	opcodeConvertFToS                             = 110
	opcodeConvertSToF                             = 111
	opcodeConvertUToF                             = 112
	opcodeUConvert                                = 113
	opcodeSConvert                                = 114
	opcodeFConvert                                = 115
	opcodeQuantizeToF16                           = 116
	opcodeConvertPtrToU                           = 117
	opcodeSatConvertSToU                          = 118
	opcodeSatConvertUToS                          = 119
	opcodeConvertUToPtr                           = 120
	opcodePtrCastToGeneric                        = 121
	opcodeGenericCastToPtr                        = 122
	opcodeGenericCastToPtrExplicit                = 123
	opcodeBitcast                                 = 124
	opcodeSNegate                                 = 126
	opcodeFNegate                                 = 127
	opcodeIAdd                                    = 128
	opcodeFAdd                                    = 129
	opcodeISub                                    = 130
	opcodeFSub                                    = 131
	opcodeIMul                                    = 132
	opcodeFMul                                    = 133
	opcodeUDiv                                    = 134
	opcodeSDiv                                    = 135
	opcodeFDiv                                    = 136
	opcodeUMod                                    = 137
	opcodeSRem                                    = 138
	opcodeSMod                                    = 139
	opcodeFRem                                    = 140
	opcodeFMod                                    = 141
	opcodeVectorTimesScalar                       = 142
	opcodeMatrixTimesScalar                       = 143
	opcodeVectorTimesMatrix                       = 144
	opcodeMatrixTimesVector                       = 145
	opcodeMatrixTimesMatrix                       = 146
	opcodeOuterProduct                            = 147
	opcodeDot                                     = 148
	opcodeIAddCarry                               = 149
	opcodeISubBorrow                              = 150
	opcodeUMulExtended                            = 151
	opcodeSMulExtended                            = 152
	opcodeAny                                     = 154
	opcodeAll                                     = 155
	opcodeIsNan                                   = 156
	opcodeIsInf                                   = 157
	opcodeIsFinite                                = 158
	opcodeIsNormal                                = 159
	opcodeSignBitSet                              = 160
	opcodeLessOrGreater                           = 161
	opcodeOrdered                                 = 162
	opcodeUnordered                               = 163
	opcodeLogicalEqual                            = 164
	opcodeLogicalNotEqual                         = 165
	opcodeLogicalOr                               = 166
	opcodeLogicalAnd                              = 167
	opcodeLogicalNot                              = 168
	opcodeSelect                                  = 169
	opcodeIEqual                                  = 170
	opcodeINotEqual                               = 171
	opcodeUGreaterThan                            = 172
	opcodeSGreaterThan                            = 173
	opcodeUGreaterThanEqual                       = 174
	opcodeSGreaterThanEqual                       = 175
	opcodeULessThan                               = 176
	opcodeSLessThan                               = 177
	opcodeULessThanEqual                          = 178
	opcodeSLessThanEqual                          = 179
	opcodeFOrdEqual                               = 180
	opcodeFUnordEqual                             = 181
	opcodeFOrdNotEqual                            = 182
	opcodeFUnordNotEqual                          = 183
	opcodeFOrdLessThan                            = 184
	opcodeFUnordLessThan                          = 185
	opcodeFOrdGreaterThan                         = 186
	opcodeFUnordGreaterThan                       = 187
	opcodeFOrdLessThanEqual                       = 188
	opcodeFUnordLessThanEqual                     = 189
	opcodeFOrdGreaterThanEqual                    = 190
	opcodeFUnordGreaterThanEqual                  = 191
	opcodeShiftRightLogical                       = 194
	opcodeShiftRightArithmetic                    = 195
	opcodeShiftLeftLogical                        = 196
	opcodeBitwiseOr                               = 197
	opcodeBitwiseXor                              = 198
	opcodeBitwiseAnd                              = 199
	opcodeNot                                     = 200
	opcodeBitFieldInsert                          = 201
	opcodeBitFieldSExtract                        = 202
	opcodeBitFieldUExtract                        = 203
	opcodeBitReverse                              = 204
	opcodeBitCount                                = 205
	opcodeDPdx                                    = 207
	opcodeDPdy                                    = 208
	opcodeFwidth                                  = 209
	opcodeDPdxFine                                = 210
	opcodeDPdyFine                                = 211
	opcodeFwidthFine                              = 212
	opcodeDPdxCoarse                              = 213
	opcodeDPdyCoarse                              = 214
	opcodeFwidthCoarse                            = 215
	opcodeEmitVertex                              = 218
	opcodeEndPrimitive                            = 219
	opcodeEmitStreamVertex                        = 220
	opcodeEndStreamPrimitive                      = 221
	opcodeControlBarrier                          = 224
	opcodeMemoryBarrier                           = 225
	opcodeAtomicLoad                              = 227
	opcodeAtomicStore                             = 228
	opcodeAtomicExchange                          = 229
	opcodeAtomicCompareExchange                   = 230
	opcodeAtomicCompareExchangeWeak               = 231
	opcodeAtomicIIncrement                        = 232
	opcodeAtomicIDecrement                        = 233
	opcodeAtomicIAdd                              = 234
	opcodeAtomicISub                              = 235
	opcodeAtomicSMin                              = 236
	opcodeAtomicUMin                              = 237
	opcodeAtomicSMax                              = 238
	opcodeAtomicUMax                              = 239
	opcodeAtomicAnd                               = 240
	opcodeAtomicOr                                = 241
	opcodeAtomicXor                               = 242
	opcodePhi                                     = 245
	opcodeLoopMerge                               = 246
	opcodeSelectionMerge                          = 247
	opcodeLabel                                   = 248
	opcodeBranch                                  = 249
	opcodeBranchConditional                       = 250
	opcodeSwitch                                  = 251
	opcodeKill                                    = 252
	opcodeReturn                                  = 253
	opcodeReturnValue                             = 254
	opcodeUnreachable                             = 255
	opcodeLifetimeStart                           = 256
	opcodeLifetimeStop                            = 257
	opcodeGroupAsyncCopy                          = 259
	opcodeGroupWaitEvents                         = 260
	opcodeGroupAll                                = 261
	opcodeGroupAny                                = 262
	opcodeGroupBroadcast                          = 263
	opcodeGroupIAdd                               = 264
	opcodeGroupFAdd                               = 265
	opcodeGroupFMin                               = 266
	opcodeGroupUMin                               = 267
	opcodeGroupSMin                               = 268
	opcodeGroupFMax                               = 269
	opcodeGroupUMax                               = 270
	opcodeGroupSMax                               = 271
	opcodeReadPipe                                = 274
	opcodeWritePipe                               = 275
	opcodeReservedReadPipe                        = 276
	opcodeReservedWritePipe                       = 277
	opcodeReserveReadPipePackets                  = 278
	opcodeReserveWritePipePackets                 = 279
	opcodeCommitReadPipe                          = 280
	opcodeCommitWritePipe                         = 281
	opcodeIsValidReserveId                        = 282
	opcodeGetNumPipePackets                       = 283
	opcodeGetMaxPipePackets                       = 284
	opcodeGroupReserveReadPipePackets             = 285
	opcodeGroupReserveWritePipePackets            = 286
	opcodeGroupCommitReadPipe                     = 287
	opcodeGroupCommitWritePipe                    = 288
	opcodeEnqueueMarker                           = 291
	opcodeEnqueueKernel                           = 292
	opcodeGetKernelNDrangeSubGroupCount           = 293
	opcodeGetKernelNDrangeMaxSubGroupSize         = 294
	opcodeGetKernelWorkGroupSize                  = 295
	opcodeGetKernelPreferredWorkGroupSizeMultiple = 296
	opcodeRetainEvent                             = 297
	opcodeReleaseEvent                            = 298
	opcodeCreateUserEvent                         = 299
	opcodeIsValidEvent                            = 300
	opcodeSetUserEventStatus                      = 301
	opcodeCaptureEventProfilingInfo               = 302
	opcodeGetDefaultQueue                         = 303
	opcodeBuildNDRange                            = 304
	opcodeImageSparseSampleImplicitLod            = 305
	opcodeImageSparseSampleExplicitLod            = 306
	opcodeImageSparseSampleDrefImplicitLod        = 307
	opcodeImageSparseSampleDrefExplicitLod        = 308
	opcodeImageSparseSampleProjImplicitLod        = 309
	opcodeImageSparseSampleProjExplicitLod        = 310
	opcodeImageSparseSampleProjDrefImplicitLod    = 311
	opcodeImageSparseSampleProjDrefExplicitLod    = 312
	opcodeImageSparseFetch                        = 313
	opcodeImageSparseGather                       = 314
	opcodeImageSparseDrefGather                   = 315
	opcodeImageSparseTexelsResident               = 316
	opcodeNoLine                                  = 317
	opcodeAtomicFlagTestAndSet                    = 318
	opcodeAtomicFlagClear                         = 319
	opcodeImageSparseRead                         = 320
	opcodeSizeOf                                  = 321
	opcodeTypePipeStorage                         = 322
	opcodeConstantPipeStorage                     = 323
	opcodeCreatePipeFromPipeStorage               = 324
	opcodeGetKernelLocalSizeForSubgroupCount      = 325
	opcodeGetKernelMaxNumSubgroups                = 326
	opcodeTypeNamedBarrier                        = 327
	opcodeNamedBarrierInitialize                  = 328
	opcodeMemoryNamedBarrier                      = 329
	opcodeModuleProcessed                         = 330
	opcodeExecutionModeId                         = 331
	opcodeDecorateId                              = 332
	opcodeGroupNonUniformElect                    = 333
	opcodeGroupNonUniformAll                      = 334
	opcodeGroupNonUniformAny                      = 335
	opcodeGroupNonUniformAllEqual                 = 336
	opcodeGroupNonUniformBroadcast                = 337
	opcodeGroupNonUniformBroadcastFirst           = 338
	opcodeGroupNonUniformBallot                   = 339
	opcodeGroupNonUniformInverseBallot            = 340
	opcodeGroupNonUniformBallotBitExtract         = 341
	opcodeGroupNonUniformBallotBitCount           = 342
	opcodeGroupNonUniformBallotFindLSB            = 343
	opcodeGroupNonUniformBallotFindMSB            = 344
	opcodeGroupNonUniformShuffle                  = 345
	opcodeGroupNonUniformShuffleXor               = 346
	opcodeGroupNonUniformShuffleUp                = 347
	opcodeGroupNonUniformShuffleDown              = 348
	opcodeGroupNonUniformIAdd                     = 349
	opcodeGroupNonUniformFAdd                     = 350
	opcodeGroupNonUniformIMul                     = 351
	opcodeGroupNonUniformFMul                     = 352
	opcodeGroupNonUniformSMin                     = 353
	opcodeGroupNonUniformUMin                     = 354
	opcodeGroupNonUniformFMin                     = 355
	opcodeGroupNonUniformSMax                     = 356
	opcodeGroupNonUniformUMax                     = 357
	opcodeGroupNonUniformFMax                     = 358
	opcodeGroupNonUniformBitwiseAnd               = 359
	opcodeGroupNonUniformBitwiseOr                = 360
	opcodeGroupNonUniformBitwiseXor               = 361
	opcodeGroupNonUniformLogicalAnd               = 362
	opcodeGroupNonUniformLogicalOr                = 363
	opcodeGroupNonUniformLogicalXor               = 364
	opcodeGroupNonUniformQuadBroadcast            = 365
	opcodeGroupNonUniformQuadSwap                 = 366
	opcodeCopyLogical                             = 400
	opcodePtrEqual                                = 401
	opcodePtrNotEqual                             = 402
	opcodePtrDiff                                 = 403
	opcodeSubgroupBallotKHR                       = 4421
	opcodeSubgroupFirstInvocationKHR              = 4422
	opcodeSubgroupAllKHR                          = 4428
	opcodeSubgroupAnyKHR                          = 4429
	opcodeSubgroupAllEqualKHR                     = 4430
	opcodeSubgroupReadInvocationKHR               = 4432
	opcodeGroupIAddNonUniformAMD                  = 5000
	opcodeGroupFAddNonUniformAMD                  = 5001
	opcodeGroupFMinNonUniformAMD                  = 5002
	opcodeGroupUMinNonUniformAMD                  = 5003
	opcodeGroupSMinNonUniformAMD                  = 5004
	opcodeGroupFMaxNonUniformAMD                  = 5005
	opcodeGroupUMaxNonUniformAMD                  = 5006
	opcodeGroupSMaxNonUniformAMD                  = 5007
	opcodeImageSampleFootprintNV                  = 5283
	opcodeGroupNonUniformPartitionNV              = 5296
	opcodeSubgroupShuffleINTEL                    = 5571
	opcodeSubgroupShuffleDownINTEL                = 5572
	opcodeSubgroupShuffleUpINTEL                  = 5573
	opcodeSubgroupShuffleXorINTEL                 = 5574
	opcodeSubgroupBlockReadINTEL                  = 5575
	opcodeSubgroupBlockWriteINTEL                 = 5576
	opcodeSubgroupImageBlockReadINTEL             = 5577
	opcodeSubgroupImageBlockWriteINTEL            = 5578
	opcodeSubgroupImageMediaBlockReadINTEL        = 5580
	opcodeSubgroupImageMediaBlockWriteINTEL       = 5581
	opcodeDecorateString                          = 5632
	opcodeMemberDecorateString                    = 5633
)
