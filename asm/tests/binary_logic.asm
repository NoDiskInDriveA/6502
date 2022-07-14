.code
    .org 0x0800

    TEST_START
    SUITE_IMMEDIATE
        LDA #$7F
        ORA #$80
        STA $00 // -> FF

        LDA #$FE
        AND #$7F
        STA $01 // -> 7E

        LDA #$FC // 11111100
        EOR #$3F // 00111111
        STA $02 // -> 11000011 -> C3

    SUITE_ZP
        LDA #$80
        STA $10
        LDA #$7F
        ORA $10
        STA $03 // -> FF

        LDA #$7F
        STA $11
        LDA #$FE
        AND $11
        STA $04 // -> 7E

        LDA #$3F
        STA $12
        LDA #$FC
        EOR $12
        STA $05 // -> C3
    TEST_END
        .byte $F2