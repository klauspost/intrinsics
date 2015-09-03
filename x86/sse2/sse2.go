package sse2

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := a[i+15:i] + b[i+15:i]
//		ENDFOR
//
// Instruction: 'PADDW'. Intrinsic: '_mm_add_epi16'.
// Requires SSE2.
func AddEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addEpi16([16]byte(a), [16]byte(b)))
}

func addEpi16(a [16]byte, b [16]byte) [16]byte


// AddEpi32: Add packed 32-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//
// Instruction: 'PADDD'. Intrinsic: '_mm_add_epi32'.
// Requires SSE2.
func AddEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addEpi32([16]byte(a), [16]byte(b)))
}

func addEpi32(a [16]byte, b [16]byte) [16]byte


// AddEpi64: Add packed 64-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//
// Instruction: 'PADDQ'. Intrinsic: '_mm_add_epi64'.
// Requires SSE2.
func AddEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addEpi64([16]byte(a), [16]byte(b)))
}

func addEpi64(a [16]byte, b [16]byte) [16]byte


// AddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := a[i+7:i] + b[i+7:i]
//		ENDFOR
//
// Instruction: 'PADDB'. Intrinsic: '_mm_add_epi8'.
// Requires SSE2.
func AddEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addEpi8([16]byte(a), [16]byte(b)))
}

func addEpi8(a [16]byte, b [16]byte) [16]byte


// AddPd: Add packed double-precision (64-bit) floating-point elements in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//
// Instruction: 'ADDPD'. Intrinsic: '_mm_add_pd'.
// Requires SSE2.
func AddPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(addPd([2]float64(a), [2]float64(b)))
}

func addPd(a [2]float64, b [2]float64) [2]float64


// AddSd: Add the lower double-precision (64-bit) floating-point element in 'a'
// and 'b', store the result in the lower element of 'dst', and copy the upper
// element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := a[63:0] + b[63:0]
//		dst[127:64] := a[127:64]
//
// Instruction: 'ADDSD'. Intrinsic: '_mm_add_sd'.
// Requires SSE2.
func AddSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(addSd([2]float64(a), [2]float64(b)))
}

func addSd(a [2]float64, b [2]float64) [2]float64


// AddSi64: Add 64-bit integers 'a' and 'b', and store the result in 'dst'. 
//
//		dst[63:0] := a[63:0] + b[63:0]
//
// Instruction: 'PADDQ'. Intrinsic: '_mm_add_si64'.
// Requires SSE2.
func AddSi64(a x86.M64, b x86.M64) (dst x86.M64) {
	return x86.M64(addSi64(a, b))
}

func addSi64(a x86.M64, b x86.M64) x86.M64


// AddsEpi16: Add packed 16-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//
// Instruction: 'PADDSW'. Intrinsic: '_mm_adds_epi16'.
// Requires SSE2.
func AddsEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addsEpi16([16]byte(a), [16]byte(b)))
}

func addsEpi16(a [16]byte, b [16]byte) [16]byte


// AddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PADDSB'. Intrinsic: '_mm_adds_epi8'.
// Requires SSE2.
func AddsEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addsEpi8([16]byte(a), [16]byte(b)))
}

func addsEpi8(a [16]byte, b [16]byte) [16]byte


// AddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//
// Instruction: 'PADDUSW'. Intrinsic: '_mm_adds_epu16'.
// Requires SSE2.
func AddsEpu16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addsEpu16([16]byte(a), [16]byte(b)))
}

func addsEpu16(a [16]byte, b [16]byte) [16]byte


// AddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PADDUSB'. Intrinsic: '_mm_adds_epu8'.
// Requires SSE2.
func AddsEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(addsEpu8([16]byte(a), [16]byte(b)))
}

func addsEpu8(a [16]byte, b [16]byte) [16]byte


// AndPd: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//		ENDFOR
//
// Instruction: 'ANDPD'. Intrinsic: '_mm_and_pd'.
// Requires SSE2.
func AndPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(andPd([2]float64(a), [2]float64(b)))
}

func andPd(a [2]float64, b [2]float64) [2]float64


// AndSi128: Compute the bitwise AND of 128 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[127:0] := (a[127:0] AND b[127:0])
//
// Instruction: 'PAND'. Intrinsic: '_mm_and_si128'.
// Requires SSE2.
func AndSi128(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(andSi128([16]byte(a), [16]byte(b)))
}

func andSi128(a [16]byte, b [16]byte) [16]byte


// AndnotPd: Compute the bitwise AND NOT of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//		ENDFOR
//
// Instruction: 'ANDNPD'. Intrinsic: '_mm_andnot_pd'.
// Requires SSE2.
func AndnotPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(andnotPd([2]float64(a), [2]float64(b)))
}

func andnotPd(a [2]float64, b [2]float64) [2]float64


// AndnotSi128: Compute the bitwise AND NOT of 128 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[127:0] := ((NOT a[127:0]) AND b[127:0])
//
// Instruction: 'PANDN'. Intrinsic: '_mm_andnot_si128'.
// Requires SSE2.
func AndnotSi128(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(andnotSi128([16]byte(a), [16]byte(b)))
}

func andnotSi128(a [16]byte, b [16]byte) [16]byte


// AvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGW'. Intrinsic: '_mm_avg_epu16'.
// Requires SSE2.
func AvgEpu16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(avgEpu16([16]byte(a), [16]byte(b)))
}

func avgEpu16(a [16]byte, b [16]byte) [16]byte


// AvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGB'. Intrinsic: '_mm_avg_epu8'.
// Requires SSE2.
func AvgEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(avgEpu8([16]byte(a), [16]byte(b)))
}

func avgEpu8(a [16]byte, b [16]byte) [16]byte


// BslliSi128: Shift 'a' left by 'imm8' bytes while shifting in zeros, and
// store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] << (tmp*8)
//
// Instruction: 'PSLLDQ'. Intrinsic: '_mm_bslli_si128'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func BslliSi128(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(bslliSi128([16]byte(a), imm8))
}

func bslliSi128(a [16]byte, imm8 byte) [16]byte


// BsrliSi128: Shift 'a' right by 'imm8' bytes while shifting in zeros, and
// store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] >> (tmp*8)
//
// Instruction: 'PSRLDQ'. Intrinsic: '_mm_bsrli_si128'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func BsrliSi128(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(bsrliSi128([16]byte(a), imm8))
}

func bsrliSi128(a [16]byte, imm8 byte) [16]byte


// CastpdPs: Cast vector of type __m128d to type __m128. This intrinsic is only
// used for compilation and does not generate any instructions, thus it has
// zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_castpd_ps'.
// Requires SSE2.
func CastpdPs(a x86.M128d) (dst x86.M128) {
	return x86.M128(castpdPs([2]float64(a)))
}

func castpdPs(a [2]float64) [4]float32


// CastpdSi128: Cast vector of type __m128d to type __m128i. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_castpd_si128'.
// Requires SSE2.
func CastpdSi128(a x86.M128d) (dst x86.M128i) {
	return x86.M128i(castpdSi128([2]float64(a)))
}

func castpdSi128(a [2]float64) [16]byte


// CastpsPd: Cast vector of type __m128 to type __m128d. This intrinsic is only
// used for compilation and does not generate any instructions, thus it has
// zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_castps_pd'.
// Requires SSE2.
func CastpsPd(a x86.M128) (dst x86.M128d) {
	return x86.M128d(castpsPd([4]float32(a)))
}

func castpsPd(a [4]float32) [2]float64


// CastpsSi128: Cast vector of type __m128 to type __m128i. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_castps_si128'.
// Requires SSE2.
func CastpsSi128(a x86.M128) (dst x86.M128i) {
	return x86.M128i(castpsSi128([4]float32(a)))
}

func castpsSi128(a [4]float32) [16]byte


// Castsi128Pd: Cast vector of type __m128i to type __m128d. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_castsi128_pd'.
// Requires SSE2.
func Castsi128Pd(a x86.M128i) (dst x86.M128d) {
	return x86.M128d(castsi128Pd([16]byte(a)))
}

func castsi128Pd(a [16]byte) [2]float64


// Castsi128Ps: Cast vector of type __m128i to type __m128. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_castsi128_ps'.
// Requires SSE2.
func Castsi128Ps(a x86.M128i) (dst x86.M128) {
	return x86.M128(castsi128Ps([16]byte(a)))
}

func castsi128Ps(a [16]byte) [4]float32


// Skipped: _mm_clflush. Contains pointer parameter.


// CmpeqEpi16: Compare packed 16-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] == b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQW'. Intrinsic: '_mm_cmpeq_epi16'.
// Requires SSE2.
func CmpeqEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpeqEpi16([16]byte(a), [16]byte(b)))
}

func cmpeqEpi16(a [16]byte, b [16]byte) [16]byte


// CmpeqEpi32: Compare packed 32-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] == b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQD'. Intrinsic: '_mm_cmpeq_epi32'.
// Requires SSE2.
func CmpeqEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpeqEpi32([16]byte(a), [16]byte(b)))
}

func cmpeqEpi32(a [16]byte, b [16]byte) [16]byte


// CmpeqEpi8: Compare packed 8-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] == b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQB'. Intrinsic: '_mm_cmpeq_epi8'.
// Requires SSE2.
func CmpeqEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpeqEpi8([16]byte(a), [16]byte(b)))
}

func cmpeqEpi8(a [16]byte, b [16]byte) [16]byte


// CmpeqPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' for equality, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] == b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpeq_pd'.
// Requires SSE2.
func CmpeqPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpeqPd([2]float64(a), [2]float64(b)))
}

func cmpeqPd(a [2]float64, b [2]float64) [2]float64


// CmpeqSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b' for equality, store the result in the lower element of 'dst',
// and copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] == b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpeq_sd'.
// Requires SSE2.
func CmpeqSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpeqSd([2]float64(a), [2]float64(b)))
}

func cmpeqSd(a [2]float64, b [2]float64) [2]float64


// CmpgePd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' for greater-than-or-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] >= b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpge_pd'.
// Requires SSE2.
func CmpgePd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpgePd([2]float64(a), [2]float64(b)))
}

func cmpgePd(a [2]float64, b [2]float64) [2]float64


// CmpgeSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b' for greater-than-or-equal, store the result in the lower
// element of 'dst', and copy the upper element from 'a' to the upper element
// of 'dst'. 
//
//		dst[63:0] := (a[63:0] >= b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpge_sd'.
// Requires SSE2.
func CmpgeSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpgeSd([2]float64(a), [2]float64(b)))
}

func cmpgeSd(a [2]float64, b [2]float64) [2]float64


// CmpgtEpi16: Compare packed 16-bit integers in 'a' and 'b' for greater-than,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] > b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTW'. Intrinsic: '_mm_cmpgt_epi16'.
// Requires SSE2.
func CmpgtEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpgtEpi16([16]byte(a), [16]byte(b)))
}

