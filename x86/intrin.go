package x86


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


// BextrU32: Extract contiguous bits from unsigned 32-bit integer 'a', and
// store the result in 'dst'. Extract the number of bits specified by 'len',
// starting at the bit specified by 'start'. 
//
//		tmp := ZERO_EXTEND_TO_512(a)
//		dst := ZERO_EXTEND(tmp[start+len-1:start])
//
// Instruction: 'BEXTR'. Intrinsic: '_bextr_u32'.
// Requires BMI1.
func BextrU32(a uint32, start uint32, len uint32) uint32 {
	return uint32(bextrU32(a, start, len))
}

func bextrU32(a uint32, start uint32, len uint32) uint32


// BextrU64: Extract contiguous bits from unsigned 64-bit integer 'a', and
// store the result in 'dst'. Extract the number of bits specified by 'len',
// starting at the bit specified by 'start'. 
//
//		tmp := ZERO_EXTEND_TO_512(a)
//		dst := ZERO_EXTEND(tmp[start+len-1:start])
//
// Instruction: 'BEXTR'. Intrinsic: '_bextr_u64'.
// Requires BMI1.
func BextrU64(a uint64, start uint32, len uint32) uint64 {
	return uint64(bextrU64(a, start, len))
}

func bextrU64(a uint64, start uint32, len uint32) uint64


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


// BlsiU32: Extract the lowest set bit from unsigned 32-bit integer 'a' and set
// the corresponding bit in 'dst'. All other bits in 'dst' are zeroed, and all
// bits are zeroed if no bits are set in 'a'. 
//
//		dst := (-a) BITWISE AND a
//
// Instruction: 'BLSI'. Intrinsic: '_blsi_u32'.
// Requires BMI1.
func BlsiU32(a uint32) uint32 {
	return uint32(blsiU32(a))
}

func blsiU32(a uint32) uint32


// BlsiU64: Extract the lowest set bit from unsigned 64-bit integer 'a' and set
// the corresponding bit in 'dst'. All other bits in 'dst' are zeroed, and all
// bits are zeroed if no bits are set in 'a'. 
//
//		dst := (-a) BITWISE AND a
//
// Instruction: 'BLSI'. Intrinsic: '_blsi_u64'.
// Requires BMI1.
func BlsiU64(a uint64) uint64 {
	return uint64(blsiU64(a))
}

func blsiU64(a uint64) uint64


// BlsmskU32: Set all the lower bits of 'dst' up to and including the lowest
// set bit in unsigned 32-bit integer 'a'. 
//
//		dst := (a - 1) XOR a
//
// Instruction: 'BLSMSK'. Intrinsic: '_blsmsk_u32'.
// Requires BMI1.
func BlsmskU32(a uint32) uint32 {
	return uint32(blsmskU32(a))
}

func blsmskU32(a uint32) uint32


// BlsmskU64: Set all the lower bits of 'dst' up to and including the lowest
// set bit in unsigned 64-bit integer 'a'. 
//
//		dst := (a - 1) XOR a
//
// Instruction: 'BLSMSK'. Intrinsic: '_blsmsk_u64'.
// Requires BMI1.
func BlsmskU64(a uint64) uint64 {
	return uint64(blsmskU64(a))
}

func blsmskU64(a uint64) uint64


// BlsrU32: Copy all bits from unsigned 32-bit integer 'a' to 'dst', and reset
// (set to 0) the bit in 'dst' that corresponds to the lowest set bit in 'a'. 
//
//		dst := (a - 1) BITWISE AND a
//
// Instruction: 'BLSR'. Intrinsic: '_blsr_u32'.
// Requires BMI1.
func BlsrU32(a uint32) uint32 {
	return uint32(blsrU32(a))
}

func blsrU32(a uint32) uint32


// BlsrU64: Copy all bits from unsigned 64-bit integer 'a' to 'dst', and reset
// (set to 0) the bit in 'dst' that corresponds to the lowest set bit in 'a'. 
//
//		dst := (a - 1) BITWISE AND a
//
// Instruction: 'BLSR'. Intrinsic: '_blsr_u64'.
// Requires BMI1.
func BlsrU64(a uint64) uint64 {
	return uint64(blsrU64(a))
}

func blsrU64(a uint64) uint64


// BndChkPtrBounds: Checks if ['q', 'q' + 'size' - 1] is within the lower and
// upper bounds of 'q' and throws a #BR if not. 
//
//		IF (q + size - 1) < q.LB OR (q + size - 1) > q.UB THEN
//			#BR;
//		FI;
//
// Instruction: 'BNDCU, BNDCN'. Intrinsic: '_bnd_chk_ptr_bounds'.
// Requires MPX.
func BndChkPtrBounds(q uintptr, size int)  {
	bndChkPtrBounds(uintptr(q), size)
}

func bndChkPtrBounds(q uintptr, size int) 


// BndChkPtrLbounds: Checks if 'q' is within its lower bound, and throws a #BR
// if not. 
//
//		IF q < q.LB THEN
//			#BR;
//		FI;
//
// Instruction: 'BNDCL'. Intrinsic: '_bnd_chk_ptr_lbounds'.
// Requires MPX.
func BndChkPtrLbounds(q uintptr)  {
	bndChkPtrLbounds(uintptr(q))
}

func bndChkPtrLbounds(q uintptr) 


// BndChkPtrUbounds: Checks if 'q' is within its upper bound, and throws a #BR
// if not. 
//
//		IF q > q.UB THEN
//			#BR;
//		FI;
//
// Instruction: 'BNDCU, BNDCN'. Intrinsic: '_bnd_chk_ptr_ubounds'.
// Requires MPX.
func BndChkPtrUbounds(q uintptr)  {
	bndChkPtrUbounds(uintptr(q))
}

func bndChkPtrUbounds(q uintptr) 


// BndCopyPtrBounds: Make a pointer with the value of 'q' and bounds set to the
// bounds of 'r' (e.g. copy the bounds of 'r' to pointer 'q'), and store the
// result in 'dst'. 
//
//		dst := q;
//		dst.LB := r.LB;
//		dst.UB := r.UB;
//
// Instruction: '...'. Intrinsic: '_bnd_copy_ptr_bounds'.
// Requires MPX.
func BndCopyPtrBounds(q uintptr, r uintptr)  {
	bndCopyPtrBounds(uintptr(q), uintptr(r))
}

func bndCopyPtrBounds(q uintptr, r uintptr) 


// BndGetPtrLbound: Return the lower bound of 'q'. 
//
//		dst := q.LB
//
// Instruction: '...'. Intrinsic: '_bnd_get_ptr_lbound'.
// Requires MPX.
func BndGetPtrLbound(q uintptr) uintptr {
	return uintptr(bndGetPtrLbound(uintptr(q)))
}

func bndGetPtrLbound(q uintptr) uintptr


// BndGetPtrUbound: Return the upper bound of 'q'. 
//
//		dst := q.UB
//
// Instruction: '...'. Intrinsic: '_bnd_get_ptr_ubound'.
// Requires MPX.
func BndGetPtrUbound(q uintptr) uintptr {
	return uintptr(bndGetPtrUbound(uintptr(q)))
}

func bndGetPtrUbound(q uintptr) uintptr


// BndInitPtrBounds: Make a pointer with the value of 'q' and open bounds,
// which allow the pointer to access the entire virtual address space, and
// store the result in 'dst'. 
//
//		dst := q;
//		dst.LB := 0;
//		dst.UB := 0;
//
// Instruction: '...'. Intrinsic: '_bnd_init_ptr_bounds'.
// Requires MPX.
func BndInitPtrBounds(q uintptr)  {
	bndInitPtrBounds(uintptr(q))
}

func bndInitPtrBounds(q uintptr) 


