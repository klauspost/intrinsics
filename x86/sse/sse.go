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
func AcosPd(a x86.M128d) x86.M128d {
	return x86.M128d(acosPd([2]float64(a)))
}

func acosPd(a [2]float64) [2]float64


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
func AcosPs(a x86.M128) x86.M128 {
	return x86.M128(acosPs([4]float32(a)))
}

func acosPs(a [4]float32) [4]float32


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
func AcoshPd(a x86.M128d) x86.M128d {
	return x86.M128d(acoshPd([2]float64(a)))
}

func acoshPd(a [2]float64) [2]float64


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
func AcoshPs(a x86.M128) x86.M128 {
	return x86.M128(acoshPs([4]float32(a)))
}

func acoshPs(a [4]float32) [4]float32


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
func AddPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(addPs([4]float32(a), [4]float32(b)))
}

func addPs(a [4]float32, b [4]float32) [4]float32


// AddSs: Add the lower single-precision (32-bit) floating-point element in 'a'
// and 'b', store the result in the lower element of 'dst', and copy the upper
// 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0] + b[31:0]
//		dst[127:32] := a[127:32]
//
// Instruction: 'ADDSS'. Intrinsic: '_mm_add_ss'.
// Requires SSE.
func AddSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(addSs([4]float32(a), [4]float32(b)))
}

func addSs(a [4]float32, b [4]float32) [4]float32


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
func AndPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(andPs([4]float32(a), [4]float32(b)))
}

func andPs(a [4]float32, b [4]float32) [4]float32


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
func AndnotPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(andnotPs([4]float32(a), [4]float32(b)))
}

func andnotPs(a [4]float32, b [4]float32) [4]float32


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
func AsinPd(a x86.M128d) x86.M128d {
	return x86.M128d(asinPd([2]float64(a)))
}

func asinPd(a [2]float64) [2]float64


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
func AsinPs(a x86.M128) x86.M128 {
	return x86.M128(asinPs([4]float32(a)))
}

func asinPs(a [4]float32) [4]float32


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
func AsinhPd(a x86.M128d) x86.M128d {
	return x86.M128d(asinhPd([2]float64(a)))
}

func asinhPd(a [2]float64) [2]float64


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
func AsinhPs(a x86.M128) x86.M128 {
	return x86.M128(asinhPs([4]float32(a)))
}

func asinhPs(a [4]float32) [4]float32


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
func AtanPd(a x86.M128d) x86.M128d {
	return x86.M128d(atanPd([2]float64(a)))
}

func atanPd(a [2]float64) [2]float64


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
func AtanPs(a x86.M128) x86.M128 {
	return x86.M128(atanPs([4]float32(a)))
}

func atanPs(a [4]float32) [4]float32


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
func Atan2Pd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(atan2Pd([2]float64(a), [2]float64(b)))
}

func atan2Pd(a [2]float64, b [2]float64) [2]float64


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
func Atan2Ps(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(atan2Ps([4]float32(a), [4]float32(b)))
}

func atan2Ps(a [4]float32, b [4]float32) [4]float32


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
func AtanhPd(a x86.M128d) x86.M128d {
	return x86.M128d(atanhPd([2]float64(a)))
}

func atanhPd(a [2]float64) [2]float64


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
func AtanhPs(a x86.M128) x86.M128 {
	return x86.M128(atanhPs([4]float32(a)))
}

func atanhPs(a [4]float32) [4]float32


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
func AvgPu16(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(avgPu16(a, b))
}

func avgPu16(a x86.M64, b x86.M64) x86.M64


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
func AvgPu8(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(avgPu8(a, b))
}

func avgPu8(a x86.M64, b x86.M64) x86.M64


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
func CbrtPd(a x86.M128d) x86.M128d {
	return x86.M128d(cbrtPd([2]float64(a)))
}

func cbrtPd(a [2]float64) [2]float64


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
func CbrtPs(a x86.M128) x86.M128 {
	return x86.M128(cbrtPs([4]float32(a)))
}

func cbrtPs(a [4]float32) [4]float32


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
func CdfnormPd(a x86.M128d) x86.M128d {
	return x86.M128d(cdfnormPd([2]float64(a)))
}

func cdfnormPd(a [2]float64) [2]float64


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
func CdfnormPs(a x86.M128) x86.M128 {
	return x86.M128(cdfnormPs([4]float32(a)))
}

func cdfnormPs(a [4]float32) [4]float32


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
func CdfnorminvPd(a x86.M128d) x86.M128d {
	return x86.M128d(cdfnorminvPd([2]float64(a)))
}

func cdfnorminvPd(a [2]float64) [2]float64


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
func CdfnorminvPs(a x86.M128) x86.M128 {
	return x86.M128(cdfnorminvPs([4]float32(a)))
}

func cdfnorminvPs(a [4]float32) [4]float32


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
func CexpPs(a x86.M128) x86.M128 {
	return x86.M128(cexpPs([4]float32(a)))
}

func cexpPs(a [4]float32) [4]float32


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
func ClogPs(a x86.M128) x86.M128 {
	return x86.M128(clogPs([4]float32(a)))
}

func clogPs(a [4]float32) [4]float32


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
func CmpeqPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpeqPs([4]float32(a), [4]float32(b)))
}

func cmpeqPs(a [4]float32, b [4]float32) [4]float32


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
func CmpeqSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpeqSs([4]float32(a), [4]float32(b)))
}

func cmpeqSs(a [4]float32, b [4]float32) [4]float32


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
func CmpgePs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpgePs([4]float32(a), [4]float32(b)))
}

func cmpgePs(a [4]float32, b [4]float32) [4]float32


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
func CmpgeSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpgeSs([4]float32(a), [4]float32(b)))
}

func cmpgeSs(a [4]float32, b [4]float32) [4]float32


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
func CmpgtPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpgtPs([4]float32(a), [4]float32(b)))
}

func cmpgtPs(a [4]float32, b [4]float32) [4]float32


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
func CmpgtSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpgtSs([4]float32(a), [4]float32(b)))
}

func cmpgtSs(a [4]float32, b [4]float32) [4]float32


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
func CmplePs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmplePs([4]float32(a), [4]float32(b)))
}

func cmplePs(a [4]float32, b [4]float32) [4]float32


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
func CmpleSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpleSs([4]float32(a), [4]float32(b)))
}

