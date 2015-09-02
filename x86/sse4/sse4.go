package sse4

import "github.com/klauspost/intrinsics/x86"


// BlendEpi16: Blend packed 16-bit integers from 'a' and 'b' using control mask
// 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF imm8[j%8]
//				dst[i+15:i] := b[i+15:i]
//			ELSE
//				dst[i+15:i] := a[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PBLENDW'. Intrinsic: '_mm_blend_epi16'.
// Requires SSE4.1.
func BlendEpi16(a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(blendEpi16([16]byte(a), [16]byte(b), imm8))
}

func blendEpi16(a [16]byte, b [16]byte, imm8 int) [16]byte


// BlendPd: Blend packed double-precision (64-bit) floating-point elements from
// 'a' and 'b' using control mask 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF imm8[j%8]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'BLENDPD'. Intrinsic: '_mm_blend_pd'.
// Requires SSE4.1.
func BlendPd(a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(blendPd([2]float64(a), [2]float64(b), imm8))
}

func blendPd(a [2]float64, b [2]float64, imm8 int) [2]float64


// BlendPs: Blend packed single-precision (32-bit) floating-point elements from
// 'a' and 'b' using control mask 'imm8', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF imm8[j%8]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'BLENDPS'. Intrinsic: '_mm_blend_ps'.
// Requires SSE4.1.
func BlendPs(a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(blendPs([4]float32(a), [4]float32(b), imm8))
}

func blendPs(a [4]float32, b [4]float32, imm8 int) [4]float32


// BlendvEpi8: Blend packed 8-bit integers from 'a' and 'b' using 'mask', and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF mask[i+7]
//				dst[i+7:i] := b[i+7:i]
//			ELSE
//				dst[i+7:i] := a[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PBLENDVB'. Intrinsic: '_mm_blendv_epi8'.
// Requires SSE4.1.
func BlendvEpi8(a x86.M128i, b x86.M128i, mask x86.M128i) x86.M128i {
	return x86.M128i(blendvEpi8([16]byte(a), [16]byte(b), [16]byte(mask)))
}

func blendvEpi8(a [16]byte, b [16]byte, mask [16]byte) [16]byte


// BlendvPd: Blend packed double-precision (64-bit) floating-point elements
// from 'a' and 'b' using 'mask', and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF mask[i+63]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'BLENDVPD'. Intrinsic: '_mm_blendv_pd'.
// Requires SSE4.1.
func BlendvPd(a x86.M128d, b x86.M128d, mask x86.M128d) x86.M128d {
	return x86.M128d(blendvPd([2]float64(a), [2]float64(b), [2]float64(mask)))
}

func blendvPd(a [2]float64, b [2]float64, mask [2]float64) [2]float64


// BlendvPs: Blend packed single-precision (32-bit) floating-point elements
// from 'a' and 'b' using 'mask', and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF mask[i+31]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'BLENDVPS'. Intrinsic: '_mm_blendv_ps'.
// Requires SSE4.1.
func BlendvPs(a x86.M128, b x86.M128, mask x86.M128) x86.M128 {
	return x86.M128(blendvPs([4]float32(a), [4]float32(b), [4]float32(mask)))
}

func blendvPs(a [4]float32, b [4]float32, mask [4]float32) [4]float32


// CeilPd: Round the packed double-precision (64-bit) floating-point elements
// in 'a' up to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := CEIL(a[i+63:i])
//		ENDFOR
//
// Instruction: 'ROUNDPD'. Intrinsic: '_mm_ceil_pd'.
// Requires SSE4.1.
func CeilPd(a x86.M128d) x86.M128d {
	return x86.M128d(ceilPd([2]float64(a)))
}

func ceilPd(a [2]float64) [2]float64


// CeilPs: Round the packed single-precision (32-bit) floating-point elements
// in 'a' up to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := CEIL(a[i+31:i])
//		ENDFOR
//
// Instruction: 'ROUNDPS'. Intrinsic: '_mm_ceil_ps'.
// Requires SSE4.1.
func CeilPs(a x86.M128) x86.M128 {
	return x86.M128(ceilPs([4]float32(a)))
}

func ceilPs(a [4]float32) [4]float32


// CeilSd: Round the lower double-precision (64-bit) floating-point element in
// 'b' up to an integer value, store the result as a double-precision
// floating-point element in the lower element of 'dst', and copy the upper
// element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := CEIL(b[63:0])
//		dst[127:64] := a[127:64]
//
// Instruction: 'ROUNDSD'. Intrinsic: '_mm_ceil_sd'.
// Requires SSE4.1.
func CeilSd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(ceilSd([2]float64(a), [2]float64(b)))
}

func ceilSd(a [2]float64, b [2]float64) [2]float64


// CeilSs: Round the lower single-precision (32-bit) floating-point element in
// 'b' up to an integer value, store the result as a single-precision
// floating-point element in the lower element of 'dst', and copy the upper 3
// packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := CEIL(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'ROUNDSS'. Intrinsic: '_mm_ceil_ss'.
// Requires SSE4.1.
func CeilSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(ceilSs([4]float32(a), [4]float32(b)))
}

func ceilSs(a [4]float32, b [4]float32) [4]float32


// CmpeqEpi64: Compare packed 64-bit integers in 'a' and 'b' for equality, and
// store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ( a[i+63:i] == b[i+63:i] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPEQQ'. Intrinsic: '_mm_cmpeq_epi64'.
// Requires SSE4.1.
func CmpeqEpi64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(cmpeqEpi64([16]byte(a), [16]byte(b)))
}

func cmpeqEpi64(a [16]byte, b [16]byte) [16]byte


