// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package avx512vl

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// MaskAddPd: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm_mask_add_pd'.
// Requires AVX512VL.
func MaskAddPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// MaskzAddPd: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm_maskz_add_pd'.
// Requires AVX512VL.
func MaskzAddPd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// M256MaskAddPd: Add packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm256_mask_add_pd'.
// Requires AVX512VL.
func M256MaskAddPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


// M256MaskzAddPd: Add packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm256_maskz_add_pd'.
// Requires AVX512VL.
func M256MaskzAddPd(k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


// MaskAddPs: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm_mask_add_ps'.
// Requires AVX512VL.
func MaskAddPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MaskzAddPs: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm_maskz_add_ps'.
// Requires AVX512VL.
func MaskzAddPs(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// M256MaskAddPs: Add packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm256_mask_add_ps'.
// Requires AVX512VL.
func M256MaskAddPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


// M256MaskzAddPs: Add packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm256_maskz_add_ps'.
// Requires AVX512VL.
func M256MaskzAddPs(k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


// AlignrEpi32: Concatenate 'a' and 'b' into a 32-byte immediate result, shift
// the result right by 'count' 32-bit elements, and store the low 16 bytes (4
// elements) in 'dst'. 
//
//		temp[255:128] := a[127:0]
//		temp[127:0] := b[127:0]
//		temp[255:0] := temp[255:0] >> (32*count)
//		dst[127:0] := temp[127:0]
//		dst[MAX:128] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm_alignr_epi32'.
// Requires AVX512VL.
func AlignrEpi32(a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskAlignrEpi32: Concatenate 'a' and 'b' into a 32-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 16
// bytes (4 elements) in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		temp[255:128] := a[127:0]
//		temp[127:0] := b[127:0]
//		temp[255:0] := temp[255:0] >> (32*count)
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := temp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm_mask_alignr_epi32'.
// Requires AVX512VL.
func MaskAlignrEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAlignrEpi32: Concatenate 'a' and 'b' into a 32-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 16
// bytes (4 elements) in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		temp[255:128] := a[127:0]
//		temp[127:0] := b[127:0]
//		temp[255:0] := temp[255:0] >> (32*count)
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := temp[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm_maskz_alignr_epi32'.
// Requires AVX512VL.
func MaskzAlignrEpi32(k x86.Mmask8, a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// M256AlignrEpi32: Concatenate 'a' and 'b' into a 64-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 32
// bytes (8 elements) in 'dst'. 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (32*count)
//		dst[255:0] := temp[255:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm256_alignr_epi32'.
// Requires AVX512VL.
func M256AlignrEpi32(a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskAlignrEpi32: Concatenate 'a' and 'b' into a 64-byte immediate
// result, shift the result right by 'count' 32-bit elements, and store the low
// 32 bytes (8 elements) in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (32*count)
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := temp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm256_mask_alignr_epi32'.
// Requires AVX512VL.
func M256MaskAlignrEpi32(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAlignrEpi32: Concatenate 'a' and 'b' into a 64-byte immediate
// result, shift the result right by 'count' 32-bit elements, and store the low
// 32 bytes (8 elements) in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (32*count)
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := temp[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm256_maskz_alignr_epi32'.
// Requires AVX512VL.
func M256MaskzAlignrEpi32(k x86.Mmask8, a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// AlignrEpi64: Concatenate 'a' and 'b' into a 32-byte immediate result, shift
// the result right by 'count' 64-bit elements, and store the low 16 bytes (2
// elements) in 'dst'. 
//
//		temp[255:128] := a[127:0]
//		temp[127:0] := b[127:0]
//		temp[255:0] := temp[255:0] >> (64*count)
//		dst[127:0] := temp[127:0]
//		dst[MAX:128] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm_alignr_epi64'.
// Requires AVX512VL.
func AlignrEpi64(a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskAlignrEpi64: Concatenate 'a' and 'b' into a 32-byte immediate result,
// shift the result right by 'count' 64-bit elements, and store the low 16
// bytes (2 elements) in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		temp[255:128] := a[127:0]
//		temp[127:0] := b[127:0]
//		temp[255:0] := temp[255:0] >> (64*count)
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := temp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm_mask_alignr_epi64'.
// Requires AVX512VL.
func MaskAlignrEpi64(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAlignrEpi64: Concatenate 'a' and 'b' into a 32-byte immediate result,
// shift the result right by 'count' 64-bit elements, and store the low 16
// bytes (2 elements) in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		temp[255:128] := a[127:0]
//		temp[127:0] := b[127:0]
//		temp[255:0] := temp[255:0] >> (64*count)
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := temp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm_maskz_alignr_epi64'.
// Requires AVX512VL.
func MaskzAlignrEpi64(k x86.Mmask8, a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// M256AlignrEpi64: Concatenate 'a' and 'b' into a 64-byte immediate result,
// shift the result right by 'count' 64-bit elements, and store the low 32
// bytes (4 elements) in 'dst'. 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (64*count)
//		dst[255:0] := temp[255:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm256_alignr_epi64'.
// Requires AVX512VL.
func M256AlignrEpi64(a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskAlignrEpi64: Concatenate 'a' and 'b' into a 64-byte immediate
// result, shift the result right by 'count' 64-bit elements, and store the low
// 32 bytes (4 elements) in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (64*count)
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := temp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm256_mask_alignr_epi64'.
// Requires AVX512VL.
func M256MaskAlignrEpi64(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAlignrEpi64: Concatenate 'a' and 'b' into a 64-byte immediate
// result, shift the result right by 'count' 64-bit elements, and store the low
// 32 bytes (4 elements) in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (64*count)
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := temp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm256_maskz_alignr_epi64'.
// Requires AVX512VL.
func M256MaskzAlignrEpi64(k x86.Mmask8, a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// Madd52hiEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//			dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm_madd52hi_epu64'.
// Requires AVX512VL.
func Madd52hiEpu64(a x86.M128i, b x86.M128i, c x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskMadd52hiEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'a' when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm_mask_madd52hi_epu64'.
// Requires AVX512VL.
func MaskMadd52hiEpu64(a x86.M128i, k x86.Mmask8, b x86.M128i, c x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMadd52hiEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm_maskz_madd52hi_epu64'.
// Requires AVX512VL.
func MaskzMadd52hiEpu64(k x86.Mmask8, a x86.M128i, b x86.M128i, c x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256Madd52hiEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//			dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm256_madd52hi_epu64'.
// Requires AVX512VL.
func M256Madd52hiEpu64(a x86.M256i, b x86.M256i, c x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskMadd52hiEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// high 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm256_mask_madd52hi_epu64'.
// Requires AVX512VL.
func M256MaskMadd52hiEpu64(a x86.M256i, k x86.Mmask8, b x86.M256i, c x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMadd52hiEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// high 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm256_maskz_madd52hi_epu64'.
// Requires AVX512VL.
func M256MaskzMadd52hiEpu64(k x86.Mmask8, a x86.M256i, b x86.M256i, c x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// Madd52loEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//			dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm_madd52lo_epu64'.
// Requires AVX512VL.
func Madd52loEpu64(a x86.M128i, b x86.M128i, c x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskMadd52loEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'a' when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm_mask_madd52lo_epu64'.
// Requires AVX512VL.
func MaskMadd52loEpu64(a x86.M128i, k x86.Mmask8, b x86.M128i, c x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMadd52loEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm_maskz_madd52lo_epu64'.
// Requires AVX512VL.
func MaskzMadd52loEpu64(k x86.Mmask8, a x86.M128i, b x86.M128i, c x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256Madd52loEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//			dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm256_madd52lo_epu64'.
// Requires AVX512VL.
func M256Madd52loEpu64(a x86.M256i, b x86.M256i, c x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskMadd52loEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// low 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm256_mask_madd52lo_epu64'.
// Requires AVX512VL.
func M256MaskMadd52loEpu64(a x86.M256i, k x86.Mmask8, b x86.M256i, c x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMadd52loEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// low 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm256_maskz_madd52lo_epu64'.
// Requires AVX512VL.
func M256MaskzMadd52loEpu64(k x86.Mmask8, a x86.M256i, b x86.M256i, c x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// MovmEpi8: Set each packed 8-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := 0xFF
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVM2B'. Intrinsic: '_mm_movm_epi8'.
// Requires AVX512VL.
func MovmEpi8(k x86.Mmask16) (dst x86.M128i) {
	panic("not implemented")
}


// MaskMultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR i := 0 to 1
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				IF k[i*8+j]
//					dst[q+j*8+7:q+j*8] := tmp8[7:0]
//				ELSE
//					dst[q+j*8+7:q+j*8] := src[q+j*8+7:q+j*8]
//				FI
//			ENDFOR
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm_mask_multishift_epi64_epi8'.
// Requires AVX512VL.
func MaskMultishiftEpi64Epi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR i := 0 to 1
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				IF k[i*8+j]
//					dst[q+j*8+7:q+j*8] := tmp8[7:0]
//				ELSE
//					dst[q+j*8+7:q+j*8] := 0
//				FI
//			ENDFOR
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm_maskz_multishift_epi64_epi8'.
// Requires AVX512VL.
func MaskzMultishiftEpi64Epi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst'. 
//
//		FOR i := 0 to 1
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				dst[q+j*8+7:q+j*8] := tmp8[7:0]
//			ENDFOR
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm_multishift_epi64_epi8'.
// Requires AVX512VL.
func MultishiftEpi64Epi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMultishiftEpi64Epi8: For each 64-bit element in 'b', select 8
// unaligned bytes using a byte-granular shift control within the corresponding
// 64-bit element of 'a', and store the 8 assembled bytes to the corresponding
// 64-bit element of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR i := 0 to 3
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				IF k[i*8+j]
//					dst[q+j*8+7:q+j*8] := tmp8[7:0]
//				ELSE
//					dst[q+j*8+7:q+j*8] := src[q+j*8+7:q+j*8]
//				FI
//			ENDFOR
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm256_mask_multishift_epi64_epi8'.
// Requires AVX512VL.
func M256MaskMultishiftEpi64Epi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMultishiftEpi64Epi8: For each 64-bit element in 'b', select 8
// unaligned bytes using a byte-granular shift control within the corresponding
// 64-bit element of 'a', and store the 8 assembled bytes to the corresponding
// 64-bit element of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR i := 0 to 3
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				IF k[i*8+j]
//					dst[q+j*8+7:q+j*8] := tmp8[7:0]
//				ELSE
//					dst[q+j*8+7:q+j*8] := 0
//				FI
//			ENDFOR
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm256_maskz_multishift_epi64_epi8'.
// Requires AVX512VL.
func M256MaskzMultishiftEpi64Epi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst'. 
//
//		FOR i := 0 to 3
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				dst[q+j*8+7:q+j*8] := tmp8[7:0]
//			ENDFOR
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm256_multishift_epi64_epi8'.
// Requires AVX512VL.
func M256MultishiftEpi64Epi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// MaskPermutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+3:i]
//				dst[i+7:i] := idx[i+4] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMT2B'. Intrinsic: '_mm_mask_permutex2var_epi8'.
// Requires AVX512VL.
func MaskPermutex2varEpi8(a x86.M128i, k x86.Mmask16, idx x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// Mask2Permutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+3:i]
//				dst[i+7:i] := idx[i+4] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMI2B'. Intrinsic: '_mm_mask2_permutex2var_epi8'.
// Requires AVX512VL.
func Mask2Permutex2varEpi8(a x86.M128i, idx x86.M128i, k x86.Mmask16, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPermutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+3:i]
//				dst[i+7:i] := idx[i+4] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMI2B, VPERMT2B'. Intrinsic: '_mm_maskz_permutex2var_epi8'.
// Requires AVX512VL.
func MaskzPermutex2varEpi8(k x86.Mmask16, a x86.M128i, idx x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// Permutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			off := 8*idx[i+3:i]
//			dst[i+7:i] := idx[i+4] ? b[off+7:off] : a[off+7:off]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMI2B'. Intrinsic: '_mm_permutex2var_epi8'.
// Requires AVX512VL.
func Permutex2varEpi8(a x86.M128i, idx x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPermutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+4:i]
//				dst[i+7:i] := idx[i+5] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMT2B'. Intrinsic: '_mm256_mask_permutex2var_epi8'.
// Requires AVX512VL.
func M256MaskPermutex2varEpi8(a x86.M256i, k x86.Mmask32, idx x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Mask2Permutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+4:i]
//				dst[i+7:i] := idx[i+5] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMI2B'. Intrinsic: '_mm256_mask2_permutex2var_epi8'.
// Requires AVX512VL.
func M256Mask2Permutex2varEpi8(a x86.M256i, idx x86.M256i, k x86.Mmask32, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPermutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+4:i]
//				dst[i+7:i] := idx[i+5] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMI2B, VPERMT2B'. Intrinsic: '_mm256_maskz_permutex2var_epi8'.
// Requires AVX512VL.
func M256MaskzPermutex2varEpi8(k x86.Mmask32, a x86.M256i, idx x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Permutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			off := 8*idx[i+4:i]
//			dst[i+7:i] := idx[i+6] ? b[off+5:off] : a[off+7:off]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMI2B'. Intrinsic: '_mm256_permutex2var_epi8'.
// Requires AVX512VL.
func M256Permutex2varEpi8(a x86.M256i, idx x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// MaskPermutexvarEpi8: Shuffle 8-bit integers in 'a' using the corresponding
// index in 'idx', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			id := idx[i+3:i]*8
//			IF k[j]
//				dst[i+7:i] := a[id+7:id]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm_mask_permutexvar_epi8'.
// Requires AVX512VL.
func MaskPermutexvarEpi8(src x86.M128i, k x86.Mmask16, idx x86.M128i, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPermutexvarEpi8: Shuffle 8-bit integers in 'a' using the corresponding
// index in 'idx', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			id := idx[i+3:i]*8
//			IF k[j]
//				dst[i+7:i] := a[id+7:id]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm_maskz_permutexvar_epi8'.
// Requires AVX512VL.
func MaskzPermutexvarEpi8(k x86.Mmask16, idx x86.M128i, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// PermutexvarEpi8: Shuffle 8-bit integers in 'a' using the corresponding index
// in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			id := idx[i+3:i]*8
//			dst[i+7:i] := a[id+7:id]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm_permutexvar_epi8'.
// Requires AVX512VL.
func PermutexvarEpi8(idx x86.M128i, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPermutexvarEpi8: Shuffle 8-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			id := idx[i+4:i]*8
//			IF k[j]
//				dst[i+7:i] := a[id+7:id]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm256_mask_permutexvar_epi8'.
// Requires AVX512VL.
func M256MaskPermutexvarEpi8(src x86.M256i, k x86.Mmask32, idx x86.M256i, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPermutexvarEpi8: Shuffle 8-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			id := idx[i+4:i]*8
//			IF k[j]
//				dst[i+7:i] := a[id+7:id]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm256_maskz_permutexvar_epi8'.
// Requires AVX512VL.
func M256MaskzPermutexvarEpi8(k x86.Mmask32, idx x86.M256i, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256PermutexvarEpi8: Shuffle 8-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			id := idx[i+4:i]*8
//			dst[i+7:i] := a[id+7:id]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm256_permutexvar_epi8'.
// Requires AVX512VL.
func M256PermutexvarEpi8(idx x86.M256i, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}

