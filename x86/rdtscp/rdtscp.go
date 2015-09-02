package rdtscp

import . "github.com/klauspost/intrinsics/x86"


// Rdtscp: Copy the current 64-bit value of the processor's time-stamp counter
// into 'dst', and store the IA32_TSC_AUX MSR (signature value) into memory at
// 'mem_addr'. 
//
//		dst[63:0] := TimeStampCounter
//		MEM[mem_addr+31:mem_addr] := IA32_TSC_AUX[31:0]
//
// Instruction: 'RDTSCP'. Intrinsic: '__rdtscp'.
// Requires RDTSCP.
func Rdtscp(mem_addr uint32) uint64 {
	return uint64(rdtscp(mem_addr))
}

func rdtscp(mem_addr uint32) uint64

