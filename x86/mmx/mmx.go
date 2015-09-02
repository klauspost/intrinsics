package mmx

import . "github.com/klauspost/intrinsics/x86"


// AddPi16: Add packed 16-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := a[i+15:i] + b[i+15:i]
//		ENDFOR
//
// Instruction: 'PADDW'. Intrinsic: '_mm_add_pi16'.
// Requires MMX.
func AddPi16(a M64, b M64) M64 {
	return M64(addPi16(a, b))
}

func addPi16(a M64, b M64) M64


// AddPi32: Add packed 32-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//
// Instruction: 'PADDD'. Intrinsic: '_mm_add_pi32'.
// Requires MMX.
func AddPi32(a M64, b M64) M64 {
	return M64(addPi32(a, b))
}

func addPi32(a M64, b M64) M64


// AddPi8: Add packed 8-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := a[i+7:i] + b[i+7:i]
//		ENDFOR
//
// Instruction: 'PADDB'. Intrinsic: '_mm_add_pi8'.
// Requires MMX.
func AddPi8(a M64, b M64) M64 {
	return M64(addPi8(a, b))
}

func addPi8(a M64, b M64) M64


// AddsPi16: Add packed 16-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//
// Instruction: 'PADDSW'. Intrinsic: '_mm_adds_pi16'.
// Requires MMX.
func AddsPi16(a M64, b M64) M64 {
	return M64(addsPi16(a, b))
}

func addsPi16(a M64, b M64) M64


// AddsPi8: Add packed 8-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PADDSB'. Intrinsic: '_mm_adds_pi8'.
// Requires MMX.
func AddsPi8(a M64, b M64) M64 {
	return M64(addsPi8(a, b))
}

func addsPi8(a M64, b M64) M64


// AddsPu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//
// Instruction: 'PADDUSW'. Intrinsic: '_mm_adds_pu16'.
// Requires MMX.
func AddsPu16(a M64, b M64) M64 {
	return M64(addsPu16(a, b))
}

func addsPu16(a M64, b M64) M64


// AddsPu8: Add packed unsigned 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PADDUSB'. Intrinsic: '_mm_adds_pu8'.
// Requires MMX.
func AddsPu8(a M64, b M64) M64 {
	return M64(addsPu8(a, b))
}

func addsPu8(a M64, b M64) M64


// AndSi64: Compute the bitwise AND of 64 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] AND b[63:0])
//
// Instruction: 'PAND'. Intrinsic: '_mm_and_si64'.
// Requires MMX.
func AndSi64(a M64, b M64) M64 {
	return M64(andSi64(a, b))
}

func andSi64(a M64, b M64) M64


// AndnotSi64: Compute the bitwise AND NOT of 64 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := ((NOT a[63:0]) AND b[63:0])
//
// Instruction: 'PANDN'. Intrinsic: '_mm_andnot_si64'.
// Requires MMX.
func AndnotSi64(a M64, b M64) M64 {
	return M64(andnotSi64(a, b))
}

func andnotSi64(a M64, b M64) M64


// CmpeqPi16: Compare packed 16-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] == b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQW'. Intrinsic: '_mm_cmpeq_pi16'.
// Requires MMX.
func CmpeqPi16(a M64, b M64) M64 {
	return M64(cmpeqPi16(a, b))
}

func cmpeqPi16(a M64, b M64) M64


// CmpeqPi32: Compare packed 32-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] == b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQD'. Intrinsic: '_mm_cmpeq_pi32'.
// Requires MMX.
func CmpeqPi32(a M64, b M64) M64 {
	return M64(cmpeqPi32(a, b))
}

func cmpeqPi32(a M64, b M64) M64


// CmpeqPi8: Compare packed 8-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] == b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQB'. Intrinsic: '_mm_cmpeq_pi8'.
// Requires MMX.
func CmpeqPi8(a M64, b M64) M64 {
	return M64(cmpeqPi8(a, b))
}

func cmpeqPi8(a M64, b M64) M64


// CmpgtPi16: Compare packed 16-bit integers in 'a' and 'b' for greater-than,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] > b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTW'. Intrinsic: '_mm_cmpgt_pi16'.
// Requires MMX.
func CmpgtPi16(a M64, b M64) M64 {
	return M64(cmpgtPi16(a, b))
}

func cmpgtPi16(a M64, b M64) M64


// CmpgtPi32: Compare packed 32-bit integers in 'a' and 'b' for greater-than,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] > b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTD'. Intrinsic: '_mm_cmpgt_pi32'.
// Requires MMX.
func CmpgtPi32(a M64, b M64) M64 {
	return M64(cmpgtPi32(a, b))
}

func cmpgtPi32(a M64, b M64) M64


// CmpgtPi8: Compare packed 8-bit integers in 'a' and 'b' for greater-than, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] > b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTB'. Intrinsic: '_mm_cmpgt_pi8'.
// Requires MMX.
func CmpgtPi8(a M64, b M64) M64 {
	return M64(cmpgtPi8(a, b))
}

func cmpgtPi8(a M64, b M64) M64


// Cvtm64Si64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtm64_si64'.
// Requires MMX.
func Cvtm64Si64(a M64) int64 {
	return int64(cvtm64Si64(a))
}

func cvtm64Si64(a M64) int64


// Cvtsi32Si64: Copy 32-bit integer 'a' to the lower elements of 'dst', and
// zero the upper element of 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_mm_cvtsi32_si64'.
// Requires MMX.
func Cvtsi32Si64(a int) M64 {
	return M64(cvtsi32Si64(a))
}

func cvtsi32Si64(a int) M64


// Cvtsi64M64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtsi64_m64'.
// Requires MMX.
func Cvtsi64M64(a int64) M64 {
	return M64(cvtsi64M64(a))
}

func cvtsi64M64(a int64) M64


