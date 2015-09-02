// func addPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·addPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDW M0, M1

	// Return size: 8
	RET

// func addPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·addPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDD M0, M1

	// Return size: 8
	RET

// func addPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·addPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDB M0, M1

	// Return size: 8
	RET

// func addsPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·addsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDSW M0, M1

	// Return size: 8
	RET

// func addsPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·addsPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDSB M0, M1

	// Return size: 8
	RET

// func addsPu16(a x86.M64, b x86.M64) x86.M64
TEXT ·addsPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDUSW M0, M1

	// Return size: 8
	RET

// func addsPu8(a x86.M64, b x86.M64) x86.M64
TEXT ·addsPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDUSB M0, M1

	// Return size: 8
	RET

// func andSi64(a x86.M64, b x86.M64) x86.M64
TEXT ·andSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PAND M0, M1

	// Return size: 8
	RET

// func andnotSi64(a x86.M64, b x86.M64) x86.M64
TEXT ·andnotSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PANDN M0, M1

	// Return size: 8
	RET

// func cmpeqPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·cmpeqPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPEQW M0, M1

	// Return size: 8
	RET

// func cmpeqPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·cmpeqPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPEQD M0, M1

	// Return size: 8
	RET

// func cmpeqPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·cmpeqPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPEQB M0, M1

	// Return size: 8
	RET

// func cmpgtPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·cmpgtPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPGTW M0, M1

	// Return size: 8
	RET

// func cmpgtPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·cmpgtPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPGTD M0, M1

	// Return size: 8
	RET

// func cmpgtPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·cmpgtPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPGTB M0, M1

	// Return size: 8
	RET

