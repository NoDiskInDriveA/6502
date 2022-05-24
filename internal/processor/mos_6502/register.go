package mos_6502

type RegisterDef string

const (
	REGISTER_A  RegisterDef = "A"
	REGISTER_X  RegisterDef = "X"
	REGISTER_Y  RegisterDef = "Y"
	REGISTER_SP RegisterDef = "SP"
	DATA_LATCH  RegisterDef = "D"
)