// Cvtsi64Si32: Copy the lower 32-bit integer in 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_cvtsi64_si32'.
// Requires MMX.
func Cvtsi64Si32(a M64) int {
	return int(cvtsi64Si32(a))
}

func cvtsi64Si32(a M64) int


// Empty: Empty the MMX state, which marks the x87 FPU registers as available
// for use by x87 instructions. This instruction must be used at the end of all
// MMX technology procedures. 
//
//		
//
// Instruction: 'EMMS'. Intrinsic: '_m_empty'.
// Requires MMX.
func Empty()  {
	empty()
}

func empty() 


// Empty1: Empty the MMX state, which marks the x87 FPU registers as available
// for use by x87 instructions. This instruction must be used at the end of all
// MMX technology procedures. 
//
//		
//
// Instruction: 'EMMS'. Intrinsic: '_mm_empty'.
// Requires MMX.
func Empty1()  {
	empty1()
}

func empty1() 


// FromInt: Copy 32-bit integer 'a' to the lower elements of 'dst', and zero
// the upper element of 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_m_from_int'.
// Requires MMX.
func FromInt(a int) M64 {
	return M64(fromInt(a))
}

func fromInt(a int) M64


// FromInt64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_m_from_int64'.
// Requires MMX.
func FromInt64(a int64) M64 {
	return M64(fromInt64(a))
}

func fromInt64(a int64) M64


// MaddPi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			st[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//		ENDFOR
//
// Instruction: 'PMADDWD'. Intrinsic: '_mm_madd_pi16'.
// Requires MMX.
func MaddPi16(a M64, b M64) M64 {
	return M64(maddPi16(a, b))
}

func maddPi16(a M64, b M64) M64


// MulhiPi16: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the high 16 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHW'. Intrinsic: '_mm_mulhi_pi16'.
// Requires MMX.
func MulhiPi16(a M64, b M64) M64 {
	return M64(mulhiPi16(a, b))
}

func mulhiPi16(a M64, b M64) M64


// MulloPi16: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the low 16 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[15:0]
//		ENDFOR
//
// Instruction: 'PMULLW'. Intrinsic: '_mm_mullo_pi16'.
// Requires MMX.
func MulloPi16(a M64, b M64) M64 {
	return M64(mulloPi16(a, b))
}

func mulloPi16(a M64, b M64) M64


// OrSi64: Compute the bitwise OR of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] OR b[63:0])
//
// Instruction: 'POR'. Intrinsic: '_mm_or_si64'.
// Requires MMX.
func OrSi64(a M64, b M64) M64 {
	return M64(orSi64(a, b))
}

func orSi64(a M64, b M64) M64


// PacksPi16: Convert packed 16-bit integers from 'a' and 'b' to packed 8-bit
// integers using signed saturation, and store the results in 'dst'. 
//
//		dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		dst[39:32] := Saturate_Int16_To_Int8 (b[15:0])
//		dst[47:40] := Saturate_Int16_To_Int8 (b[31:16])
//		dst[55:48] := Saturate_Int16_To_Int8 (b[47:32])
//		dst[63:56] := Saturate_Int16_To_Int8 (b[63:48])
//
// Instruction: 'PACKSSWB'. Intrinsic: '_mm_packs_pi16'.
// Requires MMX.
func PacksPi16(a M64, b M64) M64 {
	return M64(packsPi16(a, b))
}

func packsPi16(a M64, b M64) M64


// PacksPi32: Convert packed 32-bit integers from 'a' and 'b' to packed 16-bit
// integers using signed saturation, and store the results in 'dst'. 
//
//		dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		dst[47:32] := Saturate_Int32_To_Int16 (b[31:0])
//		dst[63:48] := Saturate_Int32_To_Int16 (b[63:32])
//
// Instruction: 'PACKSSDW'. Intrinsic: '_mm_packs_pi32'.
// Requires MMX.
func PacksPi32(a M64, b M64) M64 {
	return M64(packsPi32(a, b))
}

func packsPi32(a M64, b M64) M64


// PacksPu16: Convert packed 16-bit integers from 'a' and 'b' to packed 8-bit
// integers using unsigned saturation, and store the results in 'dst'. 
//
//		dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		dst[39:32] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		dst[47:40] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		dst[55:48] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		dst[63:56] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//
// Instruction: 'PACKUSWB'. Intrinsic: '_mm_packs_pu16'.
// Requires MMX.
func PacksPu16(a M64, b M64) M64 {
	return M64(packsPu16(a, b))
}

func packsPu16(a M64, b M64) M64


// Packssdw: Convert packed 32-bit integers from 'a' and 'b' to packed 16-bit
// integers using signed saturation, and store the results in 'dst'. 
//
//		dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		dst[47:32] := Saturate_Int32_To_Int16 (b[31:0])
//		dst[63:48] := Saturate_Int32_To_Int16 (b[63:32])
//
// Instruction: 'PACKSSDW'. Intrinsic: '_m_packssdw'.
// Requires MMX.
func Packssdw(a M64, b M64) M64 {
	return M64(packssdw(a, b))
}

func packssdw(a M64, b M64) M64


// Packsswb: Convert packed 16-bit integers from 'a' and 'b' to packed 8-bit
// integers using signed saturation, and store the results in 'dst'. 
//
//		dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		dst[39:32] := Saturate_Int16_To_Int8 (b[15:0])
//		dst[47:40] := Saturate_Int16_To_Int8 (b[31:16])
//		dst[55:48] := Saturate_Int16_To_Int8 (b[47:32])
//		dst[63:56] := Saturate_Int16_To_Int8 (b[63:48])
//
// Instruction: 'PACKSSWB'. Intrinsic: '_m_packsswb'.
// Requires MMX.
func Packsswb(a M64, b M64) M64 {
	return M64(packsswb(a, b))
}

