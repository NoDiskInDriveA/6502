package helper

import (
	"fmt"
	"io/ioutil"

	"github.com/NoDiskInDriveA/6502/internal/device"
)

func LoadIntoMemory(memory device.MemoryDevice, org uint16, binPath string) {
	bytes, err := ioutil.ReadFile(binPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load input file at %s", binPath))
	}
	memory.LoadAt(org, bytes)
}
