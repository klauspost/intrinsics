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
func M512Exp2a23Pd(a x86.M512d) x86.M512d {
	return x86.M512d(m512Exp2a23Pd([8]float64(a)))
}

func m512Exp2a23Pd(a [8]float64) [8]float64


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
func M512MaskExp2a23Pd(a x86.M512d, k x86.Mmask8, src x86.M512d) x86.M512d {
	return x86.M512d(m512MaskExp2a23Pd([8]float64(a), uint8(k), [8]float64(src)))
}

func m512MaskExp2a23Pd(a [8]float64, k uint8, src [8]float64) [8]float64


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
func M512MaskzExp2a23Pd(k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzExp2a23Pd(uint8(k), [8]float64(a)))
}

func m512MaskzExp2a23Pd(k uint8, a [8]float64) [8]float64


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
func M512Exp2a23Ps(a x86.M512) x86.M512 {
	return x86.M512(m512Exp2a23Ps([16]float32(a)))
}

func m512Exp2a23Ps(a [16]float32) [16]float32


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
func M512MaskExp2a23Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(m512MaskExp2a23Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func m512MaskExp2a23Ps(src [16]float32, k uint16, a [16]float32) [16]float32


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
func M512MaskzExp2a23Ps(k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(m512MaskzExp2a23Ps(uint16(k), [16]float32(a)))
}

func m512MaskzExp2a23Ps(k uint16, a [16]float32) [16]float32


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
func M512Exp2a23RoundPd(a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512Exp2a23RoundPd([8]float64(a), rounding))
}

func m512Exp2a23RoundPd(a [8]float64, rounding int) [8]float64


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
func M512MaskExp2a23RoundPd(a x86.M512d, k x86.Mmask8, src x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512MaskExp2a23RoundPd([8]float64(a), uint8(k), [8]float64(src), rounding))
}

func m512MaskExp2a23RoundPd(a [8]float64, k uint8, src [8]float64, rounding int) [8]float64


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
func M512MaskzExp2a23RoundPd(k x86.Mmask8, a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512MaskzExp2a23RoundPd(uint8(k), [8]float64(a), rounding))
}

func m512MaskzExp2a23RoundPd(k uint8, a [8]float64, rounding int) [8]float64


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
func M512Exp2a23RoundPs(a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512Exp2a23RoundPs([16]float32(a), rounding))
}

func m512Exp2a23RoundPs(a [16]float32, rounding int) [16]float32


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
func M512MaskExp2a23RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512MaskExp2a23RoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func m512MaskExp2a23RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


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
func M512MaskzExp2a23RoundPs(k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512MaskzExp2a23RoundPs(uint16(k), [16]float32(a), rounding))
}

func m512MaskzExp2a23RoundPs(k uint16, a [16]float32, rounding int) [16]float32


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
func M512MaskRcp28Pd(src x86.M512d, k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(m512MaskRcp28Pd([8]float64(src), uint8(k), [8]float64(a)))
}

func m512MaskRcp28Pd(src [8]float64, k uint8, a [8]float64) [8]float64


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
func M512MaskzRcp28Pd(k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzRcp28Pd(uint8(k), [8]float64(a)))
}

func m512MaskzRcp28Pd(k uint8, a [8]float64) [8]float64


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
func M512Rcp28Pd(a x86.M512d) x86.M512d {
	return x86.M512d(m512Rcp28Pd([8]float64(a)))
}

func m512Rcp28Pd(a [8]float64) [8]float64


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
func M512MaskRcp28Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(m512MaskRcp28Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func m512MaskRcp28Ps(src [16]float32, k uint16, a [16]float32) [16]float32


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
func M512MaskzRcp28Ps(k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(m512MaskzRcp28Ps(uint16(k), [16]float32(a)))
}

func m512MaskzRcp28Ps(k uint16, a [16]float32) [16]float32


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
func M512Rcp28Ps(a x86.M512) x86.M512 {
	return x86.M512(m512Rcp28Ps([16]float32(a)))
}

func m512Rcp28Ps(a [16]float32) [16]float32


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
func M512MaskRcp28RoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512MaskRcp28RoundPd([8]float64(src), uint8(k), [8]float64(a), rounding))
}

func m512MaskRcp28RoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64


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
func M512MaskzRcp28RoundPd(k x86.Mmask8, a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512MaskzRcp28RoundPd(uint8(k), [8]float64(a), rounding))
}

func m512MaskzRcp28RoundPd(k uint8, a [8]float64, rounding int) [8]float64


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
func M512Rcp28RoundPd(a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512Rcp28RoundPd([8]float64(a), rounding))
}

func m512Rcp28RoundPd(a [8]float64, rounding int) [8]float64


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
func M512MaskRcp28RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512MaskRcp28RoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func m512MaskRcp28RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


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
func M512MaskzRcp28RoundPs(k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512MaskzRcp28RoundPs(uint16(k), [16]float32(a), rounding))
}

func m512MaskzRcp28RoundPs(k uint16, a [16]float32, rounding int) [16]float32


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
func M512Rcp28RoundPs(a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512Rcp28RoundPs([16]float32(a), rounding))
}