// BndNarrowPtrBounds: Narrow the bounds for pointer 'q' to the intersection of
// the bounds of 'r' and the bounds ['q', 'q' + 'size' - 1], and store the
// result in 'dst'. 
//
//		dst := q;
//		IF r.LB > (q + size - 1) OR r.UB < q THEN
//			dst.LB := 1;
//			dst.UB := 0;
//		ELSE
//			dst.LB := MAX(r.LB, q);
//			dst.UB := MIN(r.UB, (q + size - 1));
//		FI;
//
// Instruction: '...'. Intrinsic: '_bnd_narrow_ptr_bounds'.
// Requires MPX.
func BndNarrowPtrBounds(q uintptr, r uintptr, size int)  {
	bndNarrowPtrBounds(uintptr(q), uintptr(r), size)
}

func bndNarrowPtrBounds(q uintptr, r uintptr, size int) 


// BndSetPtrBounds: Make a pointer with the value of 'srcmem' and bounds set to
// ['srcmem', 'srcmem' + 'size' - 1], and store the result in 'dst'. 
//
//		dst := srcmem;
//		dst.LB := srcmem.LB;
//		dst.UB := srcmem + size - 1;
//
// Instruction: 'BNDMK'. Intrinsic: '_bnd_set_ptr_bounds'.
// Requires MPX.
func BndSetPtrBounds(srcmem uintptr, size int)  {
	bndSetPtrBounds(uintptr(srcmem), size)
}

func bndSetPtrBounds(srcmem uintptr, size int) 


// BndStorePtrBounds: Stores the bounds of 'ptr_val' pointer in memory at
// address 'ptr_addr'. 
//
//		MEM[ptr_addr].LB := ptr_val.LB;
//		MEM[ptr_addr].UB := ptr_val.UB;
//
// Instruction: 'BNDSTX'. Intrinsic: '_bnd_store_ptr_bounds'.
// Requires MPX.
func BndStorePtrBounds(ptr_addr uintptr, ptr_val uintptr)  {
	bndStorePtrBounds(uintptr(ptr_addr), uintptr(ptr_val))
}

func bndStorePtrBounds(ptr_addr uintptr, ptr_val uintptr) 


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


// BzhiU32: Copy all bits from unsigned 32-bit integer 'a' to 'dst', and reset
// (set to 0) the high bits in 'dst' starting at 'index'. 
//
//		n := index[7:0]
//		dst := a
//		IF (n < 32)
//			dst[31:n] := 0
//		FI
//
// Instruction: 'BZHI'. Intrinsic: '_bzhi_u32'.
// Requires BMI2.
func BzhiU32(a uint32, index uint32) uint32 {
	return uint32(bzhiU32(a, index))
}

func bzhiU32(a uint32, index uint32) uint32


// BzhiU64: Copy all bits from unsigned 64-bit integer 'a' to 'dst', and reset
// (set to 0) the high bits in 'dst' starting at 'index'. 
//
//		n := index[7:0]
//		dst := a
//		IF (n < 64)
//			dst[63:n] := 0
//		FI
//
// Instruction: 'BZHI'. Intrinsic: '_bzhi_u64'.
// Requires BMI2.
func BzhiU64(a uint64, index uint32) uint64 {
	return uint64(bzhiU64(a, index))
}

func bzhiU64(a uint64, index uint32) uint64


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


// Clevict: Evicts the cache line containing the address 'ptr' from cache level
// 'level' (can be either 0 or 1). 
//
//		CacheLineEvict(ptr, level)
//
// Instruction: 'CLEVICT0, CLEVICT1'. Intrinsic: '_mm_clevict'.
// Requires KNCNI.
func Clevict(ptr uintptr, level int)  {
	clevict(uintptr(ptr), level)
}

func clevict(ptr uintptr, level int) 


// Clflush: Invalidate and flush the cache line that contains 'p' from all
// levels of the cache hierarchy. 
//
//		
//
// Instruction: 'CLFLUSH'. Intrinsic: '_mm_clflush'.
// Requires SSE2.
func Clflush(p uintptr)  {
	clflush(uintptr(p))
}

func clflush(p uintptr) 


// Clflushopt: Invalidate and flush the cache line that contains 'p' from all
// levels of the cache hierarchy. 
//
//		
//
// Instruction: 'CLFLUSHOPT'. Intrinsic: '_mm_clflushopt'.
// Requires CLFLUSHOPT.
func Clflushopt(p uintptr)  {
	clflushopt(uintptr(p))
}

func clflushopt(p uintptr) 


// Cmpestra: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if 'b' did not contain a null
// character and the resulting mask was zero, and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//					0:  // equal any
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					1:  // ranges
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					2:  // equal each
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//					3:  // equal ordered
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 1
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//			0:  // equal any
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound
//						IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//					ENDFOR
//				ENDFOR
//			1:  // ranges
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound, j += 2
//						IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//					ENDFOR
//				ENDFOR
//			2:  // equal each
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					IntRes1[i] := BoolRes[i][i]
//				ENDFOR
//			3:  // equal ordered
//				IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//				FOR i := 0 to UpperBound
//					k := i
//					FOR j := 0 to UpperBound-i
//						IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//						k++
//					ENDFOR
//				ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 == 0) AND (lb > UpperBound)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestra'.
// Requires SSE4.2.
func Cmpestra(a M128i, la int, b M128i, lb int, imm8 int) int {
	return int(cmpestra(a, la, b, lb, imm8))
}

func cmpestra(a M128i, la int, b M128i, lb int, imm8 int) int


// Cmpestrc: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if the resulting mask was
// non-zero, and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//					0:  // equal any
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					1:  // ranges
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					2:  // equal each
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//					3:  // equal ordered
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 1
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//			0:  // equal any
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound
//						IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//					ENDFOR
//				ENDFOR
//			1:  // ranges
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound, j += 2
//						IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//					ENDFOR
//				ENDFOR
//			2:  // equal each
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					IntRes1[i] := BoolRes[i][i]
//				ENDFOR
//			3:  // equal ordered
//				IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//				FOR i := 0 to UpperBound
//					k := i
//					FOR j := 0 to UpperBound-i
//						IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//						k++
//					ENDFOR
//				ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 != 0)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestrc'.
// Requires SSE4.2.
func Cmpestrc(a M128i, la int, b M128i, lb int, imm8 int) int {
	return int(cmpestrc(a, la, b, lb, imm8))
}

func cmpestrc(a M128i, la int, b M128i, lb int, imm8 int) int


// Cmpestri: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and store the generated index in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // most significant bit
//			tmp := UpperBound
//			dst := tmp
//			DO WHILE ((tmp >= 0) AND a[tmp] = 0)
//				tmp := tmp - 1
//				dst := tmp
//			OD
//		ELSE // least significant bit
//			tmp := 0
//			dst := tmp
//			DO WHILE ((tmp <= UpperBound) AND a[tmp] = 0)
//				tmp := tmp + 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestri'.
// Requires SSE4.2.
func Cmpestri(a M128i, la int, b M128i, lb int, imm8 int) int {
	return int(cmpestri(a, la, b, lb, imm8))
}

func cmpestri(a M128i, la int, b M128i, lb int, imm8 int) int


// Cmpestrm: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and store the generated mask in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // byte / word mask
//			FOR i := 0 to UpperBound
//				j := i*size
//				IF IntRes2[i]
//					dst[j+size-1:j] := (imm8[0] ? 0xFF : 0xFFFF)
//				ELSE
//					dst[j+size-1:j] := 0
//				FI
//			ENDFOR
//		ELSE // bit mask
//			dst[UpperBound:0] := IntRes[UpperBound:0]
//			dst[127:UpperBound+1] := 0
//		FI
//
// Instruction: 'PCMPESTRM'. Intrinsic: '_mm_cmpestrm'.
// Requires SSE4.2.
func Cmpestrm(a M128i, la int, b M128i, lb int, imm8 int) M128i {
	return M128i(cmpestrm(a, la, b, lb, imm8))
}

