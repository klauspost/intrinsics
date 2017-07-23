package x86

// MMX
type M64 [8]byte
type M64i [8]byte
type M64Const *M64

// SSE+
//type M128i [16]byte
//type M128 [4]float32
//type M128d [2]float64
//type M128Const *M128
//type M128iConst *M128i
//type M128dConst *M128d

type M128i _m128_
type M128 _m128_
type M128d _m128_
type M128Const _m128_
type M128iConst _m128_
type M128dConst _m128_

// AVX
type M256 [8]float32
type M256i [32]byte
type M256d [4]float64
type M256Const *M256
type M256iConst *M256i
type M256dConst *M256d

// AVX2
type M512 [16]float32
type M512i [64]byte
type M512d [8]float64
type M512Const *M512
type M512iConst *M512i
type M512dConst *M512d

type Mmask8 uint8
type Mmask16 uint16
type Mmask32 uint32
type Mmask64 uint64

// Fixme?
type MMEXPADJENUM int
type MMUPCONVEPI32ENUM int
type MMUPCONVEPI64ENUM int
type MMBROADCAST32ENUM int
type MMBROADCAST64ENUM int
type MMUPCONVPDENUM int
type MMUPCONVPSENUM int
type MMDOWNCONVEPI32ENUM int
type MMDOWNCONVEPI64ENUM int
type MMDOWNCONVPDENUM int
type MMDOWNCONVPSENUM int
type MMMANTISSANORMENUM int
type MMMANTISSASIGNENUM int
type MMPERMENUM int
type MMSWIZZLEENUM int
