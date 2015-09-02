package avx512dq

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


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
func MaskAndPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskAndPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
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
func MaskzAndPd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzAndPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzAndPd(k uint8, a [2]float64, b [2]float64) [2]float64


// M256MaskAndPd: Compute the bitwise AND of packed double-precision (64-bit)
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
func M256MaskAndPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskAndPd([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskAndPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// M256MaskzAndPd: Compute the bitwise AND of packed double-precision (64-bit)
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
func M256MaskzAndPd(k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskzAndPd(uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskzAndPd(k uint8, a [4]float64, b [4]float64) [4]float64


// M512AndPd: Compute the bitwise AND of packed double-precision (64-bit)
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
func M512AndPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512AndPd([8]float64(a), [8]float64(b)))
}

func m512AndPd(a [8]float64, b [8]float64) [8]float64


// M512MaskAndPd: Compute the bitwise AND of packed double-precision (64-bit)
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
func M512MaskAndPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskAndPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskAndPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// M512MaskzAndPd: Compute the bitwise AND of packed double-precision (64-bit)
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
func M512MaskzAndPd(k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzAndPd(uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskzAndPd(k uint8, a [8]float64, b [8]float64) [8]float64


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
func MaskAndPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskAndPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
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
func MaskzAndPs(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzAndPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzAndPs(k uint8, a [4]float32, b [4]float32) [4]float32


// M256MaskAndPs: Compute the bitwise AND of packed single-precision (32-bit)
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
func M256MaskAndPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskAndPs([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskAndPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// M256MaskzAndPs: Compute the bitwise AND of packed single-precision (32-bit)
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
func M256MaskzAndPs(k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskzAndPs(uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskzAndPs(k uint8, a [8]float32, b [8]float32) [8]float32


// M512AndPs: Compute the bitwise AND of packed single-precision (32-bit)
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
func M512AndPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512AndPs([16]float32(a), [16]float32(b)))
}

func m512AndPs(a [16]float32, b [16]float32) [16]float32


// M512MaskAndPs: Compute the bitwise AND of packed single-precision (32-bit)
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
func M512MaskAndPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskAndPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskAndPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// M512MaskzAndPs: Compute the bitwise AND of packed single-precision (32-bit)
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
func M512MaskzAndPs(k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskzAndPs(uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskzAndPs(k uint16, a [16]float32, b [16]float32) [16]float32


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
func MaskAndnotPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskAndnotPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
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
func MaskzAndnotPd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzAndnotPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzAndnotPd(k uint8, a [2]float64, b [2]float64) [2]float64


// M256MaskAndnotPd: Compute the bitwise AND NOT of packed double-precision
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
func M256MaskAndnotPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskAndnotPd([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskAndnotPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// M256MaskzAndnotPd: Compute the bitwise AND NOT of packed double-precision
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
func M256MaskzAndnotPd(k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskzAndnotPd(uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskzAndnotPd(k uint8, a [4]float64, b [4]float64) [4]float64


// M512AndnotPd: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm512_andnot_pd'.
// Requires AVX512DQ.
func M512AndnotPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512AndnotPd([8]float64(a), [8]float64(b)))
}

func m512AndnotPd(a [8]float64, b [8]float64) [8]float64


// M512MaskAndnotPd: Compute the bitwise AND NOT of packed double-precision
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
func M512MaskAndnotPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskAndnotPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskAndnotPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// M512MaskzAndnotPd: Compute the bitwise AND NOT of packed double-precision
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
func M512MaskzAndnotPd(k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzAndnotPd(uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskzAndnotPd(k uint8, a [8]float64, b [8]float64) [8]float64


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
func MaskAndnotPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskAndnotPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
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
func MaskzAndnotPs(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzAndnotPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzAndnotPs(k uint8, a [4]float32, b [4]float32) [4]float32


// M256MaskAndnotPs: Compute the bitwise AND NOT of packed single-precision
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
func M256MaskAndnotPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskAndnotPs([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskAndnotPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// M256MaskzAndnotPs: Compute the bitwise AND NOT of packed single-precision
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
func M256MaskzAndnotPs(k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskzAndnotPs(uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskzAndnotPs(k uint8, a [8]float32, b [8]float32) [8]float32


// M512AndnotPs: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm512_andnot_ps'.
// Requires AVX512DQ.
func M512AndnotPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512AndnotPs([16]float32(a), [16]float32(b)))
}

func m512AndnotPs(a [16]float32, b [16]float32) [16]float32


// M512MaskAndnotPs: Compute the bitwise AND NOT of packed single-precision
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
func M512MaskAndnotPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskAndnotPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskAndnotPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// M512MaskzAndnotPs: Compute the bitwise AND NOT of packed single-precision
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
func M512MaskzAndnotPs(k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskzAndnotPs(uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskzAndnotPs(k uint16, a [16]float32, b [16]float32) [16]float32


// M256BroadcastF32x2: Broadcast the lower 2 packed single-precision (32-bit)
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
func M256BroadcastF32x2(a x86.M128) x86.M256 {
	return x86.M256(m256BroadcastF32x2([4]float32(a)))
}

func m256BroadcastF32x2(a [4]float32) [8]float32


// M256MaskBroadcastF32x2: Broadcast the lower 2 packed single-precision
// (32-bit) floating-point elements from 'a' to all elements of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func M256MaskBroadcastF32x2(src x86.M256, k x86.Mmask8, a x86.M128) x86.M256 {
	return x86.M256(m256MaskBroadcastF32x2([8]float32(src), uint8(k), [4]float32(a)))
}

func m256MaskBroadcastF32x2(src [8]float32, k uint8, a [4]float32) [8]float32


// M256MaskzBroadcastF32x2: Broadcast the lower 2 packed single-precision
// (32-bit) floating-point elements from 'a' to all elements of 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func M256MaskzBroadcastF32x2(k x86.Mmask8, a x86.M128) x86.M256 {
	return x86.M256(m256MaskzBroadcastF32x2(uint8(k), [4]float32(a)))
}

func m256MaskzBroadcastF32x2(k uint8, a [4]float32) [8]float32


// M512BroadcastF32x2: Broadcast the lower 2 packed single-precision (32-bit)
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
func M512BroadcastF32x2(a x86.M128) x86.M512 {
	return x86.M512(m512BroadcastF32x2([4]float32(a)))
}

func m512BroadcastF32x2(a [4]float32) [16]float32


// M512MaskBroadcastF32x2: Broadcast the lower 2 packed single-precision
// (32-bit) floating-point elements from 'a' to all elements of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
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
func M512MaskBroadcastF32x2(src x86.M512, k x86.Mmask16, a x86.M128) x86.M512 {
	return x86.M512(m512MaskBroadcastF32x2([16]float32(src), uint16(k), [4]float32(a)))
}

func m512MaskBroadcastF32x2(src [16]float32, k uint16, a [4]float32) [16]float32


// M512MaskzBroadcastF32x2: Broadcast the lower 2 packed single-precision
// (32-bit) floating-point elements from 'a' to all elements of 'dst' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
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
func M512MaskzBroadcastF32x2(k x86.Mmask16, a x86.M128) x86.M512 {
	return x86.M512(m512MaskzBroadcastF32x2(uint16(k), [4]float32(a)))
}

func m512MaskzBroadcastF32x2(k uint16, a [4]float32) [16]float32


// M512BroadcastF32x8: Broadcast the 8 packed single-precision (32-bit)
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
func M512BroadcastF32x8(a x86.M256) x86.M512 {
	return x86.M512(m512BroadcastF32x8([8]float32(a)))
}

func m512BroadcastF32x8(a [8]float32) [16]float32


// M512MaskBroadcastF32x8: Broadcast the 8 packed single-precision (32-bit)
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
func M512MaskBroadcastF32x8(src x86.M512, k x86.Mmask16, a x86.M256) x86.M512 {
	return x86.M512(m512MaskBroadcastF32x8([16]float32(src), uint16(k), [8]float32(a)))
}

func m512MaskBroadcastF32x8(src [16]float32, k uint16, a [8]float32) [16]float32


// M512MaskzBroadcastF32x8: Broadcast the 8 packed single-precision (32-bit)
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
func M512MaskzBroadcastF32x8(k x86.Mmask16, a x86.M256) x86.M512 {
	return x86.M512(m512MaskzBroadcastF32x8(uint16(k), [8]float32(a)))
}

func m512MaskzBroadcastF32x8(k uint16, a [8]float32) [16]float32


// M256BroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
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
func M256BroadcastF64x2(a x86.M128d) x86.M256d {
	return x86.M256d(m256BroadcastF64x2([2]float64(a)))
}

func m256BroadcastF64x2(a [2]float64) [4]float64


// M256MaskBroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
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
func M256MaskBroadcastF64x2(src x86.M256d, k x86.Mmask8, a x86.M128d) x86.M256d {
	return x86.M256d(m256MaskBroadcastF64x2([4]float64(src), uint8(k), [2]float64(a)))
}

func m256MaskBroadcastF64x2(src [4]float64, k uint8, a [2]float64) [4]float64


// M256MaskzBroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
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
func M256MaskzBroadcastF64x2(k x86.Mmask8, a x86.M128d) x86.M256d {
	return x86.M256d(m256MaskzBroadcastF64x2(uint8(k), [2]float64(a)))
}

func m256MaskzBroadcastF64x2(k uint8, a [2]float64) [4]float64


// M512BroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
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
func M512BroadcastF64x2(a x86.M128d) x86.M512d {
	return x86.M512d(m512BroadcastF64x2([2]float64(a)))
}

func m512BroadcastF64x2(a [2]float64) [8]float64


// M512MaskBroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
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
func M512MaskBroadcastF64x2(src x86.M512d, k x86.Mmask8, a x86.M128d) x86.M512d {
	return x86.M512d(m512MaskBroadcastF64x2([8]float64(src), uint8(k), [2]float64(a)))
}

func m512MaskBroadcastF64x2(src [8]float64, k uint8, a [2]float64) [8]float64


// M512MaskzBroadcastF64x2: Broadcast the 2 packed double-precision (64-bit)
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
func M512MaskzBroadcastF64x2(k x86.Mmask8, a x86.M128d) x86.M512d {
	return x86.M512d(m512MaskzBroadcastF64x2(uint8(k), [2]float64(a)))
}

func m512MaskzBroadcastF64x2(k uint8, a [2]float64) [8]float64


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
func BroadcastI32x2(a x86.M128i) x86.M128i {
	return x86.M128i(broadcastI32x2([16]byte(a)))
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
func MaskBroadcastI32x2(src x86.M128i, k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskBroadcastI32x2([16]byte(src), uint8(k), [16]byte(a)))
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
func MaskzBroadcastI32x2(k x86.Mmask8, a x86.M128i) x86.M128i {
	return x86.M128i(maskzBroadcastI32x2(uint8(k), [16]byte(a)))
}

func maskzBroadcastI32x2(k uint8, a [16]byte) [16]byte


// M256BroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from 'a' to
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
func M256BroadcastI32x2(a x86.M128i) x86.M256i {
	return x86.M256i(m256BroadcastI32x2([16]byte(a)))
}

func m256BroadcastI32x2(a [16]byte) [32]byte


// M256MaskBroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from
// 'a' to all elements of 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func M256MaskBroadcastI32x2(src x86.M256i, k x86.Mmask8, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskBroadcastI32x2([32]byte(src), uint8(k), [16]byte(a)))
}

func m256MaskBroadcastI32x2(src [32]byte, k uint8, a [16]byte) [32]byte


// M256MaskzBroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from
// 'a' to all elements of 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func M256MaskzBroadcastI32x2(k x86.Mmask8, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzBroadcastI32x2(uint8(k), [16]byte(a)))
}

func m256MaskzBroadcastI32x2(k uint8, a [16]byte) [32]byte


// M512BroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from 'a' to
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
func M512BroadcastI32x2(a x86.M128i) x86.M512i {
	return x86.M512i(m512BroadcastI32x2([16]byte(a)))
}

func m512BroadcastI32x2(a [16]byte) [64]byte


// M512MaskBroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from
// 'a' to all elements of 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
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
func M512MaskBroadcastI32x2(src x86.M512i, k x86.Mmask16, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskBroadcastI32x2([64]byte(src), uint16(k), [16]byte(a)))
}

func m512MaskBroadcastI32x2(src [64]byte, k uint16, a [16]byte) [64]byte


// M512MaskzBroadcastI32x2: Broadcast the lower 2 packed 32-bit integers from
// 'a' to all elements of 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func M512MaskzBroadcastI32x2(k x86.Mmask16, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzBroadcastI32x2(uint16(k), [16]byte(a)))
}

func m512MaskzBroadcastI32x2(k uint16, a [16]byte) [64]byte


// M512BroadcastI32x8: Broadcast the 8 packed 32-bit integers from 'a' to all
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
func M512BroadcastI32x8(a x86.M256i) x86.M512i {
	return x86.M512i(m512BroadcastI32x8([32]byte(a)))
}

func m512BroadcastI32x8(a [32]byte) [64]byte


// M512MaskBroadcastI32x8: Broadcast the 8 packed 32-bit integers from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
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
func M512MaskBroadcastI32x8(src x86.M512i, k x86.Mmask16, a x86.M256i) x86.M512i {
	return x86.M512i(m512MaskBroadcastI32x8([64]byte(src), uint16(k), [32]byte(a)))
}

func m512MaskBroadcastI32x8(src [64]byte, k uint16, a [32]byte) [64]byte


// M512MaskzBroadcastI32x8: Broadcast the 8 packed 32-bit integers from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
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
func M512MaskzBroadcastI32x8(k x86.Mmask16, a x86.M256i) x86.M512i {
	return x86.M512i(m512MaskzBroadcastI32x8(uint16(k), [32]byte(a)))
}

func m512MaskzBroadcastI32x8(k uint16, a [32]byte) [64]byte


// M256BroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to all
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
func M256BroadcastI64x2(a x86.M128i) x86.M256i {
	return x86.M256i(m256BroadcastI64x2([16]byte(a)))
}

func m256BroadcastI64x2(a [16]byte) [32]byte


// M256MaskBroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
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
func M256MaskBroadcastI64x2(src x86.M256i, k x86.Mmask8, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskBroadcastI64x2([32]byte(src), uint8(k), [16]byte(a)))
}

func m256MaskBroadcastI64x2(src [32]byte, k uint8, a [16]byte) [32]byte


// M256MaskzBroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
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
func M256MaskzBroadcastI64x2(k x86.Mmask8, a x86.M128i) x86.M256i {
	return x86.M256i(m256MaskzBroadcastI64x2(uint8(k), [16]byte(a)))
}

func m256MaskzBroadcastI64x2(k uint8, a [16]byte) [32]byte


// M512BroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to all
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
func M512BroadcastI64x2(a x86.M128i) x86.M512i {
	return x86.M512i(m512BroadcastI64x2([16]byte(a)))
}

func m512BroadcastI64x2(a [16]byte) [64]byte


// M512MaskBroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to
// all elements of 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
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
func M512MaskBroadcastI64x2(src x86.M512i, k x86.Mmask8, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskBroadcastI64x2([64]byte(src), uint8(k), [16]byte(a)))
}

func m512MaskBroadcastI64x2(src [64]byte, k uint8, a [16]byte) [64]byte


// M512MaskzBroadcastI64x2: Broadcast the 2 packed 64-bit integers from 'a' to
// all elements of 'dst' using zeromask 'k' (elements are zeroed out when the
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
func M512MaskzBroadcastI64x2(k x86.Mmask8, a x86.M128i) x86.M512i {
	return x86.M512i(m512MaskzBroadcastI64x2(uint8(k), [16]byte(a)))
}

func m512MaskzBroadcastI64x2(k uint8, a [16]byte) [64]byte


// M512CvtRoundepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M512CvtRoundepi64Pd(a x86.M512i, rounding int) x86.M512d {
	return x86.M512d(m512CvtRoundepi64Pd([64]byte(a), rounding))
}

func m512CvtRoundepi64Pd(a [64]byte, rounding int) [8]float64


// M512MaskCvtRoundepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskCvtRoundepi64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i, rounding int) x86.M512d {
	return x86.M512d(m512MaskCvtRoundepi64Pd([8]float64(src), uint8(k), [64]byte(a), rounding))
}

func m512MaskCvtRoundepi64Pd(src [8]float64, k uint8, a [64]byte, rounding int) [8]float64


// M512MaskzCvtRoundepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskzCvtRoundepi64Pd(k x86.Mmask8, a x86.M512i, rounding int) x86.M512d {
	return x86.M512d(m512MaskzCvtRoundepi64Pd(uint8(k), [64]byte(a), rounding))
}

func m512MaskzCvtRoundepi64Pd(k uint8, a [64]byte, rounding int) [8]float64


// M512CvtRoundepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M512CvtRoundepi64Ps(a x86.M512i, rounding int) x86.M256 {
	return x86.M256(m512CvtRoundepi64Ps([64]byte(a), rounding))
}

func m512CvtRoundepi64Ps(a [64]byte, rounding int) [8]float32


// M512MaskCvtRoundepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskCvtRoundepi64Ps(src x86.M256, k x86.Mmask8, a x86.M512i, rounding int) x86.M256 {
	return x86.M256(m512MaskCvtRoundepi64Ps([8]float32(src), uint8(k), [64]byte(a), rounding))
}

func m512MaskCvtRoundepi64Ps(src [8]float32, k uint8, a [64]byte, rounding int) [8]float32


// M512MaskzCvtRoundepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskzCvtRoundepi64Ps(k x86.Mmask8, a x86.M512i, rounding int) x86.M256 {
	return x86.M256(m512MaskzCvtRoundepi64Ps(uint8(k), [64]byte(a), rounding))
}

func m512MaskzCvtRoundepi64Ps(k uint8, a [64]byte, rounding int) [8]float32


// M512CvtRoundepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
// packed double-precision (64-bit) floating-point elements, and store the
// results in 'dst'.
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
func M512CvtRoundepu64Pd(a x86.M512i, rounding int) x86.M512d {
	return x86.M512d(m512CvtRoundepu64Pd([64]byte(a), rounding))
}

func m512CvtRoundepu64Pd(a [64]byte, rounding int) [8]float64


// M512MaskCvtRoundepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
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
func M512MaskCvtRoundepu64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i, rounding int) x86.M512d {
	return x86.M512d(m512MaskCvtRoundepu64Pd([8]float64(src), uint8(k), [64]byte(a), rounding))
}

func m512MaskCvtRoundepu64Pd(src [8]float64, k uint8, a [64]byte, rounding int) [8]float64


// M512MaskzCvtRoundepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
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
func M512MaskzCvtRoundepu64Pd(k x86.Mmask8, a x86.M512i, rounding int) x86.M512d {
	return x86.M512d(m512MaskzCvtRoundepu64Pd(uint8(k), [64]byte(a), rounding))
}

func m512MaskzCvtRoundepu64Pd(k uint8, a [64]byte, rounding int) [8]float64


// M512CvtRoundepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
// packed single-precision (32-bit) floating-point elements, and store the
// results in 'dst'.
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
func M512CvtRoundepu64Ps(a x86.M512i, rounding int) x86.M256 {
	return x86.M256(m512CvtRoundepu64Ps([64]byte(a), rounding))
}

func m512CvtRoundepu64Ps(a [64]byte, rounding int) [8]float32


// M512MaskCvtRoundepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
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
func M512MaskCvtRoundepu64Ps(src x86.M256, k x86.Mmask8, a x86.M512i, rounding int) x86.M256 {
	return x86.M256(m512MaskCvtRoundepu64Ps([8]float32(src), uint8(k), [64]byte(a), rounding))
}

func m512MaskCvtRoundepu64Ps(src [8]float32, k uint8, a [64]byte, rounding int) [8]float32


// M512MaskzCvtRoundepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
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
func M512MaskzCvtRoundepu64Ps(k x86.Mmask8, a x86.M512i, rounding int) x86.M256 {
	return x86.M256(m512MaskzCvtRoundepu64Ps(uint8(k), [64]byte(a), rounding))
}

func m512MaskzCvtRoundepu64Ps(k uint8, a [64]byte, rounding int) [8]float32


// M512CvtRoundpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M512CvtRoundpdEpi64(a x86.M512d, rounding int) x86.M512i {
	return x86.M512i(m512CvtRoundpdEpi64([8]float64(a), rounding))
}

func m512CvtRoundpdEpi64(a [8]float64, rounding int) [64]byte


// M512MaskCvtRoundpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers, and store the
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
//				dst[i+63:i] := Convert_FP64_To_Int64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2QQ'. Intrinsic: '_mm512_mask_cvt_roundpd_epi64'.
// Requires AVX512DQ.
func M512MaskCvtRoundpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d, rounding int) x86.M512i {
	return x86.M512i(m512MaskCvtRoundpdEpi64([64]byte(src), uint8(k), [8]float64(a), rounding))
}

func m512MaskCvtRoundpdEpi64(src [64]byte, k uint8, a [8]float64, rounding int) [64]byte


// M512MaskzCvtRoundpdEpi64: Convert packed double-precision (64-bit)
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
func M512MaskzCvtRoundpdEpi64(k x86.Mmask8, a x86.M512d, rounding int) x86.M512i {
	return x86.M512i(m512MaskzCvtRoundpdEpi64(uint8(k), [8]float64(a), rounding))
}

func m512MaskzCvtRoundpdEpi64(k uint8, a [8]float64, rounding int) [64]byte


// M512CvtRoundpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M512CvtRoundpdEpu64(a x86.M512d, rounding int) x86.M512i {
	return x86.M512i(m512CvtRoundpdEpu64([8]float64(a), rounding))
}

func m512CvtRoundpdEpu64(a [8]float64, rounding int) [64]byte


// M512MaskCvtRoundpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers, and store
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
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
func M512MaskCvtRoundpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d, rounding int) x86.M512i {
	return x86.M512i(m512MaskCvtRoundpdEpu64([64]byte(src), uint8(k), [8]float64(a), rounding))
}

