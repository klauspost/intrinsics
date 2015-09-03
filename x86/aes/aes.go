package aes

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AesdecSi128: Perform one round of an AES decryption flow on data (state) in
// 'a' using the round key in 'RoundKey', and store the result in 'dst'." 
//
//		state := a
//		a[127:0] := InvShiftRows(a[127:0])
//		a[127:0] := InvSubBytes(a[127:0])
//		a[127:0] := InvMixColumns(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESDEC'. Intrinsic: '_mm_aesdec_si128'.
// Requires AES.
func AesdecSi128(a x86.M128i, RoundKey x86.M128i) x86.M128i {
	return x86.M128i(aesdecSi128([16]byte(a), [16]byte(RoundKey)))
}

func aesdecSi128(a [16]byte, RoundKey [16]byte) [16]byte


// AesdeclastSi128: Perform the last round of an AES decryption flow on data
// (state) in 'a' using the round key in 'RoundKey', and store the result in
// 'dst'." 
//
//		state := a
//		a[127:0] := InvShiftRows(a[127:0])
//		a[127:0] := InvSubBytes(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESDECLAST'. Intrinsic: '_mm_aesdeclast_si128'.
// Requires AES.
func AesdeclastSi128(a x86.M128i, RoundKey x86.M128i) x86.M128i {
	return x86.M128i(aesdeclastSi128([16]byte(a), [16]byte(RoundKey)))
}

func aesdeclastSi128(a [16]byte, RoundKey [16]byte) [16]byte


// AesencSi128: Perform one round of an AES encryption flow on data (state) in
// 'a' using the round key in 'RoundKey', and store the result in 'dst'." 
//
//		state := a
//		a[127:0] := ShiftRows(a[127:0])
//		a[127:0] := SubBytes(a[127:0])
//		a[127:0] := MixColumns(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESENC'. Intrinsic: '_mm_aesenc_si128'.
// Requires AES.
func AesencSi128(a x86.M128i, RoundKey x86.M128i) x86.M128i {
	return x86.M128i(aesencSi128([16]byte(a), [16]byte(RoundKey)))
}

func aesencSi128(a [16]byte, RoundKey [16]byte) [16]byte


// AesenclastSi128: Perform the last round of an AES encryption flow on data
// (state) in 'a' using the round key in 'RoundKey', and store the result in
// 'dst'." 
//
//		state := a
//		a[127:0] := ShiftRows(a[127:0])
//		a[127:0] := SubBytes(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESENCLAST'. Intrinsic: '_mm_aesenclast_si128'.
// Requires AES.
func AesenclastSi128(a x86.M128i, RoundKey x86.M128i) x86.M128i {
	return x86.M128i(aesenclastSi128([16]byte(a), [16]byte(RoundKey)))
}

func aesenclastSi128(a [16]byte, RoundKey [16]byte) [16]byte


// AesimcSi128: Perform the InvMixColumns transformation on 'a' and store the
// result in 'dst'. 
//
//		dst[127:0] := InvMixColumns(a[127:0])
//
// Instruction: 'AESIMC'. Intrinsic: '_mm_aesimc_si128'.
// Requires AES.
func AesimcSi128(a x86.M128i) x86.M128i {
	return x86.M128i(aesimcSi128([16]byte(a)))
}

func aesimcSi128(a [16]byte) [16]byte


// AeskeygenassistSi128: Assist in expanding the AES cipher key by computing
// steps towards generating a round key for encryption cipher using data from
// 'a' and an 8-bit round constant specified in 'imm8', and store the result in
// 'dst'." 
//
//		X3[31:0] := a[127:96]
//		X2[31:0] := a[95:64]
//		X1[31:0] := a[63:32]
//		X0[31:0] := a[31:0]
//		RCON[31:0] := ZeroExtend(imm8[7:0]);
//		dst[31:0] := SubWord(X1)
//		dst[63:32] := (RotWord(SubWord(X1)) XOR RCON;
//		dst[95:64] := SubWord(X3)
//		dst[127:96] := RotWord(SubWord(X3)) XOR RCON;
//
// Instruction: 'AESKEYGENASSIST'. Intrinsic: '_mm_aeskeygenassist_si128'.
// Requires AES.
//
// FIXME: Requires compiler support (has immediate)
func AeskeygenassistSi128(a x86.M128i, imm8 int) x86.M128i {
	return x86.M128i(aeskeygenassistSi128([16]byte(a), imm8))
}

func aeskeygenassistSi128(a [16]byte, imm8 int) [16]byte

