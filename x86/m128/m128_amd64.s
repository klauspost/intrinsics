// func abs16(a M128i) M128i
TEXT ·abs16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskAbs16(src M128i, k uint8, a M128i) M128i
TEXT ·maskAbs16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbs16(k uint8, a M128i) M128i
TEXT ·maskzAbs16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func abs32(a M128i) M128i
TEXT ·abs32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskAbs32(src M128i, k uint8, a M128i) M128i
TEXT ·maskAbs32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbs32(k uint8, a M128i) M128i
TEXT ·maskzAbs32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func abs64(a M128i) M128i
TEXT ·abs64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskAbs64(src M128i, k uint8, a M128i) M128i
TEXT ·maskAbs64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbs64(k uint8, a M128i) M128i
TEXT ·maskzAbs64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func abs8(a M128i) M128i
TEXT ·abs8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskAbs8(src M128i, k uint16, a M128i) M128i
TEXT ·maskAbs8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzAbs8(k uint16, a M128i) M128i
TEXT ·maskzAbs8(SB),7,$0
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

// func add16(a M128i, b M128i) M128i
TEXT ·add16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAdd16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAdd16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAdd16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAdd16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func add32(a M128i, b M128i) M128i
TEXT ·add32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAdd32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAdd32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAdd32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAdd32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func add64(a M128i, b M128i) M128i
TEXT ·add64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAdd64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAdd64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAdd64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAdd64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func add8(a M128i, b M128i) M128i
TEXT ·add8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	PADDB  X1,X0

	MOVOU X0, ret+32(FP)
	RET

// func maskAdd8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskAdd8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAdd8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzAdd8(SB),7,$0
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

// func adds16(a M128i, b M128i) M128i
TEXT ·adds16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAdds16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAdds16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAdds16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAdds16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func adds8(a M128i, b M128i) M128i
TEXT ·adds8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAdds8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskAdds8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAdds8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzAdds8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func addsEpu16(a M128i, b M128i) M128i
TEXT ·addsEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAddsEpu16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAddsEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddsEpu16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAddsEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func addsEpu8(a M128i, b M128i) M128i
TEXT ·addsEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAddsEpu8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskAddsEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAddsEpu8(k uint16, a M128i, b M128i) M128i
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

