//go:build !amd64 && !arm64

package drongo

import "github.com/pfcm/drongo/arch/any"

// TODO: kind of ugly, but seems to big to get inlined and the extra function
// call hurts a bit.
var AddFloat32 = any.AddFloat32

func AbsoluteFloat64(a, b []float64)
