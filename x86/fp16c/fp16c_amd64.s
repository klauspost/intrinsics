// func cvtphPs(a [16]byte) [4]float32
TEXT ·cvtphPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtphPs1(a [16]byte) [8]float32
TEXT ·cvtphPs1(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtpsPh(a [4]float32, rounding int) [16]byte
TEXT ·cvtpsPh(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtpsPh1(a [8]float32, rounding int) [16]byte
TEXT ·cvtpsPh1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