// func aesdecSi128(a M128i, RoundKey M128i) M128i
TEXT ·aesdecSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesdeclastSi128(a M128i, RoundKey M128i) M128i
TEXT ·aesdeclastSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesencSi128(a M128i, RoundKey M128i) M128i
TEXT ·aesencSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesenclastSi128(a M128i, RoundKey M128i) M128i
TEXT ·aesenclastSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU RoundKey+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func aesimcSi128(a M128i) M128i
TEXT ·aesimcSi128(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func aeskeygenassistSi128(a M128i, imm8 int) M128i
TEXT ·aeskeygenassistSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func alignr32(a M128i, b M128i, count int) M128i
TEXT ·alignr32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAlignr32(src M128i, k uint8, a M128i, b M128i, count int) M128i
TEXT ·maskAlignr32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAlignr32(k uint8, a M128i, b M128i, count int) M128i
TEXT ·maskzAlignr32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func alignr64(a M128i, b M128i, count int) M128i
TEXT ·alignr64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAlignr64(src M128i, k uint8, a M128i, b M128i, count int) M128i
TEXT ·maskAlignr64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAlignr64(k uint8, a M128i, b M128i, count int) M128i
TEXT ·maskzAlignr64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func alignr8(a M128i, b M128i, count int) M128i
TEXT ·alignr8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskAlignr8(src M128i, k uint16, a M128i, b M128i, count int) M128i
TEXT ·maskAlignr8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ count+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzAlignr8(k uint16, a M128i, b M128i, count int) M128i
TEXT ·maskzAlignr8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ count+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskAnd32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAnd32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAnd32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAnd32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAnd64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAnd64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAnd64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAnd64(SB),7,$0
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

// func andSi128(a M128i, b M128i) M128i
TEXT ·andSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAndnot32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAndnot32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnot32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAndnot32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskAndnot64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAndnot64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAndnot64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAndnot64(SB),7,$0
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

// func andnotSi128(a M128i, b M128i) M128i
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

// func avgEpu16(a M128i, b M128i) M128i
TEXT ·avgEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAvgEpu16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskAvgEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAvgEpu16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzAvgEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func avgEpu8(a M128i, b M128i) M128i
TEXT ·avgEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskAvgEpu8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskAvgEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzAvgEpu8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzAvgEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func blend16(a M128i, b M128i, imm8 int) M128i
TEXT ·blend16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskBlend16(k uint8, a M128i, b M128i) M128i
TEXT ·maskBlend16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func blend32(a M128i, b M128i, imm8 int) M128i
TEXT ·blend32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskBlend32(k uint8, a M128i, b M128i) M128i
TEXT ·maskBlend32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlend64(k uint8, a M128i, b M128i) M128i
TEXT ·maskBlend64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskBlend8(k uint16, a M128i, b M128i) M128i
TEXT ·maskBlend8(SB),7,$0
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

// func blendv8(a M128i, b M128i, mask M128i) M128i
TEXT ·blendv8(SB),7,$0
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

// func broadcastb8(a M128i) M128i
TEXT ·broadcastb8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskBroadcastb8(src M128i, k uint16, a M128i) M128i
TEXT ·maskBroadcastb8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastb8(k uint16, a M128i) M128i
TEXT ·maskzBroadcastb8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func broadcastd32(a M128i) M128i
TEXT ·broadcastd32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskBroadcastd32(src M128i, k uint8, a M128i) M128i
TEXT ·maskBroadcastd32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastd32(k uint8, a M128i) M128i
TEXT ·maskzBroadcastd32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func broadcastmb64(k uint8) M128i
TEXT ·broadcastmb64(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastmw32(k uint16) M128i
TEXT ·broadcastmw32(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func broadcastq64(a M128i) M128i
TEXT ·broadcastq64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskBroadcastq64(src M128i, k uint8, a M128i) M128i
TEXT ·maskBroadcastq64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastq64(k uint8, a M128i) M128i
TEXT ·maskzBroadcastq64(SB),7,$0
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

// func broadcastw16(a M128i) M128i
TEXT ·broadcastw16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskBroadcastw16(src M128i, k uint8, a M128i) M128i
TEXT ·maskBroadcastw16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzBroadcastw16(k uint8, a M128i) M128i
TEXT ·maskzBroadcastw16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func bslliSi128(a M128i, imm8 int) M128i
TEXT ·bslliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func bsrliSi128(a M128i, imm8 int) M128i
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

// func castpdSi128(a [2]float64) M128i
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

// func castpsSi128(a [4]float32) M128i
TEXT ·castpsSi128(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castsi128Pd(a M128i) [2]float64
TEXT ·castsi128Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func castsi128Ps(a M128i) [4]float32
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

// func clmulepi64Si128(a M128i, b M128i, imm8 int) M128i
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

// func cmp16Mask(a M128i, b M128i, imm8 int) uint8
TEXT ·cmp16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmp16Mask(k1 uint8, a M128i, b M128i, imm8 int) uint8
TEXT ·maskCmp16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmp32Mask(a M128i, b M128i, imm8 uint8) uint8
TEXT ·cmp32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmp32Mask(k1 uint8, a M128i, b M128i, imm8 uint8) uint8
TEXT ·maskCmp32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmp64Mask(a M128i, b M128i, imm8 uint8) uint8
TEXT ·cmp64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmp64Mask(k1 uint8, a M128i, b M128i, imm8 uint8) uint8
TEXT ·maskCmp64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmp8Mask(a M128i, b M128i, imm8 int) uint16
TEXT ·cmp8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVW $0, ret+40(FP)
	RET

// func maskCmp8Mask(k1 uint16, a M128i, b M128i, imm8 int) uint16
TEXT ·maskCmp8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVW $0, ret+44(FP)
	RET

// func cmpEpu16Mask(a M128i, b M128i, imm8 int) uint8
TEXT ·cmpEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func maskCmpEpu16Mask(k1 uint8, a M128i, b M128i, imm8 int) uint8
TEXT ·maskCmpEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+44(FP)
	RET

// func cmpEpu32Mask(a M128i, b M128i, imm8 uint8) uint8
TEXT ·cmpEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu32Mask(k1 uint8, a M128i, b M128i, imm8 uint8) uint8
TEXT ·maskCmpEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmpEpu64Mask(a M128i, b M128i, imm8 uint8) uint8
TEXT ·cmpEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVB imm8+32(FP),R10

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func maskCmpEpu64Mask(k1 uint8, a M128i, b M128i, imm8 uint8) uint8
TEXT ·maskCmpEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVB imm8+36(FP),R11

	//TODO: Code missing

	MOVB $0, ret+40(FP)
	RET

// func cmpEpu8Mask(a M128i, b M128i, imm8 int) uint16
TEXT ·cmpEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVW $0, ret+40(FP)
	RET

// func maskCmpEpu8Mask(k1 uint16, a M128i, b M128i, imm8 int) uint16
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

// func cmpeq16(a M128i, b M128i) M128i
TEXT ·cmpeq16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeq16Mask(a M128i, b M128i) uint8
TEXT ·cmpeq16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeq16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpeq16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeq32(a M128i, b M128i) M128i
TEXT ·cmpeq32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeq32Mask(a M128i, b M128i) uint8
TEXT ·cmpeq32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeq32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpeq32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeq64(a M128i, b M128i) M128i
TEXT ·cmpeq64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeq64Mask(a M128i, b M128i) uint8
TEXT ·cmpeq64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeq64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpeq64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeq8(a M128i, b M128i) M128i
TEXT ·cmpeq8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeq8Mask(a M128i, b M128i) uint16
TEXT ·cmpeq8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpeq8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskCmpeq8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func cmpeqEpu16Mask(a M128i, b M128i) uint8
TEXT ·cmpeqEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpeqEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpu32Mask(a M128i, b M128i) uint8
TEXT ·cmpeqEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpeqEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpu64Mask(a M128i, b M128i) uint8
TEXT ·cmpeqEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpeqEpu64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpeqEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpeqEpu8Mask(a M128i, b M128i) uint16
TEXT ·cmpeqEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpeqEpu8Mask(k1 uint16, a M128i, b M128i) uint16
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

// func cmpge16Mask(a M128i, b M128i) uint8
TEXT ·cmpge16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpge16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpge16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpge32Mask(a M128i, b M128i) uint8
TEXT ·cmpge32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpge32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpge32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpge64Mask(a M128i, b M128i) uint8
TEXT ·cmpge64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpge64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpge64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpge8Mask(a M128i, b M128i) uint16
TEXT ·cmpge8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpge8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskCmpge8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func cmpgeEpu16Mask(a M128i, b M128i) uint8
TEXT ·cmpgeEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgeEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpu32Mask(a M128i, b M128i) uint8
TEXT ·cmpgeEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgeEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpu64Mask(a M128i, b M128i) uint8
TEXT ·cmpgeEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgeEpu64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgeEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgeEpu8Mask(a M128i, b M128i) uint16
TEXT ·cmpgeEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgeEpu8Mask(k1 uint16, a M128i, b M128i) uint16
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

// func cmpgt16(a M128i, b M128i) M128i
TEXT ·cmpgt16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgt16Mask(a M128i, b M128i) uint8
TEXT ·cmpgt16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgt16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgt16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgt32(a M128i, b M128i) M128i
TEXT ·cmpgt32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgt32Mask(a M128i, b M128i) uint8
TEXT ·cmpgt32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgt32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgt32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgt64(a M128i, b M128i) M128i
TEXT ·cmpgt64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgt64Mask(a M128i, b M128i) uint8
TEXT ·cmpgt64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgt64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgt64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgt8(a M128i, b M128i) M128i
TEXT ·cmpgt8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgt8Mask(a M128i, b M128i) uint16
TEXT ·cmpgt8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgt8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskCmpgt8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func cmpgtEpu16Mask(a M128i, b M128i) uint8
TEXT ·cmpgtEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgtEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpu32Mask(a M128i, b M128i) uint8
TEXT ·cmpgtEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgtEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpu64Mask(a M128i, b M128i) uint8
TEXT ·cmpgtEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpgtEpu64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpgtEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpgtEpu8Mask(a M128i, b M128i) uint16
TEXT ·cmpgtEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpgtEpu8Mask(k1 uint16, a M128i, b M128i) uint16
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

// func cmple16Mask(a M128i, b M128i) uint8
TEXT ·cmple16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmple16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmple16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmple32Mask(a M128i, b M128i) uint8
TEXT ·cmple32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmple32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmple32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmple64Mask(a M128i, b M128i) uint8
TEXT ·cmple64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmple64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmple64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmple8Mask(a M128i, b M128i) uint16
TEXT ·cmple8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmple8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskCmple8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func cmpleEpu16Mask(a M128i, b M128i) uint8
TEXT ·cmpleEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpleEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpu32Mask(a M128i, b M128i) uint8
TEXT ·cmpleEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpleEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpu64Mask(a M128i, b M128i) uint8
TEXT ·cmpleEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpleEpu64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpleEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpleEpu8Mask(a M128i, b M128i) uint16
TEXT ·cmpleEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpleEpu8Mask(k1 uint16, a M128i, b M128i) uint16
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

// func cmplt16(a M128i, b M128i) M128i
TEXT ·cmplt16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmplt16Mask(a M128i, b M128i) uint8
TEXT ·cmplt16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmplt16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmplt16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmplt32(a M128i, b M128i) M128i
TEXT ·cmplt32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmplt32Mask(a M128i, b M128i) uint8
TEXT ·cmplt32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmplt32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmplt32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmplt64Mask(a M128i, b M128i) uint8
TEXT ·cmplt64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmplt64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmplt64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmplt8(a M128i, b M128i) M128i
TEXT ·cmplt8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmplt8Mask(a M128i, b M128i) uint16
TEXT ·cmplt8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmplt8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskCmplt8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func cmpltEpu16Mask(a M128i, b M128i) uint8
TEXT ·cmpltEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpltEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpu32Mask(a M128i, b M128i) uint8
TEXT ·cmpltEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpltEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpu64Mask(a M128i, b M128i) uint8
TEXT ·cmpltEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpltEpu64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpltEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpltEpu8Mask(a M128i, b M128i) uint16
TEXT ·cmpltEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpltEpu8Mask(k1 uint16, a M128i, b M128i) uint16
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

// func cmpneq16Mask(a M128i, b M128i) uint8
TEXT ·cmpneq16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneq16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpneq16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneq32Mask(a M128i, b M128i) uint8
TEXT ·cmpneq32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneq32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpneq32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneq64Mask(a M128i, b M128i) uint8
TEXT ·cmpneq64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneq64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpneq64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneq8Mask(a M128i, b M128i) uint16
TEXT ·cmpneq8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpneq8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskCmpneq8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func cmpneqEpu16Mask(a M128i, b M128i) uint8
TEXT ·cmpneqEpu16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpneqEpu16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpu32Mask(a M128i, b M128i) uint8
TEXT ·cmpneqEpu32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpneqEpu32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpu64Mask(a M128i, b M128i) uint8
TEXT ·cmpneqEpu64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskCmpneqEpu64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskCmpneqEpu64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func cmpneqEpu8Mask(a M128i, b M128i) uint16
TEXT ·cmpneqEpu8Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVW $0, ret+32(FP)
	RET

// func maskCmpneqEpu8Mask(k1 uint16, a M128i, b M128i) uint16
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

// func maskCompress32(src M128i, k uint8, a M128i) M128i
TEXT ·maskCompress32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCompress32(k uint8, a M128i) M128i
TEXT ·maskzCompress32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCompress64(src M128i, k uint8, a M128i) M128i
TEXT ·maskCompress64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCompress64(k uint8, a M128i) M128i
TEXT ·maskzCompress64(SB),7,$0
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

// func maskCompressstoreu32(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCompressstoreu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCompressstoreu64(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCompressstoreu64(SB),7,$0
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

// func conflict32(a M128i) M128i
TEXT ·conflict32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskConflict32(src M128i, k uint8, a M128i) M128i
TEXT ·maskConflict32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzConflict32(k uint8, a M128i) M128i
TEXT ·maskzConflict32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func conflict64(a M128i) M128i
TEXT ·conflict64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskConflict64(src M128i, k uint8, a M128i) M128i
TEXT ·maskConflict64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzConflict64(k uint8, a M128i) M128i
TEXT ·maskzConflict64(SB),7,$0
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

// func maskCvtRoundpsPh(src M128i, k uint8, a [4]float32, rounding int) M128i
TEXT ·maskCvtRoundpsPh(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzCvtRoundpsPh(k uint8, a [4]float32, rounding int) M128i
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

// func cvtepi1632(a M128i) M128i
TEXT ·cvtepi1632(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi1632(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi1632(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi1632(k uint8, a M128i) M128i
TEXT ·maskzCvtepi1632(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi1664(a M128i) M128i
TEXT ·cvtepi1664(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi1664(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi1664(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi1664(k uint8, a M128i) M128i
TEXT ·maskzCvtepi1664(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi168(a M128i) M128i
TEXT ·cvtepi168(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi168(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi168(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi168(k uint8, a M128i) M128i
TEXT ·maskzCvtepi168(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi16Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtepi16Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtepi3216(a M128i) M128i
TEXT ·cvtepi3216(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi3216(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi3216(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi3216(k uint8, a M128i) M128i
TEXT ·maskzCvtepi3216(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi3264(a M128i) M128i
TEXT ·cvtepi3264(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi3264(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi3264(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi3264(k uint8, a M128i) M128i
TEXT ·maskzCvtepi3264(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi328(a M128i) M128i
TEXT ·cvtepi328(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi328(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi328(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi328(k uint8, a M128i) M128i
TEXT ·maskzCvtepi328(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi32Pd(a M128i) [2]float64
TEXT ·cvtepi32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi32Pd(src [2]float64, k uint8, a M128i) [2]float64
TEXT ·maskCvtepi32Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Pd(k uint8, a M128i) [2]float64
TEXT ·maskzCvtepi32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi32Ps(a M128i) [4]float32
TEXT ·cvtepi32Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi32Ps(src [4]float32, k uint8, a M128i) [4]float32
TEXT ·maskCvtepi32Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi32Ps(k uint8, a M128i) [4]float32
TEXT ·maskzCvtepi32Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi32Storeu16(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtepi32Storeu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi32Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtepi32Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtepi6416(a M128i) M128i
TEXT ·cvtepi6416(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi6416(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi6416(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi6416(k uint8, a M128i) M128i
TEXT ·maskzCvtepi6416(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi6432(a M128i) M128i
TEXT ·cvtepi6432(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi6432(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi6432(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi6432(k uint8, a M128i) M128i
TEXT ·maskzCvtepi6432(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi648(a M128i) M128i
TEXT ·cvtepi648(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi648(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi648(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi648(k uint8, a M128i) M128i
TEXT ·maskzCvtepi648(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Pd(a M128i) [2]float64
TEXT ·cvtepi64Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Pd(src [2]float64, k uint8, a M128i) [2]float64
TEXT ·maskCvtepi64Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Pd(k uint8, a M128i) [2]float64
TEXT ·maskzCvtepi64Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi64Ps(a M128i) [4]float32
TEXT ·cvtepi64Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi64Ps(src [4]float32, k uint8, a M128i) [4]float32
TEXT ·maskCvtepi64Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi64Ps(k uint8, a M128i) [4]float32
TEXT ·maskzCvtepi64Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtepi64Storeu16(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtepi64Storeu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64Storeu32(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtepi64Storeu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtepi64Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtepi64Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtepi816(a M128i) M128i
TEXT ·cvtepi816(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi816(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi816(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi816(k uint8, a M128i) M128i
TEXT ·maskzCvtepi816(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi832(a M128i) M128i
TEXT ·cvtepi832(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi832(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi832(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi832(k uint8, a M128i) M128i
TEXT ·maskzCvtepi832(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepi864(a M128i) M128i
TEXT ·cvtepi864(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepi864(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepi864(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepi864(k uint8, a M128i) M128i
TEXT ·maskzCvtepi864(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu1632(a M128i) M128i
TEXT ·cvtepu1632(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu1632(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepu1632(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu1632(k uint8, a M128i) M128i
TEXT ·maskzCvtepu1632(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu1664(a M128i) M128i
TEXT ·cvtepu1664(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu1664(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepu1664(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu1664(k uint8, a M128i) M128i
TEXT ·maskzCvtepu1664(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu3264(a M128i) M128i
TEXT ·cvtepu3264(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu3264(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepu3264(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu3264(k uint8, a M128i) M128i
TEXT ·maskzCvtepu3264(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu32Pd(a M128i) [2]float64
TEXT ·cvtepu32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu32Pd(src [2]float64, k uint8, a M128i) [2]float64
TEXT ·maskCvtepu32Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu32Pd(k uint8, a M128i) [2]float64
TEXT ·maskzCvtepu32Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu64Pd(a M128i) [2]float64
TEXT ·cvtepu64Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu64Pd(src [2]float64, k uint8, a M128i) [2]float64
TEXT ·maskCvtepu64Pd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu64Pd(k uint8, a M128i) [2]float64
TEXT ·maskzCvtepu64Pd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu64Ps(a M128i) [4]float32
TEXT ·cvtepu64Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu64Ps(src [4]float32, k uint8, a M128i) [4]float32
TEXT ·maskCvtepu64Ps(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu64Ps(k uint8, a M128i) [4]float32
TEXT ·maskzCvtepu64Ps(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu816(a M128i) M128i
TEXT ·cvtepu816(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu816(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepu816(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu816(k uint8, a M128i) M128i
TEXT ·maskzCvtepu816(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu832(a M128i) M128i
TEXT ·cvtepu832(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu832(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepu832(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu832(k uint8, a M128i) M128i
TEXT ·maskzCvtepu832(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtepu864(a M128i) M128i
TEXT ·cvtepu864(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtepu864(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtepu864(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtepu864(k uint8, a M128i) M128i
TEXT ·maskzCvtepu864(SB),7,$0
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

// func cvtpd32(a [2]float64) M128i
TEXT ·cvtpd32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpd32(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvtpd32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpd32(k uint8, a [2]float64) M128i
TEXT ·maskzCvtpd32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpd64(a [2]float64) M128i
TEXT ·cvtpd64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpd64(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvtpd64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpd64(k uint8, a [2]float64) M128i
TEXT ·maskzCvtpd64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpdEpu32(a [2]float64) M128i
TEXT ·cvtpdEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpu32(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvtpdEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpu32(k uint8, a [2]float64) M128i
TEXT ·maskzCvtpdEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpdEpu64(a [2]float64) M128i
TEXT ·cvtpdEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpdEpu64(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvtpdEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpdEpu64(k uint8, a [2]float64) M128i
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

// func cvtphPs(a M128i) [4]float32
TEXT ·cvtphPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtphPs(src [4]float32, k uint8, a M128i) [4]float32
TEXT ·maskCvtphPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtphPs(k uint8, a M128i) [4]float32
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

// func cvtps32(a [4]float32) M128i
TEXT ·cvtps32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtps32(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvtps32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtps32(k uint8, a [4]float32) M128i
TEXT ·maskzCvtps32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtps64(a [4]float32) M128i
TEXT ·cvtps64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtps64(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvtps64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtps64(k uint8, a [4]float32) M128i
TEXT ·maskzCvtps64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpsEpu32(a [4]float32) M128i
TEXT ·cvtpsEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpu32(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvtpsEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpu32(k uint8, a [4]float32) M128i
TEXT ·maskzCvtpsEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtpsEpu64(a [4]float32) M128i
TEXT ·cvtpsEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtpsEpu64(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvtpsEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtpsEpu64(k uint8, a [4]float32) M128i
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

// func cvtpsPh(a [4]float32, rounding int) M128i
TEXT ·cvtpsPh(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskCvtpsPh(src M128i, k uint8, a [4]float32, rounding int) M128i
TEXT ·maskCvtpsPh(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ rounding+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzCvtpsPh(k uint8, a [4]float32, rounding int) M128i
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

// func cvtsepi168(a M128i) M128i
TEXT ·cvtsepi168(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi168(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtsepi168(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi168(k uint8, a M128i) M128i
TEXT ·maskzCvtsepi168(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtsepi16Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtsepi16Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtsepi3216(a M128i) M128i
TEXT ·cvtsepi3216(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi3216(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtsepi3216(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi3216(k uint8, a M128i) M128i
TEXT ·maskzCvtsepi3216(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi328(a M128i) M128i
TEXT ·cvtsepi328(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi328(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtsepi328(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi328(k uint8, a M128i) M128i
TEXT ·maskzCvtsepi328(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtsepi32Storeu16(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtsepi32Storeu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi32Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtsepi32Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtsepi6416(a M128i) M128i
TEXT ·cvtsepi6416(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi6416(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtsepi6416(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi6416(k uint8, a M128i) M128i
TEXT ·maskzCvtsepi6416(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi6432(a M128i) M128i
TEXT ·cvtsepi6432(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi6432(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtsepi6432(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi6432(k uint8, a M128i) M128i
TEXT ·maskzCvtsepi6432(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtsepi648(a M128i) M128i
TEXT ·cvtsepi648(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtsepi648(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtsepi648(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtsepi648(k uint8, a M128i) M128i
TEXT ·maskzCvtsepi648(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtsepi64Storeu16(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtsepi64Storeu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64Storeu32(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtsepi64Storeu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtsepi64Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtsepi64Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtshSs(a uint16) float32
TEXT ·cvtshSs(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func cvtsi128Si32(a M128i) int
TEXT ·cvtsi128Si32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsi128Si64(a M128i) int64
TEXT ·cvtsi128Si64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func cvtsi128Si64x(a M128i) int64
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

// func cvtsi32Si128(a int) M128i
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

// func cvtsi64Si128(a int64) M128i
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

// func cvtsi64xSi128(a int64) M128i
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

// func cvttpd32(a [2]float64) M128i
TEXT ·cvttpd32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpd32(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvttpd32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpd32(k uint8, a [2]float64) M128i
TEXT ·maskzCvttpd32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpd64(a [2]float64) M128i
TEXT ·cvttpd64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpd64(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvttpd64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpd64(k uint8, a [2]float64) M128i
TEXT ·maskzCvttpd64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpdEpu32(a [2]float64) M128i
TEXT ·cvttpdEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpu32(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvttpdEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpu32(k uint8, a [2]float64) M128i
TEXT ·maskzCvttpdEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpdEpu64(a [2]float64) M128i
TEXT ·cvttpdEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpdEpu64(src M128i, k uint8, a [2]float64) M128i
TEXT ·maskCvttpdEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpdEpu64(k uint8, a [2]float64) M128i
TEXT ·maskzCvttpdEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttps32(a [4]float32) M128i
TEXT ·cvttps32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttps32(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvttps32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttps32(k uint8, a [4]float32) M128i
TEXT ·maskzCvttps32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttps64(a [4]float32) M128i
TEXT ·cvttps64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttps64(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvttps64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttps64(k uint8, a [4]float32) M128i
TEXT ·maskzCvttps64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpsEpu32(a [4]float32) M128i
TEXT ·cvttpsEpu32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpu32(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvttpsEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpu32(k uint8, a [4]float32) M128i
TEXT ·maskzCvttpsEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvttpsEpu64(a [4]float32) M128i
TEXT ·cvttpsEpu64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvttpsEpu64(src M128i, k uint8, a [4]float32) M128i
TEXT ·maskCvttpsEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvttpsEpu64(k uint8, a [4]float32) M128i
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

// func cvtusepi168(a M128i) M128i
TEXT ·cvtusepi168(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi168(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtusepi168(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi168(k uint8, a M128i) M128i
TEXT ·maskzCvtusepi168(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtusepi16Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtusepi16Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtusepi3216(a M128i) M128i
TEXT ·cvtusepi3216(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi3216(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtusepi3216(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi3216(k uint8, a M128i) M128i
TEXT ·maskzCvtusepi3216(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi328(a M128i) M128i
TEXT ·cvtusepi328(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi328(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtusepi328(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi328(k uint8, a M128i) M128i
TEXT ·maskzCvtusepi328(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtusepi32Storeu16(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtusepi32Storeu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi32Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtusepi32Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cvtusepi6416(a M128i) M128i
TEXT ·cvtusepi6416(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi6416(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtusepi6416(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi6416(k uint8, a M128i) M128i
TEXT ·maskzCvtusepi6416(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi6432(a M128i) M128i
TEXT ·cvtusepi6432(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi6432(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtusepi6432(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi6432(k uint8, a M128i) M128i
TEXT ·maskzCvtusepi6432(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func cvtusepi648(a M128i) M128i
TEXT ·cvtusepi648(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskCvtusepi648(src M128i, k uint8, a M128i) M128i
TEXT ·maskCvtusepi648(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzCvtusepi648(k uint8, a M128i) M128i
TEXT ·maskzCvtusepi648(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskCvtusepi64Storeu16(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtusepi64Storeu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64Storeu32(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtusepi64Storeu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskCvtusepi64Storeu8(base_addr uintptr, k uint8, a M128i) 
TEXT ·maskCvtusepi64Storeu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func dbsadEpu8(a M128i, b M128i, imm8 int) M128i
TEXT ·dbsadEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskDbsadEpu8(src M128i, k uint8, a M128i, b M128i, imm8 int) M128i
TEXT ·maskDbsadEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzDbsadEpu8(k uint8, a M128i, b M128i, imm8 int) M128i
TEXT ·maskzDbsadEpu8(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func div16(a M128i, b M128i) M128i
TEXT ·div16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func div32(a M128i, b M128i) M128i
TEXT ·div32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func div64(a M128i, b M128i) M128i
TEXT ·div64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func div8(a M128i, b M128i) M128i
TEXT ·div8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu16(a M128i, b M128i) M128i
TEXT ·divEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu32(a M128i, b M128i) M128i
TEXT ·divEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu64(a M128i, b M128i) M128i
TEXT ·divEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divEpu8(a M128i, b M128i) M128i
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

// func maskExpand32(src M128i, k uint8, a M128i) M128i
TEXT ·maskExpand32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzExpand32(k uint8, a M128i) M128i
TEXT ·maskzExpand32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskExpand64(src M128i, k uint8, a M128i) M128i
TEXT ·maskExpand64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzExpand64(k uint8, a M128i) M128i
TEXT ·maskzExpand64(SB),7,$0
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

// func maskExpandloadu32(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskExpandloadu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExpandloadu32(k uint8, mem_addr uintptr) M128i
TEXT ·maskzExpandloadu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskExpandloadu64(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskExpandloadu64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzExpandloadu64(k uint8, mem_addr uintptr) M128i
TEXT ·maskzExpandloadu64(SB),7,$0
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

// func extract16(a M128i, imm8 int) int
TEXT ·extract16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extract32(a M128i, imm8 int) int
TEXT ·extract32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extract64(a M128i, imm8 int) int64
TEXT ·extract64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func extract8(a M128i, imm8 int) int
TEXT ·extract8(SB),7,$0
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

// func fixupimmPd(a [2]float64, b [2]float64, c M128i, imm8 int) [2]float64
TEXT ·fixupimmPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmPd(a [2]float64, k uint8, b [2]float64, c M128i, imm8 int) [2]float64
TEXT ·maskFixupimmPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmPd(k uint8, a [2]float64, b [2]float64, c M128i, imm8 int) [2]float64
TEXT ·maskzFixupimmPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func fixupimmPs(a [4]float32, b [4]float32, c M128i, imm8 int) [4]float32
TEXT ·fixupimmPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmPs(a [4]float32, k uint8, b [4]float32, c M128i, imm8 int) [4]float32
TEXT ·maskFixupimmPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmPs(k uint8, a [4]float32, b [4]float32, c M128i, imm8 int) [4]float32
TEXT ·maskzFixupimmPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func fixupimmRoundSd(a [2]float64, b [2]float64, c M128i, imm8 int, rounding int) [2]float64
TEXT ·fixupimmRoundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11
	MOVQ rounding+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func maskFixupimmRoundSd(a [2]float64, k uint8, b [2]float64, c M128i, imm8 int, rounding int) [2]float64
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

// func maskzFixupimmRoundSd(k uint8, a [2]float64, b [2]float64, c M128i, imm8 int, rounding int) [2]float64
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

// func fixupimmRoundSs(a [4]float32, b [4]float32, c M128i, imm8 int, rounding int) [4]float32
TEXT ·fixupimmRoundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11
	MOVQ rounding+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func maskFixupimmRoundSs(a [4]float32, k uint8, b [4]float32, c M128i, imm8 int, rounding int) [4]float32
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

// func maskzFixupimmRoundSs(k uint8, a [4]float32, b [4]float32, c M128i, imm8 int, rounding int) [4]float32
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

// func fixupimmSd(a [2]float64, b [2]float64, c M128i, imm8 int) [2]float64
TEXT ·fixupimmSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmSd(a [2]float64, k uint8, b [2]float64, c M128i, imm8 int) [2]float64
TEXT ·maskFixupimmSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmSd(k uint8, a [2]float64, b [2]float64, c M128i, imm8 int) [2]float64
TEXT ·maskzFixupimmSd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func fixupimmSs(a [4]float32, b [4]float32, c M128i, imm8 int) [4]float32
TEXT ·fixupimmSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskFixupimmSs(a [4]float32, k uint8, b [4]float32, c M128i, imm8 int) [4]float32
TEXT ·maskFixupimmSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzFixupimmSs(k uint8, a [4]float32, b [4]float32, c M128i, imm8 int) [4]float32
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

// func hadd16(a M128i, b M128i) M128i
TEXT ·hadd16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hadd32(a M128i, b M128i) M128i
TEXT ·hadd32(SB),7,$0
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

// func hadds16(a M128i, b M128i) M128i
TEXT ·hadds16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsub16(a M128i, b M128i) M128i
TEXT ·hsub16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsub32(a M128i, b M128i) M128i
TEXT ·hsub32(SB),7,$0
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

// func hsubs16(a M128i, b M128i) M128i
TEXT ·hsubs16(SB),7,$0
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

// func i32gather32(base_addr int, vindex M128i, scale int) M128i
TEXT ·i32gather32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gather32(src M128i, base_addr int, vindex M128i, mask M128i, scale int) M128i
TEXT ·maskI32gather32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI32gather32(src M128i, k uint8, vindex M128i, base_addr uintptr, scale int) M128i
TEXT ·mmaskI32gather32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32gather64(base_addr int, vindex M128i, scale int) M128i
TEXT ·i32gather64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gather64(src M128i, base_addr int, vindex M128i, mask M128i, scale int) M128i
TEXT ·maskI32gather64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI32gather64(src M128i, k uint8, vindex M128i, base_addr uintptr, scale int) M128i
TEXT ·mmaskI32gather64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32gatherPd(base_addr float64, vindex M128i, scale int) [2]float64
TEXT ·i32gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherPd(src [2]float64, base_addr float64, vindex M128i, mask [2]float64, scale int) [2]float64
TEXT ·maskI32gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI32gatherPd(src [2]float64, k uint8, vindex M128i, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32gatherPs(base_addr float32, vindex M128i, scale int) [4]float32
TEXT ·i32gatherPs(SB),7,$0
	MOVL base_addr+0(FP),R8
	MOVOU vindex+4(FP),X1
	MOVQ scale+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskI32gatherPs(src [4]float32, base_addr float32, vindex M128i, mask [4]float32, scale int) [4]float32
TEXT ·maskI32gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVL base_addr+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVOU mask+36(FP),X3
	MOVQ scale+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mmaskI32gatherPs(src [4]float32, k uint8, vindex M128i, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI32gatherPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i32scatter32(base_addr uintptr, vindex M128i, a M128i, scale int) 
TEXT ·i32scatter32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatter32(base_addr uintptr, k uint8, vindex M128i, a M128i, scale int) 
TEXT ·maskI32scatter32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatter64(base_addr uintptr, vindex M128i, a M128i, scale int) 
TEXT ·i32scatter64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatter64(base_addr uintptr, k uint8, vindex M128i, a M128i, scale int) 
TEXT ·maskI32scatter64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPd(base_addr uintptr, vindex M128i, a [2]float64, scale int) 
TEXT ·i32scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPd(base_addr uintptr, k uint8, vindex M128i, a [2]float64, scale int) 
TEXT ·maskI32scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i32scatterPs(base_addr uintptr, vindex M128i, a [4]float32, scale int) 
TEXT ·i32scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI32scatterPs(base_addr uintptr, k uint8, vindex M128i, a [4]float32, scale int) 
TEXT ·maskI32scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64gather32(base_addr int, vindex M128i, scale int) M128i
TEXT ·i64gather32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gather32(src M128i, base_addr int, vindex M128i, mask M128i, scale int) M128i
TEXT ·maskI64gather32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI64gather32(src M128i, k uint8, vindex M128i, base_addr uintptr, scale int) M128i
TEXT ·mmaskI64gather32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gather64(base_addr int, vindex M128i, scale int) M128i
TEXT ·i64gather64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gather64(src M128i, base_addr int, vindex M128i, mask M128i, scale int) M128i
TEXT ·maskI64gather64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI64gather64(src M128i, k uint8, vindex M128i, base_addr uintptr, scale int) M128i
TEXT ·mmaskI64gather64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherPd(base_addr float64, vindex M128i, scale int) [2]float64
TEXT ·i64gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherPd(src [2]float64, base_addr float64, vindex M128i, mask [2]float64, scale int) [2]float64
TEXT ·maskI64gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func mmaskI64gatherPd(src [2]float64, k uint8, vindex M128i, base_addr uintptr, scale int) [2]float64
TEXT ·mmaskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherPs(base_addr float32, vindex M128i, scale int) [4]float32
TEXT ·i64gatherPs(SB),7,$0
	MOVL base_addr+0(FP),R8
	MOVOU vindex+4(FP),X1
	MOVQ scale+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskI64gatherPs(src [4]float32, base_addr float32, vindex M128i, mask [4]float32, scale int) [4]float32
TEXT ·maskI64gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVL base_addr+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVOU mask+36(FP),X3
	MOVQ scale+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func mmaskI64gatherPs(src [4]float32, k uint8, vindex M128i, base_addr uintptr, scale int) [4]float32
TEXT ·mmaskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64scatter32(base_addr uintptr, vindex M128i, a M128i, scale int) 
TEXT ·i64scatter32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatter32(base_addr uintptr, k uint8, vindex M128i, a M128i, scale int) 
TEXT ·maskI64scatter32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatter64(base_addr uintptr, vindex M128i, a M128i, scale int) 
TEXT ·i64scatter64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatter64(base_addr uintptr, k uint8, vindex M128i, a M128i, scale int) 
TEXT ·maskI64scatter64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPd(base_addr uintptr, vindex M128i, a [2]float64, scale int) 
TEXT ·i64scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPd(base_addr uintptr, k uint8, vindex M128i, a [2]float64, scale int) 
TEXT ·maskI64scatterPd(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func i64scatterPs(base_addr uintptr, vindex M128i, a [4]float32, scale int) 
TEXT ·i64scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskI64scatterPs(base_addr uintptr, k uint8, vindex M128i, a [4]float32, scale int) 
TEXT ·maskI64scatterPs(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func idiv32(a M128i, b M128i) M128i
TEXT ·idiv32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func idivrem32(mem_addr M128i, a M128i, b M128i) M128i
TEXT ·idivrem32(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func insert16(a M128i, i int, imm8 int) M128i
TEXT ·insert16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insert32(a M128i, i int, imm8 int) M128i
TEXT ·insert32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insert64(a M128i, i int64, imm8 int) M128i
TEXT ·insert64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func insert8(a M128i, i int, imm8 int) M128i
TEXT ·insert8(SB),7,$0
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

// func irem32(a M128i, b M128i) M128i
TEXT ·irem32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func lddquSi128(mem_addr M128iConst) M128i
TEXT ·lddquSi128(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoad32(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskLoad32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoad32(k uint8, mem_addr uintptr) M128i
TEXT ·maskzLoad32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoad64(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskLoad64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoad64(k uint8, mem_addr uintptr) M128i
TEXT ·maskzLoad64(SB),7,$0
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

// func loadSi128(mem_addr M128iConst) M128i
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

// func loadl64(mem_addr M128iConst) M128i
TEXT ·loadl64(SB),7,$0
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

// func maskLoadu16(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskLoadu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadu16(k uint8, mem_addr uintptr) M128i
TEXT ·maskzLoadu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadu32(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskLoadu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadu32(k uint8, mem_addr uintptr) M128i
TEXT ·maskzLoadu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadu64(src M128i, k uint8, mem_addr uintptr) M128i
TEXT ·maskLoadu64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadu64(k uint8, mem_addr uintptr) M128i
TEXT ·maskzLoadu64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskLoadu8(src M128i, k uint16, mem_addr uintptr) M128i
TEXT ·maskLoadu8(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzLoadu8(k uint16, mem_addr uintptr) M128i
TEXT ·maskzLoadu8(SB),7,$0
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

// func loaduSi128(mem_addr M128iConst) M128i
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

// func lzcnt32(a M128i) M128i
TEXT ·lzcnt32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskLzcnt32(src M128i, k uint8, a M128i) M128i
TEXT ·maskLzcnt32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzLzcnt32(k uint8, a M128i) M128i
TEXT ·maskzLzcnt32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func lzcnt64(a M128i) M128i
TEXT ·lzcnt64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskLzcnt64(src M128i, k uint8, a M128i) M128i
TEXT ·maskLzcnt64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzLzcnt64(k uint8, a M128i) M128i
TEXT ·maskzLzcnt64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func madd16(a M128i, b M128i) M128i
TEXT ·madd16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMadd16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMadd16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMadd16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMadd16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func madd52hiEpu64(a M128i, b M128i, c M128i) M128i
TEXT ·madd52hiEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskMadd52hiEpu64(a M128i, k uint8, b M128i, c M128i) M128i
TEXT ·maskMadd52hiEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMadd52hiEpu64(k uint8, a M128i, b M128i, c M128i) M128i
TEXT ·maskzMadd52hiEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func madd52loEpu64(a M128i, b M128i, c M128i) M128i
TEXT ·madd52loEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskMadd52loEpu64(a M128i, k uint8, b M128i, c M128i) M128i
TEXT ·maskMadd52loEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMadd52loEpu64(k uint8, a M128i, b M128i, c M128i) M128i
TEXT ·maskzMadd52loEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maddubs16(a M128i, b M128i) M128i
TEXT ·maddubs16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaddubs16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMaddubs16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaddubs16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMaddubs16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskload32(mem_addr int, mask M128i) M128i
TEXT ·maskload32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskload64(mem_addr int, mask M128i) M128i
TEXT ·maskload64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskloadPd(mem_addr float64, mask M128i) [2]float64
TEXT ·maskloadPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskloadPs(mem_addr float32, mask M128i) [4]float32
TEXT ·maskloadPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU mask+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskmoveuSi128(a M128i, mask M128i, mem_addr byte) 
TEXT ·maskmoveuSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1
	MOVB mem_addr+32(FP),R10

	//TODO: Code missing

	RET

// func maskstore32(mem_addr int, mask M128i, a M128i) 
TEXT ·maskstore32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func maskstore64(mem_addr int64, mask M128i, a M128i) 
TEXT ·maskstore64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func maskstorePd(mem_addr float64, mask M128i, a [2]float64) 
TEXT ·maskstorePd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func maskstorePs(mem_addr float32, mask M128i, a [4]float32) 
TEXT ·maskstorePs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU mask+4(FP),X1
	MOVOU a+20(FP),X2

	//TODO: Code missing

	RET

// func maskMax16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMax16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMax16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMax16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func max16(a M128i, b M128i) M128i
TEXT ·max16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMax32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMax32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMax32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMax32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func max32(a M128i, b M128i) M128i
TEXT ·max32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMax64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMax64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMax64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMax64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func max64(a M128i, b M128i) M128i
TEXT ·max64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMax8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskMax8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMax8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzMax8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func max8(a M128i, b M128i) M128i
TEXT ·max8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaxEpu16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMaxEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMaxEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maxEpu16(a M128i, b M128i) M128i
TEXT ·maxEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaxEpu32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMaxEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMaxEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maxEpu32(a M128i, b M128i) M128i
TEXT ·maxEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaxEpu64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMaxEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMaxEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maxEpu64(a M128i, b M128i) M128i
TEXT ·maxEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMaxEpu8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskMaxEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMaxEpu8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzMaxEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maxEpu8(a M128i, b M128i) M128i
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

// func maskMin16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMin16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMin16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMin16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func min16(a M128i, b M128i) M128i
TEXT ·min16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMin32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMin32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMin32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMin32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func min32(a M128i, b M128i) M128i
TEXT ·min32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMin64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMin64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMin64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMin64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func min64(a M128i, b M128i) M128i
TEXT ·min64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMin8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskMin8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMin8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzMin8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func min8(a M128i, b M128i) M128i
TEXT ·min8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMinEpu16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMinEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMinEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func minEpu16(a M128i, b M128i) M128i
TEXT ·minEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMinEpu32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMinEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMinEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func minEpu32(a M128i, b M128i) M128i
TEXT ·minEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMinEpu64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMinEpu64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMinEpu64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func minEpu64(a M128i, b M128i) M128i
TEXT ·minEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMinEpu8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskMinEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMinEpu8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzMinEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func minEpu8(a M128i, b M128i) M128i
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

// func minposEpu16(a M128i) M128i
TEXT ·minposEpu16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func maskMov16(src M128i, k uint8, a M128i) M128i
TEXT ·maskMov16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMov16(k uint8, a M128i) M128i
TEXT ·maskzMov16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMov32(src M128i, k uint8, a M128i) M128i
TEXT ·maskMov32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMov32(k uint8, a M128i) M128i
TEXT ·maskzMov32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMov64(src M128i, k uint8, a M128i) M128i
TEXT ·maskMov64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMov64(k uint8, a M128i) M128i
TEXT ·maskzMov64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskMov8(src M128i, k uint16, a M128i) M128i
TEXT ·maskMov8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskzMov8(k uint16, a M128i) M128i
TEXT ·maskzMov8(SB),7,$0
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

// func move64(a M128i) M128i
TEXT ·move64(SB),7,$0
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

// func movemask8(a M128i) int
TEXT ·movemask8(SB),7,$0
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

// func movm16(k uint8) M128i
TEXT ·movm16(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movm32(k uint8) M128i
TEXT ·movm32(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movm64(k uint8) M128i
TEXT ·movm64(SB),7,$0
	MOVB k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movm8(k uint16) M128i
TEXT ·movm8(SB),7,$0
	MOVW k+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func movpi6464(a M64) M128i
TEXT ·movpi6464(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func mpsadbwEpu8(a M128i, b M128i, imm8 int) M128i
TEXT ·mpsadbwEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskMul32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMul32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMul32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMul32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mul32(a M128i, b M128i) M128i
TEXT ·mul32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMulEpu32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMulEpu32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulEpu32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMulEpu32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mulEpu32(a M128i, b M128i) M128i
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

// func maskMulhi16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMulhi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulhi16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMulhi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mulhi16(a M128i, b M128i) M128i
TEXT ·mulhi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMulhiEpu16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMulhiEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulhiEpu16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMulhiEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mulhiEpu16(a M128i, b M128i) M128i
TEXT ·mulhiEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMulhrs16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMulhrs16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMulhrs16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMulhrs16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mulhrs16(a M128i, b M128i) M128i
TEXT ·mulhrs16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMullo16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMullo16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMullo16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMullo16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mullo16(a M128i, b M128i) M128i
TEXT ·mullo16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMullo32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMullo32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMullo32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMullo32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mullo32(a M128i, b M128i) M128i
TEXT ·mullo32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMullo64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskMullo64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMullo64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzMullo64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func mullo64(a M128i, b M128i) M128i
TEXT ·mullo64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskMultishift64Epi8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskMultishift64Epi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzMultishift64Epi8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzMultishift64Epi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func multishift64Epi8(a M128i, b M128i) M128i
TEXT ·multishift64Epi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskOr32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskOr32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOr32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzOr32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskOr64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskOr64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzOr64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzOr64(SB),7,$0
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

// func orSi128(a M128i, b M128i) M128i
TEXT ·orSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPacks16(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskPacks16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPacks16(k uint16, a M128i, b M128i) M128i
TEXT ·maskzPacks16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func packs16(a M128i, b M128i) M128i
TEXT ·packs16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPacks32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskPacks32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPacks32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzPacks32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func packs32(a M128i, b M128i) M128i
TEXT ·packs32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPackus16(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskPackus16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPackus16(k uint16, a M128i, b M128i) M128i
TEXT ·maskzPackus16(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func packus16(a M128i, b M128i) M128i
TEXT ·packus16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPackus32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskPackus32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPackus32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzPackus32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func packus32(a M128i, b M128i) M128i
TEXT ·packus32(SB),7,$0
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

// func maskPermutevarPd(src [2]float64, k uint8, a [2]float64, b M128i) [2]float64
TEXT ·maskPermutevarPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutevarPd(k uint8, a [2]float64, b M128i) [2]float64
TEXT ·maskzPermutevarPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func permutevarPd(a [2]float64, b M128i) [2]float64
TEXT ·permutevarPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPermutevarPs(src [4]float32, k uint8, a [4]float32, b M128i) [4]float32
TEXT ·maskPermutevarPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutevarPs(k uint8, a [4]float32, b M128i) [4]float32
TEXT ·maskzPermutevarPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func permutevarPs(a [4]float32, b M128i) [4]float32
TEXT ·permutevarPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPermutex2var16(a M128i, k uint8, idx M128i, b M128i) M128i
TEXT ·maskPermutex2var16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2var16(a M128i, idx M128i, k uint8, b M128i) M128i
TEXT ·mask2Permutex2var16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2var16(k uint8, a M128i, idx M128i, b M128i) M128i
TEXT ·maskzPermutex2var16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2var16(a M128i, idx M128i, b M128i) M128i
TEXT ·permutex2var16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2var32(a M128i, k uint8, idx M128i, b M128i) M128i
TEXT ·maskPermutex2var32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2var32(a M128i, idx M128i, k uint8, b M128i) M128i
TEXT ·mask2Permutex2var32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2var32(k uint8, a M128i, idx M128i, b M128i) M128i
TEXT ·maskzPermutex2var32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2var32(a M128i, idx M128i, b M128i) M128i
TEXT ·permutex2var32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2var64(a M128i, k uint8, idx M128i, b M128i) M128i
TEXT ·maskPermutex2var64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2var64(a M128i, idx M128i, k uint8, b M128i) M128i
TEXT ·mask2Permutex2var64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2var64(k uint8, a M128i, idx M128i, b M128i) M128i
TEXT ·maskzPermutex2var64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2var64(a M128i, idx M128i, b M128i) M128i
TEXT ·permutex2var64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2var8(a M128i, k uint16, idx M128i, b M128i) M128i
TEXT ·maskPermutex2var8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2var8(a M128i, idx M128i, k uint16, b M128i) M128i
TEXT ·mask2Permutex2var8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVW k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2var8(k uint16, a M128i, idx M128i, b M128i) M128i
TEXT ·maskzPermutex2var8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2var8(a M128i, idx M128i, b M128i) M128i
TEXT ·permutex2var8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varPd(a [2]float64, k uint8, idx M128i, b [2]float64) [2]float64
TEXT ·maskPermutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varPd(a [2]float64, idx M128i, k uint8, b [2]float64) [2]float64
TEXT ·mask2Permutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varPd(k uint8, a [2]float64, idx M128i, b [2]float64) [2]float64
TEXT ·maskzPermutex2varPd(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varPd(a [2]float64, idx M128i, b [2]float64) [2]float64
TEXT ·permutex2varPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutex2varPs(a [4]float32, k uint8, idx M128i, b [4]float32) [4]float32
TEXT ·maskPermutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func mask2Permutex2varPs(a [4]float32, idx M128i, k uint8, b [4]float32) [4]float32
TEXT ·mask2Permutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVB k+32(FP),R10
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutex2varPs(k uint8, a [4]float32, idx M128i, b [4]float32) [4]float32
TEXT ·maskzPermutex2varPs(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU idx+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func permutex2varPs(a [4]float32, idx M128i, b [4]float32) [4]float32
TEXT ·permutex2varPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU idx+16(FP),X1
	MOVOU b+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskPermutexvar16(src M128i, k uint8, idx M128i, a M128i) M128i
TEXT ·maskPermutexvar16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU a+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutexvar16(k uint8, idx M128i, a M128i) M128i
TEXT ·maskzPermutexvar16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU idx+4(FP),X1
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func permutexvar16(idx M128i, a M128i) M128i
TEXT ·permutexvar16(SB),7,$0
	MOVOU idx+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskPermutexvar8(src M128i, k uint16, idx M128i, a M128i) M128i
TEXT ·maskPermutexvar8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU idx+20(FP),X2
	MOVOU a+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzPermutexvar8(k uint16, idx M128i, a M128i) M128i
TEXT ·maskzPermutexvar8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU idx+4(FP),X1
	MOVOU a+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func permutexvar8(idx M128i, a M128i) M128i
TEXT ·permutexvar8(SB),7,$0
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

// func rem16(a M128i, b M128i) M128i
TEXT ·rem16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func rem32(a M128i, b M128i) M128i
TEXT ·rem32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func rem64(a M128i, b M128i) M128i
TEXT ·rem64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func rem8(a M128i, b M128i) M128i
TEXT ·rem8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu16(a M128i, b M128i) M128i
TEXT ·remEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu32(a M128i, b M128i) M128i
TEXT ·remEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu64(a M128i, b M128i) M128i
TEXT ·remEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func remEpu8(a M128i, b M128i) M128i
TEXT ·remEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRol32(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskRol32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRol32(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzRol32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func rol32(a M128i, imm8 int) M128i
TEXT ·rol32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRol64(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskRol64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRol64(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzRol64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func rol64(a M128i, imm8 int) M128i
TEXT ·rol64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRolv32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskRolv32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRolv32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzRolv32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rolv32(a M128i, b M128i) M128i
TEXT ·rolv32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRolv64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskRolv64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRolv64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzRolv64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rolv64(a M128i, b M128i) M128i
TEXT ·rolv64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRor32(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskRor32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRor32(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzRor32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func ror32(a M128i, imm8 int) M128i
TEXT ·ror32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRor64(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskRor64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzRor64(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzRor64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func ror64(a M128i, imm8 int) M128i
TEXT ·ror64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskRorv32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskRorv32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRorv32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzRorv32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rorv32(a M128i, b M128i) M128i
TEXT ·rorv32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskRorv64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskRorv64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzRorv64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzRorv64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func rorv64(a M128i, b M128i) M128i
TEXT ·rorv64(SB),7,$0
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

// func sadEpu8(a M128i, b M128i) M128i
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

// func set16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) M128i
TEXT ·set16(SB),7,$0
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

// func set32(e3 int, e2 int, e1 int, e0 int) M128i
TEXT ·set32(SB),7,$0
	MOVQ e3+0(FP),R8
	MOVQ e2+8(FP),R9
	MOVQ e1+16(FP),R10
	MOVQ e0+24(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func set64(e1 M64, e0 M64) M128i
TEXT ·set64(SB),7,$0
	MOVQ e1+0(FP),M0
	MOVQ e0+8(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func set64x(e1 int64, e0 int64) M128i
TEXT ·set64x(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func set8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M128i
TEXT ·set8(SB),7,$0
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

// func maskSet116(src M128i, k uint8, a int16) M128i
TEXT ·maskSet116(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVW a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskzSet116(k uint8, a int16) M128i
TEXT ·maskzSet116(SB),7,$0
	MOVB k+0(FP),R8
	MOVW a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func set116(a int16) M128i
TEXT ·set116(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func maskSet132(src M128i, k uint8, a int) M128i
TEXT ·maskSet132(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzSet132(k uint8, a int) M128i
TEXT ·maskzSet132(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func set132(a int) M128i
TEXT ·set132(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func maskSet164(src M128i, k uint8, a int64) M128i
TEXT ·maskSet164(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVQ a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskzSet164(k uint8, a int64) M128i
TEXT ·maskzSet164(SB),7,$0
	MOVB k+0(FP),R8
	MOVQ a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+12(FP)
	RET

// func set164(a M64) M128i
TEXT ·set164(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func set164x(a int64) M128i
TEXT ·set164x(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func maskSet18(src M128i, k uint16, a byte) M128i
TEXT ·maskSet18(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVB a+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskzSet18(k uint16, a byte) M128i
TEXT ·maskzSet18(SB),7,$0
	MOVW k+0(FP),R8
	MOVB a+4(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func set18(a byte) M128i
TEXT ·set18(SB),7,$0
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

// func setr16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) M128i
TEXT ·setr16(SB),7,$0
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

// func setr32(e3 int, e2 int, e1 int, e0 int) M128i
TEXT ·setr32(SB),7,$0
	MOVQ e3+0(FP),R8
	MOVQ e2+8(FP),R9
	MOVQ e1+16(FP),R10
	MOVQ e0+24(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func setr64(e1 M64, e0 M64) M128i
TEXT ·setr64(SB),7,$0
	MOVQ e1+0(FP),M0
	MOVQ e0+8(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func setr8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M128i
TEXT ·setr8(SB),7,$0
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

// func setzeroSi128() M128i
TEXT ·setzeroSi128(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func sha1msg1Epu32(a M128i, b M128i) M128i
TEXT ·sha1msg1Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha1msg2Epu32(a M128i, b M128i) M128i
TEXT ·sha1msg2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha1nexteEpu32(a M128i, b M128i) M128i
TEXT ·sha1nexteEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha1rnds4Epu32(a M128i, b M128i, fnc int) M128i
TEXT ·sha1rnds4Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ fnc+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func sha256msg1Epu32(a M128i, b M128i) M128i
TEXT ·sha256msg1Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha256msg2Epu32(a M128i, b M128i) M128i
TEXT ·sha256msg2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sha256rnds2Epu32(a M128i, b M128i, k M128i) M128i
TEXT ·sha256rnds2Epu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU k+32(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+48(FP)
	RET

// func maskShuffle32(src M128i, k uint8, a M128i, imm8 MMPERMENUM) M128i
TEXT ·maskShuffle32(SB),7,$0
	// Unimplemented. Unknown size of type MMPERMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskzShuffle32(k uint8, a M128i, imm8 MMPERMENUM) M128i
TEXT ·maskzShuffle32(SB),7,$0
	// Unimplemented. Unknown size of type MMPERMENUM
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func shuffle32(a M128i, imm8 int) M128i
TEXT ·shuffle32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskShuffle8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskShuffle8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzShuffle8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzShuffle8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func shuffle8(a M128i, b M128i) M128i
TEXT ·shuffle8(SB),7,$0
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

// func maskShufflehi16(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskShufflehi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzShufflehi16(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzShufflehi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func shufflehi16(a M128i, imm8 int) M128i
TEXT ·shufflehi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskShufflelo16(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskShufflelo16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzShufflelo16(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzShufflelo16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func shufflelo16(a M128i, imm8 int) M128i
TEXT ·shufflelo16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func sign16(a M128i, b M128i) M128i
TEXT ·sign16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sign32(a M128i, b M128i) M128i
TEXT ·sign32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sign8(a M128i, b M128i) M128i
TEXT ·sign8(SB),7,$0
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

// func maskSll16(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSll16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSll16(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSll16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sll16(a M128i, count M128i) M128i
TEXT ·sll16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSll32(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSll32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSll32(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSll32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sll32(a M128i, count M128i) M128i
TEXT ·sll32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSll64(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSll64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSll64(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSll64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sll64(a M128i, count M128i) M128i
TEXT ·sll64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSlli16(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSlli16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSlli16(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSlli16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func slli16(a M128i, imm8 int) M128i
TEXT ·slli16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSlli32(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSlli32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSlli32(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSlli32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func slli32(a M128i, imm8 int) M128i
TEXT ·slli32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSlli64(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSlli64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSlli64(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSlli64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func slli64(a M128i, imm8 int) M128i
TEXT ·slli64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func slliSi128(a M128i, imm8 int) M128i
TEXT ·slliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSllv16(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSllv16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllv16(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSllv16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sllv16(a M128i, count M128i) M128i
TEXT ·sllv16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSllv32(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSllv32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllv32(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSllv32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sllv32(a M128i, count M128i) M128i
TEXT ·sllv32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSllv64(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSllv64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSllv64(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSllv64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sllv64(a M128i, count M128i) M128i
TEXT ·sllv64(SB),7,$0
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

// func maskSra16(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSra16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSra16(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSra16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sra16(a M128i, count M128i) M128i
TEXT ·sra16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSra32(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSra32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSra32(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSra32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sra32(a M128i, count M128i) M128i
TEXT ·sra32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSra64(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSra64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSra64(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSra64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sra64(a M128i, count M128i) M128i
TEXT ·sra64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrai16(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSrai16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrai16(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSrai16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srai16(a M128i, imm8 int) M128i
TEXT ·srai16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrai32(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSrai32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrai32(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSrai32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srai32(a M128i, imm8 int) M128i
TEXT ·srai32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrai64(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSrai64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrai64(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSrai64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srai64(a M128i, imm8 uint32) M128i
TEXT ·srai64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVL imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+20(FP)
	RET

// func maskSrav16(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrav16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrav16(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrav16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srav16(a M128i, count M128i) M128i
TEXT ·srav16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrav32(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrav32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrav32(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrav32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srav32(a M128i, count M128i) M128i
TEXT ·srav32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrav64(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrav64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrav64(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrav64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srav64(a M128i, count M128i) M128i
TEXT ·srav64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrl16(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrl16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrl16(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrl16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srl16(a M128i, count M128i) M128i
TEXT ·srl16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrl32(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrl32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrl32(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrl32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srl32(a M128i, count M128i) M128i
TEXT ·srl32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrl64(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrl64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrl64(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrl64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srl64(a M128i, count M128i) M128i
TEXT ·srl64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrli16(src M128i, k uint8, a M128i, imm8 int) M128i
TEXT ·maskSrli16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVQ imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+44(FP)
	RET

// func maskzSrli16(k uint8, a M128i, imm8 int) M128i
TEXT ·maskzSrli16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVQ imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func srli16(a M128i, imm8 int) M128i
TEXT ·srli16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrli32(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSrli32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrli32(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSrli32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srli32(a M128i, imm8 int) M128i
TEXT ·srli32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrli64(src M128i, k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskSrli64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVL imm8+36(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func maskzSrli64(k uint8, a M128i, imm8 uint32) M128i
TEXT ·maskzSrli64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVL imm8+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srli64(a M128i, imm8 int) M128i
TEXT ·srli64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func srliSi128(a M128i, imm8 int) M128i
TEXT ·srliSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func maskSrlv16(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrlv16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlv16(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrlv16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srlv16(a M128i, count M128i) M128i
TEXT ·srlv16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrlv32(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrlv32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlv32(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrlv32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srlv32(a M128i, count M128i) M128i
TEXT ·srlv32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSrlv64(src M128i, k uint8, a M128i, count M128i) M128i
TEXT ·maskSrlv64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU count+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSrlv64(k uint8, a M128i, count M128i) M128i
TEXT ·maskzSrlv64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU count+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func srlv64(a M128i, count M128i) M128i
TEXT ·srlv64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskStore32(mem_addr uintptr, k uint8, a M128i) 
TEXT ·maskStore32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStore64(mem_addr uintptr, k uint8, a M128i) 
TEXT ·maskStore64(SB),7,$0
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

// func storeSi128(mem_addr M128i, a M128i) 
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

// func storel64(mem_addr M128i, a M128i) 
TEXT ·storel64(SB),7,$0
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

// func maskStoreu16(mem_addr uintptr, k uint8, a M128i) 
TEXT ·maskStoreu16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreu32(mem_addr uintptr, k uint8, a M128i) 
TEXT ·maskStoreu32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreu64(mem_addr uintptr, k uint8, a M128i) 
TEXT ·maskStoreu64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func maskStoreu8(mem_addr uintptr, k uint16, a M128i) 
TEXT ·maskStoreu8(SB),7,$0
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

// func storeuSi128(mem_addr M128i, a M128i) 
TEXT ·storeuSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func streamLoadSi128(mem_addr M128i) M128i
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

// func streamSi128(mem_addr M128i, a M128i) 
TEXT ·streamSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func maskSub16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskSub16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSub16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzSub16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sub16(a M128i, b M128i) M128i
TEXT ·sub16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSub32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskSub32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSub32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzSub32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sub32(a M128i, b M128i) M128i
TEXT ·sub32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSub64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskSub64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSub64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzSub64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sub64(a M128i, b M128i) M128i
TEXT ·sub64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSub8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskSub8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSub8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzSub8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func sub8(a M128i, b M128i) M128i
TEXT ·sub8(SB),7,$0
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

// func maskSubs16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskSubs16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubs16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzSubs16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func subs16(a M128i, b M128i) M128i
TEXT ·subs16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSubs8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskSubs8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubs8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzSubs8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func subs8(a M128i, b M128i) M128i
TEXT ·subs8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSubsEpu16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskSubsEpu16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubsEpu16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzSubsEpu16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func subsEpu16(a M128i, b M128i) M128i
TEXT ·subsEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskSubsEpu8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskSubsEpu8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzSubsEpu8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzSubsEpu8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func subsEpu8(a M128i, b M128i) M128i
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

// func maskTernarylogic32(src M128i, k uint8, a M128i, b M128i, imm8 int) M128i
TEXT ·maskTernarylogic32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzTernarylogic32(k uint8, a M128i, b M128i, c M128i, imm8 int) M128i
TEXT ·maskzTernarylogic32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func ternarylogic32(a M128i, b M128i, c M128i, imm8 int) M128i
TEXT ·ternarylogic32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskTernarylogic64(src M128i, k uint8, a M128i, b M128i, imm8 int) M128i
TEXT ·maskTernarylogic64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func maskzTernarylogic64(k uint8, a M128i, b M128i, c M128i, imm8 int) M128i
TEXT ·maskzTernarylogic64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2
	MOVOU c+36(FP),X3
	MOVQ imm8+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func ternarylogic64(a M128i, b M128i, c M128i, imm8 int) M128i
TEXT ·ternarylogic64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU c+32(FP),X2
	MOVQ imm8+48(FP),R11

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func maskTest16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskTest16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func test16Mask(a M128i, b M128i) uint8
TEXT ·test16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTest32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskTest32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func test32Mask(a M128i, b M128i) uint8
TEXT ·test32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTest64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskTest64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func test64Mask(a M128i, b M128i) uint8
TEXT ·test64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTest8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskTest8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func test8Mask(a M128i, b M128i) uint16
TEXT ·test8Mask(SB),7,$0
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

// func testcSi128(a M128i, b M128i) int
TEXT ·testcSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func maskTestn16Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskTestn16Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testn16Mask(a M128i, b M128i) uint8
TEXT ·testn16Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestn32Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskTestn32Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testn32Mask(a M128i, b M128i) uint8
TEXT ·testn32Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestn64Mask(k1 uint8, a M128i, b M128i) uint8
TEXT ·maskTestn64Mask(SB),7,$0
	MOVB k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVB $0, ret+36(FP)
	RET

// func testn64Mask(a M128i, b M128i) uint8
TEXT ·testn64Mask(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVB $0, ret+32(FP)
	RET

// func maskTestn8Mask(k1 uint16, a M128i, b M128i) uint16
TEXT ·maskTestn8Mask(SB),7,$0
	MOVW k1+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVW $0, ret+36(FP)
	RET

// func testn8Mask(a M128i, b M128i) uint16
TEXT ·testn8Mask(SB),7,$0
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

// func testnzcSi128(a M128i, b M128i) int
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

// func testzSi128(a M128i, b M128i) int
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

// func udiv32(a M128i, b M128i) M128i
TEXT ·udiv32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func udivrem32(mem_addr M128i, a M128i, b M128i) M128i
TEXT ·udivrem32(SB),7,$0
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

// func undefinedSi128() M128i
TEXT ·undefinedSi128(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskUnpackhi16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskUnpackhi16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhi16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzUnpackhi16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpackhi16(a M128i, b M128i) M128i
TEXT ·unpackhi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskUnpackhi32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskUnpackhi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhi32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzUnpackhi32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpackhi32(a M128i, b M128i) M128i
TEXT ·unpackhi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskUnpackhi64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskUnpackhi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhi64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzUnpackhi64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpackhi64(a M128i, b M128i) M128i
TEXT ·unpackhi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskUnpackhi8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskUnpackhi8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpackhi8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzUnpackhi8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpackhi8(a M128i, b M128i) M128i
TEXT ·unpackhi8(SB),7,$0
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

// func maskUnpacklo16(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskUnpacklo16(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpacklo16(k uint8, a M128i, b M128i) M128i
TEXT ·maskzUnpacklo16(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpacklo16(a M128i, b M128i) M128i
TEXT ·unpacklo16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskUnpacklo32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskUnpacklo32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpacklo32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzUnpacklo32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpacklo32(a M128i, b M128i) M128i
TEXT ·unpacklo32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskUnpacklo64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskUnpacklo64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpacklo64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzUnpacklo64(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpacklo64(a M128i, b M128i) M128i
TEXT ·unpacklo64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskUnpacklo8(src M128i, k uint16, a M128i, b M128i) M128i
TEXT ·maskUnpacklo8(SB),7,$0
	MOVOU src+0(FP),X0
	MOVW k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzUnpacklo8(k uint16, a M128i, b M128i) M128i
TEXT ·maskzUnpacklo8(SB),7,$0
	MOVW k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func unpacklo8(a M128i, b M128i) M128i
TEXT ·unpacklo8(SB),7,$0
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

// func urem32(a M128i, b M128i) M128i
TEXT ·urem32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskXor32(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskXor32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXor32(k uint8, a M128i, b M128i) M128i
TEXT ·maskzXor32(SB),7,$0
	MOVB k+0(FP),R8
	MOVOU a+4(FP),X1
	MOVOU b+20(FP),X2

	//TODO: Code missing

	MOVOU X0, ret+36(FP)
	RET

// func maskXor64(src M128i, k uint8, a M128i, b M128i) M128i
TEXT ·maskXor64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVB k+16(FP),R9
	MOVOU a+20(FP),X2
	MOVOU b+36(FP),X3

	//TODO: Code missing

	MOVOU X0, ret+52(FP)
	RET

// func maskzXor64(k uint8, a M128i, b M128i) M128i
TEXT ·maskzXor64(SB),7,$0
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

// func xorSi128(a M128i, b M128i) M128i
TEXT ·xorSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

