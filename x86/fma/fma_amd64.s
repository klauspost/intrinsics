// func fmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmaddPd(a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256FmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmaddPs(a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256FmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmaddSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fmaddSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fmaddsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmaddsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmaddsubPd(a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256FmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmaddsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmaddsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmaddsubPs(a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256FmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmsubPd(a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256FmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmsubPs(a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256FmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmsubSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fmsubaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmsubaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmsubaddPd(a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256FmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmsubaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmsubaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FmsubaddPs(a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256FmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fnmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FnmaddPd(a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256FnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fnmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FnmaddPs(a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256FnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fnmaddSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fnmaddSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fnmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FnmsubPd(a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256FnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fnmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256FnmsubPs(a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256FnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fnmsubSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func fnmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

