package avx512bw

import . "github.com/klauspost/intrinsics/x86"


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
func MaskAbsEpi16(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskAbsEpi16([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzAbsEpi16(k Mmask8, a M128i) M128i {
	return M128i(maskzAbsEpi16(uint8(k), [16]byte(a)))
}

func maskzAbsEpi16(k uint8, a [16]byte) [16]byte


// MaskAbsEpi161: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func MaskAbsEpi161(src M256i, k Mmask16, a M256i) M256i {
	return M256i(maskAbsEpi161([32]byte(src), uint16(k), [32]byte(a)))
}

func maskAbsEpi161(src [32]byte, k uint16, a [32]byte) [32]byte


// MaskzAbsEpi161: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func MaskzAbsEpi161(k Mmask16, a M256i) M256i {
	return M256i(maskzAbsEpi161(uint16(k), [32]byte(a)))
}

func maskzAbsEpi161(k uint16, a [32]byte) [32]byte


// AbsEpi16: Compute the absolute value of packed 16-bit integers in 'a', and
// store the unsigned results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := ABS(a[i+15:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPABSW'. Intrinsic: '_mm512_abs_epi16'.
// Requires AVX512BW.
func AbsEpi16(a M512i) M512i {
	return M512i(absEpi16([64]byte(a)))
}

func absEpi16(a [64]byte) [64]byte


// MaskAbsEpi162: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func MaskAbsEpi162(src M512i, k Mmask32, a M512i) M512i {
	return M512i(maskAbsEpi162([64]byte(src), uint32(k), [64]byte(a)))
}

func maskAbsEpi162(src [64]byte, k uint32, a [64]byte) [64]byte


// MaskzAbsEpi162: Compute the absolute value of packed 16-bit integers in 'a',
// and store the unsigned results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func MaskzAbsEpi162(k Mmask32, a M512i) M512i {
	return M512i(maskzAbsEpi162(uint32(k), [64]byte(a)))
}

func maskzAbsEpi162(k uint32, a [64]byte) [64]byte


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
func MaskAbsEpi8(src M128i, k Mmask16, a M128i) M128i {
	return M128i(maskAbsEpi8([16]byte(src), uint16(k), [16]byte(a)))
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
func MaskzAbsEpi8(k Mmask16, a M128i) M128i {
	return M128i(maskzAbsEpi8(uint16(k), [16]byte(a)))
}

func maskzAbsEpi8(k uint16, a [16]byte) [16]byte


// MaskAbsEpi81: Compute the absolute value of packed 8-bit integers in 'a',
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
func MaskAbsEpi81(src M256i, k Mmask32, a M256i) M256i {
	return M256i(maskAbsEpi81([32]byte(src), uint32(k), [32]byte(a)))
}

func maskAbsEpi81(src [32]byte, k uint32, a [32]byte) [32]byte


// MaskzAbsEpi81: Compute the absolute value of packed 8-bit integers in 'a',
// and store the unsigned results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func MaskzAbsEpi81(k Mmask32, a M256i) M256i {
	return M256i(maskzAbsEpi81(uint32(k), [32]byte(a)))
}

func maskzAbsEpi81(k uint32, a [32]byte) [32]byte


// AbsEpi8: Compute the absolute value of packed 8-bit integers in 'a', and
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
func AbsEpi8(a M512i) M512i {
	return M512i(absEpi8([64]byte(a)))
}

func absEpi8(a [64]byte) [64]byte


// MaskAbsEpi82: Compute the absolute value of packed 8-bit integers in 'a',
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
func MaskAbsEpi82(src M512i, k Mmask64, a M512i) M512i {
	return M512i(maskAbsEpi82([64]byte(src), uint64(k), [64]byte(a)))
}

func maskAbsEpi82(src [64]byte, k uint64, a [64]byte) [64]byte


// MaskzAbsEpi82: Compute the absolute value of packed 8-bit integers in 'a',
// and store the unsigned results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func MaskzAbsEpi82(k Mmask64, a M512i) M512i {
	return M512i(maskzAbsEpi82(uint64(k), [64]byte(a)))
}

func maskzAbsEpi82(k uint64, a [64]byte) [64]byte


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
func MaskAddEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskAddEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzAddEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzAddEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskAddEpi161: Add packed 16-bit integers in 'a' and 'b', and store the
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
func MaskAddEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskAddEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskAddEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzAddEpi161: Add packed 16-bit integers in 'a' and 'b', and store the
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
func MaskzAddEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzAddEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzAddEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// AddEpi16: Add packed 16-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := a[i+15:i] + b[i+15:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDW'. Intrinsic: '_mm512_add_epi16'.
// Requires AVX512BW.
func AddEpi16(a M512i, b M512i) M512i {
	return M512i(addEpi16([64]byte(a), [64]byte(b)))
}

func addEpi16(a [64]byte, b [64]byte) [64]byte


// MaskAddEpi162: Add packed 16-bit integers in 'a' and 'b', and store the
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
func MaskAddEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskAddEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskAddEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzAddEpi162: Add packed 16-bit integers in 'a' and 'b', and store the
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
func MaskzAddEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzAddEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzAddEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAddEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskAddEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzAddEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzAddEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAddEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskAddEpi81: Add packed 8-bit integers in 'a' and 'b', and store the
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
func MaskAddEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskAddEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskAddEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzAddEpi81: Add packed 8-bit integers in 'a' and 'b', and store the
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
func MaskzAddEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzAddEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzAddEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// AddEpi8: Add packed 8-bit integers in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := a[i+7:i] + b[i+7:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDB'. Intrinsic: '_mm512_add_epi8'.
// Requires AVX512BW.
func AddEpi8(a M512i, b M512i) M512i {
	return M512i(addEpi8([64]byte(a), [64]byte(b)))
}

func addEpi8(a [64]byte, b [64]byte) [64]byte


// MaskAddEpi82: Add packed 8-bit integers in 'a' and 'b', and store the
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
func MaskAddEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskAddEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskAddEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzAddEpi82: Add packed 8-bit integers in 'a' and 'b', and store the
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
func MaskzAddEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzAddEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzAddEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskAddsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzAddsEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzAddsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskAddsEpi161: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskAddsEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskAddsEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskAddsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzAddsEpi161: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzAddsEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzAddsEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzAddsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// AddsEpi16: Add packed 16-bit integers in 'a' and 'b' using saturation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := Saturate_To_Int16( a[i+15:i] + b[i+15:i] )
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSW'. Intrinsic: '_mm512_adds_epi16'.
// Requires AVX512BW.
func AddsEpi16(a M512i, b M512i) M512i {
	return M512i(addsEpi16([64]byte(a), [64]byte(b)))
}

func addsEpi16(a [64]byte, b [64]byte) [64]byte


// MaskAddsEpi162: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskAddsEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskAddsEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskAddsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzAddsEpi162: Add packed 16-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzAddsEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzAddsEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzAddsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskAddsEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzAddsEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzAddsEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskAddsEpi81: Add packed 8-bit integers in 'a' and 'b' using saturation,
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
func MaskAddsEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskAddsEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskAddsEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzAddsEpi81: Add packed 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzAddsEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzAddsEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzAddsEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// AddsEpi8: Add packed 8-bit integers in 'a' and 'b' using saturation, and
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
func AddsEpi8(a M512i, b M512i) M512i {
	return M512i(addsEpi8([64]byte(a), [64]byte(b)))
}

func addsEpi8(a [64]byte, b [64]byte) [64]byte


// MaskAddsEpi82: Add packed 8-bit integers in 'a' and 'b' using saturation,
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
func MaskAddsEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskAddsEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskAddsEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzAddsEpi82: Add packed 8-bit integers in 'a' and 'b' using saturation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzAddsEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzAddsEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzAddsEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpu16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskAddsEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzAddsEpu16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzAddsEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskAddsEpu161: Add packed unsigned 16-bit integers in 'a' and 'b' using
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
func MaskAddsEpu161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskAddsEpu161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskAddsEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzAddsEpu161: Add packed unsigned 16-bit integers in 'a' and 'b' using
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
func MaskzAddsEpu161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzAddsEpu161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzAddsEpu161(k uint16, a [32]byte, b [32]byte) [32]byte


// AddsEpu16: Add packed unsigned 16-bit integers in 'a' and 'b' using
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
func AddsEpu16(a M512i, b M512i) M512i {
	return M512i(addsEpu16([64]byte(a), [64]byte(b)))
}

func addsEpu16(a [64]byte, b [64]byte) [64]byte


// MaskAddsEpu162: Add packed unsigned 16-bit integers in 'a' and 'b' using
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
func MaskAddsEpu162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskAddsEpu162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskAddsEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzAddsEpu162: Add packed unsigned 16-bit integers in 'a' and 'b' using
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
func MaskzAddsEpu162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzAddsEpu162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzAddsEpu162(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAddsEpu8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskAddsEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzAddsEpu8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzAddsEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAddsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskAddsEpu81: Add packed unsigned 8-bit integers in 'a' and 'b' using
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
func MaskAddsEpu81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskAddsEpu81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskAddsEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzAddsEpu81: Add packed unsigned 8-bit integers in 'a' and 'b' using
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
func MaskzAddsEpu81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzAddsEpu81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzAddsEpu81(k uint32, a [32]byte, b [32]byte) [32]byte


// AddsEpu8: Add packed unsigned 8-bit integers in 'a' and 'b' using
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
func AddsEpu8(a M512i, b M512i) M512i {
	return M512i(addsEpu8([64]byte(a), [64]byte(b)))
}

func addsEpu8(a [64]byte, b [64]byte) [64]byte


// MaskAddsEpu82: Add packed unsigned 8-bit integers in 'a' and 'b' using
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
func MaskAddsEpu82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskAddsEpu82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskAddsEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzAddsEpu82: Add packed unsigned 8-bit integers in 'a' and 'b' using
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
func MaskzAddsEpu82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzAddsEpu82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzAddsEpu82(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskAlignrEpi8(src M128i, k Mmask16, a M128i, b M128i, count int) M128i {
	return M128i(maskAlignrEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b), count))
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
func MaskzAlignrEpi8(k Mmask16, a M128i, b M128i, count int) M128i {
	return M128i(maskzAlignrEpi8(uint16(k), [16]byte(a), [16]byte(b), count))
}

func maskzAlignrEpi8(k uint16, a [16]byte, b [16]byte, count int) [16]byte


// MaskAlignrEpi81: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskAlignrEpi81(src M256i, k Mmask32, a M256i, b M256i, count int) M256i {
	return M256i(maskAlignrEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b), count))
}

func maskAlignrEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte, count int) [32]byte


// MaskzAlignrEpi81: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
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
func MaskzAlignrEpi81(k Mmask32, a M256i, b M256i, count int) M256i {
	return M256i(maskzAlignrEpi81(uint32(k), [32]byte(a), [32]byte(b), count))
}

func maskzAlignrEpi81(k uint32, a [32]byte, b [32]byte, count int) [32]byte


// AlignrEpi8: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
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
func AlignrEpi8(a M512i, b M512i, count int) M512i {
	return M512i(alignrEpi8([64]byte(a), [64]byte(b), count))
}

func alignrEpi8(a [64]byte, b [64]byte, count int) [64]byte


// MaskAlignrEpi82: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskAlignrEpi82(src M512i, k Mmask64, a M512i, b M512i, count int) M512i {
	return M512i(maskAlignrEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b), count))
}

func maskAlignrEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte, count int) [64]byte


// MaskzAlignrEpi82: Concatenate pairs of 16-byte blocks in 'a' and 'b' into a
// 32-byte temporary result, shift the result right by 'count' bytes, and store
// the low 16 bytes in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
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
func MaskzAlignrEpi82(k Mmask64, a M512i, b M512i, count int) M512i {
	return M512i(maskzAlignrEpi82(uint64(k), [64]byte(a), [64]byte(b), count))
}

func maskzAlignrEpi82(k uint64, a [64]byte, b [64]byte, count int) [64]byte


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
func MaskAvgEpu16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskAvgEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzAvgEpu16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzAvgEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzAvgEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskAvgEpu161: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst' using writemask 'k' (elements are copied from
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
func MaskAvgEpu161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskAvgEpu161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskAvgEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzAvgEpu161: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
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
func MaskzAvgEpu161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzAvgEpu161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzAvgEpu161(k uint16, a [32]byte, b [32]byte) [32]byte


// AvgEpu16: Average packed unsigned 16-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGW'. Intrinsic: '_mm512_avg_epu16'.
// Requires AVX512BW.
func AvgEpu16(a M512i, b M512i) M512i {
	return M512i(avgEpu16([64]byte(a), [64]byte(b)))
}

func avgEpu16(a [64]byte, b [64]byte) [64]byte


