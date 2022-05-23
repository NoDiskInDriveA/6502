package mos_6502

import "github.com/NoDiskInDriveA/6502/internal/helper/pc"

type Opcode uint8

const (
	OPCODE_HALT Opcode = 0xF2 // artificial, would crash a real cpu
	OPCODE_NOP  Opcode = 0xEA

	OPCODE_LDA_IMMEDIATE  Opcode = 0xA9
	OPCODE_LDX_IMMEDIATE  Opcode = 0xA2
	OPCODE_LDY_IMMEDIATE  Opcode = 0xA0
	OPCODE_LDA_ZP         Opcode = 0xA5
	OPCODE_LDX_ZP         Opcode = 0xA6
	OPCODE_LDY_ZP         Opcode = 0xA4
	OPCODE_LDA_ZP_X       Opcode = 0xB5
	OPCODE_LDX_ZP_Y       Opcode = 0xB6
	OPCODE_LDY_ZP_X       Opcode = 0xB4
	OPCODE_LDA_ABSOLUTE   Opcode = 0xAD
	OPCODE_LDX_ABSOLUTE   Opcode = 0xAE
	OPCODE_LDY_ABSOLUTE   Opcode = 0xAC
	OPCODE_LDA_ABSOLUTE_X Opcode = 0xBD
	OPCODE_LDA_ABSOLUTE_Y Opcode = 0xB9
	OPCODE_LDX_ABSOLUTE_Y Opcode = 0xBE
	OPCODE_LDY_ABSOLUTE_X Opcode = 0xBC

	OPCODE_STA_ZP       Opcode = 0x85
	OPCODE_STX_ZP       Opcode = 0x86
	OPCODE_STY_ZP       Opcode = 0x84
	OPCODE_STA_ZP_X     Opcode = 0x95
	OPCODE_STX_ZP_Y     Opcode = 0x96
	OPCODE_STY_ZP_X     Opcode = 0x94
	OPCODE_STA_ABSOLUTE Opcode = 0x8D
	OPCODE_STX_ABSOLUTE Opcode = 0x8E
	OPCODE_STY_ABSOLUTE Opcode = 0x8C
	//# WIP
	OPCODE_STA_ABSOLUTE_X Opcode = 0x9D
	OPCODE_STA_ABSOLUTE_Y Opcode = 0x99
	OPCODE_STA_INDIRECT_X Opcode = 0x81
	OPCODE_STA_INDIRECT_Y Opcode = 0x91
	//# WIP END
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
	//# WIP
	OPCODE_BCC_RELATIVE Opcode = 0x90
	OPCODE_BCS_RELATIVE Opcode = 0xB0
	OPCODE_BEQ_RELATIVE Opcode = 0xF0
	OPCODE_BMI_RELATIVE Opcode = 0x30
	OPCODE_BNE_RELATIVE Opcode = 0xD0
	OPCODE_BPL_RELATIVE Opcode = 0x10
	OPCODE_BVC_RELATIVE Opcode = 0x50
	OPCODE_BVS_RELATIVE Opcode = 0x70
	//# WIP END
	OPCODE_ADC_IMMEDIATE Opcode = 0x69
	OPCODE_ADC_ZP        Opcode = 0x65
	OPCODE_ADC_ABSOLUTE  Opcode = 0x6D
	OPCODE_SBC_IMMEDIATE Opcode = 0xD9
	//# WIP
	OPCODE_BIT_ABSOLUTE Opcode = 0x24
	OPCODE_BIT_ZP       Opcode = 0x2C
	//# WIP END

	OPCODE_CLC_IMPLIED Opcode = 0x18
	OPCODE_CLD_IMPLIED Opcode = 0xD8
	OPCODE_CLI_IMPLIED Opcode = 0x58
	OPCODE_CLV_IMPLIED Opcode = 0xB8
	OPCODE_SEC_IMPLIED Opcode = 0x38
	OPCODE_SED_IMPLIED Opcode = 0xF8
	OPCODE_SEI_IMPLIED Opcode = 0x78
)

type RegisterDef string

const (
	REGISTER_A  RegisterDef = "A"
	REGISTER_X  RegisterDef = "X"
	REGISTER_Y  RegisterDef = "Y"
	REGISTER_SP RegisterDef = "SP"
)

