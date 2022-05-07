package main

import (
	"github.com/NoDiskInDriveA/6502/internal/cpu"
)

func main() {
	c := cpu.NewCpu()
	c.LoadPrg("/Users/patrick.durold/Projects/GoLang/6502/asm/main.prg")
	c.Run()
}
