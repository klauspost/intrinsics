// func monitor(p uintptr, extensions uint, hints uint) 
TEXT ·monitor(SB),7,$0
	MOVQ p+0(FP),R8
	MOVQ extensions+8(FP),R9
	MOVQ hints+16(FP),R10

	// TODO: Code missing

	RET

// func mwait(extensions uint, hints uint) 
TEXT ·mwait(SB),7,$0
	MOVQ extensions+0(FP),R8
	MOVQ hints+8(FP),R9

	// TODO: Code missing - could be:
	// MWAIT R8, R9

	RET

