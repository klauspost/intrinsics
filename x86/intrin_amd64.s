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

// func addcarryxU32(c_in uint8, a uint32, b uint32, out uint32) uint8
TEXT ·addcarryxU32(SB),7,$0
	MOVB c_in+0(FP),R8
	MOVL a+4(FP),R9
	MOVL b+8(FP),R10
	MOVL out+12(FP),R11

	//TODO: Code missing

	MOVB $0, ret+16(FP)
	RET

// func addcarryxU64(c_in uint8, a uint64, b uint64, out uint64) uint8
TEXT ·addcarryxU64(SB),7,$0
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

// func bextrU32(a uint32, start uint32, len uint32) uint32
TEXT ·bextrU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL start+4(FP),R9
	MOVL len+8(FP),R10

	//TODO: Code missing

	MOVL $0, ret+12(FP)
	RET

// func bextrU64(a uint64, start uint32, len uint32) uint64
TEXT ·bextrU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVL start+8(FP),R9
	MOVL len+12(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
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

// func bitScanForward1(index uint32, mask uint32) uint8
TEXT ·bitScanForward1(SB),7,$0
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

// func bitScanReverse1(index uint32, mask uint32) uint8
TEXT ·bitScanReverse1(SB),7,$0
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

// func blsiU32(a uint32) uint32
TEXT ·blsiU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func blsiU64(a uint64) uint64
TEXT ·blsiU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func blsmskU32(a uint32) uint32
TEXT ·blsmskU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func blsmskU64(a uint64) uint64
TEXT ·blsmskU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func blsrU32(a uint32) uint32
TEXT ·blsrU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func blsrU64(a uint64) uint64
TEXT ·blsrU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func bndChkPtrBounds(q uintptr, size int) 
TEXT ·bndChkPtrBounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndChkPtrLbounds(q uintptr) 
TEXT ·bndChkPtrLbounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndChkPtrUbounds(q uintptr) 
TEXT ·bndChkPtrUbounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndCopyPtrBounds(q uintptr, r uintptr) 
TEXT ·bndCopyPtrBounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndGetPtrLbound(q uintptr) uintptr
TEXT ·bndGetPtrLbound(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	// Return size: 0
	RET

// func bndGetPtrUbound(q uintptr) uintptr
TEXT ·bndGetPtrUbound(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	// Return size: 0
	RET

// func bndInitPtrBounds(q uintptr) 
TEXT ·bndInitPtrBounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndNarrowPtrBounds(q uintptr, r uintptr, size int) 
TEXT ·bndNarrowPtrBounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndSetPtrBounds(srcmem uintptr, size int) 
TEXT ·bndSetPtrBounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func bndStorePtrBounds(ptr_addr uintptr, ptr_val uintptr) 
TEXT ·bndStorePtrBounds(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

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

// func bzhiU32(a uint32, index uint32) uint32
TEXT ·bzhiU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL index+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func bzhiU64(a uint64, index uint32) uint64
TEXT ·bzhiU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVL index+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+12(FP)
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

// func clevict(ptr uintptr, level int) 
TEXT ·clevict(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func clflush(p uintptr) 
TEXT ·clflush(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func clflushopt(p uintptr) 
TEXT ·clflushopt(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func cmpestra(a M128i, la int, b M128i, lb int, imm8 int) int
TEXT ·cmpestra(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrc(a M128i, la int, b M128i, lb int, imm8 int) int
TEXT ·cmpestrc(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVQ $0, ret+56(FP)
	RET

// func cmpestri(a M128i, la int, b M128i, lb int, imm8 int) int
TEXT ·cmpestri(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrm(a M128i, la int, b M128i, lb int, imm8 int) M128i
TEXT ·cmpestrm(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+56(FP)
	RET

// func cmpestro(a M128i, la int, b M128i, lb int, imm8 int) int
TEXT ·cmpestro(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrs(a M128i, la int, b M128i, lb int, imm8 int) int
TEXT ·cmpestrs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrz(a M128i, la int, b M128i, lb int, imm8 int) int
TEXT ·cmpestrz(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	//TODO: Code missing

	MOVQ $0, ret+56(FP)
	RET

// func cmpistra(a M128i, b M128i, imm8 int) int
TEXT ·cmpistra(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrc(a M128i, b M128i, imm8 int) int
TEXT ·cmpistrc(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+40(FP)
	RET

// func cmpistri(a M128i, b M128i, imm8 int) int
TEXT ·cmpistri(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrm(a M128i, b M128i, imm8 int) M128i
TEXT ·cmpistrm(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func cmpistro(a M128i, b M128i, imm8 int) int
TEXT ·cmpistro(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrs(a M128i, b M128i, imm8 int) int
TEXT ·cmpistrs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrz(a M128i, b M128i, imm8 int) int
TEXT ·cmpistrz(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVQ $0, ret+40(FP)
	RET

// func countbits32(r1 uint32) uint32
TEXT ·countbits32(SB),7,$0
	MOVL r1+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func countbits64(r1 uint64) uint64
TEXT ·countbits64(SB),7,$0
	MOVQ r1+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func crc32U16(crc uint32, v uint16) uint32
TEXT ·crc32U16(SB),7,$0
	MOVL crc+0(FP),R8
	MOVW v+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func crc32U32(crc uint32, v uint32) uint32
TEXT ·crc32U32(SB),7,$0
	MOVL crc+0(FP),R8
	MOVL v+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func crc32U64(crc uint64, v uint64) uint64
TEXT ·crc32U64(SB),7,$0
	MOVQ crc+0(FP),R8
	MOVQ v+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func crc32U8(crc uint32, v uint8) uint32
TEXT ·crc32U8(SB),7,$0
	MOVL crc+0(FP),R8
	MOVB v+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func delay32(r1 uint32) 
TEXT ·delay32(SB),7,$0
	MOVL r1+0(FP),R8

	//TODO: Code missing

	RET

// func delay64(r1 uint64) 
TEXT ·delay64(SB),7,$0
	MOVQ r1+0(FP),R8

	//TODO: Code missing

	RET

// func free(mem_addr uintptr) 
TEXT ·free(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func fxrstor(mem_addr uintptr) 
TEXT ·fxrstor(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func fxrstor64(mem_addr uintptr) 
TEXT ·fxrstor64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func fxsave(mem_addr uintptr) 
TEXT ·fxsave(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func fxsave64(mem_addr uintptr) 
TEXT ·fxsave64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func mMGETEXCEPTIONMASK() uint32
TEXT ·mMGETEXCEPTIONMASK(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func mMGETEXCEPTIONSTATE() uint32
TEXT ·mMGETEXCEPTIONSTATE(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func mMGETFLUSHZEROMODE() uint32
TEXT ·mMGETFLUSHZEROMODE(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func mMGETROUNDINGMODE() uint32
TEXT ·mMGETROUNDINGMODE(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func getcsr() uint32
TEXT ·getcsr(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func invpcid(typ uint32, descriptor uintptr) 
TEXT ·invpcid(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func lfence() 
TEXT ·lfence(SB),7,$0

	//TODO: Code missing

	RET

// func loadbeI16(ptr uintptr) int16
TEXT ·loadbeI16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func loadbeI32(ptr uintptr) int
TEXT ·loadbeI32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func loadbeI64(ptr uintptr) int64
TEXT ·loadbeI64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
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

// func lzcntU32(a uint32) uint32
TEXT ·lzcntU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func lzcntU64(a uint64) uint64
TEXT ·lzcntU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func malloc(size int, align int) 
TEXT ·malloc(SB),7,$0
	MOVQ size+0(FP),R8
	MOVQ align+8(FP),R9

	//TODO: Code missing

	RET

// func mayIUseCpuFeature(a uint64) int
TEXT ·mayIUseCpuFeature(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func mfence() 
TEXT ·mfence(SB),7,$0

	//TODO: Code missing

	RET

// func monitor(p uintptr, extensions uint, hints uint) 
TEXT ·monitor(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func mwait(extensions uint, hints uint) 
TEXT ·mwait(SB),7,$0
	MOVQ extensions+0(FP),R8
	MOVQ hints+8(FP),R9

	//TODO: Code missing

	RET

// func pause() 
TEXT ·pause(SB),7,$0

	//TODO: Code missing

	RET

// func pextU32(a uint32, mask uint32) uint32
TEXT ·pextU32(SB),7,$0
	MOVL a+0(FP),R8
	MOVL mask+4(FP),R9

	//TODO: Code missing

	MOVL $0, ret+8(FP)
	RET

// func pextU64(a uint64, mask uint64) uint64
TEXT ·pextU64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ mask+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func popcntU32(a uint32) int
TEXT ·popcntU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func popcntU64(a uint64) int64
TEXT ·popcntU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func popcnt32(a int) int
TEXT ·popcnt32(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func popcnt64(a int64) int
TEXT ·popcnt64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func prefetch(p byte, i int) 
TEXT ·prefetch(SB),7,$0
	MOVB p+0(FP),R8
	MOVQ i+4(FP),R9

	//TODO: Code missing

	RET

// func prefetch1(p byte, i int) 
TEXT ·prefetch1(SB),7,$0
	MOVB p+0(FP),R8
	MOVQ i+4(FP),R9

	//TODO: Code missing

	RET

// func prefetch2(p byte, i int) 
TEXT ·prefetch2(SB),7,$0
	MOVB p+0(FP),R8
	MOVQ i+4(FP),R9

	//TODO: Code missing

	RET

// func rdpmc(a int) int64
TEXT ·rdpmc(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func rdrand16Step(val uint16) int
TEXT ·rdrand16Step(SB),7,$0
	MOVW val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func rdrand32Step(val uint32) int
TEXT ·rdrand32Step(SB),7,$0
	MOVL val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func rdrand64Step(val uint64) int
TEXT ·rdrand64Step(SB),7,$0
	MOVQ val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func rdseed16Step(val uint16) int
TEXT ·rdseed16Step(SB),7,$0
	MOVW val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func rdseed32Step(val uint32) int
TEXT ·rdseed32Step(SB),7,$0
	MOVL val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func rdseed64Step(val uint64) int
TEXT ·rdseed64Step(SB),7,$0
	MOVQ val+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func rdtsc() int64
TEXT ·rdtsc(SB),7,$0

	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func rdtscp(mem_addr uint32) uint64
TEXT ·rdtscp(SB),7,$0
	MOVL mem_addr+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func readfsbaseU32() uint32
TEXT ·readfsbaseU32(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func readfsbaseU64() uint64
TEXT ·readfsbaseU64(SB),7,$0

	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func readgsbaseU32() uint32
TEXT ·readgsbaseU32(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func readgsbaseU64() uint64
TEXT ·readgsbaseU64(SB),7,$0

	//TODO: Code missing

	MOVQ $0, ret+0(FP)
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

// func mMSETEXCEPTIONMASK(a uint32) 
TEXT ·mMSETEXCEPTIONMASK(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func mMSETEXCEPTIONSTATE(a uint32) 
TEXT ·mMSETEXCEPTIONSTATE(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func mMSETFLUSHZEROMODE(a uint32) 
TEXT ·mMSETFLUSHZEROMODE(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func mMSETROUNDINGMODE(a uint32) 
TEXT ·mMSETROUNDINGMODE(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func setcsr(a uint32) 
TEXT ·setcsr(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func sfence() 
TEXT ·sfence(SB),7,$0

	//TODO: Code missing

	RET

// func spflt32(r1 uint32) 
TEXT ·spflt32(SB),7,$0
	MOVL r1+0(FP),R8

	//TODO: Code missing

	RET

// func spflt64(r1 uint64) 
TEXT ·spflt64(SB),7,$0
	MOVQ r1+0(FP),R8

	//TODO: Code missing

	RET

// func storebeI16(ptr uintptr, data int16) 
TEXT ·storebeI16(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storebeI32(ptr uintptr, data int) 
TEXT ·storebeI32(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func storebeI64(ptr uintptr, data int64) 
TEXT ·storebeI64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func testAllOnes(a M128i) int
TEXT ·testAllOnes(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func testAllZeros(a M128i, mask M128i) int
TEXT ·testAllZeros(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func testMixOnesZeros(a M128i, mask M128i) int
TEXT ·testMixOnesZeros(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1

	//TODO: Code missing

	MOVQ $0, ret+32(FP)
	RET

// func mMTRANSPOSE4PS(row0 [4]float32, row1 [4]float32, row2 [4]float32, row3 [4]float32) 
TEXT ·mMTRANSPOSE4PS(SB),7,$0
	MOVOU row0+0(FP),X0
	MOVOU row1+16(FP),X1
	MOVOU row2+32(FP),X2
	MOVOU row3+48(FP),X3

	//TODO: Code missing

	RET

// func tzcnt32(a uint32) int
TEXT ·tzcnt32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func tzcnt64(a uint64) int64
TEXT ·tzcnt64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func tzcntU32(a uint32) uint32
TEXT ·tzcntU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVL $0, ret+4(FP)
	RET

// func tzcntU64(a uint64) uint64
TEXT ·tzcntU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func tzcnti32(a int, x uint32) int
TEXT ·tzcnti32(SB),7,$0
	MOVQ a+0(FP),R8
	MOVL x+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+12(FP)
	RET

// func tzcnti64(a int64, x uint64) int64
TEXT ·tzcnti64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ x+8(FP),R9

	//TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func writefsbaseU32(a uint32) 
TEXT ·writefsbaseU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func writefsbaseU64(a uint64) 
TEXT ·writefsbaseU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	RET

// func writegsbaseU32(a uint32) 
TEXT ·writegsbaseU32(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	RET

// func writegsbaseU64(a uint64) 
TEXT ·writegsbaseU64(SB),7,$0
	MOVQ a+0(FP),R8

	//TODO: Code missing

	RET

// func xabort(imm8 uint) 
TEXT ·xabort(SB),7,$0
	MOVQ imm8+0(FP),R8

	//TODO: Code missing

	RET

// func xbegin() uint32
TEXT ·xbegin(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func xend() 
TEXT ·xend(SB),7,$0

	//TODO: Code missing

	RET

// func xgetbv(a uint32) uint64
TEXT ·xgetbv(SB),7,$0
	MOVL a+0(FP),R8

	//TODO: Code missing

	MOVQ $0, ret+4(FP)
	RET

// func xrstor(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstor(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xrstor64(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstor64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xrstors(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstors(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xrstors64(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstors64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsave(mem_addr uintptr, save_mask uint64) 
TEXT ·xsave(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsave64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsave64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsavec(mem_addr uintptr, save_mask uint64) 
TEXT ·xsavec(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsavec64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsavec64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsaveopt(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaveopt(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsaveopt64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaveopt64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsaves(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaves(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsaves64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaves64(SB),7,$0
	// Unimplemented. Unknown size of type uintptr
	//TODO: Code missing

	RET

// func xsetbv(a uint32, val uint64) 
TEXT ·xsetbv(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ val+4(FP),R9

	//TODO: Code missing

	RET

// func xtest() uint8
TEXT ·xtest(SB),7,$0

	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

