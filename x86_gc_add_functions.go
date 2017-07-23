// +build ignore

addF("github.com/klauspost/intrinsics/x86/sse2", "AddEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256AddEpi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256AddEpi32, types.Types[TM256], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddEpi64, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256AddEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256AddEpi64, types.Types[TM256], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "AddPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "AddSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddsEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddsEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddsEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddsEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddsEpu16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddsEpu16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AddsEpu8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAddsEpu8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/aes", "AesdecSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAesdecSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/aes", "AesdeclastSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAesdeclastSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/aes", "AesencSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAesencSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/aes", "AesenclastSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAesenclastSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/aes", "AesimcSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmAesimcSi128, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AndPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAndPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "AndPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAndPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AndSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAndSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256AndSi256",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256AndSi256, types.Types[TM256], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AndnotPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAndnotPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "AndnotPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAndnotPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AndnotSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAndnotSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AvgEpu16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAvgEpu16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "AvgEpu8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmAvgEpu8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "BroadcastbEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmBroadcastbEpi8, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "BroadcastsdPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmBroadcastsdPd, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "BroadcastssPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmBroadcastssPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "CeilPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCeilPd, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "CeilPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCeilPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "CeilSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCeilSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "CeilSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCeilSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpeqEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpeqEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpeqEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpeqEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256CmpeqEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256CmpeqEpi8, types.Types[TM256], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpeqPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpeqPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpeqPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpeqPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpeqSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpeqSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpeqSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpeqSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpgePd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgePd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpgePs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgePs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpgeSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgeSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpgeSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgeSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpgtEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgtEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpgtEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgtEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpgtPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgtPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpgtPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgtPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpgtSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgtSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpgtSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpgtSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmplePd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmplePd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmplePs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmplePs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpleSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpleSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpleSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpleSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpltEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpltEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpltEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpltEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpltPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpltPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpltPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpltPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpltSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpltSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpltSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpltSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpneqPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpneqPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpneqPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpneqPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpneqSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpneqSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpneqSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpneqSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpngePd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngePd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpngePs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngePs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpngeSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngeSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpngeSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngeSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpngtPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngtPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpngtPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngtPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpngtSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngtSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpngtSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpngtSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpnlePd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnlePd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpnlePs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnlePs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpnleSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnleSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpnleSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnleSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpnltPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnltPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpnltPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnltPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpnltSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnltSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpnltSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpnltSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpordPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpordPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpordPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpordPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpordSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpordSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpordSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpordSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpunordPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpunordPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpunordPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpunordPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CmpunordSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpunordSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "CmpunordSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCmpunordSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepi16Epi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepi16Epi32, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepi16Epi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepi16Epi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepi32Epi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepi32Epi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepi8Epi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepi8Epi16, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepi8Epi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepi8Epi32, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepi8Epi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepi8Epi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepu16Epi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepu16Epi32, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepu16Epi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepu16Epi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepu32Epi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepu32Epi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepu8Epi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepu8Epi16, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepu8Epi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepu8Epi32, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "Cvtepu8Epi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtepu8Epi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CvtpdPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtpdPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CvtpsPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmCvtpsPd, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CvtsdSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCvtsdSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "CvtssSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmCvtssSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "DivPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmDivPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "DivPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmDivPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "DivSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmDivSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "DivSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmDivSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "FloorPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmFloorPd, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "FloorPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmFloorPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "FloorSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmFloorSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "FloorSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmFloorSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "HaddEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHaddEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "HaddEpi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHaddEpi32, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "HaddPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHaddPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "HaddPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHaddPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "HaddsEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHaddsEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "HsubEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHsubEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "HsubEpi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHsubEpi32, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "HsubPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHsubPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "HsubPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHsubPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "HsubsEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmHsubsEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MaxEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMaxEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MaxEpu8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMaxEpu8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MaxPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMaxPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MaxPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMaxPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MaxSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMaxSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MaxSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMaxSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MinEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMinEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MinEpu8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMinEpu8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MinPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMinPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MinPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMinPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MinSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMinSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MinSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMinSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "MinposEpu16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmMinposEpu16, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MoveEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmMoveEpi64, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MoveSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMoveSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MoveSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMoveSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "MovedupPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmMovedupPd, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx", "M256MovedupPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMm256MovedupPd, types.Types[TM256], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "MovehdupPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmMovehdupPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx", "M256MovehdupPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMm256MovehdupPs, types.Types[TM256], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MovehlPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMovehlPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse3", "MoveldupPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmMoveldupPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx", "M256MoveldupPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMm256MoveldupPs, types.Types[TM256], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MovelhPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMovelhPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "MulEpi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulEpi32, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MulPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MulPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MulSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "MulSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MulhiEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulhiEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MulhiEpu16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulhiEpu16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "MulloEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulloEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse4", "MulloEpi32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmMulloEpi32, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "OrPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmOrPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "OrPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmOrPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "OrSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmOrSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256OrSi256",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256OrSi256, types.Types[TM256], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "PacksEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmPacksEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "PackusEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmPackusEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "RcpPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmRcpPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "RcpSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmRcpSs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "RsqrtPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmRsqrtPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "RsqrtSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmRsqrtSs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SadEpu8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSadEpu8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/ssse3", "ShuffleEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmShuffleEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256ShuffleEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256ShuffleEpi8, types.Types[TM256], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SllEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSllEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SllEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSllEpi64, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SqrtPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmSqrtPd, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "SqrtPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmSqrtPs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SqrtSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSqrtSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "SqrtSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue1(ssa.OpAMD64IntrinMmSqrtSs, types.Types[TM128], args[0])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SraEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSraEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SrlEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSrlEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SrlEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSrlEpi64, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubEpi64, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "SubPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubSd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubSd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "SubSs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubSs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubsEpi16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubsEpi16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubsEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubsEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubsEpu16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubsEpu16, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "SubsEpu8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmSubsEpu8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "UnpackhiEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackhiEpi64, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "UnpackhiEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackhiEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "UnpackhiPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackhiPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "UnpackhiPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackhiPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "UnpackloEpi64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackloEpi64, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "UnpackloEpi8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackloEpi8, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "UnpackloPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackloPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "UnpackloPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmUnpackloPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "XorPd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmXorPd, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse", "XorPs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmXorPs, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/sse2", "XorSi128",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMmXorSi128, types.Types[TM128], args[0], args[1])
},sys.AMD64)

addF("github.com/klauspost/intrinsics/x86/avx2", "M256XorSi256",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return s.newValue2(ssa.OpAMD64IntrinMm256XorSi256, types.Types[TM256], args[0], args[1])
},sys.AMD64)

