#include "textflag.h"

// func clipFloat32(in []float32, lower, upper float32, out []float32)
// assumes in and out are the same size. Or at least that out is larger than in.
TEXT ·clipFloat32(SB), NOSPLIT, $32-56
	MOVD R1, R2
	// R0 is already in's base, R1 is the length, and R3 is out
	// F0 and F1 are lower and upper
	VDUP V0.S[0], V30.S4  // Duplicate lower (F0/V0.S[0]) across V30
	VDUP V1.S[0], V31.S4  // Duplicate upper (F1/V1.S[0]) across V31
loop16:	// 16 per loop
	CMP $16, R2
	BLT loop4

	VLD1 (R0), [V0.S4, V1.S4, V2.S4, V3.S4]
	WORD $0b01001110011_00000_111101_11110_00000 // fmax v0.4s, v0.4s, v30.4s
	WORD $0b01001110111_00000_111101_11111_00000 // fmin v0.4s, v0.4s, v31.4s
	WORD $0b01001110011_00001_111101_11110_00001 // fmax v1.4s, v1.4s, v30.4s
	WORD $0b01001110111_00001_111101_11111_00001 // fmin v1.4s, v1.4s, v31.4s
	WORD $0b01001110011_00010_111101_11110_00010 // fmax v2.4s, v2.4s, v30.4s
	WORD $0b01001110111_00010_111101_11111_00010 // fmin v2.4s, v2.4s, v31.4s
	WORD $0b01001110011_00011_111101_11110_00011 // fmax v3.4s, v3.4s, v30.4s
	WORD $0b01001110111_00011_111101_11111_00011 // fmin v3.4s, v3.4s, v31.4s
	
	VST1 [V0.S4, V1.S4, V2.S4, V3.S4], (R3)
	
	ADD $64, R0
	ADD $64, R3
	SUB $16, R2
	B loop16
loop4:	// 4 per loop
	CMP $4, R2
	BLT loop1

	VLD1 (R0), [V0.S4]
	
	WORD $0b01001110011_00000_111101_11110_00000 // fmax v0.4s, v0.4s, v30.4s
	WORD $0b01001110111_00000_111101_11111_00000 // fmin v0.4s, v0.4s, v31.4s

	VST1 [V0.S4], (R3)

	ADD $16, R0
	ADD $16, R3
	SUB $4, R2
	B loop4
loop1:	
	CBZ R2, done

	FMOVS (R0), F2
	FMAXS F30, F2, F2
	FMINS F31, F2, F2
	FMOVS F2, (R3)
	
	ADD $4, R0
	ADD $4, R3
	SUB $1, R2
	B loop1
done:	
	RET
