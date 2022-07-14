package tests

import (
	"testing"
	"time"

	"github.com/NoDiskInDriveA/6502/internal/device"
	"github.com/NoDiskInDriveA/6502/internal/helper"
	"github.com/NoDiskInDriveA/6502/internal/processor/mos_6502"
	"github.com/stretchr/testify/assert"
)

const (
	PRG_LOCATION uint16 = 0x0800
)

func createTestSystem(binPath string) (device.SystemClock, device.ReadWriteDevice) {
	memory := device.NewRamDevice()
	helper.LoadIntoMemory(memory, PRG_LOCATION, binPath)
	memory.Write(0xFFFC, uint8(0xFF&PRG_LOCATION))
	memory.Write(0xFFFD, uint8(PRG_LOCATION>>8))
	p := mos_6502.NewProcessor(memory)
	p.EnableHaltOpcode(true)
	p.Reset()
	clock := device.NewSynchronizedClock()
	clock.AttachDevice(p, 1)
	return clock, memory
}

func TestLdAbsoluteIndexed(t *testing.T) {
	clock, ram := createTestSystem("../asm/tests/ld_absolute_indexed.bin")
	<-clock.RunClock(time.Microsecond)
	assert.EqualValues(t, uint8(0x01), ram.Read(0x00))
	assert.EqualValues(t, uint8(0x02), ram.Read(0x01))
	assert.EqualValues(t, uint8(0x03), ram.Read(0x02))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x03))
	assert.EqualValues(t, uint8(0x02), ram.Read(0x04))
	assert.EqualValues(t, uint8(0x03), ram.Read(0x05))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x06))
	assert.EqualValues(t, uint8(0x02), ram.Read(0x07))
	assert.EqualValues(t, uint8(0x03), ram.Read(0x08))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x09))
	assert.EqualValues(t, uint8(0x02), ram.Read(0x0A))
	assert.EqualValues(t, uint8(0x03), ram.Read(0x0B))
}

func TestStAbsoluteIndexed(t *testing.T) {
	clock, ram := createTestSystem("../asm/tests/st_absolute_indexed.bin")
	<-clock.RunClock(time.Microsecond)
	assert.EqualValues(t, uint8(0x01), ram.Read(0x1EFE))
	assert.EqualValues(t, uint8(0x02), ram.Read(0x1EFF))
	assert.EqualValues(t, uint8(0x03), ram.Read(0x1F00))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x1FFE))
	assert.EqualValues(t, uint8(0x02), ram.Read(0x1FFF))
	assert.EqualValues(t, uint8(0x03), ram.Read(0x2000))
}

func TestUnaryLogic(t *testing.T) {
	clock, ram := createTestSystem("../asm/tests/unary_logic_test.bin")
	<-clock.RunClock(time.Microsecond)
	assert.EqualValues(t, uint8(0x7F), ram.Read(0x4000))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x4001))
	assert.EqualValues(t, uint8(0xFF), ram.Read(0x4002))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x4003))
	assert.EqualValues(t, uint8(0xFE), ram.Read(0x4004))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x4005))
	assert.EqualValues(t, uint8(0xFF), ram.Read(0x4006))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x4007))
	assert.EqualValues(t, uint8(0x7F), ram.Read(0x4008))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x4009))
	assert.EqualValues(t, uint8(0xFF), ram.Read(0x400A))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x400B))
	assert.EqualValues(t, uint8(0xFE), ram.Read(0x400C))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x400D))
	assert.EqualValues(t, uint8(0xFF), ram.Read(0x400E))
	assert.EqualValues(t, uint8(0x01), ram.Read(0x400F))
}