func cmpestrm(a M128i, la int, b M128i, lb int, imm8 int) M128i


// Cmpestro: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns bit 0 of the resulting bit mask.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF i == la
//					aInvalid := 1
//				FI
//				IF j == lb
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//					0:  // equal any
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					1:  // ranges
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						FI
//					2:  // equal each
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 0
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//					3:  // equal ordered
//						IF (!aInvalid && bInvalid)
//							BoolRes[i][j] := 0
//						ELSE IF (aInvalid && !bInvalid)
//							BoolRes[i][j] := 1
//						ELSE If (aInvalid && bInvalid)
//							BoolRes[i][j] := 1
//						FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//			0:  // equal any
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound
//						IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//					ENDFOR
//				ENDFOR
//			1:  // ranges
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					FOR j := 0 to UpperBound, j += 2
//						IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//					ENDFOR
//				ENDFOR
//			2:  // equal each
//				IntRes1 := 0
//				FOR i := 0 to UpperBound
//					IntRes1[i] := BoolRes[i][i]
//				ENDFOR
//			3:  // equal ordered
//				IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//				FOR i := 0 to UpperBound
//					k := i
//					FOR j := 0 to UpperBound-i
//						IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//						k++
//					ENDFOR
//				ENDFOR
//		ESAC
//		
//		// optionally negate results
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF i >= lb // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := IntRes2[0
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestro'.
// Requires SSE4.2.
func Cmpestro(a M128i, la int, b M128i, lb int, imm8 int) int {
	return int(cmpestro(a, la, b, lb, imm8))
}

func cmpestro(a M128i, la int, b M128i, lb int, imm8 int) int


// Cmpestrs: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if any character in 'a' was null,
// and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		dst := (la <= UpperBound)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestrs'.
// Requires SSE4.2.
func Cmpestrs(a M128i, la int, b M128i, lb int, imm8 int) int {
	return int(cmpestrs(a, la, b, lb, imm8))
}

func cmpestrs(a M128i, la int, b M128i, lb int, imm8 int) int


// Cmpestrz: Compare packed strings in 'a' and 'b' with lengths 'la' and 'lb'
// using the control in 'imm8', and returns 1 if any character in 'b' was null,
// and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		dst := (lb <= UpperBound)
//
// Instruction: 'PCMPESTRI'. Intrinsic: '_mm_cmpestrz'.
// Requires SSE4.2.
func Cmpestrz(a M128i, la int, b M128i, lb int, imm8 int) int {
	return int(cmpestrz(a, la, b, lb, imm8))
}

func cmpestrz(a M128i, la int, b M128i, lb int, imm8 int) int


// Cmpistra: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if 'b' did not contain a null character
// and the resulting mask was zero, and 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 == 0) AND bInvalid
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistra'.
// Requires SSE4.2.
func Cmpistra(a M128i, b M128i, imm8 int) int {
	return int(cmpistra(a, b, imm8))
}

func cmpistra(a M128i, b M128i, imm8 int) int


// Cmpistrc: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if the resulting mask was non-zero, and
// 0 otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := (IntRes2 != 0)
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistrc'.
// Requires SSE4.2.
func Cmpistrc(a M128i, b M128i, imm8 int) int {
	return int(cmpistrc(a, b, imm8))
}

func cmpistrc(a M128i, b M128i, imm8 int) int


// Cmpistri: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and store the generated index in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // most significant bit
//			tmp := UpperBound
//			dst := tmp
//			DO WHILE ((tmp >= 0) AND a[tmp] = 0)
//				tmp := tmp - 1
//				dst := tmp
//			OD
//		ELSE // least significant bit
//			tmp := 0
//			dst := tmp
//			DO WHILE ((tmp <= UpperBound) AND a[tmp] = 0)
//				tmp := tmp + 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistri'.
// Requires SSE4.2.
func Cmpistri(a M128i, b M128i, imm8 int) int {
	return int(cmpistri(a, b, imm8))
}

func cmpistri(a M128i, b M128i, imm8 int) int


// Cmpistrm: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and store the generated mask in 'dst'.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		IF imm8[6] // byte / word mask
//			FOR i := 0 to UpperBound
//				j := i*size
//				IF IntRes2[i]
//					dst[j+size-1:j] := (imm8[0] ? 0xFF : 0xFFFF)
//				ELSE
//					dst[j+size-1:j] := 0
//				FI
//			ENDFOR
//		ELSE // bit mask
//			dst[UpperBound:0] := IntRes[UpperBound:0]
//			dst[127:UpperBound+1] := 0
//		FI
//
// Instruction: 'PCMPISTRM'. Intrinsic: '_mm_cmpistrm'.
// Requires SSE4.2.
func Cmpistrm(a M128i, b M128i, imm8 int) M128i {
	return M128i(cmpistrm(a, b, imm8))
}

func cmpistrm(a M128i, b M128i, imm8 int) M128i


// Cmpistro: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns bit 0 of the resulting bit mask.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		// compare all characters
//		aInvalid := 0
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			FOR j := 0 to UpperBound
//				n := j*size
//				BoolRes[i][j] := (a[m+size-1:m] == b[n+size-1:n])
//				
//				// invalidate characters after EOS
//				IF a[m+size-1:m] == 0
//					aInvalid := 1
//				FI
//				IF b[n+size-1:n] == 0
//					bInvalid := 1
//				FI
//				
//				// override comparisons for invalid characters
//				CASE (imm8[3:2]) OF
//				0:  // equal any
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				1:  // ranges
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					FI
//				2:  // equal each
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 0
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				3:  // equal ordered
//					IF (!aInvalid && bInvalid)
//						BoolRes[i][j] := 0
//					ELSE IF (aInvalid && !bInvalid)
//						BoolRes[i][j] := 1
//					ELSE If (aInvalid && bInvalid)
//						BoolRes[i][j] := 1
//					FI
//				ESAC
//			ENDFOR
//		ENDFOR
//		
//		// aggregate results
//		CASE (imm8[3:2]) OF
//		0:  // equal any
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound
//					IntRes1[i] := IntRes1[i] OR BoolRes[i][j]
//				ENDFOR
//			ENDFOR
//		1:  // ranges
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				FOR j := 0 to UpperBound, j += 2
//					IntRes1[i] := IntRes1[i] OR (BoolRes[i][j] AND BoolRes[i][j+1])
//				ENDFOR
//			ENDFOR
//		2:  // equal each
//			IntRes1 := 0
//			FOR i := 0 to UpperBound
//				IntRes1[i] := BoolRes[i][i]
//			ENDFOR
//		3:  // equal ordered
//			IntRes1 := (imm8[0] ? 0xFF : 0xFFFF)
//			FOR i := 0 to UpperBound
//				k := i
//				FOR j := 0 to UpperBound-i
//					IntRes1[i] := IntRes1[i] AND BoolRes[k][j]
//					k++
//				ENDFOR
//			ENDFOR
//		ESAC
//		
//		// optionally negate results
//		bInvalid := 0
//		FOR i := 0 to UpperBound
//			IF imm8[4]
//				IF imm8[5] // only negate valid
//					IF b[n+size-1:n] == 0
//						bInvalid := 1
//					FI
//					IF bInvalid // invalid, don't negate
//						IntRes2[i] := IntRes1[i]
//					ELSE // valid, negate
//						IntRes2[i] := -1 XOR IntRes1[i]
//					FI
//				ELSE // negate all
//					IntRes2[i] := -1 XOR IntRes1[i]
//				FI
//			ELSE // don't negate
//				IntRes2[i] := IntRes1[i]
//			FI
//		ENDFOR
//		
//		// output
//		dst := IntRes2[0]
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistro'.
// Requires SSE4.2.
func Cmpistro(a M128i, b M128i, imm8 int) int {
	return int(cmpistro(a, b, imm8))
}