type Cycle interface {
	Exec(*processor6502) Cycle // returns an additional cycle for variable duration ops, nil otherwise
}

// CycleWait

type CycleWait struct{}

func (c *CycleWait) Exec(proc *processor6502) Cycle {
	return nil
}

// CycleTrash

type CycleTrash struct{}

func (c *CycleTrash) Exec(proc *processor6502) Cycle {
	proc.Bus.Read(proc.PC)
	return nil
}

// CycleFetchImmediate

type CycleFetchImmediate struct {
	Target RegisterDef
}

func (c *CycleFetchImmediate) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(c.Target) = proc.Bus.Read(proc.PC)
	proc.Status.UpdateNZ(*proc.GetRegister(c.Target))
	proc.PC++
	return nil
}

// CycleFetchAddressLow

type CycleFetchAddressLow struct{}

func (c *CycleFetchAddressLow) Exec(proc *processor6502) Cycle {
	proc.AB = uint16(proc.Bus.Read(proc.PC))
	proc.PC++
	return nil
}

// CycleFetchAddressHigh

type CycleFetchAddressHigh struct{}

func (c *CycleFetchAddressHigh) Exec(proc *processor6502) Cycle {
	proc.AB += uint16(proc.Bus.Read(proc.PC)) << 8
	proc.PC++
	return nil
}

// CycleWriteEffective

type CycleWriteEffective struct {
	Source RegisterDef
}

func (c *CycleWriteEffective) Exec(proc *processor6502) Cycle {
	proc.Bus.Write(proc.AB, *proc.GetRegister(c.Source))
	return nil
}

// CycleFetchEffective

type CycleFetchEffective struct {
	Target RegisterDef
}

func (c *CycleFetchEffective) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(c.Target)
	*reg = proc.Bus.Read(proc.AB)
	proc.Status.UpdateNZ(*reg)
	return nil
}

// CycleFetchEffectiveAddrIndexedZp
// this will always stay on zero page

type CycleFetchAddressIndexedZp struct {
	Index RegisterDef
}

func (c *CycleFetchAddressIndexedZp) Exec(proc *processor6502) Cycle {
	proc.AB = (uint16(proc.Bus.Read(proc.AB)) + uint16(*proc.GetRegister(c.Index))) & 0xFF
	return nil
}

// CycleFetchEffectiveAddrIndexed

type CycleFetchAddressIndexed struct {
	Index RegisterDef
}

func (c *CycleFetchAddressIndexed) Exec(proc *processor6502) Cycle {
	high := proc.Bus.Read(proc.PC)
	calcAddressLsb := uint16(proc.AB) + uint16(*proc.GetRegister(c.Index))
	proc.pageChangedInfo = calcAddressLsb&0x100 != 0
	proc.AB = (uint16(high) << 8) | calcAddressLsb&0xFF
	proc.PC++
	return nil
}

// CycleFetchEffectiveAddrIndexed

type CycleFetchEffectiveIndexed struct {
	Target RegisterDef
}

func (c *CycleFetchEffectiveIndexed) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(c.Target) = proc.Bus.Read(proc.AB)
	proc.Status.UpdateNZ(*proc.GetRegister(c.Target))
	if proc.pageChangedInfo {
		proc.pageChangedInfo = false
		proc.AB = ((proc.AB & 0xFF00) + 0x100) | proc.AB&0x00FF
		return c
	}
	return nil
}

// CycleWriteEffectiveAddrIndexed

type CycleCalcEffectiveAddrIndexed struct {
}

func (c *CycleCalcEffectiveAddrIndexed) Exec(proc *processor6502) Cycle {
	proc.Bus.Read(proc.AB)
	if proc.pageChangedInfo {
		proc.pageChangedInfo = false
		proc.AB = ((proc.AB & 0xFF00) + 0x100) | proc.AB&0x00FF
	}
	return nil
}

// CycleIncImplied

type CycleIncImplied struct {
	Implied RegisterDef
}

func (c *CycleIncImplied) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(c.Implied)
	*reg += uint8(1)
	proc.Status.UpdateNZ(*reg)
	return nil
}