func packsswb(a M64, b M64) M64


// Packuswb: Convert packed 16-bit integers from 'a' and 'b' to packed 8-bit
// integers using unsigned saturation, and store the results in 'dst'. 
//
//		dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		dst[39:32] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		dst[47:40] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		dst[55:48] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		dst[63:56] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//
// Instruction: 'PACKUSWB'. Intrinsic: '_m_packuswb'.
// Requires MMX.
func Packuswb(a M64, b M64) M64 {
	return M64(packuswb(a, b))
}

func packuswb(a M64, b M64) M64


// Paddb: Add packed 8-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := a[i+7:i] + b[i+7:i]
//		ENDFOR
//
// Instruction: 'PADDB'. Intrinsic: '_m_paddb'.
// Requires MMX.
func Paddb(a M64, b M64) M64 {
	return M64(paddb(a, b))
}

func paddb(a M64, b M64) M64


// Paddd: Add packed 32-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//
// Instruction: 'PADDD'. Intrinsic: '_m_paddd'.
// Requires MMX.
func Paddd(a M64, b M64) M64 {
	return M64(paddd(a, b))
}

func paddd(a M64, b M64) M64


// Paddsb: Add packed 8-bit integers in 'a' and 'b' using saturation, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PADDSB'. Intrinsic: '_m_paddsb'.
// Requires MMX.
func Paddsb(a M64, b M64) M64 {
	return M64(paddsb(a, b))
}

func paddsb(a M64, b M64) M64


// Paddsw: Add packed 16-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//
// Instruction: 'PADDSW'. Intrinsic: '_m_paddsw'.
// Requires MMX.
func Paddsw(a M64, b M64) M64 {
	return M64(paddsw(a, b))
}

func paddsw(a M64, b M64) M64


// Paddusb: Add packed unsigned 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PADDUSB'. Intrinsic: '_m_paddusb'.
// Requires MMX.
func Paddusb(a M64, b M64) M64 {
	return M64(paddusb(a, b))
}

func paddusb(a M64, b M64) M64


// Paddusw: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//
// Instruction: 'PADDUSW'. Intrinsic: '_m_paddusw'.
// Requires MMX.
func Paddusw(a M64, b M64) M64 {
	return M64(paddusw(a, b))
}

func paddusw(a M64, b M64) M64


// Paddw: Add packed 16-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := a[i+15:i] + b[i+15:i]
//		ENDFOR
//
// Instruction: 'PADDW'. Intrinsic: '_m_paddw'.
// Requires MMX.
func Paddw(a M64, b M64) M64 {
	return M64(paddw(a, b))
}

func paddw(a M64, b M64) M64


// Pand: Compute the bitwise AND of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] AND b[63:0])
//
// Instruction: 'PAND'. Intrinsic: '_m_pand'.
// Requires MMX.
func Pand(a M64, b M64) M64 {
	return M64(pand(a, b))
}

func pand(a M64, b M64) M64


// Pandn: Compute the bitwise AND NOT of 64 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := ((NOT a[63:0]) AND b[63:0])
//
// Instruction: 'PANDN'. Intrinsic: '_m_pandn'.
// Requires MMX.
func Pandn(a M64, b M64) M64 {
	return M64(pandn(a, b))
}

func pandn(a M64, b M64) M64


// Pcmpeqb: Compare packed 8-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] == b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQB'. Intrinsic: '_m_pcmpeqb'.
// Requires MMX.
func Pcmpeqb(a M64, b M64) M64 {
	return M64(pcmpeqb(a, b))
}

func pcmpeqb(a M64, b M64) M64


// Pcmpeqd: Compare packed 32-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] == b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQD'. Intrinsic: '_m_pcmpeqd'.
// Requires MMX.
func Pcmpeqd(a M64, b M64) M64 {
	return M64(pcmpeqd(a, b))
}

func pcmpeqd(a M64, b M64) M64


// Pcmpeqw: Compare packed 16-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] == b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQW'. Intrinsic: '_m_pcmpeqw'.
// Requires MMX.
func Pcmpeqw(a M64, b M64) M64 {
	return M64(pcmpeqw(a, b))
}

func pcmpeqw(a M64, b M64) M64


// Pcmpgtb: Compare packed 8-bit integers in 'a' and 'b' for greater-than, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] > b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTB'. Intrinsic: '_m_pcmpgtb'.
// Requires MMX.
func Pcmpgtb(a M64, b M64) M64 {
	return M64(pcmpgtb(a, b))
}

func pcmpgtb(a M64, b M64) M64


// Pcmpgtd: Compare packed 32-bit integers in 'a' and 'b' for greater-than, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] > b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTD'. Intrinsic: '_m_pcmpgtd'.
// Requires MMX.
func Pcmpgtd(a M64, b M64) M64 {
	return M64(pcmpgtd(a, b))
}

func pcmpgtd(a M64, b M64) M64


// Pcmpgtw: Compare packed 16-bit integers in 'a' and 'b' for greater-than, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] > b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTW'. Intrinsic: '_m_pcmpgtw'.
// Requires MMX.
func Pcmpgtw(a M64, b M64) M64 {
	return M64(pcmpgtw(a, b))
}

func pcmpgtw(a M64, b M64) M64


// Pmaddwd: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			st[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//		ENDFOR
//
// Instruction: 'PMADDWD'. Intrinsic: '_m_pmaddwd'.
// Requires MMX.
func Pmaddwd(a M64, b M64) M64 {
	return M64(pmaddwd(a, b))
}

func pmaddwd(a M64, b M64) M64


// Pmulhw: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the high 16 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHW'. Intrinsic: '_m_pmulhw'.
// Requires MMX.
func Pmulhw(a M64, b M64) M64 {
	return M64(pmulhw(a, b))
}

func pmulhw(a M64, b M64) M64


