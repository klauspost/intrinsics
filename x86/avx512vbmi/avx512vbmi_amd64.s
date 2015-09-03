// func m512MaskMultishiftEpi64Epi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMultishiftEpi64Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPMULTISHIFTQB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzMultishiftEpi64Epi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMultishiftEpi64Epi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPMULTISHIFTQB

	MOV Z2, ret+136(FP)
	RET

// func m512MultishiftEpi64Epi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MultishiftEpi64Epi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+64(FP),Z1

	// TODO: Code missing - could be:
	// VPMULTISHIFTQB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

// func m512MaskPermutex2varEpi8(a [64]byte, k uint64, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPermutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMT2B

	MOV Z3, ret+200(FP)
	RET

// func m512Mask2Permutex2varEpi8(a [64]byte, idx [64]byte, k uint64, b [64]byte) [64]byte
TEXT ·m512Mask2Permutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+64(FP),Z1
	MOVQ k+128(FP),R10
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMI2B

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzPermutex2varEpi8(k uint64, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPermutex2varEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMI2B, VPERMT2B

	MOV Z3, ret+200(FP)
	RET

// func m512Permutex2varEpi8(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512Permutex2varEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+64(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV b+128(FP),Z2

	// TODO: Code missing - uses instrunction: VPERMI2B

	MOV Z2, ret+192(FP)
	RET

// func m512MaskPermutexvarEpi8(src [64]byte, k uint64, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutexvarEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV src+0(FP),Z0
	MOVQ k+64(FP),R9
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+72(FP),Z2
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+136(FP),Z3

	// TODO: Code missing - uses instrunction: VPERMB

	MOV Z3, ret+200(FP)
	RET

// func m512MaskzPermutexvarEpi8(k uint64, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskzPermutexvarEpi8(SB),7,$0
	MOVQ k+0(FP),R8
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+8(FP),Z1
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+72(FP),Z2

	// TODO: Code missing - uses instrunction: VPERMB

	MOV Z2, ret+136(FP)
	RET

// func m512PermutexvarEpi8(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutexvarEpi8(SB),7,$0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV idx+0(FP),Z0
	// FIXME: Unimplemented. Unknown MOVE postfix for type x86.M512i
	//	MOV a+64(FP),Z1

	// TODO: Code missing - could be:
	// VPERMB Z0, Z1

	MOV Z1, ret+128(FP)
	RET

