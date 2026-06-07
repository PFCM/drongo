//go:build arm64

package drongo

import "fmt"

func AddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic(fmt.Errorf("incompatible lengths: %d, %d, %d", len(a), len(b), len(c)))
	}
	addFloat32NEON(a, b, c)
}

// drongo_arm64.s
func addFloat32NEON(a, b, c []float32)

var AbsoluteFloat64 = unrolled32ScalarAbsoluteFloat64