// Pmullw: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the low 16 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[15:0]
//		ENDFOR
//
// Instruction: 'PMULLW'. Intrinsic: '_m_pmullw'.
// Requires MMX.
func Pmullw(a M64, b M64) M64 {
	return M64(pmullw(a, b))
}

func pmullw(a M64, b M64) M64


// Por: Compute the bitwise OR of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] OR b[63:0])
//
// Instruction: 'POR'. Intrinsic: '_m_por'.
// Requires MMX.
func Por(a M64, b M64) M64 {
	return M64(por(a, b))
}

func por(a M64, b M64) M64


// Pslld: Shift packed 32-bit integers in 'a' left by 'count' while shifting in
// zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLD'. Intrinsic: '_m_pslld'.
// Requires MMX.
func Pslld(a M64, count M64) M64 {
	return M64(pslld(a, count))
}

func pslld(a M64, count M64) M64


// Pslldi: Shift packed 32-bit integers in 'a' left by 'imm8' while shifting in
// zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLD'. Intrinsic: '_m_pslldi'.
// Requires MMX.
func Pslldi(a M64, imm8 int) M64 {
	return M64(pslldi(a, imm8))
}

func pslldi(a M64, imm8 int) M64


// Psllq: Shift 64-bit integer 'a' left by 'count' while shifting in zeros, and
// store the result in 'dst'. 
//
//		IF count[63:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] << count[63:0])
//		FI
//
// Instruction: 'PSLLQ'. Intrinsic: '_m_psllq'.
// Requires MMX.
func Psllq(a M64, count M64) M64 {
	return M64(psllq(a, count))
}

func psllq(a M64, count M64) M64


// Psllqi: Shift 64-bit integer 'a' left by 'imm8' while shifting in zeros, and
// store the result in 'dst'. 
//
//		IF imm8[7:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] << imm8[7:0])
//		FI
//
// Instruction: 'PSLLQ'. Intrinsic: '_m_psllqi'.
// Requires MMX.
func Psllqi(a M64, imm8 int) M64 {
	return M64(psllqi(a, imm8))
}

func psllqi(a M64, imm8 int) M64


// Psllw: Shift packed 16-bit integers in 'a' left by 'count' while shifting in
// zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLW'. Intrinsic: '_m_psllw'.
// Requires MMX.
func Psllw(a M64, count M64) M64 {
	return M64(psllw(a, count))
}

func psllw(a M64, count M64) M64


// Psllwi: Shift packed 16-bit integers in 'a' left by 'imm8' while shifting in
// zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLW'. Intrinsic: '_m_psllwi'.
// Requires MMX.
func Psllwi(a M64, imm8 int) M64 {
	return M64(psllwi(a, imm8))
}

func psllwi(a M64, imm8 int) M64


// Psrad: Shift packed 32-bit integers in 'a' right by 'count' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAD'. Intrinsic: '_m_psrad'.
// Requires MMX.
func Psrad(a M64, count M64) M64 {
	return M64(psrad(a, count))
}

func psrad(a M64, count M64) M64


// Psradi: Shift packed 32-bit integers in 'a' right by 'imm8' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAD'. Intrinsic: '_m_psradi'.
// Requires MMX.
func Psradi(a M64, imm8 int) M64 {
	return M64(psradi(a, imm8))
}

func psradi(a M64, imm8 int) M64


// Psraw: Shift packed 16-bit integers in 'a' right by 'count' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAW'. Intrinsic: '_m_psraw'.
// Requires MMX.
func Psraw(a M64, count M64) M64 {
	return M64(psraw(a, count))
}

func psraw(a M64, count M64) M64


// Psrawi: Shift packed 16-bit integers in 'a' right by 'imm8' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAW'. Intrinsic: '_m_psrawi'.
// Requires MMX.
func Psrawi(a M64, imm8 int) M64 {
	return M64(psrawi(a, imm8))
}

func psrawi(a M64, imm8 int) M64


// Psrld: Shift packed 32-bit integers in 'a' right by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLD'. Intrinsic: '_m_psrld'.
// Requires MMX.
func Psrld(a M64, count M64) M64 {
	return M64(psrld(a, count))
}

func psrld(a M64, count M64) M64


// Psrldi: Shift packed 32-bit integers in 'a' right by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLD'. Intrinsic: '_m_psrldi'.
// Requires MMX.
func Psrldi(a M64, imm8 int) M64 {
	return M64(psrldi(a, imm8))
}

func psrldi(a M64, imm8 int) M64


// Psrlq: Shift 64-bit integer 'a' right by 'count' while shifting in zeros,
// and store the result in 'dst'. 
//
//		IF count[63:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] >> count[63:0])
//		FI
//
// Instruction: 'PSRLQ'. Intrinsic: '_m_psrlq'.
// Requires MMX.
func Psrlq(a M64, count M64) M64 {
	return M64(psrlq(a, count))
}

func psrlq(a M64, count M64) M64


// Psrlqi: Shift 64-bit integer 'a' right by 'imm8' while shifting in zeros,
// and store the result in 'dst'. 
//
//		IF imm8[7:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] >> imm8[7:0])
//		FI
//
// Instruction: 'PSRLQ'. Intrinsic: '_m_psrlqi'.
// Requires MMX.
func Psrlqi(a M64, imm8 int) M64 {
	return M64(psrlqi(a, imm8))
}

func psrlqi(a M64, imm8 int) M64


// Psrlw: Shift packed 16-bit integers in 'a' right by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLW'. Intrinsic: '_m_psrlw'.
// Requires MMX.
func Psrlw(a M64, count M64) M64 {
	return M64(psrlw(a, count))
}

func psrlw(a M64, count M64) M64


// Psrlwi: Shift packed 16-bit integers in 'a' right by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLW'. Intrinsic: '_m_psrlwi'.
// Requires MMX.
func Psrlwi(a M64, imm8 int) M64 {
	return M64(psrlwi(a, imm8))
}

