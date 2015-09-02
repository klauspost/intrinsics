// func bextrU32(a uint32, start uint32, len uint32) uint32
TEXT ·bextrU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL start+4(FP),R9
	MOVL len+8(FP),R10

	// TODO: Code missing

	MOVL $0, ret+12(FP)
	RET

// func bextrU64(a uint64, start uint32, len uint32) uint64
TEXT ·bextrU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVL start+8(FP),R9
	MOVL len+12(FP),R10

	// TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func blsiU32(a uint32) uint32
TEXT ·blsiU32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// BLSI R8

	MOVL $0, ret+4(FP)
	RET

// func blsiU64(a uint64) uint64
TEXT ·blsiU64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BLSI R8

	MOVQ $0, ret+8(FP)
	RET

// func blsmskU32(a uint32) uint32
TEXT ·blsmskU32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// BLSMSK R8

	MOVL $0, ret+4(FP)
	RET

// func blsmskU64(a uint64) uint64
TEXT ·blsmskU64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BLSMSK R8

	MOVQ $0, ret+8(FP)
	RET

// func blsrU32(a uint32) uint32
TEXT ·blsrU32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// BLSR R8

	MOVL $0, ret+4(FP)
	RET

// func blsrU64(a uint64) uint64
TEXT ·blsrU64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BLSR R8

	MOVQ $0, ret+8(FP)
	RET

// func tzcnt32(a uint32) int
TEXT ·tzcnt32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// TZCNT R8

	MOVQ $0, ret+4(FP)
	RET

// func tzcnt64(a uint64) int64
TEXT ·tzcnt64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// TZCNT R8

	MOVQ $0, ret+8(FP)
	RET

// func tzcntU32(a uint32) uint32
TEXT ·tzcntU32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// TZCNT R8

	MOVL $0, ret+4(FP)
	RET

// func tzcntU64(a uint64) uint64
TEXT ·tzcntU64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// TZCNT R8

	MOVQ $0, ret+8(FP)
	RET

