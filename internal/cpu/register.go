package cpu

import (
	"strings"
)

const (
	STATUS_FLAG_N uint8 = 0x80
	STATUS_FLAG_V uint8 = 0x40
	STATUS_FLAG_U uint8 = 0x20
	STATUS_FLAG_B uint8 = 0x10
	STATUS_FLAG_D uint8 = 0x08
	STATUS_FLAG_I uint8 = 0x04
	STATUS_FLAG_Z uint8 = 0x02
	STATUS_FLAG_C uint8 = 0x01
)

type ProcessorStatus struct {
	value byte
}

func NewProcessorStatus() *ProcessorStatus {
	return &ProcessorStatus{uint8(STATUS_FLAG_U)}
}

func (ps *ProcessorStatus) Set(flag uint8) {
	ps.value |= uint8(flag)
}

func (ps *ProcessorStatus) Clear(flag uint8) {
	ps.value &= ^uint8(flag)
}

func (ps *ProcessorStatus) Update(flag uint8, on bool) {
	if on {
		ps.Set(flag)
	} else {
		ps.Clear(flag)
	}
}

func (ps *ProcessorStatus) Get(flag uint8) bool {
	return ps.value&uint8(flag) != 0
}

func (ps *ProcessorStatus) UpdateNZ(value uint8) {
	if value>>7 == 0 {
		ps.Clear(STATUS_FLAG_N)
	} else {
		ps.Set(STATUS_FLAG_N)
	}

	if value == 0 {
		ps.Set(STATUS_FLAG_Z)
	} else {
		ps.Clear(STATUS_FLAG_Z)
	}
}

func (ps *ProcessorStatus) String() string {
	var mask byte = 0x80
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
