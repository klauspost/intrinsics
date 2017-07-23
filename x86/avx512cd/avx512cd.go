// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package avx512cd

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// BroadcastmbEpi64: Broadcast the low 8-bits from input mask 'k' to all 64-bit
// elements of 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ZeroExtend(k[7:0])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTMB2Q'. Intrinsic: '_mm_broadcastmb_epi64'.
// Requires AVX512CD.
func BroadcastmbEpi64(k x86.Mmask8) (dst x86.M128i) {
	panic("not implemented")
}


// M256BroadcastmbEpi64: Broadcast the low 8-bits from input mask 'k' to all
// 64-bit elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			dst[i+63:i] := ZeroExtend(k[7:0])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTMB2Q'. Intrinsic: '_mm256_broadcastmb_epi64'.
// Requires AVX512CD.
func M256BroadcastmbEpi64(k x86.Mmask8) (dst x86.M256i) {
	panic("not implemented")
}


// M512BroadcastmbEpi64: Broadcast the low 8-bits from input mask 'k' to all
// 64-bit elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ZeroExtend(k[7:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTMB2Q'. Intrinsic: '_mm512_broadcastmb_epi64'.
// Requires AVX512CD.
func M512BroadcastmbEpi64(k x86.Mmask8) (dst x86.M512i) {
	panic("not implemented")
}


// BroadcastmwEpi32: Broadcast the low 16-bits from input mask 'k' to all
// 32-bit elements of 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ZeroExtend(k[15:0])
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPBROADCASTMW2D'. Intrinsic: '_mm_broadcastmw_epi32'.
// Requires AVX512CD.
func BroadcastmwEpi32(k x86.Mmask16) (dst x86.M128i) {
	panic("not implemented")
}


// M256BroadcastmwEpi32: Broadcast the low 16-bits from input mask 'k' to all
// 32-bit elements of 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			dst[i+31:i] := ZeroExtend(k[15:0])
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPBROADCASTMW2D'. Intrinsic: '_mm256_broadcastmw_epi32'.
// Requires AVX512CD.
func M256BroadcastmwEpi32(k x86.Mmask16) (dst x86.M256i) {
	panic("not implemented")
}


