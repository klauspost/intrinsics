// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package bmi1

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// BextrU32: Extract contiguous bits from unsigned 32-bit integer 'a', and
// store the result in 'dst'. Extract the number of bits specified by 'len',
// starting at the bit specified by 'start'. 
//
//		tmp := ZERO_EXTEND_TO_512(a)
//		dst := ZERO_EXTEND(tmp[start+len-1:start])
//
// Instruction: 'BEXTR'. Intrinsic: '_bextr_u32'.
// Requires BMI1.
func BextrU32(a uint32, start uint32, len uint32) uint32 {
	panic("not implemented")
}


// BextrU64: Extract contiguous bits from unsigned 64-bit integer 'a', and
// store the result in 'dst'. Extract the number of bits specified by 'len',
// starting at the bit specified by 'start'. 
//
//		tmp := ZERO_EXTEND_TO_512(a)
//		dst := ZERO_EXTEND(tmp[start+len-1:start])
//
// Instruction: 'BEXTR'. Intrinsic: '_bextr_u64'.
// Requires BMI1.
func BextrU64(a uint64, start uint32, len uint32) uint64 {
	panic("not implemented")
}


// BlsiU32: Extract the lowest set bit from unsigned 32-bit integer 'a' and set
// the corresponding bit in 'dst'. All other bits in 'dst' are zeroed, and all
// bits are zeroed if no bits are set in 'a'. 
//
//		dst := (-a) BITWISE AND a
//
// Instruction: 'BLSI'. Intrinsic: '_blsi_u32'.
// Requires BMI1.
func BlsiU32(a uint32) uint32 {
	panic("not implemented")
}


// BlsiU64: Extract the lowest set bit from unsigned 64-bit integer 'a' and set
// the corresponding bit in 'dst'. All other bits in 'dst' are zeroed, and all
// bits are zeroed if no bits are set in 'a'. 
//
//		dst := (-a) BITWISE AND a
//
// Instruction: 'BLSI'. Intrinsic: '_blsi_u64'.
// Requires BMI1.
func BlsiU64(a uint64) uint64 {
	panic("not implemented")
}


// BlsmskU32: Set all the lower bits of 'dst' up to and including the lowest
// set bit in unsigned 32-bit integer 'a'. 
//
//		dst := (a - 1) XOR a
//
// Instruction: 'BLSMSK'. Intrinsic: '_blsmsk_u32'.
// Requires BMI1.
func BlsmskU32(a uint32) uint32 {
	panic("not implemented")
}


// BlsmskU64: Set all the lower bits of 'dst' up to and including the lowest
// set bit in unsigned 64-bit integer 'a'. 
//
//		dst := (a - 1) XOR a
//
// Instruction: 'BLSMSK'. Intrinsic: '_blsmsk_u64'.
// Requires BMI1.
func BlsmskU64(a uint64) uint64 {
	panic("not implemented")
}


// BlsrU32: Copy all bits from unsigned 32-bit integer 'a' to 'dst', and reset
// (set to 0) the bit in 'dst' that corresponds to the lowest set bit in 'a'. 
//
//		dst := (a - 1) BITWISE AND a
//
// Instruction: 'BLSR'. Intrinsic: '_blsr_u32'.
// Requires BMI1.
func BlsrU32(a uint32) uint32 {
	panic("not implemented")
}


// BlsrU64: Copy all bits from unsigned 64-bit integer 'a' to 'dst', and reset
// (set to 0) the bit in 'dst' that corresponds to the lowest set bit in 'a'. 
//
//		dst := (a - 1) BITWISE AND a
//
// Instruction: 'BLSR'. Intrinsic: '_blsr_u64'.
// Requires BMI1.
func BlsrU64(a uint64) uint64 {
	panic("not implemented")
}


// Tzcnt32: Count the number of trailing zero bits in unsigned 32-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 32) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_mm_tzcnt_32'.
// Requires BMI1.
func Tzcnt32(a uint32) int {
	panic("not implemented")
}


// Tzcnt64: Count the number of trailing zero bits in unsigned 64-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 64) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_mm_tzcnt_64'.
// Requires BMI1.
func Tzcnt64(a uint64) int64 {
	panic("not implemented")
}


// TzcntU32: Count the number of trailing zero bits in unsigned 32-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 32) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_tzcnt_u32'.
// Requires BMI1.
func TzcntU32(a uint32) uint32 {
	panic("not implemented")
}


// TzcntU64: Count the number of trailing zero bits in unsigned 64-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 64) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_tzcnt_u64'.
// Requires BMI1.
func TzcntU64(a uint64) uint64 {
	panic("not implemented")
}

