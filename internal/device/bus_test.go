package device_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/device"
	"github.com/stretchr/testify/assert"
)

func createTestBus() device.AddressDecodingBus {
	bus := device.NewSystemBus()
	bus.AttachMappedDevice(0x0000, 0xFFFF, 0xFFFF, device.NewRamDevice())
	return bus
}

func TestBusReadWrite(t *testing.T) {
	bus := createTestBus()
	bus.Write(0x100, 0x40)
	assert.EqualValues(t, 0x40, bus.Read(0x100))
	bus.Write(0x100, 0x80)
	assert.EqualValues(t, 0x80, bus.Read(0x100))
}

func TestMemoryLoadAt(t *testing.T) {
	mem := device.NewRamDevice()
	prg := []uint8{0xA9, 0xFF}
	mem.LoadAt(0x4020, prg)
	assert.EqualValues(t, 0xA9, mem.Read(0x4020))
	assert.EqualValues(t, 0xFF, mem.Read(0x4021))
}
