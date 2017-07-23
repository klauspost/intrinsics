// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package avx512bw

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// MaskAbsEpi16: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ABS(a[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm_mask_abs_epi16'.
// Requires AVX512BW.
func MaskAbsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAbsEpi16: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ABS(a[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm_maskz_abs_epi16'.
// Requires AVX512BW.
func MaskzAbsEpi16(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAbsEpi16: Compute the absolute value of packed 16-bit integers in
// 'a', and store the unsigned results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ABS(a[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm256_mask_abs_epi16'.
// Requires AVX512BW.
func M256MaskAbsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAbsEpi16: Compute the absolute value of packed 16-bit integers in
// 'a', and store the unsigned results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ABS(a[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm256_maskz_abs_epi16'.
// Requires AVX512BW.
func M256MaskzAbsEpi16(k x86.Mmask16, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AbsEpi16: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := ABS(a[i+15:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm512_abs_epi16'.
// Requires AVX512BW.
func M512AbsEpi16(a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAbsEpi16: Compute the absolute value of packed 16-bit integers in
// 'a', and store the unsigned results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ABS(a[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm512_mask_abs_epi16'.
// Requires AVX512BW.
func M512MaskAbsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAbsEpi16: Compute the absolute value of packed 16-bit integers in
// 'a', and store the unsigned results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ABS(a[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm512_maskz_abs_epi16'.
// Requires AVX512BW.
func M512MaskzAbsEpi16(k x86.Mmask32, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAbsEpi8: Compute the absolute value of packed 8-bit integers in 'a', and
// store the unsigned results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := ABS(a[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm_mask_abs_epi8'.
// Requires AVX512BW.
func MaskAbsEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAbsEpi8: Compute the absolute value of packed 8-bit integers in 'a',
// and store the unsigned results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := ABS(a[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm_maskz_abs_epi8'.
// Requires AVX512BW.
func MaskzAbsEpi8(k x86.Mmask16, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAbsEpi8: Compute the absolute value of packed 8-bit integers in 'a',
// and store the unsigned results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := ABS(a[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm256_mask_abs_epi8'.
// Requires AVX512BW.
func M256MaskAbsEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAbsEpi8: Compute the absolute value of packed 8-bit integers in
// 'a', and store the unsigned results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := ABS(a[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm256_maskz_abs_epi8'.
// Requires AVX512BW.
func M256MaskzAbsEpi8(k x86.Mmask32, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AbsEpi8: Compute the absolute value of packed 8-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := ABS(a[i+7:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm512_abs_epi8'.
// Requires AVX512BW.
func M512AbsEpi8(a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAbsEpi8: Compute the absolute value of packed 8-bit integers in 'a',
// and store the unsigned results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := ABS(a[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm512_mask_abs_epi8'.
// Requires AVX512BW.
func M512MaskAbsEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAbsEpi8: Compute the absolute value of packed 8-bit integers in
// 'a', and store the unsigned results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := ABS(a[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSB'. Intrinsic: '_mm512_maskz_abs_epi8'.
// Requires AVX512BW.
func M512MaskzAbsEpi8(k x86.Mmask64, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] + b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm_mask_add_epi16'.
// Requires AVX512BW.
func MaskAddEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] + b[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm_maskz_add_epi16'.
// Requires AVX512BW.
func MaskzAddEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] + b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm256_mask_add_epi16'.
// Requires AVX512BW.
func M256MaskAddEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] + b[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm256_maskz_add_epi16'.
// Requires AVX512BW.
func M256MaskzAddEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := a[i+15:i] + b[i+15:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm512_add_epi16'.
// Requires AVX512BW.
func M512AddEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] + b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm512_mask_add_epi16'.
// Requires AVX512BW.
func M512MaskAddEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] + b[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm512_maskz_add_epi16'.
// Requires AVX512BW.
func M512MaskzAddEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] + b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm_mask_add_epi8'.
// Requires AVX512BW.
func MaskAddEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] + b[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm_maskz_add_epi8'.
// Requires AVX512BW.
func MaskzAddEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] + b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm256_mask_add_epi8'.
// Requires AVX512BW.
func M256MaskAddEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] + b[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm256_maskz_add_epi8'.
// Requires AVX512BW.
func M256MaskzAddEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := a[i+7:i] + b[i+7:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm512_add_epi8'.
// Requires AVX512BW.
func M512AddEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] + b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm512_mask_add_epi8'.
// Requires AVX512BW.
func M512MaskAddEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] + b[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm512_maskz_add_epi8'.
// Requires AVX512BW.
func M512MaskzAddEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAddsEpi16: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm_mask_adds_epi16'.
// Requires AVX512BW.
func MaskAddsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAddsEpi16: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm_maskz_adds_epi16'.
// Requires AVX512BW.
func MaskzAddsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAddsEpi16: Add packed 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm256_mask_adds_epi16'.
// Requires AVX512BW.
func M256MaskAddsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAddsEpi16: Add packed 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm256_maskz_adds_epi16'.
// Requires AVX512BW.
func M256MaskzAddsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AddsEpi16: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm512_adds_epi16'.
// Requires AVX512BW.
func M512AddsEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddsEpi16: Add packed 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm512_mask_adds_epi16'.
// Requires AVX512BW.
func M512MaskAddsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAddsEpi16: Add packed 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm512_maskz_adds_epi16'.
// Requires AVX512BW.
func M512MaskzAddsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm_mask_adds_epi8'.
// Requires AVX512BW.
func MaskAddsEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm_maskz_adds_epi8'.
// Requires AVX512BW.
func MaskzAddsEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm256_mask_adds_epi8'.
// Requires AVX512BW.
func M256MaskAddsEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAddsEpi8: Add packed 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm256_maskz_adds_epi8'.
// Requires AVX512BW.
func M256MaskzAddsEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm512_adds_epi8'.
// Requires AVX512BW.
func M512AddsEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm512_mask_adds_epi8'.
// Requires AVX512BW.
func M512MaskAddsEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAddsEpi8: Add packed 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSB'. Intrinsic: '_mm512_maskz_adds_epi8'.
// Requires AVX512BW.
func M512MaskzAddsEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm_mask_adds_epu16'.
// Requires AVX512BW.
func MaskAddsEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm_maskz_adds_epu16'.
// Requires AVX512BW.
func MaskzAddsEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm256_mask_adds_epu16'.
// Requires AVX512BW.
func M256MaskAddsEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm256_maskz_adds_epu16'.
// Requires AVX512BW.
func M256MaskzAddsEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm512_adds_epu16'.
// Requires AVX512BW.
func M512AddsEpu16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm512_mask_adds_epu16'.
// Requires AVX512BW.
func M512MaskAddsEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16( a[i+15:i] + b[i+15:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDUSW'. Intrinsic: '_mm512_maskz_adds_epu16'.
// Requires AVX512BW.
func M512MaskzAddsEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm_mask_adds_epu8'.
// Requires AVX512BW.
func MaskAddsEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm_maskz_adds_epu8'.
// Requires AVX512BW.
func MaskzAddsEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm256_mask_adds_epu8'.
// Requires AVX512BW.
func M256MaskAddsEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm256_maskz_adds_epu8'.
// Requires AVX512BW.
func M256MaskzAddsEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm512_adds_epu8'.
// Requires AVX512BW.
func M512AddsEpu8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm512_mask_adds_epu8'.
// Requires AVX512BW.
func M512MaskAddsEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
// saturation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8( a[i+7:i] + b[i+7:i] )
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDUSB'. Intrinsic: '_mm512_maskz_adds_epu8'.
// Requires AVX512BW.
func M512MaskzAddsEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[255:0] := ((a[127:0] << 128) OR b[127:0]) >> (count[7:0]*8)
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm_mask_alignr_epi8'.
// Requires AVX512BW.
func MaskAlignrEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		tmp_dst[255:0] := ((a[127:0] << 128) OR b[127:0]) >> (count[7:0]*8)
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm_maskz_alignr_epi8'.
// Requires AVX512BW.
func MaskzAlignrEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i, count int) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into
// a 32-byte temporary result, shift the result right by 'count' bytes, and
// store the low 16 bytes in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*128
//			tmp[255:0] := ((a[i+127:i] << 128) OR b[i+127:i]) >> (count[7:0]*8)
//			tmp_dst[i+127:i] := tmp[127:0]
//		ENDFOR
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm256_mask_alignr_epi8'.
// Requires AVX512BW.
func M256MaskAlignrEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into
// a 32-byte temporary result, shift the result right by 'count' bytes, and
// store the low 16 bytes in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*128
//			tmp[255:0] := ((a[i+127:i] << 128) OR b[i+127:i]) >> (count[7:0]*8)
//			tmp_dst[i+127:i] := tmp[127:0]
//		ENDFOR
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm256_maskz_alignr_epi8'.
// Requires AVX512BW.
func M256MaskzAlignrEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i, count int) (dst x86.M256i) {
	panic("not implemented")
}


// M512AlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*128
//			tmp[255:0] := ((a[i+127:i] << 128) OR b[i+127:i]) >> (count[7:0]*8)
//			dst[i+127:i] := tmp[127:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm512_alignr_epi8'.
// Requires AVX512BW.
func M512AlignrEpi8(a x86.M512i, b x86.M512i, count int) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into
// a 32-byte temporary result, shift the result right by 'count' bytes, and
// store the low 16 bytes in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*128
//			tmp[255:0] := ((a[i+127:i] << 128) OR b[i+127:i]) >> (count[7:0]*8)
//			tmp_dst[i+127:i] := tmp[127:0]
//		ENDFOR
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm512_mask_alignr_epi8'.
// Requires AVX512BW.
func M512MaskAlignrEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i, count int) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into
// a 32-byte temporary result, shift the result right by 'count' bytes, and
// store the low 16 bytes in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*128
//			tmp[255:0] := ((a[i+127:i] << 128) OR b[i+127:i]) >> (count[7:0]*8)
//			tmp_dst[i+127:i] := tmp[127:0]
//		ENDFOR
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPALIGNR'. Intrinsic: '_mm512_maskz_alignr_epi8'.
// Requires AVX512BW.
func M512MaskzAlignrEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i, count int) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm_mask_avg_epu16'.
// Requires AVX512BW.
func MaskAvgEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm_maskz_avg_epu16'.
// Requires AVX512BW.
func MaskzAvgEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b',
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm256_mask_avg_epu16'.
// Requires AVX512BW.
func M256MaskAvgEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b',
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm256_maskz_avg_epu16'.
// Requires AVX512BW.
func M256MaskzAvgEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm512_avg_epu16'.
// Requires AVX512BW.
func M512AvgEpu16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b',
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm512_mask_avg_epu16'.
// Requires AVX512BW.
func M512MaskAvgEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b',
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm512_maskz_avg_epu16'.
// Requires AVX512BW.
func M512MaskzAvgEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskAvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm_mask_avg_epu8'.
// Requires AVX512BW.
func MaskAvgEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzAvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm_maskz_avg_epu8'.
// Requires AVX512BW.
func MaskzAvgEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskAvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm256_mask_avg_epu8'.
// Requires AVX512BW.
func M256MaskAvgEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzAvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm256_maskz_avg_epu8'.
// Requires AVX512BW.
func M256MaskzAvgEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512AvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm512_avg_epu8'.
// Requires AVX512BW.
func M512AvgEpu8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm512_mask_avg_epu8'.
// Requires AVX512BW.
func M512MaskAvgEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzAvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm512_maskz_avg_epu8'.
// Requires AVX512BW.
func M512MaskzAvgEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskBlendEpi16: Blend packed 16-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := b[i+15:i]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBLENDMW'. Intrinsic: '_mm_mask_blend_epi16'.
// Requires AVX512BW.
func MaskBlendEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskBlendEpi16: Blend packed 16-bit integers from 'a' and 'b' using
// control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := b[i+15:i]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBLENDMW'. Intrinsic: '_mm256_mask_blend_epi16'.
// Requires AVX512BW.
func M256MaskBlendEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskBlendEpi16: Blend packed 16-bit integers from 'a' and 'b' using
// control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := b[i+15:i]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBLENDMW'. Intrinsic: '_mm512_mask_blend_epi16'.
// Requires AVX512BW.
func M512MaskBlendEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskBlendEpi8: Blend packed 8-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := b[i+7:i]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBLENDMB'. Intrinsic: '_mm_mask_blend_epi8'.
// Requires AVX512BW.
func MaskBlendEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskBlendEpi8: Blend packed 8-bit integers from 'a' and 'b' using
// control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := b[i+7:i]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBLENDMB'. Intrinsic: '_mm256_mask_blend_epi8'.
// Requires AVX512BW.
func M256MaskBlendEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskBlendEpi8: Blend packed 8-bit integers from 'a' and 'b' using
// control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := b[i+7:i]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBLENDMB'. Intrinsic: '_mm512_mask_blend_epi8'.
// Requires AVX512BW.
func M512MaskBlendEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskBroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm_mask_broadcastb_epi8'.
// Requires AVX512BW.
func MaskBroadcastbEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzBroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm_maskz_broadcastb_epi8'.
// Requires AVX512BW.
func MaskzBroadcastbEpi8(k x86.Mmask16, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskBroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm256_mask_broadcastb_epi8'.
// Requires AVX512BW.
func M256MaskBroadcastbEpi8(src x86.M256i, k x86.Mmask32, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzBroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm256_maskz_broadcastb_epi8'.
// Requires AVX512BW.
func M256MaskzBroadcastbEpi8(k x86.Mmask32, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512BroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm512_broadcastb_epi8'.
// Requires AVX512BW.
func M512BroadcastbEpi8(a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskBroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm512_mask_broadcastb_epi8'.
// Requires AVX512BW.
func M512MaskBroadcastbEpi8(src x86.M512i, k x86.Mmask64, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzBroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm512_maskz_broadcastb_epi8'.
// Requires AVX512BW.
func M512MaskzBroadcastbEpi8(k x86.Mmask64, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskBroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm_mask_broadcastw_epi16'.
// Requires AVX512BW.
func MaskBroadcastwEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzBroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm_maskz_broadcastw_epi16'.
// Requires AVX512BW.
func MaskzBroadcastwEpi16(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskBroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm256_mask_broadcastw_epi16'.
// Requires AVX512BW.
func M256MaskBroadcastwEpi16(src x86.M256i, k x86.Mmask16, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzBroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a'
// to all elements of 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm256_maskz_broadcastw_epi16'.
// Requires AVX512BW.
func M256MaskzBroadcastwEpi16(k x86.Mmask16, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512BroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm512_broadcastw_epi16'.
// Requires AVX512BW.
func M512BroadcastwEpi16(a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskBroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm512_mask_broadcastw_epi16'.
// Requires AVX512BW.
func M512MaskBroadcastwEpi16(src x86.M512i, k x86.Mmask32, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzBroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a'
// to all elements of 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm512_maskz_broadcastw_epi16'.
// Requires AVX512BW.
func M512MaskzBroadcastwEpi16(k x86.Mmask32, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512BslliEpi128: Shift 128-bit lanes in 'a' left by 'imm8' bytes while
// shifting in zeros, and store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] << (tmp*8)
//		dst[255:128] := a[255:128] << (tmp*8)
//		dst[383:256] := a[383:256] << (tmp*8)
//		dst[511:384] := a[511:384] << (tmp*8)
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLDQ'. Intrinsic: '_mm512_bslli_epi128'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512BslliEpi128(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512BsrliEpi128: Shift 128-bit lanes in 'a' right by 'imm8' bytes while
// shifting in zeros, and store the results in 'dst'. 
//
//		tmp := imm8[7:0]
//		IF tmp > 15
//			tmp := 16
//		FI
//		dst[127:0] := a[127:0] >> (tmp*8)
//		dst[255:128] := a[255:128] >> (tmp*8)
//		dst[383:256] := a[383:256] >> (tmp*8)
//		dst[511:384] := a[511:384] >> (tmp*8)
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLDQ'. Intrinsic: '_mm512_bsrli_epi128'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512BsrliEpi128(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// CmpEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmp_epi16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func CmpEpi16Mask(a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k1' using zeromask 'k' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmp_epi16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskCmpEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmp_epi16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256CmpEpi16Mask(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' based on
// the comparison operand specified by 'imm8', and store the results in mask
// vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmp_epi16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskCmpEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmp_epi16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpEpi16Mask(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' based on
// the comparison operand specified by 'imm8', and store the results in mask
// vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmp_epi16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmp_epi8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func CmpEpi8Mask(a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmp_epi8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskCmpEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmp_epi8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256CmpEpi8Mask(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' based on
// the comparison operand specified by 'imm8', and store the results in mask
// vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmp_epi8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskCmpEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmp_epi8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpEpi8Mask(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' based on
// the comparison operand specified by 'imm8', and store the results in mask
// vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmp_epi8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' based
// on the comparison operand specified by 'imm8', and store the results in mask
// vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmp_epu16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func CmpEpu16Mask(a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmp_epu16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskCmpEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmp_epu16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256CmpEpu16Mask(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmp_epu16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskCmpEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmp_epu16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpEpu16Mask(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmp_epu16_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' based on
// the comparison operand specified by 'imm8', and store the results in mask
// vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmp_epu8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func CmpEpu8Mask(a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' based
// on the comparison operand specified by 'imm8', and store the results in mask
// vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmp_epu8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskCmpEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' based
// on the comparison operand specified by 'imm8', and store the results in mask
// vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmp_epu8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256CmpEpu8Mask(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmp_epu8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskCmpEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i, imm8 byte) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' based
// on the comparison operand specified by 'imm8', and store the results in mask
// vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmp_epu8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpEpu8Mask(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmp_epu8_mask'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpeqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmpeq_epi16_mask'.
// Requires AVX512BW.
func CmpeqEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpeqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmpeq_epi16_mask'.
// Requires AVX512BW.
func MaskCmpeqEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpeqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmpeq_epi16_mask'.
// Requires AVX512BW.
func M256CmpeqEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpeqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmpeq_epi16_mask'.
// Requires AVX512BW.
func M256MaskCmpeqEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpeqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmpeq_epi16_mask'.
// Requires AVX512BW.
func M512CmpeqEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpeqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmpeq_epi16_mask'.
// Requires AVX512BW.
func M512MaskCmpeqEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpeqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmpeq_epi8_mask'.
// Requires AVX512BW.
func CmpeqEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpeqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmpeq_epi8_mask'.
// Requires AVX512BW.
func MaskCmpeqEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpeqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmpeq_epi8_mask'.
// Requires AVX512BW.
func M256CmpeqEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpeqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmpeq_epi8_mask'.
// Requires AVX512BW.
func M256MaskCmpeqEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpeqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmpeq_epi8_mask'.
// Requires AVX512BW.
func M512CmpeqEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpeqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmpeq_epi8_mask'.
// Requires AVX512BW.
func M512MaskCmpeqEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpeqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmpeq_epu16_mask'.
// Requires AVX512BW.
func CmpeqEpu16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpeqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmpeq_epu16_mask'.
// Requires AVX512BW.
func MaskCmpeqEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpeqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpeq_epu16_mask'.
// Requires AVX512BW.
func M256CmpeqEpu16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpeqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for equality, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpeq_epu16_mask'.
// Requires AVX512BW.
func M256MaskCmpeqEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpeqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpeq_epu16_mask'.
// Requires AVX512BW.
func M512CmpeqEpu16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpeqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for equality, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpeq_epu16_mask'.
// Requires AVX512BW.
func M512MaskCmpeqEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpeqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmpeq_epu8_mask'.
// Requires AVX512BW.
func CmpeqEpu8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpeqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmpeq_epu8_mask'.
// Requires AVX512BW.
func MaskCmpeqEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpeqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmpeq_epu8_mask'.
// Requires AVX512BW.
func M256CmpeqEpu8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpeqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmpeq_epu8_mask'.
// Requires AVX512BW.
func M256MaskCmpeqEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpeqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmpeq_epu8_mask'.
// Requires AVX512BW.
func M512CmpeqEpu8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpeqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmpeq_epu8_mask'.
// Requires AVX512BW.
func M512MaskCmpeqEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpgeEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmpge_epi16_mask'.
// Requires AVX512BW.
func CmpgeEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpgeEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmpge_epi16_mask'.
// Requires AVX512BW.
func MaskCmpgeEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpgeEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmpge_epi16_mask'.
// Requires AVX512BW.
func M256CmpgeEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpgeEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmpge_epi16_mask'.
// Requires AVX512BW.
func M256MaskCmpgeEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgeEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmpge_epi16_mask'.
// Requires AVX512BW.
func M512CmpgeEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpgeEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmpge_epi16_mask'.
// Requires AVX512BW.
func M512MaskCmpgeEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpgeEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmpge_epi8_mask'.
// Requires AVX512BW.
func CmpgeEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpgeEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmpge_epi8_mask'.
// Requires AVX512BW.
func MaskCmpgeEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpgeEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmpge_epi8_mask'.
// Requires AVX512BW.
func M256CmpgeEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpgeEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmpge_epi8_mask'.
// Requires AVX512BW.
func M256MaskCmpgeEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpgeEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmpge_epi8_mask'.
// Requires AVX512BW.
func M512CmpgeEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpgeEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmpge_epi8_mask'.
// Requires AVX512BW.
func M512MaskCmpgeEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpgeEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmpge_epu16_mask'.
// Requires AVX512BW.
func CmpgeEpu16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpgeEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmpge_epu16_mask'.
// Requires AVX512BW.
func MaskCmpgeEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpgeEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpge_epu16_mask'.
// Requires AVX512BW.
func M256CmpgeEpu16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpgeEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for greater-than-or-equal, and store the results in mask vector 'k1'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpge_epu16_mask'.
// Requires AVX512BW.
func M256MaskCmpgeEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgeEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpge_epu16_mask'.
// Requires AVX512BW.
func M512CmpgeEpu16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpgeEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for greater-than-or-equal, and store the results in mask vector 'k1'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpge_epu16_mask'.
// Requires AVX512BW.
func M512MaskCmpgeEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpgeEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmpge_epu8_mask'.
// Requires AVX512BW.
func CmpgeEpu8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpgeEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmpge_epu8_mask'.
// Requires AVX512BW.
func MaskCmpgeEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpgeEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmpge_epu8_mask'.
// Requires AVX512BW.
func M256CmpgeEpu8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpgeEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmpge_epu8_mask'.
// Requires AVX512BW.
func M256MaskCmpgeEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpgeEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmpge_epu8_mask'.
// Requires AVX512BW.
func M512CmpgeEpu8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpgeEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] >= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmpge_epu8_mask'.
// Requires AVX512BW.
func M512MaskCmpgeEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpgtEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmpgt_epi16_mask'.
// Requires AVX512BW.
func CmpgtEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpgtEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >== b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmpgt_epi16_mask'.
// Requires AVX512BW.
func MaskCmpgtEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpgtEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmpgt_epi16_mask'.
// Requires AVX512BW.
func M256CmpgtEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpgtEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmpgt_epi16_mask'.
// Requires AVX512BW.
func M256MaskCmpgtEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgtEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmpgt_epi16_mask'.
// Requires AVX512BW.
func M512CmpgtEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpgtEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmpgt_epi16_mask'.
// Requires AVX512BW.
func M512MaskCmpgtEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpgtEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmpgt_epi8_mask'.
// Requires AVX512BW.
func CmpgtEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpgtEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmpgt_epi8_mask'.
// Requires AVX512BW.
func MaskCmpgtEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpgtEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmpgt_epi8_mask'.
// Requires AVX512BW.
func M256CmpgtEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpgtEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmpgt_epi8_mask'.
// Requires AVX512BW.
func M256MaskCmpgtEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpgtEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmpgt_epi8_mask'.
// Requires AVX512BW.
func M512CmpgtEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpgtEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmpgt_epi8_mask'.
// Requires AVX512BW.
func M512MaskCmpgtEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpgtEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmpgt_epu16_mask'.
// Requires AVX512BW.
func CmpgtEpu16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpgtEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] >== b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmpgt_epu16_mask'.
// Requires AVX512BW.
func MaskCmpgtEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpgtEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpgt_epu16_mask'.
// Requires AVX512BW.
func M256CmpgtEpu16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpgtEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for greater-than, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpgt_epu16_mask'.
// Requires AVX512BW.
func M256MaskCmpgtEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgtEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpgt_epu16_mask'.
// Requires AVX512BW.
func M512CmpgtEpu16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpgtEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for greater-than, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpgt_epu16_mask'.
// Requires AVX512BW.
func M512MaskCmpgtEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpgtEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmpgt_epu8_mask'.
// Requires AVX512BW.
func CmpgtEpu8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpgtEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmpgt_epu8_mask'.
// Requires AVX512BW.
func MaskCmpgtEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpgtEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmpgt_epu8_mask'.
// Requires AVX512BW.
func M256CmpgtEpu8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpgtEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmpgt_epu8_mask'.
// Requires AVX512BW.
func M256MaskCmpgtEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpgtEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmpgt_epu8_mask'.
// Requires AVX512BW.
func M512CmpgtEpu8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpgtEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] > b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmpgt_epu8_mask'.
// Requires AVX512BW.
func M512MaskCmpgtEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpleEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmple_epi16_mask'.
// Requires AVX512BW.
func CmpleEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpleEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmple_epi16_mask'.
// Requires AVX512BW.
func MaskCmpleEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpleEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmple_epi16_mask'.
// Requires AVX512BW.
func M256CmpleEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpleEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmple_epi16_mask'.
// Requires AVX512BW.
func M256MaskCmpleEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpleEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmple_epi16_mask'.
// Requires AVX512BW.
func M512CmpleEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpleEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmple_epi16_mask'.
// Requires AVX512BW.
func M512MaskCmpleEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpleEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmple_epi8_mask'.
// Requires AVX512BW.
func CmpleEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpleEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmple_epi8_mask'.
// Requires AVX512BW.
func MaskCmpleEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpleEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmple_epi8_mask'.
// Requires AVX512BW.
func M256CmpleEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpleEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmple_epi8_mask'.
// Requires AVX512BW.
func M256MaskCmpleEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpleEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmple_epi8_mask'.
// Requires AVX512BW.
func M512CmpleEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpleEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmple_epi8_mask'.
// Requires AVX512BW.
func M512MaskCmpleEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpleEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmple_epu16_mask'.
// Requires AVX512BW.
func CmpleEpu16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpleEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmple_epu16_mask'.
// Requires AVX512BW.
func MaskCmpleEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpleEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmple_epu16_mask'.
// Requires AVX512BW.
func M256CmpleEpu16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpleEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for less-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//			ELSE 
//					k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmple_epu16_mask'.
// Requires AVX512BW.
func M256MaskCmpleEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpleEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmple_epu16_mask'.
// Requires AVX512BW.
func M512CmpleEpu16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpleEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for less-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmple_epu16_mask'.
// Requires AVX512BW.
func M512MaskCmpleEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpleEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmple_epu8_mask'.
// Requires AVX512BW.
func CmpleEpu8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpleEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmple_epu8_mask'.
// Requires AVX512BW.
func MaskCmpleEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpleEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmple_epu8_mask'.
// Requires AVX512BW.
func M256CmpleEpu8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpleEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmple_epu8_mask'.
// Requires AVX512BW.
func M256MaskCmpleEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpleEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmple_epu8_mask'.
// Requires AVX512BW.
func M512CmpleEpu8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpleEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k' using
// zeromask 'k1' (elements are zeroed out when the corresponding mask bit is
// not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] <= b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmple_epu8_mask'.
// Requires AVX512BW.
func M512MaskCmpleEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpltEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for less-than,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmplt_epi16_mask'.
// Requires AVX512BW.
func CmpltEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpltEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmplt_epi16_mask'.
// Requires AVX512BW.
func MaskCmpltEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpltEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmplt_epi16_mask'.
// Requires AVX512BW.
func M256CmpltEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpltEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmplt_epi16_mask'.
// Requires AVX512BW.
func M256MaskCmpltEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpltEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmplt_epi16_mask'.
// Requires AVX512BW.
func M512CmpltEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpltEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmplt_epi16_mask'.
// Requires AVX512BW.
func M512MaskCmpltEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpltEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for less-than,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmplt_epi8_mask'.
// Requires AVX512BW.
func CmpltEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpltEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmplt_epi8_mask'.
// Requires AVX512BW.
func MaskCmpltEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpltEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmplt_epi8_mask'.
// Requires AVX512BW.
func M256CmpltEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpltEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmplt_epi8_mask'.
// Requires AVX512BW.
func M256MaskCmpltEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpltEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmplt_epi8_mask'.
// Requires AVX512BW.
func M512CmpltEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpltEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmplt_epi8_mask'.
// Requires AVX512BW.
func M512MaskCmpltEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpltEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmplt_epu16_mask'.
// Requires AVX512BW.
func CmpltEpu16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpltEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmplt_epu16_mask'.
// Requires AVX512BW.
func MaskCmpltEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpltEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmplt_epu16_mask'.
// Requires AVX512BW.
func M256CmpltEpu16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpltEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for less-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmplt_epu16_mask'.
// Requires AVX512BW.
func M256MaskCmpltEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpltEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmplt_epu16_mask'.
// Requires AVX512BW.
func M512CmpltEpu16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpltEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for less-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmplt_epu16_mask'.
// Requires AVX512BW.
func M512MaskCmpltEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpltEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmplt_epu8_mask'.
// Requires AVX512BW.
func CmpltEpu8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpltEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmplt_epu8_mask'.
// Requires AVX512BW.
func MaskCmpltEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpltEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmplt_epu8_mask'.
// Requires AVX512BW.
func M256CmpltEpu8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpltEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmplt_epu8_mask'.
// Requires AVX512BW.
func M256MaskCmpltEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpltEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmplt_epu8_mask'.
// Requires AVX512BW.
func M512CmpltEpu8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpltEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmplt_epu8_mask'.
// Requires AVX512BW.
func M512MaskCmpltEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpneqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_cmpneq_epi16_mask'.
// Requires AVX512BW.
func CmpneqEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpneqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm_mask_cmpneq_epi16_mask'.
// Requires AVX512BW.
func MaskCmpneqEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpneqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmpneq_epi16_mask'.
// Requires AVX512BW.
func M256CmpneqEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpneqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_mask_cmpneq_epi16_mask'.
// Requires AVX512BW.
func M256MaskCmpneqEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpneqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmpneq_epi16_mask'.
// Requires AVX512BW.
func M512CmpneqEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpneqEpi16Mask: Compare packed 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_mask_cmpneq_epi16_mask'.
// Requires AVX512BW.
func M512MaskCmpneqEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpneqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for not-equal,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_cmpneq_epi8_mask'.
// Requires AVX512BW.
func CmpneqEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpneqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm_mask_cmpneq_epi8_mask'.
// Requires AVX512BW.
func MaskCmpneqEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpneqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmpneq_epi8_mask'.
// Requires AVX512BW.
func M256CmpneqEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpneqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_mask_cmpneq_epi8_mask'.
// Requires AVX512BW.
func M256MaskCmpneqEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpneqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmpneq_epi8_mask'.
// Requires AVX512BW.
func M512CmpneqEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpneqEpi8Mask: Compare packed 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_mask_cmpneq_epi8_mask'.
// Requires AVX512BW.
func M512MaskCmpneqEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// CmpneqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_cmpneq_epu16_mask'.
// Requires AVX512BW.
func CmpneqEpu16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// MaskCmpneqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm_mask_cmpneq_epu16_mask'.
// Requires AVX512BW.
func MaskCmpneqEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256CmpneqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpneq_epu16_mask'.
// Requires AVX512BW.
func M256CmpneqEpu16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskCmpneqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for not-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpneq_epu16_mask'.
// Requires AVX512BW.
func M256MaskCmpneqEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpneqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpneq_epu16_mask'.
// Requires AVX512BW.
func M512CmpneqEpu16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskCmpneqEpu16Mask: Compare packed unsigned 16-bit integers in 'a' and
// 'b' for not-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpneq_epu16_mask'.
// Requires AVX512BW.
func M512MaskCmpneqEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// CmpneqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_cmpneq_epu8_mask'.
// Requires AVX512BW.
func CmpneqEpu8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// MaskCmpneqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k' using zeromask 'k1'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm_mask_cmpneq_epu8_mask'.
// Requires AVX512BW.
func MaskCmpneqEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256CmpneqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmpneq_epu8_mask'.
// Requires AVX512BW.
func M256CmpneqEpu8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256MaskCmpneqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and
// 'b' for not-equal, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmpneq_epu8_mask'.
// Requires AVX512BW.
func M256MaskCmpneqEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512CmpneqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmpneq_epu8_mask'.
// Requires AVX512BW.
func M512CmpneqEpu8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512MaskCmpneqEpu8Mask: Compare packed unsigned 8-bit integers in 'a' and
// 'b' for not-equal, and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmpneq_epu8_mask'.
// Requires AVX512BW.
func M512MaskCmpneqEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// Cvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit integers
// with truncation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm_cvtepi16_epi8'.
// Requires AVX512BW.
func Cvtepi16Epi8(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskCvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm_mask_cvtepi16_epi8'.
// Requires AVX512BW.
func MaskCvtepi16Epi8(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzCvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm_maskz_cvtepi16_epi8'.
// Requires AVX512BW.
func MaskzCvtepi16Epi8(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256Cvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm256_cvtepi16_epi8'.
// Requires AVX512BW.
func M256Cvtepi16Epi8(a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskCvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm256_mask_cvtepi16_epi8'.
// Requires AVX512BW.
func M256MaskCvtepi16Epi8(src x86.M128i, k x86.Mmask16, a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskzCvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm256_maskz_cvtepi16_epi8'.
// Requires AVX512BW.
func M256MaskzCvtepi16Epi8(k x86.Mmask16, a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M512Cvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm512_cvtepi16_epi8'.
// Requires AVX512BW.
func M512Cvtepi16Epi8(a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskCvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm512_mask_cvtepi16_epi8'.
// Requires AVX512BW.
func M512MaskCvtepi16Epi8(src x86.M256i, k x86.Mmask32, a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskzCvtepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with truncation, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Truncate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm512_maskz_cvtepi16_epi8'.
// Requires AVX512BW.
func M512MaskzCvtepi16Epi8(k x86.Mmask32, a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// Skipped: _mm_mask_cvtepi16_storeu_epi8. Contains pointer parameter.


// Skipped: _mm256_mask_cvtepi16_storeu_epi8. Contains pointer parameter.


// Skipped: _mm512_mask_cvtepi16_storeu_epi8. Contains pointer parameter.


// MaskCvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := SignExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := src[l+15:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm_mask_cvtepi8_epi16'.
// Requires AVX512BW.
func MaskCvtepi8Epi16(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzCvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := SignExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm_maskz_cvtepi8_epi16'.
// Requires AVX512BW.
func MaskzCvtepi8Epi16(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskCvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed
// 16-bit integers, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := SignExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := src[l+15:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm256_mask_cvtepi8_epi16'.
// Requires AVX512BW.
func M256MaskCvtepi8Epi16(src x86.M256i, k x86.Mmask16, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzCvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed
// 16-bit integers, and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := SignExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm256_maskz_cvtepi8_epi16'.
// Requires AVX512BW.
func M256MaskzCvtepi8Epi16(k x86.Mmask16, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512Cvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			l := j*16
//			dst[l+15:l] := SignExtend(a[i+7:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm512_cvtepi8_epi16'.
// Requires AVX512BW.
func M512Cvtepi8Epi16(a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskCvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed
// 16-bit integers, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := SignExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := src[l+15:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm512_mask_cvtepi8_epi16'.
// Requires AVX512BW.
func M512MaskCvtepi8Epi16(src x86.M512i, k x86.Mmask32, a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzCvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed
// 16-bit integers, and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := SignExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVSXBW'. Intrinsic: '_mm512_maskz_cvtepi8_epi16'.
// Requires AVX512BW.
func M512MaskzCvtepi8Epi16(k x86.Mmask32, a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskCvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := ZeroExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := src[l+15:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm_mask_cvtepu8_epi16'.
// Requires AVX512BW.
func MaskCvtepu8Epi16(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzCvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := ZeroExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm_maskz_cvtepu8_epi16'.
// Requires AVX512BW.
func MaskzCvtepu8Epi16(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskCvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := ZeroExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := src[l+15:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm256_mask_cvtepu8_epi16'.
// Requires AVX512BW.
func M256MaskCvtepu8Epi16(src x86.M256i, k x86.Mmask16, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzCvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := ZeroExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm256_maskz_cvtepu8_epi16'.
// Requires AVX512BW.
func M256MaskzCvtepu8Epi16(k x86.Mmask16, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512Cvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			l := j*16
//			dst[l+15:l] := ZeroExtend(a[i+7:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm512_cvtepu8_epi16'.
// Requires AVX512BW.
func M512Cvtepu8Epi16(a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskCvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := ZeroExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := src[l+15:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm512_mask_cvtepu8_epi16'.
// Requires AVX512BW.
func M512MaskCvtepu8Epi16(src x86.M512i, k x86.Mmask32, a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzCvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to
// packed 16-bit integers, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			l := j*16
//			IF k[j]
//				dst[l+15:l] := ZeroExtend(a[i+7:i])
//			ELSE
//				dst[l+15:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVZXBW'. Intrinsic: '_mm512_maskz_cvtepu8_epi16'.
// Requires AVX512BW.
func M512MaskzCvtepu8Epi16(k x86.Mmask32, a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


// Cvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm_cvtsepi16_epi8'.
// Requires AVX512BW.
func Cvtsepi16Epi8(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskCvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm_mask_cvtsepi16_epi8'.
// Requires AVX512BW.
func MaskCvtsepi16Epi8(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzCvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm_maskz_cvtsepi16_epi8'.
// Requires AVX512BW.
func MaskzCvtsepi16Epi8(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256Cvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm256_cvtsepi16_epi8'.
// Requires AVX512BW.
func M256Cvtsepi16Epi8(a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskCvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm256_mask_cvtsepi16_epi8'.
// Requires AVX512BW.
func M256MaskCvtsepi16Epi8(src x86.M128i, k x86.Mmask16, a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskzCvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with signed saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm256_maskz_cvtsepi16_epi8'.
// Requires AVX512BW.
func M256MaskzCvtsepi16Epi8(k x86.Mmask16, a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M512Cvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm512_cvtsepi16_epi8'.
// Requires AVX512BW.
func M512Cvtsepi16Epi8(a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskCvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm512_mask_cvtsepi16_epi8'.
// Requires AVX512BW.
func M512MaskCvtsepi16Epi8(src x86.M256i, k x86.Mmask32, a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskzCvtsepi16Epi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with signed saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_Int16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm512_maskz_cvtsepi16_epi8'.
// Requires AVX512BW.
func M512MaskzCvtsepi16Epi8(k x86.Mmask32, a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// Skipped: _mm_mask_cvtsepi16_storeu_epi8. Contains pointer parameter.


// Skipped: _mm256_mask_cvtsepi16_storeu_epi8. Contains pointer parameter.


// Skipped: _mm512_mask_cvtsepi16_storeu_epi8. Contains pointer parameter.


// Cvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to packed
// unsigned 8-bit integers with unsigned saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm_cvtusepi16_epi8'.
// Requires AVX512BW.
func Cvtusepi16Epi8(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskCvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to packed
// unsigned 8-bit integers with unsigned saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm_mask_cvtusepi16_epi8'.
// Requires AVX512BW.
func MaskCvtusepi16Epi8(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzCvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm_maskz_cvtusepi16_epi8'.
// Requires AVX512BW.
func MaskzCvtusepi16Epi8(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256Cvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to packed
// unsigned 8-bit integers with unsigned saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm256_cvtusepi16_epi8'.
// Requires AVX512BW.
func M256Cvtusepi16Epi8(a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskCvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm256_mask_cvtusepi16_epi8'.
// Requires AVX512BW.
func M256MaskCvtusepi16Epi8(src x86.M128i, k x86.Mmask16, a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskzCvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm256_maskz_cvtusepi16_epi8'.
// Requires AVX512BW.
func M256MaskzCvtusepi16Epi8(k x86.Mmask16, a x86.M256i) (dst x86.M128i) {
	panic("not implemented")
}


// M512Cvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to packed
// unsigned 8-bit integers with unsigned saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm512_cvtusepi16_epi8'.
// Requires AVX512BW.
func M512Cvtusepi16Epi8(a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskCvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := src[l+7:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm512_mask_cvtusepi16_epi8'.
// Requires AVX512BW.
func M512MaskCvtusepi16Epi8(src x86.M256i, k x86.Mmask32, a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskzCvtusepi16Epi8: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				dst[l+7:l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			ELSE
//				dst[l+7:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm512_maskz_cvtusepi16_epi8'.
// Requires AVX512BW.
func M512MaskzCvtusepi16Epi8(k x86.Mmask32, a x86.M512i) (dst x86.M256i) {
	panic("not implemented")
}


// Skipped: _mm_mask_cvtusepi16_storeu_epi8. Contains pointer parameter.


// Skipped: _mm256_mask_cvtusepi16_storeu_epi8. Contains pointer parameter.


// Skipped: _mm512_mask_cvtusepi16_storeu_epi8. Contains pointer parameter.


// DbsadEpu8: Compute the sum of absolute differences (SADs) of quadruplets of
// unsigned 8-bit integers in 'a' compared to those in 'b', and store the
// 16-bit results in 'dst'.
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected according to the control in 'imm8', and
// each SAD in each 64-bit lane uses the selected quadruplet at 8-bit offsets. 
//
//		tmp[31:0] := select(b[127:0], imm8[1:0])
//		tmp[63:32] := select(b[127:0], imm8[3:2])
//		tmp[95:64] := select(b[127:0], imm8[5:4])
//		tmp[127:96] := select(b[127:0], imm8[7:6])
//		
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		dst[MAX:128] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func DbsadEpu8(a x86.M128i, b x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskDbsadEpu8: Compute the sum of absolute differences (SADs) of quadruplets
// of unsigned 8-bit integers in 'a' compared to those in 'b', and store the
// 16-bit results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set).
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected according to the control in 'imm8', and
// each SAD in each 64-bit lane uses the selected quadruplet at 8-bit offsets. 
//
//		tmp[31:0] := select(b[127:0], imm8[1:0])
//		tmp[63:32] := select(b[127:0], imm8[3:2])
//		tmp[95:64] := select(b[127:0], imm8[5:4])
//		tmp[127:96] := select(b[127:0], imm8[7:6])
//		
//		FOR j := 0 to 1
//			i := j*64
//			tmp_dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			tmp_dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			tmp_dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			tmp_dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm_mask_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskDbsadEpu8(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzDbsadEpu8: Compute the sum of absolute differences (SADs) of
// quadruplets of unsigned 8-bit integers in 'a' compared to those in 'b', and
// store the 16-bit results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set).
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected according to the control in 'imm8', and
// each SAD in each 64-bit lane uses the selected quadruplet at 8-bit offsets. 
//
//		tmp[31:0] := select(b[127:0], imm8[1:0])
//		tmp[63:32] := select(b[127:0], imm8[3:2])
//		tmp[95:64] := select(b[127:0], imm8[5:4])
//		tmp[127:96] := select(b[127:0], imm8[7:6])
//		
//		FOR j := 0 to 1
//			i := j*64
//			tmp_dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			tmp_dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			tmp_dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			tmp_dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm_maskz_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskzDbsadEpu8(k x86.Mmask8, a x86.M128i, b x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256DbsadEpu8: Compute the sum of absolute differences (SADs) of quadruplets
// of unsigned 8-bit integers in 'a' compared to those in 'b', and store the
// 16-bit results in 'dst'.
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected from within 128-bit lanes according to the
// control in 'imm8', and each SAD in each 64-bit lane uses the selected
// quadruplet at 8-bit offsets. 
//
//		FOR j := 0 to 1
//			i := j*128
//			tmp[i+31:i] := select(b[i+127:i], imm8[1:0])
//			tmp[i+63:i+32] := select(b[i+127:i], imm8[3:2])
//			tmp[i+95:i+64] := select(b[i+127:i], imm8[5:4])
//			tmp[i+127:i+96] := select(b[i+127:i], imm8[7:6])
//		ENDFOR
//		
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm256_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256DbsadEpu8(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskDbsadEpu8: Compute the sum of absolute differences (SADs) of
// quadruplets of unsigned 8-bit integers in 'a' compared to those in 'b', and
// store the 16-bit results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set).
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected from within 128-bit lanes according to the
// control in 'imm8', and each SAD in each 64-bit lane uses the selected
// quadruplet at 8-bit offsets. 
//
//		FOR j := 0 to 1
//			i := j*128
//			tmp[i+31:i] := select(b[i+127:i], imm8[1:0])
//			tmp[i+63:i+32] := select(b[i+127:i], imm8[3:2])
//			tmp[i+95:i+64] := select(b[i+127:i], imm8[5:4])
//			tmp[i+127:i+96] := select(b[i+127:i], imm8[7:6])
//		ENDFOR
//		
//		FOR j := 0 to 3
//			i := j*64
//			tmp_dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			tmp_dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			tmp_dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			tmp_dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm256_mask_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskDbsadEpu8(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzDbsadEpu8: Compute the sum of absolute differences (SADs) of
// quadruplets of unsigned 8-bit integers in 'a' compared to those in 'b', and
// store the 16-bit results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set).
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected from within 128-bit lanes according to the
// control in 'imm8', and each SAD in each 64-bit lane uses the selected
// quadruplet at 8-bit offsets. 
//
//		FOR j := 0 to 1
//			i := j*128
//			tmp[i+31:i] := select(b[i+127:i], imm8[1:0])
//			tmp[i+63:i+32] := select(b[i+127:i], imm8[3:2])
//			tmp[i+95:i+64] := select(b[i+127:i], imm8[5:4])
//			tmp[i+127:i+96] := select(b[i+127:i], imm8[7:6])
//		ENDFOR
//		
//		FOR j := 0 to 3
//			i := j*64
//			tmp_dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			tmp_dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			tmp_dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			tmp_dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm256_maskz_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzDbsadEpu8(k x86.Mmask16, a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512DbsadEpu8: Compute the sum of absolute differences (SADs) of quadruplets
// of unsigned 8-bit integers in 'a' compared to those in 'b', and store the
// 16-bit results in 'dst'.
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected from within 128-bit lanes according to the
// control in 'imm8', and each SAD in each 64-bit lane uses the selected
// quadruplet at 8-bit offsets. 
//
//		FOR j := 0 to 3
//			i := j*128
//			tmp[i+31:i] := select(b[i+127:i], imm8[1:0])
//			tmp[i+63:i+32] := select(b[i+127:i], imm8[3:2])
//			tmp[i+95:i+64] := select(b[i+127:i], imm8[5:4])
//			tmp[i+127:i+96] := select(b[i+127:i], imm8[7:6])
//		ENDFOR
//		
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm512_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512DbsadEpu8(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskDbsadEpu8: Compute the sum of absolute differences (SADs) of
// quadruplets of unsigned 8-bit integers in 'a' compared to those in 'b', and
// store the 16-bit results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set).
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected from within 128-bit lanes according to the
// control in 'imm8', and each SAD in each 64-bit lane uses the selected
// quadruplet at 8-bit offsets. 
//
//		FOR j := 0 to 3
//			i := j*128
//			tmp[i+31:i] := select(b[i+127:i], imm8[1:0])
//			tmp[i+63:i+32] := select(b[i+127:i], imm8[3:2])
//			tmp[i+95:i+64] := select(b[i+127:i], imm8[5:4])
//			tmp[i+127:i+96] := select(b[i+127:i], imm8[7:6])
//		ENDFOR
//		
//		FOR j := 0 to 7
//			i := j*64
//			tmp_dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			tmp_dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			tmp_dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			tmp_dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm512_mask_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskDbsadEpu8(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzDbsadEpu8: Compute the sum of absolute differences (SADs) of
// quadruplets of unsigned 8-bit integers in 'a' compared to those in 'b', and
// store the 16-bit results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set).
// 	Four SADs are performed on four 8-bit quadruplets for each 64-bit lane. The
// first two SADs use the lower 8-bit quadruplet of the lane from 'a', and the
// last two SADs use the uppper 8-bit quadruplet of the lane from 'a'.
// Quadruplets from 'b' are selected from within 128-bit lanes according to the
// control in 'imm8', and each SAD in each 64-bit lane uses the selected
// quadruplet at 8-bit offsets. 
//
//		FOR j := 0 to 3
//			i := j*128
//			tmp[i+31:i] := select(b[i+127:i], imm8[1:0])
//			tmp[i+63:i+32] := select(b[i+127:i], imm8[3:2])
//			tmp[i+95:i+64] := select(b[i+127:i], imm8[5:4])
//			tmp[i+127:i+96] := select(b[i+127:i], imm8[7:6])
//		ENDFOR
//		
//		FOR j := 0 to 7
//			i := j*64
//			tmp_dst[i+15:i] := ABS(a[i+7:i] - tmp[i+7:i]) + ABS(a[i+15:i+8] - tmp[i+15:i+8])
//						 + ABS(a[i+23:i+16] - tmp[i+23:i+16]) + ABS(a[i+31:i+24] - tmp[i+31:i+24])
//			
//			tmp_dst[i+31:i+16] := ABS(a[i+7:i] - tmp[i+15:i+8]) + ABS(a[i+15:i+8] - tmp[i+23:i+16])
//						 + ABS(a[i+23:i+16] - tmp[i+31:i+24]) + ABS(a[i+31:i+24] - tmp[i+39:i+32])
//			
//			tmp_dst[i+47:i+32] := ABS(a[i+39:i+32] - tmp[i+23:i+16]) + ABS(a[i+47:i+40] - tmp[i+31:i+24])
//						 + ABS(a[i+55:i+48] - tmp[i+39:i+32]) + ABS(a[i+63:i+56] - tmp[i+47:i+40])
//			
//			tmp_dst[i+63:i+48] := ABS(a[i+39:i+32] - tmp[i+31:i+24]) + ABS(a[i+47:i+40] - tmp[i+39:i+32])
//						 + ABS(a[i+55:i+48] - tmp[i+47:i+40]) + ABS(a[i+63:i+56] - tmp[i+55:i+48])
//		ENDFOR
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VDBPSADBW'. Intrinsic: '_mm512_maskz_dbsad_epu8'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzDbsadEpu8(k x86.Mmask32, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512Kunpackd: Unpack and interleave 32 bits from masks 'a' and 'b', and
// store the 64-bit result in 'k'. 
//
//		k[31:0] := a[31:0]
//		k[63:32] := b[31:0]
//		k[MAX:64] := 0
//
// Instruction: 'KUNPCKDQ'. Intrinsic: '_mm512_kunpackd'.
// Requires AVX512BW.
func M512Kunpackd(a x86.Mmask64, b x86.Mmask64) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512Kunpackw: Unpack and interleave 16 bits from masks 'a' and 'b', and
// store the 32-bit result in 'k'. 
//
//		k[15:0] := a[15:0]
//		k[31:16] := b[15:0]
//		k[MAX:32] := 0
//
// Instruction: 'KUNPCKWD'. Intrinsic: '_mm512_kunpackw'.
// Requires AVX512BW.
func M512Kunpackw(a x86.Mmask32, b x86.Mmask32) (dst x86.Mmask32) {
	panic("not implemented")
}


// Skipped: _mm_mask_loadu_epi16. Contains pointer parameter.


// Skipped: _mm_maskz_loadu_epi16. Contains pointer parameter.


// Skipped: _mm256_mask_loadu_epi16. Contains pointer parameter.


// Skipped: _mm256_maskz_loadu_epi16. Contains pointer parameter.


// Skipped: _mm512_mask_loadu_epi16. Contains pointer parameter.


// Skipped: _mm512_maskz_loadu_epi16. Contains pointer parameter.


// Skipped: _mm_mask_loadu_epi8. Contains pointer parameter.


// Skipped: _mm_maskz_loadu_epi8. Contains pointer parameter.


// Skipped: _mm256_mask_loadu_epi8. Contains pointer parameter.


// Skipped: _mm256_maskz_loadu_epi8. Contains pointer parameter.


// Skipped: _mm512_mask_loadu_epi8. Contains pointer parameter.


// Skipped: _mm512_maskz_loadu_epi8. Contains pointer parameter.


// MaskMaddEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm_mask_madd_epi16'.
// Requires AVX512BW.
func MaskMaddEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMaddEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm_maskz_madd_epi16'.
// Requires AVX512BW.
func MaskzMaddEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMaddEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm256_mask_madd_epi16'.
// Requires AVX512BW.
func M256MaskMaddEpi16(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMaddEpi16: Multiply packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm256_maskz_madd_epi16'.
// Requires AVX512BW.
func M256MaskzMaddEpi16(k x86.Mmask8, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaddEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			st[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm512_madd_epi16'.
// Requires AVX512BW.
func M512MaddEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMaddEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm512_mask_madd_epi16'.
// Requires AVX512BW.
func M512MaskMaddEpi16(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMaddEpi16: Multiply packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers. Horizontally add adjacent pairs of
// intermediate 32-bit integers, and pack the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i+16]*b[i+31:i+16] + a[i+15:i]*b[i+15:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADDWD'. Intrinsic: '_mm512_maskz_madd_epi16'.
// Requires AVX512BW.
func M512MaskzMaddEpi16(k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMaddubsEpi16: Multiply packed unsigned 8-bit integers in 'a' by packed
// signed 8-bit integers in 'b', producing intermediate signed 16-bit integers.
// Horizontally add adjacent pairs of intermediate signed 16-bit integers, and
// pack the saturated results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm_mask_maddubs_epi16'.
// Requires AVX512BW.
func MaskMaddubsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMaddubsEpi16: Multiply packed unsigned 8-bit integers in 'a' by packed
// signed 8-bit integers in 'b', producing intermediate signed 16-bit integers.
// Horizontally add adjacent pairs of intermediate signed 16-bit integers, and
// pack the saturated results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm_maskz_maddubs_epi16'.
// Requires AVX512BW.
func MaskzMaddubsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMaddubsEpi16: Multiply packed unsigned 8-bit integers in 'a' by
// packed signed 8-bit integers in 'b', producing intermediate signed 16-bit
// integers. Horizontally add adjacent pairs of intermediate signed 16-bit
// integers, and pack the saturated results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm256_mask_maddubs_epi16'.
// Requires AVX512BW.
func M256MaskMaddubsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMaddubsEpi16: Multiply packed unsigned 8-bit integers in 'a' by
// packed signed 8-bit integers in 'b', producing intermediate signed 16-bit
// integers. Horizontally add adjacent pairs of intermediate signed 16-bit
// integers, and pack the saturated results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm256_maskz_maddubs_epi16'.
// Requires AVX512BW.
func M256MaskzMaddubsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaddubsEpi16: Vertically multiply each unsigned 8-bit integer from 'a'
// with the corresponding signed 8-bit integer from 'b', producing intermediate
// signed 16-bit integers. Horizontally add adjacent pairs of intermediate
// signed 16-bit integers, and pack the saturated results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm512_maddubs_epi16'.
// Requires AVX512BW.
func M512MaddubsEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMaddubsEpi16: Multiply packed unsigned 8-bit integers in 'a' by
// packed signed 8-bit integers in 'b', producing intermediate signed 16-bit
// integers. Horizontally add adjacent pairs of intermediate signed 16-bit
// integers, and pack the saturated results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm512_mask_maddubs_epi16'.
// Requires AVX512BW.
func M512MaskMaddubsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMaddubsEpi16: Multiply packed unsigned 8-bit integers in 'a' by
// packed signed 8-bit integers in 'b', producing intermediate signed 16-bit
// integers. Horizontally add adjacent pairs of intermediate signed 16-bit
// integers, and pack the saturated results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16( a[i+15:i+8]*b[i+15:i+8] + a[i+7:i]*b[i+7:i] )
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADDUBSW'. Intrinsic: '_mm512_maskz_maddubs_epi16'.
// Requires AVX512BW.
func M512MaskzMaddubsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm_mask_max_epi16'.
// Requires AVX512BW.
func MaskMaxEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm_maskz_max_epi16'.
// Requires AVX512BW.
func MaskzMaxEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm256_mask_max_epi16'.
// Requires AVX512BW.
func M256MaskMaxEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm256_maskz_max_epi16'.
// Requires AVX512BW.
func M256MaskzMaxEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm512_mask_max_epi16'.
// Requires AVX512BW.
func M512MaskMaxEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm512_maskz_max_epi16'.
// Requires AVX512BW.
func M512MaskzMaxEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSW'. Intrinsic: '_mm512_max_epi16'.
// Requires AVX512BW.
func M512MaxEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm_mask_max_epi8'.
// Requires AVX512BW.
func MaskMaxEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm_maskz_max_epi8'.
// Requires AVX512BW.
func MaskzMaxEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm256_mask_max_epi8'.
// Requires AVX512BW.
func M256MaskMaxEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm256_maskz_max_epi8'.
// Requires AVX512BW.
func M256MaskzMaxEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm512_mask_max_epi8'.
// Requires AVX512BW.
func M512MaskMaxEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm512_maskz_max_epi8'.
// Requires AVX512BW.
func M512MaskzMaxEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm512_max_epi8'.
// Requires AVX512BW.
func M512MaxEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm_mask_max_epu16'.
// Requires AVX512BW.
func MaskMaxEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm_maskz_max_epu16'.
// Requires AVX512BW.
func MaskzMaxEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm256_mask_max_epu16'.
// Requires AVX512BW.
func M256MaskMaxEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed maximum values in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm256_maskz_max_epu16'.
// Requires AVX512BW.
func M256MaskzMaxEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm512_mask_max_epu16'.
// Requires AVX512BW.
func M512MaskMaxEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed maximum values in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] > b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm512_maskz_max_epu16'.
// Requires AVX512BW.
func M512MaskzMaxEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUW'. Intrinsic: '_mm512_max_epu16'.
// Requires AVX512BW.
func M512MaxEpu16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm_mask_max_epu8'.
// Requires AVX512BW.
func MaskMaxEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm_maskz_max_epu8'.
// Requires AVX512BW.
func MaskzMaxEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm256_mask_max_epu8'.
// Requires AVX512BW.
func M256MaskMaxEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm256_maskz_max_epu8'.
// Requires AVX512BW.
func M256MaskzMaxEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm512_mask_max_epu8'.
// Requires AVX512BW.
func M512MaskMaxEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] > b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm512_maskz_max_epu8'.
// Requires AVX512BW.
func M512MaskzMaxEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUB'. Intrinsic: '_mm512_max_epu8'.
// Requires AVX512BW.
func M512MaxEpu8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm_mask_min_epi16'.
// Requires AVX512BW.
func MaskMinEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm_maskz_min_epi16'.
// Requires AVX512BW.
func MaskzMinEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm256_mask_min_epi16'.
// Requires AVX512BW.
func M256MaskMinEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm256_maskz_min_epi16'.
// Requires AVX512BW.
func M256MaskzMinEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm512_mask_min_epi16'.
// Requires AVX512BW.
func M512MaskMinEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm512_maskz_min_epi16'.
// Requires AVX512BW.
func M512MaskzMinEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSW'. Intrinsic: '_mm512_min_epi16'.
// Requires AVX512BW.
func M512MinEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm_mask_min_epi8'.
// Requires AVX512BW.
func MaskMinEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm_maskz_min_epi8'.
// Requires AVX512BW.
func MaskzMinEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm256_mask_min_epi8'.
// Requires AVX512BW.
func M256MaskMinEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm256_maskz_min_epi8'.
// Requires AVX512BW.
func M256MaskzMinEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm512_mask_min_epi8'.
// Requires AVX512BW.
func M512MaskMinEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm512_maskz_min_epi8'.
// Requires AVX512BW.
func M512MaskzMinEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm512_min_epi8'.
// Requires AVX512BW.
func M512MinEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm_mask_min_epu16'.
// Requires AVX512BW.
func MaskMinEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm_maskz_min_epu16'.
// Requires AVX512BW.
func MaskzMinEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm256_mask_min_epu16'.
// Requires AVX512BW.
func M256MaskMinEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed minimum values in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm256_maskz_min_epu16'.
// Requires AVX512BW.
func M256MaskzMinEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm512_mask_min_epu16'.
// Requires AVX512BW.
func M512MaskMinEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b',
// and store packed minimum values in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF a[i+15:i] < b[i+15:i]
//					dst[i+15:i] := a[i+15:i]
//				ELSE
//					dst[i+15:i] := b[i+15:i]
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm512_maskz_min_epu16'.
// Requires AVX512BW.
func M512MaskzMinEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUW'. Intrinsic: '_mm512_min_epu16'.
// Requires AVX512BW.
func M512MinEpu16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm_mask_min_epu8'.
// Requires AVX512BW.
func MaskMinEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm_maskz_min_epu8'.
// Requires AVX512BW.
func MaskzMinEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm256_mask_min_epu8'.
// Requires AVX512BW.
func M256MaskMinEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm256_maskz_min_epu8'.
// Requires AVX512BW.
func M256MaskzMinEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm512_mask_min_epu8'.
// Requires AVX512BW.
func M512MaskMinEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF a[i+7:i] < b[i+7:i]
//					dst[i+7:i] := a[i+7:i]
//				ELSE
//					dst[i+7:i] := b[i+7:i]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm512_maskz_min_epu8'.
// Requires AVX512BW.
func M512MaskzMinEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUB'. Intrinsic: '_mm512_min_epu8'.
// Requires AVX512BW.
func M512MinEpu8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMovEpi16: Move packed 16-bit integers from 'a' into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm_mask_mov_epi16'.
// Requires AVX512BW.
func MaskMovEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMovEpi16: Move packed 16-bit integers from 'a' into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm_maskz_mov_epi16'.
// Requires AVX512BW.
func MaskzMovEpi16(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMovEpi16: Move packed 16-bit integers from 'a' into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm256_mask_mov_epi16'.
// Requires AVX512BW.
func M256MaskMovEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMovEpi16: Move packed 16-bit integers from 'a' into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm256_maskz_mov_epi16'.
// Requires AVX512BW.
func M256MaskzMovEpi16(k x86.Mmask16, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMovEpi16: Move packed 16-bit integers from 'a' into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm512_mask_mov_epi16'.
// Requires AVX512BW.
func M512MaskMovEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMovEpi16: Move packed 16-bit integers from 'a' into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm512_maskz_mov_epi16'.
// Requires AVX512BW.
func M512MaskzMovEpi16(k x86.Mmask32, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMovEpi8: Move packed 8-bit integers from 'a' into 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm_mask_mov_epi8'.
// Requires AVX512BW.
func MaskMovEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMovEpi8: Move packed 8-bit integers from 'a' into 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm_maskz_mov_epi8'.
// Requires AVX512BW.
func MaskzMovEpi8(k x86.Mmask16, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMovEpi8: Move packed 8-bit integers from 'a' into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm256_mask_mov_epi8'.
// Requires AVX512BW.
func M256MaskMovEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMovEpi8: Move packed 8-bit integers from 'a' into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm256_maskz_mov_epi8'.
// Requires AVX512BW.
func M256MaskzMovEpi8(k x86.Mmask32, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMovEpi8: Move packed 8-bit integers from 'a' into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm512_mask_mov_epi8'.
// Requires AVX512BW.
func M512MaskMovEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMovEpi8: Move packed 8-bit integers from 'a' into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm512_maskz_mov_epi8'.
// Requires AVX512BW.
func M512MaskzMovEpi8(k x86.Mmask64, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// Movepi16Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 16-bit integer in 'a'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF a[i+15]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPMOVW2M'. Intrinsic: '_mm_movepi16_mask'.
// Requires AVX512BW.
func Movepi16Mask(a x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256Movepi16Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 16-bit integer in 'a'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF a[i+15]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPMOVW2M'. Intrinsic: '_mm256_movepi16_mask'.
// Requires AVX512BW.
func M256Movepi16Mask(a x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Movepi16Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 16-bit integer in 'a'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF a[i+15]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPMOVW2M'. Intrinsic: '_mm512_movepi16_mask'.
// Requires AVX512BW.
func M512Movepi16Mask(a x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// Movepi8Mask: Set each bit of mask register 'k' based on the most significant
// bit of the corresponding packed 8-bit integer in 'a'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF a[i+7]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPMOVB2M'. Intrinsic: '_mm_movepi8_mask'.
// Requires AVX512BW.
func Movepi8Mask(a x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256Movepi8Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 8-bit integer in 'a'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF a[i+7]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPMOVB2M'. Intrinsic: '_mm256_movepi8_mask'.
// Requires AVX512BW.
func M256Movepi8Mask(a x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512Movepi8Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 8-bit integer in 'a'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF a[i+7]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPMOVB2M'. Intrinsic: '_mm512_movepi8_mask'.
// Requires AVX512BW.
func M512Movepi8Mask(a x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// MovmEpi16: Set each packed 16-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := 0xFFFF
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVM2W'. Intrinsic: '_mm_movm_epi16'.
// Requires AVX512BW.
func MovmEpi16(k x86.Mmask8) (dst x86.M128i) {
	panic("not implemented")
}


// M256MovmEpi16: Set each packed 16-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := 0xFFFF
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVM2W'. Intrinsic: '_mm256_movm_epi16'.
// Requires AVX512BW.
func M256MovmEpi16(k x86.Mmask16) (dst x86.M256i) {
	panic("not implemented")
}


// M512MovmEpi16: Set each packed 16-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := 0xFFFF
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVM2W'. Intrinsic: '_mm512_movm_epi16'.
// Requires AVX512BW.
func M512MovmEpi16(k x86.Mmask32) (dst x86.M512i) {
	panic("not implemented")
}


// M256MovmEpi8: Set each packed 8-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := 0xFF
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVM2B'. Intrinsic: '_mm256_movm_epi8'.
// Requires AVX512BW.
func M256MovmEpi8(k x86.Mmask32) (dst x86.M256i) {
	panic("not implemented")
}


// M512MovmEpi8: Set each packed 8-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := 0xFF
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVM2B'. Intrinsic: '_mm512_movm_epi8'.
// Requires AVX512BW.
func M512MovmEpi8(k x86.Mmask64) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm_mask_mulhi_epi16'.
// Requires AVX512BW.
func MaskMulhiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm_maskz_mulhi_epi16'.
// Requires AVX512BW.
func MaskzMulhiEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm256_mask_mulhi_epi16'.
// Requires AVX512BW.
func M256MaskMulhiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm256_maskz_mulhi_epi16'.
// Requires AVX512BW.
func M256MaskzMulhiEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm512_mask_mulhi_epi16'.
// Requires AVX512BW.
func M512MaskMulhiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm512_maskz_mulhi_epi16'.
// Requires AVX512BW.
func M512MaskzMulhiEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHW'. Intrinsic: '_mm512_mulhi_epi16'.
// Requires AVX512BW.
func M512MulhiEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm_mask_mulhi_epu16'.
// Requires AVX512BW.
func MaskMulhiEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and
// 'b', producing intermediate 32-bit integers, and store the high 16 bits of
// the intermediate integers in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := o
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm_maskz_mulhi_epu16'.
// Requires AVX512BW.
func MaskzMulhiEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and
// 'b', producing intermediate 32-bit integers, and store the high 16 bits of
// the intermediate integers in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm256_mask_mulhi_epu16'.
// Requires AVX512BW.
func M256MaskMulhiEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and
// 'b', producing intermediate 32-bit integers, and store the high 16 bits of
// the intermediate integers in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := o
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm256_maskz_mulhi_epu16'.
// Requires AVX512BW.
func M256MaskzMulhiEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and
// 'b', producing intermediate 32-bit integers, and store the high 16 bits of
// the intermediate integers in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm512_mask_mulhi_epu16'.
// Requires AVX512BW.
func M512MaskMulhiEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and
// 'b', producing intermediate 32-bit integers, and store the high 16 bits of
// the intermediate integers in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[31:16]
//			ELSE
//				dst[i+15:i] := o
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm512_maskz_mulhi_epu16'.
// Requires AVX512BW.
func M512MaskzMulhiEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUW'. Intrinsic: '_mm512_mulhi_epu16'.
// Requires AVX512BW.
func M512MulhiEpu16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//				dst[i+15:i] := tmp[16:1]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm_mask_mulhrs_epi16'.
// Requires AVX512BW.
func MaskMulhrsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//				dst[i+15:i] := tmp[16:1]
//			ELSE
//				dst[i+15:i] := 9
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm_maskz_mulhrs_epi16'.
// Requires AVX512BW.
func MaskzMulhrsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b',
// producing intermediate signed 32-bit integers. Truncate each intermediate
// integer to the 18 most significant bits, round by adding 1, and store bits
// [16:1] to 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//				dst[i+15:i] := tmp[16:1]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm256_mask_mulhrs_epi16'.
// Requires AVX512BW.
func M256MaskMulhrsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b',
// producing intermediate signed 32-bit integers. Truncate each intermediate
// integer to the 18 most significant bits, round by adding 1, and store bits
// [16:1] to 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//				dst[i+15:i] := tmp[16:1]
//			ELSE
//				dst[i+15:i] := 9
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm256_maskz_mulhrs_epi16'.
// Requires AVX512BW.
func M256MaskzMulhrsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b',
// producing intermediate signed 32-bit integers. Truncate each intermediate
// integer to the 18 most significant bits, round by adding 1, and store bits
// [16:1] to 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//				dst[i+15:i] := tmp[16:1]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm512_mask_mulhrs_epi16'.
// Requires AVX512BW.
func M512MaskMulhrsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b',
// producing intermediate signed 32-bit integers. Truncate each intermediate
// integer to the 18 most significant bits, round by adding 1, and store bits
// [16:1] to 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//				dst[i+15:i] := tmp[16:1]
//			ELSE
//				dst[i+15:i] := 9
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm512_maskz_mulhrs_epi16'.
// Requires AVX512BW.
func M512MaskzMulhrsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			tmp[31:0] := ((a[i+15:i] * b[i+15:i]) >> 14) + 1
//			dst[i+15:i] := tmp[16:1]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHRSW'. Intrinsic: '_mm512_mulhrs_epi16'.
// Requires AVX512BW.
func M512MulhrsEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskMulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm_mask_mullo_epi16'.
// Requires AVX512BW.
func MaskMulloEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzMulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm_maskz_mullo_epi16'.
// Requires AVX512BW.
func MaskzMulloEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskMulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm256_mask_mullo_epi16'.
// Requires AVX512BW.
func M256MaskMulloEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzMulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm256_maskz_mullo_epi16'.
// Requires AVX512BW.
func M256MaskzMulloEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskMulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm512_mask_mullo_epi16'.
// Requires AVX512BW.
func M512MaskMulloEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				tmp[31:0] := a[i+15:i] * b[i+15:i]
//				dst[i+15:i] := tmp[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm512_maskz_mullo_epi16'.
// Requires AVX512BW.
func M512MaskzMulloEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the low 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[15:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLW'. Intrinsic: '_mm512_mullo_epi16'.
// Requires AVX512BW.
func M512MulloEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskPacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
// 8-bit integers using signed saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm_mask_packs_epi16'.
// Requires AVX512BW.
func MaskPacksEpi16(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
// 8-bit integers using signed saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm_maskz_packs_epi16'.
// Requires AVX512BW.
func MaskzPacksEpi16(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using signed saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_Int8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_Int8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_Int8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_Int8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_Int8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_Int8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_Int8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_Int8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_Int8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_Int8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_Int8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_Int8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_Int8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_Int8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_Int8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_Int8 (b[255:240])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm256_mask_packs_epi16'.
// Requires AVX512BW.
func M256MaskPacksEpi16(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using signed saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_Int8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_Int8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_Int8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_Int8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_Int8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_Int8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_Int8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_Int8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_Int8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_Int8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_Int8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_Int8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_Int8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_Int8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_Int8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_Int8 (b[255:240])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm256_maskz_packs_epi16'.
// Requires AVX512BW.
func M256MaskzPacksEpi16(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskPacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using signed saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_Int8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_Int8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_Int8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_Int8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_Int8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_Int8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_Int8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_Int8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_Int8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_Int8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_Int8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_Int8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_Int8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_Int8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_Int8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_Int8 (b[255:240])
//		tmp_dst[263:256] := Saturate_Int16_To_Int8 (a[271:256])
//		tmp_dst[271:264] := Saturate_Int16_To_Int8 (a[287:272])
//		tmp_dst[279:272] := Saturate_Int16_To_Int8 (a[303:288])
//		tmp_dst[287:280] := Saturate_Int16_To_Int8 (a[319:304])
//		tmp_dst[295:288] := Saturate_Int16_To_Int8 (a[335:320])
//		tmp_dst[303:296] := Saturate_Int16_To_Int8 (a[351:336])
//		tmp_dst[311:304] := Saturate_Int16_To_Int8 (a[367:352])
//		tmp_dst[319:312] := Saturate_Int16_To_Int8 (a[383:368])
//		tmp_dst[327:320] := Saturate_Int16_To_Int8 (b[271:256])
//		tmp_dst[335:328] := Saturate_Int16_To_Int8 (b[287:272])
//		tmp_dst[343:336] := Saturate_Int16_To_Int8 (b[303:288])
//		tmp_dst[351:344] := Saturate_Int16_To_Int8 (b[319:304])
//		tmp_dst[359:352] := Saturate_Int16_To_Int8 (b[335:320])
//		tmp_dst[367:360] := Saturate_Int16_To_Int8 (b[351:336])
//		tmp_dst[375:368] := Saturate_Int16_To_Int8 (b[367:352])
//		tmp_dst[383:376] := Saturate_Int16_To_Int8 (b[383:368])
//		tmp_dst[391:384] := Saturate_Int16_To_Int8 (a[399:384])
//		tmp_dst[399:392] := Saturate_Int16_To_Int8 (a[415:400])
//		tmp_dst[407:400] := Saturate_Int16_To_Int8 (a[431:416])
//		tmp_dst[415:408] := Saturate_Int16_To_Int8 (a[447:432])
//		tmp_dst[423:416] := Saturate_Int16_To_Int8 (a[463:448])
//		tmp_dst[431:424] := Saturate_Int16_To_Int8 (a[479:464])
//		tmp_dst[439:432] := Saturate_Int16_To_Int8 (a[495:480])
//		tmp_dst[447:440] := Saturate_Int16_To_Int8 (a[511:496])
//		tmp_dst[455:448] := Saturate_Int16_To_Int8 (b[399:384])
//		tmp_dst[463:456] := Saturate_Int16_To_Int8 (b[415:400])
//		tmp_dst[471:464] := Saturate_Int16_To_Int8 (b[431:416])
//		tmp_dst[479:472] := Saturate_Int16_To_Int8 (b[447:432])
//		tmp_dst[487:480] := Saturate_Int16_To_Int8 (b[463:448])
//		tmp_dst[495:488] := Saturate_Int16_To_Int8 (b[479:464])
//		tmp_dst[503:496] := Saturate_Int16_To_Int8 (b[495:480])
//		tmp_dst[511:504] := Saturate_Int16_To_Int8 (b[511:496])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm512_mask_packs_epi16'.
// Requires AVX512BW.
func M512MaskPacksEpi16(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzPacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using signed saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_Int8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_Int8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_Int8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_Int8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_Int8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_Int8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_Int8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_Int8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_Int8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_Int8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_Int8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_Int8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_Int8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_Int8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_Int8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_Int8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_Int8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_Int8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_Int8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_Int8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_Int8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_Int8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_Int8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_Int8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_Int8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_Int8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_Int8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_Int8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_Int8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_Int8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_Int8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_Int8 (b[255:240])
//		tmp_dst[263:256] := Saturate_Int16_To_Int8 (a[271:256])
//		tmp_dst[271:264] := Saturate_Int16_To_Int8 (a[287:272])
//		tmp_dst[279:272] := Saturate_Int16_To_Int8 (a[303:288])
//		tmp_dst[287:280] := Saturate_Int16_To_Int8 (a[319:304])
//		tmp_dst[295:288] := Saturate_Int16_To_Int8 (a[335:320])
//		tmp_dst[303:296] := Saturate_Int16_To_Int8 (a[351:336])
//		tmp_dst[311:304] := Saturate_Int16_To_Int8 (a[367:352])
//		tmp_dst[319:312] := Saturate_Int16_To_Int8 (a[383:368])
//		tmp_dst[327:320] := Saturate_Int16_To_Int8 (b[271:256])
//		tmp_dst[335:328] := Saturate_Int16_To_Int8 (b[287:272])
//		tmp_dst[343:336] := Saturate_Int16_To_Int8 (b[303:288])
//		tmp_dst[351:344] := Saturate_Int16_To_Int8 (b[319:304])
//		tmp_dst[359:352] := Saturate_Int16_To_Int8 (b[335:320])
//		tmp_dst[367:360] := Saturate_Int16_To_Int8 (b[351:336])
//		tmp_dst[375:368] := Saturate_Int16_To_Int8 (b[367:352])
//		tmp_dst[383:376] := Saturate_Int16_To_Int8 (b[383:368])
//		tmp_dst[391:384] := Saturate_Int16_To_Int8 (a[399:384])
//		tmp_dst[399:392] := Saturate_Int16_To_Int8 (a[415:400])
//		tmp_dst[407:400] := Saturate_Int16_To_Int8 (a[431:416])
//		tmp_dst[415:408] := Saturate_Int16_To_Int8 (a[447:432])
//		tmp_dst[423:416] := Saturate_Int16_To_Int8 (a[463:448])
//		tmp_dst[431:424] := Saturate_Int16_To_Int8 (a[479:464])
//		tmp_dst[439:432] := Saturate_Int16_To_Int8 (a[495:480])
//		tmp_dst[447:440] := Saturate_Int16_To_Int8 (a[511:496])
//		tmp_dst[455:448] := Saturate_Int16_To_Int8 (b[399:384])
//		tmp_dst[463:456] := Saturate_Int16_To_Int8 (b[415:400])
//		tmp_dst[471:464] := Saturate_Int16_To_Int8 (b[431:416])
//		tmp_dst[479:472] := Saturate_Int16_To_Int8 (b[447:432])
//		tmp_dst[487:480] := Saturate_Int16_To_Int8 (b[463:448])
//		tmp_dst[495:488] := Saturate_Int16_To_Int8 (b[479:464])
//		tmp_dst[503:496] := Saturate_Int16_To_Int8 (b[495:480])
//		tmp_dst[511:504] := Saturate_Int16_To_Int8 (b[511:496])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm512_maskz_packs_epi16'.
// Requires AVX512BW.
func M512MaskzPacksEpi16(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512PacksEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
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
//		dst[263:256] := Saturate_Int16_To_Int8 (a[271:256])
//		dst[271:264] := Saturate_Int16_To_Int8 (a[287:272])
//		dst[279:272] := Saturate_Int16_To_Int8 (a[303:288])
//		dst[287:280] := Saturate_Int16_To_Int8 (a[319:304])
//		dst[295:288] := Saturate_Int16_To_Int8 (a[335:320])
//		dst[303:296] := Saturate_Int16_To_Int8 (a[351:336])
//		dst[311:304] := Saturate_Int16_To_Int8 (a[367:352])
//		dst[319:312] := Saturate_Int16_To_Int8 (a[383:368])
//		dst[327:320] := Saturate_Int16_To_Int8 (b[271:256])
//		dst[335:328] := Saturate_Int16_To_Int8 (b[287:272])
//		dst[343:336] := Saturate_Int16_To_Int8 (b[303:288])
//		dst[351:344] := Saturate_Int16_To_Int8 (b[319:304])
//		dst[359:352] := Saturate_Int16_To_Int8 (b[335:320])
//		dst[367:360] := Saturate_Int16_To_Int8 (b[351:336])
//		dst[375:368] := Saturate_Int16_To_Int8 (b[367:352])
//		dst[383:376] := Saturate_Int16_To_Int8 (b[383:368])
//		dst[391:384] := Saturate_Int16_To_Int8 (a[399:384])
//		dst[399:392] := Saturate_Int16_To_Int8 (a[415:400])
//		dst[407:400] := Saturate_Int16_To_Int8 (a[431:416])
//		dst[415:408] := Saturate_Int16_To_Int8 (a[447:432])
//		dst[423:416] := Saturate_Int16_To_Int8 (a[463:448])
//		dst[431:424] := Saturate_Int16_To_Int8 (a[479:464])
//		dst[439:432] := Saturate_Int16_To_Int8 (a[495:480])
//		dst[447:440] := Saturate_Int16_To_Int8 (a[511:496])
//		dst[455:448] := Saturate_Int16_To_Int8 (b[399:384])
//		dst[463:456] := Saturate_Int16_To_Int8 (b[415:400])
//		dst[471:464] := Saturate_Int16_To_Int8 (b[431:416])
//		dst[479:472] := Saturate_Int16_To_Int8 (b[447:432])
//		dst[487:480] := Saturate_Int16_To_Int8 (b[463:448])
//		dst[495:488] := Saturate_Int16_To_Int8 (b[479:464])
//		dst[503:496] := Saturate_Int16_To_Int8 (b[495:480])
//		dst[511:504] := Saturate_Int16_To_Int8 (b[511:496])
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSSWB'. Intrinsic: '_mm512_packs_epi16'.
// Requires AVX512BW.
func M512PacksEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskPacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using signed saturation, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm_mask_packs_epi32'.
// Requires AVX512BW.
func MaskPacksEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using signed saturation, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm_maskz_packs_epi32'.
// Requires AVX512BW.
func MaskzPacksEpi32(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using signed saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_Int16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_Int16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_Int16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_Int16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_Int16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_Int16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_Int16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_Int16 (b[255:224])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm256_mask_packs_epi32'.
// Requires AVX512BW.
func M256MaskPacksEpi32(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using signed saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_Int16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_Int16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_Int16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_Int16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_Int16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_Int16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_Int16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_Int16 (b[255:224])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm256_maskz_packs_epi32'.
// Requires AVX512BW.
func M256MaskzPacksEpi32(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskPacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using signed saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_Int16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_Int16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_Int16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_Int16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_Int16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_Int16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_Int16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_Int16 (b[255:224])
//		tmp_dst[271:256] := Saturate_Int32_To_Int16 (a[287:256])
//		tmp_dst[287:272] := Saturate_Int32_To_Int16 (a[319:288])
//		tmp_dst[303:288] := Saturate_Int32_To_Int16 (a[351:320])
//		tmp_dst[319:304] := Saturate_Int32_To_Int16 (a[383:352])
//		tmp_dst[335:320] := Saturate_Int32_To_Int16 (b[287:256])
//		tmp_dst[351:336] := Saturate_Int32_To_Int16 (b[319:288])
//		tmp_dst[367:352] := Saturate_Int32_To_Int16 (b[351:320])
//		tmp_dst[383:368] := Saturate_Int32_To_Int16 (b[383:352])
//		tmp_dst[399:384] := Saturate_Int32_To_Int16 (a[415:384])
//		tmp_dst[415:400] := Saturate_Int32_To_Int16 (a[447:416])
//		tmp_dst[431:416] := Saturate_Int32_To_Int16 (a[479:448])
//		tmp_dst[447:432] := Saturate_Int32_To_Int16 (a[511:480])
//		tmp_dst[463:448] := Saturate_Int32_To_Int16 (b[415:384])
//		tmp_dst[479:464] := Saturate_Int32_To_Int16 (b[447:416])
//		tmp_dst[495:480] := Saturate_Int32_To_Int16 (b[479:448])
//		tmp_dst[511:496] := Saturate_Int32_To_Int16 (b[511:480])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm512_mask_packs_epi32'.
// Requires AVX512BW.
func M512MaskPacksEpi32(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzPacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using signed saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_Int16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_Int16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_Int16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_Int16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_Int16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_Int16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_Int16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_Int16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_Int16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_Int16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_Int16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_Int16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_Int16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_Int16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_Int16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_Int16 (b[255:224])
//		tmp_dst[271:256] := Saturate_Int32_To_Int16 (a[287:256])
//		tmp_dst[287:272] := Saturate_Int32_To_Int16 (a[319:288])
//		tmp_dst[303:288] := Saturate_Int32_To_Int16 (a[351:320])
//		tmp_dst[319:304] := Saturate_Int32_To_Int16 (a[383:352])
//		tmp_dst[335:320] := Saturate_Int32_To_Int16 (b[287:256])
//		tmp_dst[351:336] := Saturate_Int32_To_Int16 (b[319:288])
//		tmp_dst[367:352] := Saturate_Int32_To_Int16 (b[351:320])
//		tmp_dst[383:368] := Saturate_Int32_To_Int16 (b[383:352])
//		tmp_dst[399:384] := Saturate_Int32_To_Int16 (a[415:384])
//		tmp_dst[415:400] := Saturate_Int32_To_Int16 (a[447:416])
//		tmp_dst[431:416] := Saturate_Int32_To_Int16 (a[479:448])
//		tmp_dst[447:432] := Saturate_Int32_To_Int16 (a[511:480])
//		tmp_dst[463:448] := Saturate_Int32_To_Int16 (b[415:384])
//		tmp_dst[479:464] := Saturate_Int32_To_Int16 (b[447:416])
//		tmp_dst[495:480] := Saturate_Int32_To_Int16 (b[479:448])
//		tmp_dst[511:496] := Saturate_Int32_To_Int16 (b[511:480])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm512_maskz_packs_epi32'.
// Requires AVX512BW.
func M512MaskzPacksEpi32(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512PacksEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
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
//		dst[271:256] := Saturate_Int32_To_Int16 (a[287:256])
//		dst[287:272] := Saturate_Int32_To_Int16 (a[319:288])
//		dst[303:288] := Saturate_Int32_To_Int16 (a[351:320])
//		dst[319:304] := Saturate_Int32_To_Int16 (a[383:352])
//		dst[335:320] := Saturate_Int32_To_Int16 (b[287:256])
//		dst[351:336] := Saturate_Int32_To_Int16 (b[319:288])
//		dst[367:352] := Saturate_Int32_To_Int16 (b[351:320])
//		dst[383:368] := Saturate_Int32_To_Int16 (b[383:352])
//		dst[399:384] := Saturate_Int32_To_Int16 (a[415:384])
//		dst[415:400] := Saturate_Int32_To_Int16 (a[447:416])
//		dst[431:416] := Saturate_Int32_To_Int16 (a[479:448])
//		dst[447:432] := Saturate_Int32_To_Int16 (a[511:480])
//		dst[463:448] := Saturate_Int32_To_Int16 (b[415:384])
//		dst[479:464] := Saturate_Int32_To_Int16 (b[447:416])
//		dst[495:480] := Saturate_Int32_To_Int16 (b[479:448])
//		dst[511:496] := Saturate_Int32_To_Int16 (b[511:480])
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSSDW'. Intrinsic: '_mm512_packs_epi32'.
// Requires AVX512BW.
func M512PacksEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskPackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
// 8-bit integers using unsigned saturation, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm_mask_packus_epi16'.
// Requires AVX512BW.
func MaskPackusEpi16(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
// 8-bit integers using unsigned saturation, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm_maskz_packus_epi16'.
// Requires AVX512BW.
func MaskzPackusEpi16(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using unsigned saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_UnsignedInt8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_UnsignedInt8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_UnsignedInt8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_UnsignedInt8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_UnsignedInt8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_UnsignedInt8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_UnsignedInt8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_UnsignedInt8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_UnsignedInt8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_UnsignedInt8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_UnsignedInt8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_UnsignedInt8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_UnsignedInt8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_UnsignedInt8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_UnsignedInt8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_UnsignedInt8 (b[255:240])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm256_mask_packus_epi16'.
// Requires AVX512BW.
func M256MaskPackusEpi16(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using unsigned saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_UnsignedInt8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_UnsignedInt8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_UnsignedInt8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_UnsignedInt8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_UnsignedInt8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_UnsignedInt8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_UnsignedInt8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_UnsignedInt8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_UnsignedInt8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_UnsignedInt8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_UnsignedInt8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_UnsignedInt8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_UnsignedInt8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_UnsignedInt8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_UnsignedInt8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_UnsignedInt8 (b[255:240])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm256_maskz_packus_epi16'.
// Requires AVX512BW.
func M256MaskzPackusEpi16(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskPackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using unsigned saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_UnsignedInt8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_UnsignedInt8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_UnsignedInt8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_UnsignedInt8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_UnsignedInt8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_UnsignedInt8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_UnsignedInt8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_UnsignedInt8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_UnsignedInt8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_UnsignedInt8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_UnsignedInt8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_UnsignedInt8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_UnsignedInt8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_UnsignedInt8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_UnsignedInt8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_UnsignedInt8 (b[255:240])
//		tmp_dst[263:256] := Saturate_Int16_To_UnsignedInt8 (a[271:256])
//		tmp_dst[271:264] := Saturate_Int16_To_UnsignedInt8 (a[287:272])
//		tmp_dst[279:272] := Saturate_Int16_To_UnsignedInt8 (a[303:288])
//		tmp_dst[287:280] := Saturate_Int16_To_UnsignedInt8 (a[319:304])
//		tmp_dst[295:288] := Saturate_Int16_To_UnsignedInt8 (a[335:320])
//		tmp_dst[303:296] := Saturate_Int16_To_UnsignedInt8 (a[351:336])
//		tmp_dst[311:304] := Saturate_Int16_To_UnsignedInt8 (a[367:352])
//		tmp_dst[319:312] := Saturate_Int16_To_UnsignedInt8 (a[383:368])
//		tmp_dst[327:320] := Saturate_Int16_To_UnsignedInt8 (b[271:256])
//		tmp_dst[335:328] := Saturate_Int16_To_UnsignedInt8 (b[287:272])
//		tmp_dst[343:336] := Saturate_Int16_To_UnsignedInt8 (b[303:288])
//		tmp_dst[351:344] := Saturate_Int16_To_UnsignedInt8 (b[319:304])
//		tmp_dst[359:352] := Saturate_Int16_To_UnsignedInt8 (b[335:320])
//		tmp_dst[367:360] := Saturate_Int16_To_UnsignedInt8 (b[351:336])
//		tmp_dst[375:368] := Saturate_Int16_To_UnsignedInt8 (b[367:352])
//		tmp_dst[383:376] := Saturate_Int16_To_UnsignedInt8 (b[383:368])
//		tmp_dst[391:384] := Saturate_Int16_To_UnsignedInt8 (a[399:384])
//		tmp_dst[399:392] := Saturate_Int16_To_UnsignedInt8 (a[415:400])
//		tmp_dst[407:400] := Saturate_Int16_To_UnsignedInt8 (a[431:416])
//		tmp_dst[415:408] := Saturate_Int16_To_UnsignedInt8 (a[447:432])
//		tmp_dst[423:416] := Saturate_Int16_To_UnsignedInt8 (a[463:448])
//		tmp_dst[431:424] := Saturate_Int16_To_UnsignedInt8 (a[479:464])
//		tmp_dst[439:432] := Saturate_Int16_To_UnsignedInt8 (a[495:480])
//		tmp_dst[447:440] := Saturate_Int16_To_UnsignedInt8 (a[511:496])
//		tmp_dst[455:448] := Saturate_Int16_To_UnsignedInt8 (b[399:384])
//		tmp_dst[463:456] := Saturate_Int16_To_UnsignedInt8 (b[415:400])
//		tmp_dst[471:464] := Saturate_Int16_To_UnsignedInt8 (b[431:416])
//		tmp_dst[479:472] := Saturate_Int16_To_UnsignedInt8 (b[447:432])
//		tmp_dst[487:480] := Saturate_Int16_To_UnsignedInt8 (b[463:448])
//		tmp_dst[495:488] := Saturate_Int16_To_UnsignedInt8 (b[479:464])
//		tmp_dst[503:496] := Saturate_Int16_To_UnsignedInt8 (b[495:480])
//		tmp_dst[511:504] := Saturate_Int16_To_UnsignedInt8 (b[511:496])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm512_mask_packus_epi16'.
// Requires AVX512BW.
func M512MaskPackusEpi16(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzPackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to
// packed 8-bit integers using unsigned saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[7:0] := Saturate_Int16_To_UnsignedInt8 (a[15:0])
//		tmp_dst[15:8] := Saturate_Int16_To_UnsignedInt8 (a[31:16])
//		tmp_dst[23:16] := Saturate_Int16_To_UnsignedInt8 (a[47:32])
//		tmp_dst[31:24] := Saturate_Int16_To_UnsignedInt8 (a[63:48])
//		tmp_dst[39:32] := Saturate_Int16_To_UnsignedInt8 (a[79:64])
//		tmp_dst[47:40] := Saturate_Int16_To_UnsignedInt8 (a[95:80])
//		tmp_dst[55:48] := Saturate_Int16_To_UnsignedInt8 (a[111:96])
//		tmp_dst[63:56] := Saturate_Int16_To_UnsignedInt8 (a[127:112])
//		tmp_dst[71:64] := Saturate_Int16_To_UnsignedInt8 (b[15:0])
//		tmp_dst[79:72] := Saturate_Int16_To_UnsignedInt8 (b[31:16])
//		tmp_dst[87:80] := Saturate_Int16_To_UnsignedInt8 (b[47:32])
//		tmp_dst[95:88] := Saturate_Int16_To_UnsignedInt8 (b[63:48])
//		tmp_dst[103:96] := Saturate_Int16_To_UnsignedInt8 (b[79:64])
//		tmp_dst[111:104] := Saturate_Int16_To_UnsignedInt8 (b[95:80])
//		tmp_dst[119:112] := Saturate_Int16_To_UnsignedInt8 (b[111:96])
//		tmp_dst[127:120] := Saturate_Int16_To_UnsignedInt8 (b[127:112])
//		tmp_dst[135:128] := Saturate_Int16_To_UnsignedInt8 (a[143:128])
//		tmp_dst[143:136] := Saturate_Int16_To_UnsignedInt8 (a[159:144])
//		tmp_dst[151:144] := Saturate_Int16_To_UnsignedInt8 (a[175:160])
//		tmp_dst[159:152] := Saturate_Int16_To_UnsignedInt8 (a[191:176])
//		tmp_dst[167:160] := Saturate_Int16_To_UnsignedInt8 (a[207:192])
//		tmp_dst[175:168] := Saturate_Int16_To_UnsignedInt8 (a[223:208])
//		tmp_dst[183:176] := Saturate_Int16_To_UnsignedInt8 (a[239:224])
//		tmp_dst[191:184] := Saturate_Int16_To_UnsignedInt8 (a[255:240])
//		tmp_dst[199:192] := Saturate_Int16_To_UnsignedInt8 (b[143:128])
//		tmp_dst[207:200] := Saturate_Int16_To_UnsignedInt8 (b[159:144])
//		tmp_dst[215:208] := Saturate_Int16_To_UnsignedInt8 (b[175:160])
//		tmp_dst[223:216] := Saturate_Int16_To_UnsignedInt8 (b[191:176])
//		tmp_dst[231:224] := Saturate_Int16_To_UnsignedInt8 (b[207:192])
//		tmp_dst[239:232] := Saturate_Int16_To_UnsignedInt8 (b[223:208])
//		tmp_dst[247:240] := Saturate_Int16_To_UnsignedInt8 (b[239:224])
//		tmp_dst[255:248] := Saturate_Int16_To_UnsignedInt8 (b[255:240])
//		tmp_dst[263:256] := Saturate_Int16_To_UnsignedInt8 (a[271:256])
//		tmp_dst[271:264] := Saturate_Int16_To_UnsignedInt8 (a[287:272])
//		tmp_dst[279:272] := Saturate_Int16_To_UnsignedInt8 (a[303:288])
//		tmp_dst[287:280] := Saturate_Int16_To_UnsignedInt8 (a[319:304])
//		tmp_dst[295:288] := Saturate_Int16_To_UnsignedInt8 (a[335:320])
//		tmp_dst[303:296] := Saturate_Int16_To_UnsignedInt8 (a[351:336])
//		tmp_dst[311:304] := Saturate_Int16_To_UnsignedInt8 (a[367:352])
//		tmp_dst[319:312] := Saturate_Int16_To_UnsignedInt8 (a[383:368])
//		tmp_dst[327:320] := Saturate_Int16_To_UnsignedInt8 (b[271:256])
//		tmp_dst[335:328] := Saturate_Int16_To_UnsignedInt8 (b[287:272])
//		tmp_dst[343:336] := Saturate_Int16_To_UnsignedInt8 (b[303:288])
//		tmp_dst[351:344] := Saturate_Int16_To_UnsignedInt8 (b[319:304])
//		tmp_dst[359:352] := Saturate_Int16_To_UnsignedInt8 (b[335:320])
//		tmp_dst[367:360] := Saturate_Int16_To_UnsignedInt8 (b[351:336])
//		tmp_dst[375:368] := Saturate_Int16_To_UnsignedInt8 (b[367:352])
//		tmp_dst[383:376] := Saturate_Int16_To_UnsignedInt8 (b[383:368])
//		tmp_dst[391:384] := Saturate_Int16_To_UnsignedInt8 (a[399:384])
//		tmp_dst[399:392] := Saturate_Int16_To_UnsignedInt8 (a[415:400])
//		tmp_dst[407:400] := Saturate_Int16_To_UnsignedInt8 (a[431:416])
//		tmp_dst[415:408] := Saturate_Int16_To_UnsignedInt8 (a[447:432])
//		tmp_dst[423:416] := Saturate_Int16_To_UnsignedInt8 (a[463:448])
//		tmp_dst[431:424] := Saturate_Int16_To_UnsignedInt8 (a[479:464])
//		tmp_dst[439:432] := Saturate_Int16_To_UnsignedInt8 (a[495:480])
//		tmp_dst[447:440] := Saturate_Int16_To_UnsignedInt8 (a[511:496])
//		tmp_dst[455:448] := Saturate_Int16_To_UnsignedInt8 (b[399:384])
//		tmp_dst[463:456] := Saturate_Int16_To_UnsignedInt8 (b[415:400])
//		tmp_dst[471:464] := Saturate_Int16_To_UnsignedInt8 (b[431:416])
//		tmp_dst[479:472] := Saturate_Int16_To_UnsignedInt8 (b[447:432])
//		tmp_dst[487:480] := Saturate_Int16_To_UnsignedInt8 (b[463:448])
//		tmp_dst[495:488] := Saturate_Int16_To_UnsignedInt8 (b[479:464])
//		tmp_dst[503:496] := Saturate_Int16_To_UnsignedInt8 (b[495:480])
//		tmp_dst[511:504] := Saturate_Int16_To_UnsignedInt8 (b[511:496])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm512_maskz_packus_epi16'.
// Requires AVX512BW.
func M512MaskzPackusEpi16(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512PackusEpi16: Convert packed 16-bit integers from 'a' and 'b' to packed
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
//		dst[263:256] := Saturate_Int16_To_UnsignedInt8 (a[271:256])
//		dst[271:264] := Saturate_Int16_To_UnsignedInt8 (a[287:272])
//		dst[279:272] := Saturate_Int16_To_UnsignedInt8 (a[303:288])
//		dst[287:280] := Saturate_Int16_To_UnsignedInt8 (a[319:304])
//		dst[295:288] := Saturate_Int16_To_UnsignedInt8 (a[335:320])
//		dst[303:296] := Saturate_Int16_To_UnsignedInt8 (a[351:336])
//		dst[311:304] := Saturate_Int16_To_UnsignedInt8 (a[367:352])
//		dst[319:312] := Saturate_Int16_To_UnsignedInt8 (a[383:368])
//		dst[327:320] := Saturate_Int16_To_UnsignedInt8 (b[271:256])
//		dst[335:328] := Saturate_Int16_To_UnsignedInt8 (b[287:272])
//		dst[343:336] := Saturate_Int16_To_UnsignedInt8 (b[303:288])
//		dst[351:344] := Saturate_Int16_To_UnsignedInt8 (b[319:304])
//		dst[359:352] := Saturate_Int16_To_UnsignedInt8 (b[335:320])
//		dst[367:360] := Saturate_Int16_To_UnsignedInt8 (b[351:336])
//		dst[375:368] := Saturate_Int16_To_UnsignedInt8 (b[367:352])
//		dst[383:376] := Saturate_Int16_To_UnsignedInt8 (b[383:368])
//		dst[391:384] := Saturate_Int16_To_UnsignedInt8 (a[399:384])
//		dst[399:392] := Saturate_Int16_To_UnsignedInt8 (a[415:400])
//		dst[407:400] := Saturate_Int16_To_UnsignedInt8 (a[431:416])
//		dst[415:408] := Saturate_Int16_To_UnsignedInt8 (a[447:432])
//		dst[423:416] := Saturate_Int16_To_UnsignedInt8 (a[463:448])
//		dst[431:424] := Saturate_Int16_To_UnsignedInt8 (a[479:464])
//		dst[439:432] := Saturate_Int16_To_UnsignedInt8 (a[495:480])
//		dst[447:440] := Saturate_Int16_To_UnsignedInt8 (a[511:496])
//		dst[455:448] := Saturate_Int16_To_UnsignedInt8 (b[399:384])
//		dst[463:456] := Saturate_Int16_To_UnsignedInt8 (b[415:400])
//		dst[471:464] := Saturate_Int16_To_UnsignedInt8 (b[431:416])
//		dst[479:472] := Saturate_Int16_To_UnsignedInt8 (b[447:432])
//		dst[487:480] := Saturate_Int16_To_UnsignedInt8 (b[463:448])
//		dst[495:488] := Saturate_Int16_To_UnsignedInt8 (b[479:464])
//		dst[503:496] := Saturate_Int16_To_UnsignedInt8 (b[495:480])
//		dst[511:504] := Saturate_Int16_To_UnsignedInt8 (b[511:496])
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKUSWB'. Intrinsic: '_mm512_packus_epi16'.
// Requires AVX512BW.
func M512PackusEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskPackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using unsigned saturation, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm_mask_packus_epi32'.
// Requires AVX512BW.
func MaskPackusEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using unsigned saturation, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm_maskz_packus_epi32'.
// Requires AVX512BW.
func MaskzPackusEpi32(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using unsigned saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_UnsignedInt16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_UnsignedInt16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_UnsignedInt16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_UnsignedInt16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_UnsignedInt16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_UnsignedInt16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_UnsignedInt16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_UnsignedInt16 (b[255:224])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm256_mask_packus_epi32'.
// Requires AVX512BW.
func M256MaskPackusEpi32(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using unsigned saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_UnsignedInt16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_UnsignedInt16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_UnsignedInt16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_UnsignedInt16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_UnsignedInt16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_UnsignedInt16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_UnsignedInt16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_UnsignedInt16 (b[255:224])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm256_maskz_packus_epi32'.
// Requires AVX512BW.
func M256MaskzPackusEpi32(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskPackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using unsigned saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_UnsignedInt16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_UnsignedInt16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_UnsignedInt16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_UnsignedInt16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_UnsignedInt16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_UnsignedInt16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_UnsignedInt16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_UnsignedInt16 (b[255:224])
//		tmp_dst[271:256] := Saturate_Int32_To_UnsignedInt16 (a[287:256])
//		tmp_dst[287:272] := Saturate_Int32_To_UnsignedInt16 (a[319:288])
//		tmp_dst[303:288] := Saturate_Int32_To_UnsignedInt16 (a[351:320])
//		tmp_dst[319:304] := Saturate_Int32_To_UnsignedInt16 (a[383:352])
//		tmp_dst[335:320] := Saturate_Int32_To_UnsignedInt16 (b[287:256])
//		tmp_dst[351:336] := Saturate_Int32_To_UnsignedInt16 (b[319:288])
//		tmp_dst[367:352] := Saturate_Int32_To_UnsignedInt16 (b[351:320])
//		tmp_dst[383:368] := Saturate_Int32_To_UnsignedInt16 (b[383:352])
//		tmp_dst[399:384] := Saturate_Int32_To_UnsignedInt16 (a[415:384])
//		tmp_dst[415:400] := Saturate_Int32_To_UnsignedInt16 (a[447:416])
//		tmp_dst[431:416] := Saturate_Int32_To_UnsignedInt16 (a[479:448])
//		tmp_dst[447:432] := Saturate_Int32_To_UnsignedInt16 (a[511:480])
//		tmp_dst[463:448] := Saturate_Int32_To_UnsignedInt16 (b[415:384])
//		tmp_dst[479:464] := Saturate_Int32_To_UnsignedInt16 (b[447:416])
//		tmp_dst[495:480] := Saturate_Int32_To_UnsignedInt16 (b[479:448])
//		tmp_dst[511:496] := Saturate_Int32_To_UnsignedInt16 (b[511:480])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm512_mask_packus_epi32'.
// Requires AVX512BW.
func M512MaskPackusEpi32(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzPackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to
// packed 16-bit integers using unsigned saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		tmp_dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		tmp_dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		tmp_dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		tmp_dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		tmp_dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		tmp_dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		tmp_dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		tmp_dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//		tmp_dst[143:128] := Saturate_Int32_To_UnsignedInt16 (a[159:128])
//		tmp_dst[159:144] := Saturate_Int32_To_UnsignedInt16 (a[191:160])
//		tmp_dst[175:160] := Saturate_Int32_To_UnsignedInt16 (a[223:192])
//		tmp_dst[191:176] := Saturate_Int32_To_UnsignedInt16 (a[255:224])
//		tmp_dst[207:192] := Saturate_Int32_To_UnsignedInt16 (b[159:128])
//		tmp_dst[223:208] := Saturate_Int32_To_UnsignedInt16 (b[191:160])
//		tmp_dst[239:224] := Saturate_Int32_To_UnsignedInt16 (b[223:192])
//		tmp_dst[255:240] := Saturate_Int32_To_UnsignedInt16 (b[255:224])
//		tmp_dst[271:256] := Saturate_Int32_To_UnsignedInt16 (a[287:256])
//		tmp_dst[287:272] := Saturate_Int32_To_UnsignedInt16 (a[319:288])
//		tmp_dst[303:288] := Saturate_Int32_To_UnsignedInt16 (a[351:320])
//		tmp_dst[319:304] := Saturate_Int32_To_UnsignedInt16 (a[383:352])
//		tmp_dst[335:320] := Saturate_Int32_To_UnsignedInt16 (b[287:256])
//		tmp_dst[351:336] := Saturate_Int32_To_UnsignedInt16 (b[319:288])
//		tmp_dst[367:352] := Saturate_Int32_To_UnsignedInt16 (b[351:320])
//		tmp_dst[383:368] := Saturate_Int32_To_UnsignedInt16 (b[383:352])
//		tmp_dst[399:384] := Saturate_Int32_To_UnsignedInt16 (a[415:384])
//		tmp_dst[415:400] := Saturate_Int32_To_UnsignedInt16 (a[447:416])
//		tmp_dst[431:416] := Saturate_Int32_To_UnsignedInt16 (a[479:448])
//		tmp_dst[447:432] := Saturate_Int32_To_UnsignedInt16 (a[511:480])
//		tmp_dst[463:448] := Saturate_Int32_To_UnsignedInt16 (b[415:384])
//		tmp_dst[479:464] := Saturate_Int32_To_UnsignedInt16 (b[447:416])
//		tmp_dst[495:480] := Saturate_Int32_To_UnsignedInt16 (b[479:448])
//		tmp_dst[511:496] := Saturate_Int32_To_UnsignedInt16 (b[511:480])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm512_maskz_packus_epi32'.
// Requires AVX512BW.
func M512MaskzPackusEpi32(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512PackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
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
//		dst[271:256] := Saturate_Int32_To_UnsignedInt16 (a[287:256])
//		dst[287:272] := Saturate_Int32_To_UnsignedInt16 (a[319:288])
//		dst[303:288] := Saturate_Int32_To_UnsignedInt16 (a[351:320])
//		dst[319:304] := Saturate_Int32_To_UnsignedInt16 (a[383:352])
//		dst[335:320] := Saturate_Int32_To_UnsignedInt16 (b[287:256])
//		dst[351:336] := Saturate_Int32_To_UnsignedInt16 (b[319:288])
//		dst[367:352] := Saturate_Int32_To_UnsignedInt16 (b[351:320])
//		dst[383:368] := Saturate_Int32_To_UnsignedInt16 (b[383:352])
//		dst[399:384] := Saturate_Int32_To_UnsignedInt16 (a[415:384])
//		dst[415:400] := Saturate_Int32_To_UnsignedInt16 (a[447:416])
//		dst[431:416] := Saturate_Int32_To_UnsignedInt16 (a[479:448])
//		dst[447:432] := Saturate_Int32_To_UnsignedInt16 (a[511:480])
//		dst[463:448] := Saturate_Int32_To_UnsignedInt16 (b[415:384])
//		dst[479:464] := Saturate_Int32_To_UnsignedInt16 (b[447:416])
//		dst[495:480] := Saturate_Int32_To_UnsignedInt16 (b[479:448])
//		dst[511:496] := Saturate_Int32_To_UnsignedInt16 (b[511:480])
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKUSDW'. Intrinsic: '_mm512_packus_epi32'.
// Requires AVX512BW.
func M512PackusEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskPermutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+2:i]
//				dst[i+15:i] := idx[i+3] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMT2W'. Intrinsic: '_mm_mask_permutex2var_epi16'.
// Requires AVX512BW.
func MaskPermutex2varEpi16(a x86.M128i, k x86.Mmask8, idx x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// Mask2Permutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'idx' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+2:i]
//				dst[i+15:i] := idx[i+3] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := idx[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMI2W'. Intrinsic: '_mm_mask2_permutex2var_epi16'.
// Requires AVX512BW.
func Mask2Permutex2varEpi16(a x86.M128i, idx x86.M128i, k x86.Mmask8, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPermutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+2:i]
//				dst[i+15:i] := idx[i+3] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMI2W, VPERMT2W'. Intrinsic: '_mm_maskz_permutex2var_epi16'.
// Requires AVX512BW.
func MaskzPermutex2varEpi16(k x86.Mmask8, a x86.M128i, idx x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// Permutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' using the
// corresponding selector and index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			off := 16*idx[i+2:i]
//			dst[i+15:i] := idx[i+4] ? b[off+15:off] : a[off+15:off]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMI2W, VPERMT2W'. Intrinsic: '_mm_permutex2var_epi16'.
// Requires AVX512BW.
func Permutex2varEpi16(a x86.M128i, idx x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPermutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+3:i]
//				dst[i+15:i] := idx[i+4] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMT2W'. Intrinsic: '_mm256_mask_permutex2var_epi16'.
// Requires AVX512BW.
func M256MaskPermutex2varEpi16(a x86.M256i, k x86.Mmask16, idx x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Mask2Permutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'idx' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+3:i]
//				dst[i+15:i] := idx[i+4] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := idx[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMI2W'. Intrinsic: '_mm256_mask2_permutex2var_epi16'.
// Requires AVX512BW.
func M256Mask2Permutex2varEpi16(a x86.M256i, idx x86.M256i, k x86.Mmask16, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPermutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+3:i]
//				dst[i+15:i] := idx[i+4] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMI2W, VPERMT2W'. Intrinsic: '_mm256_maskz_permutex2var_epi16'.
// Requires AVX512BW.
func M256MaskzPermutex2varEpi16(k x86.Mmask16, a x86.M256i, idx x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256Permutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			off := 16*idx[i+3:i]
//			dst[i+15:i] := idx[i+4] ? b[off+15:off] : a[off+15:off]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMI2W, VPERMT2W'. Intrinsic: '_mm256_permutex2var_epi16'.
// Requires AVX512BW.
func M256Permutex2varEpi16(a x86.M256i, idx x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskPermutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+4:i]
//				dst[i+15:i] := idx[i+5] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMT2W'. Intrinsic: '_mm512_mask_permutex2var_epi16'.
// Requires AVX512BW.
func M512MaskPermutex2varEpi16(a x86.M512i, k x86.Mmask32, idx x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512Mask2Permutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'idx' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+4:i]
//				dst[i+15:i] := idx[i+5] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := idx[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMI2W'. Intrinsic: '_mm512_mask2_permutex2var_epi16'.
// Requires AVX512BW.
func M512Mask2Permutex2varEpi16(a x86.M512i, idx x86.M512i, k x86.Mmask32, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzPermutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across
// lanes using the corresponding selector and index in 'idx', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				off := 16*idx[i+4:i]
//				dst[i+15:i] := idx[i+5] ? b[off+15:off] : a[off+15:off]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMI2W, VPERMT2W'. Intrinsic: '_mm512_maskz_permutex2var_epi16'.
// Requires AVX512BW.
func M512MaskzPermutex2varEpi16(k x86.Mmask32, a x86.M512i, idx x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512Permutex2varEpi16: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			off := 16*idx[i+4:i]
//			dst[i+15:i] := idx[i+5] ? b[off+15:off] : a[off+15:off]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMI2W, VPERMT2W'. Intrinsic: '_mm512_permutex2var_epi16'.
// Requires AVX512BW.
func M512Permutex2varEpi16(a x86.M512i, idx x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskPermutexvarEpi16: Shuffle 16-bit integers in 'a' using the corresponding
// index in 'idx', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			id := idx[i+2:i]*16
//			IF k[j]
//				dst[i+15:i] := a[id+15:id]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm_mask_permutexvar_epi16'.
// Requires AVX512BW.
func MaskPermutexvarEpi16(src x86.M128i, k x86.Mmask8, idx x86.M128i, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzPermutexvarEpi16: Shuffle 16-bit integers in 'a' using the
// corresponding index in 'idx', and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			id := idx[i+2:i]*16
//			IF k[j]
//				dst[i+15:i] := a[id+15:id]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm_maskz_permutexvar_epi16'.
// Requires AVX512BW.
func MaskzPermutexvarEpi16(k x86.Mmask8, idx x86.M128i, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// PermutexvarEpi16: Shuffle 16-bit integers in 'a' using the corresponding
// index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			id := idx[i+2:i]*16
//			dst[i+15:i] := a[id+15:id]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm_permutexvar_epi16'.
// Requires AVX512BW.
func PermutexvarEpi16(idx x86.M128i, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskPermutexvarEpi16: Shuffle 16-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			id := idx[i+3:i]*16
//			IF k[j]
//				dst[i+15:i] := a[id+15:id]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm256_mask_permutexvar_epi16'.
// Requires AVX512BW.
func M256MaskPermutexvarEpi16(src x86.M256i, k x86.Mmask16, idx x86.M256i, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzPermutexvarEpi16: Shuffle 16-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			id := idx[i+3:i]*16
//			IF k[j]
//				dst[i+15:i] := a[id+15:id]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm256_maskz_permutexvar_epi16'.
// Requires AVX512BW.
func M256MaskzPermutexvarEpi16(k x86.Mmask16, idx x86.M256i, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256PermutexvarEpi16: Shuffle 16-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			id := idx[i+3:i]*16
//			dst[i+15:i] := a[id+15:id]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm256_permutexvar_epi16'.
// Requires AVX512BW.
func M256PermutexvarEpi16(idx x86.M256i, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskPermutexvarEpi16: Shuffle 16-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			id := idx[i+4:i]*16
//			IF k[j]
//				dst[i+15:i] := a[id+15:id]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm512_mask_permutexvar_epi16'.
// Requires AVX512BW.
func M512MaskPermutexvarEpi16(src x86.M512i, k x86.Mmask32, idx x86.M512i, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzPermutexvarEpi16: Shuffle 16-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			id := idx[i+4:i]*16
//			IF k[j]
//				dst[i+15:i] := a[id+15:id]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm512_maskz_permutexvar_epi16'.
// Requires AVX512BW.
func M512MaskzPermutexvarEpi16(k x86.Mmask32, idx x86.M512i, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512PermutexvarEpi16: Shuffle 16-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			id := idx[i+4:i]*16
//			dst[i+15:i] := a[id+15:id]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMW'. Intrinsic: '_mm512_permutexvar_epi16'.
// Requires AVX512BW.
func M512PermutexvarEpi16(idx x86.M512i, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SadEpu8: Compute the absolute differences of packed unsigned 8-bit
// integers in 'a' and 'b', then horizontally sum each consecutive 8
// differences to produce four unsigned 16-bit integers, and pack these
// unsigned 16-bit integers in the low 16 bits of 64-bit elements in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+15:i] := tmp[i+7:i] + tmp[i+15:i+8] + tmp[i+23:i+16] + tmp[i+31:i+24] + tmp[i+39:i+32] + tmp[i+47:i+40] + tmp[i+55:i+48] + tmp[i+63:i+56]
//			dst[i+63:i+16] := 0
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSADBW'. Intrinsic: '_mm512_sad_epu8'.
// Requires AVX512BW.
func M512SadEpu8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSet1Epi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm_mask_set1_epi16'.
// Requires AVX512BW.
func MaskSet1Epi16(src x86.M128i, k x86.Mmask8, a int16) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSet1Epi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm_maskz_set1_epi16'.
// Requires AVX512BW.
func MaskzSet1Epi16(k x86.Mmask8, a int16) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSet1Epi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm256_mask_set1_epi16'.
// Requires AVX512BW.
func M256MaskSet1Epi16(src x86.M256i, k x86.Mmask16, a int16) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSet1Epi16: Broadcast 16-bit integer 'a' to all elements of 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm256_maskz_set1_epi16'.
// Requires AVX512BW.
func M256MaskzSet1Epi16(k x86.Mmask16, a int16) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSet1Epi16: Broadcast 16-bit integer 'a' to all elements of 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm512_mask_set1_epi16'.
// Requires AVX512BW.
func M512MaskSet1Epi16(src x86.M512i, k x86.Mmask32, a int16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSet1Epi16: Broadcast the low packed 16-bit integer from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[15:0]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm512_maskz_set1_epi16'.
// Requires AVX512BW.
func M512MaskzSet1Epi16(k x86.Mmask32, a int16) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSet1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm_mask_set1_epi8'.
// Requires AVX512BW.
func MaskSet1Epi8(src x86.M128i, k x86.Mmask16, a byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSet1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm_maskz_set1_epi8'.
// Requires AVX512BW.
func MaskzSet1Epi8(k x86.Mmask16, a byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSet1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm256_mask_set1_epi8'.
// Requires AVX512BW.
func M256MaskSet1Epi8(src x86.M256i, k x86.Mmask32, a byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSet1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm256_maskz_set1_epi8'.
// Requires AVX512BW.
func M256MaskzSet1Epi8(k x86.Mmask32, a byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSet1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm512_mask_set1_epi8'.
// Requires AVX512BW.
func M512MaskSet1Epi8(src x86.M512i, k x86.Mmask64, a byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSet1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[7:0]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTB'. Intrinsic: '_mm512_maskz_set1_epi8'.
// Requires AVX512BW.
func M512MaskzSet1Epi8(k x86.Mmask64, a byte) (dst x86.M512i) {
	panic("not implemented")
}


// MaskShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF b[i+7] == 1
//					dst[i+7:i] := 0
//				ELSE
//					index[3:0] := b[i+3:i]
//					dst[i+7:i] := a[index*8+7:index*8]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm_mask_shuffle_epi8'.
// Requires AVX512BW.
func MaskShuffleEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				IF b[i+7] == 1
//					dst[i+7:i] := 0
//				ELSE
//					index[3:0] := b[i+3:i]
//					dst[i+7:i] := a[index*8+7:index*8]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm_maskz_shuffle_epi8'.
// Requires AVX512BW.
func MaskzShuffleEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to
// shuffle control mask in the corresponding 8-bit element of 'b', and store
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF b[i+7] == 1
//					dst[i+7:i] := 0
//				ELSE
//					index[3:0] := b[i+3:i]
//					dst[i+7:i] := a[index*8+7:index*8]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm256_mask_shuffle_epi8'.
// Requires AVX512BW.
func M256MaskShuffleEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to
// shuffle control mask in the corresponding 8-bit element of 'b', and store
// the results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				IF b[i+7] == 1
//					dst[i+7:i] := 0
//				ELSE
//					index[3:0] := b[i+3:i]
//					dst[i+7:i] := a[index*8+7:index*8]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm256_maskz_shuffle_epi8'.
// Requires AVX512BW.
func M256MaskzShuffleEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskShuffleEpi8: Shuffle 8-bit integers in 'a' within 128-bit lanes
// using the control in the corresponding 8-bit element of 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF b[i+7] == 1
//					dst[i+7:i] := 0
//				ELSE
//					index[3:0] := b[i+3:i]
//					dst[i+7:i] := a[index*8+7:index*8]
//				FI
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm512_mask_shuffle_epi8'.
// Requires AVX512BW.
func M512MaskShuffleEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to
// shuffle control mask in the corresponding 8-bit element of 'b', and store
// the results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				IF b[i+7] == 1
//					dst[i+7:i] := 0
//				ELSE
//					index[3:0] := b[i+3:i]
//					dst[i+7:i] := a[index*8+7:index*8]
//				FI
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm512_maskz_shuffle_epi8'.
// Requires AVX512BW.
func M512MaskzShuffleEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512ShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF b[i+7] == 1
//				dst[i+7:i] := 0
//			ELSE
//				index[3:0] := b[i+3:i]
//				dst[i+7:i] := a[index*8+7:index*8]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm512_shuffle_epi8'.
// Requires AVX512BW.
func M512ShuffleEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of 'a' using
// the control in 'imm8'. Store the results in the high 64 bits of 'dst', with
// the low 64 bits being copied from from 'a' to 'dst', using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[63:0] := a[63:0]
//		tmp_dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		tmp_dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		tmp_dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		tmp_dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm_mask_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskShufflehiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of 'a'
// using the control in 'imm8'. Store the results in the high 64 bits of 'dst',
// with the low 64 bits being copied from from 'a' to 'dst', using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		tmp_dst[63:0] := a[63:0]
//		tmp_dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		tmp_dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		tmp_dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		tmp_dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm_maskz_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskzShufflehiEpi16(k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// high 64 bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[63:0] := a[63:0]
//		tmp_dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		tmp_dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		tmp_dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		tmp_dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		tmp_dst[191:128] := a[191:128]
//		tmp_dst[207:192] := (a >> (imm8[1:0] * 16))[207:192]
//		tmp_dst[223:208] := (a >> (imm8[3:2] * 16))[207:192]
//		tmp_dst[239:224] := (a >> (imm8[5:4] * 16))[207:192]
//		tmp_dst[255:240] := (a >> (imm8[7:6] * 16))[207:192]
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm256_mask_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskShufflehiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// high 64 bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		tmp_dst[63:0] := a[63:0]
//		tmp_dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		tmp_dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		tmp_dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		tmp_dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		tmp_dst[191:128] := a[191:128]
//		tmp_dst[207:192] := (a >> (imm8[1:0] * 16))[207:192]
//		tmp_dst[223:208] := (a >> (imm8[3:2] * 16))[207:192]
//		tmp_dst[239:224] := (a >> (imm8[5:4] * 16))[207:192]
//		tmp_dst[255:240] := (a >> (imm8[7:6] * 16))[207:192]
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm256_maskz_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzShufflehiEpi16(k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// high 64 bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[63:0] := a[63:0]
//		tmp_dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		tmp_dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		tmp_dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		tmp_dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		tmp_dst[191:128] := a[191:128]
//		tmp_dst[207:192] := (a >> (imm8[1:0] * 16))[207:192]
//		tmp_dst[223:208] := (a >> (imm8[3:2] * 16))[207:192]
//		tmp_dst[239:224] := (a >> (imm8[5:4] * 16))[207:192]
//		tmp_dst[255:240] := (a >> (imm8[7:6] * 16))[207:192]
//		tmp_dst[319:256] := a[319:256]
//		tmp_dst[335:320] := (a >> (imm8[1:0] * 16))[335:320]
//		tmp_dst[351:336] := (a >> (imm8[3:2] * 16))[335:320]
//		tmp_dst[367:352] := (a >> (imm8[5:4] * 16))[335:320]
//		tmp_dst[383:368] := (a >> (imm8[7:6] * 16))[335:320]
//		tmp_dst[447:384] := a[447:384]
//		tmp_dst[463:448] := (a >> (imm8[1:0] * 16))[463:448]
//		tmp_dst[479:464] := (a >> (imm8[3:2] * 16))[463:448]
//		tmp_dst[495:480] := (a >> (imm8[5:4] * 16))[463:448]
//		tmp_dst[511:496] := (a >> (imm8[7:6] * 16))[463:448]
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm512_mask_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskShufflehiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// high 64 bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		tmp_dst[63:0] := a[63:0]
//		tmp_dst[79:64] := (a >> (imm8[1:0] * 16))[79:64]
//		tmp_dst[95:80] := (a >> (imm8[3:2] * 16))[79:64]
//		tmp_dst[111:96] := (a >> (imm8[5:4] * 16))[79:64]
//		tmp_dst[127:112] := (a >> (imm8[7:6] * 16))[79:64]
//		tmp_dst[191:128] := a[191:128]
//		tmp_dst[207:192] := (a >> (imm8[1:0] * 16))[207:192]
//		tmp_dst[223:208] := (a >> (imm8[3:2] * 16))[207:192]
//		tmp_dst[239:224] := (a >> (imm8[5:4] * 16))[207:192]
//		tmp_dst[255:240] := (a >> (imm8[7:6] * 16))[207:192]
//		tmp_dst[319:256] := a[319:256]
//		tmp_dst[335:320] := (a >> (imm8[1:0] * 16))[335:320]
//		tmp_dst[351:336] := (a >> (imm8[3:2] * 16))[335:320]
//		tmp_dst[367:352] := (a >> (imm8[5:4] * 16))[335:320]
//		tmp_dst[383:368] := (a >> (imm8[7:6] * 16))[335:320]
//		tmp_dst[447:384] := a[447:384]
//		tmp_dst[463:448] := (a >> (imm8[1:0] * 16))[463:448]
//		tmp_dst[479:464] := (a >> (imm8[3:2] * 16))[463:448]
//		tmp_dst[495:480] := (a >> (imm8[5:4] * 16))[463:448]
//		tmp_dst[511:496] := (a >> (imm8[7:6] * 16))[463:448]
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm512_maskz_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzShufflehiEpi16(k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512ShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of 128-bit
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
//		dst[319:256] := a[319:256]
//		dst[335:320] := (a >> (imm8[1:0] * 16))[335:320]
//		dst[351:336] := (a >> (imm8[3:2] * 16))[335:320]
//		dst[367:352] := (a >> (imm8[5:4] * 16))[335:320]
//		dst[383:368] := (a >> (imm8[7:6] * 16))[335:320]
//		dst[447:384] := a[447:384]
//		dst[463:448] := (a >> (imm8[1:0] * 16))[463:448]
//		dst[479:464] := (a >> (imm8[3:2] * 16))[463:448]
//		dst[495:480] := (a >> (imm8[5:4] * 16))[463:448]
//		dst[511:496] := (a >> (imm8[7:6] * 16))[463:448]
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFHW'. Intrinsic: '_mm512_shufflehi_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512ShufflehiEpi16(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// MaskShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of 'a' using
// the control in 'imm8'. Store the results in the low 64 bits of 'dst', with
// the high 64 bits being copied from from 'a' to 'dst', using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		tmp_dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		tmp_dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		tmp_dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		tmp_dst[127:64] := a[127:64]
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm_mask_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskShuffleloEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of 'a' using
// the control in 'imm8'. Store the results in the low 64 bits of 'dst', with
// the high 64 bits being copied from from 'a' to 'dst', using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		tmp_dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		tmp_dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		tmp_dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		tmp_dst[127:64] := a[127:64]
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm_maskz_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskzShuffleloEpi16(k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// low 64 bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		tmp_dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		tmp_dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		tmp_dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		tmp_dst[127:64] := a[127:64]
//		tmp_dst[143:128] := (a >> (imm8[1:0] * 16))[143:128]
//		tmp_dst[159:144] := (a >> (imm8[3:2] * 16))[143:128]
//		tmp_dst[175:160] := (a >> (imm8[5:4] * 16))[143:128]
//		tmp_dst[191:176] := (a >> (imm8[7:6] * 16))[143:128]
//		tmp_dst[255:192] := a[255:192]
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm256_mask_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskShuffleloEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// low 64 bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		tmp_dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		tmp_dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		tmp_dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		tmp_dst[127:64] := a[127:64]
//		tmp_dst[143:128] := (a >> (imm8[1:0] * 16))[143:128]
//		tmp_dst[159:144] := (a >> (imm8[3:2] * 16))[143:128]
//		tmp_dst[175:160] := (a >> (imm8[5:4] * 16))[143:128]
//		tmp_dst[191:176] := (a >> (imm8[7:6] * 16))[143:128]
//		tmp_dst[255:192] := a[255:192]
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm256_maskz_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzShuffleloEpi16(k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// low 64 bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		tmp_dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		tmp_dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		tmp_dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		tmp_dst[127:64] := a[127:64]
//		tmp_dst[143:128] := (a >> (imm8[1:0] * 16))[143:128]
//		tmp_dst[159:144] := (a >> (imm8[3:2] * 16))[143:128]
//		tmp_dst[175:160] := (a >> (imm8[5:4] * 16))[143:128]
//		tmp_dst[191:176] := (a >> (imm8[7:6] * 16))[143:128]
//		tmp_dst[255:192] := a[255:192]
//		tmp_dst[271:256] := (a >> (imm8[1:0] * 16))[271:256]
//		tmp_dst[287:272] := (a >> (imm8[3:2] * 16))[271:256]
//		tmp_dst[303:288] := (a >> (imm8[5:4] * 16))[271:256]
//		tmp_dst[319:304] := (a >> (imm8[7:6] * 16))[271:256]
//		tmp_dst[383:320] := a[383:320]
//		tmp_dst[399:384] := (a >> (imm8[1:0] * 16))[399:384]
//		tmp_dst[415:400] := (a >> (imm8[3:2] * 16))[399:384]
//		tmp_dst[431:416] := (a >> (imm8[5:4] * 16))[399:384]
//		tmp_dst[447:432] := (a >> (imm8[7:6] * 16))[399:384]
//		tmp_dst[511:448] := a[511:448]
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm512_mask_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskShuffleloEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of
// 128-bit lanes of 'a' using the control in 'imm8'. Store the results in the
// low 64 bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit
// lanes being copied from from 'a' to 'dst', using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		tmp_dst[15:0] := (a >> (imm8[1:0] * 16))[15:0]
//		tmp_dst[31:16] := (a >> (imm8[3:2] * 16))[15:0]
//		tmp_dst[47:32] := (a >> (imm8[5:4] * 16))[15:0]
//		tmp_dst[63:48] := (a >> (imm8[7:6] * 16))[15:0]
//		tmp_dst[127:64] := a[127:64]
//		tmp_dst[143:128] := (a >> (imm8[1:0] * 16))[143:128]
//		tmp_dst[159:144] := (a >> (imm8[3:2] * 16))[143:128]
//		tmp_dst[175:160] := (a >> (imm8[5:4] * 16))[143:128]
//		tmp_dst[191:176] := (a >> (imm8[7:6] * 16))[143:128]
//		tmp_dst[255:192] := a[255:192]
//		tmp_dst[271:256] := (a >> (imm8[1:0] * 16))[271:256]
//		tmp_dst[287:272] := (a >> (imm8[3:2] * 16))[271:256]
//		tmp_dst[303:288] := (a >> (imm8[5:4] * 16))[271:256]
//		tmp_dst[319:304] := (a >> (imm8[7:6] * 16))[271:256]
//		tmp_dst[383:320] := a[383:320]
//		tmp_dst[399:384] := (a >> (imm8[1:0] * 16))[399:384]
//		tmp_dst[415:400] := (a >> (imm8[3:2] * 16))[399:384]
//		tmp_dst[431:416] := (a >> (imm8[5:4] * 16))[399:384]
//		tmp_dst[447:432] := (a >> (imm8[7:6] * 16))[399:384]
//		tmp_dst[511:448] := a[511:448]
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm512_maskz_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzShuffleloEpi16(k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512ShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of 128-bit
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
//		dst[271:256] := (a >> (imm8[1:0] * 16))[271:256]
//		dst[287:272] := (a >> (imm8[3:2] * 16))[271:256]
//		dst[303:288] := (a >> (imm8[5:4] * 16))[271:256]
//		dst[319:304] := (a >> (imm8[7:6] * 16))[271:256]
//		dst[383:320] := a[383:320]
//		dst[399:384] := (a >> (imm8[1:0] * 16))[399:384]
//		dst[415:400] := (a >> (imm8[3:2] * 16))[399:384]
//		dst[431:416] := (a >> (imm8[5:4] * 16))[399:384]
//		dst[447:432] := (a >> (imm8[7:6] * 16))[399:384]
//		dst[511:448] := a[511:448]
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFLW'. Intrinsic: '_mm512_shufflelo_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512ShuffleloEpi16(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm_mask_sll_epi16'.
// Requires AVX512BW.
func MaskSllEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm_maskz_sll_epi16'.
// Requires AVX512BW.
func MaskzSllEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm256_mask_sll_epi16'.
// Requires AVX512BW.
func M256MaskSllEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm256_maskz_sll_epi16'.
// Requires AVX512BW.
func M256MaskzSllEpi16(k x86.Mmask16, a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm512_mask_sll_epi16'.
// Requires AVX512BW.
func M512MaskSllEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm512_maskz_sll_epi16'.
// Requires AVX512BW.
func M512MaskzSllEpi16(k x86.Mmask32, a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm512_sll_epi16'.
// Requires AVX512BW.
func M512SllEpi16(a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm_mask_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskSlliEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm_maskz_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskzSlliEpi16(k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm256_mask_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskSlliEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm256_maskz_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzSlliEpi16(k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm512_mask_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskSlliEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm512_maskz_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzSlliEpi16(k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512SlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLW'. Intrinsic: '_mm512_slli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512SlliEpi16(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm_mask_sllv_epi16'.
// Requires AVX512BW.
func MaskSllvEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm_maskz_sllv_epi16'.
// Requires AVX512BW.
func MaskzSllvEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// SllvEpi16: Shift packed 16-bit integers in 'a' left by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm_sllv_epi16'.
// Requires AVX512BW.
func SllvEpi16(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm256_mask_sllv_epi16'.
// Requires AVX512BW.
func M256MaskSllvEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm256_maskz_sllv_epi16'.
// Requires AVX512BW.
func M256MaskzSllvEpi16(k x86.Mmask16, a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm256_sllv_epi16'.
// Requires AVX512BW.
func M256SllvEpi16(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm512_mask_sllv_epi16'.
// Requires AVX512BW.
func M512MaskSllvEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm512_maskz_sllv_epi16'.
// Requires AVX512BW.
func M512MaskzSllvEpi16(k x86.Mmask32, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SllvEpi16: Shift packed 16-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] << count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVW'. Intrinsic: '_mm512_sllv_epi16'.
// Requires AVX512BW.
func M512SllvEpi16(a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm_mask_sra_epi16'.
// Requires AVX512BW.
func MaskSraEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm_maskz_sra_epi16'.
// Requires AVX512BW.
func MaskzSraEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm256_mask_sra_epi16'.
// Requires AVX512BW.
func M256MaskSraEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSraEpi16: Shift packed 16-bit integers in 'a' right by 'count'
// while shifting in sign bits, and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm256_maskz_sra_epi16'.
// Requires AVX512BW.
func M256MaskzSraEpi16(k x86.Mmask16, a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm512_mask_sra_epi16'.
// Requires AVX512BW.
func M512MaskSraEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSraEpi16: Shift packed 16-bit integers in 'a' right by 'count'
// while shifting in sign bits, and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm512_maskz_sra_epi16'.
// Requires AVX512BW.
func M512MaskzSraEpi16(k x86.Mmask32, a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm512_sra_epi16'.
// Requires AVX512BW.
func M512SraEpi16(a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm_mask_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskSraiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm_maskz_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskzSraiEpi16(k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm256_mask_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskSraiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8'
// while shifting in sign bits, and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm256_maskz_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzSraiEpi16(k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm512_mask_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskSraiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8'
// while shifting in sign bits, and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := SignBit
//				ELSE
//					dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm512_maskz_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzSraiEpi16(k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512SraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := SignBit
//			ELSE
//				dst[i+15:i] := SignExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAW'. Intrinsic: '_mm512_srai_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512SraiEpi16(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm_mask_srav_epi16'.
// Requires AVX512BW.
func MaskSravEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm_maskz_srav_epi16'.
// Requires AVX512BW.
func MaskzSravEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// SravEpi16: Shift packed 16-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in sign bits, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm_srav_epi16'.
// Requires AVX512BW.
func SravEpi16(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm256_mask_srav_epi16'.
// Requires AVX512BW.
func M256MaskSravEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm256_maskz_srav_epi16'.
// Requires AVX512BW.
func M256MaskzSravEpi16(k x86.Mmask16, a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm256_srav_epi16'.
// Requires AVX512BW.
func M256SravEpi16(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm512_mask_srav_epi16'.
// Requires AVX512BW.
func M512MaskSravEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm512_maskz_srav_epi16'.
// Requires AVX512BW.
func M512MaskzSravEpi16(k x86.Mmask32, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SravEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := SignExtend(a[i+15:i] >> count[i+15:i])	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVW'. Intrinsic: '_mm512_srav_epi16'.
// Requires AVX512BW.
func M512SravEpi16(a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm_mask_srl_epi16'.
// Requires AVX512BW.
func MaskSrlEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm_maskz_srl_epi16'.
// Requires AVX512BW.
func MaskzSrlEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm256_mask_srl_epi16'.
// Requires AVX512BW.
func M256MaskSrlEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSrlEpi16: Shift packed 16-bit integers in 'a' right by 'count'
// while shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm256_maskz_srl_epi16'.
// Requires AVX512BW.
func M256MaskzSrlEpi16(k x86.Mmask16, a x86.M256i, count x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm512_mask_srl_epi16'.
// Requires AVX512BW.
func M512MaskSrlEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSrlEpi16: Shift packed 16-bit integers in 'a' right by 'count'
// while shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF count[63:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm512_maskz_srl_epi16'.
// Requires AVX512BW.
func M512MaskzSrlEpi16(k x86.Mmask32, a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF count[63:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[63:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm512_srl_epi16'.
// Requires AVX512BW.
func M512SrlEpi16(a x86.M512i, count x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm_mask_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskSrliEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm_maskz_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func MaskzSrliEpi16(k x86.Mmask8, a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm256_mask_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskSrliEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8'
// while shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm256_maskz_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzSrliEpi16(k x86.Mmask16, a x86.M256i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm512_mask_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskSrliEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8'
// while shifting in zeros, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				IF imm8[7:0] > 15
//					dst[i+15:i] := 0
//				ELSE
//					dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm512_maskz_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzSrliEpi16(k x86.Mmask32, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512SrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF imm8[7:0] > 15
//				dst[i+15:i] := 0
//			ELSE
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLW'. Intrinsic: '_mm512_srli_epi16'.
// Requires AVX512BW.
//
// FIXME: Requires compiler support (has immediate)
func M512SrliEpi16(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm_mask_srlv_epi16'.
// Requires AVX512BW.
func MaskSrlvEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm_maskz_srlv_epi16'.
// Requires AVX512BW.
func MaskzSrlvEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// SrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm_srlv_epi16'.
// Requires AVX512BW.
func SrlvEpi16(a x86.M128i, count x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+63:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm256_mask_srlv_epi16'.
// Requires AVX512BW.
func M256MaskSrlvEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm256_maskz_srlv_epi16'.
// Requires AVX512BW.
func M256MaskzSrlvEpi16(k x86.Mmask16, a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256SrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm256_srlv_epi16'.
// Requires AVX512BW.
func M256SrlvEpi16(a x86.M256i, count x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm512_mask_srlv_epi16'.
// Requires AVX512BW.
func M512MaskSrlvEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm512_maskz_srlv_epi16'.
// Requires AVX512BW.
func M512MaskzSrlvEpi16(k x86.Mmask32, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SrlvEpi16: Shift packed 16-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := ZeroExtend(a[i+15:i] >> count[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVW'. Intrinsic: '_mm512_srlv_epi16'.
// Requires AVX512BW.
func M512SrlvEpi16(a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// Skipped: _mm_mask_storeu_epi16. Contains pointer parameter.


// Skipped: _mm256_mask_storeu_epi16. Contains pointer parameter.


// Skipped: _mm512_mask_storeu_epi16. Contains pointer parameter.


// Skipped: _mm_mask_storeu_epi8. Contains pointer parameter.


// Skipped: _mm256_mask_storeu_epi8. Contains pointer parameter.


// Skipped: _mm512_mask_storeu_epi8. Contains pointer parameter.


// MaskSubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] - b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm_mask_sub_epi16'.
// Requires AVX512BW.
func MaskSubEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] - b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm_maskz_sub_epi16'.
// Requires AVX512BW.
func MaskzSubEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] - b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm256_mask_sub_epi16'.
// Requires AVX512BW.
func M256MaskSubEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] - b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm256_maskz_sub_epi16'.
// Requires AVX512BW.
func M256MaskzSubEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] - b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm512_mask_sub_epi16'.
// Requires AVX512BW.
func M512MaskSubEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := a[i+15:i] - b[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm512_maskz_sub_epi16'.
// Requires AVX512BW.
func M512MaskzSubEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := a[i+15:i] - b[i+15:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm512_sub_epi16'.
// Requires AVX512BW.
func M512SubEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] - b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm_mask_sub_epi8'.
// Requires AVX512BW.
func MaskSubEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] - b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm_maskz_sub_epi8'.
// Requires AVX512BW.
func MaskzSubEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] - b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm256_mask_sub_epi8'.
// Requires AVX512BW.
func M256MaskSubEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] - b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm256_maskz_sub_epi8'.
// Requires AVX512BW.
func M256MaskzSubEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] - b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm512_mask_sub_epi8'.
// Requires AVX512BW.
func M512MaskSubEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := a[i+7:i] - b[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm512_maskz_sub_epi8'.
// Requires AVX512BW.
func M512MaskzSubEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := a[i+7:i] - b[i+7:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm512_sub_epi8'.
// Requires AVX512BW.
func M512SubEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm_mask_subs_epi16'.
// Requires AVX512BW.
func MaskSubsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm_maskz_subs_epi16'.
// Requires AVX512BW.
func MaskzSubsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm256_mask_subs_epi16'.
// Requires AVX512BW.
func M256MaskSubsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSubsEpi16: Subtract packed 16-bit integers in 'b' from packed
// 16-bit integers in 'a' using saturation, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm256_maskz_subs_epi16'.
// Requires AVX512BW.
func M256MaskzSubsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm512_mask_subs_epi16'.
// Requires AVX512BW.
func M512MaskSubsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSubsEpi16: Subtract packed 16-bit integers in 'b' from packed
// 16-bit integers in 'a' using saturation, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm512_maskz_subs_epi16'.
// Requires AVX512BW.
func M512MaskzSubsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16(a[i+15:i] - b[i+15:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSW'. Intrinsic: '_mm512_subs_epi16'.
// Requires AVX512BW.
func M512SubsEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm_mask_subs_epi8'.
// Requires AVX512BW.
func MaskSubsEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm_maskz_subs_epi8'.
// Requires AVX512BW.
func MaskzSubsEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm256_mask_subs_epi8'.
// Requires AVX512BW.
func M256MaskSubsEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm256_maskz_subs_epi8'.
// Requires AVX512BW.
func M256MaskzSubsEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm512_mask_subs_epi8'.
// Requires AVX512BW.
func M512MaskSubsEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm512_maskz_subs_epi8'.
// Requires AVX512BW.
func M512MaskzSubsEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit
// integers in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm512_subs_epi8'.
// Requires AVX512BW.
func M512SubsEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm_mask_subs_epu16'.
// Requires AVX512BW.
func MaskSubsEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm_maskz_subs_epu16'.
// Requires AVX512BW.
func MaskzSubsEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from
// packed unsigned 16-bit integers in 'a' using saturation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm256_mask_subs_epu16'.
// Requires AVX512BW.
func M256MaskSubsEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from
// packed unsigned 16-bit integers in 'a' using saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm256_maskz_subs_epu16'.
// Requires AVX512BW.
func M256MaskzSubsEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from
// packed unsigned 16-bit integers in 'a' using saturation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm512_mask_subs_epu16'.
// Requires AVX512BW.
func M512MaskSubsEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from
// packed unsigned 16-bit integers in 'a' using saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])
//			ELSE
//				dst[i+15:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm512_maskz_subs_epu16'.
// Requires AVX512BW.
func M512MaskzSubsEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := Saturate_To_UnsignedInt16(a[i+15:i] - b[i+15:i])	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBUSW'. Intrinsic: '_mm512_subs_epu16'.
// Requires AVX512BW.
func M512SubsEpu16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskSubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm_mask_subs_epu8'.
// Requires AVX512BW.
func MaskSubsEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzSubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm_maskz_subs_epu8'.
// Requires AVX512BW.
func MaskzSubsEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskSubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm256_mask_subs_epu8'.
// Requires AVX512BW.
func M256MaskSubsEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzSubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from
// packed unsigned 8-bit integers in 'a' using saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm256_maskz_subs_epu8'.
// Requires AVX512BW.
func M256MaskzSubsEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskSubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm512_mask_subs_epu8'.
// Requires AVX512BW.
func M512MaskSubsEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzSubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from
// packed unsigned 8-bit integers in 'a' using saturation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])
//			ELSE
//				dst[i+7:i] := 0
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm512_maskz_subs_epu8'.
// Requires AVX512BW.
func M512MaskzSubsEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := Saturate_To_UnsignedInt8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBUSB'. Intrinsic: '_mm512_subs_epu8'.
// Requires AVX512BW.
func M512SubsEpu8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskTestEpi16Mask: Compute the bitwise AND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm_mask_test_epi16_mask'.
// Requires AVX512BW.
func MaskTestEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// TestEpi16Mask: Compute the bitwise AND of packed 16-bit integers in 'a' and
// 'b', producing intermediate 16-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm_test_epi16_mask'.
// Requires AVX512BW.
func TestEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256MaskTestEpi16Mask: Compute the bitwise AND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm256_mask_test_epi16_mask'.
// Requires AVX512BW.
func M256MaskTestEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256TestEpi16Mask: Compute the bitwise AND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm256_test_epi16_mask'.
// Requires AVX512BW.
func M256TestEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskTestEpi16Mask: Compute the bitwise AND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is non-zero. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm512_mask_test_epi16_mask'.
// Requires AVX512BW.
func M512MaskTestEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512TestEpi16Mask: Compute the bitwise AND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm512_test_epi16_mask'.
// Requires AVX512BW.
func M512TestEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// MaskTestEpi8Mask: Compute the bitwise AND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm_mask_test_epi8_mask'.
// Requires AVX512BW.
func MaskTestEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// TestEpi8Mask: Compute the bitwise AND of packed 8-bit integers in 'a' and
// 'b', producing intermediate 8-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm_test_epi8_mask'.
// Requires AVX512BW.
func TestEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskTestEpi8Mask: Compute the bitwise AND of packed 8-bit integers in
// 'a' and 'b', producing intermediate 8-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is non-zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm256_mask_test_epi8_mask'.
// Requires AVX512BW.
func M256MaskTestEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256TestEpi8Mask: Compute the bitwise AND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm256_test_epi8_mask'.
// Requires AVX512BW.
func M256TestEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskTestEpi8Mask: Compute the bitwise AND of packed 8-bit integers in
// 'a' and 'b', producing intermediate 8-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is non-zero. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm512_mask_test_epi8_mask'.
// Requires AVX512BW.
func M512MaskTestEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512TestEpi8Mask: Compute the bitwise AND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm512_test_epi8_mask'.
// Requires AVX512BW.
func M512TestEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// MaskTestnEpi16Mask: Compute the bitwise NAND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is zero. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k1[j]
//				k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm_mask_testn_epi16_mask'.
// Requires AVX512BW.
func MaskTestnEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// TestnEpi16Mask: Compute the bitwise NAND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 7
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm_testn_epi16_mask'.
// Requires AVX512BW.
func TestnEpi16Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


// M256MaskTestnEpi16Mask: Compute the bitwise NAND of packed 16-bit integers
// in 'a' and 'b', producing intermediate 16-bit values, and set the
// corresponding bit in result mask 'k' (subject to writemask 'k') if the
// intermediate value is zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k1[j]
//				k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm256_mask_testn_epi16_mask'.
// Requires AVX512BW.
func M256MaskTestnEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256TestnEpi16Mask: Compute the bitwise NAND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm256_testn_epi16_mask'.
// Requires AVX512BW.
func M256TestnEpi16Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskTestnEpi16Mask: Compute the bitwise NAND of packed 16-bit integers
// in 'a' and 'b', producing intermediate 16-bit values, and set the
// corresponding bit in result mask 'k' (subject to writemask 'k') if the
// intermediate value is zero. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k1[j]
//				k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm512_mask_testn_epi16_mask'.
// Requires AVX512BW.
func M512MaskTestnEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512TestnEpi16Mask: Compute the bitwise NAND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm512_testn_epi16_mask'.
// Requires AVX512BW.
func M512TestnEpi16Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask32) {
	panic("not implemented")
}


// MaskTestnEpi8Mask: Compute the bitwise NAND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// zero. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k1[j]
//				k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm_mask_testn_epi8_mask'.
// Requires AVX512BW.
func MaskTestnEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// TestnEpi8Mask: Compute the bitwise NAND of packed 8-bit integers in 'a' and
// 'b', producing intermediate 8-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 15
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm_testn_epi8_mask'.
// Requires AVX512BW.
func TestnEpi8Mask(a x86.M128i, b x86.M128i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M256MaskTestnEpi8Mask: Compute the bitwise NAND of packed 8-bit integers in
// 'a' and 'b', producing intermediate 8-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k1[j]
//				k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm256_mask_testn_epi8_mask'.
// Requires AVX512BW.
func M256MaskTestnEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M256TestnEpi8Mask: Compute the bitwise NAND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm256_testn_epi8_mask'.
// Requires AVX512BW.
func M256TestnEpi8Mask(a x86.M256i, b x86.M256i) (dst x86.Mmask32) {
	panic("not implemented")
}


// M512MaskTestnEpi8Mask: Compute the bitwise NAND of packed 8-bit integers in
// 'a' and 'b', producing intermediate 8-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is zero. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k1[j]
//				k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm512_mask_testn_epi8_mask'.
// Requires AVX512BW.
func M512MaskTestnEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// M512TestnEpi8Mask: Compute the bitwise NAND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm512_testn_epi8_mask'.
// Requires AVX512BW.
func M512TestnEpi8Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask64) {
	panic("not implemented")
}


// MaskUnpackhiEpi16: Unpack and interleave 16-bit integers from the high half
// of 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm_mask_unpackhi_epi16'.
// Requires AVX512BW.
func MaskUnpackhiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzUnpackhiEpi16: Unpack and interleave 16-bit integers from the high half
// of 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm_maskz_unpackhi_epi16'.
// Requires AVX512BW.
func MaskzUnpackhiEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskUnpackhiEpi16: Unpack and interleave 16-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_WORDS(a[255:128], b[255:128])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm256_mask_unpackhi_epi16'.
// Requires AVX512BW.
func M256MaskUnpackhiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzUnpackhiEpi16: Unpack and interleave 16-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_WORDS(a[255:128], b[255:128])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm256_maskz_unpackhi_epi16'.
// Requires AVX512BW.
func M256MaskzUnpackhiEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskUnpackhiEpi16: Unpack and interleave 16-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_WORDS(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_HIGH_WORDS(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_HIGH_WORDS(a[511:384], b[511:384])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm512_mask_unpackhi_epi16'.
// Requires AVX512BW.
func M512MaskUnpackhiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzUnpackhiEpi16: Unpack and interleave 16-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_WORDS(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_HIGH_WORDS(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_HIGH_WORDS(a[511:384], b[511:384])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm512_maskz_unpackhi_epi16'.
// Requires AVX512BW.
func M512MaskzUnpackhiEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512UnpackhiEpi16: Unpack and interleave 16-bit integers from the high half
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
//		dst[383:256] := INTERLEAVE_HIGH_WORDS(a[383:256], b[383:256])
//		dst[511:384] := INTERLEAVE_HIGH_WORDS(a[511:384], b[511:384])
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKHWD'. Intrinsic: '_mm512_unpackhi_epi16'.
// Requires AVX512BW.
func M512UnpackhiEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskUnpackhiEpi8: Unpack and interleave 8-bit integers from the high half of
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm_mask_unpackhi_epi8'.
// Requires AVX512BW.
func MaskUnpackhiEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzUnpackhiEpi8: Unpack and interleave 8-bit integers from the high half
// of 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm_maskz_unpackhi_epi8'.
// Requires AVX512BW.
func MaskzUnpackhiEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskUnpackhiEpi8: Unpack and interleave 8-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_BYTES(a[255:128], b[255:128])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm256_mask_unpackhi_epi8'.
// Requires AVX512BW.
func M256MaskUnpackhiEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzUnpackhiEpi8: Unpack and interleave 8-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_BYTES(a[255:128], b[255:128])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm256_maskz_unpackhi_epi8'.
// Requires AVX512BW.
func M256MaskzUnpackhiEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskUnpackhiEpi8: Unpack and interleave 8-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_BYTES(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_HIGH_BYTES(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_HIGH_BYTES(a[511:384], b[511:384])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm512_mask_unpackhi_epi8'.
// Requires AVX512BW.
func M512MaskUnpackhiEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzUnpackhiEpi8: Unpack and interleave 8-bit integers from the high
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_HIGH_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_HIGH_BYTES(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_HIGH_BYTES(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_HIGH_BYTES(a[511:384], b[511:384])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm512_maskz_unpackhi_epi8'.
// Requires AVX512BW.
func M512MaskzUnpackhiEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512UnpackhiEpi8: Unpack and interleave 8-bit integers from the high half of
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
//		dst[383:256] := INTERLEAVE_HIGH_BYTES(a[383:256], b[383:256])
//		dst[511:384] := INTERLEAVE_HIGH_BYTES(a[511:384], b[511:384])
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKHBW'. Intrinsic: '_mm512_unpackhi_epi8'.
// Requires AVX512BW.
func M512UnpackhiEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskUnpackloEpi16: Unpack and interleave 16-bit integers from the low half
// of 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm_mask_unpacklo_epi16'.
// Requires AVX512BW.
func MaskUnpackloEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzUnpackloEpi16: Unpack and interleave 16-bit integers from the low half
// of 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//		
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm_maskz_unpacklo_epi16'.
// Requires AVX512BW.
func MaskzUnpackloEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskUnpackloEpi16: Unpack and interleave 16-bit integers from the low
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_WORDS(a[255:128], b[255:128])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm256_mask_unpacklo_epi16'.
// Requires AVX512BW.
func M256MaskUnpackloEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzUnpackloEpi16: Unpack and interleave 16-bit integers from the low
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_WORDS(a[255:128], b[255:128])
//		
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm256_maskz_unpacklo_epi16'.
// Requires AVX512BW.
func M256MaskzUnpackloEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskUnpackloEpi16: Unpack and interleave 16-bit integers from the low
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_WORDS(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_WORDS(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_WORDS(a[511:384], b[511:384])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm512_mask_unpacklo_epi16'.
// Requires AVX512BW.
func M512MaskUnpackloEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzUnpackloEpi16: Unpack and interleave 16-bit integers from the low
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_WORDS(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_WORDS(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_WORDS(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_WORDS(a[511:384], b[511:384])
//		
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := tmp_dst[i+15:i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm512_maskz_unpacklo_epi16'.
// Requires AVX512BW.
func M512MaskzUnpackloEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512UnpackloEpi16: Unpack and interleave 16-bit integers from the low half
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
//		dst[383:256] := INTERLEAVE_WORDS(a[383:256], b[383:256])
//		dst[511:384] := INTERLEAVE_WORDS(a[511:384], b[511:384])
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKLWD'. Intrinsic: '_mm512_unpacklo_epi16'.
// Requires AVX512BW.
func M512UnpackloEpi16(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// MaskUnpackloEpi8: Unpack and interleave 8-bit integers from the low half of
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm_mask_unpacklo_epi8'.
// Requires AVX512BW.
func MaskUnpackloEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzUnpackloEpi8: Unpack and interleave 8-bit integers from the low half of
// 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//		
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm_maskz_unpacklo_epi8'.
// Requires AVX512BW.
func MaskzUnpackloEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256MaskUnpackloEpi8: Unpack and interleave 8-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_BYTES(a[255:128], b[255:128])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm256_mask_unpacklo_epi8'.
// Requires AVX512BW.
func M256MaskUnpackloEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzUnpackloEpi8: Unpack and interleave 8-bit integers from the low
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_BYTES(a[255:128], b[255:128])
//		
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm256_maskz_unpacklo_epi8'.
// Requires AVX512BW.
func M256MaskzUnpackloEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512MaskUnpackloEpi8: Unpack and interleave 8-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_BYTES(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_BYTES(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_BYTES(a[511:384], b[511:384])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm512_mask_unpacklo_epi8'.
// Requires AVX512BW.
func M512MaskUnpackloEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzUnpackloEpi8: Unpack and interleave 8-bit integers from the low
// half of each 128-bit lane in 'a' and 'b', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
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
//		tmp_dst[127:0] := INTERLEAVE_BYTES(a[127:0], b[127:0])
//		tmp_dst[255:128] := INTERLEAVE_BYTES(a[255:128], b[255:128])
//		tmp_dst[383:256] := INTERLEAVE_BYTES(a[383:256], b[383:256])
//		tmp_dst[511:384] := INTERLEAVE_BYTES(a[511:384], b[511:384])
//		
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := tmp_dst[i+7:i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm512_maskz_unpacklo_epi8'.
// Requires AVX512BW.
func M512MaskzUnpackloEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512UnpackloEpi8: Unpack and interleave 8-bit integers from the low half of
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
//		dst[383:256] := INTERLEAVE_BYTES(a[383:256], b[383:256])
//		dst[511:384] := INTERLEAVE_BYTES(a[511:384], b[511:384])
//		dst[MAX:512] := 0
//
// Instruction: 'VPUNPCKLBW'. Intrinsic: '_mm512_unpacklo_epi8'.
// Requires AVX512BW.
func M512UnpackloEpi8(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}

