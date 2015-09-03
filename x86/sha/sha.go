package sha

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Sha1msg1Epu32: Perform an intermediate calculation for the next four SHA1
// message values (unsigned 32-bit integers) using previous message values from
// 'a' and 'b', and store the result in 'dst'. 
//
//		W0 := a[127:96];
//		W1 := a[95:64];
//		W2 := a[63:32];
//		W3 := a[31:0];
//		W4 := b[127:96];
//		W5 := b[95:64];
//		
//		dst[127:96] := W2 XOR W0;
//		dst[95:64] := W3 XOR W1;
//		dst[63:32] := W4 XOR W2;
//		dst[31:0] := W5 XOR W3;
//
// Instruction: 'SHA1MSG1'. Intrinsic: '_mm_sha1msg1_epu32'.
// Requires SHA.
func Sha1msg1Epu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(sha1msg1Epu32([16]byte(a), [16]byte(b)))
}

func sha1msg1Epu32(a [16]byte, b [16]byte) [16]byte


// Sha1msg2Epu32: Perform the final calculation for the next four SHA1 message
// values (unsigned 32-bit integers) using the intermediate result in 'a' and
// the previous message values in 'b', and store the result in 'dst'. 
//
//		W13 := b[95:64];
//		W14 := b[63:32];
//		W15 := b[31:0];
//		W16 := (a[127:96] XOR W13) <<< 1;
//		W17 := (a[95:64] XOR W14) <<< 1;
//		W18 := (a[63:32] XOR W15) <<< 1;
//		W19 := (a[31:0] XOR W16) <<< 1;
//		
//		dst[127:96] := W16;
//		dst[95:64] := W17;
//		dst[63:32] := W18;
//		dst[31:0] := W19;
//
// Instruction: 'SHA1MSG2'. Intrinsic: '_mm_sha1msg2_epu32'.
// Requires SHA.
func Sha1msg2Epu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(sha1msg2Epu32([16]byte(a), [16]byte(b)))
}

func sha1msg2Epu32(a [16]byte, b [16]byte) [16]byte


// Sha1nexteEpu32: Calculate SHA1 state variable E after four rounds of
// operation from the current SHA1 state variable 'a', add that value to the
// scheduled values (unsigned 32-bit integers) in 'b', and store the result in
// 'dst'. 
//
//		tmp := (a[127:96] <<< 30);
//		dst[127:96] := b[127:96] + tmp;
//		dst[95:64] := b[95:64];
//		dst[63:32] := b[63:32];
//		dst[31:0] := b[31:0];
//
// Instruction: 'SHA1NEXTE'. Intrinsic: '_mm_sha1nexte_epu32'.
// Requires SHA.
func Sha1nexteEpu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(sha1nexteEpu32([16]byte(a), [16]byte(b)))
}

func sha1nexteEpu32(a [16]byte, b [16]byte) [16]byte


// Sha1rnds4Epu32: Perform four rounds of SHA1 operation using an initial SHA1
// state (A,B,C,D) from 'a' and some pre-computed sum of the next 4 round
// message values (unsigned 32-bit integers), and state variable E from 'b',
// and store the updated SHA1 state (A,B,C,D) in 'dst'. 'func' contains the
// logic functions and round constants. 
//
//		IF (func[1:0] = 0) THEN
//			f() := f0(), K := K0;
//		ELSE IF (func[1:0] = 1) THEN
//			f() := f1(), K := K1;
//		ELSE IF (func[1:0] = 2) THEN
//			f() := f2(), K := K2;
//		ELSE IF (func[1:0] = 3) THEN
//			f() := f3(), K := K3;
//		FI;
//		
//		A := a[127:96];
//		B := a[95:64];
//		C := a[63:32];
//		D := a[31:0];
//		
//		W[0] := b[127:96];
//		W[1] := b[95:64];
//		W[2] := b[63:32];
//		W[3] := b[31:0];
//		
//		A[1] := f(B, C, D) + (A <<< 5) + W[0] + K;
//		B[1] := A;
//		C[1] := B <<< 30;
//		D[1] := C;
//		E[1] := D;
//		
//		FOR i = 1 to 3
//				A[i+1] := f(B[i], C[i], D[i]) + (A[i] <<< 5) + W[i] + E[i] + K;
//				B[i+1] := A[i];
//				C[i+1] := B[i] <<< 30;
//				D[i+1] := C[i];
//				E[i+1] := D[i];
//		ENDFOR;
//		
//		dst[127:96] := A[4];
//		dst[95:64] := B[4];
//		dst[63:32] := C[4];
//		dst[31:0] := D[4];
//
// Instruction: 'SHA1RNDS4'. Intrinsic: '_mm_sha1rnds4_epu32'.
// Requires SHA.
func Sha1rnds4Epu32(a x86.M128i, b x86.M128i, fnc int) (dst x86.M128i) {
	return x86.M128i(sha1rnds4Epu32([16]byte(a), [16]byte(b), fnc))
}

