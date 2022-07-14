.code
    .org 0x0800

    TEST_START
        LDX #$FF
        LDY #$01

    TEST_CMP_IMM
        LDA #$10
        CMP #$10
        BNE TEST_CMP_IMM_FAIL
        STY $00
        JMP TEST_END

    TEST_CMP_IMM_FAIL:
        STX $00

    TEST_END
        .byte $F2