package mos_6502

type Opcode uint8

const (
	OPCODE_HALT Opcode = 0xF2 // artificial, would crash a real cpu
	OPCODE_NOP  Opcode = 0xEA

	OPCODE_LDA_IMMEDIATE Opcode = 0xA9
	OPCODE_LDX_IMMEDIATE Opcode = 0xA2
	OPCODE_LDY_IMMEDIATE Opcode = 0xA0

	OPCODE_LDA_ZP Opcode = 0xA5
	OPCODE_LDX_ZP Opcode = 0xA6
	OPCODE_LDY_ZP Opcode = 0xA4

	OPCODE_LDA_ZP_X Opcode = 0xB5
	OPCODE_LDX_ZP_Y Opcode = 0xB6
	OPCODE_LDY_ZP_X Opcode = 0xB4

	OPCODE_LDA_ABSOLUTE Opcode = 0xAD
	OPCODE_LDX_ABSOLUTE Opcode = 0xAE
	OPCODE_LDY_ABSOLUTE Opcode = 0xAC

	OPCODE_LDA_ABSOLUTE_X Opcode = 0xBD
	OPCODE_LDA_ABSOLUTE_Y Opcode = 0xB9
	OPCODE_LDX_ABSOLUTE_Y Opcode = 0xBE
	OPCODE_LDY_ABSOLUTE_X Opcode = 0xBC

	OPCODE_LDA_INDEXED_INDIRECT_X Opcode = 0xA1 // add X to address, load address value as new address, load address value
	OPCODE_LDA_INDIRECT_INDEXED_Y Opcode = 0xB1 // load address, add Y to value as new address, load address value

	OPCODE_STA_ZP       Opcode = 0x85
	OPCODE_STX_ZP       Opcode = 0x86
	OPCODE_STY_ZP       Opcode = 0x84
	OPCODE_STA_ZP_X     Opcode = 0x95
	OPCODE_STX_ZP_Y     Opcode = 0x96
	OPCODE_STY_ZP_X     Opcode = 0x94
	OPCODE_STA_ABSOLUTE Opcode = 0x8D
	OPCODE_STX_ABSOLUTE Opcode = 0x8E
	OPCODE_STY_ABSOLUTE Opcode = 0x8C

	OPCODE_STA_ABSOLUTE_X         Opcode = 0x9D
	OPCODE_STA_ABSOLUTE_Y         Opcode = 0x99
	OPCODE_STA_INDEXED_INDIRECT_X Opcode = 0x81
	//TODO OPCODE_STA_INDIRECT_INDEXED_Y Opcode = 0x91

	OPCODE_INY_IMPLIED Opcode = 0xC8
	OPCODE_INX_IMPLIED Opcode = 0xE8
	OPCODE_DEY_IMPLIED Opcode = 0x88
	OPCODE_DEX_IMPLIED Opcode = 0xCA

	OPCODE_TAX_IMPLIED Opcode = 0xAA
	OPCODE_TAY_IMPLIED Opcode = 0xA8
	OPCODE_TSX_IMPLIED Opcode = 0xBA
	OPCODE_TXA_IMPLIED Opcode = 0x8A
	OPCODE_TYA_IMPLIED Opcode = 0x98
	OPCODE_TXS_IMPLIED Opcode = 0x9A

	OPCODE_JMP_ABSOLUTE Opcode = 0x4C
	OPCODE_JSR_ABSOLUTE Opcode = 0x20
	OPCODE_RTS_IMPLIED  Opcode = 0x60

	OPCODE_PHA_IMPLIED Opcode = 0x48
	OPCODE_PHP_IMPLIED Opcode = 0x08
	OPCODE_PLA_IMPLIED Opcode = 0x68
	OPCODE_PLP_IMPLIED Opcode = 0x28

	OPCODE_BCC_RELATIVE Opcode = 0x90
	OPCODE_BCS_RELATIVE Opcode = 0xB0
	OPCODE_BEQ_RELATIVE Opcode = 0xF0
	OPCODE_BMI_RELATIVE Opcode = 0x30
	OPCODE_BNE_RELATIVE Opcode = 0xD0
	OPCODE_BPL_RELATIVE Opcode = 0x10
	OPCODE_BVC_RELATIVE Opcode = 0x50
	OPCODE_BVS_RELATIVE Opcode = 0x70

	OPCODE_ADC_IMMEDIATE Opcode = 0x69
	OPCODE_ADC_ZP        Opcode = 0x65
	//OPCODE_ADC_ZP_X        Opcode = 0x75
	OPCODE_ADC_ABSOLUTE Opcode = 0x6D
	//OPCODE_ADC_ABSOLUTE_X  Opcode = 0x7D
	//OPCODE_ADC_ABSOLUTE_Y  Opcode = 0x79
	//OPCODE_ADC_INDEXED_INDIRECT_X  Opcode = 0x61
	//OPCODE_ADC_INDIRECT_INDEXED_Y  Opcode = 0x71

	OPCODE_CMP_IMMEDIATE Opcode = 0xC9
	OPCODE_CMP_ZP        Opcode = 0xC5
	OPCODE_CMP_ABSOLUTE  Opcode = 0xCD

	OPCODE_SBC_IMMEDIATE Opcode = 0xE9
	OPCODE_SBC_ZP        Opcode = 0xE5
	OPCODE_SBC_ABSOLUTE  Opcode = 0xED

	//TODO OPCODE_BIT_ABSOLUTE Opcode = 0x24
	//TODO OPCODE_BIT_ZP       Opcode = 0x2C

	OPCODE_CLC_IMPLIED Opcode = 0x18
	OPCODE_CLD_IMPLIED Opcode = 0xD8
	OPCODE_CLI_IMPLIED Opcode = 0x58
	OPCODE_CLV_IMPLIED Opcode = 0xB8
	OPCODE_SEC_IMPLIED Opcode = 0x38
	OPCODE_SED_IMPLIED Opcode = 0xF8
	OPCODE_SEI_IMPLIED Opcode = 0x78

	OPCODE_AND_IMMEDIATE Opcode = 0x29
	OPCODE_EOR_IMMEDIATE Opcode = 0x49
	OPCODE_ORA_IMMEDIATE Opcode = 0x09

	OPCODE_AND_ZP Opcode = 0x25
	OPCODE_EOR_ZP Opcode = 0x45
	OPCODE_ORA_ZP Opcode = 0x05

	OPCODE_AND_ZP_X Opcode = 0x35
	OPCODE_EOR_ZP_X Opcode = 0x55
	OPCODE_ORA_ZP_X Opcode = 0x15

	OPCODE_AND_ABSOLUTE Opcode = 0x2D
	OPCODE_EOR_ABSOLUTE Opcode = 0x4D
	OPCODE_ORA_ABSOLUTE Opcode = 0x0D

	OPCODE_AND_ABSOLUTE_X Opcode = 0x3D
	OPCODE_EOR_ABSOLUTE_X Opcode = 0x5D
	OPCODE_ORA_ABSOLUTE_X Opcode = 0x1D

	OPCODE_AND_ABSOLUTE_Y Opcode = 0x39
	OPCODE_EOR_ABSOLUTE_Y Opcode = 0x59
	OPCODE_ORA_ABSOLUTE_Y Opcode = 0x19

	OPCODE_AND_INDEXED_INDIRECT_X Opcode = 0x21
	OPCODE_EOR_INDEXED_INDIRECT_X Opcode = 0x41
	OPCODE_ORA_INDEXED_INDIRECT_X Opcode = 0x01

	OPCODE_AND_INDIRECT_INDEXED_Y Opcode = 0x31
	OPCODE_EOR_INDIRECT_INDEXED_Y Opcode = 0x51
	OPCODE_ORA_INDIRECT_INDEXED_Y Opcode = 0x11

	OPCODE_ASL_IMPLIED    Opcode = 0x0A
	OPCODE_ASL_ZP         Opcode = 0x06
	OPCODE_ASL_ZP_X       Opcode = 0x16
	OPCODE_ASL_ABSOLUTE   Opcode = 0x0E
	OPCODE_ASL_ABSOLUTE_X Opcode = 0x1E

	OPCODE_LSR_IMPLIED    Opcode = 0x4A
	OPCODE_LSR_ZP         Opcode = 0x46
	OPCODE_LSR_ZP_X       Opcode = 0x56
	OPCODE_LSR_ABSOLUTE   Opcode = 0x4E
	OPCODE_LSR_ABSOLUTE_X Opcode = 0x5E

	OPCODE_ROL_IMPLIED    Opcode = 0x2A
	OPCODE_ROL_ZP         Opcode = 0x26
	OPCODE_ROL_ZP_X       Opcode = 0x36
	OPCODE_ROL_ABSOLUTE   Opcode = 0x2E
	OPCODE_ROL_ABSOLUTE_X Opcode = 0x3E

	OPCODE_ROR_IMPLIED    Opcode = 0x6A
	OPCODE_ROR_ZP         Opcode = 0x66
	OPCODE_ROR_ZP_X       Opcode = 0x76
	OPCODE_ROR_ABSOLUTE   Opcode = 0x6E
	OPCODE_ROR_ABSOLUTE_X Opcode = 0x7E
)