func m512MaskCvtRoundpdEpu64(src [64]byte, k uint8, a [8]float64, rounding int) [64]byte


// M512MaskzCvtRoundpdEpu64: Convert packed double-precision (64-bit)
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
func M512MaskzCvtRoundpdEpu64(k x86.Mmask8, a x86.M512d, rounding int) x86.M512i {
	return x86.M512i(m512MaskzCvtRoundpdEpu64(uint8(k), [8]float64(a), rounding))
}

func m512MaskzCvtRoundpdEpu64(k uint8, a [8]float64, rounding int) [64]byte


// M512CvtRoundpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M512CvtRoundpsEpi64(a x86.M256, rounding int) x86.M512i {
	return x86.M512i(m512CvtRoundpsEpi64([8]float32(a), rounding))
}

func m512CvtRoundpsEpi64(a [8]float32, rounding int) [64]byte


// M512MaskCvtRoundpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
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
func M512MaskCvtRoundpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256, rounding int) x86.M512i {
	return x86.M512i(m512MaskCvtRoundpsEpi64([64]byte(src), uint8(k), [8]float32(a), rounding))
}

func m512MaskCvtRoundpsEpi64(src [64]byte, k uint8, a [8]float32, rounding int) [64]byte


// M512MaskzCvtRoundpsEpi64: Convert packed single-precision (32-bit)
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
func M512MaskzCvtRoundpsEpi64(k x86.Mmask8, a x86.M256, rounding int) x86.M512i {
	return x86.M512i(m512MaskzCvtRoundpsEpi64(uint8(k), [8]float32(a), rounding))
}

func m512MaskzCvtRoundpsEpi64(k uint8, a [8]float32, rounding int) [64]byte


// M512CvtRoundpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M512CvtRoundpsEpu64(a x86.M256, rounding int) x86.M512i {
	return x86.M512i(m512CvtRoundpsEpu64([8]float32(a), rounding))
}

func m512CvtRoundpsEpu64(a [8]float32, rounding int) [64]byte


// M512MaskCvtRoundpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers, and store
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
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
func M512MaskCvtRoundpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256, rounding int) x86.M512i {
	return x86.M512i(m512MaskCvtRoundpsEpu64([64]byte(src), uint8(k), [8]float32(a), rounding))
}

func m512MaskCvtRoundpsEpu64(src [64]byte, k uint8, a [8]float32, rounding int) [64]byte


// M512MaskzCvtRoundpsEpu64: Convert packed single-precision (32-bit)
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
func M512MaskzCvtRoundpsEpu64(k x86.Mmask8, a x86.M256, rounding int) x86.M512i {
	return x86.M512i(m512MaskzCvtRoundpsEpu64(uint8(k), [8]float32(a), rounding))
}

func m512MaskzCvtRoundpsEpu64(k uint8, a [8]float32, rounding int) [64]byte


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
func Cvtepi64Pd(a x86.M128i) x86.M128d {
	return x86.M128d(cvtepi64Pd([16]byte(a)))
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
func MaskCvtepi64Pd(src x86.M128d, k x86.Mmask8, a x86.M128i) x86.M128d {
	return x86.M128d(maskCvtepi64Pd([2]float64(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepi64Pd(k x86.Mmask8, a x86.M128i) x86.M128d {
	return x86.M128d(maskzCvtepi64Pd(uint8(k), [16]byte(a)))
}

func maskzCvtepi64Pd(k uint8, a [16]byte) [2]float64


// M256Cvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M256Cvtepi64Pd(a x86.M256i) x86.M256d {
	return x86.M256d(m256Cvtepi64Pd([32]byte(a)))
}

func m256Cvtepi64Pd(a [32]byte) [4]float64


// M256MaskCvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M256MaskCvtepi64Pd(src x86.M256d, k x86.Mmask8, a x86.M256i) x86.M256d {
	return x86.M256d(m256MaskCvtepi64Pd([4]float64(src), uint8(k), [32]byte(a)))
}

func m256MaskCvtepi64Pd(src [4]float64, k uint8, a [32]byte) [4]float64


// M256MaskzCvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M256MaskzCvtepi64Pd(k x86.Mmask8, a x86.M256i) x86.M256d {
	return x86.M256d(m256MaskzCvtepi64Pd(uint8(k), [32]byte(a)))
}

func m256MaskzCvtepi64Pd(k uint8, a [32]byte) [4]float64


// M512Cvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M512Cvtepi64Pd(a x86.M512i) x86.M512d {
	return x86.M512d(m512Cvtepi64Pd([64]byte(a)))
}

func m512Cvtepi64Pd(a [64]byte) [8]float64


// M512MaskCvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskCvtepi64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i) x86.M512d {
	return x86.M512d(m512MaskCvtepi64Pd([8]float64(src), uint8(k), [64]byte(a)))
}