func cmpgtEpi16(a [16]byte, b [16]byte) [16]byte


// CmpgtEpi32: Compare packed 32-bit integers in 'a' and 'b' for greater-than,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] > b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTD'. Intrinsic: '_mm_cmpgt_epi32'.
// Requires SSE2.
func CmpgtEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpgtEpi32([16]byte(a), [16]byte(b)))
}

func cmpgtEpi32(a [16]byte, b [16]byte) [16]byte


// CmpgtEpi8: Compare packed 8-bit integers in 'a' and 'b' for greater-than,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] > b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTB'. Intrinsic: '_mm_cmpgt_epi8'.
// Requires SSE2.
func CmpgtEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpgtEpi8([16]byte(a), [16]byte(b)))
}

func cmpgtEpi8(a [16]byte, b [16]byte) [16]byte


// CmpgtPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' for greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] > b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpgt_pd'.
// Requires SSE2.
func CmpgtPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpgtPd([2]float64(a), [2]float64(b)))
}

func cmpgtPd(a [2]float64, b [2]float64) [2]float64


// CmpgtSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b' for greater-than, store the result in the lower element of
// 'dst', and copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] > b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpgt_sd'.
// Requires SSE2.
func CmpgtSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpgtSd([2]float64(a), [2]float64(b)))
}

func cmpgtSd(a [2]float64, b [2]float64) [2]float64


// CmplePd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' for less-than-or-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] <= b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmple_pd'.
// Requires SSE2.
func CmplePd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmplePd([2]float64(a), [2]float64(b)))
}

func cmplePd(a [2]float64, b [2]float64) [2]float64


// CmpleSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b' for less-than-or-equal, store the result in the lower element
// of 'dst', and copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] <= b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmple_sd'.
// Requires SSE2.
func CmpleSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpleSd([2]float64(a), [2]float64(b)))
}

func cmpleSd(a [2]float64, b [2]float64) [2]float64


// CmpltEpi16: Compare packed 16-bit integers in 'a' and 'b' for less-than, and
// store the results in 'dst'. Note: This intrinsic emits the pcmpgtw
// instruction with the order of the operands switched. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] < b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTW'. Intrinsic: '_mm_cmplt_epi16'.
// Requires SSE2.
func CmpltEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpltEpi16([16]byte(a), [16]byte(b)))
}

func cmpltEpi16(a [16]byte, b [16]byte) [16]byte


// CmpltEpi32: Compare packed 32-bit integers in 'a' and 'b' for less-than, and
// store the results in 'dst'. Note: This intrinsic emits the pcmpgtd
// instruction with the order of the operands switched. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] < b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTD'. Intrinsic: '_mm_cmplt_epi32'.
// Requires SSE2.
func CmpltEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpltEpi32([16]byte(a), [16]byte(b)))
}

func cmpltEpi32(a [16]byte, b [16]byte) [16]byte


// CmpltEpi8: Compare packed 8-bit integers in 'a' and 'b' for less-than, and
// store the results in 'dst'. Note: This intrinsic emits the pcmpgtb
// instruction with the order of the operands switched. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] < b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTB'. Intrinsic: '_mm_cmplt_epi8'.
// Requires SSE2.
func CmpltEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(cmpltEpi8([16]byte(a), [16]byte(b)))
}

func cmpltEpi8(a [16]byte, b [16]byte) [16]byte


// CmpltPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' for less-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] < b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmplt_pd'.
// Requires SSE2.
func CmpltPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpltPd([2]float64(a), [2]float64(b)))
}

func cmpltPd(a [2]float64, b [2]float64) [2]float64


// CmpltSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b' for less-than, store the result in the lower element of
// 'dst', and copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] < b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmplt_sd'.
// Requires SSE2.
func CmpltSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpltSd([2]float64(a), [2]float64(b)))
}

func cmpltSd(a [2]float64, b [2]float64) [2]float64


// CmpneqPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' for not-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] != b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpneq_pd'.
// Requires SSE2.
func CmpneqPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpneqPd([2]float64(a), [2]float64(b)))
}

func cmpneqPd(a [2]float64, b [2]float64) [2]float64


// CmpneqSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-equal, store the result in the lower element
// of 'dst', and copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] != b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpneq_sd'.
// Requires SSE2.
func CmpneqSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpneqSd([2]float64(a), [2]float64(b)))
}

func cmpneqSd(a [2]float64, b [2]float64) [2]float64


// CmpngePd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' for not-greater-than-or-equal, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := !(a[i+63:i] >= b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpnge_pd'.
// Requires SSE2.
func CmpngePd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpngePd([2]float64(a), [2]float64(b)))
}

func cmpngePd(a [2]float64, b [2]float64) [2]float64


// CmpngeSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-greater-than-or-equal, store the result in
// the lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. 
//
//		dst[63:0] := !(a[63:0] >= b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpnge_sd'.
// Requires SSE2.
func CmpngeSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpngeSd([2]float64(a), [2]float64(b)))
}

func cmpngeSd(a [2]float64, b [2]float64) [2]float64


// CmpngtPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' for not-greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := !(a[i+63:i] > b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpngt_pd'.
// Requires SSE2.
func CmpngtPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpngtPd([2]float64(a), [2]float64(b)))
}

func cmpngtPd(a [2]float64, b [2]float64) [2]float64


// CmpngtSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-greater-than, store the result in the lower
// element of 'dst', and copy the upper element from 'a' to the upper element
// of 'dst'. 
//
//		dst[63:0] := !(a[63:0] > b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpngt_sd'.
// Requires SSE2.
func CmpngtSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpngtSd([2]float64(a), [2]float64(b)))
}

func cmpngtSd(a [2]float64, b [2]float64) [2]float64


// CmpnlePd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' for not-less-than-or-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := !(a[i+63:i] <= b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpnle_pd'.
// Requires SSE2.
func CmpnlePd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpnlePd([2]float64(a), [2]float64(b)))
}

func cmpnlePd(a [2]float64, b [2]float64) [2]float64


// CmpnleSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, store the result in the
// lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. 
//
//		dst[63:0] := !(a[63:0] <= b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpnle_sd'.
// Requires SSE2.
func CmpnleSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpnleSd([2]float64(a), [2]float64(b)))
}

func cmpnleSd(a [2]float64, b [2]float64) [2]float64


// CmpnltPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' for not-less-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := !(a[i+63:i] < b[i+63:i]) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpnlt_pd'.
// Requires SSE2.
func CmpnltPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpnltPd([2]float64(a), [2]float64(b)))
}

func cmpnltPd(a [2]float64, b [2]float64) [2]float64


// CmpnltSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than, store the result in the lower
// element of 'dst', and copy the upper element from 'a' to the upper element
// of 'dst'. 
//
//		dst[63:0] := !(a[63:0] < b[63:0]) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpnlt_sd'.
// Requires SSE2.
func CmpnltSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpnltSd([2]float64(a), [2]float64(b)))
}

func cmpnltSd(a [2]float64, b [2]float64) [2]float64


// CmpordPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' to see if neither is NaN, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] != NaN AND b[i+63:i] != NaN) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpord_pd'.
// Requires SSE2.
func CmpordPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpordPd([2]float64(a), [2]float64(b)))
}

func cmpordPd(a [2]float64, b [2]float64) [2]float64


// CmpordSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, store the result in the
// lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. 
//
//		dst[63:0] := (a[63:0] != NaN AND b[63:0] != NaN) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpord_sd'.
// Requires SSE2.
func CmpordSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpordSd([2]float64(a), [2]float64(b)))
}

func cmpordSd(a [2]float64, b [2]float64) [2]float64


// CmpunordPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' to see if either is NaN, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] != NaN OR b[i+63:i] != NaN) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'CMPPD'. Intrinsic: '_mm_cmpunord_pd'.
// Requires SSE2.
func CmpunordPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpunordPd([2]float64(a), [2]float64(b)))
}

func cmpunordPd(a [2]float64, b [2]float64) [2]float64


// CmpunordSd: Compare the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, store the result in the
// lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. 
//
//		dst[63:0] := (a[63:0] != NaN OR b[63:0] != NaN) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//
// Instruction: 'CMPSD'. Intrinsic: '_mm_cmpunord_sd'.
// Requires SSE2.
func CmpunordSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(cmpunordSd([2]float64(a), [2]float64(b)))
}

func cmpunordSd(a [2]float64, b [2]float64) [2]float64


// ComieqSd: Compare the lower double-precision (64-bit) floating-point element
// in 'a' and 'b' for equality, and return the boolean result (0 or 1). 
//
//		RETURN ( a[63:0] == b[63:0] ) ? 1 : 0
//
// Instruction: 'COMISD'. Intrinsic: '_mm_comieq_sd'.
// Requires SSE2.
func ComieqSd(a x86.M128d, b x86.M128d) int {
	return int(comieqSd([2]float64(a), [2]float64(b)))
}

func comieqSd(a [2]float64, b [2]float64) int


// ComigeSd: Compare the lower double-precision (64-bit) floating-point element
// in 'a' and 'b' for greater-than-or-equal, and return the boolean result (0
// or 1). 
//
//		RETURN ( a[63:0] >= b[63:0] ) ? 1 : 0
//
// Instruction: 'COMISD'. Intrinsic: '_mm_comige_sd'.
// Requires SSE2.
func ComigeSd(a x86.M128d, b x86.M128d) int {
	return int(comigeSd([2]float64(a), [2]float64(b)))
}

func comigeSd(a [2]float64, b [2]float64) int


// ComigtSd: Compare the lower double-precision (64-bit) floating-point element
// in 'a' and 'b' for greater-than, and return the boolean result (0 or 1). 
//
//		RETURN ( a[63:0] > b[63:0] ) ? 1 : 0
//
// Instruction: 'COMISD'. Intrinsic: '_mm_comigt_sd'.
// Requires SSE2.
func ComigtSd(a x86.M128d, b x86.M128d) int {
	return int(comigtSd([2]float64(a), [2]float64(b)))
}

func comigtSd(a [2]float64, b [2]float64) int


// ComileSd: Compare the lower double-precision (64-bit) floating-point element
// in 'a' and 'b' for less-than-or-equal, and return the boolean result (0 or
// 1). 
//
//		RETURN ( a[63:0] <= b[63:0] ) ? 1 : 0
//
// Instruction: 'COMISD'. Intrinsic: '_mm_comile_sd'.
// Requires SSE2.
func ComileSd(a x86.M128d, b x86.M128d) int {
	return int(comileSd([2]float64(a), [2]float64(b)))
}

func comileSd(a [2]float64, b [2]float64) int


