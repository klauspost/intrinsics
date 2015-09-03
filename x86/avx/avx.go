package avx

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// M256AcosPd: Compute the inverse cosine of packed double-precision (64-bit)
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
func M256AcosPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AcosPd([4]float64(a)))
}

func m256AcosPd(a [4]float64) [4]float64


// M256AcosPs: Compute the inverse cosine of packed single-precision (32-bit)
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
func M256AcosPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256AcosPs([8]float32(a)))
}

func m256AcosPs(a [8]float32) [8]float32


// M256AcoshPd: Compute the inverse hyperbolic cosine of packed
// double-precision (64-bit) floating-point elements in 'a' expressed in
// radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ACOSH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_acosh_pd'.
// Requires AVX.
func M256AcoshPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AcoshPd([4]float64(a)))
}

func m256AcoshPd(a [4]float64) [4]float64


// M256AcoshPs: Compute the inverse hyperbolic cosine of packed
// single-precision (32-bit) floating-point elements in 'a' expressed in
// radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ACOSH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_acosh_ps'.
// Requires AVX.
func M256AcoshPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256AcoshPs([8]float32(a)))
}

func m256AcoshPs(a [8]float32) [8]float32


// M256AddPd: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm256_add_pd'.
// Requires AVX.
func M256AddPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AddPd([4]float64(a), [4]float64(b)))
}

func m256AddPd(a [4]float64, b [4]float64) [4]float64


// M256AddPs: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm256_add_ps'.
// Requires AVX.
func M256AddPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256AddPs([8]float32(a), [8]float32(b)))
}

func m256AddPs(a [8]float32, b [8]float32) [8]float32


// M256AddsubPd: Alternatively add and subtract packed double-precision
// (64-bit) floating-point elements in 'a' to/from packed elements in 'b', and
// store the results in 'dst'. 
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
func M256AddsubPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AddsubPd([4]float64(a), [4]float64(b)))
}

func m256AddsubPd(a [4]float64, b [4]float64) [4]float64


// M256AddsubPs: Alternatively add and subtract packed single-precision
// (32-bit) floating-point elements in 'a' to/from packed elements in 'b', and
// store the results in 'dst'. 
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
func M256AddsubPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256AddsubPs([8]float32(a), [8]float32(b)))
}

func m256AddsubPs(a [8]float32, b [8]float32) [8]float32


// M256AndPd: Compute the bitwise AND of packed double-precision (64-bit)
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
func M256AndPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AndPd([4]float64(a), [4]float64(b)))
}

func m256AndPd(a [4]float64, b [4]float64) [4]float64


// M256AndPs: Compute the bitwise AND of packed single-precision (32-bit)
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
func M256AndPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256AndPs([8]float32(a), [8]float32(b)))
}

func m256AndPs(a [8]float32, b [8]float32) [8]float32


// M256AndnotPd: Compute the bitwise AND NOT of packed double-precision
// (64-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPD'. Intrinsic: '_mm256_andnot_pd'.
// Requires AVX.
func M256AndnotPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AndnotPd([4]float64(a), [4]float64(b)))
}

func m256AndnotPd(a [4]float64, b [4]float64) [4]float64


// M256AndnotPs: Compute the bitwise AND NOT of packed single-precision
// (32-bit) floating-point elements in 'a' and 'b', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VANDNPS'. Intrinsic: '_mm256_andnot_ps'.
// Requires AVX.
func M256AndnotPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256AndnotPs([8]float32(a), [8]float32(b)))
}

func m256AndnotPs(a [8]float32, b [8]float32) [8]float32


// M256AsinPd: Compute the inverse sine of packed double-precision (64-bit)
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
func M256AsinPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AsinPd([4]float64(a)))
}

func m256AsinPd(a [4]float64) [4]float64


// M256AsinPs: Compute the inverse sine of packed single-precision (32-bit)
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
func M256AsinPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256AsinPs([8]float32(a)))
}

func m256AsinPs(a [8]float32) [8]float32


// M256AsinhPd: Compute the inverse hyperbolic sine of packed double-precision
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
func M256AsinhPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AsinhPd([4]float64(a)))
}

func m256AsinhPd(a [4]float64) [4]float64


// M256AsinhPs: Compute the inverse hyperbolic sine of packed single-precision
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
func M256AsinhPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256AsinhPs([8]float32(a)))
}

func m256AsinhPs(a [8]float32) [8]float32


// M256AtanPd: Compute the inverse tangent of packed double-precision (64-bit)
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
func M256AtanPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AtanPd([4]float64(a)))
}

func m256AtanPd(a [4]float64) [4]float64


// M256AtanPs: Compute the inverse tangent of packed single-precision (32-bit)
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
func M256AtanPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256AtanPs([8]float32(a)))
}

func m256AtanPs(a [8]float32) [8]float32


// M256Atan2Pd: Compute the inverse tangent of packed double-precision (64-bit)
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
func M256Atan2Pd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Atan2Pd([4]float64(a), [4]float64(b)))
}

func m256Atan2Pd(a [4]float64, b [4]float64) [4]float64


// M256Atan2Ps: Compute the inverse tangent of packed single-precision (32-bit)
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
func M256Atan2Ps(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256Atan2Ps([8]float32(a), [8]float32(b)))
}

func m256Atan2Ps(a [8]float32, b [8]float32) [8]float32


// M256AtanhPd: Compute the inverse hyperbolic tangent of packed
// double-precision (64-bit) floating-point elements in 'a' expressed in
// radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ATANH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atanh_pd'.
// Requires AVX.
func M256AtanhPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256AtanhPd([4]float64(a)))
}

func m256AtanhPd(a [4]float64) [4]float64


// M256AtanhPs: Compute the inverse hyperbolic tangent of packed
// single-precision (32-bit) floating-point elements in 'a' expressed in
// radians, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ATANH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_atanh_ps'.
// Requires AVX.
func M256AtanhPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256AtanhPs([8]float32(a)))
}

func m256AtanhPs(a [8]float32) [8]float32


// M256BlendPd: Blend packed double-precision (64-bit) floating-point elements
// from 'a' and 'b' using control mask 'imm8', and store the results in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256BlendPd(a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	return x86.M256d(m256BlendPd([4]float64(a), [4]float64(b), imm8))
}

func m256BlendPd(a [4]float64, b [4]float64, imm8 byte) [4]float64


// M256BlendPs: Blend packed single-precision (32-bit) floating-point elements
// from 'a' and 'b' using control mask 'imm8', and store the results in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256BlendPs(a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	return x86.M256(m256BlendPs([8]float32(a), [8]float32(b), imm8))
}

func m256BlendPs(a [8]float32, b [8]float32, imm8 byte) [8]float32


// M256BlendvPd: Blend packed double-precision (64-bit) floating-point elements
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
func M256BlendvPd(a x86.M256d, b x86.M256d, mask x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256BlendvPd([4]float64(a), [4]float64(b), [4]float64(mask)))
}

func m256BlendvPd(a [4]float64, b [4]float64, mask [4]float64) [4]float64


// M256BlendvPs: Blend packed single-precision (32-bit) floating-point elements
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
func M256BlendvPs(a x86.M256, b x86.M256, mask x86.M256) (dst x86.M256) {
	return x86.M256(m256BlendvPs([8]float32(a), [8]float32(b), [8]float32(mask)))
}

func m256BlendvPs(a [8]float32, b [8]float32, mask [8]float32) [8]float32


// M256BroadcastPd: Broadcast 128 bits from memory (composed of 2 packed
// double-precision (64-bit) floating-point elements) to all elements of 'dst'. 
//
//		tmp[127:0] = MEM[mem_addr+127:mem_addr]
//		dst[127:0] := tmp[127:0]
//		dst[255:128] := tmp[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF128'. Intrinsic: '_mm256_broadcast_pd'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256BroadcastPd(mem_addr *x86.M128dConst) (dst x86.M256d) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256d{}
}

// M256BroadcastPs: Broadcast 128 bits from memory (composed of 4 packed
// single-precision (32-bit) floating-point elements) to all elements of 'dst'. 
//
//		tmp[127:0] = MEM[mem_addr+127:mem_addr]
//		dst[127:0] := tmp[127:0]
//		dst[255:128] := tmp[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VBROADCASTF128'. Intrinsic: '_mm256_broadcast_ps'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256BroadcastPs(mem_addr *x86.M128Const) (dst x86.M256) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256{}
}

// Skipped: _mm256_broadcast_sd. Contains pointer parameter.


// Skipped: _mm_broadcast_ss. Contains pointer parameter.


// Skipped: _mm256_broadcast_ss. Contains pointer parameter.


// M256CastpdPs: Cast vector of type __m256d to type __m256.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd_ps'.
// Requires AVX.
func M256CastpdPs(a x86.M256d) (dst x86.M256) {
	return x86.M256(m256CastpdPs([4]float64(a)))
}

func m256CastpdPs(a [4]float64) [8]float32


// M256CastpdSi256: Casts vector of type __m256d to type __m256i. This
// intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd_si256'.
// Requires AVX.
func M256CastpdSi256(a x86.M256d) (dst x86.M256i) {
	return x86.M256i(m256CastpdSi256([4]float64(a)))
}

func m256CastpdSi256(a [4]float64) [32]byte


// M256Castpd128Pd256: Casts vector of type __m128d to type __m256d; the upper
// 128 bits of the result are undefined. This intrinsic is only used for
// compilation and does not generate any instructions, thus it has zero
// latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd128_pd256'.
// Requires AVX.
func M256Castpd128Pd256(a x86.M128d) (dst x86.M256d) {
	return x86.M256d(m256Castpd128Pd256([2]float64(a)))
}

func m256Castpd128Pd256(a [2]float64) [4]float64


// M256Castpd256Pd128: Casts vector of type __m256d to type __m128d. This
// intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castpd256_pd128'.
// Requires AVX.
func M256Castpd256Pd128(a x86.M256d) (dst x86.M128d) {
	return x86.M128d(m256Castpd256Pd128([4]float64(a)))
}

func m256Castpd256Pd128(a [4]float64) [2]float64


// M256CastpsPd: Cast vector of type __m256 to type __m256d.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps_pd'.
// Requires AVX.
func M256CastpsPd(a x86.M256) (dst x86.M256d) {
	return x86.M256d(m256CastpsPd([8]float32(a)))
}

func m256CastpsPd(a [8]float32) [4]float64


// M256CastpsSi256: Casts vector of type __m256 to type __m256i. This intrinsic
// is only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps_si256'.
// Requires AVX.
func M256CastpsSi256(a x86.M256) (dst x86.M256i) {
	return x86.M256i(m256CastpsSi256([8]float32(a)))
}

func m256CastpsSi256(a [8]float32) [32]byte


// M256Castps128Ps256: Casts vector of type __m128 to type __m256; the upper
// 128 bits of the result are undefined. This intrinsic is only used for
// compilation and does not generate any instructions, thus it has zero
// latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps128_ps256'.
// Requires AVX.
func M256Castps128Ps256(a x86.M128) (dst x86.M256) {
	return x86.M256(m256Castps128Ps256([4]float32(a)))
}

func m256Castps128Ps256(a [4]float32) [8]float32


