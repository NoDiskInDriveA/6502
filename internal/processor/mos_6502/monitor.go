package mos_6502

import (
	"fmt"
	"strings"

	"github.com/NoDiskInDriveA/6502/internal/device"
)

type Monitor struct {
	processor     Processor
	bus           device.ReadableDevice
	addressRanges []uint16
}

func NewMonitor(processor Processor, bus device.ReadableDevice) *Monitor {
	return &Monitor{processor: processor, bus: bus}
}

func (m *Monitor) AddAddressRange(startAddress uint16) {
	m.addressRanges = append(m.addressRanges, startAddress)
}

func (m *Monitor) GetMemoryView() string {
	var sb strings.Builder
	sb.WriteString("          00 01 02 03   04 05 06 07   08 09 0A 0B   0C 0D 0E 0F\n")
	sb.WriteString("          -----------------------------------------------------\n")
	for _, address := range m.addressRanges {
		sb.WriteString(fmt.Sprintf(" 0x%04X:", address))
		for i := uint16(0); i < 16; i++ {
			if i%4 == 0 {
				sb.WriteString("  ")
			}
			sb.WriteString(fmt.Sprintf("%02X ", m.bus.Read(address+i)))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (m *Monitor) GetStatusWord() string {
	p, _ := m.processor.(*processor6502)

	return p.Status.String()
}

func (m *Monitor) GetInternalStatus() []string {
	p, _ := m.processor.(*processor6502)

	state := func() string {
		switch p.state {
		case STATE_EXECUTE:
			return "execute"
		case STATE_FETCH:
			return "fetch"
		case STATE_HALT:
			return "halt"
		}
		return "unknown"
	}()
	return []string{
		fmt.Sprintf("A:   0x%02X", p.A),
		fmt.Sprintf("X:   0x%02X", p.X),
		fmt.Sprintf("Y:   0x%02X", p.Y),
		fmt.Sprintf("SP:  0x%02X", p.SP),
		fmt.Sprintf("PC:  0x%04X", p.PC),
		fmt.Sprintf("AB:  0x%04X", p.AB),
		fmt.Sprintf("IR:  0x%02X", p.IR),
		fmt.Sprintf("Cyc: %d", p.lifetimeCycles),
		fmt.Sprintf("STT: %s", state),
	}
}
