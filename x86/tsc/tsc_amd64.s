// func rdtsc() int64
TEXT ·rdtsc(SB),7,$0

	// TODO: Code missing - uses instrunction: RDTSC

	MOVQ $0, ret+0(FP)
	RET

