// func addEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·addEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addPd(a [2]float64, b [2]float64) [2]float64
TEXT ·addPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addSd(a [2]float64, b [2]float64) [2]float64
TEXT ·addSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addSi64(a M64, b M64) M64
TEXT ·addSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func addsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addsEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addsEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func addsEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·addsEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func andPd(a [2]float64, b [2]float64) [2]float64
TEXT ·andPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func andSi128(a [16]byte, b [16]byte) [16]byte
TEXT ·andSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func andnotPd(a [2]float64, b [2]float64) [2]float64
TEXT ·andnotPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func andnotSi128(a [16]byte, b [16]byte) [16]byte
TEXT ·andnotSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func avgEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·avgEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func avgEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·avgEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
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

// func clflush(p uintptr) 
TEXT ·clflush(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cmpeqEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeqEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeqEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpeqPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpeqPd(SB),7,$0
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

// func cmpgePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpgePd(SB),7,$0
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

// func cmpgtEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpgtPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpgtPd(SB),7,$0
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

// func cmplePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmplePd(SB),7,$0
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

// func cmpltEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpltEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpltEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpltEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cmpltPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpltPd(SB),7,$0
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

// func cmpneqPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpneqPd(SB),7,$0
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

// func cmpngePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpngePd(SB),7,$0
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

// func cmpngtPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpngtPd(SB),7,$0
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

// func cmpnlePd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpnlePd(SB),7,$0
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

// func cmpnltPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpnltPd(SB),7,$0
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

// func cmpordPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpordPd(SB),7,$0
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

// func cmpunordPd(a [2]float64, b [2]float64) [2]float64
TEXT ·cmpunordPd(SB),7,$0
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

// func comieqSd(a [2]float64, b [2]float64) int
TEXT ·comieqSd(SB),7,$0
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

// func comigtSd(a [2]float64, b [2]float64) int
TEXT ·comigtSd(SB),7,$0
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

// func comiltSd(a [2]float64, b [2]float64) int
TEXT ·comiltSd(SB),7,$0
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

// func cvtepi32Pd(a [16]byte) [2]float64
TEXT ·cvtepi32Pd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtepi32Ps(a [16]byte) [4]float32
TEXT ·cvtepi32Ps(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtpdEpi32(a [2]float64) [16]byte
TEXT ·cvtpdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtpdPi32(a [2]float64) M64
TEXT ·cvtpdPi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtpdPs(a [2]float64) [4]float32
TEXT ·cvtpdPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtpi32Pd(a M64) [2]float64
TEXT ·cvtpi32Pd(SB),7,$0
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

// func cvtpsPd(a [4]float32) [2]float64
TEXT ·cvtpsPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtsdF64(a [2]float64) float64
TEXT ·cvtsdF64(SB),7,$0
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

// func cvtssSd(a [2]float64, b [4]float32) [2]float64
TEXT ·cvtssSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func cvttpdEpi32(a [2]float64) [16]byte
TEXT ·cvttpdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvttpdPi32(a [2]float64) M64
TEXT ·cvttpdPi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvttpsEpi32(a [4]float32) [16]byte
TEXT ·cvttpsEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
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

// func divPd(a [2]float64, b [2]float64) [2]float64
TEXT ·divPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func divSd(a [2]float64, b [2]float64) [2]float64
TEXT ·divSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func extractEpi16(a [16]byte, imm8 int) int
TEXT ·extractEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+24(FP)
	RET

// func insertEpi16(a [16]byte, i int, imm8 int) [16]byte
TEXT ·insertEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func lfence() 
TEXT ·lfence(SB),7,$0

	//TODO: Code missing

	RET

// func loadPd(mem_addr float64) [2]float64
TEXT ·loadPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loadPd1(mem_addr float64) [2]float64
TEXT ·loadPd1(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loadSd(mem_addr float64) [2]float64
TEXT ·loadSd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loadSi128(mem_addr M128iConst) [16]byte
TEXT ·loadSi128(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func load1Pd(mem_addr float64) [2]float64
TEXT ·load1Pd(SB),7,$0
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

// func loaduPd(mem_addr float64) [2]float64
TEXT ·loaduPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loaduSi128(mem_addr M128iConst) [16]byte
TEXT ·loaduSi128(SB),7,$0
	// Unimplemented. Unknown size of type M128iConst
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maddEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·maddEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskmoveuSi128(a [16]byte, mask [16]byte, mem_addr byte) 
TEXT ·maskmoveuSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1
	MOVB mem_addr+32(FP),R10

	//TODO: Code missing

	RET

// func maxEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maxEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maxPd(a [2]float64, b [2]float64) [2]float64
TEXT ·maxPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maxSd(a [2]float64, b [2]float64) [2]float64
TEXT ·maxSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mfence() 
TEXT ·mfence(SB),7,$0

	//TODO: Code missing

	RET

// func minEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func minEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func minPd(a [2]float64, b [2]float64) [2]float64
TEXT ·minPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func minSd(a [2]float64, b [2]float64) [2]float64
TEXT ·minSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func moveEpi64(a [16]byte) [16]byte
TEXT ·moveEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func moveSd(a [2]float64, b [2]float64) [2]float64
TEXT ·moveSd(SB),7,$0
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

// func movepi64Pi64(a [16]byte) M64
TEXT ·movepi64Pi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func movpi64Epi64(a M64) [16]byte
TEXT ·movpi64Epi64(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func mulEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·mulEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mulPd(a [2]float64, b [2]float64) [2]float64
TEXT ·mulPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mulSd(a [2]float64, b [2]float64) [2]float64
TEXT ·mulSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mulSu32(a M64, b M64) M64
TEXT ·mulSu32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mulhiEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulhiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mulhiEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulhiEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func mulloEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·mulloEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func orPd(a [2]float64, b [2]float64) [2]float64
TEXT ·orPd(SB),7,$0
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

// func packsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·packsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func packsEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·packsEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func packusEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·packusEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func pause() 
TEXT ·pause(SB),7,$0

	//TODO: Code missing

	RET

// func sadEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·sadEpu8(SB),7,$0
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

// func setSd(a float64) [2]float64
TEXT ·setSd(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func set1Epi16(a int16) [16]byte
TEXT ·set1Epi16(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func set1Epi32(a int) [16]byte
TEXT ·set1Epi32(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
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

// func setzeroPd() [2]float64
TEXT ·setzeroPd(SB),7,$0

	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func setzeroSi128() [16]byte
TEXT ·setzeroSi128(SB),7,$0

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

// func shufflePd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·shufflePd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func shufflehiEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·shufflehiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func shuffleloEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·shuffleloEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func sllEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sllEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sllEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sllEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sllEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sllEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func slliEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·slliEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

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

// func sqrtPd(a [2]float64) [2]float64
TEXT ·sqrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sqrtSd(a [2]float64, b [2]float64) [2]float64
TEXT ·sqrtSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sraEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·sraEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sraEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sraEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func sraiEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·sraiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

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

// func srlEpi16(a [16]byte, count [16]byte) [16]byte
TEXT ·srlEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func srlEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·srlEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func srlEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·srlEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func srliEpi16(a [16]byte, imm8 int) [16]byte
TEXT ·srliEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

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

// func store1Pd(mem_addr float64, a [2]float64) 
TEXT ·store1Pd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

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

// func storeuPd(mem_addr float64, a [2]float64) 
TEXT ·storeuPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storeuSi128(mem_addr [16]byte, a [16]byte) 
TEXT ·storeuSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func streamPd(mem_addr float64, a [2]float64) 
TEXT ·streamPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func streamSi128(mem_addr [16]byte, a [16]byte) 
TEXT ·streamSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	//TODO: Code missing

	RET

// func streamSi32(mem_addr int, a int) 
TEXT ·streamSi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ a+8(FP),R9

	//TODO: Code missing

	RET

// func streamSi64(mem_addr int64, a int64) 
TEXT ·streamSi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ a+8(FP),R9

	//TODO: Code missing

	RET

// func subEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·subEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subPd(a [2]float64, b [2]float64) [2]float64
TEXT ·subPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subSd(a [2]float64, b [2]float64) [2]float64
TEXT ·subSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subSi64(a M64, b M64) M64
TEXT ·subSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subsEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subsEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subsEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func subsEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·subsEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func ucomieqSd(a [2]float64, b [2]float64) int
TEXT ·ucomieqSd(SB),7,$0
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

// func ucomigtSd(a [2]float64, b [2]float64) int
TEXT ·ucomigtSd(SB),7,$0
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

// func ucomiltSd(a [2]float64, b [2]float64) int
TEXT ·ucomiltSd(SB),7,$0
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

// func unpackhiEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackhiEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackhiEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackhiEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackhiEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackhiPd(a [2]float64, b [2]float64) [2]float64
TEXT ·unpackhiPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackloEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackloEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackloEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackloEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·unpackloEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func unpackloPd(a [2]float64, b [2]float64) [2]float64
TEXT ·unpackloPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func xorPd(a [2]float64, b [2]float64) [2]float64
TEXT ·xorPd(SB),7,$0
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

