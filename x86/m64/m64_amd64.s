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

// func addSi64(a M64, b M64) M64
TEXT ·addSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func alignrPi8(a M64, b M64, count int) M64
TEXT ·alignrPi8(SB),7,$0
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

// func cvtPi2ps(a [4]float32, b M64) [4]float32
TEXT ·cvtPi2ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),M1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func cvtpdPi32(a [2]float64) M64
TEXT ·cvtpdPi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtpsPi16(a [4]float32) M64
TEXT ·cvtpsPi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtpsPi32(a [4]float32) M64
TEXT ·cvtpsPi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtpsPi8(a [4]float32) M64
TEXT ·cvtpsPi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvttpdPi32(a [2]float64) M64
TEXT ·cvttpdPi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func cvttpsPi32(a [4]float32) M64
TEXT ·cvttpsPi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	// Return size: 8
	RET

// func extractPi16(a M64, imm8 int) int
TEXT ·extractPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
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

// func haddsPi16(a M64, b M64) M64
TEXT ·haddsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
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

// func hsubsPi16(a M64, b M64) M64
TEXT ·hsubsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func insertPi16(a M64, i int, imm8 int) M64
TEXT ·insertPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ i+8(FP),R9
	MOVQ imm8+16(FP),R10

	//TODO: Code missing

	// Return size: 8
	RET

// func loadhPi(a [4]float32, mem_addr M64Const) [4]float32
TEXT ·loadhPi(SB),7,$0
	// Unimplemented. Unknown size of type M64Const
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loadlPi(a [4]float32, mem_addr M64Const) [4]float32
TEXT ·loadlPi(SB),7,$0
	// Unimplemented. Unknown size of type M64Const
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi16(mem_addr uintptr) [16]byte
TEXT ·loaduSi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi161(mem_addr uintptr) [16]byte
TEXT ·loaduSi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi32(mem_addr uintptr) [16]byte
TEXT ·loaduSi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi321(mem_addr uintptr) [16]byte
TEXT ·loaduSi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi64(mem_addr uintptr) [16]byte
TEXT ·loaduSi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func loaduSi641(mem_addr uintptr) [16]byte
TEXT ·loaduSi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maddubsPi16(a M64, b M64) M64
TEXT ·maddubsPi16(SB),7,$0
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

// func maxPi16(a M64, b M64) M64
TEXT ·maxPi16(SB),7,$0
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

// func minPi16(a M64, b M64) M64
TEXT ·minPi16(SB),7,$0
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

// func movemaskPi8(a M64) int
TEXT ·movemaskPi8(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func movepi64Pi64(a [16]byte) M64
TEXT ·movepi64Pi64(SB),7,$0
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

// func mulhrsPi16(a M64, b M64) M64
TEXT ·mulhrsPi16(SB),7,$0
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

// func mPinsrw(a M64, i int, imm8 int) M64
TEXT ·mPinsrw(SB),7,$0
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

// func shufflePi16(a M64, imm8 int) M64
TEXT ·shufflePi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func shufflePi8(a M64, b M64) M64
TEXT ·shufflePi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
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

// func storehPi(mem_addr M64, a [4]float32) 
TEXT ·storehPi(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storelPi(mem_addr M64, a [4]float32) 
TEXT ·storelPi(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storeuSi16(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi161(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi161(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi32(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi321(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi321(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi64(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storeuSi641(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi641(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func streamPi(mem_addr M64, a M64) 
TEXT ·streamPi(SB),7,$0
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

