// binary cap prints the capabilities of the archsimd library.
package main

import (
	"fmt"
	"maps"
	"simd/archsimd"
	"slices"
)

func main() {
	x := archsimd.X86
	caps := map[string]bool{
		"AVX":              x.AVX(),
		"AVX2":             x.AVX2(),
		"AVX512":           x.AVX512(),
		"AVX512BITALG":     x.AVX512BITALG(),
		"AVX512VAES":       x.AVX512VAES(),
		"AVX512VBMI":       x.AVX512VBMI(),
		"AVX512VBMI2":      x.AVX512VBMI2(),
		"AVX512VNNI":       x.AVX512VNNI(),
		"AVX512VPCLMULQDQ": x.AVX512VPCLMULQDQ(),
		"AVX512VPOPCNTDQ":  x.AVX512VPOPCNTDQ(),
		"AVXAES":           x.AVXAES(),
		"AVXVNNI":          x.AVXVNNI(),
		"FMA":              x.FMA(),
		"VAES":             x.VAES(),
	}
	for _, name := range slices.Sorted(maps.Keys(caps)) {
		capability := caps[name]
		fmt.Printf("%s: %v\n", name, capability)
	}
}
