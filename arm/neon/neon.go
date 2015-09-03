package neon

import "github.com/klauspost/intrinsics/arm"

var _ = arm.Int8x8{}  // Make sure we use arm package


// VaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s8'.
// Requires NEON.
func VaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s16'.
// Requires NEON.
func VaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s32'.
// Requires NEON.
func VaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VaddF32: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_f32'.
// Requires NEON.
func VaddF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VaddF64: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_f64'.
// Requires NEON.
func VaddF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u8'.
// Requires NEON.
func VaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u16'.
// Requires NEON.
func VaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u32'.
// Requires NEON.
func VaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VaddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s64'.
// Requires NEON.
func VaddS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VaddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u64'.
// Requires NEON.
func VaddU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s8'.
// Requires NEON.
func VaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s16'.
// Requires NEON.
func VaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s32'.
// Requires NEON.
func VaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s64'.
// Requires NEON.
func VaddqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VaddqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_f32'.
// Requires NEON.
func VaddqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VaddqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_f64'.
// Requires NEON.
func VaddqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u8'.
// Requires NEON.
func VaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u16'.
// Requires NEON.
func VaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u32'.
// Requires NEON.
func VaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u64'.
// Requires NEON.
func VaddqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VaddlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_s8'.
// Requires NEON.
func VaddlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VaddlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_s16'.
// Requires NEON.
func VaddlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VaddlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_s32'.
// Requires NEON.
func VaddlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VaddlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_u8'.
// Requires NEON.
func VaddlU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VaddlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_u16'.
// Requires NEON.
func VaddlU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VaddlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_u32'.
// Requires NEON.
func VaddlU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VaddlHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_s8'.
// Requires NEON.
func VaddlHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VaddlHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_s16'.
// Requires NEON.
func VaddlHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VaddlHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_s32'.
// Requires NEON.
func VaddlHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VaddlHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_u8'.
// Requires NEON.
func VaddlHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VaddlHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_u16'.
// Requires NEON.
func VaddlHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VaddlHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_u32'.
// Requires NEON.
func VaddlHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VaddwS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_s8'.
// Requires NEON.
func VaddwS8(a arm.Int16x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VaddwS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_s16'.
// Requires NEON.
func VaddwS16(a arm.Int32x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VaddwS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_s32'.
// Requires NEON.
func VaddwS32(a arm.Int64x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VaddwU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_u8'.
// Requires NEON.
func VaddwU8(a arm.Uint16x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VaddwU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_u16'.
// Requires NEON.
func VaddwU16(a arm.Uint32x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VaddwU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_u32'.
// Requires NEON.
func VaddwU32(a arm.Uint64x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VaddwHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_s8'.
// Requires NEON.
func VaddwHighS8(a arm.Int16x8, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VaddwHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_s16'.
// Requires NEON.
func VaddwHighS16(a arm.Int32x4, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VaddwHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_s32'.
// Requires NEON.
func VaddwHighS32(a arm.Int64x2, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VaddwHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_u8'.
// Requires NEON.
func VaddwHighU8(a arm.Uint16x8, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VaddwHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_u16'.
// Requires NEON.
func VaddwHighU16(a arm.Uint32x4, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VaddwHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_u32'.
// Requires NEON.
func VaddwHighU32(a arm.Uint64x2, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VhaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_s8'.
// Requires NEON.
func VhaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VhaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_s16'.
// Requires NEON.
func VhaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VhaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_s32'.
// Requires NEON.
func VhaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VhaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_u8'.
// Requires NEON.
func VhaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VhaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_u16'.
// Requires NEON.
func VhaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VhaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_u32'.
// Requires NEON.
func VhaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VhaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_s8'.
// Requires NEON.
func VhaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VhaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_s16'.
// Requires NEON.
func VhaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VhaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_s32'.
// Requires NEON.
func VhaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VhaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_u8'.
// Requires NEON.
func VhaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VhaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_u16'.
// Requires NEON.
func VhaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VhaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_u32'.
// Requires NEON.
func VhaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrhaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_s8'.
// Requires NEON.
func VrhaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VrhaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_s16'.
// Requires NEON.
func VrhaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VrhaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_s32'.
// Requires NEON.
func VrhaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VrhaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_u8'.
// Requires NEON.
func VrhaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VrhaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_u16'.
// Requires NEON.
func VrhaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VrhaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_u32'.
// Requires NEON.
func VrhaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrhaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_s8'.
// Requires NEON.
func VrhaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VrhaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_s16'.
// Requires NEON.
func VrhaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VrhaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_s32'.
// Requires NEON.
func VrhaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VrhaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_u8'.
// Requires NEON.
func VrhaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VrhaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_u16'.
// Requires NEON.
func VrhaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VrhaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_u32'.
// Requires NEON.
func VrhaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VaddhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_s16'.
// Requires NEON.
func VaddhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VaddhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_s32'.
// Requires NEON.
func VaddhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VaddhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_s64'.
// Requires NEON.
func VaddhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VaddhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_u16'.
// Requires NEON.
func VaddhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VaddhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_u32'.
// Requires NEON.
func VaddhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VaddhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_u64'.
// Requires NEON.
func VaddhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VraddhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_s16'.
// Requires NEON.
func VraddhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VraddhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_s32'.
// Requires NEON.
func VraddhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VraddhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_s64'.
// Requires NEON.
func VraddhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VraddhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_u16'.
// Requires NEON.
func VraddhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VraddhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_u32'.
// Requires NEON.
func VraddhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VraddhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_u64'.
// Requires NEON.
func VraddhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VaddhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_s16'.
// Requires NEON.
func VaddhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VaddhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_s32'.
// Requires NEON.
func VaddhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VaddhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_s64'.
// Requires NEON.
func VaddhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VaddhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_u16'.
// Requires NEON.
func VaddhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VaddhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_u32'.
// Requires NEON.
func VaddhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VaddhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_u64'.
// Requires NEON.
func VaddhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VraddhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_s16'.
// Requires NEON.
func VraddhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VraddhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_s32'.
// Requires NEON.
func VraddhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VraddhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_s64'.
// Requires NEON.
func VraddhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VraddhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_u16'.
// Requires NEON.
func VraddhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VraddhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_u32'.
// Requires NEON.
func VraddhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VraddhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_u64'.
// Requires NEON.
func VraddhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VdivF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdiv_f32'.
// Requires NEON.
func VdivF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VdivF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdiv_f64'.
// Requires NEON.
func VdivF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VdivqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdivq_f32'.
// Requires NEON.
func VdivqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VdivqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdivq_f64'.
// Requires NEON.
func VdivqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmulS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_s8'.
// Requires NEON.
func VmulS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmulS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_s16'.
// Requires NEON.
func VmulS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmulS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_s32'.
// Requires NEON.
func VmulS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmulF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_f32'.
// Requires NEON.
func VmulF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmulF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_f64'.
// Requires NEON.
func VmulF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmulU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_u8'.
// Requires NEON.
func VmulU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmulU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_u16'.
// Requires NEON.
func VmulU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmulU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_u32'.
// Requires NEON.
func VmulU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmulP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_p8'.
// Requires NEON.
func VmulP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VmulqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_s8'.
// Requires NEON.
func VmulqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmulqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_s16'.
// Requires NEON.
func VmulqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmulqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_s32'.
// Requires NEON.
func VmulqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmulqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_f32'.
// Requires NEON.
func VmulqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmulqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_f64'.
// Requires NEON.
func VmulqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmulqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_u8'.
// Requires NEON.
func VmulqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmulqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_u16'.
// Requires NEON.
func VmulqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmulqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_u32'.
// Requires NEON.
func VmulqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmulqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_p8'.
// Requires NEON.
func VmulqP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VandS8: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s8'.
// Requires NEON.
func VandS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VandS16: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s16'.
// Requires NEON.
func VandS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VandS32: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s32'.
// Requires NEON.
func VandS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VandU8: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u8'.
// Requires NEON.
func VandU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VandU16: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u16'.
// Requires NEON.
func VandU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VandU32: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u32'.
// Requires NEON.
func VandU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VandS64: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s64'.
// Requires NEON.
func VandS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VandU64: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u64'.
// Requires NEON.
func VandU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VandqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s8'.
// Requires NEON.
func VandqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VandqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s16'.
// Requires NEON.
func VandqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VandqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s32'.
// Requires NEON.
func VandqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VandqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s64'.
// Requires NEON.
func VandqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VandqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u8'.
// Requires NEON.
func VandqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VandqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u16'.
// Requires NEON.
func VandqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VandqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u32'.
// Requires NEON.
func VandqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VandqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u64'.
// Requires NEON.
func VandqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VorrS8: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s8'.
// Requires NEON.
func VorrS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VorrS16: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s16'.
// Requires NEON.
func VorrS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VorrS32: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s32'.
// Requires NEON.
func VorrS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VorrU8: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u8'.
// Requires NEON.
func VorrU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VorrU16: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u16'.
// Requires NEON.
func VorrU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VorrU32: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u32'.
// Requires NEON.
func VorrU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VorrS64: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s64'.
// Requires NEON.
func VorrS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VorrU64: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u64'.
// Requires NEON.
func VorrU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VorrqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s8'.
// Requires NEON.
func VorrqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VorrqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s16'.
// Requires NEON.
func VorrqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VorrqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s32'.
// Requires NEON.
func VorrqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VorrqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s64'.
// Requires NEON.
func VorrqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VorrqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u8'.
// Requires NEON.
func VorrqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VorrqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u16'.
// Requires NEON.
func VorrqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VorrqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u32'.
// Requires NEON.
func VorrqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VorrqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u64'.
// Requires NEON.
func VorrqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VeorS8: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s8'.
// Requires NEON.
func VeorS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VeorS16: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s16'.
// Requires NEON.
func VeorS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VeorS32: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s32'.
// Requires NEON.
func VeorS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VeorU8: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u8'.
// Requires NEON.
func VeorU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VeorU16: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u16'.
// Requires NEON.
func VeorU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VeorU32: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u32'.
// Requires NEON.
func VeorU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VeorS64: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s64'.
// Requires NEON.
func VeorS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VeorU64: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u64'.
// Requires NEON.
func VeorU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VeorqS8: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s8'.
// Requires NEON.
func VeorqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VeorqS16: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s16'.
// Requires NEON.
func VeorqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VeorqS32: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s32'.
// Requires NEON.
func VeorqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VeorqS64: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s64'.
// Requires NEON.
func VeorqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VeorqU8: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u8'.
// Requires NEON.
func VeorqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VeorqU16: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u16'.
// Requires NEON.
func VeorqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VeorqU32: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u32'.
// Requires NEON.
func VeorqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VeorqU64: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u64'.
// Requires NEON.
func VeorqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VbicS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s8'.
// Requires NEON.
func VbicS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VbicS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s16'.
// Requires NEON.
func VbicS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VbicS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s32'.
// Requires NEON.
func VbicS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VbicU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u8'.
// Requires NEON.
func VbicU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VbicU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u16'.
// Requires NEON.
func VbicU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VbicU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u32'.
// Requires NEON.
func VbicU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VbicS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s64'.
// Requires NEON.
func VbicS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VbicU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u64'.
// Requires NEON.
func VbicU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VbicqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s8'.
// Requires NEON.
func VbicqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VbicqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s16'.
// Requires NEON.
func VbicqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VbicqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s32'.
// Requires NEON.
func VbicqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VbicqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s64'.
// Requires NEON.
func VbicqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VbicqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u8'.
// Requires NEON.
func VbicqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VbicqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u16'.
// Requires NEON.
func VbicqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VbicqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u32'.
// Requires NEON.
func VbicqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VbicqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u64'.
// Requires NEON.
func VbicqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VornS8: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s8'.
// Requires NEON.
func VornS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VornS16: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s16'.
// Requires NEON.
func VornS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VornS32: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s32'.
// Requires NEON.
func VornS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VornU8: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u8'.
// Requires NEON.
func VornU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VornU16: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u16'.
// Requires NEON.
func VornU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VornU32: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u32'.
// Requires NEON.
func VornU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VornS64: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s64'.
// Requires NEON.
func VornS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VornU64: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u64'.
// Requires NEON.
func VornU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VornqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s8'.
// Requires NEON.
func VornqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VornqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s16'.
// Requires NEON.
func VornqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VornqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s32'.
// Requires NEON.
func VornqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VornqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s64'.
// Requires NEON.
func VornqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VornqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u8'.
// Requires NEON.
func VornqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VornqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u16'.
// Requires NEON.
func VornqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VornqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u32'.
// Requires NEON.
func VornqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VornqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u64'.
// Requires NEON.
func VornqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsubS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s8'.
// Requires NEON.
func VsubS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VsubS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s16'.
// Requires NEON.
func VsubS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VsubS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s32'.
// Requires NEON.
func VsubS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VsubF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_f32'.
// Requires NEON.
func VsubF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VsubF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_f64'.
// Requires NEON.
func VsubF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VsubU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u8'.
// Requires NEON.
func VsubU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsubU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u16'.
// Requires NEON.
func VsubU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsubU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u32'.
// Requires NEON.
func VsubU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VsubS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s64'.
// Requires NEON.
func VsubS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VsubU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u64'.
// Requires NEON.
func VsubU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VsubqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s8'.
// Requires NEON.
func VsubqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VsubqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s16'.
// Requires NEON.
func VsubqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsubqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s32'.
// Requires NEON.
func VsubqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsubqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s64'.
// Requires NEON.
func VsubqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsubqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_f32'.
// Requires NEON.
func VsubqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VsubqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_f64'.
// Requires NEON.
func VsubqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VsubqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u8'.
// Requires NEON.
func VsubqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsubqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u16'.
// Requires NEON.
func VsubqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsubqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u32'.
// Requires NEON.
func VsubqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsubqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u64'.
// Requires NEON.
func VsubqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsublS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_s8'.
// Requires NEON.
func VsublS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsublS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_s16'.
// Requires NEON.
func VsublS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsublS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_s32'.
// Requires NEON.
func VsublS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsublU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_u8'.
// Requires NEON.
func VsublU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsublU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_u16'.
// Requires NEON.
func VsublU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsublU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_u32'.
// Requires NEON.
func VsublU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsublHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_s8'.
// Requires NEON.
func VsublHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsublHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_s16'.
// Requires NEON.
func VsublHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsublHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_s32'.
// Requires NEON.
func VsublHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsublHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_u8'.
// Requires NEON.
func VsublHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsublHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_u16'.
// Requires NEON.
func VsublHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsublHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_u32'.
// Requires NEON.
func VsublHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsubwS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_s8'.
// Requires NEON.
func VsubwS8(a arm.Int16x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsubwS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_s16'.
// Requires NEON.
func VsubwS16(a arm.Int32x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsubwS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_s32'.
// Requires NEON.
func VsubwS32(a arm.Int64x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsubwU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_u8'.
// Requires NEON.
func VsubwU8(a arm.Uint16x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsubwU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_u16'.
// Requires NEON.
func VsubwU16(a arm.Uint32x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsubwU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_u32'.
// Requires NEON.
func VsubwU32(a arm.Uint64x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsubwHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_s8'.
// Requires NEON.
func VsubwHighS8(a arm.Int16x8, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsubwHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_s16'.
// Requires NEON.
func VsubwHighS16(a arm.Int32x4, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsubwHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_s32'.
// Requires NEON.
func VsubwHighS32(a arm.Int64x2, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsubwHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_u8'.
// Requires NEON.
func VsubwHighU8(a arm.Uint16x8, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsubwHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_u16'.
// Requires NEON.
func VsubwHighU16(a arm.Uint32x4, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsubwHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_u32'.
// Requires NEON.
func VsubwHighU32(a arm.Uint64x2, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s8'.
// Requires NEON.
func VqaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s16'.
// Requires NEON.
func VqaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s32'.
// Requires NEON.
func VqaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqaddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s64'.
// Requires NEON.
func VqaddS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u8'.
// Requires NEON.
func VqaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VhsubS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_s8'.
// Requires NEON.
func VhsubS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VhsubS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_s16'.
// Requires NEON.
func VhsubS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VhsubS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_s32'.
// Requires NEON.
func VhsubS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VhsubU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_u8'.
// Requires NEON.
func VhsubU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VhsubU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_u16'.
// Requires NEON.
func VhsubU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VhsubU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_u32'.
// Requires NEON.
func VhsubU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VhsubqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_s8'.
// Requires NEON.
func VhsubqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VhsubqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_s16'.
// Requires NEON.
func VhsubqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VhsubqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_s32'.
// Requires NEON.
func VhsubqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VhsubqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_u8'.
// Requires NEON.
func VhsubqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VhsubqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_u16'.
// Requires NEON.
func VhsubqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VhsubqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_u32'.
// Requires NEON.
func VhsubqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsubhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_s16'.
// Requires NEON.
func VsubhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VsubhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_s32'.
// Requires NEON.
func VsubhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VsubhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_s64'.
// Requires NEON.
func VsubhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VsubhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_u16'.
// Requires NEON.
func VsubhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsubhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_u32'.
// Requires NEON.
func VsubhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsubhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_u64'.
// Requires NEON.
func VsubhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrsubhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_s16'.
// Requires NEON.
func VrsubhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VrsubhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_s32'.
// Requires NEON.
func VrsubhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VrsubhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_s64'.
// Requires NEON.
func VrsubhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VrsubhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_u16'.
// Requires NEON.
func VrsubhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VrsubhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_u32'.
// Requires NEON.
func VrsubhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VrsubhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_u64'.
// Requires NEON.
func VrsubhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrsubhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_s16'.
// Requires NEON.
func VrsubhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VrsubhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_s32'.
// Requires NEON.
func VrsubhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VrsubhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_s64'.
// Requires NEON.
func VrsubhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VrsubhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_u16'.
// Requires NEON.
func VrsubhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VrsubhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_u32'.
// Requires NEON.
func VrsubhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VrsubhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_u64'.
// Requires NEON.
func VrsubhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsubhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_s16'.
// Requires NEON.
func VsubhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VsubhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_s32'.
// Requires NEON.
func VsubhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsubhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_s64'.
// Requires NEON.
func VsubhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsubhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_u16'.
// Requires NEON.
func VsubhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsubhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_u32'.
// Requires NEON.
func VsubhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsubhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_u64'.
// Requires NEON.
func VsubhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u16'.
// Requires NEON.
func VqaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u32'.
// Requires NEON.
func VqaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqaddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u64'.
// Requires NEON.
func VqaddU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VqaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s8'.
// Requires NEON.
func VqaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s16'.
// Requires NEON.
func VqaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s32'.
// Requires NEON.
func VqaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s64'.
// Requires NEON.
func VqaddqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u8'.
// Requires NEON.
func VqaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u16'.
// Requires NEON.
func VqaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u32'.
// Requires NEON.
func VqaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u64'.
// Requires NEON.
func VqaddqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqsubS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s8'.
// Requires NEON.
func VqsubS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqsubS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s16'.
// Requires NEON.
func VqsubS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqsubS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s32'.
// Requires NEON.
func VqsubS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqsubS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s64'.
// Requires NEON.
func VqsubS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqsubU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u8'.
// Requires NEON.
func VqsubU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqsubU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u16'.
// Requires NEON.
func VqsubU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqsubU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u32'.
// Requires NEON.
func VqsubU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqsubU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u64'.
// Requires NEON.
func VqsubU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VqsubqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s8'.
// Requires NEON.
func VqsubqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqsubqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s16'.
// Requires NEON.
func VqsubqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqsubqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s32'.
// Requires NEON.
func VqsubqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqsubqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s64'.
// Requires NEON.
func VqsubqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqsubqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u8'.
// Requires NEON.
func VqsubqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqsubqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u16'.
// Requires NEON.
func VqsubqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqsubqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u32'.
// Requires NEON.
func VqsubqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqsubqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u64'.
// Requires NEON.
func VqsubqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqnegS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s8'.
// Requires NEON.
func VqnegS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqnegS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s16'.
// Requires NEON.
func VqnegS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqnegS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s32'.
// Requires NEON.
func VqnegS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqnegS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s64'.
// Requires NEON.
func VqnegS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqnegqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s8'.
// Requires NEON.
func VqnegqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqnegqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s16'.
// Requires NEON.
func VqnegqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqnegqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s32'.
// Requires NEON.
func VqnegqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqabsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s8'.
// Requires NEON.
func VqabsS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqabsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s16'.
// Requires NEON.
func VqabsS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqabsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s32'.
// Requires NEON.
func VqabsS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqabsS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s64'.
// Requires NEON.
func VqabsS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqabsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s8'.
// Requires NEON.
func VqabsqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqabsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s16'.
// Requires NEON.
func VqabsqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqabsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s32'.
// Requires NEON.
func VqabsqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmulhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_s16'.
// Requires NEON.
func VqdmulhS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqdmulhS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_s32'.
// Requires NEON.
func VqdmulhS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqdmulhqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_s16'.
// Requires NEON.
func VqdmulhqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqdmulhqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_s32'.
// Requires NEON.
func VqdmulhqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqrdmulhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_s16'.
// Requires NEON.
func VqrdmulhS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqrdmulhS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_s32'.
// Requires NEON.
func VqrdmulhS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqrdmulhqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_s16'.
// Requires NEON.
func VqrdmulhqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqrdmulhqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_s32'.
// Requires NEON.
func VqrdmulhqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcreateS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s8'.
// Requires NEON.
func VcreateS8(a uint64) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VcreateS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s16'.
// Requires NEON.
func VcreateS16(a uint64) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VcreateS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s32'.
// Requires NEON.
func VcreateS32(a uint64) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VcreateS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s64'.
// Requires NEON.
func VcreateS64(a uint64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VcreateF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_f32'.
// Requires NEON.
func VcreateF32(a uint64) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VcreateU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u8'.
// Requires NEON.
func VcreateU8(a uint64) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcreateU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u16'.
// Requires NEON.
func VcreateU16(a uint64) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcreateU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u32'.
// Requires NEON.
func VcreateU32(a uint64) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcreateU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u64'.
// Requires NEON.
func VcreateU64(a uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcreateF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_f64'.
// Requires NEON.
func VcreateF64(a uint64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VcreateP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_p8'.
// Requires NEON.
func VcreateP8(a uint64) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VcreateP16: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_p16'.
// Requires NEON.
func VcreateP16(a uint64) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VgetLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneF32(a arm.Float32x2, b int) float32 {
	return 0
}

// VgetLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneF64(a arm.Float64x1, b int) float64 {
	return 0
}

// VgetLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// VgetLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// VgetLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneS8(a arm.Int8x8, b int) int8 {
	return 0
}

// VgetLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneS16(a arm.Int16x4, b int) int16 {
	return 0
}

// VgetLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneS32(a arm.Int32x2, b int) int32 {
	return 0
}

// VgetLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneS64(a arm.Int64x1, b int) int64 {
	return 0
}

// VgetLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneU8(a arm.Uint8x8, b int) uint8 {
	return 0
}

// VgetLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneU16(a arm.Uint16x4, b int) uint16 {
	return 0
}

// VgetLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneU32(a arm.Uint32x2, b int) uint32 {
	return 0
}

// VgetLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetLaneU64(a arm.Uint64x1, b int) uint64 {
	return 0
}

// VgetqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneF32(a arm.Float32x4, b int) float32 {
	return 0
}

// VgetqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneF64(a arm.Float64x2, b int) float64 {
	return 0
}

// VgetqLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneP8(a arm.Poly8x16, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// VgetqLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneP16(a arm.Poly16x8, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// VgetqLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneS8(a arm.Int8x16, b int) int8 {
	return 0
}

// VgetqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneS16(a arm.Int16x8, b int) int16 {
	return 0
}

// VgetqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneS32(a arm.Int32x4, b int) int32 {
	return 0
}

// VgetqLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneS64(a arm.Int64x2, b int) int64 {
	return 0
}

// VgetqLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneU8(a arm.Uint8x16, b int) uint8 {
	return 0
}

// VgetqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneU16(a arm.Uint16x8, b int) uint16 {
	return 0
}

// VgetqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneU32(a arm.Uint32x4, b int) uint32 {
	return 0
}

// VgetqLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VgetqLaneU64(a arm.Uint64x2, b int) uint64 {
	return 0
}

// VreinterpretP8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_f64'.
// Requires NEON.
func VreinterpretP8F64(a arm.Float64x1) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s8'.
// Requires NEON.
func VreinterpretP8S8(a arm.Int8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s16'.
// Requires NEON.
func VreinterpretP8S16(a arm.Int16x4) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s32'.
// Requires NEON.
func VreinterpretP8S32(a arm.Int32x2) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s64'.
// Requires NEON.
func VreinterpretP8S64(a arm.Int64x1) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_f32'.
// Requires NEON.
func VreinterpretP8F32(a arm.Float32x2) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u8'.
// Requires NEON.
func VreinterpretP8U8(a arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u16'.
// Requires NEON.
func VreinterpretP8U16(a arm.Uint16x4) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u32'.
// Requires NEON.
func VreinterpretP8U32(a arm.Uint32x2) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u64'.
// Requires NEON.
func VreinterpretP8U64(a arm.Uint64x1) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretP8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_p16'.
// Requires NEON.
func VreinterpretP8P16(a arm.Poly16x4) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VreinterpretqP8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_f64'.
// Requires NEON.
func VreinterpretqP8F64(a arm.Float64x2) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s8'.
// Requires NEON.
func VreinterpretqP8S8(a arm.Int8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s16'.
// Requires NEON.
func VreinterpretqP8S16(a arm.Int16x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s32'.
// Requires NEON.
func VreinterpretqP8S32(a arm.Int32x4) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s64'.
// Requires NEON.
func VreinterpretqP8S64(a arm.Int64x2) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_f32'.
// Requires NEON.
func VreinterpretqP8F32(a arm.Float32x4) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u8'.
// Requires NEON.
func VreinterpretqP8U8(a arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u16'.
// Requires NEON.
func VreinterpretqP8U16(a arm.Uint16x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u32'.
// Requires NEON.
func VreinterpretqP8U32(a arm.Uint32x4) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u64'.
// Requires NEON.
func VreinterpretqP8U64(a arm.Uint64x2) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretqP8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_p16'.
// Requires NEON.
func VreinterpretqP8P16(a arm.Poly16x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VreinterpretP16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_f64'.
// Requires NEON.
func VreinterpretP16F64(a arm.Float64x1) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s8'.
// Requires NEON.
func VreinterpretP16S8(a arm.Int8x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s16'.
// Requires NEON.
func VreinterpretP16S16(a arm.Int16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s32'.
// Requires NEON.
func VreinterpretP16S32(a arm.Int32x2) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s64'.
// Requires NEON.
func VreinterpretP16S64(a arm.Int64x1) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_f32'.
// Requires NEON.
func VreinterpretP16F32(a arm.Float32x2) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u8'.
// Requires NEON.
func VreinterpretP16U8(a arm.Uint8x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u16'.
// Requires NEON.
func VreinterpretP16U16(a arm.Uint16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u32'.
// Requires NEON.
func VreinterpretP16U32(a arm.Uint32x2) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u64'.
// Requires NEON.
func VreinterpretP16U64(a arm.Uint64x1) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretP16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_p8'.
// Requires NEON.
func VreinterpretP16P8(a arm.Poly8x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VreinterpretqP16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_f64'.
// Requires NEON.
func VreinterpretqP16F64(a arm.Float64x2) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s8'.
// Requires NEON.
func VreinterpretqP16S8(a arm.Int8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s16'.
// Requires NEON.
func VreinterpretqP16S16(a arm.Int16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s32'.
// Requires NEON.
func VreinterpretqP16S32(a arm.Int32x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s64'.
// Requires NEON.
func VreinterpretqP16S64(a arm.Int64x2) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_f32'.
// Requires NEON.
func VreinterpretqP16F32(a arm.Float32x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u8'.
// Requires NEON.
func VreinterpretqP16U8(a arm.Uint8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u16'.
// Requires NEON.
func VreinterpretqP16U16(a arm.Uint16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u32'.
// Requires NEON.
func VreinterpretqP16U32(a arm.Uint32x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u64'.
// Requires NEON.
func VreinterpretqP16U64(a arm.Uint64x2) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretqP16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_p8'.
// Requires NEON.
func VreinterpretqP16P8(a arm.Poly8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VreinterpretF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_f64'.
// Requires NEON.
func VreinterpretF32F64(a arm.Float64x1) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s8'.
// Requires NEON.
func VreinterpretF32S8(a arm.Int8x8) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s16'.
// Requires NEON.
func VreinterpretF32S16(a arm.Int16x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s32'.
// Requires NEON.
func VreinterpretF32S32(a arm.Int32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s64'.
// Requires NEON.
func VreinterpretF32S64(a arm.Int64x1) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u8'.
// Requires NEON.
func VreinterpretF32U8(a arm.Uint8x8) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u16'.
// Requires NEON.
func VreinterpretF32U16(a arm.Uint16x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u32'.
// Requires NEON.
func VreinterpretF32U32(a arm.Uint32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u64'.
// Requires NEON.
func VreinterpretF32U64(a arm.Uint64x1) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_p8'.
// Requires NEON.
func VreinterpretF32P8(a arm.Poly8x8) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretF32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_p16'.
// Requires NEON.
func VreinterpretF32P16(a arm.Poly16x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VreinterpretqF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_f64'.
// Requires NEON.
func VreinterpretqF32F64(a arm.Float64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s8'.
// Requires NEON.
func VreinterpretqF32S8(a arm.Int8x16) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s16'.
// Requires NEON.
func VreinterpretqF32S16(a arm.Int16x8) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s32'.
// Requires NEON.
func VreinterpretqF32S32(a arm.Int32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s64'.
// Requires NEON.
func VreinterpretqF32S64(a arm.Int64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u8'.
// Requires NEON.
func VreinterpretqF32U8(a arm.Uint8x16) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u16'.
// Requires NEON.
func VreinterpretqF32U16(a arm.Uint16x8) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u32'.
// Requires NEON.
func VreinterpretqF32U32(a arm.Uint32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u64'.
// Requires NEON.
func VreinterpretqF32U64(a arm.Uint64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_p8'.
// Requires NEON.
func VreinterpretqF32P8(a arm.Poly8x16) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretqF32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_p16'.
// Requires NEON.
func VreinterpretqF32P16(a arm.Poly16x8) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VreinterpretF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_f32'.
// Requires NEON.
func VreinterpretF64F32(a arm.Float32x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_p8'.
// Requires NEON.
func VreinterpretF64P8(a arm.Poly8x8) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_p16'.
// Requires NEON.
func VreinterpretF64P16(a arm.Poly16x4) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s8'.
// Requires NEON.
func VreinterpretF64S8(a arm.Int8x8) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s16'.
// Requires NEON.
func VreinterpretF64S16(a arm.Int16x4) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s32'.
// Requires NEON.
func VreinterpretF64S32(a arm.Int32x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s64'.
// Requires NEON.
func VreinterpretF64S64(a arm.Int64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u8'.
// Requires NEON.
func VreinterpretF64U8(a arm.Uint8x8) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u16'.
// Requires NEON.
func VreinterpretF64U16(a arm.Uint16x4) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u32'.
// Requires NEON.
func VreinterpretF64U32(a arm.Uint32x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u64'.
// Requires NEON.
func VreinterpretF64U64(a arm.Uint64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VreinterpretqF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_f32'.
// Requires NEON.
func VreinterpretqF64F32(a arm.Float32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_p8'.
// Requires NEON.
func VreinterpretqF64P8(a arm.Poly8x16) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_p16'.
// Requires NEON.
func VreinterpretqF64P16(a arm.Poly16x8) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s8'.
// Requires NEON.
func VreinterpretqF64S8(a arm.Int8x16) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s16'.
// Requires NEON.
func VreinterpretqF64S16(a arm.Int16x8) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s32'.
// Requires NEON.
func VreinterpretqF64S32(a arm.Int32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s64'.
// Requires NEON.
func VreinterpretqF64S64(a arm.Int64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u8'.
// Requires NEON.
func VreinterpretqF64U8(a arm.Uint8x16) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u16'.
// Requires NEON.
func VreinterpretqF64U16(a arm.Uint16x8) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u32'.
// Requires NEON.
func VreinterpretqF64U32(a arm.Uint32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretqF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u64'.
// Requires NEON.
func VreinterpretqF64U64(a arm.Uint64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VreinterpretS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_f64'.
// Requires NEON.
func VreinterpretS64F64(a arm.Float64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_s8'.
// Requires NEON.
func VreinterpretS64S8(a arm.Int8x8) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_s16'.
// Requires NEON.
func VreinterpretS64S16(a arm.Int16x4) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_s32'.
// Requires NEON.
func VreinterpretS64S32(a arm.Int32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_f32'.
// Requires NEON.
func VreinterpretS64F32(a arm.Float32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u8'.
// Requires NEON.
func VreinterpretS64U8(a arm.Uint8x8) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u16'.
// Requires NEON.
func VreinterpretS64U16(a arm.Uint16x4) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u32'.
// Requires NEON.
func VreinterpretS64U32(a arm.Uint32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u64'.
// Requires NEON.
func VreinterpretS64U64(a arm.Uint64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_p8'.
// Requires NEON.
func VreinterpretS64P8(a arm.Poly8x8) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretS64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_p16'.
// Requires NEON.
func VreinterpretS64P16(a arm.Poly16x4) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VreinterpretqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_f64'.
// Requires NEON.
func VreinterpretqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_s8'.
// Requires NEON.
func VreinterpretqS64S8(a arm.Int8x16) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_s16'.
// Requires NEON.
func VreinterpretqS64S16(a arm.Int16x8) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_s32'.
// Requires NEON.
func VreinterpretqS64S32(a arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_f32'.
// Requires NEON.
func VreinterpretqS64F32(a arm.Float32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u8'.
// Requires NEON.
func VreinterpretqS64U8(a arm.Uint8x16) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u16'.
// Requires NEON.
func VreinterpretqS64U16(a arm.Uint16x8) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u32'.
// Requires NEON.
func VreinterpretqS64U32(a arm.Uint32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u64'.
// Requires NEON.
func VreinterpretqS64U64(a arm.Uint64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_p8'.
// Requires NEON.
func VreinterpretqS64P8(a arm.Poly8x16) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretqS64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_p16'.
// Requires NEON.
func VreinterpretqS64P16(a arm.Poly16x8) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VreinterpretU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_f64'.
// Requires NEON.
func VreinterpretU64F64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s8'.
// Requires NEON.
func VreinterpretU64S8(a arm.Int8x8) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s16'.
// Requires NEON.
func VreinterpretU64S16(a arm.Int16x4) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s32'.
// Requires NEON.
func VreinterpretU64S32(a arm.Int32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s64'.
// Requires NEON.
func VreinterpretU64S64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_f32'.
// Requires NEON.
func VreinterpretU64F32(a arm.Float32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_u8'.
// Requires NEON.
func VreinterpretU64U8(a arm.Uint8x8) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_u16'.
// Requires NEON.
func VreinterpretU64U16(a arm.Uint16x4) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_u32'.
// Requires NEON.
func VreinterpretU64U32(a arm.Uint32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_p8'.
// Requires NEON.
func VreinterpretU64P8(a arm.Poly8x8) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretU64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_p16'.
// Requires NEON.
func VreinterpretU64P16(a arm.Poly16x4) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VreinterpretqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_f64'.
// Requires NEON.
func VreinterpretqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s8'.
// Requires NEON.
func VreinterpretqU64S8(a arm.Int8x16) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s16'.
// Requires NEON.
func VreinterpretqU64S16(a arm.Int16x8) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s32'.
// Requires NEON.
func VreinterpretqU64S32(a arm.Int32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s64'.
// Requires NEON.
func VreinterpretqU64S64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_f32'.
// Requires NEON.
func VreinterpretqU64F32(a arm.Float32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_u8'.
// Requires NEON.
func VreinterpretqU64U8(a arm.Uint8x16) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_u16'.
// Requires NEON.
func VreinterpretqU64U16(a arm.Uint16x8) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_u32'.
// Requires NEON.
func VreinterpretqU64U32(a arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_p8'.
// Requires NEON.
func VreinterpretqU64P8(a arm.Poly8x16) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretqU64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_p16'.
// Requires NEON.
func VreinterpretqU64P16(a arm.Poly16x8) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VreinterpretS8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_f64'.
// Requires NEON.
func VreinterpretS8F64(a arm.Float64x1) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_s16'.
// Requires NEON.
func VreinterpretS8S16(a arm.Int16x4) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_s32'.
// Requires NEON.
func VreinterpretS8S32(a arm.Int32x2) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_s64'.
// Requires NEON.
func VreinterpretS8S64(a arm.Int64x1) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_f32'.
// Requires NEON.
func VreinterpretS8F32(a arm.Float32x2) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u8'.
// Requires NEON.
func VreinterpretS8U8(a arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u16'.
// Requires NEON.
func VreinterpretS8U16(a arm.Uint16x4) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u32'.
// Requires NEON.
func VreinterpretS8U32(a arm.Uint32x2) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u64'.
// Requires NEON.
func VreinterpretS8U64(a arm.Uint64x1) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_p8'.
// Requires NEON.
func VreinterpretS8P8(a arm.Poly8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretS8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_p16'.
// Requires NEON.
func VreinterpretS8P16(a arm.Poly16x4) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VreinterpretqS8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_f64'.
// Requires NEON.
func VreinterpretqS8F64(a arm.Float64x2) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_s16'.
// Requires NEON.
func VreinterpretqS8S16(a arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_s32'.
// Requires NEON.
func VreinterpretqS8S32(a arm.Int32x4) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_s64'.
// Requires NEON.
func VreinterpretqS8S64(a arm.Int64x2) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_f32'.
// Requires NEON.
func VreinterpretqS8F32(a arm.Float32x4) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u8'.
// Requires NEON.
func VreinterpretqS8U8(a arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u16'.
// Requires NEON.
func VreinterpretqS8U16(a arm.Uint16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u32'.
// Requires NEON.
func VreinterpretqS8U32(a arm.Uint32x4) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u64'.
// Requires NEON.
func VreinterpretqS8U64(a arm.Uint64x2) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_p8'.
// Requires NEON.
func VreinterpretqS8P8(a arm.Poly8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretqS8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_p16'.
// Requires NEON.
func VreinterpretqS8P16(a arm.Poly16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VreinterpretS16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_f64'.
// Requires NEON.
func VreinterpretS16F64(a arm.Float64x1) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_s8'.
// Requires NEON.
func VreinterpretS16S8(a arm.Int8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_s32'.
// Requires NEON.
func VreinterpretS16S32(a arm.Int32x2) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_s64'.
// Requires NEON.
func VreinterpretS16S64(a arm.Int64x1) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_f32'.
// Requires NEON.
func VreinterpretS16F32(a arm.Float32x2) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u8'.
// Requires NEON.
func VreinterpretS16U8(a arm.Uint8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u16'.
// Requires NEON.
func VreinterpretS16U16(a arm.Uint16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u32'.
// Requires NEON.
func VreinterpretS16U32(a arm.Uint32x2) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u64'.
// Requires NEON.
func VreinterpretS16U64(a arm.Uint64x1) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_p8'.
// Requires NEON.
func VreinterpretS16P8(a arm.Poly8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretS16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_p16'.
// Requires NEON.
func VreinterpretS16P16(a arm.Poly16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VreinterpretqS16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_f64'.
// Requires NEON.
func VreinterpretqS16F64(a arm.Float64x2) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_s8'.
// Requires NEON.
func VreinterpretqS16S8(a arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_s32'.
// Requires NEON.
func VreinterpretqS16S32(a arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_s64'.
// Requires NEON.
func VreinterpretqS16S64(a arm.Int64x2) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_f32'.
// Requires NEON.
func VreinterpretqS16F32(a arm.Float32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u8'.
// Requires NEON.
func VreinterpretqS16U8(a arm.Uint8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u16'.
// Requires NEON.
func VreinterpretqS16U16(a arm.Uint16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u32'.
// Requires NEON.
func VreinterpretqS16U32(a arm.Uint32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u64'.
// Requires NEON.
func VreinterpretqS16U64(a arm.Uint64x2) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_p8'.
// Requires NEON.
func VreinterpretqS16P8(a arm.Poly8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretqS16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_p16'.
// Requires NEON.
func VreinterpretqS16P16(a arm.Poly16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VreinterpretS32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_f64'.
// Requires NEON.
func VreinterpretS32F64(a arm.Float64x1) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_s8'.
// Requires NEON.
func VreinterpretS32S8(a arm.Int8x8) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_s16'.
// Requires NEON.
func VreinterpretS32S16(a arm.Int16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_s64'.
// Requires NEON.
func VreinterpretS32S64(a arm.Int64x1) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_f32'.
// Requires NEON.
func VreinterpretS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u8'.
// Requires NEON.
func VreinterpretS32U8(a arm.Uint8x8) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u16'.
// Requires NEON.
func VreinterpretS32U16(a arm.Uint16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u32'.
// Requires NEON.
func VreinterpretS32U32(a arm.Uint32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u64'.
// Requires NEON.
func VreinterpretS32U64(a arm.Uint64x1) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_p8'.
// Requires NEON.
func VreinterpretS32P8(a arm.Poly8x8) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretS32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_p16'.
// Requires NEON.
func VreinterpretS32P16(a arm.Poly16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VreinterpretqS32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_f64'.
// Requires NEON.
func VreinterpretqS32F64(a arm.Float64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_s8'.
// Requires NEON.
func VreinterpretqS32S8(a arm.Int8x16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_s16'.
// Requires NEON.
func VreinterpretqS32S16(a arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_s64'.
// Requires NEON.
func VreinterpretqS32S64(a arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_f32'.
// Requires NEON.
func VreinterpretqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u8'.
// Requires NEON.
func VreinterpretqS32U8(a arm.Uint8x16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u16'.
// Requires NEON.
func VreinterpretqS32U16(a arm.Uint16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u32'.
// Requires NEON.
func VreinterpretqS32U32(a arm.Uint32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u64'.
// Requires NEON.
func VreinterpretqS32U64(a arm.Uint64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_p8'.
// Requires NEON.
func VreinterpretqS32P8(a arm.Poly8x16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretqS32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_p16'.
// Requires NEON.
func VreinterpretqS32P16(a arm.Poly16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VreinterpretU8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_f64'.
// Requires NEON.
func VreinterpretU8F64(a arm.Float64x1) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s8'.
// Requires NEON.
func VreinterpretU8S8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s16'.
// Requires NEON.
func VreinterpretU8S16(a arm.Int16x4) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s32'.
// Requires NEON.
func VreinterpretU8S32(a arm.Int32x2) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s64'.
// Requires NEON.
func VreinterpretU8S64(a arm.Int64x1) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_f32'.
// Requires NEON.
func VreinterpretU8F32(a arm.Float32x2) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_u16'.
// Requires NEON.
func VreinterpretU8U16(a arm.Uint16x4) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_u32'.
// Requires NEON.
func VreinterpretU8U32(a arm.Uint32x2) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_u64'.
// Requires NEON.
func VreinterpretU8U64(a arm.Uint64x1) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_p8'.
// Requires NEON.
func VreinterpretU8P8(a arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretU8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_p16'.
// Requires NEON.
func VreinterpretU8P16(a arm.Poly16x4) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VreinterpretqU8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_f64'.
// Requires NEON.
func VreinterpretqU8F64(a arm.Float64x2) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s8'.
// Requires NEON.
func VreinterpretqU8S8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s16'.
// Requires NEON.
func VreinterpretqU8S16(a arm.Int16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s32'.
// Requires NEON.
func VreinterpretqU8S32(a arm.Int32x4) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s64'.
// Requires NEON.
func VreinterpretqU8S64(a arm.Int64x2) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_f32'.
// Requires NEON.
func VreinterpretqU8F32(a arm.Float32x4) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_u16'.
// Requires NEON.
func VreinterpretqU8U16(a arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_u32'.
// Requires NEON.
func VreinterpretqU8U32(a arm.Uint32x4) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_u64'.
// Requires NEON.
func VreinterpretqU8U64(a arm.Uint64x2) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_p8'.
// Requires NEON.
func VreinterpretqU8P8(a arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretqU8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_p16'.
// Requires NEON.
func VreinterpretqU8P16(a arm.Poly16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VreinterpretU16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_f64'.
// Requires NEON.
func VreinterpretU16F64(a arm.Float64x1) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s8'.
// Requires NEON.
func VreinterpretU16S8(a arm.Int8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s16'.
// Requires NEON.
func VreinterpretU16S16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s32'.
// Requires NEON.
func VreinterpretU16S32(a arm.Int32x2) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s64'.
// Requires NEON.
func VreinterpretU16S64(a arm.Int64x1) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_f32'.
// Requires NEON.
func VreinterpretU16F32(a arm.Float32x2) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_u8'.
// Requires NEON.
func VreinterpretU16U8(a arm.Uint8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_u32'.
// Requires NEON.
func VreinterpretU16U32(a arm.Uint32x2) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_u64'.
// Requires NEON.
func VreinterpretU16U64(a arm.Uint64x1) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_p8'.
// Requires NEON.
func VreinterpretU16P8(a arm.Poly8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretU16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_p16'.
// Requires NEON.
func VreinterpretU16P16(a arm.Poly16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VreinterpretqU16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_f64'.
// Requires NEON.
func VreinterpretqU16F64(a arm.Float64x2) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s8'.
// Requires NEON.
func VreinterpretqU16S8(a arm.Int8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s16'.
// Requires NEON.
func VreinterpretqU16S16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s32'.
// Requires NEON.
func VreinterpretqU16S32(a arm.Int32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s64'.
// Requires NEON.
func VreinterpretqU16S64(a arm.Int64x2) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_f32'.
// Requires NEON.
func VreinterpretqU16F32(a arm.Float32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_u8'.
// Requires NEON.
func VreinterpretqU16U8(a arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_u32'.
// Requires NEON.
func VreinterpretqU16U32(a arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_u64'.
// Requires NEON.
func VreinterpretqU16U64(a arm.Uint64x2) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_p8'.
// Requires NEON.
func VreinterpretqU16P8(a arm.Poly8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretqU16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_p16'.
// Requires NEON.
func VreinterpretqU16P16(a arm.Poly16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VreinterpretU32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_f64'.
// Requires NEON.
func VreinterpretU32F64(a arm.Float64x1) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s8'.
// Requires NEON.
func VreinterpretU32S8(a arm.Int8x8) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s16'.
// Requires NEON.
func VreinterpretU32S16(a arm.Int16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s32'.
// Requires NEON.
func VreinterpretU32S32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s64'.
// Requires NEON.
func VreinterpretU32S64(a arm.Int64x1) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_f32'.
// Requires NEON.
func VreinterpretU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_u8'.
// Requires NEON.
func VreinterpretU32U8(a arm.Uint8x8) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_u16'.
// Requires NEON.
func VreinterpretU32U16(a arm.Uint16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_u64'.
// Requires NEON.
func VreinterpretU32U64(a arm.Uint64x1) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_p8'.
// Requires NEON.
func VreinterpretU32P8(a arm.Poly8x8) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretU32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_p16'.
// Requires NEON.
func VreinterpretU32P16(a arm.Poly16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VreinterpretqU32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_f64'.
// Requires NEON.
func VreinterpretqU32F64(a arm.Float64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s8'.
// Requires NEON.
func VreinterpretqU32S8(a arm.Int8x16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s16'.
// Requires NEON.
func VreinterpretqU32S16(a arm.Int16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s32'.
// Requires NEON.
func VreinterpretqU32S32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s64'.
// Requires NEON.
func VreinterpretqU32S64(a arm.Int64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_f32'.
// Requires NEON.
func VreinterpretqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_u8'.
// Requires NEON.
func VreinterpretqU32U8(a arm.Uint8x16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_u16'.
// Requires NEON.
func VreinterpretqU32U16(a arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_u64'.
// Requires NEON.
func VreinterpretqU32U64(a arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_p8'.
// Requires NEON.
func VreinterpretqU32P8(a arm.Poly8x16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VreinterpretqU32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_p16'.
// Requires NEON.
func VreinterpretqU32P16(a arm.Poly16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsetLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneF32(elem float32, vec arm.Float32x2, index int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VsetLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneF64(elem float64, vec arm.Float64x1, index int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VsetLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneP8(elem arm.Poly8, vec arm.Poly8x8, index int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VsetLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneP16(elem arm.Poly16, vec arm.Poly16x4, index int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VsetLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneS8(elem int8, vec arm.Int8x8, index int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VsetLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneS16(elem int16, vec arm.Int16x4, index int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VsetLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneS32(elem int32, vec arm.Int32x2, index int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VsetLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneS64(elem int64, vec arm.Int64x1, index int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VsetLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneU8(elem uint8, vec arm.Uint8x8, index int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsetLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneU16(elem uint16, vec arm.Uint16x4, index int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsetLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneU32(elem uint32, vec arm.Uint32x2, index int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VsetLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetLaneU64(elem uint64, vec arm.Uint64x1, index int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VsetqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneF32(elem float32, vec arm.Float32x4, index int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VsetqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneF64(elem float64, vec arm.Float64x2, index int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VsetqLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneP8(elem arm.Poly8, vec arm.Poly8x16, index int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VsetqLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneP16(elem arm.Poly16, vec arm.Poly16x8, index int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VsetqLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneS8(elem int8, vec arm.Int8x16, index int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VsetqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneS16(elem int16, vec arm.Int16x8, index int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsetqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneS32(elem int32, vec arm.Int32x4, index int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsetqLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneS64(elem int64, vec arm.Int64x2, index int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsetqLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneU8(elem uint8, vec arm.Uint8x16, index int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsetqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneU16(elem uint16, vec arm.Uint16x8, index int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsetqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneU32(elem uint32, vec arm.Uint32x4, index int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsetqLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsetqLaneU64(elem uint64, vec arm.Uint64x2, index int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VgetLowF32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_f32'.
// Requires NEON.
func VgetLowF32(a arm.Float32x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VgetLowF64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_f64'.
// Requires NEON.
func VgetLowF64(a arm.Float64x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VgetLowP8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_p8'.
// Requires NEON.
func VgetLowP8(a arm.Poly8x16) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VgetLowP16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_p16'.
// Requires NEON.
func VgetLowP16(a arm.Poly16x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VgetLowS8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s8'.
// Requires NEON.
func VgetLowS8(a arm.Int8x16) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VgetLowS16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s16'.
// Requires NEON.
func VgetLowS16(a arm.Int16x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VgetLowS32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s32'.
// Requires NEON.
func VgetLowS32(a arm.Int32x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VgetLowS64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s64'.
// Requires NEON.
func VgetLowS64(a arm.Int64x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VgetLowU8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u8'.
// Requires NEON.
func VgetLowU8(a arm.Uint8x16) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VgetLowU16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u16'.
// Requires NEON.
func VgetLowU16(a arm.Uint16x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VgetLowU32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u32'.
// Requires NEON.
func VgetLowU32(a arm.Uint32x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VgetLowU64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u64'.
// Requires NEON.
func VgetLowU64(a arm.Uint64x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VgetHighF32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_f32'.
// Requires NEON.
func VgetHighF32(a arm.Float32x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VgetHighF64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_f64'.
// Requires NEON.
func VgetHighF64(a arm.Float64x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VgetHighP8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_p8'.
// Requires NEON.
func VgetHighP8(a arm.Poly8x16) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VgetHighP16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_p16'.
// Requires NEON.
func VgetHighP16(a arm.Poly16x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VgetHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s8'.
// Requires NEON.
func VgetHighS8(a arm.Int8x16) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VgetHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s16'.
// Requires NEON.
func VgetHighS16(a arm.Int16x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VgetHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s32'.
// Requires NEON.
func VgetHighS32(a arm.Int32x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VgetHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s64'.
// Requires NEON.
func VgetHighS64(a arm.Int64x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VgetHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u8'.
// Requires NEON.
func VgetHighU8(a arm.Uint8x16) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VgetHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u16'.
// Requires NEON.
func VgetHighU16(a arm.Uint16x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VgetHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u32'.
// Requires NEON.
func VgetHighU32(a arm.Uint32x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VgetHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u64'.
// Requires NEON.
func VgetHighU64(a arm.Uint64x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcombineS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s8'.
// Requires NEON.
func VcombineS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VcombineS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s16'.
// Requires NEON.
func VcombineS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VcombineS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s32'.
// Requires NEON.
func VcombineS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcombineS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s64'.
// Requires NEON.
func VcombineS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VcombineF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_f32'.
// Requires NEON.
func VcombineF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VcombineU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u8'.
// Requires NEON.
func VcombineU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcombineU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u16'.
// Requires NEON.
func VcombineU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcombineU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u32'.
// Requires NEON.
func VcombineU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcombineU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u64'.
// Requires NEON.
func VcombineU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcombineF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_f64'.
// Requires NEON.
func VcombineF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VcombineP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_p8'.
// Requires NEON.
func VcombineP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VcombineP16: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_p16'.
// Requires NEON.
func VcombineP16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VabaS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_s8'.
// Requires NEON.
func VabaS8(a arm.Int8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VabaS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_s16'.
// Requires NEON.
func VabaS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VabaS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_s32'.
// Requires NEON.
func VabaS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VabaU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_u8'.
// Requires NEON.
func VabaU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VabaU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_u16'.
// Requires NEON.
func VabaU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VabaU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_u32'.
// Requires NEON.
func VabaU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VabalHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_s8'.
// Requires NEON.
func VabalHighS8(a arm.Int16x8, b arm.Int8x16, c arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabalHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_s16'.
// Requires NEON.
func VabalHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabalHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_s32'.
// Requires NEON.
func VabalHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VabalHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_u8'.
// Requires NEON.
func VabalHighU8(a arm.Uint16x8, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VabalHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_u16'.
// Requires NEON.
func VabalHighU16(a arm.Uint32x4, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VabalHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_u32'.
// Requires NEON.
func VabalHighU32(a arm.Uint64x2, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VabalS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_s8'.
// Requires NEON.
func VabalS8(a arm.Int16x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_s16'.
// Requires NEON.
func VabalS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_s32'.
// Requires NEON.
func VabalS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VabalU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_u8'.
// Requires NEON.
func VabalU8(a arm.Uint16x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VabalU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_u16'.
// Requires NEON.
func VabalU16(a arm.Uint32x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VabalU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_u32'.
// Requires NEON.
func VabalU32(a arm.Uint64x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VabaqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_s8'.
// Requires NEON.
func VabaqS8(a arm.Int8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VabaqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_s16'.
// Requires NEON.
func VabaqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabaqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_s32'.
// Requires NEON.
func VabaqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabaqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_u8'.
// Requires NEON.
func VabaqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VabaqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_u16'.
// Requires NEON.
func VabaqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VabaqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_u32'.
// Requires NEON.
func VabaqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VabdF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_f32'.
// Requires NEON.
func VabdF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VabdS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_s8'.
// Requires NEON.
func VabdS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VabdS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_s16'.
// Requires NEON.
func VabdS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VabdS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_s32'.
// Requires NEON.
func VabdS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VabdU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_u8'.
// Requires NEON.
func VabdU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VabdU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_u16'.
// Requires NEON.
func VabdU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VabdU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_u32'.
// Requires NEON.
func VabdU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VabddF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabdd_f64'.
// Requires NEON.
func VabddF64(a float64, b float64) float64 {
	return 0
}

// VabdlHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_s8'.
// Requires NEON.
func VabdlHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabdlHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_s16'.
// Requires NEON.
func VabdlHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabdlHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_s32'.
// Requires NEON.
func VabdlHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VabdlHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_u8'.
// Requires NEON.
func VabdlHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VabdlHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_u16'.
// Requires NEON.
func VabdlHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VabdlHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_u32'.
// Requires NEON.
func VabdlHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VabdlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_s8'.
// Requires NEON.
func VabdlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabdlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_s16'.
// Requires NEON.
func VabdlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabdlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_s32'.
// Requires NEON.
func VabdlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VabdlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_u8'.
// Requires NEON.
func VabdlU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VabdlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_u16'.
// Requires NEON.
func VabdlU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VabdlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_u32'.
// Requires NEON.
func VabdlU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VabdqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_f32'.
// Requires NEON.
func VabdqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VabdqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_f64'.
// Requires NEON.
func VabdqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VabdqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_s8'.
// Requires NEON.
func VabdqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VabdqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_s16'.
// Requires NEON.
func VabdqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabdqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_s32'.
// Requires NEON.
func VabdqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabdqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_u8'.
// Requires NEON.
func VabdqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VabdqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_u16'.
// Requires NEON.
func VabdqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VabdqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_u32'.
// Requires NEON.
func VabdqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VabdsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabds_f32'.
// Requires NEON.
func VabdsF32(a float32, b float32) float32 {
	return 0
}

// VaddlvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_s8'.
// Requires NEON.
func VaddlvS8(a arm.Int8x8) int16 {
	return 0
}

// VaddlvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_s16'.
// Requires NEON.
func VaddlvS16(a arm.Int16x4) int32 {
	return 0
}

// VaddlvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_u8'.
// Requires NEON.
func VaddlvU8(a arm.Uint8x8) uint16 {
	return 0
}

// VaddlvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_u16'.
// Requires NEON.
func VaddlvU16(a arm.Uint16x4) uint32 {
	return 0
}

// VaddlvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_s8'.
// Requires NEON.
func VaddlvqS8(a arm.Int8x16) int16 {
	return 0
}

// VaddlvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_s16'.
// Requires NEON.
func VaddlvqS16(a arm.Int16x8) int32 {
	return 0
}

// VaddlvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_s32'.
// Requires NEON.
func VaddlvqS32(a arm.Int32x4) int64 {
	return 0
}

// VaddlvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_u8'.
// Requires NEON.
func VaddlvqU8(a arm.Uint8x16) uint16 {
	return 0
}

// VaddlvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_u16'.
// Requires NEON.
func VaddlvqU16(a arm.Uint16x8) uint32 {
	return 0
}

// VaddlvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_u32'.
// Requires NEON.
func VaddlvqU32(a arm.Uint32x4) uint64 {
	return 0
}

// VcvtxF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtx_f32_f64'.
// Requires NEON.
func VcvtxF32F64(a arm.Float64x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VcvtxHighF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtx_high_f32_f64'.
// Requires NEON.
func VcvtxHighF32F64(a arm.Float32x2, b arm.Float64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VcvtxdF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtxd_f32_f64'.
// Requires NEON.
func VcvtxdF32F64(a float64) float32 {
	return 0
}

// VmlaNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_f32'.
// Requires NEON.
func VmlaNF32(a arm.Float32x2, b arm.Float32x2, c float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlaNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_s16'.
// Requires NEON.
func VmlaNS16(a arm.Int16x4, b arm.Int16x4, c int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlaNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_s32'.
// Requires NEON.
func VmlaNS32(a arm.Int32x2, b arm.Int32x2, c int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlaNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_u16'.
// Requires NEON.
func VmlaNU16(a arm.Uint16x4, b arm.Uint16x4, c uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlaNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_u32'.
// Requires NEON.
func VmlaNU32(a arm.Uint32x2, b arm.Uint32x2, c uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlaS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_s8'.
// Requires NEON.
func VmlaS8(a arm.Int8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmlaS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_s16'.
// Requires NEON.
func VmlaS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlaS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_s32'.
// Requires NEON.
func VmlaS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlaU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_u8'.
// Requires NEON.
func VmlaU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmlaU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_u16'.
// Requires NEON.
func VmlaU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlaU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_u32'.
// Requires NEON.
func VmlaU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlalHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_s16'.
// Requires NEON.
func VmlalHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlalHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_s32'.
// Requires NEON.
func VmlalHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlalHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_u16'.
// Requires NEON.
func VmlalHighNU16(a arm.Uint32x4, b arm.Uint16x8, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlalHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_u32'.
// Requires NEON.
func VmlalHighNU32(a arm.Uint64x2, b arm.Uint32x4, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlalHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_s8'.
// Requires NEON.
func VmlalHighS8(a arm.Int16x8, b arm.Int8x16, c arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlalHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_s16'.
// Requires NEON.
func VmlalHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlalHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_s32'.
// Requires NEON.
func VmlalHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlalHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_u8'.
// Requires NEON.
func VmlalHighU8(a arm.Uint16x8, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlalHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_u16'.
// Requires NEON.
func VmlalHighU16(a arm.Uint32x4, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlalHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_u32'.
// Requires NEON.
func VmlalHighU32(a arm.Uint64x2, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlalNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_s16'.
// Requires NEON.
func VmlalNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlalNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_s32'.
// Requires NEON.
func VmlalNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlalNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_u16'.
// Requires NEON.
func VmlalNU16(a arm.Uint32x4, b arm.Uint16x4, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlalNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_u32'.
// Requires NEON.
func VmlalNU32(a arm.Uint64x2, b arm.Uint32x2, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlalS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_s8'.
// Requires NEON.
func VmlalS8(a arm.Int16x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_s16'.
// Requires NEON.
func VmlalS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_s32'.
// Requires NEON.
func VmlalS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlalU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_u8'.
// Requires NEON.
func VmlalU8(a arm.Uint16x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlalU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_u16'.
// Requires NEON.
func VmlalU16(a arm.Uint32x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlalU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_u32'.
// Requires NEON.
func VmlalU32(a arm.Uint64x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlaqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_f32'.
// Requires NEON.
func VmlaqNF32(a arm.Float32x4, b arm.Float32x4, c float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlaqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_s16'.
// Requires NEON.
func VmlaqNS16(a arm.Int16x8, b arm.Int16x8, c int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlaqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_s32'.
// Requires NEON.
func VmlaqNS32(a arm.Int32x4, b arm.Int32x4, c int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlaqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_u16'.
// Requires NEON.
func VmlaqNU16(a arm.Uint16x8, b arm.Uint16x8, c uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlaqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_u32'.
// Requires NEON.
func VmlaqNU32(a arm.Uint32x4, b arm.Uint32x4, c uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlaqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_s8'.
// Requires NEON.
func VmlaqS8(a arm.Int8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmlaqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_s16'.
// Requires NEON.
func VmlaqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlaqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_s32'.
// Requires NEON.
func VmlaqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlaqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_u8'.
// Requires NEON.
func VmlaqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmlaqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_u16'.
// Requires NEON.
func VmlaqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlaqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_u32'.
// Requires NEON.
func VmlaqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlsNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_f32'.
// Requires NEON.
func VmlsNF32(a arm.Float32x2, b arm.Float32x2, c float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlsNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_s16'.
// Requires NEON.
func VmlsNS16(a arm.Int16x4, b arm.Int16x4, c int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_s32'.
// Requires NEON.
func VmlsNS32(a arm.Int32x2, b arm.Int32x2, c int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlsNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_u16'.
// Requires NEON.
func VmlsNU16(a arm.Uint16x4, b arm.Uint16x4, c uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_u32'.
// Requires NEON.
func VmlsNU32(a arm.Uint32x2, b arm.Uint32x2, c uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_s8'.
// Requires NEON.
func VmlsS8(a arm.Int8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmlsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_s16'.
// Requires NEON.
func VmlsS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_s32'.
// Requires NEON.
func VmlsS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlsU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_u8'.
// Requires NEON.
func VmlsU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmlsU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_u16'.
// Requires NEON.
func VmlsU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_u32'.
// Requires NEON.
func VmlsU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlslHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_s16'.
// Requires NEON.
func VmlslHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlslHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_s32'.
// Requires NEON.
func VmlslHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlslHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_u16'.
// Requires NEON.
func VmlslHighNU16(a arm.Uint32x4, b arm.Uint16x8, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlslHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_u32'.
// Requires NEON.
func VmlslHighNU32(a arm.Uint64x2, b arm.Uint32x4, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlslHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_s8'.
// Requires NEON.
func VmlslHighS8(a arm.Int16x8, b arm.Int8x16, c arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlslHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_s16'.
// Requires NEON.
func VmlslHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlslHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_s32'.
// Requires NEON.
func VmlslHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlslHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_u8'.
// Requires NEON.
func VmlslHighU8(a arm.Uint16x8, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlslHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_u16'.
// Requires NEON.
func VmlslHighU16(a arm.Uint32x4, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlslHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_u32'.
// Requires NEON.
func VmlslHighU32(a arm.Uint64x2, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlslNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_s16'.
// Requires NEON.
func VmlslNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlslNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_s32'.
// Requires NEON.
func VmlslNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlslNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_u16'.
// Requires NEON.
func VmlslNU16(a arm.Uint32x4, b arm.Uint16x4, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlslNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_u32'.
// Requires NEON.
func VmlslNU32(a arm.Uint64x2, b arm.Uint32x2, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlslS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_s8'.
// Requires NEON.
func VmlslS8(a arm.Int16x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlslS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_s16'.
// Requires NEON.
func VmlslS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlslS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_s32'.
// Requires NEON.
func VmlslS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmlslU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_u8'.
// Requires NEON.
func VmlslU8(a arm.Uint16x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlslU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_u16'.
// Requires NEON.
func VmlslU16(a arm.Uint32x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlslU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_u32'.
// Requires NEON.
func VmlslU32(a arm.Uint64x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmlsqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_f32'.
// Requires NEON.
func VmlsqNF32(a arm.Float32x4, b arm.Float32x4, c float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlsqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_s16'.
// Requires NEON.
func VmlsqNS16(a arm.Int16x8, b arm.Int16x8, c int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlsqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_s32'.
// Requires NEON.
func VmlsqNS32(a arm.Int32x4, b arm.Int32x4, c int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlsqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_u16'.
// Requires NEON.
func VmlsqNU16(a arm.Uint16x8, b arm.Uint16x8, c uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlsqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_u32'.
// Requires NEON.
func VmlsqNU32(a arm.Uint32x4, b arm.Uint32x4, c uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_s8'.
// Requires NEON.
func VmlsqS8(a arm.Int8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmlsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_s16'.
// Requires NEON.
func VmlsqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_s32'.
// Requires NEON.
func VmlsqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlsqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_u8'.
// Requires NEON.
func VmlsqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmlsqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_u16'.
// Requires NEON.
func VmlsqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlsqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_u32'.
// Requires NEON.
func VmlsqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmovlHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_s8'.
// Requires NEON.
func VmovlHighS8(a arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmovlHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_s16'.
// Requires NEON.
func VmovlHighS16(a arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmovlHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_s32'.
// Requires NEON.
func VmovlHighS32(a arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmovlHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_u8'.
// Requires NEON.
func VmovlHighU8(a arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmovlHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_u16'.
// Requires NEON.
func VmovlHighU16(a arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmovlHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_u32'.
// Requires NEON.
func VmovlHighU32(a arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmovlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_s8'.
// Requires NEON.
func VmovlS8(a arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmovlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_s16'.
// Requires NEON.
func VmovlS16(a arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmovlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_s32'.
// Requires NEON.
func VmovlS32(a arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmovlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_u8'.
// Requires NEON.
func VmovlU8(a arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmovlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_u16'.
// Requires NEON.
func VmovlU16(a arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmovlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_u32'.
// Requires NEON.
func VmovlU32(a arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmovnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_s16'.
// Requires NEON.
func VmovnHighS16(a arm.Int8x8, b arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmovnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_s32'.
// Requires NEON.
func VmovnHighS32(a arm.Int16x4, b arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmovnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_s64'.
// Requires NEON.
func VmovnHighS64(a arm.Int32x2, b arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmovnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_u16'.
// Requires NEON.
func VmovnHighU16(a arm.Uint8x8, b arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmovnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_u32'.
// Requires NEON.
func VmovnHighU32(a arm.Uint16x4, b arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmovnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_u64'.
// Requires NEON.
func VmovnHighU64(a arm.Uint32x2, b arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmovnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_s16'.
// Requires NEON.
func VmovnS16(a arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmovnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_s32'.
// Requires NEON.
func VmovnS32(a arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmovnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_s64'.
// Requires NEON.
func VmovnS64(a arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmovnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_u16'.
// Requires NEON.
func VmovnU16(a arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmovnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_u32'.
// Requires NEON.
func VmovnU32(a arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmovnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_u64'.
// Requires NEON.
func VmovnU64(a arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmulNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_f32'.
// Requires NEON.
func VmulNF32(a arm.Float32x2, b float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmulNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_s16'.
// Requires NEON.
func VmulNS16(a arm.Int16x4, b int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmulNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_s32'.
// Requires NEON.
func VmulNS32(a arm.Int32x2, b int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmulNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_u16'.
// Requires NEON.
func VmulNU16(a arm.Uint16x4, b uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmulNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_u32'.
// Requires NEON.
func VmulNU32(a arm.Uint32x2, b uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmullHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_s16'.
// Requires NEON.
func VmullHighNS16(a arm.Int16x8, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmullHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_s32'.
// Requires NEON.
func VmullHighNS32(a arm.Int32x4, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmullHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_u16'.
// Requires NEON.
func VmullHighNU16(a arm.Uint16x8, b uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmullHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_u32'.
// Requires NEON.
func VmullHighNU32(a arm.Uint32x4, b uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmullHighP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_p8'.
// Requires NEON.
func VmullHighP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VmullHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_s8'.
// Requires NEON.
func VmullHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmullHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_s16'.
// Requires NEON.
func VmullHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmullHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_s32'.
// Requires NEON.
func VmullHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmullHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_u8'.
// Requires NEON.
func VmullHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmullHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_u16'.
// Requires NEON.
func VmullHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmullHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_u32'.
// Requires NEON.
func VmullHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmullNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_s16'.
// Requires NEON.
func VmullNS16(a arm.Int16x4, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmullNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_s32'.
// Requires NEON.
func VmullNS32(a arm.Int32x2, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmullNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_u16'.
// Requires NEON.
func VmullNU16(a arm.Uint16x4, b uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmullNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_u32'.
// Requires NEON.
func VmullNU32(a arm.Uint32x2, b uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmullP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_p8'.
// Requires NEON.
func VmullP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VmullS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_s8'.
// Requires NEON.
func VmullS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmullS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_s16'.
// Requires NEON.
func VmullS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmullS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_s32'.
// Requires NEON.
func VmullS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmullU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_u8'.
// Requires NEON.
func VmullU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmullU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_u16'.
// Requires NEON.
func VmullU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmullU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_u32'.
// Requires NEON.
func VmullU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmulqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_f32'.
// Requires NEON.
func VmulqNF32(a arm.Float32x4, b float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmulqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_f64'.
// Requires NEON.
func VmulqNF64(a arm.Float64x2, b float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmulqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_s16'.
// Requires NEON.
func VmulqNS16(a arm.Int16x8, b int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmulqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_s32'.
// Requires NEON.
func VmulqNS32(a arm.Int32x4, b int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmulqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_u16'.
// Requires NEON.
func VmulqNU16(a arm.Uint16x8, b uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmulqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_u32'.
// Requires NEON.
func VmulqNU32(a arm.Uint32x4, b uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmulxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulx_f32'.
// Requires NEON.
func VmulxF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmulxdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxd_f64'.
// Requires NEON.
func VmulxdF64(a float64, b float64) float64 {
	return 0
}

// VmulxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxq_f32'.
// Requires NEON.
func VmulxqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmulxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxq_f64'.
// Requires NEON.
func VmulxqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmulxsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxs_f32'.
// Requires NEON.
func VmulxsF32(a float32, b float32) float32 {
	return 0
}

// VmvnP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_p8'.
// Requires NEON.
func VmvnP8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VmvnS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_s8'.
// Requires NEON.
func VmvnS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmvnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_s16'.
// Requires NEON.
func VmvnS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmvnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_s32'.
// Requires NEON.
func VmvnS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmvnU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_u8'.
// Requires NEON.
func VmvnU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmvnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_u16'.
// Requires NEON.
func VmvnU16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmvnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_u32'.
// Requires NEON.
func VmvnU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmvnqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_p8'.
// Requires NEON.
func VmvnqP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VmvnqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_s8'.
// Requires NEON.
func VmvnqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmvnqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_s16'.
// Requires NEON.
func VmvnqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmvnqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_s32'.
// Requires NEON.
func VmvnqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmvnqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_u8'.
// Requires NEON.
func VmvnqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmvnqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_u16'.
// Requires NEON.
func VmvnqU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmvnqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_u32'.
// Requires NEON.
func VmvnqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpadalS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_s8'.
// Requires NEON.
func VpadalS8(a arm.Int16x4, b arm.Int8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VpadalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_s16'.
// Requires NEON.
func VpadalS16(a arm.Int32x2, b arm.Int16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VpadalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_s32'.
// Requires NEON.
func VpadalS32(a arm.Int64x1, b arm.Int32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VpadalU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_u8'.
// Requires NEON.
func VpadalU8(a arm.Uint16x4, b arm.Uint8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VpadalU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_u16'.
// Requires NEON.
func VpadalU16(a arm.Uint32x2, b arm.Uint16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VpadalU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_u32'.
// Requires NEON.
func VpadalU32(a arm.Uint64x1, b arm.Uint32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VpadalqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_s8'.
// Requires NEON.
func VpadalqS8(a arm.Int16x8, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VpadalqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_s16'.
// Requires NEON.
func VpadalqS16(a arm.Int32x4, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VpadalqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_s32'.
// Requires NEON.
func VpadalqS32(a arm.Int64x2, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VpadalqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_u8'.
// Requires NEON.
func VpadalqU8(a arm.Uint16x8, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VpadalqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_u16'.
// Requires NEON.
func VpadalqU16(a arm.Uint32x4, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpadalqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_u32'.
// Requires NEON.
func VpadalqU32(a arm.Uint64x2, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VpaddF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_f32'.
// Requires NEON.
func VpaddF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VpaddlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_s8'.
// Requires NEON.
func VpaddlS8(a arm.Int8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VpaddlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_s16'.
// Requires NEON.
func VpaddlS16(a arm.Int16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VpaddlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_s32'.
// Requires NEON.
func VpaddlS32(a arm.Int32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VpaddlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_u8'.
// Requires NEON.
func VpaddlU8(a arm.Uint8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VpaddlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_u16'.
// Requires NEON.
func VpaddlU16(a arm.Uint16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VpaddlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_u32'.
// Requires NEON.
func VpaddlU32(a arm.Uint32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VpaddlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_s8'.
// Requires NEON.
func VpaddlqS8(a arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VpaddlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_s16'.
// Requires NEON.
func VpaddlqS16(a arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VpaddlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_s32'.
// Requires NEON.
func VpaddlqS32(a arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VpaddlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_u8'.
// Requires NEON.
func VpaddlqU8(a arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VpaddlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_u16'.
// Requires NEON.
func VpaddlqU16(a arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpaddlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_u32'.
// Requires NEON.
func VpaddlqU32(a arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VpaddqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_f32'.
// Requires NEON.
func VpaddqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VpaddqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_f64'.
// Requires NEON.
func VpaddqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VpaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s8'.
// Requires NEON.
func VpaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VpaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s16'.
// Requires NEON.
func VpaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VpaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s32'.
// Requires NEON.
func VpaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VpaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s64'.
// Requires NEON.
func VpaddqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VpaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u8'.
// Requires NEON.
func VpaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VpaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u16'.
// Requires NEON.
func VpaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VpaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u32'.
// Requires NEON.
func VpaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u64'.
// Requires NEON.
func VpaddqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VpaddsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadds_f32'.
// Requires NEON.
func VpaddsF32(a arm.Float32x2) float32 {
	return 0
}

// VqdmulhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_n_s16'.
// Requires NEON.
func VqdmulhNS16(a arm.Int16x4, b int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqdmulhNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_n_s32'.
// Requires NEON.
func VqdmulhNS32(a arm.Int32x2, b int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqdmulhqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_n_s16'.
// Requires NEON.
func VqdmulhqNS16(a arm.Int16x8, b int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqdmulhqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_n_s32'.
// Requires NEON.
func VqdmulhqNS32(a arm.Int32x4, b int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqmovnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_s16'.
// Requires NEON.
func VqmovnHighS16(a arm.Int8x8, b arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqmovnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_s32'.
// Requires NEON.
func VqmovnHighS32(a arm.Int16x4, b arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqmovnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_s64'.
// Requires NEON.
func VqmovnHighS64(a arm.Int32x2, b arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqmovnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_u16'.
// Requires NEON.
func VqmovnHighU16(a arm.Uint8x8, b arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqmovnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_u32'.
// Requires NEON.
func VqmovnHighU32(a arm.Uint16x4, b arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqmovnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_u64'.
// Requires NEON.
func VqmovnHighU64(a arm.Uint32x2, b arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqmovunHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_high_s16'.
// Requires NEON.
func VqmovunHighS16(a arm.Uint8x8, b arm.Int16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqmovunHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_high_s32'.
// Requires NEON.
func VqmovunHighS32(a arm.Uint16x4, b arm.Int32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqmovunHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_high_s64'.
// Requires NEON.
func VqmovunHighS64(a arm.Uint32x2, b arm.Int64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqrdmulhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_n_s16'.
// Requires NEON.
func VqrdmulhNS16(a arm.Int16x4, b int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqrdmulhNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_n_s32'.
// Requires NEON.
func VqrdmulhNS32(a arm.Int32x2, b int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqrdmulhqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_n_s16'.
// Requires NEON.
func VqrdmulhqNS16(a arm.Int16x8, b int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqrdmulhqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_n_s32'.
// Requires NEON.
func VqrdmulhqNS32(a arm.Int32x4, b int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VrsqrteF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrte_f32'.
// Requires NEON.
func VrsqrteF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrsqrteF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrte_f64'.
// Requires NEON.
func VrsqrteF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrsqrteU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrte_u32'.
// Requires NEON.
func VrsqrteU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrsqrtedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrted_f64'.
// Requires NEON.
func VrsqrtedF64(a float64) float64 {
	return 0
}

// VrsqrteqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrteq_f32'.
// Requires NEON.
func VrsqrteqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrsqrteqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrteq_f64'.
// Requires NEON.
func VrsqrteqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrsqrteqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrteq_u32'.
// Requires NEON.
func VrsqrteqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrsqrtesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtes_f32'.
// Requires NEON.
func VrsqrtesF32(a float32) float32 {
	return 0
}

// VrsqrtsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrts_f32'.
// Requires NEON.
func VrsqrtsF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrsqrtsdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtsd_f64'.
// Requires NEON.
func VrsqrtsdF64(a float64, b float64) float64 {
	return 0
}

// VrsqrtsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtsq_f32'.
// Requires NEON.
func VrsqrtsqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrsqrtsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtsq_f64'.
// Requires NEON.
func VrsqrtsqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrsqrtssF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtss_f32'.
// Requires NEON.
func VrsqrtssF32(a float32, b float32) float32 {
	return 0
}

// VtstP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_p8'.
// Requires NEON.
func VtstP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VtstP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_p16'.
// Requires NEON.
func VtstP16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VtstqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_p8'.
// Requires NEON.
func VtstqP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VtstqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_p16'.
// Requires NEON.
func VtstqP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vst2LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneF32(ptr *float32, b [2]arm.Float32x2, c int)  {

}

// Vst2LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneF64(ptr *float64, b [2]arm.Float64x1, c int)  {

}

// Vst2LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneP8(ptr *arm.Poly8, b [2]arm.Poly8x8, c int)  {

}

// Vst2LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneP16(ptr *arm.Poly16, b [2]arm.Poly16x4, c int)  {

}

// Vst2LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneS8(ptr *int8, b [2]arm.Int8x8, c int)  {

}

// Vst2LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneS16(ptr *int16, b [2]arm.Int16x4, c int)  {

}

// Vst2LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneS32(ptr *int32, b [2]arm.Int32x2, c int)  {

}

// Vst2LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneS64(ptr *int64, b [2]arm.Int64x1, c int)  {

}

// Vst2LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneU8(ptr *uint8, b [2]arm.Uint8x8, c int)  {

}

// Vst2LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneU16(ptr *uint16, b [2]arm.Uint16x4, c int)  {

}

// Vst2LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneU32(ptr *uint32, b [2]arm.Uint32x2, c int)  {

}

// Vst2LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2LaneU64(ptr *uint64, b [2]arm.Uint64x1, c int)  {

}

// Vst2qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneF32(ptr *float32, b [2]arm.Float32x4, c int)  {

}

// Vst2qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneF64(ptr *float64, b [2]arm.Float64x2, c int)  {

}

// Vst2qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneP8(ptr *arm.Poly8, b [2]arm.Poly8x16, c int)  {

}

// Vst2qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneP16(ptr *arm.Poly16, b [2]arm.Poly16x8, c int)  {

}

// Vst2qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneS8(ptr *int8, b [2]arm.Int8x16, c int)  {

}

// Vst2qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneS16(ptr *int16, b [2]arm.Int16x8, c int)  {

}

// Vst2qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneS32(ptr *int32, b [2]arm.Int32x4, c int)  {

}

// Vst2qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneS64(ptr *int64, b [2]arm.Int64x2, c int)  {

}

// Vst2qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneU8(ptr *uint8, b [2]arm.Uint8x16, c int)  {

}

// Vst2qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneU16(ptr *uint16, b [2]arm.Uint16x8, c int)  {

}

// Vst2qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneU32(ptr *uint32, b [2]arm.Uint32x4, c int)  {

}

// Vst2qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst2qLaneU64(ptr *uint64, b [2]arm.Uint64x2, c int)  {

}

// Vst3LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneF32(ptr *float32, b [3]arm.Float32x2, c int)  {

}

// Vst3LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneF64(ptr *float64, b [3]arm.Float64x1, c int)  {

}

// Vst3LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneP8(ptr *arm.Poly8, b [3]arm.Poly8x8, c int)  {

}

// Vst3LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneP16(ptr *arm.Poly16, b [3]arm.Poly16x4, c int)  {

}

// Vst3LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneS8(ptr *int8, b [3]arm.Int8x8, c int)  {

}

// Vst3LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneS16(ptr *int16, b [3]arm.Int16x4, c int)  {

}

// Vst3LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneS32(ptr *int32, b [3]arm.Int32x2, c int)  {

}

// Vst3LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneS64(ptr *int64, b [3]arm.Int64x1, c int)  {

}

// Vst3LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneU8(ptr *uint8, b [3]arm.Uint8x8, c int)  {

}

// Vst3LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneU16(ptr *uint16, b [3]arm.Uint16x4, c int)  {

}

// Vst3LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneU32(ptr *uint32, b [3]arm.Uint32x2, c int)  {

}

// Vst3LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3LaneU64(ptr *uint64, b [3]arm.Uint64x1, c int)  {

}

// Vst3qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneF32(ptr *float32, b [3]arm.Float32x4, c int)  {

}

// Vst3qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneF64(ptr *float64, b [3]arm.Float64x2, c int)  {

}

// Vst3qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneP8(ptr *arm.Poly8, b [3]arm.Poly8x16, c int)  {

}

// Vst3qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneP16(ptr *arm.Poly16, b [3]arm.Poly16x8, c int)  {

}

// Vst3qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneS8(ptr *int8, b [3]arm.Int8x16, c int)  {

}

// Vst3qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneS16(ptr *int16, b [3]arm.Int16x8, c int)  {

}

// Vst3qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneS32(ptr *int32, b [3]arm.Int32x4, c int)  {

}

// Vst3qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneS64(ptr *int64, b [3]arm.Int64x2, c int)  {

}

// Vst3qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneU8(ptr *uint8, b [3]arm.Uint8x16, c int)  {

}

// Vst3qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneU16(ptr *uint16, b [3]arm.Uint16x8, c int)  {

}

// Vst3qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneU32(ptr *uint32, b [3]arm.Uint32x4, c int)  {

}

// Vst3qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst3qLaneU64(ptr *uint64, b [3]arm.Uint64x2, c int)  {

}

// Vst4LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneF32(ptr *float32, b [4]arm.Float32x2, c int)  {

}

// Vst4LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneF64(ptr *float64, b [4]arm.Float64x1, c int)  {

}

// Vst4LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneP8(ptr *arm.Poly8, b [4]arm.Poly8x8, c int)  {

}

// Vst4LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneP16(ptr *arm.Poly16, b [4]arm.Poly16x4, c int)  {

}

// Vst4LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneS8(ptr *int8, b [4]arm.Int8x8, c int)  {

}

// Vst4LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneS16(ptr *int16, b [4]arm.Int16x4, c int)  {

}

// Vst4LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneS32(ptr *int32, b [4]arm.Int32x2, c int)  {

}

// Vst4LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneS64(ptr *int64, b [4]arm.Int64x1, c int)  {

}

// Vst4LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneU8(ptr *uint8, b [4]arm.Uint8x8, c int)  {

}

// Vst4LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneU16(ptr *uint16, b [4]arm.Uint16x4, c int)  {

}

// Vst4LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneU32(ptr *uint32, b [4]arm.Uint32x2, c int)  {

}

// Vst4LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4LaneU64(ptr *uint64, b [4]arm.Uint64x1, c int)  {

}

// Vst4qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneF32(ptr *float32, b [4]arm.Float32x4, c int)  {

}

// Vst4qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneF64(ptr *float64, b [4]arm.Float64x2, c int)  {

}

// Vst4qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneP8(ptr *arm.Poly8, b [4]arm.Poly8x16, c int)  {

}

// Vst4qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneP16(ptr *arm.Poly16, b [4]arm.Poly16x8, c int)  {

}

// Vst4qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneS8(ptr *int8, b [4]arm.Int8x16, c int)  {

}

// Vst4qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneS16(ptr *int16, b [4]arm.Int16x8, c int)  {

}

// Vst4qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneS32(ptr *int32, b [4]arm.Int32x4, c int)  {

}

// Vst4qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneS64(ptr *int64, b [4]arm.Int64x2, c int)  {

}

// Vst4qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneU8(ptr *uint8, b [4]arm.Uint8x16, c int)  {

}

// Vst4qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneU16(ptr *uint16, b [4]arm.Uint16x8, c int)  {

}

// Vst4qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneU32(ptr *uint32, b [4]arm.Uint32x4, c int)  {

}

// Vst4qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst4qLaneU64(ptr *uint64, b [4]arm.Uint64x2, c int)  {

}

// VaddlvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_s32'.
// Requires NEON.
func VaddlvS32(a arm.Int32x2) int64 {
	return 0
}

// VaddlvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_u32'.
// Requires NEON.
func VaddlvU32(a arm.Uint32x2) uint64 {
	return 0
}

// VqdmulhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhLaneqS16(a arm.Int16x4, b arm.Int16x8, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqdmulhLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhLaneqS32(a arm.Int32x2, b arm.Int32x4, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqdmulhqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhqLaneqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqdmulhqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhqLaneqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqrdmulhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhLaneqS16(a arm.Int16x4, b arm.Int16x8, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqrdmulhLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhLaneqS32(a arm.Int32x2, b arm.Int32x4, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqrdmulhqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhqLaneqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqrdmulhqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhqLaneqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vqtbl1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1_p8'.
// Requires NEON.
func Vqtbl1P8(a arm.Poly8x16, b arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbl1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1_s8'.
// Requires NEON.
func Vqtbl1S8(a arm.Int8x16, b arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbl1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1_u8'.
// Requires NEON.
func Vqtbl1U8(a arm.Uint8x16, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbl1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1q_p8'.
// Requires NEON.
func Vqtbl1qP8(a arm.Poly8x16, b arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbl1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1q_s8'.
// Requires NEON.
func Vqtbl1qS8(a arm.Int8x16, b arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbl1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1q_u8'.
// Requires NEON.
func Vqtbl1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbl2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2_s8'.
// Requires NEON.
func Vqtbl2S8(tab [2]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbl2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2_u8'.
// Requires NEON.
func Vqtbl2U8(tab [2]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbl2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2_p8'.
// Requires NEON.
func Vqtbl2P8(tab [2]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbl2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2q_s8'.
// Requires NEON.
func Vqtbl2qS8(tab [2]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbl2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2q_u8'.
// Requires NEON.
func Vqtbl2qU8(tab [2]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbl2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2q_p8'.
// Requires NEON.
func Vqtbl2qP8(tab [2]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbl3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3_s8'.
// Requires NEON.
func Vqtbl3S8(tab [3]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbl3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3_u8'.
// Requires NEON.
func Vqtbl3U8(tab [3]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbl3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3_p8'.
// Requires NEON.
func Vqtbl3P8(tab [3]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbl3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3q_s8'.
// Requires NEON.
func Vqtbl3qS8(tab [3]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbl3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3q_u8'.
// Requires NEON.
func Vqtbl3qU8(tab [3]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbl3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3q_p8'.
// Requires NEON.
func Vqtbl3qP8(tab [3]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbl4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4_s8'.
// Requires NEON.
func Vqtbl4S8(tab [4]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbl4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4_u8'.
// Requires NEON.
func Vqtbl4U8(tab [4]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbl4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4_p8'.
// Requires NEON.
func Vqtbl4P8(tab [4]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbl4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4q_s8'.
// Requires NEON.
func Vqtbl4qS8(tab [4]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbl4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4q_u8'.
// Requires NEON.
func Vqtbl4qU8(tab [4]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbl4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4q_p8'.
// Requires NEON.
func Vqtbl4qP8(tab [4]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbx1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1_s8'.
// Requires NEON.
func Vqtbx1S8(r arm.Int8x8, tab arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbx1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1_u8'.
// Requires NEON.
func Vqtbx1U8(r arm.Uint8x8, tab arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbx1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1_p8'.
// Requires NEON.
func Vqtbx1P8(r arm.Poly8x8, tab arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbx1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1q_s8'.
// Requires NEON.
func Vqtbx1qS8(r arm.Int8x16, tab arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbx1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1q_u8'.
// Requires NEON.
func Vqtbx1qU8(r arm.Uint8x16, tab arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbx1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1q_p8'.
// Requires NEON.
func Vqtbx1qP8(r arm.Poly8x16, tab arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbx2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2_s8'.
// Requires NEON.
func Vqtbx2S8(r arm.Int8x8, tab [2]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbx2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2_u8'.
// Requires NEON.
func Vqtbx2U8(r arm.Uint8x8, tab [2]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbx2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2_p8'.
// Requires NEON.
func Vqtbx2P8(r arm.Poly8x8, tab [2]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbx2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2q_s8'.
// Requires NEON.
func Vqtbx2qS8(r arm.Int8x16, tab [2]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbx2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2q_u8'.
// Requires NEON.
func Vqtbx2qU8(r arm.Uint8x16, tab [2]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbx2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2q_p8'.
// Requires NEON.
func Vqtbx2qP8(r arm.Poly8x16, tab [2]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbx3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3_s8'.
// Requires NEON.
func Vqtbx3S8(r arm.Int8x8, tab [3]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbx3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3_u8'.
// Requires NEON.
func Vqtbx3U8(r arm.Uint8x8, tab [3]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbx3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3_p8'.
// Requires NEON.
func Vqtbx3P8(r arm.Poly8x8, tab [3]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbx3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3q_s8'.
// Requires NEON.
func Vqtbx3qS8(r arm.Int8x16, tab [3]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbx3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3q_u8'.
// Requires NEON.
func Vqtbx3qU8(r arm.Uint8x16, tab [3]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbx3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3q_p8'.
// Requires NEON.
func Vqtbx3qP8(r arm.Poly8x16, tab [3]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vqtbx4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4_s8'.
// Requires NEON.
func Vqtbx4S8(r arm.Int8x8, tab [4]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vqtbx4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4_u8'.
// Requires NEON.
func Vqtbx4U8(r arm.Uint8x8, tab [4]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vqtbx4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4_p8'.
// Requires NEON.
func Vqtbx4P8(r arm.Poly8x8, tab [4]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vqtbx4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4q_s8'.
// Requires NEON.
func Vqtbx4qS8(r arm.Int8x16, tab [4]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vqtbx4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4q_u8'.
// Requires NEON.
func Vqtbx4qU8(r arm.Uint8x16, tab [4]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vqtbx4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4q_p8'.
// Requires NEON.
func Vqtbx4qP8(r arm.Poly8x16, tab [4]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vtbl1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl1_s8'.
// Requires NEON.
func Vtbl1S8(tab arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbl1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl1_u8'.
// Requires NEON.
func Vtbl1U8(tab arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbl1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl1_p8'.
// Requires NEON.
func Vtbl1P8(tab arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtbl2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl2_s8'.
// Requires NEON.
func Vtbl2S8(tab [2]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbl2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl2_u8'.
// Requires NEON.
func Vtbl2U8(tab [2]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbl2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl2_p8'.
// Requires NEON.
func Vtbl2P8(tab [2]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtbl3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl3_s8'.
// Requires NEON.
func Vtbl3S8(tab [3]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbl3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl3_u8'.
// Requires NEON.
func Vtbl3U8(tab [3]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbl3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl3_p8'.
// Requires NEON.
func Vtbl3P8(tab [3]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtbl4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl4_s8'.
// Requires NEON.
func Vtbl4S8(tab [4]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbl4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl4_u8'.
// Requires NEON.
func Vtbl4U8(tab [4]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbl4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl4_p8'.
// Requires NEON.
func Vtbl4P8(tab [4]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtbx2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx2_s8'.
// Requires NEON.
func Vtbx2S8(r arm.Int8x8, tab [2]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbx2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx2_u8'.
// Requires NEON.
func Vtbx2U8(r arm.Uint8x8, tab [2]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbx2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx2_p8'.
// Requires NEON.
func Vtbx2P8(r arm.Poly8x8, tab [2]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtbx4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx4_s8'.
// Requires NEON.
func Vtbx4S8(r arm.Int8x8, tab [4]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbx4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx4_u8'.
// Requires NEON.
func Vtbx4U8(r arm.Uint8x8, tab [4]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbx4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx4_p8'.
// Requires NEON.
func Vtbx4P8(r arm.Poly8x8, tab [4]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VabsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_f32'.
// Requires NEON.
func VabsF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VabsF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_f64'.
// Requires NEON.
func VabsF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VabsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s8'.
// Requires NEON.
func VabsS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VabsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s16'.
// Requires NEON.
func VabsS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VabsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s32'.
// Requires NEON.
func VabsS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VabsS64: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s64'.
// Requires NEON.
func VabsS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VabsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_f32'.
// Requires NEON.
func VabsqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VabsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_f64'.
// Requires NEON.
func VabsqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VabsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s8'.
// Requires NEON.
func VabsqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VabsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s16'.
// Requires NEON.
func VabsqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VabsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s32'.
// Requires NEON.
func VabsqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VabsqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s64'.
// Requires NEON.
func VabsqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddd_s64'.
// Requires NEON.
func VadddS64(a int64, b int64) int64 {
	return 0
}

// VadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddd_u64'.
// Requires NEON.
func VadddU64(a uint64, b uint64) uint64 {
	return 0
}

// VaddvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_s8'.
// Requires NEON.
func VaddvS8(a arm.Int8x8) int8 {
	return 0
}

// VaddvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_s16'.
// Requires NEON.
func VaddvS16(a arm.Int16x4) int16 {
	return 0
}

// VaddvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_s32'.
// Requires NEON.
func VaddvS32(a arm.Int32x2) int32 {
	return 0
}

// VaddvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_u8'.
// Requires NEON.
func VaddvU8(a arm.Uint8x8) uint8 {
	return 0
}

// VaddvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_u16'.
// Requires NEON.
func VaddvU16(a arm.Uint16x4) uint16 {
	return 0
}

// VaddvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_u32'.
// Requires NEON.
func VaddvU32(a arm.Uint32x2) uint32 {
	return 0
}

// VaddvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s8'.
// Requires NEON.
func VaddvqS8(a arm.Int8x16) int8 {
	return 0
}

// VaddvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s16'.
// Requires NEON.
func VaddvqS16(a arm.Int16x8) int16 {
	return 0
}

// VaddvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s32'.
// Requires NEON.
func VaddvqS32(a arm.Int32x4) int32 {
	return 0
}

// VaddvqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s64'.
// Requires NEON.
func VaddvqS64(a arm.Int64x2) int64 {
	return 0
}

// VaddvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u8'.
// Requires NEON.
func VaddvqU8(a arm.Uint8x16) uint8 {
	return 0
}

// VaddvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u16'.
// Requires NEON.
func VaddvqU16(a arm.Uint16x8) uint16 {
	return 0
}

// VaddvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u32'.
// Requires NEON.
func VaddvqU32(a arm.Uint32x4) uint32 {
	return 0
}

// VaddvqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u64'.
// Requires NEON.
func VaddvqU64(a arm.Uint64x2) uint64 {
	return 0
}

// VaddvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_f32'.
// Requires NEON.
func VaddvF32(a arm.Float32x2) float32 {
	return 0
}

// VaddvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_f32'.
// Requires NEON.
func VaddvqF32(a arm.Float32x4) float32 {
	return 0
}

// VaddvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_f64'.
// Requires NEON.
func VaddvqF64(a arm.Float64x2) float64 {
	return 0
}

// VbslF32: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_f32'.
// Requires NEON.
func VbslF32(a arm.Uint32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VbslF64: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_f64'.
// Requires NEON.
func VbslF64(a arm.Uint64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VbslP8: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_p8'.
// Requires NEON.
func VbslP8(a arm.Uint8x8, b arm.Poly8x8, c arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VbslP16: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_p16'.
// Requires NEON.
func VbslP16(a arm.Uint16x4, b arm.Poly16x4, c arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VbslS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s8'.
// Requires NEON.
func VbslS8(a arm.Uint8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VbslS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s16'.
// Requires NEON.
func VbslS16(a arm.Uint16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VbslS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s32'.
// Requires NEON.
func VbslS32(a arm.Uint32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VbslS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s64'.
// Requires NEON.
func VbslS64(a arm.Uint64x1, b arm.Int64x1, c arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VbslU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u8'.
// Requires NEON.
func VbslU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VbslU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u16'.
// Requires NEON.
func VbslU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VbslU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u32'.
// Requires NEON.
func VbslU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VbslU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u64'.
// Requires NEON.
func VbslU64(a arm.Uint64x1, b arm.Uint64x1, c arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VbslqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_f32'.
// Requires NEON.
func VbslqF32(a arm.Uint32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VbslqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_f64'.
// Requires NEON.
func VbslqF64(a arm.Uint64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VbslqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_p8'.
// Requires NEON.
func VbslqP8(a arm.Uint8x16, b arm.Poly8x16, c arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VbslqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_p16'.
// Requires NEON.
func VbslqP16(a arm.Uint16x8, b arm.Poly16x8, c arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VbslqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s8'.
// Requires NEON.
func VbslqS8(a arm.Uint8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VbslqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s16'.
// Requires NEON.
func VbslqS16(a arm.Uint16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VbslqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s32'.
// Requires NEON.
func VbslqS32(a arm.Uint32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VbslqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s64'.
// Requires NEON.
func VbslqS64(a arm.Uint64x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VbslqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u8'.
// Requires NEON.
func VbslqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VbslqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u16'.
// Requires NEON.
func VbslqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VbslqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u32'.
// Requires NEON.
func VbslqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VbslqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u64'.
// Requires NEON.
func VbslqU64(a arm.Uint64x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VaeseqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaeseq_u8'.
// Requires NEON.
func VaeseqU8(data arm.Uint8x16, key arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VaesdqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaesdq_u8'.
// Requires NEON.
func VaesdqU8(data arm.Uint8x16, key arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VaesmcqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaesmcq_u8'.
// Requires NEON.
func VaesmcqU8(data arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VaesimcqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaesimcq_u8'.
// Requires NEON.
func VaesimcqU8(data arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcageF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcage_f64'.
// Requires NEON.
func VcageF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcagesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcages_f32'.
// Requires NEON.
func VcagesF32(a float32, b float32) uint32 {
	return 0
}

// VcageF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcage_f32'.
// Requires NEON.
func VcageF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcageqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcageq_f32'.
// Requires NEON.
func VcageqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcagedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaged_f64'.
// Requires NEON.
func VcagedF64(a float64, b float64) uint64 {
	return 0
}

// VcageqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcageq_f64'.
// Requires NEON.
func VcageqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcagtsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcagts_f32'.
// Requires NEON.
func VcagtsF32(a float32, b float32) uint32 {
	return 0
}

// VcagtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcagt_f32'.
// Requires NEON.
func VcagtF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcagtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcagt_f64'.
// Requires NEON.
func VcagtF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcagtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcagtq_f32'.
// Requires NEON.
func VcagtqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcagtdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcagtd_f64'.
// Requires NEON.
func VcagtdF64(a float64, b float64) uint64 {
	return 0
}

// VcagtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcagtq_f64'.
// Requires NEON.
func VcagtqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcaleF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcale_f32'.
// Requires NEON.
func VcaleF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcaleF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcale_f64'.
// Requires NEON.
func VcaleF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcaledF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaled_f64'.
// Requires NEON.
func VcaledF64(a float64, b float64) uint64 {
	return 0
}

// VcalesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcales_f32'.
// Requires NEON.
func VcalesF32(a float32, b float32) uint32 {
	return 0
}

// VcaleqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcaleq_f32'.
// Requires NEON.
func VcaleqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcaleqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaleq_f64'.
// Requires NEON.
func VcaleqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcaltF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcalt_f32'.
// Requires NEON.
func VcaltF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcaltF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcalt_f64'.
// Requires NEON.
func VcaltF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcaltdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaltd_f64'.
// Requires NEON.
func VcaltdF64(a float64, b float64) uint64 {
	return 0
}

// VcaltqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcaltq_f32'.
// Requires NEON.
func VcaltqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcaltqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaltq_f64'.
// Requires NEON.
func VcaltqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcaltsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcalts_f32'.
// Requires NEON.
func VcaltsF32(a float32, b float32) uint32 {
	return 0
}

// VceqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_f32'.
// Requires NEON.
func VceqF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VceqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_f64'.
// Requires NEON.
func VceqF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VceqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_p8'.
// Requires NEON.
func VceqP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VceqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s8'.
// Requires NEON.
func VceqS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VceqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s16'.
// Requires NEON.
func VceqS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VceqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s32'.
// Requires NEON.
func VceqS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VceqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s64'.
// Requires NEON.
func VceqS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VceqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u8'.
// Requires NEON.
func VceqU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VceqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u16'.
// Requires NEON.
func VceqU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VceqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u32'.
// Requires NEON.
func VceqU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VceqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u64'.
// Requires NEON.
func VceqU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VceqqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_f32'.
// Requires NEON.
func VceqqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VceqqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_f64'.
// Requires NEON.
func VceqqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VceqqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_p8'.
// Requires NEON.
func VceqqP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VceqqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s8'.
// Requires NEON.
func VceqqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VceqqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s16'.
// Requires NEON.
func VceqqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VceqqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s32'.
// Requires NEON.
func VceqqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VceqqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s64'.
// Requires NEON.
func VceqqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VceqqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u8'.
// Requires NEON.
func VceqqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VceqqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u16'.
// Requires NEON.
func VceqqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VceqqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u32'.
// Requires NEON.
func VceqqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VceqqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u64'.
// Requires NEON.
func VceqqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VceqsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqs_f32'.
// Requires NEON.
func VceqsF32(a float32, b float32) uint32 {
	return 0
}

// VceqdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqd_s64'.
// Requires NEON.
func VceqdS64(a int64, b int64) uint64 {
	return 0
}

// VceqdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqd_u64'.
// Requires NEON.
func VceqdU64(a uint64, b uint64) uint64 {
	return 0
}

// VceqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqd_f64'.
// Requires NEON.
func VceqdF64(a float64, b float64) uint64 {
	return 0
}

// VceqzF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_f32'.
// Requires NEON.
func VceqzF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VceqzF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_f64'.
// Requires NEON.
func VceqzF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VceqzP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_p8'.
// Requires NEON.
func VceqzP8(a arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VceqzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s8'.
// Requires NEON.
func VceqzS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VceqzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s16'.
// Requires NEON.
func VceqzS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VceqzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s32'.
// Requires NEON.
func VceqzS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VceqzS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s64'.
// Requires NEON.
func VceqzS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VceqzU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u8'.
// Requires NEON.
func VceqzU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VceqzU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u16'.
// Requires NEON.
func VceqzU16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VceqzU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u32'.
// Requires NEON.
func VceqzU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VceqzU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u64'.
// Requires NEON.
func VceqzU64(a arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VceqzqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_f32'.
// Requires NEON.
func VceqzqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VceqzqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_f64'.
// Requires NEON.
func VceqzqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VceqzqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_p8'.
// Requires NEON.
func VceqzqP8(a arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VceqzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s8'.
// Requires NEON.
func VceqzqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VceqzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s16'.
// Requires NEON.
func VceqzqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VceqzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s32'.
// Requires NEON.
func VceqzqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VceqzqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s64'.
// Requires NEON.
func VceqzqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VceqzqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u8'.
// Requires NEON.
func VceqzqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VceqzqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u16'.
// Requires NEON.
func VceqzqU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VceqzqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u32'.
// Requires NEON.
func VceqzqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VceqzqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u64'.
// Requires NEON.
func VceqzqU64(a arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VceqzsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzs_f32'.
// Requires NEON.
func VceqzsF32(a float32) uint32 {
	return 0
}

// VceqzdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzd_s64'.
// Requires NEON.
func VceqzdS64(a int64) uint64 {
	return 0
}

// VceqzdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzd_u64'.
// Requires NEON.
func VceqzdU64(a uint64) uint64 {
	return 0
}

// VceqzdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzd_f64'.
// Requires NEON.
func VceqzdF64(a float64) uint64 {
	return 0
}

// VcgeF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_f32'.
// Requires NEON.
func VcgeF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgeF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_f64'.
// Requires NEON.
func VcgeF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgeS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s8'.
// Requires NEON.
func VcgeS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcgeS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s16'.
// Requires NEON.
func VcgeS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcgeS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s32'.
// Requires NEON.
func VcgeS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgeS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s64'.
// Requires NEON.
func VcgeS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgeU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u8'.
// Requires NEON.
func VcgeU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcgeU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u16'.
// Requires NEON.
func VcgeU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcgeU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u32'.
// Requires NEON.
func VcgeU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgeU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u64'.
// Requires NEON.
func VcgeU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgeqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_f32'.
// Requires NEON.
func VcgeqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgeqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_f64'.
// Requires NEON.
func VcgeqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgeqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s8'.
// Requires NEON.
func VcgeqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcgeqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s16'.
// Requires NEON.
func VcgeqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcgeqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s32'.
// Requires NEON.
func VcgeqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgeqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s64'.
// Requires NEON.
func VcgeqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgeqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u8'.
// Requires NEON.
func VcgeqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcgeqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u16'.
// Requires NEON.
func VcgeqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcgeqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u32'.
// Requires NEON.
func VcgeqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgeqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u64'.
// Requires NEON.
func VcgeqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcges_f32'.
// Requires NEON.
func VcgesF32(a float32, b float32) uint32 {
	return 0
}

// VcgedS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcged_s64'.
// Requires NEON.
func VcgedS64(a int64, b int64) uint64 {
	return 0
}

// VcgedU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcged_u64'.
// Requires NEON.
func VcgedU64(a uint64, b uint64) uint64 {
	return 0
}

// VcgedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcged_f64'.
// Requires NEON.
func VcgedF64(a float64, b float64) uint64 {
	return 0
}

// VcgezF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_f32'.
// Requires NEON.
func VcgezF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgezF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_f64'.
// Requires NEON.
func VcgezF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgezS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s8'.
// Requires NEON.
func VcgezS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcgezS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s16'.
// Requires NEON.
func VcgezS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcgezS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s32'.
// Requires NEON.
func VcgezS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgezS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s64'.
// Requires NEON.
func VcgezS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgezqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_f32'.
// Requires NEON.
func VcgezqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgezqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_f64'.
// Requires NEON.
func VcgezqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgezqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s8'.
// Requires NEON.
func VcgezqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcgezqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s16'.
// Requires NEON.
func VcgezqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcgezqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s32'.
// Requires NEON.
func VcgezqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgezqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s64'.
// Requires NEON.
func VcgezqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgezsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezs_f32'.
// Requires NEON.
func VcgezsF32(a float32) uint32 {
	return 0
}

// VcgezdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezd_s64'.
// Requires NEON.
func VcgezdS64(a int64) uint64 {
	return 0
}

// VcgezdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezd_f64'.
// Requires NEON.
func VcgezdF64(a float64) uint64 {
	return 0
}

// VcgtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_f32'.
// Requires NEON.
func VcgtF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_f64'.
// Requires NEON.
func VcgtF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgtS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s8'.
// Requires NEON.
func VcgtS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcgtS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s16'.
// Requires NEON.
func VcgtS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcgtS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s32'.
// Requires NEON.
func VcgtS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgtS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s64'.
// Requires NEON.
func VcgtS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgtU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u8'.
// Requires NEON.
func VcgtU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcgtU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u16'.
// Requires NEON.
func VcgtU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcgtU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u32'.
// Requires NEON.
func VcgtU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgtU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u64'.
// Requires NEON.
func VcgtU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_f32'.
// Requires NEON.
func VcgtqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_f64'.
// Requires NEON.
func VcgtqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgtqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s8'.
// Requires NEON.
func VcgtqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcgtqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s16'.
// Requires NEON.
func VcgtqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcgtqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s32'.
// Requires NEON.
func VcgtqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgtqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s64'.
// Requires NEON.
func VcgtqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgtqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u8'.
// Requires NEON.
func VcgtqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcgtqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u16'.
// Requires NEON.
func VcgtqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcgtqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u32'.
// Requires NEON.
func VcgtqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgtqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u64'.
// Requires NEON.
func VcgtqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgtsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgts_f32'.
// Requires NEON.
func VcgtsF32(a float32, b float32) uint32 {
	return 0
}

// VcgtdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtd_s64'.
// Requires NEON.
func VcgtdS64(a int64, b int64) uint64 {
	return 0
}

// VcgtdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtd_u64'.
// Requires NEON.
func VcgtdU64(a uint64, b uint64) uint64 {
	return 0
}

// VcgtdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtd_f64'.
// Requires NEON.
func VcgtdF64(a float64, b float64) uint64 {
	return 0
}

// VcgtzF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_f32'.
// Requires NEON.
func VcgtzF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgtzF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_f64'.
// Requires NEON.
func VcgtzF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgtzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s8'.
// Requires NEON.
func VcgtzS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcgtzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s16'.
// Requires NEON.
func VcgtzS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcgtzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s32'.
// Requires NEON.
func VcgtzS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcgtzS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s64'.
// Requires NEON.
func VcgtzS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcgtzqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_f32'.
// Requires NEON.
func VcgtzqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgtzqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_f64'.
// Requires NEON.
func VcgtzqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgtzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s8'.
// Requires NEON.
func VcgtzqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcgtzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s16'.
// Requires NEON.
func VcgtzqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcgtzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s32'.
// Requires NEON.
func VcgtzqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcgtzqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s64'.
// Requires NEON.
func VcgtzqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcgtzsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzs_f32'.
// Requires NEON.
func VcgtzsF32(a float32) uint32 {
	return 0
}

// VcgtzdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzd_s64'.
// Requires NEON.
func VcgtzdS64(a int64) uint64 {
	return 0
}

// VcgtzdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzd_f64'.
// Requires NEON.
func VcgtzdF64(a float64) uint64 {
	return 0
}

// VcleF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_f32'.
// Requires NEON.
func VcleF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcleF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_f64'.
// Requires NEON.
func VcleF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcleS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s8'.
// Requires NEON.
func VcleS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcleS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s16'.
// Requires NEON.
func VcleS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcleS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s32'.
// Requires NEON.
func VcleS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcleS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s64'.
// Requires NEON.
func VcleS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcleU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u8'.
// Requires NEON.
func VcleU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcleU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u16'.
// Requires NEON.
func VcleU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcleU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u32'.
// Requires NEON.
func VcleU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcleU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u64'.
// Requires NEON.
func VcleU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcleqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_f32'.
// Requires NEON.
func VcleqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcleqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_f64'.
// Requires NEON.
func VcleqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcleqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s8'.
// Requires NEON.
func VcleqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcleqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s16'.
// Requires NEON.
func VcleqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcleqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s32'.
// Requires NEON.
func VcleqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcleqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s64'.
// Requires NEON.
func VcleqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcleqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u8'.
// Requires NEON.
func VcleqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcleqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u16'.
// Requires NEON.
func VcleqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcleqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u32'.
// Requires NEON.
func VcleqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcleqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u64'.
// Requires NEON.
func VcleqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VclesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcles_f32'.
// Requires NEON.
func VclesF32(a float32, b float32) uint32 {
	return 0
}

// VcledS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcled_s64'.
// Requires NEON.
func VcledS64(a int64, b int64) uint64 {
	return 0
}

// VcledU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcled_u64'.
// Requires NEON.
func VcledU64(a uint64, b uint64) uint64 {
	return 0
}

// VcledF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcled_f64'.
// Requires NEON.
func VcledF64(a float64, b float64) uint64 {
	return 0
}

// VclezF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_f32'.
// Requires NEON.
func VclezF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VclezF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_f64'.
// Requires NEON.
func VclezF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VclezS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s8'.
// Requires NEON.
func VclezS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VclezS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s16'.
// Requires NEON.
func VclezS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VclezS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s32'.
// Requires NEON.
func VclezS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VclezS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s64'.
// Requires NEON.
func VclezS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VclezqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_f32'.
// Requires NEON.
func VclezqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VclezqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_f64'.
// Requires NEON.
func VclezqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VclezqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s8'.
// Requires NEON.
func VclezqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VclezqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s16'.
// Requires NEON.
func VclezqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VclezqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s32'.
// Requires NEON.
func VclezqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VclezqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s64'.
// Requires NEON.
func VclezqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VclezsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclezs_f32'.
// Requires NEON.
func VclezsF32(a float32) uint32 {
	return 0
}

// VclezdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezd_s64'.
// Requires NEON.
func VclezdS64(a int64) uint64 {
	return 0
}

// VclezdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezd_f64'.
// Requires NEON.
func VclezdF64(a float64) uint64 {
	return 0
}

// VcltF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_f32'.
// Requires NEON.
func VcltF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcltF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_f64'.
// Requires NEON.
func VcltF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcltS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s8'.
// Requires NEON.
func VcltS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcltS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s16'.
// Requires NEON.
func VcltS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcltS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s32'.
// Requires NEON.
func VcltS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcltS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s64'.
// Requires NEON.
func VcltS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcltU8: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u8'.
// Requires NEON.
func VcltU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcltU16: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u16'.
// Requires NEON.
func VcltU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcltU32: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u32'.
// Requires NEON.
func VcltU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcltU64: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u64'.
// Requires NEON.
func VcltU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcltqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_f32'.
// Requires NEON.
func VcltqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcltqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_f64'.
// Requires NEON.
func VcltqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcltqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s8'.
// Requires NEON.
func VcltqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcltqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s16'.
// Requires NEON.
func VcltqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcltqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s32'.
// Requires NEON.
func VcltqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcltqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s64'.
// Requires NEON.
func VcltqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcltqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u8'.
// Requires NEON.
func VcltqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcltqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u16'.
// Requires NEON.
func VcltqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcltqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u32'.
// Requires NEON.
func VcltqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcltqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u64'.
// Requires NEON.
func VcltqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcltsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclts_f32'.
// Requires NEON.
func VcltsF32(a float32, b float32) uint32 {
	return 0
}

// VcltdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltd_s64'.
// Requires NEON.
func VcltdS64(a int64, b int64) uint64 {
	return 0
}

// VcltdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltd_u64'.
// Requires NEON.
func VcltdU64(a uint64, b uint64) uint64 {
	return 0
}

// VcltdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltd_f64'.
// Requires NEON.
func VcltdF64(a float64, b float64) uint64 {
	return 0
}

// VcltzF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_f32'.
// Requires NEON.
func VcltzF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcltzF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_f64'.
// Requires NEON.
func VcltzF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcltzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s8'.
// Requires NEON.
func VcltzS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcltzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s16'.
// Requires NEON.
func VcltzS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VcltzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s32'.
// Requires NEON.
func VcltzS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcltzS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s64'.
// Requires NEON.
func VcltzS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VcltzqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_f32'.
// Requires NEON.
func VcltzqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcltzqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_f64'.
// Requires NEON.
func VcltzqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcltzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s8'.
// Requires NEON.
func VcltzqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcltzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s16'.
// Requires NEON.
func VcltzqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VcltzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s32'.
// Requires NEON.
func VcltzqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcltzqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s64'.
// Requires NEON.
func VcltzqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcltzsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzs_f32'.
// Requires NEON.
func VcltzsF32(a float32) uint32 {
	return 0
}

// VcltzdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzd_s64'.
// Requires NEON.
func VcltzdS64(a int64) uint64 {
	return 0
}

// VcltzdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzd_f64'.
// Requires NEON.
func VcltzdF64(a float64) uint64 {
	return 0
}

// VclsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcls_s8'.
// Requires NEON.
func VclsS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VclsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcls_s16'.
// Requires NEON.
func VclsS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VclsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcls_s32'.
// Requires NEON.
func VclsS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VclsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclsq_s8'.
// Requires NEON.
func VclsqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VclsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclsq_s16'.
// Requires NEON.
func VclsqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VclsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclsq_s32'.
// Requires NEON.
func VclsqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VclzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_s8'.
// Requires NEON.
func VclzS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VclzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_s16'.
// Requires NEON.
func VclzS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VclzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_s32'.
// Requires NEON.
func VclzS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VclzU8: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_u8'.
// Requires NEON.
func VclzU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VclzU16: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_u16'.
// Requires NEON.
func VclzU16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VclzU32: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_u32'.
// Requires NEON.
func VclzU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VclzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_s8'.
// Requires NEON.
func VclzqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VclzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_s16'.
// Requires NEON.
func VclzqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VclzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_s32'.
// Requires NEON.
func VclzqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VclzqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_u8'.
// Requires NEON.
func VclzqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VclzqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_u16'.
// Requires NEON.
func VclzqU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VclzqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_u32'.
// Requires NEON.
func VclzqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcntP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcnt_p8'.
// Requires NEON.
func VcntP8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VcntS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcnt_s8'.
// Requires NEON.
func VcntS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VcntU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcnt_u8'.
// Requires NEON.
func VcntU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VcntqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcntq_p8'.
// Requires NEON.
func VcntqP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VcntqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcntq_s8'.
// Requires NEON.
func VcntqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VcntqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcntq_u8'.
// Requires NEON.
func VcntqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VcvtF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f32_f64'.
// Requires NEON.
func VcvtF32F64(a arm.Float64x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VcvtHighF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_high_f32_f64'.
// Requires NEON.
func VcvtHighF32F64(a arm.Float32x2, b arm.Float64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VcvtF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f64_f32'.
// Requires NEON.
func VcvtF64F32(a arm.Float32x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VcvtHighF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_high_f64_f32'.
// Requires NEON.
func VcvtHighF64F32(a arm.Float32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VcvtdF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_f64_s64'.
// Requires NEON.
func VcvtdF64S64(a int64) float64 {
	return 0
}

// VcvtdF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_f64_u64'.
// Requires NEON.
func VcvtdF64U64(a uint64) float64 {
	return 0
}

// VcvtsF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_f32_s32'.
// Requires NEON.
func VcvtsF32S32(a int32) float32 {
	return 0
}

// VcvtsF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_f32_u32'.
// Requires NEON.
func VcvtsF32U32(a uint32) float32 {
	return 0
}

// VcvtF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f32_s32'.
// Requires NEON.
func VcvtF32S32(a arm.Int32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VcvtF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f32_u32'.
// Requires NEON.
func VcvtF32U32(a arm.Uint32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VcvtqF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f32_s32'.
// Requires NEON.
func VcvtqF32S32(a arm.Int32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VcvtqF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f32_u32'.
// Requires NEON.
func VcvtqF32U32(a arm.Uint32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VcvtqF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f64_s64'.
// Requires NEON.
func VcvtqF64S64(a arm.Int64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VcvtqF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f64_u64'.
// Requires NEON.
func VcvtqF64U64(a arm.Uint64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VcvtdS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_s64_f64'.
// Requires NEON.
func VcvtdS64F64(a float64) int64 {
	return 0
}

// VcvtdU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_u64_f64'.
// Requires NEON.
func VcvtdU64F64(a float64) uint64 {
	return 0
}

// VcvtsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_s32_f32'.
// Requires NEON.
func VcvtsS32F32(a float32) int32 {
	return 0
}

// VcvtsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_u32_f32'.
// Requires NEON.
func VcvtsU32F32(a float32) uint32 {
	return 0
}

// VcvtS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_s32_f32'.
// Requires NEON.
func VcvtS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VcvtU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_u32_f32'.
// Requires NEON.
func VcvtU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcvtqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_s32_f32'.
// Requires NEON.
func VcvtqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcvtqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_u32_f32'.
// Requires NEON.
func VcvtqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcvtqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_s64_f64'.
// Requires NEON.
func VcvtqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VcvtqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_u64_f64'.
// Requires NEON.
func VcvtqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcvtadS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtad_s64_f64'.
// Requires NEON.
func VcvtadS64F64(a float64) int64 {
	return 0
}

// VcvtadU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtad_u64_f64'.
// Requires NEON.
func VcvtadU64F64(a float64) uint64 {
	return 0
}

// VcvtasS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtas_s32_f32'.
// Requires NEON.
func VcvtasS32F32(a float32) int32 {
	return 0
}

// VcvtasU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtas_u32_f32'.
// Requires NEON.
func VcvtasU32F32(a float32) uint32 {
	return 0
}

// VcvtaS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvta_s32_f32'.
// Requires NEON.
func VcvtaS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VcvtaU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvta_u32_f32'.
// Requires NEON.
func VcvtaU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcvtaqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_s32_f32'.
// Requires NEON.
func VcvtaqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcvtaqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_u32_f32'.
// Requires NEON.
func VcvtaqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcvtaqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_s64_f64'.
// Requires NEON.
func VcvtaqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VcvtaqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_u64_f64'.
// Requires NEON.
func VcvtaqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcvtmdS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmd_s64_f64'.
// Requires NEON.
func VcvtmdS64F64(a float64) int64 {
	return 0
}

// VcvtmdU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmd_u64_f64'.
// Requires NEON.
func VcvtmdU64F64(a float64) uint64 {
	return 0
}

// VcvtmsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtms_s32_f32'.
// Requires NEON.
func VcvtmsS32F32(a float32) int32 {
	return 0
}

// VcvtmsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtms_u32_f32'.
// Requires NEON.
func VcvtmsU32F32(a float32) uint32 {
	return 0
}

// VcvtmS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtm_s32_f32'.
// Requires NEON.
func VcvtmS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VcvtmU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtm_u32_f32'.
// Requires NEON.
func VcvtmU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcvtmqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_s32_f32'.
// Requires NEON.
func VcvtmqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcvtmqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_u32_f32'.
// Requires NEON.
func VcvtmqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcvtmqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_s64_f64'.
// Requires NEON.
func VcvtmqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VcvtmqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_u64_f64'.
// Requires NEON.
func VcvtmqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcvtndS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnd_s64_f64'.
// Requires NEON.
func VcvtndS64F64(a float64) int64 {
	return 0
}

// VcvtndU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnd_u64_f64'.
// Requires NEON.
func VcvtndU64F64(a float64) uint64 {
	return 0
}

// VcvtnsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtns_s32_f32'.
// Requires NEON.
func VcvtnsS32F32(a float32) int32 {
	return 0
}

// VcvtnsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtns_u32_f32'.
// Requires NEON.
func VcvtnsU32F32(a float32) uint32 {
	return 0
}

// VcvtnS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtn_s32_f32'.
// Requires NEON.
func VcvtnS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VcvtnU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtn_u32_f32'.
// Requires NEON.
func VcvtnU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcvtnqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_s32_f32'.
// Requires NEON.
func VcvtnqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcvtnqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_u32_f32'.
// Requires NEON.
func VcvtnqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcvtnqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_s64_f64'.
// Requires NEON.
func VcvtnqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VcvtnqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_u64_f64'.
// Requires NEON.
func VcvtnqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VcvtpdS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpd_s64_f64'.
// Requires NEON.
func VcvtpdS64F64(a float64) int64 {
	return 0
}

// VcvtpdU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpd_u64_f64'.
// Requires NEON.
func VcvtpdU64F64(a float64) uint64 {
	return 0
}

// VcvtpsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtps_s32_f32'.
// Requires NEON.
func VcvtpsS32F32(a float32) int32 {
	return 0
}

// VcvtpsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtps_u32_f32'.
// Requires NEON.
func VcvtpsU32F32(a float32) uint32 {
	return 0
}

// VcvtpS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtp_s32_f32'.
// Requires NEON.
func VcvtpS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VcvtpU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtp_u32_f32'.
// Requires NEON.
func VcvtpU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VcvtpqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_s32_f32'.
// Requires NEON.
func VcvtpqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VcvtpqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_u32_f32'.
// Requires NEON.
func VcvtpqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VcvtpqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_s64_f64'.
// Requires NEON.
func VcvtpqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VcvtpqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_u64_f64'.
// Requires NEON.
func VcvtpqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VdupNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_f32'.
// Requires NEON.
func VdupNF32(a float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VdupNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_f64'.
// Requires NEON.
func VdupNF64(a float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VdupNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_p8'.
// Requires NEON.
func VdupNP8(a arm.Poly8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VdupNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_p16'.
// Requires NEON.
func VdupNP16(a arm.Poly16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VdupNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s8'.
// Requires NEON.
func VdupNS8(a int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VdupNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s16'.
// Requires NEON.
func VdupNS16(a int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VdupNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s32'.
// Requires NEON.
func VdupNS32(a int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VdupNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s64'.
// Requires NEON.
func VdupNS64(a int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VdupNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u8'.
// Requires NEON.
func VdupNU8(a uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VdupNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u16'.
// Requires NEON.
func VdupNU16(a uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VdupNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u32'.
// Requires NEON.
func VdupNU32(a uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VdupNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u64'.
// Requires NEON.
func VdupNU64(a uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VdupqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_f32'.
// Requires NEON.
func VdupqNF32(a float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VdupqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_f64'.
// Requires NEON.
func VdupqNF64(a float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VdupqNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_p8'.
// Requires NEON.
func VdupqNP8(a uint32) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VdupqNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_p16'.
// Requires NEON.
func VdupqNP16(a uint32) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VdupqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s8'.
// Requires NEON.
func VdupqNS8(a int32) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VdupqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s16'.
// Requires NEON.
func VdupqNS16(a int32) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VdupqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s32'.
// Requires NEON.
func VdupqNS32(a int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VdupqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s64'.
// Requires NEON.
func VdupqNS64(a int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VdupqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u8'.
// Requires NEON.
func VdupqNU8(a uint32) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VdupqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u16'.
// Requires NEON.
func VdupqNU16(a uint32) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VdupqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u32'.
// Requires NEON.
func VdupqNU32(a uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VdupqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u64'.
// Requires NEON.
func VdupqNU64(a uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VdupLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneF32(a arm.Float32x2, b int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VdupLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneF64(a arm.Float64x1, b int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VdupLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VdupLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VdupLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VdupLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VdupLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VdupLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VdupLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VdupLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VdupLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VdupLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VdupLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqF32(a arm.Float32x4, b int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VdupLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqF64(a arm.Float64x2, b int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VdupLaneqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqP8(a arm.Poly8x16, b int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VdupLaneqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqP16(a arm.Poly16x8, b int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VdupLaneqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqS8(a arm.Int8x16, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VdupLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqS16(a arm.Int16x8, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VdupLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqS32(a arm.Int32x4, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VdupLaneqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqS64(a arm.Int64x2, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VdupLaneqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqU8(a arm.Uint8x16, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VdupLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqU16(a arm.Uint16x8, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VdupLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqU32(a arm.Uint32x4, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VdupLaneqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupLaneqU64(a arm.Uint64x2, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VdupqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneF32(a arm.Float32x2, b int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VdupqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneF64(a arm.Float64x1, b int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VdupqLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VdupqLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VdupqLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneS8(a arm.Int8x8, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VdupqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneS16(a arm.Int16x4, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VdupqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneS32(a arm.Int32x2, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VdupqLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneS64(a arm.Int64x1, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VdupqLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneU8(a arm.Uint8x8, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VdupqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneU16(a arm.Uint16x4, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VdupqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneU32(a arm.Uint32x2, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VdupqLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneU64(a arm.Uint64x1, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VdupqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqF32(a arm.Float32x4, b int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VdupqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqF64(a arm.Float64x2, b int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VdupqLaneqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqP8(a arm.Poly8x16, b int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VdupqLaneqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqP16(a arm.Poly16x8, b int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VdupqLaneqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VdupqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VdupqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VdupqLaneqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VdupqLaneqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VdupqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VdupqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VdupqLaneqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupqLaneqU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VdupbLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupbLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// VdupbLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupbLaneS8(a arm.Int8x8, b int) int8 {
	return 0
}

// VdupbLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupbLaneU8(a arm.Uint8x8, b int) uint8 {
	return 0
}

// VduphLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VduphLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// VduphLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VduphLaneS16(a arm.Int16x4, b int) int16 {
	return 0
}

// VduphLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VduphLaneU16(a arm.Uint16x4, b int) uint16 {
	return 0
}

// VdupsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupsLaneF32(a arm.Float32x2, b int) float32 {
	return 0
}

// VdupsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupsLaneS32(a arm.Int32x2, b int) int32 {
	return 0
}

// VdupsLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupsLaneU32(a arm.Uint32x2, b int) uint32 {
	return 0
}

// VdupdLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupdLaneF64(a arm.Float64x1, b int) float64 {
	return 0
}

// VdupdLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupdLaneS64(a arm.Int64x1, b int) int64 {
	return 0
}

// VdupdLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupdLaneU64(a arm.Uint64x1, b int) uint64 {
	return 0
}

// VdupbLaneqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_laneq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupbLaneqP8(a arm.Poly8x16, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// VdupbLaneqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_laneq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupbLaneqS8(a arm.Int8x16, b int) int8 {
	return 0
}

// VdupbLaneqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_laneq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupbLaneqU8(a arm.Uint8x16, b int) uint8 {
	return 0
}

// VduphLaneqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_laneq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VduphLaneqP16(a arm.Poly16x8, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// VduphLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VduphLaneqS16(a arm.Int16x8, b int) int16 {
	return 0
}

// VduphLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VduphLaneqU16(a arm.Uint16x8, b int) uint16 {
	return 0
}

// VdupsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupsLaneqF32(a arm.Float32x4, b int) float32 {
	return 0
}

// VdupsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupsLaneqS32(a arm.Int32x4, b int) int32 {
	return 0
}

// VdupsLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupsLaneqU32(a arm.Uint32x4, b int) uint32 {
	return 0
}

// VdupdLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupdLaneqF64(a arm.Float64x2, b int) float64 {
	return 0
}

// VdupdLaneqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_laneq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupdLaneqS64(a arm.Int64x2, b int) int64 {
	return 0
}

// VdupdLaneqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_laneq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VdupdLaneqU64(a arm.Uint64x2, b int) uint64 {
	return 0
}

// VextF32: ARM NEON intrinsic 
//
// Intrinsic: 'vext_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextF32(a arm.Float32x2, b arm.Float32x2, c int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VextF64: ARM NEON intrinsic 
//
// Intrinsic: 'vext_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextF64(a arm.Float64x1, b arm.Float64x1, c int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VextP8: ARM NEON intrinsic 
//
// Intrinsic: 'vext_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextP8(a arm.Poly8x8, b arm.Poly8x8, c int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VextP16: ARM NEON intrinsic 
//
// Intrinsic: 'vext_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextP16(a arm.Poly16x4, b arm.Poly16x4, c int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VextS8: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VextS16: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VextS32: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VextS64: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VextU8: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VextU16: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VextU32: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VextU64: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VextqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqF32(a arm.Float32x4, b arm.Float32x4, c int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VextqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqF64(a arm.Float64x2, b arm.Float64x2, c int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VextqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqP8(a arm.Poly8x16, b arm.Poly8x16, c int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VextqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqP16(a arm.Poly16x8, b arm.Poly16x8, c int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VextqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VextqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VextqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VextqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VextqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VextqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VextqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VextqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VextqU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VfmaF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_f64'.
// Requires NEON.
func VfmaF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VfmaF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_f32'.
// Requires NEON.
func VfmaF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmaqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_f32'.
// Requires NEON.
func VfmaqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmaqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_f64'.
// Requires NEON.
func VfmaqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VfmaNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_n_f32'.
// Requires NEON.
func VfmaNF32(a arm.Float32x2, b arm.Float32x2, c float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmaqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_n_f32'.
// Requires NEON.
func VfmaqNF32(a arm.Float32x4, b arm.Float32x4, c float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmaqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_n_f64'.
// Requires NEON.
func VfmaqNF64(a arm.Float64x2, b arm.Float64x2, c float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VfmaLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmaLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaLaneF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VfmadLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmad_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmadLaneF64(a float64, b float64, c arm.Float64x1, lane int) float64 {
	return 0
}

// VfmasLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmas_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmasLaneF32(a float32, b float32, c arm.Float32x2, lane int) float32 {
	return 0
}

// VfmaLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmaLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaLaneqF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x2, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VfmadLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmad_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmadLaneqF64(a float64, b float64, c arm.Float64x2, lane int) float64 {
	return 0
}

// VfmasLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmas_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmasLaneqF32(a float32, b float32, c arm.Float32x4, lane int) float32 {
	return 0
}

// VfmaqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmaqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaqLaneF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x1, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VfmaqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmaqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmaqLaneqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VfmsF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_f64'.
// Requires NEON.
func VfmsF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VfmsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_f32'.
// Requires NEON.
func VfmsF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_f32'.
// Requires NEON.
func VfmsqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_f64'.
// Requires NEON.
func VfmsqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VfmsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmsLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsLaneF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VfmsdLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsd_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsdLaneF64(a float64, b float64, c arm.Float64x1, lane int) float64 {
	return 0
}

// VfmssLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmss_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmssLaneF32(a float32, b float32, c arm.Float32x2, lane int) float32 {
	return 0
}

// VfmsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VfmsLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsLaneqF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x2, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VfmsdLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsd_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsdLaneqF64(a float64, b float64, c arm.Float64x2, lane int) float64 {
	return 0
}

// VfmssLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmss_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmssLaneqF32(a float32, b float32, c arm.Float32x4, lane int) float32 {
	return 0
}

// VfmsqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmsqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsqLaneF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x1, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VfmsqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VfmsqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VfmsqLaneqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vld1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1F32(a *float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vld1F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1F64(a *float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// Vld1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1P8(a *arm.Poly8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vld1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1P16(a *arm.Poly16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vld1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1S8(a *int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vld1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1S16(a *int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vld1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1S32(a *int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vld1S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1S64(a *int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// Vld1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1U8(a *uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vld1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1U16(a *uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vld1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1U32(a *uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vld1U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1U64(a *uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// Vld1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qF32(a *float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vld1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qF64(a *float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vld1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qP8(a *arm.Poly8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vld1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qP16(a *arm.Poly16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vld1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qS8(a *int8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vld1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qS16(a *int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vld1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qS32(a *int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vld1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qS64(a *int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vld1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qU8(a *uint8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vld1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qU16(a *uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vld1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qU32(a *uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vld1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qU64(a *uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Vld1DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupF32(a *float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vld1DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupF64(a *float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// Vld1DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupP8(a *uint8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vld1DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupP16(a *uint16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vld1DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupS8(a *int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vld1DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupS16(a *int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vld1DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupS32(a *int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vld1DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupS64(a *int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// Vld1DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupU8(a *uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vld1DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupU16(a *uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vld1DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupU32(a *uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vld1DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1DupU64(a *uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// Vld1qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupF32(a *float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vld1qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupF64(a *float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vld1qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupP8(a *uint8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vld1qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupP16(a *uint16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vld1qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupS8(a *int8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vld1qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupS16(a *int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vld1qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupS32(a *int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vld1qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupS64(a *int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vld1qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupU8(a *uint8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vld1qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupU16(a *uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vld1qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupU32(a *uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vld1qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qDupU64(a *uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Vld1LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneF32(src *float32, vec arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vld1LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneF64(src *float64, vec arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// Vld1LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneP8(src *arm.Poly8, vec arm.Poly8x8, lane int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vld1LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneP16(src *arm.Poly16, vec arm.Poly16x4, lane int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vld1LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneS8(src *int8, vec arm.Int8x8, lane int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vld1LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneS16(src *int16, vec arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vld1LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneS32(src *int32, vec arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vld1LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneS64(src *int64, vec arm.Int64x1, lane int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// Vld1LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneU8(src *uint8, vec arm.Uint8x8, lane int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vld1LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneU16(src *uint16, vec arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vld1LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneU32(src *uint32, vec arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vld1LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1LaneU64(src *uint64, vec arm.Uint64x1, lane int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// Vld1qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneF32(src *float32, vec arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vld1qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneF64(src *float64, vec arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vld1qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneP8(src *arm.Poly8, vec arm.Poly8x16, lane int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vld1qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneP16(src *arm.Poly16, vec arm.Poly16x8, lane int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vld1qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneS8(src *int8, vec arm.Int8x16, lane int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vld1qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneS16(src *int16, vec arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vld1qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneS32(src *int32, vec arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vld1qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneS64(src *int64, vec arm.Int64x2, lane int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vld1qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneU8(src *uint8, vec arm.Uint8x16, lane int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vld1qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneU16(src *uint16, vec arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vld1qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneU32(src *uint32, vec arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vld1qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld1qLaneU64(src *uint64, vec arm.Uint64x2, lane int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Vld2S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2S64(a *int64) (dst [2]arm.Int64x1) {
	return [2]arm.Int64x1{}
}

// Vld2U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2U64(a *uint64) (dst [2]arm.Uint64x1) {
	return [2]arm.Uint64x1{}
}

// Vld2F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2F64(a *float64) (dst [2]arm.Float64x1) {
	return [2]arm.Float64x1{}
}

// Vld2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2S8(a *int8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// Vld2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2P8(a *arm.Poly8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// Vld2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2S16(a *int16) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// Vld2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2P16(a *arm.Poly16) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// Vld2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2S32(a *int32) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// Vld2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2U8(a *uint8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// Vld2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2U16(a *uint16) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// Vld2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2U32(a *uint32) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// Vld2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2F32(a *float32) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// Vld2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qS8(a *int8) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// Vld2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qP8(a *arm.Poly8) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// Vld2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qS16(a *int16) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// Vld2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qP16(a *arm.Poly16) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// Vld2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qS32(a *int32) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// Vld2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qS64(a *int64) (dst [2]arm.Int64x2) {
	return [2]arm.Int64x2{}
}

// Vld2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qU8(a *uint8) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// Vld2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qU16(a *uint16) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// Vld2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qU32(a *uint32) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Vld2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qU64(a *uint64) (dst [2]arm.Uint64x2) {
	return [2]arm.Uint64x2{}
}

// Vld2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qF32(a *float32) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// Vld2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qF64(a *float64) (dst [2]arm.Float64x2) {
	return [2]arm.Float64x2{}
}

// Vld3S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3S64(a *int64) (dst [3]arm.Int64x1) {
	return [3]arm.Int64x1{}
}

// Vld3U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3U64(a *uint64) (dst [3]arm.Uint64x1) {
	return [3]arm.Uint64x1{}
}

// Vld3F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3F64(a *float64) (dst [3]arm.Float64x1) {
	return [3]arm.Float64x1{}
}

// Vld3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3S8(a *int8) (dst [3]arm.Int8x8) {
	return [3]arm.Int8x8{}
}

// Vld3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3P8(a *arm.Poly8) (dst [3]arm.Poly8x8) {
	return [3]arm.Poly8x8{}
}

// Vld3S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3S16(a *int16) (dst [3]arm.Int16x4) {
	return [3]arm.Int16x4{}
}

// Vld3P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3P16(a *arm.Poly16) (dst [3]arm.Poly16x4) {
	return [3]arm.Poly16x4{}
}

// Vld3S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3S32(a *int32) (dst [3]arm.Int32x2) {
	return [3]arm.Int32x2{}
}

// Vld3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3U8(a *uint8) (dst [3]arm.Uint8x8) {
	return [3]arm.Uint8x8{}
}

// Vld3U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3U16(a *uint16) (dst [3]arm.Uint16x4) {
	return [3]arm.Uint16x4{}
}

// Vld3U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3U32(a *uint32) (dst [3]arm.Uint32x2) {
	return [3]arm.Uint32x2{}
}

// Vld3F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3F32(a *float32) (dst [3]arm.Float32x2) {
	return [3]arm.Float32x2{}
}

// Vld3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qS8(a *int8) (dst [3]arm.Int8x16) {
	return [3]arm.Int8x16{}
}

// Vld3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qP8(a *arm.Poly8) (dst [3]arm.Poly8x16) {
	return [3]arm.Poly8x16{}
}

// Vld3qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qS16(a *int16) (dst [3]arm.Int16x8) {
	return [3]arm.Int16x8{}
}

// Vld3qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qP16(a *arm.Poly16) (dst [3]arm.Poly16x8) {
	return [3]arm.Poly16x8{}
}

// Vld3qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qS32(a *int32) (dst [3]arm.Int32x4) {
	return [3]arm.Int32x4{}
}

// Vld3qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qS64(a *int64) (dst [3]arm.Int64x2) {
	return [3]arm.Int64x2{}
}

// Vld3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qU8(a *uint8) (dst [3]arm.Uint8x16) {
	return [3]arm.Uint8x16{}
}

// Vld3qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qU16(a *uint16) (dst [3]arm.Uint16x8) {
	return [3]arm.Uint16x8{}
}

// Vld3qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qU32(a *uint32) (dst [3]arm.Uint32x4) {
	return [3]arm.Uint32x4{}
}

// Vld3qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qU64(a *uint64) (dst [3]arm.Uint64x2) {
	return [3]arm.Uint64x2{}
}

// Vld3qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qF32(a *float32) (dst [3]arm.Float32x4) {
	return [3]arm.Float32x4{}
}

// Vld3qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qF64(a *float64) (dst [3]arm.Float64x2) {
	return [3]arm.Float64x2{}
}

// Vld4S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4S64(a *int64) (dst [4]arm.Int64x1) {
	return [4]arm.Int64x1{}
}

// Vld4U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4U64(a *uint64) (dst [4]arm.Uint64x1) {
	return [4]arm.Uint64x1{}
}

// Vld4F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4F64(a *float64) (dst [4]arm.Float64x1) {
	return [4]arm.Float64x1{}
}

// Vld4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4S8(a *int8) (dst [4]arm.Int8x8) {
	return [4]arm.Int8x8{}
}

// Vld4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4P8(a *arm.Poly8) (dst [4]arm.Poly8x8) {
	return [4]arm.Poly8x8{}
}

// Vld4S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4S16(a *int16) (dst [4]arm.Int16x4) {
	return [4]arm.Int16x4{}
}

// Vld4P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4P16(a *arm.Poly16) (dst [4]arm.Poly16x4) {
	return [4]arm.Poly16x4{}
}

// Vld4S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4S32(a *int32) (dst [4]arm.Int32x2) {
	return [4]arm.Int32x2{}
}

// Vld4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4U8(a *uint8) (dst [4]arm.Uint8x8) {
	return [4]arm.Uint8x8{}
}

// Vld4U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4U16(a *uint16) (dst [4]arm.Uint16x4) {
	return [4]arm.Uint16x4{}
}

// Vld4U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4U32(a *uint32) (dst [4]arm.Uint32x2) {
	return [4]arm.Uint32x2{}
}

// Vld4F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4F32(a *float32) (dst [4]arm.Float32x2) {
	return [4]arm.Float32x2{}
}

// Vld4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qS8(a *int8) (dst [4]arm.Int8x16) {
	return [4]arm.Int8x16{}
}

// Vld4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qP8(a *arm.Poly8) (dst [4]arm.Poly8x16) {
	return [4]arm.Poly8x16{}
}

// Vld4qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qS16(a *int16) (dst [4]arm.Int16x8) {
	return [4]arm.Int16x8{}
}

// Vld4qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qP16(a *arm.Poly16) (dst [4]arm.Poly16x8) {
	return [4]arm.Poly16x8{}
}

// Vld4qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qS32(a *int32) (dst [4]arm.Int32x4) {
	return [4]arm.Int32x4{}
}

// Vld4qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qS64(a *int64) (dst [4]arm.Int64x2) {
	return [4]arm.Int64x2{}
}

// Vld4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qU8(a *uint8) (dst [4]arm.Uint8x16) {
	return [4]arm.Uint8x16{}
}

// Vld4qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qU16(a *uint16) (dst [4]arm.Uint16x8) {
	return [4]arm.Uint16x8{}
}

// Vld4qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qU32(a *uint32) (dst [4]arm.Uint32x4) {
	return [4]arm.Uint32x4{}
}

// Vld4qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qU64(a *uint64) (dst [4]arm.Uint64x2) {
	return [4]arm.Uint64x2{}
}

// Vld4qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qF32(a *float32) (dst [4]arm.Float32x4) {
	return [4]arm.Float32x4{}
}

// Vld4qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qF64(a *float64) (dst [4]arm.Float64x2) {
	return [4]arm.Float64x2{}
}

// Vld2DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupS8(a *int8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// Vld2DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupS16(a *int16) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// Vld2DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupS32(a *int32) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// Vld2DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupF32(a *float32) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// Vld2DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupF64(a *float64) (dst [2]arm.Float64x1) {
	return [2]arm.Float64x1{}
}

// Vld2DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupU8(a *uint8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// Vld2DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupU16(a *uint16) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// Vld2DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupU32(a *uint32) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// Vld2DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupP8(a *arm.Poly8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// Vld2DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupP16(a *arm.Poly16) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// Vld2DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupS64(a *int64) (dst [2]arm.Int64x1) {
	return [2]arm.Int64x1{}
}

// Vld2DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2DupU64(a *uint64) (dst [2]arm.Uint64x1) {
	return [2]arm.Uint64x1{}
}

// Vld2qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupS8(a *int8) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// Vld2qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupP8(a *arm.Poly8) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// Vld2qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupS16(a *int16) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// Vld2qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupP16(a *arm.Poly16) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// Vld2qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupS32(a *int32) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// Vld2qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupS64(a *int64) (dst [2]arm.Int64x2) {
	return [2]arm.Int64x2{}
}

// Vld2qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupU8(a *uint8) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// Vld2qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupU16(a *uint16) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// Vld2qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupU32(a *uint32) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Vld2qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupU64(a *uint64) (dst [2]arm.Uint64x2) {
	return [2]arm.Uint64x2{}
}

// Vld2qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupF32(a *float32) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// Vld2qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qDupF64(a *float64) (dst [2]arm.Float64x2) {
	return [2]arm.Float64x2{}
}

// Vld3DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupS64(a *int64) (dst [3]arm.Int64x1) {
	return [3]arm.Int64x1{}
}

// Vld3DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupU64(a *uint64) (dst [3]arm.Uint64x1) {
	return [3]arm.Uint64x1{}
}

// Vld3DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupF64(a *float64) (dst [3]arm.Float64x1) {
	return [3]arm.Float64x1{}
}

// Vld3DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupS8(a *int8) (dst [3]arm.Int8x8) {
	return [3]arm.Int8x8{}
}

// Vld3DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupP8(a *arm.Poly8) (dst [3]arm.Poly8x8) {
	return [3]arm.Poly8x8{}
}

// Vld3DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupS16(a *int16) (dst [3]arm.Int16x4) {
	return [3]arm.Int16x4{}
}

// Vld3DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupP16(a *arm.Poly16) (dst [3]arm.Poly16x4) {
	return [3]arm.Poly16x4{}
}

// Vld3DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupS32(a *int32) (dst [3]arm.Int32x2) {
	return [3]arm.Int32x2{}
}

// Vld3DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupU8(a *uint8) (dst [3]arm.Uint8x8) {
	return [3]arm.Uint8x8{}
}

// Vld3DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupU16(a *uint16) (dst [3]arm.Uint16x4) {
	return [3]arm.Uint16x4{}
}

// Vld3DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupU32(a *uint32) (dst [3]arm.Uint32x2) {
	return [3]arm.Uint32x2{}
}

// Vld3DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3DupF32(a *float32) (dst [3]arm.Float32x2) {
	return [3]arm.Float32x2{}
}

// Vld3qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupS8(a *int8) (dst [3]arm.Int8x16) {
	return [3]arm.Int8x16{}
}

// Vld3qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupP8(a *arm.Poly8) (dst [3]arm.Poly8x16) {
	return [3]arm.Poly8x16{}
}

// Vld3qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupS16(a *int16) (dst [3]arm.Int16x8) {
	return [3]arm.Int16x8{}
}

// Vld3qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupP16(a *arm.Poly16) (dst [3]arm.Poly16x8) {
	return [3]arm.Poly16x8{}
}

// Vld3qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupS32(a *int32) (dst [3]arm.Int32x4) {
	return [3]arm.Int32x4{}
}

// Vld3qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupS64(a *int64) (dst [3]arm.Int64x2) {
	return [3]arm.Int64x2{}
}

// Vld3qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupU8(a *uint8) (dst [3]arm.Uint8x16) {
	return [3]arm.Uint8x16{}
}

// Vld3qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupU16(a *uint16) (dst [3]arm.Uint16x8) {
	return [3]arm.Uint16x8{}
}

// Vld3qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupU32(a *uint32) (dst [3]arm.Uint32x4) {
	return [3]arm.Uint32x4{}
}

// Vld3qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupU64(a *uint64) (dst [3]arm.Uint64x2) {
	return [3]arm.Uint64x2{}
}

// Vld3qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupF32(a *float32) (dst [3]arm.Float32x4) {
	return [3]arm.Float32x4{}
}

// Vld3qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qDupF64(a *float64) (dst [3]arm.Float64x2) {
	return [3]arm.Float64x2{}
}

// Vld4DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupS64(a *int64) (dst [4]arm.Int64x1) {
	return [4]arm.Int64x1{}
}

// Vld4DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupU64(a *uint64) (dst [4]arm.Uint64x1) {
	return [4]arm.Uint64x1{}
}

// Vld4DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupF64(a *float64) (dst [4]arm.Float64x1) {
	return [4]arm.Float64x1{}
}

// Vld4DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupS8(a *int8) (dst [4]arm.Int8x8) {
	return [4]arm.Int8x8{}
}

// Vld4DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupP8(a *arm.Poly8) (dst [4]arm.Poly8x8) {
	return [4]arm.Poly8x8{}
}

// Vld4DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupS16(a *int16) (dst [4]arm.Int16x4) {
	return [4]arm.Int16x4{}
}

// Vld4DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupP16(a *arm.Poly16) (dst [4]arm.Poly16x4) {
	return [4]arm.Poly16x4{}
}

// Vld4DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupS32(a *int32) (dst [4]arm.Int32x2) {
	return [4]arm.Int32x2{}
}

// Vld4DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupU8(a *uint8) (dst [4]arm.Uint8x8) {
	return [4]arm.Uint8x8{}
}

// Vld4DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupU16(a *uint16) (dst [4]arm.Uint16x4) {
	return [4]arm.Uint16x4{}
}

// Vld4DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupU32(a *uint32) (dst [4]arm.Uint32x2) {
	return [4]arm.Uint32x2{}
}

// Vld4DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4DupF32(a *float32) (dst [4]arm.Float32x2) {
	return [4]arm.Float32x2{}
}

// Vld4qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupS8(a *int8) (dst [4]arm.Int8x16) {
	return [4]arm.Int8x16{}
}

// Vld4qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupP8(a *arm.Poly8) (dst [4]arm.Poly8x16) {
	return [4]arm.Poly8x16{}
}

// Vld4qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupS16(a *int16) (dst [4]arm.Int16x8) {
	return [4]arm.Int16x8{}
}

// Vld4qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupP16(a *arm.Poly16) (dst [4]arm.Poly16x8) {
	return [4]arm.Poly16x8{}
}

// Vld4qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupS32(a *int32) (dst [4]arm.Int32x4) {
	return [4]arm.Int32x4{}
}

// Vld4qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupS64(a *int64) (dst [4]arm.Int64x2) {
	return [4]arm.Int64x2{}
}

// Vld4qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupU8(a *uint8) (dst [4]arm.Uint8x16) {
	return [4]arm.Uint8x16{}
}

// Vld4qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupU16(a *uint16) (dst [4]arm.Uint16x8) {
	return [4]arm.Uint16x8{}
}

// Vld4qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupU32(a *uint32) (dst [4]arm.Uint32x4) {
	return [4]arm.Uint32x4{}
}

// Vld4qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupU64(a *uint64) (dst [4]arm.Uint64x2) {
	return [4]arm.Uint64x2{}
}

// Vld4qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupF32(a *float32) (dst [4]arm.Float32x4) {
	return [4]arm.Float32x4{}
}

// Vld4qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qDupF64(a *float64) (dst [4]arm.Float64x2) {
	return [4]arm.Float64x2{}
}

// Vld2LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneF32(ptr *float32, b [2]arm.Float32x2, c int) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// Vld2LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneF64(ptr *float64, b [2]arm.Float64x1, c int) (dst [2]arm.Float64x1) {
	return [2]arm.Float64x1{}
}

// Vld2LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneP8(ptr *arm.Poly8, b [2]arm.Poly8x8, c int) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// Vld2LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneP16(ptr *arm.Poly16, b [2]arm.Poly16x4, c int) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// Vld2LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneS8(ptr *int8, b [2]arm.Int8x8, c int) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// Vld2LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneS16(ptr *int16, b [2]arm.Int16x4, c int) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// Vld2LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneS32(ptr *int32, b [2]arm.Int32x2, c int) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// Vld2LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneS64(ptr *int64, b [2]arm.Int64x1, c int) (dst [2]arm.Int64x1) {
	return [2]arm.Int64x1{}
}

// Vld2LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneU8(ptr *uint8, b [2]arm.Uint8x8, c int) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// Vld2LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneU16(ptr *uint16, b [2]arm.Uint16x4, c int) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// Vld2LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneU32(ptr *uint32, b [2]arm.Uint32x2, c int) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// Vld2LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2LaneU64(ptr *uint64, b [2]arm.Uint64x1, c int) (dst [2]arm.Uint64x1) {
	return [2]arm.Uint64x1{}
}

// Vld2qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneF32(ptr *float32, b [2]arm.Float32x4, c int) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// Vld2qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneF64(ptr *float64, b [2]arm.Float64x2, c int) (dst [2]arm.Float64x2) {
	return [2]arm.Float64x2{}
}

// Vld2qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneP8(ptr *arm.Poly8, b [2]arm.Poly8x16, c int) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// Vld2qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneP16(ptr *arm.Poly16, b [2]arm.Poly16x8, c int) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// Vld2qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneS8(ptr *int8, b [2]arm.Int8x16, c int) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// Vld2qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneS16(ptr *int16, b [2]arm.Int16x8, c int) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// Vld2qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneS32(ptr *int32, b [2]arm.Int32x4, c int) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// Vld2qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneS64(ptr *int64, b [2]arm.Int64x2, c int) (dst [2]arm.Int64x2) {
	return [2]arm.Int64x2{}
}

// Vld2qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneU8(ptr *uint8, b [2]arm.Uint8x16, c int) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// Vld2qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneU16(ptr *uint16, b [2]arm.Uint16x8, c int) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// Vld2qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneU32(ptr *uint32, b [2]arm.Uint32x4, c int) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Vld2qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld2qLaneU64(ptr *uint64, b [2]arm.Uint64x2, c int) (dst [2]arm.Uint64x2) {
	return [2]arm.Uint64x2{}
}

// Vld3LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneF32(ptr *float32, b [3]arm.Float32x2, c int) (dst [3]arm.Float32x2) {
	return [3]arm.Float32x2{}
}

// Vld3LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneF64(ptr *float64, b [3]arm.Float64x1, c int) (dst [3]arm.Float64x1) {
	return [3]arm.Float64x1{}
}

// Vld3LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneP8(ptr *arm.Poly8, b [3]arm.Poly8x8, c int) (dst [3]arm.Poly8x8) {
	return [3]arm.Poly8x8{}
}

// Vld3LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneP16(ptr *arm.Poly16, b [3]arm.Poly16x4, c int) (dst [3]arm.Poly16x4) {
	return [3]arm.Poly16x4{}
}

// Vld3LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneS8(ptr *int8, b [3]arm.Int8x8, c int) (dst [3]arm.Int8x8) {
	return [3]arm.Int8x8{}
}

// Vld3LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneS16(ptr *int16, b [3]arm.Int16x4, c int) (dst [3]arm.Int16x4) {
	return [3]arm.Int16x4{}
}

// Vld3LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneS32(ptr *int32, b [3]arm.Int32x2, c int) (dst [3]arm.Int32x2) {
	return [3]arm.Int32x2{}
}

// Vld3LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneS64(ptr *int64, b [3]arm.Int64x1, c int) (dst [3]arm.Int64x1) {
	return [3]arm.Int64x1{}
}

// Vld3LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneU8(ptr *uint8, b [3]arm.Uint8x8, c int) (dst [3]arm.Uint8x8) {
	return [3]arm.Uint8x8{}
}

// Vld3LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneU16(ptr *uint16, b [3]arm.Uint16x4, c int) (dst [3]arm.Uint16x4) {
	return [3]arm.Uint16x4{}
}

// Vld3LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneU32(ptr *uint32, b [3]arm.Uint32x2, c int) (dst [3]arm.Uint32x2) {
	return [3]arm.Uint32x2{}
}

// Vld3LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3LaneU64(ptr *uint64, b [3]arm.Uint64x1, c int) (dst [3]arm.Uint64x1) {
	return [3]arm.Uint64x1{}
}

// Vld3qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneF32(ptr *float32, b [3]arm.Float32x4, c int) (dst [3]arm.Float32x4) {
	return [3]arm.Float32x4{}
}

// Vld3qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneF64(ptr *float64, b [3]arm.Float64x2, c int) (dst [3]arm.Float64x2) {
	return [3]arm.Float64x2{}
}

// Vld3qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneP8(ptr *arm.Poly8, b [3]arm.Poly8x16, c int) (dst [3]arm.Poly8x16) {
	return [3]arm.Poly8x16{}
}

// Vld3qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneP16(ptr *arm.Poly16, b [3]arm.Poly16x8, c int) (dst [3]arm.Poly16x8) {
	return [3]arm.Poly16x8{}
}

// Vld3qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneS8(ptr *int8, b [3]arm.Int8x16, c int) (dst [3]arm.Int8x16) {
	return [3]arm.Int8x16{}
}

// Vld3qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneS16(ptr *int16, b [3]arm.Int16x8, c int) (dst [3]arm.Int16x8) {
	return [3]arm.Int16x8{}
}

// Vld3qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneS32(ptr *int32, b [3]arm.Int32x4, c int) (dst [3]arm.Int32x4) {
	return [3]arm.Int32x4{}
}

// Vld3qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneS64(ptr *int64, b [3]arm.Int64x2, c int) (dst [3]arm.Int64x2) {
	return [3]arm.Int64x2{}
}

// Vld3qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneU8(ptr *uint8, b [3]arm.Uint8x16, c int) (dst [3]arm.Uint8x16) {
	return [3]arm.Uint8x16{}
}

// Vld3qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneU16(ptr *uint16, b [3]arm.Uint16x8, c int) (dst [3]arm.Uint16x8) {
	return [3]arm.Uint16x8{}
}

// Vld3qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneU32(ptr *uint32, b [3]arm.Uint32x4, c int) (dst [3]arm.Uint32x4) {
	return [3]arm.Uint32x4{}
}

// Vld3qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld3qLaneU64(ptr *uint64, b [3]arm.Uint64x2, c int) (dst [3]arm.Uint64x2) {
	return [3]arm.Uint64x2{}
}

// Vld4LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneF32(ptr *float32, b [4]arm.Float32x2, c int) (dst [4]arm.Float32x2) {
	return [4]arm.Float32x2{}
}

// Vld4LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneF64(ptr *float64, b [4]arm.Float64x1, c int) (dst [4]arm.Float64x1) {
	return [4]arm.Float64x1{}
}

// Vld4LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneP8(ptr *arm.Poly8, b [4]arm.Poly8x8, c int) (dst [4]arm.Poly8x8) {
	return [4]arm.Poly8x8{}
}

// Vld4LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneP16(ptr *arm.Poly16, b [4]arm.Poly16x4, c int) (dst [4]arm.Poly16x4) {
	return [4]arm.Poly16x4{}
}

// Vld4LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneS8(ptr *int8, b [4]arm.Int8x8, c int) (dst [4]arm.Int8x8) {
	return [4]arm.Int8x8{}
}

// Vld4LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneS16(ptr *int16, b [4]arm.Int16x4, c int) (dst [4]arm.Int16x4) {
	return [4]arm.Int16x4{}
}

// Vld4LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneS32(ptr *int32, b [4]arm.Int32x2, c int) (dst [4]arm.Int32x2) {
	return [4]arm.Int32x2{}
}

// Vld4LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneS64(ptr *int64, b [4]arm.Int64x1, c int) (dst [4]arm.Int64x1) {
	return [4]arm.Int64x1{}
}

// Vld4LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneU8(ptr *uint8, b [4]arm.Uint8x8, c int) (dst [4]arm.Uint8x8) {
	return [4]arm.Uint8x8{}
}

// Vld4LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneU16(ptr *uint16, b [4]arm.Uint16x4, c int) (dst [4]arm.Uint16x4) {
	return [4]arm.Uint16x4{}
}

// Vld4LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneU32(ptr *uint32, b [4]arm.Uint32x2, c int) (dst [4]arm.Uint32x2) {
	return [4]arm.Uint32x2{}
}

// Vld4LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4LaneU64(ptr *uint64, b [4]arm.Uint64x1, c int) (dst [4]arm.Uint64x1) {
	return [4]arm.Uint64x1{}
}

// Vld4qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneF32(ptr *float32, b [4]arm.Float32x4, c int) (dst [4]arm.Float32x4) {
	return [4]arm.Float32x4{}
}

// Vld4qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneF64(ptr *float64, b [4]arm.Float64x2, c int) (dst [4]arm.Float64x2) {
	return [4]arm.Float64x2{}
}

// Vld4qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneP8(ptr *arm.Poly8, b [4]arm.Poly8x16, c int) (dst [4]arm.Poly8x16) {
	return [4]arm.Poly8x16{}
}

// Vld4qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneP16(ptr *arm.Poly16, b [4]arm.Poly16x8, c int) (dst [4]arm.Poly16x8) {
	return [4]arm.Poly16x8{}
}

// Vld4qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneS8(ptr *int8, b [4]arm.Int8x16, c int) (dst [4]arm.Int8x16) {
	return [4]arm.Int8x16{}
}

// Vld4qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneS16(ptr *int16, b [4]arm.Int16x8, c int) (dst [4]arm.Int16x8) {
	return [4]arm.Int16x8{}
}

// Vld4qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneS32(ptr *int32, b [4]arm.Int32x4, c int) (dst [4]arm.Int32x4) {
	return [4]arm.Int32x4{}
}

// Vld4qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneS64(ptr *int64, b [4]arm.Int64x2, c int) (dst [4]arm.Int64x2) {
	return [4]arm.Int64x2{}
}

// Vld4qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneU8(ptr *uint8, b [4]arm.Uint8x16, c int) (dst [4]arm.Uint8x16) {
	return [4]arm.Uint8x16{}
}

// Vld4qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneU16(ptr *uint16, b [4]arm.Uint16x8, c int) (dst [4]arm.Uint16x8) {
	return [4]arm.Uint16x8{}
}

// Vld4qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneU32(ptr *uint32, b [4]arm.Uint32x4, c int) (dst [4]arm.Uint32x4) {
	return [4]arm.Uint32x4{}
}

// Vld4qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vld4qLaneU64(ptr *uint64, b [4]arm.Uint64x2, c int) (dst [4]arm.Uint64x2) {
	return [4]arm.Uint64x2{}
}

// VmaxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_f32'.
// Requires NEON.
func VmaxF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmaxS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_s8'.
// Requires NEON.
func VmaxS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmaxS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_s16'.
// Requires NEON.
func VmaxS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmaxS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_s32'.
// Requires NEON.
func VmaxS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmaxU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_u8'.
// Requires NEON.
func VmaxU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmaxU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_u16'.
// Requires NEON.
func VmaxU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmaxU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_u32'.
// Requires NEON.
func VmaxU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmaxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_f32'.
// Requires NEON.
func VmaxqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmaxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_f64'.
// Requires NEON.
func VmaxqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmaxqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_s8'.
// Requires NEON.
func VmaxqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmaxqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_s16'.
// Requires NEON.
func VmaxqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmaxqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_s32'.
// Requires NEON.
func VmaxqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmaxqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_u8'.
// Requires NEON.
func VmaxqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmaxqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_u16'.
// Requires NEON.
func VmaxqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmaxqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_u32'.
// Requires NEON.
func VmaxqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpmaxS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_s8'.
// Requires NEON.
func VpmaxS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VpmaxS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_s16'.
// Requires NEON.
func VpmaxS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VpmaxS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_s32'.
// Requires NEON.
func VpmaxS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VpmaxU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_u8'.
// Requires NEON.
func VpmaxU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VpmaxU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_u16'.
// Requires NEON.
func VpmaxU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VpmaxU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_u32'.
// Requires NEON.
func VpmaxU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VpmaxqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_s8'.
// Requires NEON.
func VpmaxqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VpmaxqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_s16'.
// Requires NEON.
func VpmaxqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VpmaxqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_s32'.
// Requires NEON.
func VpmaxqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VpmaxqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_u8'.
// Requires NEON.
func VpmaxqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VpmaxqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_u16'.
// Requires NEON.
func VpmaxqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VpmaxqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_u32'.
// Requires NEON.
func VpmaxqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpmaxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_f32'.
// Requires NEON.
func VpmaxF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VpmaxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_f32'.
// Requires NEON.
func VpmaxqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VpmaxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_f64'.
// Requires NEON.
func VpmaxqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VpmaxqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxqd_f64'.
// Requires NEON.
func VpmaxqdF64(a arm.Float64x2) float64 {
	return 0
}

// VpmaxsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxs_f32'.
// Requires NEON.
func VpmaxsF32(a arm.Float32x2) float32 {
	return 0
}

// VpmaxnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnm_f32'.
// Requires NEON.
func VpmaxnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VpmaxnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnmq_f32'.
// Requires NEON.
func VpmaxnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VpmaxnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnmq_f64'.
// Requires NEON.
func VpmaxnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VpmaxnmqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnmqd_f64'.
// Requires NEON.
func VpmaxnmqdF64(a arm.Float64x2) float64 {
	return 0
}

// VpmaxnmsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnms_f32'.
// Requires NEON.
func VpmaxnmsF32(a arm.Float32x2) float32 {
	return 0
}

// VpminS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_s8'.
// Requires NEON.
func VpminS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VpminS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_s16'.
// Requires NEON.
func VpminS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VpminS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_s32'.
// Requires NEON.
func VpminS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VpminU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_u8'.
// Requires NEON.
func VpminU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VpminU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_u16'.
// Requires NEON.
func VpminU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VpminU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_u32'.
// Requires NEON.
func VpminU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VpminqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_s8'.
// Requires NEON.
func VpminqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VpminqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_s16'.
// Requires NEON.
func VpminqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VpminqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_s32'.
// Requires NEON.
func VpminqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VpminqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_u8'.
// Requires NEON.
func VpminqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VpminqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_u16'.
// Requires NEON.
func VpminqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VpminqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_u32'.
// Requires NEON.
func VpminqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VpminF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_f32'.
// Requires NEON.
func VpminF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VpminqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_f32'.
// Requires NEON.
func VpminqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VpminqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_f64'.
// Requires NEON.
func VpminqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VpminqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminqd_f64'.
// Requires NEON.
func VpminqdF64(a arm.Float64x2) float64 {
	return 0
}

// VpminsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmins_f32'.
// Requires NEON.
func VpminsF32(a arm.Float32x2) float32 {
	return 0
}

// VpminnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnm_f32'.
// Requires NEON.
func VpminnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VpminnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnmq_f32'.
// Requires NEON.
func VpminnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VpminnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnmq_f64'.
// Requires NEON.
func VpminnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VpminnmqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnmqd_f64'.
// Requires NEON.
func VpminnmqdF64(a arm.Float64x2) float64 {
	return 0
}

// VpminnmsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnms_f32'.
// Requires NEON.
func VpminnmsF32(a arm.Float32x2) float32 {
	return 0
}

// VmaxnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnm_f32'.
// Requires NEON.
func VmaxnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmaxnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmq_f32'.
// Requires NEON.
func VmaxnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmaxnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmq_f64'.
// Requires NEON.
func VmaxnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmaxvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_f32'.
// Requires NEON.
func VmaxvF32(a arm.Float32x2) float32 {
	return 0
}

// VmaxvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_s8'.
// Requires NEON.
func VmaxvS8(a arm.Int8x8) int8 {
	return 0
}

// VmaxvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_s16'.
// Requires NEON.
func VmaxvS16(a arm.Int16x4) int16 {
	return 0
}

// VmaxvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_s32'.
// Requires NEON.
func VmaxvS32(a arm.Int32x2) int32 {
	return 0
}

// VmaxvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_u8'.
// Requires NEON.
func VmaxvU8(a arm.Uint8x8) uint8 {
	return 0
}

// VmaxvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_u16'.
// Requires NEON.
func VmaxvU16(a arm.Uint16x4) uint16 {
	return 0
}

// VmaxvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_u32'.
// Requires NEON.
func VmaxvU32(a arm.Uint32x2) uint32 {
	return 0
}

// VmaxvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_f32'.
// Requires NEON.
func VmaxvqF32(a arm.Float32x4) float32 {
	return 0
}

// VmaxvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_f64'.
// Requires NEON.
func VmaxvqF64(a arm.Float64x2) float64 {
	return 0
}

// VmaxvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_s8'.
// Requires NEON.
func VmaxvqS8(a arm.Int8x16) int8 {
	return 0
}

// VmaxvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_s16'.
// Requires NEON.
func VmaxvqS16(a arm.Int16x8) int16 {
	return 0
}

// VmaxvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_s32'.
// Requires NEON.
func VmaxvqS32(a arm.Int32x4) int32 {
	return 0
}

// VmaxvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_u8'.
// Requires NEON.
func VmaxvqU8(a arm.Uint8x16) uint8 {
	return 0
}

// VmaxvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_u16'.
// Requires NEON.
func VmaxvqU16(a arm.Uint16x8) uint16 {
	return 0
}

// VmaxvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_u32'.
// Requires NEON.
func VmaxvqU32(a arm.Uint32x4) uint32 {
	return 0
}

// VmaxnmvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmv_f32'.
// Requires NEON.
func VmaxnmvF32(a arm.Float32x2) float32 {
	return 0
}

// VmaxnmvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmvq_f32'.
// Requires NEON.
func VmaxnmvqF32(a arm.Float32x4) float32 {
	return 0
}

// VmaxnmvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmvq_f64'.
// Requires NEON.
func VmaxnmvqF64(a arm.Float64x2) float64 {
	return 0
}

// VminF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_f32'.
// Requires NEON.
func VminF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VminS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_s8'.
// Requires NEON.
func VminS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VminS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_s16'.
// Requires NEON.
func VminS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VminS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_s32'.
// Requires NEON.
func VminS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VminU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_u8'.
// Requires NEON.
func VminU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VminU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_u16'.
// Requires NEON.
func VminU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VminU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_u32'.
// Requires NEON.
func VminU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VminqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_f32'.
// Requires NEON.
func VminqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VminqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_f64'.
// Requires NEON.
func VminqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VminqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_s8'.
// Requires NEON.
func VminqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VminqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_s16'.
// Requires NEON.
func VminqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VminqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_s32'.
// Requires NEON.
func VminqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VminqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_u8'.
// Requires NEON.
func VminqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VminqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_u16'.
// Requires NEON.
func VminqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VminqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_u32'.
// Requires NEON.
func VminqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VminnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnm_f32'.
// Requires NEON.
func VminnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VminnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmq_f32'.
// Requires NEON.
func VminnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VminnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmq_f64'.
// Requires NEON.
func VminnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VminvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_f32'.
// Requires NEON.
func VminvF32(a arm.Float32x2) float32 {
	return 0
}

// VminvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_s8'.
// Requires NEON.
func VminvS8(a arm.Int8x8) int8 {
	return 0
}

// VminvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_s16'.
// Requires NEON.
func VminvS16(a arm.Int16x4) int16 {
	return 0
}

// VminvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_s32'.
// Requires NEON.
func VminvS32(a arm.Int32x2) int32 {
	return 0
}

// VminvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_u8'.
// Requires NEON.
func VminvU8(a arm.Uint8x8) uint8 {
	return 0
}

// VminvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_u16'.
// Requires NEON.
func VminvU16(a arm.Uint16x4) uint16 {
	return 0
}

// VminvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_u32'.
// Requires NEON.
func VminvU32(a arm.Uint32x2) uint32 {
	return 0
}

// VminvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_f32'.
// Requires NEON.
func VminvqF32(a arm.Float32x4) float32 {
	return 0
}

// VminvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_f64'.
// Requires NEON.
func VminvqF64(a arm.Float64x2) float64 {
	return 0
}

// VminvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_s8'.
// Requires NEON.
func VminvqS8(a arm.Int8x16) int8 {
	return 0
}

// VminvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_s16'.
// Requires NEON.
func VminvqS16(a arm.Int16x8) int16 {
	return 0
}

// VminvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_s32'.
// Requires NEON.
func VminvqS32(a arm.Int32x4) int32 {
	return 0
}

// VminvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_u8'.
// Requires NEON.
func VminvqU8(a arm.Uint8x16) uint8 {
	return 0
}

// VminvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_u16'.
// Requires NEON.
func VminvqU16(a arm.Uint16x8) uint16 {
	return 0
}

// VminvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_u32'.
// Requires NEON.
func VminvqU32(a arm.Uint32x4) uint32 {
	return 0
}

// VminnmvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmv_f32'.
// Requires NEON.
func VminnmvF32(a arm.Float32x2) float32 {
	return 0
}

// VminnmvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmvq_f32'.
// Requires NEON.
func VminnmvqF32(a arm.Float32x4) float32 {
	return 0
}

// VminnmvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmvq_f64'.
// Requires NEON.
func VminnmvqF64(a arm.Float64x2) float64 {
	return 0
}

// VmlaF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_f32'.
// Requires NEON.
func VmlaF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlaF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_f64'.
// Requires NEON.
func VmlaF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmlaqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_f32'.
// Requires NEON.
func VmlaqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlaqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_f64'.
// Requires NEON.
func VmlaqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmlaLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlaLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlaLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlaLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlaLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlaLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlaLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneqS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x8, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlaLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneqS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x4, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlaLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneqU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x8, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlaLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaLaneqU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x4, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlaqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlaqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x4, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlaqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x2, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlaqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x4, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlaqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x2, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlaqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlaqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlaqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlaqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlaqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlaqLaneqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_f32'.
// Requires NEON.
func VmlsF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlsF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_f64'.
// Requires NEON.
func VmlsF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmlsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_f32'.
// Requires NEON.
func VmlsqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_f64'.
// Requires NEON.
func VmlsqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmlsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlsLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlsLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlsLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmlsLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneqS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x8, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmlsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneqS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x4, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmlsLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneqU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x8, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmlsLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsLaneqU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x4, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmlsqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlsqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x4, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlsqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x2, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlsqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x4, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlsqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x2, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmlsqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmlsqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmlsqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmlsqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmlsqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmlsqLaneqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmovNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_f32'.
// Requires NEON.
func VmovNF32(a float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmovNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_f64'.
// Requires NEON.
func VmovNF64(a float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmovNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_p8'.
// Requires NEON.
func VmovNP8(a arm.Poly8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VmovNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_p16'.
// Requires NEON.
func VmovNP16(a arm.Poly16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// VmovNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s8'.
// Requires NEON.
func VmovNS8(a int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VmovNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s16'.
// Requires NEON.
func VmovNS16(a int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmovNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s32'.
// Requires NEON.
func VmovNS32(a int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmovNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s64'.
// Requires NEON.
func VmovNS64(a int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VmovNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u8'.
// Requires NEON.
func VmovNU8(a uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VmovNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u16'.
// Requires NEON.
func VmovNU16(a uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmovNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u32'.
// Requires NEON.
func VmovNU32(a uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmovNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u64'.
// Requires NEON.
func VmovNU64(a uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VmovqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_f32'.
// Requires NEON.
func VmovqNF32(a float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmovqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_f64'.
// Requires NEON.
func VmovqNF64(a float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmovqNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_p8'.
// Requires NEON.
func VmovqNP8(a arm.Poly8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VmovqNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_p16'.
// Requires NEON.
func VmovqNP16(a arm.Poly16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// VmovqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s8'.
// Requires NEON.
func VmovqNS8(a int8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VmovqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s16'.
// Requires NEON.
func VmovqNS16(a int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmovqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s32'.
// Requires NEON.
func VmovqNS32(a int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmovqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s64'.
// Requires NEON.
func VmovqNS64(a int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VmovqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u8'.
// Requires NEON.
func VmovqNU8(a uint8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VmovqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u16'.
// Requires NEON.
func VmovqNU16(a uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmovqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u32'.
// Requires NEON.
func VmovqNU32(a uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmovqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u64'.
// Requires NEON.
func VmovqNU64(a uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VmulLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneF32(a arm.Float32x2, b arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmulLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneF64(a arm.Float64x1, b arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmulLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneS16(a arm.Int16x4, b arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmulLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneS32(a arm.Int32x2, b arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmulLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneU16(a arm.Uint16x4, b arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmulLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneU32(a arm.Uint32x2, b arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmuldLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmuld_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmuldLaneF64(a float64, b arm.Float64x1, lane int) float64 {
	return 0
}

// VmuldLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmuld_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmuldLaneqF64(a float64, b arm.Float64x2, lane int) float64 {
	return 0
}

// VmulsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmuls_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulsLaneF32(a float32, b arm.Float32x2, lane int) float32 {
	return 0
}

// VmulsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmuls_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulsLaneqF32(a float32, b arm.Float32x4, lane int) float32 {
	return 0
}

// VmulLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneqF32(a arm.Float32x2, b arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VmulLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneqF64(a arm.Float64x1, b arm.Float64x2, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmulLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneqS16(a arm.Int16x4, b arm.Int16x8, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VmulLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneqS32(a arm.Int32x2, b arm.Int32x4, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VmulLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneqU16(a arm.Uint16x4, b arm.Uint16x8, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VmulLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulLaneqU32(a arm.Uint32x2, b arm.Uint32x4, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VmulNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_f64'.
// Requires NEON.
func VmulNF64(a arm.Float64x1, b float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VmulqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneF32(a arm.Float32x4, b arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmulqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneF64(a arm.Float64x2, b arm.Float64x1, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmulqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneS16(a arm.Int16x8, b arm.Int16x4, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmulqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneS32(a arm.Int32x4, b arm.Int32x2, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmulqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneU16(a arm.Uint16x8, b arm.Uint16x4, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmulqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneU32(a arm.Uint32x4, b arm.Uint32x2, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmulqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneqF32(a arm.Float32x4, b arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VmulqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneqF64(a arm.Float64x2, b arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VmulqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneqS16(a arm.Int16x8, b arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VmulqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneqS32(a arm.Int32x4, b arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VmulqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneqU16(a arm.Uint16x8, b arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VmulqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VmulqLaneqU32(a arm.Uint32x4, b arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VnegF32: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_f32'.
// Requires NEON.
func VnegF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VnegF64: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_f64'.
// Requires NEON.
func VnegF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VnegS8: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s8'.
// Requires NEON.
func VnegS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VnegS16: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s16'.
// Requires NEON.
func VnegS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VnegS32: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s32'.
// Requires NEON.
func VnegS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VnegS64: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s64'.
// Requires NEON.
func VnegS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VnegqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_f32'.
// Requires NEON.
func VnegqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VnegqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_f64'.
// Requires NEON.
func VnegqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VnegqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s8'.
// Requires NEON.
func VnegqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VnegqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s16'.
// Requires NEON.
func VnegqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VnegqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s32'.
// Requires NEON.
func VnegqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VnegqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s64'.
// Requires NEON.
func VnegqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VpaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_s8'.
// Requires NEON.
func VpaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VpaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_s16'.
// Requires NEON.
func VpaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VpaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_s32'.
// Requires NEON.
func VpaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VpaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_u8'.
// Requires NEON.
func VpaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VpaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_u16'.
// Requires NEON.
func VpaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VpaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_u32'.
// Requires NEON.
func VpaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VpadddF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddd_f64'.
// Requires NEON.
func VpadddF64(a arm.Float64x2) float64 {
	return 0
}

// VpadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddd_s64'.
// Requires NEON.
func VpadddS64(a arm.Int64x2) int64 {
	return 0
}

// VpadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddd_u64'.
// Requires NEON.
func VpadddU64(a arm.Uint64x2) uint64 {
	return 0
}

// VqabsqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s64'.
// Requires NEON.
func VqabsqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqabsbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsb_s8'.
// Requires NEON.
func VqabsbS8(a int8) int8 {
	return 0
}

// VqabshS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsh_s16'.
// Requires NEON.
func VqabshS16(a int16) int16 {
	return 0
}

// VqabssS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqabss_s32'.
// Requires NEON.
func VqabssS32(a int32) int32 {
	return 0
}

// VqabsdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsd_s64'.
// Requires NEON.
func VqabsdS64(a int64) int64 {
	return 0
}

// VqaddbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddb_s8'.
// Requires NEON.
func VqaddbS8(a int8, b int8) int8 {
	return 0
}

// VqaddhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddh_s16'.
// Requires NEON.
func VqaddhS16(a int16, b int16) int16 {
	return 0
}

// VqaddsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadds_s32'.
// Requires NEON.
func VqaddsS32(a int32, b int32) int32 {
	return 0
}

// VqadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddd_s64'.
// Requires NEON.
func VqadddS64(a int64, b int64) int64 {
	return 0
}

// VqaddbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddb_u8'.
// Requires NEON.
func VqaddbU8(a uint8, b uint8) uint8 {
	return 0
}

// VqaddhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddh_u16'.
// Requires NEON.
func VqaddhU16(a uint16, b uint16) uint16 {
	return 0
}

// VqaddsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadds_u32'.
// Requires NEON.
func VqaddsU32(a uint32, b uint32) uint32 {
	return 0
}

// VqadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddd_u64'.
// Requires NEON.
func VqadddU64(a uint64, b uint64) uint64 {
	return 0
}

// VqdmlalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_s16'.
// Requires NEON.
func VqdmlalS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_s16'.
// Requires NEON.
func VqdmlalHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalHighLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalHighLaneS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalHighLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalHighLaneqS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_n_s16'.
// Requires NEON.
func VqdmlalHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalLaneS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalLaneqS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_n_s16'.
// Requires NEON.
func VqdmlalNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_s32'.
// Requires NEON.
func VqdmlalS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_s32'.
// Requires NEON.
func VqdmlalHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalHighLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalHighLaneS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalHighLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalHighLaneqS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_n_s32'.
// Requires NEON.
func VqdmlalHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalLaneS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalLaneqS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_n_s32'.
// Requires NEON.
func VqdmlalNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlalhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlalh_s16'.
// Requires NEON.
func VqdmlalhS16(a int32, b int16, c int16) int32 {
	return 0
}

// VqdmlalhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlalh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalhLaneS16(a int32, b int16, c arm.Int16x4, d int) int32 {
	return 0
}

// VqdmlalhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlalh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalhLaneqS16(a int32, b int16, c arm.Int16x8, d int) int32 {
	return 0
}

// VqdmlalsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlals_s32'.
// Requires NEON.
func VqdmlalsS32(a int64, b int32, c int32) int64 {
	return 0
}

// VqdmlalsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlals_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalsLaneS32(a int64, b int32, c arm.Int32x2, d int) int64 {
	return 0
}

// VqdmlalsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlals_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlalsLaneqS32(a int64, b int32, c arm.Int32x4, d int) int64 {
	return 0
}

// VqdmlslS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_s16'.
// Requires NEON.
func VqdmlslS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_s16'.
// Requires NEON.
func VqdmlslHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslHighLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslHighLaneS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslHighLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslHighLaneqS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_n_s16'.
// Requires NEON.
func VqdmlslHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslLaneS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslLaneqS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_n_s16'.
// Requires NEON.
func VqdmlslNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmlslS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_s32'.
// Requires NEON.
func VqdmlslS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_s32'.
// Requires NEON.
func VqdmlslHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslHighLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslHighLaneS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslHighLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslHighLaneqS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_n_s32'.
// Requires NEON.
func VqdmlslHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslLaneS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslLaneqS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_n_s32'.
// Requires NEON.
func VqdmlslNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmlslhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlslh_s16'.
// Requires NEON.
func VqdmlslhS16(a int32, b int16, c int16) int32 {
	return 0
}

// VqdmlslhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlslh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslhLaneS16(a int32, b int16, c arm.Int16x4, d int) int32 {
	return 0
}

// VqdmlslhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlslh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslhLaneqS16(a int32, b int16, c arm.Int16x8, d int) int32 {
	return 0
}

// VqdmlslsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsls_s32'.
// Requires NEON.
func VqdmlslsS32(a int64, b int32, c int32) int64 {
	return 0
}

// VqdmlslsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsls_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslsLaneS32(a int64, b int32, c arm.Int32x2, d int) int64 {
	return 0
}

// VqdmlslsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsls_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmlslsLaneqS32(a int64, b int32, c arm.Int32x4, d int) int64 {
	return 0
}

// VqdmulhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhLaneS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqdmulhLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhLaneS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqdmulhqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhqLaneS16(a arm.Int16x8, b arm.Int16x4, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqdmulhqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhqLaneS32(a arm.Int32x4, b arm.Int32x2, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmulhhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhh_s16'.
// Requires NEON.
func VqdmulhhS16(a int16, b int16) int16 {
	return 0
}

// VqdmulhhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhhLaneS16(a int16, b arm.Int16x4, c int) int16 {
	return 0
}

// VqdmulhhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhhLaneqS16(a int16, b arm.Int16x8, c int) int16 {
	return 0
}

// VqdmulhsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhs_s32'.
// Requires NEON.
func VqdmulhsS32(a int32, b int32) int32 {
	return 0
}

// VqdmulhsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhs_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhsLaneS32(a int32, b arm.Int32x2, c int) int32 {
	return 0
}

// VqdmulhsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhs_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmulhsLaneqS32(a int32, b arm.Int32x4, c int) int32 {
	return 0
}

// VqdmullS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_s16'.
// Requires NEON.
func VqdmullS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_s16'.
// Requires NEON.
func VqdmullHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullHighLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullHighLaneS16(a arm.Int16x8, b arm.Int16x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullHighLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullHighLaneqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_n_s16'.
// Requires NEON.
func VqdmullHighNS16(a arm.Int16x8, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullLaneS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullLaneqS16(a arm.Int16x4, b arm.Int16x8, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_n_s16'.
// Requires NEON.
func VqdmullNS16(a arm.Int16x4, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqdmullS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_s32'.
// Requires NEON.
func VqdmullS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_s32'.
// Requires NEON.
func VqdmullHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullHighLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullHighLaneS32(a arm.Int32x4, b arm.Int32x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullHighLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullHighLaneqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_n_s32'.
// Requires NEON.
func VqdmullHighNS32(a arm.Int32x4, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullLaneS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullLaneqS32(a arm.Int32x2, b arm.Int32x4, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_n_s32'.
// Requires NEON.
func VqdmullNS32(a arm.Int32x2, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqdmullhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmullh_s16'.
// Requires NEON.
func VqdmullhS16(a int16, b int16) int32 {
	return 0
}

// VqdmullhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmullh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullhLaneS16(a int16, b arm.Int16x4, c int) int32 {
	return 0
}

// VqdmullhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmullh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullhLaneqS16(a int16, b arm.Int16x8, c int) int32 {
	return 0
}

// VqdmullsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulls_s32'.
// Requires NEON.
func VqdmullsS32(a int32, b int32) int64 {
	return 0
}

// VqdmullsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulls_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullsLaneS32(a int32, b arm.Int32x2, c int) int64 {
	return 0
}

// VqdmullsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulls_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqdmullsLaneqS32(a int32, b arm.Int32x4, c int) int64 {
	return 0
}

// VqmovnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_s16'.
// Requires NEON.
func VqmovnS16(a arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqmovnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_s32'.
// Requires NEON.
func VqmovnS32(a arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqmovnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_s64'.
// Requires NEON.
func VqmovnS64(a arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqmovnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_u16'.
// Requires NEON.
func VqmovnU16(a arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqmovnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_u32'.
// Requires NEON.
func VqmovnU32(a arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqmovnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_u64'.
// Requires NEON.
func VqmovnU64(a arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqmovnhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnh_s16'.
// Requires NEON.
func VqmovnhS16(a int16) int8 {
	return 0
}

// VqmovnsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovns_s32'.
// Requires NEON.
func VqmovnsS32(a int32) int16 {
	return 0
}

// VqmovndS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnd_s64'.
// Requires NEON.
func VqmovndS64(a int64) int32 {
	return 0
}

// VqmovnhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnh_u16'.
// Requires NEON.
func VqmovnhU16(a uint16) uint8 {
	return 0
}

// VqmovnsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovns_u32'.
// Requires NEON.
func VqmovnsU32(a uint32) uint16 {
	return 0
}

// VqmovndU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnd_u64'.
// Requires NEON.
func VqmovndU64(a uint64) uint32 {
	return 0
}

// VqmovunS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_s16'.
// Requires NEON.
func VqmovunS16(a arm.Int16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqmovunS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_s32'.
// Requires NEON.
func VqmovunS32(a arm.Int32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqmovunS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_s64'.
// Requires NEON.
func VqmovunS64(a arm.Int64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqmovunhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovunh_s16'.
// Requires NEON.
func VqmovunhS16(a int16) int8 {
	return 0
}

// VqmovunsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovuns_s32'.
// Requires NEON.
func VqmovunsS32(a int32) int16 {
	return 0
}

// VqmovundS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovund_s64'.
// Requires NEON.
func VqmovundS64(a int64) int32 {
	return 0
}

// VqnegqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s64'.
// Requires NEON.
func VqnegqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqnegbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegb_s8'.
// Requires NEON.
func VqnegbS8(a int8) int8 {
	return 0
}

// VqneghS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegh_s16'.
// Requires NEON.
func VqneghS16(a int16) int16 {
	return 0
}

// VqnegsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegs_s32'.
// Requires NEON.
func VqnegsS32(a int32) int32 {
	return 0
}

// VqnegdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegd_s64'.
// Requires NEON.
func VqnegdS64(a int64) int64 {
	return 0
}

// VqrdmulhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhLaneS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqrdmulhLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhLaneS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqrdmulhqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhqLaneS16(a arm.Int16x8, b arm.Int16x4, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqrdmulhqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhqLaneS32(a arm.Int32x4, b arm.Int32x2, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqrdmulhhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhh_s16'.
// Requires NEON.
func VqrdmulhhS16(a int16, b int16) int16 {
	return 0
}

// VqrdmulhhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhhLaneS16(a int16, b arm.Int16x4, c int) int16 {
	return 0
}

// VqrdmulhhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhhLaneqS16(a int16, b arm.Int16x8, c int) int16 {
	return 0
}

// VqrdmulhsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhs_s32'.
// Requires NEON.
func VqrdmulhsS32(a int32, b int32) int32 {
	return 0
}

// VqrdmulhsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhs_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhsLaneS32(a int32, b arm.Int32x2, c int) int32 {
	return 0
}

// VqrdmulhsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhs_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrdmulhsLaneqS32(a int32, b arm.Int32x4, c int) int32 {
	return 0
}

// VqrshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s8'.
// Requires NEON.
func VqrshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqrshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s16'.
// Requires NEON.
func VqrshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqrshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s32'.
// Requires NEON.
func VqrshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqrshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s64'.
// Requires NEON.
func VqrshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqrshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u8'.
// Requires NEON.
func VqrshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqrshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u16'.
// Requires NEON.
func VqrshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqrshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u32'.
// Requires NEON.
func VqrshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqrshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u64'.
// Requires NEON.
func VqrshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VqrshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s8'.
// Requires NEON.
func VqrshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqrshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s16'.
// Requires NEON.
func VqrshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqrshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s32'.
// Requires NEON.
func VqrshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqrshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s64'.
// Requires NEON.
func VqrshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqrshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u8'.
// Requires NEON.
func VqrshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqrshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u16'.
// Requires NEON.
func VqrshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqrshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u32'.
// Requires NEON.
func VqrshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqrshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u64'.
// Requires NEON.
func VqrshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqrshlbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlb_s8'.
// Requires NEON.
func VqrshlbS8(a int8, b int8) int8 {
	return 0
}

// VqrshlhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlh_s16'.
// Requires NEON.
func VqrshlhS16(a int16, b int16) int16 {
	return 0
}

// VqrshlsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshls_s32'.
// Requires NEON.
func VqrshlsS32(a int32, b int32) int32 {
	return 0
}

// VqrshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshld_s64'.
// Requires NEON.
func VqrshldS64(a int64, b int64) int64 {
	return 0
}

// VqrshlbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlb_u8'.
// Requires NEON.
func VqrshlbU8(a uint8, b uint8) uint8 {
	return 0
}

// VqrshlhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlh_u16'.
// Requires NEON.
func VqrshlhU16(a uint16, b uint16) uint16 {
	return 0
}

// VqrshlsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshls_u32'.
// Requires NEON.
func VqrshlsU32(a uint32, b uint32) uint32 {
	return 0
}

// VqrshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshld_u64'.
// Requires NEON.
func VqrshldU64(a uint64, b uint64) uint64 {
	return 0
}

// VqrshrnNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnNS16(a arm.Int16x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqrshrnNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnNS32(a arm.Int32x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqrshrnNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnNS64(a arm.Int64x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqrshrnNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnNU16(a arm.Uint16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqrshrnNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnNU32(a arm.Uint32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqrshrnNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnNU64(a arm.Uint64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqrshrnhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnhNS16(a int16, b int) int8 {
	return 0
}

// VqrshrnsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnsNS32(a int32, b int) int16 {
	return 0
}

// VqrshrndNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrndNS64(a int64, b int) int32 {
	return 0
}

// VqrshrnhNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnh_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnhNU16(a uint16, b int) uint8 {
	return 0
}

// VqrshrnsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrns_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrnsNU32(a uint32, b int) uint16 {
	return 0
}

// VqrshrndNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrndNU64(a uint64, b int) uint32 {
	return 0
}

// VqrshrunNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrun_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrunNS16(a arm.Int16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqrshrunNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrun_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrunNS32(a arm.Int32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqrshrunNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrun_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrunNS64(a arm.Int64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqrshrunhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrunh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrunhNS16(a int16, b int) int8 {
	return 0
}

// VqrshrunsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshruns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrunsNS32(a int32, b int) int16 {
	return 0
}

// VqrshrundNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrund_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqrshrundNS64(a int64, b int) int32 {
	return 0
}

// VqshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s8'.
// Requires NEON.
func VqshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s16'.
// Requires NEON.
func VqshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s32'.
// Requires NEON.
func VqshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s64'.
// Requires NEON.
func VqshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u8'.
// Requires NEON.
func VqshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u16'.
// Requires NEON.
func VqshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u32'.
// Requires NEON.
func VqshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u64'.
// Requires NEON.
func VqshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VqshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s8'.
// Requires NEON.
func VqshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s16'.
// Requires NEON.
func VqshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s32'.
// Requires NEON.
func VqshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s64'.
// Requires NEON.
func VqshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u8'.
// Requires NEON.
func VqshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u16'.
// Requires NEON.
func VqshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u32'.
// Requires NEON.
func VqshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u64'.
// Requires NEON.
func VqshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqshlbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_s8'.
// Requires NEON.
func VqshlbS8(a int8, b int8) int8 {
	return 0
}

// VqshlhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_s16'.
// Requires NEON.
func VqshlhS16(a int16, b int16) int16 {
	return 0
}

// VqshlsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_s32'.
// Requires NEON.
func VqshlsS32(a int32, b int32) int32 {
	return 0
}

// VqshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_s64'.
// Requires NEON.
func VqshldS64(a int64, b int64) int64 {
	return 0
}

// VqshlbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_u8'.
// Requires NEON.
func VqshlbU8(a uint8, b uint8) uint8 {
	return 0
}

// VqshlhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_u16'.
// Requires NEON.
func VqshlhU16(a uint16, b uint16) uint16 {
	return 0
}

// VqshlsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_u32'.
// Requires NEON.
func VqshlsU32(a uint32, b uint32) uint32 {
	return 0
}

// VqshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_u64'.
// Requires NEON.
func VqshldU64(a uint64, b uint64) uint64 {
	return 0
}

// VqshlNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqshlNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqshlNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqshlNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VqshlNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqshlNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqshlNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqshlNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VqshlqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VqshlqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VqshlqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VqshlqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VqshlqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqshlqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqshlqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqshlqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqshlbNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlbNS8(a int8, b int) int8 {
	return 0
}

// VqshlhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlhNS16(a int16, b int) int16 {
	return 0
}

// VqshlsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlsNS32(a int32, b int) int32 {
	return 0
}

// VqshldNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshldNS64(a int64, b int) int64 {
	return 0
}

// VqshlbNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlbNU8(a uint8, b int) uint8 {
	return 0
}

// VqshlhNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlhNU16(a uint16, b int) uint16 {
	return 0
}

// VqshlsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlsNU32(a uint32, b int) uint32 {
	return 0
}

// VqshldNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshldNU64(a uint64, b int) uint64 {
	return 0
}

// VqshluNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluNS8(a arm.Int8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqshluNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluNS16(a arm.Int16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqshluNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluNS32(a arm.Int32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqshluNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluNS64(a arm.Int64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VqshluqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluqNS8(a arm.Int8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VqshluqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluqNS16(a arm.Int16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VqshluqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluqNS32(a arm.Int32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VqshluqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluqNS64(a arm.Int64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VqshlubNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlub_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlubNS8(a int8, b int) int8 {
	return 0
}

// VqshluhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshluhNS16(a int16, b int) int16 {
	return 0
}

// VqshlusNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlus_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshlusNS32(a int32, b int) int32 {
	return 0
}

// VqshludNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlud_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshludNS64(a int64, b int) uint64 {
	return 0
}

// VqshrnNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnNS16(a arm.Int16x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VqshrnNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnNS32(a arm.Int32x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VqshrnNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnNS64(a arm.Int64x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VqshrnNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnNU16(a arm.Uint16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqshrnNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnNU32(a arm.Uint32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqshrnNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnNU64(a arm.Uint64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqshrnhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnhNS16(a int16, b int) int8 {
	return 0
}

// VqshrnsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnsNS32(a int32, b int) int16 {
	return 0
}

// VqshrndNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrndNS64(a int64, b int) int32 {
	return 0
}

// VqshrnhNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnh_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnhNU16(a uint16, b int) uint8 {
	return 0
}

// VqshrnsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrns_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrnsNU32(a uint32, b int) uint16 {
	return 0
}

// VqshrndNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrndNU64(a uint64, b int) uint32 {
	return 0
}

// VqshrunNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrun_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrunNS16(a arm.Int16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VqshrunNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrun_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrunNS32(a arm.Int32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VqshrunNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrun_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrunNS64(a arm.Int64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VqshrunhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrunh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrunhNS16(a int16, b int) int8 {
	return 0
}

// VqshrunsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshruns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrunsNS32(a int32, b int) int16 {
	return 0
}

// VqshrundNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrund_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VqshrundNS64(a int64, b int) int32 {
	return 0
}

// VqsubbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubb_s8'.
// Requires NEON.
func VqsubbS8(a int8, b int8) int8 {
	return 0
}

// VqsubhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubh_s16'.
// Requires NEON.
func VqsubhS16(a int16, b int16) int16 {
	return 0
}

// VqsubsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubs_s32'.
// Requires NEON.
func VqsubsS32(a int32, b int32) int32 {
	return 0
}

// VqsubdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubd_s64'.
// Requires NEON.
func VqsubdS64(a int64, b int64) int64 {
	return 0
}

// VqsubbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubb_u8'.
// Requires NEON.
func VqsubbU8(a uint8, b uint8) uint8 {
	return 0
}

// VqsubhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubh_u16'.
// Requires NEON.
func VqsubhU16(a uint16, b uint16) uint16 {
	return 0
}

// VqsubsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubs_u32'.
// Requires NEON.
func VqsubsU32(a uint32, b uint32) uint32 {
	return 0
}

// VqsubdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubd_u64'.
// Requires NEON.
func VqsubdU64(a uint64, b uint64) uint64 {
	return 0
}

// VrbitP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbit_p8'.
// Requires NEON.
func VrbitP8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// VrbitS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbit_s8'.
// Requires NEON.
func VrbitS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VrbitU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbit_u8'.
// Requires NEON.
func VrbitU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VrbitqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbitq_p8'.
// Requires NEON.
func VrbitqP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// VrbitqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbitq_s8'.
// Requires NEON.
func VrbitqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VrbitqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbitq_u8'.
// Requires NEON.
func VrbitqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VrecpeU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpe_u32'.
// Requires NEON.
func VrecpeU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrecpeqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpeq_u32'.
// Requires NEON.
func VrecpeqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrecpesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpes_f32'.
// Requires NEON.
func VrecpesF32(a float32) float32 {
	return 0
}

// VrecpedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecped_f64'.
// Requires NEON.
func VrecpedF64(a float64) float64 {
	return 0
}

// VrecpeF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpe_f32'.
// Requires NEON.
func VrecpeF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrecpeqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpeq_f32'.
// Requires NEON.
func VrecpeqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrecpeqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpeq_f64'.
// Requires NEON.
func VrecpeqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrecpssF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpss_f32'.
// Requires NEON.
func VrecpssF32(a float32, b float32) float32 {
	return 0
}

// VrecpsdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpsd_f64'.
// Requires NEON.
func VrecpsdF64(a float64, b float64) float64 {
	return 0
}

// VrecpsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecps_f32'.
// Requires NEON.
func VrecpsF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrecpsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpsq_f32'.
// Requires NEON.
func VrecpsqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrecpsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpsq_f64'.
// Requires NEON.
func VrecpsqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrecpxsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpxs_f32'.
// Requires NEON.
func VrecpxsF32(a float32) float32 {
	return 0
}

// VrecpxdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpxd_f64'.
// Requires NEON.
func VrecpxdF64(a float64) float64 {
	return 0
}

// Vrev16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16_p8'.
// Requires NEON.
func Vrev16P8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vrev16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16_s8'.
// Requires NEON.
func Vrev16S8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vrev16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16_u8'.
// Requires NEON.
func Vrev16U8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vrev16qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16q_p8'.
// Requires NEON.
func Vrev16qP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vrev16qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16q_s8'.
// Requires NEON.
func Vrev16qS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vrev16qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16q_u8'.
// Requires NEON.
func Vrev16qU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vrev32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_p8'.
// Requires NEON.
func Vrev32P8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vrev32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_p16'.
// Requires NEON.
func Vrev32P16(a arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vrev32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_s8'.
// Requires NEON.
func Vrev32S8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vrev32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_s16'.
// Requires NEON.
func Vrev32S16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vrev32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_u8'.
// Requires NEON.
func Vrev32U8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vrev32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_u16'.
// Requires NEON.
func Vrev32U16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vrev32qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_p8'.
// Requires NEON.
func Vrev32qP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vrev32qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_p16'.
// Requires NEON.
func Vrev32qP16(a arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vrev32qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_s8'.
// Requires NEON.
func Vrev32qS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vrev32qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_s16'.
// Requires NEON.
func Vrev32qS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vrev32qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_u8'.
// Requires NEON.
func Vrev32qU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vrev32qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_u16'.
// Requires NEON.
func Vrev32qU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vrev64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_f32'.
// Requires NEON.
func Vrev64F32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vrev64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_p8'.
// Requires NEON.
func Vrev64P8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vrev64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_p16'.
// Requires NEON.
func Vrev64P16(a arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vrev64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_s8'.
// Requires NEON.
func Vrev64S8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vrev64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_s16'.
// Requires NEON.
func Vrev64S16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vrev64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_s32'.
// Requires NEON.
func Vrev64S32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vrev64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_u8'.
// Requires NEON.
func Vrev64U8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vrev64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_u16'.
// Requires NEON.
func Vrev64U16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vrev64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_u32'.
// Requires NEON.
func Vrev64U32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vrev64qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_f32'.
// Requires NEON.
func Vrev64qF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vrev64qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_p8'.
// Requires NEON.
func Vrev64qP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vrev64qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_p16'.
// Requires NEON.
func Vrev64qP16(a arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vrev64qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_s8'.
// Requires NEON.
func Vrev64qS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vrev64qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_s16'.
// Requires NEON.
func Vrev64qS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vrev64qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_s32'.
// Requires NEON.
func Vrev64qS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vrev64qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_u8'.
// Requires NEON.
func Vrev64qU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vrev64qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_u16'.
// Requires NEON.
func Vrev64qU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vrev64qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_u32'.
// Requires NEON.
func Vrev64qU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrndF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrnd_f32'.
// Requires NEON.
func VrndF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrnd_f64'.
// Requires NEON.
func VrndF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndq_f32'.
// Requires NEON.
func VrndqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndq_f64'.
// Requires NEON.
func VrndqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrndaF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrnda_f32'.
// Requires NEON.
func VrndaF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndaF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrnda_f64'.
// Requires NEON.
func VrndaF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndaqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndaq_f32'.
// Requires NEON.
func VrndaqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndaqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndaq_f64'.
// Requires NEON.
func VrndaqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrndiF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndi_f32'.
// Requires NEON.
func VrndiF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndiF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndi_f64'.
// Requires NEON.
func VrndiF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndiqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndiq_f32'.
// Requires NEON.
func VrndiqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndiqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndiq_f64'.
// Requires NEON.
func VrndiqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrndmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndm_f32'.
// Requires NEON.
func VrndmF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndmF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndm_f64'.
// Requires NEON.
func VrndmF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndmq_f32'.
// Requires NEON.
func VrndmqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndmq_f64'.
// Requires NEON.
func VrndmqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrndnF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndn_f32'.
// Requires NEON.
func VrndnF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndnF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndn_f64'.
// Requires NEON.
func VrndnF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndnqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndnq_f32'.
// Requires NEON.
func VrndnqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndnqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndnq_f64'.
// Requires NEON.
func VrndnqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrndpF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndp_f32'.
// Requires NEON.
func VrndpF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndpF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndp_f64'.
// Requires NEON.
func VrndpF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndpqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndpq_f32'.
// Requires NEON.
func VrndpqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndpqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndpq_f64'.
// Requires NEON.
func VrndpqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrndxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndx_f32'.
// Requires NEON.
func VrndxF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VrndxF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndx_f64'.
// Requires NEON.
func VrndxF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VrndxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndxq_f32'.
// Requires NEON.
func VrndxqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VrndxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndxq_f64'.
// Requires NEON.
func VrndxqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VrshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s8'.
// Requires NEON.
func VrshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VrshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s16'.
// Requires NEON.
func VrshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VrshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s32'.
// Requires NEON.
func VrshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VrshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s64'.
// Requires NEON.
func VrshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VrshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u8'.
// Requires NEON.
func VrshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VrshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u16'.
// Requires NEON.
func VrshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VrshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u32'.
// Requires NEON.
func VrshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u64'.
// Requires NEON.
func VrshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VrshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s8'.
// Requires NEON.
func VrshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VrshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s16'.
// Requires NEON.
func VrshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VrshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s32'.
// Requires NEON.
func VrshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VrshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s64'.
// Requires NEON.
func VrshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VrshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u8'.
// Requires NEON.
func VrshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VrshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u16'.
// Requires NEON.
func VrshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VrshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u32'.
// Requires NEON.
func VrshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u64'.
// Requires NEON.
func VrshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VrshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshld_s64'.
// Requires NEON.
func VrshldS64(a int64, b int64) int64 {
	return 0
}

// VrshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshld_u64'.
// Requires NEON.
func VrshldU64(a uint64, b int64) uint64 {
	return 0
}

// VrshrNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VrshrNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VrshrNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VrshrNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VrshrNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VrshrNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VrshrNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrshrNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VrshrqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VrshrqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VrshrqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VrshrqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VrshrqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VrshrqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VrshrqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrshrqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VrshrdNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrdNS64(a int64, b int) int64 {
	return 0
}

// VrshrdNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrshrdNU64(a uint64, b int) uint64 {
	return 0
}

// VrsraNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VrsraNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VrsraNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VrsraNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VrsraNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VrsraNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VrsraNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VrsraNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VrsraqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VrsraqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VrsraqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VrsraqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VrsraqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VrsraqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VrsraqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VrsraqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsraqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VrsradNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsrad_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsradNS64(a int64, b int64, c int) int64 {
	return 0
}

// VrsradNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsrad_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VrsradNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// Vsha1cqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1cq_u32'.
// Requires NEON.
func Vsha1cqU32(hash_abcd arm.Uint32x4, hash_e uint32, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha1mqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1mq_u32'.
// Requires NEON.
func Vsha1mqU32(hash_abcd arm.Uint32x4, hash_e uint32, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha1pqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1pq_u32'.
// Requires NEON.
func Vsha1pqU32(hash_abcd arm.Uint32x4, hash_e uint32, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha1hU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1h_u32'.
// Requires NEON.
func Vsha1hU32(hash_e uint32) uint32 {
	return 0
}

// Vsha1su0qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1su0q_u32'.
// Requires NEON.
func Vsha1su0qU32(w0_3 arm.Uint32x4, w4_7 arm.Uint32x4, w8_11 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha1su1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1su1q_u32'.
// Requires NEON.
func Vsha1su1qU32(tw0_3 arm.Uint32x4, w12_15 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha256hqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256hq_u32'.
// Requires NEON.
func Vsha256hqU32(hash_abcd arm.Uint32x4, hash_efgh arm.Uint32x4, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha256h2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256h2q_u32'.
// Requires NEON.
func Vsha256h2qU32(hash_efgh arm.Uint32x4, hash_abcd arm.Uint32x4, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha256su0qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256su0q_u32'.
// Requires NEON.
func Vsha256su0qU32(w0_3 arm.Uint32x4, w4_7 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vsha256su1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256su1q_u32'.
// Requires NEON.
func Vsha256su1qU32(tw0_3 arm.Uint32x4, w8_11 arm.Uint32x4, w12_15 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VmullP64: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_p64'.
// Requires NEON.
func VmullP64(a arm.Poly64, b arm.Poly64) (dst arm.Poly128) {
	return arm.Poly128{}
}

// VmullHighP64: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_p64'.
// Requires NEON.
func VmullHighP64(a arm.Poly64x2, b arm.Poly64x2) (dst arm.Poly128) {
	return arm.Poly128{}
}

// VshlNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VshlNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VshlNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VshlNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VshlNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VshlNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VshlNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VshlNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VshlqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VshlqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VshlqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VshlqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VshlqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VshlqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VshlqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VshlqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshlqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VshldNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshldNS64(a int64, b int) int64 {
	return 0
}

// VshldNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshldNU64(a uint64, b int) uint64 {
	return 0
}

// VshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s8'.
// Requires NEON.
func VshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s16'.
// Requires NEON.
func VshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s32'.
// Requires NEON.
func VshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s64'.
// Requires NEON.
func VshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u8'.
// Requires NEON.
func VshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u16'.
// Requires NEON.
func VshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u32'.
// Requires NEON.
func VshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u64'.
// Requires NEON.
func VshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s8'.
// Requires NEON.
func VshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s16'.
// Requires NEON.
func VshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s32'.
// Requires NEON.
func VshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s64'.
// Requires NEON.
func VshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u8'.
// Requires NEON.
func VshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u16'.
// Requires NEON.
func VshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u32'.
// Requires NEON.
func VshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u64'.
// Requires NEON.
func VshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_s64'.
// Requires NEON.
func VshldS64(a int64, b int64) int64 {
	return 0
}

// VshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_u64'.
// Requires NEON.
func VshldU64(a uint64, b uint64) uint64 {
	return 0
}

// VshllHighNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllHighNS8(a arm.Int8x16, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VshllHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllHighNS16(a arm.Int16x8, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VshllHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllHighNS32(a arm.Int32x4, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VshllHighNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllHighNU8(a arm.Uint8x16, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VshllHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllHighNU16(a arm.Uint16x8, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VshllHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllHighNU32(a arm.Uint32x4, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VshllNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllNS8(a arm.Int8x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VshllNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllNS16(a arm.Int16x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VshllNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllNS32(a arm.Int32x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VshllNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllNU8(a arm.Uint8x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VshllNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllNU16(a arm.Uint16x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VshllNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshllNU32(a arm.Uint32x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VshrNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VshrNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VshrNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VshrNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VshrNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VshrNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VshrNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VshrNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VshrqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VshrqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VshrqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VshrqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VshrqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VshrqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VshrqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VshrqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VshrdNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrdNS64(a int64, b int) int64 {
	return 0
}

// VshrdNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VshrdNU64(a uint64, b int) uint64 {
	return 0
}

// VsliNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VsliNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VsliNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VsliNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VsliNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsliNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsliNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VsliNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VsliqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VsliqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsliqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsliqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsliqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsliqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsliqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsliqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsliqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VslidNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vslid_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VslidNS64(a int64, b int64, c int) int64 {
	return 0
}

// VslidNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vslid_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VslidNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// VsqaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u8'.
// Requires NEON.
func VsqaddU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsqaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u16'.
// Requires NEON.
func VsqaddU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsqaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u32'.
// Requires NEON.
func VsqaddU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VsqaddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u64'.
// Requires NEON.
func VsqaddU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VsqaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u8'.
// Requires NEON.
func VsqaddqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsqaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u16'.
// Requires NEON.
func VsqaddqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsqaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u32'.
// Requires NEON.
func VsqaddqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsqaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u64'.
// Requires NEON.
func VsqaddqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsqaddbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddb_u8'.
// Requires NEON.
func VsqaddbU8(a uint8, b int8) uint8 {
	return 0
}

// VsqaddhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddh_u16'.
// Requires NEON.
func VsqaddhU16(a uint16, b int16) uint16 {
	return 0
}

// VsqaddsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadds_u32'.
// Requires NEON.
func VsqaddsU32(a uint32, b int32) uint32 {
	return 0
}

// VsqadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddd_u64'.
// Requires NEON.
func VsqadddU64(a uint64, b int64) uint64 {
	return 0
}

// VsqrtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrt_f32'.
// Requires NEON.
func VsqrtF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// VsqrtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrtq_f32'.
// Requires NEON.
func VsqrtqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// VsqrtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrt_f64'.
// Requires NEON.
func VsqrtF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// VsqrtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrtq_f64'.
// Requires NEON.
func VsqrtqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// VsraNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VsraNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VsraNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VsraNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VsraNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsraNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsraNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VsraNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VsraqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VsraqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsraqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsraqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsraqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsraqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsraqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsraqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsraqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsradNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrad_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsradNS64(a int64, b int64, c int) int64 {
	return 0
}

// VsradNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrad_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsradNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// VsriNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VsriNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VsriNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VsriNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VsriNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VsriNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VsriNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VsriNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VsriqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VsriqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VsriqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VsriqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VsriqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VsriqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VsriqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VsriqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsriqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VsridNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrid_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsridNS64(a int64, b int64, c int) int64 {
	return 0
}

// VsridNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrid_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func VsridNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// Vst1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1F32(a *float32, b arm.Float32x2)  {

}

// Vst1F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1F64(a *float64, b arm.Float64x1)  {

}

// Vst1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1P8(a *arm.Poly8, b arm.Poly8x8)  {

}

// Vst1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1P16(a *arm.Poly16, b arm.Poly16x4)  {

}

// Vst1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1S8(a *int8, b arm.Int8x8)  {

}

// Vst1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1S16(a *int16, b arm.Int16x4)  {

}

// Vst1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1S32(a *int32, b arm.Int32x2)  {

}

// Vst1S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1S64(a *int64, b arm.Int64x1)  {

}

// Vst1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1U8(a *uint8, b arm.Uint8x8)  {

}

// Vst1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1U16(a *uint16, b arm.Uint16x4)  {

}

// Vst1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1U32(a *uint32, b arm.Uint32x2)  {

}

// Vst1U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1U64(a *uint64, b arm.Uint64x1)  {

}

// Vst1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qF32(a *float32, b arm.Float32x4)  {

}

// Vst1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qF64(a *float64, b arm.Float64x2)  {

}

// Vst1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qP8(a *arm.Poly8, b arm.Poly8x16)  {

}

// Vst1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qP16(a *arm.Poly16, b arm.Poly16x8)  {

}

// Vst1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qS8(a *int8, b arm.Int8x16)  {

}

// Vst1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qS16(a *int16, b arm.Int16x8)  {

}

// Vst1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qS32(a *int32, b arm.Int32x4)  {

}

// Vst1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qS64(a *int64, b arm.Int64x2)  {

}

// Vst1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qU8(a *uint8, b arm.Uint8x16)  {

}

// Vst1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qU16(a *uint16, b arm.Uint16x8)  {

}

// Vst1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qU32(a *uint32, b arm.Uint32x4)  {

}

// Vst1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst1qU64(a *uint64, b arm.Uint64x2)  {

}

// Vst1LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneF32(a *float32, b arm.Float32x2, lane int)  {

}

// Vst1LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneF64(a *float64, b arm.Float64x1, lane int)  {

}

// Vst1LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneP8(a *arm.Poly8, b arm.Poly8x8, lane int)  {

}

// Vst1LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneP16(a *arm.Poly16, b arm.Poly16x4, lane int)  {

}

// Vst1LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneS8(a *int8, b arm.Int8x8, lane int)  {

}

// Vst1LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneS16(a *int16, b arm.Int16x4, lane int)  {

}

// Vst1LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneS32(a *int32, b arm.Int32x2, lane int)  {

}

// Vst1LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneS64(a *int64, b arm.Int64x1, lane int)  {

}

// Vst1LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneU8(a *uint8, b arm.Uint8x8, lane int)  {

}

// Vst1LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneU16(a *uint16, b arm.Uint16x4, lane int)  {

}

// Vst1LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneU32(a *uint32, b arm.Uint32x2, lane int)  {

}

// Vst1LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1LaneU64(a *uint64, b arm.Uint64x1, lane int)  {

}

// Vst1qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneF32(a *float32, b arm.Float32x4, lane int)  {

}

// Vst1qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneF64(a *float64, b arm.Float64x2, lane int)  {

}

// Vst1qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneP8(a *arm.Poly8, b arm.Poly8x16, lane int)  {

}

// Vst1qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneP16(a *arm.Poly16, b arm.Poly16x8, lane int)  {

}

// Vst1qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneS8(a *int8, b arm.Int8x16, lane int)  {

}

// Vst1qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneS16(a *int16, b arm.Int16x8, lane int)  {

}

// Vst1qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneS32(a *int32, b arm.Int32x4, lane int)  {

}

// Vst1qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneS64(a *int64, b arm.Int64x2, lane int)  {

}

// Vst1qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneU8(a *uint8, b arm.Uint8x16, lane int)  {

}

// Vst1qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneU16(a *uint16, b arm.Uint16x8, lane int)  {

}

// Vst1qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneU32(a *uint32, b arm.Uint32x4, lane int)  {

}

// Vst1qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Vst1qLaneU64(a *uint64, b arm.Uint64x2, lane int)  {

}

// Vst2S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2S64(a *int64, val [2]arm.Int64x1)  {

}

// Vst2U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2U64(a *uint64, val [2]arm.Uint64x1)  {

}

// Vst2F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2F64(a *float64, val [2]arm.Float64x1)  {

}

// Vst2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2S8(a *int8, val [2]arm.Int8x8)  {

}

// Vst2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2P8(a *arm.Poly8, val [2]arm.Poly8x8)  {

}

// Vst2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2S16(a *int16, val [2]arm.Int16x4)  {

}

// Vst2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2P16(a *arm.Poly16, val [2]arm.Poly16x4)  {

}

// Vst2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2S32(a *int32, val [2]arm.Int32x2)  {

}

// Vst2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2U8(a *uint8, val [2]arm.Uint8x8)  {

}

// Vst2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2U16(a *uint16, val [2]arm.Uint16x4)  {

}

// Vst2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2U32(a *uint32, val [2]arm.Uint32x2)  {

}

// Vst2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2F32(a *float32, val [2]arm.Float32x2)  {

}

// Vst2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qS8(a *int8, val [2]arm.Int8x16)  {

}

// Vst2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qP8(a *arm.Poly8, val [2]arm.Poly8x16)  {

}

// Vst2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qS16(a *int16, val [2]arm.Int16x8)  {

}

// Vst2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qP16(a *arm.Poly16, val [2]arm.Poly16x8)  {

}

// Vst2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qS32(a *int32, val [2]arm.Int32x4)  {

}

// Vst2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qS64(a *int64, val [2]arm.Int64x2)  {

}

// Vst2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qU8(a *uint8, val [2]arm.Uint8x16)  {

}

// Vst2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qU16(a *uint16, val [2]arm.Uint16x8)  {

}

// Vst2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qU32(a *uint32, val [2]arm.Uint32x4)  {

}

// Vst2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qU64(a *uint64, val [2]arm.Uint64x2)  {

}

// Vst2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qF32(a *float32, val [2]arm.Float32x4)  {

}

// Vst2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst2qF64(a *float64, val [2]arm.Float64x2)  {

}

// Vst3S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3S64(a *int64, val [3]arm.Int64x1)  {

}

// Vst3U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3U64(a *uint64, val [3]arm.Uint64x1)  {

}

// Vst3F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3F64(a *float64, val [3]arm.Float64x1)  {

}

// Vst3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3S8(a *int8, val [3]arm.Int8x8)  {

}

// Vst3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3P8(a *arm.Poly8, val [3]arm.Poly8x8)  {

}

// Vst3S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3S16(a *int16, val [3]arm.Int16x4)  {

}

// Vst3P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3P16(a *arm.Poly16, val [3]arm.Poly16x4)  {

}

// Vst3S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3S32(a *int32, val [3]arm.Int32x2)  {

}

// Vst3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3U8(a *uint8, val [3]arm.Uint8x8)  {

}

// Vst3U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3U16(a *uint16, val [3]arm.Uint16x4)  {

}

// Vst3U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3U32(a *uint32, val [3]arm.Uint32x2)  {

}

// Vst3F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3F32(a *float32, val [3]arm.Float32x2)  {

}

// Vst3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qS8(a *int8, val [3]arm.Int8x16)  {

}

// Vst3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qP8(a *arm.Poly8, val [3]arm.Poly8x16)  {

}

// Vst3qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qS16(a *int16, val [3]arm.Int16x8)  {

}

// Vst3qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qP16(a *arm.Poly16, val [3]arm.Poly16x8)  {

}

// Vst3qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qS32(a *int32, val [3]arm.Int32x4)  {

}

// Vst3qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qS64(a *int64, val [3]arm.Int64x2)  {

}

// Vst3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qU8(a *uint8, val [3]arm.Uint8x16)  {

}

// Vst3qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qU16(a *uint16, val [3]arm.Uint16x8)  {

}

// Vst3qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qU32(a *uint32, val [3]arm.Uint32x4)  {

}

// Vst3qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qU64(a *uint64, val [3]arm.Uint64x2)  {

}

// Vst3qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qF32(a *float32, val [3]arm.Float32x4)  {

}

// Vst3qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst3qF64(a *float64, val [3]arm.Float64x2)  {

}

// Vst4S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4S64(a *int64, val [4]arm.Int64x1)  {

}

// Vst4U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4U64(a *uint64, val [4]arm.Uint64x1)  {

}

// Vst4F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4F64(a *float64, val [4]arm.Float64x1)  {

}

// Vst4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4S8(a *int8, val [4]arm.Int8x8)  {

}

// Vst4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4P8(a *arm.Poly8, val [4]arm.Poly8x8)  {

}

// Vst4S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4S16(a *int16, val [4]arm.Int16x4)  {

}

// Vst4P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4P16(a *arm.Poly16, val [4]arm.Poly16x4)  {

}

// Vst4S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4S32(a *int32, val [4]arm.Int32x2)  {

}

// Vst4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4U8(a *uint8, val [4]arm.Uint8x8)  {

}

// Vst4U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4U16(a *uint16, val [4]arm.Uint16x4)  {

}

// Vst4U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4U32(a *uint32, val [4]arm.Uint32x2)  {

}

// Vst4F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4F32(a *float32, val [4]arm.Float32x2)  {

}

// Vst4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qS8(a *int8, val [4]arm.Int8x16)  {

}

// Vst4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qP8(a *arm.Poly8, val [4]arm.Poly8x16)  {

}

// Vst4qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qS16(a *int16, val [4]arm.Int16x8)  {

}

// Vst4qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qP16(a *arm.Poly16, val [4]arm.Poly16x8)  {

}

// Vst4qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qS32(a *int32, val [4]arm.Int32x4)  {

}

// Vst4qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qS64(a *int64, val [4]arm.Int64x2)  {

}

// Vst4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qU8(a *uint8, val [4]arm.Uint8x16)  {

}

// Vst4qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qU16(a *uint16, val [4]arm.Uint16x8)  {

}

// Vst4qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qU32(a *uint32, val [4]arm.Uint32x4)  {

}

// Vst4qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qU64(a *uint64, val [4]arm.Uint64x2)  {

}

// Vst4qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qF32(a *float32, val [4]arm.Float32x4)  {

}

// Vst4qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Vst4qF64(a *float64, val [4]arm.Float64x2)  {

}

// VsubdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubd_s64'.
// Requires NEON.
func VsubdS64(a int64, b int64) int64 {
	return 0
}

// VsubdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubd_u64'.
// Requires NEON.
func VsubdU64(a uint64, b uint64) uint64 {
	return 0
}

// Vtbx1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx1_s8'.
// Requires NEON.
func Vtbx1S8(r arm.Int8x8, tab arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbx1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx1_u8'.
// Requires NEON.
func Vtbx1U8(r arm.Uint8x8, tab arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbx1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx1_p8'.
// Requires NEON.
func Vtbx1P8(r arm.Poly8x8, tab arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtbx3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx3_s8'.
// Requires NEON.
func Vtbx3S8(r arm.Int8x8, tab [3]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtbx3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx3_u8'.
// Requires NEON.
func Vtbx3U8(r arm.Uint8x8, tab [3]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtbx3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx3_p8'.
// Requires NEON.
func Vtbx3P8(r arm.Poly8x8, tab [3]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtrn1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_f32'.
// Requires NEON.
func Vtrn1F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vtrn1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_p8'.
// Requires NEON.
func Vtrn1P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtrn1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_p16'.
// Requires NEON.
func Vtrn1P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vtrn1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_s8'.
// Requires NEON.
func Vtrn1S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtrn1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_s16'.
// Requires NEON.
func Vtrn1S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vtrn1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_s32'.
// Requires NEON.
func Vtrn1S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vtrn1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_u8'.
// Requires NEON.
func Vtrn1U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtrn1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_u16'.
// Requires NEON.
func Vtrn1U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vtrn1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_u32'.
// Requires NEON.
func Vtrn1U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vtrn1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_f32'.
// Requires NEON.
func Vtrn1qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vtrn1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_f64'.
// Requires NEON.
func Vtrn1qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vtrn1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_p8'.
// Requires NEON.
func Vtrn1qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vtrn1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_p16'.
// Requires NEON.
func Vtrn1qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vtrn1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s8'.
// Requires NEON.
func Vtrn1qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vtrn1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s16'.
// Requires NEON.
func Vtrn1qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vtrn1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s32'.
// Requires NEON.
func Vtrn1qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vtrn1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s64'.
// Requires NEON.
func Vtrn1qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vtrn1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u8'.
// Requires NEON.
func Vtrn1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vtrn1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u16'.
// Requires NEON.
func Vtrn1qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vtrn1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u32'.
// Requires NEON.
func Vtrn1qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vtrn1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u64'.
// Requires NEON.
func Vtrn1qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Vtrn2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_f32'.
// Requires NEON.
func Vtrn2F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vtrn2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_p8'.
// Requires NEON.
func Vtrn2P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vtrn2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_p16'.
// Requires NEON.
func Vtrn2P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vtrn2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_s8'.
// Requires NEON.
func Vtrn2S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vtrn2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_s16'.
// Requires NEON.
func Vtrn2S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vtrn2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_s32'.
// Requires NEON.
func Vtrn2S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vtrn2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_u8'.
// Requires NEON.
func Vtrn2U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vtrn2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_u16'.
// Requires NEON.
func Vtrn2U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vtrn2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_u32'.
// Requires NEON.
func Vtrn2U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vtrn2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_f32'.
// Requires NEON.
func Vtrn2qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vtrn2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_f64'.
// Requires NEON.
func Vtrn2qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vtrn2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_p8'.
// Requires NEON.
func Vtrn2qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vtrn2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_p16'.
// Requires NEON.
func Vtrn2qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vtrn2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s8'.
// Requires NEON.
func Vtrn2qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vtrn2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s16'.
// Requires NEON.
func Vtrn2qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vtrn2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s32'.
// Requires NEON.
func Vtrn2qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vtrn2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s64'.
// Requires NEON.
func Vtrn2qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vtrn2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u8'.
// Requires NEON.
func Vtrn2qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vtrn2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u16'.
// Requires NEON.
func Vtrn2qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vtrn2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u32'.
// Requires NEON.
func Vtrn2qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vtrn2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u64'.
// Requires NEON.
func Vtrn2qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VtrnF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_f32'.
// Requires NEON.
func VtrnF32(a arm.Float32x2, b arm.Float32x2) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// VtrnP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_p8'.
// Requires NEON.
func VtrnP8(a arm.Poly8x8, b arm.Poly8x8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// VtrnP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_p16'.
// Requires NEON.
func VtrnP16(a arm.Poly16x4, b arm.Poly16x4) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// VtrnS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_s8'.
// Requires NEON.
func VtrnS8(a arm.Int8x8, b arm.Int8x8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// VtrnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_s16'.
// Requires NEON.
func VtrnS16(a arm.Int16x4, b arm.Int16x4) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// VtrnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_s32'.
// Requires NEON.
func VtrnS32(a arm.Int32x2, b arm.Int32x2) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// VtrnU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_u8'.
// Requires NEON.
func VtrnU8(a arm.Uint8x8, b arm.Uint8x8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// VtrnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_u16'.
// Requires NEON.
func VtrnU16(a arm.Uint16x4, b arm.Uint16x4) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// VtrnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_u32'.
// Requires NEON.
func VtrnU32(a arm.Uint32x2, b arm.Uint32x2) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// VtrnqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_f32'.
// Requires NEON.
func VtrnqF32(a arm.Float32x4, b arm.Float32x4) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// VtrnqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_p8'.
// Requires NEON.
func VtrnqP8(a arm.Poly8x16, b arm.Poly8x16) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// VtrnqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_p16'.
// Requires NEON.
func VtrnqP16(a arm.Poly16x8, b arm.Poly16x8) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// VtrnqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_s8'.
// Requires NEON.
func VtrnqS8(a arm.Int8x16, b arm.Int8x16) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// VtrnqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_s16'.
// Requires NEON.
func VtrnqS16(a arm.Int16x8, b arm.Int16x8) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// VtrnqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_s32'.
// Requires NEON.
func VtrnqS32(a arm.Int32x4, b arm.Int32x4) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// VtrnqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_u8'.
// Requires NEON.
func VtrnqU8(a arm.Uint8x16, b arm.Uint8x16) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// VtrnqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_u16'.
// Requires NEON.
func VtrnqU16(a arm.Uint16x8, b arm.Uint16x8) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// VtrnqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_u32'.
// Requires NEON.
func VtrnqU32(a arm.Uint32x4, b arm.Uint32x4) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// VtstS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s8'.
// Requires NEON.
func VtstS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VtstS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s16'.
// Requires NEON.
func VtstS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VtstS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s32'.
// Requires NEON.
func VtstS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VtstS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s64'.
// Requires NEON.
func VtstS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VtstU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u8'.
// Requires NEON.
func VtstU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// VtstU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u16'.
// Requires NEON.
func VtstU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// VtstU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u32'.
// Requires NEON.
func VtstU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// VtstU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u64'.
// Requires NEON.
func VtstU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// VtstqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s8'.
// Requires NEON.
func VtstqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VtstqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s16'.
// Requires NEON.
func VtstqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VtstqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s32'.
// Requires NEON.
func VtstqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VtstqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s64'.
// Requires NEON.
func VtstqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VtstqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u8'.
// Requires NEON.
func VtstqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// VtstqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u16'.
// Requires NEON.
func VtstqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// VtstqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u32'.
// Requires NEON.
func VtstqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// VtstqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u64'.
// Requires NEON.
func VtstqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VtstdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstd_s64'.
// Requires NEON.
func VtstdS64(a int64, b int64) uint64 {
	return 0
}

// VtstdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstd_u64'.
// Requires NEON.
func VtstdU64(a uint64, b uint64) uint64 {
	return 0
}

// VuqaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s8'.
// Requires NEON.
func VuqaddS8(a arm.Int8x8, b arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// VuqaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s16'.
// Requires NEON.
func VuqaddS16(a arm.Int16x4, b arm.Uint16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// VuqaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s32'.
// Requires NEON.
func VuqaddS32(a arm.Int32x2, b arm.Uint32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// VuqaddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s64'.
// Requires NEON.
func VuqaddS64(a arm.Int64x1, b arm.Uint64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// VuqaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s8'.
// Requires NEON.
func VuqaddqS8(a arm.Int8x16, b arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// VuqaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s16'.
// Requires NEON.
func VuqaddqS16(a arm.Int16x8, b arm.Uint16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// VuqaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s32'.
// Requires NEON.
func VuqaddqS32(a arm.Int32x4, b arm.Uint32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// VuqaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s64'.
// Requires NEON.
func VuqaddqS64(a arm.Int64x2, b arm.Uint64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// VuqaddbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddb_s8'.
// Requires NEON.
func VuqaddbS8(a int8, b uint8) int8 {
	return 0
}

// VuqaddhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddh_s16'.
// Requires NEON.
func VuqaddhS16(a int16, b uint16) int16 {
	return 0
}

// VuqaddsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadds_s32'.
// Requires NEON.
func VuqaddsS32(a int32, b uint32) int32 {
	return 0
}

// VuqadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddd_s64'.
// Requires NEON.
func VuqadddS64(a int64, b uint64) int64 {
	return 0
}

// Vuzp1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_f32'.
// Requires NEON.
func Vuzp1F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vuzp1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_p8'.
// Requires NEON.
func Vuzp1P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vuzp1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_p16'.
// Requires NEON.
func Vuzp1P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vuzp1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_s8'.
// Requires NEON.
func Vuzp1S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vuzp1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_s16'.
// Requires NEON.
func Vuzp1S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vuzp1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_s32'.
// Requires NEON.
func Vuzp1S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vuzp1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_u8'.
// Requires NEON.
func Vuzp1U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vuzp1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_u16'.
// Requires NEON.
func Vuzp1U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vuzp1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_u32'.
// Requires NEON.
func Vuzp1U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vuzp1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_f32'.
// Requires NEON.
func Vuzp1qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vuzp1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_f64'.
// Requires NEON.
func Vuzp1qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vuzp1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_p8'.
// Requires NEON.
func Vuzp1qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vuzp1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_p16'.
// Requires NEON.
func Vuzp1qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vuzp1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s8'.
// Requires NEON.
func Vuzp1qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vuzp1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s16'.
// Requires NEON.
func Vuzp1qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vuzp1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s32'.
// Requires NEON.
func Vuzp1qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vuzp1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s64'.
// Requires NEON.
func Vuzp1qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vuzp1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u8'.
// Requires NEON.
func Vuzp1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vuzp1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u16'.
// Requires NEON.
func Vuzp1qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vuzp1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u32'.
// Requires NEON.
func Vuzp1qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vuzp1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u64'.
// Requires NEON.
func Vuzp1qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Vuzp2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_f32'.
// Requires NEON.
func Vuzp2F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vuzp2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_p8'.
// Requires NEON.
func Vuzp2P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vuzp2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_p16'.
// Requires NEON.
func Vuzp2P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vuzp2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_s8'.
// Requires NEON.
func Vuzp2S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vuzp2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_s16'.
// Requires NEON.
func Vuzp2S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vuzp2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_s32'.
// Requires NEON.
func Vuzp2S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vuzp2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_u8'.
// Requires NEON.
func Vuzp2U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vuzp2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_u16'.
// Requires NEON.
func Vuzp2U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vuzp2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_u32'.
// Requires NEON.
func Vuzp2U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vuzp2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_f32'.
// Requires NEON.
func Vuzp2qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vuzp2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_f64'.
// Requires NEON.
func Vuzp2qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vuzp2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_p8'.
// Requires NEON.
func Vuzp2qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vuzp2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_p16'.
// Requires NEON.
func Vuzp2qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vuzp2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s8'.
// Requires NEON.
func Vuzp2qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vuzp2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s16'.
// Requires NEON.
func Vuzp2qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vuzp2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s32'.
// Requires NEON.
func Vuzp2qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vuzp2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s64'.
// Requires NEON.
func Vuzp2qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vuzp2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u8'.
// Requires NEON.
func Vuzp2qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vuzp2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u16'.
// Requires NEON.
func Vuzp2qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vuzp2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u32'.
// Requires NEON.
func Vuzp2qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vuzp2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u64'.
// Requires NEON.
func Vuzp2qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VuzpF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_f32'.
// Requires NEON.
func VuzpF32(a arm.Float32x2, b arm.Float32x2) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// VuzpP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_p8'.
// Requires NEON.
func VuzpP8(a arm.Poly8x8, b arm.Poly8x8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// VuzpP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_p16'.
// Requires NEON.
func VuzpP16(a arm.Poly16x4, b arm.Poly16x4) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// VuzpS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_s8'.
// Requires NEON.
func VuzpS8(a arm.Int8x8, b arm.Int8x8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// VuzpS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_s16'.
// Requires NEON.
func VuzpS16(a arm.Int16x4, b arm.Int16x4) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// VuzpS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_s32'.
// Requires NEON.
func VuzpS32(a arm.Int32x2, b arm.Int32x2) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// VuzpU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_u8'.
// Requires NEON.
func VuzpU8(a arm.Uint8x8, b arm.Uint8x8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// VuzpU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_u16'.
// Requires NEON.
func VuzpU16(a arm.Uint16x4, b arm.Uint16x4) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// VuzpU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_u32'.
// Requires NEON.
func VuzpU32(a arm.Uint32x2, b arm.Uint32x2) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// VuzpqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_f32'.
// Requires NEON.
func VuzpqF32(a arm.Float32x4, b arm.Float32x4) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// VuzpqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_p8'.
// Requires NEON.
func VuzpqP8(a arm.Poly8x16, b arm.Poly8x16) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// VuzpqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_p16'.
// Requires NEON.
func VuzpqP16(a arm.Poly16x8, b arm.Poly16x8) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// VuzpqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_s8'.
// Requires NEON.
func VuzpqS8(a arm.Int8x16, b arm.Int8x16) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// VuzpqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_s16'.
// Requires NEON.
func VuzpqS16(a arm.Int16x8, b arm.Int16x8) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// VuzpqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_s32'.
// Requires NEON.
func VuzpqS32(a arm.Int32x4, b arm.Int32x4) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// VuzpqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_u8'.
// Requires NEON.
func VuzpqU8(a arm.Uint8x16, b arm.Uint8x16) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// VuzpqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_u16'.
// Requires NEON.
func VuzpqU16(a arm.Uint16x8, b arm.Uint16x8) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// VuzpqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_u32'.
// Requires NEON.
func VuzpqU32(a arm.Uint32x4, b arm.Uint32x4) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Vzip1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_f32'.
// Requires NEON.
func Vzip1F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vzip1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_p8'.
// Requires NEON.
func Vzip1P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vzip1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_p16'.
// Requires NEON.
func Vzip1P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vzip1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_s8'.
// Requires NEON.
func Vzip1S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vzip1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_s16'.
// Requires NEON.
func Vzip1S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vzip1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_s32'.
// Requires NEON.
func Vzip1S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vzip1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_u8'.
// Requires NEON.
func Vzip1U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vzip1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_u16'.
// Requires NEON.
func Vzip1U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vzip1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_u32'.
// Requires NEON.
func Vzip1U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vzip1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_f32'.
// Requires NEON.
func Vzip1qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vzip1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_f64'.
// Requires NEON.
func Vzip1qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vzip1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_p8'.
// Requires NEON.
func Vzip1qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vzip1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_p16'.
// Requires NEON.
func Vzip1qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vzip1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s8'.
// Requires NEON.
func Vzip1qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vzip1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s16'.
// Requires NEON.
func Vzip1qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vzip1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s32'.
// Requires NEON.
func Vzip1qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vzip1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s64'.
// Requires NEON.
func Vzip1qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vzip1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u8'.
// Requires NEON.
func Vzip1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vzip1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u16'.
// Requires NEON.
func Vzip1qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vzip1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u32'.
// Requires NEON.
func Vzip1qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vzip1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u64'.
// Requires NEON.
func Vzip1qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Vzip2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_f32'.
// Requires NEON.
func Vzip2F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Vzip2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_p8'.
// Requires NEON.
func Vzip2P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Vzip2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_p16'.
// Requires NEON.
func Vzip2P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Vzip2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_s8'.
// Requires NEON.
func Vzip2S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Vzip2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_s16'.
// Requires NEON.
func Vzip2S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Vzip2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_s32'.
// Requires NEON.
func Vzip2S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Vzip2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_u8'.
// Requires NEON.
func Vzip2U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Vzip2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_u16'.
// Requires NEON.
func Vzip2U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Vzip2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_u32'.
// Requires NEON.
func Vzip2U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Vzip2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_f32'.
// Requires NEON.
func Vzip2qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Vzip2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_f64'.
// Requires NEON.
func Vzip2qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Vzip2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_p8'.
// Requires NEON.
func Vzip2qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Vzip2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_p16'.
// Requires NEON.
func Vzip2qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Vzip2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s8'.
// Requires NEON.
func Vzip2qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Vzip2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s16'.
// Requires NEON.
func Vzip2qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Vzip2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s32'.
// Requires NEON.
func Vzip2qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Vzip2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s64'.
// Requires NEON.
func Vzip2qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Vzip2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u8'.
// Requires NEON.
func Vzip2qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Vzip2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u16'.
// Requires NEON.
func Vzip2qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Vzip2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u32'.
// Requires NEON.
func Vzip2qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Vzip2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u64'.
// Requires NEON.
func Vzip2qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// VzipF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_f32'.
// Requires NEON.
func VzipF32(a arm.Float32x2, b arm.Float32x2) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// VzipP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_p8'.
// Requires NEON.
func VzipP8(a arm.Poly8x8, b arm.Poly8x8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// VzipP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_p16'.
// Requires NEON.
func VzipP16(a arm.Poly16x4, b arm.Poly16x4) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// VzipS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_s8'.
// Requires NEON.
func VzipS8(a arm.Int8x8, b arm.Int8x8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// VzipS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_s16'.
// Requires NEON.
func VzipS16(a arm.Int16x4, b arm.Int16x4) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// VzipS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_s32'.
// Requires NEON.
func VzipS32(a arm.Int32x2, b arm.Int32x2) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// VzipU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_u8'.
// Requires NEON.
func VzipU8(a arm.Uint8x8, b arm.Uint8x8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// VzipU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_u16'.
// Requires NEON.
func VzipU16(a arm.Uint16x4, b arm.Uint16x4) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// VzipU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_u32'.
// Requires NEON.
func VzipU32(a arm.Uint32x2, b arm.Uint32x2) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// VzipqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_f32'.
// Requires NEON.
func VzipqF32(a arm.Float32x4, b arm.Float32x4) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// VzipqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_p8'.
// Requires NEON.
func VzipqP8(a arm.Poly8x16, b arm.Poly8x16) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// VzipqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_p16'.
// Requires NEON.
func VzipqP16(a arm.Poly16x8, b arm.Poly16x8) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// VzipqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_s8'.
// Requires NEON.
func VzipqS8(a arm.Int8x16, b arm.Int8x16) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// VzipqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_s16'.
// Requires NEON.
func VzipqS16(a arm.Int16x8, b arm.Int16x8) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// VzipqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_s32'.
// Requires NEON.
func VzipqS32(a arm.Int32x4, b arm.Int32x4) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// VzipqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_u8'.
// Requires NEON.
func VzipqU8(a arm.Uint8x16, b arm.Uint8x16) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// VzipqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_u16'.
// Requires NEON.
func VzipqU16(a arm.Uint16x8, b arm.Uint16x8) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// VzipqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_u32'.
// Requires NEON.
func VzipqU32(a arm.Uint32x4, b arm.Uint32x4) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}
