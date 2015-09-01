// func abs16(a M64) M64
TEXT ·abs16(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	// Return size: 8
	RET

// func abs32(a M64) M64
TEXT ·abs32(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	// Return size: 8
	RET

// func abs8(a M64) M64
TEXT ·abs8(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	// Return size: 8
	RET

// func addSi64(a M64, b M64) M64
TEXT ·addSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func alignr8(a M64, b M64, count int) M64
TEXT ·alignr8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1
	MOVQ count+16(FP),R10

	//TODO: Code missing

	// Return size: 8
	RET

// func avgPu16(a M64, b M64) M64
TEXT ·avgPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func avgPu8(a M64, b M64) M64
TEXT ·avgPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func cvt2ps(a [4]float32, b M64) [4]float32
TEXT ·cvt2ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtpd32(a [2]float64) M64
TEXT ·cvtpd32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtps16(a [4]float32) M64
TEXT ·cvtps16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtps32(a [4]float32) M64
TEXT ·cvtps32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtps8(a [4]float32) M64
TEXT ·cvtps8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvttpd32(a [2]float64) M64
TEXT ·cvttpd32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvttps32(a [4]float32) M64
TEXT ·cvttps32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func extract16(a M64, imm8 int) int
TEXT ·extract16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func hadd16(a M64, b M64) M64
TEXT ·hadd16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hadd32(a M64, b M64) M64
TEXT ·hadd32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hadds16(a M64, b M64) M64
TEXT ·hadds16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hsub16(a M64, b M64) M64
TEXT ·hsub16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hsub32(a M64, b M64) M64
TEXT ·hsub32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func hsubs16(a M64, b M64) M64
TEXT ·hsubs16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func insert16(a M64, i int, imm8 int) M64
TEXT ·insert16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ i+8(FP),R9
	MOVQ imm8+16(FP),R10

	//TODO: Code missing

	// Return size: 8
	RET

// func loadh(a [4]float32, mem_addr M64Const) [4]float32
TEXT ·loadh(SB),7,$0
	// Unimplemented. Unknown size of type M64Const
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loadl(a [4]float32, mem_addr M64Const) [4]float32
TEXT ·loadl(SB),7,$0
	// Unimplemented. Unknown size of type M64Const
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi16(mem_addr uintptr) M128i
TEXT ·loaduSi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi161(mem_addr uintptr) M128i
TEXT ·loaduSi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi32(mem_addr uintptr) M128i
TEXT ·loaduSi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi321(mem_addr uintptr) M128i
TEXT ·loaduSi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi64(mem_addr uintptr) M128i
TEXT ·loaduSi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi641(mem_addr uintptr) M128i
TEXT ·loaduSi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maddubs16(a M64, b M64) M64
TEXT ·maddubs16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func maskmoveSi64(a M64, mask M64, mem_addr byte) 
TEXT ·maskmoveSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ mask+8(FP),M1
	MOVB mem_addr+16(FP),R10

	//TODO: Code missing

	RET

// func mMaskmovq(a M64, mask M64, mem_addr byte) 
TEXT ·mMaskmovq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ mask+8(FP),M1
	MOVB mem_addr+16(FP),R10

	//TODO: Code missing

	RET

// func max16(a M64, b M64) M64
TEXT ·max16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func maxPu8(a M64, b M64) M64
TEXT ·maxPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func min16(a M64, b M64) M64
TEXT ·min16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func minPu8(a M64, b M64) M64
TEXT ·minPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func movemask8(a M64) int
TEXT ·movemask8(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func movepi6464(a M128i) M64
TEXT ·movepi6464(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func mulSu32(a M64, b M64) M64
TEXT ·mulSu32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mulhiPu16(a M64, b M64) M64
TEXT ·mulhiPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mulhrs16(a M64, b M64) M64
TEXT ·mulhrs16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPavgb(a M64, b M64) M64
TEXT ·mPavgb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPavgw(a M64, b M64) M64
TEXT ·mPavgw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPextrw(a M64, imm8 int) int
TEXT ·mPextrw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func mnsrw(a M64, i int, imm8 int) M64
TEXT ·mnsrw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ i+8(FP),R9
	MOVQ imm8+16(FP),R10

	//TODO: Code missing

	// Return size: 8
	RET

// func mPmaxsw(a M64, b M64) M64
TEXT ·mPmaxsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPmaxub(a M64, b M64) M64
TEXT ·mPmaxub(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPminsw(a M64, b M64) M64
TEXT ·mPminsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPminub(a M64, b M64) M64
TEXT ·mPminub(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPmovmskb(a M64) int
TEXT ·mPmovmskb(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func mPmulhuw(a M64, b M64) M64
TEXT ·mPmulhuw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPsadbw(a M64, b M64) M64
TEXT ·mPsadbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func mPshufw(a M64, imm8 int) M64
TEXT ·mPshufw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func sadPu8(a M64, b M64) M64
TEXT ·sadPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func shuffle16(a M64, imm8 int) M64
TEXT ·shuffle16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func shuffle8(a M64, b M64) M64
TEXT ·shuffle8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sign16(a M64, b M64) M64
TEXT ·sign16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sign32(a M64, b M64) M64
TEXT ·sign32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sign8(a M64, b M64) M64
TEXT ·sign8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func storeh(mem_addr M64, a [4]float32) 
TEXT ·storeh(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storel(mem_addr M64, a [4]float32) 
TEXT ·storel(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storeuSi16(mem_addr uintptr, a M128i) 
TEXT ·storeuSi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi161(mem_addr uintptr, a M128i) 
TEXT ·storeuSi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi32(mem_addr uintptr, a M128i) 
TEXT ·storeuSi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi321(mem_addr uintptr, a M128i) 
TEXT ·storeuSi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi64(mem_addr uintptr, a M128i) 
TEXT ·storeuSi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi641(mem_addr uintptr, a M128i) 
TEXT ·storeuSi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func stream(mem_addr M64, a M64) 
TEXT ·stream(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVQ a+8(FP),M1

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

// func subSi64(a M64, b M64) M64
TEXT ·subSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subborrowU32(b_in uint8, a uint32, b uint32, out uint32) uint8
TEXT ·subborrowU32(SB),7,$0
	MOVB b_in+0(FP),R8
	MOVL a+4(FP),R9
	MOVL b+8(FP),R10
	MOVL out+12(FP),R11

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func subborrowU64(b_in uint8, a uint64, b uint64, out uint64) uint8
TEXT ·subborrowU64(SB),7,$0
	MOVB b_in+0(FP),R8
	MOVQ a+4(FP),R9
	MOVQ b+12(FP),R10
	MOVQ out+20(FP),R11

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

