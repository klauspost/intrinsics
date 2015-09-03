package fp16c

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// CvtphPs: Convert packed half-precision (16-bit) floating-point elements in
// 'a' to packed single-precision (32-bit) floating-point elements, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			m := j*16
//			dst[i+31:i] := Convert_FP16_To_FP32(a[m+15:m])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPH2PS'. Intrinsic: '_mm_cvtph_ps'.
// Requires FP16C.
func CvtphPs(a x86.M128i) (dst x86.M128) {
	return x86.M128(cvtphPs([16]byte(a)))
}

func cvtphPs(a [16]byte) [4]float32


// M256CvtphPs: Convert packed half-precision (16-bit) floating-point elements
// in 'a' to packed single-precision (32-bit) floating-point elements, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			m := j*16
//			dst[i+31:i] := Convert_FP16_To_FP32(a[m+15:m])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPH2PS'. Intrinsic: '_mm256_cvtph_ps'.
// Requires FP16C.
func M256CvtphPs(a x86.M128i) (dst x86.M256) {
	return x86.M256(m256CvtphPs([16]byte(a)))
}

func m256CvtphPs(a [16]byte) [8]float32


// CvtpsPh: Convert packed single-precision (32-bit) floating-point elements in
// 'a' to packed half-precision (16-bit) floating-point elements, and store the
// results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 3
//			i := 16*j
//			l := 32*j
//			dst[i+15:i] := Convert_FP32_To_FP16FP(a[l+31:l])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2PH'. Intrinsic: '_mm_cvtps_ph'.
// Requires FP16C.
func CvtpsPh(a x86.M128, rounding int) (dst x86.M128i) {
	return x86.M128i(cvtpsPh([4]float32(a), rounding))
}

func cvtpsPh(a [4]float32, rounding int) [16]byte


// M256CvtpsPh: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed half-precision (16-bit) floating-point elements,
// and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := 16*j
//			l := 32*j
//			dst[i+15:i] := Convert_FP32_To_FP16FP(a[l+31:l])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPS2PH'. Intrinsic: '_mm256_cvtps_ph'.
// Requires FP16C.
func M256CvtpsPh(a x86.M256, rounding int) (dst x86.M128i) {
	return x86.M128i(m256CvtpsPh([8]float32(a), rounding))
}

func m256CvtpsPh(a [8]float32, rounding int) [16]byte

