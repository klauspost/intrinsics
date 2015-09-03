// func m512Madd52hiEpu64(a [64]byte, b [64]byte, c [64]byte) [64]byte
TEXT ·m512Madd52hiEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV c+128(FP),Z2

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOV Z2, ret+192(FP)
	RET

// func m512MaskMadd52hiEpu64(a [64]byte, k uint8, b [64]byte, c [64]byte) [64]byte
TEXT ·m512MaskMadd52hiEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVB k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV c+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMadd52hiEpu64(k uint8, a [64]byte, b [64]byte, c [64]byte) [64]byte
TEXT ·m512MaskzMadd52hiEpu64(SB),7,$0
	MOVB k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV c+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMADD52HUQ

	MOV Z3, ret+196(FP)
	RET

// func m512Madd52loEpu64(a [64]byte, b [64]byte, c [64]byte) [64]byte
TEXT ·m512Madd52loEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV c+128(FP),Z2

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOV Z2, ret+192(FP)
	RET

// func m512MaskMadd52loEpu64(a [64]byte, k uint8, b [64]byte, c [64]byte) [64]byte
TEXT ·m512MaskMadd52loEpu64(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVB k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV c+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOV Z3, ret+196(FP)
	RET

// func m512MaskzMadd52loEpu64(k uint8, a [64]byte, b [64]byte, c [64]byte) [64]byte
TEXT ·m512MaskzMadd52loEpu64(SB),7,$0
	MOVB k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+4(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+68(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV c+132(FP),Z3

	// TODO: Code missing - uses instrunction: VPMADD52LUQ

	MOV Z3, ret+196(FP)
	RET

