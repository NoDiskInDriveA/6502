package cpu

import (
	"strings"
)

type ProgramCounter uint16

func (p *ProgramCounter) SetLow(v uint8) {
	*p = ProgramCounter(uint16(*p&0xFF00) | uint16(v))
}

func (p *ProgramCounter) Low() uint8 {
	return uint8(*p)
}

func (p *ProgramCounter) SetHigh(v uint8) {
	*p = ProgramCounter(uint16(*p&0x00FF) | uint16(v)<<8)
}

func (p *ProgramCounter) High() uint8 {
	return uint8(*p >> 8)
}

type ProcessorStatusFlag uint8

const (
	PROCESSOR_STATUS_FLAG_N ProcessorStatusFlag = 0x80
	PROCESSOR_STATUS_FLAG_V ProcessorStatusFlag = 0x40
	PROCESSOR_STATUS_FLAG_U ProcessorStatusFlag = 0x20
	PROCESSOR_STATUS_FLAG_B ProcessorStatusFlag = 0x10
	PROCESSOR_STATUS_FLAG_D ProcessorStatusFlag = 0x08
	PROCESSOR_STATUS_FLAG_I ProcessorStatusFlag = 0x04
	PROCESSOR_STATUS_FLAG_Z ProcessorStatusFlag = 0x02
	PROCESSOR_STATUS_FLAG_C ProcessorStatusFlag = 0x01
)

type ProcessorStatus struct {
	value uint8
}

func NewProcessorStatus() *ProcessorStatus {
	return &ProcessorStatus{uint8(PROCESSOR_STATUS_FLAG_U | PROCESSOR_STATUS_FLAG_B | PROCESSOR_STATUS_FLAG_Z | PROCESSOR_STATUS_FLAG_I)}
}

func (ps *ProcessorStatus) Set(flag ProcessorStatusFlag) {
	ps.value |= uint8(flag)
}

func (ps *ProcessorStatus) Clear(flag ProcessorStatusFlag) {
	ps.value &= ^uint8(flag)
}

func (ps *ProcessorStatus) Update(flag ProcessorStatusFlag, on bool) {
	if on {
		ps.Set(flag)
	} else {
		ps.Clear(flag)
	}
}

func (ps *ProcessorStatus) Get(flag ProcessorStatusFlag) bool {
	return ps.value&uint8(flag) != 0
}

func (ps *ProcessorStatus) UpdateNZ(value uint8) {
	if value>>7 == 0 {
		ps.Clear(PROCESSOR_STATUS_FLAG_N)
	} else {
		ps.Set(PROCESSOR_STATUS_FLAG_N)
	}

	if value == 0 {
		ps.Set(PROCESSOR_STATUS_FLAG_Z)
	} else {
		ps.Clear(PROCESSOR_STATUS_FLAG_Z)
	}
}

func (ps *ProcessorStatus) String() string {
	var mask uint8 = 0x80
	r := ""
	for _, s := range strings.Split("nvubdizc", "") {
		if ps.value&mask != 0 {
			r += strings.ToUpper(s)
		} else {
			r += s
		}
		mask >>= 1
	}

	return r
}
