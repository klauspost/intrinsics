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
	a := sse2.SetEpi32(0x03020100, 0x07060504, 0x0b0a0908, 0x0f0e0d0c)
	b := sse2.SetEpi32(0x03020100, 0x07060504, 0x0b0a0908, 0x0f0e0d0c)

	c := sse2.AddEpi8(a, b)

	expect := x86.M128i{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}

func TestAddPs(t *testing.T) {
	a := sse.SetPs(10.0, 20.0, 30.0, 40.0)
	b := sse.SetPs(100.0, 200.0, 300.0, 400.0)

	c := sse.AddPs(a, b)

	expect := x86.M128{a[0] + b[0], a[1] + b[1], a[2] + b[2], a[3] + b[3]}
	if !reflect.DeepEqual(c, expect) {
		t.Fatal("got", c, "expected", expect)
	}
	t.Log("correctly got", expect)
}

func TestMulPs(t *testing.T) {
	// Nice way to initialize
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
		line := x86.BytesToM128i(input[y*width*2 : y*width*2+width*2])

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

// Float point example.
// If intrinsics where working, this would convert 4 RGB values to HSV.
// Converted from: https://github.com/rawstudio/rawstudio/blob/master/plugins/dcp/dcp-sse4.c#L74
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
	h = zeroPs

	m := sse.MinPs(b, sse.MinPs(r, g))
	gap := sse.SubPs(v, m)
	v_mask := sse.CmpeqPs(gap, zeroPs)
	v = sse.AddPs(v, sse.AndPs(add_v, v_mask))

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

// Will likely strain the register allocator nicely.
func TestMixedSSE(t *testing.T) {
	Transform8sRGB()
}

// Mixed sse+sse2
// Applies a 3x3 matrix and sRGB gamma corrects input
// and writes it to output.
// https://github.com/rawstudio/rawstudio/blob/master/plugins/colorspace-transform/colorspace_transform_sse2.c#L205
func Transform8sRGB() {
	w := 1048
	h := 1024
	var matrix [3][3]float32

	// Input is a w * h * 4 components 16 bit per component image
	var input = make([]byte, w*h*8)

	// Output is 4 byte/pixel, RGBA image
	var output = make([]byte, w*h*4)

	// The matrix with values splatted to all registers
	var matPs = make([]x86.M128, 3*3)

	// Fill with values
	matPs[0] = sse.Set1Ps(matrix[0][0])
	matPs[1] = sse.Set1Ps(matrix[0][1])
	matPs[2] = sse.Set1Ps(matrix[0][2])
	matPs[3] = sse.Set1Ps(matrix[1][0])
	matPs[4] = sse.Set1Ps(matrix[1][1])
	matPs[5] = sse.Set1Ps(matrix[1][2])
	matPs[6] = sse.Set1Ps(matrix[2][0])
	matPs[7] = sse.Set1Ps(matrix[2][1])
	matPs[8] = sse.Set1Ps(matrix[2][2])

	for y := 0; y < h; y++ {
		i := x86.BytesToM128i(input[y*w*8 : y*w*8+w*8])
		o := x86.BytesToM128i(output[y*w*4 : y*w*4+w*4])

		// Converts 4 pixels per loop
		for x := 0; x < w/4; x++ {

			/* Load and convert to float */
			zero := sse2.SetzeroSi128()
			in := i[x*2]    // Load two pixels
			in2 := i[x*2+1] // Load two pixels
			p1 := sse2.UnpackloEpi16(in, zero)
			p2 := sse2.UnpackhiEpi16(in, zero)
			p3 := sse2.UnpackloEpi16(in2, zero)
			p4 := sse2.UnpackhiEpi16(in2, zero)
			p1f := sse2.Cvtepi32Ps(p1)
			p2f := sse2.Cvtepi32Ps(p2)
			p3f := sse2.Cvtepi32Ps(p3)
			p4f := sse2.Cvtepi32Ps(p4)

			/* Convert to planar */
			g1g0r1r0 := sse.UnpackloPs(p1f, p2f)
			b1b0 := sse.UnpackhiPs(p1f, p2f)
			g3g2r3r2 := sse.UnpackloPs(p3f, p4f)
			b3b2 := sse.UnpackhiPs(p3f, p4f)
			r := sse.MovelhPs(g1g0r1r0, g3g2r3r2)
			g := sse.MovehlPs(g3g2r3r2, g1g0r1r0)
			b := sse.MovelhPs(b1b0, b3b2)

			/* Apply matrix to convert to sRGB */
			r = sseMatrix3Mul(matPs[0:3], r, g, b)
			g = sseMatrix3Mul(matPs[3:6], r, g, b)
			b = sseMatrix3Mul(matPs[6:9], r, g, b)

			/* Normalize to 0->1 and clamp */
			normalize := sse.Set1Ps(1.0 / 65535.0)
			max_val := sse.Set1Ps(1.0)
			min_val := sse.SetzeroPs()
			r = sse.MinPs(max_val, sse.MaxPs(min_val, sse.MulPs(normalize, r)))
			g = sse.MinPs(max_val, sse.MaxPs(min_val, sse.MulPs(normalize, g)))
			b = sse.MinPs(max_val, sse.MaxPs(min_val, sse.MulPs(normalize, b)))

			/* Apply Gamma */
			/* Calculate values to be used if larger than junction point */
			mul_over := sse.Set1Ps(1.055)
			sub_over := sse.Set1Ps(0.055)
			pow_over := sse.Set1Ps(1.0 / 2.4)
			r_gam := sse.SubPs(sse.MulPs(mul_over, FastPowPs(r, pow_over)), sub_over)
			g_gam := sse.SubPs(sse.MulPs(mul_over, FastPowPs(g, pow_over)), sub_over)
			b_gam := sse.SubPs(sse.MulPs(mul_over, FastPowPs(b, pow_over)), sub_over)

			/* Create mask for values smaller than junction point */
			junction := sse.Set1Ps(0.0031308)
			mask_r := sse.CmpltPs(r, junction)
			mask_g := sse.CmpltPs(g, junction)
			mask_b := sse.CmpltPs(b, junction)

			/* Calculate value to be used if under junction */
			mul_under := sse.Set1Ps(12.92)
			r_mul := sse.AndPs(mask_r, sse.MulPs(mul_under, r))
			g_mul := sse.AndPs(mask_g, sse.MulPs(mul_under, g))
			b_mul := sse.AndPs(mask_b, sse.MulPs(mul_under, b))

			/* Select the value to be used based on the junction mask and scale to 8 bit */
			upscale := sse.Set1Ps(255.5)
			r = sse.MulPs(upscale, sse.OrPs(r_mul, sse.AndnotPs(mask_r, r_gam)))
			g = sse.MulPs(upscale, sse.OrPs(g_mul, sse.AndnotPs(mask_g, g_gam)))
			b = sse.MulPs(upscale, sse.OrPs(b_mul, sse.AndnotPs(mask_b, b_gam)))

			/* Convert to 8 bit unsigned  and interleave*/
			r_i := sse2.CvtpsEpi32(r)
			g_i := sse2.CvtpsEpi32(g)
			b_i := sse2.CvtpsEpi32(b)

			r_i = sse2.PacksEpi32(r_i, r_i)
			g_i = sse2.PacksEpi32(g_i, g_i)
			b_i = sse2.PacksEpi32(b_i, b_i)

			/* Set alpha value to 255 and store */
			alpha_mask := sse2.Set1Epi32(0xff000000)
			rg_i := sse2.UnpackloEpi16(r_i, g_i)
			bb_i := sse2.UnpackloEpi16(b_i, b_i)
			p1 = sse2.UnpackloEpi32(rg_i, bb_i)
			p2 = sse2.UnpackhiEpi32(rg_i, bb_i)

			p1 = sse2.OrSi128(alpha_mask, sse2.PackusEpi16(p1, p2))

			o[x] = p1
		}
	}
}

// Multiply planar values by a 3x3 matrix
func sseMatrix3Mul(mul []x86.M128, a, b, c x86.M128) x86.M128 {
	acc := sse.MulPs(a, mul[0])
	acc = sse.AddPs(acc, sse.MulPs(b, mul[1]))
	acc = sse.AddPs(acc, sse.MulPs(c, mul[2]))

	return acc
}

// Return x^y for all 4 elements
func FastPowPs(x, y x86.M128) x86.M128 {
	return exp2f4(sse.MulPs(log2f4(x), y))
}

/* SSE2 Polynomial pow function from Mesa3d (MIT License) */

func poly0(x x86.M128, c0 float32) x86.M128 {
	return sse.Set1Ps(c0)
}

func poly1(x x86.M128, c0, c1 float32) x86.M128 {
	return sse.AddPs(sse.MulPs(poly0(x, c1), x), sse.Set1Ps(c0))
}

func poly2(x x86.M128, c0, c1, c2 float32) x86.M128 {
	return sse.AddPs(sse.MulPs(poly1(x, c1, c2), x), sse.Set1Ps(c0))
}

func poly3(x x86.M128, c0, c1, c2, c3 float32) x86.M128 {
	return sse.AddPs(sse.MulPs(poly2(x, c1, c2, c3), x), sse.Set1Ps(c0))
}

func poly4(x x86.M128, c0, c1, c2, c3, c4 float32) x86.M128 {
	return sse.AddPs(sse.MulPs(poly3(x, c1, c2, c3, c4), x), sse.Set1Ps(c0))
}

func poly5(x x86.M128, c0, c1, c2, c3, c4, c5 float32) x86.M128 {
	return sse.AddPs(sse.MulPs(poly4(x, c1, c2, c3, c4, c5), x), sse.Set1Ps(c0))
}

const LOG_poly_DEGREE = 4

const log_p5_0 = 3.1157899
const log_p5_1 = -3.3241990
const log_p5_2 = 2.5988452
const log_p5_3 = -1.2315303
const log_p5_4 = 3.1821337e-1
const log_p5_5 = -3.4436006e-2

const log_p4_0 = 2.8882704548164776201
const log_p4_1 = -2.52074962577807006663
const log_p4_2 = 1.48116647521213171641
const log_p4_3 = -0.465725644288844778798
const log_p4_4 = 0.0596515482674574969533

const log_p3_0 = 2.61761038894603480148
const log_p3_1 = -1.75647175389045657003
const log_p3_2 = 0.688243882994381274313
const log_p3_3 = -0.107254423828329604454

const log_p2_0 = 2.28330284476918490682
const log_p2_1 = -1.04913055217340124191
const log_p2_2 = 0.204446009836232697516

const exp_mask = 0x7F800000
const mantissa_mask = 0x007FFFFF

func log2f4(x x86.M128) x86.M128 {
	exp := sse2.Set1Epi32(exp_mask)
	mant := sse2.Set1Epi32(mantissa_mask)
	one := sse.Set1Ps(1.0)
	i := sse2.CastpsSi128(x)
	e := sse2.Cvtepi32Ps(sse2.SubEpi32(sse2.SrliEpi32(sse2.AndSi128(i, exp), 23), sse2.Set1Epi32(127)))
	m := sse.OrPs(sse2.Castsi128Ps(sse2.AndSi128(i, mant)), one)
	var p x86.M128

	/* Minimax polynomial fit of log2(x)/(x - 1), for x in range [1, 2[ */

	if LOG_poly_DEGREE == 6 {
		p = poly5(m, log_p5_0, log_p5_1, log_p5_2, log_p5_3, log_p5_4, log_p5_5)
	} else if LOG_poly_DEGREE == 5 {
		p = poly4(m, log_p4_0, log_p4_1, log_p4_2, log_p4_3, log_p4_4)
	} else if LOG_poly_DEGREE == 4 {
		p = poly3(m, log_p3_0, log_p3_1, log_p3_2, log_p3_3)
	} else if LOG_poly_DEGREE == 3 {
		p = poly2(m, log_p2_0, log_p2_1, log_p2_2)
	} else {
		panic("unsupported poly degree")
	}

	/* This effectively increases the polynomial degree by one, but ensures that log2(1) == 0*/
	p = sse.MulPs(p, sse.SubPs(m, one))

	return sse.AddPs(p, e)
}

const EXP_poly_DEGREE = 2

const exp_p5_0 = 9.9999994e-1
const exp_p5_1 = 6.9315308e-1
const exp_p5_2 = 2.4015361e-1
const exp_p5_3 = 5.5826318e-2
const exp_p5_4 = 8.9893397e-3
const exp_p5_5 = 1.8775767e-3

const exp_p4_0 = 1.0000026
const exp_p4_1 = 6.9300383e-1
const exp_p4_2 = 2.4144275e-1
const exp_p4_3 = 5.2011464e-2
const exp_p4_4 = 1.3534167e-2

const exp_p3_0 = 9.9992520e-1
const exp_p3_1 = 6.9583356e-1
const exp_p3_2 = 2.2606716e-1
const exp_p3_3 = 7.8024521e-2

const exp_p2_0 = 1.0017247
const exp_p2_1 = 6.5763628e-1
const exp_p2_2 = 3.3718944e-1

func exp2f4(x x86.M128) x86.M128 {
	var ipart x86.M128i
	var fpart, expipart, expfpart x86.M128

	x = sse.MinPs(x, sse.Set1Ps(129))
	x = sse.MaxPs(x, sse.Set1Ps(-126.99999))

	/* ipart = int(x - 0.5) */
	ipart = sse2.CvtpsEpi32(sse.SubPs(x, sse.Set1Ps(0.5)))

	/* fpart = x - ipart */
	fpart = sse.SubPs(x, sse2.Cvtepi32Ps(ipart))

	/* expipart = (float) (1 << ipart) */
	expipart = sse2.Castsi128Ps(sse2.SlliEpi32(sse2.AddEpi32(ipart, sse2.Set1Epi32(127)), 23))

	/* minimax polynomial fit of 2**x, in range [-0.5, 0.5[ */
	if EXP_poly_DEGREE == 5 {
		expfpart = poly5(fpart, exp_p5_0, exp_p5_1, exp_p5_2, exp_p5_3, exp_p5_4, exp_p5_5)
	} else if EXP_poly_DEGREE == 4 {
		expfpart = poly4(fpart, exp_p4_0, exp_p4_1, exp_p4_2, exp_p4_3, exp_p4_4)
	} else if EXP_poly_DEGREE == 3 {
		expfpart = poly3(fpart, exp_p3_0, exp_p3_1, exp_p3_2, exp_p3_3)
	} else if EXP_poly_DEGREE == 2 {
		expfpart = poly2(fpart, exp_p2_0, exp_p2_1, exp_p2_2)
	} else {
		panic("invalid poly degree")
	}

	return sse.MulPs(expipart, expfpart)
}
