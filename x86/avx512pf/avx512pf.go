package avx512pf

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// M512MaskPrefetchI32gatherPd: Prefetch double-precision (64-bit)
// floating-point elements from memory using 32-bit indices. 64-bit elements
// are loaded from addresses starting at 'base_addr' and offset by each 32-bit
// element in 'vindex' (each index is scaled by the factor in 'scale').
// Gathered elements are merged in cache using writemask 'k' (elements are
// brought into cache only when their corresponding mask bits are set). 'scale'
// should be 1, 2, 4 or 8. The 'hint' parameter may be 1 (_MM_HINT_T0) for
// prefetching to L1 cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache. 
//
//		FOR j := 0 to 7
//			i := j*32;
//			IF mask[j] THEN
//				Prefetch([base_addr + SignExtend(vindex[i*31:i]) * scale], hint, RFO=0);
//			FI
//		ENDFOR;
//
// Instruction: 'VGATHERPF0DPD, VGATHERPF1DPD'. Intrinsic: '_mm512_mask_prefetch_i32gather_pd'.
// Requires AVX512PF.
func M512MaskPrefetchI32gatherPd(vindex x86.M256i, mask x86.Mmask8, base_addr uintptr, scale int, hint int)  {
	m512MaskPrefetchI32gatherPd([32]byte(vindex), uint8(mask), uintptr(base_addr), scale, hint)
}

func m512MaskPrefetchI32gatherPd(vindex [32]byte, mask uint8, base_addr uintptr, scale int, hint int) 


// M512PrefetchI32gatherPd: Prefetch double-precision (64-bit) floating-point
// elements from memory using 32-bit indices. 64-bit elements are loaded from
// addresses starting at 'base_addr' and offset by each 32-bit element in
// 'vindex' (each index is scaled by the factor in 'scale'). Gathered elements
// are merged in cache. 'scale' should be 1, 2, 4 or 8. The 'hint' parameter
// may be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 
//
//		FOR j := 0 to 7
//			i := j*32;
//			Prefetch([base_addr + SignExtend(vindex[i*31:i]) * scale], hint, RFO=0);
//		ENDFOR;
//
// Instruction: 'VGATHERPF0DPD, VGATHERPF1DPD'. Intrinsic: '_mm512_prefetch_i32gather_pd'.
// Requires AVX512PF.
func M512PrefetchI32gatherPd(vindex x86.M256i, base_addr uintptr, scale int, hint int)  {
	m512PrefetchI32gatherPd([32]byte(vindex), uintptr(base_addr), scale, hint)
}

func m512PrefetchI32gatherPd(vindex [32]byte, base_addr uintptr, scale int, hint int) 


// M512MaskPrefetchI32scatterPd: Prefetch double-precision (64-bit)
// floating-point elements with intent to write using 32-bit indices. The
// 'hint' parameter may be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2
// (_MM_HINT_T1) for prefetching to L2 cache. 64-bit elements are brought into
// cache from addresses starting at 'base_addr' and offset by each 32-bit
// element in 'vindex' (each index is scaled by the factor in 'scale') subject
// to mask 'k' (elements are not brought into cache when the corresponding mask
// bit is not set). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 TO 7
//			i := j*32;
//			IF mask[j] THEN
//				Prefetch(base_addr + SignExtend(vindex[i+31:i]) * scale], Level=hint, RFO=1);
//			FI
//		ENDFOR;
//
// Instruction: 'VSCATTERPF0DPD, VSCATTERPF1DPD'. Intrinsic: '_mm512_mask_prefetch_i32scatter_pd'.
// Requires AVX512PF.
func M512MaskPrefetchI32scatterPd(base_addr uintptr, mask x86.Mmask8, vinde x86.M256i, scale int, hint int)  {
	m512MaskPrefetchI32scatterPd(uintptr(base_addr), uint8(mask), [32]byte(vinde), scale, hint)
}

func m512MaskPrefetchI32scatterPd(base_addr uintptr, mask uint8, vinde [32]byte, scale int, hint int) 


// M512PrefetchI32scatterPd: Prefetch double-precision (64-bit) floating-point
// elements with intent to write using 32-bit indices. The 'hint' parameter may
// be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 64-bit elements are brought into cache from
// addresses starting at 'base_addr' and offset by each 32-bit element in
// 'vindex' (each index is scaled by the factor in 'scale'). 'scale' should be
// 1, 2, 4 or 8. 
//
//		FOR j := 0 TO 7
//			i := j*32;
//			Prefetch(base_addr + SignExtend(vindex[i+31:i]) * scale], Level=hint, RFO=1);
//		ENDFOR;
//
// Instruction: 'VSCATTERPF0DPD, VSCATTERPF1DPD'. Intrinsic: '_mm512_prefetch_i32scatter_pd'.
// Requires AVX512PF.
func M512PrefetchI32scatterPd(base_addr uintptr, vindex x86.M256i, scale int, hint int)  {
	m512PrefetchI32scatterPd(uintptr(base_addr), [32]byte(vindex), scale, hint)
}

func m512PrefetchI32scatterPd(base_addr uintptr, vindex [32]byte, scale int, hint int) 


