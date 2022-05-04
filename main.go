package main

import "github.com/NoDiskInDriveA/6502/internal/cpu"

func main() {
	cpu := cpu.NewCpu()
	cpu.Run()
}
