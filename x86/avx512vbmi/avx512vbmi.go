package avx512vbmi

import . "github.com/klauspost/intrinsics/x86"


// MaskMultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR i := 0 to 7
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				IF k[i*8+j]
//					dst[q+j*8+7:q+j*8] := tmp8[7:0]
//				ELSE
//					dst[q+j*8+7:q+j*8] := src[q+j*8+7:q+j*8]
//				FI
//			ENDFOR
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm512_mask_multishift_epi64_epi8'.
// Requires AVX512VBMI.
func MaskMultishiftEpi64Epi8(src M512i, k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskMultishiftEpi64Epi8([64]byte(src), uint64(k), [64]byte(a), [64]byte(b)))
}

func maskMultishiftEpi64Epi8(src [64]byte, k uint64, a [64]byte, b [64]byte) [64]byte


// MaskzMultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR i := 0 to 7
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				IF k[i*8+j]
//					dst[q+j*8+7:q+j*8] := tmp8[7:0]
//				ELSE
//					dst[q+j*8+7:q+j*8] := 0
//				FI
//			ENDFOR
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm512_maskz_multishift_epi64_epi8'.
// Requires AVX512VBMI.
func MaskzMultishiftEpi64Epi8(k Mmask64, a M512i, b M512i) M512i {
	return M512i(maskzMultishiftEpi64Epi8(uint64(k), [64]byte(a), [64]byte(b)))
}

func maskzMultishiftEpi64Epi8(k uint64, a [64]byte, b [64]byte) [64]byte


// MultishiftEpi64Epi8: For each 64-bit element in 'b', select 8 unaligned
// bytes using a byte-granular shift control within the corresponding 64-bit
// element of 'a', and store the 8 assembled bytes to the corresponding 64-bit
// element of 'dst'. 
//
//		FOR i := 0 to 7
//			q := i * 64
//			FOR j := 0 to 7
//				tmp8 := 0
//				ctrl := a[q+j*8+7:q+j*8] & 63
//				FOR l := 0 to 7
//					tmp8[k] := b[q+((ctrl+k) & 63)]
//				ENDFOR
//				dst[q+j*8+7:q+j*8] := tmp8[7:0]
//			ENDFOR
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULTISHIFTQB'. Intrinsic: '_mm512_multishift_epi64_epi8'.
// Requires AVX512VBMI.
func MultishiftEpi64Epi8(a M512i, b M512i) M512i {
	return M512i(multishiftEpi64Epi8([64]byte(a), [64]byte(b)))
}

func multishiftEpi64Epi8(a [64]byte, b [64]byte) [64]byte


// MaskPermutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+5:i]
//				dst[i+7:i] := idx[i+6] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMT2B'. Intrinsic: '_mm512_mask_permutex2var_epi8'.
// Requires AVX512VBMI.
func MaskPermutex2varEpi8(a M512i, k Mmask64, idx M512i, b M512i) M512i {
	return M512i(maskPermutex2varEpi8([64]byte(a), uint64(k), [64]byte(idx), [64]byte(b)))
}

func maskPermutex2varEpi8(a [64]byte, k uint64, idx [64]byte, b [64]byte) [64]byte


// Mask2Permutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+5:i]
//				dst[i+7:i] := idx[i+6] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMI2B'. Intrinsic: '_mm512_mask2_permutex2var_epi8'.
// Requires AVX512VBMI.
func Mask2Permutex2varEpi8(a M512i, idx M512i, k Mmask64, b M512i) M512i {
	return M512i(mask2Permutex2varEpi8([64]byte(a), [64]byte(idx), uint64(k), [64]byte(b)))
}

func mask2Permutex2varEpi8(a [64]byte, idx [64]byte, k uint64, b [64]byte) [64]byte


// MaskzPermutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across lanes
// using the corresponding selector and index in 'idx', and store the results
// in 'dst' using zeromask 'k' (elements are zeroed out when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			IF k[j]
//				off := 8*idx[i+5:i]
//				dst[i+7:i] := idx[i+6] ? b[off+7:off] : a[off+7:off]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMI2B, VPERMT2B'. Intrinsic: '_mm512_maskz_permutex2var_epi8'.
// Requires AVX512VBMI.
func MaskzPermutex2varEpi8(k Mmask64, a M512i, idx M512i, b M512i) M512i {
	return M512i(maskzPermutex2varEpi8(uint64(k), [64]byte(a), [64]byte(idx), [64]byte(b)))
}

func maskzPermutex2varEpi8(k uint64, a [64]byte, idx [64]byte, b [64]byte) [64]byte


// Permutex2varEpi8: Shuffle 8-bit integers in 'a' and 'b' across lanes using
// the corresponding selector and index in 'idx', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			off := 8*idx[i+5:i]
//			dst[i+7:i] := idx[i+6] ? b[off+7:off] : a[off+7:off]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMI2B'. Intrinsic: '_mm512_permutex2var_epi8'.
// Requires AVX512VBMI.
func Permutex2varEpi8(a M512i, idx M512i, b M512i) M512i {
	return M512i(permutex2varEpi8([64]byte(a), [64]byte(idx), [64]byte(b)))
}

func permutex2varEpi8(a [64]byte, idx [64]byte, b [64]byte) [64]byte


// MaskPermutexvarEpi8: Shuffle 8-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			id := idx[i+5:i]*8
//			IF k[j]
//				dst[i+7:i] := a[id+7:id]
//			ELSE
//				dst[i+7:i] := src[i+7:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm512_mask_permutexvar_epi8'.
// Requires AVX512VBMI.
func MaskPermutexvarEpi8(src M512i, k Mmask64, idx M512i, a M512i) M512i {
	return M512i(maskPermutexvarEpi8([64]byte(src), uint64(k), [64]byte(idx), [64]byte(a)))
}

func maskPermutexvarEpi8(src [64]byte, k uint64, idx [64]byte, a [64]byte) [64]byte


// MaskzPermutexvarEpi8: Shuffle 8-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 63
//			i := j*8
//			id := idx[i+5:i]*8
//			IF k[j]
//				dst[i+7:i] := a[id+7:id]
//			ELSE
//				dst[i+7:i] := 0
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm512_maskz_permutexvar_epi8'.
// Requires AVX512VBMI.
func MaskzPermutexvarEpi8(k Mmask64, idx M512i, a M512i) M512i {
	return M512i(maskzPermutexvarEpi8(uint64(k), [64]byte(idx), [64]byte(a)))
}

func maskzPermutexvarEpi8(k uint64, idx [64]byte, a [64]byte) [64]byte


// PermutexvarEpi8: Shuffle 8-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst'. 
//
//		FOR j := 0 to 63
//			i := j*8
//			id := idx[i+5:i]*8
//			dst[i+7:i] := a[id+7:id]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMB'. Intrinsic: '_mm512_permutexvar_epi8'.
// Requires AVX512VBMI.
func PermutexvarEpi8(idx M512i, a M512i) M512i {
	return M512i(permutexvarEpi8([64]byte(idx), [64]byte(a)))
}

func permutexvarEpi8(idx [64]byte, a [64]byte) [64]byte

