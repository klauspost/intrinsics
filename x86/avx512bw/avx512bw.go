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
func MaskAbsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskAbsEpi16([16]byte(src), uint8(k), [16]byte(a)))
}

func maskAbsEpi16(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzAbsEpi16(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzAbsEpi16(uint8(k), [16]byte(a)))
}

func maskzAbsEpi16(k uint8, a [16]byte) [16]byte


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
func M256MaskAbsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAbsEpi16([32]byte(src), uint16(k), [32]byte(a)))
}

func m256MaskAbsEpi16(src [32]byte, k uint16, a [32]byte) [32]byte


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
func M256MaskzAbsEpi16(k x86.Mmask16, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAbsEpi16(uint16(k), [32]byte(a)))
}

func m256MaskzAbsEpi16(k uint16, a [32]byte) [32]byte


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
func M512AbsEpi16(a x86.M512i) x86.M512i {
	return x86.M512i(m512AbsEpi16([64]byte(a)))
}

func m512AbsEpi16(a [64]byte) [64]byte


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
func M512MaskAbsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAbsEpi16([64]byte(src), uint32(k), [64]byte(a)))
}

func m512MaskAbsEpi16(src [64]byte, k uint32, a [64]byte) [64]byte


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
func M512MaskzAbsEpi16(k x86.Mmask32, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAbsEpi16(uint32(k), [64]byte(a)))
}

func m512MaskzAbsEpi16(k uint32, a [64]byte) [64]byte


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
func MaskAbsEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i) x86.M128i {
	return x86.M128i(maskAbsEpi8([16]byte(src), uint16(k), [16]byte(a)))
}

func maskAbsEpi8(src [16]byte, k uint16, a [16]byte) [16]byte


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
func MaskzAbsEpi8(k x86.Mmask16, a x86.M128i) x86.M128i {
	return x86.M128i(maskzAbsEpi8(uint16(k), [16]byte(a)))
}

func maskzAbsEpi8(k uint16, a [16]byte) [16]byte


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
func M256MaskAbsEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAbsEpi8([32]byte(src), uint32(k), [32]byte(a)))
}

func m256MaskAbsEpi8(src [32]byte, k uint32, a [32]byte) [32]byte


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
func M256MaskzAbsEpi8(k x86.Mmask32, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAbsEpi8(uint32(k), [32]byte(a)))
}

func m256MaskzAbsEpi8(k uint32, a [32]byte) [32]byte


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
func M512AbsEpi8(a x86.M512i) x86.M512i {
	return x86.M512i(m512AbsEpi8([64]byte(a)))
}

func m512AbsEpi8(a [64]byte) [64]byte


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
func M512MaskAbsEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAbsEpi8([64]byte(src), uint64(k), [64]byte(a)))
}

func m512MaskAbsEpi8(src [64]byte, k uint64, a [64]byte) [64]byte


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
func M512MaskzAbsEpi8(k x86.Mmask64, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAbsEpi8(uint64(k), [64]byte(a)))
}

func m512MaskzAbsEpi8(k uint64, a [64]byte) [64]byte


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
func MaskAddEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAddEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskAddEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzAddEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAddEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAddEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAddEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAddEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAddEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAddEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAddEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512AddEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AddEpi16([64]byte(a), [64]byte(b)))
}

func m512AddEpi16(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAddEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAddEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAddEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAddEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAddEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAddEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAddEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAddEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskAddEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzAddEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAddEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAddEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAddEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAddEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAddEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAddEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAddEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAddEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512AddEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AddEpi8([64]byte(a), [64]byte(b)))
}

func m512AddEpi8(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAddEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAddEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAddEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAddEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAddEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAddEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAddsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskAddsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzAddsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAddsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAddsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAddsEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAddsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAddsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAddsEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAddsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512AddsEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AddsEpi16([64]byte(a), [64]byte(b)))
}

func m512AddsEpi16(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAddsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAddsEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAddsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAddsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAddsEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAddsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAddsEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskAddsEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzAddsEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAddsEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAddsEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAddsEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAddsEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAddsEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAddsEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAddsEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512AddsEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AddsEpi8([64]byte(a), [64]byte(b)))
}

func m512AddsEpi8(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAddsEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAddsEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAddsEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAddsEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAddsEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAddsEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAddsEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskAddsEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzAddsEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAddsEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAddsEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAddsEpu16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAddsEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAddsEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAddsEpu16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAddsEpu16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512AddsEpu16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AddsEpu16([64]byte(a), [64]byte(b)))
}

func m512AddsEpu16(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAddsEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAddsEpu16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAddsEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAddsEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAddsEpu16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAddsEpu16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAddsEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskAddsEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzAddsEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAddsEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAddsEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAddsEpu8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAddsEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAddsEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAddsEpu8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAddsEpu8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512AddsEpu8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AddsEpu8([64]byte(a), [64]byte(b)))
}

func m512AddsEpu8(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAddsEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAddsEpu8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAddsEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAddsEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAddsEpu8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAddsEpu8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskAlignrEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(maskAlignrEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b), count))
}

func maskAlignrEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte, count int) [16]byte


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
func MaskzAlignrEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(maskzAlignrEpi8(uint16(k), [16]byte(a), [16]byte(b), count))
}

func maskzAlignrEpi8(k uint16, a [16]byte, b [16]byte, count int) [16]byte


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
func M256MaskAlignrEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(m256MaskAlignrEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b), count))
}

func m256MaskAlignrEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte, count int) [32]byte


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
func M256MaskzAlignrEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(m256MaskzAlignrEpi8(uint32(k), [32]byte(a), [32]byte(b), count))
}

func m256MaskzAlignrEpi8(k uint32, a [32]byte, b [32]byte, count int) [32]byte


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
func M512AlignrEpi8(a x86.M512i, b x86.M512i, count int) x86.M512i {
	return x86.M512i(m512AlignrEpi8([64]byte(a), [64]byte(b), count))
}

func m512AlignrEpi8(a [64]byte, b [64]byte, count int) [64]byte


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
func M512MaskAlignrEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i, count int) x86.M512i {
	return x86.M512i(m512MaskAlignrEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b), count))
}

func m512MaskAlignrEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte, count int) [64]byte


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
func M512MaskzAlignrEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i, count int) x86.M512i {
	return x86.M512i(m512MaskzAlignrEpi8(uint64(k), [64]byte(a), [64]byte(b), count))
}

func m512MaskzAlignrEpi8(k uint64, a [64]byte, b [64]byte, count int) [64]byte


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
func MaskAvgEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAvgEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskAvgEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzAvgEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAvgEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAvgEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAvgEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAvgEpu16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAvgEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAvgEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAvgEpu16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAvgEpu16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512AvgEpu16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AvgEpu16([64]byte(a), [64]byte(b)))
}

func m512AvgEpu16(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAvgEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAvgEpu16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAvgEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAvgEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAvgEpu16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAvgEpu16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAvgEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskAvgEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskAvgEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzAvgEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzAvgEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAvgEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskAvgEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskAvgEpu8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskAvgEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzAvgEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzAvgEpu8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzAvgEpu8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512AvgEpu8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512AvgEpu8([64]byte(a), [64]byte(b)))
}

func m512AvgEpu8(a [64]byte, b [64]byte) [64]byte


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
func M512MaskAvgEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskAvgEpu8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskAvgEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzAvgEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzAvgEpu8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzAvgEpu8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskBlendEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskBlendEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskBlendEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskBlendEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskBlendEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskBlendEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskBlendEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskBlendEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskBlendEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskBlendEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskBlendEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskBlendEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskBlendEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskBlendEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskBlendEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskBlendEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskBlendEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskBlendEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskBroadcastbEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i) x86.M128i {
	return x86.M128i(maskBroadcastbEpi8([16]byte(src), uint16(k), [16]byte(a)))
}

func maskBroadcastbEpi8(src [16]byte, k uint16, a [16]byte) [16]byte


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
func MaskzBroadcastbEpi8(k x86.Mmask16, a x86.M128i) x86.M128i {
	return x86.M128i(maskzBroadcastbEpi8(uint16(k), [16]byte(a)))
}

func maskzBroadcastbEpi8(k uint16, a [16]byte) [16]byte


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
func M256MaskBroadcastbEpi8(src x86.M256i, k x86.Mmask32, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskBroadcastbEpi8([32]byte(src), uint32(k), [16]byte(a)))
}

func m256MaskBroadcastbEpi8(src [32]byte, k uint32, a [16]byte) [32]byte


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
func M256MaskzBroadcastbEpi8(k x86.Mmask32, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzBroadcastbEpi8(uint32(k), [16]byte(a)))
}

func m256MaskzBroadcastbEpi8(k uint32, a [16]byte) [32]byte


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
func M512BroadcastbEpi8(a x86.M128i) x86.M512i {
	return x86.M512i(m512BroadcastbEpi8([16]byte(a)))
}

func m512BroadcastbEpi8(a [16]byte) [64]byte


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
func M512MaskBroadcastbEpi8(src x86.M512i, k x86.Mmask64, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskBroadcastbEpi8([64]byte(src), uint64(k), [16]byte(a)))
}

func m512MaskBroadcastbEpi8(src [64]byte, k uint64, a [16]byte) [64]byte


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
func M512MaskzBroadcastbEpi8(k x86.Mmask64, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzBroadcastbEpi8(uint64(k), [16]byte(a)))
}

func m512MaskzBroadcastbEpi8(k uint64, a [16]byte) [64]byte


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
func MaskBroadcastwEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskBroadcastwEpi16([16]byte(src), uint8(k), [16]byte(a)))
}

func maskBroadcastwEpi16(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzBroadcastwEpi16(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzBroadcastwEpi16(uint8(k), [16]byte(a)))
}

func maskzBroadcastwEpi16(k uint8, a [16]byte) [16]byte


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
func M256MaskBroadcastwEpi16(src x86.M256i, k x86.Mmask16, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskBroadcastwEpi16([32]byte(src), uint16(k), [16]byte(a)))
}

func m256MaskBroadcastwEpi16(src [32]byte, k uint16, a [16]byte) [32]byte


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
func M256MaskzBroadcastwEpi16(k x86.Mmask16, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzBroadcastwEpi16(uint16(k), [16]byte(a)))
}

func m256MaskzBroadcastwEpi16(k uint16, a [16]byte) [32]byte


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
func M512BroadcastwEpi16(a x86.M128i) x86.M512i {
	return x86.M512i(m512BroadcastwEpi16([16]byte(a)))
}

func m512BroadcastwEpi16(a [16]byte) [64]byte


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
func M512MaskBroadcastwEpi16(src x86.M512i, k x86.Mmask32, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskBroadcastwEpi16([64]byte(src), uint32(k), [16]byte(a)))
}

func m512MaskBroadcastwEpi16(src [64]byte, k uint32, a [16]byte) [64]byte


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
func M512MaskzBroadcastwEpi16(k x86.Mmask32, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzBroadcastwEpi16(uint32(k), [16]byte(a)))
}

func m512MaskzBroadcastwEpi16(k uint32, a [16]byte) [64]byte


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
func M512BslliEpi128(a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512BslliEpi128([64]byte(a), imm8))
}

func m512BslliEpi128(a [64]byte, imm8 int) [64]byte


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
func M512BsrliEpi128(a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512BsrliEpi128([64]byte(a), imm8))
}

func m512BsrliEpi128(a [64]byte, imm8 int) [64]byte


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
func CmpEpi16Mask(a x86.M128i, b x86.M128i, imm8 int) x86.Mmask8 {
	return x86.Mmask8(cmpEpi16Mask([16]byte(a), [16]byte(b), imm8))
}

func cmpEpi16Mask(a [16]byte, b [16]byte, imm8 int) uint8


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
func MaskCmpEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskCmpEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpi16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 int) uint8


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
func M256CmpEpi16Mask(a x86.M256i, b x86.M256i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(m256CmpEpi16Mask([32]byte(a), [32]byte(b), imm8))
}

func m256CmpEpi16Mask(a [32]byte, b [32]byte, imm8 int) uint16


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
func M256MaskCmpEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b), imm8))
}

func m256MaskCmpEpi16Mask(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16


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
func M512CmpEpi16Mask(a x86.M512i, b x86.M512i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m512CmpEpi16Mask([64]byte(a), [64]byte(b), imm8))
}

func m512CmpEpi16Mask(a [64]byte, b [64]byte, imm8 int) uint32


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
func M512MaskCmpEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b), imm8))
}

