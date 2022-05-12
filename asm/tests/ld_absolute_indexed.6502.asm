.code
.org 0x0800
    START
        LDA #0
        STA $00
        STA $01
        STA $02
        STA $03
        STA $04
        STA $05
        STA $06
        STA $07
        STA $08
        STA $09
        STA $0A
        STA $0B
    TEST_LDA_X
        LDX #0
        LDA DATA,X
        STA $00
        LDX #1
        LDA DATA,X
        STA $01
        LDX #2
        LDA DATA,X
        STA $02
    TEST_LDA_Y
        LDY #0
        LDA DATA,Y
        STA $03
        LDY #1
        LDA DATA,Y
        STA $04
        LDY #2
        LDA DATA,Y
        STA $05
    TEST_LDY_X
        LDX #0
        LDY DATA,X
        STY $06
        LDX #1
        LDY DATA,X
        STY $07
        LDX #2
        LDY DATA,X
        STY $08
    TEST_LDX_Y
        LDY #0
        LDX DATA,Y
        STX $09
        LDY #1
        LDX DATA,Y
        STX $0A
        LDY #2
        LDX DATA,Y
        STX $0B
    DONE
        .byte $F2

.data
.org 0x1FFE
    DATA
    .byte 0x01
    .byte 0x02
    .byte 0x03
    