func cmpistro(a M128i, b M128i, imm8 int) int


// Cmpistrs: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if any character in 'a' was null, and 0
// otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		aInvalid := 0
//		FOR i := 0 to UpperBound
//			m := i*size
//			IF b[m+size-1:m] == 0
//				aInvalid := 1
//			FI
//		ENDFOR
//		
//		dst := aInvalid
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistrs'.
// Requires SSE4.2.
func Cmpistrs(a M128i, b M128i, imm8 int) int {
	return int(cmpistrs(a, b, imm8))
}

func cmpistrs(a M128i, b M128i, imm8 int) int


// Cmpistrz: Compare packed strings with implicit lengths in 'a' and 'b' using
// the control in 'imm8', and returns 1 if any character in 'b' was null, and 0
// otherwise.
// 	'imm' can be a combination of:
//     _SIDD_UBYTE_OPS                // unsigned 8-bit characters
//     _SIDD_UWORD_OPS                // unsigned 16-bit characters
//     _SIDD_SBYTE_OPS                // signed 8-bit characters
//     _SIDD_SWORD_OPS                // signed 16-bit characters
//     _SIDD_CMP_EQUAL_ANY            // compare equal any
//     _SIDD_CMP_RANGES               // compare ranges
//     _SIDD_CMP_EQUAL_EACH           // compare equal each
//     _SIDD_CMP_EQUAL_ORDERED        // compare equal ordered
//     _SIDD_NEGATIVE_POLARITY        // negate results
//     _SIDD_MASKED_NEGATIVE_POLARITY // negate results only before end of string
//     _SIDD_LEAST_SIGNIFICANT        // index only: return last significant bit
//     _SIDD_MOST_SIGNIFICANT         // index only: return most significant bit
//     _SIDD_BIT_MASK                 // mask only: return bit mask
//     _SIDD_UNIT_MASK                // mask only: return byte/word mask 
//
//		size := (imm8[0] ? 16 : 8) // 8 or 16-bit characters
//		UpperBound := (128 / size) - 1
//		
//		bInvalid := 0
//		FOR j := 0 to UpperBound
//			n := j*size
//			IF b[n+size-1:n] == 0
//				bInvalid := 1
//			FI
//		ENDFOR
//		
//		dst := bInvalid
//
// Instruction: 'PCMPISTRI'. Intrinsic: '_mm_cmpistrz'.
// Requires SSE4.2.
func Cmpistrz(a M128i, b M128i, imm8 int) int {
	return int(cmpistrz(a, b, imm8))
}

func cmpistrz(a M128i, b M128i, imm8 int) int


// Countbits32: Counts the number of set bits in 32-bit unsigned integer 'r1',
// returning the results in 'dst'. 
//
//		dst[31:0] := PopCount(r1[31:0])
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_countbits_32'.
// Requires KNCNI.
func Countbits32(r1 uint32) uint32 {
	return uint32(countbits32(r1))
}

func countbits32(r1 uint32) uint32


// Countbits64: Counts the number of set bits in double-precision (32-bit)
// unsigned integer 'r1', returning the results in 'dst'. 
//
//		dst[63:0] := PopCount(r1[63:0])
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_countbits_64'.
// Requires KNCNI.
func Countbits64(r1 uint64) uint64 {
	return uint64(countbits64(r1))
}

func countbits64(r1 uint64) uint64


// Crc32U16: Starting with the initial value in 'crc', accumulates a CRC32
// value for unsigned 16-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[15:0] := v[0:15] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[47:0] := tmp1[15:0] << 32
//		tmp4[47:0] := tmp2[31:0] << 16
//		tmp5[47:0] := tmp3[47:0] XOR tmp4[47:0]
//		tmp6[31:0] := tmp5[47:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u16'.
// Requires SSE4.2.
func Crc32U16(crc uint32, v uint16) uint32 {
	return uint32(crc32U16(crc, v))
}

func crc32U16(crc uint32, v uint16) uint32


// Crc32U32: Starting with the initial value in 'crc', accumulates a CRC32
// value for unsigned 32-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[31:0] := v[0:31] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[63:0] := tmp1[31:0] << 32
//		tmp4[63:0] := tmp2[31:0] << 32
//		tmp5[63:0] := tmp3[63:0] XOR tmp4[63:0]
//		tmp6[31:0] := tmp5[63:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u32'.
// Requires SSE4.2.
func Crc32U32(crc uint32, v uint32) uint32 {
	return uint32(crc32U32(crc, v))
}

func crc32U32(crc uint32, v uint32) uint32


// Crc32U64: Starting with the initial value in 'crc', accumulates a CRC32
// value for unsigned 64-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[63:0] := v[0:63] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[95:0] := tmp1[31:0] << 32
//		tmp4[95:0] := tmp2[63:0] << 64
//		tmp5[95:0] := tmp3[95:0] XOR tmp4[95:0]
//		tmp6[31:0] := tmp5[95:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u64'.
// Requires SSE4.2.
func Crc32U64(crc uint64, v uint64) uint64 {
	return uint64(crc32U64(crc, v))
}

func crc32U64(crc uint64, v uint64) uint64


// Crc32U8: Starting with the initial value in 'crc', accumulates a CRC32 value
// for unsigned 8-bit integer 'v', and stores the result in 'dst'. 
//
//		tmp1[7:0] := v[0:7] // bit reflection
//		tmp2[31:0] := crc[0:31] // bit reflection
//		tmp3[39:0] := tmp1[7:0] << 32 
//		tmp4[39:0] := tmp2[31:0] << 8
//		tmp5[39:0] := tmp3[39:0] XOR tmp4[39:0]
//		tmp6[31:0] := tmp5[39:0] MOD2 0x11EDC6F41
//		dst[31:0] := tmp6[0:31] // bit reflection
//
// Instruction: 'CRC32'. Intrinsic: '_mm_crc32_u8'.
// Requires SSE4.2.
func Crc32U8(crc uint32, v uint8) uint32 {
	return uint32(crc32U8(crc, v))
}

func crc32U8(crc uint32, v uint8) uint32


// Delay32: Stalls a thread without blocking other threads for 32-bit unsigned
// integer 'r1' clock cycles. 
//
//		BlockThread(r1)
//
// Instruction: 'DELAY'. Intrinsic: '_mm_delay_32'.
// Requires KNCNI.
func Delay32(r1 uint32)  {
	delay32(r1)
}

func delay32(r1 uint32) 


// Delay64: Stalls a thread without blocking other threads for 64-bit unsigned
// integer 'r1' clock cycles. 
//
//		BlockThread(r1)
//
// Instruction: 'DELAY'. Intrinsic: '_mm_delay_64'.
// Requires KNCNI.
func Delay64(r1 uint64)  {
	delay64(r1)
}

func delay64(r1 uint64) 


// Free: Free aligned memory that was allocated with '_mm_malloc'. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_free'.
// Requires SSE.
func Free(mem_addr uintptr)  {
	free(uintptr(mem_addr))
}

func free(mem_addr uintptr) 


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


// MMGETEXCEPTIONMASK: Macro: Get the exception mask bits from the MXCSR
// control and status register. The exception mask may contain any of the
// following flags: _MM_MASK_INVALID, _MM_MASK_DIV_ZERO, _MM_MASK_DENORM,
// _MM_MASK_OVERFLOW, _MM_MASK_UNDERFLOW, _MM_MASK_INEXACT 
//
//		dst[31:0] := MXCSR & _MM_MASK_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_EXCEPTION_MASK'.
// Requires SSE.
func MMGETEXCEPTIONMASK() uint32 {
	return uint32(mMGETEXCEPTIONMASK())
}