// ComiltSd: Compare the lower double-precision (64-bit) floating-point element
// in 'a' and 'b' for less-than, and return the boolean result (0 or 1). 
//
//		RETURN ( a[63:0] < b[63:0] ) ? 1 : 0
//
// Instruction: 'COMISD'. Intrinsic: '_mm_comilt_sd'.
// Requires SSE2.
func ComiltSd(a x86.M128d, b x86.M128d) int {
	return int(comiltSd([2]float64(a), [2]float64(b)))
}

func comiltSd(a [2]float64, b [2]float64) int


// ComineqSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for not-equal, and return the boolean result (0 or
// 1). 
//
//		RETURN ( a[63:0] != b[63:0] ) ? 1 : 0
//
// Instruction: 'COMISD'. Intrinsic: '_mm_comineq_sd'.
// Requires SSE2.
func ComineqSd(a x86.M128d, b x86.M128d) int {
	return int(comineqSd([2]float64(a), [2]float64(b)))
}

func comineqSd(a [2]float64, b [2]float64) int


// Cvtepi32Pd: Convert packed 32-bit integers in 'a' to packed double-precision
// (64-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			m := j*64
//			dst[m+63:m] := Convert_Int32_To_FP64(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTDQ2PD'. Intrinsic: '_mm_cvtepi32_pd'.
// Requires SSE2.
func Cvtepi32Pd(a x86.M128i) (dst x86.M128d) {
	return x86.M128d(cvtepi32Pd([16]byte(a)))
}

func cvtepi32Pd(a [16]byte) [2]float64


// Cvtepi32Ps: Convert packed 32-bit integers in 'a' to packed single-precision
// (32-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := Convert_Int32_To_FP32(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTDQ2PS'. Intrinsic: '_mm_cvtepi32_ps'.
// Requires SSE2.
func Cvtepi32Ps(a x86.M128i) (dst x86.M128) {
	return x86.M128(cvtepi32Ps([16]byte(a)))
}

func cvtepi32Ps(a [16]byte) [4]float32


// CvtpdEpi32: Convert packed double-precision (64-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32(a[k+63:k])
//		ENDFOR
//
// Instruction: 'CVTPD2DQ'. Intrinsic: '_mm_cvtpd_epi32'.
// Requires SSE2.
func CvtpdEpi32(a x86.M128d) (dst x86.M128i) {
	return x86.M128i(cvtpdEpi32([2]float64(a)))
}

func cvtpdEpi32(a [2]float64) [16]byte


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
func CvtpdPi32(a x86.M128d) (dst x86.M64) {
	return x86.M64(cvtpdPi32([2]float64(a)))
}

func cvtpdPi32(a [2]float64) x86.M64


// CvtpdPs: Convert packed double-precision (64-bit) floating-point elements in
// 'a' to packed single-precision (32-bit) floating-point elements, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_FP32(a[k+63:k])
//		ENDFOR
//
// Instruction: 'CVTPD2PS'. Intrinsic: '_mm_cvtpd_ps'.
// Requires SSE2.
func CvtpdPs(a x86.M128d) (dst x86.M128) {
	return x86.M128(cvtpdPs([2]float64(a)))
}

func cvtpdPs(a [2]float64) [4]float32


// Cvtpi32Pd: Convert packed 32-bit integers in 'a' to packed double-precision
// (64-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*32
//			m := j*64
//			dst[m+63:m] := Convert_Int32_To_FP64(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTPI2PD'. Intrinsic: '_mm_cvtpi32_pd'.
// Requires SSE2.
func Cvtpi32Pd(a x86.M64) (dst x86.M128d) {
	return x86.M128d(cvtpi32Pd(a))
}

func cvtpi32Pd(a x86.M64) [2]float64


// CvtpsEpi32: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTPS2DQ'. Intrinsic: '_mm_cvtps_epi32'.
// Requires SSE2.
func CvtpsEpi32(a x86.M128) (dst x86.M128i) {
	return x86.M128i(cvtpsEpi32([4]float32(a)))
}

func cvtpsEpi32(a [4]float32) [16]byte


// CvtpsPd: Convert packed single-precision (32-bit) floating-point elements in
// 'a' to packed double-precision (64-bit) floating-point elements, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 32*j
//			dst[i+63:i] := Convert_FP32_To_FP64(a[k+31:k])
//		ENDFOR
//
// Instruction: 'CVTPS2PD'. Intrinsic: '_mm_cvtps_pd'.
// Requires SSE2.
func CvtpsPd(a x86.M128) (dst x86.M128d) {
	return x86.M128d(cvtpsPd([4]float32(a)))
}

func cvtpsPd(a [4]float32) [2]float64


// CvtsdF64: Copy the lower double-precision (64-bit) floating-point element of
// 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVSD'. Intrinsic: '_mm_cvtsd_f64'.
// Requires SSE2.
func CvtsdF64(a x86.M128d) float64 {
	return float64(cvtsdF64([2]float64(a)))
}

func cvtsdF64(a [2]float64) float64


// CvtsdSi32: Convert the lower double-precision (64-bit) floating-point
// element in 'a' to a 32-bit integer, and store the result in 'dst'. 
//
//		dst[31:0] := Convert_FP64_To_Int32(a[63:0])
//
// Instruction: 'CVTSD2SI'. Intrinsic: '_mm_cvtsd_si32'.
// Requires SSE2.
func CvtsdSi32(a x86.M128d) int {
	return int(cvtsdSi32([2]float64(a)))
}

func cvtsdSi32(a [2]float64) int


// CvtsdSi64: Convert the lower double-precision (64-bit) floating-point
// element in 'a' to a 64-bit integer, and store the result in 'dst'. 
//
//		dst[63:0] := Convert_FP64_To_Int64(a[63:0])
//
// Instruction: 'CVTSD2SI'. Intrinsic: '_mm_cvtsd_si64'.
// Requires SSE2.
func CvtsdSi64(a x86.M128d) int64 {
	return int64(cvtsdSi64([2]float64(a)))
}

func cvtsdSi64(a [2]float64) int64


// CvtsdSi64x: Convert the lower double-precision (64-bit) floating-point
// element in 'a' to a 64-bit integer, and store the result in 'dst'. 
//
//		dst[63:0] := Convert_FP64_To_Int64(a[63:0])
//
// Instruction: 'CVTSD2SI'. Intrinsic: '_mm_cvtsd_si64x'.
// Requires SSE2.
func CvtsdSi64x(a x86.M128d) int64 {
	return int64(cvtsdSi64x([2]float64(a)))
}

func cvtsdSi64x(a [2]float64) int64


// CvtsdSs: Convert the lower double-precision (64-bit) floating-point element
// in 'b' to a single-precision (32-bit) floating-point element, store the
// result in the lower element of 'dst', and copy the upper element from 'a' to
// the upper element of 'dst'. 
//
//		dst[31:0] := Convert_FP64_To_FP32(b[63:0])
//		dst[127:32] := a[127:31]
//		dst[MAX:64] := 0
//
// Instruction: 'CVTSD2SS'. Intrinsic: '_mm_cvtsd_ss'.
// Requires SSE2.
func CvtsdSs(a x86.M128, b x86.M128d) (dst x86.M128) {
	return x86.M128(cvtsdSs([4]float32(a), [2]float64(b)))
}

func cvtsdSs(a [4]float32, b [2]float64) [4]float32


// Cvtsi128Si32: Copy the lower 32-bit integer in 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_cvtsi128_si32'.
// Requires SSE2.
func Cvtsi128Si32(a x86.M128i) int {
	return int(cvtsi128Si32([16]byte(a)))
}

func cvtsi128Si32(a [16]byte) int


// Cvtsi128Si64: Copy the lower 64-bit integer in 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtsi128_si64'.
// Requires SSE2.
func Cvtsi128Si64(a x86.M128i) int64 {
	return int64(cvtsi128Si64([16]byte(a)))
}

func cvtsi128Si64(a [16]byte) int64


// Cvtsi128Si64x: Copy the lower 64-bit integer in 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtsi128_si64x'.
// Requires SSE2.
func Cvtsi128Si64x(a x86.M128i) int64 {
	return int64(cvtsi128Si64x([16]byte(a)))
}

func cvtsi128Si64x(a [16]byte) int64


// Cvtsi32Sd: Convert the 32-bit integer 'b' to a double-precision (64-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := Convert_Int32_To_FP64(b[31:0])
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'CVTSI2SD'. Intrinsic: '_mm_cvtsi32_sd'.
// Requires SSE2.
func Cvtsi32Sd(a x86.M128d, b int) (dst x86.M128d) {
	return x86.M128d(cvtsi32Sd([2]float64(a), b))
}

func cvtsi32Sd(a [2]float64, b int) [2]float64


// Cvtsi32Si128: Copy 32-bit integer 'a' to the lower elements of 'dst', and
// zero the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[127:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_mm_cvtsi32_si128'.
// Requires SSE2.
func Cvtsi32Si128(a int) (dst x86.M128i) {
	return x86.M128i(cvtsi32Si128(a))
}

func cvtsi32Si128(a int) [16]byte


// Cvtsi64Sd: Convert the 64-bit integer 'b' to a double-precision (64-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := Convert_Int64_To_FP64(b[63:0])
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'CVTSI2SD'. Intrinsic: '_mm_cvtsi64_sd'.
// Requires SSE2.
func Cvtsi64Sd(a x86.M128d, b int64) (dst x86.M128d) {
	return x86.M128d(cvtsi64Sd([2]float64(a), b))
}

func cvtsi64Sd(a [2]float64, b int64) [2]float64


// Cvtsi64Si128: Copy 64-bit integer 'a' to the lower element of 'dst', and
// zero the upper element. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtsi64_si128'.
// Requires SSE2.
func Cvtsi64Si128(a int64) (dst x86.M128i) {
	return x86.M128i(cvtsi64Si128(a))
}

func cvtsi64Si128(a int64) [16]byte


// Cvtsi64xSd: Convert the 64-bit integer 'b' to a double-precision (64-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := Convert_Int64_To_FP64(b[63:0])
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'CVTSI2SD'. Intrinsic: '_mm_cvtsi64x_sd'.
// Requires SSE2.
func Cvtsi64xSd(a x86.M128d, b int64) (dst x86.M128d) {
	return x86.M128d(cvtsi64xSd([2]float64(a), b))
}

func cvtsi64xSd(a [2]float64, b int64) [2]float64


// Cvtsi64xSi128: Copy 64-bit integer 'a' to the lower element of 'dst', and
// zero the upper element. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_cvtsi64x_si128'.
// Requires SSE2.
func Cvtsi64xSi128(a int64) (dst x86.M128i) {
	return x86.M128i(cvtsi64xSi128(a))
}

func cvtsi64xSi128(a int64) [16]byte