func m512MaskCmpEpi16Mask(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32


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
func CmpEpi8Mask(a x86.M128i, b x86.M128i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(cmpEpi8Mask([16]byte(a), [16]byte(b), imm8))
}

func cmpEpi8Mask(a [16]byte, b [16]byte, imm8 int) uint16


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
func MaskCmpEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(maskCmpEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpi8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 int) uint16


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
func M256CmpEpi8Mask(a x86.M256i, b x86.M256i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m256CmpEpi8Mask([32]byte(a), [32]byte(b), imm8))
}

func m256CmpEpi8Mask(a [32]byte, b [32]byte, imm8 int) uint32


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
func M256MaskCmpEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b), imm8))
}

func m256MaskCmpEpi8Mask(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32


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
func M512CmpEpi8Mask(a x86.M512i, b x86.M512i, imm8 int) x86.Mmask64 {
	return x86.Mmask64(m512CmpEpi8Mask([64]byte(a), [64]byte(b), imm8))
}

func m512CmpEpi8Mask(a [64]byte, b [64]byte, imm8 int) uint64


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
func M512MaskCmpEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i, imm8 int) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b), imm8))
}

func m512MaskCmpEpi8Mask(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64


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
func CmpEpu16Mask(a x86.M128i, b x86.M128i, imm8 int) x86.Mmask8 {
	return x86.Mmask8(cmpEpu16Mask([16]byte(a), [16]byte(b), imm8))
}

func cmpEpu16Mask(a [16]byte, b [16]byte, imm8 int) uint8


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
func MaskCmpEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskCmpEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpu16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 int) uint8


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
func M256CmpEpu16Mask(a x86.M256i, b x86.M256i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(m256CmpEpu16Mask([32]byte(a), [32]byte(b), imm8))
}

func m256CmpEpu16Mask(a [32]byte, b [32]byte, imm8 int) uint16


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
func M256MaskCmpEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b), imm8))
}

func m256MaskCmpEpu16Mask(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16


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
func M512CmpEpu16Mask(a x86.M512i, b x86.M512i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m512CmpEpu16Mask([64]byte(a), [64]byte(b), imm8))
}

func m512CmpEpu16Mask(a [64]byte, b [64]byte, imm8 int) uint32


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
func M512MaskCmpEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b), imm8))
}

func m512MaskCmpEpu16Mask(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32


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
func CmpEpu8Mask(a x86.M128i, b x86.M128i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(cmpEpu8Mask([16]byte(a), [16]byte(b), imm8))
}

func cmpEpu8Mask(a [16]byte, b [16]byte, imm8 int) uint16


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
func MaskCmpEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i, imm8 int) x86.Mmask16 {
	return x86.Mmask16(maskCmpEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpu8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 int) uint16


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
func M256CmpEpu8Mask(a x86.M256i, b x86.M256i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m256CmpEpu8Mask([32]byte(a), [32]byte(b), imm8))
}

func m256CmpEpu8Mask(a [32]byte, b [32]byte, imm8 int) uint32


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
func M256MaskCmpEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i, imm8 int) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b), imm8))
}

func m256MaskCmpEpu8Mask(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32


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
func M512CmpEpu8Mask(a x86.M512i, b x86.M512i, imm8 int) x86.Mmask64 {
	return x86.Mmask64(m512CmpEpu8Mask([64]byte(a), [64]byte(b), imm8))
}

func m512CmpEpu8Mask(a [64]byte, b [64]byte, imm8 int) uint64


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
func M512MaskCmpEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i, imm8 int) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b), imm8))
}

func m512MaskCmpEpu8Mask(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64


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
func CmpeqEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpeqEpi16Mask([16]byte(a), [16]byte(b)))
}

func cmpeqEpi16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpeqEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpeqEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpeqEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpeqEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpeqEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpeqEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpeqEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpeqEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpeqEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpeqEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpeqEpi16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpeqEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpeqEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpeqEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpeqEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpeqEpi8Mask([16]byte(a), [16]byte(b)))
}

func cmpeqEpi8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpeqEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpeqEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpeqEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpeqEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpeqEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpeqEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpeqEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpeqEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpeqEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpeqEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpeqEpi8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpeqEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpeqEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpeqEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpeqEpu16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpeqEpu16Mask([16]byte(a), [16]byte(b)))
}

func cmpeqEpu16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpeqEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpeqEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpeqEpu16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpeqEpu16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpeqEpu16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpeqEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpeqEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpeqEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpeqEpu16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpeqEpu16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpeqEpu16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpeqEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpeqEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpeqEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpeqEpu8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpeqEpu8Mask([16]byte(a), [16]byte(b)))
}

func cmpeqEpu8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpeqEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpeqEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpeqEpu8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpeqEpu8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpeqEpu8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpeqEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpeqEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpeqEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpeqEpu8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpeqEpu8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpeqEpu8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpeqEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpeqEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpeqEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgeEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpgeEpi16Mask([16]byte(a), [16]byte(b)))
}

func cmpgeEpi16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpgeEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpgeEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpgeEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpgeEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgeEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpgeEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpgeEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgeEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpgeEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpgeEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgeEpi16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpgeEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpgeEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgeEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgeEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpgeEpi8Mask([16]byte(a), [16]byte(b)))
}

func cmpgeEpi8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpgeEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgeEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpgeEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpgeEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgeEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpgeEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpgeEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgeEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpgeEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpgeEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgeEpi8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpgeEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpgeEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgeEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgeEpu16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpgeEpu16Mask([16]byte(a), [16]byte(b)))
}

func cmpgeEpu16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpgeEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpgeEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpgeEpu16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpgeEpu16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgeEpu16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpgeEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpgeEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgeEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpgeEpu16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpgeEpu16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgeEpu16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpgeEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpgeEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgeEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgeEpu8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpgeEpu8Mask([16]byte(a), [16]byte(b)))
}

func cmpgeEpu8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpgeEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgeEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpgeEpu8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpgeEpu8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgeEpu8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpgeEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpgeEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgeEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpgeEpu8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpgeEpu8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgeEpu8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpgeEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpgeEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgeEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgtEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpgtEpi16Mask([16]byte(a), [16]byte(b)))
}

func cmpgtEpi16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpgtEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpgtEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpgtEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpgtEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgtEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpgtEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpgtEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgtEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpgtEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpgtEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgtEpi16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpgtEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpgtEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgtEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgtEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpgtEpi8Mask([16]byte(a), [16]byte(b)))
}

func cmpgtEpi8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpgtEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgtEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpgtEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpgtEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgtEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpgtEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpgtEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgtEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpgtEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpgtEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgtEpi8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpgtEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpgtEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgtEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgtEpu16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpgtEpu16Mask([16]byte(a), [16]byte(b)))
}

func cmpgtEpu16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpgtEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpgtEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpgtEpu16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpgtEpu16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgtEpu16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpgtEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpgtEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgtEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpgtEpu16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpgtEpu16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgtEpu16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpgtEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpgtEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgtEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgtEpu8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpgtEpu8Mask([16]byte(a), [16]byte(b)))
}

func cmpgtEpu8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpgtEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgtEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpgtEpu8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpgtEpu8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpgtEpu8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpgtEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpgtEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpgtEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpgtEpu8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpgtEpu8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpgtEpu8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpgtEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpgtEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpgtEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpleEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpleEpi16Mask([16]byte(a), [16]byte(b)))
}

func cmpleEpi16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpleEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpleEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpleEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpleEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpleEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpleEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpleEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpleEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpleEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpleEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpleEpi16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpleEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpleEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpleEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpleEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpleEpi8Mask([16]byte(a), [16]byte(b)))
}

func cmpleEpi8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpleEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpleEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpleEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpleEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpleEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpleEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpleEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpleEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpleEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpleEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpleEpi8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpleEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpleEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpleEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpleEpu16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpleEpu16Mask([16]byte(a), [16]byte(b)))
}

func cmpleEpu16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpleEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpleEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpleEpu16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpleEpu16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpleEpu16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpleEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpleEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpleEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpleEpu16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpleEpu16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpleEpu16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpleEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpleEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpleEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpleEpu8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpleEpu8Mask([16]byte(a), [16]byte(b)))
}

func cmpleEpu8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpleEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpleEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpleEpu8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpleEpu8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpleEpu8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpleEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpleEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpleEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpleEpu8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpleEpu8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpleEpu8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpleEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpleEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpleEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpltEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpltEpi16Mask([16]byte(a), [16]byte(b)))
}

func cmpltEpi16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpltEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpltEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpltEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpltEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpltEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpltEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpltEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpltEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpltEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpltEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpltEpi16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpltEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpltEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpltEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpltEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpltEpi8Mask([16]byte(a), [16]byte(b)))
}

func cmpltEpi8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpltEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpltEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpltEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpltEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpltEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpltEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpltEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpltEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpltEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpltEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpltEpi8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpltEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpltEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpltEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpltEpu16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpltEpu16Mask([16]byte(a), [16]byte(b)))
}

func cmpltEpu16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpltEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpltEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpltEpu16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpltEpu16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpltEpu16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpltEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpltEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpltEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpltEpu16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpltEpu16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpltEpu16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpltEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpltEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpltEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpltEpu8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpltEpu8Mask([16]byte(a), [16]byte(b)))
}

func cmpltEpu8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpltEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpltEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpltEpu8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpltEpu8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpltEpu8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpltEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpltEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpltEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpltEpu8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpltEpu8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpltEpu8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpltEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpltEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpltEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpneqEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpneqEpi16Mask([16]byte(a), [16]byte(b)))
}

func cmpneqEpi16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpneqEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpneqEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpneqEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpneqEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpneqEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpneqEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpneqEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpneqEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpneqEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpneqEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpneqEpi16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpneqEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpneqEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpneqEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpneqEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpneqEpi8Mask([16]byte(a), [16]byte(b)))
}

func cmpneqEpi8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpneqEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpneqEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpneqEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpneqEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpneqEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpneqEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpneqEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpneqEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpneqEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpneqEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpneqEpi8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpneqEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpneqEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpneqEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpneqEpu16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(cmpneqEpu16Mask([16]byte(a), [16]byte(b)))
}

func cmpneqEpu16Mask(a [16]byte, b [16]byte) uint8


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
func MaskCmpneqEpu16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskCmpneqEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func M256CmpneqEpu16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256CmpneqEpu16Mask([32]byte(a), [32]byte(b)))
}

func m256CmpneqEpu16Mask(a [32]byte, b [32]byte) uint16


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
func M256MaskCmpneqEpu16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskCmpneqEpu16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpneqEpu16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M512CmpneqEpu16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512CmpneqEpu16Mask([64]byte(a), [64]byte(b)))
}

func m512CmpneqEpu16Mask(a [64]byte, b [64]byte) uint32


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
func M512MaskCmpneqEpu16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskCmpneqEpu16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpneqEpu16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpneqEpu8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(cmpneqEpu8Mask([16]byte(a), [16]byte(b)))
}

func cmpneqEpu8Mask(a [16]byte, b [16]byte) uint16


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
func MaskCmpneqEpu8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskCmpneqEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func M256CmpneqEpu8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256CmpneqEpu8Mask([32]byte(a), [32]byte(b)))
}

func m256CmpneqEpu8Mask(a [32]byte, b [32]byte) uint32


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
func M256MaskCmpneqEpu8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskCmpneqEpu8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskCmpneqEpu8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M512CmpneqEpu8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512CmpneqEpu8Mask([64]byte(a), [64]byte(b)))
}

func m512CmpneqEpu8Mask(a [64]byte, b [64]byte) uint64


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
func M512MaskCmpneqEpu8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskCmpneqEpu8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskCmpneqEpu8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func Cvtepi16Epi8(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi16Epi8([16]byte(a)))
}

func cvtepi16Epi8(a [16]byte) [16]byte


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
func MaskCvtepi16Epi8(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskCvtepi16Epi8([16]byte(src), uint8(k), [16]byte(a)))
}

func maskCvtepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzCvtepi16Epi8(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzCvtepi16Epi8(uint8(k), [16]byte(a)))
}

func maskzCvtepi16Epi8(k uint8, a [16]byte) [16]byte


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
func M256Cvtepi16Epi8(a x86.M256i) x86.M128i {
	return x86.M128i(m256Cvtepi16Epi8([32]byte(a)))
}

func m256Cvtepi16Epi8(a [32]byte) [16]byte


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
func M256MaskCvtepi16Epi8(src x86.M128i, k x86.Mmask16, a x86.M256i) x86.M128i {
	return x86.M128i(m256MaskCvtepi16Epi8([16]byte(src), uint16(k), [32]byte(a)))
}

