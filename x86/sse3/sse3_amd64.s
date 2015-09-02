// func addsubPd(a [2]float64, b [2]float64) [2]float64
TEXT ·addsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ADDSUBPD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func addsubPs(a [4]float32, b [4]float32) [4]float32
TEXT ·addsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ADDSUBPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func haddPd(a [2]float64, b [2]float64) [2]float64
TEXT ·haddPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// HADDPD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func haddPs(a [4]float32, b [4]float32) [4]float32
TEXT ·haddPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// HADDPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func hsubPd(a [2]float64, b [2]float64) [2]float64
TEXT ·hsubPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// HSUBPD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func hsubPs(a [4]float32, b [4]float32) [4]float32
TEXT ·hsubPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// HSUBPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func lddquSi128(mem_addr x86.M128iConst) [16]byte
TEXT ·lddquSi128(SB),7,$0
	// Unimplemented. Unknown size of type x86.M128iConst
	// TODO: Code missing - could be:
	// LDDQU 

	MOVOU X0, ret+0(FP)
	RET

// func loaddupPd(mem_addr float64) [2]float64
TEXT ·loaddupPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing - could be:
	// MOVDDUP R8

	MOVOU X0, ret+8(FP)
	RET

// func movedupPd(a [2]float64) [2]float64
TEXT ·movedupPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// MOVDDUP X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func movehdupPs(a [4]float32) [4]float32
TEXT ·movehdupPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// MOVSHDUP X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func moveldupPs(a [4]float32) [4]float32
TEXT ·moveldupPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// MOVSLDUP X0, X0

	MOVOU X0, ret+16(FP)
	RET

