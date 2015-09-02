// func maskAbsEpi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskAbsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbsEpi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzAbsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskAbsEpi16(src [32]byte, k uint16, a [32]byte) [32]byte
TEXT ·m256MaskAbsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAbsEpi16(k uint16, a [32]byte) [32]byte
TEXT ·m256MaskzAbsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AbsEpi16(a [64]byte) [64]byte
TEXT ·m512AbsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAbsEpi16(src [64]byte, k uint32, a [64]byte) [64]byte
TEXT ·m512MaskAbsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAbsEpi16(k uint32, a [64]byte) [64]byte
TEXT ·m512MaskzAbsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAbsEpi8(src [16]byte, k uint16, a [16]byte) [16]byte
TEXT ·maskAbsEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbsEpi8(k uint16, a [16]byte) [16]byte
TEXT ·maskzAbsEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskAbsEpi8(src [32]byte, k uint32, a [32]byte) [32]byte
TEXT ·m256MaskAbsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAbsEpi8(k uint32, a [32]byte) [32]byte
TEXT ·m256MaskzAbsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AbsEpi8(a [64]byte) [64]byte
TEXT ·m512AbsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAbsEpi8(src [64]byte, k uint64, a [64]byte) [64]byte
TEXT ·m512MaskAbsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAbsEpi8(k uint64, a [64]byte) [64]byte
TEXT ·m512MaskzAbsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAddEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAddEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AddEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAddEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAddEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAddEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAddEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AddEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAddEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAddEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAddsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAddsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AddsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAddsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAddsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAddsEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAddsEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AddsEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAddsEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAddsEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAddsEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAddsEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AddsEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAddsEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAddsEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAddsEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAddsEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AddsEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAddsEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAddsEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAlignrEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskAlignrEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAlignrEpi8(k uint16, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskzAlignrEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func m256MaskAlignrEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskAlignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAlignrEpi8(k uint32, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskzAlignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AlignrEpi8(a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512AlignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAlignrEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskAlignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAlignrEpi8(k uint64, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskzAlignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAvgEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAvgEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAvgEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAvgEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAvgEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAvgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAvgEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAvgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AvgEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AvgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAvgEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAvgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAvgEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAvgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAvgEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAvgEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAvgEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAvgEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskAvgEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAvgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzAvgEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAvgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512AvgEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AvgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAvgEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAvgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzAvgEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAvgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBlendEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskBlendEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskBlendEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskBlendEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskBlendEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBlendEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskBlendEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskBlendEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskBlendEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskBlendEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBroadcastbEpi8(src [16]byte, k uint16, a [16]byte) [16]byte
TEXT ·maskBroadcastbEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastbEpi8(k uint16, a [16]byte) [16]byte
TEXT ·maskzBroadcastbEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskBroadcastbEpi8(src [32]byte, k uint32, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastbEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzBroadcastbEpi8(k uint32, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastbEpi8(SB),7,$0
	MOVL k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func m512BroadcastbEpi8(a [16]byte) [64]byte
TEXT ·m512BroadcastbEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastbEpi8(src [64]byte, k uint64, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastbEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzBroadcastbEpi8(k uint64, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastbEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	MOV Z0, ret+24(FP)
	RET

// func maskBroadcastwEpi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastwEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastwEpi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastwEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskBroadcastwEpi16(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastwEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzBroadcastwEpi16(k uint16, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastwEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func m512BroadcastwEpi16(a [16]byte) [64]byte
TEXT ·m512BroadcastwEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastwEpi16(src [64]byte, k uint32, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastwEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzBroadcastwEpi16(k uint32, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastwEpi16(SB),7,$0
	MOVL k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func m512BslliEpi128(a [64]byte, imm8 int) [64]byte
TEXT ·m512BslliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512BsrliEpi128(a [64]byte, imm8 int) [64]byte
TEXT ·m512BsrliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cmpEpi16Mask(a [16]byte, b [16]byte, imm8 int) uint8
TEXT ·cmpEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpEpi16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 int) uint8
TEXT ·maskCmpEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func m256CmpEpi16Mask(a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·m256CmpEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpEpi16Mask(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·m256MaskCmpEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpEpi16Mask(a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·m512CmpEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpEpi16Mask(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·m512MaskCmpEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpEpi8Mask(a [16]byte, b [16]byte, imm8 int) uint16
TEXT ·cmpEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVW $0, ret+40(FP)
	RET

// func maskCmpEpi8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 int) uint16
TEXT ·maskCmpEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVW $0, ret+44(FP)
	RET

// func m256CmpEpi8Mask(a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·m256CmpEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpEpi8Mask(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·m256MaskCmpEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpEpi8Mask(a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·m512CmpEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpEpi8Mask(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·m512MaskCmpEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpEpu16Mask(a [16]byte, b [16]byte, imm8 int) uint8
TEXT ·cmpEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpEpu16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 int) uint8
TEXT ·maskCmpEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func m256CmpEpu16Mask(a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·m256CmpEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpEpu16Mask(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·m256MaskCmpEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpEpu16Mask(a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·m512CmpEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpEpu16Mask(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·m512MaskCmpEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpEpu8Mask(a [16]byte, b [16]byte, imm8 int) uint16
TEXT ·cmpEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVW $0, ret+40(FP)
	RET

// func maskCmpEpu8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 int) uint16
TEXT ·maskCmpEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVW $0, ret+44(FP)
	RET

// func m256CmpEpu8Mask(a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·m256CmpEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpEpu8Mask(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·m256MaskCmpEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpEpu8Mask(a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·m512CmpEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpEpu8Mask(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·m512MaskCmpEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpeqEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpeqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpeqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpeqEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpeqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpeqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpeqEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpeqEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpeqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpeqEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpeqEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpeqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpeqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpeqEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpeqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpeqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpeqEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpeqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpeqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpeqEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpeqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpeqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpeqEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpeqEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpeqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpeqEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpeqEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpeqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpeqEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpeqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpeqEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpeqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpeqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpgeEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgeEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgeEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgeEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgeEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgeEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgeEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgeEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgeEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgeEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgeEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgeEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgeEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpgeEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgeEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgeEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpgeEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgeEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgeEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgeEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgeEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgeEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgeEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgeEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgeEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgeEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgeEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgeEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpgeEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgeEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpgeEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgeEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgeEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpgtEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgtEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgtEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgtEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgtEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgtEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgtEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgtEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgtEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgtEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgtEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgtEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgtEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpgtEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgtEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgtEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpgtEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgtEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgtEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgtEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgtEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgtEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgtEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgtEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgtEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgtEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgtEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgtEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpgtEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgtEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpgtEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgtEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgtEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpleEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpleEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpleEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpleEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpleEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpleEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpleEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpleEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpleEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpleEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpleEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpleEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpleEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpleEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpleEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpleEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpleEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpleEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpleEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpleEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpleEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpleEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpleEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpleEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpleEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpleEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpleEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpleEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpleEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpleEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpleEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpleEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpleEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpleEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpleEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpleEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpleEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpleEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpleEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpleEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpltEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpltEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpltEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpltEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpltEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpltEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpltEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpltEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpltEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpltEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpltEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpltEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpltEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpltEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpltEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpltEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpltEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpltEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpltEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpltEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpltEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpltEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpltEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpltEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpltEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpltEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpltEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpltEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpltEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpltEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpltEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpltEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpltEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpltEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpltEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpltEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpltEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpltEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpltEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpltEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpneqEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpneqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpneqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpneqEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpneqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpneqEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpneqEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpneqEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpneqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpneqEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpneqEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpneqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpneqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpneqEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpneqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpneqEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cmpneqEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpneqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpneqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpneqEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpneqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpneqEpu16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpneqEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpneqEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpneqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpneqEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func m256CmpneqEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpneqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256MaskCmpneqEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpneqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512CmpneqEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpneqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpneqEpu8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func cvtepi16Epi8(a [16]byte) [16]byte
TEXT ·cvtepi16Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi16Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi16Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256Cvtepi16Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskzCvtepi16Epi8(k uint16, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m512Cvtepi16Epi8(a [64]byte) [32]byte
TEXT ·m512Cvtepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·m512MaskCvtepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskzCvtepi16Epi8(k uint32, a [64]byte) [32]byte
TEXT ·m512MaskzCvtepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi16StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func m256MaskCvtepi16StoreuEpi8(base_addr uintptr, k uint16, a [32]byte) 
TEXT ·m256MaskCvtepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func m512MaskCvtepi16StoreuEpi8(base_addr uintptr, k uint32, a [64]byte) 
TEXT ·m512MaskCvtepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	RET

// func maskCvtepi8Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi8Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi8Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi8Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskCvtepi8Epi16(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzCvtepi8Epi16(k uint16, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi8Epi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func m512Cvtepi8Epi16(a [32]byte) [64]byte
TEXT ·m512Cvtepi8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepi8Epi16(src [64]byte, k uint32, a [32]byte) [64]byte
TEXT ·m512MaskCvtepi8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzCvtepi8Epi16(k uint32, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepi8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu8Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu8Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu8Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu8Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskCvtepu8Epi16(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzCvtepu8Epi16(k uint16, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu8Epi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func m512Cvtepu8Epi16(a [32]byte) [64]byte
TEXT ·m512Cvtepu8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepu8Epi16(src [64]byte, k uint32, a [32]byte) [64]byte
TEXT ·m512MaskCvtepu8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzCvtepu8Epi16(k uint32, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepu8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtsepi16Epi8(a [16]byte) [16]byte
TEXT ·cvtsepi16Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi16Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi16Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256Cvtsepi16Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtsepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtsepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskzCvtsepi16Epi8(k uint16, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m512Cvtsepi16Epi8(a [64]byte) [32]byte
TEXT ·m512Cvtsepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtsepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·m512MaskCvtsepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskzCvtsepi16Epi8(k uint32, a [64]byte) [32]byte
TEXT ·m512MaskzCvtsepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtsepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi16StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func m256MaskCvtsepi16StoreuEpi8(base_addr uintptr, k uint16, a [32]byte) 
TEXT ·m256MaskCvtsepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func m512MaskCvtsepi16StoreuEpi8(base_addr uintptr, k uint32, a [64]byte) 
TEXT ·m512MaskCvtsepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	RET

// func cvtusepi16Epi8(a [16]byte) [16]byte
TEXT ·cvtusepi16Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi16Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi16Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256Cvtusepi16Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtusepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskCvtusepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskzCvtusepi16Epi8(k uint16, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m512Cvtusepi16Epi8(a [64]byte) [32]byte
TEXT ·m512Cvtusepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskCvtusepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·m512MaskCvtusepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskzCvtusepi16Epi8(k uint32, a [64]byte) [32]byte
TEXT ·m512MaskzCvtusepi16Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtusepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi16StoreuEpi8(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func m256MaskCvtusepi16StoreuEpi8(base_addr uintptr, k uint16, a [32]byte) 
TEXT ·m256MaskCvtusepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func m512MaskCvtusepi16StoreuEpi8(base_addr uintptr, k uint32, a [64]byte) 
TEXT ·m512MaskCvtusepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	RET

// func dbsadEpu8(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·dbsadEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskDbsadEpu8(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·maskDbsadEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzDbsadEpu8(k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·maskzDbsadEpu8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func m256DbsadEpu8(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256DbsadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskDbsadEpu8(src [32]byte, k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskDbsadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzDbsadEpu8(k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzDbsadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512DbsadEpu8(a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512DbsadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskDbsadEpu8(src [64]byte, k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskDbsadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzDbsadEpu8(k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzDbsadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Kunpackd(a uint64, b uint64) uint64
TEXT ·m512Kunpackd(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func m512Kunpackw(a uint32, b uint32) uint32
TEXT ·m512Kunpackw(SB),7,$0
	MOVL a+0(FP),R8
	MOVL b+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func maskLoaduEpi16(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzLoaduEpi16(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func m256MaskLoaduEpi16(src [32]byte, k uint16, mem_addr uintptr) [32]byte
TEXT ·m256MaskLoaduEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzLoaduEpi16(k uint16, mem_addr uintptr) [32]byte
TEXT ·m256MaskzLoaduEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+12(FP)
	RET

// func m512MaskLoaduEpi16(src [64]byte, k uint32, mem_addr uintptr) [64]byte
TEXT ·m512MaskLoaduEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzLoaduEpi16(k uint32, mem_addr uintptr) [64]byte
TEXT ·m512MaskzLoaduEpi16(SB),7,$0
	MOVL k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOV Z0, ret+12(FP)
	RET

// func maskLoaduEpi8(src [16]byte, k uint16, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzLoaduEpi8(k uint16, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func m256MaskLoaduEpi8(src [32]byte, k uint32, mem_addr uintptr) [32]byte
TEXT ·m256MaskLoaduEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzLoaduEpi8(k uint32, mem_addr uintptr) [32]byte
TEXT ·m256MaskzLoaduEpi8(SB),7,$0
	MOVL k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+12(FP)
	RET

// func m512MaskLoaduEpi8(src [64]byte, k uint64, mem_addr uintptr) [64]byte
TEXT ·m512MaskLoaduEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzLoaduEpi8(k uint64, mem_addr uintptr) [64]byte
TEXT ·m512MaskzLoaduEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	MOVQ mem_addr+8(FP),R9

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskMaddEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaddEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaddEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMaddEpi16(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMaddEpi16(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaddEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskMaddEpi16(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMaddEpi16(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaddubsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaddubsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaddubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaddubsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMaddubsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMaddubsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaddubsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskMaddubsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMaddubsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMaxEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMaxEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMaxEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMaxEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaxEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMaxEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMaxEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMaxEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMaxEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaxEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMaxEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMaxEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMaxEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMaxEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaxEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMaxEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMaxEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMaxEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMaxEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaxEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMinEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMinEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMinEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMinEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MinEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMinEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMinEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMinEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMinEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MinEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMinEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMinEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMinEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMinEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MinEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMinEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMinEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMinEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMinEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MinEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMovEpi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskMovEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovEpi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzMovEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskMovEpi16(src [32]byte, k uint16, a [32]byte) [32]byte
TEXT ·m256MaskMovEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMovEpi16(k uint16, a [32]byte) [32]byte
TEXT ·m256MaskzMovEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMovEpi16(src [64]byte, k uint32, a [64]byte) [64]byte
TEXT ·m512MaskMovEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMovEpi16(k uint32, a [64]byte) [64]byte
TEXT ·m512MaskzMovEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMovEpi8(src [16]byte, k uint16, a [16]byte) [16]byte
TEXT ·maskMovEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovEpi8(k uint16, a [16]byte) [16]byte
TEXT ·maskzMovEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func m256MaskMovEpi8(src [32]byte, k uint32, a [32]byte) [32]byte
TEXT ·m256MaskMovEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMovEpi8(k uint32, a [32]byte) [32]byte
TEXT ·m256MaskzMovEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMovEpi8(src [64]byte, k uint64, a [64]byte) [64]byte
TEXT ·m512MaskMovEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMovEpi8(k uint64, a [64]byte) [64]byte
TEXT ·m512MaskzMovEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func movepi16Mask(a [16]byte) uint8
TEXT ·movepi16Mask(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func m256Movepi16Mask(a [32]byte) uint16
TEXT ·m256Movepi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512Movepi16Mask(a [64]byte) uint32
TEXT ·m512Movepi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func movepi8Mask(a [16]byte) uint16
TEXT ·movepi8Mask(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVW $0, ret+16(FP)
	RET

// func m256Movepi8Mask(a [32]byte) uint32
TEXT ·m256Movepi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512Movepi8Mask(a [64]byte) uint64
TEXT ·m512Movepi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func movmEpi16(k uint8) [16]byte
TEXT ·movmEpi16(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func m256MovmEpi16(k uint16) [32]byte
TEXT ·m256MovmEpi16(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func m512MovmEpi16(k uint32) [64]byte
TEXT ·m512MovmEpi16(SB),7,$0
	MOVL k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func m256MovmEpi8(k uint32) [32]byte
TEXT ·m256MovmEpi8(SB),7,$0
	MOVL k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func m512MovmEpi8(k uint64) [64]byte
TEXT ·m512MovmEpi8(SB),7,$0
	MOVQ k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+8(FP)
	RET

// func maskMulhiEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulhiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulhiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMulhiEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMulhiEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMulhiEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMulhiEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MulhiEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulhiEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulhiEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulhiEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulhiEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMulhiEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMulhiEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMulhiEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMulhiEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MulhiEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulhrsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulhrsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulhrsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulhrsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMulhrsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMulhrsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMulhrsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMulhrsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MulhrsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulloEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulloEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulloEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskMulloEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzMulloEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskMulloEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzMulloEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MulloEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPacksEpi16(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPacksEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPacksEpi16(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPacksEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskPacksEpi16(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPacksEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzPacksEpi16(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPacksEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskPacksEpi16(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPacksEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzPacksEpi16(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPacksEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512PacksEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PacksEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPacksEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPacksEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPacksEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPacksEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskPacksEpi32(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPacksEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzPacksEpi32(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPacksEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskPacksEpi32(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPacksEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzPacksEpi32(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPacksEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512PacksEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PacksEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPackusEpi16(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPackusEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPackusEpi16(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPackusEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskPackusEpi16(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPackusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzPackusEpi16(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPackusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskPackusEpi16(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPackusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzPackusEpi16(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPackusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512PackusEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PackusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPackusEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPackusEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPackusEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPackusEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskPackusEpi32(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPackusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzPackusEpi32(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPackusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskPackusEpi32(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPackusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzPackusEpi32(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPackusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512PackusEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PackusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutex2varEpi16(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varEpi16(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varEpi16(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varEpi16(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func m256MaskPermutex2varEpi16(a [32]byte, k uint16, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPermutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Mask2Permutex2varEpi16(a [32]byte, idx [32]byte, k uint16, b [32]byte) [32]byte
TEXT ·m256Mask2Permutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzPermutex2varEpi16(k uint16, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPermutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Permutex2varEpi16(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256Permutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskPermutex2varEpi16(a [64]byte, k uint32, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPermutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Mask2Permutex2varEpi16(a [64]byte, idx [64]byte, k uint32, b [64]byte) [64]byte
TEXT ·m512Mask2Permutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzPermutex2varEpi16(k uint32, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPermutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Permutex2varEpi16(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512Permutex2varEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexvarEpi16(src [16]byte, k uint8, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskPermutexvarEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU a+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutexvarEpi16(k uint8, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskzPermutexvarEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU idx+4(FP),X1
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func permutexvarEpi16(idx [16]byte, a [16]byte) [16]byte
TEXT ·permutexvarEpi16(SB),7,$0
	MOVOU idx+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskPermutexvarEpi16(src [32]byte, k uint16, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskPermutexvarEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzPermutexvarEpi16(k uint16, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskzPermutexvarEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256PermutexvarEpi16(idx [32]byte, a [32]byte) [32]byte
TEXT ·m256PermutexvarEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskPermutexvarEpi16(src [64]byte, k uint32, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutexvarEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzPermutexvarEpi16(k uint32, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskzPermutexvarEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512PermutexvarEpi16(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutexvarEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SadEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSet1Epi16(src [16]byte, k uint8, a int16) [16]byte
TEXT ·maskSet1Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVW a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskzSet1Epi16(k uint8, a int16) [16]byte
TEXT ·maskzSet1Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVW a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func m256MaskSet1Epi16(src [32]byte, k uint16, a int16) [32]byte
TEXT ·m256MaskSet1Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSet1Epi16(k uint16, a int16) [32]byte
TEXT ·m256MaskzSet1Epi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVW a+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+8(FP)
	RET

// func m512MaskSet1Epi16(src [64]byte, k uint32, a int16) [64]byte
TEXT ·m512MaskSet1Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSet1Epi16(k uint32, a int16) [64]byte
TEXT ·m512MaskzSet1Epi16(SB),7,$0
	MOVL k+0(FP),R8
	MOVW a+4(FP),R9

	//TODO: Code missing

	MOV Z0, ret+8(FP)
	RET

// func maskSet1Epi8(src [16]byte, k uint16, a byte) [16]byte
TEXT ·maskSet1Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVB a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskzSet1Epi8(k uint16, a byte) [16]byte
TEXT ·maskzSet1Epi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVB a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func m256MaskSet1Epi8(src [32]byte, k uint32, a byte) [32]byte
TEXT ·m256MaskSet1Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSet1Epi8(k uint32, a byte) [32]byte
TEXT ·m256MaskzSet1Epi8(SB),7,$0
	MOVL k+0(FP),R8
	MOVB a+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+8(FP)
	RET

// func m512MaskSet1Epi8(src [64]byte, k uint64, a byte) [64]byte
TEXT ·m512MaskSet1Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSet1Epi8(k uint64, a byte) [64]byte
TEXT ·m512MaskzSet1Epi8(SB),7,$0
	MOVQ k+0(FP),R8
	MOVB a+8(FP),R9

	//TODO: Code missing

	MOV Z0, ret+12(FP)
	RET

// func maskShuffleEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskShuffleEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzShuffleEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzShuffleEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskShuffleEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskShuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzShuffleEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzShuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskShuffleEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskShuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzShuffleEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzShuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512ShuffleEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512ShuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShufflehiEpi16(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskShufflehiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzShufflehiEpi16(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzShufflehiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func m256MaskShufflehiEpi16(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskShufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzShufflehiEpi16(k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzShufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskShufflehiEpi16(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskShufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzShufflehiEpi16(k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzShufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512ShufflehiEpi16(a [64]byte, imm8 int) [64]byte
TEXT ·m512ShufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShuffleloEpi16(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskShuffleloEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzShuffleloEpi16(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzShuffleloEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func m256MaskShuffleloEpi16(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskShuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzShuffleloEpi16(k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzShuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskShuffleloEpi16(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskShuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzShuffleloEpi16(k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzShuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512ShuffleloEpi16(a [64]byte, imm8 int) [64]byte
TEXT ·m512ShuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSllEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSllEpi16(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSllEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSllEpi16(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SllEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSlliEpi16(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSlliEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSlliEpi16(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSlliEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func m256MaskSlliEpi16(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSlliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSlliEpi16(k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSlliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSlliEpi16(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSlliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSlliEpi16(k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSlliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SlliEpi16(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SlliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllvEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllvEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllvEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllvEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sllvEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sllvEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskSllvEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSllvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSllvEpi16(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSllvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SllvEpi16(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SllvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSllvEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSllvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSllvEpi16(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSllvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SllvEpi16(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SllvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSraEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSraEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSraEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSraEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSraEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSraEpi16(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSraEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSraEpi16(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SraEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSraiEpi16(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSraiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSraiEpi16(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSraiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func m256MaskSraiEpi16(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskSraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSraiEpi16(k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·m256MaskzSraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSraiEpi16(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSraiEpi16(k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskzSraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SraiEpi16(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSravEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSravEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSravEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSravEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sravEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sravEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskSravEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSravEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSravEpi16(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSravEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SravEpi16(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SravEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSravEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSravEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSravEpi16(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSravEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SravEpi16(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SravEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrlEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSrlEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSrlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSrlEpi16(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSrlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSrlEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSrlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSrlEpi16(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSrlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SrlEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SrlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrliEpi16(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskSrliEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzSrliEpi16(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzSrliEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func m256MaskSrliEpi16(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskSrliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSrliEpi16(k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·m256MaskzSrliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSrliEpi16(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSrliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSrliEpi16(k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·m512MaskzSrliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SrliEpi16(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SrliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrlvEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlvEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlvEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlvEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srlvEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·srlvEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskSrlvEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSrlvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSrlvEpi16(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSrlvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrlvEpi16(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SrlvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSrlvEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSrlvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSrlvEpi16(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSrlvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SrlvEpi16(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SrlvEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskStoreuEpi16(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi16(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func m256MaskStoreuEpi16(mem_addr uintptr, k uint16, a [32]byte) 
TEXT ·m256MaskStoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func m512MaskStoreuEpi16(mem_addr uintptr, k uint32, a [64]byte) 
TEXT ·m512MaskStoreuEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	RET

// func maskStoreuEpi8(mem_addr uintptr, k uint16, a [16]byte) 
TEXT ·maskStoreuEpi8(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVW k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func m256MaskStoreuEpi8(mem_addr uintptr, k uint32, a [32]byte) 
TEXT ·m256MaskStoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func m512MaskStoreuEpi8(mem_addr uintptr, k uint64, a [64]byte) 
TEXT ·m512MaskStoreuEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	RET

// func maskSubEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSubEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSubEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSubEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSubEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SubEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSubEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSubEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSubEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSubEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SubEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSubsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSubsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSubsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSubsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SubsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubsEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSubsEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSubsEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSubsEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSubsEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SubsEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubsEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSubsEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSubsEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSubsEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSubsEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SubsEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubsEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskSubsEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzSubsEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskSubsEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzSubsEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512SubsEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTestEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·testEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskTestEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256TestEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256TestEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512MaskTestEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskTestEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512TestEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512TestEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskTestEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskTestEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func testEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·testEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func m256MaskTestEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskTestEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256TestEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256TestEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskTestEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskTestEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512TestEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512TestEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskTestnEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestnEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testnEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·testnEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestnEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskTestnEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m256TestnEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256TestnEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512MaskTestnEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskTestnEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512TestnEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512TestnEpi16Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskTestnEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskTestnEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func testnEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·testnEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func m256MaskTestnEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskTestnEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m256TestnEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256TestnEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskTestnEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskTestnEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512TestnEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512TestnEpi8Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskUnpackhiEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskUnpackhiEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzUnpackhiEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskUnpackhiEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzUnpackhiEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512UnpackhiEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackhiEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhiEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskUnpackhiEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzUnpackhiEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskUnpackhiEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzUnpackhiEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512UnpackhiEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackloEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskUnpackloEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzUnpackloEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskUnpackloEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzUnpackloEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512UnpackloEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackloEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackloEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func m256MaskUnpackloEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskzUnpackloEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m512MaskUnpackloEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512MaskzUnpackloEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512UnpackloEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

