package mos_6502

import (
	"github.com/NoDiskInDriveA/6502/internal/helper/pc"
)

const (
	STACK_BASE_ADDR = 0x0100
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

type CycleFetchImmediateOperandToRegister struct {
	Target RegisterDef
}

func (c *CycleFetchImmediateOperandToRegister) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(c.Target) = proc.Bus.Read(proc.PC)
	proc.Status.UpdateNZ(*proc.GetRegister(c.Target))
	proc.PC++
	return nil
}

// CycleFetchAddressLow

type CycleFetchAddressOperandLow struct{}

func (c *CycleFetchAddressOperandLow) Exec(proc *processor6502) Cycle {
	proc.AB = uint16(proc.Bus.Read(proc.PC))
	proc.PC++
	return nil
}

// CycleFetchAddressHigh

type CycleFetchAddressOperandHigh struct{}

func (c *CycleFetchAddressOperandHigh) Exec(proc *processor6502) Cycle {
	proc.AB += uint16(proc.Bus.Read(proc.PC)) << 8
	proc.PC++
	return nil
}

// CycleWriteEffective

type CycleWriteRegisterEffective struct {
	Source RegisterDef
}

func (c *CycleWriteRegisterEffective) Exec(proc *processor6502) Cycle {
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

// CycleFetchAddressZpIndexed
// this will always stay on zero page

type CycleFetchAddressZpIndexed struct {
	Index RegisterDef
}

func (c *CycleFetchAddressZpIndexed) Exec(proc *processor6502) Cycle {
	proc.AB = (proc.AB + uint16(*proc.GetRegister(c.Index))) & 0xFF
	return nil
}

// CycleFetchEffectiveAddrIndexed

type CycleFetchAddressOperandHighIndexed struct {
	Index RegisterDef
}

func (c *CycleFetchAddressOperandHighIndexed) Exec(proc *processor6502) Cycle {
	high := proc.Bus.Read(proc.PC)
	calcAddressLsb := uint16(proc.AB) + uint16(*proc.GetRegister(c.Index))
	proc.pageChangedInfo = calcAddressLsb&0x100 != 0
	proc.AB = (uint16(high) << 8) | calcAddressLsb&0xFF
	proc.PC++
	return nil
}

// CycleFetchEffectiveIndexed

type CycleFetchEffectiveWithPageFix struct {
	Target RegisterDef
}

func (c *CycleFetchEffectiveWithPageFix) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(c.Target) = proc.Bus.Read(proc.AB)
	proc.Status.UpdateNZ(*proc.GetRegister(c.Target))
	if proc.pageChangedInfo {
		proc.pageChangedInfo = false
		proc.AB = ((proc.AB & 0xFF00) + 0x100) | proc.AB&0x00FF
		return c
	}
	return nil
}

// CycleCalcEffectiveAddrIndexed

type CycleCalcEffectiveAddrIndexed struct{}

func (c *CycleCalcEffectiveAddrIndexed) Exec(proc *processor6502) Cycle {
	proc.Bus.Read(proc.AB)
	if proc.pageChangedInfo {
		proc.pageChangedInfo = false
		proc.AB = ((proc.AB & 0xFF00) + 0x100) | proc.AB&0x00FF
	}
	return nil
}

type CycleIndexedIndirectHigh struct{}

func (c *CycleIndexedIndirectHigh) Exec(proc *processor6502) Cycle {
	proc.AB = uint16(proc.Bus.Read(proc.AB+1))<<8 | uint16(proc.DataLatch)
	return nil
}

type CycleIndirectIndexedFetchHighByte struct{}

