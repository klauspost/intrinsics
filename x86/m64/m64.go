package m64

import . "github.com/klauspost/intrinsics/x86"

// AbsPi16: Compute the absolute value of packed 16-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := ABS(a[i+15:i])
//		ENDFOR
//
// Instruction: 'PABSW'. Intrinsic: '_mm_abs_pi16'.
// Requires SSSE3.
func AbsPi16(a M64) M64 {
	return M64(absPi16(a))
}

func absPi16(a M64) M64


// AbsPi32: Compute the absolute value of packed 32-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := ABS(a[i+31:i])
//		ENDFOR
//
// Instruction: 'PABSD'. Intrinsic: '_mm_abs_pi32'.
// Requires SSSE3.
func AbsPi32(a M64) M64 {
	return M64(absPi32(a))
}

func absPi32(a M64) M64


// AbsPi8: Compute the absolute value of packed 8-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := ABS(a[i+7:i])
//		ENDFOR
//
// Instruction: 'PABSB'. Intrinsic: '_mm_abs_pi8'.
// Requires SSSE3.
func AbsPi8(a M64) M64 {
	return M64(absPi8(a))
}

func absPi8(a M64) M64


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


// AddSi64: Add 64-bit integers 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := a[63:0] + b[63:0]
//
// Instruction: 'PADDQ'. Intrinsic: '_mm_add_si64'.
// Requires SSE2.
func AddSi64(a M64, b M64) M64 {
	return M64(addSi64(a, b))
}

func addSi64(a M64, b M64) M64


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


// AlignrPi8: Concatenate 8-byte blocks in 'a' and 'b' into a 16-byte temporary
// result, shift the result right by 'count' bytes, and store the low 16 bytes
// in 'dst'. 
//
//		tmp[127:0] := ((a[63:0] << 64) OR b[63:0]) >> (count[7:0]*8)
//		dst[63:0] := tmp[63:0]
//
// Instruction: 'PALIGNR'. Intrinsic: '_mm_alignr_pi8'.
// Requires SSSE3.
func AlignrPi8(a M64, b M64, count int) M64 {
	return M64(alignrPi8(a, b, count))
}

func alignrPi8(a M64, b M64, count int) M64


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


// AvgPu16: Average packed unsigned 16-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGW'. Intrinsic: '_mm_avg_pu16'.
// Requires SSE.
func AvgPu16(a M64, b M64) M64 {
	return M64(avgPu16(a, b))
}

func avgPu16(a M64, b M64) M64


// AvgPu8: Average packed unsigned 8-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGB'. Intrinsic: '_mm_avg_pu8'.
// Requires SSE.
func AvgPu8(a M64, b M64) M64 {
	return M64(avgPu8(a, b))
}

func avgPu8(a M64, b M64) M64


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


// CvtPi2ps: Convert packed 32-bit integers in 'b' to packed single-precision
// (32-bit) floating-point elements, store the results in the lower 2 elements
// of 'dst', and copy the upper 2 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[63:32] := Convert_Int32_To_FP32(b[63:32])
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'CVTPI2PS'. Intrinsic: '_mm_cvt_pi2ps'.
// Requires SSE.
func CvtPi2ps(a M128, b M64) M128 {
	return M128(cvtPi2ps([4]float32(a), b))
}

func cvtPi2ps(a [4]float32, b M64) [4]float32


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


// CvtpdPi32: Convert packed double-precision (64-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32(a[k+63:k])
//		ENDFOR
//
// Instruction: 'CVTPD2PI'. Intrinsic: '_mm_cvtpd_pi32'.
// Requires SSE2.
func CvtpdPi32(a M128d) M64 {
	return M64(cvtpdPi32([2]float64(a)))
}

func cvtpdPi32(a [2]float64) M64


// CvtpsPi16: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 16-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 16*j
//			k := 32*j
//			dst[i+15:i] := Convert_FP32_To_Int16(a[k+31:k])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtps_pi16'.
// Requires SSE.
func CvtpsPi16(a M128) M64 {
	return M64(cvtpsPi16([4]float32(a)))
}

