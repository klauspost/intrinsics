package x86

// Slice2M128 converts a slice into M128 array.
// The number of elements is len(src) / 4.
// You can supply a destination, or supply nil to create a new one.
// FIXME: Should simply be a pointer to the original data, not a copy
func Slice2M128(src []float32, dst []M128) []M128 {
	l := len(src) / 4
	if dst == nil {
		dst = make([]M128, l)
	} else {
		dst = dst[:l]
	}
	for i := 0; i < l; i++ {
		dst[i] = M128{src[i*4], src[i*4+1], src[i*4+2], src[i*4+3]}
	}
	return dst
}

// Slice2M128i converts a byte slice into M128i array.
// The number of elements is len(src) / 16.
// You can supply a destination, or supply nil to create a new one.
// FIXME: Should simply be a pointer to the original data, not a copy
func Slice2M128i(src []byte, dst []M128i) []M128i {
	l := len(src) / 16
	if dst == nil {
		dst = make([]M128i, l)
	} else {
		dst = dst[:l]
	}
	for i := 0; i < l; i++ {
		d := M128i{}
		for j, n := range src[:16] {
			d[j] = n
		}
		dst[i] = d
	}
	return dst
}

// Slice2M128d converts a slice into M128d array.
// The number of elements is len(src) / 2.
// You can supply a destination, or supply nil to create a new one.
// FIXME: Should simply be a pointer to the original data, not a copy
func Slice2M128d(src []float64, dst []M128d) []M128d {
	l := len(src) / 2
	if dst == nil {
		dst = make([]M128d, l)
	} else {
		dst = dst[:l]
	}
	for i := 0; i < l; i++ {
		dst[i] = M128d{src[i*2], src[i*2+1]}
	}
	return dst
}
