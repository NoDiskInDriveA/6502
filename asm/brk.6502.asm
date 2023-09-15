.include "macros.6502.asm"

.code
.org 0x0800
        BRK #$7F
        LDA #$FF