func cvtpsPi16(a [4]float32) M64


// CvtpsPi32: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTPS2PI'. Intrinsic: '_mm_cvtps_pi32'.
// Requires SSE.
func CvtpsPi32(a M128) M64 {
	return M64(cvtpsPi32([4]float32(a)))
}

func cvtpsPi32(a [4]float32) M64


// CvtpsPi8: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 8-bit integers, and store the results in lower 4 elements
// of 'dst'. 
//
//		FOR j := 0 to 3
//			i := 8*j
//			k := 32*j
//			dst[i+7:i] := Convert_FP32_To_Int8(a[k+31:k])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtps_pi8'.
// Requires SSE.
func CvtpsPi8(a M128) M64 {
	return M64(cvtpsPi8([4]float32(a)))
}

func cvtpsPi8(a [4]float32) M64


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


// CvttpdPi32: Convert packed double-precision (64-bit) floating-point elements
// in 'a' to packed 32-bit integers with truncation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32_Truncate(a[k+63:k])
//		ENDFOR
//
// Instruction: 'CVTTPD2PI'. Intrinsic: '_mm_cvttpd_pi32'.
// Requires SSE2.
func CvttpdPi32(a M128d) M64 {
	return M64(cvttpdPi32([2]float64(a)))
}

func cvttpdPi32(a [2]float64) M64


// CvttpsPi32: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers with truncation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32_Truncate(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTTPS2PI'. Intrinsic: '_mm_cvttps_pi32'.
// Requires SSE.
func CvttpsPi32(a M128) M64 {
	return M64(cvttpsPi32([4]float32(a)))
}

func cvttpsPi32(a [4]float32) M64


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


// ExtractPi16: Extract a 16-bit integer from 'a', selected with 'imm8', and
// store the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_mm_extract_pi16'.
// Requires SSE.
func ExtractPi16(a M64, imm8 int) int {
	return int(extractPi16(a, imm8))
}

func extractPi16(a M64, imm8 int) int


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


// HaddPi16: Horizontally add adjacent pairs of 16-bit integers in 'a' and 'b',
// and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[31:16] + a[15:0]
//		dst[31:16] := a[63:48] + a[47:32]
//		dst[47:32] := b[31:16] + b[15:0]
//		dst[63:48] := b[63:48] + b[47:32]
//
// Instruction: 'PHADDW'. Intrinsic: '_mm_hadd_pi16'.
// Requires SSSE3.
func HaddPi16(a M64, b M64) M64 {
	return M64(haddPi16(a, b))
}

func haddPi16(a M64, b M64) M64


// HaddPi32: Horizontally add adjacent pairs of 32-bit integers in 'a' and 'b',
// and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := b[63:32] + b[31:0]
//
// Instruction: 'PHADDW'. Intrinsic: '_mm_hadd_pi32'.
// Requires SSSE3.
func HaddPi32(a M64, b M64) M64 {
	return M64(haddPi32(a, b))
}

func haddPi32(a M64, b M64) M64


// HaddsPi16: Horizontally add adjacent pairs of 16-bit integers in 'a' and 'b'
// using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[31:16] + a[15:0])
//		dst[31:16] = Saturate_To_Int16(a[63:48] + a[47:32])
//		dst[47:32] = Saturate_To_Int16(b[31:16] + b[15:0])
//		dst[63:48] = Saturate_To_Int16(b[63:48] + b[47:32])
//
// Instruction: 'PHADDSW'. Intrinsic: '_mm_hadds_pi16'.
// Requires SSSE3.
func HaddsPi16(a M64, b M64) M64 {
	return M64(haddsPi16(a, b))
}

func haddsPi16(a M64, b M64) M64


// HsubPi16: Horizontally subtract adjacent pairs of 16-bit integers in 'a' and
// 'b', and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[15:0] - a[31:16]
//		dst[31:16] := a[47:32] - a[63:48]
//		dst[47:32] := b[15:0] - b[31:16]
//		dst[63:48] := b[47:32] - b[63:48]
//
// Instruction: 'PHSUBW'. Intrinsic: '_mm_hsub_pi16'.
// Requires SSSE3.
func HsubPi16(a M64, b M64) M64 {
	return M64(hsubPi16(a, b))
}

