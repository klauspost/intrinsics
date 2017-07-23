// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package tsc

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Rdtsc: Copy the current 64-bit value of the processor's time-stamp counter
// into 'dst'. 
//
//		dst[63:0] := TimeStampCounter
//
// Instruction: 'RDTSC'. Intrinsic: '_rdtsc'.
// Requires TSC.
func Rdtsc() int64 {
	panic("not implemented")
}

