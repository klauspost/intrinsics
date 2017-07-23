// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package sse

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AcosPd: Compute the inverse cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ACOS(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_acos_pd'.
// Requires SSE.
func AcosPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// AcosPs: Compute the inverse cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ACOS(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_acos_ps'.
// Requires SSE.
func AcosPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AcoshPd: Compute the inverse hyperbolic cosine of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ACOSH(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_acosh_pd'.
// Requires SSE.
func AcoshPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// AcoshPs: Compute the inverse hyperbolic cosine of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ACOSH(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_acosh_ps'.
// Requires SSE.
func AcoshPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AddPs: Add packed single-precision (32-bit) floating-point elements in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//
// Instruction: 'ADDPS'. Intrinsic: '_mm_add_ps'.
// Requires SSE.
func AddPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AddSs: Add the lower single-precision (32-bit) floating-point element in 'a'
// and 'b', store the result in the lower element of 'dst', and copy the upper
// 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0] + b[31:0]
//		dst[127:32] := a[127:32]
//
// Instruction: 'ADDSS'. Intrinsic: '_mm_add_ss'.
// Requires SSE.
func AddSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AndPs: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//		ENDFOR
//
// Instruction: 'ANDPS'. Intrinsic: '_mm_and_ps'.
// Requires SSE.
func AndPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AndnotPs: Compute the bitwise AND NOT of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//		ENDFOR
//
// Instruction: 'ANDNPS'. Intrinsic: '_mm_andnot_ps'.
// Requires SSE.
func AndnotPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AsinPd: Compute the inverse sine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ASIN(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_asin_pd'.
// Requires SSE.
func AsinPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// AsinPs: Compute the inverse sine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ASIN(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_asin_ps'.
// Requires SSE.
func AsinPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AsinhPd: Compute the inverse hyperbolic sine of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ASINH(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_asinh_pd'.
// Requires SSE.
func AsinhPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// AsinhPs: Compute the inverse hyperbolic sine of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ASINH(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_asinh_ps'.
// Requires SSE.
func AsinhPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AtanPd: Compute the inverse tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ATAN(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_atan_pd'.
// Requires SSE.
func AtanPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// AtanPs: Compute the inverse tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ATAN(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_atan_ps'.
// Requires SSE.
func AtanPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Atan2Pd: Compute the inverse tangent of packed double-precision (64-bit)
// floating-point elements in 'a' divided by packed elements in 'b', and store
// the results in 'dst' expressed in radians. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ATAN(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_atan2_pd'.
// Requires SSE.
func Atan2Pd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Atan2Ps: Compute the inverse tangent of packed single-precision (32-bit)
// floating-point elements in 'a' divided by packed elements in 'b', and store
// the results in 'dst' expressed in radians. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ATAN(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_atan2_ps'.
// Requires SSE.
func Atan2Ps(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AtanhPd: Compute the inverse hyperbolic tangent of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ATANH(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_atanh_pd'.
// Requires SSE.
func AtanhPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// AtanhPs: Compute the inverse hyperbolic tangent of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ATANH(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_atanh_ps'.
// Requires SSE.
func AtanhPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// AvgPu16: Average packed unsigned 16-bit integers in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGW'. Intrinsic: '_mm_avg_pu16'.
// Requires SSE.
func AvgPu16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// AvgPu8: Average packed unsigned 8-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGB'. Intrinsic: '_mm_avg_pu8'.
// Requires SSE.
func AvgPu8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// CbrtPd: Compute the cube root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := CubeRoot(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cbrt_pd'.
// Requires SSE.
func CbrtPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// CbrtPs: Compute the cube root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := CubeRoot(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cbrt_ps'.
// Requires SSE.
func CbrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CdfnormPd: Compute the cumulative distribution function of packed
// double-precision (64-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := CDFNormal(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cdfnorm_pd'.
// Requires SSE.
func CdfnormPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// CdfnormPs: Compute the cumulative distribution function of packed
// single-precision (32-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := CDFNormal(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cdfnorm_ps'.
// Requires SSE.
func CdfnormPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CdfnorminvPd: Compute the inverse cumulative distribution function of packed
// double-precision (64-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := InverseCDFNormal(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cdfnorminv_pd'.
// Requires SSE.
func CdfnorminvPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// CdfnorminvPs: Compute the inverse cumulative distribution function of packed
// single-precision (32-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := InverseCDFNormal(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cdfnorminv_ps'.
// Requires SSE.
func CdfnorminvPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CexpPs: Compute the exponential value of 'e' raised to the power of packed
// complex single-precision (32-bit) floating-point elements in 'a', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cexp_ps'.
// Requires SSE.
func CexpPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ClogPs: Compute the natural logarithm of packed complex single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ln(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_clog_ps'.
// Requires SSE.
func ClogPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpeqPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' for equality, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] == b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpeq_ps'.
// Requires SSE.
func CmpeqPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpeqSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b' for equality, store the result in the lower element of 'dst',
// and copy the upper 3 packed elements from 'a' to the upper elements of
// 'dst'. 
//
//		dst[31:0] := ( a[31:0] == b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpeq_ss'.
// Requires SSE.
func CmpeqSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpgePs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' for greater-than-or-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] >= b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpge_ps'.
// Requires SSE.
func CmpgePs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpgeSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b' for greater-than-or-equal, store the result in the lower
// element of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := ( a[31:0] >= b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpge_ss'.
// Requires SSE.
func CmpgeSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpgtPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' for greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] > b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpgt_ps'.
// Requires SSE.
func CmpgtPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpgtSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b' for greater-than, store the result in the lower element of
// 'dst', and copy the upper 3 packed elements from 'a' to the upper elements
// of 'dst'. 
//
//		dst[31:0] := ( a[31:0] > b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpgt_ss'.
// Requires SSE.
func CmpgtSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmplePs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' for less-than-or-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] <= b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmple_ps'.
// Requires SSE.
func CmplePs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpleSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b' for less-than-or-equal, store the result in the lower element
// of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := ( a[31:0] <= b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmple_ss'.
// Requires SSE.
func CmpleSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpltPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' for less-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] < b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmplt_ps'.
// Requires SSE.
func CmpltPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpltSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b' for less-than, store the result in the lower element of
// 'dst', and copy the upper 3 packed elements from 'a' to the upper elements
// of 'dst'. 
//
//		dst[31:0] := ( a[31:0] < b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmplt_ss'.
// Requires SSE.
func CmpltSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpneqPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' for not-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] != b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpneq_ps'.
// Requires SSE.
func CmpneqPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpneqSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-equal, store the result in the lower element
// of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := ( a[31:0] != b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpneq_ss'.
// Requires SSE.
func CmpneqSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpngePs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' for not-greater-than-or-equal, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := !( a[i+31:i] >= b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpnge_ps'.
// Requires SSE.
func CmpngePs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpngeSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-greater-than-or-equal, store the result in
// the lower element of 'dst', and copy the upper 3 packed elements from 'a' to
// the upper elements of 'dst'. 
//
//		dst[31:0] := !( a[31:0] >= b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpnge_ss'.
// Requires SSE.
func CmpngeSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpngtPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' for not-greater-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := !( a[i+31:i] > b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpngt_ps'.
// Requires SSE.
func CmpngtPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpngtSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-greater-than, store the result in the lower
// element of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := !( a[31:0] > b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpngt_ss'.
// Requires SSE.
func CmpngtSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpnlePs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' for not-less-than-or-equal, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := !( a[i+31:i] <= b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpnle_ps'.
// Requires SSE.
func CmpnlePs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpnleSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, store the result in the
// lower element of 'dst', and copy the upper 3 packed elements from 'a' to the
// upper elements of 'dst'. 
//
//		dst[31:0] := !( a[31:0] <= b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpnle_ss'.
// Requires SSE.
func CmpnleSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpnltPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' for not-less-than, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := !( a[i+31:i] < b[i+31:i] ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpnlt_ps'.
// Requires SSE.
func CmpnltPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpnltSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than, store the result in the lower
// element of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := !( a[31:0] < b[31:0] ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpnlt_ss'.
// Requires SSE.
func CmpnltSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpordPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' to see if neither is NaN, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] != NaN AND b[i+31:i] != NaN ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpord_ps'.
// Requires SSE.
func CmpordPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpordSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, store the result in the
// lower element of 'dst', and copy the upper 3 packed elements from 'a' to the
// upper elements of 'dst'. 
//
//		dst[31:0] := ( a[31:0] != NaN AND b[31:0] != NaN ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpord_ss'.
// Requires SSE.
func CmpordSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpunordPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' to see if either is NaN, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] != NaN OR b[i+31:i] != NaN ) ? 0xffffffff : 0
//		ENDFOR
//
// Instruction: 'CMPPS'. Intrinsic: '_mm_cmpunord_ps'.
// Requires SSE.
func CmpunordPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CmpunordSs: Compare the lower single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, store the result in the
// lower element of 'dst', and copy the upper 3 packed elements from 'a' to the
// upper elements of 'dst'. 
//
//		dst[31:0] := ( a[31:0] != NaN OR b[31:0] != NaN ) ? 0xffffffff : 0
//		dst[127:32] := a[127:32]
//
// Instruction: 'CMPSS'. Intrinsic: '_mm_cmpunord_ss'.
// Requires SSE.
func CmpunordSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ComieqSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for equality, and return the boolean result (0 or 1). 
//
//		RETURN ( a[31:0] == b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comieq_ss'.
// Requires SSE.
func ComieqSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// ComigeSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for greater-than-or-equal, and return the boolean result (0
// or 1). 
//
//		RETURN ( a[31:0] >= b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comige_ss'.
// Requires SSE.
func ComigeSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// ComigtSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for greater-than, and return the boolean result (0 or 1). 
//
//		RETURN ( a[31:0] > b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comigt_ss'.
// Requires SSE.
func ComigtSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// ComileSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for less-than-or-equal, and return the boolean result (0 or
// 1). 
//
//		RETURN ( a[31:0] <= b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comile_ss'.
// Requires SSE.
func ComileSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// ComiltSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for less-than, and return the boolean result (0 or 1). 
//
//		RETURN ( a[31:0] < b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comilt_ss'.
// Requires SSE.
func ComiltSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// ComineqSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for not-equal, and return the boolean result (0 or
// 1). 
//
//		RETURN ( a[31:0] != b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comineq_ss'.
// Requires SSE.
func ComineqSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// CosPd: Compute the cosine of packed double-precision (64-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := COS(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cos_pd'.
// Requires SSE.
func CosPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// CosPs: Compute the cosine of packed single-precision (32-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := COS(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cos_ps'.
// Requires SSE.
func CosPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CosdPd: Compute the cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := COSD(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cosd_pd'.
// Requires SSE.
func CosdPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// CosdPs: Compute the cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := COSD(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cosd_ps'.
// Requires SSE.
func CosdPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CoshPd: Compute the hyperbolic cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := COSH(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cosh_pd'.
// Requires SSE.
func CoshPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// CoshPs: Compute the hyperbolic cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := COSH(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_cosh_ps'.
// Requires SSE.
func CoshPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CsqrtPs: Compute the square root of packed complex single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_csqrt_ps'.
// Requires SSE.
func CsqrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// CvtPi2ps: Convert packed 32-bit integers in 'b' to packed single-precision
// (32-bit) floating-point elements, store the results in the lower 2 elements
// of 'dst', and copy the upper 2 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[63:32] := Convert_Int32_To_FP32(b[63:32])
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'CVTPI2PS'. Intrinsic: '_mm_cvt_pi2ps'.
// Requires SSE.
func CvtPi2ps(a x86.M128, b x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// CvtPs2pi: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTPS2PI'. Intrinsic: '_mm_cvt_ps2pi'.
// Requires SSE.
func CvtPs2pi(a x86.M128) (dst x86.M64) {
	panic("not implemented")
}


// CvtSi2ss: Convert the 32-bit integer 'b' to a single-precision (32-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'CVTSI2SS'. Intrinsic: '_mm_cvt_si2ss'.
// Requires SSE.
func CvtSi2ss(a x86.M128, b int) (dst x86.M128) {
	panic("not implemented")
}


// CvtSs2si: Convert the lower single-precision (32-bit) floating-point element
// in 'a' to a 32-bit integer, and store the result in 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32(a[31:0])
//
// Instruction: 'CVTSS2SI'. Intrinsic: '_mm_cvt_ss2si'.
// Requires SSE.
func CvtSs2si(a x86.M128) int {
	panic("not implemented")
}


// Cvtpi16Ps: Convert packed 16-bit integers in 'a' to packed single-precision
// (32-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			m := j*32
//			dst[m+31:m] := Convert_Int16_To_FP32(a[i+15:i])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtpi16_ps'.
// Requires SSE.
func Cvtpi16Ps(a x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// Cvtpi32Ps: Convert packed 32-bit integers in 'b' to packed single-precision
// (32-bit) floating-point elements, store the results in the lower 2 elements
// of 'dst', and copy the upper 2 packed elements from 'a' to the upper
// elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[63:32] := Convert_Int32_To_FP32(b[63:32])
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'CVTPI2PS'. Intrinsic: '_mm_cvtpi32_ps'.
// Requires SSE.
func Cvtpi32Ps(a x86.M128, b x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// Cvtpi32x2Ps: Convert packed 32-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, store the results in the
// lower 2 elements of 'dst', then covert the packed 32-bit integers in 'a' to
// single-precision (32-bit) floating-point element, and store the results in
// the upper 2 elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(a[31:0])
//		dst[63:32] := Convert_Int32_To_FP32(a[63:32])
//		dst[95:64] := Convert_Int32_To_FP32(b[31:0])
//		dst[127:96] := Convert_Int32_To_FP32(b[63:32])
//
// Instruction: '...'. Intrinsic: '_mm_cvtpi32x2_ps'.
// Requires SSE.
func Cvtpi32x2Ps(a x86.M64, b x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// Cvtpi8Ps: Convert the lower packed 8-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*8
//			m := j*32
//			dst[m+31:m] := Convert_Int8_To_FP32(a[i+7:i])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtpi8_ps'.
// Requires SSE.
func Cvtpi8Ps(a x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// CvtpsPi16: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 16-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 16*j
//			k := 32*j
//			dst[i+15:i] := Convert_FP32_To_Int16(a[k+31:k])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtps_pi16'.
// Requires SSE.
func CvtpsPi16(a x86.M128) (dst x86.M64) {
	panic("not implemented")
}


// CvtpsPi32: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTPS2PI'. Intrinsic: '_mm_cvtps_pi32'.
// Requires SSE.
func CvtpsPi32(a x86.M128) (dst x86.M64) {
	panic("not implemented")
}


// CvtpsPi8: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 8-bit integers, and store the results in lower 4 elements
// of 'dst'. 
//
//		FOR j := 0 to 3
//			i := 8*j
//			k := 32*j
//			dst[i+7:i] := Convert_FP32_To_Int8(a[k+31:k])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtps_pi8'.
// Requires SSE.
func CvtpsPi8(a x86.M128) (dst x86.M64) {
	panic("not implemented")
}


// Cvtpu16Ps: Convert packed unsigned 16-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			m := j*32
//			dst[m+31:m] := Convert_UnsignedInt16_To_FP32(a[i+15:i])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtpu16_ps'.
// Requires SSE.
func Cvtpu16Ps(a x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// Cvtpu8Ps: Convert the lower packed unsigned 8-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*8
//			m := j*32
//			dst[m+31:m] := Convert_UnsignedInt8_To_FP32(a[i+7:i])
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_cvtpu8_ps'.
// Requires SSE.
func Cvtpu8Ps(a x86.M64) (dst x86.M128) {
	panic("not implemented")
}


// Cvtsi32Ss: Convert the 32-bit integer 'b' to a single-precision (32-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'CVTSI2SS'. Intrinsic: '_mm_cvtsi32_ss'.
// Requires SSE.
func Cvtsi32Ss(a x86.M128, b int) (dst x86.M128) {
	panic("not implemented")
}


// Cvtsi64Ss: Convert the 64-bit integer 'b' to a single-precision (32-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := Convert_Int64_To_FP32(b[63:0])
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'CVTSI2SS'. Intrinsic: '_mm_cvtsi64_ss'.
// Requires SSE.
func Cvtsi64Ss(a x86.M128, b int64) (dst x86.M128) {
	panic("not implemented")
}


// CvtssF32: Copy the lower single-precision (32-bit) floating-point element of
// 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVSS'. Intrinsic: '_mm_cvtss_f32'.
// Requires SSE.
func CvtssF32(a x86.M128) float32 {
	panic("not implemented")
}


// CvtssSi32: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 32-bit integer, and store the result in 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32(a[31:0])
//
// Instruction: 'CVTSS2SI'. Intrinsic: '_mm_cvtss_si32'.
// Requires SSE.
func CvtssSi32(a x86.M128) int {
	panic("not implemented")
}


// CvtssSi64: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 64-bit integer, and store the result in 'dst'. 
//
//		dst[63:0] := Convert_FP32_To_Int64(a[31:0])
//
// Instruction: 'CVTSS2SI'. Intrinsic: '_mm_cvtss_si64'.
// Requires SSE.
func CvtssSi64(a x86.M128) int64 {
	panic("not implemented")
}


// CvttPs2pi: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers with truncation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32_Truncate(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTTPS2PI'. Intrinsic: '_mm_cvtt_ps2pi'.
// Requires SSE.
func CvttPs2pi(a x86.M128) (dst x86.M64) {
	panic("not implemented")
}


// CvttSs2si: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 32-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32_Truncate(a[31:0])
//
// Instruction: 'CVTTSS2SI'. Intrinsic: '_mm_cvtt_ss2si'.
// Requires SSE.
func CvttSs2si(a x86.M128) int {
	panic("not implemented")
}


// CvttpsPi32: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers with truncation, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32_Truncate(a[i+31:i])
//		ENDFOR
//
// Instruction: 'CVTTPS2PI'. Intrinsic: '_mm_cvttps_pi32'.
// Requires SSE.
func CvttpsPi32(a x86.M128) (dst x86.M64) {
	panic("not implemented")
}


// CvttssSi32: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 32-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32_Truncate(a[31:0])
//
// Instruction: 'CVTTSS2SI'. Intrinsic: '_mm_cvttss_si32'.
// Requires SSE.
func CvttssSi32(a x86.M128) int {
	panic("not implemented")
}


// CvttssSi64: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 64-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[63:0] := Convert_FP64_To_Int32_Truncate(a[31:0])
//
// Instruction: 'CVTTSS2SI'. Intrinsic: '_mm_cvttss_si64'.
// Requires SSE.
func CvttssSi64(a x86.M128) int64 {
	panic("not implemented")
}


// DivEpi16: Divide packed 16-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			dst[i+15:i] := TRUNCATE(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epi16'.
// Requires SSE.
func DivEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epi32'.
// Requires SSE.
func DivEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpi64: Divide packed 64-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			dst[i+63:i] := TRUNCATE(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epi64'.
// Requires SSE.
func DivEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpi8: Divide packed 8-bit integers in 'a' by packed elements in 'b', and
// store the truncated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 8*j
//			dst[i+7:i] := TRUNCATE(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epi8'.
// Requires SSE.
func DivEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpu16: Divide packed unsigned 16-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			dst[i+15:i] := TRUNCATE(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epu16'.
// Requires SSE.
func DivEpu16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpu32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epu32'.
// Requires SSE.
func DivEpu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpu64: Divide packed unsigned 64-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			dst[i+63:i] := TRUNCATE(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epu64'.
// Requires SSE.
func DivEpu64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivEpu8: Divide packed unsigned 8-bit integers in 'a' by packed elements in
// 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 8*j
//			dst[i+7:i] := TRUNCATE(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_div_epu8'.
// Requires SSE.
func DivEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// DivPs: Divide packed single-precision (32-bit) floating-point elements in
// 'a' by packed elements in 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := a[i+31:i] / b[i+31:i]
//		ENDFOR
//
// Instruction: 'DIVPS'. Intrinsic: '_mm_div_ps'.
// Requires SSE.
func DivPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// DivSs: Divide the lower single-precision (32-bit) floating-point element in
// 'a' by the lower single-precision (32-bit) floating-point element in 'b',
// store the result in the lower element of 'dst', and copy the upper 3 packed
// elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0] / b[31:0]
//		dst[127:32] := a[127:32]
//
// Instruction: 'DIVSS'. Intrinsic: '_mm_div_ss'.
// Requires SSE.
func DivSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ErfPd: Compute the error function of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erf_pd'.
// Requires SSE.
func ErfPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// ErfPs: Compute the error function of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erf_ps'.
// Requires SSE.
func ErfPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ErfcPd: Compute the complementary error function of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := 1.0 - ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erfc_pd'.
// Requires SSE.
func ErfcPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// ErfcPs: Compute the complementary error function of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := 1.0 - ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erfc_ps'.
// Requires SSE.
func ErfcPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ErfcinvPd: Compute the inverse complementary error function of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := 1.0 / (1.0 - ERF(a[i+63:i]))
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erfcinv_pd'.
// Requires SSE.
func ErfcinvPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// ErfcinvPs: Compute the inverse complementary error function of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := 1.0 / (1.0 - ERF(a[i+31:i]))
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erfcinv_ps'.
// Requires SSE.
func ErfcinvPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ErfinvPd: Compute the inverse error function of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := 1.0 / ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erfinv_pd'.
// Requires SSE.
func ErfinvPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// ErfinvPs: Compute the inverse error function of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := 1.0 / ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_erfinv_ps'.
// Requires SSE.
func ErfinvPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ExpPd: Compute the exponential value of 'e' raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := e^(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_exp_pd'.
// Requires SSE.
func ExpPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// ExpPs: Compute the exponential value of 'e' raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_exp_ps'.
// Requires SSE.
func ExpPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Exp10Pd: Compute the exponential value of 10 raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := 10^(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_exp10_pd'.
// Requires SSE.
func Exp10Pd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Exp10Ps: Compute the exponential value of 10 raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := 10^(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_exp10_ps'.
// Requires SSE.
func Exp10Ps(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Exp2Pd: Compute the exponential value of 2 raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := 2^(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_exp2_pd'.
// Requires SSE.
func Exp2Pd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Exp2Ps: Compute the exponential value of 2 raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := 2^(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_exp2_ps'.
// Requires SSE.
func Exp2Ps(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Expm1Pd: Compute the exponential value of 'e' raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', subtract one from
// each element, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := e^(a[i+63:i]) - 1.0
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_expm1_pd'.
// Requires SSE.
func Expm1Pd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Expm1Ps: Compute the exponential value of 'e' raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', subtract one from
// each element, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i]) - 1.0
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_expm1_ps'.
// Requires SSE.
func Expm1Ps(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// ExtractPi16: Extract a 16-bit integer from 'a', selected with 'imm8', and
// store the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_mm_extract_pi16'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func ExtractPi16(a x86.M64, imm8 byte) int {
	panic("not implemented")
}


// Skipped: _mm_free. Contains pointer parameter.


// MMGETEXCEPTIONMASK: Macro: Get the exception mask bits from the MXCSR
// control and status register. The exception mask may contain any of the
// following flags: _MM_MASK_INVALID, _MM_MASK_DIV_ZERO, _MM_MASK_DENORM,
// _MM_MASK_OVERFLOW, _MM_MASK_UNDERFLOW, _MM_MASK_INEXACT 
//
//		dst[31:0] := MXCSR & _MM_MASK_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_EXCEPTION_MASK'.
// Requires SSE.
func MMGETEXCEPTIONMASK() uint32 {
	panic("not implemented")
}


// MMGETEXCEPTIONSTATE: Macro: Get the exception state bits from the MXCSR
// control and status register. The exception state may contain any of the
// following flags: _MM_EXCEPT_INVALID, _MM_EXCEPT_DIV_ZERO, _MM_EXCEPT_DENORM,
// _MM_EXCEPT_OVERFLOW, _MM_EXCEPT_UNDERFLOW, _MM_EXCEPT_INEXACT 
//
//		dst[31:0] := MXCSR & _MM_EXCEPT_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_EXCEPTION_STATE'.
// Requires SSE.
func MMGETEXCEPTIONSTATE() uint32 {
	panic("not implemented")
}


// MMGETFLUSHZEROMODE: Macro: Get the flush zero bits from the MXCSR control
// and status register. The flush zero may contain any of the following flags:
// _MM_FLUSH_ZERO_ON or _MM_FLUSH_ZERO_OFF 
//
//		dst[31:0] := MXCSR & _MM_FLUSH_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_FLUSH_ZERO_MODE'.
// Requires SSE.
func MMGETFLUSHZEROMODE() uint32 {
	panic("not implemented")
}


// MMGETROUNDINGMODE: Macro: Get the rounding mode bits from the MXCSR control
// and status register. The rounding mode may contain any of the following
// flags: _MM_ROUND_NEAREST, _MM_ROUND_DOWN, _MM_ROUND_UP,
// _MM_ROUND_TOWARD_ZERO 
//
//		dst[31:0] := MXCSR & _MM_ROUND_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_ROUNDING_MODE'.
// Requires SSE.
func MMGETROUNDINGMODE() uint32 {
	panic("not implemented")
}


// Getcsr: Get the unsigned 32-bit value of the MXCSR control and status
// register. 
//
//		dst[31:0] := MXCSR
//
// Instruction: 'STMXCSR'. Intrinsic: '_mm_getcsr'.
// Requires SSE.
func Getcsr() uint32 {
	panic("not implemented")
}


// HypotPd: Compute the length of the hypotenous of a right triangle, with the
// lengths of the other two sides of the triangle stored as packed
// double-precision (64-bit) floating-point elements in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SQRT(a[i+63:i]^2 + b[i+63:i]^2)
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_hypot_pd'.
// Requires SSE.
func HypotPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// HypotPs: Compute the length of the hypotenous of a right triangle, with the
// lengths of the other two sides of the triangle stored as packed
// single-precision (32-bit) floating-point elements in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i]^2 + b[i+31:i]^2)
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_hypot_ps'.
// Requires SSE.
func HypotPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// IdivEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_idiv_epi32'.
// Requires SSE.
func IdivEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// IdivremEpi32: Divide packed 32-bit integers in 'a' by packed elements in
// 'b', store the truncated results in 'dst', and store the remainders as
// packed 32-bit integers into memory at 'mem_addr'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//			MEM[mem_addr+i+31:mem_addr+i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_idivrem_epi32'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func IdivremEpi32(mem_addr *x86.M128i, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// InsertPi16: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst'
// at the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_mm_insert_pi16'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func InsertPi16(a x86.M64, i int, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


// InvcbrtPd: Compute the inverse cube root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := InvCubeRoot(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_invcbrt_pd'.
// Requires SSE.
func InvcbrtPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// InvcbrtPs: Compute the inverse cube root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := InvCubeRoot(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_invcbrt_ps'.
// Requires SSE.
func InvcbrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// InvsqrtPd: Compute the inverse square root of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := InvSQRT(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_invsqrt_pd'.
// Requires SSE.
func InvsqrtPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// InvsqrtPs: Compute the inverse square root of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := InvSQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_invsqrt_ps'.
// Requires SSE.
func InvsqrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// IremEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_irem_epi32'.
// Requires SSE.
func IremEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// Skipped: _mm_load_ps. Contains pointer parameter.


// Skipped: _mm_load_ps1. Contains pointer parameter.


// Skipped: _mm_load_ss. Contains pointer parameter.


// Skipped: _mm_load1_ps. Contains pointer parameter.


// LoadhPi: Load 2 single-precision (32-bit) floating-point elements from
// memory into the upper 2 elements of 'dst', and copy the lower 2 elements
// from 'a' to 'dst'. 'mem_addr' does not need to be aligned on any particular
// boundary. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := a[63:32]
//		dst[95:64] := MEM[mem_addr+31:mem_addr]
//		dst[127:96] := MEM[mem_addr+63:mem_addr+32]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_loadh_pi'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func LoadhPi(a x86.M128, mem_addr *x86.M64Const) (dst x86.M128) {
	panic("not implemented")
}


// LoadlPi: Load 2 single-precision (32-bit) floating-point elements from
// memory into the lower 2 elements of 'dst', and copy the upper 2 elements
// from 'a' to 'dst'. 'mem_addr' does not need to be aligned on any particular
// boundary. 
//
//		dst[31:0] := MEM[mem_addr+31:mem_addr]
//		dst[63:32] := MEM[mem_addr+63:mem_addr+32]
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_loadl_pi'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func LoadlPi(a x86.M128, mem_addr *x86.M64Const) (dst x86.M128) {
	panic("not implemented")
}


// Skipped: _mm_loadr_ps. Contains pointer parameter.


// Skipped: _mm_loadu_ps. Contains pointer parameter.


// Skipped: _mm_loadu_si16. Contains pointer parameter.


// Skipped: _mm_loadu_si32. Contains pointer parameter.


// Skipped: _mm_loadu_si64. Contains pointer parameter.


// LogPd: Compute the natural logarithm of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ln(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log_pd'.
// Requires SSE.
func LogPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// LogPs: Compute the natural logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ln(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log_ps'.
// Requires SSE.
func LogPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Log10Pd: Compute the base-10 logarithm of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := log10(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log10_pd'.
// Requires SSE.
func Log10Pd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Log10Ps: Compute the base-10 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := log10(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log10_ps'.
// Requires SSE.
func Log10Ps(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Log1pPd: Compute the natural logarithm of one plus packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ln(1.0 + a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log1p_pd'.
// Requires SSE.
func Log1pPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Log1pPs: Compute the natural logarithm of one plus packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ln(1.0 + a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log1p_ps'.
// Requires SSE.
func Log1pPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Log2Pd: Compute the base-2 logarithm of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := log2(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log2_pd'.
// Requires SSE.
func Log2Pd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// Log2Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := log2(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_log2_ps'.
// Requires SSE.
func Log2Ps(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// LogbPd: Convert the exponent of each packed double-precision (64-bit)
// floating-point element in 'a' to a double-precision floating-point number
// representing the integer exponent, and store the results in 'dst'. This
// intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_logb_pd'.
// Requires SSE.
func LogbPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// LogbPs: Convert the exponent of each packed single-precision (32-bit)
// floating-point element in 'a' to a single-precision floating-point number
// representing the integer exponent, and store the results in 'dst'. This
// intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_logb_ps'.
// Requires SSE.
func LogbPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Skipped: _mm_malloc. Contains pointer parameter.


// MaskmoveSi64: Conditionally store 8-bit integer elements from 'a' into
// memory using 'mask' (elements are not stored when the highest bit is not set
// in the corresponding element) and a non-temporal memory hint. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF mask[i+7]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'MASKMOVQ'. Intrinsic: '_mm_maskmove_si64'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskmoveSi64(a x86.M64, mask x86.M64, mem_addr *byte)  {
	panic("not implemented")
}


// Maskmovq: Conditionally store 8-bit integer elements from 'a' into memory
// using 'mask' (elements are not stored when the highest bit is not set in the
// corresponding element). 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF mask[i+7]
//				MEM[mem_addr+i+7:mem_addr+i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'MASKMOVQ'. Intrinsic: '_m_maskmovq'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Maskmovq(a x86.M64, mask x86.M64, mem_addr *byte)  {
	panic("not implemented")
}


// MaxPi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSW'. Intrinsic: '_mm_max_pi16'.
// Requires SSE.
func MaxPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MaxPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := MAX(a[i+31:i], b[i+31:i])
//		ENDFOR
//
// Instruction: 'MAXPS'. Intrinsic: '_mm_max_ps'.
// Requires SSE.
func MaxPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MaxPu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUB'. Intrinsic: '_mm_max_pu8'.
// Requires SSE.
func MaxPu8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MaxSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b', store the maximum value in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[31:0] := MAX(a[31:0], b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'MAXSS'. Intrinsic: '_mm_max_ss'.
// Requires SSE.
func MaxSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MinPi16: Compare packed 16-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSW'. Intrinsic: '_mm_min_pi16'.
// Requires SSE.
func MinPi16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MinPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := MIN(a[i+31:i], b[i+31:i])
//		ENDFOR
//
// Instruction: 'MINPS'. Intrinsic: '_mm_min_ps'.
// Requires SSE.
func MinPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MinPu8: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUB'. Intrinsic: '_mm_min_pu8'.
// Requires SSE.
func MinPu8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MinSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b', store the minimum value in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[31:0] := MIN(a[31:0], b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'MINSS'. Intrinsic: '_mm_min_ss'.
// Requires SSE.
func MinSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MoveSs: Move the lower single-precision (32-bit) floating-point element from
// 'b' to the lower element of 'dst', and copy the upper 3 elements from 'a' to
// the upper elements of 'dst'. 
//
//		dst[31:0] := b[31:0]
//		dst[63:32] := a[63:32]
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'MOVSS'. Intrinsic: '_mm_move_ss'.
// Requires SSE.
func MoveSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MovehlPs: Move the upper 2 single-precision (32-bit) floating-point elements
// from 'b' to the lower 2 elements of 'dst', and copy the upper 2 elements
// from 'a' to the upper 2 elements of 'dst'. 
//
//		dst[31:0] := b[95:64]
//		dst[63:32] := b[127:96]
//		dst[95:64] := a[95:64]
//		dst[127:96] := a[127:96]
//
// Instruction: 'MOVHLPS'. Intrinsic: '_mm_movehl_ps'.
// Requires SSE.
func MovehlPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MovelhPs: Move the lower 2 single-precision (32-bit) floating-point elements
// from 'b' to the upper 2 elements of 'dst', and copy the lower 2 elements
// from 'a' to the lower 2 elements of 'dst'. 
//
//		dst[31:0] := a[31:0]
//		dst[63:32] := a[63:32]
//		dst[95:64] := b[31:0]
//		dst[127:96] := b[63:32]
//
// Instruction: 'MOVLHPS'. Intrinsic: '_mm_movelh_ps'.
// Requires SSE.
func MovelhPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MovemaskPi8: Create mask from the most significant bit of each 8-bit element
// in 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//		dst[MAX:8] := 0
//
// Instruction: 'PMOVMSKB'. Intrinsic: '_mm_movemask_pi8'.
// Requires SSE.
func MovemaskPi8(a x86.M64) int {
	panic("not implemented")
}


// MovemaskPs: Set each bit of mask 'dst' based on the most significant bit of
// the corresponding packed single-precision (32-bit) floating-point element in
// 'a'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF a[i+31]
//				dst[j] := 1
//			ELSE
//				dst[j] := 0
//			FI
//		ENDFOR
//		dst[MAX:4] := 0
//
// Instruction: 'MOVMSKPS'. Intrinsic: '_mm_movemask_ps'.
// Requires SSE.
func MovemaskPs(a x86.M128) int {
	panic("not implemented")
}


// MulPs: Multiply packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//
// Instruction: 'MULPS'. Intrinsic: '_mm_mul_ps'.
// Requires SSE.
func MulPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MulSs: Multiply the lower single-precision (32-bit) floating-point element
// in 'a' and 'b', store the result in the lower element of 'dst', and copy the
// upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0] * b[31:0]
//		dst[127:32] := a[127:32]
//
// Instruction: 'MULSS'. Intrinsic: '_mm_mul_ss'.
// Requires SSE.
func MulSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MulhiPu16: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHUW'. Intrinsic: '_mm_mulhi_pu16'.
// Requires SSE.
func MulhiPu16(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// OrPs: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//		ENDFOR
//
// Instruction: 'ORPS'. Intrinsic: '_mm_or_ps'.
// Requires SSE.
func OrPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Pavgb: Average packed unsigned 8-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[i+7:i] := (a[i+7:i] + b[i+7:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGB'. Intrinsic: '_m_pavgb'.
// Requires SSE.
func Pavgb(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pavgw: Average packed unsigned 16-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			dst[i+15:i] := (a[i+15:i] + b[i+15:i] + 1) >> 1
//		ENDFOR
//
// Instruction: 'PAVGW'. Intrinsic: '_m_pavgw'.
// Requires SSE.
func Pavgw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pextrw: Extract a 16-bit integer from 'a', selected with 'imm8', and store
// the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_m_pextrw'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func Pextrw(a x86.M64, imm8 byte) int {
	panic("not implemented")
}


// Pinsrw: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_m_pinsrw'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func Pinsrw(a x86.M64, i int, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


// Pmaxsw: Compare packed 16-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSW'. Intrinsic: '_m_pmaxsw'.
// Requires SSE.
func Pmaxsw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pmaxub: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUB'. Intrinsic: '_m_pmaxub'.
// Requires SSE.
func Pmaxub(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pminsw: Compare packed 16-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSW'. Intrinsic: '_m_pminsw'.
// Requires SSE.
func Pminsw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pminub: Compare packed unsigned 8-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUB'. Intrinsic: '_m_pminub'.
// Requires SSE.
func Pminub(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pmovmskb: Create mask from the most significant bit of each 8-bit element in
// 'a', and store the result in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			dst[j] := a[i+7]
//		ENDFOR
//		dst[MAX:8] := 0
//
// Instruction: 'PMOVMSKB'. Intrinsic: '_m_pmovmskb'.
// Requires SSE.
func Pmovmskb(a x86.M64) int {
	panic("not implemented")
}


// Pmulhuw: Multiply the packed unsigned 16-bit integers in 'a' and 'b',
// producing intermediate 32-bit integers, and store the high 16 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*16
//			tmp[31:0] := a[i+15:i] * b[i+15:i]
//			dst[i+15:i] := tmp[31:16]
//		ENDFOR
//
// Instruction: 'PMULHUW'. Intrinsic: '_m_pmulhuw'.
// Requires SSE.
func Pmulhuw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// PowPd: Compute the exponential value of packed double-precision (64-bit)
// floating-point elements in 'a' raised by packed elements in 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := (a[i+63:i])^(b[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_pow_pd'.
// Requires SSE.
func PowPd(a x86.M128d, b x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// PowPs: Compute the exponential value of packed single-precision (32-bit)
// floating-point elements in 'a' raised by packed elements in 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := (a[i+31:i])^(b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_pow_ps'.
// Requires SSE.
func PowPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'PREFETCHNTA, PREFETCHT0, PREFETCHT1, PREFETCHT2'. Intrinsic: '_mm_prefetch'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Prefetch(p *byte, i int)  {
	panic("not implemented")
}


// Psadbw: Compute the absolute differences of packed unsigned 8-bit integers
// in 'a' and 'b', then horizontally sum each consecutive 8 differences to
// produce four unsigned 16-bit integers, and pack these unsigned 16-bit
// integers in the low 16 bits of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		
//		dst[15:0] := tmp[7:0] + tmp[15:8] + tmp[23:16] + tmp[31:24] + tmp[39:32] + tmp[47:40] + tmp[55:48] + tmp[63:56]
//		dst[63:16] := 0
//
// Instruction: 'PSADBW'. Intrinsic: '_m_psadbw'.
// Requires SSE.
func Psadbw(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// Pshufw: Shuffle 16-bit integers in 'a' using the control in 'imm8', and
// store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[15:0] := src[15:0]
//			1:	tmp[15:0] := src[31:16]
//			2:	tmp[15:0] := src[47:32]
//			3:	tmp[15:0] := src[63:48]
//			ESAC
//			RETURN tmp[15:0]
//		}
//		
//		dst[15:0] := SELECT4(a[63:0], imm8[1:0])
//		dst[31:16] := SELECT4(a[63:0], imm8[3:2])
//		dst[47:32] := SELECT4(a[63:0], imm8[5:4])
//		dst[63:48] := SELECT4(a[63:0], imm8[7:6])
//
// Instruction: 'PSHUFW'. Intrinsic: '_m_pshufw'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func Pshufw(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


// RcpPs: Compute the approximate reciprocal of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. The
// maximum relative error for this approximation is less than 1.5*2^-12. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := APPROXIMATE(1.0/a[i+31:i])
//		ENDFOR
//
// Instruction: 'RCPPS'. Intrinsic: '_mm_rcp_ps'.
// Requires SSE.
func RcpPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// RcpSs: Compute the approximate reciprocal of the lower single-precision
// (32-bit) floating-point element in 'a', store the result in the lower
// element of 'dst', and copy the upper 3 packed elements from 'a' to the upper
// elements of 'dst'. The maximum relative error for this approximation is less
// than 1.5*2^-12. 
//
//		dst[31:0] := APPROXIMATE(1.0/a[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'RCPSS'. Intrinsic: '_mm_rcp_ss'.
// Requires SSE.
func RcpSs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// RemEpi16: Divide packed 16-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			dst[i+15:i] := REMAINDER(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epi16'.
// Requires SSE.
func RemEpi16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epi32'.
// Requires SSE.
func RemEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpi64: Divide packed 64-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			dst[i+63:i] := REMAINDER(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epi64'.
// Requires SSE.
func RemEpi64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpi8: Divide packed 8-bit integers in 'a' by packed elements in 'b', and
// store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 8*j
//			dst[i+7:i] := REMAINDER(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epi8'.
// Requires SSE.
func RemEpi8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpu16: Divide packed unsigned 16-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := 16*j
//			dst[i+15:i] := REMAINDER(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epu16'.
// Requires SSE.
func RemEpu16(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpu32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epu32'.
// Requires SSE.
func RemEpu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpu64: Divide packed unsigned 64-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			dst[i+63:i] := REMAINDER(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epu64'.
// Requires SSE.
func RemEpu64(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RemEpu8: Divide packed unsigned 8-bit integers in 'a' by packed elements in
// 'b', and store the remainders as packed unsigned 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 8*j
//			dst[i+7:i] := REMAINDER(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_rem_epu8'.
// Requires SSE.
func RemEpu8(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// RsqrtPs: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. The maximum relative error for this approximation is less
// than 1.5*2^-12. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := APPROXIMATE(1.0 / SQRT(a[i+31:i]))
//		ENDFOR
//
// Instruction: 'RSQRTPS'. Intrinsic: '_mm_rsqrt_ps'.
// Requires SSE.
func RsqrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// RsqrtSs: Compute the approximate reciprocal square root of the lower
// single-precision (32-bit) floating-point element in 'a', store the result in
// the lower element of 'dst', and copy the upper 3 packed elements from 'a' to
// the upper elements of 'dst'. The maximum relative error for this
// approximation is less than 1.5*2^-12. 
//
//		dst[31:0] := APPROXIMATE(1.0 / SQRT(a[31:0]))
//		dst[127:32] := a[127:32]
//
// Instruction: 'RSQRTSS'. Intrinsic: '_mm_rsqrt_ss'.
// Requires SSE.
func RsqrtSs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SadPu8: Compute the absolute differences of packed unsigned 8-bit integers
// in 'a' and 'b', then horizontally sum each consecutive 8 differences to
// produce four unsigned 16-bit integers, and pack these unsigned 16-bit
// integers in the low 16 bits of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			tmp[i+7:i] := ABS(a[i+7:i] - b[i+7:i])
//		ENDFOR
//		
//		dst[15:0] := tmp[7:0] + tmp[15:8] + tmp[23:16] + tmp[31:24] + tmp[39:32] + tmp[47:40] + tmp[55:48] + tmp[63:56]
//		dst[63:16] := 0
//
// Instruction: 'PSADBW'. Intrinsic: '_mm_sad_pu8'.
// Requires SSE.
func SadPu8(a x86.M64, b x86.M64) (dst x86.M64) {
	panic("not implemented")
}


// MMSETEXCEPTIONMASK: Macro: Set the exception mask bits of the MXCSR control
// and status register to the value in unsigned 32-bit integer 'a'. The
// exception mask may contain any of the following flags: _MM_MASK_INVALID,
// _MM_MASK_DIV_ZERO, _MM_MASK_DENORM, _MM_MASK_OVERFLOW, _MM_MASK_UNDERFLOW,
// _MM_MASK_INEXACT 
//
//		MXCSR := a[31:0] AND ~_MM_MASK_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_EXCEPTION_MASK'.
// Requires SSE.
func MMSETEXCEPTIONMASK(a uint32)  {
	panic("not implemented")
}


// MMSETEXCEPTIONSTATE: Macro: Set the exception state bits of the MXCSR
// control and status register to the value in unsigned 32-bit integer 'a'. The
// exception state may contain any of the following flags: _MM_EXCEPT_INVALID,
// _MM_EXCEPT_DIV_ZERO, _MM_EXCEPT_DENORM, _MM_EXCEPT_OVERFLOW,
// _MM_EXCEPT_UNDERFLOW, _MM_EXCEPT_INEXACT 
//
//		MXCSR := a[31:0] AND ~_MM_EXCEPT_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_EXCEPTION_STATE'.
// Requires SSE.
func MMSETEXCEPTIONSTATE(a uint32)  {
	panic("not implemented")
}


// MMSETFLUSHZEROMODE: Macro: Set the flush zero bits of the MXCSR control and
// status register to the value in unsigned 32-bit integer 'a'. The flush zero
// may contain any of the following flags: _MM_FLUSH_ZERO_ON or
// _MM_FLUSH_ZERO_OFF 
//
//		MXCSR := a[31:0] AND ~_MM_FLUSH_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_FLUSH_ZERO_MODE'.
// Requires SSE.
func MMSETFLUSHZEROMODE(a uint32)  {
	panic("not implemented")
}


// SetPs: Set packed single-precision (32-bit) floating-point elements in 'dst'
// with the supplied values. 
//
//		dst[31:0] := e0
//		dst[63:32] := e1
//		dst[95:64] := e2
//		dst[127:96] := e3
//
// Instruction: '...'. Intrinsic: '_mm_set_ps'.
// Requires SSE.
func SetPs(e3 float32, e2 float32, e1 float32, e0 float32) (dst x86.M128) {
	panic("not implemented")
}


// SetPs1: Broadcast single-precision (32-bit) floating-point value 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set_ps1'.
// Requires SSE.
func SetPs1(a float32) (dst x86.M128) {
	panic("not implemented")
}


// MMSETROUNDINGMODE: Macro: Set the rounding mode bits of the MXCSR control
// and status register to the value in unsigned 32-bit integer 'a'. The
// rounding mode may contain any of the following flags: _MM_ROUND_NEAREST,
// _MM_ROUND_DOWN, _MM_ROUND_UP, _MM_ROUND_TOWARD_ZERO 
//
//		MXCSR := a[31:0] AND ~_MM_ROUND_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_ROUNDING_MODE'.
// Requires SSE.
func MMSETROUNDINGMODE(a uint32)  {
	panic("not implemented")
}


// SetSs: Copy single-precision (32-bit) floating-point element 'a' to the
// lower element of 'dst', and zero the upper 3 elements. 
//
//		dst[31:0] := a[31:0]
//		dst[127:32] := 0
//
// Instruction: '...'. Intrinsic: '_mm_set_ss'.
// Requires SSE.
func SetSs(a float32) (dst x86.M128) {
	panic("not implemented")
}


// Set1Ps: Broadcast single-precision (32-bit) floating-point value 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_mm_set1_ps'.
// Requires SSE.
func Set1Ps(a float32) (dst x86.M128) {
	panic("not implemented")
}


// Setcsr: Set the MXCSR control and status register with the value in unsigned
// 32-bit integer 'a'. 
//
//		MXCSR := a[31:0]
//
// Instruction: 'LDMXCSR'. Intrinsic: '_mm_setcsr'.
// Requires SSE.
func Setcsr(a uint32)  {
	panic("not implemented")
}


// SetrPs: Set packed single-precision (32-bit) floating-point elements in
// 'dst' with the supplied values in reverse order. 
//
//		dst[31:0] := e3
//		dst[63:32] := e2
//		dst[95:64] := e1
//		dst[127:96] := e0
//
// Instruction: '...'. Intrinsic: '_mm_setr_ps'.
// Requires SSE.
func SetrPs(e3 float32, e2 float32, e1 float32, e0 float32) (dst x86.M128) {
	panic("not implemented")
}


// SetzeroPs: Return vector of type __m128 with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'XORPS'. Intrinsic: '_mm_setzero_ps'.
// Requires SSE.
func SetzeroPs() (dst x86.M128) {
	panic("not implemented")
}


// Sfence: Perform a serializing operation on all store-to-memory instructions
// that were issued prior to this instruction. Guarantees that every store
// instruction that precedes, in program order, is globally visible before any
// store instruction which follows the fence in program order. 
//
//		
//
// Instruction: 'SFENCE'. Intrinsic: '_mm_sfence'.
// Requires SSE.
func Sfence()  {
	panic("not implemented")
}


// ShufflePi16: Shuffle 16-bit integers in 'a' using the control in 'imm8', and
// store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[15:0] := src[15:0]
//			1:	tmp[15:0] := src[31:16]
//			2:	tmp[15:0] := src[47:32]
//			3:	tmp[15:0] := src[63:48]
//			ESAC
//			RETURN tmp[15:0]
//		}
//		
//		dst[15:0] := SELECT4(a[63:0], imm8[1:0])
//		dst[31:16] := SELECT4(a[63:0], imm8[3:2])
//		dst[47:32] := SELECT4(a[63:0], imm8[5:4])
//		dst[63:48] := SELECT4(a[63:0], imm8[7:6])
//
// Instruction: 'PSHUFW'. Intrinsic: '_mm_shuffle_pi16'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func ShufflePi16(a x86.M64, imm8 byte) (dst x86.M64) {
	panic("not implemented")
}


// ShufflePs: Shuffle single-precision (32-bit) floating-point elements in 'a'
// using the control in 'imm8', and store the results in 'dst'. 
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
//		dst[95:64] := SELECT4(b[127:0], imm8[5:4])
//		dst[127:96] := SELECT4(b[127:0], imm8[7:6])
//
// Instruction: 'SHUFPS'. Intrinsic: '_mm_shuffle_ps'.
// Requires SSE.
//
// FIXME: Requires compiler support (has immediate)
func ShufflePs(a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	panic("not implemented")
}


// SinPd: Compute the sine of packed double-precision (64-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SIN(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sin_pd'.
// Requires SSE.
func SinPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SinPs: Compute the sine of packed single-precision (32-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SIN(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sin_ps'.
// Requires SSE.
func SinPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SincosPd: Compute the sine and cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, store the sine in
// 'dst', and store the cosine into memory at 'mem_addr'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SIN(a[i+63:i])
//			MEM[mem_addr+i+63:mem_addr+i] := COS(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sincos_pd'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func SincosPd(mem_addr *x86.M128d, a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SincosPs: Compute the sine and cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, store the sine in
// 'dst', and store the cosine into memory at 'mem_addr'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SIN(a[i+31:i])
//			MEM[mem_addr+i+31:mem_addr+i] := COS(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sincos_ps'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func SincosPs(mem_addr *x86.M128, a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SindPd: Compute the sine of packed double-precision (64-bit) floating-point
// elements in 'a' expressed in degrees, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SIND(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sind_pd'.
// Requires SSE.
func SindPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SindPs: Compute the sine of packed single-precision (32-bit) floating-point
// elements in 'a' expressed in degrees, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SIND(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sind_ps'.
// Requires SSE.
func SindPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SinhPd: Compute the hyperbolic sine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SINH(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sinh_pd'.
// Requires SSE.
func SinhPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SinhPs: Compute the hyperbolic sine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SINH(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_sinh_ps'.
// Requires SSE.
func SinhPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SqrtPs: Compute the square root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//
// Instruction: 'SQRTPS'. Intrinsic: '_mm_sqrt_ps'.
// Requires SSE.
func SqrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SqrtSs: Compute the square root of the lower single-precision (32-bit)
// floating-point element in 'a', store the result in the lower element of
// 'dst', and copy the upper 3 packed elements from 'a' to the upper elements
// of 'dst'. 
//
//		dst[31:0] := SQRT(a[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'SQRTSS'. Intrinsic: '_mm_sqrt_ss'.
// Requires SSE.
func SqrtSs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// StorePs: Store 128-bits (composed of 4 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVAPS'. Intrinsic: '_mm_store_ps'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorePs(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// StorePs1: Store the lower single-precision (32-bit) floating-point element
// from 'a' into 4 contiguous elements in memory. 'mem_addr' must be aligned on
// a 16-byte boundary or a general-protection exception may be generated. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//		MEM[mem_addr+63:mem_addr+32] := a[31:0]
//		MEM[mem_addr+95:mem_addr+64] := a[31:0]
//		MEM[mem_addr+127:mem_addr+96] := a[31:0]
//
// Instruction: '...'. Intrinsic: '_mm_store_ps1'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorePs1(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// StoreSs: Store the lower single-precision (32-bit) floating-point element
// from 'a' into memory. 'mem_addr' does not need to be aligned on any
// particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVSS'. Intrinsic: '_mm_store_ss'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StoreSs(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// Store1Ps: Store the lower single-precision (32-bit) floating-point element
// from 'a' into 4 contiguous elements in memory. 'mem_addr' must be aligned on
// a 16-byte boundary or a general-protection exception may be generated. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//		MEM[mem_addr+63:mem_addr+32] := a[31:0]
//		MEM[mem_addr+95:mem_addr+64] := a[31:0]
//		MEM[mem_addr+127:mem_addr+96] := a[31:0]
//
// Instruction: '...'. Intrinsic: '_mm_store1_ps'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Store1Ps(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// StorehPi: Store the upper 2 single-precision (32-bit) floating-point
// elements from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[95:64]
//		MEM[mem_addr+63:mem_addr+32] := a[127:96]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_storeh_pi'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorehPi(mem_addr *x86.M64, a x86.M128)  {
	panic("not implemented")
}


// StorelPi: Store the lower 2 single-precision (32-bit) floating-point
// elements from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//		MEM[mem_addr+63:mem_addr+32] := a[63:32]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_storel_pi'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorelPi(mem_addr *x86.M64, a x86.M128)  {
	panic("not implemented")
}


// StorerPs: Store 4 single-precision (32-bit) floating-point elements from 'a'
// into memory in reverse order.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+31:mem_addr] := a[127:96]
//		MEM[mem_addr+63:mem_addr+32] := a[95:64]
//		MEM[mem_addr+95:mem_addr+64] := a[63:32]
//		MEM[mem_addr+127:mem_addr+96] := a[31:0]
//
// Instruction: '...'. Intrinsic: '_mm_storer_ps'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StorerPs(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// StoreuPs: Store 128-bits (composed of 4 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVUPS'. Intrinsic: '_mm_storeu_ps'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StoreuPs(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// Skipped: _mm_storeu_si16. Contains pointer parameter.


// Skipped: _mm_storeu_si32. Contains pointer parameter.


// Skipped: _mm_storeu_si64. Contains pointer parameter.


// StreamPi: Store 64-bits of integer data from 'a' into memory using a
// non-temporal memory hint. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVNTQ'. Intrinsic: '_mm_stream_pi'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StreamPi(mem_addr *x86.M64, a x86.M64)  {
	panic("not implemented")
}


// StreamPs: Store 128-bits (composed of 4 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory using a non-temporal memory
// hint.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVNTPS'. Intrinsic: '_mm_stream_ps'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func StreamPs(mem_addr *float32, a x86.M128)  {
	panic("not implemented")
}


// SubPs: Subtract packed single-precision (32-bit) floating-point elements in
// 'b' from packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//
// Instruction: 'SUBPS'. Intrinsic: '_mm_sub_ps'.
// Requires SSE.
func SubPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SubSs: Subtract the lower single-precision (32-bit) floating-point element
// in 'b' from the lower single-precision (32-bit) floating-point element in
// 'a', store the result in the lower element of 'dst', and copy the upper 3
// packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0] - b[31:0]
//		dst[127:32] := a[127:32]
//
// Instruction: 'SUBSS'. Intrinsic: '_mm_sub_ss'.
// Requires SSE.
func SubSs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SvmlCeilPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' up to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := CEIL(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_ceil_pd'.
// Requires SSE.
func SvmlCeilPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SvmlCeilPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' up to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := CEIL(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_ceil_ps'.
// Requires SSE.
func SvmlCeilPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SvmlFloorPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' down to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := FLOOR(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_floor_pd'.
// Requires SSE.
func SvmlFloorPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SvmlFloorPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' down to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := FLOOR(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_floor_ps'.
// Requires SSE.
func SvmlFloorPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SvmlRoundPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' to the nearest integer value, and store the results as
// packed double-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ROUND(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_round_pd'.
// Requires SSE.
func SvmlRoundPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SvmlRoundPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' to the nearest integer value, and store the results as
// packed single-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_round_ps'.
// Requires SSE.
func SvmlRoundPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// SvmlSqrtPd: Compute the square root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. Note that
// this intrinsic is less efficient than '_mm_sqrt_pd'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := SQRT(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_sqrt_pd'.
// Requires SSE.
func SvmlSqrtPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// SvmlSqrtPs: Compute the square root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. Note that
// this intrinsic is less efficient than '_mm_sqrt_ps'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_svml_sqrt_ps'.
// Requires SSE.
func SvmlSqrtPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// TanPd: Compute the tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := TAN(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_tan_pd'.
// Requires SSE.
func TanPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// TanPs: Compute the tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := TAN(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_tan_ps'.
// Requires SSE.
func TanPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// TandPd: Compute the tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := TAND(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_tand_pd'.
// Requires SSE.
func TandPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// TandPs: Compute the tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := TAND(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_tand_ps'.
// Requires SSE.
func TandPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// TanhPd: Compute the hyperbolic tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := TANH(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_tanh_pd'.
// Requires SSE.
func TanhPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// TanhPs: Compute the hyperbolic tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := TANH(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_tanh_ps'.
// Requires SSE.
func TanhPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// MMTRANSPOSE4PS: Macro: Transpose the 4x4 matrix formed by the 4 rows of
// single-precision (32-bit) floating-point elements in 'row0', 'row1', 'row2',
// and 'row3', and store the transposed matrix in these vectors ('row0' now
// contains column 0, etc.). 
//
//		__m128 tmp3, tmp2, tmp1, tmp0;
//		tmp0 = _mm_unpacklo_ps(row0, row1);
//		tmp2 = _mm_unpacklo_ps(row2, row3);
//		tmp1 = _mm_unpackhi_ps(row0, row1);
//		tmp3 = _mm_unpackhi_ps(row2, row3);
//		row0 = _mm_movelh_ps(tmp0, tmp2);
//		row1 = _mm_movehl_ps(tmp2, tmp0);
//		row2 = _mm_movelh_ps(tmp1, tmp3);
//		row3 = _mm_movehl_ps(tmp3, tmp1);
//
// Instruction: '...'. Intrinsic: '_MM_TRANSPOSE4_PS'.
// Requires SSE.
func MMTRANSPOSE4PS(row0 x86.M128, row1 x86.M128, row2 x86.M128, row3 x86.M128)  {
	panic("not implemented")
}


// TruncPd: Truncate the packed double-precision (64-bit) floating-point
// elements in 'a', and store the results as packed double-precision
// floating-point elements in 'dst'. This intrinsic may generate the
// 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := TRUNCATE(a[i+63:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_trunc_pd'.
// Requires SSE.
func TruncPd(a x86.M128d) (dst x86.M128d) {
	panic("not implemented")
}


// TruncPs: Truncate the packed single-precision (32-bit) floating-point
// elements in 'a', and store the results as packed single-precision
// floating-point elements in 'dst'. This intrinsic may generate the
// 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := TRUNCATE(a[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_trunc_ps'.
// Requires SSE.
func TruncPs(a x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// UcomieqSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for equality, and return the boolean result (0 or 1).
// This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] == b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomieq_ss'.
// Requires SSE.
func UcomieqSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// UcomigeSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for greater-than-or-equal, and return the boolean
// result (0 or 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] >= b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomige_ss'.
// Requires SSE.
func UcomigeSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// UcomigtSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for greater-than, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] > b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomigt_ss'.
// Requires SSE.
func UcomigtSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// UcomileSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for less-than-or-equal, and return the boolean result
// (0 or 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] <= b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomile_ss'.
// Requires SSE.
func UcomileSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// UcomiltSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for less-than, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] < b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomilt_ss'.
// Requires SSE.
func UcomiltSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// UcomineqSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for not-equal, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] != b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomineq_ss'.
// Requires SSE.
func UcomineqSs(a x86.M128, b x86.M128) int {
	panic("not implemented")
}


// UdivEpi32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_udiv_epi32'.
// Requires SSE.
func UdivEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// UdivremEpi32: Divide packed unsigned 32-bit integers in 'a' by packed
// elements in 'b', store the truncated results in 'dst', and store the
// remainders as packed unsigned 32-bit integers into memory at 'mem_addr'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//			MEM[mem_addr+i+31:mem_addr+i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_udivrem_epi32'.
// Requires SSE.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func UdivremEpi32(mem_addr *x86.M128i, a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// UnpackhiPs: Unpack and interleave single-precision (32-bit) floating-point
// elements from the high half 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_HIGH_DWORDS(src1[127:0], src2[127:0]){
//			dst[31:0] := src1[95:64] 
//			dst[63:32] := src2[95:64] 
//			dst[95:64] := src1[127:96] 
//			dst[127:96] := src2[127:96] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_DWORDS(a[127:0], b[127:0])
//
// Instruction: 'UNPCKHPS'. Intrinsic: '_mm_unpackhi_ps'.
// Requires SSE.
func UnpackhiPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// UnpackloPs: Unpack and interleave single-precision (32-bit) floating-point
// elements from the low half of 'a' and 'b', and store the results in 'dst'. 
//
//		INTERLEAVE_DWORDS(src1[127:0], src2[127:0]){
//			dst[31:0] := src1[31:0] 
//			dst[63:32] := src2[31:0] 
//			dst[95:64] := src1[63:32] 
//			dst[127:96] := src2[63:32] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_DWORDS(a[127:0], b[127:0])
//
// Instruction: 'UNPCKLPS'. Intrinsic: '_mm_unpacklo_ps'.
// Requires SSE.
func UnpackloPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}


// UremEpi32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: '...'. Intrinsic: '_mm_urem_epi32'.
// Requires SSE.
func UremEpi32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// XorPs: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//		ENDFOR
//
// Instruction: 'XORPS'. Intrinsic: '_mm_xor_ps'.
// Requires SSE.
func XorPs(a x86.M128, b x86.M128) (dst x86.M128) {
	panic("not implemented")
}

