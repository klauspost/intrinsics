// func rdrand16Step(val uint16) int
TEXT ·rdrand16Step(SB),7,$0
	MOVW val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func rdrand32Step(val uint32) int
TEXT ·rdrand32Step(SB),7,$0
	MOVL val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func rdrand64Step(val uint64) int
TEXT ·rdrand64Step(SB),7,$0
	MOVQ val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