func psrlwi(a M64, imm8 int) M64


// Psubb: Subtract packed 8-bit integers in 'b' from packed 8-bit integers in
// 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := a[i+7:i] - b[i+7:i]
//		ENDFOR
//
// Instruction: 'PSUBB'. Intrinsic: '_m_psubb'.
// Requires MMX.
func Psubb(a M64, b M64) M64 {
	return M64(psubb(a, b))
}

func psubb(a M64, b M64) M64


// Psubd: Subtract packed 32-bit integers in 'b' from packed 32-bit integers in
// 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//
// Instruction: 'PSUBD'. Intrinsic: '_m_psubd'.
// Requires MMX.
func Psubd(a M64, b M64) M64 {
	return M64(psubd(a, b))
}

func psubd(a M64, b M64) M64


// Psubsb: Subtract packed 8-bit integers in 'b' from packed 8-bit integers in
// 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//
// Instruction: 'PSUBSB'. Intrinsic: '_m_psubsb'.
// Requires MMX.
func Psubsb(a M64, b M64) M64 {
	return M64(psubsb(a, b))
}

func psubsb(a M64, b M64) M64


// Psubsw: Subtract packed 16-bit integers in 'b' from packed 16-bit integers
// in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//		ENDFOR
//
// Instruction: 'PSUBSW'. Intrinsic: '_m_psubsw'.
// Requires MMX.
func Psubsw(a M64, b M64) M64 {
	return M64(psubsw(a, b))
}

func psubsw(a M64, b M64) M64


// Psubusb: Subtract packed unsigned 8-bit integers in 'b' from packed unsigned
// 8-bit integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//
// Instruction: 'PSUBUSB'. Intrinsic: '_m_psubusb'.
// Requires MMX.
func Psubusb(a M64, b M64) M64 {
	return M64(psubusb(a, b))
}

func psubusb(a M64, b M64) M64


// Psubusw: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])	
//		ENDFOR
//
// Instruction: 'PSUBUSW'. Intrinsic: '_m_psubusw'.
// Requires MMX.
func Psubusw(a M64, b M64) M64 {
	return M64(psubusw(a, b))
}

func psubusw(a M64, b M64) M64


// Psubw: Subtract packed 16-bit integers in 'b' from packed 16-bit integers in
// 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := a[i+15:i] - b[i+15:i]
//		ENDFOR
//
// Instruction: 'PSUBW'. Intrinsic: '_m_psubw'.
// Requires MMX.
func Psubw(a M64, b M64) M64 {
	return M64(psubw(a, b))
}

func psubw(a M64, b M64) M64


// Punpckhbw: Unpack and interleave 8-bit integers from the high half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_BYTES(src1[63:0], src2[63:0]){
//			dst[7:0] := src1[39:32]
//			dst[15:8] := src2[39:32] 
//			dst[23:16] := src1[47:40]
//			dst[31:24] := src2[47:40]
//			dst[39:32] := src1[55:48]
//			dst[47:40] := src2[55:48]
//			dst[55:48] := src1[63:56]
//			dst[63:56] := src2[63:56]
//			RETURN dst[63:0]
//		}	
//			
//		dst[63:0] := INTERLEAVE_HIGH_BYTES(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKHBW'. Intrinsic: '_m_punpckhbw'.
// Requires MMX.
func Punpckhbw(a M64, b M64) M64 {
	return M64(punpckhbw(a, b))
}

func punpckhbw(a M64, b M64) M64


// Punpckhdq: Unpack and interleave 32-bit integers from the high half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[63:32]
//		dst[63:32] := b[63:32]
//
// Instruction: 'PUNPCKHDQ'. Intrinsic: '_m_punpckhdq'.
// Requires MMX.
func Punpckhdq(a M64, b M64) M64 {
	return M64(punpckhdq(a, b))
}

func punpckhdq(a M64, b M64) M64


// Punpckhwd: Unpack and interleave 16-bit integers from the high half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_WORDS(src1[63:0], src2[63:0]){
//			dst[15:0] := src1[47:32]
//			dst[31:16] := src2[47:32]
//			dst[47:32] := src1[63:48]
//			dst[63:48] := src2[63:48]
//			RETURN dst[63:0]
//		}
//		
//		dst[63:0] := INTERLEAVE_HIGH_WORDS(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKLBW'. Intrinsic: '_m_punpckhwd'.
// Requires MMX.
func Punpckhwd(a M64, b M64) M64 {
	return M64(punpckhwd(a, b))
}

func punpckhwd(a M64, b M64) M64


// Punpcklbw: Unpack and interleave 8-bit integers from the low half of 'a' and
// 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_BYTES(src1[63:0], src2[63:0]){
//			dst[7:0] := src1[7:0] 
//			dst[15:8] := src2[7:0] 
//			dst[23:16] := src1[15:8] 
//			dst[31:24] := src2[15:8] 
//			dst[39:32] := src1[23:16] 
//			dst[47:40] := src2[23:16] 
//			dst[55:48] := src1[31:24] 
//			dst[63:56] := src2[31:24] 
//			RETURN dst[63:0]
//		}	
//		
//		dst[63:0] := INTERLEAVE_BYTES(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKLBW'. Intrinsic: '_m_punpcklbw'.
// Requires MMX.
func Punpcklbw(a M64, b M64) M64 {
	return M64(punpcklbw(a, b))
}

func punpcklbw(a M64, b M64) M64


// Punpckldq: Unpack and interleave 32-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := b[31:0]
//
// Instruction: 'PUNPCKLDQ'. Intrinsic: '_m_punpckldq'.
// Requires MMX.
func Punpckldq(a M64, b M64) M64 {
	return M64(punpckldq(a, b))
}

