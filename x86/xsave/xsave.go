package xsave

import "github.com/klauspost/intrinsics/x86"


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


// Xrstor: Perform a full or partial restore of the enabled processor states
// using the state information stored in memory at 'mem_addr'. State is
// restored based on bits [62:0] in 'rs_mask', 'XCR0', and
// 'mem_addr.HEADER.XSTATE_BV'. 'mem_addr' must be aligned on a 64-byte
// boundary. 
//
//		st_mask = mem_addr.HEADER.XSTATE_BV[62:0]
//		FOR i := 0 to 62
//			IF (rs_mask[i] AND XCR0[i])
//				IF st_mask[i]
//					CASE (i) OF
//					0: ProcessorState[x87 FPU] := mem_addr.FPUSSESave_Area[FPU]
//					1: ProcessorState[SSE] := mem_addr.FPUSSESaveArea[SSE]
//					DEFAULT: ProcessorState[i] := mem_addr.Ext_Save_Area[i]
//					ESAC
//				ELSE
//					// ProcessorExtendedState := Processor Supplied Values
//					CASE (i) OF
//					1: MXCSR := mem_addr.FPUSSESave_Area[SSE]
//					ESAC
//				FI
//			FI
//			i := i + 1
//		ENDFOR
//
// Instruction: 'XRSTOR'. Intrinsic: '_xrstor'.
// Requires XSAVE.
func Xrstor(mem_addr uintptr, rs_mask uint64)  {
	xrstor(uintptr(mem_addr), rs_mask)
}

func xrstor(mem_addr uintptr, rs_mask uint64) 


// Xrstor64: Perform a full or partial restore of the enabled processor states
// using the state information stored in memory at 'mem_addr'. State is
// restored based on bits [62:0] in 'rs_mask', 'XCR0', and
// 'mem_addr.HEADER.XSTATE_BV'. 'mem_addr' must be aligned on a 64-byte
// boundary. 
//
//		st_mask = mem_addr.HEADER.XSTATE_BV[62:0]
//		FOR i := 0 to 62
//			IF (rs_mask[i] AND XCR0[i])
//				IF st_mask[i]
//					CASE (i) OF
//					0: ProcessorState[x87 FPU] := mem_addr.FPUSSESave_Area[FPU]
//					1: ProcessorState[SSE] := mem_addr.FPUSSESaveArea[SSE]
//					DEFAULT: ProcessorState[i] := mem_addr.Ext_Save_Area[i]
//					ESAC
//				ELSE
//					// ProcessorExtendedState := Processor Supplied Values
//					CASE (i) OF
//					1: MXCSR := mem_addr.FPUSSESave_Area[SSE]
//					ESAC
//				FI
//			FI
//			i := i + 1
//		ENDFOR
//
// Instruction: 'XRSTOR64'. Intrinsic: '_xrstor64'.
// Requires XSAVE.
func Xrstor64(mem_addr uintptr, rs_mask uint64)  {
	xrstor64(uintptr(mem_addr), rs_mask)
}

func xrstor64(mem_addr uintptr, rs_mask uint64) 


// Xsave: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'. State is saved based on bits [62:0] in 'save_mask' and
// 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. 
//
//		mask[62:0] := save_mask[62:0] BITWISE AND XCR0[62:0]
//		FOR i := 0 to 62
//			IF mask[i]
//				CASE (i) OF
//				0: mem_addr.FPUSSESave_Area[FPU] := ProcessorState[x87 FPU]
//				1: mem_addr.FPUSSESaveArea[SSE] := ProcessorState[SSE]
//				DEFAULT: mem_addr.Ext_Save_Area[i] := ProcessorState[i]
//				ESAC
//				mem_addr.HEADER.XSTATE_BV[i] := INIT_FUNCTION[i]
//			FI
//			i := i + 1
//		ENDFOR
//
// Instruction: 'XSAVE'. Intrinsic: '_xsave'.
// Requires XSAVE.
func Xsave(mem_addr uintptr, save_mask uint64)  {
	xsave(uintptr(mem_addr), save_mask)
}

func xsave(mem_addr uintptr, save_mask uint64) 


// Xsave64: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'. State is saved based on bits [62:0] in 'save_mask' and
// 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. 
//
//		mask[62:0] := save_mask[62:0] BITWISE AND XCR0[62:0]
//		FOR i := 0 to 62
//			IF mask[i]
//				CASE (i) OF
//				0: mem_addr.FPUSSESave_Area[FPU] := ProcessorState[x87 FPU]
//				1: mem_addr.FPUSSESaveArea[SSE] := ProcessorState[SSE]
//				DEFAULT: mem_addr.Ext_Save_Area[i] := ProcessorState[i]
//				ESAC
//				mem_addr.HEADER.XSTATE_BV[i] := INIT_FUNCTION[i]
//			FI
//			i := i + 1
//		ENDFOR
//
// Instruction: 'XSAVE64'. Intrinsic: '_xsave64'.
// Requires XSAVE.
func Xsave64(mem_addr uintptr, save_mask uint64)  {
	xsave64(uintptr(mem_addr), save_mask)
}

func xsave64(mem_addr uintptr, save_mask uint64) 


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

