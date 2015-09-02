// func addcarryxU32(c_in uint8, a uint32, b uint32, out uint32) uint8
TEXT ·addcarryxU32(SB),7,$0
	MOVB c_in+0(FP),R8
	MOVL a+4(FP),R9
	MOVL b+8(FP),R10
	MOVL out+12(FP),R11

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func addcarryxU64(c_in uint8, a uint64, b uint64, out uint64) uint8
TEXT ·addcarryxU64(SB),7,$0
	MOVB c_in+0(FP),R8
	MOVQ a+4(FP),R9
	MOVQ b+12(FP),R10
	MOVQ out+20(FP),R11

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

