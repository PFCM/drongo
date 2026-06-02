//go:build !amd64

package drongo

// TODO: kind of ugly, but seems to big to get inlined and the extra function
// call hurts a bit.
var AddFloat32 = unrolled32ScalarAddFloat32

var AbsoluteFloat64 = unrolled32ScalarAbsoluteFloat64
