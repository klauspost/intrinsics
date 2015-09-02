package sse3

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AddsubPd: Alternatively add and subtract packed double-precision (64-bit)
// floating-point elements in 'a' to/from packed elements in 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 1
//			i := j*64
//			IF (j is even) 
//				dst[i+63:i] := a[i+63:i] - b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			FI
//		ENDFOR
//
// Instruction: 'ADDSUBPD'. Intrinsic: '_mm_addsub_pd'.
// Requires SSE3.
func AddsubPd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(addsubPd([2]float64(a), [2]float64(b)))
}

func addsubPd(a [2]float64, b [2]float64) [2]float64


// AddsubPs: Alternatively add and subtract packed single-precision (32-bit)
// floating-point elements in 'a' to/from packed elements in 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 3
//			i := j*32
//			IF (j is even) 
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			FI
//		ENDFOR
//
// Instruction: 'ADDSUBPS'. Intrinsic: '_mm_addsub_ps'.
// Requires SSE3.
func AddsubPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(addsubPs([4]float32(a), [4]float32(b)))
}

func addsubPs(a [4]float32, b [4]float32) [4]float32


// HaddPd: Horizontally add adjacent pairs of double-precision (64-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[63:0] := a[127:64] + a[63:0]
//		dst[127:64] := b[127:64] + b[63:0]
//
// Instruction: 'HADDPD'. Intrinsic: '_mm_hadd_pd'.
// Requires SSE3.
func HaddPd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(haddPd([2]float64(a), [2]float64(b)))
}

func haddPd(a [2]float64, b [2]float64) [2]float64


// HaddPs: Horizontally add adjacent pairs of single-precision (32-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[31:0] := a[63:32] + a[31:0]
//		dst[63:32] := a[127:96] + a[95:64]
//		dst[95:64] := b[63:32] + b[31:0]
//		dst[127:96] := b[127:96] + b[95:64]
//
// Instruction: 'HADDPS'. Intrinsic: '_mm_hadd_ps'.
// Requires SSE3.
func HaddPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(haddPs([4]float32(a), [4]float32(b)))
}

func haddPs(a [4]float32, b [4]float32) [4]float32


// HsubPd: Horizontally subtract adjacent pairs of double-precision (64-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[63:0] := a[63:0] - a[127:64]
//		dst[127:64] := b[63:0] - b[127:64]
//
// Instruction: 'HSUBPD'. Intrinsic: '_mm_hsub_pd'.
// Requires SSE3.
func HsubPd(a x86.M128d, b x86.M128d) x86.M128d {
	return x86.M128d(hsubPd([2]float64(a), [2]float64(b)))
}

func hsubPd(a [2]float64, b [2]float64) [2]float64


// HsubPs: Horizontally add adjacent pairs of single-precision (32-bit)
// floating-point elements in 'a' and 'b', and pack the results in 'dst'. 
//
//		dst[31:0] := a[31:0] - a[63:32]
//		dst[63:32] := a[95:64] - a[127:96]
//		dst[95:64] := b[31:0] - b[63:32]
//		dst[127:96] := b[95:64] - b[127:96]
//
// Instruction: 'HSUBPS'. Intrinsic: '_mm_hsub_ps'.
// Requires SSE3.
func HsubPs(a x86.M128, b x86.M128) x86.M128 {
	return x86.M128(hsubPs([4]float32(a), [4]float32(b)))
}

func hsubPs(a [4]float32, b [4]float32) [4]float32


// LddquSi128: Load 128-bits of integer data from unaligned memory into 'dst'.
// This intrinsic may perform better than '_mm_loadu_si128' when the data
// crosses a cache line boundary. 
//
//		dst[127:0] := MEM[mem_addr+127:mem_addr]
//
// Instruction: 'LDDQU'. Intrinsic: '_mm_lddqu_si128'.
// Requires SSE3.
func LddquSi128(mem_addr x86.M128iConst) x86.M128i {
	return x86.M128i(lddquSi128(mem_addr))
}

func lddquSi128(mem_addr x86.M128iConst) [16]byte


// LoaddupPd: Load a double-precision (64-bit) floating-point element from
// memory into both elements of 'dst'. 
//
//		tmp[63:0] := MEM[mem_addr+63:mem_addr]
//		tmp[127:64] := MEM[mem_addr+63:mem_addr]
//
// Instruction: 'MOVDDUP'. Intrinsic: '_mm_loaddup_pd'.
// Requires SSE3.
func LoaddupPd(mem_addr float64) x86.M128d {
	return x86.M128d(loaddupPd(mem_addr))
}

func loaddupPd(mem_addr float64) [2]float64


// MovedupPd: Duplicate the low double-precision (64-bit) floating-point
// element from 'a', and store the results in 'dst'. 
//
//		tmp[63:0] := a[63:0]
//		tmp[127:64] := a[63:0]
//
// Instruction: 'MOVDDUP'. Intrinsic: '_mm_movedup_pd'.
// Requires SSE3.
func MovedupPd(a x86.M128d) x86.M128d {
	return x86.M128d(movedupPd([2]float64(a)))
}

func movedupPd(a [2]float64) [2]float64


// MovehdupPs: Duplicate odd-indexed single-precision (32-bit) floating-point
// elements from 'a', and store the results in 'dst'. 
//
//		dst[31:0] := a[63:32] 
//		dst[63:32] := a[63:32]
//		dst[95:64] := a[127:96] 
//		dst[127:96] := a[127:96]
//
// Instruction: 'MOVSHDUP'. Intrinsic: '_mm_movehdup_ps'.
// Requires SSE3.
func MovehdupPs(a x86.M128) x86.M128 {
	return x86.M128(movehdupPs([4]float32(a)))
}

func movehdupPs(a [4]float32) [4]float32


// MoveldupPs: Duplicate even-indexed single-precision (32-bit) floating-point
// elements from 'a', and store the results in 'dst'. 
//
//		dst[31:0] := a[31:0] 
//		dst[63:32] := a[31:0]
//		dst[95:64] := a[95:64] 
//		dst[127:96] := a[95:64]
//
// Instruction: 'MOVSLDUP'. Intrinsic: '_mm_moveldup_ps'.
// Requires SSE3.
func MoveldupPs(a x86.M128) x86.M128 {
	return x86.M128(moveldupPs([4]float32(a)))
}

func moveldupPs(a [4]float32) [4]float32

