// func cvtphPs(a [16]byte) [4]float32
TEXT ·cvtphPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VCVTPH2PS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func m256CvtphPs(a [16]byte) [8]float32
TEXT ·m256CvtphPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VCVTPH2PS X0

	MOV Y0, ret+16(FP)
	RET

// func cvtpsPh(a [4]float32, rounding int) [16]byte
TEXT ·cvtpsPh(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing - could be:
	// VCVTPS2PH X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256CvtpsPh(a [8]float32, rounding int) [16]byte
TEXT ·m256CvtpsPh(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256
	//	MOV a+0(FP),Y0
	MOVQ rounding+32(FP),R9

	// TODO: Code missing - could be:
	// VCVTPS2PH Y0, R9

	MOVOU X1, ret+40(FP)
	RET

