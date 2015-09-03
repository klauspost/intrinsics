// func maskAbsEpi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskAbsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPABSW

	MOVOU X2, ret+36(FP)
	RET

// func maskzAbsEpi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzAbsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPABSW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskAbsEpi16(src [32]byte, k uint16, a [32]byte) [32]byte
TEXT ·m256MaskAbsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPABSW

	MOV Y2, ret+68(FP)
	RET

// func m256MaskzAbsEpi16(k uint16, a [32]byte) [32]byte
TEXT ·m256MaskzAbsEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPABSW R8, Y1

	MOV Y1, ret+36(FP)
	RET

// func m512AbsEpi16(a [64]byte) [64]byte
TEXT ·m512AbsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPABSW Z0

	MOV Z0, ret+64(FP)
	RET

// func m512MaskAbsEpi16(src [64]byte, k uint32, a [64]byte) [64]byte
TEXT ·m512MaskAbsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPABSW

	MOV Z2, ret+132(FP)
	RET

// func m512MaskzAbsEpi16(k uint32, a [64]byte) [64]byte
TEXT ·m512MaskzAbsEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1

	// TODO: Code missing - could be:
	// VPABSW R8, Z1

	MOV Z1, ret+68(FP)
	RET

// func maskAbsEpi8(src [16]byte, k uint16, a [16]byte) [16]byte
TEXT ·maskAbsEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPABSB

	MOVOU X2, ret+36(FP)
	RET

// func maskzAbsEpi8(k uint16, a [16]byte) [16]byte
TEXT ·maskzAbsEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPABSB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskAbsEpi8(src [32]byte, k uint32, a [32]byte) [32]byte
TEXT ·m256MaskAbsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPABSB

	MOV Y2, ret+68(FP)
	RET

// func m256MaskzAbsEpi8(k uint32, a [32]byte) [32]byte
TEXT ·m256MaskzAbsEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPABSB R8, Y1

	MOV Y1, ret+36(FP)
	RET

// func m512AbsEpi8(a [64]byte) [64]byte
TEXT ·m512AbsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPABSB Z0

	MOV Z0, ret+64(FP)
	RET

// func m512MaskAbsEpi8(src [64]byte, k uint64, a [64]byte) [64]byte
TEXT ·m512MaskAbsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPABSB

	MOV Z2, ret+136(FP)
	RET

// func m512MaskzAbsEpi8(k uint64, a [64]byte) [64]byte
TEXT ·m512MaskzAbsEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1

	// TODO: Code missing - could be:
	// VPABSB R8, Z1

	MOV Z1, ret+72(FP)
	RET

// func maskAddEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPADDW

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPADDW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPADDW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAddEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPADDW

	MOV Y2, ret+68(FP)
	RET

// func m512AddEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPADDW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAddEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPADDW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzAddEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPADDW

	MOV Z2, ret+132(FP)
	RET

// func maskAddEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPADDB

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPADDB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPADDB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAddEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPADDB

	MOV Y2, ret+68(FP)
	RET

// func m512AddEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPADDB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAddEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPADDB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzAddEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPADDB

	MOV Z2, ret+136(FP)
	RET

// func maskAddsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPADDSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPADDSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPADDSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAddsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPADDSW

	MOV Y2, ret+68(FP)
	RET

// func m512AddsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPADDSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAddsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPADDSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzAddsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPADDSW

	MOV Z2, ret+132(FP)
	RET

// func maskAddsEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPADDSB

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPADDSB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddsEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPADDSB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAddsEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPADDSB

	MOV Y2, ret+68(FP)
	RET

// func m512AddsEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPADDSB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAddsEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPADDSB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzAddsEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPADDSB

	MOV Z2, ret+136(FP)
	RET

// func maskAddsEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPADDUSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPADDUSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddsEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPADDUSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAddsEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpu16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPADDUSW

	MOV Y2, ret+68(FP)
	RET

// func m512AddsEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPADDUSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAddsEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPADDUSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzAddsEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpu16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPADDUSW

	MOV Z2, ret+132(FP)
	RET

// func maskAddsEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddsEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPADDUSB

	MOVOU X3, ret+52(FP)
	RET

// func maskzAddsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddsEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPADDUSB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAddsEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAddsEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPADDUSB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAddsEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAddsEpu8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPADDUSB

	MOV Y2, ret+68(FP)
	RET

// func m512AddsEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddsEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPADDUSB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAddsEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddsEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPADDUSB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzAddsEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAddsEpu8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPADDUSB

	MOV Z2, ret+136(FP)
	RET

// func maskAlignrEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskAlignrEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOVOU X4, ret+60(FP)
	RET

// func maskzAlignrEpi8(k uint16, a [16]byte, b [16]byte, count int) [16]byte
TEXT ·maskzAlignrEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOVOU X3, ret+44(FP)
	RET

// func m256MaskAlignrEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskAlignrEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3
	MOVQ count+100(FP),R12

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOV Y4, ret+108(FP)
	RET

// func m256MaskzAlignrEpi8(k uint32, a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256MaskzAlignrEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2
	MOVQ count+68(FP),R11

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOV Y3, ret+76(FP)
	RET

// func m512AlignrEpi8(a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512AlignrEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	MOVQ count+128(FP),R10

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOV Z2, ret+136(FP)
	RET

// func m512MaskAlignrEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskAlignrEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3
	MOVQ count+200(FP),R12

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOV Z4, ret+208(FP)
	RET

// func m512MaskzAlignrEpi8(k uint64, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskzAlignrEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2
	MOVQ count+136(FP),R11

	// TODO: Code missing - uses instrunction: VPALIGNR

	MOV Z3, ret+144(FP)
	RET

// func maskAvgEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAvgEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPAVGW

	MOVOU X3, ret+52(FP)
	RET

// func maskzAvgEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAvgEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPAVGW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAvgEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAvgEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPAVGW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAvgEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAvgEpu16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPAVGW

	MOV Y2, ret+68(FP)
	RET

// func m512AvgEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AvgEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPAVGW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAvgEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAvgEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPAVGW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzAvgEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAvgEpu16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPAVGW

	MOV Z2, ret+132(FP)
	RET

// func maskAvgEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAvgEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPAVGB

	MOVOU X3, ret+52(FP)
	RET

// func maskzAvgEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAvgEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPAVGB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskAvgEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskAvgEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPAVGB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzAvgEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzAvgEpu8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPAVGB

	MOV Y2, ret+68(FP)
	RET

// func m512AvgEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AvgEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPAVGB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskAvgEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAvgEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPAVGB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzAvgEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzAvgEpu8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPAVGB

	MOV Z2, ret+136(FP)
	RET

// func maskBlendEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPBLENDMW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskBlendEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskBlendEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPBLENDMW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskBlendEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskBlendEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPBLENDMW

	MOV Z2, ret+132(FP)
	RET

// func maskBlendEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPBLENDMB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskBlendEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskBlendEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPBLENDMB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskBlendEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskBlendEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPBLENDMB

	MOV Z2, ret+136(FP)
	RET

// func maskBroadcastbEpi8(src [16]byte, k uint16, a [16]byte) [16]byte
TEXT ·maskBroadcastbEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPBROADCASTB

	MOVOU X2, ret+36(FP)
	RET