func cmpleSs(a [4]float32, b [4]float32) [4]float32


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
func CmpltPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpltPs([4]float32(a), [4]float32(b)))
}

func cmpltPs(a [4]float32, b [4]float32) [4]float32


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
func CmpltSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpltSs([4]float32(a), [4]float32(b)))
}

func cmpltSs(a [4]float32, b [4]float32) [4]float32


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
func CmpneqPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpneqPs([4]float32(a), [4]float32(b)))
}

func cmpneqPs(a [4]float32, b [4]float32) [4]float32


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
func CmpneqSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpneqSs([4]float32(a), [4]float32(b)))
}

func cmpneqSs(a [4]float32, b [4]float32) [4]float32


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
func CmpngePs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpngePs([4]float32(a), [4]float32(b)))
}

func cmpngePs(a [4]float32, b [4]float32) [4]float32


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
func CmpngeSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpngeSs([4]float32(a), [4]float32(b)))
}

func cmpngeSs(a [4]float32, b [4]float32) [4]float32


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
func CmpngtPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpngtPs([4]float32(a), [4]float32(b)))
}

func cmpngtPs(a [4]float32, b [4]float32) [4]float32


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
func CmpngtSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpngtSs([4]float32(a), [4]float32(b)))
}

func cmpngtSs(a [4]float32, b [4]float32) [4]float32


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
func CmpnlePs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpnlePs([4]float32(a), [4]float32(b)))
}

func cmpnlePs(a [4]float32, b [4]float32) [4]float32


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
func CmpnleSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpnleSs([4]float32(a), [4]float32(b)))
}

func cmpnleSs(a [4]float32, b [4]float32) [4]float32


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
func CmpnltPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpnltPs([4]float32(a), [4]float32(b)))
}

func cmpnltPs(a [4]float32, b [4]float32) [4]float32


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
func CmpnltSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpnltSs([4]float32(a), [4]float32(b)))
}

func cmpnltSs(a [4]float32, b [4]float32) [4]float32


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
func CmpordPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpordPs([4]float32(a), [4]float32(b)))
}

func cmpordPs(a [4]float32, b [4]float32) [4]float32


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
func CmpordSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpordSs([4]float32(a), [4]float32(b)))
}

func cmpordSs(a [4]float32, b [4]float32) [4]float32


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
func CmpunordPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpunordPs([4]float32(a), [4]float32(b)))
}

func cmpunordPs(a [4]float32, b [4]float32) [4]float32


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
func CmpunordSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(cmpunordSs([4]float32(a), [4]float32(b)))
}

func cmpunordSs(a [4]float32, b [4]float32) [4]float32


// ComieqSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for equality, and return the boolean result (0 or 1). 
//
//		RETURN ( a[31:0] == b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comieq_ss'.
// Requires SSE.
func ComieqSs(a x86.M128, b x86.M128) int {
	return int(comieqSs([4]float32(a), [4]float32(b)))
}

func comieqSs(a [4]float32, b [4]float32) int


// ComigeSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for greater-than-or-equal, and return the boolean result (0
// or 1). 
//
//		RETURN ( a[31:0] >= b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comige_ss'.
// Requires SSE.
func ComigeSs(a x86.M128, b x86.M128) int {
	return int(comigeSs([4]float32(a), [4]float32(b)))
}

func comigeSs(a [4]float32, b [4]float32) int


// ComigtSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for greater-than, and return the boolean result (0 or 1). 
//
//		RETURN ( a[31:0] > b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comigt_ss'.
// Requires SSE.
func ComigtSs(a x86.M128, b x86.M128) int {
	return int(comigtSs([4]float32(a), [4]float32(b)))
}

func comigtSs(a [4]float32, b [4]float32) int


// ComileSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for less-than-or-equal, and return the boolean result (0 or
// 1). 
//
//		RETURN ( a[31:0] <= b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comile_ss'.
// Requires SSE.
func ComileSs(a x86.M128, b x86.M128) int {
	return int(comileSs([4]float32(a), [4]float32(b)))
}

func comileSs(a [4]float32, b [4]float32) int


// ComiltSs: Compare the lower single-precision (32-bit) floating-point element
// in 'a' and 'b' for less-than, and return the boolean result (0 or 1). 
//
//		RETURN ( a[31:0] < b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comilt_ss'.
// Requires SSE.
func ComiltSs(a x86.M128, b x86.M128) int {
	return int(comiltSs([4]float32(a), [4]float32(b)))
}

func comiltSs(a [4]float32, b [4]float32) int


// ComineqSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for not-equal, and return the boolean result (0 or
// 1). 
//
//		RETURN ( a[31:0] != b[31:0] ) ? 1 : 0
//
// Instruction: 'COMISS'. Intrinsic: '_mm_comineq_ss'.
// Requires SSE.
func ComineqSs(a x86.M128, b x86.M128) int {
	return int(comineqSs([4]float32(a), [4]float32(b)))
}

func comineqSs(a [4]float32, b [4]float32) int


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
func CosPd(a x86.M128d) x86.M128d {
	return x86.M128d(cosPd([2]float64(a)))
}

func cosPd(a [2]float64) [2]float64


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
func CosPs(a x86.M128) x86.M128 {
	return x86.M128(cosPs([4]float32(a)))
}

func cosPs(a [4]float32) [4]float32


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
func CosdPd(a x86.M128d) x86.M128d {
	return x86.M128d(cosdPd([2]float64(a)))
}

func cosdPd(a [2]float64) [2]float64


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
func CosdPs(a x86.M128) x86.M128 {
	return x86.M128(cosdPs([4]float32(a)))
}

func cosdPs(a [4]float32) [4]float32


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
func CoshPd(a x86.M128d) x86.M128d {
	return x86.M128d(coshPd([2]float64(a)))
}

func coshPd(a [2]float64) [2]float64


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
func CoshPs(a x86.M128) x86.M128 {
	return x86.M128(coshPs([4]float32(a)))
}

func coshPs(a [4]float32) [4]float32


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
func CsqrtPs(a x86.M128) x86.M128 {
	return x86.M128(csqrtPs([4]float32(a)))
}

func csqrtPs(a [4]float32) [4]float32


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
func CvtPi2ps(a x86.M128, b x86.M64) x86.M128 {
	return x86.M128(cvtPi2ps([4]float32(a), b))
}

