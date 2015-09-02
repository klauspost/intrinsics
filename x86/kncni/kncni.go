package kncni

import "github.com/klauspost/intrinsics/x86"


// AbsPd: Finds the absolute value of each packed double-precision (64-bit)
// floating-point element in 'v2', storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ABS(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_abs_pd'.
// Requires KNCNI.
func AbsPd(v2 x86.M512d) x86.M512d {
	return x86.M512d(absPd([8]float64(v2)))
}

func absPd(v2 [8]float64) [8]float64


// MaskAbsPd: Finds the absolute value of each packed double-precision (64-bit)
// floating-point element in 'v2', storing the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ABS(v2[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_mask_abs_pd'.
// Requires KNCNI.
func MaskAbsPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d) x86.M512d {
	return x86.M512d(maskAbsPd([8]float64(src), uint8(k), [8]float64(v2)))
}

func maskAbsPd(src [8]float64, k uint8, v2 [8]float64) [8]float64


// AbsPs: Finds the absolute value of each packed single-precision (32-bit)
// floating-point element in 'v2', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ABS(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_abs_ps'.
// Requires KNCNI.
func AbsPs(v2 x86.M512) x86.M512 {
	return x86.M512(absPs([16]float32(v2)))
}

func absPs(v2 [16]float32) [16]float32


// MaskAbsPs: Finds the absolute value of each packed single-precision (32-bit)
// floating-point element in 'v2', storing the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ABS(v2[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_mask_abs_ps'.
// Requires KNCNI.
func MaskAbsPs(src x86.M512, k x86.Mmask16, v2 x86.M512) x86.M512 {
	return x86.M512(maskAbsPs([16]float32(src), uint16(k), [16]float32(v2)))
}

func maskAbsPs(src [16]float32, k uint16, v2 [16]float32) [16]float32


// AdcEpi32: Performs element-by-element addition of packed 32-bit integers in
// 'v2' and 'v3' and the corresponding bit in 'k2', storing the result of the
// addition in 'dst' and the result of the carry in 'k2_res'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k2_res[j]   := Carry(v2[i+31:i] + v3[i+31:i] + k2[j])
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i] + k2[j]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADCD'. Intrinsic: '_mm512_adc_epi32'.
// Requires KNCNI.
func AdcEpi32(v2 x86.M512i, k2 x86.Mmask16, v3 x86.M512i, k2_res x86.Mmask16) x86.M512i {
	return x86.M512i(adcEpi32([64]byte(v2), uint16(k2), [64]byte(v3), uint16(k2_res)))
}

func adcEpi32(v2 [64]byte, k2 uint16, v3 [64]byte, k2_res uint16) [64]byte


// MaskAdcEpi32: Performs element-by-element addition of packed 32-bit integers
// in 'v2' and 'v3' and the corresponding bit in 'k2', storing the result of
// the addition in 'dst' and the result of the carry in 'k2_res' using
// writemask 'k1' (elements are copied from 'v2' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k2_res[j]   := Carry(v2[i+31:i] + v3[i+31:i] + k2[j])
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i] + k2[j]
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADCD'. Intrinsic: '_mm512_mask_adc_epi32'.
// Requires KNCNI.
func MaskAdcEpi32(v2 x86.M512i, k1 x86.Mmask16, k2 x86.Mmask16, v3 x86.M512i, k2_res x86.Mmask16) x86.M512i {
	return x86.M512i(maskAdcEpi32([64]byte(v2), uint16(k1), uint16(k2), [64]byte(v3), uint16(k2_res)))
}

func maskAdcEpi32(v2 [64]byte, k1 uint16, k2 uint16, v3 [64]byte, k2_res uint16) [64]byte


// AddEpi32: Add packed 32-bit integers in 'a' and 'b', and store the results
// in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDD'. Intrinsic: '_mm512_add_epi32'.
// Requires KNCNI.
func AddEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(addEpi32([64]byte(a), [64]byte(b)))
}

func addEpi32(a [64]byte, b [64]byte) [64]byte


// MaskAddEpi32: Add packed 32-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDD'. Intrinsic: '_mm512_mask_add_epi32'.
// Requires KNCNI.
func MaskAddEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskAddEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskAddEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// AddPd: Add packed double-precision (64-bit) floating-point elements in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_add_pd'.
// Requires KNCNI.
func AddPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(addPd([8]float64(a), [8]float64(b)))
}

func addPd(a [8]float64, b [8]float64) [8]float64


// MaskAddPd: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_mask_add_pd'.
// Requires KNCNI.
func MaskAddPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(maskAddPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskAddPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// AddPs: Add packed single-precision (32-bit) floating-point elements in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_add_ps'.
// Requires KNCNI.
func AddPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(addPs([16]float32(a), [16]float32(b)))
}

func addPs(a [16]float32, b [16]float32) [16]float32


// MaskAddPs: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_mask_add_ps'.
// Requires KNCNI.
func MaskAddPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskAddPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskAddPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// AddRoundPd: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'.
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
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_add_round_pd'.
// Requires KNCNI.
func AddRoundPd(a x86.M512d, b x86.M512d, rounding int) x86.M512d {
	return x86.M512d(addRoundPd([8]float64(a), [8]float64(b), rounding))
}

func addRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64


// MaskAddRoundPd: Add packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_mask_add_round_pd'.
// Requires KNCNI.
func MaskAddRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskAddRoundPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), rounding))
}

func maskAddRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64


// AddRoundPs: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_add_round_ps'.
// Requires KNCNI.
func AddRoundPs(a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(addRoundPs([16]float32(a), [16]float32(b), rounding))
}

func addRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32


// MaskAddRoundPs: Add packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_mask_add_round_ps'.
// Requires KNCNI.
func MaskAddRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(maskAddRoundPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), rounding))
}

func maskAddRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32


// AddnPd: Performs element-by-element addition between packed double-precision
// (64-bit) floating-point elements in 'v2' and 'v3' and negates their sum,
// storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_addn_pd'.
// Requires KNCNI.
func AddnPd(v2 x86.M512d, v3 x86.M512d) x86.M512d {
	return x86.M512d(addnPd([8]float64(v2), [8]float64(v3)))
}

func addnPd(v2 [8]float64, v3 [8]float64) [8]float64


// MaskAddnPd: Performs element-by-element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates their sum, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_mask_addn_pd'.
// Requires KNCNI.
func MaskAddnPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d) x86.M512d {
	return x86.M512d(maskAddnPd([8]float64(src), uint8(k), [8]float64(v2), [8]float64(v3)))
}

func maskAddnPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64) [8]float64


// AddnPs: Performs element-by-element addition between packed single-precision
// (32-bit) floating-point elements in 'v2' and 'v3' and negates their sum,
// storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_addn_ps'.
// Requires KNCNI.
func AddnPs(v2 x86.M512, v3 x86.M512) x86.M512 {
	return x86.M512(addnPs([16]float32(v2), [16]float32(v3)))
}

func addnPs(v2 [16]float32, v3 [16]float32) [16]float32


// MaskAddnPs: Performs element-by-element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates their sum, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_mask_addn_ps'.
// Requires KNCNI.
func MaskAddnPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512) x86.M512 {
	return x86.M512(maskAddnPs([16]float32(src), uint16(k), [16]float32(v2), [16]float32(v3)))
}

func maskAddnPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32) [16]float32


// AddnRoundPd: Performs element by element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst'.
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
//			dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_addn_round_pd'.
// Requires KNCNI.
func AddnRoundPd(v2 x86.M512d, v3 x86.M512d, rounding int) x86.M512d {
	return x86.M512d(addnRoundPd([8]float64(v2), [8]float64(v3), rounding))
}

func addnRoundPd(v2 [8]float64, v3 [8]float64, rounding int) [8]float64


// MaskAddnRoundPd: Performs element by element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).
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
//				dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_mask_addn_round_pd'.
// Requires KNCNI.
func MaskAddnRoundPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskAddnRoundPd([8]float64(src), uint8(k), [8]float64(v2), [8]float64(v3), rounding))
}

func maskAddnRoundPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64, rounding int) [8]float64


// AddnRoundPs: Performs element by element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_addn_round_ps'.
// Requires KNCNI.
func AddnRoundPs(v2 x86.M512, v3 x86.M512, rounding int) x86.M512 {
	return x86.M512(addnRoundPs([16]float32(v2), [16]float32(v3), rounding))
}

func addnRoundPs(v2 [16]float32, v3 [16]float32, rounding int) [16]float32


// MaskAddnRoundPs: Performs element by element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_mask_addn_round_ps'.
// Requires KNCNI.
func MaskAddnRoundPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, rounding int) x86.M512 {
	return x86.M512(maskAddnRoundPs([16]float32(src), uint16(k), [16]float32(v2), [16]float32(v3), rounding))
}

func maskAddnRoundPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, rounding int) [16]float32


// AddsetcEpi32: Performs element-by-element addition of packed 32-bit integer
// elements in 'v2' and 'v3', storing the resultant carry in 'k2_res' (carry
// flag) and the addition results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			k2_res[j] := Carry(v2[i+31:i] + v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETCD'. Intrinsic: '_mm512_addsetc_epi32'.
// Requires KNCNI.
func AddsetcEpi32(v2 x86.M512i, v3 x86.M512i, k2_res x86.Mmask16) x86.M512i {
	return x86.M512i(addsetcEpi32([64]byte(v2), [64]byte(v3), uint16(k2_res)))
}

func addsetcEpi32(v2 [64]byte, v3 [64]byte, k2_res uint16) [64]byte


// MaskAddsetcEpi32: Performs element-by-element addition of packed 32-bit
// integer elements in 'v2' and 'v3', storing the resultant carry in 'k2_res'
// (carry flag) and the addition results in 'dst' using writemask 'k' (elements
// are copied from 'v2' and 'k_old' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//				k2_res[j] := k_old[j]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETCD'. Intrinsic: '_mm512_mask_addsetc_epi32'.
// Requires KNCNI.
func MaskAddsetcEpi32(v2 x86.M512i, k x86.Mmask16, k_old x86.Mmask16, v3 x86.M512i, k2_res x86.Mmask16) x86.M512i {
	return x86.M512i(maskAddsetcEpi32([64]byte(v2), uint16(k), uint16(k_old), [64]byte(v3), uint16(k2_res)))
}

func maskAddsetcEpi32(v2 [64]byte, k uint16, k_old uint16, v3 [64]byte, k2_res uint16) [64]byte


// AddsetsEpi32: Performs an element-by-element addition of packed 32-bit
// integer elements in 'v2' and 'v3', storing the results in 'dst' and the sign
// of the sum in 'sign' (sign flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETSD'. Intrinsic: '_mm512_addsets_epi32'.
// Requires KNCNI.
func AddsetsEpi32(v2 x86.M512i, v3 x86.M512i, sign x86.Mmask16) x86.M512i {
	return x86.M512i(addsetsEpi32([64]byte(v2), [64]byte(v3), uint16(sign)))
}

func addsetsEpi32(v2 [64]byte, v3 [64]byte, sign uint16) [64]byte


// MaskAddsetsEpi32: Performs an element-by-element addition of packed 32-bit
// integer elements in 'v2' and 'v3', storing the results in 'dst' and the sign
// of the sum in 'sign' (sign flag). Results are stored using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//				sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETSD'. Intrinsic: '_mm512_mask_addsets_epi32'.
// Requires KNCNI.
func MaskAddsetsEpi32(src x86.M512i, k x86.Mmask16, v2 x86.M512i, v3 x86.M512i, sign x86.Mmask16) x86.M512i {
	return x86.M512i(maskAddsetsEpi32([64]byte(src), uint16(k), [64]byte(v2), [64]byte(v3), uint16(sign)))
}

func maskAddsetsEpi32(src [64]byte, k uint16, v2 [64]byte, v3 [64]byte, sign uint16) [64]byte


// AddsetsPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_addsets_ps'.
// Requires KNCNI.
func AddsetsPs(v2 x86.M512, v3 x86.M512, sign x86.Mmask16) x86.M512 {
	return x86.M512(addsetsPs([16]float32(v2), [16]float32(v3), uint16(sign)))
}

func addsetsPs(v2 [16]float32, v3 [16]float32, sign uint16) [16]float32


// MaskAddsetsPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag). Results
// are stored using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//				sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_mask_addsets_ps'.
// Requires KNCNI.
func MaskAddsetsPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, sign x86.Mmask16) x86.M512 {
	return x86.M512(maskAddsetsPs([16]float32(src), uint16(k), [16]float32(v2), [16]float32(v3), uint16(sign)))
}

func maskAddsetsPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, sign uint16) [16]float32


// AddsetsRoundPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_addsets_round_ps'.
// Requires KNCNI.
func AddsetsRoundPs(v2 x86.M512, v3 x86.M512, sign x86.Mmask16, rounding int) x86.M512 {
	return x86.M512(addsetsRoundPs([16]float32(v2), [16]float32(v3), uint16(sign), rounding))
}

func addsetsRoundPs(v2 [16]float32, v3 [16]float32, sign uint16, rounding int) [16]float32


// MaskAddsetsRoundPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag). Results
// are stored using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//				sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_mask_addsets_round_ps'.
// Requires KNCNI.
func MaskAddsetsRoundPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, sign x86.Mmask16, rounding int) x86.M512 {
	return x86.M512(maskAddsetsRoundPs([16]float32(src), uint16(k), [16]float32(v2), [16]float32(v3), uint16(sign), rounding))
}

func maskAddsetsRoundPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, sign uint16, rounding int) [16]float32


// AlignrEpi32: Concatenate 'a' and 'b' into a 128-byte immediate result, shift
// the result right by 'count' 32-bit elements, and store the low 64 bytes (16
// elements) in 'dst'. 
//
//		temp[1023:512] := a[511:0]
//		temp[511:0] := b[511:0]
//		temp[1023:0] := temp[1023:0] >> (32*count)
//		dst[511:0] := temp[511:0]
//		dst[MAX:512] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm512_alignr_epi32'.
// Requires KNCNI.
func AlignrEpi32(a x86.M512i, b x86.M512i, count int) x86.M512i {
	return x86.M512i(alignrEpi32([64]byte(a), [64]byte(b), count))
}

func alignrEpi32(a [64]byte, b [64]byte, count int) [64]byte


// MaskAlignrEpi32: Concatenate 'a' and 'b' into a 128-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 64
// bytes (16 elements) in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		temp[1023:512] := a[511:0]
//		temp[511:0] := b[511:0]
//		temp[1023:0] := temp[1023:0] >> (32*count)
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := temp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm512_mask_alignr_epi32'.
// Requires KNCNI.
func MaskAlignrEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i, count int) x86.M512i {
	return x86.M512i(maskAlignrEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b), count))
}

func maskAlignrEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte, count int) [64]byte


// AndEpi32: Compute the bitwise AND of packed 32-bit integers in 'a' and 'b',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] BITWISE AND b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_and_epi32'.
// Requires KNCNI.
func AndEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(andEpi32([64]byte(a), [64]byte(b)))
}

func andEpi32(a [64]byte, b [64]byte) [64]byte


// MaskAndEpi32: Performs element-by-element bitwise AND between packed 32-bit
// integer elements of 'v2' and 'v3', storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] & v3[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_mask_and_epi32'.
// Requires KNCNI.
func MaskAndEpi32(src x86.M512i, k x86.Mmask16, v2 x86.M512i, v3 x86.M512i) x86.M512i {
	return x86.M512i(maskAndEpi32([64]byte(src), uint16(k), [64]byte(v2), [64]byte(v3)))
}

func maskAndEpi32(src [64]byte, k uint16, v2 [64]byte, v3 [64]byte) [64]byte


// AndEpi64: Compute the bitwise AND of 512 bits (composed of packed 64-bit
// integers) in 'a' and 'b', and store the results in 'dst'. 
//
//		dst[511:0] := (a[511:0] AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_and_epi64'.
// Requires KNCNI.
func AndEpi64(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(andEpi64([64]byte(a), [64]byte(b)))
}

func andEpi64(a [64]byte, b [64]byte) [64]byte


// MaskAndEpi64: Compute the bitwise AND of packed 64-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] AND b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_mask_and_epi64'.
// Requires KNCNI.
func MaskAndEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskAndEpi64([64]byte(src), uint8(k), [64]byte(a), [64]byte(b)))
}

func maskAndEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte


// AndSi512: Compute the bitwise AND of 512 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := (a[511:0] AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_and_si512'.
// Requires KNCNI.
func AndSi512(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(andSi512([64]byte(a), [64]byte(b)))
}

func andSi512(a [64]byte, b [64]byte) [64]byte


// AndnotEpi32: Compute the bitwise AND NOT of packed 32-bit integers in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (NOT a[i+31:i]) AND b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDND'. Intrinsic: '_mm512_andnot_epi32'.
// Requires KNCNI.
func AndnotEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(andnotEpi32([64]byte(a), [64]byte(b)))
}

func andnotEpi32(a [64]byte, b [64]byte) [64]byte


// MaskAndnotEpi32: Compute the bitwise AND NOT of packed 32-bit integers in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
// Instruction: 'VPANDND'. Intrinsic: '_mm512_mask_andnot_epi32'.
// Requires KNCNI.
func MaskAndnotEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskAndnotEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskAndnotEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// AndnotEpi64: Compute the bitwise AND NOT of 512 bits (composed of packed
// 64-bit integers) in 'a' and 'b', and store the results in 'dst'. 
//
//		dst[511:0] := ((NOT a[511:0]) AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDNQ'. Intrinsic: '_mm512_andnot_epi64'.
// Requires KNCNI.
func AndnotEpi64(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(andnotEpi64([64]byte(a), [64]byte(b)))
}

func andnotEpi64(a [64]byte, b [64]byte) [64]byte


// MaskAndnotEpi64: Compute the bitwise AND NOT of packed 64-bit integers in
// 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
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
// Instruction: 'VPANDNQ'. Intrinsic: '_mm512_mask_andnot_epi64'.
// Requires KNCNI.
func MaskAndnotEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskAndnotEpi64([64]byte(src), uint8(k), [64]byte(a), [64]byte(b)))
}

func maskAndnotEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte


// AndnotSi512: Compute the bitwise AND NOT of 512 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := ((NOT a[511:0]) AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDND'. Intrinsic: '_mm512_andnot_si512'.
// Requires KNCNI.
func AndnotSi512(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(andnotSi512([64]byte(a), [64]byte(b)))
}

func andnotSi512(a [64]byte, b [64]byte) [64]byte


// MaskBlendEpi32: Blend packed 32-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBLENDMD'. Intrinsic: '_mm512_mask_blend_epi32'.
// Requires KNCNI.
func MaskBlendEpi32(k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskBlendEpi32(uint16(k), [64]byte(a), [64]byte(b)))
}

func maskBlendEpi32(k uint16, a [64]byte, b [64]byte) [64]byte


// MaskBlendEpi64: Blend packed 64-bit integers from 'a' and 'b' using control
// mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBLENDMQ'. Intrinsic: '_mm512_mask_blend_epi64'.
// Requires KNCNI.
func MaskBlendEpi64(k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskBlendEpi64(uint8(k), [64]byte(a), [64]byte(b)))
}

func maskBlendEpi64(k uint8, a [64]byte, b [64]byte) [64]byte


// MaskBlendPd: Blend packed double-precision (64-bit) floating-point elements
// from 'a' and 'b' using control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBLENDMPD'. Intrinsic: '_mm512_mask_blend_pd'.
// Requires KNCNI.
func MaskBlendPd(k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(maskBlendPd(uint8(k), [8]float64(a), [8]float64(b)))
}

func maskBlendPd(k uint8, a [8]float64, b [8]float64) [8]float64


// MaskBlendPs: Blend packed single-precision (32-bit) floating-point elements
// from 'a' and 'b' using control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBLENDMPS'. Intrinsic: '_mm512_mask_blend_ps'.
// Requires KNCNI.
func MaskBlendPs(k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskBlendPs(uint16(k), [16]float32(a), [16]float32(b)))
}

func maskBlendPs(k uint16, a [16]float32, b [16]float32) [16]float32


// CastpdPs: Cast vector of type __m512d to type __m512.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castpd_ps'.
// Requires KNCNI.
func CastpdPs(a x86.M512d) x86.M512 {
	return x86.M512(castpdPs([8]float64(a)))
}

func castpdPs(a [8]float64) [16]float32


// CastpdSi512: Cast vector of type __m512d to type __m512i.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castpd_si512'.
// Requires KNCNI.
func CastpdSi512(a x86.M512d) x86.M512i {
	return x86.M512i(castpdSi512([8]float64(a)))
}

func castpdSi512(a [8]float64) [64]byte


// CastpsPd: Cast vector of type __m512 to type __m512d.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castps_pd'.
// Requires KNCNI.
func CastpsPd(a x86.M512) x86.M512d {
	return x86.M512d(castpsPd([16]float32(a)))
}

func castpsPd(a [16]float32) [8]float64


// CastpsSi512: Cast vector of type __m512 to type __m512i.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castps_si512'.
// Requires KNCNI.
func CastpsSi512(a x86.M512) x86.M512i {
	return x86.M512i(castpsSi512([16]float32(a)))
}

func castpsSi512(a [16]float32) [64]byte


// Castsi512Pd: Cast vector of type __m512i to type __m512d.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castsi512_pd'.
// Requires KNCNI.
func Castsi512Pd(a x86.M512i) x86.M512d {
	return x86.M512d(castsi512Pd([64]byte(a)))
}

func castsi512Pd(a [64]byte) [8]float64


// Castsi512Ps: Cast vector of type __m512i to type __m512.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castsi512_ps'.
// Requires KNCNI.
func Castsi512Ps(a x86.M512i) x86.M512 {
	return x86.M512(castsi512Ps([64]byte(a)))
}

func castsi512Ps(a [64]byte) [16]float32


// Clevict: Evicts the cache line containing the address 'ptr' from cache level
// 'level' (can be either 0 or 1). 
//
//		CacheLineEvict(ptr, level)
//
// Instruction: 'CLEVICT0, CLEVICT1'. Intrinsic: '_mm_clevict'.
// Requires KNCNI.
func Clevict(ptr uintptr, level int)  {
	clevict(uintptr(ptr), level)
}

func clevict(ptr uintptr, level int) 


// CmpEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' based on the
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
//			i := j*32
//			k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmp_epi32_mask'.
// Requires KNCNI.
func CmpEpi32Mask(a x86.M512i, b x86.M512i, imm8 uint8) x86.Mmask16 {
	return x86.Mmask16(cmpEpi32Mask([64]byte(a), [64]byte(b), imm8))
}

func cmpEpi32Mask(a [64]byte, b [64]byte, imm8 uint8) uint16


// MaskCmpEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' based on the
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
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmp_epi32_mask'.
// Requires KNCNI.
func MaskCmpEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i, imm8 uint8) x86.Mmask16 {
	return x86.Mmask16(maskCmpEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b), imm8))
}

func maskCmpEpi32Mask(k1 uint16, a [64]byte, b [64]byte, imm8 uint8) uint16


// CmpEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' based
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
//			i := j*32
//			k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmp_epu32_mask'.
// Requires KNCNI.
func CmpEpu32Mask(a x86.M512i, b x86.M512i, imm8 uint8) x86.Mmask16 {
	return x86.Mmask16(cmpEpu32Mask([64]byte(a), [64]byte(b), imm8))
}

func cmpEpu32Mask(a [64]byte, b [64]byte, imm8 uint8) uint16


// MaskCmpEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
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
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmp_epu32_mask'.
// Requires KNCNI.
func MaskCmpEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i, imm8 uint8) x86.Mmask16 {
	return x86.Mmask16(maskCmpEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b), imm8))
}

func maskCmpEpu32Mask(k1 uint16, a [64]byte, b [64]byte, imm8 uint8) uint16


// CmpPdMask: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' based on the comparison operand specified by 'imm8', and
// store the results in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] OP b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmp_pd_mask'.
// Requires KNCNI.
func CmpPdMask(a x86.M512d, b x86.M512d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(cmpPdMask([8]float64(a), [8]float64(b), imm8))
}

func cmpPdMask(a [8]float64, b [8]float64, imm8 int) uint8


// MaskCmpPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k' using zeromask 'k1' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := ( a[i+63:i] OP b[i+63:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmp_pd_mask'.
// Requires KNCNI.
func MaskCmpPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d, imm8 int) x86.Mmask8 {
	return x86.Mmask8(maskCmpPdMask(uint8(k1), [8]float64(a), [8]float64(b), imm8))
}

func maskCmpPdMask(k1 uint8, a [8]float64, b [8]float64, imm8 int) uint8


// CmpPsMask: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' based on the comparison operand specified by 'imm8', and
// store the results in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] OP b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmp_ps_mask'.
// Requires KNCNI.
func CmpPsMask(a x86.M512, b x86.M512, imm8 int) x86.Mmask16 {
	return x86.Mmask16(cmpPsMask([16]float32(a), [16]float32(b), imm8))
}

func cmpPsMask(a [16]float32, b [16]float32, imm8 int) uint16


// MaskCmpPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k' using zeromask 'k1' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmp_ps_mask'.
// Requires KNCNI.
func MaskCmpPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512, imm8 int) x86.Mmask16 {
	return x86.Mmask16(maskCmpPsMask(uint16(k1), [16]float32(a), [16]float32(b), imm8))
}

func maskCmpPsMask(k1 uint16, a [16]float32, b [16]float32, imm8 int) uint16


// CmpRoundPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k'.
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] OP b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmp_round_pd_mask'.
// Requires KNCNI.
func CmpRoundPdMask(a x86.M512d, b x86.M512d, imm8 int, sae int) x86.Mmask8 {
	return x86.Mmask8(cmpRoundPdMask([8]float64(a), [8]float64(b), imm8, sae))
}

func cmpRoundPdMask(a [8]float64, b [8]float64, imm8 int, sae int) uint8


// MaskCmpRoundPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k' using zeromask 'k1' (elements are
// zeroed out when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := ( a[i+63:i] OP b[i+63:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmp_round_pd_mask'.
// Requires KNCNI.
func MaskCmpRoundPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d, imm8 int, sae int) x86.Mmask8 {
	return x86.Mmask8(maskCmpRoundPdMask(uint8(k1), [8]float64(a), [8]float64(b), imm8, sae))
}

func maskCmpRoundPdMask(k1 uint8, a [8]float64, b [8]float64, imm8 int, sae int) uint8


// CmpRoundPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k'.
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] OP b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmp_round_ps_mask'.
// Requires KNCNI.
func CmpRoundPsMask(a x86.M512, b x86.M512, imm8 int, sae int) x86.Mmask16 {
	return x86.Mmask16(cmpRoundPsMask([16]float32(a), [16]float32(b), imm8, sae))
}

func cmpRoundPsMask(a [16]float32, b [16]float32, imm8 int, sae int) uint16


// MaskCmpRoundPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k' using zeromask 'k1' (elements are
// zeroed out when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmp_round_ps_mask'.
// Requires KNCNI.
func MaskCmpRoundPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512, imm8 int, sae int) x86.Mmask16 {
	return x86.Mmask16(maskCmpRoundPsMask(uint16(k1), [16]float32(a), [16]float32(b), imm8, sae))
}

func maskCmpRoundPsMask(k1 uint16, a [16]float32, b [16]float32, imm8 int, sae int) uint16


// CmpeqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for equality,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPEQD'. Intrinsic: '_mm512_cmpeq_epi32_mask'.
// Requires KNCNI.
func CmpeqEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpeqEpi32Mask([64]byte(a), [64]byte(b)))
}

func cmpeqEpi32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpeqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPEQD'. Intrinsic: '_mm512_mask_cmpeq_epi32_mask'.
// Requires KNCNI.
func MaskCmpeqEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpeqEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpeqEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpeqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpeq_epu32_mask'.
// Requires KNCNI.
func CmpeqEpu32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpeqEpu32Mask([64]byte(a), [64]byte(b)))
}

func cmpeqEpu32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpeqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpeq_epu32_mask'.
// Requires KNCNI.
func MaskCmpeqEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpeqEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpeqEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpeqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] == b[i+63:i]) ? 1 : 0
//		ENDFOR	
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpeq_pd_mask'.
// Requires KNCNI.
func CmpeqPdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpeqPdMask([8]float64(a), [8]float64(b)))
}

func cmpeqPdMask(a [8]float64, b [8]float64) uint8


// MaskCmpeqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] == b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR	
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpeq_pd_mask'.
// Requires KNCNI.
func MaskCmpeqPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpeqPdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpeqPdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpeqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] == b[i+31:i]) ? 1 : 0
//		ENDFOR	
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpeq_ps_mask'.
// Requires KNCNI.
func CmpeqPsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpeqPsMask([16]float32(a), [16]float32(b)))
}

func cmpeqPsMask(a [16]float32, b [16]float32) uint16


// MaskCmpeqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		y
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] == b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR		
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpeq_ps_mask'.
// Requires KNCNI.
func MaskCmpeqPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpeqPsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpeqPsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpgeEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmpge_epi32_mask'.
// Requires KNCNI.
func CmpgeEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpgeEpi32Mask([64]byte(a), [64]byte(b)))
}

func cmpgeEpi32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpgeEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmpge_epi32_mask'.
// Requires KNCNI.
func MaskCmpgeEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgeEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgeEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpgeEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpge_epu32_mask'.
// Requires KNCNI.
func CmpgeEpu32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpgeEpu32Mask([64]byte(a), [64]byte(b)))
}

func cmpgeEpu32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpgeEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpge_epu32_mask'.
// Requires KNCNI.
func MaskCmpgeEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgeEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgeEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpgtEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPGTD'. Intrinsic: '_mm512_cmpgt_epi32_mask'.
// Requires KNCNI.
func CmpgtEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpgtEpi32Mask([64]byte(a), [64]byte(b)))
}

func cmpgtEpi32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpgtEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPGTD'. Intrinsic: '_mm512_mask_cmpgt_epi32_mask'.
// Requires KNCNI.
func MaskCmpgtEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgtEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgtEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpgtEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpgt_epu32_mask'.
// Requires KNCNI.
func CmpgtEpu32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpgtEpu32Mask([64]byte(a), [64]byte(b)))
}

func cmpgtEpu32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpgtEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpgt_epu32_mask'.
// Requires KNCNI.
func MaskCmpgtEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpgtEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpgtEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpleEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmple_epi32_mask'.
// Requires KNCNI.
func CmpleEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpleEpi32Mask([64]byte(a), [64]byte(b)))
}

func cmpleEpi32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpleEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmple_epi32_mask'.
// Requires KNCNI.
func MaskCmpleEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpleEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpleEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpleEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmple_epu32_mask'.
// Requires KNCNI.
func CmpleEpu32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpleEpu32Mask([64]byte(a), [64]byte(b)))
}

func cmpleEpu32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpleEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmple_epu32_mask'.
// Requires KNCNI.
func MaskCmpleEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpleEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpleEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmplePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] <= b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmple_pd_mask'.
// Requires KNCNI.
func CmplePdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmplePdMask([8]float64(a), [8]float64(b)))
}

func cmplePdMask(a [8]float64, b [8]float64) uint8


// MaskCmplePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] <= b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmple_pd_mask'.
// Requires KNCNI.
func MaskCmplePdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmplePdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmplePdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmplePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] <= b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmple_ps_mask'.
// Requires KNCNI.
func CmplePsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmplePsMask([16]float32(a), [16]float32(b)))
}

func cmplePsMask(a [16]float32, b [16]float32) uint16


// MaskCmplePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] <= b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmple_ps_mask'.
// Requires KNCNI.
func MaskCmplePsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmplePsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmplePsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpltEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for less-than,
// and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPLTD'. Intrinsic: '_mm512_cmplt_epi32_mask'.
// Requires KNCNI.
func CmpltEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpltEpi32Mask([64]byte(a), [64]byte(b)))
}

func cmpltEpi32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpltEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPLTD'. Intrinsic: '_mm512_mask_cmplt_epi32_mask'.
// Requires KNCNI.
func MaskCmpltEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpltEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpltEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpltEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmplt_epu32_mask'.
// Requires KNCNI.
func CmpltEpu32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpltEpu32Mask([64]byte(a), [64]byte(b)))
}

func cmpltEpu32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpltEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmplt_epu32_mask'.
// Requires KNCNI.
func MaskCmpltEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpltEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpltEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] < b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmplt_pd_mask'.
// Requires KNCNI.
func CmpltPdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpltPdMask([8]float64(a), [8]float64(b)))
}

func cmpltPdMask(a [8]float64, b [8]float64) uint8


// MaskCmpltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] < b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmplt_pd_mask'.
// Requires KNCNI.
func MaskCmpltPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpltPdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpltPdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] < b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmplt_ps_mask'.
// Requires KNCNI.
func CmpltPsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpltPsMask([16]float32(a), [16]float32(b)))
}

func cmpltPsMask(a [16]float32, b [16]float32) uint16


// MaskCmpltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] < b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmplt_ps_mask'.
// Requires KNCNI.
func MaskCmpltPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpltPsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpltPsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpneqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmpneq_epi32_mask'.
// Requires KNCNI.
func CmpneqEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpneqEpi32Mask([64]byte(a), [64]byte(b)))
}

func cmpneqEpi32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpneqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmpneq_epi32_mask'.
// Requires KNCNI.
func MaskCmpneqEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpneqEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpneqEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpneqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpneq_epu32_mask'.
// Requires KNCNI.
func CmpneqEpu32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(cmpneqEpu32Mask([64]byte(a), [64]byte(b)))
}

func cmpneqEpu32Mask(a [64]byte, b [64]byte) uint16


// MaskCmpneqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpneq_epu32_mask'.
// Requires KNCNI.
func MaskCmpneqEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskCmpneqEpu32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskCmpneqEpu32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// CmpneqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-equal, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] != b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpneq_pd_mask'.
// Requires KNCNI.
func CmpneqPdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpneqPdMask([8]float64(a), [8]float64(b)))
}

func cmpneqPdMask(a [8]float64, b [8]float64) uint8


// MaskCmpneqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-equal, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] != b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpneq_pd_mask'.
// Requires KNCNI.
func MaskCmpneqPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpneqPdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpneqPdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpneqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-equal, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] != b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpneq_ps_mask'.
// Requires KNCNI.
func CmpneqPsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpneqPsMask([16]float32(a), [16]float32(b)))
}

func cmpneqPsMask(a [16]float32, b [16]float32) uint16


// MaskCmpneqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-equal, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] != b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpneq_ps_mask'.
// Requires KNCNI.
func MaskCmpneqPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpneqPsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpneqPsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpnlePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := !(a[i+63:i] <= b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpnle_pd_mask'.
// Requires KNCNI.
func CmpnlePdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpnlePdMask([8]float64(a), [8]float64(b)))
}

func cmpnlePdMask(a [8]float64, b [8]float64) uint8


// MaskCmpnlePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := !(a[i+63:i] <= b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpnle_pd_mask'.
// Requires KNCNI.
func MaskCmpnlePdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpnlePdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpnlePdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpnlePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := !(a[i+31:i] <= b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpnle_ps_mask'.
// Requires KNCNI.
func CmpnlePsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpnlePsMask([16]float32(a), [16]float32(b)))
}

func cmpnlePsMask(a [16]float32, b [16]float32) uint16


// MaskCmpnlePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := !(a[i+31:i] <= b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpnle_ps_mask'.
// Requires KNCNI.
func MaskCmpnlePsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpnlePsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpnlePsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpnltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than, and store the results in mask
// vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := !(a[i+63:i] < b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpnlt_pd_mask'.
// Requires KNCNI.
func CmpnltPdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpnltPdMask([8]float64(a), [8]float64(b)))
}

func cmpnltPdMask(a [8]float64, b [8]float64) uint8


// MaskCmpnltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than, and store the results in mask
// vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := !(a[i+63:i] < b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpnlt_pd_mask'.
// Requires KNCNI.
func MaskCmpnltPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpnltPdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpnltPdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpnltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than, and store the results in mask
// vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := !(a[i+31:i] < b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpnlt_ps_mask'.
// Requires KNCNI.
func CmpnltPsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpnltPsMask([16]float32(a), [16]float32(b)))
}

func cmpnltPsMask(a [16]float32, b [16]float32) uint16


// MaskCmpnltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than, and store the results in mask
// vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := !(a[i+31:i] < b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpnlt_ps_mask'.
// Requires KNCNI.
func MaskCmpnltPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpnltPsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpnltPsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpordPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] != NaN AND b[i+63:i] != NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpord_pd_mask'.
// Requires KNCNI.
func CmpordPdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpordPdMask([8]float64(a), [8]float64(b)))
}

func cmpordPdMask(a [8]float64, b [8]float64) uint8


// MaskCmpordPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] != NaN AND b[i+63:i] != NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpord_pd_mask'.
// Requires KNCNI.
func MaskCmpordPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpordPdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpordPdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpordPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] != NaN AND b[i+31:i] != NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpord_ps_mask'.
// Requires KNCNI.
func CmpordPsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpordPsMask([16]float32(a), [16]float32(b)))
}

func cmpordPsMask(a [16]float32, b [16]float32) uint16


// MaskCmpordPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] != NaN AND b[i+31:i] != NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpord_ps_mask'.
// Requires KNCNI.
func MaskCmpordPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpordPsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpordPsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// CmpunordPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] == NaN OR b[i+63:i] == NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpunord_pd_mask'.
// Requires KNCNI.
func CmpunordPdMask(a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(cmpunordPdMask([8]float64(a), [8]float64(b)))
}

func cmpunordPdMask(a [8]float64, b [8]float64) uint8


// MaskCmpunordPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] == NaN OR b[i+63:i] == NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpunord_pd_mask'.
// Requires KNCNI.
func MaskCmpunordPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) x86.Mmask8 {
	return x86.Mmask8(maskCmpunordPdMask(uint8(k1), [8]float64(a), [8]float64(b)))
}

func maskCmpunordPdMask(k1 uint8, a [8]float64, b [8]float64) uint8


// CmpunordPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] == NaN OR b[i+31:i] == NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpunord_ps_mask'.
// Requires KNCNI.
func CmpunordPsMask(a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(cmpunordPsMask([16]float32(a), [16]float32(b)))
}

func cmpunordPsMask(a [16]float32, b [16]float32) uint16


// MaskCmpunordPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] == NaN OR b[i+31:i] == NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpunord_ps_mask'.
// Requires KNCNI.
func MaskCmpunordPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) x86.Mmask16 {
	return x86.Mmask16(maskCmpunordPsMask(uint16(k1), [16]float32(a), [16]float32(b)))
}

func maskCmpunordPsMask(k1 uint16, a [16]float32, b [16]float32) uint16


// Countbits32: Counts the number of set bits in 32-bit unsigned integer 'r1',
// returning the results in 'dst'. 
//
//		dst[31:0] := PopCount(r1[31:0])
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_countbits_32'.
// Requires KNCNI.
func Countbits32(r1 uint32) uint32 {
	return uint32(countbits32(r1))
}

func countbits32(r1 uint32) uint32


// Countbits64: Counts the number of set bits in double-precision (32-bit)
// unsigned integer 'r1', returning the results in 'dst'. 
//
//		dst[63:0] := PopCount(r1[63:0])
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_countbits_64'.
// Requires KNCNI.
func Countbits64(r1 uint64) uint64 {
	return uint64(countbits64(r1))
}

func countbits64(r1 uint64) uint64


// CvtRoundpdPslo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed
// single-precision (32-bit) floating-point elements, storing the results in
// 'dst'. Results are written to the lower half of 'dst', and the upper half
// locations are set to '0'.
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
//			k := j*32
//			dst[k+31:k] := Float64ToFloat32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_cvt_roundpd_pslo'.
// Requires KNCNI.
func CvtRoundpdPslo(v2 x86.M512d, rounding int) x86.M512 {
	return x86.M512(cvtRoundpdPslo([8]float64(v2), rounding))
}

func cvtRoundpdPslo(v2 [8]float64, rounding int) [16]float32


// MaskCvtRoundpdPslo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed
// single-precision (32-bit) floating-point elements, storing the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). Results are written to the lower half of
// 'dst', and the upper half locations are set to '0'.
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
//				dst[l+31:l] := Float64ToFloat32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_mask_cvt_roundpd_pslo'.
// Requires KNCNI.
func MaskCvtRoundpdPslo(src x86.M512, k x86.Mmask8, v2 x86.M512d, rounding int) x86.M512 {
	return x86.M512(maskCvtRoundpdPslo([16]float32(src), uint8(k), [8]float64(v2), rounding))
}

func maskCvtRoundpdPslo(src [16]float32, k uint8, v2 [8]float64, rounding int) [16]float32


// Cvtepi32loPd: Performs element-by-element conversion of the lower half of
// packed 32-bit integer elements in 'v2' to packed double-precision (64-bit)
// floating-point elements, storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			l := j*64
//			dst[l+63:l] := Int32ToFloat64(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTDQ2PD'. Intrinsic: '_mm512_cvtepi32lo_pd'.
// Requires KNCNI.
func Cvtepi32loPd(v2 x86.M512i) x86.M512d {
	return x86.M512d(cvtepi32loPd([64]byte(v2)))
}

func cvtepi32loPd(v2 [64]byte) [8]float64


// MaskCvtepi32loPd: Performs element-by-element conversion of the lower half
// of packed 32-bit integer elements in 'v2' to packed double-precision
// (64-bit) floating-point elements, storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := j*64
//			IF k[j]
//				dst[k+63:k] := Int32ToFloat64(v2[i+31:i])
//			ELSE
//				dst[n+63:n] := src[n+63:n]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTDQ2PD'. Intrinsic: '_mm512_mask_cvtepi32lo_pd'.
// Requires KNCNI.
func MaskCvtepi32loPd(src x86.M512d, k x86.Mmask8, v2 x86.M512i) x86.M512d {
	return x86.M512d(maskCvtepi32loPd([8]float64(src), uint8(k), [64]byte(v2)))
}

func maskCvtepi32loPd(src [8]float64, k uint8, v2 [64]byte) [8]float64


// Cvtepu32loPd: Performs element-by-element conversion of the lower half of
// packed 32-bit unsigned integer elements in 'v2' to packed double-precision
// (64-bit) floating-point elements, storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			k := j*64
//			dst[k+63:k] := UInt32ToFloat64(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUDQ2PD'. Intrinsic: '_mm512_cvtepu32lo_pd'.
// Requires KNCNI.
func Cvtepu32loPd(v2 x86.M512i) x86.M512d {
	return x86.M512d(cvtepu32loPd([64]byte(v2)))
}

func cvtepu32loPd(v2 [64]byte) [8]float64


// MaskCvtepu32loPd: Performs element-by-element conversion of the lower half
// of 32-bit unsigned integer elements in 'v2' to packed double-precision
// (64-bit) floating-point elements, storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			l := j*64
//			IF k[j]
//				dst[l+63:l] := UInt32ToFloat64(v2[i+31:i])
//			ELSE
//				dst[l+63:l] := src[l+63:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUDQ2PD'. Intrinsic: '_mm512_mask_cvtepu32lo_pd'.
// Requires KNCNI.
func MaskCvtepu32loPd(src x86.M512d, k x86.Mmask8, v2 x86.M512i) x86.M512d {
	return x86.M512d(maskCvtepu32loPd([8]float64(src), uint8(k), [64]byte(v2)))
}

func maskCvtepu32loPd(src [8]float64, k uint8, v2 [64]byte) [8]float64


// CvtfxpntRoundAdjustepi32Ps: Performs element-by-element conversion of packed
// 32-bit integer elements in 'v2' to packed single-precision (32-bit)
// floating-point elements and performing an optional exponent adjust using
// 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Int32ToFloat32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTDQ2PS'. Intrinsic: '_mm512_cvtfxpnt_round_adjustepi32_ps'.
// Requires KNCNI.
func CvtfxpntRoundAdjustepi32Ps(v2 x86.M512i, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(cvtfxpntRoundAdjustepi32Ps([64]byte(v2), rounding, expadj))
}

func cvtfxpntRoundAdjustepi32Ps(v2 [64]byte, rounding int, expadj MMEXPADJENUM) [16]float32


// CvtfxpntRoundAdjustepu32Ps: Performs element-by-element conversion of packed
// 32-bit unsigned integer elements in 'v2' to packed single-precision (32-bit)
// floating-point elements and performing an optional exponent adjust using
// 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := UInt32ToFloat32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTUDQ2PS'. Intrinsic: '_mm512_cvtfxpnt_round_adjustepu32_ps'.
// Requires KNCNI.
func CvtfxpntRoundAdjustepu32Ps(v2 x86.M512i, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(cvtfxpntRoundAdjustepu32Ps([64]byte(v2), rounding, expadj))
}

func cvtfxpntRoundAdjustepu32Ps(v2 [64]byte, rounding int, expadj MMEXPADJENUM) [16]float32


// MaskCvtfxpntRoundAdjustepu32Ps: Performs element-by-element conversion of
// packed 32-bit unsigned integer elements in 'v2' to packed single-precision
// (32-bit) floating-point elements and performing an optional exponent adjust
// using 'expadj', storing the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := Int32ToFloat32(v2[i+31:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTUDQ2PS'. Intrinsic: '_mm512_mask_cvtfxpnt_round_adjustepu32_ps'.
// Requires KNCNI.
func MaskCvtfxpntRoundAdjustepu32Ps(src x86.M512, k x86.Mmask16, v2 x86.M512i, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(maskCvtfxpntRoundAdjustepu32Ps([16]float32(src), uint16(k), [64]byte(v2), rounding, expadj))
}

func maskCvtfxpntRoundAdjustepu32Ps(src [16]float32, k uint16, v2 [64]byte, rounding int, expadj MMEXPADJENUM) [16]float32


// CvtfxpntRoundAdjustpsEpi32: Performs element-by-element conversion of packed
// single-precision (32-bit) floating-point elements in 'v2' to packed 32-bit
// integer elements and performs an optional exponent adjust using 'expadj',
// storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Float32ToInt32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPS2DQ'. Intrinsic: '_mm512_cvtfxpnt_round_adjustps_epi32'.
// Requires KNCNI.
func CvtfxpntRoundAdjustpsEpi32(v2 x86.M512, rounding int, expadj MMEXPADJENUM) x86.M512i {
	return x86.M512i(cvtfxpntRoundAdjustpsEpi32([16]float32(v2), rounding, expadj))
}

func cvtfxpntRoundAdjustpsEpi32(v2 [16]float32, rounding int, expadj MMEXPADJENUM) [64]byte


// CvtfxpntRoundAdjustpsEpu32: Performs element-by-element conversion of packed
// single-precision (32-bit) floating-point elements in 'v2' to packed 32-bit
// unsigned integer elements and performing an optional exponent adjust using
// 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Float32ToUInt32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i]  0
//			_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i]  4
//			_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i]  5
//			_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i]  8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i]  16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i]  24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i]  31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i]  32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPS2UDQ'. Intrinsic: '_mm512_cvtfxpnt_round_adjustps_epu32'.
// Requires KNCNI.
func CvtfxpntRoundAdjustpsEpu32(v2 x86.M512, rounding int, expadj MMEXPADJENUM) x86.M512i {
	return x86.M512i(cvtfxpntRoundAdjustpsEpu32([16]float32(v2), rounding, expadj))
}

func cvtfxpntRoundAdjustpsEpu32(v2 [16]float32, rounding int, expadj MMEXPADJENUM) [64]byte


// CvtfxpntRoundpdEpi32lo: Performs an element-by-element conversion of
// elements in packed double-precision (64-bit) floating-point vector 'v2' to
// 32-bit integer elements, storing them in the lower half of 'dst'. The
// elements in the upper half of 'dst' are set to 0.
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
//			k := j*32
//			dst[k+31:k] := Float64ToInt32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2DQ'. Intrinsic: '_mm512_cvtfxpnt_roundpd_epi32lo'.
// Requires KNCNI.
func CvtfxpntRoundpdEpi32lo(v2 x86.M512d, rounding int) x86.M512i {
	return x86.M512i(cvtfxpntRoundpdEpi32lo([8]float64(v2), rounding))
}

func cvtfxpntRoundpdEpi32lo(v2 [8]float64, rounding int) [64]byte


// MaskCvtfxpntRoundpdEpi32lo: Performs an element-by-element conversion of
// elements in packed double-precision (64-bit) floating-point vector 'v2' to
// 32-bit integer elements, storing them in the lower half of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). The elements in the upper half of 'dst' are set to 0.
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
//				dst[l+31:l] := Float64ToInt32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2DQ'. Intrinsic: '_mm512_mask_cvtfxpnt_roundpd_epi32lo'.
// Requires KNCNI.
func MaskCvtfxpntRoundpdEpi32lo(src x86.M512i, k x86.Mmask8, v2 x86.M512d, rounding int) x86.M512i {
	return x86.M512i(maskCvtfxpntRoundpdEpi32lo([64]byte(src), uint8(k), [8]float64(v2), rounding))
}

func maskCvtfxpntRoundpdEpi32lo(src [64]byte, k uint8, v2 [8]float64, rounding int) [64]byte


// CvtfxpntRoundpdEpu32lo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed 32-bit
// unsigned integer elements, storing the results in 'dst'. Results are written
// to the lower half of 'dst', and the upper half locations are set to '0'.
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
//			k := j*32
//			dst[k+31:k] := Float64ToInt32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2UDQ'. Intrinsic: '_mm512_cvtfxpnt_roundpd_epu32lo'.
// Requires KNCNI.
func CvtfxpntRoundpdEpu32lo(v2 x86.M512d, rounding int) x86.M512i {
	return x86.M512i(cvtfxpntRoundpdEpu32lo([8]float64(v2), rounding))
}

func cvtfxpntRoundpdEpu32lo(v2 [8]float64, rounding int) [64]byte


// MaskCvtfxpntRoundpdEpu32lo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed 32-bit
// unsigned integer elements, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// Results are written to the lower half of 'dst', and the upper half locations
// are set to '0'.
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
//				dst[l+31:l] := Float64ToInt32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2UDQ'. Intrinsic: '_mm512_mask_cvtfxpnt_roundpd_epu32lo'.
// Requires KNCNI.
func MaskCvtfxpntRoundpdEpu32lo(src x86.M512i, k x86.Mmask8, v2 x86.M512d, rounding int) x86.M512i {
	return x86.M512i(maskCvtfxpntRoundpdEpu32lo([64]byte(src), uint8(k), [8]float64(v2), rounding))
}

func maskCvtfxpntRoundpdEpu32lo(src [64]byte, k uint8, v2 [8]float64, rounding int) [64]byte


// CvtpdPslo: Performs an element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to
// single-precision (32-bit) floating-point elements and stores them in 'dst'.
// The elements are stored in the lower half of the results vector, while the
// remaining upper half locations are set to 0. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k := j*32
//			dst[k+31:k] := Float64ToFloat32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_cvtpd_pslo'.
// Requires KNCNI.
func CvtpdPslo(v2 x86.M512d) x86.M512 {
	return x86.M512(cvtpdPslo([8]float64(v2)))
}

func cvtpdPslo(v2 [8]float64) [16]float32


// MaskCvtpdPslo: Performs an element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to
// single-precision (32-bit) floating-point elements and stores them in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). The elements are stored in the lower half of the
// results vector, while the remaining upper half locations are set to 0. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Float64ToFloat32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_mask_cvtpd_pslo'.
// Requires KNCNI.
func MaskCvtpdPslo(src x86.M512, k x86.Mmask8, v2 x86.M512d) x86.M512 {
	return x86.M512(maskCvtpdPslo([16]float32(src), uint8(k), [8]float64(v2)))
}

func maskCvtpdPslo(src [16]float32, k uint8, v2 [8]float64) [16]float32


// CvtpsloPd: Performs element-by-element conversion of the lower half of
// packed single-precision (32-bit) floating-point elements in 'v2' to packed
// double-precision (64-bit) floating-point elements, storing the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			k := j*64
//			dst[k+63:k] := Float32ToFloat64(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2PD'. Intrinsic: '_mm512_cvtpslo_pd'.
// Requires KNCNI.
func CvtpsloPd(v2 x86.M512) x86.M512d {
	return x86.M512d(cvtpsloPd([16]float32(v2)))
}

func cvtpsloPd(v2 [16]float32) [8]float64


// MaskCvtpsloPd: Performs element-by-element conversion of the lower half of
// packed single-precision (32-bit) floating-point elements in 'v2' to packed
// double-precision (64-bit) floating-point elements, storing the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			l := j*64
//			IF k[j]
//				dst[l+63:l] := Float32ToFloat64(v2[i+31:i])
//			ELSE
//				dst[l+63:l] := src[l+63:l]:
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2PD'. Intrinsic: '_mm512_mask_cvtpslo_pd'.
// Requires KNCNI.
func MaskCvtpsloPd(src x86.M512d, k x86.Mmask8, v2 x86.M512) x86.M512d {
	return x86.M512d(maskCvtpsloPd([8]float64(src), uint8(k), [16]float32(v2)))
}

func maskCvtpsloPd(src [8]float64, k uint8, v2 [16]float32) [8]float64


// Delay32: Stalls a thread without blocking other threads for 32-bit unsigned
// integer 'r1' clock cycles. 
//
//		BlockThread(r1)
//
// Instruction: 'DELAY'. Intrinsic: '_mm_delay_32'.
// Requires KNCNI.
func Delay32(r1 uint32)  {
	delay32(r1)
}

func delay32(r1 uint32) 


// Delay64: Stalls a thread without blocking other threads for 64-bit unsigned
// integer 'r1' clock cycles. 
//
//		BlockThread(r1)
//
// Instruction: 'DELAY'. Intrinsic: '_mm_delay_64'.
// Requires KNCNI.
func Delay64(r1 uint64)  {
	delay64(r1)
}

func delay64(r1 uint64) 


// Exp223Ps: Approximates the base-2 exponent of the packed single-precision
// (32-bit) floating-point elements in 'v2' with eight bits for sign and
// magnitude and 24 bits for the fractional part. Results are stored in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := exp223(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP223PS'. Intrinsic: '_mm512_exp223_ps'.
// Requires KNCNI.
func Exp223Ps(v2 x86.M512i) x86.M512 {
	return x86.M512(exp223Ps([64]byte(v2)))
}

func exp223Ps(v2 [64]byte) [16]float32


// MaskExp223Ps: Approximates the base-2 exponent of the packed
// single-precision (32-bit) floating-point elements in 'v2' with eight bits
// for sign and magnitude and 24 bits for the fractional part. Results are
// stored in 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := exp223(v2[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP223PS'. Intrinsic: '_mm512_mask_exp223_ps'.
// Requires KNCNI.
func MaskExp223Ps(src x86.M512, k x86.Mmask16, v2 x86.M512i) x86.M512 {
	return x86.M512(maskExp223Ps([16]float32(src), uint16(k), [64]byte(v2)))
}

func maskExp223Ps(src [16]float32, k uint16, v2 [64]byte) [16]float32


// ExtloadEpi32: Depending on 'bc', loads 1, 4, or 16 elements of type and size
// determined by 'conv' from memory address 'mt' and converts all elements to
// 32-bit integer elements, storing the results in 'dst'. 'hint' indicates to
// the processor whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			CASE bc OF
//			_MM_BROADCAST32_NONE:
//				CASE conv OF
//				_MM_UPCONV_EPI32_NONE:
//					n	 := j*32
//					dst[i+31:i] := addr[n+31:n]
//				_MM_UPCONV_EPI32_UINT8:
//					n	 := j*8
//					dst[i+31:i] := UInt8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_SINT8:
//					n	 := j*8
//					dst[i+31:i] := SInt8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_UINT16:
//					n	 := j*16
//					dst[i+31:i] := UInt16ToInt32(addr[n+15:n])
//				_MM_UPCONV_EPI32_SINT16:
//					n	 := j*16
//					dst[i+31:i] := SInt16ToInt32(addr[n+15:n])
//				ESAC
//			_MM_BROADCAST_1X16:
//				CASE conv OF
//				_MM_UPCONV_EPI32_NONE:
//					n	 := j*32
//					dst[i+31:i] := addr[31:0]
//				_MM_UPCONV_EPI32_UINT8:
//					n	 := j*8
//					dst[i+31:i] := UInt8ToInt32(addr[7:0])
//				_MM_UPCONV_EPI32_SINT8:
//					n	 := j*8
//					dst[i+31:i] := SInt8ToInt32(addr[7:0])
//				_MM_UPCONV_EPI32_UINT16:
//					n	 := j*16
//					dst[i+31:i] := UInt16ToInt32(addr[15:0])
//				_MM_UPCONV_EPI32_SINT16:
//					n	 := j*16
//					dst[i+31:i] := SInt16ToInt32(addr[15:0])
//				ESAC
//			_MM_BROADCAST_4X16:
//				mod := j%4
//				CASE conv OF
//				_MM_UPCONV_EPI32_NONE:
//					n := mod*32
//					dst[i+31:i] := addr[n+31:n]
//				_MM_UPCONV_EPI32_UINT8:
//					n := mod*8
//					dst[i+31:i] := UInt8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_SINT8:
//					n := mod*8
//					dst[i+31:i] := SInt8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_UINT16:
//					n := mod*16
//					dst[i+31:i] := UInt16ToInt32(addr[n+15:n])
//				_MM_UPCONV_EPI32_SINT16:
//					n := mod*16
//					dst[i+31:i] := SInt16ToInt32(addr[n+15:n])
//				ESAC
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32, VBROADCASTI32X4, VPBROADCASTD'. Intrinsic: '_mm512_extload_epi32'.
// Requires KNCNI.
func ExtloadEpi32(mt uintptr, conv MMUPCONVEPI32ENUM, bc MMBROADCAST32ENUM, hint int) x86.M512i {
	return x86.M512i(extloadEpi32(uintptr(mt), conv, bc, hint))
}

func extloadEpi32(mt uintptr, conv MMUPCONVEPI32ENUM, bc MMBROADCAST32ENUM, hint int) [64]byte


// MaskExtloadEpi32: Depending on 'bc', loads 1, 4, or 16 elements of type and
// size determined by 'conv' from memory address 'mt' and converts all elements
// to 32-bit integer elements, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				CASE bc OF
//				_MM_BROADCAST32_NONE:
//					CASE conv OF
//					_MM_UPCONV_EPI32_NONE:
//						n	 := j*32
//						dst[i+31:i] := addr[n+31:n]
//					_MM_UPCONV_EPI32_UINT8:
//						n	 := j*8
//						dst[i+31:i] := UInt8ToInt32(addr[n+7:n])
//					_MM_UPCONV_EPI32_SINT8:
//						n	 := j*8
//						dst[i+31:i] := SInt8ToInt32(addr[n+7:n])
//					_MM_UPCONV_EPI32_UINT16:
//						n	 := j*16
//						dst[i+31:i] := UInt16ToInt32(addr[n+15:n])
//					_MM_UPCONV_EPI32_SINT16:
//						n	 := j*16
//						dst[i+31:i] := SInt16ToInt32(addr[n+15:n])
//					ESAC
//				_MM_BROADCAST_1X16:
//					CASE conv OF
//					_MM_UPCONV_EPI32_NONE:
//						n	 := j*32
//						dst[i+31:i] := addr[31:0]
//					_MM_UPCONV_EPI32_UINT8:
//						n	 := j*8
//						dst[i+31:i] := UInt8ToInt32(addr[7:0])
//					_MM_UPCONV_EPI32_SINT8:
//						n	 := j*8
//						dst[i+31:i] := SInt8ToInt32(addr[7:0])
//					_MM_UPCONV_EPI32_UINT16:
//						n	 := j*16
//						dst[i+31:i] := UInt16ToInt32(addr[15:0])
//					_MM_UPCONV_EPI32_SINT16:
//						n	 := j*16
//						dst[i+31:i] := SInt16ToInt32(addr[15:0])
//					ESAC
//				_MM_BROADCAST_4X16:
//					mod := j%4
//					CASE conv OF
//					_MM_UPCONV_EPI32_NONE:
//						n := mod*32
//						dst[i+31:i] := addr[n+31:n]
//					_MM_UPCONV_EPI32_UINT8:
//						n := mod*8
//						dst[i+31:i] := UInt8ToInt32(addr[n+7:n])
//					_MM_UPCONV_EPI32_SINT8:
//						n := mod*8
//						dst[i+31:i] := SInt8ToInt32(addr[n+7:n])
//					_MM_UPCONV_EPI32_UINT16:
//						n := mod*16
//						dst[i+31:i] := UInt16ToInt32(addr[n+15:n])
//					_MM_UPCONV_EPI32_SINT16:
//						n := mod*16
//						dst[i+31:i] := SInt16ToInt32(addr[n+15:n])
//					ESAC
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32, VBROADCASTI32X4, VPBROADCASTD'. Intrinsic: '_mm512_mask_extload_epi32'.
// Requires KNCNI.
func MaskExtloadEpi32(src x86.M512i, k x86.Mmask16, mt uintptr, conv MMUPCONVEPI32ENUM, bc MMBROADCAST32ENUM, hint int) x86.M512i {
	return x86.M512i(maskExtloadEpi32([64]byte(src), uint16(k), uintptr(mt), conv, bc, hint))
}

func maskExtloadEpi32(src [64]byte, k uint16, mt uintptr, conv MMUPCONVEPI32ENUM, bc MMBROADCAST32ENUM, hint int) [64]byte


// ExtloadEpi64: Depending on 'bc', loads 1, 4, or 8 elements of type and size
// determined by 'conv' from memory address 'mt' and converts all elements to
// 64-bit integer elements, storing the results in 'dst'. 'hint' indicates to
// the processor whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			CASE bc OF
//			_MM_BROADCAST64_NONE:
//				CASE conv OF
//				_MM_UPCONV_EPI64_NONE:
//					n := j*64
//					dst[i+63:i] := addr[n+63:n]
//				ESAC
//			_MM_BROADCAST_1X8:
//				CASE conv OF
//				_MM_UPCONV_EPI64_NONE:
//					n := j*64
//					dst[i+63:i] := addr[63:0]
//				ESAC
//			_MM_BROADCAST_4X8:
//				mod := j%4
//				CASE conv OF
//				_MM_UPCONV_EPI64_NONE:
//					n := mod*64
//					dst[i+63:i] := addr[n+63:n]
//				ESAC
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA64, VBROADCASTI64X4, VPBROADCASTQ'. Intrinsic: '_mm512_extload_epi64'.
// Requires KNCNI.
func ExtloadEpi64(mt uintptr, conv MMUPCONVEPI64ENUM, bc MMBROADCAST64ENUM, hint int) x86.M512i {
	return x86.M512i(extloadEpi64(uintptr(mt), conv, bc, hint))
}

func extloadEpi64(mt uintptr, conv MMUPCONVEPI64ENUM, bc MMBROADCAST64ENUM, hint int) [64]byte


// MaskExtloadEpi64: Depending on 'bc', loads 1, 4, or 8 elements of type and
// size determined by 'conv' from memory address 'mt' and converts all elements
// to 64-bit integer elements, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				CASE bc OF
//				_MM_BROADCAST64_NONE:
//					CASE conv OF
//					_MM_UPCONV_EPI64_NONE:
//						n := j*64
//						dst[i+63:i] := addr[n+63:n]
//					ESAC
//				_MM_BROADCAST_1X8:
//					CASE conv OF
//					_MM_UPCONV_EPI64_NONE:
//						n := j*64
//						dst[i+63:i] := addr[63:0]
//					ESAC
//				_MM_BROADCAST_4X8:
//					mod := j%4
//					CASE conv OF
//					_MM_UPCONV_EPI64_NONE:
//						n := mod*64
//						dst[i+63:i] := addr[n+63:n]
//					ESAC
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA64, VBROADCASTI64X4, VPBROADCASTQ'. Intrinsic: '_mm512_mask_extload_epi64'.
// Requires KNCNI.
func MaskExtloadEpi64(src x86.M512i, k x86.Mmask8, mt uintptr, conv MMUPCONVEPI64ENUM, bc MMBROADCAST64ENUM, hint int) x86.M512i {
	return x86.M512i(maskExtloadEpi64([64]byte(src), uint8(k), uintptr(mt), conv, bc, hint))
}

func maskExtloadEpi64(src [64]byte, k uint8, mt uintptr, conv MMUPCONVEPI64ENUM, bc MMBROADCAST64ENUM, hint int) [64]byte


// ExtloadPd: Depending on 'bc', loads 1, 4, or 8 elements of type and size
// determined by 'conv' from memory address 'mt' and converts all elements to
// double-precision (64-bit) floating-point elements, storing the results in
// 'dst'. 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			CASE bc OF
//			_MM_BROADCAST64_NONE:
//				CASE conv OF
//				_MM_UPCONV_PD_NONE:
//					n := j*64
//					dst[i+63:i] := addr[n+63:n]
//				ESAC
//			_MM_BROADCAST_1X8:
//				CASE conv OF
//				_MM_UPCONV_PD_NONE:
//					n := j*64
//					dst[i+63:i] := addr[63:0]
//				ESAC
//			_MM_BROADCAST_4X8:
//				mod := j%4
//				CASE conv OF
//				_MM_UPCONV_PD_NONE:
//					n := mod*64
//					dst[i+63:i] := addr[n+63:n]
//				ESAC
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPD, VBROADCASTF64X4, VBROADCASTSD'. Intrinsic: '_mm512_extload_pd'.
// Requires KNCNI.
func ExtloadPd(mt uintptr, conv MMUPCONVPDENUM, bc MMBROADCAST64ENUM, hint int) x86.M512d {
	return x86.M512d(extloadPd(uintptr(mt), conv, bc, hint))
}

func extloadPd(mt uintptr, conv MMUPCONVPDENUM, bc MMBROADCAST64ENUM, hint int) [8]float64


// MaskExtloadPd: Depending on 'bc', loads 1, 4, or 8 elements of type and size
// determined by 'conv' from memory address 'mt' and converts all elements to
// double-precision (64-bit) floating-point elements, storing the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 'hint' indicates to the processor
// whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				CASE bc OF
//				_MM_BROADCAST64_NONE:
//					CASE conv OF
//					_MM_UPCONV_PD_NONE:
//						n := j*64
//						dst[i+63:i] := addr[n+63:n]
//					ESAC
//				_MM_BROADCAST_1X8:
//					CASE conv OF
//					_MM_UPCONV_PD_NONE:
//						n := j*64
//						dst[i+63:i] := addr[63:0]
//					ESAC
//				_MM_BROADCAST_4X8:
//					mod := j%4
//					CASE conv OF
//					_MM_UPCONV_PD_NONE:
//						n := mod*64
//						dst[i+63:i] := addr[n+63:n]
//					ESAC
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPD, VBROADCASTF64X4, VBROADCASTSD'. Intrinsic: '_mm512_mask_extload_pd'.
// Requires KNCNI.
func MaskExtloadPd(src x86.M512d, k x86.Mmask8, mt uintptr, conv MMUPCONVPDENUM, bc MMBROADCAST64ENUM, hint int) x86.M512d {
	return x86.M512d(maskExtloadPd([8]float64(src), uint8(k), uintptr(mt), conv, bc, hint))
}

func maskExtloadPd(src [8]float64, k uint8, mt uintptr, conv MMUPCONVPDENUM, bc MMBROADCAST64ENUM, hint int) [8]float64


// ExtloadPs: Depending on 'bc', loads 1, 4, or 16 elements of type and size
// determined by 'conv' from memory address 'mt' and converts all elements to
// single-precision (32-bit) floating-point elements, storing the results in
// 'dst'. 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			CASE bc OF
//			_MM_BROADCAST32_NONE:
//				CASE conv OF
//				_MM_UPCONV_PS_NONE:
//					n	 := j*32
//					dst[i+31:i] := addr[n+31:n]
//				_MM_UPCONV_PS_FLOAT16:
//					n	 := j*16
//					dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_UINT8:
//					n	 := j*8
//					dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_SINT8:
//					n	 := j*8
//					dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_UINT16:
//					n	 := j*16
//					dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_SINT16:
//					n	 := j*16
//					dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//				ESAC
//			_MM_BROADCAST_1X16:
//				CASE conv OF
//				_MM_UPCONV_PS_NONE:
//					n	 := j*32
//					dst[i+31:i] := addr[31:0]
//				_MM_UPCONV_PS_FLOAT16:
//					n	 := j*16
//					dst[i+31:i] := Float16ToFloat32(addr[15:0])
//				_MM_UPCONV_PS_UINT8:
//					n	 := j*8
//					dst[i+31:i] := UInt8ToFloat32(addr[7:0])
//				_MM_UPCONV_PS_SINT8:
//					n	 := j*8
//					dst[i+31:i] := SInt8ToFloat32(addr[7:0])
//				_MM_UPCONV_PS_UINT16:
//					n	 := j*16
//					dst[i+31:i] := UInt16ToFloat32(addr[15:0])
//				_MM_UPCONV_PS_SINT16:
//					n	 := j*16
//					dst[i+31:i] := SInt16ToFloat32(addr[15:0])
//				ESAC
//			_MM_BROADCAST_4X16:
//				mod := j%4
//				CASE conv OF
//				_MM_UPCONV_PS_NONE:
//					n := mod*32
//					dst[i+31:i] := addr[n+31:n]
//				_MM_UPCONV_PS_FLOAT16:
//					n := mod*16
//					dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_UINT8:
//					n := mod*8
//					dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_SINT8:
//					n := mod*8
//					dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_UINT16:
//					n := mod*16
//					dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_SINT16:
//					n := mod*16
//					dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//				ESAC
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPS, VBROADCASTF32X4, VBROADCASTSS'. Intrinsic: '_mm512_extload_ps'.
// Requires KNCNI.
func ExtloadPs(mt uintptr, conv MMUPCONVPSENUM, bc MMBROADCAST32ENUM, hint int) x86.M512 {
	return x86.M512(extloadPs(uintptr(mt), conv, bc, hint))
}

func extloadPs(mt uintptr, conv MMUPCONVPSENUM, bc MMBROADCAST32ENUM, hint int) [16]float32


// MaskExtloadPs: Depending on 'bc', loads 1, 4, or 16 elements of type and
// size determined by 'conv' from memory address 'mt' and converts all elements
// to single-precision (32-bit) floating-point elements, storing the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 'hint' indicates to the processor
// whether the data is non-temporal. 
//
//		addr = MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				CASE bc OF
//				_MM_BROADCAST32_NONE:
//					CASE conv OF
//					_MM_UPCONV_PS_NONE:
//						n	 := j*32
//						dst[i+31:i] := addr[n+31:n]
//					_MM_UPCONV_PS_FLOAT16:
//						n	 := j*16
//						dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//					_MM_UPCONV_PS_UINT8:
//						n	 := j*8
//						dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//					_MM_UPCONV_PS_SINT8:
//						n	 := j*8
//						dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//					_MM_UPCONV_PS_UINT16:
//						n	 := j*16
//						dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//					_MM_UPCONV_PS_SINT16:
//						n	 := j*16
//						dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//					ESAC
//				_MM_BROADCAST_1X16:
//					CASE conv OF
//					_MM_UPCONV_PS_NONE:
//						n	 := j*32
//						dst[i+31:i] := addr[31:0]
//					_MM_UPCONV_PS_FLOAT16:
//						n	 := j*16
//						dst[i+31:i] := Float16ToFloat32(addr[15:0])
//					_MM_UPCONV_PS_UINT8:
//						n	 := j*8
//						dst[i+31:i] := UInt8ToFloat32(addr[7:0])
//					_MM_UPCONV_PS_SINT8:
//						n	 := j*8
//						dst[i+31:i] := SInt8ToFloat32(addr[7:0])
//					_MM_UPCONV_PS_UINT16:
//						n	 := j*16
//						dst[i+31:i] := UInt16ToFloat32(addr[15:0])
//					_MM_UPCONV_PS_SINT16:
//						n	 := j*16
//						dst[i+31:i] := SInt16ToFloat32(addr[15:0])
//					ESAC
//				_MM_BROADCAST_4X16:
//					mod := j%4
//					CASE conv OF
//					_MM_UPCONV_PS_NONE:
//						n := mod*32
//						dst[i+31:i] := addr[n+31:n]
//					_MM_UPCONV_PS_FLOAT16:
//						n := mod*16
//						dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//					_MM_UPCONV_PS_UINT8:
//						n := mod*8
//						dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//					_MM_UPCONV_PS_SINT8:
//						n := mod*8
//						dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//					_MM_UPCONV_PS_UINT16:
//						n := mod*16
//						dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//					_MM_UPCONV_PS_SINT16:
//						n := mod*16
//						dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//					ESAC
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPS, VBROADCASTF32X4, VBROADCASTSS'. Intrinsic: '_mm512_mask_extload_ps'.
// Requires KNCNI.
func MaskExtloadPs(src x86.M512, k x86.Mmask16, mt uintptr, conv MMUPCONVPSENUM, bc MMBROADCAST32ENUM, hint int) x86.M512 {
	return x86.M512(maskExtloadPs([16]float32(src), uint16(k), uintptr(mt), conv, bc, hint))
}

func maskExtloadPs(src [16]float32, k uint16, mt uintptr, conv MMUPCONVPSENUM, bc MMBROADCAST32ENUM, hint int) [16]float32


// ExtloadunpackhiEpi32: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64,
// up-converted depending on the value of 'conv', and expanded into packed
// 32-bit integers in 'dst'. The initial values of 'dst' are copied from 'src'.
// Only those converted doublewords that occur at or after the first
// 64-byte-aligned address following (mt-64) are loaded. Elements in the
// resulting vector that do not map to those doublewords are taken from 'src'.
// 'hint' indicates to the processor whether the loaded data is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt16ToInt32(MEM[addr + 2*offset])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt16ToInt32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*upSize % 64) == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*32
//				dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHD'. Intrinsic: '_mm512_extloadunpackhi_epi32'.
// Requires KNCNI.
func ExtloadunpackhiEpi32(src x86.M512i, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) x86.M512i {
	return x86.M512i(extloadunpackhiEpi32([64]byte(src), uintptr(mt), conv, hint))
}

func extloadunpackhiEpi32(src [64]byte, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte


// MaskExtloadunpackhiEpi32: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64,
// up-converted depending on the value of 'conv', and expanded into packed
// 32-bit integers in 'dst'. The initial values of 'dst' are copied from 'src'.
// Only those converted doublewords that occur at or after the first
// 64-byte-aligned address following (mt-64) are loaded. Elements in the
// resulting vector that do not map to those doublewords are taken from 'src'.
// 'hint' indicates to the processor whether the loaded data is non-temporal.
// Elements are copied to 'dst' according to element selector 'k' (elements are
// skipped when the corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt16ToInt32(MEM[addr + 2*offset])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt16ToInt32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*upSize % 64) == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*32
//					dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHD'. Intrinsic: '_mm512_mask_extloadunpackhi_epi32'.
// Requires KNCNI.
func MaskExtloadunpackhiEpi32(src x86.M512i, k x86.Mmask16, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) x86.M512i {
	return x86.M512i(maskExtloadunpackhiEpi32([64]byte(src), uint16(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackhiEpi32(src [64]byte, k uint16, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte


// ExtloadunpackhiEpi64: Loads the high-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt-64, up-converted depending on
// the value of 'conv', and expanded into packed 64-bit integers in 'dst'. The
// initial values of 'dst' are copied from 'src'. Only those converted
// quadwords that occur at or after the first 64-byte-aligned address following
// (mt-64) are loaded. Elements in the resulting vector that do not map to
// those quadwords are taken from 'src'. 'hint' indicates to the processor
// whether the loaded data is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*upSize) == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*64
//				dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHQ'. Intrinsic: '_mm512_extloadunpackhi_epi64'.
// Requires KNCNI.
func ExtloadunpackhiEpi64(src x86.M512i, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) x86.M512i {
	return x86.M512i(extloadunpackhiEpi64([64]byte(src), uintptr(mt), conv, hint))
}

func extloadunpackhiEpi64(src [64]byte, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte


// MaskExtloadunpackhiEpi64: Loads the high-64-byte-aligned portion of the
// quadword stream starting at element-aligned address mt-64, up-converted
// depending on the value of 'conv', and expanded into packed 64-bit integers
// in 'dst'. The initial values of 'dst' are copied from 'src'. Only those
// converted quadwords that occur at or after the first 64-byte-aligned address
// following (mt-64) are loaded. Elements in the resulting vector that do not
// map to those quadwords are taken from 'src'. 'hint' indicates to the
// processor whether the loaded data is non-temporal. Elements are copied to
// 'dst' according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*upSize) == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*64
//					dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHQ'. Intrinsic: '_mm512_mask_extloadunpackhi_epi64'.
// Requires KNCNI.
func MaskExtloadunpackhiEpi64(src x86.M512i, k x86.Mmask8, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) x86.M512i {
	return x86.M512i(maskExtloadunpackhiEpi64([64]byte(src), uint8(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackhiEpi64(src [64]byte, k uint8, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte


// ExtloadunpackhiPd: Loads the high-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt-64, up-converted depending on
// the value of 'conv', and expanded into packed double-precision (64-bit)
// floating-point values in 'dst'. The initial values of 'dst' are copied from
// 'src'. Only those converted quadwords that occur at or after the first
// 64-byte-aligned address following (mt-64) are loaded. Elements in the
// resulting vector that do not map to those quadwords are taken from 'src'.
// 'hint' indicates to the processor whether the loaded data is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE: RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE: RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*upSize) % 64 == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*64
//				dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPD'. Intrinsic: '_mm512_extloadunpackhi_pd'.
// Requires KNCNI.
func ExtloadunpackhiPd(src x86.M512d, mt uintptr, conv MMUPCONVPDENUM, hint int) x86.M512d {
	return x86.M512d(extloadunpackhiPd([8]float64(src), uintptr(mt), conv, hint))
}

func extloadunpackhiPd(src [8]float64, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64


// MaskExtloadunpackhiPd: Loads the high-64-byte-aligned portion of the
// quadword stream starting at element-aligned address mt-64, up-converted
// depending on the value of 'conv', and expanded into packed double-precision
// (64-bit) floating-point values in 'dst'. The initial values of 'dst' are
// copied from 'src'. Only those converted quadwords that occur at or after the
// first 64-byte-aligned address following (mt-64) are loaded. Elements in the
// resulting vector that do not map to those quadwords are taken from 'src'.
// 'hint' indicates to the processor whether the loaded data is non-temporal.
// Elements are copied to 'dst' according to element selector 'k' (elements are
// skipped when the corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE: RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE: RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*upSize) % 64 == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*64
//					dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPD'. Intrinsic: '_mm512_mask_extloadunpackhi_pd'.
// Requires KNCNI.
func MaskExtloadunpackhiPd(src x86.M512d, k x86.Mmask8, mt uintptr, conv MMUPCONVPDENUM, hint int) x86.M512d {
	return x86.M512d(maskExtloadunpackhiPd([8]float64(src), uint8(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackhiPd(src [8]float64, k uint8, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64


// ExtloadunpackhiPs: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64,
// up-converted depending on the value of 'conv', and expanded into packed
// single-precision (32-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted quadwords that
// occur at or after the first 64-byte-aligned address following (mt-64) are
// loaded. Elements in the resulting vector that do not map to those quadwords
// are taken from 'src'. 'hint' indicates to the processor whether the loaded
// data is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float16ToFloat32(MEM[addr + 4*offset])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt16ToFloat32(MEM[addr + 2*offset])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt16ToFloat32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*upSize % 64) == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*32
//				dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPS'. Intrinsic: '_mm512_extloadunpackhi_ps'.
// Requires KNCNI.
func ExtloadunpackhiPs(src x86.M512, mt uintptr, conv MMUPCONVPSENUM, hint int) x86.M512 {
	return x86.M512(extloadunpackhiPs([16]float32(src), uintptr(mt), conv, hint))
}

func extloadunpackhiPs(src [16]float32, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32


// MaskExtloadunpackhiPs: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64,
// up-converted depending on the value of 'conv', and expanded into packed
// single-precision (32-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted quadwords that
// occur at or after the first 64-byte-aligned address following (mt-64) are
// loaded. Elements in the resulting vector that do not map to those quadwords
// are taken from 'src'. 'hint' indicates to the processor whether the loaded
// data is non-temporal. Elements are copied to 'dst' according to element
// selector 'k' (elements are skipped when the corresponding mask bit is not
// set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float16ToFloat32(MEM[addr + 4*offset])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt16ToFloat32(MEM[addr + 2*offset])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt16ToFloat32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*upSize % 64) == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*32
//					dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPS'. Intrinsic: '_mm512_mask_extloadunpackhi_ps'.
// Requires KNCNI.
func MaskExtloadunpackhiPs(src x86.M512, k x86.Mmask16, mt uintptr, conv MMUPCONVPSENUM, hint int) x86.M512 {
	return x86.M512(maskExtloadunpackhiPs([16]float32(src), uint16(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackhiPs(src [16]float32, k uint16, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32


// ExtloadunpackloEpi32: Loads the low-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt,
// up-converted depending on the value of 'conv', and expanded into packed
// 32-bit integers in 'dst'. The initial values of 'dst' are copied from 'src'.
// Only those converted doublewords that occur before first 64-byte-aligned
// address following 'mt' are loaded. Elements in the resulting vector that do
// not map to those doublewords are taken from 'src'. 'hint' indicates to the
// processor whether the loaded data is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt16ToInt32(MEM[addr + 2*offset])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt16ToInt32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//			loadOffset := loadOffset + 1
//			IF (mt + loadOffset * upSize) % 64 == 0
//				break
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLD'. Intrinsic: '_mm512_extloadunpacklo_epi32'.
// Requires KNCNI.
func ExtloadunpackloEpi32(src x86.M512i, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) x86.M512i {
	return x86.M512i(extloadunpackloEpi32([64]byte(src), uintptr(mt), conv, hint))
}

func extloadunpackloEpi32(src [64]byte, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte


// MaskExtloadunpackloEpi32: Loads the low-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt,
// up-converted depending on the value of 'conv', and expanded into packed
// 32-bit integers in 'dst'. The initial values of 'dst' are copied from 'src'.
// Only those converted doublewords that occur before first 64-byte-aligned
// address following 'mt' are loaded. Elements in the resulting vector that do
// not map to those doublewords are taken from 'src'. 'hint' indicates to the
// processor whether the loaded data is non-temporal. Elements are copied to
// 'dst' according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt8ToInt32(MEM[addr + offset])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt16ToInt32(MEM[addr + 2*offset])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt16ToInt32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 15
//			IF k[j]
//				i := j*32
//				dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//				loadOffset := loadOffset + 1
//				IF (mt + loadOffset * upSize) % 64 == 0
//					break
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLD'. Intrinsic: '_mm512_mask_extloadunpacklo_epi32'.
// Requires KNCNI.
func MaskExtloadunpackloEpi32(src x86.M512i, k x86.Mmask16, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) x86.M512i {
	return x86.M512i(maskExtloadunpackloEpi32([64]byte(src), uint16(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackloEpi32(src [64]byte, k uint16, mt uintptr, conv MMUPCONVEPI32ENUM, hint int) [64]byte


// ExtloadunpackloEpi64: Loads the low-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt, up-converted depending on the
// value of 'conv', and expanded into packed 64-bit integers in 'dst'. The
// initial values of 'dst' are copied from 'src'. Only those converted quad
// that occur before first 64-byte-aligned address following 'mt' are loaded.
// Elements in the resulting vector that do not map to those quadwords are
// taken from 'src'. 'hint' indicates to the processor whether the loaded data
// is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//			loadOffset := loadOffset + 1
//			IF (addr + loadOffset*upSize % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLQ'. Intrinsic: '_mm512_extloadunpacklo_epi64'.
// Requires KNCNI.
func ExtloadunpackloEpi64(src x86.M512i, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) x86.M512i {
	return x86.M512i(extloadunpackloEpi64([64]byte(src), uintptr(mt), conv, hint))
}

func extloadunpackloEpi64(src [64]byte, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte


// MaskExtloadunpackloEpi64: Loads the low-64-byte-aligned portion of the
// quadword stream starting at element-aligned address mt, up-converted
// depending on the value of 'conv', and expanded into packed 64-bit integers
// in 'dst'. The initial values of 'dst' are copied from 'src'. Only those
// converted quadwords that occur before first 64-byte-aligned address
// following 'mt' are loaded. Elements in the resulting vector that do not map
// to those quadwords are taken from 'src'. 'hint' indicates to the processor
// whether the loaded data is non-temporal. Elements are copied to 'dst'
// according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE:   RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			IF k[j]
//				i := j*64
//				dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//				loadOffset := loadOffset + 1
//				IF (addr + loadOffset*upSize % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLQ'. Intrinsic: '_mm512_mask_extloadunpacklo_epi64'.
// Requires KNCNI.
func MaskExtloadunpackloEpi64(src x86.M512i, k x86.Mmask8, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) x86.M512i {
	return x86.M512i(maskExtloadunpackloEpi64([64]byte(src), uint8(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackloEpi64(src [64]byte, k uint8, mt uintptr, conv MMUPCONVEPI64ENUM, hint int) [64]byte


// ExtloadunpackloPd: Loads the low-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt, up-converted depending on the
// value of 'conv', and expanded into packed double-precision (64-bit)
// floating-point elements in 'dst'. The initial values of 'dst' are copied
// from 'src'. Only those converted quad that occur before first
// 64-byte-aligned address following 'mt' are loaded. Elements in the resulting
// vector that do not map to those quadwords are taken from 'src'. 'hint'
// indicates to the processor whether the loaded data is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE:	RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE:	RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//			loadOffset := loadOffset + 1
//			IF (mt + loadOffset * upSize) % 64 == 0
//				break
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPD'. Intrinsic: '_mm512_extloadunpacklo_pd'.
// Requires KNCNI.
func ExtloadunpackloPd(src x86.M512d, mt uintptr, conv MMUPCONVPDENUM, hint int) x86.M512d {
	return x86.M512d(extloadunpackloPd([8]float64(src), uintptr(mt), conv, hint))
}

func extloadunpackloPd(src [8]float64, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64


// MaskExtloadunpackloPd: Loads the low-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt, up-converted depending on the
// value of 'conv', and expanded into packed double-precision (64-bit)
// floating-point elements in 'dst'. The initial values of 'dst' are copied
// from 'src'. Only those converted quad that occur before first
// 64-byte-aligned address following 'mt' are loaded. Elements in the resulting
// vector that do not map to those quadwords are taken from 'src'. 'hint'
// indicates to the processor whether the loaded data is non-temporal. Elements
// are copied to 'dst' according to element selector 'k' (elemenst are skipped
// when the corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE:	RETURN MEM[addr + 8*offset]
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PD_NONE:	RETURN 8
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			IF k[j]
//				i := j*64
//				dst[i+63:i] := UPCONVERT(addr, loadOffset, conv)
//				loadOffset := loadOffset + 1
//				IF (mt + loadOffset * upSize) % 64 == 0
//					break
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPD'. Intrinsic: '_mm512_mask_extloadunpacklo_pd'.
// Requires KNCNI.
func MaskExtloadunpackloPd(src x86.M512d, k x86.Mmask8, mt uintptr, conv MMUPCONVPDENUM, hint int) x86.M512d {
	return x86.M512d(maskExtloadunpackloPd([8]float64(src), uint8(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackloPd(src [8]float64, k uint8, mt uintptr, conv MMUPCONVPDENUM, hint int) [8]float64


// ExtloadunpackloPs: Loads the low-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt,
// up-converted depending on the value of 'conv', and expanded into packed
// single-precision (32-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted doublewords that
// occur before first 64-byte-aligned address following 'mt' are loaded.
// Elements in the resulting vector that do not map to those doublewords are
// taken from 'src'. 'hint' indicates to the processor whether the loaded data
// is non-temporal. 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float16ToFloat32(MEM[addr + 4*offset])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt16ToFloat32(MEM[addr + 2*offset])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt16ToFloat32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//			loadOffset := loadOffset + 1
//			IF (mt + loadOffset * upSize) % 64 == 0
//				break
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPS'. Intrinsic: '_mm512_extloadunpacklo_ps'.
// Requires KNCNI.
func ExtloadunpackloPs(src x86.M512, mt uintptr, conv MMUPCONVPSENUM, hint int) x86.M512 {
	return x86.M512(extloadunpackloPs([16]float32(src), uintptr(mt), conv, hint))
}

func extloadunpackloPs(src [16]float32, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32


// MaskExtloadunpackloPs: Loads the low-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt,
// up-converted depending on the value of 'conv', and expanded into packed
// single-precision (32-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted doublewords that
// occur before first 64-byte-aligned address following 'mt' are loaded.
// Elements in the resulting vector that do not map to those doublewords are
// taken from 'src'. 'hint' indicates to the processor whether the loaded data
// is non-temporal. Elements are copied to 'dst' according to element selector
// 'k' (elements are skipped when the corresponding mask bit is not set). 
//
//		UPCONVERT(address, offset, convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:    RETURN MEM[addr + 4*offset]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float16ToFloat32(MEM[addr + 4*offset])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt8ToFloat32(MEM[addr + offset])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt16ToFloat32(MEM[addr + 2*offset])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt16ToFloat32(MEM[addr + 2*offset])
//			ESAC
//		}
//		
//		UPCONVERTSIZE(convertTo) {
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:	   RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		upSize := UPCONVERTSIZE(conv)
//		addr = MEM[mt]
//		FOR j := 0 to 15
//			IF k[j]
//				i := j*32
//				dst[i+31:i] := UPCONVERT(addr, loadOffset, conv)
//				loadOffset := loadOffset + 1
//				IF (mt + loadOffset * upSize) % 64 == 0
//					break
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPS'. Intrinsic: '_mm512_mask_extloadunpacklo_ps'.
// Requires KNCNI.
func MaskExtloadunpackloPs(src x86.M512, k x86.Mmask16, mt uintptr, conv MMUPCONVPSENUM, hint int) x86.M512 {
	return x86.M512(maskExtloadunpackloPs([16]float32(src), uint16(k), uintptr(mt), conv, hint))
}

func maskExtloadunpackloPs(src [16]float32, k uint16, mt uintptr, conv MMUPCONVPSENUM, hint int) [16]float32


// ExtpackstorehiEpi32: Down-converts and stores packed 32-bit integer elements
// of 'v1' into a byte/word/doubleword stream according to 'conv' at a
// logically mapped starting address (mt-64), storing the high-64-byte elements
// of that stream (those elemetns of the stream that map at or after the first
// 64-byte-aligned address following (m5-64)). 'hint' indicates to the
// processor whether the data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN element[i+31:i]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*32
//				tmp := DOWNCONVERT(v1[i+31:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				4: MEM[storeAddr] := tmp[31:0]
//				2: MEM[storeAddr] := tmp[15:0]
//				1: MEM[storeAddr] := tmp[7:0]
//				ESAC
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHD'. Intrinsic: '_mm512_extpackstorehi_epi32'.
// Requires KNCNI.
func ExtpackstorehiEpi32(mt uintptr, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, hint int)  {
	extpackstorehiEpi32(uintptr(mt), [64]byte(v1), conv, hint)
}

func extpackstorehiEpi32(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 


// MaskExtpackstorehiEpi32: Down-converts and stores packed 32-bit integer
// elements of 'v1' into a byte/word/doubleword stream according to 'conv' at a
// logically mapped starting address (mt-64), storing the high-64-byte elements
// of that stream (those elemetns of the stream that map at or after the first
// 64-byte-aligned address following (m5-64)). 'hint' indicates to the
// processor whether the data is non-temporal. Elements are stored to memory
// according to element selector 'k' (elements are skipped when the
// corresonding mask bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN element[i+31:i]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*32
//					tmp := DOWNCONVERT(v1[i+31:i], conv)
//					storeAddr := addr + storeOffset * downSize
//					CASE downSize OF
//					4: MEM[storeAddr] := tmp[31:0]
//					2: MEM[storeAddr] := tmp[15:0]
//					1: MEM[storeAddr] := tmp[7:0]
//					ESAC
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHD'. Intrinsic: '_mm512_mask_extpackstorehi_epi32'.
// Requires KNCNI.
func MaskExtpackstorehiEpi32(mt uintptr, k x86.Mmask16, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, hint int)  {
	maskExtpackstorehiEpi32(uintptr(mt), uint16(k), [64]byte(v1), conv, hint)
}

func maskExtpackstorehiEpi32(mt uintptr, k uint16, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 


// ExtpackstorehiEpi64: Down-converts and stores packed 64-bit integer elements
// of 'v1' into a quadword stream according to 'conv' at a logically mapped
// starting address (mt-64), storing the high-64-byte elements of that stream
// (those elemetns of the stream that map at or after the first 64-byte-aligned
// address following (m5-64)). 'hint' indicates to the processor whether the
// data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*64
//				tmp := DOWNCONVERT(v1[i+63:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				8: MEM[storeAddr] := tmp[63:0]
//				ESAC
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHQ'. Intrinsic: '_mm512_extpackstorehi_epi64'.
// Requires KNCNI.
func ExtpackstorehiEpi64(mt uintptr, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, hint int)  {
	extpackstorehiEpi64(uintptr(mt), [64]byte(v1), conv, hint)
}

func extpackstorehiEpi64(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 


// MaskExtpackstorehiEpi64: Down-converts and stores packed 64-bit integer
// elements of 'v1' into a quadword stream according to 'conv' at a logically
// mapped starting address (mt-64), storing the high-64-byte elements of that
// stream (those elemetns of the stream that map at or after the first
// 64-byte-aligned address following (mt-64)). 'hint' indicates to the
// processor whether the data is non-temporal. Elements are stored to memory
// according to element selector 'k' (elements are skipped when the
// corresonding mask bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*64
//					tmp := DOWNCONVERT(v1[i+63:i], conv)
//					storeAddr := addr + storeOffset * downSize
//					CASE downSize OF
//					8: MEM[storeAddr] := tmp[63:0]
//					ESAC
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHQ'. Intrinsic: '_mm512_mask_extpackstorehi_epi64'.
// Requires KNCNI.
func MaskExtpackstorehiEpi64(mt uintptr, k x86.Mmask8, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, hint int)  {
	maskExtpackstorehiEpi64(uintptr(mt), uint8(k), [64]byte(v1), conv, hint)
}

func maskExtpackstorehiEpi64(mt uintptr, k uint8, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 


// ExtpackstorehiPd: Down-converts and stores packed double-precision (64-bit)
// floating-point elements of 'v1' into a quadword stream according to 'conv'
// at a logically mapped starting address (mt-64), storing the high-64-byte
// elements of that stream (those elemetns of the stream that map at or after
// the first 64-byte-aligned address following (m5-64)). 'hint' indicates to
// the processor whether the data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*64
//				tmp := DOWNCONVERT(v1[i+63:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				8: MEM[storeAddr] := tmp[63:0]
//				ESAC
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPD'. Intrinsic: '_mm512_extpackstorehi_pd'.
// Requires KNCNI.
func ExtpackstorehiPd(mt uintptr, v1 x86.M512d, conv MMDOWNCONVPDENUM, hint int)  {
	extpackstorehiPd(uintptr(mt), [8]float64(v1), conv, hint)
}

func extpackstorehiPd(mt uintptr, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 


// MaskExtpackstorehiPd: Down-converts and stores packed double-precision
// (64-bit) floating-point elements of 'v1' into a quadword stream according to
// 'conv' at a logically mapped starting address (mt-64), storing the
// high-64-byte elements of that stream (those elemetns of the stream that map
// at or after the first 64-byte-aligned address following (m5-64)). 'hint'
// indicates to the processor whether the data is non-temporal. Elements are
// stored to memory according to element selector 'k' (elements are skipped
// when the corresponding mask bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*64
//					tmp := DOWNCONVERT(v1[i+63:i], conv)
//					storeAddr := addr + storeOffset * downSize
//					CASE downSize OF
//					8: MEM[storeAddr] := tmp[63:0]
//					ESAC
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPD'. Intrinsic: '_mm512_mask_extpackstorehi_pd'.
// Requires KNCNI.
func MaskExtpackstorehiPd(mt uintptr, k x86.Mmask8, v1 x86.M512d, conv MMDOWNCONVPDENUM, hint int)  {
	maskExtpackstorehiPd(uintptr(mt), uint8(k), [8]float64(v1), conv, hint)
}

func maskExtpackstorehiPd(mt uintptr, k uint8, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 


// ExtpackstorehiPs: Down-converts and stores packed single-precision (32-bit)
// floating-point elements of 'v1' into a byte/word/doubleword stream according
// to 'conv' at a logically mapped starting address (mt-64), storing the
// high-64-byte elements of that stream (those elemetns of the stream that map
// at or after the first 64-byte-aligned address following (m5-64)). 'hint'
// indicates to the processor whether the data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:	   RETURN element[i+31:i]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float32ToFloat16(element[i+31:i])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:	   RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*32
//				tmp := DOWNCONVERT(v1[i+31:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				4: MEM[storeAddr] := tmp[31:0]
//				2: MEM[storeAddr] := tmp[15:0]
//				1: MEM[storeAddr] := tmp[7:0]
//				ESAC
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPS'. Intrinsic: '_mm512_extpackstorehi_ps'.
// Requires KNCNI.
func ExtpackstorehiPs(mt uintptr, v1 x86.M512, conv MMDOWNCONVPSENUM, hint int)  {
	extpackstorehiPs(uintptr(mt), [16]float32(v1), conv, hint)
}

func extpackstorehiPs(mt uintptr, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 


// MaskExtpackstorehiPs: Down-converts and stores packed single-precision
// (32-bit) floating-point elements of 'v1' into a byte/word/doubleword stream
// according to 'conv' at a logically mapped starting address (mt-64), storing
// the high-64-byte elements of that stream (those elemetns of the stream that
// map at or after the first 64-byte-aligned address following (m5-64)). 'hint'
// indicates to the processor whether the data is non-temporal. Elements are
// stored to memory according to element selector 'k' (elements are skipped
// when the corresponding mask bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:	   RETURN element[i+31:i]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float32ToFloat16(element[i+31:i])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:    RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*downSize) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*32
//					tmp := DOWNCONVERT(v1[i+31:i], conv)
//					storeAddr := addr + storeOffset * downSize
//					CASE downSize OF
//					4: MEM[storeAddr] := tmp[31:0]
//					2: MEM[storeAddr] := tmp[15:0]
//					1: MEM[storeAddr] := tmp[7:0]
//					ESAC
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPS'. Intrinsic: '_mm512_mask_extpackstorehi_ps'.
// Requires KNCNI.
func MaskExtpackstorehiPs(mt uintptr, k x86.Mmask16, v1 x86.M512, conv MMDOWNCONVPSENUM, hint int)  {
	maskExtpackstorehiPs(uintptr(mt), uint16(k), [16]float32(v1), conv, hint)
}

func maskExtpackstorehiPs(mt uintptr, k uint16, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 


// ExtpackstoreloEpi32: Down-converts and stores packed 32-bit integer elements
// of 'v1' into a byte/word/doubleword stream according to 'conv' at a
// logically mapped starting address 'mt', storing the low-64-byte elements of
// that stream (those elements of the stream that map before the first
// 64-byte-aligned address follwing 'mt'). 'hint' indicates to the processor
// whether the data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN element[i+31:i]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			tmp := DOWNCONVERT(v1[i+31:i], conv)
//			storeAddr := addr + storeOffset * downSize
//			CASE downSize OF
//			4: MEM[storeAddr] := tmp[31:0]
//			2: MEM[storeAddr] := tmp[15:0]
//			1: MEM[storeAddr] := tmp[7:0]
//			ESAC
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset * downSize) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELD'. Intrinsic: '_mm512_extpackstorelo_epi32'.
// Requires KNCNI.
func ExtpackstoreloEpi32(mt uintptr, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, hint int)  {
	extpackstoreloEpi32(uintptr(mt), [64]byte(v1), conv, hint)
}

func extpackstoreloEpi32(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 


// MaskExtpackstoreloEpi32: Down-converts and stores packed 32-bit integer
// elements of 'v1' into a byte/word/doubleword stream according to 'conv' at a
// logically mapped starting address 'mt', storing the low-64-byte elements of
// that stream (those elements of the stream that map before the first
// 64-byte-aligned address follwing 'mt'). 'hint' indicates to the processor
// whether the data is non-temporal. Elements are written to memory according
// to element selector 'k' (elements are skipped when the corresponding mask
// bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN element[i+31:i]
//			_MM_UPCONV_EPI32_UINT8:  RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT8:  RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_EPI32_UINT16: RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_EPI32_SINT16: RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI32_NONE:   RETURN 4
//			_MM_UPCONV_EPI32_UINT8:  RETURN 1
//			_MM_UPCONV_EPI32_SINT8:  RETURN 1
//			_MM_UPCONV_EPI32_UINT16: RETURN 2
//			_MM_UPCONV_EPI32_SINT16: RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 15
//			IF k[j]
//				i := j*32
//				tmp := DOWNCONVERT(v1[i+31:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				4: MEM[storeAddr] := tmp[31:0]
//				2: MEM[storeAddr] := tmp[15:0]
//				1: MEM[storeAddr] := tmp[7:0]
//				ESAC
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset * downSize) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELD'. Intrinsic: '_mm512_mask_extpackstorelo_epi32'.
// Requires KNCNI.
func MaskExtpackstoreloEpi32(mt uintptr, k x86.Mmask16, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, hint int)  {
	maskExtpackstoreloEpi32(uintptr(mt), uint16(k), [64]byte(v1), conv, hint)
}

func maskExtpackstoreloEpi32(mt uintptr, k uint16, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 


// ExtpackstoreloEpi64: Down-converts and stores packed 64-bit integer elements
// of 'v1' into a quadword stream according to 'conv' at a logically mapped
// starting address 'mt', storing the low-64-byte elements of that stream
// (those elements of the stream that map before the first 64-byte-aligned
// address follwing 'mt'). 'hint' indicates to the processor whether the data
// is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			i := j*63
//			tmp := DOWNCONVERT(v1[i+63:i], conv)
//			storeAddr := addr + storeOffset * downSize
//			CASE downSize OF
//			8: MEM[storeAddr] := tmp[63:0]
//			ESAC
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset * downSize) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELQ'. Intrinsic: '_mm512_extpackstorelo_epi64'.
// Requires KNCNI.
func ExtpackstoreloEpi64(mt uintptr, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, hint int)  {
	extpackstoreloEpi64(uintptr(mt), [64]byte(v1), conv, hint)
}

func extpackstoreloEpi64(mt uintptr, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 


// MaskExtpackstoreloEpi64: Down-converts and stores packed 64-bit integer
// elements of 'v1' into a quadword stream according to 'conv' at a logically
// mapped starting address 'mt', storing the low-64-byte elements of that
// stream (those elements of the stream that map before the first
// 64-byte-aligned address follwing 'mt'). 'hint' indicates to the processor
// whether the data is non-temporal. Elements are stored to memory according to
// element selector 'k' (elements are skipped whent he corresponding mask bit
// is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_EPI64_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			IF k[j]
//				i := j*63
//				tmp := DOWNCONVERT(v1[i+63:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				8: MEM[storeAddr] := tmp[63:0]
//				ESAC
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset * downSize) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELQ'. Intrinsic: '_mm512_mask_extpackstorelo_epi64'.
// Requires KNCNI.
func MaskExtpackstoreloEpi64(mt uintptr, k x86.Mmask8, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, hint int)  {
	maskExtpackstoreloEpi64(uintptr(mt), uint8(k), [64]byte(v1), conv, hint)
}

func maskExtpackstoreloEpi64(mt uintptr, k uint8, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 


// ExtpackstoreloPd: Down-converts and stores packed double-precision (64-bit)
// floating-point elements of 'v1' into a quadword stream according to 'conv'
// at a logically mapped starting address 'mt', storing the low-64-byte
// elements of that stream (those elements of the stream that map before the
// first 64-byte-aligned address follwing 'mt'). 'hint' indicates to the
// processor whether the data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			i := j*63
//			tmp := DOWNCONVERT(v1[i+63:i], conv)
//			storeAddr := addr + storeOffset * downSize
//			CASE downSize OF
//			8: MEM[storeAddr] := tmp[63:0]
//			ESAC
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset * downSize) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPD'. Intrinsic: '_mm512_extpackstorelo_pd'.
// Requires KNCNI.
func ExtpackstoreloPd(mt uintptr, v1 x86.M512d, conv MMDOWNCONVPDENUM, hint int)  {
	extpackstoreloPd(uintptr(mt), [8]float64(v1), conv, hint)
}

func extpackstoreloPd(mt uintptr, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 


// MaskExtpackstoreloPd: Down-converts and stores packed double-precision
// (64-bit) floating-point elements of 'v1' into a quadword stream according to
// 'conv' at a logically mapped starting address 'mt', storing the low-64-byte
// elements of that stream (those elements of the stream that map before the
// first 64-byte-aligned address follwing 'mt'). 'hint' indicates to the
// processor whether the data is non-temporal. Elements are stored to memory
// according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN element[i+63:i]
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PD_NONE: RETURN 8
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 7
//			IF k[j]
//				i := j*63
//				tmp := DOWNCONVERT(v1[i+63:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				8: MEM[storeAddr] := tmp[63:0]
//				ESAC
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset * downSize) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPD'. Intrinsic: '_mm512_mask_extpackstorelo_pd'.
// Requires KNCNI.
func MaskExtpackstoreloPd(mt uintptr, k x86.Mmask8, v1 x86.M512d, conv MMDOWNCONVPDENUM, hint int)  {
	maskExtpackstoreloPd(uintptr(mt), uint8(k), [8]float64(v1), conv, hint)
}

func maskExtpackstoreloPd(mt uintptr, k uint8, v1 [8]float64, conv MMDOWNCONVPDENUM, hint int) 


// ExtpackstoreloPs: Down-converts and stores packed single-precision (32-bit)
// floating-point elements of 'v1' into a byte/word/doubleword stream according
// to 'conv' at a logically mapped starting address 'mt', storing the
// low-64-byte elements of that stream (those elements of the stream that map
// before the first 64-byte-aligned address follwing 'mt'). 'hint' indicates to
// the processor whether the data is non-temporal. 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:	   RETURN element[i+31:i]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float32ToFloat16(element[i+31:i])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:	   RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			tmp := DOWNCONVERT(v1[i+31:i], conv)
//			storeAddr := addr + storeOffset * downSize
//			CASE downSize OF
//			4: MEM[storeAddr] := tmp[31:0]
//			2: MEM[storeAddr] := tmp[15:0]
//			1: MEM[storeAddr] := tmp[7:0]
//			ESAC
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset * downSize) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPS'. Intrinsic: '_mm512_extpackstorelo_ps'.
// Requires KNCNI.
func ExtpackstoreloPs(mt uintptr, v1 x86.M512, conv MMDOWNCONVPSENUM, hint int)  {
	extpackstoreloPs(uintptr(mt), [16]float32(v1), conv, hint)
}

func extpackstoreloPs(mt uintptr, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 


// MaskExtpackstoreloPs: Down-converts and stores packed single-precision
// (32-bit) floating-point elements of 'v1' into a byte/word/doubleword stream
// according to 'conv' at a logically mapped starting address 'mt', storing the
// low-64-byte elements of that stream (those elements of the stream that map
// before the first 64-byte-aligned address follwing 'mt'). 'hint' indicates to
// the processor whether the data is non-temporal. Elements are stored to
// memory according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		DOWNCONVERT(element, convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:	   RETURN element[i+31:i]
//			_MM_UPCONV_PS_FLOAT16: RETURN Float32ToFloat16(element[i+31:i])
//			_MM_UPCONV_PS_UINT8:   RETURN UInt32ToUInt8(element[i+31:i])
//			_MM_UPCONV_PS_SINT8:   RETURN SInt32ToSInt8(element[i+31:i])
//			_MM_UPCONV_PS_UINT16:  RETURN UInt32ToUInt16(element[i+31:i])
//			_MM_UPCONV_PS_SINT16:  RETURN SInt32ToSInt16(element[i+31:i])
//			ESAC
//		}
//		
//		DOWNCONVERTSIZE(convertTo) {
//			CASE converTo OF
//			_MM_UPCONV_PS_NONE:    RETURN 4
//			_MM_UPCONV_PS_FLOAT16: RETURN 2
//			_MM_UPCONV_PS_UINT8:   RETURN 1
//			_MM_UPCONV_PS_SINT8:   RETURN 1
//			_MM_UPCONV_PS_UINT16:  RETURN 2
//			_MM_UPCONV_PS_SINT16:  RETURN 2
//			ESAC
//		}
//		
//		storeOffset := 0
//		downSize := DOWNCONVERTSIZE(conv)
//		addr = mt
//		FOR j := 0 to 15
//			IF k[j]
//				i := j*32
//				tmp := DOWNCONVERT(v1[i+31:i], conv)
//				storeAddr := addr + storeOffset * downSize
//				CASE downSize OF
//				4: MEM[storeAddr] := tmp[31:0]
//				2: MEM[storeAddr] := tmp[15:0]
//				1: MEM[storeAddr] := tmp[7:0]
//				ESAC
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset * downSize) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPS'. Intrinsic: '_mm512_mask_extpackstorelo_ps'.
// Requires KNCNI.
func MaskExtpackstoreloPs(mt uintptr, k x86.Mmask16, v1 x86.M512, conv MMDOWNCONVPSENUM, hint int)  {
	maskExtpackstoreloPs(uintptr(mt), uint16(k), [16]float32(v1), conv, hint)
}

func maskExtpackstoreloPs(mt uintptr, k uint16, v1 [16]float32, conv MMDOWNCONVPSENUM, hint int) 


// ExtstoreEpi32: Downconverts packed 32-bit integer elements stored in 'v' to
// a smaller type depending on 'conv' and stores them in memory location 'mt'.
// 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			CASE conv OF
//			_MM_DOWNCONV_EPI32_NONE:
//				addr[i+31:i] := v[i+31:i]
//			_MM_DOWNCONV_EPI32_UINT8:
//				n := j*8
//				addr[n+7:n] := Int32ToUInt8(v[i+31:i])
//			_MM_DOWNCONV_EPI32_SINT8:
//				n := j*8
//				addr[n+7:n] := Int32ToSInt8(v[i+31:i])
//			_MM_DOWNCONV_EPI32_UINT16:
//				n := j*16
//				addr[n+15:n] := Int32ToUInt16(v[i+31:i])
//			_MM_DOWNCONV_EPI32_SINT16:
//				n := j*16
//				addr[n+15:n] := Int32ToSInt16(v[i+31:i])
//			ESAC
//		ENDFOR
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_extstore_epi32'.
// Requires KNCNI.
func ExtstoreEpi32(mt uintptr, v x86.M512i, conv MMDOWNCONVEPI32ENUM, hint int)  {
	extstoreEpi32(uintptr(mt), [64]byte(v), conv, hint)
}

func extstoreEpi32(mt uintptr, v [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 


// MaskExtstoreEpi32: Downconverts packed 32-bit integer elements stored in 'v'
// to a smaller type depending on 'conv' and stores them in memory location
// 'mt' (elements in 'mt' are unaltered when the corresponding mask bit is not
// set). 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_EPI32_NONE:
//					addr[i+31:i] := v[i+31:i]
//				_MM_DOWNCONV_EPI32_UINT8:
//					n := j*8
//					addr[n+7:n] := Int32ToUInt8(v[i+31:i])
//				_MM_DOWNCONV_EPI32_SINT8:
//					n := j*8
//					addr[n+7:n] := Int32ToSInt8(v[i+31:i])
//				_MM_DOWNCONV_EPI32_UINT16:
//					n := j*16
//					addr[n+15:n] := Int32ToUInt16(v[i+31:i])
//				_MM_DOWNCONV_EPI32_SINT16:
//					n := j*16
//					addr[n+15:n] := Int32ToSInt16(v[i+31:i])
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_mask_extstore_epi32'.
// Requires KNCNI.
func MaskExtstoreEpi32(mt uintptr, k x86.Mmask16, v x86.M512i, conv MMDOWNCONVEPI32ENUM, hint int)  {
	maskExtstoreEpi32(uintptr(mt), uint16(k), [64]byte(v), conv, hint)
}

func maskExtstoreEpi32(mt uintptr, k uint16, v [64]byte, conv MMDOWNCONVEPI32ENUM, hint int) 


// ExtstoreEpi64: Downconverts packed 64-bit integer elements stored in 'v' to
// a smaller type depending on 'conv' and stores them in memory location 'mt'.
// 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_EPI64_NONE: addr[i+63:i] := v[i+63:i]
//			ESAC
//		ENDFOR
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_extstore_epi64'.
// Requires KNCNI.
func ExtstoreEpi64(mt uintptr, v x86.M512i, conv MMDOWNCONVEPI64ENUM, hint int)  {
	extstoreEpi64(uintptr(mt), [64]byte(v), conv, hint)
}

func extstoreEpi64(mt uintptr, v [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 


// MaskExtstoreEpi64: Downconverts packed 64-bit integer elements stored in 'v'
// to a smaller type depending on 'conv' and stores them in memory location
// 'mt' (elements in 'mt' are unaltered when the corresponding mask bit is not
// set). 'hint' indicates to the processor whether the data is non-temporal. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_EPI64_NONE: addr[i+63:i] := v[i+63:i]
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_mask_extstore_epi64'.
// Requires KNCNI.
func MaskExtstoreEpi64(mt uintptr, k x86.Mmask8, v x86.M512i, conv MMDOWNCONVEPI64ENUM, hint int)  {
	maskExtstoreEpi64(uintptr(mt), uint8(k), [64]byte(v), conv, hint)
}

func maskExtstoreEpi64(mt uintptr, k uint8, v [64]byte, conv MMDOWNCONVEPI64ENUM, hint int) 


// ExtstorePd: Downconverts packed double-precision (64-bit) floating-point
// elements stored in 'v' to a smaller type depending on 'conv' and stores them
// in memory location 'mt'. 'hint' indicates to the processor whether the data
// is non-temporal. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_PS_NONE:
//				addr[i+63:i] := v[i+63:i]
//			ESAC
//		ENDFOR
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_extstore_pd'.
// Requires KNCNI.
func ExtstorePd(mt uintptr, v x86.M512d, conv MMDOWNCONVPDENUM, hint int)  {
	extstorePd(uintptr(mt), [8]float64(v), conv, hint)
}

func extstorePd(mt uintptr, v [8]float64, conv MMDOWNCONVPDENUM, hint int) 


// MaskExtstorePd: Downconverts packed double-precision (64-bit) floating-point
// elements stored in 'v' to a smaller type depending on 'conv' and stores them
// in memory location 'mt' (elements in 'mt' are unaltered when the
// corresponding mask bit is not set). 'hint' indicates to the processor
// whether the data is non-temporal. 
//
//		addr := MEM[mt]		
//		FOR j := 0 to 7
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_PD_NONE:
//				IF k[j]
//					mt[i+63:i] := v[i+63:i]
//				FI
//			ESAC
//		ENDFOR
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_mask_extstore_pd'.
// Requires KNCNI.
func MaskExtstorePd(mt uintptr, k x86.Mmask8, v x86.M512d, conv MMDOWNCONVPDENUM, hint int)  {
	maskExtstorePd(uintptr(mt), uint8(k), [8]float64(v), conv, hint)
}

func maskExtstorePd(mt uintptr, k uint8, v [8]float64, conv MMDOWNCONVPDENUM, hint int) 


// ExtstorePs: Downconverts packed single-precision (32-bit) floating-point
// elements stored in 'v' to a smaller type depending on 'conv' and stores them
// in memory location 'mt'. 'hint' indicates to the processor whether the data
// is non-temporal. 
//
//		addr := MEM[mt]		
//		FOR j := 0 to 15
//			i := j*32
//			CASE conv OF
//			_MM_DOWNCONV_PS_NONE:
//				addr[i+31:i] := v[i+31:i]
//			_MM_DOWNCONV_PS_FLOAT16:
//				n := j*16
//				addr[n+15:n] := Float32ToFloat16(v[i+31:i])
//			_MM_DOWNCONV_PS_UINT8:
//				n := j*8
//				addr[n+7:n] := Float32ToUInt8(v[i+31:i])
//			_MM_DOWNCONV_PS_SINT8:
//				n := j*8
//				addr[n+7:n] := Float32ToSInt8(v[i+31:i])
//			_MM_DOWNCONV_PS_UINT16:
//				n := j*16
//				addr[n+15:n] := Float32ToUInt16(v[i+31:i])
//			_MM_DOWNCONV_PS_SINT16:
//				n := j*16
//				addr[n+15:n] := Float32ToSInt16(v[i+31:i])
//			ESAC
//		ENDFOR
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_extstore_ps'.
// Requires KNCNI.
func ExtstorePs(mt uintptr, v x86.M512, conv MMDOWNCONVPSENUM, hint int)  {
	extstorePs(uintptr(mt), [16]float32(v), conv, hint)
}

func extstorePs(mt uintptr, v [16]float32, conv MMDOWNCONVPSENUM, hint int) 


// MaskExtstorePs: Downconverts packed single-precision (32-bit) floating-point
// elements stored in 'v' to a smaller type depending on 'conv' and stores them
// in memory location 'mt' using writemask 'k' (elements are not written to
// memory when the corresponding mask bit is not set). 'hint' indicates to the
// processor whether the data is non-temporal. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_PS_NONE:
//					mt[i+31:i] := v[i+31:i]
//				_MM_DOWNCONV_PS_FLOAT16:
//					n := j*16
//					mt[n+15:n] := Float32ToFloat16(v[i+31:i])
//				_MM_DOWNCONV_PS_UINT8:
//					n := j*8
//					mt[n+7:n] := Float32ToUInt8(v[i+31:i])
//				_MM_DOWNCONV_PS_SINT8:
//					n := j*8
//					mt[n+7:n] := Float32ToSInt8(v[i+31:i])
//				_MM_DOWNCONV_PS_UINT16:
//					n := j*16
//					mt[n+15:n] := Float32ToUInt16(v[i+31:i])
//				_MM_DOWNCONV_PS_SINT16:
//					n := j*16
//					mt[n+15:n] := Float32ToSInt16(v[i+31:i])
//				ESAC
//			 FI
//		ENDFOR
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_mask_extstore_ps'.
// Requires KNCNI.
func MaskExtstorePs(mt uintptr, k x86.Mmask16, v x86.M512, conv MMDOWNCONVPSENUM, hint int)  {
	maskExtstorePs(uintptr(mt), uint16(k), [16]float32(v), conv, hint)
}

func maskExtstorePs(mt uintptr, k uint16, v [16]float32, conv MMDOWNCONVPSENUM, hint int) 


// FixupnanPd: Fixes up NaN's from packed double-precision (64-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' and
// storing the quietized NaN's from 'v1' in 'v3'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := FixupNaNs(v1[i+63:i], v2[i+63:i])
//			v3[i+63:i] := QuietizeNaNs(v1[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPD'. Intrinsic: '_mm512_fixupnan_pd'.
// Requires KNCNI.
func FixupnanPd(v1 x86.M512d, v2 x86.M512d, v3 x86.M512i) x86.M512d {
	return x86.M512d(fixupnanPd([8]float64(v1), [8]float64(v2), [64]byte(v3)))
}

func fixupnanPd(v1 [8]float64, v2 [8]float64, v3 [64]byte) [8]float64


// MaskFixupnanPd: Fixes up NaN's from packed double-precision (64-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' using
// writemask 'k' (only elements whose corresponding mask bit is set are used in
// the computation). Quietized NaN's from 'v1' are stored in 'v3'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := FixupNaNs(v1[i+63:i], v2[i+63:i])
//				v3[i+63:i] := QuietizeNaNs(v1[i+63:i])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPD'. Intrinsic: '_mm512_mask_fixupnan_pd'.
// Requires KNCNI.
func MaskFixupnanPd(v1 x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512i) x86.M512d {
	return x86.M512d(maskFixupnanPd([8]float64(v1), uint8(k), [8]float64(v2), [64]byte(v3)))
}

func maskFixupnanPd(v1 [8]float64, k uint8, v2 [8]float64, v3 [64]byte) [8]float64


// FixupnanPs: Fixes up NaN's from packed single-precision (32-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' and
// storing the quietized NaN's from 'v1' in 'v3'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FixupNaNs(v1[i+31:i], v2[i+31:i])
//			v3[i+31:i] := QuietizeNaNs(v1[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPS'. Intrinsic: '_mm512_fixupnan_ps'.
// Requires KNCNI.
func FixupnanPs(v1 x86.M512, v2 x86.M512, v3 x86.M512i) x86.M512 {
	return x86.M512(fixupnanPs([16]float32(v1), [16]float32(v2), [64]byte(v3)))
}

func fixupnanPs(v1 [16]float32, v2 [16]float32, v3 [64]byte) [16]float32


// MaskFixupnanPs: Fixes up NaN's from packed single-precision (32-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' using
// writemask 'k' (only elements whose corresponding mask bit is set are used in
// the computation). Quietized NaN's from 'v1' are stored in 'v3'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FixupNaNs(v1[i+31:i], v2[i+31:i])
//				v3[i+31:i] := QuietizeNaNs(v1[i+31:i])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPS'. Intrinsic: '_mm512_mask_fixupnan_ps'.
// Requires KNCNI.
func MaskFixupnanPs(v1 x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512i) x86.M512 {
	return x86.M512(maskFixupnanPs([16]float32(v1), uint16(k), [16]float32(v2), [64]byte(v3)))
}

func maskFixupnanPs(v1 [16]float32, k uint16, v2 [16]float32, v3 [64]byte) [16]float32


// FmaddEpi32: Multiply packed 32-bit integer elements in 'a' and 'b', add the
// intermediate result to packed elements in 'c' and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD231D'. Intrinsic: '_mm512_fmadd_epi32'.
// Requires KNCNI.
func FmaddEpi32(a x86.M512i, b x86.M512i, c x86.M512i) x86.M512i {
	return x86.M512i(fmaddEpi32([64]byte(a), [64]byte(b), [64]byte(c)))
}

func fmaddEpi32(a [64]byte, b [64]byte, c [64]byte) [64]byte


// MaskFmaddEpi32: Multiply packed 32-bit integer elements in 'a' and 'b', add
// the intermediate result to packed elements in 'c' and store the results in
// 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD231D'. Intrinsic: '_mm512_mask_fmadd_epi32'.
// Requires KNCNI.
func MaskFmaddEpi32(a x86.M512i, k x86.Mmask16, b x86.M512i, c x86.M512i) x86.M512i {
	return x86.M512i(maskFmaddEpi32([64]byte(a), uint16(k), [64]byte(b), [64]byte(c)))
}

func maskFmaddEpi32(a [64]byte, k uint16, b [64]byte, c [64]byte) [64]byte


// Mask3FmaddEpi32: Multiply packed 32-bit integer elements in 'a' and 'b', add
// the intermediate result to packed elements in 'c' and store the results in
// 'dst' using writemask 'k' (elements are copied from 'c' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD231D'. Intrinsic: '_mm512_mask3_fmadd_epi32'.
// Requires KNCNI.
func Mask3FmaddEpi32(a x86.M512i, b x86.M512i, c x86.M512i, k x86.Mmask16) x86.M512i {
	return x86.M512i(mask3FmaddEpi32([64]byte(a), [64]byte(b), [64]byte(c), uint16(k)))
}

func mask3FmaddEpi32(a [64]byte, b [64]byte, c [64]byte, k uint16) [64]byte


// FmaddPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', add the intermediate result to packed elements in 'c', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_fmadd_pd'.
// Requires KNCNI.
func FmaddPd(a x86.M512d, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(fmaddPd([8]float64(a), [8]float64(b), [8]float64(c)))
}

func fmaddPd(a [8]float64, b [8]float64, c [8]float64) [8]float64


// MaskFmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask_fmadd_pd'.
// Requires KNCNI.
func MaskFmaddPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(maskFmaddPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c)))
}

func maskFmaddPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64


// Mask3FmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask3_fmadd_pd'.
// Requires KNCNI.
func Mask3FmaddPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) x86.M512d {
	return x86.M512d(mask3FmaddPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k)))
}

func mask3FmaddPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64


// FmaddPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', add the intermediate result to packed elements in 'c', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_fmadd_ps'.
// Requires KNCNI.
func FmaddPs(a x86.M512, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(fmaddPs([16]float32(a), [16]float32(b), [16]float32(c)))
}

func fmaddPs(a [16]float32, b [16]float32, c [16]float32) [16]float32


// MaskFmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask_fmadd_ps'.
// Requires KNCNI.
func MaskFmaddPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(maskFmaddPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c)))
}

func maskFmaddPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32


// Mask3FmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask3_fmadd_ps'.
// Requires KNCNI.
func Mask3FmaddPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) x86.M512 {
	return x86.M512(mask3FmaddPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k)))
}

func mask3FmaddPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32


// FmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
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
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_fmadd_round_pd'.
// Requires KNCNI.
func FmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(fmaddRoundPd([8]float64(a), [8]float64(b), [8]float64(c), rounding))
}

func fmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64


// MaskFmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'a' when the corresponding mask bit is not set).
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
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask_fmadd_round_pd'.
// Requires KNCNI.
func MaskFmaddRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskFmaddRoundPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c), rounding))
}

func maskFmaddRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64


// Mask3FmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'c' when the corresponding mask bit is not set). 
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
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE 
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask3_fmadd_round_pd'.
// Requires KNCNI.
func Mask3FmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) x86.M512d {
	return x86.M512d(mask3FmaddRoundPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k), rounding))
}

func mask3FmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64


// FmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_fmadd_round_ps'.
// Requires KNCNI.
func FmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(fmaddRoundPs([16]float32(a), [16]float32(b), [16]float32(c), rounding))
}

func fmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32


// MaskFmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'a' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask_fmadd_round_ps'.
// Requires KNCNI.
func MaskFmaddRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(maskFmaddRoundPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c), rounding))
}

func maskFmaddRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32


// Mask3FmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'c' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask3_fmadd_round_ps'.
// Requires KNCNI.
func Mask3FmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) x86.M512 {
	return x86.M512(mask3FmaddRoundPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k), rounding))
}

func mask3FmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32


// Fmadd233Epi32: Multiply packed 32-bit integer elements in each 4-element set
// of 'a' and by element 1 of the corresponding 4-element set from 'b', add the
// intermediate result to element 0 of the corresponding 4-element set from
// 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			base := (j & ~0x3) * 32
//			scale[31:0] := b[base+63:base+32]
//			bias[31:0]  := b[base+31:base]
//			dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD233D'. Intrinsic: '_mm512_fmadd233_epi32'.
// Requires KNCNI.
func Fmadd233Epi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(fmadd233Epi32([64]byte(a), [64]byte(b)))
}

func fmadd233Epi32(a [64]byte, b [64]byte) [64]byte


// MaskFmadd233Epi32: Multiply packed 32-bit integer elements in each 4-element
// set of 'a' and by element 1 of the corresponding 4-element set from 'b', add
// the intermediate result to element 0 of the corresponding 4-element set from
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				base := (j & ~0x3) * 32
//				scale[31:0] := b[base+63:base+32]
//				bias[31:0]  := b[base+31:base]
//				dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD233D'. Intrinsic: '_mm512_mask_fmadd233_epi32'.
// Requires KNCNI.
func MaskFmadd233Epi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskFmadd233Epi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskFmadd233Epi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// Fmadd233Ps: Performs multiplication between single-precision (32-bit)
// floating-point elements in 'a' and 'b' and adds the result to the elements
// in 'b', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_fmadd233_ps'.
// Requires KNCNI.
func Fmadd233Ps(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(fmadd233Ps([16]float32(a), [16]float32(b)))
}

func fmadd233Ps(a [16]float32, b [16]float32) [16]float32


// MaskFmadd233Ps: Performs multiplication between single-precision (32-bit)
// floating-point elements in 'a' and 'b' and adds the result to the elements
// in 'b', storing the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_mask_fmadd233_ps'.
// Requires KNCNI.
func MaskFmadd233Ps(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskFmadd233Ps([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskFmadd233Ps(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// Fmadd233RoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in each 4-element set of 'a' and by element 1 of the corresponding
// 4-element set from 'b', add the intermediate result to element 0 of the
// corresponding 4-element set from 'b', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			base := (j & ~0x3) * 32
//			scale[31:0] := b[base+63:base+32]
//			bias[31:0]  := b[base+31:base]
//			dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_fmadd233_round_ps'.
// Requires KNCNI.
func Fmadd233RoundPs(a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(fmadd233RoundPs([16]float32(a), [16]float32(b), rounding))
}

func fmadd233RoundPs(a [16]float32, b [16]float32, rounding int) [16]float32


// MaskFmadd233RoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in each 4-element set of 'a' and by element 1 of the
// corresponding 4-element set from 'b', add the intermediate result to element
// 0 of the corresponding 4-element set from 'b', and store the results in
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
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				base := (j & ~0x3) * 32
//				scale[31:0] := b[base+63:base+32]
//				bias[31:0]  := b[base+31:base]
//				dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_mask_fmadd233_round_ps'.
// Requires KNCNI.
func MaskFmadd233RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(maskFmadd233RoundPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), rounding))
}

func maskFmadd233RoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32


// FmsubPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the intermediate
// result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_fmsub_pd'.
// Requires KNCNI.
func FmsubPd(a x86.M512d, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(fmsubPd([8]float64(a), [8]float64(b), [8]float64(c)))
}

func fmsubPd(a [8]float64, b [8]float64, c [8]float64) [8]float64


// MaskFmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask_fmsub_pd'.
// Requires KNCNI.
func MaskFmsubPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(maskFmsubPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c)))
}

func maskFmsubPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64


// Mask3FmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask3_fmsub_pd'.
// Requires KNCNI.
func Mask3FmsubPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) x86.M512d {
	return x86.M512d(mask3FmsubPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k)))
}

func mask3FmsubPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64


// FmsubPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the intermediate
// result, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_fmsub_ps'.
// Requires KNCNI.
func FmsubPs(a x86.M512, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(fmsubPs([16]float32(a), [16]float32(b), [16]float32(c)))
}

func fmsubPs(a [16]float32, b [16]float32, c [16]float32) [16]float32


// MaskFmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask_fmsub_ps'.
// Requires KNCNI.
func MaskFmsubPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(maskFmsubPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c)))
}

func maskFmsubPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32


// Mask3FmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask3_fmsub_ps'.
// Requires KNCNI.
func Mask3FmsubPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) x86.M512 {
	return x86.M512(mask3FmsubPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k)))
}

func mask3FmsubPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32


// FmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
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
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_fmsub_round_pd'.
// Requires KNCNI.
func FmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(fmsubRoundPd([8]float64(a), [8]float64(b), [8]float64(c), rounding))
}

func fmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64


// MaskFmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask_fmsub_round_pd'.
// Requires KNCNI.
func MaskFmsubRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskFmsubRoundPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c), rounding))
}

func maskFmsubRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64


// Mask3FmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask3_fmsub_round_pd'.
// Requires KNCNI.
func Mask3FmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) x86.M512d {
	return x86.M512d(mask3FmsubRoundPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k), rounding))
}

func mask3FmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64


// FmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_fmsub_round_ps'.
// Requires KNCNI.
func FmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(fmsubRoundPs([16]float32(a), [16]float32(b), [16]float32(c), rounding))
}

func fmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32


// MaskFmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask_fmsub_round_ps'.
// Requires KNCNI.
func MaskFmsubRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(maskFmsubRoundPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c), rounding))
}

func maskFmsubRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32


// Mask3FmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask3_fmsub_round_ps'.
// Requires KNCNI.
func Mask3FmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) x86.M512 {
	return x86.M512(mask3FmsubRoundPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k), rounding))
}

func mask3FmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32


// FnmaddPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', add the negated intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_fnmadd_pd'.
// Requires KNCNI.
func FnmaddPd(a x86.M512d, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(fnmaddPd([8]float64(a), [8]float64(b), [8]float64(c)))
}

func fnmaddPd(a [8]float64, b [8]float64, c [8]float64) [8]float64


// MaskFnmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask_fnmadd_pd'.
// Requires KNCNI.
func MaskFnmaddPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(maskFnmaddPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c)))
}

func maskFnmaddPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64


// Mask3FnmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask3_fnmadd_pd'.
// Requires KNCNI.
func Mask3FnmaddPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) x86.M512d {
	return x86.M512d(mask3FnmaddPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k)))
}

func mask3FnmaddPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64


// FnmaddPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', add the negated intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			a[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_fnmadd_ps'.
// Requires KNCNI.
func FnmaddPs(a x86.M512, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(fnmaddPs([16]float32(a), [16]float32(b), [16]float32(c)))
}

func fnmaddPs(a [16]float32, b [16]float32, c [16]float32) [16]float32


// MaskFnmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask_fnmadd_ps'.
// Requires KNCNI.
func MaskFnmaddPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(maskFnmaddPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c)))
}

func maskFnmaddPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32


// Mask3FnmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask3_fnmadd_ps'.
// Requires KNCNI.
func Mask3FnmaddPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) x86.M512 {
	return x86.M512(mask3FnmaddPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k)))
}

func mask3FnmaddPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32


// FnmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'.
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
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_fnmadd_round_pd'.
// Requires KNCNI.
func FnmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(fnmaddRoundPd([8]float64(a), [8]float64(b), [8]float64(c), rounding))
}

func fnmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64


// MaskFnmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask_fnmadd_round_pd'.
// Requires KNCNI.
func MaskFnmaddRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskFnmaddRoundPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c), rounding))
}

func maskFnmaddRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64


// Mask3FnmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask3_fnmadd_round_pd'.
// Requires KNCNI.
func Mask3FnmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) x86.M512d {
	return x86.M512d(mask3FnmaddRoundPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k), rounding))
}

func mask3FnmaddRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64


// FnmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'.  
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_fnmadd_round_ps'.
// Requires KNCNI.
func FnmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(fnmaddRoundPs([16]float32(a), [16]float32(b), [16]float32(c), rounding))
}

func fnmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32


// MaskFnmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask_fnmadd_round_ps'.
// Requires KNCNI.
func MaskFnmaddRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(maskFnmaddRoundPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c), rounding))
}

func maskFnmaddRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32


// Mask3FnmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask3_fnmadd_round_ps'.
// Requires KNCNI.
func Mask3FnmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) x86.M512 {
	return x86.M512(mask3FnmaddRoundPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k), rounding))
}

func mask3FnmaddRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32


// FnmsubPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_fnmsub_pd'.
// Requires KNCNI.
func FnmsubPd(a x86.M512d, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(fnmsubPd([8]float64(a), [8]float64(b), [8]float64(c)))
}

func fnmsubPd(a [8]float64, b [8]float64, c [8]float64) [8]float64


// MaskFnmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask_fnmsub_pd'.
// Requires KNCNI.
func MaskFnmsubPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) x86.M512d {
	return x86.M512d(maskFnmsubPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c)))
}

func maskFnmsubPd(a [8]float64, k uint8, b [8]float64, c [8]float64) [8]float64


// Mask3FnmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask3_fnmsub_pd'.
// Requires KNCNI.
func Mask3FnmsubPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) x86.M512d {
	return x86.M512d(mask3FnmsubPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k)))
}

func mask3FnmsubPd(a [8]float64, b [8]float64, c [8]float64, k uint8) [8]float64


// FnmsubPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_fnmsub_ps'.
// Requires KNCNI.
func FnmsubPs(a x86.M512, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(fnmsubPs([16]float32(a), [16]float32(b), [16]float32(c)))
}

func fnmsubPs(a [16]float32, b [16]float32, c [16]float32) [16]float32


// MaskFnmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask_fnmsub_ps'.
// Requires KNCNI.
func MaskFnmsubPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) x86.M512 {
	return x86.M512(maskFnmsubPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c)))
}

func maskFnmsubPs(a [16]float32, k uint16, b [16]float32, c [16]float32) [16]float32


// Mask3FnmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask3_fnmsub_ps'.
// Requires KNCNI.
func Mask3FnmsubPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) x86.M512 {
	return x86.M512(mask3FnmsubPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k)))
}

func mask3FnmsubPs(a [16]float32, b [16]float32, c [16]float32, k uint16) [16]float32


// FnmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'.  
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
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_fnmsub_round_pd'.
// Requires KNCNI.
func FnmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(fnmsubRoundPd([8]float64(a), [8]float64(b), [8]float64(c), rounding))
}

func fnmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, rounding int) [8]float64


// MaskFnmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
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
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask_fnmsub_round_pd'.
// Requires KNCNI.
func MaskFnmsubRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskFnmsubRoundPd([8]float64(a), uint8(k), [8]float64(b), [8]float64(c), rounding))
}

func maskFnmsubRoundPd(a [8]float64, k uint8, b [8]float64, c [8]float64, rounding int) [8]float64


// Mask3FnmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask3_fnmsub_round_pd'.
// Requires KNCNI.
func Mask3FnmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) x86.M512d {
	return x86.M512d(mask3FnmsubRoundPd([8]float64(a), [8]float64(b), [8]float64(c), uint8(k), rounding))
}

func mask3FnmsubRoundPd(a [8]float64, b [8]float64, c [8]float64, k uint8, rounding int) [8]float64


// FnmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_fnmsub_round_ps'.
// Requires KNCNI.
func FnmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(fnmsubRoundPs([16]float32(a), [16]float32(b), [16]float32(c), rounding))
}

func fnmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, rounding int) [16]float32


// MaskFnmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask_fnmsub_round_ps'.
// Requires KNCNI.
func MaskFnmsubRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) x86.M512 {
	return x86.M512(maskFnmsubRoundPs([16]float32(a), uint16(k), [16]float32(b), [16]float32(c), rounding))
}

func maskFnmsubRoundPs(a [16]float32, k uint16, b [16]float32, c [16]float32, rounding int) [16]float32


// Mask3FnmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask3_fnmsub_round_ps'.
// Requires KNCNI.
func Mask3FnmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) x86.M512 {
	return x86.M512(mask3FnmsubRoundPs([16]float32(a), [16]float32(b), [16]float32(c), uint16(k), rounding))
}

func mask3FnmsubRoundPs(a [16]float32, b [16]float32, c [16]float32, k uint16, rounding int) [16]float32


// GetexpPd: Convert the exponent of each packed double-precision (64-bit)
// floating-point element in 'a' to a double-precision (64-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'.
// This intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_getexp_pd'.
// Requires KNCNI.
func GetexpPd(a x86.M512d) x86.M512d {
	return x86.M512d(getexpPd([8]float64(a)))
}

func getexpPd(a [8]float64) [8]float64


// MaskGetexpPd: Convert the exponent of each packed double-precision (64-bit)
// floating-point element in 'a' to a double-precision (64-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates 'floor(log2(x))'
// for each element. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_mask_getexp_pd'.
// Requires KNCNI.
func MaskGetexpPd(src x86.M512d, k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(maskGetexpPd([8]float64(src), uint8(k), [8]float64(a)))
}

func maskGetexpPd(src [8]float64, k uint8, a [8]float64) [8]float64


// GetexpPs: Convert the exponent of each packed single-precision (32-bit)
// floating-point element in 'a' to a single-precision (32-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'.
// This intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_getexp_ps'.
// Requires KNCNI.
func GetexpPs(a x86.M512) x86.M512 {
	return x86.M512(getexpPs([16]float32(a)))
}

func getexpPs(a [16]float32) [16]float32


// MaskGetexpPs: Convert the exponent of each packed single-precision (32-bit)
// floating-point element in 'a' to a single-precision (32-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates 'floor(log2(x))'
// for each element. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_mask_getexp_ps'.
// Requires KNCNI.
func MaskGetexpPs(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(maskGetexpPs([16]float32(src), uint16(k), [16]float32(a)))
}

func maskGetexpPs(src [16]float32, k uint16, a [16]float32) [16]float32


// GetexpRoundPd: Convert the exponent of each packed double-precision (64-bit)
// floating-point element in 'a' to a double-precision (64-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'.
// This intrinsic essentially calculates 'floor(log2(x))' for each element.
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
//			dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_getexp_round_pd'.
// Requires KNCNI.
func GetexpRoundPd(a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(getexpRoundPd([8]float64(a), rounding))
}

func getexpRoundPd(a [8]float64, rounding int) [8]float64


// MaskGetexpRoundPd: Convert the exponent of each packed double-precision
// (64-bit) floating-point element in 'a' to a double-precision (64-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). This intrinsic essentially
// calculates 'floor(log2(x))' for each element.
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
//				dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_mask_getexp_round_pd'.
// Requires KNCNI.
func MaskGetexpRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskGetexpRoundPd([8]float64(src), uint8(k), [8]float64(a), rounding))
}

func maskGetexpRoundPd(src [8]float64, k uint8, a [8]float64, rounding int) [8]float64


// GetexpRoundPs: Convert the exponent of each packed single-precision (32-bit)
// floating-point element in 'a' to a single-precision (32-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'.
// This intrinsic essentially calculates 'floor(log2(x))' for each element.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_getexp_round_ps'.
// Requires KNCNI.
func GetexpRoundPs(a x86.M512, rounding int) x86.M512 {
	return x86.M512(getexpRoundPs([16]float32(a), rounding))
}

func getexpRoundPs(a [16]float32, rounding int) [16]float32


// MaskGetexpRoundPs: Convert the exponent of each packed single-precision
// (32-bit) floating-point element in 'a' to a single-precision (32-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). This intrinsic essentially
// calculates 'floor(log2(x))' for each element.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_mask_getexp_round_ps'.
// Requires KNCNI.
func MaskGetexpRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) x86.M512 {
	return x86.M512(maskGetexpRoundPs([16]float32(src), uint16(k), [16]float32(a), rounding))
}

func maskGetexpRoundPs(src [16]float32, k uint16, a [16]float32, rounding int) [16]float32


// GetmantPd: Normalize the mantissas of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. This
// intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_getmant_pd'.
// Requires KNCNI.
func GetmantPd(a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) x86.M512d {
	return x86.M512d(getmantPd([8]float64(a), interv, sc))
}

func getmantPd(a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float64


// MaskGetmantPd: Normalize the mantissas of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_mask_getmant_pd'.
// Requires KNCNI.
func MaskGetmantPd(src x86.M512d, k x86.Mmask8, a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) x86.M512d {
	return x86.M512d(maskGetmantPd([8]float64(src), uint8(k), [8]float64(a), interv, sc))
}

func maskGetmantPd(src [8]float64, k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [8]float64


// GetmantPs: Normalize the mantissas of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. This
// intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_getmant_ps'.
// Requires KNCNI.
func GetmantPs(a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) x86.M512 {
	return x86.M512(getmantPs([16]float32(a), interv, sc))
}

func getmantPs(a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [16]float32


// MaskGetmantPs: Normalize the mantissas of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_mask_getmant_ps'.
// Requires KNCNI.
func MaskGetmantPs(src x86.M512, k x86.Mmask16, a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) x86.M512 {
	return x86.M512(maskGetmantPs([16]float32(src), uint16(k), [16]float32(a), interv, sc))
}

func maskGetmantPs(src [16]float32, k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) [16]float32


// GetmantRoundPd: Normalize the mantissas of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. This
// intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_getmant_round_pd'.
// Requires KNCNI.
func GetmantRoundPd(a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) x86.M512d {
	return x86.M512d(getmantRoundPd([8]float64(a), interv, sc, rounding))
}

func getmantRoundPd(a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [8]float64


// MaskGetmantRoundPd: Normalize the mantissas of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_mask_getmant_round_pd'.
// Requires KNCNI.
func MaskGetmantRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) x86.M512d {
	return x86.M512d(maskGetmantRoundPd([8]float64(src), uint8(k), [8]float64(a), interv, sc, rounding))
}

func maskGetmantRoundPd(src [8]float64, k uint8, a [8]float64, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [8]float64


// GetmantRoundPs: Normalize the mantissas of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. This
// intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_getmant_round_ps'.
// Requires KNCNI.
func GetmantRoundPs(a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) x86.M512 {
	return x86.M512(getmantRoundPs([16]float32(a), interv, sc, rounding))
}

func getmantRoundPs(a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [16]float32


// MaskGetmantRoundPs: Normalize the mantissas of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_mask_getmant_round_ps'.
// Requires KNCNI.
func MaskGetmantRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) x86.M512 {
	return x86.M512(maskGetmantRoundPs([16]float32(src), uint16(k), [16]float32(a), interv, sc, rounding))
}

func maskGetmantRoundPs(src [16]float32, k uint16, a [16]float32, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) [16]float32


// GmaxPd: Determines the maximum of each pair of corresponding elements in
// packed double-precision (64-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := FpMax(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPD'. Intrinsic: '_mm512_gmax_pd'.
// Requires KNCNI.
func GmaxPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(gmaxPd([8]float64(a), [8]float64(b)))
}

func gmaxPd(a [8]float64, b [8]float64) [8]float64


// MaskGmaxPd: Determines the maximum of each pair of corresponding elements of
// packed double-precision (64-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := FpMax(a[i+63:i], b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPD'. Intrinsic: '_mm512_mask_gmax_pd'.
// Requires KNCNI.
func MaskGmaxPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(maskGmaxPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskGmaxPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// GmaxPs: Determines the maximum of each pair of corresponding elements in
// packed single-precision (32-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMax(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPS'. Intrinsic: '_mm512_gmax_ps'.
// Requires KNCNI.
func GmaxPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(gmaxPs([16]float32(a), [16]float32(b)))
}

func gmaxPs(a [16]float32, b [16]float32) [16]float32


// MaskGmaxPs: Determines the maximum of each pair of corresponding elements of
// packed single-precision (32-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMax(a[i+31:i], b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPS'. Intrinsic: '_mm512_mask_gmax_ps'.
// Requires KNCNI.
func MaskGmaxPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskGmaxPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskGmaxPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// GmaxabsPs: Determines the maximum of the absolute elements of each pair of
// corresponding elements of packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_gmaxabs_ps'.
// Requires KNCNI.
func GmaxabsPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(gmaxabsPs([16]float32(a), [16]float32(b)))
}

func gmaxabsPs(a [16]float32, b [16]float32) [16]float32


// MaskGmaxabsPs: Determines the maximum of the absolute elements of each pair
// of corresponding elements of packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_mask_gmaxabs_ps'.
// Requires KNCNI.
func MaskGmaxabsPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskGmaxabsPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskGmaxabsPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// GminPd: Determines the minimum of each pair of corresponding elements in
// packed double-precision (64-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := FpMin(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPD'. Intrinsic: '_mm512_gmin_pd'.
// Requires KNCNI.
func GminPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(gminPd([8]float64(a), [8]float64(b)))
}

func gminPd(a [8]float64, b [8]float64) [8]float64


// MaskGminPd: Determines the maximum of each pair of corresponding elements of
// packed double-precision (64-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := FpMin(a[i+63:i], b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPD'. Intrinsic: '_mm512_mask_gmin_pd'.
// Requires KNCNI.
func MaskGminPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(maskGminPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskGminPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// GminPs: Determines the minimum of each pair of corresponding elements in
// packed single-precision (32-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMin(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPS'. Intrinsic: '_mm512_gmin_ps'.
// Requires KNCNI.
func GminPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(gminPs([16]float32(a), [16]float32(b)))
}

func gminPs(a [16]float32, b [16]float32) [16]float32


// MaskGminPs: Determines the maximum of each pair of corresponding elements of
// packed single-precision (32-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMin(a[i+31:i], b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPS'. Intrinsic: '_mm512_mask_gmin_ps'.
// Requires KNCNI.
func MaskGminPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskGminPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskGminPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// I32extgatherEpi32: Up-converts 16 memory locations starting at location 'mv'
// at packed 32-bit integer indices stored in 'index' scaled by 'scale' using
// 'conv' to 32-bit integer elements and stores them in 'dst'. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:
//				dst[i+31:i] := addr[i+31:i]
//			_MM_UPCONV_EPI32_UINT8:
//				n := j*7
//				dst[i+31:i] := UInt8ToUInt32(addr[n+7:n])
//			_MM_UPCONV_EPI32_SINT8:
//				n := j*7
//				dst[i+31:i] := Int8ToInt32(addr[n+7:n])
//			_MM_UPCONV_EPI32_UINT16:
//				n := j*16
//				dst[i+31:i] := UInt16ToUInt32(addr[n+15:n])
//			_MM_UPCONV_EPI32_SINT16:
//				n := j*16
//				dst[i+31:i] := Int16ToInt32(addr[n+15:n])
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm512_i32extgather_epi32'.
// Requires KNCNI.
func I32extgatherEpi32(index x86.M512i, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(i32extgatherEpi32([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i32extgatherEpi32(index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte


// MaskI32extgatherEpi32: Up-converts 16 single-precision (32-bit) memory
// locations starting at location 'mv' at packed 32-bit integer indices stored
// in 'index' scaled by 'scale' using 'conv' to 32-bit integer elements and
// stores them in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_EPI32_NONE:
//					dst[i+31:i] := addr[i+31:i]
//				_MM_UPCONV_EPI32_UINT8:
//					n := j*7
//					dst[i+31:i] := UInt8ToUInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_SINT8:
//					n := j*7
//					dst[i+31:i] := Int8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_UINT16:
//					n := j*16
//					dst[i+31:i] := UInt16ToUInt32(addr[n+15:n])
//				_MM_UPCONV_EPI32_SINT16:
//					n := j*16
//					dst[i+31:i] := Int16ToInt32(addr[n+15:n])
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm512_mask_i32extgather_epi32'.
// Requires KNCNI.
func MaskI32extgatherEpi32(src x86.M512i, k x86.Mmask16, index x86.M512i, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(maskI32extgatherEpi32([64]byte(src), uint16(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI32extgatherEpi32(src [64]byte, k uint16, index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte


// I32extgatherPs: Up-converts 16 memory locations starting at location 'mv' at
// packed 32-bit integer indices stored in 'index' scaled by 'scale' using
// 'conv' to single-precision (32-bit) floating-point elements and stores them
// in 'dst'. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:
//				dst[i+31:i] := addr[i+31:i]
//			_MM_UPCONV_PS_FLOAT16:
//				n := j*16
//				dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//			_MM_UPCONV_PS_UINT8:
//				n := j*8
//				dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//			_MM_UPCONV_PS_SINT8:
//				n := j*8
//				dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//			_MM_UPCONV_PS_UINT16:
//				n := j*16
//				dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//			_MM_UPCONV_PS_SINT16:
//				n := j*16
//				dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPS'. Intrinsic: '_mm512_i32extgather_ps'.
// Requires KNCNI.
func I32extgatherPs(index x86.M512i, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) x86.M512 {
	return x86.M512(i32extgatherPs([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i32extgatherPs(index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32


// MaskI32extgatherPs: Up-converts 16 single-precision (32-bit) memory
// locations starting at location 'mv' at packed 32-bit integer indices stored
// in 'index' scaled by 'scale' using 'conv' to single-precision (32-bit)
// floating-point elements and stores them in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_PS_NONE:
//					dst[i+31:i] := addr[i+31:i]
//				_MM_UPCONV_PS_FLOAT16:
//					n := j*16
//					dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_UINT8:
//					n := j*8
//					dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_SINT8:
//					n := j*8
//					dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_UINT16:
//					n := j*16
//					dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_SINT16:
//					n := j*16
//					dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPS'. Intrinsic: '_mm512_mask_i32extgather_ps'.
// Requires KNCNI.
func MaskI32extgatherPs(src x86.M512, k x86.Mmask16, index x86.M512i, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) x86.M512 {
	return x86.M512(maskI32extgatherPs([16]float32(src), uint16(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI32extgatherPs(src [16]float32, k uint16, index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32


// I32extscatterEpi32: Down-converts 16 packed 32-bit integer elements in 'v1'
// using 'conv' and stores them in memory locations starting at location 'mv'
// at packed 32-bit integer indices stored in 'index' scaled by 'scale'. 'hint'
// indicates to the processor whether the data is non-temporal. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_DOWNCONV_EPI32_NONE:
//				addr[i+31:i] := v1[i+31:i]
//			_MM_DOWNCONV_EPI32_UINT8:
//				n := j*8
//				addr[n+7:n] := UInt32ToUInt8(v1[i+31:i])
//			_MM_DOWNCONV_EPI32_SINT8:
//				n := j*8
//				addr[n+7:n] := SInt32ToSInt8(v1[i+31:i])
//			_MM_DOWNCONV_EPI32_UINT16:
//				n := j*16
//				addr[n+15:n] := UInt32ToUInt16(v1[i+31:i])
//			_MM_DOWNCONV_EPI32_SINT16:
//				n := j*16
//				addr[n+15:n] := SInt32ToSInt16(v1[n+15:n])
//			ESAC
//		ENDFOR
//
// Instruction: 'VPSCATTERDD'. Intrinsic: '_mm512_i32extscatter_epi32'.
// Requires KNCNI.
func I32extscatterEpi32(mv uintptr, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, scale int, hint int)  {
	i32extscatterEpi32(uintptr(mv), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func i32extscatterEpi32(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 


// MaskI32extscatterEpi32: Down-converts 16 packed 32-bit integer elements in
// 'v1' using 'conv' and stores them in memory locations starting at location
// 'mv' at packed 32-bit integer indices stored in 'index' scaled by 'scale'.
// Elements are written using writemask 'k' (elements are only written when the
// corresponding mask bit is set; otherwise, elements are left unchanged in
// memory). 'hint' indicates to the processor whether the data is non-temporal. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_EPI32_NONE:
//					addr[i+31:i] := v1[i+31:i]
//				_MM_DOWNCONV_EPI32_UINT8:
//					n := j*8
//					addr[n+7:n] := UInt32ToUInt8(v1[i+31:i])
//				_MM_DOWNCONV_EPI32_SINT8:
//					n := j*8
//					addr[n+7:n] := SInt32ToSInt8(v1[i+31:i])
//				_MM_DOWNCONV_EPI32_UINT16:
//					n := j*16
//					addr[n+15:n] := UInt32ToUInt16(v1[i+31:i])
//				_MM_DOWNCONV_EPI32_SINT16:
//					n := j*16
//					addr[n+15:n] := SInt32ToSInt16(v1[n+15:n])
//				ESAC
//			FI 
//		ENDFOR
//
// Instruction: 'VPSCATTERDD'. Intrinsic: '_mm512_mask_i32extscatter_epi32'.
// Requires KNCNI.
func MaskI32extscatterEpi32(mv uintptr, k x86.Mmask16, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, scale int, hint int)  {
	maskI32extscatterEpi32(uintptr(mv), uint16(k), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func maskI32extscatterEpi32(mv uintptr, k uint16, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 


// I32extscatterPs: Down-converts 16 packed single-precision (32-bit)
// floating-point elements in 'v1' and stores them in memory locations starting
// at location 'mv' at packed 32-bit integer indices stored in 'index' scaled
// by 'scale' using 'conv'. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_DOWNCONV_PS_NONE:
//				n := j*32
//				addr[i+31:i] := v1[n+31:n]
//			_MM_DOWNCONV_PS_FLOAT16:
//				i := j*16
//				addr[i+15:i] := Float32ToFloat16(v1[n+31:n])
//			_MM_DOWNCONV_PS_UINT8:
//				i := j*8
//				addr[i+7:i] := Float32ToUInt8(v1[n+31:n])
//			_MM_DOWNCONV_PS_SINT8:
//				i := j*8
//				addr[i+7:i] := Float32ToSInt8(v1[n+31:n])
//			_MM_DOWNCONV_PS_UINT16:
//				i := j*8
//				addr[i+15:i] := Float32ToUInt16(v1[n+31:n])
//			_MM_DOWNCONV_PS_SINT16:
//				i := j*8
//				addr[i+15:i] := Float32ToSInt16(v1[n+31:n])
//			ESAC
//		ENDFOR
//
// Instruction: 'VSCATTERDPS'. Intrinsic: '_mm512_i32extscatter_ps'.
// Requires KNCNI.
func I32extscatterPs(mv uintptr, index x86.M512i, v1 x86.M512, conv MMDOWNCONVPSENUM, scale int)  {
	i32extscatterPs(uintptr(mv), [64]byte(index), [16]float32(v1), conv, scale)
}

func i32extscatterPs(mv uintptr, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int) 


// MaskI32extscatterPs: Down-converts 16 packed single-precision (32-bit)
// floating-point elements in 'v1' according to 'conv' and stores them in
// memory locations starting at location 'mv' at packed 32-bit integer indices
// stored in 'index' scaled by 'scale' using writemask 'k' (elements are
// written only when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				CASE conv OF
//				_MM_DOWNCONV_PS_NONE:
//					n := j*32
//					addr[i+31:i] := v1[n+31:n]
//				_MM_DOWNCONV_PS_FLOAT16:
//					i := j*16
//					addr[i+15:i] := Float32ToFloat16(v1[n+31:n])
//				_MM_DOWNCONV_PS_UINT8:
//					i := j*8
//					addr[i+7:i] := Float32ToUInt8(v1[n+31:n])
//				_MM_DOWNCONV_PS_SINT8:
//					i := j*8
//					addr[i+7:i] := Float32ToSInt8(v1[n+31:n])
//				_MM_DOWNCONV_PS_UINT16:
//					i := j*8
//					addr[i+15:i] := Float32ToUInt16(v1[n+31:n])
//				_MM_DOWNCONV_PS_SINT16:
//					i := j*8
//					addr[i+15:i] := Float32ToSInt16(v1[n+31:n])
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: 'VSCATTERDPS'. Intrinsic: '_mm512_mask_i32extscatter_ps'.
// Requires KNCNI.
func MaskI32extscatterPs(mv uintptr, k x86.Mmask16, index x86.M512i, v1 x86.M512, conv MMDOWNCONVPSENUM, scale int, hint int)  {
	maskI32extscatterPs(uintptr(mv), uint16(k), [64]byte(index), [16]float32(v1), conv, scale, hint)
}

func maskI32extscatterPs(mv uintptr, k uint16, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int, hint int) 


// I32gatherEpi32: Gather 32-bit integers from memory using 32-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst'. 'scale' should be 1, 2, 4
// or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm512_i32gather_epi32'.
// Requires KNCNI.
func I32gatherEpi32(vindex x86.M512i, base_addr uintptr, scale int) x86.M512i {
	return x86.M512i(i32gatherEpi32([64]byte(vindex), uintptr(base_addr), scale))
}

func i32gatherEpi32(vindex [64]byte, base_addr uintptr, scale int) [64]byte


// MaskI32gatherEpi32: Gather 32-bit integers from memory using 32-bit indices.
// 32-bit elements are loaded from addresses starting at 'base_addr' and offset
// by each 32-bit element in 'vindex' (each index is scaled by the factor in
// 'scale'). Gathered elements are merged into 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//				k[j] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDD'. Intrinsic: '_mm512_mask_i32gather_epi32'.
// Requires KNCNI.
func MaskI32gatherEpi32(src x86.M512i, k x86.Mmask16, vindex x86.M512i, base_addr uintptr, scale int) x86.M512i {
	return x86.M512i(maskI32gatherEpi32([64]byte(src), uint16(k), [64]byte(vindex), uintptr(base_addr), scale))
}

func maskI32gatherEpi32(src [64]byte, k uint16, vindex [64]byte, base_addr uintptr, scale int) [64]byte


// I32gatherPs: Gather single-precision (32-bit) floating-point elements from
// memory using 32-bit indices. 32-bit elements are loaded from addresses
// starting at 'base_addr' and offset by each 32-bit element in 'vindex' (each
// index is scaled by the factor in 'scale'). Gathered elements are merged into
// 'dst'. 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPS'. Intrinsic: '_mm512_i32gather_ps'.
// Requires KNCNI.
func I32gatherPs(vindex x86.M512i, base_addr uintptr, scale int) x86.M512 {
	return x86.M512(i32gatherPs([64]byte(vindex), uintptr(base_addr), scale))
}

func i32gatherPs(vindex [64]byte, base_addr uintptr, scale int) [16]float32


// MaskI32gatherPs: Gather single-precision (32-bit) floating-point elements
// from memory using 32-bit indices. 32-bit elements are loaded from addresses
// starting at 'base_addr' and offset by each 32-bit element in 'vindex' (each
// index is scaled by the factor in 'scale'). Gathered elements are merged into
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := MEM[base_addr + SignExtend(vindex[i+31:i])*scale]
//				k[j] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPS'. Intrinsic: '_mm512_mask_i32gather_ps'.
// Requires KNCNI.
func MaskI32gatherPs(src x86.M512, k x86.Mmask16, vindex x86.M512i, base_addr uintptr, scale int) x86.M512 {
	return x86.M512(maskI32gatherPs([16]float32(src), uint16(k), [64]byte(vindex), uintptr(base_addr), scale))
}

func maskI32gatherPs(src [16]float32, k uint16, vindex [64]byte, base_addr uintptr, scale int) [16]float32


// I32loextgatherEpi64: Up-converts 8 double-precision (64-bit) memory
// locations starting at location 'mv' at packed 32-bit integer indices stored
// in the lower half of 'index' scaled by 'scale' using 'conv' to 64-bit
// integer elements and stores them in 'dst'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE: dst[i+63:i] := addr[i+63:i]
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm512_i32loextgather_epi64'.
// Requires KNCNI.
func I32loextgatherEpi64(index x86.M512i, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(i32loextgatherEpi64([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i32loextgatherEpi64(index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte


// MaskI32loextgatherEpi64: Up-converts 8 double-precision (64-bit) memory
// locations starting at location 'mv' at packed 32-bit integer indices stored
// in the lower half of 'index' scaled by 'scale' using 'conv' to 64-bit
// integer elements and stores them in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_EPI64_NONE: dst[i+63:i] := addr[i+63:i]
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm512_mask_i32loextgather_epi64'.
// Requires KNCNI.
func MaskI32loextgatherEpi64(src x86.M512i, k x86.Mmask8, index x86.M512i, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(maskI32loextgatherEpi64([64]byte(src), uint8(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI32loextgatherEpi64(src [64]byte, k uint8, index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte


// I32loextgatherPd: Up-converts 8 double-precision (64-bit) floating-point
// elements in memory locations starting at location 'mv' at packed 32-bit
// integer indices stored in the lower half of 'index' scaled by 'scale' using
// 'conv' to 64-bit floating-point elements and stores them in 'dst'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_UPCONV_PD_NONE: dst[i+63:i] := addr[i+63:i]
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPD'. Intrinsic: '_mm512_i32loextgather_pd'.
// Requires KNCNI.
func I32loextgatherPd(index x86.M512i, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) x86.M512d {
	return x86.M512d(i32loextgatherPd([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i32loextgatherPd(index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64


// MaskI32loextgatherPd: Up-converts 8 double-precision (64-bit) floating-point
// elements in memory locations starting at location 'mv' at packed 32-bit
// integer indices stored in the lower half of 'index' scaled by 'scale' using
// 'conv' to 64-bit floating-point elements and stores them in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_PD_NONE: dst[i+63:i] := addr[i+63:i]
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPD'. Intrinsic: '_mm512_mask_i32loextgather_pd'.
// Requires KNCNI.
func MaskI32loextgatherPd(src x86.M512d, k x86.Mmask8, index x86.M512i, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) x86.M512d {
	return x86.M512d(maskI32loextgatherPd([8]float64(src), uint8(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI32loextgatherPd(src [8]float64, k uint8, index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64


// I32loextscatterEpi64: Down-converts 8 packed 64-bit integer elements in 'v1'
// and stores them in memory locations starting at location 'mv' at packed
// 32-bit integer indices stored in 'index' scaled by 'scale' using 'conv'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_EPI64_NONE: addr[i+63:i] := v1[i+63:i]
//			ESAC
//		ENDFOR
//
// Instruction: 'VPSCATTERDQ'. Intrinsic: '_mm512_i32loextscatter_epi64'.
// Requires KNCNI.
func I32loextscatterEpi64(mv uintptr, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, scale int, hint int)  {
	i32loextscatterEpi64(uintptr(mv), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func i32loextscatterEpi64(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 


// MaskI32loextscatterEpi64: Down-converts 8 packed 64-bit integer elements in
// 'v1' and stores them in memory locations starting at location 'mv' at packed
// 32-bit integer indices stored in 'index' scaled by 'scale' using 'conv'.
// Only those elements whose corresponding mask bit is set in writemask 'k' are
// written to memory. 
//
//		FOR j := 0 to 7
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				i := j*64
//				CASE conv OF
//				_MM_DOWNCONV_EPI64_NONE: addr[i+63:i] := v1[i+63:i]
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: 'VPSCATTERDQ'. Intrinsic: '_mm512_mask_i32loextscatter_epi64'.
// Requires KNCNI.
func MaskI32loextscatterEpi64(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, scale int, hint int)  {
	maskI32loextscatterEpi64(uintptr(mv), uint8(k), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func maskI32loextscatterEpi64(mv uintptr, k uint8, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 


// I32loextscatterPd: Down-converts 8 packed double-precision (64-bit)
// floating-point elements in 'v1' and stores them in memory locations starting
// at location 'mv' at packed 32-bit integer indices stored in 'index' scaled
// by 'scale' using 'conv'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_PD_NONE: addr[i+63:i] := v1[i+63:i]
//			ESAC
//		ENDFOR
//
// Instruction: 'VSCATTERDPD'. Intrinsic: '_mm512_i32loextscatter_pd'.
// Requires KNCNI.
func I32loextscatterPd(mv uintptr, index x86.M512i, v1 x86.M512d, conv MMDOWNCONVPDENUM, scale int, hint int)  {
	i32loextscatterPd(uintptr(mv), [64]byte(index), [8]float64(v1), conv, scale, hint)
}

func i32loextscatterPd(mv uintptr, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 


// MaskI32loextscatterPd: Down-converts 8 packed double-precision (64-bit)
// floating-point elements in 'v1' and stores them in memory locations starting
// at location 'mv' at packed 32-bit integer indices stored in 'index' scaled
// by 'scale' using 'conv'. Only those elements whose corresponding mask bit is
// set in writemask 'k' are written to memory. 
//
//		FOR j := 0 to 7
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				i := j*64
//				CASE conv OF
//				_MM_DOWNCONV_PD_NONE: addr[i+63:i] := v1[i+63:i]
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: 'VSCATTERDPD'. Intrinsic: '_mm512_mask_i32loextscatter_pd'.
// Requires KNCNI.
func MaskI32loextscatterPd(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512d, conv MMDOWNCONVPDENUM, scale int, hint int)  {
	maskI32loextscatterPd(uintptr(mv), uint8(k), [64]byte(index), [8]float64(v1), conv, scale, hint)
}

func maskI32loextscatterPd(mv uintptr, k uint8, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 


// I32logatherEpi64: Loads 8 64-bit integer elements from memory starting at
// location 'mv' at packed 32-bit integer indices stored in the lower half of
// 'index' scaled by 'scale' and stores them in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			addr := MEM[mv + index[j] * scale]
//			dst[i+63:i] := addr[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm512_i32logather_epi64'.
// Requires KNCNI.
func I32logatherEpi64(index x86.M512i, mv uintptr, scale int) x86.M512i {
	return x86.M512i(i32logatherEpi64([64]byte(index), uintptr(mv), scale))
}

func i32logatherEpi64(index [64]byte, mv uintptr, scale int) [64]byte


// MaskI32logatherEpi64: Loads 8 64-bit integer elements from memory starting
// at location 'mv' at packed 32-bit integer indices stored in the lower half
// of 'index' scaled by 'scale' and stores them in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				dst[i+63:i] := addr[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPGATHERDQ'. Intrinsic: '_mm512_mask_i32logather_epi64'.
// Requires KNCNI.
func MaskI32logatherEpi64(src x86.M512i, k x86.Mmask8, index x86.M512i, mv uintptr, scale int) x86.M512i {
	return x86.M512i(maskI32logatherEpi64([64]byte(src), uint8(k), [64]byte(index), uintptr(mv), scale))
}

func maskI32logatherEpi64(src [64]byte, k uint8, index [64]byte, mv uintptr, scale int) [64]byte


// I32logatherPd: Loads 8 double-precision (64-bit) floating-point elements
// stored at memory locations starting at location 'mv' at packed 32-bit
// integer indices stored in the lower half of 'index' scaled by 'scale' them
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			addr := MEM[mv + index[j] * scale]
//			dst[i+63:i] := addr[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPD'. Intrinsic: '_mm512_i32logather_pd'.
// Requires KNCNI.
func I32logatherPd(index x86.M512i, mv uintptr, scale int) x86.M512d {
	return x86.M512d(i32logatherPd([64]byte(index), uintptr(mv), scale))
}

func i32logatherPd(index [64]byte, mv uintptr, scale int) [8]float64


// MaskI32logatherPd: Loads 8 double-precision (64-bit) floating-point elements
// from memory starting at location 'mv' at packed 32-bit integer indices
// stored in the lower half of 'index' scaled by 'scale' into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				dst[i+63:i] := addr[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERDPD'. Intrinsic: '_mm512_mask_i32logather_pd'.
// Requires KNCNI.
func MaskI32logatherPd(src x86.M512d, k x86.Mmask8, index x86.M512i, mv uintptr, scale int) x86.M512d {
	return x86.M512d(maskI32logatherPd([8]float64(src), uint8(k), [64]byte(index), uintptr(mv), scale))
}

func maskI32logatherPd(src [8]float64, k uint8, index [64]byte, mv uintptr, scale int) [8]float64


// I32loscatterEpi64: Stores 8 packed 64-bit integer elements located in 'v1'
// and stores them in memory locations starting at location 'mv' at packed
// 32-bit integer indices stored in 'index' scaled by 'scale'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			addr[i+63:i] := v1[k+63:j]
//		ENDFOR
//
// Instruction: 'VPSCATTERDQ'. Intrinsic: '_mm512_i32loscatter_epi64'.
// Requires KNCNI.
func I32loscatterEpi64(mv uintptr, index x86.M512i, v1 x86.M512i, scale int)  {
	i32loscatterEpi64(uintptr(mv), [64]byte(index), [64]byte(v1), scale)
}

func i32loscatterEpi64(mv uintptr, index [64]byte, v1 [64]byte, scale int) 


// MaskI32loscatterEpi64: Stores 8 packed 64-bit integer elements located in
// 'v1' and stores them in memory locations starting at location 'mv' at packed
// 32-bit integer indices stored in 'index' scaled by 'scale' using writemask
// 'k' (elements whose corresponding mask bit is not set are not written to
// memory). 
//
//		FOR j := 0 to 7
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				addr[i+63:i] := v1[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VPSCATTERDQ'. Intrinsic: '_mm512_mask_i32loscatter_epi64'.
// Requires KNCNI.
func MaskI32loscatterEpi64(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512i, scale int)  {
	maskI32loscatterEpi64(uintptr(mv), uint8(k), [64]byte(index), [64]byte(v1), scale)
}

func maskI32loscatterEpi64(mv uintptr, k uint8, index [64]byte, v1 [64]byte, scale int) 


// I32loscatterPd: Stores 8 packed double-precision (64-bit) floating-point
// elements in 'v1' and to memory locations starting at location 'mv' at packed
// 32-bit integer indices stored in 'index' scaled by 'scale'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			addr[i+63:i] := v1[k+63:j]
//		ENDFOR
//
// Instruction: 'VSCATTERDPD'. Intrinsic: '_mm512_i32loscatter_pd'.
// Requires KNCNI.
func I32loscatterPd(mv uintptr, index x86.M512i, v1 x86.M512d, scale int)  {
	i32loscatterPd(uintptr(mv), [64]byte(index), [8]float64(v1), scale)
}

func i32loscatterPd(mv uintptr, index [64]byte, v1 [8]float64, scale int) 


// MaskI32loscatterPd: Stores 8 packed double-precision (64-bit) floating-point
// elements in 'v1' to memory locations starting at location 'mv' at packed
// 32-bit integer indices stored in 'index' scaled by 'scale'. Only those
// elements whose corresponding mask bit is set in writemask 'k' are written to
// memory. 
//
//		FOR j := 0 to 7
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				i := j*64
//				addr[i+63:i] := v1[k+63:j]
//			FI
//		ENDFOR
//
// Instruction: 'VSCATTERDPD'. Intrinsic: '_mm512_mask_i32loscatter_pd'.
// Requires KNCNI.
func MaskI32loscatterPd(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512d, scale int)  {
	maskI32loscatterPd(uintptr(mv), uint8(k), [64]byte(index), [8]float64(v1), scale)
}

func maskI32loscatterPd(mv uintptr, k uint8, index [64]byte, v1 [8]float64, scale int) 


// I32scatterEpi32: Scatter 32-bit integers from 'a' into memory using 32-bit
// indices. 32-bit elements are stored at addresses starting at 'base_addr' and
// offset by each 32-bit element in 'vindex' (each index is scaled by the
// factor in 'scale'). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			MEM[base_addr + SignExtend(vindex[i+31:i])*scale] := a[i+31:i]
//		ENDFOR
//
// Instruction: 'VPSCATTERDD'. Intrinsic: '_mm512_i32scatter_epi32'.
// Requires KNCNI.
func I32scatterEpi32(base_addr uintptr, vindex x86.M512i, a x86.M512i, scale int)  {
	i32scatterEpi32(uintptr(base_addr), [64]byte(vindex), [64]byte(a), scale)
}

func i32scatterEpi32(base_addr uintptr, vindex [64]byte, a [64]byte, scale int) 


// MaskI32scatterEpi32: Scatter 32-bit integers from 'a' into memory using
// 32-bit indices. 32-bit elements are stored at addresses starting at
// 'base_addr' and offset by each 32-bit element in 'vindex' (each index is
// scaled by the factor in 'scale') subject to mask 'k' (elements are not
// stored when the corresponding mask bit is not set). 'scale' should be 1, 2,
// 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				MEM[base_addr + SignExtend(vindex[i+31:i])*scale] := a[i+31:i]
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPSCATTERDD'. Intrinsic: '_mm512_mask_i32scatter_epi32'.
// Requires KNCNI.
func MaskI32scatterEpi32(base_addr uintptr, k x86.Mmask16, vindex x86.M512i, a x86.M512i, scale int)  {
	maskI32scatterEpi32(uintptr(base_addr), uint16(k), [64]byte(vindex), [64]byte(a), scale)
}

func maskI32scatterEpi32(base_addr uintptr, k uint16, vindex [64]byte, a [64]byte, scale int) 


// I32scatterPs: Scatter single-precision (32-bit) floating-point elements from
// 'a' into memory using 32-bit indices. 32-bit elements are stored at
// addresses starting at 'base_addr' and offset by each 32-bit element in
// 'vindex' (each index is scaled by the factor in 'scale'). 'scale' should be
// 1, 2, 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			MEM[base_addr + SignExtend(vindex[i+31:i])*scale] := a[i+31:i]
//		ENDFOR
//
// Instruction: 'VSCATTERDPS'. Intrinsic: '_mm512_i32scatter_ps'.
// Requires KNCNI.
func I32scatterPs(base_addr uintptr, vindex x86.M512i, a x86.M512, scale int)  {
	i32scatterPs(uintptr(base_addr), [64]byte(vindex), [16]float32(a), scale)
}

func i32scatterPs(base_addr uintptr, vindex [64]byte, a [16]float32, scale int) 


// MaskI32scatterPs: Scatter single-precision (32-bit) floating-point elements
// from 'a' into memory using 32-bit indices. 32-bit elements are stored at
// addresses starting at 'base_addr' and offset by each 32-bit element in
// 'vindex' (each index is scaled by the factor in 'scale') subject to mask 'k'
// (elements are not stored when the corresponding mask bit is not set).
// 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				MEM[base_addr + SignExtend(vindex[i+31:i])*scale] := a[i+31:i]
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VSCATTERDPS'. Intrinsic: '_mm512_mask_i32scatter_ps'.
// Requires KNCNI.
func MaskI32scatterPs(base_addr uintptr, k x86.Mmask16, vindex x86.M512i, a x86.M512, scale int)  {
	maskI32scatterPs(uintptr(base_addr), uint16(k), [64]byte(vindex), [16]float32(a), scale)
}

func maskI32scatterPs(base_addr uintptr, k uint16, vindex [64]byte, a [16]float32, scale int) 


// I64extgatherEpi32lo: Up-converts 8 single-precision (32-bit) memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale' using 'conv' to 32-bit integer elements and
// stores them in 'dst'. 'hint' indicates to the processor whether the data is
// non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_UPCONV_EPI32_NONE:
//				dst[i+31:i] := addr[i+31:i]
//			_MM_UPCONV_EPI32_UINT8:
//				n := j*8
//				dst[i+31:i] := UInt8ToInt32(addr[n+7:n])
//			_MM_UPCONV_EPI32_SINT8:
//				n := j*8
//				dst[i+31:i] := SInt8ToInt32(addr[n+7:n])
//			_MM_UPCONV_EPI32_UINT16:
//				n := j*16
//				dst[i+31:i] := UInt16ToInt32(addr[n+15:n])
//			_MM_UPCONV_EPI32_SINT16:
//				n := j*16
//				dst[i+31:i] := SInt16ToInt32(addr[n+15:n])
//			ESAC
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_i64extgather_epi32lo'.
// Requires KNCNI.
func I64extgatherEpi32lo(index x86.M512i, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(i64extgatherEpi32lo([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i64extgatherEpi32lo(index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte


// MaskI64extgatherEpi32lo: Up-converts 8 single-precision (32-bit) memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale' using 'conv' to 32-bit integer elements and
// stores them in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 'hint' indicates to the
// processor whether the data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_EPI32_NONE:
//					dst[i+31:i] := addr[i+31:i]
//				_MM_UPCONV_EPI32_UINT8:
//					n := j*8
//					dst[i+31:i] := UInt8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_SINT8:
//					n := j*8
//					dst[i+31:i] := SInt8ToInt32(addr[n+7:n])
//				_MM_UPCONV_EPI32_UINT16:
//					n := j*16
//					dst[i+31:i] := UInt16ToInt32(addr[n+15:n])
//				_MM_UPCONV_EPI32_SINT16:
//					n := j*16
//					dst[i+31:i] := SInt16ToInt32(addr[n+15:n])
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extgather_epi32lo'.
// Requires KNCNI.
func MaskI64extgatherEpi32lo(src x86.M512i, k x86.Mmask8, index x86.M512i, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(maskI64extgatherEpi32lo([64]byte(src), uint8(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI64extgatherEpi32lo(src [64]byte, k uint8, index [64]byte, mv uintptr, conv MMUPCONVEPI32ENUM, scale int, hint int) [64]byte


// I64extgatherEpi64: Up-converts 8 double-precision (64-bit) memory locations
// starting at location 'mv' at packed 64-bit integer indices stored in 'index'
// scaled by 'scale' using 'conv' to 64-bit integer elements and stores them in
// 'dst'. 'hint' indicates to the processor whether the load is non-temporal. 
//
//		FOR j := 0 to 7
//			i := j*64
//			addr := MEM[mv + index[j] * scale]
//			CASE conv OF
//			_MM_UPCONV_EPI64_NONE: dst[i+63:i] := addr[i+63:i]
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_i64extgather_epi64'.
// Requires KNCNI.
func I64extgatherEpi64(index x86.M512i, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(i64extgatherEpi64([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i64extgatherEpi64(index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte


// MaskI64extgatherEpi64: Up-converts 8 double-precision (64-bit) memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale' using 'conv' to 64-bit integer elements and
// stores them in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 'hint' indicates to the
// processor whether the load is non-temporal. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				CASE conv OF
//				_MM_UPCONV_EPI64_NONE: dst[i+63:i] := addr[i+63:i]
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extgather_epi64'.
// Requires KNCNI.
func MaskI64extgatherEpi64(src x86.M512i, k x86.Mmask8, index x86.M512i, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) x86.M512i {
	return x86.M512i(maskI64extgatherEpi64([64]byte(src), uint8(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI64extgatherEpi64(src [64]byte, k uint8, index [64]byte, mv uintptr, conv MMUPCONVEPI64ENUM, scale int, hint int) [64]byte


// I64extgatherPd: Up-converts 8 double-precision (64-bit) floating-point
// elements stored in memory starting at location 'mv' at packed 64-bit integer
// indices stored in 'index' scaled by 'scale' using 'conv' to 64-bit
// floating-point elements and stores them in 'dst'. 'hint' indicates to the
// processor whether the data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_UPCONV_PD_NONE: dst[i+63:i] := addr[i+63:i]
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_i64extgather_pd'.
// Requires KNCNI.
func I64extgatherPd(index x86.M512i, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) x86.M512d {
	return x86.M512d(i64extgatherPd([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i64extgatherPd(index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64


// MaskI64extgatherPd: Up-converts 8 double-precision (64-bit) floating-point
// elements stored in memory starting at location 'mv' at packed 64-bit integer
// indices stored in 'index' scaled by 'scale' using 'conv' to 64-bit
// floating-point elements and stores them in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 'hint' indicates to the processor whether the data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_PD_NONE: dst[i+63:i] := addr[i+63:i]
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extgather_pd'.
// Requires KNCNI.
func MaskI64extgatherPd(src x86.M512d, k x86.Mmask8, index x86.M512i, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) x86.M512d {
	return x86.M512d(maskI64extgatherPd([8]float64(src), uint8(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI64extgatherPd(src [8]float64, k uint8, index [64]byte, mv uintptr, conv MMUPCONVPDENUM, scale int, hint int) [8]float64


// I64extgatherPslo: Up-converts 8 memory locations starting at location 'mv'
// at packed 64-bit integer indices stored in 'index' scaled by 'scale' using
// 'conv' to single-precision (32-bit) floating-point elements and stores them
// in the lower half of 'dst'. 'hint' indicates to the processor whether the
// load is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_UPCONV_PS_NONE:
//				dst[i+31:i] := addr[i+31:i]
//			_MM_UPCONV_PS_FLOAT16:
//				n := j*16
//				dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//			_MM_UPCONV_PS_UINT8:
//				n := j*8
//				dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//			_MM_UPCONV_PS_SINT8:
//				n := j*8
//				dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//			_MM_UPCONV_PS_UINT16:
//				n := j*16
//				dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//			_MM_UPCONV_PS_SINT16:
//				n := j*16
//				dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//			ESAC
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_i64extgather_pslo'.
// Requires KNCNI.
func I64extgatherPslo(index x86.M512i, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) x86.M512 {
	return x86.M512(i64extgatherPslo([64]byte(index), uintptr(mv), conv, scale, hint))
}

func i64extgatherPslo(index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32


// MaskI64extgatherPslo: Up-converts 8 memory locations starting at location
// 'mv' at packed 64-bit integer indices stored in 'index' scaled by 'scale'
// using 'conv' to single-precision (32-bit) floating-point elements and stores
// them in the lower half of 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 'hint' indicates to
// the processor whether the load is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_UPCONV_PS_NONE:
//					dst[i+31:i] := addr[i+31:i]
//				_MM_UPCONV_PS_FLOAT16:
//					n := j*16
//					dst[i+31:i] := Float16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_UINT8:
//					n := j*8
//					dst[i+31:i] := UInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_SINT8:
//					n := j*8
//					dst[i+31:i] := SInt8ToFloat32(addr[n+7:n])
//				_MM_UPCONV_PS_UINT16:
//					n := j*16
//					dst[i+31:i] := UInt16ToFloat32(addr[n+15:n])
//				_MM_UPCONV_PS_SINT16:
//					n := j*16
//					dst[i+31:i] := SInt16ToFloat32(addr[n+15:n])
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extgather_pslo'.
// Requires KNCNI.
func MaskI64extgatherPslo(src x86.M512, k x86.Mmask8, index x86.M512i, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) x86.M512 {
	return x86.M512(maskI64extgatherPslo([16]float32(src), uint8(k), [64]byte(index), uintptr(mv), conv, scale, hint))
}

func maskI64extgatherPslo(src [16]float32, k uint8, index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) [16]float32


// I64extscatterEpi32lo: Down-converts the low 8 packed 32-bit integer elements
// in 'v1' using 'conv' and stores them in memory locations starting at
// location 'mv' at packed 64-bit integer indices stored in 'index' scaled by
// 'scale'. 'hint' indicates to the processor whether the data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_EPI32_NONE:
//				addr[i+31:i] := v1[i+31:i]
//			_MM_DOWNCONV_EPI32_UINT8:
//				n := j*8
//				addr[n+7:n] := UInt32ToUInt8(v1[i+31:i])
//			_MM_DOWNCONV_EPI32_SINT8:
//				n := j*8
//				addr[n+7:n] := SInt32ToSInt8(v1[i+31:i])
//			_MM_DOWNCONV_EPI32_UINT16:
//				n := j*16
//				addr[n+15:n] := UInt32ToUInt16(v1[i+31:i])
//			_MM_DOWNCONV_EPI32_SINT16:
//				n := j*16
//				addr[n+15:n] := SInt32ToSInt16(v1[n+15:n])
//			ESAC
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_i64extscatter_epi32lo'.
// Requires KNCNI.
func I64extscatterEpi32lo(mv uintptr, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, scale int, hint int)  {
	i64extscatterEpi32lo(uintptr(mv), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func i64extscatterEpi32lo(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 


// MaskI64extscatterEpi32lo: Down-converts the low 8 packed 32-bit integer
// elements in 'v1' using 'conv' and stores them in memory locations starting
// at location 'mv' at packed 64-bit integer indices stored in 'index' scaled
// by 'scale'. Elements are written to memory using writemask 'k' (elements are
// only written when the corresponding mask bit is set; otherwise, the memory
// location is left unchanged). 'hint' indicates to the processor whether the
// data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_EPI32_NONE:
//					addr[i+31:i] := v1[i+31:i]
//				_MM_DOWNCONV_EPI32_UINT8:
//					n := j*8
//					addr[n+7:n] := UInt32ToUInt8(v1[i+31:i])
//				_MM_DOWNCONV_EPI32_SINT8:
//					n := j*8
//					addr[n+7:n] := SInt32ToSInt8(v1[i+31:i])
//				_MM_DOWNCONV_EPI32_UINT16:
//					n := j*16
//					addr[n+15:n] := UInt32ToUInt16(v1[i+31:i])
//				_MM_DOWNCONV_EPI32_SINT16:
//					n := j*16
//					addr[n+15:n] := SInt32ToSInt16(v1[n+15:n])
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extscatter_epi32lo'.
// Requires KNCNI.
func MaskI64extscatterEpi32lo(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI32ENUM, scale int, hint int)  {
	maskI64extscatterEpi32lo(uintptr(mv), uint8(k), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func maskI64extscatterEpi32lo(mv uintptr, k uint8, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI32ENUM, scale int, hint int) 


// I64extscatterEpi64: Down-converts 8 packed 64-bit integer elements in 'v1'
// using 'conv' and stores them in memory locations starting at location 'mv'
// at packed 64-bit integer indices stored in 'index' scaled by 'scale'. 'hint'
// indicates to the processor whether the load is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_EPI64_NONE: addr[i+63:i] := v1[i+63:i]
//			ESAC
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_i64extscatter_epi64'.
// Requires KNCNI.
func I64extscatterEpi64(mv uintptr, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, scale int, hint int)  {
	i64extscatterEpi64(uintptr(mv), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func i64extscatterEpi64(mv uintptr, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 


// MaskI64extscatterEpi64: Down-converts 8 packed 64-bit integer elements in
// 'v1' using 'conv' and stores them in memory locations starting at location
// 'mv' at packed 64-bit integer indices stored in 'index' scaled by 'scale'.
// Only those elements whose corresponding mask bit is set in writemask 'k' are
// written to memory. 
//
//		FOR j := 0 to 7
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				i := j*64
//				CASE conv OF
//				_MM_DOWNCONV_EPI64_NONE: addr[i+63:i] := v1[i+63:i]
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extscatter_epi64'.
// Requires KNCNI.
func MaskI64extscatterEpi64(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512i, conv MMDOWNCONVEPI64ENUM, scale int, hint int)  {
	maskI64extscatterEpi64(uintptr(mv), uint8(k), [64]byte(index), [64]byte(v1), conv, scale, hint)
}

func maskI64extscatterEpi64(mv uintptr, k uint8, index [64]byte, v1 [64]byte, conv MMDOWNCONVEPI64ENUM, scale int, hint int) 


// I64extscatterPd: Down-converts 8 packed double-precision (64-bit)
// floating-point elements in 'v1' using 'conv' and stores them in memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale'. 'hint' indicates to the processor whether the
// data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			CASE conv OF
//			_MM_DOWNCONV_EPI64_NONE:
//				addr[i+63:i] := v1[i+63:i]
//			ESAC
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_i64extscatter_pd'.
// Requires KNCNI.
func I64extscatterPd(mv uintptr, index x86.M512i, v1 x86.M512d, conv MMDOWNCONVPDENUM, scale int, hint int)  {
	i64extscatterPd(uintptr(mv), [64]byte(index), [8]float64(v1), conv, scale, hint)
}

func i64extscatterPd(mv uintptr, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 


// MaskI64extscatterPd: Down-converts 8 packed double-precision (64-bit)
// floating-point elements in 'v1' using 'conv' and stores them in memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale'. Elements are written to memory using writemask
// 'k' (elements are not stored to memory when the corresponding mask bit is
// not set; the memory location is left unchagned). 'hint' indicates to the
// processor whether the data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*64
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_EPI64_NONE:
//					addr[i+63:i] := v1[i+63:i]
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extscatter_pd'.
// Requires KNCNI.
func MaskI64extscatterPd(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512d, conv MMDOWNCONVPDENUM, scale int, hint int)  {
	maskI64extscatterPd(uintptr(mv), uint8(k), [64]byte(index), [8]float64(v1), conv, scale, hint)
}

func maskI64extscatterPd(mv uintptr, k uint8, index [64]byte, v1 [8]float64, conv MMDOWNCONVPDENUM, scale int, hint int) 


// I64extscatterPslo: Down-converts 8 packed single-precision (32-bit)
// floating-point elements in 'v1' using 'conv' and stores them in memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale'. 'hint' indicates to the processor whether the
// data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE conv OF
//			_MM_DOWNCONV_PS_NONE:
//				addr[i+31:i] := v1[i+31:i]
//			_MM_DOWNCONV_PS_FLOAT16:
//				n := j*16
//				addr[n+15:n] := Float32ToFloat16(v1[i+31:i])
//			_MM_DOWNCONV_PS_UINT8:
//				n := j*8
//				addr[n+7:n] := Float32ToUInt8(v1[i+31:i])
//			_MM_DOWNCONV_PS_SINT8:
//				n := j*8
//				addr[n+7:n] := Float32ToSInt8(v1[i+31:i])
//			_MM_DOWNCONV_PS_UINT16:
//				n := j*16
//				addr[n+15:n] := Float32ToUInt16(v1[i+31:i])
//			_MM_DOWNCONV_PS_SINT16:
//				n := j*16
//				addr[n+15:n] := Float32ToSInt16(v1[i+31:i])
//			ESAC
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_i64extscatter_pslo'.
// Requires KNCNI.
func I64extscatterPslo(mv uintptr, index x86.M512i, v1 x86.M512, conv MMDOWNCONVPSENUM, scale int, hint int)  {
	i64extscatterPslo(uintptr(mv), [64]byte(index), [16]float32(v1), conv, scale, hint)
}

func i64extscatterPslo(mv uintptr, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int, hint int) 


// MaskI64extscatterPslo: Down-converts 8 packed single-precision (32-bit)
// floating-point elements in 'v1' using 'conv' and stores them in memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale'. Elements are only written when the
// corresponding mask bit is set in 'k'; otherwise, elements are unchanged in
// memory. 'hint' indicates to the processor whether the data is non-temporal. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j]
//				CASE conv OF
//				_MM_DOWNCONV_PS_NONE:
//					addr[i+31:i] := v[i+31:i]
//				_MM_DOWNCONV_PS_FLOAT16:
//					n := j*16
//					addr[n+15:n] := Float32ToFloat16(v1[i+31:i])
//				_MM_DOWNCONV_PS_UINT8:
//					n := j*8
//					addr[n+7:n] := Float32ToUInt8(v1[i+31:i])
//				_MM_DOWNCONV_PS_SINT8:
//					n := j*8
//					addr[n+7:n] := Float32ToSInt8(v1[i+31:i])
//				_MM_DOWNCONV_PS_UINT16:
//					n := j*16
//					addr[n+15:n] := Float32ToUInt16(v1[i+31:i])
//				_MM_DOWNCONV_PS_SINT16:
//					n := j*16
//					addr[n+15:n] := Float32ToSInt16(v1[i+31:i])
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64extscatter_pslo'.
// Requires KNCNI.
func MaskI64extscatterPslo(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512, conv MMDOWNCONVPSENUM, scale int, hint int)  {
	maskI64extscatterPslo(uintptr(mv), uint8(k), [64]byte(index), [16]float32(v1), conv, scale, hint)
}

func maskI64extscatterPslo(mv uintptr, k uint8, index [64]byte, v1 [16]float32, conv MMDOWNCONVPSENUM, scale int, hint int) 


// I64gatherEpi32lo: Loads 8 32-bit integer memory locations starting at
// location 'mv' at packed 64-bit integer indices stored in 'index' scaled by
// 'scale' to 'dst'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			dst[i+31:i] := addr[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_i64gather_epi32lo'.
// Requires KNCNI.
func I64gatherEpi32lo(index x86.M512i, mv uintptr, scale int) x86.M512i {
	return x86.M512i(i64gatherEpi32lo([64]byte(index), uintptr(mv), scale))
}

func i64gatherEpi32lo(index [64]byte, mv uintptr, scale int) [64]byte


// MaskI64gatherEpi32lo: Loads 8 32-bit integer memory locations starting at
// location 'mv' at packed 64-bit integer indices stored in 'index' scaled by
// 'scale' to 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				dst[i+31:i] := addr[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64gather_epi32lo'.
// Requires KNCNI.
func MaskI64gatherEpi32lo(src x86.M512i, k x86.Mmask8, index x86.M512i, mv uintptr, scale int) x86.M512i {
	return x86.M512i(maskI64gatherEpi32lo([64]byte(src), uint8(k), [64]byte(index), uintptr(mv), scale))
}

func maskI64gatherEpi32lo(src [64]byte, k uint8, index [64]byte, mv uintptr, scale int) [64]byte


// I64gatherPslo: Loads 8 single-precision (32-bit) floating-point memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale' to 'dst'. 
//
//		FOR j := 0 to 7
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			dst[i+31:i] := addr[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_i64gather_pslo'.
// Requires KNCNI.
func I64gatherPslo(index x86.M512i, mv uintptr, scale int) x86.M512 {
	return x86.M512(i64gatherPslo([64]byte(index), uintptr(mv), scale))
}

func i64gatherPslo(index [64]byte, mv uintptr, scale int) [16]float32


// MaskI64gatherPslo: Loads 8 single-precision (32-bit) floating-point memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale' to 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				dst[i+31:i] := addr[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64gather_pslo'.
// Requires KNCNI.
func MaskI64gatherPslo(src x86.M512, k x86.Mmask8, index x86.M512i, mv uintptr, scale int) x86.M512 {
	return x86.M512(maskI64gatherPslo([16]float32(src), uint8(k), [64]byte(index), uintptr(mv), scale))
}

func maskI64gatherPslo(src [16]float32, k uint8, index [64]byte, mv uintptr, scale int) [16]float32


// I64scatterEpi32lo: Stores 8 packed 32-bit integer elements in 'v1' in memory
// locations starting at location 'mv' at packed 64-bit integer indices stored
// in 'index' scaled by 'scale'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			addr := MEM[mv + index[j] * scale]
//			addr[i+31:i] := v1[i+31:i]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_i64scatter_epi32lo'.
// Requires KNCNI.
func I64scatterEpi32lo(mv uintptr, index x86.M512i, v1 x86.M512i, scale int)  {
	i64scatterEpi32lo(uintptr(mv), [64]byte(index), [64]byte(v1), scale)
}

func i64scatterEpi32lo(mv uintptr, index [64]byte, v1 [64]byte, scale int) 


// MaskI64scatterEpi32lo: Stores 8 packed 32-bit integer elements in 'v1' in
// memory locations starting at location 'mv' at packed 64-bit integer indices
// stored in 'index' scaled by 'scale' using writemask 'k' (elements are only
// written to memory when the corresponding mask bit is set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				addr[i+31:i] := v1[i+31:i]
//			FI	
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64scatter_epi32lo'.
// Requires KNCNI.
func MaskI64scatterEpi32lo(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512i, scale int)  {
	maskI64scatterEpi32lo(uintptr(mv), uint8(k), [64]byte(index), [64]byte(v1), scale)
}

func maskI64scatterEpi32lo(mv uintptr, k uint8, index [64]byte, v1 [64]byte, scale int) 


// I64scatterPslo: Stores 8 packed single-precision (32-bit) floating-point
// elements in 'v' in memory locations starting at location 'mv' at packed
// 64-bit integer indices stored in 'index' scaled by 'scale'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			addr := MEM[mv + index[j] * scale]
//			addr[i+31:i] := v[i+31:i]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_i64scatter_pslo'.
// Requires KNCNI.
func I64scatterPslo(mv uintptr, index x86.M512i, v x86.M512, scale int)  {
	i64scatterPslo(uintptr(mv), [64]byte(index), [16]float32(v), scale)
}

func i64scatterPslo(mv uintptr, index [64]byte, v [16]float32, scale int) 


// MaskI64scatterPslo: Stores 8 packed single-precision (32-bit) floating-point
// elements in 'v1' in memory locations starting at location 'mv' at packed
// 64-bit integer indices stored in 'index' scaled by 'scale' using writemask
// 'k' (elements are only written to memory when the corresponding mask bit is
// set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				addr[i+31:i] := v1[i+31:i]
//			FI	
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm512_mask_i64scatter_pslo'.
// Requires KNCNI.
func MaskI64scatterPslo(mv uintptr, k x86.Mmask8, index x86.M512i, v1 x86.M512, scale int)  {
	maskI64scatterPslo(uintptr(mv), uint8(k), [64]byte(index), [16]float32(v1), scale)
}

func maskI64scatterPslo(mv uintptr, k uint8, index [64]byte, v1 [16]float32, scale int) 


// Int2mask: Converts integer 'mask' into bitmask, storing the result in 'dst'. 
//
//		dst := mask[15:0]
//
// Instruction: 'KMOV'. Intrinsic: '_mm512_int2mask'.
// Requires KNCNI.
func Int2mask(mask int) x86.Mmask16 {
	return x86.Mmask16(int2mask(mask))
}

func int2mask(mask int) uint16


// Kand: Compute the bitwise AND of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := a[15:0] AND b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KAND'. Intrinsic: '_mm512_kand'.
// Requires KNCNI.
func Kand(a x86.Mmask16, b x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kand(uint16(a), uint16(b)))
}

func kand(a uint16, b uint16) uint16


// Kandn: Compute the bitwise AND NOT of 16-bit masks 'a' and 'b', and store
// the result in 'k'. 
//
//		k[15:0] := (NOT a[15:0]) AND b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KANDN'. Intrinsic: '_mm512_kandn'.
// Requires KNCNI.
func Kandn(a x86.Mmask16, b x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kandn(uint16(a), uint16(b)))
}

func kandn(a uint16, b uint16) uint16


// Kandnr: Performs a bitwise AND operation between NOT of 'k2' and 'k1',
// storing the result in 'dst'. 
//
//		dst[15:0] := NOT(k2[15:0]) & k1[15:0]
//
// Instruction: 'KANDNR'. Intrinsic: '_mm512_kandnr'.
// Requires KNCNI.
func Kandnr(k1 x86.Mmask16, k2 x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kandnr(uint16(k1), uint16(k2)))
}

func kandnr(k1 uint16, k2 uint16) uint16


// Kconcathi64: Packs masks 'k1' and 'k2' into the high 32 bits of 'dst'. The
// rest of 'dst' is set to 0. 
//
//		dst[63:48] := k1[15:0]
//		dst[47:32] := k2[15:0]
//		dst[31:0]  := 0
//
// Instruction: 'KCONCATH'. Intrinsic: '_mm512_kconcathi_64'.
// Requires KNCNI.
func Kconcathi64(k1 x86.Mmask16, k2 x86.Mmask16) int64 {
	return int64(kconcathi64(uint16(k1), uint16(k2)))
}

func kconcathi64(k1 uint16, k2 uint16) int64


// Kconcatlo64: Packs masks 'k1' and 'k2' into the low 32 bits of 'dst'. The
// rest of 'dst' is set to 0. 
//
//		dst[31:16] := k1[15:0]
//		dst[15:0]  := k2[15:0]
//		dst[63:32] := 0
//
// Instruction: 'KCONCATL'. Intrinsic: '_mm512_kconcatlo_64'.
// Requires KNCNI.
func Kconcatlo64(k1 x86.Mmask16, k2 x86.Mmask16) int64 {
	return int64(kconcatlo64(uint16(k1), uint16(k2)))
}

func kconcatlo64(k1 uint16, k2 uint16) int64


// Kextract64: Extracts 16-bit value 'b' from 64-bit integer 'a', storing the
// result in 'dst'. 
//
//		CASE b of
//		0: dst[15:0] := a[63:48]
//		1: dst[15:0] := a[47:32]
//		2: dst[15:0] := a[31:16]
//		3: dst[15:0] := a[15:0]
//		ESAC
//		dst[MAX:15] := 0
//
// Instruction: 'KEXTRACT'. Intrinsic: '_mm512_kextract_64'.
// Requires KNCNI.
func Kextract64(a int64, b int) x86.Mmask16 {
	return x86.Mmask16(kextract64(a, b))
}

func kextract64(a int64, b int) uint16


// Kmerge2l1h: Move the high element from 'k1' to the low element of 'k1', and
// insert the low element of 'k2' into the high element of 'k1'. 
//
//		tmp[7:0] := k1[15:8]
//		k1[15:8] := k2[7:0]
//		k1[7:0]  := tmp[7:0]
//
// Instruction: 'KMERGE2L1H'. Intrinsic: '_mm512_kmerge2l1h'.
// Requires KNCNI.
func Kmerge2l1h(k1 x86.Mmask16, k2 x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kmerge2l1h(uint16(k1), uint16(k2)))
}

func kmerge2l1h(k1 uint16, k2 uint16) uint16


// Kmerge2l1l: Insert the low element of 'k2' into the high element of 'k1'. 
//
//		k1[15:8] := k2[7:0]
//
// Instruction: 'KMERGE2L1L'. Intrinsic: '_mm512_kmerge2l1l'.
// Requires KNCNI.
func Kmerge2l1l(k1 x86.Mmask16, k2 x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kmerge2l1l(uint16(k1), uint16(k2)))
}

func kmerge2l1l(k1 uint16, k2 uint16) uint16


// Kmov: Copy 16-bit mask 'a' to 'k'. 
//
//		k[15:0] := a[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KMOV'. Intrinsic: '_mm512_kmov'.
// Requires KNCNI.
func Kmov(a x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kmov(uint16(a)))
}

func kmov(a uint16) uint16


// Kmovlhb: Inserts the low byte of mask 'k2' into the high byte of 'dst', and
// copies the low byte of 'k1' to the low byte of 'dst'. 
//
//		dst[7:0] := k1[7:0]
//		dst[15:8] := k2[7:0]
//
// Instruction: 'KMERGE2L1L'. Intrinsic: '_mm512_kmovlhb'.
// Requires KNCNI.
func Kmovlhb(k1 x86.Mmask16, k2 x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kmovlhb(uint16(k1), uint16(k2)))
}

func kmovlhb(k1 uint16, k2 uint16) uint16


// Knot: Compute the bitwise NOT of 16-bit mask 'a', and store the result in
// 'k'. 
//
//		k[15:0] := NOT a[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KNOT'. Intrinsic: '_mm512_knot'.
// Requires KNCNI.
func Knot(a x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(knot(uint16(a)))
}

func knot(a uint16) uint16


// Kor: Compute the bitwise OR of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := a[15:0] OR b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KOR'. Intrinsic: '_mm512_kor'.
// Requires KNCNI.
func Kor(a x86.Mmask16, b x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kor(uint16(a), uint16(b)))
}

func kor(a uint16, b uint16) uint16


// Kortestc: Performs bitwise OR between 'k1' and 'k2', storing the result in
// 'dst'. CF flag is set if 'dst' consists of all 1's. 
//
//		dst[15:0] := k1[15:0] | k2[15:0]
//		IF PopCount(dst[15:0]) = 16
//			SetCF()
//		FI
//
// Instruction: 'KORTEST'. Intrinsic: '_mm512_kortestc'.
// Requires KNCNI.
func Kortestc(k1 x86.Mmask16, k2 x86.Mmask16) int {
	return int(kortestc(uint16(k1), uint16(k2)))
}

func kortestc(k1 uint16, k2 uint16) int


// Kortestz: Performs bitwise OR between 'k1' and 'k2', storing the result in
// 'dst'. ZF flag is set if 'dst' is 0. 
//
//		dst[15:0] := k1[15:0] | k2[15:0]
//		IF dst = 0
//			SetZF()
//		FI
//
// Instruction: 'KORTEST'. Intrinsic: '_mm512_kortestz'.
// Requires KNCNI.
func Kortestz(k1 x86.Mmask16, k2 x86.Mmask16) int {
	return int(kortestz(uint16(k1), uint16(k2)))
}

func kortestz(k1 uint16, k2 uint16) int


// Kswapb: Moves high byte from 'k2' to low byte of 'k1', and moves low byte of
// 'k2' to high byte of 'k1'. 
//
//		tmp[7:0] := k2[15:8]
//		k2[15:8] := k1[7:0]
//		k1[7:0]  := tmp[7:0]
//		
//		tmp[7:0] := k2[7:0]
//		k2[7:0]  := k1[15:8]
//		k1[15:8] := tmp[7:0]
//
// Instruction: 'KMERGE2L1H'. Intrinsic: '_mm512_kswapb'.
// Requires KNCNI.
func Kswapb(k1 x86.Mmask16, k2 x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kswapb(uint16(k1), uint16(k2)))
}

func kswapb(k1 uint16, k2 uint16) uint16


// Kxnor: Compute the bitwise XNOR of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := NOT (a[15:0] XOR b[15:0])
//		k[MAX:16] := 0
//
// Instruction: 'KXNOR'. Intrinsic: '_mm512_kxnor'.
// Requires KNCNI.
func Kxnor(a x86.Mmask16, b x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kxnor(uint16(a), uint16(b)))
}

func kxnor(a uint16, b uint16) uint16


// Kxor: Compute the bitwise XOR of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := a[15:0] XOR b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KXOR'. Intrinsic: '_mm512_kxor'.
// Requires KNCNI.
func Kxor(a x86.Mmask16, b x86.Mmask16) x86.Mmask16 {
	return x86.Mmask16(kxor(uint16(a), uint16(b)))
}

func kxor(a uint16, b uint16) uint16


// LoadEpi32: Load 512-bits (composed of 16 packed 32-bit integers) from memory
// into 'dst'. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[511:0] := MEM[mem_addr+511:mem_addr]
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_load_epi32'.
// Requires KNCNI.
func LoadEpi32(mem_addr uintptr) x86.M512i {
	return x86.M512i(loadEpi32(uintptr(mem_addr)))
}

func loadEpi32(mem_addr uintptr) [64]byte


// MaskLoadEpi32: Load packed 32-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set).
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := MEM[mem_addr+i+31:mem_addr+i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_mask_load_epi32'.
// Requires KNCNI.
func MaskLoadEpi32(src x86.M512i, k x86.Mmask16, mem_addr uintptr) x86.M512i {
	return x86.M512i(maskLoadEpi32([64]byte(src), uint16(k), uintptr(mem_addr)))
}

func maskLoadEpi32(src [64]byte, k uint16, mem_addr uintptr) [64]byte


// LoadEpi64: Load 512-bits (composed of 8 packed 64-bit integers) from memory
// into 'dst'. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[511:0] := MEM[mem_addr+511:mem_addr]
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_load_epi64'.
// Requires KNCNI.
func LoadEpi64(mem_addr uintptr) x86.M512i {
	return x86.M512i(loadEpi64(uintptr(mem_addr)))
}

func loadEpi64(mem_addr uintptr) [64]byte


// MaskLoadEpi64: Load packed 64-bit integers from memory into 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := MEM[mem_addr+i+63:mem_addr+i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_mask_load_epi64'.
// Requires KNCNI.
func MaskLoadEpi64(src x86.M512i, k x86.Mmask8, mem_addr uintptr) x86.M512i {
	return x86.M512i(maskLoadEpi64([64]byte(src), uint8(k), uintptr(mem_addr)))
}

func maskLoadEpi64(src [64]byte, k uint8, mem_addr uintptr) [64]byte


// LoadPd: Load 512-bits (composed of 8 packed double-precision (64-bit)
// floating-point elements) from memory into 'dst'. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[511:0] := MEM[mem_addr+511:mem_addr]
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_load_pd'.
// Requires KNCNI.
func LoadPd(mem_addr uintptr) x86.M512d {
	return x86.M512d(loadPd(uintptr(mem_addr)))
}

func loadPd(mem_addr uintptr) [8]float64


// MaskLoadPd: Load packed double-precision (64-bit) floating-point elements
// from memory into 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 'mem_addr' must be aligned on a
// 64-byte boundary or a general-protection exception may be generated. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := MEM[mem_addr+i+63:mem_addr+i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_mask_load_pd'.
// Requires KNCNI.
func MaskLoadPd(src x86.M512d, k x86.Mmask8, mem_addr uintptr) x86.M512d {
	return x86.M512d(maskLoadPd([8]float64(src), uint8(k), uintptr(mem_addr)))
}

func maskLoadPd(src [8]float64, k uint8, mem_addr uintptr) [8]float64


// LoadPs: Load 512-bits (composed of 16 packed single-precision (32-bit)
// floating-point elements) from memory into 'dst'. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[511:0] := MEM[mem_addr+511:mem_addr]
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_load_ps'.
// Requires KNCNI.
func LoadPs(mem_addr uintptr) x86.M512 {
	return x86.M512(loadPs(uintptr(mem_addr)))
}

func loadPs(mem_addr uintptr) [16]float32


// MaskLoadPs: Load packed single-precision (32-bit) floating-point elements
// from memory into 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 'mem_addr' must be aligned on a
// 64-byte boundary or a general-protection exception may be generated. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := MEM[mem_addr+i+31:mem_addr+i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_mask_load_ps'.
// Requires KNCNI.
func MaskLoadPs(src x86.M512, k x86.Mmask16, mem_addr uintptr) x86.M512 {
	return x86.M512(maskLoadPs([16]float32(src), uint16(k), uintptr(mem_addr)))
}

func maskLoadPs(src [16]float32, k uint16, mem_addr uintptr) [16]float32


// LoadSi512: Load 512-bits of integer data from memory into 'dst'. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[511:0] := MEM[mem_addr+511:mem_addr]
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_load_si512'.
// Requires KNCNI.
func LoadSi512(mem_addr uintptr) x86.M512i {
	return x86.M512i(loadSi512(uintptr(mem_addr)))
}

func loadSi512(mem_addr uintptr) [64]byte


// LoadunpackhiEpi32: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64 and
// expands them into packed 32-bit integers in 'dst'. The initial values of
// 'dst' are copied from 'src'. Only those converted doublewords that occur at
// or after the first 64-byte-aligned address following (mt-64) are loaded.
// Elements in the resulting vector that do not map to those doublewords are
// taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*4 % 64) == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*32
//				tmp := MEM[addr + loadOffset*4]
//				dst[i+31:i] := tmp[i+31:i]
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHD'. Intrinsic: '_mm512_loadunpackhi_epi32'.
// Requires KNCNI.
func LoadunpackhiEpi32(src x86.M512i, mt uintptr) x86.M512i {
	return x86.M512i(loadunpackhiEpi32([64]byte(src), uintptr(mt)))
}

func loadunpackhiEpi32(src [64]byte, mt uintptr) [64]byte


// MaskLoadunpackhiEpi32: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64 and
// expands them into packed 32-bit integers in 'dst'. The initial values of
// 'dst' are copied from 'src'. Only those converted doublewords that occur at
// or after the first 64-byte-aligned address following (mt-64) are loaded.
// Elements in the resulting vector that do not map to those doublewords are
// taken from 'src'. Elements are loaded from memory according to element
// selector 'k' (elements are skipped when the corresponding mask bit is not
// set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*4 % 64) == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*32
//					tmp := MEM[addr + loadOffset*4]
//					dst[i+31:i] := tmp[i+31:i]
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHD'. Intrinsic: '_mm512_mask_loadunpackhi_epi32'.
// Requires KNCNI.
func MaskLoadunpackhiEpi32(src x86.M512i, k x86.Mmask16, mt uintptr) x86.M512i {
	return x86.M512i(maskLoadunpackhiEpi32([64]byte(src), uint16(k), uintptr(mt)))
}

func maskLoadunpackhiEpi32(src [64]byte, k uint16, mt uintptr) [64]byte


// LoadunpackhiEpi64: Loads the high-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt-64 and expands them into
// packed 64-bit integers in 'dst'. The initial values of 'dst' are copied from
// 'src'. Only those converted quadwords that occur at or after the first
// 64-byte-aligned address following (mt-64) are loaded. Elements in the
// resulting vector that do not map to those quadwords are taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*8) == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*64
//				tmp := MEM[addr + loadOffset*8]
//				dst[i+63:i] := tmp[i+63:i]
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHQ'. Intrinsic: '_mm512_loadunpackhi_epi64'.
// Requires KNCNI.
func LoadunpackhiEpi64(src x86.M512i, mt uintptr) x86.M512i {
	return x86.M512i(loadunpackhiEpi64([64]byte(src), uintptr(mt)))
}

func loadunpackhiEpi64(src [64]byte, mt uintptr) [64]byte


// MaskLoadunpackhiEpi64: Loads the high-64-byte-aligned portion of the
// quadword stream starting at element-aligned address mt-64 and expands them
// into packed 64-bit integers in 'dst'. The initial values of 'dst' are copied
// from 'src'. Only those converted quadwords that occur at or after the first
// 64-byte-aligned address following (mt-64) are loaded. Elements in the
// resulting vector that do not map to those quadwords are taken from 'src'.
// Elements are loaded from memory according to element selector 'k' (elements
// are skipped when the corresponding mask bit is not set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*8) == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*64
//					tmp := MEM[addr + loadOffset*8]
//					dst[i+63:i] := tmp[i+63:i]
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHQ'. Intrinsic: '_mm512_mask_loadunpackhi_epi64'.
// Requires KNCNI.
func MaskLoadunpackhiEpi64(src x86.M512i, k x86.Mmask8, mt uintptr) x86.M512i {
	return x86.M512i(maskLoadunpackhiEpi64([64]byte(src), uint8(k), uintptr(mt)))
}

func maskLoadunpackhiEpi64(src [64]byte, k uint8, mt uintptr) [64]byte


// LoadunpackhiPd: Loads the high-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt-64 and expands them into
// packed double-precision (64-bit) floating-point values in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted quadwords that
// occur at or after the first 64-byte-aligned address following (mt-64) are
// loaded. Elements in the resulting vector that do not map to those quadwords
// are taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*8) % 64 == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*64
//				tmp := MEM[addr + loadOffset*8]
//				dst[i+63:i] := tmp[i+63:i]
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPD'. Intrinsic: '_mm512_loadunpackhi_pd'.
// Requires KNCNI.
func LoadunpackhiPd(src x86.M512d, mt uintptr) x86.M512d {
	return x86.M512d(loadunpackhiPd([8]float64(src), uintptr(mt)))
}

func loadunpackhiPd(src [8]float64, mt uintptr) [8]float64


// MaskLoadunpackhiPd: Loads the high-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt-64 and expands them into
// packed double-precision (64-bit) floating-point values in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted quadwords that
// occur at or after the first 64-byte-aligned address following (mt-64) are
// loaded. Elements in the resulting vector that do not map to those quadwords
// are taken from 'src'. Elements are loaded from memory according to element
// selector 'k' (elements are skipped when the corresponding mask bit is not
// set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*8) % 64 == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*64
//					tmp := MEM[addr + loadOffset*8]
//					dst[i+63:i] := tmp[i+63:i]
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPD'. Intrinsic: '_mm512_mask_loadunpackhi_pd'.
// Requires KNCNI.
func MaskLoadunpackhiPd(src x86.M512d, k x86.Mmask8, mt uintptr) x86.M512d {
	return x86.M512d(maskLoadunpackhiPd([8]float64(src), uint8(k), uintptr(mt)))
}

func maskLoadunpackhiPd(src [8]float64, k uint8, mt uintptr) [8]float64


// LoadunpackhiPs: Loads the high-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt-64 and
// expands them into packed single-precision (32-bit) floating-point elements
// in 'dst'. The initial values of 'dst' are copied from 'src'. Only those
// converted quadwords that occur at or after the first 64-byte-aligned address
// following (mt-64) are loaded. Elements in the resulting vector that do not
// map to those quadwords are taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF (addr + (loadOffset + 1)*4 % 64) == 0
//					foundNext64BytesBoundary := true
//				FI
//			ELSE
//				i := j*32
//				tmp := MEM[addr + loadOffset*4]
//				dst[i+31:i] := tmp[i+31:i]
//			FI
//			loadOffset := loadOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPS'. Intrinsic: '_mm512_loadunpackhi_ps'.
// Requires KNCNI.
func LoadunpackhiPs(src x86.M512, mt uintptr) x86.M512 {
	return x86.M512(loadunpackhiPs([16]float32(src), uintptr(mt)))
}

func loadunpackhiPs(src [16]float32, mt uintptr) [16]float32


// MaskLoadunpackhiPs: Loads the high-64-byte-aligned portion of the doubleword
// stream starting at element-aligned address mt-64 and expands them into
// packed single-precision (32-bit) floating-point elements in 'dst'. The
// initial values of 'dst' are copied from 'src'. Only those converted
// quadwords that occur at or after the first 64-byte-aligned address following
// (mt-64) are loaded. Elements in the resulting vector that do not map to
// those quadwords are taken from 'src'. Elements are loaded from memory
// according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF (addr + (loadOffset + 1)*4 % 64) == 0
//						foundNext64BytesBoundary := true
//					FI
//				ELSE
//					i := j*32
//					tmp := MEM[addr + loadOffset*4]
//					dst[i+31:i] := tmp[i+31:i]
//				FI
//				loadOffset := loadOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKHPS'. Intrinsic: '_mm512_mask_loadunpackhi_ps'.
// Requires KNCNI.
func MaskLoadunpackhiPs(src x86.M512, k x86.Mmask16, mt uintptr) x86.M512 {
	return x86.M512(maskLoadunpackhiPs([16]float32(src), uint16(k), uintptr(mt)))
}

func maskLoadunpackhiPs(src [16]float32, k uint16, mt uintptr) [16]float32


// LoadunpackloEpi32: Loads the low-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt and
// expanded into packed 32-bit integers in 'dst'. The initial values of 'dst'
// are copied from 'src'. Only those converted doublewords that occur before
// first 64-byte-aligned address following 'mt' are loaded. Elements in the
// resulting vector that do not map to those doublewords are taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			tmp := MEM[addr + loadOffset*4]
//			dst[i+31:i] := tmp[i+31:i]
//			loadOffset := loadOffset + 1
//			IF (mt + loadOffset * 4) % 64 == 0
//				break
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLD'. Intrinsic: '_mm512_loadunpacklo_epi32'.
// Requires KNCNI.
func LoadunpackloEpi32(src x86.M512i, mt uintptr) x86.M512i {
	return x86.M512i(loadunpackloEpi32([64]byte(src), uintptr(mt)))
}

func loadunpackloEpi32(src [64]byte, mt uintptr) [64]byte


// MaskLoadunpackloEpi32: Loads the low-64-byte-aligned portion of the
// byte/word/doubleword stream starting at element-aligned address mt and
// expands them into packed 32-bit integers in 'dst'. The initial values of
// 'dst' are copied from 'src'. Only those converted doublewords that occur
// before first 64-byte-aligned address following 'mt' are loaded. Elements in
// the resulting vector that do not map to those doublewords are taken from
// 'src'. Elements are loaded from memory according to element selector 'k'
// (elements are skipped when the corresponding mask bit is not set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				tmp := MEM[addr + loadOffset*4]
//				dst[i+31:i] := tmp[i+31:i]
//				loadOffset := loadOffset + 1
//				IF (mt + loadOffset * 4) % 64 == 0
//					break
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLD'. Intrinsic: '_mm512_mask_loadunpacklo_epi32'.
// Requires KNCNI.
func MaskLoadunpackloEpi32(src x86.M512i, k x86.Mmask16, mt uintptr) x86.M512i {
	return x86.M512i(maskLoadunpackloEpi32([64]byte(src), uint16(k), uintptr(mt)))
}

func maskLoadunpackloEpi32(src [64]byte, k uint16, mt uintptr) [64]byte


// LoadunpackloEpi64: Loads the low-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt and expands them into packed
// 64-bit integers in 'dst'. The initial values of 'dst' are copied from 'src'.
// Only those converted quad that occur before first 64-byte-aligned address
// following 'mt' are loaded. Elements in the resulting vector that do not map
// to those quadwords are taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			tmp := MEM[addr + loadOffset*8]
//			dst[i+63:i] := tmp[i+63:i]
//			loadOffset := loadOffset + 1
//			IF (addr + loadOffset*8 % 64) == 0
//				break
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLQ'. Intrinsic: '_mm512_loadunpacklo_epi64'.
// Requires KNCNI.
func LoadunpackloEpi64(src x86.M512i, mt uintptr) x86.M512i {
	return x86.M512i(loadunpackloEpi64([64]byte(src), uintptr(mt)))
}

func loadunpackloEpi64(src [64]byte, mt uintptr) [64]byte


// MaskLoadunpackloEpi64: Loads the low-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt and expands them into packed
// 64-bit integers in 'dst'. The initial values of 'dst' are copied from 'src'.
// Only those converted quad that occur before first 64-byte-aligned address
// following 'mt' are loaded. Elements in the resulting vector that do not map
// to those quadwords are taken from 'src'. Elements are loaded from memory
// according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp := MEM[addr + loadOffset*8]
//				dst[i+63:i] := tmp[i+63:i]
//				loadOffset := loadOffset + 1
//				IF (addr + loadOffset*8 % 64) == 0
//					break
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLQ'. Intrinsic: '_mm512_mask_loadunpacklo_epi64'.
// Requires KNCNI.
func MaskLoadunpackloEpi64(src x86.M512i, k x86.Mmask8, mt uintptr) x86.M512i {
	return x86.M512i(maskLoadunpackloEpi64([64]byte(src), uint8(k), uintptr(mt)))
}

func maskLoadunpackloEpi64(src [64]byte, k uint8, mt uintptr) [64]byte


// LoadunpackloPd: Loads the low-64-byte-aligned portion of the quadword stream
// starting at element-aligned address mt and expands them into packed
// double-precision (64-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted quad that occur
// before first 64-byte-aligned address following 'mt' are loaded. Elements in
// the resulting vector that do not map to those quadwords are taken from
// 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			tmp := MEM[addr + loadOffset*8]
//			dst[i+63:i] := tmp[i+63:i]
//			loadOffset := loadOffset + 1
//			IF ((addr + 8*loadOffset) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPD'. Intrinsic: '_mm512_loadunpacklo_pd'.
// Requires KNCNI.
func LoadunpackloPd(src x86.M512d, mt uintptr) x86.M512d {
	return x86.M512d(loadunpackloPd([8]float64(src), uintptr(mt)))
}

func loadunpackloPd(src [8]float64, mt uintptr) [8]float64


// MaskLoadunpackloPd: Loads the low-64-byte-aligned portion of the quadword
// stream starting at element-aligned address mt and expands them into packed
// double-precision (64-bit) floating-point values in 'dst'. The initial values
// of 'dst' are copied from 'src'. Only those converted quad that occur before
// first 64-byte-aligned address following 'mt' are loaded. Elements in the
// resulting vector that do not map to those quadwords are taken from 'src'.
// Elements are loaded from memory according to element selector 'k' (elements
// are skipped when the corresponding mask bit is not set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp := MEM[addr + loadOffset*8]
//				dst[i+63:i] := tmp[i+63:i]
//				loadOffset := loadOffset + 1
//				IF ((addr + 8*loadOffset) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPD'. Intrinsic: '_mm512_mask_loadunpacklo_pd'.
// Requires KNCNI.
func MaskLoadunpackloPd(src x86.M512d, k x86.Mmask8, mt uintptr) x86.M512d {
	return x86.M512d(maskLoadunpackloPd([8]float64(src), uint8(k), uintptr(mt)))
}

func maskLoadunpackloPd(src [8]float64, k uint8, mt uintptr) [8]float64


// LoadunpackloPs: Loads the low-64-byte-aligned portion of the doubleword
// stream starting at element-aligned address mt and expanded into packed
// single-precision (32-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted doublewords that
// occur before first 64-byte-aligned address following 'mt' are loaded.
// Elements in the resulting vector that do not map to those doublewords are
// taken from 'src'. 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			tmp := MEM[addr + loadOffset*4]
//			dst[i+31:i] := tmp[i+31:i]
//			loadOffset := loadOffset + 1
//			IF (mt + loadOffset * 4) % 64 == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPS'. Intrinsic: '_mm512_loadunpacklo_ps'.
// Requires KNCNI.
func LoadunpackloPs(src x86.M512, mt uintptr) x86.M512 {
	return x86.M512(loadunpackloPs([16]float32(src), uintptr(mt)))
}

func loadunpackloPs(src [16]float32, mt uintptr) [16]float32


// MaskLoadunpackloPs: Loads the low-64-byte-aligned portion of the doubleword
// stream starting at element-aligned address mt and expanded into packed
// single-precision (32-bit) floating-point elements in 'dst'. The initial
// values of 'dst' are copied from 'src'. Only those converted doublewords that
// occur before first 64-byte-aligned address following 'mt' are loaded.
// Elements in the resulting vector that do not map to those doublewords are
// taken from 'src'. Elements are loaded from memory according to element
// selector 'k' (elements are skipped when the corresponding mask bit is not
// set). 
//
//		dst[511:0] := src[511:0]
//		loadOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				tmp := MEM[addr + loadOffset*4]
//				dst[i+31:i] := tmp[i+31:i]
//				loadOffset := loadOffset + 1
//				IF (mt + loadOffset * 4) % 64 == 0
//					break
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOADUNPACKLPS'. Intrinsic: '_mm512_mask_loadunpacklo_ps'.
// Requires KNCNI.
func MaskLoadunpackloPs(src x86.M512, k x86.Mmask16, mt uintptr) x86.M512 {
	return x86.M512(maskLoadunpackloPs([16]float32(src), uint16(k), uintptr(mt)))
}

func maskLoadunpackloPs(src [16]float32, k uint16, mt uintptr) [16]float32


// Log2Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := log2(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_log2_ps'.
// Requires KNCNI.
func Log2Ps(a x86.M512) x86.M512 {
	return x86.M512(log2Ps([16]float32(a)))
}

func log2Ps(a [16]float32) [16]float32


// MaskLog2Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := log2(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_mask_log2_ps'.
// Requires KNCNI.
func MaskLog2Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(maskLog2Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskLog2Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// Log2ae23Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a' with absolute error of 2^(-23) and store the
// results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Log2ae23(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_log2ae23_ps'.
// Requires KNCNI.
func Log2ae23Ps(a x86.M512) x86.M512 {
	return x86.M512(log2ae23Ps([16]float32(a)))
}

func log2ae23Ps(a [16]float32) [16]float32


// MaskLog2ae23Ps: Compute the base-2 logarithm of packed single-precision
// (32-bit) floating-point elements in 'a' with absolute error of 2^(-23) and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := Log2ae23(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_mask_log2ae23_ps'.
// Requires KNCNI.
func MaskLog2ae23Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(maskLog2ae23Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskLog2ae23Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// Mask2int: Converts bit mask 'k1' into an integer value, storing the results
// in 'dst'. 
//
//		dst := SignExtend(k1)
//
// Instruction: 'KMOV'. Intrinsic: '_mm512_mask2int'.
// Requires KNCNI.
func Mask2int(k1 x86.Mmask16) int {
	return int(mask2int(uint16(k1)))
}

func mask2int(k1 uint16) int


// MaskMaxEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] > b[i+31:i]
//					dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSD'. Intrinsic: '_mm512_mask_max_epi32'.
// Requires KNCNI.
func MaskMaxEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMaxEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMaxEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MaxEpi32: Compare packed 32-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSD'. Intrinsic: '_mm512_max_epi32'.
// Requires KNCNI.
func MaxEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maxEpi32([64]byte(a), [64]byte(b)))
}

func maxEpi32(a [64]byte, b [64]byte) [64]byte


// MaskMaxEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] > b[i+31:i]
//					dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUD'. Intrinsic: '_mm512_mask_max_epu32'.
// Requires KNCNI.
func MaskMaxEpu32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMaxEpu32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMaxEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MaxEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUD'. Intrinsic: '_mm512_max_epu32'.
// Requires KNCNI.
func MaxEpu32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maxEpu32([64]byte(a), [64]byte(b)))
}

func maxEpu32(a [64]byte, b [64]byte) [64]byte


// MaskMaxabsPs: Determines the maximum of the absolute elements of each pair
// of corresponding elements of packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_mask_maxabs_ps'.
// Requires KNCNI.
func MaskMaxabsPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskMaxabsPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskMaxabsPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// MaxabsPs: Determines the maximum of the absolute elements of each pair of
// corresponding elements of packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_maxabs_ps'.
// Requires KNCNI.
func MaxabsPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maxabsPs([16]float32(a), [16]float32(b)))
}

func maxabsPs(a [16]float32, b [16]float32) [16]float32


// MaskMinEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] < b[i+31:i]
//						dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSD'. Intrinsic: '_mm512_mask_min_epi32'.
// Requires KNCNI.
func MaskMinEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMinEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMinEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MinEpi32: Compare packed 32-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSD'. Intrinsic: '_mm512_min_epi32'.
// Requires KNCNI.
func MinEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(minEpi32([64]byte(a), [64]byte(b)))
}

func minEpi32(a [64]byte, b [64]byte) [64]byte


// MaskMinEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] < b[i+31:i]
//					dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUD'. Intrinsic: '_mm512_mask_min_epu32'.
// Requires KNCNI.
func MaskMinEpu32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMinEpu32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMinEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MinEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUD'. Intrinsic: '_mm512_min_epu32'.
// Requires KNCNI.
func MinEpu32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(minEpu32([64]byte(a), [64]byte(b)))
}

func minEpu32(a [64]byte, b [64]byte) [64]byte


// MaskMovEpi32: Move packed 32-bit integers from 'a' to 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_mask_mov_epi32'.
// Requires KNCNI.
func MaskMovEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i) x86.M512i {
	return x86.M512i(maskMovEpi32([64]byte(src), uint16(k), [64]byte(a)))
}

func maskMovEpi32(src [64]byte, k uint16, a [64]byte) [64]byte


// MaskMovEpi64: Move packed 64-bit integers from 'a' to 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_mask_mov_epi64'.
// Requires KNCNI.
func MaskMovEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i) x86.M512i {
	return x86.M512i(maskMovEpi64([64]byte(src), uint8(k), [64]byte(a)))
}

func maskMovEpi64(src [64]byte, k uint8, a [64]byte) [64]byte


// MaskMovPd: Move packed double-precision (64-bit) floating-point elements
// from 'a' to 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_mask_mov_pd'.
// Requires KNCNI.
func MaskMovPd(src x86.M512d, k x86.Mmask8, a x86.M512d) x86.M512d {
	return x86.M512d(maskMovPd([8]float64(src), uint8(k), [8]float64(a)))
}

func maskMovPd(src [8]float64, k uint8, a [8]float64) [8]float64


// MaskMovPs: Move packed single-precision (32-bit) floating-point elements
// from 'a' to 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_mask_mov_ps'.
// Requires KNCNI.
func MaskMovPs(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(maskMovPs([16]float32(src), uint16(k), [16]float32(a)))
}

func maskMovPs(src [16]float32, k uint16, a [16]float32) [16]float32


// MaskMulPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).  RM. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] * b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mask_mul_pd'.
// Requires KNCNI.
func MaskMulPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(maskMulPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskMulPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// MulPd: Multiply packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mul_pd'.
// Requires KNCNI.
func MulPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(mulPd([8]float64(a), [8]float64(b)))
}

func mulPd(a [8]float64, b [8]float64) [8]float64


// MaskMulPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).  RM. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mask_mul_ps'.
// Requires KNCNI.
func MaskMulPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskMulPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskMulPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// MulPs: Multiply packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mul_ps'.
// Requires KNCNI.
func MulPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(mulPs([16]float32(a), [16]float32(b)))
}

func mulPs(a [16]float32, b [16]float32) [16]float32


// MaskMulRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
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
//				dst[i+63:i] := a[i+63:i] * b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mask_mul_round_pd'.
// Requires KNCNI.
func MaskMulRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskMulRoundPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), rounding))
}

func maskMulRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64


// MulRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'. 
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
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mul_round_pd'.
// Requires KNCNI.
func MulRoundPd(a x86.M512d, b x86.M512d, rounding int) x86.M512d {
	return x86.M512d(mulRoundPd([8]float64(a), [8]float64(b), rounding))
}

func mulRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64


// MaskMulRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	 Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mask_mul_round_ps'.
// Requires KNCNI.
func MaskMulRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(maskMulRoundPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), rounding))
}

func maskMulRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32


// MulRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mul_round_ps'.
// Requires KNCNI.
func MulRoundPs(a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(mulRoundPs([16]float32(a), [16]float32(b), rounding))
}

func mulRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32


// MaskMulhiEpi32: Performs element-by-element multiplication between packed
// 32-bit integer elements in 'a' and 'b' and stores the high 32 bits of each
// result into 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHD'. Intrinsic: '_mm512_mask_mulhi_epi32'.
// Requires KNCNI.
func MaskMulhiEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMulhiEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMulhiEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MulhiEpi32: Performs element-by-element multiplication between packed 32-bit
// integer elements in 'a' and 'b' and stores the high 32 bits of each result
// into 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHD'. Intrinsic: '_mm512_mulhi_epi32'.
// Requires KNCNI.
func MulhiEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(mulhiEpi32([64]byte(a), [64]byte(b)))
}

func mulhiEpi32(a [64]byte, b [64]byte) [64]byte


// MaskMulhiEpu32: Performs element-by-element multiplication between packed
// unsigned 32-bit integer elements in 'a' and 'b' and stores the high 32 bits
// of each result into 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUD'. Intrinsic: '_mm512_mask_mulhi_epu32'.
// Requires KNCNI.
func MaskMulhiEpu32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMulhiEpu32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMulhiEpu32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MulhiEpu32: Performs element-by-element multiplication between packed
// unsigned 32-bit integer elements in 'a' and 'b' and stores the high 32 bits
// of each result into 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUD'. Intrinsic: '_mm512_mulhi_epu32'.
// Requires KNCNI.
func MulhiEpu32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(mulhiEpu32([64]byte(a), [64]byte(b)))
}

func mulhiEpu32(a [64]byte, b [64]byte) [64]byte


// MaskMulloEpi32: Multiply the packed 32-bit integers in 'a' and 'b',
// producing intermediate 64-bit integers, and store the low 32 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				tmp[63:0] := a[i+31:i] * b[i+31:i]
//				dst[i+31:i] := tmp[31:0]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLD'. Intrinsic: '_mm512_mask_mullo_epi32'.
// Requires KNCNI.
func MaskMulloEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskMulloEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskMulloEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// MulloEpi32: Multiply the packed 32-bit integers in 'a' and 'b', producing
// intermediate 64-bit integers, and store the low 32 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			tmp[63:0] := a[i+31:i] * b[i+31:i]
//			dst[i+31:i] := tmp[31:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLD'. Intrinsic: '_mm512_mullo_epi32'.
// Requires KNCNI.
func MulloEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(mulloEpi32([64]byte(a), [64]byte(b)))
}

func mulloEpi32(a [64]byte, b [64]byte) [64]byte


// MaskOrEpi32: Compute the bitwise OR of packed 32-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORD'. Intrinsic: '_mm512_mask_or_epi32'.
// Requires KNCNI.
func MaskOrEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskOrEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskOrEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// OrEpi32: Compute the bitwise OR of packed 32-bit integers in 'a' and 'b',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] OR b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORD'. Intrinsic: '_mm512_or_epi32'.
// Requires KNCNI.
func OrEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(orEpi32([64]byte(a), [64]byte(b)))
}

func orEpi32(a [64]byte, b [64]byte) [64]byte


// MaskOrEpi64: Compute the bitwise OR of packed 64-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORQ'. Intrinsic: '_mm512_mask_or_epi64'.
// Requires KNCNI.
func MaskOrEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskOrEpi64([64]byte(src), uint8(k), [64]byte(a), [64]byte(b)))
}

func maskOrEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte


// OrEpi64: Compute the bitwise OR of packed 64-bit integers in 'a' and 'b',
// and store the resut in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] OR b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORQ'. Intrinsic: '_mm512_or_epi64'.
// Requires KNCNI.
func OrEpi64(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(orEpi64([64]byte(a), [64]byte(b)))
}

func orEpi64(a [64]byte, b [64]byte) [64]byte


// OrSi512: Compute the bitwise OR of 512 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := (a[511:0] OR b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPORD'. Intrinsic: '_mm512_or_si512'.
// Requires KNCNI.
func OrSi512(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(orSi512([64]byte(a), [64]byte(b)))
}

func orSi512(a [64]byte, b [64]byte) [64]byte


// MaskPackstorehiEpi32: Stores packed 32-bit integer elements of 'v1' into a
// doubleword stream at a logically mapped starting address (mt-64), storing
// the high-64-byte elements of that stream (those elements of the stream that
// map at or after the first 64-byte-aligned address following (m5-64)).
// Elements are loaded from memory according to element selector 'k' (elements
// are skipped when the corresponding mask bit is not set). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*4) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*32
//					MEM[addr + storeOffset*4] := v1[i+31:i]
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHD'. Intrinsic: '_mm512_mask_packstorehi_epi32'.
// Requires KNCNI.
func MaskPackstorehiEpi32(mt uintptr, k x86.Mmask16, v1 x86.M512i)  {
	maskPackstorehiEpi32(uintptr(mt), uint16(k), [64]byte(v1))
}

func maskPackstorehiEpi32(mt uintptr, k uint16, v1 [64]byte) 


// PackstorehiEpi32: Stores packed 32-bit integer elements of 'v1' into a
// doubleword stream at a logically mapped starting address (mt-64), storing
// the high-64-byte elements of that stream (those elements of the stream that
// map at or after the first 64-byte-aligned address following (m5-64)). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*4) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*32
//				MEM[addr + storeOffset*4] := v1[i+31:i]
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHD'. Intrinsic: '_mm512_packstorehi_epi32'.
// Requires KNCNI.
func PackstorehiEpi32(mt uintptr, v1 x86.M512i)  {
	packstorehiEpi32(uintptr(mt), [64]byte(v1))
}

func packstorehiEpi32(mt uintptr, v1 [64]byte) 


// MaskPackstorehiEpi64: Stores packed 64-bit integer elements of 'v1' into a
// quadword stream at a logically mapped starting address (mt-64), storing the
// high-64-byte elements of that stream (those elemetns of the stream that map
// at or after the first 64-byte-aligned address following (m5-64)). Elements
// are loaded from memory according to element selector 'k' (elements are
// skipped when the corresponding mask bit is not set). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*8) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*64
//					MEM[addr + storeOffset*8] := v1[i+63:i]
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHQ'. Intrinsic: '_mm512_mask_packstorehi_epi64'.
// Requires KNCNI.
func MaskPackstorehiEpi64(mt uintptr, k x86.Mmask8, v1 x86.M512i)  {
	maskPackstorehiEpi64(uintptr(mt), uint8(k), [64]byte(v1))
}

func maskPackstorehiEpi64(mt uintptr, k uint8, v1 [64]byte) 


// PackstorehiEpi64: Stores packed 64-bit integer elements of 'v1' into a
// quadword stream at a logically mapped starting address (mt-64), storing the
// high-64-byte elements of that stream (those elemetns of the stream that map
// at or after the first 64-byte-aligned address following (m5-64)). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*8) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*64
//				MEM[addr + storeOffset*8] := v1[i+63:i]
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHQ'. Intrinsic: '_mm512_packstorehi_epi64'.
// Requires KNCNI.
func PackstorehiEpi64(mt uintptr, v1 x86.M512i)  {
	packstorehiEpi64(uintptr(mt), [64]byte(v1))
}

func packstorehiEpi64(mt uintptr, v1 [64]byte) 


// MaskPackstorehiPd: Stores packed double-precision (64-bit) floating-point
// elements of 'v1' into a quadword stream at a logically mapped starting
// address (mt-64), storing the high-64-byte elements of that stream (those
// elemetns of the stream that map at or after the first 64-byte-aligned
// address following (m5-64)). Elements are loaded from memory according to
// element selector 'k' (elements are skipped when the corresponding mask bit
// is not set). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*8) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*64
//					MEM[addr + storeOffset*4] := v1[i+63:i]
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPD'. Intrinsic: '_mm512_mask_packstorehi_pd'.
// Requires KNCNI.
func MaskPackstorehiPd(mt uintptr, k x86.Mmask8, v1 x86.M512d)  {
	maskPackstorehiPd(uintptr(mt), uint8(k), [8]float64(v1))
}

func maskPackstorehiPd(mt uintptr, k uint8, v1 [8]float64) 


// PackstorehiPd: Stores packed double-precision (64-bit) floating-point
// elements of 'v1' into a quadword stream at a logically mapped starting
// address (mt-64), storing the high-64-byte elements of that stream (those
// elemetns of the stream that map at or after the first 64-byte-aligned
// address following (m5-64)). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 7
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*8) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*64
//				MEM[addr + storeOffset*4] := v1[i+63:i]
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPD'. Intrinsic: '_mm512_packstorehi_pd'.
// Requires KNCNI.
func PackstorehiPd(mt uintptr, v1 x86.M512d)  {
	packstorehiPd(uintptr(mt), [8]float64(v1))
}

func packstorehiPd(mt uintptr, v1 [8]float64) 


// MaskPackstorehiPs: Stores packed single-precision (32-bit) floating-point
// elements of 'v1' into a doubleword stream at a logically mapped starting
// address (mt-64), storing the high-64-byte elements of that stream (those
// elemetns of the stream that map at or after the first 64-byte-aligned
// address following (m5-64)). Elements are loaded from memory according to
// element selector 'k' (elements are skipped when the corresponding mask bit
// is not set). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF k[j]
//				IF foundNext64BytesBoundary == false
//					IF ((addr + (storeOffset + 1)*4) % 64) == 0
//						foundNext64BytesBoundary = true
//					FI
//				ELSE
//					i := j*32
//					MEM[addr + storeOffset*4] := v1[i+31:i]
//				FI
//				storeOffset := storeOffset + 1
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPS'. Intrinsic: '_mm512_mask_packstorehi_ps'.
// Requires KNCNI.
func MaskPackstorehiPs(mt uintptr, k x86.Mmask16, v1 x86.M512)  {
	maskPackstorehiPs(uintptr(mt), uint16(k), [16]float32(v1))
}

func maskPackstorehiPs(mt uintptr, k uint16, v1 [16]float32) 


// PackstorehiPs: Stores packed single-precision (32-bit) floating-point
// elements of 'v1' into a doubleword stream at a logically mapped starting
// address (mt-64), storing the high-64-byte elements of that stream (those
// elemetns of the stream that map at or after the first 64-byte-aligned
// address following (m5-64)). 
//
//		storeOffset := 0
//		foundNext64BytesBoundary := false
//		addr = mt-64
//		FOR j := 0 to 15
//			IF foundNext64BytesBoundary == false
//				IF ((addr + (storeOffset + 1)*4) % 64) == 0
//					foundNext64BytesBoundary = true
//				FI
//			ELSE
//				i := j*32
//				MEM[addr + storeOffset*4] := v1[i+31:i]
//			FI
//			storeOffset := storeOffset + 1
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTOREHPS'. Intrinsic: '_mm512_packstorehi_ps'.
// Requires KNCNI.
func PackstorehiPs(mt uintptr, v1 x86.M512)  {
	packstorehiPs(uintptr(mt), [16]float32(v1))
}

func packstorehiPs(mt uintptr, v1 [16]float32) 


// MaskPackstoreloEpi32: Stores packed 32-bit integer elements of 'v1' into a
// doubleword stream at a logically mapped starting address 'mt', storing the
// low-64-byte elements of that stream (those elements of the stream that map
// before the first 64-byte-aligned address follwing 'mt'). Elements are loaded
// from memory according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			IF k[j]
//				i := j*32
//				MEM[addr + storeOffset*4] := v1[i+31:i]
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset*4) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELD'. Intrinsic: '_mm512_mask_packstorelo_epi32'.
// Requires KNCNI.
func MaskPackstoreloEpi32(mt uintptr, k x86.Mmask16, v1 x86.M512i)  {
	maskPackstoreloEpi32(uintptr(mt), uint16(k), [64]byte(v1))
}

func maskPackstoreloEpi32(mt uintptr, k uint16, v1 [64]byte) 


// PackstoreloEpi32: Stores packed 32-bit integer elements of 'v1' into a
// doubleword stream at a logically mapped starting address 'mt', storing the
// low-64-byte elements of that stream (those elements of the stream that map
// before the first 64-byte-aligned address follwing 'mt'). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			MEM[addr + storeOffset*4] := v1[i+31:i]
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset*4) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELD'. Intrinsic: '_mm512_packstorelo_epi32'.
// Requires KNCNI.
func PackstoreloEpi32(mt uintptr, v1 x86.M512i)  {
	packstoreloEpi32(uintptr(mt), [64]byte(v1))
}

func packstoreloEpi32(mt uintptr, v1 [64]byte) 


// MaskPackstoreloEpi64: Stores packed 64-bit integer elements of 'v1' into a
// quadword stream at a logically mapped starting address 'mt', storing the
// low-64-byte elements of that stream (those elements of the stream that map
// before the first 64-byte-aligned address follwing 'mt'). Elements are loaded
// from memory according to element selector 'k' (elements are skipped when the
// corresponding mask bit is not set). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			IF k[j]
//				i := j*64
//				MEM[addr + storeOffset*8] := v1[i+63:i]
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset*8) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELQ'. Intrinsic: '_mm512_mask_packstorelo_epi64'.
// Requires KNCNI.
func MaskPackstoreloEpi64(mt uintptr, k x86.Mmask8, v1 x86.M512i)  {
	maskPackstoreloEpi64(uintptr(mt), uint8(k), [64]byte(v1))
}

func maskPackstoreloEpi64(mt uintptr, k uint8, v1 [64]byte) 


// PackstoreloEpi64: Stores packed 64-bit integer elements of 'v1' into a
// quadword stream at a logically mapped starting address 'mt', storing the
// low-64-byte elements of that stream (those elements of the stream that map
// before the first 64-byte-aligned address follwing 'mt'). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			MEM[addr + storeOffset*8] := v1[i+63:i]
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset*8) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELQ'. Intrinsic: '_mm512_packstorelo_epi64'.
// Requires KNCNI.
func PackstoreloEpi64(mt uintptr, v1 x86.M512i)  {
	packstoreloEpi64(uintptr(mt), [64]byte(v1))
}

func packstoreloEpi64(mt uintptr, v1 [64]byte) 


// MaskPackstoreloPd: Stores packed double-precision (64-bit) floating-point
// elements of 'v1' into a quadword stream at a logically mapped starting
// address 'mt', storing the low-64-byte elements of that stream (those
// elements of the stream that map before the first 64-byte-aligned address
// follwing 'mt'). Elements are loaded from memory according to element
// selector 'k' (elements are skipped when the corresponding mask bit is not
// set). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			IF k[j]
//				i := j*64
//				MEM[addr + storeOffset*8] := v1[i+63:i]
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset*8) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPD'. Intrinsic: '_mm512_mask_packstorelo_pd'.
// Requires KNCNI.
func MaskPackstoreloPd(mt uintptr, k x86.Mmask8, v1 x86.M512d)  {
	maskPackstoreloPd(uintptr(mt), uint8(k), [8]float64(v1))
}

func maskPackstoreloPd(mt uintptr, k uint8, v1 [8]float64) 


// PackstoreloPd: Stores packed double-precision (64-bit) floating-point
// elements of 'v1' into a quadword stream at a logically mapped starting
// address 'mt', storing the low-64-byte elements of that stream (those
// elements of the stream that map before the first 64-byte-aligned address
// follwing 'mt'). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 7
//			i := j*64
//			MEM[addr + storeOffset*8] := v1[i+63:i]
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset*8) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPD'. Intrinsic: '_mm512_packstorelo_pd'.
// Requires KNCNI.
func PackstoreloPd(mt uintptr, v1 x86.M512d)  {
	packstoreloPd(uintptr(mt), [8]float64(v1))
}

func packstoreloPd(mt uintptr, v1 [8]float64) 


// MaskPackstoreloPs: Stores packed single-precision (32-bit) floating-point
// elements of 'v1' into a doubleword stream at a logically mapped starting
// address 'mt', storing the low-64-byte elements of that stream (those
// elements of the stream that map before the first 64-byte-aligned address
// follwing 'mt'). Elements are loaded from memory according to element
// selector 'k' (elements are skipped when the corresponding mask bit is not
// set). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			IF k[j]
//				i := j*32
//				MEM[addr + storeOffset*4] := v1[i+31:i]
//				storeOffset := storeOffset + 1
//				IF ((addr + storeOffset*4) % 64) == 0
//					BREAK
//				FI
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPS'. Intrinsic: '_mm512_mask_packstorelo_ps'.
// Requires KNCNI.
func MaskPackstoreloPs(mt uintptr, k x86.Mmask16, v1 x86.M512)  {
	maskPackstoreloPs(uintptr(mt), uint16(k), [16]float32(v1))
}

func maskPackstoreloPs(mt uintptr, k uint16, v1 [16]float32) 


// PackstoreloPs: Stores packed single-precision (32-bit) floating-point
// elements of 'v1' into a doubleword stream at a logically mapped starting
// address 'mt', storing the low-64-byte elements of that stream (those
// elements of the stream that map before the first 64-byte-aligned address
// follwing 'mt'). 
//
//		storeOffset := 0
//		addr = mt
//		FOR j := 0 to 15
//			i := j*32
//			MEM[addr + storeOffset*4] := v1[i+31:i]
//			storeOffset := storeOffset + 1
//			IF ((addr + storeOffset*4) % 64) == 0
//				BREAK
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPACKSTORELPS'. Intrinsic: '_mm512_packstorelo_ps'.
// Requires KNCNI.
func PackstoreloPs(mt uintptr, v1 x86.M512)  {
	packstoreloPs(uintptr(mt), [16]float32(v1))
}

func packstoreloPs(mt uintptr, v1 [16]float32) 


// MaskPermute4f128Epi32: Permutes 128-bit blocks of the packed 32-bit integer
// vector 'a' using constant 'imm8'. The results are stored in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		tmp[511:0] := 0
//		FOR j := 0 to 4
//			i := j*128
//			n := j*2
//			tmp[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		FOR j := 0 to 15
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_mask_permute4f128_epi32'.
// Requires KNCNI.
func MaskPermute4f128Epi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 MMPERMENUM) x86.M512i {
	return x86.M512i(maskPermute4f128Epi32([64]byte(src), uint16(k), [64]byte(a), imm8))
}

func maskPermute4f128Epi32(src [64]byte, k uint16, a [64]byte, imm8 MMPERMENUM) [64]byte


// Permute4f128Epi32: Permutes 128-bit blocks of the packed 32-bit integer
// vector 'a' using constant 'imm8'. The results are stored in 'dst'. 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*128
//			n := j*2
//			dst[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_permute4f128_epi32'.
// Requires KNCNI.
func Permute4f128Epi32(a x86.M512i, imm8 MMPERMENUM) x86.M512i {
	return x86.M512i(permute4f128Epi32([64]byte(a), imm8))
}

func permute4f128Epi32(a [64]byte, imm8 MMPERMENUM) [64]byte


// MaskPermute4f128Ps: Permutes 128-bit blocks of the packed single-precision
// (32-bit) floating-point elements in 'a' using constant 'imm8'. The results
// are stored in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		tmp[511:0] := 0
//		FOR j := 0 to 4
//			i := j*128
//			n := j*2
//			tmp[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		FOR j := 0 to 15
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_mask_permute4f128_ps'.
// Requires KNCNI.
func MaskPermute4f128Ps(src x86.M512, k x86.Mmask16, a x86.M512, imm8 MMPERMENUM) x86.M512 {
	return x86.M512(maskPermute4f128Ps([16]float32(src), uint16(k), [16]float32(a), imm8))
}

func maskPermute4f128Ps(src [16]float32, k uint16, a [16]float32, imm8 MMPERMENUM) [16]float32


// Permute4f128Ps: Permutes 128-bit blocks of the packed single-precision
// (32-bit) floating-point elements in 'a' using constant 'imm8'. The results
// are stored in 'dst'. 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*128
//			n := j*2
//			dst[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_permute4f128_ps'.
// Requires KNCNI.
func Permute4f128Ps(a x86.M512, imm8 MMPERMENUM) x86.M512 {
	return x86.M512(permute4f128Ps([16]float32(a), imm8))
}

func permute4f128Ps(a [16]float32, imm8 MMPERMENUM) [16]float32


// MaskPermutevarEpi32: Shuffle 32-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Note that this intrinsic shuffles across 128-bit lanes, unlike past
// intrinsics that use the 'permutevar' name. This intrinsic is identical to
// '_mm512_mask_permutexvar_epi32', and it is recommended that you use that
// intrinsic name. 
//
//		FOR j := 0 to 15
//			i := j*32
//			id := idx[i+3:i]*32
//			IF k[j]
//				dst[i+31:i] := a[id+31:id]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMD'. Intrinsic: '_mm512_mask_permutevar_epi32'.
// Requires KNCNI.
func MaskPermutevarEpi32(src x86.M512i, k x86.Mmask16, idx x86.M512i, a x86.M512i) x86.M512i {
	return x86.M512i(maskPermutevarEpi32([64]byte(src), uint16(k), [64]byte(idx), [64]byte(a)))
}

func maskPermutevarEpi32(src [64]byte, k uint16, idx [64]byte, a [64]byte) [64]byte


// PermutevarEpi32: Shuffle 32-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst'. Note that this
// intrinsic shuffles across 128-bit lanes, unlike past intrinsics that use the
// 'permutevar' name. This intrinsic is identical to
// '_mm512_permutexvar_epi32', and it is recommended that you use that
// intrinsic name. 
//
//		FOR j := 0 to 15
//			i := j*32
//			id := idx[i+3:i]*32
//			dst[i+31:i] := a[id+31:id]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMD'. Intrinsic: '_mm512_permutevar_epi32'.
// Requires KNCNI.
func PermutevarEpi32(idx x86.M512i, a x86.M512i) x86.M512i {
	return x86.M512i(permutevarEpi32([64]byte(idx), [64]byte(a)))
}

func permutevarEpi32(idx [64]byte, a [64]byte) [64]byte


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'VPREFETCH0, VPREFETCH1, VPREFETCH2, VPREFETCHNTA, VPREFETCHE0, VPREFETCHE1, VPREFETCHE2, VPREFETCHENTA'. Intrinsic: '_mm_prefetch'.
// Requires KNCNI.
func Prefetch(p byte, i int)  {
	prefetch(p, i)
}

func prefetch(p byte, i int) 


// MaskPrefetchI32extgatherPs: Prefetches a set of 16 single-precision (32-bit)
// memory locations pointed by base address 'mv' and 32-bit integer index
// vector 'index' with scale 'scale' to L1 or L2 level of cache depending on
// the value of 'hint'. Gathered elements are merged in cache using writemask
// 'k' (elements are brought into cache only when their corresponding mask bits
// are set). The 'hint' parameter may be 1 (_MM_HINT_T0) for prefetching to L1
// cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache.
// The 'conv' parameter specifies the granularity used by compilers to better
// encode the instruction. It should be the same as the 'conv' parameter
// specified for the subsequent gather intrinsic. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			IF k[j] THEN
//				CASE hint OF
//				_MM_HINT_T0: PrefetchL1WithT0Hint(addr[i+31:i])
//				_MM_HINT_T1: PrefetchL2WithT1Hint(addr[i+31:i])
//				ESAC
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERPF0DPS, VGATHERPF1DPS'. Intrinsic: '_mm512_mask_prefetch_i32extgather_ps'.
// Requires KNCNI.
func MaskPrefetchI32extgatherPs(index x86.M512i, k x86.Mmask16, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int)  {
	maskPrefetchI32extgatherPs([64]byte(index), uint16(k), uintptr(mv), conv, scale, hint)
}

func maskPrefetchI32extgatherPs(index [64]byte, k uint16, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) 


// PrefetchI32extgatherPs: Prefetches a set of 16 single-precision (32-bit)
// memory locations pointed by base address 'mv' and 32-bit integer index
// vector 'index' with scale 'scale' to L1 or L2 level of cache depending on
// the value of 'hint'. The 'hint' parameter may be 1 (_MM_HINT_T0) for
// prefetching to L1 cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache.
// The 'conv' parameter specifies the granularity used by compilers to better
// encode the instruction. It should be the same as the 'conv' parameter
// specified for the subsequent gather intrinsic. 
//
//		FOR j := 0 to 15
//			addr := MEM[mv + index[j] * scale]
//			i := j*32
//			CASE hint OF
//			_MM_HINT_T0: PrefetchL1WithT0Hint(addr[i+31:i])
//			_MM_HINT_T1: PrefetchL2WithT1Hint(addr[i+31:i])
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGATHERPF0DPS, VGATHERPF1DPS'. Intrinsic: '_mm512_prefetch_i32extgather_ps'.
// Requires KNCNI.
func PrefetchI32extgatherPs(index x86.M512i, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int)  {
	prefetchI32extgatherPs([64]byte(index), uintptr(mv), conv, scale, hint)
}

func prefetchI32extgatherPs(index [64]byte, mv uintptr, conv MMUPCONVPSENUM, scale int, hint int) 


// MaskPrefetchI32extscatterPs: Prefetches a set of 16 single-precision
// (32-bit) memory locations pointed by base address 'mv' and 32-bit integer
// index vector 'index' with scale 'scale' to L1 or L2 level of cache depending
// on the value of 'hint'. The 'hint' parameter may be 1 (_MM_HINT_T0) for
// prefetching to L1 cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache.
// The 'conv' parameter specifies the granularity used by compilers to better
// encode the instruction. It should be the same as the 'conv' parameter
// specified for the subsequent gather intrinsic. Only those elements whose
// corresponding mask bit in 'k' is set are loaded into cache. 
//
//		cachev := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				CASE hint OF
//				_MM_HINT_T0: PrefetchL1WithT0Hint(addr[i+31:i])
//				_MM_HINT_T1: PrefetchL2WithT1Hint(addr[i+31:i])
//				ESAC
//			FI
//		ENDFOR
//
// Instruction: 'VSCATTERPF0DPS, VSCATTERPF1DPS'. Intrinsic: '_mm512_mask_prefetch_i32extscatter_ps'.
// Requires KNCNI.
func MaskPrefetchI32extscatterPs(mv uintptr, k x86.Mmask16, index x86.M512i, conv MMUPCONVPSENUM, scale int, hint int)  {
	maskPrefetchI32extscatterPs(uintptr(mv), uint16(k), [64]byte(index), conv, scale, hint)
}

func maskPrefetchI32extscatterPs(mv uintptr, k uint16, index [64]byte, conv MMUPCONVPSENUM, scale int, hint int) 


// PrefetchI32extscatterPs: Prefetches a set of 16 single-precision (32-bit)
// memory locations pointed by base address 'mv' and 32-bit integer index
// vector 'index' with scale 'scale' to L1 or L2 level of cache depending on
// the value of 'hint', with a request for exclusive ownership. The 'hint'
// parameter may be one of the following: _MM_HINT_T0 = 1 for prefetching to L1
// cache, _MM_HINT_T1 = 2 for prefetching to L2 cache, _MM_HINT_T2 = 3 for
// prefetching to L2 cache non-temporal, _MM_HINT_NTA = 0 for prefetching to L1
// cache non-temporal. The 'conv' parameter specifies the granularity used by
// compilers to better encode the instruction. It should be the same as the
// 'conv' parameter specified for the subsequent scatter intrinsic. 
//
//		cachev := 0
//		FOR j := 0 to 15
//			i := j*32
//			addr := MEM[mv + index[j] * scale]
//			CASE hint OF
//			_MM_HINT_T0: PrefetchL1WithT0Hint(addr[i+31:i])
//			_MM_HINT_T1: PrefetchL2WithT1Hint(addr[i+31:i])
//			_MM_HINT_T2: PrefetchL2WithT1HintNonTemporal(addr[i+31:i])
//			_MM_HINT_NTA: PrefetchL1WithT0HintNonTemporal(addr[i+31:i])
//			ESAC
//		ENDFOR
//
// Instruction: 'VSCATTERPF0DPS, VSCATTERPF1DPS'. Intrinsic: '_mm512_prefetch_i32extscatter_ps'.
// Requires KNCNI.
func PrefetchI32extscatterPs(mv uintptr, index x86.M512i, conv MMUPCONVPSENUM, scale int, hint int)  {
	prefetchI32extscatterPs(uintptr(mv), [64]byte(index), conv, scale, hint)
}

func prefetchI32extscatterPs(mv uintptr, index [64]byte, conv MMUPCONVPSENUM, scale int, hint int) 


// MaskPrefetchI32gatherPs: Prefetch single-precision (32-bit) floating-point
// elements from memory using 32-bit indices. 32-bit elements are loaded from
// addresses starting at 'base_addr' and offset by each 32-bit element in
// 'vindex' (each index is scaled by the factor in 'scale'). Gathered elements
// are merged in cache using writemask 'k' (elements are brought into cache
// only when their corresponding mask bits are set). 'scale' should be 1, 2, 4
// or 8. The 'hint' parameter may be 1 (_MM_HINT_T0) for prefetching to L1
// cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache. 
//
//		FOR j := 0 to 15
//			i := j*16;
//			IF mask[j] THEN
//				Prefetch([base_addr + SignExtend(vindex[i*31:i]) * scale], hint, RFO=0);
//			FI
//		ENDFOR;
//
// Instruction: 'VGATHERPF0DPS, VGATHERPF1DPS'. Intrinsic: '_mm512_mask_prefetch_i32gather_ps'.
// Requires KNCNI.
func MaskPrefetchI32gatherPs(vindex x86.M512i, mask x86.Mmask16, base_addr uintptr, scale int, hint int)  {
	maskPrefetchI32gatherPs([64]byte(vindex), uint16(mask), uintptr(base_addr), scale, hint)
}

func maskPrefetchI32gatherPs(vindex [64]byte, mask uint16, base_addr uintptr, scale int, hint int) 


// PrefetchI32gatherPs: Prefetches 16 single-precision (32-bit) floating-point
// elements in memory starting at location 'mv' at packed 32-bit integer
// indices stored in 'index' scaled by 'scale'. The 'hint' parameter may be 1
// (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 
//
//		cachev := 0
//		FOR j := 0 to 15
//			i := j*32
//			addr := MEM[mv + index[j] * scale]
//			cachev[i+31:i] := addr[i+63:i]
//		ENDFOR
//
// Instruction: 'VGATHERPF0DPS, VGATHERPF1DPS'. Intrinsic: '_mm512_prefetch_i32gather_ps'.
// Requires KNCNI.
func PrefetchI32gatherPs(index x86.M512i, mv uintptr, scale int, hint int)  {
	prefetchI32gatherPs([64]byte(index), uintptr(mv), scale, hint)
}

func prefetchI32gatherPs(index [64]byte, mv uintptr, scale int, hint int) 


// MaskPrefetchI32scatterPs: Prefetches 16 single-precision (32-bit)
// floating-point elements in memory starting at location 'mv' at packed 32-bit
// integer indices stored in 'index' scaled by 'scale'. The 'hint' parameter
// may be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. Only those elements whose corresponding mask bit in
// 'k' is set are loaded into the desired cache. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				addr := MEM[mv + index[j] * scale]
//				CASE hint OF
//				_MM_HINT_T0: PrefetchL1WithT0Hint(addr[i+31:i])
//				_MM_HINT_T1: PrefetchL2WithT1Hint(addr[i+31:i])
//				_MM_HINT_T2: PrefetchL2WithT1HintNonTemporal(addr[i+31:i])
//				_MM_HINT_NTA: PrefetchL1WithT0HintNonTemporal(addr[i+31:i])
//			FI
//		ENDFOR
//
// Instruction: 'VSCATTERPF0DPS, VSCATTERPF1DPS'. Intrinsic: '_mm512_mask_prefetch_i32scatter_ps'.
// Requires KNCNI.
func MaskPrefetchI32scatterPs(mv uintptr, k x86.Mmask16, index x86.M512i, scale int, hint int)  {
	maskPrefetchI32scatterPs(uintptr(mv), uint16(k), [64]byte(index), scale, hint)
}

func maskPrefetchI32scatterPs(mv uintptr, k uint16, index [64]byte, scale int, hint int) 


// PrefetchI32scatterPs: Prefetches 16 single-precision (32-bit) floating-point
// elements in memory starting at location 'mv' at packed 32-bit integer
// indices stored in 'index' scaled by 'scale'. The 'hint' parameter may be 1
// (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 
//
//		FOR j := 0 to 15
//			i := j*32
//			addr := MEM[mv + index[j] * scale]
//			CASE hint OF
//			_MM_HINT_T0: PrefetchL1WithT0Hint(addr[i+31:i])
//			_MM_HINT_T1: PrefetchL2WithT1Hint(addr[i+31:i])
//			_MM_HINT_T2: PrefetchL2WithT1HintNonTemporal(addr[i+31:i])
//			_MM_HINT_NTA: PrefetchL1WithT0HintNonTemporal(addr[i+31:i])
//			ESAC
//		ENDFOR
//
// Instruction: 'VSCATTERPF0DPS, VSCATTERPF1DPS'. Intrinsic: '_mm512_prefetch_i32scatter_ps'.
// Requires KNCNI.
func PrefetchI32scatterPs(mv uintptr, index x86.M512i, scale int, hint int)  {
	prefetchI32scatterPs(uintptr(mv), [64]byte(index), scale, hint)
}

func prefetchI32scatterPs(mv uintptr, index [64]byte, scale int, hint int) 


// MaskRcp23Ps: Approximates the reciprocals of packed single-precision
// (32-bit) floating-point elements in 'a' to 23 bits of precision, storing the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := APPROXIMATE(1.0/a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRCP23PS'. Intrinsic: '_mm512_mask_rcp23_ps'.
// Requires KNCNI.
func MaskRcp23Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(maskRcp23Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskRcp23Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// Rcp23Ps: Approximates the reciprocals of packed single-precision (32-bit)
// floating-point elements in 'a' to 23 bits of precision, storing the results
// in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := APPROXIMATE(1.0/a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRCP23PS'. Intrinsic: '_mm512_rcp23_ps'.
// Requires KNCNI.
func Rcp23Ps(a x86.M512) x86.M512 {
	return x86.M512(rcp23Ps([16]float32(a)))
}

func rcp23Ps(a [16]float32) [16]float32


// MaskReduceAddEpi32: Reduce the packed 32-bit integers in 'a' by addition
// using mask 'k'. Returns the sum of all active elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				sum[31:0] := sum[31:0] + a[i+31:i]
//			FI
//		ENDFOR
//		RETURN sum[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_epi32'.
// Requires KNCNI.
func MaskReduceAddEpi32(k x86.Mmask16, a x86.M512i) int {
	return int(maskReduceAddEpi32(uint16(k), [64]byte(a)))
}

func maskReduceAddEpi32(k uint16, a [64]byte) int


// ReduceAddEpi32: Reduce the packed 32-bit integers in 'a' by addition.
// Returns the sum of all elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			sum[31:0] := sum[31:0] + a[i+31:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_epi32'.
// Requires KNCNI.
func ReduceAddEpi32(a x86.M512i) int {
	return int(reduceAddEpi32([64]byte(a)))
}

func reduceAddEpi32(a [64]byte) int


// MaskReduceAddEpi64: Reduce the packed 64-bit integers in 'a' by addition
// using mask 'k'. Returns the sum of all active elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				sum[63:0] := sum[63:0] + a[i+63:i]
//			FI
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_epi64'.
// Requires KNCNI.
func MaskReduceAddEpi64(k x86.Mmask8, a x86.M512i) int64 {
	return int64(maskReduceAddEpi64(uint8(k), [64]byte(a)))
}

func maskReduceAddEpi64(k uint8, a [64]byte) int64


// ReduceAddEpi64: Reduce the packed 64-bit integers in 'a' by addition.
// Returns the sum of all elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			sum[63:0] := sum[63:0] + a[i+63:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_epi64'.
// Requires KNCNI.
func ReduceAddEpi64(a x86.M512i) int64 {
	return int64(reduceAddEpi64([64]byte(a)))
}

func reduceAddEpi64(a [64]byte) int64


// MaskReduceAddPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by addition using mask 'k'. Returns the sum of all active
// elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				sum[63:0] := sum[63:0] + a[i+63:i]
//			FI
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_pd'.
// Requires KNCNI.
func MaskReduceAddPd(k x86.Mmask8, a x86.M512d) float64 {
	return float64(maskReduceAddPd(uint8(k), [8]float64(a)))
}

func maskReduceAddPd(k uint8, a [8]float64) float64


// ReduceAddPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by addition. Returns the sum of all elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			sum[63:0] := sum[63:0] + a[i+63:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_pd'.
// Requires KNCNI.
func ReduceAddPd(a x86.M512d) float64 {
	return float64(reduceAddPd([8]float64(a)))
}

func reduceAddPd(a [8]float64) float64


// MaskReduceAddPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by addition using mask 'k'. Returns the sum of all active
// elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				sum[31:0] := sum[31:0] + a[i+31:i]
//			FI
//		ENDFOR
//		RETURN sum[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_ps'.
// Requires KNCNI.
func MaskReduceAddPs(k x86.Mmask16, a x86.M512) float32 {
	return float32(maskReduceAddPs(uint16(k), [16]float32(a)))
}

func maskReduceAddPs(k uint16, a [16]float32) float32


// ReduceAddPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by addition. Returns the sum of all elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			sum[31:0] := sum[31:0] + a[i+31:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_ps'.
// Requires KNCNI.
func ReduceAddPs(a x86.M512) float32 {
	return float32(reduceAddPs([16]float32(a)))
}

func reduceAddPs(a [16]float32) float32


// MaskReduceAndEpi32: Reduce the packed 32-bit integers in 'a' by bitwise AND
// using mask 'k'. Returns the bitwise AND of all active elements in 'a'. 
//
//		reduced[31:0] := 0xFFFFFFFF
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				reduced[31:0] := reduced[31:0] AND a[i+31:i]
//			FI
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_and_epi32'.
// Requires KNCNI.
func MaskReduceAndEpi32(k x86.Mmask16, a x86.M512i) int {
	return int(maskReduceAndEpi32(uint16(k), [64]byte(a)))
}

func maskReduceAndEpi32(k uint16, a [64]byte) int


// ReduceAndEpi32: Reduce the packed 32-bit integers in 'a' by bitwise AND.
// Returns the bitwise AND of all elements in 'a'. 
//
//		reduced[31:0] := 0xFFFFFFFF
//		FOR j := 0 to 15
//			i := j*32
//			reduced[31:0] := reduced[31:0] AND a[i+31:i]
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_and_epi32'.
// Requires KNCNI.
func ReduceAndEpi32(a x86.M512i) int {
	return int(reduceAndEpi32([64]byte(a)))
}

func reduceAndEpi32(a [64]byte) int


// MaskReduceAndEpi64: Reduce the packed 64-bit integers in 'a' by bitwise AND
// using mask 'k'. Returns the bitwise AND of all active elements in 'a'. 
//
//		reduced[63:0] := 0xFFFFFFFFFFFFFFFF
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				reduced[63:0] := reduced[63:0] AND a[i+63:i]
//			FI
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_and_epi64'.
// Requires KNCNI.
func MaskReduceAndEpi64(k x86.Mmask8, a x86.M512i) int64 {
	return int64(maskReduceAndEpi64(uint8(k), [64]byte(a)))
}

func maskReduceAndEpi64(k uint8, a [64]byte) int64


// ReduceAndEpi64: Reduce the packed 64-bit integers in 'a' by bitwise AND.
// Returns the bitwise AND of all elements in 'a'. 
//
//		reduced[63:0] := 0xFFFFFFFFFFFFFFFF
//		FOR j := 0 to 7
//			i := j*64
//			reduced[63:0] := reduced[63:0] AND a[i+63:i]
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_and_epi64'.
// Requires KNCNI.
func ReduceAndEpi64(a x86.M512i) int64 {
	return int64(reduceAndEpi64([64]byte(a)))
}

func reduceAndEpi64(a [64]byte) int64


// MaskReduceGmaxPd: Determines the maximum element of the packed
// double-precision (64-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. Bitmask 'k' is used to exclude certain elements
// (elements are ignored when the corresponding mask bit is not set). 
//
//		max = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMax(max, a[i+63:i])
//			FI
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmax_pd'.
// Requires KNCNI.
func MaskReduceGmaxPd(k x86.Mmask8, a x86.M512d) float64 {
	return float64(maskReduceGmaxPd(uint8(k), [8]float64(a)))
}

func maskReduceGmaxPd(k uint8, a [8]float64) float64


// ReduceGmaxPd: Determines the maximum element of the packed double-precision
// (64-bit) floating-point elements stored in 'a' and stores the result in
// 'dst'. 
//
//		max = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			dst = FpMax(max, a[i+63:i])
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmax_pd'.
// Requires KNCNI.
func ReduceGmaxPd(a x86.M512d) float64 {
	return float64(reduceGmaxPd([8]float64(a)))
}

func reduceGmaxPd(a [8]float64) float64


// MaskReduceGmaxPs: Determines the maximum element of the packed
// single-precision (32-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. Bitmask 'k' is used to exclude certain elements
// (elements are ignored when the corresponding mask bit is not set). 
//
//		max = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMax(max, a[i+31:i])
//			FI
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmax_ps'.
// Requires KNCNI.
func MaskReduceGmaxPs(k x86.Mmask16, a x86.M512) float32 {
	return float32(maskReduceGmaxPs(uint16(k), [16]float32(a)))
}

func maskReduceGmaxPs(k uint16, a [16]float32) float32


// ReduceGmaxPs: Determines the maximum element of the packed single-precision
// (32-bit) floating-point elements stored in 'a' and stores the result in
// 'dst'. 
//
//		max = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			dst = FpMax(max, a[i+31:i])
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmax_ps'.
// Requires KNCNI.
func ReduceGmaxPs(a x86.M512) float32 {
	return float32(reduceGmaxPs([16]float32(a)))
}

func reduceGmaxPs(a [16]float32) float32


// MaskReduceGminPd: Determines the minimum element of the packed
// double-precision (64-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. Bitmask 'k' is used to exclude certain elements
// (elements are ignored when the corresponding mask bit is not set). 
//
//		min = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMin(min, a[i+63:i])
//			FI
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmin_pd'.
// Requires KNCNI.
func MaskReduceGminPd(k x86.Mmask8, a x86.M512d) float64 {
	return float64(maskReduceGminPd(uint8(k), [8]float64(a)))
}

func maskReduceGminPd(k uint8, a [8]float64) float64


// ReduceGminPd: Determines the minimum element of the packed double-precision
// (64-bit) floating-point elements stored in 'a' and stores the result in
// 'dst'. 
//
//		min = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			dst = FpMin(min, a[i+63:i])
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmin_pd'.
// Requires KNCNI.
func ReduceGminPd(a x86.M512d) float64 {
	return float64(reduceGminPd([8]float64(a)))
}

func reduceGminPd(a [8]float64) float64


// MaskReduceGminPs: Determines the minimum element of the packed
// single-precision (32-bit) floating-point elements stored in 'a' and stores
// the result in 'dst' using writemask 'k' (elements are ignored when the
// corresponding mask bit is not set). 
//
//		min = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMin(min, a[i+31:i])
//			FI
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmin_ps'.
// Requires KNCNI.
func MaskReduceGminPs(k x86.Mmask16, a x86.M512) float32 {
	return float32(maskReduceGminPs(uint16(k), [16]float32(a)))
}

func maskReduceGminPs(k uint16, a [16]float32) float32


// ReduceGminPs: Determines the minimum element of the packed single-precision
// (32-bit) floating-point elements stored in 'a' and stores the result in
// 'dst'. 
//
//		min = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			dst = FpMin(min, a[i+31:i])
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmin_ps'.
// Requires KNCNI.
func ReduceGminPs(a x86.M512) float32 {
	return float32(reduceGminPs([16]float32(a)))
}

func reduceGminPs(a [16]float32) float32


// MaskReduceMaxEpi32: Reduce the packed 32-bit integers in 'a' by maximum
// using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[31:0] := MIN_INT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epi32'.
// Requires KNCNI.
func MaskReduceMaxEpi32(k x86.Mmask16, a x86.M512i) int {
	return int(maskReduceMaxEpi32(uint16(k), [64]byte(a)))
}

func maskReduceMaxEpi32(k uint16, a [64]byte) int


// ReduceMaxEpi32: Reduce the packed 32-bit integers in 'a' by maximum. Returns
// the maximum of all elements in 'a'. 
//
//		max[31:0] := MIN_INT
//		FOR j := 0 to 15
//			i := j*32
//			max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epi32'.
// Requires KNCNI.
func ReduceMaxEpi32(a x86.M512i) int {
	return int(reduceMaxEpi32([64]byte(a)))
}

func reduceMaxEpi32(a [64]byte) int


// MaskReduceMaxEpi64: Reduce the packed 64-bit integers in 'a' by maximum
// using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[63:0] := MIN_INT
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epi64'.
// Requires KNCNI.
func MaskReduceMaxEpi64(k x86.Mmask8, a x86.M512i) int64 {
	return int64(maskReduceMaxEpi64(uint8(k), [64]byte(a)))
}

func maskReduceMaxEpi64(k uint8, a [64]byte) int64


// ReduceMaxEpi64: Reduce the packed 64-bit integers in 'a' by maximum. Returns
// the maximum of all elements in 'a'. 
//
//		max[63:0] := MIN_INT
//		FOR j := 0 to 7
//			i := j*64
//			max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epi64'.
// Requires KNCNI.
func ReduceMaxEpi64(a x86.M512i) int64 {
	return int64(reduceMaxEpi64([64]byte(a)))
}

func reduceMaxEpi64(a [64]byte) int64


// MaskReduceMaxEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// maximum using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epu32'.
// Requires KNCNI.
func MaskReduceMaxEpu32(k x86.Mmask16, a x86.M512i) uint32 {
	return uint32(maskReduceMaxEpu32(uint16(k), [64]byte(a)))
}

func maskReduceMaxEpu32(k uint16, a [64]byte) uint32


// ReduceMaxEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// maximum. Returns the maximum of all elements in 'a'. 
//
//		max[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epu32'.
// Requires KNCNI.
func ReduceMaxEpu32(a x86.M512i) uint32 {
	return uint32(reduceMaxEpu32([64]byte(a)))
}

func reduceMaxEpu32(a [64]byte) uint32


// MaskReduceMaxEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// maximum using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epu64'.
// Requires KNCNI.
func MaskReduceMaxEpu64(k x86.Mmask8, a x86.M512i) uint64 {
	return uint64(maskReduceMaxEpu64(uint8(k), [64]byte(a)))
}

func maskReduceMaxEpu64(k uint8, a [64]byte) uint64


// ReduceMaxEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// maximum. Returns the maximum of all elements in 'a'. 
//
//		max[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epu64'.
// Requires KNCNI.
func ReduceMaxEpu64(a x86.M512i) uint64 {
	return uint64(reduceMaxEpu64([64]byte(a)))
}

func reduceMaxEpu64(a [64]byte) uint64


// MaskReduceMaxPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by maximum using mask 'k'. Returns the maximum of all active
// elements in 'a'. 
//
//		max[63:0] := MIN_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_pd'.
// Requires KNCNI.
func MaskReduceMaxPd(k x86.Mmask8, a x86.M512d) float64 {
	return float64(maskReduceMaxPd(uint8(k), [8]float64(a)))
}

func maskReduceMaxPd(k uint8, a [8]float64) float64


// ReduceMaxPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by maximum. Returns the maximum of all elements in 'a'. 
//
//		max[63:0] := MIN_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_pd'.
// Requires KNCNI.
func ReduceMaxPd(a x86.M512d) float64 {
	return float64(reduceMaxPd([8]float64(a)))
}

func reduceMaxPd(a [8]float64) float64


// MaskReduceMaxPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by maximum using mask 'k'. Returns the maximum of all active
// elements in 'a'. 
//
//		max[31:0] := MIN_FLOAT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_ps'.
// Requires KNCNI.
func MaskReduceMaxPs(k x86.Mmask16, a x86.M512) float32 {
	return float32(maskReduceMaxPs(uint16(k), [16]float32(a)))
}

func maskReduceMaxPs(k uint16, a [16]float32) float32


// ReduceMaxPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by maximum. Returns the maximum of all elements in 'a'. 
//
//		max[31:0] := MIN_FLOAT
//		FOR j := 0 to 15
//			i := j*32
//			max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_ps'.
// Requires KNCNI.
func ReduceMaxPs(a x86.M512) float32 {
	return float32(reduceMaxPs([16]float32(a)))
}

func reduceMaxPs(a [16]float32) float32


// MaskReduceMinEpi32: Reduce the packed 32-bit integers in 'a' by maximum
// using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[31:0] := MAX_INT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epi32'.
// Requires KNCNI.
func MaskReduceMinEpi32(k x86.Mmask16, a x86.M512i) int {
	return int(maskReduceMinEpi32(uint16(k), [64]byte(a)))
}

func maskReduceMinEpi32(k uint16, a [64]byte) int


// ReduceMinEpi32: Reduce the packed 32-bit integers in 'a' by minimum. Returns
// the minimum of all elements in 'a'. 
//
//		min[31:0] := MAX_INT
//		FOR j := 0 to 15
//			i := j*32
//			min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epi32'.
// Requires KNCNI.
func ReduceMinEpi32(a x86.M512i) int {
	return int(reduceMinEpi32([64]byte(a)))
}

func reduceMinEpi32(a [64]byte) int


// MaskReduceMinEpi64: Reduce the packed 64-bit integers in 'a' by maximum
// using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[63:0] := MAX_INT
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epi64'.
// Requires KNCNI.
func MaskReduceMinEpi64(k x86.Mmask8, a x86.M512i) int64 {
	return int64(maskReduceMinEpi64(uint8(k), [64]byte(a)))
}

func maskReduceMinEpi64(k uint8, a [64]byte) int64


// ReduceMinEpi64: Reduce the packed 64-bit integers in 'a' by minimum. Returns
// the minimum of all elements in 'a'. 
//
//		min[63:0] := MAX_INT
//		FOR j := 0 to 7
//			i := j*64
//			min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epi64'.
// Requires KNCNI.
func ReduceMinEpi64(a x86.M512i) int64 {
	return int64(reduceMinEpi64([64]byte(a)))
}

func reduceMinEpi64(a [64]byte) int64


// MaskReduceMinEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// maximum using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[31:0] := MAX_UINT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epu32'.
// Requires KNCNI.
func MaskReduceMinEpu32(k x86.Mmask16, a x86.M512i) uint32 {
	return uint32(maskReduceMinEpu32(uint16(k), [64]byte(a)))
}

func maskReduceMinEpu32(k uint16, a [64]byte) uint32


// ReduceMinEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// minimum. Returns the minimum of all elements in 'a'. 
//
//		min[31:0] := MAX_UINT
//		FOR j := 0 to 15
//			i := j*32
//			min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epu32'.
// Requires KNCNI.
func ReduceMinEpu32(a x86.M512i) uint32 {
	return uint32(reduceMinEpu32([64]byte(a)))
}

func reduceMinEpu32(a [64]byte) uint32


// MaskReduceMinEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// minimum using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[63:0] := MAX_UINT
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epu64'.
// Requires KNCNI.
func MaskReduceMinEpu64(k x86.Mmask8, a x86.M512i) uint64 {
	return uint64(maskReduceMinEpu64(uint8(k), [64]byte(a)))
}

func maskReduceMinEpu64(k uint8, a [64]byte) uint64


// ReduceMinEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// minimum. Returns the minimum of all elements in 'a'. 
//
//		min[63:0] := MAX_UINT
//		FOR j := 0 to 7
//			i := j*64
//			min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epu64'.
// Requires KNCNI.
func ReduceMinEpu64(a x86.M512i) uint64 {
	return uint64(reduceMinEpu64([64]byte(a)))
}

func reduceMinEpu64(a [64]byte) uint64


// MaskReduceMinPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by maximum using mask 'k'. Returns the minimum of all active
// elements in 'a'. 
//
//		min[63:0] := MAX_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_pd'.
// Requires KNCNI.
func MaskReduceMinPd(k x86.Mmask8, a x86.M512d) float64 {
	return float64(maskReduceMinPd(uint8(k), [8]float64(a)))
}

func maskReduceMinPd(k uint8, a [8]float64) float64


// ReduceMinPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by minimum. Returns the minimum of all elements in 'a'. 
//
//		min[63:0] := MAX_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_pd'.
// Requires KNCNI.
func ReduceMinPd(a x86.M512d) float64 {
	return float64(reduceMinPd([8]float64(a)))
}

func reduceMinPd(a [8]float64) float64


// MaskReduceMinPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by maximum using mask 'k'. Returns the minimum of all active
// elements in 'a'. 
//
//		min[31:0] := MAX_FLOAT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_ps'.
// Requires KNCNI.
func MaskReduceMinPs(k x86.Mmask16, a x86.M512) float32 {
	return float32(maskReduceMinPs(uint16(k), [16]float32(a)))
}

func maskReduceMinPs(k uint16, a [16]float32) float32


// ReduceMinPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by minimum. Returns the minimum of all elements in 'a'. 
//
//		min[31:0] := MAX_INT
//		FOR j := 0 to 15
//			i := j*32
//			min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_ps'.
// Requires KNCNI.
func ReduceMinPs(a x86.M512) float32 {
	return float32(reduceMinPs([16]float32(a)))
}

func reduceMinPs(a [16]float32) float32


// MaskReduceMulEpi32: Reduce the packed 32-bit integers in 'a' by
// multiplication using mask 'k'. Returns the product of all active elements in
// 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				prod[31:0] := prod[31:0] * a[i+31:i]
//			FI
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_epi32'.
// Requires KNCNI.
func MaskReduceMulEpi32(k x86.Mmask16, a x86.M512i) int {
	return int(maskReduceMulEpi32(uint16(k), [64]byte(a)))
}

func maskReduceMulEpi32(k uint16, a [64]byte) int


// ReduceMulEpi32: Reduce the packed 32-bit integers in 'a' by multiplication.
// Returns the product of all elements in 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			prod[31:0] := prod[31:0] * a[i+31:i]
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_epi32'.
// Requires KNCNI.
func ReduceMulEpi32(a x86.M512i) int {
	return int(reduceMulEpi32([64]byte(a)))
}

func reduceMulEpi32(a [64]byte) int


// MaskReduceMulEpi64: Reduce the packed 64-bit integers in 'a' by
// multiplication using mask 'k'. Returns the product of all active elements in
// 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				prod[63:0] := prod[63:0] * a[i+63:i]
//			FI
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_epi64'.
// Requires KNCNI.
func MaskReduceMulEpi64(k x86.Mmask8, a x86.M512i) int64 {
	return int64(maskReduceMulEpi64(uint8(k), [64]byte(a)))
}

func maskReduceMulEpi64(k uint8, a [64]byte) int64


// ReduceMulEpi64: Reduce the packed 64-bit integers in 'a' by multiplication.
// Returns the product of all elements in 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			prod[63:0] := prod[63:0] * a[i+63:i]
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_epi64'.
// Requires KNCNI.
func ReduceMulEpi64(a x86.M512i) int64 {
	return int64(reduceMulEpi64([64]byte(a)))
}

func reduceMulEpi64(a [64]byte) int64


// MaskReduceMulPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by multiplication using mask 'k'. Returns the product of all
// active elements in 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				prod[63:0] := prod[63:0] * a[i+63:i]
//			FI
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_pd'.
// Requires KNCNI.
func MaskReduceMulPd(k x86.Mmask8, a x86.M512d) float64 {
	return float64(maskReduceMulPd(uint8(k), [8]float64(a)))
}

func maskReduceMulPd(k uint8, a [8]float64) float64


// ReduceMulPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by multiplication. Returns the product of all elements in
// 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			prod[63:0] := prod[63:0] * a[i+63:i]
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_pd'.
// Requires KNCNI.
func ReduceMulPd(a x86.M512d) float64 {
	return float64(reduceMulPd([8]float64(a)))
}

func reduceMulPd(a [8]float64) float64


// MaskReduceMulPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by multiplication using mask 'k'. Returns the product of all
// active elements in 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				prod[31:0] := prod[31:0] * a[i+31:i]
//			FI
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_ps'.
// Requires KNCNI.
func MaskReduceMulPs(k x86.Mmask16, a x86.M512) float32 {
	return float32(maskReduceMulPs(uint16(k), [16]float32(a)))
}

func maskReduceMulPs(k uint16, a [16]float32) float32


// ReduceMulPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by multiplication. Returns the product of all elements in
// 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			prod[31:0] := prod[31:0] * a[i+31:i]
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_ps'.
// Requires KNCNI.
func ReduceMulPs(a x86.M512) float32 {
	return float32(reduceMulPs([16]float32(a)))
}

func reduceMulPs(a [16]float32) float32


// MaskReduceOrEpi32: Reduce the packed 32-bit integers in 'a' by bitwise OR
// using mask 'k'. Returns the bitwise OR of all active elements in 'a'. 
//
//		reduced[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				reduced[31:0] := reduced[31:0] OR a[i+31:i]
//			FI
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_or_epi32'.
// Requires KNCNI.
func MaskReduceOrEpi32(k x86.Mmask16, a x86.M512i) int {
	return int(maskReduceOrEpi32(uint16(k), [64]byte(a)))
}

func maskReduceOrEpi32(k uint16, a [64]byte) int


// ReduceOrEpi32: Reduce the packed 32-bit integers in 'a' by bitwise OR.
// Returns the bitwise OR of all elements in 'a'. 
//
//		reduced[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			reduced[31:0] := reduced[31:0] OR a[i+31:i]
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_or_epi32'.
// Requires KNCNI.
func ReduceOrEpi32(a x86.M512i) int {
	return int(reduceOrEpi32([64]byte(a)))
}

func reduceOrEpi32(a [64]byte) int


// MaskReduceOrEpi64: Reduce the packed 64-bit integers in 'a' by bitwise OR
// using mask 'k'. Returns the bitwise OR of all active elements in 'a'. 
//
//		reduced[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				reduced[63:0] := reduced[63:0] OR a[i+63:i]
//			FI
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_or_epi64'.
// Requires KNCNI.
func MaskReduceOrEpi64(k x86.Mmask8, a x86.M512i) int64 {
	return int64(maskReduceOrEpi64(uint8(k), [64]byte(a)))
}

func maskReduceOrEpi64(k uint8, a [64]byte) int64


// ReduceOrEpi64: Reduce the packed 64-bit integers in 'a' by bitwise OR.
// Returns the bitwise OR of all elements in 'a'. 
//
//		reduced[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			reduced[63:0] := reduced[63:0] OR a[i+63:i]
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_or_epi64'.
// Requires KNCNI.
func ReduceOrEpi64(a x86.M512i) int64 {
	return int64(reduceOrEpi64([64]byte(a)))
}

func reduceOrEpi64(a [64]byte) int64


// MaskRoundPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' to the nearest integer value using 'expadj' and in the
// direction of 'rounding', and store the results as packed single-precision
// floating-point elements in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ROUND(a[i+31:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm512_mask_round_ps'.
// Requires KNCNI.
func MaskRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(maskRoundPs([16]float32(src), uint16(k), [16]float32(a), rounding, expadj))
}

func maskRoundPs(src [16]float32, k uint16, a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32


// RoundPs: Round the packed single-precision (32-bit) floating-point elements
// in 'a' to the nearest integer value using 'expadj' and in the direction of
// 'rounding', and store the results as packed single-precision floating-point
// elements in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm512_round_ps'.
// Requires KNCNI.
func RoundPs(a x86.M512, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(roundPs([16]float32(a), rounding, expadj))
}

func roundPs(a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32


// MaskRoundfxpntAdjustPd: Performs element-by-element rounding of packed
// double-precision (64-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set).
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
//				dst[i+63:i] := ROUND(a[i+63:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPD'. Intrinsic: '_mm512_mask_roundfxpnt_adjust_pd'.
// Requires KNCNI.
func MaskRoundfxpntAdjustPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int, expadj MMEXPADJENUM) x86.M512d {
	return x86.M512d(maskRoundfxpntAdjustPd([8]float64(src), uint8(k), [8]float64(a), rounding, expadj))
}

func maskRoundfxpntAdjustPd(src [8]float64, k uint8, a [8]float64, rounding int, expadj MMEXPADJENUM) [8]float64


// RoundfxpntAdjustPd: Performs element-by-element rounding of packed
// double-precision (64-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst'.
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
//			dst[i+63:i] := ROUND(a[i+63:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPD'. Intrinsic: '_mm512_roundfxpnt_adjust_pd'.
// Requires KNCNI.
func RoundfxpntAdjustPd(a x86.M512d, rounding int, expadj MMEXPADJENUM) x86.M512d {
	return x86.M512d(roundfxpntAdjustPd([8]float64(a), rounding, expadj))
}

func roundfxpntAdjustPd(a [8]float64, rounding int, expadj MMEXPADJENUM) [8]float64


// MaskRoundfxpntAdjustPs: Performs element-by-element rounding of packed
// single-precision (32-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ROUND(a[i+31:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPS'. Intrinsic: '_mm512_mask_roundfxpnt_adjust_ps'.
// Requires KNCNI.
func MaskRoundfxpntAdjustPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(maskRoundfxpntAdjustPs([16]float32(src), uint16(k), [16]float32(a), rounding, expadj))
}

func maskRoundfxpntAdjustPs(src [16]float32, k uint16, a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32


// RoundfxpntAdjustPs: Performs element-by-element rounding of packed
// single-precision (32-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPS'. Intrinsic: '_mm512_roundfxpnt_adjust_ps'.
// Requires KNCNI.
func RoundfxpntAdjustPs(a x86.M512, rounding int, expadj MMEXPADJENUM) x86.M512 {
	return x86.M512(roundfxpntAdjustPs([16]float32(a), rounding, expadj))
}

func roundfxpntAdjustPs(a [16]float32, rounding int, expadj MMEXPADJENUM) [16]float32


// MaskRsqrt23Ps: Calculates the reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a' to 23 bits of
// accuracy and stores the result in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := Sqrt(1.0 / a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRSQRT23PS'. Intrinsic: '_mm512_mask_rsqrt23_ps'.
// Requires KNCNI.
func MaskRsqrt23Ps(src x86.M512, k x86.Mmask16, a x86.M512) x86.M512 {
	return x86.M512(maskRsqrt23Ps([16]float32(src), uint16(k), [16]float32(a)))
}

func maskRsqrt23Ps(src [16]float32, k uint16, a [16]float32) [16]float32


// Rsqrt23Ps: Calculates the reciprocal square root of packed single-precision
// (32-bit) floating-point elements in 'a' to 23 bits of accuracy and stores
// the result in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Sqrt(1.0 / a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRSQRT23PS'. Intrinsic: '_mm512_rsqrt23_ps'.
// Requires KNCNI.
func Rsqrt23Ps(a x86.M512) x86.M512 {
	return x86.M512(rsqrt23Ps([16]float32(a)))
}

func rsqrt23Ps(a [16]float32) [16]float32


// MaskSbbEpi32: Performs element-by-element three-input subtraction of packed
// 32-bit integer elements of 'v3' as well as the corresponding bit from 'k2'
// from 'v2'. The borrowed value from the subtraction difference for the nth
// element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst' using writemask 'k1' (elements are copied from 'v2' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				dst[i+31:i] := v2[i+31:i] - v3[i+31:i] - k2[j]
//				borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k2[j])
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBD'. Intrinsic: '_mm512_mask_sbb_epi32'.
// Requires KNCNI.
func MaskSbbEpi32(v2 x86.M512i, k1 x86.Mmask16, k2 x86.Mmask16, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(maskSbbEpi32([64]byte(v2), uint16(k1), uint16(k2), [64]byte(v3), uint16(borrow)))
}

func maskSbbEpi32(v2 [64]byte, k1 uint16, k2 uint16, v3 [64]byte, borrow uint16) [64]byte


// SbbEpi32: Performs element-by-element three-input subtraction of packed
// 32-bit integer elements of 'v3' as well as the corresponding bit from 'k'
// from 'v2'. The borrowed value from the subtraction difference for the nth
// element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] - v3[i+31:i] - k[j]
//			borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k[j])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBD'. Intrinsic: '_mm512_sbb_epi32'.
// Requires KNCNI.
func SbbEpi32(v2 x86.M512i, k x86.Mmask16, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(sbbEpi32([64]byte(v2), uint16(k), [64]byte(v3), uint16(borrow)))
}

func sbbEpi32(v2 [64]byte, k uint16, v3 [64]byte, borrow uint16) [64]byte


// MaskSbbrEpi32: Performs element-by-element three-input subtraction of packed
// 32-bit integer elements of 'v2' as well as the corresponding bit from 'k2'
// from 'v3'. The borrowed value from the subtraction difference for the nth
// element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst' using writemask 'k1' (elements are copied from 'v2' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i] - k2[j]
//				borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k[j])
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBRD'. Intrinsic: '_mm512_mask_sbbr_epi32'.
// Requires KNCNI.
func MaskSbbrEpi32(v2 x86.M512i, k1 x86.Mmask16, k2 x86.Mmask16, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(maskSbbrEpi32([64]byte(v2), uint16(k1), uint16(k2), [64]byte(v3), uint16(borrow)))
}

func maskSbbrEpi32(v2 [64]byte, k1 uint16, k2 uint16, v3 [64]byte, borrow uint16) [64]byte


// SbbrEpi32: Performs element-by-element three-input subtraction of packed
// 32-bit integer elements of 'v2' as well as the corresponding bit from 'k'
// from 'v3'. The borrowed value from the subtraction difference for the nth
// element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i] - k[j]
//			borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k[j])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBRD'. Intrinsic: '_mm512_sbbr_epi32'.
// Requires KNCNI.
func SbbrEpi32(v2 x86.M512i, k x86.Mmask16, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(sbbrEpi32([64]byte(v2), uint16(k), [64]byte(v3), uint16(borrow)))
}

func sbbrEpi32(v2 [64]byte, k uint16, v3 [64]byte, borrow uint16) [64]byte


// MaskScalePs: Scales each single-precision (32-bit) floating-point element in
// 'a' by multiplying it by 2**exponent, where the exponenet is the
// corresponding 32-bit integer element in 'b', storing results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSCALEPS'. Intrinsic: '_mm512_mask_scale_ps'.
// Requires KNCNI.
func MaskScalePs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512i) x86.M512 {
	return x86.M512(maskScalePs([16]float32(src), uint16(k), [16]float32(a), [64]byte(b)))
}

func maskScalePs(src [16]float32, k uint16, a [16]float32, b [64]byte) [16]float32


// ScalePs: Scales each single-precision (32-bit) floating-point element in 'a'
// by multiplying it by 2**exponent, where the exponent is the corresponding
// 32-bit integer element in 'b', storing results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSCALEPS'. Intrinsic: '_mm512_scale_ps'.
// Requires KNCNI.
func ScalePs(a x86.M512, b x86.M512i) x86.M512 {
	return x86.M512(scalePs([16]float32(a), [64]byte(b)))
}

func scalePs(a [16]float32, b [64]byte) [16]float32


// MaskScaleRoundPs: Scales each single-precision (32-bit) floating-point
// element in 'a' by multiplying it by 2**exp, where the exp is the
// corresponding 32-bit integer element in 'b', storing results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). Results are rounded using constant 'rounding'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSCALEPS'. Intrinsic: '_mm512_mask_scale_round_ps'.
// Requires KNCNI.
func MaskScaleRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512i, rounding int) x86.M512 {
	return x86.M512(maskScaleRoundPs([16]float32(src), uint16(k), [16]float32(a), [64]byte(b), rounding))
}

func maskScaleRoundPs(src [16]float32, k uint16, a [16]float32, b [64]byte, rounding int) [16]float32


// ScaleRoundPs: Scales each single-precision (32-bit) floating-point element
// in 'a' by multiplying it by 2**exponent, where the exponenet is the
// corresponding 32-bit integer element in 'b', storing results in 'dst'.
// Intermediate elements are rounded using 'rounding'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_scale_round_ps'.
// Requires KNCNI.
func ScaleRoundPs(a x86.M512, b x86.M512i, rounding int) x86.M512 {
	return x86.M512(scaleRoundPs([16]float32(a), [64]byte(b), rounding))
}

func scaleRoundPs(a [16]float32, b [64]byte, rounding int) [16]float32


// MaskShuffleEpi32: Shuffle 32-bit integers in 'a' within 128-bit lanes using
// the control in 'imm8', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[31:0] := src[31:0]
//			1:	tmp[31:0] := src[63:32]
//			2:	tmp[31:0] := src[95:64]
//			3:	tmp[31:0] := src[127:96]
//			ESAC
//			RETURN tmp[31:0]
//		}
//		
//		tmp_dst[31:0] := SELECT4(a[127:0], imm8[1:0])
//		tmp_dst[63:32] := SELECT4(a[127:0], imm8[3:2])
//		tmp_dst[95:64] := SELECT4(a[127:0], imm8[5:4])
//		tmp_dst[127:96] := SELECT4(a[127:0], imm8[7:6])
//		tmp_dst[159:128] := SELECT4(a[255:128], imm8[1:0])
//		tmp_dst[191:160] := SELECT4(a[255:128], imm8[3:2])
//		tmp_dst[223:192] := SELECT4(a[255:128], imm8[5:4])
//		tmp_dst[255:224] := SELECT4(a[255:128], imm8[7:6])
//		tmp_dst[287:256] := SELECT4(a[383:256], imm8[1:0])
//		tmp_dst[319:288] := SELECT4(a[383:256], imm8[3:2])
//		tmp_dst[351:320] := SELECT4(a[383:256], imm8[5:4])
//		tmp_dst[383:352] := SELECT4(a[383:256], imm8[7:6])
//		tmp_dst[415:384] := SELECT4(a[511:384], imm8[1:0])
//		tmp_dst[447:416] := SELECT4(a[511:384], imm8[3:2])
//		tmp_dst[479:448] := SELECT4(a[511:384], imm8[5:4])
//		tmp_dst[511:480] := SELECT4(a[511:384], imm8[7:6])
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp_dst[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFD'. Intrinsic: '_mm512_mask_shuffle_epi32'.
// Requires KNCNI.
func MaskShuffleEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 MMPERMENUM) x86.M512i {
	return x86.M512i(maskShuffleEpi32([64]byte(src), uint16(k), [64]byte(a), imm8))
}

func maskShuffleEpi32(src [64]byte, k uint16, a [64]byte, imm8 MMPERMENUM) [64]byte


// ShuffleEpi32: Shuffle 32-bit integers in 'a' within 128-bit lanes using the
// control in 'imm8', and store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[31:0] := src[31:0]
//			1:	tmp[31:0] := src[63:32]
//			2:	tmp[31:0] := src[95:64]
//			3:	tmp[31:0] := src[127:96]
//			ESAC
//			RETURN tmp[31:0]
//		}
//		
//		dst[31:0] := SELECT4(a[127:0], imm8[1:0])
//		dst[63:32] := SELECT4(a[127:0], imm8[3:2])
//		dst[95:64] := SELECT4(a[127:0], imm8[5:4])
//		dst[127:96] := SELECT4(a[127:0], imm8[7:6])
//		dst[159:128] := SELECT4(a[255:128], imm8[1:0])
//		dst[191:160] := SELECT4(a[255:128], imm8[3:2])
//		dst[223:192] := SELECT4(a[255:128], imm8[5:4])
//		dst[255:224] := SELECT4(a[255:128], imm8[7:6])
//		dst[287:256] := SELECT4(a[383:256], imm8[1:0])
//		dst[319:288] := SELECT4(a[383:256], imm8[3:2])
//		dst[351:320] := SELECT4(a[383:256], imm8[5:4])
//		dst[383:352] := SELECT4(a[383:256], imm8[7:6])
//		dst[415:384] := SELECT4(a[511:384], imm8[1:0])
//		dst[447:416] := SELECT4(a[511:384], imm8[3:2])
//		dst[479:448] := SELECT4(a[511:384], imm8[5:4])
//		dst[511:480] := SELECT4(a[511:384], imm8[7:6])
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFD'. Intrinsic: '_mm512_shuffle_epi32'.
// Requires KNCNI.
func ShuffleEpi32(a x86.M512i, imm8 MMPERMENUM) x86.M512i {
	return x86.M512i(shuffleEpi32([64]byte(a), imm8))
}

func shuffleEpi32(a [64]byte, imm8 MMPERMENUM) [64]byte


// MaskSlliEpi32: Shift packed 32-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF imm8[7:0] > 31
//					dst[i+31:i] := 0
//				ELSE
//					dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLD'. Intrinsic: '_mm512_mask_slli_epi32'.
// Requires KNCNI.
func MaskSlliEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(maskSlliEpi32([64]byte(src), uint16(k), [64]byte(a), imm8))
}

func maskSlliEpi32(src [64]byte, k uint16, a [64]byte, imm8 uint32) [64]byte


// SlliEpi32: Shift packed 32-bit integers in 'a' left by 'imm8' while shifting
// in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLD'. Intrinsic: '_mm512_slli_epi32'.
// Requires KNCNI.
func SlliEpi32(a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(slliEpi32([64]byte(a), imm8))
}

func slliEpi32(a [64]byte, imm8 uint32) [64]byte


// MaskSllvEpi32: Shift packed 32-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << count[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVD'. Intrinsic: '_mm512_mask_sllv_epi32'.
// Requires KNCNI.
func MaskSllvEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(maskSllvEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(count)))
}

func maskSllvEpi32(src [64]byte, k uint16, a [64]byte, count [64]byte) [64]byte


// SllvEpi32: Shift packed 32-bit integers in 'a' left by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] << count[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVD'. Intrinsic: '_mm512_sllv_epi32'.
// Requires KNCNI.
func SllvEpi32(a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(sllvEpi32([64]byte(a), [64]byte(count)))
}

func sllvEpi32(a [64]byte, count [64]byte) [64]byte


// Spflt32: Set performance monitoring filtering mask to 32-bit unsigned
// integer 'r1'. 
//
//		SetPerfMonMask(r1[31:0])
//
// Instruction: 'SPFLT'. Intrinsic: '_mm_spflt_32'.
// Requires KNCNI.
func Spflt32(r1 uint32)  {
	spflt32(r1)
}

func spflt32(r1 uint32) 


// Spflt64: Set performance monitoring filtering mask to 64-bit unsigned
// integer 'r1'. 
//
//		SetPerfMonMask(r1[63:0])
//
// Instruction: 'SPFLT'. Intrinsic: '_mm_spflt_64'.
// Requires KNCNI.
func Spflt64(r1 uint64)  {
	spflt64(r1)
}

func spflt64(r1 uint64) 


// MaskSraiEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF imm8[7:0] > 31
//					dst[i+31:i] := SignBit
//				ELSE
//					dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAD'. Intrinsic: '_mm512_mask_srai_epi32'.
// Requires KNCNI.
func MaskSraiEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(maskSraiEpi32([64]byte(src), uint16(k), [64]byte(a), imm8))
}

func maskSraiEpi32(src [64]byte, k uint16, a [64]byte, imm8 uint32) [64]byte


// SraiEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAD'. Intrinsic: '_mm512_srai_epi32'.
// Requires KNCNI.
func SraiEpi32(a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(sraiEpi32([64]byte(a), imm8))
}

func sraiEpi32(a [64]byte, imm8 uint32) [64]byte


// MaskSravEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := SignExtend(a[i+31:i] >> count[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVD'. Intrinsic: '_mm512_mask_srav_epi32'.
// Requires KNCNI.
func MaskSravEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(maskSravEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(count)))
}

func maskSravEpi32(src [64]byte, k uint16, a [64]byte, count [64]byte) [64]byte


// SravEpi32: Shift packed 32-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in sign bits, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := SignExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVD'. Intrinsic: '_mm512_srav_epi32'.
// Requires KNCNI.
func SravEpi32(a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(sravEpi32([64]byte(a), [64]byte(count)))
}

func sravEpi32(a [64]byte, count [64]byte) [64]byte


// MaskSrliEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF imm8[7:0] > 31
//					dst[i+31:i] := 0
//				ELSE
//					dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLD'. Intrinsic: '_mm512_mask_srli_epi32'.
// Requires KNCNI.
func MaskSrliEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(maskSrliEpi32([64]byte(src), uint16(k), [64]byte(a), imm8))
}

func maskSrliEpi32(src [64]byte, k uint16, a [64]byte, imm8 uint32) [64]byte


// SrliEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLD'. Intrinsic: '_mm512_srli_epi32'.
// Requires KNCNI.
func SrliEpi32(a x86.M512i, imm8 uint32) x86.M512i {
	return x86.M512i(srliEpi32([64]byte(a), imm8))
}

func srliEpi32(a [64]byte, imm8 uint32) [64]byte


// MaskSrlvEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVD'. Intrinsic: '_mm512_mask_srlv_epi32'.
// Requires KNCNI.
func MaskSrlvEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(maskSrlvEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(count)))
}

func maskSrlvEpi32(src [64]byte, k uint16, a [64]byte, count [64]byte) [64]byte


// SrlvEpi32: Shift packed 32-bit integers in 'a' right by the amount specified
// by the corresponding element in 'count' while shifting in zeros, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVD'. Intrinsic: '_mm512_srlv_epi32'.
// Requires KNCNI.
func SrlvEpi32(a x86.M512i, count x86.M512i) x86.M512i {
	return x86.M512i(srlvEpi32([64]byte(a), [64]byte(count)))
}

func srlvEpi32(a [64]byte, count [64]byte) [64]byte


// MaskStoreEpi32: Store packed 32-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				MEM[mem_addr+i+31:mem_addr+i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_mask_store_epi32'.
// Requires KNCNI.
func MaskStoreEpi32(mem_addr uintptr, k x86.Mmask16, a x86.M512i)  {
	maskStoreEpi32(uintptr(mem_addr), uint16(k), [64]byte(a))
}

func maskStoreEpi32(mem_addr uintptr, k uint16, a [64]byte) 


// StoreEpi32: Store 512-bits (composed of 16 packed 32-bit integers) from 'a'
// into memory. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+511:mem_addr] := a[511:0]
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_store_epi32'.
// Requires KNCNI.
func StoreEpi32(mem_addr uintptr, a x86.M512i)  {
	storeEpi32(uintptr(mem_addr), [64]byte(a))
}

func storeEpi32(mem_addr uintptr, a [64]byte) 


// MaskStoreEpi64: Store packed 64-bit integers from 'a' into memory using
// writemask 'k'.
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				MEM[mem_addr+i+63:mem_addr+i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_mask_store_epi64'.
// Requires KNCNI.
func MaskStoreEpi64(mem_addr uintptr, k x86.Mmask8, a x86.M512i)  {
	maskStoreEpi64(uintptr(mem_addr), uint8(k), [64]byte(a))
}

func maskStoreEpi64(mem_addr uintptr, k uint8, a [64]byte) 


// StoreEpi64: Store 512-bits (composed of 8 packed 64-bit integers) from 'a'
// into memory. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+511:mem_addr] := a[511:0]
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_store_epi64'.
// Requires KNCNI.
func StoreEpi64(mem_addr uintptr, a x86.M512i)  {
	storeEpi64(uintptr(mem_addr), [64]byte(a))
}

func storeEpi64(mem_addr uintptr, a [64]byte) 


// MaskStorePd: Store packed double-precision (64-bit) floating-point elements
// from 'a' into memory using writemask 'k'.
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				MEM[mem_addr+i+63:mem_addr+i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_mask_store_pd'.
// Requires KNCNI.
func MaskStorePd(mem_addr uintptr, k x86.Mmask8, a x86.M512d)  {
	maskStorePd(uintptr(mem_addr), uint8(k), [8]float64(a))
}

func maskStorePd(mem_addr uintptr, k uint8, a [8]float64) 


// StorePd: Store 512-bits (composed of 8 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+511:mem_addr] := a[511:0]
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_store_pd'.
// Requires KNCNI.
func StorePd(mem_addr uintptr, a x86.M512d)  {
	storePd(uintptr(mem_addr), [8]float64(a))
}

func storePd(mem_addr uintptr, a [8]float64) 


// MaskStorePs: Store packed single-precision (32-bit) floating-point elements
// from 'a' into memory using writemask 'k'.
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				MEM[mem_addr+i+31:mem_addr+i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_mask_store_ps'.
// Requires KNCNI.
func MaskStorePs(mem_addr uintptr, k x86.Mmask16, a x86.M512)  {
	maskStorePs(uintptr(mem_addr), uint16(k), [16]float32(a))
}

func maskStorePs(mem_addr uintptr, k uint16, a [16]float32) 


// StorePs: Store 512-bits (composed of 16 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+511:mem_addr] := a[511:0]
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_store_ps'.
// Requires KNCNI.
func StorePs(mem_addr uintptr, a x86.M512)  {
	storePs(uintptr(mem_addr), [16]float32(a))
}

func storePs(mem_addr uintptr, a [16]float32) 


// StoreSi512: Store 512-bits of integer data from 'a' into memory. 
// 	'mem_addr' must be aligned on a 64-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+511:mem_addr] := a[511:0]
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_store_si512'.
// Requires KNCNI.
func StoreSi512(mem_addr uintptr, a x86.M512i)  {
	storeSi512(uintptr(mem_addr), [64]byte(a))
}

func storeSi512(mem_addr uintptr, a [64]byte) 


// StorenrPd: Stores packed double-precision (64-bit) floating-point elements
// from 'v' to memory address 'mt' with a no-read hint to the processor. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			addr[i+63:i] := v[i+63:i]
//		ENDFOR
//
// Instruction: 'VMOVNRAPD'. Intrinsic: '_mm512_storenr_pd'.
// Requires KNCNI.
func StorenrPd(mt uintptr, v x86.M512d)  {
	storenrPd(uintptr(mt), [8]float64(v))
}

func storenrPd(mt uintptr, v [8]float64) 


// StorenrPs: Stores packed single-precision (32-bit) floating-point elements
// from 'v' to memory address 'mt' with a no-read hint to the processor. 
//
//		addr := MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			addr[i+31:i] := v[i+31:i]
//		ENDFOR
//
// Instruction: 'VMOVNRAPS'. Intrinsic: '_mm512_storenr_ps'.
// Requires KNCNI.
func StorenrPs(mt uintptr, v x86.M512)  {
	storenrPs(uintptr(mt), [16]float32(v))
}

func storenrPs(mt uintptr, v [16]float32) 


// StorenrngoPd: Stores packed double-precision (64-bit) floating-point
// elements from 'v' to memory address 'mt' with a no-read hint and using a
// weakly-ordered memory consistency model (stores performed with this function
// are not globally ordered, and subsequent stores from the same thread can be
// observed before them). 
//
//		addr := MEM[mt]
//		FOR j := 0 to 7
//			i := j*64
//			addr[i+63:i] := v[i+63:i]
//		ENDFOR
//
// Instruction: 'VMOVNRNGOAPD'. Intrinsic: '_mm512_storenrngo_pd'.
// Requires KNCNI.
func StorenrngoPd(mt uintptr, v x86.M512d)  {
	storenrngoPd(uintptr(mt), [8]float64(v))
}

func storenrngoPd(mt uintptr, v [8]float64) 


// StorenrngoPs: Stores packed single-precision (32-bit) floating-point
// elements from 'v' to memory address 'mt' with a no-read hint and using a
// weakly-ordered memory consistency model (stores performed with this function
// are not globally ordered, and subsequent stores from the same thread can be
// observed before them). 
//
//		addr := MEM[mt]
//		FOR j := 0 to 15
//			i := j*32
//			addr[i+31:i] := v[i+31:i]
//		ENDFOR
//
// Instruction: 'VMOVNRNGOAPS'. Intrinsic: '_mm512_storenrngo_ps'.
// Requires KNCNI.
func StorenrngoPs(mt uintptr, v x86.M512)  {
	storenrngoPs(uintptr(mt), [16]float32(v))
}

func storenrngoPs(mt uintptr, v [16]float32) 


// MaskSubEpi32: Subtract packed 32-bit integers in 'b' from packed 32-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBD'. Intrinsic: '_mm512_mask_sub_epi32'.
// Requires KNCNI.
func MaskSubEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskSubEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskSubEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// SubEpi32: Subtract packed 32-bit integers in 'b' from packed 32-bit integers
// in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBD'. Intrinsic: '_mm512_sub_epi32'.
// Requires KNCNI.
func SubEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(subEpi32([64]byte(a), [64]byte(b)))
}

func subEpi32(a [64]byte, b [64]byte) [64]byte


// MaskSubPd: Subtract packed double-precision (64-bit) floating-point elements
// in 'b' from packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] - b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_mask_sub_pd'.
// Requires KNCNI.
func MaskSubPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(maskSubPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b)))
}

func maskSubPd(src [8]float64, k uint8, a [8]float64, b [8]float64) [8]float64


// SubPd: Subtract packed double-precision (64-bit) floating-point elements in
// 'b' from packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_sub_pd'.
// Requires KNCNI.
func SubPd(a x86.M512d, b x86.M512d) x86.M512d {
	return x86.M512d(subPd([8]float64(a), [8]float64(b)))
}

func subPd(a [8]float64, b [8]float64) [8]float64


// MaskSubPs: Subtract packed single-precision (32-bit) floating-point elements
// in 'b' from packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_mask_sub_ps'.
// Requires KNCNI.
func MaskSubPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(maskSubPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b)))
}

func maskSubPs(src [16]float32, k uint16, a [16]float32, b [16]float32) [16]float32


// SubPs: Subtract packed single-precision (32-bit) floating-point elements in
// 'b' from packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_sub_ps'.
// Requires KNCNI.
func SubPs(a x86.M512, b x86.M512) x86.M512 {
	return x86.M512(subPs([16]float32(a), [16]float32(b)))
}

func subPs(a [16]float32, b [16]float32) [16]float32


// MaskSubRoundPd: Subtract packed double-precision (64-bit) floating-point
// elements in 'b' from packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
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
//				dst[i+63:i] := a[i+63:i] - b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_mask_sub_round_pd'.
// Requires KNCNI.
func MaskSubRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskSubRoundPd([8]float64(src), uint8(k), [8]float64(a), [8]float64(b), rounding))
}

func maskSubRoundPd(src [8]float64, k uint8, a [8]float64, b [8]float64, rounding int) [8]float64


// SubRoundPd: Subtract packed double-precision (64-bit) floating-point
// elements in 'b' from packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst'.
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
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_sub_round_pd'.
// Requires KNCNI.
func SubRoundPd(a x86.M512d, b x86.M512d, rounding int) x86.M512d {
	return x86.M512d(subRoundPd([8]float64(a), [8]float64(b), rounding))
}

func subRoundPd(a [8]float64, b [8]float64, rounding int) [8]float64


// MaskSubRoundPs: Subtract packed single-precision (32-bit) floating-point
// elements in 'b' from packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_mask_sub_round_ps'.
// Requires KNCNI.
func MaskSubRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(maskSubRoundPs([16]float32(src), uint16(k), [16]float32(a), [16]float32(b), rounding))
}

func maskSubRoundPs(src [16]float32, k uint16, a [16]float32, b [16]float32, rounding int) [16]float32


// SubRoundPs: Subtract packed single-precision (32-bit) floating-point
// elements in 'b' from packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_sub_round_ps'.
// Requires KNCNI.
func SubRoundPs(a x86.M512, b x86.M512, rounding int) x86.M512 {
	return x86.M512(subRoundPs([16]float32(a), [16]float32(b), rounding))
}

func subRoundPs(a [16]float32, b [16]float32, rounding int) [16]float32


// MaskSubrEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v2' from 'v3' storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set) 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRD'. Intrinsic: '_mm512_mask_subr_epi32'.
// Requires KNCNI.
func MaskSubrEpi32(src x86.M512i, k x86.Mmask16, v2 x86.M512i, v3 x86.M512i) x86.M512i {
	return x86.M512i(maskSubrEpi32([64]byte(src), uint16(k), [64]byte(v2), [64]byte(v3)))
}

func maskSubrEpi32(src [64]byte, k uint16, v2 [64]byte, v3 [64]byte) [64]byte


// SubrEpi32: Performs element-by-element subtraction of packed 32-bit integer
// elements in 'v2' from 'v3' storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRD'. Intrinsic: '_mm512_subr_epi32'.
// Requires KNCNI.
func SubrEpi32(v2 x86.M512i, v3 x86.M512i) x86.M512i {
	return x86.M512i(subrEpi32([64]byte(v2), [64]byte(v3)))
}

func subrEpi32(v2 [64]byte, v3 [64]byte) [64]byte


// MaskSubrPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_mask_subr_pd'.
// Requires KNCNI.
func MaskSubrPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d) x86.M512d {
	return x86.M512d(maskSubrPd([8]float64(src), uint8(k), [8]float64(v2), [8]float64(v3)))
}

func maskSubrPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64) [8]float64


// SubrPd: Performs element-by-element subtraction of packed double-precision
// (64-bit) floating-point elements in 'v2' from 'v3' storing the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_subr_pd'.
// Requires KNCNI.
func SubrPd(v2 x86.M512d, v3 x86.M512d) x86.M512d {
	return x86.M512d(subrPd([8]float64(v2), [8]float64(v3)))
}

func subrPd(v2 [8]float64, v3 [8]float64) [8]float64


// MaskSubrPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_mask_subr_ps'.
// Requires KNCNI.
func MaskSubrPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512) x86.M512 {
	return x86.M512(maskSubrPs([16]float32(src), uint16(k), [16]float32(v2), [16]float32(v3)))
}

func maskSubrPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32) [16]float32


// SubrPs: Performs element-by-element subtraction of packed single-precision
// (32-bit) floating-point elements in 'v2' from 'v3' storing the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_subr_ps'.
// Requires KNCNI.
func SubrPs(v2 x86.M512, v3 x86.M512) x86.M512 {
	return x86.M512(subrPs([16]float32(v2), [16]float32(v3)))
}

func subrPs(v2 [16]float32, v3 [16]float32) [16]float32


// MaskSubrRoundPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
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
//				dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_mask_subr_round_pd'.
// Requires KNCNI.
func MaskSubrRoundPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d, rounding int) x86.M512d {
	return x86.M512d(maskSubrRoundPd([8]float64(src), uint8(k), [8]float64(v2), [8]float64(v3), rounding))
}

func maskSubrRoundPd(src [8]float64, k uint8, v2 [8]float64, v3 [8]float64, rounding int) [8]float64


// SubrRoundPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst'.
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
//			dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_subr_round_pd'.
// Requires KNCNI.
func SubrRoundPd(v2 x86.M512d, v3 x86.M512d, rounding int) x86.M512d {
	return x86.M512d(subrRoundPd([8]float64(v2), [8]float64(v3), rounding))
}

func subrRoundPd(v2 [8]float64, v3 [8]float64, rounding int) [8]float64


// MaskSubrRoundPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
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
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_mask_subr_round_ps'.
// Requires KNCNI.
func MaskSubrRoundPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, rounding int) x86.M512 {
	return x86.M512(maskSubrRoundPs([16]float32(src), uint16(k), [16]float32(v2), [16]float32(v3), rounding))
}

func maskSubrRoundPs(src [16]float32, k uint16, v2 [16]float32, v3 [16]float32, rounding int) [16]float32


// SubrRoundPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_subr_round_ps'.
// Requires KNCNI.
func SubrRoundPs(v2 x86.M512, v3 x86.M512, rounding int) x86.M512 {
	return x86.M512(subrRoundPs([16]float32(v2), [16]float32(v3), rounding))
}

func subrRoundPs(v2 [16]float32, v3 [16]float32, rounding int) [16]float32


// MaskSubrsetbEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v2' from 'v3', storing the results in 'dst' and 'v2'.
// The borrowed value from the subtraction difference for the nth element is
// written to the nth bit of 'borrow' (borrow flag). Results are written using
// writemask 'k' (elements are copied from 'k' to 'k_old' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				diff := v3[i+31:i] - v2[i+31:i]
//				borrow[j] := Borrow(v3[i+31:i] - v2[i+31:i])
//				dst[i+31:i] := diff
//				v2[i+31:i] := diff
//			ELSE
//				borrow[j] := k_old[j]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRSETBD'. Intrinsic: '_mm512_mask_subrsetb_epi32'.
// Requires KNCNI.
func MaskSubrsetbEpi32(v2 x86.M512i, k x86.Mmask16, k_old x86.Mmask16, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(maskSubrsetbEpi32([64]byte(v2), uint16(k), uint16(k_old), [64]byte(v3), uint16(borrow)))
}

func maskSubrsetbEpi32(v2 [64]byte, k uint16, k_old uint16, v3 [64]byte, borrow uint16) [64]byte


// SubrsetbEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v2' from 'v3', storing the results in 'dst' and 'v2'.
// The borrowed value from the subtraction difference for the nth element is
// written to the nth bit of 'borrow' (borrow flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			borrow[j] := Borrow(v3[i+31:i] - v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRSETBD'. Intrinsic: '_mm512_subrsetb_epi32'.
// Requires KNCNI.
func SubrsetbEpi32(v2 x86.M512i, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(subrsetbEpi32([64]byte(v2), [64]byte(v3), uint16(borrow)))
}

func subrsetbEpi32(v2 [64]byte, v3 [64]byte, borrow uint16) [64]byte


// MaskSubsetbEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v3' from 'v2', storing the results in 'dst' and the nth
// borrow bit in the nth position of 'borrow' (borrow flag). Results are stored
// using writemask 'k' (elements are copied from 'v2' and 'k_old' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] - v3[i+31:i]
//				borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//				borrow[j] := k_old[j]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSETBD'. Intrinsic: '_mm512_mask_subsetb_epi32'.
// Requires KNCNI.
func MaskSubsetbEpi32(v2 x86.M512i, k x86.Mmask16, k_old x86.Mmask16, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(maskSubsetbEpi32([64]byte(v2), uint16(k), uint16(k_old), [64]byte(v3), uint16(borrow)))
}

func maskSubsetbEpi32(v2 [64]byte, k uint16, k_old uint16, v3 [64]byte, borrow uint16) [64]byte


// SubsetbEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v3' from 'v2', storing the results in 'dst' and the nth
// borrow bit in the nth position of 'borrow' (borrow flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] - v3[i+31:i]
//			borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSETBD'. Intrinsic: '_mm512_subsetb_epi32'.
// Requires KNCNI.
func SubsetbEpi32(v2 x86.M512i, v3 x86.M512i, borrow x86.Mmask16) x86.M512i {
	return x86.M512i(subsetbEpi32([64]byte(v2), [64]byte(v3), uint16(borrow)))
}

func subsetbEpi32(v2 [64]byte, v3 [64]byte, borrow uint16) [64]byte


// MaskSwizzleEpi32: Performs a swizzle transformation of each of the four
// groups of packed 4x32-bit integer elements in 'v' using swizzle parameter
// 's', storing the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				IF k[j*2]
//					dst[i+31:i]	:= v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	:= src[i+31:i]
//				FI
//				IF k[j*2+1]
//					dst[i+63:i+32] := v[i+31:i]
//				ELSE
//					dst[i+63:i+32] := src[i+63:i+32]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+31:i]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+31:i]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+31:i]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+63:i+32]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+63:i+32]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+95:i+64]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+95:i+64]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+127:i+96]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+127:i+96]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_epi32'.
// Requires KNCNI.
func MaskSwizzleEpi32(src x86.M512i, k x86.Mmask16, v x86.M512i, s MMSWIZZLEENUM) x86.M512i {
	return x86.M512i(maskSwizzleEpi32([64]byte(src), uint16(k), [64]byte(v), s))
}

func maskSwizzleEpi32(src [64]byte, k uint16, v [64]byte, s MMSWIZZLEENUM) [64]byte


// SwizzleEpi32: Performs a swizzle transformation of each of the four groups
// of packed 4x 32-bit integer elements in 'v' using swizzle parameter 's',
// storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				dst[i+31:i]    := v[i+63:i+32]
//				dst[i+63:i+32] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+31:i]
//				dst[i+63:i+32]  := v[i+31:i]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+63:i+32]
//				dst[i+95:i+64]  := v[i+63:i+32]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+95:i+64]
//				dst[i+127:i+96] := v[i+95:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+127:i+96]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+127:i+96]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_epi32'.
// Requires KNCNI.
func SwizzleEpi32(v x86.M512i, s MMSWIZZLEENUM) x86.M512i {
	return x86.M512i(swizzleEpi32([64]byte(v), s))
}

func swizzleEpi32(v [64]byte, s MMSWIZZLEENUM) [64]byte


// MaskSwizzleEpi64: Performs a swizzle transformation of each of the four
// groups of packed 4x64-bit integer elements in 'v' using swizzle parameter
// 's', storing the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				IF k[j*2]
//					dst[i+63:i]	 := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	 := src[i+63:i]
//				FI
//				IF k[j*2+1]
//					dst[i+127:i+64] := v[i+63:i]
//				ELSE
//					dst[i+127:i+64] := src[i+127:i+64]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+64]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+63:i]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+63:i]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+63:i]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+63]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+127:i+63]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+127:i+63]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+63]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+191:i+128]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+191:i+128]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+255:i+192]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+255:i+192]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_epi64'.
// Requires KNCNI.
func MaskSwizzleEpi64(src x86.M512i, k x86.Mmask8, v x86.M512i, s MMSWIZZLEENUM) x86.M512i {
	return x86.M512i(maskSwizzleEpi64([64]byte(src), uint8(k), [64]byte(v), s))
}

func maskSwizzleEpi64(src [64]byte, k uint8, v [64]byte, s MMSWIZZLEENUM) [64]byte


// SwizzleEpi64: Performs a swizzle transformation of each of the two groups of
// packed 4x64-bit integer elements in 'v' using swizzle parameter 's', storing
// the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				dst[i+63:i]	    := v[i+127:i+64]
//				dst[i+127:i+64] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+127:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+63:i]
//				dst[i+127:i+64]  := v[i+63:i]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+127:i+63]
//				dst[i+127:i+64]  := v[i+127:i+63]
//				dst[i+191:i+128] := v[i+127:i+63]
//				dst[i+255:i+192] := v[i+127:i+63]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+191:i+128]
//				dst[i+255:i+192] := v[i+191:i+128]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+255:i+192]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+255:i+192]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+127:i+64]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_epi64'.
// Requires KNCNI.
func SwizzleEpi64(v x86.M512i, s MMSWIZZLEENUM) x86.M512i {
	return x86.M512i(swizzleEpi64([64]byte(v), s))
}

func swizzleEpi64(v [64]byte, s MMSWIZZLEENUM) [64]byte


// MaskSwizzlePd: Performs a swizzle transformation of each of the two groups
// of packed 4x double-precision (64-bit) floating-point elements in 'v' using
// swizzle parameter 's', storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				IF k[j*2]
//					dst[i+63:i]	 := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	 := src[i+63:i]
//				FI
//				IF k[j*2+1]
//					dst[i+127:i+64] := v[i+63:i]
//				ELSE
//					dst[i+127:i+64] := src[i+127:i+64]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+64]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+63:i]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+63:i]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+63:i]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+63]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+127:i+63]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+127:i+63]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+63]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+191:i+128]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+191:i+128]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+255:i+192]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+255:i+192]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_pd'.
// Requires KNCNI.
func MaskSwizzlePd(src x86.M512d, k x86.Mmask8, v x86.M512d, s MMSWIZZLEENUM) x86.M512d {
	return x86.M512d(maskSwizzlePd([8]float64(src), uint8(k), [8]float64(v), s))
}

func maskSwizzlePd(src [8]float64, k uint8, v [8]float64, s MMSWIZZLEENUM) [8]float64


// SwizzlePd: Performs a swizzle transformation of each of the two groups of
// packed 4x double-precision (64-bit) floating-point elements in 'v' using
// swizzle parameter 's', storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				dst[i+63:i]     := v[i+127:i+64]
//				dst[i+127:i+64] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+127:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+63:i]
//				dst[i+127:i+64]  := v[i+63:i]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+127:i+63]
//				dst[i+127:i+64]  := v[i+127:i+63]
//				dst[i+191:i+128] := v[i+127:i+63]
//				dst[i+255:i+192] := v[i+127:i+63]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+191:i+128]
//				dst[i+255:i+192] := v[i+191:i+128]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+255:i+192]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+255:i+192]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+127:i+64]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_pd'.
// Requires KNCNI.
func SwizzlePd(v x86.M512d, s MMSWIZZLEENUM) x86.M512d {
	return x86.M512d(swizzlePd([8]float64(v), s))
}

func swizzlePd(v [8]float64, s MMSWIZZLEENUM) [8]float64


// MaskSwizzlePs: Performs a swizzle transformation of each of the four groups
// of packed 4x single-precision (32-bit) floating-point elements in 'v' using
// swizzle parameter 's', storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				IF k[j*2]
//					dst[i+31:i]	:= v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	:= src[i+31:i]
//				FI
//				IF k[j*2+1]
//					dst[i+63:i+32] := v[i+31:i]
//				ELSE
//					dst[i+63:i+32] := src[i+63:i+32]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+31:i]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+31:i]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+31:i]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+63:i+32]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+63:i+32]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+95:i+64]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+95:i+64]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+127:i+96]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+127:i+96]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_ps'.
// Requires KNCNI.
func MaskSwizzlePs(src x86.M512, k x86.Mmask16, v x86.M512, s MMSWIZZLEENUM) x86.M512 {
	return x86.M512(maskSwizzlePs([16]float32(src), uint16(k), [16]float32(v), s))
}

func maskSwizzlePs(src [16]float32, k uint16, v [16]float32, s MMSWIZZLEENUM) [16]float32


// SwizzlePs: Performs a swizzle transformation of each of the four groups of
// packed 4xsingle-precision (32-bit) floating-point elements in 'v' using
// swizzle parameter 's', storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				dst[i+31:i]    := v[i+63:i+32]
//				dst[i+63:i+32] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+31:i]
//				dst[i+63:i+32]  := v[i+31:i]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+63:i+32]
//				dst[i+95:i+64]  := v[i+63:i+32]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+95:i+64]
//				dst[i+127:i+96] := v[i+95:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+127:i+96]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+127:i+96]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_ps'.
// Requires KNCNI.
func SwizzlePs(v x86.M512, s MMSWIZZLEENUM) x86.M512 {
	return x86.M512(swizzlePs([16]float32(v), s))
}

func swizzlePs(v [16]float32, s MMSWIZZLEENUM) [16]float32


// MaskTestEpi32Mask: Compute the bitwise AND of packed 32-bit integers in 'a'
// and 'b', producing intermediate 32-bit values, and set the corresponding bit
// in result mask 'k' (subject to writemask 'k') if the intermediate value is
// non-zero. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ((a[i+31:i] AND b[i+31:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMD'. Intrinsic: '_mm512_mask_test_epi32_mask'.
// Requires KNCNI.
func MaskTestEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(maskTestEpi32Mask(uint16(k1), [64]byte(a), [64]byte(b)))
}

func maskTestEpi32Mask(k1 uint16, a [64]byte, b [64]byte) uint16


// TestEpi32Mask: Compute the bitwise AND of packed 32-bit integers in 'a' and
// 'b', producing intermediate 32-bit values, and set the corresponding bit in
// result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ((a[i+31:i] AND b[i+31:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMD'. Intrinsic: '_mm512_test_epi32_mask'.
// Requires KNCNI.
func TestEpi32Mask(a x86.M512i, b x86.M512i) x86.Mmask16 {
	return x86.Mmask16(testEpi32Mask([64]byte(a), [64]byte(b)))
}

func testEpi32Mask(a [64]byte, b [64]byte) uint16


// Tzcnti32: Counts the number of trailing bits in unsigned 32-bit integer 'x'
// starting at bit 'a' storing the result in 'dst'. 
//
//		count := 0
//		FOR j := a to 31
//			IF NOT(x[j]  1)
//				count := count + 1
//			FI
//		ENDFOR
//		dst := count
//
// Instruction: 'TZCNTI'. Intrinsic: '_mm_tzcnti_32'.
// Requires KNCNI.
func Tzcnti32(a int, x uint32) int {
	return int(tzcnti32(a, x))
}

func tzcnti32(a int, x uint32) int


// Tzcnti64: Counts the number of trailing bits in unsigned 64-bit integer 'x'
// starting at bit 'a' storing the result in 'dst'. 
//
//		count := 0
//		FOR j := a to 63
//			IF NOT(x[j]  1)
//				count := count + 1
//			FI
//		ENDFOR
//		dst := count
//
// Instruction: 'TZCNTI'. Intrinsic: '_mm_tzcnti_64'.
// Requires KNCNI.
func Tzcnti64(a int64, x uint64) int64 {
	return int64(tzcnti64(a, x))
}

func tzcnti64(a int64, x uint64) int64


// MaskXorEpi32: Compute the bitwise XOR of packed 32-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
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
// Instruction: 'VPXORD'. Intrinsic: '_mm512_mask_xor_epi32'.
// Requires KNCNI.
func MaskXorEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskXorEpi32([64]byte(src), uint16(k), [64]byte(a), [64]byte(b)))
}

func maskXorEpi32(src [64]byte, k uint16, a [64]byte, b [64]byte) [64]byte


// XorEpi32: Compute the bitwise XOR of packed 32-bit integers in 'a' and 'b',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORD'. Intrinsic: '_mm512_xor_epi32'.
// Requires KNCNI.
func XorEpi32(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(xorEpi32([64]byte(a), [64]byte(b)))
}

func xorEpi32(a [64]byte, b [64]byte) [64]byte


// MaskXorEpi64: Compute the bitwise XOR of packed 64-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
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
// Instruction: 'VPXORQ'. Intrinsic: '_mm512_mask_xor_epi64'.
// Requires KNCNI.
func MaskXorEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(maskXorEpi64([64]byte(src), uint8(k), [64]byte(a), [64]byte(b)))
}

func maskXorEpi64(src [64]byte, k uint8, a [64]byte, b [64]byte) [64]byte


// XorEpi64: Compute the bitwise XOR of packed 64-bit integers in 'a' and 'b',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORQ'. Intrinsic: '_mm512_xor_epi64'.
// Requires KNCNI.
func XorEpi64(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(xorEpi64([64]byte(a), [64]byte(b)))
}

func xorEpi64(a [64]byte, b [64]byte) [64]byte


// XorSi512: Compute the bitwise XOR of 512 bits (representing integer data) in
// 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := (a[511:0] XOR b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORD'. Intrinsic: '_mm512_xor_si512'.
// Requires KNCNI.
func XorSi512(a x86.M512i, b x86.M512i) x86.M512i {
	return x86.M512i(xorSi512([64]byte(a), [64]byte(b)))
}

func xorSi512(a [64]byte, b [64]byte) [64]byte