func mMGETEXCEPTIONMASK() uint32


// MMGETEXCEPTIONSTATE: Macro: Get the exception state bits from the MXCSR
// control and status register. The exception state may contain any of the
// following flags: _MM_EXCEPT_INVALID, _MM_EXCEPT_DIV_ZERO, _MM_EXCEPT_DENORM,
// _MM_EXCEPT_OVERFLOW, _MM_EXCEPT_UNDERFLOW, _MM_EXCEPT_INEXACT 
//
//		dst[31:0] := MXCSR & _MM_EXCEPT_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_EXCEPTION_STATE'.
// Requires SSE.
func MMGETEXCEPTIONSTATE() uint32 {
	return uint32(mMGETEXCEPTIONSTATE())
}

func mMGETEXCEPTIONSTATE() uint32


// MMGETFLUSHZEROMODE: Macro: Get the flush zero bits from the MXCSR control
// and status register. The flush zero may contain any of the following flags:
// _MM_FLUSH_ZERO_ON or _MM_FLUSH_ZERO_OFF 
//
//		dst[31:0] := MXCSR & _MM_FLUSH_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_FLUSH_ZERO_MODE'.
// Requires SSE.
func MMGETFLUSHZEROMODE() uint32 {
	return uint32(mMGETFLUSHZEROMODE())
}

func mMGETFLUSHZEROMODE() uint32


// MMGETROUNDINGMODE: Macro: Get the rounding mode bits from the MXCSR control
// and status register. The rounding mode may contain any of the following
// flags: _MM_ROUND_NEAREST, _MM_ROUND_DOWN, _MM_ROUND_UP,
// _MM_ROUND_TOWARD_ZERO 
//
//		dst[31:0] := MXCSR & _MM_ROUND_MASK
//
// Instruction: ''. Intrinsic: '_MM_GET_ROUNDING_MODE'.
// Requires SSE.
func MMGETROUNDINGMODE() uint32 {
	return uint32(mMGETROUNDINGMODE())
}

func mMGETROUNDINGMODE() uint32


// Getcsr: Get the unsigned 32-bit value of the MXCSR control and status
// register. 
//
//		dst[31:0] := MXCSR
//
// Instruction: 'STMXCSR'. Intrinsic: '_mm_getcsr'.
// Requires SSE.
func Getcsr() uint32 {
	return uint32(getcsr())
}

func getcsr() uint32


// Invpcid: Invalidate mappings in the Translation Lookaside Buffers (TLBs) and
// paging-structure caches for the processor context identifier (PCID)
// specified by 'descriptor' based on the invalidation type specified in
// 'type'. 
// 	The PCID 'descriptor' is specified as a 16-byte memory operand (with no
// alignment restrictions) where bits [11:0] specify the PCID, and bits
// [127:64] specify the linear address; bits [63:12] are reserved.
// 	The types supported are:
// 		0) Individual-address invalidation: If 'type' is 0, the logical processor
// invalidates mappings for a single linear address and tagged with the PCID
// specified in 'descriptor', except global translations. The instruction may
// also invalidate global translations, mappings for other linear addresses, or
// mappings tagged with other PCIDs.
// 		1) Single-context invalidation: If 'type' is 1, the logical processor
// invalidates all mappings tagged with the PCID specified in 'descriptor'
// except global translations. In some cases, it may invalidate mappings for
// other PCIDs as well.
// 		2) All-context invalidation: If 'type' is 2, the logical processor
// invalidates all mappings tagged with any PCID.
// 		3) All-context invalidation, retaining global translations: If 'type' is
// 3, the logical processor invalidates all mappings tagged with any PCID
// except global translations, ignoring 'descriptor'. The instruction may also
// invalidate global translations as well. 
//
//		CASE type OF
//		0: // individual-address invalidation retaining global translations
//			OP_PCID := descriptor[11:0]
//			ADDR := descriptor[127:64]
//			BREAK
//		1: // single PCID invalidation retaining globals
//			OP_PCID := descriptor[11:0]
//			// invalidate all mappings tagged with OP_PCID except global translations
//			BREAK
//		2: // all PCID invalidation
//			// invalidate all mappings tagged with any PCID
//			BREAK
//		3: // all PCID invalidation retaining global translations
//			// invalidate all mappings tagged with any PCID except global translations
//			BREAK
//		ESAC
//
// Instruction: 'INVPCID'. Intrinsic: '_invpcid'.
// Requires INVPCID.
func Invpcid(typ uint32, descriptor uintptr)  {
	invpcid(typ, uintptr(descriptor))
}

func invpcid(typ uint32, descriptor uintptr) 


// Lfence: Perform a serializing operation on all load-from-memory instructions
// that were issued prior to this instruction. Guarantees that every load
// instruction that precedes, in program order, is globally visible before any
// load instruction which follows the fence in program order. 
//
//		
//
// Instruction: 'LFENCE'. Intrinsic: '_mm_lfence'.
// Requires SSE2.
func Lfence()  {
	lfence()
}

func lfence() 


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


// LzcntU32: Count the number of leading zero bits in unsigned 32-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 31
//		dst := 0
//		DO WHILE (tmp >= 0 AND a[tmp] = 0)
//			tmp := tmp - 1
//			dst := dst + 1
//		OD
//
// Instruction: 'LZCNT'. Intrinsic: '_lzcnt_u32'.
// Requires LZCNT.
func LzcntU32(a uint32) uint32 {
	return uint32(lzcntU32(a))
}

func lzcntU32(a uint32) uint32


// LzcntU64: Count the number of leading zero bits in unsigned 64-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 63
//		dst := 0
//		DO WHILE (tmp >= 0 AND a[tmp] = 0)
//			tmp := tmp - 1
//			dst := dst + 1
//		OD
//
// Instruction: 'LZCNT'. Intrinsic: '_lzcnt_u64'.
// Requires LZCNT.
func LzcntU64(a uint64) uint64 {
	return uint64(lzcntU64(a))
}

func lzcntU64(a uint64) uint64


// Malloc: Allocate 'size' bytes of memory, aligned to the alignment specified
// in 'align', and return a pointer to the allocated memory. '_mm_free' should
// be used to free memory that is allocated with '_mm_malloc'. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm_malloc'.
// Requires SSE.
func Malloc(size int, align int)  {
	malloc(size, align)
}

func malloc(size int, align int) 


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


// Mfence: Perform a serializing operation on all load-from-memory and
// store-to-memory instructions that were issued prior to this instruction.
// Guarantees that every memory access that precedes, in program order, the
// memory fence instruction is globally visible before any memory instruction
// which follows the fence in program order. 
//
//		
//
// Instruction: 'MFENCE'. Intrinsic: '_mm_mfence'.
// Requires SSE2.
func Mfence()  {
	mfence()
}

func mfence() 


// Monitor: Arm address monitoring hardware using the address specified in 'p'.
// A store to an address within the specified address range triggers the
// monitoring hardware. Specify optional extensions in 'extensions', and
// optional hints in 'hints'. 
//
//		
//
// Instruction: 'MONITOR'. Intrinsic: '_mm_monitor'.
// Requires MONITOR.
func Monitor(p uintptr, extensions uint, hints uint)  {
	monitor(uintptr(p), extensions, hints)
}

func monitor(p uintptr, extensions uint, hints uint) 


// Mwait: Hint to the processor that it can enter an
// implementation-dependent-optimized state while waiting for an event or store
// operation to the address range specified by MONITOR. 
//
//		
//
// Instruction: 'MWAIT'. Intrinsic: '_mm_mwait'.
// Requires MONITOR.
func Mwait(extensions uint, hints uint)  {
	mwait(extensions, hints)
}

