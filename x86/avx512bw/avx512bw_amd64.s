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

// func maskAbsEpi161(src [32]byte, k uint16, a [32]byte) [32]byte
TEXT ·maskAbsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAbsEpi161(k uint16, a [32]byte) [32]byte
TEXT ·maskzAbsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func absEpi16(a [64]byte) [64]byte
TEXT ·absEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAbsEpi162(src [64]byte, k uint32, a [64]byte) [64]byte
TEXT ·maskAbsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAbsEpi162(k uint32, a [64]byte) [64]byte
TEXT ·maskzAbsEpi162(SB),7,$0
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

// func maskAbsEpi81(src [32]byte, k uint32, a [32]byte) [32]byte
TEXT ·maskAbsEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAbsEpi81(k uint32, a [32]byte) [32]byte
TEXT ·maskzAbsEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func absEpi8(a [64]byte) [64]byte
TEXT ·absEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAbsEpi82(src [64]byte, k uint64, a [64]byte) [64]byte
TEXT ·maskAbsEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAbsEpi82(k uint64, a [64]byte) [64]byte
TEXT ·maskzAbsEpi82(SB),7,$0
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

// func maskAddEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·addEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddEpi162(SB),7,$0
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

// func maskAddEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·addEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddEpi82(SB),7,$0
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

// func maskAddsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·addsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddsEpi162(SB),7,$0
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

// func maskAddsEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddsEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddsEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddsEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·addsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddsEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddsEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddsEpi82(SB),7,$0
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

// func maskAddsEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddsEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddsEpu161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddsEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·addsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddsEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddsEpu162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddsEpu162(SB),7,$0
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

// func maskAddsEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddsEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddsEpu81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddsEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·addsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddsEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddsEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddsEpu82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddsEpu82(SB),7,$0
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

