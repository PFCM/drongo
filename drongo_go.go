//go:build !amd64 && !arm64

package drongo

import "github.com/pfcm/drongo/arch/all"

// TODO: kind of ugly, but seems to big to get inlined and the extra function
// call hurts a bit.
var (
	AddFloat32      = all.AddFloat32
	AbsoluteFloat64 = all.AbsoluteFloat64
	ClipFloat32     = all.ClipFloat32
	FillFloat32     = all.FillFloat32
)
