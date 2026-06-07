//go:build !amd64 && !arm64

package drongo

import "github.com/pfcm/drongo/gen"

// TODO: kind of ugly, but seems to big to get inlined and the extra function
// call hurts a bit.
var AddFloat32 = gen.UnrolledScalarAddFloat32

var AbsoluteFloat64 = unrolled32ScalarAbsoluteFloat64
