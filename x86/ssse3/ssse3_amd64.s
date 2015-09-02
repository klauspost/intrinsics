// func absEpi16(a [16]byte) [16]byte
TEXT ·absEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func absEpi32(a [16]byte) [16]byte
TEXT ·absEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func absEpi8(a [16]byte) [16]byte
TEXT ·absEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func absPi16(a M64) M64
TEXT ·absPi16(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	// Return size: 8
	RET

// func absPi32(a M64) M64
TEXT ·absPi32(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	// Return size: 8
	RET

// func absPi8(a M64) M64
TEXT ·absPi8(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	// Return size: 8
	RET

// func alignrEpi8(a [16]byte, b [16]byte, count int) [16]byte
TEXT ·alignrEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ count+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func alignrPi8(a M64, b M64, count int) M64
TEXT ·alignrPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1
	MOVQ count+16(FP),R10

	//TODO: Code missing

	// Return size: 8
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

// func haddPi16(a M64, b M64) M64
TEXT ·haddPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func haddPi32(a M64, b M64) M64
TEXT ·haddPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func haddsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·haddsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func haddsPi16(a M64, b M64) M64
TEXT ·haddsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
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

// func hsubPi16(a M64, b M64) M64
TEXT ·hsubPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hsubPi32(a M64, b M64) M64
TEXT ·hsubPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hsubsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·hsubsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func hsubsPi16(a M64, b M64) M64
TEXT ·hsubsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func maddubsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·maddubsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maddubsPi16(a M64, b M64) M64
TEXT ·maddubsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mulhrsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulhrsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mulhrsPi16(a M64, b M64) M64
TEXT ·mulhrsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func shuffleEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·shuffleEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func shufflePi8(a M64, b M64) M64
TEXT ·shufflePi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
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

// func signPi16(a M64, b M64) M64
TEXT ·signPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func signPi32(a M64, b M64) M64
TEXT ·signPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func signPi8(a M64, b M64) M64
TEXT ·signPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

