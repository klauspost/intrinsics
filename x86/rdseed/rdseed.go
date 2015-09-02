package rdseed

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Rdseed16Step: Read a 16-bit NIST SP800-90B and SP800-90C compliant random
// value and store in 'val'. Return 1 if a random value was generated, and 0
// otherwise. 
//
//		IF HW_NRND_GEN.ready = 1 THEN
//			val[15:0] := HW_NRND_GEN.data
//			RETURN 1
//		ELSE
//			val[15:0] := 0
//			RETURN 0
//		FI
//
// Instruction: 'RDSEED'. Intrinsic: '_rdseed16_step'.
// Requires RDSEED.
func Rdseed16Step(val uint16) int {
	return int(rdseed16Step(val))
}

func rdseed16Step(val uint16) int


// Rdseed32Step: Read a 32-bit NIST SP800-90B and SP800-90C compliant random
// value and store in 'val'. Return 1 if a random value was generated, and 0
// otherwise. 
//
//		IF HW_NRND_GEN.ready = 1 THEN
//			val[31:0] := HW_NRND_GEN.data
//			RETURN 1
//		ELSE
//			val[31:0] := 0
//			RETURN 0
//		FI
//
// Instruction: 'RDSEED'. Intrinsic: '_rdseed32_step'.
// Requires RDSEED.
func Rdseed32Step(val uint32) int {
	return int(rdseed32Step(val))
}

func rdseed32Step(val uint32) int


// Rdseed64Step: Read a 64-bit NIST SP800-90B and SP800-90C compliant random
// value and store in 'val'. Return 1 if a random value was generated, and 0
// otherwise. 
//
//		IF HW_NRND_GEN.ready = 1 THEN
//			val[63:0] := HW_NRND_GEN.data
//			RETURN 1
//		ELSE
//			val[63:0] := 0
//			RETURN 0
//		FI
//
// Instruction: 'RDSEED'. Intrinsic: '_rdseed64_step'.
// Requires RDSEED.
func Rdseed64Step(val uint64) int {
	return int(rdseed64Step(val))
}

func rdseed64Step(val uint64) int