func hsubPi16(a M64, b M64) M64


// HsubPi32: Horizontally subtract adjacent pairs of 32-bit integers in 'a' and
// 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := b[31:0] - b[63:32]
//
// Instruction: 'PHSUBD'. Intrinsic: '_mm_hsub_pi32'.
// Requires SSSE3.
func HsubPi32(a M64, b M64) M64 {
	return M64(hsubPi32(a, b))
}

func hsubPi32(a M64, b M64) M64


// HsubsPi16: Horizontally subtract adjacent pairs of 16-bit integers in 'a'
// and 'b' using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[15:0] - a[31:16])
//		dst[31:16] = Saturate_To_Int16(a[47:32] - a[63:48])
//		dst[47:32] = Saturate_To_Int16(b[15:0] - b[31:16])
//		dst[63:48] = Saturate_To_Int16(b[47:32] - b[63:48])
//
// Instruction: 'PHSUBSW'. Intrinsic: '_mm_hsubs_pi16'.
// Requires SSSE3.
func HsubsPi16(a M64, b M64) M64 {
	return M64(hsubsPi16(a, b))
}

func hsubsPi16(a M64, b M64) M64


// InsertPi16: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst'
// at the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_mm_insert_pi16'.
// Requires SSE.
func InsertPi16(a M64, i int, imm8 int) M64 {
	return M64(insertPi16(a, i, imm8))
}

func insertPi16(a M64, i int, imm8 int) M64


// LoadhPi: Load 2 single-precision (32-bit) floating-point elements from
// memory into the upper 2 elements of 'dst', and copy the lower 2 elements
// from 'a' to 'dst'. 'mem_addr' does not need to be aligned on any particular
// boundary. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := a[63:32]
//		dst[95:64] := MEM[mem_addr+31:mem_addr]
//		dst[127:96] := MEM[mem_addr+63:mem_addr+32]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_loadh_pi'.
// Requires SSE.
func LoadhPi(a M128, mem_addr M64Const) M128 {
	return M128(loadhPi([4]float32(a), mem_addr))
}

func loadhPi(a [4]float32, mem_addr M64Const) [4]float32


// LoadlPi: Load 2 single-precision (32-bit) floating-point elements from
// memory into the lower 2 elements of 'dst', and copy the upper 2 elements
// from 'a' to 'dst'. 'mem_addr' does not need to be aligned on any particular
// boundary. 
//
//		dst[31:0] := MEM[mem_addr+31:mem_addr]
//		dst[63:32] := MEM[mem_addr+63:mem_addr+32]
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_loadl_pi'.
// Requires SSE.
func LoadlPi(a M128, mem_addr M64Const) M128 {
	return M128(loadlPi([4]float32(a), mem_addr))
}

func loadlPi(a [4]float32, mem_addr M64Const) [4]float32


// LoaduSi16: Load unaligned 16-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[15:0] := MEM[mem_addr+15:mem_addr]
//		dst[MAX:16] := 0
//
// Instruction: 'MOVZWL+MOVD'. Intrinsic: '_mm_loadu_si16'.
func LoaduSi16(mem_addr uintptr) M128i {
	return M128i(loaduSi16(uintptr(mem_addr)))
}

func loaduSi16(mem_addr uintptr) [16]byte


// LoaduSi161: Load unaligned 16-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[15:0] := MEM[mem_addr+15:mem_addr]
//		dst[MAX:16] := 0
//
// Instruction: '...'. Intrinsic: '_mm_loadu_si16'.
// Requires SSE.
func LoaduSi161(mem_addr uintptr) M128i {
	return M128i(loaduSi161(uintptr(mem_addr)))
}

func loaduSi161(mem_addr uintptr) [16]byte