func mwait(extensions uint, hints uint) 


// Pause: Provide a hint to the processor that the code sequence is a spin-wait
// loop. This can help improve the performance and power consumption of
// spin-wait loops. 
//
//		
//
// Instruction: 'PAUSE'. Intrinsic: '_mm_pause'.
// Requires SSE2.
func Pause()  {
	pause()
}

func pause() 


// PextU32: Extract bits from unsigned 32-bit integer 'a' at the corresponding
// bit locations specified by 'mask' to contiguous low bits in 'dst'; the
// remaining upper bits in 'dst' are set to zero. 
//
//		tmp := a
//		dst := 0
//		m := 0
//		k := 0
//		DO WHILE m < 32
//			IF mask[m] = 1
//				dst[k] := tmp[m]
//				k := k + 1
//			FI
//			m := m + 1
//		OD
//
// Instruction: 'PEXT'. Intrinsic: '_pext_u32'.
// Requires BMI2.
func PextU32(a uint32, mask uint32) uint32 {
	return uint32(pextU32(a, mask))
}

func pextU32(a uint32, mask uint32) uint32


// PextU64: Extract bits from unsigned 64-bit integer 'a' at the corresponding
// bit locations specified by 'mask' to contiguous low bits in 'dst'; the
// remaining upper bits in 'dst' are set to zero. 
//
//		tmp := a
//		dst := 0
//		m := 0
//		k := 0
//		DO WHILE m < 64
//			IF mask[m] = 1
//				dst[k] := tmp[m]
//				k := k + 1
//			FI
//			m := m + 1
//		OD
//
// Instruction: 'PEXT'. Intrinsic: '_pext_u64'.
// Requires BMI2.
func PextU64(a uint64, mask uint64) uint64 {
	return uint64(pextU64(a, mask))
}

func pextU64(a uint64, mask uint64) uint64


// PopcntU32: Count the number of bits set to 1 in unsigned 32-bit integer 'a',
// and return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 31
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_popcnt_u32'.
// Requires POPCNT.
func PopcntU32(a uint32) int {
	return int(popcntU32(a))
}

func popcntU32(a uint32) int


// PopcntU64: Count the number of bits set to 1 in unsigned 64-bit integer 'a',
// and return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 63
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_popcnt_u64'.
// Requires POPCNT.
func PopcntU64(a uint64) int64 {
	return int64(popcntU64(a))
}

func popcntU64(a uint64) int64


// Popcnt32: Count the number of bits set to 1 in 32-bit integer 'a', and
// return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 31
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_popcnt32'.
// Requires POPCNT.
func Popcnt32(a int) int {
	return int(popcnt32(a))
}

func popcnt32(a int) int


// Popcnt64: Count the number of bits set to 1 in 64-bit integer 'a', and
// return that count in 'dst'. 
//
//		dst := 0
//		FOR i := 0 to 63
//			IF a[i]
//				dst := dst + 1
//			FI
//		ENDFOR
//
// Instruction: 'POPCNT'. Intrinsic: '_popcnt64'.
// Requires POPCNT.
func Popcnt64(a int64) int {
	return int(popcnt64(a))
}

func popcnt64(a int64) int


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'PREFETCHNTA, PREFETCHT0, PREFETCHT1, PREFETCHT2'. Intrinsic: '_mm_prefetch'.
// Requires SSE.
func Prefetch(p byte, i int)  {
	prefetch(p, i)
}

func prefetch(p byte, i int) 


// Prefetch1: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'PREFETCHWT1'. Intrinsic: '_mm_prefetch'.
// Requires PREFETCHWT1.
func Prefetch1(p byte, i int)  {
	prefetch1(p, i)
}

func prefetch1(p byte, i int) 


// Prefetch2: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'VPREFETCH0, VPREFETCH1, VPREFETCH2, VPREFETCHNTA, VPREFETCHE0, VPREFETCHE1, VPREFETCHE2, VPREFETCHENTA'. Intrinsic: '_mm_prefetch'.
// Requires KNCNI.
func Prefetch2(p byte, i int)  {
	prefetch2(p, i)
}

func prefetch2(p byte, i int) 


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


// Rdrand16Step: Read a hardware generated 16-bit random value and store the
// result in 'val'. Return 1 if a random value was generated, and 0 otherwise. 
//
//		IF HW_RND_GEN.ready = 1
//			val[15:0] := HW_RND_GEN.data;
//			RETURN 1;
//		ELSE
//			val[15:0] := 0;
//			RETURN 0;
//		FI
//
// Instruction: 'RDRAND'. Intrinsic: '_rdrand16_step'.
// Requires RDRAND.
func Rdrand16Step(val uint16) int {
	return int(rdrand16Step(val))
}

func rdrand16Step(val uint16) int


// Rdrand32Step: Read a hardware generated 32-bit random value and store the
// result in 'val'. Return 1 if a random value was generated, and 0 otherwise. 
//
//		IF HW_RND_GEN.ready = 1
//			val[31:0] := HW_RND_GEN.data;
//			RETURN 1;
//		ELSE
//			val[31:0] := 0;
//			RETURN 0;
//		FI
//
// Instruction: 'RDRAND'. Intrinsic: '_rdrand32_step'.
// Requires RDRAND.
func Rdrand32Step(val uint32) int {
	return int(rdrand32Step(val))
}

func rdrand32Step(val uint32) int


// Rdrand64Step: Read a hardware generated 64-bit random value and store the
// result in 'val'. Return 1 if a random value was generated, and 0 otherwise. 
//
//		IF HW_RND_GEN.ready = 1
//			val[63:0] := HW_RND_GEN.data;
//			RETURN 1;
//		ELSE
//			val[63:0] := 0;
//			RETURN 0;
//		FI
//
// Instruction: 'RDRAND'. Intrinsic: '_rdrand64_step'.
// Requires RDRAND.
func Rdrand64Step(val uint64) int {
	return int(rdrand64Step(val))
}

func rdrand64Step(val uint64) int


// Rdseed16Step: Read a 16-bit NIST SP800-90B and SP800-90C compliant random
// value and store in 'val'. Return 1 if a random value was generated, and 0
// otherwise. 
//
//		IF HW_NRND_GEN.ready = 1 THEN
//			val[15:0] := HW_NRND_GEN.data
//			RETURN 1
//		ELSE
//			val[15:0] := 0
//			RETURN 0
//		FI
//
// Instruction: 'RDSEED'. Intrinsic: '_rdseed16_step'.
// Requires RDSEED.
func Rdseed16Step(val uint16) int {
	return int(rdseed16Step(val))
}

func rdseed16Step(val uint16) int


// Rdseed32Step: Read a 32-bit NIST SP800-90B and SP800-90C compliant random
// value and store in 'val'. Return 1 if a random value was generated, and 0
// otherwise. 
//
//		IF HW_NRND_GEN.ready = 1 THEN
//			val[31:0] := HW_NRND_GEN.data
//			RETURN 1
//		ELSE
//			val[31:0] := 0
//			RETURN 0
//		FI
//
// Instruction: 'RDSEED'. Intrinsic: '_rdseed32_step'.
// Requires RDSEED.
func Rdseed32Step(val uint32) int {
	return int(rdseed32Step(val))
}

func rdseed32Step(val uint32) int


// Rdseed64Step: Read a 64-bit NIST SP800-90B and SP800-90C compliant random
// value and store in 'val'. Return 1 if a random value was generated, and 0
// otherwise. 
//
//		IF HW_NRND_GEN.ready = 1 THEN
//			val[63:0] := HW_NRND_GEN.data
//			RETURN 1
//		ELSE
//			val[63:0] := 0
//			RETURN 0
//		FI
//
// Instruction: 'RDSEED'. Intrinsic: '_rdseed64_step'.
// Requires RDSEED.
func Rdseed64Step(val uint64) int {
	return int(rdseed64Step(val))
}