func punpckldq(a M64, b M64) M64


// Punpcklwd: Unpack and interleave 16-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_WORDS(src1[63:0], src2[63:0]){
//			dst[15:0] := src1[15:0] 
//			dst[31:16] := src2[15:0] 
//			dst[47:32] := src1[31:16] 
//			dst[63:48] := src2[31:16] 
//			RETURN dst[63:0]
//		}	
//		
//		dst[63:0] := INTERLEAVE_WORDS(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKLWD'. Intrinsic: '_m_punpcklwd'.
// Requires MMX.
func Punpcklwd(a M64, b M64) M64 {
	return M64(punpcklwd(a, b))
}

func punpcklwd(a M64, b M64) M64


// Pxor: Compute the bitwise OR of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] XOR b[63:0])
//
// Instruction: 'PXOR'. Intrinsic: '_m_pxor'.
// Requires MMX.
func Pxor(a M64, b M64) M64 {
	return M64(pxor(a, b))
}

func pxor(a M64, b M64) M64


// SetPi16: Set packed 16-bit integers in 'dst' with the supplied values. 
//
//		dst[15:0] := e0
//		dst[31:16] := e1
//		dst[47:32] := e2
//		dst[63:48] := e3
//
// Instruction: '...'. Intrinsic: '_mm_set_pi16'.
// Requires MMX.
func SetPi16(e3 int16, e2 int16, e1 int16, e0 int16) M64 {
	return M64(setPi16(e3, e2, e1, e0))
}

func setPi16(e3 int16, e2 int16, e1 int16, e0 int16) M64


// SetPi32: Set packed 32-bit integers in 'dst' with the supplied values. 
//
//		dst[31:0] := e0
//		dst[63:32] := e1
//
// Instruction: '...'. Intrinsic: '_mm_set_pi32'.
// Requires MMX.
func SetPi32(e1 int, e0 int) M64 {
	return M64(setPi32(e1, e0))
}

func setPi32(e1 int, e0 int) M64


