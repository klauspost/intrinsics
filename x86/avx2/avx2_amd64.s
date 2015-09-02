// func absEpi16(a [32]byte) [32]byte
TEXT ·absEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func absEpi32(a [32]byte) [32]byte
TEXT ·absEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func absEpi8(a [32]byte) [32]byte
TEXT ·absEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·addEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·addEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·addEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·addEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·addsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·addsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·addsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func addsEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·addsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func alignrEpi8(a [32]byte, b [32]byte, count int) [32]byte
TEXT ·alignrEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func andSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·andSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func andnotSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·andnotSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func avgEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·avgEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func avgEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·avgEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func blendEpi16(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·blendEpi16(SB),7,$0
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

// func blendEpi321(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·blendEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func blendvEpi8(a [32]byte, b [32]byte, mask [32]byte) [32]byte
TEXT ·blendvEpi8(SB),7,$0
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

// func broadcastbEpi81(a [16]byte) [32]byte
TEXT ·broadcastbEpi81(SB),7,$0
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

// func broadcastdEpi321(a [16]byte) [32]byte
TEXT ·broadcastdEpi321(SB),7,$0
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

// func broadcastqEpi641(a [16]byte) [32]byte
TEXT ·broadcastqEpi641(SB),7,$0
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

// func broadcastsdPd1(a [2]float64) [4]float64
TEXT ·broadcastsdPd1(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func broadcastsi128Si256(a [16]byte) [32]byte
TEXT ·broadcastsi128Si256(SB),7,$0
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

// func broadcastssPs1(a [4]float32) [8]float32
TEXT ·broadcastssPs1(SB),7,$0
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

// func broadcastwEpi161(a [16]byte) [32]byte
TEXT ·broadcastwEpi161(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func bslliEpi128(a [32]byte, imm8 int) [32]byte
TEXT ·bslliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func bsrliEpi128(a [32]byte, imm8 int) [32]byte
TEXT ·bsrliEpi128(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpeqEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpeqEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpeqEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpeqEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpeqEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpeqEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpeqEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpeqEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpgtEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpgtEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpgtEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpgtEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpgtEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpgtEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cmpgtEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·cmpgtEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func cvtepi16Epi32(a [16]byte) [32]byte
TEXT ·cvtepi16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepi16Epi64(a [16]byte) [32]byte
TEXT ·cvtepi16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepi32Epi64(a [16]byte) [32]byte
TEXT ·cvtepi32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepi8Epi16(a [16]byte) [32]byte
TEXT ·cvtepi8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepi8Epi32(a [16]byte) [32]byte
TEXT ·cvtepi8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepi8Epi64(a [16]byte) [32]byte
TEXT ·cvtepi8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepu16Epi32(a [16]byte) [32]byte
TEXT ·cvtepu16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepu16Epi64(a [16]byte) [32]byte
TEXT ·cvtepu16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepu32Epi64(a [16]byte) [32]byte
TEXT ·cvtepu32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepu8Epi16(a [16]byte) [32]byte
TEXT ·cvtepu8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepu8Epi32(a [16]byte) [32]byte
TEXT ·cvtepu8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func cvtepu8Epi64(a [16]byte) [32]byte
TEXT ·cvtepu8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	//TODO: Code missing

	MOV Y0, ret+16(FP)
	RET

// func extracti128Si256(a [32]byte, imm8 int) [16]byte
TEXT ·extracti128Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func haddEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·haddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func haddEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·haddEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func haddsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·haddsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func hsubEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·hsubEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func hsubEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·hsubEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func hsubsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·hsubsEpi16(SB),7,$0
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

// func i32gatherEpi321(base_addr int, vindex [32]byte, scale int) [32]byte
TEXT ·i32gatherEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskI32gatherEpi321(src [32]byte, base_addr int, vindex [32]byte, mask [32]byte, scale int) [32]byte
TEXT ·maskI32gatherEpi321(SB),7,$0
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

// func i32gatherEpi641(base_addr int, vindex [16]byte, scale int) [32]byte
TEXT ·i32gatherEpi641(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOV Y0, ret+32(FP)
	RET

// func maskI32gatherEpi641(src [32]byte, base_addr int, vindex [16]byte, mask [32]byte, scale int) [32]byte
TEXT ·maskI32gatherEpi641(SB),7,$0
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

// func i32gatherPd1(base_addr float64, vindex [16]byte, scale int) [4]float64
TEXT ·i32gatherPd1(SB),7,$0
	MOVQ base_addr+0(FP),R8
	MOVOU vindex+8(FP),X1
	MOVQ scale+24(FP),R10

	//TODO: Code missing

	MOV Y0, ret+32(FP)
	RET

// func maskI32gatherPd1(src [4]float64, base_addr float64, vindex [16]byte, mask [4]float64, scale int) [4]float64
TEXT ·maskI32gatherPd1(SB),7,$0
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

// func i32gatherPs1(base_addr float32, vindex [32]byte, scale int) [8]float32
TEXT ·i32gatherPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskI32gatherPs1(src [8]float32, base_addr float32, vindex [32]byte, mask [8]float32, scale int) [8]float32
TEXT ·maskI32gatherPs1(SB),7,$0
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

// func i64gatherEpi321(base_addr int, vindex [32]byte, scale int) [16]byte
TEXT ·i64gatherEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskI64gatherEpi321(src [16]byte, base_addr int, vindex [32]byte, mask [16]byte, scale int) [16]byte
TEXT ·maskI64gatherEpi321(SB),7,$0
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

// func i64gatherEpi641(base_addr int, vindex [32]byte, scale int) [32]byte
TEXT ·i64gatherEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskI64gatherEpi641(src [32]byte, base_addr int, vindex [32]byte, mask [32]byte, scale int) [32]byte
TEXT ·maskI64gatherEpi641(SB),7,$0
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

// func i64gatherPd1(base_addr float64, vindex [32]byte, scale int) [4]float64
TEXT ·i64gatherPd1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maskI64gatherPd1(src [4]float64, base_addr float64, vindex [32]byte, mask [4]float64, scale int) [4]float64
TEXT ·maskI64gatherPd1(SB),7,$0
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

// func i64gatherPs1(base_addr float32, vindex [32]byte, scale int) [4]float32
TEXT ·i64gatherPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func maskI64gatherPs1(src [4]float32, base_addr float32, vindex [32]byte, mask [4]float32, scale int) [4]float32
TEXT ·maskI64gatherPs1(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVOU X0, ret+0(FP)
	RET

// func inserti128Si256(a [32]byte, b [16]byte, imm8 int) [32]byte
TEXT ·inserti128Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maddEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·maddEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maddubsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·maddubsEpi16(SB),7,$0
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

// func maskloadEpi321(mem_addr int, mask [32]byte) [32]byte
TEXT ·maskloadEpi321(SB),7,$0
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

// func maskloadEpi641(mem_addr int, mask [32]byte) [32]byte
TEXT ·maskloadEpi641(SB),7,$0
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

// func maskstoreEpi321(mem_addr int, mask [32]byte, a [32]byte) 
TEXT ·maskstoreEpi321(SB),7,$0
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

// func maskstoreEpi641(mem_addr int64, mask [32]byte, a [32]byte) 
TEXT ·maskstoreEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	RET

// func maxEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpu32(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func maxEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·maxEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpu32(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func minEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·minEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func movemaskEpi8(a [32]byte) int
TEXT ·movemaskEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOVQ $0, ret+0(FP)
	RET

// func mpsadbwEpu8(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·mpsadbwEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·mulEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulEpu32(a [32]byte, b [32]byte) [32]byte
TEXT ·mulEpu32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulhiEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·mulhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulhiEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·mulhiEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulhrsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·mulhrsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulloEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·mulloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func mulloEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·mulloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func orSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·orSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func packsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·packsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func packsEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·packsEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func packusEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·packusEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func packusEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·packusEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permute2x128Si256(a [32]byte, b [32]byte, imm8 int) [32]byte
TEXT ·permute2x128Si256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permute4x64Epi64(a [32]byte, imm8 int) [32]byte
TEXT ·permute4x64Epi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permute4x64Pd(a [4]float64, imm8 int) [4]float64
TEXT ·permute4x64Pd(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256d
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutevar8x32Epi32(a [32]byte, idx [32]byte) [32]byte
TEXT ·permutevar8x32Epi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func permutevar8x32Ps(a [8]float32, idx [32]byte) [8]float32
TEXT ·permutevar8x32Ps(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sadEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·sadEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·shuffleEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·shuffleEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shufflehiEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·shufflehiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func shuffleloEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·shuffleloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func signEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·signEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func signEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·signEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func signEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·signEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sllEpi16(a [32]byte, count [16]byte) [32]byte
TEXT ·sllEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sllEpi32(a [32]byte, count [16]byte) [32]byte
TEXT ·sllEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sllEpi64(a [32]byte, count [16]byte) [32]byte
TEXT ·sllEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func slliEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·slliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func slliEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·slliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func slliEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·slliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func slliSi256(a [32]byte, imm8 int) [32]byte
TEXT ·slliSi256(SB),7,$0
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

// func sllvEpi321(a [32]byte, count [32]byte) [32]byte
TEXT ·sllvEpi321(SB),7,$0
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

// func sllvEpi641(a [32]byte, count [32]byte) [32]byte
TEXT ·sllvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sraEpi16(a [32]byte, count [16]byte) [32]byte
TEXT ·sraEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sraEpi32(a [32]byte, count [16]byte) [32]byte
TEXT ·sraEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sraiEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·sraiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func sraiEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·sraiEpi32(SB),7,$0
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

// func sravEpi321(a [32]byte, count [32]byte) [32]byte
TEXT ·sravEpi321(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srlEpi16(a [32]byte, count [16]byte) [32]byte
TEXT ·srlEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srlEpi32(a [32]byte, count [16]byte) [32]byte
TEXT ·srlEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srlEpi64(a [32]byte, count [16]byte) [32]byte
TEXT ·srlEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srliEpi16(a [32]byte, imm8 int) [32]byte
TEXT ·srliEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srliEpi32(a [32]byte, imm8 int) [32]byte
TEXT ·srliEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srliEpi64(a [32]byte, imm8 int) [32]byte
TEXT ·srliEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func srliSi256(a [32]byte, imm8 int) [32]byte
TEXT ·srliSi256(SB),7,$0
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

// func srlvEpi321(a [32]byte, count [32]byte) [32]byte
TEXT ·srlvEpi321(SB),7,$0
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

// func srlvEpi641(a [32]byte, count [32]byte) [32]byte
TEXT ·srlvEpi641(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func streamLoadSi256(mem_addr x86.M256iConst) [32]byte
TEXT ·streamLoadSi256(SB),7,$0
	// Unimplemented. Unknown size of type x86.M256iConst
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·subEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·subEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·subEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·subEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subsEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·subsEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subsEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·subsEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subsEpu16(a [32]byte, b [32]byte) [32]byte
TEXT ·subsEpu16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func subsEpu8(a [32]byte, b [32]byte) [32]byte
TEXT ·subsEpu8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackhiEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackhiEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackhiEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackhiEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackhiEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackhiEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackhiEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackhiEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackloEpi16(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackloEpi16(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackloEpi32(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackloEpi32(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackloEpi64(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackloEpi64(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func unpackloEpi8(a [32]byte, b [32]byte) [32]byte
TEXT ·unpackloEpi8(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

// func xorSi256(a [32]byte, b [32]byte) [32]byte
TEXT ·xorSi256(SB),7,$0
	// Unimplemented. Unknown MOVE postfix for type x86.M256i
	//TODO: Code missing

	MOV Y0, ret+0(FP)
	RET

