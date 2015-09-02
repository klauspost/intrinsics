// func m512Exp2a23Pd(a [8]float64) [8]float64
TEXT ·m512Exp2a23Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VEXP2PD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp2a23Pd(a [8]float64, k uint8, src [8]float64) [8]float64
TEXT ·m512MaskExp2a23Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExp2a23Pd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzExp2a23Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VEXP2PD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Exp2a23Ps(a [16]float32) [16]float32
TEXT ·m512Exp2a23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VEXP2PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp2a23Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskExp2a23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExp2a23Ps(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzExp2a23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VEXP2PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Exp2a23RoundPd(a [8]float64, rounding int) [8]float64
TEXT ·m512Exp2a23RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VEXP2PD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskExp2a23RoundPd(a [8]float64, k uint8, src [8]float64, rounding int) [8]float64
TEXT ·m512MaskExp2a23RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzExp2a23RoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskzExp2a23RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Exp2a23RoundPs(a [16]float32, rounding int) [16]float32
TEXT ·m512Exp2a23RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VEXP2PS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskExp2a23RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskExp2a23RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzExp2a23RoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskzExp2a23RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskRcp28Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskRcp28Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRcp28Pd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzRcp28Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VRCP28PD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rcp28Pd(a [8]float64) [8]float64
TEXT ·m512Rcp28Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VRCP28PD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskRcp28Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRcp28Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRcp28Ps(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzRcp28Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VRCP28PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rcp28Ps(a [16]float32) [16]float32
TEXT ·m512Rcp28Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VRCP28PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskRcp28RoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskRcp28RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRcp28RoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskzRcp28RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Rcp28RoundPd(a [8]float64, rounding int) [8]float64
TEXT ·m512Rcp28RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VRCP28PD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskRcp28RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskRcp28RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRcp28RoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskzRcp28RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Rcp28RoundPs(a [16]float32, rounding int) [16]float32
TEXT ·m512Rcp28RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VRCP28PS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRcp28RoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskRcp28RoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzRcp28RoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzRcp28RoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func rcp28RoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·rcp28RoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskRcp28RoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskRcp28RoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzRcp28RoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzRcp28RoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func rcp28RoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·rcp28RoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskRcp28Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRcp28Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRcp28Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRcp28Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rcp28Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rcp28Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VRCP28SD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskRcp28Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRcp28Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRcp28Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRcp28Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rcp28Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rcp28Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VRCP28SS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m512MaskRsqrt28Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskRsqrt28Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRsqrt28Pd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzRsqrt28Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VRSQRT28PD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rsqrt28Pd(a [8]float64) [8]float64
TEXT ·m512Rsqrt28Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VRSQRT28PD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskRsqrt28Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRsqrt28Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRsqrt28Ps(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzRsqrt28Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VRSQRT28PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rsqrt28Ps(a [16]float32) [16]float32
TEXT ·m512Rsqrt28Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VRSQRT28PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskRsqrt28RoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskRsqrt28RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRsqrt28RoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskzRsqrt28RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Rsqrt28RoundPd(a [8]float64, rounding int) [8]float64
TEXT ·m512Rsqrt28RoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing - could be:
	// VRSQRT28PD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskRsqrt28RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskRsqrt28RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRsqrt28RoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskzRsqrt28RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Rsqrt28RoundPs(a [16]float32, rounding int) [16]float32
TEXT ·m512Rsqrt28RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing - could be:
	// VRSQRT28PS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRsqrt28RoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskRsqrt28RoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzRsqrt28RoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzRsqrt28RoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func rsqrt28RoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·rsqrt28RoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskRsqrt28RoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskRsqrt28RoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzRsqrt28RoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzRsqrt28RoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func rsqrt28RoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·rsqrt28RoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskRsqrt28Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRsqrt28Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRsqrt28Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRsqrt28Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rsqrt28Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rsqrt28Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VRSQRT28SD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskRsqrt28Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRsqrt28Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRsqrt28Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRsqrt28Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rsqrt28Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rsqrt28Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VRSQRT28SS X0, X1

	MOVOU X1, ret+32(FP)
	RET