// SetPi8: Set packed 8-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[7:0] := e0
//		dst[15:8] := e1
//		dst[23:16] := e2
//		dst[31:24] := e3
//		dst[39:32] := e4
//		dst[47:40] := e5
//		dst[55:48] := e6
//		dst[63:56] := e7
//
// Instruction: '...'. Intrinsic: '_mm_set_pi8'.
// Requires MMX.
func SetPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M64 {
	return M64(setPi8(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M64


// Set1Pi16: Broadcast 16-bit integer 'a' to all all elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pi16'.
// Requires MMX.
func Set1Pi16(a int16) M64 {
	return M64(set1Pi16(a))
}

func set1Pi16(a int16) M64


// Set1Pi32: Broadcast 32-bit integer 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pi32'.
// Requires MMX.
func Set1Pi32(a int) M64 {
	return M64(set1Pi32(a))
}

func set1Pi32(a int) M64


// Set1Pi8: Broadcast 8-bit integer 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pi8'.
// Requires MMX.
func Set1Pi8(a byte) M64 {
	return M64(set1Pi8(a))
}

func set1Pi8(a byte) M64


// SetrPi16: Set packed 16-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[15:0] := e3
//		dst[31:16] := e2
//		dst[47:32] := e1
//		dst[63:48] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_pi16'.
// Requires MMX.
func SetrPi16(e3 int16, e2 int16, e1 int16, e0 int16) M64 {
	return M64(setrPi16(e3, e2, e1, e0))
}

func setrPi16(e3 int16, e2 int16, e1 int16, e0 int16) M64


// SetrPi32: Set packed 32-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[31:0] := e1
//		dst[63:32] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_pi32'.
// Requires MMX.
func SetrPi32(e1 int, e0 int) M64 {
	return M64(setrPi32(e1, e0))
}

func setrPi32(e1 int, e0 int) M64


// SetrPi8: Set packed 8-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[7:0] := e7
//		dst[15:8] := e6
//		dst[23:16] := e5
//		dst[31:24] := e4
//		dst[39:32] := e3
//		dst[47:40] := e2
//		dst[55:48] := e1
//		dst[63:56] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_pi8'.
// Requires MMX.
func SetrPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M64 {
	return M64(setrPi8(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M64


// SetzeroSi64: Return vector of type __m64 with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'PXOR'. Intrinsic: '_mm_setzero_si64'.
// Requires MMX.
func SetzeroSi64() M64 {
	return M64(setzeroSi64())
}

func setzeroSi64() M64


// SllPi16: Shift packed 16-bit integers in 'a' left by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLW'. Intrinsic: '_mm_sll_pi16'.
// Requires MMX.
func SllPi16(a M64, count M64) M64 {
	return M64(sllPi16(a, count))
}

func sllPi16(a M64, count M64) M64


// SllPi32: Shift packed 32-bit integers in 'a' left by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLD'. Intrinsic: '_mm_sll_pi32'.
// Requires MMX.
func SllPi32(a M64, count M64) M64 {
	return M64(sllPi32(a, count))
}

func sllPi32(a M64, count M64) M64


// SllSi64: Shift 64-bit integer 'a' left by 'count' while shifting in zeros,
// and store the result in 'dst'. 
//
//		IF count[63:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] << count[63:0])
//		FI
//
// Instruction: 'PSLLQ'. Intrinsic: '_mm_sll_si64'.
// Requires MMX.
func SllSi64(a M64, count M64) M64 {
	return M64(sllSi64(a, count))
}

func sllSi64(a M64, count M64) M64


// SlliPi16: Shift packed 16-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLW'. Intrinsic: '_mm_slli_pi16'.
// Requires MMX.
func SlliPi16(a M64, imm8 int) M64 {
	return M64(slliPi16(a, imm8))
}

func slliPi16(a M64, imm8 int) M64


// SlliPi32: Shift packed 32-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLD'. Intrinsic: '_mm_slli_pi32'.
// Requires MMX.
func SlliPi32(a M64, imm8 int) M64 {
	return M64(slliPi32(a, imm8))
}

func slliPi32(a M64, imm8 int) M64


// SlliSi64: Shift 64-bit integer 'a' left by 'imm8' while shifting in zeros,
// and store the result in 'dst'. 
//
//		IF imm8[7:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] << imm8[7:0])
//		FI
//
// Instruction: 'PSLLQ'. Intrinsic: '_mm_slli_si64'.
// Requires MMX.
func SlliSi64(a M64, imm8 int) M64 {
	return M64(slliSi64(a, imm8))
}

func slliSi64(a M64, imm8 int) M64


// SraPi16: Shift packed 16-bit integers in 'a' right by 'count' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAW'. Intrinsic: '_mm_sra_pi16'.
// Requires MMX.
func SraPi16(a M64, count M64) M64 {
	return M64(sraPi16(a, count))
}

func sraPi16(a M64, count M64) M64


// SraPi32: Shift packed 32-bit integers in 'a' right by 'count' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAD'. Intrinsic: '_mm_sra_pi32'.
// Requires MMX.
func SraPi32(a M64, count M64) M64 {
	return M64(sraPi32(a, count))
}

func sraPi32(a M64, count M64) M64


// SraiPi16: Shift packed 16-bit integers in 'a' right by 'imm8' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAW'. Intrinsic: '_mm_srai_pi16'.
// Requires MMX.
func SraiPi16(a M64, imm8 int) M64 {
	return M64(sraiPi16(a, imm8))
}

func sraiPi16(a M64, imm8 int) M64


// SraiPi32: Shift packed 32-bit integers in 'a' right by 'imm8' while shifting
// in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAD'. Intrinsic: '_mm_srai_pi32'.
// Requires MMX.
func SraiPi32(a M64, imm8 int) M64 {
	return M64(sraiPi32(a, imm8))
}

func sraiPi32(a M64, imm8 int) M64


// SrlPi16: Shift packed 16-bit integers in 'a' right by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLW'. Intrinsic: '_mm_srl_pi16'.
// Requires MMX.
func SrlPi16(a M64, count M64) M64 {
	return M64(srlPi16(a, count))
}

func srlPi16(a M64, count M64) M64


// SrlPi32: Shift packed 32-bit integers in 'a' right by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLD'. Intrinsic: '_mm_srl_pi32'.
// Requires MMX.
func SrlPi32(a M64, count M64) M64 {
	return M64(srlPi32(a, count))
}

func srlPi32(a M64, count M64) M64


// SrlSi64: Shift 64-bit integer 'a' right by 'count' while shifting in zeros,
// and store the result in 'dst'. 
//
//		IF count[63:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] >> count[63:0])
//		FI
//
// Instruction: 'PSRLQ'. Intrinsic: '_mm_srl_si64'.
// Requires MMX.
func SrlSi64(a M64, count M64) M64 {
	return M64(srlSi64(a, count))
}

func srlSi64(a M64, count M64) M64


// SrliPi16: Shift packed 16-bit integers in 'a' right by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLW'. Intrinsic: '_mm_srli_pi16'.
// Requires MMX.
func SrliPi16(a M64, imm8 int) M64 {
	return M64(srliPi16(a, imm8))
}

func srliPi16(a M64, imm8 int) M64


// SrliPi32: Shift packed 32-bit integers in 'a' right by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLD'. Intrinsic: '_mm_srli_pi32'.
// Requires MMX.
func SrliPi32(a M64, imm8 int) M64 {
	return M64(srliPi32(a, imm8))
}

func srliPi32(a M64, imm8 int) M64


// SrliSi64: Shift 64-bit integer 'a' right by 'imm8' while shifting in zeros,
// and store the result in 'dst'. 
//
//		IF imm8[7:0] > 63
//			dst[63:0] := 0
//		ELSE
//			dst[63:0] := ZeroExtend(a[63:0] >> imm8[7:0])
//		FI
//
// Instruction: 'PSRLQ'. Intrinsic: '_mm_srli_si64'.
// Requires MMX.
func SrliSi64(a M64, imm8 int) M64 {
	return M64(srliSi64(a, imm8))
}

func srliSi64(a M64, imm8 int) M64


// SubPi16: Subtract packed 16-bit integers in 'b' from packed 16-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := a[i+15:i] - b[i+15:i]
//		ENDFOR
//
// Instruction: 'PSUBW'. Intrinsic: '_mm_sub_pi16'.
// Requires MMX.
func SubPi16(a M64, b M64) M64 {
	return M64(subPi16(a, b))
}

func subPi16(a M64, b M64) M64


// SubPi32: Subtract packed 32-bit integers in 'b' from packed 32-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//
// Instruction: 'PSUBD'. Intrinsic: '_mm_sub_pi32'.
// Requires MMX.
func SubPi32(a M64, b M64) M64 {
	return M64(subPi32(a, b))
}

func subPi32(a M64, b M64) M64


// SubPi8: Subtract packed 8-bit integers in 'b' from packed 8-bit integers in
// 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := a[i+7:i] - b[i+7:i]
//		ENDFOR
//
// Instruction: 'PSUBB'. Intrinsic: '_mm_sub_pi8'.
// Requires MMX.
func SubPi8(a M64, b M64) M64 {
	return M64(subPi8(a, b))
}

func subPi8(a M64, b M64) M64


// SubsPi16: Subtract packed 16-bit integers in 'b' from packed 16-bit integers
// in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//		ENDFOR
//
// Instruction: 'PSUBSW'. Intrinsic: '_mm_subs_pi16'.
// Requires MMX.
func SubsPi16(a M64, b M64) M64 {
	return M64(subsPi16(a, b))
}

func subsPi16(a M64, b M64) M64


// SubsPi8: Subtract packed 8-bit integers in 'b' from packed 8-bit integers in
// 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//
// Instruction: 'PSUBSB'. Intrinsic: '_mm_subs_pi8'.
// Requires MMX.
func SubsPi8(a M64, b M64) M64 {
	return M64(subsPi8(a, b))
}

func subsPi8(a M64, b M64) M64


// SubsPu16: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])	
//		ENDFOR
//
// Instruction: 'PSUBUSW'. Intrinsic: '_mm_subs_pu16'.
// Requires MMX.
func SubsPu16(a M64, b M64) M64 {
	return M64(subsPu16(a, b))
}

func subsPu16(a M64, b M64) M64


// SubsPu8: Subtract packed unsigned 8-bit integers in 'b' from packed unsigned
// 8-bit integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//
// Instruction: 'PSUBUSB'. Intrinsic: '_mm_subs_pu8'.
// Requires MMX.
func SubsPu8(a M64, b M64) M64 {
	return M64(subsPu8(a, b))
}

func subsPu8(a M64, b M64) M64


// ToInt: Copy the lower 32-bit integer in 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_m_to_int'.
// Requires MMX.
func ToInt(a M64) int {
	return int(toInt(a))
}

func toInt(a M64) int


// ToInt64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_m_to_int64'.
// Requires MMX.
func ToInt64(a M64) int64 {
	return int64(toInt64(a))
}

func toInt64(a M64) int64


// UnpackhiPi16: Unpack and interleave 16-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_WORDS(src1[63:0], src2[63:0]){
//			dst[15:0] := src1[47:32]
//			dst[31:16] := src2[47:32]
//			dst[47:32] := src1[63:48]
//			dst[63:48] := src2[63:48]
//			RETURN dst[63:0]
//		}
//		
//		dst[63:0] := INTERLEAVE_HIGH_WORDS(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKLBW'. Intrinsic: '_mm_unpackhi_pi16'.
// Requires MMX.
func UnpackhiPi16(a M64, b M64) M64 {
	return M64(unpackhiPi16(a, b))
}

func unpackhiPi16(a M64, b M64) M64


// UnpackhiPi32: Unpack and interleave 32-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[63:32]
//		dst[63:32] := b[63:32]
//
// Instruction: 'PUNPCKHDQ'. Intrinsic: '_mm_unpackhi_pi32'.
// Requires MMX.
func UnpackhiPi32(a M64, b M64) M64 {
	return M64(unpackhiPi32(a, b))
}

func unpackhiPi32(a M64, b M64) M64


// UnpackhiPi8: Unpack and interleave 8-bit integers from the high half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_BYTES(src1[63:0], src2[63:0]){
//			dst[7:0] := src1[39:32]
//			dst[15:8] := src2[39:32] 
//			dst[23:16] := src1[47:40]
//			dst[31:24] := src2[47:40]
//			dst[39:32] := src1[55:48]
//			dst[47:40] := src2[55:48]
//			dst[55:48] := src1[63:56]
//			dst[63:56] := src2[63:56]
//			RETURN dst[63:0]
//		}	
//		
//		dst[63:0] := INTERLEAVE_HIGH_BYTES(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKHBW'. Intrinsic: '_mm_unpackhi_pi8'.
// Requires MMX.
func UnpackhiPi8(a M64, b M64) M64 {
	return M64(unpackhiPi8(a, b))
}

func unpackhiPi8(a M64, b M64) M64


// UnpackloPi16: Unpack and interleave 16-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_WORDS(src1[63:0], src2[63:0]){
//			dst[15:0] := src1[15:0] 
//			dst[31:16] := src2[15:0] 
//			dst[47:32] := src1[31:16] 
//			dst[63:48] := src2[31:16] 
//			RETURN dst[63:0]
//		}	
//		
//		dst[63:0] := INTERLEAVE_WORDS(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKLWD'. Intrinsic: '_mm_unpacklo_pi16'.
// Requires MMX.
func UnpackloPi16(a M64, b M64) M64 {
	return M64(unpackloPi16(a, b))
}

func unpackloPi16(a M64, b M64) M64


// UnpackloPi32: Unpack and interleave 32-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := b[31:0]
//
// Instruction: 'PUNPCKLDQ'. Intrinsic: '_mm_unpacklo_pi32'.
// Requires MMX.
func UnpackloPi32(a M64, b M64) M64 {
	return M64(unpackloPi32(a, b))
}

func unpackloPi32(a M64, b M64) M64


// UnpackloPi8: Unpack and interleave 8-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_BYTES(src1[63:0], src2[63:0]){
//			dst[7:0] := src1[7:0] 
//			dst[15:8] := src2[7:0] 
//			dst[23:16] := src1[15:8] 
//			dst[31:24] := src2[15:8] 
//			dst[39:32] := src1[23:16] 
//			dst[47:40] := src2[23:16] 
//			dst[55:48] := src1[31:24] 
//			dst[63:56] := src2[31:24] 
//			RETURN dst[63:0]
//		}	
//		
//		dst[63:0] := INTERLEAVE_BYTES(a[63:0], b[63:0])
//
// Instruction: 'PUNPCKLBW'. Intrinsic: '_mm_unpacklo_pi8'.
// Requires MMX.
func UnpackloPi8(a M64, b M64) M64 {
	return M64(unpackloPi8(a, b))
}

func unpackloPi8(a M64, b M64) M64


// XorSi64: Compute the bitwise XOR of 64 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] XOR b[63:0])
//
// Instruction: 'PXOR'. Intrinsic: '_mm_xor_si64'.
// Requires MMX.
func XorSi64(a M64, b M64) M64 {
	return M64(xorSi64(a, b))
}

func xorSi64(a M64, b M64) M64

