.include "macros.6502.asm"

// check whether value in ADDR is divisible by 3, after macro ADDR contains 1 if divisible, 0 otherwise
.macro DIVISIBLE_BY_THREE(ADDR)
    @ODD_BITS .equ $08
    @EVEN_BITS .equ $09
    // X -> LOOP VAR, Y -> COUNT VAR
        LDA #$00
        STA @ODD_BITS
        STA @EVEN_BITS

    @CALC_ODD
        LDA ADDR
        LDX #$4 
        LDY #$0 
        AND #%01010101
    @LOOP_ODD
        LSR A
        BCC @CONT_ODD
        INY
    @CONT_ODD
        LSR A
        DEX
        BNE @LOOP_ODD
    
        STY @ODD_BITS

    @CALC_EVEN
        LDA ADDR
        LDX #$4 
        LDY #$0 
        AND #%10101010
    @LOOP_EVEN
        LSR A
        LSR A
        BCC @CONT_EVEN
        INY
    @CONT_EVEN
        DEX
        BNE @LOOP_EVEN
    
        STY @EVEN_BITS
    
    @CHECK_DIFF
        LDA @ODD_BITS
        SEC
        SBC @EVEN_BITS
        STA ADDR+1 //DEBUG
        BEQ @IS_DIVISIBLE
        SEC
        SBC #$03
        STA ADDR+2 //DEBUG
        BEQ @IS_DIVISIBLE
        LDA #$00
        BEQ @WRITE_RESULT

    @IS_DIVISIBLE
        LDA #$01

    @WRITE_RESULT
        STA ADDR

.endmacro

.code
.org 0x0800
    _START
        DIVISIBLE_BY_THREE(INPUT1)
        DIVISIBLE_BY_THREE(INPUT2)
    _END
        HALT()

.segment "DATA"
.org 0x2000
    INPUT1
    .byte $64,$00,$00
    INPUT2
    .byte $63,$00,$00