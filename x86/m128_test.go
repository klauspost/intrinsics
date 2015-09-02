package x86_test

import (
	"testing"

	"github.com/klauspost/intrinsics/x86"
	"github.com/klauspost/intrinsics/x86/sse"
	"github.com/klauspost/intrinsics/x86/sse2"
	"reflect"
)

func TestAddEpi8(t *testing.T) {
	a := x86.M128i{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b := x86.M128i{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	c := sse2.AddEpi8(a, b)

	expect := x86.M128i{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}

func TestAddPs(t *testing.T) {
	a := x86.M128{10.0, 20.0, 30.0, 40.0}
	b := x86.M128{100.0, 200.0, 300.0, 400.0}

	c := sse.AddPs(a, b)

	expect := x86.M128{a[0] + b[0], a[1] + b[1], a[2] + b[2], a[3] + b[3]}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}

func TestMulPs(t *testing.T) {
	a := x86.M128{1.0, 2.0, 3.0, 4.0}
	b := x86.M128{1000.0, 100.0, 10.0, 1.0}

	c := sse.MulPs(a, b)

	expect := x86.M128{a[0] * b[0], a[1] * b[1], a[2] * b[2], a[3] * b[3]}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}

func TestMulAddPs(t *testing.T) {
	a := x86.M128{1.0, 2.0, 3.0, 4.0}
	b := x86.M128{1000.0, 100.0, 10.0, 1.0}

	c := sse.MulPs(sse.AddPs(a, a), b)

	expect := x86.M128{(a[0] + a[0]) * b[0], (a[1] + a[1]) * b[1], (a[2] + a[2]) * b[2], (a[3] + a[3]) * b[3]}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}
