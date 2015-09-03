package rtm

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// Xabort: Force an RTM abort. The EAX register is updated to reflect an XABORT
// instruction caused the abort, and the 'imm8' parameter will be provided in
// bits [31:24] of EAX.
// 	Following an RTM abort, the logical processor resumes execution at the
// fallback address computed through the outermost XBEGIN instruction. 
//
//		IF RTM_ACTIVE = 0
//			// nop
//		ELSE
//			// restore architectural register state
//			// discard memory updates performed in transaction
//			// update EAX with status and imm8 value
//			RTM_NEST_COUNT := 0
//			RTM_ACTIVE := 0
//			IF 64-bit Mode
//				RIP := fallbackRIP
//			ELSE
//				EIP := fallbackEIP
//			FI
//		FI
//
// Instruction: 'XABORT'. Intrinsic: '_xabort'.
// Requires RTM.
//
// FIXME: Requires compiler support (has immediate)
func Xabort(imm8 byte)  {
	xabort(imm8)
}

func xabort(imm8 byte) 


// Xbegin: Specify the start of an RTM code region. 
// 	If the logical processor was not already in transactional execution, then
// this call causes the logical processor to transition into transactional
// execution. 
// 	On an RTM abort, the logical processor discards all architectural register
// and memory updates performed during the RTM execution, restores
// architectural state, and starts execution beginning at the fallback address
// computed from the outermost XBEGIN instruction. 
//
//		IF RTM_NEST_COUNT < MAX_RTM_NEST_COUNT
//			RTM_NEST_COUNT := RTM_NEST_COUNT + 1
//			IF RTM_NEST_COUNT = 1
//				IF 64-bit Mode
//					fallbackRIP := RIP + SignExtend(IMM)
//				ELSE IF 32-bit Mode
//					fallbackEIP := EIP + SignExtend(IMM)
//				ELSE // 16-bit Mode
//					fallbackEIP := (EIP + SignExtend(IMM)) AND 0x0000FFFF
//				FI
//				
//				RTM_ACTIVE := 1
//				// enter RTM execution, record register state, start tracking memory state
//			FI
//		ELSE
//			// RTM abort (see _xabort)
//		FI
//
// Instruction: 'XBEGIN'. Intrinsic: '_xbegin'.
// Requires RTM.
func Xbegin() (dst uint32) {
	return uint32(xbegin())
}

func xbegin() uint32


// Xend: Specify the end of an RTM code region.
// 	If this corresponds to the outermost scope, the logical processor will
// attempt to commit the logical processor state atomically. 
// 	If the commit fails, the logical processor will perform an RTM abort. 
//
//		IF RTM_ACTIVE = 1
//			RTM_NEST_COUNT := RTM_NEST_COUNT - 1
//			IF RTM_NEST_COUNT = 0
//				// try to commit transaction
//				IF fail to commit transaction
//					// RTM abort (see _xabort)
//				ELSE
//					RTM_ACTIVE = 0
//				FI
//			FI
//		FI
//
// Instruction: 'XEND'. Intrinsic: '_xend'.
// Requires RTM.
func Xend()  {
	xend()
}

func xend() 


// Xtest: Query the transactional execution status, return 0 if inside a
// transactionally executing RTM or HLE region, and return 1 otherwise. 
//
//		IF (RTM_ACTIVE = 1 OR HLE_ACTIVE = 1)
//			dst := 0
//		ELSE
//			dst := 1
//		FI
//
// Instruction: 'XTEST'. Intrinsic: '_xtest'.
// Requires RTM.
func Xtest() (dst uint8) {
	return uint8(xtest())
}

func xtest() uint8