func m256MaskCvtepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte


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
func M256MaskzCvtepi16Epi8(k x86.Mmask16, a x86.M256i) x86.M128i {
	return x86.M128i(m256MaskzCvtepi16Epi8(uint16(k), [32]byte(a)))
}

func m256MaskzCvtepi16Epi8(k uint16, a [32]byte) [16]byte


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
func M512Cvtepi16Epi8(a x86.M512i) x86.M256i {
	return x86.M256i(m512Cvtepi16Epi8([64]byte(a)))
}

func m512Cvtepi16Epi8(a [64]byte) [32]byte


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
func M512MaskCvtepi16Epi8(src x86.M256i, k x86.Mmask32, a x86.M512i) x86.M256i {
	return x86.M256i(m512MaskCvtepi16Epi8([32]byte(src), uint32(k), [64]byte(a)))
}

func m512MaskCvtepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte


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
func M512MaskzCvtepi16Epi8(k x86.Mmask32, a x86.M512i) x86.M256i {
	return x86.M256i(m512MaskzCvtepi16Epi8(uint32(k), [64]byte(a)))
}

func m512MaskzCvtepi16Epi8(k uint32, a [64]byte) [32]byte


// MaskCvtepi16StoreuEpi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with truncation, and store the active results (those with
// their respective bit set in writemask 'k') to unaligned memory at
// 'base_addr'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Truncate_Int16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm_mask_cvtepi16_storeu_epi8'.
// Requires AVX512BW.
func MaskCvtepi16StoreuEpi8(base_addr uintptr, k x86.Mmask8, a x86.M128i)  {
	maskCvtepi16StoreuEpi8(uintptr(base_addr), uint8(k), [16]byte(a))
}

func maskCvtepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 


// M256MaskCvtepi16StoreuEpi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with truncation, and store the active results (those with
// their respective bit set in writemask 'k') to unaligned memory at
// 'base_addr'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Truncate_Int16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm256_mask_cvtepi16_storeu_epi8'.
// Requires AVX512BW.
func M256MaskCvtepi16StoreuEpi8(base_addr uintptr, k x86.Mmask16, a x86.M256i)  {
	m256MaskCvtepi16StoreuEpi8(uintptr(base_addr), uint16(k), [32]byte(a))
}

func m256MaskCvtepi16StoreuEpi8(base_addr uintptr, k uint16, a [32]byte) 


// M512MaskCvtepi16StoreuEpi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with truncation, and store the active results (those with
// their respective bit set in writemask 'k') to unaligned memory at
// 'base_addr'. 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Truncate_Int16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVWB'. Intrinsic: '_mm512_mask_cvtepi16_storeu_epi8'.
// Requires AVX512BW.
func M512MaskCvtepi16StoreuEpi8(base_addr uintptr, k x86.Mmask32, a x86.M512i)  {
	m512MaskCvtepi16StoreuEpi8(uintptr(base_addr), uint32(k), [64]byte(a))
}

func m512MaskCvtepi16StoreuEpi8(base_addr uintptr, k uint32, a [64]byte) 


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
func MaskCvtepi8Epi16(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskCvtepi8Epi16([16]byte(src), uint8(k), [16]byte(a)))
}

func maskCvtepi8Epi16(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzCvtepi8Epi16(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzCvtepi8Epi16(uint8(k), [16]byte(a)))
}

func maskzCvtepi8Epi16(k uint8, a [16]byte) [16]byte


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
func M256MaskCvtepi8Epi16(src x86.M256i, k x86.Mmask16, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskCvtepi8Epi16([32]byte(src), uint16(k), [16]byte(a)))
}

func m256MaskCvtepi8Epi16(src [32]byte, k uint16, a [16]byte) [32]byte


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
func M256MaskzCvtepi8Epi16(k x86.Mmask16, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzCvtepi8Epi16(uint16(k), [16]byte(a)))
}

func m256MaskzCvtepi8Epi16(k uint16, a [16]byte) [32]byte


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
func M512Cvtepi8Epi16(a x86.M256i) x86.M512i {
	return x86.M512i(m512Cvtepi8Epi16([32]byte(a)))
}

func m512Cvtepi8Epi16(a [32]byte) [64]byte


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
func M512MaskCvtepi8Epi16(src x86.M512i, k x86.Mmask32, a x86.M256i) x86.M512i {
	return x86.M512i(m512MaskCvtepi8Epi16([64]byte(src), uint32(k), [32]byte(a)))
}

func m512MaskCvtepi8Epi16(src [64]byte, k uint32, a [32]byte) [64]byte


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
func M512MaskzCvtepi8Epi16(k x86.Mmask32, a x86.M256i) x86.M512i {
	return x86.M512i(m512MaskzCvtepi8Epi16(uint32(k), [32]byte(a)))
}

func m512MaskzCvtepi8Epi16(k uint32, a [32]byte) [64]byte


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
func MaskCvtepu8Epi16(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskCvtepu8Epi16([16]byte(src), uint8(k), [16]byte(a)))
}

func maskCvtepu8Epi16(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzCvtepu8Epi16(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzCvtepu8Epi16(uint8(k), [16]byte(a)))
}

func maskzCvtepu8Epi16(k uint8, a [16]byte) [16]byte


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
func M256MaskCvtepu8Epi16(src x86.M256i, k x86.Mmask16, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskCvtepu8Epi16([32]byte(src), uint16(k), [16]byte(a)))
}

func m256MaskCvtepu8Epi16(src [32]byte, k uint16, a [16]byte) [32]byte


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
func M256MaskzCvtepu8Epi16(k x86.Mmask16, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzCvtepu8Epi16(uint16(k), [16]byte(a)))
}

func m256MaskzCvtepu8Epi16(k uint16, a [16]byte) [32]byte


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
func M512Cvtepu8Epi16(a x86.M256i) x86.M512i {
	return x86.M512i(m512Cvtepu8Epi16([32]byte(a)))
}

func m512Cvtepu8Epi16(a [32]byte) [64]byte


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
func M512MaskCvtepu8Epi16(src x86.M512i, k x86.Mmask32, a x86.M256i) x86.M512i {
	return x86.M512i(m512MaskCvtepu8Epi16([64]byte(src), uint32(k), [32]byte(a)))
}

func m512MaskCvtepu8Epi16(src [64]byte, k uint32, a [32]byte) [64]byte


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
func M512MaskzCvtepu8Epi16(k x86.Mmask32, a x86.M256i) x86.M512i {
	return x86.M512i(m512MaskzCvtepu8Epi16(uint32(k), [32]byte(a)))
}

func m512MaskzCvtepu8Epi16(k uint32, a [32]byte) [64]byte


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
func Cvtsepi16Epi8(a x86.M128i) x86.M128i {
	return x86.M128i(cvtsepi16Epi8([16]byte(a)))
}

func cvtsepi16Epi8(a [16]byte) [16]byte


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
func MaskCvtsepi16Epi8(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskCvtsepi16Epi8([16]byte(src), uint8(k), [16]byte(a)))
}

func maskCvtsepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzCvtsepi16Epi8(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzCvtsepi16Epi8(uint8(k), [16]byte(a)))
}

func maskzCvtsepi16Epi8(k uint8, a [16]byte) [16]byte


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
func M256Cvtsepi16Epi8(a x86.M256i) x86.M128i {
	return x86.M128i(m256Cvtsepi16Epi8([32]byte(a)))
}

func m256Cvtsepi16Epi8(a [32]byte) [16]byte


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
func M256MaskCvtsepi16Epi8(src x86.M128i, k x86.Mmask16, a x86.M256i) x86.M128i {
	return x86.M128i(m256MaskCvtsepi16Epi8([16]byte(src), uint16(k), [32]byte(a)))
}

func m256MaskCvtsepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte


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
func M256MaskzCvtsepi16Epi8(k x86.Mmask16, a x86.M256i) x86.M128i {
	return x86.M128i(m256MaskzCvtsepi16Epi8(uint16(k), [32]byte(a)))
}

func m256MaskzCvtsepi16Epi8(k uint16, a [32]byte) [16]byte


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
func M512Cvtsepi16Epi8(a x86.M512i) x86.M256i {
	return x86.M256i(m512Cvtsepi16Epi8([64]byte(a)))
}

func m512Cvtsepi16Epi8(a [64]byte) [32]byte


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
func M512MaskCvtsepi16Epi8(src x86.M256i, k x86.Mmask32, a x86.M512i) x86.M256i {
	return x86.M256i(m512MaskCvtsepi16Epi8([32]byte(src), uint32(k), [64]byte(a)))
}

func m512MaskCvtsepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte


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
func M512MaskzCvtsepi16Epi8(k x86.Mmask32, a x86.M512i) x86.M256i {
	return x86.M256i(m512MaskzCvtsepi16Epi8(uint32(k), [64]byte(a)))
}

func m512MaskzCvtsepi16Epi8(k uint32, a [64]byte) [32]byte


// MaskCvtsepi16StoreuEpi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with signed saturation, and store the active results (those
// with their respective bit set in writemask 'k') to unaligned memory at
// 'base_addr'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Saturate_Int16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm_mask_cvtsepi16_storeu_epi8'.
// Requires AVX512BW.
func MaskCvtsepi16StoreuEpi8(base_addr uintptr, k x86.Mmask8, a x86.M128i)  {
	maskCvtsepi16StoreuEpi8(uintptr(base_addr), uint8(k), [16]byte(a))
}

func maskCvtsepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 


// M256MaskCvtsepi16StoreuEpi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with signed saturation, and store the active results (those
// with their respective bit set in writemask 'k') to unaligned memory at
// 'base_addr'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Saturate_Int16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm256_mask_cvtsepi16_storeu_epi8'.
// Requires AVX512BW.
func M256MaskCvtsepi16StoreuEpi8(base_addr uintptr, k x86.Mmask16, a x86.M256i)  {
	m256MaskCvtsepi16StoreuEpi8(uintptr(base_addr), uint16(k), [32]byte(a))
}

func m256MaskCvtsepi16StoreuEpi8(base_addr uintptr, k uint16, a [32]byte) 


// M512MaskCvtsepi16StoreuEpi8: Convert packed 16-bit integers in 'a' to packed
// 8-bit integers with signed saturation, and store the active results (those
// with their respective bit set in writemask 'k') to unaligned memory at
// 'base_addr'. 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Saturate_Int16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVSWB'. Intrinsic: '_mm512_mask_cvtsepi16_storeu_epi8'.
// Requires AVX512BW.
func M512MaskCvtsepi16StoreuEpi8(base_addr uintptr, k x86.Mmask32, a x86.M512i)  {
	m512MaskCvtsepi16StoreuEpi8(uintptr(base_addr), uint32(k), [64]byte(a))
}

func m512MaskCvtsepi16StoreuEpi8(base_addr uintptr, k uint32, a [64]byte) 


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
func Cvtusepi16Epi8(a x86.M128i) x86.M128i {
	return x86.M128i(cvtusepi16Epi8([16]byte(a)))
}

func cvtusepi16Epi8(a [16]byte) [16]byte


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
func MaskCvtusepi16Epi8(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskCvtusepi16Epi8([16]byte(src), uint8(k), [16]byte(a)))
}

func maskCvtusepi16Epi8(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzCvtusepi16Epi8(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzCvtusepi16Epi8(uint8(k), [16]byte(a)))
}

func maskzCvtusepi16Epi8(k uint8, a [16]byte) [16]byte


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
func M256Cvtusepi16Epi8(a x86.M256i) x86.M128i {
	return x86.M128i(m256Cvtusepi16Epi8([32]byte(a)))
}

func m256Cvtusepi16Epi8(a [32]byte) [16]byte


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
func M256MaskCvtusepi16Epi8(src x86.M128i, k x86.Mmask16, a x86.M256i) x86.M128i {
	return x86.M128i(m256MaskCvtusepi16Epi8([16]byte(src), uint16(k), [32]byte(a)))
}

func m256MaskCvtusepi16Epi8(src [16]byte, k uint16, a [32]byte) [16]byte


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
func M256MaskzCvtusepi16Epi8(k x86.Mmask16, a x86.M256i) x86.M128i {
	return x86.M128i(m256MaskzCvtusepi16Epi8(uint16(k), [32]byte(a)))
}

