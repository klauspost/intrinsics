// func rdtscp(mem_addr uint32) uint64
TEXT ·rdtscp(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// RDTSCP R8

	MOVQ $0, ret+4(FP)
	RET

