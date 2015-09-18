// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package monitor

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Skipped: _mm_monitor. Contains pointer parameter.


// Mwait: Hint to the processor that it can enter an
// implementation-dependent-optimized state while waiting for an event or store
// operation to the address range specified by MONITOR. 
//
//		
//
// Instruction: 'MWAIT'. Intrinsic: '_mm_mwait'.
// Requires MONITOR.
func Mwait(extensions uint, hints uint)  {
	mwait(extensions, hints)
}

func mwait(extensions uint, hints uint) 