// LoaduSi32: Load unaligned 32-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[31:0] := MEM[mem_addr+31:mem_addr]
//		dst[MAX:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_mm_loadu_si32'.
// Requires SSE.
func LoaduSi32(mem_addr uintptr) M128i {
	return M128i(loaduSi32(uintptr(mem_addr)))
}

func loaduSi32(mem_addr uintptr) [16]byte


// LoaduSi321: Load unaligned 32-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[31:0] := MEM[mem_addr+31:mem_addr]
//		dst[MAX:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_mm_loadu_si32'.
func LoaduSi321(mem_addr uintptr) M128i {
	return M128i(loaduSi321(uintptr(mem_addr)))
}

func loaduSi321(mem_addr uintptr) [16]byte


// LoaduSi64: Load unaligned 64-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[63:0] := MEM[mem_addr+63:mem_addr]
//		dst[MAX:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_loadu_si64'.
// Requires SSE.
func LoaduSi64(mem_addr uintptr) M128i {
	return M128i(loaduSi64(uintptr(mem_addr)))
}

func loaduSi64(mem_addr uintptr) [16]byte


// LoaduSi641: Load unaligned 64-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[63:0] := MEM[mem_addr+63:mem_addr]
//		dst[MAX:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_loadu_si64'.
func LoaduSi641(mem_addr uintptr) M128i {
	return M128i(loaduSi641(uintptr(mem_addr)))
}

func loaduSi641(mem_addr uintptr) [16]byte


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


// MaddubsPi16: Vertically multiply each unsigned 8-bit integer from 'a' with
// the corresponding signed 8-bit integer from 'b', producing intermediate
// signed 16-bit integers. Horizontally add adjacent pairs of intermediate
// signed 16-bit integers, and pack the saturated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PMADDUBSW'. Intrinsic: '_mm_maddubs_pi16'.
// Requires SSSE3.
func MaddubsPi16(a M64, b M64) M64 {
	return M64(maddubsPi16(a, b))
}

func maddubsPi16(a M64, b M64) M64


// MaskmoveSi64: Conditionally store 8-bit integer elements from 'a' into
// memory using 'mask' (elements are not stored when the highest bit is not set
// in the corresponding element) and a non-temporal memory hint. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF mask[i+7]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'MASKMOVQ'. Intrinsic: '_mm_maskmove_si64'.
// Requires SSE.
func MaskmoveSi64(a M64, mask M64, mem_addr byte)  {
	maskmoveSi64(a, mask, mem_addr)
}

func maskmoveSi64(a M64, mask M64, mem_addr byte) 


// Maskmovq: Conditionally store 8-bit integer elements from 'a' into memory
// using 'mask' (elements are not stored when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF mask[i+7]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'MASKMOVQ'. Intrinsic: '_m_maskmovq'.
// Requires SSE.
func Maskmovq(a M64, mask M64, mem_addr byte)  {
	maskmovq(a, mask, mem_addr)
}

func maskmovq(a M64, mask M64, mem_addr byte) 


// MaxPi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSW'. Intrinsic: '_mm_max_pi16'.
// Requires SSE.
func MaxPi16(a M64, b M64) M64 {
	return M64(maxPi16(a, b))
}

func maxPi16(a M64, b M64) M64


// MaxPu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUB'. Intrinsic: '_mm_max_pu8'.
// Requires SSE.
func MaxPu8(a M64, b M64) M64 {
	return M64(maxPu8(a, b))
}

func maxPu8(a M64, b M64) M64


// MinPi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSW'. Intrinsic: '_mm_min_pi16'.
// Requires SSE.
func MinPi16(a M64, b M64) M64 {
	return M64(minPi16(a, b))
}

func minPi16(a M64, b M64) M64


// MinPu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUB'. Intrinsic: '_mm_min_pu8'.
// Requires SSE.
func MinPu8(a M64, b M64) M64 {
	return M64(minPu8(a, b))
}

func minPu8(a M64, b M64) M64


// MovemaskPi8: Create mask from the most significant bit of each 8-bit element
// in 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//		dst[MAX:8] := 0
//
// Instruction: 'PMOVMSKB'. Intrinsic: '_mm_movemask_pi8'.
// Requires SSE.
func MovemaskPi8(a M64) int {
	return int(movemaskPi8(a))
}

