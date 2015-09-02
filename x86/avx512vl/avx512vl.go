package avx512vl

import "github.com/klauspost/intrinsics/x86"


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
func MaskAddPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskAddPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskAddPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


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
func MaskzAddPd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzAddPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzAddPd(k uint8, a [2]float64, b [2]float64) [2]float64


// MaskAddPd1: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
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
func MaskAddPd1(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(maskAddPd1([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func maskAddPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// MaskzAddPd1: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func MaskzAddPd1(k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(maskzAddPd1(uint8(k), [4]float64(a), [4]float64(b)))
}

func maskzAddPd1(k uint8, a [4]float64, b [4]float64) [4]float64


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
func MaskAddPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskAddPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskAddPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


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
func MaskzAddPs(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzAddPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzAddPs(k uint8, a [4]float32, b [4]float32) [4]float32


// MaskAddPs1: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
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
func MaskAddPs1(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(maskAddPs1([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func maskAddPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// MaskzAddPs1: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func MaskzAddPs1(k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(maskzAddPs1(uint8(k), [8]float32(a), [8]float32(b)))
}

func maskzAddPs1(k uint8, a [8]float32, b [8]float32) [8]float32


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
func AlignrEpi32(a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(alignrEpi32([16]byte(a), [16]byte(b), count))
}

func alignrEpi32(a [16]byte, b [16]byte, count int) [16]byte


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
func MaskAlignrEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(maskAlignrEpi32([16]byte(src), uint8(k), [16]byte(a), [16]byte(b), count))
}

func maskAlignrEpi32(src [16]byte, k uint8, a [16]byte, b [16]byte, count int) [16]byte


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
func MaskzAlignrEpi32(k x86.Mmask8, a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(maskzAlignrEpi32(uint8(k), [16]byte(a), [16]byte(b), count))
}

func maskzAlignrEpi32(k uint8, a [16]byte, b [16]byte, count int) [16]byte


// AlignrEpi321: Concatenate 'a' and 'b' into a 64-byte immediate result, shift
// the result right by 'count' 32-bit elements, and store the low 32 bytes (8
// elements) in 'dst'. 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (32*count)
//		dst[255:0] := temp[255:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm256_alignr_epi32'.
// Requires AVX512VL.
func AlignrEpi321(a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(alignrEpi321([32]byte(a), [32]byte(b), count))
}

func alignrEpi321(a [32]byte, b [32]byte, count int) [32]byte


// MaskAlignrEpi321: Concatenate 'a' and 'b' into a 64-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 32
// bytes (8 elements) in 'dst' using writemask 'k' (elements are copied from
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
func MaskAlignrEpi321(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(maskAlignrEpi321([32]byte(src), uint8(k), [32]byte(a), [32]byte(b), count))
}

func maskAlignrEpi321(src [32]byte, k uint8, a [32]byte, b [32]byte, count int) [32]byte


// MaskzAlignrEpi321: Concatenate 'a' and 'b' into a 64-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 32
// bytes (8 elements) in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
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
func MaskzAlignrEpi321(k x86.Mmask8, a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(maskzAlignrEpi321(uint8(k), [32]byte(a), [32]byte(b), count))
}

func maskzAlignrEpi321(k uint8, a [32]byte, b [32]byte, count int) [32]byte


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
func AlignrEpi64(a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(alignrEpi64([16]byte(a), [16]byte(b), count))
}

func alignrEpi64(a [16]byte, b [16]byte, count int) [16]byte


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
func MaskAlignrEpi64(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(maskAlignrEpi64([16]byte(src), uint8(k), [16]byte(a), [16]byte(b), count))
}

func maskAlignrEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte, count int) [16]byte


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
func MaskzAlignrEpi64(k x86.Mmask8, a x86.M128i, b x86.M128i, count int) x86.M128i {
	return x86.M128i(maskzAlignrEpi64(uint8(k), [16]byte(a), [16]byte(b), count))
}

func maskzAlignrEpi64(k uint8, a [16]byte, b [16]byte, count int) [16]byte


// AlignrEpi641: Concatenate 'a' and 'b' into a 64-byte immediate result, shift
// the result right by 'count' 64-bit elements, and store the low 32 bytes (4
// elements) in 'dst'. 
//
//		temp[511:256] := a[255:0]
//		temp[255:0] := b[255:0]
//		temp[511:0] := temp[511:0] >> (64*count)
//		dst[255:0] := temp[255:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VALIGNQ'. Intrinsic: '_mm256_alignr_epi64'.
// Requires AVX512VL.
func AlignrEpi641(a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(alignrEpi641([32]byte(a), [32]byte(b), count))
}

func alignrEpi641(a [32]byte, b [32]byte, count int) [32]byte


// MaskAlignrEpi641: Concatenate 'a' and 'b' into a 64-byte immediate result,
// shift the result right by 'count' 64-bit elements, and store the low 32
// bytes (4 elements) in 'dst' using writemask 'k' (elements are copied from
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
func MaskAlignrEpi641(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(maskAlignrEpi641([32]byte(src), uint8(k), [32]byte(a), [32]byte(b), count))
}

func maskAlignrEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte, count int) [32]byte


// MaskzAlignrEpi641: Concatenate 'a' and 'b' into a 64-byte immediate result,
// shift the result right by 'count' 64-bit elements, and store the low 32
// bytes (4 elements) in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
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
func MaskzAlignrEpi641(k x86.Mmask8, a x86.M256i, b x86.M256i, count int) x86.M256i {
	return x86.M256i(maskzAlignrEpi641(uint8(k), [32]byte(a), [32]byte(b), count))
}

func maskzAlignrEpi641(k uint8, a [32]byte, b [32]byte, count int) [32]byte


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
func Madd52hiEpu64(a x86.M128i, b x86.M128i, c x86.M128i) x86.M128i {
	return x86.M128i(madd52hiEpu64([16]byte(a), [16]byte(b), [16]byte(c)))
}

func madd52hiEpu64(a [16]byte, b [16]byte, c [16]byte) [16]byte


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
func MaskMadd52hiEpu64(a x86.M128i, k x86.Mmask8, b x86.M128i, c x86.M128i) x86.M128i {
	return x86.M128i(maskMadd52hiEpu64([16]byte(a), uint8(k), [16]byte(b), [16]byte(c)))
}

func maskMadd52hiEpu64(a [16]byte, k uint8, b [16]byte, c [16]byte) [16]byte


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
func MaskzMadd52hiEpu64(k x86.Mmask8, a x86.M128i, b x86.M128i, c x86.M128i) x86.M128i {
	return x86.M128i(maskzMadd52hiEpu64(uint8(k), [16]byte(a), [16]byte(b), [16]byte(c)))
}

func maskzMadd52hiEpu64(k uint8, a [16]byte, b [16]byte, c [16]byte) [16]byte


// Madd52hiEpu641: Multiply packed unsigned 52-bit integers in each 64-bit
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
func Madd52hiEpu641(a x86.M256i, b x86.M256i, c x86.M256i) x86.M256i {
	return x86.M256i(madd52hiEpu641([32]byte(a), [32]byte(b), [32]byte(c)))
}

func madd52hiEpu641(a [32]byte, b [32]byte, c [32]byte) [32]byte


// MaskMadd52hiEpu641: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'a' when the corresponding mask bit
// is not set). 
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
func MaskMadd52hiEpu641(a x86.M256i, k x86.Mmask8, b x86.M256i, c x86.M256i) x86.M256i {
	return x86.M256i(maskMadd52hiEpu641([32]byte(a), uint8(k), [32]byte(b), [32]byte(c)))
}

func maskMadd52hiEpu641(a [32]byte, k uint8, b [32]byte, c [32]byte) [32]byte


// MaskzMadd52hiEpu641: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzMadd52hiEpu641(k x86.Mmask8, a x86.M256i, b x86.M256i, c x86.M256i) x86.M256i {
	return x86.M256i(maskzMadd52hiEpu641(uint8(k), [32]byte(a), [32]byte(b), [32]byte(c)))
}

func maskzMadd52hiEpu641(k uint8, a [32]byte, b [32]byte, c [32]byte) [32]byte


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
func Madd52loEpu64(a x86.M128i, b x86.M128i, c x86.M128i) x86.M128i {
	return x86.M128i(madd52loEpu64([16]byte(a), [16]byte(b), [16]byte(c)))
}

func madd52loEpu64(a [16]byte, b [16]byte, c [16]byte) [16]byte


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
func MaskMadd52loEpu64(a x86.M128i, k x86.Mmask8, b x86.M128i, c x86.M128i) x86.M128i {
	return x86.M128i(maskMadd52loEpu64([16]byte(a), uint8(k), [16]byte(b), [16]byte(c)))
}

func maskMadd52loEpu64(a [16]byte, k uint8, b [16]byte, c [16]byte) [16]byte


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
func MaskzMadd52loEpu64(k x86.Mmask8, a x86.M128i, b x86.M128i, c x86.M128i) x86.M128i {
	return x86.M128i(maskzMadd52loEpu64(uint8(k), [16]byte(a), [16]byte(b), [16]byte(c)))
}

func maskzMadd52loEpu64(k uint8, a [16]byte, b [16]byte, c [16]byte) [16]byte


// Madd52loEpu641: Multiply packed unsigned 52-bit integers in each 64-bit
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
func Madd52loEpu641(a x86.M256i, b x86.M256i, c x86.M256i) x86.M256i {
	return x86.M256i(madd52loEpu641([32]byte(a), [32]byte(b), [32]byte(c)))
}

func madd52loEpu641(a [32]byte, b [32]byte, c [32]byte) [32]byte


// MaskMadd52loEpu641: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'a' when the corresponding mask bit
// is not set). 
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
func MaskMadd52loEpu641(a x86.M256i, k x86.Mmask8, b x86.M256i, c x86.M256i) x86.M256i {
	return x86.M256i(maskMadd52loEpu641([32]byte(a), uint8(k), [32]byte(b), [32]byte(c)))
}

func maskMadd52loEpu641(a [32]byte, k uint8, b [32]byte, c [32]byte) [32]byte


// MaskzMadd52loEpu641: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func MaskzMadd52loEpu641(k x86.Mmask8, a x86.M256i, b x86.M256i, c x86.M256i) x86.M256i {
	return x86.M256i(maskzMadd52loEpu641(uint8(k), [32]byte(a), [32]byte(b), [32]byte(c)))
}

func maskzMadd52loEpu641(k uint8, a [32]byte, b [32]byte, c [32]byte) [32]byte


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
func MovmEpi8(k x86.Mmask16) x86.M128i {
	return x86.M128i(movmEpi8(uint16(k)))
}

func movmEpi8(k uint16) [16]byte


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
func MaskMultishiftEpi64Epi8(src x86.M128i, k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMultishiftEpi64Epi8([16]byte(src), uint16(k), [16]byte(a), [16]byte(b)))
}

func maskMultishiftEpi64Epi8(src [16]byte, k uint16, a [16]byte, b [16]byte) [16]byte


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
func MaskzMultishiftEpi64Epi8(k x86.Mmask16, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMultishiftEpi64Epi8(uint16(k), [16]byte(a), [16]byte(b)))
}

