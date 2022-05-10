.include "macros.6502.asm"

.code
.org 0x0800
    _START
        LDA #00
        STA $00
        STA $01
        LDX #30
    _ITER
        CLC
        TXA
        ADC $00
        STA $00
        LDA $01
        ADC #00
        STA $01
        DEX
        BNE _ITER
    _END
        HALT()