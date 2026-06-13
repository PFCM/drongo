#include "textflag.h"

// func addFloat32NEON(a, b, c []float32)
// assumes a b and c are all the same size.
TEXT ·addFloat32NEON(SB), NOSPLIT, $0-72
	MOVD a_len+8(FP), R0
	MOVD a_base+0(FP), R1
	MOVD b_base+24(FP), R2
	MOVD c_base+48(FP), R3

	// 4 × 128-bit adds per loop, aka 16 adds.
loop16:
	CMP  $16, R0
	BLT  loop4

	// Load 4 sets of 4 each from a (R1) and b (R2)
	VLD1 (R1), [V0.S4, V1.S4, V2.S4, V3.S4]
	VLD1 (R2), [V4.S4, V5.S4, V6.S4, V7.S4]

	// This is deeply cursed, but the Go assembler doesn't have mnemonics
	// for NEON, so we just have to encode it ourselves.
	// TODO: generate these.
	WORD $0x4e24d400 // fadd v0.4s, v0.4s, v4.4s
	WORD $0x4e25d421 // fadd v1.4s, v1.4s, v5.4s
	WORD $0x4e26d442 // fadd v2.4s, v2.4s, v6.4s
	WORD $0x4e27d463 // fadd v3.4s, v3.4s, v7.4s
	// Something like this would be better, but the assembler generates
	// some kind of integer adds instead of what we're after.
	// VADD V0.S4, V4.S4, V0.S4
	// VADD V1.S4, V5.S4, V1.S4
	// VADD V2.S4, V6.S4, V2.S4
	// VADD V3.S4, V7.S4, V3.S4

	// store results back into c (R3)
	VST1 [V0.S4, V1.S4, V2.S4, V3.S4], (R3)

	// Move along: bump the addresses and decrement the remaining.
	ADD $64, R1
	ADD $64, R2
	ADD $64, R3
	SUB $16, R0
	B   loop16

	// Not unrolled, but still SIMD: just 4 per loop to handle as much of
	// the tail as we can.
loop4:
	CMP  $4, R0
	BLT  loop1

	VLD1 (R1), [V0.S4]
	VLD1 (R2), [V1.S4]
	WORD $0x4e21d400 // fadd v0.4s, v0.4s, v1.4s
	VST1 [V0.S4], (R3)

	ADD $16, R1
	ADD $16, R2
	ADD $16, R3
	SUB $4, R0
	B   loop4

	// Any leftovers the old fashioned way. There's at most 3, we could
	// probably shave a few cycles if we wanted to.
loop1:
	CBZ   R0, done
	FMOVS (R1), F0
	FMOVS (R2), F1
	FADDS F1, F0, F0
	FMOVS F0, (R3)
	ADD   $4, R1
	ADD   $4, R2
	ADD   $4, R3
	SUB   $1, R0
	B     loop1

done:
	RET