func cvtPi2ps(a [4]float32, b x86.M64) [4]float32


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
func CvtPs2pi(a x86.M128) x86.M64 {
	return x86.M64(cvtPs2pi([4]float32(a)))
}

func cvtPs2pi(a [4]float32) x86.M64


// CvtSi2ss: Convert the 32-bit integer 'b' to a single-precision (32-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'CVTSI2SS'. Intrinsic: '_mm_cvt_si2ss'.
// Requires SSE.
func CvtSi2ss(a x86.M128, b int) x86.M128 {
	return x86.M128(cvtSi2ss([4]float32(a), b))
}

func cvtSi2ss(a [4]float32, b int) [4]float32


// CvtSs2si: Convert the lower single-precision (32-bit) floating-point element
// in 'a' to a 32-bit integer, and store the result in 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32(a[31:0])
//
// Instruction: 'CVTSS2SI'. Intrinsic: '_mm_cvt_ss2si'.
// Requires SSE.
func CvtSs2si(a x86.M128) int {
	return int(cvtSs2si([4]float32(a)))
}

func cvtSs2si(a [4]float32) int


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
func Cvtpi16Ps(a x86.M64) x86.M128 {
	return x86.M128(cvtpi16Ps(a))
}

func cvtpi16Ps(a x86.M64) [4]float32


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
func Cvtpi32Ps(a x86.M128, b x86.M64) x86.M128 {
	return x86.M128(cvtpi32Ps([4]float32(a), b))
}

func cvtpi32Ps(a [4]float32, b x86.M64) [4]float32


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
func Cvtpi32x2Ps(a x86.M64, b x86.M64) x86.M128 {
	return x86.M128(cvtpi32x2Ps(a, b))
}

func cvtpi32x2Ps(a x86.M64, b x86.M64) [4]float32


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
func Cvtpi8Ps(a x86.M64) x86.M128 {
	return x86.M128(cvtpi8Ps(a))
}

func cvtpi8Ps(a x86.M64) [4]float32


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
func CvtpsPi16(a x86.M128) x86.M64 {
	return x86.M64(cvtpsPi16([4]float32(a)))
}

func cvtpsPi16(a [4]float32) x86.M64


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
func CvtpsPi32(a x86.M128) x86.M64 {
	return x86.M64(cvtpsPi32([4]float32(a)))
}

func cvtpsPi32(a [4]float32) x86.M64


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
func CvtpsPi8(a x86.M128) x86.M64 {
	return x86.M64(cvtpsPi8([4]float32(a)))
}

func cvtpsPi8(a [4]float32) x86.M64


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
func Cvtpu16Ps(a x86.M64) x86.M128 {
	return x86.M128(cvtpu16Ps(a))
}

func cvtpu16Ps(a x86.M64) [4]float32


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
func Cvtpu8Ps(a x86.M64) x86.M128 {
	return x86.M128(cvtpu8Ps(a))
}

func cvtpu8Ps(a x86.M64) [4]float32


// Cvtsi32Ss: Convert the 32-bit integer 'b' to a single-precision (32-bit)
// floating-point element, store the result in the lower element of 'dst', and
// copy the upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := Convert_Int32_To_FP32(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'CVTSI2SS'. Intrinsic: '_mm_cvtsi32_ss'.
// Requires SSE.
func Cvtsi32Ss(a x86.M128, b int) x86.M128 {
	return x86.M128(cvtsi32Ss([4]float32(a), b))
}

func cvtsi32Ss(a [4]float32, b int) [4]float32


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
func Cvtsi64Ss(a x86.M128, b int64) x86.M128 {
	return x86.M128(cvtsi64Ss([4]float32(a), b))
}

func cvtsi64Ss(a [4]float32, b int64) [4]float32


// CvtssF32: Copy the lower single-precision (32-bit) floating-point element of
// 'a' to 'dst'. 
//
//		dst[31:0] := a[31:0]
//
// Instruction: 'MOVSS'. Intrinsic: '_mm_cvtss_f32'.
// Requires SSE.
func CvtssF32(a x86.M128) float32 {
	return float32(cvtssF32([4]float32(a)))
}

func cvtssF32(a [4]float32) float32


// CvtssSi32: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 32-bit integer, and store the result in 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32(a[31:0])
//
// Instruction: 'CVTSS2SI'. Intrinsic: '_mm_cvtss_si32'.
// Requires SSE.
func CvtssSi32(a x86.M128) int {
	return int(cvtssSi32([4]float32(a)))
}

func cvtssSi32(a [4]float32) int


// CvtssSi64: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 64-bit integer, and store the result in 'dst'. 
//
//		dst[63:0] := Convert_FP32_To_Int64(a[31:0])
//
// Instruction: 'CVTSS2SI'. Intrinsic: '_mm_cvtss_si64'.
// Requires SSE.
func CvtssSi64(a x86.M128) int64 {
	return int64(cvtssSi64([4]float32(a)))
}

func cvtssSi64(a [4]float32) int64


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
func CvttPs2pi(a x86.M128) x86.M64 {
	return x86.M64(cvttPs2pi([4]float32(a)))
}

func cvttPs2pi(a [4]float32) x86.M64


// CvttSs2si: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 32-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32_Truncate(a[31:0])
//
// Instruction: 'CVTTSS2SI'. Intrinsic: '_mm_cvtt_ss2si'.
// Requires SSE.
func CvttSs2si(a x86.M128) int {
	return int(cvttSs2si([4]float32(a)))
}

func cvttSs2si(a [4]float32) int


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
func CvttpsPi32(a x86.M128) x86.M64 {
	return x86.M64(cvttpsPi32([4]float32(a)))
}

func cvttpsPi32(a [4]float32) x86.M64


// CvttssSi32: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 32-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP32_To_Int32_Truncate(a[31:0])
//
// Instruction: 'CVTTSS2SI'. Intrinsic: '_mm_cvttss_si32'.
// Requires SSE.
func CvttssSi32(a x86.M128) int {
	return int(cvttssSi32([4]float32(a)))
}

func cvttssSi32(a [4]float32) int


// CvttssSi64: Convert the lower single-precision (32-bit) floating-point
// element in 'a' to a 64-bit integer with truncation, and store the result in
// 'dst'. 
//
//		dst[63:0] := Convert_FP64_To_Int32_Truncate(a[31:0])
//
// Instruction: 'CVTTSS2SI'. Intrinsic: '_mm_cvttss_si64'.
// Requires SSE.
func CvttssSi64(a x86.M128) int64 {
	return int64(cvttssSi64([4]float32(a)))
}

