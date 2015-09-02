// func blendEpi16(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·blendEpi16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PBLENDW

	MOVOU X2, ret+40(FP)
	RET

// func blendPd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·blendPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: BLENDPD

	MOVOU X2, ret+40(FP)
	RET

// func blendPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·blendPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: BLENDPS

	MOVOU X2, ret+40(FP)
	RET

// func blendvEpi8(a [16]byte, b [16]byte, mask [16]byte) [16]byte
TEXT ·blendvEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU mask+32(FP),X2

	// TODO: Code missing - uses instrunction: PBLENDVB

	MOVOU X2, ret+48(FP)
	RET

// func blendvPd(a [2]float64, b [2]float64, mask [2]float64) [2]float64
TEXT ·blendvPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU mask+32(FP),X2

	// TODO: Code missing - uses instrunction: BLENDVPD

	MOVOU X2, ret+48(FP)
	RET

// func blendvPs(a [4]float32, b [4]float32, mask [4]float32) [4]float32
TEXT ·blendvPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVOU mask+32(FP),X2

	// TODO: Code missing - uses instrunction: BLENDVPS

	MOVOU X2, ret+48(FP)
	RET

// func ceilPd(a [2]float64) [2]float64
TEXT ·ceilPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// ROUNDPD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func ceilPs(a [4]float32) [4]float32
TEXT ·ceilPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// ROUNDPS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func ceilSd(a [2]float64, b [2]float64) [2]float64
TEXT ·ceilSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ROUNDSD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func ceilSs(a [4]float32, b [4]float32) [4]float32
TEXT ·ceilSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ROUNDSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpeqEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpeqEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PCMPEQQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpestra(a [16]byte, la int, b [16]byte, lb int, imm8 int) int
TEXT ·cmpestra(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRI

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrc(a [16]byte, la int, b [16]byte, lb int, imm8 int) int
TEXT ·cmpestrc(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRI

	MOVQ $0, ret+56(FP)
	RET

// func cmpestri(a [16]byte, la int, b [16]byte, lb int, imm8 int) int
TEXT ·cmpestri(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRI

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrm(a [16]byte, la int, b [16]byte, lb int, imm8 int) [16]byte
TEXT ·cmpestrm(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRM

	MOVOU X4, ret+56(FP)
	RET

// func cmpestro(a [16]byte, la int, b [16]byte, lb int, imm8 int) int
TEXT ·cmpestro(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRI

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrs(a [16]byte, la int, b [16]byte, lb int, imm8 int) int
TEXT ·cmpestrs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRI

	MOVQ $0, ret+56(FP)
	RET

// func cmpestrz(a [16]byte, la int, b [16]byte, lb int, imm8 int) int
TEXT ·cmpestrz(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ la+16(FP),R9
	MOVOU b+24(FP),X2
	MOVQ lb+40(FP),R11
	MOVQ imm8+48(FP),R12

	// TODO: Code missing - uses instrunction: PCMPESTRI

	MOVQ $0, ret+56(FP)
	RET

// func cmpgtEpi64(a [16]byte, b [16]byte) [16]byte
TEXT ·cmpgtEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PCMPGTQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func cmpistra(a [16]byte, b [16]byte, imm8 int) int
TEXT ·cmpistra(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRI

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrc(a [16]byte, b [16]byte, imm8 int) int
TEXT ·cmpistrc(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRI

	MOVQ $0, ret+40(FP)
	RET

// func cmpistri(a [16]byte, b [16]byte, imm8 int) int
TEXT ·cmpistri(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRI

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrm(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·cmpistrm(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRM

	MOVOU X2, ret+40(FP)
	RET

// func cmpistro(a [16]byte, b [16]byte, imm8 int) int
TEXT ·cmpistro(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRI

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrs(a [16]byte, b [16]byte, imm8 int) int
TEXT ·cmpistrs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRI

	MOVQ $0, ret+40(FP)
	RET

// func cmpistrz(a [16]byte, b [16]byte, imm8 int) int
TEXT ·cmpistrz(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: PCMPISTRI

	MOVQ $0, ret+40(FP)
	RET

// func crc32U16(crc uint32, v uint16) uint32
TEXT ·crc32U16(SB),7,$0
	MOVL crc+0(FP),R8
	MOVW v+4(FP),R9

	// TODO: Code missing - could be:
	// CRC32 R8, R9

	MOVL $0, ret+8(FP)
	RET

// func crc32U32(crc uint32, v uint32) uint32
TEXT ·crc32U32(SB),7,$0
	MOVL crc+0(FP),R8
	MOVL v+4(FP),R9

	// TODO: Code missing - could be:
	// CRC32 R8, R9

	MOVL $0, ret+8(FP)
	RET

// func crc32U64(crc uint64, v uint64) uint64
TEXT ·crc32U64(SB),7,$0
	MOVQ crc+0(FP),R8
	MOVQ v+8(FP),R9

	// TODO: Code missing - could be:
	// CRC32 R8, R9

	MOVQ $0, ret+16(FP)
	RET

// func crc32U8(crc uint32, v uint8) uint32
TEXT ·crc32U8(SB),7,$0
	MOVL crc+0(FP),R8
	MOVB v+4(FP),R9

	// TODO: Code missing - could be:
	// CRC32 R8, R9

	MOVL $0, ret+8(FP)
	RET

// func cvtepi16Epi32(a [16]byte) [16]byte
TEXT ·cvtepi16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVSXWD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepi16Epi64(a [16]byte) [16]byte
TEXT ·cvtepi16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVSXWQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepi32Epi64(a [16]byte) [16]byte
TEXT ·cvtepi32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVSXDQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepi8Epi16(a [16]byte) [16]byte
TEXT ·cvtepi8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVSXBW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepi8Epi32(a [16]byte) [16]byte
TEXT ·cvtepi8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVSXBD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepi8Epi64(a [16]byte) [16]byte
TEXT ·cvtepi8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVSXBQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepu16Epi32(a [16]byte) [16]byte
TEXT ·cvtepu16Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVZXWD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepu16Epi64(a [16]byte) [16]byte
TEXT ·cvtepu16Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVZXWQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepu32Epi64(a [16]byte) [16]byte
TEXT ·cvtepu32Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVZXDQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepu8Epi16(a [16]byte) [16]byte
TEXT ·cvtepu8Epi16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVZXBW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepu8Epi32(a [16]byte) [16]byte
TEXT ·cvtepu8Epi32(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVZXBD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func cvtepu8Epi64(a [16]byte) [16]byte
TEXT ·cvtepu8Epi64(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PMOVZXBQ X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func dpPd(a [2]float64, b [2]float64, imm8 int) [2]float64
TEXT ·dpPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: DPPD

	MOVOU X2, ret+40(FP)
	RET

// func dpPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·dpPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: DPPS

	MOVOU X2, ret+40(FP)
	RET

// func extractEpi32(a [16]byte, imm8 int) int
TEXT ·extractEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing - could be:
	// PEXTRD X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func extractEpi64(a [16]byte, imm8 int) int64
TEXT ·extractEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing - could be:
	// PEXTRQ X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func extractEpi8(a [16]byte, imm8 int) int
TEXT ·extractEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing - could be:
	// PEXTRB X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func extractPs(a [4]float32, imm8 int) int
TEXT ·extractPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ imm8+16(FP),R9

	// TODO: Code missing - could be:
	// EXTRACTPS X0, R9

	MOVQ $0, ret+24(FP)
	RET

// func floorPd(a [2]float64) [2]float64
TEXT ·floorPd(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// ROUNDPD X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func floorPs(a [4]float32) [4]float32
TEXT ·floorPs(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// ROUNDPS X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func floorSd(a [2]float64, b [2]float64) [2]float64
TEXT ·floorSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ROUNDSD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func floorSs(a [4]float32, b [4]float32) [4]float32
TEXT ·floorSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// ROUNDSS X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func insertEpi32(a [16]byte, i int, imm8 int) [16]byte
TEXT ·insertEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	// TODO: Code missing - uses instrunction: PINSRD

	MOVOU X2, ret+32(FP)
	RET

// func insertEpi64(a [16]byte, i int64, imm8 int) [16]byte
TEXT ·insertEpi64(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	// TODO: Code missing - uses instrunction: PINSRQ

	MOVOU X2, ret+32(FP)
	RET

// func insertEpi8(a [16]byte, i int, imm8 int) [16]byte
TEXT ·insertEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ i+16(FP),R9
	MOVQ imm8+24(FP),R10

	// TODO: Code missing - uses instrunction: PINSRB

	MOVOU X2, ret+32(FP)
	RET

// func insertPs(a [4]float32, b [4]float32, imm8 int) [4]float32
TEXT ·insertPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: INSERTPS

	MOVOU X2, ret+40(FP)
	RET

// func maxEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMAXSD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maxEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMAXSB X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maxEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMAXUW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func maxEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·maxEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMAXUD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMINSD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minEpi8(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpi8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMINSB X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minEpu16(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu16(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMINUW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minEpu32(a [16]byte, b [16]byte) [16]byte
TEXT ·minEpu32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMINUD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func minposEpu16(a [16]byte) [16]byte
TEXT ·minposEpu16(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing - could be:
	// PHMINPOSUW X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func mpsadbwEpu8(a [16]byte, b [16]byte, imm8 int) [16]byte
TEXT ·mpsadbwEpu8(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ imm8+32(FP),R10

	// TODO: Code missing - uses instrunction: MPSADBW

	MOVOU X2, ret+40(FP)
	RET

// func mulEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·mulEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMULDQ X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func mulloEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·mulloEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PMULLD X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func packusEpi32(a [16]byte, b [16]byte) [16]byte
TEXT ·packusEpi32(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PACKUSDW X0, X1

	MOVOU X1, ret+32(FP)
	RET

// func roundPd(a [2]float64, rounding int) [2]float64
TEXT ·roundPd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing - could be:
	// ROUNDPD X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func roundPs(a [4]float32, rounding int) [4]float32
TEXT ·roundPs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVQ rounding+16(FP),R9

	// TODO: Code missing - could be:
	// ROUNDPS X0, R9

	MOVOU X1, ret+24(FP)
	RET

// func roundSd(a [2]float64, b [2]float64, rounding int) [2]float64
TEXT ·roundSd(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing - uses instrunction: ROUNDSD

	MOVOU X2, ret+40(FP)
	RET

// func roundSs(a [4]float32, b [4]float32, rounding int) [4]float32
TEXT ·roundSs(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1
	MOVQ rounding+32(FP),R10

	// TODO: Code missing - uses instrunction: ROUNDSS

	MOVOU X2, ret+40(FP)
	RET

// func streamLoadSi128(mem_addr [16]byte) [16]byte
TEXT ·streamLoadSi128(SB),7,$0
	MOVOU mem_addr+0(FP),X0

	// TODO: Code missing - could be:
	// MOVNTDQA X0, X0

	MOVOU X0, ret+16(FP)
	RET

// func testAllOnes(a [16]byte) int
TEXT ·testAllOnes(SB),7,$0
	MOVOU a+0(FP),X0

	// TODO: Code missing

	MOVQ $0, ret+16(FP)
	RET

// func testAllZeros(a [16]byte, mask [16]byte) int
TEXT ·testAllZeros(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1

	// TODO: Code missing - could be:
	// PTEST X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func testMixOnesZeros(a [16]byte, mask [16]byte) int
TEXT ·testMixOnesZeros(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU mask+16(FP),X1

	// TODO: Code missing - could be:
	// PTEST X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func testcSi128(a [16]byte, b [16]byte) int
TEXT ·testcSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PTEST X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func testnzcSi128(a [16]byte, b [16]byte) int
TEXT ·testnzcSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PTEST X0, X1

	MOVQ $0, ret+32(FP)
	RET

// func testzSi128(a [16]byte, b [16]byte) int
TEXT ·testzSi128(SB),7,$0
	MOVOU a+0(FP),X0
	MOVOU b+16(FP),X1

	// TODO: Code missing - could be:
	// PTEST X0, X1

	MOVQ $0, ret+32(FP)
	RET

