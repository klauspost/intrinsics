// func mwait(extensions uint, hints uint) 
TEXT ·mwait(SB),7,$0
	MOVQ extensions+0(FP),R8
	MOVQ hints+8(FP),R9

	// TODO: Code missing - could be:
	// MWAIT R8, R9

	RET