func cvttssSi64(a [4]float32) int64


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
func DivEpi16(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpi16([16]byte(a), [16]byte(b)))
}

func divEpi16(a [16]byte, b [16]byte) [16]byte


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
func DivEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpi32([16]byte(a), [16]byte(b)))
}

func divEpi32(a [16]byte, b [16]byte) [16]byte


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
func DivEpi64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpi64([16]byte(a), [16]byte(b)))
}

func divEpi64(a [16]byte, b [16]byte) [16]byte


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
func DivEpi8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpi8([16]byte(a), [16]byte(b)))
}

func divEpi8(a [16]byte, b [16]byte) [16]byte


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
func DivEpu16(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpu16([16]byte(a), [16]byte(b)))
}

func divEpu16(a [16]byte, b [16]byte) [16]byte


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
func DivEpu32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpu32([16]byte(a), [16]byte(b)))
}

func divEpu32(a [16]byte, b [16]byte) [16]byte


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
func DivEpu64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpu64([16]byte(a), [16]byte(b)))
}

func divEpu64(a [16]byte, b [16]byte) [16]byte


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
func DivEpu8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(divEpu8([16]byte(a), [16]byte(b)))
}

func divEpu8(a [16]byte, b [16]byte) [16]byte


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
func DivPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(divPs([4]float32(a), [4]float32(b)))
}

func divPs(a [4]float32, b [4]float32) [4]float32


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
func DivSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(divSs([4]float32(a), [4]float32(b)))
}

func divSs(a [4]float32, b [4]float32) [4]float32


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
func ErfPd(a x86.M128d) x86.M128d {
	return x86.M128d(erfPd([2]float64(a)))
}

func erfPd(a [2]float64) [2]float64


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
func ErfPs(a x86.M128) x86.M128 {
	return x86.M128(erfPs([4]float32(a)))
}

func erfPs(a [4]float32) [4]float32


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
func ErfcPd(a x86.M128d) x86.M128d {
	return x86.M128d(erfcPd([2]float64(a)))
}

func erfcPd(a [2]float64) [2]float64


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
func ErfcPs(a x86.M128) x86.M128 {
	return x86.M128(erfcPs([4]float32(a)))
}

func erfcPs(a [4]float32) [4]float32


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
func ErfcinvPd(a x86.M128d) x86.M128d {
	return x86.M128d(erfcinvPd([2]float64(a)))
}

func erfcinvPd(a [2]float64) [2]float64


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
func ErfcinvPs(a x86.M128) x86.M128 {
	return x86.M128(erfcinvPs([4]float32(a)))
}

func erfcinvPs(a [4]float32) [4]float32


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
func ErfinvPd(a x86.M128d) x86.M128d {
	return x86.M128d(erfinvPd([2]float64(a)))
}

func erfinvPd(a [2]float64) [2]float64


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
func ErfinvPs(a x86.M128) x86.M128 {
	return x86.M128(erfinvPs([4]float32(a)))
}

func erfinvPs(a [4]float32) [4]float32


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
func ExpPd(a x86.M128d) x86.M128d {
	return x86.M128d(expPd([2]float64(a)))
}

func expPd(a [2]float64) [2]float64


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
func ExpPs(a x86.M128) x86.M128 {
	return x86.M128(expPs([4]float32(a)))
}

func expPs(a [4]float32) [4]float32


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
func Exp10Pd(a x86.M128d) x86.M128d {
	return x86.M128d(exp10Pd([2]float64(a)))
}

func exp10Pd(a [2]float64) [2]float64


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
func Exp10Ps(a x86.M128) x86.M128 {
	return x86.M128(exp10Ps([4]float32(a)))
}

func exp10Ps(a [4]float32) [4]float32


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
func Exp2Pd(a x86.M128d) x86.M128d {
	return x86.M128d(exp2Pd([2]float64(a)))
}

func exp2Pd(a [2]float64) [2]float64


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
func Exp2Ps(a x86.M128) x86.M128 {
	return x86.M128(exp2Ps([4]float32(a)))
}

func exp2Ps(a [4]float32) [4]float32


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
func Expm1Pd(a x86.M128d) x86.M128d {
	return x86.M128d(expm1Pd([2]float64(a)))
}

func expm1Pd(a [2]float64) [2]float64


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
func Expm1Ps(a x86.M128) x86.M128 {
	return x86.M128(expm1Ps([4]float32(a)))
}

func expm1Ps(a [4]float32) [4]float32


// ExtractPi16: Extract a 16-bit integer from 'a', selected with 'imm8', and
// store the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_mm_extract_pi16'.
// Requires SSE.
func ExtractPi16(a x86.M64, imm8 int) int {
	return int(extractPi16(a, imm8))
}

func extractPi16(a x86.M64, imm8 int) int


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
	return uint32(mMGETEXCEPTIONMASK())
}

func mMGETEXCEPTIONMASK() uint32


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
	return uint32(mMGETEXCEPTIONSTATE())
}

func mMGETEXCEPTIONSTATE() uint32


// MMGETFLUSHZEROMODE: Macro: Get the flush zero bits from the MXCSR control
// and status register. The flush zero may contain any of the following flags:
// _MM_FLUSH_ZERO_ON or _MM_FLUSH_ZERO_OFF 
//
//		dst[31:0] := MXCSR & _MM_FLUSH_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_FLUSH_ZERO_MODE'.
// Requires SSE.
func MMGETFLUSHZEROMODE() uint32 {
	return uint32(mMGETFLUSHZEROMODE())
}

func mMGETFLUSHZEROMODE() uint32


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
	return uint32(mMGETROUNDINGMODE())
}

func mMGETROUNDINGMODE() uint32


// Getcsr: Get the unsigned 32-bit value of the MXCSR control and status
// register. 
//
//		dst[31:0] := MXCSR
//
// Instruction: 'STMXCSR'. Intrinsic: '_mm_getcsr'.
// Requires SSE.
func Getcsr() uint32 {
	return uint32(getcsr())
}

func getcsr() uint32


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
func HypotPd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(hypotPd([2]float64(a), [2]float64(b)))
}

func hypotPd(a [2]float64, b [2]float64) [2]float64


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
func HypotPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(hypotPs([4]float32(a), [4]float32(b)))
}

