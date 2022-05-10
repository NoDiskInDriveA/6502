package pc

type PageCrossing uint8

const (
	PAGE_NOT_CROSSED       PageCrossing = iota
	PAGE_CROSSED_UNDERFLOW PageCrossing = iota
	PAGE_CROSSED_OVERFLOW  PageCrossing = iota
)

// func AddPCL(pc uint16, add uint8) (uint16, pcOpInfo) {
// 	calc := pc + uint16(add)
// 	return calc, pc>>8 != calc>>8
// }

// func SetPCL(pc uint16, add uint8) (uint16, pcOpInfo) {
// 	calc := pc + uint16(add)
// 	return (pc & 0xFF00) | (calc & 0x00FF), pc>>8 != calc>>8
// }

func AddPcSigned(pc uint16, add uint8) (uint16, PageCrossing) {
	calc := int32(int16(pc)) + int32(int8(add))
	newPc := uint16(calc)
	pchOld, pchNew := pc>>8, newPc>>8

	if pchNew < pchOld {
		return newPc, PAGE_CROSSED_UNDERFLOW
	}
	if pchNew > pchOld {
		return newPc, PAGE_CROSSED_OVERFLOW
	}

	return newPc, PAGE_NOT_CROSSED
}

func AddPcUnSigned(pc uint16, add uint8) (uint16, PageCrossing) {
	calc := int32(int16(pc)) + int32(int8(add))
	newPc := uint16(calc)
	pchOld, pchNew := pc>>8, newPc>>8

	if pchNew < pchOld {
		return newPc, PAGE_CROSSED_UNDERFLOW
	}
	if pchNew > pchOld {
		return newPc, PAGE_CROSSED_OVERFLOW
	}

	return newPc, PAGE_NOT_CROSSED
}

// func AddPCH(pc uint16, add uint8) (uint16, pcOpInfo) {
// 	calc := uint32(pc + uint16(add)<<8)
// 	return uint16(calc), calc&0x1FFFF != 0
// }
