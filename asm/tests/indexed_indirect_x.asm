.code
.org 0x0800
    TEST_START
        LDA #$80
        STA $08
        LDA #$20
        STA $09

        LDA #$7F
        STA $2080
        LDA #$F7
        STA $8020
        
    NORMAL_LD
        LDX #$04
        LDA ($4,X)
        STA $00
    
    PAGE_OVERFLOW_LD
        LDX #$FC
        LDA ($0C,X)
        STA $01

    NORMAL_ST
        LDX #$04
        LDA #$FF
        STA ($4,X)
        
    TEST_END
        .byte $F2