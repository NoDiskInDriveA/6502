package cpu

import (
	"fmt"
	"io/ioutil"
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
	ir       Opcode
	opCycles []Cycle
	Data     uint8
	AB       uint16
	cycles   int64
}

func NewCpu() *Cpu {
	return &Cpu{
		Status: NewProcessorStatus(),
		Bus:    NewBus(),
		Clock:  NewTimedClock(),
	}
}

func (cpu *Cpu) LoadPrg(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not load input file")
	}
	cpu.Bus.Memory.Set(RESET_VECTOR_LSB, bytes[0])
	cpu.Bus.Memory.Set(RESET_VECTOR_MSB, bytes[1])
	var org uint16 = uint16(bytes[0]) + uint16(bytes[1])<<8
	cpu.Bus.Memory.LoadAt(org, bytes[2:])
}

// power up reset is non time critical
func (cpu *Cpu) Run() {
	cpu.SP = 0xFD
	cpu.PC = uint16(cpu.Bus.Read(RESET_VECTOR_LSB)) + uint16(cpu.Bus.Read(RESET_VECTOR_MSB))<<8
	cpu.loop(64)
}

func (cpu *Cpu) loop(maxCycles int64) {
	for ; cpu.cycles < maxCycles; cpu.cycles++ {
		<-cpu.Clock.Ticker()
		if len(cpu.opCycles) < 1 {
			// fetch and "decode" happens on last logical instruction cycle
			cpu.ir = Opcode(cpu.Bus.Read(cpu.PC))
			cpu.opCycles = InstructionMap[cpu.ir]
			cpu.PC++
		} else {
			next := cpu.opCycles[0]
			cpu.opCycles = cpu.opCycles[1:]
			next.Exec(cpu)
		}
		cpu.dumpStatus()
		cpu.dumpMemoryLine(0x0000)
		cpu.dumpMemoryLine(0x01F0)
		cpu.dumpMemoryLine(0x4000)
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
	case REGISTER_SP:
		return &cpu.SP
	default:
		panic("Unhandled register!")
	}
}

func (cpu *Cpu) dumpMemoryLine(startAddress uint16) {
	fmt.Printf(
		"0x%04X: "+strings.Repeat(" %02X", 8)+" |"+strings.Repeat(" %02X", 8)+"\n",
		startAddress,
		cpu.Bus.Memory.Get(startAddress),
		cpu.Bus.Memory.Get(startAddress+1),
		cpu.Bus.Memory.Get(startAddress+2),
		cpu.Bus.Memory.Get(startAddress+3),
		cpu.Bus.Memory.Get(startAddress+4),
		cpu.Bus.Memory.Get(startAddress+5),
		cpu.Bus.Memory.Get(startAddress+6),
		cpu.Bus.Memory.Get(startAddress+7),
		cpu.Bus.Memory.Get(startAddress+8),
		cpu.Bus.Memory.Get(startAddress+9),
		cpu.Bus.Memory.Get(startAddress+10),
		cpu.Bus.Memory.Get(startAddress+11),
		cpu.Bus.Memory.Get(startAddress+12),
		cpu.Bus.Memory.Get(startAddress+13),
		cpu.Bus.Memory.Get(startAddress+14),
		cpu.Bus.Memory.Get(startAddress+15),
	)
}

func (cpu *Cpu) dumpStatus() {
	fmt.Printf(">>> A: 0x%02X X: 0x%02X Y: 0x%02X SR: %s\tCycles: %08d\tPC: 0x%04X\tSP: 0x%02X\tIR: 0x%02X\n", cpu.A, cpu.X, cpu.Y, cpu.Status, cpu.cycles, cpu.PC, cpu.SP, cpu.ir)
}
