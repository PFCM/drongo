//go:build amd64

package drongo

import "simd/archsimd"

// AddFloat32 adds a to b elementwise, storing the result in c. Only really
// makes sense to use if you have at least 16 floats to add.
// TODO: can this be in-place?
func AddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic("ohno")
	}
	// TODO: figure out what's actually available to us.
	var ax, bx, cx archsimd.Float32x8
	for i := 0; i < len(a); i += 16 {
		// TODO: benchmark SlicePart vs just Slice
		ax = archsimd.LoadFloat32x8SlicePart(a[i:min(i+16, len(a))])
		bx = archsimd.LoadFloat32x8SlicePart(b[i:min(i+16, len(b))])
		cx = ax.Add(bx)
		cx.StoreSlicePart(c[i:])
	}
}
