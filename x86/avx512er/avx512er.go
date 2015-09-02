package avx512er

import . "github.com/klauspost/intrinsics/x86"


// Exp2a23Pd: Compute the approximate exponential value of 2 raised to the
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
func Exp2a23Pd(a M512d) M512d {
	return M512d(exp2a23Pd([8]float64(a)))
}

func exp2a23Pd(a [8]float64) [8]float64


// MaskExp2a23Pd: Compute the approximate exponential value of 2 raised to the
// power of packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). The maximum relative
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
func MaskExp2a23Pd(a M512d, k Mmask8, src M512d) M512d {
	return M512d(maskExp2a23Pd([8]float64(a), uint8(k), [8]float64(src)))
}

func maskExp2a23Pd(a [8]float64, k uint8, src [8]float64) [8]float64


// MaskzExp2a23Pd: Compute the approximate exponential value of 2 raised to the
// power of packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). The maximum relative error for
// this approximation is less than 2^-23. 
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
func MaskzExp2a23Pd(k Mmask8, a M512d) M512d {
	return M512d(maskzExp2a23Pd(uint8(k), [8]float64(a)))
}

func maskzExp2a23Pd(k uint8, a [8]float64) [8]float64


// Exp2a23Ps: Compute the approximate exponential value of 2 raised to the
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
func Exp2a23Ps(a M512) M512 {
	return M512(exp2a23Ps([16]float32(a)))
}

func exp2a23Ps(a [16]float32) [16]float32


