// func maskAndPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAndPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAndPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskAndPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzAndPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func andPd(a [8]float64, b [8]float64) [8]float64
TEXT ·andPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskAndPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAndPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzAndPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAndPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAndPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskAndPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzAndPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func andPs(a [16]float32, b [16]float32) [16]float32
TEXT ·andPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskAndPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAndPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzAndPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndnotPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAndnotPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnotPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAndnotPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndnotPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskAndnotPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndnotPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzAndnotPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func andnotPd(a [8]float64, b [8]float64) [8]float64
TEXT ·andnotPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndnotPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskAndnotPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAndnotPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzAndnotPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndnotPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAndnotPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnotPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAndnotPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndnotPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskAndnotPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndnotPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzAndnotPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func andnotPs(a [16]float32, b [16]float32) [16]float32
TEXT ·andnotPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndnotPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskAndnotPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAndnotPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzAndnotPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func broadcastF32x2(a [4]float32) [8]float32
TEXT ·broadcastF32x2(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskBroadcastF32x2(src [8]float32, k uint8, a [4]float32) [8]float32
TEXT ·maskBroadcastF32x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastF32x2(k uint8, a [4]float32) [8]float32
TEXT ·maskzBroadcastF32x2(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastF32x21(a [4]float32) [16]float32
TEXT ·broadcastF32x21(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastF32x21(src [16]float32, k uint16, a [4]float32) [16]float32
TEXT ·maskBroadcastF32x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastF32x21(k uint16, a [4]float32) [16]float32
TEXT ·maskzBroadcastF32x21(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func broadcastF32x8(a [8]float32) [16]float32
TEXT ·broadcastF32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBroadcastF32x8(src [16]float32, k uint16, a [8]float32) [16]float32
TEXT ·maskBroadcastF32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastF32x8(k uint16, a [8]float32) [16]float32
TEXT ·maskzBroadcastF32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func broadcastF64x2(a [2]float64) [4]float64
TEXT ·broadcastF64x2(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskBroadcastF64x2(src [4]float64, k uint8, a [2]float64) [4]float64
TEXT ·maskBroadcastF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastF64x2(k uint8, a [2]float64) [4]float64
TEXT ·maskzBroadcastF64x2(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastF64x21(a [2]float64) [8]float64
TEXT ·broadcastF64x21(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastF64x21(src [8]float64, k uint8, a [2]float64) [8]float64
TEXT ·maskBroadcastF64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastF64x21(k uint8, a [2]float64) [8]float64
TEXT ·maskzBroadcastF64x21(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func broadcastI32x2(a [16]byte) [16]byte
TEXT ·broadcastI32x2(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskBroadcastI32x2(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastI32x2(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastI32x2(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastI32x2(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func broadcastI32x21(a [16]byte) [32]byte
TEXT ·broadcastI32x21(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskBroadcastI32x21(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskBroadcastI32x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastI32x21(k uint8, a [16]byte) [32]byte
TEXT ·maskzBroadcastI32x21(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastI32x22(a [16]byte) [64]byte
TEXT ·broadcastI32x22(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastI32x22(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·maskBroadcastI32x22(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastI32x22(k uint16, a [16]byte) [64]byte
TEXT ·maskzBroadcastI32x22(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func broadcastI32x8(a [32]byte) [64]byte
TEXT ·broadcastI32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBroadcastI32x8(src [64]byte, k uint16, a [32]byte) [64]byte
TEXT ·maskBroadcastI32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastI32x8(k uint16, a [32]byte) [64]byte
TEXT ·maskzBroadcastI32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func broadcastI64x2(a [16]byte) [32]byte
TEXT ·broadcastI64x2(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskBroadcastI64x2(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskBroadcastI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastI64x2(k uint8, a [16]byte) [32]byte
TEXT ·maskzBroadcastI64x2(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastI64x21(a [16]byte) [64]byte
TEXT ·broadcastI64x21(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastI64x21(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·maskBroadcastI64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastI64x21(k uint8, a [16]byte) [64]byte
TEXT ·maskzBroadcastI64x21(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func cvtRoundepi64Pd(a [64]byte, rounding int) [8]float64
TEXT ·cvtRoundepi64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundepi64Pd(src [8]float64, k uint8, a [64]byte, rounding int) [8]float64
TEXT ·maskCvtRoundepi64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundepi64Pd(k uint8, a [64]byte, rounding int) [8]float64
TEXT ·maskzCvtRoundepi64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundepi64Ps(a [64]byte, rounding int) [8]float32
TEXT ·cvtRoundepi64Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtRoundepi64Ps(src [8]float32, k uint8, a [64]byte, rounding int) [8]float32
TEXT ·maskCvtRoundepi64Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtRoundepi64Ps(k uint8, a [64]byte, rounding int) [8]float32
TEXT ·maskzCvtRoundepi64Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtRoundepu64Pd(a [64]byte, rounding int) [8]float64
TEXT ·cvtRoundepu64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundepu64Pd(src [8]float64, k uint8, a [64]byte, rounding int) [8]float64
TEXT ·maskCvtRoundepu64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundepu64Pd(k uint8, a [64]byte, rounding int) [8]float64
TEXT ·maskzCvtRoundepu64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundepu64Ps(a [64]byte, rounding int) [8]float32
TEXT ·cvtRoundepu64Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtRoundepu64Ps(src [8]float32, k uint8, a [64]byte, rounding int) [8]float32
TEXT ·maskCvtRoundepu64Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtRoundepu64Ps(k uint8, a [64]byte, rounding int) [8]float32
TEXT ·maskzCvtRoundepu64Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtRoundpdEpi64(a [8]float64, rounding int) [64]byte
TEXT ·cvtRoundpdEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpdEpi64(src [64]byte, k uint8, a [8]float64, rounding int) [64]byte
TEXT ·maskCvtRoundpdEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpdEpi64(k uint8, a [8]float64, rounding int) [64]byte
TEXT ·maskzCvtRoundpdEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundpdEpu64(a [8]float64, rounding int) [64]byte
TEXT ·cvtRoundpdEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpdEpu64(src [64]byte, k uint8, a [8]float64, rounding int) [64]byte
TEXT ·maskCvtRoundpdEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpdEpu64(k uint8, a [8]float64, rounding int) [64]byte
TEXT ·maskzCvtRoundpdEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundpsEpi64(a [8]float32, rounding int) [64]byte
TEXT ·cvtRoundpsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpsEpi64(src [64]byte, k uint8, a [8]float32, rounding int) [64]byte
TEXT ·maskCvtRoundpsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpsEpi64(k uint8, a [8]float32, rounding int) [64]byte
TEXT ·maskzCvtRoundpsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundpsEpu64(a [8]float32, rounding int) [64]byte
TEXT ·cvtRoundpsEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpsEpu64(src [64]byte, k uint8, a [8]float32, rounding int) [64]byte
TEXT ·maskCvtRoundpsEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpsEpu64(k uint8, a [8]float32, rounding int) [64]byte
TEXT ·maskzCvtRoundpsEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtepi64Pd(a [16]byte) [2]float64
TEXT ·cvtepi64Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepi64Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepi64Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Pd1(a [32]byte) [4]float64
TEXT ·cvtepi64Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi64Pd1(src [4]float64, k uint8, a [32]byte) [4]float64
TEXT ·maskCvtepi64Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi64Pd1(k uint8, a [32]byte) [4]float64
TEXT ·maskzCvtepi64Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtepi64Pd2(a [64]byte) [8]float64
TEXT ·cvtepi64Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi64Pd2(src [8]float64, k uint8, a [64]byte) [8]float64
TEXT ·maskCvtepi64Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi64Pd2(k uint8, a [64]byte) [8]float64
TEXT ·maskzCvtepi64Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtepi64Ps(a [16]byte) [4]float32
TEXT ·cvtepi64Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Ps(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtepi64Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Ps(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtepi64Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Ps1(a [32]byte) [4]float32
TEXT ·cvtepi64Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64Ps1(src [4]float32, k uint8, a [32]byte) [4]float32
TEXT ·maskCvtepi64Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi64Ps1(k uint8, a [32]byte) [4]float32
TEXT ·maskzCvtepi64Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi64Ps2(a [64]byte) [8]float32
TEXT ·cvtepi64Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi64Ps2(src [8]float32, k uint8, a [64]byte) [8]float32
TEXT ·maskCvtepi64Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi64Ps2(k uint8, a [64]byte) [8]float32
TEXT ·maskzCvtepi64Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtepu64Pd(a [16]byte) [2]float64
TEXT ·cvtepu64Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu64Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepu64Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu64Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepu64Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu64Pd1(a [32]byte) [4]float64
TEXT ·cvtepu64Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepu64Pd1(src [4]float64, k uint8, a [32]byte) [4]float64
TEXT ·maskCvtepu64Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu64Pd1(k uint8, a [32]byte) [4]float64
TEXT ·maskzCvtepu64Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtepu64Pd2(a [64]byte) [8]float64
TEXT ·cvtepu64Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu64Pd2(src [8]float64, k uint8, a [64]byte) [8]float64
TEXT ·maskCvtepu64Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu64Pd2(k uint8, a [64]byte) [8]float64
TEXT ·maskzCvtepu64Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtepu64Ps(a [16]byte) [4]float32
TEXT ·cvtepu64Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu64Ps(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtepu64Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu64Ps(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtepu64Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu64Ps1(a [32]byte) [4]float32
TEXT ·cvtepu64Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepu64Ps1(src [4]float32, k uint8, a [32]byte) [4]float32
TEXT ·maskCvtepu64Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepu64Ps1(k uint8, a [32]byte) [4]float32
TEXT ·maskzCvtepu64Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepu64Ps2(a [64]byte) [8]float32
TEXT ·cvtepu64Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepu64Ps2(src [8]float32, k uint8, a [64]byte) [8]float32
TEXT ·maskCvtepu64Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu64Ps2(k uint8, a [64]byte) [8]float32
TEXT ·maskzCvtepu64Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtpdEpi64(a [2]float64) [16]byte
TEXT ·cvtpdEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpi64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpi64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpdEpi641(a [4]float64) [32]byte
TEXT ·cvtpdEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpdEpi641(src [32]byte, k uint8, a [4]float64) [32]byte
TEXT ·maskCvtpdEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpdEpi641(k uint8, a [4]float64) [32]byte
TEXT ·maskzCvtpdEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtpdEpi642(a [8]float64) [64]byte
TEXT ·cvtpdEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpdEpi642(src [64]byte, k uint8, a [8]float64) [64]byte
TEXT ·maskCvtpdEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpdEpi642(k uint8, a [8]float64) [64]byte
TEXT ·maskzCvtpdEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtpdEpu64(a [2]float64) [16]byte
TEXT ·cvtpdEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpu64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpu64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpdEpu641(a [4]float64) [32]byte
TEXT ·cvtpdEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpdEpu641(src [32]byte, k uint8, a [4]float64) [32]byte
TEXT ·maskCvtpdEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpdEpu641(k uint8, a [4]float64) [32]byte
TEXT ·maskzCvtpdEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtpdEpu642(a [8]float64) [64]byte
TEXT ·cvtpdEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpdEpu642(src [64]byte, k uint8, a [8]float64) [64]byte
TEXT ·maskCvtpdEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpdEpu642(k uint8, a [8]float64) [64]byte
TEXT ·maskzCvtpdEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtpsEpi64(a [4]float32) [16]byte
TEXT ·cvtpsEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpi64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpi64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpsEpi641(a [4]float32) [32]byte
TEXT ·cvtpsEpi641(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskCvtpsEpi641(src [32]byte, k uint8, a [4]float32) [32]byte
TEXT ·maskCvtpsEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpsEpi641(k uint8, a [4]float32) [32]byte
TEXT ·maskzCvtpsEpi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtpsEpi642(a [8]float32) [64]byte
TEXT ·cvtpsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsEpi642(src [64]byte, k uint8, a [8]float32) [64]byte
TEXT ·maskCvtpsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpsEpi642(k uint8, a [8]float32) [64]byte
TEXT ·maskzCvtpsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtpsEpu64(a [4]float32) [16]byte
TEXT ·cvtpsEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpu64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpu64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpsEpu641(a [4]float32) [32]byte
TEXT ·cvtpsEpu641(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskCvtpsEpu641(src [32]byte, k uint8, a [4]float32) [32]byte
TEXT ·maskCvtpsEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpsEpu641(k uint8, a [4]float32) [32]byte
TEXT ·maskzCvtpsEpu641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtpsEpu642(a [8]float32) [64]byte
TEXT ·cvtpsEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsEpu642(src [64]byte, k uint8, a [8]float32) [64]byte
TEXT ·maskCvtpsEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpsEpu642(k uint8, a [8]float32) [64]byte
TEXT ·maskzCvtpsEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttRoundpdEpi64(a [8]float64, sae int) [64]byte
TEXT ·cvttRoundpdEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttRoundpdEpi64(src [64]byte, k uint8, a [8]float64, sae int) [64]byte
TEXT ·maskCvttRoundpdEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttRoundpdEpi64(k uint8, a [8]float64, sae int) [64]byte
TEXT ·maskzCvttRoundpdEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttRoundpdEpu64(a [8]float64, sae int) [64]byte
TEXT ·cvttRoundpdEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttRoundpdEpu64(src [64]byte, k uint8, a [8]float64, sae int) [64]byte
TEXT ·maskCvttRoundpdEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttRoundpdEpu64(k uint8, a [8]float64, sae int) [64]byte
TEXT ·maskzCvttRoundpdEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttRoundpsEpi64(a [8]float32, sae int) [64]byte
TEXT ·cvttRoundpsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttRoundpsEpi64(src [64]byte, k uint8, a [8]float32, sae int) [64]byte
TEXT ·maskCvttRoundpsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttRoundpsEpi64(k uint8, a [8]float32, sae int) [64]byte
TEXT ·maskzCvttRoundpsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttRoundpsEpu64(a [8]float32, sae int) [64]byte
TEXT ·cvttRoundpsEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttRoundpsEpu64(src [64]byte, k uint8, a [8]float32, sae int) [64]byte
TEXT ·maskCvttRoundpsEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttRoundpsEpu64(k uint8, a [8]float32, sae int) [64]byte
TEXT ·maskzCvttRoundpsEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttpdEpi64(a [2]float64) [16]byte
TEXT ·cvttpdEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpi64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpi64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpdEpi641(a [4]float64) [32]byte
TEXT ·cvttpdEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttpdEpi641(src [32]byte, k uint8, a [4]float64) [32]byte
TEXT ·maskCvttpdEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpdEpi641(k uint8, a [4]float64) [32]byte
TEXT ·maskzCvttpdEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttpdEpi642(a [8]float64) [64]byte
TEXT ·cvttpdEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttpdEpi642(src [64]byte, k uint8, a [8]float64) [64]byte
TEXT ·maskCvttpdEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttpdEpi642(k uint8, a [8]float64) [64]byte
TEXT ·maskzCvttpdEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttpdEpu64(a [2]float64) [16]byte
TEXT ·cvttpdEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpu64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpu64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpdEpu641(a [4]float64) [32]byte
TEXT ·cvttpdEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttpdEpu641(src [32]byte, k uint8, a [4]float64) [32]byte
TEXT ·maskCvttpdEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpdEpu641(k uint8, a [4]float64) [32]byte
TEXT ·maskzCvttpdEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttpdEpu642(a [8]float64) [64]byte
TEXT ·cvttpdEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttpdEpu642(src [64]byte, k uint8, a [8]float64) [64]byte
TEXT ·maskCvttpdEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttpdEpu642(k uint8, a [8]float64) [64]byte
TEXT ·maskzCvttpdEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttpsEpi64(a [4]float32) [16]byte
TEXT ·cvttpsEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpi64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpi64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpsEpi641(a [4]float32) [32]byte
TEXT ·cvttpsEpi641(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskCvttpsEpi641(src [32]byte, k uint8, a [4]float32) [32]byte
TEXT ·maskCvttpsEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpsEpi641(k uint8, a [4]float32) [32]byte
TEXT ·maskzCvttpsEpi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvttpsEpi642(a [8]float32) [64]byte
TEXT ·cvttpsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttpsEpi642(src [64]byte, k uint8, a [8]float32) [64]byte
TEXT ·maskCvttpsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttpsEpi642(k uint8, a [8]float32) [64]byte
TEXT ·maskzCvttpsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttpsEpu64(a [4]float32) [16]byte
TEXT ·cvttpsEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpu64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpu64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpsEpu641(a [4]float32) [32]byte
TEXT ·cvttpsEpu641(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskCvttpsEpu641(src [32]byte, k uint8, a [4]float32) [32]byte
TEXT ·maskCvttpsEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpsEpu641(k uint8, a [4]float32) [32]byte
TEXT ·maskzCvttpsEpu641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvttpsEpu642(a [8]float32) [64]byte
TEXT ·cvttpsEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttpsEpu642(src [64]byte, k uint8, a [8]float32) [64]byte
TEXT ·maskCvttpsEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttpsEpu642(k uint8, a [8]float32) [64]byte
TEXT ·maskzCvttpsEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func extractf32x8Ps(a [16]float32, imm8 int) [8]float32
TEXT ·extractf32x8Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExtractf32x8Ps(src [8]float32, k uint8, a [16]float32, imm8 int) [8]float32
TEXT ·maskExtractf32x8Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExtractf32x8Ps(k uint8, a [16]float32, imm8 int) [8]float32
TEXT ·maskzExtractf32x8Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func extractf64x2Pd(a [4]float64, imm8 int) [2]float64
TEXT ·extractf64x2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtractf64x2Pd(src [2]float64, k uint8, a [4]float64, imm8 int) [2]float64
TEXT ·maskExtractf64x2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtractf64x2Pd(k uint8, a [4]float64, imm8 int) [2]float64
TEXT ·maskzExtractf64x2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extractf64x2Pd1(a [8]float64, imm8 int) [2]float64
TEXT ·extractf64x2Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtractf64x2Pd1(src [2]float64, k uint8, a [8]float64, imm8 int) [2]float64
TEXT ·maskExtractf64x2Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtractf64x2Pd1(k uint8, a [8]float64, imm8 int) [2]float64
TEXT ·maskzExtractf64x2Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extracti32x8Epi32(a [64]byte, imm8 int) [32]byte
TEXT ·extracti32x8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExtracti32x8Epi32(src [32]byte, k uint8, a [64]byte, imm8 int) [32]byte
TEXT ·maskExtracti32x8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExtracti32x8Epi32(k uint8, a [64]byte, imm8 int) [32]byte
TEXT ·maskzExtracti32x8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func extracti64x2Epi64(a [32]byte, imm8 int) [16]byte
TEXT ·extracti64x2Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtracti64x2Epi64(src [16]byte, k uint8, a [32]byte, imm8 int) [16]byte
TEXT ·maskExtracti64x2Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtracti64x2Epi64(k uint8, a [32]byte, imm8 int) [16]byte
TEXT ·maskzExtracti64x2Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extracti64x2Epi641(a [64]byte, imm8 int) [16]byte
TEXT ·extracti64x2Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtracti64x2Epi641(src [16]byte, k uint8, a [64]byte, imm8 int) [16]byte
TEXT ·maskExtracti64x2Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtracti64x2Epi641(k uint8, a [64]byte, imm8 int) [16]byte
TEXT ·maskzExtracti64x2Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func fpclassPdMask(a [2]float64, imm8 int) uint8
TEXT ·fpclassPdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassPdMask(k1 uint8, a [2]float64, imm8 int) uint8
TEXT ·maskFpclassPdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func fpclassPdMask1(a [4]float64, imm8 int) uint8
TEXT ·fpclassPdMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskFpclassPdMask1(k1 uint8, a [4]float64, imm8 int) uint8
TEXT ·maskFpclassPdMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func fpclassPdMask2(a [8]float64, imm8 int) uint8
TEXT ·fpclassPdMask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskFpclassPdMask2(k1 uint8, a [8]float64, imm8 int) uint8
TEXT ·maskFpclassPdMask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func fpclassPsMask(a [4]float32, imm8 int) uint8
TEXT ·fpclassPsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassPsMask(k1 uint8, a [4]float32, imm8 int) uint8
TEXT ·maskFpclassPsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func fpclassPsMask1(a [8]float32, imm8 int) uint8
TEXT ·fpclassPsMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskFpclassPsMask1(k1 uint8, a [8]float32, imm8 int) uint8
TEXT ·maskFpclassPsMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func fpclassPsMask2(a [16]float32, imm8 int) uint16
TEXT ·fpclassPsMask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskFpclassPsMask2(k1 uint16, a [16]float32, imm8 int) uint16
TEXT ·maskFpclassPsMask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func fpclassSdMask(a [2]float64, imm8 int) uint8
TEXT ·fpclassSdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassSdMask(k1 uint8, a [2]float64, imm8 int) uint8
TEXT ·maskFpclassSdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func fpclassSsMask(a [4]float32, imm8 int) uint8
TEXT ·fpclassSsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassSsMask(k1 uint8, a [4]float32, imm8 int) uint8
TEXT ·maskFpclassSsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func insertf32x8(a [16]float32, b [8]float32, imm8 int) [16]float32
TEXT ·insertf32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInsertf32x8(src [16]float32, k uint16, a [16]float32, b [8]float32, imm8 int) [16]float32
TEXT ·maskInsertf32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInsertf32x8(k uint16, a [16]float32, b [8]float32, imm8 int) [16]float32
TEXT ·maskzInsertf32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func insertf64x2(a [4]float64, b [2]float64, imm8 int) [4]float64
TEXT ·insertf64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskInsertf64x2(src [4]float64, k uint8, a [4]float64, b [2]float64, imm8 int) [4]float64
TEXT ·maskInsertf64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzInsertf64x2(k uint8, a [4]float64, b [2]float64, imm8 int) [4]float64
TEXT ·maskzInsertf64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func insertf64x21(a [8]float64, b [2]float64, imm8 int) [8]float64
TEXT ·insertf64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInsertf64x21(src [8]float64, k uint8, a [8]float64, b [2]float64, imm8 int) [8]float64
TEXT ·maskInsertf64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInsertf64x21(k uint8, a [8]float64, b [2]float64, imm8 int) [8]float64
TEXT ·maskzInsertf64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func inserti32x8(a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·inserti32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInserti32x8(src [64]byte, k uint16, a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·maskInserti32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInserti32x8(k uint16, a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·maskzInserti32x8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func inserti64x2(a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·inserti64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskInserti64x2(src [32]byte, k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·maskInserti64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzInserti64x2(k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·maskzInserti64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func inserti64x21(a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·inserti64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInserti64x21(src [64]byte, k uint8, a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·maskInserti64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInserti64x21(k uint8, a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·maskzInserti64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func movepi32Mask(a [16]byte) uint8
TEXT ·movepi32Mask(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func movepi32Mask1(a [32]byte) uint8
TEXT ·movepi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func movepi32Mask2(a [64]byte) uint16
TEXT ·movepi32Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func movepi64Mask(a [16]byte) uint8
TEXT ·movepi64Mask(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func movepi64Mask1(a [32]byte) uint8
TEXT ·movepi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func movepi64Mask2(a [64]byte) uint8
TEXT ·movepi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func movmEpi32(k uint8) [16]byte
TEXT ·movmEpi32(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movmEpi321(k uint8) [32]byte
TEXT ·movmEpi321(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func movmEpi322(k uint16) [64]byte
TEXT ·movmEpi322(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func movmEpi64(k uint8) [16]byte
TEXT ·movmEpi64(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movmEpi641(k uint8) [32]byte
TEXT ·movmEpi641(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func movmEpi642(k uint8) [64]byte
TEXT ·movmEpi642(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func maskMulloEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulloEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulloEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulloEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mulloEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·mulloEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMulloEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulloEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulloEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulloEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulloEpi641(a [32]byte, b [32]byte) [32]byte
TEXT ·mulloEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulloEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulloEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulloEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulloEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulloEpi642(a [64]byte, b [64]byte) [64]byte
TEXT ·mulloEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskOrPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskOrPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOrPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzOrPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskOrPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskOrPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzOrPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzOrPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskOrPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskOrPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzOrPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzOrPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func orPd(a [8]float64, b [8]float64) [8]float64
TEXT ·orPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskOrPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskOrPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOrPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzOrPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskOrPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskOrPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzOrPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzOrPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskOrPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskOrPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzOrPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzOrPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func orPs(a [16]float32, b [16]float32) [16]float32
TEXT ·orPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRangePd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskRangePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangePd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzRangePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rangePd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·rangePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRangePd1(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·maskRangePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRangePd1(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·maskzRangePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rangePd1(a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·rangePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRangePd2(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·maskRangePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRangePd2(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·maskzRangePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rangePd2(a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·rangePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRangePs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskRangePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangePs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzRangePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rangePs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·rangePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRangePs1(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·maskRangePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRangePs1(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·maskzRangePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rangePs1(a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·rangePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRangePs2(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·maskRangePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRangePs2(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·maskzRangePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rangePs2(a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·rangePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRangeRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64
TEXT ·maskRangeRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRangeRoundPd(k uint8, a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64
TEXT ·maskzRangeRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rangeRoundPd(a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64
TEXT ·rangeRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRangeRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32
TEXT ·maskRangeRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRangeRoundPs(k uint16, a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32
TEXT ·maskzRangeRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rangeRoundPs(a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32
TEXT ·rangeRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRangeRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskRangeRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzRangeRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskzRangeRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func rangeRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·rangeRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskRangeRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskRangeRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzRangeRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskzRangeRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func rangeRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·rangeRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskRangeSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskRangeSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangeSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzRangeSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskRangeSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskRangeSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangeSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzRangeSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskReducePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskReducePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzReducePd(k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskzReducePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func reducePd(a [2]float64, imm8 int) [2]float64
TEXT ·reducePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskReducePd1(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskReducePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzReducePd1(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskzReducePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func reducePd1(a [4]float64, imm8 int) [4]float64
TEXT ·reducePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskReducePd2(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskReducePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzReducePd2(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskzReducePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func reducePd2(a [8]float64, imm8 int) [8]float64
TEXT ·reducePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskReducePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskReducePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzReducePs(k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskzReducePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func reducePs(a [4]float32, imm8 int) [4]float32
TEXT ·reducePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskReducePs1(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·maskReducePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzReducePs1(k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·maskzReducePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func reducePs1(a [8]float32, imm8 int) [8]float32
TEXT ·reducePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskReducePs2(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·maskReducePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzReducePs2(k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·maskzReducePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func reducePs2(a [16]float32, imm8 int) [16]float32
TEXT ·reducePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskReduceRoundPd(src [8]float64, k uint8, a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·maskReduceRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzReduceRoundPd(k uint8, a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·maskzReduceRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func reduceRoundPd(a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·reduceRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskReduceRoundPs(src [16]float32, k uint16, a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·maskReduceRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzReduceRoundPs(k uint16, a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·maskzReduceRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func reduceRoundPs(a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·reduceRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskReduceRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskReduceRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzReduceRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskzReduceRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func reduceRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·reduceRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskReduceRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskReduceRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzReduceRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskzReduceRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func reduceRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·reduceRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskReduceSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskReduceSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzReduceSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzReduceSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func reduceSd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·reduceSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskReduceSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskReduceSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzReduceSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzReduceSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func reduceSs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·reduceSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskXorPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskXorPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXorPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzXorPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskXorPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskXorPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzXorPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzXorPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskXorPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskXorPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzXorPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzXorPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func xorPd(a [8]float64, b [8]float64) [8]float64
TEXT ·xorPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskXorPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskXorPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXorPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzXorPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskXorPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskXorPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzXorPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzXorPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskXorPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskXorPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzXorPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzXorPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func xorPs(a [16]float32, b [16]float32) [16]float32
TEXT ·xorPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