func m256MaskzCvtusepi16Epi8(k uint16, a [32]byte) [16]byte


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
func M512Cvtusepi16Epi8(a x86.M512i) x86.M256i {
	return x86.M256i(m512Cvtusepi16Epi8([64]byte(a)))
}

func m512Cvtusepi16Epi8(a [64]byte) [32]byte


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
func M512MaskCvtusepi16Epi8(src x86.M256i, k x86.Mmask32, a x86.M512i) x86.M256i {
	return x86.M256i(m512MaskCvtusepi16Epi8([32]byte(src), uint32(k), [64]byte(a)))
}

func m512MaskCvtusepi16Epi8(src [32]byte, k uint32, a [64]byte) [32]byte


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
func M512MaskzCvtusepi16Epi8(k x86.Mmask32, a x86.M512i) x86.M256i {
	return x86.M256i(m512MaskzCvtusepi16Epi8(uint32(k), [64]byte(a)))
}

func m512MaskzCvtusepi16Epi8(k uint32, a [64]byte) [32]byte


// MaskCvtusepi16StoreuEpi8: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
// active results (those with their respective bit set in writemask 'k') to
// unaligned memory at 'base_addr'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm_mask_cvtusepi16_storeu_epi8'.
// Requires AVX512BW.
func MaskCvtusepi16StoreuEpi8(base_addr uintptr, k x86.Mmask8, a x86.M128i)  {
	maskCvtusepi16StoreuEpi8(uintptr(base_addr), uint8(k), [16]byte(a))
}

func maskCvtusepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 


// M256MaskCvtusepi16StoreuEpi8: Convert packed unsigned 16-bit integers in 'a'
// to packed unsigned 8-bit integers with unsigned saturation, and store the
// active results (those with their respective bit set in writemask 'k') to
// unaligned memory at 'base_addr'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm256_mask_cvtusepi16_storeu_epi8'.
// Requires AVX512BW.
func M256MaskCvtusepi16StoreuEpi8(base_addr uintptr, k x86.Mmask16, a x86.M256i)  {
	m256MaskCvtusepi16StoreuEpi8(uintptr(base_addr), uint16(k), [32]byte(a))
}

func m256MaskCvtusepi16StoreuEpi8(base_addr uintptr, k uint16, a [32]byte) 


// M512MaskCvtusepi16StoreuEpi8: Convert packed unsigned 16-bit integers in 'a'
// to packed unsigned 8-bit integers with unsigned saturation, and store the
// active results (those with their respective bit set in writemask 'k') to
// unaligned memory at 'base_addr'. 
//
//		FOR j := 0 to 31
//			i := 16*j
//			l := 8*j
//			IF k[j]
//				MEM[base_addr+l+7:base_addr+l] := Saturate_UnsignedInt16_To_Int8(a[i+15:i])
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVUSWB'. Intrinsic: '_mm512_mask_cvtusepi16_storeu_epi8'.
// Requires AVX512BW.
func M512MaskCvtusepi16StoreuEpi8(base_addr uintptr, k x86.Mmask32, a x86.M512i)  {
	m512MaskCvtusepi16StoreuEpi8(uintptr(base_addr), uint32(k), [64]byte(a))
}

func m512MaskCvtusepi16StoreuEpi8(base_addr uintptr, k uint32, a [64]byte) 


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
func DbsadEpu8(a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(dbsadEpu8([16]byte(a), [16]byte(b), imm8))
}

func dbsadEpu8(a [16]byte, b [16]byte, imm8 int) [16]byte


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
func MaskDbsadEpu8(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskDbsadEpu8([16]byte(src), uint8(k), [16]byte(a), [16]byte(b), imm8))
}

func maskDbsadEpu8(src [16]byte, k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte


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
func MaskzDbsadEpu8(k x86.Mmask8, a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskzDbsadEpu8(uint8(k), [16]byte(a), [16]byte(b), imm8))
}

func maskzDbsadEpu8(k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte


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
func M256DbsadEpu8(a x86.M256i, b x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256DbsadEpu8([32]byte(a), [32]byte(b), imm8))
}

func m256DbsadEpu8(a [32]byte, b [32]byte, imm8 int) [32]byte


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
func M256MaskDbsadEpu8(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskDbsadEpu8([32]byte(src), uint16(k), [32]byte(a), [32]byte(b), imm8))
}

func m256MaskDbsadEpu8(src [32]byte, k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte


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
func M256MaskzDbsadEpu8(k x86.Mmask16, a x86.M256i, b x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskzDbsadEpu8(uint16(k), [32]byte(a), [32]byte(b), imm8))
}

func m256MaskzDbsadEpu8(k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte


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
func M512DbsadEpu8(a x86.M512i, b x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512DbsadEpu8([64]byte(a), [64]byte(b), imm8))
}

func m512DbsadEpu8(a [64]byte, b [64]byte, imm8 int) [64]byte


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
func M512MaskDbsadEpu8(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskDbsadEpu8([64]byte(src), uint32(k), [64]byte(a), [64]byte(b), imm8))
}

func m512MaskDbsadEpu8(src [64]byte, k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte


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
func M512MaskzDbsadEpu8(k x86.Mmask32, a x86.M512i, b x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskzDbsadEpu8(uint32(k), [64]byte(a), [64]byte(b), imm8))
}

func m512MaskzDbsadEpu8(k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte


// M512Kunpackd: Unpack and interleave 32 bits from masks 'a' and 'b', and
// store the 64-bit result in 'k'. 
//
//		k[31:0] := a[31:0]
//		k[63:32] := b[31:0]
//		k[MAX:64] := 0
//
// Instruction: 'KUNPCKDQ'. Intrinsic: '_mm512_kunpackd'.
// Requires AVX512BW.
func M512Kunpackd(a x86.Mmask64, b x86.Mmask64) x86.Mmask64 {
	return x86.Mmask64(m512Kunpackd(uint64(a), uint64(b)))
}

func m512Kunpackd(a uint64, b uint64) uint64


// M512Kunpackw: Unpack and interleave 16 bits from masks 'a' and 'b', and
// store the 32-bit result in 'k'. 
//
//		k[15:0] := a[15:0]
//		k[31:16] := b[15:0]
//		k[MAX:32] := 0
//
// Instruction: 'KUNPCKWD'. Intrinsic: '_mm512_kunpackw'.
// Requires AVX512BW.
func M512Kunpackw(a x86.Mmask32, b x86.Mmask32) x86.Mmask32 {
	return x86.Mmask32(m512Kunpackw(uint32(a), uint32(b)))
}

func m512Kunpackw(a uint32, b uint32) uint32


// MaskLoaduEpi16: Load packed 16-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := MEM[mem_addr+i+15:mem_addr+i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm_mask_loadu_epi16'.
// Requires AVX512BW.
func MaskLoaduEpi16(src x86.M128i, k x86.Mmask8, mem_addr uintptr) x86.M128i {
	return x86.M128i(maskLoaduEpi16([16]byte(src), uint8(k), uintptr(mem_addr)))
}

func maskLoaduEpi16(src [16]byte, k uint8, mem_addr uintptr) [16]byte


// MaskzLoaduEpi16: Load packed 16-bit integers from memory into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := MEM[mem_addr+i+15:mem_addr+i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm_maskz_loadu_epi16'.
// Requires AVX512BW.
func MaskzLoaduEpi16(k x86.Mmask8, mem_addr uintptr) x86.M128i {
	return x86.M128i(maskzLoaduEpi16(uint8(k), uintptr(mem_addr)))
}

func maskzLoaduEpi16(k uint8, mem_addr uintptr) [16]byte


// M256MaskLoaduEpi16: Load packed 16-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := MEM[mem_addr+i+15:mem_addr+i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm256_mask_loadu_epi16'.
// Requires AVX512BW.
func M256MaskLoaduEpi16(src x86.M256i, k x86.Mmask16, mem_addr uintptr) x86.M256i {
	return x86.M256i(m256MaskLoaduEpi16([32]byte(src), uint16(k), uintptr(mem_addr)))
}

func m256MaskLoaduEpi16(src [32]byte, k uint16, mem_addr uintptr) [32]byte


// M256MaskzLoaduEpi16: Load packed 16-bit integers from memory into 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set).
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := MEM[mem_addr+i+15:mem_addr+i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm256_maskz_loadu_epi16'.
// Requires AVX512BW.
func M256MaskzLoaduEpi16(k x86.Mmask16, mem_addr uintptr) x86.M256i {
	return x86.M256i(m256MaskzLoaduEpi16(uint16(k), uintptr(mem_addr)))
}

func m256MaskzLoaduEpi16(k uint16, mem_addr uintptr) [32]byte


// M512MaskLoaduEpi16: Load packed 16-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := MEM[mem_addr+i+15:mem_addr+i]
//			ELSE
//				dst[i+15:i] := src[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm512_mask_loadu_epi16'.
// Requires AVX512BW.
func M512MaskLoaduEpi16(src x86.M512i, k x86.Mmask32, mem_addr uintptr) x86.M512i {
	return x86.M512i(m512MaskLoaduEpi16([64]byte(src), uint32(k), uintptr(mem_addr)))
}

func m512MaskLoaduEpi16(src [64]byte, k uint32, mem_addr uintptr) [64]byte


// M512MaskzLoaduEpi16: Load packed 16-bit integers from memory into 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set).
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				dst[i+15:i] := MEM[mem_addr+i+15:mem_addr+i]
//			ELSE
//				dst[i+15:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm512_maskz_loadu_epi16'.
// Requires AVX512BW.
func M512MaskzLoaduEpi16(k x86.Mmask32, mem_addr uintptr) x86.M512i {
	return x86.M512i(m512MaskzLoaduEpi16(uint32(k), uintptr(mem_addr)))
}

func m512MaskzLoaduEpi16(k uint32, mem_addr uintptr) [64]byte


// MaskLoaduEpi8: Load packed 8-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := MEM[mem_addr+i+7:mem_addr+i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm_mask_loadu_epi8'.
// Requires AVX512BW.
func MaskLoaduEpi8(src x86.M128i, k x86.Mmask16, mem_addr uintptr) x86.M128i {
	return x86.M128i(maskLoaduEpi8([16]byte(src), uint16(k), uintptr(mem_addr)))
}

func maskLoaduEpi8(src [16]byte, k uint16, mem_addr uintptr) [16]byte


// MaskzLoaduEpi8: Load packed 8-bit integers from memory into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := MEM[mem_addr+i+7:mem_addr+i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm_maskz_loadu_epi8'.
// Requires AVX512BW.
func MaskzLoaduEpi8(k x86.Mmask16, mem_addr uintptr) x86.M128i {
	return x86.M128i(maskzLoaduEpi8(uint16(k), uintptr(mem_addr)))
}

func maskzLoaduEpi8(k uint16, mem_addr uintptr) [16]byte


// M256MaskLoaduEpi8: Load packed 8-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := MEM[mem_addr+i+7:mem_addr+i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm256_mask_loadu_epi8'.
// Requires AVX512BW.
func M256MaskLoaduEpi8(src x86.M256i, k x86.Mmask32, mem_addr uintptr) x86.M256i {
	return x86.M256i(m256MaskLoaduEpi8([32]byte(src), uint32(k), uintptr(mem_addr)))
}

func m256MaskLoaduEpi8(src [32]byte, k uint32, mem_addr uintptr) [32]byte


// M256MaskzLoaduEpi8: Load packed 8-bit integers from memory into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := MEM[mem_addr+i+7:mem_addr+i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm256_maskz_loadu_epi8'.
// Requires AVX512BW.
func M256MaskzLoaduEpi8(k x86.Mmask32, mem_addr uintptr) x86.M256i {
	return x86.M256i(m256MaskzLoaduEpi8(uint32(k), uintptr(mem_addr)))
}

func m256MaskzLoaduEpi8(k uint32, mem_addr uintptr) [32]byte


// M512MaskLoaduEpi8: Load packed 8-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := MEM[mem_addr+i+7:mem_addr+i]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm512_mask_loadu_epi8'.
// Requires AVX512BW.
func M512MaskLoaduEpi8(src x86.M512i, k x86.Mmask64, mem_addr uintptr) x86.M512i {
	return x86.M512i(m512MaskLoaduEpi8([64]byte(src), uint64(k), uintptr(mem_addr)))
}

func m512MaskLoaduEpi8(src [64]byte, k uint64, mem_addr uintptr) [64]byte


// M512MaskzLoaduEpi8: Load packed 8-bit integers from memory into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				dst[i+7:i] := MEM[mem_addr+i+7:mem_addr+i]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm512_maskz_loadu_epi8'.
// Requires AVX512BW.
func M512MaskzLoaduEpi8(k x86.Mmask64, mem_addr uintptr) x86.M512i {
	return x86.M512i(m512MaskzLoaduEpi8(uint64(k), uintptr(mem_addr)))
}

func m512MaskzLoaduEpi8(k uint64, mem_addr uintptr) [64]byte


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
func MaskMaddEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMaddEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMaddEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMaddEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMaddEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMaddEpi16(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMaddEpi16([32]byte(src), uint8(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMaddEpi16(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMaddEpi16(k x86.Mmask8, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMaddEpi16(uint8(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMaddEpi16(k uint8, a [32]byte, b [32]byte) [32]byte


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
func M512MaddEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaddEpi16([64]byte(a), [64]byte(b)))
}

func m512MaddEpi16(a [64]byte, b [64]byte) [64]byte


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
func M512MaskMaddEpi16(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMaddEpi16([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMaddEpi16(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMaddEpi16(k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMaddEpi16(uint16(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMaddEpi16(k uint16, a [64]byte, b [64]byte) [64]byte


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
func MaskMaddubsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMaddubsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMaddubsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMaddubsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMaddubsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaddubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMaddubsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMaddubsEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMaddubsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMaddubsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMaddubsEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMaddubsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaddubsEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaddubsEpi16([64]byte(a), [64]byte(b)))
}

func m512MaddubsEpi16(a [64]byte, b [64]byte) [64]byte


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
func M512MaskMaddubsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMaddubsEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMaddubsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMaddubsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMaddubsEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMaddubsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMaxEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMaxEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMaxEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMaxEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMaxEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMaxEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMaxEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMaxEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMaxEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMaxEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMaxEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMaxEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMaxEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMaxEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMaxEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMaxEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaxEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaxEpi16([64]byte(a), [64]byte(b)))
}

func m512MaxEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMaxEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskMaxEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzMaxEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMaxEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMaxEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMaxEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMaxEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMaxEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMaxEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMaxEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMaxEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMaxEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMaxEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMaxEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMaxEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMaxEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaxEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaxEpi8([64]byte(a), [64]byte(b)))
}

func m512MaxEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMaxEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMaxEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMaxEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMaxEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMaxEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMaxEpu16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMaxEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMaxEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMaxEpu16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMaxEpu16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMaxEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMaxEpu16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMaxEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMaxEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMaxEpu16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMaxEpu16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaxEpu16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaxEpu16([64]byte(a), [64]byte(b)))
}

func m512MaxEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMaxEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskMaxEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzMaxEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMaxEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMaxEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMaxEpu8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMaxEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMaxEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMaxEpu8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMaxEpu8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMaxEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMaxEpu8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMaxEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMaxEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMaxEpu8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMaxEpu8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaxEpu8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaxEpu8([64]byte(a), [64]byte(b)))
}

func m512MaxEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMinEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMinEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMinEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMinEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMinEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMinEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMinEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMinEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMinEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMinEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMinEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMinEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMinEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMinEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMinEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMinEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MinEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MinEpi16([64]byte(a), [64]byte(b)))
}

func m512MinEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMinEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskMinEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzMinEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMinEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMinEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMinEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMinEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMinEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMinEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMinEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMinEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMinEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMinEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMinEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMinEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMinEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MinEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MinEpi8([64]byte(a), [64]byte(b)))
}

