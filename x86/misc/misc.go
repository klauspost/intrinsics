package misc

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AddcarryU32: Add unsigned 32-bit integers 'a' and 'b' with unsigned 8-bit
// carry-in 'c_in' (carry flag), and store the unsigned 32-bit result in 'out',
// and the carry-out in 'dst' (carry or overflow flag). 
//
//		dst:out[31:0] := a[31:0] + b[31:0] + c_in;
//
// Instruction: 'ADC'. Intrinsic: '_addcarry_u32'.
func AddcarryU32(c_in uint8, a uint32, b uint32, out uint32) uint8 {
	return uint8(addcarryU32(c_in, a, b, out))
}

func addcarryU32(c_in uint8, a uint32, b uint32, out uint32) uint8


// AddcarryU64: Add unsigned 64-bit integers 'a' and 'b' with unsigned 8-bit
// carry-in 'c_in' (carry flag), and store the unsigned 64-bit result in 'out',
// and the carry-out in 'dst' (carry or overflow flag). 
//
//		dst:out[63:0] := a[63:0] + b[63:0] + c_in;
//
// Instruction: 'ADC'. Intrinsic: '_addcarry_u64'.
func AddcarryU64(c_in uint8, a uint64, b uint64, out uint64) uint8 {
	return uint8(addcarryU64(c_in, a, b, out))
}

func addcarryU64(c_in uint8, a uint64, b uint64, out uint64) uint8


// AllowCpuFeatures: Treat the processor-specific feature(s) specified in 'a'
// as available. Multiple features may be OR'd together. See the valid feature
// flags below: 
//
//		_FEATURE_GENERIC_IA32
//		_FEATURE_FPU
//		_FEATURE_CMOV
//		_FEATURE_MMX
//		_FEATURE_FXSAVE
//		_FEATURE_SSE
//		_FEATURE_SSE2
//		_FEATURE_SSE3
//		_FEATURE_SSSE3
//		_FEATURE_SSE4_1
//		_FEATURE_SSE4_2
//		_FEATURE_MOVBE
//		_FEATURE_POPCNT
//		_FEATURE_PCLMULQDQ
//		_FEATURE_AES
//		_FEATURE_F16C
//		_FEATURE_AVX
//		_FEATURE_RDRND
//		_FEATURE_FMA
//		_FEATURE_BMI
//		_FEATURE_LZCNT
//		_FEATURE_HLE
//		_FEATURE_RTM
//		_FEATURE_AVX2
//		_FEATURE_KNCNI
//		_FEATURE_AVX512F
//		_FEATURE_ADX
//		_FEATURE_RDSEED
//		_FEATURE_AVX512ER
//		_FEATURE_AVX512PF
//		_FEATURE_AVX512CD
//		_FEATURE_SHA
//		_FEATURE_MPX
//
// Instruction: '...'. Intrinsic: '_allow_cpu_features'.
func AllowCpuFeatures(a uint64)  {
	allowCpuFeatures(a)
}

func allowCpuFeatures(a uint64) 


// BitScanForward: Set 'dst' to the index of the lowest set bit in 32-bit
// integer 'a'. If no bits are set in 'a' then 'dst' is undefined. 
//
//		tmp := 0
//		IF a = 0
//			dst := undefined
//		ELSE
//			DO WHILE ((tmp < 32) AND a[tmp] = 0)
//				tmp := tmp + 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'BSF'. Intrinsic: '_bit_scan_forward'.
func BitScanForward(a int) int {
	return int(bitScanForward(a))
}

func bitScanForward(a int) int


// BitScanReverse: Set 'dst' to the index of the highest set bit in 32-bit
// integer 'a'. If no bits are set in 'a' then 'dst' is undefined. 
//
//		tmp := 31
//		IF a = 0
//			dst := undefined
//		ELSE
//			DO WHILE ((tmp > 0) AND a[tmp] = 0)
//				tmp := tmp - 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'BSR'. Intrinsic: '_bit_scan_reverse'.
func BitScanReverse(a int) int {
	return int(bitScanReverse(a))
}

