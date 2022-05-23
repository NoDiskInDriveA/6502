package arch

import (
	"fmt"
	"time"

	"github.com/NoDiskInDriveA/6502/internal/device"
	"github.com/NoDiskInDriveA/6502/internal/helper"
	"github.com/NoDiskInDriveA/6502/internal/processor/mos_6502"
)

const (
	PRG_LOCATION uint16 = 0x0800
	ROM_LOCATION uint16 = 0xE000
)

type fantasyArchitecture struct {
	clock device.SystemClock
}

func (fa *fantasyArchitecture) Run() {
	start := time.Now()
	<-fa.clock.StartClock(time.Microsecond)
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s", elapsed)
}

func NewDebugFantasyArchitecture() Architecture {
	bus := device.NewSystemBus()
	memory := device.NewRamDevice()
	helper.LoadIntoMemory(memory, PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/sum.bin")
	helper.LoadIntoMemory(memory, ROM_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/rom.bin")
	bus.AttachMappedDevice(0x0000, 0xFFFF, 0xFFFF, memory)
	p := mos_6502.NewProcessor(bus)
	p.EnableHaltOpcode(true)
	p.Reset()

	clock := device.NewSynchronizedClock()
	clock.AttachDevice(p, 1)
	clock.AttachDebugger(mos_6502.NewFantasyArchDebugger(p, bus))

	return &fantasyArchitecture{clock}
}

func NewFantasyArchitecture() Architecture {
	bus := device.NewSystemBus()
	memory := device.NewRamDevice()
	helper.LoadIntoMemory(memory, PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/sum.bin")
	helper.LoadIntoMemory(memory, ROM_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/rom.bin")
	bus.AttachMappedDevice(0x0000, 0xFFFF, 0xFFFF, memory)
	p := mos_6502.NewProcessor(bus)
	p.EnableHaltOpcode(true)
	p.Reset()

	clock := device.NewSynchronizedClock()
	clock.AttachDevice(p, 1)

	return &fantasyArchitecture{clock}
}
