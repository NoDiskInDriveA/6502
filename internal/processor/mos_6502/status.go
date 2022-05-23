package mos_6502

import (
	"strings"
)

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

type processorStatus struct {
	value uint8
}

func NewProcessorStatus() *processorStatus {
	return &processorStatus{uint8(PROCESSOR_STATUS_FLAG_U | PROCESSOR_STATUS_FLAG_B | PROCESSOR_STATUS_FLAG_Z | PROCESSOR_STATUS_FLAG_I)}
}

func (ps *processorStatus) Set(flag ProcessorStatusFlag) {
	ps.value |= uint8(flag)
}

func (ps *processorStatus) Clear(flag ProcessorStatusFlag) {
	ps.value &= ^uint8(flag)
}

func (ps *processorStatus) Update(flag ProcessorStatusFlag, on bool) {
	if on {
		ps.Set(flag)
	} else {
		ps.Clear(flag)
	}
}

func (ps *processorStatus) Get(flag ProcessorStatusFlag) bool {
	return ps.value&uint8(flag) != 0
}

func (ps *processorStatus) UpdateNZ(value uint8) {
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

func (ps *processorStatus) String() string {
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