func movemaskPi8(a M64) int


// Movepi64Pi64: Copy the lower 64-bit integer in 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVDQ2Q'. Intrinsic: '_mm_movepi64_pi64'.
// Requires SSE2.
func Movepi64Pi64(a M128i) M64 {
	return M64(movepi64Pi64([16]byte(a)))
}

func movepi64Pi64(a [16]byte) M64


// MulSu32: Multiply the low unsigned 32-bit integers from 'a' and 'b', and
// store the unsigned 64-bit result in 'dst'. 
//
//		dst[63:0] := a[31:0] * b[31:0]
//
// Instruction: 'PMULUDQ'. Intrinsic: '_mm_mul_su32'.
// Requires SSE2.
func MulSu32(a M64, b M64) M64 {
	return M64(mulSu32(a, b))
}

func mulSu32(a M64, b M64) M64


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


// MulhiPu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHUW'. Intrinsic: '_mm_mulhi_pu16'.
// Requires SSE.
func MulhiPu16(a M64, b M64) M64 {
	return M64(mulhiPu16(a, b))
}

func mulhiPu16(a M64, b M64) M64


// MulhrsPi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//			dst[i+15:i] := tmp[16:1]
//		ENDFOR
//
// Instruction: 'PMULHRSW'. Intrinsic: '_mm_mulhrs_pi16'.
// Requires SSSE3.
func MulhrsPi16(a M64, b M64) M64 {
	return M64(mulhrsPi16(a, b))
}

func mulhrsPi16(a M64, b M64) M64


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


// Pavgb: Average packed unsigned 8-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGB'. Intrinsic: '_m_pavgb'.
// Requires SSE.
func Pavgb(a M64, b M64) M64 {
	return M64(pavgb(a, b))
}

func pavgb(a M64, b M64) M64


// Pavgw: Average packed unsigned 16-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGW'. Intrinsic: '_m_pavgw'.
// Requires SSE.
func Pavgw(a M64, b M64) M64 {
	return M64(pavgw(a, b))
}

func pavgw(a M64, b M64) M64


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


// Pextrw: Extract a 16-bit integer from 'a', selected with 'imm8', and store
// the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_m_pextrw'.
// Requires SSE.
func Pextrw(a M64, imm8 int) int {
	return int(pextrw(a, imm8))
}

func pextrw(a M64, imm8 int) int


// Pinsrw: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_m_pinsrw'.
// Requires SSE.
func Pinsrw(a M64, i int, imm8 int) M64 {
	return M64(pinsrw(a, i, imm8))
}

func pinsrw(a M64, i int, imm8 int) M64


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


// Pmaxsw: Compare packed 16-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSW'. Intrinsic: '_m_pmaxsw'.
// Requires SSE.
func Pmaxsw(a M64, b M64) M64 {
	return M64(pmaxsw(a, b))
}

func pmaxsw(a M64, b M64) M64


// Pmaxub: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUB'. Intrinsic: '_m_pmaxub'.
// Requires SSE.
func Pmaxub(a M64, b M64) M64 {
	return M64(pmaxub(a, b))
}

func pmaxub(a M64, b M64) M64


// Pminsw: Compare packed 16-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSW'. Intrinsic: '_m_pminsw'.
// Requires SSE.
func Pminsw(a M64, b M64) M64 {
	return M64(pminsw(a, b))
}

func pminsw(a M64, b M64) M64


// Pminub: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUB'. Intrinsic: '_m_pminub'.
// Requires SSE.
func Pminub(a M64, b M64) M64 {
	return M64(pminub(a, b))
}

func pminub(a M64, b M64) M64


// Pmovmskb: Create mask from the most significant bit of each 8-bit element in
// 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//		dst[MAX:8] := 0
//
// Instruction: 'PMOVMSKB'. Intrinsic: '_m_pmovmskb'.
// Requires SSE.
func Pmovmskb(a M64) int {
	return int(pmovmskb(a))
}

