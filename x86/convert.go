package x86

import (
    "reflect"
    "unsafe"
)

// FloatToM128 converts a slice into M128 array.
// The number of elements is len(src) / 4.
func FloatToM128(src []float32) []M128 {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&src))

	// The length and capacity of the slice are different.
	header.Len /= 4
	header.Cap /= 4

	// Convert slice header to an []int32
	dst := *(*[]M128)(unsafe.Pointer(&header))

	return dst
}

// M128ToFloat converts M128 slice to a float32 slice.
// The number of elements is len(src) * 4.
func M128ToFloat(src []M128) []float32 {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&src))

	// The length and capacity of the slice are different.
	header.Len *= 4
	header.Cap *= 4

	// Convert slice header to an []int32
	dst := *(*[]float32)(unsafe.Pointer(&header))

	return dst
}

// BytesToM128i converts a byte slice into M128i array.
// The number of elements is len(src) / 16.
func BytesToM128i(src []byte) []M128i {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&src))

	// The length and capacity of the slice are different.
	header.Len /= 16
	header.Cap /= 16

	// Convert slice header to an []int32
	dst := *(*[]M128i)(unsafe.Pointer(&header))

	return dst
}

// M128iToBytes converts M128i slice to a byte slice.
// The number of elements is len(src) * 16.
func M128iToBytes(src []M128i) []byte {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&src))

	// The length and capacity of the slice are different.
	header.Len *= 16
	header.Cap *= 16

	// Convert slice header to an []int32
	dst := *(*[]byte)(unsafe.Pointer(&header))

	return dst
}


// DoubleToM128d converts a slice into M128d array.
// The number of elements is len(src) / 2.
func DoubleToM128d(src []float64) []M128d {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&src))

	// The length and capacity of the slice are different.
	header.Len /= 2
	header.Cap /= 2

	// Convert slice header to an []int32
	dst := *(*[]M128d)(unsafe.Pointer(&header))

	return dst
}

// M128ToFloat converts M128 slice to a float64 slice.
// The number of elements is len(src) * 2.
func M128dToDouble(src []M128) []float64 {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&src))

	// The length and capacity of the slice are different.
	header.Len *= 2
	header.Cap *= 2

	// Convert slice header to an []int32
	dst := *(*[]float64)(unsafe.Pointer(&header))

	return dst
}
