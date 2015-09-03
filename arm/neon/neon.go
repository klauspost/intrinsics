package neon

import "github.com/klauspost/intrinsics/arm"

var _ = arm.Int8x8{}  // Make sure we use arm package


// AddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s8'.
// Requires NEON.
func AddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// AddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s16'.
// Requires NEON.
func AddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// AddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s32'.
// Requires NEON.
func AddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// AddF32: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_f32'.
// Requires NEON.
func AddF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// AddF64: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_f64'.
// Requires NEON.
func AddF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// AddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u8'.
// Requires NEON.
func AddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// AddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u16'.
// Requires NEON.
func AddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// AddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u32'.
// Requires NEON.
func AddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// AddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_s64'.
// Requires NEON.
func AddS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// AddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vadd_u64'.
// Requires NEON.
func AddU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// AddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s8'.
// Requires NEON.
func AddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// AddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s16'.
// Requires NEON.
func AddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s32'.
// Requires NEON.
func AddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_s64'.
// Requires NEON.
func AddqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AddqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_f32'.
// Requires NEON.
func AddqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// AddqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_f64'.
// Requires NEON.
func AddqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// AddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u8'.
// Requires NEON.
func AddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u16'.
// Requires NEON.
func AddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u32'.
// Requires NEON.
func AddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddq_u64'.
// Requires NEON.
func AddqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AddlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_s8'.
// Requires NEON.
func AddlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AddlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_s16'.
// Requires NEON.
func AddlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AddlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_s32'.
// Requires NEON.
func AddlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AddlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_u8'.
// Requires NEON.
func AddlU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AddlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_u16'.
// Requires NEON.
func AddlU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AddlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_u32'.
// Requires NEON.
func AddlU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AddlHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_s8'.
// Requires NEON.
func AddlHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AddlHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_s16'.
// Requires NEON.
func AddlHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AddlHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_s32'.
// Requires NEON.
func AddlHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AddlHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_u8'.
// Requires NEON.
func AddlHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AddlHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_u16'.
// Requires NEON.
func AddlHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AddlHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddl_high_u32'.
// Requires NEON.
func AddlHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AddwS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_s8'.
// Requires NEON.
func AddwS8(a arm.Int16x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AddwS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_s16'.
// Requires NEON.
func AddwS16(a arm.Int32x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AddwS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_s32'.
// Requires NEON.
func AddwS32(a arm.Int64x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AddwU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_u8'.
// Requires NEON.
func AddwU8(a arm.Uint16x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AddwU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_u16'.
// Requires NEON.
func AddwU16(a arm.Uint32x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AddwU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_u32'.
// Requires NEON.
func AddwU32(a arm.Uint64x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AddwHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_s8'.
// Requires NEON.
func AddwHighS8(a arm.Int16x8, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AddwHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_s16'.
// Requires NEON.
func AddwHighS16(a arm.Int32x4, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AddwHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_s32'.
// Requires NEON.
func AddwHighS32(a arm.Int64x2, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AddwHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_u8'.
// Requires NEON.
func AddwHighU8(a arm.Uint16x8, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AddwHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_u16'.
// Requires NEON.
func AddwHighU16(a arm.Uint32x4, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AddwHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddw_high_u32'.
// Requires NEON.
func AddwHighU32(a arm.Uint64x2, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// HaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_s8'.
// Requires NEON.
func HaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// HaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_s16'.
// Requires NEON.
func HaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// HaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_s32'.
// Requires NEON.
func HaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// HaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_u8'.
// Requires NEON.
func HaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// HaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_u16'.
// Requires NEON.
func HaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// HaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhadd_u32'.
// Requires NEON.
func HaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// HaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_s8'.
// Requires NEON.
func HaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// HaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_s16'.
// Requires NEON.
func HaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// HaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_s32'.
// Requires NEON.
func HaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// HaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_u8'.
// Requires NEON.
func HaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// HaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_u16'.
// Requires NEON.
func HaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// HaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhaddq_u32'.
// Requires NEON.
func HaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RhaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_s8'.
// Requires NEON.
func RhaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RhaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_s16'.
// Requires NEON.
func RhaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// RhaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_s32'.
// Requires NEON.
func RhaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// RhaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_u8'.
// Requires NEON.
func RhaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RhaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_u16'.
// Requires NEON.
func RhaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// RhaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhadd_u32'.
// Requires NEON.
func RhaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RhaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_s8'.
// Requires NEON.
func RhaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RhaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_s16'.
// Requires NEON.
func RhaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// RhaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_s32'.
// Requires NEON.
func RhaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RhaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_u8'.
// Requires NEON.
func RhaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RhaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_u16'.
// Requires NEON.
func RhaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// RhaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrhaddq_u32'.
// Requires NEON.
func RhaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AddhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_s16'.
// Requires NEON.
func AddhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// AddhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_s32'.
// Requires NEON.
func AddhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// AddhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_s64'.
// Requires NEON.
func AddhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// AddhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_u16'.
// Requires NEON.
func AddhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// AddhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_u32'.
// Requires NEON.
func AddhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// AddhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_u64'.
// Requires NEON.
func AddhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RaddhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_s16'.
// Requires NEON.
func RaddhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RaddhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_s32'.
// Requires NEON.
func RaddhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// RaddhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_s64'.
// Requires NEON.
func RaddhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// RaddhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_u16'.
// Requires NEON.
func RaddhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RaddhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_u32'.
// Requires NEON.
func RaddhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// RaddhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_u64'.
// Requires NEON.
func RaddhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// AddhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_s16'.
// Requires NEON.
func AddhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// AddhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_s32'.
// Requires NEON.
func AddhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AddhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_s64'.
// Requires NEON.
func AddhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AddhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_u16'.
// Requires NEON.
func AddhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AddhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_u32'.
// Requires NEON.
func AddhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AddhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddhn_high_u64'.
// Requires NEON.
func AddhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RaddhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_s16'.
// Requires NEON.
func RaddhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RaddhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_s32'.
// Requires NEON.
func RaddhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// RaddhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_s64'.
// Requires NEON.
func RaddhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RaddhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_u16'.
// Requires NEON.
func RaddhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RaddhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_u32'.
// Requires NEON.
func RaddhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// RaddhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vraddhn_high_u64'.
// Requires NEON.
func RaddhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// DivF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdiv_f32'.
// Requires NEON.
func DivF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// DivF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdiv_f64'.
// Requires NEON.
func DivF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// DivqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdivq_f32'.
// Requires NEON.
func DivqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// DivqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdivq_f64'.
// Requires NEON.
func DivqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MulS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_s8'.
// Requires NEON.
func MulS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MulS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_s16'.
// Requires NEON.
func MulS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MulS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_s32'.
// Requires NEON.
func MulS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MulF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_f32'.
// Requires NEON.
func MulF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MulF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_f64'.
// Requires NEON.
func MulF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MulU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_u8'.
// Requires NEON.
func MulU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MulU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_u16'.
// Requires NEON.
func MulU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MulU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_u32'.
// Requires NEON.
func MulU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MulP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_p8'.
// Requires NEON.
func MulP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// MulqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_s8'.
// Requires NEON.
func MulqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MulqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_s16'.
// Requires NEON.
func MulqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MulqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_s32'.
// Requires NEON.
func MulqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MulqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_f32'.
// Requires NEON.
func MulqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MulqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_f64'.
// Requires NEON.
func MulqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MulqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_u8'.
// Requires NEON.
func MulqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MulqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_u16'.
// Requires NEON.
func MulqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MulqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_u32'.
// Requires NEON.
func MulqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MulqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_p8'.
// Requires NEON.
func MulqP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// AndS8: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s8'.
// Requires NEON.
func AndS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// AndS16: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s16'.
// Requires NEON.
func AndS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// AndS32: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s32'.
// Requires NEON.
func AndS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// AndU8: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u8'.
// Requires NEON.
func AndU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// AndU16: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u16'.
// Requires NEON.
func AndU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// AndU32: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u32'.
// Requires NEON.
func AndU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// AndS64: ARM NEON intrinsic 
//
// Intrinsic: 'vand_s64'.
// Requires NEON.
func AndS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// AndU64: ARM NEON intrinsic 
//
// Intrinsic: 'vand_u64'.
// Requires NEON.
func AndU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// AndqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s8'.
// Requires NEON.
func AndqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// AndqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s16'.
// Requires NEON.
func AndqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AndqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s32'.
// Requires NEON.
func AndqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AndqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_s64'.
// Requires NEON.
func AndqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AndqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u8'.
// Requires NEON.
func AndqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AndqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u16'.
// Requires NEON.
func AndqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AndqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u32'.
// Requires NEON.
func AndqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AndqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vandq_u64'.
// Requires NEON.
func AndqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// OrrS8: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s8'.
// Requires NEON.
func OrrS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// OrrS16: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s16'.
// Requires NEON.
func OrrS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// OrrS32: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s32'.
// Requires NEON.
func OrrS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// OrrU8: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u8'.
// Requires NEON.
func OrrU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// OrrU16: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u16'.
// Requires NEON.
func OrrU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// OrrU32: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u32'.
// Requires NEON.
func OrrU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// OrrS64: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_s64'.
// Requires NEON.
func OrrS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// OrrU64: ARM NEON intrinsic 
//
// Intrinsic: 'vorr_u64'.
// Requires NEON.
func OrrU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// OrrqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s8'.
// Requires NEON.
func OrrqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// OrrqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s16'.
// Requires NEON.
func OrrqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// OrrqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s32'.
// Requires NEON.
func OrrqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// OrrqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_s64'.
// Requires NEON.
func OrrqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// OrrqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u8'.
// Requires NEON.
func OrrqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// OrrqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u16'.
// Requires NEON.
func OrrqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// OrrqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u32'.
// Requires NEON.
func OrrqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// OrrqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vorrq_u64'.
// Requires NEON.
func OrrqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// EorS8: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s8'.
// Requires NEON.
func EorS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// EorS16: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s16'.
// Requires NEON.
func EorS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// EorS32: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s32'.
// Requires NEON.
func EorS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// EorU8: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u8'.
// Requires NEON.
func EorU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// EorU16: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u16'.
// Requires NEON.
func EorU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// EorU32: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u32'.
// Requires NEON.
func EorU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// EorS64: ARM NEON intrinsic 
//
// Intrinsic: 'veor_s64'.
// Requires NEON.
func EorS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// EorU64: ARM NEON intrinsic 
//
// Intrinsic: 'veor_u64'.
// Requires NEON.
func EorU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// EorqS8: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s8'.
// Requires NEON.
func EorqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// EorqS16: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s16'.
// Requires NEON.
func EorqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// EorqS32: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s32'.
// Requires NEON.
func EorqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// EorqS64: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_s64'.
// Requires NEON.
func EorqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// EorqU8: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u8'.
// Requires NEON.
func EorqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// EorqU16: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u16'.
// Requires NEON.
func EorqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// EorqU32: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u32'.
// Requires NEON.
func EorqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// EorqU64: ARM NEON intrinsic 
//
// Intrinsic: 'veorq_u64'.
// Requires NEON.
func EorqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// BicS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s8'.
// Requires NEON.
func BicS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// BicS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s16'.
// Requires NEON.
func BicS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// BicS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s32'.
// Requires NEON.
func BicS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// BicU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u8'.
// Requires NEON.
func BicU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// BicU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u16'.
// Requires NEON.
func BicU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// BicU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u32'.
// Requires NEON.
func BicU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// BicS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_s64'.
// Requires NEON.
func BicS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// BicU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbic_u64'.
// Requires NEON.
func BicU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// BicqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s8'.
// Requires NEON.
func BicqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// BicqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s16'.
// Requires NEON.
func BicqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// BicqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s32'.
// Requires NEON.
func BicqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// BicqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_s64'.
// Requires NEON.
func BicqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// BicqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u8'.
// Requires NEON.
func BicqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// BicqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u16'.
// Requires NEON.
func BicqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// BicqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u32'.
// Requires NEON.
func BicqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// BicqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbicq_u64'.
// Requires NEON.
func BicqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// OrnS8: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s8'.
// Requires NEON.
func OrnS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// OrnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s16'.
// Requires NEON.
func OrnS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// OrnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s32'.
// Requires NEON.
func OrnS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// OrnU8: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u8'.
// Requires NEON.
func OrnU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// OrnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u16'.
// Requires NEON.
func OrnU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// OrnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u32'.
// Requires NEON.
func OrnU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// OrnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_s64'.
// Requires NEON.
func OrnS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// OrnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vorn_u64'.
// Requires NEON.
func OrnU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// OrnqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s8'.
// Requires NEON.
func OrnqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// OrnqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s16'.
// Requires NEON.
func OrnqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// OrnqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s32'.
// Requires NEON.
func OrnqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// OrnqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_s64'.
// Requires NEON.
func OrnqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// OrnqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u8'.
// Requires NEON.
func OrnqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// OrnqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u16'.
// Requires NEON.
func OrnqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// OrnqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u32'.
// Requires NEON.
func OrnqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// OrnqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vornq_u64'.
// Requires NEON.
func OrnqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SubS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s8'.
// Requires NEON.
func SubS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// SubS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s16'.
// Requires NEON.
func SubS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// SubS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s32'.
// Requires NEON.
func SubS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// SubF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_f32'.
// Requires NEON.
func SubF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// SubF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_f64'.
// Requires NEON.
func SubF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// SubU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u8'.
// Requires NEON.
func SubU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SubU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u16'.
// Requires NEON.
func SubU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SubU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u32'.
// Requires NEON.
func SubU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// SubS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_s64'.
// Requires NEON.
func SubS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// SubU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsub_u64'.
// Requires NEON.
func SubU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// SubqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s8'.
// Requires NEON.
func SubqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// SubqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s16'.
// Requires NEON.
func SubqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SubqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s32'.
// Requires NEON.
func SubqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SubqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_s64'.
// Requires NEON.
func SubqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SubqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_f32'.
// Requires NEON.
func SubqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// SubqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_f64'.
// Requires NEON.
func SubqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// SubqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u8'.
// Requires NEON.
func SubqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SubqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u16'.
// Requires NEON.
func SubqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SubqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u32'.
// Requires NEON.
func SubqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SubqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubq_u64'.
// Requires NEON.
func SubqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SublS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_s8'.
// Requires NEON.
func SublS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SublS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_s16'.
// Requires NEON.
func SublS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SublS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_s32'.
// Requires NEON.
func SublS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SublU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_u8'.
// Requires NEON.
func SublU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SublU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_u16'.
// Requires NEON.
func SublU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SublU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_u32'.
// Requires NEON.
func SublU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SublHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_s8'.
// Requires NEON.
func SublHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SublHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_s16'.
// Requires NEON.
func SublHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SublHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_s32'.
// Requires NEON.
func SublHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SublHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_u8'.
// Requires NEON.
func SublHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SublHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_u16'.
// Requires NEON.
func SublHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SublHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubl_high_u32'.
// Requires NEON.
func SublHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SubwS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_s8'.
// Requires NEON.
func SubwS8(a arm.Int16x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SubwS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_s16'.
// Requires NEON.
func SubwS16(a arm.Int32x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SubwS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_s32'.
// Requires NEON.
func SubwS32(a arm.Int64x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SubwU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_u8'.
// Requires NEON.
func SubwU8(a arm.Uint16x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SubwU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_u16'.
// Requires NEON.
func SubwU16(a arm.Uint32x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SubwU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_u32'.
// Requires NEON.
func SubwU32(a arm.Uint64x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SubwHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_s8'.
// Requires NEON.
func SubwHighS8(a arm.Int16x8, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SubwHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_s16'.
// Requires NEON.
func SubwHighS16(a arm.Int32x4, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SubwHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_s32'.
// Requires NEON.
func SubwHighS32(a arm.Int64x2, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SubwHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_u8'.
// Requires NEON.
func SubwHighU8(a arm.Uint16x8, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SubwHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_u16'.
// Requires NEON.
func SubwHighU16(a arm.Uint32x4, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SubwHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubw_high_u32'.
// Requires NEON.
func SubwHighU32(a arm.Uint64x2, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s8'.
// Requires NEON.
func QaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s16'.
// Requires NEON.
func QaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s32'.
// Requires NEON.
func QaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QaddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_s64'.
// Requires NEON.
func QaddS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u8'.
// Requires NEON.
func QaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// HsubS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_s8'.
// Requires NEON.
func HsubS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// HsubS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_s16'.
// Requires NEON.
func HsubS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// HsubS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_s32'.
// Requires NEON.
func HsubS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// HsubU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_u8'.
// Requires NEON.
func HsubU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// HsubU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_u16'.
// Requires NEON.
func HsubU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// HsubU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsub_u32'.
// Requires NEON.
func HsubU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// HsubqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_s8'.
// Requires NEON.
func HsubqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// HsubqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_s16'.
// Requires NEON.
func HsubqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// HsubqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_s32'.
// Requires NEON.
func HsubqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// HsubqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_u8'.
// Requires NEON.
func HsubqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// HsubqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_u16'.
// Requires NEON.
func HsubqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// HsubqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vhsubq_u32'.
// Requires NEON.
func HsubqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SubhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_s16'.
// Requires NEON.
func SubhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// SubhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_s32'.
// Requires NEON.
func SubhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// SubhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_s64'.
// Requires NEON.
func SubhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// SubhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_u16'.
// Requires NEON.
func SubhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SubhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_u32'.
// Requires NEON.
func SubhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SubhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_u64'.
// Requires NEON.
func SubhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RsubhnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_s16'.
// Requires NEON.
func RsubhnS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RsubhnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_s32'.
// Requires NEON.
func RsubhnS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// RsubhnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_s64'.
// Requires NEON.
func RsubhnS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// RsubhnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_u16'.
// Requires NEON.
func RsubhnU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RsubhnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_u32'.
// Requires NEON.
func RsubhnU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// RsubhnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_u64'.
// Requires NEON.
func RsubhnU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RsubhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_s16'.
// Requires NEON.
func RsubhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RsubhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_s32'.
// Requires NEON.
func RsubhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// RsubhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_s64'.
// Requires NEON.
func RsubhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RsubhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_u16'.
// Requires NEON.
func RsubhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RsubhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_u32'.
// Requires NEON.
func RsubhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// RsubhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsubhn_high_u64'.
// Requires NEON.
func RsubhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SubhnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_s16'.
// Requires NEON.
func SubhnHighS16(a arm.Int8x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// SubhnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_s32'.
// Requires NEON.
func SubhnHighS32(a arm.Int16x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SubhnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_s64'.
// Requires NEON.
func SubhnHighS64(a arm.Int32x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SubhnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_u16'.
// Requires NEON.
func SubhnHighU16(a arm.Uint8x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SubhnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_u32'.
// Requires NEON.
func SubhnHighU32(a arm.Uint16x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SubhnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubhn_high_u64'.
// Requires NEON.
func SubhnHighU64(a arm.Uint32x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u16'.
// Requires NEON.
func QaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u32'.
// Requires NEON.
func QaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QaddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqadd_u64'.
// Requires NEON.
func QaddU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// QaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s8'.
// Requires NEON.
func QaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s16'.
// Requires NEON.
func QaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s32'.
// Requires NEON.
func QaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_s64'.
// Requires NEON.
func QaddqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u8'.
// Requires NEON.
func QaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u16'.
// Requires NEON.
func QaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u32'.
// Requires NEON.
func QaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddq_u64'.
// Requires NEON.
func QaddqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QsubS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s8'.
// Requires NEON.
func QsubS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QsubS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s16'.
// Requires NEON.
func QsubS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QsubS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s32'.
// Requires NEON.
func QsubS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QsubS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_s64'.
// Requires NEON.
func QsubS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QsubU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u8'.
// Requires NEON.
func QsubU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QsubU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u16'.
// Requires NEON.
func QsubU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QsubU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u32'.
// Requires NEON.
func QsubU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QsubU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsub_u64'.
// Requires NEON.
func QsubU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// QsubqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s8'.
// Requires NEON.
func QsubqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QsubqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s16'.
// Requires NEON.
func QsubqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QsubqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s32'.
// Requires NEON.
func QsubqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QsubqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_s64'.
// Requires NEON.
func QsubqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QsubqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u8'.
// Requires NEON.
func QsubqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QsubqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u16'.
// Requires NEON.
func QsubqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QsubqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u32'.
// Requires NEON.
func QsubqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QsubqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubq_u64'.
// Requires NEON.
func QsubqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QnegS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s8'.
// Requires NEON.
func QnegS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QnegS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s16'.
// Requires NEON.
func QnegS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QnegS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s32'.
// Requires NEON.
func QnegS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QnegS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqneg_s64'.
// Requires NEON.
func QnegS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QnegqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s8'.
// Requires NEON.
func QnegqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QnegqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s16'.
// Requires NEON.
func QnegqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QnegqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s32'.
// Requires NEON.
func QnegqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QabsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s8'.
// Requires NEON.
func QabsS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QabsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s16'.
// Requires NEON.
func QabsS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QabsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s32'.
// Requires NEON.
func QabsS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QabsS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqabs_s64'.
// Requires NEON.
func QabsS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QabsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s8'.
// Requires NEON.
func QabsqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QabsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s16'.
// Requires NEON.
func QabsqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QabsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s32'.
// Requires NEON.
func QabsqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmulhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_s16'.
// Requires NEON.
func QdmulhS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QdmulhS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_s32'.
// Requires NEON.
func QdmulhS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QdmulhqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_s16'.
// Requires NEON.
func QdmulhqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QdmulhqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_s32'.
// Requires NEON.
func QdmulhqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QrdmulhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_s16'.
// Requires NEON.
func QrdmulhS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QrdmulhS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_s32'.
// Requires NEON.
func QrdmulhS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QrdmulhqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_s16'.
// Requires NEON.
func QrdmulhqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QrdmulhqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_s32'.
// Requires NEON.
func QrdmulhqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CreateS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s8'.
// Requires NEON.
func CreateS8(a uint64) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// CreateS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s16'.
// Requires NEON.
func CreateS16(a uint64) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// CreateS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s32'.
// Requires NEON.
func CreateS32(a uint64) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// CreateS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_s64'.
// Requires NEON.
func CreateS64(a uint64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// CreateF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_f32'.
// Requires NEON.
func CreateF32(a uint64) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// CreateU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u8'.
// Requires NEON.
func CreateU8(a uint64) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CreateU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u16'.
// Requires NEON.
func CreateU16(a uint64) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CreateU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u32'.
// Requires NEON.
func CreateU32(a uint64) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CreateU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_u64'.
// Requires NEON.
func CreateU64(a uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CreateF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_f64'.
// Requires NEON.
func CreateF64(a uint64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// CreateP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_p8'.
// Requires NEON.
func CreateP8(a uint64) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// CreateP16: ARM NEON intrinsic 
//
// Intrinsic: 'vcreate_p16'.
// Requires NEON.
func CreateP16(a uint64) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// GetLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneF32(a arm.Float32x2, b int) float32 {
	return 0
}

// GetLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneF64(a arm.Float64x1, b int) float64 {
	return 0
}

// GetLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// GetLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// GetLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneS8(a arm.Int8x8, b int) int8 {
	return 0
}

// GetLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneS16(a arm.Int16x4, b int) int16 {
	return 0
}

// GetLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneS32(a arm.Int32x2, b int) int32 {
	return 0
}

// GetLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneS64(a arm.Int64x1, b int) int64 {
	return 0
}

// GetLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneU8(a arm.Uint8x8, b int) uint8 {
	return 0
}

// GetLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneU16(a arm.Uint16x4, b int) uint16 {
	return 0
}

// GetLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneU32(a arm.Uint32x2, b int) uint32 {
	return 0
}

// GetLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetLaneU64(a arm.Uint64x1, b int) uint64 {
	return 0
}

// GetqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneF32(a arm.Float32x4, b int) float32 {
	return 0
}

// GetqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneF64(a arm.Float64x2, b int) float64 {
	return 0
}

// GetqLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneP8(a arm.Poly8x16, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// GetqLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneP16(a arm.Poly16x8, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// GetqLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneS8(a arm.Int8x16, b int) int8 {
	return 0
}

// GetqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneS16(a arm.Int16x8, b int) int16 {
	return 0
}

// GetqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneS32(a arm.Int32x4, b int) int32 {
	return 0
}

// GetqLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneS64(a arm.Int64x2, b int) int64 {
	return 0
}

// GetqLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneU8(a arm.Uint8x16, b int) uint8 {
	return 0
}

// GetqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneU16(a arm.Uint16x8, b int) uint16 {
	return 0
}

// GetqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneU32(a arm.Uint32x4, b int) uint32 {
	return 0
}

// GetqLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vgetq_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func GetqLaneU64(a arm.Uint64x2, b int) uint64 {
	return 0
}

// ReinterpretP8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_f64'.
// Requires NEON.
func ReinterpretP8F64(a arm.Float64x1) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s8'.
// Requires NEON.
func ReinterpretP8S8(a arm.Int8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s16'.
// Requires NEON.
func ReinterpretP8S16(a arm.Int16x4) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s32'.
// Requires NEON.
func ReinterpretP8S32(a arm.Int32x2) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_s64'.
// Requires NEON.
func ReinterpretP8S64(a arm.Int64x1) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_f32'.
// Requires NEON.
func ReinterpretP8F32(a arm.Float32x2) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u8'.
// Requires NEON.
func ReinterpretP8U8(a arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u16'.
// Requires NEON.
func ReinterpretP8U16(a arm.Uint16x4) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u32'.
// Requires NEON.
func ReinterpretP8U32(a arm.Uint32x2) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_u64'.
// Requires NEON.
func ReinterpretP8U64(a arm.Uint64x1) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretP8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p8_p16'.
// Requires NEON.
func ReinterpretP8P16(a arm.Poly16x4) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ReinterpretqP8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_f64'.
// Requires NEON.
func ReinterpretqP8F64(a arm.Float64x2) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s8'.
// Requires NEON.
func ReinterpretqP8S8(a arm.Int8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s16'.
// Requires NEON.
func ReinterpretqP8S16(a arm.Int16x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s32'.
// Requires NEON.
func ReinterpretqP8S32(a arm.Int32x4) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_s64'.
// Requires NEON.
func ReinterpretqP8S64(a arm.Int64x2) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_f32'.
// Requires NEON.
func ReinterpretqP8F32(a arm.Float32x4) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u8'.
// Requires NEON.
func ReinterpretqP8U8(a arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u16'.
// Requires NEON.
func ReinterpretqP8U16(a arm.Uint16x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u32'.
// Requires NEON.
func ReinterpretqP8U32(a arm.Uint32x4) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_u64'.
// Requires NEON.
func ReinterpretqP8U64(a arm.Uint64x2) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretqP8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p8_p16'.
// Requires NEON.
func ReinterpretqP8P16(a arm.Poly16x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ReinterpretP16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_f64'.
// Requires NEON.
func ReinterpretP16F64(a arm.Float64x1) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s8'.
// Requires NEON.
func ReinterpretP16S8(a arm.Int8x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s16'.
// Requires NEON.
func ReinterpretP16S16(a arm.Int16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s32'.
// Requires NEON.
func ReinterpretP16S32(a arm.Int32x2) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_s64'.
// Requires NEON.
func ReinterpretP16S64(a arm.Int64x1) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_f32'.
// Requires NEON.
func ReinterpretP16F32(a arm.Float32x2) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u8'.
// Requires NEON.
func ReinterpretP16U8(a arm.Uint8x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u16'.
// Requires NEON.
func ReinterpretP16U16(a arm.Uint16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u32'.
// Requires NEON.
func ReinterpretP16U32(a arm.Uint32x2) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_u64'.
// Requires NEON.
func ReinterpretP16U64(a arm.Uint64x1) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretP16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_p16_p8'.
// Requires NEON.
func ReinterpretP16P8(a arm.Poly8x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ReinterpretqP16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_f64'.
// Requires NEON.
func ReinterpretqP16F64(a arm.Float64x2) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s8'.
// Requires NEON.
func ReinterpretqP16S8(a arm.Int8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s16'.
// Requires NEON.
func ReinterpretqP16S16(a arm.Int16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s32'.
// Requires NEON.
func ReinterpretqP16S32(a arm.Int32x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_s64'.
// Requires NEON.
func ReinterpretqP16S64(a arm.Int64x2) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_f32'.
// Requires NEON.
func ReinterpretqP16F32(a arm.Float32x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u8'.
// Requires NEON.
func ReinterpretqP16U8(a arm.Uint8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u16'.
// Requires NEON.
func ReinterpretqP16U16(a arm.Uint16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u32'.
// Requires NEON.
func ReinterpretqP16U32(a arm.Uint32x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_u64'.
// Requires NEON.
func ReinterpretqP16U64(a arm.Uint64x2) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretqP16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_p16_p8'.
// Requires NEON.
func ReinterpretqP16P8(a arm.Poly8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ReinterpretF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_f64'.
// Requires NEON.
func ReinterpretF32F64(a arm.Float64x1) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s8'.
// Requires NEON.
func ReinterpretF32S8(a arm.Int8x8) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s16'.
// Requires NEON.
func ReinterpretF32S16(a arm.Int16x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s32'.
// Requires NEON.
func ReinterpretF32S32(a arm.Int32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_s64'.
// Requires NEON.
func ReinterpretF32S64(a arm.Int64x1) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u8'.
// Requires NEON.
func ReinterpretF32U8(a arm.Uint8x8) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u16'.
// Requires NEON.
func ReinterpretF32U16(a arm.Uint16x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u32'.
// Requires NEON.
func ReinterpretF32U32(a arm.Uint32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_u64'.
// Requires NEON.
func ReinterpretF32U64(a arm.Uint64x1) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_p8'.
// Requires NEON.
func ReinterpretF32P8(a arm.Poly8x8) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretF32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f32_p16'.
// Requires NEON.
func ReinterpretF32P16(a arm.Poly16x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ReinterpretqF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_f64'.
// Requires NEON.
func ReinterpretqF32F64(a arm.Float64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s8'.
// Requires NEON.
func ReinterpretqF32S8(a arm.Int8x16) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s16'.
// Requires NEON.
func ReinterpretqF32S16(a arm.Int16x8) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s32'.
// Requires NEON.
func ReinterpretqF32S32(a arm.Int32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_s64'.
// Requires NEON.
func ReinterpretqF32S64(a arm.Int64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u8'.
// Requires NEON.
func ReinterpretqF32U8(a arm.Uint8x16) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u16'.
// Requires NEON.
func ReinterpretqF32U16(a arm.Uint16x8) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u32'.
// Requires NEON.
func ReinterpretqF32U32(a arm.Uint32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_u64'.
// Requires NEON.
func ReinterpretqF32U64(a arm.Uint64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_p8'.
// Requires NEON.
func ReinterpretqF32P8(a arm.Poly8x16) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretqF32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f32_p16'.
// Requires NEON.
func ReinterpretqF32P16(a arm.Poly16x8) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ReinterpretF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_f32'.
// Requires NEON.
func ReinterpretF64F32(a arm.Float32x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_p8'.
// Requires NEON.
func ReinterpretF64P8(a arm.Poly8x8) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_p16'.
// Requires NEON.
func ReinterpretF64P16(a arm.Poly16x4) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s8'.
// Requires NEON.
func ReinterpretF64S8(a arm.Int8x8) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s16'.
// Requires NEON.
func ReinterpretF64S16(a arm.Int16x4) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s32'.
// Requires NEON.
func ReinterpretF64S32(a arm.Int32x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_s64'.
// Requires NEON.
func ReinterpretF64S64(a arm.Int64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u8'.
// Requires NEON.
func ReinterpretF64U8(a arm.Uint8x8) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u16'.
// Requires NEON.
func ReinterpretF64U16(a arm.Uint16x4) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u32'.
// Requires NEON.
func ReinterpretF64U32(a arm.Uint32x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_f64_u64'.
// Requires NEON.
func ReinterpretF64U64(a arm.Uint64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ReinterpretqF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_f32'.
// Requires NEON.
func ReinterpretqF64F32(a arm.Float32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_p8'.
// Requires NEON.
func ReinterpretqF64P8(a arm.Poly8x16) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_p16'.
// Requires NEON.
func ReinterpretqF64P16(a arm.Poly16x8) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s8'.
// Requires NEON.
func ReinterpretqF64S8(a arm.Int8x16) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s16'.
// Requires NEON.
func ReinterpretqF64S16(a arm.Int16x8) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s32'.
// Requires NEON.
func ReinterpretqF64S32(a arm.Int32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_s64'.
// Requires NEON.
func ReinterpretqF64S64(a arm.Int64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u8'.
// Requires NEON.
func ReinterpretqF64U8(a arm.Uint8x16) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u16'.
// Requires NEON.
func ReinterpretqF64U16(a arm.Uint16x8) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u32'.
// Requires NEON.
func ReinterpretqF64U32(a arm.Uint32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretqF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_f64_u64'.
// Requires NEON.
func ReinterpretqF64U64(a arm.Uint64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ReinterpretS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_f64'.
// Requires NEON.
func ReinterpretS64F64(a arm.Float64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_s8'.
// Requires NEON.
func ReinterpretS64S8(a arm.Int8x8) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_s16'.
// Requires NEON.
func ReinterpretS64S16(a arm.Int16x4) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_s32'.
// Requires NEON.
func ReinterpretS64S32(a arm.Int32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_f32'.
// Requires NEON.
func ReinterpretS64F32(a arm.Float32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u8'.
// Requires NEON.
func ReinterpretS64U8(a arm.Uint8x8) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u16'.
// Requires NEON.
func ReinterpretS64U16(a arm.Uint16x4) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u32'.
// Requires NEON.
func ReinterpretS64U32(a arm.Uint32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_u64'.
// Requires NEON.
func ReinterpretS64U64(a arm.Uint64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_p8'.
// Requires NEON.
func ReinterpretS64P8(a arm.Poly8x8) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretS64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s64_p16'.
// Requires NEON.
func ReinterpretS64P16(a arm.Poly16x4) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ReinterpretqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_f64'.
// Requires NEON.
func ReinterpretqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_s8'.
// Requires NEON.
func ReinterpretqS64S8(a arm.Int8x16) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_s16'.
// Requires NEON.
func ReinterpretqS64S16(a arm.Int16x8) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_s32'.
// Requires NEON.
func ReinterpretqS64S32(a arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_f32'.
// Requires NEON.
func ReinterpretqS64F32(a arm.Float32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u8'.
// Requires NEON.
func ReinterpretqS64U8(a arm.Uint8x16) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u16'.
// Requires NEON.
func ReinterpretqS64U16(a arm.Uint16x8) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u32'.
// Requires NEON.
func ReinterpretqS64U32(a arm.Uint32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_u64'.
// Requires NEON.
func ReinterpretqS64U64(a arm.Uint64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_p8'.
// Requires NEON.
func ReinterpretqS64P8(a arm.Poly8x16) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretqS64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s64_p16'.
// Requires NEON.
func ReinterpretqS64P16(a arm.Poly16x8) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ReinterpretU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_f64'.
// Requires NEON.
func ReinterpretU64F64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s8'.
// Requires NEON.
func ReinterpretU64S8(a arm.Int8x8) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s16'.
// Requires NEON.
func ReinterpretU64S16(a arm.Int16x4) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s32'.
// Requires NEON.
func ReinterpretU64S32(a arm.Int32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_s64'.
// Requires NEON.
func ReinterpretU64S64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_f32'.
// Requires NEON.
func ReinterpretU64F32(a arm.Float32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_u8'.
// Requires NEON.
func ReinterpretU64U8(a arm.Uint8x8) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_u16'.
// Requires NEON.
func ReinterpretU64U16(a arm.Uint16x4) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_u32'.
// Requires NEON.
func ReinterpretU64U32(a arm.Uint32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_p8'.
// Requires NEON.
func ReinterpretU64P8(a arm.Poly8x8) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretU64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u64_p16'.
// Requires NEON.
func ReinterpretU64P16(a arm.Poly16x4) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ReinterpretqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_f64'.
// Requires NEON.
func ReinterpretqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s8'.
// Requires NEON.
func ReinterpretqU64S8(a arm.Int8x16) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s16'.
// Requires NEON.
func ReinterpretqU64S16(a arm.Int16x8) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s32'.
// Requires NEON.
func ReinterpretqU64S32(a arm.Int32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_s64'.
// Requires NEON.
func ReinterpretqU64S64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_f32'.
// Requires NEON.
func ReinterpretqU64F32(a arm.Float32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_u8'.
// Requires NEON.
func ReinterpretqU64U8(a arm.Uint8x16) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_u16'.
// Requires NEON.
func ReinterpretqU64U16(a arm.Uint16x8) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_u32'.
// Requires NEON.
func ReinterpretqU64U32(a arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_p8'.
// Requires NEON.
func ReinterpretqU64P8(a arm.Poly8x16) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretqU64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u64_p16'.
// Requires NEON.
func ReinterpretqU64P16(a arm.Poly16x8) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ReinterpretS8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_f64'.
// Requires NEON.
func ReinterpretS8F64(a arm.Float64x1) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_s16'.
// Requires NEON.
func ReinterpretS8S16(a arm.Int16x4) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_s32'.
// Requires NEON.
func ReinterpretS8S32(a arm.Int32x2) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_s64'.
// Requires NEON.
func ReinterpretS8S64(a arm.Int64x1) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_f32'.
// Requires NEON.
func ReinterpretS8F32(a arm.Float32x2) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u8'.
// Requires NEON.
func ReinterpretS8U8(a arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u16'.
// Requires NEON.
func ReinterpretS8U16(a arm.Uint16x4) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u32'.
// Requires NEON.
func ReinterpretS8U32(a arm.Uint32x2) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_u64'.
// Requires NEON.
func ReinterpretS8U64(a arm.Uint64x1) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_p8'.
// Requires NEON.
func ReinterpretS8P8(a arm.Poly8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretS8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s8_p16'.
// Requires NEON.
func ReinterpretS8P16(a arm.Poly16x4) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ReinterpretqS8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_f64'.
// Requires NEON.
func ReinterpretqS8F64(a arm.Float64x2) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_s16'.
// Requires NEON.
func ReinterpretqS8S16(a arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_s32'.
// Requires NEON.
func ReinterpretqS8S32(a arm.Int32x4) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_s64'.
// Requires NEON.
func ReinterpretqS8S64(a arm.Int64x2) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_f32'.
// Requires NEON.
func ReinterpretqS8F32(a arm.Float32x4) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u8'.
// Requires NEON.
func ReinterpretqS8U8(a arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u16'.
// Requires NEON.
func ReinterpretqS8U16(a arm.Uint16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u32'.
// Requires NEON.
func ReinterpretqS8U32(a arm.Uint32x4) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_u64'.
// Requires NEON.
func ReinterpretqS8U64(a arm.Uint64x2) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_p8'.
// Requires NEON.
func ReinterpretqS8P8(a arm.Poly8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretqS8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s8_p16'.
// Requires NEON.
func ReinterpretqS8P16(a arm.Poly16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ReinterpretS16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_f64'.
// Requires NEON.
func ReinterpretS16F64(a arm.Float64x1) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_s8'.
// Requires NEON.
func ReinterpretS16S8(a arm.Int8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_s32'.
// Requires NEON.
func ReinterpretS16S32(a arm.Int32x2) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_s64'.
// Requires NEON.
func ReinterpretS16S64(a arm.Int64x1) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_f32'.
// Requires NEON.
func ReinterpretS16F32(a arm.Float32x2) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u8'.
// Requires NEON.
func ReinterpretS16U8(a arm.Uint8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u16'.
// Requires NEON.
func ReinterpretS16U16(a arm.Uint16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u32'.
// Requires NEON.
func ReinterpretS16U32(a arm.Uint32x2) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_u64'.
// Requires NEON.
func ReinterpretS16U64(a arm.Uint64x1) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_p8'.
// Requires NEON.
func ReinterpretS16P8(a arm.Poly8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretS16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s16_p16'.
// Requires NEON.
func ReinterpretS16P16(a arm.Poly16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ReinterpretqS16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_f64'.
// Requires NEON.
func ReinterpretqS16F64(a arm.Float64x2) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_s8'.
// Requires NEON.
func ReinterpretqS16S8(a arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_s32'.
// Requires NEON.
func ReinterpretqS16S32(a arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_s64'.
// Requires NEON.
func ReinterpretqS16S64(a arm.Int64x2) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_f32'.
// Requires NEON.
func ReinterpretqS16F32(a arm.Float32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u8'.
// Requires NEON.
func ReinterpretqS16U8(a arm.Uint8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u16'.
// Requires NEON.
func ReinterpretqS16U16(a arm.Uint16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u32'.
// Requires NEON.
func ReinterpretqS16U32(a arm.Uint32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_u64'.
// Requires NEON.
func ReinterpretqS16U64(a arm.Uint64x2) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_p8'.
// Requires NEON.
func ReinterpretqS16P8(a arm.Poly8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretqS16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s16_p16'.
// Requires NEON.
func ReinterpretqS16P16(a arm.Poly16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ReinterpretS32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_f64'.
// Requires NEON.
func ReinterpretS32F64(a arm.Float64x1) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_s8'.
// Requires NEON.
func ReinterpretS32S8(a arm.Int8x8) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_s16'.
// Requires NEON.
func ReinterpretS32S16(a arm.Int16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_s64'.
// Requires NEON.
func ReinterpretS32S64(a arm.Int64x1) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_f32'.
// Requires NEON.
func ReinterpretS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u8'.
// Requires NEON.
func ReinterpretS32U8(a arm.Uint8x8) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u16'.
// Requires NEON.
func ReinterpretS32U16(a arm.Uint16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u32'.
// Requires NEON.
func ReinterpretS32U32(a arm.Uint32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_u64'.
// Requires NEON.
func ReinterpretS32U64(a arm.Uint64x1) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_p8'.
// Requires NEON.
func ReinterpretS32P8(a arm.Poly8x8) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretS32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_s32_p16'.
// Requires NEON.
func ReinterpretS32P16(a arm.Poly16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ReinterpretqS32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_f64'.
// Requires NEON.
func ReinterpretqS32F64(a arm.Float64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_s8'.
// Requires NEON.
func ReinterpretqS32S8(a arm.Int8x16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_s16'.
// Requires NEON.
func ReinterpretqS32S16(a arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_s64'.
// Requires NEON.
func ReinterpretqS32S64(a arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_f32'.
// Requires NEON.
func ReinterpretqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u8'.
// Requires NEON.
func ReinterpretqS32U8(a arm.Uint8x16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u16'.
// Requires NEON.
func ReinterpretqS32U16(a arm.Uint16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u32'.
// Requires NEON.
func ReinterpretqS32U32(a arm.Uint32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_u64'.
// Requires NEON.
func ReinterpretqS32U64(a arm.Uint64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_p8'.
// Requires NEON.
func ReinterpretqS32P8(a arm.Poly8x16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretqS32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_s32_p16'.
// Requires NEON.
func ReinterpretqS32P16(a arm.Poly16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ReinterpretU8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_f64'.
// Requires NEON.
func ReinterpretU8F64(a arm.Float64x1) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s8'.
// Requires NEON.
func ReinterpretU8S8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s16'.
// Requires NEON.
func ReinterpretU8S16(a arm.Int16x4) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s32'.
// Requires NEON.
func ReinterpretU8S32(a arm.Int32x2) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_s64'.
// Requires NEON.
func ReinterpretU8S64(a arm.Int64x1) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_f32'.
// Requires NEON.
func ReinterpretU8F32(a arm.Float32x2) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_u16'.
// Requires NEON.
func ReinterpretU8U16(a arm.Uint16x4) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_u32'.
// Requires NEON.
func ReinterpretU8U32(a arm.Uint32x2) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_u64'.
// Requires NEON.
func ReinterpretU8U64(a arm.Uint64x1) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_p8'.
// Requires NEON.
func ReinterpretU8P8(a arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretU8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u8_p16'.
// Requires NEON.
func ReinterpretU8P16(a arm.Poly16x4) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ReinterpretqU8F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_f64'.
// Requires NEON.
func ReinterpretqU8F64(a arm.Float64x2) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s8'.
// Requires NEON.
func ReinterpretqU8S8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s16'.
// Requires NEON.
func ReinterpretqU8S16(a arm.Int16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s32'.
// Requires NEON.
func ReinterpretqU8S32(a arm.Int32x4) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_s64'.
// Requires NEON.
func ReinterpretqU8S64(a arm.Int64x2) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_f32'.
// Requires NEON.
func ReinterpretqU8F32(a arm.Float32x4) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_u16'.
// Requires NEON.
func ReinterpretqU8U16(a arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_u32'.
// Requires NEON.
func ReinterpretqU8U32(a arm.Uint32x4) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_u64'.
// Requires NEON.
func ReinterpretqU8U64(a arm.Uint64x2) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_p8'.
// Requires NEON.
func ReinterpretqU8P8(a arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretqU8P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u8_p16'.
// Requires NEON.
func ReinterpretqU8P16(a arm.Poly16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ReinterpretU16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_f64'.
// Requires NEON.
func ReinterpretU16F64(a arm.Float64x1) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s8'.
// Requires NEON.
func ReinterpretU16S8(a arm.Int8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s16'.
// Requires NEON.
func ReinterpretU16S16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s32'.
// Requires NEON.
func ReinterpretU16S32(a arm.Int32x2) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_s64'.
// Requires NEON.
func ReinterpretU16S64(a arm.Int64x1) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_f32'.
// Requires NEON.
func ReinterpretU16F32(a arm.Float32x2) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_u8'.
// Requires NEON.
func ReinterpretU16U8(a arm.Uint8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_u32'.
// Requires NEON.
func ReinterpretU16U32(a arm.Uint32x2) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_u64'.
// Requires NEON.
func ReinterpretU16U64(a arm.Uint64x1) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_p8'.
// Requires NEON.
func ReinterpretU16P8(a arm.Poly8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretU16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u16_p16'.
// Requires NEON.
func ReinterpretU16P16(a arm.Poly16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ReinterpretqU16F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_f64'.
// Requires NEON.
func ReinterpretqU16F64(a arm.Float64x2) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s8'.
// Requires NEON.
func ReinterpretqU16S8(a arm.Int8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s16'.
// Requires NEON.
func ReinterpretqU16S16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s32'.
// Requires NEON.
func ReinterpretqU16S32(a arm.Int32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_s64'.
// Requires NEON.
func ReinterpretqU16S64(a arm.Int64x2) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_f32'.
// Requires NEON.
func ReinterpretqU16F32(a arm.Float32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_u8'.
// Requires NEON.
func ReinterpretqU16U8(a arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16U32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_u32'.
// Requires NEON.
func ReinterpretqU16U32(a arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_u64'.
// Requires NEON.
func ReinterpretqU16U64(a arm.Uint64x2) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_p8'.
// Requires NEON.
func ReinterpretqU16P8(a arm.Poly8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretqU16P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u16_p16'.
// Requires NEON.
func ReinterpretqU16P16(a arm.Poly16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ReinterpretU32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_f64'.
// Requires NEON.
func ReinterpretU32F64(a arm.Float64x1) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s8'.
// Requires NEON.
func ReinterpretU32S8(a arm.Int8x8) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s16'.
// Requires NEON.
func ReinterpretU32S16(a arm.Int16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s32'.
// Requires NEON.
func ReinterpretU32S32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_s64'.
// Requires NEON.
func ReinterpretU32S64(a arm.Int64x1) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_f32'.
// Requires NEON.
func ReinterpretU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_u8'.
// Requires NEON.
func ReinterpretU32U8(a arm.Uint8x8) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_u16'.
// Requires NEON.
func ReinterpretU32U16(a arm.Uint16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_u64'.
// Requires NEON.
func ReinterpretU32U64(a arm.Uint64x1) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_p8'.
// Requires NEON.
func ReinterpretU32P8(a arm.Poly8x8) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretU32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpret_u32_p16'.
// Requires NEON.
func ReinterpretU32P16(a arm.Poly16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ReinterpretqU32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_f64'.
// Requires NEON.
func ReinterpretqU32F64(a arm.Float64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s8'.
// Requires NEON.
func ReinterpretqU32S8(a arm.Int8x16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s16'.
// Requires NEON.
func ReinterpretqU32S16(a arm.Int16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s32'.
// Requires NEON.
func ReinterpretqU32S32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32S64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_s64'.
// Requires NEON.
func ReinterpretqU32S64(a arm.Int64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_f32'.
// Requires NEON.
func ReinterpretqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_u8'.
// Requires NEON.
func ReinterpretqU32U8(a arm.Uint8x16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_u16'.
// Requires NEON.
func ReinterpretqU32U16(a arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32U64: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_u64'.
// Requires NEON.
func ReinterpretqU32U64(a arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_p8'.
// Requires NEON.
func ReinterpretqU32P8(a arm.Poly8x16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ReinterpretqU32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vreinterpretq_u32_p16'.
// Requires NEON.
func ReinterpretqU32P16(a arm.Poly16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SetLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneF32(elem float32, vec arm.Float32x2, index int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// SetLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneF64(elem float64, vec arm.Float64x1, index int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// SetLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneP8(elem arm.Poly8, vec arm.Poly8x8, index int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// SetLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneP16(elem arm.Poly16, vec arm.Poly16x4, index int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// SetLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneS8(elem int8, vec arm.Int8x8, index int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// SetLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneS16(elem int16, vec arm.Int16x4, index int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// SetLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneS32(elem int32, vec arm.Int32x2, index int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// SetLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneS64(elem int64, vec arm.Int64x1, index int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// SetLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneU8(elem uint8, vec arm.Uint8x8, index int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SetLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneU16(elem uint16, vec arm.Uint16x4, index int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SetLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneU32(elem uint32, vec arm.Uint32x2, index int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// SetLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vset_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetLaneU64(elem uint64, vec arm.Uint64x1, index int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// SetqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneF32(elem float32, vec arm.Float32x4, index int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// SetqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneF64(elem float64, vec arm.Float64x2, index int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// SetqLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneP8(elem arm.Poly8, vec arm.Poly8x16, index int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// SetqLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneP16(elem arm.Poly16, vec arm.Poly16x8, index int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// SetqLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneS8(elem int8, vec arm.Int8x16, index int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// SetqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneS16(elem int16, vec arm.Int16x8, index int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SetqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneS32(elem int32, vec arm.Int32x4, index int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SetqLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneS64(elem int64, vec arm.Int64x2, index int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SetqLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneU8(elem uint8, vec arm.Uint8x16, index int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SetqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneU16(elem uint16, vec arm.Uint16x8, index int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SetqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneU32(elem uint32, vec arm.Uint32x4, index int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SetqLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsetq_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SetqLaneU64(elem uint64, vec arm.Uint64x2, index int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// GetLowF32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_f32'.
// Requires NEON.
func GetLowF32(a arm.Float32x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// GetLowF64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_f64'.
// Requires NEON.
func GetLowF64(a arm.Float64x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// GetLowP8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_p8'.
// Requires NEON.
func GetLowP8(a arm.Poly8x16) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// GetLowP16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_p16'.
// Requires NEON.
func GetLowP16(a arm.Poly16x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// GetLowS8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s8'.
// Requires NEON.
func GetLowS8(a arm.Int8x16) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// GetLowS16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s16'.
// Requires NEON.
func GetLowS16(a arm.Int16x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// GetLowS32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s32'.
// Requires NEON.
func GetLowS32(a arm.Int32x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// GetLowS64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_s64'.
// Requires NEON.
func GetLowS64(a arm.Int64x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// GetLowU8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u8'.
// Requires NEON.
func GetLowU8(a arm.Uint8x16) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// GetLowU16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u16'.
// Requires NEON.
func GetLowU16(a arm.Uint16x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// GetLowU32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u32'.
// Requires NEON.
func GetLowU32(a arm.Uint32x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// GetLowU64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_low_u64'.
// Requires NEON.
func GetLowU64(a arm.Uint64x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// GetHighF32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_f32'.
// Requires NEON.
func GetHighF32(a arm.Float32x4) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// GetHighF64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_f64'.
// Requires NEON.
func GetHighF64(a arm.Float64x2) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// GetHighP8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_p8'.
// Requires NEON.
func GetHighP8(a arm.Poly8x16) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// GetHighP16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_p16'.
// Requires NEON.
func GetHighP16(a arm.Poly16x8) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// GetHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s8'.
// Requires NEON.
func GetHighS8(a arm.Int8x16) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// GetHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s16'.
// Requires NEON.
func GetHighS16(a arm.Int16x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// GetHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s32'.
// Requires NEON.
func GetHighS32(a arm.Int32x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// GetHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_s64'.
// Requires NEON.
func GetHighS64(a arm.Int64x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// GetHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u8'.
// Requires NEON.
func GetHighU8(a arm.Uint8x16) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// GetHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u16'.
// Requires NEON.
func GetHighU16(a arm.Uint16x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// GetHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u32'.
// Requires NEON.
func GetHighU32(a arm.Uint32x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// GetHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vget_high_u64'.
// Requires NEON.
func GetHighU64(a arm.Uint64x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CombineS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s8'.
// Requires NEON.
func CombineS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// CombineS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s16'.
// Requires NEON.
func CombineS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// CombineS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s32'.
// Requires NEON.
func CombineS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CombineS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_s64'.
// Requires NEON.
func CombineS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// CombineF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_f32'.
// Requires NEON.
func CombineF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// CombineU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u8'.
// Requires NEON.
func CombineU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CombineU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u16'.
// Requires NEON.
func CombineU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CombineU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u32'.
// Requires NEON.
func CombineU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CombineU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_u64'.
// Requires NEON.
func CombineU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CombineF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_f64'.
// Requires NEON.
func CombineF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// CombineP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_p8'.
// Requires NEON.
func CombineP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// CombineP16: ARM NEON intrinsic 
//
// Intrinsic: 'vcombine_p16'.
// Requires NEON.
func CombineP16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// AbaS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_s8'.
// Requires NEON.
func AbaS8(a arm.Int8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// AbaS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_s16'.
// Requires NEON.
func AbaS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// AbaS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_s32'.
// Requires NEON.
func AbaS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// AbaU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_u8'.
// Requires NEON.
func AbaU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// AbaU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_u16'.
// Requires NEON.
func AbaU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// AbaU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaba_u32'.
// Requires NEON.
func AbaU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// AbalHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_s8'.
// Requires NEON.
func AbalHighS8(a arm.Int16x8, b arm.Int8x16, c arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbalHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_s16'.
// Requires NEON.
func AbalHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbalHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_s32'.
// Requires NEON.
func AbalHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AbalHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_u8'.
// Requires NEON.
func AbalHighU8(a arm.Uint16x8, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AbalHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_u16'.
// Requires NEON.
func AbalHighU16(a arm.Uint32x4, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AbalHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_high_u32'.
// Requires NEON.
func AbalHighU32(a arm.Uint64x2, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AbalS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_s8'.
// Requires NEON.
func AbalS8(a arm.Int16x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_s16'.
// Requires NEON.
func AbalS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_s32'.
// Requires NEON.
func AbalS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AbalU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_u8'.
// Requires NEON.
func AbalU8(a arm.Uint16x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AbalU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_u16'.
// Requires NEON.
func AbalU16(a arm.Uint32x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AbalU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabal_u32'.
// Requires NEON.
func AbalU32(a arm.Uint64x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AbaqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_s8'.
// Requires NEON.
func AbaqS8(a arm.Int8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// AbaqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_s16'.
// Requires NEON.
func AbaqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbaqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_s32'.
// Requires NEON.
func AbaqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbaqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_u8'.
// Requires NEON.
func AbaqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AbaqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_u16'.
// Requires NEON.
func AbaqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AbaqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabaq_u32'.
// Requires NEON.
func AbaqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AbdF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_f32'.
// Requires NEON.
func AbdF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// AbdS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_s8'.
// Requires NEON.
func AbdS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// AbdS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_s16'.
// Requires NEON.
func AbdS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// AbdS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_s32'.
// Requires NEON.
func AbdS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// AbdU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_u8'.
// Requires NEON.
func AbdU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// AbdU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_u16'.
// Requires NEON.
func AbdU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// AbdU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabd_u32'.
// Requires NEON.
func AbdU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// AbddF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabdd_f64'.
// Requires NEON.
func AbddF64(a float64, b float64) float64 {
	return 0
}

// AbdlHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_s8'.
// Requires NEON.
func AbdlHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbdlHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_s16'.
// Requires NEON.
func AbdlHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbdlHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_s32'.
// Requires NEON.
func AbdlHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AbdlHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_u8'.
// Requires NEON.
func AbdlHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AbdlHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_u16'.
// Requires NEON.
func AbdlHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AbdlHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_high_u32'.
// Requires NEON.
func AbdlHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AbdlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_s8'.
// Requires NEON.
func AbdlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbdlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_s16'.
// Requires NEON.
func AbdlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbdlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_s32'.
// Requires NEON.
func AbdlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AbdlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_u8'.
// Requires NEON.
func AbdlU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AbdlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_u16'.
// Requires NEON.
func AbdlU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AbdlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdl_u32'.
// Requires NEON.
func AbdlU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AbdqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_f32'.
// Requires NEON.
func AbdqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// AbdqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_f64'.
// Requires NEON.
func AbdqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// AbdqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_s8'.
// Requires NEON.
func AbdqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// AbdqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_s16'.
// Requires NEON.
func AbdqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbdqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_s32'.
// Requires NEON.
func AbdqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbdqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_u8'.
// Requires NEON.
func AbdqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AbdqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_u16'.
// Requires NEON.
func AbdqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// AbdqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vabdq_u32'.
// Requires NEON.
func AbdqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// AbdsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabds_f32'.
// Requires NEON.
func AbdsF32(a float32, b float32) float32 {
	return 0
}

// AddlvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_s8'.
// Requires NEON.
func AddlvS8(a arm.Int8x8) int16 {
	return 0
}

// AddlvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_s16'.
// Requires NEON.
func AddlvS16(a arm.Int16x4) int32 {
	return 0
}

// AddlvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_u8'.
// Requires NEON.
func AddlvU8(a arm.Uint8x8) uint16 {
	return 0
}

// AddlvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_u16'.
// Requires NEON.
func AddlvU16(a arm.Uint16x4) uint32 {
	return 0
}

// AddlvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_s8'.
// Requires NEON.
func AddlvqS8(a arm.Int8x16) int16 {
	return 0
}

// AddlvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_s16'.
// Requires NEON.
func AddlvqS16(a arm.Int16x8) int32 {
	return 0
}

// AddlvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_s32'.
// Requires NEON.
func AddlvqS32(a arm.Int32x4) int64 {
	return 0
}

// AddlvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_u8'.
// Requires NEON.
func AddlvqU8(a arm.Uint8x16) uint16 {
	return 0
}

// AddlvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_u16'.
// Requires NEON.
func AddlvqU16(a arm.Uint16x8) uint32 {
	return 0
}

// AddlvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlvq_u32'.
// Requires NEON.
func AddlvqU32(a arm.Uint32x4) uint64 {
	return 0
}

// CvtxF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtx_f32_f64'.
// Requires NEON.
func CvtxF32F64(a arm.Float64x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// CvtxHighF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtx_high_f32_f64'.
// Requires NEON.
func CvtxHighF32F64(a arm.Float32x2, b arm.Float64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// CvtxdF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtxd_f32_f64'.
// Requires NEON.
func CvtxdF32F64(a float64) float32 {
	return 0
}

// MlaNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_f32'.
// Requires NEON.
func MlaNF32(a arm.Float32x2, b arm.Float32x2, c float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlaNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_s16'.
// Requires NEON.
func MlaNS16(a arm.Int16x4, b arm.Int16x4, c int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlaNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_s32'.
// Requires NEON.
func MlaNS32(a arm.Int32x2, b arm.Int32x2, c int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlaNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_u16'.
// Requires NEON.
func MlaNU16(a arm.Uint16x4, b arm.Uint16x4, c uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlaNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_n_u32'.
// Requires NEON.
func MlaNU32(a arm.Uint32x2, b arm.Uint32x2, c uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlaS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_s8'.
// Requires NEON.
func MlaS8(a arm.Int8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MlaS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_s16'.
// Requires NEON.
func MlaS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlaS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_s32'.
// Requires NEON.
func MlaS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlaU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_u8'.
// Requires NEON.
func MlaU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MlaU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_u16'.
// Requires NEON.
func MlaU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlaU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_u32'.
// Requires NEON.
func MlaU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlalHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_s16'.
// Requires NEON.
func MlalHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlalHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_s32'.
// Requires NEON.
func MlalHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlalHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_u16'.
// Requires NEON.
func MlalHighNU16(a arm.Uint32x4, b arm.Uint16x8, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlalHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_n_u32'.
// Requires NEON.
func MlalHighNU32(a arm.Uint64x2, b arm.Uint32x4, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlalHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_s8'.
// Requires NEON.
func MlalHighS8(a arm.Int16x8, b arm.Int8x16, c arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlalHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_s16'.
// Requires NEON.
func MlalHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlalHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_s32'.
// Requires NEON.
func MlalHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlalHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_u8'.
// Requires NEON.
func MlalHighU8(a arm.Uint16x8, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlalHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_u16'.
// Requires NEON.
func MlalHighU16(a arm.Uint32x4, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlalHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_high_u32'.
// Requires NEON.
func MlalHighU32(a arm.Uint64x2, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlalNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_s16'.
// Requires NEON.
func MlalNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlalNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_s32'.
// Requires NEON.
func MlalNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlalNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_u16'.
// Requires NEON.
func MlalNU16(a arm.Uint32x4, b arm.Uint16x4, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlalNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_n_u32'.
// Requires NEON.
func MlalNU32(a arm.Uint64x2, b arm.Uint32x2, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlalS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_s8'.
// Requires NEON.
func MlalS8(a arm.Int16x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_s16'.
// Requires NEON.
func MlalS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_s32'.
// Requires NEON.
func MlalS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlalU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_u8'.
// Requires NEON.
func MlalU8(a arm.Uint16x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlalU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_u16'.
// Requires NEON.
func MlalU16(a arm.Uint32x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlalU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlal_u32'.
// Requires NEON.
func MlalU32(a arm.Uint64x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlaqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_f32'.
// Requires NEON.
func MlaqNF32(a arm.Float32x4, b arm.Float32x4, c float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlaqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_s16'.
// Requires NEON.
func MlaqNS16(a arm.Int16x8, b arm.Int16x8, c int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlaqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_s32'.
// Requires NEON.
func MlaqNS32(a arm.Int32x4, b arm.Int32x4, c int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlaqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_u16'.
// Requires NEON.
func MlaqNU16(a arm.Uint16x8, b arm.Uint16x8, c uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlaqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_n_u32'.
// Requires NEON.
func MlaqNU32(a arm.Uint32x4, b arm.Uint32x4, c uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlaqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_s8'.
// Requires NEON.
func MlaqS8(a arm.Int8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MlaqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_s16'.
// Requires NEON.
func MlaqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlaqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_s32'.
// Requires NEON.
func MlaqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlaqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_u8'.
// Requires NEON.
func MlaqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MlaqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_u16'.
// Requires NEON.
func MlaqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlaqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_u32'.
// Requires NEON.
func MlaqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlsNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_f32'.
// Requires NEON.
func MlsNF32(a arm.Float32x2, b arm.Float32x2, c float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlsNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_s16'.
// Requires NEON.
func MlsNS16(a arm.Int16x4, b arm.Int16x4, c int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_s32'.
// Requires NEON.
func MlsNS32(a arm.Int32x2, b arm.Int32x2, c int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlsNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_u16'.
// Requires NEON.
func MlsNU16(a arm.Uint16x4, b arm.Uint16x4, c uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_n_u32'.
// Requires NEON.
func MlsNU32(a arm.Uint32x2, b arm.Uint32x2, c uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_s8'.
// Requires NEON.
func MlsS8(a arm.Int8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MlsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_s16'.
// Requires NEON.
func MlsS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_s32'.
// Requires NEON.
func MlsS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlsU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_u8'.
// Requires NEON.
func MlsU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MlsU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_u16'.
// Requires NEON.
func MlsU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_u32'.
// Requires NEON.
func MlsU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlslHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_s16'.
// Requires NEON.
func MlslHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlslHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_s32'.
// Requires NEON.
func MlslHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlslHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_u16'.
// Requires NEON.
func MlslHighNU16(a arm.Uint32x4, b arm.Uint16x8, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlslHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_n_u32'.
// Requires NEON.
func MlslHighNU32(a arm.Uint64x2, b arm.Uint32x4, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlslHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_s8'.
// Requires NEON.
func MlslHighS8(a arm.Int16x8, b arm.Int8x16, c arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlslHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_s16'.
// Requires NEON.
func MlslHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlslHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_s32'.
// Requires NEON.
func MlslHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlslHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_u8'.
// Requires NEON.
func MlslHighU8(a arm.Uint16x8, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlslHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_u16'.
// Requires NEON.
func MlslHighU16(a arm.Uint32x4, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlslHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_high_u32'.
// Requires NEON.
func MlslHighU32(a arm.Uint64x2, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlslNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_s16'.
// Requires NEON.
func MlslNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlslNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_s32'.
// Requires NEON.
func MlslNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlslNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_u16'.
// Requires NEON.
func MlslNU16(a arm.Uint32x4, b arm.Uint16x4, c uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlslNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_n_u32'.
// Requires NEON.
func MlslNU32(a arm.Uint64x2, b arm.Uint32x2, c uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlslS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_s8'.
// Requires NEON.
func MlslS8(a arm.Int16x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlslS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_s16'.
// Requires NEON.
func MlslS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlslS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_s32'.
// Requires NEON.
func MlslS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MlslU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_u8'.
// Requires NEON.
func MlslU8(a arm.Uint16x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlslU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_u16'.
// Requires NEON.
func MlslU16(a arm.Uint32x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlslU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsl_u32'.
// Requires NEON.
func MlslU32(a arm.Uint64x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MlsqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_f32'.
// Requires NEON.
func MlsqNF32(a arm.Float32x4, b arm.Float32x4, c float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlsqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_s16'.
// Requires NEON.
func MlsqNS16(a arm.Int16x8, b arm.Int16x8, c int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlsqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_s32'.
// Requires NEON.
func MlsqNS32(a arm.Int32x4, b arm.Int32x4, c int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlsqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_u16'.
// Requires NEON.
func MlsqNU16(a arm.Uint16x8, b arm.Uint16x8, c uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlsqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_n_u32'.
// Requires NEON.
func MlsqNU32(a arm.Uint32x4, b arm.Uint32x4, c uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_s8'.
// Requires NEON.
func MlsqS8(a arm.Int8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MlsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_s16'.
// Requires NEON.
func MlsqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_s32'.
// Requires NEON.
func MlsqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlsqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_u8'.
// Requires NEON.
func MlsqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MlsqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_u16'.
// Requires NEON.
func MlsqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlsqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_u32'.
// Requires NEON.
func MlsqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MovlHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_s8'.
// Requires NEON.
func MovlHighS8(a arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MovlHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_s16'.
// Requires NEON.
func MovlHighS16(a arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MovlHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_s32'.
// Requires NEON.
func MovlHighS32(a arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MovlHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_u8'.
// Requires NEON.
func MovlHighU8(a arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MovlHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_u16'.
// Requires NEON.
func MovlHighU16(a arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MovlHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_high_u32'.
// Requires NEON.
func MovlHighU32(a arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MovlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_s8'.
// Requires NEON.
func MovlS8(a arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MovlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_s16'.
// Requires NEON.
func MovlS16(a arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MovlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_s32'.
// Requires NEON.
func MovlS32(a arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MovlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_u8'.
// Requires NEON.
func MovlU8(a arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MovlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_u16'.
// Requires NEON.
func MovlU16(a arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MovlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovl_u32'.
// Requires NEON.
func MovlU32(a arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MovnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_s16'.
// Requires NEON.
func MovnHighS16(a arm.Int8x8, b arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MovnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_s32'.
// Requires NEON.
func MovnHighS32(a arm.Int16x4, b arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MovnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_s64'.
// Requires NEON.
func MovnHighS64(a arm.Int32x2, b arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MovnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_u16'.
// Requires NEON.
func MovnHighU16(a arm.Uint8x8, b arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MovnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_u32'.
// Requires NEON.
func MovnHighU32(a arm.Uint16x4, b arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MovnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_high_u64'.
// Requires NEON.
func MovnHighU64(a arm.Uint32x2, b arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MovnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_s16'.
// Requires NEON.
func MovnS16(a arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MovnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_s32'.
// Requires NEON.
func MovnS32(a arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MovnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_s64'.
// Requires NEON.
func MovnS64(a arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MovnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_u16'.
// Requires NEON.
func MovnU16(a arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MovnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_u32'.
// Requires NEON.
func MovnU32(a arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MovnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovn_u64'.
// Requires NEON.
func MovnU64(a arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MulNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_f32'.
// Requires NEON.
func MulNF32(a arm.Float32x2, b float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MulNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_s16'.
// Requires NEON.
func MulNS16(a arm.Int16x4, b int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MulNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_s32'.
// Requires NEON.
func MulNS32(a arm.Int32x2, b int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MulNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_u16'.
// Requires NEON.
func MulNU16(a arm.Uint16x4, b uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MulNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_u32'.
// Requires NEON.
func MulNU32(a arm.Uint32x2, b uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MullHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_s16'.
// Requires NEON.
func MullHighNS16(a arm.Int16x8, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MullHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_s32'.
// Requires NEON.
func MullHighNS32(a arm.Int32x4, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MullHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_u16'.
// Requires NEON.
func MullHighNU16(a arm.Uint16x8, b uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MullHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_n_u32'.
// Requires NEON.
func MullHighNU32(a arm.Uint32x4, b uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MullHighP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_p8'.
// Requires NEON.
func MullHighP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// MullHighS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_s8'.
// Requires NEON.
func MullHighS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MullHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_s16'.
// Requires NEON.
func MullHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MullHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_s32'.
// Requires NEON.
func MullHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MullHighU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_u8'.
// Requires NEON.
func MullHighU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MullHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_u16'.
// Requires NEON.
func MullHighU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MullHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_u32'.
// Requires NEON.
func MullHighU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MullNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_s16'.
// Requires NEON.
func MullNS16(a arm.Int16x4, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MullNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_s32'.
// Requires NEON.
func MullNS32(a arm.Int32x2, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MullNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_u16'.
// Requires NEON.
func MullNU16(a arm.Uint16x4, b uint16) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MullNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_n_u32'.
// Requires NEON.
func MullNU32(a arm.Uint32x2, b uint32) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MullP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_p8'.
// Requires NEON.
func MullP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// MullS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_s8'.
// Requires NEON.
func MullS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MullS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_s16'.
// Requires NEON.
func MullS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MullS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_s32'.
// Requires NEON.
func MullS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MullU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_u8'.
// Requires NEON.
func MullU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MullU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_u16'.
// Requires NEON.
func MullU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MullU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_u32'.
// Requires NEON.
func MullU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MulqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_f32'.
// Requires NEON.
func MulqNF32(a arm.Float32x4, b float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MulqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_f64'.
// Requires NEON.
func MulqNF64(a arm.Float64x2, b float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MulqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_s16'.
// Requires NEON.
func MulqNS16(a arm.Int16x8, b int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MulqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_s32'.
// Requires NEON.
func MulqNS32(a arm.Int32x4, b int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MulqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_u16'.
// Requires NEON.
func MulqNU16(a arm.Uint16x8, b uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MulqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_n_u32'.
// Requires NEON.
func MulqNU32(a arm.Uint32x4, b uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MulxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulx_f32'.
// Requires NEON.
func MulxF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MulxdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxd_f64'.
// Requires NEON.
func MulxdF64(a float64, b float64) float64 {
	return 0
}

// MulxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxq_f32'.
// Requires NEON.
func MulxqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MulxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxq_f64'.
// Requires NEON.
func MulxqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MulxsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulxs_f32'.
// Requires NEON.
func MulxsF32(a float32, b float32) float32 {
	return 0
}

// MvnP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_p8'.
// Requires NEON.
func MvnP8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// MvnS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_s8'.
// Requires NEON.
func MvnS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MvnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_s16'.
// Requires NEON.
func MvnS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MvnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_s32'.
// Requires NEON.
func MvnS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MvnU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_u8'.
// Requires NEON.
func MvnU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MvnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_u16'.
// Requires NEON.
func MvnU16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MvnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvn_u32'.
// Requires NEON.
func MvnU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MvnqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_p8'.
// Requires NEON.
func MvnqP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// MvnqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_s8'.
// Requires NEON.
func MvnqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MvnqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_s16'.
// Requires NEON.
func MvnqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MvnqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_s32'.
// Requires NEON.
func MvnqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MvnqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_u8'.
// Requires NEON.
func MvnqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MvnqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_u16'.
// Requires NEON.
func MvnqU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MvnqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmvnq_u32'.
// Requires NEON.
func MvnqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PadalS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_s8'.
// Requires NEON.
func PadalS8(a arm.Int16x4, b arm.Int8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// PadalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_s16'.
// Requires NEON.
func PadalS16(a arm.Int32x2, b arm.Int16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// PadalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_s32'.
// Requires NEON.
func PadalS32(a arm.Int64x1, b arm.Int32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// PadalU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_u8'.
// Requires NEON.
func PadalU8(a arm.Uint16x4, b arm.Uint8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// PadalU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_u16'.
// Requires NEON.
func PadalU16(a arm.Uint32x2, b arm.Uint16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// PadalU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadal_u32'.
// Requires NEON.
func PadalU32(a arm.Uint64x1, b arm.Uint32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// PadalqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_s8'.
// Requires NEON.
func PadalqS8(a arm.Int16x8, b arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// PadalqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_s16'.
// Requires NEON.
func PadalqS16(a arm.Int32x4, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// PadalqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_s32'.
// Requires NEON.
func PadalqS32(a arm.Int64x2, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// PadalqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_u8'.
// Requires NEON.
func PadalqU8(a arm.Uint16x8, b arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// PadalqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_u16'.
// Requires NEON.
func PadalqU16(a arm.Uint32x4, b arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PadalqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadalq_u32'.
// Requires NEON.
func PadalqU32(a arm.Uint64x2, b arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// PaddF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_f32'.
// Requires NEON.
func PaddF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// PaddlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_s8'.
// Requires NEON.
func PaddlS8(a arm.Int8x8) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// PaddlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_s16'.
// Requires NEON.
func PaddlS16(a arm.Int16x4) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// PaddlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_s32'.
// Requires NEON.
func PaddlS32(a arm.Int32x2) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// PaddlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_u8'.
// Requires NEON.
func PaddlU8(a arm.Uint8x8) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// PaddlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_u16'.
// Requires NEON.
func PaddlU16(a arm.Uint16x4) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// PaddlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddl_u32'.
// Requires NEON.
func PaddlU32(a arm.Uint32x2) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// PaddlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_s8'.
// Requires NEON.
func PaddlqS8(a arm.Int8x16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// PaddlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_s16'.
// Requires NEON.
func PaddlqS16(a arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// PaddlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_s32'.
// Requires NEON.
func PaddlqS32(a arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// PaddlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_u8'.
// Requires NEON.
func PaddlqU8(a arm.Uint8x16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// PaddlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_u16'.
// Requires NEON.
func PaddlqU16(a arm.Uint16x8) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PaddlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddlq_u32'.
// Requires NEON.
func PaddlqU32(a arm.Uint32x4) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// PaddqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_f32'.
// Requires NEON.
func PaddqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// PaddqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_f64'.
// Requires NEON.
func PaddqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// PaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s8'.
// Requires NEON.
func PaddqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// PaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s16'.
// Requires NEON.
func PaddqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// PaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s32'.
// Requires NEON.
func PaddqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// PaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_s64'.
// Requires NEON.
func PaddqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// PaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u8'.
// Requires NEON.
func PaddqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// PaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u16'.
// Requires NEON.
func PaddqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// PaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u32'.
// Requires NEON.
func PaddqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddq_u64'.
// Requires NEON.
func PaddqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// PaddsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadds_f32'.
// Requires NEON.
func PaddsF32(a arm.Float32x2) float32 {
	return 0
}

// QdmulhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_n_s16'.
// Requires NEON.
func QdmulhNS16(a arm.Int16x4, b int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QdmulhNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_n_s32'.
// Requires NEON.
func QdmulhNS32(a arm.Int32x2, b int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QdmulhqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_n_s16'.
// Requires NEON.
func QdmulhqNS16(a arm.Int16x8, b int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QdmulhqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_n_s32'.
// Requires NEON.
func QdmulhqNS32(a arm.Int32x4, b int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QmovnHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_s16'.
// Requires NEON.
func QmovnHighS16(a arm.Int8x8, b arm.Int16x8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QmovnHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_s32'.
// Requires NEON.
func QmovnHighS32(a arm.Int16x4, b arm.Int32x4) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QmovnHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_s64'.
// Requires NEON.
func QmovnHighS64(a arm.Int32x2, b arm.Int64x2) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QmovnHighU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_u16'.
// Requires NEON.
func QmovnHighU16(a arm.Uint8x8, b arm.Uint16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QmovnHighU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_u32'.
// Requires NEON.
func QmovnHighU32(a arm.Uint16x4, b arm.Uint32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QmovnHighU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_high_u64'.
// Requires NEON.
func QmovnHighU64(a arm.Uint32x2, b arm.Uint64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QmovunHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_high_s16'.
// Requires NEON.
func QmovunHighS16(a arm.Uint8x8, b arm.Int16x8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QmovunHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_high_s32'.
// Requires NEON.
func QmovunHighS32(a arm.Uint16x4, b arm.Int32x4) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QmovunHighS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_high_s64'.
// Requires NEON.
func QmovunHighS64(a arm.Uint32x2, b arm.Int64x2) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QrdmulhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_n_s16'.
// Requires NEON.
func QrdmulhNS16(a arm.Int16x4, b int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QrdmulhNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_n_s32'.
// Requires NEON.
func QrdmulhNS32(a arm.Int32x2, b int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QrdmulhqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_n_s16'.
// Requires NEON.
func QrdmulhqNS16(a arm.Int16x8, b int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QrdmulhqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_n_s32'.
// Requires NEON.
func QrdmulhqNS32(a arm.Int32x4, b int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RsqrteF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrte_f32'.
// Requires NEON.
func RsqrteF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RsqrteF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrte_f64'.
// Requires NEON.
func RsqrteF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RsqrteU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrte_u32'.
// Requires NEON.
func RsqrteU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RsqrtedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrted_f64'.
// Requires NEON.
func RsqrtedF64(a float64) float64 {
	return 0
}

// RsqrteqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrteq_f32'.
// Requires NEON.
func RsqrteqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RsqrteqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrteq_f64'.
// Requires NEON.
func RsqrteqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RsqrteqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrteq_u32'.
// Requires NEON.
func RsqrteqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RsqrtesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtes_f32'.
// Requires NEON.
func RsqrtesF32(a float32) float32 {
	return 0
}

// RsqrtsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrts_f32'.
// Requires NEON.
func RsqrtsF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RsqrtsdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtsd_f64'.
// Requires NEON.
func RsqrtsdF64(a float64, b float64) float64 {
	return 0
}

// RsqrtsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtsq_f32'.
// Requires NEON.
func RsqrtsqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RsqrtsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtsq_f64'.
// Requires NEON.
func RsqrtsqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RsqrtssF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsqrtss_f32'.
// Requires NEON.
func RsqrtssF32(a float32, b float32) float32 {
	return 0
}

// TstP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_p8'.
// Requires NEON.
func TstP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// TstP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_p16'.
// Requires NEON.
func TstP16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// TstqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_p8'.
// Requires NEON.
func TstqP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// TstqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_p16'.
// Requires NEON.
func TstqP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// St2LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneF32(ptr *float32, b [2]arm.Float32x2, c int)  {

}

// St2LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneF64(ptr *float64, b [2]arm.Float64x1, c int)  {

}

// St2LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneP8(ptr *arm.Poly8, b [2]arm.Poly8x8, c int)  {

}

// St2LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneP16(ptr *arm.Poly16, b [2]arm.Poly16x4, c int)  {

}

// St2LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneS8(ptr *int8, b [2]arm.Int8x8, c int)  {

}

// St2LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneS16(ptr *int16, b [2]arm.Int16x4, c int)  {

}

// St2LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneS32(ptr *int32, b [2]arm.Int32x2, c int)  {

}

// St2LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneS64(ptr *int64, b [2]arm.Int64x1, c int)  {

}

// St2LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneU8(ptr *uint8, b [2]arm.Uint8x8, c int)  {

}

// St2LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneU16(ptr *uint16, b [2]arm.Uint16x4, c int)  {

}

// St2LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneU32(ptr *uint32, b [2]arm.Uint32x2, c int)  {

}

// St2LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2LaneU64(ptr *uint64, b [2]arm.Uint64x1, c int)  {

}

// St2qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneF32(ptr *float32, b [2]arm.Float32x4, c int)  {

}

// St2qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneF64(ptr *float64, b [2]arm.Float64x2, c int)  {

}

// St2qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneP8(ptr *arm.Poly8, b [2]arm.Poly8x16, c int)  {

}

// St2qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneP16(ptr *arm.Poly16, b [2]arm.Poly16x8, c int)  {

}

// St2qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneS8(ptr *int8, b [2]arm.Int8x16, c int)  {

}

// St2qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneS16(ptr *int16, b [2]arm.Int16x8, c int)  {

}

// St2qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneS32(ptr *int32, b [2]arm.Int32x4, c int)  {

}

// St2qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneS64(ptr *int64, b [2]arm.Int64x2, c int)  {

}

// St2qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneU8(ptr *uint8, b [2]arm.Uint8x16, c int)  {

}

// St2qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneU16(ptr *uint16, b [2]arm.Uint16x8, c int)  {

}

// St2qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneU32(ptr *uint32, b [2]arm.Uint32x4, c int)  {

}

// St2qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St2qLaneU64(ptr *uint64, b [2]arm.Uint64x2, c int)  {

}

// St3LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneF32(ptr *float32, b [3]arm.Float32x2, c int)  {

}

// St3LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneF64(ptr *float64, b [3]arm.Float64x1, c int)  {

}

// St3LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneP8(ptr *arm.Poly8, b [3]arm.Poly8x8, c int)  {

}

// St3LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneP16(ptr *arm.Poly16, b [3]arm.Poly16x4, c int)  {

}

// St3LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneS8(ptr *int8, b [3]arm.Int8x8, c int)  {

}

// St3LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneS16(ptr *int16, b [3]arm.Int16x4, c int)  {

}

// St3LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneS32(ptr *int32, b [3]arm.Int32x2, c int)  {

}

// St3LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneS64(ptr *int64, b [3]arm.Int64x1, c int)  {

}

// St3LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneU8(ptr *uint8, b [3]arm.Uint8x8, c int)  {

}

// St3LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneU16(ptr *uint16, b [3]arm.Uint16x4, c int)  {

}

// St3LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneU32(ptr *uint32, b [3]arm.Uint32x2, c int)  {

}

// St3LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3LaneU64(ptr *uint64, b [3]arm.Uint64x1, c int)  {

}

// St3qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneF32(ptr *float32, b [3]arm.Float32x4, c int)  {

}

// St3qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneF64(ptr *float64, b [3]arm.Float64x2, c int)  {

}

// St3qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneP8(ptr *arm.Poly8, b [3]arm.Poly8x16, c int)  {

}

// St3qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneP16(ptr *arm.Poly16, b [3]arm.Poly16x8, c int)  {

}

// St3qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneS8(ptr *int8, b [3]arm.Int8x16, c int)  {

}

// St3qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneS16(ptr *int16, b [3]arm.Int16x8, c int)  {

}

// St3qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneS32(ptr *int32, b [3]arm.Int32x4, c int)  {

}

// St3qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneS64(ptr *int64, b [3]arm.Int64x2, c int)  {

}

// St3qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneU8(ptr *uint8, b [3]arm.Uint8x16, c int)  {

}

// St3qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneU16(ptr *uint16, b [3]arm.Uint16x8, c int)  {

}

// St3qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneU32(ptr *uint32, b [3]arm.Uint32x4, c int)  {

}

// St3qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St3qLaneU64(ptr *uint64, b [3]arm.Uint64x2, c int)  {

}

// St4LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneF32(ptr *float32, b [4]arm.Float32x2, c int)  {

}

// St4LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneF64(ptr *float64, b [4]arm.Float64x1, c int)  {

}

// St4LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneP8(ptr *arm.Poly8, b [4]arm.Poly8x8, c int)  {

}

// St4LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneP16(ptr *arm.Poly16, b [4]arm.Poly16x4, c int)  {

}

// St4LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneS8(ptr *int8, b [4]arm.Int8x8, c int)  {

}

// St4LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneS16(ptr *int16, b [4]arm.Int16x4, c int)  {

}

// St4LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneS32(ptr *int32, b [4]arm.Int32x2, c int)  {

}

// St4LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneS64(ptr *int64, b [4]arm.Int64x1, c int)  {

}

// St4LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneU8(ptr *uint8, b [4]arm.Uint8x8, c int)  {

}

// St4LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneU16(ptr *uint16, b [4]arm.Uint16x4, c int)  {

}

// St4LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneU32(ptr *uint32, b [4]arm.Uint32x2, c int)  {

}

// St4LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4LaneU64(ptr *uint64, b [4]arm.Uint64x1, c int)  {

}

// St4qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneF32(ptr *float32, b [4]arm.Float32x4, c int)  {

}

// St4qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneF64(ptr *float64, b [4]arm.Float64x2, c int)  {

}

// St4qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneP8(ptr *arm.Poly8, b [4]arm.Poly8x16, c int)  {

}

// St4qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneP16(ptr *arm.Poly16, b [4]arm.Poly16x8, c int)  {

}

// St4qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneS8(ptr *int8, b [4]arm.Int8x16, c int)  {

}

// St4qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneS16(ptr *int16, b [4]arm.Int16x8, c int)  {

}

// St4qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneS32(ptr *int32, b [4]arm.Int32x4, c int)  {

}

// St4qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneS64(ptr *int64, b [4]arm.Int64x2, c int)  {

}

// St4qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneU8(ptr *uint8, b [4]arm.Uint8x16, c int)  {

}

// St4qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneU16(ptr *uint16, b [4]arm.Uint16x8, c int)  {

}

// St4qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneU32(ptr *uint32, b [4]arm.Uint32x4, c int)  {

}

// St4qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St4qLaneU64(ptr *uint64, b [4]arm.Uint64x2, c int)  {

}

// AddlvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_s32'.
// Requires NEON.
func AddlvS32(a arm.Int32x2) int64 {
	return 0
}

// AddlvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddlv_u32'.
// Requires NEON.
func AddlvU32(a arm.Uint32x2) uint64 {
	return 0
}

// QdmulhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhLaneqS16(a arm.Int16x4, b arm.Int16x8, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QdmulhLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhLaneqS32(a arm.Int32x2, b arm.Int32x4, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QdmulhqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhqLaneqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QdmulhqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhqLaneqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QrdmulhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhLaneqS16(a arm.Int16x4, b arm.Int16x8, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QrdmulhLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhLaneqS32(a arm.Int32x2, b arm.Int32x4, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QrdmulhqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhqLaneqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QrdmulhqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhqLaneqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Qtbl1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1_p8'.
// Requires NEON.
func Qtbl1P8(a arm.Poly8x16, b arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbl1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1_s8'.
// Requires NEON.
func Qtbl1S8(a arm.Int8x16, b arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbl1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1_u8'.
// Requires NEON.
func Qtbl1U8(a arm.Uint8x16, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbl1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1q_p8'.
// Requires NEON.
func Qtbl1qP8(a arm.Poly8x16, b arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbl1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1q_s8'.
// Requires NEON.
func Qtbl1qS8(a arm.Int8x16, b arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbl1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl1q_u8'.
// Requires NEON.
func Qtbl1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbl2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2_s8'.
// Requires NEON.
func Qtbl2S8(tab [2]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbl2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2_u8'.
// Requires NEON.
func Qtbl2U8(tab [2]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbl2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2_p8'.
// Requires NEON.
func Qtbl2P8(tab [2]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbl2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2q_s8'.
// Requires NEON.
func Qtbl2qS8(tab [2]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbl2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2q_u8'.
// Requires NEON.
func Qtbl2qU8(tab [2]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbl2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl2q_p8'.
// Requires NEON.
func Qtbl2qP8(tab [2]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbl3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3_s8'.
// Requires NEON.
func Qtbl3S8(tab [3]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbl3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3_u8'.
// Requires NEON.
func Qtbl3U8(tab [3]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbl3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3_p8'.
// Requires NEON.
func Qtbl3P8(tab [3]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbl3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3q_s8'.
// Requires NEON.
func Qtbl3qS8(tab [3]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbl3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3q_u8'.
// Requires NEON.
func Qtbl3qU8(tab [3]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbl3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl3q_p8'.
// Requires NEON.
func Qtbl3qP8(tab [3]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbl4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4_s8'.
// Requires NEON.
func Qtbl4S8(tab [4]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbl4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4_u8'.
// Requires NEON.
func Qtbl4U8(tab [4]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbl4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4_p8'.
// Requires NEON.
func Qtbl4P8(tab [4]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbl4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4q_s8'.
// Requires NEON.
func Qtbl4qS8(tab [4]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbl4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4q_u8'.
// Requires NEON.
func Qtbl4qU8(tab [4]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbl4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbl4q_p8'.
// Requires NEON.
func Qtbl4qP8(tab [4]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbx1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1_s8'.
// Requires NEON.
func Qtbx1S8(r arm.Int8x8, tab arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbx1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1_u8'.
// Requires NEON.
func Qtbx1U8(r arm.Uint8x8, tab arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbx1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1_p8'.
// Requires NEON.
func Qtbx1P8(r arm.Poly8x8, tab arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbx1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1q_s8'.
// Requires NEON.
func Qtbx1qS8(r arm.Int8x16, tab arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbx1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1q_u8'.
// Requires NEON.
func Qtbx1qU8(r arm.Uint8x16, tab arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbx1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx1q_p8'.
// Requires NEON.
func Qtbx1qP8(r arm.Poly8x16, tab arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbx2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2_s8'.
// Requires NEON.
func Qtbx2S8(r arm.Int8x8, tab [2]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbx2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2_u8'.
// Requires NEON.
func Qtbx2U8(r arm.Uint8x8, tab [2]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbx2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2_p8'.
// Requires NEON.
func Qtbx2P8(r arm.Poly8x8, tab [2]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbx2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2q_s8'.
// Requires NEON.
func Qtbx2qS8(r arm.Int8x16, tab [2]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbx2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2q_u8'.
// Requires NEON.
func Qtbx2qU8(r arm.Uint8x16, tab [2]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbx2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx2q_p8'.
// Requires NEON.
func Qtbx2qP8(r arm.Poly8x16, tab [2]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbx3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3_s8'.
// Requires NEON.
func Qtbx3S8(r arm.Int8x8, tab [3]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbx3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3_u8'.
// Requires NEON.
func Qtbx3U8(r arm.Uint8x8, tab [3]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbx3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3_p8'.
// Requires NEON.
func Qtbx3P8(r arm.Poly8x8, tab [3]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbx3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3q_s8'.
// Requires NEON.
func Qtbx3qS8(r arm.Int8x16, tab [3]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbx3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3q_u8'.
// Requires NEON.
func Qtbx3qU8(r arm.Uint8x16, tab [3]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbx3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx3q_p8'.
// Requires NEON.
func Qtbx3qP8(r arm.Poly8x16, tab [3]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Qtbx4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4_s8'.
// Requires NEON.
func Qtbx4S8(r arm.Int8x8, tab [4]arm.Int8x16, idx arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Qtbx4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4_u8'.
// Requires NEON.
func Qtbx4U8(r arm.Uint8x8, tab [4]arm.Uint8x16, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Qtbx4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4_p8'.
// Requires NEON.
func Qtbx4P8(r arm.Poly8x8, tab [4]arm.Poly8x16, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Qtbx4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4q_s8'.
// Requires NEON.
func Qtbx4qS8(r arm.Int8x16, tab [4]arm.Int8x16, idx arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Qtbx4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4q_u8'.
// Requires NEON.
func Qtbx4qU8(r arm.Uint8x16, tab [4]arm.Uint8x16, idx arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Qtbx4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vqtbx4q_p8'.
// Requires NEON.
func Qtbx4qP8(r arm.Poly8x16, tab [4]arm.Poly8x16, idx arm.Uint8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Tbl1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl1_s8'.
// Requires NEON.
func Tbl1S8(tab arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbl1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl1_u8'.
// Requires NEON.
func Tbl1U8(tab arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbl1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl1_p8'.
// Requires NEON.
func Tbl1P8(tab arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Tbl2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl2_s8'.
// Requires NEON.
func Tbl2S8(tab [2]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbl2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl2_u8'.
// Requires NEON.
func Tbl2U8(tab [2]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbl2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl2_p8'.
// Requires NEON.
func Tbl2P8(tab [2]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Tbl3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl3_s8'.
// Requires NEON.
func Tbl3S8(tab [3]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbl3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl3_u8'.
// Requires NEON.
func Tbl3U8(tab [3]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbl3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl3_p8'.
// Requires NEON.
func Tbl3P8(tab [3]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Tbl4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl4_s8'.
// Requires NEON.
func Tbl4S8(tab [4]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbl4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl4_u8'.
// Requires NEON.
func Tbl4U8(tab [4]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbl4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbl4_p8'.
// Requires NEON.
func Tbl4P8(tab [4]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Tbx2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx2_s8'.
// Requires NEON.
func Tbx2S8(r arm.Int8x8, tab [2]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbx2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx2_u8'.
// Requires NEON.
func Tbx2U8(r arm.Uint8x8, tab [2]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbx2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx2_p8'.
// Requires NEON.
func Tbx2P8(r arm.Poly8x8, tab [2]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Tbx4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx4_s8'.
// Requires NEON.
func Tbx4S8(r arm.Int8x8, tab [4]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbx4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx4_u8'.
// Requires NEON.
func Tbx4U8(r arm.Uint8x8, tab [4]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbx4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx4_p8'.
// Requires NEON.
func Tbx4P8(r arm.Poly8x8, tab [4]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// AbsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_f32'.
// Requires NEON.
func AbsF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// AbsF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_f64'.
// Requires NEON.
func AbsF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// AbsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s8'.
// Requires NEON.
func AbsS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// AbsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s16'.
// Requires NEON.
func AbsS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// AbsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s32'.
// Requires NEON.
func AbsS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// AbsS64: ARM NEON intrinsic 
//
// Intrinsic: 'vabs_s64'.
// Requires NEON.
func AbsS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// AbsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_f32'.
// Requires NEON.
func AbsqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// AbsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_f64'.
// Requires NEON.
func AbsqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// AbsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s8'.
// Requires NEON.
func AbsqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// AbsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s16'.
// Requires NEON.
func AbsqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// AbsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s32'.
// Requires NEON.
func AbsqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// AbsqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vabsq_s64'.
// Requires NEON.
func AbsqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// AdddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddd_s64'.
// Requires NEON.
func AdddS64(a int64, b int64) int64 {
	return 0
}

// AdddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddd_u64'.
// Requires NEON.
func AdddU64(a uint64, b uint64) uint64 {
	return 0
}

// AddvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_s8'.
// Requires NEON.
func AddvS8(a arm.Int8x8) int8 {
	return 0
}

// AddvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_s16'.
// Requires NEON.
func AddvS16(a arm.Int16x4) int16 {
	return 0
}

// AddvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_s32'.
// Requires NEON.
func AddvS32(a arm.Int32x2) int32 {
	return 0
}

// AddvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_u8'.
// Requires NEON.
func AddvU8(a arm.Uint8x8) uint8 {
	return 0
}

// AddvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_u16'.
// Requires NEON.
func AddvU16(a arm.Uint16x4) uint16 {
	return 0
}

// AddvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_u32'.
// Requires NEON.
func AddvU32(a arm.Uint32x2) uint32 {
	return 0
}

// AddvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s8'.
// Requires NEON.
func AddvqS8(a arm.Int8x16) int8 {
	return 0
}

// AddvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s16'.
// Requires NEON.
func AddvqS16(a arm.Int16x8) int16 {
	return 0
}

// AddvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s32'.
// Requires NEON.
func AddvqS32(a arm.Int32x4) int32 {
	return 0
}

// AddvqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_s64'.
// Requires NEON.
func AddvqS64(a arm.Int64x2) int64 {
	return 0
}

// AddvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u8'.
// Requires NEON.
func AddvqU8(a arm.Uint8x16) uint8 {
	return 0
}

// AddvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u16'.
// Requires NEON.
func AddvqU16(a arm.Uint16x8) uint16 {
	return 0
}

// AddvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u32'.
// Requires NEON.
func AddvqU32(a arm.Uint32x4) uint32 {
	return 0
}

// AddvqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_u64'.
// Requires NEON.
func AddvqU64(a arm.Uint64x2) uint64 {
	return 0
}

// AddvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddv_f32'.
// Requires NEON.
func AddvF32(a arm.Float32x2) float32 {
	return 0
}

// AddvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_f32'.
// Requires NEON.
func AddvqF32(a arm.Float32x4) float32 {
	return 0
}

// AddvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vaddvq_f64'.
// Requires NEON.
func AddvqF64(a arm.Float64x2) float64 {
	return 0
}

// BslF32: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_f32'.
// Requires NEON.
func BslF32(a arm.Uint32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// BslF64: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_f64'.
// Requires NEON.
func BslF64(a arm.Uint64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// BslP8: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_p8'.
// Requires NEON.
func BslP8(a arm.Uint8x8, b arm.Poly8x8, c arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// BslP16: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_p16'.
// Requires NEON.
func BslP16(a arm.Uint16x4, b arm.Poly16x4, c arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// BslS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s8'.
// Requires NEON.
func BslS8(a arm.Uint8x8, b arm.Int8x8, c arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// BslS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s16'.
// Requires NEON.
func BslS16(a arm.Uint16x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// BslS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s32'.
// Requires NEON.
func BslS32(a arm.Uint32x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// BslS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_s64'.
// Requires NEON.
func BslS64(a arm.Uint64x1, b arm.Int64x1, c arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// BslU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u8'.
// Requires NEON.
func BslU8(a arm.Uint8x8, b arm.Uint8x8, c arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// BslU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u16'.
// Requires NEON.
func BslU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// BslU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u32'.
// Requires NEON.
func BslU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// BslU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbsl_u64'.
// Requires NEON.
func BslU64(a arm.Uint64x1, b arm.Uint64x1, c arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// BslqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_f32'.
// Requires NEON.
func BslqF32(a arm.Uint32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// BslqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_f64'.
// Requires NEON.
func BslqF64(a arm.Uint64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// BslqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_p8'.
// Requires NEON.
func BslqP8(a arm.Uint8x16, b arm.Poly8x16, c arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// BslqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_p16'.
// Requires NEON.
func BslqP16(a arm.Uint16x8, b arm.Poly16x8, c arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// BslqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s8'.
// Requires NEON.
func BslqS8(a arm.Uint8x16, b arm.Int8x16, c arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// BslqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s16'.
// Requires NEON.
func BslqS16(a arm.Uint16x8, b arm.Int16x8, c arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// BslqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s32'.
// Requires NEON.
func BslqS32(a arm.Uint32x4, b arm.Int32x4, c arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// BslqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_s64'.
// Requires NEON.
func BslqS64(a arm.Uint64x2, b arm.Int64x2, c arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// BslqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u8'.
// Requires NEON.
func BslqU8(a arm.Uint8x16, b arm.Uint8x16, c arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// BslqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u16'.
// Requires NEON.
func BslqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// BslqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u32'.
// Requires NEON.
func BslqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// BslqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vbslq_u64'.
// Requires NEON.
func BslqU64(a arm.Uint64x2, b arm.Uint64x2, c arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// AeseqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaeseq_u8'.
// Requires NEON.
func AeseqU8(data arm.Uint8x16, key arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AesdqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaesdq_u8'.
// Requires NEON.
func AesdqU8(data arm.Uint8x16, key arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AesmcqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaesmcq_u8'.
// Requires NEON.
func AesmcqU8(data arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// AesimcqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vaesimcq_u8'.
// Requires NEON.
func AesimcqU8(data arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CageF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcage_f64'.
// Requires NEON.
func CageF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CagesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcages_f32'.
// Requires NEON.
func CagesF32(a float32, b float32) uint32 {
	return 0
}

// CageF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcage_f32'.
// Requires NEON.
func CageF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CageqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcageq_f32'.
// Requires NEON.
func CageqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CagedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaged_f64'.
// Requires NEON.
func CagedF64(a float64, b float64) uint64 {
	return 0
}

// CageqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcageq_f64'.
// Requires NEON.
func CageqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CagtsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcagts_f32'.
// Requires NEON.
func CagtsF32(a float32, b float32) uint32 {
	return 0
}

// CagtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcagt_f32'.
// Requires NEON.
func CagtF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CagtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcagt_f64'.
// Requires NEON.
func CagtF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CagtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcagtq_f32'.
// Requires NEON.
func CagtqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CagtdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcagtd_f64'.
// Requires NEON.
func CagtdF64(a float64, b float64) uint64 {
	return 0
}

// CagtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcagtq_f64'.
// Requires NEON.
func CagtqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CaleF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcale_f32'.
// Requires NEON.
func CaleF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CaleF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcale_f64'.
// Requires NEON.
func CaleF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CaledF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaled_f64'.
// Requires NEON.
func CaledF64(a float64, b float64) uint64 {
	return 0
}

// CalesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcales_f32'.
// Requires NEON.
func CalesF32(a float32, b float32) uint32 {
	return 0
}

// CaleqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcaleq_f32'.
// Requires NEON.
func CaleqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CaleqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaleq_f64'.
// Requires NEON.
func CaleqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CaltF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcalt_f32'.
// Requires NEON.
func CaltF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CaltF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcalt_f64'.
// Requires NEON.
func CaltF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CaltdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaltd_f64'.
// Requires NEON.
func CaltdF64(a float64, b float64) uint64 {
	return 0
}

// CaltqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcaltq_f32'.
// Requires NEON.
func CaltqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CaltqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcaltq_f64'.
// Requires NEON.
func CaltqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CaltsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcalts_f32'.
// Requires NEON.
func CaltsF32(a float32, b float32) uint32 {
	return 0
}

// CeqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_f32'.
// Requires NEON.
func CeqF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CeqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_f64'.
// Requires NEON.
func CeqF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CeqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_p8'.
// Requires NEON.
func CeqP8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CeqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s8'.
// Requires NEON.
func CeqS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CeqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s16'.
// Requires NEON.
func CeqS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CeqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s32'.
// Requires NEON.
func CeqS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CeqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_s64'.
// Requires NEON.
func CeqS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CeqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u8'.
// Requires NEON.
func CeqU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CeqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u16'.
// Requires NEON.
func CeqU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CeqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u32'.
// Requires NEON.
func CeqU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CeqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceq_u64'.
// Requires NEON.
func CeqU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CeqqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_f32'.
// Requires NEON.
func CeqqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CeqqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_f64'.
// Requires NEON.
func CeqqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CeqqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_p8'.
// Requires NEON.
func CeqqP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CeqqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s8'.
// Requires NEON.
func CeqqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CeqqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s16'.
// Requires NEON.
func CeqqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CeqqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s32'.
// Requires NEON.
func CeqqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CeqqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_s64'.
// Requires NEON.
func CeqqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CeqqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u8'.
// Requires NEON.
func CeqqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CeqqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u16'.
// Requires NEON.
func CeqqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CeqqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u32'.
// Requires NEON.
func CeqqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CeqqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqq_u64'.
// Requires NEON.
func CeqqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CeqsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqs_f32'.
// Requires NEON.
func CeqsF32(a float32, b float32) uint32 {
	return 0
}

// CeqdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqd_s64'.
// Requires NEON.
func CeqdS64(a int64, b int64) uint64 {
	return 0
}

// CeqdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqd_u64'.
// Requires NEON.
func CeqdU64(a uint64, b uint64) uint64 {
	return 0
}

// CeqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqd_f64'.
// Requires NEON.
func CeqdF64(a float64, b float64) uint64 {
	return 0
}

// CeqzF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_f32'.
// Requires NEON.
func CeqzF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CeqzF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_f64'.
// Requires NEON.
func CeqzF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CeqzP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_p8'.
// Requires NEON.
func CeqzP8(a arm.Poly8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CeqzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s8'.
// Requires NEON.
func CeqzS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CeqzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s16'.
// Requires NEON.
func CeqzS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CeqzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s32'.
// Requires NEON.
func CeqzS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CeqzS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_s64'.
// Requires NEON.
func CeqzS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CeqzU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u8'.
// Requires NEON.
func CeqzU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CeqzU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u16'.
// Requires NEON.
func CeqzU16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CeqzU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u32'.
// Requires NEON.
func CeqzU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CeqzU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqz_u64'.
// Requires NEON.
func CeqzU64(a arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CeqzqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_f32'.
// Requires NEON.
func CeqzqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CeqzqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_f64'.
// Requires NEON.
func CeqzqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CeqzqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_p8'.
// Requires NEON.
func CeqzqP8(a arm.Poly8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CeqzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s8'.
// Requires NEON.
func CeqzqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CeqzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s16'.
// Requires NEON.
func CeqzqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CeqzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s32'.
// Requires NEON.
func CeqzqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CeqzqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_s64'.
// Requires NEON.
func CeqzqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CeqzqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u8'.
// Requires NEON.
func CeqzqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CeqzqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u16'.
// Requires NEON.
func CeqzqU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CeqzqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u32'.
// Requires NEON.
func CeqzqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CeqzqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzq_u64'.
// Requires NEON.
func CeqzqU64(a arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CeqzsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzs_f32'.
// Requires NEON.
func CeqzsF32(a float32) uint32 {
	return 0
}

// CeqzdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzd_s64'.
// Requires NEON.
func CeqzdS64(a int64) uint64 {
	return 0
}

// CeqzdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzd_u64'.
// Requires NEON.
func CeqzdU64(a uint64) uint64 {
	return 0
}

// CeqzdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vceqzd_f64'.
// Requires NEON.
func CeqzdF64(a float64) uint64 {
	return 0
}

// CgeF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_f32'.
// Requires NEON.
func CgeF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgeF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_f64'.
// Requires NEON.
func CgeF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgeS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s8'.
// Requires NEON.
func CgeS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CgeS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s16'.
// Requires NEON.
func CgeS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CgeS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s32'.
// Requires NEON.
func CgeS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgeS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_s64'.
// Requires NEON.
func CgeS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgeU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u8'.
// Requires NEON.
func CgeU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CgeU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u16'.
// Requires NEON.
func CgeU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CgeU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u32'.
// Requires NEON.
func CgeU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgeU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcge_u64'.
// Requires NEON.
func CgeU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgeqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_f32'.
// Requires NEON.
func CgeqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgeqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_f64'.
// Requires NEON.
func CgeqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgeqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s8'.
// Requires NEON.
func CgeqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CgeqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s16'.
// Requires NEON.
func CgeqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CgeqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s32'.
// Requires NEON.
func CgeqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgeqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_s64'.
// Requires NEON.
func CgeqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgeqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u8'.
// Requires NEON.
func CgeqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CgeqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u16'.
// Requires NEON.
func CgeqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CgeqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u32'.
// Requires NEON.
func CgeqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgeqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgeq_u64'.
// Requires NEON.
func CgeqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcges_f32'.
// Requires NEON.
func CgesF32(a float32, b float32) uint32 {
	return 0
}

// CgedS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcged_s64'.
// Requires NEON.
func CgedS64(a int64, b int64) uint64 {
	return 0
}

// CgedU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcged_u64'.
// Requires NEON.
func CgedU64(a uint64, b uint64) uint64 {
	return 0
}

// CgedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcged_f64'.
// Requires NEON.
func CgedF64(a float64, b float64) uint64 {
	return 0
}

// CgezF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_f32'.
// Requires NEON.
func CgezF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgezF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_f64'.
// Requires NEON.
func CgezF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgezS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s8'.
// Requires NEON.
func CgezS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CgezS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s16'.
// Requires NEON.
func CgezS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CgezS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s32'.
// Requires NEON.
func CgezS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgezS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgez_s64'.
// Requires NEON.
func CgezS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgezqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_f32'.
// Requires NEON.
func CgezqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgezqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_f64'.
// Requires NEON.
func CgezqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgezqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s8'.
// Requires NEON.
func CgezqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CgezqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s16'.
// Requires NEON.
func CgezqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CgezqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s32'.
// Requires NEON.
func CgezqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgezqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezq_s64'.
// Requires NEON.
func CgezqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgezsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezs_f32'.
// Requires NEON.
func CgezsF32(a float32) uint32 {
	return 0
}

// CgezdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezd_s64'.
// Requires NEON.
func CgezdS64(a int64) uint64 {
	return 0
}

// CgezdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgezd_f64'.
// Requires NEON.
func CgezdF64(a float64) uint64 {
	return 0
}

// CgtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_f32'.
// Requires NEON.
func CgtF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_f64'.
// Requires NEON.
func CgtF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgtS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s8'.
// Requires NEON.
func CgtS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CgtS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s16'.
// Requires NEON.
func CgtS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CgtS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s32'.
// Requires NEON.
func CgtS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgtS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_s64'.
// Requires NEON.
func CgtS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgtU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u8'.
// Requires NEON.
func CgtU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CgtU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u16'.
// Requires NEON.
func CgtU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CgtU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u32'.
// Requires NEON.
func CgtU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgtU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgt_u64'.
// Requires NEON.
func CgtU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_f32'.
// Requires NEON.
func CgtqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_f64'.
// Requires NEON.
func CgtqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgtqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s8'.
// Requires NEON.
func CgtqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CgtqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s16'.
// Requires NEON.
func CgtqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CgtqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s32'.
// Requires NEON.
func CgtqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgtqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_s64'.
// Requires NEON.
func CgtqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgtqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u8'.
// Requires NEON.
func CgtqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CgtqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u16'.
// Requires NEON.
func CgtqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CgtqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u32'.
// Requires NEON.
func CgtqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgtqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtq_u64'.
// Requires NEON.
func CgtqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgtsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgts_f32'.
// Requires NEON.
func CgtsF32(a float32, b float32) uint32 {
	return 0
}

// CgtdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtd_s64'.
// Requires NEON.
func CgtdS64(a int64, b int64) uint64 {
	return 0
}

// CgtdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtd_u64'.
// Requires NEON.
func CgtdU64(a uint64, b uint64) uint64 {
	return 0
}

// CgtdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtd_f64'.
// Requires NEON.
func CgtdF64(a float64, b float64) uint64 {
	return 0
}

// CgtzF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_f32'.
// Requires NEON.
func CgtzF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgtzF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_f64'.
// Requires NEON.
func CgtzF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgtzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s8'.
// Requires NEON.
func CgtzS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CgtzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s16'.
// Requires NEON.
func CgtzS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CgtzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s32'.
// Requires NEON.
func CgtzS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CgtzS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtz_s64'.
// Requires NEON.
func CgtzS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CgtzqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_f32'.
// Requires NEON.
func CgtzqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgtzqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_f64'.
// Requires NEON.
func CgtzqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgtzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s8'.
// Requires NEON.
func CgtzqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CgtzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s16'.
// Requires NEON.
func CgtzqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CgtzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s32'.
// Requires NEON.
func CgtzqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CgtzqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzq_s64'.
// Requires NEON.
func CgtzqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CgtzsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzs_f32'.
// Requires NEON.
func CgtzsF32(a float32) uint32 {
	return 0
}

// CgtzdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzd_s64'.
// Requires NEON.
func CgtzdS64(a int64) uint64 {
	return 0
}

// CgtzdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcgtzd_f64'.
// Requires NEON.
func CgtzdF64(a float64) uint64 {
	return 0
}

// CleF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_f32'.
// Requires NEON.
func CleF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CleF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_f64'.
// Requires NEON.
func CleF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CleS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s8'.
// Requires NEON.
func CleS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CleS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s16'.
// Requires NEON.
func CleS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CleS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s32'.
// Requires NEON.
func CleS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CleS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_s64'.
// Requires NEON.
func CleS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CleU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u8'.
// Requires NEON.
func CleU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CleU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u16'.
// Requires NEON.
func CleU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CleU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u32'.
// Requires NEON.
func CleU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CleU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcle_u64'.
// Requires NEON.
func CleU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CleqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_f32'.
// Requires NEON.
func CleqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CleqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_f64'.
// Requires NEON.
func CleqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CleqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s8'.
// Requires NEON.
func CleqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CleqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s16'.
// Requires NEON.
func CleqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CleqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s32'.
// Requires NEON.
func CleqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CleqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_s64'.
// Requires NEON.
func CleqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CleqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u8'.
// Requires NEON.
func CleqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CleqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u16'.
// Requires NEON.
func CleqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CleqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u32'.
// Requires NEON.
func CleqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CleqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcleq_u64'.
// Requires NEON.
func CleqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ClesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcles_f32'.
// Requires NEON.
func ClesF32(a float32, b float32) uint32 {
	return 0
}

// CledS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcled_s64'.
// Requires NEON.
func CledS64(a int64, b int64) uint64 {
	return 0
}

// CledU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcled_u64'.
// Requires NEON.
func CledU64(a uint64, b uint64) uint64 {
	return 0
}

// CledF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcled_f64'.
// Requires NEON.
func CledF64(a float64, b float64) uint64 {
	return 0
}

// ClezF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_f32'.
// Requires NEON.
func ClezF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ClezF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_f64'.
// Requires NEON.
func ClezF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ClezS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s8'.
// Requires NEON.
func ClezS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ClezS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s16'.
// Requires NEON.
func ClezS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ClezS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s32'.
// Requires NEON.
func ClezS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ClezS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclez_s64'.
// Requires NEON.
func ClezS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ClezqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_f32'.
// Requires NEON.
func ClezqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ClezqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_f64'.
// Requires NEON.
func ClezqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ClezqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s8'.
// Requires NEON.
func ClezqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ClezqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s16'.
// Requires NEON.
func ClezqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ClezqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s32'.
// Requires NEON.
func ClezqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ClezqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezq_s64'.
// Requires NEON.
func ClezqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ClezsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclezs_f32'.
// Requires NEON.
func ClezsF32(a float32) uint32 {
	return 0
}

// ClezdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezd_s64'.
// Requires NEON.
func ClezdS64(a int64) uint64 {
	return 0
}

// ClezdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclezd_f64'.
// Requires NEON.
func ClezdF64(a float64) uint64 {
	return 0
}

// CltF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_f32'.
// Requires NEON.
func CltF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CltF64: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_f64'.
// Requires NEON.
func CltF64(a arm.Float64x1, b arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CltS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s8'.
// Requires NEON.
func CltS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CltS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s16'.
// Requires NEON.
func CltS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CltS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s32'.
// Requires NEON.
func CltS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CltS64: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_s64'.
// Requires NEON.
func CltS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CltU8: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u8'.
// Requires NEON.
func CltU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CltU16: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u16'.
// Requires NEON.
func CltU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CltU32: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u32'.
// Requires NEON.
func CltU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CltU64: ARM NEON intrinsic 
//
// Intrinsic: 'vclt_u64'.
// Requires NEON.
func CltU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CltqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_f32'.
// Requires NEON.
func CltqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CltqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_f64'.
// Requires NEON.
func CltqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CltqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s8'.
// Requires NEON.
func CltqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CltqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s16'.
// Requires NEON.
func CltqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CltqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s32'.
// Requires NEON.
func CltqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CltqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_s64'.
// Requires NEON.
func CltqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CltqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u8'.
// Requires NEON.
func CltqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CltqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u16'.
// Requires NEON.
func CltqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CltqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u32'.
// Requires NEON.
func CltqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CltqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltq_u64'.
// Requires NEON.
func CltqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CltsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vclts_f32'.
// Requires NEON.
func CltsF32(a float32, b float32) uint32 {
	return 0
}

// CltdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltd_s64'.
// Requires NEON.
func CltdS64(a int64, b int64) uint64 {
	return 0
}

// CltdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltd_u64'.
// Requires NEON.
func CltdU64(a uint64, b uint64) uint64 {
	return 0
}

// CltdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltd_f64'.
// Requires NEON.
func CltdF64(a float64, b float64) uint64 {
	return 0
}

// CltzF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_f32'.
// Requires NEON.
func CltzF32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CltzF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_f64'.
// Requires NEON.
func CltzF64(a arm.Float64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CltzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s8'.
// Requires NEON.
func CltzS8(a arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CltzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s16'.
// Requires NEON.
func CltzS16(a arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// CltzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s32'.
// Requires NEON.
func CltzS32(a arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CltzS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltz_s64'.
// Requires NEON.
func CltzS64(a arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// CltzqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_f32'.
// Requires NEON.
func CltzqF32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CltzqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_f64'.
// Requires NEON.
func CltzqF64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CltzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s8'.
// Requires NEON.
func CltzqS8(a arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CltzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s16'.
// Requires NEON.
func CltzqS16(a arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// CltzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s32'.
// Requires NEON.
func CltzqS32(a arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CltzqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzq_s64'.
// Requires NEON.
func CltzqS64(a arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CltzsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzs_f32'.
// Requires NEON.
func CltzsF32(a float32) uint32 {
	return 0
}

// CltzdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzd_s64'.
// Requires NEON.
func CltzdS64(a int64) uint64 {
	return 0
}

// CltzdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vcltzd_f64'.
// Requires NEON.
func CltzdF64(a float64) uint64 {
	return 0
}

// ClsS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcls_s8'.
// Requires NEON.
func ClsS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ClsS16: ARM NEON intrinsic 
//
// Intrinsic: 'vcls_s16'.
// Requires NEON.
func ClsS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ClsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vcls_s32'.
// Requires NEON.
func ClsS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ClsqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclsq_s8'.
// Requires NEON.
func ClsqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ClsqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclsq_s16'.
// Requires NEON.
func ClsqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ClsqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclsq_s32'.
// Requires NEON.
func ClsqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ClzS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_s8'.
// Requires NEON.
func ClzS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ClzS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_s16'.
// Requires NEON.
func ClzS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ClzS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_s32'.
// Requires NEON.
func ClzS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ClzU8: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_u8'.
// Requires NEON.
func ClzU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ClzU16: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_u16'.
// Requires NEON.
func ClzU16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ClzU32: ARM NEON intrinsic 
//
// Intrinsic: 'vclz_u32'.
// Requires NEON.
func ClzU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ClzqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_s8'.
// Requires NEON.
func ClzqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ClzqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_s16'.
// Requires NEON.
func ClzqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ClzqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_s32'.
// Requires NEON.
func ClzqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ClzqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_u8'.
// Requires NEON.
func ClzqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ClzqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_u16'.
// Requires NEON.
func ClzqU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ClzqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vclzq_u32'.
// Requires NEON.
func ClzqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CntP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcnt_p8'.
// Requires NEON.
func CntP8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// CntS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcnt_s8'.
// Requires NEON.
func CntS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// CntU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcnt_u8'.
// Requires NEON.
func CntU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// CntqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vcntq_p8'.
// Requires NEON.
func CntqP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// CntqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vcntq_s8'.
// Requires NEON.
func CntqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// CntqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vcntq_u8'.
// Requires NEON.
func CntqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// CvtF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f32_f64'.
// Requires NEON.
func CvtF32F64(a arm.Float64x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// CvtHighF32F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_high_f32_f64'.
// Requires NEON.
func CvtHighF32F64(a arm.Float32x2, b arm.Float64x2) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// CvtF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f64_f32'.
// Requires NEON.
func CvtF64F32(a arm.Float32x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// CvtHighF64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_high_f64_f32'.
// Requires NEON.
func CvtHighF64F32(a arm.Float32x4) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// CvtdF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_f64_s64'.
// Requires NEON.
func CvtdF64S64(a int64) float64 {
	return 0
}

// CvtdF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_f64_u64'.
// Requires NEON.
func CvtdF64U64(a uint64) float64 {
	return 0
}

// CvtsF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_f32_s32'.
// Requires NEON.
func CvtsF32S32(a int32) float32 {
	return 0
}

// CvtsF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_f32_u32'.
// Requires NEON.
func CvtsF32U32(a uint32) float32 {
	return 0
}

// CvtF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f32_s32'.
// Requires NEON.
func CvtF32S32(a arm.Int32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// CvtF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_f32_u32'.
// Requires NEON.
func CvtF32U32(a arm.Uint32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// CvtqF32S32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f32_s32'.
// Requires NEON.
func CvtqF32S32(a arm.Int32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// CvtqF32U32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f32_u32'.
// Requires NEON.
func CvtqF32U32(a arm.Uint32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// CvtqF64S64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f64_s64'.
// Requires NEON.
func CvtqF64S64(a arm.Int64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// CvtqF64U64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_f64_u64'.
// Requires NEON.
func CvtqF64U64(a arm.Uint64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// CvtdS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_s64_f64'.
// Requires NEON.
func CvtdS64F64(a float64) int64 {
	return 0
}

// CvtdU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtd_u64_f64'.
// Requires NEON.
func CvtdU64F64(a float64) uint64 {
	return 0
}

// CvtsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_s32_f32'.
// Requires NEON.
func CvtsS32F32(a float32) int32 {
	return 0
}

// CvtsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvts_u32_f32'.
// Requires NEON.
func CvtsU32F32(a float32) uint32 {
	return 0
}

// CvtS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_s32_f32'.
// Requires NEON.
func CvtS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// CvtU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvt_u32_f32'.
// Requires NEON.
func CvtU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CvtqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_s32_f32'.
// Requires NEON.
func CvtqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CvtqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_u32_f32'.
// Requires NEON.
func CvtqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CvtqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_s64_f64'.
// Requires NEON.
func CvtqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// CvtqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtq_u64_f64'.
// Requires NEON.
func CvtqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CvtadS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtad_s64_f64'.
// Requires NEON.
func CvtadS64F64(a float64) int64 {
	return 0
}

// CvtadU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtad_u64_f64'.
// Requires NEON.
func CvtadU64F64(a float64) uint64 {
	return 0
}

// CvtasS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtas_s32_f32'.
// Requires NEON.
func CvtasS32F32(a float32) int32 {
	return 0
}

// CvtasU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtas_u32_f32'.
// Requires NEON.
func CvtasU32F32(a float32) uint32 {
	return 0
}

// CvtaS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvta_s32_f32'.
// Requires NEON.
func CvtaS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// CvtaU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvta_u32_f32'.
// Requires NEON.
func CvtaU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CvtaqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_s32_f32'.
// Requires NEON.
func CvtaqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CvtaqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_u32_f32'.
// Requires NEON.
func CvtaqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CvtaqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_s64_f64'.
// Requires NEON.
func CvtaqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// CvtaqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtaq_u64_f64'.
// Requires NEON.
func CvtaqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CvtmdS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmd_s64_f64'.
// Requires NEON.
func CvtmdS64F64(a float64) int64 {
	return 0
}

// CvtmdU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmd_u64_f64'.
// Requires NEON.
func CvtmdU64F64(a float64) uint64 {
	return 0
}

// CvtmsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtms_s32_f32'.
// Requires NEON.
func CvtmsS32F32(a float32) int32 {
	return 0
}

// CvtmsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtms_u32_f32'.
// Requires NEON.
func CvtmsU32F32(a float32) uint32 {
	return 0
}

// CvtmS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtm_s32_f32'.
// Requires NEON.
func CvtmS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// CvtmU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtm_u32_f32'.
// Requires NEON.
func CvtmU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CvtmqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_s32_f32'.
// Requires NEON.
func CvtmqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CvtmqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_u32_f32'.
// Requires NEON.
func CvtmqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CvtmqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_s64_f64'.
// Requires NEON.
func CvtmqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// CvtmqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtmq_u64_f64'.
// Requires NEON.
func CvtmqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CvtndS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnd_s64_f64'.
// Requires NEON.
func CvtndS64F64(a float64) int64 {
	return 0
}

// CvtndU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnd_u64_f64'.
// Requires NEON.
func CvtndU64F64(a float64) uint64 {
	return 0
}

// CvtnsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtns_s32_f32'.
// Requires NEON.
func CvtnsS32F32(a float32) int32 {
	return 0
}

// CvtnsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtns_u32_f32'.
// Requires NEON.
func CvtnsU32F32(a float32) uint32 {
	return 0
}

// CvtnS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtn_s32_f32'.
// Requires NEON.
func CvtnS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// CvtnU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtn_u32_f32'.
// Requires NEON.
func CvtnU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CvtnqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_s32_f32'.
// Requires NEON.
func CvtnqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CvtnqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_u32_f32'.
// Requires NEON.
func CvtnqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CvtnqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_s64_f64'.
// Requires NEON.
func CvtnqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// CvtnqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtnq_u64_f64'.
// Requires NEON.
func CvtnqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// CvtpdS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpd_s64_f64'.
// Requires NEON.
func CvtpdS64F64(a float64) int64 {
	return 0
}

// CvtpdU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpd_u64_f64'.
// Requires NEON.
func CvtpdU64F64(a float64) uint64 {
	return 0
}

// CvtpsS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtps_s32_f32'.
// Requires NEON.
func CvtpsS32F32(a float32) int32 {
	return 0
}

// CvtpsU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtps_u32_f32'.
// Requires NEON.
func CvtpsU32F32(a float32) uint32 {
	return 0
}

// CvtpS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtp_s32_f32'.
// Requires NEON.
func CvtpS32F32(a arm.Float32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// CvtpU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtp_u32_f32'.
// Requires NEON.
func CvtpU32F32(a arm.Float32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// CvtpqS32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_s32_f32'.
// Requires NEON.
func CvtpqS32F32(a arm.Float32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// CvtpqU32F32: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_u32_f32'.
// Requires NEON.
func CvtpqU32F32(a arm.Float32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// CvtpqS64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_s64_f64'.
// Requires NEON.
func CvtpqS64F64(a arm.Float64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// CvtpqU64F64: ARM NEON intrinsic 
//
// Intrinsic: 'vcvtpq_u64_f64'.
// Requires NEON.
func CvtpqU64F64(a arm.Float64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// DupNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_f32'.
// Requires NEON.
func DupNF32(a float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// DupNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_f64'.
// Requires NEON.
func DupNF64(a float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// DupNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_p8'.
// Requires NEON.
func DupNP8(a arm.Poly8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// DupNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_p16'.
// Requires NEON.
func DupNP16(a arm.Poly16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// DupNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s8'.
// Requires NEON.
func DupNS8(a int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// DupNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s16'.
// Requires NEON.
func DupNS16(a int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// DupNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s32'.
// Requires NEON.
func DupNS32(a int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// DupNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_s64'.
// Requires NEON.
func DupNS64(a int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// DupNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u8'.
// Requires NEON.
func DupNU8(a uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// DupNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u16'.
// Requires NEON.
func DupNU16(a uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// DupNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u32'.
// Requires NEON.
func DupNU32(a uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// DupNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_n_u64'.
// Requires NEON.
func DupNU64(a uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// DupqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_f32'.
// Requires NEON.
func DupqNF32(a float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// DupqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_f64'.
// Requires NEON.
func DupqNF64(a float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// DupqNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_p8'.
// Requires NEON.
func DupqNP8(a uint32) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// DupqNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_p16'.
// Requires NEON.
func DupqNP16(a uint32) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// DupqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s8'.
// Requires NEON.
func DupqNS8(a int32) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// DupqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s16'.
// Requires NEON.
func DupqNS16(a int32) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// DupqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s32'.
// Requires NEON.
func DupqNS32(a int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// DupqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_s64'.
// Requires NEON.
func DupqNS64(a int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// DupqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u8'.
// Requires NEON.
func DupqNU8(a uint32) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// DupqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u16'.
// Requires NEON.
func DupqNU16(a uint32) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// DupqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u32'.
// Requires NEON.
func DupqNU32(a uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// DupqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_n_u64'.
// Requires NEON.
func DupqNU64(a uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// DupLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneF32(a arm.Float32x2, b int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// DupLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneF64(a arm.Float64x1, b int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// DupLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// DupLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// DupLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// DupLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// DupLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// DupLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// DupLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// DupLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// DupLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// DupLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// DupLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqF32(a arm.Float32x4, b int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// DupLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqF64(a arm.Float64x2, b int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// DupLaneqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqP8(a arm.Poly8x16, b int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// DupLaneqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqP16(a arm.Poly16x8, b int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// DupLaneqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqS8(a arm.Int8x16, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// DupLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqS16(a arm.Int16x8, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// DupLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqS32(a arm.Int32x4, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// DupLaneqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqS64(a arm.Int64x2, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// DupLaneqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqU8(a arm.Uint8x16, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// DupLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqU16(a arm.Uint16x8, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// DupLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqU32(a arm.Uint32x4, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// DupLaneqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdup_laneq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupLaneqU64(a arm.Uint64x2, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// DupqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneF32(a arm.Float32x2, b int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// DupqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneF64(a arm.Float64x1, b int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// DupqLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// DupqLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// DupqLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneS8(a arm.Int8x8, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// DupqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneS16(a arm.Int16x4, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// DupqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneS32(a arm.Int32x2, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// DupqLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneS64(a arm.Int64x1, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// DupqLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneU8(a arm.Uint8x8, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// DupqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneU16(a arm.Uint16x4, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// DupqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneU32(a arm.Uint32x2, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// DupqLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneU64(a arm.Uint64x1, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// DupqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqF32(a arm.Float32x4, b int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// DupqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqF64(a arm.Float64x2, b int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// DupqLaneqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqP8(a arm.Poly8x16, b int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// DupqLaneqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqP16(a arm.Poly16x8, b int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// DupqLaneqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// DupqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// DupqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// DupqLaneqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// DupqLaneqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// DupqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// DupqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// DupqLaneqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupq_laneq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupqLaneqU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// DupbLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_lane_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupbLaneP8(a arm.Poly8x8, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// DupbLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_lane_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupbLaneS8(a arm.Int8x8, b int) int8 {
	return 0
}

// DupbLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_lane_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupbLaneU8(a arm.Uint8x8, b int) uint8 {
	return 0
}

// DuphLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_lane_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DuphLaneP16(a arm.Poly16x4, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// DuphLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DuphLaneS16(a arm.Int16x4, b int) int16 {
	return 0
}

// DuphLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DuphLaneU16(a arm.Uint16x4, b int) uint16 {
	return 0
}

// DupsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupsLaneF32(a arm.Float32x2, b int) float32 {
	return 0
}

// DupsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupsLaneS32(a arm.Int32x2, b int) int32 {
	return 0
}

// DupsLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupsLaneU32(a arm.Uint32x2, b int) uint32 {
	return 0
}

// DupdLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupdLaneF64(a arm.Float64x1, b int) float64 {
	return 0
}

// DupdLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_lane_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupdLaneS64(a arm.Int64x1, b int) int64 {
	return 0
}

// DupdLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_lane_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupdLaneU64(a arm.Uint64x1, b int) uint64 {
	return 0
}

// DupbLaneqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_laneq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupbLaneqP8(a arm.Poly8x16, b int) (dst arm.Poly8) {
	return arm.Poly8{}
}

// DupbLaneqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_laneq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupbLaneqS8(a arm.Int8x16, b int) int8 {
	return 0
}

// DupbLaneqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vdupb_laneq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupbLaneqU8(a arm.Uint8x16, b int) uint8 {
	return 0
}

// DuphLaneqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_laneq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DuphLaneqP16(a arm.Poly16x8, b int) (dst arm.Poly16) {
	return arm.Poly16{}
}

// DuphLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DuphLaneqS16(a arm.Int16x8, b int) int16 {
	return 0
}

// DuphLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vduph_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DuphLaneqU16(a arm.Uint16x8, b int) uint16 {
	return 0
}

// DupsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupsLaneqF32(a arm.Float32x4, b int) float32 {
	return 0
}

// DupsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupsLaneqS32(a arm.Int32x4, b int) int32 {
	return 0
}

// DupsLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vdups_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupsLaneqU32(a arm.Uint32x4, b int) uint32 {
	return 0
}

// DupdLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupdLaneqF64(a arm.Float64x2, b int) float64 {
	return 0
}

// DupdLaneqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_laneq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupdLaneqS64(a arm.Int64x2, b int) int64 {
	return 0
}

// DupdLaneqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vdupd_laneq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func DupdLaneqU64(a arm.Uint64x2, b int) uint64 {
	return 0
}

// ExtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vext_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtF32(a arm.Float32x2, b arm.Float32x2, c int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// ExtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vext_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtF64(a arm.Float64x1, b arm.Float64x1, c int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// ExtP8: ARM NEON intrinsic 
//
// Intrinsic: 'vext_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtP8(a arm.Poly8x8, b arm.Poly8x8, c int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// ExtP16: ARM NEON intrinsic 
//
// Intrinsic: 'vext_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtP16(a arm.Poly16x4, b arm.Poly16x4, c int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// ExtS8: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ExtS16: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ExtS32: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ExtS64: ARM NEON intrinsic 
//
// Intrinsic: 'vext_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ExtU8: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ExtU16: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ExtU32: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ExtU64: ARM NEON intrinsic 
//
// Intrinsic: 'vext_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ExtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqF32(a arm.Float32x4, b arm.Float32x4, c int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// ExtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqF64(a arm.Float64x2, b arm.Float64x2, c int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// ExtqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_p8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqP8(a arm.Poly8x16, b arm.Poly8x16, c int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// ExtqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_p16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqP16(a arm.Poly16x8, b arm.Poly16x8, c int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// ExtqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ExtqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ExtqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ExtqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ExtqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ExtqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ExtqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ExtqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vextq_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ExtqU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// FmaF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_f64'.
// Requires NEON.
func FmaF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// FmaF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_f32'.
// Requires NEON.
func FmaF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmaqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_f32'.
// Requires NEON.
func FmaqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmaqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_f64'.
// Requires NEON.
func FmaqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// FmaNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_n_f32'.
// Requires NEON.
func FmaNF32(a arm.Float32x2, b arm.Float32x2, c float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmaqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_n_f32'.
// Requires NEON.
func FmaqNF32(a arm.Float32x4, b arm.Float32x4, c float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmaqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_n_f64'.
// Requires NEON.
func FmaqNF64(a arm.Float64x2, b arm.Float64x2, c float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// FmaLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmaLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaLaneF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// FmadLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmad_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmadLaneF64(a float64, b float64, c arm.Float64x1, lane int) float64 {
	return 0
}

// FmasLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmas_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmasLaneF32(a float32, b float32, c arm.Float32x2, lane int) float32 {
	return 0
}

// FmaLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmaLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfma_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaLaneqF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x2, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// FmadLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmad_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmadLaneqF64(a float64, b float64, c arm.Float64x2, lane int) float64 {
	return 0
}

// FmasLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmas_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmasLaneqF32(a float32, b float32, c arm.Float32x4, lane int) float32 {
	return 0
}

// FmaqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmaqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaqLaneF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x1, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// FmaqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmaqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmaq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmaqLaneqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// FmsF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_f64'.
// Requires NEON.
func FmsF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// FmsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_f32'.
// Requires NEON.
func FmsF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_f32'.
// Requires NEON.
func FmsqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_f64'.
// Requires NEON.
func FmsqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// FmsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmsLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsLaneF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// FmsdLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsd_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsdLaneF64(a float64, b float64, c arm.Float64x1, lane int) float64 {
	return 0
}

// FmssLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmss_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmssLaneF32(a float32, b float32, c arm.Float32x2, lane int) float32 {
	return 0
}

// FmsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// FmsLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfms_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsLaneqF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x2, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// FmsdLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsd_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsdLaneqF64(a float64, b float64, c arm.Float64x2, lane int) float64 {
	return 0
}

// FmssLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmss_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmssLaneqF32(a float32, b float32, c arm.Float32x4, lane int) float32 {
	return 0
}

// FmsqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmsqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsqLaneF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x1, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// FmsqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// FmsqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vfmsq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func FmsqLaneqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Ld1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1F32(a *float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Ld1F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1F64(a *float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// Ld1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1P8(a *arm.Poly8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Ld1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1P16(a *arm.Poly16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Ld1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1S8(a *int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Ld1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1S16(a *int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Ld1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1S32(a *int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Ld1S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1S64(a *int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// Ld1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1U8(a *uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Ld1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1U16(a *uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Ld1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1U32(a *uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Ld1U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1U64(a *uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// Ld1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qF32(a *float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Ld1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qF64(a *float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Ld1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qP8(a *arm.Poly8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Ld1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qP16(a *arm.Poly16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Ld1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qS8(a *int8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Ld1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qS16(a *int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Ld1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qS32(a *int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Ld1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qS64(a *int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Ld1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qU8(a *uint8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Ld1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qU16(a *uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Ld1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qU32(a *uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Ld1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qU64(a *uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Ld1DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupF32(a *float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Ld1DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupF64(a *float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// Ld1DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupP8(a *uint8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Ld1DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupP16(a *uint16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Ld1DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupS8(a *int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Ld1DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupS16(a *int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Ld1DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupS32(a *int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Ld1DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupS64(a *int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// Ld1DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupU8(a *uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Ld1DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupU16(a *uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Ld1DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupU32(a *uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Ld1DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1DupU64(a *uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// Ld1qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupF32(a *float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Ld1qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupF64(a *float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Ld1qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupP8(a *uint8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Ld1qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupP16(a *uint16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Ld1qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupS8(a *int8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Ld1qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupS16(a *int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Ld1qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupS32(a *int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Ld1qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupS64(a *int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Ld1qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupU8(a *uint8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Ld1qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupU16(a *uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Ld1qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupU32(a *uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Ld1qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qDupU64(a *uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Ld1LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneF32(src *float32, vec arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Ld1LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneF64(src *float64, vec arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// Ld1LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneP8(src *arm.Poly8, vec arm.Poly8x8, lane int) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Ld1LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneP16(src *arm.Poly16, vec arm.Poly16x4, lane int) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Ld1LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneS8(src *int8, vec arm.Int8x8, lane int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Ld1LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneS16(src *int16, vec arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Ld1LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneS32(src *int32, vec arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Ld1LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneS64(src *int64, vec arm.Int64x1, lane int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// Ld1LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneU8(src *uint8, vec arm.Uint8x8, lane int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Ld1LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneU16(src *uint16, vec arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Ld1LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneU32(src *uint32, vec arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Ld1LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1LaneU64(src *uint64, vec arm.Uint64x1, lane int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// Ld1qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneF32(src *float32, vec arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Ld1qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneF64(src *float64, vec arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Ld1qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneP8(src *arm.Poly8, vec arm.Poly8x16, lane int) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Ld1qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneP16(src *arm.Poly16, vec arm.Poly16x8, lane int) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Ld1qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneS8(src *int8, vec arm.Int8x16, lane int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Ld1qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneS16(src *int16, vec arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Ld1qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneS32(src *int32, vec arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Ld1qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneS64(src *int64, vec arm.Int64x2, lane int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Ld1qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneU8(src *uint8, vec arm.Uint8x16, lane int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Ld1qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneU16(src *uint16, vec arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Ld1qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneU32(src *uint32, vec arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Ld1qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld1q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld1qLaneU64(src *uint64, vec arm.Uint64x2, lane int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Ld2S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2S64(a *int64) (dst [2]arm.Int64x1) {
	return [2]arm.Int64x1{}
}

// Ld2U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2U64(a *uint64) (dst [2]arm.Uint64x1) {
	return [2]arm.Uint64x1{}
}

// Ld2F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2F64(a *float64) (dst [2]arm.Float64x1) {
	return [2]arm.Float64x1{}
}

// Ld2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2S8(a *int8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// Ld2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2P8(a *arm.Poly8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// Ld2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2S16(a *int16) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// Ld2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2P16(a *arm.Poly16) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// Ld2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2S32(a *int32) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// Ld2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2U8(a *uint8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// Ld2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2U16(a *uint16) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// Ld2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2U32(a *uint32) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// Ld2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2F32(a *float32) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// Ld2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qS8(a *int8) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// Ld2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qP8(a *arm.Poly8) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// Ld2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qS16(a *int16) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// Ld2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qP16(a *arm.Poly16) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// Ld2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qS32(a *int32) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// Ld2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qS64(a *int64) (dst [2]arm.Int64x2) {
	return [2]arm.Int64x2{}
}

// Ld2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qU8(a *uint8) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// Ld2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qU16(a *uint16) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// Ld2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qU32(a *uint32) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Ld2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qU64(a *uint64) (dst [2]arm.Uint64x2) {
	return [2]arm.Uint64x2{}
}

// Ld2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qF32(a *float32) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// Ld2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qF64(a *float64) (dst [2]arm.Float64x2) {
	return [2]arm.Float64x2{}
}

// Ld3S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3S64(a *int64) (dst [3]arm.Int64x1) {
	return [3]arm.Int64x1{}
}

// Ld3U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3U64(a *uint64) (dst [3]arm.Uint64x1) {
	return [3]arm.Uint64x1{}
}

// Ld3F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3F64(a *float64) (dst [3]arm.Float64x1) {
	return [3]arm.Float64x1{}
}

// Ld3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3S8(a *int8) (dst [3]arm.Int8x8) {
	return [3]arm.Int8x8{}
}

// Ld3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3P8(a *arm.Poly8) (dst [3]arm.Poly8x8) {
	return [3]arm.Poly8x8{}
}

// Ld3S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3S16(a *int16) (dst [3]arm.Int16x4) {
	return [3]arm.Int16x4{}
}

// Ld3P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3P16(a *arm.Poly16) (dst [3]arm.Poly16x4) {
	return [3]arm.Poly16x4{}
}

// Ld3S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3S32(a *int32) (dst [3]arm.Int32x2) {
	return [3]arm.Int32x2{}
}

// Ld3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3U8(a *uint8) (dst [3]arm.Uint8x8) {
	return [3]arm.Uint8x8{}
}

// Ld3U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3U16(a *uint16) (dst [3]arm.Uint16x4) {
	return [3]arm.Uint16x4{}
}

// Ld3U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3U32(a *uint32) (dst [3]arm.Uint32x2) {
	return [3]arm.Uint32x2{}
}

// Ld3F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3F32(a *float32) (dst [3]arm.Float32x2) {
	return [3]arm.Float32x2{}
}

// Ld3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qS8(a *int8) (dst [3]arm.Int8x16) {
	return [3]arm.Int8x16{}
}

// Ld3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qP8(a *arm.Poly8) (dst [3]arm.Poly8x16) {
	return [3]arm.Poly8x16{}
}

// Ld3qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qS16(a *int16) (dst [3]arm.Int16x8) {
	return [3]arm.Int16x8{}
}

// Ld3qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qP16(a *arm.Poly16) (dst [3]arm.Poly16x8) {
	return [3]arm.Poly16x8{}
}

// Ld3qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qS32(a *int32) (dst [3]arm.Int32x4) {
	return [3]arm.Int32x4{}
}

// Ld3qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qS64(a *int64) (dst [3]arm.Int64x2) {
	return [3]arm.Int64x2{}
}

// Ld3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qU8(a *uint8) (dst [3]arm.Uint8x16) {
	return [3]arm.Uint8x16{}
}

// Ld3qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qU16(a *uint16) (dst [3]arm.Uint16x8) {
	return [3]arm.Uint16x8{}
}

// Ld3qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qU32(a *uint32) (dst [3]arm.Uint32x4) {
	return [3]arm.Uint32x4{}
}

// Ld3qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qU64(a *uint64) (dst [3]arm.Uint64x2) {
	return [3]arm.Uint64x2{}
}

// Ld3qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qF32(a *float32) (dst [3]arm.Float32x4) {
	return [3]arm.Float32x4{}
}

// Ld3qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qF64(a *float64) (dst [3]arm.Float64x2) {
	return [3]arm.Float64x2{}
}

// Ld4S64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4S64(a *int64) (dst [4]arm.Int64x1) {
	return [4]arm.Int64x1{}
}

// Ld4U64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4U64(a *uint64) (dst [4]arm.Uint64x1) {
	return [4]arm.Uint64x1{}
}

// Ld4F64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4F64(a *float64) (dst [4]arm.Float64x1) {
	return [4]arm.Float64x1{}
}

// Ld4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4S8(a *int8) (dst [4]arm.Int8x8) {
	return [4]arm.Int8x8{}
}

// Ld4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4P8(a *arm.Poly8) (dst [4]arm.Poly8x8) {
	return [4]arm.Poly8x8{}
}

// Ld4S16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4S16(a *int16) (dst [4]arm.Int16x4) {
	return [4]arm.Int16x4{}
}

// Ld4P16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4P16(a *arm.Poly16) (dst [4]arm.Poly16x4) {
	return [4]arm.Poly16x4{}
}

// Ld4S32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4S32(a *int32) (dst [4]arm.Int32x2) {
	return [4]arm.Int32x2{}
}

// Ld4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4U8(a *uint8) (dst [4]arm.Uint8x8) {
	return [4]arm.Uint8x8{}
}

// Ld4U16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4U16(a *uint16) (dst [4]arm.Uint16x4) {
	return [4]arm.Uint16x4{}
}

// Ld4U32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4U32(a *uint32) (dst [4]arm.Uint32x2) {
	return [4]arm.Uint32x2{}
}

// Ld4F32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4F32(a *float32) (dst [4]arm.Float32x2) {
	return [4]arm.Float32x2{}
}

// Ld4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qS8(a *int8) (dst [4]arm.Int8x16) {
	return [4]arm.Int8x16{}
}

// Ld4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qP8(a *arm.Poly8) (dst [4]arm.Poly8x16) {
	return [4]arm.Poly8x16{}
}

// Ld4qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qS16(a *int16) (dst [4]arm.Int16x8) {
	return [4]arm.Int16x8{}
}

// Ld4qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qP16(a *arm.Poly16) (dst [4]arm.Poly16x8) {
	return [4]arm.Poly16x8{}
}

// Ld4qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qS32(a *int32) (dst [4]arm.Int32x4) {
	return [4]arm.Int32x4{}
}

// Ld4qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qS64(a *int64) (dst [4]arm.Int64x2) {
	return [4]arm.Int64x2{}
}

// Ld4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qU8(a *uint8) (dst [4]arm.Uint8x16) {
	return [4]arm.Uint8x16{}
}

// Ld4qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qU16(a *uint16) (dst [4]arm.Uint16x8) {
	return [4]arm.Uint16x8{}
}

// Ld4qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qU32(a *uint32) (dst [4]arm.Uint32x4) {
	return [4]arm.Uint32x4{}
}

// Ld4qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qU64(a *uint64) (dst [4]arm.Uint64x2) {
	return [4]arm.Uint64x2{}
}

// Ld4qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qF32(a *float32) (dst [4]arm.Float32x4) {
	return [4]arm.Float32x4{}
}

// Ld4qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qF64(a *float64) (dst [4]arm.Float64x2) {
	return [4]arm.Float64x2{}
}

// Ld2DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupS8(a *int8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// Ld2DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupS16(a *int16) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// Ld2DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupS32(a *int32) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// Ld2DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupF32(a *float32) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// Ld2DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupF64(a *float64) (dst [2]arm.Float64x1) {
	return [2]arm.Float64x1{}
}

// Ld2DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupU8(a *uint8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// Ld2DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupU16(a *uint16) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// Ld2DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupU32(a *uint32) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// Ld2DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupP8(a *arm.Poly8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// Ld2DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupP16(a *arm.Poly16) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// Ld2DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupS64(a *int64) (dst [2]arm.Int64x1) {
	return [2]arm.Int64x1{}
}

// Ld2DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2DupU64(a *uint64) (dst [2]arm.Uint64x1) {
	return [2]arm.Uint64x1{}
}

// Ld2qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupS8(a *int8) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// Ld2qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupP8(a *arm.Poly8) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// Ld2qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupS16(a *int16) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// Ld2qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupP16(a *arm.Poly16) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// Ld2qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupS32(a *int32) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// Ld2qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupS64(a *int64) (dst [2]arm.Int64x2) {
	return [2]arm.Int64x2{}
}

// Ld2qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupU8(a *uint8) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// Ld2qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupU16(a *uint16) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// Ld2qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupU32(a *uint32) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Ld2qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupU64(a *uint64) (dst [2]arm.Uint64x2) {
	return [2]arm.Uint64x2{}
}

// Ld2qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupF32(a *float32) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// Ld2qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qDupF64(a *float64) (dst [2]arm.Float64x2) {
	return [2]arm.Float64x2{}
}

// Ld3DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupS64(a *int64) (dst [3]arm.Int64x1) {
	return [3]arm.Int64x1{}
}

// Ld3DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupU64(a *uint64) (dst [3]arm.Uint64x1) {
	return [3]arm.Uint64x1{}
}

// Ld3DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupF64(a *float64) (dst [3]arm.Float64x1) {
	return [3]arm.Float64x1{}
}

// Ld3DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupS8(a *int8) (dst [3]arm.Int8x8) {
	return [3]arm.Int8x8{}
}

// Ld3DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupP8(a *arm.Poly8) (dst [3]arm.Poly8x8) {
	return [3]arm.Poly8x8{}
}

// Ld3DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupS16(a *int16) (dst [3]arm.Int16x4) {
	return [3]arm.Int16x4{}
}

// Ld3DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupP16(a *arm.Poly16) (dst [3]arm.Poly16x4) {
	return [3]arm.Poly16x4{}
}

// Ld3DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupS32(a *int32) (dst [3]arm.Int32x2) {
	return [3]arm.Int32x2{}
}

// Ld3DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupU8(a *uint8) (dst [3]arm.Uint8x8) {
	return [3]arm.Uint8x8{}
}

// Ld3DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupU16(a *uint16) (dst [3]arm.Uint16x4) {
	return [3]arm.Uint16x4{}
}

// Ld3DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupU32(a *uint32) (dst [3]arm.Uint32x2) {
	return [3]arm.Uint32x2{}
}

// Ld3DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3DupF32(a *float32) (dst [3]arm.Float32x2) {
	return [3]arm.Float32x2{}
}

// Ld3qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupS8(a *int8) (dst [3]arm.Int8x16) {
	return [3]arm.Int8x16{}
}

// Ld3qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupP8(a *arm.Poly8) (dst [3]arm.Poly8x16) {
	return [3]arm.Poly8x16{}
}

// Ld3qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupS16(a *int16) (dst [3]arm.Int16x8) {
	return [3]arm.Int16x8{}
}

// Ld3qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupP16(a *arm.Poly16) (dst [3]arm.Poly16x8) {
	return [3]arm.Poly16x8{}
}

// Ld3qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupS32(a *int32) (dst [3]arm.Int32x4) {
	return [3]arm.Int32x4{}
}

// Ld3qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupS64(a *int64) (dst [3]arm.Int64x2) {
	return [3]arm.Int64x2{}
}

// Ld3qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupU8(a *uint8) (dst [3]arm.Uint8x16) {
	return [3]arm.Uint8x16{}
}

// Ld3qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupU16(a *uint16) (dst [3]arm.Uint16x8) {
	return [3]arm.Uint16x8{}
}

// Ld3qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupU32(a *uint32) (dst [3]arm.Uint32x4) {
	return [3]arm.Uint32x4{}
}

// Ld3qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupU64(a *uint64) (dst [3]arm.Uint64x2) {
	return [3]arm.Uint64x2{}
}

// Ld3qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupF32(a *float32) (dst [3]arm.Float32x4) {
	return [3]arm.Float32x4{}
}

// Ld3qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qDupF64(a *float64) (dst [3]arm.Float64x2) {
	return [3]arm.Float64x2{}
}

// Ld4DupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupS64(a *int64) (dst [4]arm.Int64x1) {
	return [4]arm.Int64x1{}
}

// Ld4DupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupU64(a *uint64) (dst [4]arm.Uint64x1) {
	return [4]arm.Uint64x1{}
}

// Ld4DupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupF64(a *float64) (dst [4]arm.Float64x1) {
	return [4]arm.Float64x1{}
}

// Ld4DupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupS8(a *int8) (dst [4]arm.Int8x8) {
	return [4]arm.Int8x8{}
}

// Ld4DupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupP8(a *arm.Poly8) (dst [4]arm.Poly8x8) {
	return [4]arm.Poly8x8{}
}

// Ld4DupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupS16(a *int16) (dst [4]arm.Int16x4) {
	return [4]arm.Int16x4{}
}

// Ld4DupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupP16(a *arm.Poly16) (dst [4]arm.Poly16x4) {
	return [4]arm.Poly16x4{}
}

// Ld4DupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupS32(a *int32) (dst [4]arm.Int32x2) {
	return [4]arm.Int32x2{}
}

// Ld4DupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupU8(a *uint8) (dst [4]arm.Uint8x8) {
	return [4]arm.Uint8x8{}
}

// Ld4DupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupU16(a *uint16) (dst [4]arm.Uint16x4) {
	return [4]arm.Uint16x4{}
}

// Ld4DupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupU32(a *uint32) (dst [4]arm.Uint32x2) {
	return [4]arm.Uint32x2{}
}

// Ld4DupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4DupF32(a *float32) (dst [4]arm.Float32x2) {
	return [4]arm.Float32x2{}
}

// Ld4qDupS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupS8(a *int8) (dst [4]arm.Int8x16) {
	return [4]arm.Int8x16{}
}

// Ld4qDupP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupP8(a *arm.Poly8) (dst [4]arm.Poly8x16) {
	return [4]arm.Poly8x16{}
}

// Ld4qDupS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupS16(a *int16) (dst [4]arm.Int16x8) {
	return [4]arm.Int16x8{}
}

// Ld4qDupP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupP16(a *arm.Poly16) (dst [4]arm.Poly16x8) {
	return [4]arm.Poly16x8{}
}

// Ld4qDupS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupS32(a *int32) (dst [4]arm.Int32x4) {
	return [4]arm.Int32x4{}
}

// Ld4qDupS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupS64(a *int64) (dst [4]arm.Int64x2) {
	return [4]arm.Int64x2{}
}

// Ld4qDupU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupU8(a *uint8) (dst [4]arm.Uint8x16) {
	return [4]arm.Uint8x16{}
}

// Ld4qDupU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupU16(a *uint16) (dst [4]arm.Uint16x8) {
	return [4]arm.Uint16x8{}
}

// Ld4qDupU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupU32(a *uint32) (dst [4]arm.Uint32x4) {
	return [4]arm.Uint32x4{}
}

// Ld4qDupU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupU64(a *uint64) (dst [4]arm.Uint64x2) {
	return [4]arm.Uint64x2{}
}

// Ld4qDupF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupF32(a *float32) (dst [4]arm.Float32x4) {
	return [4]arm.Float32x4{}
}

// Ld4qDupF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_dup_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qDupF64(a *float64) (dst [4]arm.Float64x2) {
	return [4]arm.Float64x2{}
}

// Ld2LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneF32(ptr *float32, b [2]arm.Float32x2, c int) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// Ld2LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneF64(ptr *float64, b [2]arm.Float64x1, c int) (dst [2]arm.Float64x1) {
	return [2]arm.Float64x1{}
}

// Ld2LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneP8(ptr *arm.Poly8, b [2]arm.Poly8x8, c int) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// Ld2LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneP16(ptr *arm.Poly16, b [2]arm.Poly16x4, c int) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// Ld2LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneS8(ptr *int8, b [2]arm.Int8x8, c int) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// Ld2LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneS16(ptr *int16, b [2]arm.Int16x4, c int) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// Ld2LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneS32(ptr *int32, b [2]arm.Int32x2, c int) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// Ld2LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneS64(ptr *int64, b [2]arm.Int64x1, c int) (dst [2]arm.Int64x1) {
	return [2]arm.Int64x1{}
}

// Ld2LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneU8(ptr *uint8, b [2]arm.Uint8x8, c int) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// Ld2LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneU16(ptr *uint16, b [2]arm.Uint16x4, c int) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// Ld2LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneU32(ptr *uint32, b [2]arm.Uint32x2, c int) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// Ld2LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2LaneU64(ptr *uint64, b [2]arm.Uint64x1, c int) (dst [2]arm.Uint64x1) {
	return [2]arm.Uint64x1{}
}

// Ld2qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneF32(ptr *float32, b [2]arm.Float32x4, c int) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// Ld2qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneF64(ptr *float64, b [2]arm.Float64x2, c int) (dst [2]arm.Float64x2) {
	return [2]arm.Float64x2{}
}

// Ld2qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneP8(ptr *arm.Poly8, b [2]arm.Poly8x16, c int) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// Ld2qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneP16(ptr *arm.Poly16, b [2]arm.Poly16x8, c int) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// Ld2qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneS8(ptr *int8, b [2]arm.Int8x16, c int) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// Ld2qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneS16(ptr *int16, b [2]arm.Int16x8, c int) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// Ld2qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneS32(ptr *int32, b [2]arm.Int32x4, c int) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// Ld2qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneS64(ptr *int64, b [2]arm.Int64x2, c int) (dst [2]arm.Int64x2) {
	return [2]arm.Int64x2{}
}

// Ld2qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneU8(ptr *uint8, b [2]arm.Uint8x16, c int) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// Ld2qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneU16(ptr *uint16, b [2]arm.Uint16x8, c int) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// Ld2qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneU32(ptr *uint32, b [2]arm.Uint32x4, c int) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Ld2qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld2q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld2qLaneU64(ptr *uint64, b [2]arm.Uint64x2, c int) (dst [2]arm.Uint64x2) {
	return [2]arm.Uint64x2{}
}

// Ld3LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneF32(ptr *float32, b [3]arm.Float32x2, c int) (dst [3]arm.Float32x2) {
	return [3]arm.Float32x2{}
}

// Ld3LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneF64(ptr *float64, b [3]arm.Float64x1, c int) (dst [3]arm.Float64x1) {
	return [3]arm.Float64x1{}
}

// Ld3LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneP8(ptr *arm.Poly8, b [3]arm.Poly8x8, c int) (dst [3]arm.Poly8x8) {
	return [3]arm.Poly8x8{}
}

// Ld3LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneP16(ptr *arm.Poly16, b [3]arm.Poly16x4, c int) (dst [3]arm.Poly16x4) {
	return [3]arm.Poly16x4{}
}

// Ld3LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneS8(ptr *int8, b [3]arm.Int8x8, c int) (dst [3]arm.Int8x8) {
	return [3]arm.Int8x8{}
}

// Ld3LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneS16(ptr *int16, b [3]arm.Int16x4, c int) (dst [3]arm.Int16x4) {
	return [3]arm.Int16x4{}
}

// Ld3LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneS32(ptr *int32, b [3]arm.Int32x2, c int) (dst [3]arm.Int32x2) {
	return [3]arm.Int32x2{}
}

// Ld3LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneS64(ptr *int64, b [3]arm.Int64x1, c int) (dst [3]arm.Int64x1) {
	return [3]arm.Int64x1{}
}

// Ld3LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneU8(ptr *uint8, b [3]arm.Uint8x8, c int) (dst [3]arm.Uint8x8) {
	return [3]arm.Uint8x8{}
}

// Ld3LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneU16(ptr *uint16, b [3]arm.Uint16x4, c int) (dst [3]arm.Uint16x4) {
	return [3]arm.Uint16x4{}
}

// Ld3LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneU32(ptr *uint32, b [3]arm.Uint32x2, c int) (dst [3]arm.Uint32x2) {
	return [3]arm.Uint32x2{}
}

// Ld3LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3LaneU64(ptr *uint64, b [3]arm.Uint64x1, c int) (dst [3]arm.Uint64x1) {
	return [3]arm.Uint64x1{}
}

// Ld3qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneF32(ptr *float32, b [3]arm.Float32x4, c int) (dst [3]arm.Float32x4) {
	return [3]arm.Float32x4{}
}

// Ld3qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneF64(ptr *float64, b [3]arm.Float64x2, c int) (dst [3]arm.Float64x2) {
	return [3]arm.Float64x2{}
}

// Ld3qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneP8(ptr *arm.Poly8, b [3]arm.Poly8x16, c int) (dst [3]arm.Poly8x16) {
	return [3]arm.Poly8x16{}
}

// Ld3qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneP16(ptr *arm.Poly16, b [3]arm.Poly16x8, c int) (dst [3]arm.Poly16x8) {
	return [3]arm.Poly16x8{}
}

// Ld3qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneS8(ptr *int8, b [3]arm.Int8x16, c int) (dst [3]arm.Int8x16) {
	return [3]arm.Int8x16{}
}

// Ld3qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneS16(ptr *int16, b [3]arm.Int16x8, c int) (dst [3]arm.Int16x8) {
	return [3]arm.Int16x8{}
}

// Ld3qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneS32(ptr *int32, b [3]arm.Int32x4, c int) (dst [3]arm.Int32x4) {
	return [3]arm.Int32x4{}
}

// Ld3qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneS64(ptr *int64, b [3]arm.Int64x2, c int) (dst [3]arm.Int64x2) {
	return [3]arm.Int64x2{}
}

// Ld3qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneU8(ptr *uint8, b [3]arm.Uint8x16, c int) (dst [3]arm.Uint8x16) {
	return [3]arm.Uint8x16{}
}

// Ld3qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneU16(ptr *uint16, b [3]arm.Uint16x8, c int) (dst [3]arm.Uint16x8) {
	return [3]arm.Uint16x8{}
}

// Ld3qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneU32(ptr *uint32, b [3]arm.Uint32x4, c int) (dst [3]arm.Uint32x4) {
	return [3]arm.Uint32x4{}
}

// Ld3qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld3q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld3qLaneU64(ptr *uint64, b [3]arm.Uint64x2, c int) (dst [3]arm.Uint64x2) {
	return [3]arm.Uint64x2{}
}

// Ld4LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneF32(ptr *float32, b [4]arm.Float32x2, c int) (dst [4]arm.Float32x2) {
	return [4]arm.Float32x2{}
}

// Ld4LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneF64(ptr *float64, b [4]arm.Float64x1, c int) (dst [4]arm.Float64x1) {
	return [4]arm.Float64x1{}
}

// Ld4LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneP8(ptr *arm.Poly8, b [4]arm.Poly8x8, c int) (dst [4]arm.Poly8x8) {
	return [4]arm.Poly8x8{}
}

// Ld4LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneP16(ptr *arm.Poly16, b [4]arm.Poly16x4, c int) (dst [4]arm.Poly16x4) {
	return [4]arm.Poly16x4{}
}

// Ld4LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneS8(ptr *int8, b [4]arm.Int8x8, c int) (dst [4]arm.Int8x8) {
	return [4]arm.Int8x8{}
}

// Ld4LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneS16(ptr *int16, b [4]arm.Int16x4, c int) (dst [4]arm.Int16x4) {
	return [4]arm.Int16x4{}
}

// Ld4LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneS32(ptr *int32, b [4]arm.Int32x2, c int) (dst [4]arm.Int32x2) {
	return [4]arm.Int32x2{}
}

// Ld4LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneS64(ptr *int64, b [4]arm.Int64x1, c int) (dst [4]arm.Int64x1) {
	return [4]arm.Int64x1{}
}

// Ld4LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneU8(ptr *uint8, b [4]arm.Uint8x8, c int) (dst [4]arm.Uint8x8) {
	return [4]arm.Uint8x8{}
}

// Ld4LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneU16(ptr *uint16, b [4]arm.Uint16x4, c int) (dst [4]arm.Uint16x4) {
	return [4]arm.Uint16x4{}
}

// Ld4LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneU32(ptr *uint32, b [4]arm.Uint32x2, c int) (dst [4]arm.Uint32x2) {
	return [4]arm.Uint32x2{}
}

// Ld4LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4LaneU64(ptr *uint64, b [4]arm.Uint64x1, c int) (dst [4]arm.Uint64x1) {
	return [4]arm.Uint64x1{}
}

// Ld4qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneF32(ptr *float32, b [4]arm.Float32x4, c int) (dst [4]arm.Float32x4) {
	return [4]arm.Float32x4{}
}

// Ld4qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneF64(ptr *float64, b [4]arm.Float64x2, c int) (dst [4]arm.Float64x2) {
	return [4]arm.Float64x2{}
}

// Ld4qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneP8(ptr *arm.Poly8, b [4]arm.Poly8x16, c int) (dst [4]arm.Poly8x16) {
	return [4]arm.Poly8x16{}
}

// Ld4qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneP16(ptr *arm.Poly16, b [4]arm.Poly16x8, c int) (dst [4]arm.Poly16x8) {
	return [4]arm.Poly16x8{}
}

// Ld4qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneS8(ptr *int8, b [4]arm.Int8x16, c int) (dst [4]arm.Int8x16) {
	return [4]arm.Int8x16{}
}

// Ld4qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneS16(ptr *int16, b [4]arm.Int16x8, c int) (dst [4]arm.Int16x8) {
	return [4]arm.Int16x8{}
}

// Ld4qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneS32(ptr *int32, b [4]arm.Int32x4, c int) (dst [4]arm.Int32x4) {
	return [4]arm.Int32x4{}
}

// Ld4qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneS64(ptr *int64, b [4]arm.Int64x2, c int) (dst [4]arm.Int64x2) {
	return [4]arm.Int64x2{}
}

// Ld4qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneU8(ptr *uint8, b [4]arm.Uint8x16, c int) (dst [4]arm.Uint8x16) {
	return [4]arm.Uint8x16{}
}

// Ld4qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneU16(ptr *uint16, b [4]arm.Uint16x8, c int) (dst [4]arm.Uint16x8) {
	return [4]arm.Uint16x8{}
}

// Ld4qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneU32(ptr *uint32, b [4]arm.Uint32x4, c int) (dst [4]arm.Uint32x4) {
	return [4]arm.Uint32x4{}
}

// Ld4qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vld4q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func Ld4qLaneU64(ptr *uint64, b [4]arm.Uint64x2, c int) (dst [4]arm.Uint64x2) {
	return [4]arm.Uint64x2{}
}

// MaxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_f32'.
// Requires NEON.
func MaxF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MaxS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_s8'.
// Requires NEON.
func MaxS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MaxS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_s16'.
// Requires NEON.
func MaxS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MaxS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_s32'.
// Requires NEON.
func MaxS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MaxU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_u8'.
// Requires NEON.
func MaxU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MaxU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_u16'.
// Requires NEON.
func MaxU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MaxU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmax_u32'.
// Requires NEON.
func MaxU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MaxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_f32'.
// Requires NEON.
func MaxqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MaxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_f64'.
// Requires NEON.
func MaxqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MaxqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_s8'.
// Requires NEON.
func MaxqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MaxqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_s16'.
// Requires NEON.
func MaxqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MaxqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_s32'.
// Requires NEON.
func MaxqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MaxqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_u8'.
// Requires NEON.
func MaxqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MaxqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_u16'.
// Requires NEON.
func MaxqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MaxqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxq_u32'.
// Requires NEON.
func MaxqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PmaxS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_s8'.
// Requires NEON.
func PmaxS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// PmaxS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_s16'.
// Requires NEON.
func PmaxS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// PmaxS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_s32'.
// Requires NEON.
func PmaxS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// PmaxU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_u8'.
// Requires NEON.
func PmaxU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// PmaxU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_u16'.
// Requires NEON.
func PmaxU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// PmaxU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_u32'.
// Requires NEON.
func PmaxU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// PmaxqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_s8'.
// Requires NEON.
func PmaxqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// PmaxqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_s16'.
// Requires NEON.
func PmaxqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// PmaxqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_s32'.
// Requires NEON.
func PmaxqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// PmaxqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_u8'.
// Requires NEON.
func PmaxqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// PmaxqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_u16'.
// Requires NEON.
func PmaxqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// PmaxqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_u32'.
// Requires NEON.
func PmaxqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PmaxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmax_f32'.
// Requires NEON.
func PmaxF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// PmaxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_f32'.
// Requires NEON.
func PmaxqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// PmaxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxq_f64'.
// Requires NEON.
func PmaxqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// PmaxqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxqd_f64'.
// Requires NEON.
func PmaxqdF64(a arm.Float64x2) float64 {
	return 0
}

// PmaxsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxs_f32'.
// Requires NEON.
func PmaxsF32(a arm.Float32x2) float32 {
	return 0
}

// PmaxnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnm_f32'.
// Requires NEON.
func PmaxnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// PmaxnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnmq_f32'.
// Requires NEON.
func PmaxnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// PmaxnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnmq_f64'.
// Requires NEON.
func PmaxnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// PmaxnmqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnmqd_f64'.
// Requires NEON.
func PmaxnmqdF64(a arm.Float64x2) float64 {
	return 0
}

// PmaxnmsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmaxnms_f32'.
// Requires NEON.
func PmaxnmsF32(a arm.Float32x2) float32 {
	return 0
}

// PminS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_s8'.
// Requires NEON.
func PminS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// PminS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_s16'.
// Requires NEON.
func PminS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// PminS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_s32'.
// Requires NEON.
func PminS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// PminU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_u8'.
// Requires NEON.
func PminU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// PminU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_u16'.
// Requires NEON.
func PminU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// PminU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_u32'.
// Requires NEON.
func PminU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// PminqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_s8'.
// Requires NEON.
func PminqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// PminqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_s16'.
// Requires NEON.
func PminqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// PminqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_s32'.
// Requires NEON.
func PminqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// PminqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_u8'.
// Requires NEON.
func PminqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// PminqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_u16'.
// Requires NEON.
func PminqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// PminqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_u32'.
// Requires NEON.
func PminqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// PminF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmin_f32'.
// Requires NEON.
func PminF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// PminqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_f32'.
// Requires NEON.
func PminqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// PminqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminq_f64'.
// Requires NEON.
func PminqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// PminqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminqd_f64'.
// Requires NEON.
func PminqdF64(a arm.Float64x2) float64 {
	return 0
}

// PminsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpmins_f32'.
// Requires NEON.
func PminsF32(a arm.Float32x2) float32 {
	return 0
}

// PminnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnm_f32'.
// Requires NEON.
func PminnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// PminnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnmq_f32'.
// Requires NEON.
func PminnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// PminnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnmq_f64'.
// Requires NEON.
func PminnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// PminnmqdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnmqd_f64'.
// Requires NEON.
func PminnmqdF64(a arm.Float64x2) float64 {
	return 0
}

// PminnmsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vpminnms_f32'.
// Requires NEON.
func PminnmsF32(a arm.Float32x2) float32 {
	return 0
}

// MaxnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnm_f32'.
// Requires NEON.
func MaxnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MaxnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmq_f32'.
// Requires NEON.
func MaxnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MaxnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmq_f64'.
// Requires NEON.
func MaxnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MaxvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_f32'.
// Requires NEON.
func MaxvF32(a arm.Float32x2) float32 {
	return 0
}

// MaxvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_s8'.
// Requires NEON.
func MaxvS8(a arm.Int8x8) int8 {
	return 0
}

// MaxvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_s16'.
// Requires NEON.
func MaxvS16(a arm.Int16x4) int16 {
	return 0
}

// MaxvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_s32'.
// Requires NEON.
func MaxvS32(a arm.Int32x2) int32 {
	return 0
}

// MaxvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_u8'.
// Requires NEON.
func MaxvU8(a arm.Uint8x8) uint8 {
	return 0
}

// MaxvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_u16'.
// Requires NEON.
func MaxvU16(a arm.Uint16x4) uint16 {
	return 0
}

// MaxvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxv_u32'.
// Requires NEON.
func MaxvU32(a arm.Uint32x2) uint32 {
	return 0
}

// MaxvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_f32'.
// Requires NEON.
func MaxvqF32(a arm.Float32x4) float32 {
	return 0
}

// MaxvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_f64'.
// Requires NEON.
func MaxvqF64(a arm.Float64x2) float64 {
	return 0
}

// MaxvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_s8'.
// Requires NEON.
func MaxvqS8(a arm.Int8x16) int8 {
	return 0
}

// MaxvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_s16'.
// Requires NEON.
func MaxvqS16(a arm.Int16x8) int16 {
	return 0
}

// MaxvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_s32'.
// Requires NEON.
func MaxvqS32(a arm.Int32x4) int32 {
	return 0
}

// MaxvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_u8'.
// Requires NEON.
func MaxvqU8(a arm.Uint8x16) uint8 {
	return 0
}

// MaxvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_u16'.
// Requires NEON.
func MaxvqU16(a arm.Uint16x8) uint16 {
	return 0
}

// MaxvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxvq_u32'.
// Requires NEON.
func MaxvqU32(a arm.Uint32x4) uint32 {
	return 0
}

// MaxnmvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmv_f32'.
// Requires NEON.
func MaxnmvF32(a arm.Float32x2) float32 {
	return 0
}

// MaxnmvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmvq_f32'.
// Requires NEON.
func MaxnmvqF32(a arm.Float32x4) float32 {
	return 0
}

// MaxnmvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmaxnmvq_f64'.
// Requires NEON.
func MaxnmvqF64(a arm.Float64x2) float64 {
	return 0
}

// MinF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_f32'.
// Requires NEON.
func MinF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MinS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_s8'.
// Requires NEON.
func MinS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MinS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_s16'.
// Requires NEON.
func MinS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MinS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_s32'.
// Requires NEON.
func MinS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MinU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_u8'.
// Requires NEON.
func MinU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MinU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_u16'.
// Requires NEON.
func MinU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MinU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmin_u32'.
// Requires NEON.
func MinU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MinqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_f32'.
// Requires NEON.
func MinqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MinqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_f64'.
// Requires NEON.
func MinqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MinqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_s8'.
// Requires NEON.
func MinqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MinqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_s16'.
// Requires NEON.
func MinqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MinqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_s32'.
// Requires NEON.
func MinqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MinqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_u8'.
// Requires NEON.
func MinqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MinqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_u16'.
// Requires NEON.
func MinqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MinqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vminq_u32'.
// Requires NEON.
func MinqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MinnmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnm_f32'.
// Requires NEON.
func MinnmF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MinnmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmq_f32'.
// Requires NEON.
func MinnmqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MinnmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmq_f64'.
// Requires NEON.
func MinnmqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MinvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_f32'.
// Requires NEON.
func MinvF32(a arm.Float32x2) float32 {
	return 0
}

// MinvS8: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_s8'.
// Requires NEON.
func MinvS8(a arm.Int8x8) int8 {
	return 0
}

// MinvS16: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_s16'.
// Requires NEON.
func MinvS16(a arm.Int16x4) int16 {
	return 0
}

// MinvS32: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_s32'.
// Requires NEON.
func MinvS32(a arm.Int32x2) int32 {
	return 0
}

// MinvU8: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_u8'.
// Requires NEON.
func MinvU8(a arm.Uint8x8) uint8 {
	return 0
}

// MinvU16: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_u16'.
// Requires NEON.
func MinvU16(a arm.Uint16x4) uint16 {
	return 0
}

// MinvU32: ARM NEON intrinsic 
//
// Intrinsic: 'vminv_u32'.
// Requires NEON.
func MinvU32(a arm.Uint32x2) uint32 {
	return 0
}

// MinvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_f32'.
// Requires NEON.
func MinvqF32(a arm.Float32x4) float32 {
	return 0
}

// MinvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_f64'.
// Requires NEON.
func MinvqF64(a arm.Float64x2) float64 {
	return 0
}

// MinvqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_s8'.
// Requires NEON.
func MinvqS8(a arm.Int8x16) int8 {
	return 0
}

// MinvqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_s16'.
// Requires NEON.
func MinvqS16(a arm.Int16x8) int16 {
	return 0
}

// MinvqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_s32'.
// Requires NEON.
func MinvqS32(a arm.Int32x4) int32 {
	return 0
}

// MinvqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_u8'.
// Requires NEON.
func MinvqU8(a arm.Uint8x16) uint8 {
	return 0
}

// MinvqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_u16'.
// Requires NEON.
func MinvqU16(a arm.Uint16x8) uint16 {
	return 0
}

// MinvqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vminvq_u32'.
// Requires NEON.
func MinvqU32(a arm.Uint32x4) uint32 {
	return 0
}

// MinnmvF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmv_f32'.
// Requires NEON.
func MinnmvF32(a arm.Float32x2) float32 {
	return 0
}

// MinnmvqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmvq_f32'.
// Requires NEON.
func MinnmvqF32(a arm.Float32x4) float32 {
	return 0
}

// MinnmvqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vminnmvq_f64'.
// Requires NEON.
func MinnmvqF64(a arm.Float64x2) float64 {
	return 0
}

// MlaF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_f32'.
// Requires NEON.
func MlaF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlaF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_f64'.
// Requires NEON.
func MlaF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MlaqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_f32'.
// Requires NEON.
func MlaqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlaqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_f64'.
// Requires NEON.
func MlaqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MlaLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlaLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlaLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlaLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlaLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlaLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlaLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneqS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x8, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlaLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneqS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x4, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlaLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneqU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x8, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlaLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmla_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaLaneqU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x4, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlaqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlaqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x4, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlaqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x2, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlaqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x4, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlaqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x2, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlaqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlaqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlaqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlaqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlaqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlaq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlaqLaneqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_f32'.
// Requires NEON.
func MlsF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlsF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_f64'.
// Requires NEON.
func MlsF64(a arm.Float64x1, b arm.Float64x1, c arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MlsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_f32'.
// Requires NEON.
func MlsqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_f64'.
// Requires NEON.
func MlsqF64(a arm.Float64x2, b arm.Float64x2, c arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MlsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlsLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlsLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlsLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneqF32(a arm.Float32x2, b arm.Float32x2, c arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MlsLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneqS16(a arm.Int16x4, b arm.Int16x4, c arm.Int16x8, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MlsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneqS32(a arm.Int32x2, b arm.Int32x2, c arm.Int32x4, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MlsLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneqU16(a arm.Uint16x4, b arm.Uint16x4, c arm.Uint16x8, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MlsLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmls_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsLaneqU32(a arm.Uint32x2, b arm.Uint32x2, c arm.Uint32x4, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MlsqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlsqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x4, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlsqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x2, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlsqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x4, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlsqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x2, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MlsqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneqF32(a arm.Float32x4, b arm.Float32x4, c arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MlsqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneqS16(a arm.Int16x8, b arm.Int16x8, c arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MlsqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneqS32(a arm.Int32x4, b arm.Int32x4, c arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MlsqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneqU16(a arm.Uint16x8, b arm.Uint16x8, c arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MlsqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmlsq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MlsqLaneqU32(a arm.Uint32x4, b arm.Uint32x4, c arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MovNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_f32'.
// Requires NEON.
func MovNF32(a float32) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MovNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_f64'.
// Requires NEON.
func MovNF64(a float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MovNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_p8'.
// Requires NEON.
func MovNP8(a arm.Poly8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// MovNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_p16'.
// Requires NEON.
func MovNP16(a arm.Poly16) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// MovNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s8'.
// Requires NEON.
func MovNS8(a int8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// MovNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s16'.
// Requires NEON.
func MovNS16(a int16) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MovNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s32'.
// Requires NEON.
func MovNS32(a int32) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MovNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_s64'.
// Requires NEON.
func MovNS64(a int64) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// MovNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u8'.
// Requires NEON.
func MovNU8(a uint8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// MovNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u16'.
// Requires NEON.
func MovNU16(a uint16) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MovNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u32'.
// Requires NEON.
func MovNU32(a uint32) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MovNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmov_n_u64'.
// Requires NEON.
func MovNU64(a uint64) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// MovqNF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_f32'.
// Requires NEON.
func MovqNF32(a float32) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MovqNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_f64'.
// Requires NEON.
func MovqNF64(a float64) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MovqNP8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_p8'.
// Requires NEON.
func MovqNP8(a arm.Poly8) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// MovqNP16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_p16'.
// Requires NEON.
func MovqNP16(a arm.Poly16) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// MovqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s8'.
// Requires NEON.
func MovqNS8(a int8) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// MovqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s16'.
// Requires NEON.
func MovqNS16(a int16) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MovqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s32'.
// Requires NEON.
func MovqNS32(a int32) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MovqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_s64'.
// Requires NEON.
func MovqNS64(a int64) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// MovqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u8'.
// Requires NEON.
func MovqNU8(a uint8) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// MovqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u16'.
// Requires NEON.
func MovqNU16(a uint16) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MovqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u32'.
// Requires NEON.
func MovqNU32(a uint32) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MovqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vmovq_n_u64'.
// Requires NEON.
func MovqNU64(a uint64) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// MulLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneF32(a arm.Float32x2, b arm.Float32x2, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MulLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneF64(a arm.Float64x1, b arm.Float64x1, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MulLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneS16(a arm.Int16x4, b arm.Int16x4, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MulLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneS32(a arm.Int32x2, b arm.Int32x2, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MulLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneU16(a arm.Uint16x4, b arm.Uint16x4, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MulLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneU32(a arm.Uint32x2, b arm.Uint32x2, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MuldLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmuld_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MuldLaneF64(a float64, b arm.Float64x1, lane int) float64 {
	return 0
}

// MuldLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmuld_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MuldLaneqF64(a float64, b arm.Float64x2, lane int) float64 {
	return 0
}

// MulsLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmuls_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulsLaneF32(a float32, b arm.Float32x2, lane int) float32 {
	return 0
}

// MulsLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmuls_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulsLaneqF32(a float32, b arm.Float32x4, lane int) float32 {
	return 0
}

// MulLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneqF32(a arm.Float32x2, b arm.Float32x4, lane int) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// MulLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneqF64(a arm.Float64x1, b arm.Float64x2, lane int) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MulLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneqS16(a arm.Int16x4, b arm.Int16x8, lane int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// MulLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneqS32(a arm.Int32x2, b arm.Int32x4, lane int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// MulLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneqU16(a arm.Uint16x4, b arm.Uint16x8, lane int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// MulLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulLaneqU32(a arm.Uint32x2, b arm.Uint32x4, lane int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// MulNF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmul_n_f64'.
// Requires NEON.
func MulNF64(a arm.Float64x1, b float64) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// MulqLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneF32(a arm.Float32x4, b arm.Float32x2, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MulqLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneF64(a arm.Float64x2, b arm.Float64x1, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MulqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneS16(a arm.Int16x8, b arm.Int16x4, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MulqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneS32(a arm.Int32x4, b arm.Int32x2, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MulqLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneU16(a arm.Uint16x8, b arm.Uint16x4, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MulqLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_lane_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneU32(a arm.Uint32x4, b arm.Uint32x2, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MulqLaneqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_f32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneqF32(a arm.Float32x4, b arm.Float32x4, lane int) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// MulqLaneqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_f64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneqF64(a arm.Float64x2, b arm.Float64x2, lane int) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// MulqLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneqS16(a arm.Int16x8, b arm.Int16x8, lane int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// MulqLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneqS32(a arm.Int32x4, b arm.Int32x4, lane int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// MulqLaneqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneqU16(a arm.Uint16x8, b arm.Uint16x8, lane int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// MulqLaneqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vmulq_laneq_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func MulqLaneqU32(a arm.Uint32x4, b arm.Uint32x4, lane int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// NegF32: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_f32'.
// Requires NEON.
func NegF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// NegF64: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_f64'.
// Requires NEON.
func NegF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// NegS8: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s8'.
// Requires NEON.
func NegS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// NegS16: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s16'.
// Requires NEON.
func NegS16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// NegS32: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s32'.
// Requires NEON.
func NegS32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// NegS64: ARM NEON intrinsic 
//
// Intrinsic: 'vneg_s64'.
// Requires NEON.
func NegS64(a arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// NegqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_f32'.
// Requires NEON.
func NegqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// NegqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_f64'.
// Requires NEON.
func NegqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// NegqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s8'.
// Requires NEON.
func NegqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// NegqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s16'.
// Requires NEON.
func NegqS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// NegqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s32'.
// Requires NEON.
func NegqS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// NegqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vnegq_s64'.
// Requires NEON.
func NegqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// PaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_s8'.
// Requires NEON.
func PaddS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// PaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_s16'.
// Requires NEON.
func PaddS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// PaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_s32'.
// Requires NEON.
func PaddS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// PaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_u8'.
// Requires NEON.
func PaddU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// PaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_u16'.
// Requires NEON.
func PaddU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// PaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vpadd_u32'.
// Requires NEON.
func PaddU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// PadddF64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddd_f64'.
// Requires NEON.
func PadddF64(a arm.Float64x2) float64 {
	return 0
}

// PadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddd_s64'.
// Requires NEON.
func PadddS64(a arm.Int64x2) int64 {
	return 0
}

// PadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vpaddd_u64'.
// Requires NEON.
func PadddU64(a arm.Uint64x2) uint64 {
	return 0
}

// QabsqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsq_s64'.
// Requires NEON.
func QabsqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QabsbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsb_s8'.
// Requires NEON.
func QabsbS8(a int8) int8 {
	return 0
}

// QabshS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsh_s16'.
// Requires NEON.
func QabshS16(a int16) int16 {
	return 0
}

// QabssS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqabss_s32'.
// Requires NEON.
func QabssS32(a int32) int32 {
	return 0
}

// QabsdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqabsd_s64'.
// Requires NEON.
func QabsdS64(a int64) int64 {
	return 0
}

// QaddbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddb_s8'.
// Requires NEON.
func QaddbS8(a int8, b int8) int8 {
	return 0
}

// QaddhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddh_s16'.
// Requires NEON.
func QaddhS16(a int16, b int16) int16 {
	return 0
}

// QaddsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadds_s32'.
// Requires NEON.
func QaddsS32(a int32, b int32) int32 {
	return 0
}

// QadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddd_s64'.
// Requires NEON.
func QadddS64(a int64, b int64) int64 {
	return 0
}

// QaddbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddb_u8'.
// Requires NEON.
func QaddbU8(a uint8, b uint8) uint8 {
	return 0
}

// QaddhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddh_u16'.
// Requires NEON.
func QaddhU16(a uint16, b uint16) uint16 {
	return 0
}

// QaddsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqadds_u32'.
// Requires NEON.
func QaddsU32(a uint32, b uint32) uint32 {
	return 0
}

// QadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqaddd_u64'.
// Requires NEON.
func QadddU64(a uint64, b uint64) uint64 {
	return 0
}

// QdmlalS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_s16'.
// Requires NEON.
func QdmlalS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_s16'.
// Requires NEON.
func QdmlalHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalHighLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalHighLaneS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalHighLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalHighLaneqS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_n_s16'.
// Requires NEON.
func QdmlalHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalLaneS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalLaneqS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_n_s16'.
// Requires NEON.
func QdmlalNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlalS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_s32'.
// Requires NEON.
func QdmlalS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_s32'.
// Requires NEON.
func QdmlalHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalHighLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalHighLaneS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalHighLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalHighLaneqS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_high_n_s32'.
// Requires NEON.
func QdmlalHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalLaneS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalLaneqS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlal_n_s32'.
// Requires NEON.
func QdmlalNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlalhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlalh_s16'.
// Requires NEON.
func QdmlalhS16(a int32, b int16, c int16) int32 {
	return 0
}

// QdmlalhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlalh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalhLaneS16(a int32, b int16, c arm.Int16x4, d int) int32 {
	return 0
}

// QdmlalhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlalh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalhLaneqS16(a int32, b int16, c arm.Int16x8, d int) int32 {
	return 0
}

// QdmlalsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlals_s32'.
// Requires NEON.
func QdmlalsS32(a int64, b int32, c int32) int64 {
	return 0
}

// QdmlalsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlals_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalsLaneS32(a int64, b int32, c arm.Int32x2, d int) int64 {
	return 0
}

// QdmlalsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlals_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlalsLaneqS32(a int64, b int32, c arm.Int32x4, d int) int64 {
	return 0
}

// QdmlslS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_s16'.
// Requires NEON.
func QdmlslS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_s16'.
// Requires NEON.
func QdmlslHighS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslHighLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslHighLaneS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslHighLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslHighLaneqS16(a arm.Int32x4, b arm.Int16x8, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_n_s16'.
// Requires NEON.
func QdmlslHighNS16(a arm.Int32x4, b arm.Int16x8, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslLaneS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x4, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslLaneqS16(a arm.Int32x4, b arm.Int16x4, c arm.Int16x8, d int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_n_s16'.
// Requires NEON.
func QdmlslNS16(a arm.Int32x4, b arm.Int16x4, c int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmlslS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_s32'.
// Requires NEON.
func QdmlslS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_s32'.
// Requires NEON.
func QdmlslHighS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslHighLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslHighLaneS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslHighLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslHighLaneqS32(a arm.Int64x2, b arm.Int32x4, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_high_n_s32'.
// Requires NEON.
func QdmlslHighNS32(a arm.Int64x2, b arm.Int32x4, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslLaneS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x2, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslLaneqS32(a arm.Int64x2, b arm.Int32x2, c arm.Int32x4, d int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsl_n_s32'.
// Requires NEON.
func QdmlslNS32(a arm.Int64x2, b arm.Int32x2, c int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmlslhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlslh_s16'.
// Requires NEON.
func QdmlslhS16(a int32, b int16, c int16) int32 {
	return 0
}

// QdmlslhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlslh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslhLaneS16(a int32, b int16, c arm.Int16x4, d int) int32 {
	return 0
}

// QdmlslhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlslh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslhLaneqS16(a int32, b int16, c arm.Int16x8, d int) int32 {
	return 0
}

// QdmlslsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsls_s32'.
// Requires NEON.
func QdmlslsS32(a int64, b int32, c int32) int64 {
	return 0
}

// QdmlslsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsls_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslsLaneS32(a int64, b int32, c arm.Int32x2, d int) int64 {
	return 0
}

// QdmlslsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmlsls_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmlslsLaneqS32(a int64, b int32, c arm.Int32x4, d int) int64 {
	return 0
}

// QdmulhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhLaneS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QdmulhLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulh_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhLaneS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QdmulhqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhqLaneS16(a arm.Int16x8, b arm.Int16x4, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QdmulhqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhqLaneS32(a arm.Int32x4, b arm.Int32x2, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmulhhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhh_s16'.
// Requires NEON.
func QdmulhhS16(a int16, b int16) int16 {
	return 0
}

// QdmulhhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhhLaneS16(a int16, b arm.Int16x4, c int) int16 {
	return 0
}

// QdmulhhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhhLaneqS16(a int16, b arm.Int16x8, c int) int16 {
	return 0
}

// QdmulhsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhs_s32'.
// Requires NEON.
func QdmulhsS32(a int32, b int32) int32 {
	return 0
}

// QdmulhsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhs_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhsLaneS32(a int32, b arm.Int32x2, c int) int32 {
	return 0
}

// QdmulhsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulhs_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmulhsLaneqS32(a int32, b arm.Int32x4, c int) int32 {
	return 0
}

// QdmullS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_s16'.
// Requires NEON.
func QdmullS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullHighS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_s16'.
// Requires NEON.
func QdmullHighS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullHighLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullHighLaneS16(a arm.Int16x8, b arm.Int16x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullHighLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullHighLaneqS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_n_s16'.
// Requires NEON.
func QdmullHighNS16(a arm.Int16x8, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullLaneS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullLaneqS16(a arm.Int16x4, b arm.Int16x8, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_n_s16'.
// Requires NEON.
func QdmullNS16(a arm.Int16x4, b int16) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QdmullS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_s32'.
// Requires NEON.
func QdmullS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullHighS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_s32'.
// Requires NEON.
func QdmullHighS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullHighLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullHighLaneS32(a arm.Int32x4, b arm.Int32x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullHighLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullHighLaneqS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_high_n_s32'.
// Requires NEON.
func QdmullHighNS32(a arm.Int32x4, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullLaneS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullLaneqS32(a arm.Int32x2, b arm.Int32x4, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmull_n_s32'.
// Requires NEON.
func QdmullNS32(a arm.Int32x2, b int32) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QdmullhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmullh_s16'.
// Requires NEON.
func QdmullhS16(a int16, b int16) int32 {
	return 0
}

// QdmullhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmullh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullhLaneS16(a int16, b arm.Int16x4, c int) int32 {
	return 0
}

// QdmullhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmullh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullhLaneqS16(a int16, b arm.Int16x8, c int) int32 {
	return 0
}

// QdmullsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulls_s32'.
// Requires NEON.
func QdmullsS32(a int32, b int32) int64 {
	return 0
}

// QdmullsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulls_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullsLaneS32(a int32, b arm.Int32x2, c int) int64 {
	return 0
}

// QdmullsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqdmulls_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QdmullsLaneqS32(a int32, b arm.Int32x4, c int) int64 {
	return 0
}

// QmovnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_s16'.
// Requires NEON.
func QmovnS16(a arm.Int16x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QmovnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_s32'.
// Requires NEON.
func QmovnS32(a arm.Int32x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QmovnS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_s64'.
// Requires NEON.
func QmovnS64(a arm.Int64x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QmovnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_u16'.
// Requires NEON.
func QmovnU16(a arm.Uint16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QmovnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_u32'.
// Requires NEON.
func QmovnU32(a arm.Uint32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QmovnU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovn_u64'.
// Requires NEON.
func QmovnU64(a arm.Uint64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QmovnhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnh_s16'.
// Requires NEON.
func QmovnhS16(a int16) int8 {
	return 0
}

// QmovnsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovns_s32'.
// Requires NEON.
func QmovnsS32(a int32) int16 {
	return 0
}

// QmovndS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnd_s64'.
// Requires NEON.
func QmovndS64(a int64) int32 {
	return 0
}

// QmovnhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnh_u16'.
// Requires NEON.
func QmovnhU16(a uint16) uint8 {
	return 0
}

// QmovnsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovns_u32'.
// Requires NEON.
func QmovnsU32(a uint32) uint16 {
	return 0
}

// QmovndU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovnd_u64'.
// Requires NEON.
func QmovndU64(a uint64) uint32 {
	return 0
}

// QmovunS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_s16'.
// Requires NEON.
func QmovunS16(a arm.Int16x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QmovunS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_s32'.
// Requires NEON.
func QmovunS32(a arm.Int32x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QmovunS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovun_s64'.
// Requires NEON.
func QmovunS64(a arm.Int64x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QmovunhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovunh_s16'.
// Requires NEON.
func QmovunhS16(a int16) int8 {
	return 0
}

// QmovunsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovuns_s32'.
// Requires NEON.
func QmovunsS32(a int32) int16 {
	return 0
}

// QmovundS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqmovund_s64'.
// Requires NEON.
func QmovundS64(a int64) int32 {
	return 0
}

// QnegqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegq_s64'.
// Requires NEON.
func QnegqS64(a arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QnegbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegb_s8'.
// Requires NEON.
func QnegbS8(a int8) int8 {
	return 0
}

// QneghS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegh_s16'.
// Requires NEON.
func QneghS16(a int16) int16 {
	return 0
}

// QnegsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegs_s32'.
// Requires NEON.
func QnegsS32(a int32) int32 {
	return 0
}

// QnegdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqnegd_s64'.
// Requires NEON.
func QnegdS64(a int64) int64 {
	return 0
}

// QrdmulhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhLaneS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QrdmulhLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulh_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhLaneS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QrdmulhqLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhqLaneS16(a arm.Int16x8, b arm.Int16x4, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QrdmulhqLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhq_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhqLaneS32(a arm.Int32x4, b arm.Int32x2, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QrdmulhhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhh_s16'.
// Requires NEON.
func QrdmulhhS16(a int16, b int16) int16 {
	return 0
}

// QrdmulhhLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhh_lane_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhhLaneS16(a int16, b arm.Int16x4, c int) int16 {
	return 0
}

// QrdmulhhLaneqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhh_laneq_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhhLaneqS16(a int16, b arm.Int16x8, c int) int16 {
	return 0
}

// QrdmulhsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhs_s32'.
// Requires NEON.
func QrdmulhsS32(a int32, b int32) int32 {
	return 0
}

// QrdmulhsLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhs_lane_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhsLaneS32(a int32, b arm.Int32x2, c int) int32 {
	return 0
}

// QrdmulhsLaneqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrdmulhs_laneq_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrdmulhsLaneqS32(a int32, b arm.Int32x4, c int) int32 {
	return 0
}

// QrshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s8'.
// Requires NEON.
func QrshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QrshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s16'.
// Requires NEON.
func QrshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QrshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s32'.
// Requires NEON.
func QrshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QrshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_s64'.
// Requires NEON.
func QrshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QrshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u8'.
// Requires NEON.
func QrshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QrshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u16'.
// Requires NEON.
func QrshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QrshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u32'.
// Requires NEON.
func QrshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QrshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshl_u64'.
// Requires NEON.
func QrshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// QrshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s8'.
// Requires NEON.
func QrshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QrshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s16'.
// Requires NEON.
func QrshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QrshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s32'.
// Requires NEON.
func QrshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QrshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_s64'.
// Requires NEON.
func QrshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QrshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u8'.
// Requires NEON.
func QrshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QrshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u16'.
// Requires NEON.
func QrshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QrshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u32'.
// Requires NEON.
func QrshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QrshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlq_u64'.
// Requires NEON.
func QrshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QrshlbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlb_s8'.
// Requires NEON.
func QrshlbS8(a int8, b int8) int8 {
	return 0
}

// QrshlhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlh_s16'.
// Requires NEON.
func QrshlhS16(a int16, b int16) int16 {
	return 0
}

// QrshlsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshls_s32'.
// Requires NEON.
func QrshlsS32(a int32, b int32) int32 {
	return 0
}

// QrshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshld_s64'.
// Requires NEON.
func QrshldS64(a int64, b int64) int64 {
	return 0
}

// QrshlbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlb_u8'.
// Requires NEON.
func QrshlbU8(a uint8, b uint8) uint8 {
	return 0
}

// QrshlhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshlh_u16'.
// Requires NEON.
func QrshlhU16(a uint16, b uint16) uint16 {
	return 0
}

// QrshlsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshls_u32'.
// Requires NEON.
func QrshlsU32(a uint32, b uint32) uint32 {
	return 0
}

// QrshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshld_u64'.
// Requires NEON.
func QrshldU64(a uint64, b uint64) uint64 {
	return 0
}

// QrshrnNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnNS16(a arm.Int16x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QrshrnNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnNS32(a arm.Int32x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QrshrnNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnNS64(a arm.Int64x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QrshrnNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnNU16(a arm.Uint16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QrshrnNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnNU32(a arm.Uint32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QrshrnNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrn_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnNU64(a arm.Uint64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QrshrnhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnhNS16(a int16, b int) int8 {
	return 0
}

// QrshrnsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnsNS32(a int32, b int) int16 {
	return 0
}

// QrshrndNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrndNS64(a int64, b int) int32 {
	return 0
}

// QrshrnhNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnh_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnhNU16(a uint16, b int) uint8 {
	return 0
}

// QrshrnsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrns_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrnsNU32(a uint32, b int) uint16 {
	return 0
}

// QrshrndNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrnd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrndNU64(a uint64, b int) uint32 {
	return 0
}

// QrshrunNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrun_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrunNS16(a arm.Int16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QrshrunNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrun_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrunNS32(a arm.Int32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QrshrunNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrun_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrunNS64(a arm.Int64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QrshrunhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrunh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrunhNS16(a int16, b int) int8 {
	return 0
}

// QrshrunsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshruns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrunsNS32(a int32, b int) int16 {
	return 0
}

// QrshrundNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqrshrund_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QrshrundNS64(a int64, b int) int32 {
	return 0
}

// QshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s8'.
// Requires NEON.
func QshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s16'.
// Requires NEON.
func QshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s32'.
// Requires NEON.
func QshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_s64'.
// Requires NEON.
func QshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u8'.
// Requires NEON.
func QshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u16'.
// Requires NEON.
func QshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u32'.
// Requires NEON.
func QshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_u64'.
// Requires NEON.
func QshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// QshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s8'.
// Requires NEON.
func QshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s16'.
// Requires NEON.
func QshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s32'.
// Requires NEON.
func QshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_s64'.
// Requires NEON.
func QshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u8'.
// Requires NEON.
func QshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u16'.
// Requires NEON.
func QshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u32'.
// Requires NEON.
func QshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_u64'.
// Requires NEON.
func QshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QshlbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_s8'.
// Requires NEON.
func QshlbS8(a int8, b int8) int8 {
	return 0
}

// QshlhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_s16'.
// Requires NEON.
func QshlhS16(a int16, b int16) int16 {
	return 0
}

// QshlsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_s32'.
// Requires NEON.
func QshlsS32(a int32, b int32) int32 {
	return 0
}

// QshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_s64'.
// Requires NEON.
func QshldS64(a int64, b int64) int64 {
	return 0
}

// QshlbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_u8'.
// Requires NEON.
func QshlbU8(a uint8, b uint8) uint8 {
	return 0
}

// QshlhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_u16'.
// Requires NEON.
func QshlhU16(a uint16, b uint16) uint16 {
	return 0
}

// QshlsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_u32'.
// Requires NEON.
func QshlsU32(a uint32, b uint32) uint32 {
	return 0
}

// QshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_u64'.
// Requires NEON.
func QshldU64(a uint64, b uint64) uint64 {
	return 0
}

// QshlNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QshlNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QshlNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QshlNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// QshlNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QshlNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QshlNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QshlNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshl_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// QshlqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// QshlqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// QshlqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// QshlqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// QshlqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QshlqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QshlqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QshlqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QshlbNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlbNS8(a int8, b int) int8 {
	return 0
}

// QshlhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlhNS16(a int16, b int) int16 {
	return 0
}

// QshlsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlsNS32(a int32, b int) int32 {
	return 0
}

// QshldNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshldNS64(a int64, b int) int64 {
	return 0
}

// QshlbNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlb_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlbNU8(a uint8, b int) uint8 {
	return 0
}

// QshlhNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlh_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlhNU16(a uint16, b int) uint16 {
	return 0
}

// QshlsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshls_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlsNU32(a uint32, b int) uint32 {
	return 0
}

// QshldNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshld_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshldNU64(a uint64, b int) uint64 {
	return 0
}

// QshluNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluNS8(a arm.Int8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QshluNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluNS16(a arm.Int16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QshluNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluNS32(a arm.Int32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QshluNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlu_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluNS64(a arm.Int64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// QshluqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluqNS8(a arm.Int8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// QshluqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluqNS16(a arm.Int16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// QshluqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluqNS32(a arm.Int32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// QshluqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluqNS64(a arm.Int64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// QshlubNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlub_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlubNS8(a int8, b int) int8 {
	return 0
}

// QshluhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshluh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshluhNS16(a int16, b int) int16 {
	return 0
}

// QshlusNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlus_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshlusNS32(a int32, b int) int32 {
	return 0
}

// QshludNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshlud_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshludNS64(a int64, b int) uint64 {
	return 0
}

// QshrnNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnNS16(a arm.Int16x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// QshrnNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnNS32(a arm.Int32x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// QshrnNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnNS64(a arm.Int64x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// QshrnNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnNU16(a arm.Uint16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QshrnNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnNU32(a arm.Uint32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QshrnNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrn_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnNU64(a arm.Uint64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QshrnhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnhNS16(a int16, b int) int8 {
	return 0
}

// QshrnsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnsNS32(a int32, b int) int16 {
	return 0
}

// QshrndNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrndNS64(a int64, b int) int32 {
	return 0
}

// QshrnhNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnh_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnhNU16(a uint16, b int) uint8 {
	return 0
}

// QshrnsNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrns_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrnsNU32(a uint32, b int) uint16 {
	return 0
}

// QshrndNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrnd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrndNU64(a uint64, b int) uint32 {
	return 0
}

// QshrunNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrun_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrunNS16(a arm.Int16x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// QshrunNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrun_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrunNS32(a arm.Int32x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// QshrunNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrun_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrunNS64(a arm.Int64x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// QshrunhNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrunh_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrunhNS16(a int16, b int) int8 {
	return 0
}

// QshrunsNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqshruns_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrunsNS32(a int32, b int) int16 {
	return 0
}

// QshrundNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqshrund_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func QshrundNS64(a int64, b int) int32 {
	return 0
}

// QsubbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubb_s8'.
// Requires NEON.
func QsubbS8(a int8, b int8) int8 {
	return 0
}

// QsubhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubh_s16'.
// Requires NEON.
func QsubhS16(a int16, b int16) int16 {
	return 0
}

// QsubsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubs_s32'.
// Requires NEON.
func QsubsS32(a int32, b int32) int32 {
	return 0
}

// QsubdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubd_s64'.
// Requires NEON.
func QsubdS64(a int64, b int64) int64 {
	return 0
}

// QsubbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubb_u8'.
// Requires NEON.
func QsubbU8(a uint8, b uint8) uint8 {
	return 0
}

// QsubhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubh_u16'.
// Requires NEON.
func QsubhU16(a uint16, b uint16) uint16 {
	return 0
}

// QsubsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubs_u32'.
// Requires NEON.
func QsubsU32(a uint32, b uint32) uint32 {
	return 0
}

// QsubdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vqsubd_u64'.
// Requires NEON.
func QsubdU64(a uint64, b uint64) uint64 {
	return 0
}

// RbitP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbit_p8'.
// Requires NEON.
func RbitP8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// RbitS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbit_s8'.
// Requires NEON.
func RbitS8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RbitU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbit_u8'.
// Requires NEON.
func RbitU8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RbitqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbitq_p8'.
// Requires NEON.
func RbitqP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// RbitqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbitq_s8'.
// Requires NEON.
func RbitqS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RbitqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrbitq_u8'.
// Requires NEON.
func RbitqU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RecpeU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpe_u32'.
// Requires NEON.
func RecpeU32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RecpeqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpeq_u32'.
// Requires NEON.
func RecpeqU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RecpesF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpes_f32'.
// Requires NEON.
func RecpesF32(a float32) float32 {
	return 0
}

// RecpedF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecped_f64'.
// Requires NEON.
func RecpedF64(a float64) float64 {
	return 0
}

// RecpeF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpe_f32'.
// Requires NEON.
func RecpeF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RecpeqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpeq_f32'.
// Requires NEON.
func RecpeqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RecpeqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpeq_f64'.
// Requires NEON.
func RecpeqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RecpssF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpss_f32'.
// Requires NEON.
func RecpssF32(a float32, b float32) float32 {
	return 0
}

// RecpsdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpsd_f64'.
// Requires NEON.
func RecpsdF64(a float64, b float64) float64 {
	return 0
}

// RecpsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecps_f32'.
// Requires NEON.
func RecpsF32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RecpsqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpsq_f32'.
// Requires NEON.
func RecpsqF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RecpsqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpsq_f64'.
// Requires NEON.
func RecpsqF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RecpxsF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpxs_f32'.
// Requires NEON.
func RecpxsF32(a float32) float32 {
	return 0
}

// RecpxdF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrecpxd_f64'.
// Requires NEON.
func RecpxdF64(a float64) float64 {
	return 0
}

// Rev16P8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16_p8'.
// Requires NEON.
func Rev16P8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Rev16S8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16_s8'.
// Requires NEON.
func Rev16S8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Rev16U8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16_u8'.
// Requires NEON.
func Rev16U8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Rev16qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16q_p8'.
// Requires NEON.
func Rev16qP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Rev16qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16q_s8'.
// Requires NEON.
func Rev16qS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Rev16qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev16q_u8'.
// Requires NEON.
func Rev16qU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Rev32P8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_p8'.
// Requires NEON.
func Rev32P8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Rev32P16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_p16'.
// Requires NEON.
func Rev32P16(a arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Rev32S8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_s8'.
// Requires NEON.
func Rev32S8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Rev32S16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_s16'.
// Requires NEON.
func Rev32S16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Rev32U8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_u8'.
// Requires NEON.
func Rev32U8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Rev32U16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32_u16'.
// Requires NEON.
func Rev32U16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Rev32qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_p8'.
// Requires NEON.
func Rev32qP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Rev32qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_p16'.
// Requires NEON.
func Rev32qP16(a arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Rev32qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_s8'.
// Requires NEON.
func Rev32qS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Rev32qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_s16'.
// Requires NEON.
func Rev32qS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Rev32qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_u8'.
// Requires NEON.
func Rev32qU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Rev32qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev32q_u16'.
// Requires NEON.
func Rev32qU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Rev64F32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_f32'.
// Requires NEON.
func Rev64F32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Rev64P8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_p8'.
// Requires NEON.
func Rev64P8(a arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Rev64P16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_p16'.
// Requires NEON.
func Rev64P16(a arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Rev64S8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_s8'.
// Requires NEON.
func Rev64S8(a arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Rev64S16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_s16'.
// Requires NEON.
func Rev64S16(a arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Rev64S32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_s32'.
// Requires NEON.
func Rev64S32(a arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Rev64U8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_u8'.
// Requires NEON.
func Rev64U8(a arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Rev64U16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_u16'.
// Requires NEON.
func Rev64U16(a arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Rev64U32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64_u32'.
// Requires NEON.
func Rev64U32(a arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Rev64qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_f32'.
// Requires NEON.
func Rev64qF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Rev64qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_p8'.
// Requires NEON.
func Rev64qP8(a arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Rev64qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_p16'.
// Requires NEON.
func Rev64qP16(a arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Rev64qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_s8'.
// Requires NEON.
func Rev64qS8(a arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Rev64qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_s16'.
// Requires NEON.
func Rev64qS16(a arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Rev64qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_s32'.
// Requires NEON.
func Rev64qS32(a arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Rev64qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_u8'.
// Requires NEON.
func Rev64qU8(a arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Rev64qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_u16'.
// Requires NEON.
func Rev64qU16(a arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Rev64qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrev64q_u32'.
// Requires NEON.
func Rev64qU32(a arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RndF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrnd_f32'.
// Requires NEON.
func RndF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrnd_f64'.
// Requires NEON.
func RndF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndq_f32'.
// Requires NEON.
func RndqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndq_f64'.
// Requires NEON.
func RndqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RndaF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrnda_f32'.
// Requires NEON.
func RndaF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndaF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrnda_f64'.
// Requires NEON.
func RndaF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndaqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndaq_f32'.
// Requires NEON.
func RndaqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndaqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndaq_f64'.
// Requires NEON.
func RndaqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RndiF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndi_f32'.
// Requires NEON.
func RndiF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndiF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndi_f64'.
// Requires NEON.
func RndiF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndiqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndiq_f32'.
// Requires NEON.
func RndiqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndiqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndiq_f64'.
// Requires NEON.
func RndiqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RndmF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndm_f32'.
// Requires NEON.
func RndmF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndmF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndm_f64'.
// Requires NEON.
func RndmF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndmqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndmq_f32'.
// Requires NEON.
func RndmqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndmqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndmq_f64'.
// Requires NEON.
func RndmqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RndnF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndn_f32'.
// Requires NEON.
func RndnF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndnF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndn_f64'.
// Requires NEON.
func RndnF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndnqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndnq_f32'.
// Requires NEON.
func RndnqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndnqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndnq_f64'.
// Requires NEON.
func RndnqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RndpF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndp_f32'.
// Requires NEON.
func RndpF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndpF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndp_f64'.
// Requires NEON.
func RndpF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndpqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndpq_f32'.
// Requires NEON.
func RndpqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndpqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndpq_f64'.
// Requires NEON.
func RndpqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RndxF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndx_f32'.
// Requires NEON.
func RndxF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// RndxF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndx_f64'.
// Requires NEON.
func RndxF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// RndxqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vrndxq_f32'.
// Requires NEON.
func RndxqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// RndxqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vrndxq_f64'.
// Requires NEON.
func RndxqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// RshlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s8'.
// Requires NEON.
func RshlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RshlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s16'.
// Requires NEON.
func RshlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// RshlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s32'.
// Requires NEON.
func RshlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// RshlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_s64'.
// Requires NEON.
func RshlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// RshlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u8'.
// Requires NEON.
func RshlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RshlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u16'.
// Requires NEON.
func RshlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// RshlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u32'.
// Requires NEON.
func RshlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RshlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshl_u64'.
// Requires NEON.
func RshlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// RshlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s8'.
// Requires NEON.
func RshlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RshlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s16'.
// Requires NEON.
func RshlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// RshlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s32'.
// Requires NEON.
func RshlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RshlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_s64'.
// Requires NEON.
func RshlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// RshlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u8'.
// Requires NEON.
func RshlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RshlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u16'.
// Requires NEON.
func RshlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// RshlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u32'.
// Requires NEON.
func RshlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RshlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshlq_u64'.
// Requires NEON.
func RshlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// RshldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshld_s64'.
// Requires NEON.
func RshldS64(a int64, b int64) int64 {
	return 0
}

// RshldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshld_u64'.
// Requires NEON.
func RshldU64(a uint64, b int64) uint64 {
	return 0
}

// RshrNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RshrNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// RshrNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// RshrNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// RshrNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RshrNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// RshrNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RshrNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshr_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// RshrqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RshrqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// RshrqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RshrqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// RshrqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RshrqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// RshrqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RshrqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// RshrdNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrdNS64(a int64, b int) int64 {
	return 0
}

// RshrdNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrshrd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RshrdNU64(a uint64, b int) uint64 {
	return 0
}

// RsraNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// RsraNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// RsraNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// RsraNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// RsraNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// RsraNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// RsraNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// RsraNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsra_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// RsraqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// RsraqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// RsraqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// RsraqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// RsraqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// RsraqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// RsraqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// RsraqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsraq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsraqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// RsradNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsrad_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsradNS64(a int64, b int64, c int) int64 {
	return 0
}

// RsradNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vrsrad_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func RsradNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// Sha1cqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1cq_u32'.
// Requires NEON.
func Sha1cqU32(hash_abcd arm.Uint32x4, hash_e uint32, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha1mqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1mq_u32'.
// Requires NEON.
func Sha1mqU32(hash_abcd arm.Uint32x4, hash_e uint32, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha1pqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1pq_u32'.
// Requires NEON.
func Sha1pqU32(hash_abcd arm.Uint32x4, hash_e uint32, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha1hU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1h_u32'.
// Requires NEON.
func Sha1hU32(hash_e uint32) uint32 {
	return 0
}

// Sha1su0qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1su0q_u32'.
// Requires NEON.
func Sha1su0qU32(w0_3 arm.Uint32x4, w4_7 arm.Uint32x4, w8_11 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha1su1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha1su1q_u32'.
// Requires NEON.
func Sha1su1qU32(tw0_3 arm.Uint32x4, w12_15 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha256hqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256hq_u32'.
// Requires NEON.
func Sha256hqU32(hash_abcd arm.Uint32x4, hash_efgh arm.Uint32x4, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha256h2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256h2q_u32'.
// Requires NEON.
func Sha256h2qU32(hash_efgh arm.Uint32x4, hash_abcd arm.Uint32x4, wk arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha256su0qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256su0q_u32'.
// Requires NEON.
func Sha256su0qU32(w0_3 arm.Uint32x4, w4_7 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Sha256su1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsha256su1q_u32'.
// Requires NEON.
func Sha256su1qU32(tw0_3 arm.Uint32x4, w8_11 arm.Uint32x4, w12_15 arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// MullP64: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_p64'.
// Requires NEON.
func MullP64(a arm.Poly64, b arm.Poly64) (dst arm.Poly128) {
	return arm.Poly128{}
}

// MullHighP64: ARM NEON intrinsic 
//
// Intrinsic: 'vmull_high_p64'.
// Requires NEON.
func MullHighP64(a arm.Poly64x2, b arm.Poly64x2) (dst arm.Poly128) {
	return arm.Poly128{}
}

// ShlNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ShlNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ShlNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ShlNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ShlNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ShlNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ShlNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ShlNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ShlqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ShlqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ShlqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ShlqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ShlqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ShlqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ShlqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ShlqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShlqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ShldNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShldNS64(a int64, b int) int64 {
	return 0
}

// ShldNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShldNU64(a uint64, b int) uint64 {
	return 0
}

// ShlS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s8'.
// Requires NEON.
func ShlS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ShlS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s16'.
// Requires NEON.
func ShlS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ShlS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s32'.
// Requires NEON.
func ShlS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ShlS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_s64'.
// Requires NEON.
func ShlS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ShlU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u8'.
// Requires NEON.
func ShlU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ShlU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u16'.
// Requires NEON.
func ShlU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ShlU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u32'.
// Requires NEON.
func ShlU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ShlU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshl_u64'.
// Requires NEON.
func ShlU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ShlqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s8'.
// Requires NEON.
func ShlqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ShlqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s16'.
// Requires NEON.
func ShlqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ShlqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s32'.
// Requires NEON.
func ShlqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ShlqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_s64'.
// Requires NEON.
func ShlqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ShlqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u8'.
// Requires NEON.
func ShlqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ShlqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u16'.
// Requires NEON.
func ShlqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ShlqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u32'.
// Requires NEON.
func ShlqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ShlqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshlq_u64'.
// Requires NEON.
func ShlqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ShldS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_s64'.
// Requires NEON.
func ShldS64(a int64, b int64) int64 {
	return 0
}

// ShldU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshld_u64'.
// Requires NEON.
func ShldU64(a uint64, b uint64) uint64 {
	return 0
}

// ShllHighNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllHighNS8(a arm.Int8x16, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ShllHighNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllHighNS16(a arm.Int16x8, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ShllHighNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllHighNS32(a arm.Int32x4, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ShllHighNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllHighNU8(a arm.Uint8x16, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ShllHighNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllHighNU16(a arm.Uint16x8, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ShllHighNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_high_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllHighNU32(a arm.Uint32x4, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ShllNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllNS8(a arm.Int8x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ShllNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllNS16(a arm.Int16x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ShllNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllNS32(a arm.Int32x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ShllNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllNU8(a arm.Uint8x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ShllNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllNU16(a arm.Uint16x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ShllNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshll_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShllNU32(a arm.Uint32x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ShrNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNS8(a arm.Int8x8, b int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// ShrNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNS16(a arm.Int16x4, b int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// ShrNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNS32(a arm.Int32x2, b int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// ShrNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNS64(a arm.Int64x1, b int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// ShrNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNU8(a arm.Uint8x8, b int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// ShrNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNU16(a arm.Uint16x4, b int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// ShrNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNU32(a arm.Uint32x2, b int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// ShrNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshr_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrNU64(a arm.Uint64x1, b int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// ShrqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNS8(a arm.Int8x16, b int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// ShrqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNS16(a arm.Int16x8, b int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// ShrqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNS32(a arm.Int32x4, b int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// ShrqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNS64(a arm.Int64x2, b int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// ShrqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNU8(a arm.Uint8x16, b int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// ShrqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNU16(a arm.Uint16x8, b int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// ShrqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNU32(a arm.Uint32x4, b int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// ShrqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrqNU64(a arm.Uint64x2, b int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ShrdNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrd_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrdNS64(a int64, b int) int64 {
	return 0
}

// ShrdNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vshrd_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func ShrdNU64(a uint64, b int) uint64 {
	return 0
}

// SliNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// SliNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// SliNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// SliNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// SliNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SliNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SliNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// SliNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsli_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// SliqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// SliqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SliqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SliqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SliqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SliqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SliqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SliqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsliq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SliqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SlidNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vslid_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SlidNS64(a int64, b int64, c int) int64 {
	return 0
}

// SlidNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vslid_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SlidNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// SqaddU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u8'.
// Requires NEON.
func SqaddU8(a arm.Uint8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SqaddU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u16'.
// Requires NEON.
func SqaddU16(a arm.Uint16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SqaddU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u32'.
// Requires NEON.
func SqaddU32(a arm.Uint32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// SqaddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadd_u64'.
// Requires NEON.
func SqaddU64(a arm.Uint64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// SqaddqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u8'.
// Requires NEON.
func SqaddqU8(a arm.Uint8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SqaddqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u16'.
// Requires NEON.
func SqaddqU16(a arm.Uint16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SqaddqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u32'.
// Requires NEON.
func SqaddqU32(a arm.Uint32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SqaddqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddq_u64'.
// Requires NEON.
func SqaddqU64(a arm.Uint64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SqaddbU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddb_u8'.
// Requires NEON.
func SqaddbU8(a uint8, b int8) uint8 {
	return 0
}

// SqaddhU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddh_u16'.
// Requires NEON.
func SqaddhU16(a uint16, b int16) uint16 {
	return 0
}

// SqaddsU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqadds_u32'.
// Requires NEON.
func SqaddsU32(a uint32, b int32) uint32 {
	return 0
}

// SqadddU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqaddd_u64'.
// Requires NEON.
func SqadddU64(a uint64, b int64) uint64 {
	return 0
}

// SqrtF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrt_f32'.
// Requires NEON.
func SqrtF32(a arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// SqrtqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrtq_f32'.
// Requires NEON.
func SqrtqF32(a arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// SqrtF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrt_f64'.
// Requires NEON.
func SqrtF64(a arm.Float64x1) (dst arm.Float64x1) {
	return arm.Float64x1{}
}

// SqrtqF64: ARM NEON intrinsic 
//
// Intrinsic: 'vsqrtq_f64'.
// Requires NEON.
func SqrtqF64(a arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// SraNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// SraNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// SraNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// SraNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// SraNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SraNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SraNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// SraNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsra_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// SraqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// SraqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SraqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SraqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SraqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SraqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SraqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SraqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsraq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SraqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SradNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrad_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SradNS64(a int64, b int64, c int) int64 {
	return 0
}

// SradNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrad_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SradNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// SriNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNS8(a arm.Int8x8, b arm.Int8x8, c int) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// SriNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNS16(a arm.Int16x4, b arm.Int16x4, c int) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// SriNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNS32(a arm.Int32x2, b arm.Int32x2, c int) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// SriNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNS64(a arm.Int64x1, b arm.Int64x1, c int) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// SriNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNU8(a arm.Uint8x8, b arm.Uint8x8, c int) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// SriNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNU16(a arm.Uint16x4, b arm.Uint16x4, c int) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// SriNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNU32(a arm.Uint32x2, b arm.Uint32x2, c int) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// SriNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsri_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriNU64(a arm.Uint64x1, b arm.Uint64x1, c int) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// SriqNS8: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNS8(a arm.Int8x16, b arm.Int8x16, c int) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// SriqNS16: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNS16(a arm.Int16x8, b arm.Int16x8, c int) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// SriqNS32: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNS32(a arm.Int32x4, b arm.Int32x4, c int) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// SriqNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNS64(a arm.Int64x2, b arm.Int64x2, c int) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// SriqNU8: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u8'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNU8(a arm.Uint8x16, b arm.Uint8x16, c int) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// SriqNU16: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u16'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNU16(a arm.Uint16x8, b arm.Uint16x8, c int) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// SriqNU32: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u32'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNU32(a arm.Uint32x4, b arm.Uint32x4, c int) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// SriqNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsriq_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SriqNU64(a arm.Uint64x2, b arm.Uint64x2, c int) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// SridNS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrid_n_s64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SridNS64(a int64, b int64, c int) int64 {
	return 0
}

// SridNU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsrid_n_u64'.
// Requires NEON.
//
// FIXME: Requires compiler support (has immediate)
func SridNU64(a uint64, b uint64, c int) uint64 {
	return 0
}

// St1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1F32(a *float32, b arm.Float32x2)  {

}

// St1F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1F64(a *float64, b arm.Float64x1)  {

}

// St1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1P8(a *arm.Poly8, b arm.Poly8x8)  {

}

// St1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1P16(a *arm.Poly16, b arm.Poly16x4)  {

}

// St1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1S8(a *int8, b arm.Int8x8)  {

}

// St1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1S16(a *int16, b arm.Int16x4)  {

}

// St1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1S32(a *int32, b arm.Int32x2)  {

}

// St1S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1S64(a *int64, b arm.Int64x1)  {

}

// St1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1U8(a *uint8, b arm.Uint8x8)  {

}

// St1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1U16(a *uint16, b arm.Uint16x4)  {

}

// St1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1U32(a *uint32, b arm.Uint32x2)  {

}

// St1U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1U64(a *uint64, b arm.Uint64x1)  {

}

// St1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qF32(a *float32, b arm.Float32x4)  {

}

// St1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qF64(a *float64, b arm.Float64x2)  {

}

// St1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qP8(a *arm.Poly8, b arm.Poly8x16)  {

}

// St1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qP16(a *arm.Poly16, b arm.Poly16x8)  {

}

// St1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qS8(a *int8, b arm.Int8x16)  {

}

// St1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qS16(a *int16, b arm.Int16x8)  {

}

// St1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qS32(a *int32, b arm.Int32x4)  {

}

// St1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qS64(a *int64, b arm.Int64x2)  {

}

// St1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qU8(a *uint8, b arm.Uint8x16)  {

}

// St1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qU16(a *uint16, b arm.Uint16x8)  {

}

// St1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qU32(a *uint32, b arm.Uint32x4)  {

}

// St1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St1qU64(a *uint64, b arm.Uint64x2)  {

}

// St1LaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneF32(a *float32, b arm.Float32x2, lane int)  {

}

// St1LaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneF64(a *float64, b arm.Float64x1, lane int)  {

}

// St1LaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneP8(a *arm.Poly8, b arm.Poly8x8, lane int)  {

}

// St1LaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneP16(a *arm.Poly16, b arm.Poly16x4, lane int)  {

}

// St1LaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneS8(a *int8, b arm.Int8x8, lane int)  {

}

// St1LaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneS16(a *int16, b arm.Int16x4, lane int)  {

}

// St1LaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneS32(a *int32, b arm.Int32x2, lane int)  {

}

// St1LaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneS64(a *int64, b arm.Int64x1, lane int)  {

}

// St1LaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneU8(a *uint8, b arm.Uint8x8, lane int)  {

}

// St1LaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneU16(a *uint16, b arm.Uint16x4, lane int)  {

}

// St1LaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneU32(a *uint32, b arm.Uint32x2, lane int)  {

}

// St1LaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1LaneU64(a *uint64, b arm.Uint64x1, lane int)  {

}

// St1qLaneF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneF32(a *float32, b arm.Float32x4, lane int)  {

}

// St1qLaneF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneF64(a *float64, b arm.Float64x2, lane int)  {

}

// St1qLaneP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneP8(a *arm.Poly8, b arm.Poly8x16, lane int)  {

}

// St1qLaneP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneP16(a *arm.Poly16, b arm.Poly16x8, lane int)  {

}

// St1qLaneS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneS8(a *int8, b arm.Int8x16, lane int)  {

}

// St1qLaneS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneS16(a *int16, b arm.Int16x8, lane int)  {

}

// St1qLaneS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneS32(a *int32, b arm.Int32x4, lane int)  {

}

// St1qLaneS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneS64(a *int64, b arm.Int64x2, lane int)  {

}

// St1qLaneU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneU8(a *uint8, b arm.Uint8x16, lane int)  {

}

// St1qLaneU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneU16(a *uint16, b arm.Uint16x8, lane int)  {

}

// St1qLaneU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneU32(a *uint32, b arm.Uint32x4, lane int)  {

}

// St1qLaneU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst1q_lane_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
//
// FIXME: Requires compiler support (has immediate)
func St1qLaneU64(a *uint64, b arm.Uint64x2, lane int)  {

}

// St2S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2S64(a *int64, val [2]arm.Int64x1)  {

}

// St2U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2U64(a *uint64, val [2]arm.Uint64x1)  {

}

// St2F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2F64(a *float64, val [2]arm.Float64x1)  {

}

// St2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2S8(a *int8, val [2]arm.Int8x8)  {

}

// St2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2P8(a *arm.Poly8, val [2]arm.Poly8x8)  {

}

// St2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2S16(a *int16, val [2]arm.Int16x4)  {

}

// St2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2P16(a *arm.Poly16, val [2]arm.Poly16x4)  {

}

// St2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2S32(a *int32, val [2]arm.Int32x2)  {

}

// St2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2U8(a *uint8, val [2]arm.Uint8x8)  {

}

// St2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2U16(a *uint16, val [2]arm.Uint16x4)  {

}

// St2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2U32(a *uint32, val [2]arm.Uint32x2)  {

}

// St2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2F32(a *float32, val [2]arm.Float32x2)  {

}

// St2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qS8(a *int8, val [2]arm.Int8x16)  {

}

// St2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qP8(a *arm.Poly8, val [2]arm.Poly8x16)  {

}

// St2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qS16(a *int16, val [2]arm.Int16x8)  {

}

// St2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qP16(a *arm.Poly16, val [2]arm.Poly16x8)  {

}

// St2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qS32(a *int32, val [2]arm.Int32x4)  {

}

// St2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qS64(a *int64, val [2]arm.Int64x2)  {

}

// St2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qU8(a *uint8, val [2]arm.Uint8x16)  {

}

// St2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qU16(a *uint16, val [2]arm.Uint16x8)  {

}

// St2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qU32(a *uint32, val [2]arm.Uint32x4)  {

}

// St2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qU64(a *uint64, val [2]arm.Uint64x2)  {

}

// St2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qF32(a *float32, val [2]arm.Float32x4)  {

}

// St2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst2q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St2qF64(a *float64, val [2]arm.Float64x2)  {

}

// St3S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3S64(a *int64, val [3]arm.Int64x1)  {

}

// St3U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3U64(a *uint64, val [3]arm.Uint64x1)  {

}

// St3F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3F64(a *float64, val [3]arm.Float64x1)  {

}

// St3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3S8(a *int8, val [3]arm.Int8x8)  {

}

// St3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3P8(a *arm.Poly8, val [3]arm.Poly8x8)  {

}

// St3S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3S16(a *int16, val [3]arm.Int16x4)  {

}

// St3P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3P16(a *arm.Poly16, val [3]arm.Poly16x4)  {

}

// St3S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3S32(a *int32, val [3]arm.Int32x2)  {

}

// St3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3U8(a *uint8, val [3]arm.Uint8x8)  {

}

// St3U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3U16(a *uint16, val [3]arm.Uint16x4)  {

}

// St3U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3U32(a *uint32, val [3]arm.Uint32x2)  {

}

// St3F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3F32(a *float32, val [3]arm.Float32x2)  {

}

// St3qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qS8(a *int8, val [3]arm.Int8x16)  {

}

// St3qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qP8(a *arm.Poly8, val [3]arm.Poly8x16)  {

}

// St3qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qS16(a *int16, val [3]arm.Int16x8)  {

}

// St3qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qP16(a *arm.Poly16, val [3]arm.Poly16x8)  {

}

// St3qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qS32(a *int32, val [3]arm.Int32x4)  {

}

// St3qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qS64(a *int64, val [3]arm.Int64x2)  {

}

// St3qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qU8(a *uint8, val [3]arm.Uint8x16)  {

}

// St3qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qU16(a *uint16, val [3]arm.Uint16x8)  {

}

// St3qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qU32(a *uint32, val [3]arm.Uint32x4)  {

}

// St3qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qU64(a *uint64, val [3]arm.Uint64x2)  {

}

// St3qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qF32(a *float32, val [3]arm.Float32x4)  {

}

// St3qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst3q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St3qF64(a *float64, val [3]arm.Float64x2)  {

}

// St4S64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4S64(a *int64, val [4]arm.Int64x1)  {

}

// St4U64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4U64(a *uint64, val [4]arm.Uint64x1)  {

}

// St4F64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4F64(a *float64, val [4]arm.Float64x1)  {

}

// St4S8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4S8(a *int8, val [4]arm.Int8x8)  {

}

// St4P8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4P8(a *arm.Poly8, val [4]arm.Poly8x8)  {

}

// St4S16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4S16(a *int16, val [4]arm.Int16x4)  {

}

// St4P16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4P16(a *arm.Poly16, val [4]arm.Poly16x4)  {

}

// St4S32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4S32(a *int32, val [4]arm.Int32x2)  {

}

// St4U8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4U8(a *uint8, val [4]arm.Uint8x8)  {

}

// St4U16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4U16(a *uint16, val [4]arm.Uint16x4)  {

}

// St4U32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4U32(a *uint32, val [4]arm.Uint32x2)  {

}

// St4F32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4F32(a *float32, val [4]arm.Float32x2)  {

}

// St4qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qS8(a *int8, val [4]arm.Int8x16)  {

}

// St4qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_p8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qP8(a *arm.Poly8, val [4]arm.Poly8x16)  {

}

// St4qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qS16(a *int16, val [4]arm.Int16x8)  {

}

// St4qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_p16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qP16(a *arm.Poly16, val [4]arm.Poly16x8)  {

}

// St4qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qS32(a *int32, val [4]arm.Int32x4)  {

}

// St4qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_s64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qS64(a *int64, val [4]arm.Int64x2)  {

}

// St4qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u8'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qU8(a *uint8, val [4]arm.Uint8x16)  {

}

// St4qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u16'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qU16(a *uint16, val [4]arm.Uint16x8)  {

}

// St4qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qU32(a *uint32, val [4]arm.Uint32x4)  {

}

// St4qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_u64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qU64(a *uint64, val [4]arm.Uint64x2)  {

}

// St4qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_f32'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qF32(a *float32, val [4]arm.Float32x4)  {

}

// St4qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vst4q_f64'.
// Requires NEON.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func St4qF64(a *float64, val [4]arm.Float64x2)  {

}

// SubdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubd_s64'.
// Requires NEON.
func SubdS64(a int64, b int64) int64 {
	return 0
}

// SubdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vsubd_u64'.
// Requires NEON.
func SubdU64(a uint64, b uint64) uint64 {
	return 0
}

// Tbx1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx1_s8'.
// Requires NEON.
func Tbx1S8(r arm.Int8x8, tab arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbx1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx1_u8'.
// Requires NEON.
func Tbx1U8(r arm.Uint8x8, tab arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbx1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx1_p8'.
// Requires NEON.
func Tbx1P8(r arm.Poly8x8, tab arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Tbx3S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx3_s8'.
// Requires NEON.
func Tbx3S8(r arm.Int8x8, tab [3]arm.Int8x8, idx arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Tbx3U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx3_u8'.
// Requires NEON.
func Tbx3U8(r arm.Uint8x8, tab [3]arm.Uint8x8, idx arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Tbx3P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtbx3_p8'.
// Requires NEON.
func Tbx3P8(r arm.Poly8x8, tab [3]arm.Poly8x8, idx arm.Uint8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Trn1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_f32'.
// Requires NEON.
func Trn1F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Trn1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_p8'.
// Requires NEON.
func Trn1P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Trn1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_p16'.
// Requires NEON.
func Trn1P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Trn1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_s8'.
// Requires NEON.
func Trn1S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Trn1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_s16'.
// Requires NEON.
func Trn1S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Trn1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_s32'.
// Requires NEON.
func Trn1S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Trn1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_u8'.
// Requires NEON.
func Trn1U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Trn1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_u16'.
// Requires NEON.
func Trn1U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Trn1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1_u32'.
// Requires NEON.
func Trn1U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Trn1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_f32'.
// Requires NEON.
func Trn1qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Trn1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_f64'.
// Requires NEON.
func Trn1qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Trn1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_p8'.
// Requires NEON.
func Trn1qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Trn1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_p16'.
// Requires NEON.
func Trn1qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Trn1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s8'.
// Requires NEON.
func Trn1qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Trn1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s16'.
// Requires NEON.
func Trn1qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Trn1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s32'.
// Requires NEON.
func Trn1qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Trn1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_s64'.
// Requires NEON.
func Trn1qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Trn1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u8'.
// Requires NEON.
func Trn1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Trn1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u16'.
// Requires NEON.
func Trn1qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Trn1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u32'.
// Requires NEON.
func Trn1qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Trn1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn1q_u64'.
// Requires NEON.
func Trn1qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Trn2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_f32'.
// Requires NEON.
func Trn2F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Trn2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_p8'.
// Requires NEON.
func Trn2P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Trn2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_p16'.
// Requires NEON.
func Trn2P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Trn2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_s8'.
// Requires NEON.
func Trn2S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Trn2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_s16'.
// Requires NEON.
func Trn2S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Trn2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_s32'.
// Requires NEON.
func Trn2S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Trn2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_u8'.
// Requires NEON.
func Trn2U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Trn2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_u16'.
// Requires NEON.
func Trn2U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Trn2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2_u32'.
// Requires NEON.
func Trn2U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Trn2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_f32'.
// Requires NEON.
func Trn2qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Trn2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_f64'.
// Requires NEON.
func Trn2qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Trn2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_p8'.
// Requires NEON.
func Trn2qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Trn2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_p16'.
// Requires NEON.
func Trn2qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Trn2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s8'.
// Requires NEON.
func Trn2qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Trn2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s16'.
// Requires NEON.
func Trn2qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Trn2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s32'.
// Requires NEON.
func Trn2qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Trn2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_s64'.
// Requires NEON.
func Trn2qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Trn2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u8'.
// Requires NEON.
func Trn2qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Trn2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u16'.
// Requires NEON.
func Trn2qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Trn2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u32'.
// Requires NEON.
func Trn2qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Trn2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn2q_u64'.
// Requires NEON.
func Trn2qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// TrnF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_f32'.
// Requires NEON.
func TrnF32(a arm.Float32x2, b arm.Float32x2) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// TrnP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_p8'.
// Requires NEON.
func TrnP8(a arm.Poly8x8, b arm.Poly8x8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// TrnP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_p16'.
// Requires NEON.
func TrnP16(a arm.Poly16x4, b arm.Poly16x4) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// TrnS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_s8'.
// Requires NEON.
func TrnS8(a arm.Int8x8, b arm.Int8x8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// TrnS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_s16'.
// Requires NEON.
func TrnS16(a arm.Int16x4, b arm.Int16x4) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// TrnS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_s32'.
// Requires NEON.
func TrnS32(a arm.Int32x2, b arm.Int32x2) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// TrnU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_u8'.
// Requires NEON.
func TrnU8(a arm.Uint8x8, b arm.Uint8x8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// TrnU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_u16'.
// Requires NEON.
func TrnU16(a arm.Uint16x4, b arm.Uint16x4) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// TrnU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrn_u32'.
// Requires NEON.
func TrnU32(a arm.Uint32x2, b arm.Uint32x2) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// TrnqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_f32'.
// Requires NEON.
func TrnqF32(a arm.Float32x4, b arm.Float32x4) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// TrnqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_p8'.
// Requires NEON.
func TrnqP8(a arm.Poly8x16, b arm.Poly8x16) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// TrnqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_p16'.
// Requires NEON.
func TrnqP16(a arm.Poly16x8, b arm.Poly16x8) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// TrnqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_s8'.
// Requires NEON.
func TrnqS8(a arm.Int8x16, b arm.Int8x16) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// TrnqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_s16'.
// Requires NEON.
func TrnqS16(a arm.Int16x8, b arm.Int16x8) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// TrnqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_s32'.
// Requires NEON.
func TrnqS32(a arm.Int32x4, b arm.Int32x4) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// TrnqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_u8'.
// Requires NEON.
func TrnqU8(a arm.Uint8x16, b arm.Uint8x16) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// TrnqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_u16'.
// Requires NEON.
func TrnqU16(a arm.Uint16x8, b arm.Uint16x8) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// TrnqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtrnq_u32'.
// Requires NEON.
func TrnqU32(a arm.Uint32x4, b arm.Uint32x4) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// TstS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s8'.
// Requires NEON.
func TstS8(a arm.Int8x8, b arm.Int8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// TstS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s16'.
// Requires NEON.
func TstS16(a arm.Int16x4, b arm.Int16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// TstS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s32'.
// Requires NEON.
func TstS32(a arm.Int32x2, b arm.Int32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// TstS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_s64'.
// Requires NEON.
func TstS64(a arm.Int64x1, b arm.Int64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// TstU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u8'.
// Requires NEON.
func TstU8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// TstU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u16'.
// Requires NEON.
func TstU16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// TstU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u32'.
// Requires NEON.
func TstU32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// TstU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtst_u64'.
// Requires NEON.
func TstU64(a arm.Uint64x1, b arm.Uint64x1) (dst arm.Uint64x1) {
	return arm.Uint64x1{}
}

// TstqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s8'.
// Requires NEON.
func TstqS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// TstqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s16'.
// Requires NEON.
func TstqS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// TstqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s32'.
// Requires NEON.
func TstqS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// TstqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_s64'.
// Requires NEON.
func TstqS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// TstqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u8'.
// Requires NEON.
func TstqU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// TstqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u16'.
// Requires NEON.
func TstqU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// TstqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u32'.
// Requires NEON.
func TstqU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// TstqU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstq_u64'.
// Requires NEON.
func TstqU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// TstdS64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstd_s64'.
// Requires NEON.
func TstdS64(a int64, b int64) uint64 {
	return 0
}

// TstdU64: ARM NEON intrinsic 
//
// Intrinsic: 'vtstd_u64'.
// Requires NEON.
func TstdU64(a uint64, b uint64) uint64 {
	return 0
}

// UqaddS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s8'.
// Requires NEON.
func UqaddS8(a arm.Int8x8, b arm.Uint8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// UqaddS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s16'.
// Requires NEON.
func UqaddS16(a arm.Int16x4, b arm.Uint16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// UqaddS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s32'.
// Requires NEON.
func UqaddS32(a arm.Int32x2, b arm.Uint32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// UqaddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadd_s64'.
// Requires NEON.
func UqaddS64(a arm.Int64x1, b arm.Uint64x1) (dst arm.Int64x1) {
	return arm.Int64x1{}
}

// UqaddqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s8'.
// Requires NEON.
func UqaddqS8(a arm.Int8x16, b arm.Uint8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// UqaddqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s16'.
// Requires NEON.
func UqaddqS16(a arm.Int16x8, b arm.Uint16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// UqaddqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s32'.
// Requires NEON.
func UqaddqS32(a arm.Int32x4, b arm.Uint32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// UqaddqS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddq_s64'.
// Requires NEON.
func UqaddqS64(a arm.Int64x2, b arm.Uint64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// UqaddbS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddb_s8'.
// Requires NEON.
func UqaddbS8(a int8, b uint8) int8 {
	return 0
}

// UqaddhS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddh_s16'.
// Requires NEON.
func UqaddhS16(a int16, b uint16) int16 {
	return 0
}

// UqaddsS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuqadds_s32'.
// Requires NEON.
func UqaddsS32(a int32, b uint32) int32 {
	return 0
}

// UqadddS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuqaddd_s64'.
// Requires NEON.
func UqadddS64(a int64, b uint64) int64 {
	return 0
}

// Uzp1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_f32'.
// Requires NEON.
func Uzp1F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Uzp1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_p8'.
// Requires NEON.
func Uzp1P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Uzp1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_p16'.
// Requires NEON.
func Uzp1P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Uzp1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_s8'.
// Requires NEON.
func Uzp1S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Uzp1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_s16'.
// Requires NEON.
func Uzp1S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Uzp1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_s32'.
// Requires NEON.
func Uzp1S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Uzp1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_u8'.
// Requires NEON.
func Uzp1U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Uzp1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_u16'.
// Requires NEON.
func Uzp1U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Uzp1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1_u32'.
// Requires NEON.
func Uzp1U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Uzp1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_f32'.
// Requires NEON.
func Uzp1qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Uzp1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_f64'.
// Requires NEON.
func Uzp1qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Uzp1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_p8'.
// Requires NEON.
func Uzp1qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Uzp1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_p16'.
// Requires NEON.
func Uzp1qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Uzp1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s8'.
// Requires NEON.
func Uzp1qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Uzp1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s16'.
// Requires NEON.
func Uzp1qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Uzp1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s32'.
// Requires NEON.
func Uzp1qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Uzp1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_s64'.
// Requires NEON.
func Uzp1qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Uzp1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u8'.
// Requires NEON.
func Uzp1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Uzp1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u16'.
// Requires NEON.
func Uzp1qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Uzp1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u32'.
// Requires NEON.
func Uzp1qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Uzp1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp1q_u64'.
// Requires NEON.
func Uzp1qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Uzp2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_f32'.
// Requires NEON.
func Uzp2F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Uzp2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_p8'.
// Requires NEON.
func Uzp2P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Uzp2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_p16'.
// Requires NEON.
func Uzp2P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Uzp2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_s8'.
// Requires NEON.
func Uzp2S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Uzp2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_s16'.
// Requires NEON.
func Uzp2S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Uzp2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_s32'.
// Requires NEON.
func Uzp2S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Uzp2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_u8'.
// Requires NEON.
func Uzp2U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Uzp2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_u16'.
// Requires NEON.
func Uzp2U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Uzp2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2_u32'.
// Requires NEON.
func Uzp2U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Uzp2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_f32'.
// Requires NEON.
func Uzp2qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Uzp2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_f64'.
// Requires NEON.
func Uzp2qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Uzp2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_p8'.
// Requires NEON.
func Uzp2qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Uzp2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_p16'.
// Requires NEON.
func Uzp2qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Uzp2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s8'.
// Requires NEON.
func Uzp2qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Uzp2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s16'.
// Requires NEON.
func Uzp2qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Uzp2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s32'.
// Requires NEON.
func Uzp2qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Uzp2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_s64'.
// Requires NEON.
func Uzp2qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Uzp2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u8'.
// Requires NEON.
func Uzp2qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Uzp2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u16'.
// Requires NEON.
func Uzp2qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Uzp2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u32'.
// Requires NEON.
func Uzp2qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Uzp2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp2q_u64'.
// Requires NEON.
func Uzp2qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// UzpF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_f32'.
// Requires NEON.
func UzpF32(a arm.Float32x2, b arm.Float32x2) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// UzpP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_p8'.
// Requires NEON.
func UzpP8(a arm.Poly8x8, b arm.Poly8x8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// UzpP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_p16'.
// Requires NEON.
func UzpP16(a arm.Poly16x4, b arm.Poly16x4) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// UzpS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_s8'.
// Requires NEON.
func UzpS8(a arm.Int8x8, b arm.Int8x8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// UzpS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_s16'.
// Requires NEON.
func UzpS16(a arm.Int16x4, b arm.Int16x4) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// UzpS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_s32'.
// Requires NEON.
func UzpS32(a arm.Int32x2, b arm.Int32x2) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// UzpU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_u8'.
// Requires NEON.
func UzpU8(a arm.Uint8x8, b arm.Uint8x8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// UzpU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_u16'.
// Requires NEON.
func UzpU16(a arm.Uint16x4, b arm.Uint16x4) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// UzpU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzp_u32'.
// Requires NEON.
func UzpU32(a arm.Uint32x2, b arm.Uint32x2) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// UzpqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_f32'.
// Requires NEON.
func UzpqF32(a arm.Float32x4, b arm.Float32x4) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// UzpqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_p8'.
// Requires NEON.
func UzpqP8(a arm.Poly8x16, b arm.Poly8x16) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// UzpqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_p16'.
// Requires NEON.
func UzpqP16(a arm.Poly16x8, b arm.Poly16x8) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// UzpqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_s8'.
// Requires NEON.
func UzpqS8(a arm.Int8x16, b arm.Int8x16) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// UzpqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_s16'.
// Requires NEON.
func UzpqS16(a arm.Int16x8, b arm.Int16x8) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// UzpqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_s32'.
// Requires NEON.
func UzpqS32(a arm.Int32x4, b arm.Int32x4) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// UzpqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_u8'.
// Requires NEON.
func UzpqU8(a arm.Uint8x16, b arm.Uint8x16) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// UzpqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_u16'.
// Requires NEON.
func UzpqU16(a arm.Uint16x8, b arm.Uint16x8) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// UzpqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vuzpq_u32'.
// Requires NEON.
func UzpqU32(a arm.Uint32x4, b arm.Uint32x4) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}

// Zip1F32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_f32'.
// Requires NEON.
func Zip1F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Zip1P8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_p8'.
// Requires NEON.
func Zip1P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Zip1P16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_p16'.
// Requires NEON.
func Zip1P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Zip1S8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_s8'.
// Requires NEON.
func Zip1S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Zip1S16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_s16'.
// Requires NEON.
func Zip1S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Zip1S32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_s32'.
// Requires NEON.
func Zip1S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Zip1U8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_u8'.
// Requires NEON.
func Zip1U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Zip1U16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_u16'.
// Requires NEON.
func Zip1U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Zip1U32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1_u32'.
// Requires NEON.
func Zip1U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Zip1qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_f32'.
// Requires NEON.
func Zip1qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Zip1qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_f64'.
// Requires NEON.
func Zip1qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Zip1qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_p8'.
// Requires NEON.
func Zip1qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Zip1qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_p16'.
// Requires NEON.
func Zip1qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Zip1qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s8'.
// Requires NEON.
func Zip1qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Zip1qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s16'.
// Requires NEON.
func Zip1qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Zip1qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s32'.
// Requires NEON.
func Zip1qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Zip1qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_s64'.
// Requires NEON.
func Zip1qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Zip1qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u8'.
// Requires NEON.
func Zip1qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Zip1qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u16'.
// Requires NEON.
func Zip1qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Zip1qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u32'.
// Requires NEON.
func Zip1qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Zip1qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip1q_u64'.
// Requires NEON.
func Zip1qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// Zip2F32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_f32'.
// Requires NEON.
func Zip2F32(a arm.Float32x2, b arm.Float32x2) (dst arm.Float32x2) {
	return arm.Float32x2{}
}

// Zip2P8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_p8'.
// Requires NEON.
func Zip2P8(a arm.Poly8x8, b arm.Poly8x8) (dst arm.Poly8x8) {
	return arm.Poly8x8{}
}

// Zip2P16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_p16'.
// Requires NEON.
func Zip2P16(a arm.Poly16x4, b arm.Poly16x4) (dst arm.Poly16x4) {
	return arm.Poly16x4{}
}

// Zip2S8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_s8'.
// Requires NEON.
func Zip2S8(a arm.Int8x8, b arm.Int8x8) (dst arm.Int8x8) {
	return arm.Int8x8{}
}

// Zip2S16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_s16'.
// Requires NEON.
func Zip2S16(a arm.Int16x4, b arm.Int16x4) (dst arm.Int16x4) {
	return arm.Int16x4{}
}

// Zip2S32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_s32'.
// Requires NEON.
func Zip2S32(a arm.Int32x2, b arm.Int32x2) (dst arm.Int32x2) {
	return arm.Int32x2{}
}

// Zip2U8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_u8'.
// Requires NEON.
func Zip2U8(a arm.Uint8x8, b arm.Uint8x8) (dst arm.Uint8x8) {
	return arm.Uint8x8{}
}

// Zip2U16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_u16'.
// Requires NEON.
func Zip2U16(a arm.Uint16x4, b arm.Uint16x4) (dst arm.Uint16x4) {
	return arm.Uint16x4{}
}

// Zip2U32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2_u32'.
// Requires NEON.
func Zip2U32(a arm.Uint32x2, b arm.Uint32x2) (dst arm.Uint32x2) {
	return arm.Uint32x2{}
}

// Zip2qF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_f32'.
// Requires NEON.
func Zip2qF32(a arm.Float32x4, b arm.Float32x4) (dst arm.Float32x4) {
	return arm.Float32x4{}
}

// Zip2qF64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_f64'.
// Requires NEON.
func Zip2qF64(a arm.Float64x2, b arm.Float64x2) (dst arm.Float64x2) {
	return arm.Float64x2{}
}

// Zip2qP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_p8'.
// Requires NEON.
func Zip2qP8(a arm.Poly8x16, b arm.Poly8x16) (dst arm.Poly8x16) {
	return arm.Poly8x16{}
}

// Zip2qP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_p16'.
// Requires NEON.
func Zip2qP16(a arm.Poly16x8, b arm.Poly16x8) (dst arm.Poly16x8) {
	return arm.Poly16x8{}
}

// Zip2qS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s8'.
// Requires NEON.
func Zip2qS8(a arm.Int8x16, b arm.Int8x16) (dst arm.Int8x16) {
	return arm.Int8x16{}
}

// Zip2qS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s16'.
// Requires NEON.
func Zip2qS16(a arm.Int16x8, b arm.Int16x8) (dst arm.Int16x8) {
	return arm.Int16x8{}
}

// Zip2qS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s32'.
// Requires NEON.
func Zip2qS32(a arm.Int32x4, b arm.Int32x4) (dst arm.Int32x4) {
	return arm.Int32x4{}
}

// Zip2qS64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_s64'.
// Requires NEON.
func Zip2qS64(a arm.Int64x2, b arm.Int64x2) (dst arm.Int64x2) {
	return arm.Int64x2{}
}

// Zip2qU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u8'.
// Requires NEON.
func Zip2qU8(a arm.Uint8x16, b arm.Uint8x16) (dst arm.Uint8x16) {
	return arm.Uint8x16{}
}

// Zip2qU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u16'.
// Requires NEON.
func Zip2qU16(a arm.Uint16x8, b arm.Uint16x8) (dst arm.Uint16x8) {
	return arm.Uint16x8{}
}

// Zip2qU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u32'.
// Requires NEON.
func Zip2qU32(a arm.Uint32x4, b arm.Uint32x4) (dst arm.Uint32x4) {
	return arm.Uint32x4{}
}

// Zip2qU64: ARM NEON intrinsic 
//
// Intrinsic: 'vzip2q_u64'.
// Requires NEON.
func Zip2qU64(a arm.Uint64x2, b arm.Uint64x2) (dst arm.Uint64x2) {
	return arm.Uint64x2{}
}

// ZipF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_f32'.
// Requires NEON.
func ZipF32(a arm.Float32x2, b arm.Float32x2) (dst [2]arm.Float32x2) {
	return [2]arm.Float32x2{}
}

// ZipP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_p8'.
// Requires NEON.
func ZipP8(a arm.Poly8x8, b arm.Poly8x8) (dst [2]arm.Poly8x8) {
	return [2]arm.Poly8x8{}
}

// ZipP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_p16'.
// Requires NEON.
func ZipP16(a arm.Poly16x4, b arm.Poly16x4) (dst [2]arm.Poly16x4) {
	return [2]arm.Poly16x4{}
}

// ZipS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_s8'.
// Requires NEON.
func ZipS8(a arm.Int8x8, b arm.Int8x8) (dst [2]arm.Int8x8) {
	return [2]arm.Int8x8{}
}

// ZipS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_s16'.
// Requires NEON.
func ZipS16(a arm.Int16x4, b arm.Int16x4) (dst [2]arm.Int16x4) {
	return [2]arm.Int16x4{}
}

// ZipS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_s32'.
// Requires NEON.
func ZipS32(a arm.Int32x2, b arm.Int32x2) (dst [2]arm.Int32x2) {
	return [2]arm.Int32x2{}
}

// ZipU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_u8'.
// Requires NEON.
func ZipU8(a arm.Uint8x8, b arm.Uint8x8) (dst [2]arm.Uint8x8) {
	return [2]arm.Uint8x8{}
}

// ZipU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_u16'.
// Requires NEON.
func ZipU16(a arm.Uint16x4, b arm.Uint16x4) (dst [2]arm.Uint16x4) {
	return [2]arm.Uint16x4{}
}

// ZipU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzip_u32'.
// Requires NEON.
func ZipU32(a arm.Uint32x2, b arm.Uint32x2) (dst [2]arm.Uint32x2) {
	return [2]arm.Uint32x2{}
}

// ZipqF32: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_f32'.
// Requires NEON.
func ZipqF32(a arm.Float32x4, b arm.Float32x4) (dst [2]arm.Float32x4) {
	return [2]arm.Float32x4{}
}

// ZipqP8: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_p8'.
// Requires NEON.
func ZipqP8(a arm.Poly8x16, b arm.Poly8x16) (dst [2]arm.Poly8x16) {
	return [2]arm.Poly8x16{}
}

// ZipqP16: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_p16'.
// Requires NEON.
func ZipqP16(a arm.Poly16x8, b arm.Poly16x8) (dst [2]arm.Poly16x8) {
	return [2]arm.Poly16x8{}
}

// ZipqS8: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_s8'.
// Requires NEON.
func ZipqS8(a arm.Int8x16, b arm.Int8x16) (dst [2]arm.Int8x16) {
	return [2]arm.Int8x16{}
}

// ZipqS16: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_s16'.
// Requires NEON.
func ZipqS16(a arm.Int16x8, b arm.Int16x8) (dst [2]arm.Int16x8) {
	return [2]arm.Int16x8{}
}

// ZipqS32: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_s32'.
// Requires NEON.
func ZipqS32(a arm.Int32x4, b arm.Int32x4) (dst [2]arm.Int32x4) {
	return [2]arm.Int32x4{}
}

// ZipqU8: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_u8'.
// Requires NEON.
func ZipqU8(a arm.Uint8x16, b arm.Uint8x16) (dst [2]arm.Uint8x16) {
	return [2]arm.Uint8x16{}
}

// ZipqU16: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_u16'.
// Requires NEON.
func ZipqU16(a arm.Uint16x8, b arm.Uint16x8) (dst [2]arm.Uint16x8) {
	return [2]arm.Uint16x8{}
}

// ZipqU32: ARM NEON intrinsic 
//
// Intrinsic: 'vzipq_u32'.
// Requires NEON.
func ZipqU32(a arm.Uint32x4, b arm.Uint32x4) (dst [2]arm.Uint32x4) {
	return [2]arm.Uint32x4{}
}
