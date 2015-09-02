// func prefetch(p byte, i int) 
TEXT ·prefetch(SB),7,$0
	MOVB p+0(FP),R8
	MOVQ i+4(FP),R9

	//TODO: Code missing

	RET

