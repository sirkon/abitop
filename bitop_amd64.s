TEXT ·BitSet(SB), $0-16
    MOVQ    p + 0(FP), SI
    MOVQ    p + 8(FP), AX
    LOCK
    BTSQ    AX, (SI)
    RET

TEXT ·BitUnset(SB), $0-16
    MOVQ    p + 0(FP), SI
    MOVQ    p + 8(FP), AX
    LOCK
    BTRQ    AX, (SI)
    RET