func pmovmskb(a M64) int


// Pmulhuw: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHUW'. Intrinsic: '_m_pmulhuw'.
// Requires SSE.
func Pmulhuw(a M64, b M64) M64 {
	return M64(pmulhuw(a, b))
}

func pmulhuw(a M64, b M64) M64


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


// Psadbw: Compute the absolute differences of packed unsigned 8-bit integers
// in 'a' and 'b', then horizontally sum each consecutive 8 differences to
// produce four unsigned 16-bit integers, and pack these unsigned 16-bit
// integers in the low 16 bits of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		
//		dst[15:0] := tmp[7:0] + tmp[15:8] + tmp[23:16] + tmp[31:24] + tmp[39:32] + tmp[47:40] + tmp[55:48] + tmp[63:56]
//		dst[63:16] := 0
//
// Instruction: 'PSADBW'. Intrinsic: '_m_psadbw'.
// Requires SSE.
func Psadbw(a M64, b M64) M64 {
	return M64(psadbw(a, b))
}

func psadbw(a M64, b M64) M64


// Pshufw: Shuffle 16-bit integers in 'a' using the control in 'imm8', and
// store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[15:0] := src[15:0]
//			1:	tmp[15:0] := src[31:16]
//			2:	tmp[15:0] := src[47:32]
//			3:	tmp[15:0] := src[63:48]
//			ESAC
//			RETURN tmp[15:0]
//		}
//		
//		dst[15:0] := SELECT4(a[63:0], imm8[1:0])
//		dst[31:16] := SELECT4(a[63:0], imm8[3:2])
//		dst[47:32] := SELECT4(a[63:0], imm8[5:4])
//		dst[63:48] := SELECT4(a[63:0], imm8[7:6])
//
// Instruction: 'PSHUFW'. Intrinsic: '_m_pshufw'.
// Requires SSE.
func Pshufw(a M64, imm8 int) M64 {
	return M64(pshufw(a, imm8))
}

func pshufw(a M64, imm8 int) M64


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


// SadPu8: Compute the absolute differences of packed unsigned 8-bit integers
// in 'a' and 'b', then horizontally sum each consecutive 8 differences to
// produce four unsigned 16-bit integers, and pack these unsigned 16-bit
// integers in the low 16 bits of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		
//		dst[15:0] := tmp[7:0] + tmp[15:8] + tmp[23:16] + tmp[31:24] + tmp[39:32] + tmp[47:40] + tmp[55:48] + tmp[63:56]
//		dst[63:16] := 0
//
// Instruction: 'PSADBW'. Intrinsic: '_mm_sad_pu8'.
// Requires SSE.
func SadPu8(a M64, b M64) M64 {
	return M64(sadPu8(a, b))
}

func sadPu8(a M64, b M64) M64


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


// ShufflePi16: Shuffle 16-bit integers in 'a' using the control in 'imm8', and
// store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[15:0] := src[15:0]
//			1:	tmp[15:0] := src[31:16]
//			2:	tmp[15:0] := src[47:32]
//			3:	tmp[15:0] := src[63:48]
//			ESAC
//			RETURN tmp[15:0]
//		}
//		
//		dst[15:0] := SELECT4(a[63:0], imm8[1:0])
//		dst[31:16] := SELECT4(a[63:0], imm8[3:2])
//		dst[47:32] := SELECT4(a[63:0], imm8[5:4])
//		dst[63:48] := SELECT4(a[63:0], imm8[7:6])
//
// Instruction: 'PSHUFW'. Intrinsic: '_mm_shuffle_pi16'.
// Requires SSE.
func ShufflePi16(a M64, imm8 int) M64 {
	return M64(shufflePi16(a, imm8))
}

func shufflePi16(a M64, imm8 int) M64


// ShufflePi8: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF b[i+7] == 1
//				dst[i+7:i] := 0
//			ELSE
//				index[2:0] := b[i+2:i]
//				dst[i+7:i] := a[index*8+7:index*8]
//			FI
//		ENDFOR
//
// Instruction: 'PSHUFB'. Intrinsic: '_mm_shuffle_pi8'.
// Requires SSSE3.
func ShufflePi8(a M64, b M64) M64 {
	return M64(shufflePi8(a, b))
}