func m512MaskCvtepi64Pd(src [8]float64, k uint8, a [64]byte) [8]float64


// M512MaskzCvtepi64Pd: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskzCvtepi64Pd(k x86.Mmask8, a x86.M512i) x86.M512d {
	return x86.M512d(m512MaskzCvtepi64Pd(uint8(k), [64]byte(a)))
}

func m512MaskzCvtepi64Pd(k uint8, a [64]byte) [8]float64


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
func Cvtepi64Ps(a x86.M128i) x86.M128 {
	return x86.M128(cvtepi64Ps([16]byte(a)))
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
func MaskCvtepi64Ps(src x86.M128, k x86.Mmask8, a x86.M128i) x86.M128 {
	return x86.M128(maskCvtepi64Ps([4]float32(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepi64Ps(k x86.Mmask8, a x86.M128i) x86.M128 {
	return x86.M128(maskzCvtepi64Ps(uint8(k), [16]byte(a)))
}

func maskzCvtepi64Ps(k uint8, a [16]byte) [4]float32


// M256Cvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M256Cvtepi64Ps(a x86.M256i) x86.M128 {
	return x86.M128(m256Cvtepi64Ps([32]byte(a)))
}

func m256Cvtepi64Ps(a [32]byte) [4]float32


// M256MaskCvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M256MaskCvtepi64Ps(src x86.M128, k x86.Mmask8, a x86.M256i) x86.M128 {
	return x86.M128(m256MaskCvtepi64Ps([4]float32(src), uint8(k), [32]byte(a)))
}

func m256MaskCvtepi64Ps(src [4]float32, k uint8, a [32]byte) [4]float32


// M256MaskzCvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M256MaskzCvtepi64Ps(k x86.Mmask8, a x86.M256i) x86.M128 {
	return x86.M128(m256MaskzCvtepi64Ps(uint8(k), [32]byte(a)))
}

func m256MaskzCvtepi64Ps(k uint8, a [32]byte) [4]float32


// M512Cvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M512Cvtepi64Ps(a x86.M512i) x86.M256 {
	return x86.M256(m512Cvtepi64Ps([64]byte(a)))
}

func m512Cvtepi64Ps(a [64]byte) [8]float32


// M512MaskCvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskCvtepi64Ps(src x86.M256, k x86.Mmask8, a x86.M512i) x86.M256 {
	return x86.M256(m512MaskCvtepi64Ps([8]float32(src), uint8(k), [64]byte(a)))
}

func m512MaskCvtepi64Ps(src [8]float32, k uint8, a [64]byte) [8]float32


// M512MaskzCvtepi64Ps: Convert packed 64-bit integers in 'a' to packed
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
func M512MaskzCvtepi64Ps(k x86.Mmask8, a x86.M512i) x86.M256 {
	return x86.M256(m512MaskzCvtepi64Ps(uint8(k), [64]byte(a)))
}

func m512MaskzCvtepi64Ps(k uint8, a [64]byte) [8]float32


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
func Cvtepu64Pd(a x86.M128i) x86.M128d {
	return x86.M128d(cvtepu64Pd([16]byte(a)))
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
func MaskCvtepu64Pd(src x86.M128d, k x86.Mmask8, a x86.M128i) x86.M128d {
	return x86.M128d(maskCvtepu64Pd([2]float64(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepu64Pd(k x86.Mmask8, a x86.M128i) x86.M128d {
	return x86.M128d(maskzCvtepu64Pd(uint8(k), [16]byte(a)))
}

func maskzCvtepu64Pd(k uint8, a [16]byte) [2]float64


// M256Cvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M256Cvtepu64Pd(a x86.M256i) x86.M256d {
	return x86.M256d(m256Cvtepu64Pd([32]byte(a)))
}

func m256Cvtepu64Pd(a [32]byte) [4]float64


// M256MaskCvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M256MaskCvtepu64Pd(src x86.M256d, k x86.Mmask8, a x86.M256i) x86.M256d {
	return x86.M256d(m256MaskCvtepu64Pd([4]float64(src), uint8(k), [32]byte(a)))
}

func m256MaskCvtepu64Pd(src [4]float64, k uint8, a [32]byte) [4]float64


// M256MaskzCvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
// packed double-precision (64-bit) floating-point elements, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
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
func M256MaskzCvtepu64Pd(k x86.Mmask8, a x86.M256i) x86.M256d {
	return x86.M256d(m256MaskzCvtepu64Pd(uint8(k), [32]byte(a)))
}

func m256MaskzCvtepu64Pd(k uint8, a [32]byte) [4]float64


// M512Cvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M512Cvtepu64Pd(a x86.M512i) x86.M512d {
	return x86.M512d(m512Cvtepu64Pd([64]byte(a)))
}

func m512Cvtepu64Pd(a [64]byte) [8]float64


// M512MaskCvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M512MaskCvtepu64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i) x86.M512d {
	return x86.M512d(m512MaskCvtepu64Pd([8]float64(src), uint8(k), [64]byte(a)))
}

func m512MaskCvtepu64Pd(src [8]float64, k uint8, a [64]byte) [8]float64


// M512MaskzCvtepu64Pd: Convert packed unsigned 64-bit integers in 'a' to
// packed double-precision (64-bit) floating-point elements, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
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
func M512MaskzCvtepu64Pd(k x86.Mmask8, a x86.M512i) x86.M512d {
	return x86.M512d(m512MaskzCvtepu64Pd(uint8(k), [64]byte(a)))
}

func m512MaskzCvtepu64Pd(k uint8, a [64]byte) [8]float64


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
func Cvtepu64Ps(a x86.M128i) x86.M128 {
	return x86.M128(cvtepu64Ps([16]byte(a)))
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
func MaskCvtepu64Ps(src x86.M128, k x86.Mmask8, a x86.M128i) x86.M128 {
	return x86.M128(maskCvtepu64Ps([4]float32(src), uint8(k), [16]byte(a)))
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
func MaskzCvtepu64Ps(k x86.Mmask8, a x86.M128i) x86.M128 {
	return x86.M128(maskzCvtepu64Ps(uint8(k), [16]byte(a)))
}

func maskzCvtepu64Ps(k uint8, a [16]byte) [4]float32


// M256Cvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M256Cvtepu64Ps(a x86.M256i) x86.M128 {
	return x86.M128(m256Cvtepu64Ps([32]byte(a)))
}

func m256Cvtepu64Ps(a [32]byte) [4]float32


// M256MaskCvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M256MaskCvtepu64Ps(src x86.M128, k x86.Mmask8, a x86.M256i) x86.M128 {
	return x86.M128(m256MaskCvtepu64Ps([4]float32(src), uint8(k), [32]byte(a)))
}

func m256MaskCvtepu64Ps(src [4]float32, k uint8, a [32]byte) [4]float32


// M256MaskzCvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
// packed single-precision (32-bit) floating-point elements, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
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
func M256MaskzCvtepu64Ps(k x86.Mmask8, a x86.M256i) x86.M128 {
	return x86.M128(m256MaskzCvtepu64Ps(uint8(k), [32]byte(a)))
}

func m256MaskzCvtepu64Ps(k uint8, a [32]byte) [4]float32


// M512Cvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M512Cvtepu64Ps(a x86.M512i) x86.M256 {
	return x86.M256(m512Cvtepu64Ps([64]byte(a)))
}

func m512Cvtepu64Ps(a [64]byte) [8]float32


// M512MaskCvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to packed
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
func M512MaskCvtepu64Ps(src x86.M256, k x86.Mmask8, a x86.M512i) x86.M256 {
	return x86.M256(m512MaskCvtepu64Ps([8]float32(src), uint8(k), [64]byte(a)))
}

func m512MaskCvtepu64Ps(src [8]float32, k uint8, a [64]byte) [8]float32


// M512MaskzCvtepu64Ps: Convert packed unsigned 64-bit integers in 'a' to
// packed single-precision (32-bit) floating-point elements, and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
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
func M512MaskzCvtepu64Ps(k x86.Mmask8, a x86.M512i) x86.M256 {
	return x86.M256(m512MaskzCvtepu64Ps(uint8(k), [64]byte(a)))
}

func m512MaskzCvtepu64Ps(k uint8, a [64]byte) [8]float32


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
func CvtpdEpi64(a x86.M128d) x86.M128i {
	return x86.M128i(cvtpdEpi64([2]float64(a)))
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
func MaskCvtpdEpi64(src x86.M128i, k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskCvtpdEpi64([16]byte(src), uint8(k), [2]float64(a)))
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
func MaskzCvtpdEpi64(k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskzCvtpdEpi64(uint8(k), [2]float64(a)))
}

func maskzCvtpdEpi64(k uint8, a [2]float64) [16]byte


// M256CvtpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M256CvtpdEpi64(a x86.M256d) x86.M256i {
	return x86.M256i(m256CvtpdEpi64([4]float64(a)))
}

func m256CvtpdEpi64(a [4]float64) [32]byte


// M256MaskCvtpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M256MaskCvtpdEpi64(src x86.M256i, k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskCvtpdEpi64([32]byte(src), uint8(k), [4]float64(a)))
}

func m256MaskCvtpdEpi64(src [32]byte, k uint8, a [4]float64) [32]byte


// M256MaskzCvtpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M256MaskzCvtpdEpi64(k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskzCvtpdEpi64(uint8(k), [4]float64(a)))
}

func m256MaskzCvtpdEpi64(k uint8, a [4]float64) [32]byte


// M512CvtpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M512CvtpdEpi64(a x86.M512d) x86.M512i {
	return x86.M512i(m512CvtpdEpi64([8]float64(a)))
}

func m512CvtpdEpi64(a [8]float64) [64]byte


// M512MaskCvtpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M512MaskCvtpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskCvtpdEpi64([64]byte(src), uint8(k), [8]float64(a)))
}

func m512MaskCvtpdEpi64(src [64]byte, k uint8, a [8]float64) [64]byte


// M512MaskzCvtpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M512MaskzCvtpdEpi64(k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskzCvtpdEpi64(uint8(k), [8]float64(a)))
}

func m512MaskzCvtpdEpi64(k uint8, a [8]float64) [64]byte


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
func CvtpdEpu64(a x86.M128d) x86.M128i {
	return x86.M128i(cvtpdEpu64([2]float64(a)))
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
func MaskCvtpdEpu64(src x86.M128i, k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskCvtpdEpu64([16]byte(src), uint8(k), [2]float64(a)))
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
func MaskzCvtpdEpu64(k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskzCvtpdEpu64(uint8(k), [2]float64(a)))
}

func maskzCvtpdEpu64(k uint8, a [2]float64) [16]byte


// M256CvtpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M256CvtpdEpu64(a x86.M256d) x86.M256i {
	return x86.M256i(m256CvtpdEpu64([4]float64(a)))
}

func m256CvtpdEpu64(a [4]float64) [32]byte


// M256MaskCvtpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M256MaskCvtpdEpu64(src x86.M256i, k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskCvtpdEpu64([32]byte(src), uint8(k), [4]float64(a)))
}

func m256MaskCvtpdEpu64(src [32]byte, k uint8, a [4]float64) [32]byte


// M256MaskzCvtpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M256MaskzCvtpdEpu64(k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskzCvtpdEpu64(uint8(k), [4]float64(a)))
}

func m256MaskzCvtpdEpu64(k uint8, a [4]float64) [32]byte


// M512CvtpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M512CvtpdEpu64(a x86.M512d) x86.M512i {
	return x86.M512i(m512CvtpdEpu64([8]float64(a)))
}

func m512CvtpdEpu64(a [8]float64) [64]byte


// M512MaskCvtpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M512MaskCvtpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskCvtpdEpu64([64]byte(src), uint8(k), [8]float64(a)))
}

func m512MaskCvtpdEpu64(src [64]byte, k uint8, a [8]float64) [64]byte


// M512MaskzCvtpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M512MaskzCvtpdEpu64(k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskzCvtpdEpu64(uint8(k), [8]float64(a)))
}

func m512MaskzCvtpdEpu64(k uint8, a [8]float64) [64]byte


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
func CvtpsEpi64(a x86.M128) x86.M128i {
	return x86.M128i(cvtpsEpi64([4]float32(a)))
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
func MaskCvtpsEpi64(src x86.M128i, k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskCvtpsEpi64([16]byte(src), uint8(k), [4]float32(a)))
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
func MaskzCvtpsEpi64(k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskzCvtpsEpi64(uint8(k), [4]float32(a)))
}