// MaskAvgEpu162: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst' using writemask 'k' (elements are copied from
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
func MaskAvgEpu162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskAvgEpu162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskAvgEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzAvgEpu162: Average packed unsigned 16-bit integers in 'a' and 'b', and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
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
func MaskzAvgEpu162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzAvgEpu162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzAvgEpu162(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskAvgEpu8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskAvgEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzAvgEpu8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzAvgEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzAvgEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskAvgEpu81: Average packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskAvgEpu81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskAvgEpu81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskAvgEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzAvgEpu81: Average packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskzAvgEpu81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzAvgEpu81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzAvgEpu81(k uint32, a [32]byte, b [32]byte) [32]byte


// AvgEpu8: Average packed unsigned 8-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPAVGB'. Intrinsic: '_mm512_avg_epu8'.
// Requires AVX512BW.
func AvgEpu8(a M512i, b M512i) M512i {
	return M512i(avgEpu8([64]byte(a), [64]byte(b)))
}

func avgEpu8(a [64]byte, b [64]byte) [64]byte


// MaskAvgEpu82: Average packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskAvgEpu82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskAvgEpu82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskAvgEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzAvgEpu82: Average packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskzAvgEpu82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzAvgEpu82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzAvgEpu82(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskBlendEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskBlendEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskBlendEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskBlendEpi161: Blend packed 16-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
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
func MaskBlendEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskBlendEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskBlendEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskBlendEpi162: Blend packed 16-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
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
func MaskBlendEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskBlendEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskBlendEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskBlendEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskBlendEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskBlendEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskBlendEpi81: Blend packed 8-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
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
func MaskBlendEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskBlendEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskBlendEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskBlendEpi82: Blend packed 8-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
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
func MaskBlendEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskBlendEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskBlendEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


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
func MaskBroadcastbEpi8(src M128i, k Mmask16, a M128i) M128i {
	return M128i(maskBroadcastbEpi8([16]byte(src), uint16(k), [16]byte(a)))
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
func MaskzBroadcastbEpi8(k Mmask16, a M128i) M128i {
	return M128i(maskzBroadcastbEpi8(uint16(k), [16]byte(a)))
}

func maskzBroadcastbEpi8(k uint16, a [16]byte) [16]byte


// MaskBroadcastbEpi81: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
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
func MaskBroadcastbEpi81(src M256i, k Mmask32, a M128i) M256i {
	return M256i(maskBroadcastbEpi81([32]byte(src), uint32(k), [16]byte(a)))
}

func maskBroadcastbEpi81(src [32]byte, k uint32, a [16]byte) [32]byte


// MaskzBroadcastbEpi81: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
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
func MaskzBroadcastbEpi81(k Mmask32, a M128i) M256i {
	return M256i(maskzBroadcastbEpi81(uint32(k), [16]byte(a)))
}

func maskzBroadcastbEpi81(k uint32, a [16]byte) [32]byte


// BroadcastbEpi8: Broadcast the low packed 8-bit integer from 'a' to all
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
func BroadcastbEpi8(a M128i) M512i {
	return M512i(broadcastbEpi8([16]byte(a)))
}

func broadcastbEpi8(a [16]byte) [64]byte


// MaskBroadcastbEpi82: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
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
func MaskBroadcastbEpi82(src M512i, k Mmask64, a M128i) M512i {
	return M512i(maskBroadcastbEpi82([64]byte(src), uint64(k), [16]byte(a)))
}

func maskBroadcastbEpi82(src [64]byte, k uint64, a [16]byte) [64]byte


// MaskzBroadcastbEpi82: Broadcast the low packed 8-bit integer from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
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
func MaskzBroadcastbEpi82(k Mmask64, a M128i) M512i {
	return M512i(maskzBroadcastbEpi82(uint64(k), [16]byte(a)))
}

func maskzBroadcastbEpi82(k uint64, a [16]byte) [64]byte


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
func MaskBroadcastwEpi16(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskBroadcastwEpi16([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzBroadcastwEpi16(k Mmask8, a M128i) M128i {
	return M128i(maskzBroadcastwEpi16(uint8(k), [16]byte(a)))
}

func maskzBroadcastwEpi16(k uint8, a [16]byte) [16]byte


// MaskBroadcastwEpi161: Broadcast the low packed 16-bit integer from 'a' to
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
func MaskBroadcastwEpi161(src M256i, k Mmask16, a M128i) M256i {
	return M256i(maskBroadcastwEpi161([32]byte(src), uint16(k), [16]byte(a)))
}

func maskBroadcastwEpi161(src [32]byte, k uint16, a [16]byte) [32]byte


// MaskzBroadcastwEpi161: Broadcast the low packed 16-bit integer from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
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
func MaskzBroadcastwEpi161(k Mmask16, a M128i) M256i {
	return M256i(maskzBroadcastwEpi161(uint16(k), [16]byte(a)))
}

func maskzBroadcastwEpi161(k uint16, a [16]byte) [32]byte


// BroadcastwEpi16: Broadcast the low packed 16-bit integer from 'a' to all
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
func BroadcastwEpi16(a M128i) M512i {
	return M512i(broadcastwEpi16([16]byte(a)))
}

func broadcastwEpi16(a [16]byte) [64]byte


// MaskBroadcastwEpi162: Broadcast the low packed 16-bit integer from 'a' to
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
func MaskBroadcastwEpi162(src M512i, k Mmask32, a M128i) M512i {
	return M512i(maskBroadcastwEpi162([64]byte(src), uint32(k), [16]byte(a)))
}

func maskBroadcastwEpi162(src [64]byte, k uint32, a [16]byte) [64]byte


// MaskzBroadcastwEpi162: Broadcast the low packed 16-bit integer from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
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
// Instruction: 'VPBROADCASTW'. Intrinsic: '_mm512_maskz_broadcastw_epi16'.
// Requires AVX512BW.
func MaskzBroadcastwEpi162(k Mmask32, a M128i) M512i {
	return M512i(maskzBroadcastwEpi162(uint32(k), [16]byte(a)))
}

func maskzBroadcastwEpi162(k uint32, a [16]byte) [64]byte


// BslliEpi128: Shift 128-bit lanes in 'a' left by 'imm8' bytes while shifting
// in zeros, and store the results in 'dst'. 
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
func BslliEpi128(a M512i, imm8 int) M512i {
	return M512i(bslliEpi128([64]byte(a), imm8))
}

func bslliEpi128(a [64]byte, imm8 int) [64]byte


// BsrliEpi128: Shift 128-bit lanes in 'a' right by 'imm8' bytes while shifting
// in zeros, and store the results in 'dst'. 
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
func BsrliEpi128(a M512i, imm8 int) M512i {
	return M512i(bsrliEpi128([64]byte(a), imm8))
}

func bsrliEpi128(a [64]byte, imm8 int) [64]byte


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
func CmpEpi16Mask(a M128i, b M128i, imm8 int) Mmask8 {
	return Mmask8(cmpEpi16Mask([16]byte(a), [16]byte(b), imm8))
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
func MaskCmpEpi16Mask(k1 Mmask8, a M128i, b M128i, imm8 int) Mmask8 {
	return Mmask8(maskCmpEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpi16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 int) uint8


// CmpEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' based on the
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
func CmpEpi16Mask1(a M256i, b M256i, imm8 int) Mmask16 {
	return Mmask16(cmpEpi16Mask1([32]byte(a), [32]byte(b), imm8))
}

func cmpEpi16Mask1(a [32]byte, b [32]byte, imm8 int) uint16


// MaskCmpEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' based on
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
func MaskCmpEpi16Mask1(k1 Mmask16, a M256i, b M256i, imm8 int) Mmask16 {
	return Mmask16(maskCmpEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b), imm8))
}

func maskCmpEpi16Mask1(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16


// CmpEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' based on the
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
func CmpEpi16Mask2(a M512i, b M512i, imm8 int) Mmask32 {
	return Mmask32(cmpEpi16Mask2([64]byte(a), [64]byte(b), imm8))
}

func cmpEpi16Mask2(a [64]byte, b [64]byte, imm8 int) uint32


// MaskCmpEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' based on
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
func MaskCmpEpi16Mask2(k1 Mmask32, a M512i, b M512i, imm8 int) Mmask32 {
	return Mmask32(maskCmpEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b), imm8))
}

func maskCmpEpi16Mask2(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32


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
func CmpEpi8Mask(a M128i, b M128i, imm8 int) Mmask16 {
	return Mmask16(cmpEpi8Mask([16]byte(a), [16]byte(b), imm8))
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
func MaskCmpEpi8Mask(k1 Mmask16, a M128i, b M128i, imm8 int) Mmask16 {
	return Mmask16(maskCmpEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpi8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 int) uint16


// CmpEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' based on the
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
func CmpEpi8Mask1(a M256i, b M256i, imm8 int) Mmask32 {
	return Mmask32(cmpEpi8Mask1([32]byte(a), [32]byte(b), imm8))
}

func cmpEpi8Mask1(a [32]byte, b [32]byte, imm8 int) uint32


// MaskCmpEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' based on the
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
func MaskCmpEpi8Mask1(k1 Mmask32, a M256i, b M256i, imm8 int) Mmask32 {
	return Mmask32(maskCmpEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b), imm8))
}

func maskCmpEpi8Mask1(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32


// CmpEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' based on the
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
func CmpEpi8Mask2(a M512i, b M512i, imm8 int) Mmask64 {
	return Mmask64(cmpEpi8Mask2([64]byte(a), [64]byte(b), imm8))
}

func cmpEpi8Mask2(a [64]byte, b [64]byte, imm8 int) uint64


// MaskCmpEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' based on the
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
func MaskCmpEpi8Mask2(k1 Mmask64, a M512i, b M512i, imm8 int) Mmask64 {
	return Mmask64(maskCmpEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b), imm8))
}

func maskCmpEpi8Mask2(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64


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
func CmpEpu16Mask(a M128i, b M128i, imm8 int) Mmask8 {
	return Mmask8(cmpEpu16Mask([16]byte(a), [16]byte(b), imm8))
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
func MaskCmpEpu16Mask(k1 Mmask8, a M128i, b M128i, imm8 int) Mmask8 {
	return Mmask8(maskCmpEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpu16Mask(k1 uint8, a [16]byte, b [16]byte, imm8 int) uint8


// CmpEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' based
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
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmp_epu16_mask'.
// Requires AVX512BW.
func CmpEpu16Mask1(a M256i, b M256i, imm8 int) Mmask16 {
	return Mmask16(cmpEpu16Mask1([32]byte(a), [32]byte(b), imm8))
}

func cmpEpu16Mask1(a [32]byte, b [32]byte, imm8 int) uint16


// MaskCmpEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
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
func MaskCmpEpu16Mask1(k1 Mmask16, a M256i, b M256i, imm8 int) Mmask16 {
	return Mmask16(maskCmpEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b), imm8))
}

func maskCmpEpu16Mask1(k1 uint16, a [32]byte, b [32]byte, imm8 int) uint16


// CmpEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' based
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
//			i := j*16
//			k[j] := ( a[i+15:i] OP b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmp_epu16_mask'.
// Requires AVX512BW.
func CmpEpu16Mask2(a M512i, b M512i, imm8 int) Mmask32 {
	return Mmask32(cmpEpu16Mask2([64]byte(a), [64]byte(b), imm8))
}

func cmpEpu16Mask2(a [64]byte, b [64]byte, imm8 int) uint32


// MaskCmpEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
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
func MaskCmpEpu16Mask2(k1 Mmask32, a M512i, b M512i, imm8 int) Mmask32 {
	return Mmask32(maskCmpEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b), imm8))
}

func maskCmpEpu16Mask2(k1 uint32, a [64]byte, b [64]byte, imm8 int) uint32


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
func CmpEpu8Mask(a M128i, b M128i, imm8 int) Mmask16 {
	return Mmask16(cmpEpu8Mask([16]byte(a), [16]byte(b), imm8))
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
func MaskCmpEpu8Mask(k1 Mmask16, a M128i, b M128i, imm8 int) Mmask16 {
	return Mmask16(maskCmpEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b), imm8))
}

func maskCmpEpu8Mask(k1 uint16, a [16]byte, b [16]byte, imm8 int) uint16


// CmpEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' based on
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
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmp_epu8_mask'.
// Requires AVX512BW.
func CmpEpu8Mask1(a M256i, b M256i, imm8 int) Mmask32 {
	return Mmask32(cmpEpu8Mask1([32]byte(a), [32]byte(b), imm8))
}

func cmpEpu8Mask1(a [32]byte, b [32]byte, imm8 int) uint32


// MaskCmpEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpEpu8Mask1(k1 Mmask32, a M256i, b M256i, imm8 int) Mmask32 {
	return Mmask32(maskCmpEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b), imm8))
}

func maskCmpEpu8Mask1(k1 uint32, a [32]byte, b [32]byte, imm8 int) uint32


// CmpEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' based on
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
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] OP b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmp_epu8_mask'.
// Requires AVX512BW.
func CmpEpu8Mask2(a M512i, b M512i, imm8 int) Mmask64 {
	return Mmask64(cmpEpu8Mask2([64]byte(a), [64]byte(b), imm8))
}

func cmpEpu8Mask2(a [64]byte, b [64]byte, imm8 int) uint64


// MaskCmpEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpEpu8Mask2(k1 Mmask64, a M512i, b M512i, imm8 int) Mmask64 {
	return Mmask64(maskCmpEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b), imm8))
}

func maskCmpEpu8Mask2(k1 uint64, a [64]byte, b [64]byte, imm8 int) uint64


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
func CmpeqEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpeqEpi16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpeqEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpeqEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpeqEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm256_cmpeq_epi16_mask'.
// Requires AVX512BW.
func CmpeqEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpeqEpi16Mask1([32]byte(a), [32]byte(b)))
}

func cmpeqEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpeqEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpeqEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpeqEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpeqEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpeqEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPW'. Intrinsic: '_mm512_cmpeq_epi16_mask'.
// Requires AVX512BW.
func CmpeqEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpeqEpi16Mask2([64]byte(a), [64]byte(b)))
}

func cmpeqEpi16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpeqEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpeqEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpeqEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpeqEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpeqEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpeqEpi8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpeqEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpeqEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpeqEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmpeq_epi8_mask'.
// Requires AVX512BW.
func CmpeqEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpeqEpi8Mask1([32]byte(a), [32]byte(b)))
}

func cmpeqEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpeqEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpeqEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpeqEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpeqEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpeqEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] == b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmpeq_epi8_mask'.
// Requires AVX512BW.
func CmpeqEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpeqEpi8Mask2([64]byte(a), [64]byte(b)))
}

func cmpeqEpi8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpeqEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpeqEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpeqEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpeqEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpeqEpu16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpeqEpu16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpeqEpu16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpeqEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpeqEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpeq_epu16_mask'.
// Requires AVX512BW.
func CmpeqEpu16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpeqEpu16Mask1([32]byte(a), [32]byte(b)))
}

func cmpeqEpu16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpeqEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k1' using zeromask 'k'
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpeq_epu16_mask'.
// Requires AVX512BW.
func MaskCmpeqEpu16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpeqEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpeqEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpeqEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] == b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpeq_epu16_mask'.
// Requires AVX512BW.
func CmpeqEpu16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpeqEpu16Mask2([64]byte(a), [64]byte(b)))
}

func cmpeqEpu16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpeqEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k1' using zeromask 'k'
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpeq_epu16_mask'.
// Requires AVX512BW.
func MaskCmpeqEpu16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpeqEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpeqEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpeqEpu8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpeqEpu8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpeqEpu8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpeqEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpeqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpeqEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpeqEpu8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpeqEpu8Mask1([32]byte(a), [32]byte(b)))
}

func cmpeqEpu8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpeqEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpeqEpu8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpeqEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpeqEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpeqEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpeqEpu8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpeqEpu8Mask2([64]byte(a), [64]byte(b)))
}

func cmpeqEpu8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpeqEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpeqEpu8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpeqEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpeqEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgeEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpgeEpi16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgeEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpgeEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpgeEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpgeEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpgeEpi16Mask1([32]byte(a), [32]byte(b)))
}

func cmpgeEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpgeEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpgeEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpgeEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgeEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpgeEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpgeEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpgeEpi16Mask2([64]byte(a), [64]byte(b)))
}

func cmpgeEpi16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpgeEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpgeEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpgeEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgeEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgeEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpgeEpi8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgeEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpgeEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpgeEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func CmpgeEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpgeEpi8Mask1([32]byte(a), [32]byte(b)))
}

func cmpgeEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpgeEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpgeEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpgeEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgeEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpgeEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func CmpgeEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpgeEpi8Mask2([64]byte(a), [64]byte(b)))
}

func cmpgeEpi8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpgeEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpgeEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpgeEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgeEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgeEpu16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpgeEpu16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgeEpu16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpgeEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpgeEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpge_epu16_mask'.
// Requires AVX512BW.
func CmpgeEpu16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpgeEpu16Mask1([32]byte(a), [32]byte(b)))
}

func cmpgeEpu16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpgeEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k1' using
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpge_epu16_mask'.
// Requires AVX512BW.
func MaskCmpgeEpu16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpgeEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgeEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpgeEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] >= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpge_epu16_mask'.
// Requires AVX512BW.
func CmpgeEpu16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpgeEpu16Mask2([64]byte(a), [64]byte(b)))
}

func cmpgeEpu16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpgeEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k1' using
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpge_epu16_mask'.
// Requires AVX512BW.
func MaskCmpgeEpu16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpgeEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgeEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgeEpu8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpgeEpu8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgeEpu8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpgeEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgeEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpgeEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpgeEpu8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpgeEpu8Mask1([32]byte(a), [32]byte(b)))
}

func cmpgeEpu8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpgeEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpgeEpu8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpgeEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgeEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpgeEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpgeEpu8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpgeEpu8Mask2([64]byte(a), [64]byte(b)))
}

func cmpgeEpu8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpgeEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpgeEpu8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpgeEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgeEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgtEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpgtEpi16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgtEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpgtEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpgtEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpgtEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpgtEpi16Mask1([32]byte(a), [32]byte(b)))
}

func cmpgtEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpgtEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpgtEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpgtEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgtEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpgtEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpgtEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpgtEpi16Mask2([64]byte(a), [64]byte(b)))
}

func cmpgtEpi16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpgtEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpgtEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpgtEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgtEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgtEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpgtEpi8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgtEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpgtEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpgtEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func CmpgtEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpgtEpi8Mask1([32]byte(a), [32]byte(b)))
}

func cmpgtEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpgtEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpgtEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpgtEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgtEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpgtEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func CmpgtEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpgtEpi8Mask2([64]byte(a), [64]byte(b)))
}

func cmpgtEpi8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpgtEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpgtEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpgtEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgtEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpgtEpu16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpgtEpu16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgtEpu16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpgtEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpgtEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpgt_epu16_mask'.
// Requires AVX512BW.
func CmpgtEpu16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpgtEpu16Mask1([32]byte(a), [32]byte(b)))
}

func cmpgtEpu16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpgtEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskCmpgtEpu16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpgtEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgtEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpgtEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] > b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpgt_epu16_mask'.
// Requires AVX512BW.
func CmpgtEpu16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpgtEpu16Mask2([64]byte(a), [64]byte(b)))
}

func cmpgtEpu16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpgtEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskCmpgtEpu16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpgtEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgtEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpgtEpu8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpgtEpu8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpgtEpu8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpgtEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpgtEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpgtEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpgtEpu8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpgtEpu8Mask1([32]byte(a), [32]byte(b)))
}

func cmpgtEpu8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpgtEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpgtEpu8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpgtEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpgtEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpgtEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpgtEpu8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpgtEpu8Mask2([64]byte(a), [64]byte(b)))
}

func cmpgtEpu8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpgtEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpgtEpu8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpgtEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgtEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpleEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpleEpi16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpleEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpleEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpleEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpleEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpleEpi16Mask1([32]byte(a), [32]byte(b)))
}

func cmpleEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpleEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpleEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpleEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpleEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpleEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpleEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpleEpi16Mask2([64]byte(a), [64]byte(b)))
}

func cmpleEpi16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpleEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpleEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpleEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpleEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpleEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpleEpi8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpleEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpleEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpleEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func CmpleEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpleEpi8Mask1([32]byte(a), [32]byte(b)))
}

func cmpleEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpleEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpleEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpleEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpleEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpleEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func CmpleEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpleEpi8Mask2([64]byte(a), [64]byte(b)))
}

func cmpleEpi8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpleEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpleEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpleEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpleEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpleEpu16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpleEpu16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpleEpu16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpleEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpleEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmple_epu16_mask'.
// Requires AVX512BW.
func CmpleEpu16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpleEpu16Mask1([32]byte(a), [32]byte(b)))
}

func cmpleEpu16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpleEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k1' using
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
func MaskCmpleEpu16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpleEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpleEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpleEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] <= b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmple_epu16_mask'.
// Requires AVX512BW.
func CmpleEpu16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpleEpu16Mask2([64]byte(a), [64]byte(b)))
}

func cmpleEpu16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpleEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k1' using
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
func MaskCmpleEpu16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpleEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpleEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpleEpu8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpleEpu8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpleEpu8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpleEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpleEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpleEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpleEpu8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpleEpu8Mask1([32]byte(a), [32]byte(b)))
}

func cmpleEpu8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpleEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpleEpu8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpleEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpleEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpleEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpleEpu8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpleEpu8Mask2([64]byte(a), [64]byte(b)))
}

func cmpleEpu8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpleEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpleEpu8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpleEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpleEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpltEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpltEpi16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpltEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpltEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpltEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpltEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpltEpi16Mask1([32]byte(a), [32]byte(b)))
}

func cmpltEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpltEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpltEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpltEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpltEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpltEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpltEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpltEpi16Mask2([64]byte(a), [64]byte(b)))
}

func cmpltEpi16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpltEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpltEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpltEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpltEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpltEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpltEpi8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpltEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpltEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpltEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for less-than,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmplt_epi8_mask'.
// Requires AVX512BW.
func CmpltEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpltEpi8Mask1([32]byte(a), [32]byte(b)))
}

func cmpltEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpltEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpltEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpltEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpltEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpltEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for less-than,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] < b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmplt_epi8_mask'.
// Requires AVX512BW.
func CmpltEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpltEpi8Mask2([64]byte(a), [64]byte(b)))
}

func cmpltEpi8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpltEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpltEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpltEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpltEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpltEpu16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpltEpu16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpltEpu16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpltEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpltEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmplt_epu16_mask'.
// Requires AVX512BW.
func CmpltEpu16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpltEpu16Mask1([32]byte(a), [32]byte(b)))
}

func cmpltEpu16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpltEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k1' using zeromask 'k'
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmplt_epu16_mask'.
// Requires AVX512BW.
func MaskCmpltEpu16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpltEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpltEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpltEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] < b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmplt_epu16_mask'.
// Requires AVX512BW.
func CmpltEpu16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpltEpu16Mask2([64]byte(a), [64]byte(b)))
}

func cmpltEpu16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpltEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k1' using zeromask 'k'
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmplt_epu16_mask'.
// Requires AVX512BW.
func MaskCmpltEpu16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpltEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpltEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpltEpu8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpltEpu8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpltEpu8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpltEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpltEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpltEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpltEpu8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpltEpu8Mask1([32]byte(a), [32]byte(b)))
}

func cmpltEpu8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpltEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpltEpu8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpltEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpltEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpltEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' for
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
func CmpltEpu8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpltEpu8Mask2([64]byte(a), [64]byte(b)))
}

func cmpltEpu8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpltEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
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
func MaskCmpltEpu8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpltEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpltEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpneqEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpneqEpi16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpneqEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpneqEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpi16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpneqEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpneqEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpneqEpi16Mask1([32]byte(a), [32]byte(b)))
}

func cmpneqEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpneqEpi16Mask1: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpneqEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpneqEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpneqEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpneqEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func CmpneqEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpneqEpi16Mask2([64]byte(a), [64]byte(b)))
}

func cmpneqEpi16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpneqEpi16Mask2: Compare packed 16-bit integers in 'a' and 'b' for
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
func MaskCmpneqEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpneqEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpneqEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpneqEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpneqEpi8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpneqEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpneqEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpi8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpneqEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for not-equal,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm256_cmpneq_epi8_mask'.
// Requires AVX512BW.
func CmpneqEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpneqEpi8Mask1([32]byte(a), [32]byte(b)))
}

func cmpneqEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpneqEpi8Mask1: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpneqEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpneqEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpneqEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpneqEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for not-equal,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPB'. Intrinsic: '_mm512_cmpneq_epi8_mask'.
// Requires AVX512BW.
func CmpneqEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpneqEpi8Mask2([64]byte(a), [64]byte(b)))
}

func cmpneqEpi8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpneqEpi8Mask2: Compare packed 8-bit integers in 'a' and 'b' for
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
func MaskCmpneqEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpneqEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpneqEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func CmpneqEpu16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(cmpneqEpu16Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpneqEpu16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskCmpneqEpu16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpu16Mask(k1 uint8, a [16]byte, b [16]byte) uint8


// CmpneqEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_cmpneq_epu16_mask'.
// Requires AVX512BW.
func CmpneqEpu16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(cmpneqEpu16Mask1([32]byte(a), [32]byte(b)))
}

func cmpneqEpu16Mask1(a [32]byte, b [32]byte) uint16


// MaskCmpneqEpu16Mask1: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k1' using zeromask 'k'
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm256_mask_cmpneq_epu16_mask'.
// Requires AVX512BW.
func MaskCmpneqEpu16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskCmpneqEpu16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpneqEpu16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// CmpneqEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ( a[i+15:i] != b[i+15:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_cmpneq_epu16_mask'.
// Requires AVX512BW.
func CmpneqEpu16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(cmpneqEpu16Mask2([64]byte(a), [64]byte(b)))
}

func cmpneqEpu16Mask2(a [64]byte, b [64]byte) uint32


// MaskCmpneqEpu16Mask2: Compare packed unsigned 16-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k1' using zeromask 'k'
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
// Instruction: 'VPCMPUW'. Intrinsic: '_mm512_mask_cmpneq_epu16_mask'.
// Requires AVX512BW.
func MaskCmpneqEpu16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskCmpneqEpu16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpneqEpu16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


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
func CmpneqEpu8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(cmpneqEpu8Mask([16]byte(a), [16]byte(b)))
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
func MaskCmpneqEpu8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskCmpneqEpu8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
}

func maskCmpneqEpu8Mask(k1 uint16, a [16]byte, b [16]byte) uint16


// CmpneqEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_cmpneq_epu8_mask'.
// Requires AVX512BW.
func CmpneqEpu8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(cmpneqEpu8Mask1([32]byte(a), [32]byte(b)))
}

func cmpneqEpu8Mask1(a [32]byte, b [32]byte) uint32


// MaskCmpneqEpu8Mask1: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k' using zeromask 'k1'
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
// Instruction: 'VPCMPUB'. Intrinsic: '_mm256_mask_cmpneq_epu8_mask'.
// Requires AVX512BW.
func MaskCmpneqEpu8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskCmpneqEpu8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskCmpneqEpu8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// CmpneqEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ( a[i+7:i] != b[i+7:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_cmpneq_epu8_mask'.
// Requires AVX512BW.
func CmpneqEpu8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(cmpneqEpu8Mask2([64]byte(a), [64]byte(b)))
}

func cmpneqEpu8Mask2(a [64]byte, b [64]byte) uint64


// MaskCmpneqEpu8Mask2: Compare packed unsigned 8-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k' using zeromask 'k1'
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
// Instruction: 'VPCMPUB'. Intrinsic: '_mm512_mask_cmpneq_epu8_mask'.
// Requires AVX512BW.
func MaskCmpneqEpu8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskCmpneqEpu8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpneqEpu8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


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
func Cvtepi16Epi8(a M128i) M128i {
	return M128i(cvtepi16Epi8([16]byte(a)))
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
func MaskCvtepi16Epi8(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskCvtepi16Epi8([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepi16Epi8(k Mmask8, a M128i) M128i {
	return M128i(maskzCvtepi16Epi8(uint8(k), [16]byte(a)))
}

func maskzCvtepi16Epi8(k uint8, a [16]byte) [16]byte


// Cvtepi16Epi81: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func Cvtepi16Epi81(a M256i) M128i {
	return M128i(cvtepi16Epi81([32]byte(a)))
}

func cvtepi16Epi81(a [32]byte) [16]byte


// MaskCvtepi16Epi81: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func MaskCvtepi16Epi81(src M128i, k Mmask16, a M256i) M128i {
	return M128i(maskCvtepi16Epi81([16]byte(src), uint16(k), [32]byte(a)))
}

func maskCvtepi16Epi81(src [16]byte, k uint16, a [32]byte) [16]byte


// MaskzCvtepi16Epi81: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func MaskzCvtepi16Epi81(k Mmask16, a M256i) M128i {
	return M128i(maskzCvtepi16Epi81(uint16(k), [32]byte(a)))
}

func maskzCvtepi16Epi81(k uint16, a [32]byte) [16]byte


// Cvtepi16Epi82: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func Cvtepi16Epi82(a M512i) M256i {
	return M256i(cvtepi16Epi82([64]byte(a)))
}

func cvtepi16Epi82(a [64]byte) [32]byte


// MaskCvtepi16Epi82: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func MaskCvtepi16Epi82(src M256i, k Mmask32, a M512i) M256i {
	return M256i(maskCvtepi16Epi82([32]byte(src), uint32(k), [64]byte(a)))
}

func maskCvtepi16Epi82(src [32]byte, k uint32, a [64]byte) [32]byte


// MaskzCvtepi16Epi82: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func MaskzCvtepi16Epi82(k Mmask32, a M512i) M256i {
	return M256i(maskzCvtepi16Epi82(uint32(k), [64]byte(a)))
}

func maskzCvtepi16Epi82(k uint32, a [64]byte) [32]byte


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
func MaskCvtepi16StoreuEpi8(base_addr uintptr, k Mmask8, a M128i)  {
	maskCvtepi16StoreuEpi8(uintptr(base_addr), uint8(k), [16]byte(a))
}

func maskCvtepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 


// MaskCvtepi16StoreuEpi81: Convert packed 16-bit integers in 'a' to packed
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
func MaskCvtepi16StoreuEpi81(base_addr uintptr, k Mmask16, a M256i)  {
	maskCvtepi16StoreuEpi81(uintptr(base_addr), uint16(k), [32]byte(a))
}

func maskCvtepi16StoreuEpi81(base_addr uintptr, k uint16, a [32]byte) 


// MaskCvtepi16StoreuEpi82: Convert packed 16-bit integers in 'a' to packed
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
func MaskCvtepi16StoreuEpi82(base_addr uintptr, k Mmask32, a M512i)  {
	maskCvtepi16StoreuEpi82(uintptr(base_addr), uint32(k), [64]byte(a))
}

