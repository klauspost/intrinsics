// func xsaveopt(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaveopt(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	//TODO: Code missing

	RET

// func xsaveopt64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaveopt64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	//TODO: Code missing

	RET

