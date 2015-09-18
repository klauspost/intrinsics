// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package bmi2

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// BzhiU32: Copy all bits from unsigned 32-bit integer 'a' to 'dst', and reset
// (set to 0) the high bits in 'dst' starting at 'index'. 
//
//		n := index[7:0]
//		dst := a
//		IF (n < 32)
//			dst[31:n] := 0
//		FI
//
// Instruction: 'BZHI'. Intrinsic: '_bzhi_u32'.
// Requires BMI2.
func BzhiU32(a uint32, index uint32) uint32 {
	return uint32(bzhiU32(a, index))
}

func bzhiU32(a uint32, index uint32) uint32


// BzhiU64: Copy all bits from unsigned 64-bit integer 'a' to 'dst', and reset
// (set to 0) the high bits in 'dst' starting at 'index'. 
//
//		n := index[7:0]
//		dst := a
//		IF (n < 64)
//			dst[63:n] := 0
//		FI
//
// Instruction: 'BZHI'. Intrinsic: '_bzhi_u64'.
// Requires BMI2.
func BzhiU64(a uint64, index uint32) uint64 {
	return uint64(bzhiU64(a, index))
}

func bzhiU64(a uint64, index uint32) uint64


// PdepU32: Deposit contiguous low bits from unsigned 32-bit integer 'a' to
// 'dst' at the corresponding bit locations specified by 'mask'; all other bits
// in 'dst' are set to zero. 
//
//		tmp := a
//		dst := 0
//		m := 0
//		k := 0
//		DO WHILE m < 32
//			IF mask[m] = 1
//				dst[m] := tmp[k]
//				k := k + 1
//			FI
//			m := m + 1
//		OD
//
// Instruction: 'PDEP'. Intrinsic: '_pdep_u32'.
// Requires BMI2.
func PdepU32(a uint32, mask uint32) uint32 {
	return uint32(pdepU32(a, mask))
}

func pdepU32(a uint32, mask uint32) uint32


// PdepU64: Deposit contiguous low bits from unsigned 64-bit integer 'a' to
// 'dst' at the corresponding bit locations specified by 'mask'; all other bits
// in 'dst' are set to zero. 
//
//		tmp := a
//		dst := 0
//		m := 0
//		k := 0
//		DO WHILE m < 64
//			IF mask[m] = 1
//				dst[m] := tmp[k]
//				k := k + 1
//			FI
//			m := m + 1
//		OD
//
// Instruction: 'PDEP'. Intrinsic: '_pdep_u64'.
// Requires BMI2.
func PdepU64(a uint64, mask uint64) uint64 {
	return uint64(pdepU64(a, mask))
}

func pdepU64(a uint64, mask uint64) uint64


// PextU32: Extract bits from unsigned 32-bit integer 'a' at the corresponding
// bit locations specified by 'mask' to contiguous low bits in 'dst'; the
// remaining upper bits in 'dst' are set to zero. 
//
//		tmp := a
//		dst := 0
//		m := 0
//		k := 0
//		DO WHILE m < 32
//			IF mask[m] = 1
//				dst[k] := tmp[m]
//				k := k + 1
//			FI
//			m := m + 1
//		OD
//
// Instruction: 'PEXT'. Intrinsic: '_pext_u32'.
// Requires BMI2.
func PextU32(a uint32, mask uint32) uint32 {
	return uint32(pextU32(a, mask))
}

func pextU32(a uint32, mask uint32) uint32


// PextU64: Extract bits from unsigned 64-bit integer 'a' at the corresponding
// bit locations specified by 'mask' to contiguous low bits in 'dst'; the
// remaining upper bits in 'dst' are set to zero. 
//
//		tmp := a
//		dst := 0
//		m := 0
//		k := 0
//		DO WHILE m < 64
//			IF mask[m] = 1
//				dst[k] := tmp[m]
//				k := k + 1
//			FI
//			m := m + 1
//		OD
//
// Instruction: 'PEXT'. Intrinsic: '_pext_u64'.
// Requires BMI2.
func PextU64(a uint64, mask uint64) uint64 {
	return uint64(pextU64(a, mask))
}

func pextU64(a uint64, mask uint64) uint64