func bitScanReverse(a int) int


// BitScanForward1: Set 'index' to the index of the lowest set bit in 32-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 0
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp < 32) AND mask[tmp] = 0)
//				tmp := tmp + 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSF'. Intrinsic: '_BitScanForward'.
func BitScanForward1(index uint32, mask uint32) uint8 {
	return uint8(bitScanForward1(index, mask))
}

func bitScanForward1(index uint32, mask uint32) uint8


// BitScanForward64: Set 'index' to the index of the lowest set bit in 64-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 0
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp < 64) AND mask[tmp] = 0)
//				tmp := tmp + 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSF'. Intrinsic: '_BitScanForward64'.
func BitScanForward64(index uint32, mask uint64) uint8 {
	return uint8(bitScanForward64(index, mask))
}

func bitScanForward64(index uint32, mask uint64) uint8


// BitScanReverse1: Set 'index' to the index of the highest set bit in 32-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 31
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp > 0) AND mask[tmp] = 0)
//				tmp := tmp - 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSR'. Intrinsic: '_BitScanReverse'.
func BitScanReverse1(index uint32, mask uint32) uint8 {
	return uint8(bitScanReverse1(index, mask))
}

func bitScanReverse1(index uint32, mask uint32) uint8


// BitScanReverse64: Set 'index' to the index of the highest set bit in 64-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 31
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp > 0) AND mask[tmp] = 0)
//				tmp := tmp - 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSR'. Intrinsic: '_BitScanReverse64'.
func BitScanReverse64(index uint32, mask uint64) uint8 {
	return uint8(bitScanReverse64(index, mask))
}

func bitScanReverse64(index uint32, mask uint64) uint8


// Bittest: Return the bit at index 'b' of 32-bit integer 'a'. 
//
//		dst := a[b]
//
// Instruction: 'BT'. Intrinsic: '_bittest'.
func Bittest(a int32, b int32) uint8 {
	return uint8(bittest(a, b))
}

func bittest(a int32, b int32) uint8


// Bittest64: Return the bit at index 'b' of 64-bit integer 'a'. 
//
//		dst := a[b]
//
// Instruction: 'BT'. Intrinsic: '_bittest64'.
func Bittest64(a int64, b int64) uint8 {
	return uint8(bittest64(a, b))
}

func bittest64(a int64, b int64) uint8


// Bittestandcomplement: Return the bit at index 'b' of 32-bit integer 'a', and
// set that bit to its complement. 
//
//		dst := a[b]
//		a[b] := ~a[b]
//
// Instruction: 'BTC'. Intrinsic: '_bittestandcomplement'.
func Bittestandcomplement(a int32, b int32) uint8 {
	return uint8(bittestandcomplement(a, b))
}

func bittestandcomplement(a int32, b int32) uint8


// Bittestandcomplement64: Return the bit at index 'b' of 64-bit integer 'a',
// and set that bit to its complement. 
//
//		dst := a[b]
//		a[b] := ~a[b]
//
// Instruction: 'BTC'. Intrinsic: '_bittestandcomplement64'.
func Bittestandcomplement64(a int64, b int64) uint8 {
	return uint8(bittestandcomplement64(a, b))
}

func bittestandcomplement64(a int64, b int64) uint8


// Bittestandreset: Return the bit at index 'b' of 32-bit integer 'a', and set
// that bit to zero. 
//
//		dst := a[b]
//		a[b] := 0
//
// Instruction: 'BTR'. Intrinsic: '_bittestandreset'.
func Bittestandreset(a int32, b int32) uint8 {
	return uint8(bittestandreset(a, b))
}

func bittestandreset(a int32, b int32) uint8


// Bittestandreset64: Return the bit at index 'b' of 64-bit integer 'a', and
// set that bit to zero. 
//
//		dst := a[b]
//		a[b] := 0
//
// Instruction: 'BTR'. Intrinsic: '_bittestandreset64'.
func Bittestandreset64(a int64, b int64) uint8 {
	return uint8(bittestandreset64(a, b))
}

func bittestandreset64(a int64, b int64) uint8


