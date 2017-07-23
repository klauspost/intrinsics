// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
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
func MaskAndPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskzAndPd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func M256MaskAndPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzAndPd(k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512AndPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskAndPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzAndPd(k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func MaskAndPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskzAndPs(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskAndPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M256MaskzAndPs(k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M512AndPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskAndPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskzAndPs(k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func MaskAndnotPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskzAndnotPd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func M256MaskAndnotPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzAndnotPd(k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512AndnotPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskAndnotPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzAndnotPd(k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func MaskAndnotPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskzAndnotPs(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskAndnotPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M256MaskzAndnotPs(k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M512AndnotPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskAndnotPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskzAndnotPs(k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M256BroadcastF32x2(a x86.M128) (dst x86.M256) {
	panic("not implemented")
}


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
func M256MaskBroadcastF32x2(src x86.M256, k x86.Mmask8, a x86.M128) (dst x86.M256) {
	panic("not implemented")
}


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
func M256MaskzBroadcastF32x2(k x86.Mmask8, a x86.M128) (dst x86.M256) {
	panic("not implemented")
}


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
func M512BroadcastF32x2(a x86.M128) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskBroadcastF32x2(src x86.M512, k x86.Mmask16, a x86.M128) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskzBroadcastF32x2(k x86.Mmask16, a x86.M128) (dst x86.M512) {
	panic("not implemented")
}


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
func M512BroadcastF32x8(a x86.M256) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskBroadcastF32x8(src x86.M512, k x86.Mmask16, a x86.M256) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskzBroadcastF32x8(k x86.Mmask16, a x86.M256) (dst x86.M512) {
	panic("not implemented")
}


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
func M256BroadcastF64x2(a x86.M128d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskBroadcastF64x2(src x86.M256d, k x86.Mmask8, a x86.M128d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzBroadcastF64x2(k x86.Mmask8, a x86.M128d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512BroadcastF64x2(a x86.M128d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskBroadcastF64x2(src x86.M512d, k x86.Mmask8, a x86.M128d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzBroadcastF64x2(k x86.Mmask8, a x86.M128d) (dst x86.M512d) {
	panic("not implemented")
}


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
func BroadcastI32x2(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskBroadcastI32x2(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzBroadcastI32x2(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256BroadcastI32x2(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskBroadcastI32x2(src x86.M256i, k x86.Mmask8, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzBroadcastI32x2(k x86.Mmask8, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512BroadcastI32x2(a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskBroadcastI32x2(src x86.M512i, k x86.Mmask16, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzBroadcastI32x2(k x86.Mmask16, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512BroadcastI32x8(a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskBroadcastI32x8(src x86.M512i, k x86.Mmask16, a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzBroadcastI32x8(k x86.Mmask16, a x86.M256i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M256BroadcastI64x2(a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskBroadcastI64x2(src x86.M256i, k x86.Mmask8, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzBroadcastI64x2(k x86.Mmask8, a x86.M128i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512BroadcastI64x2(a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskBroadcastI64x2(src x86.M512i, k x86.Mmask8, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzBroadcastI64x2(k x86.Mmask8, a x86.M128i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvtRoundepi64Pd(a x86.M512i, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskCvtRoundepi64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundepi64Pd(k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512CvtRoundepi64Ps(a x86.M512i, rounding int) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskCvtRoundepi64Ps(src x86.M256, k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundepi64Ps(k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M256) {
	panic("not implemented")
}


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
func M512CvtRoundepu64Pd(a x86.M512i, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskCvtRoundepu64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundepu64Pd(k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512CvtRoundepu64Ps(a x86.M512i, rounding int) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskCvtRoundepu64Ps(src x86.M256, k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundepu64Ps(k x86.Mmask8, a x86.M512i, rounding int) (dst x86.M256) {
	panic("not implemented")
}


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
func M512CvtRoundpdEpi64(a x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtRoundpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundpdEpi64(k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvtRoundpdEpu64(a x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtRoundpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundpdEpu64(k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvtRoundpsEpi64(a x86.M256, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtRoundpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundpsEpi64(k x86.Mmask8, a x86.M256, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvtRoundpsEpu64(a x86.M256, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtRoundpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtRoundpsEpu64(k x86.Mmask8, a x86.M256, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


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
func Cvtepi64Pd(a x86.M128i) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskCvtepi64Pd(src x86.M128d, k x86.Mmask8, a x86.M128i) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskzCvtepi64Pd(k x86.Mmask8, a x86.M128i) (dst x86.M128d) {
	panic("not implemented")
}


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
func M256Cvtepi64Pd(a x86.M256i) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskCvtepi64Pd(src x86.M256d, k x86.Mmask8, a x86.M256i) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzCvtepi64Pd(k x86.Mmask8, a x86.M256i) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512Cvtepi64Pd(a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskCvtepi64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzCvtepi64Pd(k x86.Mmask8, a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


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
func Cvtepi64Ps(a x86.M128i) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskCvtepi64Ps(src x86.M128, k x86.Mmask8, a x86.M128i) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskzCvtepi64Ps(k x86.Mmask8, a x86.M128i) (dst x86.M128) {
	panic("not implemented")
}


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
func M256Cvtepi64Ps(a x86.M256i) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskCvtepi64Ps(src x86.M128, k x86.Mmask8, a x86.M256i) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskzCvtepi64Ps(k x86.Mmask8, a x86.M256i) (dst x86.M128) {
	panic("not implemented")
}


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
func M512Cvtepi64Ps(a x86.M512i) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskCvtepi64Ps(src x86.M256, k x86.Mmask8, a x86.M512i) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskzCvtepi64Ps(k x86.Mmask8, a x86.M512i) (dst x86.M256) {
	panic("not implemented")
}


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
func Cvtepu64Pd(a x86.M128i) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskCvtepu64Pd(src x86.M128d, k x86.Mmask8, a x86.M128i) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskzCvtepu64Pd(k x86.Mmask8, a x86.M128i) (dst x86.M128d) {
	panic("not implemented")
}


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
func M256Cvtepu64Pd(a x86.M256i) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskCvtepu64Pd(src x86.M256d, k x86.Mmask8, a x86.M256i) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzCvtepu64Pd(k x86.Mmask8, a x86.M256i) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512Cvtepu64Pd(a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskCvtepu64Pd(src x86.M512d, k x86.Mmask8, a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzCvtepu64Pd(k x86.Mmask8, a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


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
func Cvtepu64Ps(a x86.M128i) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskCvtepu64Ps(src x86.M128, k x86.Mmask8, a x86.M128i) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskzCvtepu64Ps(k x86.Mmask8, a x86.M128i) (dst x86.M128) {
	panic("not implemented")
}


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
func M256Cvtepu64Ps(a x86.M256i) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskCvtepu64Ps(src x86.M128, k x86.Mmask8, a x86.M256i) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskzCvtepu64Ps(k x86.Mmask8, a x86.M256i) (dst x86.M128) {
	panic("not implemented")
}


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
func M512Cvtepu64Ps(a x86.M512i) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskCvtepu64Ps(src x86.M256, k x86.Mmask8, a x86.M512i) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskzCvtepu64Ps(k x86.Mmask8, a x86.M512i) (dst x86.M256) {
	panic("not implemented")
}


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
func CvtpdEpi64(a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvtpdEpi64(src x86.M128i, k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvtpdEpi64(k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvtpdEpi64(a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvtpdEpi64(src x86.M256i, k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvtpdEpi64(k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvtpdEpi64(a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtpdEpi64(k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvtpdEpu64(a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvtpdEpu64(src x86.M128i, k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvtpdEpu64(k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvtpdEpu64(a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvtpdEpu64(src x86.M256i, k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvtpdEpu64(k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvtpdEpu64(a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtpdEpu64(k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvtpsEpi64(a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvtpsEpi64(src x86.M128i, k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvtpsEpi64(k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvtpsEpi64(a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvtpsEpi64(src x86.M256i, k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvtpsEpi64(k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvtpsEpi64(a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtpsEpi64(k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvtpsEpu64(a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvtpsEpu64(src x86.M128i, k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvtpsEpu64(k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvtpsEpu64(a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvtpsEpu64(src x86.M256i, k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvtpsEpu64(k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvtpsEpu64(a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvtpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvtpsEpu64(k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvttRoundpdEpi64(a x86.M512d, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttRoundpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttRoundpdEpi64(k x86.Mmask8, a x86.M512d, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvttRoundpdEpu64(a x86.M512d, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttRoundpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttRoundpdEpu64(k x86.Mmask8, a x86.M512d, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvttRoundpsEpi64(a x86.M256, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttRoundpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttRoundpsEpi64(k x86.Mmask8, a x86.M256, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512CvttRoundpsEpu64(a x86.M256, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttRoundpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttRoundpsEpu64(k x86.Mmask8, a x86.M256, sae int) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvttpdEpi64(a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvttpdEpi64(src x86.M128i, k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvttpdEpi64(k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvttpdEpi64(a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvttpdEpi64(src x86.M256i, k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvttpdEpi64(k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvttpdEpi64(a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttpdEpi64(src x86.M512i, k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttpdEpi64(k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvttpdEpu64(a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvttpdEpu64(src x86.M128i, k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvttpdEpu64(k x86.Mmask8, a x86.M128d) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvttpdEpu64(a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvttpdEpu64(src x86.M256i, k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvttpdEpu64(k x86.Mmask8, a x86.M256d) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvttpdEpu64(a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttpdEpu64(src x86.M512i, k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttpdEpu64(k x86.Mmask8, a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvttpsEpi64(a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvttpsEpi64(src x86.M128i, k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvttpsEpi64(k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvttpsEpi64(a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvttpsEpi64(src x86.M256i, k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvttpsEpi64(k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvttpsEpi64(a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttpsEpi64(src x86.M512i, k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttpsEpi64(k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func CvttpsEpu64(a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskCvttpsEpu64(src x86.M128i, k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzCvttpsEpu64(k x86.Mmask8, a x86.M128) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256CvttpsEpu64(a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskCvttpsEpu64(src x86.M256i, k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzCvttpsEpu64(k x86.Mmask8, a x86.M128) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512CvttpsEpu64(a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskCvttpsEpu64(src x86.M512i, k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzCvttpsEpu64(k x86.Mmask8, a x86.M256) (dst x86.M512i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Extractf32x8Ps(a x86.M512, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskExtractf32x8Ps(src x86.M256, k x86.Mmask8, a x86.M512, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzExtractf32x8Ps(k x86.Mmask8, a x86.M512, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256Extractf64x2Pd(a x86.M256d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskExtractf64x2Pd(src x86.M128d, k x86.Mmask8, a x86.M256d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzExtractf64x2Pd(k x86.Mmask8, a x86.M256d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Extractf64x2Pd(a x86.M512d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskExtractf64x2Pd(src x86.M128d, k x86.Mmask8, a x86.M512d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzExtractf64x2Pd(k x86.Mmask8, a x86.M512d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Extracti32x8Epi32(a x86.M512i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskExtracti32x8Epi32(src x86.M256i, k x86.Mmask8, a x86.M512i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzExtracti32x8Epi32(k x86.Mmask8, a x86.M512i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256Extracti64x2Epi64(a x86.M256i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskExtracti64x2Epi64(src x86.M128i, k x86.Mmask8, a x86.M256i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzExtracti64x2Epi64(k x86.Mmask8, a x86.M256i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Extracti64x2Epi64(a x86.M512i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskExtracti64x2Epi64(src x86.M128i, k x86.Mmask8, a x86.M512i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzExtracti64x2Epi64(k x86.Mmask8, a x86.M512i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func FpclassPdMask(a x86.M128d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskFpclassPdMask(k1 x86.Mmask8, a x86.M128d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256FpclassPdMask(a x86.M256d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskFpclassPdMask(k1 x86.Mmask8, a x86.M256d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512FpclassPdMask(a x86.M512d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskFpclassPdMask(k1 x86.Mmask8, a x86.M512d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func FpclassPsMask(a x86.M128, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskFpclassPsMask(k1 x86.Mmask8, a x86.M128, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256FpclassPsMask(a x86.M256, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskFpclassPsMask(k1 x86.Mmask8, a x86.M256, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512FpclassPsMask(a x86.M512, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskFpclassPsMask(k1 x86.Mmask16, a x86.M512, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func FpclassSdMask(a x86.M128d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskFpclassSdMask(k1 x86.Mmask8, a x86.M128d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func FpclassSsMask(a x86.M128, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskFpclassSsMask(k1 x86.Mmask8, a x86.M128, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Insertf32x8(a x86.M512, b x86.M256, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskInsertf32x8(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M256, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzInsertf32x8(k x86.Mmask16, a x86.M512, b x86.M256, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256Insertf64x2(a x86.M256d, b x86.M128d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskInsertf64x2(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M128d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzInsertf64x2(k x86.Mmask8, a x86.M256d, b x86.M128d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Insertf64x2(a x86.M512d, b x86.M128d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskInsertf64x2(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M128d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzInsertf64x2(k x86.Mmask8, a x86.M512d, b x86.M128d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Inserti32x8(a x86.M512i, b x86.M256i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskInserti32x8(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M256i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzInserti32x8(k x86.Mmask16, a x86.M512i, b x86.M256i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256Inserti64x2(a x86.M256i, b x86.M128i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskInserti64x2(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M128i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzInserti64x2(k x86.Mmask8, a x86.M256i, b x86.M128i, imm8 byte) (dst x86.M256i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512Inserti64x2(a x86.M512i, b x86.M128i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskInserti64x2(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M128i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzInserti64x2(k x86.Mmask8, a x86.M512i, b x86.M128i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


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
func Movepi32Mask(a x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


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
func M256Movepi32Mask(a x86.M256i) (dst x86.Mmask8) {
	panic("not implemented")
}


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
func M512Movepi32Mask(a x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


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
func Movepi64Mask(a x86.M128i) (dst x86.Mmask8) {
	panic("not implemented")
}


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
func M256Movepi64Mask(a x86.M256i) (dst x86.Mmask8) {
	panic("not implemented")
}


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
func M512Movepi64Mask(a x86.M512i) (dst x86.Mmask8) {
	panic("not implemented")
}


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
func MovmEpi32(k x86.Mmask8) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256MovmEpi32(k x86.Mmask8) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512MovmEpi32(k x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


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
func MovmEpi64(k x86.Mmask8) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256MovmEpi64(k x86.Mmask8) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512MovmEpi64(k x86.Mmask8) (dst x86.M512i) {
	panic("not implemented")
}


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
func MaskMulloEpi64(src x86.M128i, k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func MaskzMulloEpi64(k x86.Mmask8, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func MulloEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


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
func M256MaskMulloEpi64(src x86.M256i, k x86.Mmask8, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MaskzMulloEpi64(k x86.Mmask8, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M256MulloEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


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
func M512MaskMulloEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MaskzMulloEpi64(k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


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
func M512MulloEpi64(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


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
func MaskOrPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskzOrPd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func M256MaskOrPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzOrPd(k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512MaskOrPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzOrPd(k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512OrPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func MaskOrPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskzOrPs(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskOrPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M256MaskzOrPs(k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskOrPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskzOrPs(k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512OrPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskRangePd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzRangePd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func RangePd(a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskRangePd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzRangePd(k x86.Mmask8, a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256RangePd(a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskRangePd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzRangePd(k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512RangePd(a x86.M512d, b x86.M512d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskRangePs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzRangePs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func RangePs(a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskRangePs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzRangePs(k x86.Mmask8, a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256RangePs(a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskRangePs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzRangePs(k x86.Mmask16, a x86.M512, b x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512RangePs(a x86.M512, b x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskRangeRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 byte, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzRangeRoundPd(k x86.Mmask8, a x86.M512d, b x86.M512d, imm8 byte, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512RangeRoundPd(a x86.M512d, b x86.M512d, imm8 byte, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskRangeRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, imm8 byte, rounding int) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzRangeRoundPs(k x86.Mmask16, a x86.M512, b x86.M512, imm8 byte, rounding int) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512RangeRoundPs(a x86.M512, b x86.M512, imm8 byte, rounding int) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskRangeRoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzRangeRoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func RangeRoundSd(a x86.M128d, b x86.M128d, imm8 byte, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskRangeRoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte, rounding int) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzRangeRoundSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte, rounding int) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func RangeRoundSs(a x86.M128, b x86.M128, imm8 byte, rounding int) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskRangeSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzRangeSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskRangeSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzRangeSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskReducePd(src x86.M128d, k x86.Mmask8, a x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzReducePd(k x86.Mmask8, a x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func ReducePd(a x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskReducePd(src x86.M256d, k x86.Mmask8, a x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzReducePd(k x86.Mmask8, a x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256ReducePd(a x86.M256d, imm8 byte) (dst x86.M256d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskReducePd(src x86.M512d, k x86.Mmask8, a x86.M512d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzReducePd(k x86.Mmask8, a x86.M512d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512ReducePd(a x86.M512d, imm8 byte) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskReducePs(src x86.M128, k x86.Mmask8, a x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzReducePs(k x86.Mmask8, a x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func ReducePs(a x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskReducePs(src x86.M256, k x86.Mmask8, a x86.M256, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256MaskzReducePs(k x86.Mmask8, a x86.M256, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M256ReducePs(a x86.M256, imm8 byte) (dst x86.M256) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskReducePs(src x86.M512, k x86.Mmask16, a x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzReducePs(k x86.Mmask16, a x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512ReducePs(a x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskReduceRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, imm8 byte, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzReduceRoundPd(k x86.Mmask8, a x86.M512d, imm8 byte, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512ReduceRoundPd(a x86.M512d, imm8 byte, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskReduceRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, imm8 byte, rounding int) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512MaskzReduceRoundPs(k x86.Mmask16, a x86.M512, imm8 byte, rounding int) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func M512ReduceRoundPs(a x86.M512, imm8 byte, rounding int) (dst x86.M512) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskReduceRoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzReduceRoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func ReduceRoundSd(a x86.M128d, b x86.M128d, imm8 byte, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskReduceRoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte, rounding int) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzReduceRoundSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte, rounding int) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func ReduceRoundSs(a x86.M128, b x86.M128, imm8 byte, rounding int) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskReduceSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzReduceSd(k x86.Mmask8, a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func ReduceSd(a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskReduceSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func MaskzReduceSs(k x86.Mmask8, a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
//
// FIXME: Requires compiler support (has immediate)
func ReduceSs(a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskXorPd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func MaskzXorPd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


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
func M256MaskXorPd(src x86.M256d, k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M256MaskzXorPd(k x86.Mmask8, a x86.M256d, b x86.M256d) (dst x86.M256d) {
	panic("not implemented")
}


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
func M512MaskXorPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512MaskzXorPd(k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func M512XorPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


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
func MaskXorPs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func MaskzXorPs(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


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
func M256MaskXorPs(src x86.M256, k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M256MaskzXorPs(k x86.Mmask8, a x86.M256, b x86.M256) (dst x86.M256) {
	panic("not implemented")
}


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
func M512MaskXorPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512MaskzXorPs(k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


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
func M512XorPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}

