// func broadcastmbEpi64(k uint8) [16]byte
TEXT ·broadcastmbEpi64(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastmbEpi641(k uint8) [32]byte
TEXT ·broadcastmbEpi641(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func broadcastmbEpi642(k uint8) [64]byte
TEXT ·broadcastmbEpi642(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func broadcastmwEpi32(k uint16) [16]byte
TEXT ·broadcastmwEpi32(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastmwEpi321(k uint16) [32]byte
TEXT ·broadcastmwEpi321(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOV Y0, ret+4(FP)
	RET

// func broadcastmwEpi322(k uint16) [64]byte
TEXT ·broadcastmwEpi322(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func conflictEpi32(a [16]byte) [16]byte
TEXT ·conflictEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskConflictEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskConflictEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzConflictEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzConflictEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func conflictEpi321(a [32]byte) [32]byte
TEXT ·conflictEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskConflictEpi321(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskConflictEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzConflictEpi321(k uint8, a [32]byte) [32]byte
TEXT ·maskzConflictEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func conflictEpi322(a [64]byte) [64]byte
TEXT ·conflictEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskConflictEpi322(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·maskConflictEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzConflictEpi322(k uint16, a [64]byte) [64]byte
TEXT ·maskzConflictEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func conflictEpi64(a [16]byte) [16]byte
TEXT ·conflictEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskConflictEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskConflictEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzConflictEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzConflictEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func conflictEpi641(a [32]byte) [32]byte
TEXT ·conflictEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskConflictEpi641(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskConflictEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzConflictEpi641(k uint8, a [32]byte) [32]byte
TEXT ·maskzConflictEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func conflictEpi642(a [64]byte) [64]byte
TEXT ·conflictEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskConflictEpi642(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·maskConflictEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzConflictEpi642(k uint8, a [64]byte) [64]byte
TEXT ·maskzConflictEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func lzcntEpi32(a [16]byte) [16]byte
TEXT ·lzcntEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskLzcntEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskLzcntEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzLzcntEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzLzcntEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func lzcntEpi321(a [32]byte) [32]byte
TEXT ·lzcntEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskLzcntEpi321(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskLzcntEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLzcntEpi321(k uint8, a [32]byte) [32]byte
TEXT ·maskzLzcntEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func lzcntEpi322(a [64]byte) [64]byte
TEXT ·lzcntEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLzcntEpi322(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·maskLzcntEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLzcntEpi322(k uint16, a [64]byte) [64]byte
TEXT ·maskzLzcntEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func lzcntEpi64(a [16]byte) [16]byte
TEXT ·lzcntEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskLzcntEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskLzcntEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzLzcntEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzLzcntEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func lzcntEpi641(a [32]byte) [32]byte
TEXT ·lzcntEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskLzcntEpi641(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskLzcntEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLzcntEpi641(k uint8, a [32]byte) [32]byte
TEXT ·maskzLzcntEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func lzcntEpi642(a [64]byte) [64]byte
TEXT ·lzcntEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLzcntEpi642(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·maskLzcntEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLzcntEpi642(k uint8, a [64]byte) [64]byte
TEXT ·maskzLzcntEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