// CvtssSd: Convert the lower single-precision (32-bit) floating-point element
// in 'b' to a double-precision (64-bit) floating-point element, store the
// result in the lower element of 'dst', and copy the upper element from 'a' to
// the upper element of 'dst'. 
//
//		dst[63:0] := Convert_FP32_To_FP64(b[31:0])
//		dst[127:64] := a[127:64]
//		dst[MAX:64] := 0
//
// Instruction: 'CVTSS2SD'. Intrinsic: '_mm_cvtss_sd'.
// Requires SSE2.
func CvtssSd(a x86.M128d, b x86.M128) (dst x86.M128d) {
	return x86.M128d(cvtssSd([2]float64(a), [4]float32(b)))
}

func cvtssSd(a [2]float64, b [4]float32) [2]float64


// CvttpdEpi32: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 32-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32_Truncate(a[k+63:k])
//		ENDFOR
//
// Instruction: 'CVTTPD2DQ'. Intrinsic: '_mm_cvttpd_epi32'.
// Requires SSE2.
func CvttpdEpi32(a x86.M128d) (dst x86.M128i) {
	return x86.M128i(cvttpdEpi32([2]float64(a)))
}

func cvttpdEpi32(a [2]float64) [16]byte


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
func CvttpdPi32(a x86.M128d) (dst x86.M64) {
	return x86.M64(cvttpdPi32([2]float64(a)))
}

func cvttpdPi32(a [2]float64) x86.M64


// CvttpsEpi32: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 32-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32_Truncate(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTTPS2DQ'. Intrinsic: '_mm_cvttps_epi32'.
// Requires SSE2.
func CvttpsEpi32(a x86.M128) (dst x86.M128i) {
	return x86.M128i(cvttpsEpi32([4]float32(a)))
}

func cvttpsEpi32(a [4]float32) [16]byte


// CvttsdSi32: Convert the lower double-precision (64-bit) floating-point
// element in 'a' to a 32-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP64_To_Int32_Truncate(a[63:0])
//
// Instruction: 'CVTTSD2SI'. Intrinsic: '_mm_cvttsd_si32'.
// Requires SSE2.
func CvttsdSi32(a x86.M128d) int {
	return int(cvttsdSi32([2]float64(a)))
}

func cvttsdSi32(a [2]float64) int


// CvttsdSi64: Convert the lower double-precision (64-bit) floating-point
// element in 'a' to a 64-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[63:0] := Convert_FP64_To_Int64_Truncate(a[63:0])
//
// Instruction: 'CVTTSD2SI'. Intrinsic: '_mm_cvttsd_si64'.
// Requires SSE2.
func CvttsdSi64(a x86.M128d) int64 {
	return int64(cvttsdSi64([2]float64(a)))
}

func cvttsdSi64(a [2]float64) int64


// CvttsdSi64x: Convert the lower double-precision (64-bit) floating-point
// element in 'a' to a 64-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[63:0] := Convert_FP64_To_Int64_Truncate(a[63:0])
//
// Instruction: 'CVTTSD2SI'. Intrinsic: '_mm_cvttsd_si64x'.
// Requires SSE2.
func CvttsdSi64x(a x86.M128d) int64 {
	return int64(cvttsdSi64x([2]float64(a)))
}

func cvttsdSi64x(a [2]float64) int64


// DivPd: Divide packed double-precision (64-bit) floating-point elements in
// 'a' by packed elements in 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			dst[i+63:i] := a[i+63:i] / b[i+63:i]
//		ENDFOR
//
// Instruction: 'DIVPD'. Intrinsic: '_mm_div_pd'.
// Requires SSE2.
func DivPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(divPd([2]float64(a), [2]float64(b)))
}

func divPd(a [2]float64, b [2]float64) [2]float64


// DivSd: Divide the lower double-precision (64-bit) floating-point element in
// 'a' by the lower double-precision (64-bit) floating-point element in 'b',
// store the result in the lower element of 'dst', and copy the upper element
// from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := a[63:0] 0 b[63:0]
//		dst[127:64] := a[127:64]
//
// Instruction: 'DIVSD'. Intrinsic: '_mm_div_sd'.
// Requires SSE2.
func DivSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(divSd([2]float64(a), [2]float64(b)))
}

func divSd(a [2]float64, b [2]float64) [2]float64


// ExtractEpi16: Extract a 16-bit integer from 'a', selected with 'imm8', and
// store the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[127:0] >> (imm8[2:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_mm_extract_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func ExtractEpi16(a x86.M128i, imm8 byte) int {
	return int(extractEpi16([16]byte(a), imm8))
}

func extractEpi16(a [16]byte, imm8 byte) int


// InsertEpi16: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst'
// at the location specified by 'imm8'. 
//
//		dst[127:0] := a[127:0]
//		sel := imm8[2:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_mm_insert_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func InsertEpi16(a x86.M128i, i int, imm8 byte) (dst x86.M128i) {
	return x86.M128i(insertEpi16([16]byte(a), i, imm8))
}

func insertEpi16(a [16]byte, i int, imm8 byte) [16]byte


// Lfence: Perform a serializing operation on all load-from-memory instructions
// that were issued prior to this instruction. Guarantees that every load
// instruction that precedes, in program order, is globally visible before any
// load instruction which follows the fence in program order. 
//
//		
//
// Instruction: 'LFENCE'. Intrinsic: '_mm_lfence'.
// Requires SSE2.
func Lfence()  {
	lfence()
}

func lfence() 


// Skipped: _mm_load_pd. Contains pointer parameter.


// Skipped: _mm_load_pd1. Contains pointer parameter.


// Skipped: _mm_load_sd. Contains pointer parameter.


// LoadSi128: Load 128-bits of integer data from memory into 'dst'. 
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[127:0] := MEM[mem_addr+127:mem_addr]
//
// Instruction: 'MOVDQA'. Intrinsic: '_mm_load_si128'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func LoadSi128(mem_addr *x86.M128iConst) (dst x86.M128i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128i{}
}

// Skipped: _mm_load1_pd. Contains pointer parameter.


// Skipped: _mm_loadh_pd. Contains pointer parameter.


// LoadlEpi64: Load 64-bit integer from memory into the first element of 'dst'. 
//
//		dst[63:0] := MEM[mem_addr+63:mem_addr]
//		dst[MAX:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_loadl_epi64'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func LoadlEpi64(mem_addr *x86.M128iConst) (dst x86.M128i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128i{}
}

// Skipped: _mm_loadl_pd. Contains pointer parameter.


// Skipped: _mm_loadr_pd. Contains pointer parameter.


// Skipped: _mm_loadu_pd. Contains pointer parameter.


// LoaduSi128: Load 128-bits of integer data from memory into 'dst'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[127:0] := MEM[mem_addr+127:mem_addr]
//
// Instruction: 'MOVDQU'. Intrinsic: '_mm_loadu_si128'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func LoaduSi128(mem_addr *x86.M128iConst) (dst x86.M128i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128i{}
}

// MaddEpi16: Multiply packed signed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			st[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//		ENDFOR
//
// Instruction: 'PMADDWD'. Intrinsic: '_mm_madd_epi16'.
// Requires SSE2.
func MaddEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(maddEpi16([16]byte(a), [16]byte(b)))
}

func maddEpi16(a [16]byte, b [16]byte) [16]byte


// MaskmoveuSi128: Conditionally store 8-bit integer elements from 'a' into
// memory using 'mask' (elements are not stored when the highest bit is not set
// in the corresponding element) and a non-temporal memory hint. 'mem_addr'
// does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF mask[i+7]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'MASKMOVDQU'. Intrinsic: '_mm_maskmoveu_si128'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskmoveuSi128(a x86.M128i, mask x86.M128i, mem_addr *byte)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// MaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSW'. Intrinsic: '_mm_max_epi16'.
// Requires SSE2.
func MaxEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(maxEpi16([16]byte(a), [16]byte(b)))
}

func maxEpi16(a [16]byte, b [16]byte) [16]byte


// MaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUB'. Intrinsic: '_mm_max_epu8'.
// Requires SSE2.
func MaxEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(maxEpu8([16]byte(a), [16]byte(b)))
}

func maxEpu8(a [16]byte, b [16]byte) [16]byte


// MaxPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := MAX(a[i+63:i], b[i+63:i])
//		ENDFOR
//
// Instruction: 'MAXPD'. Intrinsic: '_mm_max_pd'.
// Requires SSE2.
func MaxPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(maxPd([2]float64(a), [2]float64(b)))
}

func maxPd(a [2]float64, b [2]float64) [2]float64


// MaxSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b', store the maximum value in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := MAX(a[63:0], b[63:0])
//		dst[127:64] := a[127:64]
//
// Instruction: 'MAXSD'. Intrinsic: '_mm_max_sd'.
// Requires SSE2.
func MaxSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(maxSd([2]float64(a), [2]float64(b)))
}

func maxSd(a [2]float64, b [2]float64) [2]float64


// Mfence: Perform a serializing operation on all load-from-memory and
// store-to-memory instructions that were issued prior to this instruction.
// Guarantees that every memory access that precedes, in program order, the
// memory fence instruction is globally visible before any memory instruction
// which follows the fence in program order. 
//
//		
//
// Instruction: 'MFENCE'. Intrinsic: '_mm_mfence'.
// Requires SSE2.
func Mfence()  {
	mfence()
}

func mfence() 


// MinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSW'. Intrinsic: '_mm_min_epi16'.
// Requires SSE2.
func MinEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(minEpi16([16]byte(a), [16]byte(b)))
}

func minEpi16(a [16]byte, b [16]byte) [16]byte


// MinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUB'. Intrinsic: '_mm_min_epu8'.
// Requires SSE2.
func MinEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(minEpu8([16]byte(a), [16]byte(b)))
}

func minEpu8(a [16]byte, b [16]byte) [16]byte


// MinPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := MIN(a[i+63:i], b[i+63:i])
//		ENDFOR
//
// Instruction: 'MINPD'. Intrinsic: '_mm_min_pd'.
// Requires SSE2.
func MinPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(minPd([2]float64(a), [2]float64(b)))
}

func minPd(a [2]float64, b [2]float64) [2]float64


// MinSd: Compare the lower double-precision (64-bit) floating-point elements
// in 'a' and 'b', store the minimum value in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := MIN(a[63:0], b[63:0])
//		dst[127:64] := a[127:64]
//
// Instruction: 'MINSD'. Intrinsic: '_mm_min_sd'.
// Requires SSE2.
func MinSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(minSd([2]float64(a), [2]float64(b)))
}

func minSd(a [2]float64, b [2]float64) [2]float64


// MoveEpi64: Copy the lower 64-bit integer in 'a' to the lower element of
// 'dst', and zero the upper element. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_move_epi64'.
// Requires SSE2.
func MoveEpi64(a x86.M128i) (dst x86.M128i) {
	return x86.M128i(moveEpi64([16]byte(a)))
}

func moveEpi64(a [16]byte) [16]byte


