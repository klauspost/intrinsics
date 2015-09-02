package xsavec

import . "github.com/klauspost/intrinsics/x86"


// Xsavec: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'; xsavec differs from xsave in that it uses compaction
// and that it may use init optimization. State is saved based on bits [62:0]
// in 'save_mask' and 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. 
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
// Instruction: 'XSAVEC'. Intrinsic: '_xsavec'.
// Requires XSAVEC.
func Xsavec(mem_addr uintptr, save_mask uint64)  {
	xsavec(uintptr(mem_addr), save_mask)
}

func xsavec(mem_addr uintptr, save_mask uint64) 


// Xsavec64: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'; xsavec differs from xsave in that it uses compaction
// and that it may use init optimization. State is saved based on bits [62:0]
// in 'save_mask' and 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. 
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
// Instruction: 'XSAVEC64'. Intrinsic: '_xsavec64'.
// Requires XSAVEC.
func Xsavec64(mem_addr uintptr, save_mask uint64)  {
	xsavec64(uintptr(mem_addr), save_mask)
}

func xsavec64(mem_addr uintptr, save_mask uint64) 