// CycleDecImplied

type CycleDecImplied struct {
	Implied RegisterDef
}

func (c *CycleDecImplied) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(c.Implied)
	*reg -= uint8(1)
	proc.Status.UpdateNZ(*reg)
	return nil
}

// CycleCopyRegister

type CycleCopyRegister struct {
	Source RegisterDef
	Target RegisterDef
}

func (c *CycleCopyRegister) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(c.Target) = *proc.GetRegister(c.Source)
	proc.Status.UpdateNZ(*proc.GetRegister(c.Target))
	return nil
}

// CycleAddWithCarryImmediate
// TODO dec mode

type CycleAddWithCarryImmediate struct{}

func (c *CycleAddWithCarryImmediate) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(proc.Bus.Read(proc.PC))
	sum := op1 + op2
	if proc.Status.Get(PROCESSOR_STATUS_FLAG_C) {
		sum += 1
	}
	carry := sum>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit == op2&signBit) && (op1&signBit != sum&signBit)
	*reg = uint8(sum)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	proc.Status.UpdateNZ(*reg)
	proc.PC++
	return nil
}

// CycleAddWithCarry
// TODO dec mode

type CycleAddWithCarryEffective struct{}

func (c *CycleAddWithCarryEffective) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(proc.Bus.Read(proc.AB))
	sum := op1 + op2
	if proc.Status.Get(PROCESSOR_STATUS_FLAG_C) {
		sum += 1
	}
	carry := sum>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit == op2&signBit) && (op1&signBit != sum&signBit)
	*reg = uint8(sum)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	proc.Status.UpdateNZ(*reg)
	return nil
}

type CycleSubWithBorrowImmediate struct{}

func (c *CycleSubWithBorrowImmediate) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(proc.Bus.Read(proc.PC))
	diff := op1 - op2
	if proc.Status.Get(PROCESSOR_STATUS_FLAG_C) {
		diff -= 1
	}
	carry := diff>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit != op2&signBit) && (op2&signBit == diff&signBit)
	*reg = uint8(diff)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	proc.Status.UpdateNZ(*reg)
	proc.PC++
	return nil
}

// CycleProcFlagSet

type CycleProcFlagSet struct {
	Flag ProcessorStatusFlag
}

func (c *CycleProcFlagSet) Exec(proc *processor6502) Cycle {
	proc.Status.Set(c.Flag)
	return nil
}

// CycleProcFlagClear

type CycleProcFlagClear struct {
	Flag ProcessorStatusFlag
}

func (c *CycleProcFlagClear) Exec(proc *processor6502) Cycle {
	proc.Status.Clear(c.Flag)
	return nil
}

// CycleCopyPclFetchPch

type CycleCopyPclFetchPch struct {
}

func (c *CycleCopyPclFetchPch) Exec(proc *processor6502) Cycle {
	proc.PC = (proc.AB & 0xFF) | uint16(proc.Bus.Read(proc.PC))<<8
	return nil
}

// Stackery

type CycleJsrPchPush struct{}

func (c *CycleJsrPchPush) Exec(proc *processor6502) Cycle {
	proc.Bus.Write(0x100+uint16(proc.SP), uint8(proc.PC>>8))
	proc.SP--
	return nil
}

type CycleJsrPclPush struct{}

func (c *CycleJsrPclPush) Exec(proc *processor6502) Cycle {
	proc.Bus.Write(0x100+uint16(proc.SP), uint8(proc.PC))
	proc.SP--
	return nil
}

type CycleRtIncSp struct{}

func (c *CycleRtIncSp) Exec(proc *processor6502) Cycle {
	proc.SP++
	return nil
}

type CycleRtPullPcl struct{}

func (c *CycleRtPullPcl) Exec(proc *processor6502) Cycle {
	proc.PC = proc.PC&0xFF00 | uint16(proc.Bus.Read(0x100+uint16(proc.SP)))
	proc.SP++
	return nil
}

type CycleRtPullPch struct{}

func (c *CycleRtPullPch) Exec(proc *processor6502) Cycle {
	proc.PC = proc.PC&0x00FF | uint16(proc.Bus.Read(0x100+uint16(proc.SP)))<<8
	return nil
}

