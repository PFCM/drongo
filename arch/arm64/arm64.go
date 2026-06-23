//go:build arm64

// package arm64 holds arm64-specific versions of things.
package arm64

import "fmt"

// addFloat32NEON is a slightly unrolled elementwise add using NEON vector
// FADDs. It is defined in add_float32.s.
func addFloat32NEON(a, b, c []float32)

func AddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic(fmt.Errorf("incompatible lengths: %d, %d, %d", len(a), len(b), len(c)))
	}
	addFloat32NEON(a, b, c)
}

// absFloat64NEON does elementwise absolute values, defined in abs_float64.s
func absFloat64NEON(a, b []float64)

func AbsoluteFloat64(a, b []float64) {
	if len(a) != len(b) {
		panic(fmt.Errorf("incompatible lengths: %d, %d", len(a), len(b)))
	}
	absFloat64NEON(a, b)
}

func clipFloat32(in []float32, lower, upper float32, out []float32) // clip.s

func ClipFloat32(in []float32, lower, upper float32, out []float32) {
	if len(in) != len(out) {
		panic(fmt.Errorf("incompatible lengths: %d, %d", len(in), len(out)))
	}
	clipFloat32(in, lower, upper, out)
}

func fillFloat32(v float32, out []float32)

var FillFloat32 = fillFloat32

func fillFloat32STNP(v float32, out []float32)

var FillFloat32STNP = fillFloat32STNP