// InstructionSet
// Instructions are one cycle longer than array elements as fetch op is part of the cpu
var InstructionSet = map[Opcode][]Cycle{
	OPCODE_NOP: {&CycleWait{}},
	// LD*
	OPCODE_LDA_IMMEDIATE: {&CycleFetchImmediateOperandToRegister{REGISTER_A}},
	OPCODE_LDX_IMMEDIATE: {&CycleFetchImmediateOperandToRegister{REGISTER_X}},
	OPCODE_LDY_IMMEDIATE: {&CycleFetchImmediateOperandToRegister{REGISTER_Y}},

	OPCODE_LDA_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{REGISTER_Y}},

	OPCODE_LDA_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ZP_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_Y}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{REGISTER_Y}},

	OPCODE_LDA_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{REGISTER_Y}},

	OPCODE_LDA_ABSOLUTE_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleFetchEffectiveWithPageFix{REGISTER_A}},
	OPCODE_LDA_ABSOLUTE_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_Y}, &CycleFetchEffectiveWithPageFix{REGISTER_A}},
	OPCODE_LDX_ABSOLUTE_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_Y}, &CycleFetchEffectiveWithPageFix{REGISTER_X}},
	OPCODE_LDY_ABSOLUTE_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleFetchEffectiveWithPageFix{REGISTER_Y}},

	OPCODE_LDA_INDEXED_INDIRECT_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleIndexedIndirectHigh{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDA_INDIRECT_INDEXED_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleIndirectIndexedFetchHighByte{}, &CycleFetchEffectiveWithPageFix{REGISTER_A}},

	// IN*/DE*
	OPCODE_INX_IMPLIED: {&CycleIncImplied{REGISTER_X}},
	OPCODE_INY_IMPLIED: {&CycleIncImplied{REGISTER_Y}},
	OPCODE_DEX_IMPLIED: {&CycleDecImplied{REGISTER_X}},
	OPCODE_DEY_IMPLIED: {&CycleDecImplied{REGISTER_Y}},

	// ST*
	OPCODE_STA_ZP: {&CycleFetchAddressOperandLow{}, &CycleWriteRegisterEffective{REGISTER_A}},
	OPCODE_STX_ZP: {&CycleFetchAddressOperandLow{}, &CycleWriteRegisterEffective{REGISTER_X}},
	OPCODE_STY_ZP: {&CycleFetchAddressOperandLow{}, &CycleWriteRegisterEffective{REGISTER_Y}},

	OPCODE_STA_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleWriteRegisterEffective{REGISTER_A}},
	OPCODE_STX_ZP_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_Y}, &CycleWriteRegisterEffective{REGISTER_X}},
	OPCODE_STY_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleWriteRegisterEffective{REGISTER_Y}},

	OPCODE_STA_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleWriteRegisterEffective{REGISTER_A}},
	OPCODE_STX_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleWriteRegisterEffective{REGISTER_X}},
	OPCODE_STY_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleWriteRegisterEffective{REGISTER_Y}},

	OPCODE_STA_ABSOLUTE_X:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleCalcEffectiveAddrIndexed{}, &CycleWriteRegisterEffective{REGISTER_A}},
	OPCODE_STA_ABSOLUTE_Y:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_Y}, &CycleCalcEffectiveAddrIndexed{}, &CycleWriteRegisterEffective{REGISTER_A}},
	OPCODE_STA_INDEXED_INDIRECT_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleIndexedIndirectHigh{}, &CycleWriteRegisterEffective{REGISTER_A}},

	// Branch
	OPCODE_BCC_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_C, false}}, // 2-4 cycles!
	OPCODE_BCS_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_C, true}},  // 2-4 cycles!
	OPCODE_BNE_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_Z, false}}, // 2-4 cycles!
	OPCODE_BEQ_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_Z, true}},  // 2-4 cycles!
	OPCODE_BPL_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_N, false}}, // 2-4 cycles!
	OPCODE_BMI_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_N, true}},  // 2-4 cycles!
	OPCODE_BVC_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_V, false}}, // 2-4 cycles!
	OPCODE_BVS_RELATIVE: {&CycleBranchFetchOp{PROCESSOR_STATUS_FLAG_V, true}},  // 2-4 cycles!

	// TX
	OPCODE_TAX_IMPLIED: {&CycleCopyRegister{Source: REGISTER_A, Target: REGISTER_X}},
	OPCODE_TXA_IMPLIED: {&CycleCopyRegister{Source: REGISTER_X, Target: REGISTER_A}},
	OPCODE_TSX_IMPLIED: {&CycleCopyRegister{Source: REGISTER_SP, Target: REGISTER_X}},
	OPCODE_TXS_IMPLIED: {&CycleCopyRegister{Source: REGISTER_X, Target: REGISTER_SP}},
	OPCODE_TAY_IMPLIED: {&CycleCopyRegister{Source: REGISTER_A, Target: REGISTER_Y}},
	OPCODE_TYA_IMPLIED: {&CycleCopyRegister{Source: REGISTER_Y, Target: REGISTER_A}},

	// JMP
	OPCODE_JMP_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleCopyPclFetchPch{}},
	OPCODE_JSR_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleWait{}, &CycleJsrPchPush{}, &CycleJsrPclPush{}, &CycleCopyPclFetchPch{}},
	OPCODE_RTS_IMPLIED:  {&CycleTrash{}, &CycleIncSp{}, &CycleRtPullPcl{}, &CycleRtPullPch{}, &CycleRtIncPc{}},

	// STACK MEM
	OPCODE_PHA_IMPLIED: {&CycleDummyReadInstruction{}, &CyclePushAcc{}},
	OPCODE_PHP_IMPLIED: {&CycleDummyReadInstruction{}, &CyclePushStatus{}},
	OPCODE_PLA_IMPLIED: {&CycleDummyReadInstruction{}, &CycleDecSp{}, &CyclePullAcc{}},
	OPCODE_PLP_IMPLIED: {&CycleDummyReadInstruction{}, &CycleDecSp{}, &CyclePullStatus{}},

	// ARITH
	// TODO refactor
	OPCODE_ADC_IMMEDIATE: {&CycleAddWithCarryImmediate{}},
	OPCODE_ADC_ZP:        {&CycleFetchAddressOperandLow{}, &CycleAddWithCarryEffective{}},
	OPCODE_ADC_ABSOLUTE:  {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleAddWithCarryEffective{}},

	OPCODE_SBC_IMMEDIATE: {&CycleSubWithBorrowImmediate{}},
	OPCODE_SBC_ZP:        {&CycleFetchAddressOperandLow{}, &CycleSubWithBorrowEffective{}},
	OPCODE_SBC_ABSOLUTE:  {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleSubWithBorrowEffective{}},

	OPCODE_CMP_IMMEDIATE: {&CycleCmpImmediate{}},
	OPCODE_CMP_ZP:        {&CycleFetchAddressOperandLow{}, &CycleCmpEffective{}},
	OPCODE_CMP_ABSOLUTE:  {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleCmpEffective{}},

	// BINARY LOGIC
	OPCODE_AND_IMMEDIATE: {&CycleBinaryLogicOpApplyImmediate{LogicOps.LogicOpAnd}},
	OPCODE_ORA_IMMEDIATE: {&CycleBinaryLogicOpApplyImmediate{LogicOps.LogicOpOr}},
	OPCODE_EOR_IMMEDIATE: {&CycleBinaryLogicOpApplyImmediate{LogicOps.LogicOpXor}},

	OPCODE_AND_ZP: {&CycleFetchAddressOperandLow{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_ZP: {&CycleFetchAddressOperandLow{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_ZP: {&CycleFetchAddressOperandLow{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},

	OPCODE_AND_ZP_X:               {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_ZP_X:               {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_ZP_X:               {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},
	OPCODE_AND_ABSOLUTE:           {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_ABSOLUTE:           {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_ABSOLUTE:           {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},
	OPCODE_AND_ABSOLUTE_X:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_ABSOLUTE_X:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_ABSOLUTE_X:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},
	OPCODE_AND_ABSOLUTE_Y:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_Y}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_ABSOLUTE_Y:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_Y}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_ABSOLUTE_Y:         {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHighIndexed{REGISTER_Y}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},
	OPCODE_AND_INDEXED_INDIRECT_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleIndexedIndirectHigh{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_INDEXED_INDIRECT_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleIndexedIndirectHigh{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_INDEXED_INDIRECT_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleIndexedIndirectHigh{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},
	OPCODE_AND_INDIRECT_INDEXED_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleIndirectIndexedFetchHighByte{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpAnd}},
	OPCODE_EOR_INDIRECT_INDEXED_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleIndirectIndexedFetchHighByte{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpXor}},
	OPCODE_ORA_INDIRECT_INDEXED_Y: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleIndirectIndexedFetchHighByte{}, &CycleBinaryLogicOpApplyEffectiveWithPageFix{LogicOps.LogicOpOr}},

	// UNARY LOGIC
	OPCODE_ASL_IMPLIED: {&CycleUnaryLogicOpApply{LogicOps.LogicOpAsl, REGISTER_A}},
	OPCODE_LSR_IMPLIED: {&CycleUnaryLogicOpApply{LogicOps.LogicOpLsr, REGISTER_A}},
	OPCODE_ROL_IMPLIED: {&CycleUnaryLogicOpApply{LogicOps.LogicOpRol, REGISTER_A}},
	OPCODE_ROR_IMPLIED: {&CycleUnaryLogicOpApply{LogicOps.LogicOpRor, REGISTER_A}},

	OPCODE_ASL_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpAsl, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_LSR_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpLsr, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROL_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRol, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROR_ZP: {&CycleFetchAddressOperandLow{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRor, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},

	OPCODE_ASL_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpAsl, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_LSR_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpLsr, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROL_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRol, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROR_ZP_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressZpIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRor, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},

	OPCODE_ASL_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpAsl, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_LSR_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpLsr, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROL_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRol, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROR_ABSOLUTE: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRor, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},

	OPCODE_ASL_ABSOLUTE_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpAsl, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_LSR_ABSOLUTE_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpLsr, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROL_ABSOLUTE_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRol, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},
	OPCODE_ROR_ABSOLUTE_X: {&CycleFetchAddressOperandLow{}, &CycleFetchAddressOperandHigh{}, &CycleFetchAddressOperandHighIndexed{REGISTER_X}, &CycleFetchEffective{DATA_LATCH}, &CycleUnaryLogicOpApply{LogicOps.LogicOpRor, DATA_LATCH}, &CycleWriteRegisterEffective{DATA_LATCH}},

	// STATUS
	OPCODE_CLC_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_C}},
	OPCODE_CLD_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_D}},
	OPCODE_CLI_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_I}},
	OPCODE_CLV_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_V}},
	OPCODE_SEC_IMPLIED: {&CycleProcFlagSet{PROCESSOR_STATUS_FLAG_C}},
	OPCODE_SED_IMPLIED: {&CycleProcFlagSet{PROCESSOR_STATUS_FLAG_D}},
	OPCODE_SEI_IMPLIED: {&CycleProcFlagSet{PROCESSOR_STATUS_FLAG_I}},
}