// Bittestandset: Return the bit at index 'b' of 32-bit integer 'a', and set
// that bit to one. 
//
//		dst := a[b]
//		a[b] := 1
//
// Instruction: 'BTS'. Intrinsic: '_bittestandset'.
func Bittestandset(a int32, b int32) uint8 {
	return uint8(bittestandset(a, b))
}

func bittestandset(a int32, b int32) uint8


// Bittestandset64: Return the bit at index 'b' of 64-bit integer 'a', and set
// that bit to one. 
//
//		dst := a[b]
//		a[b] := 1
//
// Instruction: 'BTS'. Intrinsic: '_bittestandset64'.
func Bittestandset64(a int64, b int64) uint8 {
	return uint8(bittestandset64(a, b))
}

func bittestandset64(a int64, b int64) uint8


// Bswap: Reverse the byte order of 32-bit integer 'a', and store the result in
// 'dst'. This intrinsic is provided for conversion between little and big
// endian values. 
//
//		dst[7:0] := a[31:24]
//		dst[15:8] := a[23:16]
//		dst[23:16] := a[15:8]
//		dst[31:24] := a[7:0]
//
// Instruction: 'BSWAP'. Intrinsic: '_bswap'.
func Bswap(a int) int {
	return int(bswap(a))
}

func bswap(a int) int


// Bswap64: Reverse the byte order of 64-bit integer 'a', and store the result
// in 'dst'. This intrinsic is provided for conversion between little and big
// endian values. 
//
//		dst[7:0] := a[63:56]
//		dst[15:8] := a[55:48]
//		dst[23:16] := a[47:40]
//		dst[31:24] := a[39:32]
//		dst[39:32] := a[31:24]
//		dst[47:40] := a[23:16]
//		dst[55:48] := a[15:8]
//		dst[63:56] := a[7:0]
//
// Instruction: 'BSWAP'. Intrinsic: '_bswap64'.
func Bswap64(a int64) int64 {
	return int64(bswap64(a))
}

func bswap64(a int64) int64


// Castf32U32: Cast from type float to type unsigned __int32 without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castf32_u32'.
func Castf32U32(a float32) uint32 {
	return uint32(castf32U32(a))
}

func castf32U32(a float32) uint32


// Castf64U64: Cast from type double to type unsigned __int64 without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castf64_u64'.
func Castf64U64(a float64) uint64 {
	return uint64(castf64U64(a))
}

func castf64U64(a float64) uint64


// Castu32F32: Cast from type unsigned __int32 to type float without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castu32_f32'.
func Castu32F32(a uint32) float32 {
	return float32(castu32F32(a))
}

func castu32F32(a uint32) float32


// Castu64F64: Cast from type unsigned __int64 to type double without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castu64_f64'.
func Castu64F64(a uint64) float64 {
	return float64(castu64F64(a))
}

func castu64F64(a uint64) float64


// CvtshSs: Convert the half-precision (16-bit) floating-point value 'a' to a
// single-precision (32-bit) floating-point value, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP16_To_FP32(a[15:0])
//
// Instruction: '...'. Intrinsic: '_cvtsh_ss'.
func CvtshSs(a uint16) float32 {
	return float32(cvtshSs(a))
}

func cvtshSs(a uint16) float32


// CvtssSh: Convert the single-precision (32-bit) floating-point value 'a' to a
// half-precision (16-bit) floating-point value, and store the result in 'dst'. 
//
//		dst[15:0] := Convert_FP32_To_FP16(a[31:0])
//
// Instruction: '...'. Intrinsic: '_cvtss_sh'.
func CvtssSh(a float32, imm8 int) uint16 {
	return uint16(cvtssSh(a, imm8))
}

func cvtssSh(a float32, imm8 int) uint16


// LoadbeI16: Loads a big-endian word (16-bit) value from address 'ptr' and
// stores the result in 'dst'. 
//
//		addr := MEM[ptr]
//		FOR j := 0 to 1
//			i := j*8
//			dst[i+7:i] := addr[15-i:15-i-7]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_loadbe_i16'.
func LoadbeI16(ptr uintptr) int16 {
	return int16(loadbeI16(uintptr(ptr)))
}

