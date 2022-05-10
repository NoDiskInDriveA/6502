.code
.org $0800
            LDA #$10
            STA $00
            JSR Func
            LDA #$30
            STA $02

    End     NOP
            JMP End

    Func    LDA #$20
            STA $01
            RTS
    