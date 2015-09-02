package avx512cd

import . "github.com/klauspost/intrinsics/x86"


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
func BroadcastmbEpi64(k Mmask8) M128i {
	return M128i(broadcastmbEpi64(uint8(k)))
}

func broadcastmbEpi64(k uint8) [16]byte


// BroadcastmbEpi641: Broadcast the low 8-bits from input mask 'k' to all
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
func BroadcastmbEpi641(k Mmask8) M256i {
	return M256i(broadcastmbEpi641(uint8(k)))
}

func broadcastmbEpi641(k uint8) [32]byte


// BroadcastmbEpi642: Broadcast the low 8-bits from input mask 'k' to all
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
func BroadcastmbEpi642(k Mmask8) M512i {
	return M512i(broadcastmbEpi642(uint8(k)))
}

func broadcastmbEpi642(k uint8) [64]byte


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
func BroadcastmwEpi32(k Mmask16) M128i {
	return M128i(broadcastmwEpi32(uint16(k)))
}

func broadcastmwEpi32(k uint16) [16]byte


// BroadcastmwEpi321: Broadcast the low 16-bits from input mask 'k' to all
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
func BroadcastmwEpi321(k Mmask16) M256i {
	return M256i(broadcastmwEpi321(uint16(k)))
}

func broadcastmwEpi321(k uint16) [32]byte


// BroadcastmwEpi322: Broadcast the low 16-bits from input mask 'k' to all
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
func BroadcastmwEpi322(k Mmask16) M512i {
	return M512i(broadcastmwEpi322(uint16(k)))
}

func broadcastmwEpi322(k uint16) [64]byte


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
func ConflictEpi32(a M128i) M128i {
	return M128i(conflictEpi32([16]byte(a)))
}

func conflictEpi32(a [16]byte) [16]byte


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
func MaskConflictEpi32(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskConflictEpi32([16]byte(src), uint8(k), [16]byte(a)))
}

func maskConflictEpi32(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzConflictEpi32(k Mmask8, a M128i) M128i {
	return M128i(maskzConflictEpi32(uint8(k), [16]byte(a)))
}

func maskzConflictEpi32(k uint8, a [16]byte) [16]byte


// ConflictEpi321: Test each 32-bit element of 'a' for equality with all other
// elements in 'a' closer to the least significant bit. Each element's
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
func ConflictEpi321(a M256i) M256i {
	return M256i(conflictEpi321([32]byte(a)))
}

func conflictEpi321(a [32]byte) [32]byte


// MaskConflictEpi321: Test each 32-bit element of 'a' for equality with all
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
func MaskConflictEpi321(src M256i, k Mmask8, a M256i) M256i {
	return M256i(maskConflictEpi321([32]byte(src), uint8(k), [32]byte(a)))
}

func maskConflictEpi321(src [32]byte, k uint8, a [32]byte) [32]byte


// MaskzConflictEpi321: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). Each
// element's comparison forms a zero extended bit vector in 'dst'. 
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
func MaskzConflictEpi321(k Mmask8, a M256i) M256i {
	return M256i(maskzConflictEpi321(uint8(k), [32]byte(a)))
}

func maskzConflictEpi321(k uint8, a [32]byte) [32]byte


// ConflictEpi322: Test each 32-bit element of 'a' for equality with all other
// elements in 'a' closer to the least significant bit. Each element's
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
func ConflictEpi322(a M512i) M512i {
	return M512i(conflictEpi322([64]byte(a)))
}

func conflictEpi322(a [64]byte) [64]byte


// MaskConflictEpi322: Test each 32-bit element of 'a' for equality with all
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
func MaskConflictEpi322(src M512i, k Mmask16, a M512i) M512i {
	return M512i(maskConflictEpi322([64]byte(src), uint16(k), [64]byte(a)))
}

func maskConflictEpi322(src [64]byte, k uint16, a [64]byte) [64]byte


// MaskzConflictEpi322: Test each 32-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). Each
// element's comparison forms a zero extended bit vector in 'dst'. 
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
func MaskzConflictEpi322(k Mmask16, a M512i) M512i {
	return M512i(maskzConflictEpi322(uint16(k), [64]byte(a)))
}

func maskzConflictEpi322(k uint16, a [64]byte) [64]byte


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
func ConflictEpi64(a M128i) M128i {
	return M128i(conflictEpi64([16]byte(a)))
}

func conflictEpi64(a [16]byte) [16]byte


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
func MaskConflictEpi64(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskConflictEpi64([16]byte(src), uint8(k), [16]byte(a)))
}

