package main

import (
	"github.com/NoDiskInDriveA/6502/internal/arch"
)

const (
	PRG_LOCATION uint16 = 0x0800
	ROM_LOCATION uint16 = 0xE000
)

func main() {
	arch.NewDebugFantasyArchitecture().Run()
}
