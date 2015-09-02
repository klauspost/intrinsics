package rdrand

import "github.com/klauspost/intrinsics/x86"


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
func Rdrand16Step(val uint16) int {
	return int(rdrand16Step(val))
}

func rdrand16Step(val uint16) int


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
func Rdrand32Step(val uint32) int {
	return int(rdrand32Step(val))
}

func rdrand32Step(val uint32) int


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
func Rdrand64Step(val uint64) int {
	return int(rdrand64Step(val))
}

func rdrand64Step(val uint64) int

