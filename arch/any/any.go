// package any has fallbacks that should build for anything. It's also mostly
// generated code, because unrolling loops is boring. See cmd/gen for the actual
// generation.
package any

import (
	"fmt"
	"math"
)

//go:generate go run ../../cmd/gen -package-name=any -unrolls=16 -export -output=./unrolled.gen.go

func AbsoluteFloat64(a, b []float64) {
	if len(a) != len(b) {
		panic(fmt.Errorf("incompatible lengths: %d, %d", len(a), len(b)))
	}
	for len(a) >= 32 {
		b[0] = math.Abs(a[0])
		b[1] = math.Abs(a[1])
		b[2] = math.Abs(a[2])
		b[3] = math.Abs(a[3])
		b[4] = math.Abs(a[4])
		b[5] = math.Abs(a[5])
		b[6] = math.Abs(a[6])
		b[7] = math.Abs(a[7])

		b[8] = math.Abs(a[8])
		b[9] = math.Abs(a[9])
		b[10] = math.Abs(a[10])
		b[11] = math.Abs(a[11])
		b[12] = math.Abs(a[12])
		b[13] = math.Abs(a[13])
		b[14] = math.Abs(a[14])
		b[15] = math.Abs(a[15])

		b[16] = math.Abs(a[16])
		b[17] = math.Abs(a[17])
		b[18] = math.Abs(a[18])
		b[19] = math.Abs(a[19])
		b[20] = math.Abs(a[20])
		b[21] = math.Abs(a[21])
		b[22] = math.Abs(a[22])
		b[23] = math.Abs(a[23])

		b[24] = math.Abs(a[24])
		b[25] = math.Abs(a[25])
		b[26] = math.Abs(a[26])
		b[27] = math.Abs(a[27])
		b[28] = math.Abs(a[28])
		b[29] = math.Abs(a[29])
		b[30] = math.Abs(a[30])
		b[31] = math.Abs(a[31])

		a = a[32:]
		b = b[32:]
	}

	for len(a) >= 16 {
		b[0] = math.Abs(a[0])
		b[1] = math.Abs(a[1])
		b[2] = math.Abs(a[2])
		b[3] = math.Abs(a[3])
		b[4] = math.Abs(a[4])
		b[5] = math.Abs(a[5])
		b[6] = math.Abs(a[6])
		b[7] = math.Abs(a[7])

		b[8] = math.Abs(a[8])
		b[9] = math.Abs(a[9])
		b[10] = math.Abs(a[10])
		b[11] = math.Abs(a[11])
		b[12] = math.Abs(a[12])
		b[13] = math.Abs(a[13])
		b[14] = math.Abs(a[14])
		b[15] = math.Abs(a[15])

		a = a[16:]
		b = b[16:]
	}

	for len(a) >= 8 {
		b[0] = math.Abs(a[0])
		b[1] = math.Abs(a[1])
		b[2] = math.Abs(a[2])
		b[3] = math.Abs(a[3])
		b[4] = math.Abs(a[4])
		b[5] = math.Abs(a[5])
		b[6] = math.Abs(a[6])
		b[7] = math.Abs(a[7])

		a = a[8:]
		b = b[8:]
	}

	for i := range a {
		b[i] = math.Abs(a[i])
	}
}
