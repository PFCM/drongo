//go:build amd64

package drongo

import (
	"github.com/pfcm/drongo/arch/amd64"
	"github.com/pfcm/drongo/arch/any"
)

var AddFloat32 = amd64.AddFloat32

var AbsoluteFloat64 = any.AbsoluteFloat64
