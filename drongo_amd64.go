//go:build amd64

package drongo

import (
	"github.com/pfcm/drongo/arch/all"
	"github.com/pfcm/drongo/arch/amd64"
)

var (
	AddFloat32      = amd64.AddFloat32
	AbsoluteFloat64 = all.AbsoluteFloat64
	ClipFloat32     = all.ClipFloat32
	FillFloat32     = all.FillFloat32
)
