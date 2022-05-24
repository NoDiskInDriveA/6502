.include "../macros.6502.asm"

.code
.org 0x0800
    TEST_ACC_ROR_CO_NI // -> $7F
        LDA #$FF
        CLC
        ROR A
        STA RESULTS
        BCC TEST_ACC_ROR_NO_CI
        LDA #$01
        STA RESULTS+1
    TEST_ACC_ROR_NO_CI // -> $FF
        LDA #$FE
        SEC
        ROR A
        STA RESULTS+2
        BCS TEST_ACC_ROL_CO_NI
        LDA #$01
        STA RESULTS+3
    TEST_ACC_ROL_CO_NI // -> $FE
        LDA #$FF
        CLC
        ROL A
        STA RESULTS+4
        BCC TEST_ACC_ROL_NO_CI
        LDA #$01
        STA RESULTS+5
    TEST_ACC_ROL_NO_CI // -> $FF
        LDA #$7F
        SEC
        ROL A
        STA RESULTS+6
        BCS TEST_MEM_ROR_CO_NI
        LDA #$01
        STA RESULTS+7
    TEST_MEM_ROR_CO_NI // -> $7F
        CLC
        ROR RESULTS+8
        BCC TEST_MEM_ROR_NO_CI
        LDA #$01
        STA RESULTS+9
    TEST_MEM_ROR_NO_CI // -> $FF
        LDA #$FE
        SEC
        ROR RESULTS+10
        BCS TEST_MEM_ROL_CO_NI
        LDA #$01
        STA RESULTS+11
    TEST_MEM_ROL_CO_NI // -> $7F
        CLC
        ROL RESULTS+12
        BCC TEST_MEM_ROL_NO_CI
        LDA #$01
        STA RESULTS+13
    TEST_MEM_ROL_NO_CI // -> $FF
        LDA #$FE
        SEC
        ROL RESULTS+14
        BCS TEST_END
        LDA #$01
        STA RESULTS+15
    TEST_END
        HALT()

.segment "DATA"
.org 0x4000
    RESULTS
    .byte $00,$00,$00,$00,$00,$00,$00,$00
    .byte $FF,$00,$FE,$00,$FF,$00,$7F,$00
