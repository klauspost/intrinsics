package adx

import . "github.com/klauspost/intrinsics/x86"


// AddcarryxU32: Add unsigned 32-bit integers 'a' and 'b' with unsigned 8-bit
// carry-in 'c_in' (carry or overflow flag), and store the unsigned 32-bit
// result in 'out', and the carry-out in 'dst' (carry or overflow flag). 
//
//		dst:out[31:0] := a[31:0] + b[31:0] + c_in;
//
// Instruction: 'ADCX, ADOX'. Intrinsic: '_addcarryx_u32'.
// Requires ADX.
func AddcarryxU32(c_in uint8, a uint32, b uint32, out uint32) uint8 {
	return uint8(addcarryxU32(c_in, a, b, out))
}

func addcarryxU32(c_in uint8, a uint32, b uint32, out uint32) uint8


// AddcarryxU64: Add unsigned 64-bit integers 'a' and 'b' with unsigned 8-bit
// carry-in 'c_in' (carry or overflow flag), and store the unsigned 64-bit
// result in 'out', and the carry-out in 'dst' (carry or overflow flag). 
//
//		dst:out[63:0] := a[63:0] + b[63:0] + c_in;
//
// Instruction: 'ADCX, ADOX'. Intrinsic: '_addcarryx_u64'.
// Requires ADX.
func AddcarryxU64(c_in uint8, a uint64, b uint64, out uint64) uint8 {
	return uint8(addcarryxU64(c_in, a, b, out))
}

func addcarryxU64(c_in uint8, a uint64, b uint64, out uint64) uint8

