// func maskAbsEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskAbsEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzAbsEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzAbsEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPABSD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskAbsEpi32(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskAbsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzAbsEpi32(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzAbsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPABSD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512AbsEpi32(a [64]byte) [64]byte
TEXT ·m512AbsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPABSD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAbsEpi32(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·m512MaskAbsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzAbsEpi32(k uint16, a [64]byte) [64]byte
TEXT ·m512MaskzAbsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPABSD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func absEpi64(a [16]byte) [16]byte
TEXT ·absEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPABSQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskAbsEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskAbsEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzAbsEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzAbsEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPABSQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256AbsEpi64(a [32]byte) [32]byte
TEXT ·m256AbsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPABSQ Y0

	MOV Y0, ret+0(FP)
	RET

// func m256MaskAbsEpi64(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskAbsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzAbsEpi64(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzAbsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPABSQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512AbsEpi64(a [64]byte) [64]byte
TEXT ·m512AbsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPABSQ Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAbsEpi64(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·m512MaskAbsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzAbsEpi64(k uint8, a [64]byte) [64]byte
TEXT ·m512MaskzAbsEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPABSQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512AcosPd(a [8]float64) [8]float64
TEXT ·m512AcosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAcosPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskAcosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AcosPs(a [16]float32) [16]float32
TEXT ·m512AcosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAcosPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskAcosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AcoshPd(a [8]float64) [8]float64
TEXT ·m512AcoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAcoshPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskAcoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AcoshPs(a [16]float32) [16]float32
TEXT ·m512AcoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAcoshPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskAcoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskAddEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAddEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzAddEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskAddEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAddEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512AddEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPADDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAddEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzAddEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzAddPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzAddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzAddPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzAddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzAddRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskzAddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzAddRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskzAddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func addRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·addRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskAddRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskAddRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzAddRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzAddRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func addRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·addRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskAddRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskAddRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzAddRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzAddRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskAddSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAddSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAddSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskAddSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAddSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAddSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m512MaskzAlignrEpi32(k uint16, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskzAlignrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AlignrEpi64(a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512AlignrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAlignrEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskAlignrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzAlignrEpi64(k uint8, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskzAlignrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskAndEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAndEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAndEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAndEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzAndEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskAndEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAndEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAndEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAndEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzAndEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskAndnotEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndnotEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAndnotEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndnotEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAndnotEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAndnotEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAndnotEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAndnotEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzAndnotEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAndnotEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskAndnotEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndnotEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzAndnotEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndnotEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAndnotEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAndnotEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAndnotEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAndnotEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzAndnotEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAndnotEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AsinPd(a [8]float64) [8]float64
TEXT ·m512AsinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAsinPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskAsinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AsinPs(a [16]float32) [16]float32
TEXT ·m512AsinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAsinPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskAsinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AsinhPd(a [8]float64) [8]float64
TEXT ·m512AsinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAsinhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskAsinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AsinhPs(a [16]float32) [16]float32
TEXT ·m512AsinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAsinhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskAsinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AtanPd(a [8]float64) [8]float64
TEXT ·m512AtanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAtanPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskAtanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AtanPs(a [16]float32) [16]float32
TEXT ·m512AtanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAtanPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskAtanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Atan2Pd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512Atan2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAtan2Pd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskAtan2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Atan2Ps(a [16]float32, b [16]float32) [16]float32
TEXT ·m512Atan2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAtan2Ps(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskAtan2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AtanhPd(a [8]float64) [8]float64
TEXT ·m512AtanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAtanhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskAtanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AtanhPs(a [16]float32) [16]float32
TEXT ·m512AtanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAtanhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskAtanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskBlendEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskBlendEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskBlendEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func maskBlendEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskBlendEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskBlendEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func maskBlendPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskBlendPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskBlendPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskBlendPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func maskBlendPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskBlendPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskBlendPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskBlendPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256BroadcastF32x4(a [4]float32) [8]float32
TEXT ·m256BroadcastF32x4(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VBROADCASTF32X4 X0

	MOV Y0, ret+16(FP)
	RET

// func m256MaskBroadcastF32x4(src [8]float32, k uint8, a [4]float32) [8]float32
TEXT ·m256MaskBroadcastF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzBroadcastF32x4(k uint8, a [4]float32) [8]float32
TEXT ·m256MaskzBroadcastF32x4(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTF32X4 R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastF32x4(a [4]float32) [16]float32
TEXT ·m512BroadcastF32x4(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VBROADCASTF32X4 X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastF32x4(src [16]float32, k uint16, a [4]float32) [16]float32
TEXT ·m512MaskBroadcastF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastF32x4(k uint16, a [4]float32) [16]float32
TEXT ·m512MaskzBroadcastF32x4(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTF32X4 R8, X1

	MOV Z1, ret+20(FP)
	RET

// func m512BroadcastF64x4(a [4]float64) [8]float64
TEXT ·m512BroadcastF64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VBROADCASTF64X4 Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskBroadcastF64x4(src [8]float64, k uint8, a [4]float64) [8]float64
TEXT ·m512MaskBroadcastF64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastF64x4(k uint8, a [4]float64) [8]float64
TEXT ·m512MaskzBroadcastF64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VBROADCASTF64X4 R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func m256BroadcastI32x4(a [16]byte) [32]byte
TEXT ·m256BroadcastI32x4(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VBROADCASTI32X4 X0

	MOV Y0, ret+16(FP)
	RET

// func m256MaskBroadcastI32x4(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzBroadcastI32x4(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastI32x4(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTI32X4 R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastI32x4(a [16]byte) [64]byte
TEXT ·m512BroadcastI32x4(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VBROADCASTI32X4 X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastI32x4(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastI32x4(k uint16, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastI32x4(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTI32X4 R8, X1

	MOV Z1, ret+20(FP)
	RET

// func m512BroadcastI64x4(a [32]byte) [64]byte
TEXT ·m512BroadcastI64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VBROADCASTI64X4 Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskBroadcastI64x4(src [64]byte, k uint8, a [32]byte) [64]byte
TEXT ·m512MaskBroadcastI64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastI64x4(k uint8, a [32]byte) [64]byte
TEXT ·m512MaskzBroadcastI64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VBROADCASTI64X4 R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func maskBroadcastdEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastdEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzBroadcastdEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskBroadcastdEpi32(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzBroadcastdEpi32(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastdEpi32(a [16]byte) [64]byte
TEXT ·m512BroadcastdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastdEpi32(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastdEpi32(k uint16, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastdEpi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8, X1

	MOV Z1, ret+20(FP)
	RET

// func maskBroadcastqEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastqEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzBroadcastqEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastqEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskBroadcastqEpi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastqEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzBroadcastqEpi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastqEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastqEpi64(a [16]byte) [64]byte
TEXT ·m512BroadcastqEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastqEpi64(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastqEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastqEpi64(k uint8, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastqEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8, X1

	MOV Z1, ret+20(FP)
	RET

// func m256MaskBroadcastsdPd(src [4]float64, k uint8, a [2]float64) [4]float64
TEXT ·m256MaskBroadcastsdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzBroadcastsdPd(k uint8, a [2]float64) [4]float64
TEXT ·m256MaskzBroadcastsdPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTSD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastsdPd(a [2]float64) [8]float64
TEXT ·m512BroadcastsdPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VBROADCASTSD X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastsdPd(src [8]float64, k uint8, a [2]float64) [8]float64
TEXT ·m512MaskBroadcastsdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastsdPd(k uint8, a [2]float64) [8]float64
TEXT ·m512MaskzBroadcastsdPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTSD R8, X1

	MOV Z1, ret+20(FP)
	RET

// func maskBroadcastssPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskBroadcastssPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzBroadcastssPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzBroadcastssPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTSS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskBroadcastssPs(src [8]float32, k uint8, a [4]float32) [8]float32
TEXT ·m256MaskBroadcastssPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzBroadcastssPs(k uint8, a [4]float32) [8]float32
TEXT ·m256MaskzBroadcastssPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTSS R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastssPs(a [4]float32) [16]float32
TEXT ·m512BroadcastssPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VBROADCASTSS X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastssPs(src [16]float32, k uint16, a [4]float32) [16]float32
TEXT ·m512MaskBroadcastssPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzBroadcastssPs(k uint16, a [4]float32) [16]float32
TEXT ·m512MaskzBroadcastssPs(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VBROADCASTSS R8, X1

	MOV Z1, ret+20(FP)
	RET

// func m512Castpd128Pd512(a [2]float64) [8]float64
TEXT ·m512Castpd128Pd512(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func m512Castpd256Pd512(a [4]float64) [8]float64
TEXT ·m512Castpd256Pd512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Castpd512Pd128(a [8]float64) [2]float64
TEXT ·m512Castpd512Pd128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m512Castpd512Pd256(a [8]float64) [4]float64
TEXT ·m512Castpd512Pd256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512Castps128Ps512(a [4]float32) [16]float32
TEXT ·m512Castps128Ps512(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func m512Castps256Ps512(a [8]float32) [16]float32
TEXT ·m512Castps256Ps512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Castps512Ps128(a [16]float32) [4]float32
TEXT ·m512Castps512Ps128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m512Castps512Ps256(a [16]float32) [8]float32
TEXT ·m512Castps512Ps256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512Castsi128Si512(a [16]byte) [64]byte
TEXT ·m512Castsi128Si512(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func m512Castsi256Si512(a [32]byte) [64]byte
TEXT ·m512Castsi256Si512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Castsi512Si128(a [64]byte) [16]byte
TEXT ·m512Castsi512Si128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m512Castsi512Si256(a [64]byte) [32]byte
TEXT ·m512Castsi512Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512CbrtPd(a [8]float64) [8]float64
TEXT ·m512CbrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCbrtPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCbrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CbrtPs(a [16]float32) [16]float32
TEXT ·m512CbrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCbrtPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCbrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CdfnormPd(a [8]float64) [8]float64
TEXT ·m512CdfnormPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCdfnormPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCdfnormPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CdfnormPs(a [16]float32) [16]float32
TEXT ·m512CdfnormPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCdfnormPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCdfnormPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CdfnorminvPd(a [8]float64) [8]float64
TEXT ·m512CdfnorminvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCdfnorminvPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCdfnorminvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CdfnorminvPs(a [16]float32) [16]float32
TEXT ·m512CdfnorminvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCdfnorminvPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCdfnorminvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CeilPd(a [8]float64) [8]float64
TEXT ·m512CeilPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCeilPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCeilPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CeilPs(a [16]float32) [16]float32
TEXT ·m512CeilPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCeilPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCeilPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func cmpEpi32Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpi32Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func m256CmpEpi32Mask(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256CmpEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpEpi32Mask(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256MaskCmpEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpi64Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpi64Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func m256CmpEpi64Mask(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256CmpEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpEpi64Mask(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256MaskCmpEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpEpi64Mask(a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·m512CmpEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpEpi64Mask(k1 uint8, a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·m512MaskCmpEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpu32Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu32Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func m256CmpEpu32Mask(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256CmpEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpEpu32Mask(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256MaskCmpEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpu64Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu64Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func m256CmpEpu64Mask(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256CmpEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpEpu64Mask(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·m256MaskCmpEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpEpu64Mask(a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·m512CmpEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpEpu64Mask(k1 uint8, a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·m512MaskCmpEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpPdMask(a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·cmpPdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpPdMask(k1 uint8, a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·maskCmpPdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func m256CmpPdMask(a [4]float64, b [4]float64, imm8 int) uint8
TEXT ·m256CmpPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpPdMask(k1 uint8, a [4]float64, b [4]float64, imm8 int) uint8
TEXT ·m256MaskCmpPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpPsMask(a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·cmpPsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpPsMask(k1 uint8, a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·maskCmpPsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func m256CmpPsMask(a [8]float32, b [8]float32, imm8 int) uint8
TEXT ·m256CmpPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpPsMask(k1 uint8, a [8]float32, b [8]float32, imm8 int) uint8
TEXT ·m256MaskCmpPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpRoundSdMask(a [2]float64, b [2]float64, imm8 int, sae int) uint8
TEXT ·cmpRoundSdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	// TODO: Code missing

	MOVB $0, ret+48(FP)
	RET

// func maskCmpRoundSdMask(k1 uint8, a [2]float64, b [2]float64, imm8 int, sae int) uint8
TEXT ·maskCmpRoundSdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ sae+44(FP),R12

	// TODO: Code missing

	MOVB $0, ret+52(FP)
	RET

// func cmpRoundSsMask(a [4]float32, b [4]float32, imm8 int, sae int) uint8
TEXT ·cmpRoundSsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	// TODO: Code missing

	MOVB $0, ret+48(FP)
	RET

// func maskCmpRoundSsMask(k1 uint8, a [4]float32, b [4]float32, imm8 int, sae int) uint8
TEXT ·maskCmpRoundSsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ sae+44(FP),R12

	// TODO: Code missing

	MOVB $0, ret+52(FP)
	RET

// func cmpSdMask(a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·cmpSdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpSdMask(k1 uint8, a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·maskCmpSdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpSsMask(a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·cmpSsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpSsMask(k1 uint8, a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·maskCmpSsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpeqEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpeqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpeqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpeqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpeqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpeqEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpeqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPEQQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpeqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpu32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpeqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpu32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpeqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpu64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpeqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpu64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpeqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpeqEpu64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpeqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpu64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpeqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgeEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgeEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgeEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgeEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpgeEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpgeEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpgeEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpu32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgeEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpu32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgeEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpu64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgeEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpu64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgeEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpgeEpu64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpgeEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpu64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpgeEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgtEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgtEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgtEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgtEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpgtEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpgtEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPGTQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpgtEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpu32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgtEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpu32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgtEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpu64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpgtEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpu64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpgtEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpgtEpu64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpgtEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpu64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpgtEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpleEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpleEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpleEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpleEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpleEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpleEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpleEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpleEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpleEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpleEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpu32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpleEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpleEpu32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpleEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpu64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpleEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpleEpu64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpleEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpleEpu64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpleEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpleEpu64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpleEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpltEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpltEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpltEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpltEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpltEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpltEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpltEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpltEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpltEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpltEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpltEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpltEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpltEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpltEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpltEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpu32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpltEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpltEpu32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpltEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpu64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpltEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpltEpu64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpltEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpltEpu64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpltEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpltEpu64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpltEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpneqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpneqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpneqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpneqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpneqEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpneqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpneqEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpu32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpneqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpu32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpneqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCMPUQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpu64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256CmpneqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpu64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskCmpneqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpneqEpu64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512CmpneqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpu64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskCmpneqEpu64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func comiRoundSd(a [2]float64, b [2]float64, imm8 int, sae int) int
TEXT ·comiRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	// TODO: Code missing

	MOVQ $0, ret+48(FP)
	RET

// func comiRoundSs(a [4]float32, b [4]float32, imm8 int, sae int) int
TEXT ·comiRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	// TODO: Code missing

	MOVQ $0, ret+48(FP)
	RET

// func maskCompressEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCompressEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCompressEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCompressEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCOMPRESSD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCompressEpi32(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskCompressEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCompressEpi32(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzCompressEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCOMPRESSD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCompressEpi32(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·m512MaskCompressEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCompressEpi32(k uint16, a [64]byte) [64]byte
TEXT ·m512MaskzCompressEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCOMPRESSD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskCompressEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCompressEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCompressEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCompressEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPCOMPRESSQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCompressEpi64(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskCompressEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCompressEpi64(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzCompressEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPCOMPRESSQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCompressEpi64(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·m512MaskCompressEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCompressEpi64(k uint8, a [64]byte) [64]byte
TEXT ·m512MaskzCompressEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCOMPRESSQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskCompressPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskCompressPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCompressPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzCompressPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCOMPRESSPD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCompressPd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskCompressPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCompressPd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzCompressPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCOMPRESSPD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCompressPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCompressPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCompressPd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzCompressPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCOMPRESSPD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskCompressPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskCompressPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCompressPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzCompressPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCOMPRESSPS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCompressPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskCompressPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCompressPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzCompressPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCOMPRESSPS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCompressPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCompressPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCompressPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzCompressPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCOMPRESSPS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskCompressstoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCompressstoreuEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCompressstoreuEpi32(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCompressstoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCompressstoreuEpi32(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCompressstoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCompressstoreuEpi64(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCompressstoreuEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCompressstoreuEpi64(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCompressstoreuEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCompressstoreuEpi64(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCompressstoreuEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCompressstoreuPd(base_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskCompressstoreuPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCompressstoreuPd(base_addr uintptr, k uint8, a [4]float64) 
TEXT ·m256MaskCompressstoreuPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	RET

// func m512MaskCompressstoreuPd(base_addr uintptr, k uint8, a [8]float64) 
TEXT ·m512MaskCompressstoreuPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func maskCompressstoreuPs(base_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskCompressstoreuPs(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCompressstoreuPs(base_addr uintptr, k uint8, a [8]float32) 
TEXT ·m256MaskCompressstoreuPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	RET

// func m512MaskCompressstoreuPs(base_addr uintptr, k uint16, a [16]float32) 
TEXT ·m512MaskCompressstoreuPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512CosPd(a [8]float64) [8]float64
TEXT ·m512CosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCosPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CosPs(a [16]float32) [16]float32
TEXT ·m512CosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCosPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CosdPd(a [8]float64) [8]float64
TEXT ·m512CosdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCosdPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCosdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CosdPs(a [16]float32) [16]float32
TEXT ·m512CosdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCosdPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCosdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CoshPd(a [8]float64) [8]float64
TEXT ·m512CoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCoshPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskCoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CoshPs(a [16]float32) [16]float32
TEXT ·m512CoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCoshPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskCoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtRoundepi32Ps(a [64]byte, rounding int) [16]float32
TEXT ·m512CvtRoundepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundepi32Ps(src [16]float32, k uint16, a [64]byte, rounding int) [16]float32
TEXT ·m512MaskCvtRoundepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvtRoundepi32Ps(k uint16, a [64]byte, rounding int) [16]float32
TEXT ·m512MaskzCvtRoundepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtRoundepu32Ps(a [64]byte, rounding int) [16]float32
TEXT ·m512CvtRoundepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundepu32Ps(src [16]float32, k uint16, a [64]byte, rounding int) [16]float32
TEXT ·m512MaskCvtRoundepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvtRoundepu32Ps(k uint16, a [64]byte, rounding int) [16]float32
TEXT ·m512MaskzCvtRoundepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func cvtRoundi32Ss(a [4]float32, b int, rounding int) [4]float32
TEXT ·cvtRoundi32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func cvtRoundi64Sd(a [2]float64, b int64, rounding int) [2]float64
TEXT ·cvtRoundi64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func cvtRoundi64Ss(a [4]float32, b int64, rounding int) [4]float32
TEXT ·cvtRoundi64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func m512CvtRoundpdEpi32(a [8]float64, rounding int) [32]byte
TEXT ·m512CvtRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2DQ Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvtRoundpdEpi32(src [32]byte, k uint8, a [8]float64, rounding int) [32]byte
TEXT ·m512MaskCvtRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpdEpi32(k uint8, a [8]float64, rounding int) [32]byte
TEXT ·m512MaskzCvtRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512CvtRoundpdEpu32(a [8]float64, rounding int) [32]byte
TEXT ·m512CvtRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvtRoundpdEpu32(src [32]byte, k uint8, a [8]float64, rounding int) [32]byte
TEXT ·m512MaskCvtRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpdEpu32(k uint8, a [8]float64, rounding int) [32]byte
TEXT ·m512MaskzCvtRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512CvtRoundpdPs(a [8]float64, rounding int) [8]float32
TEXT ·m512CvtRoundpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2PS Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvtRoundpdPs(src [8]float32, k uint8, a [8]float64, rounding int) [8]float32
TEXT ·m512MaskCvtRoundpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpdPs(k uint8, a [8]float64, rounding int) [8]float32
TEXT ·m512MaskzCvtRoundpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512CvtRoundphPs(a [32]byte, sae int) [16]float32
TEXT ·m512CvtRoundphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTPH2PS Y0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundphPs(src [16]float32, k uint16, a [32]byte, sae int) [16]float32
TEXT ·m512MaskCvtRoundphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvtRoundphPs(k uint16, a [32]byte, sae int) [16]float32
TEXT ·m512MaskzCvtRoundphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtRoundpsEpi32(a [16]float32, rounding int) [64]byte
TEXT ·m512CvtRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2DQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundpsEpi32(src [64]byte, k uint16, a [16]float32, rounding int) [64]byte
TEXT ·m512MaskCvtRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpsEpi32(k uint16, a [16]float32, rounding int) [64]byte
TEXT ·m512MaskzCvtRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtRoundpsEpu32(a [16]float32, rounding int) [64]byte
TEXT ·m512CvtRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundpsEpu32(src [64]byte, k uint16, a [16]float32, rounding int) [64]byte
TEXT ·m512MaskCvtRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpsEpu32(k uint16, a [16]float32, rounding int) [64]byte
TEXT ·m512MaskzCvtRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtRoundpsPd(a [8]float32, sae int) [8]float64
TEXT ·m512CvtRoundpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTPS2PD Y0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundpsPd(src [8]float64, k uint8, a [8]float32, sae int) [8]float64
TEXT ·m512MaskCvtRoundpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpsPd(k uint8, a [8]float32, sae int) [8]float64
TEXT ·m512MaskzCvtRoundpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskCvtRoundpsPh(src [16]byte, k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskCvtRoundpsPh(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzCvtRoundpsPh(k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskzCvtRoundpsPh(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ rounding+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func m256MaskCvtRoundpsPh(src [16]byte, k uint8, a [8]float32, rounding int) [16]byte
TEXT ·m256MaskCvtRoundpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m256MaskzCvtRoundpsPh(k uint8, a [8]float32, rounding int) [16]byte
TEXT ·m256MaskzCvtRoundpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512CvtRoundpsPh(a [16]float32, rounding int) [32]byte
TEXT ·m512CvtRoundpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2PH Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvtRoundpsPh(src [32]byte, k uint16, a [16]float32, rounding int) [32]byte
TEXT ·m512MaskCvtRoundpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvtRoundpsPh(k uint16, a [16]float32, rounding int) [32]byte
TEXT ·m512MaskzCvtRoundpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func cvtRoundsdI32(a [2]float64, rounding int) int
TEXT ·cvtRoundsdI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdI64(a [2]float64, rounding int) int64
TEXT ·cvtRoundsdI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdSi32(a [2]float64, rounding int) int
TEXT ·cvtRoundsdSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdSi64(a [2]float64, rounding int) int64
TEXT ·cvtRoundsdSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdSs(a [4]float32, b [2]float64, rounding int) [4]float32
TEXT ·cvtRoundsdSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskCvtRoundsdSs(src [4]float32, k uint8, a [4]float32, b [2]float64, rounding int) [4]float32
TEXT ·maskCvtRoundsdSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzCvtRoundsdSs(k uint8, a [4]float32, b [2]float64, rounding int) [4]float32
TEXT ·maskzCvtRoundsdSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func cvtRoundsdU32(a [2]float64, rounding int) uint32
TEXT ·cvtRoundsdU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSD2USI X0, R9

	MOVL $0, ret+24(FP)
	RET

// func cvtRoundsdU64(a [2]float64, rounding int) uint64
TEXT ·cvtRoundsdU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSD2USI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsi32Ss(a [4]float32, b int, rounding int) [4]float32
TEXT ·cvtRoundsi32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func cvtRoundsi64Sd(a [2]float64, b int64, rounding int) [2]float64
TEXT ·cvtRoundsi64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func cvtRoundsi64Ss(a [4]float32, b int64, rounding int) [4]float32
TEXT ·cvtRoundsi64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func cvtRoundssI32(a [4]float32, rounding int) int
TEXT ·cvtRoundssI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssI64(a [4]float32, rounding int) int64
TEXT ·cvtRoundssI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssSd(a [2]float64, b [4]float32, rounding int) [2]float64
TEXT ·cvtRoundssSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskCvtRoundssSd(src [2]float64, k uint8, a [2]float64, b [4]float32, rounding int) [2]float64
TEXT ·maskCvtRoundssSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzCvtRoundssSd(k uint8, a [2]float64, b [4]float32, rounding int) [2]float64
TEXT ·maskzCvtRoundssSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func cvtRoundssSi32(a [4]float32, rounding int) int
TEXT ·cvtRoundssSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssSi64(a [4]float32, rounding int) int64
TEXT ·cvtRoundssSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssU32(a [4]float32, rounding int) uint32
TEXT ·cvtRoundssU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSS2USI X0, R9

	MOVL $0, ret+24(FP)
	RET

// func cvtRoundssU64(a [4]float32, rounding int) uint64
TEXT ·cvtRoundssU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSS2USI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundu32Ss(a [4]float32, b uint32, rounding int) [4]float32
TEXT ·cvtRoundu32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL b+16(FP),R9
	MOVQ rounding+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func cvtRoundu64Sd(a [2]float64, b uint64, rounding int) [2]float64
TEXT ·cvtRoundu64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func cvtRoundu64Ss(a [4]float32, b uint64, rounding int) [4]float32
TEXT ·cvtRoundu64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+32(FP)
	RET

// func maskCvtepi16Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi16Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi16Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi16Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXWD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi16Epi32(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi16Epi32(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi16Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXWD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi16Epi32(a [32]byte) [64]byte
TEXT ·m512Cvtepi16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSXWD Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepi16Epi32(src [64]byte, k uint16, a [32]byte) [64]byte
TEXT ·m512MaskCvtepi16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi16Epi32(k uint16, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepi16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSXWD R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtepi16Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi16Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi16Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXWQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi16Epi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi16Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi16Epi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXWQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi16Epi64(a [16]byte) [64]byte
TEXT ·m512Cvtepi16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSXWQ X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskCvtepi16Epi64(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·m512MaskCvtepi16Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi16Epi64(k uint8, a [16]byte) [64]byte
TEXT ·m512MaskzCvtepi16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXWQ R8, X1

	MOV Z1, ret+20(FP)
	RET

// func cvtepi32Epi16(a [16]byte) [16]byte
TEXT ·cvtepi32Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVDW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi32Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi32Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi32Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi32Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVDW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepi32Epi16(a [32]byte) [16]byte
TEXT ·m256Cvtepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVDW Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtepi32Epi16(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtepi32Epi16(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVDW R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtepi32Epi16(a [64]byte) [32]byte
TEXT ·m512Cvtepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVDW Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtepi32Epi16(src [32]byte, k uint16, a [64]byte) [32]byte
TEXT ·m512MaskCvtepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtepi32Epi16(k uint16, a [64]byte) [32]byte
TEXT ·m512MaskzCvtepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVDW R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func maskCvtepi32Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi32Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi32Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi32Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi32Epi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi32Epi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi32Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXDQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi32Epi64(a [32]byte) [64]byte
TEXT ·m512Cvtepi32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSXDQ Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepi32Epi64(src [64]byte, k uint8, a [32]byte) [64]byte
TEXT ·m512MaskCvtepi32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi32Epi64(k uint8, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepi32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSXDQ R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func cvtepi32Epi8(a [16]byte) [16]byte
TEXT ·cvtepi32Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVDB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi32Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi32Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi32Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi32Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVDB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepi32Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVDB Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtepi32Epi8(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtepi32Epi8(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVDB R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtepi32Epi8(a [64]byte) [16]byte
TEXT ·m512Cvtepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVDB Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtepi32Epi8(src [16]byte, k uint16, a [64]byte) [16]byte
TEXT ·m512MaskCvtepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtepi32Epi8(k uint16, a [64]byte) [16]byte
TEXT ·m512MaskzCvtepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVDB R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func maskCvtepi32Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepi32Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi32Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepi32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTDQ2PD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi32Pd(src [4]float64, k uint8, a [16]byte) [4]float64
TEXT ·m256MaskCvtepi32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi32Pd(k uint8, a [16]byte) [4]float64
TEXT ·m256MaskzCvtepi32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTDQ2PD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi32Pd(a [32]byte) [8]float64
TEXT ·m512Cvtepi32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PD Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepi32Pd(src [8]float64, k uint8, a [32]byte) [8]float64
TEXT ·m512MaskCvtepi32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi32Pd(k uint8, a [32]byte) [8]float64
TEXT ·m512MaskzCvtepi32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PD R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtepi32Ps(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtepi32Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi32Ps(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtepi32Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTDQ2PS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi32Ps(src [8]float32, k uint8, a [32]byte) [8]float32
TEXT ·m256MaskCvtepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi32Ps(k uint8, a [32]byte) [8]float32
TEXT ·m256MaskzCvtepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512Cvtepi32Ps(a [64]byte) [16]float32
TEXT ·m512Cvtepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepi32Ps(src [16]float32, k uint16, a [64]byte) [16]float32
TEXT ·m512MaskCvtepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi32Ps(k uint16, a [64]byte) [16]float32
TEXT ·m512MaskzCvtepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi32StoreuEpi16(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtepi32StoreuEpi16(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtepi32StoreuEpi16(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCvtepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi32StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtepi32StoreuEpi8(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtepi32StoreuEpi8(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCvtepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func cvtepi64Epi16(a [16]byte) [16]byte
TEXT ·cvtepi64Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVQW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi64Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi64Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi64Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVQW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepi64Epi16(a [32]byte) [16]byte
TEXT ·m256Cvtepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVQW Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtepi64Epi16(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtepi64Epi16(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVQW R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtepi64Epi16(a [64]byte) [16]byte
TEXT ·m512Cvtepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVQW Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtepi64Epi16(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·m512MaskCvtepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtepi64Epi16(k uint8, a [64]byte) [16]byte
TEXT ·m512MaskzCvtepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVQW R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func cvtepi64Epi32(a [16]byte) [16]byte
TEXT ·cvtepi64Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVQD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi64Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi64Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi64Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVQD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepi64Epi32(a [32]byte) [16]byte
TEXT ·m256Cvtepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVQD Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtepi64Epi32(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtepi64Epi32(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVQD R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtepi64Epi32(a [64]byte) [32]byte
TEXT ·m512Cvtepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVQD Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtepi64Epi32(src [32]byte, k uint8, a [64]byte) [32]byte
TEXT ·m512MaskCvtepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtepi64Epi32(k uint8, a [64]byte) [32]byte
TEXT ·m512MaskzCvtepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVQD R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvtepi64Epi8(a [16]byte) [16]byte
TEXT ·cvtepi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVQB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi64Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi64Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVQB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepi64Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVQB Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtepi64Epi8(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtepi64Epi8(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVQB R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtepi64Epi8(a [64]byte) [16]byte
TEXT ·m512Cvtepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVQB Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtepi64Epi8(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·m512MaskCvtepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtepi64Epi8(k uint8, a [64]byte) [16]byte
TEXT ·m512MaskzCvtepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVQB R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func maskCvtepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi16(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtepi64StoreuEpi16(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtepi64StoreuEpi16(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtepi64StoreuEpi32(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtepi64StoreuEpi32(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtepi64StoreuEpi8(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtepi64StoreuEpi8(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtepi8Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi8Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi8Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi8Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXBD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi8Epi32(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi8Epi32(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi8Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXBD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi8Epi32(a [16]byte) [64]byte
TEXT ·m512Cvtepi8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSXBD X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskCvtepi8Epi32(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·m512MaskCvtepi8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi8Epi32(k uint16, a [16]byte) [64]byte
TEXT ·m512MaskzCvtepi8Epi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXBD R8, X1

	MOV Z1, ret+20(FP)
	RET

// func maskCvtepi8Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi8Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi8Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXBQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi8Epi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi8Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepi8Epi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXBQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi8Epi64(a [16]byte) [64]byte
TEXT ·m512Cvtepi8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSXBQ X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskCvtepi8Epi64(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·m512MaskCvtepi8Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepi8Epi64(k uint8, a [16]byte) [64]byte
TEXT ·m512MaskzCvtepi8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSXBQ R8, X1

	MOV Z1, ret+20(FP)
	RET

// func maskCvtepu16Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu16Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu16Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu16Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXWD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepu16Epi32(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepu16Epi32(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu16Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXWD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu16Epi32(a [32]byte) [64]byte
TEXT ·m512Cvtepu16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVZXWD Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepu16Epi32(src [64]byte, k uint16, a [32]byte) [64]byte
TEXT ·m512MaskCvtepu16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu16Epi32(k uint16, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepu16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVZXWD R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtepu16Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu16Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu16Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXWQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepu16Epi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu16Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepu16Epi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXWQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu16Epi64(a [16]byte) [64]byte
TEXT ·m512Cvtepu16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVZXWQ X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskCvtepu16Epi64(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·m512MaskCvtepu16Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu16Epi64(k uint8, a [16]byte) [64]byte
TEXT ·m512MaskzCvtepu16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXWQ R8, X1

	MOV Z1, ret+20(FP)
	RET

// func maskCvtepu32Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu32Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu32Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu32Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepu32Epi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepu32Epi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu32Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXDQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu32Epi64(a [32]byte) [64]byte
TEXT ·m512Cvtepu32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVZXDQ Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepu32Epi64(src [64]byte, k uint8, a [32]byte) [64]byte
TEXT ·m512MaskCvtepu32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu32Epi64(k uint8, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepu32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVZXDQ R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func cvtepu32Pd(a [16]byte) [2]float64
TEXT ·cvtepu32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu32Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepu32Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu32Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepu32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepu32Pd(a [16]byte) [4]float64
TEXT ·m256Cvtepu32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD X0

	MOV Y0, ret+16(FP)
	RET

// func m256MaskCvtepu32Pd(src [4]float64, k uint8, a [16]byte) [4]float64
TEXT ·m256MaskCvtepu32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepu32Pd(k uint8, a [16]byte) [4]float64
TEXT ·m256MaskzCvtepu32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu32Pd(a [32]byte) [8]float64
TEXT ·m512Cvtepu32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepu32Pd(src [8]float64, k uint8, a [32]byte) [8]float64
TEXT ·m512MaskCvtepu32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu32Pd(k uint8, a [32]byte) [8]float64
TEXT ·m512MaskzCvtepu32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func m512Cvtepu32Ps(a [64]byte) [16]float32
TEXT ·m512Cvtepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepu32Ps(src [16]float32, k uint16, a [64]byte) [16]float32
TEXT ·m512MaskCvtepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu32Ps(k uint16, a [64]byte) [16]float32
TEXT ·m512MaskzCvtepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtepu8Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu8Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu8Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu8Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXBD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepu8Epi32(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepu8Epi32(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu8Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXBD R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu8Epi32(a [16]byte) [64]byte
TEXT ·m512Cvtepu8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVZXBD X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskCvtepu8Epi32(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·m512MaskCvtepu8Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu8Epi32(k uint16, a [16]byte) [64]byte
TEXT ·m512MaskzCvtepu8Epi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXBD R8, X1

	MOV Z1, ret+20(FP)
	RET

// func maskCvtepu8Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu8Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu8Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXBQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepu8Epi64(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu8Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtepu8Epi64(k uint8, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXBQ R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu8Epi64(a [16]byte) [64]byte
TEXT ·m512Cvtepu8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVZXBQ X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskCvtepu8Epi64(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·m512MaskCvtepu8Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtepu8Epi64(k uint8, a [16]byte) [64]byte
TEXT ·m512MaskzCvtepu8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVZXBQ R8, X1

	MOV Z1, ret+20(FP)
	RET

// func cvti32Sd(a [2]float64, b int) [2]float64
TEXT ·cvti32Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSI2SD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvti32Ss(a [4]float32, b int) [4]float32
TEXT ·cvti32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSI2SS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvti64Sd(a [2]float64, b int64) [2]float64
TEXT ·cvti64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSI2SD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvti64Ss(a [4]float32, b int64) [4]float32
TEXT ·cvti64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTSI2SS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func maskCvtpdEpi32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtpdEpi32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPD2DQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtpdEpi32(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·m256MaskCvtpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtpdEpi32(k uint8, a [4]float64) [16]byte
TEXT ·m256MaskzCvtpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTPD2DQ R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512CvtpdEpi32(a [8]float64) [32]byte
TEXT ·m512CvtpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2DQ Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtpdEpi32(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·m512MaskCvtpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtpdEpi32(k uint8, a [8]float64) [32]byte
TEXT ·m512MaskzCvtpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2DQ R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvtpdEpu32(a [2]float64) [16]byte
TEXT ·cvtpdEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpu32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtpdEpu32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256CvtpdEpu32(a [4]float64) [16]byte
TEXT ·m256CvtpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtpdEpu32(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·m256MaskCvtpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtpdEpu32(k uint8, a [4]float64) [16]byte
TEXT ·m256MaskzCvtpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512CvtpdEpu32(a [8]float64) [32]byte
TEXT ·m512CvtpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtpdEpu32(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·m512MaskCvtpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtpdEpu32(k uint8, a [8]float64) [32]byte
TEXT ·m512MaskzCvtpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2UDQ R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func maskCvtpdPs(src [4]float32, k uint8, a [2]float64) [4]float32
TEXT ·maskCvtpdPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtpdPs(k uint8, a [2]float64) [4]float32
TEXT ·maskzCvtpdPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPD2PS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtpdPs(src [4]float32, k uint8, a [4]float64) [4]float32
TEXT ·m256MaskCvtpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtpdPs(k uint8, a [4]float64) [4]float32
TEXT ·m256MaskzCvtpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTPD2PS R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512CvtpdPs(a [8]float64) [8]float32
TEXT ·m512CvtpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2PS Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtpdPs(src [8]float32, k uint8, a [8]float64) [8]float32
TEXT ·m512MaskCvtpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtpdPs(k uint8, a [8]float64) [8]float32
TEXT ·m512MaskzCvtpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2PS R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func maskCvtphPs(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtphPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtphPs(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtphPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPH2PS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtphPs(src [8]float32, k uint8, a [16]byte) [8]float32
TEXT ·m256MaskCvtphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtphPs(k uint8, a [16]byte) [8]float32
TEXT ·m256MaskzCvtphPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPH2PS R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512CvtphPs(a [32]byte) [16]float32
TEXT ·m512CvtphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTPH2PS Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtphPs(src [16]float32, k uint16, a [32]byte) [16]float32
TEXT ·m512MaskCvtphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtphPs(k uint16, a [32]byte) [16]float32
TEXT ·m512MaskzCvtphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VCVTPH2PS R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtpsEpi32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtpsEpi32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPS2DQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtpsEpi32(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·m256MaskCvtpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtpsEpi32(k uint8, a [8]float32) [32]byte
TEXT ·m256MaskzCvtpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTPS2DQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512CvtpsEpi32(a [16]float32) [64]byte
TEXT ·m512CvtpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2DQ Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtpsEpi32(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·m512MaskCvtpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtpsEpi32(k uint16, a [16]float32) [64]byte
TEXT ·m512MaskzCvtpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2DQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func cvtpsEpu32(a [4]float32) [16]byte
TEXT ·cvtpsEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpu32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtpsEpu32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256CvtpsEpu32(a [8]float32) [32]byte
TEXT ·m256CvtpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ Y0

	MOV Y0, ret+0(FP)
	RET

// func m256MaskCvtpsEpu32(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·m256MaskCvtpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvtpsEpu32(k uint8, a [8]float32) [32]byte
TEXT ·m256MaskzCvtpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512CvtpsEpu32(a [16]float32) [64]byte
TEXT ·m512CvtpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtpsEpu32(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·m512MaskCvtpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtpsEpu32(k uint16, a [16]float32) [64]byte
TEXT ·m512MaskzCvtpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2UDQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512CvtpsPd(a [8]float32) [8]float64
TEXT ·m512CvtpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTPS2PD Y0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtpsPd(src [8]float64, k uint8, a [8]float32) [8]float64
TEXT ·m512MaskCvtpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvtpsPd(k uint8, a [8]float32) [8]float64
TEXT ·m512MaskzCvtpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTPS2PD R8, Y1

	MOV Z1, ret+0(FP)
	RET

// func maskCvtpsPh(src [16]byte, k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskCvtpsPh(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzCvtpsPh(k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskzCvtpsPh(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ rounding+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func m256MaskCvtpsPh(src [16]byte, k uint8, a [8]float32, rounding int) [16]byte
TEXT ·m256MaskCvtpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m256MaskzCvtpsPh(k uint8, a [8]float32, rounding int) [16]byte
TEXT ·m256MaskzCvtpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512CvtpsPh(a [16]float32, rounding int) [32]byte
TEXT ·m512CvtpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2PH Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvtpsPh(src [32]byte, k uint16, a [16]float32, rounding int) [32]byte
TEXT ·m512MaskCvtpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvtpsPh(k uint16, a [16]float32, rounding int) [32]byte
TEXT ·m512MaskzCvtpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func cvtsdI32(a [2]float64) int
TEXT ·cvtsdI32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSD2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvtsdI64(a [2]float64) int64
TEXT ·cvtsdI64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSD2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func maskCvtsdSs(src [4]float32, k uint8, a [4]float32, b [2]float64) [4]float32
TEXT ·maskCvtsdSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzCvtsdSs(k uint8, a [4]float32, b [2]float64) [4]float32
TEXT ·maskzCvtsdSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func cvtsdU32(a [2]float64) uint32
TEXT ·cvtsdU32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSD2USI X0

	MOVL $0, ret+16(FP)
	RET

// func cvtsdU64(a [2]float64) uint64
TEXT ·cvtsdU64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSD2USI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvtsepi32Epi16(a [16]byte) [16]byte
TEXT ·cvtsepi32Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSDW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi32Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi32Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtsepi32Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi32Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSDW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtsepi32Epi16(a [32]byte) [16]byte
TEXT ·m256Cvtsepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSDW Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtsepi32Epi16(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtsepi32Epi16(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSDW R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtsepi32Epi16(a [64]byte) [32]byte
TEXT ·m512Cvtsepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSDW Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtsepi32Epi16(src [32]byte, k uint16, a [64]byte) [32]byte
TEXT ·m512MaskCvtsepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtsepi32Epi16(k uint16, a [64]byte) [32]byte
TEXT ·m512MaskzCvtsepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSDW R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvtsepi32Epi8(a [16]byte) [16]byte
TEXT ·cvtsepi32Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSDB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi32Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi32Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtsepi32Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi32Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSDB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtsepi32Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtsepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSDB Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtsepi32Epi8(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtsepi32Epi8(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSDB R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtsepi32Epi8(a [64]byte) [16]byte
TEXT ·m512Cvtsepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSDB Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtsepi32Epi8(src [16]byte, k uint16, a [64]byte) [16]byte
TEXT ·m512MaskCvtsepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtsepi32Epi8(k uint16, a [64]byte) [16]byte
TEXT ·m512MaskzCvtsepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSDB R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func maskCvtsepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi32StoreuEpi16(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtsepi32StoreuEpi16(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtsepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtsepi32StoreuEpi16(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCvtsepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi32StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtsepi32StoreuEpi8(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtsepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtsepi32StoreuEpi8(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCvtsepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func cvtsepi64Epi16(a [16]byte) [16]byte
TEXT ·cvtsepi64Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSQW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi64Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi64Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtsepi64Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi64Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSQW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtsepi64Epi16(a [32]byte) [16]byte
TEXT ·m256Cvtsepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSQW Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtsepi64Epi16(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtsepi64Epi16(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSQW R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtsepi64Epi16(a [64]byte) [16]byte
TEXT ·m512Cvtsepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSQW Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtsepi64Epi16(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·m512MaskCvtsepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtsepi64Epi16(k uint8, a [64]byte) [16]byte
TEXT ·m512MaskzCvtsepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSQW R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func cvtsepi64Epi32(a [16]byte) [16]byte
TEXT ·cvtsepi64Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSQD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi64Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi64Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtsepi64Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi64Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSQD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtsepi64Epi32(a [32]byte) [16]byte
TEXT ·m256Cvtsepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSQD Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtsepi64Epi32(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtsepi64Epi32(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSQD R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtsepi64Epi32(a [64]byte) [32]byte
TEXT ·m512Cvtsepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSQD Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtsepi64Epi32(src [32]byte, k uint8, a [64]byte) [32]byte
TEXT ·m512MaskCvtsepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtsepi64Epi32(k uint8, a [64]byte) [32]byte
TEXT ·m512MaskzCvtsepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSQD R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvtsepi64Epi8(a [16]byte) [16]byte
TEXT ·cvtsepi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVSQB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi64Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtsepi64Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi64Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVSQB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtsepi64Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtsepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSQB Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtsepi64Epi8(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtsepi64Epi8(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVSQB R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtsepi64Epi8(a [64]byte) [16]byte
TEXT ·m512Cvtsepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSQB Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtsepi64Epi8(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·m512MaskCvtsepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtsepi64Epi8(k uint8, a [64]byte) [16]byte
TEXT ·m512MaskzCvtsepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVSQB R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func maskCvtsepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi16(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtsepi64StoreuEpi16(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtsepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtsepi64StoreuEpi16(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtsepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtsepi64StoreuEpi32(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtsepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtsepi64StoreuEpi32(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtsepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtsepi64StoreuEpi8(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtsepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtsepi64StoreuEpi8(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtsepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func cvtssI32(a [4]float32) int
TEXT ·cvtssI32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvtssI64(a [4]float32) int64
TEXT ·cvtssI64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func maskCvtssSd(src [2]float64, k uint8, a [2]float64, b [4]float32) [2]float64
TEXT ·maskCvtssSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzCvtssSd(k uint8, a [2]float64, b [4]float32) [2]float64
TEXT ·maskzCvtssSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func cvtssU32(a [4]float32) uint32
TEXT ·cvtssU32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSS2USI X0

	MOVL $0, ret+16(FP)
	RET

// func cvtssU64(a [4]float32) uint64
TEXT ·cvtssU64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTSS2USI X0

	MOVQ $0, ret+16(FP)
	RET

// func m512CvttRoundpdEpi32(a [8]float64, sae int) [32]byte
TEXT ·m512CvttRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2DQ Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvttRoundpdEpi32(src [32]byte, k uint8, a [8]float64, sae int) [32]byte
TEXT ·m512MaskCvttRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvttRoundpdEpi32(k uint8, a [8]float64, sae int) [32]byte
TEXT ·m512MaskzCvttRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512CvttRoundpdEpu32(a [8]float64, sae int) [32]byte
TEXT ·m512CvttRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskCvttRoundpdEpu32(src [32]byte, k uint8, a [8]float64, sae int) [32]byte
TEXT ·m512MaskCvttRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzCvttRoundpdEpu32(k uint8, a [8]float64, sae int) [32]byte
TEXT ·m512MaskzCvttRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512CvttRoundpsEpi32(a [16]float32, sae int) [64]byte
TEXT ·m512CvttRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTTPS2DQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvttRoundpsEpi32(src [64]byte, k uint16, a [16]float32, sae int) [64]byte
TEXT ·m512MaskCvttRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvttRoundpsEpi32(k uint16, a [16]float32, sae int) [64]byte
TEXT ·m512MaskzCvttRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvttRoundpsEpu32(a [16]float32, sae int) [64]byte
TEXT ·m512CvttRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvttRoundpsEpu32(src [64]byte, k uint16, a [16]float32, sae int) [64]byte
TEXT ·m512MaskCvttRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzCvttRoundpsEpu32(k uint16, a [16]float32, sae int) [64]byte
TEXT ·m512MaskzCvttRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func cvttRoundsdI32(a [2]float64, rounding int) int
TEXT ·cvttRoundsdI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdI64(a [2]float64, rounding int) int64
TEXT ·cvttRoundsdI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdSi32(a [2]float64, rounding int) int
TEXT ·cvttRoundsdSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdSi64(a [2]float64, rounding int) int64
TEXT ·cvttRoundsdSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSD2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdU32(a [2]float64, rounding int) uint32
TEXT ·cvttRoundsdU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSD2USI X0, R9

	MOVL $0, ret+24(FP)
	RET

// func cvttRoundsdU64(a [2]float64, rounding int) uint64
TEXT ·cvttRoundsdU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSD2USI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssI32(a [4]float32, rounding int) int
TEXT ·cvttRoundssI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssI64(a [4]float32, rounding int) int64
TEXT ·cvttRoundssI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssSi32(a [4]float32, rounding int) int
TEXT ·cvttRoundssSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssSi64(a [4]float32, rounding int) int64
TEXT ·cvttRoundssSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSS2SI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssU32(a [4]float32, rounding int) uint32
TEXT ·cvttRoundssU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSS2USI X0, R9

	MOVL $0, ret+24(FP)
	RET

// func cvttRoundssU64(a [4]float32, rounding int) uint64
TEXT ·cvttRoundssU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTTSS2USI X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func maskCvttpdEpi32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvttpdEpi32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTTPD2DQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvttpdEpi32(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·m256MaskCvttpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvttpdEpi32(k uint8, a [4]float64) [16]byte
TEXT ·m256MaskzCvttpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2DQ R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512CvttpdEpi32(a [8]float64) [32]byte
TEXT ·m512CvttpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2DQ Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvttpdEpi32(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·m512MaskCvttpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvttpdEpi32(k uint8, a [8]float64) [32]byte
TEXT ·m512MaskzCvttpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2DQ R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvttpdEpu32(a [2]float64) [16]byte
TEXT ·cvttpdEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpu32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvttpdEpu32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256CvttpdEpu32(a [4]float64) [16]byte
TEXT ·m256CvttpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvttpdEpu32(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·m256MaskCvttpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvttpdEpu32(k uint8, a [4]float64) [16]byte
TEXT ·m256MaskzCvttpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512CvttpdEpu32(a [8]float64) [32]byte
TEXT ·m512CvttpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvttpdEpu32(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·m512MaskCvttpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvttpdEpu32(k uint8, a [8]float64) [32]byte
TEXT ·m512MaskzCvttpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTTPD2UDQ R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func maskCvttpsEpi32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvttpsEpi32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTTPS2DQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvttpsEpi32(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·m256MaskCvttpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvttpsEpi32(k uint8, a [8]float32) [32]byte
TEXT ·m256MaskzCvttpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTTPS2DQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512CvttpsEpi32(a [16]float32) [64]byte
TEXT ·m512CvttpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTTPS2DQ Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvttpsEpi32(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·m512MaskCvttpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvttpsEpi32(k uint16, a [16]float32) [64]byte
TEXT ·m512MaskzCvttpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTTPS2DQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func cvttpsEpu32(a [4]float32) [16]byte
TEXT ·cvttpsEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpu32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvttpsEpu32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256CvttpsEpu32(a [8]float32) [32]byte
TEXT ·m256CvttpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ Y0

	MOV Y0, ret+0(FP)
	RET

// func m256MaskCvttpsEpu32(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·m256MaskCvttpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzCvttpsEpu32(k uint8, a [8]float32) [32]byte
TEXT ·m256MaskzCvttpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512CvttpsEpu32(a [16]float32) [64]byte
TEXT ·m512CvttpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvttpsEpu32(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·m512MaskCvttpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzCvttpsEpu32(k uint16, a [16]float32) [64]byte
TEXT ·m512MaskzCvttpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTTPS2UDQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func cvttsdI32(a [2]float64) int
TEXT ·cvttsdI32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSD2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttsdI64(a [2]float64) int64
TEXT ·cvttsdI64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSD2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttsdU32(a [2]float64) uint32
TEXT ·cvttsdU32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSD2USI X0

	MOVL $0, ret+16(FP)
	RET

// func cvttsdU64(a [2]float64) uint64
TEXT ·cvttsdU64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSD2USI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttssI32(a [4]float32) int
TEXT ·cvttssI32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttssI64(a [4]float32) int64
TEXT ·cvttssI64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttssU32(a [4]float32) uint32
TEXT ·cvttssU32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSS2USI X0

	MOVL $0, ret+16(FP)
	RET

// func cvttssU64(a [4]float32) uint64
TEXT ·cvttssU64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VCVTTSS2USI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvtu32Sd(a [2]float64, b uint32) [2]float64
TEXT ·cvtu32Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTUSI2SD X0, R9

	MOVOU X1, ret+20(FP)
	RET

// func cvtu32Ss(a [4]float32, b uint32) [4]float32
TEXT ·cvtu32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTUSI2SS X0, R9

	MOVOU X1, ret+20(FP)
	RET

// func cvtu64Sd(a [2]float64, b uint64) [2]float64
TEXT ·cvtu64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTUSI2SD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvtu64Ss(a [4]float32, b uint64) [4]float32
TEXT ·cvtu64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VCVTUSI2SS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvtusepi32Epi16(a [16]byte) [16]byte
TEXT ·cvtusepi32Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVUSDW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi32Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi32Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtusepi32Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi32Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVUSDW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtusepi32Epi16(a [32]byte) [16]byte
TEXT ·m256Cvtusepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDW Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtusepi32Epi16(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtusepi32Epi16(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDW R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtusepi32Epi16(a [64]byte) [32]byte
TEXT ·m512Cvtusepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDW Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtusepi32Epi16(src [32]byte, k uint16, a [64]byte) [32]byte
TEXT ·m512MaskCvtusepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtusepi32Epi16(k uint16, a [64]byte) [32]byte
TEXT ·m512MaskzCvtusepi32Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDW R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvtusepi32Epi8(a [16]byte) [16]byte
TEXT ·cvtusepi32Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVUSDB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi32Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi32Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtusepi32Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi32Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVUSDB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtusepi32Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtusepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDB Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtusepi32Epi8(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtusepi32Epi8(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDB R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtusepi32Epi8(a [64]byte) [16]byte
TEXT ·m512Cvtusepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDB Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtusepi32Epi8(src [16]byte, k uint16, a [64]byte) [16]byte
TEXT ·m512MaskCvtusepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtusepi32Epi8(k uint16, a [64]byte) [16]byte
TEXT ·m512MaskzCvtusepi32Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSDB R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func maskCvtusepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi32StoreuEpi16(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtusepi32StoreuEpi16(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtusepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtusepi32StoreuEpi16(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCvtusepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi32StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtusepi32StoreuEpi8(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtusepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtusepi32StoreuEpi8(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskCvtusepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func cvtusepi64Epi16(a [16]byte) [16]byte
TEXT ·cvtusepi64Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVUSQW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi64Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi64Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtusepi64Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi64Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVUSQW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtusepi64Epi16(a [32]byte) [16]byte
TEXT ·m256Cvtusepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQW Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtusepi64Epi16(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtusepi64Epi16(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQW R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtusepi64Epi16(a [64]byte) [16]byte
TEXT ·m512Cvtusepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQW Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtusepi64Epi16(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·m512MaskCvtusepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtusepi64Epi16(k uint8, a [64]byte) [16]byte
TEXT ·m512MaskzCvtusepi64Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQW R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func cvtusepi64Epi32(a [16]byte) [16]byte
TEXT ·cvtusepi64Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVUSQD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi64Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi64Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtusepi64Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi64Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVUSQD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtusepi64Epi32(a [32]byte) [16]byte
TEXT ·m256Cvtusepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQD Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtusepi64Epi32(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtusepi64Epi32(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQD R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtusepi64Epi32(a [64]byte) [32]byte
TEXT ·m512Cvtusepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQD Z0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtusepi64Epi32(src [32]byte, k uint8, a [64]byte) [32]byte
TEXT ·m512MaskCvtusepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzCvtusepi64Epi32(k uint8, a [64]byte) [32]byte
TEXT ·m512MaskzCvtusepi64Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQD R8, Z1

	MOV Y1, ret+0(FP)
	RET

// func cvtusepi64Epi8(a [16]byte) [16]byte
TEXT ·cvtusepi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VPMOVUSQB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi64Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtusepi64Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi64Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMOVUSQB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtusepi64Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtusepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQB Y0

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtusepi64Epi8(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskzCvtusepi64Epi8(k uint8, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQB R8, Y1

	MOVOU X1, ret+0(FP)
	RET

// func m512Cvtusepi64Epi8(a [64]byte) [16]byte
TEXT ·m512Cvtusepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQB Z0

	MOVOU X0, ret+0(FP)
	RET

// func m512MaskCvtusepi64Epi8(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·m512MaskCvtusepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512MaskzCvtusepi64Epi8(k uint8, a [64]byte) [16]byte
TEXT ·m512MaskzCvtusepi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMOVUSQB R8, Z1

	MOVOU X1, ret+0(FP)
	RET

// func maskCvtusepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi16(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtusepi64StoreuEpi16(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtusepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtusepi64StoreuEpi16(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtusepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtusepi64StoreuEpi32(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtusepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtusepi64StoreuEpi32(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtusepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskCvtusepi64StoreuEpi8(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskCvtusepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskCvtusepi64StoreuEpi8(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskCvtusepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512DivEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512DivEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskDivEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskDivEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512DivEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512DivEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512DivEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512DivEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskDivEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskDivEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512DivEpu64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512DivEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512DivEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func maskDivPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskDivPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzDivPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzDivPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskDivPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskDivPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzDivPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzDivPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512DivPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512DivPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VDIVPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskDivPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskDivPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzDivPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzDivPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskDivPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskDivPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzDivPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzDivPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskDivPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskDivPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzDivPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzDivPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512DivPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512DivPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VDIVPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskDivPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskDivPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzDivPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzDivPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512DivRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512DivRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskDivRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskDivRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzDivRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskzDivRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512DivRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512DivRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskDivRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskDivRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzDivRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskzDivRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func divRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·divRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskDivRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskDivRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzDivRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzDivRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func divRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·divRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskDivRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskDivRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzDivRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzDivRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskDivSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskDivSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzDivSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzDivSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskDivSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskDivSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzDivSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzDivSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m512ErfPd(a [8]float64) [8]float64
TEXT ·m512ErfPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskErfPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfPs(a [16]float32) [16]float32
TEXT ·m512ErfPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskErfPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfcPd(a [8]float64) [8]float64
TEXT ·m512ErfcPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfcPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskErfcPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfcPs(a [16]float32) [16]float32
TEXT ·m512ErfcPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfcPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskErfcPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfcinvPd(a [8]float64) [8]float64
TEXT ·m512ErfcinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfcinvPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskErfcinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfcinvPs(a [16]float32) [16]float32
TEXT ·m512ErfcinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfcinvPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskErfcinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfinvPd(a [8]float64) [8]float64
TEXT ·m512ErfinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfinvPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskErfinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ErfinvPs(a [16]float32) [16]float32
TEXT ·m512ErfinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskErfinvPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskErfinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ExpPd(a [8]float64) [8]float64
TEXT ·m512ExpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExpPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskExpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ExpPs(a [16]float32) [16]float32
TEXT ·m512ExpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExpPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskExpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Exp10Pd(a [8]float64) [8]float64
TEXT ·m512Exp10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp10Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskExp10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Exp10Ps(a [16]float32) [16]float32
TEXT ·m512Exp10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp10Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskExp10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Exp2Pd(a [8]float64) [8]float64
TEXT ·m512Exp2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp2Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskExp2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Exp2Ps(a [16]float32) [16]float32
TEXT ·m512Exp2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp2Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskExp2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskExpandEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskExpandEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzExpandEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzExpandEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPEXPANDD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskExpandEpi32(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskExpandEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandEpi32(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzExpandEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPEXPANDD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskExpandEpi32(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·m512MaskExpandEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandEpi32(k uint16, a [64]byte) [64]byte
TEXT ·m512MaskzExpandEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPEXPANDD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskExpandEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskExpandEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzExpandEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzExpandEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VPEXPANDQ R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskExpandEpi64(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskExpandEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandEpi64(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzExpandEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPEXPANDQ R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskExpandEpi64(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·m512MaskExpandEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandEpi64(k uint8, a [64]byte) [64]byte
TEXT ·m512MaskzExpandEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPEXPANDQ R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskExpandPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskExpandPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzExpandPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzExpandPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VEXPANDPD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskExpandPd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskExpandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandPd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzExpandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VEXPANDPD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskExpandPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskExpandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandPd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzExpandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VEXPANDPD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskExpandPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskExpandPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzExpandPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzExpandPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VEXPANDPS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskExpandPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskExpandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzExpandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VEXPANDPS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskExpandPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskExpandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzExpandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VEXPANDPS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskExpandloaduEpi32(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskExpandloaduEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzExpandloaduEpi32(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzExpandloaduEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPEXPANDD R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskExpandloaduEpi32(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskExpandloaduEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandloaduEpi32(k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskzExpandloaduEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPEXPANDD R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskExpandloaduEpi32(src [64]byte, k uint16, mem_addr uintptr) [64]byte
TEXT ·m512MaskExpandloaduEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandloaduEpi32(k uint16, mem_addr uintptr) [64]byte
TEXT ·m512MaskzExpandloaduEpi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPEXPANDD R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskExpandloaduEpi64(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskExpandloaduEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzExpandloaduEpi64(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzExpandloaduEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPEXPANDQ R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskExpandloaduEpi64(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskExpandloaduEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandloaduEpi64(k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskzExpandloaduEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPEXPANDQ R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskExpandloaduEpi64(src [64]byte, k uint8, mem_addr uintptr) [64]byte
TEXT ·m512MaskExpandloaduEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandloaduEpi64(k uint8, mem_addr uintptr) [64]byte
TEXT ·m512MaskzExpandloaduEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPEXPANDQ R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskExpandloaduPd(src [2]float64, k uint8, mem_addr uintptr) [2]float64
TEXT ·maskExpandloaduPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzExpandloaduPd(k uint8, mem_addr uintptr) [2]float64
TEXT ·maskzExpandloaduPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VEXPANDPD R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskExpandloaduPd(src [4]float64, k uint8, mem_addr uintptr) [4]float64
TEXT ·m256MaskExpandloaduPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandloaduPd(k uint8, mem_addr uintptr) [4]float64
TEXT ·m256MaskzExpandloaduPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VEXPANDPD R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskExpandloaduPd(src [8]float64, k uint8, mem_addr uintptr) [8]float64
TEXT ·m512MaskExpandloaduPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandloaduPd(k uint8, mem_addr uintptr) [8]float64
TEXT ·m512MaskzExpandloaduPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VEXPANDPD R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskExpandloaduPs(src [4]float32, k uint8, mem_addr uintptr) [4]float32
TEXT ·maskExpandloaduPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzExpandloaduPs(k uint8, mem_addr uintptr) [4]float32
TEXT ·maskzExpandloaduPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VEXPANDPS R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskExpandloaduPs(src [8]float32, k uint8, mem_addr uintptr) [8]float32
TEXT ·m256MaskExpandloaduPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzExpandloaduPs(k uint8, mem_addr uintptr) [8]float32
TEXT ·m256MaskzExpandloaduPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VEXPANDPS R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskExpandloaduPs(src [16]float32, k uint16, mem_addr uintptr) [16]float32
TEXT ·m512MaskExpandloaduPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzExpandloaduPs(k uint16, mem_addr uintptr) [16]float32
TEXT ·m512MaskzExpandloaduPs(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VEXPANDPS R8, R9

	MOV Z1, ret+12(FP)
	RET

// func m512Expm1Pd(a [8]float64) [8]float64
TEXT ·m512Expm1Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExpm1Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskExpm1Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Expm1Ps(a [16]float32) [16]float32
TEXT ·m512Expm1Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExpm1Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskExpm1Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m256Extractf32x4Ps(a [8]float32, imm8 int) [4]float32
TEXT ·m256Extractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VEXTRACTF32X4 Y0, R9

	MOVOU X1, ret+0(FP)
	RET

// func m256MaskExtractf32x4Ps(src [4]float32, k uint8, a [8]float32, imm8 int) [4]float32
TEXT ·m256MaskExtractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m256MaskzExtractf32x4Ps(k uint8, a [8]float32, imm8 int) [4]float32
TEXT ·m256MaskzExtractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512Extractf32x4Ps(a [16]float32, imm8 int) [4]float32
TEXT ·m512Extractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VEXTRACTF32X4 Z0, R9

	MOVOU X1, ret+0(FP)
	RET

// func m512MaskExtractf32x4Ps(src [4]float32, k uint8, a [16]float32, imm8 int) [4]float32
TEXT ·m512MaskExtractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m512MaskzExtractf32x4Ps(k uint8, a [16]float32, imm8 int) [4]float32
TEXT ·m512MaskzExtractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512Extractf64x4Pd(a [8]float64, imm8 int) [4]float64
TEXT ·m512Extractf64x4Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VEXTRACTF64X4 Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskExtractf64x4Pd(src [4]float64, k uint8, a [8]float64, imm8 int) [4]float64
TEXT ·m512MaskExtractf64x4Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzExtractf64x4Pd(k uint8, a [8]float64, imm8 int) [4]float64
TEXT ·m512MaskzExtractf64x4Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256Extracti32x4Epi32(a [32]byte, imm8 int) [16]byte
TEXT ·m256Extracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VEXTRACTI32X4 Y0, R9

	MOVOU X1, ret+0(FP)
	RET

// func m256MaskExtracti32x4Epi32(src [16]byte, k uint8, a [32]byte, imm8 int) [16]byte
TEXT ·m256MaskExtracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m256MaskzExtracti32x4Epi32(k uint8, a [32]byte, imm8 int) [16]byte
TEXT ·m256MaskzExtracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512Extracti32x4Epi32(a [64]byte, imm8 int) [16]byte
TEXT ·m512Extracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VEXTRACTI32X4 Z0, R9

	MOVOU X1, ret+0(FP)
	RET

// func m512MaskExtracti32x4Epi32(src [16]byte, k uint8, a [64]byte, imm8 int) [16]byte
TEXT ·m512MaskExtracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m512MaskzExtracti32x4Epi32(k uint8, a [64]byte, imm8 int) [16]byte
TEXT ·m512MaskzExtracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m512Extracti64x4Epi64(a [64]byte, imm8 int) [32]byte
TEXT ·m512Extracti64x4Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VEXTRACTI64X4 Z0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskExtracti64x4Epi64(src [32]byte, k uint8, a [64]byte, imm8 int) [32]byte
TEXT ·m512MaskExtracti64x4Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzExtracti64x4Epi64(k uint8, a [64]byte, imm8 int) [32]byte
TEXT ·m512MaskzExtracti64x4Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func fixupimmPd(a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·fixupimmPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+56(FP)
	RET

// func maskFixupimmPd(a [2]float64, k uint8, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskFixupimmPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFixupimmPd(k uint8, a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskzFixupimmPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func m256FixupimmPd(a [4]float64, b [4]float64, c [32]byte, imm8 int) [4]float64
TEXT ·m256FixupimmPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskFixupimmPd(a [4]float64, k uint8, b [4]float64, c [32]byte, imm8 int) [4]float64
TEXT ·m256MaskFixupimmPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzFixupimmPd(k uint8, a [4]float64, b [4]float64, c [32]byte, imm8 int) [4]float64
TEXT ·m256MaskzFixupimmPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m512FixupimmPd(a [8]float64, b [8]float64, c [64]byte, imm8 int) [8]float64
TEXT ·m512FixupimmPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFixupimmPd(a [8]float64, k uint8, b [8]float64, c [64]byte, imm8 int) [8]float64
TEXT ·m512MaskFixupimmPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFixupimmPd(k uint8, a [8]float64, b [8]float64, c [64]byte, imm8 int) [8]float64
TEXT ·m512MaskzFixupimmPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func fixupimmPs(a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·fixupimmPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+56(FP)
	RET

// func maskFixupimmPs(a [4]float32, k uint8, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskFixupimmPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFixupimmPs(k uint8, a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskzFixupimmPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func m256FixupimmPs(a [8]float32, b [8]float32, c [32]byte, imm8 int) [8]float32
TEXT ·m256FixupimmPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskFixupimmPs(a [8]float32, k uint8, b [8]float32, c [32]byte, imm8 int) [8]float32
TEXT ·m256MaskFixupimmPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzFixupimmPs(k uint8, a [8]float32, b [8]float32, c [32]byte, imm8 int) [8]float32
TEXT ·m256MaskzFixupimmPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m512FixupimmPs(a [16]float32, b [16]float32, c [64]byte, imm8 int) [16]float32
TEXT ·m512FixupimmPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFixupimmPs(a [16]float32, k uint16, b [16]float32, c [64]byte, imm8 int) [16]float32
TEXT ·m512MaskFixupimmPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFixupimmPs(k uint16, a [16]float32, b [16]float32, c [64]byte, imm8 int) [16]float32
TEXT ·m512MaskzFixupimmPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FixupimmRoundPd(a [8]float64, b [8]float64, c [64]byte, imm8 int, rounding int) [8]float64
TEXT ·m512FixupimmRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskFixupimmRoundPd(a [8]float64, k uint8, b [8]float64, c [64]byte, imm8 int, rounding int) [8]float64
TEXT ·m512MaskFixupimmRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512MaskzFixupimmRoundPd(k uint8, a [8]float64, b [8]float64, c [64]byte, imm8 int, rounding int) [8]float64
TEXT ·m512MaskzFixupimmRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512FixupimmRoundPs(a [16]float32, b [16]float32, c [64]byte, imm8 int, rounding int) [16]float32
TEXT ·m512FixupimmRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskFixupimmRoundPs(a [16]float32, k uint16, b [16]float32, c [64]byte, imm8 int, rounding int) [16]float32
TEXT ·m512MaskFixupimmRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512MaskzFixupimmRoundPs(k uint16, a [16]float32, b [16]float32, c [64]byte, imm8 int, rounding int) [16]float32
TEXT ·m512MaskzFixupimmRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func fixupimmRoundSd(a [2]float64, b [2]float64, c [16]byte, imm8 int, rounding int) [2]float64
TEXT ·fixupimmRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11
	MOVQ rounding+56(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+64(FP)
	RET

// func maskFixupimmRoundSd(a [2]float64, k uint8, b [2]float64, c [16]byte, imm8 int, rounding int) [2]float64
TEXT ·maskFixupimmRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	// TODO: Code missing

	MOVOU X5, ret+68(FP)
	RET

// func maskzFixupimmRoundSd(k uint8, a [2]float64, b [2]float64, c [16]byte, imm8 int, rounding int) [2]float64
TEXT ·maskzFixupimmRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	// TODO: Code missing

	MOVOU X5, ret+68(FP)
	RET

// func fixupimmRoundSs(a [4]float32, b [4]float32, c [16]byte, imm8 int, rounding int) [4]float32
TEXT ·fixupimmRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11
	MOVQ rounding+56(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+64(FP)
	RET

// func maskFixupimmRoundSs(a [4]float32, k uint8, b [4]float32, c [16]byte, imm8 int, rounding int) [4]float32
TEXT ·maskFixupimmRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	// TODO: Code missing

	MOVOU X5, ret+68(FP)
	RET

// func maskzFixupimmRoundSs(k uint8, a [4]float32, b [4]float32, c [16]byte, imm8 int, rounding int) [4]float32
TEXT ·maskzFixupimmRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	// TODO: Code missing

	MOVOU X5, ret+68(FP)
	RET

// func fixupimmSd(a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·fixupimmSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+56(FP)
	RET

// func maskFixupimmSd(a [2]float64, k uint8, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskFixupimmSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFixupimmSd(k uint8, a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskzFixupimmSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func fixupimmSs(a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·fixupimmSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+56(FP)
	RET

// func maskFixupimmSs(a [4]float32, k uint8, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskFixupimmSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFixupimmSs(k uint8, a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskzFixupimmSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func m512FloorPd(a [8]float64) [8]float64
TEXT ·m512FloorPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskFloorPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskFloorPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512FloorPs(a [16]float32) [16]float32
TEXT ·m512FloorPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskFloorPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskFloorPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskFmaddPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmaddPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmaddPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmaddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmaddPd(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskFmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmaddPd(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·m256Mask3FmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmaddPd(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskzFmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFmaddPd(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskzFmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskFmaddPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmaddPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmaddPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmaddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmaddPs(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskFmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmaddPs(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·m256Mask3FmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmaddPs(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskzFmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFmaddPs(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskzFmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFmaddRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskzFmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFmaddRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskzFmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func maskFmaddRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FmaddRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFmaddRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFmaddRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFmaddRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FmaddRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFmaddRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFmaddRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFmaddSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmaddSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmaddSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmaddSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFmaddSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmaddSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmaddSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmaddSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFmaddsubPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmaddsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmaddsubPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmaddsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmaddsubPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmaddsubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmaddsubPd(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskFmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmaddsubPd(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·m256Mask3FmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmaddsubPd(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskzFmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512FmaddsubPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512FmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmaddsubPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskFmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmaddsubPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·m512Mask3FmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFmaddsubPd(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskzFmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskFmaddsubPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmaddsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmaddsubPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmaddsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmaddsubPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmaddsubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmaddsubPs(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskFmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmaddsubPs(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·m256Mask3FmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmaddsubPs(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskzFmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512FmaddsubPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512FmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmaddsubPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskFmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmaddsubPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·m512Mask3FmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFmaddsubPs(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskzFmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmaddsubRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512FmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmaddsubRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskFmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmaddsubRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·m512Mask3FmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFmaddsubRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskzFmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FmaddsubRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512FmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmaddsubRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskFmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmaddsubRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·m512Mask3FmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFmaddsubRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskzFmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func maskFmsubPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmsubPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmsubPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmsubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmsubPd(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskFmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmsubPd(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·m256Mask3FmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmsubPd(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskzFmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFmsubPd(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskzFmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskFmsubPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmsubPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmsubPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmsubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmsubPs(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskFmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmsubPs(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·m256Mask3FmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmsubPs(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskzFmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFmsubPs(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskzFmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFmsubRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskzFmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFmsubRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskzFmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func maskFmsubRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FmsubRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFmsubRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFmsubRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFmsubRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FmsubRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFmsubRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFmsubRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFmsubSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmsubSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmsubSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmsubSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFmsubSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmsubSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmsubSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmsubSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFmsubaddPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmsubaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmsubaddPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmsubaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmsubaddPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmsubaddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmsubaddPd(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskFmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmsubaddPd(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·m256Mask3FmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmsubaddPd(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskzFmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512FmsubaddPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512FmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmsubaddPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskFmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmsubaddPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·m512Mask3FmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFmsubaddPd(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskzFmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskFmsubaddPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmsubaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FmsubaddPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmsubaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFmsubaddPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmsubaddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFmsubaddPs(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskFmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FmsubaddPs(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·m256Mask3FmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFmsubaddPs(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskzFmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512FmsubaddPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512FmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmsubaddPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskFmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmsubaddPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·m512Mask3FmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFmsubaddPs(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskzFmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmsubaddRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512FmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmsubaddRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskFmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmsubaddRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·m512Mask3FmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFmsubaddRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskzFmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FmsubaddRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512FmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmsubaddRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskFmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmsubaddRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·m512Mask3FmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFmsubaddRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskzFmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func maskFnmaddPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmaddPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmaddPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmaddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFnmaddPd(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskFnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FnmaddPd(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·m256Mask3FnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFnmaddPd(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskzFnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFnmaddPd(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskzFnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskFnmaddPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmaddPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmaddPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmaddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFnmaddPs(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskFnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FnmaddPs(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·m256Mask3FnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFnmaddPs(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskzFnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFnmaddPs(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskzFnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFnmaddRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskzFnmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFnmaddRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskzFnmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func maskFnmaddRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFnmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FnmaddRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FnmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFnmaddRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFnmaddRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFnmaddRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFnmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FnmaddRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FnmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFnmaddRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFnmaddRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFnmaddSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmaddSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmaddSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmaddSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFnmaddSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmaddSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmaddSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmaddSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFnmsubPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmsubPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmsubPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmsubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFnmsubPd(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskFnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FnmsubPd(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·m256Mask3FnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFnmsubPd(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·m256MaskzFnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFnmsubPd(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskzFnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskFnmsubPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmsubPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmsubPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmsubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func m256MaskFnmsubPs(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskFnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask3FnmsubPs(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·m256Mask3FnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzFnmsubPs(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·m256MaskzFnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzFnmsubPs(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskzFnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzFnmsubRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskzFnmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzFnmsubRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskzFnmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func maskFnmsubRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFnmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FnmsubRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FnmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFnmsubRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFnmsubRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFnmsubRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFnmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func mask3FnmsubRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FnmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzFnmsubRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFnmsubRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskFnmsubSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmsubSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmsubSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmsubSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskFnmsubSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask3FnmsubSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzFnmsubSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmsubSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func getexpPd(a [2]float64) [2]float64
TEXT ·getexpPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VGETEXPPD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskGetexpPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskGetexpPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzGetexpPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzGetexpPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VGETEXPPD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256GetexpPd(a [4]float64) [4]float64
TEXT ·m256GetexpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VGETEXPPD Y0

	MOV Y0, ret+0(FP)
	RET

// func m256MaskGetexpPd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskGetexpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzGetexpPd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzGetexpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VGETEXPPD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskzGetexpPd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzGetexpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VGETEXPPD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func getexpPs(a [4]float32) [4]float32
TEXT ·getexpPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VGETEXPPS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskGetexpPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskGetexpPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzGetexpPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzGetexpPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VGETEXPPS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256GetexpPs(a [8]float32) [8]float32
TEXT ·m256GetexpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VGETEXPPS Y0

	MOV Y0, ret+0(FP)
	RET

// func m256MaskGetexpPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskGetexpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzGetexpPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzGetexpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VGETEXPPS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskzGetexpPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzGetexpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGETEXPPS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskzGetexpRoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskzGetexpRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzGetexpRoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskzGetexpRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func getexpRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·getexpRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskGetexpRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskGetexpRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzGetexpRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzGetexpRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func getexpRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·getexpRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskGetexpRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskGetexpRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzGetexpRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzGetexpRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func getexpSd(a [2]float64, b [2]float64) [2]float64
TEXT ·getexpSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VGETEXPSD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskGetexpSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskGetexpSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzGetexpSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzGetexpSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func getexpSs(a [4]float32, b [4]float32) [4]float32
TEXT ·getexpSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VGETEXPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskGetexpSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskGetexpSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzGetexpSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzGetexpSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func getmantPd(a [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·getmantPd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func maskGetmantPd(src [2]float64, k uint8, a [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskGetmantPd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func maskzGetmantPd(k uint8, a [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskzGetmantPd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m256GetmantPd(a [4]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float64
TEXT ·m256GetmantPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskGetmantPd(src [4]float64, k uint8, a [4]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float64
TEXT ·m256MaskGetmantPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzGetmantPd(k uint8, a [4]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float64
TEXT ·m256MaskzGetmantPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzGetmantPd(k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float64
TEXT ·m512MaskzGetmantPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func getmantPs(a [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·getmantPs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func maskGetmantPs(src [4]float32, k uint8, a [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskGetmantPs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func maskzGetmantPs(k uint8, a [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskzGetmantPs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func m256GetmantPs(a [8]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float32
TEXT ·m256GetmantPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskGetmantPs(src [8]float32, k uint8, a [8]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float32
TEXT ·m256MaskGetmantPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzGetmantPs(k uint8, a [8]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float32
TEXT ·m256MaskzGetmantPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskzGetmantPs(k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [16]float32
TEXT ·m512MaskzGetmantPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzGetmantRoundPd(k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [8]float64
TEXT ·m512MaskzGetmantRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzGetmantRoundPs(k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [16]float32
TEXT ·m512MaskzGetmantRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func getmantRoundSd(a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [2]float64
TEXT ·getmantRoundSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func maskGetmantRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [2]float64
TEXT ·maskGetmantRoundSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X6, ret+0(FP)
	RET

// func maskzGetmantRoundSd(k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [2]float64
TEXT ·maskzGetmantRoundSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X5, ret+0(FP)
	RET

// func getmantRoundSs(a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [4]float32
TEXT ·getmantRoundSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func maskGetmantRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [4]float32
TEXT ·maskGetmantRoundSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X6, ret+0(FP)
	RET

// func maskzGetmantRoundSs(k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [4]float32
TEXT ·maskzGetmantRoundSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X5, ret+0(FP)
	RET

// func getmantSd(a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·getmantSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func maskGetmantSd(src [2]float64, k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskGetmantSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X5, ret+0(FP)
	RET

// func maskzGetmantSd(k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskzGetmantSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func getmantSs(a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·getmantSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func maskGetmantSs(src [4]float32, k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskGetmantSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X5, ret+0(FP)
	RET

// func maskzGetmantSs(k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskzGetmantSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func m512HypotPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512HypotPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskHypotPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskHypotPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512HypotPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512HypotPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskHypotPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskHypotPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func mmaskI32gatherEpi32(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI32gatherEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI32gatherEpi32(src [32]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [32]byte
TEXT ·m256MmaskI32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func mmaskI32gatherEpi64(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI32gatherEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI32gatherEpi64(src [32]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [32]byte
TEXT ·m256MmaskI32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m512I32gatherEpi64(vindex [32]byte, base_addr uintptr, scale int) [64]byte
TEXT ·m512I32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI32gatherEpi64(src [64]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [64]byte
TEXT ·m512MaskI32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func mmaskI32gatherPd(src [2]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI32gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI32gatherPd(src [4]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float64
TEXT ·m256MmaskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m512I32gatherPd(vindex [32]byte, base_addr uintptr, scale int) [8]float64
TEXT ·m512I32gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI32gatherPd(src [8]float64, k uint8, vindex [32]byte, base_addr uintptr, scale int) [8]float64
TEXT ·m512MaskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func mmaskI32gatherPs(src [4]float32, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI32gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI32gatherPs(src [8]float32, k uint8, vindex [32]byte, base_addr uintptr, scale int) [8]float32
TEXT ·m256MmaskI32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func i32scatterEpi32(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i32scatterEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI32scatterEpi32(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI32scatterEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I32scatterEpi32(base_addr uintptr, vindex [32]byte, a [32]byte, scale int) 
TEXT ·m256I32scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI32scatterEpi32(base_addr uintptr, k uint8, vindex [32]byte, a [32]byte, scale int) 
TEXT ·m256MaskI32scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func i32scatterEpi64(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i32scatterEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI32scatterEpi64(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI32scatterEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I32scatterEpi64(base_addr uintptr, vindex [16]byte, a [32]byte, scale int) 
TEXT ·m256I32scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI32scatterEpi64(base_addr uintptr, k uint8, vindex [16]byte, a [32]byte, scale int) 
TEXT ·m256MaskI32scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512I32scatterEpi64(base_addr uintptr, vindex [32]byte, a [64]byte, scale int) 
TEXT ·m512I32scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskI32scatterEpi64(base_addr uintptr, k uint8, vindex [32]byte, a [64]byte, scale int) 
TEXT ·m512MaskI32scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func i32scatterPd(base_addr uintptr, vindex [16]byte, a [2]float64, scale int) 
TEXT ·i32scatterPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI32scatterPd(base_addr uintptr, k uint8, vindex [16]byte, a [2]float64, scale int) 
TEXT ·maskI32scatterPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I32scatterPd(base_addr uintptr, vindex [16]byte, a [4]float64, scale int) 
TEXT ·m256I32scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	RET

// func m256MaskI32scatterPd(base_addr uintptr, k uint8, vindex [16]byte, a [4]float64, scale int) 
TEXT ·m256MaskI32scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	RET

// func m512I32scatterPd(base_addr uintptr, vindex [32]byte, a [8]float64, scale int) 
TEXT ·m512I32scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskI32scatterPd(base_addr uintptr, k uint8, vindex [32]byte, a [8]float64, scale int) 
TEXT ·m512MaskI32scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func i32scatterPs(base_addr uintptr, vindex [16]byte, a [4]float32, scale int) 
TEXT ·i32scatterPs(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI32scatterPs(base_addr uintptr, k uint8, vindex [16]byte, a [4]float32, scale int) 
TEXT ·maskI32scatterPs(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I32scatterPs(base_addr uintptr, vindex [32]byte, a [8]float32, scale int) 
TEXT ·m256I32scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI32scatterPs(base_addr uintptr, k uint8, vindex [32]byte, a [8]float32, scale int) 
TEXT ·m256MaskI32scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func mmaskI64gatherEpi32(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI64gatherEpi32(src [16]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [16]byte
TEXT ·m256MmaskI64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func m512I64gatherEpi32(vindex [64]byte, base_addr uintptr, scale int) [32]byte
TEXT ·m512I64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskI64gatherEpi32(src [32]byte, k uint8, vindex [64]byte, base_addr uintptr, scale int) [32]byte
TEXT ·m512MaskI64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func mmaskI64gatherEpi64(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI64gatherEpi64(src [32]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [32]byte
TEXT ·m256MmaskI64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m512I64gatherEpi64(vindex [64]byte, base_addr uintptr, scale int) [64]byte
TEXT ·m512I64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI64gatherEpi64(src [64]byte, k uint8, vindex [64]byte, base_addr uintptr, scale int) [64]byte
TEXT ·m512MaskI64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func mmaskI64gatherPd(src [2]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI64gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI64gatherPd(src [4]float64, k uint8, vindex [32]byte, base_addr uintptr, scale int) [4]float64
TEXT ·m256MmaskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m512I64gatherPd(vindex [64]byte, base_addr uintptr, scale int) [8]float64
TEXT ·m512I64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI64gatherPd(src [8]float64, k uint8, vindex [64]byte, base_addr uintptr, scale int) [8]float64
TEXT ·m512MaskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func mmaskI64gatherPs(src [4]float32, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI64gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVQ base_addr+36(FP),R11
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func m256MmaskI64gatherPs(src [4]float32, k uint8, vindex [32]byte, base_addr uintptr, scale int) [4]float32
TEXT ·m256MmaskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVOU X4, ret+0(FP)
	RET

// func m512I64gatherPs(vindex [64]byte, base_addr uintptr, scale int) [8]float32
TEXT ·m512I64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskI64gatherPs(src [8]float32, k uint8, vindex [64]byte, base_addr uintptr, scale int) [8]float32
TEXT ·m512MaskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func i64scatterEpi32(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i64scatterEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI64scatterEpi32(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI64scatterEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I64scatterEpi32(base_addr uintptr, vindex [32]byte, a [16]byte, scale int) 
TEXT ·m256I64scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI64scatterEpi32(base_addr uintptr, k uint8, vindex [32]byte, a [16]byte, scale int) 
TEXT ·m256MaskI64scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512I64scatterEpi32(base_addr uintptr, vindex [64]byte, a [32]byte, scale int) 
TEXT ·m512I64scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64scatterEpi32(base_addr uintptr, k uint8, vindex [64]byte, a [32]byte, scale int) 
TEXT ·m512MaskI64scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func i64scatterEpi64(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i64scatterEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI64scatterEpi64(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI64scatterEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I64scatterEpi64(base_addr uintptr, vindex [32]byte, a [32]byte, scale int) 
TEXT ·m256I64scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI64scatterEpi64(base_addr uintptr, k uint8, vindex [32]byte, a [32]byte, scale int) 
TEXT ·m256MaskI64scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512I64scatterEpi64(base_addr uintptr, vindex [64]byte, a [64]byte, scale int) 
TEXT ·m512I64scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64scatterEpi64(base_addr uintptr, k uint8, vindex [64]byte, a [64]byte, scale int) 
TEXT ·m512MaskI64scatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func i64scatterPd(base_addr uintptr, vindex [16]byte, a [2]float64, scale int) 
TEXT ·i64scatterPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI64scatterPd(base_addr uintptr, k uint8, vindex [16]byte, a [2]float64, scale int) 
TEXT ·maskI64scatterPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I64scatterPd(base_addr uintptr, vindex [32]byte, a [4]float64, scale int) 
TEXT ·m256I64scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI64scatterPd(base_addr uintptr, k uint8, vindex [32]byte, a [4]float64, scale int) 
TEXT ·m256MaskI64scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512I64scatterPd(base_addr uintptr, vindex [64]byte, a [8]float64, scale int) 
TEXT ·m512I64scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64scatterPd(base_addr uintptr, k uint8, vindex [64]byte, a [8]float64, scale int) 
TEXT ·m512MaskI64scatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func i64scatterPs(base_addr uintptr, vindex [16]byte, a [4]float32, scale int) 
TEXT ·i64scatterPs(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVOU a+24(FP),X2
	MOVQ scale+40(FP),R11

	// TODO: Code missing

	RET

// func maskI64scatterPs(base_addr uintptr, k uint8, vindex [16]byte, a [4]float32, scale int) 
TEXT ·maskI64scatterPs(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU vindex+12(FP),X2
	MOVOU a+28(FP),X3
	MOVQ scale+44(FP),R12

	// TODO: Code missing

	RET

// func m256I64scatterPs(base_addr uintptr, vindex [32]byte, a [4]float32, scale int) 
TEXT ·m256I64scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m256MaskI64scatterPs(base_addr uintptr, k uint8, vindex [32]byte, a [4]float32, scale int) 
TEXT ·m256MaskI64scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512I64scatterPs(base_addr uintptr, vindex [64]byte, a [8]float32, scale int) 
TEXT ·m512I64scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64scatterPs(base_addr uintptr, k uint8, vindex [64]byte, a [8]float32, scale int) 
TEXT ·m512MaskI64scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m256Insertf32x4(a [8]float32, b [4]float32, imm8 int) [8]float32
TEXT ·m256Insertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskInsertf32x4(src [8]float32, k uint8, a [8]float32, b [4]float32, imm8 int) [8]float32
TEXT ·m256MaskInsertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzInsertf32x4(k uint8, a [8]float32, b [4]float32, imm8 int) [8]float32
TEXT ·m256MaskzInsertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512Insertf32x4(a [16]float32, b [4]float32, imm8 int) [16]float32
TEXT ·m512Insertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskInsertf32x4(src [16]float32, k uint16, a [16]float32, b [4]float32, imm8 int) [16]float32
TEXT ·m512MaskInsertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzInsertf32x4(k uint16, a [16]float32, b [4]float32, imm8 int) [16]float32
TEXT ·m512MaskzInsertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Insertf64x4(a [8]float64, b [4]float64, imm8 int) [8]float64
TEXT ·m512Insertf64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskInsertf64x4(src [8]float64, k uint8, a [8]float64, b [4]float64, imm8 int) [8]float64
TEXT ·m512MaskInsertf64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzInsertf64x4(k uint8, a [8]float64, b [4]float64, imm8 int) [8]float64
TEXT ·m512MaskzInsertf64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m256Inserti32x4(a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·m256Inserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskInserti32x4(src [32]byte, k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·m256MaskInserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzInserti32x4(k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·m256MaskzInserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512Inserti32x4(a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·m512Inserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskInserti32x4(src [64]byte, k uint16, a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·m512MaskInserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzInserti32x4(k uint16, a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·m512MaskzInserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Inserti64x4(a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·m512Inserti64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskInserti64x4(src [64]byte, k uint8, a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·m512MaskInserti64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzInserti64x4(k uint8, a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·m512MaskzInserti64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512InvsqrtPd(a [8]float64) [8]float64
TEXT ·m512InvsqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskInvsqrtPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskInvsqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512InvsqrtPs(a [16]float32) [16]float32
TEXT ·m512InvsqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskInvsqrtPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskInvsqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Kand(a uint16, b uint16) uint16
TEXT ·m512Kand(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KANDW R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kandn(a uint16, b uint16) uint16
TEXT ·m512Kandn(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KANDNW R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kmov(a uint16) uint16
TEXT ·m512Kmov(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// KMOVW R8

	MOVW $0, ret+4(FP)
	RET

// func m512Knot(a uint16) uint16
TEXT ·m512Knot(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// KNOTW R8

	MOVW $0, ret+4(FP)
	RET

// func m512Kor(a uint16, b uint16) uint16
TEXT ·m512Kor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KORW R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kortestc(k1 uint16, k2 uint16) int
TEXT ·m512Kortestc(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KORTESTW R8, R9

	MOVQ $0, ret+8(FP)
	RET

// func m512Kortestz(k1 uint16, k2 uint16) int
TEXT ·m512Kortestz(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KORTESTW R8, R9

	MOVQ $0, ret+8(FP)
	RET

// func m512Kunpackb(a uint16, b uint16) uint16
TEXT ·m512Kunpackb(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KUNPCKBW R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kxnor(a uint16, b uint16) uint16
TEXT ·m512Kxnor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KXNORW R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kxor(a uint16, b uint16) uint16
TEXT ·m512Kxor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KXORW R8, R9

	MOVW $0, ret+8(FP)
	RET

// func maskLoadEpi32(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoadEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoadEpi32(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoadEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoadEpi32(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskLoadEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoadEpi32(k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskzLoadEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskzLoadEpi32(k uint16, mem_addr uintptr) [64]byte
TEXT ·m512MaskzLoadEpi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoadEpi64(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoadEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoadEpi64(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoadEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoadEpi64(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskLoadEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoadEpi64(k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskzLoadEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskzLoadEpi64(k uint8, mem_addr uintptr) [64]byte
TEXT ·m512MaskzLoadEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoadPd(src [2]float64, k uint8, mem_addr uintptr) [2]float64
TEXT ·maskLoadPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoadPd(k uint8, mem_addr uintptr) [2]float64
TEXT ·maskzLoadPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoadPd(src [4]float64, k uint8, mem_addr uintptr) [4]float64
TEXT ·m256MaskLoadPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoadPd(k uint8, mem_addr uintptr) [4]float64
TEXT ·m256MaskzLoadPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskzLoadPd(k uint8, mem_addr uintptr) [8]float64
TEXT ·m512MaskzLoadPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoadPs(src [4]float32, k uint8, mem_addr uintptr) [4]float32
TEXT ·maskLoadPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoadPs(k uint8, mem_addr uintptr) [4]float32
TEXT ·maskzLoadPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoadPs(src [8]float32, k uint8, mem_addr uintptr) [8]float32
TEXT ·m256MaskLoadPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoadPs(k uint8, mem_addr uintptr) [8]float32
TEXT ·m256MaskzLoadPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskzLoadPs(k uint16, mem_addr uintptr) [16]float32
TEXT ·m512MaskzLoadPs(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoadSd(src [2]float64, k uint8, mem_addr float64) [2]float64
TEXT ·maskLoadSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoadSd(k uint8, mem_addr float64) [2]float64
TEXT ·maskzLoadSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVSD R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func maskLoadSs(src [4]float32, k uint8, mem_addr float32) [4]float32
TEXT ·maskLoadSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVL mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func maskzLoadSs(k uint8, mem_addr float32) [4]float32
TEXT ·maskzLoadSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVL mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVSS R8, R9

	MOVOU X1, ret+8(FP)
	RET

// func maskLoaduEpi32(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoaduEpi32(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQU32 R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoaduEpi32(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskLoaduEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoaduEpi32(k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskzLoaduEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQU32 R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskLoaduEpi32(src [64]byte, k uint16, mem_addr uintptr) [64]byte
TEXT ·m512MaskLoaduEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzLoaduEpi32(k uint16, mem_addr uintptr) [64]byte
TEXT ·m512MaskzLoaduEpi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQU32 R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoaduEpi64(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoaduEpi64(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQU64 R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoaduEpi64(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskLoaduEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoaduEpi64(k uint8, mem_addr uintptr) [32]byte
TEXT ·m256MaskzLoaduEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQU64 R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskLoaduEpi64(src [64]byte, k uint8, mem_addr uintptr) [64]byte
TEXT ·m512MaskLoaduEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzLoaduEpi64(k uint8, mem_addr uintptr) [64]byte
TEXT ·m512MaskzLoaduEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVDQU64 R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoaduPd(src [2]float64, k uint8, mem_addr uintptr) [2]float64
TEXT ·maskLoaduPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoaduPd(k uint8, mem_addr uintptr) [2]float64
TEXT ·maskzLoaduPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVUPD R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoaduPd(src [4]float64, k uint8, mem_addr uintptr) [4]float64
TEXT ·m256MaskLoaduPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoaduPd(k uint8, mem_addr uintptr) [4]float64
TEXT ·m256MaskzLoaduPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVUPD R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512LoaduPd(mem_addr uintptr) [8]float64
TEXT ·m512LoaduPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVUPD R8

	MOV Z0, ret+8(FP)
	RET

// func m512MaskLoaduPd(src [8]float64, k uint8, mem_addr uintptr) [8]float64
TEXT ·m512MaskLoaduPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzLoaduPd(k uint8, mem_addr uintptr) [8]float64
TEXT ·m512MaskzLoaduPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVUPD R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskLoaduPs(src [4]float32, k uint8, mem_addr uintptr) [4]float32
TEXT ·maskLoaduPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzLoaduPs(k uint8, mem_addr uintptr) [4]float32
TEXT ·maskzLoaduPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVUPS R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskLoaduPs(src [8]float32, k uint8, mem_addr uintptr) [8]float32
TEXT ·m256MaskLoaduPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzLoaduPs(k uint8, mem_addr uintptr) [8]float32
TEXT ·m256MaskzLoaduPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVUPS R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512LoaduPs(mem_addr uintptr) [16]float32
TEXT ·m512LoaduPs(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVUPS R8

	MOV Z0, ret+8(FP)
	RET

// func m512MaskLoaduPs(src [16]float32, k uint16, mem_addr uintptr) [16]float32
TEXT ·m512MaskLoaduPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzLoaduPs(k uint16, mem_addr uintptr) [16]float32
TEXT ·m512MaskzLoaduPs(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VMOVUPS R8, R9

	MOV Z1, ret+12(FP)
	RET

// func m512LoaduSi512(mem_addr uintptr) [64]byte
TEXT ·m512LoaduSi512(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVDQU32 R8

	MOV Z0, ret+8(FP)
	RET

// func m512LogPd(a [8]float64) [8]float64
TEXT ·m512LogPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLogPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskLogPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LogPs(a [16]float32) [16]float32
TEXT ·m512LogPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLogPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskLogPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log10Pd(a [8]float64) [8]float64
TEXT ·m512Log10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog10Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskLog10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log10Ps(a [16]float32) [16]float32
TEXT ·m512Log10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog10Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskLog10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log1pPd(a [8]float64) [8]float64
TEXT ·m512Log1pPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog1pPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskLog1pPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log1pPs(a [16]float32) [16]float32
TEXT ·m512Log1pPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog1pPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskLog1pPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log2Pd(a [8]float64) [8]float64
TEXT ·m512Log2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog2Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskLog2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LogbPd(a [8]float64) [8]float64
TEXT ·m512LogbPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLogbPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskLogbPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LogbPs(a [16]float32) [16]float32
TEXT ·m512LogbPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLogbPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskLogbPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMaxEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMaxEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMaxEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMaxEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maxEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMAXSQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskMaxEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMaxEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaxEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMAXSQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMaxEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMaxEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaxEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMAXSQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMaxEpu32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpu32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxEpu32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMaxEpu32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMaxEpu32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMaxEpu64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpu64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maxEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMAXUQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskMaxEpu64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMaxEpu64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaxEpu64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMAXUQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMaxEpu64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMaxEpu64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaxEpu64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMAXUQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMaxPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMaxPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMaxPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskMaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMaxPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzMaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskMaxPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskMaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMaxPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzMaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaxPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMAXPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMaxPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMaxPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMaxPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskMaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMaxPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzMaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskMaxPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskMaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMaxPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzMaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaxPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMAXPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMaxRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·m512MaskMaxRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzMaxRoundPd(k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·m512MaskzMaxRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaxRoundPd(a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·m512MaxRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMaxRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·m512MaskMaxRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzMaxRoundPs(k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·m512MaskzMaxRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaxRoundPs(a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·m512MaxRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMaxRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskMaxRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzMaxRoundSd(k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskzMaxRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maxRoundSd(a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maxRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskMaxRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskMaxRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzMaxRoundSs(k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskzMaxRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maxRoundSs(a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maxRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskMaxSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMaxSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMaxSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMaxSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMaxSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMaxSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMinEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMinEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMinEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMinEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func minEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMINSQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskMinEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMinEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MinEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMINSQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMinEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMinEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MinEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMINSQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMinEpu32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpu32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinEpu32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMinEpu32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMinEpu32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMinEpu64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpu64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func minEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPMINUQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskMinEpu64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMinEpu64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MinEpu64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPMINUQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMinEpu64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMinEpu64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MinEpu64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMINUQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMinPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMinPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMinPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskMinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMinPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzMinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskMinPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskMinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMinPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzMinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MinPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512MinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMINPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMinPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMinPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMinPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskMinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMinPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzMinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskMinPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskMinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMinPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzMinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MinPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512MinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMINPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMinRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·m512MaskMinRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzMinRoundPd(k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·m512MaskzMinRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MinRoundPd(a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·m512MinRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMinRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·m512MaskMinRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzMinRoundPs(k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·m512MaskzMinRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MinRoundPs(a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·m512MinRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMinRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskMinRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzMinRoundSd(k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskzMinRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func minRoundSd(a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·minRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskMinRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskMinRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzMinRoundSs(k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskzMinRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func minRoundSs(a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·minRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskMinSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMinSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMinSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMinSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMinSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMinSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMovEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskMovEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzMovEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovEpi32(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskMovEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMovEpi32(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzMovEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskzMovEpi32(k uint16, a [64]byte) [64]byte
TEXT ·m512MaskzMovEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMovEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskMovEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzMovEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovEpi64(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·m256MaskMovEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMovEpi64(k uint8, a [32]byte) [32]byte
TEXT ·m256MaskzMovEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskzMovEpi64(k uint8, a [64]byte) [64]byte
TEXT ·m512MaskzMovEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMovPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskMovPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzMovPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovPd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskMovPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMovPd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzMovPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskzMovPd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzMovPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMovPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskMovPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzMovPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskMovPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMovPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzMovPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskzMovPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzMovPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMoveSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMoveSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMoveSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMoveSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMoveSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMoveSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMoveSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMoveSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMovedupPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskMovedupPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovedupPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzMovedupPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVDDUP R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovedupPd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskMovedupPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMovedupPd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzMovedupPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VMOVDDUP R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMovedupPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskMovedupPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzMovedupPd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzMovedupPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVDDUP R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MovedupPd(a [8]float64) [8]float64
TEXT ·m512MovedupPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVDDUP Z0

	MOV Z0, ret+0(FP)
	RET

// func maskMovehdupPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskMovehdupPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovehdupPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzMovehdupPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVSHDUP R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovehdupPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskMovehdupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMovehdupPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzMovehdupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VMOVSHDUP R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMovehdupPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskMovehdupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzMovehdupPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzMovehdupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVSHDUP R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MovehdupPs(a [16]float32) [16]float32
TEXT ·m512MovehdupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVSHDUP Z0

	MOV Z0, ret+0(FP)
	RET

// func maskMoveldupPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskMoveldupPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzMoveldupPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzMoveldupPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VMOVSLDUP R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMoveldupPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskMoveldupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzMoveldupPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzMoveldupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VMOVSLDUP R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskMoveldupPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskMoveldupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzMoveldupPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzMoveldupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVSLDUP R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MoveldupPs(a [16]float32) [16]float32
TEXT ·m512MoveldupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVSLDUP Z0

	MOV Z0, ret+0(FP)
	RET

// func maskMulEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMulEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskMulEpi32(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMulEpi32(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MulEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMULDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMulEpu32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulEpu32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulEpu32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMulEpu32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskMulEpu32(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMulEpu32(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MulEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMULUDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskMulPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMulPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMulPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskMulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMulPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzMulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMulPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzMulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskMulPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMulPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMulPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskMulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMulPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzMulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMulPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzMulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzMulRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskzMulRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMulRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskzMulRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskMulRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskMulRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzMulRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzMulRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func mulRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·mulRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskMulRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskMulRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzMulRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzMulRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func mulRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·mulRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskMulSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMulSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMulSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMulSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMulSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMulSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskMulloEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulloEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulloEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulloEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulloEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMulloEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzMulloEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMulloxEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulloxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MulloxEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulloxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskNearbyintPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskNearbyintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512NearbyintPd(a [8]float64) [8]float64
TEXT ·m512NearbyintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskNearbyintPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskNearbyintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512NearbyintPs(a [16]float32) [16]float32
TEXT ·m512NearbyintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskOrEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskOrEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzOrEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzOrEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskOrEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskOrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzOrEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzOrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzOrEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzOrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskOrEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskOrEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzOrEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzOrEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskOrEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskOrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzOrEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzOrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzOrEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzOrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskPermutePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskPermutePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzPermutePd(k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskzPermutePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func m256MaskPermutePd(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·m256MaskPermutePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutePd(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·m256MaskzPermutePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutePd(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·m512MaskPermutePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutePd(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·m512MaskzPermutePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutePd(a [8]float64, imm8 int) [8]float64
TEXT ·m512PermutePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VPERMILPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskPermutePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskPermutePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzPermutePs(k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskzPermutePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func m256MaskPermutePs(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·m256MaskPermutePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutePs(k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·m256MaskzPermutePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutePs(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·m512MaskPermutePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutePs(k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·m512MaskzPermutePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutePs(a [16]float32, imm8 int) [16]float32
TEXT ·m512PermutePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VPERMILPS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskPermutevarPd(src [2]float64, k uint8, a [2]float64, b [16]byte) [2]float64
TEXT ·maskPermutevarPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutevarPd(k uint8, a [2]float64, b [16]byte) [2]float64
TEXT ·maskzPermutevarPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskPermutevarPd(src [4]float64, k uint8, a [4]float64, b [32]byte) [4]float64
TEXT ·m256MaskPermutevarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutevarPd(k uint8, a [4]float64, b [32]byte) [4]float64
TEXT ·m256MaskzPermutevarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutevarPd(src [8]float64, k uint8, a [8]float64, b [64]byte) [8]float64
TEXT ·m512MaskPermutevarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutevarPd(k uint8, a [8]float64, b [64]byte) [8]float64
TEXT ·m512MaskzPermutevarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutevarPd(a [8]float64, b [64]byte) [8]float64
TEXT ·m512PermutevarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VPERMILPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskPermutevarPs(src [4]float32, k uint8, a [4]float32, b [16]byte) [4]float32
TEXT ·maskPermutevarPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutevarPs(k uint8, a [4]float32, b [16]byte) [4]float32
TEXT ·maskzPermutevarPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskPermutevarPs(src [8]float32, k uint8, a [8]float32, b [32]byte) [8]float32
TEXT ·m256MaskPermutevarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutevarPs(k uint8, a [8]float32, b [32]byte) [8]float32
TEXT ·m256MaskzPermutevarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutevarPs(src [16]float32, k uint16, a [16]float32, b [64]byte) [16]float32
TEXT ·m512MaskPermutevarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutevarPs(k uint16, a [16]float32, b [64]byte) [16]float32
TEXT ·m512MaskzPermutevarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutevarPs(a [16]float32, b [64]byte) [16]float32
TEXT ·m512PermutevarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VPERMILPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m256MaskPermutexEpi64(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskPermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexEpi64(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzPermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·m256PermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPERMQ Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskPermutexEpi64(src [64]byte, k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskPermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexEpi64(k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzPermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexEpi64(a [64]byte, imm8 int) [64]byte
TEXT ·m512PermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m256MaskPermutexPd(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·m256MaskPermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexPd(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·m256MaskzPermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexPd(a [4]float64, imm8 int) [4]float64
TEXT ·m256PermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VPERMPD Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskPermutexPd(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·m512MaskPermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexPd(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·m512MaskzPermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexPd(a [8]float64, imm8 int) [8]float64
TEXT ·m512PermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VPERMPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskPermutex2varEpi32(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask2Permutex2varEpi32(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutex2varEpi32(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func permutex2varEpi32(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

// func m256MaskPermutex2varEpi32(a [32]byte, k uint8, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPermutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask2Permutex2varEpi32(a [32]byte, idx [32]byte, k uint8, b [32]byte) [32]byte
TEXT ·m256Mask2Permutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutex2varEpi32(k uint8, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPermutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Permutex2varEpi32(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256Permutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutex2varEpi32(a [64]byte, k uint16, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPermutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask2Permutex2varEpi32(a [64]byte, idx [64]byte, k uint16, b [64]byte) [64]byte
TEXT ·m512Mask2Permutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutex2varEpi32(k uint16, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPermutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permutex2varEpi32(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512Permutex2varEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskPermutex2varEpi64(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask2Permutex2varEpi64(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutex2varEpi64(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func permutex2varEpi64(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

// func m256MaskPermutex2varEpi64(a [32]byte, k uint8, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPermutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask2Permutex2varEpi64(a [32]byte, idx [32]byte, k uint8, b [32]byte) [32]byte
TEXT ·m256Mask2Permutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutex2varEpi64(k uint8, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPermutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Permutex2varEpi64(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256Permutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutex2varEpi64(a [64]byte, k uint8, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPermutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask2Permutex2varEpi64(a [64]byte, idx [64]byte, k uint8, b [64]byte) [64]byte
TEXT ·m512Mask2Permutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutex2varEpi64(k uint8, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPermutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permutex2varEpi64(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512Permutex2varEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskPermutex2varPd(a [2]float64, k uint8, idx [16]byte, b [2]float64) [2]float64
TEXT ·maskPermutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask2Permutex2varPd(a [2]float64, idx [16]byte, k uint8, b [2]float64) [2]float64
TEXT ·mask2Permutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutex2varPd(k uint8, a [2]float64, idx [16]byte, b [2]float64) [2]float64
TEXT ·maskzPermutex2varPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func permutex2varPd(a [2]float64, idx [16]byte, b [2]float64) [2]float64
TEXT ·permutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

// func m256MaskPermutex2varPd(a [4]float64, k uint8, idx [32]byte, b [4]float64) [4]float64
TEXT ·m256MaskPermutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask2Permutex2varPd(a [4]float64, idx [32]byte, k uint8, b [4]float64) [4]float64
TEXT ·m256Mask2Permutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutex2varPd(k uint8, a [4]float64, idx [32]byte, b [4]float64) [4]float64
TEXT ·m256MaskzPermutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Permutex2varPd(a [4]float64, idx [32]byte, b [4]float64) [4]float64
TEXT ·m256Permutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutex2varPd(a [8]float64, k uint8, idx [64]byte, b [8]float64) [8]float64
TEXT ·m512MaskPermutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask2Permutex2varPd(a [8]float64, idx [64]byte, k uint8, b [8]float64) [8]float64
TEXT ·m512Mask2Permutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutex2varPd(k uint8, a [8]float64, idx [64]byte, b [8]float64) [8]float64
TEXT ·m512MaskzPermutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permutex2varPd(a [8]float64, idx [64]byte, b [8]float64) [8]float64
TEXT ·m512Permutex2varPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskPermutex2varPs(a [4]float32, k uint8, idx [16]byte, b [4]float32) [4]float32
TEXT ·maskPermutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func mask2Permutex2varPs(a [4]float32, idx [16]byte, k uint8, b [4]float32) [4]float32
TEXT ·mask2Permutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutex2varPs(k uint8, a [4]float32, idx [16]byte, b [4]float32) [4]float32
TEXT ·maskzPermutex2varPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func permutex2varPs(a [4]float32, idx [16]byte, b [4]float32) [4]float32
TEXT ·permutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

// func m256MaskPermutex2varPs(a [8]float32, k uint8, idx [32]byte, b [8]float32) [8]float32
TEXT ·m256MaskPermutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Mask2Permutex2varPs(a [8]float32, idx [32]byte, k uint8, b [8]float32) [8]float32
TEXT ·m256Mask2Permutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutex2varPs(k uint8, a [8]float32, idx [32]byte, b [8]float32) [8]float32
TEXT ·m256MaskzPermutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256Permutex2varPs(a [8]float32, idx [32]byte, b [8]float32) [8]float32
TEXT ·m256Permutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskPermutex2varPs(a [16]float32, k uint16, idx [64]byte, b [16]float32) [16]float32
TEXT ·m512MaskPermutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask2Permutex2varPs(a [16]float32, idx [64]byte, k uint16, b [16]float32) [16]float32
TEXT ·m512Mask2Permutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutex2varPs(k uint16, a [16]float32, idx [64]byte, b [16]float32) [16]float32
TEXT ·m512MaskzPermutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permutex2varPs(a [16]float32, idx [64]byte, b [16]float32) [16]float32
TEXT ·m512Permutex2varPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m256MaskPermutexvarEpi32(src [32]byte, k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskPermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexvarEpi32(k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskzPermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexvarEpi32(idx [32]byte, a [32]byte) [32]byte
TEXT ·m256PermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPERMD Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskPermutexvarEpi32(src [64]byte, k uint16, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexvarEpi32(k uint16, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskzPermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexvarEpi32(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m256MaskPermutexvarEpi64(src [32]byte, k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskPermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexvarEpi64(k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskzPermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexvarEpi64(idx [32]byte, a [32]byte) [32]byte
TEXT ·m256PermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPERMQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskPermutexvarEpi64(src [64]byte, k uint8, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexvarEpi64(k uint8, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskzPermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexvarEpi64(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m256MaskPermutexvarPd(src [4]float64, k uint8, idx [32]byte, a [4]float64) [4]float64
TEXT ·m256MaskPermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexvarPd(k uint8, idx [32]byte, a [4]float64) [4]float64
TEXT ·m256MaskzPermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexvarPd(idx [32]byte, a [4]float64) [4]float64
TEXT ·m256PermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPERMPD Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskPermutexvarPd(src [8]float64, k uint8, idx [64]byte, a [8]float64) [8]float64
TEXT ·m512MaskPermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexvarPd(k uint8, idx [64]byte, a [8]float64) [8]float64
TEXT ·m512MaskzPermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexvarPd(idx [64]byte, a [8]float64) [8]float64
TEXT ·m512PermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m256MaskPermutexvarPs(src [8]float32, k uint8, idx [32]byte, a [8]float32) [8]float32
TEXT ·m256MaskPermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexvarPs(k uint8, idx [32]byte, a [8]float32) [8]float32
TEXT ·m256MaskzPermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexvarPs(idx [32]byte, a [8]float32) [8]float32
TEXT ·m256PermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPERMPS Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskPermutexvarPs(src [16]float32, k uint16, idx [64]byte, a [16]float32) [16]float32
TEXT ·m512MaskPermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexvarPs(k uint16, idx [64]byte, a [16]float32) [16]float32
TEXT ·m512MaskzPermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexvarPs(idx [64]byte, a [16]float32) [16]float32
TEXT ·m512PermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskPowPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskPowPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512PowPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512PowPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskPowPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskPowPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512PowPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512PowPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func maskRcp14Pd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskRcp14Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzRcp14Pd(k uint8, a [2]float64) [2]float64
TEXT ·maskzRcp14Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VRCP14PD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func rcp14Pd(a [2]float64) [2]float64
TEXT ·rcp14Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VRCP14PD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func m256MaskRcp14Pd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskRcp14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzRcp14Pd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzRcp14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VRCP14PD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m256Rcp14Pd(a [4]float64) [4]float64
TEXT ·m256Rcp14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VRCP14PD Y0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskRcp14Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskRcp14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRcp14Pd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzRcp14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VRCP14PD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rcp14Pd(a [8]float64) [8]float64
TEXT ·m512Rcp14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VRCP14PD Z0

	MOV Z0, ret+0(FP)
	RET

// func maskRcp14Ps(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskRcp14Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzRcp14Ps(k uint8, a [4]float32) [4]float32
TEXT ·maskzRcp14Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VRCP14PS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func rcp14Ps(a [4]float32) [4]float32
TEXT ·rcp14Ps(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing
	// Could be:
	// VRCP14PS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func m256MaskRcp14Ps(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskRcp14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzRcp14Ps(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzRcp14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VRCP14PS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m256Rcp14Ps(a [8]float32) [8]float32
TEXT ·m256Rcp14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VRCP14PS Y0

	MOV Y0, ret+0(FP)
	RET

// func m512MaskRcp14Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRcp14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRcp14Ps(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzRcp14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRCP14PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rcp14Ps(a [16]float32) [16]float32
TEXT ·m512Rcp14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRCP14PS Z0

	MOV Z0, ret+0(FP)
	RET

// func maskRcp14Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRcp14Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRcp14Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRcp14Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rcp14Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rcp14Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VRCP14SD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskRcp14Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRcp14Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRcp14Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRcp14Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rcp14Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rcp14Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VRCP14SS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m512MaskRecipPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskRecipPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RecipPd(a [8]float64) [8]float64
TEXT ·m512RecipPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskRecipPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRecipPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RecipPs(a [16]float32) [16]float32
TEXT ·m512RecipPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512RemEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskRemEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskRemEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512RemEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512RemEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512RemEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512RemEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskRemEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskRemEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512RemEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512RemEpu64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512RemEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RemEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskRintPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskRintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RintPd(a [8]float64) [8]float64
TEXT ·m512RintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskRintPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RintPs(a [16]float32) [16]float32
TEXT ·m512RintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRolEpi32(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRolEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzRolEpi32(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRolEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func rolEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·rolEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VPROLD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256MaskRolEpi32(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskRolEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRolEpi32(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzRolEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RolEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·m256RolEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPROLD Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRolEpi32(src [64]byte, k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskRolEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRolEpi32(k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzRolEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RolEpi32(a [64]byte, imm8 int) [64]byte
TEXT ·m512RolEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPROLD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRolEpi64(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRolEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzRolEpi64(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRolEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func rolEpi64(a [16]byte, imm8 int) [16]byte
TEXT ·rolEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VPROLQ X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256MaskRolEpi64(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskRolEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRolEpi64(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzRolEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RolEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·m256RolEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPROLQ Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRolEpi64(src [64]byte, k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskRolEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRolEpi64(k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzRolEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RolEpi64(a [64]byte, imm8 int) [64]byte
TEXT ·m512RolEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPROLQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRolvEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRolvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRolvEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRolvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rolvEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·rolvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPROLVD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskRolvEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskRolvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRolvEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzRolvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RolvEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256RolvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPROLVD Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRolvEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskRolvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRolvEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzRolvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RolvEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RolvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPROLVD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskRolvEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRolvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRolvEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRolvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rolvEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·rolvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPROLVQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskRolvEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskRolvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRolvEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzRolvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RolvEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256RolvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPROLVQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRolvEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskRolvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRolvEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzRolvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RolvEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RolvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPROLVQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskRorEpi32(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRorEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzRorEpi32(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRorEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func rorEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·rorEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VPRORD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256MaskRorEpi32(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskRorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRorEpi32(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzRorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RorEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·m256RorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPRORD Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRorEpi32(src [64]byte, k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskRorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRorEpi32(k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzRorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RorEpi32(a [64]byte, imm8 int) [64]byte
TEXT ·m512RorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPRORD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRorEpi64(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRorEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzRorEpi64(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRorEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func rorEpi64(a [16]byte, imm8 int) [16]byte
TEXT ·rorEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VPRORQ X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256MaskRorEpi64(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskRorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRorEpi64(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzRorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RorEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·m256RorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPRORQ Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRorEpi64(src [64]byte, k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskRorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRorEpi64(k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzRorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RorEpi64(a [64]byte, imm8 int) [64]byte
TEXT ·m512RorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPRORQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRorvEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRorvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRorvEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRorvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rorvEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·rorvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPRORVD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskRorvEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskRorvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRorvEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzRorvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RorvEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256RorvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPRORVD Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRorvEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskRorvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRorvEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzRorvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RorvEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RorvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPRORVD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskRorvEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRorvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRorvEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRorvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rorvEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·rorvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPRORVQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskRorvEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskRorvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRorvEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzRorvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RorvEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256RorvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPRORVQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRorvEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskRorvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRorvEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzRorvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RorvEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512RorvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPRORVQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskRoundscalePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskRoundscalePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzRoundscalePd(k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskzRoundscalePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func roundscalePd(a [2]float64, imm8 int) [2]float64
TEXT ·roundscalePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VRNDSCALEPD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256MaskRoundscalePd(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·m256MaskRoundscalePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRoundscalePd(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·m256MaskzRoundscalePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RoundscalePd(a [4]float64, imm8 int) [4]float64
TEXT ·m256RoundscalePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VRNDSCALEPD Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRoundscalePd(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·m512MaskRoundscalePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRoundscalePd(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·m512MaskzRoundscalePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RoundscalePd(a [8]float64, imm8 int) [8]float64
TEXT ·m512RoundscalePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VRNDSCALEPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskRoundscalePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskRoundscalePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func maskzRoundscalePs(k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskzRoundscalePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func roundscalePs(a [4]float32, imm8 int) [4]float32
TEXT ·roundscalePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VRNDSCALEPS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func m256MaskRoundscalePs(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·m256MaskRoundscalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzRoundscalePs(k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·m256MaskzRoundscalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256RoundscalePs(a [8]float32, imm8 int) [8]float32
TEXT ·m256RoundscalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VRNDSCALEPS Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRoundscalePs(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·m512MaskRoundscalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzRoundscalePs(k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·m512MaskzRoundscalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512RoundscalePs(a [16]float32, imm8 int) [16]float32
TEXT ·m512RoundscalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRNDSCALEPS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskRoundscaleRoundPd(src [8]float64, k uint8, a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·m512MaskRoundscaleRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzRoundscaleRoundPd(k uint8, a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·m512MaskzRoundscaleRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512RoundscaleRoundPd(a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·m512RoundscaleRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskRoundscaleRoundPs(src [16]float32, k uint16, a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·m512MaskRoundscaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzRoundscaleRoundPs(k uint16, a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·m512MaskzRoundscaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512RoundscaleRoundPs(a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·m512RoundscaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskRoundscaleRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskRoundscaleRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	// TODO: Code missing

	MOVOU X5, ret+68(FP)
	RET

// func maskzRoundscaleRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskzRoundscaleRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func roundscaleRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·roundscaleRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+48(FP)
	RET

// func maskRoundscaleRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskRoundscaleRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	// TODO: Code missing

	MOVOU X5, ret+68(FP)
	RET

// func maskzRoundscaleRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskzRoundscaleRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+52(FP)
	RET

// func roundscaleRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·roundscaleRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+48(FP)
	RET

// func maskRoundscaleSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskRoundscaleSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzRoundscaleSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzRoundscaleSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func roundscaleSd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·roundscaleSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskRoundscaleSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskRoundscaleSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzRoundscaleSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzRoundscaleSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func roundscaleSs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·roundscaleSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskRsqrt14Pd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskRsqrt14Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzRsqrt14Pd(k uint8, a [2]float64) [2]float64
TEXT ·maskzRsqrt14Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VRSQRT14PD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskRsqrt14Pd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskRsqrt14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzRsqrt14Pd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzRsqrt14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VRSQRT14PD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRsqrt14Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskRsqrt14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRsqrt14Pd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzRsqrt14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VRSQRT14PD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rsqrt14Pd(a [8]float64) [8]float64
TEXT ·m512Rsqrt14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VRSQRT14PD Z0

	MOV Z0, ret+0(FP)
	RET

// func maskRsqrt14Ps(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskRsqrt14Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzRsqrt14Ps(k uint8, a [4]float32) [4]float32
TEXT ·maskzRsqrt14Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VRSQRT14PS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskRsqrt14Ps(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskRsqrt14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzRsqrt14Ps(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzRsqrt14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VRSQRT14PS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskRsqrt14Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRsqrt14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzRsqrt14Ps(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzRsqrt14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRSQRT14PS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512Rsqrt14Ps(a [16]float32) [16]float32
TEXT ·m512Rsqrt14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRSQRT14PS Z0

	MOV Z0, ret+0(FP)
	RET

// func maskRsqrt14Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRsqrt14Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRsqrt14Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRsqrt14Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rsqrt14Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rsqrt14Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VRSQRT14SD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskRsqrt14Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRsqrt14Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzRsqrt14Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRsqrt14Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func rsqrt14Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rsqrt14Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VRSQRT14SS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskScalefPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskScalefPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzScalefPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzScalefPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func scalefPd(a [2]float64, b [2]float64) [2]float64
TEXT ·scalefPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VSCALEFPD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskScalefPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskScalefPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzScalefPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzScalefPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256ScalefPd(a [4]float64, b [4]float64) [4]float64
TEXT ·m256ScalefPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VSCALEFPD Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskScalefPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskScalefPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzScalefPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzScalefPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ScalefPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512ScalefPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VSCALEFPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskScalefPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskScalefPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzScalefPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzScalefPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func scalefPs(a [4]float32, b [4]float32) [4]float32
TEXT ·scalefPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VSCALEFPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskScalefPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskScalefPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzScalefPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzScalefPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256ScalefPs(a [8]float32, b [8]float32) [8]float32
TEXT ·m256ScalefPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VSCALEFPS Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskScalefPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskScalefPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzScalefPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzScalefPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ScalefPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512ScalefPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSCALEFPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskScalefRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskScalefRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzScalefRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskzScalefRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ScalefRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512ScalefRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskScalefRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskScalefRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzScalefRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskzScalefRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ScalefRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512ScalefRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskScalefRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskScalefRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzScalefRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzScalefRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func scalefRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·scalefRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskScalefRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskScalefRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzScalefRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzScalefRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func scalefRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·scalefRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskScalefSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskScalefSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzScalefSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzScalefSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func scalefSd(a [2]float64, b [2]float64) [2]float64
TEXT ·scalefSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VSCALEFSD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maskScalefSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskScalefSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzScalefSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzScalefSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func scalefSs(a [4]float32, b [4]float32) [4]float32
TEXT ·scalefSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VSCALEFSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m512SetEpi32(e15 int, e14 int, e13 int, e12 int, e11 int, e10 int, e9 int, e8 int, e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [64]byte
TEXT ·m512SetEpi32(SB),7,$0
	// Unimplemented. Unknown register for type int
	// TODO: Code missing

	MOV , ret+0(FP)
	RET

// func m512SetEpi64(e7 int64, e6 int64, e5 int64, e4 int64, e3 int64, e2 int64, e1 int64, e0 int64) [64]byte
TEXT ·m512SetEpi64(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	// TODO: Code missing

	MOV Z7, ret+64(FP)
	RET

// func m512SetPd(e7 float64, e6 float64, e5 float64, e4 float64, e3 float64, e2 float64, e1 float64, e0 float64) [8]float64
TEXT ·m512SetPd(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	// TODO: Code missing

	MOV Z7, ret+64(FP)
	RET

// func m512SetPs(e15 float32, e14 float32, e13 float32, e12 float32, e11 float32, e10 float32, e9 float32, e8 float32, e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [16]float32
TEXT ·m512SetPs(SB),7,$0
	// Unimplemented. Unknown register for type float32
	// TODO: Code missing

	MOV , ret+0(FP)
	RET

// func m512Set1Epi16(a int16) [64]byte
TEXT ·m512Set1Epi16(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func maskSet1Epi32(src [16]byte, k uint8, a int) [16]byte
TEXT ·maskSet1Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ a+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzSet1Epi32(k uint8, a int) [16]byte
TEXT ·maskzSet1Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskSet1Epi32(src [32]byte, k uint8, a int) [32]byte
TEXT ·m256MaskSet1Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzSet1Epi32(k uint8, a int) [32]byte
TEXT ·m256MaskzSet1Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskSet1Epi32(src [64]byte, k uint16, a int) [64]byte
TEXT ·m512MaskSet1Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzSet1Epi32(k uint16, a int) [64]byte
TEXT ·m512MaskzSet1Epi32(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ a+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8, R9

	MOV Z1, ret+12(FP)
	RET

// func m512Set1Epi32(a int) [64]byte
TEXT ·m512Set1Epi32(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VPBROADCASTD R8

	MOV Z0, ret+8(FP)
	RET

// func maskSet1Epi64(src [16]byte, k uint8, a int64) [16]byte
TEXT ·maskSet1Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ a+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+28(FP)
	RET

// func maskzSet1Epi64(k uint8, a int64) [16]byte
TEXT ·maskzSet1Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8, R9

	MOVOU X1, ret+12(FP)
	RET

// func m256MaskSet1Epi64(src [32]byte, k uint8, a int64) [32]byte
TEXT ·m256MaskSet1Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzSet1Epi64(k uint8, a int64) [32]byte
TEXT ·m256MaskzSet1Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8, R9

	MOV Y1, ret+12(FP)
	RET

// func m512MaskSet1Epi64(src [64]byte, k uint8, a int64) [64]byte
TEXT ·m512MaskSet1Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzSet1Epi64(k uint8, a int64) [64]byte
TEXT ·m512MaskzSet1Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8, R9

	MOV Z1, ret+12(FP)
	RET

// func m512Set1Epi64(a int64) [64]byte
TEXT ·m512Set1Epi64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VPBROADCASTQ R8

	MOV Z0, ret+8(FP)
	RET

// func m512Set1Epi8(a byte) [64]byte
TEXT ·m512Set1Epi8(SB),7,$0
	MOVB a+0(FP),R8

	// TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func m512Set1Pd(a float64) [8]float64
TEXT ·m512Set1Pd(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing

	MOV Z0, ret+8(FP)
	RET

// func m512Set1Ps(a float32) [16]float32
TEXT ·m512Set1Ps(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func m512Set4Epi32(d int, c int, b int, a int) [64]byte
TEXT ·m512Set4Epi32(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	// TODO: Code missing

	MOV Z3, ret+32(FP)
	RET

// func m512Set4Epi64(d int64, c int64, b int64, a int64) [64]byte
TEXT ·m512Set4Epi64(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	// TODO: Code missing

	MOV Z3, ret+32(FP)
	RET

// func m512Set4Pd(d float64, c float64, b float64, a float64) [8]float64
TEXT ·m512Set4Pd(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	// TODO: Code missing

	MOV Z3, ret+32(FP)
	RET

// func m512Set4Ps(d float32, c float32, b float32, a float32) [16]float32
TEXT ·m512Set4Ps(SB),7,$0
	MOVL d+0(FP),R8
	MOVL c+4(FP),R9
	MOVL b+8(FP),R10
	MOVL a+12(FP),R11

	// TODO: Code missing

	MOV Z3, ret+16(FP)
	RET

// func m512SetrEpi32(e15 int, e14 int, e13 int, e12 int, e11 int, e10 int, e9 int, e8 int, e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [64]byte
TEXT ·m512SetrEpi32(SB),7,$0
	// Unimplemented. Unknown register for type int
	// TODO: Code missing

	MOV , ret+0(FP)
	RET

// func m512SetrEpi64(e7 int64, e6 int64, e5 int64, e4 int64, e3 int64, e2 int64, e1 int64, e0 int64) [64]byte
TEXT ·m512SetrEpi64(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	// TODO: Code missing

	MOV Z7, ret+64(FP)
	RET

// func m512SetrPd(e7 float64, e6 float64, e5 float64, e4 float64, e3 float64, e2 float64, e1 float64, e0 float64) [8]float64
TEXT ·m512SetrPd(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	// TODO: Code missing

	MOV Z7, ret+64(FP)
	RET

// func m512SetrPs(e15 float32, e14 float32, e13 float32, e12 float32, e11 float32, e10 float32, e9 float32, e8 float32, e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [16]float32
TEXT ·m512SetrPs(SB),7,$0
	// Unimplemented. Unknown register for type float32
	// TODO: Code missing

	MOV , ret+0(FP)
	RET

// func m512Setr4Epi32(d int, c int, b int, a int) [64]byte
TEXT ·m512Setr4Epi32(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	// TODO: Code missing

	MOV Z3, ret+32(FP)
	RET

// func m512Setr4Epi64(d int64, c int64, b int64, a int64) [64]byte
TEXT ·m512Setr4Epi64(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	// TODO: Code missing

	MOV Z3, ret+32(FP)
	RET

// func m512Setr4Pd(d float64, c float64, b float64, a float64) [8]float64
TEXT ·m512Setr4Pd(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	// TODO: Code missing

	MOV Z3, ret+32(FP)
	RET

// func m512Setr4Ps(d float32, c float32, b float32, a float32) [16]float32
TEXT ·m512Setr4Ps(SB),7,$0
	MOVL d+0(FP),R8
	MOVL c+4(FP),R9
	MOVL b+8(FP),R10
	MOVL a+12(FP),R11

	// TODO: Code missing

	MOV Z3, ret+16(FP)
	RET

// func m512Setzero() [16]float32
TEXT ·m512Setzero(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512SetzeroEpi32() [64]byte
TEXT ·m512SetzeroEpi32(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512SetzeroPd() [8]float64
TEXT ·m512SetzeroPd(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512SetzeroPs() [16]float32
TEXT ·m512SetzeroPs(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512SetzeroSi512() [64]byte
TEXT ·m512SetzeroSi512(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func maskShuffleEpi32(src [16]byte, k uint8, a [16]byte, imm8 MMPERMENUM) [16]byte
TEXT ·maskShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown size of type MMPERMENUM
	// TODO: Code missing

	MOVOU X3, ret+0(FP)
	RET

// func maskzShuffleEpi32(k uint8, a [16]byte, imm8 MMPERMENUM) [16]byte
TEXT ·maskzShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown size of type MMPERMENUM
	// TODO: Code missing

	MOVOU X2, ret+0(FP)
	RET

// func m256MaskShuffleEpi32(src [32]byte, k uint8, a [32]byte, imm8 MMPERMENUM) [32]byte
TEXT ·m256MaskShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzShuffleEpi32(k uint8, a [32]byte, imm8 MMPERMENUM) [32]byte
TEXT ·m256MaskzShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzShuffleEpi32(k uint16, a [64]byte, imm8 MMPERMENUM) [64]byte
TEXT ·m512MaskzShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m256MaskShuffleF32x4(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·m256MaskShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzShuffleF32x4(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·m256MaskzShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256ShuffleF32x4(a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·m256ShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskShuffleF32x4(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·m512MaskShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzShuffleF32x4(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·m512MaskzShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShuffleF32x4(a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·m512ShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m256MaskShuffleF64x2(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·m256MaskShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzShuffleF64x2(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·m256MaskzShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256ShuffleF64x2(a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·m256ShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskShuffleF64x2(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·m512MaskShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzShuffleF64x2(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·m512MaskzShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShuffleF64x2(a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·m512ShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m256MaskShuffleI32x4(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzShuffleI32x4(k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256ShuffleI32x4(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256ShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskShuffleI32x4(src [64]byte, k uint16, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzShuffleI32x4(k uint16, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShuffleI32x4(a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512ShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m256MaskShuffleI64x2(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzShuffleI64x2(k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256ShuffleI64x2(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256ShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskShuffleI64x2(src [64]byte, k uint8, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzShuffleI64x2(k uint8, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShuffleI64x2(a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512ShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskShufflePd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskShufflePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzShufflePd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzShufflePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func m256MaskShufflePd(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·m256MaskShufflePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzShufflePd(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·m256MaskzShufflePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskShufflePd(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·m512MaskShufflePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzShufflePd(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·m512MaskzShufflePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShufflePd(a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·m512ShufflePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskShufflePs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskShufflePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzShufflePs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzShufflePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func m256MaskShufflePs(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·m256MaskShufflePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzShufflePs(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·m256MaskzShufflePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskShufflePs(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·m512MaskShufflePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzShufflePs(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·m512MaskzShufflePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShufflePs(a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·m512ShufflePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSinPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskSinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SinPd(a [8]float64) [8]float64
TEXT ·m512SinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSinPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskSinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SinPs(a [16]float32) [16]float32
TEXT ·m512SinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSincosPd(cos_res [8]float64, sin_src [8]float64, cos_src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskSincosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SincosPd(cos_res [8]float64, a [8]float64) [8]float64
TEXT ·m512SincosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSincosPs(cos_res [16]float32, sin_src [16]float32, cos_src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskSincosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SincosPs(cos_res [16]float32, a [16]float32) [16]float32
TEXT ·m512SincosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSindPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskSindPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SindPd(a [8]float64) [8]float64
TEXT ·m512SindPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSindPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskSindPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SindPs(a [16]float32) [16]float32
TEXT ·m512SindPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSinhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskSinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SinhPd(a [8]float64) [8]float64
TEXT ·m512SinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSinhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskSinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SinhPs(a [16]float32) [16]float32
TEXT ·m512SinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSllEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSllEpi32(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSllEpi32(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSllEpi32(src [64]byte, k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSllEpi32(k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SllEpi32(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSLLD Z0, X1

	MOV Z1, ret+0(FP)
	RET

// func maskSllEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSllEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSllEpi64(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSllEpi64(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSllEpi64(src [64]byte, k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSllEpi64(k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SllEpi64(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSLLQ Z0, X1

	MOV Z1, ret+0(FP)
	RET

// func maskSlliEpi32(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSlliEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+40(FP)
	RET

// func maskzSlliEpi32(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSlliEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSlliEpi32(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSlliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSlliEpi32(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSlliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSlliEpi32(k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSlliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSlliEpi64(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSlliEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+40(FP)
	RET

// func maskzSlliEpi64(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSlliEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSlliEpi64(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSlliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSlliEpi64(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSlliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSlliEpi64(src [64]byte, k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSlliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSlliEpi64(k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSlliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SlliEpi64(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SlliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSLLQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskSllvEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSllvEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSllvEpi32(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSllvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSllvEpi32(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSllvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSllvEpi32(k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSllvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSllvEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSllvEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSllvEpi64(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSllvEpi64(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSllvEpi64(src [64]byte, k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSllvEpi64(k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SllvEpi64(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSLLVQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskSqrtPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskSqrtPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzSqrtPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzSqrtPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VSQRTPD R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskSqrtPd(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·m256MaskSqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzSqrtPd(k uint8, a [4]float64) [4]float64
TEXT ·m256MaskzSqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing
	// Could be:
	// VSQRTPD R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskSqrtPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskSqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzSqrtPd(k uint8, a [8]float64) [8]float64
TEXT ·m512MaskzSqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VSQRTPD R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512SqrtPd(a [8]float64) [8]float64
TEXT ·m512SqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VSQRTPD Z0

	MOV Z0, ret+0(FP)
	RET

// func maskSqrtPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskSqrtPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskzSqrtPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzSqrtPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing
	// Could be:
	// VSQRTPS R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskSqrtPs(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·m256MaskSqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256MaskzSqrtPs(k uint8, a [8]float32) [8]float32
TEXT ·m256MaskzSqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing
	// Could be:
	// VSQRTPS R8, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskSqrtPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskSqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzSqrtPs(k uint16, a [16]float32) [16]float32
TEXT ·m512MaskzSqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSQRTPS R8, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512SqrtPs(a [16]float32) [16]float32
TEXT ·m512SqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSQRTPS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSqrtRoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskSqrtRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSqrtRoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskzSqrtRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SqrtRoundPd(a [8]float64, rounding int) [8]float64
TEXT ·m512SqrtRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VSQRTPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSqrtRoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskSqrtRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSqrtRoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskzSqrtRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SqrtRoundPs(a [16]float32, rounding int) [16]float32
TEXT ·m512SqrtRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSQRTPS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskSqrtRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskSqrtRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzSqrtRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzSqrtRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func sqrtRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·sqrtRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskSqrtRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskSqrtRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzSqrtRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzSqrtRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func sqrtRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·sqrtRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskSqrtSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskSqrtSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSqrtSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzSqrtSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskSqrtSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskSqrtSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSqrtSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzSqrtSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskSraEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSraEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSraEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSraEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSraEpi32(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSraEpi32(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSraEpi32(src [64]byte, k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSraEpi32(k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SraEpi32(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRAD Z0, X1

	MOV Z1, ret+0(FP)
	RET

// func maskSraEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSraEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSraEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSraEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func sraEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sraEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPSRAQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskSraEpi64(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSraEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSraEpi64(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSraEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256SraEpi64(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SraEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPSRAQ Y0, X1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskSraEpi64(src [64]byte, k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSraEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSraEpi64(k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSraEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SraEpi64(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SraEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRAQ Z0, X1

	MOV Z1, ret+0(FP)
	RET

// func maskSraiEpi32(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSraiEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+40(FP)
	RET

// func maskzSraiEpi32(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSraiEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSraiEpi32(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSraiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSraiEpi32(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSraiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSraiEpi32(k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSraiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSraiEpi64(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSraiEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+40(FP)
	RET

// func maskzSraiEpi64(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSraiEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func sraiEpi64(a [16]byte, imm8 uint32) [16]byte
TEXT ·sraiEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL imm8+16(FP),R9

	// TODO: Code missing
	// Could be:
	// VPSRAQ X0, R9

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskSraiEpi64(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSraiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSraiEpi64(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSraiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256SraiEpi64(a [32]byte, imm8 uint32) [32]byte
TEXT ·m256SraiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPSRAQ Y0, R9

	MOV Y1, ret+0(FP)
	RET

// func m512MaskSraiEpi64(src [64]byte, k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSraiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSraiEpi64(k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSraiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SraiEpi64(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SraiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRAQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskSravEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSravEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSravEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSravEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSravEpi32(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSravEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSravEpi32(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSravEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSravEpi32(k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSravEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSravEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSravEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSravEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSravEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func sravEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sravEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPSRAVQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskSravEpi64(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSravEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSravEpi64(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSravEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m256SravEpi64(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SravEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPSRAVQ Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func m512MaskSravEpi64(src [64]byte, k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSravEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSravEpi64(k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSravEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SravEpi64(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SravEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRAVQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskSrlEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSrlEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSrlEpi32(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSrlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSrlEpi32(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSrlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSrlEpi32(src [64]byte, k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSrlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSrlEpi32(k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSrlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SrlEpi32(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SrlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRLD Z0, X1

	MOV Z1, ret+0(FP)
	RET

// func maskSrlEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSrlEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSrlEpi64(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSrlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSrlEpi64(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSrlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSrlEpi64(src [64]byte, k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSrlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSrlEpi64(k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSrlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SrlEpi64(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SrlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRLQ Z0, X1

	MOV Z1, ret+0(FP)
	RET

// func maskSrliEpi32(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSrliEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+40(FP)
	RET

// func maskzSrliEpi32(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSrliEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSrliEpi32(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSrliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSrliEpi32(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSrliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSrliEpi32(k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSrliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSrliEpi64(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSrliEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+40(FP)
	RET

// func maskzSrliEpi64(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSrliEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSrliEpi64(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSrliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSrliEpi64(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSrliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSrliEpi64(src [64]byte, k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSrliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSrliEpi64(k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSrliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SrliEpi64(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SrliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRLQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func maskSrlvEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSrlvEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSrlvEpi32(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSrlvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSrlvEpi32(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSrlvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSrlvEpi32(k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSrlvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSrlvEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSrlvEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSrlvEpi64(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSrlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSrlvEpi64(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSrlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSrlvEpi64(src [64]byte, k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSrlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSrlvEpi64(k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSrlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SrlvEpi64(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SrlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRLVQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskStoreEpi32(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStoreEpi32(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskStoreEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func maskStoreEpi64(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStoreEpi64(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskStoreEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func maskStorePd(mem_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskStorePd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStorePd(mem_addr uintptr, k uint8, a [4]float64) 
TEXT ·m256MaskStorePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	RET

// func maskStorePs(mem_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskStorePs(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStorePs(mem_addr uintptr, k uint8, a [8]float32) 
TEXT ·m256MaskStorePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	RET

// func maskStoreSd(mem_addr float64, k uint8, a [2]float64) 
TEXT ·maskStoreSd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func maskStoreSs(mem_addr float32, k uint8, a [4]float32) 
TEXT ·maskStoreSs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVB k+4(FP),R9
	MOVOU a+8(FP),X2

	// TODO: Code missing

	RET

// func maskStoreuEpi32(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStoreuEpi32(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskStoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskStoreuEpi32(mem_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskStoreuEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskStoreuEpi64(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStoreuEpi64(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·m256MaskStoreuEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	RET

// func m512MaskStoreuEpi64(mem_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskStoreuEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func maskStoreuPd(mem_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskStoreuPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStoreuPd(mem_addr uintptr, k uint8, a [4]float64) 
TEXT ·m256MaskStoreuPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	RET

// func m512MaskStoreuPd(mem_addr uintptr, k uint8, a [8]float64) 
TEXT ·m512MaskStoreuPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512StoreuPd(mem_addr uintptr, a [8]float64) 
TEXT ·m512StoreuPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVUPD R8, Z1

	RET

// func maskStoreuPs(mem_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskStoreuPs(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	// TODO: Code missing

	RET

// func m256MaskStoreuPs(mem_addr uintptr, k uint8, a [8]float32) 
TEXT ·m256MaskStoreuPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	RET

// func m512MaskStoreuPs(mem_addr uintptr, k uint16, a [16]float32) 
TEXT ·m512MaskStoreuPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512StoreuPs(mem_addr uintptr, a [16]float32) 
TEXT ·m512StoreuPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVUPS R8, Z1

	RET

// func m512StoreuSi512(mem_addr uintptr, a [64]byte) 
TEXT ·m512StoreuSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVDQU32 R8, Z1

	RET

// func m512StreamLoadSi512(mem_addr uintptr) [64]byte
TEXT ·m512StreamLoadSi512(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVNTDQA R8

	MOV Z0, ret+8(FP)
	RET

// func m512StreamPd(mem_addr uintptr, a [8]float64) 
TEXT ·m512StreamPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVNTPD R8, Z1

	RET

// func m512StreamPs(mem_addr uintptr, a [16]float32) 
TEXT ·m512StreamPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVNTPS R8, Z1

	RET

// func m512StreamSi512(mem_addr uintptr, a [64]byte) 
TEXT ·m512StreamSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVNTDQA R8, Z1

	RET

// func maskSubEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSubEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSubEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSubEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSubEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskSubEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSubEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SubEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSUBQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskSubPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskSubPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzSubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskSubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSubPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzSubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSubPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzSubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskSubPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskSubPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzSubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskSubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzSubPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzSubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzSubPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzSubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskzSubRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskzSubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzSubRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskzSubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskSubRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskSubRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzSubRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzSubRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func subRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·subRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskSubRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskSubRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzSubRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzSubRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+44(FP)
	RET

// func subRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·subRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing

	MOVOU X2, ret+40(FP)
	RET

// func maskSubSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskSubSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzSubSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func maskSubSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskSubSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzSubSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m512MaskSvmlRoundPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskSvmlRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512SvmlRoundPd(a [8]float64) [8]float64
TEXT ·m512SvmlRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTanPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskTanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TanPd(a [8]float64) [8]float64
TEXT ·m512TanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTanPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskTanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TanPs(a [16]float32) [16]float32
TEXT ·m512TanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTandPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskTandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TandPd(a [8]float64) [8]float64
TEXT ·m512TandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTandPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskTandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TandPs(a [16]float32) [16]float32
TEXT ·m512TandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTanhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskTanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TanhPd(a [8]float64) [8]float64
TEXT ·m512TanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTanhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskTanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TanhPs(a [16]float32) [16]float32
TEXT ·m512TanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTernarylogicEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·maskTernarylogicEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzTernarylogicEpi32(k uint8, a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·maskzTernarylogicEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func ternarylogicEpi32(a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·ternarylogicEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+56(FP)
	RET

// func m256MaskTernarylogicEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskTernarylogicEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzTernarylogicEpi32(k uint8, a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzTernarylogicEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256TernarylogicEpi32(a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·m256TernarylogicEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskTernarylogicEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskTernarylogicEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzTernarylogicEpi32(k uint16, a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzTernarylogicEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512TernarylogicEpi32(a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·m512TernarylogicEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskTernarylogicEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·maskTernarylogicEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func maskzTernarylogicEpi64(k uint8, a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·maskzTernarylogicEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	// TODO: Code missing

	MOVOU X4, ret+60(FP)
	RET

// func ternarylogicEpi64(a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·ternarylogicEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+56(FP)
	RET

// func m256MaskTernarylogicEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskTernarylogicEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzTernarylogicEpi64(k uint8, a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzTernarylogicEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y4, ret+0(FP)
	RET

// func m256TernarylogicEpi64(a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·m256TernarylogicEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m512MaskTernarylogicEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskTernarylogicEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskzTernarylogicEpi64(k uint8, a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzTernarylogicEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512TernarylogicEpi64(a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·m512TernarylogicEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func maskTestEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·testEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPTESTMD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskTestEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256TestEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256TestEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPTESTMD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func maskTestEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·testEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPTESTMQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskTestEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256TestEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256TestEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPTESTMQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskTestEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskTestEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512TestEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512TestEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPTESTMQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func maskTestnEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestnEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testnEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·testnEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPTESTNMD X0, X1

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestnEpi32Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskTestnEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256TestnEpi32Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256TestnEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPTESTNMD Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskTestnEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskTestnEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512TestnEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512TestnEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPTESTNMD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func maskTestnEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestnEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testnEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·testnEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing
	// Could be:
	// VPTESTNMQ X0, X1

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestnEpi64Mask(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·m256MaskTestnEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m256TestnEpi64Mask(a [32]byte, b [32]byte) uint8
TEXT ·m256TestnEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing
	// Could be:
	// VPTESTNMQ Y0, Y1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskTestnEpi64Mask(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·m512MaskTestnEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512TestnEpi64Mask(a [64]byte, b [64]byte) uint8
TEXT ·m512TestnEpi64Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPTESTNMQ Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskTruncPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskTruncPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TruncPd(a [8]float64) [8]float64
TEXT ·m512TruncPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskTruncPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskTruncPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512TruncPs(a [16]float32) [16]float32
TEXT ·m512TruncPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Undefined() [16]float32
TEXT ·m512Undefined(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512UndefinedEpi32() [64]byte
TEXT ·m512UndefinedEpi32(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512UndefinedPd() [8]float64
TEXT ·m512UndefinedPd(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func m512UndefinedPs() [16]float32
TEXT ·m512UndefinedPs(SB),7,$0

	// TODO: Code missing

	MOV Z-1, ret+0(FP)
	RET

// func maskUnpackhiEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackhiEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackhiEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackhiEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackhiEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackhiEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackhiEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPUNPCKHDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackhiEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackhiEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackhiEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackhiEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackhiEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackhiEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackhiEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPUNPCKHQDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackhiPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskUnpackhiPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackhiPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzUnpackhiPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackhiPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskUnpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackhiPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzUnpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackhiPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskUnpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackhiPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzUnpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackhiPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512UnpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VUNPCKHPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackhiPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskUnpackhiPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackhiPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzUnpackhiPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackhiPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskUnpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackhiPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzUnpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackhiPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskUnpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackhiPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzUnpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackhiPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512UnpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VUNPCKHPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackloEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackloEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackloEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackloEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackloEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackloEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackloEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPUNPCKLDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackloEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackloEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackloEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackloEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackloEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackloEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackloEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPUNPCKLQDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackloPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskUnpackloPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackloPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzUnpackloPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackloPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskUnpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackloPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzUnpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackloPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskUnpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackloPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskzUnpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackloPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512UnpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VUNPCKLPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskUnpackloPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskUnpackloPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackloPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzUnpackloPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackloPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskUnpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzUnpackloPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzUnpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskUnpackloPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskUnpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzUnpackloPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskzUnpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512UnpackloPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512UnpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VUNPCKLPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func maskXorEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskXorEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzXorEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzXorEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskXorEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskXorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzXorEpi32(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzXorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzXorEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzXorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func maskXorEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskXorEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing

	MOVOU X3, ret+52(FP)
	RET

// func maskzXorEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzXorEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskXorEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskXorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzXorEpi64(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzXorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	// TODO: Code missing

	MOV Y2, ret+0(FP)
	RET

// func m512MaskzXorEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzXorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