// MoveSd: Move the lower double-precision (64-bit) floating-point element from
// 'b' to the lower element of 'dst', and copy the upper element from 'a' to
// the upper element of 'dst'. 
//
//		dst[63:0] := b[63:0]
//		dst[127:64] := a[127:64]
//
// Instruction: 'MOVSD'. Intrinsic: '_mm_move_sd'.
// Requires SSE2.
func MoveSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(moveSd([2]float64(a), [2]float64(b)))
}

func moveSd(a [2]float64, b [2]float64) [2]float64


// MovemaskEpi8: Create mask from the most significant bit of each 8-bit
// element in 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//		dst[MAX:16] := 0
//
// Instruction: 'PMOVMSKB'. Intrinsic: '_mm_movemask_epi8'.
// Requires SSE2.
func MovemaskEpi8(a x86.M128i) int {
	return int(movemaskEpi8([16]byte(a)))
}

func movemaskEpi8(a [16]byte) int


// MovemaskPd: Set each bit of mask 'dst' based on the most significant bit of
// the corresponding packed double-precision (64-bit) floating-point element in
// 'a'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF a[i+63]
//				dst[j] := 1
//			ELSE
//				dst[j] := 0
//			FI
//		ENDFOR
//		dst[MAX:2] := 0
//
// Instruction: 'MOVMSKPD'. Intrinsic: '_mm_movemask_pd'.
// Requires SSE2.
func MovemaskPd(a x86.M128d) int {
	return int(movemaskPd([2]float64(a)))
}

func movemaskPd(a [2]float64) int


// Movepi64Pi64: Copy the lower 64-bit integer in 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//
// Instruction: 'MOVDQ2Q'. Intrinsic: '_mm_movepi64_pi64'.
// Requires SSE2.
func Movepi64Pi64(a x86.M128i) (dst x86.M64) {
	return x86.M64(movepi64Pi64([16]byte(a)))
}

func movepi64Pi64(a [16]byte) x86.M64


// Movpi64Epi64: Copy the 64-bit integer 'a' to the lower element of 'dst', and
// zero the upper element. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := 0
//
// Instruction: 'MOVQ2DQ'. Intrinsic: '_mm_movpi64_epi64'.
// Requires SSE2.
func Movpi64Epi64(a x86.M64) (dst x86.M128i) {
	return x86.M128i(movpi64Epi64(a))
}

func movpi64Epi64(a x86.M64) [16]byte


// MulEpu32: Multiply the low unsigned 32-bit integers from each packed 64-bit
// element in 'a' and 'b', and store the unsigned 64-bit results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//
// Instruction: 'PMULUDQ'. Intrinsic: '_mm_mul_epu32'.
// Requires SSE2.
func MulEpu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(mulEpu32([16]byte(a), [16]byte(b)))
}

func mulEpu32(a [16]byte, b [16]byte) [16]byte


// MulPd: Multiply packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//
// Instruction: 'MULPD'. Intrinsic: '_mm_mul_pd'.
// Requires SSE2.
func MulPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(mulPd([2]float64(a), [2]float64(b)))
}

func mulPd(a [2]float64, b [2]float64) [2]float64


// MulSd: Multiply the lower double-precision (64-bit) floating-point element
// in 'a' and 'b', store the result in the lower element of 'dst', and copy the
// upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := a[63:0] * b[63:0]
//		dst[127:64] := a[127:64]
//
// Instruction: 'MULSD'. Intrinsic: '_mm_mul_sd'.
// Requires SSE2.
func MulSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(mulSd([2]float64(a), [2]float64(b)))
}

func mulSd(a [2]float64, b [2]float64) [2]float64


// MulSu32: Multiply the low unsigned 32-bit integers from 'a' and 'b', and
// store the unsigned 64-bit result in 'dst'. 
//
//		dst[63:0] := a[31:0] * b[31:0]
//
// Instruction: 'PMULUDQ'. Intrinsic: '_mm_mul_su32'.
// Requires SSE2.
func MulSu32(a x86.M64, b x86.M64) (dst x86.M64) {
	return x86.M64(mulSu32(a, b))
}

func mulSu32(a x86.M64, b x86.M64) x86.M64


// MulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the high 16 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHW'. Intrinsic: '_mm_mulhi_epi16'.
// Requires SSE2.
func MulhiEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(mulhiEpi16([16]byte(a), [16]byte(b)))
}

func mulhiEpi16(a [16]byte, b [16]byte) [16]byte


// MulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHUW'. Intrinsic: '_mm_mulhi_epu16'.
// Requires SSE2.
func MulhiEpu16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(mulhiEpu16([16]byte(a), [16]byte(b)))
}

func mulhiEpu16(a [16]byte, b [16]byte) [16]byte


// MulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the low 16 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[15:0]
//		ENDFOR
//
// Instruction: 'PMULLW'. Intrinsic: '_mm_mullo_epi16'.
// Requires SSE2.
func MulloEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(mulloEpi16([16]byte(a), [16]byte(b)))
}

func mulloEpi16(a [16]byte, b [16]byte) [16]byte


// OrPd: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//		ENDFOR
//
// Instruction: 'ORPD'. Intrinsic: '_mm_or_pd'.
// Requires SSE2.
func OrPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(orPd([2]float64(a), [2]float64(b)))
}

func orPd(a [2]float64, b [2]float64) [2]float64


// OrSi128: Compute the bitwise OR of 128 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[127:0] := (a[127:0] OR b[127:0])
//
// Instruction: 'POR'. Intrinsic: '_mm_or_si128'.
// Requires SSE2.
func OrSi128(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(orSi128([16]byte(a), [16]byte(b)))
}

func orSi128(a [16]byte, b [16]byte) [16]byte


// PacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed 8-bit
// integers using signed saturation, and store the results in 'dst'. 
//
//		dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//
// Instruction: 'PACKSSWB'. Intrinsic: '_mm_packs_epi16'.
// Requires SSE2.
func PacksEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(packsEpi16([16]byte(a), [16]byte(b)))
}

func packsEpi16(a [16]byte, b [16]byte) [16]byte


// PacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed 16-bit
// integers using signed saturation, and store the results in 'dst'. 
//
//		dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//
// Instruction: 'PACKSSDW'. Intrinsic: '_mm_packs_epi32'.
// Requires SSE2.
func PacksEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(packsEpi32([16]byte(a), [16]byte(b)))
}

func packsEpi32(a [16]byte, b [16]byte) [16]byte


// PackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed 8-bit
// integers using unsigned saturation, and store the results in 'dst'. 
//
//		dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//
// Instruction: 'PACKUSWB'. Intrinsic: '_mm_packus_epi16'.
// Requires SSE2.
func PackusEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(packusEpi16([16]byte(a), [16]byte(b)))
}

func packusEpi16(a [16]byte, b [16]byte) [16]byte


// Pause: Provide a hint to the processor that the code sequence is a spin-wait
// loop. This can help improve the performance and power consumption of
// spin-wait loops. 
//
//		
//
// Instruction: 'PAUSE'. Intrinsic: '_mm_pause'.
// Requires SSE2.
func Pause()  {
	pause()
}

func pause() 


// SadEpu8: Compute the absolute differences of packed unsigned 8-bit integers
// in 'a' and 'b', then horizontally sum each consecutive 8 differences to
// produce two unsigned 16-bit integers, and pack these unsigned 16-bit
// integers in the low 16 bits of 64-bit elements in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+15:i] := tmp[i+7:i] + tmp[i+15:i+8] + tmp[i+23:i+16] + tmp[i+31:i+24] +
//			               tmp[i+39:i+32] + tmp[i+47:i+40] + tmp[i+55:i+48] + tmp[i+63:i+56]
//			dst[i+63:i+16] := 0
//		ENDFOR
//
// Instruction: 'PSADBW'. Intrinsic: '_mm_sad_epu8'.
// Requires SSE2.
func SadEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(sadEpu8([16]byte(a), [16]byte(b)))
}

func sadEpu8(a [16]byte, b [16]byte) [16]byte