func maskzCvtpsEpi64(k uint8, a [4]float32) [16]byte


// M256CvtpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M256CvtpsEpi64(a x86.M128) x86.M256i {
	return x86.M256i(m256CvtpsEpi64([4]float32(a)))
}

func m256CvtpsEpi64(a [4]float32) [32]byte


// M256MaskCvtpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M256MaskCvtpsEpi64(src x86.M256i, k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskCvtpsEpi64([32]byte(src), uint8(k), [4]float32(a)))
}

func m256MaskCvtpsEpi64(src [32]byte, k uint8, a [4]float32) [32]byte


// M256MaskzCvtpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M256MaskzCvtpsEpi64(k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskzCvtpsEpi64(uint8(k), [4]float32(a)))
}

func m256MaskzCvtpsEpi64(k uint8, a [4]float32) [32]byte


// M512CvtpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M512CvtpsEpi64(a x86.M256) x86.M512i {
	return x86.M512i(m512CvtpsEpi64([8]float32(a)))
}

func m512CvtpsEpi64(a [8]float32) [64]byte


// M512MaskCvtpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M512MaskCvtpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskCvtpsEpi64([64]byte(src), uint8(k), [8]float32(a)))
}

func m512MaskCvtpsEpi64(src [64]byte, k uint8, a [8]float32) [64]byte


// M512MaskzCvtpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M512MaskzCvtpsEpi64(k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskzCvtpsEpi64(uint8(k), [8]float32(a)))
}

func m512MaskzCvtpsEpi64(k uint8, a [8]float32) [64]byte


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
func CvtpsEpu64(a x86.M128) x86.M128i {
	return x86.M128i(cvtpsEpu64([4]float32(a)))
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
func MaskCvtpsEpu64(src x86.M128i, k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskCvtpsEpu64([16]byte(src), uint8(k), [4]float32(a)))
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
func MaskzCvtpsEpu64(k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskzCvtpsEpu64(uint8(k), [4]float32(a)))
}

func maskzCvtpsEpu64(k uint8, a [4]float32) [16]byte


// M256CvtpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M256CvtpsEpu64(a x86.M128) x86.M256i {
	return x86.M256i(m256CvtpsEpu64([4]float32(a)))
}

func m256CvtpsEpu64(a [4]float32) [32]byte


// M256MaskCvtpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M256MaskCvtpsEpu64(src x86.M256i, k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskCvtpsEpu64([32]byte(src), uint8(k), [4]float32(a)))
}

func m256MaskCvtpsEpu64(src [32]byte, k uint8, a [4]float32) [32]byte


// M256MaskzCvtpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M256MaskzCvtpsEpu64(k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskzCvtpsEpu64(uint8(k), [4]float32(a)))
}

func m256MaskzCvtpsEpu64(k uint8, a [4]float32) [32]byte


// M512CvtpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M512CvtpsEpu64(a x86.M256) x86.M512i {
	return x86.M512i(m512CvtpsEpu64([8]float32(a)))
}

func m512CvtpsEpu64(a [8]float32) [64]byte


// M512MaskCvtpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M512MaskCvtpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskCvtpsEpu64([64]byte(src), uint8(k), [8]float32(a)))
}

func m512MaskCvtpsEpu64(src [64]byte, k uint8, a [8]float32) [64]byte


// M512MaskzCvtpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M512MaskzCvtpsEpu64(k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskzCvtpsEpu64(uint8(k), [8]float32(a)))
}

func m512MaskzCvtpsEpu64(k uint8, a [8]float32) [64]byte


// M512CvttRoundpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae' to suppress
// all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_Int64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2QQ'. Intrinsic: '_mm512_cvtt_roundpd_epi64'.
// Requires AVX512DQ.
func M512CvttRoundpdEpi64(a x86.M512d, sae int) x86.M512i {
	return x86.M512i(m512CvttRoundpdEpi64([8]float64(a), sae))
}

func m512CvttRoundpdEpi64(a [8]float64, sae int) [64]byte


// M512MaskCvttRoundpdEpi64: Convert packed double-precision (64-bit)
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
func M512MaskCvttRoundpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d, sae int) x86.M512i {
	return x86.M512i(m512MaskCvttRoundpdEpi64([64]byte(src), uint8(k), [8]float64(a), sae))
}

func m512MaskCvttRoundpdEpi64(src [64]byte, k uint8, a [8]float64, sae int) [64]byte


// M512MaskzCvttRoundpdEpi64: Convert packed double-precision (64-bit)
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
func M512MaskzCvttRoundpdEpi64(k x86.Mmask8, a x86.M512d, sae int) x86.M512i {
	return x86.M512i(m512MaskzCvttRoundpdEpi64(uint8(k), [8]float64(a), sae))
}

func m512MaskzCvttRoundpdEpi64(k uint8, a [8]float64, sae int) [64]byte


// M512CvttRoundpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae'
// to suppress all exceptions. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := Convert_FP64_To_UnsignedInt64_Truncate(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTTPD2UQQ'. Intrinsic: '_mm512_cvtt_roundpd_epu64'.
// Requires AVX512DQ.
func M512CvttRoundpdEpu64(a x86.M512d, sae int) x86.M512i {
	return x86.M512i(m512CvttRoundpdEpu64([8]float64(a), sae))
}

func m512CvttRoundpdEpu64(a [8]float64, sae int) [64]byte


// M512MaskCvttRoundpdEpu64: Convert packed double-precision (64-bit)
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
func M512MaskCvttRoundpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d, sae int) x86.M512i {
	return x86.M512i(m512MaskCvttRoundpdEpu64([64]byte(src), uint8(k), [8]float64(a), sae))
}

func m512MaskCvttRoundpdEpu64(src [64]byte, k uint8, a [8]float64, sae int) [64]byte


// M512MaskzCvttRoundpdEpu64: Convert packed double-precision (64-bit)
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
func M512MaskzCvttRoundpdEpu64(k x86.Mmask8, a x86.M512d, sae int) x86.M512i {
	return x86.M512i(m512MaskzCvttRoundpdEpu64(uint8(k), [8]float64(a), sae))
}

func m512MaskzCvttRoundpdEpu64(k uint8, a [8]float64, sae int) [64]byte


// M512CvttRoundpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae' to suppress
// all exceptions. 
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
func M512CvttRoundpsEpi64(a x86.M256, sae int) x86.M512i {
	return x86.M512i(m512CvttRoundpsEpi64([8]float32(a), sae))
}

func m512CvttRoundpsEpi64(a [8]float32, sae int) [64]byte


// M512MaskCvttRoundpsEpi64: Convert packed single-precision (32-bit)
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
func M512MaskCvttRoundpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256, sae int) x86.M512i {
	return x86.M512i(m512MaskCvttRoundpsEpi64([64]byte(src), uint8(k), [8]float32(a), sae))
}

func m512MaskCvttRoundpsEpi64(src [64]byte, k uint8, a [8]float32, sae int) [64]byte


// M512MaskzCvttRoundpsEpi64: Convert packed single-precision (32-bit)
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
func M512MaskzCvttRoundpsEpi64(k x86.Mmask8, a x86.M256, sae int) x86.M512i {
	return x86.M512i(m512MaskzCvttRoundpsEpi64(uint8(k), [8]float32(a), sae))
}

func m512MaskzCvttRoundpsEpi64(k uint8, a [8]float32, sae int) [64]byte


// M512CvttRoundpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst'. Pass __MM_FROUND_NO_EXC to 'sae'
// to suppress all exceptions. 
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
func M512CvttRoundpsEpu64(a x86.M256, sae int) x86.M512i {
	return x86.M512i(m512CvttRoundpsEpu64([8]float32(a), sae))
}

func m512CvttRoundpsEpu64(a [8]float32, sae int) [64]byte


// M512MaskCvttRoundpsEpu64: Convert packed single-precision (32-bit)
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
func M512MaskCvttRoundpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256, sae int) x86.M512i {
	return x86.M512i(m512MaskCvttRoundpsEpu64([64]byte(src), uint8(k), [8]float32(a), sae))
}

func m512MaskCvttRoundpsEpu64(src [64]byte, k uint8, a [8]float32, sae int) [64]byte


// M512MaskzCvttRoundpsEpu64: Convert packed single-precision (32-bit)
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
func M512MaskzCvttRoundpsEpu64(k x86.Mmask8, a x86.M256, sae int) x86.M512i {
	return x86.M512i(m512MaskzCvttRoundpsEpu64(uint8(k), [8]float32(a), sae))
}

func m512MaskzCvttRoundpsEpu64(k uint8, a [8]float32, sae int) [64]byte


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
func CvttpdEpi64(a x86.M128d) x86.M128i {
	return x86.M128i(cvttpdEpi64([2]float64(a)))
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
func MaskCvttpdEpi64(src x86.M128i, k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskCvttpdEpi64([16]byte(src), uint8(k), [2]float64(a)))
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
func MaskzCvttpdEpi64(k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskzCvttpdEpi64(uint8(k), [2]float64(a)))
}

func maskzCvttpdEpi64(k uint8, a [2]float64) [16]byte


// M256CvttpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M256CvttpdEpi64(a x86.M256d) x86.M256i {
	return x86.M256i(m256CvttpdEpi64([4]float64(a)))
}

func m256CvttpdEpi64(a [4]float64) [32]byte


// M256MaskCvttpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M256MaskCvttpdEpi64(src x86.M256i, k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskCvttpdEpi64([32]byte(src), uint8(k), [4]float64(a)))
}

func m256MaskCvttpdEpi64(src [32]byte, k uint8, a [4]float64) [32]byte


// M256MaskzCvttpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func M256MaskzCvttpdEpi64(k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskzCvttpdEpi64(uint8(k), [4]float64(a)))
}

func m256MaskzCvttpdEpi64(k uint8, a [4]float64) [32]byte


// M512CvttpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M512CvttpdEpi64(a x86.M512d) x86.M512i {
	return x86.M512i(m512CvttpdEpi64([8]float64(a)))
}

func m512CvttpdEpi64(a [8]float64) [64]byte


// M512MaskCvttpdEpi64: Convert packed double-precision (64-bit) floating-point
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
func M512MaskCvttpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskCvttpdEpi64([64]byte(src), uint8(k), [8]float64(a)))
}

func m512MaskCvttpdEpi64(src [64]byte, k uint8, a [8]float64) [64]byte


// M512MaskzCvttpdEpi64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func M512MaskzCvttpdEpi64(k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskzCvttpdEpi64(uint8(k), [8]float64(a)))
}

func m512MaskzCvttpdEpi64(k uint8, a [8]float64) [64]byte


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
func CvttpdEpu64(a x86.M128d) x86.M128i {
	return x86.M128i(cvttpdEpu64([2]float64(a)))
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
func MaskCvttpdEpu64(src x86.M128i, k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskCvttpdEpu64([16]byte(src), uint8(k), [2]float64(a)))
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
func MaskzCvttpdEpu64(k x86.Mmask8, a x86.M128d) x86.M128i {
	return x86.M128i(maskzCvttpdEpu64(uint8(k), [2]float64(a)))
}

func maskzCvttpdEpu64(k uint8, a [2]float64) [16]byte


// M256CvttpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M256CvttpdEpu64(a x86.M256d) x86.M256i {
	return x86.M256i(m256CvttpdEpu64([4]float64(a)))
}

func m256CvttpdEpu64(a [4]float64) [32]byte


// M256MaskCvttpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M256MaskCvttpdEpu64(src x86.M256i, k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskCvttpdEpu64([32]byte(src), uint8(k), [4]float64(a)))
}

func m256MaskCvttpdEpu64(src [32]byte, k uint8, a [4]float64) [32]byte


// M256MaskzCvttpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func M256MaskzCvttpdEpu64(k x86.Mmask8, a x86.M256d) x86.M256i {
	return x86.M256i(m256MaskzCvttpdEpu64(uint8(k), [4]float64(a)))
}

func m256MaskzCvttpdEpu64(k uint8, a [4]float64) [32]byte


// M512CvttpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M512CvttpdEpu64(a x86.M512d) x86.M512i {
	return x86.M512i(m512CvttpdEpu64([8]float64(a)))
}

func m512CvttpdEpu64(a [8]float64) [64]byte


// M512MaskCvttpdEpu64: Convert packed double-precision (64-bit) floating-point
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
func M512MaskCvttpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskCvttpdEpu64([64]byte(src), uint8(k), [8]float64(a)))
}

func m512MaskCvttpdEpu64(src [64]byte, k uint8, a [8]float64) [64]byte


// M512MaskzCvttpdEpu64: Convert packed double-precision (64-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func M512MaskzCvttpdEpu64(k x86.Mmask8, a x86.M512d) x86.M512i {
	return x86.M512i(m512MaskzCvttpdEpu64(uint8(k), [8]float64(a)))
}

