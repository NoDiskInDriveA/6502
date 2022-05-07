package cpu

type Opcode uint8

const (
	OPCODE_HALT Opcode = 0xFF // artificial
	OPCODE_NOP  Opcode = 0xEA

	OPCODE_LDA_IMMEDIATE Opcode = 0xA9
	OPCODE_LDX_IMMEDIATE Opcode = 0xA2
	OPCODE_LDY_IMMEDIATE Opcode = 0xA0
	OPCODE_LDA_ZP        Opcode = 0xA5
	OPCODE_LDX_ZP        Opcode = 0xA6
	OPCODE_LDY_ZP        Opcode = 0xA4
	OPCODE_LDA_ABSOLUTE  Opcode = 0xAD
	OPCODE_LDX_ABSOLUTE  Opcode = 0xAE
	OPCODE_LDY_ABSOLUTE  Opcode = 0xAC

	OPCODE_STA_ZP       Opcode = 0x85
	OPCODE_STX_ZP       Opcode = 0x86
	OPCODE_STY_ZP       Opcode = 0x84
	OPCODE_STA_ABSOLUTE Opcode = 0x8D
	OPCODE_STX_ABSOLUTE Opcode = 0x8E
	OPCODE_STY_ABSOLUTE Opcode = 0x8C

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

	OPCODE_ADC_IMMEDIATE Opcode = 0x69
	OPCODE_ADC_ZP        Opcode = 0x65
	OPCODE_ADC_ABSOLUTE  Opcode = 0x6D
	OPCODE_SBC_IMMEDIATE Opcode = 0xD9

	OPCODE_CLC_IMPLIED Opcode = 0x18
	OPCODE_CLD_IMPLIED Opcode = 0xD8
	OPCODE_CLI_IMPLIED Opcode = 0x58
	OPCODE_CLV_IMPLIED Opcode = 0xB8
	OPCODE_SEC_IMPLIED Opcode = 0x38
	OPCODE_SED_IMPLIED Opcode = 0xF8
	OPCODE_SEI_IMPLIED Opcode = 0x78
)

type InstructionType uint8

const (
	INSTRUCTION_TYPE_JMP       InstructionType = iota
	INSTRUCTION_TYPE_READ      InstructionType = iota
	INSTRUCTION_TYPE_WRITE     InstructionType = iota
	INSTRUCTION_TYPE_READWRITE InstructionType = iota
	INSTRUCTION_TYPE_ANY       InstructionType = iota
)

type InstructionMode uint8

const (
	INSTRUCTION_MODE_IMMEDIATE InstructionMode = iota
	INSTRUCTION_MODE_IMPLIED   InstructionMode = iota
	INSTRUCTION_MODE_ABSOLUTE  InstructionMode = iota
	INSTRUCTION_MODE_ZP        InstructionMode = iota
	INSTRUCTION_MODE_RELATIVE  InstructionMode = iota
)

type InstructionParams struct {
	Source   RegisterDef
	Target   RegisterDef
	Index    RegisterDef
	Flag     ProcessorStatusFlag
	FlagTrue bool
}

type RegisterDef string

const (
	REGISTER_A  RegisterDef = "A"
	REGISTER_X  RegisterDef = "X"
	REGISTER_Y  RegisterDef = "Y"
	REGISTER_SP RegisterDef = "SP"
)

type Cycle interface {
	Exec(*Cpu)
}

// CycleWait

type CycleWait struct{}

func (c *CycleWait) Exec(cpu *Cpu) {
	// do nothing
}

// CycleTrash

type CycleTrash struct{}

func (c *CycleTrash) Exec(cpu *Cpu) {
	cpu.Bus.Read(cpu.PC)
}

// CycleFetchImmediate

type CycleFetchImmediate struct {
	Target RegisterDef
}

func (c *CycleFetchImmediate) Exec(cpu *Cpu) {
	*cpu.getRegister(c.Target) = cpu.Bus.Read(cpu.PC)
	cpu.Status.UpdateNZ(*cpu.getRegister(c.Target))
	cpu.PC++
}

// CycleFetchAddressLow

type CycleFetchAddressLow struct{}

