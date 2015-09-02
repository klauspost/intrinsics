package mpx

import "github.com/klauspost/intrinsics/x86"


// BndChkPtrBounds: Checks if ['q', 'q' + 'size' - 1] is within the lower and
// upper bounds of 'q' and throws a #BR if not. 
//
//		IF (q + size - 1) < q.LB OR (q + size - 1) > q.UB THEN
//			#BR;
//		FI;
//
// Instruction: 'BNDCU, BNDCN'. Intrinsic: '_bnd_chk_ptr_bounds'.
// Requires MPX.
func BndChkPtrBounds(q uintptr, size int)  {
	bndChkPtrBounds(uintptr(q), size)
}

func bndChkPtrBounds(q uintptr, size int) 


// BndChkPtrLbounds: Checks if 'q' is within its lower bound, and throws a #BR
// if not. 
//
//		IF q < q.LB THEN
//			#BR;
//		FI;
//
// Instruction: 'BNDCL'. Intrinsic: '_bnd_chk_ptr_lbounds'.
// Requires MPX.
func BndChkPtrLbounds(q uintptr)  {
	bndChkPtrLbounds(uintptr(q))
}

func bndChkPtrLbounds(q uintptr) 


// BndChkPtrUbounds: Checks if 'q' is within its upper bound, and throws a #BR
// if not. 
//
//		IF q > q.UB THEN
//			#BR;
//		FI;
//
// Instruction: 'BNDCU, BNDCN'. Intrinsic: '_bnd_chk_ptr_ubounds'.
// Requires MPX.
func BndChkPtrUbounds(q uintptr)  {
	bndChkPtrUbounds(uintptr(q))
}

func bndChkPtrUbounds(q uintptr) 


// BndCopyPtrBounds: Make a pointer with the value of 'q' and bounds set to the
// bounds of 'r' (e.g. copy the bounds of 'r' to pointer 'q'), and store the
// result in 'dst'. 
//
//		dst := q;
//		dst.LB := r.LB;
//		dst.UB := r.UB;
//
// Instruction: '...'. Intrinsic: '_bnd_copy_ptr_bounds'.
// Requires MPX.
func BndCopyPtrBounds(q uintptr, r uintptr)  {
	bndCopyPtrBounds(uintptr(q), uintptr(r))
}

func bndCopyPtrBounds(q uintptr, r uintptr) 


// BndGetPtrLbound: Return the lower bound of 'q'. 
//
//		dst := q.LB
//
// Instruction: '...'. Intrinsic: '_bnd_get_ptr_lbound'.
// Requires MPX.
func BndGetPtrLbound(q uintptr) uintptr {
	return uintptr(bndGetPtrLbound(uintptr(q)))
}

func bndGetPtrLbound(q uintptr) uintptr


// BndGetPtrUbound: Return the upper bound of 'q'. 
//
//		dst := q.UB
//
// Instruction: '...'. Intrinsic: '_bnd_get_ptr_ubound'.
// Requires MPX.
func BndGetPtrUbound(q uintptr) uintptr {
	return uintptr(bndGetPtrUbound(uintptr(q)))
}

func bndGetPtrUbound(q uintptr) uintptr


// BndInitPtrBounds: Make a pointer with the value of 'q' and open bounds,
// which allow the pointer to access the entire virtual address space, and
// store the result in 'dst'. 
//
//		dst := q;
//		dst.LB := 0;
//		dst.UB := 0;
//
// Instruction: '...'. Intrinsic: '_bnd_init_ptr_bounds'.
// Requires MPX.
func BndInitPtrBounds(q uintptr)  {
	bndInitPtrBounds(uintptr(q))
}

func bndInitPtrBounds(q uintptr) 


// BndNarrowPtrBounds: Narrow the bounds for pointer 'q' to the intersection of
// the bounds of 'r' and the bounds ['q', 'q' + 'size' - 1], and store the
// result in 'dst'. 
//
//		dst := q;
//		IF r.LB > (q + size - 1) OR r.UB < q THEN
//			dst.LB := 1;
//			dst.UB := 0;
//		ELSE
//			dst.LB := MAX(r.LB, q);
//			dst.UB := MIN(r.UB, (q + size - 1));
//		FI;
//
// Instruction: '...'. Intrinsic: '_bnd_narrow_ptr_bounds'.
// Requires MPX.
func BndNarrowPtrBounds(q uintptr, r uintptr, size int)  {
	bndNarrowPtrBounds(uintptr(q), uintptr(r), size)
}

func bndNarrowPtrBounds(q uintptr, r uintptr, size int) 


// BndSetPtrBounds: Make a pointer with the value of 'srcmem' and bounds set to
// ['srcmem', 'srcmem' + 'size' - 1], and store the result in 'dst'. 
//
//		dst := srcmem;
//		dst.LB := srcmem.LB;
//		dst.UB := srcmem + size - 1;
//
// Instruction: 'BNDMK'. Intrinsic: '_bnd_set_ptr_bounds'.
// Requires MPX.
func BndSetPtrBounds(srcmem uintptr, size int)  {
	bndSetPtrBounds(uintptr(srcmem), size)
}

func bndSetPtrBounds(srcmem uintptr, size int) 


// BndStorePtrBounds: Stores the bounds of 'ptr_val' pointer in memory at
// address 'ptr_addr'. 
//
//		MEM[ptr_addr].LB := ptr_val.LB;
//		MEM[ptr_addr].UB := ptr_val.UB;
//
// Instruction: 'BNDSTX'. Intrinsic: '_bnd_store_ptr_bounds'.
// Requires MPX.
func BndStorePtrBounds(ptr_addr uintptr, ptr_val uintptr)  {
	bndStorePtrBounds(uintptr(ptr_addr), uintptr(ptr_val))
}

func bndStorePtrBounds(ptr_addr uintptr, ptr_val uintptr) 

