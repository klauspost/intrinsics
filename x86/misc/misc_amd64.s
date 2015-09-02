// func addcarryU32(c_in uint8, a uint32, b uint32, out uint32) uint8
TEXT ·addcarryU32(SB),7,$0
	MOVB c_in+0(FP),R8
	MOVL a+4(FP),R9
	MOVL b+8(FP),R10
	MOVL out+12(FP),R11

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func addcarryU64(c_in uint8, a uint64, b uint64, out uint64) uint8
TEXT ·addcarryU64(SB),7,$0
	MOVB c_in+0(FP),R8
	MOVQ a+4(FP),R9
	MOVQ b+12(FP),R10
	MOVQ out+20(FP),R11

	//TODO: Code missing

	MOVB $0, ret+28(FP)
	RET

// func allowCpuFeatures(a uint64) 
TEXT ·allowCpuFeatures(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	RET

// func bitScanForward(a int) int
TEXT ·bitScanForward(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func bitScanReverse(a int) int
TEXT ·bitScanReverse(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func bitScanForward2(index uint32, mask uint32) uint8
TEXT ·bitScanForward2(SB),7,$0
	MOVL index+0(FP),R8
	MOVL mask+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+8(FP)
	RET

// func bitScanForward64(index uint32, mask uint64) uint8
TEXT ·bitScanForward64(SB),7,$0
	MOVL index+0(FP),R8
	MOVQ mask+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+12(FP)
	RET

// func bitScanReverse2(index uint32, mask uint32) uint8
TEXT ·bitScanReverse2(SB),7,$0
	MOVL index+0(FP),R8
	MOVL mask+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+8(FP)
	RET

// func bitScanReverse64(index uint32, mask uint64) uint8
TEXT ·bitScanReverse64(SB),7,$0
	MOVL index+0(FP),R8
	MOVQ mask+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+12(FP)
	RET

// func bittest(a int32, b int32) uint8
TEXT ·bittest(SB),7,$0
	MOVL a+0(FP),R8
	MOVL b+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+8(FP)
	RET

// func bittest64(a int64, b int64) uint8
TEXT ·bittest64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func bittestandcomplement(a int32, b int32) uint8
TEXT ·bittestandcomplement(SB),7,$0
	MOVL a+0(FP),R8
	MOVL b+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+8(FP)
	RET

// func bittestandcomplement64(a int64, b int64) uint8
TEXT ·bittestandcomplement64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func bittestandreset(a int32, b int32) uint8
TEXT ·bittestandreset(SB),7,$0
	MOVL a+0(FP),R8
	MOVL b+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+8(FP)
	RET

// func bittestandreset64(a int64, b int64) uint8
TEXT ·bittestandreset64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func bittestandset(a int32, b int32) uint8
TEXT ·bittestandset(SB),7,$0
	MOVL a+0(FP),R8
	MOVL b+4(FP),R9

	//TODO: Code missing

	MOVB $0, ret+8(FP)
	RET

// func bittestandset64(a int64, b int64) uint8
TEXT ·bittestandset64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func bswap(a int) int
TEXT ·bswap(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func bswap64(a int64) int64
TEXT ·bswap64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func castf32U32(a float32) uint32
TEXT ·castf32U32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func castf64U64(a float64) uint64
TEXT ·castf64U64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func castu32F32(a uint32) float32
TEXT ·castu32F32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func castu64F64(a uint64) float64
TEXT ·castu64F64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func cvtshSs(a uint16) float32
TEXT ·cvtshSs(SB),7,$0
	MOVW a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func cvtssSh(a float32, imm8 int) uint16
TEXT ·cvtssSh(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ imm8+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+12(FP)
	RET

// func loadbeI16(ptr uintptr) int16
TEXT ·loadbeI16(SB),7,$0
	MOVQ ptr+0(FP),R8

	//TODO: Code missing

	MOVW $0, ret+8(FP)
	RET

// func loadbeI32(ptr uintptr) int
TEXT ·loadbeI32(SB),7,$0
	MOVQ ptr+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func loadbeI64(ptr uintptr) int64
TEXT ·loadbeI64(SB),7,$0
	MOVQ ptr+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func loaduSi16(mem_addr uintptr) [16]byte
TEXT ·loaduSi16(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loaduSi32(mem_addr uintptr) [16]byte
TEXT ·loaduSi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loaduSi64(mem_addr uintptr) [16]byte
TEXT ·loaduSi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func lrotl(a uint32, shift int) uint32
TEXT ·lrotl(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+12(FP)
	RET

// func lrotr(a uint32, shift int) uint32
TEXT ·lrotr(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+12(FP)
	RET

// func mayIUseCpuFeature(a uint64) int
TEXT ·mayIUseCpuFeature(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func rdpmc(a int) int64
TEXT ·rdpmc(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func rotl(a uint32, shift int) uint32
TEXT ·rotl(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+12(FP)
	RET

// func rotr(a uint32, shift int) uint32
TEXT ·rotr(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ shift+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+12(FP)
	RET

// func rotwl(a uint16, shift int) uint16
TEXT ·rotwl(SB),7,$0
	MOVW a+0(FP),R8
	MOVQ shift+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+12(FP)
	RET

// func rotwr(a uint16, shift int) uint16
TEXT ·rotwr(SB),7,$0
	MOVW a+0(FP),R8
	MOVQ shift+4(FP),R9

	//TODO: Code missing

	MOVW $0, ret+12(FP)
	RET

// func storebeI16(ptr uintptr, data int16) 
TEXT ·storebeI16(SB),7,$0
	MOVQ ptr+0(FP),R8
	MOVW data+8(FP),R9

	//TODO: Code missing

	RET

// func storebeI32(ptr uintptr, data int) 
TEXT ·storebeI32(SB),7,$0
	MOVQ ptr+0(FP),R8
	MOVQ data+8(FP),R9

	//TODO: Code missing

	RET

// func storebeI64(ptr uintptr, data int64) 
TEXT ·storebeI64(SB),7,$0
	MOVQ ptr+0(FP),R8
	MOVQ data+8(FP),R9

	//TODO: Code missing

	RET

// func storeuSi16(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi16(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storeuSi32(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

	RET

// func storeuSi64(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	//TODO: Code missing

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

