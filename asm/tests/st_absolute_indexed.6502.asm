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
    TEST_STA_X
        LDX #0
        LDA #$01
        STA DATA_X,X
        LDX #1
        LDA #$02
        STA DATA_X,X
        LDX #2
        LDA #$03
        STA DATA_X,X
    TEST_STA_Y
        LDY #0
        LDA #$01
        STA DATA_Y,Y
        LDY #1
        LDA #$02
        STA DATA_Y,Y
        LDY #2
        LDA #$03
        STA DATA_Y,Y
    DONE
        .byte $F2

.segment "X"
.org 0x1EFE
    DATA_X
    .byte 0x00
    .byte 0x00
    .byte 0x00
.segment "Y"
.org 0x1FFE
    DATA_Y
    .byte 0x00
    .byte 0x00
    .byte 0x00