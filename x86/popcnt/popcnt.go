// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package popcnt

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// PopcntU32: Count the number of bits set to 1 in unsigned 32-bit integer 'a',
// and return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 31
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_popcnt_u32'.
// Requires POPCNT.
func PopcntU32(a uint32) int {
	return int(popcntU32(a))
}

func popcntU32(a uint32) int


// PopcntU64: Count the number of bits set to 1 in unsigned 64-bit integer 'a',
// and return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 63
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_popcnt_u64'.
// Requires POPCNT.
func PopcntU64(a uint64) int64 {
	return int64(popcntU64(a))
}

func popcntU64(a uint64) int64


// Popcnt32: Count the number of bits set to 1 in 32-bit integer 'a', and
// return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 31
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_popcnt32'.
// Requires POPCNT.
func Popcnt32(a int) int {
	return int(popcnt32(a))
}

func popcnt32(a int) int


// Popcnt64: Count the number of bits set to 1 in 64-bit integer 'a', and
// return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 63
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_popcnt64'.
// Requires POPCNT.
func Popcnt64(a int64) int {
	return int(popcnt64(a))
}

func popcnt64(a int64) int

