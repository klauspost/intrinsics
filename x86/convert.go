package x86

// FloatToM128 converts a slice into M128 array.
// The number of elements is len(src) / 4.
// FIXME: Should simply be a pointer to the original data, not a copy
func FloatToM128(src []float32) []M128 {
	l := len(src) / 4
	dst := make([]M128, l)
	for i := 0; i < l; i++ {
		dst[i] = M128{src[i*4], src[i*4+1], src[i*4+2], src[i*4+3]}
	}
	return dst
}

// BytesToM128i converts a byte slice into M128i array.
// The number of elements is len(src) / 16.
// FIXME: Should simply be a pointer to the original data, not a copy
func BytesToM128i(src []byte) []M128i {
	l := len(src) / 16
	dst := make([]M128i, l)
	for i := 0; i < l; i++ {
		d := M128i{}
		for j, n := range src[:16] {
			d[j] = n
		}
		dst[i] = d
		src = src[16:]
	}
	return dst
}

// M128iToBytes converts M128i slice to a byte slice.
// The number of elements is len(src) * 16.
// FIXME: Should simply be a pointer to the original data, not a copy
func M128iToBytes(src []M128i) []byte {
	l := len(src) * 16
	dst := make([]byte, l)
	for i := 0; i < l; i++ {
		for j, n := range src[i] {
			dst[i*16+j] = n
		}
	}
	return dst
}

// DoubleToM128d converts a slice into M128d array.
// The number of elements is len(src) / 2.
// FIXME: Should simply be a pointer to the original data, not a copy
func DoubleToM128d(src []float64) []M128d {
	l := len(src) / 2
	dst := make([]M128d, l)
	for i := 0; i < l; i++ {
		dst[i] = M128d{src[i*2], src[i*2+1]}
	}
	return dst
}