// M512MaskPrefetchI64gatherPd: Prefetch double-precision (64-bit)
// floating-point elements from memory into cache level specified by 'hint'
// using 64-bit indices. 64-bit elements are loaded from addresses starting at
// 'base_addr' and offset by each 64-bit element in 'vindex' (each index is
// scaled by the factor in 'scale'). Prefetched elements are merged in cache
// using writemask 'k' (elements are copied from memory when the corresponding
// mask bit is set). 'scale' should be 1, 2, 4 or 8. The 'hint' parameter may
// be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF mask[j] THEN
//				Prefetch([base_addr + SignExtend(vindex[i*63:i] * scale]), Level=hint, RFO=0);
//			FI
//		ENDFOR;
//
// Instruction: 'VGATHERPF0QPD, VGATHERPF1QPD'. Intrinsic: '_mm512_mask_prefetch_i64gather_pd'.
// Requires AVX512PF.
func M512MaskPrefetchI64gatherPd(vindex x86.M512i, mask x86.Mmask8, base_addr uintptr, scale int, hint int)  {
	m512MaskPrefetchI64gatherPd([64]byte(vindex), uint8(mask), uintptr(base_addr), scale, hint)
}

func m512MaskPrefetchI64gatherPd(vindex [64]byte, mask uint8, base_addr uintptr, scale int, hint int) 


// M512PrefetchI64gatherPd: Prefetch double-precision (64-bit) floating-point
// elements from memory into cache level specified by 'hint' using 64-bit
// indices. 64-bit elements are loaded from addresses starting at 'base_addr'
// and offset by each 64-bit element in 'vindex' (each index is scaled by the
// factor in 'scale'). 'scale' should be 1, 2, 4 or 8. The 'hint' parameter may
// be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			Prefetch([base_addr + SignExtend(vindex[i*63:i] * scale]), Level=hint, RFO=0);
//		ENDFOR;
//
// Instruction: 'VGATHERPF0QPD, VGATHERPF1QPD'. Intrinsic: '_mm512_prefetch_i64gather_pd'.
// Requires AVX512PF.
func M512PrefetchI64gatherPd(vindex x86.M512i, base_addr uintptr, scale int, hint int)  {
	m512PrefetchI64gatherPd([64]byte(vindex), uintptr(base_addr), scale, hint)
}

func m512PrefetchI64gatherPd(vindex [64]byte, base_addr uintptr, scale int, hint int) 


// M512MaskPrefetchI64gatherPs: Prefetch single-precision (32-bit)
// floating-point elements from memory using 64-bit indices. 32-bit elements
// are loaded from addresses starting at 'base_addr' and offset by each 64-bit
// element in 'vindex' (each index is scaled by the factor in 'scale').
// Gathered elements are merged in cache using writemask 'k' (elements are only
// brought into cache when their corresponding mask bit is set). 'scale' should
// be 1, 2, 4 or 8.. The 'hint' parameter may be 1 (_MM_HINT_T0) for
// prefetching to L1 cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache. 
//
//		FOR j:= 0 to 7
//			i := j*64;
//			IF mask[j] THEN
//				Prefetch([base_addr + SignExtend(vindex[i+63:i]) * scale], hint, RFO=0);
//			FI
//		ENDFOR;
//
// Instruction: 'VGATHERPF0QPS, VGATHERPF1QPS'. Intrinsic: '_mm512_mask_prefetch_i64gather_ps'.
// Requires AVX512PF.
func M512MaskPrefetchI64gatherPs(vindex x86.M512i, mask x86.Mmask8, base_addr uintptr, scale int, hint int)  {
	m512MaskPrefetchI64gatherPs([64]byte(vindex), uint8(mask), uintptr(base_addr), scale, hint)
}

func m512MaskPrefetchI64gatherPs(vindex [64]byte, mask uint8, base_addr uintptr, scale int, hint int) 


// M512PrefetchI64gatherPs: Prefetch single-precision (32-bit) floating-point
// elements from memory using 64-bit indices. 32-bit elements are loaded from
// addresses starting at 'base_addr' and offset by each 64-bit element in
// 'vindex' (each index is scaled by the factor in 'scale'). Gathered elements
// are merged in cache. 'scale' should be 1, 2, 4 or 8. The 'hint' parameter
// may be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2 (_MM_HINT_T1) for
// prefetching to L2 cache. 
//
//		FOR j:= 0 to 7
//			i := j*64;
//			Prefetch([base_addr + SignExtend(vindex[i+63:i]) * scale], hint, RFO=0);
//		ENDFOR;
//
// Instruction: 'VGATHERPF0QPS, VGATHERPF1QPS'. Intrinsic: '_mm512_prefetch_i64gather_ps'.
// Requires AVX512PF.
func M512PrefetchI64gatherPs(vindex x86.M512i, base_addr uintptr, scale int, hint int)  {
	m512PrefetchI64gatherPs([64]byte(vindex), uintptr(base_addr), scale, hint)
}

func m512PrefetchI64gatherPs(vindex [64]byte, base_addr uintptr, scale int, hint int) 