func maskCvtepi16StoreuEpi82(base_addr uintptr, k uint32, a [64]byte) 


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
func MaskCvtepi8Epi16(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskCvtepi8Epi16([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepi8Epi16(k Mmask8, a M128i) M128i {
	return M128i(maskzCvtepi8Epi16(uint8(k), [16]byte(a)))
}

func maskzCvtepi8Epi16(k uint8, a [16]byte) [16]byte


// MaskCvtepi8Epi161: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func MaskCvtepi8Epi161(src M256i, k Mmask16, a M128i) M256i {
	return M256i(maskCvtepi8Epi161([32]byte(src), uint16(k), [16]byte(a)))
}

func maskCvtepi8Epi161(src [32]byte, k uint16, a [16]byte) [32]byte


// MaskzCvtepi8Epi161: Sign extend packed 8-bit integers in 'a' to packed
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
func MaskzCvtepi8Epi161(k Mmask16, a M128i) M256i {
	return M256i(maskzCvtepi8Epi161(uint16(k), [16]byte(a)))
}

func maskzCvtepi8Epi161(k uint16, a [16]byte) [32]byte


// Cvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed 16-bit
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
func Cvtepi8Epi16(a M256i) M512i {
	return M512i(cvtepi8Epi16([32]byte(a)))
}

func cvtepi8Epi16(a [32]byte) [64]byte


// MaskCvtepi8Epi162: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func MaskCvtepi8Epi162(src M512i, k Mmask32, a M256i) M512i {
	return M512i(maskCvtepi8Epi162([64]byte(src), uint32(k), [32]byte(a)))
}

func maskCvtepi8Epi162(src [64]byte, k uint32, a [32]byte) [64]byte


// MaskzCvtepi8Epi162: Sign extend packed 8-bit integers in 'a' to packed
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
func MaskzCvtepi8Epi162(k Mmask32, a M256i) M512i {
	return M512i(maskzCvtepi8Epi162(uint32(k), [32]byte(a)))
}

func maskzCvtepi8Epi162(k uint32, a [32]byte) [64]byte


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
func MaskCvtepu8Epi16(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskCvtepu8Epi16([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepu8Epi16(k Mmask8, a M128i) M128i {
	return M128i(maskzCvtepu8Epi16(uint8(k), [16]byte(a)))
}

func maskzCvtepu8Epi16(k uint8, a [16]byte) [16]byte


// MaskCvtepu8Epi161: Zero extend packed unsigned 8-bit integers in 'a' to
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
func MaskCvtepu8Epi161(src M256i, k Mmask16, a M128i) M256i {
	return M256i(maskCvtepu8Epi161([32]byte(src), uint16(k), [16]byte(a)))
}

func maskCvtepu8Epi161(src [32]byte, k uint16, a [16]byte) [32]byte


// MaskzCvtepu8Epi161: Zero extend packed unsigned 8-bit integers in 'a' to
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
func MaskzCvtepu8Epi161(k Mmask16, a M128i) M256i {
	return M256i(maskzCvtepu8Epi161(uint16(k), [16]byte(a)))
}

func maskzCvtepu8Epi161(k uint16, a [16]byte) [32]byte


// Cvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to packed
// 16-bit integers, and store the results in 'dst'. 
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
func Cvtepu8Epi16(a M256i) M512i {
	return M512i(cvtepu8Epi16([32]byte(a)))
}

func cvtepu8Epi16(a [32]byte) [64]byte


// MaskCvtepu8Epi162: Zero extend packed unsigned 8-bit integers in 'a' to
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
func MaskCvtepu8Epi162(src M512i, k Mmask32, a M256i) M512i {
	return M512i(maskCvtepu8Epi162([64]byte(src), uint32(k), [32]byte(a)))
}

func maskCvtepu8Epi162(src [64]byte, k uint32, a [32]byte) [64]byte


// MaskzCvtepu8Epi162: Zero extend packed unsigned 8-bit integers in 'a' to
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
func MaskzCvtepu8Epi162(k Mmask32, a M256i) M512i {
	return M512i(maskzCvtepu8Epi162(uint32(k), [32]byte(a)))
}

func maskzCvtepu8Epi162(k uint32, a [32]byte) [64]byte


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
func Cvtsepi16Epi8(a M128i) M128i {
	return M128i(cvtsepi16Epi8([16]byte(a)))
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
func MaskCvtsepi16Epi8(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskCvtsepi16Epi8([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzCvtsepi16Epi8(k Mmask8, a M128i) M128i {
	return M128i(maskzCvtsepi16Epi8(uint8(k), [16]byte(a)))
}

func maskzCvtsepi16Epi8(k uint8, a [16]byte) [16]byte


// Cvtsepi16Epi81: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func Cvtsepi16Epi81(a M256i) M128i {
	return M128i(cvtsepi16Epi81([32]byte(a)))
}

func cvtsepi16Epi81(a [32]byte) [16]byte


// MaskCvtsepi16Epi81: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func MaskCvtsepi16Epi81(src M128i, k Mmask16, a M256i) M128i {
	return M128i(maskCvtsepi16Epi81([16]byte(src), uint16(k), [32]byte(a)))
}

func maskCvtsepi16Epi81(src [16]byte, k uint16, a [32]byte) [16]byte


// MaskzCvtsepi16Epi81: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst' using
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
func MaskzCvtsepi16Epi81(k Mmask16, a M256i) M128i {
	return M128i(maskzCvtsepi16Epi81(uint16(k), [32]byte(a)))
}

func maskzCvtsepi16Epi81(k uint16, a [32]byte) [16]byte


// Cvtsepi16Epi82: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func Cvtsepi16Epi82(a M512i) M256i {
	return M256i(cvtsepi16Epi82([64]byte(a)))
}

func cvtsepi16Epi82(a [64]byte) [32]byte


// MaskCvtsepi16Epi82: Convert packed 16-bit integers in 'a' to packed 8-bit
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
func MaskCvtsepi16Epi82(src M256i, k Mmask32, a M512i) M256i {
	return M256i(maskCvtsepi16Epi82([32]byte(src), uint32(k), [64]byte(a)))
}

func maskCvtsepi16Epi82(src [32]byte, k uint32, a [64]byte) [32]byte


// MaskzCvtsepi16Epi82: Convert packed 16-bit integers in 'a' to packed 8-bit
// integers with signed saturation, and store the results in 'dst' using
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
func MaskzCvtsepi16Epi82(k Mmask32, a M512i) M256i {
	return M256i(maskzCvtsepi16Epi82(uint32(k), [64]byte(a)))
}

func maskzCvtsepi16Epi82(k uint32, a [64]byte) [32]byte


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
func MaskCvtsepi16StoreuEpi8(base_addr uintptr, k Mmask8, a M128i)  {
	maskCvtsepi16StoreuEpi8(uintptr(base_addr), uint8(k), [16]byte(a))
}

func maskCvtsepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 


// MaskCvtsepi16StoreuEpi81: Convert packed 16-bit integers in 'a' to packed
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
func MaskCvtsepi16StoreuEpi81(base_addr uintptr, k Mmask16, a M256i)  {
	maskCvtsepi16StoreuEpi81(uintptr(base_addr), uint16(k), [32]byte(a))
}

func maskCvtsepi16StoreuEpi81(base_addr uintptr, k uint16, a [32]byte) 


// MaskCvtsepi16StoreuEpi82: Convert packed 16-bit integers in 'a' to packed
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
func MaskCvtsepi16StoreuEpi82(base_addr uintptr, k Mmask32, a M512i)  {
	maskCvtsepi16StoreuEpi82(uintptr(base_addr), uint32(k), [64]byte(a))
}

func maskCvtsepi16StoreuEpi82(base_addr uintptr, k uint32, a [64]byte) 


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
func Cvtusepi16Epi8(a M128i) M128i {
	return M128i(cvtusepi16Epi8([16]byte(a)))
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
func MaskCvtusepi16Epi8(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskCvtusepi16Epi8([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzCvtusepi16Epi8(k Mmask8, a M128i) M128i {
	return M128i(maskzCvtusepi16Epi8(uint8(k), [16]byte(a)))
}

func maskzCvtusepi16Epi8(k uint8, a [16]byte) [16]byte


// Cvtusepi16Epi81: Convert packed unsigned 16-bit integers in 'a' to packed
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
func Cvtusepi16Epi81(a M256i) M128i {
	return M128i(cvtusepi16Epi81([32]byte(a)))
}

func cvtusepi16Epi81(a [32]byte) [16]byte


// MaskCvtusepi16Epi81: Convert packed unsigned 16-bit integers in 'a' to
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
func MaskCvtusepi16Epi81(src M128i, k Mmask16, a M256i) M128i {
	return M128i(maskCvtusepi16Epi81([16]byte(src), uint16(k), [32]byte(a)))
}

func maskCvtusepi16Epi81(src [16]byte, k uint16, a [32]byte) [16]byte


// MaskzCvtusepi16Epi81: Convert packed unsigned 16-bit integers in 'a' to
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
func MaskzCvtusepi16Epi81(k Mmask16, a M256i) M128i {
	return M128i(maskzCvtusepi16Epi81(uint16(k), [32]byte(a)))
}

func maskzCvtusepi16Epi81(k uint16, a [32]byte) [16]byte


// Cvtusepi16Epi82: Convert packed unsigned 16-bit integers in 'a' to packed
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
func Cvtusepi16Epi82(a M512i) M256i {
	return M256i(cvtusepi16Epi82([64]byte(a)))
}

func cvtusepi16Epi82(a [64]byte) [32]byte


// MaskCvtusepi16Epi82: Convert packed unsigned 16-bit integers in 'a' to
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
func MaskCvtusepi16Epi82(src M256i, k Mmask32, a M512i) M256i {
	return M256i(maskCvtusepi16Epi82([32]byte(src), uint32(k), [64]byte(a)))
}

func maskCvtusepi16Epi82(src [32]byte, k uint32, a [64]byte) [32]byte


// MaskzCvtusepi16Epi82: Convert packed unsigned 16-bit integers in 'a' to
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
func MaskzCvtusepi16Epi82(k Mmask32, a M512i) M256i {
	return M256i(maskzCvtusepi16Epi82(uint32(k), [64]byte(a)))
}

func maskzCvtusepi16Epi82(k uint32, a [64]byte) [32]byte


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
func MaskCvtusepi16StoreuEpi8(base_addr uintptr, k Mmask8, a M128i)  {
	maskCvtusepi16StoreuEpi8(uintptr(base_addr), uint8(k), [16]byte(a))
}

func maskCvtusepi16StoreuEpi8(base_addr uintptr, k uint8, a [16]byte) 


// MaskCvtusepi16StoreuEpi81: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
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
func MaskCvtusepi16StoreuEpi81(base_addr uintptr, k Mmask16, a M256i)  {
	maskCvtusepi16StoreuEpi81(uintptr(base_addr), uint16(k), [32]byte(a))
}

func maskCvtusepi16StoreuEpi81(base_addr uintptr, k uint16, a [32]byte) 


// MaskCvtusepi16StoreuEpi82: Convert packed unsigned 16-bit integers in 'a' to
// packed unsigned 8-bit integers with unsigned saturation, and store the
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
func MaskCvtusepi16StoreuEpi82(base_addr uintptr, k Mmask32, a M512i)  {
	maskCvtusepi16StoreuEpi82(uintptr(base_addr), uint32(k), [64]byte(a))
}

func maskCvtusepi16StoreuEpi82(base_addr uintptr, k uint32, a [64]byte) 


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
func DbsadEpu8(a M128i, b M128i, imm8 int) M128i {
	return M128i(dbsadEpu8([16]byte(a), [16]byte(b), imm8))
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
func MaskDbsadEpu8(src M128i, k Mmask8, a M128i, b M128i, imm8 int) M128i {
	return M128i(maskDbsadEpu8([16]byte(src), uint8(k), [16]byte(a), [16]byte(b), imm8))
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
func MaskzDbsadEpu8(k Mmask8, a M128i, b M128i, imm8 int) M128i {
	return M128i(maskzDbsadEpu8(uint8(k), [16]byte(a), [16]byte(b), imm8))
}

func maskzDbsadEpu8(k uint8, a [16]byte, b [16]byte, imm8 int) [16]byte


// DbsadEpu81: Compute the sum of absolute differences (SADs) of quadruplets of
// unsigned 8-bit integers in 'a' compared to those in 'b', and store the
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
func DbsadEpu81(a M256i, b M256i, imm8 int) M256i {
	return M256i(dbsadEpu81([32]byte(a), [32]byte(b), imm8))
}

func dbsadEpu81(a [32]byte, b [32]byte, imm8 int) [32]byte


// MaskDbsadEpu81: Compute the sum of absolute differences (SADs) of
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
func MaskDbsadEpu81(src M256i, k Mmask16, a M256i, b M256i, imm8 int) M256i {
	return M256i(maskDbsadEpu81([32]byte(src), uint16(k), [32]byte(a), [32]byte(b), imm8))
}

func maskDbsadEpu81(src [32]byte, k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte


// MaskzDbsadEpu81: Compute the sum of absolute differences (SADs) of
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
func MaskzDbsadEpu81(k Mmask16, a M256i, b M256i, imm8 int) M256i {
	return M256i(maskzDbsadEpu81(uint16(k), [32]byte(a), [32]byte(b), imm8))
}

func maskzDbsadEpu81(k uint16, a [32]byte, b [32]byte, imm8 int) [32]byte


// DbsadEpu82: Compute the sum of absolute differences (SADs) of quadruplets of
// unsigned 8-bit integers in 'a' compared to those in 'b', and store the
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
func DbsadEpu82(a M512i, b M512i, imm8 int) M512i {
	return M512i(dbsadEpu82([64]byte(a), [64]byte(b), imm8))
}

func dbsadEpu82(a [64]byte, b [64]byte, imm8 int) [64]byte


// MaskDbsadEpu82: Compute the sum of absolute differences (SADs) of
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
func MaskDbsadEpu82(src M512i, k Mmask32, a M512i, b M512i, imm8 int) M512i {
	return M512i(maskDbsadEpu82([64]byte(src), uint32(k), [64]byte(a), [64]byte(b), imm8))
}

func maskDbsadEpu82(src [64]byte, k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte


// MaskzDbsadEpu82: Compute the sum of absolute differences (SADs) of
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
func MaskzDbsadEpu82(k Mmask32, a M512i, b M512i, imm8 int) M512i {
	return M512i(maskzDbsadEpu82(uint32(k), [64]byte(a), [64]byte(b), imm8))
}

func maskzDbsadEpu82(k uint32, a [64]byte, b [64]byte, imm8 int) [64]byte


// Kunpackd: Unpack and interleave 32 bits from masks 'a' and 'b', and store
// the 64-bit result in 'k'. 
//
//		k[31:0] := a[31:0]
//		k[63:32] := b[31:0]
//		k[MAX:64] := 0
//
// Instruction: 'KUNPCKDQ'. Intrinsic: '_mm512_kunpackd'.
// Requires AVX512BW.
func Kunpackd(a Mmask64, b Mmask64) Mmask64 {
	return Mmask64(kunpackd(uint64(a), uint64(b)))
}

func kunpackd(a uint64, b uint64) uint64


// Kunpackw: Unpack and interleave 16 bits from masks 'a' and 'b', and store
// the 32-bit result in 'k'. 
//
//		k[15:0] := a[15:0]
//		k[31:16] := b[15:0]
//		k[MAX:32] := 0
//
// Instruction: 'KUNPCKWD'. Intrinsic: '_mm512_kunpackw'.
// Requires AVX512BW.
func Kunpackw(a Mmask32, b Mmask32) Mmask32 {
	return Mmask32(kunpackw(uint32(a), uint32(b)))
}

func kunpackw(a uint32, b uint32) uint32


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
func MaskLoaduEpi16(src M128i, k Mmask8, mem_addr uintptr) M128i {
	return M128i(maskLoaduEpi16([16]byte(src), uint8(k), uintptr(mem_addr)))
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
func MaskzLoaduEpi16(k Mmask8, mem_addr uintptr) M128i {
	return M128i(maskzLoaduEpi16(uint8(k), uintptr(mem_addr)))
}

func maskzLoaduEpi16(k uint8, mem_addr uintptr) [16]byte


// MaskLoaduEpi161: Load packed 16-bit integers from memory into 'dst' using
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
func MaskLoaduEpi161(src M256i, k Mmask16, mem_addr uintptr) M256i {
	return M256i(maskLoaduEpi161([32]byte(src), uint16(k), uintptr(mem_addr)))
}

func maskLoaduEpi161(src [32]byte, k uint16, mem_addr uintptr) [32]byte


// MaskzLoaduEpi161: Load packed 16-bit integers from memory into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
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
func MaskzLoaduEpi161(k Mmask16, mem_addr uintptr) M256i {
	return M256i(maskzLoaduEpi161(uint16(k), uintptr(mem_addr)))
}

func maskzLoaduEpi161(k uint16, mem_addr uintptr) [32]byte


// MaskLoaduEpi162: Load packed 16-bit integers from memory into 'dst' using
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
func MaskLoaduEpi162(src M512i, k Mmask32, mem_addr uintptr) M512i {
	return M512i(maskLoaduEpi162([64]byte(src), uint32(k), uintptr(mem_addr)))
}

func maskLoaduEpi162(src [64]byte, k uint32, mem_addr uintptr) [64]byte


// MaskzLoaduEpi162: Load packed 16-bit integers from memory into 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
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
func MaskzLoaduEpi162(k Mmask32, mem_addr uintptr) M512i {
	return M512i(maskzLoaduEpi162(uint32(k), uintptr(mem_addr)))
}

func maskzLoaduEpi162(k uint32, mem_addr uintptr) [64]byte


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
func MaskLoaduEpi8(src M128i, k Mmask16, mem_addr uintptr) M128i {
	return M128i(maskLoaduEpi8([16]byte(src), uint16(k), uintptr(mem_addr)))
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
func MaskzLoaduEpi8(k Mmask16, mem_addr uintptr) M128i {
	return M128i(maskzLoaduEpi8(uint16(k), uintptr(mem_addr)))
}

func maskzLoaduEpi8(k uint16, mem_addr uintptr) [16]byte


// MaskLoaduEpi81: Load packed 8-bit integers from memory into 'dst' using
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
func MaskLoaduEpi81(src M256i, k Mmask32, mem_addr uintptr) M256i {
	return M256i(maskLoaduEpi81([32]byte(src), uint32(k), uintptr(mem_addr)))
}

func maskLoaduEpi81(src [32]byte, k uint32, mem_addr uintptr) [32]byte


// MaskzLoaduEpi81: Load packed 8-bit integers from memory into 'dst' using
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
func MaskzLoaduEpi81(k Mmask32, mem_addr uintptr) M256i {
	return M256i(maskzLoaduEpi81(uint32(k), uintptr(mem_addr)))
}

func maskzLoaduEpi81(k uint32, mem_addr uintptr) [32]byte


// MaskLoaduEpi82: Load packed 8-bit integers from memory into 'dst' using
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
func MaskLoaduEpi82(src M512i, k Mmask64, mem_addr uintptr) M512i {
	return M512i(maskLoaduEpi82([64]byte(src), uint64(k), uintptr(mem_addr)))
}

func maskLoaduEpi82(src [64]byte, k uint64, mem_addr uintptr) [64]byte


// MaskzLoaduEpi82: Load packed 8-bit integers from memory into 'dst' using
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
func MaskzLoaduEpi82(k Mmask64, mem_addr uintptr) M512i {
	return M512i(maskzLoaduEpi82(uint64(k), uintptr(mem_addr)))
}

func maskzLoaduEpi82(k uint64, mem_addr uintptr) [64]byte


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
func MaskMaddEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMaddEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMaddEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMaddEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaddEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMaddEpi161: Multiply packed 16-bit integers in 'a' and 'b', producing
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
func MaskMaddEpi161(src M256i, k Mmask8, a M256i, b M256i) M256i {
	return M256i(maskMaddEpi161([32]byte(src), uint8(k), [32]byte(a), [32]byte(b)))
}

func maskMaddEpi161(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte


// MaskzMaddEpi161: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
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
func MaskzMaddEpi161(k Mmask8, a M256i, b M256i) M256i {
	return M256i(maskzMaddEpi161(uint8(k), [32]byte(a), [32]byte(b)))
}

func maskzMaddEpi161(k uint8, a [32]byte, b [32]byte) [32]byte


// MaddEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
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
func MaddEpi16(a M512i, b M512i) M512i {
	return M512i(maddEpi16([64]byte(a), [64]byte(b)))
}

func maddEpi16(a [64]byte, b [64]byte) [64]byte


// MaskMaddEpi162: Multiply packed 16-bit integers in 'a' and 'b', producing
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
func MaskMaddEpi162(src M512i, k Mmask16, a M512i, b M512i) M512i {
	return M512i(maskMaddEpi162([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMaddEpi162(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MaskzMaddEpi162: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers. Horizontally add adjacent pairs of
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
func MaskzMaddEpi162(k Mmask16, a M512i, b M512i) M512i {
	return M512i(maskzMaddEpi162(uint16(k), [64]byte(a), [64]byte(b)))
}

func maskzMaddEpi162(k uint16, a [64]byte, b [64]byte) [64]byte


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
func MaskMaddubsEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMaddubsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMaddubsEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMaddubsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaddubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMaddubsEpi161: Multiply packed unsigned 8-bit integers in 'a' by packed
// signed 8-bit integers in 'b', producing intermediate signed 16-bit integers.
// Horizontally add adjacent pairs of intermediate signed 16-bit integers, and
// pack the saturated results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
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
func MaskMaddubsEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMaddubsEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMaddubsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMaddubsEpi161: Multiply packed unsigned 8-bit integers in 'a' by packed
// signed 8-bit integers in 'b', producing intermediate signed 16-bit integers.
// Horizontally add adjacent pairs of intermediate signed 16-bit integers, and
// pack the saturated results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
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
func MaskzMaddubsEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMaddubsEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMaddubsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaddubsEpi16: Vertically multiply each unsigned 8-bit integer from 'a' with
// the corresponding signed 8-bit integer from 'b', producing intermediate
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
func MaddubsEpi16(a M512i, b M512i) M512i {
	return M512i(maddubsEpi16([64]byte(a), [64]byte(b)))
}

func maddubsEpi16(a [64]byte, b [64]byte) [64]byte


// MaskMaddubsEpi162: Multiply packed unsigned 8-bit integers in 'a' by packed
// signed 8-bit integers in 'b', producing intermediate signed 16-bit integers.
// Horizontally add adjacent pairs of intermediate signed 16-bit integers, and
// pack the saturated results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
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
func MaskMaddubsEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMaddubsEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMaddubsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMaddubsEpi162: Multiply packed unsigned 8-bit integers in 'a' by packed
// signed 8-bit integers in 'b', producing intermediate signed 16-bit integers.
// Horizontally add adjacent pairs of intermediate signed 16-bit integers, and
// pack the saturated results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
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
func MaskzMaddubsEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMaddubsEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMaddubsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMaxEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMaxEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMaxEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMaxEpi161: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskMaxEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMaxEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMaxEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMaxEpi161: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskzMaxEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMaxEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMaxEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMaxEpi162: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskMaxEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMaxEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMaxEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMaxEpi162: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskzMaxEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMaxEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMaxEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// MaxEpi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
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
func MaxEpi16(a M512i, b M512i) M512i {
	return M512i(maxEpi16([64]byte(a), [64]byte(b)))
}

func maxEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskMaxEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzMaxEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzMaxEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskMaxEpi81: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst' using writemask 'k' (elements are copied from 'src'
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
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm256_mask_max_epi8'.
// Requires AVX512BW.
func MaskMaxEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskMaxEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskMaxEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzMaxEpi81: Compare packed 8-bit integers in 'a' and 'b', and store
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
func MaskzMaxEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzMaxEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzMaxEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskMaxEpi82: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst' using writemask 'k' (elements are copied from 'src'
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
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSB'. Intrinsic: '_mm512_mask_max_epi8'.
// Requires AVX512BW.
func MaskMaxEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskMaxEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskMaxEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzMaxEpi82: Compare packed 8-bit integers in 'a' and 'b', and store
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
func MaskzMaxEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzMaxEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzMaxEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// MaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
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
func MaxEpi8(a M512i, b M512i) M512i {
	return M512i(maxEpi8([64]byte(a), [64]byte(b)))
}

func maxEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpu16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMaxEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMaxEpu16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMaxEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMaxEpu161: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
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
func MaskMaxEpu161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMaxEpu161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMaxEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMaxEpu161: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
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
func MaskzMaxEpu161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMaxEpu161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMaxEpu161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMaxEpu162: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
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
func MaskMaxEpu162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMaxEpu162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMaxEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMaxEpu162: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
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
func MaskzMaxEpu162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMaxEpu162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMaxEpu162(k uint32, a [64]byte, b [64]byte) [64]byte


// MaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and store
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
// Instruction: 'VPMAXUW'. Intrinsic: '_mm512_max_epu16'.
// Requires AVX512BW.
func MaxEpu16(a M512i, b M512i) M512i {
	return M512i(maxEpu16([64]byte(a), [64]byte(b)))
}

func maxEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskMaxEpu8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskMaxEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzMaxEpu8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzMaxEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMaxEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskMaxEpu81: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskMaxEpu81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskMaxEpu81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskMaxEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzMaxEpu81: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskzMaxEpu81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzMaxEpu81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzMaxEpu81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskMaxEpu82: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskMaxEpu82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskMaxEpu82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskMaxEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzMaxEpu82: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskzMaxEpu82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzMaxEpu82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzMaxEpu82(k uint64, a [64]byte, b [64]byte) [64]byte


// MaxEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
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
func MaxEpu8(a M512i, b M512i) M512i {
	return M512i(maxEpu8([64]byte(a), [64]byte(b)))
}

func maxEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMinEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMinEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMinEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMinEpi161: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskMinEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMinEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMinEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMinEpi161: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskzMinEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMinEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMinEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMinEpi162: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskMinEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMinEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMinEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMinEpi162: Compare packed 16-bit integers in 'a' and 'b', and store
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
func MaskzMinEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMinEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMinEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// MinEpi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
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
func MinEpi16(a M512i, b M512i) M512i {
	return M512i(minEpi16([64]byte(a), [64]byte(b)))
}

func minEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskMinEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzMinEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzMinEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskMinEpi81: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst' using writemask 'k' (elements are copied from 'src'
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
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm256_mask_min_epi8'.
// Requires AVX512BW.
func MaskMinEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskMinEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskMinEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzMinEpi81: Compare packed 8-bit integers in 'a' and 'b', and store
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
func MaskzMinEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzMinEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzMinEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskMinEpi82: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst' using writemask 'k' (elements are copied from 'src'
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
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSB'. Intrinsic: '_mm512_mask_min_epi8'.
// Requires AVX512BW.
func MaskMinEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskMinEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskMinEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzMinEpi82: Compare packed 8-bit integers in 'a' and 'b', and store
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
func MaskzMinEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzMinEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzMinEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// MinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
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
func MinEpi8(a M512i, b M512i) M512i {
	return M512i(minEpi8([64]byte(a), [64]byte(b)))
}

func minEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpu16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMinEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMinEpu16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMinEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMinEpu161: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
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
func MaskMinEpu161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMinEpu161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMinEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMinEpu161: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
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
func MaskzMinEpu161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMinEpu161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMinEpu161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMinEpu162: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
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
func MaskMinEpu162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMinEpu162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMinEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMinEpu162: Compare packed unsigned 16-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). 
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
func MaskzMinEpu162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMinEpu162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMinEpu162(k uint32, a [64]byte, b [64]byte) [64]byte


// MinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and store
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
// Instruction: 'VPMINUW'. Intrinsic: '_mm512_min_epu16'.
// Requires AVX512BW.
func MinEpu16(a M512i, b M512i) M512i {
	return M512i(minEpu16([64]byte(a), [64]byte(b)))
}

func minEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskMinEpu8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskMinEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzMinEpu8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzMinEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMinEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskMinEpu81: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskMinEpu81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskMinEpu81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskMinEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzMinEpu81: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskzMinEpu81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzMinEpu81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzMinEpu81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskMinEpu82: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskMinEpu82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskMinEpu82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskMinEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzMinEpu82: Compare packed unsigned 8-bit integers in 'a' and 'b', and
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
func MaskzMinEpu82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzMinEpu82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzMinEpu82(k uint64, a [64]byte, b [64]byte) [64]byte


// MinEpu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
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
func MinEpu8(a M512i, b M512i) M512i {
	return M512i(minEpu8([64]byte(a), [64]byte(b)))
}

func minEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskMovEpi16(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskMovEpi16([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzMovEpi16(k Mmask8, a M128i) M128i {
	return M128i(maskzMovEpi16(uint8(k), [16]byte(a)))
}

func maskzMovEpi16(k uint8, a [16]byte) [16]byte


// MaskMovEpi161: Move packed 16-bit integers from 'a' into 'dst' using
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
func MaskMovEpi161(src M256i, k Mmask16, a M256i) M256i {
	return M256i(maskMovEpi161([32]byte(src), uint16(k), [32]byte(a)))
}

func maskMovEpi161(src [32]byte, k uint16, a [32]byte) [32]byte


// MaskzMovEpi161: Move packed 16-bit integers from 'a' into 'dst' using
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
func MaskzMovEpi161(k Mmask16, a M256i) M256i {
	return M256i(maskzMovEpi161(uint16(k), [32]byte(a)))
}

func maskzMovEpi161(k uint16, a [32]byte) [32]byte


// MaskMovEpi162: Move packed 16-bit integers from 'a' into 'dst' using
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
func MaskMovEpi162(src M512i, k Mmask32, a M512i) M512i {
	return M512i(maskMovEpi162([64]byte(src), uint32(k), [64]byte(a)))
}

func maskMovEpi162(src [64]byte, k uint32, a [64]byte) [64]byte


// MaskzMovEpi162: Move packed 16-bit integers from 'a' into 'dst' using
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
func MaskzMovEpi162(k Mmask32, a M512i) M512i {
	return M512i(maskzMovEpi162(uint32(k), [64]byte(a)))
}

func maskzMovEpi162(k uint32, a [64]byte) [64]byte


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
func MaskMovEpi8(src M128i, k Mmask16, a M128i) M128i {
	return M128i(maskMovEpi8([16]byte(src), uint16(k), [16]byte(a)))
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
func MaskzMovEpi8(k Mmask16, a M128i) M128i {
	return M128i(maskzMovEpi8(uint16(k), [16]byte(a)))
}

func maskzMovEpi8(k uint16, a [16]byte) [16]byte


// MaskMovEpi81: Move packed 8-bit integers from 'a' into 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
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
func MaskMovEpi81(src M256i, k Mmask32, a M256i) M256i {
	return M256i(maskMovEpi81([32]byte(src), uint32(k), [32]byte(a)))
}

func maskMovEpi81(src [32]byte, k uint32, a [32]byte) [32]byte


// MaskzMovEpi81: Move packed 8-bit integers from 'a' into 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzMovEpi81(k Mmask32, a M256i) M256i {
	return M256i(maskzMovEpi81(uint32(k), [32]byte(a)))
}

func maskzMovEpi81(k uint32, a [32]byte) [32]byte


// MaskMovEpi82: Move packed 8-bit integers from 'a' into 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
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
func MaskMovEpi82(src M512i, k Mmask64, a M512i) M512i {
	return M512i(maskMovEpi82([64]byte(src), uint64(k), [64]byte(a)))
}

func maskMovEpi82(src [64]byte, k uint64, a [64]byte) [64]byte


// MaskzMovEpi82: Move packed 8-bit integers from 'a' into 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzMovEpi82(k Mmask64, a M512i) M512i {
	return M512i(maskzMovEpi82(uint64(k), [64]byte(a)))
}

func maskzMovEpi82(k uint64, a [64]byte) [64]byte


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
func Movepi16Mask(a M128i) Mmask8 {
	return Mmask8(movepi16Mask([16]byte(a)))
}

func movepi16Mask(a [16]byte) uint8


// Movepi16Mask1: Set each bit of mask register 'k' based on the most
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
func Movepi16Mask1(a M256i) Mmask16 {
	return Mmask16(movepi16Mask1([32]byte(a)))
}

func movepi16Mask1(a [32]byte) uint16


// Movepi16Mask2: Set each bit of mask register 'k' based on the most
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
func Movepi16Mask2(a M512i) Mmask32 {
	return Mmask32(movepi16Mask2([64]byte(a)))
}

func movepi16Mask2(a [64]byte) uint32


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
func Movepi8Mask(a M128i) Mmask16 {
	return Mmask16(movepi8Mask([16]byte(a)))
}

func movepi8Mask(a [16]byte) uint16


// Movepi8Mask1: Set each bit of mask register 'k' based on the most
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
func Movepi8Mask1(a M256i) Mmask32 {
	return Mmask32(movepi8Mask1([32]byte(a)))
}

func movepi8Mask1(a [32]byte) uint32


// Movepi8Mask2: Set each bit of mask register 'k' based on the most
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
func Movepi8Mask2(a M512i) Mmask64 {
	return Mmask64(movepi8Mask2([64]byte(a)))
}

func movepi8Mask2(a [64]byte) uint64


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
func MovmEpi16(k Mmask8) M128i {
	return M128i(movmEpi16(uint8(k)))
}

func movmEpi16(k uint8) [16]byte


// MovmEpi161: Set each packed 16-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
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
func MovmEpi161(k Mmask16) M256i {
	return M256i(movmEpi161(uint16(k)))
}

func movmEpi161(k uint16) [32]byte


// MovmEpi162: Set each packed 16-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
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
func MovmEpi162(k Mmask32) M512i {
	return M512i(movmEpi162(uint32(k)))
}

func movmEpi162(k uint32) [64]byte


// MovmEpi8: Set each packed 8-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
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
func MovmEpi8(k Mmask32) M256i {
	return M256i(movmEpi8(uint32(k)))
}

func movmEpi8(k uint32) [32]byte


// MovmEpi81: Set each packed 8-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
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
func MovmEpi81(k Mmask64) M512i {
	return M512i(movmEpi81(uint64(k)))
}

func movmEpi81(k uint64) [64]byte


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
func MaskMulhiEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMulhiEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMulhiEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMulhiEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMulhiEpi161: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskMulhiEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMulhiEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMulhiEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMulhiEpi161: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskzMulhiEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMulhiEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMulhiEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMulhiEpi162: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskMulhiEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMulhiEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMulhiEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMulhiEpi162: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskzMulhiEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMulhiEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMulhiEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// MulhiEpi16: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the high 16 bits of the intermediate
// integers in 'dst'. 
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
func MulhiEpi16(a M512i, b M512i) M512i {
	return M512i(mulhiEpi16([64]byte(a), [64]byte(b)))
}

func mulhiEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMulhiEpu16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMulhiEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMulhiEpu16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMulhiEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulhiEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMulhiEpu161: Multiply the packed unsigned 16-bit integers in 'a' and
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
func MaskMulhiEpu161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMulhiEpu161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMulhiEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMulhiEpu161: Multiply the packed unsigned 16-bit integers in 'a' and
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
func MaskzMulhiEpu161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMulhiEpu161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMulhiEpu161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMulhiEpu162: Multiply the packed unsigned 16-bit integers in 'a' and
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
func MaskMulhiEpu162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMulhiEpu162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMulhiEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMulhiEpu162: Multiply the packed unsigned 16-bit integers in 'a' and
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
func MaskzMulhiEpu162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMulhiEpu162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMulhiEpu162(k uint32, a [64]byte, b [64]byte) [64]byte


// MulhiEpu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
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
func MulhiEpu16(a M512i, b M512i) M512i {
	return M512i(mulhiEpu16([64]byte(a), [64]byte(b)))
}

func mulhiEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskMulhrsEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMulhrsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMulhrsEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMulhrsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulhrsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMulhrsEpi161: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst' using writemask 'k' (elements are copied from 'src' when the
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
func MaskMulhrsEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMulhrsEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMulhrsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMulhrsEpi161: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzMulhrsEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMulhrsEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMulhrsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMulhrsEpi162: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst' using writemask 'k' (elements are copied from 'src' when the
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
func MaskMulhrsEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMulhrsEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMulhrsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMulhrsEpi162: Multiply packed 16-bit integers in 'a' and 'b', producing
// intermediate signed 32-bit integers. Truncate each intermediate integer to
// the 18 most significant bits, round by adding 1, and store bits [16:1] to
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzMulhrsEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMulhrsEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMulhrsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// MulhrsEpi16: Multiply packed 16-bit integers in 'a' and 'b', producing
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
func MulhrsEpi16(a M512i, b M512i) M512i {
	return M512i(mulhrsEpi16([64]byte(a), [64]byte(b)))
}

func mulhrsEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskMulloEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMulloEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMulloEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMulloEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskMulloEpi161: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskMulloEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskMulloEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskMulloEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzMulloEpi161: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskzMulloEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzMulloEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzMulloEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskMulloEpi162: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskMulloEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskMulloEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskMulloEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzMulloEpi162: Multiply the packed 16-bit integers in 'a' and 'b',
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
func MaskzMulloEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzMulloEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzMulloEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// MulloEpi16: Multiply the packed 16-bit integers in 'a' and 'b', producing
// intermediate 32-bit integers, and store the low 16 bits of the intermediate
// integers in 'dst'. 
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
func MulloEpi16(a M512i, b M512i) M512i {
	return M512i(mulloEpi16([64]byte(a), [64]byte(b)))
}

func mulloEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskPacksEpi16(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskPacksEpi16([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzPacksEpi16(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzPacksEpi16(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzPacksEpi16(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskPacksEpi161: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskPacksEpi161(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskPacksEpi161([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskPacksEpi161(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzPacksEpi161: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskzPacksEpi161(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzPacksEpi161(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzPacksEpi161(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskPacksEpi162: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskPacksEpi162(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskPacksEpi162([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskPacksEpi162(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzPacksEpi162: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskzPacksEpi162(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzPacksEpi162(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzPacksEpi162(k uint64, a [64]byte, b [64]byte) [64]byte


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
func PacksEpi16(a M512i, b M512i) M512i {
	return M512i(packsEpi16([64]byte(a), [64]byte(b)))
}

func packsEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskPacksEpi32(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskPacksEpi32([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzPacksEpi32(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzPacksEpi32(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzPacksEpi32(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskPacksEpi321: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskPacksEpi321(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskPacksEpi321([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskPacksEpi321(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzPacksEpi321: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskzPacksEpi321(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzPacksEpi321(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzPacksEpi321(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskPacksEpi322: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskPacksEpi322(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskPacksEpi322([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskPacksEpi322(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzPacksEpi322: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskzPacksEpi322(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzPacksEpi322(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzPacksEpi322(k uint32, a [64]byte, b [64]byte) [64]byte


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
func PacksEpi32(a M512i, b M512i) M512i {
	return M512i(packsEpi32([64]byte(a), [64]byte(b)))
}

func packsEpi32(a [64]byte, b [64]byte) [64]byte


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
func MaskPackusEpi16(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskPackusEpi16([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzPackusEpi16(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzPackusEpi16(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzPackusEpi16(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskPackusEpi161: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskPackusEpi161(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskPackusEpi161([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskPackusEpi161(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzPackusEpi161: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskzPackusEpi161(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzPackusEpi161(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzPackusEpi161(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskPackusEpi162: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskPackusEpi162(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskPackusEpi162([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskPackusEpi162(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzPackusEpi162: Convert packed 16-bit integers from 'a' and 'b' to packed
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
func MaskzPackusEpi162(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzPackusEpi162(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzPackusEpi162(k uint64, a [64]byte, b [64]byte) [64]byte


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
func PackusEpi16(a M512i, b M512i) M512i {
	return M512i(packusEpi16([64]byte(a), [64]byte(b)))
}

func packusEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskPackusEpi32(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskPackusEpi32([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzPackusEpi32(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzPackusEpi32(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzPackusEpi32(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskPackusEpi321: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskPackusEpi321(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskPackusEpi321([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskPackusEpi321(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzPackusEpi321: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskzPackusEpi321(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzPackusEpi321(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzPackusEpi321(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskPackusEpi322: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskPackusEpi322(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskPackusEpi322([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskPackusEpi322(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzPackusEpi322: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func MaskzPackusEpi322(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzPackusEpi322(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzPackusEpi322(k uint32, a [64]byte, b [64]byte) [64]byte


// PackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
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
func PackusEpi32(a M512i, b M512i) M512i {
	return M512i(packusEpi32([64]byte(a), [64]byte(b)))
}

func packusEpi32(a [64]byte, b [64]byte) [64]byte


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
func MaskPermutex2varEpi16(a M128i, k Mmask8, idx M128i, b M128i) M128i {
	return M128i(maskPermutex2varEpi16([16]byte(a), uint8(k), [16]byte(idx), [16]byte(b)))
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
func Mask2Permutex2varEpi16(a M128i, idx M128i, k Mmask8, b M128i) M128i {
	return M128i(mask2Permutex2varEpi16([16]byte(a), [16]byte(idx), uint8(k), [16]byte(b)))
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
func MaskzPermutex2varEpi16(k Mmask8, a M128i, idx M128i, b M128i) M128i {
	return M128i(maskzPermutex2varEpi16(uint8(k), [16]byte(a), [16]byte(idx), [16]byte(b)))
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
func Permutex2varEpi16(a M128i, idx M128i, b M128i) M128i {
	return M128i(permutex2varEpi16([16]byte(a), [16]byte(idx), [16]byte(b)))
}

func permutex2varEpi16(a [16]byte, idx [16]byte, b [16]byte) [16]byte


// MaskPermutex2varEpi161: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'a' when the
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
func MaskPermutex2varEpi161(a M256i, k Mmask16, idx M256i, b M256i) M256i {
	return M256i(maskPermutex2varEpi161([32]byte(a), uint16(k), [32]byte(idx), [32]byte(b)))
}

func maskPermutex2varEpi161(a [32]byte, k uint16, idx [32]byte, b [32]byte) [32]byte


// Mask2Permutex2varEpi161: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'idx' when the
// corresponding mask bit is not set). 
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
func Mask2Permutex2varEpi161(a M256i, idx M256i, k Mmask16, b M256i) M256i {
	return M256i(mask2Permutex2varEpi161([32]byte(a), [32]byte(idx), uint16(k), [32]byte(b)))
}

func mask2Permutex2varEpi161(a [32]byte, idx [32]byte, k uint16, b [32]byte) [32]byte


// MaskzPermutex2varEpi161: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzPermutex2varEpi161(k Mmask16, a M256i, idx M256i, b M256i) M256i {
	return M256i(maskzPermutex2varEpi161(uint16(k), [32]byte(a), [32]byte(idx), [32]byte(b)))
}

func maskzPermutex2varEpi161(k uint16, a [32]byte, idx [32]byte, b [32]byte) [32]byte


// Permutex2varEpi161: Shuffle 16-bit integers in 'a' and 'b' across lanes
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
func Permutex2varEpi161(a M256i, idx M256i, b M256i) M256i {
	return M256i(permutex2varEpi161([32]byte(a), [32]byte(idx), [32]byte(b)))
}

func permutex2varEpi161(a [32]byte, idx [32]byte, b [32]byte) [32]byte


// MaskPermutex2varEpi162: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'a' when the
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
func MaskPermutex2varEpi162(a M512i, k Mmask32, idx M512i, b M512i) M512i {
	return M512i(maskPermutex2varEpi162([64]byte(a), uint32(k), [64]byte(idx), [64]byte(b)))
}

func maskPermutex2varEpi162(a [64]byte, k uint32, idx [64]byte, b [64]byte) [64]byte


// Mask2Permutex2varEpi162: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'idx' when the
// corresponding mask bit is not set). 
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
func Mask2Permutex2varEpi162(a M512i, idx M512i, k Mmask32, b M512i) M512i {
	return M512i(mask2Permutex2varEpi162([64]byte(a), [64]byte(idx), uint32(k), [64]byte(b)))
}

func mask2Permutex2varEpi162(a [64]byte, idx [64]byte, k uint32, b [64]byte) [64]byte


// MaskzPermutex2varEpi162: Shuffle 16-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzPermutex2varEpi162(k Mmask32, a M512i, idx M512i, b M512i) M512i {
	return M512i(maskzPermutex2varEpi162(uint32(k), [64]byte(a), [64]byte(idx), [64]byte(b)))
}

func maskzPermutex2varEpi162(k uint32, a [64]byte, idx [64]byte, b [64]byte) [64]byte


// Permutex2varEpi162: Shuffle 16-bit integers in 'a' and 'b' across lanes
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
func Permutex2varEpi162(a M512i, idx M512i, b M512i) M512i {
	return M512i(permutex2varEpi162([64]byte(a), [64]byte(idx), [64]byte(b)))
}

func permutex2varEpi162(a [64]byte, idx [64]byte, b [64]byte) [64]byte


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
func MaskPermutexvarEpi16(src M128i, k Mmask8, idx M128i, a M128i) M128i {
	return M128i(maskPermutexvarEpi16([16]byte(src), uint8(k), [16]byte(idx), [16]byte(a)))
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
func MaskzPermutexvarEpi16(k Mmask8, idx M128i, a M128i) M128i {
	return M128i(maskzPermutexvarEpi16(uint8(k), [16]byte(idx), [16]byte(a)))
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
func PermutexvarEpi16(idx M128i, a M128i) M128i {
	return M128i(permutexvarEpi16([16]byte(idx), [16]byte(a)))
}

func permutexvarEpi16(idx [16]byte, a [16]byte) [16]byte


// MaskPermutexvarEpi161: Shuffle 16-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
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
func MaskPermutexvarEpi161(src M256i, k Mmask16, idx M256i, a M256i) M256i {
	return M256i(maskPermutexvarEpi161([32]byte(src), uint16(k), [32]byte(idx), [32]byte(a)))
}

func maskPermutexvarEpi161(src [32]byte, k uint16, idx [32]byte, a [32]byte) [32]byte


// MaskzPermutexvarEpi161: Shuffle 16-bit integers in 'a' across lanes using
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
func MaskzPermutexvarEpi161(k Mmask16, idx M256i, a M256i) M256i {
	return M256i(maskzPermutexvarEpi161(uint16(k), [32]byte(idx), [32]byte(a)))
}

func maskzPermutexvarEpi161(k uint16, idx [32]byte, a [32]byte) [32]byte


// PermutexvarEpi161: Shuffle 16-bit integers in 'a' across lanes using the
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
func PermutexvarEpi161(idx M256i, a M256i) M256i {
	return M256i(permutexvarEpi161([32]byte(idx), [32]byte(a)))
}

func permutexvarEpi161(idx [32]byte, a [32]byte) [32]byte


// MaskPermutexvarEpi162: Shuffle 16-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
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
func MaskPermutexvarEpi162(src M512i, k Mmask32, idx M512i, a M512i) M512i {
	return M512i(maskPermutexvarEpi162([64]byte(src), uint32(k), [64]byte(idx), [64]byte(a)))
}

func maskPermutexvarEpi162(src [64]byte, k uint32, idx [64]byte, a [64]byte) [64]byte


// MaskzPermutexvarEpi162: Shuffle 16-bit integers in 'a' across lanes using
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
func MaskzPermutexvarEpi162(k Mmask32, idx M512i, a M512i) M512i {
	return M512i(maskzPermutexvarEpi162(uint32(k), [64]byte(idx), [64]byte(a)))
}

func maskzPermutexvarEpi162(k uint32, idx [64]byte, a [64]byte) [64]byte


// PermutexvarEpi162: Shuffle 16-bit integers in 'a' across lanes using the
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
func PermutexvarEpi162(idx M512i, a M512i) M512i {
	return M512i(permutexvarEpi162([64]byte(idx), [64]byte(a)))
}

func permutexvarEpi162(idx [64]byte, a [64]byte) [64]byte


// SadEpu8: Compute the absolute differences of packed unsigned 8-bit integers
// in 'a' and 'b', then horizontally sum each consecutive 8 differences to
// produce four unsigned 16-bit integers, and pack these unsigned 16-bit
// integers in the low 16 bits of 64-bit elements in 'dst'. 
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
func SadEpu8(a M512i, b M512i) M512i {
	return M512i(sadEpu8([64]byte(a), [64]byte(b)))
}

func sadEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskSet1Epi16(src M128i, k Mmask8, a int16) M128i {
	return M128i(maskSet1Epi16([16]byte(src), uint8(k), a))
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
func MaskzSet1Epi16(k Mmask8, a int16) M128i {
	return M128i(maskzSet1Epi16(uint8(k), a))
}

func maskzSet1Epi16(k uint8, a int16) [16]byte


// MaskSet1Epi161: Broadcast the low packed 16-bit integer from 'a' to all
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
func MaskSet1Epi161(src M256i, k Mmask16, a int16) M256i {
	return M256i(maskSet1Epi161([32]byte(src), uint16(k), a))
}

func maskSet1Epi161(src [32]byte, k uint16, a int16) [32]byte


// MaskzSet1Epi161: Broadcast 16-bit integer 'a' to all elements of 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzSet1Epi161(k Mmask16, a int16) M256i {
	return M256i(maskzSet1Epi161(uint16(k), a))
}

func maskzSet1Epi161(k uint16, a int16) [32]byte


// MaskSet1Epi162: Broadcast 16-bit integer 'a' to all elements of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskSet1Epi162(src M512i, k Mmask32, a int16) M512i {
	return M512i(maskSet1Epi162([64]byte(src), uint32(k), a))
}

func maskSet1Epi162(src [64]byte, k uint32, a int16) [64]byte


// MaskzSet1Epi162: Broadcast the low packed 16-bit integer from 'a' to all
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
func MaskzSet1Epi162(k Mmask32, a int16) M512i {
	return M512i(maskzSet1Epi162(uint32(k), a))
}

func maskzSet1Epi162(k uint32, a int16) [64]byte


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
func MaskSet1Epi8(src M128i, k Mmask16, a byte) M128i {
	return M128i(maskSet1Epi8([16]byte(src), uint16(k), a))
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
func MaskzSet1Epi8(k Mmask16, a byte) M128i {
	return M128i(maskzSet1Epi8(uint16(k), a))
}

func maskzSet1Epi8(k uint16, a byte) [16]byte


// MaskSet1Epi81: Broadcast 8-bit integer 'a' to all elements of 'dst' using
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
func MaskSet1Epi81(src M256i, k Mmask32, a byte) M256i {
	return M256i(maskSet1Epi81([32]byte(src), uint32(k), a))
}

func maskSet1Epi81(src [32]byte, k uint32, a byte) [32]byte


// MaskzSet1Epi81: Broadcast 8-bit integer 'a' to all elements of 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzSet1Epi81(k Mmask32, a byte) M256i {
	return M256i(maskzSet1Epi81(uint32(k), a))
}

func maskzSet1Epi81(k uint32, a byte) [32]byte


// MaskSet1Epi82: Broadcast 8-bit integer 'a' to all elements of 'dst' using
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
func MaskSet1Epi82(src M512i, k Mmask64, a byte) M512i {
	return M512i(maskSet1Epi82([64]byte(src), uint64(k), a))
}

func maskSet1Epi82(src [64]byte, k uint64, a byte) [64]byte


// MaskzSet1Epi82: Broadcast 8-bit integer 'a' to all elements of 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzSet1Epi82(k Mmask64, a byte) M512i {
	return M512i(maskzSet1Epi82(uint64(k), a))
}

func maskzSet1Epi82(k uint64, a byte) [64]byte


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
func MaskShuffleEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskShuffleEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzShuffleEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzShuffleEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzShuffleEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskShuffleEpi81: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
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
func MaskShuffleEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskShuffleEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskShuffleEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzShuffleEpi81: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
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
func MaskzShuffleEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzShuffleEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzShuffleEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskShuffleEpi82: Shuffle 8-bit integers in 'a' within 128-bit lanes using
// the control in the corresponding 8-bit element of 'b', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'src' when the
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
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFB'. Intrinsic: '_mm512_mask_shuffle_epi8'.
// Requires AVX512BW.
func MaskShuffleEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskShuffleEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskShuffleEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzShuffleEpi82: Shuffle packed 8-bit integers in 'a' according to shuffle
// control mask in the corresponding 8-bit element of 'b', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
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
func MaskzShuffleEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzShuffleEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzShuffleEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// ShuffleEpi8: Shuffle packed 8-bit integers in 'a' according to shuffle
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
func ShuffleEpi8(a M512i, b M512i) M512i {
	return M512i(shuffleEpi8([64]byte(a), [64]byte(b)))
}

func shuffleEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskShufflehiEpi16(src M128i, k Mmask8, a M128i, imm8 int) M128i {
	return M128i(maskShufflehiEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
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
func MaskzShufflehiEpi16(k Mmask8, a M128i, imm8 int) M128i {
	return M128i(maskzShufflehiEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzShufflehiEpi16(k uint8, a [16]byte, imm8 int) [16]byte


// MaskShufflehiEpi161: Shuffle 16-bit integers in the high 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the high 64
// bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskShufflehiEpi161(src M256i, k Mmask16, a M256i, imm8 int) M256i {
	return M256i(maskShufflehiEpi161([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func maskShufflehiEpi161(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte


// MaskzShufflehiEpi161: Shuffle 16-bit integers in the high 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the high 64
// bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzShufflehiEpi161(k Mmask16, a M256i, imm8 int) M256i {
	return M256i(maskzShufflehiEpi161(uint16(k), [32]byte(a), imm8))
}

func maskzShufflehiEpi161(k uint16, a [32]byte, imm8 int) [32]byte


// MaskShufflehiEpi162: Shuffle 16-bit integers in the high 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the high 64
// bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskShufflehiEpi162(src M512i, k Mmask32, a M512i, imm8 int) M512i {
	return M512i(maskShufflehiEpi162([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func maskShufflehiEpi162(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte


// MaskzShufflehiEpi162: Shuffle 16-bit integers in the high 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the high 64
// bits of 128-bit lanes of 'dst', with the low 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzShufflehiEpi162(k Mmask32, a M512i, imm8 int) M512i {
	return M512i(maskzShufflehiEpi162(uint32(k), [64]byte(a), imm8))
}

func maskzShufflehiEpi162(k uint32, a [64]byte, imm8 int) [64]byte


// ShufflehiEpi16: Shuffle 16-bit integers in the high 64 bits of 128-bit lanes
// of 'a' using the control in 'imm8'. Store the results in the high 64 bits of
// 128-bit lanes of 'dst', with the low 64 bits of 128-bit lanes being copied
// from from 'a' to 'dst'. 
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
func ShufflehiEpi16(a M512i, imm8 int) M512i {
	return M512i(shufflehiEpi16([64]byte(a), imm8))
}

func shufflehiEpi16(a [64]byte, imm8 int) [64]byte


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
func MaskShuffleloEpi16(src M128i, k Mmask8, a M128i, imm8 int) M128i {
	return M128i(maskShuffleloEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
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
func MaskzShuffleloEpi16(k Mmask8, a M128i, imm8 int) M128i {
	return M128i(maskzShuffleloEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzShuffleloEpi16(k uint8, a [16]byte, imm8 int) [16]byte


// MaskShuffleloEpi161: Shuffle 16-bit integers in the low 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the low 64
// bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskShuffleloEpi161(src M256i, k Mmask16, a M256i, imm8 int) M256i {
	return M256i(maskShuffleloEpi161([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func maskShuffleloEpi161(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte


// MaskzShuffleloEpi161: Shuffle 16-bit integers in the low 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the low 64
// bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzShuffleloEpi161(k Mmask16, a M256i, imm8 int) M256i {
	return M256i(maskzShuffleloEpi161(uint16(k), [32]byte(a), imm8))
}

func maskzShuffleloEpi161(k uint16, a [32]byte, imm8 int) [32]byte


// MaskShuffleloEpi162: Shuffle 16-bit integers in the low 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the low 64
// bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func MaskShuffleloEpi162(src M512i, k Mmask32, a M512i, imm8 int) M512i {
	return M512i(maskShuffleloEpi162([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func maskShuffleloEpi162(src [64]byte, k uint32, a [64]byte, imm8 int) [64]byte


// MaskzShuffleloEpi162: Shuffle 16-bit integers in the low 64 bits of 128-bit
// lanes of 'a' using the control in 'imm8'. Store the results in the low 64
// bits of 128-bit lanes of 'dst', with the high 64 bits of 128-bit lanes being
// copied from from 'a' to 'dst', using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func MaskzShuffleloEpi162(k Mmask32, a M512i, imm8 int) M512i {
	return M512i(maskzShuffleloEpi162(uint32(k), [64]byte(a), imm8))
}

func maskzShuffleloEpi162(k uint32, a [64]byte, imm8 int) [64]byte


// ShuffleloEpi16: Shuffle 16-bit integers in the low 64 bits of 128-bit lanes
// of 'a' using the control in 'imm8'. Store the results in the low 64 bits of
// 128-bit lanes of 'dst', with the high 64 bits of 128-bit lanes being copied
// from from 'a' to 'dst'. 
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
func ShuffleloEpi16(a M512i, imm8 int) M512i {
	return M512i(shuffleloEpi16([64]byte(a), imm8))
}

func shuffleloEpi16(a [64]byte, imm8 int) [64]byte


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
func MaskSllEpi16(src M128i, k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskSllEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
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
func MaskzSllEpi16(k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskzSllEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSllEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


// MaskSllEpi161: Shift packed 16-bit integers in 'a' left by 'count' while
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
func MaskSllEpi161(src M256i, k Mmask16, a M256i, count M128i) M256i {
	return M256i(maskSllEpi161([32]byte(src), uint16(k), [32]byte(a), [16]byte(count)))
}

func maskSllEpi161(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte


// MaskzSllEpi161: Shift packed 16-bit integers in 'a' left by 'count' while
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
func MaskzSllEpi161(k Mmask16, a M256i, count M128i) M256i {
	return M256i(maskzSllEpi161(uint16(k), [32]byte(a), [16]byte(count)))
}

func maskzSllEpi161(k uint16, a [32]byte, count [16]byte) [32]byte


// MaskSllEpi162: Shift packed 16-bit integers in 'a' left by 'count' while
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
func MaskSllEpi162(src M512i, k Mmask32, a M512i, count M128i) M512i {
	return M512i(maskSllEpi162([64]byte(src), uint32(k), [64]byte(a), [16]byte(count)))
}

func maskSllEpi162(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte


// MaskzSllEpi162: Shift packed 16-bit integers in 'a' left by 'count' while
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
func MaskzSllEpi162(k Mmask32, a M512i, count M128i) M512i {
	return M512i(maskzSllEpi162(uint32(k), [64]byte(a), [16]byte(count)))
}

func maskzSllEpi162(k uint32, a [64]byte, count [16]byte) [64]byte


// SllEpi16: Shift packed 16-bit integers in 'a' left by 'count' while shifting
// in zeros, and store the results in 'dst'. 
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
func SllEpi16(a M512i, count M128i) M512i {
	return M512i(sllEpi16([64]byte(a), [16]byte(count)))
}

func sllEpi16(a [64]byte, count [16]byte) [64]byte


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
func MaskSlliEpi16(src M128i, k Mmask8, a M128i, imm8 uint32) M128i {
	return M128i(maskSlliEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
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
func MaskzSlliEpi16(k Mmask8, a M128i, imm8 uint32) M128i {
	return M128i(maskzSlliEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzSlliEpi16(k uint8, a [16]byte, imm8 uint32) [16]byte


// MaskSlliEpi161: Shift packed 16-bit integers in 'a' left by 'imm8' while
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
func MaskSlliEpi161(src M256i, k Mmask16, a M256i, imm8 uint32) M256i {
	return M256i(maskSlliEpi161([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func maskSlliEpi161(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte


// MaskzSlliEpi161: Shift packed 16-bit integers in 'a' left by 'imm8' while
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
func MaskzSlliEpi161(k Mmask16, a M256i, imm8 uint32) M256i {
	return M256i(maskzSlliEpi161(uint16(k), [32]byte(a), imm8))
}

func maskzSlliEpi161(k uint16, a [32]byte, imm8 uint32) [32]byte


// MaskSlliEpi162: Shift packed 16-bit integers in 'a' left by 'imm8' while
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
func MaskSlliEpi162(src M512i, k Mmask32, a M512i, imm8 uint32) M512i {
	return M512i(maskSlliEpi162([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func maskSlliEpi162(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte


// MaskzSlliEpi162: Shift packed 16-bit integers in 'a' left by 'imm8' while
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
func MaskzSlliEpi162(k Mmask32, a M512i, imm8 uint32) M512i {
	return M512i(maskzSlliEpi162(uint32(k), [64]byte(a), imm8))
}

func maskzSlliEpi162(k uint32, a [64]byte, imm8 uint32) [64]byte


// SlliEpi16: Shift packed 16-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
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
func SlliEpi16(a M512i, imm8 uint32) M512i {
	return M512i(slliEpi16([64]byte(a), imm8))
}

func slliEpi16(a [64]byte, imm8 uint32) [64]byte


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
func MaskSllvEpi16(src M128i, k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskSllvEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
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
func MaskzSllvEpi16(k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskzSllvEpi16(uint8(k), [16]byte(a), [16]byte(count)))
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
func SllvEpi16(a M128i, count M128i) M128i {
	return M128i(sllvEpi16([16]byte(a), [16]byte(count)))
}

func sllvEpi16(a [16]byte, count [16]byte) [16]byte


// MaskSllvEpi161: Shift packed 16-bit integers in 'a' left by the amount
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
func MaskSllvEpi161(src M256i, k Mmask16, a M256i, count M256i) M256i {
	return M256i(maskSllvEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(count)))
}

func maskSllvEpi161(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte


// MaskzSllvEpi161: Shift packed 16-bit integers in 'a' left by the amount
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
func MaskzSllvEpi161(k Mmask16, a M256i, count M256i) M256i {
	return M256i(maskzSllvEpi161(uint16(k), [32]byte(a), [32]byte(count)))
}

func maskzSllvEpi161(k uint16, a [32]byte, count [32]byte) [32]byte


// SllvEpi161: Shift packed 16-bit integers in 'a' left by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
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
func SllvEpi161(a M256i, count M256i) M256i {
	return M256i(sllvEpi161([32]byte(a), [32]byte(count)))
}

func sllvEpi161(a [32]byte, count [32]byte) [32]byte


// MaskSllvEpi162: Shift packed 16-bit integers in 'a' left by the amount
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
func MaskSllvEpi162(src M512i, k Mmask32, a M512i, count M512i) M512i {
	return M512i(maskSllvEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(count)))
}

func maskSllvEpi162(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte


// MaskzSllvEpi162: Shift packed 16-bit integers in 'a' left by the amount
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
func MaskzSllvEpi162(k Mmask32, a M512i, count M512i) M512i {
	return M512i(maskzSllvEpi162(uint32(k), [64]byte(a), [64]byte(count)))
}

func maskzSllvEpi162(k uint32, a [64]byte, count [64]byte) [64]byte


// SllvEpi162: Shift packed 16-bit integers in 'a' left by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
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
func SllvEpi162(a M512i, count M512i) M512i {
	return M512i(sllvEpi162([64]byte(a), [64]byte(count)))
}

func sllvEpi162(a [64]byte, count [64]byte) [64]byte


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
func MaskSraEpi16(src M128i, k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskSraEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
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
func MaskzSraEpi16(k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskzSraEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSraEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


// MaskSraEpi161: Shift packed 16-bit integers in 'a' right by 'count' while
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
func MaskSraEpi161(src M256i, k Mmask16, a M256i, count M128i) M256i {
	return M256i(maskSraEpi161([32]byte(src), uint16(k), [32]byte(a), [16]byte(count)))
}

func maskSraEpi161(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte


// MaskzSraEpi161: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzSraEpi161(k Mmask16, a M256i, count M128i) M256i {
	return M256i(maskzSraEpi161(uint16(k), [32]byte(a), [16]byte(count)))
}

func maskzSraEpi161(k uint16, a [32]byte, count [16]byte) [32]byte


// MaskSraEpi162: Shift packed 16-bit integers in 'a' right by 'count' while
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
func MaskSraEpi162(src M512i, k Mmask32, a M512i, count M128i) M512i {
	return M512i(maskSraEpi162([64]byte(src), uint32(k), [64]byte(a), [16]byte(count)))
}

func maskSraEpi162(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte


// MaskzSraEpi162: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in sign bits, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzSraEpi162(k Mmask32, a M512i, count M128i) M512i {
	return M512i(maskzSraEpi162(uint32(k), [64]byte(a), [16]byte(count)))
}

func maskzSraEpi162(k uint32, a [64]byte, count [16]byte) [64]byte


// SraEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
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
func SraEpi16(a M512i, count M128i) M512i {
	return M512i(sraEpi16([64]byte(a), [16]byte(count)))
}

func sraEpi16(a [64]byte, count [16]byte) [64]byte


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
func MaskSraiEpi16(src M128i, k Mmask8, a M128i, imm8 uint32) M128i {
	return M128i(maskSraiEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
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
func MaskzSraiEpi16(k Mmask8, a M128i, imm8 uint32) M128i {
	return M128i(maskzSraiEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzSraiEpi16(k uint8, a [16]byte, imm8 uint32) [16]byte


// MaskSraiEpi161: Shift packed 16-bit integers in 'a' right by 'imm8' while
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
func MaskSraiEpi161(src M256i, k Mmask16, a M256i, imm8 uint32) M256i {
	return M256i(maskSraiEpi161([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func maskSraiEpi161(src [32]byte, k uint16, a [32]byte, imm8 uint32) [32]byte


// MaskzSraiEpi161: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzSraiEpi161(k Mmask16, a M256i, imm8 uint32) M256i {
	return M256i(maskzSraiEpi161(uint16(k), [32]byte(a), imm8))
}

func maskzSraiEpi161(k uint16, a [32]byte, imm8 uint32) [32]byte


// MaskSraiEpi162: Shift packed 16-bit integers in 'a' right by 'imm8' while
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
func MaskSraiEpi162(src M512i, k Mmask32, a M512i, imm8 uint32) M512i {
	return M512i(maskSraiEpi162([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func maskSraiEpi162(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte


// MaskzSraiEpi162: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzSraiEpi162(k Mmask32, a M512i, imm8 uint32) M512i {
	return M512i(maskzSraiEpi162(uint32(k), [64]byte(a), imm8))
}

func maskzSraiEpi162(k uint32, a [64]byte, imm8 uint32) [64]byte


// SraiEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
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
func SraiEpi16(a M512i, imm8 uint32) M512i {
	return M512i(sraiEpi16([64]byte(a), imm8))
}

func sraiEpi16(a [64]byte, imm8 uint32) [64]byte


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
func MaskSravEpi16(src M128i, k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskSravEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
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
func MaskzSravEpi16(k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskzSravEpi16(uint8(k), [16]byte(a), [16]byte(count)))
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
func SravEpi16(a M128i, count M128i) M128i {
	return M128i(sravEpi16([16]byte(a), [16]byte(count)))
}

func sravEpi16(a [16]byte, count [16]byte) [16]byte


// MaskSravEpi161: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskSravEpi161(src M256i, k Mmask16, a M256i, count M256i) M256i {
	return M256i(maskSravEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(count)))
}

func maskSravEpi161(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte


// MaskzSravEpi161: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskzSravEpi161(k Mmask16, a M256i, count M256i) M256i {
	return M256i(maskzSravEpi161(uint16(k), [32]byte(a), [32]byte(count)))
}

func maskzSravEpi161(k uint16, a [32]byte, count [32]byte) [32]byte


// SravEpi161: Shift packed 16-bit integers in 'a' right by the amount
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
func SravEpi161(a M256i, count M256i) M256i {
	return M256i(sravEpi161([32]byte(a), [32]byte(count)))
}

func sravEpi161(a [32]byte, count [32]byte) [32]byte


// MaskSravEpi162: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskSravEpi162(src M512i, k Mmask32, a M512i, count M512i) M512i {
	return M512i(maskSravEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(count)))
}

func maskSravEpi162(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte


// MaskzSravEpi162: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskzSravEpi162(k Mmask32, a M512i, count M512i) M512i {
	return M512i(maskzSravEpi162(uint32(k), [64]byte(a), [64]byte(count)))
}

func maskzSravEpi162(k uint32, a [64]byte, count [64]byte) [64]byte


// SravEpi162: Shift packed 16-bit integers in 'a' right by the amount
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
func SravEpi162(a M512i, count M512i) M512i {
	return M512i(sravEpi162([64]byte(a), [64]byte(count)))
}

func sravEpi162(a [64]byte, count [64]byte) [64]byte


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
func MaskSrlEpi16(src M128i, k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskSrlEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
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
func MaskzSrlEpi16(k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskzSrlEpi16(uint8(k), [16]byte(a), [16]byte(count)))
}

func maskzSrlEpi16(k uint8, a [16]byte, count [16]byte) [16]byte


// MaskSrlEpi161: Shift packed 16-bit integers in 'a' right by 'count' while
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
func MaskSrlEpi161(src M256i, k Mmask16, a M256i, count M128i) M256i {
	return M256i(maskSrlEpi161([32]byte(src), uint16(k), [32]byte(a), [16]byte(count)))
}

func maskSrlEpi161(src [32]byte, k uint16, a [32]byte, count [16]byte) [32]byte


// MaskzSrlEpi161: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
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
func MaskzSrlEpi161(k Mmask16, a M256i, count M128i) M256i {
	return M256i(maskzSrlEpi161(uint16(k), [32]byte(a), [16]byte(count)))
}

func maskzSrlEpi161(k uint16, a [32]byte, count [16]byte) [32]byte


// MaskSrlEpi162: Shift packed 16-bit integers in 'a' right by 'count' while
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
func MaskSrlEpi162(src M512i, k Mmask32, a M512i, count M128i) M512i {
	return M512i(maskSrlEpi162([64]byte(src), uint32(k), [64]byte(a), [16]byte(count)))
}

func maskSrlEpi162(src [64]byte, k uint32, a [64]byte, count [16]byte) [64]byte


// MaskzSrlEpi162: Shift packed 16-bit integers in 'a' right by 'count' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
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
func MaskzSrlEpi162(k Mmask32, a M512i, count M128i) M512i {
	return M512i(maskzSrlEpi162(uint32(k), [64]byte(a), [16]byte(count)))
}

func maskzSrlEpi162(k uint32, a [64]byte, count [16]byte) [64]byte


// SrlEpi16: Shift packed 16-bit integers in 'a' right by 'count' while
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
func SrlEpi16(a M512i, count M128i) M512i {
	return M512i(srlEpi16([64]byte(a), [16]byte(count)))
}

func srlEpi16(a [64]byte, count [16]byte) [64]byte


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
func MaskSrliEpi16(src M128i, k Mmask8, a M128i, imm8 int) M128i {
	return M128i(maskSrliEpi16([16]byte(src), uint8(k), [16]byte(a), imm8))
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
func MaskzSrliEpi16(k Mmask8, a M128i, imm8 int) M128i {
	return M128i(maskzSrliEpi16(uint8(k), [16]byte(a), imm8))
}

func maskzSrliEpi16(k uint8, a [16]byte, imm8 int) [16]byte


// MaskSrliEpi161: Shift packed 16-bit integers in 'a' right by 'imm8' while
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
func MaskSrliEpi161(src M256i, k Mmask16, a M256i, imm8 int) M256i {
	return M256i(maskSrliEpi161([32]byte(src), uint16(k), [32]byte(a), imm8))
}

func maskSrliEpi161(src [32]byte, k uint16, a [32]byte, imm8 int) [32]byte


// MaskzSrliEpi161: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
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
func MaskzSrliEpi161(k Mmask16, a M256i, imm8 int) M256i {
	return M256i(maskzSrliEpi161(uint16(k), [32]byte(a), imm8))
}

func maskzSrliEpi161(k uint16, a [32]byte, imm8 int) [32]byte


// MaskSrliEpi162: Shift packed 16-bit integers in 'a' right by 'imm8' while
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
func MaskSrliEpi162(src M512i, k Mmask32, a M512i, imm8 uint32) M512i {
	return M512i(maskSrliEpi162([64]byte(src), uint32(k), [64]byte(a), imm8))
}

func maskSrliEpi162(src [64]byte, k uint32, a [64]byte, imm8 uint32) [64]byte


// MaskzSrliEpi162: Shift packed 16-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using zeromask 'k'
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
func MaskzSrliEpi162(k Mmask32, a M512i, imm8 int) M512i {
	return M512i(maskzSrliEpi162(uint32(k), [64]byte(a), imm8))
}

func maskzSrliEpi162(k uint32, a [64]byte, imm8 int) [64]byte


// SrliEpi16: Shift packed 16-bit integers in 'a' right by 'imm8' while
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
func SrliEpi16(a M512i, imm8 uint32) M512i {
	return M512i(srliEpi16([64]byte(a), imm8))
}

func srliEpi16(a [64]byte, imm8 uint32) [64]byte


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
func MaskSrlvEpi16(src M128i, k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskSrlvEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(count)))
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
func MaskzSrlvEpi16(k Mmask8, a M128i, count M128i) M128i {
	return M128i(maskzSrlvEpi16(uint8(k), [16]byte(a), [16]byte(count)))
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
func SrlvEpi16(a M128i, count M128i) M128i {
	return M128i(srlvEpi16([16]byte(a), [16]byte(count)))
}

func srlvEpi16(a [16]byte, count [16]byte) [16]byte


// MaskSrlvEpi161: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskSrlvEpi161(src M256i, k Mmask16, a M256i, count M256i) M256i {
	return M256i(maskSrlvEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(count)))
}

func maskSrlvEpi161(src [32]byte, k uint16, a [32]byte, count [32]byte) [32]byte


// MaskzSrlvEpi161: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskzSrlvEpi161(k Mmask16, a M256i, count M256i) M256i {
	return M256i(maskzSrlvEpi161(uint16(k), [32]byte(a), [32]byte(count)))
}

func maskzSrlvEpi161(k uint16, a [32]byte, count [32]byte) [32]byte


// SrlvEpi161: Shift packed 16-bit integers in 'a' right by the amount
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
func SrlvEpi161(a M256i, count M256i) M256i {
	return M256i(srlvEpi161([32]byte(a), [32]byte(count)))
}

func srlvEpi161(a [32]byte, count [32]byte) [32]byte


// MaskSrlvEpi162: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskSrlvEpi162(src M512i, k Mmask32, a M512i, count M512i) M512i {
	return M512i(maskSrlvEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(count)))
}

func maskSrlvEpi162(src [64]byte, k uint32, a [64]byte, count [64]byte) [64]byte


// MaskzSrlvEpi162: Shift packed 16-bit integers in 'a' right by the amount
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
func MaskzSrlvEpi162(k Mmask32, a M512i, count M512i) M512i {
	return M512i(maskzSrlvEpi162(uint32(k), [64]byte(a), [64]byte(count)))
}

func maskzSrlvEpi162(k uint32, a [64]byte, count [64]byte) [64]byte


// SrlvEpi162: Shift packed 16-bit integers in 'a' right by the amount
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
func SrlvEpi162(a M512i, count M512i) M512i {
	return M512i(srlvEpi162([64]byte(a), [64]byte(count)))
}

func srlvEpi162(a [64]byte, count [64]byte) [64]byte


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
func MaskStoreuEpi16(mem_addr uintptr, k Mmask8, a M128i)  {
	maskStoreuEpi16(uintptr(mem_addr), uint8(k), [16]byte(a))
}

func maskStoreuEpi16(mem_addr uintptr, k uint8, a [16]byte) 


// MaskStoreuEpi161: Store packed 16-bit integers from 'a' into memory using
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
func MaskStoreuEpi161(mem_addr uintptr, k Mmask16, a M256i)  {
	maskStoreuEpi161(uintptr(mem_addr), uint16(k), [32]byte(a))
}

func maskStoreuEpi161(mem_addr uintptr, k uint16, a [32]byte) 


// MaskStoreuEpi162: Store packed 16-bit integers from 'a' into memory using
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
func MaskStoreuEpi162(mem_addr uintptr, k Mmask32, a M512i)  {
	maskStoreuEpi162(uintptr(mem_addr), uint32(k), [64]byte(a))
}

func maskStoreuEpi162(mem_addr uintptr, k uint32, a [64]byte) 


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
func MaskStoreuEpi8(mem_addr uintptr, k Mmask16, a M128i)  {
	maskStoreuEpi8(uintptr(mem_addr), uint16(k), [16]byte(a))
}

func maskStoreuEpi8(mem_addr uintptr, k uint16, a [16]byte) 


// MaskStoreuEpi81: Store packed 8-bit integers from 'a' into memory using
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
func MaskStoreuEpi81(mem_addr uintptr, k Mmask32, a M256i)  {
	maskStoreuEpi81(uintptr(mem_addr), uint32(k), [32]byte(a))
}

func maskStoreuEpi81(mem_addr uintptr, k uint32, a [32]byte) 


// MaskStoreuEpi82: Store packed 8-bit integers from 'a' into memory using
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
func MaskStoreuEpi82(mem_addr uintptr, k Mmask64, a M512i)  {
	maskStoreuEpi82(uintptr(mem_addr), uint64(k), [64]byte(a))
}

func maskStoreuEpi82(mem_addr uintptr, k uint64, a [64]byte) 


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
func MaskSubEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskSubEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzSubEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzSubEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzSubEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskSubEpi161: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func MaskSubEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskSubEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskSubEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzSubEpi161: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func MaskzSubEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzSubEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzSubEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskSubEpi162: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func MaskSubEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskSubEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskSubEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzSubEpi162: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func MaskzSubEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzSubEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzSubEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// SubEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := j*16
//			dst[i+15:i] := a[i+15:i] - b[i+15:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBW'. Intrinsic: '_mm512_sub_epi16'.
// Requires AVX512BW.
func SubEpi16(a M512i, b M512i) M512i {
	return M512i(subEpi16([64]byte(a), [64]byte(b)))
}

func subEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskSubEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskSubEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzSubEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzSubEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzSubEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskSubEpi81: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskSubEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskSubEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskSubEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzSubEpi81: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskzSubEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzSubEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzSubEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskSubEpi82: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskSubEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskSubEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskSubEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzSubEpi82: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskzSubEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzSubEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzSubEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// SubEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit integers in
// 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := a[i+7:i] - b[i+7:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBB'. Intrinsic: '_mm512_sub_epi8'.
// Requires AVX512BW.
func SubEpi8(a M512i, b M512i) M512i {
	return M512i(subEpi8([64]byte(a), [64]byte(b)))
}

func subEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskSubsEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzSubsEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzSubsEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskSubsEpi161: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func MaskSubsEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskSubsEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskSubsEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzSubsEpi161: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzSubsEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzSubsEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzSubsEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskSubsEpi162: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func MaskSubsEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskSubsEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskSubsEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzSubsEpi162: Subtract packed 16-bit integers in 'b' from packed 16-bit
// integers in 'a' using saturation, and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzSubsEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzSubsEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzSubsEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// SubsEpi16: Subtract packed 16-bit integers in 'b' from packed 16-bit
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
func SubsEpi16(a M512i, b M512i) M512i {
	return M512i(subsEpi16([64]byte(a), [64]byte(b)))
}

func subsEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskSubsEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzSubsEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzSubsEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskSubsEpi81: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskSubsEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskSubsEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskSubsEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzSubsEpi81: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskzSubsEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzSubsEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzSubsEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskSubsEpi82: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskSubsEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskSubsEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskSubsEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzSubsEpi82: Subtract packed 8-bit integers in 'b' from packed 8-bit
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
func MaskzSubsEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzSubsEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzSubsEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// SubsEpi8: Subtract packed 8-bit integers in 'b' from packed 8-bit integers
// in 'a' using saturation, and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			dst[i+7:i] := Saturate_To_Int8(a[i+7:i] - b[i+7:i])	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSB'. Intrinsic: '_mm512_subs_epi8'.
// Requires AVX512BW.
func SubsEpi8(a M512i, b M512i) M512i {
	return M512i(subsEpi8([64]byte(a), [64]byte(b)))
}

func subsEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpu16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskSubsEpu16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzSubsEpu16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzSubsEpu16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpu16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskSubsEpu161: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
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
func MaskSubsEpu161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskSubsEpu161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskSubsEpu161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzSubsEpu161: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzSubsEpu161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzSubsEpu161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzSubsEpu161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskSubsEpu162: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
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
func MaskSubsEpu162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskSubsEpu162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskSubsEpu162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzSubsEpu162: Subtract packed unsigned 16-bit integers in 'b' from packed
// unsigned 16-bit integers in 'a' using saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzSubsEpu162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzSubsEpu162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzSubsEpu162(k uint32, a [64]byte, b [64]byte) [64]byte


// SubsEpu16: Subtract packed unsigned 16-bit integers in 'b' from packed
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
func SubsEpu16(a M512i, b M512i) M512i {
	return M512i(subsEpu16([64]byte(a), [64]byte(b)))
}

func subsEpu16(a [64]byte, b [64]byte) [64]byte


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
func MaskSubsEpu8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskSubsEpu8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzSubsEpu8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzSubsEpu8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzSubsEpu8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskSubsEpu81: Subtract packed unsigned 8-bit integers in 'b' from packed
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
func MaskSubsEpu81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskSubsEpu81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskSubsEpu81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzSubsEpu81: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzSubsEpu81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzSubsEpu81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzSubsEpu81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskSubsEpu82: Subtract packed unsigned 8-bit integers in 'b' from packed
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
func MaskSubsEpu82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskSubsEpu82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskSubsEpu82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzSubsEpu82: Subtract packed unsigned 8-bit integers in 'b' from packed
// unsigned 8-bit integers in 'a' using saturation, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzSubsEpu82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzSubsEpu82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzSubsEpu82(k uint64, a [64]byte, b [64]byte) [64]byte


// SubsEpu8: Subtract packed unsigned 8-bit integers in 'b' from packed
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
func SubsEpu8(a M512i, b M512i) M512i {
	return M512i(subsEpu8([64]byte(a), [64]byte(b)))
}

func subsEpu8(a [64]byte, b [64]byte) [64]byte


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
func MaskTestEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskTestEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
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
func TestEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(testEpi16Mask([16]byte(a), [16]byte(b)))
}

func testEpi16Mask(a [16]byte, b [16]byte) uint8


// MaskTestEpi16Mask1: Compute the bitwise AND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
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
func MaskTestEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskTestEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskTestEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// TestEpi16Mask1: Compute the bitwise AND of packed 16-bit integers in 'a' and
// 'b', producing intermediate 16-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm256_test_epi16_mask'.
// Requires AVX512BW.
func TestEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(testEpi16Mask1([32]byte(a), [32]byte(b)))
}

func testEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskTestEpi16Mask2: Compute the bitwise AND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
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
func MaskTestEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskTestEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskTestEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


// TestEpi16Mask2: Compute the bitwise AND of packed 16-bit integers in 'a' and
// 'b', producing intermediate 16-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTMW'. Intrinsic: '_mm512_test_epi16_mask'.
// Requires AVX512BW.
func TestEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(testEpi16Mask2([64]byte(a), [64]byte(b)))
}

func testEpi16Mask2(a [64]byte, b [64]byte) uint32


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
func MaskTestEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskTestEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
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
func TestEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(testEpi8Mask([16]byte(a), [16]byte(b)))
}

func testEpi8Mask(a [16]byte, b [16]byte) uint16


// MaskTestEpi8Mask1: Compute the bitwise AND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
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
func MaskTestEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskTestEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskTestEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// TestEpi8Mask1: Compute the bitwise AND of packed 8-bit integers in 'a' and
// 'b', producing intermediate 8-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm256_test_epi8_mask'.
// Requires AVX512BW.
func TestEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(testEpi8Mask1([32]byte(a), [32]byte(b)))
}

func testEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskTestEpi8Mask2: Compute the bitwise AND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
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
func MaskTestEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskTestEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskTestEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


// TestEpi8Mask2: Compute the bitwise AND of packed 8-bit integers in 'a' and
// 'b', producing intermediate 8-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPTESTMB'. Intrinsic: '_mm512_test_epi8_mask'.
// Requires AVX512BW.
func TestEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(testEpi8Mask2([64]byte(a), [64]byte(b)))
}

func testEpi8Mask2(a [64]byte, b [64]byte) uint64


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
func MaskTestnEpi16Mask(k1 Mmask8, a M128i, b M128i) Mmask8 {
	return Mmask8(maskTestnEpi16Mask(uint8(k1), [16]byte(a), [16]byte(b)))
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
func TestnEpi16Mask(a M128i, b M128i) Mmask8 {
	return Mmask8(testnEpi16Mask([16]byte(a), [16]byte(b)))
}

func testnEpi16Mask(a [16]byte, b [16]byte) uint8


// MaskTestnEpi16Mask1: Compute the bitwise NAND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is zero. 
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
func MaskTestnEpi16Mask1(k1 Mmask16, a M256i, b M256i) Mmask16 {
	return Mmask16(maskTestnEpi16Mask1(uint16(k1), [32]byte(a), [32]byte(b)))
}

func maskTestnEpi16Mask1(k1 uint16, a [32]byte, b [32]byte) uint16


// TestnEpi16Mask1: Compute the bitwise NAND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 15
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm256_testn_epi16_mask'.
// Requires AVX512BW.
func TestnEpi16Mask1(a M256i, b M256i) Mmask16 {
	return Mmask16(testnEpi16Mask1([32]byte(a), [32]byte(b)))
}

func testnEpi16Mask1(a [32]byte, b [32]byte) uint16


// MaskTestnEpi16Mask2: Compute the bitwise NAND of packed 16-bit integers in
// 'a' and 'b', producing intermediate 16-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is zero. 
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
func MaskTestnEpi16Mask2(k1 Mmask32, a M512i, b M512i) Mmask32 {
	return Mmask32(maskTestnEpi16Mask2(uint32(k1), [64]byte(a), [64]byte(b)))
}

func maskTestnEpi16Mask2(k1 uint32, a [64]byte, b [64]byte) uint32


// TestnEpi16Mask2: Compute the bitwise NAND of packed 16-bit integers in 'a'
// and 'b', producing intermediate 16-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 31
//			i := j*16
//			k[j] := ((a[i+15:i] AND b[i+15:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTNMW'. Intrinsic: '_mm512_testn_epi16_mask'.
// Requires AVX512BW.
func TestnEpi16Mask2(a M512i, b M512i) Mmask32 {
	return Mmask32(testnEpi16Mask2([64]byte(a), [64]byte(b)))
}

func testnEpi16Mask2(a [64]byte, b [64]byte) uint32


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
func MaskTestnEpi8Mask(k1 Mmask16, a M128i, b M128i) Mmask16 {
	return Mmask16(maskTestnEpi8Mask(uint16(k1), [16]byte(a), [16]byte(b)))
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
func TestnEpi8Mask(a M128i, b M128i) Mmask16 {
	return Mmask16(testnEpi8Mask([16]byte(a), [16]byte(b)))
}

func testnEpi8Mask(a [16]byte, b [16]byte) uint16


// MaskTestnEpi8Mask1: Compute the bitwise NAND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// zero. 
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
func MaskTestnEpi8Mask1(k1 Mmask32, a M256i, b M256i) Mmask32 {
	return Mmask32(maskTestnEpi8Mask1(uint32(k1), [32]byte(a), [32]byte(b)))
}

func maskTestnEpi8Mask1(k1 uint32, a [32]byte, b [32]byte) uint32


// TestnEpi8Mask1: Compute the bitwise NAND of packed 8-bit integers in 'a' and
// 'b', producing intermediate 8-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 31
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:32] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm256_testn_epi8_mask'.
// Requires AVX512BW.
func TestnEpi8Mask1(a M256i, b M256i) Mmask32 {
	return Mmask32(testnEpi8Mask1([32]byte(a), [32]byte(b)))
}

func testnEpi8Mask1(a [32]byte, b [32]byte) uint32


// MaskTestnEpi8Mask2: Compute the bitwise NAND of packed 8-bit integers in 'a'
// and 'b', producing intermediate 8-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// zero. 
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
func MaskTestnEpi8Mask2(k1 Mmask64, a M512i, b M512i) Mmask64 {
	return Mmask64(maskTestnEpi8Mask2(uint64(k1), [64]byte(a), [64]byte(b)))
}

func maskTestnEpi8Mask2(k1 uint64, a [64]byte, b [64]byte) uint64


// TestnEpi8Mask2: Compute the bitwise NAND of packed 8-bit integers in 'a' and
// 'b', producing intermediate 8-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is zero. 
//
//		FOR j := 0 to 63
//			i := j*8
//			k[j] := ((a[i+7:i] AND b[i+7:i]) == 0) ? 1 : 0
//		ENDFOR
//		k[MAX:64] := 0
//
// Instruction: 'VPTESTNMB'. Intrinsic: '_mm512_testn_epi8_mask'.
// Requires AVX512BW.
func TestnEpi8Mask2(a M512i, b M512i) Mmask64 {
	return Mmask64(testnEpi8Mask2([64]byte(a), [64]byte(b)))
}

func testnEpi8Mask2(a [64]byte, b [64]byte) uint64


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
func MaskUnpackhiEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskUnpackhiEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzUnpackhiEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzUnpackhiEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackhiEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskUnpackhiEpi161: Unpack and interleave 16-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskUnpackhiEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskUnpackhiEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskUnpackhiEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzUnpackhiEpi161: Unpack and interleave 16-bit integers from the high
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
func MaskzUnpackhiEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzUnpackhiEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzUnpackhiEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskUnpackhiEpi162: Unpack and interleave 16-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskUnpackhiEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskUnpackhiEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskUnpackhiEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzUnpackhiEpi162: Unpack and interleave 16-bit integers from the high
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
func MaskzUnpackhiEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzUnpackhiEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzUnpackhiEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// UnpackhiEpi16: Unpack and interleave 16-bit integers from the high half of
// each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
func UnpackhiEpi16(a M512i, b M512i) M512i {
	return M512i(unpackhiEpi16([64]byte(a), [64]byte(b)))
}

func unpackhiEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskUnpackhiEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskUnpackhiEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzUnpackhiEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzUnpackhiEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackhiEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskUnpackhiEpi81: Unpack and interleave 8-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskUnpackhiEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskUnpackhiEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskUnpackhiEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzUnpackhiEpi81: Unpack and interleave 8-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzUnpackhiEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzUnpackhiEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzUnpackhiEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskUnpackhiEpi82: Unpack and interleave 8-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskUnpackhiEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskUnpackhiEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskUnpackhiEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzUnpackhiEpi82: Unpack and interleave 8-bit integers from the high half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzUnpackhiEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzUnpackhiEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzUnpackhiEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// UnpackhiEpi8: Unpack and interleave 8-bit integers from the high half of
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
func UnpackhiEpi8(a M512i, b M512i) M512i {
	return M512i(unpackhiEpi8([64]byte(a), [64]byte(b)))
}

func unpackhiEpi8(a [64]byte, b [64]byte) [64]byte


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
func MaskUnpackloEpi16(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskUnpackloEpi16([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzUnpackloEpi16(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzUnpackloEpi16(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackloEpi16(k uint8, a [16]byte, b [16]byte) [16]byte


// MaskUnpackloEpi161: Unpack and interleave 16-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskUnpackloEpi161(src M256i, k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskUnpackloEpi161([32]byte(src), uint16(k), [32]byte(a), [32]byte(b)))
}

func maskUnpackloEpi161(src [32]byte, k uint16, a [32]byte, b [32]byte) [32]byte


// MaskzUnpackloEpi161: Unpack and interleave 16-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzUnpackloEpi161(k Mmask16, a M256i, b M256i) M256i {
	return M256i(maskzUnpackloEpi161(uint16(k), [32]byte(a), [32]byte(b)))
}

func maskzUnpackloEpi161(k uint16, a [32]byte, b [32]byte) [32]byte


// MaskUnpackloEpi162: Unpack and interleave 16-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func MaskUnpackloEpi162(src M512i, k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskUnpackloEpi162([64]byte(src), uint32(k), [64]byte(a), [64]byte(b)))
}

func maskUnpackloEpi162(src [64]byte, k uint32, a [64]byte, b [64]byte) [64]byte


// MaskzUnpackloEpi162: Unpack and interleave 16-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzUnpackloEpi162(k Mmask32, a M512i, b M512i) M512i {
	return M512i(maskzUnpackloEpi162(uint32(k), [64]byte(a), [64]byte(b)))
}

func maskzUnpackloEpi162(k uint32, a [64]byte, b [64]byte) [64]byte


// UnpackloEpi16: Unpack and interleave 16-bit integers from the low half of
// each 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
func UnpackloEpi16(a M512i, b M512i) M512i {
	return M512i(unpackloEpi16([64]byte(a), [64]byte(b)))
}

func unpackloEpi16(a [64]byte, b [64]byte) [64]byte


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
func MaskUnpackloEpi8(src M128i, k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskUnpackloEpi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
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
func MaskzUnpackloEpi8(k Mmask16, a M128i, b M128i) M128i {
	return M128i(maskzUnpackloEpi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzUnpackloEpi8(k uint16, a [16]byte, b [16]byte) [16]byte


// MaskUnpackloEpi81: Unpack and interleave 8-bit integers from the low half of
// each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
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
func MaskUnpackloEpi81(src M256i, k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskUnpackloEpi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskUnpackloEpi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzUnpackloEpi81: Unpack and interleave 8-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzUnpackloEpi81(k Mmask32, a M256i, b M256i) M256i {
	return M256i(maskzUnpackloEpi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzUnpackloEpi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MaskUnpackloEpi82: Unpack and interleave 8-bit integers from the low half of
// each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
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
func MaskUnpackloEpi82(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskUnpackloEpi82([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskUnpackloEpi82(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzUnpackloEpi82: Unpack and interleave 8-bit integers from the low half
// of each 128-bit lane in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzUnpackloEpi82(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzUnpackloEpi82(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzUnpackloEpi82(k uint64, a [64]byte, b [64]byte) [64]byte


// UnpackloEpi8: Unpack and interleave 8-bit integers from the low half of each
// 128-bit lane in 'a' and 'b', and store the results in 'dst'. 
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
func UnpackloEpi8(a M512i, b M512i) M512i {
	return M512i(unpackloEpi8([64]byte(a), [64]byte(b)))
}

func unpackloEpi8(a [64]byte, b [64]byte) [64]byte

