//go:build arm64

// package arm64 holds arm64-specific versions of things.
package arm64

import "fmt"

func AddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic(fmt.Errorf("incompatible lengths: %d, %d, %d", len(a), len(b), len(c)))
	}
	addFloat32NEON(a, b, c)
}

// addFloat32NEON is a slightly unrolled elementwise add using NEON vector
// FADDs. It is defined in add_float32.s.
func addFloat32NEON(a, b, c []float32)

func AbsoluteFloat64(a, b []float64) {
	if len(a) != len(b) {
		panic(fmt.Errorf("incompatible lengths: %d, %d", len(a), len(b)))
	}
	absFloat64NEON(a, b)
}

// absFloat64NEON does elementwise absolute values, defined in abs_float64.s
func absFloat64NEON(a, b []float64)