// func maskzBroadcastbEpi8(k uint16, a [16]byte) [16]byte
TEXT ·maskzBroadcastbEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPBROADCASTB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskBroadcastbEpi8(src [32]byte, k uint32, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastbEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	MOVOU a+36(FP),X2

	// TODO: Code missing - uses instrunction: VPBROADCASTB

	MOV Y2, ret+52(FP)
	RET

// func m256MaskzBroadcastbEpi8(k uint32, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastbEpi8(SB),7,$0
	MOVL k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPBROADCASTB R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastbEpi8(a [16]byte) [64]byte
TEXT ·m512BroadcastbEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPBROADCASTB X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastbEpi8(src [64]byte, k uint64, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastbEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	MOVOU a+72(FP),X2

	// TODO: Code missing - uses instrunction: VPBROADCASTB

	MOV Z2, ret+88(FP)
	RET

// func m512MaskzBroadcastbEpi8(k uint64, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastbEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	MOVOU a+8(FP),X1

	// TODO: Code missing - could be:
	// VPBROADCASTB R8, X1

	MOV Z1, ret+24(FP)
	RET

// func maskBroadcastwEpi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastwEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPBROADCASTW

	MOVOU X2, ret+36(FP)
	RET

// func maskzBroadcastwEpi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastwEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPBROADCASTW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskBroadcastwEpi16(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·m256MaskBroadcastwEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	MOVOU a+36(FP),X2

	// TODO: Code missing - uses instrunction: VPBROADCASTW

	MOV Y2, ret+52(FP)
	RET

// func m256MaskzBroadcastwEpi16(k uint16, a [16]byte) [32]byte
TEXT ·m256MaskzBroadcastwEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPBROADCASTW R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512BroadcastwEpi16(a [16]byte) [64]byte
TEXT ·m512BroadcastwEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPBROADCASTW X0

	MOV Z0, ret+16(FP)
	RET

// func m512MaskBroadcastwEpi16(src [64]byte, k uint32, a [16]byte) [64]byte
TEXT ·m512MaskBroadcastwEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	MOVOU a+68(FP),X2

	// TODO: Code missing - uses instrunction: VPBROADCASTW

	MOV Z2, ret+84(FP)
	RET

// func m512MaskzBroadcastwEpi16(k uint32, a [16]byte) [64]byte
TEXT ·m512MaskzBroadcastwEpi16(SB),7,$0
	MOVL k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPBROADCASTW R8, X1

	MOV Z1, ret+20(FP)
	RET

// func m512BslliEpi128(a [64]byte, imm8 byte) [64]byte
TEXT ·m512BslliEpi128(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSLLDQ Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func m512BsrliEpi128(a [64]byte, imm8 byte) [64]byte
TEXT ·m512BsrliEpi128(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSRLDQ Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func cmpEpi16Mask(a [16]byte, b [16]byte, imm8 byte) uint8
TEXT ·cmpEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpi16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 byte) uint8
TEXT ·maskCmpEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+40(FP)
	RET

// func m256CmpEpi16Mask(a [32]byte, b [32]byte, imm8 byte) uint16
TEXT ·m256CmpEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m256MaskCmpEpi16Mask(k1 uint16, a [32]byte, b [32]byte, imm8 byte) uint16
TEXT ·m256MaskCmpEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+72(FP)
	RET

// func m512CmpEpi16Mask(a [64]byte, b [64]byte, imm8 byte) uint32
TEXT ·m512CmpEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func m512MaskCmpEpi16Mask(k1 uint32, a [64]byte, b [64]byte, imm8 byte) uint32
TEXT ·m512MaskCmpEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+136(FP)
	RET

// func cmpEpi8Mask(a [16]byte, b [16]byte, imm8 byte) uint16
TEXT ·cmpEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func maskCmpEpi8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 byte) uint16
TEXT ·maskCmpEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+40(FP)
	RET

// func m256CmpEpi8Mask(a [32]byte, b [32]byte, imm8 byte) uint32
TEXT ·m256CmpEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m256MaskCmpEpi8Mask(k1 uint32, a [32]byte, b [32]byte, imm8 byte) uint32
TEXT ·m256MaskCmpEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+72(FP)
	RET

// func m512CmpEpi8Mask(a [64]byte, b [64]byte, imm8 byte) uint64
TEXT ·m512CmpEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+132(FP)
	RET

// func m512MaskCmpEpi8Mask(k1 uint64, a [64]byte, b [64]byte, imm8 byte) uint64
TEXT ·m512MaskCmpEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+140(FP)
	RET

// func cmpEpu16Mask(a [16]byte, b [16]byte, imm8 byte) uint8
TEXT ·cmpEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 byte) uint8
TEXT ·maskCmpEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+40(FP)
	RET

// func m256CmpEpu16Mask(a [32]byte, b [32]byte, imm8 byte) uint16
TEXT ·m256CmpEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m256MaskCmpEpu16Mask(k1 uint16, a [32]byte, b [32]byte, imm8 byte) uint16
TEXT ·m256MaskCmpEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+72(FP)
	RET

// func m512CmpEpu16Mask(a [64]byte, b [64]byte, imm8 byte) uint32
TEXT ·m512CmpEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func m512MaskCmpEpu16Mask(k1 uint32, a [64]byte, b [64]byte, imm8 byte) uint32
TEXT ·m512MaskCmpEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+136(FP)
	RET

// func cmpEpu8Mask(a [16]byte, b [16]byte, imm8 byte) uint16
TEXT ·cmpEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func maskCmpEpu8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 byte) uint16
TEXT ·maskCmpEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+40(FP)
	RET

// func m256CmpEpu8Mask(a [32]byte, b [32]byte, imm8 byte) uint32
TEXT ·m256CmpEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m256MaskCmpEpu8Mask(k1 uint32, a [32]byte, b [32]byte, imm8 byte) uint32
TEXT ·m256MaskCmpEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+72(FP)
	RET

// func m512CmpEpu8Mask(a [64]byte, b [64]byte, imm8 byte) uint64
TEXT ·m512CmpEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+132(FP)
	RET

// func m512MaskCmpEpu8Mask(k1 uint64, a [64]byte, b [64]byte, imm8 byte) uint64
TEXT ·m512MaskCmpEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+140(FP)
	RET

// func cmpeqEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpeqEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpeqEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpeqEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpeqEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpeqEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpeqEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpeqEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func cmpeqEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpeqEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpeqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpeqEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpeqEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpeqEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpeqEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpeqEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpeqEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpeqEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpeqEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpeqEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+136(FP)
	RET

// func cmpeqEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpeqEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpeqEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpeqEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpeqEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpeqEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpeqEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpeqEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpeqEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func cmpeqEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpeqEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpeqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpeqEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpeqEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpeqEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpeqEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpeqEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpeqEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpeqEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpeqEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpeqEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+136(FP)
	RET

// func cmpgeEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgeEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpgeEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgeEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpgeEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgeEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpgeEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgeEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func cmpgeEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgeEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgeEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgeEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgeEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgeEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpgeEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgeEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpgeEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgeEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpgeEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgeEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+136(FP)
	RET

// func cmpgeEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgeEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgeEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpgeEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgeEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpgeEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgeEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpgeEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgeEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func cmpgeEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgeEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgeEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgeEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgeEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgeEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpgeEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgeEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpgeEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgeEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpgeEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgeEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+136(FP)
	RET

// func cmpgtEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgtEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpgtEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgtEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpgtEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgtEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpgtEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgtEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func cmpgtEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgtEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgtEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgtEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgtEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgtEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpgtEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgtEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpgtEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgtEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpgtEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgtEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+136(FP)
	RET

// func cmpgtEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpgtEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpgtEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpgtEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpgtEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpgtEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpgtEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpgtEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpgtEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func cmpgtEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpgtEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgtEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpgtEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpgtEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpgtEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpgtEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpgtEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpgtEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpgtEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpgtEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpgtEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+136(FP)
	RET

// func cmpleEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpleEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpleEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpleEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpleEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpleEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpleEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpleEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func cmpleEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpleEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpleEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpleEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpleEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpleEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpleEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpleEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpleEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpleEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpleEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpleEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+136(FP)
	RET

// func cmpleEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpleEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpleEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpleEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpleEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpleEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpleEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpleEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpleEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func cmpleEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpleEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpleEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpleEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpleEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpleEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpleEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpleEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpleEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpleEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpleEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpleEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+136(FP)
	RET

// func cmpltEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpltEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpltEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpltEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpltEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpltEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpltEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpltEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func cmpltEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpltEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpltEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpltEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpltEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpltEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpltEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpltEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpltEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpltEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpltEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpltEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+136(FP)
	RET

