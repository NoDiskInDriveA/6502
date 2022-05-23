package mos_6502

import (
	"fmt"
	"strings"

	"github.com/NoDiskInDriveA/6502/internal/device"
)

type debugger struct {
	processor Processor
	bus       device.AddressDecodingBus
}

func NewFantasyArchDebugger(processor Processor, bus device.AddressDecodingBus) device.ClockAwareDevice {
	return &debugger{processor, bus}
}

func (d debugger) Tick() bool {
	d.dumpStatus()
	d.dumpMemoryLine(0x0000)
	d.dumpMemoryLine(0x01F0)
	d.dumpMemoryLine(0x0800)
	d.dumpMemoryLine(0x0810)
	d.dumpMemoryLine(0x1EF0)
	d.dumpMemoryLine(0x1FF0)
	d.dumpMemoryLine(0x2000)
	return true
}

func (d debugger) dumpStatus() {
	p, _ := d.processor.(*processor6502)
	fmt.Printf(">>> A: 0x%02X X: 0x%02X Y: 0x%02X SR: %s\tCycles: %08d\tPC: 0x%04X\tSP: 0x%02X\tIR: 0x%02X\tAB: 0x%04X\n", p.A, p.X, p.Y, p.Status, p.lifetimeCycles, p.PC, p.SP, p.IR, p.AB)
}

func (d debugger) dumpMemoryLine(startAddress uint16) {
	fmt.Printf(
		"0x%04X: "+strings.Repeat(" %02X", 8)+" |"+strings.Repeat(" %02X", 8)+"\n",
		startAddress,
		d.bus.Read(startAddress),
		d.bus.Read(startAddress+1),
		d.bus.Read(startAddress+2),
		d.bus.Read(startAddress+3),
		d.bus.Read(startAddress+4),
		d.bus.Read(startAddress+5),
		d.bus.Read(startAddress+6),
		d.bus.Read(startAddress+7),
		d.bus.Read(startAddress+8),
		d.bus.Read(startAddress+9),
		d.bus.Read(startAddress+10),
		d.bus.Read(startAddress+11),
		d.bus.Read(startAddress+12),
		d.bus.Read(startAddress+13),
		d.bus.Read(startAddress+14),
		d.bus.Read(startAddress+15),
	)
}
