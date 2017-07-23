// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package ssse3

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AbsEpi16: Compute the absolute value of packed 16-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := ABS(a[i+15:i])
//		ENDFOR
//
// Instruction: 'PABSW'. Intrinsic: '_mm_abs_epi16'.
// Requires SSSE3.
func AbsEpi16(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AbsEpi32: Compute the absolute value of packed 32-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ABS(a[i+31:i])
//		ENDFOR
//
// Instruction: 'PABSD'. Intrinsic: '_mm_abs_epi32'.
// Requires SSSE3.
func AbsEpi32(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AbsEpi8: Compute the absolute value of packed 8-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			dst[i+7:i] := ABS(a[i+7:i])
//		ENDFOR
//
// Instruction: 'PABSB'. Intrinsic: '_mm_abs_epi8'.
// Requires SSSE3.
func AbsEpi8(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func AbsPi16(a x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AbsPi32(a x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func AbsPi8(a x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// AlignrEpi8: Concatenate 16-byte blocks in 'a' and 'b' into a 32-byte
// temporary result, shift the result right by 'count' bytes, and store the low
// 16 bytes in 'dst'. 
//
//		tmp[255:0] := ((a[127:0] << 128) OR b[127:0]) >> (count[7:0]*8)
//		dst[127:0] := tmp[127:0]
//
// Instruction: 'PALIGNR'. Intrinsic: '_mm_alignr_epi8'.
// Requires SSSE3.
func AlignrEpi8(a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// AlignrPi8: Concatenate 8-byte blocks in 'a' and 'b' into a 16-byte temporary
// result, shift the result right by 'count' bytes, and store the low 16 bytes
// in 'dst'. 
//
//		tmp[127:0] := ((a[63:0] << 64) OR b[63:0]) >> (count[7:0]*8)
//		dst[63:0] := tmp[63:0]
//
// Instruction: 'PALIGNR'. Intrinsic: '_mm_alignr_pi8'.
// Requires SSSE3.
func AlignrPi8(a x86.M64, b x86.M64, count int) (dst x86.M64) {
	panic("not implemented")
}


// HaddEpi16: Horizontally add adjacent pairs of 16-bit integers in 'a' and
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
//
// Instruction: 'PHADDW'. Intrinsic: '_mm_hadd_epi16'.
// Requires SSSE3.
func HaddEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// HaddEpi32: Horizontally add adjacent pairs of 32-bit integers in 'a' and
// 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := a[127:96] + a[95:64]
//		dst[95:64] := b[63:32] + b[31:0]
//		dst[127:96] := b[127:96] + b[95:64]
//
// Instruction: 'PHADDD'. Intrinsic: '_mm_hadd_epi32'.
// Requires SSSE3.
func HaddEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func HaddPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// HaddPi32: Horizontally add adjacent pairs of 32-bit integers in 'a' and 'b',
// and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := b[63:32] + b[31:0]
//
// Instruction: 'PHADDW'. Intrinsic: '_mm_hadd_pi32'.
// Requires SSSE3.
func HaddPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// HaddsEpi16: Horizontally add adjacent pairs of 16-bit integers in 'a' and
// 'b' using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[31:16] + a[15:0])
//		dst[31:16] = Saturate_To_Int16(a[63:48] + a[47:32])
//		dst[47:32] = Saturate_To_Int16(a[95:80] + a[79:64])
//		dst[63:48] = Saturate_To_Int16(a[127:112] + a[111:96])
//		dst[79:64] = Saturate_To_Int16(b[31:16] + b[15:0])
//		dst[95:80] = Saturate_To_Int16(b[63:48] + b[47:32])
//		dst[111:96] = Saturate_To_Int16(b[95:80] + b[79:64])
//		dst[127:112] = Saturate_To_Int16(b[127:112] + b[111:96])
//
// Instruction: 'PHADDSW'. Intrinsic: '_mm_hadds_epi16'.
// Requires SSSE3.
func HaddsEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func HaddsPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// HsubEpi16: Horizontally subtract adjacent pairs of 16-bit integers in 'a'
// and 'b', and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0] := a[15:0] - a[31:16]
//		dst[31:16] := a[47:32] - a[63:48]
//		dst[47:32] := a[79:64] - a[95:80]
//		dst[63:48] := a[111:96] - a[127:112]
//		dst[79:64] := b[15:0] - b[31:16]
//		dst[95:80] := b[47:32] - b[63:48]
//		dst[111:96] := b[79:64] - b[95:80]
//		dst[127:112] := b[111:96] - b[127:112]
//
// Instruction: 'PHSUBW'. Intrinsic: '_mm_hsub_epi16'.
// Requires SSSE3.
func HsubEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// HsubEpi32: Horizontally subtract adjacent pairs of 32-bit integers in 'a'
// and 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := a[95:64] - a[127:96]
//		dst[95:64] := b[31:0] - b[63:32]
//		dst[127:96] := b[95:64] - b[127:96]
//
// Instruction: 'PHSUBD'. Intrinsic: '_mm_hsub_epi32'.
// Requires SSSE3.
func HsubEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func HsubPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// HsubPi32: Horizontally subtract adjacent pairs of 32-bit integers in 'a' and
// 'b', and pack the signed 32-bit results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := b[31:0] - b[63:32]
//
// Instruction: 'PHSUBD'. Intrinsic: '_mm_hsub_pi32'.
// Requires SSSE3.
func HsubPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// HsubsEpi16: Horizontally subtract adjacent pairs of 16-bit integers in 'a'
// and 'b' using saturation, and pack the signed 16-bit results in 'dst'. 
//
//		dst[15:0]= Saturate_To_Int16(a[15:0] - a[31:16])
//		dst[31:16] = Saturate_To_Int16(a[47:32] - a[63:48])
//		dst[47:32] = Saturate_To_Int16(a[79:64] - a[95:80])
//		dst[63:48] = Saturate_To_Int16(a[111:96] - a[127:112])
//		dst[79:64] = Saturate_To_Int16(b[15:0] - b[31:16])
//		dst[95:80] = Saturate_To_Int16(b[47:32] - b[63:48])
//		dst[111:96] = Saturate_To_Int16(b[79:64] - b[95:80])
//		dst[127:112] = Saturate_To_Int16(b[111:96] - b[127:112])
//
// Instruction: 'PHSUBSW'. Intrinsic: '_mm_hsubs_epi16'.
// Requires SSSE3.
func HsubsEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func HsubsPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MaddubsEpi16: Vertically multiply each unsigned 8-bit integer from 'a' with
// the corresponding signed 8-bit integer from 'b', producing intermediate
// signed 16-bit integers. Horizontally add adjacent pairs of intermediate
// signed 16-bit integers, and pack the saturated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//		ENDFOR
//
// Instruction: 'PMADDUBSW'. Intrinsic: '_mm_maddubs_epi16'.
// Requires SSSE3.
func MaddubsEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaddubsPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//			dst[i+15:i] := tmp[16:1]
//		ENDFOR
//
// Instruction: 'PMULHRSW'. Intrinsic: '_mm_mulhrs_epi16'.
// Requires SSSE3.
func MulhrsEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func MulhrsPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// ShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF b[i+7] == 1
//				dst[i+7:i] := 0
//			ELSE
//				index[3:0] := b[i+3:i]
//				dst[i+7:i] := a[index*8+7:index*8]
//			FI
//		ENDFOR
//
// Instruction: 'PSHUFB'. Intrinsic: '_mm_shuffle_epi8'.
// Requires SSSE3.
func ShuffleEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func ShufflePi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// SignEpi16: Negate packed 16-bit integers in 'a' when the corresponding
// signed 16-bit integer in 'b' is negative, and store the results in 'dst'.
// Element in 'dst' are zeroed out when the corresponding element in 'b' is
// zero. 
//
//		FOR j := 0 to 7
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
// Instruction: 'PSIGNW'. Intrinsic: '_mm_sign_epi16'.
// Requires SSSE3.
func SignEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// SignEpi32: Negate packed 32-bit integers in 'a' when the corresponding
// signed 32-bit integer in 'b' is negative, and store the results in 'dst'.
// Element in 'dst' are zeroed out when the corresponding element in 'b' is
// zero. 
//
//		FOR j := 0 to 3
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
// Instruction: 'PSIGND'. Intrinsic: '_mm_sign_epi32'.
// Requires SSSE3.
func SignEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// SignEpi8: Negate packed 8-bit integers in 'a' when the corresponding signed
// 8-bit integer in 'b' is negative, and store the results in 'dst'. Element in
// 'dst' are zeroed out when the corresponding element in 'b' is zero. 
//
//		FOR j := 0 to 15
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
// Instruction: 'PSIGNB'. Intrinsic: '_mm_sign_epi8'.
// Requires SSSE3.
func SignEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func SignPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SignPi32(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


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
func SignPi8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}

