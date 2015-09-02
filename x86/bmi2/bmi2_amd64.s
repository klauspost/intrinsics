// func bzhiU32(a uint32, index uint32) uint32
TEXT ·bzhiU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL index+4(FP),R9

	// TODO: Code missing
	// Could be:
	// BZHI R8, R9

	MOVL $0, ret+8(FP)
	RET

// func bzhiU64(a uint64, index uint32) uint64
TEXT ·bzhiU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVL index+8(FP),R9

	// TODO: Code missing
	// Could be:
	// BZHI R8, R9

	MOVQ $0, ret+12(FP)
	RET

// func pdepU32(a uint32, mask uint32) uint32
TEXT ·pdepU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL mask+4(FP),R9

	// TODO: Code missing
	// Could be:
	// PDEP R8, R9

	MOVL $0, ret+8(FP)
	RET

// func pdepU64(a uint64, mask uint64) uint64
TEXT ·pdepU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// PDEP R8, R9

	MOVQ $0, ret+16(FP)
	RET

// func pextU32(a uint32, mask uint32) uint32
TEXT ·pextU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL mask+4(FP),R9

	// TODO: Code missing
	// Could be:
	// PEXT R8, R9

	MOVL $0, ret+8(FP)
	RET

// func pextU64(a uint64, mask uint64) uint64
TEXT ·pextU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// PEXT R8, R9

	MOVQ $0, ret+16(FP)
	RET

