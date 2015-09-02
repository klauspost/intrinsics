// func acosPd(a [2]float64) [2]float64
TEXT ·acosPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func acosPs(a [4]float32) [4]float32
TEXT ·acosPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func acoshPd(a [2]float64) [2]float64
TEXT ·acoshPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func acoshPs(a [4]float32) [4]float32
TEXT ·acoshPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func addPs(a [4]float32, b [4]float32) [4]float32
TEXT ·addPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ADDPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func addSs(a [4]float32, b [4]float32) [4]float32
TEXT ·addSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ADDSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func andPs(a [4]float32, b [4]float32) [4]float32
TEXT ·andPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ANDPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func andnotPs(a [4]float32, b [4]float32) [4]float32
TEXT ·andnotPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ANDNPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func asinPd(a [2]float64) [2]float64
TEXT ·asinPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func asinPs(a [4]float32) [4]float32
TEXT ·asinPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func asinhPd(a [2]float64) [2]float64
TEXT ·asinhPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func asinhPs(a [4]float32) [4]float32
TEXT ·asinhPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atanPd(a [2]float64) [2]float64
TEXT ·atanPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atanPs(a [4]float32) [4]float32
TEXT ·atanPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atan2Pd(a [2]float64, b [2]float64) [2]float64
TEXT ·atan2Pd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func atan2Ps(a [4]float32, b [4]float32) [4]float32
TEXT ·atan2Ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func atanhPd(a [2]float64) [2]float64
TEXT ·atanhPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func atanhPs(a [4]float32) [4]float32
TEXT ·atanhPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func avgPu16(a x86.M64, b x86.M64) x86.M64
TEXT ·avgPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PAVGW M0, M1

	// Return size: 8
	RET

// func avgPu8(a x86.M64, b x86.M64) x86.M64
TEXT ·avgPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PAVGB M0, M1

	// Return size: 8
	RET

// func cbrtPd(a [2]float64) [2]float64
TEXT ·cbrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cbrtPs(a [4]float32) [4]float32
TEXT ·cbrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnormPd(a [2]float64) [2]float64
TEXT ·cdfnormPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnormPs(a [4]float32) [4]float32
TEXT ·cdfnormPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnorminvPd(a [2]float64) [2]float64
TEXT ·cdfnorminvPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cdfnorminvPs(a [4]float32) [4]float32
TEXT ·cdfnorminvPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cexpPs(a [4]float32) [4]float32
TEXT ·cexpPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func clogPs(a [4]float32) [4]float32
TEXT ·clogPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cmpeqPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpeqPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpeqSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpeqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpgePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpgeSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpgtPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgtPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpgtSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpgtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmplePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmplePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpleSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpleSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpltPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpltPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpltSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpneqPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpneqPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpneqSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpneqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpngePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpngeSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpngtPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngtPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpngtSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpngtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpnlePs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnlePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpnleSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnleSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpnltPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnltPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpnltSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpnltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpordPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpordPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpordSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpordSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpunordPs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpunordPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpunordSs(a [4]float32, b [4]float32) [4]float32
TEXT ·cmpunordSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// CMPSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func comieqSs(a [4]float32, b [4]float32) int
TEXT ·comieqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// COMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func comigeSs(a [4]float32, b [4]float32) int
TEXT ·comigeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// COMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func comigtSs(a [4]float32, b [4]float32) int
TEXT ·comigtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// COMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func comileSs(a [4]float32, b [4]float32) int
TEXT ·comileSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// COMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func comiltSs(a [4]float32, b [4]float32) int
TEXT ·comiltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// COMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func comineqSs(a [4]float32, b [4]float32) int
TEXT ·comineqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// COMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func cosPd(a [2]float64) [2]float64
TEXT ·cosPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cosPs(a [4]float32) [4]float32
TEXT ·cosPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cosdPd(a [2]float64) [2]float64
TEXT ·cosdPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cosdPs(a [4]float32) [4]float32
TEXT ·cosdPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func coshPd(a [2]float64) [2]float64
TEXT ·coshPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func coshPs(a [4]float32) [4]float32
TEXT ·coshPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func csqrtPs(a [4]float32) [4]float32
TEXT ·csqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func cvtPi2ps(a [4]float32, b x86.M64) [4]float32
TEXT ·cvtPi2ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),M1

	// TODO: Code missing - could be:
	// CVTPI2PS X0, M1

	MOVOU X1, ret+24(FP)
	RET

