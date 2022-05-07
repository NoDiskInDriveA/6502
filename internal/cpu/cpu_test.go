package cpu_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/cpu"
	"github.com/stretchr/testify/assert"
)

func TestCpuInit(t *testing.T) {
	t.Skip()
}

func TestSetPcBytes(t *testing.T) {
	var pc cpu.ProgramCounter = 0x1000
	assert.EqualValues(t, 0x1000, pc)
	pc.SetHigh(0x20)
	assert.EqualValues(t, 0x2000, pc)
	pc.SetLow(0x40)
	assert.EqualValues(t, 0x2040, pc)
}