func m512MaskzCvttpdEpu64(k uint8, a [8]float64) [64]byte


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
func CvttpsEpi64(a x86.M128) x86.M128i {
	return x86.M128i(cvttpsEpi64([4]float32(a)))
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
func MaskCvttpsEpi64(src x86.M128i, k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskCvttpsEpi64([16]byte(src), uint8(k), [4]float32(a)))
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
func MaskzCvttpsEpi64(k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskzCvttpsEpi64(uint8(k), [4]float32(a)))
}

func maskzCvttpsEpi64(k uint8, a [4]float32) [16]byte


// M256CvttpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M256CvttpsEpi64(a x86.M128) x86.M256i {
	return x86.M256i(m256CvttpsEpi64([4]float32(a)))
}

func m256CvttpsEpi64(a [4]float32) [32]byte


// M256MaskCvttpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M256MaskCvttpsEpi64(src x86.M256i, k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskCvttpsEpi64([32]byte(src), uint8(k), [4]float32(a)))
}

func m256MaskCvttpsEpi64(src [32]byte, k uint8, a [4]float32) [32]byte


// M256MaskzCvttpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func M256MaskzCvttpsEpi64(k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskzCvttpsEpi64(uint8(k), [4]float32(a)))
}

func m256MaskzCvttpsEpi64(k uint8, a [4]float32) [32]byte


// M512CvttpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M512CvttpsEpi64(a x86.M256) x86.M512i {
	return x86.M512i(m512CvttpsEpi64([8]float32(a)))
}

func m512CvttpsEpi64(a [8]float32) [64]byte


// M512MaskCvttpsEpi64: Convert packed single-precision (32-bit) floating-point
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
func M512MaskCvttpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskCvttpsEpi64([64]byte(src), uint8(k), [8]float32(a)))
}

func m512MaskCvttpsEpi64(src [64]byte, k uint8, a [8]float32) [64]byte


// M512MaskzCvttpsEpi64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed 64-bit integers with truncation,
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). 
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
func M512MaskzCvttpsEpi64(k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskzCvttpsEpi64(uint8(k), [8]float32(a)))
}

func m512MaskzCvttpsEpi64(k uint8, a [8]float32) [64]byte


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
func CvttpsEpu64(a x86.M128) x86.M128i {
	return x86.M128i(cvttpsEpu64([4]float32(a)))
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
func MaskCvttpsEpu64(src x86.M128i, k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskCvttpsEpu64([16]byte(src), uint8(k), [4]float32(a)))
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
func MaskzCvttpsEpu64(k x86.Mmask8, a x86.M128) x86.M128i {
	return x86.M128i(maskzCvttpsEpu64(uint8(k), [4]float32(a)))
}

func maskzCvttpsEpu64(k uint8, a [4]float32) [16]byte


// M256CvttpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M256CvttpsEpu64(a x86.M128) x86.M256i {
	return x86.M256i(m256CvttpsEpu64([4]float32(a)))
}

func m256CvttpsEpu64(a [4]float32) [32]byte


// M256MaskCvttpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M256MaskCvttpsEpu64(src x86.M256i, k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskCvttpsEpu64([32]byte(src), uint8(k), [4]float32(a)))
}

func m256MaskCvttpsEpu64(src [32]byte, k uint8, a [4]float32) [32]byte


// M256MaskzCvttpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func M256MaskzCvttpsEpu64(k x86.Mmask8, a x86.M128) x86.M256i {
	return x86.M256i(m256MaskzCvttpsEpu64(uint8(k), [4]float32(a)))
}

func m256MaskzCvttpsEpu64(k uint8, a [4]float32) [32]byte


// M512CvttpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M512CvttpsEpu64(a x86.M256) x86.M512i {
	return x86.M512i(m512CvttpsEpu64([8]float32(a)))
}

func m512CvttpsEpu64(a [8]float32) [64]byte


// M512MaskCvttpsEpu64: Convert packed single-precision (32-bit) floating-point
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
func M512MaskCvttpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskCvttpsEpu64([64]byte(src), uint8(k), [8]float32(a)))
}

func m512MaskCvttpsEpu64(src [64]byte, k uint8, a [8]float32) [64]byte


// M512MaskzCvttpsEpu64: Convert packed single-precision (32-bit)
// floating-point elements in 'a' to packed unsigned 64-bit integers with
// truncation, and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func M512MaskzCvttpsEpu64(k x86.Mmask8, a x86.M256) x86.M512i {
	return x86.M512i(m512MaskzCvttpsEpu64(uint8(k), [8]float32(a)))
}

func m512MaskzCvttpsEpu64(k uint8, a [8]float32) [64]byte


// M512Extractf32x8Ps: Extract 256 bits (composed of 8 packed single-precision
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
func M512Extractf32x8Ps(a x86.M512, imm8 int) x86.M256 {
	return x86.M256(m512Extractf32x8Ps([16]float32(a), imm8))
}

func m512Extractf32x8Ps(a [16]float32, imm8 int) [8]float32


// M512MaskExtractf32x8Ps: Extract 256 bits (composed of 8 packed
// single-precision (32-bit) floating-point elements) from 'a', selected with
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func M512MaskExtractf32x8Ps(src x86.M256, k x86.Mmask8, a x86.M512, imm8 int) x86.M256 {
	return x86.M256(m512MaskExtractf32x8Ps([8]float32(src), uint8(k), [16]float32(a), imm8))
}

func m512MaskExtractf32x8Ps(src [8]float32, k uint8, a [16]float32, imm8 int) [8]float32


// M512MaskzExtractf32x8Ps: Extract 256 bits (composed of 8 packed
// single-precision (32-bit) floating-point elements) from 'a', selected with
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func M512MaskzExtractf32x8Ps(k x86.Mmask8, a x86.M512, imm8 int) x86.M256 {
	return x86.M256(m512MaskzExtractf32x8Ps(uint8(k), [16]float32(a), imm8))
}

func m512MaskzExtractf32x8Ps(k uint8, a [16]float32, imm8 int) [8]float32


// M256Extractf64x2Pd: Extract 128 bits (composed of 2 packed double-precision
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
func M256Extractf64x2Pd(a x86.M256d, imm8 int) x86.M128d {
	return x86.M128d(m256Extractf64x2Pd([4]float64(a), imm8))
}

func m256Extractf64x2Pd(a [4]float64, imm8 int) [2]float64


// M256MaskExtractf64x2Pd: Extract 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'a', selected with
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func M256MaskExtractf64x2Pd(src x86.M128d, k x86.Mmask8, a x86.M256d, imm8 int) x86.M128d {
	return x86.M128d(m256MaskExtractf64x2Pd([2]float64(src), uint8(k), [4]float64(a), imm8))
}

func m256MaskExtractf64x2Pd(src [2]float64, k uint8, a [4]float64, imm8 int) [2]float64


// M256MaskzExtractf64x2Pd: Extract 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'a', selected with
// 'imm8', and store the results in 'dst' using zeromask 'k' (elements are
// zeroed out when the corresponding mask bit is not set). 
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
func M256MaskzExtractf64x2Pd(k x86.Mmask8, a x86.M256d, imm8 int) x86.M128d {
	return x86.M128d(m256MaskzExtractf64x2Pd(uint8(k), [4]float64(a), imm8))
}

func m256MaskzExtractf64x2Pd(k uint8, a [4]float64, imm8 int) [2]float64


// M512Extractf64x2Pd: Extract 128 bits (composed of 2 packed double-precision
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
func M512Extractf64x2Pd(a x86.M512d, imm8 int) x86.M128d {
	return x86.M128d(m512Extractf64x2Pd([8]float64(a), imm8))
}

func m512Extractf64x2Pd(a [8]float64, imm8 int) [2]float64


// M512MaskExtractf64x2Pd: Extract 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'a', selected with
// 'imm8', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
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
func M512MaskExtractf64x2Pd(src x86.M128d, k x86.Mmask8, a x86.M512d, imm8 int) x86.M128d {
	return x86.M128d(m512MaskExtractf64x2Pd([2]float64(src), uint8(k), [8]float64(a), imm8))
}

func m512MaskExtractf64x2Pd(src [2]float64, k uint8, a [8]float64, imm8 int) [2]float64


// M512MaskzExtractf64x2Pd: Extract 128 bits (composed of 2 packed
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
func M512MaskzExtractf64x2Pd(k x86.Mmask8, a x86.M512d, imm8 int) x86.M128d {
	return x86.M128d(m512MaskzExtractf64x2Pd(uint8(k), [8]float64(a), imm8))
}

func m512MaskzExtractf64x2Pd(k uint8, a [8]float64, imm8 int) [2]float64


// M512Extracti32x8Epi32: Extract 256 bits (composed of 8 packed 32-bit
// integers) from 'a', selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[255:0] := a[255:0]
//		1: dst[255:0] := a[511:256]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VEXTRACTI32X8'. Intrinsic: '_mm512_extracti32x8_epi32'.
// Requires AVX512DQ.
func M512Extracti32x8Epi32(a x86.M512i, imm8 int) x86.M256i {
	return x86.M256i(m512Extracti32x8Epi32([64]byte(a), imm8))
}

func m512Extracti32x8Epi32(a [64]byte, imm8 int) [32]byte


// M512MaskExtracti32x8Epi32: Extract 256 bits (composed of 8 packed 32-bit
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
func M512MaskExtracti32x8Epi32(src x86.M256i, k x86.Mmask8, a x86.M512i, imm8 int) x86.M256i {
	return x86.M256i(m512MaskExtracti32x8Epi32([32]byte(src), uint8(k), [64]byte(a), imm8))
}

func m512MaskExtracti32x8Epi32(src [32]byte, k uint8, a [64]byte, imm8 int) [32]byte


// M512MaskzExtracti32x8Epi32: Extract 256 bits (composed of 8 packed 32-bit
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
func M512MaskzExtracti32x8Epi32(k x86.Mmask8, a x86.M512i, imm8 int) x86.M256i {
	return x86.M256i(m512MaskzExtracti32x8Epi32(uint8(k), [64]byte(a), imm8))
}

func m512MaskzExtracti32x8Epi32(k uint8, a [64]byte, imm8 int) [32]byte


// M256Extracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
// integers) from 'a', selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTI64X2'. Intrinsic: '_mm256_extracti64x2_epi64'.
// Requires AVX512DQ.
func M256Extracti64x2Epi64(a x86.M256i, imm8 int) x86.M128i {
	return x86.M128i(m256Extracti64x2Epi64([32]byte(a), imm8))
}

func m256Extracti64x2Epi64(a [32]byte, imm8 int) [16]byte


// M256MaskExtracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
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
func M256MaskExtracti64x2Epi64(src x86.M128i, k x86.Mmask8, a x86.M256i, imm8 int) x86.M128i {
	return x86.M128i(m256MaskExtracti64x2Epi64([16]byte(src), uint8(k), [32]byte(a), imm8))
}

func m256MaskExtracti64x2Epi64(src [16]byte, k uint8, a [32]byte, imm8 int) [16]byte


// M256MaskzExtracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
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
func M256MaskzExtracti64x2Epi64(k x86.Mmask8, a x86.M256i, imm8 int) x86.M128i {
	return x86.M128i(m256MaskzExtracti64x2Epi64(uint8(k), [32]byte(a), imm8))
}

func m256MaskzExtracti64x2Epi64(k uint8, a [32]byte, imm8 int) [16]byte


// M512Extracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
// integers) from 'a', selected with 'imm8', and store the result in 'dst'. 
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
func M512Extracti64x2Epi64(a x86.M512i, imm8 int) x86.M128i {
	return x86.M128i(m512Extracti64x2Epi64([64]byte(a), imm8))
}

func m512Extracti64x2Epi64(a [64]byte, imm8 int) [16]byte


// M512MaskExtracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
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
func M512MaskExtracti64x2Epi64(src x86.M128i, k x86.Mmask8, a x86.M512i, imm8 int) x86.M128i {
	return x86.M128i(m512MaskExtracti64x2Epi64([16]byte(src), uint8(k), [64]byte(a), imm8))
}

func m512MaskExtracti64x2Epi64(src [16]byte, k uint8, a [64]byte, imm8 int) [16]byte


// M512MaskzExtracti64x2Epi64: Extract 128 bits (composed of 2 packed 64-bit
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
func M512MaskzExtracti64x2Epi64(k x86.Mmask8, a x86.M512i, imm8 int) x86.M128i {
	return x86.M128i(m512MaskzExtracti64x2Epi64(uint8(k), [64]byte(a), imm8))
}

func m512MaskzExtracti64x2Epi64(k uint8, a [64]byte, imm8 int) [16]byte


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
func FpclassPdMask(a x86.M128d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(fpclassPdMask([2]float64(a), imm8))
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
func MaskFpclassPdMask(k1 x86.Mmask8, a x86.M128d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskFpclassPdMask(uint8(k1), [2]float64(a), imm8))
}

func maskFpclassPdMask(k1 uint8, a [2]float64, imm8 int) uint8


// M256FpclassPdMask: Test packed double-precision (64-bit) floating-point
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
func M256FpclassPdMask(a x86.M256d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(m256FpclassPdMask([4]float64(a), imm8))
}

func m256FpclassPdMask(a [4]float64, imm8 int) uint8


