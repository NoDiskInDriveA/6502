package cpu_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/cpu"
	"github.com/stretchr/testify/assert"
)

func TestBusReadWrite(t *testing.T) {
	bus := cpu.NewBus()
	bus.Write(0x100, 0x40)
	assert.Equal(t, 0x40, bus.Read(0x100))
	bus.Write(0x100, 0x80)
	assert.Equal(t, 0x80, bus.Read(0x100))
}

func TestMemoryLoadAt(t *testing.T) {
	mem := cpu.NewMemory()
	prg := []uint8{0xA9, 0xFF}
	mem.LoadAt(0x4020, prg)
	assert.EqualValues(t, 0xA9, mem.Get(0x4020))
	assert.EqualValues(t, 0xFF, mem.Get(0x4021))
}
