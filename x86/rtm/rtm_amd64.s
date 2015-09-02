// func xabort(imm8 uint) 
TEXT ·xabort(SB),7,$0
	MOVQ imm8+0(FP),R8

	//TODO: Code missing

	RET

// func xbegin() uint32
TEXT ·xbegin(SB),7,$0

	//TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func xend() 
TEXT ·xend(SB),7,$0

	//TODO: Code missing

	RET

// func xtest() uint8
TEXT ·xtest(SB),7,$0

	//TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