func (c *CycleFetchAddressLow) Exec(cpu *Cpu) {
	cpu.AB = uint16(cpu.Bus.Read(cpu.PC))
	cpu.PC++
}

// CycleFetchAddressHigh

type CycleFetchAddressHigh struct{}

func (c *CycleFetchAddressHigh) Exec(cpu *Cpu) {
	cpu.AB += uint16(cpu.Bus.Read(cpu.PC)) << 8
	cpu.PC++
}

// CycleWriteEffective

type CycleWriteEffective struct {
	Source RegisterDef
}

func (c *CycleWriteEffective) Exec(cpu *Cpu) {
	cpu.Bus.Write(cpu.AB, *cpu.getRegister(c.Source))
}

// CycleFetchEffective

type CycleFetchEffective struct {
	Target RegisterDef
}

func (c *CycleFetchEffective) Exec(cpu *Cpu) {
	reg := cpu.getRegister(c.Target)
	*reg = cpu.Bus.Read(cpu.AB)
	cpu.Status.UpdateNZ(*reg)
}

// CycleIncImplied

type CycleIncImplied struct {
	Implied RegisterDef
}

func (c *CycleIncImplied) Exec(cpu *Cpu) {
	reg := cpu.getRegister(c.Implied)
	*reg += 1
	cpu.Status.UpdateNZ(*reg)
}

// CycleDecImplied

type CycleDecImplied struct {
	Implied RegisterDef
}

func (c *CycleDecImplied) Exec(cpu *Cpu) {
	reg := cpu.getRegister(c.Implied)
	*reg -= 1
	cpu.Status.UpdateNZ(*reg)
}

// CycleCopyRegister

type CycleCopyRegister struct {
	Source RegisterDef
	Target RegisterDef
}

func (c *CycleCopyRegister) Exec(cpu *Cpu) {
	*cpu.getRegister(c.Target) = *cpu.getRegister(c.Source)
}

// CycleAddWithCarryImmediate
// TODO dec mode

type CycleAddWithCarryImmediate struct{}

func (c *CycleAddWithCarryImmediate) Exec(cpu *Cpu) {
	reg := cpu.getRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(cpu.Bus.Read(cpu.PC))
	sum := op1 + op2
	if cpu.Status.Get(PROCESSOR_STATUS_FLAG_C) {
		sum += 1
	}
	carry := sum>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit == op2&signBit) && (op1&signBit != sum&signBit)
	*reg = uint8(sum)
	cpu.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	cpu.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	cpu.Status.UpdateNZ(*reg)
	cpu.PC++
}

// CycleAddWithCarry
// TODO dec mode

type CycleAddWithCarryEffective struct{}

func (c *CycleAddWithCarryEffective) Exec(cpu *Cpu) {
	reg := cpu.getRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(cpu.Bus.Read(cpu.AB))
	sum := op1 + op2
	if cpu.Status.Get(PROCESSOR_STATUS_FLAG_C) {
		sum += 1
	}
	carry := sum>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit == op2&signBit) && (op1&signBit != sum&signBit)
	*reg = uint8(sum)
	cpu.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	cpu.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	cpu.Status.UpdateNZ(*reg)
}

type CycleSubWithBorrowImmediate struct{}

func (c *CycleSubWithBorrowImmediate) Exec(cpu *Cpu) {
	reg := cpu.getRegister(REGISTER_A)
	op1 := uint16(*reg)
	op2 := uint16(cpu.Bus.Read(cpu.PC))
	diff := op1 - op2
	if cpu.Status.Get(PROCESSOR_STATUS_FLAG_C) {
		diff -= 1
	}
	carry := diff>>8 != 0
	signBit := uint16(0x0080)
	overflow := (op1&signBit != op2&signBit) && (op2&signBit == diff&signBit)
	*reg = uint8(diff)
	cpu.Status.Update(PROCESSOR_STATUS_FLAG_C, carry)
	cpu.Status.Update(PROCESSOR_STATUS_FLAG_V, overflow)
	cpu.Status.UpdateNZ(*reg)
	cpu.PC++
}

// CycleProcFlagSet

type CycleProcFlagSet struct {
	Flag ProcessorStatusFlag
}