func rdseed64Step(val uint64) int


// Rdtsc: Copy the current 64-bit value of the processor's time-stamp counter
// into 'dst'. 
//
//		dst[63:0] := TimeStampCounter
//
// Instruction: 'RDTSC'. Intrinsic: '_rdtsc'.
// Requires TSC.
func Rdtsc() int64 {
	return int64(rdtsc())
}

func rdtsc() int64


// Rdtscp: Copy the current 64-bit value of the processor's time-stamp counter
// into 'dst', and store the IA32_TSC_AUX MSR (signature value) into memory at
// 'mem_addr'. 
//
//		dst[63:0] := TimeStampCounter
//		MEM[mem_addr+31:mem_addr] := IA32_TSC_AUX[31:0]
//
// Instruction: 'RDTSCP'. Intrinsic: '__rdtscp'.
// Requires RDTSCP.
func Rdtscp(mem_addr uint32) uint64 {
	return uint64(rdtscp(mem_addr))
}

func rdtscp(mem_addr uint32) uint64


// ReadfsbaseU32: Read the FS segment base register and store the 32-bit result
// in 'dst'. 
//
//		dst[31:0] := FS_Segment_Base_Register;
//		dst[63:32] := 0
//
// Instruction: 'RDFSBASE'. Intrinsic: '_readfsbase_u32'.
// Requires FSGSBASE.
func ReadfsbaseU32() uint32 {
	return uint32(readfsbaseU32())
}

func readfsbaseU32() uint32


// ReadfsbaseU64: Read the FS segment base register and store the 64-bit result
// in 'dst'. 
//
//		dst[63:0] := FS_Segment_Base_Register;
//
// Instruction: 'RDFSBASE'. Intrinsic: '_readfsbase_u64'.
// Requires FSGSBASE.
func ReadfsbaseU64() uint64 {
	return uint64(readfsbaseU64())
}

func readfsbaseU64() uint64


// ReadgsbaseU32: Read the GS segment base register and store the 32-bit result
// in 'dst'. 
//
//		dst[31:0] := GS_Segment_Base_Register;
//		dst[63:32] := 0
//
// Instruction: 'RDGSBASE'. Intrinsic: '_readgsbase_u32'.
// Requires FSGSBASE.
func ReadgsbaseU32() uint32 {
	return uint32(readgsbaseU32())
}

func readgsbaseU32() uint32


// ReadgsbaseU64: Read the GS segment base register and store the 64-bit result
// in 'dst'. 
//
//		dst[63:0] := GS_Segment_Base_Register;
//
// Instruction: 'RDGSBASE'. Intrinsic: '_readgsbase_u64'.
// Requires FSGSBASE.
func ReadgsbaseU64() uint64 {
	return uint64(readgsbaseU64())
}

func readgsbaseU64() uint64


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


// MMSETEXCEPTIONMASK: Macro: Set the exception mask bits of the MXCSR control
// and status register to the value in unsigned 32-bit integer 'a'. The
// exception mask may contain any of the following flags: _MM_MASK_INVALID,
// _MM_MASK_DIV_ZERO, _MM_MASK_DENORM, _MM_MASK_OVERFLOW, _MM_MASK_UNDERFLOW,
// _MM_MASK_INEXACT 
//
//		MXCSR := a[31:0] AND ~_MM_MASK_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_EXCEPTION_MASK'.
// Requires SSE.
func MMSETEXCEPTIONMASK(a uint32)  {
	mMSETEXCEPTIONMASK(a)
}

func mMSETEXCEPTIONMASK(a uint32) 


// MMSETEXCEPTIONSTATE: Macro: Set the exception state bits of the MXCSR
// control and status register to the value in unsigned 32-bit integer 'a'. The
// exception state may contain any of the following flags: _MM_EXCEPT_INVALID,
// _MM_EXCEPT_DIV_ZERO, _MM_EXCEPT_DENORM, _MM_EXCEPT_OVERFLOW,
// _MM_EXCEPT_UNDERFLOW, _MM_EXCEPT_INEXACT 
//
//		MXCSR := a[31:0] AND ~_MM_EXCEPT_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_EXCEPTION_STATE'.
// Requires SSE.
func MMSETEXCEPTIONSTATE(a uint32)  {
	mMSETEXCEPTIONSTATE(a)
}

func mMSETEXCEPTIONSTATE(a uint32) 


// MMSETFLUSHZEROMODE: Macro: Set the flush zero bits of the MXCSR control and
// status register to the value in unsigned 32-bit integer 'a'. The flush zero
// may contain any of the following flags: _MM_FLUSH_ZERO_ON or
// _MM_FLUSH_ZERO_OFF 
//
//		MXCSR := a[31:0] AND ~_MM_FLUSH_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_FLUSH_ZERO_MODE'.
// Requires SSE.
func MMSETFLUSHZEROMODE(a uint32)  {
	mMSETFLUSHZEROMODE(a)
}

func mMSETFLUSHZEROMODE(a uint32) 


// MMSETROUNDINGMODE: Macro: Set the rounding mode bits of the MXCSR control
// and status register to the value in unsigned 32-bit integer 'a'. The
// rounding mode may contain any of the following flags: _MM_ROUND_NEAREST,
// _MM_ROUND_DOWN, _MM_ROUND_UP, _MM_ROUND_TOWARD_ZERO 
//
//		MXCSR := a[31:0] AND ~_MM_ROUND_MASK
//
// Instruction: ''. Intrinsic: '_MM_SET_ROUNDING_MODE'.
// Requires SSE.
func MMSETROUNDINGMODE(a uint32)  {
	mMSETROUNDINGMODE(a)
}

func mMSETROUNDINGMODE(a uint32) 


// Setcsr: Set the MXCSR control and status register with the value in unsigned
// 32-bit integer 'a'. 
//
//		MXCSR := a[31:0]
//
// Instruction: 'LDMXCSR'. Intrinsic: '_mm_setcsr'.
// Requires SSE.
func Setcsr(a uint32)  {
	setcsr(a)
}

func setcsr(a uint32) 


// Sfence: Perform a serializing operation on all store-to-memory instructions
// that were issued prior to this instruction. Guarantees that every store
// instruction that precedes, in program order, is globally visible before any
// store instruction which follows the fence in program order. 
//
//		
//
// Instruction: 'SFENCE'. Intrinsic: '_mm_sfence'.
// Requires SSE.
func Sfence()  {
	sfence()
}

func sfence() 


// Spflt32: Set performance monitoring filtering mask to 32-bit unsigned
// integer 'r1'. 
//
//		SetPerfMonMask(r1[31:0])
//
// Instruction: 'SPFLT'. Intrinsic: '_mm_spflt_32'.
// Requires KNCNI.
func Spflt32(r1 uint32)  {
	spflt32(r1)
}

func spflt32(r1 uint32) 


// Spflt64: Set performance monitoring filtering mask to 64-bit unsigned
// integer 'r1'. 
//
//		SetPerfMonMask(r1[63:0])
//
// Instruction: 'SPFLT'. Intrinsic: '_mm_spflt_64'.
// Requires KNCNI.
func Spflt64(r1 uint64)  {
	spflt64(r1)
}

func spflt64(r1 uint64) 


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


// TestAllOnes: Compute the complement of 'a' and 0xFFFFFFFF, and return 1 if
// the result is zero, otherwise return 0. 
//
//		IF (a[127:0] AND NOT 0xFFFFFFFF == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		RETURN CF
//
// Instruction: '...'. Intrinsic: '_mm_test_all_ones'.
// Requires SSE4.1.
func TestAllOnes(a M128i) int {
	return int(testAllOnes(a))
}