func hypotPs(a [4]float32, b [4]float32) [4]float32


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
func IdivEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(idivEpi32([16]byte(a), [16]byte(b)))
}

func idivEpi32(a [16]byte, b [16]byte) [16]byte


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
// FIXME: Will likely need to be reworked.
func IdivremEpi32(mem_addr *x86.M128i, a x86.M128i, b x86.M128i) x86.M128i {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128i{}
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
func InsertPi16(a x86.M64, i int, imm8 int) x86.M64 {
	return x86.M64(insertPi16(a, i, imm8))
}

func insertPi16(a x86.M64, i int, imm8 int) x86.M64


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
func InvcbrtPd(a x86.M128d) x86.M128d {
	return x86.M128d(invcbrtPd([2]float64(a)))
}

func invcbrtPd(a [2]float64) [2]float64


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
func InvcbrtPs(a x86.M128) x86.M128 {
	return x86.M128(invcbrtPs([4]float32(a)))
}

func invcbrtPs(a [4]float32) [4]float32


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
func InvsqrtPd(a x86.M128d) x86.M128d {
	return x86.M128d(invsqrtPd([2]float64(a)))
}

func invsqrtPd(a [2]float64) [2]float64


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
func InvsqrtPs(a x86.M128) x86.M128 {
	return x86.M128(invsqrtPs([4]float32(a)))
}

func invsqrtPs(a [4]float32) [4]float32


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
func IremEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(iremEpi32([16]byte(a), [16]byte(b)))
}

func iremEpi32(a [16]byte, b [16]byte) [16]byte


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
// FIXME: Will likely need to be reworked.
func LoadhPi(a x86.M128, mem_addr *x86.M64Const) x86.M128 {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128{}
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
// FIXME: Will likely need to be reworked.
func LoadlPi(a x86.M128, mem_addr *x86.M64Const) x86.M128 {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128{}
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
func LogPd(a x86.M128d) x86.M128d {
	return x86.M128d(logPd([2]float64(a)))
}

func logPd(a [2]float64) [2]float64


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
func LogPs(a x86.M128) x86.M128 {
	return x86.M128(logPs([4]float32(a)))
}

func logPs(a [4]float32) [4]float32


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
func Log10Pd(a x86.M128d) x86.M128d {
	return x86.M128d(log10Pd([2]float64(a)))
}

func log10Pd(a [2]float64) [2]float64


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
func Log10Ps(a x86.M128) x86.M128 {
	return x86.M128(log10Ps([4]float32(a)))
}

func log10Ps(a [4]float32) [4]float32


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
func Log1pPd(a x86.M128d) x86.M128d {
	return x86.M128d(log1pPd([2]float64(a)))
}

func log1pPd(a [2]float64) [2]float64


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
func Log1pPs(a x86.M128) x86.M128 {
	return x86.M128(log1pPs([4]float32(a)))
}

func log1pPs(a [4]float32) [4]float32


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
func Log2Pd(a x86.M128d) x86.M128d {
	return x86.M128d(log2Pd([2]float64(a)))
}

func log2Pd(a [2]float64) [2]float64


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
func Log2Ps(a x86.M128) x86.M128 {
	return x86.M128(log2Ps([4]float32(a)))
}

func log2Ps(a [4]float32) [4]float32


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
func LogbPd(a x86.M128d) x86.M128d {
	return x86.M128d(logbPd([2]float64(a)))
}

func logbPd(a [2]float64) [2]float64


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
func LogbPs(a x86.M128) x86.M128 {
	return x86.M128(logbPs([4]float32(a)))
}

func logbPs(a [4]float32) [4]float32


// Malloc: Allocate 'size' bytes of memory, aligned to the alignment specified
// in 'align', and return a pointer to the allocated memory. '_mm_free' should
// be used to free memory that is allocated with '_mm_malloc'. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_malloc'.
// Requires SSE.
func Malloc(size int, align int)  {
	malloc(size, align)
}

func malloc(size int, align int) 


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
// FIXME: Will likely need to be reworked.
func MaskmoveSi64(a x86.M64, mask x86.M64, mem_addr *byte)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
// FIXME: Will likely need to be reworked.
func Maskmovq(a x86.M64, mask x86.M64, mem_addr *byte)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
func MaxPi16(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(maxPi16(a, b))
}

func maxPi16(a x86.M64, b x86.M64) x86.M64


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
func MaxPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maxPs([4]float32(a), [4]float32(b)))
}

func maxPs(a [4]float32, b [4]float32) [4]float32


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
func MaxPu8(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(maxPu8(a, b))
}

func maxPu8(a x86.M64, b x86.M64) x86.M64


// MaxSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b', store the maximum value in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[31:0] := MAX(a[31:0], b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'MAXSS'. Intrinsic: '_mm_max_ss'.
// Requires SSE.
func MaxSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(maxSs([4]float32(a), [4]float32(b)))
}

func maxSs(a [4]float32, b [4]float32) [4]float32


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
func MinPi16(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(minPi16(a, b))
}

func minPi16(a x86.M64, b x86.M64) x86.M64


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
func MinPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(minPs([4]float32(a), [4]float32(b)))
}

func minPs(a [4]float32, b [4]float32) [4]float32


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
func MinPu8(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(minPu8(a, b))
}

func minPu8(a x86.M64, b x86.M64) x86.M64


// MinSs: Compare the lower single-precision (32-bit) floating-point elements
// in 'a' and 'b', store the minimum value in the lower element of 'dst', and
// copy the upper element from 'a' to the upper element of 'dst'. 
//
//		dst[31:0] := MIN(a[31:0], b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'MINSS'. Intrinsic: '_mm_min_ss'.
// Requires SSE.
func MinSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(minSs([4]float32(a), [4]float32(b)))
}

func minSs(a [4]float32, b [4]float32) [4]float32


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
func MoveSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(moveSs([4]float32(a), [4]float32(b)))
}

func moveSs(a [4]float32, b [4]float32) [4]float32


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
func MovehlPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(movehlPs([4]float32(a), [4]float32(b)))
}

func movehlPs(a [4]float32, b [4]float32) [4]float32


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
func MovelhPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(movelhPs([4]float32(a), [4]float32(b)))
}

func movelhPs(a [4]float32, b [4]float32) [4]float32


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
	return int(movemaskPi8(a))
}

func movemaskPi8(a x86.M64) int


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
	return int(movemaskPs([4]float32(a)))
}

