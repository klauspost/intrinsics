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

// func addPi16(a M64, b M64) M64
TEXT ·addPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func addPi32(a M64, b M64) M64
TEXT ·addPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func addPi8(a M64, b M64) M64
TEXT ·addPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

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

// func addsPi16(a M64, b M64) M64
TEXT ·addsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func addsPi8(a M64, b M64) M64
TEXT ·addsPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func addsPu16(a M64, b M64) M64
TEXT ·addsPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func addsPu8(a M64, b M64) M64
TEXT ·addsPu8(SB),7,$0
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

// func andSi64(a M64, b M64) M64
TEXT ·andSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func andnotSi64(a M64, b M64) M64
TEXT ·andnotSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

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

// func cmpeqPi16(a M64, b M64) M64
TEXT ·cmpeqPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func cmpeqPi32(a M64, b M64) M64
TEXT ·cmpeqPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func cmpeqPi8(a M64, b M64) M64
TEXT ·cmpeqPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func cmpgtPi16(a M64, b M64) M64
TEXT ·cmpgtPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func cmpgtPi32(a M64, b M64) M64
TEXT ·cmpgtPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func cmpgtPi8(a M64, b M64) M64
TEXT ·cmpgtPi8(SB),7,$0
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

// func cvtm64Si64(a M64) int64
TEXT ·cvtm64Si64(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
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

// func cvtsi32Si64(a int) M64
TEXT ·cvtsi32Si64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtsi64M64(a int64) M64
TEXT ·cvtsi64M64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
	RET

// func cvtsi64Si32(a M64) int
TEXT ·cvtsi64Si32(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
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

// func empty() 
TEXT ·empty(SB),7,$0

	//TODO: Code missing

	RET

// func empty1() 
TEXT ·empty1(SB),7,$0

	//TODO: Code missing

	RET

// func extractPi16(a M64, imm8 int) int
TEXT ·extractPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func fromInt(a int) M64
TEXT ·fromInt(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
	RET

// func fromInt64(a int64) M64
TEXT ·fromInt64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
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

// func maddPi16(a M64, b M64) M64
TEXT ·maddPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
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

// func maskmovq(a M64, mask M64, mem_addr byte) 
TEXT ·maskmovq(SB),7,$0
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

// func mulhiPi16(a M64, b M64) M64
TEXT ·mulhiPi16(SB),7,$0
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

// func mulloPi16(a M64, b M64) M64
TEXT ·mulloPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func orSi64(a M64, b M64) M64
TEXT ·orSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func packsPi16(a M64, b M64) M64
TEXT ·packsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func packsPi32(a M64, b M64) M64
TEXT ·packsPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func packsPu16(a M64, b M64) M64
TEXT ·packsPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func packssdw(a M64, b M64) M64
TEXT ·packssdw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func packsswb(a M64, b M64) M64
TEXT ·packsswb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func packuswb(a M64, b M64) M64
TEXT ·packuswb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddb(a M64, b M64) M64
TEXT ·paddb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddd(a M64, b M64) M64
TEXT ·paddd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddsb(a M64, b M64) M64
TEXT ·paddsb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddsw(a M64, b M64) M64
TEXT ·paddsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddusb(a M64, b M64) M64
TEXT ·paddusb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddusw(a M64, b M64) M64
TEXT ·paddusw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func paddw(a M64, b M64) M64
TEXT ·paddw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pand(a M64, b M64) M64
TEXT ·pand(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pandn(a M64, b M64) M64
TEXT ·pandn(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pavgb(a M64, b M64) M64
TEXT ·pavgb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pavgw(a M64, b M64) M64
TEXT ·pavgw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pcmpeqb(a M64, b M64) M64
TEXT ·pcmpeqb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pcmpeqd(a M64, b M64) M64
TEXT ·pcmpeqd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pcmpeqw(a M64, b M64) M64
TEXT ·pcmpeqw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pcmpgtb(a M64, b M64) M64
TEXT ·pcmpgtb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pcmpgtd(a M64, b M64) M64
TEXT ·pcmpgtd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pcmpgtw(a M64, b M64) M64
TEXT ·pcmpgtw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pextrw(a M64, imm8 int) int
TEXT ·pextrw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func pinsrw(a M64, i int, imm8 int) M64
TEXT ·pinsrw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ i+8(FP),R9
	MOVQ imm8+16(FP),R10

	//TODO: Code missing

	// Return size: 8
	RET

// func pmaddwd(a M64, b M64) M64
TEXT ·pmaddwd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pmaxsw(a M64, b M64) M64
TEXT ·pmaxsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pmaxub(a M64, b M64) M64
TEXT ·pmaxub(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pminsw(a M64, b M64) M64
TEXT ·pminsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pminub(a M64, b M64) M64
TEXT ·pminub(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pmovmskb(a M64) int
TEXT ·pmovmskb(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func pmulhuw(a M64, b M64) M64
TEXT ·pmulhuw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pmulhw(a M64, b M64) M64
TEXT ·pmulhw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pmullw(a M64, b M64) M64
TEXT ·pmullw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func por(a M64, b M64) M64
TEXT ·por(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psadbw(a M64, b M64) M64
TEXT ·psadbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pshufw(a M64, imm8 int) M64
TEXT ·pshufw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func pslld(a M64, count M64) M64
TEXT ·pslld(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pslldi(a M64, imm8 int) M64
TEXT ·pslldi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psllq(a M64, count M64) M64
TEXT ·psllq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psllqi(a M64, imm8 int) M64
TEXT ·psllqi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psllw(a M64, count M64) M64
TEXT ·psllw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psllwi(a M64, imm8 int) M64
TEXT ·psllwi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psrad(a M64, count M64) M64
TEXT ·psrad(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psradi(a M64, imm8 int) M64
TEXT ·psradi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psraw(a M64, count M64) M64
TEXT ·psraw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psrawi(a M64, imm8 int) M64
TEXT ·psrawi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psrld(a M64, count M64) M64
TEXT ·psrld(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psrldi(a M64, imm8 int) M64
TEXT ·psrldi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psrlq(a M64, count M64) M64
TEXT ·psrlq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psrlqi(a M64, imm8 int) M64
TEXT ·psrlqi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psrlw(a M64, count M64) M64
TEXT ·psrlw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psrlwi(a M64, imm8 int) M64
TEXT ·psrlwi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func psubb(a M64, b M64) M64
TEXT ·psubb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psubd(a M64, b M64) M64
TEXT ·psubd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psubsb(a M64, b M64) M64
TEXT ·psubsb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psubsw(a M64, b M64) M64
TEXT ·psubsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psubusb(a M64, b M64) M64
TEXT ·psubusb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psubusw(a M64, b M64) M64
TEXT ·psubusw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func psubw(a M64, b M64) M64
TEXT ·psubw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func punpckhbw(a M64, b M64) M64
TEXT ·punpckhbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func punpckhdq(a M64, b M64) M64
TEXT ·punpckhdq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func punpckhwd(a M64, b M64) M64
TEXT ·punpckhwd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func punpcklbw(a M64, b M64) M64
TEXT ·punpcklbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func punpckldq(a M64, b M64) M64
TEXT ·punpckldq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func punpcklwd(a M64, b M64) M64
TEXT ·punpcklwd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func pxor(a M64, b M64) M64
TEXT ·pxor(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

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

// func setPi16(e3 int16, e2 int16, e1 int16, e0 int16) M64
TEXT ·setPi16(SB),7,$0
	MOVW e3+0(FP),R8
	MOVW e2+4(FP),R9
	MOVW e1+8(FP),R10
	MOVW e0+12(FP),R11

	//TODO: Code missing

	// Return size: 8
	RET

// func setPi32(e1 int, e0 int) M64
TEXT ·setPi32(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func setPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M64
TEXT ·setPi8(SB),7,$0
	MOVB e7+0(FP),R8
	MOVB e6+4(FP),R9
	MOVB e5+8(FP),R10
	MOVB e4+12(FP),R11
	MOVB e3+16(FP),R12
	MOVB e2+20(FP),R13
	MOVB e1+24(FP),R14
	MOVB e0+28(FP),R15

	//TODO: Code missing

	// Return size: 8
	RET

// func set1Pi16(a int16) M64
TEXT ·set1Pi16(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
	RET

// func set1Pi32(a int) M64
TEXT ·set1Pi32(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
	RET

// func set1Pi8(a byte) M64
TEXT ·set1Pi8(SB),7,$0
	MOVB a+0(FP),R8

	//TODO: Code missing

	// Return size: 8
	RET

// func setrPi16(e3 int16, e2 int16, e1 int16, e0 int16) M64
TEXT ·setrPi16(SB),7,$0
	MOVW e3+0(FP),R8
	MOVW e2+4(FP),R9
	MOVW e1+8(FP),R10
	MOVW e0+12(FP),R11

	//TODO: Code missing

	// Return size: 8
	RET

// func setrPi32(e1 int, e0 int) M64
TEXT ·setrPi32(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func setrPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M64
TEXT ·setrPi8(SB),7,$0
	MOVB e7+0(FP),R8
	MOVB e6+4(FP),R9
	MOVB e5+8(FP),R10
	MOVB e4+12(FP),R11
	MOVB e3+16(FP),R12
	MOVB e2+20(FP),R13
	MOVB e1+24(FP),R14
	MOVB e0+28(FP),R15

	//TODO: Code missing

	// Return size: 8
	RET

// func setzeroSi64() M64
TEXT ·setzeroSi64(SB),7,$0

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

// func sllPi16(a M64, count M64) M64
TEXT ·sllPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sllPi32(a M64, count M64) M64
TEXT ·sllPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sllSi64(a M64, count M64) M64
TEXT ·sllSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func slliPi16(a M64, imm8 int) M64
TEXT ·slliPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func slliPi32(a M64, imm8 int) M64
TEXT ·slliPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func slliSi64(a M64, imm8 int) M64
TEXT ·slliSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func sraPi16(a M64, count M64) M64
TEXT ·sraPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sraPi32(a M64, count M64) M64
TEXT ·sraPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func sraiPi16(a M64, imm8 int) M64
TEXT ·sraiPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func sraiPi32(a M64, imm8 int) M64
TEXT ·sraiPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func srlPi16(a M64, count M64) M64
TEXT ·srlPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func srlPi32(a M64, count M64) M64
TEXT ·srlPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func srlSi64(a M64, count M64) M64
TEXT ·srlSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func srliPi16(a M64, imm8 int) M64
TEXT ·srliPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func srliPi32(a M64, imm8 int) M64
TEXT ·srliPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	//TODO: Code missing

	// Return size: 8
	RET

// func srliSi64(a M64, imm8 int) M64
TEXT ·srliSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

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

// func subPi16(a M64, b M64) M64
TEXT ·subPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subPi32(a M64, b M64) M64
TEXT ·subPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subPi8(a M64, b M64) M64
TEXT ·subPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
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

// func subsPi16(a M64, b M64) M64
TEXT ·subsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subsPi8(a M64, b M64) M64
TEXT ·subsPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subsPu16(a M64, b M64) M64
TEXT ·subsPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func subsPu8(a M64, b M64) M64
TEXT ·subsPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func toInt(a M64) int
TEXT ·toInt(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func toInt64(a M64) int64
TEXT ·toInt64(SB),7,$0
	MOVQ a+0(FP),M0

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func unpackhiPi16(a M64, b M64) M64
TEXT ·unpackhiPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func unpackhiPi32(a M64, b M64) M64
TEXT ·unpackhiPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func unpackhiPi8(a M64, b M64) M64
TEXT ·unpackhiPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func unpackloPi16(a M64, b M64) M64
TEXT ·unpackloPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func unpackloPi32(a M64, b M64) M64
TEXT ·unpackloPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func unpackloPi8(a M64, b M64) M64
TEXT ·unpackloPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

// func xorSi64(a M64, b M64) M64
TEXT ·xorSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	//TODO: Code missing

	// Return size: 8
	RET

