// func xsavec(mem_addr uintptr, save_mask uint64) 
TEXT ·xsavec(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// XSAVEC R8, R9

	RET

// func xsavec64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsavec64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	// TODO: Code missing
	// Could be:
	// XSAVEC64 R8, R9

	RET