func movemaskPs(a [4]float32) int


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
func MulPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(mulPs([4]float32(a), [4]float32(b)))
}

func mulPs(a [4]float32, b [4]float32) [4]float32


// MulSs: Multiply the lower single-precision (32-bit) floating-point element
// in 'a' and 'b', store the result in the lower element of 'dst', and copy the
// upper 3 packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := a[31:0] * b[31:0]
//		dst[127:32] := a[127:32]
//
// Instruction: 'MULSS'. Intrinsic: '_mm_mul_ss'.
// Requires SSE.
func MulSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(mulSs([4]float32(a), [4]float32(b)))
}

func mulSs(a [4]float32, b [4]float32) [4]float32


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
func MulhiPu16(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(mulhiPu16(a, b))
}

func mulhiPu16(a x86.M64, b x86.M64) x86.M64


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
func OrPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(orPs([4]float32(a), [4]float32(b)))
}

func orPs(a [4]float32, b [4]float32) [4]float32


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
func Pavgb(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pavgb(a, b))
}

func pavgb(a x86.M64, b x86.M64) x86.M64


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
func Pavgw(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pavgw(a, b))
}

func pavgw(a x86.M64, b x86.M64) x86.M64


// Pextrw: Extract a 16-bit integer from 'a', selected with 'imm8', and store
// the result in the lower element of 'dst'. 
//
//		dst[15:0] := (a[63:0] >> (imm8[1:0] * 16))[15:0]
//		dst[31:16] := 0
//
// Instruction: 'PEXTRW'. Intrinsic: '_m_pextrw'.
// Requires SSE.
func Pextrw(a x86.M64, imm8 int) int {
	return int(pextrw(a, imm8))
}

func pextrw(a x86.M64, imm8 int) int


// Pinsrw: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[63:0] := a[63:0]
//		sel := imm8[1:0]*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: 'PINSRW'. Intrinsic: '_m_pinsrw'.
// Requires SSE.
func Pinsrw(a x86.M64, i int, imm8 int) x86.M64 {
	return x86.M64(pinsrw(a, i, imm8))
}

func pinsrw(a x86.M64, i int, imm8 int) x86.M64


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
func Pmaxsw(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pmaxsw(a, b))
}

func pmaxsw(a x86.M64, b x86.M64) x86.M64


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
func Pmaxub(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pmaxub(a, b))
}

func pmaxub(a x86.M64, b x86.M64) x86.M64


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
func Pminsw(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pminsw(a, b))
}

func pminsw(a x86.M64, b x86.M64) x86.M64


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
func Pminub(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pminub(a, b))
}

func pminub(a x86.M64, b x86.M64) x86.M64


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
	return int(pmovmskb(a))
}

func pmovmskb(a x86.M64) int


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
func Pmulhuw(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(pmulhuw(a, b))
}

func pmulhuw(a x86.M64, b x86.M64) x86.M64


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
func PowPd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(powPd([2]float64(a), [2]float64(b)))
}

func powPd(a [2]float64, b [2]float64) [2]float64


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
func PowPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(powPs([4]float32(a), [4]float32(b)))
}

func powPs(a [4]float32, b [4]float32) [4]float32


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'PREFETCHNTA, PREFETCHT0, PREFETCHT1, PREFETCHT2'. Intrinsic: '_mm_prefetch'.
// Requires SSE.
// FIXME: Will likely need to be reworked.
func Prefetch(p *byte, i int)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
func Psadbw(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(psadbw(a, b))
}

func psadbw(a x86.M64, b x86.M64) x86.M64


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
func Pshufw(a x86.M64, imm8 int) x86.M64 {
	return x86.M64(pshufw(a, imm8))
}

func pshufw(a x86.M64, imm8 int) x86.M64


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
func RcpPs(a x86.M128) x86.M128 {
	return x86.M128(rcpPs([4]float32(a)))
}

func rcpPs(a [4]float32) [4]float32


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
func RcpSs(a x86.M128) x86.M128 {
	return x86.M128(rcpSs([4]float32(a)))
}

func rcpSs(a [4]float32) [4]float32


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
func RemEpi16(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpi16([16]byte(a), [16]byte(b)))
}

func remEpi16(a [16]byte, b [16]byte) [16]byte


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
func RemEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpi32([16]byte(a), [16]byte(b)))
}

func remEpi32(a [16]byte, b [16]byte) [16]byte


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
func RemEpi64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpi64([16]byte(a), [16]byte(b)))
}

func remEpi64(a [16]byte, b [16]byte) [16]byte


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
func RemEpi8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpi8([16]byte(a), [16]byte(b)))
}

func remEpi8(a [16]byte, b [16]byte) [16]byte


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
func RemEpu16(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpu16([16]byte(a), [16]byte(b)))
}

func remEpu16(a [16]byte, b [16]byte) [16]byte


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
func RemEpu32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpu32([16]byte(a), [16]byte(b)))
}

func remEpu32(a [16]byte, b [16]byte) [16]byte


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
func RemEpu64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpu64([16]byte(a), [16]byte(b)))
}

func remEpu64(a [16]byte, b [16]byte) [16]byte


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
func RemEpu8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(remEpu8([16]byte(a), [16]byte(b)))
}

func remEpu8(a [16]byte, b [16]byte) [16]byte


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
func RsqrtPs(a x86.M128) x86.M128 {
	return x86.M128(rsqrtPs([4]float32(a)))
}

func rsqrtPs(a [4]float32) [4]float32


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
func RsqrtSs(a x86.M128) x86.M128 {
	return x86.M128(rsqrtSs([4]float32(a)))
}

func rsqrtSs(a [4]float32) [4]float32


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
func SadPu8(a x86.M64, b x86.M64) x86.M64 {
	return x86.M64(sadPu8(a, b))
}

func sadPu8(a x86.M64, b x86.M64) x86.M64


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
	mMSETEXCEPTIONMASK(a)
}

func mMSETEXCEPTIONMASK(a uint32) 


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
	mMSETEXCEPTIONSTATE(a)
}

func mMSETEXCEPTIONSTATE(a uint32) 


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
	mMSETFLUSHZEROMODE(a)
}

func mMSETFLUSHZEROMODE(a uint32) 


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
func SetPs(e3 float32, e2 float32, e1 float32, e0 float32) x86.M128 {
	return x86.M128(setPs(e3, e2, e1, e0))
}