// M256Castps256Ps128: Casts vector of type __m256 to type __m128. This
// intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castps256_ps128'.
// Requires AVX.
func M256Castps256Ps128(a x86.M256) (dst x86.M128) {
	return x86.M128(m256Castps256Ps128([8]float32(a)))
}

func m256Castps256Ps128(a [8]float32) [4]float32


// M256Castsi128Si256: Casts vector of type __m128i to type __m256i; the upper
// 128 bits of the result are undefined. This intrinsic is only used for
// compilation and does not generate any instructions, thus it has zero
// latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi128_si256'.
// Requires AVX.
func M256Castsi128Si256(a x86.M128i) (dst x86.M256i) {
	return x86.M256i(m256Castsi128Si256([16]byte(a)))
}

func m256Castsi128Si256(a [16]byte) [32]byte


// M256Castsi256Pd: Casts vector of type __m256i to type __m256d. This
// intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi256_pd'.
// Requires AVX.
func M256Castsi256Pd(a x86.M256i) (dst x86.M256d) {
	return x86.M256d(m256Castsi256Pd([32]byte(a)))
}

func m256Castsi256Pd(a [32]byte) [4]float64


// M256Castsi256Ps: Casts vector of type __m256i to type __m256. This intrinsic
// is only used for compilation and does not generate any instructions, thus it
// has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi256_ps'.
// Requires AVX.
func M256Castsi256Ps(a x86.M256i) (dst x86.M256) {
	return x86.M256(m256Castsi256Ps([32]byte(a)))
}

func m256Castsi256Ps(a [32]byte) [8]float32


// M256Castsi256Si128: Casts vector of type __m256i to type __m128i. This
// intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_castsi256_si128'.
// Requires AVX.
func M256Castsi256Si128(a x86.M256i) (dst x86.M128i) {
	return x86.M128i(m256Castsi256Si128([32]byte(a)))
}

func m256Castsi256Si128(a [32]byte) [16]byte


// M256CbrtPd: Compute the cube root of packed double-precision (64-bit)
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
func M256CbrtPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CbrtPd([4]float64(a)))
}

func m256CbrtPd(a [4]float64) [4]float64


// M256CbrtPs: Compute the cube root of packed single-precision (32-bit)
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
func M256CbrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CbrtPs([8]float32(a)))
}

func m256CbrtPs(a [8]float32) [8]float32


// M256CdfnormPd: Compute the cumulative distribution function of packed
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
func M256CdfnormPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CdfnormPd([4]float64(a)))
}

func m256CdfnormPd(a [4]float64) [4]float64


// M256CdfnormPs: Compute the cumulative distribution function of packed
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
func M256CdfnormPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CdfnormPs([8]float32(a)))
}

func m256CdfnormPs(a [8]float32) [8]float32


// M256CdfnorminvPd: Compute the inverse cumulative distribution function of
// packed double-precision (64-bit) floating-point elements in 'a' using the
// normal distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := InverseCDFNormal(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cdfnorminv_pd'.
// Requires AVX.
func M256CdfnorminvPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CdfnorminvPd([4]float64(a)))
}

func m256CdfnorminvPd(a [4]float64) [4]float64


// M256CdfnorminvPs: Compute the inverse cumulative distribution function of
// packed single-precision (32-bit) floating-point elements in 'a' using the
// normal distribution, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := InverseCDFNormal(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cdfnorminv_ps'.
// Requires AVX.
func M256CdfnorminvPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CdfnorminvPs([8]float32(a)))
}

func m256CdfnorminvPs(a [8]float32) [8]float32


// M256CeilPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' up to an integer value, and store the results as packed
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
func M256CeilPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CeilPd([4]float64(a)))
}

func m256CeilPd(a [4]float64) [4]float64


// M256CeilPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' up to an integer value, and store the results as packed
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
func M256CeilPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CeilPs([8]float32(a)))
}

func m256CeilPs(a [8]float32) [8]float32


// M256CexpPs: Compute the exponential value of 'e' raised to the power of
// packed complex single-precision (32-bit) floating-point elements in 'a', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cexp_ps'.
// Requires AVX.
func M256CexpPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CexpPs([8]float32(a)))
}

func m256CexpPs(a [8]float32) [8]float32


// M256ClogPs: Compute the natural logarithm of packed complex single-precision
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
func M256ClogPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256ClogPs([8]float32(a)))
}

func m256ClogPs(a [8]float32) [8]float32


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
//
// FIXME: Requires compiler support (has immediate)
func CmpPd(a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	return x86.M128d(cmpPd([2]float64(a), [2]float64(b), imm8))
}

func cmpPd(a [2]float64, b [2]float64, imm8 byte) [2]float64


// M256CmpPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b' based on the comparison operand specified by 'imm8', and
// store the results in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256CmpPd(a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	return x86.M256d(m256CmpPd([4]float64(a), [4]float64(b), imm8))
}

func m256CmpPd(a [4]float64, b [4]float64, imm8 byte) [4]float64


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
//
// FIXME: Requires compiler support (has immediate)
func CmpPs(a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	return x86.M128(cmpPs([4]float32(a), [4]float32(b), imm8))
}

func cmpPs(a [4]float32, b [4]float32, imm8 byte) [4]float32


// M256CmpPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b' based on the comparison operand specified by 'imm8', and
// store the results in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256CmpPs(a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	return x86.M256(m256CmpPs([8]float32(a), [8]float32(b), imm8))
}

func m256CmpPs(a [8]float32, b [8]float32, imm8 byte) [8]float32


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
//
// FIXME: Requires compiler support (has immediate)
func CmpSd(a x86.M128d, b x86.M128d, imm8 byte) (dst x86.M128d) {
	return x86.M128d(cmpSd([2]float64(a), [2]float64(b), imm8))
}

func cmpSd(a [2]float64, b [2]float64, imm8 byte) [2]float64


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
//
// FIXME: Requires compiler support (has immediate)
func CmpSs(a x86.M128, b x86.M128, imm8 byte) (dst x86.M128) {
	return x86.M128(cmpSs([4]float32(a), [4]float32(b), imm8))
}

func cmpSs(a [4]float32, b [4]float32, imm8 byte) [4]float32


// M256CosPd: Compute the cosine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := COS(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cos_pd'.
// Requires AVX.
func M256CosPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CosPd([4]float64(a)))
}

func m256CosPd(a [4]float64) [4]float64


// M256CosPs: Compute the cosine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := COS(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cos_ps'.
// Requires AVX.
func M256CosPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CosPs([8]float32(a)))
}

func m256CosPs(a [8]float32) [8]float32


// M256CosdPd: Compute the cosine of packed double-precision (64-bit)
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
func M256CosdPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CosdPd([4]float64(a)))
}

func m256CosdPd(a [4]float64) [4]float64


// M256CosdPs: Compute the cosine of packed single-precision (32-bit)
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
func M256CosdPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CosdPs([8]float32(a)))
}

func m256CosdPs(a [8]float32) [8]float32


// M256CoshPd: Compute the hyperbolic cosine of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := COSH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cosh_pd'.
// Requires AVX.
func M256CoshPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256CoshPd([4]float64(a)))
}

func m256CoshPd(a [4]float64) [4]float64


// M256CoshPs: Compute the hyperbolic cosine of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := COSH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_cosh_ps'.
// Requires AVX.
func M256CoshPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CoshPs([8]float32(a)))
}

func m256CoshPs(a [8]float32) [8]float32


// M256CsqrtPs: Compute the square root of packed complex single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SQRT(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_csqrt_ps'.
// Requires AVX.
func M256CsqrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256CsqrtPs([8]float32(a)))
}

func m256CsqrtPs(a [8]float32) [8]float32


// M256Cvtepi32Pd: Convert packed 32-bit integers in 'a' to packed
// double-precision (64-bit) floating-point elements, and store the results in
// 'dst'. 
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
func M256Cvtepi32Pd(a x86.M128i) (dst x86.M256d) {
	return x86.M256d(m256Cvtepi32Pd([16]byte(a)))
}

func m256Cvtepi32Pd(a [16]byte) [4]float64


// M256Cvtepi32Ps: Convert packed 32-bit integers in 'a' to packed
// single-precision (32-bit) floating-point elements, and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := Convert_Int32_To_FP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTDQ2PS'. Intrinsic: '_mm256_cvtepi32_ps'.
// Requires AVX.
func M256Cvtepi32Ps(a x86.M256i) (dst x86.M256) {
	return x86.M256(m256Cvtepi32Ps([32]byte(a)))
}

func m256Cvtepi32Ps(a [32]byte) [8]float32


// M256CvtpdEpi32: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed 32-bit integers, and store the results in 'dst'. 
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
func M256CvtpdEpi32(a x86.M256d) (dst x86.M128i) {
	return x86.M128i(m256CvtpdEpi32([4]float64(a)))
}

func m256CvtpdEpi32(a [4]float64) [16]byte


// M256CvtpdPs: Convert packed double-precision (64-bit) floating-point
// elements in 'a' to packed single-precision (32-bit) floating-point elements,
// and store the results in 'dst'. 
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
func M256CvtpdPs(a x86.M256d) (dst x86.M128) {
	return x86.M128(m256CvtpdPs([4]float64(a)))
}

func m256CvtpdPs(a [4]float64) [4]float32


// M256CvtpsEpi32: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := Convert_FP32_To_Int32(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VCVTPS2DQ'. Intrinsic: '_mm256_cvtps_epi32'.
// Requires AVX.
func M256CvtpsEpi32(a x86.M256) (dst x86.M256i) {
	return x86.M256i(m256CvtpsEpi32([8]float32(a)))
}

func m256CvtpsEpi32(a [8]float32) [32]byte


// M256CvtpsPd: Convert packed single-precision (32-bit) floating-point
// elements in 'a' to packed double-precision (64-bit) floating-point elements,
// and store the results in 'dst'. 
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
func M256CvtpsPd(a x86.M128) (dst x86.M256d) {
	return x86.M256d(m256CvtpsPd([4]float32(a)))
}

func m256CvtpsPd(a [4]float32) [4]float64


// M256CvttpdEpi32: Convert packed double-precision (64-bit) floating-point
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
func M256CvttpdEpi32(a x86.M256d) (dst x86.M128i) {
	return x86.M128i(m256CvttpdEpi32([4]float64(a)))
}

func m256CvttpdEpi32(a [4]float64) [16]byte


// M256CvttpsEpi32: Convert packed single-precision (32-bit) floating-point
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
func M256CvttpsEpi32(a x86.M256) (dst x86.M256i) {
	return x86.M256i(m256CvttpsEpi32([8]float32(a)))
}

func m256CvttpsEpi32(a [8]float32) [32]byte


// M256DivEpi16: Divide packed 16-bit integers in 'a' by packed elements in
// 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := TRUNCATE(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi16'.
// Requires AVX.
func M256DivEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpi16([32]byte(a), [32]byte(b)))
}

func m256DivEpi16(a [32]byte, b [32]byte) [32]byte


// M256DivEpi32: Divide packed 32-bit integers in 'a' by packed elements in
// 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi32'.
// Requires AVX.
func M256DivEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpi32([32]byte(a), [32]byte(b)))
}

