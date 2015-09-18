// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package fma

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// FmaddPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', add the intermediate result to packed elements in 'c', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm_fmadd_pd'.
// Requires FMA.
func FmaddPd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fmaddPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// M256FmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm256_fmadd_pd'.
// Requires FMA.
func M256FmaddPd(a x86.M256d, b x86.M256d, c x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FmaddPd([4]float64(a), [4]float64(b), [4]float64(c)))
}

func m256FmaddPd(a [4]float64, b [4]float64, c [4]float64) [4]float64


// FmaddPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', add the intermediate result to packed elements in 'c', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm_fmadd_ps'.
// Requires FMA.
func FmaddPs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fmaddPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// M256FmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm256_fmadd_ps'.
// Requires FMA.
func M256FmaddPs(a x86.M256, b x86.M256, c x86.M256) (dst x86.M256) {
	return x86.M256(m256FmaddPs([8]float32(a), [8]float32(b), [8]float32(c)))
}

func m256FmaddPs(a [8]float32, b [8]float32, c [8]float32) [8]float32


// FmaddSd: Multiply the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b', and add the intermediate result to the lower
// element in 'c'. Store the result in the lower element of 'dst', and copy the
// upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] * b[63:0]) + c[63:0]
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VFMADD132SD, VFMADD213SD, VFMADD231SD'. Intrinsic: '_mm_fmadd_sd'.
// Requires FMA.
func FmaddSd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fmaddSd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmaddSd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FmaddSs: Multiply the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b', and add the intermediate result to the lower
// element in 'c'. Store the result in the lower element of 'dst', and copy the
// upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := (a[31:0] * b[31:0]) + c[31:0]
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VFMADD132SS, VFMADD213SS, VFMADD231SS'. Intrinsic: '_mm_fmadd_ss'.
// Requires FMA.
func FmaddSs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fmaddSs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmaddSs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FmaddsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', alternatively add and subtract packed elements in
// 'c' to/from the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF (j is even) 
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMADDSUB132PD, VFMADDSUB213PD, VFMADDSUB231PD'. Intrinsic: '_mm_fmaddsub_pd'.
// Requires FMA.
func FmaddsubPd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fmaddsubPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmaddsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// M256FmaddsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', alternatively add and subtract packed elements in
// 'c' to/from the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF (j is even) 
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMADDSUB132PD, VFMADDSUB213PD, VFMADDSUB231PD'. Intrinsic: '_mm256_fmaddsub_pd'.
// Requires FMA.
func M256FmaddsubPd(a x86.M256d, b x86.M256d, c x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FmaddsubPd([4]float64(a), [4]float64(b), [4]float64(c)))
}

func m256FmaddsubPd(a [4]float64, b [4]float64, c [4]float64) [4]float64


// FmaddsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', alternatively add and subtract packed elements in
// 'c' to/from the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF (j is even) 
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMADDSUB132PS, VFMADDSUB213PS, VFMADDSUB231PS'. Intrinsic: '_mm_fmaddsub_ps'.
// Requires FMA.
func FmaddsubPs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fmaddsubPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmaddsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// M256FmaddsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', alternatively add and subtract packed elements in
// 'c' to/from the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF (j is even) 
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMADDSUB132PS, VFMADDSUB213PS, VFMADDSUB231PS'. Intrinsic: '_mm256_fmaddsub_ps'.
// Requires FMA.
func M256FmaddsubPs(a x86.M256, b x86.M256, c x86.M256) (dst x86.M256) {
	return x86.M256(m256FmaddsubPs([8]float32(a), [8]float32(b), [8]float32(c)))
}

func m256FmaddsubPs(a [8]float32, b [8]float32, c [8]float32) [8]float32


// FmsubPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the intermediate
// result, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm_fmsub_pd'.
// Requires FMA.
func FmsubPd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fmsubPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// M256FmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm256_fmsub_pd'.
// Requires FMA.
func M256FmsubPd(a x86.M256d, b x86.M256d, c x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FmsubPd([4]float64(a), [4]float64(b), [4]float64(c)))
}

func m256FmsubPd(a [4]float64, b [4]float64, c [4]float64) [4]float64


// FmsubPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the intermediate
// result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm_fmsub_ps'.
// Requires FMA.
func FmsubPs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fmsubPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// M256FmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm256_fmsub_ps'.
// Requires FMA.
func M256FmsubPs(a x86.M256, b x86.M256, c x86.M256) (dst x86.M256) {
	return x86.M256(m256FmsubPs([8]float32(a), [8]float32(b), [8]float32(c)))
}

func m256FmsubPs(a [8]float32, b [8]float32, c [8]float32) [8]float32


// FmsubSd: Multiply the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b', and subtract the lower element in 'c' from the
// intermediate result. Store the result in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := (a[63:0] * b[63:0]) - c[63:0]
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VFMSUB132SD, VFMSUB213SD, VFMSUB231SD'. Intrinsic: '_mm_fmsub_sd'.
// Requires FMA.
func FmsubSd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fmsubSd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmsubSd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FmsubSs: Multiply the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b', and subtract the lower element in 'c' from the
// intermediate result. Store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := (a[31:0] * b[31:0]) - c[31:0]
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VFMSUB132SS, VFMSUB213SS, VFMSUB231SS'. Intrinsic: '_mm_fmsub_ss'.
// Requires FMA.
func FmsubSs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fmsubSs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FmsubaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', alternatively subtract and add packed elements in
// 'c' from/to the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF (j is even) 
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMSUBADD132PD, VFMSUBADD213PD, VFMSUBADD231PD'. Intrinsic: '_mm_fmsubadd_pd'.
// Requires FMA.
func FmsubaddPd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fmsubaddPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmsubaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// M256FmsubaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', alternatively subtract and add packed elements in
// 'c' from/to the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF (j is even) 
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMSUBADD132PD, VFMSUBADD213PD, VFMSUBADD231PD'. Intrinsic: '_mm256_fmsubadd_pd'.
// Requires FMA.
func M256FmsubaddPd(a x86.M256d, b x86.M256d, c x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FmsubaddPd([4]float64(a), [4]float64(b), [4]float64(c)))
}