func shufflePi8(a M64, b M64) M64


// SignPi16: Negate packed 16-bit integers in 'a' when the corresponding signed
// 16-bit integer in 'b' is negative, and store the results in 'dst'. Element
// in 'dst' are zeroed out when the corresponding element in 'b' is zero. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF b[i+15:i] < 0
//				dst[i+15:i] := NEG(a[i+15:i])
//			ELSE IF b[i+15:i] = 0
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PSIGNW'. Intrinsic: '_mm_sign_pi16'.
// Requires SSSE3.
func SignPi16(a M64, b M64) M64 {
	return M64(signPi16(a, b))
}

func signPi16(a M64, b M64) M64


// SignPi32: Negate packed 32-bit integers in 'a' when the corresponding signed
// 32-bit integer in 'b' is negative, and store the results in 'dst'. Element
// in 'dst' are zeroed out when the corresponding element in 'b' is zero. 
//
//		FOR j := 0 to 1
//			i := j*32
//			IF b[i+31:i] < 0
//				dst[i+31:i] := NEG(a[i+31:i])
//			ELSE IF b[i+31:i] = 0
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'PSIGND'. Intrinsic: '_mm_sign_pi32'.
// Requires SSSE3.
func SignPi32(a M64, b M64) M64 {
	return M64(signPi32(a, b))
}

func signPi32(a M64, b M64) M64


// SignPi8: Negate packed 8-bit integers in 'a' when the corresponding signed
// 8-bit integer in 'b' is negative, and store the results in 'dst'. Element in
// 'dst' are zeroed out when the corresponding element in 'b' is zero. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF b[i+7:i] < 0
//				dst[i+7:i] := NEG(a[i+7:i])
//			ELSE IF b[i+7:i] = 0
//				dst[i+7:i] := 0
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PSIGNB'. Intrinsic: '_mm_sign_pi8'.
// Requires SSSE3.
func SignPi8(a M64, b M64) M64 {
	return M64(signPi8(a, b))
}

func signPi8(a M64, b M64) M64


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


// StorehPi: Store the upper 2 single-precision (32-bit) floating-point
// elements from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[95:64]
//		MEM[mem_addr+63:mem_addr+32] := a[127:96]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_storeh_pi'.
// Requires SSE.
func StorehPi(mem_addr M64, a M128)  {
	storehPi(mem_addr, [4]float32(a))
}

func storehPi(mem_addr M64, a [4]float32) 


// StorelPi: Store the lower 2 single-precision (32-bit) floating-point
// elements from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//		MEM[mem_addr+63:mem_addr+32] := a[63:32]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_storel_pi'.
// Requires SSE.
func StorelPi(mem_addr M64, a M128)  {
	storelPi(mem_addr, [4]float32(a))
}

func storelPi(mem_addr M64, a [4]float32) 


// StoreuSi16: Store 16-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+15:mem_addr] := a[15:0]
//
// Instruction: '...'. Intrinsic: '_mm_storeu_si16'.
// Requires SSE.
func StoreuSi16(mem_addr uintptr, a M128i)  {
	storeuSi16(uintptr(mem_addr), [16]byte(a))
}

func storeuSi16(mem_addr uintptr, a [16]byte) 


// StoreuSi161: Store 16-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+15:mem_addr] := a[15:0]
//
// Instruction: 'MOVD+MOVW'. Intrinsic: '_mm_storeu_si16'.
func StoreuSi161(mem_addr uintptr, a M128i)  {
	storeuSi161(uintptr(mem_addr), [16]byte(a))
}

func storeuSi161(mem_addr uintptr, a [16]byte) 


// StoreuSi32: Store 32-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_storeu_si32'.
func StoreuSi32(mem_addr uintptr, a M128i)  {
	storeuSi32(uintptr(mem_addr), [16]byte(a))
}