// M256MaskFpclassPdMask: Test packed double-precision (64-bit) floating-point
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
func M256MaskFpclassPdMask(k1 x86.Mmask8, a x86.M256d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(m256MaskFpclassPdMask(uint8(k1), [4]float64(a), imm8))
}

func m256MaskFpclassPdMask(k1 uint8, a [4]float64, imm8 int) uint8


// M512FpclassPdMask: Test packed double-precision (64-bit) floating-point
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
func M512FpclassPdMask(a x86.M512d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(m512FpclassPdMask([8]float64(a), imm8))
}

func m512FpclassPdMask(a [8]float64, imm8 int) uint8


// M512MaskFpclassPdMask: Test packed double-precision (64-bit) floating-point
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
func M512MaskFpclassPdMask(k1 x86.Mmask8, a x86.M512d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(m512MaskFpclassPdMask(uint8(k1), [8]float64(a), imm8))
}

func m512MaskFpclassPdMask(k1 uint8, a [8]float64, imm8 int) uint8


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
func FpclassPsMask(a x86.M128, imm8 int) x86.Mmask8 {
	return x86.Mmask8(fpclassPsMask([4]float32(a), imm8))
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
func MaskFpclassPsMask(k1 x86.Mmask8, a x86.M128, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskFpclassPsMask(uint8(k1), [4]float32(a), imm8))
}

func maskFpclassPsMask(k1 uint8, a [4]float32, imm8 int) uint8


// M256FpclassPsMask: Test packed single-precision (32-bit) floating-point
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
func M256FpclassPsMask(a x86.M256, imm8 int) x86.Mmask8 {
	return x86.Mmask8(m256FpclassPsMask([8]float32(a), imm8))
}

func m256FpclassPsMask(a [8]float32, imm8 int) uint8


// M256MaskFpclassPsMask: Test packed single-precision (32-bit) floating-point
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
func M256MaskFpclassPsMask(k1 x86.Mmask8, a x86.M256, imm8 int) x86.Mmask8 {
	return x86.Mmask8(m256MaskFpclassPsMask(uint8(k1), [8]float32(a), imm8))
}

func m256MaskFpclassPsMask(k1 uint8, a [8]float32, imm8 int) uint8


// M512FpclassPsMask: Test packed single-precision (32-bit) floating-point
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
func M512FpclassPsMask(a x86.M512, imm8 int) x86.Mmask16 {
	return x86.Mmask16(m512FpclassPsMask([16]float32(a), imm8))
}

func m512FpclassPsMask(a [16]float32, imm8 int) uint16


// M512MaskFpclassPsMask: Test packed single-precision (32-bit) floating-point
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
func M512MaskFpclassPsMask(k1 x86.Mmask16, a x86.M512, imm8 int) x86.Mmask16 {
	return x86.Mmask16(m512MaskFpclassPsMask(uint16(k1), [16]float32(a), imm8))
}

func m512MaskFpclassPsMask(k1 uint16, a [16]float32, imm8 int) uint16


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
func FpclassSdMask(a x86.M128d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(fpclassSdMask([2]float64(a), imm8))
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
func MaskFpclassSdMask(k1 x86.Mmask8, a x86.M128d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskFpclassSdMask(uint8(k1), [2]float64(a), imm8))
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
func FpclassSsMask(a x86.M128, imm8 int) x86.Mmask8 {
	return x86.Mmask8(fpclassSsMask([4]float32(a), imm8))
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
func MaskFpclassSsMask(k1 x86.Mmask8, a x86.M128, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskFpclassSsMask(uint8(k1), [4]float32(a), imm8))
}

func maskFpclassSsMask(k1 uint8, a [4]float32, imm8 int) uint8


// M512Insertf32x8: Copy 'a' to 'dst', then insert 256 bits (composed of 8
// packed single-precision (32-bit) floating-point elements) from 'b' into
// 'dst' at the location specified by 'imm8'. 
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
func M512Insertf32x8(a x86.M512, b x86.M256, imm8 int) x86.M512 {
	return x86.M512(m512Insertf32x8([16]float32(a), [8]float32(b), imm8))
}

func m512Insertf32x8(a [16]float32, b [8]float32, imm8 int) [16]float32


// M512MaskInsertf32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
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
func M512MaskInsertf32x8(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M256, imm8 int) x86.M512 {
	return x86.M512(m512MaskInsertf32x8([16]float32(src), uint16(k), [16]float32(a), [8]float32(b), imm8))
}

func m512MaskInsertf32x8(src [16]float32, k uint16, a [16]float32, b [8]float32, imm8 int) [16]float32


// M512MaskzInsertf32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
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
func M512MaskzInsertf32x8(k x86.Mmask16, a x86.M512, b x86.M256, imm8 int) x86.M512 {
	return x86.M512(m512MaskzInsertf32x8(uint16(k), [16]float32(a), [8]float32(b), imm8))
}

func m512MaskzInsertf32x8(k uint16, a [16]float32, b [8]float32, imm8 int) [16]float32


// M256Insertf64x2: Copy 'a' to 'dst', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'dst' at the location specified by 'imm8'. 
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
func M256Insertf64x2(a x86.M256d, b x86.M128d, imm8 int) x86.M256d {
	return x86.M256d(m256Insertf64x2([4]float64(a), [2]float64(b), imm8))
}

func m256Insertf64x2(a [4]float64, b [2]float64, imm8 int) [4]float64


// M256MaskInsertf64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M256MaskInsertf64x2(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M128d, imm8 int) x86.M256d {
	return x86.M256d(m256MaskInsertf64x2([4]float64(src), uint8(k), [4]float64(a), [2]float64(b), imm8))
}

func m256MaskInsertf64x2(src [4]float64, k uint8, a [4]float64, b [2]float64, imm8 int) [4]float64


// M256MaskzInsertf64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M256MaskzInsertf64x2(k x86.Mmask8, a x86.M256d, b x86.M128d, imm8 int) x86.M256d {
	return x86.M256d(m256MaskzInsertf64x2(uint8(k), [4]float64(a), [2]float64(b), imm8))
}

func m256MaskzInsertf64x2(k uint8, a [4]float64, b [2]float64, imm8 int) [4]float64


// M512Insertf64x2: Copy 'a' to 'dst', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'dst' at the location specified by 'imm8'. 
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
func M512Insertf64x2(a x86.M512d, b x86.M128d, imm8 int) x86.M512d {
	return x86.M512d(m512Insertf64x2([8]float64(a), [2]float64(b), imm8))
}

func m512Insertf64x2(a [8]float64, b [2]float64, imm8 int) [8]float64


// M512MaskInsertf64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M512MaskInsertf64x2(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M128d, imm8 int) x86.M512d {
	return x86.M512d(m512MaskInsertf64x2([8]float64(src), uint8(k), [8]float64(a), [2]float64(b), imm8))
}

func m512MaskInsertf64x2(src [8]float64, k uint8, a [8]float64, b [2]float64, imm8 int) [8]float64


// M512MaskzInsertf64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M512MaskzInsertf64x2(k x86.Mmask8, a x86.M512d, b x86.M128d, imm8 int) x86.M512d {
	return x86.M512d(m512MaskzInsertf64x2(uint8(k), [8]float64(a), [2]float64(b), imm8))
}

func m512MaskzInsertf64x2(k uint8, a [8]float64, b [2]float64, imm8 int) [8]float64


// M512Inserti32x8: Copy 'a' to 'dst', then insert 256 bits (composed of 8
// packed 32-bit integers) from 'b' into 'dst' at the location specified by
// 'imm8'. 
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
func M512Inserti32x8(a x86.M512i, b x86.M256i, imm8 int) x86.M512i {
	return x86.M512i(m512Inserti32x8([64]byte(a), [32]byte(b), imm8))
}

func m512Inserti32x8(a [64]byte, b [32]byte, imm8 int) [64]byte


// M512MaskInserti32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
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
func M512MaskInserti32x8(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M256i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskInserti32x8([64]byte(src), uint16(k), [64]byte(a), [32]byte(b), imm8))
}

func m512MaskInserti32x8(src [64]byte, k uint16, a [64]byte, b [32]byte, imm8 int) [64]byte


// M512MaskzInserti32x8: Copy 'a' to 'tmp', then insert 256 bits (composed of 8
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
func M512MaskzInserti32x8(k x86.Mmask16, a x86.M512i, b x86.M256i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskzInserti32x8(uint16(k), [64]byte(a), [32]byte(b), imm8))
}

func m512MaskzInserti32x8(k uint16, a [64]byte, b [32]byte, imm8 int) [64]byte


// M256Inserti64x2: Copy 'a' to 'dst', then insert 128 bits (composed of 2
// packed 64-bit integers) from 'b' into 'dst' at the location specified by
// 'imm8'. 
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
func M256Inserti64x2(a x86.M256i, b x86.M128i, imm8 int) x86.M256i {
	return x86.M256i(m256Inserti64x2([32]byte(a), [16]byte(b), imm8))
}

func m256Inserti64x2(a [32]byte, b [16]byte, imm8 int) [32]byte


// M256MaskInserti64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M256MaskInserti64x2(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M128i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskInserti64x2([32]byte(src), uint8(k), [32]byte(a), [16]byte(b), imm8))
}

func m256MaskInserti64x2(src [32]byte, k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte


// M256MaskzInserti64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M256MaskzInserti64x2(k x86.Mmask8, a x86.M256i, b x86.M128i, imm8 int) x86.M256i {
	return x86.M256i(m256MaskzInserti64x2(uint8(k), [32]byte(a), [16]byte(b), imm8))
}

func m256MaskzInserti64x2(k uint8, a [32]byte, b [16]byte, imm8 int) [32]byte


// M512Inserti64x2: Copy 'a' to 'dst', then insert 128 bits (composed of 2
// packed 64-bit integers) from 'b' into 'dst' at the location specified by
// 'imm8'. 
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
func M512Inserti64x2(a x86.M512i, b x86.M128i, imm8 int) x86.M512i {
	return x86.M512i(m512Inserti64x2([64]byte(a), [16]byte(b), imm8))
}

func m512Inserti64x2(a [64]byte, b [16]byte, imm8 int) [64]byte


// M512MaskInserti64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M512MaskInserti64x2(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M128i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskInserti64x2([64]byte(src), uint8(k), [64]byte(a), [16]byte(b), imm8))
}

func m512MaskInserti64x2(src [64]byte, k uint8, a [64]byte, b [16]byte, imm8 int) [64]byte


// M512MaskzInserti64x2: Copy 'a' to 'tmp', then insert 128 bits (composed of 2
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
func M512MaskzInserti64x2(k x86.Mmask8, a x86.M512i, b x86.M128i, imm8 int) x86.M512i {
	return x86.M512i(m512MaskzInserti64x2(uint8(k), [64]byte(a), [16]byte(b), imm8))
}

func m512MaskzInserti64x2(k uint8, a [64]byte, b [16]byte, imm8 int) [64]byte


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
func Movepi32Mask(a x86.M128i) x86.Mmask8 {
	return x86.Mmask8(movepi32Mask([16]byte(a)))
}

func movepi32Mask(a [16]byte) uint8


// M256Movepi32Mask: Set each bit of mask register 'k' based on the most
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
func M256Movepi32Mask(a x86.M256i) x86.Mmask8 {
	return x86.Mmask8(m256Movepi32Mask([32]byte(a)))
}

func m256Movepi32Mask(a [32]byte) uint8


// M512Movepi32Mask: Set each bit of mask register 'k' based on the most
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
func M512Movepi32Mask(a x86.M512i) x86.Mmask16 {
	return x86.Mmask16(m512Movepi32Mask([64]byte(a)))
}

func m512Movepi32Mask(a [64]byte) uint16


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
func Movepi64Mask(a x86.M128i) x86.Mmask8 {
	return x86.Mmask8(movepi64Mask([16]byte(a)))
}

func movepi64Mask(a [16]byte) uint8


// M256Movepi64Mask: Set each bit of mask register 'k' based on the most
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
func M256Movepi64Mask(a x86.M256i) x86.Mmask8 {
	return x86.Mmask8(m256Movepi64Mask([32]byte(a)))
}

func m256Movepi64Mask(a [32]byte) uint8


// M512Movepi64Mask: Set each bit of mask register 'k' based on the most
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
func M512Movepi64Mask(a x86.M512i) x86.Mmask8 {
	return x86.Mmask8(m512Movepi64Mask([64]byte(a)))
}

func m512Movepi64Mask(a [64]byte) uint8


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
func MovmEpi32(k x86.Mmask8) x86.M128i {
	return x86.M128i(movmEpi32(uint8(k)))
}

func movmEpi32(k uint8) [16]byte


// M256MovmEpi32: Set each packed 32-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
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
func M256MovmEpi32(k x86.Mmask8) x86.M256i {
	return x86.M256i(m256MovmEpi32(uint8(k)))
}

func m256MovmEpi32(k uint8) [32]byte


// M512MovmEpi32: Set each packed 32-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
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
func M512MovmEpi32(k x86.Mmask16) x86.M512i {
	return x86.M512i(m512MovmEpi32(uint16(k)))
}

