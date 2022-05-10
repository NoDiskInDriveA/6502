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

type state int8

const (
	STATE_FETCH   state = iota
	STATE_EXECUTE state = iota
)

type Cpu struct {
	A                        uint8
	X                        uint8
	Y                        uint8
	SP                       uint8
	PC                       uint16
	Status                   *ProcessorStatus
	Bus                      *Bus
	Clock                    Clock
	ir                       Opcode
	currentInstructionCycles []Cycle
	lifetimeCycles           uint64
	Data                     uint8
	AB                       uint16
	enableHaltOpcode         bool
	state                    state
}

func NewCpu() *Cpu {
	return &Cpu{
		Status:           NewProcessorStatus(),
		Bus:              NewBus(),
		Clock:            NewTimedClock(),
		enableHaltOpcode: false,
		state:            STATE_FETCH,
	}
}

func (cpu *Cpu) LoadPrgAt(org uint16, path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not load input file")
	}
	cpu.Bus.Memory.LoadAt(org, bytes)
}

func (cpu *Cpu) EnableHaltOpcode(enable bool) {
	cpu.enableHaltOpcode = enable
}

// power up reset is non time critical
func (cpu *Cpu) Run() {
	cpu.SP = 0xFD
	cpu.PC = uint16(cpu.Bus.Read(RESET_VECTOR_LSB)) + uint16(cpu.Bus.Read(RESET_VECTOR_MSB))<<8
	cpu.Loop(10000)
}

func (cpu *Cpu) Loop(maxCycles uint64) {
	var currentInstructionCycle int
	var nextCycle Cycle
	for ; cpu.lifetimeCycles < maxCycles; cpu.lifetimeCycles++ {
		<-cpu.Clock.Ticker()
		switch cpu.state {
		case STATE_FETCH:
			cpu.ir = Opcode(cpu.Bus.Read(cpu.PC))
			if cpu.ir == OPCODE_HALT {
				if cpu.enableHaltOpcode {
					fmt.Print("Encountered HALT instruction, exiting")
					return
				} else {
					cpu.PC++
					continue
				}
			}
			cpu.state = STATE_EXECUTE
			cpu.currentInstructionCycles = InstructionMap[cpu.ir]
			currentInstructionCycle = 0
			nextCycle = cpu.currentInstructionCycles[currentInstructionCycle]
			cpu.PC++
		case STATE_EXECUTE:
			additionalCycle := nextCycle.Exec(cpu)
			if currentInstructionCycle == len(cpu.currentInstructionCycles) {
				// interrupts are latched on the last logical ("hardcoded") cycle,
				// which is also the only one that can return additional cycles
				cpu.latchInterrupts()
			}
			currentInstructionCycle++
			if additionalCycle != nil {
				nextCycle = additionalCycle
			} else if currentInstructionCycle < len(cpu.currentInstructionCycles) {
				nextCycle = cpu.currentInstructionCycles[currentInstructionCycle]
			} else {
				cpu.state = STATE_FETCH
			}
		}

		cpu.dumpStatus()
		cpu.dumpMemoryLine(0x0000)
		cpu.dumpMemoryLine(0x01F0)
		cpu.dumpMemoryLine(0x0800)
		cpu.dumpMemoryLine(0x0810)
	}
}

func (cpu *Cpu) GetRegister(reg RegisterDef) *uint8 {
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

func (cpu *Cpu) latchInterrupts() {
	// TODO handle interrupts
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
	fmt.Printf(">>> A: 0x%02X X: 0x%02X Y: 0x%02X SR: %s\tCycles: %08d\tPC: 0x%04X\tSP: 0x%02X\tIR: 0x%02X\n", cpu.A, cpu.X, cpu.Y, cpu.Status, cpu.lifetimeCycles, cpu.PC, cpu.SP, cpu.ir)
}