func m256DivEpi32(a [32]byte, b [32]byte) [32]byte


// M256DivEpi64: Divide packed 64-bit integers in 'a' by packed elements in
// 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := TRUNCATE(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi64'.
// Requires AVX.
func M256DivEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpi64([32]byte(a), [32]byte(b)))
}

func m256DivEpi64(a [32]byte, b [32]byte) [32]byte


// M256DivEpi8: Divide packed 8-bit integers in 'a' by packed elements in 'b',
// and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := TRUNCATE(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epi8'.
// Requires AVX.
func M256DivEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpi8([32]byte(a), [32]byte(b)))
}

func m256DivEpi8(a [32]byte, b [32]byte) [32]byte


// M256DivEpu16: Divide packed unsigned 16-bit integers in 'a' by packed
// elements in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := TRUNCATE(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu16'.
// Requires AVX.
func M256DivEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpu16([32]byte(a), [32]byte(b)))
}

func m256DivEpu16(a [32]byte, b [32]byte) [32]byte


// M256DivEpu32: Divide packed unsigned 32-bit integers in 'a' by packed
// elements in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu32'.
// Requires AVX.
func M256DivEpu32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpu32([32]byte(a), [32]byte(b)))
}

func m256DivEpu32(a [32]byte, b [32]byte) [32]byte


// M256DivEpu64: Divide packed unsigned 64-bit integers in 'a' by packed
// elements in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := TRUNCATE(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu64'.
// Requires AVX.
func M256DivEpu64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpu64([32]byte(a), [32]byte(b)))
}

func m256DivEpu64(a [32]byte, b [32]byte) [32]byte


// M256DivEpu8: Divide packed unsigned 8-bit integers in 'a' by packed elements
// in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := TRUNCATE(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_div_epu8'.
// Requires AVX.
func M256DivEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256DivEpu8([32]byte(a), [32]byte(b)))
}

func m256DivEpu8(a [32]byte, b [32]byte) [32]byte


// M256DivPd: Divide packed double-precision (64-bit) floating-point elements
// in 'a' by packed elements in 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := a[i+63:i] / b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDIVPD'. Intrinsic: '_mm256_div_pd'.
// Requires AVX.
func M256DivPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256DivPd([4]float64(a), [4]float64(b)))
}

func m256DivPd(a [4]float64, b [4]float64) [4]float64


// M256DivPs: Divide packed single-precision (32-bit) floating-point elements
// in 'a' by packed elements in 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := a[i+31:i] / b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VDIVPS'. Intrinsic: '_mm256_div_ps'.
// Requires AVX.
func M256DivPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256DivPs([8]float32(a), [8]float32(b)))
}

func m256DivPs(a [8]float32, b [8]float32) [8]float32


// M256DpPs: Conditionally multiply the packed single-precision (32-bit)
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
//
// FIXME: Requires compiler support (has immediate)
func M256DpPs(a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	return x86.M256(m256DpPs([8]float32(a), [8]float32(b), imm8))
}

func m256DpPs(a [8]float32, b [8]float32, imm8 byte) [8]float32


// M256ErfPd: Compute the error function of packed double-precision (64-bit)
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
func M256ErfPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256ErfPd([4]float64(a)))
}

func m256ErfPd(a [4]float64) [4]float64


// M256ErfPs: Compute the error function of packed single-precision (32-bit)
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
func M256ErfPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256ErfPs([8]float32(a)))
}

func m256ErfPs(a [8]float32) [8]float32


// M256ErfcPd: Compute the complementary error function of packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 1.0 - ERF(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfc_pd'.
// Requires AVX.
func M256ErfcPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256ErfcPd([4]float64(a)))
}

func m256ErfcPd(a [4]float64) [4]float64


// M256ErfcPs: Compute the complementary error function of packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 1.0 - ERF(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_erfc_ps'.
// Requires AVX.
func M256ErfcPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256ErfcPs([8]float32(a)))
}

func m256ErfcPs(a [8]float32) [8]float32


// M256ErfcinvPd: Compute the inverse complementary error function of packed
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
func M256ErfcinvPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256ErfcinvPd([4]float64(a)))
}

func m256ErfcinvPd(a [4]float64) [4]float64


// M256ErfcinvPs: Compute the inverse complementary error function of packed
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
func M256ErfcinvPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256ErfcinvPs([8]float32(a)))
}

func m256ErfcinvPs(a [8]float32) [8]float32


// M256ErfinvPd: Compute the inverse error function of packed double-precision
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
func M256ErfinvPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256ErfinvPd([4]float64(a)))
}

func m256ErfinvPd(a [4]float64) [4]float64


// M256ErfinvPs: Compute the inverse error function of packed single-precision
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
func M256ErfinvPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256ErfinvPs([8]float32(a)))
}

func m256ErfinvPs(a [8]float32) [8]float32


// M256ExpPd: Compute the exponential value of 'e' raised to the power of
// packed double-precision (64-bit) floating-point elements in 'a', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := e^(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp_pd'.
// Requires AVX.
func M256ExpPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256ExpPd([4]float64(a)))
}

func m256ExpPd(a [4]float64) [4]float64


// M256ExpPs: Compute the exponential value of 'e' raised to the power of
// packed single-precision (32-bit) floating-point elements in 'a', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp_ps'.
// Requires AVX.
func M256ExpPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256ExpPs([8]float32(a)))
}

func m256ExpPs(a [8]float32) [8]float32


// M256Exp10Pd: Compute the exponential value of 10 raised to the power of
// packed double-precision (64-bit) floating-point elements in 'a', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := 10^(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp10_pd'.
// Requires AVX.
func M256Exp10Pd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Exp10Pd([4]float64(a)))
}

func m256Exp10Pd(a [4]float64) [4]float64


// M256Exp10Ps: Compute the exponential value of 10 raised to the power of
// packed single-precision (32-bit) floating-point elements in 'a', and store
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := 10^(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_exp10_ps'.
// Requires AVX.
func M256Exp10Ps(a x86.M256) (dst x86.M256) {
	return x86.M256(m256Exp10Ps([8]float32(a)))
}

func m256Exp10Ps(a [8]float32) [8]float32


// M256Exp2Pd: Compute the exponential value of 2 raised to the power of packed
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
func M256Exp2Pd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Exp2Pd([4]float64(a)))
}

func m256Exp2Pd(a [4]float64) [4]float64


// M256Exp2Ps: Compute the exponential value of 2 raised to the power of packed
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
func M256Exp2Ps(a x86.M256) (dst x86.M256) {
	return x86.M256(m256Exp2Ps([8]float32(a)))
}

func m256Exp2Ps(a [8]float32) [8]float32


// M256Expm1Pd: Compute the exponential value of 'e' raised to the power of
// packed double-precision (64-bit) floating-point elements in 'a', subtract
// one from each element, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := e^(a[i+63:i]) - 1.0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_expm1_pd'.
// Requires AVX.
func M256Expm1Pd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Expm1Pd([4]float64(a)))
}

func m256Expm1Pd(a [4]float64) [4]float64


// M256Expm1Ps: Compute the exponential value of 'e' raised to the power of
// packed single-precision (32-bit) floating-point elements in 'a', subtract
// one from each element, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := e^(a[i+31:i]) - 1.0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_expm1_ps'.
// Requires AVX.
func M256Expm1Ps(a x86.M256) (dst x86.M256) {
	return x86.M256(m256Expm1Ps([8]float32(a)))
}

func m256Expm1Ps(a [8]float32) [8]float32


// M256ExtractEpi16: Extract a 16-bit integer from 'a', selected with 'index',
// and store the result in 'dst'. 
//
//		dst[15:0] := (a[255:0] >> (index * 16))[15:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi16'.
// Requires AVX.
func M256ExtractEpi16(a x86.M256i, index int) int16 {
	return int16(m256ExtractEpi16([32]byte(a), index))
}

func m256ExtractEpi16(a [32]byte, index int) int16


// M256ExtractEpi32: Extract a 32-bit integer from 'a', selected with 'index',
// and store the result in 'dst'. 
//
//		dst[31:0] := (a[255:0] >> (index * 32))[31:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi32'.
// Requires AVX.
func M256ExtractEpi32(a x86.M256i, index int) int32 {
	return int32(m256ExtractEpi32([32]byte(a), index))
}

func m256ExtractEpi32(a [32]byte, index int) int32


// M256ExtractEpi64: Extract a 64-bit integer from 'a', selected with 'index',
// and store the result in 'dst'. 
//
//		dst[63:0] := (a[255:0] >> (index * 64))[63:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi64'.
// Requires AVX.
func M256ExtractEpi64(a x86.M256i, index int) int64 {
	return int64(m256ExtractEpi64([32]byte(a), index))
}

func m256ExtractEpi64(a [32]byte, index int) int64


// M256ExtractEpi8: Extract an 8-bit integer from 'a', selected with 'index',
// and store the result in 'dst'. 
//
//		dst[7:0] := (a[255:0] >> (index * 8))[7:0]
//
// Instruction: '...'. Intrinsic: '_mm256_extract_epi8'.
// Requires AVX.
func M256ExtractEpi8(a x86.M256i, index int) int8 {
	return int8(m256ExtractEpi8([32]byte(a), index))
}

func m256ExtractEpi8(a [32]byte, index int) int8


// M256Extractf128Pd: Extract 128 bits (composed of 2 packed double-precision
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
//
// FIXME: Requires compiler support (has immediate)
func M256Extractf128Pd(a x86.M256d, imm8 byte) (dst x86.M128d) {
	return x86.M128d(m256Extractf128Pd([4]float64(a), imm8))
}

func m256Extractf128Pd(a [4]float64, imm8 byte) [2]float64


// M256Extractf128Ps: Extract 128 bits (composed of 4 packed single-precision
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
//
// FIXME: Requires compiler support (has immediate)
func M256Extractf128Ps(a x86.M256, imm8 byte) (dst x86.M128) {
	return x86.M128(m256Extractf128Ps([8]float32(a), imm8))
}

func m256Extractf128Ps(a [8]float32, imm8 byte) [4]float32


// M256Extractf128Si256: Extract 128 bits (composed of integer data) from 'a',
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
//
// FIXME: Requires compiler support (has immediate)
func M256Extractf128Si256(a x86.M256i, imm8 byte) (dst x86.M128i) {
	return x86.M128i(m256Extractf128Si256([32]byte(a), imm8))
}

func m256Extractf128Si256(a [32]byte, imm8 byte) [16]byte


// M256FloorPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' down to an integer value, and store the results as packed
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
func M256FloorPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256FloorPd([4]float64(a)))
}

func m256FloorPd(a [4]float64) [4]float64


// M256FloorPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' down to an integer value, and store the results as packed
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
func M256FloorPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256FloorPs([8]float32(a)))
}

func m256FloorPs(a [8]float32) [8]float32


// M256HaddPd: Horizontally add adjacent pairs of double-precision (64-bit)
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
func M256HaddPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256HaddPd([4]float64(a), [4]float64(b)))
}

func m256HaddPd(a [4]float64, b [4]float64) [4]float64


// M256HaddPs: Horizontally add adjacent pairs of single-precision (32-bit)
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
func M256HaddPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256HaddPs([8]float32(a), [8]float32(b)))
}

