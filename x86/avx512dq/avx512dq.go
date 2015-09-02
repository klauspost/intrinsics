package avx512dq

import . "github.com/klauspost/intrinsics/x86"


// MaskAndPd: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm_mask_and_pd'.
// Requires AVX512DQ.
func MaskAndPd(src M128d, k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskAndPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskAndPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


// MaskzAndPd: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm_maskz_and_pd'.
// Requires AVX512DQ.
func MaskzAndPd(k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskzAndPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzAndPd(k uint8, a [2]float64, b [2]float64) [2]float64


// MaskAndPd1: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm256_mask_and_pd'.
// Requires AVX512DQ.
func MaskAndPd1(src M256d, k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskAndPd1([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func maskAndPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// MaskzAndPd1: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := 0 
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm256_maskz_and_pd'.
// Requires AVX512DQ.
func MaskzAndPd1(k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskzAndPd1(uint8(k), [4]float64(a), [4]float64(b)))
}

func maskzAndPd1(k uint8, a [4]float64, b [4]float64) [4]float64


// AndPd: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm512_and_pd'.
// Requires AVX512DQ.
func AndPd(a M512d, b M512d) M512d {
	return M512d(andPd([8]float64(a), [8]float64(b)))
}

func andPd(a [8]float64, b [8]float64) [8]float64


// MaskAndPd2: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm512_mask_and_pd'.
// Requires AVX512DQ.
func MaskAndPd2(src M512d, k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskAndPd2([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskAndPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// MaskzAndPd2: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm512_maskz_and_pd'.
// Requires AVX512DQ.
func MaskzAndPd2(k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskzAndPd2(uint8(k), [8]float64(a), [8]float64(b)))
}

func maskzAndPd2(k uint8, a [8]float64, b [8]float64) [8]float64


// MaskAndPs: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm_mask_and_ps'.
// Requires AVX512DQ.
func MaskAndPs(src M128, k Mmask8, a M128, b M128) M128 {
	return M128(maskAndPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskAndPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


// MaskzAndPs: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm_maskz_and_ps'.
// Requires AVX512DQ.
func MaskzAndPs(k Mmask8, a M128, b M128) M128 {
	return M128(maskzAndPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzAndPs(k uint8, a [4]float32, b [4]float32) [4]float32


// MaskAndPs1: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm256_mask_and_ps'.
// Requires AVX512DQ.
func MaskAndPs1(src M256, k Mmask8, a M256, b M256) M256 {
	return M256(maskAndPs1([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func maskAndPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// MaskzAndPs1: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm256_maskz_and_ps'.
// Requires AVX512DQ.
func MaskzAndPs1(k Mmask8, a M256, b M256) M256 {
	return M256(maskzAndPs1(uint8(k), [8]float32(a), [8]float32(b)))
}

func maskzAndPs1(k uint8, a [8]float32, b [8]float32) [8]float32


// AndPs: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm512_and_ps'.
// Requires AVX512DQ.
func AndPs(a M512, b M512) M512 {
	return M512(andPs([16]float32(a), [16]float32(b)))
}

func andPs(a [16]float32, b [16]float32) [16]float32


// MaskAndPs2: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm512_mask_and_ps'.
// Requires AVX512DQ.
func MaskAndPs2(src M512, k Mmask16, a M512, b M512) M512 {
	return M512(maskAndPs2([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskAndPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// MaskzAndPs2: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm512_maskz_and_ps'.
// Requires AVX512DQ.
func MaskzAndPs2(k Mmask16, a M512, b M512) M512 {
	return M512(maskzAndPs2(uint16(k), [16]float32(a), [16]float32(b)))
}

func maskzAndPs2(k uint16, a [16]float32, b [16]float32) [16]float32


// MaskAndnotPd: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm_mask_andnot_pd'.
// Requires AVX512DQ.
func MaskAndnotPd(src M128d, k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskAndnotPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskAndnotPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


// MaskzAndnotPd: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm_maskz_andnot_pd'.
// Requires AVX512DQ.
func MaskzAndnotPd(k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskzAndnotPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzAndnotPd(k uint8, a [2]float64, b [2]float64) [2]float64


// MaskAndnotPd1: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm256_mask_andnot_pd'.
// Requires AVX512DQ.
func MaskAndnotPd1(src M256d, k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskAndnotPd1([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func maskAndnotPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// MaskzAndnotPd1: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm256_maskz_andnot_pd'.
// Requires AVX512DQ.
func MaskzAndnotPd1(k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskzAndnotPd1(uint8(k), [4]float64(a), [4]float64(b)))
}

func maskzAndnotPd1(k uint8, a [4]float64, b [4]float64) [4]float64


// AndnotPd: Compute the bitwise AND NOT of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm512_andnot_pd'.
// Requires AVX512DQ.
func AndnotPd(a M512d, b M512d) M512d {
	return M512d(andnotPd([8]float64(a), [8]float64(b)))
}

func andnotPd(a [8]float64, b [8]float64) [8]float64


// MaskAndnotPd2: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm512_mask_andnot_pd'.
// Requires AVX512DQ.
func MaskAndnotPd2(src M512d, k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskAndnotPd2([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskAndnotPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// MaskzAndnotPd2: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm512_maskz_andnot_pd'.
// Requires AVX512DQ.
func MaskzAndnotPd2(k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskzAndnotPd2(uint8(k), [8]float64(a), [8]float64(b)))
}

func maskzAndnotPd2(k uint8, a [8]float64, b [8]float64) [8]float64


// MaskAndnotPs: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm_mask_andnot_ps'.
// Requires AVX512DQ.
func MaskAndnotPs(src M128, k Mmask8, a M128, b M128) M128 {
	return M128(maskAndnotPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskAndnotPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


// MaskzAndnotPs: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm_maskz_andnot_ps'.
// Requires AVX512DQ.
func MaskzAndnotPs(k Mmask8, a M128, b M128) M128 {
	return M128(maskzAndnotPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzAndnotPs(k uint8, a [4]float32, b [4]float32) [4]float32


// MaskAndnotPs1: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm256_mask_andnot_ps'.
// Requires AVX512DQ.
func MaskAndnotPs1(src M256, k Mmask8, a M256, b M256) M256 {
	return M256(maskAndnotPs1([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func maskAndnotPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// MaskzAndnotPs1: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm256_maskz_andnot_ps'.
// Requires AVX512DQ.
func MaskzAndnotPs1(k Mmask8, a M256, b M256) M256 {
	return M256(maskzAndnotPs1(uint8(k), [8]float32(a), [8]float32(b)))
}

func maskzAndnotPs1(k uint8, a [8]float32, b [8]float32) [8]float32


// AndnotPs: Compute the bitwise AND NOT of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm512_andnot_ps'.
// Requires AVX512DQ.
func AndnotPs(a M512, b M512) M512 {
	return M512(andnotPs([16]float32(a), [16]float32(b)))
}

func andnotPs(a [16]float32, b [16]float32) [16]float32


// MaskAndnotPs2: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm512_mask_andnot_ps'.
// Requires AVX512DQ.
func MaskAndnotPs2(src M512, k Mmask16, a M512, b M512) M512 {
	return M512(maskAndnotPs2([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskAndnotPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// MaskzAndnotPs2: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm512_maskz_andnot_ps'.
// Requires AVX512DQ.
func MaskzAndnotPs2(k Mmask16, a M512, b M512) M512 {
	return M512(maskzAndnotPs2(uint16(k), [16]float32(a), [16]float32(b)))
}

func maskzAndnotPs2(k uint16, a [16]float32, b [16]float32) [16]float32


// BroadcastF32x2: Broadcast the lower 2 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := (j mod 2)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF32X2'. Intrinsic: '_mm256_broadcast_f32x2'.
// Requires AVX512DQ.
func BroadcastF32x2(a M128) M256 {
	return M256(broadcastF32x2([4]float32(a)))
}

func broadcastF32x2(a [4]float32) [8]float32


// MaskBroadcastF32x2: Broadcast the lower 2 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF32X2'. Intrinsic: '_mm256_mask_broadcast_f32x2'.
// Requires AVX512DQ.
func MaskBroadcastF32x2(src M256, k Mmask8, a M128) M256 {
	return M256(maskBroadcastF32x2([8]float32(src), uint8(k), [4]float32(a)))
}

func maskBroadcastF32x2(src [8]float32, k uint8, a [4]float32) [8]float32


// MaskzBroadcastF32x2: Broadcast the lower 2 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF32X2'. Intrinsic: '_mm256_maskz_broadcast_f32x2'.
// Requires AVX512DQ.
func MaskzBroadcastF32x2(k Mmask8, a M128) M256 {
	return M256(maskzBroadcastF32x2(uint8(k), [4]float32(a)))
}

func maskzBroadcastF32x2(k uint8, a [4]float32) [8]float32


// BroadcastF32x21: Broadcast the lower 2 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 2)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF32X2'. Intrinsic: '_mm512_broadcast_f32x2'.
// Requires AVX512DQ.
func BroadcastF32x21(a M128) M512 {
	return M512(broadcastF32x21([4]float32(a)))
}

func broadcastF32x21(a [4]float32) [16]float32


// MaskBroadcastF32x21: Broadcast the lower 2 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF32X2'. Intrinsic: '_mm512_mask_broadcast_f32x2'.
// Requires AVX512DQ.
func MaskBroadcastF32x21(src M512, k Mmask16, a M128) M512 {
	return M512(maskBroadcastF32x21([16]float32(src), uint16(k), [4]float32(a)))
}

func maskBroadcastF32x21(src [16]float32, k uint16, a [4]float32) [16]float32


// MaskzBroadcastF32x21: Broadcast the lower 2 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF32X2'. Intrinsic: '_mm512_maskz_broadcast_f32x2'.
// Requires AVX512DQ.
func MaskzBroadcastF32x21(k Mmask16, a M128) M512 {
	return M512(maskzBroadcastF32x21(uint16(k), [4]float32(a)))
}

func maskzBroadcastF32x21(k uint16, a [4]float32) [16]float32


// BroadcastF32x8: Broadcast the 8 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 8)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF32X8'. Intrinsic: '_mm512_broadcast_f32x8'.
// Requires AVX512DQ.
func BroadcastF32x8(a M256) M512 {
	return M512(broadcastF32x8([8]float32(a)))
}

func broadcastF32x8(a [8]float32) [16]float32


// MaskBroadcastF32x8: Broadcast the 8 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 8)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF32X8'. Intrinsic: '_mm512_mask_broadcast_f32x8'.
// Requires AVX512DQ.
func MaskBroadcastF32x8(src M512, k Mmask16, a M256) M512 {
	return M512(maskBroadcastF32x8([16]float32(src), uint16(k), [8]float32(a)))
}

func maskBroadcastF32x8(src [16]float32, k uint16, a [8]float32) [16]float32


// MaskzBroadcastF32x8: Broadcast the 8 packed single-precision (32-bit)
// floating-point elements from 'a' to all elements of 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 8)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF32X8'. Intrinsic: '_mm512_maskz_broadcast_f32x8'.
// Requires AVX512DQ.
func MaskzBroadcastF32x8(k Mmask16, a M256) M512 {
	return M512(maskzBroadcastF32x8(uint16(k), [8]float32(a)))
}

func maskzBroadcastF32x8(k uint16, a [8]float32) [16]float32


// BroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
// floating-point elements from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			n := (j mod 2)*64
//			dst[i+63:i] := a[n+63:n]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF64X2'. Intrinsic: '_mm256_broadcast_f64x2'.
// Requires AVX512DQ.
func BroadcastF64x2(a M128d) M256d {
	return M256d(broadcastF64x2([2]float64(a)))
}

func broadcastF64x2(a [2]float64) [4]float64


// MaskBroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
// floating-point elements from 'a' to all elements of 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := src[n+63:n]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF64X2'. Intrinsic: '_mm256_mask_broadcast_f64x2'.
// Requires AVX512DQ.
func MaskBroadcastF64x2(src M256d, k Mmask8, a M128d) M256d {
	return M256d(maskBroadcastF64x2([4]float64(src), uint8(k), [2]float64(a)))
}

func maskBroadcastF64x2(src [4]float64, k uint8, a [2]float64) [4]float64


// MaskzBroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
// floating-point elements from 'a' to all elements of 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF64X2'. Intrinsic: '_mm256_maskz_broadcast_f64x2'.
// Requires AVX512DQ.
func MaskzBroadcastF64x2(k Mmask8, a M128d) M256d {
	return M256d(maskzBroadcastF64x2(uint8(k), [2]float64(a)))
}

func maskzBroadcastF64x2(k uint8, a [2]float64) [4]float64


// BroadcastF64x21: Broadcast the 2 packed double-precision (64-bit)
// floating-point elements from 'a' to all elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			n := (j mod 2)*64
//			dst[i+63:i] := a[n+63:n]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF64X2'. Intrinsic: '_mm512_broadcast_f64x2'.
// Requires AVX512DQ.
func BroadcastF64x21(a M128d) M512d {
	return M512d(broadcastF64x21([2]float64(a)))
}

func broadcastF64x21(a [2]float64) [8]float64


// MaskBroadcastF64x21: Broadcast the 2 packed double-precision (64-bit)
// floating-point elements from 'a' to all elements of 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := src[n+63:n]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF64X2'. Intrinsic: '_mm512_mask_broadcast_f64x2'.
// Requires AVX512DQ.
func MaskBroadcastF64x21(src M512d, k Mmask8, a M128d) M512d {
	return M512d(maskBroadcastF64x21([8]float64(src), uint8(k), [2]float64(a)))
}

func maskBroadcastF64x21(src [8]float64, k uint8, a [2]float64) [8]float64


// MaskzBroadcastF64x21: Broadcast the 2 packed double-precision (64-bit)
// floating-point elements from 'a' to all elements of 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTF64X2'. Intrinsic: '_mm512_maskz_broadcast_f64x2'.
// Requires AVX512DQ.
func MaskzBroadcastF64x21(k Mmask8, a M128d) M512d {
	return M512d(maskzBroadcastF64x21(uint8(k), [2]float64(a)))
}

func maskzBroadcastF64x21(k uint8, a [2]float64) [8]float64


// BroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from 'a' to all
// elements of "dst. 
//
//		FOR j := 0 to 3
//			i := j*32
//			n := (j mod 2)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm_broadcast_i32x2'.
// Requires AVX512DQ.
func BroadcastI32x2(a M128i) M128i {
	return M128i(broadcastI32x2([16]byte(a)))
}

func broadcastI32x2(a [16]byte) [16]byte


// MaskBroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[n+31:n]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm_mask_broadcast_i32x2'.
// Requires AVX512DQ.
func MaskBroadcastI32x2(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskBroadcastI32x2([16]byte(src), uint8(k), [16]byte(a)))
}

func maskBroadcastI32x2(src [16]byte, k uint8, a [16]byte) [16]byte


// MaskzBroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from 'a'
// to all elements of 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm_maskz_broadcast_i32x2'.
// Requires AVX512DQ.
func MaskzBroadcastI32x2(k Mmask8, a M128i) M128i {
	return M128i(maskzBroadcastI32x2(uint8(k), [16]byte(a)))
}

func maskzBroadcastI32x2(k uint8, a [16]byte) [16]byte


// BroadcastI32x21: Broadcast the lower 2 packed 32-bit integers from 'a' to
// all elements of "dst. 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := (j mod 2)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm256_broadcast_i32x2'.
// Requires AVX512DQ.
func BroadcastI32x21(a M128i) M256i {
	return M256i(broadcastI32x21([16]byte(a)))
}

func broadcastI32x21(a [16]byte) [32]byte


// MaskBroadcastI32x21: Broadcast the lower 2 packed 32-bit integers from 'a'
// to all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[n+31:n]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm256_mask_broadcast_i32x2'.
// Requires AVX512DQ.
func MaskBroadcastI32x21(src M256i, k Mmask8, a M128i) M256i {
	return M256i(maskBroadcastI32x21([32]byte(src), uint8(k), [16]byte(a)))
}

func maskBroadcastI32x21(src [32]byte, k uint8, a [16]byte) [32]byte


// MaskzBroadcastI32x21: Broadcast the lower 2 packed 32-bit integers from 'a'
// to all elements of 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm256_maskz_broadcast_i32x2'.
// Requires AVX512DQ.
func MaskzBroadcastI32x21(k Mmask8, a M128i) M256i {
	return M256i(maskzBroadcastI32x21(uint8(k), [16]byte(a)))
}

func maskzBroadcastI32x21(k uint8, a [16]byte) [32]byte


// BroadcastI32x22: Broadcast the lower 2 packed 32-bit integers from 'a' to
// all elements of "dst. 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 2)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm512_broadcast_i32x2'.
// Requires AVX512DQ.
func BroadcastI32x22(a M128i) M512i {
	return M512i(broadcastI32x22([16]byte(a)))
}

func broadcastI32x22(a [16]byte) [64]byte


// MaskBroadcastI32x22: Broadcast the lower 2 packed 32-bit integers from 'a'
// to all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[n+31:n]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm512_mask_broadcast_i32x2'.
// Requires AVX512DQ.
func MaskBroadcastI32x22(src M512i, k Mmask16, a M128i) M512i {
	return M512i(maskBroadcastI32x22([64]byte(src), uint16(k), [16]byte(a)))
}

func maskBroadcastI32x22(src [64]byte, k uint16, a [16]byte) [64]byte


// MaskzBroadcastI32x22: Broadcast the lower 2 packed 32-bit integers from 'a'
// to all elements of 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 2)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI32X2'. Intrinsic: '_mm512_maskz_broadcast_i32x2'.
// Requires AVX512DQ.
func MaskzBroadcastI32x22(k Mmask16, a M128i) M512i {
	return M512i(maskzBroadcastI32x22(uint16(k), [16]byte(a)))
}

func maskzBroadcastI32x22(k uint16, a [16]byte) [64]byte


// BroadcastI32x8: Broadcast the 8 packed 32-bit integers from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 8)*32
//			dst[i+31:i] := a[n+31:n]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI32X8'. Intrinsic: '_mm512_broadcast_i32x8'.
// Requires AVX512DQ.
func BroadcastI32x8(a M256i) M512i {
	return M512i(broadcastI32x8([32]byte(a)))
}

func broadcastI32x8(a [32]byte) [64]byte


// MaskBroadcastI32x8: Broadcast the 8 packed 32-bit integers from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 8)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := src[n+31:n]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI32X8'. Intrinsic: '_mm512_mask_broadcast_i32x8'.
// Requires AVX512DQ.
func MaskBroadcastI32x8(src M512i, k Mmask16, a M256i) M512i {
	return M512i(maskBroadcastI32x8([64]byte(src), uint16(k), [32]byte(a)))
}

func maskBroadcastI32x8(src [64]byte, k uint16, a [32]byte) [64]byte


// MaskzBroadcastI32x8: Broadcast the 8 packed 32-bit integers from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			n := (j mod 8)*32
//			IF k[j]
//				dst[i+31:i] := a[n+31:n]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI32X8'. Intrinsic: '_mm512_maskz_broadcast_i32x8'.
// Requires AVX512DQ.
func MaskzBroadcastI32x8(k Mmask16, a M256i) M512i {
	return M512i(maskzBroadcastI32x8(uint16(k), [32]byte(a)))
}

func maskzBroadcastI32x8(k uint16, a [32]byte) [64]byte


// BroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			n := (j mod 2)*64
//			dst[i+63:i] := a[n+63:n]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI64X2'. Intrinsic: '_mm256_broadcast_i64x2'.
// Requires AVX512DQ.
func BroadcastI64x2(a M128i) M256i {
	return M256i(broadcastI64x2([16]byte(a)))
}

func broadcastI64x2(a [16]byte) [32]byte


// MaskBroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := src[n+63:n]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI64X2'. Intrinsic: '_mm256_mask_broadcast_i64x2'.
// Requires AVX512DQ.
func MaskBroadcastI64x2(src M256i, k Mmask8, a M128i) M256i {
	return M256i(maskBroadcastI64x2([32]byte(src), uint8(k), [16]byte(a)))
}

func maskBroadcastI64x2(src [32]byte, k uint8, a [16]byte) [32]byte


// MaskzBroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTI64X2'. Intrinsic: '_mm256_maskz_broadcast_i64x2'.
// Requires AVX512DQ.
func MaskzBroadcastI64x2(k Mmask8, a M128i) M256i {
	return M256i(maskzBroadcastI64x2(uint8(k), [16]byte(a)))
}

func maskzBroadcastI64x2(k uint8, a [16]byte) [32]byte


// BroadcastI64x21: Broadcast the 2 packed 64-bit integers from 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			n := (j mod 2)*64
//			dst[i+63:i] := a[n+63:n]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI64X2'. Intrinsic: '_mm512_broadcast_i64x2'.
// Requires AVX512DQ.
func BroadcastI64x21(a M128i) M512i {
	return M512i(broadcastI64x21([16]byte(a)))
}

func broadcastI64x21(a [16]byte) [64]byte


// MaskBroadcastI64x21: Broadcast the 2 packed 64-bit integers from 'a' to all
// elements of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := src[n+63:n]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI64X2'. Intrinsic: '_mm512_mask_broadcast_i64x2'.
// Requires AVX512DQ.
func MaskBroadcastI64x21(src M512i, k Mmask8, a M128i) M512i {
	return M512i(maskBroadcastI64x21([64]byte(src), uint8(k), [16]byte(a)))
}

func maskBroadcastI64x21(src [64]byte, k uint8, a [16]byte) [64]byte


// MaskzBroadcastI64x21: Broadcast the 2 packed 64-bit integers from 'a' to all
// elements of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			n := (j mod 2)*64
//			IF k[j]
//				dst[i+63:i] := a[n+63:n]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBROADCASTI64X2'. Intrinsic: '_mm512_maskz_broadcast_i64x2'.
// Requires AVX512DQ.
func MaskzBroadcastI64x21(k Mmask8, a M128i) M512i {
	return M512i(maskzBroadcastI64x21(uint8(k), [16]byte(a)))
}

func maskzBroadcastI64x21(k uint8, a [16]byte) [64]byte


// CvtRoundepi64Pd: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm512_cvt_roundepi64_pd'.
// Requires AVX512DQ.
func CvtRoundepi64Pd(a M512i, rounding int) M512d {
	return M512d(cvtRoundepi64Pd([64]byte(a), rounding))
}

func cvtRoundepi64Pd(a [64]byte, rounding int) [8]float64


// MaskCvtRoundepi64Pd: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm512_mask_cvt_roundepi64_pd'.
// Requires AVX512DQ.
func MaskCvtRoundepi64Pd(src M512d, k Mmask8, a M512i, rounding int) M512d {
	return M512d(maskCvtRoundepi64Pd([8]float64(src), uint8(k), [64]byte(a), rounding))
}

func maskCvtRoundepi64Pd(src [8]float64, k uint8, a [64]byte, rounding int) [8]float64


// MaskzCvtRoundepi64Pd: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm512_maskz_cvt_roundepi64_pd'.
// Requires AVX512DQ.
func MaskzCvtRoundepi64Pd(k Mmask8, a M512i, rounding int) M512d {
	return M512d(maskzCvtRoundepi64Pd(uint8(k), [64]byte(a), rounding))
}

func maskzCvtRoundepi64Pd(k uint8, a [64]byte, rounding int) [8]float64


// CvtRoundepi64Ps: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm512_cvt_roundepi64_ps'.
// Requires AVX512DQ.
func CvtRoundepi64Ps(a M512i, rounding int) M256 {
	return M256(cvtRoundepi64Ps([64]byte(a), rounding))
}

func cvtRoundepi64Ps(a [64]byte, rounding int) [8]float32


// MaskCvtRoundepi64Ps: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm512_mask_cvt_roundepi64_ps'.
// Requires AVX512DQ.
func MaskCvtRoundepi64Ps(src M256, k Mmask8, a M512i, rounding int) M256 {
	return M256(maskCvtRoundepi64Ps([8]float32(src), uint8(k), [64]byte(a), rounding))
}

func maskCvtRoundepi64Ps(src [8]float32, k uint8, a [64]byte, rounding int) [8]float32


// MaskzCvtRoundepi64Ps: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm512_maskz_cvt_roundepi64_ps'.
// Requires AVX512DQ.
func MaskzCvtRoundepi64Ps(k Mmask8, a M512i, rounding int) M256 {
	return M256(maskzCvtRoundepi64Ps(uint8(k), [64]byte(a), rounding))
}

func maskzCvtRoundepi64Ps(k uint8, a [64]byte, rounding int) [8]float32


// CvtRoundepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm512_cvt_roundepu64_pd'.
// Requires AVX512DQ.
func CvtRoundepu64Pd(a M512i, rounding int) M512d {
	return M512d(cvtRoundepu64Pd([64]byte(a), rounding))
}

func cvtRoundepu64Pd(a [64]byte, rounding int) [8]float64


// MaskCvtRoundepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
// packed double-precision (64-bit) floating-point elements, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm512_mask_cvt_roundepu64_pd'.
// Requires AVX512DQ.
func MaskCvtRoundepu64Pd(src M512d, k Mmask8, a M512i, rounding int) M512d {
	return M512d(maskCvtRoundepu64Pd([8]float64(src), uint8(k), [64]byte(a), rounding))
}

func maskCvtRoundepu64Pd(src [8]float64, k uint8, a [64]byte, rounding int) [8]float64


// MaskzCvtRoundepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
// packed double-precision (64-bit) floating-point elements, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm512_maskz_cvt_roundepu64_pd'.
// Requires AVX512DQ.
func MaskzCvtRoundepu64Pd(k Mmask8, a M512i, rounding int) M512d {
	return M512d(maskzCvtRoundepu64Pd(uint8(k), [64]byte(a), rounding))
}

func maskzCvtRoundepu64Pd(k uint8, a [64]byte, rounding int) [8]float64


// CvtRoundepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm512_cvt_roundepu64_ps'.
// Requires AVX512DQ.
func CvtRoundepu64Ps(a M512i, rounding int) M256 {
	return M256(cvtRoundepu64Ps([64]byte(a), rounding))
}

func cvtRoundepu64Ps(a [64]byte, rounding int) [8]float32


// MaskCvtRoundepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
// packed single-precision (32-bit) floating-point elements, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm512_mask_cvt_roundepu64_ps'.
// Requires AVX512DQ.
func MaskCvtRoundepu64Ps(src M256, k Mmask8, a M512i, rounding int) M256 {
	return M256(maskCvtRoundepu64Ps([8]float32(src), uint8(k), [64]byte(a), rounding))
}

func maskCvtRoundepu64Ps(src [8]float32, k uint8, a [64]byte, rounding int) [8]float32


// MaskzCvtRoundepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
// packed single-precision (32-bit) floating-point elements, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm512_maskz_cvt_roundepu64_ps'.
// Requires AVX512DQ.
func MaskzCvtRoundepu64Ps(k Mmask8, a M512i, rounding int) M256 {
	return M256(maskzCvtRoundepu64Ps(uint8(k), [64]byte(a), rounding))
}

func maskzCvtRoundepu64Ps(k uint8, a [64]byte, rounding int) [8]float32


// CvtRoundpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_cvt_roundpd_epi64'.
// Requires AVX512DQ.
func CvtRoundpdEpi64(a M512d, rounding int) M512i {
	return M512i(cvtRoundpdEpi64([8]float64(a), rounding))
}

func cvtRoundpdEpi64(a [8]float64, rounding int) [64]byte


// MaskCvtRoundpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_mask_cvt_roundpd_epi64'.
// Requires AVX512DQ.
func MaskCvtRoundpdEpi64(src M512i, k Mmask8, a M512d, rounding int) M512i {
	return M512i(maskCvtRoundpdEpi64([64]byte(src), uint8(k), [8]float64(a), rounding))
}

func maskCvtRoundpdEpi64(src [64]byte, k uint8, a [8]float64, rounding int) [64]byte


// MaskzCvtRoundpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_maskz_cvt_roundpd_epi64'.
// Requires AVX512DQ.
func MaskzCvtRoundpdEpi64(k Mmask8, a M512d, rounding int) M512i {
	return M512i(maskzCvtRoundpdEpi64(uint8(k), [8]float64(a), rounding))
}

func maskzCvtRoundpdEpi64(k uint8, a [8]float64, rounding int) [64]byte


// CvtRoundpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm512_cvt_roundpd_epu64'.
// Requires AVX512DQ.
func CvtRoundpdEpu64(a M512d, rounding int) M512i {
	return M512i(cvtRoundpdEpu64([8]float64(a), rounding))
}

func cvtRoundpdEpu64(a [8]float64, rounding int) [64]byte


// MaskCvtRoundpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm512_mask_cvt_roundpd_epu64'.
// Requires AVX512DQ.
func MaskCvtRoundpdEpu64(src M512i, k Mmask8, a M512d, rounding int) M512i {
	return M512i(maskCvtRoundpdEpu64([64]byte(src), uint8(k), [8]float64(a), rounding))
}

func maskCvtRoundpdEpu64(src [64]byte, k uint8, a [8]float64, rounding int) [64]byte


// MaskzCvtRoundpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers, and store
// the results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm512_maskz_cvt_roundpd_epu64'.
// Requires AVX512DQ.
func MaskzCvtRoundpdEpu64(k Mmask8, a M512d, rounding int) M512i {
	return M512i(maskzCvtRoundpdEpu64(uint8(k), [8]float64(a), rounding))
}

func maskzCvtRoundpdEpu64(k uint8, a [8]float64, rounding int) [64]byte


// CvtRoundpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm512_cvt_roundps_epi64'.
// Requires AVX512DQ.
func CvtRoundpsEpi64(a M256, rounding int) M512i {
	return M512i(cvtRoundpsEpi64([8]float32(a), rounding))
}

func cvtRoundpsEpi64(a [8]float32, rounding int) [64]byte


// MaskCvtRoundpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
// 	 Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm512_mask_cvt_roundps_epi64'.
// Requires AVX512DQ.
func MaskCvtRoundpsEpi64(src M512i, k Mmask8, a M256, rounding int) M512i {
	return M512i(maskCvtRoundpsEpi64([64]byte(src), uint8(k), [8]float32(a), rounding))
}

func maskCvtRoundpsEpi64(src [64]byte, k uint8, a [8]float32, rounding int) [64]byte


// MaskzCvtRoundpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm512_maskz_cvt_roundps_epi64'.
// Requires AVX512DQ.
func MaskzCvtRoundpsEpi64(k Mmask8, a M256, rounding int) M512i {
	return M512i(maskzCvtRoundpsEpi64(uint8(k), [8]float32(a), rounding))
}

func maskzCvtRoundpsEpi64(k uint8, a [8]float32, rounding int) [64]byte


// CvtRoundpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm512_cvt_roundps_epu64'.
// Requires AVX512DQ.
func CvtRoundpsEpu64(a M256, rounding int) M512i {
	return M512i(cvtRoundpsEpu64([8]float32(a), rounding))
}

func cvtRoundpsEpu64(a [8]float32, rounding int) [64]byte


// MaskCvtRoundpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm512_mask_cvt_roundps_epu64'.
// Requires AVX512DQ.
func MaskCvtRoundpsEpu64(src M512i, k Mmask8, a M256, rounding int) M512i {
	return M512i(maskCvtRoundpsEpu64([64]byte(src), uint8(k), [8]float32(a), rounding))
}

func maskCvtRoundpsEpu64(src [64]byte, k uint8, a [8]float32, rounding int) [64]byte


// MaskzCvtRoundpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers, and store
// the results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm512_maskz_cvt_roundps_epu64'.
// Requires AVX512DQ.
func MaskzCvtRoundpsEpu64(k Mmask8, a M256, rounding int) M512i {
	return M512i(maskzCvtRoundpsEpu64(uint8(k), [8]float32(a), rounding))
}

func maskzCvtRoundpsEpu64(k uint8, a [8]float32, rounding int) [64]byte


// Cvtepi64Pd: Convert packed 64-bit integers in 'a' to packed double-precision
// (64-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm_cvtepi64_pd'.
// Requires AVX512DQ.
func Cvtepi64Pd(a M128i) M128d {
	return M128d(cvtepi64Pd([16]byte(a)))
}

func cvtepi64Pd(a [16]byte) [2]float64


// MaskCvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm_mask_cvtepi64_pd'.
// Requires AVX512DQ.
func MaskCvtepi64Pd(src M128d, k Mmask8, a M128i) M128d {
	return M128d(maskCvtepi64Pd([2]float64(src), uint8(k), [16]byte(a)))
}

func maskCvtepi64Pd(src [2]float64, k uint8, a [16]byte) [2]float64


// MaskzCvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm_maskz_cvtepi64_pd'.
// Requires AVX512DQ.
func MaskzCvtepi64Pd(k Mmask8, a M128i) M128d {
	return M128d(maskzCvtepi64Pd(uint8(k), [16]byte(a)))
}

func maskzCvtepi64Pd(k uint8, a [16]byte) [2]float64


// Cvtepi64Pd1: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm256_cvtepi64_pd'.
// Requires AVX512DQ.
func Cvtepi64Pd1(a M256i) M256d {
	return M256d(cvtepi64Pd1([32]byte(a)))
}

func cvtepi64Pd1(a [32]byte) [4]float64


// MaskCvtepi64Pd1: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm256_mask_cvtepi64_pd'.
// Requires AVX512DQ.
func MaskCvtepi64Pd1(src M256d, k Mmask8, a M256i) M256d {
	return M256d(maskCvtepi64Pd1([4]float64(src), uint8(k), [32]byte(a)))
}

func maskCvtepi64Pd1(src [4]float64, k uint8, a [32]byte) [4]float64


// MaskzCvtepi64Pd1: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm256_maskz_cvtepi64_pd'.
// Requires AVX512DQ.
func MaskzCvtepi64Pd1(k Mmask8, a M256i) M256d {
	return M256d(maskzCvtepi64Pd1(uint8(k), [32]byte(a)))
}

func maskzCvtepi64Pd1(k uint8, a [32]byte) [4]float64


// Cvtepi64Pd2: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm512_cvtepi64_pd'.
// Requires AVX512DQ.
func Cvtepi64Pd2(a M512i) M512d {
	return M512d(cvtepi64Pd2([64]byte(a)))
}

func cvtepi64Pd2(a [64]byte) [8]float64


// MaskCvtepi64Pd2: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm512_mask_cvtepi64_pd'.
// Requires AVX512DQ.
func MaskCvtepi64Pd2(src M512d, k Mmask8, a M512i) M512d {
	return M512d(maskCvtepi64Pd2([8]float64(src), uint8(k), [64]byte(a)))
}

func maskCvtepi64Pd2(src [8]float64, k uint8, a [64]byte) [8]float64


// MaskzCvtepi64Pd2: Convert packed 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_Int64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTQQ2PD'. Intrinsic: '_mm512_maskz_cvtepi64_pd'.
// Requires AVX512DQ.
func MaskzCvtepi64Pd2(k Mmask8, a M512i) M512d {
	return M512d(maskzCvtepi64Pd2(uint8(k), [64]byte(a)))
}

func maskzCvtepi64Pd2(k uint8, a [64]byte) [8]float64


// Cvtepi64Ps: Convert packed 64-bit integers in 'a' to packed single-precision
// (32-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm_cvtepi64_ps'.
// Requires AVX512DQ.
func Cvtepi64Ps(a M128i) M128 {
	return M128(cvtepi64Ps([16]byte(a)))
}

func cvtepi64Ps(a [16]byte) [4]float32


// MaskCvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm_mask_cvtepi64_ps'.
// Requires AVX512DQ.
func MaskCvtepi64Ps(src M128, k Mmask8, a M128i) M128 {
	return M128(maskCvtepi64Ps([4]float32(src), uint8(k), [16]byte(a)))
}

func maskCvtepi64Ps(src [4]float32, k uint8, a [16]byte) [4]float32


// MaskzCvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm_maskz_cvtepi64_ps'.
// Requires AVX512DQ.
func MaskzCvtepi64Ps(k Mmask8, a M128i) M128 {
	return M128(maskzCvtepi64Ps(uint8(k), [16]byte(a)))
}

func maskzCvtepi64Ps(k uint8, a [16]byte) [4]float32


// Cvtepi64Ps1: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm256_cvtepi64_ps'.
// Requires AVX512DQ.
func Cvtepi64Ps1(a M256i) M128 {
	return M128(cvtepi64Ps1([32]byte(a)))
}

func cvtepi64Ps1(a [32]byte) [4]float32


// MaskCvtepi64Ps1: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm256_mask_cvtepi64_ps'.
// Requires AVX512DQ.
func MaskCvtepi64Ps1(src M128, k Mmask8, a M256i) M128 {
	return M128(maskCvtepi64Ps1([4]float32(src), uint8(k), [32]byte(a)))
}

func maskCvtepi64Ps1(src [4]float32, k uint8, a [32]byte) [4]float32


// MaskzCvtepi64Ps1: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm256_maskz_cvtepi64_ps'.
// Requires AVX512DQ.
func MaskzCvtepi64Ps1(k Mmask8, a M256i) M128 {
	return M128(maskzCvtepi64Ps1(uint8(k), [32]byte(a)))
}

func maskzCvtepi64Ps1(k uint8, a [32]byte) [4]float32


// Cvtepi64Ps2: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm512_cvtepi64_ps'.
// Requires AVX512DQ.
func Cvtepi64Ps2(a M512i) M256 {
	return M256(cvtepi64Ps2([64]byte(a)))
}

func cvtepi64Ps2(a [64]byte) [8]float32


// MaskCvtepi64Ps2: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm512_mask_cvtepi64_ps'.
// Requires AVX512DQ.
func MaskCvtepi64Ps2(src M256, k Mmask8, a M512i) M256 {
	return M256(maskCvtepi64Ps2([8]float32(src), uint8(k), [64]byte(a)))
}

func maskCvtepi64Ps2(src [8]float32, k uint8, a [64]byte) [8]float32


// MaskzCvtepi64Ps2: Convert packed 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Convert_Int64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTQQ2PS'. Intrinsic: '_mm512_maskz_cvtepi64_ps'.
// Requires AVX512DQ.
func MaskzCvtepi64Ps2(k Mmask8, a M512i) M256 {
	return M256(maskzCvtepi64Ps2(uint8(k), [64]byte(a)))
}

func maskzCvtepi64Ps2(k uint8, a [64]byte) [8]float32


// Cvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm_cvtepu64_pd'.
// Requires AVX512DQ.
func Cvtepu64Pd(a M128i) M128d {
	return M128d(cvtepu64Pd([16]byte(a)))
}

func cvtepu64Pd(a [16]byte) [2]float64


// MaskCvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm_mask_cvtepu64_pd'.
// Requires AVX512DQ.
func MaskCvtepu64Pd(src M128d, k Mmask8, a M128i) M128d {
	return M128d(maskCvtepu64Pd([2]float64(src), uint8(k), [16]byte(a)))
}

func maskCvtepu64Pd(src [2]float64, k uint8, a [16]byte) [2]float64


// MaskzCvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm_maskz_cvtepu64_pd'.
// Requires AVX512DQ.
func MaskzCvtepu64Pd(k Mmask8, a M128i) M128d {
	return M128d(maskzCvtepu64Pd(uint8(k), [16]byte(a)))
}

func maskzCvtepu64Pd(k uint8, a [16]byte) [2]float64


// Cvtepu64Pd1: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm256_cvtepu64_pd'.
// Requires AVX512DQ.
func Cvtepu64Pd1(a M256i) M256d {
	return M256d(cvtepu64Pd1([32]byte(a)))
}

func cvtepu64Pd1(a [32]byte) [4]float64


// MaskCvtepu64Pd1: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm256_mask_cvtepu64_pd'.
// Requires AVX512DQ.
func MaskCvtepu64Pd1(src M256d, k Mmask8, a M256i) M256d {
	return M256d(maskCvtepu64Pd1([4]float64(src), uint8(k), [32]byte(a)))
}

func maskCvtepu64Pd1(src [4]float64, k uint8, a [32]byte) [4]float64


// MaskzCvtepu64Pd1: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm256_maskz_cvtepu64_pd'.
// Requires AVX512DQ.
func MaskzCvtepu64Pd1(k Mmask8, a M256i) M256d {
	return M256d(maskzCvtepu64Pd1(uint8(k), [32]byte(a)))
}

func maskzCvtepu64Pd1(k uint8, a [32]byte) [4]float64


// Cvtepu64Pd2: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm512_cvtepu64_pd'.
// Requires AVX512DQ.
func Cvtepu64Pd2(a M512i) M512d {
	return M512d(cvtepu64Pd2([64]byte(a)))
}

func cvtepu64Pd2(a [64]byte) [8]float64


// MaskCvtepu64Pd2: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm512_mask_cvtepu64_pd'.
// Requires AVX512DQ.
func MaskCvtepu64Pd2(src M512d, k Mmask8, a M512i) M512d {
	return M512d(maskCvtepu64Pd2([8]float64(src), uint8(k), [64]byte(a)))
}

func maskCvtepu64Pd2(src [8]float64, k uint8, a [64]byte) [8]float64


// MaskzCvtepu64Pd2: Convert packed unsigned 64-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertUnsignedInt64_To_FP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUQQ2PD'. Intrinsic: '_mm512_maskz_cvtepu64_pd'.
// Requires AVX512DQ.
func MaskzCvtepu64Pd2(k Mmask8, a M512i) M512d {
	return M512d(maskzCvtepu64Pd2(uint8(k), [64]byte(a)))
}

func maskzCvtepu64Pd2(k uint8, a [64]byte) [8]float64


// Cvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm_cvtepu64_ps'.
// Requires AVX512DQ.
func Cvtepu64Ps(a M128i) M128 {
	return M128(cvtepu64Ps([16]byte(a)))
}

func cvtepu64Ps(a [16]byte) [4]float32


// MaskCvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm_mask_cvtepu64_ps'.
// Requires AVX512DQ.
func MaskCvtepu64Ps(src M128, k Mmask8, a M128i) M128 {
	return M128(maskCvtepu64Ps([4]float32(src), uint8(k), [16]byte(a)))
}

func maskCvtepu64Ps(src [4]float32, k uint8, a [16]byte) [4]float32


// MaskzCvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:64] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm_maskz_cvtepu64_ps'.
// Requires AVX512DQ.
func MaskzCvtepu64Ps(k Mmask8, a M128i) M128 {
	return M128(maskzCvtepu64Ps(uint8(k), [16]byte(a)))
}

func maskzCvtepu64Ps(k uint8, a [16]byte) [4]float32


// Cvtepu64Ps1: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm256_cvtepu64_ps'.
// Requires AVX512DQ.
func Cvtepu64Ps1(a M256i) M128 {
	return M128(cvtepu64Ps1([32]byte(a)))
}

func cvtepu64Ps1(a [32]byte) [4]float32


// MaskCvtepu64Ps1: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm256_mask_cvtepu64_ps'.
// Requires AVX512DQ.
func MaskCvtepu64Ps1(src M128, k Mmask8, a M256i) M128 {
	return M128(maskCvtepu64Ps1([4]float32(src), uint8(k), [32]byte(a)))
}

func maskCvtepu64Ps1(src [4]float32, k uint8, a [32]byte) [4]float32


// MaskzCvtepu64Ps1: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm256_maskz_cvtepu64_ps'.
// Requires AVX512DQ.
func MaskzCvtepu64Ps1(k Mmask8, a M256i) M128 {
	return M128(maskzCvtepu64Ps1(uint8(k), [32]byte(a)))
}

func maskzCvtepu64Ps1(k uint8, a [32]byte) [4]float32


// Cvtepu64Ps2: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm512_cvtepu64_ps'.
// Requires AVX512DQ.
func Cvtepu64Ps2(a M512i) M256 {
	return M256(cvtepu64Ps2([64]byte(a)))
}

func cvtepu64Ps2(a [64]byte) [8]float32


// MaskCvtepu64Ps2: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm512_mask_cvtepu64_ps'.
// Requires AVX512DQ.
func MaskCvtepu64Ps2(src M256, k Mmask8, a M512i) M256 {
	return M256(maskCvtepu64Ps2([8]float32(src), uint8(k), [64]byte(a)))
}

func maskCvtepu64Ps2(src [8]float32, k uint8, a [64]byte) [8]float32


// MaskzCvtepu64Ps2: Convert packed unsigned 64-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := ConvertUnsignedInt64_To_FP32(a[i+63:i])
//			ELSE
//				dst[l+31:l] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTUQQ2PS'. Intrinsic: '_mm512_maskz_cvtepu64_ps'.
// Requires AVX512DQ.
func MaskzCvtepu64Ps2(k Mmask8, a M512i) M256 {
	return M256(maskzCvtepu64Ps2(uint8(k), [64]byte(a)))
}

func maskzCvtepu64Ps2(k uint8, a [64]byte) [8]float32


// CvtpdEpi64: Convert packed double-precision (64-bit) floating-point elements
// in 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm_cvtpd_epi64'.
// Requires AVX512DQ.
func CvtpdEpi64(a M128d) M128i {
	return M128i(cvtpdEpi64([2]float64(a)))
}

func cvtpdEpi64(a [2]float64) [16]byte


// MaskCvtpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm_mask_cvtpd_epi64'.
// Requires AVX512DQ.
func MaskCvtpdEpi64(src M128i, k Mmask8, a M128d) M128i {
	return M128i(maskCvtpdEpi64([16]byte(src), uint8(k), [2]float64(a)))
}

func maskCvtpdEpi64(src [16]byte, k uint8, a [2]float64) [16]byte


// MaskzCvtpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm_maskz_cvtpd_epi64'.
// Requires AVX512DQ.
func MaskzCvtpdEpi64(k Mmask8, a M128d) M128i {
	return M128i(maskzCvtpdEpi64(uint8(k), [2]float64(a)))
}

func maskzCvtpdEpi64(k uint8, a [2]float64) [16]byte


// CvtpdEpi641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm256_cvtpd_epi64'.
// Requires AVX512DQ.
func CvtpdEpi641(a M256d) M256i {
	return M256i(cvtpdEpi641([4]float64(a)))
}

func cvtpdEpi641(a [4]float64) [32]byte


// MaskCvtpdEpi641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm256_mask_cvtpd_epi64'.
// Requires AVX512DQ.
func MaskCvtpdEpi641(src M256i, k Mmask8, a M256d) M256i {
	return M256i(maskCvtpdEpi641([32]byte(src), uint8(k), [4]float64(a)))
}

func maskCvtpdEpi641(src [32]byte, k uint8, a [4]float64) [32]byte


// MaskzCvtpdEpi641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm256_maskz_cvtpd_epi64'.
// Requires AVX512DQ.
func MaskzCvtpdEpi641(k Mmask8, a M256d) M256i {
	return M256i(maskzCvtpdEpi641(uint8(k), [4]float64(a)))
}

func maskzCvtpdEpi641(k uint8, a [4]float64) [32]byte


// CvtpdEpi642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_cvtpd_epi64'.
// Requires AVX512DQ.
func CvtpdEpi642(a M512d) M512i {
	return M512i(cvtpdEpi642([8]float64(a)))
}

func cvtpdEpi642(a [8]float64) [64]byte


// MaskCvtpdEpi642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_mask_cvtpd_epi64'.
// Requires AVX512DQ.
func MaskCvtpdEpi642(src M512i, k Mmask8, a M512d) M512i {
	return M512i(maskCvtpdEpi642([64]byte(src), uint8(k), [8]float64(a)))
}

func maskCvtpdEpi642(src [64]byte, k uint8, a [8]float64) [64]byte


// MaskzCvtpdEpi642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_maskz_cvtpd_epi64'.
// Requires AVX512DQ.
func MaskzCvtpdEpi642(k Mmask8, a M512d) M512i {
	return M512i(maskzCvtpdEpi642(uint8(k), [8]float64(a)))
}

func maskzCvtpdEpi642(k uint8, a [8]float64) [64]byte


// CvtpdEpu64: Convert packed double-precision (64-bit) floating-point elements
// in 'a' to packed unsigned 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm_cvtpd_epu64'.
// Requires AVX512DQ.
func CvtpdEpu64(a M128d) M128i {
	return M128i(cvtpdEpu64([2]float64(a)))
}

func cvtpdEpu64(a [2]float64) [16]byte


// MaskCvtpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm_mask_cvtpd_epu64'.
// Requires AVX512DQ.
func MaskCvtpdEpu64(src M128i, k Mmask8, a M128d) M128i {
	return M128i(maskCvtpdEpu64([16]byte(src), uint8(k), [2]float64(a)))
}

func maskCvtpdEpu64(src [16]byte, k uint8, a [2]float64) [16]byte


// MaskzCvtpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm_maskz_cvtpd_epu64'.
// Requires AVX512DQ.
func MaskzCvtpdEpu64(k Mmask8, a M128d) M128i {
	return M128i(maskzCvtpdEpu64(uint8(k), [2]float64(a)))
}

func maskzCvtpdEpu64(k uint8, a [2]float64) [16]byte


// CvtpdEpu641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm256_cvtpd_epu64'.
// Requires AVX512DQ.
func CvtpdEpu641(a M256d) M256i {
	return M256i(cvtpdEpu641([4]float64(a)))
}

func cvtpdEpu641(a [4]float64) [32]byte


// MaskCvtpdEpu641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm256_mask_cvtpd_epu64'.
// Requires AVX512DQ.
func MaskCvtpdEpu641(src M256i, k Mmask8, a M256d) M256i {
	return M256i(maskCvtpdEpu641([32]byte(src), uint8(k), [4]float64(a)))
}

func maskCvtpdEpu641(src [32]byte, k uint8, a [4]float64) [32]byte


// MaskzCvtpdEpu641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm256_maskz_cvtpd_epu64'.
// Requires AVX512DQ.
func MaskzCvtpdEpu641(k Mmask8, a M256d) M256i {
	return M256i(maskzCvtpdEpu641(uint8(k), [4]float64(a)))
}

func maskzCvtpdEpu641(k uint8, a [4]float64) [32]byte


// CvtpdEpu642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm512_cvtpd_epu64'.
// Requires AVX512DQ.
func CvtpdEpu642(a M512d) M512i {
	return M512i(cvtpdEpu642([8]float64(a)))
}

func cvtpdEpu642(a [8]float64) [64]byte


// MaskCvtpdEpu642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm512_mask_cvtpd_epu64'.
// Requires AVX512DQ.
func MaskCvtpdEpu642(src M512i, k Mmask8, a M512d) M512i {
	return M512i(maskCvtpdEpu642([64]byte(src), uint8(k), [8]float64(a)))
}

func maskCvtpdEpu642(src [64]byte, k uint8, a [8]float64) [64]byte


// MaskzCvtpdEpu642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2UQQ'. Intrinsic: '_mm512_maskz_cvtpd_epu64'.
// Requires AVX512DQ.
func MaskzCvtpdEpu642(k Mmask8, a M512d) M512i {
	return M512i(maskzCvtpdEpu642(uint8(k), [8]float64(a)))
}

func maskzCvtpdEpu642(k uint8, a [8]float64) [64]byte


// CvtpsEpi64: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm_cvtps_epi64'.
// Requires AVX512DQ.
func CvtpsEpi64(a M128) M128i {
	return M128i(cvtpsEpi64([4]float32(a)))
}

func cvtpsEpi64(a [4]float32) [16]byte


// MaskCvtpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm_mask_cvtps_epi64'.
// Requires AVX512DQ.
func MaskCvtpsEpi64(src M128i, k Mmask8, a M128) M128i {
	return M128i(maskCvtpsEpi64([16]byte(src), uint8(k), [4]float32(a)))
}

func maskCvtpsEpi64(src [16]byte, k uint8, a [4]float32) [16]byte


// MaskzCvtpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm_maskz_cvtps_epi64'.
// Requires AVX512DQ.
func MaskzCvtpsEpi64(k Mmask8, a M128) M128i {
	return M128i(maskzCvtpsEpi64(uint8(k), [4]float32(a)))
}

func maskzCvtpsEpi64(k uint8, a [4]float32) [16]byte


// CvtpsEpi641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm256_cvtps_epi64'.
// Requires AVX512DQ.
func CvtpsEpi641(a M128) M256i {
	return M256i(cvtpsEpi641([4]float32(a)))
}

func cvtpsEpi641(a [4]float32) [32]byte


// MaskCvtpsEpi641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm256_mask_cvtps_epi64'.
// Requires AVX512DQ.
func MaskCvtpsEpi641(src M256i, k Mmask8, a M128) M256i {
	return M256i(maskCvtpsEpi641([32]byte(src), uint8(k), [4]float32(a)))
}

func maskCvtpsEpi641(src [32]byte, k uint8, a [4]float32) [32]byte


// MaskzCvtpsEpi641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm256_maskz_cvtps_epi64'.
// Requires AVX512DQ.
func MaskzCvtpsEpi641(k Mmask8, a M128) M256i {
	return M256i(maskzCvtpsEpi641(uint8(k), [4]float32(a)))
}

func maskzCvtpsEpi641(k uint8, a [4]float32) [32]byte


// CvtpsEpi642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm512_cvtps_epi64'.
// Requires AVX512DQ.
func CvtpsEpi642(a M256) M512i {
	return M512i(cvtpsEpi642([8]float32(a)))
}

func cvtpsEpi642(a [8]float32) [64]byte


// MaskCvtpsEpi642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm512_mask_cvtps_epi64'.
// Requires AVX512DQ.
func MaskCvtpsEpi642(src M512i, k Mmask8, a M256) M512i {
	return M512i(maskCvtpsEpi642([64]byte(src), uint8(k), [8]float32(a)))
}

func maskCvtpsEpi642(src [64]byte, k uint8, a [8]float32) [64]byte


// MaskzCvtpsEpi642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers, and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2QQ'. Intrinsic: '_mm512_maskz_cvtps_epi64'.
// Requires AVX512DQ.
func MaskzCvtpsEpi642(k Mmask8, a M256) M512i {
	return M512i(maskzCvtpsEpi642(uint8(k), [8]float32(a)))
}

func maskzCvtpsEpi642(k uint8, a [8]float32) [64]byte


// CvtpsEpu64: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed unsigned 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm_cvtps_epu64'.
// Requires AVX512DQ.
func CvtpsEpu64(a M128) M128i {
	return M128i(cvtpsEpu64([4]float32(a)))
}

func cvtpsEpu64(a [4]float32) [16]byte


// MaskCvtpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm_mask_cvtps_epu64'.
// Requires AVX512DQ.
func MaskCvtpsEpu64(src M128i, k Mmask8, a M128) M128i {
	return M128i(maskCvtpsEpu64([16]byte(src), uint8(k), [4]float32(a)))
}

func maskCvtpsEpu64(src [16]byte, k uint8, a [4]float32) [16]byte


// MaskzCvtpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm_maskz_cvtps_epu64'.
// Requires AVX512DQ.
func MaskzCvtpsEpu64(k Mmask8, a M128) M128i {
	return M128i(maskzCvtpsEpu64(uint8(k), [4]float32(a)))
}

func maskzCvtpsEpu64(k uint8, a [4]float32) [16]byte


// CvtpsEpu641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm256_cvtps_epu64'.
// Requires AVX512DQ.
func CvtpsEpu641(a M128) M256i {
	return M256i(cvtpsEpu641([4]float32(a)))
}

func cvtpsEpu641(a [4]float32) [32]byte


// MaskCvtpsEpu641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm256_mask_cvtps_epu64'.
// Requires AVX512DQ.
func MaskCvtpsEpu641(src M256i, k Mmask8, a M128) M256i {
	return M256i(maskCvtpsEpu641([32]byte(src), uint8(k), [4]float32(a)))
}

func maskCvtpsEpu641(src [32]byte, k uint8, a [4]float32) [32]byte


// MaskzCvtpsEpu641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm256_maskz_cvtps_epu64'.
// Requires AVX512DQ.
func MaskzCvtpsEpu641(k Mmask8, a M128) M256i {
	return M256i(maskzCvtpsEpu641(uint8(k), [4]float32(a)))
}

func maskzCvtpsEpu641(k uint8, a [4]float32) [32]byte


// CvtpsEpu642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm512_cvtps_epu64'.
// Requires AVX512DQ.
func CvtpsEpu642(a M256) M512i {
	return M512i(cvtpsEpu642([8]float32(a)))
}

func cvtpsEpu642(a [8]float32) [64]byte


// MaskCvtpsEpu642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm512_mask_cvtps_epu64'.
// Requires AVX512DQ.
func MaskCvtpsEpu642(src M512i, k Mmask8, a M256) M512i {
	return M512i(maskCvtpsEpu642([64]byte(src), uint8(k), [8]float32(a)))
}

func maskCvtpsEpu642(src [64]byte, k uint8, a [8]float32) [64]byte


// MaskzCvtpsEpu642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers, and store the results in
// 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2UQQ'. Intrinsic: '_mm512_maskz_cvtps_epu64'.
// Requires AVX512DQ.
func MaskzCvtpsEpu642(k Mmask8, a M256) M512i {
	return M512i(maskzCvtpsEpu642(uint8(k), [8]float32(a)))
}

func maskzCvtpsEpu642(k uint8, a [8]float32) [64]byte


// CvttRoundpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae' to suppress all
// exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_cvtt_roundpd_epi64'.
// Requires AVX512DQ.
func CvttRoundpdEpi64(a M512d, sae int) M512i {
	return M512i(cvttRoundpdEpi64([8]float64(a), sae))
}

func cvttRoundpdEpi64(a [8]float64, sae int) [64]byte


// MaskCvttRoundpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_mask_cvtt_roundpd_epi64'.
// Requires AVX512DQ.
func MaskCvttRoundpdEpi64(src M512i, k Mmask8, a M512d, sae int) M512i {
	return M512i(maskCvttRoundpdEpi64([64]byte(src), uint8(k), [8]float64(a), sae))
}

func maskCvttRoundpdEpi64(src [64]byte, k uint8, a [8]float64, sae int) [64]byte


// MaskzCvttRoundpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). Pass __MM_FROUND_NO_EXC to
// 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_maskz_cvtt_roundpd_epi64'.
// Requires AVX512DQ.
func MaskzCvttRoundpdEpi64(k Mmask8, a M512d, sae int) M512i {
	return M512i(maskzCvttRoundpdEpi64(uint8(k), [8]float64(a), sae))
}

func maskzCvttRoundpdEpi64(k uint8, a [8]float64, sae int) [64]byte


// CvttRoundpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae' to suppress all
// exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_cvtt_roundpd_epu64'.
// Requires AVX512DQ.
func CvttRoundpdEpu64(a M512d, sae int) M512i {
	return M512i(cvttRoundpdEpu64([8]float64(a), sae))
}

func cvttRoundpdEpu64(a [8]float64, sae int) [64]byte


// MaskCvttRoundpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_mask_cvtt_roundpd_epu64'.
// Requires AVX512DQ.
func MaskCvttRoundpdEpu64(src M512i, k Mmask8, a M512d, sae int) M512i {
	return M512i(maskCvttRoundpdEpu64([64]byte(src), uint8(k), [8]float64(a), sae))
}

func maskCvttRoundpdEpu64(src [64]byte, k uint8, a [8]float64, sae int) [64]byte


// MaskzCvttRoundpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). Pass
// __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_maskz_cvtt_roundpd_epu64'.
// Requires AVX512DQ.
func MaskzCvttRoundpdEpu64(k Mmask8, a M512d, sae int) M512i {
	return M512i(maskzCvttRoundpdEpu64(uint8(k), [8]float64(a), sae))
}

func maskzCvttRoundpdEpu64(k uint8, a [8]float64, sae int) [64]byte


// CvttRoundpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae' to suppress all
// exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm512_cvtt_roundps_epi64'.
// Requires AVX512DQ.
func CvttRoundpsEpi64(a M256, sae int) M512i {
	return M512i(cvttRoundpsEpi64([8]float32(a), sae))
}

func cvttRoundpsEpi64(a [8]float32, sae int) [64]byte


// MaskCvttRoundpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm512_mask_cvtt_roundps_epi64'.
// Requires AVX512DQ.
func MaskCvttRoundpsEpi64(src M512i, k Mmask8, a M256, sae int) M512i {
	return M512i(maskCvttRoundpsEpi64([64]byte(src), uint8(k), [8]float32(a), sae))
}

func maskCvttRoundpsEpi64(src [64]byte, k uint8, a [8]float32, sae int) [64]byte


// MaskzCvttRoundpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). Pass __MM_FROUND_NO_EXC to
// 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm512_maskz_cvtt_roundps_epi64'.
// Requires AVX512DQ.
func MaskzCvttRoundpsEpi64(k Mmask8, a M256, sae int) M512i {
	return M512i(maskzCvttRoundpsEpi64(uint8(k), [8]float32(a), sae))
}

func maskzCvttRoundpsEpi64(k uint8, a [8]float32, sae int) [64]byte


// CvttRoundpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae' to suppress all
// exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm512_cvtt_roundps_epu64'.
// Requires AVX512DQ.
func CvttRoundpsEpu64(a M256, sae int) M512i {
	return M512i(cvttRoundpsEpu64([8]float32(a), sae))
}

func cvttRoundpsEpu64(a [8]float32, sae int) [64]byte


// MaskCvttRoundpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm512_mask_cvtt_roundps_epu64'.
// Requires AVX512DQ.
func MaskCvttRoundpsEpu64(src M512i, k Mmask8, a M256, sae int) M512i {
	return M512i(maskCvttRoundpsEpu64([64]byte(src), uint8(k), [8]float32(a), sae))
}

func maskCvttRoundpsEpu64(src [64]byte, k uint8, a [8]float32, sae int) [64]byte


// MaskzCvttRoundpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). Pass
// __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm512_maskz_cvtt_roundps_epu64'.
// Requires AVX512DQ.
func MaskzCvttRoundpsEpu64(k Mmask8, a M256, sae int) M512i {
	return M512i(maskzCvttRoundpsEpu64(uint8(k), [8]float32(a), sae))
}

func maskzCvttRoundpsEpu64(k uint8, a [8]float32, sae int) [64]byte


// CvttpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm_cvttpd_epi64'.
// Requires AVX512DQ.
func CvttpdEpi64(a M128d) M128i {
	return M128i(cvttpdEpi64([2]float64(a)))
}

func cvttpdEpi64(a [2]float64) [16]byte


// MaskCvttpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm_mask_cvttpd_epi64'.
// Requires AVX512DQ.
func MaskCvttpdEpi64(src M128i, k Mmask8, a M128d) M128i {
	return M128i(maskCvttpdEpi64([16]byte(src), uint8(k), [2]float64(a)))
}

func maskCvttpdEpi64(src [16]byte, k uint8, a [2]float64) [16]byte


// MaskzCvttpdEpi64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm_maskz_cvttpd_epi64'.
// Requires AVX512DQ.
func MaskzCvttpdEpi64(k Mmask8, a M128d) M128i {
	return M128i(maskzCvttpdEpi64(uint8(k), [2]float64(a)))
}

func maskzCvttpdEpi64(k uint8, a [2]float64) [16]byte


// CvttpdEpi641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm256_cvttpd_epi64'.
// Requires AVX512DQ.
func CvttpdEpi641(a M256d) M256i {
	return M256i(cvttpdEpi641([4]float64(a)))
}

func cvttpdEpi641(a [4]float64) [32]byte


// MaskCvttpdEpi641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm256_mask_cvttpd_epi64'.
// Requires AVX512DQ.
func MaskCvttpdEpi641(src M256i, k Mmask8, a M256d) M256i {
	return M256i(maskCvttpdEpi641([32]byte(src), uint8(k), [4]float64(a)))
}

func maskCvttpdEpi641(src [32]byte, k uint8, a [4]float64) [32]byte


// MaskzCvttpdEpi641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm256_maskz_cvttpd_epi64'.
// Requires AVX512DQ.
func MaskzCvttpdEpi641(k Mmask8, a M256d) M256i {
	return M256i(maskzCvttpdEpi641(uint8(k), [4]float64(a)))
}

func maskzCvttpdEpi641(k uint8, a [4]float64) [32]byte


// CvttpdEpi642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_cvttpd_epi64'.
// Requires AVX512DQ.
func CvttpdEpi642(a M512d) M512i {
	return M512i(cvttpdEpi642([8]float64(a)))
}

func cvttpdEpi642(a [8]float64) [64]byte


// MaskCvttpdEpi642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_mask_cvttpd_epi64'.
// Requires AVX512DQ.
func MaskCvttpdEpi642(src M512i, k Mmask8, a M512d) M512i {
	return M512i(maskCvttpdEpi642([64]byte(src), uint8(k), [8]float64(a)))
}

func maskCvttpdEpi642(src [64]byte, k uint8, a [8]float64) [64]byte


// MaskzCvttpdEpi642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_maskz_cvttpd_epi64'.
// Requires AVX512DQ.
func MaskzCvttpdEpi642(k Mmask8, a M512d) M512i {
	return M512i(maskzCvttpdEpi642(uint8(k), [8]float64(a)))
}

func maskzCvttpdEpi642(k uint8, a [8]float64) [64]byte


// CvttpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm_cvttpd_epu64'.
// Requires AVX512DQ.
func CvttpdEpu64(a M128d) M128i {
	return M128i(cvttpdEpu64([2]float64(a)))
}

func cvttpdEpu64(a [2]float64) [16]byte


// MaskCvttpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm_mask_cvttpd_epu64'.
// Requires AVX512DQ.
func MaskCvttpdEpu64(src M128i, k Mmask8, a M128d) M128i {
	return M128i(maskCvttpdEpu64([16]byte(src), uint8(k), [2]float64(a)))
}

func maskCvttpdEpu64(src [16]byte, k uint8, a [2]float64) [16]byte


// MaskzCvttpdEpu64: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm_maskz_cvttpd_epu64'.
// Requires AVX512DQ.
func MaskzCvttpdEpu64(k Mmask8, a M128d) M128i {
	return M128i(maskzCvttpdEpu64(uint8(k), [2]float64(a)))
}

func maskzCvttpdEpu64(k uint8, a [2]float64) [16]byte


// CvttpdEpu641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm256_cvttpd_epu64'.
// Requires AVX512DQ.
func CvttpdEpu641(a M256d) M256i {
	return M256i(cvttpdEpu641([4]float64(a)))
}

func cvttpdEpu641(a [4]float64) [32]byte


// MaskCvttpdEpu641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm256_mask_cvttpd_epu64'.
// Requires AVX512DQ.
func MaskCvttpdEpu641(src M256i, k Mmask8, a M256d) M256i {
	return M256i(maskCvttpdEpu641([32]byte(src), uint8(k), [4]float64(a)))
}

func maskCvttpdEpu641(src [32]byte, k uint8, a [4]float64) [32]byte


// MaskzCvttpdEpu641: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm256_maskz_cvttpd_epu64'.
// Requires AVX512DQ.
func MaskzCvttpdEpu641(k Mmask8, a M256d) M256i {
	return M256i(maskzCvttpdEpu641(uint8(k), [4]float64(a)))
}

func maskzCvttpdEpu641(k uint8, a [4]float64) [32]byte


// CvttpdEpu642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_cvttpd_epu64'.
// Requires AVX512DQ.
func CvttpdEpu642(a M512d) M512i {
	return M512i(cvttpdEpu642([8]float64(a)))
}

func cvttpdEpu642(a [8]float64) [64]byte


// MaskCvttpdEpu642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_mask_cvttpd_epu64'.
// Requires AVX512DQ.
func MaskCvttpdEpu642(src M512i, k Mmask8, a M512d) M512i {
	return M512i(maskCvttpdEpu642([64]byte(src), uint8(k), [8]float64(a)))
}

func maskCvttpdEpu642(src [64]byte, k uint8, a [8]float64) [64]byte


// MaskzCvttpdEpu642: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_maskz_cvttpd_epu64'.
// Requires AVX512DQ.
func MaskzCvttpdEpu642(k Mmask8, a M512d) M512i {
	return M512i(maskzCvttpdEpu642(uint8(k), [8]float64(a)))
}

func maskzCvttpdEpu642(k uint8, a [8]float64) [64]byte


// CvttpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm_cvttps_epi64'.
// Requires AVX512DQ.
func CvttpsEpi64(a M128) M128i {
	return M128i(cvttpsEpi64([4]float32(a)))
}

func cvttpsEpi64(a [4]float32) [16]byte


// MaskCvttpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm_mask_cvttps_epi64'.
// Requires AVX512DQ.
func MaskCvttpsEpi64(src M128i, k Mmask8, a M128) M128i {
	return M128i(maskCvttpsEpi64([16]byte(src), uint8(k), [4]float32(a)))
}

func maskCvttpsEpi64(src [16]byte, k uint8, a [4]float32) [16]byte


// MaskzCvttpsEpi64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm_maskz_cvttps_epi64'.
// Requires AVX512DQ.
func MaskzCvttpsEpi64(k Mmask8, a M128) M128i {
	return M128i(maskzCvttpsEpi64(uint8(k), [4]float32(a)))
}

func maskzCvttpsEpi64(k uint8, a [4]float32) [16]byte


// CvttpsEpi641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm256_cvttps_epi64'.
// Requires AVX512DQ.
func CvttpsEpi641(a M128) M256i {
	return M256i(cvttpsEpi641([4]float32(a)))
}

func cvttpsEpi641(a [4]float32) [32]byte


// MaskCvttpsEpi641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm256_mask_cvttps_epi64'.
// Requires AVX512DQ.
func MaskCvttpsEpi641(src M256i, k Mmask8, a M128) M256i {
	return M256i(maskCvttpsEpi641([32]byte(src), uint8(k), [4]float32(a)))
}

func maskCvttpsEpi641(src [32]byte, k uint8, a [4]float32) [32]byte


// MaskzCvttpsEpi641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm256_maskz_cvttps_epi64'.
// Requires AVX512DQ.
func MaskzCvttpsEpi641(k Mmask8, a M128) M256i {
	return M256i(maskzCvttpsEpi641(uint8(k), [4]float32(a)))
}

func maskzCvttpsEpi641(k uint8, a [4]float32) [32]byte


// CvttpsEpi642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm512_cvttps_epi64'.
// Requires AVX512DQ.
func CvttpsEpi642(a M256) M512i {
	return M512i(cvttpsEpi642([8]float32(a)))
}

func cvttpsEpi642(a [8]float32) [64]byte


// MaskCvttpsEpi642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm512_mask_cvttps_epi64'.
// Requires AVX512DQ.
func MaskCvttpsEpi642(src M512i, k Mmask8, a M256) M512i {
	return M512i(maskCvttpsEpi642([64]byte(src), uint8(k), [8]float32(a)))
}

func maskCvttpsEpi642(src [64]byte, k uint8, a [8]float32) [64]byte


// MaskzCvttpsEpi642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 64-bit integers with truncation, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_Int64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2QQ'. Intrinsic: '_mm512_maskz_cvttps_epi64'.
// Requires AVX512DQ.
func MaskzCvttpsEpi642(k Mmask8, a M256) M512i {
	return M512i(maskzCvttpsEpi642(uint8(k), [8]float32(a)))
}

func maskzCvttpsEpi642(k uint8, a [8]float32) [64]byte


// CvttpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm_cvttps_epu64'.
// Requires AVX512DQ.
func CvttpsEpu64(a M128) M128i {
	return M128i(cvttpsEpu64([4]float32(a)))
}

func cvttpsEpu64(a [4]float32) [16]byte


// MaskCvttpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm_mask_cvttps_epu64'.
// Requires AVX512DQ.
func MaskCvttpsEpu64(src M128i, k Mmask8, a M128) M128i {
	return M128i(maskCvttpsEpu64([16]byte(src), uint8(k), [4]float32(a)))
}

func maskCvttpsEpu64(src [16]byte, k uint8, a [4]float32) [16]byte


// MaskzCvttpsEpu64: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm_maskz_cvttps_epu64'.
// Requires AVX512DQ.
func MaskzCvttpsEpu64(k Mmask8, a M128) M128i {
	return M128i(maskzCvttpsEpu64(uint8(k), [4]float32(a)))
}

func maskzCvttpsEpu64(k uint8, a [4]float32) [16]byte


// CvttpsEpu641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm256_cvttps_epu64'.
// Requires AVX512DQ.
func CvttpsEpu641(a M128) M256i {
	return M256i(cvttpsEpu641([4]float32(a)))
}

func cvttpsEpu641(a [4]float32) [32]byte


// MaskCvttpsEpu641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm256_mask_cvttps_epu64'.
// Requires AVX512DQ.
func MaskCvttpsEpu641(src M256i, k Mmask8, a M128) M256i {
	return M256i(maskCvttpsEpu641([32]byte(src), uint8(k), [4]float32(a)))
}

func maskCvttpsEpu641(src [32]byte, k uint8, a [4]float32) [32]byte


// MaskzCvttpsEpu641: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm256_maskz_cvttps_epu64'.
// Requires AVX512DQ.
func MaskzCvttpsEpu641(k Mmask8, a M128) M256i {
	return M256i(maskzCvttpsEpu641(uint8(k), [4]float32(a)))
}

func maskzCvttpsEpu641(k uint8, a [4]float32) [32]byte


// CvttpsEpu642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm512_cvttps_epu64'.
// Requires AVX512DQ.
func CvttpsEpu642(a M256) M512i {
	return M512i(cvttpsEpu642([8]float32(a)))
}

func cvttpsEpu642(a [8]float32) [64]byte


// MaskCvttpsEpu642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm512_mask_cvttps_epu64'.
// Requires AVX512DQ.
func MaskCvttpsEpu642(src M512i, k Mmask8, a M256) M512i {
	return M512i(maskCvttpsEpu642([64]byte(src), uint8(k), [8]float32(a)))
}

func maskCvttpsEpu642(src [64]byte, k uint8, a [8]float32) [64]byte


// MaskzCvttpsEpu642: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed unsigned 64-bit integers with truncation, and
// store the results in 'dst' using zeromask 'k' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[i+63:i] := Convert_FP32_To_UnsignedInt64_Truncate(a[l+31:l])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPS2UQQ'. Intrinsic: '_mm512_maskz_cvttps_epu64'.
// Requires AVX512DQ.
func MaskzCvttpsEpu642(k Mmask8, a M256) M512i {
	return M512i(maskzCvttpsEpu642(uint8(k), [8]float32(a)))
}

func maskzCvttpsEpu642(k uint8, a [8]float32) [64]byte


// Extractf32x8Ps: Extract 256 bits (composed of 8 packed single-precision
// (32-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[255:0] := a[255:0]
//		1: dst[255:0] := a[511:256]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTF32X8'. Intrinsic: '_mm512_extractf32x8_ps'.
// Requires AVX512DQ.
func Extractf32x8Ps(a M512, imm8 int) M256 {
	return M256(extractf32x8Ps([16]float32(a), imm8))
}

func extractf32x8Ps(a [16]float32, imm8 int) [8]float32


// MaskExtractf32x8Ps: Extract 256 bits (composed of 8 packed single-precision
// (32-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[255:0] := a[255:0]
//		1: tmp[255:0] := a[511:256]
//		ESAC
//		
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTF32X8'. Intrinsic: '_mm512_mask_extractf32x8_ps'.
// Requires AVX512DQ.
func MaskExtractf32x8Ps(src M256, k Mmask8, a M512, imm8 int) M256 {
	return M256(maskExtractf32x8Ps([8]float32(src), uint8(k), [16]float32(a), imm8))
}

func maskExtractf32x8Ps(src [8]float32, k uint8, a [16]float32, imm8 int) [8]float32


// MaskzExtractf32x8Ps: Extract 256 bits (composed of 8 packed single-precision
// (32-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[255:0] := a[255:0]
//		1: tmp[255:0] := a[511:256]
//		ESAC
//		
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTF32X8'. Intrinsic: '_mm512_maskz_extractf32x8_ps'.
// Requires AVX512DQ.
func MaskzExtractf32x8Ps(k Mmask8, a M512, imm8 int) M256 {
	return M256(maskzExtractf32x8Ps(uint8(k), [16]float32(a), imm8))
}

func maskzExtractf32x8Ps(k uint8, a [16]float32, imm8 int) [8]float32


// Extractf64x2Pd: Extract 128 bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF64X2'. Intrinsic: '_mm256_extractf64x2_pd'.
// Requires AVX512DQ.
func Extractf64x2Pd(a M256d, imm8 int) M128d {
	return M128d(extractf64x2Pd([4]float64(a), imm8))
}

func extractf64x2Pd(a [4]float64, imm8 int) [2]float64


// MaskExtractf64x2Pd: Extract 128 bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF64X2'. Intrinsic: '_mm256_mask_extractf64x2_pd'.
// Requires AVX512DQ.
func MaskExtractf64x2Pd(src M128d, k Mmask8, a M256d, imm8 int) M128d {
	return M128d(maskExtractf64x2Pd([2]float64(src), uint8(k), [4]float64(a), imm8))
}

func maskExtractf64x2Pd(src [2]float64, k uint8, a [4]float64, imm8 int) [2]float64


// MaskzExtractf64x2Pd: Extract 128 bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF64X2'. Intrinsic: '_mm256_maskz_extractf64x2_pd'.
// Requires AVX512DQ.
func MaskzExtractf64x2Pd(k Mmask8, a M256d, imm8 int) M128d {
	return M128d(maskzExtractf64x2Pd(uint8(k), [4]float64(a), imm8))
}

func maskzExtractf64x2Pd(k uint8, a [4]float64, imm8 int) [2]float64


// Extractf64x2Pd1: Extract 128 bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		2: dst[127:0] := a[383:256]
//		3: dst[127:0] := a[511:384]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF64X2'. Intrinsic: '_mm512_extractf64x2_pd'.
// Requires AVX512DQ.
func Extractf64x2Pd1(a M512d, imm8 int) M128d {
	return M128d(extractf64x2Pd1([8]float64(a), imm8))
}

func extractf64x2Pd1(a [8]float64, imm8 int) [2]float64


// MaskExtractf64x2Pd1: Extract 128 bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		2: tmp[127:0] := a[383:256]
//		3: tmp[127:0] := a[511:384]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF64X2'. Intrinsic: '_mm512_mask_extractf64x2_pd'.
// Requires AVX512DQ.
func MaskExtractf64x2Pd1(src M128d, k Mmask8, a M512d, imm8 int) M128d {
	return M128d(maskExtractf64x2Pd1([2]float64(src), uint8(k), [8]float64(a), imm8))
}

func maskExtractf64x2Pd1(src [2]float64, k uint8, a [8]float64, imm8 int) [2]float64


// MaskzExtractf64x2Pd1: Extract 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'a', selected with
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		2: tmp[127:0] := a[383:256]
//		3: tmp[127:0] := a[511:384]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF64X2'. Intrinsic: '_mm512_maskz_extractf64x2_pd'.
// Requires AVX512DQ.
func MaskzExtractf64x2Pd1(k Mmask8, a M512d, imm8 int) M128d {
	return M128d(maskzExtractf64x2Pd1(uint8(k), [8]float64(a), imm8))
}

func maskzExtractf64x2Pd1(k uint8, a [8]float64, imm8 int) [2]float64


// Extracti32x8Epi32: Extract 256 bits (composed of 8 packed 32-bit integers)
// from 'a', selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[255:0] := a[255:0]
//		1: dst[255:0] := a[511:256]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTI32X8'. Intrinsic: '_mm512_extracti32x8_epi32'.
// Requires AVX512DQ.
func Extracti32x8Epi32(a M512i, imm8 int) M256i {
	return M256i(extracti32x8Epi32([64]byte(a), imm8))
}

func extracti32x8Epi32(a [64]byte, imm8 int) [32]byte


// MaskExtracti32x8Epi32: Extract 256 bits (composed of 8 packed 32-bit
// integers) from 'a', selected with 'imm8', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[255:0] := a[255:0]
//		1: tmp[255:0] := a[511:256]
//		ESAC
//		
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTI32X8'. Intrinsic: '_mm512_mask_extracti32x8_epi32'.
// Requires AVX512DQ.
func MaskExtracti32x8Epi32(src M256i, k Mmask8, a M512i, imm8 int) M256i {
	return M256i(maskExtracti32x8Epi32([32]byte(src), uint8(k), [64]byte(a), imm8))
}

func maskExtracti32x8Epi32(src [32]byte, k uint8, a [64]byte, imm8 int) [32]byte


// MaskzExtracti32x8Epi32: Extract 256 bits (composed of 8 packed 32-bit
// integers) from 'a', selected with 'imm8', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[255:0] := a[255:0]
//		1: tmp[255:0] := a[511:256]
//		ESAC
//		
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTI32X8'. Intrinsic: '_mm512_maskz_extracti32x8_epi32'.
// Requires AVX512DQ.
func MaskzExtracti32x8Epi32(k Mmask8, a M512i, imm8 int) M256i {
	return M256i(maskzExtracti32x8Epi32(uint8(k), [64]byte(a), imm8))
}

func maskzExtracti32x8Epi32(k uint8, a [64]byte, imm8 int) [32]byte


// Extracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit integers)
// from 'a', selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm256_extracti64x2_epi64'.
// Requires AVX512DQ.
func Extracti64x2Epi64(a M256i, imm8 int) M128i {
	return M128i(extracti64x2Epi64([32]byte(a), imm8))
}

func extracti64x2Epi64(a [32]byte, imm8 int) [16]byte


// MaskExtracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
// integers) from 'a', selected with 'imm8', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm256_mask_extracti64x2_epi64'.
// Requires AVX512DQ.
func MaskExtracti64x2Epi64(src M128i, k Mmask8, a M256i, imm8 int) M128i {
	return M128i(maskExtracti64x2Epi64([16]byte(src), uint8(k), [32]byte(a), imm8))
}

func maskExtracti64x2Epi64(src [16]byte, k uint8, a [32]byte, imm8 int) [16]byte


// MaskzExtracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
// integers) from 'a', selected with 'imm8', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm256_maskz_extracti64x2_epi64'.
// Requires AVX512DQ.
func MaskzExtracti64x2Epi64(k Mmask8, a M256i, imm8 int) M128i {
	return M128i(maskzExtracti64x2Epi64(uint8(k), [32]byte(a), imm8))
}

func maskzExtracti64x2Epi64(k uint8, a [32]byte, imm8 int) [16]byte


// Extracti64x2Epi641: Extract 128 bits (composed of 2 packed 64-bit integers)
// from 'a', selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		2: dst[127:0] := a[383:256]
//		3: dst[127:0] := a[511:384]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm512_extracti64x2_epi64'.
// Requires AVX512DQ.
func Extracti64x2Epi641(a M512i, imm8 int) M128i {
	return M128i(extracti64x2Epi641([64]byte(a), imm8))
}

func extracti64x2Epi641(a [64]byte, imm8 int) [16]byte


// MaskExtracti64x2Epi641: Extract 128 bits (composed of 2 packed 64-bit
// integers) from 'a', selected with 'imm8', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		2: tmp[127:0] := a[383:256]
//		3: tmp[127:0] := a[511:384]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm512_mask_extracti64x2_epi64'.
// Requires AVX512DQ.
func MaskExtracti64x2Epi641(src M128i, k Mmask8, a M512i, imm8 int) M128i {
	return M128i(maskExtracti64x2Epi641([16]byte(src), uint8(k), [64]byte(a), imm8))
}

func maskExtracti64x2Epi641(src [16]byte, k uint8, a [64]byte, imm8 int) [16]byte


// MaskzExtracti64x2Epi641: Extract 128 bits (composed of 2 packed 64-bit
// integers) from 'a', selected with 'imm8', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		CASE imm8[7:0] of
//		0: tmp[127:0] := a[127:0]
//		1: tmp[127:0] := a[255:128]
//		2: tmp[127:0] := a[383:256]
//		3: tmp[127:0] := a[511:384]
//		ESAC
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm512_maskz_extracti64x2_epi64'.
// Requires AVX512DQ.
func MaskzExtracti64x2Epi641(k Mmask8, a M512i, imm8 int) M128i {
	return M128i(maskzExtracti64x2Epi641(uint8(k), [64]byte(a), imm8))
}

func maskzExtracti64x2Epi641(k uint8, a [64]byte, imm8 int) [16]byte


// FpclassPdMask: Test packed double-precision (64-bit) floating-point elements
// in 'a' for special categories specified by 'imm8', and store the results in
// mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 1
//			i := j*64
//			k[j] := CheckFPClass_FP64(a[i+63:i], imm8[7:0])
//		ENDFOR
//		k[MAX:2] := 0
//
// Instruction: 'VFPCLASSPD'. Intrinsic: '_mm_fpclass_pd_mask'.
// Requires AVX512DQ.
func FpclassPdMask(a M128d, imm8 int) Mmask8 {
	return Mmask8(fpclassPdMask([2]float64(a), imm8))
}

func fpclassPdMask(a [2]float64, imm8 int) uint8


// MaskFpclassPdMask: Test packed double-precision (64-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k1[j]
//				k[j] := CheckFPClass_FP64(a[i+63:i], imm8[7:0])
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:2] := 0
//
// Instruction: 'VFPCLASSPD'. Intrinsic: '_mm_mask_fpclass_pd_mask'.
// Requires AVX512DQ.
func MaskFpclassPdMask(k1 Mmask8, a M128d, imm8 int) Mmask8 {
	return Mmask8(maskFpclassPdMask(uint8(k1), [2]float64(a), imm8))
}

func maskFpclassPdMask(k1 uint8, a [2]float64, imm8 int) uint8


// FpclassPdMask1: Test packed double-precision (64-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 3
//			i := j*64
//			k[j] := CheckFPClass_FP64(a[i+63:i], imm8[7:0])
//		ENDFOR
//		k[MAX:4] := 0
//
// Instruction: 'VFPCLASSPD'. Intrinsic: '_mm256_fpclass_pd_mask'.
// Requires AVX512DQ.
func FpclassPdMask1(a M256d, imm8 int) Mmask8 {
	return Mmask8(fpclassPdMask1([4]float64(a), imm8))
}

func fpclassPdMask1(a [4]float64, imm8 int) uint8


// MaskFpclassPdMask1: Test packed double-precision (64-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k1[j]
//				k[j] := CheckFPClass_FP64(a[i+63:i], imm8[7:0])
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:4] := 0
//
// Instruction: 'VFPCLASSPD'. Intrinsic: '_mm256_mask_fpclass_pd_mask'.
// Requires AVX512DQ.
func MaskFpclassPdMask1(k1 Mmask8, a M256d, imm8 int) Mmask8 {
	return Mmask8(maskFpclassPdMask1(uint8(k1), [4]float64(a), imm8))
}

func maskFpclassPdMask1(k1 uint8, a [4]float64, imm8 int) uint8


// FpclassPdMask2: Test packed double-precision (64-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := CheckFPClass_FP64(a[i+63:i], imm8[7:0])
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VFPCLASSPD'. Intrinsic: '_mm512_fpclass_pd_mask'.
// Requires AVX512DQ.
func FpclassPdMask2(a M512d, imm8 int) Mmask8 {
	return Mmask8(fpclassPdMask2([8]float64(a), imm8))
}

func fpclassPdMask2(a [8]float64, imm8 int) uint8


// MaskFpclassPdMask2: Test packed double-precision (64-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := CheckFPClass_FP64(a[i+63:i], imm8[7:0])
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VFPCLASSPD'. Intrinsic: '_mm512_mask_fpclass_pd_mask'.
// Requires AVX512DQ.
func MaskFpclassPdMask2(k1 Mmask8, a M512d, imm8 int) Mmask8 {
	return Mmask8(maskFpclassPdMask2(uint8(k1), [8]float64(a), imm8))
}

func maskFpclassPdMask2(k1 uint8, a [8]float64, imm8 int) uint8


// FpclassPsMask: Test packed single-precision (32-bit) floating-point elements
// in 'a' for special categories specified by 'imm8', and store the results in
// mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 3
//			i := j*32
//			k[j] := CheckFPClass_FP32(a[i+31:i], imm8[7:0])
//		ENDFOR
//		k[MAX:4] := 0
//
// Instruction: 'VFPCLASSPS'. Intrinsic: '_mm_fpclass_ps_mask'.
// Requires AVX512DQ.
func FpclassPsMask(a M128, imm8 int) Mmask8 {
	return Mmask8(fpclassPsMask([4]float32(a), imm8))
}

func fpclassPsMask(a [4]float32, imm8 int) uint8


// MaskFpclassPsMask: Test packed single-precision (32-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k1[j]
//				k[j] := CheckFPClass_FP32(a[i+31:i], imm8[7:0])
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:4] := 0
//
// Instruction: 'VFPCLASSPS'. Intrinsic: '_mm_mask_fpclass_ps_mask'.
// Requires AVX512DQ.
func MaskFpclassPsMask(k1 Mmask8, a M128, imm8 int) Mmask8 {
	return Mmask8(maskFpclassPsMask(uint8(k1), [4]float32(a), imm8))
}

func maskFpclassPsMask(k1 uint8, a [4]float32, imm8 int) uint8


// FpclassPsMask1: Test packed single-precision (32-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 7
//			i := j*32
//			k[j] := CheckFPClass_FP32(a[i+31:i], imm8[7:0])
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VFPCLASSPS'. Intrinsic: '_mm256_fpclass_ps_mask'.
// Requires AVX512DQ.
func FpclassPsMask1(a M256, imm8 int) Mmask8 {
	return Mmask8(fpclassPsMask1([8]float32(a), imm8))
}

func fpclassPsMask1(a [8]float32, imm8 int) uint8


// MaskFpclassPsMask1: Test packed single-precision (32-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k1[j]
//				k[j] := CheckFPClass_FP32(a[i+31:i], imm8[7:0])
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VFPCLASSPS'. Intrinsic: '_mm256_mask_fpclass_ps_mask'.
// Requires AVX512DQ.
func MaskFpclassPsMask1(k1 Mmask8, a M256, imm8 int) Mmask8 {
	return Mmask8(maskFpclassPsMask1(uint8(k1), [8]float32(a), imm8))
}

func maskFpclassPsMask1(k1 uint8, a [8]float32, imm8 int) uint8


// FpclassPsMask2: Test packed single-precision (32-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := CheckFPClass_FP32(a[i+31:i], imm8[7:0])
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VFPCLASSPS'. Intrinsic: '_mm512_fpclass_ps_mask'.
// Requires AVX512DQ.
func FpclassPsMask2(a M512, imm8 int) Mmask16 {
	return Mmask16(fpclassPsMask2([16]float32(a), imm8))
}

func fpclassPsMask2(a [16]float32, imm8 int) uint16


// MaskFpclassPsMask2: Test packed single-precision (32-bit) floating-point
// elements in 'a' for special categories specified by 'imm8', and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := CheckFPClass_FP32(a[i+31:i], imm8[7:0])
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VFPCLASSPS'. Intrinsic: '_mm512_mask_fpclass_ps_mask'.
// Requires AVX512DQ.
func MaskFpclassPsMask2(k1 Mmask16, a M512, imm8 int) Mmask16 {
	return Mmask16(maskFpclassPsMask2(uint16(k1), [16]float32(a), imm8))
}

func maskFpclassPsMask2(k1 uint16, a [16]float32, imm8 int) uint16


// FpclassSdMask: Test the lower double-precision (64-bit) floating-point
// element in 'a' for special categories specified by 'imm8', and store the
// result in mask vector 'k'.
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		k[0] := CheckFPClass_FP64(a[63:0], imm8[7:0])
//		k[MAX:1] := 0
//
// Instruction: 'VFPCLASSSD'. Intrinsic: '_mm_fpclass_sd_mask'.
// Requires AVX512DQ.
func FpclassSdMask(a M128d, imm8 int) Mmask8 {
	return Mmask8(fpclassSdMask([2]float64(a), imm8))
}

func fpclassSdMask(a [2]float64, imm8 int) uint8


// MaskFpclassSdMask: Test the lower double-precision (64-bit) floating-point
// element in 'a' for special categories specified by 'imm8', and store the
// result in mask vector 'k' using zeromask 'k1' (the element is zeroed out
// when mask bit 0 is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		IF k1[0]
//			k[0] := CheckFPClass_FP64(a[63:0], imm8[7:0])
//		ELSE
//			k[0] := 0
//		FI
//		k[MAX:1] := 0
//
// Instruction: 'VFPCLASSSD'. Intrinsic: '_mm_mask_fpclass_sd_mask'.
// Requires AVX512DQ.
func MaskFpclassSdMask(k1 Mmask8, a M128d, imm8 int) Mmask8 {
	return Mmask8(maskFpclassSdMask(uint8(k1), [2]float64(a), imm8))
}

func maskFpclassSdMask(k1 uint8, a [2]float64, imm8 int) uint8


// FpclassSsMask: Test the lower single-precision (32-bit) floating-point
// element in 'a' for special categories specified by 'imm8', and store the
// result in mask vector 'k.
// 	'imm" can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		k[0] := CheckFPClass_FP32(a[31:0], imm8[7:0])
//		k[MAX:1] := 0
//
// Instruction: 'VFPCLASSSS'. Intrinsic: '_mm_fpclass_ss_mask'.
// Requires AVX512DQ.
func FpclassSsMask(a M128, imm8 int) Mmask8 {
	return Mmask8(fpclassSsMask([4]float32(a), imm8))
}

func fpclassSsMask(a [4]float32, imm8 int) uint8


// MaskFpclassSsMask: Test the lower single-precision (32-bit) floating-point
// element in 'a' for special categories specified by 'imm8', and store the
// result in mask vector 'k' using zeromask 'k1' (the element is zeroed out
// when mask bit 0 is not set).
// 	'imm' can be a combination of:
//     0x01 // QNaN
//     0x02 // Positive Zero
//     0x04 // Negative Zero
//     0x08 // Positive Infinity
//     0x10 // Negative Infinity
//     0x20 // Denormal
//     0x40 // Negative
//     0x80 // SNaN 
//
//		IF k1[0]
//			k[0] := CheckFPClass_FP32(a[31:0], imm8[7:0])
//		ELSE
//			k[0] := 0
//		FI
//		k[MAX:1] := 0
//
// Instruction: 'VFPCLASSSS'. Intrinsic: '_mm_mask_fpclass_ss_mask'.
// Requires AVX512DQ.
func MaskFpclassSsMask(k1 Mmask8, a M128, imm8 int) Mmask8 {
	return Mmask8(maskFpclassSsMask(uint8(k1), [4]float32(a), imm8))
}

func maskFpclassSsMask(k1 uint8, a [4]float32, imm8 int) uint8


// Insertf32x8: Copy 'a' to 'dst', then insert 256 bits (composed of 8 packed
// single-precision (32-bit) floating-point elements) from 'b' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[511:0] := a[511:0]
//		CASE (imm8[7:0]) OF
//		0: dst[255:0] := b[255:0]
//		1: dst[511:256] := b[255:0]
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTF32X8'. Intrinsic: '_mm512_insertf32x8'.
// Requires AVX512DQ.
func Insertf32x8(a M512, b M256, imm8 int) M512 {
	return M512(insertf32x8([16]float32(a), [8]float32(b), imm8))
}

func insertf32x8(a [16]float32, b [8]float32, imm8 int) [16]float32


// MaskInsertf32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
// packed single-precision (32-bit) floating-point elements) from 'b' into
// 'tmp' at the location specified by 'imm8'.  Store 'tmp' to 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[7:0]) OF
//		0: tmp[255:0] := b[255:0]
//		1: tmp[511:256] := b[255:0]
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTF32X8'. Intrinsic: '_mm512_mask_insertf32x8'.
// Requires AVX512DQ.
func MaskInsertf32x8(src M512, k Mmask16, a M512, b M256, imm8 int) M512 {
	return M512(maskInsertf32x8([16]float32(src), uint16(k), [16]float32(a), [8]float32(b), imm8))
}

func maskInsertf32x8(src [16]float32, k uint16, a [16]float32, b [8]float32, imm8 int) [16]float32


// MaskzInsertf32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
// packed single-precision (32-bit) floating-point elements) from 'b' into
// 'tmp' at the location specified by 'imm8'.  Store 'tmp' to 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[7:0]) OF
//		0: tmp[255:0] := b[255:0]
//		1: tmp[511:256] := b[255:0]
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTF32X8'. Intrinsic: '_mm512_maskz_insertf32x8'.
// Requires AVX512DQ.
func MaskzInsertf32x8(k Mmask16, a M512, b M256, imm8 int) M512 {
	return M512(maskzInsertf32x8(uint16(k), [16]float32(a), [8]float32(b), imm8))
}

func maskzInsertf32x8(k uint16, a [16]float32, b [8]float32, imm8 int) [16]float32


// Insertf64x2: Copy 'a' to 'dst', then insert 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'b' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[255:0] := a[255:0]
//		CASE imm8[7:0] of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF64X2'. Intrinsic: '_mm256_insertf64x2'.
// Requires AVX512DQ.
func Insertf64x2(a M256d, b M128d, imm8 int) M256d {
	return M256d(insertf64x2([4]float64(a), [2]float64(b), imm8))
}

func insertf64x2(a [4]float64, b [2]float64, imm8 int) [4]float64


// MaskInsertf64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'tmp' at the location specified by 'imm8'.  Store 'tmp' to 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		tmp[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		ESAC
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF64X2'. Intrinsic: '_mm256_mask_insertf64x2'.
// Requires AVX512DQ.
func MaskInsertf64x2(src M256d, k Mmask8, a M256d, b M128d, imm8 int) M256d {
	return M256d(maskInsertf64x2([4]float64(src), uint8(k), [4]float64(a), [2]float64(b), imm8))
}

func maskInsertf64x2(src [4]float64, k uint8, a [4]float64, b [2]float64, imm8 int) [4]float64


// MaskzInsertf64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'tmp' at the location specified by 'imm8'.  Store 'tmp' to 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		tmp[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		ESAC
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF64X2'. Intrinsic: '_mm256_maskz_insertf64x2'.
// Requires AVX512DQ.
func MaskzInsertf64x2(k Mmask8, a M256d, b M128d, imm8 int) M256d {
	return M256d(maskzInsertf64x2(uint8(k), [4]float64(a), [2]float64(b), imm8))
}

func maskzInsertf64x2(k uint8, a [4]float64, b [2]float64, imm8 int) [4]float64


// Insertf64x21: Copy 'a' to 'dst', then insert 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'b' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[511:0] := a[511:0]
//		CASE imm8[7:0] of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		2: dst[383:256] := b[127:0]
//		3: dst[511:384] := b[127:0]
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTF64X2'. Intrinsic: '_mm512_insertf64x2'.
// Requires AVX512DQ.
func Insertf64x21(a M512d, b M128d, imm8 int) M512d {
	return M512d(insertf64x21([8]float64(a), [2]float64(b), imm8))
}

func insertf64x21(a [8]float64, b [2]float64, imm8 int) [8]float64


// MaskInsertf64x21: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'tmp' at the location specified by 'imm8'.  Store 'tmp' to 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		2: tmp[383:256] := b[127:0]
//		3: tmp[511:384] := b[127:0]
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTF64X2'. Intrinsic: '_mm512_mask_insertf64x2'.
// Requires AVX512DQ.
func MaskInsertf64x21(src M512d, k Mmask8, a M512d, b M128d, imm8 int) M512d {
	return M512d(maskInsertf64x21([8]float64(src), uint8(k), [8]float64(a), [2]float64(b), imm8))
}

func maskInsertf64x21(src [8]float64, k uint8, a [8]float64, b [2]float64, imm8 int) [8]float64


// MaskzInsertf64x21: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'tmp' at the location specified by 'imm8'.  Store 'tmp' to 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		2: tmp[383:256] := b[127:0]
//		3: tmp[511:384] := b[127:0]
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTF64X2'. Intrinsic: '_mm512_maskz_insertf64x2'.
// Requires AVX512DQ.
func MaskzInsertf64x21(k Mmask8, a M512d, b M128d, imm8 int) M512d {
	return M512d(maskzInsertf64x21(uint8(k), [8]float64(a), [2]float64(b), imm8))
}

func maskzInsertf64x21(k uint8, a [8]float64, b [2]float64, imm8 int) [8]float64


// Inserti32x8: Copy 'a' to 'dst', then insert 256 bits (composed of 8 packed
// 32-bit integers) from 'b' into 'dst' at the location specified by 'imm8'. 
//
//		dst[511:0] := a[511:0]
//		CASE imm8[7:0] of
//		0: dst[255:0] := b[255:0]
//		1: dst[511:256] := b[255:0]
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTI32X8'. Intrinsic: '_mm512_inserti32x8'.
// Requires AVX512DQ.
func Inserti32x8(a M512i, b M256i, imm8 int) M512i {
	return M512i(inserti32x8([64]byte(a), [32]byte(b), imm8))
}

func inserti32x8(a [64]byte, b [32]byte, imm8 int) [64]byte


// MaskInserti32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
// packed 32-bit integers) from 'b' into 'tmp' at the location specified by
// 'imm8'.  Store 'tmp' to 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[7:0]) OF
//		0: tmp[255:0] := b[255:0]
//		1: tmp[511:256] := b[255:0]
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTI32X8'. Intrinsic: '_mm512_mask_inserti32x8'.
// Requires AVX512DQ.
func MaskInserti32x8(src M512i, k Mmask16, a M512i, b M256i, imm8 int) M512i {
	return M512i(maskInserti32x8([64]byte(src), uint16(k), [64]byte(a), [32]byte(b), imm8))
}

func maskInserti32x8(src [64]byte, k uint16, a [64]byte, b [32]byte, imm8 int) [64]byte


// MaskzInserti32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
// packed 32-bit integers) from 'b' into 'tmp' at the location specified by
// 'imm8'.  Store 'tmp' to 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[7:0]) OF
//		0: tmp[255:0] := b[255:0]
//		1: tmp[511:256] := b[255:0]
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTI32X8'. Intrinsic: '_mm512_maskz_inserti32x8'.
// Requires AVX512DQ.
func MaskzInserti32x8(k Mmask16, a M512i, b M256i, imm8 int) M512i {
	return M512i(maskzInserti32x8(uint16(k), [64]byte(a), [32]byte(b), imm8))
}

func maskzInserti32x8(k uint16, a [64]byte, b [32]byte, imm8 int) [64]byte


// Inserti64x2: Copy 'a' to 'dst', then insert 128 bits (composed of 2 packed
// 64-bit integers) from 'b' into 'dst' at the location specified by 'imm8'. 
//
//		dst[255:0] := a[255:0]
//		CASE imm8[7:0] of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTI64X2'. Intrinsic: '_mm256_inserti64x2'.
// Requires AVX512DQ.
func Inserti64x2(a M256i, b M128i, imm8 int) M256i {
	return M256i(inserti64x2([32]byte(a), [16]byte(b), imm8))
}

func inserti64x2(a [32]byte, b [16]byte, imm8 int) [32]byte


// MaskInserti64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed 64-bit integers) from 'b' into 'tmp' at the location specified by
// 'imm8'.  Store 'tmp' to 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		tmp[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		ESAC
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTI64X2'. Intrinsic: '_mm256_mask_inserti64x2'.
// Requires AVX512DQ.
func MaskInserti64x2(src M256i, k Mmask8, a M256i, b M128i, imm8 int) M256i {
	return M256i(maskInserti64x2([32]byte(src), uint8(k), [32]byte(a), [16]byte(b), imm8))
}

func maskInserti64x2(src [32]byte, k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte


// MaskzInserti64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed 64-bit integers) from 'b' into 'tmp' at the location specified by
// 'imm8'.  Store 'tmp' to 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		tmp[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		ESAC
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTI64X2'. Intrinsic: '_mm256_maskz_inserti64x2'.
// Requires AVX512DQ.
func MaskzInserti64x2(k Mmask8, a M256i, b M128i, imm8 int) M256i {
	return M256i(maskzInserti64x2(uint8(k), [32]byte(a), [16]byte(b), imm8))
}

func maskzInserti64x2(k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte


// Inserti64x21: Copy 'a' to 'dst', then insert 128 bits (composed of 2 packed
// 64-bit integers) from 'b' into 'dst' at the location specified by 'imm8'. 
//
//		dst[511:0] := a[511:0]
//		CASE imm8[7:0] of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		2: dst[383:256] := b[127:0]
//		3: dst[511:384] := b[127:0]
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTI64X2'. Intrinsic: '_mm512_inserti64x2'.
// Requires AVX512DQ.
func Inserti64x21(a M512i, b M128i, imm8 int) M512i {
	return M512i(inserti64x21([64]byte(a), [16]byte(b), imm8))
}

func inserti64x21(a [64]byte, b [16]byte, imm8 int) [64]byte


// MaskInserti64x21: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed 64-bit integers) from 'b' into 'tmp' at the location specified by
// 'imm8'.  Store 'tmp' to 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		2: tmp[383:256] := b[127:0]
//		3: tmp[511:384] := b[127:0]
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTI64X2'. Intrinsic: '_mm512_mask_inserti64x2'.
// Requires AVX512DQ.
func MaskInserti64x21(src M512i, k Mmask8, a M512i, b M128i, imm8 int) M512i {
	return M512i(maskInserti64x21([64]byte(src), uint8(k), [64]byte(a), [16]byte(b), imm8))
}

func maskInserti64x21(src [64]byte, k uint8, a [64]byte, b [16]byte, imm8 int) [64]byte


// MaskzInserti64x21: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
// packed 64-bit integers) from 'b' into 'tmp' at the location specified by
// 'imm8'.  Store 'tmp' to 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		tmp[511:0] := a[511:0]
//		CASE (imm8[1:0]) of
//		0: tmp[127:0] := b[127:0]
//		1: tmp[255:128] := b[127:0]
//		2: tmp[383:256] := b[127:0]
//		3: tmp[511:384] := b[127:0]
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := tmp[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VINSERTI64X2'. Intrinsic: '_mm512_maskz_inserti64x2'.
// Requires AVX512DQ.
func MaskzInserti64x21(k Mmask8, a M512i, b M128i, imm8 int) M512i {
	return M512i(maskzInserti64x21(uint8(k), [64]byte(a), [16]byte(b), imm8))
}

func maskzInserti64x21(k uint8, a [64]byte, b [16]byte, imm8 int) [64]byte


// Movepi32Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 32-bit integer in 'a'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF a[i+31]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:4] := 0
//
// Instruction: 'VPMOVD2M'. Intrinsic: '_mm_movepi32_mask'.
// Requires AVX512DQ.
func Movepi32Mask(a M128i) Mmask8 {
	return Mmask8(movepi32Mask([16]byte(a)))
}

func movepi32Mask(a [16]byte) uint8


// Movepi32Mask1: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 32-bit integer in 'a'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF a[i+31]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPMOVD2M'. Intrinsic: '_mm256_movepi32_mask'.
// Requires AVX512DQ.
func Movepi32Mask1(a M256i) Mmask8 {
	return Mmask8(movepi32Mask1([32]byte(a)))
}

func movepi32Mask1(a [32]byte) uint8


// Movepi32Mask2: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 32-bit integer in 'a'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPMOVD2M'. Intrinsic: '_mm512_movepi32_mask'.
// Requires AVX512DQ.
func Movepi32Mask2(a M512i) Mmask16 {
	return Mmask16(movepi32Mask2([64]byte(a)))
}

func movepi32Mask2(a [64]byte) uint16


// Movepi64Mask: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 64-bit integer in 'a'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF a[i+63]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:2] := 0
//
// Instruction: 'VPMOVQ2M'. Intrinsic: '_mm_movepi64_mask'.
// Requires AVX512DQ.
func Movepi64Mask(a M128i) Mmask8 {
	return Mmask8(movepi64Mask([16]byte(a)))
}

func movepi64Mask(a [16]byte) uint8


// Movepi64Mask1: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 64-bit integer in 'a'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF a[i+63]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:4] := 0
//
// Instruction: 'VPMOVQ2M'. Intrinsic: '_mm256_movepi64_mask'.
// Requires AVX512DQ.
func Movepi64Mask1(a M256i) Mmask8 {
	return Mmask8(movepi64Mask1([32]byte(a)))
}

func movepi64Mask1(a [32]byte) uint8


// Movepi64Mask2: Set each bit of mask register 'k' based on the most
// significant bit of the corresponding packed 64-bit integer in 'a'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF a[i+63]
//				k[j] := 1
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VPMOVQ2M'. Intrinsic: '_mm512_movepi64_mask'.
// Requires AVX512DQ.
func Movepi64Mask2(a M512i) Mmask8 {
	return Mmask8(movepi64Mask2([64]byte(a)))
}

func movepi64Mask2(a [64]byte) uint8


// MovmEpi32: Set each packed 32-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := 0xFFFFFFFF
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVM2D'. Intrinsic: '_mm_movm_epi32'.
// Requires AVX512DQ.
func MovmEpi32(k Mmask8) M128i {
	return M128i(movmEpi32(uint8(k)))
}

func movmEpi32(k uint8) [16]byte


// MovmEpi321: Set each packed 32-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := 0xFFFFFFFF
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVM2D'. Intrinsic: '_mm256_movm_epi32'.
// Requires AVX512DQ.
func MovmEpi321(k Mmask8) M256i {
	return M256i(movmEpi321(uint8(k)))
}

func movmEpi321(k uint8) [32]byte


// MovmEpi322: Set each packed 32-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := 0xFFFFFFFF
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVM2D'. Intrinsic: '_mm512_movm_epi32'.
// Requires AVX512DQ.
func MovmEpi322(k Mmask16) M512i {
	return M512i(movmEpi322(uint16(k)))
}

func movmEpi322(k uint16) [64]byte


// MovmEpi64: Set each packed 64-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := 0xFFFFFFFFffffffff
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMOVM2Q'. Intrinsic: '_mm_movm_epi64'.
// Requires AVX512DQ.
func MovmEpi64(k Mmask8) M128i {
	return M128i(movmEpi64(uint8(k)))
}

func movmEpi64(k uint8) [16]byte


// MovmEpi641: Set each packed 64-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := 0xFFFFFFFFffffffff
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMOVM2Q'. Intrinsic: '_mm256_movm_epi64'.
// Requires AVX512DQ.
func MovmEpi641(k Mmask8) M256i {
	return M256i(movmEpi641(uint8(k)))
}

func movmEpi641(k uint8) [32]byte


// MovmEpi642: Set each packed 64-bit integer in 'dst' to all ones or all zeros
// based on the value of the corresponding bit in 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := 0xFFFFFFFFffffffff
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMOVM2Q'. Intrinsic: '_mm512_movm_epi64'.
// Requires AVX512DQ.
func MovmEpi642(k Mmask8) M512i {
	return M512i(movmEpi642(uint8(k)))
}

func movmEpi642(k uint8) [64]byte


// MaskMulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp[127:0] := a[i+63:i] * b[i+63:i]
//				dst[i+63:i] := tmp[63:0]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm_mask_mullo_epi64'.
// Requires AVX512DQ.
func MaskMulloEpi64(src M128i, k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskMulloEpi64([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
}

func maskMulloEpi64(src [16]byte, k uint8, a [16]byte, b [16]byte) [16]byte


// MaskzMulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp[127:0] := a[i+63:i] * b[i+63:i]
//				dst[i+63:i] := tmp[63:0]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm_maskz_mullo_epi64'.
// Requires AVX512DQ.
func MaskzMulloEpi64(k Mmask8, a M128i, b M128i) M128i {
	return M128i(maskzMulloEpi64(uint8(k), [16]byte(a), [16]byte(b)))
}

func maskzMulloEpi64(k uint8, a [16]byte, b [16]byte) [16]byte


// MulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b', producing
// intermediate 128-bit integers, and store the low 64 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			tmp[127:0] := a[i+63:i] * b[i+63:i]
//			dst[i+63:i] := tmp[63:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm_mullo_epi64'.
// Requires AVX512DQ.
func MulloEpi64(a M128i, b M128i) M128i {
	return M128i(mulloEpi64([16]byte(a), [16]byte(b)))
}

func mulloEpi64(a [16]byte, b [16]byte) [16]byte


// MaskMulloEpi641: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp[127:0] := a[i+63:i] * b[i+63:i]
//				dst[i+63:i] := tmp[63:0]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm256_mask_mullo_epi64'.
// Requires AVX512DQ.
func MaskMulloEpi641(src M256i, k Mmask8, a M256i, b M256i) M256i {
	return M256i(maskMulloEpi641([32]byte(src), uint8(k), [32]byte(a), [32]byte(b)))
}

func maskMulloEpi641(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte


// MaskzMulloEpi641: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp[127:0] := a[i+63:i] * b[i+63:i]
//				dst[i+63:i] := tmp[63:0]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm256_maskz_mullo_epi64'.
// Requires AVX512DQ.
func MaskzMulloEpi641(k Mmask8, a M256i, b M256i) M256i {
	return M256i(maskzMulloEpi641(uint8(k), [32]byte(a), [32]byte(b)))
}

func maskzMulloEpi641(k uint8, a [32]byte, b [32]byte) [32]byte


// MulloEpi641: Multiply the packed 64-bit integers in 'a' and 'b', producing
// intermediate 128-bit integers, and store the low 64 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			tmp[127:0] := a[i+63:i] * b[i+63:i]
//			dst[i+63:i] := tmp[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm256_mullo_epi64'.
// Requires AVX512DQ.
func MulloEpi641(a M256i, b M256i) M256i {
	return M256i(mulloEpi641([32]byte(a), [32]byte(b)))
}

func mulloEpi641(a [32]byte, b [32]byte) [32]byte


// MaskMulloEpi642: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp[127:0] := a[i+63:i] * b[i+63:i]
//				dst[i+63:i] := tmp[63:0]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm512_mask_mullo_epi64'.
// Requires AVX512DQ.
func MaskMulloEpi642(src M512i, k Mmask8, a M512i, b M512i) M512i {
	return M512i(maskMulloEpi642([64]byte(src), uint8(k), [64]byte(a), [64]byte(b)))
}

func maskMulloEpi642(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte


// MaskzMulloEpi642: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp[127:0] := a[i+63:i] * b[i+63:i]
//				dst[i+63:i] := tmp[63:0]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm512_maskz_mullo_epi64'.
// Requires AVX512DQ.
func MaskzMulloEpi642(k Mmask8, a M512i, b M512i) M512i {
	return M512i(maskzMulloEpi642(uint8(k), [64]byte(a), [64]byte(b)))
}

func maskzMulloEpi642(k uint8, a [64]byte, b [64]byte) [64]byte


// MulloEpi642: Multiply the packed 64-bit integers in 'a' and 'b', producing
// intermediate 128-bit integers, and store the low 64 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			tmp[127:0] := a[i+63:i] * b[i+63:i]
//			dst[i+63:i] := tmp[63:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLQ'. Intrinsic: '_mm512_mullo_epi64'.
// Requires AVX512DQ.
func MulloEpi642(a M512i, b M512i) M512i {
	return M512i(mulloEpi642([64]byte(a), [64]byte(b)))
}

func mulloEpi642(a [64]byte, b [64]byte) [64]byte


// MaskOrPd: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm_mask_or_pd'.
// Requires AVX512DQ.
func MaskOrPd(src M128d, k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskOrPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskOrPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


// MaskzOrPd: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm_maskz_or_pd'.
// Requires AVX512DQ.
func MaskzOrPd(k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskzOrPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzOrPd(k uint8, a [2]float64, b [2]float64) [2]float64


// MaskOrPd1: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm256_mask_or_pd'.
// Requires AVX512DQ.
func MaskOrPd1(src M256d, k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskOrPd1([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func maskOrPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// MaskzOrPd1: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm256_maskz_or_pd'.
// Requires AVX512DQ.
func MaskzOrPd1(k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskzOrPd1(uint8(k), [4]float64(a), [4]float64(b)))
}

func maskzOrPd1(k uint8, a [4]float64, b [4]float64) [4]float64


// MaskOrPd2: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm512_mask_or_pd'.
// Requires AVX512DQ.
func MaskOrPd2(src M512d, k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskOrPd2([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskOrPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// MaskzOrPd2: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm512_maskz_or_pd'.
// Requires AVX512DQ.
func MaskzOrPd2(k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskzOrPd2(uint8(k), [8]float64(a), [8]float64(b)))
}

func maskzOrPd2(k uint8, a [8]float64, b [8]float64) [8]float64


// OrPd: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm512_or_pd'.
// Requires AVX512DQ.
func OrPd(a M512d, b M512d) M512d {
	return M512d(orPd([8]float64(a), [8]float64(b)))
}

func orPd(a [8]float64, b [8]float64) [8]float64


// MaskOrPs: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm_mask_or_ps'.
// Requires AVX512DQ.
func MaskOrPs(src M128, k Mmask8, a M128, b M128) M128 {
	return M128(maskOrPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskOrPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


// MaskzOrPs: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm_maskz_or_ps'.
// Requires AVX512DQ.
func MaskzOrPs(k Mmask8, a M128, b M128) M128 {
	return M128(maskzOrPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzOrPs(k uint8, a [4]float32, b [4]float32) [4]float32


// MaskOrPs1: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm256_mask_or_ps'.
// Requires AVX512DQ.
func MaskOrPs1(src M256, k Mmask8, a M256, b M256) M256 {
	return M256(maskOrPs1([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func maskOrPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// MaskzOrPs1: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm256_maskz_or_ps'.
// Requires AVX512DQ.
func MaskzOrPs1(k Mmask8, a M256, b M256) M256 {
	return M256(maskzOrPs1(uint8(k), [8]float32(a), [8]float32(b)))
}

func maskzOrPs1(k uint8, a [8]float32, b [8]float32) [8]float32


// MaskOrPs2: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm512_mask_or_ps'.
// Requires AVX512DQ.
func MaskOrPs2(src M512, k Mmask16, a M512, b M512) M512 {
	return M512(maskOrPs2([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskOrPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// MaskzOrPs2: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm512_maskz_or_ps'.
// Requires AVX512DQ.
func MaskzOrPs2(k Mmask16, a M512, b M512) M512 {
	return M512(maskzOrPs2(uint16(k), [16]float32(a), [16]float32(b)))
}

func maskzOrPs2(k uint16, a [16]float32, b [16]float32) [16]float32


// OrPs: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm512_or_ps'.
// Requires AVX512DQ.
func OrPs(a M512, b M512) M512 {
	return M512(orPs([16]float32(a), [16]float32(b)))
}

func orPs(a [16]float32, b [16]float32) [16]float32


// MaskRangePd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm_mask_range_pd'.
// Requires AVX512DQ.
func MaskRangePd(src M128d, k Mmask8, a M128d, b M128d, imm8 int) M128d {
	return M128d(maskRangePd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8))
}

func maskRangePd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64


// MaskzRangePd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm_maskz_range_pd'.
// Requires AVX512DQ.
func MaskzRangePd(k Mmask8, a M128d, b M128d, imm8 int) M128d {
	return M128d(maskzRangePd(uint8(k), [2]float64(a), [2]float64(b), imm8))
}

func maskzRangePd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64


// RangePd: Calculate the max, min, absolute max, or absolute min (depending on
// control in 'imm8') for packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm_range_pd'.
// Requires AVX512DQ.
func RangePd(a M128d, b M128d, imm8 int) M128d {
	return M128d(rangePd([2]float64(a), [2]float64(b), imm8))
}

func rangePd(a [2]float64, b [2]float64, imm8 int) [2]float64


// MaskRangePd1: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm256_mask_range_pd'.
// Requires AVX512DQ.
func MaskRangePd1(src M256d, k Mmask8, a M256d, b M256d, imm8 int) M256d {
	return M256d(maskRangePd1([4]float64(src), uint8(k), [4]float64(a), [4]float64(b), imm8))
}

func maskRangePd1(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64


// MaskzRangePd1: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm256_maskz_range_pd'.
// Requires AVX512DQ.
func MaskzRangePd1(k Mmask8, a M256d, b M256d, imm8 int) M256d {
	return M256d(maskzRangePd1(uint8(k), [4]float64(a), [4]float64(b), imm8))
}

func maskzRangePd1(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64


// RangePd1: Calculate the max, min, absolute max, or absolute min (depending
// on control in 'imm8') for packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm256_range_pd'.
// Requires AVX512DQ.
func RangePd1(a M256d, b M256d, imm8 int) M256d {
	return M256d(rangePd1([4]float64(a), [4]float64(b), imm8))
}

func rangePd1(a [4]float64, b [4]float64, imm8 int) [4]float64


// MaskRangePd2: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm512_mask_range_pd'.
// Requires AVX512DQ.
func MaskRangePd2(src M512d, k Mmask8, a M512d, b M512d, imm8 int) M512d {
	return M512d(maskRangePd2([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), imm8))
}

func maskRangePd2(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64


// MaskzRangePd2: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm512_maskz_range_pd'.
// Requires AVX512DQ.
func MaskzRangePd2(k Mmask8, a M512d, b M512d, imm8 int) M512d {
	return M512d(maskzRangePd2(uint8(k), [8]float64(a), [8]float64(b), imm8))
}

func maskzRangePd2(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64


// RangePd2: Calculate the max, min, absolute max, or absolute min (depending
// on control in 'imm8') for packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm512_range_pd'.
// Requires AVX512DQ.
func RangePd2(a M512d, b M512d, imm8 int) M512d {
	return M512d(rangePd2([8]float64(a), [8]float64(b), imm8))
}

func rangePd2(a [8]float64, b [8]float64, imm8 int) [8]float64


// MaskRangePs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm_mask_range_ps'.
// Requires AVX512DQ.
func MaskRangePs(src M128, k Mmask8, a M128, b M128, imm8 int) M128 {
	return M128(maskRangePs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8))
}

func maskRangePs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32


// MaskzRangePs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm_maskz_range_ps'.
// Requires AVX512DQ.
func MaskzRangePs(k Mmask8, a M128, b M128, imm8 int) M128 {
	return M128(maskzRangePs(uint8(k), [4]float32(a), [4]float32(b), imm8))
}

func maskzRangePs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32


// RangePs: Calculate the max, min, absolute max, or absolute min (depending on
// control in 'imm8') for packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm_range_ps'.
// Requires AVX512DQ.
func RangePs(a M128, b M128, imm8 int) M128 {
	return M128(rangePs([4]float32(a), [4]float32(b), imm8))
}

func rangePs(a [4]float32, b [4]float32, imm8 int) [4]float32


// MaskRangePs1: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm256_mask_range_ps'.
// Requires AVX512DQ.
func MaskRangePs1(src M256, k Mmask8, a M256, b M256, imm8 int) M256 {
	return M256(maskRangePs1([8]float32(src), uint8(k), [8]float32(a), [8]float32(b), imm8))
}

func maskRangePs1(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32


// MaskzRangePs1: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm256_maskz_range_ps'.
// Requires AVX512DQ.
func MaskzRangePs1(k Mmask8, a M256, b M256, imm8 int) M256 {
	return M256(maskzRangePs1(uint8(k), [8]float32(a), [8]float32(b), imm8))
}

func maskzRangePs1(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32


// RangePs1: Calculate the max, min, absolute max, or absolute min (depending
// on control in 'imm8') for packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm256_range_ps'.
// Requires AVX512DQ.
func RangePs1(a M256, b M256, imm8 int) M256 {
	return M256(rangePs1([8]float32(a), [8]float32(b), imm8))
}

func rangePs1(a [8]float32, b [8]float32, imm8 int) [8]float32


// MaskRangePs2: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm512_mask_range_ps'.
// Requires AVX512DQ.
func MaskRangePs2(src M512, k Mmask16, a M512, b M512, imm8 int) M512 {
	return M512(maskRangePs2([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), imm8))
}

func maskRangePs2(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32


// MaskzRangePs2: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm512_maskz_range_ps'.
// Requires AVX512DQ.
func MaskzRangePs2(k Mmask16, a M512, b M512, imm8 int) M512 {
	return M512(maskzRangePs2(uint16(k), [16]float32(a), [16]float32(b), imm8))
}

func maskzRangePs2(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32


// RangePs2: Calculate the max, min, absolute max, or absolute min (depending
// on control in 'imm8') for packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm512_range_ps'.
// Requires AVX512DQ.
func RangePs2(a M512, b M512, imm8 int) M512 {
	return M512(rangePs2([16]float32(a), [16]float32(b), imm8))
}

func rangePs2(a [16]float32, b [16]float32, imm8 int) [16]float32


// MaskRangeRoundPd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm512_mask_range_round_pd'.
// Requires AVX512DQ.
func MaskRangeRoundPd(src M512d, k Mmask8, a M512d, b M512d, imm8 int, rounding int) M512d {
	return M512d(maskRangeRoundPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), imm8, rounding))
}

func maskRangeRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64


// MaskzRangeRoundPd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm512_maskz_range_round_pd'.
// Requires AVX512DQ.
func MaskzRangeRoundPd(k Mmask8, a M512d, b M512d, imm8 int, rounding int) M512d {
	return M512d(maskzRangeRoundPd(uint8(k), [8]float64(a), [8]float64(b), imm8, rounding))
}

func maskzRangeRoundPd(k uint8, a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64


// RangeRoundPd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := RANGE(a[i+63:i], b[i+63:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPD'. Intrinsic: '_mm512_range_round_pd'.
// Requires AVX512DQ.
func RangeRoundPd(a M512d, b M512d, imm8 int, rounding int) M512d {
	return M512d(rangeRoundPd([8]float64(a), [8]float64(b), imm8, rounding))
}

func rangeRoundPd(a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64


// MaskRangeRoundPs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm512_mask_range_round_ps'.
// Requires AVX512DQ.
func MaskRangeRoundPs(src M512, k Mmask16, a M512, b M512, imm8 int, rounding int) M512 {
	return M512(maskRangeRoundPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), imm8, rounding))
}

func maskRangeRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32


// MaskzRangeRoundPs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set).
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm512_maskz_range_round_ps'.
// Requires AVX512DQ.
func MaskzRangeRoundPs(k Mmask16, a M512, b M512, imm8 int, rounding int) M512 {
	return M512(maskzRangeRoundPs(uint16(k), [16]float32(a), [16]float32(b), imm8, rounding))
}

func maskzRangeRoundPs(k uint16, a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32


// RangeRoundPs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[63:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := RANGE(a[i+31:i], b[i+31:i], imm8[1:0], imm8[3:2])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRANGEPS'. Intrinsic: '_mm512_range_round_ps'.
// Requires AVX512DQ.
func RangeRoundPs(a M512, b M512, imm8 int, rounding int) M512 {
	return M512(rangeRoundPs([16]float32(a), [16]float32(b), imm8, rounding))
}

func rangeRoundPs(a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32


// MaskRangeRoundSd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower double-precision (64-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using writemask 'k' (the element is copied from 'src' when mask bit
// 0 is not set), and copy the upper element from 'a' to the upper element of
// 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[63:0]] := RANGE(a[63:0], b[63:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[63:0] := src[63:0]
//		FI
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESD'. Intrinsic: '_mm_mask_range_round_sd'.
// Requires AVX512DQ.
func MaskRangeRoundSd(src M128d, k Mmask8, a M128d, b M128d, imm8 int, rounding int) M128d {
	return M128d(maskRangeRoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
}

func maskRangeRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64


// MaskzRangeRoundSd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower double-precision (64-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using zeromask 'k' (the element is zeroed out when mask bit 0 is
// not set), and copy the upper element from 'a' to the upper element of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[63:0]] := RANGE(a[63:0], b[63:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[63:0] := 0
//		FI
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESD'. Intrinsic: '_mm_maskz_range_round_sd'.
// Requires AVX512DQ.
func MaskzRangeRoundSd(k Mmask8, a M128d, b M128d, imm8 int, rounding int) M128d {
	return M128d(maskzRangeRoundSd(uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
}

func maskzRangeRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64


// RangeRoundSd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower double-precision (64-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst', and copy the upper element from 'a' to the upper element of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		dst[63:0]] := RANGE(a[63:0], b[63:0], imm8[1:0], imm8[3:2])
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESD'. Intrinsic: '_mm_range_round_sd'.
// Requires AVX512DQ.
func RangeRoundSd(a M128d, b M128d, imm8 int, rounding int) M128d {
	return M128d(rangeRoundSd([2]float64(a), [2]float64(b), imm8, rounding))
}

func rangeRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64


// MaskRangeRoundSs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower single-precision (32-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using writemask 'k' (the element is copied from 'src' when mask bit
// 0 is not set), and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[31:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[31:0]] := RANGE(a[31:0], b[31:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[31:0] := src[31:0]
//		FI
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESS'. Intrinsic: '_mm_mask_range_round_ss'.
// Requires AVX512DQ.
func MaskRangeRoundSs(src M128, k Mmask8, a M128, b M128, imm8 int, rounding int) M128 {
	return M128(maskRangeRoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
}

func maskRangeRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32


// MaskzRangeRoundSs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower single-precision (32-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using zeromask 'k' (the element is zeroed out when mask bit 0 is
// not set), and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[31:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[31:0]] := RANGE(a[31:0], b[31:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[31:0] := 0
//		FI
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESS'. Intrinsic: '_mm_maskz_range_round_ss'.
// Requires AVX512DQ.
func MaskzRangeRoundSs(k Mmask8, a M128, b M128, imm8 int, rounding int) M128 {
	return M128(maskzRangeRoundSs(uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
}

func maskzRangeRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32


// RangeRoundSs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower single-precision (32-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[31:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		dst[31:0]] := RANGE(a[31:0], b[31:0], imm8[1:0], imm8[3:2])
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESS'. Intrinsic: '_mm_range_round_ss'.
// Requires AVX512DQ.
func RangeRoundSs(a M128, b M128, imm8 int, rounding int) M128 {
	return M128(rangeRoundSs([4]float32(a), [4]float32(b), imm8, rounding))
}

func rangeRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32


// MaskRangeSd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower double-precision (64-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using writemask 'k' (the element is copied from 'src' when mask bit
// 0 is not set), and copy the upper element from 'a' to the upper element of
// 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[63:0]] := RANGE(a[63:0], b[63:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[63:0] := src[63:0]
//		FI
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESD'. Intrinsic: '_mm_mask_range_sd'.
// Requires AVX512DQ.
func MaskRangeSd(src M128d, k Mmask8, a M128d, b M128d, imm8 int) M128d {
	return M128d(maskRangeSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8))
}

func maskRangeSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64


// MaskzRangeSd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower double-precision (64-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using zeromask 'k' (the element is zeroed out when mask bit 0 is
// not set), and copy the upper element from 'a' to the upper element of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[63:0], src2[63:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src1[63:0] : src2[63:0]
//			1: tmp[63:0] := (src1[63:0] <= src2[63:0]) ? src2[63:0] : src1[63:0]
//			2: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src1[63:0] : src2[63:0]
//			3: tmp[63:0] := (ABS(src1[63:0]) <= ABS(src2[63:0])) ? src2[63:0] : src1[63:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[63:0] := (src1[63] << 63) OR (tmp[62:0])
//			1: dst[63:0] := tmp[63:0]
//			2: dst[63:0] := (0 << 63) OR (tmp[62:0])
//			3: dst[63:0] := (1 << 63) OR (tmp[62:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[63:0]] := RANGE(a[63:0], b[63:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[63:0] := 0
//		FI
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESD'. Intrinsic: '_mm_maskz_range_sd'.
// Requires AVX512DQ.
func MaskzRangeSd(k Mmask8, a M128d, b M128d, imm8 int) M128d {
	return M128d(maskzRangeSd(uint8(k), [2]float64(a), [2]float64(b), imm8))
}

func maskzRangeSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64


// MaskRangeSs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower single-precision (32-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using writemask 'k' (the element is copied from 'src' when mask bit
// 0 is not set), and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[31:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[31:0]] := RANGE(a[31:0], b[31:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[31:0] := src[31:0]
//		FI
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESS'. Intrinsic: '_mm_mask_range_ss'.
// Requires AVX512DQ.
func MaskRangeSs(src M128, k Mmask8, a M128, b M128, imm8 int) M128 {
	return M128(maskRangeSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8))
}

func maskRangeSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32


// MaskzRangeSs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for the lower single-precision (32-bit)
// floating-point element in 'a' and 'b', store the result in the lower element
// of 'dst' using zeromask 'k' (the element is zeroed out when mask bit 0 is
// not set), and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'.
// 	imm8[1:0] specifies the operation control: 00 = min, 01 = max, 10 =
// absolute max, 11 = absolute min.
// 	imm8[3:2] specifies the sign control: 00 = sign from a, 01 = sign from
// compare result, 10 = clear sign bit, 11 = set sign bit. 
//
//		RANGE(src1[31:0], src2[31:0], opCtl[1:0], signSelCtl[1:0])
//		{
//			CASE opCtl[1:0]
//			0: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src1[31:0] : src2[31:0]
//			1: tmp[31:0] := (src1[31:0] <= src2[31:0]) ? src2[31:0] : src1[31:0]
//			2: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src1[31:0] : src2[31:0]
//			3: tmp[31:0] := (ABS(src1[31:0]) <= ABS(src2[31:0])) ? src2[31:0] : src1[31:0]
//			ESAC
//			
//			CASE signSelCtl[1:0]
//			0: dst[31:0] := (src1[31] << 31) OR (tmp[30:0])
//			1: dst[31:0] := tmp[31:0]
//			2: dst[31:0] := (0 << 31) OR (tmp[30:0])
//			3: dst[31:0] := (1 << 31) OR (tmp[30:0])
//			ESAC
//			
//			RETURN dst
//		}
//		
//		IF k[0]
//			dst[31:0]] := RANGE(a[31:0], b[31:0], imm8[1:0], imm8[3:2])
//		ELSE
//			dst[31:0] := 0
//		FI
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VRANGESS'. Intrinsic: '_mm_maskz_range_ss'.
// Requires AVX512DQ.
func MaskzRangeSs(k Mmask8, a M128, b M128, imm8 int) M128 {
	return M128(maskzRangeSs(uint8(k), [4]float32(a), [4]float32(b), imm8))
}

func maskzRangeSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32


// MaskReducePd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm_mask_reduce_pd'.
// Requires AVX512DQ.
func MaskReducePd(src M128d, k Mmask8, a M128d, imm8 int) M128d {
	return M128d(maskReducePd([2]float64(src), uint8(k), [2]float64(a), imm8))
}

func maskReducePd(src [2]float64, k uint8, a [2]float64, imm8 int) [2]float64


// MaskzReducePd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm_maskz_reduce_pd'.
// Requires AVX512DQ.
func MaskzReducePd(k Mmask8, a M128d, imm8 int) M128d {
	return M128d(maskzReducePd(uint8(k), [2]float64(a), imm8))
}

func maskzReducePd(k uint8, a [2]float64, imm8 int) [2]float64


// ReducePd: Extract the reduced argument of packed double-precision (64-bit)
// floating-point elements in 'a' by the number of bits specified by 'imm8',
// and store the results in 'dst'. 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm_reduce_pd'.
// Requires AVX512DQ.
func ReducePd(a M128d, imm8 int) M128d {
	return M128d(reducePd([2]float64(a), imm8))
}

func reducePd(a [2]float64, imm8 int) [2]float64


// MaskReducePd1: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm256_mask_reduce_pd'.
// Requires AVX512DQ.
func MaskReducePd1(src M256d, k Mmask8, a M256d, imm8 int) M256d {
	return M256d(maskReducePd1([4]float64(src), uint8(k), [4]float64(a), imm8))
}

func maskReducePd1(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64


// MaskzReducePd1: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm256_maskz_reduce_pd'.
// Requires AVX512DQ.
func MaskzReducePd1(k Mmask8, a M256d, imm8 int) M256d {
	return M256d(maskzReducePd1(uint8(k), [4]float64(a), imm8))
}

func maskzReducePd1(k uint8, a [4]float64, imm8 int) [4]float64


// ReducePd1: Extract the reduced argument of packed double-precision (64-bit)
// floating-point elements in 'a' by the number of bits specified by 'imm8',
// and store the results in 'dst'. 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm256_reduce_pd'.
// Requires AVX512DQ.
func ReducePd1(a M256d, imm8 int) M256d {
	return M256d(reducePd1([4]float64(a), imm8))
}

func reducePd1(a [4]float64, imm8 int) [4]float64


// MaskReducePd2: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm512_mask_reduce_pd'.
// Requires AVX512DQ.
func MaskReducePd2(src M512d, k Mmask8, a M512d, imm8 int) M512d {
	return M512d(maskReducePd2([8]float64(src), uint8(k), [8]float64(a), imm8))
}

func maskReducePd2(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64


// MaskzReducePd2: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm512_maskz_reduce_pd'.
// Requires AVX512DQ.
func MaskzReducePd2(k Mmask8, a M512d, imm8 int) M512d {
	return M512d(maskzReducePd2(uint8(k), [8]float64(a), imm8))
}

func maskzReducePd2(k uint8, a [8]float64, imm8 int) [8]float64


// ReducePd2: Extract the reduced argument of packed double-precision (64-bit)
// floating-point elements in 'a' by the number of bits specified by 'imm8',
// and store the results in 'dst'. 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm512_reduce_pd'.
// Requires AVX512DQ.
func ReducePd2(a M512d, imm8 int) M512d {
	return M512d(reducePd2([8]float64(a), imm8))
}

func reducePd2(a [8]float64, imm8 int) [8]float64


// MaskReducePs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm_mask_reduce_ps'.
// Requires AVX512DQ.
func MaskReducePs(src M128, k Mmask8, a M128, imm8 int) M128 {
	return M128(maskReducePs([4]float32(src), uint8(k), [4]float32(a), imm8))
}

func maskReducePs(src [4]float32, k uint8, a [4]float32, imm8 int) [4]float32


// MaskzReducePs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm_maskz_reduce_ps'.
// Requires AVX512DQ.
func MaskzReducePs(k Mmask8, a M128, imm8 int) M128 {
	return M128(maskzReducePs(uint8(k), [4]float32(a), imm8))
}

func maskzReducePs(k uint8, a [4]float32, imm8 int) [4]float32


// ReducePs: Extract the reduced argument of packed single-precision (32-bit)
// floating-point elements in 'a' by the number of bits specified by 'imm8',
// and store the results in 'dst'. 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm_reduce_ps'.
// Requires AVX512DQ.
func ReducePs(a M128, imm8 int) M128 {
	return M128(reducePs([4]float32(a), imm8))
}

func reducePs(a [4]float32, imm8 int) [4]float32


// MaskReducePs1: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm256_mask_reduce_ps'.
// Requires AVX512DQ.
func MaskReducePs1(src M256, k Mmask8, a M256, imm8 int) M256 {
	return M256(maskReducePs1([8]float32(src), uint8(k), [8]float32(a), imm8))
}

func maskReducePs1(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32


// MaskzReducePs1: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm256_maskz_reduce_ps'.
// Requires AVX512DQ.
func MaskzReducePs1(k Mmask8, a M256, imm8 int) M256 {
	return M256(maskzReducePs1(uint8(k), [8]float32(a), imm8))
}

func maskzReducePs1(k uint8, a [8]float32, imm8 int) [8]float32


// ReducePs1: Extract the reduced argument of packed single-precision (32-bit)
// floating-point elements in 'a' by the number of bits specified by 'imm8',
// and store the results in 'dst'. 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm256_reduce_ps'.
// Requires AVX512DQ.
func ReducePs1(a M256, imm8 int) M256 {
	return M256(reducePs1([8]float32(a), imm8))
}

func reducePs1(a [8]float32, imm8 int) [8]float32


// MaskReducePs2: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm512_mask_reduce_ps'.
// Requires AVX512DQ.
func MaskReducePs2(src M512, k Mmask16, a M512, imm8 int) M512 {
	return M512(maskReducePs2([16]float32(src), uint16(k), [16]float32(a), imm8))
}

func maskReducePs2(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32


// MaskzReducePs2: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm512_maskz_reduce_ps'.
// Requires AVX512DQ.
func MaskzReducePs2(k Mmask16, a M512, imm8 int) M512 {
	return M512(maskzReducePs2(uint16(k), [16]float32(a), imm8))
}

func maskzReducePs2(k uint16, a [16]float32, imm8 int) [16]float32


// ReducePs2: Extract the reduced argument of packed single-precision (32-bit)
// floating-point elements in 'a' by the number of bits specified by 'imm8',
// and store the results in 'dst'. 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm512_reduce_ps'.
// Requires AVX512DQ.
func ReducePs2(a M512, imm8 int) M512 {
	return M512(reducePs2([16]float32(a), imm8))
}

func reducePs2(a [16]float32, imm8 int) [16]float32


// MaskReduceRoundPd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm512_mask_reduce_round_pd'.
// Requires AVX512DQ.
func MaskReduceRoundPd(src M512d, k Mmask8, a M512d, imm8 int, rounding int) M512d {
	return M512d(maskReduceRoundPd([8]float64(src), uint8(k), [8]float64(a), imm8, rounding))
}

func maskReduceRoundPd(src [8]float64, k uint8, a [8]float64, imm8 int, rounding int) [8]float64


// MaskzReduceRoundPd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm512_maskz_reduce_round_pd'.
// Requires AVX512DQ.
func MaskzReduceRoundPd(k Mmask8, a M512d, imm8 int, rounding int) M512d {
	return M512d(maskzReduceRoundPd(uint8(k), [8]float64(a), imm8, rounding))
}

func maskzReduceRoundPd(k uint8, a [8]float64, imm8 int, rounding int) [8]float64


// ReduceRoundPd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ReduceArgumentPD(src[i+63:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPD'. Intrinsic: '_mm512_reduce_round_pd'.
// Requires AVX512DQ.
func ReduceRoundPd(a M512d, imm8 int, rounding int) M512d {
	return M512d(reduceRoundPd([8]float64(a), imm8, rounding))
}

func reduceRoundPd(a [8]float64, imm8 int, rounding int) [8]float64


// MaskReduceRoundPs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm512_mask_reduce_round_ps'.
// Requires AVX512DQ.
func MaskReduceRoundPs(src M512, k Mmask16, a M512, imm8 int, rounding int) M512 {
	return M512(maskReduceRoundPs([16]float32(src), uint16(k), [16]float32(a), imm8, rounding))
}

func maskReduceRoundPs(src [16]float32, k uint16, a [16]float32, imm8 int, rounding int) [16]float32


// MaskzReduceRoundPs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm512_maskz_reduce_round_ps'.
// Requires AVX512DQ.
func MaskzReduceRoundPs(k Mmask16, a M512, imm8 int, rounding int) M512 {
	return M512(maskzReduceRoundPs(uint16(k), [16]float32(a), imm8, rounding))
}

func maskzReduceRoundPs(k uint16, a [16]float32, imm8 int, rounding int) [16]float32


// ReduceRoundPs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ReduceArgumentPS(src[i+31:i], imm8[7:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VREDUCEPS'. Intrinsic: '_mm512_reduce_round_ps'.
// Requires AVX512DQ.
func ReduceRoundPs(a M512, imm8 int, rounding int) M512 {
	return M512(reduceRoundPs([16]float32(a), imm8, rounding))
}

func reduceRoundPs(a [16]float32, imm8 int, rounding int) [16]float32


// MaskReduceRoundSd: Extract the reduced argument of the lower
// double-precision (64-bit) floating-point element in 'a' by the number of
// bits specified by 'imm8', store the result in the lower element of 'dst'
// using writemask 'k' (the element is copied from 'src' when mask bit 0 is not
// set), and copy the upper element from 'b' to the upper element of 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		IF k[0]
//			dst[63:0] := ReduceArgumentPD(a[63:0], imm8[7:0])
//		ELSE
//			dst[63:0] := src[63:0]
//		FI
//		dst[127:64] := b[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESD'. Intrinsic: '_mm_mask_reduce_round_sd'.
// Requires AVX512DQ.
func MaskReduceRoundSd(src M128d, k Mmask8, a M128d, b M128d, imm8 int, rounding int) M128d {
	return M128d(maskReduceRoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
}

func maskReduceRoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64


// MaskzReduceRoundSd: Extract the reduced argument of the lower
// double-precision (64-bit) floating-point element in 'a' by the number of
// bits specified by 'imm8', store the result in the lower element of 'dst'
// using zeromask 'k' (the element is zeroed out when mask bit 0 is not set),
// and copy the upper element from 'b' to the upper element of 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		IF k[0]
//			dst[63:0] := ReduceArgumentPD(a[63:0], imm8[7:0])
//		ELSE
//			dst[63:0] := 0
//		FI
//		dst[127:64] := b[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESD'. Intrinsic: '_mm_maskz_reduce_round_sd'.
// Requires AVX512DQ.
func MaskzReduceRoundSd(k Mmask8, a M128d, b M128d, imm8 int, rounding int) M128d {
	return M128d(maskzReduceRoundSd(uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
}

func maskzReduceRoundSd(k uint8, a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64


// ReduceRoundSd: Extract the reduced argument of the lower double-precision
// (64-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst', and copy the upper
// element from 'b' to the upper element of 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		dst[63:0] := ReduceArgumentPD(a[63:0], imm8[7:0])
//		dst[127:64] := b[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESD'. Intrinsic: '_mm_reduce_round_sd'.
// Requires AVX512DQ.
func ReduceRoundSd(a M128d, b M128d, imm8 int, rounding int) M128d {
	return M128d(reduceRoundSd([2]float64(a), [2]float64(b), imm8, rounding))
}

func reduceRoundSd(a [2]float64, b [2]float64, imm8 int, rounding int) [2]float64


// MaskReduceRoundSs: Extract the reduced argument of the lower
// single-precision (32-bit) floating-point element in 'a' by the number of
// bits specified by 'imm8', store the result in the lower element of 'dst'
// using writemask 'k' (the element is copied from 'src' when mask bit 0 is not
// set), and copy the upper 3 packed elements from 'b' to the upper elements of
// 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		
//		IF k[0]
//			dst[31:0] := ReduceArgumentPS(a[31:0], imm8[7:0])
//		ELSE
//			dst[31:0] := src[31:0]
//		FI
//		dst[127:64] := b[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESS'. Intrinsic: '_mm_mask_reduce_round_ss'.
// Requires AVX512DQ.
func MaskReduceRoundSs(src M128, k Mmask8, a M128, b M128, imm8 int, rounding int) M128 {
	return M128(maskReduceRoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
}

func maskReduceRoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32


// MaskzReduceRoundSs: Extract the reduced argument of the lower
// single-precision (32-bit) floating-point element in 'a' by the number of
// bits specified by 'imm8', store the result in the lower element of 'dst'
// using zeromask 'k' (the element is zeroed out when mask bit 0 is not set),
// and copy the upper 3 packed elements from 'b' to the upper elements of
// 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		
//		IF k[0]
//			dst[31:0] := ReduceArgumentPS(a[31:0], imm8[7:0])
//		ELSE
//			dst[31:0] := 0
//		FI
//		dst[127:64] := b[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESS'. Intrinsic: '_mm_maskz_reduce_round_ss'.
// Requires AVX512DQ.
func MaskzReduceRoundSs(k Mmask8, a M128, b M128, imm8 int, rounding int) M128 {
	return M128(maskzReduceRoundSs(uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
}

func maskzReduceRoundSs(k uint8, a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32


// ReduceRoundSs: Extract the reduced argument of the lower single-precision
// (32-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst', and copy the upper 3
// packed elements from 'b' to the upper elements of 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		
//		dst[31:0] := ReduceArgumentPS(a[31:0], imm8[7:0])
//		dst[127:64] := b[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESS'. Intrinsic: '_mm_reduce_round_ss'.
// Requires AVX512DQ.
func ReduceRoundSs(a M128, b M128, imm8 int, rounding int) M128 {
	return M128(reduceRoundSs([4]float32(a), [4]float32(b), imm8, rounding))
}

func reduceRoundSs(a [4]float32, b [4]float32, imm8 int, rounding int) [4]float32


// MaskReduceSd: Extract the reduced argument of the lower double-precision
// (64-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst' using writemask 'k'
// (the element is copied from 'src' when mask bit 0 is not set), and copy the
// upper element from 'b' to the upper element of 'dst'. 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		IF k[0]
//			dst[63:0] := ReduceArgumentPD(a[63:0], imm8[7:0])
//		ELSE
//			dst[63:0] := src[63:0]
//		FI
//		dst[127:64] := b[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESD'. Intrinsic: '_mm_mask_reduce_sd'.
// Requires AVX512DQ.
func MaskReduceSd(src M128d, k Mmask8, a M128d, b M128d, imm8 int) M128d {
	return M128d(maskReduceSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8))
}

func maskReduceSd(src [2]float64, k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64


// MaskzReduceSd: Extract the reduced argument of the lower double-precision
// (64-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst' using zeromask 'k'
// (the element is zeroed out when mask bit 0 is not set), and copy the upper
// element from 'b' to the upper element of 'dst'. 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		IF k[0]
//			dst[63:0] := ReduceArgumentPD(a[63:0], imm8[7:0])
//		ELSE
//			dst[63:0] := 0
//		FI
//		dst[127:64] := b[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESD'. Intrinsic: '_mm_maskz_reduce_sd'.
// Requires AVX512DQ.
func MaskzReduceSd(k Mmask8, a M128d, b M128d, imm8 int) M128d {
	return M128d(maskzReduceSd(uint8(k), [2]float64(a), [2]float64(b), imm8))
}

func maskzReduceSd(k uint8, a [2]float64, b [2]float64, imm8 int) [2]float64


// ReduceSd: Extract the reduced argument of the lower double-precision
// (64-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst', and copy the upper
// element from 'b' to the upper element of 'dst'. 
//
//		ReduceArgumentPD(src1[63:0], imm8[7:0])
//		{
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[63:0] := pow(2, -m) * ROUND(pow(2, m) * src1[63:0], spe, rc_source, rc)
//			tmp[63:0] := src1[63:0] - tmp[63:0]
//			RETURN tmp[63:0]
//		}
//		
//		dst[63:0] := ReduceArgumentPD(a[63:0], imm8[7:0])
//		dst[127:64] := b[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESD'. Intrinsic: '_mm_reduce_sd'.
// Requires AVX512DQ.
func ReduceSd(a M128d, b M128d, imm8 int) M128d {
	return M128d(reduceSd([2]float64(a), [2]float64(b), imm8))
}

func reduceSd(a [2]float64, b [2]float64, imm8 int) [2]float64


// MaskReduceSs: Extract the reduced argument of the lower single-precision
// (32-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst' using writemask 'k'
// (the element is copied from 'src' when mask bit 0 is not set), and copy the
// upper 3 packed elements from 'b' to the upper elements of 'dst'. 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		
//		IF k[0]
//			dst[31:0] := ReduceArgumentPS(a[31:0], imm8[7:0])
//		ELSE
//			dst[31:0] := src[31:0]
//		FI
//		dst[127:64] := b[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESS'. Intrinsic: '_mm_mask_reduce_ss'.
// Requires AVX512DQ.
func MaskReduceSs(src M128, k Mmask8, a M128, b M128, imm8 int) M128 {
	return M128(maskReduceSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8))
}

func maskReduceSs(src [4]float32, k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32


// MaskzReduceSs: Extract the reduced argument of the lower single-precision
// (32-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst' using zeromask 'k'
// (the element is zeroed out when mask bit 0 is not set), and copy the upper 3
// packed elements from 'b' to the upper elements of 'dst'. 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		
//		IF k[0]
//			dst[31:0] := ReduceArgumentPS(a[31:0], imm8[7:0])
//		ELSE
//			dst[31:0] := 0
//		FI
//		dst[127:64] := b[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESS'. Intrinsic: '_mm_maskz_reduce_ss'.
// Requires AVX512DQ.
func MaskzReduceSs(k Mmask8, a M128, b M128, imm8 int) M128 {
	return M128(maskzReduceSs(uint8(k), [4]float32(a), [4]float32(b), imm8))
}

func maskzReduceSs(k uint8, a [4]float32, b [4]float32, imm8 int) [4]float32


// ReduceSs: Extract the reduced argument of the lower single-precision
// (32-bit) floating-point element in 'a' by the number of bits specified by
// 'imm8', store the result in the lower element of 'dst', and copy the upper 3
// packed elements from 'b' to the upper elements of 'dst'. 
//
//		ReduceArgumentPS(src1[31:0], imm8[7:0])
//		{
//			IF src1[31:0] == NAN
//				RETURN (convert src1[31:0] to QNaN)
//			FI
//			
//			m := imm8[7:4] // number of fraction bits after the binary point to be preserved
//			rc := imm8[1:0] // round control
//			rc_src := imm8[2] // round ccontrol source
//			spe := 0
//			tmp[31:0] := pow(2, -m)*ROUND(pow(2, m)*src1[31:0], spe, rc_source, rc)
//			tmp[31:0] := src1[31:0] - tmp[31:0]
//			RETURN tmp[31:0]
//		}
//		
//		dst[31:0] := ReduceArgumentPS(a[31:0], imm8[7:0])
//		dst[127:64] := b[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VREDUCESS'. Intrinsic: '_mm_reduce_ss'.
// Requires AVX512DQ.
func ReduceSs(a M128, b M128, imm8 int) M128 {
	return M128(reduceSs([4]float32(a), [4]float32(b), imm8))
}

func reduceSs(a [4]float32, b [4]float32, imm8 int) [4]float32


// MaskXorPd: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm_mask_xor_pd'.
// Requires AVX512DQ.
func MaskXorPd(src M128d, k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskXorPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskXorPd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


// MaskzXorPd: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm_maskz_xor_pd'.
// Requires AVX512DQ.
func MaskzXorPd(k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskzXorPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzXorPd(k uint8, a [2]float64, b [2]float64) [2]float64


// MaskXorPd1: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm256_mask_xor_pd'.
// Requires AVX512DQ.
func MaskXorPd1(src M256d, k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskXorPd1([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func maskXorPd1(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// MaskzXorPd1: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm256_maskz_xor_pd'.
// Requires AVX512DQ.
func MaskzXorPd1(k Mmask8, a M256d, b M256d) M256d {
	return M256d(maskzXorPd1(uint8(k), [4]float64(a), [4]float64(b)))
}

func maskzXorPd1(k uint8, a [4]float64, b [4]float64) [4]float64


// MaskXorPd2: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm512_mask_xor_pd'.
// Requires AVX512DQ.
func MaskXorPd2(src M512d, k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskXorPd2([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskXorPd2(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// MaskzXorPd2: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm512_maskz_xor_pd'.
// Requires AVX512DQ.
func MaskzXorPd2(k Mmask8, a M512d, b M512d) M512d {
	return M512d(maskzXorPd2(uint8(k), [8]float64(a), [8]float64(b)))
}

func maskzXorPd2(k uint8, a [8]float64, b [8]float64) [8]float64


// XorPd: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm512_xor_pd'.
// Requires AVX512DQ.
func XorPd(a M512d, b M512d) M512d {
	return M512d(xorPd([8]float64(a), [8]float64(b)))
}

func xorPd(a [8]float64, b [8]float64) [8]float64


// MaskXorPs: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm_mask_xor_ps'.
// Requires AVX512DQ.
func MaskXorPs(src M128, k Mmask8, a M128, b M128) M128 {
	return M128(maskXorPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskXorPs(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


// MaskzXorPs: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm_maskz_xor_ps'.
// Requires AVX512DQ.
func MaskzXorPs(k Mmask8, a M128, b M128) M128 {
	return M128(maskzXorPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzXorPs(k uint8, a [4]float32, b [4]float32) [4]float32


// MaskXorPs1: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm256_mask_xor_ps'.
// Requires AVX512DQ.
func MaskXorPs1(src M256, k Mmask8, a M256, b M256) M256 {
	return M256(maskXorPs1([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func maskXorPs1(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// MaskzXorPs1: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm256_maskz_xor_ps'.
// Requires AVX512DQ.
func MaskzXorPs1(k Mmask8, a M256, b M256) M256 {
	return M256(maskzXorPs1(uint8(k), [8]float32(a), [8]float32(b)))
}

func maskzXorPs1(k uint8, a [8]float32, b [8]float32) [8]float32


// MaskXorPs2: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm512_mask_xor_ps'.
// Requires AVX512DQ.
func MaskXorPs2(src M512, k Mmask16, a M512, b M512) M512 {
	return M512(maskXorPs2([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskXorPs2(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// MaskzXorPs2: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm512_maskz_xor_ps'.
// Requires AVX512DQ.
func MaskzXorPs2(k Mmask16, a M512, b M512) M512 {
	return M512(maskzXorPs2(uint16(k), [16]float32(a), [16]float32(b)))
}

func maskzXorPs2(k uint16, a [16]float32, b [16]float32) [16]float32


// XorPs: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm512_xor_ps'.
// Requires AVX512DQ.
func XorPs(a M512, b M512) M512 {
	return M512(xorPs([16]float32(a), [16]float32(b)))
}

func xorPs(a [16]float32, b [16]float32) [16]float32