func setPs(e3 float32, e2 float32, e1 float32, e0 float32) [4]float32


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
func SetPs1(a float32) x86.M128 {
	return x86.M128(setPs1(a))
}

func setPs1(a float32) [4]float32


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
	mMSETROUNDINGMODE(a)
}

func mMSETROUNDINGMODE(a uint32) 


// SetSs: Copy single-precision (32-bit) floating-point element 'a' to the
// lower element of 'dst', and zero the upper 3 elements. 
//
//		dst[31:0] := a[31:0]
//		dst[127:32] := 0
//
// Instruction: '...'. Intrinsic: '_mm_set_ss'.
// Requires SSE.
func SetSs(a float32) x86.M128 {
	return x86.M128(setSs(a))
}

func setSs(a float32) [4]float32


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
func Set1Ps(a float32) x86.M128 {
	return x86.M128(set1Ps(a))
}

func set1Ps(a float32) [4]float32


// Setcsr: Set the MXCSR control and status register with the value in unsigned
// 32-bit integer 'a'. 
//
//		MXCSR := a[31:0]
//
// Instruction: 'LDMXCSR'. Intrinsic: '_mm_setcsr'.
// Requires SSE.
func Setcsr(a uint32)  {
	setcsr(a)
}

func setcsr(a uint32) 


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
func SetrPs(e3 float32, e2 float32, e1 float32, e0 float32) x86.M128 {
	return x86.M128(setrPs(e3, e2, e1, e0))
}

func setrPs(e3 float32, e2 float32, e1 float32, e0 float32) [4]float32


// SetzeroPs: Return vector of type __m128 with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'XORPS'. Intrinsic: '_mm_setzero_ps'.
// Requires SSE.
func SetzeroPs() x86.M128 {
	return x86.M128(setzeroPs())
}

func setzeroPs() [4]float32


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
	sfence()
}

func sfence() 


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
func ShufflePi16(a x86.M64, imm8 int) x86.M64 {
	return x86.M64(shufflePi16(a, imm8))
}

func shufflePi16(a x86.M64, imm8 int) x86.M64


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
func ShufflePs(a x86.M128, b x86.M128, imm8 uint32) x86.M128 {
	return x86.M128(shufflePs([4]float32(a), [4]float32(b), imm8))
}

func shufflePs(a [4]float32, b [4]float32, imm8 uint32) [4]float32


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
func SinPd(a x86.M128d) x86.M128d {
	return x86.M128d(sinPd([2]float64(a)))
}

func sinPd(a [2]float64) [2]float64


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
func SinPs(a x86.M128) x86.M128 {
	return x86.M128(sinPs([4]float32(a)))
}