func m256HaddPs(a [8]float32, b [8]float32) [8]float32


// M256HsubPd: Horizontally subtract adjacent pairs of double-precision
// (64-bit) floating-point elements in 'a' and 'b', and pack the results in
// 'dst'. 
//
//		dst[63:0] := a[63:0] - a[127:64]
//		dst[127:64] := b[63:0] - b[127:64]
//		dst[191:128] := a[191:128] - a[255:192]
//		dst[255:192] := b[191:128] - b[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VHSUBPD'. Intrinsic: '_mm256_hsub_pd'.
// Requires AVX.
func M256HsubPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256HsubPd([4]float64(a), [4]float64(b)))
}

func m256HsubPd(a [4]float64, b [4]float64) [4]float64


// M256HsubPs: Horizontally add adjacent pairs of single-precision (32-bit)
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
func M256HsubPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256HsubPs([8]float32(a), [8]float32(b)))
}

func m256HsubPs(a [8]float32, b [8]float32) [8]float32


// M256HypotPd: Compute the length of the hypotenous of a right triangle, with
// the lengths of the other two sides of the triangle stored as packed
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
func M256HypotPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256HypotPd([4]float64(a), [4]float64(b)))
}

func m256HypotPd(a [4]float64, b [4]float64) [4]float64


// M256HypotPs: Compute the length of the hypotenous of a right triangle, with
// the lengths of the other two sides of the triangle stored as packed
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
func M256HypotPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256HypotPs([8]float32(a), [8]float32(b)))
}

func m256HypotPs(a [8]float32, b [8]float32) [8]float32


// M256IdivEpi32: Divide packed 32-bit integers in 'a' by packed elements in
// 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_idiv_epi32'.
// Requires AVX.
func M256IdivEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256IdivEpi32([32]byte(a), [32]byte(b)))
}

func m256IdivEpi32(a [32]byte, b [32]byte) [32]byte


// M256IdivremEpi32: Divide packed 32-bit integers in 'a' by packed elements in
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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256IdivremEpi32(mem_addr *x86.M256i, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256i{}
}

// M256InsertEpi16: Copy 'a' to 'dst', and insert the 16-bit integer 'i' into
// 'dst' at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*16
//		dst[sel+15:sel] := i[15:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi16'.
// Requires AVX.
func M256InsertEpi16(a x86.M256i, i int16, index int) (dst x86.M256i) {
	return x86.M256i(m256InsertEpi16([32]byte(a), i, index))
}

func m256InsertEpi16(a [32]byte, i int16, index int) [32]byte


// M256InsertEpi32: Copy 'a' to 'dst', and insert the 32-bit integer 'i' into
// 'dst' at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*32
//		dst[sel+31:sel] := i[31:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi32'.
// Requires AVX.
func M256InsertEpi32(a x86.M256i, i int32, index int) (dst x86.M256i) {
	return x86.M256i(m256InsertEpi32([32]byte(a), i, index))
}

func m256InsertEpi32(a [32]byte, i int32, index int) [32]byte


// M256InsertEpi64: Copy 'a' to 'dst', and insert the 64-bit integer 'i' into
// 'dst' at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*64
//		dst[sel+63:sel] := i[63:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi64'.
// Requires AVX.
func M256InsertEpi64(a x86.M256i, i int64, index int) (dst x86.M256i) {
	return x86.M256i(m256InsertEpi64([32]byte(a), i, index))
}

func m256InsertEpi64(a [32]byte, i int64, index int) [32]byte


// M256InsertEpi8: Copy 'a' to 'dst', and insert the 8-bit integer 'i' into
// 'dst' at the location specified by 'index'. 
//
//		dst[255:0] := a[255:0]
//		sel := index*8
//		dst[sel+7:sel] := i[7:0]
//
// Instruction: '...'. Intrinsic: '_mm256_insert_epi8'.
// Requires AVX.
func M256InsertEpi8(a x86.M256i, i int8, index int) (dst x86.M256i) {
	return x86.M256i(m256InsertEpi8([32]byte(a), i, index))
}

func m256InsertEpi8(a [32]byte, i int8, index int) [32]byte


// M256Insertf128Pd: Copy 'a' to 'dst', then insert 128 bits (composed of 2
// packed double-precision (64-bit) floating-point elements) from 'b' into
// 'dst' at the location specified by 'imm8'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256Insertf128Pd(a x86.M256d, b x86.M128d, imm8 byte) (dst x86.M256d) {
	return x86.M256d(m256Insertf128Pd([4]float64(a), [2]float64(b), imm8))
}

func m256Insertf128Pd(a [4]float64, b [2]float64, imm8 byte) [4]float64


// M256Insertf128Ps: Copy 'a' to 'dst', then insert 128 bits (composed of 4
// packed single-precision (32-bit) floating-point elements) from 'b' into
// 'dst' at the location specified by 'imm8'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256Insertf128Ps(a x86.M256, b x86.M128, imm8 byte) (dst x86.M256) {
	return x86.M256(m256Insertf128Ps([8]float32(a), [4]float32(b), imm8))
}

func m256Insertf128Ps(a [8]float32, b [4]float32, imm8 byte) [8]float32


// M256Insertf128Si256: Copy 'a' to 'dst', then insert 128 bits from 'b' into
// 'dst' at the location specified by 'imm8'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256Insertf128Si256(a x86.M256i, b x86.M128i, imm8 byte) (dst x86.M256i) {
	return x86.M256i(m256Insertf128Si256([32]byte(a), [16]byte(b), imm8))
}

func m256Insertf128Si256(a [32]byte, b [16]byte, imm8 byte) [32]byte


// M256InvcbrtPd: Compute the inverse cube root of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := InvCubeRoot(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_invcbrt_pd'.
// Requires AVX.
func M256InvcbrtPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256InvcbrtPd([4]float64(a)))
}

func m256InvcbrtPd(a [4]float64) [4]float64


// M256InvcbrtPs: Compute the inverse cube root of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := InvCubeRoot(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_invcbrt_ps'.
// Requires AVX.
func M256InvcbrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256InvcbrtPs([8]float32(a)))
}

func m256InvcbrtPs(a [8]float32) [8]float32


// M256InvsqrtPd: Compute the inverse square root of packed double-precision
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
func M256InvsqrtPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256InvsqrtPd([4]float64(a)))
}

func m256InvsqrtPd(a [4]float64) [4]float64


// M256InvsqrtPs: Compute the inverse square root of packed single-precision
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
func M256InvsqrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256InvsqrtPs([8]float32(a)))
}

func m256InvsqrtPs(a [8]float32) [8]float32


// M256IremEpi32: Divide packed 32-bit integers in 'a' by packed elements in
// 'b', and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_irem_epi32'.
// Requires AVX.
func M256IremEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256IremEpi32([32]byte(a), [32]byte(b)))
}

func m256IremEpi32(a [32]byte, b [32]byte) [32]byte


// M256LddquSi256: Load 256-bits of integer data from unaligned memory into
// 'dst'. This intrinsic may perform better than '_mm256_loadu_si256' when the
// data crosses a cache line boundary. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VLDDQU'. Intrinsic: '_mm256_lddqu_si256'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256LddquSi256(mem_addr *x86.M256iConst) (dst x86.M256i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256i{}
}

// Skipped: _mm256_load_pd. Contains pointer parameter.


// Skipped: _mm256_load_ps. Contains pointer parameter.


// M256LoadSi256: Load 256-bits of integer data from memory into 'dst'.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQA'. Intrinsic: '_mm256_load_si256'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256LoadSi256(mem_addr *x86.M256iConst) (dst x86.M256i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256i{}
}

// Skipped: _mm256_loadu_pd. Contains pointer parameter.


// Skipped: _mm256_loadu_ps. Contains pointer parameter.


// M256LoaduSi256: Load 256-bits of integer data from memory into 'dst'.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		dst[255:0] := MEM[mem_addr+255:mem_addr]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDQU'. Intrinsic: '_mm256_loadu_si256'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256LoaduSi256(mem_addr *x86.M256iConst) (dst x86.M256i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256i{}
}

// Skipped: _mm256_loadu2_m128. Contains pointer parameter.


// Skipped: _mm256_loadu2_m128d. Contains pointer parameter.


// M256Loadu2M128i: Load two 128-bit values (composed of integer data) from
// memory, and combine them into a 256-bit value in 'dst'.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		dst[127:0] := MEM[loaddr+127:loaddr]
//		dst[255:128] := MEM[hiaddr+127:hiaddr]
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_loadu2_m128i'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256Loadu2M128i(hiaddr *x86.M128iConst, loaddr *x86.M128iConst) (dst x86.M256i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256i{}
}

// M256LogPd: Compute the natural logarithm of packed double-precision (64-bit)
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
func M256LogPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256LogPd([4]float64(a)))
}

func m256LogPd(a [4]float64) [4]float64


// M256LogPs: Compute the natural logarithm of packed single-precision (32-bit)
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
func M256LogPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256LogPs([8]float32(a)))
}

func m256LogPs(a [8]float32) [8]float32


// M256Log10Pd: Compute the base-10 logarithm of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := log10(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log10_pd'.
// Requires AVX.
func M256Log10Pd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Log10Pd([4]float64(a)))
}

func m256Log10Pd(a [4]float64) [4]float64


// M256Log10Ps: Compute the base-10 logarithm of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := log10(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log10_ps'.
// Requires AVX.
func M256Log10Ps(a x86.M256) (dst x86.M256) {
	return x86.M256(m256Log10Ps([8]float32(a)))
}

func m256Log10Ps(a [8]float32) [8]float32


// M256Log1pPd: Compute the natural logarithm of one plus packed
// double-precision (64-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ln(1.0 + a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log1p_pd'.
// Requires AVX.
func M256Log1pPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Log1pPd([4]float64(a)))
}

func m256Log1pPd(a [4]float64) [4]float64


// M256Log1pPs: Compute the natural logarithm of one plus packed
// single-precision (32-bit) floating-point elements in 'a', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ln(1.0 + a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_log1p_ps'.
// Requires AVX.
func M256Log1pPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256Log1pPs([8]float32(a)))
}

func m256Log1pPs(a [8]float32) [8]float32


// M256Log2Pd: Compute the base-2 logarithm of packed double-precision (64-bit)
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
func M256Log2Pd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256Log2Pd([4]float64(a)))
}

func m256Log2Pd(a [4]float64) [4]float64


// M256Log2Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
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
func M256Log2Ps(a x86.M256) (dst x86.M256) {
	return x86.M256(m256Log2Ps([8]float32(a)))
}

func m256Log2Ps(a [8]float32) [8]float32


// M256LogbPd: Convert the exponent of each packed double-precision (64-bit)
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
func M256LogbPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256LogbPd([4]float64(a)))
}

func m256LogbPd(a [4]float64) [4]float64


// M256LogbPs: Convert the exponent of each packed single-precision (32-bit)
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
func M256LogbPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256LogbPs([8]float32(a)))
}

func m256LogbPs(a [8]float32) [8]float32


// Skipped: _mm_maskload_pd. Contains pointer parameter.


// Skipped: _mm256_maskload_pd. Contains pointer parameter.