// func maskAlignrEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·maskAlignrEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAlignrEpi81(k uint32, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·maskzAlignrEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func alignrEpi8(a [64]byte, b [64]byte, count int) [64]byte
TEXT ·alignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAlignrEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·maskAlignrEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAlignrEpi82(k uint64, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·maskzAlignrEpi82(SB),7,$0
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

// func maskAvgEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAvgEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAvgEpu161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAvgEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func avgEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·avgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAvgEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAvgEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAvgEpu162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAvgEpu162(SB),7,$0
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

// func maskAvgEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAvgEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAvgEpu81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAvgEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func avgEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·avgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAvgEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAvgEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAvgEpu82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAvgEpu82(SB),7,$0
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

// func maskBlendEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskBlendEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskBlendEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskBlendEpi162(SB),7,$0
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

// func maskBlendEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskBlendEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskBlendEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskBlendEpi82(SB),7,$0
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

// func maskBroadcastbEpi81(src [32]byte, k uint32, a [16]byte) [32]byte
TEXT ·maskBroadcastbEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastbEpi81(k uint32, a [16]byte) [32]byte
TEXT ·maskzBroadcastbEpi81(SB),7,$0
	MOVL k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastbEpi8(a [16]byte) [64]byte
TEXT ·broadcastbEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastbEpi82(src [64]byte, k uint64, a [16]byte) [64]byte
TEXT ·maskBroadcastbEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastbEpi82(k uint64, a [16]byte) [64]byte
TEXT ·maskzBroadcastbEpi82(SB),7,$0
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

// func maskBroadcastwEpi161(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·maskBroadcastwEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastwEpi161(k uint16, a [16]byte) [32]byte
TEXT ·maskzBroadcastwEpi161(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastwEpi16(a [16]byte) [64]byte
TEXT ·broadcastwEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastwEpi162(src [64]byte, k uint32, a [16]byte) [64]byte
TEXT ·maskBroadcastwEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastwEpi162(k uint32, a [16]byte) [64]byte
TEXT ·maskzBroadcastwEpi162(SB),7,$0
	MOVL k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func bslliEpi128(a [64]byte, imm8 int) [64]byte
TEXT ·bslliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func bsrliEpi128(a [64]byte, imm8 int) [64]byte
TEXT ·bsrliEpi128(SB),7,$0
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

// func cmpEpi16Mask1(a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·cmpEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpEpi16Mask1(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·maskCmpEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpEpi16Mask2(a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·cmpEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpEpi16Mask2(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·maskCmpEpi16Mask2(SB),7,$0
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

// func cmpEpi8Mask1(a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·cmpEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpEpi8Mask1(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·maskCmpEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpEpi8Mask2(a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·cmpEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpEpi8Mask2(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·maskCmpEpi8Mask2(SB),7,$0
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

// func cmpEpu16Mask1(a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·cmpEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpEpu16Mask1(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16
TEXT ·maskCmpEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpEpu16Mask2(a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·cmpEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpEpu16Mask2(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32
TEXT ·maskCmpEpu16Mask2(SB),7,$0
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

// func cmpEpu8Mask1(a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·cmpEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpEpu8Mask1(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32
TEXT ·maskCmpEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpEpu8Mask2(a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·cmpEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpEpu8Mask2(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64
TEXT ·maskCmpEpu8Mask2(SB),7,$0
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

// func cmpeqEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpeqEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpeqEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpeqEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpeqEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpeqEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpeqEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpeqEpi16Mask2(SB),7,$0
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

// func cmpeqEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpeqEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpeqEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpeqEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpeqEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpeqEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpeqEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpeqEpi8Mask2(SB),7,$0
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

// func cmpeqEpu16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpeqEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpeqEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpeqEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpeqEpu16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpeqEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpeqEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpeqEpu16Mask2(SB),7,$0
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

// func cmpeqEpu8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpeqEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpeqEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpeqEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpeqEpu8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpeqEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpeqEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpeqEpu8Mask2(SB),7,$0
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

// func cmpgeEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpgeEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpgeEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpgeEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpgeEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpgeEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgeEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpgeEpi16Mask2(SB),7,$0
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

// func cmpgeEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpgeEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgeEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpgeEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgeEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpgeEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpgeEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpgeEpi8Mask2(SB),7,$0
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

// func cmpgeEpu16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpgeEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpgeEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpgeEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpgeEpu16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpgeEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgeEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpgeEpu16Mask2(SB),7,$0
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

// func cmpgeEpu8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpgeEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgeEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpgeEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgeEpu8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpgeEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpgeEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpgeEpu8Mask2(SB),7,$0
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

// func cmpgtEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpgtEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpgtEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpgtEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpgtEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpgtEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgtEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpgtEpi16Mask2(SB),7,$0
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

// func cmpgtEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpgtEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgtEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpgtEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgtEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpgtEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpgtEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpgtEpi8Mask2(SB),7,$0
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

// func cmpgtEpu16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpgtEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpgtEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpgtEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpgtEpu16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpgtEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgtEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpgtEpu16Mask2(SB),7,$0
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

// func cmpgtEpu8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpgtEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpgtEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpgtEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpgtEpu8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpgtEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpgtEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpgtEpu8Mask2(SB),7,$0
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

// func cmpleEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpleEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpleEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpleEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpleEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpleEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpleEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpleEpi16Mask2(SB),7,$0
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

// func cmpleEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpleEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpleEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpleEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpleEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpleEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpleEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpleEpi8Mask2(SB),7,$0
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

// func cmpleEpu16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpleEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpleEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpleEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpleEpu16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpleEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpleEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpleEpu16Mask2(SB),7,$0
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

// func cmpleEpu8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpleEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpleEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpleEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpleEpu8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpleEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpleEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpleEpu8Mask2(SB),7,$0
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

// func cmpltEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpltEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpltEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpltEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpltEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpltEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpltEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpltEpi16Mask2(SB),7,$0
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

// func cmpltEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpltEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpltEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpltEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpltEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpltEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpltEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpltEpi8Mask2(SB),7,$0
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

// func cmpltEpu16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpltEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpltEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpltEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpltEpu16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpltEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpltEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpltEpu16Mask2(SB),7,$0
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

// func cmpltEpu8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpltEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpltEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpltEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpltEpu8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpltEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpltEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpltEpu8Mask2(SB),7,$0
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

// func cmpneqEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpneqEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpneqEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpneqEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpneqEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpneqEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpneqEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpneqEpi16Mask2(SB),7,$0
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

// func cmpneqEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpneqEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpneqEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpneqEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpneqEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpneqEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpneqEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpneqEpi8Mask2(SB),7,$0
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

// func cmpneqEpu16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·cmpneqEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpneqEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskCmpneqEpu16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpneqEpu16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·cmpneqEpu16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpneqEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskCmpneqEpu16Mask2(SB),7,$0
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

// func cmpneqEpu8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·cmpneqEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskCmpneqEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskCmpneqEpu8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func cmpneqEpu8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·cmpneqEpu8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func maskCmpneqEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskCmpneqEpu8Mask2(SB),7,$0
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

// func cvtepi16Epi81(a [32]byte) [16]byte
TEXT ·cvtepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi16Epi81(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·maskCvtepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi16Epi81(k uint16, a [32]byte) [16]byte
TEXT ·maskzCvtepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi16Epi82(a [64]byte) [32]byte
TEXT ·cvtepi16Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi16Epi82(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·maskCvtepi16Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi16Epi82(k uint32, a [64]byte) [32]byte
TEXT ·maskzCvtepi16Epi82(SB),7,$0
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

// func maskCvtepi16StoreuEpi81(base_addr uintptr, k uint16, a [32]byte) 
TEXT ·maskCvtepi16StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maskCvtepi16StoreuEpi82(base_addr uintptr, k uint32, a [64]byte) 
TEXT ·maskCvtepi16StoreuEpi82(SB),7,$0
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

// func maskCvtepi8Epi161(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·maskCvtepi8Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi8Epi161(k uint16, a [16]byte) [32]byte
TEXT ·maskzCvtepi8Epi161(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi8Epi16(a [32]byte) [64]byte
TEXT ·cvtepi8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi8Epi162(src [64]byte, k uint32, a [32]byte) [64]byte
TEXT ·maskCvtepi8Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi8Epi162(k uint32, a [32]byte) [64]byte
TEXT ·maskzCvtepi8Epi162(SB),7,$0
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

// func maskCvtepu8Epi161(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·maskCvtepu8Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu8Epi161(k uint16, a [16]byte) [32]byte
TEXT ·maskzCvtepu8Epi161(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu8Epi16(a [32]byte) [64]byte
TEXT ·cvtepu8Epi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu8Epi162(src [64]byte, k uint32, a [32]byte) [64]byte
TEXT ·maskCvtepu8Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu8Epi162(k uint32, a [32]byte) [64]byte
TEXT ·maskzCvtepu8Epi162(SB),7,$0
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

// func cvtsepi16Epi81(a [32]byte) [16]byte
TEXT ·cvtsepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi16Epi81(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·maskCvtsepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi16Epi81(k uint16, a [32]byte) [16]byte
TEXT ·maskzCvtsepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi16Epi82(a [64]byte) [32]byte
TEXT ·cvtsepi16Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtsepi16Epi82(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·maskCvtsepi16Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtsepi16Epi82(k uint32, a [64]byte) [32]byte
TEXT ·maskzCvtsepi16Epi82(SB),7,$0
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

// func maskCvtsepi16StoreuEpi81(base_addr uintptr, k uint16, a [32]byte) 
TEXT ·maskCvtsepi16StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maskCvtsepi16StoreuEpi82(base_addr uintptr, k uint32, a [64]byte) 
TEXT ·maskCvtsepi16StoreuEpi82(SB),7,$0
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

// func cvtusepi16Epi81(a [32]byte) [16]byte
TEXT ·cvtusepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi16Epi81(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·maskCvtusepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi16Epi81(k uint16, a [32]byte) [16]byte
TEXT ·maskzCvtusepi16Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi16Epi82(a [64]byte) [32]byte
TEXT ·cvtusepi16Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtusepi16Epi82(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·maskCvtusepi16Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtusepi16Epi82(k uint32, a [64]byte) [32]byte
TEXT ·maskzCvtusepi16Epi82(SB),7,$0
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

// func maskCvtusepi16StoreuEpi81(base_addr uintptr, k uint16, a [32]byte) 
TEXT ·maskCvtusepi16StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maskCvtusepi16StoreuEpi82(base_addr uintptr, k uint32, a [64]byte) 
TEXT ·maskCvtusepi16StoreuEpi82(SB),7,$0
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

// func dbsadEpu81(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·dbsadEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskDbsadEpu81(src [32]byte, k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskDbsadEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzDbsadEpu81(k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskzDbsadEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func dbsadEpu82(a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·dbsadEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDbsadEpu82(src [64]byte, k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskDbsadEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzDbsadEpu82(k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskzDbsadEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func kunpackd(a uint64, b uint64) uint64
TEXT ·kunpackd(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func kunpackw(a uint32, b uint32) uint32
TEXT ·kunpackw(SB),7,$0
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

// func maskLoaduEpi161(src [32]byte, k uint16, mem_addr uintptr) [32]byte
TEXT ·maskLoaduEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoaduEpi161(k uint16, mem_addr uintptr) [32]byte
TEXT ·maskzLoaduEpi161(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+12(FP)
	RET

// func maskLoaduEpi162(src [64]byte, k uint32, mem_addr uintptr) [64]byte
TEXT ·maskLoaduEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLoaduEpi162(k uint32, mem_addr uintptr) [64]byte
TEXT ·maskzLoaduEpi162(SB),7,$0
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

// func maskLoaduEpi81(src [32]byte, k uint32, mem_addr uintptr) [32]byte
TEXT ·maskLoaduEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoaduEpi81(k uint32, mem_addr uintptr) [32]byte
TEXT ·maskzLoaduEpi81(SB),7,$0
	MOVL k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+12(FP)
	RET

// func maskLoaduEpi82(src [64]byte, k uint64, mem_addr uintptr) [64]byte
TEXT ·maskLoaduEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLoaduEpi82(k uint64, mem_addr uintptr) [64]byte
TEXT ·maskzLoaduEpi82(SB),7,$0
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

// func maskMaddEpi161(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaddEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaddEpi161(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaddEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maddEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·maddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaddEpi162(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaddEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaddEpi162(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaddEpi162(SB),7,$0
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

// func maskMaddubsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaddubsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaddubsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaddubsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maddubsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·maddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaddubsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaddubsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaddubsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaddubsEpi162(SB),7,$0
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

// func maskMaxEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaxEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·maxEpi16(SB),7,$0
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

// func maskMaxEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaxEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·maxEpi8(SB),7,$0
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

// func maskMaxEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpu161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaxEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxEpu162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·maxEpu16(SB),7,$0
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

// func maskMaxEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpu81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaxEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxEpu82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·maxEpu8(SB),7,$0
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

// func maskMinEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMinEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·minEpi16(SB),7,$0
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

// func maskMinEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMinEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·minEpi8(SB),7,$0
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

// func maskMinEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpu161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMinEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinEpu162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·minEpu16(SB),7,$0
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

// func maskMinEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpu81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMinEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinEpu82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·minEpu8(SB),7,$0
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

// func maskMovEpi161(src [32]byte, k uint16, a [32]byte) [32]byte
TEXT ·maskMovEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovEpi161(k uint16, a [32]byte) [32]byte
TEXT ·maskzMovEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMovEpi162(src [64]byte, k uint32, a [64]byte) [64]byte
TEXT ·maskMovEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMovEpi162(k uint32, a [64]byte) [64]byte
TEXT ·maskzMovEpi162(SB),7,$0
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

// func maskMovEpi81(src [32]byte, k uint32, a [32]byte) [32]byte
TEXT ·maskMovEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovEpi81(k uint32, a [32]byte) [32]byte
TEXT ·maskzMovEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMovEpi82(src [64]byte, k uint64, a [64]byte) [64]byte
TEXT ·maskMovEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMovEpi82(k uint64, a [64]byte) [64]byte
TEXT ·maskzMovEpi82(SB),7,$0
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

// func movepi16Mask1(a [32]byte) uint16
TEXT ·movepi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func movepi16Mask2(a [64]byte) uint32
TEXT ·movepi16Mask2(SB),7,$0
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

// func movepi8Mask1(a [32]byte) uint32
TEXT ·movepi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func movepi8Mask2(a [64]byte) uint64
TEXT ·movepi8Mask2(SB),7,$0
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

// func movmEpi161(k uint16) [32]byte
TEXT ·movmEpi161(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func movmEpi162(k uint32) [64]byte
TEXT ·movmEpi162(SB),7,$0
	MOVL k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func movmEpi8(k uint32) [32]byte
TEXT ·movmEpi8(SB),7,$0
	MOVL k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func movmEpi81(k uint64) [64]byte
TEXT ·movmEpi81(SB),7,$0
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

// func maskMulhiEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulhiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulhiEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulhiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulhiEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulhiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulhiEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulhiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulhiEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·mulhiEpi16(SB),7,$0
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

// func maskMulhiEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulhiEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulhiEpu161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulhiEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulhiEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulhiEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulhiEpu162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulhiEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulhiEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·mulhiEpu16(SB),7,$0
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

// func maskMulhrsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulhrsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulhrsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulhrsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulhrsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulhrsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulhrsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulhrsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulhrsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·mulhrsEpi16(SB),7,$0
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

// func maskMulloEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulloEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulloEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulloEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulloEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulloEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulloEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulloEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulloEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·mulloEpi16(SB),7,$0
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

// func maskPacksEpi161(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskPacksEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPacksEpi161(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzPacksEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPacksEpi162(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskPacksEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPacksEpi162(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzPacksEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func packsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·packsEpi16(SB),7,$0
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

// func maskPacksEpi321(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskPacksEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPacksEpi321(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzPacksEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPacksEpi322(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskPacksEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPacksEpi322(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzPacksEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func packsEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·packsEpi32(SB),7,$0
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

// func maskPackusEpi161(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskPackusEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPackusEpi161(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzPackusEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPackusEpi162(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskPackusEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPackusEpi162(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzPackusEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func packusEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·packusEpi16(SB),7,$0
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

// func maskPackusEpi321(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskPackusEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPackusEpi321(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzPackusEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPackusEpi322(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskPackusEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPackusEpi322(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzPackusEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func packusEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·packusEpi32(SB),7,$0
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

// func maskPermutex2varEpi161(a [32]byte, k uint16, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskPermutex2varEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask2Permutex2varEpi161(a [32]byte, idx [32]byte, k uint16, b [32]byte) [32]byte
TEXT ·mask2Permutex2varEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutex2varEpi161(k uint16, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskzPermutex2varEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutex2varEpi161(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·permutex2varEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutex2varEpi162(a [64]byte, k uint32, idx [64]byte, b [64]byte) [64]byte
TEXT ·maskPermutex2varEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask2Permutex2varEpi162(a [64]byte, idx [64]byte, k uint32, b [64]byte) [64]byte
TEXT ·mask2Permutex2varEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutex2varEpi162(k uint32, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·maskzPermutex2varEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutex2varEpi162(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·permutex2varEpi162(SB),7,$0
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

// func maskPermutexvarEpi161(src [32]byte, k uint16, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskPermutexvarEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexvarEpi161(k uint16, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskzPermutexvarEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexvarEpi161(idx [32]byte, a [32]byte) [32]byte
TEXT ·permutexvarEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexvarEpi162(src [64]byte, k uint32, idx [64]byte, a [64]byte) [64]byte
TEXT ·maskPermutexvarEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexvarEpi162(k uint32, idx [64]byte, a [64]byte) [64]byte
TEXT ·maskzPermutexvarEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexvarEpi162(idx [64]byte, a [64]byte) [64]byte
TEXT ·permutexvarEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sadEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·sadEpu8(SB),7,$0
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

// func maskSet1Epi161(src [32]byte, k uint16, a int16) [32]byte
TEXT ·maskSet1Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSet1Epi161(k uint16, a int16) [32]byte
TEXT ·maskzSet1Epi161(SB),7,$0
	MOVW k+0(FP),R8
	MOVW a+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+8(FP)
	RET

// func maskSet1Epi162(src [64]byte, k uint32, a int16) [64]byte
TEXT ·maskSet1Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSet1Epi162(k uint32, a int16) [64]byte
TEXT ·maskzSet1Epi162(SB),7,$0
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

// func maskSet1Epi81(src [32]byte, k uint32, a byte) [32]byte
TEXT ·maskSet1Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSet1Epi81(k uint32, a byte) [32]byte
TEXT ·maskzSet1Epi81(SB),7,$0
	MOVL k+0(FP),R8
	MOVB a+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+8(FP)
	RET

// func maskSet1Epi82(src [64]byte, k uint64, a byte) [64]byte
TEXT ·maskSet1Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSet1Epi82(k uint64, a byte) [64]byte
TEXT ·maskzSet1Epi82(SB),7,$0
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

// func maskShuffleEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskShuffleEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzShuffleEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShuffleEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskShuffleEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShuffleEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzShuffleEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shuffleEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·shuffleEpi8(SB),7,$0
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

// func maskShufflehiEpi161(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·maskShufflehiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShufflehiEpi161(k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·maskzShufflehiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShufflehiEpi162(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·maskShufflehiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShufflehiEpi162(k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·maskzShufflehiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shufflehiEpi16(a [64]byte, imm8 int) [64]byte
TEXT ·shufflehiEpi16(SB),7,$0
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

// func maskShuffleloEpi161(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·maskShuffleloEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleloEpi161(k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·maskzShuffleloEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShuffleloEpi162(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·maskShuffleloEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShuffleloEpi162(k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·maskzShuffleloEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shuffleloEpi16(a [64]byte, imm8 int) [64]byte
TEXT ·shuffleloEpi16(SB),7,$0
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

// func maskSllEpi161(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSllEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllEpi161(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSllEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSllEpi162(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSllEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSllEpi162(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSllEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sllEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·sllEpi16(SB),7,$0
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

// func maskSlliEpi161(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSlliEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSlliEpi161(k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSlliEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSlliEpi162(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskSlliEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSlliEpi162(k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSlliEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func slliEpi16(a [64]byte, imm8 uint32) [64]byte
TEXT ·slliEpi16(SB),7,$0
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

// func maskSllvEpi161(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSllvEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllvEpi161(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSllvEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sllvEpi161(a [32]byte, count [32]byte) [32]byte
TEXT ·sllvEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSllvEpi162(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·maskSllvEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSllvEpi162(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSllvEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sllvEpi162(a [64]byte, count [64]byte) [64]byte
TEXT ·sllvEpi162(SB),7,$0
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

// func maskSraEpi161(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSraEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraEpi161(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSraEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSraEpi162(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSraEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSraEpi162(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSraEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sraEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·sraEpi16(SB),7,$0
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

// func maskSraiEpi161(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSraiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraiEpi161(k uint16, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSraiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSraiEpi162(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskSraiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSraiEpi162(k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSraiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sraiEpi16(a [64]byte, imm8 uint32) [64]byte
TEXT ·sraiEpi16(SB),7,$0
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

// func maskSravEpi161(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSravEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSravEpi161(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSravEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sravEpi161(a [32]byte, count [32]byte) [32]byte
TEXT ·sravEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSravEpi162(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·maskSravEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSravEpi162(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSravEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sravEpi162(a [64]byte, count [64]byte) [64]byte
TEXT ·sravEpi162(SB),7,$0
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

// func maskSrlEpi161(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSrlEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlEpi161(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSrlEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrlEpi162(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSrlEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrlEpi162(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSrlEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srlEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·srlEpi16(SB),7,$0
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

// func maskSrliEpi161(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·maskSrliEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrliEpi161(k uint16, a [32]byte, imm8 int) [32]byte
TEXT ·maskzSrliEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrliEpi162(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskSrliEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrliEpi162(k uint32, a [64]byte, imm8 int) [64]byte
TEXT ·maskzSrliEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srliEpi16(a [64]byte, imm8 uint32) [64]byte
TEXT ·srliEpi16(SB),7,$0
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

// func maskSrlvEpi161(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSrlvEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlvEpi161(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSrlvEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srlvEpi161(a [32]byte, count [32]byte) [32]byte
TEXT ·srlvEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrlvEpi162(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·maskSrlvEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrlvEpi162(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSrlvEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srlvEpi162(a [64]byte, count [64]byte) [64]byte
TEXT ·srlvEpi162(SB),7,$0
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

// func maskStoreuEpi161(mem_addr uintptr, k uint16, a [32]byte) 
TEXT ·maskStoreuEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maskStoreuEpi162(mem_addr uintptr, k uint32, a [64]byte) 
TEXT ·maskStoreuEpi162(SB),7,$0
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

// func maskStoreuEpi81(mem_addr uintptr, k uint32, a [32]byte) 
TEXT ·maskStoreuEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maskStoreuEpi82(mem_addr uintptr, k uint64, a [64]byte) 
TEXT ·maskStoreuEpi82(SB),7,$0
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

// func maskSubEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·subEpi16(SB),7,$0
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

// func maskSubEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·subEpi8(SB),7,$0
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

// func maskSubsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubsEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubsEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·subsEpi16(SB),7,$0
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

// func maskSubsEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubsEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubsEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubsEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubsEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubsEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubsEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubsEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subsEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·subsEpi8(SB),7,$0
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

// func maskSubsEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubsEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubsEpu161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubsEpu161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubsEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubsEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubsEpu162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubsEpu162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subsEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·subsEpu16(SB),7,$0
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

// func maskSubsEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubsEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubsEpu81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubsEpu81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubsEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubsEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubsEpu82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubsEpu82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subsEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·subsEpu8(SB),7,$0
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

// func maskTestEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskTestEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func testEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·testEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskTestEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskTestEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func testEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·testEpi16Mask2(SB),7,$0
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

// func maskTestEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskTestEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func testEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·testEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskTestEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskTestEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func testEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·testEpi8Mask2(SB),7,$0
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

// func maskTestnEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·maskTestnEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func testnEpi16Mask1(a [32]byte, b [32]byte) uint16
TEXT ·testnEpi16Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskTestnEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·maskTestnEpi16Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func testnEpi16Mask2(a [64]byte, b [64]byte) uint32
TEXT ·testnEpi16Mask2(SB),7,$0
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

// func maskTestnEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·maskTestnEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func testnEpi8Mask1(a [32]byte, b [32]byte) uint32
TEXT ·testnEpi8Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func maskTestnEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·maskTestnEpi8Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func testnEpi8Mask2(a [64]byte, b [64]byte) uint64
TEXT ·testnEpi8Mask2(SB),7,$0
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

// func maskUnpackhiEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackhiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackhiEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackhiEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackhiEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackhiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackhiEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackhiEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackhiEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackhiEpi16(SB),7,$0
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

// func maskUnpackhiEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackhiEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackhiEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackhiEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackhiEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackhiEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackhiEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackhiEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackhiEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackhiEpi8(SB),7,$0
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

// func maskUnpackloEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackloEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackloEpi161(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackloEpi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackloEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackloEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackloEpi162(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackloEpi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackloEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackloEpi16(SB),7,$0
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

// func maskUnpackloEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackloEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackloEpi81(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackloEpi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackloEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackloEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackloEpi82(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackloEpi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackloEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

