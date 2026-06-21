package drongo

import (
	"fmt"
	"math"
	"math/rand/v2"
	"reflect"
	"testing"

	"github.com/pfcm/drongo/arch/all"
)

var benchmarkSizes = []int{16, 65, 259, 1013, 10001, 100000}

func simpleAddFloat32(a, b, c []float32) {
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func lessSimpleAddFloat32(a, b, c []float32) {
	// Let the compile drop bounds checks.
	if len(a) != len(b) || len(a) != len(c) {
		panic("no good")
	}
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func unrolled32ScalarAddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic(fmt.Errorf("incompatible lengths: %d, %d, %d", len(a), len(b), len(c)))
	}
	for len(a) >= 32 {
		c[0] = a[0] + b[0]
		c[1] = a[1] + b[1]
		c[2] = a[2] + b[2]
		c[3] = a[3] + b[3]
		c[4] = a[4] + b[4]
		c[5] = a[5] + b[5]
		c[6] = a[6] + b[6]
		c[7] = a[7] + b[7]

		c[8] = a[8] + b[8]
		c[9] = a[9] + b[9]
		c[10] = a[10] + b[10]
		c[11] = a[11] + b[11]
		c[12] = a[12] + b[12]
		c[13] = a[13] + b[13]
		c[14] = a[14] + b[14]
		c[15] = a[15] + b[15]

		c[16] = a[16] + b[16]
		c[17] = a[17] + b[17]
		c[18] = a[18] + b[18]
		c[19] = a[19] + b[19]
		c[20] = a[20] + b[20]
		c[21] = a[21] + b[21]
		c[22] = a[22] + b[22]
		c[23] = a[23] + b[23]

		c[24] = a[24] + b[24]
		c[25] = a[25] + b[25]
		c[26] = a[26] + b[26]
		c[27] = a[27] + b[27]
		c[28] = a[28] + b[28]
		c[29] = a[29] + b[29]
		c[30] = a[30] + b[30]
		c[31] = a[31] + b[31]

		a = a[32:]
		b = b[32:]
		c = c[32:]
	}
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func randFloat32(n int) []float32 {
	f := make([]float32, n)
	for i := range f {
		f[i] = rand.Float32()
	}
	return f
}

func randFloat64(n int) []float64 {
	f := make([]float64, n)
	for i := range f {
		f[i] = rand.Float64()
	}
	return f
}

var addFloat32s = []struct {
	name string
	f    func(a, b, c []float32)
}{{
	name: "less_simple",
	f:    lessSimpleAddFloat32,
}, {
	name: "least_simple",
	f:    unrolled32ScalarAddFloat32,
}, {
	name: "fallback",
	f:    all.AddFloat32,
}, {
	name: "arch-specific",
	f:    AddFloat32,
}}

func TestAddFloat32(t *testing.T) {
	var (
		n         = 1001
		a, b      = randFloat32(n), randFloat32(n)
		want, got = make([]float32, n), make([]float32, n)
	)
	simpleAddFloat32(a, b, want)

	for _, c := range addFloat32s {
		t.Run(c.name, func(t *testing.T) {
			for i := range got {
				got[i] = 0
			}
			c.f(a, b, got)
			if !reflect.DeepEqual(want, got) {
				for i := range want {
					if want[i] != got[i] {
						t.Fatalf("mismatch, first at %d:\n%f + %f = %f (want %f)", i, a[i], b[i], got[i], want[i])
					}
				}
			}
		})
	}
}

func BenchmarkAddFloat32(b *testing.B) {
	for _, size := range benchmarkSizes {
		var (
			x, y = randFloat32(size), randFloat32(size)
			z    = make([]float32, size)
		)
		for _, c := range addFloat32s {
			b.Run(fmt.Sprintf("%05d/%s", size, c.name), func(b *testing.B) {
				for b.Loop() {
					c.f(x, y, z)
				}
			})
		}
	}
}

func simpleAbsFloat64(a, b []float64) {
	for i := range a {
		b[i] = math.Abs(a[i])
	}
}

func noboundsAbsFloat64(a, b []float64) {
	if len(a) != len(b) {
		panic("nope")
	}
	for i := range a {
		b[i] = math.Abs(a[i])
	}
}

func unrolled32AbsFloat64(a, b []float64) {
	if len(a) != len(b) {
		panic("no thank you")
	}
	for len(a) >= 32 {
		b[0] = math.Abs(a[0])
		b[1] = math.Abs(a[1])
		b[2] = math.Abs(a[2])
		b[3] = math.Abs(a[3])
		b[4] = math.Abs(a[4])
		b[5] = math.Abs(a[5])
		b[6] = math.Abs(a[6])
		b[7] = math.Abs(a[7])

		b[8] = math.Abs(a[8])
		b[9] = math.Abs(a[9])
		b[10] = math.Abs(a[10])
		b[11] = math.Abs(a[11])
		b[12] = math.Abs(a[12])
		b[13] = math.Abs(a[13])
		b[14] = math.Abs(a[14])
		b[15] = math.Abs(a[15])

		b[16] = math.Abs(a[16])
		b[17] = math.Abs(a[17])
		b[18] = math.Abs(a[18])
		b[19] = math.Abs(a[19])
		b[20] = math.Abs(a[20])
		b[21] = math.Abs(a[21])
		b[22] = math.Abs(a[22])
		b[23] = math.Abs(a[23])

		b[24] = math.Abs(a[24])
		b[25] = math.Abs(a[25])
		b[26] = math.Abs(a[26])
		b[27] = math.Abs(a[27])
		b[28] = math.Abs(a[28])
		b[29] = math.Abs(a[29])
		b[30] = math.Abs(a[30])
		b[31] = math.Abs(a[31])

		a = a[32:]
		b = b[32:]
	}
	for i := range a {
		b[i] = math.Abs(a[i])
	}
}

var absoluteFloat64s = []struct {
	name string
	f    func(a, b []float64)
}{{
	name: "simple",
	f:    simpleAbsFloat64,
}, {
	name: "no-bounds-checks",
	f:    noboundsAbsFloat64,
}, {
	name: "unrolled-32",
	f:    unrolled32AbsFloat64,
}, {
	name: "fallback",
	f:    all.AbsoluteFloat64,
}, {
	name: "arch-specific",
	f:    AbsoluteFloat64,
}}

func TestAbsoluteFloat64(t *testing.T) {
	var (
		n         = 999
		in        = randFloat64(n)
		want, got = make([]float64, n), make([]float64, n)
	)
	simpleAbsFloat64(in, want)
	for _, c := range absoluteFloat64s {
		t.Run(c.name, func(t *testing.T) {
			for i := range got {
				got[i] = 0
			}
			c.f(in, got)
			for i := range want {
				if want[i] != got[i] {
					t.Fatalf("mismatch: first difference at %d: %f != %f", i, want[i], got[i])
				}
			}
		})
	}
}

func BenchmarkAbsoluteFloat64(b *testing.B) {
	for _, size := range benchmarkSizes {
		var (
			x = randFloat64(size)
			y = make([]float64, size)
		)
		for _, c := range absoluteFloat64s {
			b.Run(fmt.Sprintf("%d/%s", size, c.name), func(b *testing.B) {
				for b.Loop() {
					c.f(x, y)
				}
			})
		}
	}
}

func simpleClip32(in []float32, lower, upper float32, out []float32) {
	for i, v := range in {
		out[i] = min(max(v, lower), upper)
	}
}

func branchClip32(in []float32, lower, upper float32, out []float32) {
	for i, v := range in {
		switch {
		case v < lower:
			out[i] = lower
		case v > upper:
			out[i] = upper
		default:
			out[i] = v
		}
	}
}

var clip32s = []struct {
	name string
	f    func([]float32, float32, float32, []float32)
}{{
	name: "simple",
	f:    simpleClip32,
}, {
	name: "branch",
	f:    branchClip32,
}, {
	name: "fallback",
	f:    all.ClipFloat32,
}, {
	name: "arch-specific",
	f:    ClipFloat32,
}}

func TestClipFloat32(t *testing.T) {
	var (
		n    = 1013
		in   = randFloat32(n)
		want = randFloat32(n)
	)
	for _, bounds := range [][2]float32{
		{-math.MaxFloat32, math.MaxFloat32},
		{0, 1},
		{-1, 1},
		{0, math.SmallestNonzeroFloat32},
		// TODO: might be cute to test with the smallest non-subnormal
		// float.
	} {
		simpleClip32(in, bounds[0], bounds[1], want)
		for _, c := range clip32s[1:] {
			t.Run(fmt.Sprintf("[%f-%f]/%s", bounds[0], bounds[1], c.name),
				func(t *testing.T) {
					// Fill the output with random numbers,
					// because there's probably cases where
					// the expected output is either exactly
					// the same as the input, or all zeros
					// etc.
					got := randFloat32(n)
					c.f(in, bounds[0], bounds[1], got)
					for i := range in {
						if got[i] != want[i] {
							t.Fatalf("mismatch at %d: %f != %f", i, want[i], got[i])
						}
					}
				})
		}
	}
}

func BenchmarkClipFloat32(b *testing.B) {
	for _, size := range benchmarkSizes {
		var (
			in  = randFloat32(size)
			out = make([]float32, size)
		)
		for _, c := range clip32s {
			b.Run(fmt.Sprintf("%d/%s", size, c.name), func(b *testing.B) {
				for b.Loop() {
					c.f(in, -0.5, 0.5, out)
				}
			})
		}
	}
}