// Skipped: _mm_maskload_ps. Contains pointer parameter.


// Skipped: _mm256_maskload_ps. Contains pointer parameter.


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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskstorePd(mem_addr *float64, mask x86.M128i, a x86.M128d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256MaskstorePd: Store packed double-precision (64-bit) floating-point
// elements from 'a' into memory using 'mask'. 
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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskstorePd(mem_addr *float64, mask x86.M256i, a x86.M256d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func MaskstorePs(mem_addr *float32, mask x86.M128i, a x86.M128)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256MaskstorePs: Store packed single-precision (32-bit) floating-point
// elements from 'a' into memory using 'mask'. 
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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256MaskstorePs(mem_addr *float32, mask x86.M256i, a x86.M256)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256MaxPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := MAX(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMAXPD'. Intrinsic: '_mm256_max_pd'.
// Requires AVX.
func M256MaxPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256MaxPd([4]float64(a), [4]float64(b)))
}

func m256MaxPd(a [4]float64, b [4]float64) [4]float64


// M256MaxPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := MAX(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMAXPS'. Intrinsic: '_mm256_max_ps'.
// Requires AVX.
func M256MaxPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256MaxPs([8]float32(a), [8]float32(b)))
}

func m256MaxPs(a [8]float32, b [8]float32) [8]float32


// M256MinPd: Compare packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := MIN(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMINPD'. Intrinsic: '_mm256_min_pd'.
// Requires AVX.
func M256MinPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256MinPd([4]float64(a), [4]float64(b)))
}

func m256MinPd(a [4]float64, b [4]float64) [4]float64


// M256MinPs: Compare packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := MIN(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMINPS'. Intrinsic: '_mm256_min_ps'.
// Requires AVX.
func M256MinPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256MinPs([8]float32(a), [8]float32(b)))
}

func m256MinPs(a [8]float32, b [8]float32) [8]float32


// M256MovedupPd: Duplicate even-indexed double-precision (64-bit)
// floating-point elements from 'a', and store the results in 'dst'. 
//
//		dst[63:0] := a[63:0]
//		dst[127:64] := a[63:0]
//		dst[191:128] := a[191:128]
//		dst[255:192] := a[191:128]
//		dst[MAX:256] := 0
//
// Instruction: 'VMOVDDUP'. Intrinsic: '_mm256_movedup_pd'.
// Requires AVX.
func M256MovedupPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256MovedupPd([4]float64(a)))
}

func m256MovedupPd(a [4]float64) [4]float64


// M256MovehdupPs: Duplicate odd-indexed single-precision (32-bit)
// floating-point elements from 'a', and store the results in 'dst'. 
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
func M256MovehdupPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256MovehdupPs([8]float32(a)))
}

func m256MovehdupPs(a [8]float32) [8]float32


// M256MoveldupPs: Duplicate even-indexed single-precision (32-bit)
// floating-point elements from 'a', and store the results in 'dst'. 
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
func M256MoveldupPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256MoveldupPs([8]float32(a)))
}

func m256MoveldupPs(a [8]float32) [8]float32


// M256MovemaskPd: Set each bit of mask 'dst' based on the most significant bit
// of the corresponding packed double-precision (64-bit) floating-point element
// in 'a'. 
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
func M256MovemaskPd(a x86.M256d) int {
	return int(m256MovemaskPd([4]float64(a)))
}

func m256MovemaskPd(a [4]float64) int


// M256MovemaskPs: Set each bit of mask 'dst' based on the most significant bit
// of the corresponding packed single-precision (32-bit) floating-point element
// in 'a'. 
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
func M256MovemaskPs(a x86.M256) int {
	return int(m256MovemaskPs([8]float32(a)))
}

func m256MovemaskPs(a [8]float32) int


// M256MulPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm256_mul_pd'.
// Requires AVX.
func M256MulPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256MulPd([4]float64(a), [4]float64(b)))
}

func m256MulPd(a [4]float64, b [4]float64) [4]float64


// M256MulPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm256_mul_ps'.
// Requires AVX.
func M256MulPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256MulPs([8]float32(a), [8]float32(b)))
}

func m256MulPs(a [8]float32, b [8]float32) [8]float32


// M256OrPd: Compute the bitwise OR of packed double-precision (64-bit)
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
func M256OrPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256OrPd([4]float64(a), [4]float64(b)))
}

func m256OrPd(a [4]float64, b [4]float64) [4]float64


// M256OrPs: Compute the bitwise OR of packed single-precision (32-bit)
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
func M256OrPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256OrPs([8]float32(a), [8]float32(b)))
}

func m256OrPs(a [8]float32, b [8]float32) [8]float32


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
//
// FIXME: Requires compiler support (has immediate)
func PermutePd(a x86.M128d, imm8 byte) (dst x86.M128d) {
	return x86.M128d(permutePd([2]float64(a), imm8))
}

func permutePd(a [2]float64, imm8 byte) [2]float64


// M256PermutePd: Shuffle double-precision (64-bit) floating-point elements in
// 'a' within 128-bit lanes using the control in 'imm8', and store the results
// in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256PermutePd(a x86.M256d, imm8 byte) (dst x86.M256d) {
	return x86.M256d(m256PermutePd([4]float64(a), imm8))
}

func m256PermutePd(a [4]float64, imm8 byte) [4]float64


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
//
// FIXME: Requires compiler support (has immediate)
func PermutePs(a x86.M128, imm8 byte) (dst x86.M128) {
	return x86.M128(permutePs([4]float32(a), imm8))
}

func permutePs(a [4]float32, imm8 byte) [4]float32


// M256PermutePs: Shuffle single-precision (32-bit) floating-point elements in
// 'a' within 128-bit lanes using the control in 'imm8', and store the results
// in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256PermutePs(a x86.M256, imm8 byte) (dst x86.M256) {
	return x86.M256(m256PermutePs([8]float32(a), imm8))
}

func m256PermutePs(a [8]float32, imm8 byte) [8]float32


// M256Permute2f128Pd: Shuffle 128-bits (composed of 2 packed double-precision
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
//
// FIXME: Requires compiler support (has immediate)
func M256Permute2f128Pd(a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	return x86.M256d(m256Permute2f128Pd([4]float64(a), [4]float64(b), imm8))
}

func m256Permute2f128Pd(a [4]float64, b [4]float64, imm8 byte) [4]float64


// M256Permute2f128Ps: Shuffle 128-bits (composed of 4 packed single-precision
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
//
// FIXME: Requires compiler support (has immediate)
func M256Permute2f128Ps(a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	return x86.M256(m256Permute2f128Ps([8]float32(a), [8]float32(b), imm8))
}

func m256Permute2f128Ps(a [8]float32, b [8]float32, imm8 byte) [8]float32


// M256Permute2f128Si256: Shuffle 128-bits (composed of integer data) selected
// by 'imm8' from 'a' and 'b', and store the results in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256Permute2f128Si256(a x86.M256i, b x86.M256i, imm8 byte) (dst x86.M256i) {
	return x86.M256i(m256Permute2f128Si256([32]byte(a), [32]byte(b), imm8))
}

func m256Permute2f128Si256(a [32]byte, b [32]byte, imm8 byte) [32]byte


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
func PermutevarPd(a x86.M128d, b x86.M128i) (dst x86.M128d) {
	return x86.M128d(permutevarPd([2]float64(a), [16]byte(b)))
}

func permutevarPd(a [2]float64, b [16]byte) [2]float64


// M256PermutevarPd: Shuffle double-precision (64-bit) floating-point elements
// in 'a' within 128-bit lanes using the control in 'b', and store the results
// in 'dst'. 
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
func M256PermutevarPd(a x86.M256d, b x86.M256i) (dst x86.M256d) {
	return x86.M256d(m256PermutevarPd([4]float64(a), [32]byte(b)))
}

func m256PermutevarPd(a [4]float64, b [32]byte) [4]float64


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
func PermutevarPs(a x86.M128, b x86.M128i) (dst x86.M128) {
	return x86.M128(permutevarPs([4]float32(a), [16]byte(b)))
}

func permutevarPs(a [4]float32, b [16]byte) [4]float32


// M256PermutevarPs: Shuffle single-precision (32-bit) floating-point elements
// in 'a' within 128-bit lanes using the control in 'b', and store the results
// in 'dst'. 
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
func M256PermutevarPs(a x86.M256, b x86.M256i) (dst x86.M256) {
	return x86.M256(m256PermutevarPs([8]float32(a), [32]byte(b)))
}

func m256PermutevarPs(a [8]float32, b [32]byte) [8]float32


// M256PowPd: Compute the exponential value of packed double-precision (64-bit)
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
func M256PowPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256PowPd([4]float64(a), [4]float64(b)))
}

func m256PowPd(a [4]float64, b [4]float64) [4]float64


// M256PowPs: Compute the exponential value of packed single-precision (32-bit)
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
func M256PowPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256PowPs([8]float32(a), [8]float32(b)))
}

func m256PowPs(a [8]float32, b [8]float32) [8]float32


// M256RcpPs: Compute the approximate reciprocal of packed single-precision
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
func M256RcpPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256RcpPs([8]float32(a)))
}

func m256RcpPs(a [8]float32) [8]float32


// M256RemEpi16: Divide packed 16-bit integers in 'a' by packed elements in
// 'b', and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := REMAINDER(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi16'.
// Requires AVX.
func M256RemEpi16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpi16([32]byte(a), [32]byte(b)))
}

func m256RemEpi16(a [32]byte, b [32]byte) [32]byte


// M256RemEpi32: Divide packed 32-bit integers in 'a' by packed elements in
// 'b', and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi32'.
// Requires AVX.
func M256RemEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpi32([32]byte(a), [32]byte(b)))
}

func m256RemEpi32(a [32]byte, b [32]byte) [32]byte


// M256RemEpi64: Divide packed 64-bit integers in 'a' by packed elements in
// 'b', and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := REMAINDER(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi64'.
// Requires AVX.
func M256RemEpi64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpi64([32]byte(a), [32]byte(b)))
}

func m256RemEpi64(a [32]byte, b [32]byte) [32]byte


// M256RemEpi8: Divide packed 8-bit integers in 'a' by packed elements in 'b',
// and store the remainders as packed 32-bit integers in 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := REMAINDER(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epi8'.
// Requires AVX.
func M256RemEpi8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpi8([32]byte(a), [32]byte(b)))
}

func m256RemEpi8(a [32]byte, b [32]byte) [32]byte


// M256RemEpu16: Divide packed unsigned 16-bit integers in 'a' by packed
// elements in 'b', and store the remainders as packed unsigned 32-bit integers
// in 'dst'. 
//
//		FOR j := 0 to 15
//			i := 16*j
//			dst[i+15:i] := REMAINDER(a[i+15:i] / b[i+15:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu16'.
// Requires AVX.
func M256RemEpu16(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpu16([32]byte(a), [32]byte(b)))
}

func m256RemEpu16(a [32]byte, b [32]byte) [32]byte


// M256RemEpu32: Divide packed unsigned 32-bit integers in 'a' by packed
// elements in 'b', and store the remainders as packed unsigned 32-bit integers
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu32'.
// Requires AVX.
func M256RemEpu32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpu32([32]byte(a), [32]byte(b)))
}