// func cmpltEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpltEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpltEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpltEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpltEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpltEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpltEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpltEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpltEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func cmpltEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpltEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpltEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpltEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpltEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpltEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpltEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpltEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpltEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpltEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpltEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpltEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+136(FP)
	RET

// func cmpneqEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpneqEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpneqEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpneqEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpneqEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpneqEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpneqEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpneqEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPW

	MOVL $0, ret+132(FP)
	RET

// func cmpneqEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpneqEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpneqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpneqEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpneqEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpneqEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpneqEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpneqEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpneqEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpneqEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpneqEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpneqEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPB

	MOVQ $0, ret+136(FP)
	RET

// func cmpneqEpu16Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVB $0, ret+36(FP)
	RET

// func m256CmpneqEpu16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256CmpneqEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m256MaskCmpneqEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskCmpneqEpu16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVW $0, ret+68(FP)
	RET

// func m512CmpneqEpu16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512CmpneqEpu16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func m512MaskCmpneqEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskCmpneqEpu16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUW

	MOVL $0, ret+132(FP)
	RET

// func cmpneqEpu8Mask(a [16]byte, b [16]byte) uint16
TEXT ·cmpneqEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPCMPUB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func maskCmpneqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskCmpneqEpu8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVW $0, ret+36(FP)
	RET

// func m256CmpneqEpu8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256CmpneqEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPCMPUB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m256MaskCmpneqEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskCmpneqEpu8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVL $0, ret+68(FP)
	RET

// func m512CmpneqEpu8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512CmpneqEpu8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPCMPUB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func m512MaskCmpneqEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskCmpneqEpu8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPCMPUB

	MOVQ $0, ret+136(FP)
	RET

// func cvtepi16Epi8(a [16]byte) [16]byte
TEXT ·cvtepi16Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPMOVWB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVWB

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi16Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi16Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVWB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtepi16Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVWB Y0

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskCvtepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·m256MaskCvtepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+20(FP),Y2

	// TODO: Code missing - uses instrunction: VPMOVWB

	MOVOU X2, ret+52(FP)
	RET

