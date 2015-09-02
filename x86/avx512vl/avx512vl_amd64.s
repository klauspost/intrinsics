// func maskAddPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAddPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAddPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskAddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzAddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskAddPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAddPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAddPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskAddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzAddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func alignrEpi32(a [16]byte, b [16]byte, count int) [16]byte
TEXT ·alignrEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAlignrEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskAlignrEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAlignrEpi32(k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskzAlignrEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func alignrEpi321(a [32]byte, b [32]byte, count int) [32]byte
TEXT ·alignrEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskAlignrEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·maskAlignrEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAlignrEpi321(k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·maskzAlignrEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func alignrEpi64(a [16]byte, b [16]byte, count int) [16]byte
TEXT ·alignrEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAlignrEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskAlignrEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAlignrEpi64(k uint8, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskzAlignrEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func alignrEpi641(a [32]byte, b [32]byte, count int) [32]byte
TEXT ·alignrEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskAlignrEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·maskAlignrEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAlignrEpi641(k uint8, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·maskzAlignrEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func madd52hiEpu64(a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·madd52hiEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskMadd52hiEpu64(a [16]byte, k uint8, b [16]byte, c [16]byte) [16]byte
TEXT ·maskMadd52hiEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMadd52hiEpu64(k uint8, a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·maskzMadd52hiEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func madd52hiEpu641(a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·madd52hiEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMadd52hiEpu641(a [32]byte, k uint8, b [32]byte, c [32]byte) [32]byte
TEXT ·maskMadd52hiEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMadd52hiEpu641(k uint8, a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·maskzMadd52hiEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func madd52loEpu64(a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·madd52loEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskMadd52loEpu64(a [16]byte, k uint8, b [16]byte, c [16]byte) [16]byte
TEXT ·maskMadd52loEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMadd52loEpu64(k uint8, a [16]byte, b [16]byte, c [16]byte) [16]byte
TEXT ·maskzMadd52loEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func madd52loEpu641(a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·madd52loEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMadd52loEpu641(a [32]byte, k uint8, b [32]byte, c [32]byte) [32]byte
TEXT ·maskMadd52loEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMadd52loEpu641(k uint8, a [32]byte, b [32]byte, c [32]byte) [32]byte
TEXT ·maskzMadd52loEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func movmEpi8(k uint16) [16]byte
TEXT ·movmEpi8(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func maskMultishiftEpi64Epi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMultishiftEpi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMultishiftEpi64Epi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMultishiftEpi64Epi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func multishiftEpi64Epi8(a [16]byte, b [16]byte) [16]byte
TEXT ·multishiftEpi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMultishiftEpi64Epi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMultishiftEpi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMultishiftEpi64Epi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMultishiftEpi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func multishiftEpi64Epi81(a [32]byte, b [32]byte) [32]byte
TEXT ·multishiftEpi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutex2varEpi8(a [16]byte, k uint16, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varEpi8(a [16]byte, idx [16]byte, k uint16, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVW k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varEpi8(k uint16, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varEpi8(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varEpi81(a [32]byte, k uint32, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskPermutex2varEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask2Permutex2varEpi81(a [32]byte, idx [32]byte, k uint32, b [32]byte) [32]byte
TEXT ·mask2Permutex2varEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutex2varEpi81(k uint32, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskzPermutex2varEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutex2varEpi81(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·permutex2varEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexvarEpi8(src [16]byte, k uint16, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskPermutexvarEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU a+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutexvarEpi8(k uint16, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskzPermutexvarEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU idx+4(FP),X1
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func permutexvarEpi8(idx [16]byte, a [16]byte) [16]byte
TEXT ·permutexvarEpi8(SB),7,$0
	MOVOU idx+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPermutexvarEpi81(src [32]byte, k uint32, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskPermutexvarEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexvarEpi81(k uint32, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskzPermutexvarEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexvarEpi81(idx [32]byte, a [32]byte) [32]byte
TEXT ·permutexvarEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

