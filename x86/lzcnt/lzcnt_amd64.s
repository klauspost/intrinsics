// func lzcntU32(a uint32) uint32
TEXT ·lzcntU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func lzcntU64(a uint64) uint64
TEXT ·lzcntU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