func m512MinEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMinEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMinEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMinEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMinEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMinEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMinEpu16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMinEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMinEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMinEpu16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMinEpu16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMinEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMinEpu16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMinEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMinEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMinEpu16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMinEpu16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MinEpu16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MinEpu16([64]byte(a), [64]byte(b)))
}

func m512MinEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMinEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskMinEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzMinEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMinEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMinEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMinEpu8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMinEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMinEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMinEpu8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMinEpu8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMinEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMinEpu8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMinEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMinEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMinEpu8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMinEpu8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MinEpu8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MinEpu8([64]byte(a), [64]byte(b)))
}

func m512MinEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskMovEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskMovEpi16([16]byte(src), uint8(k), [16]byte(a)))
}

func maskMovEpi16(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzMovEpi16(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzMovEpi16(uint8(k), [16]byte(a)))
}

func maskzMovEpi16(k uint8, a [16]byte) [16]byte


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
func M256MaskMovEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMovEpi16([32]byte(src), uint16(k), [32]byte(a)))
}

func m256MaskMovEpi16(src [32]byte, k uint16, a [32]byte) [32]byte


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
func M256MaskzMovEpi16(k x86.Mmask16, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMovEpi16(uint16(k), [32]byte(a)))
}

func m256MaskzMovEpi16(k uint16, a [32]byte) [32]byte


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
func M512MaskMovEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMovEpi16([64]byte(src), uint32(k), [64]byte(a)))
}

func m512MaskMovEpi16(src [64]byte, k uint32, a [64]byte) [64]byte


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
func M512MaskzMovEpi16(k x86.Mmask32, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMovEpi16(uint32(k), [64]byte(a)))
}

func m512MaskzMovEpi16(k uint32, a [64]byte) [64]byte


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
func MaskMovEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i) x86.M128i {
	return x86.M128i(maskMovEpi8([16]byte(src), uint16(k), [16]byte(a)))
}

func maskMovEpi8(src [16]byte, k uint16, a [16]byte) [16]byte


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
func MaskzMovEpi8(k x86.Mmask16, a x86.M128i) x86.M128i {
	return x86.M128i(maskzMovEpi8(uint16(k), [16]byte(a)))
}

func maskzMovEpi8(k uint16, a [16]byte) [16]byte


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
func M256MaskMovEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMovEpi8([32]byte(src), uint32(k), [32]byte(a)))
}

func m256MaskMovEpi8(src [32]byte, k uint32, a [32]byte) [32]byte


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
func M256MaskzMovEpi8(k x86.Mmask32, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMovEpi8(uint32(k), [32]byte(a)))
}

func m256MaskzMovEpi8(k uint32, a [32]byte) [32]byte


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
func M512MaskMovEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMovEpi8([64]byte(src), uint64(k), [64]byte(a)))
}

func m512MaskMovEpi8(src [64]byte, k uint64, a [64]byte) [64]byte


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
func M512MaskzMovEpi8(k x86.Mmask64, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMovEpi8(uint64(k), [64]byte(a)))
}

func m512MaskzMovEpi8(k uint64, a [64]byte) [64]byte


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
func Movepi16Mask(a x86.M128i) x86.Mmask8 {
	return x86.Mmask8(movepi16Mask([16]byte(a)))
}

func movepi16Mask(a [16]byte) uint8


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
func M256Movepi16Mask(a x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256Movepi16Mask([32]byte(a)))
}

func m256Movepi16Mask(a [32]byte) uint16


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
func M512Movepi16Mask(a x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512Movepi16Mask([64]byte(a)))
}

func m512Movepi16Mask(a [64]byte) uint32


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
func Movepi8Mask(a x86.M128i) x86.Mmask16 {
	return x86.Mmask16(movepi8Mask([16]byte(a)))
}

func movepi8Mask(a [16]byte) uint16


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
func M256Movepi8Mask(a x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256Movepi8Mask([32]byte(a)))
}

func m256Movepi8Mask(a [32]byte) uint32


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
func M512Movepi8Mask(a x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512Movepi8Mask([64]byte(a)))
}

func m512Movepi8Mask(a [64]byte) uint64


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
func MovmEpi16(k x86.Mmask8) x86.M128i {
	return x86.M128i(movmEpi16(uint8(k)))
}

func movmEpi16(k uint8) [16]byte


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
func M256MovmEpi16(k x86.Mmask16) x86.M256i {
	return x86.M256i(m256MovmEpi16(uint16(k)))
}

func m256MovmEpi16(k uint16) [32]byte


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
func M512MovmEpi16(k x86.Mmask32) x86.M512i {
	return x86.M512i(m512MovmEpi16(uint32(k)))
}

func m512MovmEpi16(k uint32) [64]byte


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
func M256MovmEpi8(k x86.Mmask32) x86.M256i {
	return x86.M256i(m256MovmEpi8(uint32(k)))
}

func m256MovmEpi8(k uint32) [32]byte


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
func M512MovmEpi8(k x86.Mmask64) x86.M512i {
	return x86.M512i(m512MovmEpi8(uint64(k)))
}

func m512MovmEpi8(k uint64) [64]byte


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
func MaskMulhiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMulhiEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMulhiEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMulhiEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMulhiEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMulhiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMulhiEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMulhiEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMulhiEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMulhiEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMulhiEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMulhiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMulhiEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMulhiEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMulhiEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMulhiEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMulhiEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MulhiEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MulhiEpi16([64]byte(a), [64]byte(b)))
}

func m512MulhiEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMulhiEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMulhiEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMulhiEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMulhiEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMulhiEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulhiEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMulhiEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMulhiEpu16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMulhiEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMulhiEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMulhiEpu16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMulhiEpu16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMulhiEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMulhiEpu16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMulhiEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMulhiEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMulhiEpu16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMulhiEpu16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MulhiEpu16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MulhiEpu16([64]byte(a), [64]byte(b)))
}

func m512MulhiEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskMulhrsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMulhrsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMulhrsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMulhrsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMulhrsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulhrsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMulhrsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMulhrsEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMulhrsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMulhrsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMulhrsEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMulhrsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMulhrsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMulhrsEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMulhrsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMulhrsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMulhrsEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMulhrsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MulhrsEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MulhrsEpi16([64]byte(a), [64]byte(b)))
}

func m512MulhrsEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMulloEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMulloEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMulloEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzMulloEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMulloEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskMulloEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMulloEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMulloEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzMulloEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMulloEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMulloEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskMulloEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMulloEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMulloEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzMulloEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMulloEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMulloEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MulloEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MulloEpi16([64]byte(a), [64]byte(b)))
}

