.include "macros.6502.asm"

.code
.org 0x0800
    _START
        LDA INPUT1
        DIVISIBLE_BY_THREE()
        STA INPUT1
        LDA INPUT2
        DIVISIBLE_BY_THREE()
        STA INPUT2
    _END
        HALT()

// check whether value in A is divisible by 3, after routine A contains 1 if divisible, 0 otherwise
.function DIVISIBLE_BY_THREE()
    @INPUT .equ $00
    @ODD_BITS .equ $08
    @EVEN_BITS .equ $09
    // X -> LOOP VAR, Y -> COUNT VAR
        STA @INPUT
    @CALC_ODD
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
        LDA @INPUT
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
        BEQ @IS_DIVISIBLE
        SEC
        SBC #$03
        BEQ @IS_DIVISIBLE
        LDA #$00
        BEQ @END_ROUTINE

    @IS_DIVISIBLE
        LDA #$01

    @END_ROUTINE
.endfunction

.segment "DATA"
.org 0x2000
    INPUT1
    .byte $64
    INPUT2
    .byte $63