func m512Rcp28RoundPs(a [16]float32, rounding int) [16]float32


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
func MaskRcp28RoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(maskRcp28RoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), rounding))
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
func MaskzRcp28RoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(maskzRcp28RoundSd(uint8(k), [2]float64(a), [2]float64(b), rounding))
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
func Rcp28RoundSd(a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(rcp28RoundSd([2]float64(a), [2]float64(b), rounding))
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
func MaskRcp28RoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(maskRcp28RoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), rounding))
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
func MaskzRcp28RoundSs(k x86.Mmask8, a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(maskzRcp28RoundSs(uint8(k), [4]float32(a), [4]float32(b), rounding))
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
func Rcp28RoundSs(a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(rcp28RoundSs([4]float32(a), [4]float32(b), rounding))
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
func MaskRcp28Sd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskRcp28Sd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
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
func MaskzRcp28Sd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzRcp28Sd(uint8(k), [2]float64(a), [2]float64(b)))
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
func Rcp28Sd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(rcp28Sd([2]float64(a), [2]float64(b)))
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
func MaskRcp28Ss(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskRcp28Ss([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
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
func MaskzRcp28Ss(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzRcp28Ss(uint8(k), [4]float32(a), [4]float32(b)))
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
func Rcp28Ss(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(rcp28Ss([4]float32(a), [4]float32(b)))
}

func rcp28Ss(a [4]float32, b [4]float32) [4]float32


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
func M512MaskRsqrt28Pd(src x86.M512d, k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(m512MaskRsqrt28Pd([8]float64(src), uint8(k), [8]float64(a)))
}

func m512MaskRsqrt28Pd(src [8]float64, k uint8, a [8]float64) [8]float64


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
func M512MaskzRsqrt28Pd(k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(m512MaskzRsqrt28Pd(uint8(k), [8]float64(a)))
}

func m512MaskzRsqrt28Pd(k uint8, a [8]float64) [8]float64


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
func M512Rsqrt28Pd(a x86.M512d) x86.M512d {
	return x86.M512d(m512Rsqrt28Pd([8]float64(a)))
}

func m512Rsqrt28Pd(a [8]float64) [8]float64


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
func M512MaskRsqrt28Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(m512MaskRsqrt28Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func m512MaskRsqrt28Ps(src [16]float32, k uint16, a [16]float32) [16]float32


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
func M512MaskzRsqrt28Ps(k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(m512MaskzRsqrt28Ps(uint16(k), [16]float32(a)))
}

func m512MaskzRsqrt28Ps(k uint16, a [16]float32) [16]float32


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
func M512Rsqrt28Ps(a x86.M512) x86.M512 {
	return x86.M512(m512Rsqrt28Ps([16]float32(a)))
}

func m512Rsqrt28Ps(a [16]float32) [16]float32


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
func M512MaskRsqrt28RoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512MaskRsqrt28RoundPd([8]float64(src), uint8(k), [8]float64(a), rounding))
}

func m512MaskRsqrt28RoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64


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
func M512MaskzRsqrt28RoundPd(k x86.Mmask8, a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512MaskzRsqrt28RoundPd(uint8(k), [8]float64(a), rounding))
}

func m512MaskzRsqrt28RoundPd(k uint8, a [8]float64, rounding int) [8]float64


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
func M512Rsqrt28RoundPd(a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(m512Rsqrt28RoundPd([8]float64(a), rounding))
}

func m512Rsqrt28RoundPd(a [8]float64, rounding int) [8]float64


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
func M512MaskRsqrt28RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512MaskRsqrt28RoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func m512MaskRsqrt28RoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


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
func M512MaskzRsqrt28RoundPs(k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512MaskzRsqrt28RoundPs(uint16(k), [16]float32(a), rounding))
}

func m512MaskzRsqrt28RoundPs(k uint16, a [16]float32, rounding int) [16]float32


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
func M512Rsqrt28RoundPs(a x86.M512, rounding int) x86.M512 {
	return x86.M512(m512Rsqrt28RoundPs([16]float32(a), rounding))
}

func m512Rsqrt28RoundPs(a [16]float32, rounding int) [16]float32


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
func MaskRsqrt28RoundSd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(maskRsqrt28RoundSd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b), rounding))
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
func MaskzRsqrt28RoundSd(k x86.Mmask8, a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(maskzRsqrt28RoundSd(uint8(k), [2]float64(a), [2]float64(b), rounding))
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
func Rsqrt28RoundSd(a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(rsqrt28RoundSd([2]float64(a), [2]float64(b), rounding))
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
func MaskRsqrt28RoundSs(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(maskRsqrt28RoundSs([4]float32(src), uint8(k), [4]float32(a), [4]float32(b), rounding))
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
func MaskzRsqrt28RoundSs(k x86.Mmask8, a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(maskzRsqrt28RoundSs(uint8(k), [4]float32(a), [4]float32(b), rounding))
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
func Rsqrt28RoundSs(a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(rsqrt28RoundSs([4]float32(a), [4]float32(b), rounding))
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
func MaskRsqrt28Sd(src x86.M128d, k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskRsqrt28Sd([2]float64(src), uint8(k), [2]float64(a), [2]float64(b)))
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
func MaskzRsqrt28Sd(k x86.Mmask8, a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(maskzRsqrt28Sd(uint8(k), [2]float64(a), [2]float64(b)))
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
func Rsqrt28Sd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(rsqrt28Sd([2]float64(a), [2]float64(b)))
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
func MaskRsqrt28Ss(src x86.M128, k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskRsqrt28Ss([4]float32(src), uint8(k), [4]float32(a), [4]float32(b)))
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
func MaskzRsqrt28Ss(k x86.Mmask8, a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maskzRsqrt28Ss(uint8(k), [4]float32(a), [4]float32(b)))
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
func Rsqrt28Ss(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(rsqrt28Ss([4]float32(a), [4]float32(b)))
}

func rsqrt28Ss(a [4]float32, b [4]float32) [4]float32