func sinPs(a [4]float32) [4]float32


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
// FIXME: Will likely need to be reworked.
func SincosPd(mem_addr *x86.M128d, a x86.M128d) x86.M128d {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128d{}
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
// FIXME: Will likely need to be reworked.
func SincosPs(mem_addr *x86.M128, a x86.M128) x86.M128 {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128{}
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
func SindPd(a x86.M128d) x86.M128d {
	return x86.M128d(sindPd([2]float64(a)))
}

func sindPd(a [2]float64) [2]float64


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
func SindPs(a x86.M128) x86.M128 {
	return x86.M128(sindPs([4]float32(a)))
}

func sindPs(a [4]float32) [4]float32


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
func SinhPd(a x86.M128d) x86.M128d {
	return x86.M128d(sinhPd([2]float64(a)))
}

func sinhPd(a [2]float64) [2]float64


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
func SinhPs(a x86.M128) x86.M128 {
	return x86.M128(sinhPs([4]float32(a)))
}

func sinhPs(a [4]float32) [4]float32


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
func SqrtPs(a x86.M128) x86.M128 {
	return x86.M128(sqrtPs([4]float32(a)))
}

func sqrtPs(a [4]float32) [4]float32


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
func SqrtSs(a x86.M128) x86.M128 {
	return x86.M128(sqrtSs([4]float32(a)))
}

func sqrtSs(a [4]float32) [4]float32


// StorePs: Store 128-bits (composed of 4 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVAPS'. Intrinsic: '_mm_store_ps'.
// Requires SSE.
// FIXME: Will likely need to be reworked.
func StorePs(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
// FIXME: Will likely need to be reworked.
func StorePs1(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StoreSs: Store the lower single-precision (32-bit) floating-point element
// from 'a' into memory. 'mem_addr' does not need to be aligned on any
// particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVSS'. Intrinsic: '_mm_store_ss'.
// Requires SSE.
// FIXME: Will likely need to be reworked.
func StoreSs(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
// FIXME: Will likely need to be reworked.
func Store1Ps(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorehPi: Store the upper 2 single-precision (32-bit) floating-point
// elements from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[95:64]
//		MEM[mem_addr+63:mem_addr+32] := a[127:96]
//
// Instruction: 'MOVHPS'. Intrinsic: '_mm_storeh_pi'.
// Requires SSE.
// FIXME: Will likely need to be reworked.
func StorehPi(mem_addr *x86.M64, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StorelPi: Store the lower 2 single-precision (32-bit) floating-point
// elements from 'a' into memory. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//		MEM[mem_addr+63:mem_addr+32] := a[63:32]
//
// Instruction: 'MOVLPS'. Intrinsic: '_mm_storel_pi'.
// Requires SSE.
// FIXME: Will likely need to be reworked.
func StorelPi(mem_addr *x86.M64, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
// FIXME: Will likely need to be reworked.
func StorerPs(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// StoreuPs: Store 128-bits (composed of 4 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+127:mem_addr] := a[127:0]
//
// Instruction: 'MOVUPS'. Intrinsic: '_mm_storeu_ps'.
// Requires SSE.
// FIXME: Will likely need to be reworked.
func StoreuPs(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
// FIXME: Will likely need to be reworked.
func StreamPi(mem_addr *x86.M64, a x86.M64)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
// FIXME: Will likely need to be reworked.
func StreamPs(mem_addr *float32, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

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
func SubPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(subPs([4]float32(a), [4]float32(b)))
}

func subPs(a [4]float32, b [4]float32) [4]float32


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
func SubSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(subSs([4]float32(a), [4]float32(b)))
}

func subSs(a [4]float32, b [4]float32) [4]float32


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
func SvmlCeilPd(a x86.M128d) x86.M128d {
	return x86.M128d(svmlCeilPd([2]float64(a)))
}

func svmlCeilPd(a [2]float64) [2]float64


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
func SvmlCeilPs(a x86.M128) x86.M128 {
	return x86.M128(svmlCeilPs([4]float32(a)))
}

func svmlCeilPs(a [4]float32) [4]float32


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
func SvmlFloorPd(a x86.M128d) x86.M128d {
	return x86.M128d(svmlFloorPd([2]float64(a)))
}

func svmlFloorPd(a [2]float64) [2]float64


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
func SvmlFloorPs(a x86.M128) x86.M128 {
	return x86.M128(svmlFloorPs([4]float32(a)))
}

func svmlFloorPs(a [4]float32) [4]float32


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
func SvmlRoundPd(a x86.M128d) x86.M128d {
	return x86.M128d(svmlRoundPd([2]float64(a)))
}

func svmlRoundPd(a [2]float64) [2]float64


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
func SvmlRoundPs(a x86.M128) x86.M128 {
	return x86.M128(svmlRoundPs([4]float32(a)))
}

func svmlRoundPs(a [4]float32) [4]float32


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
func SvmlSqrtPd(a x86.M128d) x86.M128d {
	return x86.M128d(svmlSqrtPd([2]float64(a)))
}

func svmlSqrtPd(a [2]float64) [2]float64


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
func SvmlSqrtPs(a x86.M128) x86.M128 {
	return x86.M128(svmlSqrtPs([4]float32(a)))
}

func svmlSqrtPs(a [4]float32) [4]float32


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
func TanPd(a x86.M128d) x86.M128d {
	return x86.M128d(tanPd([2]float64(a)))
}

func tanPd(a [2]float64) [2]float64


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
func TanPs(a x86.M128) x86.M128 {
	return x86.M128(tanPs([4]float32(a)))
}

func tanPs(a [4]float32) [4]float32


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
func TandPd(a x86.M128d) x86.M128d {
	return x86.M128d(tandPd([2]float64(a)))
}

func tandPd(a [2]float64) [2]float64


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
func TandPs(a x86.M128) x86.M128 {
	return x86.M128(tandPs([4]float32(a)))
}

func tandPs(a [4]float32) [4]float32


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
func TanhPd(a x86.M128d) x86.M128d {
	return x86.M128d(tanhPd([2]float64(a)))
}

func tanhPd(a [2]float64) [2]float64


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
func TanhPs(a x86.M128) x86.M128 {
	return x86.M128(tanhPs([4]float32(a)))
}

func tanhPs(a [4]float32) [4]float32


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
	mMTRANSPOSE4PS([4]float32(row0), [4]float32(row1), [4]float32(row2), [4]float32(row3))
}

func mMTRANSPOSE4PS(row0 [4]float32, row1 [4]float32, row2 [4]float32, row3 [4]float32) 


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
func TruncPd(a x86.M128d) x86.M128d {
	return x86.M128d(truncPd([2]float64(a)))
}

func truncPd(a [2]float64) [2]float64


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
func TruncPs(a x86.M128) x86.M128 {
	return x86.M128(truncPs([4]float32(a)))
}

func truncPs(a [4]float32) [4]float32


// UcomieqSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for equality, and return the boolean result (0 or 1).
// This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] == b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomieq_ss'.
// Requires SSE.
func UcomieqSs(a x86.M128, b x86.M128) int {
	return int(ucomieqSs([4]float32(a), [4]float32(b)))
}

func ucomieqSs(a [4]float32, b [4]float32) int


// UcomigeSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for greater-than-or-equal, and return the boolean
// result (0 or 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] >= b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomige_ss'.
// Requires SSE.
func UcomigeSs(a x86.M128, b x86.M128) int {
	return int(ucomigeSs([4]float32(a), [4]float32(b)))
}

func ucomigeSs(a [4]float32, b [4]float32) int


// UcomigtSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for greater-than, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] > b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomigt_ss'.
// Requires SSE.
func UcomigtSs(a x86.M128, b x86.M128) int {
	return int(ucomigtSs([4]float32(a), [4]float32(b)))
}

func ucomigtSs(a [4]float32, b [4]float32) int


// UcomileSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for less-than-or-equal, and return the boolean result
// (0 or 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] <= b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomile_ss'.
// Requires SSE.
func UcomileSs(a x86.M128, b x86.M128) int {
	return int(ucomileSs([4]float32(a), [4]float32(b)))
}

func ucomileSs(a [4]float32, b [4]float32) int


// UcomiltSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for less-than, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] < b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomilt_ss'.
// Requires SSE.
func UcomiltSs(a x86.M128, b x86.M128) int {
	return int(ucomiltSs([4]float32(a), [4]float32(b)))
}

func ucomiltSs(a [4]float32, b [4]float32) int


// UcomineqSs: Compare the lower single-precision (32-bit) floating-point
// element in 'a' and 'b' for not-equal, and return the boolean result (0 or
// 1). This instruction will not signal an exception for QNaNs. 
//
//		RETURN ( a[31:0] != b[31:0] ) ? 1 : 0
//
// Instruction: 'UCOMISS'. Intrinsic: '_mm_ucomineq_ss'.
// Requires SSE.
func UcomineqSs(a x86.M128, b x86.M128) int {
	return int(ucomineqSs([4]float32(a), [4]float32(b)))
}

func ucomineqSs(a [4]float32, b [4]float32) int


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
func UdivEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(udivEpi32([16]byte(a), [16]byte(b)))
}

func udivEpi32(a [16]byte, b [16]byte) [16]byte


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
// FIXME: Will likely need to be reworked.
func UdivremEpi32(mem_addr *x86.M128i, a x86.M128i, b x86.M128i) x86.M128i {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M128i{}
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
func UnpackhiPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(unpackhiPs([4]float32(a), [4]float32(b)))
}

func unpackhiPs(a [4]float32, b [4]float32) [4]float32


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
func UnpackloPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(unpackloPs([4]float32(a), [4]float32(b)))
}

func unpackloPs(a [4]float32, b [4]float32) [4]float32


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
func UremEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(uremEpi32([16]byte(a), [16]byte(b)))
}

func uremEpi32(a [16]byte, b [16]byte) [16]byte


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
func XorPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(xorPs([4]float32(a), [4]float32(b)))
}

func xorPs(a [4]float32, b [4]float32) [4]float32

