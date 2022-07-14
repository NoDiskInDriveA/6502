.code
.org 0x0800
    TEST_START
        LDA #$80
        STA $08
        LDA #$20
        STA $09

        LDA #$7F
        STA $2084
        LDA #$F7
        STA $8024

    NORMAL_LD
        LDY #$04
        LDA ($08),Y
        STA $00
        
    TEST_END
        .byte $F2