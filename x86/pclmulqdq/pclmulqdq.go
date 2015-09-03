package pclmulqdq

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Clmulepi64Si128: Perform a carry-less multiplication of two 64-bit integers,
// selected from 'a' and 'b' according to 'imm8', and store the results in
// 'dst'. 
//
//		IF (imm8[0] = 0)
//			TEMP1 := a[63:0];
//		ELSE
//			TEMP1 := a[127:64];
//		FI 
//		IF (imm8[4] = 0)
//			TEMP2 := b[63:0];
//		ELSE 
//			TEMP2 := b[127:64];
//		FI
//		
//		FOR i := 0 to 63
//			TEMP[i] := (TEMP1[0] and TEMP2[i]);
//			FOR j := 1 to i
//				TEMP [i] := TEMP [i] XOR (TEMP1[j] AND TEMP2[i-j])
//			ENDFOR 
//			dst[i] := TEMP[i];
//		ENDFOR
//		FOR i := 64 to 127
//			TEMP [i] := 0;
//			FOR j := (i - 63) to 63
//				TEMP [i] := TEMP [i] XOR (TEMP1[j] AND TEMP2[i-j])
//			ENDFOR
//			dst[i] := TEMP[i];
//		ENDFOR
//		dst[127] := 0
//
// Instruction: 'PCLMULQDQ'. Intrinsic: '_mm_clmulepi64_si128'.
// Requires PCLMULQDQ.
//
// FIXME: Requires compiler support (has immediate)
func Clmulepi64Si128(a x86.M128i, b x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(clmulepi64Si128([16]byte(a), [16]byte(b), imm8))
}

func clmulepi64Si128(a [16]byte, b [16]byte, imm8 int) [16]byte

