            * = $0800
.code
            LDA #00
            STA $00
            STA $01
            LDX #30
    LOOP
            CLC
            TXA
            ADC $00
            STA $00
            LDA $01
            ADC #00
            DEX
            BNE LOOP