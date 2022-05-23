package device

import "fmt"

type ReadableDevice interface {
	Read(address uint16) uint8
}

type WritableDevice interface {
	Write(address uint16, value uint8)
}

type ReadWriteDevice interface {
	ReadableDevice
	WritableDevice
}

type AddressDecodingBus interface {
	ReadableDevice
	WritableDevice
	AttachMappedDevice(startAddress uint16, length uint16, effectiveMask uint16, device ReadWriteDevice)
}

type systemBus struct {
	mappedDevices []memoryMappedDevice
}

type memoryMappedDevice struct {
	startAddress  uint16
	endAddress    uint16
	effectiveMask uint16
	device        ReadWriteDevice
}

func NewSystemBus() AddressDecodingBus {
	m := make([]memoryMappedDevice, 0, 8)
	return &systemBus{m}
}

func (sb *systemBus) Read(address uint16) uint8 {
	md := sb.decodeDevice(address)
	return md.device.Read(address & md.effectiveMask)
}

func (sb *systemBus) Write(address uint16, value uint8) {
	md := sb.decodeDevice(address)
	md.device.Write(address&md.effectiveMask, value)
}

func (sb *systemBus) AttachMappedDevice(startAddress uint16, length uint16, effectiveMask uint16, device ReadWriteDevice) {
	sb.mappedDevices = append(sb.mappedDevices, memoryMappedDevice{startAddress, startAddress + length, effectiveMask, device})
}

func (sb *systemBus) decodeDevice(address uint16) memoryMappedDevice {
	for _, device := range sb.mappedDevices {
		if address >= device.startAddress && address < device.endAddress {
			return device
		}
	}
	panic(fmt.Sprintf("No device mapped at 0x%04X", address))
}