func loadbeI16(ptr uintptr) int16


// LoadbeI32: Loads a big-endian double word (32-bit) value from address 'ptr'
// and stores the result in 'dst'. 
//
//		addr := MEM[ptr]
//		FOR j := 0 to 4
//			i := j*8
//			dst[i+7:i] := addr[31-i:31-i-7]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_loadbe_i32'.
func LoadbeI32(ptr uintptr) int {
	return int(loadbeI32(uintptr(ptr)))
}

func loadbeI32(ptr uintptr) int


// LoadbeI64: Loads a big-endian quad word (64-bit) value from address 'ptr'
// and stores the result in 'dst'. 
//
//		addr := MEM[ptr]
//		FOR j := 0 to 8
//			i := j*8
//			dst[i+7:i] := addr[63-i:63-i-7]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_loadbe_i64'.
func LoadbeI64(ptr uintptr) int64 {
	return int64(loadbeI64(uintptr(ptr)))
}

func loadbeI64(ptr uintptr) int64


// LoaduSi16: Load unaligned 16-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[15:0] := MEM[mem_addr+15:mem_addr]
//		dst[MAX:16] := 0
//
// Instruction: 'MOVZWL+MOVD'. Intrinsic: '_mm_loadu_si16'.
func LoaduSi16(mem_addr uintptr) x86.M128i {
	return x86.M128i(loaduSi16(uintptr(mem_addr)))
}

func loaduSi16(mem_addr uintptr) [16]byte


// LoaduSi32: Load unaligned 32-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[31:0] := MEM[mem_addr+31:mem_addr]
//		dst[MAX:32] := 0
//
// Instruction: 'MOVD'. Intrinsic: '_mm_loadu_si32'.
func LoaduSi32(mem_addr uintptr) x86.M128i {
	return x86.M128i(loaduSi32(uintptr(mem_addr)))
}

func loaduSi32(mem_addr uintptr) [16]byte


// LoaduSi64: Load unaligned 64-bit integer from memory into the first element
// of 'dst'. 
//
//		dst[63:0] := MEM[mem_addr+63:mem_addr]
//		dst[MAX:64] := 0
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_loadu_si64'.
func LoaduSi64(mem_addr uintptr) x86.M128i {
	return x86.M128i(loaduSi64(uintptr(mem_addr)))
}

func loaduSi64(mem_addr uintptr) [16]byte


// Lrotl: Shift the bits of unsigned 64-bit integer 'a' left by the number of
// bits specified in 'shift', rotating the most-significant bit to the
// least-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 63
//		DO WHILE (count > 0)
//			tmp[0] := dst[63]
//			dst := (dst << 1) OR tmp[0]
//			count := count - 1
//		OD
//
// Instruction: 'ROL'. Intrinsic: '_lrotl'.
func Lrotl(a uint32, shift int) uint32 {
	return uint32(lrotl(a, shift))
}

func lrotl(a uint32, shift int) uint32


// Lrotr: Shift the bits of unsigned 64-bit integer 'a' right by the number of
// bits specified in 'shift', rotating the least-significant bit to the
// most-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 63
//		DO WHILE (count > 0)
//			tmp[63] := dst[0]
//			dst := (dst >> 1) OR tmp[63]
//			count := count - 1
//		OD
//
// Instruction: 'ROR'. Intrinsic: '_lrotr'.
func Lrotr(a uint32, shift int) uint32 {
	return uint32(lrotr(a, shift))
}

func lrotr(a uint32, shift int) uint32


