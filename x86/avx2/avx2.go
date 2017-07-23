// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package avx2

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// M256AbsEpi16: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := ABS(a[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm256_abs_epi16'.
// Requires AVX2.
func M256AbsEpi16(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AbsEpi32: Compute the absolute value of packed 32-bit integers in 'a',
// and store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ABS(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSD'. Intrinsic: '_mm256_abs_epi32'.
// Requires AVX2.
func M256AbsEpi32(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AbsEpi8: Compute the absolute value of packed 8-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := ABS(a[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm256_abs_epi8'.
// Requires AVX2.
func M256AbsEpi8(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := a[i+15:i] + b[i+15:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm256_add_epi16'.
// Requires AVX2.
func M256AddEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddEpi32: Add packed 32-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDD'. Intrinsic: '_mm256_add_epi32'.
// Requires AVX2.
func M256AddEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddEpi64: Add packed 64-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDQ'. Intrinsic: '_mm256_add_epi64'.
// Requires AVX2.
func M256AddEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := a[i+7:i] + b[i+7:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm256_add_epi8'.
// Requires AVX2.
func M256AddEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddsEpi16: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm256_adds_epi16'.
// Requires AVX2.
func M256AddsEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm256_adds_epi8'.
// Requires AVX2.
func M256AddsEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm256_adds_epu16'.
// Requires AVX2.
func M256AddsEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm256_adds_epu8'.
// Requires AVX2.
func M256AddsEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*128
//			tmp[255:0] := ((a[i+127:i] << 128) OR b[i+127:i]) >> (count[7:0]*8)
//			dst[i+127:i] := tmp[127:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm256_alignr_epi8'.
// Requires AVX2.
func M256AlignrEpi8(a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M256AndSi256: Compute the bitwise AND of 256 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[255:0] := (a[255:0] AND b[255:0])
//		dst[MAX:256] := 0
//
// Instruction: 'VPAND'. Intrinsic: '_mm256_and_si256'.
// Requires AVX2.
func M256AndSi256(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AndnotSi256: Compute the bitwise AND NOT of 256 bits (representing
// integer data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[255:0] := ((NOT a[255:0]) AND b[255:0])
//		dst[MAX:256] := 0
//
// Instruction: 'VPANDN'. Intrinsic: '_mm256_andnot_si256'.
// Requires AVX2.
func M256AndnotSi256(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm256_avg_epu16'.
// Requires AVX2.
func M256AvgEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256AvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm256_avg_epu8'.
// Requires AVX2.
func M256AvgEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256BlendEpi16: Blend packed 16-bit integers from 'a' and 'b' using control
// mask 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF imm8[j%8]
//				dst[i+15:i] := b[i+15:i]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBLENDW'. Intrinsic: '_mm256_blend_epi16'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256BlendEpi16(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// BlendEpi32: Blend packed 32-bit integers from 'a' and 'b' using control mask
// 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF imm8[j%8]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBLENDD'. Intrinsic: '_mm_blend_epi32'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func BlendEpi32(a x86.M128i, b x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256BlendEpi32: Blend packed 32-bit integers from 'a' and 'b' using control
// mask 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF imm8[j%8]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBLENDD'. Intrinsic: '_mm256_blend_epi32'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256BlendEpi32(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256BlendvEpi8: Blend packed 8-bit integers from 'a' and 'b' using 'mask',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF mask[i+7]
//				dst[i+7:i] := b[i+7:i]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBLENDVB'. Intrinsic: '_mm256_blendv_epi8'.
// Requires AVX2.
func M256BlendvEpi8(a x86.M256i, b x86.M256i, mask x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// BroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm_broadcastb_epi8'.
// Requires AVX2.
func BroadcastbEpi8(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256BroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm256_broadcastb_epi8'.
// Requires AVX2.
func M256BroadcastbEpi8(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// BroadcastdEpi32: Broadcast the low packed 32-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTD'. Intrinsic: '_mm_broadcastd_epi32'.
// Requires AVX2.
func BroadcastdEpi32(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256BroadcastdEpi32: Broadcast the low packed 32-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTD'. Intrinsic: '_mm256_broadcastd_epi32'.
// Requires AVX2.
func M256BroadcastdEpi32(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// BroadcastqEpi64: Broadcast the low packed 64-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTQ'. Intrinsic: '_mm_broadcastq_epi64'.
// Requires AVX2.
func BroadcastqEpi64(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256BroadcastqEpi64: Broadcast the low packed 64-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTQ'. Intrinsic: '_mm256_broadcastq_epi64'.
// Requires AVX2.
func M256BroadcastqEpi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// BroadcastsdPd: Broadcast the low double-precision (64-bit) floating-point
// element from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'MOVDDUP'. Intrinsic: '_mm_broadcastsd_pd'.
// Requires AVX2.
func BroadcastsdPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// M256BroadcastsdPd: Broadcast the low double-precision (64-bit)
// floating-point element from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTSD'. Intrinsic: '_mm256_broadcastsd_pd'.
// Requires AVX2.
func M256BroadcastsdPd(a x86.M128d) (dst x86.M256d) {
	panic("not implemented")
}


// M256Broadcastsi128Si256: Broadcast 128 bits of integer data from 'a' to all
// 128-bit lanes in 'dst'. 
//
//		dst[127:0] := a[127:0]
//		dst[255:128] := a[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI128'. Intrinsic: '_mm256_broadcastsi128_si256'.
// Requires AVX2.
func M256Broadcastsi128Si256(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// BroadcastssPs: Broadcast the low single-precision (32-bit) floating-point
// element from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VBROADCASTSS'. Intrinsic: '_mm_broadcastss_ps'.
// Requires AVX2.
func BroadcastssPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// M256BroadcastssPs: Broadcast the low single-precision (32-bit)
// floating-point element from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTSS'. Intrinsic: '_mm256_broadcastss_ps'.
// Requires AVX2.
func M256BroadcastssPs(a x86.M128) (dst x86.M256) {
	panic("not implemented")
}


// BroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm_broadcastw_epi16'.
// Requires AVX2.
func BroadcastwEpi16(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256BroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm256_broadcastw_epi16'.
// Requires AVX2.
func M256BroadcastwEpi16(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256BslliEpi128: Shift 128-bit lanes in 'a' left by 'imm8' bytes while
// shifting in zeros, and store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] << (tmp*8)
//		dst[255:128] := a[255:128] << (tmp*8)
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLDQ'. Intrinsic: '_mm256_bslli_epi128'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256BslliEpi128(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256BsrliEpi128: Shift 128-bit lanes in 'a' right by 'imm8' bytes while
// shifting in zeros, and store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] >> (tmp*8)
//		dst[255:128] := a[255:128] >> (tmp*8)
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLDQ'. Intrinsic: '_mm256_bsrli_epi128'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256BsrliEpi128(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpeqEpi16: Compare packed 16-bit integers in 'a' and 'b' for equality,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] == b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPEQW'. Intrinsic: '_mm256_cmpeq_epi16'.
// Requires AVX2.
func M256CmpeqEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpeqEpi32: Compare packed 32-bit integers in 'a' and 'b' for equality,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] == b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPEQD'. Intrinsic: '_mm256_cmpeq_epi32'.
// Requires AVX2.
func M256CmpeqEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpeqEpi64: Compare packed 64-bit integers in 'a' and 'b' for equality,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ( a[i+63:i] == b[i+63:i] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPEQQ'. Intrinsic: '_mm256_cmpeq_epi64'.
// Requires AVX2.
func M256CmpeqEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpeqEpi8: Compare packed 8-bit integers in 'a' and 'b' for equality,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] == b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPEQB'. Intrinsic: '_mm256_cmpeq_epi8'.
// Requires AVX2.
func M256CmpeqEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpgtEpi16: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := ( a[i+15:i] > b[i+15:i] ) ? 0xFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPGTW'. Intrinsic: '_mm256_cmpgt_epi16'.
// Requires AVX2.
func M256CmpgtEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpgtEpi32: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] > b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPGTD'. Intrinsic: '_mm256_cmpgt_epi32'.
// Requires AVX2.
func M256CmpgtEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpgtEpi64: Compare packed 64-bit integers in 'a' and 'b' for
// greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ( a[i+63:i] > b[i+63:i] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPGTQ'. Intrinsic: '_mm256_cmpgt_epi64'.
// Requires AVX2.
func M256CmpgtEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256CmpgtEpi8: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := ( a[i+7:i] > b[i+7:i] ) ? 0xFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCMPGTB'. Intrinsic: '_mm256_cmpgt_epi8'.
// Requires AVX2.
func M256CmpgtEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepi16Epi32: Sign extend packed 16-bit integers in 'a' to packed
// 32-bit integers, and store the results in 'dst'. 
//
//		FOR j:= 0 to 7
//			i := 32*j
//			k := 16*j
//			dst[i+31:i] := SignExtend(a[k+15:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXWD'. Intrinsic: '_mm256_cvtepi16_epi32'.
// Requires AVX2.
func M256Cvtepi16Epi32(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepi16Epi64: Sign extend packed 16-bit integers in 'a' to packed
// 64-bit integers, and store the results in 'dst'. 
//
//		FOR j:= 0 to 3
//			i := 64*j
//			k := 16*j
//			dst[i+63:i] := SignExtend(a[k+15:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXWQ'. Intrinsic: '_mm256_cvtepi16_epi64'.
// Requires AVX2.
func M256Cvtepi16Epi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepi32Epi64: Sign extend packed 32-bit integers in 'a' to packed
// 64-bit integers, and store the results in 'dst'. 
//
//		FOR j:= 0 to 3
//			i := 64*j
//			k := 32*j
//			dst[i+63:i] := SignExtend(a[k+31:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXDQ'. Intrinsic: '_mm256_cvtepi32_epi64'.
// Requires AVX2.
func M256Cvtepi32Epi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			l := j*16
//			dst[l+15:l] := SignExtend(a[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm256_cvtepi8_epi16'.
// Requires AVX2.
func M256Cvtepi8Epi16(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepi8Epi32: Sign extend packed 8-bit integers in 'a' to packed 32-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			k := 8*j
//			dst[i+31:i] := SignExtend(a[k+7:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXBD'. Intrinsic: '_mm256_cvtepi8_epi32'.
// Requires AVX2.
func M256Cvtepi8Epi32(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepi8Epi64: Sign extend packed 8-bit integers in the low 8 bytes of
// 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			k := 8*j
//			dst[i+63:i] := SignExtend(a[k+7:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXBQ'. Intrinsic: '_mm256_cvtepi8_epi64'.
// Requires AVX2.
func M256Cvtepi8Epi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepu16Epi32: Zero extend packed unsigned 16-bit integers in 'a' to
// packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			k := 16*j
//			dst[i+31:i] := ZeroExtend(a[k+15:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXWD'. Intrinsic: '_mm256_cvtepu16_epi32'.
// Requires AVX2.
func M256Cvtepu16Epi32(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepu16Epi64: Zero extend packed unsigned 16-bit integers in 'a' to
// packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j:= 0 to 3
//			i := 64*j
//			k := 16*j
//			dst[i+63:i] := ZeroExtend(a[k+15:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXWQ'. Intrinsic: '_mm256_cvtepu16_epi64'.
// Requires AVX2.
func M256Cvtepu16Epi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepu32Epi64: Zero extend packed unsigned 32-bit integers in 'a' to
// packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j:= 0 to 3
//			i := 64*j
//			k := 32*j
//			dst[i+63:i] := ZeroExtend(a[k+31:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXDQ'. Intrinsic: '_mm256_cvtepu32_epi64'.
// Requires AVX2.
func M256Cvtepu32Epi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			l := j*16
//			dst[l+15:l] := ZeroExtend(a[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm256_cvtepu8_epi16'.
// Requires AVX2.
func M256Cvtepu8Epi16(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepu8Epi32: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			k := 8*j
//			dst[i+31:i] := ZeroExtend(a[k+7:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXBD'. Intrinsic: '_mm256_cvtepu8_epi32'.
// Requires AVX2.
func M256Cvtepu8Epi32(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Cvtepu8Epi64: Zero extend packed unsigned 8-bit integers in the low 8
// byte sof 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			k := 8*j
//			dst[i+63:i] := ZeroExtend(a[k+7:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXBQ'. Intrinsic: '_mm256_cvtepu8_epi64'.
// Requires AVX2.
func M256Cvtepu8Epi64(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Extracti128Si256: Extract 128 bits (composed of integer data) from 'a',
// selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI128'. Intrinsic: '_mm256_extracti128_si256'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256Extracti128Si256(a x86.M256i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256HaddEpi16: Horizontally add adjacent pairs of 16-bit integers in 'a' and
// 'b', and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[31:16] + a[15:0]
//		dst[31:16] := a[63:48] + a[47:32]
//		dst[47:32] := a[95:80] + a[79:64]
//		dst[63:48] := a[127:112] + a[111:96]
//		dst[79:64] := b[31:16] + b[15:0]
//		dst[95:80] := b[63:48] + b[47:32]
//		dst[111:96] := b[95:80] + b[79:64]
//		dst[127:112] := b[127:112] + b[111:96]
//		dst[143:128] := a[159:144] + a[143:128]
//		dst[159:144] := a[191:176] + a[175:160]
//		dst[175:160] := a[223:208] + a[207:192]
//		dst[191:176] := a[255:240] + a[239:224]
//		dst[207:192] := b[127:112] + b[143:128]
//		dst[223:208] := b[159:144] + b[175:160]
//		dst[239:224] := b[191:176] + b[207:192]
//		dst[255:240] := b[223:208] + b[239:224]
//		dst[MAX:256] := 0
//
// Instruction: 'VPHADDW'. Intrinsic: '_mm256_hadd_epi16'.
// Requires AVX2.
func M256HaddEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256HaddEpi32: Horizontally add adjacent pairs of 32-bit integers in 'a' and
// 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := a[127:96] + a[95:64]
//		dst[95:64] := b[63:32] + b[31:0]
//		dst[127:96] := b[127:96] + b[95:64]
//		dst[159:128] := a[191:160] + a[159:128]
//		dst[191:160] := a[255:224] + a[223:192]
//		dst[223:192] := b[191:160] + b[159:128]
//		dst[255:224] := b[255:224] + b[223:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VPHADDD'. Intrinsic: '_mm256_hadd_epi32'.
// Requires AVX2.
func M256HaddEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256HaddsEpi16: Horizontally add adjacent pairs of 16-bit integers in 'a'
// and 'b' using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[31:16] + a[15:0])
//		dst[31:16] = Saturate_To_Int16(a[63:48] + a[47:32])
//		dst[47:32] = Saturate_To_Int16(a[95:80] + a[79:64])
//		dst[63:48] = Saturate_To_Int16(a[127:112] + a[111:96])
//		dst[79:64] = Saturate_To_Int16(b[31:16] + b[15:0])
//		dst[95:80] = Saturate_To_Int16(b[63:48] + b[47:32])
//		dst[111:96] = Saturate_To_Int16(b[95:80] + b[79:64])
//		dst[127:112] = Saturate_To_Int16(b[127:112] + b[111:96])
//		dst[143:128] = Saturate_To_Int16(a[159:144] + a[143:128])
//		dst[159:144] = Saturate_To_Int16(a[191:176] + a[175:160])
//		dst[175:160] = Saturate_To_Int16( a[223:208] + a[207:192])
//		dst[191:176] = Saturate_To_Int16(a[255:240] + a[239:224])
//		dst[207:192] = Saturate_To_Int16(b[127:112] + b[143:128])
//		dst[223:208] = Saturate_To_Int16(b[159:144] + b[175:160])
//		dst[239:224] = Saturate_To_Int16(b[191-160] + b[159-128])
//		dst[255:240] = Saturate_To_Int16(b[255:240] + b[239:224])
//		dst[MAX:256] := 0
//
// Instruction: 'VPHADDSW'. Intrinsic: '_mm256_hadds_epi16'.
// Requires AVX2.
func M256HaddsEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256HsubEpi16: Horizontally subtract adjacent pairs of 16-bit integers in
// 'a' and 'b', and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[15:0] - a[31:16]
//		dst[31:16] := a[47:32] - a[63:48]
//		dst[47:32] := a[79:64] - a[95:80]
//		dst[63:48] := a[111:96] - a[127:112]
//		dst[79:64] := b[15:0] - b[31:16]
//		dst[95:80] := b[47:32] - b[63:48]
//		dst[111:96] := b[79:64] - b[95:80]
//		dst[127:112] := b[111:96] - b[127:112]
//		dst[143:128] := a[143:128] - a[159:144]
//		dst[159:144] := a[175:160] - a[191:176]
//		dst[175:160] := a[207:192] - a[223:208]
//		dst[191:176] := a[239:224] - a[255:240]
//		dst[207:192] := b[143:128] - b[159:144]
//		dst[223:208] := b[175:160] - b[191:176]
//		dst[239:224] := b[207:192] - b[223:208]
//		dst[255:240] := b[239:224] - b[255:240]
//		dst[MAX:256] := 0
//
// Instruction: 'VPHSUBW'. Intrinsic: '_mm256_hsub_epi16'.
// Requires AVX2.
func M256HsubEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256HsubEpi32: Horizontally subtract adjacent pairs of 32-bit integers in
// 'a' and 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := a[95:64] - a[127:96]
//		dst[95:64] := b[31:0] - b[63:32]
//		dst[127:96] := b[95:64] - b[127:96]
//		dst[159:128] := a[159:128] - a[191:160]
//		dst[191:160] := a[223:192] - a[255:224]
//		dst[223:192] := b[159:128] - b[191:160]
//		dst[255:224] := b[223:192] - b[255:224]
//		dst[MAX:256] := 0
//
// Instruction: 'VPHSUBD'. Intrinsic: '_mm256_hsub_epi32'.
// Requires AVX2.
func M256HsubEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256HsubsEpi16: Horizontally subtract adjacent pairs of 16-bit integers in
// 'a' and 'b' using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[15:0] - a[31:16])
//		dst[31:16] = Saturate_To_Int16(a[47:32] - a[63:48])
//		dst[47:32] = Saturate_To_Int16(a[79:64] - a[95:80])
//		dst[63:48] = Saturate_To_Int16(a[111:96] - a[127:112])
//		dst[79:64] = Saturate_To_Int16(b[15:0] - b[31:16])
//		dst[95:80] = Saturate_To_Int16(b[47:32] - b[63:48])
//		dst[111:96] = Saturate_To_Int16(b[79:64] - b[95:80])
//		dst[127:112] = Saturate_To_Int16(b[111:96] - b[127:112])
//		dst[143:128]= Saturate_To_Int16(a[143:128] - a[159:144])
//		dst[159:144] = Saturate_To_Int16(a[175:160] - a[191:176])
//		dst[175:160] = Saturate_To_Int16(a[207:192] - a[223:208])
//		dst[191:176] = Saturate_To_Int16(a[239:224] - a[255:240])
//		dst[207:192] = Saturate_To_Int16(b[143:128] - b[159:144])
//		dst[223:208] = Saturate_To_Int16(b[175:160] - b[191:176])
//		dst[239:224] = Saturate_To_Int16(b[207:192] - b[223:208])
//		dst[255:240] = Saturate_To_Int16(b[239:224] - b[255:240])
//		dst[MAX:256] := 0
//
// Instruction: 'VPHSUBSW'. Intrinsic: '_mm256_hsubs_epi16'.
// Requires AVX2.
func M256HsubsEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// I32gatherEpi32: Gather 32-bit integers from memory using 32-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm_i32gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func I32gatherEpi32(base_addr *int, vindex x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskI32gatherEpi32: Gather 32-bit integers from memory using 32-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst' using 'mask' (elements are
// copied from 'src' when the highest bit is not set in the corresponding
// element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//				mask[i+31] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		mask[MAX:128] := 0
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm_mask_i32gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskI32gatherEpi32(src x86.M128i, base_addr *int, vindex x86.M128i, mask x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// M256I32gatherEpi32: Gather 32-bit integers from memory using 32-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm256_i32gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256I32gatherEpi32(base_addr *int, vindex x86.M256i, scale int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskI32gatherEpi32: Gather 32-bit integers from memory using 32-bit
// indices. 32-bit elements are loaded from addresses starting at 'base_addr'
// and offset by each 32-bit element in 'vindex' (each index is scaled by the
// factor in 'scale'). Gathered elements are merged into 'dst' using 'mask'
// (elements are copied from 'src' when the highest bit is not set in the
// corresponding element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//				mask[i+31] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		mask[MAX:256] := 0
//		dst[MAX:256] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm256_mask_i32gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskI32gatherEpi32(src x86.M256i, base_addr *int, vindex x86.M256i, mask x86.M256i, scale int) (dst x86.M256i) {
	panic("not implemented")
}


// I32gatherEpi64: Gather 64-bit integers from memory using 32-bit indices.
// 64-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 1
//			i := j*64
//			m := j*32
//			dst[i+63:i] := MEM[base_addr + SignExtend(vindex[m+31:m])*scale]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm_i32gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func I32gatherEpi64(base_addr *int, vindex x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskI32gatherEpi64: Gather 64-bit integers from memory using 32-bit indices.
// 64-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst' using 'mask' (elements are
// copied from 'src' when the highest bit is not set in the corresponding
// element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 1
//			i := j*64
//			m := j*32
//			IF mask[i+63]
//				dst[i+63:i] := MEM[base_addr + SignExtend(vindex[m+31:m])*scale]
//				mask[i+63] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		mask[MAX:128] := 0
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm_mask_i32gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskI32gatherEpi64(src x86.M128i, base_addr *int, vindex x86.M128i, mask x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// M256I32gatherEpi64: Gather 64-bit integers from memory using 32-bit indices.
// 64-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 3
//			i := j*64
//			m := j*32
//			dst[i+63:i] := MEM[base_addr + SignExtend(vindex[m+31:m])*scale]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm256_i32gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256I32gatherEpi64(base_addr *int, vindex x86.M128i, scale int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskI32gatherEpi64: Gather 64-bit integers from memory using 32-bit
// indices. 64-bit elements are loaded from addresses starting at 'base_addr'
// and offset by each 32-bit element in 'vindex' (each index is scaled by the
// factor in 'scale'). Gathered elements are merged into 'dst' using 'mask'
// (elements are copied from 'src' when the highest bit is not set in the
// corresponding element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 3
//			i := j*64
//			m := j*32
//			IF mask[i+63]
//				dst[i+63:i] := MEM[base_addr + SignExtend(vindex[m+31:m])*scale]
//				mask[i+63] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		mask[MAX:256] := 0
//		dst[MAX:256] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm256_mask_i32gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskI32gatherEpi64(src x86.M256i, base_addr *int, vindex x86.M128i, mask x86.M256i, scale int) (dst x86.M256i) {
	panic("not implemented")
}


// Skipped: _mm_i32gather_pd. Contains pointer parameter.


// Skipped: _mm_mask_i32gather_pd. Contains pointer parameter.


// Skipped: _mm256_i32gather_pd. Contains pointer parameter.


// Skipped: _mm256_mask_i32gather_pd. Contains pointer parameter.


// Skipped: _mm_i32gather_ps. Contains pointer parameter.


// Skipped: _mm_mask_i32gather_ps. Contains pointer parameter.


// Skipped: _mm256_i32gather_ps. Contains pointer parameter.


// Skipped: _mm256_mask_i32gather_ps. Contains pointer parameter.


// I64gatherEpi32: Gather 32-bit integers from memory using 64-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 64-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 1
//			i := j*32
//			m := j*64
//			dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPGATHERQD'. Intrinsic: '_mm_i64gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func I64gatherEpi32(base_addr *int, vindex x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskI64gatherEpi32: Gather 32-bit integers from memory using 64-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 64-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst' using 'mask' (elements are
// copied from 'src' when the highest bit is not set in the corresponding
// element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 1
//			i := j*32
//			m := j*64
//			IF mask[i+31]
//				dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//				mask[i+31] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		mask[MAX:64] := 0
//		dst[MAX:64] := 0
//
// Instruction: 'VPGATHERQD'. Intrinsic: '_mm_mask_i64gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskI64gatherEpi32(src x86.M128i, base_addr *int, vindex x86.M128i, mask x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// M256I64gatherEpi32: Gather 32-bit integers from memory using 64-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 64-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 3
//			i := j*32
//			m := j*64
//			dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERQD'. Intrinsic: '_mm256_i64gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256I64gatherEpi32(base_addr *int, vindex x86.M256i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskI64gatherEpi32: Gather 32-bit integers from memory using 64-bit
// indices. 32-bit elements are loaded from addresses starting at 'base_addr'
// and offset by each 64-bit element in 'vindex' (each index is scaled by the
// factor in 'scale'). Gathered elements are merged into 'dst' using 'mask'
// (elements are copied from 'src' when the highest bit is not set in the
// corresponding element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 3
//			i := j*32
//			m := j*64
//			IF mask[i+31]
//				dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//				mask[i+31] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		mask[MAX:128] := 0
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERQD'. Intrinsic: '_mm256_mask_i64gather_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskI64gatherEpi32(src x86.M128i, base_addr *int, vindex x86.M256i, mask x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// I64gatherEpi64: Gather 64-bit integers from memory using 64-bit indices.
// 64-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 64-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERQQ'. Intrinsic: '_mm_i64gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func I64gatherEpi64(base_addr *int, vindex x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskI64gatherEpi64: Gather 64-bit integers from memory using 64-bit indices.
// 64-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 64-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst' using 'mask' (elements are
// copied from 'src' when the highest bit is not set in the corresponding
// element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//				mask[i+63] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		mask[MAX:128] := 0
//		dst[MAX:128] := 0
//
// Instruction: 'VPGATHERQQ'. Intrinsic: '_mm_mask_i64gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskI64gatherEpi64(src x86.M128i, base_addr *int, vindex x86.M128i, mask x86.M128i, scale int) (dst x86.M128i) {
	panic("not implemented")
}


// M256I64gatherEpi64: Gather 64-bit integers from memory using 64-bit indices.
// 64-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 64-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPGATHERQQ'. Intrinsic: '_mm256_i64gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256I64gatherEpi64(base_addr *int, vindex x86.M256i, scale int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskI64gatherEpi64: Gather 64-bit integers from memory using 64-bit
// indices. 64-bit elements are loaded from addresses starting at 'base_addr'
// and offset by each 64-bit element in 'vindex' (each index is scaled by the
// factor in 'scale'). Gathered elements are merged into 'dst' using 'mask'
// (elements are copied from 'src' when the highest bit is not set in the
// corresponding element). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := MEM[base_addr + SignExtend(vindex[i+63:i])*scale]
//				mask[i+63] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		mask[MAX:256] := 0
//		dst[MAX:256] := 0
//
// Instruction: 'VPGATHERQQ'. Intrinsic: '_mm256_mask_i64gather_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskI64gatherEpi64(src x86.M256i, base_addr *int, vindex x86.M256i, mask x86.M256i, scale int) (dst x86.M256i) {
	panic("not implemented")
}


// Skipped: _mm_i64gather_pd. Contains pointer parameter.


// Skipped: _mm_mask_i64gather_pd. Contains pointer parameter.


// Skipped: _mm256_i64gather_pd. Contains pointer parameter.


// Skipped: _mm256_mask_i64gather_pd. Contains pointer parameter.


// Skipped: _mm_i64gather_ps. Contains pointer parameter.


// Skipped: _mm_mask_i64gather_ps. Contains pointer parameter.


// Skipped: _mm256_i64gather_ps. Contains pointer parameter.


// Skipped: _mm256_mask_i64gather_ps. Contains pointer parameter.


// M256Inserti128Si256: Copy 'a' to 'dst', then insert 128 bits (composed of
// integer data) from 'b' into 'dst' at the location specified by 'imm8'. 
//
//		dst[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTI128'. Intrinsic: '_mm256_inserti128_si256'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256Inserti128Si256(a x86.M256i, b x86.M128i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaddEpi16: Multiply packed signed 16-bit integers in 'a' and 'b',
// producing intermediate signed 32-bit integers. Horizontally add adjacent
// pairs of intermediate 32-bit integers, and pack the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			st[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm256_madd_epi16'.
// Requires AVX2.
func M256MaddEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaddubsEpi16: Vertically multiply each unsigned 8-bit integer from 'a'
// with the corresponding signed 8-bit integer from 'b', producing intermediate
// signed 16-bit integers. Horizontally add adjacent pairs of intermediate
// signed 16-bit integers, and pack the saturated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm256_maddubs_epi16'.
// Requires AVX2.
func M256MaddubsEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// MaskloadEpi32: Load packed 32-bit integers from memory into 'dst' using
// 'mask' (elements are zeroed out when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := MEM[mem_addr+i+31:mem_addr+i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMASKMOVD'. Intrinsic: '_mm_maskload_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskloadEpi32(mem_addr *int, mask x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskloadEpi32: Load packed 32-bit integers from memory into 'dst' using
// 'mask' (elements are zeroed out when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := MEM[mem_addr+i+31:mem_addr+i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMASKMOVD'. Intrinsic: '_mm256_maskload_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskloadEpi32(mem_addr *int, mask x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// MaskloadEpi64: Load packed 64-bit integers from memory into 'dst' using
// 'mask' (elements are zeroed out when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := MEM[mem_addr+i+63:mem_addr+i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMASKMOVQ'. Intrinsic: '_mm_maskload_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskloadEpi64(mem_addr *int, mask x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskloadEpi64: Load packed 64-bit integers from memory into 'dst' using
// 'mask' (elements are zeroed out when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := MEM[mem_addr+i+63:mem_addr+i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMASKMOVQ'. Intrinsic: '_mm256_maskload_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskloadEpi64(mem_addr *int, mask x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// MaskstoreEpi32: Store packed 32-bit integers from 'a' into memory using
// 'mask' (elements are not stored when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF mask[i+31]
//				MEM[mem_addr+i+31:mem_addr+i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'VPMASKMOVD'. Intrinsic: '_mm_maskstore_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskstoreEpi32(mem_addr *int, mask x86.M128i, a x86.M128i)  {
	panic("not implemented")
}


// M256MaskstoreEpi32: Store packed 32-bit integers from 'a' into memory using
// 'mask' (elements are not stored when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF mask[i+31]
//				MEM[mem_addr+i+31:mem_addr+i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'VPMASKMOVD'. Intrinsic: '_mm256_maskstore_epi32'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskstoreEpi32(mem_addr *int, mask x86.M256i, a x86.M256i)  {
	panic("not implemented")
}


// MaskstoreEpi64: Store packed 64-bit integers from 'a' into memory using
// 'mask' (elements are not stored when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF mask[i+63]
//				MEM[mem_addr+i+63:mem_addr+i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VPMASKMOVQ'. Intrinsic: '_mm_maskstore_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskstoreEpi64(mem_addr *int64, mask x86.M128i, a x86.M128i)  {
	panic("not implemented")
}


// M256MaskstoreEpi64: Store packed 64-bit integers from 'a' into memory using
// 'mask' (elements are not stored when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF mask[i+63]
//				MEM[mem_addr+i+63:mem_addr+i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VPMASKMOVQ'. Intrinsic: '_mm256_maskstore_epi64'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskstoreEpi64(mem_addr *int64, mask x86.M256i, a x86.M256i)  {
	panic("not implemented")
}


// M256MaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm256_max_epi16'.
// Requires AVX2.
func M256MaxEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaxEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSD'. Intrinsic: '_mm256_max_epi32'.
// Requires AVX2.
func M256MaxEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm256_max_epi8'.
// Requires AVX2.
func M256MaxEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm256_max_epu16'.
// Requires AVX2.
func M256MaxEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaxEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUD'. Intrinsic: '_mm256_max_epu32'.
// Requires AVX2.
func M256MaxEpu32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm256_max_epu8'.
// Requires AVX2.
func M256MaxEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm256_min_epi16'.
// Requires AVX2.
func M256MinEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MinEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSD'. Intrinsic: '_mm256_min_epi32'.
// Requires AVX2.
func M256MinEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm256_min_epi8'.
// Requires AVX2.
func M256MinEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm256_min_epu16'.
// Requires AVX2.
func M256MinEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MinEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUD'. Intrinsic: '_mm256_min_epu32'.
// Requires AVX2.
func M256MinEpu32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm256_min_epu8'.
// Requires AVX2.
func M256MinEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MovemaskEpi8: Create mask from the most significant bit of each 8-bit
// element in 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//
// Instruction: 'VPMOVMSKB'. Intrinsic: '_mm256_movemask_epi8'.
// Requires AVX2.
func M256MovemaskEpi8(a x86.M256i) int {
	panic("not implemented")
}


// M256MpsadbwEpu8: Compute the sum of absolute differences (SADs) of
// quadruplets of unsigned 8-bit integers in 'a' compared to those in 'b', and
// store the 16-bit results in 'dst'.
// 	Eight SADs are performed for each 128-bit lane using one quadruplet from
// 'b' and eight quadruplets from 'a'. One quadruplet is selected from 'b'
// starting at on the offset specified in 'imm8'. Eight quadruplets are formed
// from sequential 8-bit integers selected from 'a' starting at the offset
// specified in 'imm8'. 
//
//		MPSADBW(a[127:0], b[127:0], imm8[2:0]) {
//			i := imm8[2]*32
//			b_offset := imm8[1:0]*32
//			FOR j := 0 to 7
//				i := j*8
//				k := a_offset+i
//				l := b_offset
//				tmp[i+15:i] := ABS(a[k+7:k] - b[l+7:l]) + ABS(a[k+15:k+8] - b[l+15:l+8]) + ABS(a[k+23:k+16] - b[l+23:l+16]) + ABS(a[k+31:k+24] - b[l+31:l+24])
//			ENDFOR
//			RETURN tmp[127:0]
//		}
//		
//		dst[127:0] := MPSADBW(a[127:0], b[127:0], imm8[2:0])
//		dst[255:128] := MPSADBW(a[255:128], b[255:128], imm8[5:3])
//		dst[MAX:256] := 0
//
// Instruction: 'VMPSADBW'. Intrinsic: '_mm256_mpsadbw_epu8'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256MpsadbwEpu8(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulEpi32: Multiply the low 32-bit integers from each packed 64-bit
// element in 'a' and 'b', and store the signed 64-bit results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULDQ'. Intrinsic: '_mm256_mul_epi32'.
// Requires AVX2.
func M256MulEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulEpu32: Multiply the low unsigned 32-bit integers from each packed
// 64-bit element in 'a' and 'b', and store the unsigned 64-bit results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULUDQ'. Intrinsic: '_mm256_mul_epu32'.
// Requires AVX2.
func M256MulEpu32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm256_mulhi_epi16'.
// Requires AVX2.
func M256MulhiEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm256_mulhi_epu16'.
// Requires AVX2.
func M256MulhiEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//			dst[i+15:i] := tmp[16:1]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm256_mulhrs_epi16'.
// Requires AVX2.
func M256MulhrsEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[15:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm256_mullo_epi16'.
// Requires AVX2.
func M256MulloEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MulloEpi32: Multiply the packed 32-bit integers in 'a' and 'b',
// producing intermediate 64-bit integers, and store the low 32 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			tmp[63:0] := a[i+31:i] * b[i+31:i]
//			dst[i+31:i] := tmp[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLD'. Intrinsic: '_mm256_mullo_epi32'.
// Requires AVX2.
func M256MulloEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256OrSi256: Compute the bitwise OR of 256 bits (representing integer data)
// in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[255:0] := (a[255:0] OR b[255:0])
//		dst[MAX:256] := 0
//
// Instruction: 'VPOR'. Intrinsic: '_mm256_or_si256'.
// Requires AVX2.
func M256OrSi256(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256PacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
// 8-bit integers using signed saturation, and store the results in 'dst'. 
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
//		dst[135:128] := Saturate_Int16_To_Int8 (a[143:128])
//		dst[143:136] := Saturate_Int16_To_Int8 (a[159:144])
//		dst[151:144] := Saturate_Int16_To_Int8 (a[175:160])
//		dst[159:152] := Saturate_Int16_To_Int8 (a[191:176])
//		dst[167:160] := Saturate_Int16_To_Int8 (a[207:192])
//		dst[175:168] := Saturate_Int16_To_Int8 (a[223:208])
//		dst[183:176] := Saturate_Int16_To_Int8 (a[239:224])
//		dst[191:184] := Saturate_Int16_To_Int8 (a[255:240])
//		dst[199:192] := Saturate_Int16_To_Int8 (b[143:128])
//		dst[207:200] := Saturate_Int16_To_Int8 (b[159:144])
//		dst[215:208] := Saturate_Int16_To_Int8 (b[175:160])
//		dst[223:216] := Saturate_Int16_To_Int8 (b[191:176])
//		dst[231:224] := Saturate_Int16_To_Int8 (b[207:192])
//		dst[239:232] := Saturate_Int16_To_Int8 (b[223:208])
//		dst[247:240] := Saturate_Int16_To_Int8 (b[239:224])
//		dst[255:248] := Saturate_Int16_To_Int8 (b[255:240])
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm256_packs_epi16'.
// Requires AVX2.
func M256PacksEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256PacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using signed saturation, and store the results in 'dst'. 
//
//		dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		dst[143:128] := Saturate_Int32_To_Int16 (a[159:128])
//		dst[159:144] := Saturate_Int32_To_Int16 (a[191:160])
//		dst[175:160] := Saturate_Int32_To_Int16 (a[223:192])
//		dst[191:176] := Saturate_Int32_To_Int16 (a[255:224])
//		dst[207:192] := Saturate_Int32_To_Int16 (b[159:128])
//		dst[223:208] := Saturate_Int32_To_Int16 (b[191:160])
//		dst[239:224] := Saturate_Int32_To_Int16 (b[223:192])
//		dst[255:240] := Saturate_Int32_To_Int16 (b[255:224])
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm256_packs_epi32'.
// Requires AVX2.
func M256PacksEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256PackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
// 8-bit integers using unsigned saturation, and store the results in 'dst'. 
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
//		dst[135:128] := Saturate_Int16_To_UnsignedInt8 (a[143:128])
//		dst[143:136] := Saturate_Int16_To_UnsignedInt8 (a[159:144])
//		dst[151:144] := Saturate_Int16_To_UnsignedInt8 (a[175:160])
//		dst[159:152] := Saturate_Int16_To_UnsignedInt8 (a[191:176])
//		dst[167:160] := Saturate_Int16_To_UnsignedInt8 (a[207:192])
//		dst[175:168] := Saturate_Int16_To_UnsignedInt8 (a[223:208])
//		dst[183:176] := Saturate_Int16_To_UnsignedInt8 (a[239:224])
//		dst[191:184] := Saturate_Int16_To_UnsignedInt8 (a[255:240])
//		dst[199:192] := Saturate_Int16_To_UnsignedInt8 (b[143:128])
//		dst[207:200] := Saturate_Int16_To_UnsignedInt8 (b[159:144])
//		dst[215:208] := Saturate_Int16_To_UnsignedInt8 (b[175:160])
//		dst[223:216] := Saturate_Int16_To_UnsignedInt8 (b[191:176])
//		dst[231:224] := Saturate_Int16_To_UnsignedInt8 (b[207:192])
//		dst[239:232] := Saturate_Int16_To_UnsignedInt8 (b[223:208])
//		dst[247:240] := Saturate_Int16_To_UnsignedInt8 (b[239:224])
//		dst[255:248] := Saturate_Int16_To_UnsignedInt8 (b[255:240])
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm256_packus_epi16'.
// Requires AVX2.
func M256PackusEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256PackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using unsigned saturation, and store the results in 'dst'. 
//
//		dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		dst[143:128] := Saturate_Int32_To_UnsignedInt16 (a[159:128])
//		dst[159:144] := Saturate_Int32_To_UnsignedInt16 (a[191:160])
//		dst[175:160] := Saturate_Int32_To_UnsignedInt16 (a[223:192])
//		dst[191:176] := Saturate_Int32_To_UnsignedInt16 (a[255:224])
//		dst[207:192] := Saturate_Int32_To_UnsignedInt16 (b[159:128])
//		dst[223:208] := Saturate_Int32_To_UnsignedInt16 (b[191:160])
//		dst[239:224] := Saturate_Int32_To_UnsignedInt16 (b[223:192])
//		dst[255:240] := Saturate_Int32_To_UnsignedInt16 (b[255:224])
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm256_packus_epi32'.
// Requires AVX2.
func M256PackusEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Permute2x128Si256: Shuffle 128-bits (composed of integer data) selected
// by 'imm8' from 'a' and 'b', and store the results in 'dst'. 
//
//		SELECT4(src1, src2, control){
//			CASE(control[1:0])
//			0:	tmp[127:0] := src1[127:0]
//			1:	tmp[127:0] := src1[255:128]
//			2:	tmp[127:0] := src2[127:0]
//			3:	tmp[127:0] := src2[255:128]
//			ESAC
//			IF control[3]
//				tmp[127:0] := 0
//			FI
//			RETURN tmp[127:0]
//		}
//		
//		dst[127:0] := SELECT4(a[255:0], b[255:0], imm8[3:0])
//		dst[255:128] := SELECT4(a[255:0], b[255:0], imm8[7:4])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERM2I128'. Intrinsic: '_mm256_permute2x128_si256'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256Permute2x128Si256(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256Permute4x64Epi64: Shuffle 64-bit integers in 'a' across lanes using the
// control in 'imm8', and store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[63:0] := src[63:0]
//			1:	tmp[63:0] := src[127:64]
//			2:	tmp[63:0] := src[191:128]
//			3:	tmp[63:0] := src[255:192]
//			ESAC
//			RETURN tmp[63:0]
//		}
//		
//		dst[63:0] := SELECT4(a[255:0], imm8[1:0])
//		dst[127:64] := SELECT4(a[255:0], imm8[3:2])
//		dst[191:128] := SELECT4(a[255:0], imm8[5:4])
//		dst[255:192] := SELECT4(a[255:0], imm8[7:6])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMQ'. Intrinsic: '_mm256_permute4x64_epi64'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256Permute4x64Epi64(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256Permute4x64Pd: Shuffle double-precision (64-bit) floating-point elements
// in 'a' across lanes using the control in 'imm8', and store the results in
// 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[63:0] := src[63:0]
//			1:	tmp[63:0] := src[127:64]
//			2:	tmp[63:0] := src[191:128]
//			3:	tmp[63:0] := src[255:192]
//			ESAC
//			RETURN tmp[63:0]
//		}
//		
//		dst[63:0] := SELECT4(a[255:0], imm8[1:0])
//		dst[127:64] := SELECT4(a[255:0], imm8[3:2])
//		dst[191:128] := SELECT4(a[255:0], imm8[5:4])
//		dst[255:192] := SELECT4(a[255:0], imm8[7:6])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMPD'. Intrinsic: '_mm256_permute4x64_pd'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256Permute4x64Pd(a x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


// M256Permutevar8x32Epi32: Shuffle 32-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			id := idx[i+2:i]*32
//			dst[i+31:i] := a[id+31:id]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMD'. Intrinsic: '_mm256_permutevar8x32_epi32'.
// Requires AVX2.
func M256Permutevar8x32Epi32(a x86.M256i, idx x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Permutevar8x32Ps: Shuffle single-precision (32-bit) floating-point
// elements in 'a' across lanes using the corresponding index in 'idx'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			id := idx[i+2:i]*32
//			dst[i+31:i] := a[id+31:id]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMPS'. Intrinsic: '_mm256_permutevar8x32_ps'.
// Requires AVX2.
func M256Permutevar8x32Ps(a x86.M256, idx x86.M256i) (dst x86.M256) {
	panic("not implemented")
}


// M256SadEpu8: Compute the absolute differences of packed unsigned 8-bit
// integers in 'a' and 'b', then horizontally sum each consecutive 8
// differences to produce four unsigned 16-bit integers, and pack these
// unsigned 16-bit integers in the low 16 bits of 64-bit elements in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		FOR j := 0 to 4
//			i := j*64
//			dst[i+15:i] := tmp[i+7:i] + tmp[i+15:i+8] + tmp[i+23:i+16] + tmp[i+31:i+24] + tmp[i+39:i+32] + tmp[i+47:i+40] + tmp[i+55:i+48] + tmp[i+63:i+56]
//			dst[i+63:i+16] := 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSADBW'. Intrinsic: '_mm256_sad_epu8'.
// Requires AVX2.
func M256SadEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256ShuffleEpi32: Shuffle 32-bit integers in 'a' within 128-bit lanes using
// the control in 'imm8', and store the results in 'dst'. 
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
//		dst[159:128] := SELECT4(a[255:128], imm8[1:0])
//		dst[191:160] := SELECT4(a[255:128], imm8[3:2])
//		dst[223:192] := SELECT4(a[255:128], imm8[5:4])
//		dst[255:224] := SELECT4(a[255:128], imm8[7:6])
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFD'. Intrinsic: '_mm256_shuffle_epi32'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256ShuffleEpi32(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256ShuffleEpi8: Shuffle 8-bit integers in 'a' within 128-bit lanes
// according to shuffle control mask in the corresponding 8-bit element of 'b',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF b[i+7] == 1
//				dst[i+7:i] := 0
//			ELSE
//				index[3:0] := b[i+3:i]
//				dst[i+7:i] := a[index*8+7:index*8]
//			FI
//			IF b[128+i+7] == 1
//				dst[128+i+7:i] := 0
//			ELSE
//				index[3:0] := b[128+i+3:128+i]
//				dst[128+i+7:i] := a[128+index*8+7:128+index*8]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm256_shuffle_epi8'.
// Requires AVX2.
func M256ShuffleEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256ShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the high 64
// bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst'. 
//
//		dst[63:0] := a[63:0]
//		dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		dst[191:128] := a[191:128]
//		dst[207:192] := (a >> (imm8[1:0] * 16))[207:192]
//		dst[223:208] := (a >> (imm8[3:2] * 16))[207:192]
//		dst[239:224] := (a >> (imm8[5:4] * 16))[207:192]
//		dst[255:240] := (a >> (imm8[7:6] * 16))[207:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm256_shufflehi_epi16'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256ShufflehiEpi16(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256ShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the low 64
// bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst'. 
//
//		dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		dst[127:64] := a[127:64]
//		dst[143:128] := (a >> (imm8[1:0] * 16))[143:128]
//		dst[159:144] := (a >> (imm8[3:2] * 16))[143:128]
//		dst[175:160] := (a >> (imm8[5:4] * 16))[143:128]
//		dst[191:176] := (a >> (imm8[7:6] * 16))[143:128]
//		dst[255:192] := a[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm256_shufflelo_epi16'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256ShuffleloEpi16(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SignEpi16: Negate packed 16-bit integers in 'a' when the corresponding
// signed 16-bit integer in 'b' is negative, and store the results in 'dst'.
// Element in 'dst' are zeroed out when the corresponding element in 'b' is
// zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF b[i+15:i] < 0
//				dst[i+15:i] := NEG(a[i+15:i])
//			ELSE IF b[i+15:i] = 0
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSIGNW'. Intrinsic: '_mm256_sign_epi16'.
// Requires AVX2.
func M256SignEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SignEpi32: Negate packed 32-bit integers in 'a' when the corresponding
// signed 32-bit integer in 'b' is negative, and store the results in 'dst'.
// Element in 'dst' are zeroed out when the corresponding element in 'b' is
// zero. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF b[i+31:i] < 0
//				dst[i+31:i] := NEG(a[i+31:i])
//			ELSE IF b[i+31:i] = 0
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSIGND'. Intrinsic: '_mm256_sign_epi32'.
// Requires AVX2.
func M256SignEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SignEpi8: Negate packed 8-bit integers in 'a' when the corresponding
// signed 8-bit integer in 'b' is negative, and store the results in 'dst'.
// Element in 'dst' are zeroed out when the corresponding element in 'b' is
// zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF b[i+7:i] < 0
//				dst[i+7:i] := NEG(a[i+7:i])
//			ELSE IF b[i+7:i] = 0
//				dst[i+7:i] := 0
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSIGNB'. Intrinsic: '_mm256_sign_epi8'.
// Requires AVX2.
func M256SignEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm256_sll_epi16'.
// Requires AVX2.
func M256SllEpi16(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SllEpi32: Shift packed 32-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLD'. Intrinsic: '_mm256_sll_epi32'.
// Requires AVX2.
func M256SllEpi32(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SllEpi64: Shift packed 64-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF count[63:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] << count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLQ'. Intrinsic: '_mm256_sll_epi64'.
// Requires AVX2.
func M256SllEpi64(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm256_slli_epi16'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SlliEpi16(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SlliEpi32: Shift packed 32-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLD'. Intrinsic: '_mm256_slli_epi32'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SlliEpi32(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SlliEpi64: Shift packed 64-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF imm8[7:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] << imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLQ'. Intrinsic: '_mm256_slli_epi64'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SlliEpi64(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SlliSi256: Shift 128-bit lanes in 'a' left by 'imm8' bytes while
// shifting in zeros, and store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] << (tmp*8)
//		dst[255:128] := a[255:128] << (tmp*8)
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLDQ'. Intrinsic: '_mm256_slli_si256'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SlliSi256(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// SllvEpi32: Shift packed 32-bit integers in 'a' left by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] << count[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLVD'. Intrinsic: '_mm_sllv_epi32'.
// Requires AVX2.
func SllvEpi32(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256SllvEpi32: Shift packed 32-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] << count[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLVD'. Intrinsic: '_mm256_sllv_epi32'.
// Requires AVX2.
func M256SllvEpi32(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// SllvEpi64: Shift packed 64-bit integers in 'a' left by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ZeroExtend(a[i+63:i] << count[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLVQ'. Intrinsic: '_mm_sllv_epi64'.
// Requires AVX2.
func SllvEpi64(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256SllvEpi64: Shift packed 64-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ZeroExtend(a[i+63:i] << count[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLVQ'. Intrinsic: '_mm256_sllv_epi64'.
// Requires AVX2.
func M256SllvEpi64(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm256_sra_epi16'.
// Requires AVX2.
func M256SraEpi16(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SraEpi32: Shift packed 32-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAD'. Intrinsic: '_mm256_sra_epi32'.
// Requires AVX2.
func M256SraEpi32(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm256_srai_epi16'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SraiEpi16(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SraiEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAD'. Intrinsic: '_mm256_srai_epi32'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SraiEpi32(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// SravEpi32: Shift packed 32-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in sign bits, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SignExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAVD'. Intrinsic: '_mm_srav_epi32'.
// Requires AVX2.
func SravEpi32(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256SravEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SignExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAVD'. Intrinsic: '_mm256_srav_epi32'.
// Requires AVX2.
func M256SravEpi32(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm256_srl_epi16'.
// Requires AVX2.
func M256SrlEpi16(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrlEpi32: Shift packed 32-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF count[63:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLD'. Intrinsic: '_mm256_srl_epi32'.
// Requires AVX2.
func M256SrlEpi32(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrlEpi64: Shift packed 64-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF count[63:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLQ'. Intrinsic: '_mm256_srl_epi64'.
// Requires AVX2.
func M256SrlEpi64(a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm256_srli_epi16'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SrliEpi16(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrliEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLD'. Intrinsic: '_mm256_srli_epi32'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SrliEpi32(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrliEpi64: Shift packed 64-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF imm8[7:0] > 63
//				dst[i+63:i] := 0
//			ELSE
//				dst[i+63:i] := ZeroExtend(a[i+63:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLQ'. Intrinsic: '_mm256_srli_epi64'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SrliEpi64(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrliSi256: Shift 128-bit lanes in 'a' right by 'imm8' bytes while
// shifting in zeros, and store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] >> (tmp*8)
//		dst[255:128] := a[255:128] >> (tmp*8)
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLDQ'. Intrinsic: '_mm256_srli_si256'.
// Requires AVX2.
//
// FIXME: Requires compiler support (has immediate)
func M256SrliSi256(a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// SrlvEpi32: Shift packed 32-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLVD'. Intrinsic: '_mm_srlv_epi32'.
// Requires AVX2.
func SrlvEpi32(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256SrlvEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLVD'. Intrinsic: '_mm256_srlv_epi32'.
// Requires AVX2.
func M256SrlvEpi32(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// SrlvEpi64: Shift packed 64-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ZeroExtend(a[i+63:i] >> count[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLVQ'. Intrinsic: '_mm_srlv_epi64'.
// Requires AVX2.
func SrlvEpi64(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256SrlvEpi64: Shift packed 64-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ZeroExtend(a[i+63:i] >> count[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLVQ'. Intrinsic: '_mm256_srlv_epi64'.
// Requires AVX2.
func M256SrlvEpi64(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256StreamLoadSi256: Load 256-bits of integer data from memory into 'dst'
// using a non-temporal memory hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVNTDQA'. Intrinsic: '_mm256_stream_load_si256'.
// Requires AVX2.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StreamLoadSi256(mem_addr *x86.M256iConst) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := a[i+15:i] - b[i+15:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm256_sub_epi16'.
// Requires AVX2.
func M256SubEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubEpi32: Subtract packed 32-bit integers in 'b' from packed 32-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBD'. Intrinsic: '_mm256_sub_epi32'.
// Requires AVX2.
func M256SubEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubEpi64: Subtract packed 64-bit integers in 'b' from packed 64-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBQ'. Intrinsic: '_mm256_sub_epi64'.
// Requires AVX2.
func M256SubEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := a[i+7:i] - b[i+7:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm256_sub_epi8'.
// Requires AVX2.
func M256SubEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm256_subs_epi16'.
// Requires AVX2.
func M256SubsEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm256_subs_epi8'.
// Requires AVX2.
func M256SubsEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm256_subs_epu16'.
// Requires AVX2.
func M256SubsEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm256_subs_epu8'.
// Requires AVX2.
func M256SubsEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackhiEpi16: Unpack and interleave 16-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_HIGH_WORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm256_unpackhi_epi16'.
// Requires AVX2.
func M256UnpackhiEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackhiEpi32: Unpack and interleave 32-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_HIGH_DWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHDQ'. Intrinsic: '_mm256_unpackhi_epi32'.
// Requires AVX2.
func M256UnpackhiEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackhiEpi64: Unpack and interleave 64-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[127:64] 
//			dst[127:64] := src2[127:64] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_QWORDS(a[127:0], b[127:0])
//		dst[255:128] := INTERLEAVE_HIGH_QWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHQDQ'. Intrinsic: '_mm256_unpackhi_epi64'.
// Requires AVX2.
func M256UnpackhiEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackhiEpi8: Unpack and interleave 8-bit integers from the high half of
// each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_HIGH_BYTES(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm256_unpackhi_epi8'.
// Requires AVX2.
func M256UnpackhiEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackloEpi16: Unpack and interleave 16-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_WORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm256_unpacklo_epi16'.
// Requires AVX2.
func M256UnpackloEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackloEpi32: Unpack and interleave 32-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_DWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLDQ'. Intrinsic: '_mm256_unpacklo_epi32'.
// Requires AVX2.
func M256UnpackloEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackloEpi64: Unpack and interleave 64-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[63:0] 
//			dst[127:64] := src2[63:0] 
//			RETURN dst[127:0]
//		}
//		
//		dst[127:0] := INTERLEAVE_QWORDS(a[127:0], b[127:0])
//		dst[255:128] := INTERLEAVE_QWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLQDQ'. Intrinsic: '_mm256_unpacklo_epi64'.
// Requires AVX2.
func M256UnpackloEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256UnpackloEpi8: Unpack and interleave 8-bit integers from the low half of
// each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_BYTES(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm256_unpacklo_epi8'.
// Requires AVX2.
func M256UnpackloEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256XorSi256: Compute the bitwise XOR of 256 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[255:0] := (a[255:0] XOR b[255:0])
//		dst[MAX:256] := 0
//
// Instruction: 'VPXOR'. Intrinsic: '_mm256_xor_si256'.
// Requires AVX2.
func M256XorSi256(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}