func testAllOnes(a M128i) int


// TestAllZeros: Compute the bitwise AND of 128 bits (representing integer
// data) in 'a' and 'mask', and return 1 if the result is zero, otherwise
// return 0. 
//
//		IF (a[127:0] AND mask[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		RETURN ZF
//
// Instruction: 'PTEST'. Intrinsic: '_mm_test_all_zeros'.
// Requires SSE4.1.
func TestAllZeros(a M128i, mask M128i) int {
	return int(testAllZeros(a, mask))
}

func testAllZeros(a M128i, mask M128i) int


// TestMixOnesZeros: Compute the bitwise AND of 128 bits (representing integer
// data) in 'a' and 'mask', and set 'ZF' to 1 if the result is zero, otherwise
// set 'ZF' to 0. Compute the bitwise AND NOT of 'a' and 'mask', and set 'CF'
// to 1 if the result is zero, otherwise set 'CF' to 0. Return 1 if both the
// 'ZF' and 'CF' values are zero, otherwise return 0. 
//
//		IF (a[127:0] AND mask[127:0] == 0)
//			ZF := 1
//		ELSE
//			ZF := 0
//		FI
//		IF (a[127:0] AND NOT mask[127:0] == 0)
//			CF := 1
//		ELSE
//			CF := 0
//		FI
//		IF (ZF == 0 && CF == 0)
//			RETURN 1
//		ELSE
//			RETURN 0
//		FI
//
// Instruction: 'PTEST'. Intrinsic: '_mm_test_mix_ones_zeros'.
// Requires SSE4.1.
func TestMixOnesZeros(a M128i, mask M128i) int {
	return int(testMixOnesZeros(a, mask))
}

func testMixOnesZeros(a M128i, mask M128i) int


// MMTRANSPOSE4PS: Macro: Transpose the 4x4 matrix formed by the 4 rows of
// single-precision (32-bit) floating-point elements in 'row0', 'row1', 'row2',
// and 'row3', and store the transposed matrix in these vectors ('row0' now
// contains column 0, etc.). 
//
//		__m128 tmp3, tmp2, tmp1, tmp0;
//		tmp0 = _mm_unpacklo_ps(row0, row1);
//		tmp2 = _mm_unpacklo_ps(row2, row3);
//		tmp1 = _mm_unpackhi_ps(row0, row1);
//		tmp3 = _mm_unpackhi_ps(row2, row3);
//		row0 = _mm_movelh_ps(tmp0, tmp2);
//		row1 = _mm_movehl_ps(tmp2, tmp0);
//		row2 = _mm_movelh_ps(tmp1, tmp3);
//		row3 = _mm_movehl_ps(tmp3, tmp1);
//
// Instruction: '...'. Intrinsic: '_MM_TRANSPOSE4_PS'.
// Requires SSE.
func MMTRANSPOSE4PS(row0 M128, row1 M128, row2 M128, row3 M128)  {
	mMTRANSPOSE4PS([4]float32(row0), [4]float32(row1), [4]float32(row2), [4]float32(row3))
}

func mMTRANSPOSE4PS(row0 [4]float32, row1 [4]float32, row2 [4]float32, row3 [4]float32) 


// Tzcnt32: Count the number of trailing zero bits in unsigned 32-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 32) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_mm_tzcnt_32'.
// Requires BMI1.
func Tzcnt32(a uint32) int {
	return int(tzcnt32(a))
}

func tzcnt32(a uint32) int


// Tzcnt64: Count the number of trailing zero bits in unsigned 64-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 64) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_mm_tzcnt_64'.
// Requires BMI1.
func Tzcnt64(a uint64) int64 {
	return int64(tzcnt64(a))
}

func tzcnt64(a uint64) int64


// TzcntU32: Count the number of trailing zero bits in unsigned 32-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 32) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_tzcnt_u32'.
// Requires BMI1.
func TzcntU32(a uint32) uint32 {
	return uint32(tzcntU32(a))
}

func tzcntU32(a uint32) uint32


// TzcntU64: Count the number of trailing zero bits in unsigned 64-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 0
//		dst := 0
//		DO WHILE ((tmp < 64) AND a[tmp] = 0)
//			tmp := tmp + 1
//			dst := dst + 1
//		OD
//
// Instruction: 'TZCNT'. Intrinsic: '_tzcnt_u64'.
// Requires BMI1.
func TzcntU64(a uint64) uint64 {
	return uint64(tzcntU64(a))
}

func tzcntU64(a uint64) uint64


// Tzcnti32: Counts the number of trailing bits in unsigned 32-bit integer 'x'
// starting at bit 'a' storing the result in 'dst'. 
//
//		count := 0
//		FOR j := a to 31
//			IF NOT(x[j]  1)
//				count := count + 1
//			FI
//		ENDFOR
//		dst := count
//
// Instruction: 'TZCNTI'. Intrinsic: '_mm_tzcnti_32'.
// Requires KNCNI.
func Tzcnti32(a int, x uint32) int {
	return int(tzcnti32(a, x))
}

func tzcnti32(a int, x uint32) int


// Tzcnti64: Counts the number of trailing bits in unsigned 64-bit integer 'x'
// starting at bit 'a' storing the result in 'dst'. 
//
//		count := 0
//		FOR j := a to 63
//			IF NOT(x[j]  1)
//				count := count + 1
//			FI
//		ENDFOR
//		dst := count
//
// Instruction: 'TZCNTI'. Intrinsic: '_mm_tzcnti_64'.
// Requires KNCNI.
func Tzcnti64(a int64, x uint64) int64 {
	return int64(tzcnti64(a, x))
}

func tzcnti64(a int64, x uint64) int64


// WritefsbaseU32: Write the unsigned 32-bit integer 'a' to the FS segment base
// register. 
//
//		FS_Segment_Base_Register[31:0] := a[31:0];
//		FS_Segment_Base_Register[63:32] := 0
//
// Instruction: 'WRFSBASE'. Intrinsic: '_writefsbase_u32'.
// Requires FSGSBASE.
func WritefsbaseU32(a uint32)  {
	writefsbaseU32(a)
}

func writefsbaseU32(a uint32) 


// WritefsbaseU64: Write the unsigned 64-bit integer 'a' to the FS segment base
// register. 
//
//		FS_Segment_Base_Register[63:0] := a[63:0];
//
// Instruction: 'WRFSBASE'. Intrinsic: '_writefsbase_u64'.
// Requires FSGSBASE.
func WritefsbaseU64(a uint64)  {
	writefsbaseU64(a)
}

func writefsbaseU64(a uint64) 


// WritegsbaseU32: Write the unsigned 32-bit integer 'a' to the GS segment base
// register. 
//
//		GS_Segment_Base_Register[31:0] := a[31:0];
//		GS_Segment_Base_Register[63:32] := 0
//
// Instruction: 'WRGSBASE'. Intrinsic: '_writegsbase_u32'.
// Requires FSGSBASE.
func WritegsbaseU32(a uint32)  {
	writegsbaseU32(a)
}

func writegsbaseU32(a uint32) 


// WritegsbaseU64: Write the unsigned 64-bit integer 'a' to the GS segment base
// register. 
//
//		GS_Segment_Base_Register[63:0] := a[63:0];
//
// Instruction: 'WRGSBASE'. Intrinsic: '_writegsbase_u64'.
// Requires FSGSBASE.
func WritegsbaseU64(a uint64)  {
	writegsbaseU64(a)
}

func writegsbaseU64(a uint64) 


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
func Xabort(imm8 uint)  {
	xabort(imm8)
}

func xabort(imm8 uint) 


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
func Xbegin() uint32 {
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
func Xtest() uint8 {
	return uint8(xtest())
}

func xtest() uint8

