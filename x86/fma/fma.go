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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}


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
	panic("not implemented")
}

