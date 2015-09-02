// func xgetbv(a uint32) uint64
TEXT ·xgetbv(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// XGETBV R8

	MOVQ $0, ret+4(FP)
	RET

// func xsetbv(a uint32, val uint64) 
TEXT ·xsetbv(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ val+4(FP),R9

	// TODO: Code missing - could be:
	// XSETBV R8, R9

	RET

