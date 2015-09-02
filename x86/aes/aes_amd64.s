// func aesdecSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesdecSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	// TODO: Code missing - could be:
	// AESDEC X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func aesdeclastSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesdeclastSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	// TODO: Code missing - could be:
	// AESDECLAST X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func aesencSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesencSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	// TODO: Code missing - could be:
	// AESENC X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func aesenclastSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesenclastSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	// TODO: Code missing - could be:
	// AESENCLAST X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func aesimcSi128(a [16]byte) [16]byte
TEXT ·aesimcSi128(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// AESIMC X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func aeskeygenassistSi128(a [16]byte, imm8 int) [16]byte
TEXT ·aeskeygenassistSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing - could be:
	// AESKEYGENASSIST X0, R9

	MOVOU X1, ret+24(FP)
	RET

