// func clmulepi64Si128(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·clmulepi64Si128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

