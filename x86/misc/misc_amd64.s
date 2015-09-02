// func allowCpuFeatures(a uint64) 
TEXT ·allowCpuFeatures(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing

	RET

// func bitScanForward(a int) int
TEXT ·bitScanForward(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BSF R8

	MOVQ $0, ret+8(FP)
	RET

// func bitScanReverse(a int) int
TEXT ·bitScanReverse(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BSR R8

	MOVQ $0, ret+8(FP)
	RET

// func bswap(a int) int
TEXT ·bswap(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BSWAP R8

	MOVQ $0, ret+8(FP)
	RET

// func bswap64(a int64) int64
TEXT ·bswap64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// BSWAP R8

	MOVQ $0, ret+8(FP)
	RET

// func castf32U32(a float32) uint32
TEXT ·castf32U32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func castf64U64(a float64) uint64
TEXT ·castf64U64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func castu32F32(a uint32) float32
TEXT ·castu32F32(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func castu64F64(a uint64) float64
TEXT ·castu64F64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func cvtshSs(a uint16) float32
TEXT ·cvtshSs(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func cvtssSh(a float32, imm8 int) uint16
TEXT ·cvtssSh(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ imm8+4(FP),R9

	// TODO: Code missing

	MOVW $0, ret+12(FP)
	RET

// func lrotl(a uint32, shift int) uint32
TEXT ·lrotl(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	// TODO: Code missing - could be:
	// ROL R8, R9

	MOVL $0, ret+12(FP)
	RET

// func lrotr(a uint32, shift int) uint32
TEXT ·lrotr(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	// TODO: Code missing - could be:
	// ROR R8, R9

	MOVL $0, ret+12(FP)
	RET

// func mayIUseCpuFeature(a uint64) int
TEXT ·mayIUseCpuFeature(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func rdpmc(a int) int64
TEXT ·rdpmc(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// RDPMC R8

	MOVQ $0, ret+8(FP)
	RET

// func rotl(a uint32, shift int) uint32
TEXT ·rotl(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	// TODO: Code missing - could be:
	// ROL R8, R9

	MOVL $0, ret+12(FP)
	RET

// func rotr(a uint32, shift int) uint32
TEXT ·rotr(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	// TODO: Code missing - could be:
	// ROR R8, R9

	MOVL $0, ret+12(FP)
	RET

// func rotwl(a uint16, shift int) uint16
TEXT ·rotwl(SB),7,$0
	MOVW a+0(FP),R8
	MOVQ shift+4(FP),R9

	// TODO: Code missing - could be:
	// ROL R8, R9

	MOVW $0, ret+12(FP)
	RET

// func rotwr(a uint16, shift int) uint16
TEXT ·rotwr(SB),7,$0
	MOVW a+0(FP),R8
	MOVQ shift+4(FP),R9

	// TODO: Code missing - could be:
	// ROR R8, R9

	MOVW $0, ret+12(FP)
	RET

