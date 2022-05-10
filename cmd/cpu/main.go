package main

import (
	"github.com/NoDiskInDriveA/6502/internal/cpu"
)

const (
	PRG_LOCATION uint16 = 0x0800
	ROM_LOCATION uint16 = 0xE000
)

func main() {
	c := cpu.NewCpu()
	// c.LoadPrgAt(PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/main.bin")
	c.LoadPrgAt(PRG_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/sum.bin")
	// loading ROM last so it can't be overwritten by PRG (until there's an address decoder implemented)
	c.LoadPrgAt(ROM_LOCATION, "/Users/patrick.durold/Projects/GoLang/6502/asm/rom.bin")
	c.EnableHaltOpcode(true)
	c.Run()
}