// M512BroadcastmwEpi32: Broadcast the low 16-bits from input mask 'k' to all
// 32-bit elements of 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ZeroExtend(k[15:0])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBROADCASTMW2D'. Intrinsic: '_mm512_broadcastmw_epi32'.
// Requires AVX512CD.
func M512BroadcastmwEpi32(k x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// ConflictEpi32: Test each 32-bit element of 'a' for equality with all other
// elements in 'a' closer to the least significant bit. Each element's
// comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			FOR k := 0 to j-1
//				m := k*32
//				dst[i+k] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//			ENDFOR
//			dst[i+31:i+j] := 0
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm_conflict_epi32'.
// Requires AVX512CD.
func ConflictEpi32(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskConflictEpi32: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[i]
//				FOR l := 0 to j-1
//					m := l*32
//					dst[i+l] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//				ENDFOR
//				dst[i+31:i+j] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm_mask_conflict_epi32'.
// Requires AVX512CD.
func MaskConflictEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzConflictEpi32: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). Each
// element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[i]
//				FOR l := 0 to j-1
//					m := l*32
//					dst[i+l] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//				ENDFOR
//				dst[i+31:i+j] := 0
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm_maskz_conflict_epi32'.
// Requires AVX512CD.
func MaskzConflictEpi32(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256ConflictEpi32: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit. Each element's
// comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			FOR k := 0 to j-1
//				m := k*32
//				dst[i+k] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//			ENDFOR
//			dst[i+31:i+j] := 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm256_conflict_epi32'.
// Requires AVX512CD.
func M256ConflictEpi32(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskConflictEpi32: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[i]
//				FOR l := 0 to j-1
//					m := l*32
//					dst[i+l] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//				ENDFOR
//				dst[i+31:i+j] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm256_mask_conflict_epi32'.
// Requires AVX512CD.
func M256MaskConflictEpi32(src x86.M256i, k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzConflictEpi32: Test each 32-bit element of 'a' for equality with
// all other elements in 'a' closer to the least significant bit using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set).
// Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[i]
//				FOR l := 0 to j-1
//					m := l*32
//					dst[i+l] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//				ENDFOR
//				dst[i+31:i+j] := 0
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm256_maskz_conflict_epi32'.
// Requires AVX512CD.
func M256MaskzConflictEpi32(k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512ConflictEpi32: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit. Each element's
// comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			FOR k := 0 to j-1
//				m := k*32
//				dst[i+k] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//			ENDFOR
//			dst[i+31:i+j] := 0
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm512_conflict_epi32'.
// Requires AVX512CD.
func M512ConflictEpi32(a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskConflictEpi32: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[i]
//				FOR l := 0 to j-1
//					m := l*32
//					dst[i+l] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//				ENDFOR
//				dst[i+31:i+j] := 0
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm512_mask_conflict_epi32'.
// Requires AVX512CD.
func M512MaskConflictEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzConflictEpi32: Test each 32-bit element of 'a' for equality with
// all other elements in 'a' closer to the least significant bit using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set).
// Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[i]
//				FOR l := 0 to j-1
//					m := l*32
//					dst[i+l] := (a[i+31:i] == a[m+31:m]) ? 1 : 0
//				ENDFOR
//				dst[i+31:i+j] := 0
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPCONFLICTD'. Intrinsic: '_mm512_maskz_conflict_epi32'.
// Requires AVX512CD.
func M512MaskzConflictEpi32(k x86.Mmask16, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// ConflictEpi64: Test each 64-bit element of 'a' for equality with all other
// elements in 'a' closer to the least significant bit. Each element's
// comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			FOR k := 0 to j-1
//				m := k*64
//				dst[i+k] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//			ENDFOR
//			dst[i+63:i+j] := 0
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm_conflict_epi64'.
// Requires AVX512CD.
func ConflictEpi64(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskConflictEpi64: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				FOR l := 0 to j-1
//					m := l*64
//					dst[i+l] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//				ENDFOR
//				dst[i+63:i+j] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm_mask_conflict_epi64'.
// Requires AVX512CD.
func MaskConflictEpi64(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzConflictEpi64: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). Each
// element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				FOR l := 0 to j-1
//					m := l*64
//					dst[i+l] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//				ENDFOR
//				dst[i+63:i+j] := 0
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm_maskz_conflict_epi64'.
// Requires AVX512CD.
func MaskzConflictEpi64(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256ConflictEpi64: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit. Each element's
// comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			FOR k := 0 to j-1
//				m := k*64
//				dst[i+k] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//			ENDFOR
//			dst[i+63:i+j] := 0
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm256_conflict_epi64'.
// Requires AVX512CD.
func M256ConflictEpi64(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskConflictEpi64: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				FOR l := 0 to j-1
//					m := l*64
//					dst[i+l] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//				ENDFOR
//				dst[i+63:i+j] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm256_mask_conflict_epi64'.
// Requires AVX512CD.
func M256MaskConflictEpi64(src x86.M256i, k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzConflictEpi64: Test each 64-bit element of 'a' for equality with
// all other elements in 'a' closer to the least significant bit using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set).
// Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				FOR l := 0 to j-1
//					m := l*64
//					dst[i+l] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//				ENDFOR
//				dst[i+63:i+j] := 0
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm256_maskz_conflict_epi64'.
// Requires AVX512CD.
func M256MaskzConflictEpi64(k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512ConflictEpi64: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit. Each element's
// comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			FOR k := 0 to j-1
//				m := k*64
//				dst[i+k] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//			ENDFOR
//			dst[i+63:i+j] := 0
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm512_conflict_epi64'.
// Requires AVX512CD.
func M512ConflictEpi64(a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskConflictEpi64: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				FOR l := 0 to j-1
//					m := l*64
//					dst[i+l] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//				ENDFOR
//				dst[i+63:i+j] := 0
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm512_mask_conflict_epi64'.
// Requires AVX512CD.
func M512MaskConflictEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzConflictEpi64: Test each 64-bit element of 'a' for equality with
// all other elements in 'a' closer to the least significant bit using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set).
// Each element's comparison forms a zero extended bit vector in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				FOR l := 0 to j-1
//					m := l*64
//					dst[i+l] := (a[i+63:i] == a[m+63:m]) ? 1 : 0
//				ENDFOR
//				dst[i+63:i+j] := 0
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPCONFLICTQ'. Intrinsic: '_mm512_maskz_conflict_epi64'.
// Requires AVX512CD.
func M512MaskzConflictEpi64(k x86.Mmask8, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// LzcntEpi32: Counts the number of leading zero bits in each packed 32-bit
// integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			tmp := 31
//			dst[i+31:i] := 0
//			DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//				tmp := tmp - 1
//				dst[i+31:i] := dst[i+31:i] + 1
//			OD
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm_lzcnt_epi32'.
// Requires AVX512CD.
func LzcntEpi32(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskLzcntEpi32: Counts the number of leading zero bits in each packed 32-bit
// integer in 'a', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				tmp := 31
//				dst[i+31:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+31:i] := dst[i+31:i] + 1
//				OD
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm_mask_lzcnt_epi32'.
// Requires AVX512CD.
func MaskLzcntEpi32(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzLzcntEpi32: Counts the number of leading zero bits in each packed
// 32-bit integer in 'a', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF k[j]
//				tmp := 31
//				dst[i+31:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+31:i] := dst[i+31:i] + 1
//				OD
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm_maskz_lzcnt_epi32'.
// Requires AVX512CD.
func MaskzLzcntEpi32(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256LzcntEpi32: Counts the number of leading zero bits in each packed 32-bit
// integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			tmp := 31
//			dst[i+31:i] := 0
//			DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//				tmp := tmp - 1
//				dst[i+31:i] := dst[i+31:i] + 1
//			OD
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm256_lzcnt_epi32'.
// Requires AVX512CD.
func M256LzcntEpi32(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskLzcntEpi32: Counts the number of leading zero bits in each packed
// 32-bit integer in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				tmp := 31
//				dst[i+31:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+31:i] := dst[i+31:i] + 1
//				OD
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm256_mask_lzcnt_epi32'.
// Requires AVX512CD.
func M256MaskLzcntEpi32(src x86.M256i, k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzLzcntEpi32: Counts the number of leading zero bits in each packed
// 32-bit integer in 'a', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			IF k[j]
//				tmp := 31
//				dst[i+31:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+31:i] := dst[i+31:i] + 1
//				OD
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm256_maskz_lzcnt_epi32'.
// Requires AVX512CD.
func M256MaskzLzcntEpi32(k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512LzcntEpi32: Counts the number of leading zero bits in each packed 32-bit
// integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			tmp := 31
//			dst[i+31:i] := 0
//			DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//				tmp := tmp - 1
//				dst[i+31:i] := dst[i+31:i] + 1
//			OD
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm512_lzcnt_epi32'.
// Requires AVX512CD.
func M512LzcntEpi32(a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskLzcntEpi32: Counts the number of leading zero bits in each packed
// 32-bit integer in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				tmp := 31
//				dst[i+31:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+31:i] := dst[i+31:i] + 1
//				OD
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm512_mask_lzcnt_epi32'.
// Requires AVX512CD.
func M512MaskLzcntEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzLzcntEpi32: Counts the number of leading zero bits in each packed
// 32-bit integer in 'a', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				tmp := 31
//				dst[i+31:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+31:i] := dst[i+31:i] + 1
//				OD
//			ELSE
//				dst[i+31:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPLZCNTD'. Intrinsic: '_mm512_maskz_lzcnt_epi32'.
// Requires AVX512CD.
func M512MaskzLzcntEpi32(k x86.Mmask16, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// LzcntEpi64: Counts the number of leading zero bits in each packed 64-bit
// integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			tmp := 63
//			dst[i+63:i] := 0
//			DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//				tmp := tmp - 1
//				dst[i+63:i] := dst[i+63:i] + 1
//			OD
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm_lzcnt_epi64'.
// Requires AVX512CD.
func LzcntEpi64(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskLzcntEpi64: Counts the number of leading zero bits in each packed 64-bit
// integer in 'a', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp := 63
//				dst[i+63:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+63:i] := dst[i+63:i] + 1
//				OD
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm_mask_lzcnt_epi64'.
// Requires AVX512CD.
func MaskLzcntEpi64(src x86.M128i, k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// MaskzLzcntEpi64: Counts the number of leading zero bits in each packed
// 64-bit integer in 'a', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF k[j]
//				tmp := 63
//				dst[i+63:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+63:i] := dst[i+63:i] + 1
//				OD
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:128] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm_maskz_lzcnt_epi64'.
// Requires AVX512CD.
func MaskzLzcntEpi64(k x86.Mmask8, a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// M256LzcntEpi64: Counts the number of leading zero bits in each packed 64-bit
// integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*64
//			tmp := 63
//			dst[i+63:i] := 0
//			DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//				tmp := tmp - 1
//				dst[i+63:i] := dst[i+63:i] + 1
//			OD
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm256_lzcnt_epi64'.
// Requires AVX512CD.
func M256LzcntEpi64(a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskLzcntEpi64: Counts the number of leading zero bits in each packed
// 64-bit integer in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp := 63
//				dst[i+63:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+63:i] := dst[i+63:i] + 1
//				OD
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm256_mask_lzcnt_epi64'.
// Requires AVX512CD.
func M256MaskLzcntEpi64(src x86.M256i, k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M256MaskzLzcntEpi64: Counts the number of leading zero bits in each packed
// 64-bit integer in 'a', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 3
//			i := j*64
//			IF k[j]
//				tmp := 63
//				dst[i+63:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+63:i] := dst[i+63:i] + 1
//				OD
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:256] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm256_maskz_lzcnt_epi64'.
// Requires AVX512CD.
func M256MaskzLzcntEpi64(k x86.Mmask8, a x86.M256i) (dst x86.M256i) {
	panic("not implemented")
}


// M512LzcntEpi64: Counts the number of leading zero bits in each packed 64-bit
// integer in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			tmp := 63
//			dst[i+63:i] := 0
//			DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//				tmp := tmp - 1
//				dst[i+63:i] := dst[i+63:i] + 1
//			OD
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm512_lzcnt_epi64'.
// Requires AVX512CD.
func M512LzcntEpi64(a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskLzcntEpi64: Counts the number of leading zero bits in each packed
// 64-bit integer in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp := 63
//				dst[i+63:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+63:i] := dst[i+63:i] + 1
//				OD
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm512_mask_lzcnt_epi64'.
// Requires AVX512CD.
func M512MaskLzcntEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskzLzcntEpi64: Counts the number of leading zero bits in each packed
// 64-bit integer in 'a', and store the results in 'dst' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				tmp := 63
//				dst[i+63:i] := 0
//				DO WHILE (tmp >= 0 AND a[i+tmp] == 0)
//					tmp := tmp - 1
//					dst[i+63:i] := dst[i+63:i] + 1
//				OD
//			ELSE
//				dst[i+63:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPLZCNTQ'. Intrinsic: '_mm512_maskz_lzcnt_epi64'.
// Requires AVX512CD.
func M512MaskzLzcntEpi64(k x86.Mmask8, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}

