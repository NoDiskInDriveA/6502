.include "macros.6502.asm"

.code
.org 0x0800
    _START
        LDA #00
        STA $00
        STA $01
        LDX #255
        LDY #255
    _ITER
        DEX
        BNE _ITER
    _ITER2
        DEY
        BNE _ITER2
    _END
        HALT()