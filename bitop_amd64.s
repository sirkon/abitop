TEXT ·bitSet(SB), $0-24
    XORQ    BX, BX
    MOVQ    p + 0(FP), SI
    MOVQ    p + 8(FP), AX
    LOCK
    BTSQ    AX, (SI)
    ADCQ    $0, BX
    MOVQ    BX, p + 16(FP)
    RET

TEXT ·bitUnset(SB), $0-16
    XORQ    BX, BX
    MOVQ    p + 0(FP), SI
    MOVQ    p + 8(FP), AX
    LOCK
    BTRQ    AX, (SI)
    ADCQ    $0, BX
    MOVQ    BX, p + 16(FP)
    RET

