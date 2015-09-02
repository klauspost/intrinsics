package xss

import . "github.com/klauspost/intrinsics/x86"


// Xrstors: Perform a full or partial restore of the enabled processor states
// using the state information stored in memory at 'mem_addr'. xrstors differs
// from xrstor in that it can restore state components corresponding to bits
// set in the IA32_XSS MSR; xrstors cannot restore from an xsave area in which
// the extended region is in the standard form. State is restored based on bits
// [62:0] in 'rs_mask', 'XCR0', and 'mem_addr.HEADER.XSTATE_BV'. 'mem_addr'
// must be aligned on a 64-byte boundary. 
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
// Instruction: 'XRSTORS'. Intrinsic: '_xrstors'.
// Requires XSS.
func Xrstors(mem_addr uintptr, rs_mask uint64)  {
	xrstors(uintptr(mem_addr), rs_mask)
}

func xrstors(mem_addr uintptr, rs_mask uint64) 


// Xrstors64: Perform a full or partial restore of the enabled processor states
// using the state information stored in memory at 'mem_addr'. xrstors differs
// from xrstor in that it can restore state components corresponding to bits
// set in the IA32_XSS MSR; xrstors cannot restore from an xsave area in which
// the extended region is in the standard form. State is restored based on bits
// [62:0] in 'rs_mask', 'XCR0', and 'mem_addr.HEADER.XSTATE_BV'. 'mem_addr'
// must be aligned on a 64-byte boundary. 
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
// Instruction: 'XRSTORS64'. Intrinsic: '_xrstors64'.
// Requires XSS.
func Xrstors64(mem_addr uintptr, rs_mask uint64)  {
	xrstors64(uintptr(mem_addr), rs_mask)
}

func xrstors64(mem_addr uintptr, rs_mask uint64) 


// Xsaves: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'; xsaves differs from xsave in that it can save state
// components corresponding to bits set in IA32_XSS MSR and that it may use the
// modified optimization. State is saved based on bits [62:0] in 'save_mask'
// and 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. 
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
// Instruction: 'XSAVES'. Intrinsic: '_xsaves'.
// Requires XSS.
func Xsaves(mem_addr uintptr, save_mask uint64)  {
	xsaves(uintptr(mem_addr), save_mask)
}

func xsaves(mem_addr uintptr, save_mask uint64) 


// Xsaves64: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'; xsaves differs from xsave in that it can save state
// components corresponding to bits set in IA32_XSS MSR and that it may use the
// modified optimization. State is saved based on bits [62:0] in 'save_mask'
// and 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. 
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
// Instruction: 'XSAVEC64'. Intrinsic: '_xsaves64'.
// Requires XSS.
func Xsaves64(mem_addr uintptr, save_mask uint64)  {
	xsaves64(uintptr(mem_addr), save_mask)
}

func xsaves64(mem_addr uintptr, save_mask uint64) 