// MayIUseCpuFeature: Dynamically query the processor to determine if the
// processor-specific feature(s) specified in 'a' are available, and return
// true or false (1 or 0) if the set of features is available. Multiple
// features may be OR'd together. This intrinsic does not check the processor
// vendor. See the valid feature flags below: 
//
//		_FEATURE_GENERIC_IA32
//		_FEATURE_FPU
//		_FEATURE_CMOV
//		_FEATURE_MMX
//		_FEATURE_FXSAVE
//		_FEATURE_SSE
//		_FEATURE_SSE2
//		_FEATURE_SSE3
//		_FEATURE_SSSE3
//		_FEATURE_SSE4_1
//		_FEATURE_SSE4_2
//		_FEATURE_MOVBE
//		_FEATURE_POPCNT
//		_FEATURE_PCLMULQDQ
//		_FEATURE_AES
//		_FEATURE_F16C
//		_FEATURE_AVX
//		_FEATURE_RDRND
//		_FEATURE_FMA
//		_FEATURE_BMI
//		_FEATURE_LZCNT
//		_FEATURE_HLE
//		_FEATURE_RTM
//		_FEATURE_AVX2
//		_FEATURE_KNCNI
//		_FEATURE_AVX512F
//		_FEATURE_ADX
//		_FEATURE_RDSEED
//		_FEATURE_AVX512ER
//		_FEATURE_AVX512PF
//		_FEATURE_AVX512CD
//		_FEATURE_SHA
//		_FEATURE_MPX
//
// Instruction: '...'. Intrinsic: '_may_i_use_cpu_feature'.
func MayIUseCpuFeature(a uint64) int {
	return int(mayIUseCpuFeature(a))
}

func mayIUseCpuFeature(a uint64) int


// Rdpmc: Read the Performance Monitor Counter (PMC) specified by 'a', and
// store up to 64-bits in 'dst'. The width of performance counters is
// implementation specific. 
//
//		dst[63:0] := ReadPMC(a)
//
// Instruction: 'RDPMC'. Intrinsic: '_rdpmc'.
func Rdpmc(a int) int64 {
	return int64(rdpmc(a))
}

func rdpmc(a int) int64


// Rotl: Shift the bits of unsigned 32-bit integer 'a' left by the number of
// bits specified in 'shift', rotating the most-significant bit to the
// least-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 31
//		DO WHILE (count > 0)
//			tmp[0] := dst[31]
//			dst := (dst << 1) OR tmp[0]
//			count := count - 1
//		OD
//
// Instruction: 'ROL'. Intrinsic: '_rotl'.
func Rotl(a uint32, shift int) uint32 {
	return uint32(rotl(a, shift))
}

func rotl(a uint32, shift int) uint32


// Rotr: Shift the bits of unsigned 32-bit integer 'a' right by the number of
// bits specified in 'shift', rotating the least-significant bit to the
// most-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 31
//		DO WHILE (count > 0)
//			tmp[31] := dst[0]
//			dst := (dst >> 1) OR tmp[31]
//			count := count - 1
//		OD
//
// Instruction: 'ROR'. Intrinsic: '_rotr'.
func Rotr(a uint32, shift int) uint32 {
	return uint32(rotr(a, shift))
}

func rotr(a uint32, shift int) uint32


// Rotwl: Shift the bits of unsigned 16-bit integer 'a' left by the number of
// bits specified in 'shift', rotating the most-significant bit to the
// least-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 15
//		DO WHILE (count > 0)
//			tmp[0] := dst[15]
//			dst := (dst << 1) OR tmp[0]
//			count := count - 1
//		OD
//
// Instruction: 'ROL'. Intrinsic: '_rotwl'.
func Rotwl(a uint16, shift int) uint16 {
	return uint16(rotwl(a, shift))
}

func rotwl(a uint16, shift int) uint16


// Rotwr: Shift the bits of unsigned 16-bit integer 'a' right by the number of
// bits specified in 'shift', rotating the least-significant bit to the
// most-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 15
//		DO WHILE (count > 0)
//			tmp[15] := dst[0]
//			dst := (dst >> 1) OR tmp[15]
//			count := count - 1
//		OD
//
// Instruction: 'ROR'. Intrinsic: '_rotwr'.
func Rotwr(a uint16, shift int) uint16 {
	return uint16(rotwr(a, shift))
}

func rotwr(a uint16, shift int) uint16