func m256FmsubaddPd(a [4]float64, b [4]float64, c [4]float64) [4]float64


// FmsubaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', alternatively subtract and add packed elements in
// 'c' from/to the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF (j is even) 
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VFMSUBADD132PS, VFMSUBADD213PS, VFMSUBADD231PS'. Intrinsic: '_mm_fmsubadd_ps'.
// Requires FMA.
func FmsubaddPs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fmsubaddPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmsubaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// M256FmsubaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', alternatively subtract and add packed elements in
// 'c' from/to the intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF (j is even) 
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMSUBADD132PS, VFMSUBADD213PS, VFMSUBADD231PS'. Intrinsic: '_mm256_fmsubadd_ps'.
// Requires FMA.
func M256FmsubaddPs(a x86.M256, b x86.M256, c x86.M256) (dst x86.M256) {
	return x86.M256(m256FmsubaddPs([8]float32(a), [8]float32(b), [8]float32(c)))
}

func m256FmsubaddPs(a [8]float32, b [8]float32, c [8]float32) [8]float32


// FnmaddPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', add the negated intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm_fnmadd_pd'.
// Requires FMA.
func FnmaddPd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fnmaddPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fnmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// M256FnmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:256] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm256_fnmadd_pd'.
// Requires FMA.
func M256FnmaddPd(a x86.M256d, b x86.M256d, c x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FnmaddPd([4]float64(a), [4]float64(b), [4]float64(c)))
}

func m256FnmaddPd(a [4]float64, b [4]float64, c [4]float64) [4]float64


// FnmaddPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', add the negated intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			a[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm_fnmadd_ps'.
// Requires FMA.
func FnmaddPs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fnmaddPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// M256FnmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			a[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:256] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm256_fnmadd_ps'.
// Requires FMA.
func M256FnmaddPs(a x86.M256, b x86.M256, c x86.M256) (dst x86.M256) {
	return x86.M256(m256FnmaddPs([8]float32(a), [8]float32(b), [8]float32(c)))
}

func m256FnmaddPs(a [8]float32, b [8]float32, c [8]float32) [8]float32


// FnmaddSd: Multiply the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b', and add the negated intermediate result to the
// lower element in 'c'. Store the result in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := -(a[63:0] * b[63:0]) + c[63:0]
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMADD132SD, VFNMADD213SD, VFNMADD231SD'. Intrinsic: '_mm_fnmadd_sd'.
// Requires FMA.
func FnmaddSd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fnmaddSd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fnmaddSd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FnmaddSs: Multiply the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b', and add the negated intermediate result to the
// lower element in 'c'. Store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := -(a[31:0] * b[31:0]) + c[31:0]
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMADD132SS, VFNMADD213SS, VFNMADD231SS'. Intrinsic: '_mm_fnmadd_ss'.
// Requires FMA.
func FnmaddSs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fnmaddSs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmaddSs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FnmsubPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR	
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm_fnmsub_pd'.
// Requires FMA.
func FnmsubPd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fnmsubPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fnmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// M256FnmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR	
//		dst[MAX:256] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm256_fnmsub_pd'.
// Requires FMA.
func M256FnmsubPd(a x86.M256d, b x86.M256d, c x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FnmsubPd([4]float64(a), [4]float64(b), [4]float64(c)))
}

func m256FnmsubPd(a [4]float64, b [4]float64, c [4]float64) [4]float64


// FnmsubPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR	
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm_fnmsub_ps'.
// Requires FMA.
func FnmsubPs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fnmsubPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// M256FnmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR	
//		dst[MAX:256] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm256_fnmsub_ps'.
// Requires FMA.
func M256FnmsubPs(a x86.M256, b x86.M256, c x86.M256) (dst x86.M256) {
	return x86.M256(m256FnmsubPs([8]float32(a), [8]float32(b), [8]float32(c)))
}

func m256FnmsubPs(a [8]float32, b [8]float32, c [8]float32) [8]float32


// FnmsubSd: Multiply the lower double-precision (64-bit) floating-point
// elements in 'a' and 'b', and subtract the lower element in 'c' from the
// negated intermediate result. Store the result in the lower element of 'dst',
// and copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := -(a[63:0] * b[63:0]) - c[63:0]
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMSUB132SD, VFNMSUB213SD, VFNMSUB231SD'. Intrinsic: '_mm_fnmsub_sd'.
// Requires FMA.
func FnmsubSd(a x86.M128d, b x86.M128d, c x86.M128d) (dst x86.M128d) {
	return x86.M128d(fnmsubSd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fnmsubSd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FnmsubSs: Multiply the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b', and subtract the lower element in 'c' from the
// negated intermediate result. Store the result in the lower element of 'dst',
// and copy the upper 3 packed elements from 'a' to the upper elements of
// 'dst'. 
//
//		dst[31:0] := -(a[31:0] * b[31:0]) - c[31:0]
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VFNMSUB132SS, VFNMSUB213SS, VFNMSUB231SS'. Intrinsic: '_mm_fnmsub_ss'.
// Requires FMA.
func FnmsubSs(a x86.M128, b x86.M128, c x86.M128) (dst x86.M128) {
	return x86.M128(fnmsubSs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32