// func m256MaskzCvtepi16Epi8(k uint16, a [32]byte) [16]byte
TEXT ·m256MaskzCvtepi16Epi8(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPMOVWB R8, Y1

	MOVOU X1, ret+36(FP)
	RET

// func m512Cvtepi16Epi8(a [64]byte) [32]byte
TEXT ·m512Cvtepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPMOVWB Z0

	MOV Y0, ret+64(FP)
	RET

// func m512MaskCvtepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·m512MaskCvtepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+36(FP),Z2

	// TODO: Code missing - uses instrunction: VPMOVWB

	MOV Y2, ret+100(FP)
	RET

// func m512MaskzCvtepi16Epi8(k uint32, a [64]byte) [32]byte
TEXT ·m512MaskzCvtepi16Epi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1

	// TODO: Code missing - could be:
	// VPMOVWB R8, Z1

	MOV Y1, ret+68(FP)
	RET

// func maskCvtepi8Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi8Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVSXBW

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepi8Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi8Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVSXBW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepi8Epi16(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·m256MaskCvtepi8Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	MOVOU a+36(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVSXBW

	MOV Y2, ret+52(FP)
	RET

// func m256MaskzCvtepi8Epi16(k uint16, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepi8Epi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVSXBW R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepi8Epi16(a [32]byte) [64]byte
TEXT ·m512Cvtepi8Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVSXBW Y0

	MOV Z0, ret+32(FP)
	RET

// func m512MaskCvtepi8Epi16(src [64]byte, k uint32, a [32]byte) [64]byte
TEXT ·m512MaskCvtepi8Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+68(FP),Y2

	// TODO: Code missing - uses instrunction: VPMOVSXBW

	MOV Z2, ret+100(FP)
	RET

// func m512MaskzCvtepi8Epi16(k uint32, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepi8Epi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPMOVSXBW R8, Y1

	MOV Z1, ret+36(FP)
	RET

// func maskCvtepu8Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu8Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVZXBW

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtepu8Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu8Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVZXBW R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskCvtepu8Epi16(src [32]byte, k uint16, a [16]byte) [32]byte
TEXT ·m256MaskCvtepu8Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	MOVOU a+36(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVZXBW

	MOV Y2, ret+52(FP)
	RET

// func m256MaskzCvtepu8Epi16(k uint16, a [16]byte) [32]byte
TEXT ·m256MaskzCvtepu8Epi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVZXBW R8, X1

	MOV Y1, ret+20(FP)
	RET

// func m512Cvtepu8Epi16(a [32]byte) [64]byte
TEXT ·m512Cvtepu8Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVZXBW Y0

	MOV Z0, ret+32(FP)
	RET

// func m512MaskCvtepu8Epi16(src [64]byte, k uint32, a [32]byte) [64]byte
TEXT ·m512MaskCvtepu8Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+68(FP),Y2

	// TODO: Code missing - uses instrunction: VPMOVZXBW

	MOV Z2, ret+100(FP)
	RET

// func m512MaskzCvtepu8Epi16(k uint32, a [32]byte) [64]byte
TEXT ·m512MaskzCvtepu8Epi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPMOVZXBW R8, Y1

	MOV Z1, ret+36(FP)
	RET

// func cvtsepi16Epi8(a [16]byte) [16]byte
TEXT ·cvtsepi16Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPMOVSWB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVSWB

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtsepi16Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi16Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVSWB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtsepi16Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtsepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVSWB Y0

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskCvtsepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·m256MaskCvtsepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+20(FP),Y2

	// TODO: Code missing - uses instrunction: VPMOVSWB

	MOVOU X2, ret+52(FP)
	RET

// func m256MaskzCvtsepi16Epi8(k uint16, a [32]byte) [16]byte
TEXT ·m256MaskzCvtsepi16Epi8(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPMOVSWB R8, Y1

	MOVOU X1, ret+36(FP)
	RET

// func m512Cvtsepi16Epi8(a [64]byte) [32]byte
TEXT ·m512Cvtsepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPMOVSWB Z0

	MOV Y0, ret+64(FP)
	RET

// func m512MaskCvtsepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·m512MaskCvtsepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+36(FP),Z2

	// TODO: Code missing - uses instrunction: VPMOVSWB

	MOV Y2, ret+100(FP)
	RET

// func m512MaskzCvtsepi16Epi8(k uint32, a [64]byte) [32]byte
TEXT ·m512MaskzCvtsepi16Epi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1

	// TODO: Code missing - could be:
	// VPMOVSWB R8, Z1

	MOV Y1, ret+68(FP)
	RET

// func cvtusepi16Epi8(a [16]byte) [16]byte
TEXT ·cvtusepi16Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPMOVUSWB X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMOVUSWB

	MOVOU X2, ret+36(FP)
	RET

// func maskzCvtusepi16Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi16Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VPMOVUSWB R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256Cvtusepi16Epi8(a [32]byte) [16]byte
TEXT ·m256Cvtusepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVUSWB Y0

	MOVOU X0, ret+32(FP)
	RET

// func m256MaskCvtusepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte
TEXT ·m256MaskCvtusepi16Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+20(FP),Y2

	// TODO: Code missing - uses instrunction: VPMOVUSWB

	MOVOU X2, ret+52(FP)
	RET

// func m256MaskzCvtusepi16Epi8(k uint16, a [32]byte) [16]byte
TEXT ·m256MaskzCvtusepi16Epi8(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VPMOVUSWB R8, Y1

	MOVOU X1, ret+36(FP)
	RET

// func m512Cvtusepi16Epi8(a [64]byte) [32]byte
TEXT ·m512Cvtusepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPMOVUSWB Z0

	MOV Y0, ret+64(FP)
	RET

// func m512MaskCvtusepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte
TEXT ·m512MaskCvtusepi16Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+36(FP),Z2

	// TODO: Code missing - uses instrunction: VPMOVUSWB

	MOV Y2, ret+100(FP)
	RET

// func m512MaskzCvtusepi16Epi8(k uint32, a [64]byte) [32]byte
TEXT ·m512MaskzCvtusepi16Epi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1

	// TODO: Code missing - could be:
	// VPMOVUSWB R8, Z1

	MOV Y1, ret+68(FP)
	RET

// func dbsadEpu8(a [16]byte, b [16]byte, imm8 byte) [16]byte
TEXT ·dbsadEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOVOU X2, ret+36(FP)
	RET

// func maskDbsadEpu8(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 byte) [16]byte
TEXT ·maskDbsadEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOVOU X4, ret+56(FP)
	RET

// func maskzDbsadEpu8(k uint8, a [16]byte, b [16]byte, imm8 byte) [16]byte
TEXT ·maskzDbsadEpu8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOVOU X3, ret+40(FP)
	RET

// func m256DbsadEpu8(a [32]byte, b [32]byte, imm8 byte) [32]byte
TEXT ·m256DbsadEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOV Y2, ret+68(FP)
	RET

// func m256MaskDbsadEpu8(src [32]byte, k uint16, a [32]byte, b [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskDbsadEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOV Y4, ret+104(FP)
	RET

// func m256MaskzDbsadEpu8(k uint16, a [32]byte, b [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskzDbsadEpu8(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOV Y3, ret+72(FP)
	RET

// func m512DbsadEpu8(a [64]byte, b [64]byte, imm8 byte) [64]byte
TEXT ·m512DbsadEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOV Z2, ret+132(FP)
	RET

// func m512MaskDbsadEpu8(src [64]byte, k uint32, a [64]byte, b [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskDbsadEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOV Z4, ret+200(FP)
	RET

// func m512MaskzDbsadEpu8(k uint32, a [64]byte, b [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskzDbsadEpu8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VDBPSADBW

	MOV Z3, ret+136(FP)
	RET

// func m512Kunpackd(a uint64, b uint64) uint64
TEXT ·m512Kunpackd(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	// TODO: Code missing - could be:
	// KUNPCKDQ R8, R9

	MOVQ $0, ret+16(FP)
	RET

// func m512Kunpackw(a uint32, b uint32) uint32
TEXT ·m512Kunpackw(SB),7,$0
	MOVL a+0(FP),R8
	MOVL b+4(FP),R9

	// TODO: Code missing - could be:
	// KUNPCKWD R8, R9

	MOVL $0, ret+8(FP)
	RET

// func maskMaddEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaddEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMADDWD

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaddEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMADDWD

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaddEpi16(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaddEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVB k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMADDWD

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMaddEpi16(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaddEpi16(SB),7,$0
	MOVB k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMADDWD

	MOV Y2, ret+68(FP)
	RET

// func m512MaddEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaddEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMADDWD Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskMaddEpi16(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaddEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVW k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMADDWD

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMaddEpi16(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaddEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMADDWD

	MOV Z2, ret+132(FP)
	RET

// func maskMaddubsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaddubsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMADDUBSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaddubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaddubsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMADDUBSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaddubsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaddubsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMADDUBSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMaddubsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaddubsEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMADDUBSW

	MOV Y2, ret+68(FP)
	RET

// func m512MaddubsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaddubsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMADDUBSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskMaddubsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaddubsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMADDUBSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMaddubsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaddubsEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMADDUBSW

	MOV Z2, ret+132(FP)
	RET

// func maskMaxEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMAXSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMAXSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMAXSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMaxEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMAXSW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMaxEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMAXSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMaxEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMAXSW

	MOV Z2, ret+132(FP)
	RET

// func m512MaxEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMAXSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMaxEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMAXSB

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMAXSB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMAXSB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMaxEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMAXSB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMaxEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPMAXSB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzMaxEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPMAXSB

	MOV Z2, ret+136(FP)
	RET

// func m512MaxEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMAXSB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMaxEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMAXUW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMAXUW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMAXUW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMaxEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpu16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMAXUW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMaxEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMAXUW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMaxEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpu16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMAXUW

	MOV Z2, ret+132(FP)
	RET

// func m512MaxEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMAXUW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMaxEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMAXUB

	MOVOU X3, ret+52(FP)
	RET

// func maskzMaxEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMAXUB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMaxEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMaxEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMAXUB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMaxEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMaxEpu8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMAXUB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMaxEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPMAXUB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzMaxEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMaxEpu8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPMAXUB

	MOV Z2, ret+136(FP)
	RET

// func m512MaxEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMAXUB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMinEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMINSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMINSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMINSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMinEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMINSW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMinEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMINSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMinEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMINSW

	MOV Z2, ret+132(FP)
	RET

// func m512MinEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMINSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMinEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMINSB

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMINSB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMINSB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMinEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMINSB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMinEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPMINSB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzMinEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPMINSB

	MOV Z2, ret+136(FP)
	RET

// func m512MinEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMINSB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMinEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMINUW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMINUW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMINUW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMinEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpu16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMINUW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMinEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMINUW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMinEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpu16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMINUW

	MOV Z2, ret+132(FP)
	RET

// func m512MinEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMINUW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMinEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMINUB

	MOVOU X3, ret+52(FP)
	RET

// func maskzMinEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMINUB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMinEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMinEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMINUB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMinEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMinEpu8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMINUB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMinEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPMINUB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzMinEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMinEpu8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPMINUB

	MOV Z2, ret+136(FP)
	RET

// func m512MinEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMINUB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMovEpi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskMovEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VMOVDQU16

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovEpi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzMovEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VMOVDQU16 R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovEpi16(src [32]byte, k uint16, a [32]byte) [32]byte
TEXT ·m256MaskMovEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2

	// TODO: Code missing - uses instrunction: VMOVDQU16

	MOV Y2, ret+68(FP)
	RET

// func m256MaskzMovEpi16(k uint16, a [32]byte) [32]byte
TEXT ·m256MaskzMovEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VMOVDQU16 R8, Y1

	MOV Y1, ret+36(FP)
	RET

// func m512MaskMovEpi16(src [64]byte, k uint32, a [64]byte) [64]byte
TEXT ·m512MaskMovEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2

	// TODO: Code missing - uses instrunction: VMOVDQU16

	MOV Z2, ret+132(FP)
	RET

// func m512MaskzMovEpi16(k uint32, a [64]byte) [64]byte
TEXT ·m512MaskzMovEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1

	// TODO: Code missing - could be:
	// VMOVDQU16 R8, Z1

	MOV Z1, ret+68(FP)
	RET

// func maskMovEpi8(src [16]byte, k uint16, a [16]byte) [16]byte
TEXT ·maskMovEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VMOVDQU8

	MOVOU X2, ret+36(FP)
	RET

// func maskzMovEpi8(k uint16, a [16]byte) [16]byte
TEXT ·maskzMovEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// VMOVDQU8 R8, X1

	MOVOU X1, ret+20(FP)
	RET

// func m256MaskMovEpi8(src [32]byte, k uint32, a [32]byte) [32]byte
TEXT ·m256MaskMovEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2

	// TODO: Code missing - uses instrunction: VMOVDQU8

	MOV Y2, ret+68(FP)
	RET

// func m256MaskzMovEpi8(k uint32, a [32]byte) [32]byte
TEXT ·m256MaskzMovEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1

	// TODO: Code missing - could be:
	// VMOVDQU8 R8, Y1

	MOV Y1, ret+36(FP)
	RET

// func m512MaskMovEpi8(src [64]byte, k uint64, a [64]byte) [64]byte
TEXT ·m512MaskMovEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2

	// TODO: Code missing - uses instrunction: VMOVDQU8

	MOV Z2, ret+136(FP)
	RET

// func m512MaskzMovEpi8(k uint64, a [64]byte) [64]byte
TEXT ·m512MaskzMovEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1

	// TODO: Code missing - could be:
	// VMOVDQU8 R8, Z1

	MOV Z1, ret+72(FP)
	RET

// func movepi16Mask(a [16]byte) uint8
TEXT ·movepi16Mask(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPMOVW2M X0

	MOVB $0, ret+16(FP)
	RET

// func m256Movepi16Mask(a [32]byte) uint16
TEXT ·m256Movepi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVW2M Y0

	MOVW $0, ret+32(FP)
	RET

// func m512Movepi16Mask(a [64]byte) uint32
TEXT ·m512Movepi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPMOVW2M Z0

	MOVL $0, ret+64(FP)
	RET

// func movepi8Mask(a [16]byte) uint16
TEXT ·movepi8Mask(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// VPMOVB2M X0

	MOVW $0, ret+16(FP)
	RET

// func m256Movepi8Mask(a [32]byte) uint32
TEXT ·m256Movepi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0

	// TODO: Code missing - could be:
	// VPMOVB2M Y0

	MOVL $0, ret+32(FP)
	RET

// func m512Movepi8Mask(a [64]byte) uint64
TEXT ·m512Movepi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0

	// TODO: Code missing - could be:
	// VPMOVB2M Z0

	MOVQ $0, ret+64(FP)
	RET

// func movmEpi16(k uint8) [16]byte
TEXT ·movmEpi16(SB),7,$0
	MOVB k+0(FP),R8

	// TODO: Code missing - could be:
	// VPMOVM2W R8

	MOVOU X0, ret+4(FP)
	RET

// func m256MovmEpi16(k uint16) [32]byte
TEXT ·m256MovmEpi16(SB),7,$0
	MOVW k+0(FP),R8

	// TODO: Code missing - could be:
	// VPMOVM2W R8

	MOV Y0, ret+4(FP)
	RET

// func m512MovmEpi16(k uint32) [64]byte
TEXT ·m512MovmEpi16(SB),7,$0
	MOVL k+0(FP),R8

	// TODO: Code missing - could be:
	// VPMOVM2W R8

	MOV Z0, ret+4(FP)
	RET

// func m256MovmEpi8(k uint32) [32]byte
TEXT ·m256MovmEpi8(SB),7,$0
	MOVL k+0(FP),R8

	// TODO: Code missing - could be:
	// VPMOVM2B R8

	MOV Y0, ret+4(FP)
	RET

// func m512MovmEpi8(k uint64) [64]byte
TEXT ·m512MovmEpi8(SB),7,$0
	MOVQ k+0(FP),R8

	// TODO: Code missing - could be:
	// VPMOVM2B R8

	MOV Z0, ret+8(FP)
	RET

// func maskMulhiEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulhiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMULHW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulhiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMULHW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulhiEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulhiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMULHW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMulhiEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulhiEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMULHW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMulhiEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMULHW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMulhiEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulhiEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMULHW

	MOV Z2, ret+132(FP)
	RET

// func m512MulhiEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMULHW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMulhiEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulhiEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMULHUW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulhiEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulhiEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMULHUW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulhiEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulhiEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMULHUW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMulhiEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulhiEpu16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMULHUW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMulhiEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhiEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMULHUW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMulhiEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulhiEpu16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMULHUW

	MOV Z2, ret+132(FP)
	RET

// func m512MulhiEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhiEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMULHUW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMulhrsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulhrsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMULHRSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulhrsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulhrsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMULHRSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulhrsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulhrsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMULHRSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMulhrsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulhrsEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMULHRSW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMulhrsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhrsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMULHRSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMulhrsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulhrsEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMULHRSW

	MOV Z2, ret+132(FP)
	RET

// func m512MulhrsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhrsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMULHRSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskMulloEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulloEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPMULLW

	MOVOU X3, ret+52(FP)
	RET

// func maskzMulloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulloEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPMULLW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskMulloEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskMulloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPMULLW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzMulloEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzMulloEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPMULLW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskMulloEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMULLW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMulloEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMulloEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPMULLW

	MOV Z2, ret+132(FP)
	RET

// func m512MulloEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMULLW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskPacksEpi16(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPacksEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPACKSSWB

	MOVOU X3, ret+52(FP)
	RET

// func maskzPacksEpi16(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPacksEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPACKSSWB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskPacksEpi16(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPacksEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPACKSSWB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzPacksEpi16(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPacksEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPACKSSWB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskPacksEpi16(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPacksEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPACKSSWB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzPacksEpi16(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPacksEpi16(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPACKSSWB

	MOV Z2, ret+136(FP)
	RET

// func m512PacksEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PacksEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPACKSSWB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskPacksEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPacksEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPACKSSDW

	MOVOU X3, ret+52(FP)
	RET

// func maskzPacksEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPacksEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPACKSSDW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskPacksEpi32(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPacksEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPACKSSDW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzPacksEpi32(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPacksEpi32(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPACKSSDW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskPacksEpi32(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPacksEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPACKSSDW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzPacksEpi32(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPacksEpi32(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPACKSSDW

	MOV Z2, ret+132(FP)
	RET

// func m512PacksEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PacksEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPACKSSDW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskPackusEpi16(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPackusEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPACKUSWB

	MOVOU X3, ret+52(FP)
	RET

// func maskzPackusEpi16(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPackusEpi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPACKUSWB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskPackusEpi16(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPackusEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPACKUSWB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzPackusEpi16(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPackusEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPACKUSWB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskPackusEpi16(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPackusEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPACKUSWB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzPackusEpi16(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPackusEpi16(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPACKUSWB

	MOV Z2, ret+136(FP)
	RET

// func m512PackusEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PackusEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPACKUSWB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskPackusEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskPackusEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPACKUSDW

	MOVOU X3, ret+52(FP)
	RET

// func maskzPackusEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzPackusEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPACKUSDW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskPackusEpi32(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPackusEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPACKUSDW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzPackusEpi32(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPackusEpi32(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPACKUSDW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskPackusEpi32(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPackusEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPACKUSDW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzPackusEpi32(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPackusEpi32(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPACKUSDW

	MOV Z2, ret+132(FP)
	RET

// func m512PackusEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512PackusEpi32(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPACKUSDW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskPermutex2varEpi16(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMT2W

	MOVOU X3, ret+52(FP)
	RET

// func mask2Permutex2varEpi16(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMI2W

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutex2varEpi16(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMI2W, VPERMT2W

	MOVOU X3, ret+52(FP)
	RET

// func permutex2varEpi16(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing - uses instrunction: VPERMI2W, VPERMT2W

	MOVOU X2, ret+48(FP)
	RET

// func m256MaskPermutex2varEpi16(a [32]byte, k uint16, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskPermutex2varEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPERMT2W

	MOV Y3, ret+100(FP)
	RET

// func m256Mask2Permutex2varEpi16(a [32]byte, idx [32]byte, k uint16, b [32]byte) [32]byte
TEXT ·m256Mask2Permutex2varEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+32(FP),Y1
	MOVW k+64(FP),R10
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPERMI2W

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzPermutex2varEpi16(k uint16, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzPermutex2varEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPERMI2W, VPERMT2W

	MOV Y3, ret+100(FP)
	RET

// func m256Permutex2varEpi16(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·m256Permutex2varEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+32(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+64(FP),Y2

	// TODO: Code missing - uses instrunction: VPERMI2W, VPERMT2W

	MOV Y2, ret+96(FP)
	RET

// func m512MaskPermutex2varEpi16(a [64]byte, k uint32, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPermutex2varEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMT2W

	MOV Z3, ret+196(FP)
	RET

// func m512Mask2Permutex2varEpi16(a [64]byte, idx [64]byte, k uint32, b [64]byte) [64]byte
TEXT ·m512Mask2Permutex2varEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+64(FP),Z1
	MOVL k+128(FP),R10
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMI2W

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzPermutex2varEpi16(k uint32, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPermutex2varEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMI2W, VPERMT2W

	MOV Z3, ret+196(FP)
	RET

// func m512Permutex2varEpi16(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512Permutex2varEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+64(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+128(FP),Z2

	// TODO: Code missing - uses instrunction: VPERMI2W, VPERMT2W

	MOV Z2, ret+192(FP)
	RET

// func maskPermutexvarEpi16(src [16]byte, k uint8, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskPermutexvarEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU a+36(FP),X3

	// TODO: Code missing - uses instrunction: VPERMW

	MOVOU X3, ret+52(FP)
	RET

// func maskzPermutexvarEpi16(k uint8, idx [16]byte, a [16]byte) [16]byte
TEXT ·maskzPermutexvarEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU idx+4(FP),X1
	MOVOU a+20(FP),X2

	// TODO: Code missing - uses instrunction: VPERMW

	MOVOU X2, ret+36(FP)
	RET

// func permutexvarEpi16(idx [16]byte, a [16]byte) [16]byte
TEXT ·permutexvarEpi16(SB),7,$0
	MOVOU idx+0(FP),X0
	MOVOU a+16(FP),X1

	// TODO: Code missing - could be:
	// VPERMW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskPermutexvarEpi16(src [32]byte, k uint16, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskPermutexvarEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPERMW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzPermutexvarEpi16(k uint16, idx [32]byte, a [32]byte) [32]byte
TEXT ·m256MaskzPermutexvarEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPERMW

	MOV Y2, ret+68(FP)
	RET

// func m256PermutexvarEpi16(idx [32]byte, a [32]byte) [32]byte
TEXT ·m256PermutexvarEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV idx+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+32(FP),Y1

	// TODO: Code missing - could be:
	// VPERMW Y0, Y1

	MOV Y1, ret+64(FP)
	RET

// func m512MaskPermutexvarEpi16(src [64]byte, k uint32, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutexvarEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzPermutexvarEpi16(k uint32, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskzPermutexvarEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPERMW

	MOV Z2, ret+132(FP)
	RET

// func m512PermutexvarEpi16(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutexvarEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+64(FP),Z1

	// TODO: Code missing - could be:
	// VPERMW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512SadEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SadEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSADBW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSet1Epi16(src [16]byte, k uint8, a int16) [16]byte
TEXT ·maskSet1Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVW a+20(FP),R10

	// TODO: Code missing - uses instrunction: VPBROADCASTW

	MOVOU X2, ret+24(FP)
	RET

// func maskzSet1Epi16(k uint8, a int16) [16]byte
TEXT ·maskzSet1Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVW a+4(FP),R9

	// TODO: Code missing - could be:
	// VPBROADCASTW R8, R9

	MOVOU X1, ret+8(FP)
	RET

// func m256MaskSet1Epi16(src [32]byte, k uint16, a int16) [32]byte
TEXT ·m256MaskSet1Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	MOVW a+36(FP),R10

	// TODO: Code missing - uses instrunction: VPBROADCASTW

	MOV Y2, ret+40(FP)
	RET

// func m256MaskzSet1Epi16(k uint16, a int16) [32]byte
TEXT ·m256MaskzSet1Epi16(SB),7,$0
	MOVW k+0(FP),R8
	MOVW a+4(FP),R9

	// TODO: Code missing - could be:
	// VPBROADCASTW R8, R9

	MOV Y1, ret+8(FP)
	RET

// func m512MaskSet1Epi16(src [64]byte, k uint32, a int16) [64]byte
TEXT ·m512MaskSet1Epi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	MOVW a+68(FP),R10

	// TODO: Code missing - uses instrunction: VPBROADCASTW

	MOV Z2, ret+72(FP)
	RET

// func m512MaskzSet1Epi16(k uint32, a int16) [64]byte
TEXT ·m512MaskzSet1Epi16(SB),7,$0
	MOVL k+0(FP),R8
	MOVW a+4(FP),R9

	// TODO: Code missing - could be:
	// VPBROADCASTW R8, R9

	MOV Z1, ret+8(FP)
	RET

// func maskSet1Epi8(src [16]byte, k uint16, a byte) [16]byte
TEXT ·maskSet1Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVB a+20(FP),R10

	// TODO: Code missing - uses instrunction: VPBROADCASTB

	MOVOU X2, ret+24(FP)
	RET

// func maskzSet1Epi8(k uint16, a byte) [16]byte
TEXT ·maskzSet1Epi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVB a+4(FP),R9

	// TODO: Code missing - could be:
	// VPBROADCASTB R8, R9

	MOVOU X1, ret+8(FP)
	RET

// func m256MaskSet1Epi8(src [32]byte, k uint32, a byte) [32]byte
TEXT ·m256MaskSet1Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	MOVB a+36(FP),R10

	// TODO: Code missing - uses instrunction: VPBROADCASTB

	MOV Y2, ret+40(FP)
	RET

// func m256MaskzSet1Epi8(k uint32, a byte) [32]byte
TEXT ·m256MaskzSet1Epi8(SB),7,$0
	MOVL k+0(FP),R8
	MOVB a+4(FP),R9

	// TODO: Code missing - could be:
	// VPBROADCASTB R8, R9

	MOV Y1, ret+8(FP)
	RET

// func m512MaskSet1Epi8(src [64]byte, k uint64, a byte) [64]byte
TEXT ·m512MaskSet1Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	MOVB a+72(FP),R10

	// TODO: Code missing - uses instrunction: VPBROADCASTB

	MOV Z2, ret+76(FP)
	RET

// func m512MaskzSet1Epi8(k uint64, a byte) [64]byte
TEXT ·m512MaskzSet1Epi8(SB),7,$0
	MOVQ k+0(FP),R8
	MOVB a+8(FP),R9

	// TODO: Code missing - could be:
	// VPBROADCASTB R8, R9

	MOV Z1, ret+12(FP)
	RET

// func maskShuffleEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskShuffleEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSHUFB

	MOVOU X3, ret+52(FP)
	RET

// func maskzShuffleEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzShuffleEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSHUFB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskShuffleEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskShuffleEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSHUFB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzShuffleEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzShuffleEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSHUFB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskShuffleEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskShuffleEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPSHUFB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzShuffleEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzShuffleEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPSHUFB

	MOV Z2, ret+136(FP)
	RET

// func m512ShuffleEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512ShuffleEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSHUFB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskShufflehiEpi16(src [16]byte, k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskShufflehiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFHW

	MOVOU X3, ret+40(FP)
	RET

// func maskzShufflehiEpi16(k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskzShufflehiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFHW

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskShufflehiEpi16(src [32]byte, k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskShufflehiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFHW

	MOV Y3, ret+72(FP)
	RET

// func m256MaskzShufflehiEpi16(k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskzShufflehiEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFHW

	MOV Y2, ret+40(FP)
	RET

// func m512MaskShufflehiEpi16(src [64]byte, k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskShufflehiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFHW

	MOV Z3, ret+136(FP)
	RET

// func m512MaskzShufflehiEpi16(k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskzShufflehiEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFHW

	MOV Z2, ret+72(FP)
	RET

// func m512ShufflehiEpi16(a [64]byte, imm8 byte) [64]byte
TEXT ·m512ShufflehiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSHUFHW Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func maskShuffleloEpi16(src [16]byte, k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskShuffleloEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFLW

	MOVOU X3, ret+40(FP)
	RET

// func maskzShuffleloEpi16(k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskzShuffleloEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFLW

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskShuffleloEpi16(src [32]byte, k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskShuffleloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFLW

	MOV Y3, ret+72(FP)
	RET

// func m256MaskzShuffleloEpi16(k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskzShuffleloEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFLW

	MOV Y2, ret+40(FP)
	RET

// func m512MaskShuffleloEpi16(src [64]byte, k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskShuffleloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFLW

	MOV Z3, ret+136(FP)
	RET

// func m512MaskzShuffleloEpi16(k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskzShuffleloEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSHUFLW

	MOV Z2, ret+72(FP)
	RET

// func m512ShuffleloEpi16(a [64]byte, imm8 byte) [64]byte
TEXT ·m512ShuffleloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSHUFLW Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func maskSllEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSLLW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSllEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSLLW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSllEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSllEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	MOVOU count+68(FP),X3

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Y3, ret+84(FP)
	RET

// func m256MaskzSllEpi16(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSllEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	MOVOU count+36(FP),X2

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Y2, ret+52(FP)
	RET

// func m512MaskSllEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSllEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	MOVOU count+132(FP),X3

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Z3, ret+148(FP)
	RET

// func m512MaskzSllEpi16(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSllEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	MOVOU count+68(FP),X2

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Z2, ret+84(FP)
	RET

// func m512SllEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SllEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVOU count+64(FP),X1

	// TODO: Code missing - could be:
	// VPSLLW Z0, X1

	MOV Z1, ret+80(FP)
	RET

// func maskSlliEpi16(src [16]byte, k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskSlliEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSLLW

	MOVOU X3, ret+40(FP)
	RET

// func maskzSlliEpi16(k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskzSlliEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSLLW

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSlliEpi16(src [32]byte, k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskSlliEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Y3, ret+72(FP)
	RET

// func m256MaskzSlliEpi16(k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskzSlliEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Y2, ret+40(FP)
	RET

// func m512MaskSlliEpi16(src [64]byte, k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskSlliEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Z3, ret+136(FP)
	RET

// func m512MaskzSlliEpi16(k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskzSlliEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSLLW

	MOV Z2, ret+72(FP)
	RET

// func m512SlliEpi16(a [64]byte, imm8 byte) [64]byte
TEXT ·m512SlliEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSLLW Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func maskSllvEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllvEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSLLVW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSllvEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllvEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSLLVW

	MOVOU X2, ret+36(FP)
	RET

// func sllvEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sllvEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	// TODO: Code missing - could be:
	// VPSLLVW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskSllvEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSllvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSLLVW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSllvEpi16(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSllvEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSLLVW

	MOV Y2, ret+68(FP)
	RET

// func m256SllvEpi16(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SllvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+32(FP),Y1

	// TODO: Code missing - could be:
	// VPSLLVW Y0, Y1

	MOV Y1, ret+64(FP)
	RET

// func m512MaskSllvEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSllvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPSLLVW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzSllvEpi16(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSllvEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPSLLVW

	MOV Z2, ret+132(FP)
	RET

// func m512SllvEpi16(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SllvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSLLVW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSraEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSraEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSRAW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSraEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSraEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSRAW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSraEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSraEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	MOVOU count+68(FP),X3

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Y3, ret+84(FP)
	RET

// func m256MaskzSraEpi16(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSraEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	MOVOU count+36(FP),X2

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Y2, ret+52(FP)
	RET

// func m512MaskSraEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSraEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	MOVOU count+132(FP),X3

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Z3, ret+148(FP)
	RET

// func m512MaskzSraEpi16(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSraEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	MOVOU count+68(FP),X2

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Z2, ret+84(FP)
	RET

// func m512SraEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SraEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVOU count+64(FP),X1

	// TODO: Code missing - could be:
	// VPSRAW Z0, X1

	MOV Z1, ret+80(FP)
	RET

// func maskSraiEpi16(src [16]byte, k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskSraiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRAW

	MOVOU X3, ret+40(FP)
	RET

// func maskzSraiEpi16(k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskzSraiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRAW

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSraiEpi16(src [32]byte, k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskSraiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Y3, ret+72(FP)
	RET

// func m256MaskzSraiEpi16(k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskzSraiEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Y2, ret+40(FP)
	RET

// func m512MaskSraiEpi16(src [64]byte, k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskSraiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Z3, ret+136(FP)
	RET

// func m512MaskzSraiEpi16(k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskzSraiEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRAW

	MOV Z2, ret+72(FP)
	RET

// func m512SraiEpi16(a [64]byte, imm8 byte) [64]byte
TEXT ·m512SraiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSRAW Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func maskSravEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSravEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSRAVW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSravEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSravEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSRAVW

	MOVOU X2, ret+36(FP)
	RET

// func sravEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sravEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	// TODO: Code missing - could be:
	// VPSRAVW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskSravEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSravEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSRAVW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSravEpi16(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSravEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSRAVW

	MOV Y2, ret+68(FP)
	RET

// func m256SravEpi16(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SravEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+32(FP),Y1

	// TODO: Code missing - could be:
	// VPSRAVW Y0, Y1

	MOV Y1, ret+64(FP)
	RET

// func m512MaskSravEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSravEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPSRAVW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzSravEpi16(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSravEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPSRAVW

	MOV Z2, ret+132(FP)
	RET

// func m512SravEpi16(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SravEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSRAVW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSrlEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSRLW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSrlEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSRLW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSrlEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskSrlEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	MOVOU count+68(FP),X3

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Y3, ret+84(FP)
	RET

// func m256MaskzSrlEpi16(k uint16, a [32]byte, count [16]byte) [32]byte
TEXT ·m256MaskzSrlEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	MOVOU count+36(FP),X2

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Y2, ret+52(FP)
	RET

// func m512MaskSrlEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskSrlEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	MOVOU count+132(FP),X3

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Z3, ret+148(FP)
	RET

// func m512MaskzSrlEpi16(k uint32, a [64]byte, count [16]byte) [64]byte
TEXT ·m512MaskzSrlEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	MOVOU count+68(FP),X2

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Z2, ret+84(FP)
	RET

// func m512SrlEpi16(a [64]byte, count [16]byte) [64]byte
TEXT ·m512SrlEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVOU count+64(FP),X1

	// TODO: Code missing - could be:
	// VPSRLW Z0, X1

	MOV Z1, ret+80(FP)
	RET

// func maskSrliEpi16(src [16]byte, k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskSrliEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRLW

	MOVOU X3, ret+40(FP)
	RET

// func maskzSrliEpi16(k uint8, a [16]byte, imm8 byte) [16]byte
TEXT ·maskzSrliEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRLW

	MOVOU X2, ret+24(FP)
	RET

// func m256MaskSrliEpi16(src [32]byte, k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskSrliEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Y3, ret+72(FP)
	RET

// func m256MaskzSrliEpi16(k uint16, a [32]byte, imm8 byte) [32]byte
TEXT ·m256MaskzSrliEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Y2, ret+40(FP)
	RET

// func m512MaskSrliEpi16(src [64]byte, k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskSrliEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Z3, ret+136(FP)
	RET

// func m512MaskzSrliEpi16(k uint32, a [64]byte, imm8 byte) [64]byte
TEXT ·m512MaskzSrliEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - uses instrunction: VPSRLW

	MOV Z2, ret+72(FP)
	RET

// func m512SrliEpi16(a [64]byte, imm8 byte) [64]byte
TEXT ·m512SrliEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Immediate parameter should be removed (imm8 byte)

	// TODO: Code missing - could be:
	// VPSRLW Z0, R9

	MOV Z1, ret+68(FP)
	RET

// func maskSrlvEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlvEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSRLVW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSrlvEpi16(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlvEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSRLVW

	MOVOU X2, ret+36(FP)
	RET

// func srlvEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·srlvEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	// TODO: Code missing - could be:
	// VPSRLVW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func m256MaskSrlvEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskSrlvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSRLVW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSrlvEpi16(k uint16, a [32]byte, count [32]byte) [32]byte
TEXT ·m256MaskzSrlvEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSRLVW

	MOV Y2, ret+68(FP)
	RET

// func m256SrlvEpi16(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SrlvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV count+32(FP),Y1

	// TODO: Code missing - could be:
	// VPSRLVW Y0, Y1

	MOV Y1, ret+64(FP)
	RET

// func m512MaskSrlvEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSrlvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPSRLVW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzSrlvEpi16(k uint32, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskzSrlvEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPSRLVW

	MOV Z2, ret+132(FP)
	RET

// func m512SrlvEpi16(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SrlvEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV count+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSRLVW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSubEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSUBW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSUBW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSUBW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSubEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSUBW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskSubEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPSUBW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzSubEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPSUBW

	MOV Z2, ret+132(FP)
	RET

// func m512SubEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSUBW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSubEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSUBB

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSUBB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSUBB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSubEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSUBB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskSubEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPSUBB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzSubEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPSUBB

	MOV Z2, ret+136(FP)
	RET

// func m512SubEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSUBB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSubsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSUBSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSUBSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSUBSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSubsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSUBSW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskSubsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPSUBSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzSubsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPSUBSW

	MOV Z2, ret+132(FP)
	RET

// func m512SubsEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSUBSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSubsEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSUBSB

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSUBSB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubsEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSUBSB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSubsEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSUBSB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskSubsEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPSUBSB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzSubsEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPSUBSB

	MOV Z2, ret+136(FP)
	RET

// func m512SubsEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSUBSB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSubsEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSUBUSW

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSUBUSW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubsEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSUBUSW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSubsEpu16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpu16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSUBUSW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskSubsEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPSUBUSW

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzSubsEpu16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpu16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPSUBUSW

	MOV Z2, ret+132(FP)
	RET

// func m512SubsEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpu16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSUBUSW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskSubsEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubsEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPSUBUSB

	MOVOU X3, ret+52(FP)
	RET

// func maskzSubsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubsEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPSUBUSB

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskSubsEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskSubsEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPSUBUSB

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzSubsEpu8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzSubsEpu8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPSUBUSB

	MOV Y2, ret+68(FP)
	RET

// func m512MaskSubsEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubsEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPSUBUSB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzSubsEpu8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzSubsEpu8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPSUBUSB

	MOV Z2, ret+136(FP)
	RET

// func m512SubsEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubsEpu8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPSUBUSB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskTestEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPTESTMW

	MOVB $0, ret+36(FP)
	RET

// func testEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·testEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPTESTMW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskTestEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPTESTMW

	MOVW $0, ret+68(FP)
	RET

// func m256TestEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256TestEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPTESTMW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m512MaskTestEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskTestEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPTESTMW

	MOVL $0, ret+132(FP)
	RET

// func m512TestEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512TestEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPTESTMW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func maskTestEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskTestEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPTESTMB

	MOVW $0, ret+36(FP)
	RET

// func testEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·testEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPTESTMB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func m256MaskTestEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskTestEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPTESTMB

	MOVL $0, ret+68(FP)
	RET

// func m256TestEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256TestEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPTESTMB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m512MaskTestEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskTestEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPTESTMB

	MOVQ $0, ret+136(FP)
	RET

// func m512TestEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512TestEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPTESTMB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func maskTestnEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestnEpi16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPTESTNMW

	MOVB $0, ret+36(FP)
	RET

// func testnEpi16Mask(a [16]byte, b [16]byte) uint8
TEXT ·testnEpi16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPTESTNMW X0, X1

	MOVB $0, ret+32(FP)
	RET

// func m256MaskTestnEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16
TEXT ·m256MaskTestnEpi16Mask(SB),7,$0
	MOVW k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPTESTNMW

	MOVW $0, ret+68(FP)
	RET

// func m256TestnEpi16Mask(a [32]byte, b [32]byte) uint16
TEXT ·m256TestnEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPTESTNMW Y0, Y1

	MOVW $0, ret+64(FP)
	RET

// func m512MaskTestnEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32
TEXT ·m512MaskTestnEpi16Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPTESTNMW

	MOVL $0, ret+132(FP)
	RET

// func m512TestnEpi16Mask(a [64]byte, b [64]byte) uint32
TEXT ·m512TestnEpi16Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPTESTNMW Z0, Z1

	MOVL $0, ret+128(FP)
	RET

// func maskTestnEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16
TEXT ·maskTestnEpi8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPTESTNMB

	MOVW $0, ret+36(FP)
	RET

// func testnEpi8Mask(a [16]byte, b [16]byte) uint16
TEXT ·testnEpi8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// VPTESTNMB X0, X1

	MOVW $0, ret+32(FP)
	RET

// func m256MaskTestnEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32
TEXT ·m256MaskTestnEpi8Mask(SB),7,$0
	MOVL k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPTESTNMB

	MOVL $0, ret+68(FP)
	RET

// func m256TestnEpi8Mask(a [32]byte, b [32]byte) uint32
TEXT ·m256TestnEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+0(FP),Y0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+32(FP),Y1

	// TODO: Code missing - could be:
	// VPTESTNMB Y0, Y1

	MOVL $0, ret+64(FP)
	RET

// func m512MaskTestnEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64
TEXT ·m512MaskTestnEpi8Mask(SB),7,$0
	MOVQ k1+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPTESTNMB

	MOVQ $0, ret+136(FP)
	RET

// func m512TestnEpi8Mask(a [64]byte, b [64]byte) uint64
TEXT ·m512TestnEpi8Mask(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPTESTNMB Z0, Z1

	MOVQ $0, ret+128(FP)
	RET

// func maskUnpackhiEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPUNPCKHWD

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPUNPCKHWD

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackhiEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackhiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPUNPCKHWD

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzUnpackhiEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackhiEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPUNPCKHWD

	MOV Y2, ret+68(FP)
	RET

// func m512MaskUnpackhiEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackhiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPUNPCKHWD

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzUnpackhiEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackhiEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPUNPCKHWD

	MOV Z2, ret+132(FP)
	RET

// func m512UnpackhiEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackhiEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPUNPCKHWD Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskUnpackhiEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPUNPCKHBW

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackhiEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPUNPCKHBW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackhiEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackhiEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPUNPCKHBW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzUnpackhiEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackhiEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPUNPCKHBW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskUnpackhiEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackhiEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPUNPCKHBW

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzUnpackhiEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackhiEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPUNPCKHBW

	MOV Z2, ret+136(FP)
	RET

// func m512UnpackhiEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackhiEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPUNPCKHBW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskUnpackloEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPUNPCKLWD

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPUNPCKLWD

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackloEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVW k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPUNPCKLWD

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzUnpackloEpi16(k uint16, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackloEpi16(SB),7,$0
	MOVW k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPUNPCKLWD

	MOV Y2, ret+68(FP)
	RET

// func m512MaskUnpackloEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVL k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPUNPCKLWD

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzUnpackloEpi16(k uint32, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackloEpi16(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2

	// TODO: Code missing - uses instrunction: VPUNPCKLWD

	MOV Z2, ret+132(FP)
	RET

// func m512UnpackloEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackloEpi16(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPUNPCKLWD Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func maskUnpackloEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	// TODO: Code missing - uses instrunction: VPUNPCKLBW

	MOVOU X3, ret+52(FP)
	RET

// func maskzUnpackloEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	// TODO: Code missing - uses instrunction: VPUNPCKLBW

	MOVOU X2, ret+36(FP)
	RET

// func m256MaskUnpackloEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskUnpackloEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV src+0(FP),Y0
	MOVL k+32(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+36(FP),Y2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+68(FP),Y3

	// TODO: Code missing - uses instrunction: VPUNPCKLBW

	MOV Y3, ret+100(FP)
	RET

// func m256MaskzUnpackloEpi8(k uint32, a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaskzUnpackloEpi8(SB),7,$0
	MOVL k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV a+4(FP),Y1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M256i
	//	MOV b+36(FP),Y2

	// TODO: Code missing - uses instrunction: VPUNPCKLBW

	MOV Y2, ret+68(FP)
	RET

// func m512MaskUnpackloEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskUnpackloEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPUNPCKLBW

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzUnpackloEpi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzUnpackloEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPUNPCKLBW

	MOV Z2, ret+136(FP)
	RET

// func m512UnpackloEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512UnpackloEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPUNPCKLBW Z0, Z1

	MOV Z1, ret+128(FP)
	RET

