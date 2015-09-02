// func bndChkPtrBounds(q uintptr, size int) 
TEXT ·bndChkPtrBounds(SB),7,$0
	MOVQ q+0(FP),R8
	MOVQ size+8(FP),R9

	// TODO: Code missing
	// Could be:
	// BNDCU, BNDCN R8, R9

	RET

// func bndChkPtrLbounds(q uintptr) 
TEXT ·bndChkPtrLbounds(SB),7,$0
	MOVQ q+0(FP),R8

	// TODO: Code missing
	// Could be:
	// BNDCL R8

	RET

// func bndChkPtrUbounds(q uintptr) 
TEXT ·bndChkPtrUbounds(SB),7,$0
	MOVQ q+0(FP),R8

	// TODO: Code missing
	// Could be:
	// BNDCU, BNDCN R8

	RET

// func bndCopyPtrBounds(q uintptr, r uintptr) 
TEXT ·bndCopyPtrBounds(SB),7,$0
	MOVQ q+0(FP),R8
	MOVQ r+8(FP),R9

	// TODO: Code missing

	RET

// func bndGetPtrLbound(q uintptr) uintptr
TEXT ·bndGetPtrLbound(SB),7,$0
	MOVQ q+0(FP),R8

	// TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func bndGetPtrUbound(q uintptr) uintptr
TEXT ·bndGetPtrUbound(SB),7,$0
	MOVQ q+0(FP),R8

	// TODO: Code missing

	MOVQ $0, ret+8(FP)
	RET

// func bndInitPtrBounds(q uintptr) 
TEXT ·bndInitPtrBounds(SB),7,$0
	MOVQ q+0(FP),R8

	// TODO: Code missing

	RET

// func bndNarrowPtrBounds(q uintptr, r uintptr, size int) 
TEXT ·bndNarrowPtrBounds(SB),7,$0
	MOVQ q+0(FP),R8
	MOVQ r+8(FP),R9
	MOVQ size+16(FP),R10

	// TODO: Code missing

	RET

// func bndSetPtrBounds(srcmem uintptr, size int) 
TEXT ·bndSetPtrBounds(SB),7,$0
	MOVQ srcmem+0(FP),R8
	MOVQ size+8(FP),R9

	// TODO: Code missing
	// Could be:
	// BNDMK R8, R9

	RET

// func bndStorePtrBounds(ptr_addr uintptr, ptr_val uintptr) 
TEXT ·bndStorePtrBounds(SB),7,$0
	MOVQ ptr_addr+0(FP),R8
	MOVQ ptr_val+8(FP),R9

	// TODO: Code missing
	// Could be:
	// BNDSTX R8, R9

	RET

