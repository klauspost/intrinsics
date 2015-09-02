// func xgetbv(a uint32) uint64
TEXT ·xgetbv(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// XGETBV R8

	MOVQ $0, ret+4(FP)
	RET

// func xrstor(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstor(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ rs_mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// XRSTOR R8, R9

	RET

// func xrstor64(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstor64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ rs_mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// XRSTOR64 R8, R9

	RET

// func xsave(mem_addr uintptr, save_mask uint64) 
TEXT ·xsave(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// XSAVE R8, R9

	RET

// func xsave64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsave64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// XSAVE64 R8, R9

	RET

// func xsetbv(a uint32, val uint64) 
TEXT ·xsetbv(SB),7,$0
	MOVL a+0(FP),R8
	MOVQ val+4(FP),R9

	// TODO: Code missing
	// Could be:
	// XSETBV R8, R9

	RET

