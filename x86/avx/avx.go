package avx

import . "github.com/klauspost/intrinsics/x86"


// AcosPd: Compute the inverse cosine of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ACOS(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_acos_pd'.
// Requires AVX.
func AcosPd(a M256d) M256d {
	return M256d(acosPd([4]float64(a)))
}

func acosPd(a [4]float64) [4]float64


// AcosPs: Compute the inverse cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ACOS(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_acos_ps'.
// Requires AVX.
func AcosPs(a M256) M256 {
	return M256(acosPs([8]float32(a)))
}

func acosPs(a [8]float32) [8]float32


// AcoshPd: Compute the inverse hyperbolic cosine of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ACOSH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_acosh_pd'.
// Requires AVX.
func AcoshPd(a M256d) M256d {
	return M256d(acoshPd([4]float64(a)))
}

func acoshPd(a [4]float64) [4]float64


// AcoshPs: Compute the inverse hyperbolic cosine of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ACOSH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_acosh_ps'.
// Requires AVX.
func AcoshPs(a M256) M256 {
	return M256(acoshPs([8]float32(a)))
}

func acoshPs(a [8]float32) [8]float32


// AddPd: Add packed double-precision (64-bit) floating-point elements in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm256_add_pd'.
// Requires AVX.
func AddPd(a M256d, b M256d) M256d {
	return M256d(addPd([4]float64(a), [4]float64(b)))
}

func addPd(a [4]float64, b [4]float64) [4]float64


// AddPs: Add packed single-precision (32-bit) floating-point elements in 'a'
// and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm256_add_ps'.
// Requires AVX.
func AddPs(a M256, b M256) M256 {
	return M256(addPs([8]float32(a), [8]float32(b)))
}

func addPs(a [8]float32, b [8]float32) [8]float32


// AddsubPd: Alternatively add and subtract packed double-precision (64-bit)
// floating-point elements in 'a' to/from packed elements in 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF (j is even) 
//				dst[i+63:i] := a[i+63:i] - b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDSUBPD'. Intrinsic: '_mm256_addsub_pd'.
// Requires AVX.
func AddsubPd(a M256d, b M256d) M256d {
	return M256d(addsubPd([4]float64(a), [4]float64(b)))
}

func addsubPd(a [4]float64, b [4]float64) [4]float64


// AddsubPs: Alternatively add and subtract packed single-precision (32-bit)
// floating-point elements in 'a' to/from packed elements in 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF (j is even) 
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDSUBPS'. Intrinsic: '_mm256_addsub_ps'.
// Requires AVX.
func AddsubPs(a M256, b M256) M256 {
	return M256(addsubPs([8]float32(a), [8]float32(b)))
}

func addsubPs(a [8]float32, b [8]float32) [8]float32


// AndPd: Compute the bitwise AND of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] AND b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDPD'. Intrinsic: '_mm256_and_pd'.
// Requires AVX.
func AndPd(a M256d, b M256d) M256d {
	return M256d(andPd([4]float64(a), [4]float64(b)))
}

func andPd(a [4]float64, b [4]float64) [4]float64


// AndPs: Compute the bitwise AND of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] AND b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDPS'. Intrinsic: '_mm256_and_ps'.
// Requires AVX.
func AndPs(a M256, b M256) M256 {
	return M256(andPs([8]float32(a), [8]float32(b)))
}

func andPs(a [8]float32, b [8]float32) [8]float32


// AndnotPd: Compute the bitwise AND NOT of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm256_andnot_pd'.
// Requires AVX.
func AndnotPd(a M256d, b M256d) M256d {
	return M256d(andnotPd([4]float64(a), [4]float64(b)))
}

func andnotPd(a [4]float64, b [4]float64) [4]float64


// AndnotPs: Compute the bitwise AND NOT of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm256_andnot_ps'.
// Requires AVX.
func AndnotPs(a M256, b M256) M256 {
	return M256(andnotPs([8]float32(a), [8]float32(b)))
}

func andnotPs(a [8]float32, b [8]float32) [8]float32


// AsinPd: Compute the inverse sine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ASIN(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_asin_pd'.
// Requires AVX.
func AsinPd(a M256d) M256d {
	return M256d(asinPd([4]float64(a)))
}

func asinPd(a [4]float64) [4]float64


// AsinPs: Compute the inverse sine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ASIN(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_asin_ps'.
// Requires AVX.
func AsinPs(a M256) M256 {
	return M256(asinPs([8]float32(a)))
}

func asinPs(a [8]float32) [8]float32


// AsinhPd: Compute the inverse hyperbolic sine of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ASINH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_asinh_pd'.
// Requires AVX.
func AsinhPd(a M256d) M256d {
	return M256d(asinhPd([4]float64(a)))
}

func asinhPd(a [4]float64) [4]float64


// AsinhPs: Compute the inverse hyperbolic sine of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ASINH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_asinh_ps'.
// Requires AVX.
func AsinhPs(a M256) M256 {
	return M256(asinhPs([8]float32(a)))
}

func asinhPs(a [8]float32) [8]float32


// AtanPd: Compute the inverse tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ATAN(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atan_pd'.
// Requires AVX.
func AtanPd(a M256d) M256d {
	return M256d(atanPd([4]float64(a)))
}

func atanPd(a [4]float64) [4]float64


// AtanPs: Compute the inverse tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ATAN(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atan_ps'.
// Requires AVX.
func AtanPs(a M256) M256 {
	return M256(atanPs([8]float32(a)))
}

func atanPs(a [8]float32) [8]float32


// Atan2Pd: Compute the inverse tangent of packed double-precision (64-bit)
// floating-point elements in 'a' divided by packed elements in 'b', and store
// the results in 'dst' expressed in radians. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ATAN(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atan2_pd'.
// Requires AVX.
func Atan2Pd(a M256d, b M256d) M256d {
	return M256d(atan2Pd([4]float64(a), [4]float64(b)))
}

func atan2Pd(a [4]float64, b [4]float64) [4]float64


// Atan2Ps: Compute the inverse tangent of packed single-precision (32-bit)
// floating-point elements in 'a' divided by packed elements in 'b', and store
// the results in 'dst' expressed in radians. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ATAN(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atan2_ps'.
// Requires AVX.
func Atan2Ps(a M256, b M256) M256 {
	return M256(atan2Ps([8]float32(a), [8]float32(b)))
}

func atan2Ps(a [8]float32, b [8]float32) [8]float32


// AtanhPd: Compute the inverse hyperbolic tangent of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ATANH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atanh_pd'.
// Requires AVX.
func AtanhPd(a M256d) M256d {
	return M256d(atanhPd([4]float64(a)))
}

func atanhPd(a [4]float64) [4]float64


// AtanhPs: Compute the inverse hyperbolic tangent of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ATANH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atanh_ps'.
// Requires AVX.
func AtanhPs(a M256) M256 {
	return M256(atanhPs([8]float32(a)))
}

func atanhPs(a [8]float32) [8]float32


// BlendPd: Blend packed double-precision (64-bit) floating-point elements from
// 'a' and 'b' using control mask 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF imm8[j%8]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBLENDPD'. Intrinsic: '_mm256_blend_pd'.
// Requires AVX.
func BlendPd(a M256d, b M256d, imm8 int) M256d {
	return M256d(blendPd([4]float64(a), [4]float64(b), imm8))
}

func blendPd(a [4]float64, b [4]float64, imm8 int) [4]float64


// BlendPs: Blend packed single-precision (32-bit) floating-point elements from
// 'a' and 'b' using control mask 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF imm8[j%8]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBLENDPS'. Intrinsic: '_mm256_blend_ps'.
// Requires AVX.
func BlendPs(a M256, b M256, imm8 int) M256 {
	return M256(blendPs([8]float32(a), [8]float32(b), imm8))
}

func blendPs(a [8]float32, b [8]float32, imm8 int) [8]float32


// BlendvPd: Blend packed double-precision (64-bit) floating-point elements
// from 'a' and 'b' using 'mask', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBLENDVPD'. Intrinsic: '_mm256_blendv_pd'.
// Requires AVX.
func BlendvPd(a M256d, b M256d, mask M256d) M256d {
	return M256d(blendvPd([4]float64(a), [4]float64(b), [4]float64(mask)))
}

func blendvPd(a [4]float64, b [4]float64, mask [4]float64) [4]float64


// BlendvPs: Blend packed single-precision (32-bit) floating-point elements
// from 'a' and 'b' using 'mask', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBLENDVPS'. Intrinsic: '_mm256_blendv_ps'.
// Requires AVX.
func BlendvPs(a M256, b M256, mask M256) M256 {
	return M256(blendvPs([8]float32(a), [8]float32(b), [8]float32(mask)))
}

func blendvPs(a [8]float32, b [8]float32, mask [8]float32) [8]float32


// BroadcastPd: Broadcast 128 bits from memory (composed of 2 packed
// double-precision (64-bit) floating-point elements) to all elements of 'dst'. 
//
//		tmp[127:0] = MEM[mem_addr+127:mem_addr]
//		dst[127:0] := tmp[127:0]
//		dst[255:128] := tmp[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF128'. Intrinsic: '_mm256_broadcast_pd'.
// Requires AVX.
func BroadcastPd(mem_addr M128dConst) M256d {
	return M256d(broadcastPd(mem_addr))
}

func broadcastPd(mem_addr M128dConst) [4]float64


// BroadcastPs: Broadcast 128 bits from memory (composed of 4 packed
// single-precision (32-bit) floating-point elements) to all elements of 'dst'. 
//
//		tmp[127:0] = MEM[mem_addr+127:mem_addr]
//		dst[127:0] := tmp[127:0]
//		dst[255:128] := tmp[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF128'. Intrinsic: '_mm256_broadcast_ps'.
// Requires AVX.
func BroadcastPs(mem_addr M128Const) M256 {
	return M256(broadcastPs(mem_addr))
}

func broadcastPs(mem_addr M128Const) [8]float32


