// func fxrstor(mem_addr uintptr) 
TEXT ·fxrstor(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// FXRSTOR R8

	RET

// func fxrstor64(mem_addr uintptr) 
TEXT ·fxrstor64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// FXRSTOR64 R8

	RET

// func fxsave(mem_addr uintptr) 
TEXT ·fxsave(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// FXSAVE R8

	RET

// func fxsave64(mem_addr uintptr) 
TEXT ·fxsave64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// FXSAVE64 R8

	RET

