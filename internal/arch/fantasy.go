package arch

import (
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
	clock   device.SystemClock
	monitor *mos_6502.Monitor
}

func (fa *fantasyArchitecture) Run() {
	<-fa.clock.RunClock(time.Microsecond)
}

func (fa *fantasyArchitecture) Step() {
	fa.clock.TickClock()
}

func (fa *fantasyArchitecture) GetMonitor() *mos_6502.Monitor {
	return fa.monitor
}

func NewDebugFantasyArchitecture() Architecture {
	bus := device.NewSystemBus()
	memory := device.NewRamDevice()
	helper.LoadIntoMemory(memory, PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/tests/indirect_indexed_y.bin")
	helper.LoadIntoMemory(memory, ROM_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/rom.bin")
	bus.AttachMappedDevice(0x0000, 0xFFFF, 0xFFFF, memory)
	p := mos_6502.NewProcessor(bus)
	p.EnableHaltOpcode(true)
	p.Reset()

	clock := device.NewSynchronizedClock()
	clock.AttachDevice(p, 1)
	clock.AttachDebugger(mos_6502.NewFantasyArchDebugger(p, bus))

	return &fantasyArchitecture{clock: clock}
}

func NewFantasyArchitecture() Architecture {
	bus := device.NewSystemBus()
	memory := device.NewRamDevice()
	helper.LoadIntoMemory(memory, PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/loop.bin")
	helper.LoadIntoMemory(memory, ROM_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/rom.bin")
	bus.AttachMappedDevice(0x0000, 0xFFFF, 0xFFFF, memory)
	p := mos_6502.NewProcessor(bus)
	p.EnableHaltOpcode(true)
	p.Reset()

	clock := device.NewSynchronizedClock()
	clock.AttachDevice(p, 1)

	return &fantasyArchitecture{clock: clock}
}

func NewMonitoredFantasyArchitecture() MonitoredArchitecture {
	bus := device.NewSystemBus()
	memory := device.NewRamDevice()
	helper.LoadIntoMemory(memory, PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/sum.bin")
	// helper.LoadIntoMemory(memory, PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/fizzbuzz.bin")
	helper.LoadIntoMemory(memory, ROM_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/rom.bin")
	bus.AttachMappedDevice(0x0000, 0xFFFF, 0xFFFF, memory)
	p := mos_6502.NewProcessor(bus)
	p.EnableHaltOpcode(true)
	p.Reset()

	monitor := mos_6502.NewMonitor(p, bus)
	monitor.AddAddressRange(0x0000)
	monitor.AddAddressRange(0x0010)
	monitor.AddAddressRange(0x0020)
	monitor.AddAddressRange(0x0030)
	monitor.AddAddressRange(0x0040)
	monitor.AddAddressRange(0x0050)
	monitor.AddAddressRange(0x0060)
	monitor.AddAddressRange(0x01E0)
	monitor.AddAddressRange(0x01F0)
	monitor.AddAddressRange(0x0800)
	monitor.AddAddressRange(0x0810)
	monitor.AddAddressRange(0x0820)
	monitor.AddAddressRange(0x0830)
	monitor.AddAddressRange(0x0840)
	monitor.AddAddressRange(0x0850)
	monitor.AddAddressRange(0x2000)
	monitor.AddAddressRange(0x2010)
	clock := device.NewSynchronizedClock()
	clock.AttachDevice(p, 1)

	return &fantasyArchitecture{clock, monitor}
}
