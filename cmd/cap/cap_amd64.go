//go:build amd64

package main

import "simd/archsimd"

func getCaps() map[string]bool {
	x := archsimd.X86
	return map[string]bool{
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
}
