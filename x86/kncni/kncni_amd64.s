// func m512AbsPd(v2 [8]float64) [8]float64
TEXT ·m512AbsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VPANDQ Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAbsPd(src [8]float64, k uint8, v2 [8]float64) [8]float64
TEXT ·m512MaskAbsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AbsPs(v2 [16]float32) [16]float32
TEXT ·m512AbsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VPANDD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskAbsPs(src [16]float32, k uint16, v2 [16]float32) [16]float32
TEXT ·m512MaskAbsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512AdcEpi32(v2 [64]byte, k2 uint16, v3 [64]byte, k2_res uint16) [64]byte
TEXT ·m512AdcEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskAdcEpi32(v2 [64]byte, k1 uint16, k2 uint16, v3 [64]byte, k2_res uint16) [64]byte
TEXT ·m512MaskAdcEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPADDD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAddEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AddPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512AddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VADDPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAddPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskAddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AddPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512AddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VADDPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAddPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskAddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AddRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512AddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskAddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512AddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskAddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddnPd(v2 [8]float64, v3 [8]float64) [8]float64
TEXT ·m512AddnPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VADDNPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAddnPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64) [8]float64
TEXT ·m512MaskAddnPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AddnPs(v2 [16]float32, v3 [16]float32) [16]float32
TEXT ·m512AddnPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VADDNPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAddnPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32) [16]float32
TEXT ·m512MaskAddnPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AddnRoundPd(v2 [8]float64, v3 [8]float64, rounding int) [8]float64
TEXT ·m512AddnRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddnRoundPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64, rounding int) [8]float64
TEXT ·m512MaskAddnRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddnRoundPs(v2 [16]float32, v3 [16]float32, rounding int) [16]float32
TEXT ·m512AddnRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddnRoundPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, rounding int) [16]float32
TEXT ·m512MaskAddnRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddsetcEpi32(v2 [64]byte, v3 [64]byte, k2_res uint16) [64]byte
TEXT ·m512AddsetcEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddsetcEpi32(v2 [64]byte, k uint16, k_old uint16, v3 [64]byte, k2_res uint16) [64]byte
TEXT ·m512MaskAddsetcEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddsetsEpi32(v2 [64]byte, v3 [64]byte, sign uint16) [64]byte
TEXT ·m512AddsetsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddsetsEpi32(src [64]byte, k uint16, v2 [64]byte, v3 [64]byte, sign uint16) [64]byte
TEXT ·m512MaskAddsetsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddsetsPs(v2 [16]float32, v3 [16]float32, sign uint16) [16]float32
TEXT ·m512AddsetsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAddsetsPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, sign uint16) [16]float32
TEXT ·m512MaskAddsetsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AddsetsRoundPs(v2 [16]float32, v3 [16]float32, sign uint16, rounding int) [16]float32
TEXT ·m512AddsetsRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskAddsetsRoundPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, sign uint16, rounding int) [16]float32
TEXT ·m512MaskAddsetsRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512AlignrEpi32(a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512AlignrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskAlignrEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte, count int) [64]byte
TEXT ·m512MaskAlignrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512AndEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPANDD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAndEpi32(src [64]byte, k uint16, v2 [64]byte, v3 [64]byte) [64]byte
TEXT ·m512MaskAndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AndEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPANDQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAndEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AndSi512(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AndSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPANDD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512AndnotEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AndnotEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPANDND Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAndnotEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAndnotEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AndnotEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AndnotEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPANDNQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskAndnotEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskAndnotEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512AndnotSi512(a [64]byte, b [64]byte) [64]byte
TEXT ·m512AndnotSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPANDND Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskBlendEpi32(k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskBlendEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskBlendEpi64(k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskBlendEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskBlendPd(k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskBlendPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskBlendPs(k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskBlendPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CastpdPs(a [8]float64) [16]float32
TEXT ·m512CastpdPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512CastpdSi512(a [8]float64) [64]byte
TEXT ·m512CastpdSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512CastpsPd(a [16]float32) [8]float64
TEXT ·m512CastpsPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512CastpsSi512(a [16]float32) [64]byte
TEXT ·m512CastpsSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Castsi512Pd(a [64]byte) [8]float64
TEXT ·m512Castsi512Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func m512Castsi512Ps(a [64]byte) [16]float32
TEXT ·m512Castsi512Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z0, ret+0(FP)
	RET

// func clevict(ptr uintptr, level int) 
TEXT ·clevict(SB),7,$0
	MOVQ ptr+0(FP),R8
	MOVQ level+8(FP),R9

	// TODO: Code missing
	// Could be:
	// CLEVICT0, CLEVICT1 R8, R9

	RET

// func m512CmpEpi32Mask(a [64]byte, b [64]byte, imm8 uint8) uint16
TEXT ·m512CmpEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpEpi32Mask(k1 uint16, a [64]byte, b [64]byte, imm8 uint8) uint16
TEXT ·m512MaskCmpEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpEpu32Mask(a [64]byte, b [64]byte, imm8 uint8) uint16
TEXT ·m512CmpEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpEpu32Mask(k1 uint16, a [64]byte, b [64]byte, imm8 uint8) uint16
TEXT ·m512MaskCmpEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpPdMask(a [8]float64, b [8]float64, imm8 int) uint8
TEXT ·m512CmpPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpPdMask(k1 uint8, a [8]float64, b [8]float64, imm8 int) uint8
TEXT ·m512MaskCmpPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpPsMask(a [16]float32, b [16]float32, imm8 int) uint16
TEXT ·m512CmpPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpPsMask(k1 uint16, a [16]float32, b [16]float32, imm8 int) uint16
TEXT ·m512MaskCmpPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpRoundPdMask(a [8]float64, b [8]float64, imm8 int, sae int) uint8
TEXT ·m512CmpRoundPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpRoundPdMask(k1 uint8, a [8]float64, b [8]float64, imm8 int, sae int) uint8
TEXT ·m512MaskCmpRoundPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpRoundPsMask(a [16]float32, b [16]float32, imm8 int, sae int) uint16
TEXT ·m512CmpRoundPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpRoundPsMask(k1 uint16, a [16]float32, b [16]float32, imm8 int, sae int) uint16
TEXT ·m512MaskCmpRoundPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpeqEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpeqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPEQD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpeqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpeqEpu32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpeqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpeqEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpeqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpeqPdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpeqPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpeqPdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpeqPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpeqPsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpeqPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpeqPsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpeqPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgeEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpgeEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpgeEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgeEpu32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpgeEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpgeEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpgeEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgtEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpgtEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPGTD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpgtEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpgtEpu32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpgtEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpgtEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpgtEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpleEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpleEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpleEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpleEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpleEpu32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpleEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpleEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpleEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmplePdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmplePdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmplePdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmplePdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmplePsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmplePsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmplePsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmplePsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpltEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpltEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPLTD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpltEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpltEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpltEpu32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpltEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpltEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpltEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpltPdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpltPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpltPdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpltPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpltPsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpltPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpltPsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpltPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpneqEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpneqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpneqEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpneqEpu32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512CmpneqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPCMPUD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpneqEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskCmpneqEpu32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpneqPdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpneqPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpneqPdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpneqPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpneqPsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpneqPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpneqPsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpneqPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpnlePdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpnlePdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpnlePdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpnlePdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpnlePsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpnlePsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpnlePsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpnlePsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpnltPdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpnltPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpnltPdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpnltPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpnltPsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpnltPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpnltPsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpnltPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpordPdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpordPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpordPdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpordPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpordPsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpordPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpordPsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpordPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512CmpunordPdMask(a [8]float64, b [8]float64) uint8
TEXT ·m512CmpunordPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCMPPD Z0, Z1

	MOVB $0, ret+0(FP)
	RET

// func m512MaskCmpunordPdMask(k1 uint8, a [8]float64, b [8]float64) uint8
TEXT ·m512MaskCmpunordPdMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVB $0, ret+0(FP)
	RET

// func m512CmpunordPsMask(a [16]float32, b [16]float32) uint16
TEXT ·m512CmpunordPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCMPPS Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func m512MaskCmpunordPsMask(k1 uint16, a [16]float32, b [16]float32) uint16
TEXT ·m512MaskCmpunordPsMask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func countbits32(r1 uint32) uint32
TEXT ·countbits32(SB),7,$0
	MOVL r1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// POPCNT R8

	MOVL $0, ret+4(FP)
	RET

// func countbits64(r1 uint64) uint64
TEXT ·countbits64(SB),7,$0
	MOVQ r1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// POPCNT R8

	MOVQ $0, ret+8(FP)
	RET

// func m512CvtRoundpdPslo(v2 [8]float64, rounding int) [16]float32
TEXT ·m512CvtRoundpdPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2PS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtRoundpdPslo(src [16]float32, k uint8, v2 [8]float64, rounding int) [16]float32
TEXT ·m512MaskCvtRoundpdPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Cvtepi32loPd(v2 [64]byte) [8]float64
TEXT ·m512Cvtepi32loPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTDQ2PD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepi32loPd(src [8]float64, k uint8, v2 [64]byte) [8]float64
TEXT ·m512MaskCvtepi32loPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Cvtepu32loPd(v2 [64]byte) [8]float64
TEXT ·m512Cvtepu32loPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VCVTUDQ2PD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtepu32loPd(src [8]float64, k uint8, v2 [64]byte) [8]float64
TEXT ·m512MaskCvtepu32loPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtfxpntRoundAdjustepi32Ps(v2 [64]byte, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512CvtfxpntRoundAdjustepi32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtfxpntRoundAdjustepu32Ps(v2 [64]byte, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512CvtfxpntRoundAdjustepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskCvtfxpntRoundAdjustepu32Ps(src [16]float32, k uint16, v2 [64]byte, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512MaskCvtfxpntRoundAdjustepu32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512CvtfxpntRoundAdjustpsEpi32(v2 [16]float32, rounding int, expadj MMEXPADJENUM) [64]byte
TEXT ·m512CvtfxpntRoundAdjustpsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtfxpntRoundAdjustpsEpu32(v2 [16]float32, rounding int, expadj MMEXPADJENUM) [64]byte
TEXT ·m512CvtfxpntRoundAdjustpsEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtfxpntRoundpdEpi32lo(v2 [8]float64, rounding int) [64]byte
TEXT ·m512CvtfxpntRoundpdEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTFXPNTPD2DQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtfxpntRoundpdEpi32lo(src [64]byte, k uint8, v2 [8]float64, rounding int) [64]byte
TEXT ·m512MaskCvtfxpntRoundpdEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512CvtfxpntRoundpdEpu32lo(v2 [8]float64, rounding int) [64]byte
TEXT ·m512CvtfxpntRoundpdEpu32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTFXPNTPD2UDQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskCvtfxpntRoundpdEpu32lo(src [64]byte, k uint8, v2 [8]float64, rounding int) [64]byte
TEXT ·m512MaskCvtfxpntRoundpdEpu32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512CvtpdPslo(v2 [8]float64) [16]float32
TEXT ·m512CvtpdPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VCVTPD2PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtpdPslo(src [16]float32, k uint8, v2 [8]float64) [16]float32
TEXT ·m512MaskCvtpdPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512CvtpsloPd(v2 [16]float32) [8]float64
TEXT ·m512CvtpsloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VCVTPS2PD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskCvtpsloPd(src [8]float64, k uint8, v2 [16]float32) [8]float64
TEXT ·m512MaskCvtpsloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func delay32(r1 uint32) 
TEXT ·delay32(SB),7,$0
	MOVL r1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// DELAY R8

	RET

// func delay64(r1 uint64) 
TEXT ·delay64(SB),7,$0
	MOVQ r1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// DELAY R8

	RET

// func m512Exp223Ps(v2 [64]byte) [16]float32
TEXT ·m512Exp223Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VEXP223PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskExp223Ps(src [16]float32, k uint16, v2 [64]byte) [16]float32
TEXT ·m512MaskExp223Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512ExtloadEpi32(mt uintptr, conv MMUPCONVEPI32ENUM, bc MMBROADCAST32ENUM, hint int) [64]byte
TEXT ·m512ExtloadEpi32(SB),7,$0
	// Unimplemented. Unknown size of type MMUPCONVEPI32ENUM
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadEpi32(src [64]byte, k uint16, mt uintptr, conv MMUPCONVEPI32ENUM, bc MMBROADCAST32ENUM, hint int) [64]byte
TEXT ·m512MaskExtloadEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512ExtloadEpi64(mt uintptr, conv MMUPCONVEPI64ENUM, bc MMBROADCAST64ENUM, hint int) [64]byte
TEXT ·m512ExtloadEpi64(SB),7,$0
	// Unimplemented. Unknown size of type MMUPCONVEPI64ENUM
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadEpi64(src [64]byte, k uint8, mt uintptr, conv MMUPCONVEPI64ENUM, bc MMBROADCAST64ENUM, hint int) [64]byte
TEXT ·m512MaskExtloadEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512ExtloadPd(mt uintptr, conv MMUPCONVPDENUM, bc MMBROADCAST64ENUM, hint int) [8]float64
TEXT ·m512ExtloadPd(SB),7,$0
	// Unimplemented. Unknown size of type MMUPCONVPDENUM
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadPd(src [8]float64, k uint8, mt uintptr, conv MMUPCONVPDENUM, bc MMBROADCAST64ENUM, hint int) [8]float64
TEXT ·m512MaskExtloadPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512ExtloadPs(mt uintptr, conv MMUPCONVPSENUM, bc MMBROADCAST32ENUM, hint int) [16]float32
TEXT ·m512ExtloadPs(SB),7,$0
	// Unimplemented. Unknown size of type MMUPCONVPSENUM
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadPs(src [16]float32, k uint16, mt uintptr, conv MMUPCONVPSENUM, bc MMBROADCAST32ENUM, hint int) [16]float32
TEXT ·m512MaskExtloadPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512ExtloadunpackhiEpi32(src [64]byte, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte
TEXT ·m512ExtloadunpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackhiEpi32(src [64]byte, k uint16, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte
TEXT ·m512MaskExtloadunpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackhiEpi64(src [64]byte, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte
TEXT ·m512ExtloadunpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackhiEpi64(src [64]byte, k uint8, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte
TEXT ·m512MaskExtloadunpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackhiPd(src [8]float64, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64
TEXT ·m512ExtloadunpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackhiPd(src [8]float64, k uint8, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64
TEXT ·m512MaskExtloadunpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackhiPs(src [16]float32, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32
TEXT ·m512ExtloadunpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackhiPs(src [16]float32, k uint16, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32
TEXT ·m512MaskExtloadunpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackloEpi32(src [64]byte, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte
TEXT ·m512ExtloadunpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackloEpi32(src [64]byte, k uint16, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte
TEXT ·m512MaskExtloadunpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackloEpi64(src [64]byte, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte
TEXT ·m512ExtloadunpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackloEpi64(src [64]byte, k uint8, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte
TEXT ·m512MaskExtloadunpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackloPd(src [8]float64, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64
TEXT ·m512ExtloadunpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackloPd(src [8]float64, k uint8, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64
TEXT ·m512MaskExtloadunpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtloadunpackloPs(src [16]float32, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32
TEXT ·m512ExtloadunpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskExtloadunpackloPs(src [16]float32, k uint16, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32
TEXT ·m512MaskExtloadunpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ExtpackstorehiEpi32(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 
TEXT ·m512ExtpackstorehiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskExtpackstorehiEpi32(mt uintptr, k uint16, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 
TEXT ·m512MaskExtpackstorehiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512ExtpackstorehiEpi64(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 
TEXT ·m512ExtpackstorehiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskExtpackstorehiEpi64(mt uintptr, k uint8, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 
TEXT ·m512MaskExtpackstorehiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512ExtpackstorehiPd(mt uintptr, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 
TEXT ·m512ExtpackstorehiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512MaskExtpackstorehiPd(mt uintptr, k uint8, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 
TEXT ·m512MaskExtpackstorehiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512ExtpackstorehiPs(mt uintptr, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 
TEXT ·m512ExtpackstorehiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512MaskExtpackstorehiPs(mt uintptr, k uint16, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 
TEXT ·m512MaskExtpackstorehiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512ExtpackstoreloEpi32(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 
TEXT ·m512ExtpackstoreloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskExtpackstoreloEpi32(mt uintptr, k uint16, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 
TEXT ·m512MaskExtpackstoreloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512ExtpackstoreloEpi64(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 
TEXT ·m512ExtpackstoreloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskExtpackstoreloEpi64(mt uintptr, k uint8, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 
TEXT ·m512MaskExtpackstoreloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512ExtpackstoreloPd(mt uintptr, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 
TEXT ·m512ExtpackstoreloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512MaskExtpackstoreloPd(mt uintptr, k uint8, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 
TEXT ·m512MaskExtpackstoreloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512ExtpackstoreloPs(mt uintptr, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 
TEXT ·m512ExtpackstoreloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512MaskExtpackstoreloPs(mt uintptr, k uint16, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 
TEXT ·m512MaskExtpackstoreloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512ExtstoreEpi32(mt uintptr, v [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 
TEXT ·m512ExtstoreEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskExtstoreEpi32(mt uintptr, k uint16, v [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 
TEXT ·m512MaskExtstoreEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512ExtstoreEpi64(mt uintptr, v [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 
TEXT ·m512ExtstoreEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskExtstoreEpi64(mt uintptr, k uint8, v [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 
TEXT ·m512MaskExtstoreEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512ExtstorePd(mt uintptr, v [8]float64, conv MMDOWNCONVPDENUM, hint int) 
TEXT ·m512ExtstorePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512MaskExtstorePd(mt uintptr, k uint8, v [8]float64, conv MMDOWNCONVPDENUM, hint int) 
TEXT ·m512MaskExtstorePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512ExtstorePs(mt uintptr, v [16]float32, conv MMDOWNCONVPSENUM, hint int) 
TEXT ·m512ExtstorePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512MaskExtstorePs(mt uintptr, k uint16, v [16]float32, conv MMDOWNCONVPSENUM, hint int) 
TEXT ·m512MaskExtstorePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512FixupnanPd(v1 [8]float64, v2 [8]float64, v3 [64]byte) [8]float64
TEXT ·m512FixupnanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFixupnanPd(v1 [8]float64, k uint8, v2 [8]float64, v3 [64]byte) [8]float64
TEXT ·m512MaskFixupnanPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FixupnanPs(v1 [16]float32, v2 [16]float32, v3 [64]byte) [16]float32
TEXT ·m512FixupnanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFixupnanPs(v1 [16]float32, k uint16, v2 [16]float32, v3 [64]byte) [16]float32
TEXT ·m512MaskFixupnanPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmaddEpi32(a [64]byte, b [64]byte, c [64]byte) [64]byte
TEXT ·m512FmaddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmaddEpi32(a [64]byte, k uint16, b [64]byte, c [64]byte) [64]byte
TEXT ·m512MaskFmaddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmaddEpi32(a [64]byte, b [64]byte, c [64]byte, k uint16) [64]byte
TEXT ·m512Mask3FmaddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmaddPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512FmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmaddPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskFmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmaddPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·m512Mask3FmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmaddPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512FmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmaddPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskFmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmaddPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·m512Mask3FmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512FmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmaddRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskFmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·m512Mask3FmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512FmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmaddRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskFmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·m512Mask3FmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Fmadd233Epi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512Fmadd233Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMADD233D Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskFmadd233Epi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskFmadd233Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Fmadd233Ps(a [16]float32, b [16]float32) [16]float32
TEXT ·m512Fmadd233Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VFMADD233PS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskFmadd233Ps(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskFmadd233Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Fmadd233RoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512Fmadd233RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmadd233RoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskFmadd233RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FmsubPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512FmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmsubPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskFmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmsubPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·m512Mask3FmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmsubPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512FmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFmsubPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskFmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FmsubPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·m512Mask3FmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512FmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmsubRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskFmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·m512Mask3FmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512FmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFmsubRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskFmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·m512Mask3FmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FnmaddPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512FnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFnmaddPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskFnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FnmaddPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·m512Mask3FnmaddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FnmaddPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512FnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFnmaddPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskFnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FnmaddPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·m512Mask3FnmaddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FnmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512FnmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFnmaddRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskFnmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FnmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·m512Mask3FnmaddRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FnmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512FnmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFnmaddRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskFnmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FnmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·m512Mask3FnmaddRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FnmsubPd(a [8]float64, b [8]float64, c [8]float64) [8]float64
TEXT ·m512FnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFnmsubPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64
TEXT ·m512MaskFnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FnmsubPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64
TEXT ·m512Mask3FnmsubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FnmsubPs(a [16]float32, b [16]float32, c [16]float32) [16]float32
TEXT ·m512FnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskFnmsubPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32
TEXT ·m512MaskFnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Mask3FnmsubPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32
TEXT ·m512Mask3FnmsubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512FnmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512FnmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFnmsubRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64
TEXT ·m512MaskFnmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FnmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64
TEXT ·m512Mask3FnmsubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512FnmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512FnmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskFnmsubRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32
TEXT ·m512MaskFnmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512Mask3FnmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32
TEXT ·m512Mask3FnmsubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512GetexpPd(a [8]float64) [8]float64
TEXT ·m512GetexpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VGETEXPPD Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskGetexpPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskGetexpPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512GetexpPs(a [16]float32) [16]float32
TEXT ·m512GetexpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGETEXPPS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskGetexpPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskGetexpPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512GetexpRoundPd(a [8]float64, rounding int) [8]float64
TEXT ·m512GetexpRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VGETEXPPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGetexpRoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64
TEXT ·m512MaskGetexpRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512GetexpRoundPs(a [16]float32, rounding int) [16]float32
TEXT ·m512GetexpRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGETEXPPS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGetexpRoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32
TEXT ·m512MaskGetexpRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512GetmantPd(a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float64
TEXT ·m512GetmantPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskGetmantPd(src [8]float64, k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float64
TEXT ·m512MaskGetmantPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512GetmantPs(a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [16]float32
TEXT ·m512GetmantPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskGetmantPs(src [16]float32, k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [16]float32
TEXT ·m512MaskGetmantPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512GetmantRoundPd(a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [8]float64
TEXT ·m512GetmantRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskGetmantRoundPd(src [8]float64, k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [8]float64
TEXT ·m512MaskGetmantRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512GetmantRoundPs(a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [16]float32
TEXT ·m512GetmantRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskGetmantRoundPs(src [16]float32, k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [16]float32
TEXT ·m512MaskGetmantRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z5, ret+0(FP)
	RET

// func m512GmaxPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512GmaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VGMAXPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGmaxPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskGmaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512GmaxPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512GmaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGMAXPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGmaxPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskGmaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512GmaxabsPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512GmaxabsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGMAXABSPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGmaxabsPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskGmaxabsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512GminPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512GminPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VGMINPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGminPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskGminPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512GminPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512GminPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGMINPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskGminPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskGminPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512I32extgatherEpi32(index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte
TEXT ·m512I32extgatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI32extgatherEpi32(src [64]byte, k uint16, index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte
TEXT ·m512MaskI32extgatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I32extgatherPs(index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32
TEXT ·m512I32extgatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI32extgatherPs(src [16]float32, k uint16, index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32
TEXT ·m512MaskI32extgatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I32extscatterEpi32(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 
TEXT ·m512I32extscatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32extscatterEpi32(mv uintptr, k uint16, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 
TEXT ·m512MaskI32extscatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32extscatterPs(mv uintptr, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int) 
TEXT ·m512I32extscatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32extscatterPs(mv uintptr, k uint16, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int, hint int) 
TEXT ·m512MaskI32extscatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32gatherEpi32(vindex [64]byte, base_addr uintptr, scale int) [64]byte
TEXT ·m512I32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI32gatherEpi32(src [64]byte, k uint16, vindex [64]byte, base_addr uintptr, scale int) [64]byte
TEXT ·m512MaskI32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512I32gatherPs(vindex [64]byte, base_addr uintptr, scale int) [16]float32
TEXT ·m512I32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI32gatherPs(src [16]float32, k uint16, vindex [64]byte, base_addr uintptr, scale int) [16]float32
TEXT ·m512MaskI32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512I32loextgatherEpi64(index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte
TEXT ·m512I32loextgatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI32loextgatherEpi64(src [64]byte, k uint8, index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte
TEXT ·m512MaskI32loextgatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I32loextgatherPd(index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64
TEXT ·m512I32loextgatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI32loextgatherPd(src [8]float64, k uint8, index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64
TEXT ·m512MaskI32loextgatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I32loextscatterEpi64(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 
TEXT ·m512I32loextscatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32loextscatterEpi64(mv uintptr, k uint8, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 
TEXT ·m512MaskI32loextscatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32loextscatterPd(mv uintptr, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 
TEXT ·m512I32loextscatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32loextscatterPd(mv uintptr, k uint8, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 
TEXT ·m512MaskI32loextscatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32logatherEpi64(index [64]byte, mv uintptr, scale int) [64]byte
TEXT ·m512I32logatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI32logatherEpi64(src [64]byte, k uint8, index [64]byte, mv uintptr, scale int) [64]byte
TEXT ·m512MaskI32logatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512I32logatherPd(index [64]byte, mv uintptr, scale int) [8]float64
TEXT ·m512I32logatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI32logatherPd(src [8]float64, k uint8, index [64]byte, mv uintptr, scale int) [8]float64
TEXT ·m512MaskI32logatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512I32loscatterEpi64(mv uintptr, index [64]byte, v1 [64]byte, scale int) 
TEXT ·m512I32loscatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32loscatterEpi64(mv uintptr, k uint8, index [64]byte, v1 [64]byte, scale int) 
TEXT ·m512MaskI32loscatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32loscatterPd(mv uintptr, index [64]byte, v1 [8]float64, scale int) 
TEXT ·m512I32loscatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32loscatterPd(mv uintptr, k uint8, index [64]byte, v1 [8]float64, scale int) 
TEXT ·m512MaskI32loscatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32scatterEpi32(base_addr uintptr, vindex [64]byte, a [64]byte, scale int) 
TEXT ·m512I32scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32scatterEpi32(base_addr uintptr, k uint16, vindex [64]byte, a [64]byte, scale int) 
TEXT ·m512MaskI32scatterEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I32scatterPs(base_addr uintptr, vindex [64]byte, a [16]float32, scale int) 
TEXT ·m512I32scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI32scatterPs(base_addr uintptr, k uint16, vindex [64]byte, a [16]float32, scale int) 
TEXT ·m512MaskI32scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I64extgatherEpi32lo(index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte
TEXT ·m512I64extgatherEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI64extgatherEpi32lo(src [64]byte, k uint8, index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte
TEXT ·m512MaskI64extgatherEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I64extgatherEpi64(index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte
TEXT ·m512I64extgatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI64extgatherEpi64(src [64]byte, k uint8, index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte
TEXT ·m512MaskI64extgatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I64extgatherPd(index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64
TEXT ·m512I64extgatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI64extgatherPd(src [8]float64, k uint8, index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64
TEXT ·m512MaskI64extgatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I64extgatherPslo(index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32
TEXT ·m512I64extgatherPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MaskI64extgatherPslo(src [16]float32, k uint8, index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32
TEXT ·m512MaskI64extgatherPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z6, ret+0(FP)
	RET

// func m512I64extscatterEpi32lo(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 
TEXT ·m512I64extscatterEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64extscatterEpi32lo(mv uintptr, k uint8, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 
TEXT ·m512MaskI64extscatterEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I64extscatterEpi64(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 
TEXT ·m512I64extscatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64extscatterEpi64(mv uintptr, k uint8, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 
TEXT ·m512MaskI64extscatterEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I64extscatterPd(mv uintptr, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 
TEXT ·m512I64extscatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64extscatterPd(mv uintptr, k uint8, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 
TEXT ·m512MaskI64extscatterPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I64extscatterPslo(mv uintptr, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int, hint int) 
TEXT ·m512I64extscatterPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64extscatterPslo(mv uintptr, k uint8, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int, hint int) 
TEXT ·m512MaskI64extscatterPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I64gatherEpi32lo(index [64]byte, mv uintptr, scale int) [64]byte
TEXT ·m512I64gatherEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI64gatherEpi32lo(src [64]byte, k uint8, index [64]byte, mv uintptr, scale int) [64]byte
TEXT ·m512MaskI64gatherEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512I64gatherPslo(index [64]byte, mv uintptr, scale int) [16]float32
TEXT ·m512I64gatherPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskI64gatherPslo(src [16]float32, k uint8, index [64]byte, mv uintptr, scale int) [16]float32
TEXT ·m512MaskI64gatherPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512I64scatterEpi32lo(mv uintptr, index [64]byte, v1 [64]byte, scale int) 
TEXT ·m512I64scatterEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64scatterEpi32lo(mv uintptr, k uint8, index [64]byte, v1 [64]byte, scale int) 
TEXT ·m512MaskI64scatterEpi32lo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512I64scatterPslo(mv uintptr, index [64]byte, v [16]float32, scale int) 
TEXT ·m512I64scatterPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskI64scatterPslo(mv uintptr, k uint8, index [64]byte, v1 [16]float32, scale int) 
TEXT ·m512MaskI64scatterPslo(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512Int2mask(mask int) uint16
TEXT ·m512Int2mask(SB),7,$0
	MOVQ mask+0(FP),R8

	// TODO: Code missing
	// Could be:
	// KMOV R8

	MOVW $0, ret+8(FP)
	RET

// func m512Kand(a uint16, b uint16) uint16
TEXT ·m512Kand(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KAND R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kandn(a uint16, b uint16) uint16
TEXT ·m512Kandn(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KANDN R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kandnr(k1 uint16, k2 uint16) uint16
TEXT ·m512Kandnr(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KANDNR R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kconcathi64(k1 uint16, k2 uint16) int64
TEXT ·m512Kconcathi64(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KCONCATH R8, R9

	MOVQ $0, ret+8(FP)
	RET

// func m512Kconcatlo64(k1 uint16, k2 uint16) int64
TEXT ·m512Kconcatlo64(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KCONCATL R8, R9

	MOVQ $0, ret+8(FP)
	RET

// func m512Kextract64(a int64, b int) uint16
TEXT ·m512Kextract64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ b+8(FP),R9

	// TODO: Code missing
	// Could be:
	// KEXTRACT R8, R9

	MOVW $0, ret+16(FP)
	RET

// func m512Kmerge2l1h(k1 uint16, k2 uint16) uint16
TEXT ·m512Kmerge2l1h(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KMERGE2L1H R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kmerge2l1l(k1 uint16, k2 uint16) uint16
TEXT ·m512Kmerge2l1l(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KMERGE2L1L R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kmov(a uint16) uint16
TEXT ·m512Kmov(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// KMOV R8

	MOVW $0, ret+4(FP)
	RET

// func m512Kmovlhb(k1 uint16, k2 uint16) uint16
TEXT ·m512Kmovlhb(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KMERGE2L1L R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Knot(a uint16) uint16
TEXT ·m512Knot(SB),7,$0
	MOVW a+0(FP),R8

	// TODO: Code missing
	// Could be:
	// KNOT R8

	MOVW $0, ret+4(FP)
	RET

// func m512Kor(a uint16, b uint16) uint16
TEXT ·m512Kor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KOR R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kortestc(k1 uint16, k2 uint16) int
TEXT ·m512Kortestc(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KORTEST R8, R9

	MOVQ $0, ret+8(FP)
	RET

// func m512Kortestz(k1 uint16, k2 uint16) int
TEXT ·m512Kortestz(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KORTEST R8, R9

	MOVQ $0, ret+8(FP)
	RET

// func m512Kswapb(k1 uint16, k2 uint16) uint16
TEXT ·m512Kswapb(SB),7,$0
	MOVW k1+0(FP),R8
	MOVW k2+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KMERGE2L1H R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kxnor(a uint16, b uint16) uint16
TEXT ·m512Kxnor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KXNOR R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512Kxor(a uint16, b uint16) uint16
TEXT ·m512Kxor(SB),7,$0
	MOVW a+0(FP),R8
	MOVW b+4(FP),R9

	// TODO: Code missing
	// Could be:
	// KXOR R8, R9

	MOVW $0, ret+8(FP)
	RET

// func m512LoadEpi32(mem_addr uintptr) [64]byte
TEXT ·m512LoadEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8

	MOV Z0, ret+8(FP)
	RET

// func m512MaskLoadEpi32(src [64]byte, k uint16, mem_addr uintptr) [64]byte
TEXT ·m512MaskLoadEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadEpi64(mem_addr uintptr) [64]byte
TEXT ·m512LoadEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8

	MOV Z0, ret+8(FP)
	RET

// func m512MaskLoadEpi64(src [64]byte, k uint8, mem_addr uintptr) [64]byte
TEXT ·m512MaskLoadEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadPd(mem_addr uintptr) [8]float64
TEXT ·m512LoadPd(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVAPD R8

	MOV Z0, ret+8(FP)
	RET

// func m512MaskLoadPd(src [8]float64, k uint8, mem_addr uintptr) [8]float64
TEXT ·m512MaskLoadPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadPs(mem_addr uintptr) [16]float32
TEXT ·m512LoadPs(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVAPS R8

	MOV Z0, ret+8(FP)
	RET

// func m512MaskLoadPs(src [16]float32, k uint16, mem_addr uintptr) [16]float32
TEXT ·m512MaskLoadPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadSi512(mem_addr uintptr) [64]byte
TEXT ·m512LoadSi512(SB),7,$0
	MOVQ mem_addr+0(FP),R8

	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8

	MOV Z0, ret+8(FP)
	RET

// func m512LoadunpackhiEpi32(src [64]byte, mt uintptr) [64]byte
TEXT ·m512LoadunpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKHD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackhiEpi32(src [64]byte, k uint16, mt uintptr) [64]byte
TEXT ·m512MaskLoadunpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackhiEpi64(src [64]byte, mt uintptr) [64]byte
TEXT ·m512LoadunpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKHQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackhiEpi64(src [64]byte, k uint8, mt uintptr) [64]byte
TEXT ·m512MaskLoadunpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackhiPd(src [8]float64, mt uintptr) [8]float64
TEXT ·m512LoadunpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKHPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackhiPd(src [8]float64, k uint8, mt uintptr) [8]float64
TEXT ·m512MaskLoadunpackhiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackhiPs(src [16]float32, mt uintptr) [16]float32
TEXT ·m512LoadunpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKHPS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackhiPs(src [16]float32, k uint16, mt uintptr) [16]float32
TEXT ·m512MaskLoadunpackhiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackloEpi32(src [64]byte, mt uintptr) [64]byte
TEXT ·m512LoadunpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKLD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackloEpi32(src [64]byte, k uint16, mt uintptr) [64]byte
TEXT ·m512MaskLoadunpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackloEpi64(src [64]byte, mt uintptr) [64]byte
TEXT ·m512LoadunpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKLQ Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackloEpi64(src [64]byte, k uint8, mt uintptr) [64]byte
TEXT ·m512MaskLoadunpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackloPd(src [8]float64, mt uintptr) [8]float64
TEXT ·m512LoadunpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKLPD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackloPd(src [8]float64, k uint8, mt uintptr) [8]float64
TEXT ·m512MaskLoadunpackloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512LoadunpackloPs(src [16]float32, mt uintptr) [16]float32
TEXT ·m512LoadunpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VLOADUNPACKLPS Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskLoadunpackloPs(src [16]float32, k uint16, mt uintptr) [16]float32
TEXT ·m512MaskLoadunpackloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log2Ps(a [16]float32) [16]float32
TEXT ·m512Log2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VLOG2PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog2Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskLog2Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Log2ae23Ps(a [16]float32) [16]float32
TEXT ·m512Log2ae23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VLOG2PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskLog2ae23Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskLog2ae23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Mask2int(k1 uint16) int
TEXT ·m512Mask2int(SB),7,$0
	MOVW k1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// KMOV R8

	MOVQ $0, ret+4(FP)
	RET

// func m512MaskMaxEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaxEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMAXSD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMaxEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaxEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMAXUD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMaxabsPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskMaxabsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaxabsPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaxabsPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VGMAXABSPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMinEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MinEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMINSD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMinEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MinEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMINUD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMovEpi32(src [64]byte, k uint16, a [64]byte) [64]byte
TEXT ·m512MaskMovEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMovEpi64(src [64]byte, k uint8, a [64]byte) [64]byte
TEXT ·m512MaskMovEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMovPd(src [8]float64, k uint8, a [8]float64) [8]float64
TEXT ·m512MaskMovPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMovPs(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskMovPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMulPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskMulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MulPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512MulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMULPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMulPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskMulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MulPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512MulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMULPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMulRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskMulRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MulRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MulRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMulRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskMulRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512MulRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MulRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskMulhiEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MulhiEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMULHD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMulhiEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulhiEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MulhiEpu32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulhiEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMULHUD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskMulloEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskMulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MulloEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512MulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPMULLD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskOrEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskOrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512OrEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512OrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPORD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskOrEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskOrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512OrEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512OrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPORQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512OrSi512(a [64]byte, b [64]byte) [64]byte
TEXT ·m512OrSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPORD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskPackstorehiEpi32(mt uintptr, k uint16, v1 [64]byte) 
TEXT ·m512MaskPackstorehiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PackstorehiEpi32(mt uintptr, v1 [64]byte) 
TEXT ·m512PackstorehiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPACKSTOREHD R8, Z1

	RET

// func m512MaskPackstorehiEpi64(mt uintptr, k uint8, v1 [64]byte) 
TEXT ·m512MaskPackstorehiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PackstorehiEpi64(mt uintptr, v1 [64]byte) 
TEXT ·m512PackstorehiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPACKSTOREHQ R8, Z1

	RET

// func m512MaskPackstorehiPd(mt uintptr, k uint8, v1 [8]float64) 
TEXT ·m512MaskPackstorehiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512PackstorehiPd(mt uintptr, v1 [8]float64) 
TEXT ·m512PackstorehiPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VPACKSTOREHPD R8, Z1

	RET

// func m512MaskPackstorehiPs(mt uintptr, k uint16, v1 [16]float32) 
TEXT ·m512MaskPackstorehiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512PackstorehiPs(mt uintptr, v1 [16]float32) 
TEXT ·m512PackstorehiPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VPACKSTOREHPS R8, Z1

	RET

// func m512MaskPackstoreloEpi32(mt uintptr, k uint16, v1 [64]byte) 
TEXT ·m512MaskPackstoreloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PackstoreloEpi32(mt uintptr, v1 [64]byte) 
TEXT ·m512PackstoreloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPACKSTORELD R8, Z1

	RET

// func m512MaskPackstoreloEpi64(mt uintptr, k uint8, v1 [64]byte) 
TEXT ·m512MaskPackstoreloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PackstoreloEpi64(mt uintptr, v1 [64]byte) 
TEXT ·m512PackstoreloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPACKSTORELQ R8, Z1

	RET

// func m512MaskPackstoreloPd(mt uintptr, k uint8, v1 [8]float64) 
TEXT ·m512MaskPackstoreloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512PackstoreloPd(mt uintptr, v1 [8]float64) 
TEXT ·m512PackstoreloPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VPACKSTORELPD R8, Z1

	RET

// func m512MaskPackstoreloPs(mt uintptr, k uint16, v1 [16]float32) 
TEXT ·m512MaskPackstoreloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512PackstoreloPs(mt uintptr, v1 [16]float32) 
TEXT ·m512PackstoreloPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VPACKSTORELPS R8, Z1

	RET

// func m512MaskPermute4f128Epi32(src [64]byte, k uint16, a [64]byte, imm8 MMPERMENUM) [64]byte
TEXT ·m512MaskPermute4f128Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permute4f128Epi32(a [64]byte, imm8 MMPERMENUM) [64]byte
TEXT ·m512Permute4f128Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMF32X4 Z0, 

	MOV Z1, ret+0(FP)
	RET

// func m512MaskPermute4f128Ps(src [16]float32, k uint16, a [16]float32, imm8 MMPERMENUM) [16]float32
TEXT ·m512MaskPermute4f128Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512Permute4f128Ps(a [16]float32, imm8 MMPERMENUM) [16]float32
TEXT ·m512Permute4f128Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VPERMF32X4 Z0, 

	MOV Z1, ret+0(FP)
	RET

// func m512MaskPermutevarEpi32(src [64]byte, k uint16, idx [64]byte, a [64]byte) [64]byte
TEXT ·m512MaskPermutevarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512PermutevarEpi32(idx [64]byte, a [64]byte) [64]byte
TEXT ·m512PermutevarEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPERMD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func prefetch(p byte, i int) 
TEXT ·prefetch(SB),7,$0
	MOVB p+0(FP),R8
	MOVQ i+4(FP),R9

	// TODO: Code missing
	// Could be:
	// VPREFETCH0, VPREFETCH1, VPREFETCH2, VPREFETCHNTA, VPREFETCHE0, VPREFETCHE1, VPREFETCHE2, VPREFETCHENTA R8, R9

	RET

// func m512MaskPrefetchI32extgatherPs(index [64]byte, k uint16, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) 
TEXT ·m512MaskPrefetchI32extgatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PrefetchI32extgatherPs(index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) 
TEXT ·m512PrefetchI32extgatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskPrefetchI32extscatterPs(mv uintptr, k uint16, index [64]byte, conv MMUPCONVPSENUM, scale int, hint int) 
TEXT ·m512MaskPrefetchI32extscatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PrefetchI32extscatterPs(mv uintptr, index [64]byte, conv MMUPCONVPSENUM, scale int, hint int) 
TEXT ·m512PrefetchI32extscatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskPrefetchI32gatherPs(vindex [64]byte, mask uint16, base_addr uintptr, scale int, hint int) 
TEXT ·m512MaskPrefetchI32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PrefetchI32gatherPs(index [64]byte, mv uintptr, scale int, hint int) 
TEXT ·m512PrefetchI32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskPrefetchI32scatterPs(mv uintptr, k uint16, index [64]byte, scale int, hint int) 
TEXT ·m512MaskPrefetchI32scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512PrefetchI32scatterPs(mv uintptr, index [64]byte, scale int, hint int) 
TEXT ·m512PrefetchI32scatterPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512MaskRcp23Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRcp23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Rcp23Ps(a [16]float32) [16]float32
TEXT ·m512Rcp23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRCP23PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskReduceAddEpi32(k uint16, a [64]byte) int
TEXT ·m512MaskReduceAddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceAddEpi32(a [64]byte) int
TEXT ·m512ReduceAddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceAddEpi64(k uint8, a [64]byte) int64
TEXT ·m512MaskReduceAddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceAddEpi64(a [64]byte) int64
TEXT ·m512ReduceAddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceAddPd(k uint8, a [8]float64) float64
TEXT ·m512MaskReduceAddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceAddPd(a [8]float64) float64
TEXT ·m512ReduceAddPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceAddPs(k uint16, a [16]float32) float32
TEXT ·m512MaskReduceAddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceAddPs(a [16]float32) float32
TEXT ·m512ReduceAddPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceAndEpi32(k uint16, a [64]byte) int
TEXT ·m512MaskReduceAndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceAndEpi32(a [64]byte) int
TEXT ·m512ReduceAndEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceAndEpi64(k uint8, a [64]byte) int64
TEXT ·m512MaskReduceAndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceAndEpi64(a [64]byte) int64
TEXT ·m512ReduceAndEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceGmaxPd(k uint8, a [8]float64) float64
TEXT ·m512MaskReduceGmaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceGmaxPd(a [8]float64) float64
TEXT ·m512ReduceGmaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceGmaxPs(k uint16, a [16]float32) float32
TEXT ·m512MaskReduceGmaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceGmaxPs(a [16]float32) float32
TEXT ·m512ReduceGmaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceGminPd(k uint8, a [8]float64) float64
TEXT ·m512MaskReduceGminPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceGminPd(a [8]float64) float64
TEXT ·m512ReduceGminPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceGminPs(k uint16, a [16]float32) float32
TEXT ·m512MaskReduceGminPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceGminPs(a [16]float32) float32
TEXT ·m512ReduceGminPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceMaxEpi32(k uint16, a [64]byte) int
TEXT ·m512MaskReduceMaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMaxEpi32(a [64]byte) int
TEXT ·m512ReduceMaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMaxEpi64(k uint8, a [64]byte) int64
TEXT ·m512MaskReduceMaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMaxEpi64(a [64]byte) int64
TEXT ·m512ReduceMaxEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMaxEpu32(k uint16, a [64]byte) uint32
TEXT ·m512MaskReduceMaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceMaxEpu32(a [64]byte) uint32
TEXT ·m512ReduceMaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceMaxEpu64(k uint8, a [64]byte) uint64
TEXT ·m512MaskReduceMaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMaxEpu64(a [64]byte) uint64
TEXT ·m512ReduceMaxEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMaxPd(k uint8, a [8]float64) float64
TEXT ·m512MaskReduceMaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMaxPd(a [8]float64) float64
TEXT ·m512ReduceMaxPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMaxPs(k uint16, a [16]float32) float32
TEXT ·m512MaskReduceMaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceMaxPs(a [16]float32) float32
TEXT ·m512ReduceMaxPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceMinEpi32(k uint16, a [64]byte) int
TEXT ·m512MaskReduceMinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMinEpi32(a [64]byte) int
TEXT ·m512ReduceMinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMinEpi64(k uint8, a [64]byte) int64
TEXT ·m512MaskReduceMinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMinEpi64(a [64]byte) int64
TEXT ·m512ReduceMinEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMinEpu32(k uint16, a [64]byte) uint32
TEXT ·m512MaskReduceMinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceMinEpu32(a [64]byte) uint32
TEXT ·m512ReduceMinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceMinEpu64(k uint8, a [64]byte) uint64
TEXT ·m512MaskReduceMinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMinEpu64(a [64]byte) uint64
TEXT ·m512ReduceMinEpu64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMinPd(k uint8, a [8]float64) float64
TEXT ·m512MaskReduceMinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMinPd(a [8]float64) float64
TEXT ·m512ReduceMinPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMinPs(k uint16, a [16]float32) float32
TEXT ·m512MaskReduceMinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceMinPs(a [16]float32) float32
TEXT ·m512ReduceMinPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceMulEpi32(k uint16, a [64]byte) int
TEXT ·m512MaskReduceMulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMulEpi32(a [64]byte) int
TEXT ·m512ReduceMulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMulEpi64(k uint8, a [64]byte) int64
TEXT ·m512MaskReduceMulEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMulEpi64(a [64]byte) int64
TEXT ·m512ReduceMulEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMulPd(k uint8, a [8]float64) float64
TEXT ·m512MaskReduceMulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceMulPd(a [8]float64) float64
TEXT ·m512ReduceMulPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceMulPs(k uint16, a [16]float32) float32
TEXT ·m512MaskReduceMulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512ReduceMulPs(a [16]float32) float32
TEXT ·m512ReduceMulPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOVL $0, ret+0(FP)
	RET

// func m512MaskReduceOrEpi32(k uint16, a [64]byte) int
TEXT ·m512MaskReduceOrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceOrEpi32(a [64]byte) int
TEXT ·m512ReduceOrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskReduceOrEpi64(k uint8, a [64]byte) int64
TEXT ·m512MaskReduceOrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512ReduceOrEpi64(a [64]byte) int64
TEXT ·m512ReduceOrEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m512MaskRoundPs(src [16]float32, k uint16, a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512MaskRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512RoundPs(a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512RoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskRoundfxpntAdjustPd(src [8]float64, k uint8, a [8]float64, rounding int, expadj MMEXPADJENUM) [8]float64
TEXT ·m512MaskRoundfxpntAdjustPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512RoundfxpntAdjustPd(a [8]float64, rounding int, expadj MMEXPADJENUM) [8]float64
TEXT ·m512RoundfxpntAdjustPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskRoundfxpntAdjustPs(src [16]float32, k uint16, a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512MaskRoundfxpntAdjustPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512RoundfxpntAdjustPs(a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32
TEXT ·m512RoundfxpntAdjustPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskRsqrt23Ps(src [16]float32, k uint16, a [16]float32) [16]float32
TEXT ·m512MaskRsqrt23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512Rsqrt23Ps(a [16]float32) [16]float32
TEXT ·m512Rsqrt23Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VRSQRT23PS Z0

	MOV Z0, ret+0(FP)
	RET

// func m512MaskSbbEpi32(v2 [64]byte, k1 uint16, k2 uint16, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512MaskSbbEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SbbEpi32(v2 [64]byte, k uint16, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512SbbEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskSbbrEpi32(v2 [64]byte, k1 uint16, k2 uint16, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512MaskSbbrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SbbrEpi32(v2 [64]byte, k uint16, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512SbbrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512MaskScalePs(src [16]float32, k uint16, a [16]float32, b [64]byte) [16]float32
TEXT ·m512MaskScalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ScalePs(a [16]float32, b [64]byte) [16]float32
TEXT ·m512ScalePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSCALEPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskScaleRoundPs(src [16]float32, k uint16, a [16]float32, b [64]byte, rounding int) [16]float32
TEXT ·m512MaskScaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512ScaleRoundPs(a [16]float32, b [64]byte, rounding int) [16]float32
TEXT ·m512ScaleRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskShuffleEpi32(src [64]byte, k uint16, a [64]byte, imm8 MMPERMENUM) [64]byte
TEXT ·m512MaskShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512ShuffleEpi32(a [64]byte, imm8 MMPERMENUM) [64]byte
TEXT ·m512ShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSHUFD Z0, 

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSlliEpi32(src [64]byte, k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSlliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SlliEpi32(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SlliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSLLD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSllvEpi32(src [64]byte, k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSllvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SllvEpi32(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SllvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSLLVD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func spflt32(r1 uint32) 
TEXT ·spflt32(SB),7,$0
	MOVL r1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// SPFLT R8

	RET

// func spflt64(r1 uint64) 
TEXT ·spflt64(SB),7,$0
	MOVQ r1+0(FP),R8

	// TODO: Code missing
	// Could be:
	// SPFLT R8

	RET

// func m512MaskSraiEpi32(src [64]byte, k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSraiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SraiEpi32(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SraiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRAD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSravEpi32(src [64]byte, k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSravEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SravEpi32(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SravEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRAVD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSrliEpi32(src [64]byte, k uint16, a [64]byte, imm8 uint32) [64]byte
TEXT ·m512MaskSrliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SrliEpi32(a [64]byte, imm8 uint32) [64]byte
TEXT ·m512SrliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRLD Z0, R9

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSrlvEpi32(src [64]byte, k uint16, a [64]byte, count [64]byte) [64]byte
TEXT ·m512MaskSrlvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SrlvEpi32(a [64]byte, count [64]byte) [64]byte
TEXT ·m512SrlvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSRLVD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskStoreEpi32(mem_addr uintptr, k uint16, a [64]byte) 
TEXT ·m512MaskStoreEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512StoreEpi32(mem_addr uintptr, a [64]byte) 
TEXT ·m512StoreEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, Z1

	RET

// func m512MaskStoreEpi64(mem_addr uintptr, k uint8, a [64]byte) 
TEXT ·m512MaskStoreEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	RET

// func m512StoreEpi64(mem_addr uintptr, a [64]byte) 
TEXT ·m512StoreEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVDQA64 R8, Z1

	RET

// func m512MaskStorePd(mem_addr uintptr, k uint8, a [8]float64) 
TEXT ·m512MaskStorePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	RET

// func m512StorePd(mem_addr uintptr, a [8]float64) 
TEXT ·m512StorePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVAPD R8, Z1

	RET

// func m512MaskStorePs(mem_addr uintptr, k uint16, a [16]float32) 
TEXT ·m512MaskStorePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	RET

// func m512StorePs(mem_addr uintptr, a [16]float32) 
TEXT ·m512StorePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVAPS R8, Z1

	RET

// func m512StoreSi512(mem_addr uintptr, a [64]byte) 
TEXT ·m512StoreSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VMOVDQA32 R8, Z1

	RET

// func m512StorenrPd(mt uintptr, v [8]float64) 
TEXT ·m512StorenrPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVNRAPD R8, Z1

	RET

// func m512StorenrPs(mt uintptr, v [16]float32) 
TEXT ·m512StorenrPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVNRAPS R8, Z1

	RET

// func m512StorenrngoPd(mt uintptr, v [8]float64) 
TEXT ·m512StorenrngoPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VMOVNRNGOAPD R8, Z1

	RET

// func m512StorenrngoPs(mt uintptr, v [16]float32) 
TEXT ·m512StorenrngoPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VMOVNRNGOAPS R8, Z1

	RET

// func m512MaskSubEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskSubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SubEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512SubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSUBD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSubPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64
TEXT ·m512MaskSubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SubPd(a [8]float64, b [8]float64) [8]float64
TEXT ·m512SubPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VSUBPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSubPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32
TEXT ·m512MaskSubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SubPs(a [16]float32, b [16]float32) [16]float32
TEXT ·m512SubPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSUBPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSubRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512MaskSubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SubRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64
TEXT ·m512SubRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSubRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512MaskSubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SubRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32
TEXT ·m512SubRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSubrEpi32(src [64]byte, k uint16, v2 [64]byte, v3 [64]byte) [64]byte
TEXT ·m512MaskSubrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SubrEpi32(v2 [64]byte, v3 [64]byte) [64]byte
TEXT ·m512SubrEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPSUBRD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSubrPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64) [8]float64
TEXT ·m512MaskSubrPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SubrPd(v2 [8]float64, v3 [8]float64) [8]float64
TEXT ·m512SubrPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing
	// Could be:
	// VSUBRPD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSubrPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32) [16]float32
TEXT ·m512MaskSubrPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SubrPs(v2 [16]float32, v3 [16]float32) [16]float32
TEXT ·m512SubrPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing
	// Could be:
	// VSUBRPS Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSubrRoundPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64, rounding int) [8]float64
TEXT ·m512MaskSubrRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SubrRoundPd(v2 [8]float64, v3 [8]float64, rounding int) [8]float64
TEXT ·m512SubrRoundPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSubrRoundPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, rounding int) [16]float32
TEXT ·m512MaskSubrRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SubrRoundPs(v2 [16]float32, v3 [16]float32, rounding int) [16]float32
TEXT ·m512SubrRoundPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSubrsetbEpi32(v2 [64]byte, k uint16, k_old uint16, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512MaskSubrsetbEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SubrsetbEpi32(v2 [64]byte, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512SubrsetbEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSubsetbEpi32(v2 [64]byte, k uint16, k_old uint16, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512MaskSubsetbEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z4, ret+0(FP)
	RET

// func m512SubsetbEpi32(v2 [64]byte, v3 [64]byte, borrow uint16) [64]byte
TEXT ·m512SubsetbEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z2, ret+0(FP)
	RET

// func m512MaskSwizzleEpi32(src [64]byte, k uint16, v [64]byte, s MMSWIZZLEENUM) [64]byte
TEXT ·m512MaskSwizzleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SwizzleEpi32(v [64]byte, s MMSWIZZLEENUM) [64]byte
TEXT ·m512SwizzleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSwizzleEpi64(src [64]byte, k uint8, v [64]byte, s MMSWIZZLEENUM) [64]byte
TEXT ·m512MaskSwizzleEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SwizzleEpi64(v [64]byte, s MMSWIZZLEENUM) [64]byte
TEXT ·m512SwizzleEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSwizzlePd(src [8]float64, k uint8, v [8]float64, s MMSWIZZLEENUM) [8]float64
TEXT ·m512MaskSwizzlePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SwizzlePd(v [8]float64, s MMSWIZZLEENUM) [8]float64
TEXT ·m512SwizzlePd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512d
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskSwizzlePs(src [16]float32, k uint16, v [16]float32, s MMSWIZZLEENUM) [16]float32
TEXT ·m512MaskSwizzlePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512SwizzlePs(v [16]float32, s MMSWIZZLEENUM) [16]float32
TEXT ·m512SwizzlePs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512
	// TODO: Code missing

	MOV Z1, ret+0(FP)
	RET

// func m512MaskTestEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16
TEXT ·m512MaskTestEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOVW $0, ret+0(FP)
	RET

// func m512TestEpi32Mask(a [64]byte, b [64]byte) uint16
TEXT ·m512TestEpi32Mask(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPTESTMD Z0, Z1

	MOVW $0, ret+0(FP)
	RET

// func tzcnti32(a int, x uint32) int
TEXT ·tzcnti32(SB),7,$0
	MOVQ a+0(FP),R8
	MOVL x+8(FP),R9

	// TODO: Code missing
	// Could be:
	// TZCNTI R8, R9

	MOVQ $0, ret+12(FP)
	RET

// func tzcnti64(a int64, x uint64) int64
TEXT ·tzcnti64(SB),7,$0
	MOVQ a+0(FP),R8
	MOVQ x+8(FP),R9

	// TODO: Code missing
	// Could be:
	// TZCNTI R8, R9

	MOVQ $0, ret+16(FP)
	RET

// func m512MaskXorEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskXorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512XorEpi32(a [64]byte, b [64]byte) [64]byte
TEXT ·m512XorEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPXORD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512MaskXorEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte
TEXT ·m512MaskXorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing

	MOV Z3, ret+0(FP)
	RET

// func m512XorEpi64(a [64]byte, b [64]byte) [64]byte
TEXT ·m512XorEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPXORQ Z0, Z1

	MOV Z1, ret+0(FP)
	RET

// func m512XorSi512(a [64]byte, b [64]byte) [64]byte
TEXT ·m512XorSi512(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M512i
	// TODO: Code missing
	// Could be:
	// VPXORD Z0, Z1

	MOV Z1, ret+0(FP)
	RET

