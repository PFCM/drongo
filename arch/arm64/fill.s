#include "textflag.h"

// func fillFloat32(v float32, out []float32)
TEXT ·fillFloat32(SB), NOSPLIT, $0-32
	MOVD out_len+16(FP), R0
	MOVD out_base+8(FP), R1

	VDUP V0.S[0], V0.S4
	VORR V0.B16, V0.B16, V1.B16
	VORR V0.B16, V0.B16, V2.B16
	VORR V0.B16, V0.B16, V3.B16
loop16:
	CMP $16, R0
	BLT loop4

	VST1 [V0.S4, V1.S4, V2.S4, V3.S4], (R1)
	// we could do something drastic instead like:	
	// WORD $0xac000020   // stnp q0, q0, [x1]      (writes first 32 bytes)
        // WORD $0xac010020   // stnp q0, q0, [x1, #32] (writes next 32 bytes)
	// which could be faster in some cases by skipping some cache stuff.
	// But we probably want cache stuff, and this VST1 is pretty good
	// just about all the time.
	
	ADD $64, R1
	SUB $16, R0
	
	B loop16
loop4:	
	CMP $4, R0
	BLT loop1

	VST1 [V0.S4], (R1)
	
	ADD $16, R1
	SUB $4, R0
	
	B loop4
loop1:
	CBZ R0, done
	
	FMOVS F0, (R1)
	
	ADD $4, R1
	SUB $1, R0
	B loop1

done:
	RET
