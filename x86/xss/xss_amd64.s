// func xrstors(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstors(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ rs_mask+8(FP),R9

	//TODO: Code missing

	RET

// func xrstors64(mem_addr uintptr, rs_mask uint64) 
TEXT ·xrstors64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ rs_mask+8(FP),R9

	//TODO: Code missing

	RET

// func xsaves(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaves(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	//TODO: Code missing

	RET

// func xsaves64(mem_addr uintptr, save_mask uint64) 
TEXT ·xsaves64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVQ save_mask+8(FP),R9

	//TODO: Code missing

	RET

