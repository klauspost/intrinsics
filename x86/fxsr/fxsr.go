package fxsr

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Fxrstor: Reload the x87 FPU, MMX technology, XMM, and MXCSR registers from
// the 512-byte memory image at 'mem_addr'. This data should have been written
// to memory previously using the FXSAVE instruction, and in the same format as
// required by the operating mode. 'mem_addr' must be aligned on a 16-byte
// boundary. 
//
//		(x87 FPU, MMX, XMM7-XMM0, MXCSR) := Load(MEM[mem_addr])
//
// Instruction: 'FXRSTOR'. Intrinsic: '_fxrstor'.
// Requires FXSR.
func Fxrstor(mem_addr uintptr)  {
	fxrstor(uintptr(mem_addr))
}

func fxrstor(mem_addr uintptr) 


// Fxrstor64: Reload the x87 FPU, MMX technology, XMM, and MXCSR registers from
// the 512-byte memory image at 'mem_addr'. This data should have been written
// to memory previously using the FXSAVE64 instruction, and in the same format
// as required by the operating mode. 'mem_addr' must be aligned on a 16-byte
// boundary. 
//
//		(x87 FPU, MMX, XMM7-XMM0, MXCSR) := Load(MEM[mem_addr])
//
// Instruction: 'FXRSTOR64'. Intrinsic: '_fxrstor64'.
// Requires FXSR.
func Fxrstor64(mem_addr uintptr)  {
	fxrstor64(uintptr(mem_addr))
}

func fxrstor64(mem_addr uintptr) 


// Fxsave: Save the current state of the x87 FPU, MMX technology, XMM, and
// MXCSR registers to a 512-byte memory location at 'mem_addr'. The clayout of
// the 512-byte region depends on the operating mode. Bytes [511:464] are
// available for software use and will not be overwritten by the processor. 
//
//		MEM[mem_addr+511*8:mem_addr] := Fxsave(x87 FPU, MMX, XMM7-XMM0, MXCSR)
//
// Instruction: 'FXSAVE'. Intrinsic: '_fxsave'.
// Requires FXSR.
func Fxsave(mem_addr uintptr)  {
	fxsave(uintptr(mem_addr))
}

func fxsave(mem_addr uintptr) 


// Fxsave64: Save the current state of the x87 FPU, MMX technology, XMM, and
// MXCSR registers to a 512-byte memory location at 'mem_addr'. The layout of
// the 512-byte region depends on the operating mode. Bytes [511:464] are
// available for software use and will not be overwritten by the processor. 
//
//		MEM[mem_addr+511*8:mem_addr] := Fxsave64(x87 FPU, MMX, XMM7-XMM0, MXCSR)
//
// Instruction: 'FXSAVE64'. Intrinsic: '_fxsave64'.
// Requires FXSR.
func Fxsave64(mem_addr uintptr)  {
	fxsave64(uintptr(mem_addr))
}

func fxsave64(mem_addr uintptr) 

