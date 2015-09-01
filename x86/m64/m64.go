package m64

import . "github.com/klauspost/intrinsics/x86"

// Abs16: Compute the absolute value of packed 16-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := ABS(a[i+15:i])
//		ENDFOR
//
// Instruction: 'PABSW'. Intrinsic: '_mm_abs_pi16'.
// Requires SSSE3.
func Abs16(a M64) M64 {
	return M64(abs16(a))
}

func abs16(a M64) M64


// Abs32: Compute the absolute value of packed 32-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			dst[i+31:i] := ABS(a[i+31:i])
//		ENDFOR
//
// Instruction: 'PABSD'. Intrinsic: '_mm_abs_pi32'.
// Requires SSSE3.
func Abs32(a M64) M64 {
	return M64(abs32(a))
}

func abs32(a M64) M64


// Abs8: Compute the absolute value of packed 8-bit integers in 'a', and store
// the unsigned results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := ABS(a[i+7:i])
//		ENDFOR
//
// Instruction: 'PABSB'. Intrinsic: '_mm_abs_pi8'.
// Requires SSSE3.
func Abs8(a M64) M64 {
	return M64(abs8(a))
}

func abs8(a M64) M64


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


// Alignr8: Concatenate 8-byte blocks in 'a' and 'b' into a 16-byte temporary
// result, shift the result right by 'count' bytes, and store the low 16 bytes
// in 'dst'. 
//
//		tmp[127:0] := ((a[63:0] << 64) OR b[63:0]) >> (count[7:0]*8)
//		dst[63:0] := tmp[63:0]
//
// Instruction: 'PALIGNR'. Intrinsic: '_mm_alignr_pi8'.
// Requires SSSE3.
func Alignr8(a M64, b M64, count int) M64 {
	return M64(alignr8(a, b, count))
}

func alignr8(a M64, b M64, count int) M64


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


// Cvt2ps: Convert packed 32-bit integers in 'b' to packed single-precision
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
func Cvt2ps(a M128, b M64) M128 {
	return M128(cvt2ps([4]float32(a), b))
}

func cvt2ps(a [4]float32, b M64) [4]float32


// Cvtpd32: Convert packed double-precision (64-bit) floating-point elements in
// 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32(a[k+63:k])
//		ENDFOR
//
// Instruction: 'CVTPD2PI'. Intrinsic: '_mm_cvtpd_pi32'.
// Requires SSE2.
func Cvtpd32(a M128d) M64 {
	return M64(cvtpd32([2]float64(a)))
}

func cvtpd32(a [2]float64) M64


// Cvtps16: Convert packed single-precision (32-bit) floating-point elements in
// 'a' to packed 16-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 16*j
//			k := 32*j
//			dst[i+15:i] := Convert_FP32_To_Int16(a[k+31:k])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtps_pi16'.
// Requires SSE.
func Cvtps16(a M128) M64 {
	return M64(cvtps16([4]float32(a)))
}

func cvtps16(a [4]float32) M64


// Cvtps32: Convert packed single-precision (32-bit) floating-point elements in
// 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTPS2PI'. Intrinsic: '_mm_cvtps_pi32'.
// Requires SSE.
func Cvtps32(a M128) M64 {
	return M64(cvtps32([4]float32(a)))
}

func cvtps32(a [4]float32) M64


// Cvtps8: Convert packed single-precision (32-bit) floating-point elements in
// 'a' to packed 8-bit integers, and store the results in lower 4 elements of
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := 8*j
//			k := 32*j
//			dst[i+7:i] := Convert_FP32_To_Int8(a[k+31:k])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtps_pi8'.
// Requires SSE.
func Cvtps8(a M128) M64 {
	return M64(cvtps8([4]float32(a)))
}

func cvtps8(a [4]float32) M64


// Cvttpd32: Convert packed double-precision (64-bit) floating-point elements
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
func Cvttpd32(a M128d) M64 {
	return M64(cvttpd32([2]float64(a)))
}

func cvttpd32(a [2]float64) M64


// Cvttps32: Convert packed single-precision (32-bit) floating-point elements
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
func Cvttps32(a M128) M64 {
	return M64(cvttps32([4]float32(a)))
}

func cvttps32(a [4]float32) M64


// Extract16: Extract a 16-bit integer from 'a', selected with 'imm8', and
// store the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_mm_extract_pi16'.
// Requires SSE.
func Extract16(a M64, imm8 int) int {
	return int(extract16(a, imm8))
}

func extract16(a M64, imm8 int) int


