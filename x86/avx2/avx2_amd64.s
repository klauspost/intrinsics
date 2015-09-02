// func m256AbsEpi16(a [32]byte) [32]byte
TEXT ·m256AbsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AbsEpi32(a [32]byte) [32]byte
TEXT ·m256AbsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AbsEpi8(a [32]byte) [32]byte
TEXT ·m256AbsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddsEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddsEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AddsEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AddsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AlignrEpi8(a [32]byte, b [32]byte, count int) [32]byte
TEXT ·m256AlignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AndSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AndSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AndnotSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AndnotSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AvgEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AvgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256AvgEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256AvgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256BlendEpi16(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256BlendEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func blendEpi32(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·blendEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+40(FP)
	RET

// func m256BlendEpi32(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256BlendEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256BlendvEpi8(a [32]byte, b [32]byte, mask [32]byte) [32]byte
TEXT ·m256BlendvEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func broadcastbEpi8(a [16]byte) [16]byte
TEXT ·broadcastbEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func m256BroadcastbEpi8(a [16]byte) [32]byte
TEXT ·m256BroadcastbEpi8(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func broadcastdEpi32(a [16]byte) [16]byte
TEXT ·broadcastdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func m256BroadcastdEpi32(a [16]byte) [32]byte
TEXT ·m256BroadcastdEpi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func broadcastqEpi64(a [16]byte) [16]byte
TEXT ·broadcastqEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func m256BroadcastqEpi64(a [16]byte) [32]byte
TEXT ·m256BroadcastqEpi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func broadcastsdPd(a [2]float64) [2]float64
TEXT ·broadcastsdPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func m256BroadcastsdPd(a [2]float64) [4]float64
TEXT ·m256BroadcastsdPd(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Broadcastsi128Si256(a [16]byte) [32]byte
TEXT ·m256Broadcastsi128Si256(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func broadcastssPs(a [4]float32) [4]float32
TEXT ·broadcastssPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func m256BroadcastssPs(a [4]float32) [8]float32
TEXT ·m256BroadcastssPs(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func broadcastwEpi16(a [16]byte) [16]byte
TEXT ·broadcastwEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOVOU X0, ret+16(FP)
	RET

// func m256BroadcastwEpi16(a [16]byte) [32]byte
TEXT ·m256BroadcastwEpi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256BslliEpi128(a [32]byte, imm8 int) [32]byte
TEXT ·m256BslliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256BsrliEpi128(a [32]byte, imm8 int) [32]byte
TEXT ·m256BsrliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpeqEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpeqEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpeqEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpeqEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpeqEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpeqEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpeqEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpeqEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpgtEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpgtEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpgtEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpgtEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpgtEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpgtEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256CmpgtEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256CmpgtEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Cvtepi16Epi32(a [16]byte) [32]byte
TEXT ·m256Cvtepi16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepi16Epi64(a [16]byte) [32]byte
TEXT ·m256Cvtepi16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepi32Epi64(a [16]byte) [32]byte
TEXT ·m256Cvtepi32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepi8Epi16(a [16]byte) [32]byte
TEXT ·m256Cvtepi8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepi8Epi32(a [16]byte) [32]byte
TEXT ·m256Cvtepi8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepi8Epi64(a [16]byte) [32]byte
TEXT ·m256Cvtepi8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepu16Epi32(a [16]byte) [32]byte
TEXT ·m256Cvtepu16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepu16Epi64(a [16]byte) [32]byte
TEXT ·m256Cvtepu16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepu32Epi64(a [16]byte) [32]byte
TEXT ·m256Cvtepu32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepu8Epi16(a [16]byte) [32]byte
TEXT ·m256Cvtepu8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepu8Epi32(a [16]byte) [32]byte
TEXT ·m256Cvtepu8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Cvtepu8Epi64(a [16]byte) [32]byte
TEXT ·m256Cvtepu8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func m256Extracti128Si256(a [32]byte, imm8 int) [16]byte
TEXT ·m256Extracti128Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256HaddEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256HaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256HaddEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256HaddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256HaddsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256HaddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256HsubEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256HsubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256HsubEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256HsubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256HsubsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256HsubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32gatherEpi32(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i32gatherEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherEpi32(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI32gatherEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func m256I32gatherEpi32(base_addr int, vindex [32]byte, scale int) [32]byte
TEXT ·m256I32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskI32gatherEpi32(src [32]byte, base_addr int, vindex [32]byte, mask [32]byte, scale int) [32]byte
TEXT ·m256MaskI32gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32gatherEpi64(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i32gatherEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherEpi64(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI32gatherEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func m256I32gatherEpi64(base_addr int, vindex [16]byte, scale int) [32]byte
TEXT ·m256I32gatherEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOV Y0, ret+32(FP)
	RET

// func m256MaskI32gatherEpi64(src [32]byte, base_addr int, vindex [16]byte, mask [32]byte, scale int) [32]byte
TEXT ·m256MaskI32gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32gatherPd(base_addr float64, vindex [16]byte, scale int) [2]float64
TEXT ·i32gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI32gatherPd(src [2]float64, base_addr float64, vindex [16]byte, mask [2]float64, scale int) [2]float64
TEXT ·maskI32gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func m256I32gatherPd(base_addr float64, vindex [16]byte, scale int) [4]float64
TEXT ·m256I32gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOV Y0, ret+32(FP)
	RET

// func m256MaskI32gatherPd(src [4]float64, base_addr float64, vindex [16]byte, mask [4]float64, scale int) [4]float64
TEXT ·m256MaskI32gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i32gatherPs(base_addr float32, vindex [16]byte, scale int) [4]float32
TEXT ·i32gatherPs(SB),7,$0
	MOVL base_addr+0(FP),R8
	MOVOU vindex+4(FP),X1
	MOVQ scale+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskI32gatherPs(src [4]float32, base_addr float32, vindex [16]byte, mask [4]float32, scale int) [4]float32
TEXT ·maskI32gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVL base_addr+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVOU mask+36(FP),X3
	MOVQ scale+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func m256I32gatherPs(base_addr float32, vindex [32]byte, scale int) [8]float32
TEXT ·m256I32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskI32gatherPs(src [8]float32, base_addr float32, vindex [32]byte, mask [8]float32, scale int) [8]float32
TEXT ·m256MaskI32gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i64gatherEpi32(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i64gatherEpi32(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherEpi32(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI64gatherEpi32(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func m256I64gatherEpi32(base_addr int, vindex [32]byte, scale int) [16]byte
TEXT ·m256I64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskI64gatherEpi32(src [16]byte, base_addr int, vindex [32]byte, mask [16]byte, scale int) [16]byte
TEXT ·m256MaskI64gatherEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func i64gatherEpi64(base_addr int, vindex [16]byte, scale int) [16]byte
TEXT ·i64gatherEpi64(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherEpi64(src [16]byte, base_addr int, vindex [16]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI64gatherEpi64(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func m256I64gatherEpi64(base_addr int, vindex [32]byte, scale int) [32]byte
TEXT ·m256I64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskI64gatherEpi64(src [32]byte, base_addr int, vindex [32]byte, mask [32]byte, scale int) [32]byte
TEXT ·m256MaskI64gatherEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i64gatherPd(base_addr float64, vindex [16]byte, scale int) [2]float64
TEXT ·i64gatherPd(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func maskI64gatherPd(src [2]float64, base_addr float64, vindex [16]byte, mask [2]float64, scale int) [2]float64
TEXT ·maskI64gatherPd(SB),7,$0
	MOVOU src+0(FP),X0
	MOVQ base_addr+16(FP),R9
	MOVOU vindex+24(FP),X2
	MOVOU mask+40(FP),X3
	MOVQ scale+56(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+64(FP)
	RET

// func m256I64gatherPd(base_addr float64, vindex [32]byte, scale int) [4]float64
TEXT ·m256I64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaskI64gatherPd(src [4]float64, base_addr float64, vindex [32]byte, mask [4]float64, scale int) [4]float64
TEXT ·m256MaskI64gatherPd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func i64gatherPs(base_addr float32, vindex [16]byte, scale int) [4]float32
TEXT ·i64gatherPs(SB),7,$0
	MOVL base_addr+0(FP),R8
	MOVOU vindex+4(FP),X1
	MOVQ scale+20(FP),R10

	//TODO: Code missing

	MOVOU X0, ret+28(FP)
	RET

// func maskI64gatherPs(src [4]float32, base_addr float32, vindex [16]byte, mask [4]float32, scale int) [4]float32
TEXT ·maskI64gatherPs(SB),7,$0
	MOVOU src+0(FP),X0
	MOVL base_addr+16(FP),R9
	MOVOU vindex+20(FP),X2
	MOVOU mask+36(FP),X3
	MOVQ scale+52(FP),R12

	//TODO: Code missing

	MOVOU X0, ret+60(FP)
	RET

// func m256I64gatherPs(base_addr float32, vindex [32]byte, scale int) [4]float32
TEXT ·m256I64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256MaskI64gatherPs(src [4]float32, base_addr float32, vindex [32]byte, mask [4]float32, scale int) [4]float32
TEXT ·m256MaskI64gatherPs(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func m256Inserti128Si256(a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·m256Inserti128Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaddEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaddubsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaddubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskloadEpi32(mem_addr int, mask [16]byte) [16]byte
TEXT ·maskloadEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func m256MaskloadEpi32(mem_addr int, mask [32]byte) [32]byte
TEXT ·m256MaskloadEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskloadEpi64(mem_addr int, mask [16]byte) [16]byte
TEXT ·maskloadEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+24(FP)
	RET

// func m256MaskloadEpi64(mem_addr int, mask [32]byte) [32]byte
TEXT ·m256MaskloadEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskstoreEpi32(mem_addr int, mask [16]byte, a [16]byte) 
TEXT ·maskstoreEpi32(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func m256MaskstoreEpi32(mem_addr int, mask [32]byte, a [32]byte) 
TEXT ·m256MaskstoreEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maskstoreEpi64(mem_addr int64, mask [16]byte, a [16]byte) 
TEXT ·maskstoreEpi64(SB),7,$0
	MOVQ mem_addr+0(FP),R8
	MOVOU mask+8(FP),X1
	MOVOU a+24(FP),X2

	//TODO: Code missing

	RET

// func m256MaskstoreEpi64(mem_addr int64, mask [32]byte, a [32]byte) 
TEXT ·m256MaskstoreEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func m256MaxEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaxEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaxEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaxEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaxEpu32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MaxEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MaxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MinEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MinEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MinEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MinEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MinEpu32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MinEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MinEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MovemaskEpi8(a [32]byte) int
TEXT ·m256MovemaskEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func m256MpsadbwEpu8(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256MpsadbwEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulEpu32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulhiEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulhiEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulhrsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulloEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256MulloEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256MulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256OrSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·m256OrSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256PacksEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256PacksEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256PacksEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256PacksEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256PackusEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256PackusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256PackusEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256PackusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Permute2x128Si256(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·m256Permute2x128Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Permute4x64Epi64(a [32]byte, imm8 int) [32]byte
TEXT ·m256Permute4x64Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Permute4x64Pd(a [4]float64, imm8 int) [4]float64
TEXT ·m256Permute4x64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Permutevar8x32Epi32(a [32]byte, idx [32]byte) [32]byte
TEXT ·m256Permutevar8x32Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256Permutevar8x32Ps(a [8]float32, idx [32]byte) [8]float32
TEXT ·m256Permutevar8x32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SadEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256ShuffleEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·m256ShuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256ShuffleEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256ShuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256ShufflehiEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·m256ShufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256ShuffleloEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·m256ShuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SignEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SignEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SignEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SignEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SignEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SignEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SllEpi16(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SllEpi32(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SllEpi64(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SlliEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·m256SlliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SlliEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·m256SlliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SlliEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·m256SlliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SlliSi256(a [32]byte, imm8 int) [32]byte
TEXT ·m256SlliSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sllvEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sllvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256SllvEpi32(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SllvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sllvEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·sllvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256SllvEpi64(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SllvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SraEpi16(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SraEpi32(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SraiEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·m256SraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SraiEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·m256SraiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sravEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·sravEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256SravEpi32(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SravEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrlEpi16(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SrlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrlEpi32(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SrlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrlEpi64(a [32]byte, count [16]byte) [32]byte
TEXT ·m256SrlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrliEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·m256SrliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrliEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·m256SrliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrliEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·m256SrliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SrliSi256(a [32]byte, imm8 int) [32]byte
TEXT ·m256SrliSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srlvEpi32(a [16]byte, count [16]byte) [16]byte
TEXT ·srlvEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256SrlvEpi32(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SrlvEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srlvEpi64(a [16]byte, count [16]byte) [16]byte
TEXT ·srlvEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU count+16(FP),X1

	//TODO: Code missing

	MOVOU X0, ret+32(FP)
	RET

// func m256SrlvEpi64(a [32]byte, count [32]byte) [32]byte
TEXT ·m256SrlvEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256StreamLoadSi256(mem_addr x86.M256iConst) [32]byte
TEXT ·m256StreamLoadSi256(SB),7,$0
	// Unimplemented. Unknown size of type x86.M256iConst
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubsEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubsEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256SubsEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256SubsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackhiEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackhiEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackhiEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackhiEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackloEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackloEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackloEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256UnpackloEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·m256UnpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func m256XorSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·m256XorSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

