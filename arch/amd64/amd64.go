//go:build amd64

// package amd64 is for 64 bit x86 etc.
package amd64

import (
	"simd/archsimd"
)

// AddFloat32 adds a to b elementwise, storing the result in c.
func AddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic("ohno")
	}
	var ax, bx archsimd.Float32x8
	for len(a) > 8 {
		ax = archsimd.LoadFloat32x8Slice(a)
		bx = archsimd.LoadFloat32x8Slice(b)
		ax.Add(bx).StoreSlice(c)

		a = a[8:]
		b = b[8:]
		c = c[8:]
	}
	for i := range a {
		c[i] = a[i] + b[i]
	}
}