func maskzMultishiftEpi64Epi8(k uint16, a [16]byte, b [16]byte) [16]byte


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
func MultishiftEpi64Epi8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(multishiftEpi64Epi8([16]byte(a), [16]byte(b)))
}

func multishiftEpi64Epi8(a [16]byte, b [16]byte) [16]byte


// MaskMultishiftEpi64Epi81: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
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
func MaskMultishiftEpi64Epi81(src x86.M256i, k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(maskMultishiftEpi64Epi81([32]byte(src), uint32(k), [32]byte(a), [32]byte(b)))
}

func maskMultishiftEpi64Epi81(src [32]byte, k uint32, a [32]byte, b [32]byte) [32]byte


// MaskzMultishiftEpi64Epi81: For each 64-bit element in 'b', select 8
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
func MaskzMultishiftEpi64Epi81(k x86.Mmask32, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(maskzMultishiftEpi64Epi81(uint32(k), [32]byte(a), [32]byte(b)))
}

func maskzMultishiftEpi64Epi81(k uint32, a [32]byte, b [32]byte) [32]byte


// MultishiftEpi64Epi81: For each 64-bit element in 'b', select 8 unaligned
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
func MultishiftEpi64Epi81(a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(multishiftEpi64Epi81([32]byte(a), [32]byte(b)))
}

