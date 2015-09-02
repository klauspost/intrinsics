package x86_test

import (
	"testing"

	"github.com/klauspost/intrinsics/x86"
	"github.com/klauspost/intrinsics/x86/sse"
	"github.com/klauspost/intrinsics/x86/sse2"
	"github.com/klauspost/intrinsics/x86/sse4"
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
	a := sse.SetPs(1.0, 2.0, 3.0, 4.0)
	b := sse.SetPs(1000.0, 100.0, 10.0, 1.0)

	c := sse.MulPs(sse.AddPs(a, a), b)

	expect := x86.M128{(a[0] + a[0]) * b[0], (a[1] + a[1]) * b[1], (a[2] + a[2]) * b[2], (a[3] + a[3]) * b[3]}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}

// Try some complex intrinsics (SSE2)
// Doesn't test any values (since intrisics are unimplemented)
// Converted from https://github.com/klauspost/rawspeed/blob/develop/RawSpeed/RawImageDataU16.cpp#L152
func TestComplex(t *testing.T) {
	full_scale_fp := 1000
	half_scale_fp := 500
	mDitherScale := true
	var sub_mul [4]x86.M128i

	var rand_mul x86.M128i
	sseround := sse2.SetEpi32(512, 512, 512, 512)
	ssesub2 := sse2.SetEpi32(32768, 32768, 32768, 32768)
	ssesign := sse2.SetEpi32(0x80008000, 0x80008000, 0x80008000, 0x80008000)
	sse_full_scale_fp := sse2.Set1Epi32(full_scale_fp | (full_scale_fp << 16))
	sse_half_scale_fp := sse2.Set1Epi32(half_scale_fp >> 4)

	if mDitherScale {
		rand_mul = sse2.Set1Epi32(0x4d9f1d32)
	} else {
		rand_mul = sse2.SetzeroSi128()
	}
	rand_mask := sse2.Set1Epi32(0x00ff00ff) // 8 random bits

	width := 1024
	height := 1024

	// Emulate 1024 x 1024 x 16bpp
	input := make([]byte, 1024*1024*2)

	for y := 0; y < height; y++ {
		// Convert current line to []M128i
		line := x86.Slice2M128i(input[y*width*2:y*width*2+width*2], nil)

		var sserandom x86.M128i
		if mDitherScale {
			sserandom = sse2.SetEpi32(width*1676+y*18000, width*2342+y*34311, width*4272+y*12123, width*1234+y*23464)
		} else {
			sserandom = sse2.SetzeroSi128()
		}

		var ssescale, ssesub x86.M128i
		if (y & 1) == 0 {
			ssesub = sub_mul[0]
			ssescale = sub_mul[1]
		} else {
			ssesub = sub_mul[2]
			ssescale = sub_mul[3]
		}

		for x, pix_low := range line {
			// Subtract black
			pix_low = sse2.SubsEpu16(pix_low, ssesub)
			// Multiply the two unsigned shorts and combine it to 32 bit result
			pix_high := sse2.MulhiEpu16(pix_low, ssescale)

			temp := sse2.MulloEpi16(pix_low, ssescale)

			pix_low = sse2.UnpackloEpi16(temp, pix_high)
			pix_high = sse2.UnpackhiEpi16(temp, pix_high)

			// Add rounder
			pix_low = sse2.AddEpi32(pix_low, sseround)
			pix_high = sse2.AddEpi32(pix_high, sseround)

			sserandom = sse2.XorSi128(sse2.MulhiEpi16(sserandom, rand_mul), sse2.MulloEpi16(sserandom, rand_mul))
			rand_masked := sse2.AndSi128(sserandom, rand_mask) // Get 8 random bits
			rand_masked = sse2.MulloEpi16(rand_masked, sse_full_scale_fp)

			zero := sse2.SetzeroSi128()
			rand_lo := sse2.SubEpi32(sse_half_scale_fp, sse2.UnpackloEpi16(rand_masked, zero))
			rand_hi := sse2.SubEpi32(sse_half_scale_fp, sse2.UnpackhiEpi16(rand_masked, zero))

			pix_low = sse2.AddEpi32(pix_low, rand_lo)
			pix_high = sse2.AddEpi32(pix_high, rand_hi)

			// Shift down
			pix_low = sse2.SraiEpi32(pix_low, 10)
			pix_high = sse2.SraiEpi32(pix_high, 10)

			// Subtract to avoid clipping
			pix_low = sse2.SubEpi32(pix_low, ssesub2)
			pix_high = sse2.SubEpi32(pix_high, ssesub2)

			// Pack
			pix_low = sse2.PacksEpi32(pix_low, pix_high)

			// Shift sign off
			pix_low = sse2.XorSi128(pix_low, ssesign)
			line[x] = pix_low
		}
	}
}