// SetEpi16: Set packed 16-bit integers in 'dst' with the supplied values. 
//
//		dst[15:0] := e0
//		dst[31:16] := e1
//		dst[47:32] := e2
//		dst[63:48] := e3
//		dst[79:64] := e4
//		dst[95:80] := e5
//		dst[111:96] := e6
//		dst[127:112] := e7
//
// Instruction: '...'. Intrinsic: '_mm_set_epi16'.
// Requires SSE2.
func SetEpi16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) (dst x86.M128i) {
	return x86.M128i(setEpi16(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setEpi16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [16]byte


// SetEpi32: Set packed 32-bit integers in 'dst' with the supplied values. 
//
//		dst[31:0] := e0
//		dst[63:32] := e1
//		dst[95:64] := e2
//		dst[127:96] := e3
//
// Instruction: '...'. Intrinsic: '_mm_set_epi32'.
// Requires SSE2.
func SetEpi32(e3 int, e2 int, e1 int, e0 int) (dst x86.M128i) {
	return x86.M128i(setEpi32(e3, e2, e1, e0))
}

func setEpi32(e3 int, e2 int, e1 int, e0 int) [16]byte


// SetEpi64: Set packed 64-bit integers in 'dst' with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//
// Instruction: '...'. Intrinsic: '_mm_set_epi64'.
// Requires SSE2.
func SetEpi64(e1 x86.M64, e0 x86.M64) (dst x86.M128i) {
	return x86.M128i(setEpi64(e1, e0))
}

func setEpi64(e1 x86.M64, e0 x86.M64) [16]byte


// SetEpi64x: Set packed 64-bit integers in 'dst' with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//
// Instruction: '...'. Intrinsic: '_mm_set_epi64x'.
// Requires SSE2.
func SetEpi64x(e1 int64, e0 int64) (dst x86.M128i) {
	return x86.M128i(setEpi64x(e1, e0))
}

func setEpi64x(e1 int64, e0 int64) [16]byte


// SetEpi8: Set packed 8-bit integers in 'dst' with the supplied values in
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
//		dst[71:64] := e8
//		dst[79:72] := e9
//		dst[87:80] := e10
//		dst[95:88] := e11
//		dst[103:96] := e12
//		dst[111:104] := e13
//		dst[119:112] := e14
//		dst[127:120] := e15
//
// Instruction: '...'. Intrinsic: '_mm_set_epi8'.
// Requires SSE2.
func SetEpi8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) (dst x86.M128i) {
	return x86.M128i(setEpi8(e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func setEpi8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [16]byte


// SetPd: Set packed double-precision (64-bit) floating-point elements in 'dst'
// with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//
// Instruction: '...'. Intrinsic: '_mm_set_pd'.
// Requires SSE2.
func SetPd(e1 float64, e0 float64) (dst x86.M128d) {
	return x86.M128d(setPd(e1, e0))
}

func setPd(e1 float64, e0 float64) [2]float64


// SetPd1: Broadcast double-precision (64-bit) floating-point value 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set_pd1'.
// Requires SSE2.
func SetPd1(a float64) (dst x86.M128d) {
	return x86.M128d(setPd1(a))
}

func setPd1(a float64) [2]float64


// SetSd: Copy double-precision (64-bit) floating-point element 'a' to the
// lower element of 'dst', and zero the upper element. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := 0
//
// Instruction: '...'. Intrinsic: '_mm_set_sd'.
// Requires SSE2.
func SetSd(a float64) (dst x86.M128d) {
	return x86.M128d(setSd(a))
}

func setSd(a float64) [2]float64


// Set1Epi16: Broadcast 16-bit integer 'a' to all all elements of 'dst'. This
// intrinsic may generate 'vpbroadcastw'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_epi16'.
// Requires SSE2.
func Set1Epi16(a int16) (dst x86.M128i) {
	return x86.M128i(set1Epi16(a))
}

func set1Epi16(a int16) [16]byte


// Set1Epi32: Broadcast 32-bit integer 'a' to all elements of 'dst'. This
// intrinsic may generate 'vpbroadcastd'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_epi32'.
// Requires SSE2.
func Set1Epi32(a int) (dst x86.M128i) {
	return x86.M128i(set1Epi32(a))
}

func set1Epi32(a int) [16]byte


// Set1Epi64: Broadcast 64-bit integer 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_epi64'.
// Requires SSE2.
func Set1Epi64(a x86.M64) (dst x86.M128i) {
	return x86.M128i(set1Epi64(a))
}

func set1Epi64(a x86.M64) [16]byte


// Set1Epi64x: Broadcast 64-bit integer 'a' to all elements of 'dst'. This
// intrinsic may generate the 'vpbroadcastq'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_epi64x'.
// Requires SSE2.
func Set1Epi64x(a int64) (dst x86.M128i) {
	return x86.M128i(set1Epi64x(a))
}

func set1Epi64x(a int64) [16]byte


// Set1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst'. This
// intrinsic may generate 'vpbroadcastb'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_epi8'.
// Requires SSE2.
func Set1Epi8(a byte) (dst x86.M128i) {
	return x86.M128i(set1Epi8(a))
}

func set1Epi8(a byte) [16]byte


// Set1Pd: Broadcast double-precision (64-bit) floating-point value 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_pd'.
// Requires SSE2.
func Set1Pd(a float64) (dst x86.M128d) {
	return x86.M128d(set1Pd(a))
}

func set1Pd(a float64) [2]float64


// SetrEpi16: Set packed 16-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[15:0] := e7
//		dst[31:16] := e6
//		dst[47:32] := e5
//		dst[63:48] := e4
//		dst[79:64] := e3
//		dst[95:80] := e2
//		dst[111:96] := e1
//		dst[127:112] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_epi16'.
// Requires SSE2.
func SetrEpi16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) (dst x86.M128i) {
	return x86.M128i(setrEpi16(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrEpi16(e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [16]byte


// SetrEpi32: Set packed 32-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[31:0] := e3
//		dst[63:32] := e2
//		dst[95:64] := e1
//		dst[127:96] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_epi32'.
// Requires SSE2.
func SetrEpi32(e3 int, e2 int, e1 int, e0 int) (dst x86.M128i) {
	return x86.M128i(setrEpi32(e3, e2, e1, e0))
}

func setrEpi32(e3 int, e2 int, e1 int, e0 int) [16]byte


// SetrEpi64: Set packed 64-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[63:0] := e1
//		dst[127:64] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_epi64'.
// Requires SSE2.
func SetrEpi64(e1 x86.M64, e0 x86.M64) (dst x86.M128i) {
	return x86.M128i(setrEpi64(e1, e0))
}

func setrEpi64(e1 x86.M64, e0 x86.M64) [16]byte


// SetrEpi8: Set packed 8-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[7:0] := e15
//		dst[15:8] := e14
//		dst[23:16] := e13
//		dst[31:24] := e12
//		dst[39:32] := e11
//		dst[47:40] := e10
//		dst[55:48] := e9
//		dst[63:56] := e8
//		dst[71:64] := e7
//		dst[79:72] := e6
//		dst[87:80] := e5
//		dst[95:88] := e4
//		dst[103:96] := e3
//		dst[111:104] := e2
//		dst[119:112] := e1
//		dst[127:120] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_epi8'.
// Requires SSE2.
func SetrEpi8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) (dst x86.M128i) {
	return x86.M128i(setrEpi8(e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrEpi8(e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [16]byte


// SetrPd: Set packed double-precision (64-bit) floating-point elements in
// 'dst' with the supplied values in reverse order. 
//
//		dst[63:0] := e1
//		dst[127:64] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_pd'.
// Requires SSE2.
func SetrPd(e1 float64, e0 float64) (dst x86.M128d) {
	return x86.M128d(setrPd(e1, e0))
}

func setrPd(e1 float64, e0 float64) [2]float64


// SetzeroPd: Return vector of type __m128d with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'XORPD'. Intrinsic: '_mm_setzero_pd'.
// Requires SSE2.
func SetzeroPd() (dst x86.M128d) {
	return x86.M128d(setzeroPd())
}

func setzeroPd() [2]float64


// SetzeroSi128: Return vector of type __m128i with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'PXOR'. Intrinsic: '_mm_setzero_si128'.
// Requires SSE2.
func SetzeroSi128() (dst x86.M128i) {
	return x86.M128i(setzeroSi128())
}

func setzeroSi128() [16]byte


// ShuffleEpi32: Shuffle 32-bit integers in 'a' using the control in 'imm8',
// and store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[31:0] := src[31:0]
//			1:	tmp[31:0] := src[63:32]
//			2:	tmp[31:0] := src[95:64]
//			3:	tmp[31:0] := src[127:96]
//			ESAC
//			RETURN tmp[31:0]
//		}
//		
//		dst[31:0] := SELECT4(a[127:0], imm8[1:0])
//		dst[63:32] := SELECT4(a[127:0], imm8[3:2])
//		dst[95:64] := SELECT4(a[127:0], imm8[5:4])
//		dst[127:96] := SELECT4(a[127:0], imm8[7:6])
//
// Instruction: 'PSHUFD'. Intrinsic: '_mm_shuffle_epi32'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func ShuffleEpi32(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(shuffleEpi32([16]byte(a), imm8))
}

func shuffleEpi32(a [16]byte, imm8 byte) [16]byte


// ShufflePd: Shuffle double-precision (64-bit) floating-point elements using
// the control in 'imm8', and store the results in 'dst'. 
//
//		dst[63:0] := (imm8[0] == 0) ? a[63:0] : a[127:64]
//		dst[127:64] := (imm8[1] == 0) ? b[63:0] : b[127:64]
//
// Instruction: 'SHUFPD'. Intrinsic: '_mm_shuffle_pd'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func ShufflePd(a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	return x86.M128d(shufflePd([2]float64(a), [2]float64(b), imm8))
}

func shufflePd(a [2]float64, b [2]float64, imm8 byte) [2]float64


// ShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of 'a' using the
// control in 'imm8'. Store the results in the high 64 bits of 'dst', with the
// low 64 bits being copied from from 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//		dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//
// Instruction: 'PSHUFHW'. Intrinsic: '_mm_shufflehi_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func ShufflehiEpi16(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(shufflehiEpi16([16]byte(a), imm8))
}

func shufflehiEpi16(a [16]byte, imm8 byte) [16]byte


// ShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of 'a' using the
// control in 'imm8'. Store the results in the low 64 bits of 'dst', with the
// high 64 bits being copied from from 'a' to 'dst'. 
//
//		dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		dst[127:64] := a[127:64]
//
// Instruction: 'PSHUFLW'. Intrinsic: '_mm_shufflelo_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func ShuffleloEpi16(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(shuffleloEpi16([16]byte(a), imm8))
}

func shuffleloEpi16(a [16]byte, imm8 byte) [16]byte


// SllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLW'. Intrinsic: '_mm_sll_epi16'.
// Requires SSE2.
func SllEpi16(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(sllEpi16([16]byte(a), [16]byte(count)))
}

func sllEpi16(a [16]byte, count [16]byte) [16]byte


// SllEpi32: Shift packed 32-bit integers in 'a' left by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLD'. Intrinsic: '_mm_sll_epi32'.
// Requires SSE2.
func SllEpi32(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(sllEpi32([16]byte(a), [16]byte(count)))
}

func sllEpi32(a [16]byte, count [16]byte) [16]byte


// SllEpi64: Shift packed 64-bit integers in 'a' left by 'count' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF count[63:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] << count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLQ'. Intrinsic: '_mm_sll_epi64'.
// Requires SSE2.
func SllEpi64(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(sllEpi64([16]byte(a), [16]byte(count)))
}

func sllEpi64(a [16]byte, count [16]byte) [16]byte


// SlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLW'. Intrinsic: '_mm_slli_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SlliEpi16(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(slliEpi16([16]byte(a), imm8))
}

func slliEpi16(a [16]byte, imm8 byte) [16]byte


// SlliEpi32: Shift packed 32-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLD'. Intrinsic: '_mm_slli_epi32'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SlliEpi32(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(slliEpi32([16]byte(a), imm8))
}

func slliEpi32(a [16]byte, imm8 byte) [16]byte


// SlliEpi64: Shift packed 64-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF imm8[7:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] << imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSLLQ'. Intrinsic: '_mm_slli_epi64'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SlliEpi64(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(slliEpi64([16]byte(a), imm8))
}

func slliEpi64(a [16]byte, imm8 byte) [16]byte


// SlliSi128: Shift 'a' left by 'imm8' bytes while shifting in zeros, and store
// the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] << (tmp*8)
//
// Instruction: 'PSLLDQ'. Intrinsic: '_mm_slli_si128'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SlliSi128(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(slliSi128([16]byte(a), imm8))
}

func slliSi128(a [16]byte, imm8 byte) [16]byte


// SqrtPd: Compute the square root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SQRT(a[i+63:i])
//		ENDFOR
//
// Instruction: 'SQRTPD'. Intrinsic: '_mm_sqrt_pd'.
// Requires SSE2.
func SqrtPd(a x86.M128d) (dst x86.M128d) {
	return x86.M128d(sqrtPd([2]float64(a)))
}

func sqrtPd(a [2]float64) [2]float64


// SqrtSd: Compute the square root of the lower double-precision (64-bit)
// floating-point element in 'a', store the result in the lower element of
// 'dst', and copy the upper element from 'b' to the upper element of 'dst'. 
//
//		dst[63:0] := SQRT(a[63:0])
//		dst[127:64] := b[127:64]
//
// Instruction: 'SQRTSD'. Intrinsic: '_mm_sqrt_sd'.
// Requires SSE2.
func SqrtSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(sqrtSd([2]float64(a), [2]float64(b)))
}

func sqrtSd(a [2]float64, b [2]float64) [2]float64


// SraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAW'. Intrinsic: '_mm_sra_epi16'.
// Requires SSE2.
func SraEpi16(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(sraEpi16([16]byte(a), [16]byte(count)))
}

func sraEpi16(a [16]byte, count [16]byte) [16]byte


// SraEpi32: Shift packed 32-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAD'. Intrinsic: '_mm_sra_epi32'.
// Requires SSE2.
func SraEpi32(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(sraEpi32([16]byte(a), [16]byte(count)))
}

func sraEpi32(a [16]byte, count [16]byte) [16]byte


// SraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAW'. Intrinsic: '_mm_srai_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SraiEpi16(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(sraiEpi16([16]byte(a), imm8))
}

func sraiEpi16(a [16]byte, imm8 byte) [16]byte


// SraiEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRAD'. Intrinsic: '_mm_srai_epi32'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SraiEpi32(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(sraiEpi32([16]byte(a), imm8))
}