func maskConflictEpi64(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzConflictEpi64(k Mmask8, a M128i) M128i {
	return M128i(maskzConflictEpi64(uint8(k), [16]byte(a)))
}

func maskzConflictEpi64(k uint8, a [16]byte) [16]byte


// ConflictEpi641: Test each 64-bit element of 'a' for equality with all other
// elements in 'a' closer to the least significant bit. Each element's
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
func ConflictEpi641(a M256i) M256i {
	return M256i(conflictEpi641([32]byte(a)))
}

func conflictEpi641(a [32]byte) [32]byte


// MaskConflictEpi641: Test each 64-bit element of 'a' for equality with all
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
func MaskConflictEpi641(src M256i, k Mmask8, a M256i) M256i {
	return M256i(maskConflictEpi641([32]byte(src), uint8(k), [32]byte(a)))
}

func maskConflictEpi641(src [32]byte, k uint8, a [32]byte) [32]byte


// MaskzConflictEpi641: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). Each
// element's comparison forms a zero extended bit vector in 'dst'. 
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
func MaskzConflictEpi641(k Mmask8, a M256i) M256i {
	return M256i(maskzConflictEpi641(uint8(k), [32]byte(a)))
}

func maskzConflictEpi641(k uint8, a [32]byte) [32]byte


// ConflictEpi642: Test each 64-bit element of 'a' for equality with all other
// elements in 'a' closer to the least significant bit. Each element's
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
func ConflictEpi642(a M512i) M512i {
	return M512i(conflictEpi642([64]byte(a)))
}

func conflictEpi642(a [64]byte) [64]byte


// MaskConflictEpi642: Test each 64-bit element of 'a' for equality with all
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
func MaskConflictEpi642(src M512i, k Mmask8, a M512i) M512i {
	return M512i(maskConflictEpi642([64]byte(src), uint8(k), [64]byte(a)))
}

func maskConflictEpi642(src [64]byte, k uint8, a [64]byte) [64]byte


// MaskzConflictEpi642: Test each 64-bit element of 'a' for equality with all
// other elements in 'a' closer to the least significant bit using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). Each
// element's comparison forms a zero extended bit vector in 'dst'. 
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
func MaskzConflictEpi642(k Mmask8, a M512i) M512i {
	return M512i(maskzConflictEpi642(uint8(k), [64]byte(a)))
}

func maskzConflictEpi642(k uint8, a [64]byte) [64]byte


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
func LzcntEpi32(a M128i) M128i {
	return M128i(lzcntEpi32([16]byte(a)))
}

func lzcntEpi32(a [16]byte) [16]byte


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
func MaskLzcntEpi32(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskLzcntEpi32([16]byte(src), uint8(k), [16]byte(a)))
}