func multishiftEpi64Epi81(a [32]byte, b [32]byte) [32]byte


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
func MaskPermutex2varEpi8(a x86.M128i, k x86.Mmask16, idx x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskPermutex2varEpi8([16]byte(a), uint16(k), [16]byte(idx), [16]byte(b)))
}

func maskPermutex2varEpi8(a [16]byte, k uint16, idx [16]byte, b [16]byte) [16]byte


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
func Mask2Permutex2varEpi8(a x86.M128i, idx x86.M128i, k x86.Mmask16, b x86.M128i) x86.M128i {
	return x86.M128i(mask2Permutex2varEpi8([16]byte(a), [16]byte(idx), uint16(k), [16]byte(b)))
}

func mask2Permutex2varEpi8(a [16]byte, idx [16]byte, k uint16, b [16]byte) [16]byte


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
func MaskzPermutex2varEpi8(k x86.Mmask16, a x86.M128i, idx x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzPermutex2varEpi8(uint16(k), [16]byte(a), [16]byte(idx), [16]byte(b)))
}

func maskzPermutex2varEpi8(k uint16, a [16]byte, idx [16]byte, b [16]byte) [16]byte


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
func Permutex2varEpi8(a x86.M128i, idx x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(permutex2varEpi8([16]byte(a), [16]byte(idx), [16]byte(b)))
}