type CycleRtIncPc struct{}

func (c *CycleRtIncPc) Exec(proc *processor6502) Cycle {
	proc.PC++
	return nil
}

// Branches

// this is not a 100% percent correct, as the check would occur during instruction fetch,
// but that is not part of the intruction abstraction right now so do it here
var tmpCycleBranchTake, tmpCycleBranchTakeUnderflow, tmpCycleBranchTakeOverflow = &CycleBranchTake{PageCrossing: pc.PAGE_NOT_CROSSED}, &CycleBranchTake{PageCrossing: pc.PAGE_CROSSED_UNDERFLOW}, &CycleBranchTake{PageCrossing: pc.PAGE_CROSSED_OVERFLOW}

type CycleBranchTake struct {
	pc.PageCrossing
}

func (c *CycleBranchTake) Exec(proc *processor6502) Cycle {
	if c.PageCrossing == pc.PAGE_CROSSED_OVERFLOW {
		proc.PC = (proc.PC & 0xFF00) + 1 | (proc.PC & 0xFF)
		return tmpCycleBranchTake
	}
	if c.PageCrossing == pc.PAGE_CROSSED_UNDERFLOW {
		proc.PC = (proc.PC & 0xFF00) - uint16(1) | (proc.PC & 0xFF)
		return tmpCycleBranchTake
	}
	proc.PC++
	return nil
}

type CycleBranchFetchOp struct {
	Flag     ProcessorStatusFlag
	FlagTest bool
}

func (c *CycleBranchFetchOp) Exec(proc *processor6502) Cycle {
	branchAddress := proc.Bus.Read(proc.PC)
	if c.FlagTest != proc.Status.Get(c.Flag) {
		proc.PC++
		return nil
	}
	newPC, pageCross := pc.AddPcSigned(proc.PC, branchAddress)
	proc.PC = proc.PC&0xFF00 | newPC&0xFF

	switch pageCross {
	case pc.PAGE_CROSSED_OVERFLOW:
		return tmpCycleBranchTakeOverflow
	case pc.PAGE_CROSSED_UNDERFLOW:
		return tmpCycleBranchTakeUnderflow
	default:
		return tmpCycleBranchTake
	}

}

