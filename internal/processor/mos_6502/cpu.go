package mos_6502

import (
	"github.com/NoDiskInDriveA/6502/internal/device"
	"github.com/NoDiskInDriveA/6502/internal/processor"
)

const (
	NMI_VECTOR_LSB   uint16 = 0xFFFA
	NMI_VECTOR_MSB   uint16 = 0xFFFB
	RESET_VECTOR_LSB uint16 = 0xFFFC
	RESET_VECTOR_MSB uint16 = 0xFFFD
	IRQ_VECTOR_LSB   uint16 = 0xFFFE
	IRQ_VECTOR_MSB   uint16 = 0xFFFF
)

type state int8

const (
	STATE_FETCH   state = iota
	STATE_EXECUTE state = iota
	STATE_HALT    state = iota
)

type Processor interface {
	processor.Processor
	EnableHaltOpcode(enabled bool)
}

type processor6502 struct {
	A                            uint8
	X                            uint8
	Y                            uint8
	SP                           uint8
	PC                           uint16
	Status                       *processorStatus
	Bus                          device.ReadWriteDevice
	DataLatch                    uint8
	AB                           uint16
	IR                           Opcode
	currentInstructionCycles     []Cycle
	currentInstructionCycleIndex int
	nextInstructionCycle         Cycle
	lifetimeCycles               uint64
	enableHaltOpcode             bool
	state                        state
	pageChangedInfo              bool
}

func NewProcessor(bus device.ReadWriteDevice) Processor {
	return &processor6502{Bus: bus, enableHaltOpcode: false}
}

func (p *processor6502) EnableHaltOpcode(enabled bool) {
	p.enableHaltOpcode = enabled
}

// power up reset is non time critical
func (p *processor6502) Reset() {
	p.Status = NewProcessorStatus()
	p.state = STATE_FETCH
	p.pageChangedInfo = false
	p.SP = 0xFD
	p.PC = uint16(p.Bus.Read(RESET_VECTOR_LSB)) + uint16(p.Bus.Read(RESET_VECTOR_MSB))<<8
	p.lifetimeCycles = 7
	p.currentInstructionCycleIndex = 0
	p.nextInstructionCycle = nil
}

func (p *processor6502) Tick() bool {
	switch p.state {
	case STATE_FETCH:
		p.IR = Opcode(p.Bus.Read(p.PC))
		if p.IR == OPCODE_HALT {
			if p.enableHaltOpcode {
				p.state = STATE_HALT
				return true
			} else {
				p.PC++
				return true
			}
		}
		p.state = STATE_EXECUTE
		p.currentInstructionCycles = InstructionSet[p.IR]
		p.currentInstructionCycleIndex = 0
		p.nextInstructionCycle = p.currentInstructionCycles[p.currentInstructionCycleIndex]
		p.PC++
	case STATE_EXECUTE:
		additionalCycle := p.nextInstructionCycle.Exec(p)
		if p.currentInstructionCycleIndex == len(p.currentInstructionCycles) {
			// interrupts are latched on the last logical ("hardcoded") cycle,
			// which is also the only one that can return additional cycles
			// TODO:
			// - need to check whether this info is correct - might be "penultimate" cycle
			// - also check whether status (like in SEI/CLI) is latched before or after instruction
			p.latchInterrupts()
		}
		p.currentInstructionCycleIndex++
		if additionalCycle != nil {
			p.nextInstructionCycle = additionalCycle
		} else if p.currentInstructionCycleIndex < len(p.currentInstructionCycles) {
			p.nextInstructionCycle = p.currentInstructionCycles[p.currentInstructionCycleIndex]
		} else {
			p.state = STATE_FETCH
		}
	case STATE_HALT:
		return false
	}

	p.lifetimeCycles++

	return true
}

func (p *processor6502) GetRegister(reg RegisterDef) *uint8 {
	switch reg {
	case REGISTER_A:
		return &p.A
	case REGISTER_X:
		return &p.X
	case REGISTER_Y:
		return &p.Y
	case REGISTER_SP:
		return &p.SP
	case DATA_LATCH:
		return &p.DataLatch
	default:
		panic("Unhandled register!")
	}
}

func (p *processor6502) latchInterrupts() {
	// TODO handle interrupts
}