func m256RemEpu32(a [32]byte, b [32]byte) [32]byte


// M256RemEpu64: Divide packed unsigned 64-bit integers in 'a' by packed
// elements in 'b', and store the remainders as packed unsigned 32-bit integers
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 64*j
//			dst[i+63:i] := REMAINDER(a[i+63:i] / b[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu64'.
// Requires AVX.
func M256RemEpu64(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpu64([32]byte(a), [32]byte(b)))
}

func m256RemEpu64(a [32]byte, b [32]byte) [32]byte


// M256RemEpu8: Divide packed unsigned 8-bit integers in 'a' by packed elements
// in 'b', and store the remainders as packed unsigned 32-bit integers in
// 'dst'. 
//
//		FOR j := 0 to 31
//			i := 8*j
//			dst[i+7:i] := REMAINDER(a[i+7:i] / b[i+7:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_rem_epu8'.
// Requires AVX.
func M256RemEpu8(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256RemEpu8([32]byte(a), [32]byte(b)))
}

func m256RemEpu8(a [32]byte, b [32]byte) [32]byte


// M256RoundPd: Round the packed double-precision (64-bit) floating-point
// elements in 'a' using the 'rounding' parameter, and store the results as
// packed double-precision floating-point elements in 'dst'.
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
func M256RoundPd(a x86.M256d, rounding int) (dst x86.M256d) {
	return x86.M256d(m256RoundPd([4]float64(a), rounding))
}

func m256RoundPd(a [4]float64, rounding int) [4]float64


// M256RoundPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' using the 'rounding' parameter, and store the results as
// packed single-precision floating-point elements in 'dst'.
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
func M256RoundPs(a x86.M256, rounding int) (dst x86.M256) {
	return x86.M256(m256RoundPs([8]float32(a), rounding))
}

func m256RoundPs(a [8]float32, rounding int) [8]float32


// M256RsqrtPs: Compute the approximate reciprocal square root of packed
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
func M256RsqrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256RsqrtPs([8]float32(a)))
}

func m256RsqrtPs(a [8]float32) [8]float32


// M256SetEpi16: Set packed 16-bit integers in 'dst' with the supplied values. 
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
func M256SetEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) (dst x86.M256i) {
	return x86.M256i(m256SetEpi16(e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [32]byte


// M256SetEpi32: Set packed 32-bit integers in 'dst' with the supplied values. 
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
func M256SetEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) (dst x86.M256i) {
	return x86.M256i(m256SetEpi32(e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [32]byte


// M256SetEpi64x: Set packed 64-bit integers in 'dst' with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//		dst[191:128] := e2
//		dst[255:192] := e3
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_epi64x'.
// Requires AVX.
func M256SetEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) (dst x86.M256i) {
	return x86.M256i(m256SetEpi64x(e3, e2, e1, e0))
}

func m256SetEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) [32]byte


// M256SetEpi8: Set packed 8-bit integers in 'dst' with the supplied values in
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
func M256SetEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) (dst x86.M256i) {
	return x86.M256i(m256SetEpi8(e31, e30, e29, e28, e27, e26, e25, e24, e23, e22, e21, e20, e19, e18, e17, e16, e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [32]byte


// M256SetM128: Set packed __m256 vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_set_m128'.
// Requires AVX.
func M256SetM128(hi x86.M128, lo x86.M128) (dst x86.M256) {
	return x86.M256(m256SetM128([4]float32(hi), [4]float32(lo)))
}

func m256SetM128(hi [4]float32, lo [4]float32) [8]float32


// M256SetM128d: Set packed __m256d vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_set_m128d'.
// Requires AVX.
func M256SetM128d(hi x86.M128d, lo x86.M128d) (dst x86.M256d) {
	return x86.M256d(m256SetM128d([2]float64(hi), [2]float64(lo)))
}

func m256SetM128d(hi [2]float64, lo [2]float64) [4]float64


// M256SetM128i: Set packed __m256i vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_set_m128i'.
// Requires AVX.
func M256SetM128i(hi x86.M128i, lo x86.M128i) (dst x86.M256i) {
	return x86.M256i(m256SetM128i([16]byte(hi), [16]byte(lo)))
}

func m256SetM128i(hi [16]byte, lo [16]byte) [32]byte


// M256SetPd: Set packed double-precision (64-bit) floating-point elements in
// 'dst' with the supplied values. 
//
//		dst[63:0] := e0
//		dst[127:64] := e1
//		dst[191:128] := e2
//		dst[255:192] := e3
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set_pd'.
// Requires AVX.
func M256SetPd(e3 float64, e2 float64, e1 float64, e0 float64) (dst x86.M256d) {
	return x86.M256d(m256SetPd(e3, e2, e1, e0))
}

func m256SetPd(e3 float64, e2 float64, e1 float64, e0 float64) [4]float64


// M256SetPs: Set packed single-precision (32-bit) floating-point elements in
// 'dst' with the supplied values. 
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
func M256SetPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) (dst x86.M256) {
	return x86.M256(m256SetPs(e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [8]float32


// M256Set1Epi16: Broadcast 16-bit integer 'a' to all all elements of 'dst'.
// This intrinsic may generate the 'vpbroadcastw'. 
//
//		FOR j := 0 to 15
//			i := j*16
//			dst[i+15:i] := a[15:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_epi16'.
// Requires AVX.
func M256Set1Epi16(a int16) (dst x86.M256i) {
	return x86.M256i(m256Set1Epi16(a))
}

func m256Set1Epi16(a int16) [32]byte


// M256Set1Epi32: Broadcast 32-bit integer 'a' to all elements of 'dst'. This
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
func M256Set1Epi32(a int) (dst x86.M256i) {
	return x86.M256i(m256Set1Epi32(a))
}

func m256Set1Epi32(a int) [32]byte


// M256Set1Epi64x: Broadcast 64-bit integer 'a' to all elements of 'dst'. This
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
func M256Set1Epi64x(a int64) (dst x86.M256i) {
	return x86.M256i(m256Set1Epi64x(a))
}

func m256Set1Epi64x(a int64) [32]byte


// M256Set1Epi8: Broadcast 8-bit integer 'a' to all elements of 'dst'. This
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
func M256Set1Epi8(a byte) (dst x86.M256i) {
	return x86.M256i(m256Set1Epi8(a))
}

func m256Set1Epi8(a byte) [32]byte


// M256Set1Pd: Broadcast double-precision (64-bit) floating-point value 'a' to
// all elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := a[63:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_pd'.
// Requires AVX.
func M256Set1Pd(a float64) (dst x86.M256d) {
	return x86.M256d(m256Set1Pd(a))
}

func m256Set1Pd(a float64) [4]float64


// M256Set1Ps: Broadcast single-precision (32-bit) floating-point value 'a' to
// all elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := a[31:0]
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_set1_ps'.
// Requires AVX.
func M256Set1Ps(a float32) (dst x86.M256) {
	return x86.M256(m256Set1Ps(a))
}

func m256Set1Ps(a float32) [8]float32


// M256SetrEpi16: Set packed 16-bit integers in 'dst' with the supplied values
// in reverse order. 
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
func M256SetrEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) (dst x86.M256i) {
	return x86.M256i(m256SetrEpi16(e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetrEpi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) [32]byte


// M256SetrEpi32: Set packed 32-bit integers in 'dst' with the supplied values
// in reverse order. 
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
func M256SetrEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) (dst x86.M256i) {
	return x86.M256i(m256SetrEpi32(e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetrEpi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) [32]byte


// M256SetrEpi64x: Set packed 64-bit integers in 'dst' with the supplied values
// in reverse order. 
//
//		dst[63:0] := e3
//		dst[127:64] := e2
//		dst[191:128] := e1
//		dst[255:192] := e0
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_setr_epi64x'.
// Requires AVX.
func M256SetrEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) (dst x86.M256i) {
	return x86.M256i(m256SetrEpi64x(e3, e2, e1, e0))
}

func m256SetrEpi64x(e3 int64, e2 int64, e1 int64, e0 int64) [32]byte


// M256SetrEpi8: Set packed 8-bit integers in 'dst' with the supplied values in
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
func M256SetrEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) (dst x86.M256i) {
	return x86.M256i(m256SetrEpi8(e31, e30, e29, e28, e27, e26, e25, e24, e23, e22, e21, e20, e19, e18, e17, e16, e15, e14, e13, e12, e11, e10, e9, e8, e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetrEpi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) [32]byte


// M256SetrM128: Set packed __m256 vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_setr_m128'.
// Requires AVX.
func M256SetrM128(lo x86.M128, hi x86.M128) (dst x86.M256) {
	return x86.M256(m256SetrM128([4]float32(lo), [4]float32(hi)))
}

func m256SetrM128(lo [4]float32, hi [4]float32) [8]float32


// M256SetrM128d: Set packed __m256d vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_setr_m128d'.
// Requires AVX.
func M256SetrM128d(lo x86.M128d, hi x86.M128d) (dst x86.M256d) {
	return x86.M256d(m256SetrM128d([2]float64(lo), [2]float64(hi)))
}

func m256SetrM128d(lo [2]float64, hi [2]float64) [4]float64


// M256SetrM128i: Set packed __m256i vector 'dst' with the supplied values. 
//
//		dst[127:0] := lo[127:0]
//		dst[255:128] := hi[127:0]
//		dst[MAX:256] := 0
//
// Instruction: 'VINSERTF128'. Intrinsic: '_mm256_setr_m128i'.
// Requires AVX.
func M256SetrM128i(lo x86.M128i, hi x86.M128i) (dst x86.M256i) {
	return x86.M256i(m256SetrM128i([16]byte(lo), [16]byte(hi)))
}

func m256SetrM128i(lo [16]byte, hi [16]byte) [32]byte


// M256SetrPd: Set packed double-precision (64-bit) floating-point elements in
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
func M256SetrPd(e3 float64, e2 float64, e1 float64, e0 float64) (dst x86.M256d) {
	return x86.M256d(m256SetrPd(e3, e2, e1, e0))
}

func m256SetrPd(e3 float64, e2 float64, e1 float64, e0 float64) [4]float64


// M256SetrPs: Set packed single-precision (32-bit) floating-point elements in
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
func M256SetrPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) (dst x86.M256) {
	return x86.M256(m256SetrPs(e7, e6, e5, e4, e3, e2, e1, e0))
}

func m256SetrPs(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) [8]float32


// M256SetzeroPd: Return vector of type __m256d with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'VXORPD'. Intrinsic: '_mm256_setzero_pd'.
// Requires AVX.
func M256SetzeroPd() (dst x86.M256d) {
	return x86.M256d(m256SetzeroPd())
}

func m256SetzeroPd() [4]float64


// M256SetzeroPs: Return vector of type __m256 with all elements set to zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'VXORPS'. Intrinsic: '_mm256_setzero_ps'.
// Requires AVX.
func M256SetzeroPs() (dst x86.M256) {
	return x86.M256(m256SetzeroPs())
}

func m256SetzeroPs() [8]float32


// M256SetzeroSi256: Return vector of type __m256i with all elements set to
// zero. 
//
//		dst[MAX:0] := 0
//
// Instruction: 'VPXOR'. Intrinsic: '_mm256_setzero_si256'.
// Requires AVX.
func M256SetzeroSi256() (dst x86.M256i) {
	return x86.M256i(m256SetzeroSi256())
}

func m256SetzeroSi256() [32]byte


// M256ShufflePd: Shuffle double-precision (64-bit) floating-point elements
// within 128-bit lanes using the control in 'imm8', and store the results in
// 'dst'. 
//
//		dst[63:0] := (imm8[0] == 0) ? a[63:0] : a[127:64]
//		dst[127:64] := (imm8[1] == 0) ? b[63:0] : b[127:64]
//		dst[191:128] := (imm8[2] == 0) ? a[191:128] : a[255:192]
//		dst[255:192] := (imm8[3] == 0) ? b[191:128] : b[255:192]
//		dst[MAX:256] := 0
//
// Instruction: 'VSHUFPD'. Intrinsic: '_mm256_shuffle_pd'.
// Requires AVX.
//
// FIXME: Requires compiler support (has immediate)
func M256ShufflePd(a x86.M256d, b x86.M256d, imm8 byte) (dst x86.M256d) {
	return x86.M256d(m256ShufflePd([4]float64(a), [4]float64(b), imm8))
}

func m256ShufflePd(a [4]float64, b [4]float64, imm8 byte) [4]float64


// M256ShufflePs: Shuffle single-precision (32-bit) floating-point elements in
// 'a' within 128-bit lanes using the control in 'imm8', and store the results
// in 'dst'. 
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
//
// FIXME: Requires compiler support (has immediate)
func M256ShufflePs(a x86.M256, b x86.M256, imm8 byte) (dst x86.M256) {
	return x86.M256(m256ShufflePs([8]float32(a), [8]float32(b), imm8))
}

func m256ShufflePs(a [8]float32, b [8]float32, imm8 byte) [8]float32


// M256SinPd: Compute the sine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SIN(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sin_pd'.
// Requires AVX.
func M256SinPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SinPd([4]float64(a)))
}

func m256SinPd(a [4]float64) [4]float64


// M256SinPs: Compute the sine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in radians, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SIN(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sin_ps'.
// Requires AVX.
func M256SinPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SinPs([8]float32(a)))
}

func m256SinPs(a [8]float32) [8]float32


// M256SincosPd: Compute the sine and cosine of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, store the sine
// in 'dst', and store the cosine into memory at 'mem_addr'. 
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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256SincosPd(mem_addr *x86.M256d, a x86.M256d) (dst x86.M256d) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256d{}
}

// M256SincosPs: Compute the sine and cosine of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, store the sine
// in 'dst', and store the cosine into memory at 'mem_addr'. 
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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256SincosPs(mem_addr *x86.M256, a x86.M256) (dst x86.M256) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256{}
}

// M256SindPd: Compute the sine of packed double-precision (64-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := SIND(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sind_pd'.
// Requires AVX.
func M256SindPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SindPd([4]float64(a)))
}

