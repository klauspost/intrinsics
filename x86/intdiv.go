package x86

// The instruction set does not support integer vector division. Instead, we
// are using a method for fast integer division based on multiplication and
// shift operations. This method is faster than simple integer division if the
// same divisor is used multiple times.
//
// All elements in a vector are divided by the same divisor. It is not possible
// to divide different elements of the same vector by different divisors.
//
// The parameters used for fast division are stored in an object of a
// Divisor class. This object can be created implicitly, for example in:
//        a, M128i b; int c;
//        a = b / c;
// or explicitly as:
//        a = b / UintDivisor(c);
// encapsulate parameters for fast division on vector of 4 32-bit unsigned integers
type Uint32Div struct {
	multiplier M128i // multiplier used in fast division
	shift1     M128i // shift count 1 used in fast division
	shift2     M128i // shift count 2 used in fast division
}

func (d *Uint32Div) Set(div uint32) {
	var L, L2, sh1, sh2, m uint32
	switch div {
	case 0:
		panic("cannot divide by 0")
	case 1:
		m = 1 // parameters for d = 1
	case 2:
		m = 1
		sh1 = 1 // parameters for d = 2
	default: // general case for d > 2
		L = misc.BitScanReverse(div-1) + 1 // ceil(log2(d))

		// 2^L, overflow to 0 if L = 32
		if L < 32 {
			L2 = 1 << L
		}
		m = 1 + uint32((uint64(L2-div)<<32)/d) // multiplier
		sh1 = 1
		sh2 = L - 1 // shift counts
	}
	d.multiplier = sse2.Set1Epi32(m)
	d.shift1 = sse2.SetEpi32(sh1, 0, 0, 0)
	d.shift2 = sse2.SetEpi32(sh2, 0, 0, 0)
}

// vector of 4 32-bit unsigned integers
func (d Uint32Div) Div(a M128i) M128i {
	t1 := sse2.MulEpu32(a, u.multiplier)  // 32x32->64 bit unsigned multiplication of a[0] and a[2]
	t2 := sse2.SrliEpi64(t1, 32)          // high dword of result 0 and 2
	t3 := sse2.SrliEpi64(a, 32)           // get a[1] and a[3] into position for multiplication
	t4 := sse2.MulEpu32(t3, u.multiplier) // 32x32->64 bit unsigned multiplication of a[1] and a[3]
	t5 := sse2.SetEpi32(-1, 0, -1, 0)     // mask of dword 1 and 3
	t6 := sse2.AndSi128(t4, t5)           // high dword of result 1 and 3
	t7 := sse2.OrSi128(t2, t6)            // combine all four results into one vector
	t8 := sse2.SubEpi32(a, t7)            // subtract
	t9 := sse2.SrlEpi32(t8, u.shift1)     // shift right logical
	t10 := sse2.AddEpi32(t7, t9)          // add
	return sse2.SrlEpi32(t10, d.u.shift2) // shift right logical
}

// vector of 4 32-bit unsigned integers
func (u Uint32Div) DivSSE4(a M128i) M128i {
	t1 := sse2.MulEpu32(a, u.multiplier)  // 32x32->64 bit unsigned multiplication of a[0] and a[2]
	t2 := sse2.SrliEpi64(t1, 32)          // high dword of result 0 and 2
	t3 := sse2.SrliEpi64(a, 32)           // get a[1] and a[3] into position for multiplication
	t4 := sse2.MulEpu32(t3, u.multiplier) // 32x32->64 bit unsigned multiplication of a[1] and a[3]
	t5 := sse2.SetEpi32(-1, 0, -1, 0)     // mask of dword 1 and 3
	t7 := sse4.BlendvEpi8(t2, t4, t5)     // blend two results
	t8 := sse2.SubEpi32(a, t7)            // subtract
	t9 := sse2.SrlEpi32(t8, u.shift1)     // shift right logical
	t10 := sse2.AddEpi32(t7, t9)          // add
	return sse2.SrlEpi32(t10, d.u.shift2) // shift right logical
}
