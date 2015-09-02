// func maskAbsEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskAbsEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbsEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzAbsEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskAbsEpi321(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskAbsEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAbsEpi321(k uint8, a [32]byte) [32]byte
TEXT ·maskzAbsEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func absEpi32(a [64]byte) [64]byte
TEXT ·absEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAbsEpi322(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·maskAbsEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAbsEpi322(k uint16, a [64]byte) [64]byte
TEXT ·maskzAbsEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func absEpi64(a [16]byte) [16]byte
TEXT ·absEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskAbsEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskAbsEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbsEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzAbsEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func absEpi641(a [32]byte) [32]byte
TEXT ·absEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskAbsEpi641(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskAbsEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAbsEpi641(k uint8, a [32]byte) [32]byte
TEXT ·maskzAbsEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func absEpi642(a [64]byte) [64]byte
TEXT ·absEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAbsEpi642(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·maskAbsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAbsEpi642(k uint8, a [64]byte) [64]byte
TEXT ·maskzAbsEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func acosPd(a [8]float64) [8]float64
TEXT ·acosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAcosPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskAcosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func acosPs(a [16]float32) [16]float32
TEXT ·acosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAcosPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskAcosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func acoshPd(a [8]float64) [8]float64
TEXT ·acoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAcoshPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskAcoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func acoshPs(a [16]float32) [16]float32
TEXT ·acoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAcoshPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskAcoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAddEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAddEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAddEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAddEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAddEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAddEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAddEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·addEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAddEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskAddEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAddEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzAddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzAddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskzAddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAddRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskzAddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func addRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·addRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAddRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskAddRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAddRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzAddRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func addRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·addRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAddRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskAddRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAddRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzAddRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskAddSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAddSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAddSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAddSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAddSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAddSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAlignrEpi32(k uint16, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·maskzAlignrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func alignrEpi64(a [64]byte, b [64]byte, count int) [64]byte
TEXT ·alignrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAlignrEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·maskAlignrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzAlignrEpi64(k uint8, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·maskzAlignrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAndEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAndEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAndEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAndEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAndEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAndEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndnotEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndnotEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnotEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndnotEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndnotEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAndnotEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndnotEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAndnotEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndnotEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAndnotEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAndnotEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskAndnotEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnotEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzAndnotEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndnotEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskAndnotEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndnotEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzAndnotEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzAndnotEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzAndnotEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func asinPd(a [8]float64) [8]float64
TEXT ·asinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAsinPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskAsinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func asinPs(a [16]float32) [16]float32
TEXT ·asinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAsinPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskAsinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func asinhPd(a [8]float64) [8]float64
TEXT ·asinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAsinhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskAsinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func asinhPs(a [16]float32) [16]float32
TEXT ·asinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAsinhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskAsinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func atanPd(a [8]float64) [8]float64
TEXT ·atanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAtanPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskAtanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func atanPs(a [16]float32) [16]float32
TEXT ·atanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAtanPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskAtanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func atan2Pd(a [8]float64, b [8]float64) [8]float64
TEXT ·atan2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAtan2Pd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskAtan2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func atan2Ps(a [16]float32, b [16]float32) [16]float32
TEXT ·atan2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAtan2Ps(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskAtan2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func atanhPd(a [8]float64) [8]float64
TEXT ·atanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAtanhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskAtanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func atanhPs(a [16]float32) [16]float32
TEXT ·atanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskAtanhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskAtanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBlendEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlendEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskBlendEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskBlendEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlendEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskBlendEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskBlendPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskBlendPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlendPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskBlendPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskBlendPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskBlendPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlendPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskBlendPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func broadcastF32x4(a [4]float32) [8]float32
TEXT ·broadcastF32x4(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskBroadcastF32x4(src [8]float32, k uint8, a [4]float32) [8]float32
TEXT ·maskBroadcastF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastF32x4(k uint8, a [4]float32) [8]float32
TEXT ·maskzBroadcastF32x4(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastF32x41(a [4]float32) [16]float32
TEXT ·broadcastF32x41(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastF32x41(src [16]float32, k uint16, a [4]float32) [16]float32
TEXT ·maskBroadcastF32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastF32x41(k uint16, a [4]float32) [16]float32
TEXT ·maskzBroadcastF32x41(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func broadcastF64x4(a [4]float64) [8]float64
TEXT ·broadcastF64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBroadcastF64x4(src [8]float64, k uint8, a [4]float64) [8]float64
TEXT ·maskBroadcastF64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastF64x4(k uint8, a [4]float64) [8]float64
TEXT ·maskzBroadcastF64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func broadcastI32x4(a [16]byte) [32]byte
TEXT ·broadcastI32x4(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskBroadcastI32x4(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskBroadcastI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastI32x4(k uint8, a [16]byte) [32]byte
TEXT ·maskzBroadcastI32x4(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastI32x41(a [16]byte) [64]byte
TEXT ·broadcastI32x41(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastI32x41(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·maskBroadcastI32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastI32x41(k uint16, a [16]byte) [64]byte
TEXT ·maskzBroadcastI32x41(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func broadcastI64x4(a [32]byte) [64]byte
TEXT ·broadcastI64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBroadcastI64x4(src [64]byte, k uint8, a [32]byte) [64]byte
TEXT ·maskBroadcastI64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastI64x4(k uint8, a [32]byte) [64]byte
TEXT ·maskzBroadcastI64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskBroadcastdEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastdEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastdEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskBroadcastdEpi321(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskBroadcastdEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastdEpi321(k uint8, a [16]byte) [32]byte
TEXT ·maskzBroadcastdEpi321(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastdEpi32(a [16]byte) [64]byte
TEXT ·broadcastdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastdEpi322(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·maskBroadcastdEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastdEpi322(k uint16, a [16]byte) [64]byte
TEXT ·maskzBroadcastdEpi322(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskBroadcastqEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskBroadcastqEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastqEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzBroadcastqEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskBroadcastqEpi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskBroadcastqEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastqEpi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzBroadcastqEpi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastqEpi64(a [16]byte) [64]byte
TEXT ·broadcastqEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastqEpi642(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·maskBroadcastqEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastqEpi642(k uint8, a [16]byte) [64]byte
TEXT ·maskzBroadcastqEpi642(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskBroadcastsdPd(src [4]float64, k uint8, a [2]float64) [4]float64
TEXT ·maskBroadcastsdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastsdPd(k uint8, a [2]float64) [4]float64
TEXT ·maskzBroadcastsdPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastsdPd(a [2]float64) [8]float64
TEXT ·broadcastsdPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastsdPd1(src [8]float64, k uint8, a [2]float64) [8]float64
TEXT ·maskBroadcastsdPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastsdPd1(k uint8, a [2]float64) [8]float64
TEXT ·maskzBroadcastsdPd1(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskBroadcastssPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskBroadcastssPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastssPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzBroadcastssPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskBroadcastssPs1(src [8]float32, k uint8, a [4]float32) [8]float32
TEXT ·maskBroadcastssPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzBroadcastssPs1(k uint8, a [4]float32) [8]float32
TEXT ·maskzBroadcastssPs1(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func broadcastssPs(a [4]float32) [16]float32
TEXT ·broadcastssPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskBroadcastssPs2(src [16]float32, k uint16, a [4]float32) [16]float32
TEXT ·maskBroadcastssPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzBroadcastssPs2(k uint16, a [4]float32) [16]float32
TEXT ·maskzBroadcastssPs2(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func castpd128Pd512(a [2]float64) [8]float64
TEXT ·castpd128Pd512(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func castpd256Pd512(a [4]float64) [8]float64
TEXT ·castpd256Pd512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func castpd512Pd128(a [8]float64) [2]float64
TEXT ·castpd512Pd128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func castpd512Pd256(a [8]float64) [4]float64
TEXT ·castpd512Pd256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func castps128Ps512(a [4]float32) [16]float32
TEXT ·castps128Ps512(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func castps256Ps512(a [8]float32) [16]float32
TEXT ·castps256Ps512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func castps512Ps128(a [16]float32) [4]float32
TEXT ·castps512Ps128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func castps512Ps256(a [16]float32) [8]float32
TEXT ·castps512Ps256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func castsi128Si512(a [16]byte) [64]byte
TEXT ·castsi128Si512(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func castsi256Si512(a [32]byte) [64]byte
TEXT ·castsi256Si512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func castsi512Si128(a [64]byte) [16]byte
TEXT ·castsi512Si128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func castsi512Si256(a [64]byte) [32]byte
TEXT ·castsi512Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cbrtPd(a [8]float64) [8]float64
TEXT ·cbrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCbrtPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCbrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cbrtPs(a [16]float32) [16]float32
TEXT ·cbrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCbrtPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCbrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cdfnormPd(a [8]float64) [8]float64
TEXT ·cdfnormPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCdfnormPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCdfnormPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cdfnormPs(a [16]float32) [16]float32
TEXT ·cdfnormPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCdfnormPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCdfnormPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cdfnorminvPd(a [8]float64) [8]float64
TEXT ·cdfnorminvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCdfnorminvPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCdfnorminvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cdfnorminvPs(a [16]float32) [16]float32
TEXT ·cdfnorminvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCdfnorminvPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCdfnorminvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func ceilPd(a [8]float64) [8]float64
TEXT ·ceilPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCeilPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCeilPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func ceilPs(a [16]float32) [16]float32
TEXT ·ceilPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCeilPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCeilPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cmpEpi32Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpi32Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmpEpi32Mask1(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·cmpEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpEpi32Mask1(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpi64Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpi64Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmpEpi64Mask1(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·cmpEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpEpi64Mask1(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpi64Mask2(a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·cmpEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpEpi64Mask2(k1 uint8, a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·maskCmpEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpu32Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu32Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmpEpu32Mask1(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·cmpEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpEpu32Mask1(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpu64Mask(a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·cmpEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu64Mask(k1 uint8, a [16]byte, b [16]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmpEpu64Mask1(a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·cmpEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpEpu64Mask1(k1 uint8, a [32]byte, b [32]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpEpu64Mask2(a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·cmpEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpEpu64Mask2(k1 uint8, a [64]byte, b [64]byte, imm8 uint8) uint8
TEXT ·maskCmpEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpPdMask(a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·cmpPdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpPdMask(k1 uint8, a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·maskCmpPdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpPdMask1(a [4]float64, b [4]float64, imm8 int) uint8
TEXT ·cmpPdMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpPdMask1(k1 uint8, a [4]float64, b [4]float64, imm8 int) uint8
TEXT ·maskCmpPdMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpPsMask(a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·cmpPsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpPsMask(k1 uint8, a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·maskCmpPsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpPsMask1(a [8]float32, b [8]float32, imm8 int) uint8
TEXT ·cmpPsMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpPsMask1(k1 uint8, a [8]float32, b [8]float32, imm8 int) uint8
TEXT ·maskCmpPsMask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpRoundSdMask(a [2]float64, b [2]float64, imm8 int, sae int) uint8
TEXT ·cmpRoundSdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	//TODO: Code missing

	MOVB $0, ret+48(FP)
	RET

// func maskCmpRoundSdMask(k1 uint8, a [2]float64, b [2]float64, imm8 int, sae int) uint8
TEXT ·maskCmpRoundSdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ sae+44(FP),R12

	//TODO: Code missing

	MOVB $0, ret+52(FP)
	RET

// func cmpRoundSsMask(a [4]float32, b [4]float32, imm8 int, sae int) uint8
TEXT ·cmpRoundSsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	//TODO: Code missing

	MOVB $0, ret+48(FP)
	RET

// func maskCmpRoundSsMask(k1 uint8, a [4]float32, b [4]float32, imm8 int, sae int) uint8
TEXT ·maskCmpRoundSsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ sae+44(FP),R12

	//TODO: Code missing

	MOVB $0, ret+52(FP)
	RET

// func cmpSdMask(a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·cmpSdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpSdMask(k1 uint8, a [2]float64, b [2]float64, imm8 int) uint8
TEXT ·maskCmpSdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpSsMask(a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·cmpSsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpSsMask(k1 uint8, a [4]float32, b [4]float32, imm8 int) uint8
TEXT ·maskCmpSsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpeqEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpeqEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpeqEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpeqEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpeqEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpeqEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpeqEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpeqEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpeqEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpeqEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpu32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpeqEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpeqEpu32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpeqEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpeqEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpeqEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpu64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpeqEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpeqEpu64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpeqEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpeqEpu64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpeqEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpeqEpu64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpeqEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgeEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgeEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgeEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgeEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgeEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgeEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpgeEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgeEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpgeEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpu32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgeEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgeEpu32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgeEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgeEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgeEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpu64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgeEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgeEpu64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgeEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgeEpu64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpgeEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgeEpu64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpgeEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgtEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgtEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgtEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgtEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgtEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgtEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpgtEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgtEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpgtEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpu32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgtEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgtEpu32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgtEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpgtEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpgtEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpu64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpgtEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgtEpu64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpgtEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpgtEpu64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpgtEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpgtEpu64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpgtEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpleEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpleEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpleEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpleEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpleEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpleEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpleEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpleEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpleEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpu32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpleEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpleEpu32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpleEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpleEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpleEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpu64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpleEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpleEpu64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpleEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpleEpu64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpleEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpleEpu64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpleEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpltEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpltEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpltEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpi32Mask2(a [64]byte, b [64]byte) uint16
TEXT ·cmpltEpi32Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskCmpltEpi32Mask2(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·maskCmpltEpi32Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func cmpltEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpltEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpltEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpltEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpltEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpltEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpltEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpu32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpltEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpltEpu32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpltEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpltEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpltEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpu64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpltEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpltEpu64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpltEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpltEpu64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpltEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpltEpu64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpltEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpneqEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpneqEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpneqEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpneqEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpneqEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpneqEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpneqEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpneqEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpneqEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpu32Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpu32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpneqEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpneqEpu32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpneqEpu32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpu64Mask(a [16]byte, b [16]byte) uint8
TEXT ·cmpneqEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskCmpneqEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpu64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·cmpneqEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpneqEpu64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskCmpneqEpu64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func cmpneqEpu64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·cmpneqEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskCmpneqEpu64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskCmpneqEpu64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func comiRoundSd(a [2]float64, b [2]float64, imm8 int, sae int) int
TEXT ·comiRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	//TODO: Code missing

	MOVQ $0, ret+48(FP)
	RET

// func comiRoundSs(a [4]float32, b [4]float32, imm8 int, sae int) int
TEXT ·comiRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ sae+40(FP),R11

	//TODO: Code missing

	MOVQ $0, ret+48(FP)
	RET

// func maskCompressEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCompressEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCompressEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCompressEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCompressEpi321(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskCompressEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCompressEpi321(k uint8, a [32]byte) [32]byte
TEXT ·maskzCompressEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCompressEpi322(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·maskCompressEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCompressEpi322(k uint16, a [64]byte) [64]byte
TEXT ·maskzCompressEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCompressEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCompressEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCompressEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCompressEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCompressEpi641(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskCompressEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCompressEpi641(k uint8, a [32]byte) [32]byte
TEXT ·maskzCompressEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCompressEpi642(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·maskCompressEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCompressEpi642(k uint8, a [64]byte) [64]byte
TEXT ·maskzCompressEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCompressPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskCompressPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCompressPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzCompressPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCompressPd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskCompressPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCompressPd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzCompressPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCompressPd2(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCompressPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCompressPd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzCompressPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCompressPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskCompressPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCompressPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzCompressPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCompressPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskCompressPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCompressPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzCompressPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCompressPs2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCompressPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCompressPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzCompressPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCompressstoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCompressstoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuEpi321(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCompressstoreuEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuEpi322(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCompressstoreuEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuEpi64(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCompressstoreuEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuEpi641(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCompressstoreuEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuEpi642(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCompressstoreuEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPd(base_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskCompressstoreuPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPd1(base_addr uintptr, k uint8, a [4]float64) 
TEXT ·maskCompressstoreuPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPd2(base_addr uintptr, k uint8, a [8]float64) 
TEXT ·maskCompressstoreuPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPs(base_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskCompressstoreuPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPs1(base_addr uintptr, k uint8, a [8]float32) 
TEXT ·maskCompressstoreuPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPs2(base_addr uintptr, k uint16, a [16]float32) 
TEXT ·maskCompressstoreuPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cosPd(a [8]float64) [8]float64
TEXT ·cosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCosPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cosPs(a [16]float32) [16]float32
TEXT ·cosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCosPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cosdPd(a [8]float64) [8]float64
TEXT ·cosdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCosdPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCosdPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cosdPs(a [16]float32) [16]float32
TEXT ·cosdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCosdPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCosdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func coshPd(a [8]float64) [8]float64
TEXT ·coshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCoshPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskCoshPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func coshPs(a [16]float32) [16]float32
TEXT ·coshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCoshPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskCoshPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundepi32Ps(a [64]byte, rounding int) [16]float32
TEXT ·cvtRoundepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundepi32Ps(src [16]float32, k uint16, a [64]byte, rounding int) [16]float32
TEXT ·maskCvtRoundepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundepi32Ps(k uint16, a [64]byte, rounding int) [16]float32
TEXT ·maskzCvtRoundepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundepu32Ps(a [64]byte, rounding int) [16]float32
TEXT ·cvtRoundepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundepu32Ps(src [16]float32, k uint16, a [64]byte, rounding int) [16]float32
TEXT ·maskCvtRoundepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundepu32Ps(k uint16, a [64]byte, rounding int) [16]float32
TEXT ·maskzCvtRoundepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundi32Ss(a [4]float32, b int, rounding int) [4]float32
TEXT ·cvtRoundi32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundi64Sd(a [2]float64, b int64, rounding int) [2]float64
TEXT ·cvtRoundi64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundi64Ss(a [4]float32, b int64, rounding int) [4]float32
TEXT ·cvtRoundi64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundpdEpi32(a [8]float64, rounding int) [32]byte
TEXT ·cvtRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtRoundpdEpi32(src [32]byte, k uint8, a [8]float64, rounding int) [32]byte
TEXT ·maskCvtRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtRoundpdEpi32(k uint8, a [8]float64, rounding int) [32]byte
TEXT ·maskzCvtRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtRoundpdEpu32(a [8]float64, rounding int) [32]byte
TEXT ·cvtRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtRoundpdEpu32(src [32]byte, k uint8, a [8]float64, rounding int) [32]byte
TEXT ·maskCvtRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtRoundpdEpu32(k uint8, a [8]float64, rounding int) [32]byte
TEXT ·maskzCvtRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtRoundpdPs(a [8]float64, rounding int) [8]float32
TEXT ·cvtRoundpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtRoundpdPs(src [8]float32, k uint8, a [8]float64, rounding int) [8]float32
TEXT ·maskCvtRoundpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtRoundpdPs(k uint8, a [8]float64, rounding int) [8]float32
TEXT ·maskzCvtRoundpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtRoundphPs(a [32]byte, sae int) [16]float32
TEXT ·cvtRoundphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundphPs(src [16]float32, k uint16, a [32]byte, sae int) [16]float32
TEXT ·maskCvtRoundphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundphPs(k uint16, a [32]byte, sae int) [16]float32
TEXT ·maskzCvtRoundphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundpsEpi32(a [16]float32, rounding int) [64]byte
TEXT ·cvtRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpsEpi32(src [64]byte, k uint16, a [16]float32, rounding int) [64]byte
TEXT ·maskCvtRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpsEpi32(k uint16, a [16]float32, rounding int) [64]byte
TEXT ·maskzCvtRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundpsEpu32(a [16]float32, rounding int) [64]byte
TEXT ·cvtRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpsEpu32(src [64]byte, k uint16, a [16]float32, rounding int) [64]byte
TEXT ·maskCvtRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpsEpu32(k uint16, a [16]float32, rounding int) [64]byte
TEXT ·maskzCvtRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtRoundpsPd(a [8]float32, sae int) [8]float64
TEXT ·cvtRoundpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpsPd(src [8]float64, k uint8, a [8]float32, sae int) [8]float64
TEXT ·maskCvtRoundpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtRoundpsPd(k uint8, a [8]float32, sae int) [8]float64
TEXT ·maskzCvtRoundpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtRoundpsPh(src [16]byte, k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskCvtRoundpsPh(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzCvtRoundpsPh(k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskzCvtRoundpsPh(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ rounding+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskCvtRoundpsPh1(src [16]byte, k uint8, a [8]float32, rounding int) [16]byte
TEXT ·maskCvtRoundpsPh1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtRoundpsPh1(k uint8, a [8]float32, rounding int) [16]byte
TEXT ·maskzCvtRoundpsPh1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtRoundpsPh(a [16]float32, rounding int) [32]byte
TEXT ·cvtRoundpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtRoundpsPh2(src [32]byte, k uint16, a [16]float32, rounding int) [32]byte
TEXT ·maskCvtRoundpsPh2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtRoundpsPh2(k uint16, a [16]float32, rounding int) [32]byte
TEXT ·maskzCvtRoundpsPh2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtRoundsdI32(a [2]float64, rounding int) int
TEXT ·cvtRoundsdI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdI64(a [2]float64, rounding int) int64
TEXT ·cvtRoundsdI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdSi32(a [2]float64, rounding int) int
TEXT ·cvtRoundsdSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdSi64(a [2]float64, rounding int) int64
TEXT ·cvtRoundsdSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsdSs(a [4]float32, b [2]float64, rounding int) [4]float32
TEXT ·cvtRoundsdSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskCvtRoundsdSs(src [4]float32, k uint8, a [4]float32, b [2]float64, rounding int) [4]float32
TEXT ·maskCvtRoundsdSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzCvtRoundsdSs(k uint8, a [4]float32, b [2]float64, rounding int) [4]float32
TEXT ·maskzCvtRoundsdSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func cvtRoundsdU32(a [2]float64, rounding int) uint32
TEXT ·cvtRoundsdU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVL $0, ret+24(FP)
	RET

// func cvtRoundsdU64(a [2]float64, rounding int) uint64
TEXT ·cvtRoundsdU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundsi32Ss(a [4]float32, b int, rounding int) [4]float32
TEXT ·cvtRoundsi32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundsi64Sd(a [2]float64, b int64, rounding int) [2]float64
TEXT ·cvtRoundsi64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundsi64Ss(a [4]float32, b int64, rounding int) [4]float32
TEXT ·cvtRoundsi64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundssI32(a [4]float32, rounding int) int
TEXT ·cvtRoundssI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssI64(a [4]float32, rounding int) int64
TEXT ·cvtRoundssI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssSd(a [2]float64, b [4]float32, rounding int) [2]float64
TEXT ·cvtRoundssSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskCvtRoundssSd(src [2]float64, k uint8, a [2]float64, b [4]float32, rounding int) [2]float64
TEXT ·maskCvtRoundssSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzCvtRoundssSd(k uint8, a [2]float64, b [4]float32, rounding int) [2]float64
TEXT ·maskzCvtRoundssSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func cvtRoundssSi32(a [4]float32, rounding int) int
TEXT ·cvtRoundssSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssSi64(a [4]float32, rounding int) int64
TEXT ·cvtRoundssSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundssU32(a [4]float32, rounding int) uint32
TEXT ·cvtRoundssU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVL $0, ret+24(FP)
	RET

// func cvtRoundssU64(a [4]float32, rounding int) uint64
TEXT ·cvtRoundssU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvtRoundu32Ss(a [4]float32, b uint32, rounding int) [4]float32
TEXT ·cvtRoundu32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL b+16(FP),R9
	MOVQ rounding+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func cvtRoundu64Sd(a [2]float64, b uint64, rounding int) [2]float64
TEXT ·cvtRoundu64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvtRoundu64Ss(a [4]float32, b uint64, rounding int) [4]float32
TEXT ·cvtRoundu64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9
	MOVQ rounding+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskCvtepi16Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi16Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi16Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi16Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi16Epi321(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepi16Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi16Epi321(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepi16Epi321(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi16Epi32(a [32]byte) [64]byte
TEXT ·cvtepi16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi16Epi322(src [64]byte, k uint16, a [32]byte) [64]byte
TEXT ·maskCvtepi16Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi16Epi322(k uint16, a [32]byte) [64]byte
TEXT ·maskzCvtepi16Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi16Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi16Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi16Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi16Epi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepi16Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi16Epi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepi16Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi16Epi64(a [16]byte) [64]byte
TEXT ·cvtepi16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskCvtepi16Epi642(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·maskCvtepi16Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi16Epi642(k uint8, a [16]byte) [64]byte
TEXT ·maskzCvtepi16Epi642(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func cvtepi32Epi16(a [16]byte) [16]byte
TEXT ·cvtepi32Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi32Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi32Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi32Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi32Epi161(a [32]byte) [16]byte
TEXT ·cvtepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi32Epi161(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi32Epi161(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi32Epi162(a [64]byte) [32]byte
TEXT ·cvtepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi32Epi162(src [32]byte, k uint16, a [64]byte) [32]byte
TEXT ·maskCvtepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi32Epi162(k uint16, a [64]byte) [32]byte
TEXT ·maskzCvtepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi32Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi32Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi32Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi32Epi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepi32Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi32Epi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepi32Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi32Epi64(a [32]byte) [64]byte
TEXT ·cvtepi32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi32Epi642(src [64]byte, k uint8, a [32]byte) [64]byte
TEXT ·maskCvtepi32Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi32Epi642(k uint8, a [32]byte) [64]byte
TEXT ·maskzCvtepi32Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtepi32Epi8(a [16]byte) [16]byte
TEXT ·cvtepi32Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi32Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi32Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi32Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi32Epi81(a [32]byte) [16]byte
TEXT ·cvtepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi32Epi81(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi32Epi81(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi32Epi82(a [64]byte) [16]byte
TEXT ·cvtepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi32Epi82(src [16]byte, k uint16, a [64]byte) [16]byte
TEXT ·maskCvtepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi32Epi82(k uint16, a [64]byte) [16]byte
TEXT ·maskzCvtepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi32Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepi32Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepi32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi32Pd1(src [4]float64, k uint8, a [16]byte) [4]float64
TEXT ·maskCvtepi32Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi32Pd1(k uint8, a [16]byte) [4]float64
TEXT ·maskzCvtepi32Pd1(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi32Pd(a [32]byte) [8]float64
TEXT ·cvtepi32Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi32Pd2(src [8]float64, k uint8, a [32]byte) [8]float64
TEXT ·maskCvtepi32Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi32Pd2(k uint8, a [32]byte) [8]float64
TEXT ·maskzCvtepi32Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi32Ps(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtepi32Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Ps(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtepi32Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi32Ps1(src [8]float32, k uint8, a [32]byte) [8]float32
TEXT ·maskCvtepi32Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi32Ps1(k uint8, a [32]byte) [8]float32
TEXT ·maskzCvtepi32Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtepi32Ps(a [64]byte) [16]float32
TEXT ·cvtepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi32Ps2(src [16]float32, k uint16, a [64]byte) [16]float32
TEXT ·maskCvtepi32Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi32Ps2(k uint16, a [64]byte) [16]float32
TEXT ·maskzCvtepi32Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi161(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtepi32StoreuEpi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi162(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCvtepi32StoreuEpi162(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi81(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtepi32StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi82(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCvtepi32StoreuEpi82(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtepi64Epi16(a [16]byte) [16]byte
TEXT ·cvtepi64Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi64Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi64Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Epi161(a [32]byte) [16]byte
TEXT ·cvtepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64Epi161(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi64Epi161(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi64Epi162(a [64]byte) [16]byte
TEXT ·cvtepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64Epi162(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·maskCvtepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi64Epi162(k uint8, a [64]byte) [16]byte
TEXT ·maskzCvtepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi64Epi32(a [16]byte) [16]byte
TEXT ·cvtepi64Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi64Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi64Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Epi321(a [32]byte) [16]byte
TEXT ·cvtepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64Epi321(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi64Epi321(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi64Epi322(a [64]byte) [32]byte
TEXT ·cvtepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtepi64Epi322(src [32]byte, k uint8, a [64]byte) [32]byte
TEXT ·maskCvtepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi64Epi322(k uint8, a [64]byte) [32]byte
TEXT ·maskzCvtepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtepi64Epi8(a [16]byte) [16]byte
TEXT ·cvtepi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi64Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Epi81(a [32]byte) [16]byte
TEXT ·cvtepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64Epi81(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi64Epi81(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtepi64Epi82(a [64]byte) [16]byte
TEXT ·cvtepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64Epi82(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·maskCvtepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtepi64Epi82(k uint8, a [64]byte) [16]byte
TEXT ·maskzCvtepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi161(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtepi64StoreuEpi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi162(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtepi64StoreuEpi162(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi321(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtepi64StoreuEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi322(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtepi64StoreuEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi81(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtepi64StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi82(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtepi64StoreuEpi82(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi8Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi8Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi8Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi8Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi8Epi321(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepi8Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi8Epi321(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepi8Epi321(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi8Epi32(a [16]byte) [64]byte
TEXT ·cvtepi8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskCvtepi8Epi322(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·maskCvtepi8Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi8Epi322(k uint16, a [16]byte) [64]byte
TEXT ·maskzCvtepi8Epi322(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskCvtepi8Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepi8Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi8Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepi8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi8Epi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepi8Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepi8Epi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepi8Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepi8Epi64(a [16]byte) [64]byte
TEXT ·cvtepi8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskCvtepi8Epi642(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·maskCvtepi8Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepi8Epi642(k uint8, a [16]byte) [64]byte
TEXT ·maskzCvtepi8Epi642(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskCvtepu16Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu16Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu16Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu16Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepu16Epi321(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepu16Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu16Epi321(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepu16Epi321(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu16Epi32(a [32]byte) [64]byte
TEXT ·cvtepu16Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu16Epi322(src [64]byte, k uint16, a [32]byte) [64]byte
TEXT ·maskCvtepu16Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu16Epi322(k uint16, a [32]byte) [64]byte
TEXT ·maskzCvtepu16Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu16Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu16Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu16Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu16Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepu16Epi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepu16Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu16Epi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepu16Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu16Epi64(a [16]byte) [64]byte
TEXT ·cvtepu16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskCvtepu16Epi642(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·maskCvtepu16Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu16Epi642(k uint8, a [16]byte) [64]byte
TEXT ·maskzCvtepu16Epi642(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskCvtepu32Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu32Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu32Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu32Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepu32Epi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepu32Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu32Epi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepu32Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu32Epi64(a [32]byte) [64]byte
TEXT ·cvtepu32Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu32Epi642(src [64]byte, k uint8, a [32]byte) [64]byte
TEXT ·maskCvtepu32Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu32Epi642(k uint8, a [32]byte) [64]byte
TEXT ·maskzCvtepu32Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtepu32Pd(a [16]byte) [2]float64
TEXT ·cvtepu32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu32Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepu32Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu32Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepu32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu32Pd1(a [16]byte) [4]float64
TEXT ·cvtepu32Pd1(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func maskCvtepu32Pd1(src [4]float64, k uint8, a [16]byte) [4]float64
TEXT ·maskCvtepu32Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu32Pd1(k uint8, a [16]byte) [4]float64
TEXT ·maskzCvtepu32Pd1(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu32Pd2(a [32]byte) [8]float64
TEXT ·cvtepu32Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu32Pd2(src [8]float64, k uint8, a [32]byte) [8]float64
TEXT ·maskCvtepu32Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu32Pd2(k uint8, a [32]byte) [8]float64
TEXT ·maskzCvtepu32Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtepu32Ps(a [64]byte) [16]float32
TEXT ·cvtepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu32Ps(src [16]float32, k uint16, a [64]byte) [16]float32
TEXT ·maskCvtepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu32Ps(k uint16, a [64]byte) [16]float32
TEXT ·maskzCvtepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtepu8Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu8Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu8Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu8Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepu8Epi321(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepu8Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu8Epi321(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepu8Epi321(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu8Epi32(a [16]byte) [64]byte
TEXT ·cvtepu8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskCvtepu8Epi322(src [64]byte, k uint16, a [16]byte) [64]byte
TEXT ·maskCvtepu8Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu8Epi322(k uint16, a [16]byte) [64]byte
TEXT ·maskzCvtepu8Epi322(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func maskCvtepu8Epi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtepu8Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu8Epi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtepu8Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepu8Epi641(src [32]byte, k uint8, a [16]byte) [32]byte
TEXT ·maskCvtepu8Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtepu8Epi641(k uint8, a [16]byte) [32]byte
TEXT ·maskzCvtepu8Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtepu8Epi64(a [16]byte) [64]byte
TEXT ·cvtepu8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func maskCvtepu8Epi642(src [64]byte, k uint8, a [16]byte) [64]byte
TEXT ·maskCvtepu8Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtepu8Epi642(k uint8, a [16]byte) [64]byte
TEXT ·maskzCvtepu8Epi642(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Z0, ret+20(FP)
	RET

// func cvti32Sd(a [2]float64, b int) [2]float64
TEXT ·cvti32Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvti32Ss(a [4]float32, b int) [4]float32
TEXT ·cvti32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvti64Sd(a [2]float64, b int64) [2]float64
TEXT ·cvti64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvti64Ss(a [4]float32, b int64) [4]float32
TEXT ·cvti64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskCvtpdEpi32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpi32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtpdEpi321(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·maskCvtpdEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtpdEpi321(k uint8, a [4]float64) [16]byte
TEXT ·maskzCvtpdEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtpdEpi32(a [8]float64) [32]byte
TEXT ·cvtpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpdEpi322(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·maskCvtpdEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpdEpi322(k uint8, a [8]float64) [32]byte
TEXT ·maskzCvtpdEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtpdEpu32(a [2]float64) [16]byte
TEXT ·cvtpdEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpu32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpu32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpdEpu321(a [4]float64) [16]byte
TEXT ·cvtpdEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtpdEpu321(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·maskCvtpdEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtpdEpu321(k uint8, a [4]float64) [16]byte
TEXT ·maskzCvtpdEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtpdEpu322(a [8]float64) [32]byte
TEXT ·cvtpdEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpdEpu322(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·maskCvtpdEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpdEpu322(k uint8, a [8]float64) [32]byte
TEXT ·maskzCvtpdEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpdPs(src [4]float32, k uint8, a [2]float64) [4]float32
TEXT ·maskCvtpdPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdPs(k uint8, a [2]float64) [4]float32
TEXT ·maskzCvtpdPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtpdPs1(src [4]float32, k uint8, a [4]float64) [4]float32
TEXT ·maskCvtpdPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtpdPs1(k uint8, a [4]float64) [4]float32
TEXT ·maskzCvtpdPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtpdPs(a [8]float64) [8]float32
TEXT ·cvtpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpdPs2(src [8]float32, k uint8, a [8]float64) [8]float32
TEXT ·maskCvtpdPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpdPs2(k uint8, a [8]float64) [8]float32
TEXT ·maskzCvtpdPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtphPs(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtphPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtphPs(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtphPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtphPs1(src [8]float32, k uint8, a [16]byte) [8]float32
TEXT ·maskCvtphPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtphPs1(k uint8, a [16]byte) [8]float32
TEXT ·maskzCvtphPs1(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOV Y0, ret+20(FP)
	RET

// func cvtphPs(a [32]byte) [16]float32
TEXT ·cvtphPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtphPs2(src [16]float32, k uint16, a [32]byte) [16]float32
TEXT ·maskCvtphPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtphPs2(k uint16, a [32]byte) [16]float32
TEXT ·maskzCvtphPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsEpi32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpi32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtpsEpi321(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·maskCvtpsEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpsEpi321(k uint8, a [8]float32) [32]byte
TEXT ·maskzCvtpsEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtpsEpi32(a [16]float32) [64]byte
TEXT ·cvtpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsEpi322(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·maskCvtpsEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpsEpi322(k uint16, a [16]float32) [64]byte
TEXT ·maskzCvtpsEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtpsEpu32(a [4]float32) [16]byte
TEXT ·cvtpsEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpu32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpu32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpsEpu321(a [8]float32) [32]byte
TEXT ·cvtpsEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpsEpu321(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·maskCvtpsEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpsEpu321(k uint8, a [8]float32) [32]byte
TEXT ·maskzCvtpsEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtpsEpu322(a [16]float32) [64]byte
TEXT ·cvtpsEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsEpu322(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·maskCvtpsEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpsEpu322(k uint16, a [16]float32) [64]byte
TEXT ·maskzCvtpsEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvtpsPd(a [8]float32) [8]float64
TEXT ·cvtpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsPd(src [8]float64, k uint8, a [8]float32) [8]float64
TEXT ·maskCvtpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvtpsPd(k uint8, a [8]float32) [8]float64
TEXT ·maskzCvtpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvtpsPh(src [16]byte, k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskCvtpsPh(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzCvtpsPh(k uint8, a [4]float32, rounding int) [16]byte
TEXT ·maskzCvtpsPh(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ rounding+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskCvtpsPh1(src [16]byte, k uint8, a [8]float32, rounding int) [16]byte
TEXT ·maskCvtpsPh1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtpsPh1(k uint8, a [8]float32, rounding int) [16]byte
TEXT ·maskzCvtpsPh1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtpsPh(a [16]float32, rounding int) [32]byte
TEXT ·cvtpsPh(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtpsPh2(src [32]byte, k uint16, a [16]float32, rounding int) [32]byte
TEXT ·maskCvtpsPh2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtpsPh2(k uint16, a [16]float32, rounding int) [32]byte
TEXT ·maskzCvtpsPh2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtsdI32(a [2]float64) int
TEXT ·cvtsdI32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsdI64(a [2]float64) int64
TEXT ·cvtsdI64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func maskCvtsdSs(src [4]float32, k uint8, a [4]float32, b [2]float64) [4]float32
TEXT ·maskCvtsdSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzCvtsdSs(k uint8, a [4]float32, b [2]float64) [4]float32
TEXT ·maskzCvtsdSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func cvtsdU32(a [2]float64) uint32
TEXT ·cvtsdU32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVL $0, ret+16(FP)
	RET

// func cvtsdU64(a [2]float64) uint64
TEXT ·cvtsdU64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsepi32Epi16(a [16]byte) [16]byte
TEXT ·cvtsepi32Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi32Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi32Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi32Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi32Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi32Epi161(a [32]byte) [16]byte
TEXT ·cvtsepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi32Epi161(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtsepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi32Epi161(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtsepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi32Epi162(a [64]byte) [32]byte
TEXT ·cvtsepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtsepi32Epi162(src [32]byte, k uint16, a [64]byte) [32]byte
TEXT ·maskCvtsepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtsepi32Epi162(k uint16, a [64]byte) [32]byte
TEXT ·maskzCvtsepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtsepi32Epi8(a [16]byte) [16]byte
TEXT ·cvtsepi32Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi32Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi32Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi32Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi32Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi32Epi81(a [32]byte) [16]byte
TEXT ·cvtsepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi32Epi81(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtsepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi32Epi81(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtsepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi32Epi82(a [64]byte) [16]byte
TEXT ·cvtsepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi32Epi82(src [16]byte, k uint16, a [64]byte) [16]byte
TEXT ·maskCvtsepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi32Epi82(k uint16, a [64]byte) [16]byte
TEXT ·maskzCvtsepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi161(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtsepi32StoreuEpi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi162(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCvtsepi32StoreuEpi162(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi81(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtsepi32StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi82(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCvtsepi32StoreuEpi82(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtsepi64Epi16(a [16]byte) [16]byte
TEXT ·cvtsepi64Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi64Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi64Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi64Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi64Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi64Epi161(a [32]byte) [16]byte
TEXT ·cvtsepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi64Epi161(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtsepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi64Epi161(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtsepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi64Epi162(a [64]byte) [16]byte
TEXT ·cvtsepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi64Epi162(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·maskCvtsepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi64Epi162(k uint8, a [64]byte) [16]byte
TEXT ·maskzCvtsepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi64Epi32(a [16]byte) [16]byte
TEXT ·cvtsepi64Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi64Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi64Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi64Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi64Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi64Epi321(a [32]byte) [16]byte
TEXT ·cvtsepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi64Epi321(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtsepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi64Epi321(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtsepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi64Epi322(a [64]byte) [32]byte
TEXT ·cvtsepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtsepi64Epi322(src [32]byte, k uint8, a [64]byte) [32]byte
TEXT ·maskCvtsepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtsepi64Epi322(k uint8, a [64]byte) [32]byte
TEXT ·maskzCvtsepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtsepi64Epi8(a [16]byte) [16]byte
TEXT ·cvtsepi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi64Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtsepi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi64Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtsepi64Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi64Epi81(a [32]byte) [16]byte
TEXT ·cvtsepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi64Epi81(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtsepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi64Epi81(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtsepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtsepi64Epi82(a [64]byte) [16]byte
TEXT ·cvtsepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi64Epi82(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·maskCvtsepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtsepi64Epi82(k uint8, a [64]byte) [16]byte
TEXT ·maskzCvtsepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtsepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi161(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtsepi64StoreuEpi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi162(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtsepi64StoreuEpi162(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi321(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtsepi64StoreuEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi322(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtsepi64StoreuEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi81(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtsepi64StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi82(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtsepi64StoreuEpi82(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtssI32(a [4]float32) int
TEXT ·cvtssI32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtssI64(a [4]float32) int64
TEXT ·cvtssI64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func maskCvtssSd(src [2]float64, k uint8, a [2]float64, b [4]float32) [2]float64
TEXT ·maskCvtssSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzCvtssSd(k uint8, a [2]float64, b [4]float32) [2]float64
TEXT ·maskzCvtssSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func cvtssU32(a [4]float32) uint32
TEXT ·cvtssU32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVL $0, ret+16(FP)
	RET

// func cvtssU64(a [4]float32) uint64
TEXT ·cvtssU64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttRoundpdEpi32(a [8]float64, sae int) [32]byte
TEXT ·cvttRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttRoundpdEpi32(src [32]byte, k uint8, a [8]float64, sae int) [32]byte
TEXT ·maskCvttRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttRoundpdEpi32(k uint8, a [8]float64, sae int) [32]byte
TEXT ·maskzCvttRoundpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttRoundpdEpu32(a [8]float64, sae int) [32]byte
TEXT ·cvttRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttRoundpdEpu32(src [32]byte, k uint8, a [8]float64, sae int) [32]byte
TEXT ·maskCvttRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttRoundpdEpu32(k uint8, a [8]float64, sae int) [32]byte
TEXT ·maskzCvttRoundpdEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttRoundpsEpi32(a [16]float32, sae int) [64]byte
TEXT ·cvttRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttRoundpsEpi32(src [64]byte, k uint16, a [16]float32, sae int) [64]byte
TEXT ·maskCvttRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttRoundpsEpi32(k uint16, a [16]float32, sae int) [64]byte
TEXT ·maskzCvttRoundpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttRoundpsEpu32(a [16]float32, sae int) [64]byte
TEXT ·cvttRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttRoundpsEpu32(src [64]byte, k uint16, a [16]float32, sae int) [64]byte
TEXT ·maskCvttRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttRoundpsEpu32(k uint16, a [16]float32, sae int) [64]byte
TEXT ·maskzCvttRoundpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttRoundsdI32(a [2]float64, rounding int) int
TEXT ·cvttRoundsdI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdI64(a [2]float64, rounding int) int64
TEXT ·cvttRoundsdI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdSi32(a [2]float64, rounding int) int
TEXT ·cvttRoundsdSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdSi64(a [2]float64, rounding int) int64
TEXT ·cvttRoundsdSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundsdU32(a [2]float64, rounding int) uint32
TEXT ·cvttRoundsdU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVL $0, ret+24(FP)
	RET

// func cvttRoundsdU64(a [2]float64, rounding int) uint64
TEXT ·cvttRoundsdU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssI32(a [4]float32, rounding int) int
TEXT ·cvttRoundssI32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssI64(a [4]float32, rounding int) int64
TEXT ·cvttRoundssI64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssSi32(a [4]float32, rounding int) int
TEXT ·cvttRoundssSi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssSi64(a [4]float32, rounding int) int64
TEXT ·cvttRoundssSi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func cvttRoundssU32(a [4]float32, rounding int) uint32
TEXT ·cvttRoundssU32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVL $0, ret+24(FP)
	RET

// func cvttRoundssU64(a [4]float32, rounding int) uint64
TEXT ·cvttRoundssU64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func maskCvttpdEpi32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpi32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvttpdEpi321(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·maskCvttpdEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvttpdEpi321(k uint8, a [4]float64) [16]byte
TEXT ·maskzCvttpdEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvttpdEpi32(a [8]float64) [32]byte
TEXT ·cvttpdEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttpdEpi322(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·maskCvttpdEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpdEpi322(k uint8, a [8]float64) [32]byte
TEXT ·maskzCvttpdEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttpdEpu32(a [2]float64) [16]byte
TEXT ·cvttpdEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpu32(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpu32(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpdEpu321(a [4]float64) [16]byte
TEXT ·cvttpdEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvttpdEpu321(src [16]byte, k uint8, a [4]float64) [16]byte
TEXT ·maskCvttpdEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvttpdEpu321(k uint8, a [4]float64) [16]byte
TEXT ·maskzCvttpdEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvttpdEpu322(a [8]float64) [32]byte
TEXT ·cvttpdEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttpdEpu322(src [32]byte, k uint8, a [8]float64) [32]byte
TEXT ·maskCvttpdEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpdEpu322(k uint8, a [8]float64) [32]byte
TEXT ·maskzCvttpdEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttpsEpi32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpi32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvttpsEpi321(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·maskCvttpsEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpsEpi321(k uint8, a [8]float32) [32]byte
TEXT ·maskzCvttpsEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttpsEpi32(a [16]float32) [64]byte
TEXT ·cvttpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttpsEpi322(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·maskCvttpsEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttpsEpi322(k uint16, a [16]float32) [64]byte
TEXT ·maskzCvttpsEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttpsEpu32(a [4]float32) [16]byte
TEXT ·cvttpsEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpu32(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpu32(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpsEpu321(a [8]float32) [32]byte
TEXT ·cvttpsEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvttpsEpu321(src [32]byte, k uint8, a [8]float32) [32]byte
TEXT ·maskCvttpsEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvttpsEpu321(k uint8, a [8]float32) [32]byte
TEXT ·maskzCvttpsEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvttpsEpu322(a [16]float32) [64]byte
TEXT ·cvttpsEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskCvttpsEpu322(src [64]byte, k uint16, a [16]float32) [64]byte
TEXT ·maskCvttpsEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzCvttpsEpu322(k uint16, a [16]float32) [64]byte
TEXT ·maskzCvttpsEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func cvttsdI32(a [2]float64) int
TEXT ·cvttsdI32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttsdI64(a [2]float64) int64
TEXT ·cvttsdI64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttsdU32(a [2]float64) uint32
TEXT ·cvttsdU32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVL $0, ret+16(FP)
	RET

// func cvttsdU64(a [2]float64) uint64
TEXT ·cvttsdU64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttssI32(a [4]float32) int
TEXT ·cvttssI32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttssI64(a [4]float32) int64
TEXT ·cvttssI64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttssU32(a [4]float32) uint32
TEXT ·cvttssU32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVL $0, ret+16(FP)
	RET

// func cvttssU64(a [4]float32) uint64
TEXT ·cvttssU64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtu32Sd(a [2]float64, b uint32) [2]float64
TEXT ·cvtu32Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtu32Ss(a [4]float32, b uint32) [4]float32
TEXT ·cvtu32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtu64Sd(a [2]float64, b uint64) [2]float64
TEXT ·cvtu64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtu64Ss(a [4]float32, b uint64) [4]float32
TEXT ·cvtu64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtusepi32Epi16(a [16]byte) [16]byte
TEXT ·cvtusepi32Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi32Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi32Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi32Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi32Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi32Epi161(a [32]byte) [16]byte
TEXT ·cvtusepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi32Epi161(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtusepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi32Epi161(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtusepi32Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi32Epi162(a [64]byte) [32]byte
TEXT ·cvtusepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtusepi32Epi162(src [32]byte, k uint16, a [64]byte) [32]byte
TEXT ·maskCvtusepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtusepi32Epi162(k uint16, a [64]byte) [32]byte
TEXT ·maskzCvtusepi32Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtusepi32Epi8(a [16]byte) [16]byte
TEXT ·cvtusepi32Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi32Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi32Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi32Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi32Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi32Epi81(a [32]byte) [16]byte
TEXT ·cvtusepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi32Epi81(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtusepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi32Epi81(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtusepi32Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi32Epi82(a [64]byte) [16]byte
TEXT ·cvtusepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi32Epi82(src [16]byte, k uint16, a [64]byte) [16]byte
TEXT ·maskCvtusepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi32Epi82(k uint16, a [64]byte) [16]byte
TEXT ·maskzCvtusepi32Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi161(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtusepi32StoreuEpi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi162(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCvtusepi32StoreuEpi162(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi32StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi81(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtusepi32StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi82(base_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskCvtusepi32StoreuEpi82(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtusepi64Epi16(a [16]byte) [16]byte
TEXT ·cvtusepi64Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi64Epi16(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi64Epi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi64Epi16(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi64Epi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi64Epi161(a [32]byte) [16]byte
TEXT ·cvtusepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi64Epi161(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtusepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi64Epi161(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtusepi64Epi161(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi64Epi162(a [64]byte) [16]byte
TEXT ·cvtusepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi64Epi162(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·maskCvtusepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi64Epi162(k uint8, a [64]byte) [16]byte
TEXT ·maskzCvtusepi64Epi162(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi64Epi32(a [16]byte) [16]byte
TEXT ·cvtusepi64Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi64Epi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi64Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi64Epi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi64Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi64Epi321(a [32]byte) [16]byte
TEXT ·cvtusepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi64Epi321(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtusepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi64Epi321(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtusepi64Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi64Epi322(a [64]byte) [32]byte
TEXT ·cvtusepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskCvtusepi64Epi322(src [32]byte, k uint8, a [64]byte) [32]byte
TEXT ·maskCvtusepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzCvtusepi64Epi322(k uint8, a [64]byte) [32]byte
TEXT ·maskzCvtusepi64Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtusepi64Epi8(a [16]byte) [16]byte
TEXT ·cvtusepi64Epi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi64Epi8(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskCvtusepi64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi64Epi8(k uint8, a [16]byte) [16]byte
TEXT ·maskzCvtusepi64Epi8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi64Epi81(a [32]byte) [16]byte
TEXT ·cvtusepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi64Epi81(src [16]byte, k uint8, a [32]byte) [16]byte
TEXT ·maskCvtusepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi64Epi81(k uint8, a [32]byte) [16]byte
TEXT ·maskzCvtusepi64Epi81(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func cvtusepi64Epi82(a [64]byte) [16]byte
TEXT ·cvtusepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi64Epi82(src [16]byte, k uint8, a [64]byte) [16]byte
TEXT ·maskCvtusepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzCvtusepi64Epi82(k uint8, a [64]byte) [16]byte
TEXT ·maskzCvtusepi64Epi82(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskCvtusepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi161(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtusepi64StoreuEpi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi162(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtusepi64StoreuEpi162(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi321(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtusepi64StoreuEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi322(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtusepi64StoreuEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi81(base_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskCvtusepi64StoreuEpi81(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi82(base_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskCvtusepi64StoreuEpi82(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func divEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskDivEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskDivEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpu64(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·divEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskDivPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzDivPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzDivPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskDivPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskDivPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzDivPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzDivPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func divPd(a [8]float64, b [8]float64) [8]float64
TEXT ·divPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskDivPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzDivPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzDivPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskDivPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzDivPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzDivPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskDivPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskDivPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzDivPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzDivPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func divPs(a [16]float32, b [16]float32) [16]float32
TEXT ·divPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskDivPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzDivPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzDivPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·divRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskDivRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzDivRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskzDivRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·divRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskDivRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskDivRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzDivRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskzDivRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func divRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·divRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskDivRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskDivRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzDivRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzDivRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func divRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·divRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskDivRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskDivRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzDivRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzDivRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskDivSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskDivSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzDivSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzDivSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskDivSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskDivSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzDivSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzDivSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func erfPd(a [8]float64) [8]float64
TEXT ·erfPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskErfPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfPs(a [16]float32) [16]float32
TEXT ·erfPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskErfPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfcPd(a [8]float64) [8]float64
TEXT ·erfcPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfcPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskErfcPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfcPs(a [16]float32) [16]float32
TEXT ·erfcPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfcPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskErfcPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfcinvPd(a [8]float64) [8]float64
TEXT ·erfcinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfcinvPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskErfcinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfcinvPs(a [16]float32) [16]float32
TEXT ·erfcinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfcinvPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskErfcinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfinvPd(a [8]float64) [8]float64
TEXT ·erfinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfinvPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskErfinvPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func erfinvPs(a [16]float32) [16]float32
TEXT ·erfinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskErfinvPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskErfinvPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func expPd(a [8]float64) [8]float64
TEXT ·expPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskExpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func expPs(a [16]float32) [16]float32
TEXT ·expPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskExpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func exp10Pd(a [8]float64) [8]float64
TEXT ·exp10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExp10Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskExp10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func exp10Ps(a [16]float32) [16]float32
TEXT ·exp10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExp10Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskExp10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func exp2Pd(a [8]float64) [8]float64
TEXT ·exp2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExp2Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskExp2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func exp2Ps(a [16]float32) [16]float32
TEXT ·exp2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExp2Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskExp2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskExpandEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzExpandEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzExpandEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskExpandEpi321(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskExpandEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandEpi321(k uint8, a [32]byte) [32]byte
TEXT ·maskzExpandEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandEpi322(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·maskExpandEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandEpi322(k uint16, a [64]byte) [64]byte
TEXT ·maskzExpandEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskExpandEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzExpandEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzExpandEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskExpandEpi641(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskExpandEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandEpi641(k uint8, a [32]byte) [32]byte
TEXT ·maskzExpandEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandEpi642(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·maskExpandEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandEpi642(k uint8, a [64]byte) [64]byte
TEXT ·maskzExpandEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskExpandPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzExpandPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzExpandPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskExpandPd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskExpandPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandPd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzExpandPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandPd2(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskExpandPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandPd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzExpandPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskExpandPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzExpandPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzExpandPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskExpandPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskExpandPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzExpandPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandPs2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskExpandPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzExpandPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandloaduEpi32(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskExpandloaduEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExpandloaduEpi32(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzExpandloaduEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExpandloaduEpi321(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·maskExpandloaduEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandloaduEpi321(k uint8, mem_addr uintptr) [32]byte
TEXT ·maskzExpandloaduEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandloaduEpi322(src [64]byte, k uint16, mem_addr uintptr) [64]byte
TEXT ·maskExpandloaduEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandloaduEpi322(k uint16, mem_addr uintptr) [64]byte
TEXT ·maskzExpandloaduEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandloaduEpi64(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskExpandloaduEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExpandloaduEpi64(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzExpandloaduEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExpandloaduEpi641(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·maskExpandloaduEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandloaduEpi641(k uint8, mem_addr uintptr) [32]byte
TEXT ·maskzExpandloaduEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandloaduEpi642(src [64]byte, k uint8, mem_addr uintptr) [64]byte
TEXT ·maskExpandloaduEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandloaduEpi642(k uint8, mem_addr uintptr) [64]byte
TEXT ·maskzExpandloaduEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandloaduPd(src [2]float64, k uint8, mem_addr uintptr) [2]float64
TEXT ·maskExpandloaduPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExpandloaduPd(k uint8, mem_addr uintptr) [2]float64
TEXT ·maskzExpandloaduPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExpandloaduPd1(src [4]float64, k uint8, mem_addr uintptr) [4]float64
TEXT ·maskExpandloaduPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandloaduPd1(k uint8, mem_addr uintptr) [4]float64
TEXT ·maskzExpandloaduPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandloaduPd2(src [8]float64, k uint8, mem_addr uintptr) [8]float64
TEXT ·maskExpandloaduPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandloaduPd2(k uint8, mem_addr uintptr) [8]float64
TEXT ·maskzExpandloaduPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpandloaduPs(src [4]float32, k uint8, mem_addr uintptr) [4]float32
TEXT ·maskExpandloaduPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExpandloaduPs(k uint8, mem_addr uintptr) [4]float32
TEXT ·maskzExpandloaduPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExpandloaduPs1(src [8]float32, k uint8, mem_addr uintptr) [8]float32
TEXT ·maskExpandloaduPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExpandloaduPs1(k uint8, mem_addr uintptr) [8]float32
TEXT ·maskzExpandloaduPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExpandloaduPs2(src [16]float32, k uint16, mem_addr uintptr) [16]float32
TEXT ·maskExpandloaduPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzExpandloaduPs2(k uint16, mem_addr uintptr) [16]float32
TEXT ·maskzExpandloaduPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func expm1Pd(a [8]float64) [8]float64
TEXT ·expm1Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpm1Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskExpm1Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func expm1Ps(a [16]float32) [16]float32
TEXT ·expm1Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskExpm1Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskExpm1Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func extractf32x4Ps(a [8]float32, imm8 int) [4]float32
TEXT ·extractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtractf32x4Ps(src [4]float32, k uint8, a [8]float32, imm8 int) [4]float32
TEXT ·maskExtractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtractf32x4Ps(k uint8, a [8]float32, imm8 int) [4]float32
TEXT ·maskzExtractf32x4Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extractf32x4Ps1(a [16]float32, imm8 int) [4]float32
TEXT ·extractf32x4Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtractf32x4Ps1(src [4]float32, k uint8, a [16]float32, imm8 int) [4]float32
TEXT ·maskExtractf32x4Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtractf32x4Ps1(k uint8, a [16]float32, imm8 int) [4]float32
TEXT ·maskzExtractf32x4Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extractf64x4Pd(a [8]float64, imm8 int) [4]float64
TEXT ·extractf64x4Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExtractf64x4Pd(src [4]float64, k uint8, a [8]float64, imm8 int) [4]float64
TEXT ·maskExtractf64x4Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExtractf64x4Pd(k uint8, a [8]float64, imm8 int) [4]float64
TEXT ·maskzExtractf64x4Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func extracti32x4Epi32(a [32]byte, imm8 int) [16]byte
TEXT ·extracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtracti32x4Epi32(src [16]byte, k uint8, a [32]byte, imm8 int) [16]byte
TEXT ·maskExtracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtracti32x4Epi32(k uint8, a [32]byte, imm8 int) [16]byte
TEXT ·maskzExtracti32x4Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extracti32x4Epi321(a [64]byte, imm8 int) [16]byte
TEXT ·extracti32x4Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExtracti32x4Epi321(src [16]byte, k uint8, a [64]byte, imm8 int) [16]byte
TEXT ·maskExtracti32x4Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExtracti32x4Epi321(k uint8, a [64]byte, imm8 int) [16]byte
TEXT ·maskzExtracti32x4Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func extracti64x4Epi64(a [64]byte, imm8 int) [32]byte
TEXT ·extracti64x4Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskExtracti64x4Epi64(src [32]byte, k uint8, a [64]byte, imm8 int) [32]byte
TEXT ·maskExtracti64x4Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzExtracti64x4Epi64(k uint8, a [64]byte, imm8 int) [32]byte
TEXT ·maskzExtracti64x4Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fixupimmPd(a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·fixupimmPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmPd(a [2]float64, k uint8, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskFixupimmPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmPd(k uint8, a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskzFixupimmPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func fixupimmPd1(a [4]float64, b [4]float64, c [32]byte, imm8 int) [4]float64
TEXT ·fixupimmPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskFixupimmPd1(a [4]float64, k uint8, b [4]float64, c [32]byte, imm8 int) [4]float64
TEXT ·maskFixupimmPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFixupimmPd1(k uint8, a [4]float64, b [4]float64, c [32]byte, imm8 int) [4]float64
TEXT ·maskzFixupimmPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fixupimmPd2(a [8]float64, b [8]float64, c [64]byte, imm8 int) [8]float64
TEXT ·fixupimmPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFixupimmPd2(a [8]float64, k uint8, b [8]float64, c [64]byte, imm8 int) [8]float64
TEXT ·maskFixupimmPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFixupimmPd2(k uint8, a [8]float64, b [8]float64, c [64]byte, imm8 int) [8]float64
TEXT ·maskzFixupimmPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fixupimmPs(a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·fixupimmPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmPs(a [4]float32, k uint8, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskFixupimmPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmPs(k uint8, a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskzFixupimmPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func fixupimmPs1(a [8]float32, b [8]float32, c [32]byte, imm8 int) [8]float32
TEXT ·fixupimmPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskFixupimmPs1(a [8]float32, k uint8, b [8]float32, c [32]byte, imm8 int) [8]float32
TEXT ·maskFixupimmPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFixupimmPs1(k uint8, a [8]float32, b [8]float32, c [32]byte, imm8 int) [8]float32
TEXT ·maskzFixupimmPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fixupimmPs2(a [16]float32, b [16]float32, c [64]byte, imm8 int) [16]float32
TEXT ·fixupimmPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFixupimmPs2(a [16]float32, k uint16, b [16]float32, c [64]byte, imm8 int) [16]float32
TEXT ·maskFixupimmPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFixupimmPs2(k uint16, a [16]float32, b [16]float32, c [64]byte, imm8 int) [16]float32
TEXT ·maskzFixupimmPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fixupimmRoundPd(a [8]float64, b [8]float64, c [64]byte, imm8 int, rounding int) [8]float64
TEXT ·fixupimmRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFixupimmRoundPd(a [8]float64, k uint8, b [8]float64, c [64]byte, imm8 int, rounding int) [8]float64
TEXT ·maskFixupimmRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFixupimmRoundPd(k uint8, a [8]float64, b [8]float64, c [64]byte, imm8 int, rounding int) [8]float64
TEXT ·maskzFixupimmRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fixupimmRoundPs(a [16]float32, b [16]float32, c [64]byte, imm8 int, rounding int) [16]float32
TEXT ·fixupimmRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFixupimmRoundPs(a [16]float32, k uint16, b [16]float32, c [64]byte, imm8 int, rounding int) [16]float32
TEXT ·maskFixupimmRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFixupimmRoundPs(k uint16, a [16]float32, b [16]float32, c [64]byte, imm8 int, rounding int) [16]float32
TEXT ·maskzFixupimmRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fixupimmRoundSd(a [2]float64, b [2]float64, c [16]byte, imm8 int, rounding int) [2]float64
TEXT ·fixupimmRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11
	MOVQ rounding+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func maskFixupimmRoundSd(a [2]float64, k uint8, b [2]float64, c [16]byte, imm8 int, rounding int) [2]float64
TEXT ·maskFixupimmRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzFixupimmRoundSd(k uint8, a [2]float64, b [2]float64, c [16]byte, imm8 int, rounding int) [2]float64
TEXT ·maskzFixupimmRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func fixupimmRoundSs(a [4]float32, b [4]float32, c [16]byte, imm8 int, rounding int) [4]float32
TEXT ·fixupimmRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11
	MOVQ rounding+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func maskFixupimmRoundSs(a [4]float32, k uint8, b [4]float32, c [16]byte, imm8 int, rounding int) [4]float32
TEXT ·maskFixupimmRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzFixupimmRoundSs(k uint8, a [4]float32, b [4]float32, c [16]byte, imm8 int, rounding int) [4]float32
TEXT ·maskzFixupimmRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func fixupimmSd(a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·fixupimmSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmSd(a [2]float64, k uint8, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskFixupimmSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmSd(k uint8, a [2]float64, b [2]float64, c [16]byte, imm8 int) [2]float64
TEXT ·maskzFixupimmSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func fixupimmSs(a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·fixupimmSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmSs(a [4]float32, k uint8, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskFixupimmSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmSs(k uint8, a [4]float32, b [4]float32, c [16]byte, imm8 int) [4]float32
TEXT ·maskzFixupimmSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func floorPd(a [8]float64) [8]float64
TEXT ·floorPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFloorPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskFloorPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func floorPs(a [16]float32) [16]float32
TEXT ·floorPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFloorPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskFloorPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmaddPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmaddPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmaddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmaddPd1(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·maskFmaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmaddPd1(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·mask3FmaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmaddPd1(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·maskzFmaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmaddPd2(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·maskzFmaddPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmaddPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmaddPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmaddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmaddPs1(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·maskFmaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmaddPs1(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·mask3FmaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmaddPs1(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·maskzFmaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmaddPs2(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·maskzFmaddPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmaddRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskzFmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmaddRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskzFmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FmaddRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFmaddRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFmaddRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFmaddRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FmaddRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFmaddRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFmaddRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFmaddSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmaddSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmaddSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmaddSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmaddSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmaddSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmaddSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmaddSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmaddsubPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmaddsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmaddsubPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmaddsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmaddsubPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmaddsubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmaddsubPd1(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·maskFmaddsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmaddsubPd1(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·mask3FmaddsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmaddsubPd1(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·maskzFmaddsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmaddsubPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·fmaddsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddsubPd2(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·maskFmaddsubPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmaddsubPd2(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·mask3FmaddsubPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmaddsubPd2(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·maskzFmaddsubPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddsubPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmaddsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmaddsubPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmaddsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmaddsubPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmaddsubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmaddsubPs1(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·maskFmaddsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmaddsubPs1(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·mask3FmaddsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmaddsubPs1(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·maskzFmaddsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmaddsubPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·fmaddsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddsubPs2(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·maskFmaddsubPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmaddsubPs2(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·mask3FmaddsubPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmaddsubPs2(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·maskzFmaddsubPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fmaddsubRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·fmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddsubRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskFmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmaddsubRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·mask3FmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmaddsubRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskzFmaddsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fmaddsubRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·fmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmaddsubRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskFmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmaddsubRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·mask3FmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmaddsubRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskzFmaddsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmsubPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmsubPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmsubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmsubPd1(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·maskFmsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmsubPd1(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·mask3FmsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmsubPd1(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·maskzFmsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmsubPd2(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·maskzFmsubPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmsubPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmsubPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmsubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmsubPs1(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·maskFmsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmsubPs1(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·mask3FmsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmsubPs1(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·maskzFmsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmsubPs2(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·maskzFmsubPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmsubRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskzFmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmsubRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskzFmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FmsubRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFmsubRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFmsubRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFmsubRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FmsubRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFmsubRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFmsubRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFmsubSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmsubSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmsubSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmsubSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmsubSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmsubSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmsubSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmsubSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmsubaddPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFmsubaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmsubaddPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FmsubaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmsubaddPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFmsubaddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmsubaddPd1(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·maskFmsubaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmsubaddPd1(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·mask3FmsubaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmsubaddPd1(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·maskzFmsubaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmsubaddPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·fmsubaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubaddPd2(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·maskFmsubaddPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmsubaddPd2(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·mask3FmsubaddPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmsubaddPd2(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·maskzFmsubaddPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubaddPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFmsubaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FmsubaddPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FmsubaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFmsubaddPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFmsubaddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFmsubaddPs1(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·maskFmsubaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FmsubaddPs1(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·mask3FmsubaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFmsubaddPs1(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·maskzFmsubaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func fmsubaddPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·fmsubaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubaddPs2(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·maskFmsubaddPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmsubaddPs2(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·mask3FmsubaddPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmsubaddPs2(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·maskzFmsubaddPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fmsubaddRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·fmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubaddRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskFmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmsubaddRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·mask3FmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmsubaddRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskzFmsubaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func fmsubaddRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·fmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFmsubaddRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskFmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask3FmsubaddRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·mask3FmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFmsubaddRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskzFmsubaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFnmaddPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmaddPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmaddPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmaddPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmaddPd1(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·maskFnmaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FnmaddPd1(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·mask3FnmaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmaddPd1(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·maskzFnmaddPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmaddPd2(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·maskzFnmaddPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFnmaddPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmaddPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmaddPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmaddPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmaddPs1(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·maskFnmaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FnmaddPs1(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·mask3FnmaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmaddPs1(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·maskzFnmaddPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmaddPs2(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·maskzFnmaddPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFnmaddRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskzFnmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFnmaddRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskzFnmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFnmaddRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFnmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FnmaddRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FnmaddRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFnmaddRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFnmaddRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFnmaddRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFnmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FnmaddRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FnmaddRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFnmaddRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFnmaddRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFnmaddSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmaddSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmaddSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmaddSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmaddSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmaddSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmaddSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmaddSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmsubPd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmsubPd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmsubPd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmsubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmsubPd1(a [4]float64, k uint8, b [4]float64, c [4]float64) [4]float64
TEXT ·maskFnmsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FnmsubPd1(a [4]float64, b [4]float64, c [4]float64, k uint8) [4]float64
TEXT ·mask3FnmsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmsubPd1(k uint8, a [4]float64, b [4]float64, c [4]float64) [4]float64
TEXT ·maskzFnmsubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmsubPd2(k uint8, a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·maskzFnmsubPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFnmsubPs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmsubPs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmsubPs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmsubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmsubPs1(a [8]float32, k uint8, b [8]float32, c [8]float32) [8]float32
TEXT ·maskFnmsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask3FnmsubPs1(a [8]float32, b [8]float32, c [8]float32, k uint8) [8]float32
TEXT ·mask3FnmsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmsubPs1(k uint8, a [8]float32, b [8]float32, c [8]float32) [8]float32
TEXT ·maskzFnmsubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzFnmsubPs2(k uint16, a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·maskzFnmsubPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFnmsubRoundPd(k uint8, a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·maskzFnmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzFnmsubRoundPs(k uint16, a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·maskzFnmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskFnmsubRoundSd(a [2]float64, k uint8, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskFnmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FnmsubRoundSd(a [2]float64, b [2]float64, c [2]float64, k uint8, rounding int) [2]float64
TEXT ·mask3FnmsubRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFnmsubRoundSd(k uint8, a [2]float64, b [2]float64, c [2]float64, rounding int) [2]float64
TEXT ·maskzFnmsubRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFnmsubRoundSs(a [4]float32, k uint8, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskFnmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mask3FnmsubRoundSs(a [4]float32, b [4]float32, c [4]float32, k uint8, rounding int) [4]float32
TEXT ·mask3FnmsubRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFnmsubRoundSs(k uint8, a [4]float32, b [4]float32, c [4]float32, rounding int) [4]float32
TEXT ·maskzFnmsubRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskFnmsubSd(a [2]float64, k uint8, b [2]float64, c [2]float64) [2]float64
TEXT ·maskFnmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmsubSd(a [2]float64, b [2]float64, c [2]float64, k uint8) [2]float64
TEXT ·mask3FnmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmsubSd(k uint8, a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·maskzFnmsubSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskFnmsubSs(a [4]float32, k uint8, b [4]float32, c [4]float32) [4]float32
TEXT ·maskFnmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask3FnmsubSs(a [4]float32, b [4]float32, c [4]float32, k uint8) [4]float32
TEXT ·mask3FnmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVB k+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzFnmsubSs(k uint8, a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·maskzFnmsubSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func getexpPd(a [2]float64) [2]float64
TEXT ·getexpPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskGetexpPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskGetexpPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzGetexpPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzGetexpPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func getexpPd1(a [4]float64) [4]float64
TEXT ·getexpPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskGetexpPd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskGetexpPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetexpPd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzGetexpPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetexpPd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzGetexpPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func getexpPs(a [4]float32) [4]float32
TEXT ·getexpPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskGetexpPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskGetexpPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzGetexpPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzGetexpPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func getexpPs1(a [8]float32) [8]float32
TEXT ·getexpPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskGetexpPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskGetexpPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetexpPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzGetexpPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetexpPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzGetexpPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzGetexpRoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·maskzGetexpRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzGetexpRoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·maskzGetexpRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func getexpRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·getexpRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskGetexpRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskGetexpRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzGetexpRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzGetexpRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func getexpRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·getexpRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskGetexpRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskGetexpRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzGetexpRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzGetexpRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func getexpSd(a [2]float64, b [2]float64) [2]float64
TEXT ·getexpSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskGetexpSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskGetexpSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzGetexpSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzGetexpSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func getexpSs(a [4]float32, b [4]float32) [4]float32
TEXT ·getexpSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskGetexpSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskGetexpSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzGetexpSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzGetexpSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func getmantPd(a [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·getmantPd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskGetmantPd(src [2]float64, k uint8, a [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskGetmantPd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzGetmantPd(k uint8, a [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskzGetmantPd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func getmantPd1(a [4]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float64
TEXT ·getmantPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskGetmantPd1(src [4]float64, k uint8, a [4]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float64
TEXT ·maskGetmantPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetmantPd1(k uint8, a [4]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float64
TEXT ·maskzGetmantPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetmantPd2(k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float64
TEXT ·maskzGetmantPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func getmantPs(a [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·getmantPs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskGetmantPs(src [4]float32, k uint8, a [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskGetmantPs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzGetmantPs(k uint8, a [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskzGetmantPs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func getmantPs1(a [8]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float32
TEXT ·getmantPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskGetmantPs1(src [8]float32, k uint8, a [8]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float32
TEXT ·maskGetmantPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetmantPs1(k uint8, a [8]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float32
TEXT ·maskzGetmantPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzGetmantPs2(k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [16]float32
TEXT ·maskzGetmantPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzGetmantRoundPd(k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [8]float64
TEXT ·maskzGetmantRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzGetmantRoundPs(k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [16]float32
TEXT ·maskzGetmantRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func getmantRoundSd(a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [2]float64
TEXT ·getmantRoundSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskGetmantRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [2]float64
TEXT ·maskGetmantRoundSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzGetmantRoundSd(k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [2]float64
TEXT ·maskzGetmantRoundSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func getmantRoundSs(a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [4]float32
TEXT ·getmantRoundSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskGetmantRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [4]float32
TEXT ·maskGetmantRoundSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzGetmantRoundSs(k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [4]float32
TEXT ·maskzGetmantRoundSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func getmantSd(a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·getmantSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskGetmantSd(src [2]float64, k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskGetmantSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzGetmantSd(k uint8, a [2]float64, b [2]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [2]float64
TEXT ·maskzGetmantSd(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func getmantSs(a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·getmantSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskGetmantSs(src [4]float32, k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskGetmantSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzGetmantSs(k uint8, a [4]float32, b [4]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [4]float32
TEXT ·maskzGetmantSs(SB),7,$0
	// Unimplemented. Unknown size of type MMMANTISSANORMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func hypotPd(a [8]float64, b [8]float64) [8]float64
TEXT ·hypotPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskHypotPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskHypotPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func hypotPs(a [16]float32, b [16]float32) [16]float32
TEXT ·hypotPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskHypotPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskHypotPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mmaskI32gatherEpi32(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI32gatherEpi321(src [32]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [32]byte
TEXT ·mmaskI32gatherEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mmaskI32gatherEpi64(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI32gatherEpi641(src [32]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [32]byte
TEXT ·mmaskI32gatherEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32gatherEpi64(vindex [32]byte, base_addr uintptr, scale int) [64]byte
TEXT ·i32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskI32gatherEpi64(src [64]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [64]byte
TEXT ·maskI32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mmaskI32gatherPd(src [2]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI32gatherPd1(src [4]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float64
TEXT ·mmaskI32gatherPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32gatherPd(vindex [32]byte, base_addr uintptr, scale int) [8]float64
TEXT ·i32gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskI32gatherPd(src [8]float64, k uint8, vindex [32]byte, base_addr uintptr, scale int) [8]float64
TEXT ·maskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mmaskI32gatherPs(src [4]float32, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI32gatherPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI32gatherPs1(src [8]float32, k uint8, vindex [32]byte, base_addr uintptr, scale int) [8]float32
TEXT ·mmaskI32gatherPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32scatterEpi32(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i32scatterEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterEpi32(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI32scatterEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterEpi321(base_addr uintptr, vindex [32]byte, a [32]byte, scale int) 
TEXT ·i32scatterEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterEpi321(base_addr uintptr, k uint8, vindex [32]byte, a [32]byte, scale int) 
TEXT ·maskI32scatterEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterEpi64(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i32scatterEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterEpi64(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI32scatterEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterEpi641(base_addr uintptr, vindex [16]byte, a [32]byte, scale int) 
TEXT ·i32scatterEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterEpi641(base_addr uintptr, k uint8, vindex [16]byte, a [32]byte, scale int) 
TEXT ·maskI32scatterEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterEpi642(base_addr uintptr, vindex [32]byte, a [64]byte, scale int) 
TEXT ·i32scatterEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterEpi642(base_addr uintptr, k uint8, vindex [32]byte, a [64]byte, scale int) 
TEXT ·maskI32scatterEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPd(base_addr uintptr, vindex [16]byte, a [2]float64, scale int) 
TEXT ·i32scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPd(base_addr uintptr, k uint8, vindex [16]byte, a [2]float64, scale int) 
TEXT ·maskI32scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPd1(base_addr uintptr, vindex [16]byte, a [4]float64, scale int) 
TEXT ·i32scatterPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPd1(base_addr uintptr, k uint8, vindex [16]byte, a [4]float64, scale int) 
TEXT ·maskI32scatterPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPd2(base_addr uintptr, vindex [32]byte, a [8]float64, scale int) 
TEXT ·i32scatterPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPd2(base_addr uintptr, k uint8, vindex [32]byte, a [8]float64, scale int) 
TEXT ·maskI32scatterPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPs(base_addr uintptr, vindex [16]byte, a [4]float32, scale int) 
TEXT ·i32scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPs(base_addr uintptr, k uint8, vindex [16]byte, a [4]float32, scale int) 
TEXT ·maskI32scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPs1(base_addr uintptr, vindex [32]byte, a [8]float32, scale int) 
TEXT ·i32scatterPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPs1(base_addr uintptr, k uint8, vindex [32]byte, a [8]float32, scale int) 
TEXT ·maskI32scatterPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func mmaskI64gatherEpi32(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI64gatherEpi321(src [16]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherEpi32(vindex [64]byte, base_addr uintptr, scale int) [32]byte
TEXT ·i64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskI64gatherEpi32(src [32]byte, k uint8, vindex [64]byte, base_addr uintptr, scale int) [32]byte
TEXT ·maskI64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mmaskI64gatherEpi64(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI64gatherEpi641(src [32]byte, k uint8, vindex [32]byte, base_addr uintptr, scale int) [32]byte
TEXT ·mmaskI64gatherEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i64gatherEpi64(vindex [64]byte, base_addr uintptr, scale int) [64]byte
TEXT ·i64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskI64gatherEpi64(src [64]byte, k uint8, vindex [64]byte, base_addr uintptr, scale int) [64]byte
TEXT ·maskI64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mmaskI64gatherPd(src [2]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI64gatherPd1(src [4]float64, k uint8, vindex [32]byte, base_addr uintptr, scale int) [4]float64
TEXT ·mmaskI64gatherPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i64gatherPd(vindex [64]byte, base_addr uintptr, scale int) [8]float64
TEXT ·i64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskI64gatherPd(src [8]float64, k uint8, vindex [64]byte, base_addr uintptr, scale int) [8]float64
TEXT ·maskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mmaskI64gatherPs(src [4]float32, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func mmaskI64gatherPs1(src [4]float32, k uint8, vindex [32]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI64gatherPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherPs(vindex [64]byte, base_addr uintptr, scale int) [8]float32
TEXT ·i64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskI64gatherPs(src [8]float32, k uint8, vindex [64]byte, base_addr uintptr, scale int) [8]float32
TEXT ·maskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i64scatterEpi32(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i64scatterEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterEpi32(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI64scatterEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterEpi321(base_addr uintptr, vindex [32]byte, a [16]byte, scale int) 
TEXT ·i64scatterEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterEpi321(base_addr uintptr, k uint8, vindex [32]byte, a [16]byte, scale int) 
TEXT ·maskI64scatterEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterEpi322(base_addr uintptr, vindex [64]byte, a [32]byte, scale int) 
TEXT ·i64scatterEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterEpi322(base_addr uintptr, k uint8, vindex [64]byte, a [32]byte, scale int) 
TEXT ·maskI64scatterEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterEpi64(base_addr uintptr, vindex [16]byte, a [16]byte, scale int) 
TEXT ·i64scatterEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterEpi64(base_addr uintptr, k uint8, vindex [16]byte, a [16]byte, scale int) 
TEXT ·maskI64scatterEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterEpi641(base_addr uintptr, vindex [32]byte, a [32]byte, scale int) 
TEXT ·i64scatterEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterEpi641(base_addr uintptr, k uint8, vindex [32]byte, a [32]byte, scale int) 
TEXT ·maskI64scatterEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterEpi642(base_addr uintptr, vindex [64]byte, a [64]byte, scale int) 
TEXT ·i64scatterEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterEpi642(base_addr uintptr, k uint8, vindex [64]byte, a [64]byte, scale int) 
TEXT ·maskI64scatterEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPd(base_addr uintptr, vindex [16]byte, a [2]float64, scale int) 
TEXT ·i64scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPd(base_addr uintptr, k uint8, vindex [16]byte, a [2]float64, scale int) 
TEXT ·maskI64scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPd1(base_addr uintptr, vindex [32]byte, a [4]float64, scale int) 
TEXT ·i64scatterPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPd1(base_addr uintptr, k uint8, vindex [32]byte, a [4]float64, scale int) 
TEXT ·maskI64scatterPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPd2(base_addr uintptr, vindex [64]byte, a [8]float64, scale int) 
TEXT ·i64scatterPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPd2(base_addr uintptr, k uint8, vindex [64]byte, a [8]float64, scale int) 
TEXT ·maskI64scatterPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPs(base_addr uintptr, vindex [16]byte, a [4]float32, scale int) 
TEXT ·i64scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPs(base_addr uintptr, k uint8, vindex [16]byte, a [4]float32, scale int) 
TEXT ·maskI64scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPs1(base_addr uintptr, vindex [32]byte, a [4]float32, scale int) 
TEXT ·i64scatterPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPs1(base_addr uintptr, k uint8, vindex [32]byte, a [4]float32, scale int) 
TEXT ·maskI64scatterPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPs2(base_addr uintptr, vindex [64]byte, a [8]float32, scale int) 
TEXT ·i64scatterPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPs2(base_addr uintptr, k uint8, vindex [64]byte, a [8]float32, scale int) 
TEXT ·maskI64scatterPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func insertf32x4(a [8]float32, b [4]float32, imm8 int) [8]float32
TEXT ·insertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskInsertf32x4(src [8]float32, k uint8, a [8]float32, b [4]float32, imm8 int) [8]float32
TEXT ·maskInsertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzInsertf32x4(k uint8, a [8]float32, b [4]float32, imm8 int) [8]float32
TEXT ·maskzInsertf32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func insertf32x41(a [16]float32, b [4]float32, imm8 int) [16]float32
TEXT ·insertf32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInsertf32x41(src [16]float32, k uint16, a [16]float32, b [4]float32, imm8 int) [16]float32
TEXT ·maskInsertf32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInsertf32x41(k uint16, a [16]float32, b [4]float32, imm8 int) [16]float32
TEXT ·maskzInsertf32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func insertf64x4(a [8]float64, b [4]float64, imm8 int) [8]float64
TEXT ·insertf64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInsertf64x4(src [8]float64, k uint8, a [8]float64, b [4]float64, imm8 int) [8]float64
TEXT ·maskInsertf64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInsertf64x4(k uint8, a [8]float64, b [4]float64, imm8 int) [8]float64
TEXT ·maskzInsertf64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func inserti32x4(a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·inserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskInserti32x4(src [32]byte, k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·maskInserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzInserti32x4(k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·maskzInserti32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func inserti32x41(a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·inserti32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInserti32x41(src [64]byte, k uint16, a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·maskInserti32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInserti32x41(k uint16, a [64]byte, b [16]byte, imm8 int) [64]byte
TEXT ·maskzInserti32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func inserti64x4(a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·inserti64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInserti64x4(src [64]byte, k uint8, a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·maskInserti64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzInserti64x4(k uint8, a [64]byte, b [32]byte, imm8 int) [64]byte
TEXT ·maskzInserti64x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func invsqrtPd(a [8]float64) [8]float64
TEXT ·invsqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInvsqrtPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskInvsqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func invsqrtPs(a [16]float32) [16]float32
TEXT ·invsqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskInvsqrtPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskInvsqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func kand(a uint16, b uint16) uint16
TEXT ·kand(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func kandn(a uint16, b uint16) uint16
TEXT ·kandn(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func kmov(a uint16) uint16
TEXT ·kmov(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVW $0, ret+4(FP)
	RET

// func knot(a uint16) uint16
TEXT ·knot(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVW $0, ret+4(FP)
	RET

// func kor(a uint16, b uint16) uint16
TEXT ·kor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func kortestc(k1 uint16, k2 uint16) int
TEXT ·kortestc(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func kortestz(k1 uint16, k2 uint16) int
TEXT ·kortestz(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func kunpackb(a uint16, b uint16) uint16
TEXT ·kunpackb(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func kxnor(a uint16, b uint16) uint16
TEXT ·kxnor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func kxor(a uint16, b uint16) uint16
TEXT ·kxor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func maskLoadEpi32(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoadEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadEpi32(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoadEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadEpi321(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·maskLoadEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadEpi321(k uint8, mem_addr uintptr) [32]byte
TEXT ·maskzLoadEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadEpi322(k uint16, mem_addr uintptr) [64]byte
TEXT ·maskzLoadEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoadEpi64(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoadEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadEpi64(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoadEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadEpi641(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·maskLoadEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadEpi641(k uint8, mem_addr uintptr) [32]byte
TEXT ·maskzLoadEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadEpi642(k uint8, mem_addr uintptr) [64]byte
TEXT ·maskzLoadEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoadPd(src [2]float64, k uint8, mem_addr uintptr) [2]float64
TEXT ·maskLoadPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadPd(k uint8, mem_addr uintptr) [2]float64
TEXT ·maskzLoadPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadPd1(src [4]float64, k uint8, mem_addr uintptr) [4]float64
TEXT ·maskLoadPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadPd1(k uint8, mem_addr uintptr) [4]float64
TEXT ·maskzLoadPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadPd2(k uint8, mem_addr uintptr) [8]float64
TEXT ·maskzLoadPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoadPs(src [4]float32, k uint8, mem_addr uintptr) [4]float32
TEXT ·maskLoadPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadPs(k uint8, mem_addr uintptr) [4]float32
TEXT ·maskzLoadPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadPs1(src [8]float32, k uint8, mem_addr uintptr) [8]float32
TEXT ·maskLoadPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadPs1(k uint8, mem_addr uintptr) [8]float32
TEXT ·maskzLoadPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoadPs2(k uint16, mem_addr uintptr) [16]float32
TEXT ·maskzLoadPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoadSd(src [2]float64, k uint8, mem_addr float64) [2]float64
TEXT ·maskLoadSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ mem_addr+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzLoadSd(k uint8, mem_addr float64) [2]float64
TEXT ·maskzLoadSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ mem_addr+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func maskLoadSs(src [4]float32, k uint8, mem_addr float32) [4]float32
TEXT ·maskLoadSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVL mem_addr+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskzLoadSs(k uint8, mem_addr float32) [4]float32
TEXT ·maskzLoadSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVL mem_addr+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func maskLoaduEpi32(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoaduEpi32(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoaduEpi321(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·maskLoaduEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoaduEpi321(k uint8, mem_addr uintptr) [32]byte
TEXT ·maskzLoaduEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskLoaduEpi322(src [64]byte, k uint16, mem_addr uintptr) [64]byte
TEXT ·maskLoaduEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLoaduEpi322(k uint16, mem_addr uintptr) [64]byte
TEXT ·maskzLoaduEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoaduEpi64(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoaduEpi64(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoaduEpi641(src [32]byte, k uint8, mem_addr uintptr) [32]byte
TEXT ·maskLoaduEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoaduEpi641(k uint8, mem_addr uintptr) [32]byte
TEXT ·maskzLoaduEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskLoaduEpi642(src [64]byte, k uint8, mem_addr uintptr) [64]byte
TEXT ·maskLoaduEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLoaduEpi642(k uint8, mem_addr uintptr) [64]byte
TEXT ·maskzLoaduEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoaduPd(src [2]float64, k uint8, mem_addr uintptr) [2]float64
TEXT ·maskLoaduPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoaduPd(k uint8, mem_addr uintptr) [2]float64
TEXT ·maskzLoaduPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoaduPd1(src [4]float64, k uint8, mem_addr uintptr) [4]float64
TEXT ·maskLoaduPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoaduPd1(k uint8, mem_addr uintptr) [4]float64
TEXT ·maskzLoaduPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func loaduPd(mem_addr uintptr) [8]float64
TEXT ·loaduPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoaduPd2(src [8]float64, k uint8, mem_addr uintptr) [8]float64
TEXT ·maskLoaduPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLoaduPd2(k uint8, mem_addr uintptr) [8]float64
TEXT ·maskzLoaduPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoaduPs(src [4]float32, k uint8, mem_addr uintptr) [4]float32
TEXT ·maskLoaduPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoaduPs(k uint8, mem_addr uintptr) [4]float32
TEXT ·maskzLoaduPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoaduPs1(src [8]float32, k uint8, mem_addr uintptr) [8]float32
TEXT ·maskLoaduPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzLoaduPs1(k uint8, mem_addr uintptr) [8]float32
TEXT ·maskzLoaduPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func loaduPs(mem_addr uintptr) [16]float32
TEXT ·loaduPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLoaduPs2(src [16]float32, k uint16, mem_addr uintptr) [16]float32
TEXT ·maskLoaduPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzLoaduPs2(k uint16, mem_addr uintptr) [16]float32
TEXT ·maskzLoaduPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func loaduSi512(mem_addr uintptr) [64]byte
TEXT ·loaduSi512(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func logPd(a [8]float64) [8]float64
TEXT ·logPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLogPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskLogPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func logPs(a [16]float32) [16]float32
TEXT ·logPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLogPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskLogPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func log10Pd(a [8]float64) [8]float64
TEXT ·log10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLog10Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskLog10Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func log10Ps(a [16]float32) [16]float32
TEXT ·log10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLog10Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskLog10Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func log1pPd(a [8]float64) [8]float64
TEXT ·log1pPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLog1pPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskLog1pPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func log1pPs(a [16]float32) [16]float32
TEXT ·log1pPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLog1pPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskLog1pPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func log2Pd(a [8]float64) [8]float64
TEXT ·log2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLog2Pd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskLog2Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func logbPd(a [8]float64) [8]float64
TEXT ·logbPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLogbPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskLogbPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func logbPs(a [16]float32) [16]float32
TEXT ·logbPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskLogbPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskLogbPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMaxEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maxEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaxEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpi641(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaxEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxEpi642(a [64]byte, b [64]byte) [64]byte
TEXT ·maxEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpu32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMaxEpu321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpu321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpu322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxEpu64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMaxEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMaxEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maxEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaxEpu641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMaxEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxEpu641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMaxEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpu641(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxEpu642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMaxEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxEpu642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMaxEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxEpu642(a [64]byte, b [64]byte) [64]byte
TEXT ·maxEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMaxPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMaxPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMaxPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskMaxPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzMaxPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskMaxPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzMaxPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxPd(a [8]float64, b [8]float64) [8]float64
TEXT ·maxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMaxPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMaxPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMaxPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskMaxPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMaxPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzMaxPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMaxPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskMaxPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzMaxPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxPs(a [16]float32, b [16]float32) [16]float32
TEXT ·maxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·maskMaxRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxRoundPd(k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·maskzMaxRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxRoundPd(a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·maxRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·maskMaxRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMaxRoundPs(k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·maskzMaxRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maxRoundPs(a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·maxRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMaxRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskMaxRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzMaxRoundSd(k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskzMaxRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maxRoundSd(a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maxRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMaxRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskMaxRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzMaxRoundSs(k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskzMaxRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maxRoundSs(a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maxRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMaxSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMaxSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMaxSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMaxSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMaxSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMaxSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMinEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMinEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func minEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMinEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpi641(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMinEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minEpi642(a [64]byte, b [64]byte) [64]byte
TEXT ·minEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpu32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMinEpu321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpu321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpu322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinEpu64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMinEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMinEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func minEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMinEpu641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMinEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinEpu641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMinEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpu641(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpu641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinEpu642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMinEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinEpu642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMinEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minEpu642(a [64]byte, b [64]byte) [64]byte
TEXT ·minEpu642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMinPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMinPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMinPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskMinPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzMinPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskMinPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzMinPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minPd(a [8]float64, b [8]float64) [8]float64
TEXT ·minPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMinPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMinPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMinPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskMinPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMinPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzMinPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMinPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskMinPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzMinPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minPs(a [16]float32, b [16]float32) [16]float32
TEXT ·minPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·maskMinRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinRoundPd(k uint8, a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·maskzMinRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minRoundPd(a [8]float64, b [8]float64, sae int) [8]float64
TEXT ·minRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·maskMinRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMinRoundPs(k uint16, a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·maskzMinRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func minRoundPs(a [16]float32, b [16]float32, sae int) [16]float32
TEXT ·minRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMinRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskMinRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzMinRoundSd(k uint8, a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·maskzMinRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func minRoundSd(a [2]float64, b [2]float64, sae int) [2]float64
TEXT ·minRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMinRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskMinRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ sae+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzMinRoundSs(k uint8, a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·maskzMinRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ sae+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func minRoundSs(a [4]float32, b [4]float32, sae int) [4]float32
TEXT ·minRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ sae+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMinSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMinSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMinSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMinSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMinSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMinSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMovEpi32(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskMovEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovEpi32(k uint8, a [16]byte) [16]byte
TEXT ·maskzMovEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMovEpi321(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskMovEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovEpi321(k uint8, a [32]byte) [32]byte
TEXT ·maskzMovEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovEpi322(k uint16, a [64]byte) [64]byte
TEXT ·maskzMovEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMovEpi64(src [16]byte, k uint8, a [16]byte) [16]byte
TEXT ·maskMovEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovEpi64(k uint8, a [16]byte) [16]byte
TEXT ·maskzMovEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMovEpi641(src [32]byte, k uint8, a [32]byte) [32]byte
TEXT ·maskMovEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovEpi641(k uint8, a [32]byte) [32]byte
TEXT ·maskzMovEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovEpi642(k uint8, a [64]byte) [64]byte
TEXT ·maskzMovEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMovPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskMovPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzMovPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMovPd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskMovPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovPd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzMovPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovPd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzMovPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMovPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskMovPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzMovPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMovPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskMovPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzMovPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzMovPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMoveSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMoveSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMoveSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMoveSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMoveSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMoveSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMoveSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMoveSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMovedupPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskMovedupPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovedupPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzMovedupPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMovedupPd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskMovedupPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovedupPd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzMovedupPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMovedupPd2(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskMovedupPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMovedupPd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzMovedupPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func movedupPd(a [8]float64) [8]float64
TEXT ·movedupPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMovehdupPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskMovehdupPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMovehdupPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzMovehdupPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMovehdupPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskMovehdupPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMovehdupPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzMovehdupPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMovehdupPs2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskMovehdupPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMovehdupPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzMovehdupPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func movehdupPs(a [16]float32) [16]float32
TEXT ·movehdupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMoveldupPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskMoveldupPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMoveldupPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzMoveldupPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMoveldupPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskMoveldupPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMoveldupPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzMoveldupPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMoveldupPs2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskMoveldupPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMoveldupPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzMoveldupPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func moveldupPs(a [16]float32) [16]float32
TEXT ·moveldupPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulEpi322(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulEpi322(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·mulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulEpu32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulEpu32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulEpu321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulEpu321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulEpu321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskMulEpu322(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulEpu322(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulEpu322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·mulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMulPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMulPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskMulPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzMulPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzMulPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMulPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMulPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskMulPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzMulPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzMulPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskzMulRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzMulRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskzMulRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskMulRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzMulRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzMulRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func mulRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·mulRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMulRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskMulRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzMulRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzMulRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func mulRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·mulRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMulSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskMulSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzMulSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskMulSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzMulSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulloEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulloEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulloEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulloEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskMulloEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskMulloEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulloEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzMulloEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzMulloEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzMulloEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskMulloxEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskMulloxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mulloxEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·mulloxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskNearbyintPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskNearbyintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func nearbyintPd(a [8]float64) [8]float64
TEXT ·nearbyintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskNearbyintPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskNearbyintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func nearbyintPs(a [16]float32) [16]float32
TEXT ·nearbyintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskOrEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskOrEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOrEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzOrEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskOrEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskOrEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzOrEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzOrEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzOrEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzOrEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskOrEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskOrEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOrEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzOrEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskOrEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskOrEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzOrEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzOrEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzOrEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzOrEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskPermutePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzPermutePd(k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskzPermutePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskPermutePd1(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskPermutePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutePd1(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskzPermutePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutePd2(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskPermutePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutePd2(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskzPermutePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutePd(a [8]float64, imm8 int) [8]float64
TEXT ·permutePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskPermutePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzPermutePs(k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskzPermutePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskPermutePs1(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·maskPermutePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutePs1(k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·maskzPermutePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutePs2(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·maskPermutePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutePs2(k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·maskzPermutePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutePs(a [16]float32, imm8 int) [16]float32
TEXT ·permutePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutevarPd(src [2]float64, k uint8, a [2]float64, b [16]byte) [2]float64
TEXT ·maskPermutevarPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutevarPd(k uint8, a [2]float64, b [16]byte) [2]float64
TEXT ·maskzPermutevarPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskPermutevarPd1(src [4]float64, k uint8, a [4]float64, b [32]byte) [4]float64
TEXT ·maskPermutevarPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutevarPd1(k uint8, a [4]float64, b [32]byte) [4]float64
TEXT ·maskzPermutevarPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutevarPd2(src [8]float64, k uint8, a [8]float64, b [64]byte) [8]float64
TEXT ·maskPermutevarPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutevarPd2(k uint8, a [8]float64, b [64]byte) [8]float64
TEXT ·maskzPermutevarPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutevarPd(a [8]float64, b [64]byte) [8]float64
TEXT ·permutevarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutevarPs(src [4]float32, k uint8, a [4]float32, b [16]byte) [4]float32
TEXT ·maskPermutevarPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutevarPs(k uint8, a [4]float32, b [16]byte) [4]float32
TEXT ·maskzPermutevarPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskPermutevarPs1(src [8]float32, k uint8, a [8]float32, b [32]byte) [8]float32
TEXT ·maskPermutevarPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutevarPs1(k uint8, a [8]float32, b [32]byte) [8]float32
TEXT ·maskzPermutevarPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutevarPs2(src [16]float32, k uint16, a [16]float32, b [64]byte) [16]float32
TEXT ·maskPermutevarPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutevarPs2(k uint16, a [16]float32, b [64]byte) [16]float32
TEXT ·maskzPermutevarPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutevarPs(a [16]float32, b [64]byte) [16]float32
TEXT ·permutevarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexEpi64(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskPermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexEpi64(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskzPermutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·permutexEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexEpi641(src [64]byte, k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·maskPermutexEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexEpi641(k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·maskzPermutexEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexEpi641(a [64]byte, imm8 int) [64]byte
TEXT ·permutexEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexPd(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskPermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexPd(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskzPermutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexPd(a [4]float64, imm8 int) [4]float64
TEXT ·permutexPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexPd1(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskPermutexPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexPd1(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskzPermutexPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexPd1(a [8]float64, imm8 int) [8]float64
TEXT ·permutexPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutex2varEpi32(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varEpi32(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varEpi32(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varEpi32(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varEpi321(a [32]byte, k uint8, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskPermutex2varEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask2Permutex2varEpi321(a [32]byte, idx [32]byte, k uint8, b [32]byte) [32]byte
TEXT ·mask2Permutex2varEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutex2varEpi321(k uint8, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskzPermutex2varEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutex2varEpi321(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·permutex2varEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutex2varEpi322(a [64]byte, k uint16, idx [64]byte, b [64]byte) [64]byte
TEXT ·maskPermutex2varEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask2Permutex2varEpi322(a [64]byte, idx [64]byte, k uint16, b [64]byte) [64]byte
TEXT ·mask2Permutex2varEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutex2varEpi322(k uint16, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·maskzPermutex2varEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutex2varEpi322(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·permutex2varEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutex2varEpi64(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskPermutex2varEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varEpi64(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte
TEXT ·mask2Permutex2varEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varEpi64(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·maskzPermutex2varEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varEpi64(a [16]byte, idx [16]byte, b [16]byte) [16]byte
TEXT ·permutex2varEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varEpi641(a [32]byte, k uint8, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskPermutex2varEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask2Permutex2varEpi641(a [32]byte, idx [32]byte, k uint8, b [32]byte) [32]byte
TEXT ·mask2Permutex2varEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutex2varEpi641(k uint8, a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·maskzPermutex2varEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutex2varEpi641(a [32]byte, idx [32]byte, b [32]byte) [32]byte
TEXT ·permutex2varEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutex2varEpi642(a [64]byte, k uint8, idx [64]byte, b [64]byte) [64]byte
TEXT ·maskPermutex2varEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask2Permutex2varEpi642(a [64]byte, idx [64]byte, k uint8, b [64]byte) [64]byte
TEXT ·mask2Permutex2varEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutex2varEpi642(k uint8, a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·maskzPermutex2varEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutex2varEpi642(a [64]byte, idx [64]byte, b [64]byte) [64]byte
TEXT ·permutex2varEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutex2varPd(a [2]float64, k uint8, idx [16]byte, b [2]float64) [2]float64
TEXT ·maskPermutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varPd(a [2]float64, idx [16]byte, k uint8, b [2]float64) [2]float64
TEXT ·mask2Permutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varPd(k uint8, a [2]float64, idx [16]byte, b [2]float64) [2]float64
TEXT ·maskzPermutex2varPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varPd(a [2]float64, idx [16]byte, b [2]float64) [2]float64
TEXT ·permutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varPd1(a [4]float64, k uint8, idx [32]byte, b [4]float64) [4]float64
TEXT ·maskPermutex2varPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask2Permutex2varPd1(a [4]float64, idx [32]byte, k uint8, b [4]float64) [4]float64
TEXT ·mask2Permutex2varPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutex2varPd1(k uint8, a [4]float64, idx [32]byte, b [4]float64) [4]float64
TEXT ·maskzPermutex2varPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutex2varPd1(a [4]float64, idx [32]byte, b [4]float64) [4]float64
TEXT ·permutex2varPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutex2varPd2(a [8]float64, k uint8, idx [64]byte, b [8]float64) [8]float64
TEXT ·maskPermutex2varPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask2Permutex2varPd2(a [8]float64, idx [64]byte, k uint8, b [8]float64) [8]float64
TEXT ·mask2Permutex2varPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutex2varPd2(k uint8, a [8]float64, idx [64]byte, b [8]float64) [8]float64
TEXT ·maskzPermutex2varPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutex2varPd2(a [8]float64, idx [64]byte, b [8]float64) [8]float64
TEXT ·permutex2varPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutex2varPs(a [4]float32, k uint8, idx [16]byte, b [4]float32) [4]float32
TEXT ·maskPermutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varPs(a [4]float32, idx [16]byte, k uint8, b [4]float32) [4]float32
TEXT ·mask2Permutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varPs(k uint8, a [4]float32, idx [16]byte, b [4]float32) [4]float32
TEXT ·maskzPermutex2varPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varPs(a [4]float32, idx [16]byte, b [4]float32) [4]float32
TEXT ·permutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varPs1(a [8]float32, k uint8, idx [32]byte, b [8]float32) [8]float32
TEXT ·maskPermutex2varPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mask2Permutex2varPs1(a [8]float32, idx [32]byte, k uint8, b [8]float32) [8]float32
TEXT ·mask2Permutex2varPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutex2varPs1(k uint8, a [8]float32, idx [32]byte, b [8]float32) [8]float32
TEXT ·maskzPermutex2varPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutex2varPs1(a [8]float32, idx [32]byte, b [8]float32) [8]float32
TEXT ·permutex2varPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutex2varPs2(a [16]float32, k uint16, idx [64]byte, b [16]float32) [16]float32
TEXT ·maskPermutex2varPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func mask2Permutex2varPs2(a [16]float32, idx [64]byte, k uint16, b [16]float32) [16]float32
TEXT ·mask2Permutex2varPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutex2varPs2(k uint16, a [16]float32, idx [64]byte, b [16]float32) [16]float32
TEXT ·maskzPermutex2varPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutex2varPs2(a [16]float32, idx [64]byte, b [16]float32) [16]float32
TEXT ·permutex2varPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexvarEpi32(src [32]byte, k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskPermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexvarEpi32(k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskzPermutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexvarEpi32(idx [32]byte, a [32]byte) [32]byte
TEXT ·permutexvarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexvarEpi321(src [64]byte, k uint16, idx [64]byte, a [64]byte) [64]byte
TEXT ·maskPermutexvarEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexvarEpi321(k uint16, idx [64]byte, a [64]byte) [64]byte
TEXT ·maskzPermutexvarEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexvarEpi321(idx [64]byte, a [64]byte) [64]byte
TEXT ·permutexvarEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexvarEpi64(src [32]byte, k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskPermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexvarEpi64(k uint8, idx [32]byte, a [32]byte) [32]byte
TEXT ·maskzPermutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexvarEpi64(idx [32]byte, a [32]byte) [32]byte
TEXT ·permutexvarEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexvarEpi641(src [64]byte, k uint8, idx [64]byte, a [64]byte) [64]byte
TEXT ·maskPermutexvarEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexvarEpi641(k uint8, idx [64]byte, a [64]byte) [64]byte
TEXT ·maskzPermutexvarEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexvarEpi641(idx [64]byte, a [64]byte) [64]byte
TEXT ·permutexvarEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexvarPd(src [4]float64, k uint8, idx [32]byte, a [4]float64) [4]float64
TEXT ·maskPermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexvarPd(k uint8, idx [32]byte, a [4]float64) [4]float64
TEXT ·maskzPermutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexvarPd(idx [32]byte, a [4]float64) [4]float64
TEXT ·permutexvarPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexvarPd1(src [8]float64, k uint8, idx [64]byte, a [8]float64) [8]float64
TEXT ·maskPermutexvarPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexvarPd1(k uint8, idx [64]byte, a [8]float64) [8]float64
TEXT ·maskzPermutexvarPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexvarPd1(idx [64]byte, a [8]float64) [8]float64
TEXT ·permutexvarPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPermutexvarPs(src [8]float32, k uint8, idx [32]byte, a [8]float32) [8]float32
TEXT ·maskPermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzPermutexvarPs(k uint8, idx [32]byte, a [8]float32) [8]float32
TEXT ·maskzPermutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutexvarPs(idx [32]byte, a [8]float32) [8]float32
TEXT ·permutexvarPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskPermutexvarPs1(src [16]float32, k uint16, idx [64]byte, a [16]float32) [16]float32
TEXT ·maskPermutexvarPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzPermutexvarPs1(k uint16, idx [64]byte, a [16]float32) [16]float32
TEXT ·maskzPermutexvarPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func permutexvarPs1(idx [64]byte, a [16]float32) [16]float32
TEXT ·permutexvarPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPowPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskPowPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func powPd(a [8]float64, b [8]float64) [8]float64
TEXT ·powPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskPowPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskPowPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func powPs(a [16]float32, b [16]float32) [16]float32
TEXT ·powPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRcp14Pd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskRcp14Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzRcp14Pd(k uint8, a [2]float64) [2]float64
TEXT ·maskzRcp14Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func rcp14Pd(a [2]float64) [2]float64
TEXT ·rcp14Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskRcp14Pd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskRcp14Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRcp14Pd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzRcp14Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rcp14Pd1(a [4]float64) [4]float64
TEXT ·rcp14Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRcp14Pd2(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskRcp14Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRcp14Pd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzRcp14Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rcp14Pd2(a [8]float64) [8]float64
TEXT ·rcp14Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRcp14Ps(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskRcp14Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzRcp14Ps(k uint8, a [4]float32) [4]float32
TEXT ·maskzRcp14Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func rcp14Ps(a [4]float32) [4]float32
TEXT ·rcp14Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskRcp14Ps1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskRcp14Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRcp14Ps1(k uint8, a [8]float32) [8]float32
TEXT ·maskzRcp14Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rcp14Ps1(a [8]float32) [8]float32
TEXT ·rcp14Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRcp14Ps2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskRcp14Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRcp14Ps2(k uint16, a [16]float32) [16]float32
TEXT ·maskzRcp14Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rcp14Ps2(a [16]float32) [16]float32
TEXT ·rcp14Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRcp14Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRcp14Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRcp14Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRcp14Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rcp14Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rcp14Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRcp14Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRcp14Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRcp14Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRcp14Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rcp14Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rcp14Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRecipPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskRecipPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func recipPd(a [8]float64) [8]float64
TEXT ·recipPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRecipPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskRecipPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func recipPs(a [16]float32) [16]float32
TEXT ·recipPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpi16(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRemEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskRemEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpi8(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpu16(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRemEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskRemEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpu64(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func remEpu8(a [64]byte, b [64]byte) [64]byte
TEXT ·remEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRintPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskRintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rintPd(a [8]float64) [8]float64
TEXT ·rintPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRintPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskRintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rintPs(a [16]float32) [16]float32
TEXT ·rintPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRolEpi32(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRolEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRolEpi32(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRolEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func rolEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·rolEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRolEpi321(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskRolEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRolEpi321(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskzRolEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rolEpi321(a [32]byte, imm8 int) [32]byte
TEXT ·rolEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRolEpi322(src [64]byte, k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·maskRolEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRolEpi322(k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·maskzRolEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rolEpi322(a [64]byte, imm8 int) [64]byte
TEXT ·rolEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRolEpi64(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRolEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRolEpi64(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRolEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func rolEpi64(a [16]byte, imm8 int) [16]byte
TEXT ·rolEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRolEpi641(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskRolEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRolEpi641(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskzRolEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rolEpi641(a [32]byte, imm8 int) [32]byte
TEXT ·rolEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRolEpi642(src [64]byte, k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·maskRolEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRolEpi642(k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·maskzRolEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rolEpi642(a [64]byte, imm8 int) [64]byte
TEXT ·rolEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRolvEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRolvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRolvEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRolvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rolvEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·rolvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRolvEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskRolvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRolvEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzRolvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rolvEpi321(a [32]byte, b [32]byte) [32]byte
TEXT ·rolvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRolvEpi322(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskRolvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRolvEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzRolvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rolvEpi322(a [64]byte, b [64]byte) [64]byte
TEXT ·rolvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRolvEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRolvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRolvEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRolvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rolvEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·rolvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRolvEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskRolvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRolvEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzRolvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rolvEpi641(a [32]byte, b [32]byte) [32]byte
TEXT ·rolvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRolvEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskRolvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRolvEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzRolvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rolvEpi642(a [64]byte, b [64]byte) [64]byte
TEXT ·rolvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRorEpi32(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRorEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRorEpi32(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRorEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func rorEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·rorEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRorEpi321(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskRorEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRorEpi321(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskzRorEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rorEpi321(a [32]byte, imm8 int) [32]byte
TEXT ·rorEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRorEpi322(src [64]byte, k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·maskRorEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRorEpi322(k uint16, a [64]byte, imm8 int) [64]byte
TEXT ·maskzRorEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rorEpi322(a [64]byte, imm8 int) [64]byte
TEXT ·rorEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRorEpi64(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskRorEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRorEpi64(k uint8, a [16]byte, imm8 int) [16]byte
TEXT ·maskzRorEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func rorEpi64(a [16]byte, imm8 int) [16]byte
TEXT ·rorEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRorEpi641(src [32]byte, k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskRorEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRorEpi641(k uint8, a [32]byte, imm8 int) [32]byte
TEXT ·maskzRorEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rorEpi641(a [32]byte, imm8 int) [32]byte
TEXT ·rorEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRorEpi642(src [64]byte, k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·maskRorEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRorEpi642(k uint8, a [64]byte, imm8 int) [64]byte
TEXT ·maskzRorEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rorEpi642(a [64]byte, imm8 int) [64]byte
TEXT ·rorEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRorvEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRorvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRorvEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRorvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rorvEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·rorvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRorvEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskRorvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRorvEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzRorvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rorvEpi321(a [32]byte, b [32]byte) [32]byte
TEXT ·rorvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRorvEpi322(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskRorvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRorvEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzRorvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rorvEpi322(a [64]byte, b [64]byte) [64]byte
TEXT ·rorvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRorvEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskRorvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRorvEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzRorvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rorvEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·rorvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRorvEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskRorvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRorvEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzRorvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func rorvEpi641(a [32]byte, b [32]byte) [32]byte
TEXT ·rorvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRorvEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskRorvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRorvEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzRorvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rorvEpi642(a [64]byte, b [64]byte) [64]byte
TEXT ·rorvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRoundscalePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskRoundscalePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRoundscalePd(k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskzRoundscalePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func roundscalePd(a [2]float64, imm8 int) [2]float64
TEXT ·roundscalePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRoundscalePd1(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskRoundscalePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRoundscalePd1(k uint8, a [4]float64, imm8 int) [4]float64
TEXT ·maskzRoundscalePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func roundscalePd1(a [4]float64, imm8 int) [4]float64
TEXT ·roundscalePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRoundscalePd2(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskRoundscalePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRoundscalePd2(k uint8, a [8]float64, imm8 int) [8]float64
TEXT ·maskzRoundscalePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func roundscalePd2(a [8]float64, imm8 int) [8]float64
TEXT ·roundscalePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRoundscalePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskRoundscalePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRoundscalePs(k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskzRoundscalePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func roundscalePs(a [4]float32, imm8 int) [4]float32
TEXT ·roundscalePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRoundscalePs1(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·maskRoundscalePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRoundscalePs1(k uint8, a [8]float32, imm8 int) [8]float32
TEXT ·maskzRoundscalePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func roundscalePs1(a [8]float32, imm8 int) [8]float32
TEXT ·roundscalePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRoundscalePs2(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·maskRoundscalePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRoundscalePs2(k uint16, a [16]float32, imm8 int) [16]float32
TEXT ·maskzRoundscalePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func roundscalePs2(a [16]float32, imm8 int) [16]float32
TEXT ·roundscalePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRoundscaleRoundPd(src [8]float64, k uint8, a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·maskRoundscaleRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRoundscaleRoundPd(k uint8, a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·maskzRoundscaleRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func roundscaleRoundPd(a [8]float64, imm8 int, rounding int) [8]float64
TEXT ·roundscaleRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRoundscaleRoundPs(src [16]float32, k uint16, a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·maskRoundscaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRoundscaleRoundPs(k uint16, a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·maskzRoundscaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func roundscaleRoundPs(a [16]float32, imm8 int, rounding int) [16]float32
TEXT ·roundscaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRoundscaleRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskRoundscaleRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzRoundscaleRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskzRoundscaleRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func roundscaleRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·roundscaleRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskRoundscaleRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskRoundscaleRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzRoundscaleRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskzRoundscaleRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func roundscaleRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·roundscaleRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskRoundscaleSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskRoundscaleSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRoundscaleSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzRoundscaleSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func roundscaleSd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·roundscaleSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRoundscaleSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskRoundscaleSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRoundscaleSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzRoundscaleSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func roundscaleSs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·roundscaleSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRsqrt14Pd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskRsqrt14Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzRsqrt14Pd(k uint8, a [2]float64) [2]float64
TEXT ·maskzRsqrt14Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskRsqrt14Pd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskRsqrt14Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRsqrt14Pd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzRsqrt14Pd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRsqrt14Pd2(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskRsqrt14Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRsqrt14Pd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzRsqrt14Pd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rsqrt14Pd(a [8]float64) [8]float64
TEXT ·rsqrt14Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRsqrt14Ps(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskRsqrt14Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzRsqrt14Ps(k uint8, a [4]float32) [4]float32
TEXT ·maskzRsqrt14Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskRsqrt14Ps1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskRsqrt14Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzRsqrt14Ps1(k uint8, a [8]float32) [8]float32
TEXT ·maskzRsqrt14Ps1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskRsqrt14Ps2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskRsqrt14Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzRsqrt14Ps2(k uint16, a [16]float32) [16]float32
TEXT ·maskzRsqrt14Ps2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func rsqrt14Ps(a [16]float32) [16]float32
TEXT ·rsqrt14Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskRsqrt14Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRsqrt14Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRsqrt14Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRsqrt14Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rsqrt14Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rsqrt14Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRsqrt14Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRsqrt14Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRsqrt14Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRsqrt14Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rsqrt14Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rsqrt14Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskScalefPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskScalefPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzScalefPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzScalefPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func scalefPd(a [2]float64, b [2]float64) [2]float64
TEXT ·scalefPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskScalefPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskScalefPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzScalefPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzScalefPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func scalefPd1(a [4]float64, b [4]float64) [4]float64
TEXT ·scalefPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskScalefPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskScalefPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzScalefPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzScalefPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func scalefPd2(a [8]float64, b [8]float64) [8]float64
TEXT ·scalefPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskScalefPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskScalefPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzScalefPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzScalefPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func scalefPs(a [4]float32, b [4]float32) [4]float32
TEXT ·scalefPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskScalefPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskScalefPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzScalefPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzScalefPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func scalefPs1(a [8]float32, b [8]float32) [8]float32
TEXT ·scalefPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskScalefPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskScalefPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzScalefPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzScalefPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func scalefPs2(a [16]float32, b [16]float32) [16]float32
TEXT ·scalefPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskScalefRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskScalefRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzScalefRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskzScalefRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func scalefRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·scalefRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskScalefRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskScalefRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzScalefRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskzScalefRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func scalefRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·scalefRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskScalefRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskScalefRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzScalefRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzScalefRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func scalefRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·scalefRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskScalefRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskScalefRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzScalefRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzScalefRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func scalefRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·scalefRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskScalefSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskScalefSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzScalefSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzScalefSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func scalefSd(a [2]float64, b [2]float64) [2]float64
TEXT ·scalefSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskScalefSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskScalefSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzScalefSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzScalefSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func scalefSs(a [4]float32, b [4]float32) [4]float32
TEXT ·scalefSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func setEpi32(e15 int, e14 int, e13 int, e12 int, e11 int, e10 int, e9 int, e8 int, e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [64]byte
TEXT ·setEpi32(SB),7,$0
	// Unimplemented. Unknown register for type int
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setEpi64(e7 int64, e6 int64, e5 int64, e4 int64, e3 int64, e2 int64, e1 int64, e0 int64) [64]byte
TEXT ·setEpi64(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	//TODO: Code missing

	MOV Z0, ret+64(FP)
	RET

// func setPd(e7 float64, e6 float64, e5 float64, e4 float64, e3 float64, e2 float64, e1 float64, e0 float64) [8]float64
TEXT ·setPd(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	//TODO: Code missing

	MOV Z0, ret+64(FP)
	RET

// func setPs(e15 float32, e14 float32, e13 float32, e12 float32, e11 float32, e10 float32, e9 float32, e8 float32, e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [16]float32
TEXT ·setPs(SB),7,$0
	// Unimplemented. Unknown register for type float32
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func set1Epi16(a int16) [64]byte
TEXT ·set1Epi16(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func maskSet1Epi32(src [16]byte, k uint8, a int) [16]byte
TEXT ·maskSet1Epi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzSet1Epi32(k uint8, a int) [16]byte
TEXT ·maskzSet1Epi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func maskSet1Epi321(src [32]byte, k uint8, a int) [32]byte
TEXT ·maskSet1Epi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSet1Epi321(k uint8, a int) [32]byte
TEXT ·maskzSet1Epi321(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+12(FP)
	RET

// func maskSet1Epi322(src [64]byte, k uint16, a int) [64]byte
TEXT ·maskSet1Epi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSet1Epi322(k uint16, a int) [64]byte
TEXT ·maskzSet1Epi322(SB),7,$0
	MOVW k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOV Z0, ret+12(FP)
	RET

// func set1Epi32(a int) [64]byte
TEXT ·set1Epi32(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+8(FP)
	RET

// func maskSet1Epi64(src [16]byte, k uint8, a int64) [16]byte
TEXT ·maskSet1Epi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzSet1Epi64(k uint8, a int64) [16]byte
TEXT ·maskzSet1Epi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func maskSet1Epi641(src [32]byte, k uint8, a int64) [32]byte
TEXT ·maskSet1Epi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSet1Epi641(k uint8, a int64) [32]byte
TEXT ·maskzSet1Epi641(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOV Y0, ret+12(FP)
	RET

// func maskSet1Epi642(src [64]byte, k uint8, a int64) [64]byte
TEXT ·maskSet1Epi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSet1Epi642(k uint8, a int64) [64]byte
TEXT ·maskzSet1Epi642(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOV Z0, ret+12(FP)
	RET

// func set1Epi64(a int64) [64]byte
TEXT ·set1Epi64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+8(FP)
	RET

// func set1Epi8(a byte) [64]byte
TEXT ·set1Epi8(SB),7,$0
	MOVB a+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func set1Pd(a float64) [8]float64
TEXT ·set1Pd(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+8(FP)
	RET

// func set1Ps(a float32) [16]float32
TEXT ·set1Ps(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOV Z0, ret+4(FP)
	RET

// func set4Epi32(d int, c int, b int, a int) [64]byte
TEXT ·set4Epi32(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	//TODO: Code missing

	MOV Z0, ret+32(FP)
	RET

// func set4Epi64(d int64, c int64, b int64, a int64) [64]byte
TEXT ·set4Epi64(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	//TODO: Code missing

	MOV Z0, ret+32(FP)
	RET

// func set4Pd(d float64, c float64, b float64, a float64) [8]float64
TEXT ·set4Pd(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	//TODO: Code missing

	MOV Z0, ret+32(FP)
	RET

// func set4Ps(d float32, c float32, b float32, a float32) [16]float32
TEXT ·set4Ps(SB),7,$0
	MOVL d+0(FP),R8
	MOVL c+4(FP),R9
	MOVL b+8(FP),R10
	MOVL a+12(FP),R11

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func setrEpi32(e15 int, e14 int, e13 int, e12 int, e11 int, e10 int, e9 int, e8 int, e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [64]byte
TEXT ·setrEpi32(SB),7,$0
	// Unimplemented. Unknown register for type int
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setrEpi64(e7 int64, e6 int64, e5 int64, e4 int64, e3 int64, e2 int64, e1 int64, e0 int64) [64]byte
TEXT ·setrEpi64(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	//TODO: Code missing

	MOV Z0, ret+64(FP)
	RET

// func setrPd(e7 float64, e6 float64, e5 float64, e4 float64, e3 float64, e2 float64, e1 float64, e0 float64) [8]float64
TEXT ·setrPd(SB),7,$0
	MOVQ e7+0(FP),R8
	MOVQ e6+8(FP),R9
	MOVQ e5+16(FP),R10
	MOVQ e4+24(FP),R11
	MOVQ e3+32(FP),R12
	MOVQ e2+40(FP),R13
	MOVQ e1+48(FP),R14
	MOVQ e0+56(FP),R15

	//TODO: Code missing

	MOV Z0, ret+64(FP)
	RET

// func setrPs(e15 float32, e14 float32, e13 float32, e12 float32, e11 float32, e10 float32, e9 float32, e8 float32, e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [16]float32
TEXT ·setrPs(SB),7,$0
	// Unimplemented. Unknown register for type float32
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setr4Epi32(d int, c int, b int, a int) [64]byte
TEXT ·setr4Epi32(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	//TODO: Code missing

	MOV Z0, ret+32(FP)
	RET

// func setr4Epi64(d int64, c int64, b int64, a int64) [64]byte
TEXT ·setr4Epi64(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	//TODO: Code missing

	MOV Z0, ret+32(FP)
	RET

// func setr4Pd(d float64, c float64, b float64, a float64) [8]float64
TEXT ·setr4Pd(SB),7,$0
	MOVQ d+0(FP),R8
	MOVQ c+8(FP),R9
	MOVQ b+16(FP),R10
	MOVQ a+24(FP),R11

	//TODO: Code missing

	MOV Z0, ret+32(FP)
	RET

// func setr4Ps(d float32, c float32, b float32, a float32) [16]float32
TEXT ·setr4Ps(SB),7,$0
	MOVL d+0(FP),R8
	MOVL c+4(FP),R9
	MOVL b+8(FP),R10
	MOVL a+12(FP),R11

	//TODO: Code missing

	MOV Z0, ret+16(FP)
	RET

// func setzero() [16]float32
TEXT ·setzero(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setzeroEpi32() [64]byte
TEXT ·setzeroEpi32(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setzeroPd() [8]float64
TEXT ·setzeroPd(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setzeroPs() [16]float32
TEXT ·setzeroPs(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func setzeroSi512() [64]byte
TEXT ·setzeroSi512(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShuffleEpi32(src [16]byte, k uint8, a [16]byte, imm8 MMPERMENUM) [16]byte
TEXT ·maskShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown size of type MMPERMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzShuffleEpi32(k uint8, a [16]byte, imm8 MMPERMENUM) [16]byte
TEXT ·maskzShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown size of type MMPERMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskShuffleEpi321(src [32]byte, k uint8, a [32]byte, imm8 MMPERMENUM) [32]byte
TEXT ·maskShuffleEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleEpi321(k uint8, a [32]byte, imm8 MMPERMENUM) [32]byte
TEXT ·maskzShuffleEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleEpi322(k uint16, a [64]byte, imm8 MMPERMENUM) [64]byte
TEXT ·maskzShuffleEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShuffleF32x4(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·maskShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleF32x4(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·maskzShuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleF32x4(a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·shuffleF32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShuffleF32x41(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·maskShuffleF32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShuffleF32x41(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·maskzShuffleF32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shuffleF32x41(a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·shuffleF32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShuffleF64x2(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·maskShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleF64x2(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·maskzShuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleF64x2(a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·shuffleF64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShuffleF64x21(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·maskShuffleF64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShuffleF64x21(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·maskzShuffleF64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shuffleF64x21(a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·shuffleF64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShuffleI32x4(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleI32x4(k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskzShuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleI32x4(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·shuffleI32x4(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShuffleI32x41(src [64]byte, k uint16, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskShuffleI32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShuffleI32x41(k uint16, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskzShuffleI32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shuffleI32x41(a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·shuffleI32x41(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShuffleI64x2(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShuffleI64x2(k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskzShuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleI64x2(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·shuffleI64x2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShuffleI64x21(src [64]byte, k uint8, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskShuffleI64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShuffleI64x21(k uint8, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskzShuffleI64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shuffleI64x21(a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·shuffleI64x21(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShufflePd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskShufflePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzShufflePd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzShufflePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskShufflePd1(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·maskShufflePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShufflePd1(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64
TEXT ·maskzShufflePd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShufflePd2(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·maskShufflePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShufflePd2(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·maskzShufflePd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shufflePd(a [8]float64, b [8]float64, imm8 int) [8]float64
TEXT ·shufflePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskShufflePs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskShufflePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzShufflePs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzShufflePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskShufflePs1(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·maskShufflePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzShufflePs1(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32
TEXT ·maskzShufflePs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskShufflePs2(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·maskShufflePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzShufflePs2(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·maskzShufflePs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func shufflePs(a [16]float32, b [16]float32, imm8 int) [16]float32
TEXT ·shufflePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSinPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskSinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sinPd(a [8]float64) [8]float64
TEXT ·sinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSinPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskSinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sinPs(a [16]float32) [16]float32
TEXT ·sinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSincosPd(cos_res [8]float64, sin_src [8]float64, cos_src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskSincosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sincosPd(cos_res [8]float64, a [8]float64) [8]float64
TEXT ·sincosPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSincosPs(cos_res [16]float32, sin_src [16]float32, cos_src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskSincosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sincosPs(cos_res [16]float32, a [16]float32) [16]float32
TEXT ·sincosPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSindPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskSindPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sindPd(a [8]float64) [8]float64
TEXT ·sindPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSindPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskSindPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sindPs(a [16]float32) [16]float32
TEXT ·sindPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSinhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskSinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sinhPd(a [8]float64) [8]float64
TEXT ·sinhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSinhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskSinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sinhPs(a [16]float32) [16]float32
TEXT ·sinhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSllEpi321(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSllEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllEpi321(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSllEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSllEpi322(src [64]byte, k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSllEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSllEpi322(k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSllEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sllEpi32(a [64]byte, count [16]byte) [64]byte
TEXT ·sllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSllEpi641(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSllEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllEpi641(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSllEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSllEpi642(src [64]byte, k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSllEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSllEpi642(k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSllEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sllEpi64(a [64]byte, count [16]byte) [64]byte
TEXT ·sllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSlliEpi32(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSlliEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSlliEpi32(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSlliEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSlliEpi321(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSlliEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSlliEpi321(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSlliEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSlliEpi322(k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSlliEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSlliEpi64(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSlliEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSlliEpi64(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSlliEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSlliEpi641(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSlliEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSlliEpi641(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSlliEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSlliEpi642(src [64]byte, k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskSlliEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSlliEpi642(k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSlliEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func slliEpi64(a [64]byte, imm8 uint32) [64]byte
TEXT ·slliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllvEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllvEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSllvEpi321(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSllvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllvEpi321(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSllvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllvEpi322(k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSllvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSllvEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSllvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllvEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSllvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSllvEpi641(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSllvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSllvEpi641(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSllvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSllvEpi642(src [64]byte, k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·maskSllvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSllvEpi642(k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSllvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sllvEpi64(a [64]byte, count [64]byte) [64]byte
TEXT ·sllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSqrtPd(src [2]float64, k uint8, a [2]float64) [2]float64
TEXT ·maskSqrtPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzSqrtPd(k uint8, a [2]float64) [2]float64
TEXT ·maskzSqrtPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskSqrtPd1(src [4]float64, k uint8, a [4]float64) [4]float64
TEXT ·maskSqrtPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSqrtPd1(k uint8, a [4]float64) [4]float64
TEXT ·maskzSqrtPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSqrtPd2(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskSqrtPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSqrtPd2(k uint8, a [8]float64) [8]float64
TEXT ·maskzSqrtPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sqrtPd(a [8]float64) [8]float64
TEXT ·sqrtPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSqrtPs(src [4]float32, k uint8, a [4]float32) [4]float32
TEXT ·maskSqrtPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzSqrtPs(k uint8, a [4]float32) [4]float32
TEXT ·maskzSqrtPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskSqrtPs1(src [8]float32, k uint8, a [8]float32) [8]float32
TEXT ·maskSqrtPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSqrtPs1(k uint8, a [8]float32) [8]float32
TEXT ·maskzSqrtPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSqrtPs2(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskSqrtPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSqrtPs2(k uint16, a [16]float32) [16]float32
TEXT ·maskzSqrtPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sqrtPs(a [16]float32) [16]float32
TEXT ·sqrtPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSqrtRoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64
TEXT ·maskSqrtRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSqrtRoundPd(k uint8, a [8]float64, rounding int) [8]float64
TEXT ·maskzSqrtRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sqrtRoundPd(a [8]float64, rounding int) [8]float64
TEXT ·sqrtRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSqrtRoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32
TEXT ·maskSqrtRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSqrtRoundPs(k uint16, a [16]float32, rounding int) [16]float32
TEXT ·maskzSqrtRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sqrtRoundPs(a [16]float32, rounding int) [16]float32
TEXT ·sqrtRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSqrtRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskSqrtRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzSqrtRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzSqrtRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func sqrtRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·sqrtRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskSqrtRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskSqrtRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzSqrtRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzSqrtRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func sqrtRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·sqrtRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskSqrtSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskSqrtSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSqrtSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzSqrtSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSqrtSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskSqrtSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSqrtSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzSqrtSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSraEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSraEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSraEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSraEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSraEpi321(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSraEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraEpi321(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSraEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSraEpi322(src [64]byte, k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSraEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSraEpi322(k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSraEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sraEpi32(a [64]byte, count [16]byte) [64]byte
TEXT ·sraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSraEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSraEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSraEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSraEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sraEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sraEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSraEpi641(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSraEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraEpi641(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSraEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sraEpi641(a [32]byte, count [16]byte) [32]byte
TEXT ·sraEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSraEpi642(src [64]byte, k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSraEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSraEpi642(k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSraEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sraEpi642(a [64]byte, count [16]byte) [64]byte
TEXT ·sraEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSraiEpi32(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSraiEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSraiEpi32(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSraiEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSraiEpi321(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSraiEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraiEpi321(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSraiEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraiEpi322(k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSraiEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSraiEpi64(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSraiEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSraiEpi64(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSraiEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func sraiEpi64(a [16]byte, imm8 uint32) [16]byte
TEXT ·sraiEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskSraiEpi641(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSraiEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSraiEpi641(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSraiEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sraiEpi641(a [32]byte, imm8 uint32) [32]byte
TEXT ·sraiEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSraiEpi642(src [64]byte, k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskSraiEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSraiEpi642(k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSraiEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sraiEpi642(a [64]byte, imm8 uint32) [64]byte
TEXT ·sraiEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSravEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSravEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSravEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSravEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSravEpi321(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSravEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSravEpi321(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSravEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSravEpi322(k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSravEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSravEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSravEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSravEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSravEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sravEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sravEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSravEpi641(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSravEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSravEpi641(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSravEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sravEpi641(a [32]byte, count [32]byte) [32]byte
TEXT ·sravEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSravEpi642(src [64]byte, k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·maskSravEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSravEpi642(k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSravEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func sravEpi642(a [64]byte, count [64]byte) [64]byte
TEXT ·sravEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrlEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSrlEpi321(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSrlEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlEpi321(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSrlEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrlEpi322(src [64]byte, k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSrlEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrlEpi322(k uint16, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSrlEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srlEpi32(a [64]byte, count [16]byte) [64]byte
TEXT ·srlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrlEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSrlEpi641(src [32]byte, k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskSrlEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlEpi641(k uint8, a [32]byte, count [16]byte) [32]byte
TEXT ·maskzSrlEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrlEpi642(src [64]byte, k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·maskSrlEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrlEpi642(k uint8, a [64]byte, count [16]byte) [64]byte
TEXT ·maskzSrlEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srlEpi64(a [64]byte, count [16]byte) [64]byte
TEXT ·srlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrliEpi32(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSrliEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrliEpi32(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSrliEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrliEpi321(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSrliEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrliEpi321(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSrliEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrliEpi322(k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSrliEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrliEpi64(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskSrliEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrliEpi64(k uint8, a [16]byte, imm8 uint32) [16]byte
TEXT ·maskzSrliEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrliEpi641(src [32]byte, k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskSrliEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrliEpi641(k uint8, a [32]byte, imm8 uint32) [32]byte
TEXT ·maskzSrliEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrliEpi642(src [64]byte, k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskSrliEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrliEpi642(k uint8, a [64]byte, imm8 uint32) [64]byte
TEXT ·maskzSrliEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srliEpi64(a [64]byte, imm8 uint32) [64]byte
TEXT ·srliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrlvEpi32(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlvEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlvEpi32(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlvEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSrlvEpi321(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSrlvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlvEpi321(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSrlvEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlvEpi322(k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSrlvEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSrlvEpi64(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskSrlvEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlvEpi64(k uint8, a [16]byte, count [16]byte) [16]byte
TEXT ·maskzSrlvEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSrlvEpi641(src [32]byte, k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskSrlvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSrlvEpi641(k uint8, a [32]byte, count [32]byte) [32]byte
TEXT ·maskzSrlvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSrlvEpi642(src [64]byte, k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·maskSrlvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSrlvEpi642(k uint8, a [64]byte, count [64]byte) [64]byte
TEXT ·maskzSrlvEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func srlvEpi64(a [64]byte, count [64]byte) [64]byte
TEXT ·srlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskStoreEpi32(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreEpi321(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskStoreEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreEpi64(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreEpi641(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskStoreEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStorePd(mem_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskStorePd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStorePd1(mem_addr uintptr, k uint8, a [4]float64) 
TEXT ·maskStorePd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStorePs(mem_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskStorePs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStorePs1(mem_addr uintptr, k uint8, a [8]float32) 
TEXT ·maskStorePs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreSd(mem_addr float64, k uint8, a [2]float64) 
TEXT ·maskStoreSd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func maskStoreSs(mem_addr float32, k uint8, a [4]float32) 
TEXT ·maskStoreSs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVB k+4(FP),R9
	MOVOU a+8(FP),X2

	//TODO: Code missing

	RET

// func maskStoreuEpi32(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi321(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskStoreuEpi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi322(mem_addr uintptr, k uint16, a [64]byte) 
TEXT ·maskStoreuEpi322(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi64(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi641(mem_addr uintptr, k uint8, a [32]byte) 
TEXT ·maskStoreuEpi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi642(mem_addr uintptr, k uint8, a [64]byte) 
TEXT ·maskStoreuEpi642(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPd(mem_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskStoreuPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPd1(mem_addr uintptr, k uint8, a [4]float64) 
TEXT ·maskStoreuPd1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPd2(mem_addr uintptr, k uint8, a [8]float64) 
TEXT ·maskStoreuPd2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuPd(mem_addr uintptr, a [8]float64) 
TEXT ·storeuPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPs(mem_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskStoreuPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPs1(mem_addr uintptr, k uint8, a [8]float32) 
TEXT ·maskStoreuPs1(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPs2(mem_addr uintptr, k uint16, a [16]float32) 
TEXT ·maskStoreuPs2(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuPs(mem_addr uintptr, a [16]float32) 
TEXT ·storeuPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi512(mem_addr uintptr, a [64]byte) 
TEXT ·storeuSi512(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func streamLoadSi512(mem_addr uintptr) [64]byte
TEXT ·streamLoadSi512(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func streamPd(mem_addr uintptr, a [8]float64) 
TEXT ·streamPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func streamPs(mem_addr uintptr, a [16]float32) 
TEXT ·streamPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func streamSi512(mem_addr uintptr, a [64]byte) 
TEXT ·streamSi512(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskSubEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSubEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskSubEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzSubEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSubEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskSubEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzSubEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskSubEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskSubEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzSubEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func subEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·subEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskSubPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzSubPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSubPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskSubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzSubPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzSubPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskSubPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzSubPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSubPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskSubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzSubPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzSubPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzSubPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubRoundPd(k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·maskzSubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzSubRoundPs(k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·maskzSubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskSubRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskSubRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzSubRoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzSubRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func subRoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·subRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskSubRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskSubRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzSubRoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzSubRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func subRoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·subRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskSubSd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskSubSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubSd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzSubSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSubSs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskSubSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubSs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzSubSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskSvmlRoundPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskSvmlRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func svmlRoundPd(a [8]float64) [8]float64
TEXT ·svmlRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTanPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskTanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func tanPd(a [8]float64) [8]float64
TEXT ·tanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTanPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskTanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func tanPs(a [16]float32) [16]float32
TEXT ·tanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTandPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskTandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func tandPd(a [8]float64) [8]float64
TEXT ·tandPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTandPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskTandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func tandPs(a [16]float32) [16]float32
TEXT ·tandPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTanhPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskTanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func tanhPd(a [8]float64) [8]float64
TEXT ·tanhPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTanhPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskTanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func tanhPs(a [16]float32) [16]float32
TEXT ·tanhPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTernarylogicEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·maskTernarylogicEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzTernarylogicEpi32(k uint8, a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·maskzTernarylogicEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func ternarylogicEpi32(a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·ternarylogicEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskTernarylogicEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskTernarylogicEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzTernarylogicEpi321(k uint8, a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·maskzTernarylogicEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func ternarylogicEpi321(a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·ternarylogicEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskTernarylogicEpi322(src [64]byte, k uint16, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskTernarylogicEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzTernarylogicEpi322(k uint16, a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·maskzTernarylogicEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func ternarylogicEpi322(a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·ternarylogicEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTernarylogicEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·maskTernarylogicEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzTernarylogicEpi64(k uint8, a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·maskzTernarylogicEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func ternarylogicEpi64(a [16]byte, b [16]byte, c [16]byte, imm8 int) [16]byte
TEXT ·ternarylogicEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskTernarylogicEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·maskTernarylogicEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzTernarylogicEpi641(k uint8, a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·maskzTernarylogicEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func ternarylogicEpi641(a [32]byte, b [32]byte, c [32]byte, imm8 int) [32]byte
TEXT ·ternarylogicEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskTernarylogicEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte, imm8 int) [64]byte
TEXT ·maskTernarylogicEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzTernarylogicEpi642(k uint8, a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·maskzTernarylogicEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func ternarylogicEpi642(a [64]byte, b [64]byte, c [64]byte, imm8 int) [64]byte
TEXT ·ternarylogicEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTestEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·testEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskTestEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func testEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·testEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskTestEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·testEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskTestEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func testEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·testEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskTestEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskTestEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func testEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·testEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskTestnEpi32Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestnEpi32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testnEpi32Mask(a [16]byte, b [16]byte) uint8
TEXT ·testnEpi32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestnEpi32Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskTestnEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func testnEpi32Mask1(a [32]byte, b [32]byte) uint8
TEXT ·testnEpi32Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskTestnEpi32Mask2(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·maskTestnEpi32Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func testnEpi32Mask2(a [64]byte, b [64]byte) uint16
TEXT ·testnEpi32Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func maskTestnEpi64Mask(k1 uint8, a [16]byte, b [16]byte) uint8
TEXT ·maskTestnEpi64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testnEpi64Mask(a [16]byte, b [16]byte) uint8
TEXT ·testnEpi64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestnEpi64Mask1(k1 uint8, a [32]byte, b [32]byte) uint8
TEXT ·maskTestnEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func testnEpi64Mask1(a [32]byte, b [32]byte) uint8
TEXT ·testnEpi64Mask1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskTestnEpi64Mask2(k1 uint8, a [64]byte, b [64]byte) uint8
TEXT ·maskTestnEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func testnEpi64Mask2(a [64]byte, b [64]byte) uint8
TEXT ·testnEpi64Mask2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func maskTruncPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·maskTruncPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func truncPd(a [8]float64) [8]float64
TEXT ·truncPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskTruncPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·maskTruncPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func truncPs(a [16]float32) [16]float32
TEXT ·truncPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func undefined() [16]float32
TEXT ·undefined(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func undefinedEpi32() [64]byte
TEXT ·undefinedEpi32(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func undefinedPd() [8]float64
TEXT ·undefinedPd(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func undefinedPs() [16]float32
TEXT ·undefinedPs(SB),7,$0

	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackhiEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhiEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackhiEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackhiEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackhiEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackhiEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackhiEpi322(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackhiEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackhiEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackhiEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackhiEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackhiEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackhiEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhiEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackhiEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackhiEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackhiEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackhiEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackhiEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackhiEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackhiEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackhiEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackhiEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackhiEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackhiPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskUnpackhiPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhiPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzUnpackhiPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackhiPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskUnpackhiPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackhiPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzUnpackhiPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackhiPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskUnpackhiPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackhiPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzUnpackhiPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackhiPd(a [8]float64, b [8]float64) [8]float64
TEXT ·unpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackhiPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskUnpackhiPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhiPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzUnpackhiPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackhiPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskUnpackhiPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackhiPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzUnpackhiPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackhiPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskUnpackhiPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackhiPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzUnpackhiPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackhiPs(a [16]float32, b [16]float32) [16]float32
TEXT ·unpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackloEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackloEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackloEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackloEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackloEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackloEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackloEpi322(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackloEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackloEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackloEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackloEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackloEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskUnpackloEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackloEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzUnpackloEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackloEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskUnpackloEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackloEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzUnpackloEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackloEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskUnpackloEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackloEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzUnpackloEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackloEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·unpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackloPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskUnpackloPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackloPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzUnpackloPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackloPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskUnpackloPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackloPd1(k uint8, a [4]float64, b [4]float64) [4]float64
TEXT ·maskzUnpackloPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackloPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskUnpackloPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackloPd2(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·maskzUnpackloPd2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackloPd(a [8]float64, b [8]float64) [8]float64
TEXT ·unpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512d
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskUnpackloPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskUnpackloPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackloPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzUnpackloPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskUnpackloPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskUnpackloPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzUnpackloPs1(k uint8, a [8]float32, b [8]float32) [8]float32
TEXT ·maskzUnpackloPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskUnpackloPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskUnpackloPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskzUnpackloPs2(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·maskzUnpackloPs2(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func unpackloPs(a [16]float32, b [16]float32) [16]float32
TEXT ·unpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskXorEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskXorEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXorEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzXorEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskXorEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskXorEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzXorEpi321(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzXorEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzXorEpi322(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzXorEpi322(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func maskXorEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskXorEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXorEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzXorEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskXorEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskXorEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzXorEpi641(k uint8, a [32]byte, b [32]byte) [32]byte
TEXT ·maskzXorEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskzXorEpi642(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·maskzXorEpi642(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type M512i
	//TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

