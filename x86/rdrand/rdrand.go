package rdrand

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Rdrand16Step: Read a hardware generated 16-bit random value and store the
// result in 'val'. Return 1 if a random value was generated, and 0 otherwise. 
//
//		IF HW_RND_GEN.ready = 1
//			val[15:0] := HW_RND_GEN.data;
//			RETURN 1;
//		ELSE
//			val[15:0] := 0;
//			RETURN 0;
//		FI
//
// Instruction: 'RDRAND'. Intrinsic: '_rdrand16_step'.
// Requires RDRAND.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Rdrand16Step(val *uint16) (dst int) {
	// FIXME: Rework to avoid possible return value as parameter.
	return 0
}

// Rdrand32Step: Read a hardware generated 32-bit random value and store the
// result in 'val'. Return 1 if a random value was generated, and 0 otherwise. 
//
//		IF HW_RND_GEN.ready = 1
//			val[31:0] := HW_RND_GEN.data;
//			RETURN 1;
//		ELSE
//			val[31:0] := 0;
//			RETURN 0;
//		FI
//
// Instruction: 'RDRAND'. Intrinsic: '_rdrand32_step'.
// Requires RDRAND.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Rdrand32Step(val *uint32) (dst int) {
	// FIXME: Rework to avoid possible return value as parameter.
	return 0
}

// Rdrand64Step: Read a hardware generated 64-bit random value and store the
// result in 'val'. Return 1 if a random value was generated, and 0 otherwise. 
//
//		IF HW_RND_GEN.ready = 1
//			val[63:0] := HW_RND_GEN.data;
//			RETURN 1;
//		ELSE
//			val[63:0] := 0;
//			RETURN 0;
//		FI
//
// Instruction: 'RDRAND'. Intrinsic: '_rdrand64_step'.
// Requires RDRAND.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Rdrand64Step(val *uint64) (dst int) {
	// FIXME: Rework to avoid possible return value as parameter.
	return 0
}
