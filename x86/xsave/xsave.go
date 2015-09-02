package xsave

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Xgetbv: Copy up to 64-bits from the value of the extended control register
// (XCR) specified by 'a' into 'dst'. Currently only XFEATURE_ENABLED_MASK XCR
// is supported. 
//
//		dst[63:0] := XCR[a]
//
// Instruction: 'XGETBV'. Intrinsic: '_xgetbv'.
// Requires XSAVE.
func Xgetbv(a uint32) uint64 {
	return uint64(xgetbv(a))
}

func xgetbv(a uint32) uint64


// Skipped: _xrstor. Contains pointer parameter.


// Skipped: _xrstor64. Contains pointer parameter.


// Skipped: _xsave. Contains pointer parameter.


// Skipped: _xsave64. Contains pointer parameter.


// Xsetbv: Copy 64-bits from 'val' to the extended control register (XCR)
// specified by 'a'. Currently only XFEATURE_ENABLED_MASK XCR is supported. 
//
//		XCR[a] := val[63:0]
//
// Instruction: 'XSETBV'. Intrinsic: '_xsetbv'.
// Requires XSAVE.
func Xsetbv(a uint32, val uint64)  {
	xsetbv(a, val)
}

func xsetbv(a uint32, val uint64) 

