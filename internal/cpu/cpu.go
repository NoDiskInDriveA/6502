package cpu

import (
	"fmt"
	"strings"
)

const (
	NMI_VECTOR_LSB   uint16 = 0xFFFA
	NMI_VECTOR_MSB   uint16 = 0xFFFB
	RESET_VECTOR_LSB uint16 = 0xFFFC
	RESET_VECTOR_MSB uint16 = 0xFFFD
	IRQ_VECTOR_LSB   uint16 = 0xFFFE
	IRQ_VECTOR_MSB   uint16 = 0xFFFF
)

type Cpu struct {
	A        uint8
	X        uint8
	Y        uint8
	SP       uint8
	PC       uint16
	Status   *ProcessorStatus
	Bus      *Bus
	Clock    Clock
	ir       uint8
	data     uint8
	ab       uint16
	pipeline []CycleDef
	cycles   int64
}

func NewCpu() *Cpu {
	return &Cpu{
		A:      0,
		X:      0,
		Y:      0,
		SP:     0,
		PC:     0,
		Status: NewProcessorStatus(),
		Bus:    NewBus(),
		Clock:  NewTimedClock(),
		ir:     0,
		data:   0,
		ab:     0,
		cycles: 0,
	}
}

// power up reset is non time critical
func (cpu *Cpu) Run() {
	cpu.SP = 0xFD
	cpu.PC = uint16(cpu.Bus.Read(RESET_VECTOR_LSB)) + uint16(cpu.Bus.Read(RESET_VECTOR_MSB))<<8
	cpu.loop(64)
}

func (cpu *Cpu) loop(maxCycles int64) {
	for ; cpu.cycles < maxCycles; cpu.cycles++ {
		if len(cpu.pipeline) < 1 {
			// fetch and "decode" happens on last logical instruction cycle
			cpu.ir = cpu.Bus.Read(cpu.PC)
			cpu.pipeline = CycleDefinitionMap[cpu.ir]
			cpu.PC++
		} else {
			next := cpu.pipeline[0]
			cpu.pipeline = cpu.pipeline[1:]
			switch next.C {
			case CYCLE_FETCH_ADDRESS_LOW:
				cpu.ab = uint16(cpu.Bus.Read(cpu.PC))
				cpu.PC++
			case CYCLE_FETCH_ADDRESS_HIGH:
				cpu.ab += uint16(cpu.Bus.Read(cpu.PC)) << 8
				cpu.PC++
			case CYCLE_FETCH_IMMEDIATE:
				cpu.data = cpu.Bus.Read(cpu.PC)
				cpu.PC++
			case CYCLE_FETCH_PCH_SET_PC:
				cpu.PC = uint16(cpu.Bus.Read(cpu.PC))<<8 + uint16(cpu.data)
			case CYCLE_READ_D_TO_REG:
				reg := cpu.getRegister(next.Reg1)
				*reg = cpu.data
				cpu.Status.UpdateNZ(*reg)
			case CYCLE_WRITE_REG_TO_EFFECTIVE:
				cpu.Bus.Write(cpu.ab, *cpu.getRegister(next.Reg1))
			case CYCLE_ALU_ADC:
				var tmp uint16 = uint16(cpu.A) + uint16(cpu.data)
				cpu.A = uint8(tmp)
				cpu.Status.Update(STATUS_FLAG_C, tmp>>8 != 0)
				cpu.Status.UpdateNZ(cpu.A)
			case CYCLE_WAIT:
				cpu.PC = cpu.PC
			default:
				panic(fmt.Sprintf("Unhandled cycle %d!", next.C))
			}
		}
		cpu.dumpStatus()
		cpu.dumpMemoryLine(0x0000)
		cpu.dumpMemoryLine(0x8000)
	}
}

func (cpu *Cpu) getRegister(reg RegisterDef) *uint8 {
	switch reg {
	case REGISTER_A:
		return &cpu.A
	case REGISTER_X:
		return &cpu.X
	case REGISTER_Y:
		return &cpu.Y
	default:
		panic("Unhandled register!")
	}
}

func (cpu *Cpu) dumpMemoryLine(startAddress uint16) {
	fmt.Printf(
		"0x%04X: "+strings.Repeat(" %02X", 16)+"\n",
		startAddress,
		cpu.Bus.memory.Get(startAddress),
		cpu.Bus.memory.Get(startAddress+1),
		cpu.Bus.memory.Get(startAddress+2),
		cpu.Bus.memory.Get(startAddress+3),
		cpu.Bus.memory.Get(startAddress+4),
		cpu.Bus.memory.Get(startAddress+5),
		cpu.Bus.memory.Get(startAddress+6),
		cpu.Bus.memory.Get(startAddress+7),
		cpu.Bus.memory.Get(startAddress+8),
		cpu.Bus.memory.Get(startAddress+9),
		cpu.Bus.memory.Get(startAddress+10),
		cpu.Bus.memory.Get(startAddress+11),
		cpu.Bus.memory.Get(startAddress+12),
		cpu.Bus.memory.Get(startAddress+13),
		cpu.Bus.memory.Get(startAddress+14),
		cpu.Bus.memory.Get(startAddress+15),
	)
}

func (cpu *Cpu) dumpStatus() {
	fmt.Printf("A: %02X\tX: %02X\tY: %02X\tSR: %s\tCycles: %d\tPC: 0x%04X\tSP: 0x%02X\tIR: %02X\n", cpu.A, cpu.X, cpu.Y, cpu.Status, cpu.cycles, cpu.PC, cpu.SP, cpu.ir)
}

// func (cpu *Cpu) push(value uint8) {
// 	cpu.Bus.Write(0x100+uint16(cpu.SP), value)
// 	cpu.SP--
// }

// func (cpu *Cpu) pop(value uint8) {
// 	cpu.Bus.Write(0x100+uint16(cpu.SP), value)
// 	cpu.SP++
// }