// StorebeI16: Stores word-sized (16-bit) 'data' to address 'ptr' in big-endian
// format. 
//
//		addr := MEM[ptr]
//		FOR j := 0 to 1
//			i := j*8
//			addr[i+7:i] := data[15-i:15-i-7]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_storebe_i16'.
func StorebeI16(ptr uintptr, data int16)  {
	storebeI16(uintptr(ptr), data)
}

func storebeI16(ptr uintptr, data int16) 


// StorebeI32: Stores double word-sized (32-bit) 'data' to address 'ptr' in
// big-endian format. 
//
//		addr := MEM[ptr]
//		FOR j := 0 to 4
//			i := j*8
//			addr[i+7:i] := data[31-i:31-i-7]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_storebe_i32'.
func StorebeI32(ptr uintptr, data int)  {
	storebeI32(uintptr(ptr), data)
}

func storebeI32(ptr uintptr, data int) 


// StorebeI64: Stores quad word-sized (64-bit) 'data' to address 'ptr' in
// big-endian format. 
//
//		addr := MEM[ptr]
//		FOR j := 0 to 7
//			i := j*8
//			addr[i+7:i] := data[63-i:63-i-7]
//		ENDFOR
//
// Instruction: '...'. Intrinsic: '_storebe_i64'.
func StorebeI64(ptr uintptr, data int64)  {
	storebeI64(uintptr(ptr), data)
}

func storebeI64(ptr uintptr, data int64) 


// StoreuSi16: Store 16-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+15:mem_addr] := a[15:0]
//
// Instruction: 'MOVD+MOVW'. Intrinsic: '_mm_storeu_si16'.
func StoreuSi16(mem_addr uintptr, a x86.M128i)  {
	storeuSi16(uintptr(mem_addr), [16]byte(a))
}

func storeuSi16(mem_addr uintptr, a [16]byte) 


// StoreuSi32: Store 32-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+31:mem_addr] := a[31:0]
//
// Instruction: 'MOVD'. Intrinsic: '_mm_storeu_si32'.
func StoreuSi32(mem_addr uintptr, a x86.M128i)  {
	storeuSi32(uintptr(mem_addr), [16]byte(a))
}

func storeuSi32(mem_addr uintptr, a [16]byte) 


// StoreuSi64: Store 64-bit integer from the first element of 'a' into memory.
// 'mem_addr' does not need to be aligned on any particular boundary. 
//
//		MEM[mem_addr+63:mem_addr] := a[63:0]
//
// Instruction: 'MOVQ'. Intrinsic: '_mm_storeu_si64'.
func StoreuSi64(mem_addr uintptr, a x86.M128i)  {
	storeuSi64(uintptr(mem_addr), [16]byte(a))
}

func storeuSi64(mem_addr uintptr, a [16]byte) 


// SubborrowU32: Add unsigned 8-bit borrow 'b_in' (carry flag) to unsigned
// 32-bit integer 'a', and subtract the result from unsigned 32-bit integer
// 'b'. Store the unsigned 32-bit result in 'out', and the carry-out in 'dst'
// (carry or overflow flag). 
//
//		dst:out[31:0] := (b[31:0] - (a[31:0] + b_in));
//
// Instruction: 'SBB'. Intrinsic: '_subborrow_u32'.
func SubborrowU32(b_in uint8, a uint32, b uint32, out uint32) uint8 {
	return uint8(subborrowU32(b_in, a, b, out))
}

func subborrowU32(b_in uint8, a uint32, b uint32, out uint32) uint8


// SubborrowU64: Add unsigned 8-bit borrow 'b_in' (carry flag) to unsigned
// 64-bit integer 'a', and subtract the result from unsigned 64-bit integer
// 'b'. Store the unsigned 64-bit result in 'out', and the carry-out in 'dst'
// (carry or overflow flag). 
//
//		dst:out[63:0] := (b[63:0] - (a[63:0] + b_in));
//
// Instruction: 'SBB'. Intrinsic: '_subborrow_u64'.
func SubborrowU64(b_in uint8, a uint64, b uint64, out uint64) uint8 {
	return uint8(subborrowU64(b_in, a, b, out))
}

func subborrowU64(b_in uint8, a uint64, b uint64, out uint64) uint8