// BroadcastSd: Broadcast a double-precision (64-bit) floating-point element
// from memory to all elements of 'dst'. 
//
//		tmp[63:0] = MEM[mem_addr+63:mem_addr]
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := tmp[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTSD'. Intrinsic: '_mm256_broadcast_sd'.
// Requires AVX.
func BroadcastSd(mem_addr float64) M256d {
	return M256d(broadcastSd(mem_addr))
}

func broadcastSd(mem_addr float64) [4]float64


// BroadcastSs: Broadcast a single-precision (32-bit) floating-point element
// from memory to all elements of 'dst'. 
//
//		tmp[31:0] = MEM[mem_addr+31:mem_addr]
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := tmp[31:0]
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VBROADCASTSS'. Intrinsic: '_mm_broadcast_ss'.
// Requires AVX.
func BroadcastSs(mem_addr float32) M128 {
	return M128(broadcastSs(mem_addr))
}

func broadcastSs(mem_addr float32) [4]float32


// BroadcastSs1: Broadcast a single-precision (32-bit) floating-point element
// from memory to all elements of 'dst'. 
//
//		tmp[31:0] = MEM[mem_addr+31:mem_addr]
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := tmp[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTSS'. Intrinsic: '_mm256_broadcast_ss'.
// Requires AVX.
func BroadcastSs1(mem_addr float32) M256 {
	return M256(broadcastSs1(mem_addr))
}

func broadcastSs1(mem_addr float32) [8]float32


// CastpdPs: Cast vector of type __m256d to type __m256.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd_ps'.
// Requires AVX.
func CastpdPs(a M256d) M256 {
	return M256(castpdPs([4]float64(a)))
}

func castpdPs(a [4]float64) [8]float32


// CastpdSi256: Casts vector of type __m256d to type __m256i. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd_si256'.
// Requires AVX.
func CastpdSi256(a M256d) M256i {
	return M256i(castpdSi256([4]float64(a)))
}

func castpdSi256(a [4]float64) [32]byte


// Castpd128Pd256: Casts vector of type __m128d to type __m256d; the upper 128
// bits of the result are undefined. This intrinsic is only used for
// compilation and does not generate any instructions, thus it has zero
// latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd128_pd256'.
// Requires AVX.
func Castpd128Pd256(a M128d) M256d {
	return M256d(castpd128Pd256([2]float64(a)))
}

func castpd128Pd256(a [2]float64) [4]float64


// Castpd256Pd128: Casts vector of type __m256d to type __m128d. This intrinsic
// is only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd256_pd128'.
// Requires AVX.
func Castpd256Pd128(a M256d) M128d {
	return M128d(castpd256Pd128([4]float64(a)))
}

func castpd256Pd128(a [4]float64) [2]float64


// CastpsPd: Cast vector of type __m256 to type __m256d.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps_pd'.
// Requires AVX.
func CastpsPd(a M256) M256d {
	return M256d(castpsPd([8]float32(a)))
}

func castpsPd(a [8]float32) [4]float64


// CastpsSi256: Casts vector of type __m256 to type __m256i. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps_si256'.
// Requires AVX.
func CastpsSi256(a M256) M256i {
	return M256i(castpsSi256([8]float32(a)))
}

func castpsSi256(a [8]float32) [32]byte


// Castps128Ps256: Casts vector of type __m128 to type __m256; the upper 128
// bits of the result are undefined. This intrinsic is only used for
// compilation and does not generate any instructions, thus it has zero
// latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps128_ps256'.
// Requires AVX.
func Castps128Ps256(a M128) M256 {
	return M256(castps128Ps256([4]float32(a)))
}

func castps128Ps256(a [4]float32) [8]float32


// Castps256Ps128: Casts vector of type __m256 to type __m128. This intrinsic
// is only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps256_ps128'.
// Requires AVX.
func Castps256Ps128(a M256) M128 {
	return M128(castps256Ps128([8]float32(a)))
}

func castps256Ps128(a [8]float32) [4]float32


// Castsi128Si256: Casts vector of type __m128i to type __m256i; the upper 128
// bits of the result are undefined. This intrinsic is only used for
// compilation and does not generate any instructions, thus it has zero
// latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi128_si256'.
// Requires AVX.
func Castsi128Si256(a M128i) M256i {
	return M256i(castsi128Si256([16]byte(a)))
}

func castsi128Si256(a [16]byte) [32]byte


// Castsi256Pd: Casts vector of type __m256i to type __m256d. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi256_pd'.
// Requires AVX.
func Castsi256Pd(a M256i) M256d {
	return M256d(castsi256Pd([32]byte(a)))
}

func castsi256Pd(a [32]byte) [4]float64


// Castsi256Ps: Casts vector of type __m256i to type __m256. This intrinsic is
// only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi256_ps'.
// Requires AVX.
func Castsi256Ps(a M256i) M256 {
	return M256(castsi256Ps([32]byte(a)))
}

func castsi256Ps(a [32]byte) [8]float32


// Castsi256Si128: Casts vector of type __m256i to type __m128i. This intrinsic
// is only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi256_si128'.
// Requires AVX.
func Castsi256Si128(a M256i) M128i {
	return M128i(castsi256Si128([32]byte(a)))
}

func castsi256Si128(a [32]byte) [16]byte


// CbrtPd: Compute the cube root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := CubeRoot(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cbrt_pd'.
// Requires AVX.
func CbrtPd(a M256d) M256d {
	return M256d(cbrtPd([4]float64(a)))
}

func cbrtPd(a [4]float64) [4]float64


// CbrtPs: Compute the cube root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := CubeRoot(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cbrt_ps'.
// Requires AVX.
func CbrtPs(a M256) M256 {
	return M256(cbrtPs([8]float32(a)))
}

func cbrtPs(a [8]float32) [8]float32


// CdfnormPd: Compute the cumulative distribution function of packed
// double-precision (64-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := CDFNormal(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cdfnorm_pd'.
// Requires AVX.
func CdfnormPd(a M256d) M256d {
	return M256d(cdfnormPd([4]float64(a)))
}

func cdfnormPd(a [4]float64) [4]float64


// CdfnormPs: Compute the cumulative distribution function of packed
// single-precision (32-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := CDFNormal(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cdfnorm_ps'.
// Requires AVX.
func CdfnormPs(a M256) M256 {
	return M256(cdfnormPs([8]float32(a)))
}

func cdfnormPs(a [8]float32) [8]float32


// CdfnorminvPd: Compute the inverse cumulative distribution function of packed
// double-precision (64-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := InverseCDFNormal(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cdfnorminv_pd'.
// Requires AVX.
func CdfnorminvPd(a M256d) M256d {
	return M256d(cdfnorminvPd([4]float64(a)))
}

func cdfnorminvPd(a [4]float64) [4]float64


// CdfnorminvPs: Compute the inverse cumulative distribution function of packed
// single-precision (32-bit) floating-point elements in 'a' using the normal
// distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := InverseCDFNormal(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cdfnorminv_ps'.
// Requires AVX.
func CdfnorminvPs(a M256) M256 {
	return M256(cdfnorminvPs([8]float32(a)))
}

func cdfnorminvPs(a [8]float32) [8]float32


// CeilPd: Round the packed double-precision (64-bit) floating-point elements
// in 'a' up to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := CEIL(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VROUNDPD'. Intrinsic: '_mm256_ceil_pd'.
// Requires AVX.
func CeilPd(a M256d) M256d {
	return M256d(ceilPd([4]float64(a)))
}

func ceilPd(a [4]float64) [4]float64


// CeilPs: Round the packed single-precision (32-bit) floating-point elements
// in 'a' up to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := CEIL(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm256_ceil_ps'.
// Requires AVX.
func CeilPs(a M256) M256 {
	return M256(ceilPs([8]float32(a)))
}

func ceilPs(a [8]float32) [8]float32


// CexpPs: Compute the exponential value of 'e' raised to the power of packed
// complex single-precision (32-bit) floating-point elements in 'a', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cexp_ps'.
// Requires AVX.
func CexpPs(a M256) M256 {
	return M256(cexpPs([8]float32(a)))
}

func cexpPs(a [8]float32) [8]float32


// ClogPs: Compute the natural logarithm of packed complex single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ln(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_clog_ps'.
// Requires AVX.
func ClogPs(a M256) M256 {
	return M256(clogPs([8]float32(a)))
}

func clogPs(a [8]float32) [8]float32


// CmpPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' based on the comparison operand specified by 'imm8', and store
// the results in 'dst'. 
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
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ( a[i+63:i] OP b[i+63:i] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm_cmp_pd'.
// Requires AVX.
func CmpPd(a M128d, b M128d, imm8 int) M128d {
	return M128d(cmpPd([2]float64(a), [2]float64(b), imm8))
}

func cmpPd(a [2]float64, b [2]float64, imm8 int) [2]float64


// CmpPd1: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b' based on the comparison operand specified by 'imm8', and store
// the results in 'dst'. 
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
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ( a[i+63:i] OP b[i+63:i] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm256_cmp_pd'.
// Requires AVX.
func CmpPd1(a M256d, b M256d, imm8 int) M256d {
	return M256d(cmpPd1([4]float64(a), [4]float64(b), imm8))
}

func cmpPd1(a [4]float64, b [4]float64, imm8 int) [4]float64


// CmpPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' based on the comparison operand specified by 'imm8', and store
// the results in 'dst'. 
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
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] OP b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm_cmp_ps'.
// Requires AVX.
func CmpPs(a M128, b M128, imm8 int) M128 {
	return M128(cmpPs([4]float32(a), [4]float32(b), imm8))
}

func cmpPs(a [4]float32, b [4]float32, imm8 int) [4]float32


// CmpPs1: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b' based on the comparison operand specified by 'imm8', and store
// the results in 'dst'. 
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
//			i := j*32
//			dst[i+31:i] := ( a[i+31:i] OP b[i+31:i] ) ? 0xFFFFFFFF : 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm256_cmp_ps'.
// Requires AVX.
func CmpPs1(a M256, b M256, imm8 int) M256 {
	return M256(cmpPs1([8]float32(a), [8]float32(b), imm8))
}

func cmpPs1(a [8]float32, b [8]float32, imm8 int) [8]float32


// CmpSd: Compare the lower double-precision (64-bit) floating-point element in
// 'a' and 'b' based on the comparison operand specified by 'imm8', store the
// result in the lower element of 'dst', and copy the upper element from 'a' to
// the upper element of 'dst'. 
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
//		
//		dst[63:0] := ( a[63:0] OP b[63:0] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VCMPSD'. Intrinsic: '_mm_cmp_sd'.
// Requires AVX.
func CmpSd(a M128d, b M128d, imm8 int) M128d {
	return M128d(cmpSd([2]float64(a), [2]float64(b), imm8))
}

func cmpSd(a [2]float64, b [2]float64, imm8 int) [2]float64


// CmpSs: Compare the lower single-precision (32-bit) floating-point element in
// 'a' and 'b' based on the comparison operand specified by 'imm8', store the
// result in the lower element of 'dst', and copy the upper 3 packed elements
// from 'a' to the upper elements of 'dst'. 
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
//		
//		dst[31:0] := ( a[31:0] OP b[31:0] ) ? 0xFFFFFFFF : 0
//		dst[127:32] := a[127:32]
//		dst[MAX:128] := 0
//
// Instruction: 'VCMPSS'. Intrinsic: '_mm_cmp_ss'.
// Requires AVX.
func CmpSs(a M128, b M128, imm8 int) M128 {
	return M128(cmpSs([4]float32(a), [4]float32(b), imm8))
}

func cmpSs(a [4]float32, b [4]float32, imm8 int) [4]float32


// CosPd: Compute the cosine of packed double-precision (64-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := COS(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cos_pd'.
// Requires AVX.
func CosPd(a M256d) M256d {
	return M256d(cosPd([4]float64(a)))
}

func cosPd(a [4]float64) [4]float64


// CosPs: Compute the cosine of packed single-precision (32-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := COS(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cos_ps'.
// Requires AVX.
func CosPs(a M256) M256 {
	return M256(cosPs([8]float32(a)))
}

func cosPs(a [8]float32) [8]float32


// CosdPd: Compute the cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := COSD(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cosd_pd'.
// Requires AVX.
func CosdPd(a M256d) M256d {
	return M256d(cosdPd([4]float64(a)))
}

func cosdPd(a [4]float64) [4]float64


// CosdPs: Compute the cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := COSD(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cosd_ps'.
// Requires AVX.
func CosdPs(a M256) M256 {
	return M256(cosdPs([8]float32(a)))
}

func cosdPs(a [8]float32) [8]float32


// CoshPd: Compute the hyperbolic cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := COSH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cosh_pd'.
// Requires AVX.
func CoshPd(a M256d) M256d {
	return M256d(coshPd([4]float64(a)))
}

func coshPd(a [4]float64) [4]float64


// CoshPs: Compute the hyperbolic cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := COSH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cosh_ps'.
// Requires AVX.
func CoshPs(a M256) M256 {
	return M256(coshPs([8]float32(a)))
}

func coshPs(a [8]float32) [8]float32


// CsqrtPs: Compute the square root of packed complex single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_csqrt_ps'.
// Requires AVX.
func CsqrtPs(a M256) M256 {
	return M256(csqrtPs([8]float32(a)))
}

func csqrtPs(a [8]float32) [8]float32


// Cvtepi32Pd: Convert packed 32-bit integers in 'a' to packed double-precision
// (64-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			m := j*64
//			dst[m+63:m] := Convert_Int32_To_FP64(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTDQ2PD'. Intrinsic: '_mm256_cvtepi32_pd'.
// Requires AVX.
func Cvtepi32Pd(a M128i) M256d {
	return M256d(cvtepi32Pd([16]byte(a)))
}

func cvtepi32Pd(a [16]byte) [4]float64


// Cvtepi32Ps: Convert packed 32-bit integers in 'a' to packed single-precision
// (32-bit) floating-point elements, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := Convert_Int32_To_FP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTDQ2PS'. Intrinsic: '_mm256_cvtepi32_ps'.
// Requires AVX.
func Cvtepi32Ps(a M256i) M256 {
	return M256(cvtepi32Ps([32]byte(a)))
}

func cvtepi32Ps(a [32]byte) [8]float32


// CvtpdEpi32: Convert packed double-precision (64-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32(a[k+63:k])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2DQ'. Intrinsic: '_mm256_cvtpd_epi32'.
// Requires AVX.
func CvtpdEpi32(a M256d) M128i {
	return M128i(cvtpdEpi32([4]float64(a)))
}

func cvtpdEpi32(a [4]float64) [16]byte


// CvtpdPs: Convert packed double-precision (64-bit) floating-point elements in
// 'a' to packed single-precision (32-bit) floating-point elements, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_FP32(a[k+63:k])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm256_cvtpd_ps'.
// Requires AVX.
func CvtpdPs(a M256d) M128 {
	return M128(cvtpdPs([4]float64(a)))
}

func cvtpdPs(a [4]float64) [4]float32


// CvtpsEpi32: Convert packed single-precision (32-bit) floating-point elements
// in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2DQ'. Intrinsic: '_mm256_cvtps_epi32'.
// Requires AVX.
func CvtpsEpi32(a M256) M256i {
	return M256i(cvtpsEpi32([8]float32(a)))
}

func cvtpsEpi32(a [8]float32) [32]byte


// CvtpsPd: Convert packed single-precision (32-bit) floating-point elements in
// 'a' to packed double-precision (64-bit) floating-point elements, and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			k := 32*j
//			dst[i+63:i] := Convert_FP32_To_FP64(a[k+31:k])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2PD'. Intrinsic: '_mm256_cvtps_pd'.
// Requires AVX.
func CvtpsPd(a M128) M256d {
	return M256d(cvtpsPd([4]float32(a)))
}

func cvtpsPd(a [4]float32) [4]float64


// CvttpdEpi32: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 32-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 64*j
//			dst[i+31:i] := Convert_FP64_To_Int32_Truncate(a[k+63:k])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VCVTTPD2DQ'. Intrinsic: '_mm256_cvttpd_epi32'.
// Requires AVX.
func CvttpdEpi32(a M256d) M128i {
	return M128i(cvttpdEpi32([4]float64(a)))
}

func cvttpdEpi32(a [4]float64) [16]byte


// CvttpsEpi32: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 32-bit integers with truncation, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32_Truncate(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTTPS2DQ'. Intrinsic: '_mm256_cvttps_epi32'.
// Requires AVX.
func CvttpsEpi32(a M256) M256i {
	return M256i(cvttpsEpi32([8]float32(a)))
}

func cvttpsEpi32(a [8]float32) [32]byte


// DivEpi16: Divide packed 16-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := TRUNCATE(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi16'.
// Requires AVX.
func DivEpi16(a M256i, b M256i) M256i {
	return M256i(divEpi16([32]byte(a), [32]byte(b)))
}

func divEpi16(a [32]byte, b [32]byte) [32]byte


// DivEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi32'.
// Requires AVX.
func DivEpi32(a M256i, b M256i) M256i {
	return M256i(divEpi32([32]byte(a), [32]byte(b)))
}

func divEpi32(a [32]byte, b [32]byte) [32]byte


// DivEpi64: Divide packed 64-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := TRUNCATE(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi64'.
// Requires AVX.
func DivEpi64(a M256i, b M256i) M256i {
	return M256i(divEpi64([32]byte(a), [32]byte(b)))
}

func divEpi64(a [32]byte, b [32]byte) [32]byte


// DivEpi8: Divide packed 8-bit integers in 'a' by packed elements in 'b', and
// store the truncated results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := TRUNCATE(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi8'.
// Requires AVX.
func DivEpi8(a M256i, b M256i) M256i {
	return M256i(divEpi8([32]byte(a), [32]byte(b)))
}

func divEpi8(a [32]byte, b [32]byte) [32]byte


// DivEpu16: Divide packed unsigned 16-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := TRUNCATE(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu16'.
// Requires AVX.
func DivEpu16(a M256i, b M256i) M256i {
	return M256i(divEpu16([32]byte(a), [32]byte(b)))
}

func divEpu16(a [32]byte, b [32]byte) [32]byte


// DivEpu32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu32'.
// Requires AVX.
func DivEpu32(a M256i, b M256i) M256i {
	return M256i(divEpu32([32]byte(a), [32]byte(b)))
}

func divEpu32(a [32]byte, b [32]byte) [32]byte


// DivEpu64: Divide packed unsigned 64-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := TRUNCATE(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu64'.
// Requires AVX.
func DivEpu64(a M256i, b M256i) M256i {
	return M256i(divEpu64([32]byte(a), [32]byte(b)))
}

func divEpu64(a [32]byte, b [32]byte) [32]byte


// DivEpu8: Divide packed unsigned 8-bit integers in 'a' by packed elements in
// 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := TRUNCATE(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu8'.
// Requires AVX.
func DivEpu8(a M256i, b M256i) M256i {
	return M256i(divEpu8([32]byte(a), [32]byte(b)))
}

func divEpu8(a [32]byte, b [32]byte) [32]byte


// DivPd: Divide packed double-precision (64-bit) floating-point elements in
// 'a' by packed elements in 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := a[i+63:i] / b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDIVPD'. Intrinsic: '_mm256_div_pd'.
// Requires AVX.
func DivPd(a M256d, b M256d) M256d {
	return M256d(divPd([4]float64(a), [4]float64(b)))
}

func divPd(a [4]float64, b [4]float64) [4]float64


// DivPs: Divide packed single-precision (32-bit) floating-point elements in
// 'a' by packed elements in 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := a[i+31:i] / b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDIVPS'. Intrinsic: '_mm256_div_ps'.
// Requires AVX.
func DivPs(a M256, b M256) M256 {
	return M256(divPs([8]float32(a), [8]float32(b)))
}

func divPs(a [8]float32, b [8]float32) [8]float32


// DpPs: Conditionally multiply the packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' using the high 4 bits in 'imm8', sum
// the four products, and conditionally store the sum in 'dst' using the low 4
// bits of 'imm8'. 
//
//		DP(a[127:0], b[127:0], imm8[7:0]) {
//			FOR j := 0 to 3
//				i := j*32
//				IF imm8[(4+j)%8]
//					temp[i+31:i] := a[i+31:i] * b[i+31:i]
//				ELSE
//					temp[i+31:i] := 0
//				FI
//			ENDFOR
//			
//			sum[31:0] := (temp[127:96] + temp[95:64]) + (temp[63:32] + temp[31:0])
//			
//			FOR j := 0 to 3
//				i := j*32
//				IF imm8[j%8]
//					tmpdst[i+31:i] := sum[31:0]
//				ELSE
//					tmpdst[i+31:i] := 0
//				FI
//			ENDFOR
//			RETURN tmpdst[127:0]
//		}
//		
//		dst[127:0] := DP(a[127:0], b[127:0], imm8[7:0])
//		dst[255:128] := DP(a[255:128], b[255:128], imm8[7:0])
//		dst[MAX:256] := 0
//
// Instruction: 'VDPPS'. Intrinsic: '_mm256_dp_ps'.
// Requires AVX.
func DpPs(a M256, b M256, imm8 int) M256 {
	return M256(dpPs([8]float32(a), [8]float32(b), imm8))
}

func dpPs(a [8]float32, b [8]float32, imm8 int) [8]float32


// ErfPd: Compute the error function of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erf_pd'.
// Requires AVX.
func ErfPd(a M256d) M256d {
	return M256d(erfPd([4]float64(a)))
}

func erfPd(a [4]float64) [4]float64


// ErfPs: Compute the error function of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erf_ps'.
// Requires AVX.
func ErfPs(a M256) M256 {
	return M256(erfPs([8]float32(a)))
}

func erfPs(a [8]float32) [8]float32


// ErfcPd: Compute the complementary error function of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 1.0 - ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfc_pd'.
// Requires AVX.
func ErfcPd(a M256d) M256d {
	return M256d(erfcPd([4]float64(a)))
}

func erfcPd(a [4]float64) [4]float64


// ErfcPs: Compute the complementary error function of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 1.0 - ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfc_ps'.
// Requires AVX.
func ErfcPs(a M256) M256 {
	return M256(erfcPs([8]float32(a)))
}

func erfcPs(a [8]float32) [8]float32


// ErfcinvPd: Compute the inverse complementary error function of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 1.0 / (1.0 - ERF(a[i+63:i]))
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfcinv_pd'.
// Requires AVX.
func ErfcinvPd(a M256d) M256d {
	return M256d(erfcinvPd([4]float64(a)))
}

func erfcinvPd(a [4]float64) [4]float64


// ErfcinvPs: Compute the inverse complementary error function of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 1.0 / (1.0 - ERF(a[i+31:i]))
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfcinv_ps'.
// Requires AVX.
func ErfcinvPs(a M256) M256 {
	return M256(erfcinvPs([8]float32(a)))
}

func erfcinvPs(a [8]float32) [8]float32


// ErfinvPd: Compute the inverse error function of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 1.0 / ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfinv_pd'.
// Requires AVX.
func ErfinvPd(a M256d) M256d {
	return M256d(erfinvPd([4]float64(a)))
}

func erfinvPd(a [4]float64) [4]float64


// ErfinvPs: Compute the inverse error function of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 1.0 / ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfinv_ps'.
// Requires AVX.
func ErfinvPs(a M256) M256 {
	return M256(erfinvPs([8]float32(a)))
}

func erfinvPs(a [8]float32) [8]float32


// ExpPd: Compute the exponential value of 'e' raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := e^(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp_pd'.
// Requires AVX.
func ExpPd(a M256d) M256d {
	return M256d(expPd([4]float64(a)))
}

func expPd(a [4]float64) [4]float64


// ExpPs: Compute the exponential value of 'e' raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp_ps'.
// Requires AVX.
func ExpPs(a M256) M256 {
	return M256(expPs([8]float32(a)))
}

func expPs(a [8]float32) [8]float32


// Exp10Pd: Compute the exponential value of 10 raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 10^(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp10_pd'.
// Requires AVX.
func Exp10Pd(a M256d) M256d {
	return M256d(exp10Pd([4]float64(a)))
}

func exp10Pd(a [4]float64) [4]float64


// Exp10Ps: Compute the exponential value of 10 raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 10^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp10_ps'.
// Requires AVX.
func Exp10Ps(a M256) M256 {
	return M256(exp10Ps([8]float32(a)))
}

func exp10Ps(a [8]float32) [8]float32


// Exp2Pd: Compute the exponential value of 2 raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 2^(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp2_pd'.
// Requires AVX.
func Exp2Pd(a M256d) M256d {
	return M256d(exp2Pd([4]float64(a)))
}

func exp2Pd(a [4]float64) [4]float64


// Exp2Ps: Compute the exponential value of 2 raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 2^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp2_ps'.
// Requires AVX.
func Exp2Ps(a M256) M256 {
	return M256(exp2Ps([8]float32(a)))
}

func exp2Ps(a [8]float32) [8]float32


// Expm1Pd: Compute the exponential value of 'e' raised to the power of packed
// double-precision (64-bit) floating-point elements in 'a', subtract one from
// each element, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := e^(a[i+63:i]) - 1.0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_expm1_pd'.
// Requires AVX.
func Expm1Pd(a M256d) M256d {
	return M256d(expm1Pd([4]float64(a)))
}

func expm1Pd(a [4]float64) [4]float64


// Expm1Ps: Compute the exponential value of 'e' raised to the power of packed
// single-precision (32-bit) floating-point elements in 'a', subtract one from
// each element, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i]) - 1.0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_expm1_ps'.
// Requires AVX.
func Expm1Ps(a M256) M256 {
	return M256(expm1Ps([8]float32(a)))
}

func expm1Ps(a [8]float32) [8]float32


// ExtractEpi16: Extract a 16-bit integer from 'a', selected with 'index', and
// store the result in 'dst'. 
//
//		dst[15:0] := (a[255:0] >> (index * 16))[15:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi16'.
// Requires AVX.
func ExtractEpi16(a M256i, index int) int16 {
	return int16(extractEpi16([32]byte(a), index))
}

func extractEpi16(a [32]byte, index int) int16


// ExtractEpi32: Extract a 32-bit integer from 'a', selected with 'index', and
// store the result in 'dst'. 
//
//		dst[31:0] := (a[255:0] >> (index * 32))[31:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi32'.
// Requires AVX.
func ExtractEpi32(a M256i, index int) int32 {
	return int32(extractEpi32([32]byte(a), index))
}

func extractEpi32(a [32]byte, index int) int32


// ExtractEpi64: Extract a 64-bit integer from 'a', selected with 'index', and
// store the result in 'dst'. 
//
//		dst[63:0] := (a[255:0] >> (index * 64))[63:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi64'.
// Requires AVX.
func ExtractEpi64(a M256i, index int) int64 {
	return int64(extractEpi64([32]byte(a), index))
}

func extractEpi64(a [32]byte, index int) int64


// ExtractEpi8: Extract an 8-bit integer from 'a', selected with 'index', and
// store the result in 'dst'. 
//
//		dst[7:0] := (a[255:0] >> (index * 8))[7:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi8'.
// Requires AVX.
func ExtractEpi8(a M256i, index int) int8 {
	return int8(extractEpi8([32]byte(a), index))
}

func extractEpi8(a [32]byte, index int) int8


// Extractf128Pd: Extract 128 bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF128'. Intrinsic: '_mm256_extractf128_pd'.
// Requires AVX.
func Extractf128Pd(a M256d, imm8 int) M128d {
	return M128d(extractf128Pd([4]float64(a), imm8))
}

func extractf128Pd(a [4]float64, imm8 int) [2]float64


// Extractf128Ps: Extract 128 bits (composed of 4 packed single-precision
// (32-bit) floating-point elements) from 'a', selected with 'imm8', and store
// the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF128'. Intrinsic: '_mm256_extractf128_ps'.
// Requires AVX.
func Extractf128Ps(a M256, imm8 int) M128 {
	return M128(extractf128Ps([8]float32(a), imm8))
}

func extractf128Ps(a [8]float32, imm8 int) [4]float32


// Extractf128Si256: Extract 128 bits (composed of integer data) from 'a',
// selected with 'imm8', and store the result in 'dst'. 
//
//		CASE imm8[7:0] of
//		0: dst[127:0] := a[127:0]
//		1: dst[127:0] := a[255:128]
//		ESAC
//		dst[MAX:128] := 0
//
// Instruction: 'VEXTRACTF128'. Intrinsic: '_mm256_extractf128_si256'.
// Requires AVX.
func Extractf128Si256(a M256i, imm8 int) M128i {
	return M128i(extractf128Si256([32]byte(a), imm8))
}

func extractf128Si256(a [32]byte, imm8 int) [16]byte


// FloorPd: Round the packed double-precision (64-bit) floating-point elements
// in 'a' down to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := FLOOR(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VROUNDPD'. Intrinsic: '_mm256_floor_pd'.
// Requires AVX.
func FloorPd(a M256d) M256d {
	return M256d(floorPd([4]float64(a)))
}

func floorPd(a [4]float64) [4]float64


// FloorPs: Round the packed single-precision (32-bit) floating-point elements
// in 'a' down to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := FLOOR(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm256_floor_ps'.
// Requires AVX.
func FloorPs(a M256) M256 {
	return M256(floorPs([8]float32(a)))
}

func floorPs(a [8]float32) [8]float32


// HaddPd: Horizontally add adjacent pairs of double-precision (64-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[63:0] := a[127:64] + a[63:0]
//		dst[127:64] := b[127:64] + b[63:0]
//		dst[191:128] := a[255:192] + a[191:128]
//		dst[255:192] := b[255:192] + b[191:128]
//		dst[MAX:256] := 0
//
// Instruction: 'VHADDPD'. Intrinsic: '_mm256_hadd_pd'.
// Requires AVX.
func HaddPd(a M256d, b M256d) M256d {
	return M256d(haddPd([4]float64(a), [4]float64(b)))
}

func haddPd(a [4]float64, b [4]float64) [4]float64


// HaddPs: Horizontally add adjacent pairs of single-precision (32-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := a[127:96] + a[95:64]
//		dst[95:64] := b[63:32] + b[31:0]
//		dst[127:96] := b[127:96] + b[95:64]
//		dst[159:128] := a[191:160] + a[159:128]
//		dst[191:160] := a[255:224] + a[223:192]
//		dst[223:192] := b[191:160] + b[159:128]
//		dst[255:224] := b[255:224] + b[223:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VHADDPS'. Intrinsic: '_mm256_hadd_ps'.
// Requires AVX.
func HaddPs(a M256, b M256) M256 {
	return M256(haddPs([8]float32(a), [8]float32(b)))
}

func haddPs(a [8]float32, b [8]float32) [8]float32


// HsubPd: Horizontally subtract adjacent pairs of double-precision (64-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[63:0] := a[63:0] - a[127:64]
//		dst[127:64] := b[63:0] - b[127:64]
//		dst[191:128] := a[191:128] - a[255:192]
//		dst[255:192] := b[191:128] - b[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VHSUBPD'. Intrinsic: '_mm256_hsub_pd'.
// Requires AVX.
func HsubPd(a M256d, b M256d) M256d {
	return M256d(hsubPd([4]float64(a), [4]float64(b)))
}

func hsubPd(a [4]float64, b [4]float64) [4]float64


// HsubPs: Horizontally add adjacent pairs of single-precision (32-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := a[95:64] - a[127:96]
//		dst[95:64] := b[31:0] - b[63:32]
//		dst[127:96] := b[95:64] - b[127:96]
//		dst[159:128] := a[159:128] - a[191:160]
//		dst[191:160] := a[223:192] - a[255:224]
//		dst[223:192] := b[159:128] - b[191:160]
//		dst[255:224] := b[223:192] - b[255:224]
//		dst[MAX:256] := 0
//
// Instruction: 'VHSUBPS'. Intrinsic: '_mm256_hsub_ps'.
// Requires AVX.
func HsubPs(a M256, b M256) M256 {
	return M256(hsubPs([8]float32(a), [8]float32(b)))
}

func hsubPs(a [8]float32, b [8]float32) [8]float32


// HypotPd: Compute the length of the hypotenous of a right triangle, with the
// lengths of the other two sides of the triangle stored as packed
// double-precision (64-bit) floating-point elements in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SQRT(a[i+63:i]^2 + b[i+63:i]^2)
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_hypot_pd'.
// Requires AVX.
func HypotPd(a M256d, b M256d) M256d {
	return M256d(hypotPd([4]float64(a), [4]float64(b)))
}

func hypotPd(a [4]float64, b [4]float64) [4]float64


// HypotPs: Compute the length of the hypotenous of a right triangle, with the
// lengths of the other two sides of the triangle stored as packed
// single-precision (32-bit) floating-point elements in 'a' and 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i]^2 + b[i+31:i]^2)
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_hypot_ps'.
// Requires AVX.
func HypotPs(a M256, b M256) M256 {
	return M256(hypotPs([8]float32(a), [8]float32(b)))
}

func hypotPs(a [8]float32, b [8]float32) [8]float32


// IdivEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_idiv_epi32'.
// Requires AVX.
func IdivEpi32(a M256i, b M256i) M256i {
	return M256i(idivEpi32([32]byte(a), [32]byte(b)))
}

func idivEpi32(a [32]byte, b [32]byte) [32]byte


// IdivremEpi32: Divide packed 32-bit integers in 'a' by packed elements in
// 'b', store the truncated results in 'dst', and store the remainders as
// packed 32-bit integers into memory at 'mem_addr'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//			MEM[mem_addr+i+31:mem_addr+i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_idivrem_epi32'.
// Requires AVX.
func IdivremEpi32(mem_addr M256i, a M256i, b M256i) M256i {
	return M256i(idivremEpi32([32]byte(mem_addr), [32]byte(a), [32]byte(b)))
}

func idivremEpi32(mem_addr [32]byte, a [32]byte, b [32]byte) [32]byte


// InsertEpi16: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into 'dst'
// at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi16'.
// Requires AVX.
func InsertEpi16(a M256i, i int16, index int) M256i {
	return M256i(insertEpi16([32]byte(a), i, index))
}

func insertEpi16(a [32]byte, i int16, index int) [32]byte


// InsertEpi32: Copy 'a' to 'dst', and insert the 32-bit integer 'i' into 'dst'
// at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*32
//		dst[sel+31:sel] := i[31:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi32'.
// Requires AVX.
func InsertEpi32(a M256i, i int32, index int) M256i {
	return M256i(insertEpi32([32]byte(a), i, index))
}

func insertEpi32(a [32]byte, i int32, index int) [32]byte


// InsertEpi64: Copy 'a' to 'dst', and insert the 64-bit integer 'i' into 'dst'
// at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*64
//		dst[sel+63:sel] := i[63:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi64'.
// Requires AVX.
func InsertEpi64(a M256i, i int64, index int) M256i {
	return M256i(insertEpi64([32]byte(a), i, index))
}

func insertEpi64(a [32]byte, i int64, index int) [32]byte


// InsertEpi8: Copy 'a' to 'dst', and insert the 8-bit integer 'i' into 'dst'
// at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*8
//		dst[sel+7:sel] := i[7:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi8'.
// Requires AVX.
func InsertEpi8(a M256i, i int8, index int) M256i {
	return M256i(insertEpi8([32]byte(a), i, index))
}

func insertEpi8(a [32]byte, i int8, index int) [32]byte


// Insertf128Pd: Copy 'a' to 'dst', then insert 128 bits (composed of 2 packed
// double-precision (64-bit) floating-point elements) from 'b' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[255:0] := a[255:0]
//		CASE imm8[7:0] of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_insertf128_pd'.
// Requires AVX.
func Insertf128Pd(a M256d, b M128d, imm8 int) M256d {
	return M256d(insertf128Pd([4]float64(a), [2]float64(b), imm8))
}

func insertf128Pd(a [4]float64, b [2]float64, imm8 int) [4]float64


// Insertf128Ps: Copy 'a' to 'dst', then insert 128 bits (composed of 4 packed
// single-precision (32-bit) floating-point elements) from 'b' into 'dst' at
// the location specified by 'imm8'. 
//
//		dst[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_insertf128_ps'.
// Requires AVX.
func Insertf128Ps(a M256, b M128, imm8 int) M256 {
	return M256(insertf128Ps([8]float32(a), [4]float32(b), imm8))
}

func insertf128Ps(a [8]float32, b [4]float32, imm8 int) [8]float32


// Insertf128Si256: Copy 'a' to 'dst', then insert 128 bits from 'b' into 'dst'
// at the location specified by 'imm8'. 
//
//		dst[255:0] := a[255:0]
//		CASE (imm8[1:0]) of
//		0: dst[127:0] := b[127:0]
//		1: dst[255:128] := b[127:0]
//		ESAC
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_insertf128_si256'.
// Requires AVX.
func Insertf128Si256(a M256i, b M128i, imm8 int) M256i {
	return M256i(insertf128Si256([32]byte(a), [16]byte(b), imm8))
}

func insertf128Si256(a [32]byte, b [16]byte, imm8 int) [32]byte


// InvcbrtPd: Compute the inverse cube root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := InvCubeRoot(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_invcbrt_pd'.
// Requires AVX.
func InvcbrtPd(a M256d) M256d {
	return M256d(invcbrtPd([4]float64(a)))
}

func invcbrtPd(a [4]float64) [4]float64


// InvcbrtPs: Compute the inverse cube root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := InvCubeRoot(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_invcbrt_ps'.
// Requires AVX.
func InvcbrtPs(a M256) M256 {
	return M256(invcbrtPs([8]float32(a)))
}

func invcbrtPs(a [8]float32) [8]float32


// InvsqrtPd: Compute the inverse square root of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := InvSQRT(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_invsqrt_pd'.
// Requires AVX.
func InvsqrtPd(a M256d) M256d {
	return M256d(invsqrtPd([4]float64(a)))
}

func invsqrtPd(a [4]float64) [4]float64


// InvsqrtPs: Compute the inverse square root of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := InvSQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_invsqrt_ps'.
// Requires AVX.
func InvsqrtPs(a M256) M256 {
	return M256(invsqrtPs([8]float32(a)))
}

func invsqrtPs(a [8]float32) [8]float32


// IremEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_irem_epi32'.
// Requires AVX.
func IremEpi32(a M256i, b M256i) M256i {
	return M256i(iremEpi32([32]byte(a), [32]byte(b)))
}

func iremEpi32(a [32]byte, b [32]byte) [32]byte


// LddquSi256: Load 256-bits of integer data from unaligned memory into 'dst'.
// This intrinsic may perform better than '_mm256_loadu_si256' when the data
// crosses a cache line boundary. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VLDDQU'. Intrinsic: '_mm256_lddqu_si256'.
// Requires AVX.
func LddquSi256(mem_addr M256iConst) M256i {
	return M256i(lddquSi256(mem_addr))
}

func lddquSi256(mem_addr M256iConst) [32]byte


// LoadPd: Load 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from memory into 'dst'.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm256_load_pd'.
// Requires AVX.
func LoadPd(mem_addr float64) M256d {
	return M256d(loadPd(mem_addr))
}

func loadPd(mem_addr float64) [4]float64


// LoadPs: Load 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from memory into 'dst'.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm256_load_ps'.
// Requires AVX.
func LoadPs(mem_addr float32) M256 {
	return M256(loadPs(mem_addr))
}

func loadPs(mem_addr float32) [8]float32


// LoadSi256: Load 256-bits of integer data from memory into 'dst'.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQA'. Intrinsic: '_mm256_load_si256'.
// Requires AVX.
func LoadSi256(mem_addr M256iConst) M256i {
	return M256i(loadSi256(mem_addr))
}

func loadSi256(mem_addr M256iConst) [32]byte


// LoaduPd: Load 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from memory into 'dst'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVUPD'. Intrinsic: '_mm256_loadu_pd'.
// Requires AVX.
func LoaduPd(mem_addr float64) M256d {
	return M256d(loaduPd(mem_addr))
}

func loaduPd(mem_addr float64) [4]float64


// LoaduPs: Load 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from memory into 'dst'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVUPS'. Intrinsic: '_mm256_loadu_ps'.
// Requires AVX.
func LoaduPs(mem_addr float32) M256 {
	return M256(loaduPs(mem_addr))
}

func loaduPs(mem_addr float32) [8]float32


// LoaduSi256: Load 256-bits of integer data from memory into 'dst'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU'. Intrinsic: '_mm256_loadu_si256'.
// Requires AVX.
func LoaduSi256(mem_addr M256iConst) M256i {
	return M256i(loaduSi256(mem_addr))
}

func loaduSi256(mem_addr M256iConst) [32]byte


// Loadu2M128: Load two 128-bit values (composed of 4 packed single-precision
// (32-bit) floating-point elements) from memory, and combine them into a
// 256-bit value in 'dst'.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		dst[127:0] := MEM[loaddr+127:loaddr]
//		dst[255:128] := MEM[hiaddr+127:hiaddr]
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_loadu2_m128'.
// Requires AVX.
func Loadu2M128(hiaddr float32, loaddr float32) M256 {
	return M256(loadu2M128(hiaddr, loaddr))
}

func loadu2M128(hiaddr float32, loaddr float32) [8]float32


// Loadu2M128d: Load two 128-bit values (composed of 2 packed double-precision
// (64-bit) floating-point elements) from memory, and combine them into a
// 256-bit value in 'dst'.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		dst[127:0] := MEM[loaddr+127:loaddr]
//		dst[255:128] := MEM[hiaddr+127:hiaddr]
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_loadu2_m128d'.
// Requires AVX.
func Loadu2M128d(hiaddr float64, loaddr float64) M256d {
	return M256d(loadu2M128d(hiaddr, loaddr))
}

func loadu2M128d(hiaddr float64, loaddr float64) [4]float64


// Loadu2M128i: Load two 128-bit values (composed of integer data) from memory,
// and combine them into a 256-bit value in 'dst'.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		dst[127:0] := MEM[loaddr+127:loaddr]
//		dst[255:128] := MEM[hiaddr+127:hiaddr]
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_loadu2_m128i'.
// Requires AVX.
func Loadu2M128i(hiaddr M128iConst, loaddr M128iConst) M256i {
	return M256i(loadu2M128i(hiaddr, loaddr))
}

func loadu2M128i(hiaddr M128iConst, loaddr M128iConst) [32]byte


// LogPd: Compute the natural logarithm of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ln(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log_pd'.
// Requires AVX.
func LogPd(a M256d) M256d {
	return M256d(logPd([4]float64(a)))
}

func logPd(a [4]float64) [4]float64


// LogPs: Compute the natural logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ln(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log_ps'.
// Requires AVX.
func LogPs(a M256) M256 {
	return M256(logPs([8]float32(a)))
}

func logPs(a [8]float32) [8]float32


// Log10Pd: Compute the base-10 logarithm of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := log10(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log10_pd'.
// Requires AVX.
func Log10Pd(a M256d) M256d {
	return M256d(log10Pd([4]float64(a)))
}

func log10Pd(a [4]float64) [4]float64


// Log10Ps: Compute the base-10 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := log10(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log10_ps'.
// Requires AVX.
func Log10Ps(a M256) M256 {
	return M256(log10Ps([8]float32(a)))
}

func log10Ps(a [8]float32) [8]float32


// Log1pPd: Compute the natural logarithm of one plus packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ln(1.0 + a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log1p_pd'.
// Requires AVX.
func Log1pPd(a M256d) M256d {
	return M256d(log1pPd([4]float64(a)))
}

func log1pPd(a [4]float64) [4]float64


// Log1pPs: Compute the natural logarithm of one plus packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ln(1.0 + a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log1p_ps'.
// Requires AVX.
func Log1pPs(a M256) M256 {
	return M256(log1pPs([8]float32(a)))
}

func log1pPs(a [8]float32) [8]float32


// Log2Pd: Compute the base-2 logarithm of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := log2(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log2_pd'.
// Requires AVX.
func Log2Pd(a M256d) M256d {
	return M256d(log2Pd([4]float64(a)))
}

func log2Pd(a [4]float64) [4]float64


// Log2Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := log2(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log2_ps'.
// Requires AVX.
func Log2Ps(a M256) M256 {
	return M256(log2Ps([8]float32(a)))
}

func log2Ps(a [8]float32) [8]float32


// LogbPd: Convert the exponent of each packed double-precision (64-bit)
// floating-point element in 'a' to a double-precision floating-point number
// representing the integer exponent, and store the results in 'dst'. This
// intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_logb_pd'.
// Requires AVX.
func LogbPd(a M256d) M256d {
	return M256d(logbPd([4]float64(a)))
}

func logbPd(a [4]float64) [4]float64


// LogbPs: Convert the exponent of each packed single-precision (32-bit)
// floating-point element in 'a' to a single-precision floating-point number
// representing the integer exponent, and store the results in 'dst'. This
// intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_logb_ps'.
// Requires AVX.
func LogbPs(a M256) M256 {
	return M256(logbPs([8]float32(a)))
}

func logbPs(a [8]float32) [8]float32


// MaskloadPd: Load packed double-precision (64-bit) floating-point elements
// from memory into 'dst' using 'mask' (elements are zeroed out when the high
// bit of the corresponding element is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := MEM[mem_addr+i+63:mem_addr+i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMASKMOVPD'. Intrinsic: '_mm_maskload_pd'.
// Requires AVX.
func MaskloadPd(mem_addr float64, mask M128i) M128d {
	return M128d(maskloadPd(mem_addr, [16]byte(mask)))
}

func maskloadPd(mem_addr float64, mask [16]byte) [2]float64


// MaskloadPd1: Load packed double-precision (64-bit) floating-point elements
// from memory into 'dst' using 'mask' (elements are zeroed out when the high
// bit of the corresponding element is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := MEM[mem_addr+i+63:mem_addr+i]
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMASKMOVPD'. Intrinsic: '_mm256_maskload_pd'.
// Requires AVX.
func MaskloadPd1(mem_addr float64, mask M256i) M256d {
	return M256d(maskloadPd1(mem_addr, [32]byte(mask)))
}

func maskloadPd1(mem_addr float64, mask [32]byte) [4]float64


// MaskloadPs: Load packed single-precision (32-bit) floating-point elements
// from memory into 'dst' using 'mask' (elements are zeroed out when the high
// bit of the corresponding element is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := MEM[mem_addr+i+31:mem_addr+i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VMASKMOVPS'. Intrinsic: '_mm_maskload_ps'.
// Requires AVX.
func MaskloadPs(mem_addr float32, mask M128i) M128 {
	return M128(maskloadPs(mem_addr, [16]byte(mask)))
}

func maskloadPs(mem_addr float32, mask [16]byte) [4]float32


// MaskloadPs1: Load packed single-precision (32-bit) floating-point elements
// from memory into 'dst' using 'mask' (elements are zeroed out when the high
// bit of the corresponding element is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := MEM[mem_addr+i+31:mem_addr+i]
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMASKMOVPS'. Intrinsic: '_mm256_maskload_ps'.
// Requires AVX.
func MaskloadPs1(mem_addr float32, mask M256i) M256 {
	return M256(maskloadPs1(mem_addr, [32]byte(mask)))
}

func maskloadPs1(mem_addr float32, mask [32]byte) [8]float32


// MaskstorePd: Store packed double-precision (64-bit) floating-point elements
// from 'a' into memory using 'mask'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF mask[i+63]
//				MEM[mem_addr+i+63:mem_addr+i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMASKMOVPD'. Intrinsic: '_mm_maskstore_pd'.
// Requires AVX.
func MaskstorePd(mem_addr float64, mask M128i, a M128d)  {
	maskstorePd(mem_addr, [16]byte(mask), [2]float64(a))
}

func maskstorePd(mem_addr float64, mask [16]byte, a [2]float64) 


// MaskstorePd1: Store packed double-precision (64-bit) floating-point elements
// from 'a' into memory using 'mask'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF mask[i+63]
//				MEM[mem_addr+i+63:mem_addr+i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMASKMOVPD'. Intrinsic: '_mm256_maskstore_pd'.
// Requires AVX.
func MaskstorePd1(mem_addr float64, mask M256i, a M256d)  {
	maskstorePd1(mem_addr, [32]byte(mask), [4]float64(a))
}

func maskstorePd1(mem_addr float64, mask [32]byte, a [4]float64) 


// MaskstorePs: Store packed single-precision (32-bit) floating-point elements
// from 'a' into memory using 'mask'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF mask[i+31]
//				MEM[mem_addr+i+31:mem_addr+i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMASKMOVPS'. Intrinsic: '_mm_maskstore_ps'.
// Requires AVX.
func MaskstorePs(mem_addr float32, mask M128i, a M128)  {
	maskstorePs(mem_addr, [16]byte(mask), [4]float32(a))
}

func maskstorePs(mem_addr float32, mask [16]byte, a [4]float32) 


// MaskstorePs1: Store packed single-precision (32-bit) floating-point elements
// from 'a' into memory using 'mask'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF mask[i+31]
//				MEM[mem_addr+i+31:mem_addr+i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'VMASKMOVPS'. Intrinsic: '_mm256_maskstore_ps'.
// Requires AVX.
func MaskstorePs1(mem_addr float32, mask M256i, a M256)  {
	maskstorePs1(mem_addr, [32]byte(mask), [8]float32(a))
}

func maskstorePs1(mem_addr float32, mask [32]byte, a [8]float32) 


// MaxPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := MAX(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMAXPD'. Intrinsic: '_mm256_max_pd'.
// Requires AVX.
func MaxPd(a M256d, b M256d) M256d {
	return M256d(maxPd([4]float64(a), [4]float64(b)))
}

func maxPd(a [4]float64, b [4]float64) [4]float64


// MaxPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := MAX(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMAXPS'. Intrinsic: '_mm256_max_ps'.
// Requires AVX.
func MaxPs(a M256, b M256) M256 {
	return M256(maxPs([8]float32(a), [8]float32(b)))
}

func maxPs(a [8]float32, b [8]float32) [8]float32


// MinPd: Compare packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := MIN(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMINPD'. Intrinsic: '_mm256_min_pd'.
// Requires AVX.
func MinPd(a M256d, b M256d) M256d {
	return M256d(minPd([4]float64(a), [4]float64(b)))
}

func minPd(a [4]float64, b [4]float64) [4]float64


// MinPs: Compare packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := MIN(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMINPS'. Intrinsic: '_mm256_min_ps'.
// Requires AVX.
func MinPs(a M256, b M256) M256 {
	return M256(minPs([8]float32(a), [8]float32(b)))
}

func minPs(a [8]float32, b [8]float32) [8]float32


// MovedupPd: Duplicate even-indexed double-precision (64-bit) floating-point
// elements from 'a', and store the results in 'dst'. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := a[63:0]
//		dst[191:128] := a[191:128]
//		dst[255:192] := a[191:128]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDDUP'. Intrinsic: '_mm256_movedup_pd'.
// Requires AVX.
func MovedupPd(a M256d) M256d {
	return M256d(movedupPd([4]float64(a)))
}

func movedupPd(a [4]float64) [4]float64


// MovehdupPs: Duplicate odd-indexed single-precision (32-bit) floating-point
// elements from 'a', and store the results in 'dst'. 
//
//		dst[31:0] := a[63:32] 
//		dst[63:32] := a[63:32] 
//		dst[95:64] := a[127:96] 
//		dst[127:96] := a[127:96]
//		dst[159:128] := a[191:160] 
//		dst[191:160] := a[191:160] 
//		dst[223:192] := a[255:224] 
//		dst[255:224] := a[255:224]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVSHDUP'. Intrinsic: '_mm256_movehdup_ps'.
// Requires AVX.
func MovehdupPs(a M256) M256 {
	return M256(movehdupPs([8]float32(a)))
}

func movehdupPs(a [8]float32) [8]float32


// MoveldupPs: Duplicate even-indexed single-precision (32-bit) floating-point
// elements from 'a', and store the results in 'dst'. 
//
//		dst[31:0] := a[31:0] 
//		dst[63:32] := a[31:0] 
//		dst[95:64] := a[95:64] 
//		dst[127:96] := a[95:64]
//		dst[159:128] := a[159:128] 
//		dst[191:160] := a[159:128] 
//		dst[223:192] := a[223:192] 
//		dst[255:224] := a[223:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVSLDUP'. Intrinsic: '_mm256_moveldup_ps'.
// Requires AVX.
func MoveldupPs(a M256) M256 {
	return M256(moveldupPs([8]float32(a)))
}

func moveldupPs(a [8]float32) [8]float32


// MovemaskPd: Set each bit of mask 'dst' based on the most significant bit of
// the corresponding packed double-precision (64-bit) floating-point element in
// 'a'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF a[i+63]
//				dst[j] := 1
//			ELSE
//				dst[j] := 0
//			FI
//		ENDFOR
//		dst[MAX:4] := 0
//
// Instruction: 'VMOVMSKPD'. Intrinsic: '_mm256_movemask_pd'.
// Requires AVX.
func MovemaskPd(a M256d) int {
	return int(movemaskPd([4]float64(a)))
}

func movemaskPd(a [4]float64) int


// MovemaskPs: Set each bit of mask 'dst' based on the most significant bit of
// the corresponding packed single-precision (32-bit) floating-point element in
// 'a'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF a[i+31]
//				dst[j] := 1
//			ELSE
//				dst[j] := 0
//			FI
//		ENDFOR
//		dst[MAX:8] := 0
//
// Instruction: 'VMOVMSKPS'. Intrinsic: '_mm256_movemask_ps'.
// Requires AVX.
func MovemaskPs(a M256) int {
	return int(movemaskPs([8]float32(a)))
}

func movemaskPs(a [8]float32) int


// MulPd: Multiply packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm256_mul_pd'.
// Requires AVX.
func MulPd(a M256d, b M256d) M256d {
	return M256d(mulPd([4]float64(a), [4]float64(b)))
}

func mulPd(a [4]float64, b [4]float64) [4]float64


// MulPs: Multiply packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm256_mul_ps'.
// Requires AVX.
func MulPs(a M256, b M256) M256 {
	return M256(mulPs([8]float32(a), [8]float32(b)))
}

func mulPs(a [8]float32, b [8]float32) [8]float32


// OrPd: Compute the bitwise OR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] BITWISE OR b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VORPD'. Intrinsic: '_mm256_or_pd'.
// Requires AVX.
func OrPd(a M256d, b M256d) M256d {
	return M256d(orPd([4]float64(a), [4]float64(b)))
}

func orPd(a [4]float64, b [4]float64) [4]float64


// OrPs: Compute the bitwise OR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] BITWISE OR b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VORPS'. Intrinsic: '_mm256_or_ps'.
// Requires AVX.
func OrPs(a M256, b M256) M256 {
	return M256(orPs([8]float32(a), [8]float32(b)))
}

func orPs(a [8]float32, b [8]float32) [8]float32


// PermutePd: Shuffle double-precision (64-bit) floating-point elements in 'a'
// using the control in 'imm8', and store the results in 'dst'. 
//
//		IF (imm8[0] == 0) dst[63:0] := a[63:0]
//		IF (imm8[0] == 1) dst[63:0] := a[127:64]
//		IF (imm8[1] == 0) dst[127:64] := a[63:0]
//		IF (imm8[1] == 1) dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMILPD'. Intrinsic: '_mm_permute_pd'.
// Requires AVX.
func PermutePd(a M128d, imm8 int) M128d {
	return M128d(permutePd([2]float64(a), imm8))
}

func permutePd(a [2]float64, imm8 int) [2]float64


// PermutePd1: Shuffle double-precision (64-bit) floating-point elements in 'a'
// within 128-bit lanes using the control in 'imm8', and store the results in
// 'dst'. 
//
//		IF (imm8[0] == 0) dst[63:0] := a[63:0]
//		IF (imm8[0] == 1) dst[63:0] := a[127:64]
//		IF (imm8[1] == 0) dst[127:64] := a[63:0]
//		IF (imm8[1] == 1) dst[127:64] := a[127:64]
//		IF (imm8[2] == 0) dst[191:128] := a[191:128]
//		IF (imm8[2] == 1) dst[191:128] := a[255:192]
//		IF (imm8[3] == 0) dst[255:192] := a[191:128]
//		IF (imm8[3] == 1) dst[255:192] := a[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMILPD'. Intrinsic: '_mm256_permute_pd'.
// Requires AVX.
func PermutePd1(a M256d, imm8 int) M256d {
	return M256d(permutePd1([4]float64(a), imm8))
}

func permutePd1(a [4]float64, imm8 int) [4]float64


// PermutePs: Shuffle single-precision (32-bit) floating-point elements in 'a'
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
//		dst[95:64] := SELECT4(a[127:0], imm8[5:4])
//		dst[127:96] := SELECT4(a[127:0], imm8[7:6])
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMILPS'. Intrinsic: '_mm_permute_ps'.
// Requires AVX.
func PermutePs(a M128, imm8 int) M128 {
	return M128(permutePs([4]float32(a), imm8))
}

func permutePs(a [4]float32, imm8 int) [4]float32


// PermutePs1: Shuffle single-precision (32-bit) floating-point elements in 'a'
// within 128-bit lanes using the control in 'imm8', and store the results in
// 'dst'. 
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
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMILPS'. Intrinsic: '_mm256_permute_ps'.
// Requires AVX.
func PermutePs1(a M256, imm8 int) M256 {
	return M256(permutePs1([8]float32(a), imm8))
}

func permutePs1(a [8]float32, imm8 int) [8]float32


// Permute2f128Pd: Shuffle 128-bits (composed of 2 packed double-precision
// (64-bit) floating-point elements) selected by 'imm8' from 'a' and 'b', and
// store the results in 'dst'. 
//
//		SELECT4(src1, src2, control){
//			CASE(control[1:0])
//			0:	tmp[127:0] := src1[127:0]
//			1:	tmp[127:0] := src1[255:128]
//			2:	tmp[127:0] := src2[127:0]
//			3:	tmp[127:0] := src2[255:128]
//			ESAC
//			IF control[3]
//				tmp[127:0] := 0
//			FI
//			RETURN tmp[127:0]
//		}
//		
//		dst[127:0] := SELECT4(a[255:0], b[255:0], imm8[3:0])
//		dst[255:128] := SELECT4(a[255:0], b[255:0], imm8[7:4])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERM2F128'. Intrinsic: '_mm256_permute2f128_pd'.
// Requires AVX.
func Permute2f128Pd(a M256d, b M256d, imm8 int) M256d {
	return M256d(permute2f128Pd([4]float64(a), [4]float64(b), imm8))
}

func permute2f128Pd(a [4]float64, b [4]float64, imm8 int) [4]float64


// Permute2f128Ps: Shuffle 128-bits (composed of 4 packed single-precision
// (32-bit) floating-point elements) selected by 'imm8' from 'a' and 'b', and
// store the results in 'dst'. 
//
//		SELECT4(src1, src2, control){
//			CASE(control[1:0])
//			0:	tmp[127:0] := src1[127:0]
//			1:	tmp[127:0] := src1[255:128]
//			2:	tmp[127:0] := src2[127:0]
//			3:	tmp[127:0] := src2[255:128]
//			ESAC
//			IF control[3]
//				tmp[127:0] := 0
//			FI
//			RETURN tmp[127:0]
//		}
//		
//		dst[127:0] := SELECT4(a[255:0], b[255:0], imm8[3:0])
//		dst[255:128] := SELECT4(a[255:0], b[255:0], imm8[7:4])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERM2F128'. Intrinsic: '_mm256_permute2f128_ps'.
// Requires AVX.
func Permute2f128Ps(a M256, b M256, imm8 int) M256 {
	return M256(permute2f128Ps([8]float32(a), [8]float32(b), imm8))
}

func permute2f128Ps(a [8]float32, b [8]float32, imm8 int) [8]float32


// Permute2f128Si256: Shuffle 128-bits (composed of integer data) selected by
// 'imm8' from 'a' and 'b', and store the results in 'dst'. 
//
//		SELECT4(src1, src2, control){
//			CASE(control[1:0])
//			0:	tmp[127:0] := src1[127:0]
//			1:	tmp[127:0] := src1[255:128]
//			2:	tmp[127:0] := src2[127:0]
//			3:	tmp[127:0] := src2[255:128]
//			ESAC
//			IF control[3]
//				tmp[127:0] := 0
//			FI
//			RETURN tmp[127:0]
//		}
//		
//		dst[127:0] := SELECT4(a[255:0], b[255:0], imm8[3:0])
//		dst[255:128] := SELECT4(a[255:0], b[255:0], imm8[7:4])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERM2F128'. Intrinsic: '_mm256_permute2f128_si256'.
// Requires AVX.
func Permute2f128Si256(a M256i, b M256i, imm8 int) M256i {
	return M256i(permute2f128Si256([32]byte(a), [32]byte(b), imm8))
}

func permute2f128Si256(a [32]byte, b [32]byte, imm8 int) [32]byte


// PermutevarPd: Shuffle double-precision (64-bit) floating-point elements in
// 'a' using the control in 'b', and store the results in 'dst'. 
//
//		IF (b[1] == 0) dst[63:0] := a[63:0]
//		IF (b[1] == 1) dst[63:0] := a[127:64]
//		IF (b[65] == 0) dst[127:64] := a[63:0]
//		IF (b[65] == 1) dst[127:64] := a[127:64]
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMILPD'. Intrinsic: '_mm_permutevar_pd'.
// Requires AVX.
func PermutevarPd(a M128d, b M128i) M128d {
	return M128d(permutevarPd([2]float64(a), [16]byte(b)))
}

func permutevarPd(a [2]float64, b [16]byte) [2]float64


// PermutevarPd1: Shuffle double-precision (64-bit) floating-point elements in
// 'a' within 128-bit lanes using the control in 'b', and store the results in
// 'dst'. 
//
//		IF (b[1] == 0) dst[63:0] := a[63:0]
//		IF (b[1] == 1) dst[63:0] := a[127:64]
//		IF (b[65] == 0) dst[127:64] := a[63:0]
//		IF (b[65] == 1) dst[127:64] := a[127:64]
//		IF (b[129] == 0) dst[191:128] := a[191:128]
//		IF (b[129] == 1) dst[191:128] := a[255:192]
//		IF (b[193] == 0) dst[255:192] := a[191:128]
//		IF (b[193] == 1) dst[255:192] := a[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMILPD'. Intrinsic: '_mm256_permutevar_pd'.
// Requires AVX.
func PermutevarPd1(a M256d, b M256i) M256d {
	return M256d(permutevarPd1([4]float64(a), [32]byte(b)))
}

func permutevarPd1(a [4]float64, b [32]byte) [4]float64


// PermutevarPs: Shuffle single-precision (32-bit) floating-point elements in
// 'a' using the control in 'b', and store the results in 'dst'. 
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
//		dst[31:0] := SELECT4(a[127:0], b[1:0])
//		dst[63:32] := SELECT4(a[127:0], b[33:32])
//		dst[95:64] := SELECT4(a[127:0], b[65:64])
//		dst[127:96] := SELECT4(a[127:0], b[97:96])
//		dst[MAX:128] := 0
//
// Instruction: 'VPERMILPS'. Intrinsic: '_mm_permutevar_ps'.
// Requires AVX.
func PermutevarPs(a M128, b M128i) M128 {
	return M128(permutevarPs([4]float32(a), [16]byte(b)))
}

func permutevarPs(a [4]float32, b [16]byte) [4]float32


// PermutevarPs1: Shuffle single-precision (32-bit) floating-point elements in
// 'a' within 128-bit lanes using the control in 'b', and store the results in
// 'dst'. 
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
//		dst[31:0] := SELECT4(a[127:0], b[1:0])
//		dst[63:32] := SELECT4(a[127:0], b[33:32])
//		dst[95:64] := SELECT4(a[127:0], b[65:64])
//		dst[127:96] := SELECT4(a[127:0], b[97:96])
//		dst[159:128] := SELECT4(a[255:128], b[129:128])
//		dst[191:160] := SELECT4(a[255:128], b[161:160])
//		dst[223:192] := SELECT4(a[255:128], b[193:192])
//		dst[255:224] := SELECT4(a[255:128], b[225:224])
//		dst[MAX:256] := 0
//
// Instruction: 'VPERMILPS'. Intrinsic: '_mm256_permutevar_ps'.
// Requires AVX.
func PermutevarPs1(a M256, b M256i) M256 {
	return M256(permutevarPs1([8]float32(a), [32]byte(b)))
}

func permutevarPs1(a [8]float32, b [32]byte) [8]float32


// PowPd: Compute the exponential value of packed double-precision (64-bit)
// floating-point elements in 'a' raised by packed elements in 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := (a[i+63:i])^(b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_pow_pd'.
// Requires AVX.
func PowPd(a M256d, b M256d) M256d {
	return M256d(powPd([4]float64(a), [4]float64(b)))
}

func powPd(a [4]float64, b [4]float64) [4]float64


// PowPs: Compute the exponential value of packed single-precision (32-bit)
// floating-point elements in 'a' raised by packed elements in 'b', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := (a[i+31:i])^(b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_pow_ps'.
// Requires AVX.
func PowPs(a M256, b M256) M256 {
	return M256(powPs([8]float32(a), [8]float32(b)))
}

func powPs(a [8]float32, b [8]float32) [8]float32


// RcpPs: Compute the approximate reciprocal of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. The
// maximum relative error for this approximation is less than 1.5*2^-12. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := APPROXIMATE(1.0/a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRCPPS'. Intrinsic: '_mm256_rcp_ps'.
// Requires AVX.
func RcpPs(a M256) M256 {
	return M256(rcpPs([8]float32(a)))
}

func rcpPs(a [8]float32) [8]float32


// RemEpi16: Divide packed 16-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := REMAINDER(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi16'.
// Requires AVX.
func RemEpi16(a M256i, b M256i) M256i {
	return M256i(remEpi16([32]byte(a), [32]byte(b)))
}

func remEpi16(a [32]byte, b [32]byte) [32]byte


// RemEpi32: Divide packed 32-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi32'.
// Requires AVX.
func RemEpi32(a M256i, b M256i) M256i {
	return M256i(remEpi32([32]byte(a), [32]byte(b)))
}

func remEpi32(a [32]byte, b [32]byte) [32]byte


// RemEpi64: Divide packed 64-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := REMAINDER(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi64'.
// Requires AVX.
func RemEpi64(a M256i, b M256i) M256i {
	return M256i(remEpi64([32]byte(a), [32]byte(b)))
}

func remEpi64(a [32]byte, b [32]byte) [32]byte


// RemEpi8: Divide packed 8-bit integers in 'a' by packed elements in 'b', and
// store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := REMAINDER(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi8'.
// Requires AVX.
func RemEpi8(a M256i, b M256i) M256i {
	return M256i(remEpi8([32]byte(a), [32]byte(b)))
}

func remEpi8(a [32]byte, b [32]byte) [32]byte


// RemEpu16: Divide packed unsigned 16-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := REMAINDER(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu16'.
// Requires AVX.
func RemEpu16(a M256i, b M256i) M256i {
	return M256i(remEpu16([32]byte(a), [32]byte(b)))
}

func remEpu16(a [32]byte, b [32]byte) [32]byte


// RemEpu32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu32'.
// Requires AVX.
func RemEpu32(a M256i, b M256i) M256i {
	return M256i(remEpu32([32]byte(a), [32]byte(b)))
}

func remEpu32(a [32]byte, b [32]byte) [32]byte


// RemEpu64: Divide packed unsigned 64-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := REMAINDER(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu64'.
// Requires AVX.
func RemEpu64(a M256i, b M256i) M256i {
	return M256i(remEpu64([32]byte(a), [32]byte(b)))
}

func remEpu64(a [32]byte, b [32]byte) [32]byte


// RemEpu8: Divide packed unsigned 8-bit integers in 'a' by packed elements in
// 'b', and store the remainders as packed unsigned 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := REMAINDER(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu8'.
// Requires AVX.
func RemEpu8(a M256i, b M256i) M256i {
	return M256i(remEpu8([32]byte(a), [32]byte(b)))
}

func remEpu8(a [32]byte, b [32]byte) [32]byte


// RoundPd: Round the packed double-precision (64-bit) floating-point elements
// in 'a' using the 'rounding' parameter, and store the results as packed
// double-precision floating-point elements in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ROUND(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VROUNDPD'. Intrinsic: '_mm256_round_pd'.
// Requires AVX.
func RoundPd(a M256d, rounding int) M256d {
	return M256d(roundPd([4]float64(a), rounding))
}

func roundPd(a [4]float64, rounding int) [4]float64


// RoundPs: Round the packed single-precision (32-bit) floating-point elements
// in 'a' using the 'rounding' parameter, and store the results as packed
// single-precision floating-point elements in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm256_round_ps'.
// Requires AVX.
func RoundPs(a M256, rounding int) M256 {
	return M256(roundPs([8]float32(a), rounding))
}

func roundPs(a [8]float32, rounding int) [8]float32


// RsqrtPs: Compute the approximate reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. The maximum relative error for this approximation is less
// than 1.5*2^-12. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := APPROXIMATE(1.0 / SQRT(a[i+31:i]))
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VRSQRTPS'. Intrinsic: '_mm256_rsqrt_ps'.
// Requires AVX.
func RsqrtPs(a M256) M256 {
	return M256(rsqrtPs([8]float32(a)))
}

func rsqrtPs(a [8]float32) [8]float32


// SetEpi16: Set packed 16-bit integers in 'dst' with the supplied values. 
//
//		dst[15:0] := e0
//		dst[31:16] := e1
//		dst[47:32] := e2
//		dst[63:48] := e3
//		dst[79:64] := e4
//		dst[95:80] := e5
//		dst[111:96] := e6
//		dst[127:112] := e7
//		dst[145:128] := e8
//		dst[159:144] := e9
//		dst[175:160] := e10
//		dst[191:176] := e11
//		dst[207:192] := e12
//		dst[223:208] := e13
//		dst[239:224] := e14
//		dst[255:240] := e15
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_epi16'.
// Requires AVX.
func SetEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) M256i {
	return M256i(setEpi16(e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func setEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [32]byte


// SetEpi32: Set packed 32-bit integers in 'dst' with the supplied values. 
//
//		dst[31:0] := e0
//		dst[63:32] := e1
//		dst[95:64] := e2
//		dst[127:96] := e3
//		dst[159:128] := e4
//		dst[191:160] := e5
//		dst[223:192] := e6
//		dst[255:224] := e7
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_epi32'.
// Requires AVX.
func SetEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) M256i {
	return M256i(setEpi32(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [32]byte


// SetEpi64x: Set packed 64-bit integers in 'dst' with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//		dst[191:128] := e2
//		dst[255:192] := e3
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_epi64x'.
// Requires AVX.
func SetEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) M256i {
	return M256i(setEpi64x(e3, e2, e1, e0))
}

func setEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) [32]byte


// SetEpi8: Set packed 8-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[7:0] := e0
//		dst[15:8] := e1
//		dst[23:16] := e2
//		dst[31:24] := e3
//		dst[39:32] := e4
//		dst[47:40] := e5
//		dst[55:48] := e6
//		dst[63:56] := e7
//		dst[71:64] := e8
//		dst[79:72] := e9
//		dst[87:80] := e10
//		dst[95:88] := e11
//		dst[103:96] := e12
//		dst[111:104] := e13
//		dst[119:112] := e14
//		dst[127:120] := e15
//		dst[135:128] := e16
//		dst[143:136] := e17
//		dst[151:144] := e18
//		dst[159:152] := e19
//		dst[167:160] := e20
//		dst[175:168] := e21
//		dst[183:176] := e22
//		dst[191:184] := e23
//		dst[199:192] := e24
//		dst[207:200] := e25
//		dst[215:208] := e26
//		dst[223:216] := e27
//		dst[231:224] := e28
//		dst[239:232] := e29
//		dst[247:240] := e30
//		dst[255:248] := e31
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_epi8'.
// Requires AVX.
func SetEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M256i {
	return M256i(setEpi8(e31, e30, e29, e28, e27, e26, e25, e24, e23, e22, e21, e20, e19, e18, e17, e16, e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func setEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [32]byte


// SetM128: Set packed __m256 vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_set_m128'.
// Requires AVX.
func SetM128(hi M128, lo M128) M256 {
	return M256(setM128([4]float32(hi), [4]float32(lo)))
}

func setM128(hi [4]float32, lo [4]float32) [8]float32


// SetM128d: Set packed __m256d vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_set_m128d'.
// Requires AVX.
func SetM128d(hi M128d, lo M128d) M256d {
	return M256d(setM128d([2]float64(hi), [2]float64(lo)))
}

func setM128d(hi [2]float64, lo [2]float64) [4]float64


// SetM128i: Set packed __m256i vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_set_m128i'.
// Requires AVX.
func SetM128i(hi M128i, lo M128i) M256i {
	return M256i(setM128i([16]byte(hi), [16]byte(lo)))
}

func setM128i(hi [16]byte, lo [16]byte) [32]byte


// SetPd: Set packed double-precision (64-bit) floating-point elements in 'dst'
// with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//		dst[191:128] := e2
//		dst[255:192] := e3
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_pd'.
// Requires AVX.
func SetPd(e3 float64, e2 float64, e1 float64, e0 float64) M256d {
	return M256d(setPd(e3, e2, e1, e0))
}

func setPd(e3 float64, e2 float64, e1 float64, e0 float64) [4]float64


// SetPs: Set packed single-precision (32-bit) floating-point elements in 'dst'
// with the supplied values. 
//
//		dst[31:0] := e0
//		dst[63:32] := e1
//		dst[95:64] := e2
//		dst[127:96] := e3
//		dst[159:128] := e4
//		dst[191:160] := e5
//		dst[223:192] := e6
//		dst[255:224] := e7
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_ps'.
// Requires AVX.
func SetPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) M256 {
	return M256(setPs(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [8]float32


// Set1Epi16: Broadcast 16-bit integer 'a' to all all elements of 'dst'. This
// intrinsic may generate the 'vpbroadcastw'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_epi16'.
// Requires AVX.
func Set1Epi16(a int16) M256i {
	return M256i(set1Epi16(a))
}

func set1Epi16(a int16) [32]byte


// Set1Epi32: Broadcast 32-bit integer 'a' to all elements of 'dst'. This
// intrinsic may generate the 'vpbroadcastd'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_epi32'.
// Requires AVX.
func Set1Epi32(a int) M256i {
	return M256i(set1Epi32(a))
}

func set1Epi32(a int) [32]byte


// Set1Epi64x: Broadcast 64-bit integer 'a' to all elements of 'dst'. This
// intrinsic may generate the 'vpbroadcastq'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_epi64x'.
// Requires AVX.
func Set1Epi64x(a int64) M256i {
	return M256i(set1Epi64x(a))
}

func set1Epi64x(a int64) [32]byte


// Set1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst'. This
// intrinsic may generate the 'vpbroadcastb'. 
//
//		FOR j := 0 to 31
//			i := j*8
//			dst[i+7:i] := a[7:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_epi8'.
// Requires AVX.
func Set1Epi8(a byte) M256i {
	return M256i(set1Epi8(a))
}

func set1Epi8(a byte) [32]byte


// Set1Pd: Broadcast double-precision (64-bit) floating-point value 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_pd'.
// Requires AVX.
func Set1Pd(a float64) M256d {
	return M256d(set1Pd(a))
}

func set1Pd(a float64) [4]float64


// Set1Ps: Broadcast single-precision (32-bit) floating-point value 'a' to all
// elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_ps'.
// Requires AVX.
func Set1Ps(a float32) M256 {
	return M256(set1Ps(a))
}

func set1Ps(a float32) [8]float32


// SetrEpi16: Set packed 16-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[15:0] := e15
//		dst[31:16] := e14
//		dst[47:32] := e13
//		dst[63:48] := e12
//		dst[79:64] := e11
//		dst[95:80] := e10
//		dst[111:96] := e9
//		dst[127:112] := e8
//		dst[145:128] := e7
//		dst[159:144] := e6
//		dst[175:160] := e5
//		dst[191:176] := e4
//		dst[207:192] := e3
//		dst[223:208] := e2
//		dst[239:224] := e1
//		dst[255:240] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_epi16'.
// Requires AVX.
func SetrEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) M256i {
	return M256i(setrEpi16(e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [32]byte


// SetrEpi32: Set packed 32-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[31:0] := e7
//		dst[63:32] := e6
//		dst[95:64] := e5
//		dst[127:96] := e4
//		dst[159:128] := e3
//		dst[191:160] := e2
//		dst[223:192] := e1
//		dst[255:224] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_epi32'.
// Requires AVX.
func SetrEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) M256i {
	return M256i(setrEpi32(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [32]byte


// SetrEpi64x: Set packed 64-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[63:0] := e3
//		dst[127:64] := e2
//		dst[191:128] := e1
//		dst[255:192] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_epi64x'.
// Requires AVX.
func SetrEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) M256i {
	return M256i(setrEpi64x(e3, e2, e1, e0))
}

func setrEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) [32]byte


// SetrEpi8: Set packed 8-bit integers in 'dst' with the supplied values in
// reverse order. 
//
//		dst[7:0] := e31
//		dst[15:8] := e30
//		dst[23:16] := e29
//		dst[31:24] := e28
//		dst[39:32] := e27
//		dst[47:40] := e26
//		dst[55:48] := e25
//		dst[63:56] := e24
//		dst[71:64] := e23
//		dst[79:72] := e22
//		dst[87:80] := e21
//		dst[95:88] := e20
//		dst[103:96] := e19
//		dst[111:104] := e18
//		dst[119:112] := e17
//		dst[127:120] := e16
//		dst[135:128] := e15
//		dst[143:136] := e14
//		dst[151:144] := e13
//		dst[159:152] := e12
//		dst[167:160] := e11
//		dst[175:168] := e10
//		dst[183:176] := e9
//		dst[191:184] := e8
//		dst[199:192] := e7
//		dst[207:200] := e6
//		dst[215:208] := e5
//		dst[223:216] := e4
//		dst[231:224] := e3
//		dst[239:232] := e2
//		dst[247:240] := e1
//		dst[255:248] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_epi8'.
// Requires AVX.
func SetrEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) M256i {
	return M256i(setrEpi8(e31, e30, e29, e28, e27, e26, e25, e24, e23, e22, e21, e20, e19, e18, e17, e16, e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [32]byte


// SetrM128: Set packed __m256 vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_setr_m128'.
// Requires AVX.
func SetrM128(lo M128, hi M128) M256 {
	return M256(setrM128([4]float32(lo), [4]float32(hi)))
}

func setrM128(lo [4]float32, hi [4]float32) [8]float32


// SetrM128d: Set packed __m256d vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_setr_m128d'.
// Requires AVX.
func SetrM128d(lo M128d, hi M128d) M256d {
	return M256d(setrM128d([2]float64(lo), [2]float64(hi)))
}

func setrM128d(lo [2]float64, hi [2]float64) [4]float64


// SetrM128i: Set packed __m256i vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_setr_m128i'.
// Requires AVX.
func SetrM128i(lo M128i, hi M128i) M256i {
	return M256i(setrM128i([16]byte(lo), [16]byte(hi)))
}

func setrM128i(lo [16]byte, hi [16]byte) [32]byte


// SetrPd: Set packed double-precision (64-bit) floating-point elements in
// 'dst' with the supplied values in reverse order. 
//
//		dst[63:0] := e3
//		dst[127:64] := e2
//		dst[191:128] := e1
//		dst[255:192] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_pd'.
// Requires AVX.
func SetrPd(e3 float64, e2 float64, e1 float64, e0 float64) M256d {
	return M256d(setrPd(e3, e2, e1, e0))
}

func setrPd(e3 float64, e2 float64, e1 float64, e0 float64) [4]float64


// SetrPs: Set packed single-precision (32-bit) floating-point elements in
// 'dst' with the supplied values in reverse order. 
//
//		dst[31:0] := e7
//		dst[63:32] := e6
//		dst[95:64] := e5
//		dst[127:96] := e4
//		dst[159:128] := e3
//		dst[191:160] := e2
//		dst[223:192] := e1
//		dst[255:224] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_ps'.
// Requires AVX.
func SetrPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) M256 {
	return M256(setrPs(e7, e6, e5, e4, e3, e2, e1, e0))
}

func setrPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [8]float32


// SetzeroPd: Return vector of type __m256d with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm256_setzero_pd'.
// Requires AVX.
func SetzeroPd() M256d {
	return M256d(setzeroPd())
}

func setzeroPd() [4]float64


// SetzeroPs: Return vector of type __m256 with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm256_setzero_ps'.
// Requires AVX.
func SetzeroPs() M256 {
	return M256(setzeroPs())
}

func setzeroPs() [8]float32


// SetzeroSi256: Return vector of type __m256i with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'VPXOR'. Intrinsic: '_mm256_setzero_si256'.
// Requires AVX.
func SetzeroSi256() M256i {
	return M256i(setzeroSi256())
}

func setzeroSi256() [32]byte


// ShufflePd: Shuffle double-precision (64-bit) floating-point elements within
// 128-bit lanes using the control in 'imm8', and store the results in 'dst'. 
//
//		dst[63:0] := (imm8[0] == 0) ? a[63:0] : a[127:64]
//		dst[127:64] := (imm8[1] == 0) ? b[63:0] : b[127:64]
//		dst[191:128] := (imm8[2] == 0) ? a[191:128] : a[255:192]
//		dst[255:192] := (imm8[3] == 0) ? b[191:128] : b[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VSHUFPD'. Intrinsic: '_mm256_shuffle_pd'.
// Requires AVX.
func ShufflePd(a M256d, b M256d, imm8 int) M256d {
	return M256d(shufflePd([4]float64(a), [4]float64(b), imm8))
}

func shufflePd(a [4]float64, b [4]float64, imm8 int) [4]float64


// ShufflePs: Shuffle single-precision (32-bit) floating-point elements in 'a'
// within 128-bit lanes using the control in 'imm8', and store the results in
// 'dst'. 
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
//		dst[159:128] := SELECT4(a[255:128], imm8[1:0])
//		dst[191:160] := SELECT4(a[255:128], imm8[3:2])
//		dst[223:192] := SELECT4(b[255:128], imm8[5:4])
//		dst[255:224] := SELECT4(b[255:128], imm8[7:6])
//		dst[MAX:256] := 0
//
// Instruction: 'VSHUFPS'. Intrinsic: '_mm256_shuffle_ps'.
// Requires AVX.
func ShufflePs(a M256, b M256, imm8 int) M256 {
	return M256(shufflePs([8]float32(a), [8]float32(b), imm8))
}

func shufflePs(a [8]float32, b [8]float32, imm8 int) [8]float32


// SinPd: Compute the sine of packed double-precision (64-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SIN(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sin_pd'.
// Requires AVX.
func SinPd(a M256d) M256d {
	return M256d(sinPd([4]float64(a)))
}

func sinPd(a [4]float64) [4]float64


// SinPs: Compute the sine of packed single-precision (32-bit) floating-point
// elements in 'a' expressed in radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SIN(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sin_ps'.
// Requires AVX.
func SinPs(a M256) M256 {
	return M256(sinPs([8]float32(a)))
}

func sinPs(a [8]float32) [8]float32


// SincosPd: Compute the sine and cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, store the sine in
// 'dst', and store the cosine into memory at 'mem_addr'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SIN(a[i+63:i])
//			MEM[mem_addr+i+63:mem_addr+i] := COS(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sincos_pd'.
// Requires AVX.
func SincosPd(mem_addr M256d, a M256d) M256d {
	return M256d(sincosPd([4]float64(mem_addr), [4]float64(a)))
}

func sincosPd(mem_addr [4]float64, a [4]float64) [4]float64


// SincosPs: Compute the sine and cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, store the sine in
// 'dst', and store the cosine into memory at 'mem_addr'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SIN(a[i+31:i])
//			MEM[mem_addr+i+31:mem_addr+i] := COS(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sincos_ps'.
// Requires AVX.
func SincosPs(mem_addr M256, a M256) M256 {
	return M256(sincosPs([8]float32(mem_addr), [8]float32(a)))
}

func sincosPs(mem_addr [8]float32, a [8]float32) [8]float32


// SindPd: Compute the sine of packed double-precision (64-bit) floating-point
// elements in 'a' expressed in degrees, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SIND(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sind_pd'.
// Requires AVX.
func SindPd(a M256d) M256d {
	return M256d(sindPd([4]float64(a)))
}

func sindPd(a [4]float64) [4]float64


// SindPs: Compute the sine of packed single-precision (32-bit) floating-point
// elements in 'a' expressed in degrees, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SIND(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sind_ps'.
// Requires AVX.
func SindPs(a M256) M256 {
	return M256(sindPs([8]float32(a)))
}

func sindPs(a [8]float32) [8]float32


// SinhPd: Compute the hyperbolic sine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SINH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sinh_pd'.
// Requires AVX.
func SinhPd(a M256d) M256d {
	return M256d(sinhPd([4]float64(a)))
}

func sinhPd(a [4]float64) [4]float64


// SinhPs: Compute the hyperbolic sine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SINH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sinh_ps'.
// Requires AVX.
func SinhPs(a M256) M256 {
	return M256(sinhPs([8]float32(a)))
}

func sinhPs(a [8]float32) [8]float32


// SqrtPd: Compute the square root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SQRT(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VSQRTPD'. Intrinsic: '_mm256_sqrt_pd'.
// Requires AVX.
func SqrtPd(a M256d) M256d {
	return M256d(sqrtPd([4]float64(a)))
}

func sqrtPd(a [4]float64) [4]float64


// SqrtPs: Compute the square root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VSQRTPS'. Intrinsic: '_mm256_sqrt_ps'.
// Requires AVX.
func SqrtPs(a M256) M256 {
	return M256(sqrtPs([8]float32(a)))
}

func sqrtPs(a [8]float32) [8]float32


// StorePd: Store 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm256_store_pd'.
// Requires AVX.
func StorePd(mem_addr float64, a M256d)  {
	storePd(mem_addr, [4]float64(a))
}

func storePd(mem_addr float64, a [4]float64) 


// StorePs: Store 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm256_store_ps'.
// Requires AVX.
func StorePs(mem_addr float32, a M256)  {
	storePs(mem_addr, [8]float32(a))
}

func storePs(mem_addr float32, a [8]float32) 


// StoreSi256: Store 256-bits of integer data from 'a' into memory.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVDQA'. Intrinsic: '_mm256_store_si256'.
// Requires AVX.
func StoreSi256(mem_addr M256i, a M256i)  {
	storeSi256([32]byte(mem_addr), [32]byte(a))
}

func storeSi256(mem_addr [32]byte, a [32]byte) 


// StoreuPd: Store 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVUPD'. Intrinsic: '_mm256_storeu_pd'.
// Requires AVX.
func StoreuPd(mem_addr float64, a M256d)  {
	storeuPd(mem_addr, [4]float64(a))
}

func storeuPd(mem_addr float64, a [4]float64) 


// StoreuPs: Store 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVUPS'. Intrinsic: '_mm256_storeu_ps'.
// Requires AVX.
func StoreuPs(mem_addr float32, a M256)  {
	storeuPs(mem_addr, [8]float32(a))
}

func storeuPs(mem_addr float32, a [8]float32) 


// StoreuSi256: Store 256-bits of integer data from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVDQU'. Intrinsic: '_mm256_storeu_si256'.
// Requires AVX.
func StoreuSi256(mem_addr M256i, a M256i)  {
	storeuSi256([32]byte(mem_addr), [32]byte(a))
}

func storeuSi256(mem_addr [32]byte, a [32]byte) 


// Storeu2M128: Store the high and low 128-bit halves (each composed of 4
// packed single-precision (32-bit) floating-point elements) from 'a' into
// memory two different 128-bit locations.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		MEM[loaddr+127:loaddr] := a[127:0]
//		MEM[hiaddr+127:hiaddr] := a[255:128]
//
// Instruction: '...'. Intrinsic: '_mm256_storeu2_m128'.
// Requires AVX.
func Storeu2M128(hiaddr float32, loaddr float32, a M256)  {
	storeu2M128(hiaddr, loaddr, [8]float32(a))
}

func storeu2M128(hiaddr float32, loaddr float32, a [8]float32) 


// Storeu2M128d: Store the high and low 128-bit halves (each composed of 2
// packed double-precision (64-bit) floating-point elements) from 'a' into
// memory two different 128-bit locations.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		MEM[loaddr+127:loaddr] := a[127:0]
//		MEM[hiaddr+127:hiaddr] := a[255:128]
//
// Instruction: '...'. Intrinsic: '_mm256_storeu2_m128d'.
// Requires AVX.
func Storeu2M128d(hiaddr float64, loaddr float64, a M256d)  {
	storeu2M128d(hiaddr, loaddr, [4]float64(a))
}

func storeu2M128d(hiaddr float64, loaddr float64, a [4]float64) 


// Storeu2M128i: Store the high and low 128-bit halves (each composed of
// integer data) from 'a' into memory two different 128-bit locations.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		MEM[loaddr+127:loaddr] := a[127:0]
//		MEM[hiaddr+127:hiaddr] := a[255:128]
//
// Instruction: '...'. Intrinsic: '_mm256_storeu2_m128i'.
// Requires AVX.
func Storeu2M128i(hiaddr M128i, loaddr M128i, a M256i)  {
	storeu2M128i([16]byte(hiaddr), [16]byte(loaddr), [32]byte(a))
}

func storeu2M128i(hiaddr [16]byte, loaddr [16]byte, a [32]byte) 


// StreamPd: Store 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory using a non-temporal memory
// hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVNTPD'. Intrinsic: '_mm256_stream_pd'.
// Requires AVX.
func StreamPd(mem_addr float64, a M256d)  {
	streamPd(mem_addr, [4]float64(a))
}

func streamPd(mem_addr float64, a [4]float64) 


// StreamPs: Store 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory using a non-temporal memory
// hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVNTPS'. Intrinsic: '_mm256_stream_ps'.
// Requires AVX.
func StreamPs(mem_addr float32, a M256)  {
	streamPs(mem_addr, [8]float32(a))
}

func streamPs(mem_addr float32, a [8]float32) 


// StreamSi256: Store 256-bits of integer data from 'a' into memory using a
// non-temporal memory hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVNTDQ'. Intrinsic: '_mm256_stream_si256'.
// Requires AVX.
func StreamSi256(mem_addr M256i, a M256i)  {
	streamSi256([32]byte(mem_addr), [32]byte(a))
}

func streamSi256(mem_addr [32]byte, a [32]byte) 


// SubPd: Subtract packed double-precision (64-bit) floating-point elements in
// 'b' from packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm256_sub_pd'.
// Requires AVX.
func SubPd(a M256d, b M256d) M256d {
	return M256d(subPd([4]float64(a), [4]float64(b)))
}

func subPd(a [4]float64, b [4]float64) [4]float64


// SubPs: Subtract packed single-precision (32-bit) floating-point elements in
// 'b' from packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm256_sub_ps'.
// Requires AVX.
func SubPs(a M256, b M256) M256 {
	return M256(subPs([8]float32(a), [8]float32(b)))
}

func subPs(a [8]float32, b [8]float32) [8]float32


// SvmlCeilPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' up to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := CEIL(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_ceil_pd'.
// Requires AVX.
func SvmlCeilPd(a M256d) M256d {
	return M256d(svmlCeilPd([4]float64(a)))
}

func svmlCeilPd(a [4]float64) [4]float64


// SvmlCeilPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' up to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := CEIL(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_ceil_ps'.
// Requires AVX.
func SvmlCeilPs(a M256) M256 {
	return M256(svmlCeilPs([8]float32(a)))
}

func svmlCeilPs(a [8]float32) [8]float32


// SvmlFloorPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' down to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := FLOOR(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_floor_pd'.
// Requires AVX.
func SvmlFloorPd(a M256d) M256d {
	return M256d(svmlFloorPd([4]float64(a)))
}

func svmlFloorPd(a [4]float64) [4]float64


// SvmlFloorPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' down to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := FLOOR(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_floor_ps'.
// Requires AVX.
func SvmlFloorPs(a M256) M256 {
	return M256(svmlFloorPs([8]float32(a)))
}

func svmlFloorPs(a [8]float32) [8]float32


// SvmlRoundPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' to the nearest integer value, and store the results as
// packed double-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ROUND(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_round_pd'.
// Requires AVX.
func SvmlRoundPd(a M256d) M256d {
	return M256d(svmlRoundPd([4]float64(a)))
}

func svmlRoundPd(a [4]float64) [4]float64


// SvmlRoundPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' to the nearest integer value, and store the results as
// packed single-precision floating-point elements in 'dst'. This intrinsic may
// generate the 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_round_ps'.
// Requires AVX.
func SvmlRoundPs(a M256) M256 {
	return M256(svmlRoundPs([8]float32(a)))
}

func svmlRoundPs(a [8]float32) [8]float32


// SvmlSqrtPd: Compute the square root of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. Note that
// this intrinsic is less efficient than '_mm_sqrt_pd'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SQRT(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_sqrt_pd'.
// Requires AVX.
func SvmlSqrtPd(a M256d) M256d {
	return M256d(svmlSqrtPd([4]float64(a)))
}

func svmlSqrtPd(a [4]float64) [4]float64


// SvmlSqrtPs: Compute the square root of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. Note that
// this intrinsic is less efficient than '_mm_sqrt_ps'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_svml_sqrt_ps'.
// Requires AVX.
func SvmlSqrtPs(a M256) M256 {
	return M256(svmlSqrtPs([8]float32(a)))
}

func svmlSqrtPs(a [8]float32) [8]float32


// TanPd: Compute the tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := TAN(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tan_pd'.
// Requires AVX.
func TanPd(a M256d) M256d {
	return M256d(tanPd([4]float64(a)))
}

func tanPd(a [4]float64) [4]float64


// TanPs: Compute the tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := TAN(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tan_ps'.
// Requires AVX.
func TanPs(a M256) M256 {
	return M256(tanPs([8]float32(a)))
}

func tanPs(a [8]float32) [8]float32


// TandPd: Compute the tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := TAND(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tand_pd'.
// Requires AVX.
func TandPd(a M256d) M256d {
	return M256d(tandPd([4]float64(a)))
}

func tandPd(a [4]float64) [4]float64


// TandPs: Compute the tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := TAND(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tand_ps'.
// Requires AVX.
func TandPs(a M256) M256 {
	return M256(tandPs([8]float32(a)))
}

func tandPs(a [8]float32) [8]float32


// TanhPd: Compute the hyperbolic tangent of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := TANH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tanh_pd'.
// Requires AVX.
func TanhPd(a M256d) M256d {
	return M256d(tanhPd([4]float64(a)))
}

func tanhPd(a [4]float64) [4]float64


// TanhPs: Compute the hyperbolic tangent of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := TANH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tanh_ps'.
// Requires AVX.
func TanhPs(a M256) M256 {
	return M256(tanhPs([8]float32(a)))
}

func tanhPs(a [8]float32) [8]float32


// TestcPd: Compute the bitwise AND of 128 bits (representing double-precision
// (64-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 128-bit value, and set 'ZF' to 1 if the sign bit of each 64-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 64-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'CF' value. 
//
//		tmp[127:0] := a[127:0] AND b[127:0]
//		IF (tmp[63] == tmp[127] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[127:0] := a[127:0] AND NOT b[127:0]
//		IF (tmp[63] == tmp[127] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: 'VTESTPD'. Intrinsic: '_mm_testc_pd'.
// Requires AVX.
func TestcPd(a M128d, b M128d) int {
	return int(testcPd([2]float64(a), [2]float64(b)))
}

func testcPd(a [2]float64, b [2]float64) int


// TestcPd1: Compute the bitwise AND of 256 bits (representing double-precision
// (64-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 256-bit value, and set 'ZF' to 1 if the sign bit of each 64-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 64-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'CF' value. 
//
//		tmp[255:0] := a[255:0] AND b[255:0]
//		IF (tmp[63] == tmp[127] == tmp[191] == tmp[255] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[255:0] := a[255:0] AND NOT b[255:0]
//		IF (tmp[63] == tmp[127] == tmp[191] == tmp[255] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: 'VTESTPD'. Intrinsic: '_mm256_testc_pd'.
// Requires AVX.
func TestcPd1(a M256d, b M256d) int {
	return int(testcPd1([4]float64(a), [4]float64(b)))
}

func testcPd1(a [4]float64, b [4]float64) int


// TestcPs: Compute the bitwise AND of 128 bits (representing single-precision
// (32-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 128-bit value, and set 'ZF' to 1 if the sign bit of each 32-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 32-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'CF' value. 
//
//		tmp[127:0] := a[127:0] AND b[127:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[127:0] := a[127:0] AND NOT b[127:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: 'VTESTPS'. Intrinsic: '_mm_testc_ps'.
// Requires AVX.
func TestcPs(a M128, b M128) int {
	return int(testcPs([4]float32(a), [4]float32(b)))
}

func testcPs(a [4]float32, b [4]float32) int


// TestcPs1: Compute the bitwise AND of 256 bits (representing single-precision
// (32-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 256-bit value, and set 'ZF' to 1 if the sign bit of each 32-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 32-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'CF' value. 
//
//		tmp[255:0] := a[255:0] AND b[255:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == tmp[159] == tmp[191] == tmp[223] == tmp[255] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[255:0] := a[255:0] AND NOT b[255:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == tmp[159] == tmp[191] == tmp[223] == tmp[255] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: 'VTESTPS'. Intrinsic: '_mm256_testc_ps'.
// Requires AVX.
func TestcPs1(a M256, b M256) int {
	return int(testcPs1([8]float32(a), [8]float32(b)))
}

func testcPs1(a [8]float32, b [8]float32) int


// TestcSi256: Compute the bitwise AND of 256 bits (representing integer data)
// in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set 'ZF'
// to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if the
// result is zero, otherwise set 'CF' to 0. Return the 'CF' value. 
//
//		IF (a[255:0] AND b[255:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[255:0] AND NOT b[255:0] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: 'VPTEST'. Intrinsic: '_mm256_testc_si256'.
// Requires AVX.
func TestcSi256(a M256i, b M256i) int {
	return int(testcSi256([32]byte(a), [32]byte(b)))
}

func testcSi256(a [32]byte, b [32]byte) int


// TestnzcPd: Compute the bitwise AND of 128 bits (representing
// double-precision (64-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 128-bit value, and set 'ZF' to 1 if the sign bit of each
// 64-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 64-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return 1 if both the 'ZF' and 'CF'
// values are zero, otherwise return 0. 
//
//		tmp[127:0] := a[127:0] AND b[127:0]
//		IF (tmp[63] == tmp[127] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[127:0] := a[127:0] AND NOT b[127:0]
//		IF (tmp[63] == tmp[127] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		IF (ZF == 0 && CF == 0)
//			RETURN 1
//		ELSE
//			RETURN 0
//		FI
//
// Instruction: 'VTESTPD'. Intrinsic: '_mm_testnzc_pd'.
// Requires AVX.
func TestnzcPd(a M128d, b M128d) int {
	return int(testnzcPd([2]float64(a), [2]float64(b)))
}

func testnzcPd(a [2]float64, b [2]float64) int


// TestnzcPd1: Compute the bitwise AND of 256 bits (representing
// double-precision (64-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 256-bit value, and set 'ZF' to 1 if the sign bit of each
// 64-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 64-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return 1 if both the 'ZF' and 'CF'
// values are zero, otherwise return 0. 
//
//		tmp[255:0] := a[255:0] AND b[255:0]
//		IF (tmp[63] == tmp[127] == tmp[191] == tmp[255] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[255:0] := a[255:0] AND NOT b[255:0]
//		IF (tmp[63] == tmp[127] == tmp[191] == tmp[255] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		IF (ZF == 0 && CF == 0)
//			RETURN 1
//		ELSE
//			RETURN 0
//		FI
//
// Instruction: 'VTESTPD'. Intrinsic: '_mm256_testnzc_pd'.
// Requires AVX.
func TestnzcPd1(a M256d, b M256d) int {
	return int(testnzcPd1([4]float64(a), [4]float64(b)))
}

func testnzcPd1(a [4]float64, b [4]float64) int


// TestnzcPs: Compute the bitwise AND of 128 bits (representing
// single-precision (32-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 128-bit value, and set 'ZF' to 1 if the sign bit of each
// 32-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 32-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return 1 if both the 'ZF' and 'CF'
// values are zero, otherwise return 0. 
//
//		tmp[127:0] := a[127:0] AND b[127:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[127:0] := a[127:0] AND NOT b[127:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		IF (ZF == 0 && CF == 0)
//			RETURN 1
//		ELSE
//			RETURN 0
//		FI
//
// Instruction: 'VTESTPS'. Intrinsic: '_mm_testnzc_ps'.
// Requires AVX.
func TestnzcPs(a M128, b M128) int {
	return int(testnzcPs([4]float32(a), [4]float32(b)))
}

func testnzcPs(a [4]float32, b [4]float32) int


// TestnzcPs1: Compute the bitwise AND of 256 bits (representing
// single-precision (32-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 256-bit value, and set 'ZF' to 1 if the sign bit of each
// 32-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 32-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return 1 if both the 'ZF' and 'CF'
// values are zero, otherwise return 0. 
//
//		tmp[255:0] := a[255:0] AND b[255:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == tmp[159] == tmp[191] == tmp[223] == tmp[255]  == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[255:0] := a[255:0] AND NOT b[255:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == tmp[159] == tmp[191] == tmp[223] == tmp[255]  == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		IF (ZF == 0 && CF == 0)
//			RETURN 1
//		ELSE
//			RETURN 0
//		FI
//
// Instruction: 'VTESTPS'. Intrinsic: '_mm256_testnzc_ps'.
// Requires AVX.
func TestnzcPs1(a M256, b M256) int {
	return int(testnzcPs1([8]float32(a), [8]float32(b)))
}

func testnzcPs1(a [8]float32, b [8]float32) int


// TestnzcSi256: Compute the bitwise AND of 256 bits (representing integer
// data) in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set
// 'ZF' to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if
// the result is zero, otherwise set 'CF' to 0. Return 1 if both the 'ZF' and
// 'CF' values are zero, otherwise return 0. 
//
//		IF (a[255:0] AND b[255:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[255:0] AND NOT b[255:0] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		IF (ZF == 0 && CF == 0)
//			RETURN 1
//		ELSE
//			RETURN 0
//		FI
//
// Instruction: 'VPTEST'. Intrinsic: '_mm256_testnzc_si256'.
// Requires AVX.
func TestnzcSi256(a M256i, b M256i) int {
	return int(testnzcSi256([32]byte(a), [32]byte(b)))
}

func testnzcSi256(a [32]byte, b [32]byte) int


// TestzPd: Compute the bitwise AND of 128 bits (representing double-precision
// (64-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 128-bit value, and set 'ZF' to 1 if the sign bit of each 64-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 64-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'ZF' value. 
//
//		tmp[127:0] := a[127:0] AND b[127:0]
//		IF (tmp[63] == tmp[127] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[127:0] := a[127:0] AND NOT b[127:0]
//		IF (tmp[63] == tmp[127] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'VTESTPD'. Intrinsic: '_mm_testz_pd'.
// Requires AVX.
func TestzPd(a M128d, b M128d) int {
	return int(testzPd([2]float64(a), [2]float64(b)))
}

func testzPd(a [2]float64, b [2]float64) int


// TestzPd1: Compute the bitwise AND of 256 bits (representing double-precision
// (64-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 256-bit value, and set 'ZF' to 1 if the sign bit of each 64-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 64-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'ZF' value. 
//
//		tmp[255:0] := a[255:0] AND b[255:0]
//		IF (tmp[63] == tmp[127] == tmp[191] == tmp[255] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[255:0] := a[255:0] AND NOT b[255:0]
//		IF (tmp[63] == tmp[127] == tmp[191] == tmp[255] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'VTESTPD'. Intrinsic: '_mm256_testz_pd'.
// Requires AVX.
func TestzPd1(a M256d, b M256d) int {
	return int(testzPd1([4]float64(a), [4]float64(b)))
}

func testzPd1(a [4]float64, b [4]float64) int


// TestzPs: Compute the bitwise AND of 128 bits (representing single-precision
// (32-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 128-bit value, and set 'ZF' to 1 if the sign bit of each 32-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 32-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'ZF' value. 
//
//		tmp[127:0] := a[127:0] AND b[127:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[127:0] := a[127:0] AND NOT b[127:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'VTESTPS'. Intrinsic: '_mm_testz_ps'.
// Requires AVX.
func TestzPs(a M128, b M128) int {
	return int(testzPs([4]float32(a), [4]float32(b)))
}

func testzPs(a [4]float32, b [4]float32) int


// TestzPs1: Compute the bitwise AND of 256 bits (representing single-precision
// (32-bit) floating-point elements) in 'a' and 'b', producing an intermediate
// 256-bit value, and set 'ZF' to 1 if the sign bit of each 32-bit element in
// the intermediate value is zero, otherwise set 'ZF' to 0. Compute the bitwise
// AND NOT of 'a' and 'b', producing an intermediate value, and set 'CF' to 1
// if the sign bit of each 32-bit element in the intermediate value is zero,
// otherwise set 'CF' to 0. Return the 'ZF' value. 
//
//		tmp[255:0] := a[255:0] AND b[255:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == tmp[159] == tmp[191] == tmp[223] == tmp[255] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		tmp[255:0] := a[255:0] AND NOT b[255:0]
//		IF (tmp[31] == tmp[63] == tmp[95] == tmp[127] == tmp[159] == tmp[191] == tmp[223] == tmp[255] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'VTESTPS'. Intrinsic: '_mm256_testz_ps'.
// Requires AVX.
func TestzPs1(a M256, b M256) int {
	return int(testzPs1([8]float32(a), [8]float32(b)))
}

func testzPs1(a [8]float32, b [8]float32) int


// TestzSi256: Compute the bitwise AND of 256 bits (representing integer data)
// in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set 'ZF'
// to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if the
// result is zero, otherwise set 'CF' to 0. Return the 'ZF' value. 
//
//		IF (a[255:0] AND b[255:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[255:0] AND NOT b[255:0] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'VPTEST'. Intrinsic: '_mm256_testz_si256'.
// Requires AVX.
func TestzSi256(a M256i, b M256i) int {
	return int(testzSi256([32]byte(a), [32]byte(b)))
}

func testzSi256(a [32]byte, b [32]byte) int


// TruncPd: Truncate the packed double-precision (64-bit) floating-point
// elements in 'a', and store the results as packed double-precision
// floating-point elements in 'dst'. This intrinsic may generate the
// 'roundpd'/'vroundpd' instruction. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := TRUNCATE(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_trunc_pd'.
// Requires AVX.
func TruncPd(a M256d) M256d {
	return M256d(truncPd([4]float64(a)))
}

func truncPd(a [4]float64) [4]float64


// TruncPs: Truncate the packed single-precision (32-bit) floating-point
// elements in 'a', and store the results as packed single-precision
// floating-point elements in 'dst'. This intrinsic may generate the
// 'roundps'/'vroundps' instruction. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := TRUNCATE(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_trunc_ps'.
// Requires AVX.
func TruncPs(a M256) M256 {
	return M256(truncPs([8]float32(a)))
}

func truncPs(a [8]float32) [8]float32


// UdivEpi32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_udiv_epi32'.
// Requires AVX.
func UdivEpi32(a M256i, b M256i) M256i {
	return M256i(udivEpi32([32]byte(a), [32]byte(b)))
}

func udivEpi32(a [32]byte, b [32]byte) [32]byte


// UdivremEpi32: Divide packed unsigned 32-bit integers in 'a' by packed
// elements in 'b', store the truncated results in 'dst', and store the
// remainders as packed unsigned 32-bit integers into memory at 'mem_addr'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//			MEM[mem_addr+i+31:mem_addr+i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_udivrem_epi32'.
// Requires AVX.
func UdivremEpi32(mem_addr M256i, a M256i, b M256i) M256i {
	return M256i(udivremEpi32([32]byte(mem_addr), [32]byte(a), [32]byte(b)))
}

func udivremEpi32(mem_addr [32]byte, a [32]byte, b [32]byte) [32]byte


// UndefinedPd: Return vector of type __m128d with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_undefined_pd'.
// Requires AVX.
func UndefinedPd() M128d {
	return M128d(undefinedPd())
}

func undefinedPd() [2]float64


// UndefinedPd1: Return vector of type __m256d with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_undefined_pd'.
// Requires AVX.
func UndefinedPd1() M256d {
	return M256d(undefinedPd1())
}

func undefinedPd1() [4]float64


// UndefinedPs: Return vector of type __m128 with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_undefined_ps'.
// Requires AVX.
func UndefinedPs() M128 {
	return M128(undefinedPs())
}

func undefinedPs() [4]float32


// UndefinedPs1: Return vector of type __m256 with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_undefined_ps'.
// Requires AVX.
func UndefinedPs1() M256 {
	return M256(undefinedPs1())
}

func undefinedPs1() [8]float32


// UndefinedSi128: Return vector of type __m128i with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_undefined_si128'.
// Requires AVX.
func UndefinedSi128() M128i {
	return M128i(undefinedSi128())
}

func undefinedSi128() [16]byte


// UndefinedSi256: Return vector of type __m256i with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_undefined_si256'.
// Requires AVX.
func UndefinedSi256() M256i {
	return M256i(undefinedSi256())
}

func undefinedSi256() [32]byte


// UnpackhiPd: Unpack and interleave double-precision (64-bit) floating-point
// elements from the high half of each 128-bit lane in 'a' and 'b', and store
// the results in 'dst'. 
//
//		INTERLEAVE_HIGH_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[127:64] 
//			dst[127:64] := src2[127:64] 
//			RETURN dst[127:0]
//		}	
//		
//		dst[127:0] := INTERLEAVE_HIGH_QWORDS(a[127:0], b[127:0])
//		dst[255:128] := INTERLEAVE_HIGH_QWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VUNPCKHPD'. Intrinsic: '_mm256_unpackhi_pd'.
// Requires AVX.
func UnpackhiPd(a M256d, b M256d) M256d {
	return M256d(unpackhiPd([4]float64(a), [4]float64(b)))
}

func unpackhiPd(a [4]float64, b [4]float64) [4]float64


// UnpackhiPs: Unpack and interleave single-precision (32-bit) floating-point
// elements from the high half of each 128-bit lane in 'a' and 'b', and store
// the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_HIGH_DWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VUNPCKHPS'. Intrinsic: '_mm256_unpackhi_ps'.
// Requires AVX.
func UnpackhiPs(a M256, b M256) M256 {
	return M256(unpackhiPs([8]float32(a), [8]float32(b)))
}

func unpackhiPs(a [8]float32, b [8]float32) [8]float32


// UnpackloPd: Unpack and interleave double-precision (64-bit) floating-point
// elements from the low half of each 128-bit lane in 'a' and 'b', and store
// the results in 'dst'. 
//
//		INTERLEAVE_QWORDS(src1[127:0], src2[127:0]){
//			dst[63:0] := src1[63:0] 
//			dst[127:64] := src2[63:0] 
//			RETURN dst[127:0]
//		}
//		
//		dst[127:0] := INTERLEAVE_QWORDS(a[127:0], b[127:0])
//		dst[255:128] := INTERLEAVE_QWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VUNPCKLPD'. Intrinsic: '_mm256_unpacklo_pd'.
// Requires AVX.
func UnpackloPd(a M256d, b M256d) M256d {
	return M256d(unpackloPd([4]float64(a), [4]float64(b)))
}

func unpackloPd(a [4]float64, b [4]float64) [4]float64


// UnpackloPs: Unpack and interleave single-precision (32-bit) floating-point
// elements from the low half of each 128-bit lane in 'a' and 'b', and store
// the results in 'dst'. 
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
//		dst[255:128] := INTERLEAVE_DWORDS(a[255:128], b[255:128])
//		dst[MAX:256] := 0
//
// Instruction: 'VUNPCKLPS'. Intrinsic: '_mm256_unpacklo_ps'.
// Requires AVX.
func UnpackloPs(a M256, b M256) M256 {
	return M256(unpackloPs([8]float32(a), [8]float32(b)))
}

func unpackloPs(a [8]float32, b [8]float32) [8]float32


// UremEpi32: Divide packed unsigned 32-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_urem_epi32'.
// Requires AVX.
func UremEpi32(a M256i, b M256i) M256i {
	return M256i(uremEpi32([32]byte(a), [32]byte(b)))
}

func uremEpi32(a [32]byte, b [32]byte) [32]byte


// XorPd: Compute the bitwise XOR of packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm256_xor_pd'.
// Requires AVX.
func XorPd(a M256d, b M256d) M256d {
	return M256d(xorPd([4]float64(a), [4]float64(b)))
}

func xorPd(a [4]float64, b [4]float64) [4]float64


// XorPs: Compute the bitwise XOR of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm256_xor_ps'.
// Requires AVX.
func XorPs(a M256, b M256) M256 {
	return M256(xorPs([8]float32(a), [8]float32(b)))
}

func xorPs(a [8]float32, b [8]float32) [8]float32


// Zeroall: Zero the contents of all XMM or YMM registers. 
//
//		YMM0[MAX:0] := 0
//		YMM1[MAX:0] := 0
//		YMM2[MAX:0] := 0
//		YMM3[MAX:0] := 0
//		YMM4[MAX:0] := 0
//		YMM5[MAX:0] := 0
//		YMM6[MAX:0] := 0
//		YMM7[MAX:0] := 0
//		IF 64-bit mode
//			YMM8[MAX:0] := 0
//			YMM9[MAX:0] := 0
//			YMM10[MAX:0] := 0
//			YMM11[MAX:0] := 0
//			YMM12[MAX:0] := 0
//			YMM13[MAX:0] := 0
//			YMM14[MAX:0] := 0
//			YMM15[MAX:0] := 0
//		FI
//
// Instruction: 'VZEROALL'. Intrinsic: '_mm256_zeroall'.
// Requires AVX.
func Zeroall()  {
	zeroall()
}

func zeroall() 


// Zeroupper: Zero the upper 128 bits of all YMM registers; the lower 128-bits
// of the registers are unmodified. 
//
//		YMM0[MAX:128] := 0
//		YMM1[MAX:128] := 0
//		YMM2[MAX:128] := 0
//		YMM3[MAX:128] := 0
//		YMM4[MAX:128] := 0
//		YMM5[MAX:128] := 0
//		YMM6[MAX:128] := 0
//		YMM7[MAX:128] := 0
//		IF 64-bit mode
//			YMM8[MAX:128] := 0
//			YMM9[MAX:128] := 0
//			YMM10[MAX:128] := 0
//			YMM11[MAX:128] := 0
//			YMM12[MAX:128] := 0
//			YMM13[MAX:128] := 0
//			YMM14[MAX:128] := 0
//			YMM15[MAX:128] := 0
//		FI
//
// Instruction: 'VZEROUPPER'. Intrinsic: '_mm256_zeroupper'.
// Requires AVX.
func Zeroupper()  {
	zeroupper()
}

func zeroupper() 