func storeuSi32(mem_addr uintptr, a [16]byte) 


// StoreuSi321: Store 32-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_storeu_si32'.
// Requires SSE.
func StoreuSi321(mem_addr uintptr, a M128i)  {
	storeuSi321(uintptr(mem_addr), [16]byte(a))
}

func storeuSi321(mem_addr uintptr, a [16]byte) 


// StoreuSi64: Store 64-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_storeu_si64'.
// Requires SSE.
func StoreuSi64(mem_addr uintptr, a M128i)  {
	storeuSi64(uintptr(mem_addr), [16]byte(a))
}

func storeuSi64(mem_addr uintptr, a [16]byte) 


// StoreuSi641: Store 64-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_storeu_si64'.
func StoreuSi641(mem_addr uintptr, a M128i)  {
	storeuSi641(uintptr(mem_addr), [16]byte(a))
}

func storeuSi641(mem_addr uintptr, a [16]byte) 


// StreamPi: Store 64-bits of integer data from 'a' into memory using a
// non-temporal memory hint. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVNTQ'. Intrinsic: '_mm_stream_pi'.
// Requires SSE.
func StreamPi(mem_addr M64, a M64)  {
	streamPi(mem_addr, a)
}

func streamPi(mem_addr M64, a M64) 


// StreamSi32: Store 32-bit integer 'a' into memory using a non-temporal hint
// to minimize cache pollution. If the cache line containing address 'mem_addr'
// is already in the cache, the cache will be updated. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVNTI'. Intrinsic: '_mm_stream_si32'.
// Requires SSE2.
func StreamSi32(mem_addr int, a int)  {
	streamSi32(mem_addr, a)
}

func streamSi32(mem_addr int, a int) 


// StreamSi64: Store 64-bit integer 'a' into memory using a non-temporal hint
// to minimize cache pollution. If the cache line containing address 'mem_addr'
// is already in the cache, the cache will be updated. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVNTI'. Intrinsic: '_mm_stream_si64'.
// Requires SSE2.
func StreamSi64(mem_addr int64, a int64)  {
	streamSi64(mem_addr, a)
}

func streamSi64(mem_addr int64, a int64) 


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


// SubSi64: Subtract 64-bit integer 'b' from 64-bit integer 'a', and store the
// result in 'dst'. 
//
//		dst[63:0] := a[63:0] - b[63:0]
//
// Instruction: 'PSUBQ'. Intrinsic: '_mm_sub_si64'.
// Requires SSE2.
func SubSi64(a M64, b M64) M64 {
	return M64(subSi64(a, b))
}

func subSi64(a M64, b M64) M64


// SubborrowU32: Add unsigned 8-bit borrow 'b_in' (carry flag) to unsigned
// 32-bit integer 'a', and subtract the result from unsigned 32-bit integer
// 'b'. Store the unsigned 32-bit result in 'out', and the carry-out in 'dst'
// (carry or overflow flag). 
//
//		dst:out[31:0] := (b[31:0] - (a[31:0] + b_in));
//
// Instruction: 'SBB'. Intrinsic: '_subborrow_u32'.
func SubborrowU32(b_in uint8, a uint32, b uint32, out uint32) uint8 {
	return uint8(subborrowU32(b_in, a, b, out))
}

func subborrowU32(b_in uint8, a uint32, b uint32, out uint32) uint8


// SubborrowU64: Add unsigned 8-bit borrow 'b_in' (carry flag) to unsigned
// 64-bit integer 'a', and subtract the result from unsigned 64-bit integer
// 'b'. Store the unsigned 64-bit result in 'out', and the carry-out in 'dst'
// (carry or overflow flag). 
//
//		dst:out[63:0] := (b[63:0] - (a[63:0] + b_in));
//
// Instruction: 'SBB'. Intrinsic: '_subborrow_u64'.
func SubborrowU64(b_in uint8, a uint64, b uint64, out uint64) uint8 {
	return uint8(subborrowU64(b_in, a, b, out))
}

func subborrowU64(b_in uint8, a uint64, b uint64, out uint64) uint8


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

