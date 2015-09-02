package monitor

import . "github.com/klauspost/intrinsics/x86"


// Monitor: Arm address monitoring hardware using the address specified in 'p'.
// A store to an address within the specified address range triggers the
// monitoring hardware. Specify optional extensions in 'extensions', and
// optional hints in 'hints'. 
//
//		
//
// Instruction: 'MONITOR'. Intrinsic: '_mm_monitor'.
// Requires MONITOR.
func Monitor(p uintptr, extensions uint, hints uint)  {
	monitor(uintptr(p), extensions, hints)
}

func monitor(p uintptr, extensions uint, hints uint) 


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

