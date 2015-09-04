# DISCLAIMERS
* This is not a finished proposal. 
* This is a brainstorm to see what is feasible, and measure general interest.
* This is not a simple project, it will take a lot of work to implement by people that are better than me.


# WHY INTRINSICS

Go has conquered the server space because of the elegant language, and the high performing implementation. However if you work with images, sound or video, Go doesn't have a big presence, and most rely on external libraries/executables. This is natural, since these types of work have huge gains from SIMD code.

Go has an assembler. However, in terms of *time investment* in my subjective experience, intrinsics are only a small fraction of the time investment of assembler. My personal estimate is that it is 5x more time effective to write intrinsics, and performance is rarely more than 10% below handcrafted assembler.

Furthermore intrinsics are much safer than pure assembler. You do not have to manage registers, so there is no risk of reading uninitialized data. I have an experimentational implementation where we even maintain memory safety, since we remove pointer math.

Also, if intrinsics are done using the existing language, we get all the tools of the language. This will enable thing like ensuring that CPU features are detected correctly, refactoring and all the good things we know and love.


# EXPERIMENTS

I have hacked together to programs that generate function signatures for all existing Intel x86 intrinsics, as well as ARM NEON intrinsics.

To get the generated code, use:

go get github.com/klauspost/intrinsics

If you just want to browse the godoc for the generated code, go to https://godoc.org/github.com/klauspost/intrinsics

To see an example, see https://github.com/klauspost/intrinsics/blob/master/x86/m128_test.go
This is actual working code (if the intrinsics are implemented). TestAddEpi8, TestAddPs, TestMulPs and TestMulAddPs will actually pass if you put in the instruction in the assembler for them.

Function names are simplified, but the aim is to make it simple to make it easy to convert existing intrinsic code.

* `_mm_` prefix is dropped.
* Underscore -> CamelCase.

This means that `_mm_and_si128(...) -> AndSi128(...)`, `_mm_add_epi8() -> AddEpi8(...)`, etc.

All intrinsics are separated into packages based on the CPUID features they require. This has the advantage that you can see your cpu requirements in your imports, and it gives reasonably sized packages.

All instruction that receives a pointer are skipped. Grep source for '// Skipped:' to find these.
Instructions that have an immediate parameter, or returns a value in a parameter pointer are marked with a "FIXME:".
I generate a rather crude stub assembler for each function that loads each parameter from the stack and writes a return value, but no "operation" code is executed. However for basic instructions it will actually work if you uncomment the "proposed" instruction.

MMX has been included for completeness. Maybe that should just be ignored.


# TYPES

The typical x86 assembler uses only a few types. Here are the 128 bit ones:
  * `M128i (__m128i)`, which is a 128 bit register with integer content (8x16, 16x8, 32x4, 64x2 bits packed).
  * `M128 (__m128)`, 128 bit register containing `[4]float32`
  * `M128d (__m128d)`, 128 bit register containing `[2]float64`

These all refer to a REGISTER, they are traditionally not settable directly. In traditional intrinsics you need to call a function to cast between them, set and get values. I think we maybe can do that better in Go.

In my opinion it should be allowed to do `var m M128 = M128(M128i{})`.

Proposed type: `[]M128, []M128i, []M128d`

This references a *safe* piece of memory. These reference the content of an underlying slice. There is no alignment guarantee. Here is an example conversion function:

```Go
// FloatToM128 converts a slice into M128 array.
// The number of elements is len(src) / 4.
// Will be a pointer to the original data, not a copy
func FloatToM128(src []float32) []M128
```
There should be similar for `[]float64 -> []M128d, int8, uint8, int16, etc -> []M128i`.

It behaves as a traditional slice, and looking up an entry gives a "M128x" as you would expect. Since the size is rounded down, only valid memory can be refernced. Writing to the []M128 will write to the original slice.

All of the above is on 128 bit registers (SSEx), there are similar types for 64, 256 and 512 bits.


# ISSUES UNCOVERED

## Immediates

Some intrinsics have "immediate" values, which needs to be compiled into the opcode. Without compiler support there is no real way of achieving this. Example of this: https://godoc.org/github.com/klauspost/intrinsics/x86/sse#Pshufw

## VEX prefixed 3+ register instructions

This is a conceptual issue. A very nice feature of intrisics is that it can switch from SSE codes `(MULPS x1, x2 // x2=x2*x1)` to VEX encoded equivalents `(MULPS x1, x2, x3 // x3=x2*x1)` without needing to rewrite any code. In GCC this is done with compiler flags. IMO we don't want that.

The best solution I have been able to come up with is to duplicate "sse" package into a "sse.vex" package. That will allow enabling VEX encoding by simply changing an import.

However, this only partially solves the problem. The compiler also needs to know that it is ok to use the extra registers (XMM16->XMM31) on AVX-512.

## Destination parameters

There are some intrinsics that return multiple values, and where you supply destination in a parameter. For instance "sse.IdivremEpi32". These should IMO be reworked to return multiple values instead.

## AVX Gather/Scatter

Currently all Gather/Scatter intrinsics are skipped, since they have a pure pointer parameter. Gather/Scatter is very powerful, and not having access to them would be a big annoyance.

## ARM NEON notes

ARM NEON is very crude. I have not yet written ARM intrinsics, so I cannot give any real examples for that. 
Maybe someone here has tried it and can give feedback on it. The current intrinsics are generated based on the GCC "arm_neon.h" file.

## M256 / M512 function names

I considered replacing "M256" -> "V" and "M512" -> "VP" function prefixes in AVX/AVX512 to make the function names shorter.


# COMPILER ADDITIONS

This is of course where the big amount of work lies, and ultimately this will be where the tradeoff between complexity and added features lie. Our aim should be to make the implementation of this as light as possible on the compiler, so we gain the flexibility we need, but don't need to change the compiler every time we need to add/change an intrinsic.

## New type(s)

This is the biggest issue I see, but I currently cannot see any way to avoid having to put in a new type per register size (64, 128, 256, 512). 

On x86 they would represent MM/XMM/YMM/ZMM, on ARM they would represent other registers. This should act as the base types that M128, M128i, etc are created from. They don't need to have a name the user ever sees, so "_vector_register_64_bits_", or something similar that is very unlikely to cause collisions is perfectly fine.

## Inlining

It is a must that intrinsic functions are inlined. The compiler must recognize an intrinsic function and know how to inline them. The usual register assignment optimizations can of course still be applied.

## Specifying intrinsics

There should be a "neat" way to specify intrinsics, so the compiler knows what to do with it, encode registers & immediates. Maybe someone has some input on this?

## Creating aliased slices

I don't know if the proposed aliased slices []M128, etc would require compiler support. It could be done with a simple assembler function that just copies the pointer and adjusts the size, but that seems a bit "hackish".


# OTHER

* Emulation: We could offer pure Go emulation on other platforms. It will be slow, take a lot of work but it is feasible as a long-term goal.
* CPU Feature ensurance: It will be possible to print all used CPU features.
* 32/64 Bits: Intrinsics mostly don't care if they are on a 32 or 64 bit platform. Especially on ARM (which is transitioning now) that could be an advantage.


I would personally *love* to see intrinsics in Go, so that is why I took out the time to research before this brainstorm. Now I would like to hear from you!

* Would this help you? 
* Would you like to see it (eventually)?
* Would you be annoyed that Go developers used time on this? 
* What have I overlooked?
* What am I doing that is silly?
* Is there a showstopper?
* Should it be less/more like C intrinsics?