// M512MaskPrefetchI64scatterPd: Prefetch double-precision (64-bit)
// floating-point elements with intent to write into memory using 64-bit
// indices. The 'hint' parameter may be 1 (_MM_HINT_T0) for prefetching to L1
// cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache. 64-bit elements are
// brought into cache from addresses starting at 'base_addr' and offset by each
// 64-bit element in 'vindex' (each index is scaled by the factor in 'scale')
// subject to mask 'k' (elements are not brought into cache when the
// corresponding mask bit is not set). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF mask[j] THEN
//				Prefetch([base_addr + SignExtend(vindex[i+63:i]) * scale], Level=hint, RFO=1);
//			FI
//		ENDFOR;
//
// Instruction: 'VSCATTERPF0QPD, VSCATTERPF1QPD'. Intrinsic: '_mm512_mask_prefetch_i64scatter_pd'.
// Requires AVX512PF.
func M512MaskPrefetchI64scatterPd(base_addr uintptr, mask x86.Mmask8, vindex x86.M512i, scale int, hint int)  {
	m512MaskPrefetchI64scatterPd(uintptr(base_addr), uint8(mask), [64]byte(vindex), scale, hint)
}

func m512MaskPrefetchI64scatterPd(base_addr uintptr, mask uint8, vindex [64]byte, scale int, hint int) 


// M512PrefetchI64scatterPd: Prefetch double-precision (64-bit) floating-point
// elements with intent to write into memory using 64-bit indices. The 'hint'
// parameter may be 1 (_MM_HINT_T0) for prefetching to L1 cache, or 2
// (_MM_HINT_T1) for prefetching to L2 cache. 64-bit elements are brought into
// cache from addresses starting at 'base_addr' and offset by each 64-bit
// element in 'vindex' (each index is scaled by the factor in 'scale'). 'scale'
// should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			Prefetch([base_addr + SignExtend(vindex[i+63:i]) * scale], Level=hint, RFO=1);
//		ENDFOR;
//
// Instruction: 'VSCATTERPF0QPD, VSCATTERPF1QPD'. Intrinsic: '_mm512_prefetch_i64scatter_pd'.
// Requires AVX512PF.
func M512PrefetchI64scatterPd(base_addr uintptr, vindex x86.M512i, scale int, hint int)  {
	m512PrefetchI64scatterPd(uintptr(base_addr), [64]byte(vindex), scale, hint)
}

func m512PrefetchI64scatterPd(base_addr uintptr, vindex [64]byte, scale int, hint int) 


// M512MaskPrefetchI64scatterPs: Prefetch single-precision (32-bit)
// floating-point elements with intent to write into memory using 64-bit
// indices. The 'hint' parameter may be 1 (_MM_HINT_T0) for prefetching to L1
// cache, or 2 (_MM_HINT_T1) for prefetching to L2 cache. 32-bit elements are
// stored at addresses starting at 'base_addr' and offset by each 64-bit
// element in 'vindex' (each index is scaled by the factor in 'scale') subject
// to mask 'k' (elements are not brought into cache when the corresponding mask
// bit is not set). 'scale' should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			IF mask[j] THEN
//				Prefetch([base_addr + SignExtend(vindex[i+63:i]) * scale], Level=hint, RFO=1);
//			FI
//		ENDFOR;
//
// Instruction: 'VSCATTERPF0QPS, VSCATTERPF1QPS'. Intrinsic: '_mm512_mask_prefetch_i64scatter_ps'.
// Requires AVX512PF.
func M512MaskPrefetchI64scatterPs(base_addr uintptr, mask x86.Mmask8, vindex x86.M512i, scale int, hint int)  {
	m512MaskPrefetchI64scatterPs(uintptr(base_addr), uint8(mask), [64]byte(vindex), scale, hint)
}

func m512MaskPrefetchI64scatterPs(base_addr uintptr, mask uint8, vindex [64]byte, scale int, hint int) 


// M512PrefetchI64scatterPs: Prefetch single-precision (32-bit) floating-point
// elements with intent to write into memory using 64-bit indices. Elements are
// prefetched into cache level 'hint', where 'hint' is 0 or 1. 32-bit elements
// are stored at addresses starting at 'base_addr' and offset by each 64-bit
// element in 'vindex' (each index is scaled by the factor in 'scale'). 'scale'
// should be 1, 2, 4 or 8. 
//
//		FOR j := 0 to 7
//			i := j*64;
//			Prefetch([base_addr + SignExtend(vindex[i+63:i]) * scale], Level=hint, RFO=1);
//		ENDFOR;
//
// Instruction: 'VSCATTERPF0QPS, VSCATTERPF1QPS'. Intrinsic: '_mm512_prefetch_i64scatter_ps'.
// Requires AVX512PF.
func M512PrefetchI64scatterPs(base_addr uintptr, vindex x86.M512i, scale int, hint int)  {
	m512PrefetchI64scatterPs(uintptr(base_addr), [64]byte(vindex), scale, hint)
}

func m512PrefetchI64scatterPs(base_addr uintptr, vindex [64]byte, scale int, hint int) 

