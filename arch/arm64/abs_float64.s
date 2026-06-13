#include "textflag.h"

// func absFloat64NEON(a, b []float64)
// assumes a and b are the same size.
TEXT ·absFloat64NEON(SB), NOSPLIT, $0-48
	MOVD a_len+8(FP), R0
	MOVD a_base+0(FP), R1
	MOVD b_base+24(FP), R2

loop16:	// 8 per loop
	CMP $8, R0
	BLT loop4
	
	VLD1 (R1), [V0.D2, V1.D2, V2.D2, V3.D2]

	WORD $0b0100111011100000101110_00000_00000 // abs v0.2d, v0.2d
	WORD $0b0100111011100000101110_00001_00001 // abs v1.2d, v1.2d
	WORD $0b0100111011100000101110_00010_00010 // abs v2.2d, v2.2d
	WORD $0b0100111011100000101110_00011_00011 // abs v3.2d, v3.2d

	VST1 [V0.D2, V1.D2, V2.D2, V3.D2], (R2)
	
	ADD $64, R1
	ADD $64, R2
	SUB $8, R0
	B loop16

loop4:	// 4 per loop.
	CMP $4, R0
	BLT loop1

	VLD1 (R1), [V0.D2, V1.D2]
	
	WORD $0b0100111011100000101110_00000_00000 // abs v0.2d, v0.2d
	WORD $0b0100111011100000101110_00001_00001 // abs v1.2d, v1.2d

	VST1 [V0.D2, V1.D2], (R2)
	
	ADD $32, R1
	ADD $32, R2
	SUB $4, R0
	B loop4

loop1:	// 1 per loop
	CBZ R0, done
	FMOVD (R1), F0
	FABSD F0, F0
	FMOVD F0, (R2)
	ADD $8, R1
	ADD $8, R2
	SUB $1, R0
	B loop1
done:	
	RET
