.target "6502"

// all values in MM_ are adresses
    MM_CPU_VECTORS      .equ $FFFA
    MM_CPU_VECTOR_NMI   .equ $FFFA
    MM_CPU_VECTOR_RESET .equ $FFFC
    MM_CPU_VECTOR_INT   .equ $FFFE

.macro HALT()
    .byte $F2
.endmacro