func m256SindPd(a [4]float64) [4]float64


// M256SindPs: Compute the sine of packed single-precision (32-bit)
// floating-point elements in 'a' expressed in degrees, and store the results
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := SIND(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_sind_ps'.
// Requires AVX.
func M256SindPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SindPs([8]float32(a)))
}

func m256SindPs(a [8]float32) [8]float32


// M256SinhPd: Compute the hyperbolic sine of packed double-precision (64-bit)
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
func M256SinhPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SinhPd([4]float64(a)))
}

func m256SinhPd(a [4]float64) [4]float64


// M256SinhPs: Compute the hyperbolic sine of packed single-precision (32-bit)
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
func M256SinhPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SinhPs([8]float32(a)))
}

func m256SinhPs(a [8]float32) [8]float32


// M256SqrtPd: Compute the square root of packed double-precision (64-bit)
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
func M256SqrtPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SqrtPd([4]float64(a)))
}

func m256SqrtPd(a [4]float64) [4]float64


// M256SqrtPs: Compute the square root of packed single-precision (32-bit)
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
func M256SqrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SqrtPs([8]float32(a)))
}

func m256SqrtPs(a [8]float32) [8]float32


// M256StorePd: Store 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm256_store_pd'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StorePd(mem_addr *float64, a x86.M256d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StorePs: Store 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm256_store_ps'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StorePs(mem_addr *float32, a x86.M256)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StoreSi256: Store 256-bits of integer data from 'a' into memory.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVDQA'. Intrinsic: '_mm256_store_si256'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StoreSi256(mem_addr *x86.M256i, a x86.M256i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StoreuPd: Store 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVUPD'. Intrinsic: '_mm256_storeu_pd'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StoreuPd(mem_addr *float64, a x86.M256d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StoreuPs: Store 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVUPS'. Intrinsic: '_mm256_storeu_ps'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StoreuPs(mem_addr *float32, a x86.M256)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StoreuSi256: Store 256-bits of integer data from 'a' into memory.
// 	'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVDQU'. Intrinsic: '_mm256_storeu_si256'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StoreuSi256(mem_addr *x86.M256i, a x86.M256i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256Storeu2M128: Store the high and low 128-bit halves (each composed of 4
// packed single-precision (32-bit) floating-point elements) from 'a' into
// memory two different 128-bit locations.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		MEM[loaddr+127:loaddr] := a[127:0]
//		MEM[hiaddr+127:hiaddr] := a[255:128]
//
// Instruction: '...'. Intrinsic: '_mm256_storeu2_m128'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256Storeu2M128(hiaddr *float32, loaddr *float32, a x86.M256)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256Storeu2M128d: Store the high and low 128-bit halves (each composed of 2
// packed double-precision (64-bit) floating-point elements) from 'a' into
// memory two different 128-bit locations.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		MEM[loaddr+127:loaddr] := a[127:0]
//		MEM[hiaddr+127:hiaddr] := a[255:128]
//
// Instruction: '...'. Intrinsic: '_mm256_storeu2_m128d'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256Storeu2M128d(hiaddr *float64, loaddr *float64, a x86.M256d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256Storeu2M128i: Store the high and low 128-bit halves (each composed of
// integer data) from 'a' into memory two different 128-bit locations.
// 	'hiaddr' and 'loaddr' do not need to be aligned on any particular boundary. 
//
//		MEM[loaddr+127:loaddr] := a[127:0]
//		MEM[hiaddr+127:hiaddr] := a[255:128]
//
// Instruction: '...'. Intrinsic: '_mm256_storeu2_m128i'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256Storeu2M128i(hiaddr *x86.M128i, loaddr *x86.M128i, a x86.M256i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StreamPd: Store 256-bits (composed of 4 packed double-precision (64-bit)
// floating-point elements) from 'a' into memory using a non-temporal memory
// hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVNTPD'. Intrinsic: '_mm256_stream_pd'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StreamPd(mem_addr *float64, a x86.M256d)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StreamPs: Store 256-bits (composed of 8 packed single-precision (32-bit)
// floating-point elements) from 'a' into memory using a non-temporal memory
// hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVNTPS'. Intrinsic: '_mm256_stream_ps'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StreamPs(mem_addr *float32, a x86.M256)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256StreamSi256: Store 256-bits of integer data from 'a' into memory using a
// non-temporal memory hint.
// 	'mem_addr' must be aligned on a 32-byte boundary or a general-protection
// exception may be generated. 
//
//		MEM[mem_addr+255:mem_addr] := a[255:0]
//
// Instruction: 'VMOVNTDQ'. Intrinsic: '_mm256_stream_si256'.
// Requires AVX.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256StreamSi256(mem_addr *x86.M256i, a x86.M256i)  {
	// FIXME: Rework to avoid possible return value as parameter.

}

// M256SubPd: Subtract packed double-precision (64-bit) floating-point elements
// in 'b' from packed double-precision (64-bit) floating-point elements in 'a',
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
func M256SubPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SubPd([4]float64(a), [4]float64(b)))
}

func m256SubPd(a [4]float64, b [4]float64) [4]float64


// M256SubPs: Subtract packed single-precision (32-bit) floating-point elements
// in 'b' from packed single-precision (32-bit) floating-point elements in 'a',
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
func M256SubPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256SubPs([8]float32(a), [8]float32(b)))
}

func m256SubPs(a [8]float32, b [8]float32) [8]float32


// M256SvmlCeilPd: Round the packed double-precision (64-bit) floating-point
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
func M256SvmlCeilPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SvmlCeilPd([4]float64(a)))
}

func m256SvmlCeilPd(a [4]float64) [4]float64


// M256SvmlCeilPs: Round the packed single-precision (32-bit) floating-point
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
func M256SvmlCeilPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SvmlCeilPs([8]float32(a)))
}

func m256SvmlCeilPs(a [8]float32) [8]float32


// M256SvmlFloorPd: Round the packed double-precision (64-bit) floating-point
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
func M256SvmlFloorPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SvmlFloorPd([4]float64(a)))
}

func m256SvmlFloorPd(a [4]float64) [4]float64


// M256SvmlFloorPs: Round the packed single-precision (32-bit) floating-point
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
func M256SvmlFloorPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SvmlFloorPs([8]float32(a)))
}

func m256SvmlFloorPs(a [8]float32) [8]float32


// M256SvmlRoundPd: Round the packed double-precision (64-bit) floating-point
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
func M256SvmlRoundPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SvmlRoundPd([4]float64(a)))
}

func m256SvmlRoundPd(a [4]float64) [4]float64


// M256SvmlRoundPs: Round the packed single-precision (32-bit) floating-point
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
func M256SvmlRoundPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SvmlRoundPs([8]float32(a)))
}

func m256SvmlRoundPs(a [8]float32) [8]float32


// M256SvmlSqrtPd: Compute the square root of packed double-precision (64-bit)
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
func M256SvmlSqrtPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256SvmlSqrtPd([4]float64(a)))
}

func m256SvmlSqrtPd(a [4]float64) [4]float64


// M256SvmlSqrtPs: Compute the square root of packed single-precision (32-bit)
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
func M256SvmlSqrtPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256SvmlSqrtPs([8]float32(a)))
}

func m256SvmlSqrtPs(a [8]float32) [8]float32


// M256TanPd: Compute the tangent of packed double-precision (64-bit)
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
func M256TanPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256TanPd([4]float64(a)))
}

func m256TanPd(a [4]float64) [4]float64


// M256TanPs: Compute the tangent of packed single-precision (32-bit)
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
func M256TanPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256TanPs([8]float32(a)))
}

func m256TanPs(a [8]float32) [8]float32


// M256TandPd: Compute the tangent of packed double-precision (64-bit)
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
func M256TandPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256TandPd([4]float64(a)))
}

func m256TandPd(a [4]float64) [4]float64


// M256TandPs: Compute the tangent of packed single-precision (32-bit)
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
func M256TandPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256TandPs([8]float32(a)))
}

func m256TandPs(a [8]float32) [8]float32


// M256TanhPd: Compute the hyperbolic tangent of packed double-precision
// (64-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := TANH(a[i+63:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tanh_pd'.
// Requires AVX.
func M256TanhPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256TanhPd([4]float64(a)))
}

func m256TanhPd(a [4]float64) [4]float64


