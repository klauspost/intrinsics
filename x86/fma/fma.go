package fma

import . "github.com/klauspost/intrinsics/x86"


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
func FmaddPd(a M128d, b M128d, c M128d) M128d {
	return M128d(fmaddPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FmaddPd1: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', add the intermediate result to packed elements in 'c', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm256_fmadd_pd'.
// Requires FMA.
func FmaddPd1(a M256d, b M256d, c M256d) M256d {
	return M256d(fmaddPd1([4]float64(a), [4]float64(b), [4]float64(c)))
}

func fmaddPd1(a [4]float64, b [4]float64, c [4]float64) [4]float64


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
func FmaddPs(a M128, b M128, c M128) M128 {
	return M128(fmaddPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FmaddPs1: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', add the intermediate result to packed elements in 'c', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm256_fmadd_ps'.
// Requires FMA.
func FmaddPs1(a M256, b M256, c M256) M256 {
	return M256(fmaddPs1([8]float32(a), [8]float32(b), [8]float32(c)))
}

func fmaddPs1(a [8]float32, b [8]float32, c [8]float32) [8]float32


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
func FmaddSd(a M128d, b M128d, c M128d) M128d {
	return M128d(fmaddSd([2]float64(a), [2]float64(b), [2]float64(c)))
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
func FmaddSs(a M128, b M128, c M128) M128 {
	return M128(fmaddSs([4]float32(a), [4]float32(b), [4]float32(c)))
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
func FmaddsubPd(a M128d, b M128d, c M128d) M128d {
	return M128d(fmaddsubPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmaddsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FmaddsubPd1: Multiply packed double-precision (64-bit) floating-point
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
func FmaddsubPd1(a M256d, b M256d, c M256d) M256d {
	return M256d(fmaddsubPd1([4]float64(a), [4]float64(b), [4]float64(c)))
}

func fmaddsubPd1(a [4]float64, b [4]float64, c [4]float64) [4]float64


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
func FmaddsubPs(a M128, b M128, c M128) M128 {
	return M128(fmaddsubPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmaddsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FmaddsubPs1: Multiply packed single-precision (32-bit) floating-point
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
func FmaddsubPs1(a M256, b M256, c M256) M256 {
	return M256(fmaddsubPs1([8]float32(a), [8]float32(b), [8]float32(c)))
}

func fmaddsubPs1(a [8]float32, b [8]float32, c [8]float32) [8]float32


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
func FmsubPd(a M128d, b M128d, c M128d) M128d {
	return M128d(fmsubPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FmsubPd1: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the intermediate
// result, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm256_fmsub_pd'.
// Requires FMA.
func FmsubPd1(a M256d, b M256d, c M256d) M256d {
	return M256d(fmsubPd1([4]float64(a), [4]float64(b), [4]float64(c)))
}

func fmsubPd1(a [4]float64, b [4]float64, c [4]float64) [4]float64


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
func FmsubPs(a M128, b M128, c M128) M128 {
	return M128(fmsubPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FmsubPs1: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the intermediate
// result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm256_fmsub_ps'.
// Requires FMA.
func FmsubPs1(a M256, b M256, c M256) M256 {
	return M256(fmsubPs1([8]float32(a), [8]float32(b), [8]float32(c)))
}

func fmsubPs1(a [8]float32, b [8]float32, c [8]float32) [8]float32


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
func FmsubSd(a M128d, b M128d, c M128d) M128d {
	return M128d(fmsubSd([2]float64(a), [2]float64(b), [2]float64(c)))
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
func FmsubSs(a M128, b M128, c M128) M128 {
	return M128(fmsubSs([4]float32(a), [4]float32(b), [4]float32(c)))
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
func FmsubaddPd(a M128d, b M128d, c M128d) M128d {
	return M128d(fmsubaddPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fmsubaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FmsubaddPd1: Multiply packed double-precision (64-bit) floating-point
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
func FmsubaddPd1(a M256d, b M256d, c M256d) M256d {
	return M256d(fmsubaddPd1([4]float64(a), [4]float64(b), [4]float64(c)))
}

func fmsubaddPd1(a [4]float64, b [4]float64, c [4]float64) [4]float64


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
func FmsubaddPs(a M128, b M128, c M128) M128 {
	return M128(fmsubaddPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fmsubaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FmsubaddPs1: Multiply packed single-precision (32-bit) floating-point
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
func FmsubaddPs1(a M256, b M256, c M256) M256 {
	return M256(fmsubaddPs1([8]float32(a), [8]float32(b), [8]float32(c)))
}

func fmsubaddPs1(a [8]float32, b [8]float32, c [8]float32) [8]float32


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
func FnmaddPd(a M128d, b M128d, c M128d) M128d {
	return M128d(fnmaddPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fnmaddPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FnmaddPd1: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', add the negated intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:256] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm256_fnmadd_pd'.
// Requires FMA.
func FnmaddPd1(a M256d, b M256d, c M256d) M256d {
	return M256d(fnmaddPd1([4]float64(a), [4]float64(b), [4]float64(c)))
}

func fnmaddPd1(a [4]float64, b [4]float64, c [4]float64) [4]float64


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
func FnmaddPs(a M128, b M128, c M128) M128 {
	return M128(fnmaddPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmaddPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FnmaddPs1: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', add the negated intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			a[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:256] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm256_fnmadd_ps'.
// Requires FMA.
func FnmaddPs1(a M256, b M256, c M256) M256 {
	return M256(fnmaddPs1([8]float32(a), [8]float32(b), [8]float32(c)))
}

func fnmaddPs1(a [8]float32, b [8]float32, c [8]float32) [8]float32


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
func FnmaddSd(a M128d, b M128d, c M128d) M128d {
	return M128d(fnmaddSd([2]float64(a), [2]float64(b), [2]float64(c)))
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
func FnmaddSs(a M128, b M128, c M128) M128 {
	return M128(fnmaddSs([4]float32(a), [4]float32(b), [4]float32(c)))
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
func FnmsubPd(a M128d, b M128d, c M128d) M128d {
	return M128d(fnmsubPd([2]float64(a), [2]float64(b), [2]float64(c)))
}

func fnmsubPd(a [2]float64, b [2]float64, c [2]float64) [2]float64


// FnmsubPd1: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the negated
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
func FnmsubPd1(a M256d, b M256d, c M256d) M256d {
	return M256d(fnmsubPd1([4]float64(a), [4]float64(b), [4]float64(c)))
}

func fnmsubPd1(a [4]float64, b [4]float64, c [4]float64) [4]float64


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
func FnmsubPs(a M128, b M128, c M128) M128 {
	return M128(fnmsubPs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmsubPs(a [4]float32, b [4]float32, c [4]float32) [4]float32


// FnmsubPs1: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', subtract packed elements in 'c' from the negated
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
func FnmsubPs1(a M256, b M256, c M256) M256 {
	return M256(fnmsubPs1([8]float32(a), [8]float32(b), [8]float32(c)))
}

func fnmsubPs1(a [8]float32, b [8]float32, c [8]float32) [8]float32


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
func FnmsubSd(a M128d, b M128d, c M128d) M128d {
	return M128d(fnmsubSd([2]float64(a), [2]float64(b), [2]float64(c)))
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
func FnmsubSs(a M128, b M128, c M128) M128 {
	return M128(fnmsubSs([4]float32(a), [4]float32(b), [4]float32(c)))
}

func fnmsubSs(a [4]float32, b [4]float32, c [4]float32) [4]float32