// InstructionSet
// Instructions are one cycle longer than array elements as fetch op is part of the cpu
var InstructionSet = map[Opcode][]Cycle{
	OPCODE_NOP: {&CycleWait{}},
	// LD*
	OPCODE_LDA_IMMEDIATE:  {&CycleFetchImmediate{REGISTER_A}},
	OPCODE_LDX_IMMEDIATE:  {&CycleFetchImmediate{REGISTER_X}},
	OPCODE_LDY_IMMEDIATE:  {&CycleFetchImmediate{REGISTER_Y}},
	OPCODE_LDA_ZP:         {&CycleFetchAddressLow{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ZP:         {&CycleFetchAddressLow{}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ZP:         {&CycleFetchAddressLow{}, &CycleFetchEffective{REGISTER_Y}},
	OPCODE_LDA_ZP_X:       {&CycleFetchAddressLow{}, &CycleFetchAddressIndexedZp{REGISTER_X}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ZP_Y:       {&CycleFetchAddressLow{}, &CycleFetchAddressIndexedZp{REGISTER_Y}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ZP_X:       {&CycleFetchAddressLow{}, &CycleFetchAddressIndexedZp{REGISTER_X}, &CycleFetchEffective{REGISTER_Y}},
	OPCODE_LDA_ABSOLUTE:   {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ABSOLUTE:   {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ABSOLUTE:   {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleFetchEffective{REGISTER_Y}},
	OPCODE_LDA_ABSOLUTE_X: {&CycleFetchAddressLow{}, &CycleFetchAddressIndexed{REGISTER_X}, &CycleFetchEffectiveIndexed{REGISTER_A}},
	OPCODE_LDA_ABSOLUTE_Y: {&CycleFetchAddressLow{}, &CycleFetchAddressIndexed{REGISTER_Y}, &CycleFetchEffectiveIndexed{REGISTER_A}},
	OPCODE_LDX_ABSOLUTE_Y: {&CycleFetchAddressLow{}, &CycleFetchAddressIndexed{REGISTER_Y}, &CycleFetchEffectiveIndexed{REGISTER_X}},
	OPCODE_LDY_ABSOLUTE_X: {&CycleFetchAddressLow{}, &CycleFetchAddressIndexed{REGISTER_X}, &CycleFetchEffectiveIndexed{REGISTER_Y}},
	// IN*/DE*
	OPCODE_INX_IMPLIED: {&CycleIncImplied{REGISTER_X}},
	OPCODE_INY_IMPLIED: {&CycleIncImplied{REGISTER_Y}},
	OPCODE_DEX_IMPLIED: {&CycleDecImplied{REGISTER_X}},
	OPCODE_DEY_IMPLIED: {&CycleDecImplied{REGISTER_Y}},
	// ST*
	OPCODE_STA_ZP:       {&CycleFetchAddressLow{}, &CycleWriteEffective{REGISTER_A}},
	OPCODE_STX_ZP:       {&CycleFetchAddressLow{}, &CycleWriteEffective{REGISTER_X}},
	OPCODE_STY_ZP:       {&CycleFetchAddressLow{}, &CycleWriteEffective{REGISTER_Y}},
	OPCODE_STA_ZP_X:     {&CycleFetchAddressLow{}, &CycleFetchAddressIndexedZp{REGISTER_X}, &CycleWriteEffective{REGISTER_A}},
	OPCODE_STX_ZP_Y:     {&CycleFetchAddressLow{}, &CycleFetchAddressIndexedZp{REGISTER_Y}, &CycleWriteEffective{REGISTER_X}},
	OPCODE_STY_ZP_X:     {&CycleFetchAddressLow{}, &CycleFetchAddressIndexedZp{REGISTER_X}, &CycleWriteEffective{REGISTER_Y}},
	OPCODE_STA_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleWriteEffective{REGISTER_A}},
	OPCODE_STX_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleWriteEffective{REGISTER_X}},
	OPCODE_STY_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleWriteEffective{REGISTER_Y}},
	//# WIP
	OPCODE_STA_ABSOLUTE_X: {&CycleFetchAddressLow{}, &CycleFetchAddressIndexed{REGISTER_X}, &CycleCalcEffectiveAddrIndexed{}, &CycleWriteEffective{REGISTER_A}},
	OPCODE_STA_ABSOLUTE_Y: {&CycleFetchAddressLow{}, &CycleFetchAddressIndexed{REGISTER_Y}, &CycleCalcEffectiveAddrIndexed{}, &CycleWriteEffective{REGISTER_A}},
	//# WIP END
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
	OPCODE_JMP_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleCopyPclFetchPch{}},
	OPCODE_JSR_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleWait{}, &CycleJsrPchPush{}, &CycleJsrPclPush{}, &CycleCopyPclFetchPch{}},
	OPCODE_RTS_IMPLIED:  {&CycleTrash{}, &CycleRtIncSp{}, &CycleRtPullPcl{}, &CycleRtPullPch{}, &CycleRtIncPc{}},
	// ARITH
	OPCODE_ADC_IMMEDIATE: {&CycleAddWithCarryImmediate{}},
	OPCODE_ADC_ZP:        {&CycleFetchAddressLow{}, &CycleAddWithCarryEffective{}},
	OPCODE_ADC_ABSOLUTE:  {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleAddWithCarryEffective{}},
	OPCODE_SBC_IMMEDIATE: {&CycleSubWithBorrowImmediate{}},
	// STATUS
	OPCODE_CLC_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_C}},
	OPCODE_CLD_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_D}},
	OPCODE_CLI_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_I}},
	OPCODE_CLV_IMPLIED: {&CycleProcFlagClear{PROCESSOR_STATUS_FLAG_V}},
	OPCODE_SEC_IMPLIED: {&CycleProcFlagSet{PROCESSOR_STATUS_FLAG_C}},
	OPCODE_SED_IMPLIED: {&CycleProcFlagSet{PROCESSOR_STATUS_FLAG_D}},
	OPCODE_SEI_IMPLIED: {&CycleProcFlagSet{PROCESSOR_STATUS_FLAG_I}},
}