// M256TanhPs: Compute the hyperbolic tangent of packed single-precision
// (32-bit) floating-point elements in 'a' expressed in radians, and store the
// results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := TANH(a[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_tanh_ps'.
// Requires AVX.
func M256TanhPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256TanhPs([8]float32(a)))
}

func m256TanhPs(a [8]float32) [8]float32


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
func TestcPd(a x86.M128d, b x86.M128d) int {
	return int(testcPd([2]float64(a), [2]float64(b)))
}

func testcPd(a [2]float64, b [2]float64) int


// M256TestcPd: Compute the bitwise AND of 256 bits (representing
// double-precision (64-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 256-bit value, and set 'ZF' to 1 if the sign bit of each
// 64-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 64-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return the 'CF' value. 
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
func M256TestcPd(a x86.M256d, b x86.M256d) int {
	return int(m256TestcPd([4]float64(a), [4]float64(b)))
}

func m256TestcPd(a [4]float64, b [4]float64) int


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
func TestcPs(a x86.M128, b x86.M128) int {
	return int(testcPs([4]float32(a), [4]float32(b)))
}

func testcPs(a [4]float32, b [4]float32) int


// M256TestcPs: Compute the bitwise AND of 256 bits (representing
// single-precision (32-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 256-bit value, and set 'ZF' to 1 if the sign bit of each
// 32-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 32-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return the 'CF' value. 
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
func M256TestcPs(a x86.M256, b x86.M256) int {
	return int(m256TestcPs([8]float32(a), [8]float32(b)))
}

func m256TestcPs(a [8]float32, b [8]float32) int


// M256TestcSi256: Compute the bitwise AND of 256 bits (representing integer
// data) in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set
// 'ZF' to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if
// the result is zero, otherwise set 'CF' to 0. Return the 'CF' value. 
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
func M256TestcSi256(a x86.M256i, b x86.M256i) int {
	return int(m256TestcSi256([32]byte(a), [32]byte(b)))
}

func m256TestcSi256(a [32]byte, b [32]byte) int


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
func TestnzcPd(a x86.M128d, b x86.M128d) int {
	return int(testnzcPd([2]float64(a), [2]float64(b)))
}

func testnzcPd(a [2]float64, b [2]float64) int


// M256TestnzcPd: Compute the bitwise AND of 256 bits (representing
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
func M256TestnzcPd(a x86.M256d, b x86.M256d) int {
	return int(m256TestnzcPd([4]float64(a), [4]float64(b)))
}

func m256TestnzcPd(a [4]float64, b [4]float64) int


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
func TestnzcPs(a x86.M128, b x86.M128) int {
	return int(testnzcPs([4]float32(a), [4]float32(b)))
}

func testnzcPs(a [4]float32, b [4]float32) int


// M256TestnzcPs: Compute the bitwise AND of 256 bits (representing
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
func M256TestnzcPs(a x86.M256, b x86.M256) int {
	return int(m256TestnzcPs([8]float32(a), [8]float32(b)))
}

func m256TestnzcPs(a [8]float32, b [8]float32) int


// M256TestnzcSi256: Compute the bitwise AND of 256 bits (representing integer
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
func M256TestnzcSi256(a x86.M256i, b x86.M256i) int {
	return int(m256TestnzcSi256([32]byte(a), [32]byte(b)))
}

func m256TestnzcSi256(a [32]byte, b [32]byte) int


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
func TestzPd(a x86.M128d, b x86.M128d) int {
	return int(testzPd([2]float64(a), [2]float64(b)))
}

func testzPd(a [2]float64, b [2]float64) int


// M256TestzPd: Compute the bitwise AND of 256 bits (representing
// double-precision (64-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 256-bit value, and set 'ZF' to 1 if the sign bit of each
// 64-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 64-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return the 'ZF' value. 
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
func M256TestzPd(a x86.M256d, b x86.M256d) int {
	return int(m256TestzPd([4]float64(a), [4]float64(b)))
}

func m256TestzPd(a [4]float64, b [4]float64) int


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
func TestzPs(a x86.M128, b x86.M128) int {
	return int(testzPs([4]float32(a), [4]float32(b)))
}

func testzPs(a [4]float32, b [4]float32) int


// M256TestzPs: Compute the bitwise AND of 256 bits (representing
// single-precision (32-bit) floating-point elements) in 'a' and 'b', producing
// an intermediate 256-bit value, and set 'ZF' to 1 if the sign bit of each
// 32-bit element in the intermediate value is zero, otherwise set 'ZF' to 0.
// Compute the bitwise AND NOT of 'a' and 'b', producing an intermediate value,
// and set 'CF' to 1 if the sign bit of each 32-bit element in the intermediate
// value is zero, otherwise set 'CF' to 0. Return the 'ZF' value. 
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
func M256TestzPs(a x86.M256, b x86.M256) int {
	return int(m256TestzPs([8]float32(a), [8]float32(b)))
}

func m256TestzPs(a [8]float32, b [8]float32) int


// M256TestzSi256: Compute the bitwise AND of 256 bits (representing integer
// data) in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set
// 'ZF' to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if
// the result is zero, otherwise set 'CF' to 0. Return the 'ZF' value. 
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
func M256TestzSi256(a x86.M256i, b x86.M256i) int {
	return int(m256TestzSi256([32]byte(a), [32]byte(b)))
}

func m256TestzSi256(a [32]byte, b [32]byte) int


// M256TruncPd: Truncate the packed double-precision (64-bit) floating-point
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
func M256TruncPd(a x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256TruncPd([4]float64(a)))
}

func m256TruncPd(a [4]float64) [4]float64


// M256TruncPs: Truncate the packed single-precision (32-bit) floating-point
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
func M256TruncPs(a x86.M256) (dst x86.M256) {
	return x86.M256(m256TruncPs([8]float32(a)))
}

func m256TruncPs(a [8]float32) [8]float32


// M256UdivEpi32: Divide packed unsigned 32-bit integers in 'a' by packed
// elements in 'b', and store the truncated results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := TRUNCATE(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_udiv_epi32'.
// Requires AVX.
func M256UdivEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256UdivEpi32([32]byte(a), [32]byte(b)))
}

func m256UdivEpi32(a [32]byte, b [32]byte) [32]byte


// M256UdivremEpi32: Divide packed unsigned 32-bit integers in 'a' by packed
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
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M256UdivremEpi32(mem_addr *x86.M256i, a x86.M256i, b x86.M256i) (dst x86.M256i) {
	// FIXME: Rework to avoid possible return value as parameter.
	return x86.M256i{}
}

// UndefinedPd: Return vector of type __m128d with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_undefined_pd'.
// Requires AVX.
func UndefinedPd() (dst x86.M128d) {
	return x86.M128d(undefinedPd())
}

func undefinedPd() [2]float64


// M256UndefinedPd: Return vector of type __m256d with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_undefined_pd'.
// Requires AVX.
func M256UndefinedPd() (dst x86.M256d) {
	return x86.M256d(m256UndefinedPd())
}

func m256UndefinedPd() [4]float64


// UndefinedPs: Return vector of type __m128 with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_undefined_ps'.
// Requires AVX.
func UndefinedPs() (dst x86.M128) {
	return x86.M128(undefinedPs())
}

func undefinedPs() [4]float32


// M256UndefinedPs: Return vector of type __m256 with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_undefined_ps'.
// Requires AVX.
func M256UndefinedPs() (dst x86.M256) {
	return x86.M256(m256UndefinedPs())
}

func m256UndefinedPs() [8]float32


// UndefinedSi128: Return vector of type __m128i with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_undefined_si128'.
// Requires AVX.
func UndefinedSi128() (dst x86.M128i) {
	return x86.M128i(undefinedSi128())
}

func undefinedSi128() [16]byte


// M256UndefinedSi256: Return vector of type __m256i with undefined elements. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm256_undefined_si256'.
// Requires AVX.
func M256UndefinedSi256() (dst x86.M256i) {
	return x86.M256i(m256UndefinedSi256())
}

func m256UndefinedSi256() [32]byte


// M256UnpackhiPd: Unpack and interleave double-precision (64-bit)
// floating-point elements from the high half of each 128-bit lane in 'a' and
// 'b', and store the results in 'dst'. 
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
func M256UnpackhiPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256UnpackhiPd([4]float64(a), [4]float64(b)))
}

func m256UnpackhiPd(a [4]float64, b [4]float64) [4]float64


// M256UnpackhiPs: Unpack and interleave single-precision (32-bit)
// floating-point elements from the high half of each 128-bit lane in 'a' and
// 'b', and store the results in 'dst'. 
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
func M256UnpackhiPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256UnpackhiPs([8]float32(a), [8]float32(b)))
}

func m256UnpackhiPs(a [8]float32, b [8]float32) [8]float32


// M256UnpackloPd: Unpack and interleave double-precision (64-bit)
// floating-point elements from the low half of each 128-bit lane in 'a' and
// 'b', and store the results in 'dst'. 
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
func M256UnpackloPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256UnpackloPd([4]float64(a), [4]float64(b)))
}

func m256UnpackloPd(a [4]float64, b [4]float64) [4]float64


// M256UnpackloPs: Unpack and interleave single-precision (32-bit)
// floating-point elements from the low half of each 128-bit lane in 'a' and
// 'b', and store the results in 'dst'. 
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
func M256UnpackloPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256UnpackloPs([8]float32(a), [8]float32(b)))
}

func m256UnpackloPs(a [8]float32, b [8]float32) [8]float32


// M256UremEpi32: Divide packed unsigned 32-bit integers in 'a' by packed
// elements in 'b', and store the remainders as packed unsigned 32-bit integers
// in 'dst'. 
//
//		FOR j := 0 to 7
//			i := 32*j
//			dst[i+31:i] := REMAINDER(a[i+31:i] / b[i+31:i])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: '...'. Intrinsic: '_mm256_urem_epi32'.
// Requires AVX.
func M256UremEpi32(a x86.M256i, b x86.M256i) (dst x86.M256i) {
	return x86.M256i(m256UremEpi32([32]byte(a), [32]byte(b)))
}

func m256UremEpi32(a [32]byte, b [32]byte) [32]byte


// M256XorPd: Compute the bitwise XOR of packed double-precision (64-bit)
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
func M256XorPd(a x86.M256d, b x86.M256d) (dst x86.M256d) {
	return x86.M256d(m256XorPd([4]float64(a), [4]float64(b)))
}

func m256XorPd(a [4]float64, b [4]float64) [4]float64


// M256XorPs: Compute the bitwise XOR of packed single-precision (32-bit)
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
func M256XorPs(a x86.M256, b x86.M256) (dst x86.M256) {
	return x86.M256(m256XorPs([8]float32(a), [8]float32(b)))
}

func m256XorPs(a [8]float32, b [8]float32) [8]float32


// M256Zeroall: Zero the contents of all XMM or YMM registers. 
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
func M256Zeroall()  {
	m256Zeroall()
}

func m256Zeroall() 


// M256Zeroupper: Zero the upper 128 bits of all YMM registers; the lower
// 128-bits of the registers are unmodified. 
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
func M256Zeroupper()  {
	m256Zeroupper()
}

func m256Zeroupper() 

