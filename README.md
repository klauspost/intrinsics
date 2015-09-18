# intrinsics
Experiment with Go intrinsics (NOT USABLE)

I have hacked together to programs that generate function **signatures** for all existing Intel x86 intrinsics, as well as ARM NEON intrinsics. When used they will not do anything, but
they can be used to test the concept & compilation.

To get the generated code, use:
```
go get github.com/klauspost/intrinsics
```

[Discussion](https://groups.google.com/forum/#!searchin/golang-nuts/intrinsics/golang-nuts/yVOfeHYCIT4/7L_mffslT84J)

All intrinsics are separated into packages based on the CPUID features they require. This has the advantage that you can see your cpu requirements in your imports, and it gives reasonably sized packages.

All instruction that receives a pointer are skipped. Grep sources for '// Skipped:' to find these.

Instructions that have an immediate parameter, or returns a value in a parameter pointer are marked with a "FIXME:".

I generate a rather crude stub assembler for each function that loads each parameter from the stack and writes a return value, but no "operation" code is executed. However for basic instructions it will actually work if you uncomment the "proposed" instruction.


[![GoDoc](https://godoc.org/github.com/klauspost/intrinsics?status.svg)](https://godoc.org/github.com/klauspost/intrinsics)

# links

* [example](https://github.com/klauspost/intrinsics/blob/master/x86/m128_test.go)
* [another example](https://gist.github.com/klauspost/64b36e9904d76d6fc122#file-crc32-intrin-go-L60). [Assembler equivalent](https://go-review.googlesource.com/#/c/14080/7/src/hash/crc32/crc32_amd64.s)
* [inline example](http://play.golang.org/p/dPotG_e2FD), [assembler equivalent](https://github.com/klauspost/compress/blob/master/flate/crc32_amd64.s#L97)
* [types](https://github.com/klauspost/intrinsics/blob/master/x86/types.go)
* [x86 intrinsics godoc](https://godoc.org/github.com/klauspost/intrinsics/x86)
* [ARM NEON intrisics godoc](https://godoc.org/github.com/klauspost/intrinsics/arm/neon)
* [intrinsic data source](https://software.intel.com/sites/landingpage/IntrinsicsGuide/#techs=MMX,SSE,SSE2,SSE3,SSSE3,SSE4_1,SSE4_2,AVX,AVX2,FMA,KNC,SVML,Other&avx512techs=AVX512F,AVX512BW,AVX512CD,AVX512DQ,AVX512ER,AVX512IFMA52,AVX512PF,AVX512VBMI)

Copyright 2015 Klaus Post.

Released under MIT license. See [LICENCE](https://raw.githubusercontent.com/klauspost/intrinsics/master/LICENSE) for more information.
