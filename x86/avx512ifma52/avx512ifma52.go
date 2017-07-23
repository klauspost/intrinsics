// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package avx512ifma52

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// M512Madd52hiEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the high
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//			dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm512_madd52hi_epu64'.
// Requires AVX512IFMA52.
func M512Madd52hiEpu64(a x86.M512i, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMadd52hiEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// high 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm512_mask_madd52hi_epu64'.
// Requires AVX512IFMA52.
func M512MaskMadd52hiEpu64(a x86.M512i, k x86.Mmask8, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMadd52hiEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// high 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[103:52])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD52HUQ'. Intrinsic: '_mm512_maskz_madd52hi_epu64'.
// Requires AVX512IFMA52.
func M512MaskzMadd52hiEpu64(k x86.Mmask8, a x86.M512i, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512Madd52loEpu64: Multiply packed unsigned 52-bit integers in each 64-bit
// element of 'b' and 'c' to form a 104-bit intermediate result. Add the low
// 52-bit unsigned integer from the intermediate result with the corresponding
// unsigned 64-bit integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//			dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm512_madd52lo_epu64'.
// Requires AVX512IFMA52.
func M512Madd52loEpu64(a x86.M512i, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMadd52loEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// low 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'a' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm512_mask_madd52lo_epu64'.
// Requires AVX512IFMA52.
func M512MaskMadd52loEpu64(a x86.M512i, k x86.Mmask8, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzMadd52loEpu64: Multiply packed unsigned 52-bit integers in each
// 64-bit element of 'b' and 'c' to form a 104-bit intermediate result. Add the
// low 52-bit unsigned integer from the intermediate result with the
// corresponding unsigned 64-bit integer in 'a', and store the results in 'dst'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp[127:0] := ZeroExtend64(b[i+51:i]) * ZeroExtend64(c[i+51:i])
//				dst[i+63:i] := a[i+63:i] + ZeroExtend64(tmp[51:0])
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD52LUQ'. Intrinsic: '_mm512_maskz_madd52lo_epu64'.
// Requires AVX512IFMA52.
func M512MaskzMadd52loEpu64(k x86.Mmask8, a x86.M512i, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}

