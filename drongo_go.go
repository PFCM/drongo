//go:build !amd64 && !arm64

package drongo

import "github.com/pfcm/drongo/arch/all"

// TODO: kind of ugly, but seems to big to get inlined and the extra function
// call hurts a bit.
var AddFloat32 = all.AddFloat32

func AbsoluteFloat64(a, b []float64)

func ClipFloat32(in []float32, l, w float32, out []float32)
