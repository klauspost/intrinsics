// func fxrstor(mem_addr uintptr) 
TEXT ·fxrstor(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	RET

// func fxrstor64(mem_addr uintptr) 
TEXT ·fxrstor64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	RET

// func fxsave(mem_addr uintptr) 
TEXT ·fxsave(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	RET

// func fxsave64(mem_addr uintptr) 
TEXT ·fxsave64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	//TODO: Code missing

	RET

