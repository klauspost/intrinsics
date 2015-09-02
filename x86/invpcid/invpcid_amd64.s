// func invpcid(typ uint32, descriptor uintptr) 
TEXT ·invpcid(SB),7,$0
	MOVL typ+0(FP),R8
	MOVQ descriptor+4(FP),R9

	//TODO: Code missing

	RET

