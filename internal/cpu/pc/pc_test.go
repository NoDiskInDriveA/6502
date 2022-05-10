package pc_test

// import (
// 	"testing"

// 	"github.com/NoDiskInDriveA/6502/internal/cpu/pc"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAddPCLWithOverflow(t *testing.T) {
// 	newPc, info := pc.AddPCL(0x00FF, 0xFF)
// 	assert.EqualValues(t, 0x01FE, newPc)
// 	assert.EqualValues(t, info, pc.OVERFLOW)
// }

// func TestAddPCLNoOverflow(t *testing.T) {
// 	newPc, info := pc.AddPCL(0x2040, 0x40)
// 	assert.EqualValues(t, 0x2080, newPc)
// 	assert.EqualValues(t, info, pc.PAGE_NOT_CROSSED)
// }

// func TestSetPCLWithOverflow(t *testing.T) {
// 	newPc, info := pc.SetPCL(0x00FF, 0xFF)
// 	assert.EqualValues(t, 0x00FE, newPc)
// 	assert.EqualValues(t, info, pc.OVERFLOW)
// }

// func TestSetPCLNoOverflow(t *testing.T) {
// 	newPc, info := pc.SetPCL(0x20FF, 0xFF)
// 	assert.EqualValues(t, 0x00FE, newPc)
// 	assert.EqualValues(t, info, pc.OVERFLOW)
// }
