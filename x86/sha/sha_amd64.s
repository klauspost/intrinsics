// func sha1msg1Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha1msg1Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// SHA1MSG1 X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func sha1msg2Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha1msg2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// SHA1MSG2 X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func sha1nexteEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha1nexteEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// SHA1NEXTE X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func sha1rnds4Epu32(a [16]byte, b [16]byte, fnc int) [16]byte
TEXT ·sha1rnds4Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ fnc+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func sha256msg1Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha256msg1Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// SHA256MSG1 X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func sha256msg2Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha256msg2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// SHA256MSG2 X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func sha256rnds2Epu32(a [16]byte, b [16]byte, k [16]byte) [16]byte
TEXT ·sha256rnds2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU k+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