// Hadd16: Horizontally add adjacent pairs of 16-bit integers in 'a' and 'b',
// and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[31:16] + a[15:0]
//		dst[31:16] := a[63:48] + a[47:32]
//		dst[47:32] := b[31:16] + b[15:0]
//		dst[63:48] := b[63:48] + b[47:32]
//
// Instruction: 'PHADDW'. Intrinsic: '_mm_hadd_pi16'.
// Requires SSSE3.
func Hadd16(a M64, b M64) M64 {
	return M64(hadd16(a, b))
}

func hadd16(a M64, b M64) M64


// Hadd32: Horizontally add adjacent pairs of 32-bit integers in 'a' and 'b',
// and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := b[63:32] + b[31:0]
//
// Instruction: 'PHADDW'. Intrinsic: '_mm_hadd_pi32'.
// Requires SSSE3.
func Hadd32(a M64, b M64) M64 {
	return M64(hadd32(a, b))
}

func hadd32(a M64, b M64) M64


// Hadds16: Horizontally add adjacent pairs of 16-bit integers in 'a' and 'b'
// using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[31:16] + a[15:0])
//		dst[31:16] = Saturate_To_Int16(a[63:48] + a[47:32])
//		dst[47:32] = Saturate_To_Int16(b[31:16] + b[15:0])
//		dst[63:48] = Saturate_To_Int16(b[63:48] + b[47:32])
//
// Instruction: 'PHADDSW'. Intrinsic: '_mm_hadds_pi16'.
// Requires SSSE3.
func Hadds16(a M64, b M64) M64 {
	return M64(hadds16(a, b))
}

func hadds16(a M64, b M64) M64


// Hsub16: Horizontally subtract adjacent pairs of 16-bit integers in 'a' and
// 'b', and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[15:0] - a[31:16]
//		dst[31:16] := a[47:32] - a[63:48]
//		dst[47:32] := b[15:0] - b[31:16]
//		dst[63:48] := b[47:32] - b[63:48]
//
// Instruction: 'PHSUBW'. Intrinsic: '_mm_hsub_pi16'.
// Requires SSSE3.
func Hsub16(a M64, b M64) M64 {
	return M64(hsub16(a, b))
}

func hsub16(a M64, b M64) M64


// Hsub32: Horizontally subtract adjacent pairs of 32-bit integers in 'a' and
// 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := b[31:0] - b[63:32]
//
// Instruction: 'PHSUBD'. Intrinsic: '_mm_hsub_pi32'.
// Requires SSSE3.
func Hsub32(a M64, b M64) M64 {
	return M64(hsub32(a, b))
}

func hsub32(a M64, b M64) M64


// Hsubs16: Horizontally subtract adjacent pairs of 16-bit integers in 'a' and
// 'b' using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[15:0] - a[31:16])
//		dst[31:16] = Saturate_To_Int16(a[47:32] - a[63:48])
//		dst[47:32] = Saturate_To_Int16(b[15:0] - b[31:16])
//		dst[63:48] = Saturate_To_Int16(b[47:32] - b[63:48])
//
// Instruction: 'PHSUBSW'. Intrinsic: '_mm_hsubs_pi16'.
// Requires SSSE3.
func Hsubs16(a M64, b M64) M64 {
	return M64(hsubs16(a, b))
}

func hsubs16(a M64, b M64) M64


// Insert16: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_mm_insert_pi16'.
// Requires SSE.
func Insert16(a M64, i int, imm8 int) M64 {
	return M64(insert16(a, i, imm8))
}

func insert16(a M64, i int, imm8 int) M64


// Loadh: Load 2 single-precision (32-bit) floating-point elements from memory
// into the upper 2 elements of 'dst', and copy the lower 2 elements from 'a'
// to 'dst'. 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := a[63:32]
//		dst[95:64] := MEM[mem_addr+31:mem_addr]
//		dst[127:96] := MEM[mem_addr+63:mem_addr+32]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_loadh_pi'.
// Requires SSE.
func Loadh(a M128, mem_addr M64Const) M128 {
	return M128(loadh([4]float32(a), mem_addr))
}

func loadh(a [4]float32, mem_addr M64Const) [4]float32


// Loadl: Load 2 single-precision (32-bit) floating-point elements from memory
// into the lower 2 elements of 'dst', and copy the upper 2 elements from 'a'
// to 'dst'. 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[31:0] := MEM[mem_addr+31:mem_addr]
//		dst[63:32] := MEM[mem_addr+63:mem_addr+32]
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_loadl_pi'.
// Requires SSE.
func Loadl(a M128, mem_addr M64Const) M128 {
	return M128(loadl([4]float32(a), mem_addr))
}

func loadl(a [4]float32, mem_addr M64Const) [4]float32


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

func loaduSi16(mem_addr uintptr) M128i


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

