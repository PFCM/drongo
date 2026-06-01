package drongo

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"
)

func simpleAddFloat32(a, b, c []float32) {
	// TODO: pre-check so the compiler has a shot at eliding bounds checks.
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func lessSimpleAddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic("no good")
	}
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func leastSimpleAddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic("not the same")
	}
	for len(a) > 32 {
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

func TestAddFloat32(t *testing.T) {
	var (
		n         = 1001
		a, b      = randFloat32(n), randFloat32(n)
		want, got = make([]float32, n), make([]float32, n)
	)
	simpleAddFloat32(a, b, want)
	leastSimpleAddFloat32(a, b, got)
	// AddFloat32(a, b, got)

	if !reflect.DeepEqual(want, got) {
		t.Fatal("mismatch")
	}
}

func BenchmarkAddFloat32(b *testing.B) {
	for _, size := range []int{16, 65, 259, 1013, 10001, 100000} {
		var (
			x, y = randFloat32(size), randFloat32(size)
			z    = make([]float32, size)
		)
		for _, c := range []struct {
			name string
			f    func(a, b, c []float32)
		}{{
			name: "simple",
			f:    simpleAddFloat32,
		}, {
			name: "less_simple",
			f:    lessSimpleAddFloat32,
		}, {
			name: "least_simple",
			f:    leastSimpleAddFloat32,
		}} {
			b.Run(fmt.Sprintf("%05d/%s", size, c.name), func(b *testing.B) {
				for b.Loop() {
					c.f(x, y, z)
				}
			})
		}
	}
}