func maskLzcntEpi32(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzLzcntEpi32(k Mmask8, a M128i) M128i {
	return M128i(maskzLzcntEpi32(uint8(k), [16]byte(a)))
}

func maskzLzcntEpi32(k uint8, a [16]byte) [16]byte


// LzcntEpi321: Counts the number of leading zero bits in each packed 32-bit
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
func LzcntEpi321(a M256i) M256i {
	return M256i(lzcntEpi321([32]byte(a)))
}

func lzcntEpi321(a [32]byte) [32]byte


// MaskLzcntEpi321: Counts the number of leading zero bits in each packed
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
func MaskLzcntEpi321(src M256i, k Mmask8, a M256i) M256i {
	return M256i(maskLzcntEpi321([32]byte(src), uint8(k), [32]byte(a)))
}

func maskLzcntEpi321(src [32]byte, k uint8, a [32]byte) [32]byte


// MaskzLzcntEpi321: Counts the number of leading zero bits in each packed
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
func MaskzLzcntEpi321(k Mmask8, a M256i) M256i {
	return M256i(maskzLzcntEpi321(uint8(k), [32]byte(a)))
}

func maskzLzcntEpi321(k uint8, a [32]byte) [32]byte


// LzcntEpi322: Counts the number of leading zero bits in each packed 32-bit
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
func LzcntEpi322(a M512i) M512i {
	return M512i(lzcntEpi322([64]byte(a)))
}

func lzcntEpi322(a [64]byte) [64]byte


// MaskLzcntEpi322: Counts the number of leading zero bits in each packed
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
func MaskLzcntEpi322(src M512i, k Mmask16, a M512i) M512i {
	return M512i(maskLzcntEpi322([64]byte(src), uint16(k), [64]byte(a)))
}

func maskLzcntEpi322(src [64]byte, k uint16, a [64]byte) [64]byte


// MaskzLzcntEpi322: Counts the number of leading zero bits in each packed
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
func MaskzLzcntEpi322(k Mmask16, a M512i) M512i {
	return M512i(maskzLzcntEpi322(uint16(k), [64]byte(a)))
}

func maskzLzcntEpi322(k uint16, a [64]byte) [64]byte


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
func LzcntEpi64(a M128i) M128i {
	return M128i(lzcntEpi64([16]byte(a)))
}

func lzcntEpi64(a [16]byte) [16]byte


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
func MaskLzcntEpi64(src M128i, k Mmask8, a M128i) M128i {
	return M128i(maskLzcntEpi64([16]byte(src), uint8(k), [16]byte(a)))
}

func maskLzcntEpi64(src [16]byte, k uint8, a [16]byte) [16]byte


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
func MaskzLzcntEpi64(k Mmask8, a M128i) M128i {
	return M128i(maskzLzcntEpi64(uint8(k), [16]byte(a)))
}

func maskzLzcntEpi64(k uint8, a [16]byte) [16]byte


// LzcntEpi641: Counts the number of leading zero bits in each packed 64-bit
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
func LzcntEpi641(a M256i) M256i {
	return M256i(lzcntEpi641([32]byte(a)))
}

func lzcntEpi641(a [32]byte) [32]byte


// MaskLzcntEpi641: Counts the number of leading zero bits in each packed
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
func MaskLzcntEpi641(src M256i, k Mmask8, a M256i) M256i {
	return M256i(maskLzcntEpi641([32]byte(src), uint8(k), [32]byte(a)))
}

func maskLzcntEpi641(src [32]byte, k uint8, a [32]byte) [32]byte


// MaskzLzcntEpi641: Counts the number of leading zero bits in each packed
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
func MaskzLzcntEpi641(k Mmask8, a M256i) M256i {
	return M256i(maskzLzcntEpi641(uint8(k), [32]byte(a)))
}

func maskzLzcntEpi641(k uint8, a [32]byte) [32]byte


// LzcntEpi642: Counts the number of leading zero bits in each packed 64-bit
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
func LzcntEpi642(a M512i) M512i {
	return M512i(lzcntEpi642([64]byte(a)))
}

func lzcntEpi642(a [64]byte) [64]byte


// MaskLzcntEpi642: Counts the number of leading zero bits in each packed
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
func MaskLzcntEpi642(src M512i, k Mmask8, a M512i) M512i {
	return M512i(maskLzcntEpi642([64]byte(src), uint8(k), [64]byte(a)))
}

func maskLzcntEpi642(src [64]byte, k uint8, a [64]byte) [64]byte


// MaskzLzcntEpi642: Counts the number of leading zero bits in each packed
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
func MaskzLzcntEpi642(k Mmask8, a M512i) M512i {
	return M512i(maskzLzcntEpi642(uint8(k), [64]byte(a)))
}

func maskzLzcntEpi642(k uint8, a [64]byte) [64]byte

