// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package mmx

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


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
func AddPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AddPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AddPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AddsPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AddsPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AddsPu16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AddsPu8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// AndSi64: Compute the bitwise AND of 64 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] AND b[63:0])
//
// Instruction: 'PAND'. Intrinsic: '_mm_and_si64'.
// Requires MMX.
func AndSi64(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// AndnotSi64: Compute the bitwise AND NOT of 64 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := ((NOT a[63:0]) AND b[63:0])
//
// Instruction: 'PANDN'. Intrinsic: '_mm_andnot_si64'.
// Requires MMX.
func AndnotSi64(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func CmpeqPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func CmpeqPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func CmpeqPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func CmpgtPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func CmpgtPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func CmpgtPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Cvtm64Si64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtm64_si64'.
// Requires MMX.
func Cvtm64Si64(a x86.M64) int64 {
	panic("not implemented")
}


// Cvtsi32Si64: Copy 32-bit integer 'a' to the lower elements of 'dst', and
// zero the upper element of 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_mm_cvtsi32_si64'.
// Requires MMX.
func Cvtsi32Si64(a int) (dst x86.M64) {
	panic("not implemented")
}


// Cvtsi64M64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtsi64_m64'.
// Requires MMX.
func Cvtsi64M64(a int64) (dst x86.M64) {
	panic("not implemented")
}


// Cvtsi64Si32: Copy the lower 32-bit integer in 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_cvtsi64_si32'.
// Requires MMX.
func Cvtsi64Si32(a x86.M64) int {
	panic("not implemented")
}


// Empty: Empty the MMX state, which marks the x87 FPU registers as available
// for use by x87 instructions. This instruction must be used at the end of all
// MMX technology procedures. 
//
//		
//
// Instruction: 'EMMS'. Intrinsic: '_m_empty'.
// Requires MMX.
func Empty()  {
	panic("not implemented")
}


// Empty2: Empty the MMX state, which marks the x87 FPU registers as available
// for use by x87 instructions. This instruction must be used at the end of all
// MMX technology procedures. 
//
//		
//
// Instruction: 'EMMS'. Intrinsic: '_mm_empty'.
// Requires MMX.
func Empty2()  {
	panic("not implemented")
}


// FromInt: Copy 32-bit integer 'a' to the lower elements of 'dst', and zero
// the upper element of 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_m_from_int'.
// Requires MMX.
func FromInt(a int) (dst x86.M64) {
	panic("not implemented")
}


// FromInt64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_m_from_int64'.
// Requires MMX.
func FromInt64(a int64) (dst x86.M64) {
	panic("not implemented")
}


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
func MaddPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func MulhiPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func MulloPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// OrSi64: Compute the bitwise OR of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] OR b[63:0])
//
// Instruction: 'POR'. Intrinsic: '_mm_or_si64'.
// Requires MMX.
func OrSi64(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func PacksPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func PacksPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func PacksPu16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Packssdw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Packsswb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Packuswb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddsb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddsw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddusb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddusw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Paddw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pand: Compute the bitwise AND of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] AND b[63:0])
//
// Instruction: 'PAND'. Intrinsic: '_m_pand'.
// Requires MMX.
func Pand(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pandn: Compute the bitwise AND NOT of 64 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := ((NOT a[63:0]) AND b[63:0])
//
// Instruction: 'PANDN'. Intrinsic: '_m_pandn'.
// Requires MMX.
func Pandn(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pcmpeqb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pcmpeqd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pcmpeqw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pcmpgtb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pcmpgtd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pcmpgtw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pmaddwd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pmulhw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pmullw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Por: Compute the bitwise OR of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] OR b[63:0])
//
// Instruction: 'POR'. Intrinsic: '_m_por'.
// Requires MMX.
func Por(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Pslld(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Pslldi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psllq(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psllqi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psllw(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psllwi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psrad(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psradi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psraw(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psrawi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psrld(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psrldi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psrlq(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psrlqi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psrlw(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func Psrlwi(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubsb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubsw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubusb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubusw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Psubw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Punpckhbw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Punpckhdq: Unpack and interleave 32-bit integers from the high half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[63:32]
//		dst[63:32] := b[63:32]
//
// Instruction: 'PUNPCKHDQ'. Intrinsic: '_m_punpckhdq'.
// Requires MMX.
func Punpckhdq(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Punpckhwd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Punpcklbw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Punpckldq: Unpack and interleave 32-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := b[31:0]
//
// Instruction: 'PUNPCKLDQ'. Intrinsic: '_m_punpckldq'.
// Requires MMX.
func Punpckldq(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func Punpcklwd(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pxor: Compute the bitwise OR of 64 bits (representing integer data) in 'a'
// and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] XOR b[63:0])
//
// Instruction: 'PXOR'. Intrinsic: '_m_pxor'.
// Requires MMX.
func Pxor(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// SetPi16: Set packed 16-bit integers in 'dst' with the supplied values. 
//
//		dst[15:0] := e0
//		dst[31:16] := e1
//		dst[47:32] := e2
//		dst[63:48] := e3
//
// Instruction: '...'. Intrinsic: '_mm_set_pi16'.
// Requires MMX.
func SetPi16(e3 int16, e2 int16, e1 int16, e0 int16) (dst x86.M64) {
	panic("not implemented")
}


// SetPi32: Set packed 32-bit integers in 'dst' with the supplied values. 
//
//		dst[31:0] := e0
//		dst[63:32] := e1
//
// Instruction: '...'. Intrinsic: '_mm_set_pi32'.
// Requires MMX.
func SetPi32(e1 int, e0 int) (dst x86.M64) {
	panic("not implemented")
}


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
func SetPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) (dst x86.M64) {
	panic("not implemented")
}


// Set1Pi16: Broadcast 16-bit integer 'a' to all all elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pi16'.
// Requires MMX.
func Set1Pi16(a int16) (dst x86.M64) {
	panic("not implemented")
}


// Set1Pi32: Broadcast 32-bit integer 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pi32'.
// Requires MMX.
func Set1Pi32(a int) (dst x86.M64) {
	panic("not implemented")
}


// Set1Pi8: Broadcast 8-bit integer 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pi8'.
// Requires MMX.
func Set1Pi8(a byte) (dst x86.M64) {
	panic("not implemented")
}


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
func SetrPi16(e3 int16, e2 int16, e1 int16, e0 int16) (dst x86.M64) {
	panic("not implemented")
}


// SetrPi32: Set packed 32-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[31:0] := e1
//		dst[63:32] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_pi32'.
// Requires MMX.
func SetrPi32(e1 int, e0 int) (dst x86.M64) {
	panic("not implemented")
}


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
func SetrPi8(e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) (dst x86.M64) {
	panic("not implemented")
}


// SetzeroSi64: Return vector of type __m64 with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'PXOR'. Intrinsic: '_mm_setzero_si64'.
// Requires MMX.
func SetzeroSi64() (dst x86.M64) {
	panic("not implemented")
}


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
func SllPi16(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SllPi32(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SllSi64(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SlliPi16(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SlliPi32(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SlliSi64(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func SraPi16(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SraPi32(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SraiPi16(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SraiPi32(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func SrlPi16(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SrlPi32(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SrlSi64(a x86.M64, count x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SrliPi16(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SrliPi32(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func SrliSi64(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


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
func SubPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SubPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SubPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SubsPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SubsPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SubsPu16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SubsPu8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// ToInt: Copy the lower 32-bit integer in 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_m_to_int'.
// Requires MMX.
func ToInt(a x86.M64) int {
	panic("not implemented")
}


// ToInt64: Copy 64-bit integer 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_m_to_int64'.
// Requires MMX.
func ToInt64(a x86.M64) int64 {
	panic("not implemented")
}


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
func UnpackhiPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// UnpackhiPi32: Unpack and interleave 32-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[63:32]
//		dst[63:32] := b[63:32]
//
// Instruction: 'PUNPCKHDQ'. Intrinsic: '_mm_unpackhi_pi32'.
// Requires MMX.
func UnpackhiPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func UnpackhiPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func UnpackloPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// UnpackloPi32: Unpack and interleave 32-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := b[31:0]
//
// Instruction: 'PUNPCKLDQ'. Intrinsic: '_mm_unpacklo_pi32'.
// Requires MMX.
func UnpackloPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func UnpackloPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// XorSi64: Compute the bitwise XOR of 64 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := (a[63:0] XOR b[63:0])
//
// Instruction: 'PXOR'. Intrinsic: '_mm_xor_si64'.
// Requires MMX.
func XorSi64(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}

