package cpu

type Memory []uint8

func NewMemory() *Memory {
	m := make(Memory, 65536)
	p := NewProgrammer(&m)
	// LDA #10
	p.Program(OPCODE_LDA_IMMEDIATE)
	p.Program(0x10)
	// STA $08
	p.Program(OPCODE_STA_DIRECT_ZP)
	p.Program(0x08)
	// ADC #20
	p.Program(OPCODE_ADC_IMMEDIATE)
	p.Program(0x20)
	// STA $09
	p.Program(OPCODE_STA_DIRECT_ZP)
	p.Program(0x09)
	// ADC #FF
	p.Program(OPCODE_ADC_IMMEDIATE)
	p.Program(0xFF)
	// STA $0A
	p.Program(OPCODE_STA_DIRECT_ZP)
	p.Program(0x0A)
	// :START
	p.Label("START")
	// NOP
	p.Program(OPCODE_NOP_IMPLIED)
	// JMP :start
	p.Program(OPCODE_JMP_ABSOLUTE)
	p.LabeledAddress("START")
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

type Bus struct {
	memory *Memory
}

func NewBus() *Bus {
	memory := NewMemory()
	// hardwire reset vector to 0x8000
	memory.Set(RESET_VECTOR_LSB, 0x00)
	memory.Set(RESET_VECTOR_MSB, 0x80)

	return &Bus{memory}
}

func (b *Bus) Read(address uint16) uint8 {
	return b.memory.Get(address)
}

func (b *Bus) Write(address uint16, value uint8) {
	b.memory.Set(address, value)
}

type Programmer struct {
	memory  *Memory
	address uint16
	labels  map[string]uint16
}

func NewProgrammer(memory *Memory) *Programmer {
	return &Programmer{memory, 0x8000, make(map[string]uint16)}
}

func (p *Programmer) StartAddress(address uint16) {
	p.address = address
}

func (p *Programmer) Program(value uint8) {
	p.memory.Set(p.address, value)
	p.address++
}

func (p *Programmer) Label(label string) {
	p.labels[label] = p.address + 1
}

func (p *Programmer) LabeledAddress(label string) {
	addr, found := p.labels[label]
	if !found {
		panic("missing label")
	}
	p.Program(uint8(addr))
	p.Program(uint8(addr >> 8))
}