func (c *CycleIndirectIndexedFetchHighByte) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_Y)
	indexedLow := uint16(proc.DataLatch) + uint16(*reg)
	if indexedLow > 0xFF {
		proc.pageChangedInfo = true
		indexedLow &= 0xFF
	}

	proc.AB = uint16(proc.Bus.Read(proc.AB+1))<<8 | uint16(indexedLow)
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
	if proc.Status.GetFlag(PROCESSOR_STATUS_FLAG_C) {
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
	if proc.Status.GetFlag(PROCESSOR_STATUS_FLAG_C) {
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
	if !proc.Status.GetFlag(PROCESSOR_STATUS_FLAG_C) {
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

type CycleCmpEffective struct{}

func (c *CycleCmpEffective) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(proc.Bus.Read(proc.AB))
	diff := op1 - op2
	carry := diff>>8 != 0
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	proc.Status.UpdateNZ(uint8(diff))
	return nil
}

type CycleCmpImmediate struct{}

func (c *CycleCmpImmediate) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(proc.Bus.Read(proc.PC))
	diff := op1 - op2
	carry := diff>>8 != 0
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	proc.Status.UpdateNZ(uint8(diff))
	proc.PC++
	return nil
}

type CycleSubWithBorrowEffective struct{}

func (c *CycleSubWithBorrowEffective) Exec(proc *processor6502) Cycle {
	reg := proc.GetRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(proc.Bus.Read(proc.AB))
	diff := op1 - op2
	if !proc.Status.GetFlag(PROCESSOR_STATUS_FLAG_C) {
		diff -= 1
	}
	carry := diff>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit != op2&signBit) && (op2&signBit == diff&signBit)
	*reg = uint8(diff)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	proc.Status.UpdateNZ(*reg)
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

type CycleIncSp struct{}

func (c *CycleIncSp) Exec(proc *processor6502) Cycle {
	proc.SP++
	return nil
}

type CycleDecSp struct{}

func (c *CycleDecSp) Exec(proc *processor6502) Cycle {
	proc.SP--
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
	if c.FlagTest != proc.Status.GetFlag(c.Flag) {
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

// logic

type CycleUnaryLogicOpApply struct {
	op           UnaryLogicOp
	sourceTarget RegisterDef
}

func (c *CycleUnaryLogicOpApply) Exec(proc *processor6502) Cycle {
	c.op.Apply(proc, c.sourceTarget, c.sourceTarget)
	return nil
}

type CycleBinaryLogicOpApplyImmediate struct {
	op BinaryLogicOp
}

func (c *CycleBinaryLogicOpApplyImmediate) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(DATA_LATCH) = proc.Bus.Read(proc.PC)
	c.op.Apply(proc, DATA_LATCH)
	proc.PC++
	return nil
}

type CycleBinaryLogicOpApplyEffectiveWithPageFix struct {
	op BinaryLogicOp
}

func (c *CycleBinaryLogicOpApplyEffectiveWithPageFix) Exec(proc *processor6502) Cycle {
	*proc.GetRegister(DATA_LATCH) = proc.Bus.Read(proc.AB)
	c.op.Apply(proc, DATA_LATCH)
	if proc.pageChangedInfo {
		proc.pageChangedInfo = false
		proc.AB = ((proc.AB & 0xFF00) + 0x100) | proc.AB&0x00FF
		return c
	}
	return nil
}

type CycleDummyReadInstruction struct{}

func (c *CycleDummyReadInstruction) Exec(proc *processor6502) Cycle {
	proc.Bus.Read(proc.PC)
	return nil
}

type CyclePushAcc struct{}

func (c *CyclePushAcc) Exec(proc *processor6502) Cycle {
	proc.Bus.Write(STACK_BASE_ADDR+uint16(proc.SP), proc.A)
	proc.SP--
	return nil
}

type CyclePushStatus struct{}

func (c *CyclePushStatus) Exec(proc *processor6502) Cycle {
	proc.Bus.Write(STACK_BASE_ADDR+uint16(proc.SP), proc.Status.value|uint8(PROCESSOR_STATUS_FLAG_B)|uint8(PROCESSOR_STATUS_FLAG_U))
	proc.SP--
	return nil
}

type CyclePullAcc struct{}

func (c *CyclePullAcc) Exec(proc *processor6502) Cycle {
	proc.A = proc.Bus.Read(STACK_BASE_ADDR + uint16(proc.SP))
	proc.Status.UpdateNZ(proc.A)
	return nil
}

type CyclePullStatus struct{}

func (c *CyclePullStatus) Exec(proc *processor6502) Cycle {
	status := proc.Bus.Read(STACK_BASE_ADDR + uint16(proc.SP))
	proc.Status.value = status & (proc.Status.value & uint8(PROCESSOR_STATUS_FLAG_B|PROCESSOR_STATUS_FLAG_U))

	return nil
}