func loaduSi161(mem_addr uintptr) M128i


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

func loaduSi32(mem_addr uintptr) M128i


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

func loaduSi321(mem_addr uintptr) M128i


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

func loaduSi64(mem_addr uintptr) M128i


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

func loaduSi641(mem_addr uintptr) M128i


// Maddubs16: Vertically multiply each unsigned 8-bit integer from 'a' with the
// corresponding signed 8-bit integer from 'b', producing intermediate signed
// 16-bit integers. Horizontally add adjacent pairs of intermediate signed
// 16-bit integers, and pack the saturated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PMADDUBSW'. Intrinsic: '_mm_maddubs_pi16'.
// Requires SSSE3.
func Maddubs16(a M64, b M64) M64 {
	return M64(maddubs16(a, b))
}

func maddubs16(a M64, b M64) M64


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


// MMaskmovq: Conditionally store 8-bit integer elements from 'a' into memory
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
func MMaskmovq(a M64, mask M64, mem_addr byte)  {
	mMaskmovq(a, mask, mem_addr)
}

func mMaskmovq(a M64, mask M64, mem_addr byte) 


// Max16: Compare packed 16-bit integers in 'a' and 'b', and store packed
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
func Max16(a M64, b M64) M64 {
	return M64(max16(a, b))
}

func max16(a M64, b M64) M64


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


// Min16: Compare packed 16-bit integers in 'a' and 'b', and store packed
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
func Min16(a M64, b M64) M64 {
	return M64(min16(a, b))
}

func min16(a M64, b M64) M64


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


// Movemask8: Create mask from the most significant bit of each 8-bit element
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
func Movemask8(a M64) int {
	return int(movemask8(a))
}

func movemask8(a M64) int


// Movepi6464: Copy the lower 64-bit integer in 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVDQ2Q'. Intrinsic: '_mm_movepi64_pi64'.
// Requires SSE2.
func Movepi6464(a M128i) M64 {
	return M64(movepi6464(a))
}

func movepi6464(a M128i) M64


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


// Mulhrs16: Multiply packed 16-bit integers in 'a' and 'b', producing
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
func Mulhrs16(a M64, b M64) M64 {
	return M64(mulhrs16(a, b))
}

func mulhrs16(a M64, b M64) M64


// MPavgb: Average packed unsigned 8-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGB'. Intrinsic: '_m_pavgb'.
// Requires SSE.
func MPavgb(a M64, b M64) M64 {
	return M64(mPavgb(a, b))
}

func mPavgb(a M64, b M64) M64


// MPavgw: Average packed unsigned 16-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGW'. Intrinsic: '_m_pavgw'.
// Requires SSE.
func MPavgw(a M64, b M64) M64 {
	return M64(mPavgw(a, b))
}

func mPavgw(a M64, b M64) M64


// MPextrw: Extract a 16-bit integer from 'a', selected with 'imm8', and store
// the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_m_pextrw'.
// Requires SSE.
func MPextrw(a M64, imm8 int) int {
	return int(mPextrw(a, imm8))
}

func mPextrw(a M64, imm8 int) int


// Mnsrw: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_m_pinsrw'.
// Requires SSE.
func Mnsrw(a M64, i int, imm8 int) M64 {
	return M64(mnsrw(a, i, imm8))
}

func mnsrw(a M64, i int, imm8 int) M64


// MPmaxsw: Compare packed 16-bit integers in 'a' and 'b', and store packed
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
func MPmaxsw(a M64, b M64) M64 {
	return M64(mPmaxsw(a, b))
}

func mPmaxsw(a M64, b M64) M64


// MPmaxub: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
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
func MPmaxub(a M64, b M64) M64 {
	return M64(mPmaxub(a, b))
}

func mPmaxub(a M64, b M64) M64


// MPminsw: Compare packed 16-bit integers in 'a' and 'b', and store packed
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
func MPminsw(a M64, b M64) M64 {
	return M64(mPminsw(a, b))
}

func mPminsw(a M64, b M64) M64


// MPminub: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
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
func MPminub(a M64, b M64) M64 {
	return M64(mPminub(a, b))
}

func mPminub(a M64, b M64) M64


// MPmovmskb: Create mask from the most significant bit of each 8-bit element
// in 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//		dst[MAX:8] := 0
//
// Instruction: 'PMOVMSKB'. Intrinsic: '_m_pmovmskb'.
// Requires SSE.
func MPmovmskb(a M64) int {
	return int(mPmovmskb(a))
}

func mPmovmskb(a M64) int


// MPmulhuw: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
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
func MPmulhuw(a M64, b M64) M64 {
	return M64(mPmulhuw(a, b))
}

