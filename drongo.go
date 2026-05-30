// package drongo does SIMD things, kind of inspired by Apple's Accelerate
// framework.
package drongo

// AddFloat32 adds a to b elementwise, storing the result in c. Only really
// makes sense to use if you have at least 16 floats to add.
// TODO: can this be in-place?
// func AddFloat32(a, b, c []float32) {
// 	var ax, bx, cx archsimd.Float32x16
// 	for i := 0; i < len(a); i += 16 {
// 		ax = archsimd.LoadFloat32x16Slice(a[i:min(i+16, len(a))])
// 		bx = archsimd.LoadFloat32x16Slice(b[i:min(i+16, len(b))])
// 		cx = ax.Add(bx)
// 		cx.StoreSlice(c[i:])
// 	}
// }
