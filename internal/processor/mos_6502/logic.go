package mos_6502

type UnaryLogicOp interface {
	Apply(proc *processor6502, source RegisterDef, target RegisterDef)
}

type BinaryLogicOp interface {
	Apply(proc *processor6502, register RegisterDef)
}

type logicAsl struct{}

func (l *logicAsl) Apply(proc *processor6502, source RegisterDef, target RegisterDef) {
	result := uint16(*proc.GetRegister(source)) << 1
	*proc.GetRegister(target) = uint8(result)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, result>>8 != 0)
	proc.Status.UpdateNZ(uint8(result))
}

type logicLsr struct{}

func (l *logicLsr) Apply(proc *processor6502, source RegisterDef, target RegisterDef) {
	result := uint16(*proc.GetRegister(source))
	nextCarry := (result & 1) == 1
	result = result >> 1
	*proc.GetRegister(target) = uint8(result)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, nextCarry)
	proc.Status.UpdateNZ(uint8(result))
}

type logicRol struct{}

func (l *logicRol) Apply(proc *processor6502, source RegisterDef, target RegisterDef) {
	result := (uint16(*proc.GetRegister(source)) << 1) | uint16(proc.Status.Get(PROCESSOR_STATUS_FLAG_C))
	*proc.GetRegister(target) = uint8(result)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, result>>8 != 0)
	proc.Status.UpdateNZ(uint8(result))
}

type logicRor struct{}

func (l *logicRor) Apply(proc *processor6502, source RegisterDef, target RegisterDef) {
	result := uint16(*proc.GetRegister(source))
	nextCarry := (result & 1) == 1
	result = (result >> 1) | (uint16(proc.Status.Get(PROCESSOR_STATUS_FLAG_C)) << 7)
	*proc.GetRegister(target) = uint8(result)
	proc.Status.Update(PROCESSOR_STATUS_FLAG_C, nextCarry)
	proc.Status.UpdateNZ(uint8(result))
}

// binary

type logicAnd struct{}

func (l *logicAnd) Apply(proc *processor6502, register RegisterDef) {
	regA := proc.GetRegister(REGISTER_A)
	op := proc.GetRegister(register)
	*regA = *regA & *op
	proc.Status.UpdateNZ(uint8(*regA))
}

type logicOr struct{}

func (l *logicOr) Apply(proc *processor6502, register RegisterDef) {
	regA := proc.GetRegister(REGISTER_A)
	op := proc.GetRegister(register)
	*regA = *regA | *op
	proc.Status.UpdateNZ(uint8(*regA))
}

type logicXor struct{}

func (l *logicXor) Apply(proc *processor6502, register RegisterDef) {
	regA := proc.GetRegister(REGISTER_A)
	op := proc.GetRegister(register)
	*regA = *regA ^ *op
	proc.Status.UpdateNZ(uint8(*regA))
}

type logicInc struct{}

func (l *logicInc) Apply(proc *processor6502, register RegisterDef) {
	*proc.GetRegister(DATA_LATCH) = *proc.GetRegister(DATA_LATCH) + 1
	proc.Status.UpdateNZ(*proc.GetRegister(DATA_LATCH))
}

type logicDec struct{}

func (l *logicDec) Apply(proc *processor6502, register RegisterDef) {
	*proc.GetRegister(DATA_LATCH) = *proc.GetRegister(DATA_LATCH) - 1
	proc.Status.UpdateNZ(*proc.GetRegister(DATA_LATCH))
}

type logicOps struct {
	LogicOpAsl UnaryLogicOp
	LogicOpLsr UnaryLogicOp
	LogicOpRol UnaryLogicOp
	LogicOpRor UnaryLogicOp
	LogicOpAnd BinaryLogicOp
	LogicOpOr  BinaryLogicOp
	LogicOpXor BinaryLogicOp
	LogicOpInc BinaryLogicOp
	LogicOpDec BinaryLogicOp
}

var LogicOps = logicOps{
	LogicOpAsl: &logicAsl{},
	LogicOpLsr: &logicLsr{},
	LogicOpRol: &logicRol{},
	LogicOpRor: &logicRor{},
	LogicOpAnd: &logicAnd{},
	LogicOpOr:  &logicOr{},
	LogicOpXor: &logicXor{},
	LogicOpInc: &logicInc{},
	LogicOpDec: &logicDec{},
}
