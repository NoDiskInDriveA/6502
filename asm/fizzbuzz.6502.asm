.include "macros.6502.asm"

.code
.org 0x0800
    _START
        LDX #100
    LOOP
        LDY #$00
    CHECK_THREE
        TXA
        DIVISIBLE_BY_THREE()
        CMP #$01
        BNE CHECK_FIVE
        INY
    CHECK_FIVE
        TXA
        DIVISIBLE_BY_FIVE()
        CMP #$01
        BNE DEC_LOOP
        INY
        INY
    DEC_LOOP
        STY $10,X
        DEX
        BNE LOOP
    _END
        HALT()

// check whether value in A is divisible by 3, after routine A contains 1 if divisible, 0 otherwise
.function DIVISIBLE_BY_THREE()
    @INPUT .equ $00
    @ODD_BITS .equ $08
    @EVEN_BITS .equ $09
    // BACKUP
        STX $01
        STY $02
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
        LDX $01
        LDY $02
.endfunction

// check whether value in A is divisible by 5, after routine A contains 1 if divisible, 0 otherwise
.function DIVISIBLE_BY_FIVE()
    @PARITY_N .equ $08
    @PARITY_N_1 .equ $09
    @PARITY_N_2 .equ $0A
    @PARITY_N_3 .equ $0B
        // BACKUP
        STA $00
        STX $01
        STY $02
        // INIT
        LDX #$00
        LDY #$02
        // zero out
        STX @PARITY_N
        STX @PARITY_N_1
        STX @PARITY_N_2
        STX @PARITY_N_3

        LDA $00
    @BIT_1
        LDX @PARITY_N
        LSR A
        BCC @BIT_2
        INX
        STX @PARITY_N
    @BIT_2
        LDX @PARITY_N_1
        LSR A
        BCC @BIT_3
        INX
        STX @PARITY_N_1
    @BIT_3
        LDX @PARITY_N_2
        LSR A
        BCC @BIT_4
        INX
        STX @PARITY_N_2
    @BIT_4
        LDX @PARITY_N_3
        LSR A
        BCC @BIT_END
        INX
        STX @PARITY_N_3
    @BIT_END
        DEY
        BNE @BIT_1
    @COMPARE
        LDA @PARITY_N
        CMP @PARITY_N_2
        BNE @NOT_DIVISIBLE
        LDA @PARITY_N_1
        CMP @PARITY_N_3
        BNE @NOT_DIVISIBLE
        LDA #$01
        BNE @RETURN

    @NOT_DIVISIBLE
        LDA #$00
     @RETURN
        // RESTORE
        LDX $01
        LDY $02
.endfunction

.segment "DATA"
.org 0x2000
    INPUT1
    .byte $64
    INPUT2
    .byte $63
    INPUT3
    .byte 235