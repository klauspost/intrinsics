package invpcid

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Invpcid: Invalidate mappings in the Translation Lookaside Buffers (TLBs) and
// paging-structure caches for the processor context identifier (PCID)
// specified by 'descriptor' based on the invalidation type specified in
// 'type'. 
// 	The PCID 'descriptor' is specified as a 16-byte memory operand (with no
// alignment restrictions) where bits [11:0] specify the PCID, and bits
// [127:64] specify the linear address; bits [63:12] are reserved.
// 	The types supported are:
// 		0) Individual-address invalidation: If 'type' is 0, the logical processor
// invalidates mappings for a single linear address and tagged with the PCID
// specified in 'descriptor', except global translations. The instruction may
// also invalidate global translations, mappings for other linear addresses, or
// mappings tagged with other PCIDs.
// 		1) Single-context invalidation: If 'type' is 1, the logical processor
// invalidates all mappings tagged with the PCID specified in 'descriptor'
// except global translations. In some cases, it may invalidate mappings for
// other PCIDs as well.
// 		2) All-context invalidation: If 'type' is 2, the logical processor
// invalidates all mappings tagged with any PCID.
// 		3) All-context invalidation, retaining global translations: If 'type' is
// 3, the logical processor invalidates all mappings tagged with any PCID
// except global translations, ignoring 'descriptor'. The instruction may also
// invalidate global translations as well. 
//
//		CASE type OF
//		0: // individual-address invalidation retaining global translations
//			OP_PCID := descriptor[11:0]
//			ADDR := descriptor[127:64]
//			BREAK
//		1: // single PCID invalidation retaining globals
//			OP_PCID := descriptor[11:0]
//			// invalidate all mappings tagged with OP_PCID except global translations
//			BREAK
//		2: // all PCID invalidation
//			// invalidate all mappings tagged with any PCID
//			BREAK
//		3: // all PCID invalidation retaining global translations
//			// invalidate all mappings tagged with any PCID except global translations
//			BREAK
//		ESAC
//
// Instruction: 'INVPCID'. Intrinsic: '_invpcid'.
// Requires INVPCID.
func Invpcid(typ uint32, descriptor uintptr)  {
	invpcid(typ, uintptr(descriptor))
}

func invpcid(typ uint32, descriptor uintptr) 