func m512MovmEpi32(k uint16) [64]byte


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
func MovmEpi64(k x86.Mmask8) x86.M128i {
	return x86.M128i(movmEpi64(uint8(k)))
}

func movmEpi64(k uint8) [16]byte


// M256MovmEpi64: Set each packed 64-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
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
func M256MovmEpi64(k x86.Mmask8) x86.M256i {
	return x86.M256i(m256MovmEpi64(uint8(k)))
}

func m256MovmEpi64(k uint8) [32]byte


// M512MovmEpi64: Set each packed 64-bit integer in 'dst' to all ones or all
// zeros based on the value of the corresponding bit in 'k'. 
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
func M512MovmEpi64(k x86.Mmask8) x86.M512i {
	return x86.M512i(m512MovmEpi64(uint8(k)))
}

func m512MovmEpi64(k uint8) [64]byte


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
func MaskMulloEpi64(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskMulloEpi64([16]byte(src), uint8(k), [16]byte(a), [16]byte(b)))
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
func MaskzMulloEpi64(k x86.Mmask8, a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maskzMulloEpi64(uint8(k), [16]byte(a), [16]byte(b)))
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
func MulloEpi64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(mulloEpi64([16]byte(a), [16]byte(b)))
}

func mulloEpi64(a [16]byte, b [16]byte) [16]byte


// M256MaskMulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
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
func M256MaskMulloEpi64(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskMulloEpi64([32]byte(src), uint8(k), [32]byte(a), [32]byte(b)))
}

func m256MaskMulloEpi64(src [32]byte, k uint8, a [32]byte, b [32]byte) [32]byte


// M256MaskzMulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
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
func M256MaskzMulloEpi64(k x86.Mmask8, a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MaskzMulloEpi64(uint8(k), [32]byte(a), [32]byte(b)))
}

func m256MaskzMulloEpi64(k uint8, a [32]byte, b [32]byte) [32]byte


// M256MulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst'. 
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
func M256MulloEpi64(a x86.M256i, b x86.M256i) x86.M256i {
	return x86.M256i(m256MulloEpi64([32]byte(a), [32]byte(b)))
}

func m256MulloEpi64(a [32]byte, b [32]byte) [32]byte


// M512MaskMulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
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
func M512MaskMulloEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskMulloEpi64([64]byte(src), uint8(k), [64]byte(a), [64]byte(b)))
}

func m512MaskMulloEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte


// M512MaskzMulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
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
func M512MaskzMulloEpi64(k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MaskzMulloEpi64(uint8(k), [64]byte(a), [64]byte(b)))
}

func m512MaskzMulloEpi64(k uint8, a [64]byte, b [64]byte) [64]byte


// M512MulloEpi64: Multiply the packed 64-bit integers in 'a' and 'b',
// producing intermediate 128-bit integers, and store the low 64 bits of the
// intermediate integers in 'dst'. 
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
func M512MulloEpi64(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(m512MulloEpi64([64]byte(a), [64]byte(b)))
}

func m512MulloEpi64(a [64]byte, b [64]byte) [64]byte


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
func MaskOrPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskOrPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
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
func MaskzOrPd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzOrPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzOrPd(k uint8, a [2]float64, b [2]float64) [2]float64


// M256MaskOrPd: Compute the bitwise OR of packed double-precision (64-bit)
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
func M256MaskOrPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskOrPd([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskOrPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// M256MaskzOrPd: Compute the bitwise OR of packed double-precision (64-bit)
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
func M256MaskzOrPd(k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskzOrPd(uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskzOrPd(k uint8, a [4]float64, b [4]float64) [4]float64


// M512MaskOrPd: Compute the bitwise OR of packed double-precision (64-bit)
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
func M512MaskOrPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskOrPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskOrPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// M512MaskzOrPd: Compute the bitwise OR of packed double-precision (64-bit)
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
func M512MaskzOrPd(k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzOrPd(uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskzOrPd(k uint8, a [8]float64, b [8]float64) [8]float64


// M512OrPd: Compute the bitwise OR of packed double-precision (64-bit)
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
func M512OrPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512OrPd([8]float64(a), [8]float64(b)))
}

func m512OrPd(a [8]float64, b [8]float64) [8]float64


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
func MaskOrPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskOrPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
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
func MaskzOrPs(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzOrPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzOrPs(k uint8, a [4]float32, b [4]float32) [4]float32


// M256MaskOrPs: Compute the bitwise OR of packed single-precision (32-bit)
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
func M256MaskOrPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskOrPs([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskOrPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// M256MaskzOrPs: Compute the bitwise OR of packed single-precision (32-bit)
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
func M256MaskzOrPs(k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskzOrPs(uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskzOrPs(k uint8, a [8]float32, b [8]float32) [8]float32


// M512MaskOrPs: Compute the bitwise OR of packed single-precision (32-bit)
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
func M512MaskOrPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskOrPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskOrPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// M512MaskzOrPs: Compute the bitwise OR of packed single-precision (32-bit)
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
func M512MaskzOrPs(k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskzOrPs(uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskzOrPs(k uint16, a [16]float32, b [16]float32) [16]float32


// M512OrPs: Compute the bitwise OR of packed single-precision (32-bit)
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
func M512OrPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512OrPs([16]float32(a), [16]float32(b)))
}

func m512OrPs(a [16]float32, b [16]float32) [16]float32


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
func MaskRangePd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskRangePd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8))
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
func MaskzRangePd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskzRangePd(uint8(k), [2]float64(a), [2]float64(b), imm8))
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
func RangePd(a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(rangePd([2]float64(a), [2]float64(b), imm8))
}

func rangePd(a [2]float64, b [2]float64, imm8 int) [2]float64


// M256MaskRangePd: Calculate the max, min, absolute max, or absolute min
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
func M256MaskRangePd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d, imm8 int) x86.M256d {
	return x86.M256d(m256MaskRangePd([4]float64(src), uint8(k), [4]float64(a), [4]float64(b), imm8))
}

func m256MaskRangePd(src [4]float64, k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64


// M256MaskzRangePd: Calculate the max, min, absolute max, or absolute min
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
func M256MaskzRangePd(k x86.Mmask8, a x86.M256d, b x86.M256d, imm8 int) x86.M256d {
	return x86.M256d(m256MaskzRangePd(uint8(k), [4]float64(a), [4]float64(b), imm8))
}

func m256MaskzRangePd(k uint8, a [4]float64, b [4]float64, imm8 int) [4]float64


// M256RangePd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'.
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
func M256RangePd(a x86.M256d, b x86.M256d, imm8 int) x86.M256d {
	return x86.M256d(m256RangePd([4]float64(a), [4]float64(b), imm8))
}

func m256RangePd(a [4]float64, b [4]float64, imm8 int) [4]float64


// M512MaskRangePd: Calculate the max, min, absolute max, or absolute min
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
func M512MaskRangePd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 int) x86.M512d {
	return x86.M512d(m512MaskRangePd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), imm8))
}

func m512MaskRangePd(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64


// M512MaskzRangePd: Calculate the max, min, absolute max, or absolute min
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
func M512MaskzRangePd(k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 int) x86.M512d {
	return x86.M512d(m512MaskzRangePd(uint8(k), [8]float64(a), [8]float64(b), imm8))
}

func m512MaskzRangePd(k uint8, a [8]float64, b [8]float64, imm8 int) [8]float64


// M512RangePd: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'.
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
func M512RangePd(a x86.M512d, b x86.M512d, imm8 int) x86.M512d {
	return x86.M512d(m512RangePd([8]float64(a), [8]float64(b), imm8))
}

func m512RangePd(a [8]float64, b [8]float64, imm8 int) [8]float64


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
func MaskRangePs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskRangePs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8))
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
func MaskzRangePs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskzRangePs(uint8(k), [4]float32(a), [4]float32(b), imm8))
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
func RangePs(a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(rangePs([4]float32(a), [4]float32(b), imm8))
}

func rangePs(a [4]float32, b [4]float32, imm8 int) [4]float32


// M256MaskRangePs: Calculate the max, min, absolute max, or absolute min
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
func M256MaskRangePs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256, imm8 int) x86.M256 {
	return x86.M256(m256MaskRangePs([8]float32(src), uint8(k), [8]float32(a), [8]float32(b), imm8))
}

func m256MaskRangePs(src [8]float32, k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32


// M256MaskzRangePs: Calculate the max, min, absolute max, or absolute min
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
func M256MaskzRangePs(k x86.Mmask8, a x86.M256, b x86.M256, imm8 int) x86.M256 {
	return x86.M256(m256MaskzRangePs(uint8(k), [8]float32(a), [8]float32(b), imm8))
}

func m256MaskzRangePs(k uint8, a [8]float32, b [8]float32, imm8 int) [8]float32


// M256RangePs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'.
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
func M256RangePs(a x86.M256, b x86.M256, imm8 int) x86.M256 {
	return x86.M256(m256RangePs([8]float32(a), [8]float32(b), imm8))
}

func m256RangePs(a [8]float32, b [8]float32, imm8 int) [8]float32


// M512MaskRangePs: Calculate the max, min, absolute max, or absolute min
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
func M512MaskRangePs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, imm8 int) x86.M512 {
	return x86.M512(m512MaskRangePs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), imm8))
}

func m512MaskRangePs(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32


// M512MaskzRangePs: Calculate the max, min, absolute max, or absolute min
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
func M512MaskzRangePs(k x86.Mmask16, a x86.M512, b x86.M512, imm8 int) x86.M512 {
	return x86.M512(m512MaskzRangePs(uint16(k), [16]float32(a), [16]float32(b), imm8))
}

func m512MaskzRangePs(k uint16, a [16]float32, b [16]float32, imm8 int) [16]float32


// M512RangePs: Calculate the max, min, absolute max, or absolute min
// (depending on control in 'imm8') for packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'.
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
func M512RangePs(a x86.M512, b x86.M512, imm8 int) x86.M512 {
	return x86.M512(m512RangePs([16]float32(a), [16]float32(b), imm8))
}

func m512RangePs(a [16]float32, b [16]float32, imm8 int) [16]float32


// M512MaskRangeRoundPd: Calculate the max, min, absolute max, or absolute min
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
func M512MaskRangeRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 int, rounding int) x86.M512d {
	return x86.M512d(m512MaskRangeRoundPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), imm8, rounding))
}

func m512MaskRangeRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64


// M512MaskzRangeRoundPd: Calculate the max, min, absolute max, or absolute min
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
func M512MaskzRangeRoundPd(k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 int, rounding int) x86.M512d {
	return x86.M512d(m512MaskzRangeRoundPd(uint8(k), [8]float64(a), [8]float64(b), imm8, rounding))
}

func m512MaskzRangeRoundPd(k uint8, a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64


// M512RangeRoundPd: Calculate the max, min, absolute max, or absolute min
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
func M512RangeRoundPd(a x86.M512d, b x86.M512d, imm8 int, rounding int) x86.M512d {
	return x86.M512d(m512RangeRoundPd([8]float64(a), [8]float64(b), imm8, rounding))
}

func m512RangeRoundPd(a [8]float64, b [8]float64, imm8 int, rounding int) [8]float64


// M512MaskRangeRoundPs: Calculate the max, min, absolute max, or absolute min
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
func M512MaskRangeRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, imm8 int, rounding int) x86.M512 {
	return x86.M512(m512MaskRangeRoundPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), imm8, rounding))
}

func m512MaskRangeRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32


// M512MaskzRangeRoundPs: Calculate the max, min, absolute max, or absolute min
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
func M512MaskzRangeRoundPs(k x86.Mmask16, a x86.M512, b x86.M512, imm8 int, rounding int) x86.M512 {
	return x86.M512(m512MaskzRangeRoundPs(uint16(k), [16]float32(a), [16]float32(b), imm8, rounding))
}

