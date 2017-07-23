// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package avx512er

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// M512Exp2a23Pd: Compute the approximate exponential value of 2 raised to the
// power of packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst'. The maximum relative error for this
// approximation is less than 2^-23. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_exp2a23_pd'.
// Requires AVX512ER.
func M512Exp2a23Pd(a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskExp2a23Pd: Compute the approximate exponential value of 2 raised to
// the power of packed double-precision (64-bit) floating-point elements in
// 'a', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). The maximum relative
// error for this approximation is less than 2^-23. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//			ELSE
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_mask_exp2a23_pd'.
// Requires AVX512ER.
func M512MaskExp2a23Pd(a x86.M512d, k x86.Mmask8, src x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskzExp2a23Pd: Compute the approximate exponential value of 2 raised to
// the power of packed double-precision (64-bit) floating-point elements in
// 'a', and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). The maximum relative error
// for this approximation is less than 2^-23. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_maskz_exp2a23_pd'.
// Requires AVX512ER.
func M512MaskzExp2a23Pd(k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Exp2a23Ps: Compute the approximate exponential value of 2 raised to the
// power of packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst'. The maximum relative error for this
// approximation is less than 2^-23. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_exp2a23_ps'.
// Requires AVX512ER.
func M512Exp2a23Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskExp2a23Ps: Compute the approximate exponential value of 2 raised to
// the power of packed single-precision (32-bit) floating-point elements in
// 'a', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). The maximum relative
// error for this approximation is less than 2^-23. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//			ELSE
//				dst[i*31:i] := src[i*31:i];
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_mask_exp2a23_ps'.
// Requires AVX512ER.
func M512MaskExp2a23Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskzExp2a23Ps: Compute the approximate exponential value of 2 raised to
// the power of packed single-precision (32-bit) floating-point elements in
// 'a', and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). The maximum relative error
// for this approximation is less than 2^-23. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//			ELSE
//				dst[i*31:i] := 0;
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_maskz_exp2a23_ps'.
// Requires AVX512ER.
func M512MaskzExp2a23Ps(k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Exp2a23RoundPd: Compute the approximate exponential value of 2 raised to
// the power of packed double-precision (64-bit) floating-point elements in
// 'a', and store the results in 'dst'. The maximum relative error for this
// approximation is less than 2^-23. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_exp2a23_round_pd'.
// Requires AVX512ER.
func M512Exp2a23RoundPd(a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskExp2a23RoundPd: Compute the approximate exponential value of 2
// raised to the power of packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// The maximum relative error for this approximation is less than 2^-23.
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//			ELSE
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_mask_exp2a23_round_pd'.
// Requires AVX512ER.
func M512MaskExp2a23RoundPd(a x86.M512d, k x86.Mmask8, src x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskzExp2a23RoundPd: Compute the approximate exponential value of 2
// raised to the power of packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). The maximum
// relative error for this approximation is less than 2^-23. Rounding is done
// according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_maskz_exp2a23_round_pd'.
// Requires AVX512ER.
func M512MaskzExp2a23RoundPd(k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Exp2a23RoundPs: Compute the approximate exponential value of 2 raised to
// the power of packed single-precision (32-bit) floating-point elements in
// 'a', and store the results in 'dst'. The maximum relative error for this
// approximation is less than 2^-23. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_exp2a23_round_ps'.
// Requires AVX512ER.
func M512Exp2a23RoundPs(a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskExp2a23RoundPs: Compute the approximate exponential value of 2
// raised to the power of packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// The maximum relative error for this approximation is less than 2^-23.
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//			ELSE
//				dst[i*31:i] := src[i*31:i];
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_mask_exp2a23_round_ps'.
// Requires AVX512ER.
func M512MaskExp2a23RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskzExp2a23RoundPs: Compute the approximate exponential value of 2
// raised to the power of packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst' using zeromask 'k' (elements
// are zeroed out when the corresponding mask bit is not set). The maximum
// relative error for this approximation is less than 2^-23. Rounding is done
// according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//			ELSE
//				dst[i*31:i] := 0;
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_maskz_exp2a23_round_ps'.
// Requires AVX512ER.
func M512MaskzExp2a23RoundPs(k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskRcp28Pd: Compute the approximate reciprocal of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := RCP_28_SP(1.0/a[i+63:i];
//			ELSE
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PD'. Intrinsic: '_mm512_mask_rcp28_pd'.
// Requires AVX512ER.
func M512MaskRcp28Pd(src x86.M512d, k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskzRcp28Pd: Compute the approximate reciprocal of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := RCP_28_SP(1.0/a[i+63:i];
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PD'. Intrinsic: '_mm512_maskz_rcp28_pd'.
// Requires AVX512ER.
func M512MaskzRcp28Pd(k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Rcp28Pd: Compute the approximate reciprocal of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. The
// maximum relative error for this approximation is less than 2^-28. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			dst[i+63:i] := RCP_28_SP(1.0/a[i+63:i];
//		ENDFOR;
//
// Instruction: 'VRCP28PD'. Intrinsic: '_mm512_rcp28_pd'.
// Requires AVX512ER.
func M512Rcp28Pd(a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskRcp28Ps: Compute the approximate reciprocal of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := RCP_28_SP(1.0/a[i+31:i];
//			ELSE
//				dst[i+31:i] := src[i+31:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PS'. Intrinsic: '_mm512_mask_rcp28_ps'.
// Requires AVX512ER.
func M512MaskRcp28Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskzRcp28Ps: Compute the approximate reciprocal of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := RCP_28_SP(1.0/a[i+31:i];
//			ELSE
//				dst[i+31:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PS'. Intrinsic: '_mm512_maskz_rcp28_ps'.
// Requires AVX512ER.
func M512MaskzRcp28Ps(k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Rcp28Ps: Compute the approximate reciprocal of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. The
// maximum relative error for this approximation is less than 2^-28. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			dst[i+31:i] := RCP_28_SP(1.0/a[i+31:i];
//		ENDFOR;
//
// Instruction: 'VRCP28PS'. Intrinsic: '_mm512_rcp28_ps'.
// Requires AVX512ER.
func M512Rcp28Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskRcp28RoundPd: Compute the approximate reciprocal of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := RCP_28_SP(1.0/a[i+63:i];
//			ELSE
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PD'. Intrinsic: '_mm512_mask_rcp28_round_pd'.
// Requires AVX512ER.
func M512MaskRcp28RoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskzRcp28RoundPd: Compute the approximate reciprocal of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := RCP_28_SP(1.0/a[i+63:i];
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PD'. Intrinsic: '_mm512_maskz_rcp28_round_pd'.
// Requires AVX512ER.
func M512MaskzRcp28RoundPd(k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Rcp28RoundPd: Compute the approximate reciprocal of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. The maximum relative error for this approximation is less
// than 2^-28. Rounding is done according to the 'rounding' parameter, which
// can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			dst[i+63:i] := RCP_28_SP(1.0/a[i+63:i];
//		ENDFOR;
//
// Instruction: 'VRCP28PD'. Intrinsic: '_mm512_rcp28_round_pd'.
// Requires AVX512ER.
func M512Rcp28RoundPd(a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskRcp28RoundPs: Compute the approximate reciprocal of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := RCP_28_SP(1.0/a[i+31:i];
//			ELSE
//				dst[i+31:i] := src[i+31:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PS'. Intrinsic: '_mm512_mask_rcp28_round_ps'.
// Requires AVX512ER.
func M512MaskRcp28RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskzRcp28RoundPs: Compute the approximate reciprocal of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := RCP_28_SP(1.0/a[i+31:i];
//			ELSE
//				dst[i+31:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRCP28PS'. Intrinsic: '_mm512_maskz_rcp28_round_ps'.
// Requires AVX512ER.
func M512MaskzRcp28RoundPs(k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Rcp28RoundPs: Compute the approximate reciprocal of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. The maximum relative error for this approximation is less
// than 2^-28. Rounding is done according to the 'rounding' parameter, which
// can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			dst[i+31:i] := RCP_28_SP(1.0/a[i+31:i];
//		ENDFOR;
//
// Instruction: 'VRCP28PS'. Intrinsic: '_mm512_rcp28_round_ps'.
// Requires AVX512ER.
func M512Rcp28RoundPs(a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// MaskRcp28RoundSd: Compute the approximate reciprocal of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using writemask 'k' (the element is copied from
// 'src' when mask bit 0 is not set), and copy the upper element from 'a' to
// the upper element of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[63:0] := RCP_28_DP(1.0/b[63:0];
//		ELSE
//			dst[63:0] := src[63:0];
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SD'. Intrinsic: '_mm_mask_rcp28_round_sd'.
// Requires AVX512ER.
func MaskRcp28RoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


// MaskzRcp28RoundSd: Compute the approximate reciprocal of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using zeromask 'k' (the element is zeroed out
// when mask bit 0 is not set), and copy the upper element from 'a' to the
// upper element of 'dst'. The maximum relative error for this approximation is
// less than 2^-28. Rounding is done according to the 'rounding' parameter,
// which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[63:0] := RCP_28_DP(1.0/b[63:0];
//		ELSE
//			dst[63:0] := 0;
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SD'. Intrinsic: '_mm_maskz_rcp28_round_sd'.
// Requires AVX512ER.
func MaskzRcp28RoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


// Rcp28RoundSd: Compute the approximate reciprocal of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. The maximum relative error for this approximation is less
// than 2^-28. Rounding is done according to the 'rounding' parameter, which
// can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		dst[63:0] := RCP_28_DP(1.0/b[63:0];
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SD'. Intrinsic: '_mm_rcp28_round_sd'.
// Requires AVX512ER.
func Rcp28RoundSd(a x86.M128d, b x86.M128d, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


// MaskRcp28RoundSs: Compute the approximate reciprocal of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using writemask 'k' (the element is copied from
// 'src' when mask bit 0 is not set), and copy the upper 3 packed elements from
// 'a' to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[31:0] := RCP_28_DP(1.0/b[31:0];
//		ELSE
//			dst[31:0] := src[31:0];
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SS'. Intrinsic: '_mm_mask_rcp28_round_ss'.
// Requires AVX512ER.
func MaskRcp28RoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, rounding int) (dst x86.M128) {
	panic("not implemented")
}


// MaskzRcp28RoundSs: Compute the approximate reciprocal of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using zeromask 'k' (the element is zeroed out
// when mask bit 0 is not set), and copy the upper 3 packed elements from 'a'
// to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[31:0] := RCP_28_DP(1.0/b[31:0];
//		ELSE
//			dst[31:0] := 0;
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SS'. Intrinsic: '_mm_maskz_rcp28_round_ss'.
// Requires AVX512ER.
func MaskzRcp28RoundSs(k x86.Mmask8, a x86.M128, b x86.M128, rounding int) (dst x86.M128) {
	panic("not implemented")
}


// Rcp28RoundSs: Compute the approximate reciprocal of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst'. The maximum relative error for this
// approximation is less than 2^-28, and copy the upper 3 packed elements from
// 'a' to the upper elements of 'dst'. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		dst[31:0] := RCP_28_DP(1.0/b[31:0];
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SS'. Intrinsic: '_mm_rcp28_round_ss'.
// Requires AVX512ER.
func Rcp28RoundSs(a x86.M128, b x86.M128, rounding int) (dst x86.M128) {
	panic("not implemented")
}


// MaskRcp28Sd: Compute the approximate reciprocal of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using writemask 'k' (the element is copied from
// 'src' when mask bit 0 is not set), and copy the upper element from 'a' to
// the upper element of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		IF k[0] THEN
//			dst[63:0] := RCP_28_DP(1.0/b[63:0];
//		ELSE
//			dst[63:0] := src[63:0];
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SD'. Intrinsic: '_mm_mask_rcp28_sd'.
// Requires AVX512ER.
func MaskRcp28Sd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// MaskzRcp28Sd: Compute the approximate reciprocal of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using zeromask 'k' (the element is zeroed out
// when mask bit 0 is not set), and copy the upper element from 'a' to the
// upper element of 'dst'. The maximum relative error for this approximation is
// less than 2^-28. 
//
//		IF k[0] THEN
//			dst[63:0] := RCP_28_DP(1.0/b[63:0];
//		ELSE
//			dst[63:0] := 0;
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SD'. Intrinsic: '_mm_maskz_rcp28_sd'.
// Requires AVX512ER.
func MaskzRcp28Sd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Rcp28Sd: Compute the approximate reciprocal of the lower double-precision
// (64-bit) floating-point element in 'b', store the result in the lower
// element of 'dst', and copy the upper element from 'a' to the upper element
// of 'dst'. The maximum relative error for this approximation is less than
// 2^-28. 
//
//		dst[63:0] := RCP_28_DP(1.0/b[63:0];
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SD'. Intrinsic: '_mm_rcp28_sd'.
// Requires AVX512ER.
func Rcp28Sd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// MaskRcp28Ss: Compute the approximate reciprocal of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using writemask 'k' (the element is copied from
// 'src' when mask bit 0 is not set), and copy the upper 3 packed elements from
// 'a' to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		IF k[0] THEN
//			dst[31:0] := RCP_28_DP(1.0/b[31:0];
//		ELSE
//			dst[31:0] := src[31:0];
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SS'. Intrinsic: '_mm_mask_rcp28_ss'.
// Requires AVX512ER.
func MaskRcp28Ss(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MaskzRcp28Ss: Compute the approximate reciprocal of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using zeromask 'k' (the element is zeroed out
// when mask bit 0 is not set), and copy the upper 3 packed elements from 'a'
// to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		IF k[0] THEN
//			dst[31:0] := RCP_28_DP(1.0/b[31:0];
//		ELSE
//			dst[31:0] := 0;
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SS'. Intrinsic: '_mm_maskz_rcp28_ss'.
// Requires AVX512ER.
func MaskzRcp28Ss(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Rcp28Ss: Compute the approximate reciprocal of the lower single-precision
// (32-bit) floating-point element in 'b', store the result in the lower
// element of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. The maximum relative error for this approximation is less
// than 2^-28. 
//
//		dst[31:0] := RCP_28_DP(1.0/b[31:0];
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRCP28SS'. Intrinsic: '_mm_rcp28_ss'.
// Requires AVX512ER.
func Rcp28Ss(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// M512MaskRsqrt28Pd: Compute the approximate reciprocal square root of packed
// double-precision (64-bit) floating-point elements in 'a', store the results
// in 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := (1.0/SQRT(a[i+63:i]));
//			ELSE
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_mask_rsqrt28_pd'.
// Requires AVX512ER.
func M512MaskRsqrt28Pd(src x86.M512d, k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskzRsqrt28Pd: Compute the approximate reciprocal square root of packed
// double-precision (64-bit) floating-point elements in 'a', store the results
// in 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). The maximum relative error for this approximation is
// less than 2^-28. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := (1.0/SQRT(a[i+63:i]));
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_maskz_rsqrt28_pd'.
// Requires AVX512ER.
func M512MaskzRsqrt28Pd(k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Rsqrt28Pd: Compute the approximate reciprocal square root of packed
// double-precision (64-bit) floating-point elements in 'a', store the results
// in 'dst'. The maximum relative error for this approximation is less than
// 2^-28. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			dst[i+63:i] := (1.0/SQRT(a[i+63:i]));
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_rsqrt28_pd'.
// Requires AVX512ER.
func M512Rsqrt28Pd(a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskRsqrt28Ps: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', store the results
// in 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := (1.0/SQRT(a[i+31:i]));
//			ELSE
//				dst[i+31:i] := src[i+31:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_mask_rsqrt28_ps'.
// Requires AVX512ER.
func M512MaskRsqrt28Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskzRsqrt28Ps: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', store the results
// in 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). The maximum relative error for this approximation is
// less than 2^-28. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := (1.0/SQRT(a[i+31:i]));
//			ELSE
//				dst[i+31:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_maskz_rsqrt28_ps'.
// Requires AVX512ER.
func M512MaskzRsqrt28Ps(k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Rsqrt28Ps: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', store the results
// in 'dst'. The maximum relative error for this approximation is less than
// 2^-28. 
//
//		FOR j := 0 to 15
//			i := j*32;
//			dst[i+31:i] := (1.0/SQRT(a[i+31:i]));
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_rsqrt28_ps'.
// Requires AVX512ER.
func M512Rsqrt28Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskRsqrt28RoundPd: Compute the approximate reciprocal square root of
// packed double-precision (64-bit) floating-point elements in 'a', store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := (1.0/SQRT(a[i+63:i]));
//			ELSE
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_mask_rsqrt28_round_pd'.
// Requires AVX512ER.
func M512MaskRsqrt28RoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskzRsqrt28RoundPd: Compute the approximate reciprocal square root of
// packed double-precision (64-bit) floating-point elements in 'a', store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF k[j] THEN
//				dst[i+63:i] := (1.0/SQRT(a[i+63:i]));
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_maskz_rsqrt28_round_pd'.
// Requires AVX512ER.
func M512MaskzRsqrt28RoundPd(k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Rsqrt28RoundPd: Compute the approximate reciprocal square root of packed
// double-precision (64-bit) floating-point elements in 'a', store the results
// in 'dst'. The maximum relative error for this approximation is less than
// 2^-28. Rounding is done according to the 'rounding' parameter, which can be
// one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64;
//			dst[i+63:i] := (1.0/SQRT(a[i+63:i]));
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_rsqrt28_round_pd'.
// Requires AVX512ER.
func M512Rsqrt28RoundPd(a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskRsqrt28RoundPs: Compute the approximate reciprocal square root of
// packed single-precision (32-bit) floating-point elements in 'a', store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := (1.0/SQRT(a[i+31:i]));
//			ELSE
//				dst[i+31:i] := src[i+31:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_mask_rsqrt28_round_ps'.
// Requires AVX512ER.
func M512MaskRsqrt28RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskzRsqrt28RoundPs: Compute the approximate reciprocal square root of
// packed single-precision (32-bit) floating-point elements in 'a', store the
// results in 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			IF k[j] THEN
//				dst[i+31:i] := (1.0/SQRT(a[i+31:i]));
//			ELSE
//				dst[i+31:i] := 0;
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_maskz_rsqrt28_round_ps'.
// Requires AVX512ER.
func M512MaskzRsqrt28RoundPs(k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Rsqrt28RoundPs: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', store the results
// in 'dst'. The maximum relative error for this approximation is less than
// 2^-28. Rounding is done according to the 'rounding' parameter, which can be
// one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32;
//			dst[i+31:i] := (1.0/SQRT(a[i+31:i]));
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_rsqrt28_round_ps'.
// Requires AVX512ER.
func M512Rsqrt28RoundPs(a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// MaskRsqrt28RoundSd: Compute the approximate reciprocal square root of the
// lower double-precision (64-bit) floating-point element in 'b', store the
// result in the lower element of 'dst' using writemask 'k' (the element is
// copied from 'src' when mask bit 0 is not set), and copy the upper element
// from 'a' to the upper element of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[63:0] := (1.0/SQRT(b[63:0]));
//		ELSE
//			dst[63:0] := src[63:0];
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SD'. Intrinsic: '_mm_mask_rsqrt28_round_sd'.
// Requires AVX512ER.
func MaskRsqrt28RoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


// MaskzRsqrt28RoundSd: Compute the approximate reciprocal square root of the
// lower double-precision (64-bit) floating-point element in 'b', store the
// result in the lower element of 'dst' using zeromask 'k' (the element is
// zeroed out when mask bit 0 is not set), and copy the upper element from 'a'
// to the upper element of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[63:0] := (1.0/SQRT(b[63:0]));
//		ELSE
//			dst[63:0] := 0;
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SD'. Intrinsic: '_mm_maskz_rsqrt28_round_sd'.
// Requires AVX512ER.
func MaskzRsqrt28RoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


// Rsqrt28RoundSd: Compute the approximate reciprocal square root of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. The maximum relative error for this approximation is less
// than 2^-28. Rounding is done according to the 'rounding' parameter, which
// can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		dst[63:0] := (1.0/SQRT(b[63:0]));
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SD'. Intrinsic: '_mm_rsqrt28_round_sd'.
// Requires AVX512ER.
func Rsqrt28RoundSd(a x86.M128d, b x86.M128d, rounding int) (dst x86.M128d) {
	panic("not implemented")
}


// MaskRsqrt28RoundSs: Compute the approximate reciprocal square root of the
// lower single-precision (32-bit) floating-point element in 'b', store the
// result in the lower element of 'dst' using writemask 'k' (the element is
// copied from 'src' when mask bit 0 is not set), and copy the upper 3 packed
// elements from 'a' to the upper elements of 'dst'. The maximum relative error
// for this approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[31:0] := (1.0/SQRT(b[31:0]));
//		ELSE
//			dst[31:0] := src[31:0];
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SS'. Intrinsic: '_mm_mask_rsqrt28_round_ss'.
// Requires AVX512ER.
func MaskRsqrt28RoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, rounding int) (dst x86.M128) {
	panic("not implemented")
}


// MaskzRsqrt28RoundSs: Compute the approximate reciprocal square root of the
// lower single-precision (32-bit) floating-point element in 'b', store the
// result in the lower element of 'dst' using zeromask 'k' (the element is
// zeroed out when mask bit 0 is not set), and copy the upper 3 packed elements
// from 'a' to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		IF k[0] THEN
//			dst[31:0] := (1.0/SQRT(b[31:0]));
//		ELSE
//			dst[31:0] := 0;
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SS'. Intrinsic: '_mm_maskz_rsqrt28_round_ss'.
// Requires AVX512ER.
func MaskzRsqrt28RoundSs(k x86.Mmask8, a x86.M128, b x86.M128, rounding int) (dst x86.M128) {
	panic("not implemented")
}


// Rsqrt28RoundSs: Compute the approximate reciprocal square root of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst', and copy the upper 3 packed elements from 'a' to
// the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. Rounding is done according to the
// 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		dst[31:0] := (1.0/SQRT(b[31:0]));
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SS'. Intrinsic: '_mm_rsqrt28_round_ss'.
// Requires AVX512ER.
func Rsqrt28RoundSs(a x86.M128, b x86.M128, rounding int) (dst x86.M128) {
	panic("not implemented")
}


// MaskRsqrt28Sd: Compute the approximate reciprocal square root of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using writemask 'k' (the element is copied from
// 'src' when mask bit 0 is not set), and copy the upper element from 'a' to
// the upper element of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		IF k[0] THEN
//			dst[63:0] := (1.0/SQRT(b[63:0]));
//		ELSE
//			dst[63:0] := src[63:0];
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SD'. Intrinsic: '_mm_mask_rsqrt28_sd'.
// Requires AVX512ER.
func MaskRsqrt28Sd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// MaskzRsqrt28Sd: Compute the approximate reciprocal square root of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using zeromask 'k' (the element is zeroed out
// when mask bit 0 is not set), and copy the upper element from 'a' to the
// upper element of 'dst'. The maximum relative error for this approximation is
// less than 2^-28. 
//
//		IF k[0] THEN
//			dst[63:0] := (1.0/SQRT(b[63:0]));
//		ELSE
//			dst[63:0] := 0;
//		FI
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SD'. Intrinsic: '_mm_maskz_rsqrt28_sd'.
// Requires AVX512ER.
func MaskzRsqrt28Sd(k x86.Mmask8, a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Rsqrt28Sd: Compute the approximate reciprocal square root of the lower
// double-precision (64-bit) floating-point element in 'b', store the result in
// the lower element of 'dst', and copy the upper element from 'a' to the upper
// element of 'dst'. The maximum relative error for this approximation is less
// than 2^-28. 
//
//		dst[63:0] := (1.0/SQRT(b[63:0]));
//		dst[127:64] := a[127:64];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SD'. Intrinsic: '_mm_rsqrt28_sd'.
// Requires AVX512ER.
func Rsqrt28Sd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// MaskRsqrt28Ss: Compute the approximate reciprocal square root of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using writemask 'k' (the element is copied from
// 'src' when mask bit 0 is not set), and copy the upper 3 packed elements from
// 'a' to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		IF k[0] THEN
//			dst[31:0] := (1.0/SQRT(b[31:0]));
//		ELSE
//			dst[31:0] := src[31:0];
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SS'. Intrinsic: '_mm_mask_rsqrt28_ss'.
// Requires AVX512ER.
func MaskRsqrt28Ss(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MaskzRsqrt28Ss: Compute the approximate reciprocal square root of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst' using zeromask 'k' (the element is zeroed out
// when mask bit 0 is not set), and copy the upper 3 packed elements from 'a'
// to the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		IF k[0] THEN
//			dst[31:0] := (1.0/SQRT(b[31:0]));
//		ELSE
//			dst[31:0] := 0;
//		FI
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SS'. Intrinsic: '_mm_maskz_rsqrt28_ss'.
// Requires AVX512ER.
func MaskzRsqrt28Ss(k x86.Mmask8, a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Rsqrt28Ss: Compute the approximate reciprocal square root of the lower
// single-precision (32-bit) floating-point element in 'b', store the result in
// the lower element of 'dst', and copy the upper 3 packed elements from 'a' to
// the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 2^-28. 
//
//		dst[31:0] := (1.0/SQRT(b[31:0]));
//		dst[127:32] := a[127:32];
//		dst[MAX:128] := 0;
//
// Instruction: 'VRSQRT28SS'. Intrinsic: '_mm_rsqrt28_ss'.
// Requires AVX512ER.
func Rsqrt28Ss(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}