// MaskExp2a23Ps: Compute the approximate exponential value of 2 raised to the
// power of packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). The maximum relative
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
func MaskExp2a23Ps(src M512, k Mmask16, a M512) M512 {
	return M512(maskExp2a23Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskExp2a23Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// MaskzExp2a23Ps: Compute the approximate exponential value of 2 raised to the
// power of packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst' using zeromask 'k' (elements are zeroed out
// when the corresponding mask bit is not set). The maximum relative error for
// this approximation is less than 2^-23. 
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
func MaskzExp2a23Ps(k Mmask16, a M512) M512 {
	return M512(maskzExp2a23Ps(uint16(k), [16]float32(a)))
}

func maskzExp2a23Ps(k uint16, a [16]float32) [16]float32


// Exp2a23RoundPd: Compute the approximate exponential value of 2 raised to the
// power of packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst'. The maximum relative error for this
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
func Exp2a23RoundPd(a M512d, rounding int) M512d {
	return M512d(exp2a23RoundPd([8]float64(a), rounding))
}

func exp2a23RoundPd(a [8]float64, rounding int) [8]float64


// MaskExp2a23RoundPd: Compute the approximate exponential value of 2 raised to
// the power of packed double-precision (64-bit) floating-point elements in
// 'a', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). The maximum relative
// error for this approximation is less than 2^-23. Rounding is done according
// to the 'rounding' parameter, which can be one of:
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
func MaskExp2a23RoundPd(a M512d, k Mmask8, src M512d, rounding int) M512d {
	return M512d(maskExp2a23RoundPd([8]float64(a), uint8(k), [8]float64(src), rounding))
}

func maskExp2a23RoundPd(a [8]float64, k uint8, src [8]float64, rounding int) [8]float64


// MaskzExp2a23RoundPd: Compute the approximate exponential value of 2 raised
// to the power of packed double-precision (64-bit) floating-point elements in
// 'a', and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). The maximum relative error
// for this approximation is less than 2^-23. Rounding is done according to the
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
//				dst[i+63:i] := EXP_2_23_DP(a[i+63:i]);
//			ELSE
//				dst[i+63:i] := 0;
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PD'. Intrinsic: '_mm512_maskz_exp2a23_round_pd'.
// Requires AVX512ER.
func MaskzExp2a23RoundPd(k Mmask8, a M512d, rounding int) M512d {
	return M512d(maskzExp2a23RoundPd(uint8(k), [8]float64(a), rounding))
}

func maskzExp2a23RoundPd(k uint8, a [8]float64, rounding int) [8]float64


// Exp2a23RoundPs: Compute the approximate exponential value of 2 raised to the
// power of packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst'. The maximum relative error for this
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
func Exp2a23RoundPs(a M512, rounding int) M512 {
	return M512(exp2a23RoundPs([16]float32(a), rounding))
}

func exp2a23RoundPs(a [16]float32, rounding int) [16]float32


// MaskExp2a23RoundPs: Compute the approximate exponential value of 2 raised to
// the power of packed single-precision (32-bit) floating-point elements in
// 'a', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). The maximum relative
// error for this approximation is less than 2^-23. Rounding is done according
// to the 'rounding' parameter, which can be one of:
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
func MaskExp2a23RoundPs(src M512, k Mmask16, a M512, rounding int) M512 {
	return M512(maskExp2a23RoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func maskExp2a23RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


// MaskzExp2a23RoundPs: Compute the approximate exponential value of 2 raised
// to the power of packed single-precision (32-bit) floating-point elements in
// 'a', and store the results in 'dst' using zeromask 'k' (elements are zeroed
// out when the corresponding mask bit is not set). The maximum relative error
// for this approximation is less than 2^-23. Rounding is done according to the
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
//				dst[i+31:i] := EXP_2_23_SP(a[i+31:i]);
//			ELSE
//				dst[i*31:i] := 0;
//			FI
//		ENDFOR;
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP2PS'. Intrinsic: '_mm512_maskz_exp2a23_round_ps'.
// Requires AVX512ER.
func MaskzExp2a23RoundPs(k Mmask16, a M512, rounding int) M512 {
	return M512(maskzExp2a23RoundPs(uint16(k), [16]float32(a), rounding))
}

func maskzExp2a23RoundPs(k uint16, a [16]float32, rounding int) [16]float32


// MaskRcp28Pd: Compute the approximate reciprocal of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). The maximum relative error for this approximation is
// less than 2^-28. 
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
func MaskRcp28Pd(src M512d, k Mmask8, a M512d) M512d {
	return M512d(maskRcp28Pd([8]float64(src), uint8(k), [8]float64(a)))
}

func maskRcp28Pd(src [8]float64, k uint8, a [8]float64) [8]float64


// MaskzRcp28Pd: Compute the approximate reciprocal of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). The maximum relative error for this approximation is less than
// 2^-28. 
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
func MaskzRcp28Pd(k Mmask8, a M512d) M512d {
	return M512d(maskzRcp28Pd(uint8(k), [8]float64(a)))
}

func maskzRcp28Pd(k uint8, a [8]float64) [8]float64


// Rcp28Pd: Compute the approximate reciprocal of packed double-precision
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
func Rcp28Pd(a M512d) M512d {
	return M512d(rcp28Pd([8]float64(a)))
}

func rcp28Pd(a [8]float64) [8]float64


// MaskRcp28Ps: Compute the approximate reciprocal of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). The maximum relative error for this approximation is
// less than 2^-28. 
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
func MaskRcp28Ps(src M512, k Mmask16, a M512) M512 {
	return M512(maskRcp28Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskRcp28Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// MaskzRcp28Ps: Compute the approximate reciprocal of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). The maximum relative error for this approximation is less than
// 2^-28. 
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
func MaskzRcp28Ps(k Mmask16, a M512) M512 {
	return M512(maskzRcp28Ps(uint16(k), [16]float32(a)))
}

func maskzRcp28Ps(k uint16, a [16]float32) [16]float32


// Rcp28Ps: Compute the approximate reciprocal of packed single-precision
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
func Rcp28Ps(a M512) M512 {
	return M512(rcp28Ps([16]float32(a)))
}

func rcp28Ps(a [16]float32) [16]float32


// MaskRcp28RoundPd: Compute the approximate reciprocal of packed
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
func MaskRcp28RoundPd(src M512d, k Mmask8, a M512d, rounding int) M512d {
	return M512d(maskRcp28RoundPd([8]float64(src), uint8(k), [8]float64(a), rounding))
}

func maskRcp28RoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64


// MaskzRcp28RoundPd: Compute the approximate reciprocal of packed
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
func MaskzRcp28RoundPd(k Mmask8, a M512d, rounding int) M512d {
	return M512d(maskzRcp28RoundPd(uint8(k), [8]float64(a), rounding))
}

func maskzRcp28RoundPd(k uint8, a [8]float64, rounding int) [8]float64


// Rcp28RoundPd: Compute the approximate reciprocal of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. The
// maximum relative error for this approximation is less than 2^-28. Rounding
// is done according to the 'rounding' parameter, which can be one of:
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
func Rcp28RoundPd(a M512d, rounding int) M512d {
	return M512d(rcp28RoundPd([8]float64(a), rounding))
}

func rcp28RoundPd(a [8]float64, rounding int) [8]float64


// MaskRcp28RoundPs: Compute the approximate reciprocal of packed
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
func MaskRcp28RoundPs(src M512, k Mmask16, a M512, rounding int) M512 {
	return M512(maskRcp28RoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func maskRcp28RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


// MaskzRcp28RoundPs: Compute the approximate reciprocal of packed
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
func MaskzRcp28RoundPs(k Mmask16, a M512, rounding int) M512 {
	return M512(maskzRcp28RoundPs(uint16(k), [16]float32(a), rounding))
}

func maskzRcp28RoundPs(k uint16, a [16]float32, rounding int) [16]float32


// Rcp28RoundPs: Compute the approximate reciprocal of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. The
// maximum relative error for this approximation is less than 2^-28. Rounding
// is done according to the 'rounding' parameter, which can be one of:
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
func Rcp28RoundPs(a M512, rounding int) M512 {
	return M512(rcp28RoundPs([16]float32(a), rounding))
}

func rcp28RoundPs(a [16]float32, rounding int) [16]float32


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
func MaskRcp28RoundSd(src M128d, k Mmask8, a M128d, b M128d, rounding int) M128d {
	return M128d(maskRcp28RoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), rounding))
}

func maskRcp28RoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64


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
func MaskzRcp28RoundSd(k Mmask8, a M128d, b M128d, rounding int) M128d {
	return M128d(maskzRcp28RoundSd(uint8(k), [2]float64(a), [2]float64(b), rounding))
}

func maskzRcp28RoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64


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
func Rcp28RoundSd(a M128d, b M128d, rounding int) M128d {
	return M128d(rcp28RoundSd([2]float64(a), [2]float64(b), rounding))
}

func rcp28RoundSd(a [2]float64, b [2]float64, rounding int) [2]float64


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
func MaskRcp28RoundSs(src M128, k Mmask8, a M128, b M128, rounding int) M128 {
	return M128(maskRcp28RoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), rounding))
}

func maskRcp28RoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32


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
func MaskzRcp28RoundSs(k Mmask8, a M128, b M128, rounding int) M128 {
	return M128(maskzRcp28RoundSs(uint8(k), [4]float32(a), [4]float32(b), rounding))
}

func maskzRcp28RoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32


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
func Rcp28RoundSs(a M128, b M128, rounding int) M128 {
	return M128(rcp28RoundSs([4]float32(a), [4]float32(b), rounding))
}

func rcp28RoundSs(a [4]float32, b [4]float32, rounding int) [4]float32


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
func MaskRcp28Sd(src M128d, k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskRcp28Sd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskRcp28Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


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
func MaskzRcp28Sd(k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskzRcp28Sd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzRcp28Sd(k uint8, a [2]float64, b [2]float64) [2]float64


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
func Rcp28Sd(a M128d, b M128d) M128d {
	return M128d(rcp28Sd([2]float64(a), [2]float64(b)))
}

func rcp28Sd(a [2]float64, b [2]float64) [2]float64


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
func MaskRcp28Ss(src M128, k Mmask8, a M128, b M128) M128 {
	return M128(maskRcp28Ss([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskRcp28Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


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
func MaskzRcp28Ss(k Mmask8, a M128, b M128) M128 {
	return M128(maskzRcp28Ss(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzRcp28Ss(k uint8, a [4]float32, b [4]float32) [4]float32


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
func Rcp28Ss(a M128, b M128) M128 {
	return M128(rcp28Ss([4]float32(a), [4]float32(b)))
}

func rcp28Ss(a [4]float32, b [4]float32) [4]float32


// MaskRsqrt28Pd: Compute the approximate reciprocal square root of packed
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
func MaskRsqrt28Pd(src M512d, k Mmask8, a M512d) M512d {
	return M512d(maskRsqrt28Pd([8]float64(src), uint8(k), [8]float64(a)))
}

func maskRsqrt28Pd(src [8]float64, k uint8, a [8]float64) [8]float64


// MaskzRsqrt28Pd: Compute the approximate reciprocal square root of packed
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
func MaskzRsqrt28Pd(k Mmask8, a M512d) M512d {
	return M512d(maskzRsqrt28Pd(uint8(k), [8]float64(a)))
}

func maskzRsqrt28Pd(k uint8, a [8]float64) [8]float64


// Rsqrt28Pd: Compute the approximate reciprocal square root of packed
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
func Rsqrt28Pd(a M512d) M512d {
	return M512d(rsqrt28Pd([8]float64(a)))
}

func rsqrt28Pd(a [8]float64) [8]float64


// MaskRsqrt28Ps: Compute the approximate reciprocal square root of packed
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
func MaskRsqrt28Ps(src M512, k Mmask16, a M512) M512 {
	return M512(maskRsqrt28Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskRsqrt28Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// MaskzRsqrt28Ps: Compute the approximate reciprocal square root of packed
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
func MaskzRsqrt28Ps(k Mmask16, a M512) M512 {
	return M512(maskzRsqrt28Ps(uint16(k), [16]float32(a)))
}

func maskzRsqrt28Ps(k uint16, a [16]float32) [16]float32


// Rsqrt28Ps: Compute the approximate reciprocal square root of packed
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
func Rsqrt28Ps(a M512) M512 {
	return M512(rsqrt28Ps([16]float32(a)))
}

func rsqrt28Ps(a [16]float32) [16]float32


// MaskRsqrt28RoundPd: Compute the approximate reciprocal square root of packed
// double-precision (64-bit) floating-point elements in 'a', store the results
// in 'dst' using writemask 'k' (elements are copied from 'src' when the
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
//				dst[i+63:i] := src[i+63:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PD'. Intrinsic: '_mm512_mask_rsqrt28_round_pd'.
// Requires AVX512ER.
func MaskRsqrt28RoundPd(src M512d, k Mmask8, a M512d, rounding int) M512d {
	return M512d(maskRsqrt28RoundPd([8]float64(src), uint8(k), [8]float64(a), rounding))
}

func maskRsqrt28RoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64


// MaskzRsqrt28RoundPd: Compute the approximate reciprocal square root of
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
func MaskzRsqrt28RoundPd(k Mmask8, a M512d, rounding int) M512d {
	return M512d(maskzRsqrt28RoundPd(uint8(k), [8]float64(a), rounding))
}

func maskzRsqrt28RoundPd(k uint8, a [8]float64, rounding int) [8]float64


// Rsqrt28RoundPd: Compute the approximate reciprocal square root of packed
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
func Rsqrt28RoundPd(a M512d, rounding int) M512d {
	return M512d(rsqrt28RoundPd([8]float64(a), rounding))
}

func rsqrt28RoundPd(a [8]float64, rounding int) [8]float64


// MaskRsqrt28RoundPs: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', store the results
// in 'dst' using writemask 'k' (elements are copied from 'src' when the
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
//				dst[i+31:i] := src[i+31:i];
//			FI
//		ENDFOR;
//
// Instruction: 'VRSQRT28PS'. Intrinsic: '_mm512_mask_rsqrt28_round_ps'.
// Requires AVX512ER.
func MaskRsqrt28RoundPs(src M512, k Mmask16, a M512, rounding int) M512 {
	return M512(maskRsqrt28RoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func maskRsqrt28RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


// MaskzRsqrt28RoundPs: Compute the approximate reciprocal square root of
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
func MaskzRsqrt28RoundPs(k Mmask16, a M512, rounding int) M512 {
	return M512(maskzRsqrt28RoundPs(uint16(k), [16]float32(a), rounding))
}

func maskzRsqrt28RoundPs(k uint16, a [16]float32, rounding int) [16]float32


// Rsqrt28RoundPs: Compute the approximate reciprocal square root of packed
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
func Rsqrt28RoundPs(a M512, rounding int) M512 {
	return M512(rsqrt28RoundPs([16]float32(a), rounding))
}

func rsqrt28RoundPs(a [16]float32, rounding int) [16]float32


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
func MaskRsqrt28RoundSd(src M128d, k Mmask8, a M128d, b M128d, rounding int) M128d {
	return M128d(maskRsqrt28RoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), rounding))
}

func maskRsqrt28RoundSd(src [2]float64, k uint8, a [2]float64, b [2]float64, rounding int) [2]float64


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
func MaskzRsqrt28RoundSd(k Mmask8, a M128d, b M128d, rounding int) M128d {
	return M128d(maskzRsqrt28RoundSd(uint8(k), [2]float64(a), [2]float64(b), rounding))
}

func maskzRsqrt28RoundSd(k uint8, a [2]float64, b [2]float64, rounding int) [2]float64


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
func Rsqrt28RoundSd(a M128d, b M128d, rounding int) M128d {
	return M128d(rsqrt28RoundSd([2]float64(a), [2]float64(b), rounding))
}

func rsqrt28RoundSd(a [2]float64, b [2]float64, rounding int) [2]float64


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
func MaskRsqrt28RoundSs(src M128, k Mmask8, a M128, b M128, rounding int) M128 {
	return M128(maskRsqrt28RoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), rounding))
}

func maskRsqrt28RoundSs(src [4]float32, k uint8, a [4]float32, b [4]float32, rounding int) [4]float32


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
func MaskzRsqrt28RoundSs(k Mmask8, a M128, b M128, rounding int) M128 {
	return M128(maskzRsqrt28RoundSs(uint8(k), [4]float32(a), [4]float32(b), rounding))
}

func maskzRsqrt28RoundSs(k uint8, a [4]float32, b [4]float32, rounding int) [4]float32


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
func Rsqrt28RoundSs(a M128, b M128, rounding int) M128 {
	return M128(rsqrt28RoundSs([4]float32(a), [4]float32(b), rounding))
}

func rsqrt28RoundSs(a [4]float32, b [4]float32, rounding int) [4]float32


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
func MaskRsqrt28Sd(src M128d, k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskRsqrt28Sd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
}

func maskRsqrt28Sd(src [2]float64, k uint8, a [2]float64, b [2]float64) [2]float64


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
func MaskzRsqrt28Sd(k Mmask8, a M128d, b M128d) M128d {
	return M128d(maskzRsqrt28Sd(uint8(k), [2]float64(a), [2]float64(b)))
}

func maskzRsqrt28Sd(k uint8, a [2]float64, b [2]float64) [2]float64


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
func Rsqrt28Sd(a M128d, b M128d) M128d {
	return M128d(rsqrt28Sd([2]float64(a), [2]float64(b)))
}

func rsqrt28Sd(a [2]float64, b [2]float64) [2]float64


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
func MaskRsqrt28Ss(src M128, k Mmask8, a M128, b M128) M128 {
	return M128(maskRsqrt28Ss([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
}

func maskRsqrt28Ss(src [4]float32, k uint8, a [4]float32, b [4]float32) [4]float32


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
func MaskzRsqrt28Ss(k Mmask8, a M128, b M128) M128 {
	return M128(maskzRsqrt28Ss(uint8(k), [4]float32(a), [4]float32(b)))
}

func maskzRsqrt28Ss(k uint8, a [4]float32, b [4]float32) [4]float32


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
func Rsqrt28Ss(a M128, b M128) M128 {
	return M128(rsqrt28Ss([4]float32(a), [4]float32(b)))
}

func rsqrt28Ss(a [4]float32, b [4]float32) [4]float32

