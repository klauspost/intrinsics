// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package fsgsbase

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// ReadfsbaseU32: Read the FS segment base register and store the 32-bit result
// in 'dst'. 
//
//		dst[31:0] := FS_Segment_Base_Register;
//		dst[63:32] := 0
//
// Instruction: 'RDFSBASE'. Intrinsic: '_readfsbase_u32'.
// Requires FSGSBASE.
func ReadfsbaseU32() uint32 {
	panic("not implemented")
}


// ReadfsbaseU64: Read the FS segment base register and store the 64-bit result
// in 'dst'. 
//
//		dst[63:0] := FS_Segment_Base_Register;
//
// Instruction: 'RDFSBASE'. Intrinsic: '_readfsbase_u64'.
// Requires FSGSBASE.
func ReadfsbaseU64() uint64 {
	panic("not implemented")
}


// ReadgsbaseU32: Read the GS segment base register and store the 32-bit result
// in 'dst'. 
//
//		dst[31:0] := GS_Segment_Base_Register;
//		dst[63:32] := 0
//
// Instruction: 'RDGSBASE'. Intrinsic: '_readgsbase_u32'.
// Requires FSGSBASE.
func ReadgsbaseU32() uint32 {
	panic("not implemented")
}


// ReadgsbaseU64: Read the GS segment base register and store the 64-bit result
// in 'dst'. 
//
//		dst[63:0] := GS_Segment_Base_Register;
//
// Instruction: 'RDGSBASE'. Intrinsic: '_readgsbase_u64'.
// Requires FSGSBASE.
func ReadgsbaseU64() uint64 {
	panic("not implemented")
}


// WritefsbaseU32: Write the unsigned 32-bit integer 'a' to the FS segment base
// register. 
//
//		FS_Segment_Base_Register[31:0] := a[31:0];
//		FS_Segment_Base_Register[63:32] := 0
//
// Instruction: 'WRFSBASE'. Intrinsic: '_writefsbase_u32'.
// Requires FSGSBASE.
func WritefsbaseU32(a uint32)  {
	panic("not implemented")
}


// WritefsbaseU64: Write the unsigned 64-bit integer 'a' to the FS segment base
// register. 
//
//		FS_Segment_Base_Register[63:0] := a[63:0];
//
// Instruction: 'WRFSBASE'. Intrinsic: '_writefsbase_u64'.
// Requires FSGSBASE.
func WritefsbaseU64(a uint64)  {
	panic("not implemented")
}


// WritegsbaseU32: Write the unsigned 32-bit integer 'a' to the GS segment base
// register. 
//
//		GS_Segment_Base_Register[31:0] := a[31:0];
//		GS_Segment_Base_Register[63:32] := 0
//
// Instruction: 'WRGSBASE'. Intrinsic: '_writegsbase_u32'.
// Requires FSGSBASE.
func WritegsbaseU32(a uint32)  {
	panic("not implemented")
}


// WritegsbaseU64: Write the unsigned 64-bit integer 'a' to the GS segment base
// register. 
//
//		GS_Segment_Base_Register[63:0] := a[63:0];
//
// Instruction: 'WRGSBASE'. Intrinsic: '_writegsbase_u64'.
// Requires FSGSBASE.
func WritegsbaseU64(a uint64)  {
	panic("not implemented")
}