func sraiEpi32(a [16]byte, imm8 byte) [16]byte


// SrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLW'. Intrinsic: '_mm_srl_epi16'.
// Requires SSE2.
func SrlEpi16(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(srlEpi16([16]byte(a), [16]byte(count)))
}

func srlEpi16(a [16]byte, count [16]byte) [16]byte


// SrlEpi32: Shift packed 32-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLD'. Intrinsic: '_mm_srl_epi32'.
// Requires SSE2.
func SrlEpi32(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(srlEpi32([16]byte(a), [16]byte(count)))
}

func srlEpi32(a [16]byte, count [16]byte) [16]byte


// SrlEpi64: Shift packed 64-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF count[63:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] >> count[63:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLQ'. Intrinsic: '_mm_srl_epi64'.
// Requires SSE2.
func SrlEpi64(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	return x86.M128i(srlEpi64([16]byte(a), [16]byte(count)))
}

func srlEpi64(a [16]byte, count [16]byte) [16]byte


// SrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLW'. Intrinsic: '_mm_srli_epi16'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SrliEpi16(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(srliEpi16([16]byte(a), imm8))
}

func srliEpi16(a [16]byte, imm8 byte) [16]byte


// SrliEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLD'. Intrinsic: '_mm_srli_epi32'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SrliEpi32(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(srliEpi32([16]byte(a), imm8))
}

func srliEpi32(a [16]byte, imm8 byte) [16]byte


// SrliEpi64: Shift packed 64-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF imm8[7:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] >> imm8[7:0])
//			FI
//		ENDFOR
//
// Instruction: 'PSRLQ'. Intrinsic: '_mm_srli_epi64'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SrliEpi64(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(srliEpi64([16]byte(a), imm8))
}

func srliEpi64(a [16]byte, imm8 byte) [16]byte


// SrliSi128: Shift 'a' right by 'imm8' bytes while shifting in zeros, and
// store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] >> (tmp*8)
//
// Instruction: 'PSRLDQ'. Intrinsic: '_mm_srli_si128'.
// Requires SSE2.
//
// FIXME: Requires compiler support (has immediate)
func SrliSi128(a x86.M128i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(srliSi128([16]byte(a), imm8))
}

func srliSi128(a [16]byte, imm8 byte) [16]byte


// StorePd: Store 128-bits (composed of 2 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVAPD'. Intrinsic: '_mm_store_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorePd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorePd1: Store the lower double-precision (64-bit) floating-point element
// from 'a' into 2 contiguous elements in memory. 'mem_addr' must be aligned on
// a 16-byte boundary or a general-protection exception may be generated. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//		MEM[mem_addr+127:mem_addr+64] := a[63:0]
//
// Instruction: '...'. Intrinsic: '_mm_store_pd1'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorePd1(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StoreSd: Store the lower double-precision (64-bit) floating-point element
// from 'a' into memory. 'mem_addr' does not need to be aligned on any
// particular boundary. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVSD'. Intrinsic: '_mm_store_sd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StoreSd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StoreSi128: Store 128-bits of integer data from 'a' into memory. 
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVDQA'. Intrinsic: '_mm_store_si128'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StoreSi128(mem_addr *x86.M128i, a x86.M128i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// Store1Pd: Store the lower double-precision (64-bit) floating-point element
// from 'a' into 2 contiguous elements in memory. 'mem_addr' must be aligned on
// a 16-byte boundary or a general-protection exception may be generated. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//		MEM[mem_addr+127:mem_addr+64] := a[63:0]
//
// Instruction: '...'. Intrinsic: '_mm_store1_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Store1Pd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorehPd: Store the upper double-precision (64-bit) floating-point element
// from 'a' into memory. 
//
//		MEM[mem_addr+63:mem_addr] := a[127:64]
//
// Instruction: 'MOVHPD'. Intrinsic: '_mm_storeh_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorehPd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorelEpi64: Store 64-bit integer from the first element of 'a' into memory. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_storel_epi64'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorelEpi64(mem_addr *x86.M128i, a x86.M128i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorelPd: Store the lower double-precision (64-bit) floating-point element
// from 'a' into memory. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVLPD'. Intrinsic: '_mm_storel_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorelPd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorerPd: Store 2 double-precision (64-bit) floating-point elements from 'a'
// into memory in reverse order.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+63:mem_addr] := a[127:64]
//		MEM[mem_addr+127:mem_addr+64] := a[63:0]
//
// Instruction: '...'. Intrinsic: '_mm_storer_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorerPd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StoreuPd: Store 128-bits (composed of 2 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVUPD'. Intrinsic: '_mm_storeu_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StoreuPd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StoreuSi128: Store 128-bits of integer data from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVDQU'. Intrinsic: '_mm_storeu_si128'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StoreuSi128(mem_addr *x86.M128i, a x86.M128i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StreamPd: Store 128-bits (composed of 2 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory using a non-temporal memory
// hint.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVNTPD'. Intrinsic: '_mm_stream_pd'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StreamPd(mem_addr *float64, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StreamSi128: Store 128-bits of integer data from 'a' into memory using a
// non-temporal memory hint. 
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVNTDQ'. Intrinsic: '_mm_stream_si128'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StreamSi128(mem_addr *x86.M128i, a x86.M128i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StreamSi32: Store 32-bit integer 'a' into memory using a non-temporal hint
// to minimize cache pollution. If the cache line containing address 'mem_addr'
// is already in the cache, the cache will be updated. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVNTI'. Intrinsic: '_mm_stream_si32'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StreamSi32(mem_addr *int, a int)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StreamSi64: Store 64-bit integer 'a' into memory using a non-temporal hint
// to minimize cache pollution. If the cache line containing address 'mem_addr'
// is already in the cache, the cache will be updated. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVNTI'. Intrinsic: '_mm_stream_si64'.
// Requires SSE2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StreamSi64(mem_addr *int64, a int64)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// SubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := a[i+15:i] - b[i+15:i]
//		ENDFOR
//
// Instruction: 'PSUBW'. Intrinsic: '_mm_sub_epi16'.
// Requires SSE2.
func SubEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subEpi16([16]byte(a), [16]byte(b)))
}

func subEpi16(a [16]byte, b [16]byte) [16]byte


// SubEpi32: Subtract packed 32-bit integers in 'b' from packed 32-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//
// Instruction: 'PSUBD'. Intrinsic: '_mm_sub_epi32'.
// Requires SSE2.
func SubEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subEpi32([16]byte(a), [16]byte(b)))
}

func subEpi32(a [16]byte, b [16]byte) [16]byte


// SubEpi64: Subtract packed 64-bit integers in 'b' from packed 64-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//
// Instruction: 'PSUBQ'. Intrinsic: '_mm_sub_epi64'.
// Requires SSE2.
func SubEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subEpi64([16]byte(a), [16]byte(b)))
}

func subEpi64(a [16]byte, b [16]byte) [16]byte


// SubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit integers in
// 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := a[i+7:i] - b[i+7:i]
//		ENDFOR
//
// Instruction: 'PSUBB'. Intrinsic: '_mm_sub_epi8'.
// Requires SSE2.
func SubEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subEpi8([16]byte(a), [16]byte(b)))
}

func subEpi8(a [16]byte, b [16]byte) [16]byte


// SubPd: Subtract packed double-precision (64-bit) floating-point elements in
// 'b' from packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//
// Instruction: 'SUBPD'. Intrinsic: '_mm_sub_pd'.
// Requires SSE2.
func SubPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(subPd([2]float64(a), [2]float64(b)))
}

func subPd(a [2]float64, b [2]float64) [2]float64


// SubSd: Subtract the lower double-precision (64-bit) floating-point element
// in 'b' from the lower double-precision (64-bit) floating-point element in
// 'a', store the result in the lower element of 'dst', and copy the upper
// element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := a[63:0] - b[63:0]
//		dst[127:64] := a[127:64]
//
// Instruction: 'SUBSD'. Intrinsic: '_mm_sub_sd'.
// Requires SSE2.
func SubSd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(subSd([2]float64(a), [2]float64(b)))
}

func subSd(a [2]float64, b [2]float64) [2]float64


// SubSi64: Subtract 64-bit integer 'b' from 64-bit integer 'a', and store the
// result in 'dst'. 
//
//		dst[63:0] := a[63:0] - b[63:0]
//
// Instruction: 'PSUBQ'. Intrinsic: '_mm_sub_si64'.
// Requires SSE2.
func SubSi64(a x86.M64, b x86.M64) (dst x86.M64) {
	return x86.M64(subSi64(a, b))
}

func subSi64(a x86.M64, b x86.M64) x86.M64


// SubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//		ENDFOR
//
// Instruction: 'PSUBSW'. Intrinsic: '_mm_subs_epi16'.
// Requires SSE2.
func SubsEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subsEpi16([16]byte(a), [16]byte(b)))
}

func subsEpi16(a [16]byte, b [16]byte) [16]byte


// SubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit integers
// in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//
// Instruction: 'PSUBSB'. Intrinsic: '_mm_subs_epi8'.
// Requires SSE2.
func SubsEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subsEpi8([16]byte(a), [16]byte(b)))
}

func subsEpi8(a [16]byte, b [16]byte) [16]byte


// SubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])	
//		ENDFOR
//
// Instruction: 'PSUBUSW'. Intrinsic: '_mm_subs_epu16'.
// Requires SSE2.
func SubsEpu16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subsEpu16([16]byte(a), [16]byte(b)))
}

func subsEpu16(a [16]byte, b [16]byte) [16]byte


// SubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//
// Instruction: 'PSUBUSB'. Intrinsic: '_mm_subs_epu8'.
// Requires SSE2.
func SubsEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(subsEpu8([16]byte(a), [16]byte(b)))
}