func TestSSE4(t *testing.T) {
	r := x86.M128{0.1, 0.2, 0.3, 0.4}
	g := x86.M128{0.2, 0.3, 0.4, 0.5}
	b := x86.M128{0.3, 0.4, 0.5, 0.6}
	h, s, v := RGBtoHSV(r, g, b)
	t.Log(h, s, v)
}

const verySmall = 1e-15

// Float point example
// If intrinsics where working, this would convert 4 rgb
// values to HSV.
func RGBtoHSV(r, g, b x86.M128) (h, s, v x86.M128) {
	zeroPs := sse.SetzeroPs()
	smallPs := sse.Set1Ps(verySmall)
	onesPs := sse.Set1Ps(1.0)

	// Any number > 1
	add_v := sse.Set1Ps(2.0)

	// Clamp
	r = sse.MinPs(sse.MaxPs(r, smallPs), onesPs)
	g = sse.MinPs(sse.MaxPs(g, smallPs), onesPs)
	b = sse.MinPs(sse.MaxPs(b, smallPs), onesPs)

	v = sse.MaxPs(b, sse.MaxPs(r, g))

	m := sse.MinPs(b, sse.MinPs(r, g))
	gap := sse.SubPs(v, m)
	v_mask := sse.CmpeqPs(gap, zeroPs)
	v = sse.AddPs(v, sse.AndPs(add_v, v_mask))

	h = sse.SetzeroPs()

	// Set gap to one where sat = 0, this will avoid divisions by zero, these values will not be used
	onesPs = sse.AndPs(onesPs, v_mask)
	gap = sse.OrPs(gap, onesPs)

	//  gap_inv = 1.0 / gap
	gap_inv := sse.RcpPs(gap)

	// if r == v
	// h = (g - b) / gap;
	mask := sse.CmpeqPs(r, v)
	val := sse.MulPs(gap_inv, sse.SubPs(g, b))

	// fill h
	v = sse.AddPs(v, sse.AndPs(add_v, mask))
	h = sse4.BlendvPs(h, val, mask)

	// if g == v
	// h = 2.0f + (b - r) / gap;
	twoPs := sse.Set1Ps(2.0)
	mask = sse.CmpeqPs(g, v)
	val = sse.SubPs(b, r)
	val = sse.MulPs(val, gap_inv)
	val = sse.AddPs(val, twoPs)

	v = sse.AddPs(v, sse.AndPs(add_v, mask))
	h = sse4.BlendvPs(h, val, mask)

	// If (b == v)
	// h = 4.0f + (r - g) / gap;
	fourPs := sse.AddPs(twoPs, twoPs)
	mask = sse.CmpeqPs(b, v)
	val = sse.AddPs(fourPs, sse.MulPs(gap_inv, sse.SubPs(r, g)))

	v = sse.AddPs(v, sse.AndPs(add_v, mask))
	h = sse4.BlendvPs(h, val, mask)

	// Fill s, if gap > 0
	v = sse.SubPs(v, add_v)
	val = sse.MulPs(gap, sse.RcpPs(v))
	s = sse.AndnotPs(v_mask, val)

	// Check if h < 0
	zeroPs = sse.SetzeroPs()
	sixPs := sse.Set1Ps(6.0 - verySmall)
	mask = sse.CmpltPs(h, zeroPs)
	h = sse.AddPs(h, sse.AndPs(mask, sixPs))
	return
}
