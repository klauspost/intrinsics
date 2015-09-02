// func m512MaskMultishiftEpi64Epi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMultishiftEpi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzMultishiftEpi64Epi8(k uint64, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzMultishiftEpi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MultishiftEpi64Epi8(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MultishiftEpi64Epi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing - could be:
	// VPMULTISHIFTQB Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskPermutex2varEpi8(a [64]byte, k uint64, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskPermutex2varEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask2Permutex2varEpi8(a [64]byte, idx [64]byte, k uint64, b [64]byte) [64]byte
TEXT ·m512Mask2Permutex2varEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutex2varEpi8(k uint64, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskzPermutex2varEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permutex2varEpi8(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·m512Permutex2varEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskPermutexvarEpi8(src [64]byte, k uint64, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutexvarEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskzPermutexvarEpi8(k uint64, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskzPermutexvarEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512PermutexvarEpi8(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutexvarEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing - could be:
	// VPERMB Z0, Z1

	MOV Z1, ret+0(FP)
	RET