// Cmpestra: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if 'b' did not contain a null
// character and the resulting mask was zero, and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//					0:  // equal any
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					1:  // ranges
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					2:  // equal each
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//					3:  // equal ordered
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 1
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//			0:  // equal any
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound
//						IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//					ENDFOR
//				ENDFOR
//			1:  // ranges
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound, j += 2
//						IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//					ENDFOR
//				ENDFOR
//			2:  // equal each
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					IntRes1[i] := BoolRes[i][i]
//				ENDFOR
//			3:  // equal ordered
//				IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//				FOR i := 0 to UpperBound
//					k := i
//					FOR j := 0 to UpperBound-i
//						IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//						k++
//					ENDFOR
//				ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 == 0) AND (lb > UpperBound)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestra'.
// Requires SSE4.2.
func Cmpestra(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) int {
	return int(cmpestra([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestra(a [16]byte, la int, b [16]byte, lb int, imm8 int) int


// Cmpestrc: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if the resulting mask was
// non-zero, and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//					0:  // equal any
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					1:  // ranges
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					2:  // equal each
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//					3:  // equal ordered
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 1
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//			0:  // equal any
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound
//						IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//					ENDFOR
//				ENDFOR
//			1:  // ranges
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound, j += 2
//						IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//					ENDFOR
//				ENDFOR
//			2:  // equal each
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					IntRes1[i] := BoolRes[i][i]
//				ENDFOR
//			3:  // equal ordered
//				IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//				FOR i := 0 to UpperBound
//					k := i
//					FOR j := 0 to UpperBound-i
//						IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//						k++
//					ENDFOR
//				ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 != 0)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestrc'.
// Requires SSE4.2.
func Cmpestrc(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) int {
	return int(cmpestrc([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestrc(a [16]byte, la int, b [16]byte, lb int, imm8 int) int


// Cmpestri: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and store the generated index in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // most significant bit
//			tmp := UpperBound
//			dst := tmp
//			DO WHILE ((tmp >= 0) AND a[tmp] = 0)
//				tmp := tmp - 1
//				dst := tmp
//			OD
//		ELSE // least significant bit
//			tmp := 0
//			dst := tmp
//			DO WHILE ((tmp <= UpperBound) AND a[tmp] = 0)
//				tmp := tmp + 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestri'.
// Requires SSE4.2.
func Cmpestri(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) int {
	return int(cmpestri([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestri(a [16]byte, la int, b [16]byte, lb int, imm8 int) int


// Cmpestrm: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and store the generated mask in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // byte / word mask
//			FOR i := 0 to UpperBound
//				j := i*size
//				IF IntRes2[i]
//					dst[j+size-1:j] := (imm8[0] ? 0xFF : 0xFFFF)
//				ELSE
//					dst[j+size-1:j] := 0
//				FI
//			ENDFOR
//		ELSE // bit mask
//			dst[UpperBound:0] := IntRes[UpperBound:0]
//			dst[127:UpperBound+1] := 0
//		FI
//
// Instruction: 'PCMPESTRM'. Intrinsic: '_mm_cmpestrm'.
// Requires SSE4.2.
func Cmpestrm(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) x86.M128i {
	return x86.M128i(cmpestrm([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestrm(a [16]byte, la int, b [16]byte, lb int, imm8 int) [16]byte


// Cmpestro: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns bit 0 of the resulting bit mask.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//					0:  // equal any
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					1:  // ranges
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					2:  // equal each
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//					3:  // equal ordered
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 1
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//			0:  // equal any
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound
//						IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//					ENDFOR
//				ENDFOR
//			1:  // ranges
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound, j += 2
//						IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//					ENDFOR
//				ENDFOR
//			2:  // equal each
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					IntRes1[i] := BoolRes[i][i]
//				ENDFOR
//			3:  // equal ordered
//				IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//				FOR i := 0 to UpperBound
//					k := i
//					FOR j := 0 to UpperBound-i
//						IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//						k++
//					ENDFOR
//				ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := IntRes2[0
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestro'.
// Requires SSE4.2.
func Cmpestro(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) int {
	return int(cmpestro([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestro(a [16]byte, la int, b [16]byte, lb int, imm8 int) int


// Cmpestrs: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if any character in 'a' was null,
// and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		dst := (la <= UpperBound)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestrs'.
// Requires SSE4.2.
func Cmpestrs(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) int {
	return int(cmpestrs([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestrs(a [16]byte, la int, b [16]byte, lb int, imm8 int) int


// Cmpestrz: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if any character in 'b' was null,
// and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		dst := (lb <= UpperBound)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestrz'.
// Requires SSE4.2.
func Cmpestrz(a x86.M128i, la int, b x86.M128i, lb int, imm8 int) int {
	return int(cmpestrz([16]byte(a), la, [16]byte(b), lb, imm8))
}

func cmpestrz(a [16]byte, la int, b [16]byte, lb int, imm8 int) int


// CmpgtEpi64: Compare packed 64-bit integers in 'a' and 'b' for greater-than,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ( a[i+63:i] > b[i+63:i] ) ? 0xFFFFFFFFFFFFFFFF : 0
//		ENDFOR
//
// Instruction: 'PCMPGTQ'. Intrinsic: '_mm_cmpgt_epi64'.
// Requires SSE4.2.
func CmpgtEpi64(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(cmpgtEpi64([16]byte(a), [16]byte(b)))
}

func cmpgtEpi64(a [16]byte, b [16]byte) [16]byte


// Cmpistra: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if 'b' did not contain a null character
// and the resulting mask was zero, and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 == 0) AND bInvalid
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistra'.
// Requires SSE4.2.
func Cmpistra(a x86.M128i, b x86.M128i, imm8 int) int {
	return int(cmpistra([16]byte(a), [16]byte(b), imm8))
}

func cmpistra(a [16]byte, b [16]byte, imm8 int) int


// Cmpistrc: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if the resulting mask was non-zero, and
// 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 != 0)
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistrc'.
// Requires SSE4.2.
func Cmpistrc(a x86.M128i, b x86.M128i, imm8 int) int {
	return int(cmpistrc([16]byte(a), [16]byte(b), imm8))
}

func cmpistrc(a [16]byte, b [16]byte, imm8 int) int


// Cmpistri: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and store the generated index in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // most significant bit
//			tmp := UpperBound
//			dst := tmp
//			DO WHILE ((tmp >= 0) AND a[tmp] = 0)
//				tmp := tmp - 1
//				dst := tmp
//			OD
//		ELSE // least significant bit
//			tmp := 0
//			dst := tmp
//			DO WHILE ((tmp <= UpperBound) AND a[tmp] = 0)
//				tmp := tmp + 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistri'.
// Requires SSE4.2.
func Cmpistri(a x86.M128i, b x86.M128i, imm8 int) int {
	return int(cmpistri([16]byte(a), [16]byte(b), imm8))
}

func cmpistri(a [16]byte, b [16]byte, imm8 int) int


// Cmpistrm: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and store the generated mask in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // byte / word mask
//			FOR i := 0 to UpperBound
//				j := i*size
//				IF IntRes2[i]
//					dst[j+size-1:j] := (imm8[0] ? 0xFF : 0xFFFF)
//				ELSE
//					dst[j+size-1:j] := 0
//				FI
//			ENDFOR
//		ELSE // bit mask
//			dst[UpperBound:0] := IntRes[UpperBound:0]
//			dst[127:UpperBound+1] := 0
//		FI
//
// Instruction: 'PCMPISTRM'. Intrinsic: '_mm_cmpistrm'.
// Requires SSE4.2.
func Cmpistrm(a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(cmpistrm([16]byte(a), [16]byte(b), imm8))
}

func cmpistrm(a [16]byte, b [16]byte, imm8 int) [16]byte


// Cmpistro: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns bit 0 of the resulting bit mask.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := IntRes2[0]
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistro'.
// Requires SSE4.2.
func Cmpistro(a x86.M128i, b x86.M128i, imm8 int) int {
	return int(cmpistro([16]byte(a), [16]byte(b), imm8))
}

func cmpistro(a [16]byte, b [16]byte, imm8 int) int


// Cmpistrs: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if any character in 'a' was null, and 0
// otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		aInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			IF b[m+size-1:m] == 0
//				aInvalid := 1
//			FI
//		ENDFOR
//		
//		dst := aInvalid
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistrs'.
// Requires SSE4.2.
func Cmpistrs(a x86.M128i, b x86.M128i, imm8 int) int {
	return int(cmpistrs([16]byte(a), [16]byte(b), imm8))
}

func cmpistrs(a [16]byte, b [16]byte, imm8 int) int


// Cmpistrz: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if any character in 'b' was null, and 0
// otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		bInvalid := 0
//		FOR j := 0 to UpperBound
//			n := j*size
//			IF b[n+size-1:n] == 0
//				bInvalid := 1
//			FI
//		ENDFOR
//		
//		dst := bInvalid
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistrz'.
// Requires SSE4.2.
func Cmpistrz(a x86.M128i, b x86.M128i, imm8 int) int {
	return int(cmpistrz([16]byte(a), [16]byte(b), imm8))
}

func cmpistrz(a [16]byte, b [16]byte, imm8 int) int


// Crc32U16: Starting with the initial value in 'crc', accumulates a CRC32
// value for unsigned 16-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[15:0] := v[0:15] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[47:0] := tmp1[15:0] << 32
//		tmp4[47:0] := tmp2[31:0] << 16
//		tmp5[47:0] := tmp3[47:0] XOR tmp4[47:0]
//		tmp6[31:0] := tmp5[47:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u16'.
// Requires SSE4.2.
func Crc32U16(crc uint32, v uint16) uint32 {
	return uint32(crc32U16(crc, v))
}

func crc32U16(crc uint32, v uint16) uint32


// Crc32U32: Starting with the initial value in 'crc', accumulates a CRC32
// value for unsigned 32-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[31:0] := v[0:31] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[63:0] := tmp1[31:0] << 32
//		tmp4[63:0] := tmp2[31:0] << 32
//		tmp5[63:0] := tmp3[63:0] XOR tmp4[63:0]
//		tmp6[31:0] := tmp5[63:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u32'.
// Requires SSE4.2.
func Crc32U32(crc uint32, v uint32) uint32 {
	return uint32(crc32U32(crc, v))
}

func crc32U32(crc uint32, v uint32) uint32


// Crc32U64: Starting with the initial value in 'crc', accumulates a CRC32
// value for unsigned 64-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[63:0] := v[0:63] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[95:0] := tmp1[31:0] << 32
//		tmp4[95:0] := tmp2[63:0] << 64
//		tmp5[95:0] := tmp3[95:0] XOR tmp4[95:0]
//		tmp6[31:0] := tmp5[95:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u64'.
// Requires SSE4.2.
func Crc32U64(crc uint64, v uint64) uint64 {
	return uint64(crc32U64(crc, v))
}

func crc32U64(crc uint64, v uint64) uint64


// Crc32U8: Starting with the initial value in 'crc', accumulates a CRC32 value
// for unsigned 8-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[7:0] := v[0:7] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[39:0] := tmp1[7:0] << 32 
//		tmp4[39:0] := tmp2[31:0] << 8
//		tmp5[39:0] := tmp3[39:0] XOR tmp4[39:0]
//		tmp6[31:0] := tmp5[39:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u8'.
// Requires SSE4.2.
func Crc32U8(crc uint32, v uint8) uint32 {
	return uint32(crc32U8(crc, v))
}

func crc32U8(crc uint32, v uint8) uint32


// Cvtepi16Epi32: Sign extend packed 16-bit integers in 'a' to packed 32-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 16*j
//			dst[i+31:i] := SignExtend(a[k+15:k])
//		ENDFOR
//
// Instruction: 'PMOVSXWD'. Intrinsic: '_mm_cvtepi16_epi32'.
// Requires SSE4.1.
func Cvtepi16Epi32(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi16Epi32([16]byte(a)))
}

func cvtepi16Epi32(a [16]byte) [16]byte


// Cvtepi16Epi64: Sign extend packed 16-bit integers in 'a' to packed 64-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 16*j
//			dst[i+63:i] := SignExtend(a[k+15:k])
//		ENDFOR
//
// Instruction: 'PMOVSXWQ'. Intrinsic: '_mm_cvtepi16_epi64'.
// Requires SSE4.1.
func Cvtepi16Epi64(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi16Epi64([16]byte(a)))
}

func cvtepi16Epi64(a [16]byte) [16]byte


// Cvtepi32Epi64: Sign extend packed 32-bit integers in 'a' to packed 64-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 32*j
//			dst[i+63:i] := SignExtend(a[k+31:k])
//		ENDFOR
//
// Instruction: 'PMOVSXDQ'. Intrinsic: '_mm_cvtepi32_epi64'.
// Requires SSE4.1.
func Cvtepi32Epi64(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi32Epi64([16]byte(a)))
}

func cvtepi32Epi64(a [16]byte) [16]byte


// Cvtepi8Epi16: Sign extend packed 8-bit integers in 'a' to packed 16-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			l := j*16
//			dst[l+15:l] := SignExtend(a[i+7:i])
//		ENDFOR
//
// Instruction: 'PMOVSXBW'. Intrinsic: '_mm_cvtepi8_epi16'.
// Requires SSE4.1.
func Cvtepi8Epi16(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi8Epi16([16]byte(a)))
}

func cvtepi8Epi16(a [16]byte) [16]byte


// Cvtepi8Epi32: Sign extend packed 8-bit integers in 'a' to packed 32-bit
// integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 8*j
//			dst[i+31:i] := SignExtend(a[k+7:k])
//		ENDFOR
//
// Instruction: 'PMOVSXBD'. Intrinsic: '_mm_cvtepi8_epi32'.
// Requires SSE4.1.
func Cvtepi8Epi32(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi8Epi32([16]byte(a)))
}

func cvtepi8Epi32(a [16]byte) [16]byte


// Cvtepi8Epi64: Sign extend packed 8-bit integers in the low 8 bytes of 'a' to
// packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 8*j
//			dst[i+63:i] := SignExtend(a[k+7:k])
//		ENDFOR
//
// Instruction: 'PMOVSXBQ'. Intrinsic: '_mm_cvtepi8_epi64'.
// Requires SSE4.1.
func Cvtepi8Epi64(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepi8Epi64([16]byte(a)))
}

func cvtepi8Epi64(a [16]byte) [16]byte


// Cvtepu16Epi32: Zero extend packed unsigned 16-bit integers in 'a' to packed
// 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 16*j
//			dst[i+31:i] := ZeroExtend(a[k+15:k])
//		ENDFOR
//
// Instruction: 'PMOVZXWD'. Intrinsic: '_mm_cvtepu16_epi32'.
// Requires SSE4.1.
func Cvtepu16Epi32(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepu16Epi32([16]byte(a)))
}

func cvtepu16Epi32(a [16]byte) [16]byte


// Cvtepu16Epi64: Zero extend packed unsigned 16-bit integers in 'a' to packed
// 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 16*j
//			dst[i+63:i] := ZeroExtend(a[k+15:k])
//		ENDFOR
//
// Instruction: 'PMOVZXWQ'. Intrinsic: '_mm_cvtepu16_epi64'.
// Requires SSE4.1.
func Cvtepu16Epi64(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepu16Epi64([16]byte(a)))
}

func cvtepu16Epi64(a [16]byte) [16]byte


// Cvtepu32Epi64: Zero extend packed unsigned 32-bit integers in 'a' to packed
// 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 32*j
//			dst[i+63:i] := ZeroExtend(a[k+31:k])
//		ENDFOR
//
// Instruction: 'PMOVZXDQ'. Intrinsic: '_mm_cvtepu32_epi64'.
// Requires SSE4.1.
func Cvtepu32Epi64(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepu32Epi64([16]byte(a)))
}

func cvtepu32Epi64(a [16]byte) [16]byte


// Cvtepu8Epi16: Zero extend packed unsigned 8-bit integers in 'a' to packed
// 16-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*8
//			l := j*16
//			dst[l+15:l] := ZeroExtend(a[i+7:i])
//		ENDFOR
//
// Instruction: 'PMOVZXBW'. Intrinsic: '_mm_cvtepu8_epi16'.
// Requires SSE4.1.
func Cvtepu8Epi16(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepu8Epi16([16]byte(a)))
}

func cvtepu8Epi16(a [16]byte) [16]byte


// Cvtepu8Epi32: Zero extend packed unsigned 8-bit integers in 'a' to packed
// 32-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := 32*j
//			k := 8*j
//			dst[i+31:i] := ZeroExtend(a[k+7:k])
//		ENDFOR
//
// Instruction: 'PMOVZXBD'. Intrinsic: '_mm_cvtepu8_epi32'.
// Requires SSE4.1.
func Cvtepu8Epi32(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepu8Epi32([16]byte(a)))
}

func cvtepu8Epi32(a [16]byte) [16]byte


// Cvtepu8Epi64: Zero extend packed unsigned 8-bit integers in the low 8 byte
// sof 'a' to packed 64-bit integers, and store the results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := 64*j
//			k := 8*j
//			dst[i+63:i] := ZeroExtend(a[k+7:k])
//		ENDFOR
//
// Instruction: 'PMOVZXBQ'. Intrinsic: '_mm_cvtepu8_epi64'.
// Requires SSE4.1.
func Cvtepu8Epi64(a x86.M128i) x86.M128i {
	return x86.M128i(cvtepu8Epi64([16]byte(a)))
}

func cvtepu8Epi64(a [16]byte) [16]byte


// DpPd: Conditionally multiply the packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' using the high 4 bits in 'imm8', sum
// the four products, and conditionally store the sum in 'dst' using the low 4
// bits of 'imm8'. 
//
//		DP(a[127:0], b[127:0], imm8[7:0]) {
//			FOR j := 0 to 1
//				i := j*64
//				IF imm8[(4+j)%8]]
//					temp[i+63:i] := a[i+63:i] * b[i+63:i]
//				ELSE
//					temp[i+63:i] := 0
//				FI
//			ENDFOR
//			
//			sum[63:0] := temp[127:64] + temp[63:0]
//			
//			FOR j := 0 to 1
//				i := j*64
//				IF imm8[j%8]
//					tmpdst[i+63:i] := sum[63:0]
//				ELSE
//					tmpdst[i+63:i] := 0
//				FI
//			ENDFOR
//			RETURN tmpdst[127:0]
//		}
//		
//		dst[127:0] := DP(a[127:0], b[127:0], imm8[7:0])
//
// Instruction: 'DPPD'. Intrinsic: '_mm_dp_pd'.
// Requires SSE4.1.
func DpPd(a x86.M128d, b x86.M128d, imm8 int) x86.M128d {
	return x86.M128d(dpPd([2]float64(a), [2]float64(b), imm8))
}

func dpPd(a [2]float64, b [2]float64, imm8 int) [2]float64


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
//
// Instruction: 'DPPS'. Intrinsic: '_mm_dp_ps'.
// Requires SSE4.1.
func DpPs(a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(dpPs([4]float32(a), [4]float32(b), imm8))
}

func dpPs(a [4]float32, b [4]float32, imm8 int) [4]float32


// ExtractEpi32: Extract a 32-bit integer from 'a', selected with 'imm8', and
// store the result in 'dst'. 
//
//		dst[31:0] := (a[127:0] >> (imm8[1:0] * 32))[31:0]
//
// Instruction: 'PEXTRD'. Intrinsic: '_mm_extract_epi32'.
// Requires SSE4.1.
func ExtractEpi32(a x86.M128i, imm8 int) int {
	return int(extractEpi32([16]byte(a), imm8))
}

func extractEpi32(a [16]byte, imm8 int) int


// ExtractEpi64: Extract a 64-bit integer from 'a', selected with 'imm8', and
// store the result in 'dst'. 
//
//		dst[63:0] := (a[127:0] >> (imm8[0] * 64))[63:0]
//
// Instruction: 'PEXTRQ'. Intrinsic: '_mm_extract_epi64'.
// Requires SSE4.1.
func ExtractEpi64(a x86.M128i, imm8 int) int64 {
	return int64(extractEpi64([16]byte(a), imm8))
}

func extractEpi64(a [16]byte, imm8 int) int64


// ExtractEpi8: Extract an 8-bit integer from 'a', selected with 'imm8', and
// store the result in the lower element of 'dst'. 
//
//		dst[7:0] := (a[127:0] >> (imm8[3:0] * 8))[7:0]
//		dst[31:8] := 0
//
// Instruction: 'PEXTRB'. Intrinsic: '_mm_extract_epi8'.
// Requires SSE4.1.
func ExtractEpi8(a x86.M128i, imm8 int) int {
	return int(extractEpi8([16]byte(a), imm8))
}

func extractEpi8(a [16]byte, imm8 int) int


// ExtractPs: Extract a single-precision (32-bit) floating-point element from
// 'a', selected with 'imm8', and store the result in 'dst'. 
//
//		dst[31:0] := (a[127:0] >> (imm8[1:0] * 32))[31:0]
//
// Instruction: 'EXTRACTPS'. Intrinsic: '_mm_extract_ps'.
// Requires SSE4.1.
func ExtractPs(a x86.M128, imm8 int) int {
	return int(extractPs([4]float32(a), imm8))
}

func extractPs(a [4]float32, imm8 int) int


// FloorPd: Round the packed double-precision (64-bit) floating-point elements
// in 'a' down to an integer value, and store the results as packed
// double-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := FLOOR(a[i+63:i])
//		ENDFOR
//
// Instruction: 'ROUNDPD'. Intrinsic: '_mm_floor_pd'.
// Requires SSE4.1.
func FloorPd(a x86.M128d) x86.M128d {
	return x86.M128d(floorPd([2]float64(a)))
}

func floorPd(a [2]float64) [2]float64


// FloorPs: Round the packed single-precision (32-bit) floating-point elements
// in 'a' down to an integer value, and store the results as packed
// single-precision floating-point elements in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := FLOOR(a[i+31:i])
//		ENDFOR
//
// Instruction: 'ROUNDPS'. Intrinsic: '_mm_floor_ps'.
// Requires SSE4.1.
func FloorPs(a x86.M128) x86.M128 {
	return x86.M128(floorPs([4]float32(a)))
}

func floorPs(a [4]float32) [4]float32


// FloorSd: Round the lower double-precision (64-bit) floating-point element in
// 'b' down to an integer value, store the result as a double-precision
// floating-point element in the lower element of 'dst', and copy the upper
// element from 'a' to the upper element of 'dst'. 
//
//		dst[63:0] := FLOOR(b[63:0])
//		dst[127:64] := a[127:64]
//
// Instruction: 'ROUNDSD'. Intrinsic: '_mm_floor_sd'.
// Requires SSE4.1.
func FloorSd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(floorSd([2]float64(a), [2]float64(b)))
}

func floorSd(a [2]float64, b [2]float64) [2]float64


// FloorSs: Round the lower single-precision (32-bit) floating-point element in
// 'b' down to an integer value, store the result as a single-precision
// floating-point element in the lower element of 'dst', and copy the upper 3
// packed elements from 'a' to the upper elements of 'dst'. 
//
//		dst[31:0] := FLOOR(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'ROUNDSS'. Intrinsic: '_mm_floor_ss'.
// Requires SSE4.1.
func FloorSs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(floorSs([4]float32(a), [4]float32(b)))
}

func floorSs(a [4]float32, b [4]float32) [4]float32


// InsertEpi32: Copy 'a' to 'dst', and insert the 32-bit integer 'i' into 'dst'
// at the location specified by 'imm8'. 
//
//		dst[127:0] := a[127:0]
//		sel := imm8[1:0]*32
//		dst[sel+31:sel] := i[31:0]
//
// Instruction: 'PINSRD'. Intrinsic: '_mm_insert_epi32'.
// Requires SSE4.1.
func InsertEpi32(a x86.M128i, i int, imm8 int) x86.M128i {
	return x86.M128i(insertEpi32([16]byte(a), i, imm8))
}

func insertEpi32(a [16]byte, i int, imm8 int) [16]byte


// InsertEpi64: Copy 'a' to 'dst', and insert the 64-bit integer 'i' into 'dst'
// at the location specified by 'imm8'. 
//
//		dst[127:0] := a[127:0]
//		sel := imm8[0]*64
//		dst[sel+63:sel] := i[63:0]
//
// Instruction: 'PINSRQ'. Intrinsic: '_mm_insert_epi64'.
// Requires SSE4.1.
func InsertEpi64(a x86.M128i, i int64, imm8 int) x86.M128i {
	return x86.M128i(insertEpi64([16]byte(a), i, imm8))
}

func insertEpi64(a [16]byte, i int64, imm8 int) [16]byte


// InsertEpi8: Copy 'a' to 'dst', and insert the lower 8-bit integer from 'i'
// into 'dst' at the location specified by 'imm8'. 
//
//		dst[127:0] := a[127:0]
//		sel := imm8[3:0]*8
//		dst[sel+7:sel] := i[7:0]
//
// Instruction: 'PINSRB'. Intrinsic: '_mm_insert_epi8'.
// Requires SSE4.1.
func InsertEpi8(a x86.M128i, i int, imm8 int) x86.M128i {
	return x86.M128i(insertEpi8([16]byte(a), i, imm8))
}

func insertEpi8(a [16]byte, i int, imm8 int) [16]byte


// InsertPs: Copy 'a' to 'tmp', then insert a single-precision (32-bit)
// floating-point element from 'b' into 'tmp' using the control in 'imm8'.
// Store 'tmp' to 'dst' using the mask in 'imm8' (elements are zeroed out when
// the corresponding bit is set). 
//
//		tmp2[127:0] := a[127:0]
//		CASE (imm8[7:6]) of
//		0: tmp1[31:0] := b[31:0]
//		1: tmp1[31:0] := b[63:32]
//		2: tmp1[31:0] := b[95:64]
//		3: tmp1[31:0] := b[127:96]
//		ESAC
//		CASE (imm8[5:4]) of
//		0: tmp2[31:0] := tmp1[31:0]
//		1: tmp2[63:32] := tmp1[31:0]
//		2: tmp2[95:64] := tmp1[31:0]
//		3: tmp2[127:96] := tmp1[31:0]
//		ESAC
//		FOR j := 0 to 3
//			i := j*32
//			IF imm8[j%8]
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := tmp2[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'INSERTPS'. Intrinsic: '_mm_insert_ps'.
// Requires SSE4.1.
func InsertPs(a x86.M128, b x86.M128, imm8 int) x86.M128 {
	return x86.M128(insertPs([4]float32(a), [4]float32(b), imm8))
}

func insertPs(a [4]float32, b [4]float32, imm8 int) [4]float32


// MaxEpi32: Compare packed 32-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSD'. Intrinsic: '_mm_max_epi32'.
// Requires SSE4.1.
func MaxEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maxEpi32([16]byte(a), [16]byte(b)))
}

func maxEpi32(a [16]byte, b [16]byte) [16]byte


// MaxEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF a[i+7:i] > b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXSB'. Intrinsic: '_mm_max_epi8'.
// Requires SSE4.1.
func MaxEpi8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maxEpi8([16]byte(a), [16]byte(b)))
}

func maxEpi8(a [16]byte, b [16]byte) [16]byte


// MaxEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF a[i+15:i] > b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUW'. Intrinsic: '_mm_max_epu16'.
// Requires SSE4.1.
func MaxEpu16(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maxEpu16([16]byte(a), [16]byte(b)))
}

func maxEpu16(a [16]byte, b [16]byte) [16]byte


// MaxEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMAXUD'. Intrinsic: '_mm_max_epu32'.
// Requires SSE4.1.
func MaxEpu32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(maxEpu32([16]byte(a), [16]byte(b)))
}

func maxEpu32(a [16]byte, b [16]byte) [16]byte


// MinEpi32: Compare packed 32-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSD'. Intrinsic: '_mm_min_epi32'.
// Requires SSE4.1.
func MinEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(minEpi32([16]byte(a), [16]byte(b)))
}

func minEpi32(a [16]byte, b [16]byte) [16]byte


// MinEpi8: Compare packed 8-bit integers in 'a' and 'b', and store packed
// minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*8
//			IF a[i+7:i] < b[i+7:i]
//				dst[i+7:i] := a[i+7:i]
//			ELSE
//				dst[i+7:i] := b[i+7:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINSB'. Intrinsic: '_mm_min_epi8'.
// Requires SSE4.1.
func MinEpi8(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(minEpi8([16]byte(a), [16]byte(b)))
}

func minEpi8(a [16]byte, b [16]byte) [16]byte


// MinEpu16: Compare packed unsigned 16-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*16
//			IF a[i+15:i] < b[i+15:i]
//				dst[i+15:i] := a[i+15:i]
//			ELSE
//				dst[i+15:i] := b[i+15:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUW'. Intrinsic: '_mm_min_epu16'.
// Requires SSE4.1.
func MinEpu16(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(minEpu16([16]byte(a), [16]byte(b)))
}

func minEpu16(a [16]byte, b [16]byte) [16]byte


// MinEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'PMINUD'. Intrinsic: '_mm_min_epu32'.
// Requires SSE4.1.
func MinEpu32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(minEpu32([16]byte(a), [16]byte(b)))
}

func minEpu32(a [16]byte, b [16]byte) [16]byte


// MinposEpu16: Horizontally compute the minimum amongst the packed unsigned
// 16-bit integers in 'a', store the minimum and index in 'dst', and zero the
// remaining bits in 'dst'. 
//
//		index[2:0] := 0
//		min[15:0] := a[15:0]
//		FOR j := 0 to 7
//			i := j*16
//			IF a[i+15:i] < min[15:0]
//				index[2:0] := j
//				min[15:0] := a[i+15:i]
//			FI
//		ENDFOR
//		dst[15:0] := min[15:0]
//		dst[18:16] := index[2:0]
//		dst[127:19] := 0
//
// Instruction: 'PHMINPOSUW'. Intrinsic: '_mm_minpos_epu16'.
// Requires SSE4.1.
func MinposEpu16(a x86.M128i) x86.M128i {
	return x86.M128i(minposEpu16([16]byte(a)))
}

func minposEpu16(a [16]byte) [16]byte


// MpsadbwEpu8: Compute the sum of absolute differences (SADs) of quadruplets
// of unsigned 8-bit integers in 'a' compared to those in 'b', and store the
// 16-bit results in 'dst'.
// 	Eight SADs are performed using one quadruplet from 'b' and eight
// quadruplets from 'a'. One quadruplet is selected from 'b' starting at on the
// offset specified in 'imm8'. Eight quadruplets are formed from sequential
// 8-bit integers selected from 'a' starting at the offset specified in 'imm8'. 
//
//		MPSADBW(a[127:0], b[127:0], imm8[2:0]) {
//			a_offset := imm8[2]*32
//			b_offset := imm8[1:0]*32
//			FOR j := 0 to 7
//				i := j*8
//				k := a_offset+i
//				l := b_offset
//				tmp[i+15:i] := ABS(a[k+7:k] - b[l+7:l]) + ABS(a[k+15:k+8] - b[l+15:l+8]) + ABS(a[k+23:k+16] - b[l+23:l+16]) + ABS(a[k+31:k+24] - b[l+31:l+24])
//			ENDFOR
//			RETURN tmp[127:0]
//		}
//		
//		dst[127:0] := MPSADBW(a[127:0], b[127:0], imm8[2:0])
//
// Instruction: 'MPSADBW'. Intrinsic: '_mm_mpsadbw_epu8'.
// Requires SSE4.1.
func MpsadbwEpu8(a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(mpsadbwEpu8([16]byte(a), [16]byte(b), imm8))
}

func mpsadbwEpu8(a [16]byte, b [16]byte, imm8 int) [16]byte


// MulEpi32: Multiply the low 32-bit integers from each packed 64-bit element
// in 'a' and 'b', and store the signed 64-bit results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//
// Instruction: 'PMULDQ'. Intrinsic: '_mm_mul_epi32'.
// Requires SSE4.1.
func MulEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(mulEpi32([16]byte(a), [16]byte(b)))
}

func mulEpi32(a [16]byte, b [16]byte) [16]byte


// MulloEpi32: Multiply the packed 32-bit integers in 'a' and 'b', producing
// intermediate 64-bit integers, and store the low 32 bits of the intermediate
// integers in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			tmp[63:0] := a[i+31:i] * b[i+31:i]
//			dst[i+31:i] := tmp[31:0]
//		ENDFOR
//
// Instruction: 'PMULLD'. Intrinsic: '_mm_mullo_epi32'.
// Requires SSE4.1.
func MulloEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(mulloEpi32([16]byte(a), [16]byte(b)))
}

func mulloEpi32(a [16]byte, b [16]byte) [16]byte


// PackusEpi32: Convert packed 32-bit integers from 'a' and 'b' to packed
// 16-bit integers using unsigned saturation, and store the results in 'dst'. 
//
//		dst[15:0] := Saturate_Int32_To_UnsignedInt16 (a[31:0])
//		dst[31:16] := Saturate_Int32_To_UnsignedInt16 (a[63:32])
//		dst[47:32] := Saturate_Int32_To_UnsignedInt16 (a[95:64])
//		dst[63:48] := Saturate_Int32_To_UnsignedInt16 (a[127:96])
//		dst[79:64] := Saturate_Int32_To_UnsignedInt16 (b[31:0])
//		dst[95:80] := Saturate_Int32_To_UnsignedInt16 (b[63:32])
//		dst[111:96] := Saturate_Int32_To_UnsignedInt16 (b[95:64])
//		dst[127:112] := Saturate_Int32_To_UnsignedInt16 (b[127:96])
//
// Instruction: 'PACKUSDW'. Intrinsic: '_mm_packus_epi32'.
// Requires SSE4.1.
func PackusEpi32(a x86.M128i, b x86.M128i) x86.M128i {
	return x86.M128i(packusEpi32([16]byte(a), [16]byte(b)))
}

func packusEpi32(a [16]byte, b [16]byte) [16]byte


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
//		FOR j := 0 to 1
//			i := j*64
//			dst[i+63:i] := ROUND(a[i+63:i])
//		ENDFOR
//
// Instruction: 'ROUNDPD'. Intrinsic: '_mm_round_pd'.
// Requires SSE4.1.
func RoundPd(a x86.M128d, rounding int) x86.M128d {
	return x86.M128d(roundPd([2]float64(a), rounding))
}

func roundPd(a [2]float64, rounding int) [2]float64


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
//		FOR j := 0 to 3
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//		ENDFOR
//
// Instruction: 'ROUNDPS'. Intrinsic: '_mm_round_ps'.
// Requires SSE4.1.
func RoundPs(a x86.M128, rounding int) x86.M128 {
	return x86.M128(roundPs([4]float32(a), rounding))
}

func roundPs(a [4]float32, rounding int) [4]float32


// RoundSd: Round the lower double-precision (64-bit) floating-point element in
// 'b' using the 'rounding' parameter, store the result as a double-precision
// floating-point element in the lower element of 'dst', and copy the upper
// element from 'a' to the upper element of 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		dst[63:0] := ROUND(b[63:0])
//		dst[127:64] := a[127:64]
//
// Instruction: 'ROUNDSD'. Intrinsic: '_mm_round_sd'.
// Requires SSE4.1.
func RoundSd(a x86.M128d, b x86.M128d, rounding int) x86.M128d {
	return x86.M128d(roundSd([2]float64(a), [2]float64(b), rounding))
}

func roundSd(a [2]float64, b [2]float64, rounding int) [2]float64


// RoundSs: Round the lower single-precision (32-bit) floating-point element in
// 'b' using the 'rounding' parameter, store the result as a single-precision
// floating-point element in the lower element of 'dst', and copy the upper 3
// packed elements from 'a' to the upper elements of 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		dst[31:0] := ROUND(b[31:0])
//		dst[127:32] := a[127:32]
//
// Instruction: 'ROUNDSS'. Intrinsic: '_mm_round_ss'.
// Requires SSE4.1.
func RoundSs(a x86.M128, b x86.M128, rounding int) x86.M128 {
	return x86.M128(roundSs([4]float32(a), [4]float32(b), rounding))
}

func roundSs(a [4]float32, b [4]float32, rounding int) [4]float32


// StreamLoadSi128: Load 128-bits of integer data from memory into 'dst' using
// a non-temporal memory hint.
// 	'mem_addr' must be aligned on a 16-byte boundary or a general-protection
// exception may be generated. 
//
//		dst[127:0] := MEM[mem_addr+127:mem_addr]
//
// Instruction: 'MOVNTDQA'. Intrinsic: '_mm_stream_load_si128'.
// Requires SSE4.1.
func StreamLoadSi128(mem_addr x86.M128i) x86.M128i {
	return x86.M128i(streamLoadSi128([16]byte(mem_addr)))
}

func streamLoadSi128(mem_addr [16]byte) [16]byte


// TestAllOnes: Compute the complement of 'a' and 0xFFFFFFFF, and return 1 if
// the result is zero, otherwise return 0. 
//
//		IF (a[127:0] AND NOT 0xFFFFFFFF == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: '...'. Intrinsic: '_mm_test_all_ones'.
// Requires SSE4.1.
func TestAllOnes(a x86.M128i) int {
	return int(testAllOnes([16]byte(a)))
}

func testAllOnes(a [16]byte) int


// TestAllZeros: Compute the bitwise AND of 128 bits (representing integer
// data) in 'a' and 'mask', and return 1 if the result is zero, otherwise
// return 0. 
//
//		IF (a[127:0] AND mask[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'PTEST'. Intrinsic: '_mm_test_all_zeros'.
// Requires SSE4.1.
func TestAllZeros(a x86.M128i, mask x86.M128i) int {
	return int(testAllZeros([16]byte(a), [16]byte(mask)))
}

func testAllZeros(a [16]byte, mask [16]byte) int


// TestMixOnesZeros: Compute the bitwise AND of 128 bits (representing integer
// data) in 'a' and 'mask', and set 'ZF' to 1 if the result is zero, otherwise
// set 'ZF' to 0. Compute the bitwise AND NOT of 'a' and 'mask', and set 'CF'
// to 1 if the result is zero, otherwise set 'CF' to 0. Return 1 if both the
// 'ZF' and 'CF' values are zero, otherwise return 0. 
//
//		IF (a[127:0] AND mask[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[127:0] AND NOT mask[127:0] == 0)
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
// Instruction: 'PTEST'. Intrinsic: '_mm_test_mix_ones_zeros'.
// Requires SSE4.1.
func TestMixOnesZeros(a x86.M128i, mask x86.M128i) int {
	return int(testMixOnesZeros([16]byte(a), [16]byte(mask)))
}

func testMixOnesZeros(a [16]byte, mask [16]byte) int


// TestcSi128: Compute the bitwise AND of 128 bits (representing integer data)
// in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set 'ZF'
// to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if the
// result is zero, otherwise set 'CF' to 0. Return the 'CF' value. 
//
//		IF (a[127:0] AND b[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[127:0] AND NOT b[127:0] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: 'PTEST'. Intrinsic: '_mm_testc_si128'.
// Requires SSE4.1.
func TestcSi128(a x86.M128i, b x86.M128i) int {
	return int(testcSi128([16]byte(a), [16]byte(b)))
}

func testcSi128(a [16]byte, b [16]byte) int


// TestnzcSi128: Compute the bitwise AND of 128 bits (representing integer
// data) in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set
// 'ZF' to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if
// the result is zero, otherwise set 'CF' to 0. Return 1 if both the 'ZF' and
// 'CF' values are zero, otherwise return 0. 
//
//		IF (a[127:0] AND b[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[127:0] AND NOT b[127:0] == 0)
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
// Instruction: 'PTEST'. Intrinsic: '_mm_testnzc_si128'.
// Requires SSE4.1.
func TestnzcSi128(a x86.M128i, b x86.M128i) int {
	return int(testnzcSi128([16]byte(a), [16]byte(b)))
}

func testnzcSi128(a [16]byte, b [16]byte) int


// TestzSi128: Compute the bitwise AND of 128 bits (representing integer data)
// in 'a' and 'b', and set 'ZF' to 1 if the result is zero, otherwise set 'ZF'
// to 0. Compute the bitwise AND NOT of 'a' and 'b', and set 'CF' to 1 if the
// result is zero, otherwise set 'CF' to 0. Return the 'ZF' value. 
//
//		IF (a[127:0] AND b[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[127:0] AND NOT b[127:0] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'PTEST'. Intrinsic: '_mm_testz_si128'.
// Requires SSE4.1.
func TestzSi128(a x86.M128i, b x86.M128i) int {
	return int(testzSi128([16]byte(a), [16]byte(b)))
}

func testzSi128(a [16]byte, b [16]byte) int

