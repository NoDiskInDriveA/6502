.code
.org 0x0800
    START
        LDA #$FF
        STA $00
        CLC
        BCS TEST_GOOD
        JMP TEST_BAD
    TEST_GOOD
        LDA #$01
        STA $00
        JMP TEST_END
    TEST_BAD
        LDA #$02
        STA $00
        JMP TEST_END
    TEST_END
        JMP DONE
    DONE
        JMP DONE