func (c *CycleProcFlagSet) Exec(cpu *Cpu) {
	cpu.Status.Set(c.Flag)
}

// CycleProcFlagClear

type CycleProcFlagClear struct {
	Flag ProcessorStatusFlag
}

func (c *CycleProcFlagClear) Exec(cpu *Cpu) {
	cpu.Status.Clear(c.Flag)
}

// CycleCopyPclFetchPch

type CycleCopyPclFetchPch struct {
}

func (c *CycleCopyPclFetchPch) Exec(cpu *Cpu) {
	cpu.PC = (cpu.AB & 0xFF) | uint16(cpu.Bus.Read(cpu.PC))<<8
}

// Stackery

type CycleJsrPchPush struct{}

func (c *CycleJsrPchPush) Exec(cpu *Cpu) {
	cpu.Bus.Write(0x100+uint16(cpu.SP), uint8(cpu.PC>>8))
	cpu.SP--
}

type CycleJsrPclPush struct{}

func (c *CycleJsrPclPush) Exec(cpu *Cpu) {
	cpu.Bus.Write(0x100+uint16(cpu.SP), uint8(cpu.PC))
	cpu.SP--
}

type CycleRtIncSp struct{}

func (c *CycleRtIncSp) Exec(cpu *Cpu) {
	cpu.SP++
}

type CycleRtPullPcl struct{}

func (c *CycleRtPullPcl) Exec(cpu *Cpu) {
	cpu.PC = cpu.PC&0xFF00 | uint16(cpu.Bus.Read(0x100+uint16(cpu.SP)))
	cpu.SP++
}

type CycleRtPullPch struct{}

func (c *CycleRtPullPch) Exec(cpu *Cpu) {
	cpu.PC = cpu.PC&0x00FF | uint16(cpu.Bus.Read(0x100+uint16(cpu.SP)))<<8
}

type CycleRtIncPc struct{}

func (c *CycleRtIncPc) Exec(cpu *Cpu) {
	cpu.PC++
}

// InstructionMap

var InstructionMap = map[Opcode][]Cycle{
	OPCODE_NOP: {&CycleWait{}},
	// LD*
	OPCODE_LDA_IMMEDIATE: {&CycleFetchImmediate{REGISTER_A}},
	OPCODE_LDX_IMMEDIATE: {&CycleFetchImmediate{REGISTER_X}},
	OPCODE_LDY_IMMEDIATE: {&CycleFetchImmediate{REGISTER_Y}},
	OPCODE_LDA_ZP:        {&CycleFetchAddressLow{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ZP:        {&CycleFetchAddressLow{}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ZP:        {&CycleFetchAddressLow{}, &CycleFetchEffective{REGISTER_Y}},
	OPCODE_LDA_ABSOLUTE:  {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleFetchEffective{REGISTER_A}},
	OPCODE_LDX_ABSOLUTE:  {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleFetchEffective{REGISTER_X}},
	OPCODE_LDY_ABSOLUTE:  {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleFetchEffective{REGISTER_Y}},
	// IN*/DE*
	OPCODE_INX_IMPLIED: {&CycleIncImplied{REGISTER_X}},
	OPCODE_INY_IMPLIED: {&CycleIncImplied{REGISTER_Y}},
	OPCODE_DEX_IMPLIED: {&CycleDecImplied{REGISTER_X}},
	OPCODE_DEY_IMPLIED: {&CycleDecImplied{REGISTER_Y}},
	// ST*
	OPCODE_STA_ZP:       {&CycleFetchAddressLow{}, &CycleWriteEffective{REGISTER_A}},
	OPCODE_STX_ZP:       {&CycleFetchAddressLow{}, &CycleWriteEffective{REGISTER_X}},
	OPCODE_STY_ZP:       {&CycleFetchAddressLow{}, &CycleWriteEffective{REGISTER_Y}},
	OPCODE_STA_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleWriteEffective{REGISTER_A}},
	OPCODE_STX_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleWriteEffective{REGISTER_X}},
	OPCODE_STY_ABSOLUTE: {&CycleFetchAddressLow{}, &CycleFetchAddressHigh{}, &CycleWriteEffective{REGISTER_Y}},
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
