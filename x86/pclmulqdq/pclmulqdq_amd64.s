// func clmulepi64Si128(a [16]byte, b [16]byte, imm8 byte) [16]byte
TEXT ·clmulepi64Si128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: PCLMULQDQ

	MOVOU X2, ret+36(FP)
	RET

