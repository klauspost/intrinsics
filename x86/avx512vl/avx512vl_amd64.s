// func maskAddPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAddPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VADDPD

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VADDPD

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskAddPd(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256d

	// TODO: Code missing - uses instrunction: VADDPD

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAddPd(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·m256MaskzAddPd(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256d

	// TODO: Code missing - uses instrunction: VADDPD

	MOV Y2, ret+0(FP)
	RET

// func maskAddPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAddPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VADDPS

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VADDPS

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskAddPs(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256

	// TODO: Code missing - uses instrunction: VADDPS

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzAddPs(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·m256MaskzAddPs(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256

	// TODO: Code missing - uses instrunction: VADDPS

	MOV Y2, ret+0(FP)
	RET

// func alignrEpi32(a [16]byte, b [16]byte, count int) [16]byte
TEXT ·alignrEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	// TODO: Code missing - uses instrunction: VALIGND

	MOVOU X2, ret+40(FP)
	RET

// func maskAlignrEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskAlignrEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	// TODO: Code missing - uses instrunction: VALIGND

	MOVOU X4, ret+60(FP)
	RET

// func maskzAlignrEpi32(k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskzAlignrEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	// TODO: Code missing - uses instrunction: VALIGND

	MOVOU X3, ret+44(FP)
	RET

// func m256AlignrEpi32(a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256AlignrEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VALIGND

	MOV Y2, ret+0(FP)
	RET

// func m256MaskAlignrEpi32(src [32]byte, k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskAlignrEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VALIGND

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzAlignrEpi32(k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskzAlignrEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VALIGND

	MOV Y3, ret+0(FP)
	RET

// func alignrEpi64(a [16]byte, b [16]byte, count int) [16]byte
TEXT ·alignrEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	// TODO: Code missing - uses instrunction: VALIGNQ

	MOVOU X2, ret+40(FP)
	RET

// func maskAlignrEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskAlignrEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	// TODO: Code missing - uses instrunction: VALIGNQ

	MOVOU X4, ret+60(FP)
	RET

// func maskzAlignrEpi64(k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskzAlignrEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	// TODO: Code missing - uses instrunction: VALIGNQ

	MOVOU X3, ret+44(FP)
	RET

// func m256AlignrEpi64(a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256AlignrEpi64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VALIGNQ

	MOV Y2, ret+0(FP)
	RET

// func m256MaskAlignrEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskAlignrEpi64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VALIGNQ

	MOV Y4, ret+0(FP)
	RET

// func m256MaskzAlignrEpi64(k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskzAlignrEpi64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VALIGNQ

	MOV Y3, ret+0(FP)
	RET

// func madd52hiEpu64(a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·madd52hiEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOVOU X2, ret+48(FP)
	RET

// func maskMadd52hiEpu64(a [16]byte, k uint8, b [16]byte, c [16]byte) [16]byte
TEXT ·maskMadd52hiEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOVOU X3, ret+52(FP)
	RET

// func maskzMadd52hiEpu64(k uint8, a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·maskzMadd52hiEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOVOU X3, ret+52(FP)
	RET

// func m256Madd52hiEpu64(a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·m256Madd52hiEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOV Y2, ret+0(FP)
	RET

// func m256MaskMadd52hiEpu64(a [32]byte, k uint8, b [32]byte, c [32]byte) [32]byte
TEXT ·m256MaskMadd52hiEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMadd52hiEpu64(k uint8, a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·m256MaskzMadd52hiEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOV Y3, ret+0(FP)
	RET

// func madd52loEpu64(a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·madd52loEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOVOU X2, ret+48(FP)
	RET

// func maskMadd52loEpu64(a [16]byte, k uint8, b [16]byte, c [16]byte) [16]byte
TEXT ·maskMadd52loEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOVOU X3, ret+52(FP)
	RET

// func maskzMadd52loEpu64(k uint8, a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·maskzMadd52loEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOVOU X3, ret+52(FP)
	RET

// func m256Madd52loEpu64(a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·m256Madd52loEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOV Y2, ret+0(FP)
	RET

// func m256MaskMadd52loEpu64(a [32]byte, k uint8, b [32]byte, c [32]byte) [32]byte
TEXT ·m256MaskMadd52loEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMadd52loEpu64(k uint8, a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·m256MaskzMadd52loEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOV Y3, ret+0(FP)
	RET

// func movmEpi8(k uint16) [16]byte
TEXT ·movmEpi8(SB),7,$0
	MOVW k+0(FP),R8

	// TODO: Code missing - could be:
	// VPMOVM2B R8

	MOVOU X0, ret+4(FP)
	RET

// func maskMultishiftEpi64Epi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMultishiftEpi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMULTISHIFTQB

	MOVOU X3, ret+52(FP)
	RET

// func maskzMultishiftEpi64Epi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMultishiftEpi64Epi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMULTISHIFTQB

	MOVOU X2, ret+36(FP)
	RET

// func multishiftEpi64Epi8(a [16]byte, b [16]byte) [16]byte
TEXT ·multishiftEpi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPMULTISHIFTQB X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskMultishiftEpi64Epi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMultishiftEpi64Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMULTISHIFTQB

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzMultishiftEpi64Epi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMultishiftEpi64Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPMULTISHIFTQB

	MOV Y2, ret+0(FP)
	RET

// func m256MultishiftEpi64Epi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MultishiftEpi64Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - could be:
	// VPMULTISHIFTQB Y0, Y1

	MOV Y1, ret+0(FP)
	RET

// func maskPermutex2varEpi8(a [16]byte, k uint16, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMT2B

	MOVOU X3, ret+52(FP)
	RET

// func mask2Permutex2varEpi8(a [16]byte, idx [16]byte, k uint16, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVW k+32(FP),R10
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMI2B

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutex2varEpi8(k uint16, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMI2B, VPERMT2B

	MOVOU X3, ret+52(FP)
	RET

// func permutex2varEpi8(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing - uses instrunction: VPERMI2B

	MOVOU X2, ret+48(FP)
	RET

// func m256MaskPermutex2varEpi8(a [32]byte, k uint32, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPermutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPERMT2B

	MOV Y3, ret+0(FP)
	RET

// func m256Mask2Permutex2varEpi8(a [32]byte, idx [32]byte, k uint32, b [32]byte) [32]byte
TEXT ·m256Mask2Permutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPERMI2B

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutex2varEpi8(k uint32, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPermutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPERMI2B, VPERMT2B

	MOV Y3, ret+0(FP)
	RET

// func m256Permutex2varEpi8(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256Permutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPERMI2B

	MOV Y2, ret+0(FP)
	RET

// func maskPermutexvarEpi8(src [16]byte, k uint16, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskPermutexvarEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU a+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMB

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutexvarEpi8(k uint16, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskzPermutexvarEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU idx+4(FP),X1
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPERMB

	MOVOU X2, ret+36(FP)
	RET

// func permutexvarEpi8(idx [16]byte, a [16]byte) [16]byte
TEXT ·permutexvarEpi8(SB),7,$0
	MOVOU idx+0(FP),X0
	MOVOU a+16(FP),X1

	// TODO: Code missing - could be:
	// VPERMB X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskPermutexvarEpi8(src [32]byte, k uint32, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskPermutexvarEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPERMB

	MOV Y3, ret+0(FP)
	RET

// func m256MaskzPermutexvarEpi8(k uint32, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskzPermutexvarEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - uses instrunction: VPERMB

	MOV Y2, ret+0(FP)
	RET

// func m256PermutexvarEpi8(idx [32]byte, a [32]byte) [32]byte
TEXT ·m256PermutexvarEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i

	// TODO: Code missing - could be:
	// VPERMB Y0, Y1

	MOV Y1, ret+0(FP)
	RET