func m512MulloEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskPacksEpi16(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskPacksEpi16([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskPacksEpi16(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzPacksEpi16(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzPacksEpi16(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzPacksEpi16(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskPacksEpi16(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskPacksEpi16([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskPacksEpi16(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzPacksEpi16(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzPacksEpi16(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzPacksEpi16(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskPacksEpi16(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskPacksEpi16([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskPacksEpi16(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzPacksEpi16(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzPacksEpi16(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzPacksEpi16(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512PacksEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512PacksEpi16([64]byte(a), [64]byte(b)))
}

func m512PacksEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskPacksEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskPacksEpi32([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskPacksEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzPacksEpi32(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzPacksEpi32(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzPacksEpi32(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskPacksEpi32(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskPacksEpi32([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskPacksEpi32(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzPacksEpi32(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzPacksEpi32(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzPacksEpi32(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskPacksEpi32(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskPacksEpi32([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskPacksEpi32(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzPacksEpi32(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzPacksEpi32(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzPacksEpi32(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512PacksEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512PacksEpi32([64]byte(a), [64]byte(b)))
}

func m512PacksEpi32(a [64]byte, b [64]byte) [64]byte


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
func MaskPackusEpi16(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskPackusEpi16([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskPackusEpi16(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzPackusEpi16(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzPackusEpi16(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzPackusEpi16(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskPackusEpi16(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskPackusEpi16([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskPackusEpi16(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzPackusEpi16(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzPackusEpi16(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzPackusEpi16(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskPackusEpi16(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskPackusEpi16([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskPackusEpi16(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzPackusEpi16(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzPackusEpi16(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzPackusEpi16(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512PackusEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512PackusEpi16([64]byte(a), [64]byte(b)))
}

func m512PackusEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskPackusEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskPackusEpi32([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskPackusEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzPackusEpi32(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzPackusEpi32(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzPackusEpi32(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskPackusEpi32(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskPackusEpi32([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskPackusEpi32(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzPackusEpi32(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzPackusEpi32(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzPackusEpi32(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskPackusEpi32(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskPackusEpi32([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskPackusEpi32(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzPackusEpi32(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzPackusEpi32(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzPackusEpi32(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512PackusEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512PackusEpi32([64]byte(a), [64]byte(b)))
}

func m512PackusEpi32(a [64]byte, b [64]byte) [64]byte


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
func MaskPermutex2varEpi16(a x86.M128i, k x86.Mmask8, idx x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskPermutex2varEpi16([16]byte(a), uint8(k), [16]byte(idx), [16]byte(b)))
}

func maskPermutex2varEpi16(a [16]byte, k uint8, idx [16]byte, b [16]byte) [16]byte


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
func Mask2Permutex2varEpi16(a x86.M128i, idx x86.M128i, k x86.Mmask8, b x86.M128i) x86.M128i {
	return x86.M128i(mask2Permutex2varEpi16([16]byte(a), [16]byte(idx), uint8(k), [16]byte(b)))
}

func mask2Permutex2varEpi16(a [16]byte, idx [16]byte, k uint8, b [16]byte) [16]byte


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
func MaskzPermutex2varEpi16(k x86.Mmask8, a x86.M128i, idx x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzPermutex2varEpi16(uint8(k), [16]byte(a), [16]byte(idx), [16]byte(b)))
}

func maskzPermutex2varEpi16(k uint8, a [16]byte, idx [16]byte, b [16]byte) [16]byte


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
func Permutex2varEpi16(a x86.M128i, idx x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(permutex2varEpi16([16]byte(a), [16]byte(idx), [16]byte(b)))
}

func permutex2varEpi16(a [16]byte, idx [16]byte, b [16]byte) [16]byte


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
func M256MaskPermutex2varEpi16(a x86.M256i, k x86.Mmask16, idx x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskPermutex2varEpi16([32]byte(a), uint16(k), [32]byte(idx), [32]byte(b)))
}

func m256MaskPermutex2varEpi16(a [32]byte, k uint16, idx [32]byte, b [32]byte) [32]byte


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
func M256Mask2Permutex2varEpi16(a x86.M256i, idx x86.M256i, k x86.Mmask16, b x86.M256i) x86.M256i {
	return x86.M256i(m256Mask2Permutex2varEpi16([32]byte(a), [32]byte(idx), uint16(k), [32]byte(b)))
}

func m256Mask2Permutex2varEpi16(a [32]byte, idx [32]byte, k uint16, b [32]byte) [32]byte


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
func M256MaskzPermutex2varEpi16(k x86.Mmask16, a x86.M256i, idx x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzPermutex2varEpi16(uint16(k), [32]byte(a), [32]byte(idx), [32]byte(b)))
}

func m256MaskzPermutex2varEpi16(k uint16, a [32]byte, idx [32]byte, b [32]byte) [32]byte


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
func M256Permutex2varEpi16(a x86.M256i, idx x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256Permutex2varEpi16([32]byte(a), [32]byte(idx), [32]byte(b)))
}

func m256Permutex2varEpi16(a [32]byte, idx [32]byte, b [32]byte) [32]byte


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
func M512MaskPermutex2varEpi16(a x86.M512i, k x86.Mmask32, idx x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskPermutex2varEpi16([64]byte(a), uint32(k), [64]byte(idx), [64]byte(b)))
}

func m512MaskPermutex2varEpi16(a [64]byte, k uint32, idx [64]byte, b [64]byte) [64]byte


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
func M512Mask2Permutex2varEpi16(a x86.M512i, idx x86.M512i, k x86.Mmask32, b x86.M512i) x86.M512i {
	return x86.M512i(m512Mask2Permutex2varEpi16([64]byte(a), [64]byte(idx), uint32(k), [64]byte(b)))
}

func m512Mask2Permutex2varEpi16(a [64]byte, idx [64]byte, k uint32, b [64]byte) [64]byte


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
func M512MaskzPermutex2varEpi16(k x86.Mmask32, a x86.M512i, idx x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzPermutex2varEpi16(uint32(k), [64]byte(a), [64]byte(idx), [64]byte(b)))
}

func m512MaskzPermutex2varEpi16(k uint32, a [64]byte, idx [64]byte, b [64]byte) [64]byte


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
func M512Permutex2varEpi16(a x86.M512i, idx x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512Permutex2varEpi16([64]byte(a), [64]byte(idx), [64]byte(b)))
}

func m512Permutex2varEpi16(a [64]byte, idx [64]byte, b [64]byte) [64]byte


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
func MaskPermutexvarEpi16(src x86.M128i, k x86.Mmask8, idx x86.M128i, a x86.M128i) x86.M128i {
	return x86.M128i(maskPermutexvarEpi16([16]byte(src), uint8(k), [16]byte(idx), [16]byte(a)))
}

func maskPermutexvarEpi16(src [16]byte, k uint8, idx [16]byte, a [16]byte) [16]byte


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
func MaskzPermutexvarEpi16(k x86.Mmask8, idx x86.M128i, a x86.M128i) x86.M128i {
	return x86.M128i(maskzPermutexvarEpi16(uint8(k), [16]byte(idx), [16]byte(a)))
}

func maskzPermutexvarEpi16(k uint8, idx [16]byte, a [16]byte) [16]byte


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
func PermutexvarEpi16(idx x86.M128i, a x86.M128i) x86.M128i {
	return x86.M128i(permutexvarEpi16([16]byte(idx), [16]byte(a)))
}

func permutexvarEpi16(idx [16]byte, a [16]byte) [16]byte


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
func M256MaskPermutexvarEpi16(src x86.M256i, k x86.Mmask16, idx x86.M256i, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskPermutexvarEpi16([32]byte(src), uint16(k), [32]byte(idx), [32]byte(a)))
}

func m256MaskPermutexvarEpi16(src [32]byte, k uint16, idx [32]byte, a [32]byte) [32]byte


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
func M256MaskzPermutexvarEpi16(k x86.Mmask16, idx x86.M256i, a x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzPermutexvarEpi16(uint16(k), [32]byte(idx), [32]byte(a)))
}

func m256MaskzPermutexvarEpi16(k uint16, idx [32]byte, a [32]byte) [32]byte


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
func M256PermutexvarEpi16(idx x86.M256i, a x86.M256i) x86.M256i {
	return x86.M256i(m256PermutexvarEpi16([32]byte(idx), [32]byte(a)))
}

func m256PermutexvarEpi16(idx [32]byte, a [32]byte) [32]byte


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
func M512MaskPermutexvarEpi16(src x86.M512i, k x86.Mmask32, idx x86.M512i, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskPermutexvarEpi16([64]byte(src), uint32(k), [64]byte(idx), [64]byte(a)))
}

func m512MaskPermutexvarEpi16(src [64]byte, k uint32, idx [64]byte, a [64]byte) [64]byte


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
func M512MaskzPermutexvarEpi16(k x86.Mmask32, idx x86.M512i, a x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzPermutexvarEpi16(uint32(k), [64]byte(idx), [64]byte(a)))
}

func m512MaskzPermutexvarEpi16(k uint32, idx [64]byte, a [64]byte) [64]byte


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
func M512PermutexvarEpi16(idx x86.M512i, a x86.M512i) x86.M512i {
	return x86.M512i(m512PermutexvarEpi16([64]byte(idx), [64]byte(a)))
}

func m512PermutexvarEpi16(idx [64]byte, a [64]byte) [64]byte


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
func M512SadEpu8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SadEpu8([64]byte(a), [64]byte(b)))
}

func m512SadEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskSet1Epi16(src x86.M128i, k x86.Mmask8, a int16) x86.M128i {
	return x86.M128i(maskSet1Epi16([16]byte(src), uint8(k), a))
}

func maskSet1Epi16(src [16]byte, k uint8, a int16) [16]byte


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
func MaskzSet1Epi16(k x86.Mmask8, a int16) x86.M128i {
	return x86.M128i(maskzSet1Epi16(uint8(k), a))
}

func maskzSet1Epi16(k uint8, a int16) [16]byte


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
func M256MaskSet1Epi16(src x86.M256i, k x86.Mmask16, a int16) x86.M256i {
	return x86.M256i(m256MaskSet1Epi16([32]byte(src), uint16(k), a))
}

func m256MaskSet1Epi16(src [32]byte, k uint16, a int16) [32]byte


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
func M256MaskzSet1Epi16(k x86.Mmask16, a int16) x86.M256i {
	return x86.M256i(m256MaskzSet1Epi16(uint16(k), a))
}

func m256MaskzSet1Epi16(k uint16, a int16) [32]byte


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
func M512MaskSet1Epi16(src x86.M512i, k x86.Mmask32, a int16) x86.M512i {
	return x86.M512i(m512MaskSet1Epi16([64]byte(src), uint32(k), a))
}

func m512MaskSet1Epi16(src [64]byte, k uint32, a int16) [64]byte


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
func M512MaskzSet1Epi16(k x86.Mmask32, a int16) x86.M512i {
	return x86.M512i(m512MaskzSet1Epi16(uint32(k), a))
}

func m512MaskzSet1Epi16(k uint32, a int16) [64]byte


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
func MaskSet1Epi8(src x86.M128i, k x86.Mmask16, a byte) x86.M128i {
	return x86.M128i(maskSet1Epi8([16]byte(src), uint16(k), a))
}

func maskSet1Epi8(src [16]byte, k uint16, a byte) [16]byte


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
func MaskzSet1Epi8(k x86.Mmask16, a byte) x86.M128i {
	return x86.M128i(maskzSet1Epi8(uint16(k), a))
}

func maskzSet1Epi8(k uint16, a byte) [16]byte


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
func M256MaskSet1Epi8(src x86.M256i, k x86.Mmask32, a byte) x86.M256i {
	return x86.M256i(m256MaskSet1Epi8([32]byte(src), uint32(k), a))
}

func m256MaskSet1Epi8(src [32]byte, k uint32, a byte) [32]byte


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
func M256MaskzSet1Epi8(k x86.Mmask32, a byte) x86.M256i {
	return x86.M256i(m256MaskzSet1Epi8(uint32(k), a))
}

func m256MaskzSet1Epi8(k uint32, a byte) [32]byte


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
func M512MaskSet1Epi8(src x86.M512i, k x86.Mmask64, a byte) x86.M512i {
	return x86.M512i(m512MaskSet1Epi8([64]byte(src), uint64(k), a))
}

func m512MaskSet1Epi8(src [64]byte, k uint64, a byte) [64]byte


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
func M512MaskzSet1Epi8(k x86.Mmask64, a byte) x86.M512i {
	return x86.M512i(m512MaskzSet1Epi8(uint64(k), a))
}

func m512MaskzSet1Epi8(k uint64, a byte) [64]byte


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
func MaskShuffleEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskShuffleEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskShuffleEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzShuffleEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzShuffleEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzShuffleEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskShuffleEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskShuffleEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskShuffleEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzShuffleEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzShuffleEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzShuffleEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskShuffleEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskShuffleEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskShuffleEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzShuffleEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzShuffleEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzShuffleEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512ShuffleEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512ShuffleEpi8([64]byte(a), [64]byte(b)))
}

func m512ShuffleEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskShufflehiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskShufflehiEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
}

func maskShufflehiEpi16(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte


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
func MaskzShufflehiEpi16(k x86.Mmask8, a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskzShufflehiEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzShufflehiEpi16(k uint8, a [16]byte, imm8 int) [16]byte


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
func M256MaskShufflehiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskShufflehiEpi16([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func m256MaskShufflehiEpi16(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte


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
func M256MaskzShufflehiEpi16(k x86.Mmask16, a x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskzShufflehiEpi16(uint16(k), [32]byte(a), imm8))
}

func m256MaskzShufflehiEpi16(k uint16, a [32]byte, imm8 int) [32]byte


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
func M512MaskShufflehiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskShufflehiEpi16([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func m512MaskShufflehiEpi16(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte


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
func M512MaskzShufflehiEpi16(k x86.Mmask32, a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskzShufflehiEpi16(uint32(k), [64]byte(a), imm8))
}

func m512MaskzShufflehiEpi16(k uint32, a [64]byte, imm8 int) [64]byte


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
func M512ShufflehiEpi16(a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512ShufflehiEpi16([64]byte(a), imm8))
}

func m512ShufflehiEpi16(a [64]byte, imm8 int) [64]byte


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
func MaskShuffleloEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskShuffleloEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
}

func maskShuffleloEpi16(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte


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
func MaskzShuffleloEpi16(k x86.Mmask8, a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskzShuffleloEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzShuffleloEpi16(k uint8, a [16]byte, imm8 int) [16]byte


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
func M256MaskShuffleloEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskShuffleloEpi16([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func m256MaskShuffleloEpi16(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte


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
func M256MaskzShuffleloEpi16(k x86.Mmask16, a x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskzShuffleloEpi16(uint16(k), [32]byte(a), imm8))
}

func m256MaskzShuffleloEpi16(k uint16, a [32]byte, imm8 int) [32]byte


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
func M512MaskShuffleloEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskShuffleloEpi16([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func m512MaskShuffleloEpi16(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte


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
func M512MaskzShuffleloEpi16(k x86.Mmask32, a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskzShuffleloEpi16(uint32(k), [64]byte(a), imm8))
}

func m512MaskzShuffleloEpi16(k uint32, a [64]byte, imm8 int) [64]byte


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
func M512ShuffleloEpi16(a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512ShuffleloEpi16([64]byte(a), imm8))
}

func m512ShuffleloEpi16(a [64]byte, imm8 int) [64]byte


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
func MaskSllEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskSllEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
}

func maskSllEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte


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
func MaskzSllEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskzSllEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSllEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


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
func M256MaskSllEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M128i) x86.M256i {
	return x86.M256i(m256MaskSllEpi16([32]byte(src), uint16(k), [32]byte(a), [16]byte(count)))
}

func m256MaskSllEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte


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
func M256MaskzSllEpi16(k x86.Mmask16, a x86.M256i, count x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzSllEpi16(uint16(k), [32]byte(a), [16]byte(count)))
}

func m256MaskzSllEpi16(k uint16, a [32]byte, count [16]byte) [32]byte


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
func M512MaskSllEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512MaskSllEpi16([64]byte(src), uint32(k), [64]byte(a), [16]byte(count)))
}

func m512MaskSllEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte


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
func M512MaskzSllEpi16(k x86.Mmask32, a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzSllEpi16(uint32(k), [64]byte(a), [16]byte(count)))
}

func m512MaskzSllEpi16(k uint32, a [64]byte, count [16]byte) [64]byte


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
func M512SllEpi16(a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512SllEpi16([64]byte(a), [16]byte(count)))
}

func m512SllEpi16(a [64]byte, count [16]byte) [64]byte


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
func MaskSlliEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 uint32) x86.M128i {
	return x86.M128i(maskSlliEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
}

func maskSlliEpi16(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte


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
func MaskzSlliEpi16(k x86.Mmask8, a x86.M128i, imm8 uint32) x86.M128i {
	return x86.M128i(maskzSlliEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzSlliEpi16(k uint8, a [16]byte, imm8 uint32) [16]byte


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
func M256MaskSlliEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 uint32) x86.M256i {
	return x86.M256i(m256MaskSlliEpi16([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func m256MaskSlliEpi16(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte


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
func M256MaskzSlliEpi16(k x86.Mmask16, a x86.M256i, imm8 uint32) x86.M256i {
	return x86.M256i(m256MaskzSlliEpi16(uint16(k), [32]byte(a), imm8))
}

func m256MaskzSlliEpi16(k uint16, a [32]byte, imm8 uint32) [32]byte


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
func M512MaskSlliEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512MaskSlliEpi16([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func m512MaskSlliEpi16(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte


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
func M512MaskzSlliEpi16(k x86.Mmask32, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512MaskzSlliEpi16(uint32(k), [64]byte(a), imm8))
}

func m512MaskzSlliEpi16(k uint32, a [64]byte, imm8 uint32) [64]byte


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
func M512SlliEpi16(a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512SlliEpi16([64]byte(a), imm8))
}

func m512SlliEpi16(a [64]byte, imm8 uint32) [64]byte


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
func MaskSllvEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskSllvEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
}

func maskSllvEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte


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
func MaskzSllvEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskzSllvEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSllvEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


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
func SllvEpi16(a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(sllvEpi16([16]byte(a), [16]byte(count)))
}

func sllvEpi16(a [16]byte, count [16]byte) [16]byte


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
func M256MaskSllvEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSllvEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(count)))
}

func m256MaskSllvEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte


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
func M256MaskzSllvEpi16(k x86.Mmask16, a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSllvEpi16(uint16(k), [32]byte(a), [32]byte(count)))
}

func m256MaskzSllvEpi16(k uint16, a [32]byte, count [32]byte) [32]byte


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
func M256SllvEpi16(a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256SllvEpi16([32]byte(a), [32]byte(count)))
}

func m256SllvEpi16(a [32]byte, count [32]byte) [32]byte


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
func M512MaskSllvEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSllvEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(count)))
}

func m512MaskSllvEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte


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
func M512MaskzSllvEpi16(k x86.Mmask32, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSllvEpi16(uint32(k), [64]byte(a), [64]byte(count)))
}

func m512MaskzSllvEpi16(k uint32, a [64]byte, count [64]byte) [64]byte


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
func M512SllvEpi16(a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512SllvEpi16([64]byte(a), [64]byte(count)))
}

func m512SllvEpi16(a [64]byte, count [64]byte) [64]byte


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
func MaskSraEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskSraEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
}

func maskSraEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte


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
func MaskzSraEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskzSraEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSraEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


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
func M256MaskSraEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M128i) x86.M256i {
	return x86.M256i(m256MaskSraEpi16([32]byte(src), uint16(k), [32]byte(a), [16]byte(count)))
}

func m256MaskSraEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte


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
func M256MaskzSraEpi16(k x86.Mmask16, a x86.M256i, count x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzSraEpi16(uint16(k), [32]byte(a), [16]byte(count)))
}

func m256MaskzSraEpi16(k uint16, a [32]byte, count [16]byte) [32]byte


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
func M512MaskSraEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512MaskSraEpi16([64]byte(src), uint32(k), [64]byte(a), [16]byte(count)))
}

func m512MaskSraEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte


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
func M512MaskzSraEpi16(k x86.Mmask32, a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzSraEpi16(uint32(k), [64]byte(a), [16]byte(count)))
}

func m512MaskzSraEpi16(k uint32, a [64]byte, count [16]byte) [64]byte


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
func M512SraEpi16(a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512SraEpi16([64]byte(a), [16]byte(count)))
}

func m512SraEpi16(a [64]byte, count [16]byte) [64]byte


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
func MaskSraiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 uint32) x86.M128i {
	return x86.M128i(maskSraiEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
}

func maskSraiEpi16(src [16]byte, k uint8, a [16]byte, imm8 uint32) [16]byte


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
func MaskzSraiEpi16(k x86.Mmask8, a x86.M128i, imm8 uint32) x86.M128i {
	return x86.M128i(maskzSraiEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzSraiEpi16(k uint8, a [16]byte, imm8 uint32) [16]byte


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
func M256MaskSraiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 uint32) x86.M256i {
	return x86.M256i(m256MaskSraiEpi16([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func m256MaskSraiEpi16(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte


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
func M256MaskzSraiEpi16(k x86.Mmask16, a x86.M256i, imm8 uint32) x86.M256i {
	return x86.M256i(m256MaskzSraiEpi16(uint16(k), [32]byte(a), imm8))
}

func m256MaskzSraiEpi16(k uint16, a [32]byte, imm8 uint32) [32]byte


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
func M512MaskSraiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512MaskSraiEpi16([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func m512MaskSraiEpi16(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte


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
func M512MaskzSraiEpi16(k x86.Mmask32, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512MaskzSraiEpi16(uint32(k), [64]byte(a), imm8))
}

func m512MaskzSraiEpi16(k uint32, a [64]byte, imm8 uint32) [64]byte


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
func M512SraiEpi16(a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512SraiEpi16([64]byte(a), imm8))
}

func m512SraiEpi16(a [64]byte, imm8 uint32) [64]byte


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
func MaskSravEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskSravEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
}

func maskSravEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte


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
func MaskzSravEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskzSravEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSravEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


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
func SravEpi16(a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(sravEpi16([16]byte(a), [16]byte(count)))
}

func sravEpi16(a [16]byte, count [16]byte) [16]byte


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
func M256MaskSravEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSravEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(count)))
}

func m256MaskSravEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte


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
func M256MaskzSravEpi16(k x86.Mmask16, a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSravEpi16(uint16(k), [32]byte(a), [32]byte(count)))
}

func m256MaskzSravEpi16(k uint16, a [32]byte, count [32]byte) [32]byte


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
func M256SravEpi16(a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256SravEpi16([32]byte(a), [32]byte(count)))
}

func m256SravEpi16(a [32]byte, count [32]byte) [32]byte


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
func M512MaskSravEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSravEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(count)))
}

func m512MaskSravEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte


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
func M512MaskzSravEpi16(k x86.Mmask32, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSravEpi16(uint32(k), [64]byte(a), [64]byte(count)))
}

func m512MaskzSravEpi16(k uint32, a [64]byte, count [64]byte) [64]byte


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
func M512SravEpi16(a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512SravEpi16([64]byte(a), [64]byte(count)))
}

func m512SravEpi16(a [64]byte, count [64]byte) [64]byte


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
func MaskSrlEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskSrlEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
}

func maskSrlEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte


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
func MaskzSrlEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskzSrlEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSrlEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


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
func M256MaskSrlEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M128i) x86.M256i {
	return x86.M256i(m256MaskSrlEpi16([32]byte(src), uint16(k), [32]byte(a), [16]byte(count)))
}

func m256MaskSrlEpi16(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte


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
func M256MaskzSrlEpi16(k x86.Mmask16, a x86.M256i, count x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzSrlEpi16(uint16(k), [32]byte(a), [16]byte(count)))
}

func m256MaskzSrlEpi16(k uint16, a [32]byte, count [16]byte) [32]byte


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
func M512MaskSrlEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512MaskSrlEpi16([64]byte(src), uint32(k), [64]byte(a), [16]byte(count)))
}

func m512MaskSrlEpi16(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte


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
func M512MaskzSrlEpi16(k x86.Mmask32, a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzSrlEpi16(uint32(k), [64]byte(a), [16]byte(count)))
}

func m512MaskzSrlEpi16(k uint32, a [64]byte, count [16]byte) [64]byte


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
func M512SrlEpi16(a x86.M512i, count x86.M128i) x86.M512i {
	return x86.M512i(m512SrlEpi16([64]byte(a), [16]byte(count)))
}

func m512SrlEpi16(a [64]byte, count [16]byte) [64]byte


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
func MaskSrliEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskSrliEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
}

func maskSrliEpi16(src [16]byte, k uint8, a [16]byte, imm8 int) [16]byte


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
func MaskzSrliEpi16(k x86.Mmask8, a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(maskzSrliEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzSrliEpi16(k uint8, a [16]byte, imm8 int) [16]byte


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
func M256MaskSrliEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskSrliEpi16([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func m256MaskSrliEpi16(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte


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
func M256MaskzSrliEpi16(k x86.Mmask16, a x86.M256i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskzSrliEpi16(uint16(k), [32]byte(a), imm8))
}

func m256MaskzSrliEpi16(k uint16, a [32]byte, imm8 int) [32]byte


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
func M512MaskSrliEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512MaskSrliEpi16([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func m512MaskSrliEpi16(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte


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
func M512MaskzSrliEpi16(k x86.Mmask32, a x86.M512i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskzSrliEpi16(uint32(k), [64]byte(a), imm8))
}

func m512MaskzSrliEpi16(k uint32, a [64]byte, imm8 int) [64]byte


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
func M512SrliEpi16(a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(m512SrliEpi16([64]byte(a), imm8))
}

func m512SrliEpi16(a [64]byte, imm8 uint32) [64]byte


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
func MaskSrlvEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskSrlvEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
}

func maskSrlvEpi16(src [16]byte, k uint8, a [16]byte, count [16]byte) [16]byte


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
func MaskzSrlvEpi16(k x86.Mmask8, a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(maskzSrlvEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSrlvEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


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
func SrlvEpi16(a x86.M128i, count x86.M128i) x86.M128i {
	return x86.M128i(srlvEpi16([16]byte(a), [16]byte(count)))
}

func srlvEpi16(a [16]byte, count [16]byte) [16]byte


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
func M256MaskSrlvEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSrlvEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(count)))
}

func m256MaskSrlvEpi16(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte


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
func M256MaskzSrlvEpi16(k x86.Mmask16, a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSrlvEpi16(uint16(k), [32]byte(a), [32]byte(count)))
}

func m256MaskzSrlvEpi16(k uint16, a [32]byte, count [32]byte) [32]byte


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
func M256SrlvEpi16(a x86.M256i, count x86.M256i) x86.M256i {
	return x86.M256i(m256SrlvEpi16([32]byte(a), [32]byte(count)))
}

func m256SrlvEpi16(a [32]byte, count [32]byte) [32]byte


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
func M512MaskSrlvEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSrlvEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(count)))
}

func m512MaskSrlvEpi16(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte


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
func M512MaskzSrlvEpi16(k x86.Mmask32, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSrlvEpi16(uint32(k), [64]byte(a), [64]byte(count)))
}

func m512MaskzSrlvEpi16(k uint32, a [64]byte, count [64]byte) [64]byte


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
func M512SrlvEpi16(a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(m512SrlvEpi16([64]byte(a), [64]byte(count)))
}

func m512SrlvEpi16(a [64]byte, count [64]byte) [64]byte


// MaskStoreuEpi16: Store packed 16-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF k[j]
//				MEM[mem_addr+i+15:mem_addr+i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm_mask_storeu_epi16'.
// Requires AVX512BW.
func MaskStoreuEpi16(mem_addr uintptr, k x86.Mmask8, a x86.M128i)  {
	maskStoreuEpi16(uintptr(mem_addr), uint8(k), [16]byte(a))
}

func maskStoreuEpi16(mem_addr uintptr, k uint8, a [16]byte) 


// M256MaskStoreuEpi16: Store packed 16-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*16
//			IF k[j]
//				MEM[mem_addr+i+15:mem_addr+i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm256_mask_storeu_epi16'.
// Requires AVX512BW.
func M256MaskStoreuEpi16(mem_addr uintptr, k x86.Mmask16, a x86.M256i)  {
	m256MaskStoreuEpi16(uintptr(mem_addr), uint16(k), [32]byte(a))
}

func m256MaskStoreuEpi16(mem_addr uintptr, k uint16, a [32]byte) 


// M512MaskStoreuEpi16: Store packed 16-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 31
//			i := j*16
//			IF k[j]
//				MEM[mem_addr+i+15:mem_addr+i] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQU16'. Intrinsic: '_mm512_mask_storeu_epi16'.
// Requires AVX512BW.
func M512MaskStoreuEpi16(mem_addr uintptr, k x86.Mmask32, a x86.M512i)  {
	m512MaskStoreuEpi16(uintptr(mem_addr), uint32(k), [64]byte(a))
}

func m512MaskStoreuEpi16(mem_addr uintptr, k uint32, a [64]byte) 


// MaskStoreuEpi8: Store packed 8-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF k[j]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm_mask_storeu_epi8'.
// Requires AVX512BW.
func MaskStoreuEpi8(mem_addr uintptr, k x86.Mmask16, a x86.M128i)  {
	maskStoreuEpi8(uintptr(mem_addr), uint16(k), [16]byte(a))
}

func maskStoreuEpi8(mem_addr uintptr, k uint16, a [16]byte) 


// M256MaskStoreuEpi8: Store packed 8-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 31
//			i := j*8
//			IF k[j]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm256_mask_storeu_epi8'.
// Requires AVX512BW.
func M256MaskStoreuEpi8(mem_addr uintptr, k x86.Mmask32, a x86.M256i)  {
	m256MaskStoreuEpi8(uintptr(mem_addr), uint32(k), [32]byte(a))
}

func m256MaskStoreuEpi8(mem_addr uintptr, k uint32, a [32]byte) 


// M512MaskStoreuEpi8: Store packed 8-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQU8'. Intrinsic: '_mm512_mask_storeu_epi8'.
// Requires AVX512BW.
func M512MaskStoreuEpi8(mem_addr uintptr, k x86.Mmask64, a x86.M512i)  {
	m512MaskStoreuEpi8(uintptr(mem_addr), uint64(k), [64]byte(a))
}

func m512MaskStoreuEpi8(mem_addr uintptr, k uint64, a [64]byte) 


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
func MaskSubEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskSubEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskSubEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzSubEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzSubEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzSubEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskSubEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSubEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskSubEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzSubEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSubEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzSubEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskSubEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSubEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskSubEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzSubEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSubEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzSubEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512SubEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SubEpi16([64]byte(a), [64]byte(b)))
}

func m512SubEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskSubEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskSubEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskSubEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzSubEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzSubEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzSubEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskSubEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSubEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskSubEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzSubEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSubEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzSubEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskSubEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSubEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskSubEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzSubEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSubEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzSubEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512SubEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SubEpi8([64]byte(a), [64]byte(b)))
}

func m512SubEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskSubsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskSubsEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzSubsEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzSubsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskSubsEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSubsEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskSubsEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzSubsEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSubsEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzSubsEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskSubsEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSubsEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskSubsEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzSubsEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSubsEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzSubsEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512SubsEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SubsEpi16([64]byte(a), [64]byte(b)))
}

func m512SubsEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskSubsEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskSubsEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzSubsEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzSubsEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskSubsEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSubsEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskSubsEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzSubsEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSubsEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzSubsEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskSubsEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSubsEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskSubsEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzSubsEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSubsEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzSubsEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512SubsEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SubsEpi8([64]byte(a), [64]byte(b)))
}

func m512SubsEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpu16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskSubsEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskSubsEpu16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzSubsEpu16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzSubsEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskSubsEpu16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSubsEpu16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskSubsEpu16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzSubsEpu16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSubsEpu16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzSubsEpu16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskSubsEpu16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSubsEpu16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskSubsEpu16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzSubsEpu16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSubsEpu16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzSubsEpu16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512SubsEpu16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SubsEpu16([64]byte(a), [64]byte(b)))
}

func m512SubsEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpu8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskSubsEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskSubsEpu8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzSubsEpu8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzSubsEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskSubsEpu8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskSubsEpu8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskSubsEpu8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzSubsEpu8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzSubsEpu8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzSubsEpu8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskSubsEpu8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskSubsEpu8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskSubsEpu8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzSubsEpu8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzSubsEpu8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzSubsEpu8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512SubsEpu8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512SubsEpu8([64]byte(a), [64]byte(b)))
}

func m512SubsEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskTestEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskTestEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskTestEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func TestEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(testEpi16Mask([16]byte(a), [16]byte(b)))
}

func testEpi16Mask(a [16]byte, b [16]byte) uint8


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
func M256MaskTestEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskTestEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskTestEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M256TestEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256TestEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256TestEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M512MaskTestEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskTestEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskTestEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func M512TestEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512TestEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512TestEpi16Mask(a [64]byte, b [64]byte) uint32


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
func MaskTestEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskTestEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskTestEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func TestEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(testEpi8Mask([16]byte(a), [16]byte(b)))
}

func testEpi8Mask(a [16]byte, b [16]byte) uint16


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
func M256MaskTestEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskTestEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskTestEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M256TestEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256TestEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256TestEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M512MaskTestEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskTestEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskTestEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func M512TestEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512TestEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512TestEpi8Mask(a [64]byte, b [64]byte) uint64


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
func MaskTestnEpi16Mask(k1 x86.Mmask8, a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(maskTestnEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskTestnEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


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
func TestnEpi16Mask(a x86.M128i, b x86.M128i) x86.Mmask8 {
	return x86.Mmask8(testnEpi16Mask([16]byte(a), [16]byte(b)))
}

func testnEpi16Mask(a [16]byte, b [16]byte) uint8


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
func M256MaskTestnEpi16Mask(k1 x86.Mmask16, a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256MaskTestnEpi16Mask(uint16(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskTestnEpi16Mask(k1 uint16, a [32]byte, b [32]byte) uint16


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
func M256TestnEpi16Mask(a x86.M256i, b x86.M256i) x86.Mmask16 {
	return x86.Mmask16(m256TestnEpi16Mask([32]byte(a), [32]byte(b)))
}

func m256TestnEpi16Mask(a [32]byte, b [32]byte) uint16


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
func M512MaskTestnEpi16Mask(k1 x86.Mmask32, a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512MaskTestnEpi16Mask(uint32(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskTestnEpi16Mask(k1 uint32, a [64]byte, b [64]byte) uint32


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
func M512TestnEpi16Mask(a x86.M512i, b x86.M512i) x86.Mmask32 {
	return x86.Mmask32(m512TestnEpi16Mask([64]byte(a), [64]byte(b)))
}

func m512TestnEpi16Mask(a [64]byte, b [64]byte) uint32


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
func MaskTestnEpi8Mask(k1 x86.Mmask16, a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(maskTestnEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskTestnEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


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
func TestnEpi8Mask(a x86.M128i, b x86.M128i) x86.Mmask16 {
	return x86.Mmask16(testnEpi8Mask([16]byte(a), [16]byte(b)))
}

func testnEpi8Mask(a [16]byte, b [16]byte) uint16


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
func M256MaskTestnEpi8Mask(k1 x86.Mmask32, a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256MaskTestnEpi8Mask(uint32(k1), [32]byte(a), [32]byte(b)))
}

func m256MaskTestnEpi8Mask(k1 uint32, a [32]byte, b [32]byte) uint32


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
func M256TestnEpi8Mask(a x86.M256i, b x86.M256i) x86.Mmask32 {
	return x86.Mmask32(m256TestnEpi8Mask([32]byte(a), [32]byte(b)))
}

func m256TestnEpi8Mask(a [32]byte, b [32]byte) uint32


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
func M512MaskTestnEpi8Mask(k1 x86.Mmask64, a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512MaskTestnEpi8Mask(uint64(k1), [64]byte(a), [64]byte(b)))
}

func m512MaskTestnEpi8Mask(k1 uint64, a [64]byte, b [64]byte) uint64


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
func M512TestnEpi8Mask(a x86.M512i, b x86.M512i) x86.Mmask64 {
	return x86.Mmask64(m512TestnEpi8Mask([64]byte(a), [64]byte(b)))
}

func m512TestnEpi8Mask(a [64]byte, b [64]byte) uint64


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
func MaskUnpackhiEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskUnpackhiEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskUnpackhiEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzUnpackhiEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzUnpackhiEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskUnpackhiEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskUnpackhiEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskUnpackhiEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzUnpackhiEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzUnpackhiEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzUnpackhiEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskUnpackhiEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskUnpackhiEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskUnpackhiEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzUnpackhiEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzUnpackhiEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzUnpackhiEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512UnpackhiEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512UnpackhiEpi16([64]byte(a), [64]byte(b)))
}

func m512UnpackhiEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskUnpackhiEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskUnpackhiEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskUnpackhiEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzUnpackhiEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzUnpackhiEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackhiEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskUnpackhiEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskUnpackhiEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskUnpackhiEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzUnpackhiEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzUnpackhiEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzUnpackhiEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskUnpackhiEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskUnpackhiEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskUnpackhiEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzUnpackhiEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzUnpackhiEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzUnpackhiEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512UnpackhiEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512UnpackhiEpi8([64]byte(a), [64]byte(b)))
}

func m512UnpackhiEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskUnpackloEpi16(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskUnpackloEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskUnpackloEpi16(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


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
func MaskzUnpackloEpi16(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzUnpackloEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


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
func M256MaskUnpackloEpi16(src x86.M256i, k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskUnpackloEpi16([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskUnpackloEpi16(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzUnpackloEpi16(k x86.Mmask16, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzUnpackloEpi16(uint16(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzUnpackloEpi16(k uint16, a [32]byte, b [32]byte) [32]byte


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
func M512MaskUnpackloEpi16(src x86.M512i, k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskUnpackloEpi16([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskUnpackloEpi16(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzUnpackloEpi16(k x86.Mmask32, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzUnpackloEpi16(uint32(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzUnpackloEpi16(k uint32, a [64]byte, b [64]byte) [64]byte


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
func M512UnpackloEpi16(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512UnpackloEpi16([64]byte(a), [64]byte(b)))
}

func m512UnpackloEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskUnpackloEpi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskUnpackloEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskUnpackloEpi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzUnpackloEpi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzUnpackloEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackloEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func M256MaskUnpackloEpi8(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskUnpackloEpi8([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskUnpackloEpi8(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


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
func M256MaskzUnpackloEpi8(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzUnpackloEpi8(uint32(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzUnpackloEpi8(k uint32, a [32]byte, b [32]byte) [32]byte


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
func M512MaskUnpackloEpi8(src x86.M512i, k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskUnpackloEpi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskUnpackloEpi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512MaskzUnpackloEpi8(k x86.Mmask64, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzUnpackloEpi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzUnpackloEpi8(k uint64, a [64]byte, b [64]byte) [64]byte


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
func M512UnpackloEpi8(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512UnpackloEpi8([64]byte(a), [64]byte(b)))
}

func m512UnpackloEpi8(a [64]byte, b [64]byte) [64]byte

