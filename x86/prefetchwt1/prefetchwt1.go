// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package prefetchwt1

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'PREFETCHWT1'. Intrinsic: '_mm_prefetch'.
// Requires PREFETCHWT1.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Prefetch(p *byte, i int)  {
	panic("not implemented")
}

