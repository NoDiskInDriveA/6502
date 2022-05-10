package cpu_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/cpu"
	"github.com/stretchr/testify/assert"
)

func createTestCpu(content []cpu.Opcode) *cpu.Cpu {
	m := cpu.NewMemory()
	m.Set(0xFFFC, 0x00)
	m.Set(0xFFFD, 0x08)
	for offset := range content {
		m.Set(0x0800+uint16(offset), uint8(content[offset]))
	}
	c := cpu.NewCpu()
	c.EnableHaltOpcode(true)
	c.Bus.Memory = m
	return c
}

func TestCpuInit(t *testing.T) {
	t.Skip()
}

func TestOpcodeBneTakeBranch(t *testing.T) {
	program := []cpu.Opcode{
		cpu.OPCODE_LDA_IMMEDIATE,
		0xFF,
		cpu.OPCODE_STA_ZP,
		0x00,
		cpu.OPCODE_SEC_IMPLIED,
		cpu.OPCODE_BCS_RELATIVE,
		cpu.Opcode(int8(0x03)),
		cpu.OPCODE_JMP_ABSOLUTE,
		0x0E,
		0x08,
		cpu.OPCODE_LDA_IMMEDIATE,
		0x01,
		cpu.OPCODE_STA_ZP,
		0x00,
		cpu.OPCODE_HALT,
	}
	c := createTestCpu(program)
	c.Run()
	assert.EqualValues(t, uint8(0x01), c.Bus.Read(0x00))
}

func TestOpcodeBneSkipBranch(t *testing.T) {
	program := []cpu.Opcode{
		cpu.OPCODE_LDA_IMMEDIATE,
		0xFF,
		cpu.OPCODE_STA_ZP,
		0x00,
		cpu.OPCODE_SEC_IMPLIED,
		cpu.OPCODE_BCC_RELATIVE,
		cpu.Opcode(int8(0x03)),
		cpu.OPCODE_JMP_ABSOLUTE,
		0x0E,
		0x08,
		cpu.OPCODE_LDA_IMMEDIATE,
		0x01,
		cpu.OPCODE_STA_ZP,
		0x00,
		cpu.OPCODE_HALT,
	}
	c := createTestCpu(program)
	c.Run()
	assert.EqualValues(t, uint8(0xFF), c.Bus.Read(0x00))
}
