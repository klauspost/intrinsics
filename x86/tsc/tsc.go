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
func Rdtsc() (dst int64) {
	return int64(rdtsc())
}

func rdtsc() int64

