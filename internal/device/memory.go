package device

type ram struct {
	memory []uint8
}

type MemoryDevice interface {
	ReadWriteDevice
	Slice(startAddress uint16, endAddress uint16) []uint8
	LoadAt(startAddress uint16, bytes []uint8)
}

func NewRamDevice() MemoryDevice {
	m := make([]uint8, 65536)
	return &ram{m}
}

func (r *ram) Read(address uint16) uint8 {
	return r.memory[address]
}

func (r *ram) Write(address uint16, value uint8) {
	r.memory[address] = value
}

func (r *ram) Slice(startAddress uint16, endAddress uint16) []uint8 {
	return r.memory[startAddress:endAddress]
}

func (r *ram) LoadAt(startAddress uint16, bytes []uint8) {
	copy(r.memory[startAddress:], bytes)
}