func subsEpu8(a [16]byte, b [16]byte) [16]byte


// UcomieqSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for equality, and return the boolean result (0 or 1).
// This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[63:0] == b[63:0] ) ? 1 : 0
//
// Instruction: 'UCOMISD'. Intrinsic: '_mm_ucomieq_sd'.
// Requires SSE2.
func UcomieqSd(a x86.M128d, b x86.M128d) int {
	return int(ucomieqSd([2]float64(a), [2]float64(b)))
}

func ucomieqSd(a [2]float64, b [2]float64) int


// UcomigeSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for greater-than-or-equal, and return the boolean
// result (0 or 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[63:0] >= b[63:0] ) ? 1 : 0
//
// Instruction: 'UCOMISD'. Intrinsic: '_mm_ucomige_sd'.
// Requires SSE2.
func UcomigeSd(a x86.M128d, b x86.M128d) int {
	return int(ucomigeSd([2]float64(a), [2]float64(b)))
}

func ucomigeSd(a [2]float64, b [2]float64) int


// UcomigtSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for greater-than, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[63:0] > b[63:0] ) ? 1 : 0
//
// Instruction: 'UCOMISD'. Intrinsic: '_mm_ucomigt_sd'.
// Requires SSE2.
func UcomigtSd(a x86.M128d, b x86.M128d) int {
	return int(ucomigtSd([2]float64(a), [2]float64(b)))
}

func ucomigtSd(a [2]float64, b [2]float64) int


// UcomileSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for less-than-or-equal, and return the boolean result
// (0 or 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[63:0] <= b[63:0] ) ? 1 : 0
//
// Instruction: 'UCOMISD'. Intrinsic: '_mm_ucomile_sd'.
// Requires SSE2.
func UcomileSd(a x86.M128d, b x86.M128d) int {
	return int(ucomileSd([2]float64(a), [2]float64(b)))
}

func ucomileSd(a [2]float64, b [2]float64) int


// UcomiltSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for less-than, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[63:0] < b[63:0] ) ? 1 : 0
//
// Instruction: 'UCOMISD'. Intrinsic: '_mm_ucomilt_sd'.
// Requires SSE2.
func UcomiltSd(a x86.M128d, b x86.M128d) int {
	return int(ucomiltSd([2]float64(a), [2]float64(b)))
}

func ucomiltSd(a [2]float64, b [2]float64) int


// UcomineqSd: Compare the lower double-precision (64-bit) floating-point
// element in 'a' and 'b' for not-equal, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[63:0] != b[63:0] ) ? 1 : 0
//
// Instruction: 'UCOMISD'. Intrinsic: '_mm_ucomineq_sd'.
// Requires SSE2.
func UcomineqSd(a x86.M128d, b x86.M128d) int {
	return int(ucomineqSd([2]float64(a), [2]float64(b)))
}

func ucomineqSd(a [2]float64, b [2]float64) int


// UnpackhiEpi16: Unpack and interleave 16-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_WORDS(src1[127:0], src2[127:0]){
//			dst[15:0] := src1[79:64]
//			dst[31:16] := src2[79:64] 
//			dst[47:32] := src1[95:80] 
//			dst[63:48] := src2[95:80] 
//			dst[79:64] := src1[111:96] 
//			dst[95:80] := src2[111:96] 
//			dst[111:96] := src1[127:112] 
//			dst[127:112] := src2[127:112] 
//			RETURN dst[127:0]
//		}
//		
//		dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKHWD'. Intrinsic: '_mm_unpackhi_epi16'.
// Requires SSE2.
func UnpackhiEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackhiEpi16([16]byte(a), [16]byte(b)))
}

func unpackhiEpi16(a [16]byte, b [16]byte) [16]byte


// UnpackhiEpi32: Unpack and interleave 32-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_DWORDS(src1[127:0], src2[127:0]){
//			dst[31:0] := src1[95:64] 
//			dst[63:32] := src2[95:64] 
//			dst[95:64] := src1[127:96] 
//			dst[127:96] := src2[127:96] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_DWORDS(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKHDQ'. Intrinsic: '_mm_unpackhi_epi32'.
// Requires SSE2.
func UnpackhiEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackhiEpi32([16]byte(a), [16]byte(b)))
}

func unpackhiEpi32(a [16]byte, b [16]byte) [16]byte


// UnpackhiEpi64: Unpack and interleave 64-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[127:64] 
//			dst[127:64] := src2[127:64] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_QWORDS(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKHQDQ'. Intrinsic: '_mm_unpackhi_epi64'.
// Requires SSE2.
func UnpackhiEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackhiEpi64([16]byte(a), [16]byte(b)))
}

func unpackhiEpi64(a [16]byte, b [16]byte) [16]byte


// UnpackhiEpi8: Unpack and interleave 8-bit integers from the high half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_BYTES(src1[127:0], src2[127:0]){
//			dst[7:0] := src1[71:64] 
//			dst[15:8] := src2[71:64] 
//			dst[23:16] := src1[79:72] 
//			dst[31:24] := src2[79:72] 
//			dst[39:32] := src1[87:80] 
//			dst[47:40] := src2[87:80] 
//			dst[55:48] := src1[95:88] 
//			dst[63:56] := src2[95:88] 
//			dst[71:64] := src1[103:96] 
//			dst[79:72] := src2[103:96] 
//			dst[87:80] := src1[111:104] 
//			dst[95:88] := src2[111:104] 
//			dst[103:96] := src1[119:112] 
//			dst[111:104] := src2[119:112] 
//			dst[119:112] := src1[127:120] 
//			dst[127:120] := src2[127:120] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKHBW'. Intrinsic: '_mm_unpackhi_epi8'.
// Requires SSE2.
func UnpackhiEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackhiEpi8([16]byte(a), [16]byte(b)))
}

func unpackhiEpi8(a [16]byte, b [16]byte) [16]byte


// UnpackhiPd: Unpack and interleave double-precision (64-bit) floating-point
// elements from the high half of 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[127:64] 
//			dst[127:64] := src2[127:64] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_QWORDS(a[127:0], b[127:0])
//
// Instruction: 'UNPCKHPD'. Intrinsic: '_mm_unpackhi_pd'.
// Requires SSE2.
func UnpackhiPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(unpackhiPd([2]float64(a), [2]float64(b)))
}

func unpackhiPd(a [2]float64, b [2]float64) [2]float64


// UnpackloEpi16: Unpack and interleave 16-bit integers from the low half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_WORDS(src1[127:0], src2[127:0]){
//			dst[15:0] := src1[15:0] 
//			dst[31:16] := src2[15:0] 
//			dst[47:32] := src1[31:16] 
//			dst[63:48] := src2[31:16] 
//			dst[79:64] := src1[47:32] 
//			dst[95:80] := src2[47:32] 
//			dst[111:96] := src1[63:48] 
//			dst[127:112] := src2[63:48] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKLWD'. Intrinsic: '_mm_unpacklo_epi16'.
// Requires SSE2.
func UnpackloEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackloEpi16([16]byte(a), [16]byte(b)))
}

func unpackloEpi16(a [16]byte, b [16]byte) [16]byte


// UnpackloEpi32: Unpack and interleave 32-bit integers from the low half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_DWORDS(src1[127:0], src2[127:0]){
//			dst[31:0] := src1[31:0] 
//			dst[63:32] := src2[31:0] 
//			dst[95:64] := src1[63:32] 
//			dst[127:96] := src2[63:32] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_DWORDS(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKLDQ'. Intrinsic: '_mm_unpacklo_epi32'.
// Requires SSE2.
func UnpackloEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackloEpi32([16]byte(a), [16]byte(b)))
}

func unpackloEpi32(a [16]byte, b [16]byte) [16]byte


// UnpackloEpi64: Unpack and interleave 64-bit integers from the low half of
// 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[63:0] 
//			dst[127:64] := src2[63:0] 
//			RETURN dst[127:0]
//		}
//		
//		dst[127:0] := INTERLEAVE_QWORDS(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKLQDQ'. Intrinsic: '_mm_unpacklo_epi64'.
// Requires SSE2.
func UnpackloEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackloEpi64([16]byte(a), [16]byte(b)))
}

func unpackloEpi64(a [16]byte, b [16]byte) [16]byte


// UnpackloEpi8: Unpack and interleave 8-bit integers from the low half of 'a'
// and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_BYTES(src1[127:0], src2[127:0]){
//			dst[7:0] := src1[7:0] 
//			dst[15:8] := src2[7:0] 
//			dst[23:16] := src1[15:8] 
//			dst[31:24] := src2[15:8] 
//			dst[39:32] := src1[23:16] 
//			dst[47:40] := src2[23:16] 
//			dst[55:48] := src1[31:24] 
//			dst[63:56] := src2[31:24] 
//			dst[71:64] := src1[39:32]
//			dst[79:72] := src2[39:32] 
//			dst[87:80] := src1[47:40] 
//			dst[95:88] := src2[47:40] 
//			dst[103:96] := src1[55:48] 
//			dst[111:104] := src2[55:48] 
//			dst[119:112] := src1[63:56] 
//			dst[127:120] := src2[63:56] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//
// Instruction: 'PUNPCKLBW'. Intrinsic: '_mm_unpacklo_epi8'.
// Requires SSE2.
func UnpackloEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(unpackloEpi8([16]byte(a), [16]byte(b)))
}

func unpackloEpi8(a [16]byte, b [16]byte) [16]byte


// UnpackloPd: Unpack and interleave double-precision (64-bit) floating-point
// elements from the low half of 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[63:0] 
//			dst[127:64] := src2[63:0] 
//			RETURN dst[127:0]
//		}
//		
//		dst[127:0] := INTERLEAVE_QWORDS(a[127:0], b[127:0])
//
// Instruction: 'UNPCKLPD'. Intrinsic: '_mm_unpacklo_pd'.
// Requires SSE2.
func UnpackloPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(unpackloPd([2]float64(a), [2]float64(b)))
}

func unpackloPd(a [2]float64, b [2]float64) [2]float64


// XorPd: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//		ENDFOR
//
// Instruction: 'XORPD'. Intrinsic: '_mm_xor_pd'.
// Requires SSE2.
func XorPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	return x86.M128d(xorPd([2]float64(a), [2]float64(b)))
}

func xorPd(a [2]float64, b [2]float64) [2]float64


// XorSi128: Compute the bitwise XOR of 128 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[127:0] := (a[127:0] XOR b[127:0])
//
// Instruction: 'PXOR'. Intrinsic: '_mm_xor_si128'.
// Requires SSE2.
func XorSi128(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(xorSi128([16]byte(a), [16]byte(b)))
}

func xorSi128(a [16]byte, b [16]byte) [16]byte

