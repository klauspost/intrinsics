package clflushopt

import "github.com/klauspost/intrinsics/x86"


// Clflushopt: Invalidate and flush the cache line that contains 'p' from all
// levels of the cache hierarchy. 
//
//		
//
// Instruction: 'CLFLUSHOPT'. Intrinsic: '_mm_clflushopt'.
// Requires CLFLUSHOPT.
func Clflushopt(p uintptr)  {
	clflushopt(uintptr(p))
}

func clflushopt(p uintptr) 

