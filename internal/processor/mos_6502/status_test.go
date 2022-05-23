package mos_6502_test

import (
	"testing"

	"github.com/NoDiskInDriveA/6502/internal/processor/mos_6502"
	"github.com/stretchr/testify/assert"
)

func TestSetAndClearBit(t *testing.T) {
	st := mos_6502.NewProcessorStatus()
	assert.False(t, st.Get(mos_6502.PROCESSOR_STATUS_FLAG_C))
	assert.True(t, st.Get(mos_6502.PROCESSOR_STATUS_FLAG_Z))

	st.Set(mos_6502.PROCESSOR_STATUS_FLAG_C)
	assert.True(t, st.Get(mos_6502.PROCESSOR_STATUS_FLAG_C))
	assert.True(t, st.Get(mos_6502.PROCESSOR_STATUS_FLAG_Z))

	st.Clear(mos_6502.PROCESSOR_STATUS_FLAG_Z)
	assert.True(t, st.Get(mos_6502.PROCESSOR_STATUS_FLAG_C))
	assert.False(t, st.Get(mos_6502.PROCESSOR_STATUS_FLAG_Z))
}

func TestStringOutput(t *testing.T) {
	st := mos_6502.NewProcessorStatus()
	st.Set(mos_6502.PROCESSOR_STATUS_FLAG_B | mos_6502.PROCESSOR_STATUS_FLAG_C)
	assert.Equal(t, "nvUBdIZC", st.String())
}
