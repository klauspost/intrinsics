package xsaveopt

import . "github.com/klauspost/intrinsics/x86"


// Xsaveopt: Perform a full or partial save of the enabled processor states to
// memory at 'mem_addr'. State is saved based on bits [62:0] in 'save_mask' and
// 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. The hardware may
// optimize the manner in which data is saved. The performance of this
// instruction will be equal to or better than using the XSAVE instruction. 
//
//		mask[62:0] := save_mask[62:0] BITWISE AND XCR0[62:0]
//		FOR i := 0 to 62
//			IF mask[i]
//				CASE (i) OF
//				0: mem_addr.FPUSSESave_Area[FPU] := ProcessorState[x87 FPU]
//				1: mem_addr.FPUSSESaveArea[SSE] := ProcessorState[SSE]
//				2: mem_addr.EXT_SAVE_Area2[YMM] := ProcessorState[YMM]
//				DEFAULT: mem_addr.Ext_Save_Area[i] := ProcessorState[i]
//				ESAC
//				mem_addr.HEADER.XSTATE_BV[i] := INIT_FUNCTION[i]
//			FI
//			i := i + 1
//		ENDFOR
//
// Instruction: 'XSAVEOPT'. Intrinsic: '_xsaveopt'.
// Requires XSAVEOPT.
func Xsaveopt(mem_addr uintptr, save_mask uint64)  {
	xsaveopt(uintptr(mem_addr), save_mask)
}

func xsaveopt(mem_addr uintptr, save_mask uint64) 


// Xsaveopt64: Perform a full or partial save of the enabled processor states
// to memory at 'mem_addr'. State is saved based on bits [62:0] in 'save_mask'
// and 'XCR0'. 'mem_addr' must be aligned on a 64-byte boundary. The hardware
// may optimize the manner in which data is saved. The performance of this
// instruction will be equal to or better than using the XSAVE64 instruction. 
//
//		mask[62:0] := save_mask[62:0] BITWISE AND XCR0[62:0]
//		FOR i := 0 to 62
//			IF mask[i]
//				CASE (i) OF
//				0: mem_addr.FPUSSESave_Area[FPU] := ProcessorState[x87 FPU]
//				1: mem_addr.FPUSSESaveArea[SSE] := ProcessorState[SSE]
//				2: mem_addr.EXT_SAVE_Area2[YMM] := ProcessorState[YMM]
//				DEFAULT: mem_addr.Ext_Save_Area[i] := ProcessorState[i]
//				ESAC
//				mem_addr.HEADER.XSTATE_BV[i] := INIT_FUNCTION[i]
//			FI
//			i := i + 1
//		ENDFOR
//
// Instruction: 'XSAVEOPT64'. Intrinsic: '_xsaveopt64'.
// Requires XSAVEOPT.
func Xsaveopt64(mem_addr uintptr, save_mask uint64)  {
	xsaveopt64(uintptr(mem_addr), save_mask)
}

func xsaveopt64(mem_addr uintptr, save_mask uint64) 