func mPmulhuw(a M64, b M64) M64


// MPsadbw: Compute the absolute differences of packed unsigned 8-bit integers
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
func MPsadbw(a M64, b M64) M64 {
	return M64(mPsadbw(a, b))
}

func mPsadbw(a M64, b M64) M64


// MPshufw: Shuffle 16-bit integers in 'a' using the control in 'imm8', and
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
func MPshufw(a M64, imm8 int) M64 {
	return M64(mPshufw(a, imm8))
}

func mPshufw(a M64, imm8 int) M64


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


// Shuffle16: Shuffle 16-bit integers in 'a' using the control in 'imm8', and
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
func Shuffle16(a M64, imm8 int) M64 {
	return M64(shuffle16(a, imm8))
}

func shuffle16(a M64, imm8 int) M64


// Shuffle8: Shuffle packed 8-bit integers in 'a' according to shuffle control
// mask in the corresponding 8-bit element of 'b', and store the results in
// 'dst'. 
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
func Shuffle8(a M64, b M64) M64 {
	return M64(shuffle8(a, b))
}

func shuffle8(a M64, b M64) M64


// Sign16: Negate packed 16-bit integers in 'a' when the corresponding signed
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
func Sign16(a M64, b M64) M64 {
	return M64(sign16(a, b))
}

func sign16(a M64, b M64) M64


// Sign32: Negate packed 32-bit integers in 'a' when the corresponding signed
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
func Sign32(a M64, b M64) M64 {
	return M64(sign32(a, b))
}

func sign32(a M64, b M64) M64


// Sign8: Negate packed 8-bit integers in 'a' when the corresponding signed
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
func Sign8(a M64, b M64) M64 {
	return M64(sign8(a, b))
}

func sign8(a M64, b M64) M64


// Storeh: Store the upper 2 single-precision (32-bit) floating-point elements
// from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[95:64]
//		MEM[mem_addr+63:mem_addr+32] := a[127:96]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_storeh_pi'.
// Requires SSE.
func Storeh(mem_addr M64, a M128)  {
	storeh(mem_addr, [4]float32(a))
}

func storeh(mem_addr M64, a [4]float32) 


// Storel: Store the lower 2 single-precision (32-bit) floating-point elements
// from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//		MEM[mem_addr+63:mem_addr+32] := a[63:32]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_storel_pi'.
// Requires SSE.
func Storel(mem_addr M64, a M128)  {
	storel(mem_addr, [4]float32(a))
}

func storel(mem_addr M64, a [4]float32) 


// StoreuSi16: Store 16-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+15:mem_addr] := a[15:0]
//
// Instruction: '...'. Intrinsic: '_mm_storeu_si16'.
// Requires SSE.
func StoreuSi16(mem_addr uintptr, a M128i)  {
	storeuSi16(uintptr(mem_addr), a)
}

func storeuSi16(mem_addr uintptr, a M128i) 


// StoreuSi161: Store 16-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+15:mem_addr] := a[15:0]
//
// Instruction: 'MOVD+MOVW'. Intrinsic: '_mm_storeu_si16'.
func StoreuSi161(mem_addr uintptr, a M128i)  {
	storeuSi161(uintptr(mem_addr), a)
}

func storeuSi161(mem_addr uintptr, a M128i) 


// StoreuSi32: Store 32-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_storeu_si32'.
func StoreuSi32(mem_addr uintptr, a M128i)  {
	storeuSi32(uintptr(mem_addr), a)
}

func storeuSi32(mem_addr uintptr, a M128i) 


// StoreuSi321: Store 32-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_storeu_si32'.
// Requires SSE.
func StoreuSi321(mem_addr uintptr, a M128i)  {
	storeuSi321(uintptr(mem_addr), a)
}

func storeuSi321(mem_addr uintptr, a M128i) 


// StoreuSi64: Store 64-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_storeu_si64'.
// Requires SSE.
func StoreuSi64(mem_addr uintptr, a M128i)  {
	storeuSi64(uintptr(mem_addr), a)
}

func storeuSi64(mem_addr uintptr, a M128i) 


// StoreuSi641: Store 64-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_storeu_si64'.
func StoreuSi641(mem_addr uintptr, a M128i)  {
	storeuSi641(uintptr(mem_addr), a)
}

func storeuSi641(mem_addr uintptr, a M128i) 


// Stream: Store 64-bits of integer data from 'a' into memory using a
// non-temporal memory hint. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVNTQ'. Intrinsic: '_mm_stream_pi'.
// Requires SSE.
func Stream(mem_addr M64, a M64)  {
	stream(mem_addr, a)
}

func stream(mem_addr M64, a M64) 


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

