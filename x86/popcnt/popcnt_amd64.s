// func popcntU32(a uint32) int
TEXT ·popcntU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func popcntU64(a uint64) int64
TEXT ·popcntU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func popcnt32(a int) int
TEXT ·popcnt32(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func popcnt64(a int64) int
TEXT ·popcnt64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

