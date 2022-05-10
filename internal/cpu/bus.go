package cpu

type Memory []uint8

func NewMemory() *Memory {
	m := make(Memory, 65536)
	return &m
}

func (m *Memory) Get(address uint16) uint8 {
	return (*m)[address]
}

func (m *Memory) Set(address uint16, value uint8) {
	(*m)[address] = value
}

func (m *Memory) Slice(startAddress uint16, endAddress uint16) []uint8 {
	return (*m)[startAddress:endAddress]
}

func (m *Memory) LoadAt(startAddress uint16, bytes []uint8) {
	copy((*m)[startAddress:], bytes)
}

type Bus struct {
	Memory *Memory
}

func NewBus() *Bus {
	memory := NewMemory()
	return &Bus{memory}
}

func (b *Bus) Read(address uint16) uint8 {
	return b.Memory.Get(address)
}

func (b *Bus) Write(address uint16, value uint8) {
	b.Memory.Set(address, value)
}
