.include "macros.6502.asm"

VECTOR_NMI .equ $0800
VECTOR_RESET .equ $0800
VECTOR_INT .equ $0800

.segment "DATA"
.org 0xE000
    .byte 0b11111111
    .byte 0b11000011
    .byte 0b11000011
    .byte 0b11000011
    .byte 0b11000011
    .byte 0b11000011
    .byte 0b11000011
    .byte 0b11111111

    .byte 0b00011000
    .byte 0b00111000
    .byte 0b01111000
    .byte 0b00011000
    .byte 0b00011000
    .byte 0b00011000
    .byte 0b00011000
    .byte 0b00011000

.segment "VECTORS"
.org CPU_VECTORS
    .word VECTOR_NMI
    .word VECTOR_RESET
    .word VECTOR_INT
