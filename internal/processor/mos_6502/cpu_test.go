package mos_6502_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/device"
	"github.com/NoDiskInDriveA/6502/internal/processor/mos_6502"
	"github.com/stretchr/testify/assert"
)

func createTestSystem(content []mos_6502.Opcode) (mos_6502.Processor, device.ReadWriteDevice) {
	ram := device.NewRamDevice()

	ram.Write(0xFFFC, 0x00)
	ram.Write(0xFFFD, 0x08)
	for offset := range content {
		ram.Write(0x0800+uint16(offset), uint8(content[offset]))
	}
	proc := mos_6502.NewProcessor(ram)
	proc.EnableHaltOpcode(true)
	proc.Reset()
	return proc, ram
}

func TestCpuInit(t *testing.T) {
	t.Skip()
}

func TestOpcodeBneTakeBranch(t *testing.T) {
	program := []mos_6502.Opcode{
		mos_6502.OPCODE_LDA_IMMEDIATE,
		0xFF,
		mos_6502.OPCODE_STA_ZP,
		0x00,
		mos_6502.OPCODE_SEC_IMPLIED,
		mos_6502.OPCODE_BCS_RELATIVE,
		mos_6502.Opcode(int8(0x03)),
		mos_6502.OPCODE_JMP_ABSOLUTE,
		0x0E,
		0x08,
		mos_6502.OPCODE_LDA_IMMEDIATE,
		0x01,
		mos_6502.OPCODE_STA_ZP,
		0x00,
		mos_6502.OPCODE_HALT,
	}
	proc, ram := createTestSystem(program)
	for range program {
		proc.Tick()
	}
	assert.EqualValues(t, uint8(0x01), ram.Read(0x00))
}

func TestOpcodeBneSkipBranch(t *testing.T) {
	program := []mos_6502.Opcode{
		mos_6502.OPCODE_LDA_IMMEDIATE,
		0xFF,
		mos_6502.OPCODE_STA_ZP,
		0x00,
		mos_6502.OPCODE_SEC_IMPLIED,
		mos_6502.OPCODE_BCC_RELATIVE,
		mos_6502.Opcode(int8(0x03)),
		mos_6502.OPCODE_JMP_ABSOLUTE,
		0x0E,
		0x08,
		mos_6502.OPCODE_LDA_IMMEDIATE,
		0x01,
		mos_6502.OPCODE_STA_ZP,
		0x00,
		mos_6502.OPCODE_HALT,
	}
	proc, ram := createTestSystem(program)
	for range program {
		proc.Tick()
	}
	assert.EqualValues(t, uint8(0xFF), ram.Read(0x00))
}