func permutex2varEpi8(a [16]byte, idx [16]byte, b [16]byte) [16]byte


// MaskPermutex2varEpi81: Shuffle 8-bit integers in 'a' and 'b' across lanes
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
func MaskPermutex2varEpi81(a x86.M256i, k x86.Mmask32, idx x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(maskPermutex2varEpi81([32]byte(a), uint32(k), [32]byte(idx), [32]byte(b)))
}

func maskPermutex2varEpi81(a [32]byte, k uint32, idx [32]byte, b [32]byte) [32]byte


// Mask2Permutex2varEpi81: Shuffle 8-bit integers in 'a' and 'b' across lanes
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
// Instruction: 'VPERMI2B'. Intrinsic: '_mm256_mask2_permutex2var_epi8'.
// Requires AVX512VL.
func Mask2Permutex2varEpi81(a x86.M256i, idx x86.M256i, k x86.Mmask32, b x86.M256i) x86.M256i {
	return x86.M256i(mask2Permutex2varEpi81([32]byte(a), [32]byte(idx), uint32(k), [32]byte(b)))
}

func mask2Permutex2varEpi81(a [32]byte, idx [32]byte, k uint32, b [32]byte) [32]byte


// MaskzPermutex2varEpi81: Shuffle 8-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
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
func MaskzPermutex2varEpi81(k x86.Mmask32, a x86.M256i, idx x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(maskzPermutex2varEpi81(uint32(k), [32]byte(a), [32]byte(idx), [32]byte(b)))
}

func maskzPermutex2varEpi81(k uint32, a [32]byte, idx [32]byte, b [32]byte) [32]byte


// Permutex2varEpi81: Shuffle 8-bit integers in 'a' and 'b' across lanes using
// the corresponding selector and index in 'idx', and store the results in
// 'dst'. 
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
func Permutex2varEpi81(a x86.M256i, idx x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(permutex2varEpi81([32]byte(a), [32]byte(idx), [32]byte(b)))
}

func permutex2varEpi81(a [32]byte, idx [32]byte, b [32]byte) [32]byte


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
func MaskPermutexvarEpi8(src x86.M128i, k x86.Mmask16, idx x86.M128i, a x86.M128i) x86.M128i {
	return x86.M128i(maskPermutexvarEpi8([16]byte(src), uint16(k), [16]byte(idx), [16]byte(a)))
}

func maskPermutexvarEpi8(src [16]byte, k uint16, idx [16]byte, a [16]byte) [16]byte


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
func MaskzPermutexvarEpi8(k x86.Mmask16, idx x86.M128i, a x86.M128i) x86.M128i {
	return x86.M128i(maskzPermutexvarEpi8(uint16(k), [16]byte(idx), [16]byte(a)))
}

func maskzPermutexvarEpi8(k uint16, idx [16]byte, a [16]byte) [16]byte


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
func PermutexvarEpi8(idx x86.M128i, a x86.M128i) x86.M128i {
	return x86.M128i(permutexvarEpi8([16]byte(idx), [16]byte(a)))
}

func permutexvarEpi8(idx [16]byte, a [16]byte) [16]byte


// MaskPermutexvarEpi81: Shuffle 8-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
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
func MaskPermutexvarEpi81(src x86.M256i, k x86.Mmask32, idx x86.M256i, a x86.M256i) x86.M256i {
	return x86.M256i(maskPermutexvarEpi81([32]byte(src), uint32(k), [32]byte(idx), [32]byte(a)))
}

func maskPermutexvarEpi81(src [32]byte, k uint32, idx [32]byte, a [32]byte) [32]byte


// MaskzPermutexvarEpi81: Shuffle 8-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
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
func MaskzPermutexvarEpi81(k x86.Mmask32, idx x86.M256i, a x86.M256i) x86.M256i {
	return x86.M256i(maskzPermutexvarEpi81(uint32(k), [32]byte(idx), [32]byte(a)))
}

func maskzPermutexvarEpi81(k uint32, idx [32]byte, a [32]byte) [32]byte


// PermutexvarEpi81: Shuffle 8-bit integers in 'a' across lanes using the
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
func PermutexvarEpi81(idx x86.M256i, a x86.M256i) x86.M256i {
	return x86.M256i(permutexvarEpi81([32]byte(idx), [32]byte(a)))
}

func permutexvarEpi81(idx [32]byte, a [32]byte) [32]byte

