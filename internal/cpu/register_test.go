package cpu_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/cpu"
	"github.com/stretchr/testify/assert"
)

func TestNewInstance(t *testing.T) {
	st := cpu.NewProcessorStatus()
	assert.False(t, st.Get(cpu.STATUS_FLAG_C))
}

func TestSetAndClearBit(t *testing.T) {
	st := cpu.NewProcessorStatus()
	assert.False(t, st.Get(cpu.STATUS_FLAG_C))
	assert.False(t, st.Get(cpu.STATUS_FLAG_Z))
	st.Set(cpu.STATUS_FLAG_C)
	assert.True(t, st.Get(cpu.STATUS_FLAG_C))
	assert.False(t, st.Get(cpu.STATUS_FLAG_Z))
	st.Set(cpu.STATUS_FLAG_C | cpu.STATUS_FLAG_Z)
	assert.True(t, st.Get(cpu.STATUS_FLAG_C))
	assert.True(t, st.Get(cpu.STATUS_FLAG_Z))
	st.Clear(cpu.STATUS_FLAG_C)
	assert.False(t, st.Get(cpu.STATUS_FLAG_C))
	assert.True(t, st.Get(cpu.STATUS_FLAG_Z))
}

func TestStringOutput(t *testing.T) {
	st := cpu.NewProcessorStatus()
	st.Set(cpu.STATUS_FLAG_B | cpu.STATUS_FLAG_C)
	assert.Equal(t, "nvUBdizC", st.String())
}