// func cvtPs2pi(a [4]float32) x86.M64
TEXT ·cvtPs2pi(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTPS2PI X0

	// Return size: 8
	RET

// func cvtSi2ss(a [4]float32, b int) [4]float32
TEXT ·cvtSi2ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing - could be:
	// CVTSI2SS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvtSs2si(a [4]float32) int
TEXT ·cvtSs2si(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvtpi16Ps(a x86.M64) [4]float32
TEXT ·cvtpi16Ps(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpi32Ps(a [4]float32, b x86.M64) [4]float32
TEXT ·cvtpi32Ps(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),M1

	// TODO: Code missing - could be:
	// CVTPI2PS X0, M1

	MOVOU X1, ret+24(FP)
	RET

// func cvtpi32x2Ps(a x86.M64, b x86.M64) [4]float32
TEXT ·cvtpi32x2Ps(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing

	MOVOU X1, ret+16(FP)
	RET

// func cvtpi8Ps(a x86.M64) [4]float32
TEXT ·cvtpi8Ps(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpsPi16(a [4]float32) x86.M64
TEXT ·cvtpsPi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	// Return size: 8
	RET

// func cvtpsPi32(a [4]float32) x86.M64
TEXT ·cvtpsPi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTPS2PI X0

	// Return size: 8
	RET

// func cvtpsPi8(a [4]float32) x86.M64
TEXT ·cvtpsPi8(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	// Return size: 8
	RET

// func cvtpu16Ps(a x86.M64) [4]float32
TEXT ·cvtpu16Ps(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtpu8Ps(a x86.M64) [4]float32
TEXT ·cvtpu8Ps(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func cvtsi32Ss(a [4]float32, b int) [4]float32
TEXT ·cvtsi32Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing - could be:
	// CVTSI2SS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvtsi64Ss(a [4]float32, b int64) [4]float32
TEXT ·cvtsi64Ss(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ b+16(FP),R9

	// TODO: Code missing - could be:
	// CVTSI2SS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func cvtssF32(a [4]float32) float32
TEXT ·cvtssF32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// MOVSS X0

	MOVL $0, ret+16(FP)
	RET

// func cvtssSi32(a [4]float32) int
TEXT ·cvtssSi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvtssSi64(a [4]float32) int64
TEXT ·cvtssSi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttPs2pi(a [4]float32) x86.M64
TEXT ·cvttPs2pi(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTTPS2PI X0

	// Return size: 8
	RET

// func cvttSs2si(a [4]float32) int
TEXT ·cvttSs2si(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttpsPi32(a [4]float32) x86.M64
TEXT ·cvttpsPi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTTPS2PI X0

	// Return size: 8
	RET

// func cvttssSi32(a [4]float32) int
TEXT ·cvttssSi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func cvttssSi64(a [4]float32) int64
TEXT ·cvttssSi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// CVTTSS2SI X0

	MOVQ $0, ret+16(FP)
	RET

// func divEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·divEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func divPs(a [4]float32, b [4]float32) [4]float32
TEXT ·divPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// DIVPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func divSs(a [4]float32, b [4]float32) [4]float32
TEXT ·divSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// DIVSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func erfPd(a [2]float64) [2]float64
TEXT ·erfPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfPs(a [4]float32) [4]float32
TEXT ·erfPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcPd(a [2]float64) [2]float64
TEXT ·erfcPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcPs(a [4]float32) [4]float32
TEXT ·erfcPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcinvPd(a [2]float64) [2]float64
TEXT ·erfcinvPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfcinvPs(a [4]float32) [4]float32
TEXT ·erfcinvPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfinvPd(a [2]float64) [2]float64
TEXT ·erfinvPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func erfinvPs(a [4]float32) [4]float32
TEXT ·erfinvPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expPd(a [2]float64) [2]float64
TEXT ·expPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expPs(a [4]float32) [4]float32
TEXT ·expPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp10Pd(a [2]float64) [2]float64
TEXT ·exp10Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp10Ps(a [4]float32) [4]float32
TEXT ·exp10Ps(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp2Pd(a [2]float64) [2]float64
TEXT ·exp2Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func exp2Ps(a [4]float32) [4]float32
TEXT ·exp2Ps(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expm1Pd(a [2]float64) [2]float64
TEXT ·expm1Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func expm1Ps(a [4]float32) [4]float32
TEXT ·expm1Ps(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func extractPi16(a x86.M64, imm8 int) int
TEXT ·extractPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PEXTRW M0, R9

	MOVQ $0, ret+16(FP)
	RET

// func free(mem_addr uintptr) 
TEXT ·free(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing

	RET

// func mMGETEXCEPTIONMASK() uint32
TEXT ·mMGETEXCEPTIONMASK(SB),7,$0

	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func mMGETEXCEPTIONSTATE() uint32
TEXT ·mMGETEXCEPTIONSTATE(SB),7,$0

	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func mMGETFLUSHZEROMODE() uint32
TEXT ·mMGETFLUSHZEROMODE(SB),7,$0

	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func mMGETROUNDINGMODE() uint32
TEXT ·mMGETROUNDINGMODE(SB),7,$0

	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func getcsr() uint32
TEXT ·getcsr(SB),7,$0

	// TODO: Code missing - uses instrunction: STMXCSR

	MOVL $0, ret+0(FP)
	RET

// func hypotPd(a [2]float64, b [2]float64) [2]float64
TEXT ·hypotPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func hypotPs(a [4]float32, b [4]float32) [4]float32
TEXT ·hypotPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func idivEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·idivEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func idivremEpi32(mem_addr [16]byte, a [16]byte, b [16]byte) [16]byte
TEXT ·idivremEpi32(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

// func insertPi16(a x86.M64, i int, imm8 int) x86.M64
TEXT ·insertPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ i+8(FP),R9
	MOVQ imm8+16(FP),R10

	// TODO: Code missing - uses instrunction: PINSRW

	// Return size: 8
	RET

// func invcbrtPd(a [2]float64) [2]float64
TEXT ·invcbrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func invcbrtPs(a [4]float32) [4]float32
TEXT ·invcbrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func invsqrtPd(a [2]float64) [2]float64
TEXT ·invsqrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func invsqrtPs(a [4]float32) [4]float32
TEXT ·invsqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func iremEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·iremEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func loadPs(mem_addr float32) [4]float32
TEXT ·loadPs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing - could be:
	// MOVAPS R8

	MOVOU X0, ret+4(FP)
	RET

// func loadPs1(mem_addr float32) [4]float32
TEXT ·loadPs1(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func loadSs(mem_addr float32) [4]float32
TEXT ·loadSs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing - could be:
	// MOVSS R8

	MOVOU X0, ret+4(FP)
	RET

// func load1Ps(mem_addr float32) [4]float32
TEXT ·load1Ps(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func loadhPi(a [4]float32, mem_addr x86.M64Const) [4]float32
TEXT ·loadhPi(SB),7,$0
	// FIXME: Unimplemented. Unknown size of type x86.M64Const

	// TODO: Code missing - could be:
	// MOVHPS X0, 

	MOVOU X1, ret+0(FP)
	RET

// func loadlPi(a [4]float32, mem_addr x86.M64Const) [4]float32
TEXT ·loadlPi(SB),7,$0
	// FIXME: Unimplemented. Unknown size of type x86.M64Const

	// TODO: Code missing - could be:
	// MOVLPS X0, 

	MOVOU X1, ret+0(FP)
	RET

// func loadrPs(mem_addr float32) [4]float32
TEXT ·loadrPs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func loaduPs(mem_addr float32) [4]float32
TEXT ·loaduPs(SB),7,$0
	MOVL mem_addr+0(FP),R8

	// TODO: Code missing - could be:
	// MOVUPS R8

	MOVOU X0, ret+4(FP)
	RET

// func loaduSi16(mem_addr uintptr) [16]byte
TEXT ·loaduSi16(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+8(FP)
	RET

// func loaduSi32(mem_addr uintptr) [16]byte
TEXT ·loaduSi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing - could be:
	// MOVD R8

	MOVOU X0, ret+8(FP)
	RET

// func loaduSi64(mem_addr uintptr) [16]byte
TEXT ·loaduSi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing - could be:
	// MOVQ R8

	MOVOU X0, ret+8(FP)
	RET

// func logPd(a [2]float64) [2]float64
TEXT ·logPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func logPs(a [4]float32) [4]float32
TEXT ·logPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log10Pd(a [2]float64) [2]float64
TEXT ·log10Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log10Ps(a [4]float32) [4]float32
TEXT ·log10Ps(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log1pPd(a [2]float64) [2]float64
TEXT ·log1pPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log1pPs(a [4]float32) [4]float32
TEXT ·log1pPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log2Pd(a [2]float64) [2]float64
TEXT ·log2Pd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func log2Ps(a [4]float32) [4]float32
TEXT ·log2Ps(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func logbPd(a [2]float64) [2]float64
TEXT ·logbPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func logbPs(a [4]float32) [4]float32
TEXT ·logbPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func malloc(size int, align int) 
TEXT ·malloc(SB),7,$0
	MOVQ size+0(FP),R8
	MOVQ align+8(FP),R9

	// TODO: Code missing

	RET

// func maskmoveSi64(a x86.M64, mask x86.M64, mem_addr byte) 
TEXT ·maskmoveSi64(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ mask+8(FP),M1
	MOVB mem_addr+16(FP),R10

	// TODO: Code missing - uses instrunction: MASKMOVQ

	RET

// func maskmovq(a x86.M64, mask x86.M64, mem_addr byte) 
TEXT ·maskmovq(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ mask+8(FP),M1
	MOVB mem_addr+16(FP),R10

	// TODO: Code missing - uses instrunction: MASKMOVQ

	RET

// func maxPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·maxPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMAXSW M0, M1

	// Return size: 8
	RET

// func maxPs(a [4]float32, b [4]float32) [4]float32
TEXT ·maxPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MAXPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maxPu8(a x86.M64, b x86.M64) x86.M64
TEXT ·maxPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMAXUB M0, M1

	// Return size: 8
	RET

// func maxSs(a [4]float32, b [4]float32) [4]float32
TEXT ·maxSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MAXSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minPi16(a x86.M64, b x86.M64) x86.M64
TEXT ·minPi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMINSW M0, M1

	// Return size: 8
	RET

// func minPs(a [4]float32, b [4]float32) [4]float32
TEXT ·minPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MINPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minPu8(a x86.M64, b x86.M64) x86.M64
TEXT ·minPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMINUB M0, M1

	// Return size: 8
	RET

// func minSs(a [4]float32, b [4]float32) [4]float32
TEXT ·minSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MINSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func moveSs(a [4]float32, b [4]float32) [4]float32
TEXT ·moveSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MOVSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func movehlPs(a [4]float32, b [4]float32) [4]float32
TEXT ·movehlPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MOVHLPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func movelhPs(a [4]float32, b [4]float32) [4]float32
TEXT ·movelhPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MOVLHPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func movemaskPi8(a x86.M64) int
TEXT ·movemaskPi8(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing - could be:
	// PMOVMSKB M0

	MOVQ $0, ret+8(FP)
	RET

// func movemaskPs(a [4]float32) int
TEXT ·movemaskPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// MOVMSKPS X0

	MOVQ $0, ret+16(FP)
	RET

// func mulPs(a [4]float32, b [4]float32) [4]float32
TEXT ·mulPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MULPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func mulSs(a [4]float32, b [4]float32) [4]float32
TEXT ·mulSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// MULSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func mulhiPu16(a x86.M64, b x86.M64) x86.M64
TEXT ·mulhiPu16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMULHUW M0, M1

	// Return size: 8
	RET

// func orPs(a [4]float32, b [4]float32) [4]float32
TEXT ·orPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ORPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func pavgb(a x86.M64, b x86.M64) x86.M64
TEXT ·pavgb(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PAVGB M0, M1

	// Return size: 8
	RET

// func pavgw(a x86.M64, b x86.M64) x86.M64
TEXT ·pavgw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PAVGW M0, M1

	// Return size: 8
	RET

// func pextrw(a x86.M64, imm8 int) int
TEXT ·pextrw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PEXTRW M0, R9

	MOVQ $0, ret+16(FP)
	RET

// func pinsrw(a x86.M64, i int, imm8 int) x86.M64
TEXT ·pinsrw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ i+8(FP),R9
	MOVQ imm8+16(FP),R10

	// TODO: Code missing - uses instrunction: PINSRW

	// Return size: 8
	RET

// func pmaxsw(a x86.M64, b x86.M64) x86.M64
TEXT ·pmaxsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMAXSW M0, M1

	// Return size: 8
	RET

// func pmaxub(a x86.M64, b x86.M64) x86.M64
TEXT ·pmaxub(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMAXUB M0, M1

	// Return size: 8
	RET

// func pminsw(a x86.M64, b x86.M64) x86.M64
TEXT ·pminsw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMINSW M0, M1

	// Return size: 8
	RET

// func pminub(a x86.M64, b x86.M64) x86.M64
TEXT ·pminub(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMINUB M0, M1

	// Return size: 8
	RET

// func pmovmskb(a x86.M64) int
TEXT ·pmovmskb(SB),7,$0
	MOVQ a+0(FP),M0

	// TODO: Code missing - could be:
	// PMOVMSKB M0

	MOVQ $0, ret+8(FP)
	RET

// func pmulhuw(a x86.M64, b x86.M64) x86.M64
TEXT ·pmulhuw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PMULHUW M0, M1

	// Return size: 8
	RET

// func powPd(a [2]float64, b [2]float64) [2]float64
TEXT ·powPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func powPs(a [4]float32, b [4]float32) [4]float32
TEXT ·powPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func prefetch(p byte, i int) 
TEXT ·prefetch(SB),7,$0
	MOVB p+0(FP),R8
	MOVQ i+4(FP),R9

	// TODO: Code missing - could be:
	// PREFETCHNTA, PREFETCHT0, PREFETCHT1, PREFETCHT2 R8, R9

	RET

// func psadbw(a x86.M64, b x86.M64) x86.M64
TEXT ·psadbw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSADBW M0, M1

	// Return size: 8
	RET

// func pshufw(a x86.M64, imm8 int) x86.M64
TEXT ·pshufw(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSHUFW M0, R9

	// Return size: 8
	RET

// func rcpPs(a [4]float32) [4]float32
TEXT ·rcpPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// RCPPS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func rcpSs(a [4]float32) [4]float32
TEXT ·rcpSs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// RCPSS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func remEpi16(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpu64(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func remEpu8(a [16]byte, b [16]byte) [16]byte
TEXT ·remEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func rsqrtPs(a [4]float32) [4]float32
TEXT ·rsqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// RSQRTPS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func rsqrtSs(a [4]float32) [4]float32
TEXT ·rsqrtSs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// RSQRTSS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func sadPu8(a x86.M64, b x86.M64) x86.M64
TEXT ·sadPu8(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ b+8(FP),M1

	// TODO: Code missing - could be:
	// PSADBW M0, M1

	// Return size: 8
	RET

// func mMSETEXCEPTIONMASK(a uint32) 
TEXT ·mMSETEXCEPTIONMASK(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	RET

// func mMSETEXCEPTIONSTATE(a uint32) 
TEXT ·mMSETEXCEPTIONSTATE(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	RET

// func mMSETFLUSHZEROMODE(a uint32) 
TEXT ·mMSETFLUSHZEROMODE(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	RET

// func setPs(e3 float32, e2 float32, e1 float32, e0 float32) [4]float32
TEXT ·setPs(SB),7,$0
	MOVL e3+0(FP),R8
	MOVL e2+4(FP),R9
	MOVL e1+8(FP),R10
	MOVL e0+12(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+16(FP)
	RET

// func setPs1(a float32) [4]float32
TEXT ·setPs1(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func mMSETROUNDINGMODE(a uint32) 
TEXT ·mMSETROUNDINGMODE(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	RET

// func setSs(a float32) [4]float32
TEXT ·setSs(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func set1Ps(a float32) [4]float32
TEXT ·set1Ps(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing

	MOVOU X0, ret+4(FP)
	RET

// func setcsr(a uint32) 
TEXT ·setcsr(SB),7,$0
	MOVL a+0(FP),R8

	// TODO: Code missing - could be:
	// LDMXCSR R8

	RET

// func setrPs(e3 float32, e2 float32, e1 float32, e0 float32) [4]float32
TEXT ·setrPs(SB),7,$0
	MOVL e3+0(FP),R8
	MOVL e2+4(FP),R9
	MOVL e1+8(FP),R10
	MOVL e0+12(FP),R11

	// TODO: Code missing

	MOVOU X3, ret+16(FP)
	RET

// func setzeroPs() [4]float32
TEXT ·setzeroPs(SB),7,$0

	// TODO: Code missing - uses instrunction: XORPS

	MOVOU X-1, ret+0(FP)
	RET

// func sfence() 
TEXT ·sfence(SB),7,$0

	// TODO: Code missing - uses instrunction: SFENCE

	RET

// func shufflePi16(a x86.M64, imm8 int) x86.M64
TEXT ·shufflePi16(SB),7,$0
	MOVQ a+0(FP),M0
	MOVQ imm8+8(FP),R9

	// TODO: Code missing - could be:
	// PSHUFW M0, R9

	// Return size: 8
	RET

// func shufflePs(a [4]float32, b [4]float32, imm8 uint32) [4]float32
TEXT ·shufflePs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVL imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: SHUFPS

	MOVOU X2, ret+36(FP)
	RET

// func sinPd(a [2]float64) [2]float64
TEXT ·sinPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sinPs(a [4]float32) [4]float32
TEXT ·sinPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sincosPd(mem_addr [2]float64, a [2]float64) [2]float64
TEXT ·sincosPd(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func sincosPs(mem_addr [4]float32, a [4]float32) [4]float32
TEXT ·sincosPs(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func sindPd(a [2]float64) [2]float64
TEXT ·sindPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sindPs(a [4]float32) [4]float32
TEXT ·sindPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sinhPd(a [2]float64) [2]float64
TEXT ·sinhPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sinhPs(a [4]float32) [4]float32
TEXT ·sinhPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func sqrtPs(a [4]float32) [4]float32
TEXT ·sqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// SQRTPS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func sqrtSs(a [4]float32) [4]float32
TEXT ·sqrtSs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// SQRTSS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func storePs(mem_addr float32, a [4]float32) 
TEXT ·storePs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// MOVAPS R8, X1

	RET

// func storePs1(mem_addr float32, a [4]float32) 
TEXT ·storePs1(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing

	RET

// func storeSs(mem_addr float32, a [4]float32) 
TEXT ·storeSs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// MOVSS R8, X1

	RET

// func store1Ps(mem_addr float32, a [4]float32) 
TEXT ·store1Ps(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing

	RET

// func storehPi(mem_addr x86.M64, a [4]float32) 
TEXT ·storehPi(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVOU a+8(FP),X1

	// TODO: Code missing - could be:
	// MOVHPS M0, X1

	RET

// func storelPi(mem_addr x86.M64, a [4]float32) 
TEXT ·storelPi(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVOU a+8(FP),X1

	// TODO: Code missing - could be:
	// MOVLPS M0, X1

	RET

// func storerPs(mem_addr float32, a [4]float32) 
TEXT ·storerPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing

	RET

// func storeuPs(mem_addr float32, a [4]float32) 
TEXT ·storeuPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// MOVUPS R8, X1

	RET

// func storeuSi16(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi16(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	// TODO: Code missing

	RET

// func storeuSi32(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	// TODO: Code missing - could be:
	// MOVD R8, X1

	RET

// func storeuSi64(mem_addr uintptr, a [16]byte) 
TEXT ·storeuSi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU a+8(FP),X1

	// TODO: Code missing - could be:
	// MOVQ R8, X1

	RET

// func streamPi(mem_addr x86.M64, a x86.M64) 
TEXT ·streamPi(SB),7,$0
	MOVQ mem_addr+0(FP),M0
	MOVQ a+8(FP),M1

	// TODO: Code missing - could be:
	// MOVNTQ M0, M1

	RET

// func streamPs(mem_addr float32, a [4]float32) 
TEXT ·streamPs(SB),7,$0
	MOVL mem_addr+0(FP),R8
	MOVOU a+4(FP),X1

	// TODO: Code missing - could be:
	// MOVNTPS R8, X1

	RET

// func subPs(a [4]float32, b [4]float32) [4]float32
TEXT ·subPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// SUBPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func subSs(a [4]float32, b [4]float32) [4]float32
TEXT ·subSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// SUBSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func svmlCeilPd(a [2]float64) [2]float64
TEXT ·svmlCeilPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlCeilPs(a [4]float32) [4]float32
TEXT ·svmlCeilPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlFloorPd(a [2]float64) [2]float64
TEXT ·svmlFloorPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlFloorPs(a [4]float32) [4]float32
TEXT ·svmlFloorPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlRoundPd(a [2]float64) [2]float64
TEXT ·svmlRoundPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlRoundPs(a [4]float32) [4]float32
TEXT ·svmlRoundPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlSqrtPd(a [2]float64) [2]float64
TEXT ·svmlSqrtPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func svmlSqrtPs(a [4]float32) [4]float32
TEXT ·svmlSqrtPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanPd(a [2]float64) [2]float64
TEXT ·tanPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanPs(a [4]float32) [4]float32
TEXT ·tanPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tandPd(a [2]float64) [2]float64
TEXT ·tandPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tandPs(a [4]float32) [4]float32
TEXT ·tandPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanhPd(a [2]float64) [2]float64
TEXT ·tanhPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func tanhPs(a [4]float32) [4]float32
TEXT ·tanhPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func mMTRANSPOSE4PS(row0 [4]float32, row1 [4]float32, row2 [4]float32, row3 [4]float32) 
TEXT ·mMTRANSPOSE4PS(SB),7,$0
	MOVOU row0+0(FP),X0
	MOVOU row1+16(FP),X1
	MOVOU row2+32(FP),X2
	MOVOU row3+48(FP),X3

	// TODO: Code missing

	RET

// func truncPd(a [2]float64) [2]float64
TEXT ·truncPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func truncPs(a [4]float32) [4]float32
TEXT ·truncPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func ucomieqSs(a [4]float32, b [4]float32) int
TEXT ·ucomieqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UCOMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func ucomigeSs(a [4]float32, b [4]float32) int
TEXT ·ucomigeSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UCOMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func ucomigtSs(a [4]float32, b [4]float32) int
TEXT ·ucomigtSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UCOMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func ucomileSs(a [4]float32, b [4]float32) int
TEXT ·ucomileSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UCOMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func ucomiltSs(a [4]float32, b [4]float32) int
TEXT ·ucomiltSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UCOMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func ucomineqSs(a [4]float32, b [4]float32) int
TEXT ·ucomineqSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UCOMISS X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func udivEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·udivEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func udivremEpi32(mem_addr [16]byte, a [16]byte, b [16]byte) [16]byte
TEXT ·udivremEpi32(SB),7,$0
	MOVOU mem_addr+0(FP),X0
	MOVOU a+16(FP),X1
	MOVOU b+32(FP),X2

	// TODO: Code missing

	MOVOU X2, ret+48(FP)
	RET

// func unpackhiPs(a [4]float32, b [4]float32) [4]float32
TEXT ·unpackhiPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UNPCKHPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func unpackloPs(a [4]float32, b [4]float32) [4]float32
TEXT ·unpackloPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// UNPCKLPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func uremEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·uremEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing

	MOVOU X1, ret+32(FP)
	RET

// func xorPs(a [4]float32, b [4]float32) [4]float32
TEXT ·xorPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// XORPS X0, X1

	MOVOU X1, ret+32(FP)
	RET