func m512MaskzRangeRoundPs(k uint16, a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32


// M512RangeRoundPs: Calculate the max, min, absolute max, or absolute min
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
func M512RangeRoundPs(a x86.M512, b x86.M512, imm8 int, rounding int) x86.M512 {
	return x86.M512(m512RangeRoundPs([16]float32(a), [16]float32(b), imm8, rounding))
}

func m512RangeRoundPs(a [16]float32, b [16]float32, imm8 int, rounding int) [16]float32


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
func MaskRangeRoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int, rounding int) x86.M128d {
	return x86.M128d(maskRangeRoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
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
func MaskzRangeRoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int, rounding int) x86.M128d {
	return x86.M128d(maskzRangeRoundSd(uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
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
func RangeRoundSd(a x86.M128d, b x86.M128d, imm8 int, rounding int) x86.M128d {
	return x86.M128d(rangeRoundSd([2]float64(a), [2]float64(b), imm8, rounding))
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
func MaskRangeRoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 int, rounding int) x86.M128 {
	return x86.M128(maskRangeRoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
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
func MaskzRangeRoundSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 int, rounding int) x86.M128 {
	return x86.M128(maskzRangeRoundSs(uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
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
func RangeRoundSs(a x86.M128, b x86.M128, imm8 int, rounding int) x86.M128 {
	return x86.M128(rangeRoundSs([4]float32(a), [4]float32(b), imm8, rounding))
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
func MaskRangeSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskRangeSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8))
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
func MaskzRangeSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskzRangeSd(uint8(k), [2]float64(a), [2]float64(b), imm8))
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
func MaskRangeSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskRangeSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8))
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
func MaskzRangeSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskzRangeSs(uint8(k), [4]float32(a), [4]float32(b), imm8))
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
func MaskReducePd(src x86.M128d, k x86.Mmask8, a x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskReducePd([2]float64(src), uint8(k), [2]float64(a), imm8))
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
func MaskzReducePd(k x86.Mmask8, a x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskzReducePd(uint8(k), [2]float64(a), imm8))
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
func ReducePd(a x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(reducePd([2]float64(a), imm8))
}

func reducePd(a [2]float64, imm8 int) [2]float64


// M256MaskReducePd: Extract the reduced argument of packed double-precision
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
func M256MaskReducePd(src x86.M256d, k x86.Mmask8, a x86.M256d, imm8 int) x86.M256d {
	return x86.M256d(m256MaskReducePd([4]float64(src), uint8(k), [4]float64(a), imm8))
}

func m256MaskReducePd(src [4]float64, k uint8, a [4]float64, imm8 int) [4]float64


// M256MaskzReducePd: Extract the reduced argument of packed double-precision
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
func M256MaskzReducePd(k x86.Mmask8, a x86.M256d, imm8 int) x86.M256d {
	return x86.M256d(m256MaskzReducePd(uint8(k), [4]float64(a), imm8))
}

func m256MaskzReducePd(k uint8, a [4]float64, imm8 int) [4]float64


// M256ReducePd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst'. 
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
func M256ReducePd(a x86.M256d, imm8 int) x86.M256d {
	return x86.M256d(m256ReducePd([4]float64(a), imm8))
}

func m256ReducePd(a [4]float64, imm8 int) [4]float64


// M512MaskReducePd: Extract the reduced argument of packed double-precision
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
func M512MaskReducePd(src x86.M512d, k x86.Mmask8, a x86.M512d, imm8 int) x86.M512d {
	return x86.M512d(m512MaskReducePd([8]float64(src), uint8(k), [8]float64(a), imm8))
}

func m512MaskReducePd(src [8]float64, k uint8, a [8]float64, imm8 int) [8]float64


// M512MaskzReducePd: Extract the reduced argument of packed double-precision
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
func M512MaskzReducePd(k x86.Mmask8, a x86.M512d, imm8 int) x86.M512d {
	return x86.M512d(m512MaskzReducePd(uint8(k), [8]float64(a), imm8))
}

func m512MaskzReducePd(k uint8, a [8]float64, imm8 int) [8]float64


// M512ReducePd: Extract the reduced argument of packed double-precision
// (64-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst'. 
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
func M512ReducePd(a x86.M512d, imm8 int) x86.M512d {
	return x86.M512d(m512ReducePd([8]float64(a), imm8))
}

func m512ReducePd(a [8]float64, imm8 int) [8]float64


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
func MaskReducePs(src x86.M128, k x86.Mmask8, a x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskReducePs([4]float32(src), uint8(k), [4]float32(a), imm8))
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
func MaskzReducePs(k x86.Mmask8, a x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskzReducePs(uint8(k), [4]float32(a), imm8))
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
func ReducePs(a x86.M128, imm8 int) x86.M128 {
	return x86.M128(reducePs([4]float32(a), imm8))
}

func reducePs(a [4]float32, imm8 int) [4]float32


// M256MaskReducePs: Extract the reduced argument of packed single-precision
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
func M256MaskReducePs(src x86.M256, k x86.Mmask8, a x86.M256, imm8 int) x86.M256 {
	return x86.M256(m256MaskReducePs([8]float32(src), uint8(k), [8]float32(a), imm8))
}

func m256MaskReducePs(src [8]float32, k uint8, a [8]float32, imm8 int) [8]float32


// M256MaskzReducePs: Extract the reduced argument of packed single-precision
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
func M256MaskzReducePs(k x86.Mmask8, a x86.M256, imm8 int) x86.M256 {
	return x86.M256(m256MaskzReducePs(uint8(k), [8]float32(a), imm8))
}

func m256MaskzReducePs(k uint8, a [8]float32, imm8 int) [8]float32


// M256ReducePs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst'. 
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
func M256ReducePs(a x86.M256, imm8 int) x86.M256 {
	return x86.M256(m256ReducePs([8]float32(a), imm8))
}

func m256ReducePs(a [8]float32, imm8 int) [8]float32


// M512MaskReducePs: Extract the reduced argument of packed single-precision
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
func M512MaskReducePs(src x86.M512, k x86.Mmask16, a x86.M512, imm8 int) x86.M512 {
	return x86.M512(m512MaskReducePs([16]float32(src), uint16(k), [16]float32(a), imm8))
}

func m512MaskReducePs(src [16]float32, k uint16, a [16]float32, imm8 int) [16]float32


// M512MaskzReducePs: Extract the reduced argument of packed single-precision
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
func M512MaskzReducePs(k x86.Mmask16, a x86.M512, imm8 int) x86.M512 {
	return x86.M512(m512MaskzReducePs(uint16(k), [16]float32(a), imm8))
}

func m512MaskzReducePs(k uint16, a [16]float32, imm8 int) [16]float32


// M512ReducePs: Extract the reduced argument of packed single-precision
// (32-bit) floating-point elements in 'a' by the number of bits specified by
// 'imm8', and store the results in 'dst'. 
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
func M512ReducePs(a x86.M512, imm8 int) x86.M512 {
	return x86.M512(m512ReducePs([16]float32(a), imm8))
}

func m512ReducePs(a [16]float32, imm8 int) [16]float32


// M512MaskReduceRoundPd: Extract the reduced argument of packed
// double-precision (64-bit) floating-point elements in 'a' by the number of
// bits specified by 'imm8', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
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
func M512MaskReduceRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, imm8 int, rounding int) x86.M512d {
	return x86.M512d(m512MaskReduceRoundPd([8]float64(src), uint8(k), [8]float64(a), imm8, rounding))
}

func m512MaskReduceRoundPd(src [8]float64, k uint8, a [8]float64, imm8 int, rounding int) [8]float64


// M512MaskzReduceRoundPd: Extract the reduced argument of packed
// double-precision (64-bit) floating-point elements in 'a' by the number of
// bits specified by 'imm8', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set).
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
func M512MaskzReduceRoundPd(k x86.Mmask8, a x86.M512d, imm8 int, rounding int) x86.M512d {
	return x86.M512d(m512MaskzReduceRoundPd(uint8(k), [8]float64(a), imm8, rounding))
}

func m512MaskzReduceRoundPd(k uint8, a [8]float64, imm8 int, rounding int) [8]float64


// M512ReduceRoundPd: Extract the reduced argument of packed double-precision
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
func M512ReduceRoundPd(a x86.M512d, imm8 int, rounding int) x86.M512d {
	return x86.M512d(m512ReduceRoundPd([8]float64(a), imm8, rounding))
}

func m512ReduceRoundPd(a [8]float64, imm8 int, rounding int) [8]float64


// M512MaskReduceRoundPs: Extract the reduced argument of packed
// single-precision (32-bit) floating-point elements in 'a' by the number of
// bits specified by 'imm8', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
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
func M512MaskReduceRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, imm8 int, rounding int) x86.M512 {
	return x86.M512(m512MaskReduceRoundPs([16]float32(src), uint16(k), [16]float32(a), imm8, rounding))
}

func m512MaskReduceRoundPs(src [16]float32, k uint16, a [16]float32, imm8 int, rounding int) [16]float32


// M512MaskzReduceRoundPs: Extract the reduced argument of packed
// single-precision (32-bit) floating-point elements in 'a' by the number of
// bits specified by 'imm8', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set).
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
func M512MaskzReduceRoundPs(k x86.Mmask16, a x86.M512, imm8 int, rounding int) x86.M512 {
	return x86.M512(m512MaskzReduceRoundPs(uint16(k), [16]float32(a), imm8, rounding))
}

func m512MaskzReduceRoundPs(k uint16, a [16]float32, imm8 int, rounding int) [16]float32


// M512ReduceRoundPs: Extract the reduced argument of packed single-precision
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
func M512ReduceRoundPs(a x86.M512, imm8 int, rounding int) x86.M512 {
	return x86.M512(m512ReduceRoundPs([16]float32(a), imm8, rounding))
}

func m512ReduceRoundPs(a [16]float32, imm8 int, rounding int) [16]float32


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
func MaskReduceRoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int, rounding int) x86.M128d {
	return x86.M128d(maskReduceRoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
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
func MaskzReduceRoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int, rounding int) x86.M128d {
	return x86.M128d(maskzReduceRoundSd(uint8(k), [2]float64(a), [2]float64(b), imm8, rounding))
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
func ReduceRoundSd(a x86.M128d, b x86.M128d, imm8 int, rounding int) x86.M128d {
	return x86.M128d(reduceRoundSd([2]float64(a), [2]float64(b), imm8, rounding))
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
func MaskReduceRoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 int, rounding int) x86.M128 {
	return x86.M128(maskReduceRoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
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
func MaskzReduceRoundSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 int, rounding int) x86.M128 {
	return x86.M128(maskzReduceRoundSs(uint8(k), [4]float32(a), [4]float32(b), imm8, rounding))
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
func ReduceRoundSs(a x86.M128, b x86.M128, imm8 int, rounding int) x86.M128 {
	return x86.M128(reduceRoundSs([4]float32(a), [4]float32(b), imm8, rounding))
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
func MaskReduceSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskReduceSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), imm8))
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
func MaskzReduceSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(maskzReduceSd(uint8(k), [2]float64(a), [2]float64(b), imm8))
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
func ReduceSd(a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(reduceSd([2]float64(a), [2]float64(b), imm8))
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
func MaskReduceSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskReduceSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), imm8))
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
func MaskzReduceSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(maskzReduceSs(uint8(k), [4]float32(a), [4]float32(b), imm8))
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
func ReduceSs(a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(reduceSs([4]float32(a), [4]float32(b), imm8))
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
func MaskXorPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskXorPd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
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
func MaskzXorPd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzXorPd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzXorPd(k uint8, a [2]float64, b [2]float64) [2]float64


// M256MaskXorPd: Compute the bitwise XOR of packed double-precision (64-bit)
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
func M256MaskXorPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskXorPd([4]float64(src), uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskXorPd(src [4]float64, k uint8, a [4]float64, b [4]float64) [4]float64


// M256MaskzXorPd: Compute the bitwise XOR of packed double-precision (64-bit)
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
func M256MaskzXorPd(k x86.Mmask8, a x86.M256d, b x86.M256d) x86.M256d {
	return x86.M256d(m256MaskzXorPd(uint8(k), [4]float64(a), [4]float64(b)))
}

func m256MaskzXorPd(k uint8, a [4]float64, b [4]float64) [4]float64


// M512MaskXorPd: Compute the bitwise XOR of packed double-precision (64-bit)
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
func M512MaskXorPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskXorPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskXorPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// M512MaskzXorPd: Compute the bitwise XOR of packed double-precision (64-bit)
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
func M512MaskzXorPd(k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzXorPd(uint8(k), [8]float64(a), [8]float64(b)))
}

func m512MaskzXorPd(k uint8, a [8]float64, b [8]float64) [8]float64


// M512XorPd: Compute the bitwise XOR of packed double-precision (64-bit)
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
func M512XorPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(m512XorPd([8]float64(a), [8]float64(b)))
}

func m512XorPd(a [8]float64, b [8]float64) [8]float64


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
func MaskXorPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskXorPs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
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
func MaskzXorPs(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzXorPs(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzXorPs(k uint8, a [4]float32, b [4]float32) [4]float32


// M256MaskXorPs: Compute the bitwise XOR of packed single-precision (32-bit)
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
func M256MaskXorPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskXorPs([8]float32(src), uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskXorPs(src [8]float32, k uint8, a [8]float32, b [8]float32) [8]float32


// M256MaskzXorPs: Compute the bitwise XOR of packed single-precision (32-bit)
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
func M256MaskzXorPs(k x86.Mmask8, a x86.M256, b x86.M256) x86.M256 {
	return x86.M256(m256MaskzXorPs(uint8(k), [8]float32(a), [8]float32(b)))
}

func m256MaskzXorPs(k uint8, a [8]float32, b [8]float32) [8]float32


// M512MaskXorPs: Compute the bitwise XOR of packed single-precision (32-bit)
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
func M512MaskXorPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskXorPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskXorPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// M512MaskzXorPs: Compute the bitwise XOR of packed single-precision (32-bit)
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
func M512MaskzXorPs(k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512MaskzXorPs(uint16(k), [16]float32(a), [16]float32(b)))
}

func m512MaskzXorPs(k uint16, a [16]float32, b [16]float32) [16]float32


// M512XorPs: Compute the bitwise XOR of packed single-precision (32-bit)
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
func M512XorPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(m512XorPs([16]float32(a), [16]float32(b)))
}

func m512XorPs(a [16]float32, b [16]float32) [16]float32