func sha1rnds4Epu32(a [16]byte, b [16]byte, fnc int) [16]byte


// Sha256msg1Epu32: Perform an intermediate calculation for the next four
// SHA256 message values (unsigned 32-bit integers) using previous message
// values from 'a' and 'b', and store the result in 'dst'. 
//
//		W4 := b[31:0];
//		W3 := a[127:96];
//		W2 := a[95:64];
//		W1 := a[63:32];
//		W0 := a[31:0];
//		
//		dst[127:96] := W3 + sigma0(W4);
//		dst[95:64] := W2 + sigma0(W3);
//		dst[63:32] := W1 + sigma0(W2);
//		dst[31:0] := W0 + sigma0(W1);
//
// Instruction: 'SHA256MSG1'. Intrinsic: '_mm_sha256msg1_epu32'.
// Requires SHA.
func Sha256msg1Epu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(sha256msg1Epu32([16]byte(a), [16]byte(b)))
}

func sha256msg1Epu32(a [16]byte, b [16]byte) [16]byte


// Sha256msg2Epu32: Perform the final calculation for the next four SHA256
// message values (unsigned 32-bit integers) using previous message values from
// 'a' and 'b', and store the result in 'dst'." 
//
//		W14 := b[95:64];
//		W15 := b[127:96];
//		W16 := a[31:0] + sigma1(W14);
//		W17 := a[63:32] + sigma1(W15);
//		W18 := a[95:64] + sigma1(W16);
//		W19 := a[127:96] + sigma1(W17);
//		
//		dst[127:96] := W19;
//		dst[95:64] := W18;
//		dst[63:32] := W17;
//		dst[31:0] := W16;
//
// Instruction: 'SHA256MSG2'. Intrinsic: '_mm_sha256msg2_epu32'.
// Requires SHA.
func Sha256msg2Epu32(a x86.M128i, b x86.M128i) (dst x86.M128i) {
	return x86.M128i(sha256msg2Epu32([16]byte(a), [16]byte(b)))
}

func sha256msg2Epu32(a [16]byte, b [16]byte) [16]byte


// Sha256rnds2Epu32: Perform 2 rounds of SHA256 operation using an initial
// SHA256 state (C,D,G,H) from 'a', an initial SHA256 state (A,B,E,F) from 'b',
// and a pre-computed sum of the next 2 round message values (unsigned 32-bit
// integers) and the corresponding round constants from 'k', and store the
// updated SHA256 state (A,B,E,F) in 'dst'. 
//
//		A[0] := b[127:96];
//		B[0] := b[95:64];
//		C[0] := a[127:96];
//		D[0] := a[95:64];
//		E[0] := b[63:32];
//		F[0] := b[31:0];
//		G[0] := a[63:32];
//		H[0] := a[31:0];
//		
//		W_K0 := k[31:0];
//		W_K1 := k[63:32];
//		
//		FOR i = 0 to 1
//				A_(i+1) := Ch(E[i], F[i], G[i]) + sum1(E[i]) + WKi + H[i] + Maj(A[i], B[i], C[i]) + sum0(A[i]);
//				B_(i+1) := A[i];
//				C_(i+1) := B[i];
//				D_(i+1) := C[i];
//				E_(i+1) := Ch(E[i], F[i], G[i]) + sum1(E[i]) + WKi + H[i] + D[i];
//				F_(i+1) := E[i];
//				G_(i+1) := F[i];
//				H_(i+1) := G[i];
//		ENDFOR;
//		
//		dst[127:96] := A[2];
//		dst[95:64] := B[2];
//		dst[63:32] := E[2];
//		dst[31:0] := F[2];
//
// Instruction: 'SHA256RNDS2'. Intrinsic: '_mm_sha256rnds2_epu32'.
// Requires SHA.
func Sha256rnds2Epu32(a x86.M128i, b x86.M128i, k x86.M128i) (dst x86.M128i) {
	return x86.M128i(sha256rnds2Epu32([16]byte(a), [16]byte(b), [16]byte(k)))
}

func sha256rnds2Epu32(a [16]byte, b [16]byte, k [16]byte) [16]byte

