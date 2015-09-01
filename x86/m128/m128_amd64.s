// func absEpi16(a [16]byte) [16]byte
TEXT ·absEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

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

// func absEpi32(a [16]byte) [16]byte
TEXT ·absEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

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

// func absEpi8(a [16]byte) [16]byte
TEXT ·absEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func acosPd(a [2]float64) [2]float64
TEXT ·acosPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func acosPs(a [4]float32) [4]float32
TEXT ·acosPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func acoshPd(a [2]float64) [2]float64
TEXT ·acoshPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func acoshPs(a [4]float32) [4]float32
TEXT ·acoshPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func addEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	PADDB  X1,X0

	MOVOU X0, ret+32(FP)
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

// func addPd(a [2]float64, b [2]float64) [2]float64
TEXT ·addPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

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

// func addPs(a [4]float32, b [4]float32) [4]float32
TEXT ·addPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	ADDPS X1, X0

	MOVOU X0, ret+32(FP)
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

// func addSd(a [2]float64, b [2]float64) [2]float64
TEXT ·addSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addSs(a [4]float32, b [4]float32) [4]float32
TEXT ·addSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addsEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addsEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addsEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func addsubPd(a [2]float64, b [2]float64) [2]float64
TEXT ·addsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addsubPs(a [4]float32, b [4]float32) [4]float32
TEXT ·addsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesdecSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesdecSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesdeclastSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesdeclastSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesencSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesencSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesenclastSi128(a [16]byte, RoundKey [16]byte) [16]byte
TEXT ·aesenclastSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesimcSi128(a [16]byte) [16]byte
TEXT ·aesimcSi128(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func aeskeygenassistSi128(a [16]byte, imm8 int) [16]byte
TEXT ·aeskeygenassistSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func alignrEpi8(a [16]byte, b [16]byte, count int) [16]byte
TEXT ·alignrEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func andPd(a [2]float64, b [2]float64) [2]float64
TEXT ·andPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAndPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAndPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAndPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func andPs(a [4]float32, b [4]float32) [4]float32
TEXT ·andPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAndPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAndPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAndPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func andSi128(a [16]byte, b [16]byte) [16]byte
TEXT ·andSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func andnotPd(a [2]float64, b [2]float64) [2]float64
TEXT ·andnotPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAndnotPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskAndnotPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnotPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzAndnotPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func andnotPs(a [4]float32, b [4]float32) [4]float32
TEXT ·andnotPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAndnotPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskAndnotPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnotPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzAndnotPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func andnotSi128(a [16]byte, b [16]byte) [16]byte
TEXT ·andnotSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func asinPd(a [2]float64) [2]float64
TEXT ·asinPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func asinPs(a [4]float32) [4]float32
TEXT ·asinPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func asinhPd(a [2]float64) [2]float64
TEXT ·asinhPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func asinhPs(a [4]float32) [4]float32
TEXT ·asinhPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atanPd(a [2]float64) [2]float64
TEXT ·atanPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atanPs(a [4]float32) [4]float32
TEXT ·atanPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atan2Pd(a [2]float64, b [2]float64) [2]float64
TEXT ·atan2Pd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func atan2Ps(a [4]float32, b [4]float32) [4]float32
TEXT ·atan2Ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func atanhPd(a [2]float64) [2]float64
TEXT ·atanhPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atanhPs(a [4]float32) [4]float32
TEXT ·atanhPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func avgEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·avgEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func avgEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·avgEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func blendEpi16(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·blendEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskBlendEpi16(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func blendEpi32(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·blendEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskBlendEpi32(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlendEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlendEpi8(k uint16, a [16]byte, b [16]byte) [16]byte
TEXT ·maskBlendEpi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func blendPd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·blendPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskBlendPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskBlendPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func blendPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·blendPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskBlendPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskBlendPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func blendvEpi8(a [16]byte, b [16]byte, mask [16]byte) [16]byte
TEXT ·blendvEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU mask+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func blendvPd(a [2]float64, b [2]float64, mask [2]float64) [2]float64
TEXT ·blendvPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU mask+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func blendvPs(a [4]float32, b [4]float32, mask [4]float32) [4]float32
TEXT ·blendvPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU mask+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func broadcastSs(mem_addr float32) [4]float32
TEXT ·broadcastSs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastbEpi8(a [16]byte) [16]byte
TEXT ·broadcastbEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func broadcastdEpi32(a [16]byte) [16]byte
TEXT ·broadcastdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func broadcastmbEpi64(k uint8) [16]byte
TEXT ·broadcastmbEpi64(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastmwEpi32(k uint16) [16]byte
TEXT ·broadcastmwEpi32(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastqEpi64(a [16]byte) [16]byte
TEXT ·broadcastqEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func broadcastsdPd(a [2]float64) [2]float64
TEXT ·broadcastsdPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func broadcastssPs(a [4]float32) [4]float32
TEXT ·broadcastssPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func broadcastwEpi16(a [16]byte) [16]byte
TEXT ·broadcastwEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func bslliSi128(a [16]byte, imm8 int) [16]byte
TEXT ·bslliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func bsrliSi128(a [16]byte, imm8 int) [16]byte
TEXT ·bsrliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func castpdPs(a [2]float64) [4]float32
TEXT ·castpdPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castpdSi128(a [2]float64) [16]byte
TEXT ·castpdSi128(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castpsPd(a [4]float32) [2]float64
TEXT ·castpsPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castpsSi128(a [4]float32) [16]byte
TEXT ·castpsSi128(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castsi128Pd(a [16]byte) [2]float64
TEXT ·castsi128Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castsi128Ps(a [16]byte) [4]float32
TEXT ·castsi128Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cbrtPd(a [2]float64) [2]float64
TEXT ·cbrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cbrtPs(a [4]float32) [4]float32
TEXT ·cbrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnormPd(a [2]float64) [2]float64
TEXT ·cdfnormPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnormPs(a [4]float32) [4]float32
TEXT ·cdfnormPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnorminvPd(a [2]float64) [2]float64
TEXT ·cdfnorminvPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnorminvPs(a [4]float32) [4]float32
TEXT ·cdfnorminvPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func ceilPd(a [2]float64) [2]float64
TEXT ·ceilPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func ceilPs(a [4]float32) [4]float32
TEXT ·ceilPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func ceilSd(a [2]float64, b [2]float64) [2]float64
TEXT ·ceilSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func ceilSs(a [4]float32, b [4]float32) [4]float32
TEXT ·ceilSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cexpPs(a [4]float32) [4]float32
TEXT ·cexpPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func clmulepi64Si128(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·clmulepi64Si128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func clogPs(a [4]float32) [4]float32
TEXT ·clogPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cmpPd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·cmpPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func cmpPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·cmpPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func cmpSd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·cmpSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func cmpSs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·cmpSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func cmpeqEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpeqEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpeqEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpeqEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpeqPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpeqPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeqPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpeqPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeqSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpeqSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeqSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpeqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpgePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpgePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgeSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpgeSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgeSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpgtEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpgtEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpgtEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpgtPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpgtPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgtPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpgtSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmplePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmplePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmplePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmplePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpleSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpleSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpleSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpleSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpltEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpltEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpltEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpltEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpltEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpltPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpltPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpltPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpltSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cmpneqPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpneqPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpneqPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpneqPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpneqSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpneqSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpneqSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpneqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpngePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngeSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpngeSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngeSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngtPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpngtPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngtPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngtPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngtSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpngtSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpngtSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnlePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpnlePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnlePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnlePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnleSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpnleSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnleSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnleSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnltPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpnltPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnltPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnltPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnltSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpnltSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpnltSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpordPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpordPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpordPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpordPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpordSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpordSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpordSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpordSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpunordPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpunordPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpunordPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpunordPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpunordSd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpunordSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpunordSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpunordSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func comieqSd(a [2]float64, b [2]float64) int
TEXT ·comieqSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comieqSs(a [4]float32, b [4]float32) int
TEXT ·comieqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comigeSd(a [2]float64, b [2]float64) int
TEXT ·comigeSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comigeSs(a [4]float32, b [4]float32) int
TEXT ·comigeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comigtSd(a [2]float64, b [2]float64) int
TEXT ·comigtSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comigtSs(a [4]float32, b [4]float32) int
TEXT ·comigtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comileSd(a [2]float64, b [2]float64) int
TEXT ·comileSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comileSs(a [4]float32, b [4]float32) int
TEXT ·comileSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comiltSd(a [2]float64, b [2]float64) int
TEXT ·comiltSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comiltSs(a [4]float32, b [4]float32) int
TEXT ·comiltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comineqSd(a [2]float64, b [2]float64) int
TEXT ·comineqSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func comineqSs(a [4]float32, b [4]float32) int
TEXT ·comineqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
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

// func maskCompressstoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCompressstoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuEpi64(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCompressstoreuEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPd(base_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskCompressstoreuPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreuPs(base_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskCompressstoreuPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

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

// func cosPd(a [2]float64) [2]float64
TEXT ·cosPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cosPs(a [4]float32) [4]float32
TEXT ·cosPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cosdPd(a [2]float64) [2]float64
TEXT ·cosdPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cosdPs(a [4]float32) [4]float32
TEXT ·cosdPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func coshPd(a [2]float64) [2]float64
TEXT ·coshPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func coshPs(a [4]float32) [4]float32
TEXT ·coshPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func csqrtPs(a [4]float32) [4]float32
TEXT ·csqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtPs2pi(a [4]float32) M64
TEXT ·cvtPs2pi(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
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

// func cvtSi2ss(a [4]float32, b int) [4]float32
TEXT ·cvtSi2ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtSs2si(a [4]float32) int
TEXT ·cvtSs2si(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtepi16Epi32(a [16]byte) [16]byte
TEXT ·cvtepi16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepi16Epi64(a [16]byte) [16]byte
TEXT ·cvtepi16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func maskCvtepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

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

// func cvtepi32Epi64(a [16]byte) [16]byte
TEXT ·cvtepi32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepi32Pd(a [16]byte) [2]float64
TEXT ·cvtepi32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepi32Ps(a [16]byte) [4]float32
TEXT ·cvtepi32Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func maskCvtepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi32StoreuEpi8(SB),7,$0
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

// func cvtepi64Pd(a [16]byte) [2]float64
TEXT ·cvtepi64Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepi64Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepi64Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Ps(a [16]byte) [4]float32
TEXT ·cvtepi64Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Ps(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtepi64Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Ps(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtepi64Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtepi8Epi16(a [16]byte) [16]byte
TEXT ·cvtepi8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepi8Epi32(a [16]byte) [16]byte
TEXT ·cvtepi8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepi8Epi64(a [16]byte) [16]byte
TEXT ·cvtepi8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepu16Epi32(a [16]byte) [16]byte
TEXT ·cvtepu16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepu16Epi64(a [16]byte) [16]byte
TEXT ·cvtepu16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepu32Epi64(a [16]byte) [16]byte
TEXT ·cvtepu32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepu64Pd(a [16]byte) [2]float64
TEXT ·cvtepu64Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu64Pd(src [2]float64, k uint8, a [16]byte) [2]float64
TEXT ·maskCvtepu64Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu64Pd(k uint8, a [16]byte) [2]float64
TEXT ·maskzCvtepu64Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu64Ps(a [16]byte) [4]float32
TEXT ·cvtepu64Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu64Ps(src [4]float32, k uint8, a [16]byte) [4]float32
TEXT ·maskCvtepu64Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu64Ps(k uint8, a [16]byte) [4]float32
TEXT ·maskzCvtepu64Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu8Epi16(a [16]byte) [16]byte
TEXT ·cvtepu8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepu8Epi32(a [16]byte) [16]byte
TEXT ·cvtepu8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtepu8Epi64(a [16]byte) [16]byte
TEXT ·cvtepu8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtpdEpi32(a [2]float64) [16]byte
TEXT ·cvtpdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtpdEpi64(a [2]float64) [16]byte
TEXT ·cvtpdEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpi64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpi64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
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

// func cvtpdEpu64(a [2]float64) [16]byte
TEXT ·cvtpdEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpu64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvtpdEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpu64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvtpdEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpdPs(a [2]float64) [4]float32
TEXT ·cvtpdPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtphPs(a [16]byte) [4]float32
TEXT ·cvtphPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtpi16Ps(a M64) [4]float32
TEXT ·cvtpi16Ps(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpi32Pd(a M64) [2]float64
TEXT ·cvtpi32Pd(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpi32Ps(a [4]float32, b M64) [4]float32
TEXT ·cvtpi32Ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtpi32x2Ps(a M64, b M64) [4]float32
TEXT ·cvtpi32x2Ps(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtpi8Ps(a M64) [4]float32
TEXT ·cvtpi8Ps(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpsEpi32(a [4]float32) [16]byte
TEXT ·cvtpsEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvtpsEpi64(a [4]float32) [16]byte
TEXT ·cvtpsEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpi64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpi64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
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

// func cvtpsEpu64(a [4]float32) [16]byte
TEXT ·cvtpsEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpu64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvtpsEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpu64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvtpsEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpsPd(a [4]float32) [2]float64
TEXT ·cvtpsPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtpsPh(a [4]float32, rounding int) [16]byte
TEXT ·cvtpsPh(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func cvtpu16Ps(a M64) [4]float32
TEXT ·cvtpu16Ps(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpu8Ps(a M64) [4]float32
TEXT ·cvtpu8Ps(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtsdF64(a [2]float64) float64
TEXT ·cvtsdF64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
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

// func cvtsdSi32(a [2]float64) int
TEXT ·cvtsdSi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsdSi64(a [2]float64) int64
TEXT ·cvtsdSi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsdSi64x(a [2]float64) int64
TEXT ·cvtsdSi64x(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsdSs(a [4]float32, b [2]float64) [4]float32
TEXT ·cvtsdSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maskCvtsepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

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

// func maskCvtsepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi32StoreuEpi8(SB),7,$0
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

// func maskCvtsepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtsepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtshSs(a uint16) float32
TEXT ·cvtshSs(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func cvtsi128Si32(a [16]byte) int
TEXT ·cvtsi128Si32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsi128Si64(a [16]byte) int64
TEXT ·cvtsi128Si64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsi128Si64x(a [16]byte) int64
TEXT ·cvtsi128Si64x(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsi32Sd(a [2]float64, b int) [2]float64
TEXT ·cvtsi32Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtsi32Si128(a int) [16]byte
TEXT ·cvtsi32Si128(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtsi32Ss(a [4]float32, b int) [4]float32
TEXT ·cvtsi32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtsi64Sd(a [2]float64, b int64) [2]float64
TEXT ·cvtsi64Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtsi64Si128(a int64) [16]byte
TEXT ·cvtsi64Si128(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtsi64Ss(a [4]float32, b int64) [4]float32
TEXT ·cvtsi64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtsi64xSd(a [2]float64, b int64) [2]float64
TEXT ·cvtsi64xSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtsi64xSi128(a int64) [16]byte
TEXT ·cvtsi64xSi128(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtssF32(a [4]float32) float32
TEXT ·cvtssF32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVL $0, ret+16(FP)
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

// func cvtssSd(a [2]float64, b [4]float32) [2]float64
TEXT ·cvtssSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func cvtssSh(a float32, imm8 int) uint16
TEXT ·cvtssSh(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ imm8+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+12(FP)
	RET

// func cvtssSi32(a [4]float32) int
TEXT ·cvtssSi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtssSi64(a [4]float32) int64
TEXT ·cvtssSi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
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

// func cvttPs2pi(a [4]float32) M64
TEXT ·cvttPs2pi(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
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

// func cvttSs2si(a [4]float32) int
TEXT ·cvttSs2si(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttpdEpi32(a [2]float64) [16]byte
TEXT ·cvttpdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvttpdEpi64(a [2]float64) [16]byte
TEXT ·cvttpdEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpi64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpi64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
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

// func cvttpdEpu64(a [2]float64) [16]byte
TEXT ·cvttpdEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpu64(src [16]byte, k uint8, a [2]float64) [16]byte
TEXT ·maskCvttpdEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpu64(k uint8, a [2]float64) [16]byte
TEXT ·maskzCvttpdEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpsEpi32(a [4]float32) [16]byte
TEXT ·cvttpsEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func cvttpsEpi64(a [4]float32) [16]byte
TEXT ·cvttpsEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpi64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpi64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
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

// func cvttpsEpu64(a [4]float32) [16]byte
TEXT ·cvttpsEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpu64(src [16]byte, k uint8, a [4]float32) [16]byte
TEXT ·maskCvttpsEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpu64(k uint8, a [4]float32) [16]byte
TEXT ·maskzCvttpsEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
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

// func cvttsdSi32(a [2]float64) int
TEXT ·cvttsdSi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttsdSi64(a [2]float64) int64
TEXT ·cvttsdSi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttsdSi64x(a [2]float64) int64
TEXT ·cvttsdSi64x(SB),7,$0
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

// func cvttssSi32(a [4]float32) int
TEXT ·cvttssSi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvttssSi64(a [4]float32) int64
TEXT ·cvttssSi64(SB),7,$0
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

// func maskCvtusepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi16StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

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

// func maskCvtusepi32StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi32StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi32StoreuEpi8(SB),7,$0
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

// func maskCvtusepi64StoreuEpi16(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi32(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskCvtusepi64StoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
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

// func divEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divPd(a [2]float64, b [2]float64) [2]float64
TEXT ·divPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func divPs(a [4]float32, b [4]float32) [4]float32
TEXT ·divPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func divSd(a [2]float64, b [2]float64) [2]float64
TEXT ·divSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func divSs(a [4]float32, b [4]float32) [4]float32
TEXT ·divSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func dpPd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·dpPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func dpPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·dpPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func erfPd(a [2]float64) [2]float64
TEXT ·erfPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfPs(a [4]float32) [4]float32
TEXT ·erfPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcPd(a [2]float64) [2]float64
TEXT ·erfcPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcPs(a [4]float32) [4]float32
TEXT ·erfcPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcinvPd(a [2]float64) [2]float64
TEXT ·erfcinvPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcinvPs(a [4]float32) [4]float32
TEXT ·erfcinvPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfinvPd(a [2]float64) [2]float64
TEXT ·erfinvPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfinvPs(a [4]float32) [4]float32
TEXT ·erfinvPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expPd(a [2]float64) [2]float64
TEXT ·expPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expPs(a [4]float32) [4]float32
TEXT ·expPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp10Pd(a [2]float64) [2]float64
TEXT ·exp10Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp10Ps(a [4]float32) [4]float32
TEXT ·exp10Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp2Pd(a [2]float64) [2]float64
TEXT ·exp2Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp2Ps(a [4]float32) [4]float32
TEXT ·exp2Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func expm1Pd(a [2]float64) [2]float64
TEXT ·expm1Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expm1Ps(a [4]float32) [4]float32
TEXT ·expm1Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func extractEpi16(a [16]byte, imm8 int) int
TEXT ·extractEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extractEpi32(a [16]byte, imm8 int) int
TEXT ·extractEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extractEpi64(a [16]byte, imm8 int) int64
TEXT ·extractEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extractEpi8(a [16]byte, imm8 int) int
TEXT ·extractEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extractPs(a [4]float32, imm8 int) int
TEXT ·extractPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
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

// func floorPd(a [2]float64) [2]float64
TEXT ·floorPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func floorPs(a [4]float32) [4]float32
TEXT ·floorPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func floorSd(a [2]float64, b [2]float64) [2]float64
TEXT ·floorSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func floorSs(a [4]float32, b [4]float32) [4]float32
TEXT ·floorSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func fmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmaddSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmaddSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmaddsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmaddsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmaddsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmaddsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmsubSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmsubaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fmsubaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fmsubaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fmsubaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmaddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmaddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmaddSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmaddSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmaddSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmaddSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmsubSd(a [2]float64, b [2]float64, c [2]float64) [2]float64
TEXT ·fnmsubSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fnmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32
TEXT ·fnmsubSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func fpclassPdMask(a [2]float64, imm8 int) uint8
TEXT ·fpclassPdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassPdMask(k1 uint8, a [2]float64, imm8 int) uint8
TEXT ·maskFpclassPdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func fpclassPsMask(a [4]float32, imm8 int) uint8
TEXT ·fpclassPsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassPsMask(k1 uint8, a [4]float32, imm8 int) uint8
TEXT ·maskFpclassPsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func fpclassSdMask(a [2]float64, imm8 int) uint8
TEXT ·fpclassSdMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassSdMask(k1 uint8, a [2]float64, imm8 int) uint8
TEXT ·maskFpclassSdMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func fpclassSsMask(a [4]float32, imm8 int) uint8
TEXT ·fpclassSsMask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVB $0, ret+24(FP)
	RET

// func maskFpclassSsMask(k1 uint8, a [4]float32, imm8 int) uint8
TEXT ·maskFpclassSsMask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVB $0, ret+28(FP)
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

// func haddEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·haddEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func haddEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·haddEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func haddPd(a [2]float64, b [2]float64) [2]float64
TEXT ·haddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func haddPs(a [4]float32, b [4]float32) [4]float32
TEXT ·haddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func haddsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·haddsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsubEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·hsubEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsubEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·hsubEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsubPd(a [2]float64, b [2]float64) [2]float64
TEXT ·hsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsubPs(a [4]float32, b [4]float32) [4]float32
TEXT ·hsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsubsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·hsubsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hypotPd(a [2]float64, b [2]float64) [2]float64
TEXT ·hypotPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hypotPs(a [4]float32, b [4]float32) [4]float32
TEXT ·hypotPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func i32gatherEpi32(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i32gatherEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherEpi32(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI32gatherEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI32gatherEpi32(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32gatherEpi64(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i32gatherEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherEpi64(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI32gatherEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI32gatherEpi64(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32gatherPd(base_addr float64, vindex [16]byte, scale int) [2]float64
TEXT ·i32gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherPd(src [2]float64, base_addr float64, vindex [16]byte, mask [2]float64, scale int) [2]float64
TEXT ·maskI32gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI32gatherPd(src [2]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32gatherPs(base_addr float32, vindex [16]byte, scale int) [4]float32
TEXT ·i32gatherPs(SB),7,$0
	MOVL base_addr+0(FP),R8
	MOVOU vindex+4(FP),X1
	MOVQ scale+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskI32gatherPs(src [4]float32, base_addr float32, vindex [16]byte, mask [4]float32, scale int) [4]float32
TEXT ·maskI32gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVL base_addr+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVOU mask+36(FP),X3
	MOVQ scale+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mmaskI32gatherPs(src [4]float32, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI32gatherPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
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

// func i64gatherEpi32(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i64gatherEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherEpi32(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI64gatherEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI64gatherEpi32(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherEpi64(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i64gatherEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherEpi64(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI64gatherEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI64gatherEpi64(src [16]byte, k uint8, vindex [16]byte, base_addr uintptr, scale int) [16]byte
TEXT ·mmaskI64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherPd(base_addr float64, vindex [16]byte, scale int) [2]float64
TEXT ·i64gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherPd(src [2]float64, base_addr float64, vindex [16]byte, mask [2]float64, scale int) [2]float64
TEXT ·maskI64gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI64gatherPd(src [2]float64, k uint8, vindex [16]byte, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherPs(base_addr float32, vindex [16]byte, scale int) [4]float32
TEXT ·i64gatherPs(SB),7,$0
	MOVL base_addr+0(FP),R8
	MOVOU vindex+4(FP),X1
	MOVQ scale+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskI64gatherPs(src [4]float32, base_addr float32, vindex [16]byte, mask [4]float32, scale int) [4]float32
TEXT ·maskI64gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVL base_addr+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVOU mask+36(FP),X3
	MOVQ scale+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mmaskI64gatherPs(src [4]float32, k uint8, vindex [16]byte, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
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

// func idivEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·idivEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func idivremEpi32(mem_addr [16]byte, a [16]byte, b [16]byte) [16]byte
TEXT ·idivremEpi32(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func insertEpi16(a [16]byte, i int, imm8 int) [16]byte
TEXT ·insertEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insertEpi32(a [16]byte, i int, imm8 int) [16]byte
TEXT ·insertEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insertEpi64(a [16]byte, i int64, imm8 int) [16]byte
TEXT ·insertEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insertEpi8(a [16]byte, i int, imm8 int) [16]byte
TEXT ·insertEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insertPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·insertPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func invcbrtPd(a [2]float64) [2]float64
TEXT ·invcbrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func invcbrtPs(a [4]float32) [4]float32
TEXT ·invcbrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func invsqrtPd(a [2]float64) [2]float64
TEXT ·invsqrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func invsqrtPs(a [4]float32) [4]float32
TEXT ·invsqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func iremEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·iremEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func lddquSi128(mem_addr M128iConst) [16]byte
TEXT ·lddquSi128(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
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

// func loadPd(mem_addr float64) [2]float64
TEXT ·loadPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
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

// func loadPd1(mem_addr float64) [2]float64
TEXT ·loadPd1(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loadPs(mem_addr float32) [4]float32
TEXT ·loadPs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
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

// func loadPs1(mem_addr float32) [4]float32
TEXT ·loadPs1(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func loadSd(mem_addr float64) [2]float64
TEXT ·loadSd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
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

// func loadSi128(mem_addr M128iConst) [16]byte
TEXT ·loadSi128(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loadSs(mem_addr float32) [4]float32
TEXT ·loadSs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
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

// func load1Pd(mem_addr float64) [2]float64
TEXT ·load1Pd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func load1Ps(mem_addr float32) [4]float32
TEXT ·load1Ps(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func loaddupPd(mem_addr float64) [2]float64
TEXT ·loaddupPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loadhPd(a [2]float64, mem_addr float64) [2]float64
TEXT ·loadhPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ mem_addr+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func loadlEpi64(mem_addr M128iConst) [16]byte
TEXT ·loadlEpi64(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loadlPd(a [2]float64, mem_addr float64) [2]float64
TEXT ·loadlPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ mem_addr+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func loadrPd(mem_addr float64) [2]float64
TEXT ·loadrPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loadrPs(mem_addr float32) [4]float32
TEXT ·loadrPs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func maskLoaduEpi16(src [16]byte, k uint8, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoaduEpi16(k uint8, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
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

// func maskLoaduEpi8(src [16]byte, k uint16, mem_addr uintptr) [16]byte
TEXT ·maskLoaduEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoaduEpi8(k uint16, mem_addr uintptr) [16]byte
TEXT ·maskzLoaduEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduPd(mem_addr float64) [2]float64
TEXT ·loaduPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
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

// func loaduPs(mem_addr float32) [4]float32
TEXT ·loaduPs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
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

// func loaduSi128(mem_addr M128iConst) [16]byte
TEXT ·loaduSi128(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func logPd(a [2]float64) [2]float64
TEXT ·logPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func logPs(a [4]float32) [4]float32
TEXT ·logPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log10Pd(a [2]float64) [2]float64
TEXT ·log10Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log10Ps(a [4]float32) [4]float32
TEXT ·log10Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log1pPd(a [2]float64) [2]float64
TEXT ·log1pPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log1pPs(a [4]float32) [4]float32
TEXT ·log1pPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log2Pd(a [2]float64) [2]float64
TEXT ·log2Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log2Ps(a [4]float32) [4]float32
TEXT ·log2Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func logbPd(a [2]float64) [2]float64
TEXT ·logbPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func logbPs(a [4]float32) [4]float32
TEXT ·logbPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func maddEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·maddEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maddubsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·maddubsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maskloadEpi32(mem_addr int, mask [16]byte) [16]byte
TEXT ·maskloadEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskloadEpi64(mem_addr int, mask [16]byte) [16]byte
TEXT ·maskloadEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskloadPd(mem_addr float64, mask [16]byte) [2]float64
TEXT ·maskloadPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskloadPs(mem_addr float32, mask [16]byte) [4]float32
TEXT ·maskloadPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU mask+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskmoveuSi128(a [16]byte, mask [16]byte, mem_addr byte) 
TEXT ·maskmoveuSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1
	MOVB mem_addr+32(FP),R10

	//TODO: Code missing

	RET

// func maskstoreEpi32(mem_addr int, mask [16]byte, a [16]byte) 
TEXT ·maskstoreEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func maskstoreEpi64(mem_addr int64, mask [16]byte, a [16]byte) 
TEXT ·maskstoreEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func maskstorePd(mem_addr float64, mask [16]byte, a [2]float64) 
TEXT ·maskstorePd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func maskstorePs(mem_addr float32, mask [16]byte, a [4]float32) 
TEXT ·maskstorePs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU mask+4(FP),X1
	MOVOU a+20(FP),X2

	//TODO: Code missing

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

// func maxEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxPd(a [2]float64, b [2]float64) [2]float64
TEXT ·maxPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxPs(a [4]float32, b [4]float32) [4]float32
TEXT ·maxPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxSd(a [2]float64, b [2]float64) [2]float64
TEXT ·maxSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maxSs(a [4]float32, b [4]float32) [4]float32
TEXT ·maxSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minPd(a [2]float64, b [2]float64) [2]float64
TEXT ·minPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minPs(a [4]float32, b [4]float32) [4]float32
TEXT ·minPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minSd(a [2]float64, b [2]float64) [2]float64
TEXT ·minSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func minSs(a [4]float32, b [4]float32) [4]float32
TEXT ·minSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func minposEpu16(a [16]byte) [16]byte
TEXT ·minposEpu16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func moveEpi64(a [16]byte) [16]byte
TEXT ·moveEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func moveSd(a [2]float64, b [2]float64) [2]float64
TEXT ·moveSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func moveSs(a [4]float32, b [4]float32) [4]float32
TEXT ·moveSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func movedupPd(a [2]float64) [2]float64
TEXT ·movedupPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func movehdupPs(a [4]float32) [4]float32
TEXT ·movehdupPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func movehlPs(a [4]float32, b [4]float32) [4]float32
TEXT ·movehlPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func moveldupPs(a [4]float32) [4]float32
TEXT ·moveldupPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func movelhPs(a [4]float32, b [4]float32) [4]float32
TEXT ·movelhPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func movemaskEpi8(a [16]byte) int
TEXT ·movemaskEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func movemaskPd(a [2]float64) int
TEXT ·movemaskPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func movemaskPs(a [4]float32) int
TEXT ·movemaskPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func movmEpi16(k uint8) [16]byte
TEXT ·movmEpi16(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movmEpi32(k uint8) [16]byte
TEXT ·movmEpi32(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movmEpi64(k uint8) [16]byte
TEXT ·movmEpi64(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movmEpi8(k uint16) [16]byte
TEXT ·movmEpi8(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movpi64Epi64(a M64) [16]byte
TEXT ·movpi64Epi64(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func mpsadbwEpu8(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·mpsadbwEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func mulEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·mulEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·mulEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulPd(a [2]float64, b [2]float64) [2]float64
TEXT ·mulPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulPs(a [4]float32, b [4]float32) [4]float32
TEXT ·mulPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	MULPS X1, X0

	MOVOU X0, ret+32(FP)
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

// func mulSd(a [2]float64, b [2]float64) [2]float64
TEXT ·mulSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulSs(a [4]float32, b [4]float32) [4]float32
TEXT ·mulSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulhiEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulhiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulhiEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulhiEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulhrsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulhrsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulloEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulloEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func mulloEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·mulloEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMulloEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskMulloEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulloEpi64(k uint8, a [16]byte, b [16]byte) [16]byte
TEXT ·maskzMulloEpi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mulloEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·mulloEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maskOrPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskOrPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOrPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzOrPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func orPd(a [2]float64, b [2]float64) [2]float64
TEXT ·orPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskOrPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskOrPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOrPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzOrPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func orPs(a [4]float32, b [4]float32) [4]float32
TEXT ·orPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func orSi128(a [16]byte, b [16]byte) [16]byte
TEXT ·orSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func packsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·packsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func packsEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·packsEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func packusEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·packusEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func packusEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·packusEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func pdepU32(a uint32, mask uint32) uint32
TEXT ·pdepU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL mask+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func pdepU64(a uint64, mask uint64) uint64
TEXT ·pdepU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ mask+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
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

// func permutePd(a [2]float64, imm8 int) [2]float64
TEXT ·permutePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func permutePs(a [4]float32, imm8 int) [4]float32
TEXT ·permutePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func permutevarPd(a [2]float64, b [16]byte) [2]float64
TEXT ·permutevarPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func permutevarPs(a [4]float32, b [16]byte) [4]float32
TEXT ·permutevarPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func powPd(a [2]float64, b [2]float64) [2]float64
TEXT ·powPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func powPs(a [4]float32, b [4]float32) [4]float32
TEXT ·powPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRangePd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskRangePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangePd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzRangePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rangePd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·rangePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRangePs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskRangePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangePs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzRangePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rangePs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·rangePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRangeRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskRangeRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzRangeRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskzRangeRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func rangeRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·rangeRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskRangeRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskRangeRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzRangeRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskzRangeRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func rangeRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·rangeRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskRangeSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskRangeSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangeSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzRangeSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskRangeSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskRangeSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRangeSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzRangeSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rcpPs(a [4]float32) [4]float32
TEXT ·rcpPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func rcpSs(a [4]float32) [4]float32
TEXT ·rcpSs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func maskRcp28RoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskRcp28RoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRcp28RoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzRcp28RoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rcp28RoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·rcp28RoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRcp28RoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskRcp28RoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRcp28RoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzRcp28RoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rcp28RoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·rcp28RoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRcp28Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRcp28Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRcp28Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRcp28Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rcp28Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rcp28Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRcp28Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRcp28Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRcp28Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRcp28Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rcp28Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rcp28Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskReducePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskReducePd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzReducePd(k uint8, a [2]float64, imm8 int) [2]float64
TEXT ·maskzReducePd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func reducePd(a [2]float64, imm8 int) [2]float64
TEXT ·reducePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskReducePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskReducePs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzReducePs(k uint8, a [4]float32, imm8 int) [4]float32
TEXT ·maskzReducePs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func reducePs(a [4]float32, imm8 int) [4]float32
TEXT ·reducePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskReduceRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskReduceRoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzReduceRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·maskzReduceRoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func reduceRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64
TEXT ·reduceRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskReduceRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskReduceRoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12
	MOVQ rounding+60(FP),R13

	//TODO: Code missing

	MOVOU X0, ret+68(FP)
	RET

// func maskzReduceRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·maskzReduceRoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11
	MOVQ rounding+44(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func reduceRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32
TEXT ·reduceRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10
	MOVQ rounding+40(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskReduceSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskReduceSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzReduceSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·maskzReduceSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func reduceSd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·reduceSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskReduceSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskReduceSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzReduceSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·maskzReduceSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func reduceSs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·reduceSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func remEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func roundPd(a [2]float64, rounding int) [2]float64
TEXT ·roundPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func roundPs(a [4]float32, rounding int) [4]float32
TEXT ·roundPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func roundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·roundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func roundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·roundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func rsqrtPs(a [4]float32) [4]float32
TEXT ·rsqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func rsqrtSs(a [4]float32) [4]float32
TEXT ·rsqrtSs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func maskRsqrt28RoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskRsqrt28RoundSd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRsqrt28RoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·maskzRsqrt28RoundSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rsqrt28RoundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·rsqrt28RoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRsqrt28RoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskRsqrt28RoundSs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ rounding+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzRsqrt28RoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·maskzRsqrt28RoundSs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func rsqrt28RoundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·rsqrt28RoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskRsqrt28Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskRsqrt28Sd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRsqrt28Sd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzRsqrt28Sd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rsqrt28Sd(a [2]float64, b [2]float64) [2]float64
TEXT ·rsqrt28Sd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRsqrt28Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskRsqrt28Ss(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRsqrt28Ss(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzRsqrt28Ss(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rsqrt28Ss(a [4]float32, b [4]float32) [4]float32
TEXT ·rsqrt28Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sadEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·sadEpu8(SB),7,$0
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

// func setEpi16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [16]byte
TEXT ·setEpi16(SB),7,$0
	MOVW e7+0(FP),R8
	MOVW e6+4(FP),R9
	MOVW e5+8(FP),R10
	MOVW e4+12(FP),R11
	MOVW e3+16(FP),R12
	MOVW e2+20(FP),R13
	MOVW e1+24(FP),R14
	MOVW e0+28(FP),R15

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func setEpi32(e3 int, e2 int, e1 int, e0 int) [16]byte
TEXT ·setEpi32(SB),7,$0
	MOVQ e3+0(FP),R8
	MOVQ e2+8(FP),R9
	MOVQ e1+16(FP),R10
	MOVQ e0+24(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func setEpi64(e1 M64, e0 M64) [16]byte
TEXT ·setEpi64(SB),7,$0
	MOVQ e1+0(FP),M0
	MOVQ e0+8(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setEpi64x(e1 int64, e0 int64) [16]byte
TEXT ·setEpi64x(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setEpi8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [16]byte
TEXT ·setEpi8(SB),7,$0
	// Unimplemented. Unknown register for type byte
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func setPd(e1 float64, e0 float64) [2]float64
TEXT ·setPd(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setPd1(a float64) [2]float64
TEXT ·setPd1(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func setPs(e3 float32, e2 float32, e1 float32, e0 float32) [4]float32
TEXT ·setPs(SB),7,$0
	MOVL e3+0(FP),R8
	MOVL e2+4(FP),R9
	MOVL e1+8(FP),R10
	MOVL e0+12(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setPs1(a float32) [4]float32
TEXT ·setPs1(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func setSd(a float64) [2]float64
TEXT ·setSd(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func setSs(a float32) [4]float32
TEXT ·setSs(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
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

// func set1Epi16(a int16) [16]byte
TEXT ·set1Epi16(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
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

// func set1Epi32(a int) [16]byte
TEXT ·set1Epi32(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
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

// func set1Epi64(a M64) [16]byte
TEXT ·set1Epi64(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func set1Epi64x(a int64) [16]byte
TEXT ·set1Epi64x(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
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

// func set1Epi8(a byte) [16]byte
TEXT ·set1Epi8(SB),7,$0
	MOVB a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func set1Pd(a float64) [2]float64
TEXT ·set1Pd(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func set1Ps(a float32) [4]float32
TEXT ·set1Ps(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func setrEpi16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [16]byte
TEXT ·setrEpi16(SB),7,$0
	MOVW e7+0(FP),R8
	MOVW e6+4(FP),R9
	MOVW e5+8(FP),R10
	MOVW e4+12(FP),R11
	MOVW e3+16(FP),R12
	MOVW e2+20(FP),R13
	MOVW e1+24(FP),R14
	MOVW e0+28(FP),R15

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func setrEpi32(e3 int, e2 int, e1 int, e0 int) [16]byte
TEXT ·setrEpi32(SB),7,$0
	MOVQ e3+0(FP),R8
	MOVQ e2+8(FP),R9
	MOVQ e1+16(FP),R10
	MOVQ e0+24(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func setrEpi64(e1 M64, e0 M64) [16]byte
TEXT ·setrEpi64(SB),7,$0
	MOVQ e1+0(FP),M0
	MOVQ e0+8(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setrEpi8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [16]byte
TEXT ·setrEpi8(SB),7,$0
	// Unimplemented. Unknown register for type byte
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func setrPd(e1 float64, e0 float64) [2]float64
TEXT ·setrPd(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setrPs(e3 float32, e2 float32, e1 float32, e0 float32) [4]float32
TEXT ·setrPs(SB),7,$0
	MOVL e3+0(FP),R8
	MOVL e2+4(FP),R9
	MOVL e1+8(FP),R10
	MOVL e0+12(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setzeroPd() [2]float64
TEXT ·setzeroPd(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func setzeroPs() [4]float32
TEXT ·setzeroPs(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func setzeroSi128() [16]byte
TEXT ·setzeroSi128(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func sha1msg1Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha1msg1Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha1msg2Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha1msg2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha1nexteEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha1nexteEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha1rnds4Epu32(a [16]byte, b [16]byte, fnc int) [16]byte
TEXT ·sha1rnds4Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ fnc+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func sha256msg1Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha256msg1Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha256msg2Epu32(a [16]byte, b [16]byte) [16]byte
TEXT ·sha256msg2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha256rnds2Epu32(a [16]byte, b [16]byte, k [16]byte) [16]byte
TEXT ·sha256rnds2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU k+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
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

// func shuffleEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·shuffleEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func shuffleEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·shuffleEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func shufflePd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·shufflePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
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

// func shufflePs(a [4]float32, b [4]float32, imm8 uint32) [4]float32
TEXT ·shufflePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVL imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
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

// func shufflehiEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·shufflehiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func shuffleloEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·shuffleloEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func signEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·signEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func signEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·signEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func signEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·signEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sinPd(a [2]float64) [2]float64
TEXT ·sinPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sinPs(a [4]float32) [4]float32
TEXT ·sinPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sincosPd(mem_addr [2]float64, a [2]float64) [2]float64
TEXT ·sincosPd(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sincosPs(mem_addr [4]float32, a [4]float32) [4]float32
TEXT ·sincosPs(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sindPd(a [2]float64) [2]float64
TEXT ·sindPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sindPs(a [4]float32) [4]float32
TEXT ·sindPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sinhPd(a [2]float64) [2]float64
TEXT ·sinhPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sinhPs(a [4]float32) [4]float32
TEXT ·sinhPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func sllEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sllEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sllEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sllEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sllEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sllEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func slliEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·slliEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func slliEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·slliEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func slliEpi64(a [16]byte, imm8 int) [16]byte
TEXT ·slliEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func slliSi128(a [16]byte, imm8 int) [16]byte
TEXT ·slliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func sllvEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sllvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sllvEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sllvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sqrtPd(a [2]float64) [2]float64
TEXT ·sqrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func sqrtPs(a [4]float32) [4]float32
TEXT ·sqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func sqrtSd(a [2]float64, b [2]float64) [2]float64
TEXT ·sqrtSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sqrtSs(a [4]float32) [4]float32
TEXT ·sqrtSs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func sraEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sraEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sraEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sraEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func sraiEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·sraiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func sraiEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·sraiEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func sravEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sravEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func srlEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·srlEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func srlEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·srlEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func srlEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·srlEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func srliEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·srliEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func srliEpi32(a [16]byte, imm8 int) [16]byte
TEXT ·srliEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func srliEpi64(a [16]byte, imm8 int) [16]byte
TEXT ·srliEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srliSi128(a [16]byte, imm8 int) [16]byte
TEXT ·srliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
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

// func srlvEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·srlvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func srlvEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·srlvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskStoreEpi32(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreEpi64(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStorePd(mem_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskStorePd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storePd(mem_addr float64, a [2]float64) 
TEXT ·storePd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storePd1(mem_addr float64, a [2]float64) 
TEXT ·storePd1(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func maskStorePs(mem_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskStorePs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storePs(mem_addr float32, a [4]float32) 
TEXT ·storePs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func storePs1(mem_addr float32, a [4]float32) 
TEXT ·storePs1(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func maskStoreSd(mem_addr float64, k uint8, a [2]float64) 
TEXT ·maskStoreSd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVB k+8(FP),R9
	MOVOU a+12(FP),X2

	//TODO: Code missing

	RET

// func storeSd(mem_addr float64, a [2]float64) 
TEXT ·storeSd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storeSi128(mem_addr [16]byte, a [16]byte) 
TEXT ·storeSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func maskStoreSs(mem_addr float32, k uint8, a [4]float32) 
TEXT ·maskStoreSs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVB k+4(FP),R9
	MOVOU a+8(FP),X2

	//TODO: Code missing

	RET

// func storeSs(mem_addr float32, a [4]float32) 
TEXT ·storeSs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func store1Pd(mem_addr float64, a [2]float64) 
TEXT ·store1Pd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func store1Ps(mem_addr float32, a [4]float32) 
TEXT ·store1Ps(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func storehPd(mem_addr float64, a [2]float64) 
TEXT ·storehPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storelEpi64(mem_addr [16]byte, a [16]byte) 
TEXT ·storelEpi64(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func storelPd(mem_addr float64, a [2]float64) 
TEXT ·storelPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storerPd(mem_addr float64, a [2]float64) 
TEXT ·storerPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storerPs(mem_addr float32, a [4]float32) 
TEXT ·storerPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func maskStoreuEpi16(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi32(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi64(mem_addr uintptr, k uint8, a [16]byte) 
TEXT ·maskStoreuEpi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuEpi8(mem_addr uintptr, k uint16, a [16]byte) 
TEXT ·maskStoreuEpi8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreuPd(mem_addr uintptr, k uint8, a [2]float64) 
TEXT ·maskStoreuPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuPd(mem_addr float64, a [2]float64) 
TEXT ·storeuPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func maskStoreuPs(mem_addr uintptr, k uint8, a [4]float32) 
TEXT ·maskStoreuPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuPs(mem_addr float32, a [4]float32) 
TEXT ·storeuPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func storeuSi128(mem_addr [16]byte, a [16]byte) 
TEXT ·storeuSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func streamLoadSi128(mem_addr [16]byte) [16]byte
TEXT ·streamLoadSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func streamPd(mem_addr float64, a [2]float64) 
TEXT ·streamPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func streamPs(mem_addr float32, a [4]float32) 
TEXT ·streamPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	RET

// func streamSi128(mem_addr [16]byte, a [16]byte) 
TEXT ·streamSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

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

// func subEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subPd(a [2]float64, b [2]float64) [2]float64
TEXT ·subPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subPs(a [4]float32, b [4]float32) [4]float32
TEXT ·subPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subSd(a [2]float64, b [2]float64) [2]float64
TEXT ·subSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subSs(a [4]float32, b [4]float32) [4]float32
TEXT ·subSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subsEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subsEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func subsEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func svmlCeilPd(a [2]float64) [2]float64
TEXT ·svmlCeilPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlCeilPs(a [4]float32) [4]float32
TEXT ·svmlCeilPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlFloorPd(a [2]float64) [2]float64
TEXT ·svmlFloorPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlFloorPs(a [4]float32) [4]float32
TEXT ·svmlFloorPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlRoundPd(a [2]float64) [2]float64
TEXT ·svmlRoundPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlRoundPs(a [4]float32) [4]float32
TEXT ·svmlRoundPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlSqrtPd(a [2]float64) [2]float64
TEXT ·svmlSqrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlSqrtPs(a [4]float32) [4]float32
TEXT ·svmlSqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanPd(a [2]float64) [2]float64
TEXT ·tanPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanPs(a [4]float32) [4]float32
TEXT ·tanPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tandPd(a [2]float64) [2]float64
TEXT ·tandPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tandPs(a [4]float32) [4]float32
TEXT ·tandPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanhPd(a [2]float64) [2]float64
TEXT ·tanhPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanhPs(a [4]float32) [4]float32
TEXT ·tanhPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func testcPd(a [2]float64, b [2]float64) int
TEXT ·testcPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testcPs(a [4]float32, b [4]float32) int
TEXT ·testcPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testcSi128(a [16]byte, b [16]byte) int
TEXT ·testcSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
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

// func testnzcPd(a [2]float64, b [2]float64) int
TEXT ·testnzcPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testnzcPs(a [4]float32, b [4]float32) int
TEXT ·testnzcPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testnzcSi128(a [16]byte, b [16]byte) int
TEXT ·testnzcSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testzPd(a [2]float64, b [2]float64) int
TEXT ·testzPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testzPs(a [4]float32, b [4]float32) int
TEXT ·testzPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testzSi128(a [16]byte, b [16]byte) int
TEXT ·testzSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func truncPd(a [2]float64) [2]float64
TEXT ·truncPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func truncPs(a [4]float32) [4]float32
TEXT ·truncPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func ucomieqSd(a [2]float64, b [2]float64) int
TEXT ·ucomieqSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomieqSs(a [4]float32, b [4]float32) int
TEXT ·ucomieqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomigeSd(a [2]float64, b [2]float64) int
TEXT ·ucomigeSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomigeSs(a [4]float32, b [4]float32) int
TEXT ·ucomigeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomigtSd(a [2]float64, b [2]float64) int
TEXT ·ucomigtSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomigtSs(a [4]float32, b [4]float32) int
TEXT ·ucomigtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomileSd(a [2]float64, b [2]float64) int
TEXT ·ucomileSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomileSs(a [4]float32, b [4]float32) int
TEXT ·ucomileSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomiltSd(a [2]float64, b [2]float64) int
TEXT ·ucomiltSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomiltSs(a [4]float32, b [4]float32) int
TEXT ·ucomiltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomineqSd(a [2]float64, b [2]float64) int
TEXT ·ucomineqSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func ucomineqSs(a [4]float32, b [4]float32) int
TEXT ·ucomineqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func udivEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·udivEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func udivremEpi32(mem_addr [16]byte, a [16]byte, b [16]byte) [16]byte
TEXT ·udivremEpi32(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func undefinedPd() [2]float64
TEXT ·undefinedPd(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func undefinedPs() [4]float32
TEXT ·undefinedPs(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func undefinedSi128() [16]byte
TEXT ·undefinedSi128(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
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

// func unpackhiEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackhiEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackhiEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackhiEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackhiPd(a [2]float64, b [2]float64) [2]float64
TEXT ·unpackhiPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackhiPs(a [4]float32, b [4]float32) [4]float32
TEXT ·unpackhiPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackloEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackloEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackloEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackloEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackloPd(a [2]float64, b [2]float64) [2]float64
TEXT ·unpackloPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func unpackloPs(a [4]float32, b [4]float32) [4]float32
TEXT ·unpackloPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func uremEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·uremEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func maskXorPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskXorPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXorPd(k uint8, a [2]float64, b [2]float64) [2]float64
TEXT ·maskzXorPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func xorPd(a [2]float64, b [2]float64) [2]float64
TEXT ·xorPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskXorPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskXorPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXorPs(k uint8, a [4]float32, b [4]float32) [4]float32
TEXT ·maskzXorPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func xorPs(a [4]float32, b [4]float32) [4]float32
TEXT ·xorPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func xorSi128(a [16]byte, b [16]byte) [16]byte
TEXT ·xorSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