// func cvtm64Si64(a x86.M64) int64
TEXT ·cvtm64Si64(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing - could be:
	// MOVQ M0

	MOVQ $0, ret+8(FP)
	RET

// func cvtsi32Si64(a int) x86.M64
TEXT ·cvtsi32Si64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// MOVD R8

	// Return size: 8
	RET

// func cvtsi64M64(a int64) x86.M64
TEXT ·cvtsi64M64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// MOVQ R8

	// Return size: 8
	RET

// func cvtsi64Si32(a x86.M64) int
TEXT ·cvtsi64Si32(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing - could be:
	// MOVD M0

	MOVQ $0, ret+8(FP)
	RET

// func empty() 
TEXT ·empty(SB),7,$0

	// TODO: Code missing

	RET

// func empty2() 
TEXT ·empty2(SB),7,$0

	// TODO: Code missing

	RET

// func fromInt(a int) x86.M64
TEXT ·fromInt(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// MOVD R8

	// Return size: 8
	RET

// func fromInt64(a int64) x86.M64
TEXT ·fromInt64(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing - could be:
	// MOVQ R8

	// Return size: 8
	RET

// func maddPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·maddPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMADDWD M0, M1

	// Return size: 8
	RET

// func mulhiPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·mulhiPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMULHW M0, M1

	// Return size: 8
	RET

// func mulloPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·mulloPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMULLW M0, M1

	// Return size: 8
	RET

// func orSi64(a x86.M64, b x86.M64) x86.M64
TEXT ·orSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// POR M0, M1

	// Return size: 8
	RET

// func packsPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·packsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PACKSSWB M0, M1

	// Return size: 8
	RET

// func packsPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·packsPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PACKSSDW M0, M1

	// Return size: 8
	RET

// func packsPu16(a x86.M64, b x86.M64) x86.M64
TEXT ·packsPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PACKUSWB M0, M1

	// Return size: 8
	RET

// func packssdw(a x86.M64, b x86.M64) x86.M64
TEXT ·packssdw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PACKSSDW M0, M1

	// Return size: 8
	RET

// func packsswb(a x86.M64, b x86.M64) x86.M64
TEXT ·packsswb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PACKSSWB M0, M1

	// Return size: 8
	RET

// func packuswb(a x86.M64, b x86.M64) x86.M64
TEXT ·packuswb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PACKUSWB M0, M1

	// Return size: 8
	RET

// func paddb(a x86.M64, b x86.M64) x86.M64
TEXT ·paddb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDB M0, M1

	// Return size: 8
	RET

// func paddd(a x86.M64, b x86.M64) x86.M64
TEXT ·paddd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDD M0, M1

	// Return size: 8
	RET

// func paddsb(a x86.M64, b x86.M64) x86.M64
TEXT ·paddsb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDSB M0, M1

	// Return size: 8
	RET

// func paddsw(a x86.M64, b x86.M64) x86.M64
TEXT ·paddsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDSW M0, M1

	// Return size: 8
	RET

// func paddusb(a x86.M64, b x86.M64) x86.M64
TEXT ·paddusb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDUSB M0, M1

	// Return size: 8
	RET

// func paddusw(a x86.M64, b x86.M64) x86.M64
TEXT ·paddusw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDUSW M0, M1

	// Return size: 8
	RET

// func paddw(a x86.M64, b x86.M64) x86.M64
TEXT ·paddw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PADDW M0, M1

	// Return size: 8
	RET

// func pand(a x86.M64, b x86.M64) x86.M64
TEXT ·pand(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PAND M0, M1

	// Return size: 8
	RET

// func pandn(a x86.M64, b x86.M64) x86.M64
TEXT ·pandn(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PANDN M0, M1

	// Return size: 8
	RET

// func pcmpeqb(a x86.M64, b x86.M64) x86.M64
TEXT ·pcmpeqb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPEQB M0, M1

	// Return size: 8
	RET

// func pcmpeqd(a x86.M64, b x86.M64) x86.M64
TEXT ·pcmpeqd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPEQD M0, M1

	// Return size: 8
	RET

// func pcmpeqw(a x86.M64, b x86.M64) x86.M64
TEXT ·pcmpeqw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPEQW M0, M1

	// Return size: 8
	RET

// func pcmpgtb(a x86.M64, b x86.M64) x86.M64
TEXT ·pcmpgtb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPGTB M0, M1

	// Return size: 8
	RET

// func pcmpgtd(a x86.M64, b x86.M64) x86.M64
TEXT ·pcmpgtd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPGTD M0, M1

	// Return size: 8
	RET

// func pcmpgtw(a x86.M64, b x86.M64) x86.M64
TEXT ·pcmpgtw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PCMPGTW M0, M1

	// Return size: 8
	RET

// func pmaddwd(a x86.M64, b x86.M64) x86.M64
TEXT ·pmaddwd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMADDWD M0, M1

	// Return size: 8
	RET

// func pmulhw(a x86.M64, b x86.M64) x86.M64
TEXT ·pmulhw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMULHW M0, M1

	// Return size: 8
	RET

// func pmullw(a x86.M64, b x86.M64) x86.M64
TEXT ·pmullw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMULLW M0, M1

	// Return size: 8
	RET

// func por(a x86.M64, b x86.M64) x86.M64
TEXT ·por(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// POR M0, M1

	// Return size: 8
	RET

// func pslld(a x86.M64, count x86.M64) x86.M64
TEXT ·pslld(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSLLD M0, M1

	// Return size: 8
	RET

// func pslldi(a x86.M64, imm8 int) x86.M64
TEXT ·pslldi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSLLD M0, R9

	// Return size: 8
	RET

// func psllq(a x86.M64, count x86.M64) x86.M64
TEXT ·psllq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSLLQ M0, M1

	// Return size: 8
	RET

// func psllqi(a x86.M64, imm8 int) x86.M64
TEXT ·psllqi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSLLQ M0, R9

	// Return size: 8
	RET

// func psllw(a x86.M64, count x86.M64) x86.M64
TEXT ·psllw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSLLW M0, M1

	// Return size: 8
	RET

// func psllwi(a x86.M64, imm8 int) x86.M64
TEXT ·psllwi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSLLW M0, R9

	// Return size: 8
	RET

// func psrad(a x86.M64, count x86.M64) x86.M64
TEXT ·psrad(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRAD M0, M1

	// Return size: 8
	RET

// func psradi(a x86.M64, imm8 int) x86.M64
TEXT ·psradi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRAD M0, R9

	// Return size: 8
	RET

// func psraw(a x86.M64, count x86.M64) x86.M64
TEXT ·psraw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRAW M0, M1

	// Return size: 8
	RET

// func psrawi(a x86.M64, imm8 int) x86.M64
TEXT ·psrawi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRAW M0, R9

	// Return size: 8
	RET

// func psrld(a x86.M64, count x86.M64) x86.M64
TEXT ·psrld(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRLD M0, M1

	// Return size: 8
	RET

// func psrldi(a x86.M64, imm8 int) x86.M64
TEXT ·psrldi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRLD M0, R9

	// Return size: 8
	RET

// func psrlq(a x86.M64, count x86.M64) x86.M64
TEXT ·psrlq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRLQ M0, M1

	// Return size: 8
	RET

// func psrlqi(a x86.M64, imm8 int) x86.M64
TEXT ·psrlqi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRLQ M0, R9

	// Return size: 8
	RET

// func psrlw(a x86.M64, count x86.M64) x86.M64
TEXT ·psrlw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRLW M0, M1

	// Return size: 8
	RET

// func psrlwi(a x86.M64, imm8 int) x86.M64
TEXT ·psrlwi(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRLW M0, R9

	// Return size: 8
	RET

// func psubb(a x86.M64, b x86.M64) x86.M64
TEXT ·psubb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBB M0, M1

	// Return size: 8
	RET

// func psubd(a x86.M64, b x86.M64) x86.M64
TEXT ·psubd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBD M0, M1

	// Return size: 8
	RET

// func psubsb(a x86.M64, b x86.M64) x86.M64
TEXT ·psubsb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBSB M0, M1

	// Return size: 8
	RET

// func psubsw(a x86.M64, b x86.M64) x86.M64
TEXT ·psubsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBSW M0, M1

	// Return size: 8
	RET

// func psubusb(a x86.M64, b x86.M64) x86.M64
TEXT ·psubusb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBUSB M0, M1

	// Return size: 8
	RET

// func psubusw(a x86.M64, b x86.M64) x86.M64
TEXT ·psubusw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBUSW M0, M1

	// Return size: 8
	RET

// func psubw(a x86.M64, b x86.M64) x86.M64
TEXT ·psubw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBW M0, M1

	// Return size: 8
	RET

// func punpckhbw(a x86.M64, b x86.M64) x86.M64
TEXT ·punpckhbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKHBW M0, M1

	// Return size: 8
	RET

// func punpckhdq(a x86.M64, b x86.M64) x86.M64
TEXT ·punpckhdq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKHDQ M0, M1

	// Return size: 8
	RET

// func punpckhwd(a x86.M64, b x86.M64) x86.M64
TEXT ·punpckhwd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLBW M0, M1

	// Return size: 8
	RET

// func punpcklbw(a x86.M64, b x86.M64) x86.M64
TEXT ·punpcklbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLBW M0, M1

	// Return size: 8
	RET

// func punpckldq(a x86.M64, b x86.M64) x86.M64
TEXT ·punpckldq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLDQ M0, M1

	// Return size: 8
	RET

// func punpcklwd(a x86.M64, b x86.M64) x86.M64
TEXT ·punpcklwd(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLWD M0, M1

	// Return size: 8
	RET

// func pxor(a x86.M64, b x86.M64) x86.M64
TEXT ·pxor(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PXOR M0, M1

	// Return size: 8
	RET

// func setPi16(e3 int16, e2 int16, e1 int16, e0 int16) x86.M64
TEXT ·setPi16(SB),7,$0
	MOVW e3+0(FP),R8
	MOVW e2+4(FP),R9
	MOVW e1+8(FP),R10
	MOVW e0+12(FP),R11

	// TODO: Code missing

	// Return size: 8
	RET

// func setPi32(e1 int, e0 int) x86.M64
TEXT ·setPi32(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	// TODO: Code missing

	// Return size: 8
	RET

// func setPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) x86.M64
TEXT ·setPi8(SB),7,$0
	MOVB e7+0(FP),R8
	MOVB e6+4(FP),R9
	MOVB e5+8(FP),R10
	MOVB e4+12(FP),R11
	MOVB e3+16(FP),R12
	MOVB e2+20(FP),R13
	MOVB e1+24(FP),R14
	MOVB e0+28(FP),R15

	// TODO: Code missing

	// Return size: 8
	RET

// func set1Pi16(a int16) x86.M64
TEXT ·set1Pi16(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing

	// Return size: 8
	RET

// func set1Pi32(a int) x86.M64
TEXT ·set1Pi32(SB),7,$0
	MOVQ a+0(FP),R8

	// TODO: Code missing

	// Return size: 8
	RET

// func set1Pi8(a byte) x86.M64
TEXT ·set1Pi8(SB),7,$0
	MOVB a+0(FP),R8

	// TODO: Code missing

	// Return size: 8
	RET

// func setrPi16(e3 int16, e2 int16, e1 int16, e0 int16) x86.M64
TEXT ·setrPi16(SB),7,$0
	MOVW e3+0(FP),R8
	MOVW e2+4(FP),R9
	MOVW e1+8(FP),R10
	MOVW e0+12(FP),R11

	// TODO: Code missing

	// Return size: 8
	RET

// func setrPi32(e1 int, e0 int) x86.M64
TEXT ·setrPi32(SB),7,$0
	MOVQ e1+0(FP),R8
	MOVQ e0+8(FP),R9

	// TODO: Code missing

	// Return size: 8
	RET

// func setrPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) x86.M64
TEXT ·setrPi8(SB),7,$0
	MOVB e7+0(FP),R8
	MOVB e6+4(FP),R9
	MOVB e5+8(FP),R10
	MOVB e4+12(FP),R11
	MOVB e3+16(FP),R12
	MOVB e2+20(FP),R13
	MOVB e1+24(FP),R14
	MOVB e0+28(FP),R15

	// TODO: Code missing

	// Return size: 8
	RET

// func setzeroSi64() x86.M64
TEXT ·setzeroSi64(SB),7,$0

	// TODO: Code missing

	// Return size: 8
	RET

// func sllPi16(a x86.M64, count x86.M64) x86.M64
TEXT ·sllPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSLLW M0, M1

	// Return size: 8
	RET

// func sllPi32(a x86.M64, count x86.M64) x86.M64
TEXT ·sllPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSLLD M0, M1

	// Return size: 8
	RET

// func sllSi64(a x86.M64, count x86.M64) x86.M64
TEXT ·sllSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSLLQ M0, M1

	// Return size: 8
	RET

// func slliPi16(a x86.M64, imm8 int) x86.M64
TEXT ·slliPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSLLW M0, R9

	// Return size: 8
	RET

// func slliPi32(a x86.M64, imm8 int) x86.M64
TEXT ·slliPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSLLD M0, R9

	// Return size: 8
	RET

// func slliSi64(a x86.M64, imm8 int) x86.M64
TEXT ·slliSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSLLQ M0, R9

	// Return size: 8
	RET

// func sraPi16(a x86.M64, count x86.M64) x86.M64
TEXT ·sraPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRAW M0, M1

	// Return size: 8
	RET

// func sraPi32(a x86.M64, count x86.M64) x86.M64
TEXT ·sraPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRAD M0, M1

	// Return size: 8
	RET

// func sraiPi16(a x86.M64, imm8 int) x86.M64
TEXT ·sraiPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRAW M0, R9

	// Return size: 8
	RET

// func sraiPi32(a x86.M64, imm8 int) x86.M64
TEXT ·sraiPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRAD M0, R9

	// Return size: 8
	RET

// func srlPi16(a x86.M64, count x86.M64) x86.M64
TEXT ·srlPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRLW M0, M1

	// Return size: 8
	RET

// func srlPi32(a x86.M64, count x86.M64) x86.M64
TEXT ·srlPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRLD M0, M1

	// Return size: 8
	RET

// func srlSi64(a x86.M64, count x86.M64) x86.M64
TEXT ·srlSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ count+8(FP),M1

	// TODO: Code missing - could be:
	// PSRLQ M0, M1

	// Return size: 8
	RET

// func srliPi16(a x86.M64, imm8 int) x86.M64
TEXT ·srliPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRLW M0, R9

	// Return size: 8
	RET

// func srliPi32(a x86.M64, imm8 int) x86.M64
TEXT ·srliPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRLD M0, R9

	// Return size: 8
	RET

// func srliSi64(a x86.M64, imm8 int) x86.M64
TEXT ·srliSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSRLQ M0, R9

	// Return size: 8
	RET

// func subPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·subPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBW M0, M1

	// Return size: 8
	RET

// func subPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·subPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBD M0, M1

	// Return size: 8
	RET

// func subPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·subPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBB M0, M1

	// Return size: 8
	RET

// func subsPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·subsPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBSW M0, M1

	// Return size: 8
	RET

// func subsPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·subsPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBSB M0, M1

	// Return size: 8
	RET

// func subsPu16(a x86.M64, b x86.M64) x86.M64
TEXT ·subsPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBUSW M0, M1

	// Return size: 8
	RET

// func subsPu8(a x86.M64, b x86.M64) x86.M64
TEXT ·subsPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSUBUSB M0, M1

	// Return size: 8
	RET

// func toInt(a x86.M64) int
TEXT ·toInt(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing - could be:
	// MOVD M0

	MOVQ $0, ret+8(FP)
	RET

// func toInt64(a x86.M64) int64
TEXT ·toInt64(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing - could be:
	// MOVQ M0

	MOVQ $0, ret+8(FP)
	RET

// func unpackhiPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·unpackhiPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLBW M0, M1

	// Return size: 8
	RET

// func unpackhiPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·unpackhiPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKHDQ M0, M1

	// Return size: 8
	RET

// func unpackhiPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·unpackhiPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKHBW M0, M1

	// Return size: 8
	RET

// func unpackloPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·unpackloPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLWD M0, M1

	// Return size: 8
	RET

// func unpackloPi32(a x86.M64, b x86.M64) x86.M64
TEXT ·unpackloPi32(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLDQ M0, M1

	// Return size: 8
	RET

// func unpackloPi8(a x86.M64, b x86.M64) x86.M64
TEXT ·unpackloPi8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PUNPCKLBW M0, M1

	// Return size: 8
	RET

// func xorSi64(a x86.M64, b x86.M64) x86.M64
TEXT ·xorSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PXOR M0, M1

	// Return size: 8
	RET

