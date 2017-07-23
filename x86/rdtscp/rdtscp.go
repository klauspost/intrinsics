// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package rdtscp

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Rdtscp: Copy the current 64-bit value of the processor's time-stamp counter
// into 'dst', and store the IA32_TSC_AUX MSR (signature value) into memory at
// 'mem_addr'. 
//
//		dst[63:0] := TimeStampCounter
//		MEM[mem_addr+31:mem_addr] := IA32_TSC_AUX[31:0]
//
// Instruction: 'RDTSCP'. Intrinsic: '__rdtscp'.
// Requires RDTSCP.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Rdtscp(mem_addr *uint32) uint64 {
	panic("not implemented")
}

